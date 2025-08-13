package link_service

import (
	"bytes"
	"fmt"
	"link_service/internal/config"
	"link_service/internal/db"
	"link_service/internal/model"
	"strings"
)

const (
	MAX_TASKS = 3
	MAX_LINKS = 3
)

type linkService struct {
	storage db.Storage
	config  *config.Config
}

func NewService(config *config.Config, storage db.Storage) LinkServiceInterface {
	return &linkService{
		config:  config,
		storage: storage,
	}
}

func (s *linkService) NewTask() (id int, err error) {
	if s.storage.CountCurrentTasks() >= MAX_TASKS {
		return id, fmt.Errorf("unable to create new task, limit exceeded: %d", MAX_TASKS)
	}
	tasks, err := s.storage.GetTaskList()
	if err != nil { // if task list is empty
		id = 1
	} else {
		id = tasks[len(tasks)-1].Id + 1 // getting last id in task list
	}
	err = s.storage.AddTask(id)
	if err != nil {
		return id, fmt.Errorf("unable to create new task: %v", err)
	}
	return id, nil
}

func (s *linkService) List() ([]model.Task, error) {
	tasks, err := s.storage.GetTaskList()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *linkService) AddLink(id int, link string) error {
	if s.storage.CountLinks(id) >= MAX_LINKS {
		return fmt.Errorf("unable to add new link, limit exceeded: %d", MAX_LINKS)
	}
	task, err := s.storage.GetTask(id)
	if err != nil {
		return err
	}
	for _, l := range task.Links {
		if l == link {
			return fmt.Errorf("service add link error, re-adding link")
		}
	}
	s.storage.AddLink(id, link)
	return nil
}

func (s *linkService) GetStatus(id int) (model.TaskStatus, error) {
	resp := model.TaskStatus{}
	task, err := s.storage.GetTask(id)
	if err != nil {
		return resp, fmt.Errorf("unable to get status: %v", err)

	}
	resp.Task = task
	if s.storage.CountLinks(id) == MAX_LINKS {
		bytes, err := s.generateArchive(task)
		if err != nil {
			return resp, err
		}
		resp.ArchiveBytes = bytes
		resp.Task.Status = model.FINISHED
		s.storage.MarkTaskAsFinished(id)
	}
	return resp, nil
}

func (s *linkService) generateArchive(task model.Task) ([]byte, error) {
	buffers := make([]*bytes.Buffer, 0)
	filenames := make([]string, 0)
	for _, link := range task.Links {
		buffer, err := s.getFile(link)
		if err != nil {
			return nil, err
		}
		linkSplited := strings.Split(link, "/")
		buffers = append(buffers, buffer)
		filenames = append(filenames, linkSplited[len(linkSplited)-1])
	}
	res, err := s.createArchive(buffers, filenames)
	if err != nil {
		return nil, err
	}
	return res, nil
}
