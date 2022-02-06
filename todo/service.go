package todo

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */
type TodoService struct {
	Store TodoStore
	//TodoService *todo.TodoService
}

//
func (s *TodoService) All() ([]Todo, error) {
	todos, err := s.Store.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

//
func (s *TodoService) Create(todo *Todo) error {
	err := s.Store.Create(todo)
	if err != nil {
		return err
	}
	return nil
}
