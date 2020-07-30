package domain

type ToDoListItem struct {
  Id string
  Name string
  Status ToDoListItemStatus
}

type ToDoListItemStatus string

const(
  NotStarted ToDoListItemStatus = "NotStarted"
  InProgress = "InProgress"
  Finished = "Finished"
)
