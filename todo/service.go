package todo

import "fmt"

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */
type TodoService struct {
	Store TodoStore
}

//
func (s *TodoService) All() ([]Todo, error) {
	fmt.Print("All()")
	todos, err := s.Store.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

//
func (s *TodoService) Create(todo *Todo) error {
	fmt.Print("All()")
	err := s.Store.Create(todo)
	if err != nil {
		return err
	}
	return nil
}
