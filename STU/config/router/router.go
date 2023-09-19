package router

import (
	"STU/app/controllers/studentControllers"
	"STU/app/controllers/teacherControllers"
	"STU/app/controllers/userControllers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/user/login", userControllers.Login)
		api.POST("/user/reg", userControllers.Register)

		contact := api.Group("/teacher/course")
		{
			contact.GET("", teacherControllers.GetCourse)
			contact.POST("", teacherControllers.CreateCourse)
			contact.PUT("", teacherControllers.UpdateCourse)
			contact.DELETE("", teacherControllers.DeleteCourse)
		}

		stucontact := api.Group("/student")
		{

			stucontact.GET("/optional-course", studentControllers.GetList)
			stucontact.GET("/select-course", studentControllers.GetTakenList)
			stucontact.POST("/course", studentControllers.ChoiceClass)
			stucontact.DELETE("/course", studentControllers.DeleteClass)
		}

	}
}
