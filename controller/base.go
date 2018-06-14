package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rs/zerolog/log"
)

func renderJSONWithError(c *gin.Context, error string) {
	log.Error().Msg(error)
	c.JSON(http.StatusBadRequest, gin.H{"error": error})
	return
}
