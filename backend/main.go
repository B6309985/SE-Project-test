package main

import (
	"github.com/B6309985/sa-65-example/controller"

	"github.com/B6309985/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)
 
func main() {
 
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
 
	// Treatment Routes
	
	r.GET("/treatments", controller.ListTreatmentShow)
	r.GET("/treatments/:id", controller.GetTreatment)
	r.POST("/treatments", controller.CreateTreatment) 
	r.PATCH("/treatments/:id", controller.UpdateTreatment)  
	r.DELETE("/treatments/:id", controller.DeleteTreatment)

	// Type of number of treatment Routes
	
	r.GET("/type_of_number_of_treatments", controller.List_Type_of_number_of_treatment)
	r.GET("/type_of_number_of_treatments/:id", controller.Get_Type_of_number_of_treatment)
	r.POST("/type_of_number_of_treatments", controller.Create_Type_of_number_of_treatment)
	r.PATCH("/type_of_number_of_treatments", controller.Update_Type_of_number_of_treatment)
	r.DELETE("/type_of_number_of_treatments/:id", controller.Delete_Type_of_number_of_treatment)
	
	// Type of treatment Routes
 
	r.GET("/type_of_treatments", controller.List_Type_of_treatment)
	r.GET("/type_of_treatments/:id", controller.Get_Type_of_treatment)
	r.POST("/type_of_treatments", controller.Create_Type_of_treatment)
	r.PATCH("/type_of_treatments", controller.Update_Type_of_treatment)
	r.DELETE("/type_of_treatments/:id", controller.Delete_Type_of_treatment)

	

	// Treatment Plan Routes
	
	r.GET("/treatment_plans", controller.ListTreatment_plan_show)
	r.GET("/treatment_plans/:id", controller.GetTreatment_plan)
	r.POST("/treatment_plans", controller.CreateTreatment_plan) 
	r.PATCH("/treatment_plans", controller.UpdateTreatment_plan)  
	r.DELETE("/treatment_plans/:id", controller.DeleteTreatment_plan)   
	

	r.GET("/dentists", controller.ListDentist)
	r.GET("/dentists/:id", controller.GetDentist) 
	r.POST("/dentists", controller.CreateDentist)
	r.PATCH("/dentists", controller.UpdateDentist)
	r.DELETE("/dentists/:id", controller.DeleteDentist)

	r.GET("/patients", controller.ListPatient)
	r.GET("/patients/:id", controller.GetPatient)
	r.POST("/patients", controller.CreatePatient)
	r.PATCH("/patients", controller.UpdatePatient)
	r.DELETE("/patients/:id", controller.DeletePatient)


	// Run the server

	r.Run() 

}
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH,DELETE")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return 

		}

		c.Next()

	}

}
