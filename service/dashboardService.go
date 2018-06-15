package service

// NewDashBoardService ...
func NewDashBoardService() DashBoardService {
	return &dashBoardService{}
}

// DashBoardService ...
type DashBoardService interface {
	Greeting() string
}

type dashBoardService struct{}

func (s *dashBoardService) Greeting() string {
	return "Hello! API"
}
