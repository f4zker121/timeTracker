package handler

import (
	"net/http"
	track "timeTracker"

	"github.com/gin-gonic/gin"
)

// @Summary createTask
// @Tags tasks
// @Description create a new task
// @ID create-task
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body track.Task true "Task info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/users/{id}/tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	var input track.Task

	userId, err := getId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.TrackTask.CreateTask(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary startTask
// @Tags tasks
// @Description start a task
// @ID start-task
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/tasks/start/{id} [put]
func (h *Handler) startTask(c *gin.Context) {
	taskId, err := getId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TrackTask.StartTask(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary endTask
// @Tags tasks
// @Description end a task
// @ID end-task
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/tasks/end/{id} [put]
func (h *Handler) endTask(c *gin.Context) {
	taskId, err := getId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TrackTask.EndTask(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary deleteTask
// @Tags tasks
// @Description delete a task
// @ID delete-task
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/tasks/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	taskId, err := getId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TrackTask.DeleteTask(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type GetTasksRespomse struct {
	Data []track.Task `json:"data"`
}

// @Summary getTasksByLaborCosts
// @Tags tasks
// @Description get tasks by labor costs
// @ID get-tasks-by-labor-costs
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} GetTasksRespomse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/tasks/{id} [get]
func (h *Handler) getTasksByLaborCosts(c *gin.Context) {
	userId, err := getId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	tasks, err := h.services.TrackTask.GetTasksByLaborCosts(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetTasksRespomse{
		Data: tasks,
	})
}
