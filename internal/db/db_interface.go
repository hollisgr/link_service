package db

import "link_service/internal/model"

type Storage interface {
	AddTask(id int) error
	AddLink(id int, link string)
	GetTask(id int) (task model.Task, err error)
	GetTaskList() (tasks []model.Task, err error)
	MarkTaskAsFinished(id int)
	CountCurrentTasks() int
	CountLinks(id int) int
}
