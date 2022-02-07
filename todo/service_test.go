package todo

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */
import (
	"reflect"
	"testing"
)

func TestTodoService_GetAll(t *testing.T) {
	type fields struct {
		Store TodoStore
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				Store: tt.fields.Store,
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
