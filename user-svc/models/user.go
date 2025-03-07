package models

import _ "gorm.io/gorm"

type UserType string

const (
	Individual UserType = "individual"
	Company    UserType = "company"
)

type User struct {
	ID        uint     `gorm:"primaryKey"`
	Type      UserType `json:"type"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Address   string   `json:"address,omitempty"`
	CreatedAt string   `gorm:"autoCreateTime"`
}

type Provider struct {
	ID             uint     `gorm:"primaryKey"`
	Type           UserType `json:"type"`
	CompanyName    *string  `json:"company_name,omitempty"`
	BusinessTaxNo  *string  `json:"business_tax_no,omitempty"`
	Representative *string  `json:"representative,omitempty"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Address        *string  `json:"address,omitempty"`
	CreatedAt      string   `gorm:"autoCreateTime"`
}
