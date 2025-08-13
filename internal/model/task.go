package model

const (
	IN_PROGRESS = "in_progress"
	FINISHED    = "finished"
)

type Task struct {
	Id     int      `json:"id"`
	Links  []string `json:"links"`
	Status string   `json:"status"`
}

type TaskStatus struct {
	Task         Task
	ArchiveBytes []byte
}
