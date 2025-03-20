package main

import (
    "fmt"
    "log"
    "net/http"
    "go-appointment-system/internal/handlers"
    "go-appointment-system/internal/repository"
    "github.com/gin-gonic/gin"
)

func main() {
    repo := repository.NewInMemoryBookingRepository()
    r := gin.Default()

    r.POST("/appointments/create", handlers.CreateAppointmentHandler(repo))
    r.GET("/appointments/get", handlers.GetAppointmentHandler(repo))
    r.POST("/custom-fields/create", handlers.CreateCustomFieldHandler(repo))
    r.GET("/dashboard/stats", handlers.GetDashboardStatsHandler(repo))

    fmt.Println("Server running on port 8080...")
    log.Fatal(r.Run(":8080"))
}
