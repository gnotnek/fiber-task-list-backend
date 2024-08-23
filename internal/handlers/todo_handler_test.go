package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetTodos(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetTodos(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetTodos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTodoByID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetTodoByID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetTodoByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTodo(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteTodo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateTodoByID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateTodoByID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTodoByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompleteAllTodos(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompleteAllTodos(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CompleteAllTodos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCompletedTodos(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetCompletedTodos(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetCompletedTodos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
