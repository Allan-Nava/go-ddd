package todo

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 04/03/2025
*
 */

type Todo struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Tabler interface {
	TableName() string
}

func (Todo) TableName() string {
	return "todo"
}
