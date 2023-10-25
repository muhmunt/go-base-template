package history

// example request
type StoreHistoryRequest struct {
	PracticingID  string `json:"practicing_id"`
	AppointmentID *int   `json:"appointment_id"`
	CreatedAt     *int   `json:"created_at" binding:"required"`
	UpdatedAt     *int   `json:"updated_at"`
}
