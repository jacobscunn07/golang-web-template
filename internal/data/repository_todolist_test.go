package data_test

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/jacobscunn07/golang-web-template/internal/data"
	"github.com/jacobscunn07/golang-web-template/internal/domain"
	"os"
  "reflect"
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
		  temp := tc.todoList
			count, err := repository.Save(&tc.todoList)

			if count != tc.expected.count {
				t.Errorf(fmt.Sprintf("Expected to affect %d rows, but only affected %d", tc.expected.count, count))
			}

			if err != tc.expected.err {
				t.Errorf(fmt.Sprintf("Expected to get error of type %T, but got %T instead", tc.expected.err, err))
			}

			if !reflect.DeepEqual(tc.todoList, temp) {
			  t.Errorf("Expected todolist and actual were not equal")
      }
		})
	}
}

func TestExists(t *testing.T) {
  tests := map[string]struct {
    key string
    save bool
    expected struct {
      exists bool
    }
  }{
    "simple exists": {
      key: uuid.New().String(),
      save: true,
      expected: struct {
        exists bool
      }{exists: true},
    },
    "simple does not exist": {
      key: uuid.New().String(),
      save: false,
      expected: struct {
        exists bool
      }{exists: false},
    },
  }

  repository := data.NewToDoListRepository(DB)

  for name, tc := range tests {
    t.Run(name, func(t *testing.T) {
      if tc.save {
        repository.Save(&domain.ToDoList{
          AggregateRoot: domain.AggregateRoot{Key: tc.key},
          Id:            uuid.New().String(),
          Name:          uuid.New().String(),
          Description:   uuid.New().String(),
        })
      }

      exists, _ := repository.Exists(tc.key)

      if exists != tc.expected.exists {
        t.Errorf(fmt.Sprintf("Expected the todolist to exist with %v, but got %v", tc.expected.exists, exists))
      }
    })
  }
}

func TestDelete(t *testing.T) {
  tests := map[string]struct {
    key string
    save bool
    expected struct {
      count int
      err error
    }
  }{
    "simple": {
      key: uuid.New().String(),
      save: true,
      expected: struct {
        count int
        err error
      }{count: 1, err: nil},
    },
  }

  repository := data.NewToDoListRepository(DB)

  for name, tc := range tests {
    t.Run(name, func(t *testing.T) {
      if tc.save {
        repository.Save(&domain.ToDoList{
          AggregateRoot: domain.AggregateRoot{Key: tc.key},
          Id:            uuid.New().String(),
          Name:          uuid.New().String(),
          Description:   uuid.New().String(),
        })
      }

      count, err := repository.Delete(tc.key)
      // TODO: Try to read todolist item with key that was deleted - should not exist

      if count != tc.expected.count {
        t.Errorf(fmt.Sprintf("Expected to affect %d rows, but only affected %d", tc.expected.count, count))
      }

      if err != tc.expected.err {
        t.Errorf(fmt.Sprintf("Expected to get error of type %T, but got %T instead", tc.expected.err, err))
      }
    })
  }
}
