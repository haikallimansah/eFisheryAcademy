package usecase

import (
	"clean_arch/entity"
	"clean_arch/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	createUser(user entity.UserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateUser(userRequest entity.UpdateUserRequest, id int) (entity.User, error)
	DeleteUser(id int) error
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (useCase UserUsecase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	users, err := useCase.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Phone:    users.Phone,
	}

	return userRes, nil
}

func (useCase UserUsecase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := useCase.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	userRes := []entity.UserResponse{}
	for _, user := range users {
		appendUser := entity.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		}
		userRes = append(userRes, appendUser)
	}

	return userRes, nil
}

func (useCase UserUsecase) GetUserById(id int) (entity.UserResponse, error) {
	user, err := useCase.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return userRes, nil
}

func (useCase UserUsecase) UpdateUser(userReq entity.UpdateUserRequest, id int) (entity.UserResponse, error) {
	user, err := useCase.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}

	user.Username = userReq.Username
	user.Email = userReq.Email
	user.Phone = userReq.Phone

	copier.CopyWithOption(&user, &userReq, copier.Option{IgnoreEmpty: true})
	user, err = useCase.userRepository.Update(user)
	userRes := entity.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, nil

	// updatedUser := entity.User{
	// 	ID:       user.ID,
	// 	Username: user.Username,
	// 	Email:    user.Email,
	// 	Phone:    user.Phone,
	// }

	// user, err = useCase.userRepository.Update(updatedUser)
	// if err != nil {
	// 	return entity.UserResponse{}, err
	// }

	// userRes := entity.UserResponse{
	// 	ID:       user.ID,
	// 	Username: user.Username,
	// 	Email:    user.Email,
	// 	Phone:    user.Phone,
	// }

	// return userRes, nil
}

func (useCase UserUsecase) DeleteUser(id int) error {
	_, err := useCase.userRepository.FindById(id)
	if err != nil {
		return err
	}

	err = useCase.userRepository.Delete(id)
	return err
}
