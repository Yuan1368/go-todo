package main

import (
	"github.com/gin-gonic/gin"
	"go_todo/controller"
	"go_todo/dao"
	"go_todo/models"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项

		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看
		v1Group.GET("/todos", controller.GetAllTodos)
		// 查看其中一条
		v1Group.GET("/todo/:id", controller.GetTodo)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	r.Run(":8088")
}
