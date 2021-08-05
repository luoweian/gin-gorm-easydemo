package dao

import (
	"gin-gorm-bublle/models"
)



type ITodoDao interface{
	CreateATodo(*models.Todo) error
	GetAllTodo ()([]*models.Todo, error)
	GetATodo(int)(*models.Todo, error)
	UpdateATodo(*models.Todo) error
	DeleteTodo(string)error
}

type TodoDao struct{

}

func NewTodoDao() ITodoDao{
	return &TodoDao{}
}

func (t *TodoDao) CreateATodo(todo *models.Todo)(err error){
	err = DB.Create(&todo).Error
	return
}

func (t *TodoDao) GetAllTodo()(todoList []*models.Todo, err error){
	err = DB.Find(&todoList).Error
	return
}

func (t *TodoDao) GetATodo(id int)(todo *models.Todo, err error){
	err = DB.First(id).Error
	return
}

func (t *TodoDao) UpdateATodo(todo *models.Todo)(err error){
	err = DB.Save(&todo).Error
	return
}

func (t *TodoDao) DeleteTodo(id string)(err error){
	return DB.Where("id = ?", id).Delete(&models.Todo{}).Error
}
