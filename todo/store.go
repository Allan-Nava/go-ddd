package todo
/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */

type TodoStore interface {
	GetByClientId(clientId string) (*Todo, error)
	GetAll() ([]Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
}