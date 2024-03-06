package users

type userService struct{}

func NewService() *userService {
	return &userService{}
}
