package domain

type ToDoList struct {
  Id string
  Name string
  Description string
  Items []*ToDoListItem
}
