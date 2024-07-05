package handler

import (
	"net/http"
	track "timeTracker"

	"github.com/gin-gonic/gin"
)

// @Summary createUser
// @Tags users
// @Description create user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body track.People true "user info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input track.People

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.TrackPeople.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary deleteUser
// @Tags users
// @Description delete user
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/{id} [delete]

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TrackPeople.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary updateUser
// @Tags users
// @Description update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body track.People true "user info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {

	var input track.People

	userId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.TrackPeople.UpdateUser(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type GetAllPeopleRespomse struct {
	Data []track.People `json:"data"`
}

// @Summary getPeopleByFilters
// @Tags users
// @Description get people by filters
// @ID get-people-by-filters
// @Accept  json
// @Produce  json
// @Param input body track.Filter true "filters"
// @Success 200 {object} GetAllPeopleRespomse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users [get]
func (h *Handler) getPeopleByFilters(c *gin.Context) {
	var filters track.Filter

	if err := c.BindJSON(&filters); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	people, err := h.services.TrackPeople.GetPeopleByFilters(filters)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllPeopleRespomse{
		Data: people,
	})
}
