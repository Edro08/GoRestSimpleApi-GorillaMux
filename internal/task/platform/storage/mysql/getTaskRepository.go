package mysql

import (
	"GoRestSimpleApi/internal/task"
)

func (r TaskRepository) ReadID(Id string) ([]task.TaskDB, error) {

	sqlquery := "Select * from task where Id = " + Id
	rows, err := r.db.Query(sqlquery)

	if err != nil {
		return nil, err
	}

	tasks := make([]task.TaskDB, 0)

	for rows.Next() {
		var t task.TaskDB
		err := rows.Scan(&t.ID, &t.Name, &t.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
