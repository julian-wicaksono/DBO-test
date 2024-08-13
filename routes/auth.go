package routes

import (
	"go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.GET("/home", controllers.Home)
	r.GET("/logout", controllers.Logout)

	r.GET("/customer_detail", controllers.GetDetailCustomer)
	r.GET("/search_customer", controllers.SearchCustomer)
	r.POST("/create_customer", controllers.InsertCustomer)
	r.PATCH("/update_customer", controllers.UpdateCustomer)
	r.DELETE("/delete_customer", controllers.DeleteCustomer)

	r.GET("/order_detail", controllers.GetDetailOrder)
	r.GET("/search_order", controllers.SearchOrder)
	r.POST("/create_order", controllers.InsertOrder)
	r.PATCH("/update_order", controllers.UpdateOrder)
	r.DELETE("/delete_order", controllers.DeleteOrder)
}
