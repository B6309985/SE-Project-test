package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/B6309985/sa-65-example/entity"
)

// POST /users
func CreatePatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})

}

// GET /BehaviorType/:id 

func GetPatient(c *gin.Context) { 
	var patient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})

} 

// GET /BehaviorTypes

func ListPatient(c *gin.Context) {
	var patient []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM patients").Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})

}

// DELETE /BehaviorTypes/:id

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patients WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type of treatment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

} 

// PATCH /BehaviorTypes

func UpdatePatient(c *gin.Context) {
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
