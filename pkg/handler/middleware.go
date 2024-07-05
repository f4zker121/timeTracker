package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	idCtx = "id"
)

func getId(c *gin.Context) (int, error) {
	id := c.Param(idCtx)

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return 0, err
	}

	return idInt, nil
}
