package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/B6309985/sa-65-example/entity"
)

// POST /users
func CreateDentist(c *gin.Context) {
	var dentist entity.Dentist
	if err := c.ShouldBindJSON(&dentist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&dentist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dentist})

}

// GET /BehaviorType/:id 

func GetDentist(c *gin.Context) { 
	var dentist entity.Dentist
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM dentists WHERE id = ?", id).Scan(&dentist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dentist})

} 

// GET /BehaviorTypes

func ListDentist(c *gin.Context) {
	var dentist []entity.Dentist
	if err := entity.DB().Raw("SELECT * FROM dentists").Scan(&dentist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dentist})

}

// DELETE /BehaviorTypes/:id

func DeleteDentist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM dentists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type of treatment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

} 

// PATCH /BehaviorTypes

func UpdateDentist(c *gin.Context) {
	var dentist entity.Type_of_treatment
	if err := c.ShouldBindJSON(&dentist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", dentist.ID).First(&dentist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type of treatment not found"})
		return
	}

	if err := entity.DB().Save(&dentist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dentist})

}
