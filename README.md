# Go Appointment Booking API

This is a **full-featured appointment booking API** built using **CQRS (Command Query Responsibility Segregation)** in Go.

## Features
- **Create Appointments** (`POST /appointments/create`)
- **Get Appointment Details** (`GET /appointments/get?id=appointment_id`)
- **Custom Fields Management** (`POST /custom-fields/create`)
- **Dashboard Statistics** (`GET /dashboard/stats`)
- **Custom Availability** for staff and locations.
- **Appointment Buffering** to prevent last-minute bookings.
- **User Account Management** for self-service bookings.

## Running the API
```sh
go run cmd/main.go
```

### Example Requests:

#### Create an Appointment
```sh
curl -X POST http://localhost:8080/appointments/create -H "Content-Type: application/json" -d '{
  "UserID": "user-123",
  "StaffID": "staff-456",
  "Location": "Room A",
  "StartTime": "2025-04-01 10:00",
  "EndTime": "2025-04-01 11:00",
  "CustomFields": {"note": "Bring documents"}
}'
```

#### Get Appointment by ID
```sh
curl -X GET "http://localhost:8080/appointments/get?id=appointment_id"
```
