package data

import (
  "errors"
  "github.com/go-pg/pg/v10"
  "github.com/jacobscunn07/golang-web-template/internal/domain"
)

type Repository interface{
  Read(string, interface{}) error
  Delete(string) (int, error)
  Exists(string) (bool, error)
  Save(interface{}) (int, error)
}

type ToDoListRepository struct {
  db *pg.DB
}

func NewToDoListRepository(db *pg.DB) *ToDoListRepository {
  return &ToDoListRepository{db: db}
}

func (r *ToDoListRepository) Read(key string, object interface{}) error {
  todoList, _ := object.(*domain.ToDoList)

  if _, err := r.db.
   Query(todoList, `SELECT id, key, name, description FROM todolist WHERE key = ?`, key);
  err != nil {
    return err
  }

  if _, err := r.db.
   Query(&todoList.TodoListItems,
     `SELECT name, is_complete FROM todolistitem WHERE todolist_id = ?`, todoList.Id);
  err != nil {
   return err
  }

  object = todoList

  return nil
}

func (r *ToDoListRepository) Delete(key string) (int, error) {
  return 0, errors.New("")
}

func (r *ToDoListRepository) Exists(key string) (bool, error) {
  return false, errors.New("")
}

func (r *ToDoListRepository) Save(object interface{}) (int, error){
  return 0, errors.New("")
}
