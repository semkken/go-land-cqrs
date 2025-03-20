package aggregates

import (
    "errors"
    "time"
    "github.com/google/uuid"
)

type Appointment struct {
    ID          string
    UserID      string
    StaffID     string
    Location    string
    StartTime   time.Time
    EndTime     time.Time
    CreatedAt   time.Time
    CustomFields map[string]string
}

func CreateNewAppointment(userID, staffID, location, start, end string, customFields map[string]string) (*Appointment, error) {
    startTime, err := time.Parse("2006-01-02 15:04", start)
    if err != nil {
        return nil, errors.New("invalid start time format")
    }

    endTime, err := time.Parse("2006-01-02 15:04", end)
    if err != nil {
        return nil, errors.New("invalid end time format")
    }

    if endTime.Before(startTime) {
        return nil, errors.New("end time must be after start time")
    }

    return &Appointment{
        ID:           uuid.New().String(),
        UserID:       userID,
        StaffID:      staffID,
        Location:     location,
        StartTime:    startTime,
        EndTime:      endTime,
        CreatedAt:    time.Now(),
        CustomFields: customFields,
    }, nil
}
