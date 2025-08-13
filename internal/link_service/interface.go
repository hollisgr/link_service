package link_service

import "link_service/internal/model"

type LinkServiceInterface interface {
	NewTask() (id int, err error)
	AddLink(id int, link string) error
	GetStatus(id int) (model.TaskStatus, error)
	List() ([]model.Task, error)
}
