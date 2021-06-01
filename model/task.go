package model

import "gorm.io/gorm"

const (
	// TaskDisable TaskDisable
	TaskDisable = 0
	// TaskAble TaskAble
	TaskAble = 1
	// TaskDone TaskDone
	TaskDone = 2
)

// Task Tassk
type Task struct {
	Code          string `json:"code ,omitempty"`
	MaskAmount    int    `json:"maskamount ,omitempty"`
	NextTimeStamp string `json:"nexttimestamp,omitempty"`
}

// TasksTable TasksTable, that's where we need to create and initialize all column
const TasksTable = "tasks"

// TaskModel TaskModel
var TaskModel = TasksModelObj{
	Table: TasksTable,
}

// TasksModelObj TasksModelObj
type TasksModelObj struct {
	Table string
}

// Create Create
func (m TasksModelObj) Create(task Task) error {
	return DBConn.Table(m.Table).Create(&task).Error
}

// GetTasks GetTasks
func (m TasksModelObj) GetTasks() ([]*Task, error) {
	var tasks []*Task
	// res := DBConn.Table(m.Table).
	// 	Select("code, mask_amount, next_time_stamp").
	// 	Order("next_time_stamp ASC").
	// 	Scan(&tasks)

	res := DBConn.Table(m.Table).
		Select("code, mask_amount, next_time_stamp").
		Scan(&tasks)

	if res.Error == gorm.ErrRecordNotFound {
		return tasks, nil
	}

	return tasks, res.Error
}
