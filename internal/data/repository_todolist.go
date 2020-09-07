package data

import (
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
  todolist := object.(*domain.ToDoList)
  _, err := r.db.Query(todolist, `SELECT id, key, name, description FROM todolist where key = ?`, key)
  if err != nil {
    return err
  }

  _, err = r.db.Query(&todolist.TodoListItems, `SELECT id, name, is_complete FROM todolistitem WHERE todolist_id = ?`, todolist.Id)
  if err != nil {
    return err
  }

  return nil
}

func (r *ToDoListRepository) Delete(key string) (int, error) {
  tx, err := r.db.Begin()
  if err != nil {
    return 0, err
  }

  result, err := tx.Exec(`DELETE FROM todolist WHERE key = ?`, key)
  if err != nil {
    return 0, err
  }

  err = tx.Commit()
  if err != nil {
    return 0, err
  }

  return result.RowsAffected(), nil
}

func (r *ToDoListRepository) Exists(key string) (bool, error) {
  var exists bool
  if _, err := r.db.QueryOne(pg.Scan(&exists), `SELECT 1 FROM todolist WHERE key = ?`, key); err != nil {
    return false, err
  }
  return exists, nil
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

  // TODO: Read back from database and set to object parameter memory address

  return count, nil
}
