package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"net/http"
)

func renderJSONWithError(c *gin.Context, error string) {
	logger.Error(error)
	c.JSON(http.StatusBadRequest, gin.H{"error": error})
	return
}
