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
func (s *TodoService) GetAll() ([]Todo, error) {
	fmt.Print("GetAll()")
	todos, err := s.Store.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

//
/*func (s *TodoService) Create(todo *Todo) error {
	fmt.Print("Create()")
	err := s.Store.Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) Update(todo *Todo) error {
	fmt.Print("Updated()")
	err := s.Store.Update(todo)
	if err != nil {
		return err
	}
	return nil
}
*/
