package todo
/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */

type Todo struct {
	ID                      int    `json:"id" gorm:"primaryKey"`
	Name 					string `json:"name"` 				
}