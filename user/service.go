package user

// Service represent the users' service
type Service interface {
	Login(email, password string) (token string, err error)
	Register(user *User) error
}

type service struct {
	repo Repository
}

// NewService create an object that represent the Service interface
func NewService(userRepo Repository) Service {
	return service{
		repo: userRepo,
	}
}

func (us service) Login(email, password string) (string, error) {
	return "", nil
}

func (us service) Register(user *User) error {
	if err := us.repo.Create(user); err != nil {
		return err
	}
	return nil
}
