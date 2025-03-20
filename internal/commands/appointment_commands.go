package commands

import "go-appointment-system/internal/aggregates"

type CreateAppointmentCommand struct {
    UserID       string
    StaffID      string
    Location     string
    StartTime    string
    EndTime      string
    CustomFields map[string]string
}

type AppointmentRepository interface {
    SaveAppointment(appointment *aggregates.Appointment) error
    FindAppointmentByID(id string) (*aggregates.Appointment, error)
}

func HandleCreateAppointment(repo AppointmentRepository, cmd CreateAppointmentCommand) (*aggregates.Appointment, error) {
    appointment, err := aggregates.CreateNewAppointment(cmd.UserID, cmd.StaffID, cmd.Location, cmd.StartTime, cmd.EndTime, cmd.CustomFields)
    if err != nil {
        return nil, err
    }
    err = repo.SaveAppointment(appointment)
    return appointment, err
}
