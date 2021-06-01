package handler

import (
	"Backend/model"
	"Backend/pkg/res"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Ping Ping
func Ping(c *gin.Context) {
	res.Success(c, gin.H{
		"msg": "pong",
	})
}

// CreateTask CreateTask
func CreateTask(c *gin.Context) {
	var taskData model.Task
	c.BindJSON(&taskData)

	if err := model.TaskModel.Create(taskData); err != nil {
		log.WithFields(log.Fields{
			"task_data":  taskData,
			"origin_err": err.Error(),
		}).Error("db error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	res.Success(c, gin.H{})
}

// GetGroups GetGroups
func GetReturn(c *gin.Context) {

	tasks, err := model.TaskModel.GetTasks()

	if err != nil {
		log.WithFields(log.Fields{
			"task_data":  tasks,
			"origin_err": err.Error(),
		}).Error("db error")
		res.SystemError(c, res.ErrSystemCode, gin.H{})
		return
	}

	res.Success(c, gin.H{
		"tasks": tasks,
	})
}
