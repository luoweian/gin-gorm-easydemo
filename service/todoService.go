package service

import (
	"gin-gorm-bublle/dao"
	"gin-gorm-bublle/models"
)

type ITodoService interface{
	CreateATodo(*models.Todo) error
	GetAllTodo ()([]*models.Todo, error)
	GetATodo(int)(*models.Todo, error)
	UpdateATodo(*models.Todo) error
	DeleteTodo(string)error
}

type TodoService struct{
	TodoDao dao.ITodoDao
}

func (t *TodoService) CreateATodo(todo *models.Todo)(err error){
	err = t.TodoDao.CreateATodo(todo)
	return
}

func (t *TodoService) GetAllTodo()(todoList []*models.Todo, err error){
	todoList, err = t.TodoDao.GetAllTodo()
	return
}

func (t *TodoService) GetATodo(id int)(todo *models.Todo, err error){
	todo, err = t.TodoDao.GetATodo(id)
	return
}

func (t *TodoService) UpdateATodo(todo *models.Todo)(err error){
	err = t.TodoDao.UpdateATodo(todo)
	return
}

func (t *TodoService) DeleteTodo(id string)(err error){
	return t.TodoDao.DeleteTodo(id)
}
