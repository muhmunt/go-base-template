package model

// example entity
type History struct {
	Id            string `json:"_id"`
	PracticingID  string `json:"practicing_id"`
	AppointmentID *int   `json:"appointment_id"`
	CreatedAt     *int   `json:"created_at"`
	UpdatedAt     *int   `json:"updated_at"`
}
