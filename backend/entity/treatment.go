package entity

import (
	"time"

	"gorm.io/gorm"
)

type Dentist struct {
	gorm.Model
    
	Dentist_name string

	Treatment        []Treatment  `gorm:"foreignkey:DentistID"`
	
}


type Type_of_treatment struct {
	gorm.Model
	Type_of_treatment_name       string
	Price 						 int
	Treatment        []Treatment  `gorm:"foreignkey:Type_Of_TreatmentID"`
} 

type Type_of_number_of_treatment struct {
	gorm.Model
	Type_of_number_of_treatment_name       string
	Treatment        []Treatment  `gorm:"foreignkey:Type_Of_Number_Of_TreatmentID"`
} 

type Treatment struct {
	gorm.Model		
 
	DentistID *uint   
	Dentist   Dentist

	PatientID *uint  
	Patient   Patient

	Number_of_cavities int    
 
	Number_of_swollen_gums int  

	Other_teeth_problems string

	Type_Of_TreatmentID *uint 
	Type_Of_Treatment	Type_of_treatment
 
	Number_of_treatment int

	Type_Of_Number_Of_TreatmentID *uint
	Type_Of_Number_Of_Treatment Type_of_number_of_treatment
		
	Treatment_detail string

	Treatment_time time.Time

	Treatment_code string
}
 