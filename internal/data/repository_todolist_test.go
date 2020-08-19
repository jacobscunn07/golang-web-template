package data_test

import (
  "github.com/go-pg/pg/v10"
  "github.com/jacobscunn07/golang-web-template/internal/data"
  "github.com/jacobscunn07/golang-web-template/internal/domain"
  "testing"
)

func TestRead_HappyPath(t *testing.T) {
  t.Parallel()
  db := pg.Connect(&pg.Options{
    Addr:     "localhost:5432",
    User:     "postgres",
    Password: "postgres",
    Database: "postgres",
  })
  defer db.Close()

  repository := data.NewToDoListRepository(db)
  var todoList domain.ToDoList
  if err := repository.Read("John", &todoList); err != nil {
    t.Error(err)
  }

  if todoList.Key == "" {
   t.Error("Failed to find a TodoList with key ", "John")
  }

  if len(todoList.TodoListItems) != 1 {
    t.Error("Failed to retrieve TodoListItems for John TodoList")
  }
}
