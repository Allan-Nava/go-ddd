package todo

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TodoStore interface {
	//
	GetAll() ([]Todo, error)
	Create(todo *Todo) error
	//Update(todo *Todo) error
}

type MySqlTodoStore struct {
	DB *gorm.DB
}

//
//Create(todo *Todo) error
//Update(todo *Todo) error
func (s *MySqlTodoStore) GetAll() ([]Todo, error) {
	var todo []Todo
	query := s.DB.Table("customer").Select("*").Find(&todo)
	if query.Error != nil {
		logrus.Error(query.Error.Error())
		return nil, query.Error
	}
	return todo, nil
}

//
func (s *MySqlTodoStore) Create(todo *Todo) error {
	query := s.DB.Save(*todo)
	if query.Error != nil {
		logrus.Error(query.Error.Error())
		return query.Error
	}
	return nil
}

//

func (s *MySqlTodoStore) Update(todo *Todo) error {
	query := s.DB.Model(&Todo{}).Where("id = ?", todo.ID).Update("name", todo.Name)
	if query.Error != nil {
		logrus.Error(query.Error.Error())
		return query.Error
	}
	return nil
}
