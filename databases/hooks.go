package databases

import (
	"fmt"

	"gorm.io/gorm"
	"neohub.asia/mod/databases/models"
)

func RegisterHooks(db *gorm.DB) {
	db.Callback().Create().After("gorm:create").Register("audit_after_create", func(db *gorm.DB) {
		// Check if the model is a Book
		if book, ok := db.Statement.Dest.(*models.Book); ok {
			log := models.AuditLog{
				Action:      "CREATE",
				Resource:    "Book",
				ResourceID:  fmt.Sprintf("%d", book.ID),
				Description: "Book created",
			}
			db.Session(&gorm.Session{NewDB: true}).Create(&log)
		}
	})
}
