package todo

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
	todos, err := s.Store.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

//
