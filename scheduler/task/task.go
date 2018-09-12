package task

import (
	"github.com/jinzhu/gorm"
	"time"
	"yugle/common/pagination"
	"yugle/dbutils"
)

const (
	IDLE    = iota //空闲
	RUNNING        //执行中
)

type Task struct {
	gorm.Model
	TaskName string `gorm:"unique;not null"`
	State    int8   `gorm:"not null"`
	Prev     time.Time
	Next     time.Time
}

func SaveTask(task *Task) {
	db := dbutils.Connect()
	defer db.Close()
	db.Save(task)
}

func GetTaskByTaskName(taskName string) *Task {
	db := dbutils.Connect()
	defer db.Close()
	task := &Task{}
	db.Where("task_name = ?", taskName).Find(task)
	return task
}

func InsertTaskIfAbsent(taskName string) *Task {
	db := dbutils.Connect()
	defer db.Close()
	task := GetTaskByTaskName(taskName)
	if task.TaskName == "" {
		task.TaskName = taskName
		task.State = IDLE
		db.Create(task)
	}
	return task
}

func GetTasks(page int, size int) *pagination.Pagination {
	db := dbutils.Connect()
	defer db.Close()
	tasks := &[]Task{}
	db.Limit(size).Offset((page - 1) * size).Order("next").Find(tasks)
	var total int
	db.Table("tasks").Count(&total)
	return pagination.GenPage(page, size, total, tasks)
}
