package entity

import (
	//"time"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(&Treatment{}, &Treatment_plan{})

	db = database

	moyong := Dentist{
		Dentist_name: "moyong",
	}
	db.Model(&Dentist{}).Create(&moyong)

	somchai := Patient{
		Patient_name: "somchai",
	}
	db.Model(&Patient{}).Create(&somchai)

	tonfun := Type_of_treatment{
		Type_of_treatment_name: "tonfun",
	}
	db.Model(&Type_of_treatment{}).Create(&tonfun)

	Root_canal_treatment_anterior_teeth := Type_of_treatment{
		Type_of_treatment_name: "รักษารากฟัน(ฟันหน้า)",
		Price:                  5000,
	}
	db.Model(&Type_of_treatment{}).Create(&Root_canal_treatment_anterior_teeth)

	teeth := Type_of_number_of_treatment{
		Type_of_number_of_treatment_name: "ซี่",
	}
	db.Model(&Type_of_number_of_treatment{}).Create(&teeth)

	film := Type_of_number_of_treatment{
		Type_of_number_of_treatment_name: "ฟิล์ม",
	}
	db.Model(&Type_of_number_of_treatment{}).Create(&film)


	  ////////////////////////////////////////////

	// db.Model(&Treatment{}).Create(&Treatment{
	// 	Dentist:           moyong,
	// 	Patient:           somchai,
	// 	Number_of_cavities:     2,
	// 	Number_of_swollen_gums: 10,
	// 	Other_teeth_problems:   "nothing",
	// 	Type_Of_Treatment: tonfun,

		
		
	// 	Treatment_detail:       "tonfun pai 2 see",
	// 	Treatment_code:         "64rgrfweoij",
	// })

	// db.Model(&Treatment{}).Create(&Treatment_plan{
	// 	Dentist:           moyong,
	// 	Patient:           somchai,
	// 	Order_of_treatment: 1,
	// 	Type_Of_Treatment: tonfun,
	// 	Treatment: "tonfun1",
	// 	Treatment_Time: time.Date(),

	// })

	
	

	
}
