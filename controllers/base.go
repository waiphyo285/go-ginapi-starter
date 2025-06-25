package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController[T any] struct {
	DB    *gorm.DB
	Model T
}

func NewBaseController[T any](db *gorm.DB, model T) *BaseController[T] {
	return &BaseController[T]{DB: db, Model: model}
}

func (ctl *BaseController[T]) List(c *gin.Context) {
	var items []T
	if err := ctl.DB.Find(&items).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusNotFound,
			"data": "Failed to list records",
		})
		return
	}
	c.Set("response", items)
}

func (ctl *BaseController[T]) Get(c *gin.Context) {
	var item T
	if err := ctl.DB.First(&item, c.Param("id")).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusNotFound,
			"data": "Not found",
		})
		return
	}
	c.Set("response", item)
}

func (ctl *BaseController[T]) Create(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": err.Error(),
		})
		return
	}
	if err := ctl.DB.Create(&item).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusInternalServerError,
			"data": "Failed to create",
		})
		return
	}
	c.Set("response", item)
	c.Set("status", http.StatusCreated)
}

func (ctl *BaseController[T]) Update(c *gin.Context) {
	var item T
	if err := ctl.DB.First(&item, c.Param("id")).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusNotFound,
			"data": "Not found",
		})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": err.Error(),
		})
		return
	}
	if err := ctl.DB.Updates(&item).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusInternalServerError,
			"data": "Failed to update",
		})
		return
	}
	c.Set("response", item)
}

func (ctl *BaseController[T]) Delete(c *gin.Context) {
	var item T
	if err := ctl.DB.First(&item, c.Param("id")).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusNotFound,
			"data": "Not found",
		})
		return
	}
	if err := ctl.DB.Delete(&item).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusInternalServerError,
			"data": "Failed to delete",
		})
		return
	}
	c.Set("response", item)
}
