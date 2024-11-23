package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := parseEventID(context)
	if err != nil {
		utils.RespondWithError(context, http.StatusBadRequest, "Could not parse event ID.")
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		utils.RespondWithError(context, http.StatusInternalServerError, "Could not fetch the event.")
		return
	}

	if err := event.Register(userId); err != nil {
		utils.RespondWithError(context, http.StatusInternalServerError, "Could not register user for event.")
		return
	}

	utils.RespondWithMessage(context, http.StatusCreated, "Registered successfully.")
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := parseEventID(context)
	if err != nil {
		utils.RespondWithError(context, http.StatusBadRequest, "Could not parse event ID.")
		return
	}

	event := models.Event{ID: eventId}
	if err := event.CancelRegistration(userId); err != nil {
		utils.RespondWithError(context, http.StatusInternalServerError, "Could not cancel the registration.")
		return
	}

	utils.RespondWithMessage(context, http.StatusOK, "Registration cancelled successfully.")
}

func parseEventID(context *gin.Context) (int64, error) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return 0, utils.HandleError(err)
	}
	return eventId, nil
}
