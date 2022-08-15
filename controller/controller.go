package controller

import (
	"github.com/gin-gonic/gin"
	"go_todo/dao"
	"go_todo/models"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	// 从请求中拿数据
	var todo models.Todo
	c.BindJSON(&todo)

	// 存入数据库、并返回响应
	if err := dao.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetAllTodos(c *gin.Context) {
	var todoList []models.Todo
	if err := dao.DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var todo models.Todo
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var todo models.Todo
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err := dao.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	if err := dao.DB.Where("id=?", id).Delete(models.Todo{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
