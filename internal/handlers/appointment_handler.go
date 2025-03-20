package handlers

import (
    "go-appointment-system/internal/commands"
    "go-appointment-system/internal/queries"
    "go-appointment-system/internal/repository"
    "net/http"
    "github.com/gin-gonic/gin"
)

func CreateAppointmentHandler(repo *repository.InMemoryAppointmentRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        var cmd commands.CreateAppointmentCommand
        if err := c.ShouldBindJSON(&cmd); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        appointment, err := commands.HandleCreateAppointment(repo, cmd)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, appointment)
    }
}

func GetAppointmentHandler(repo *repository.InMemoryAppointmentRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Query("id")
        if id == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing appointment ID"})
            return
        }

        appointment, err := queries.HandleGetAppointment(repo, queries.GetAppointmentQuery{ID: id})
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, appointment)
    }
}
