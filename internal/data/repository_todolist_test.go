package data_test

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/jacobscunn07/golang-web-template/internal/data"
	"github.com/jacobscunn07/golang-web-template/internal/domain"
	"os"
	"testing"
)

var DB *pg.DB

func TestMain(m *testing.M) {
	DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	defer DB.Close()
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestSave(t *testing.T) {
	tests := map[string]struct {
		todoList domain.ToDoList
		expected struct {
			count int
			err   error
		}
	}{
		"simple": {
			todoList: domain.ToDoList{
				AggregateRoot: domain.AggregateRoot{Key: "simple"},
				Id:            uuid.New().String(),
				Name:          "Golang",
				Description:   "learn golang",
			},
			expected: struct {
				count int
				err   error
			}{count: 1, err: nil},
		},
		"simple w/ items": {
			todoList: domain.ToDoList{
				AggregateRoot: domain.AggregateRoot{Key: "simple-with-items"},
				Id:            uuid.New().String(),
				Name:          "Golang",
				Description:   "learn golang",
				TodoListItems: []*domain.ToDoListItem{{
					Id:         uuid.New().String(),
					Name:       "write go code",
					IsComplete: false,
				}},
			},
			expected: struct {
				count int
				err   error
			}{count: 2, err: nil},
		},
	}

	repository := data.NewToDoListRepository(DB)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			count, err := repository.Save(&tc.todoList)

			if count != tc.expected.count {
				t.Errorf(fmt.Sprintf("Expected to effect %d rows, but only effected %d", tc.expected.count, count))
			}

			if err != tc.expected.err {
				t.Errorf(fmt.Sprintf("Expected to get error of type %T, but got %T instead", tc.expected.err, err))
			}
		})
	}
}
