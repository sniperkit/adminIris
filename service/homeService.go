package service

// NewHomeService ...
func NewHomeService() HomeService {
	return &homeService{}
}

// HomeService ...
type HomeService interface {
	Greeting() string
}

type homeService struct{}

func (s *homeService) Greeting() string {
	return "Hello! API"
}
