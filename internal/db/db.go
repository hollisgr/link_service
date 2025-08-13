package db

import (
	"fmt"
	"link_service/internal/model"
)

const (
	DB_MAX_TASKS = 50
)

type database struct {
	db map[int]model.Task
}

func NewStorage() Storage {
	return &database{
		db: make(map[int]model.Task, 0),
	}
}

func (d *database) AddTask(id int) error {
	if len(d.db) > DB_MAX_TASKS {
		return fmt.Errorf("database add task error, max available tasks: %d", DB_MAX_TASKS)
	}
	_, exists := d.db[id]
	if exists {
		return fmt.Errorf("database add task error, id already exists")
	}
	newTask := model.Task{
		Id:     id,
		Status: model.IN_PROGRESS,
	}
	d.db[id] = newTask
	return nil
}

func (d *database) AddLink(id int, link string) {
	tempTask := d.db[id]
	tempTask.Links = append(tempTask.Links, link)
	d.db[id] = tempTask
}

func (d *database) GetTask(id int) (task model.Task, err error) {
	_, exists := d.db[id]
	if !exists {
		return task, fmt.Errorf("database get task error, not found")
	}
	return d.db[id], nil
}

func (d *database) GetTaskList() (tasks []model.Task, err error) {
	if len(d.db) == 0 {
		return nil, fmt.Errorf("database get task list error, list is empty")
	}
	for key := range d.db {
		tasks = append(tasks, d.db[key])
	}
	return tasks, nil
}

func (d *database) MarkTaskAsFinished(id int) {
	tempTask := d.db[id]
	tempTask.Status = model.FINISHED
	d.db[id] = tempTask
}

func (d *database) CountCurrentTasks() int {
	counter := 0
	for key := range d.db {
		if d.db[key].Status != model.FINISHED {
			counter++
		}
	}
	return counter
}

func (d *database) CountLinks(id int) int {
	return len(d.db[id].Links)
}
