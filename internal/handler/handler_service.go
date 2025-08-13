package handler

import (
	"encoding/base64"
	"fmt"
	"link_service/internal/config"
	"link_service/internal/link_service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	router      *gin.Engine
	linkService link_service.LinkServiceInterface
	config      *config.Config
	fileTypes   []string
}

func NewHandler(router *gin.Engine, linkService link_service.LinkServiceInterface, config *config.Config) HandlerInterface {
	return &handler{
		router:      router,
		linkService: linkService,
		config:      config,
		fileTypes:   strings.Split(config.FileTypes, ","),
	}
}

func (h *handler) Register() {
	h.router.GET("/task/:id/status", h.Status)
	h.router.GET("/task", h.List)
	h.router.POST("/task", h.CreateTask)
	h.router.POST("/task/:id", h.AddLink)

}

func (h *handler) Status(c *gin.Context) {
	id, err := h.getIdFromQuery(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id param is required",
		})
		return
	}

	data, err := h.linkService.GetStatus(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprint(err),
		})
		return
	}

	archiveLink := ""
	if len(data.ArchiveBytes) != 0 {
		encoded := base64.StdEncoding.EncodeToString(data.ArchiveBytes)
		archiveLink = fmt.Sprintf("data:application/zip;base64,%s", encoded)
	}

	c.JSON(http.StatusOK, gin.H{
		"download_link": archiveLink,
		"task":          data.Task,
	})
}

func (h *handler) CreateTask(c *gin.Context) {
	id, err := h.linkService.NewTask()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprint(err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("task id: %d", id),
	})
}

func (h *handler) AddLink(c *gin.Context) {
	id, err := h.getIdFromQuery(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id param is required",
		})
		return
	}

	data := make(map[string]string, 0)
	err = c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "required link in request body",
		})
		return
	}

	err = h.checkContentType(data["link"])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprint(err),
		})
		return
	}

	err = h.linkService.AddLink(id, data["link"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "link added",
	})
}

func (h *handler) List(c *gin.Context) {
	tasks, err := h.linkService.List()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "task list is empty",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": tasks,
		})
	}
}

func (h *handler) getIdFromQuery(c *gin.Context) (int, error) {
	idParams := c.Params.ByName("id")
	id := 0
	_, err := fmt.Sscanf(idParams, "%d", &id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (h *handler) checkContentType(link string) error {
	resp, err := http.Head(link)
	if err != nil {
		return fmt.Errorf("can't reach file info")
	}
	contentType := resp.Header.Get("Content-Type")
	for _, t := range h.fileTypes {
		if t == contentType {
			return nil
		}
	}
	return fmt.Errorf("file type not allowed")
}
