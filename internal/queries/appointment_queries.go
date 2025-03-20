package queries

import "go-appointment-system/internal/aggregates"

type GetAppointmentQuery struct {
    ID string
}

type AppointmentRepository interface {
    FindAppointmentByID(id string) (*aggregates.Appointment, error)
}

func HandleGetAppointment(repo AppointmentRepository, query GetAppointmentQuery) (*aggregates.Appointment, error) {
    return repo.FindAppointmentByID(query.ID)
}
