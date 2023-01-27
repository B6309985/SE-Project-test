package controller

import (
	"net/http"

	"github.com/gin-gonic/gin" 
	"github.com/B6309985/sa-65-example/entity" 
)

// POST /users
func CreateTreatment(c *gin.Context) {
	var treatments entity.Treatment
	var dentists entity.Dentist
	var patients entity.Patient
	var type_of_treatments entity.Type_of_treatment
	var type_of_number_of_treatments entity.Type_of_number_of_treatment


	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร branch
	if err := c.ShouldBindJSON(&treatments); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา admin ด้วย id 
	if tx := entity.DB().Where("id = ?", treatments.DentistID).First(&dentists); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dentist not found"})
		return
	}

	// 10: ค้นหา academy ด้วย id
	if tx := entity.DB().Where("id = ?", treatments.PatientID).First(&patients); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// 11: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", treatments.Type_Of_TreatmentID).First(&type_of_treatments); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", treatments.Type_Of_Number_Of_TreatmentID).First(&type_of_number_of_treatments); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	// 12: สร้าง branch
	treatment := entity.Treatment{
		Dentist:                dentists,           
		Patient:                patients,           
		Number_of_cavities:     treatments.Number_of_cavities,
		Number_of_swollen_gums: treatments.Number_of_swollen_gums,
		Other_teeth_problems:   treatments.Other_teeth_problems,
		Type_Of_Treatment:      type_of_treatments, 
		Number_of_treatment: treatments.Number_of_treatment,
		Type_Of_Number_Of_Treatment: type_of_number_of_treatments,
		Treatment_detail:       treatments.Treatment_detail,
		Treatment_time: treatments.Treatment_time,
		Treatment_code:         treatments.Treatment_code,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": treatment})
}
 
// GET /user/:id
func GetTreatment(c *gin.Context) {
	var treatments entity.Treatment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM treatments WHERE id = ?", id).Scan(&treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatments})
}

func ListTreatments(c *gin.Context) { 
 
	var treatments []entity.Treatment

	if err := entity.DB().Raw("SELECT * FROM treatments").Scan(&treatments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": treatments}) 

}

// //////////////////////////////////////////////////////////////////////////maybe change Listshow/////////////////////////////////////////
// GET /users
func ListTreatmentShow(c *gin.Context) {
	// var bp []entity.Behavior_Point
	result := []map[string]interface{}{}
	entity.DB().Table("treatments").
		Select("treatments.id, dentists.dentist_name, patients.patient_name, type_of_treatments.type_of_treatment_name, treatments.number_of_cavities, treatments.number_of_swollen_gums, treatments.number_of_treatment, treatments.other_teeth_problems, treatments.treatment_detail, treatments.treatment_time, treatments.treatment_code").
		Joins("left join dentists on dentists.id = treatments.dentist_id").
		Joins("left join patients on patients.id = treatments.patient_id").
		Joins("left join type_of_treatments on type_of_treatments.id = treatments.type_of_treatment_id").
		Find(&result)
 
	c.JSON(http.StatusOK, gin.H{"data": result})

}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// DELETE /users/:id
// 
func DeleteTreatment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM treatments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /users
func UpdateTreatment(c *gin.Context) {
	// var treatments entity.Treatment
	// if err := c.ShouldBindJSON(&treatments); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// if tx := entity.DB().Where("id = ?", treatments.ID).First(&treatments); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "treatment not found"})
	// 	return
	// }
	// if err := entity.DB().Save(&treatments).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": treatments})
	var treatments entity.Treatment
	id := c.Param("id")

	var dentists entity.Dentist
	var patients entity.Patient
	var type_of_treatments entity.Type_of_treatment
	var type_of_number_of_treatments entity.Type_of_number_of_treatment

	if err := c.ShouldBindJSON(&treatments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", treatments.Type_Of_TreatmentID).First(&type_of_treatments); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", treatments.Type_Of_Number_Of_TreatmentID).First(&type_of_number_of_treatments); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inventory not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", treatments.DentistID).First(&dentists); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discount_type not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", treatments.PatientID).First(&patients); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "discount_type not found"})
		return
	}
	treatment := entity.Treatment{
		Dentist:                dentists,           
		Patient:                patients,           
		Number_of_cavities:     treatments.Number_of_cavities,
		Number_of_swollen_gums: treatments.Number_of_swollen_gums,
		Other_teeth_problems:   treatments.Other_teeth_problems,
		Type_Of_Treatment:      type_of_treatments, 
		Number_of_treatment: treatments.Number_of_treatment,
		Type_Of_Number_Of_Treatment: type_of_number_of_treatments,
		Treatment_detail:       treatments.Treatment_detail,
		Treatment_time: treatments.Treatment_time,
		Treatment_code:         treatments.Treatment_code,
	}
	if err := entity.DB().Where("id = ?", id).Updates(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})
}
