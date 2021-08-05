package main

import (
	"gin-gorm-bublle/controller"
	"gin-gorm-bublle/dao"
	"gin-gorm-bublle/models"
	"github.com/gin-gonic/gin"
)


func main() {
	//创建数据库
	// sql: CREATE DATABASE bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil{
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	r.Run()
}
