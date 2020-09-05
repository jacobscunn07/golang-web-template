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
  return nil
}

func (r *ToDoListRepository) Delete(key string) (int, error) {
  return 0, errors.New("")
}

func (r *ToDoListRepository) Exists(key string) (bool, error) {
  return false, nil
}

func (r *ToDoListRepository) Save(object interface{}) (int, error){
  todoList, _ := object.(*domain.ToDoList)
  count := 0

  tx, err := r.db.Begin()
  if err != nil {
    return 0, err
  }

  result, err := tx.Exec(`
    INSERT INTO todolist (
    id,
    key,
    name,
    description
    )
    VALUES (?, ?, ?, ?)`, todoList.Id, todoList.Key, todoList.Name, todoList.Description)
  if err != nil {
    return 0, err
  }

  count += result.RowsAffected()

  for _, tdli := range todoList.TodoListItems {
    result, err = tx.Exec(`
    INSERT INTO todolistitem (
    id,
    todolist_id,
    name,
    is_complete
    )
    VALUES (?, ?, ?, ?)`, tdli.Id, todoList.Id, tdli.Name, tdli.IsComplete)
    if err != nil {
      return 0, err
    }
    count += result.RowsAffected()
  }

  err = tx.Commit()
  if err != nil {
    return 0, err
  }

  return count, nil
}
