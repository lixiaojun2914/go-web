package service

import (
	"code/dao"
	"code/model"
	"code/service/dto"
	"errors"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error

	iUser := m.Dao.GetUserByNameAndPassword(iUserDTO.Name, iUserDTO.Password)
	if iUser.ID == 0 {
		errResult = errors.New("invalid username or password")
	}
	return iUser, errResult
}

func (m *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if m.Dao.CheckUserNameExist(iUserAddDTO.Name) {
		return errors.New("user name exist")
	}
	return m.Dao.AddUser(iUserAddDTO)
}

func (m *UserService) GetUserByID(iCommonIDDTO dto.CommonIDDTO) (model.User, error) {
	return m.Dao.GetUserByID(iCommonIDDTO.ID)
}

func (m *UserService) GetUserList(iUserListDTO dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(iUserListDTO)
}

func (m *UserService) UpdateUser(iUserUpdateDDTO dto.UserUpdateDTO) error {
	if iUserUpdateDDTO.ID == 0 {
		return errors.New("invalid user ID")
	}
	return m.Dao.UpdateUser(iUserUpdateDDTO)
}

func (m *UserService) DeleteUserByID(iCommonIDDTO dto.CommonIDDTO) error {
	return m.Dao.DeleteUserByID(iCommonIDDTO.ID)
}
