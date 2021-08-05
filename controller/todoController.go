package controller

import (
	"gin-gorm-bublle/dao"
	"gin-gorm-bublle/models"
	"gin-gorm-bublle/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	TodoService service.ITodoService = &service.TodoService{&dao.TodoDao{}}
)

func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}


func CreateATodo(c *gin.Context) {
	// 前端页面填写待办事项， 点击提交 会发请求到这里
	var todo models.Todo
	c.BindJSON(&todo)
	if err := TodoService.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	var todoList []*models.Todo
	var err error
	if todoList, err = TodoService.GetAllTodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todoList)
	}
}


func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	var todo *models.Todo
	var err error
	var idInt, _ = strconv.Atoi(id)
	if todo, err = TodoService.GetATodo(idInt); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = dao.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}


func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := TodoService.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id: "删除成功"})
	}
}