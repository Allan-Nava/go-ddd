package todo

import "fmt"

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */

type ITodoService interface {
	GetAll() ([]Todo, error)
	Create(name string) error
}


type TodoService struct {
	store TodoStore
}

func NewService(todoStore TodoStore) ITodoService{
	return &TodoService{

	}
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
func (s *TodoService) Create(name string) error {
	fmt.Printf("Create() %v", name)
	todo := &Todo{Name: name}
	err := s.Store.Create(todo)
	if err != nil {
		return err
	}
	return nil
}
/*
func (s *TodoService) Update(todo *Todo) error {
	fmt.Print("Updated()")
	err := s.Store.Update(todo)
	if err != nil {
		return err
	}
	return nil
}
*/
