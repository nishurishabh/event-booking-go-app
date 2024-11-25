package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func heartbeat(context *gin.Context) {
	context.JSONP(http.StatusOK, gin.H{"message": "Heart beating for the server."})
}
