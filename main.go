package main

import (
	"todolist-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// This line is VERY important!
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/", controllers.ShowTasks)
	r.POST("/add", controllers.AddTask)
	r.GET("/done/:id", controllers.MarkDone)
	r.GET("/delete/:id", controllers.DeleteTask)

	// Run app
	r.Run(":8080")
}
