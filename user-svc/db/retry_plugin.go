package db

import (
	"context"
	"log"
	"math"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Configurable retry settings
const (
	maxRetries     = 5                      // Maximum retry attempts
	initialBackoff = 100 * time.Millisecond // Initial backoff duration
	backoffFactor  = 2.0                    // Exponential backoff multiplier
)

// RetryPlugin is a custom GORM plugin for automatic retries
type RetryPlugin struct{}

// Name returns the name of the plugin
func (p *RetryPlugin) Name() string {
	return "RetryPlugin"
}

// Initialize registers the retry logic into GORM callbacks
func (p *RetryPlugin) Initialize(db *gorm.DB) error {
	// Register retry hooks for common DB operations
	_ = db.Callback().Query().Before("gorm:query").Register("retry_before_query", p.retryHook)
	_ = db.Callback().Row().Before("gorm:row").Register("retry_before_row", p.retryHook)
	_ = db.Callback().Create().Before("gorm:create").Register("retry_before_create", p.retryHook)
	_ = db.Callback().Update().Before("gorm:update").Register("retry_before_update", p.retryHook)
	_ = db.Callback().Delete().Before("gorm:delete").Register("retry_before_delete", p.retryHook)

	return nil
}

// retryHook applies automatic retries for transient errors
func (p *RetryPlugin) retryHook(db *gorm.DB) {
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Execute the query and ignore the result
		_, err = db.Statement.ConnPool.ExecContext(context.Background(), db.Statement.SQL.String(), db.Statement.Vars...)
		if err == nil {
			return // Success
		}

		// Check if error is retryable
		if !isRetryableError(err) {
			db.AddError(err)
			return
		}

		// Apply exponential backoff before retrying
		backoffDuration := time.Duration(float64(initialBackoff) * math.Pow(backoffFactor, float64(attempt)))
		log.Printf("Retrying DB operation (%d/%d) after error: %v", attempt, maxRetries, err)
		time.Sleep(backoffDuration)
	}

	// If retries exhausted, log and return error
	db.AddError(err)
}

// isRetryableError checks if an error is retryable
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}
	errStr := strings.ToLower(err.Error())
	retryableErrors := []string{
		"deadlock detected",
		"could not serialize access due to concurrent update",
		"the database system is in recovery mode",
		"network error",
		"connection refused",
		"connection reset by peer",
	}

	for _, retryErr := range retryableErrors {
		if strings.Contains(errStr, retryErr) {
			return true
		}
	}
	return false
}
