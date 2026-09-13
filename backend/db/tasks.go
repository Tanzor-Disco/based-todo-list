package db

import (
	"context"
	"fmt"

	"github.com/Tanzor-Disco/based-todo-list/backend/models"
)

func (db Database) CreateTask(ctx context.Context, task models.Task) (models.TaskID, error) {
	var taskID models.TaskID
	err := db.pool.QueryRow(ctx,
		`
	INSERT INTO tasks(user_id,description)
	VALUES ($1,$2)
	RETURNING id
	`,
		task.UserID, task.Description).Scan(&taskID)
	if err != nil {
		return 0, fmt.Errorf("CreateTask: %w", err)
	}
	return taskID, nil
}

func (db Database) GetTasksByUserID(ctx context.Context, userID models.UserID) ([]models.Task, error) {
	rows, err := db.pool.Query(ctx,
		`
	SELECT * FROM tasks WHERE user_id = $1
	`,
		userID)
	if err != nil {
		return make([]models.Task, 0), fmt.Errorf("GetTasksByUserID: %w", err)
	}

	tasks := make([]models.Task, 0, 10)

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.TaskID, &task.UserID, &task.Description)
		if err != nil {
			return make([]models.Task, 0), fmt.Errorf("GetTasksByUserID: %w", err)
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}
