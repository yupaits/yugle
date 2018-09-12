package yugle

import (
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
	"time"
)

const (
	IDLE    = iota //空闲
	RUNNING        //执行中
)

type Task struct {
	gorm.Model
	TaskName    string `gorm:"unique;not null"`
	State       int8   `gorm:"not null"`
	Prev        time.Time
	Next        time.Time
	LastSuccess bool
}

func SaveTask(task *Task) {
	db := DbConnect()
	defer db.Close()
	db.Save(task)
}

func GetTaskByTaskName(taskName string) *Task {
	db := DbConnect()
	defer db.Close()
	task := &Task{}
	db.Where("task_name = ?", taskName).Find(task)
	return task
}

func InsertTaskIfAbsent(taskName string) *Task {
	db := DbConnect()
	defer db.Close()
	task := GetTaskByTaskName(taskName)
	if task.TaskName == "" {
		task.TaskName = taskName
		task.State = IDLE
		db.Create(task)
	}
	return task
}

func UpdateTask(task *Task, cron *cron.Cron) {
	entry := cron.Entries()[0]
	if entry != nil {
		task.Prev = entry.Prev
		task.Next = entry.Next
		SaveTask(task)
	}
}

func GetTasks(page int, size int) *Pagination {
	db := DbConnect()
	defer db.Close()
	tasks := &[]Task{}
	db.Limit(size).Offset((page - 1) * size).Order("next").Find(tasks)
	var total int
	db.Table("tasks").Count(&total)
	return GenPage(page, size, total, tasks)
}
