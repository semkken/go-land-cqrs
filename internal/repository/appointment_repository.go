package repository

import (
    "errors"
    "go-appointment-system/internal/aggregates"
)

type InMemoryAppointmentRepository struct {
    appointments map[string]*aggregates.Appointment
}

func NewInMemoryAppointmentRepository() *InMemoryAppointmentRepository {
    return &InMemoryAppointmentRepository{appointments: make(map[string]*aggregates.Appointment)}
}

func (r *InMemoryAppointmentRepository) SaveAppointment(appointment *aggregates.Appointment) error {
    r.appointments[appointment.ID] = appointment
    return nil
}

func (r *InMemoryAppointmentRepository) FindAppointmentByID(id string) (*aggregates.Appointment, error) {
    if appointment, exists := r.appointments[id]; exists {
        return appointment, nil
    }
    return nil, errors.New("appointment not found")
}
