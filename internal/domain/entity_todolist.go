package domain

type ToDoList struct {
  AggregateRoot
  Id string
  Name string
  Description string
  TodoListItems []*ToDoListItem
}
