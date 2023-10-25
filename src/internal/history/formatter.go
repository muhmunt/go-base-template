package history

import "go-base-template/src/model"

// example formatter
type HistoryFormatter struct {
	PracticingID  string `json:"practicing_id"`
	AppointmentID *int   `json:"appointment_id"`
	CreatedAt     *int   `json:"created_at"`
	UpdatedAt     *int   `json:"updated_at"`
}

func FormatHistory(history model.History) HistoryFormatter {
	formatter := HistoryFormatter{
		PracticingID:  history.PracticingID,
		AppointmentID: history.AppointmentID,
		CreatedAt:     history.CreatedAt,
		UpdatedAt:     history.UpdatedAt,
	}

	return formatter
}

func FormatHistories(histories []model.History) []HistoryFormatter {
	historiesFormatter := []HistoryFormatter{}

	for _, history := range histories {
		historyFormatter := FormatHistory(history)
		historiesFormatter = append(historiesFormatter, historyFormatter)
	}

	return historiesFormatter
}
