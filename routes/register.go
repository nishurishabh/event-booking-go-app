package routes

import (
	"example.com/event-booking-backend-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id to valid event id"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id to valid event id"})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.Cancel(userID)
	if err != nil {
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel for the event"})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
