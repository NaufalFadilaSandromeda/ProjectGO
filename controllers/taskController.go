package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"
	"todolist-gin/models"

	"github.com/gin-gonic/gin"
)

var tasks, _ = models.LoadTasks()

func ShowTasks(c *gin.Context) {
	fmt.Println("🔧 ShowTasks handler is working!")
	// Sort tasks by deadline (empty deadline goes to the bottom)
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Deadline.IsZero() {
			return false
		}
		if tasks[j].Deadline.IsZero() {
			return true
		}
		return tasks[i].Deadline.Before(tasks[j].Deadline)
	})

	c.HTML(http.StatusOK, "index", gin.H{
		"tasks": tasks,
	})
}

func AddTask(c *gin.Context) {
	description := c.PostForm("description")
	deadlineStr := c.PostForm("deadline")
	var deadline time.Time

	if deadlineStr != "" {
		parsed, err := time.Parse("2006-01-02", deadlineStr)
		if err == nil {
			deadline = parsed
		}
	}

	if description != "" {
		tasks = append(tasks, models.Task{
			Description: description,
			Deadline:    deadline,
		})
		models.SaveTasks(tasks)
	}
	c.Redirect(http.StatusFound, "/")
}

func MarkDone(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil || index < 0 || index >= len(tasks) {
		c.String(http.StatusBadRequest, "Invalid task ID")
		return
	}

	tasks[index].Done = !tasks[index].Done // toggle
	models.SaveTasks(tasks)
	c.Redirect(http.StatusFound, "/")
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil || index < 0 || index >= len(tasks) {
		c.String(http.StatusBadRequest, "Invalid task ID")
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	models.SaveTasks(tasks)
	c.Redirect(http.StatusFound, "/")
}
