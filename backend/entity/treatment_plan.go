package entity

import (
	"time"

	"gorm.io/gorm"
)

type Treatment_plan struct { 
	gorm.Model

	DentistID *uint
	Dentist   Dentist

	PatientID *uint  
	Patient   Patient
	
	Order_of_treatment int 
 
	Type_Of_TreatmentID *uint
	Type_Of_Treatment    Type_of_treatment

	Number_of_treatment int

	Type_Of_Number_Of_TreatmentID *uint
	Type_Of_Number_Of_Treatment Type_of_number_of_treatment

	Treatment_detail string

	Treatment_explain string

	Treatment_time time.Time
}
