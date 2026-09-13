package models

type TaskID int

type Task struct {
	TaskID      TaskID
	UserID      UserID
	Description string
}

func NewTask(userID UserID, description string) Task {
	return Task{
		UserID:      userID,
		Description: description,
	}
}
