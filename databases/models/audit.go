package models

import "gorm.io/gorm"

type AuditLog struct {
	gorm.Model
	Action      string
	Resource    string
	ResourceID  string
	Description string
}