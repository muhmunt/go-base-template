package history

type Service interface {
	// example service function
	// StoreHistory(request StoreHistoryRequest) (History, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// example service function
// func (s *service) StoreHistory(request StoreHistoryRequest) (History, error) {

// }
