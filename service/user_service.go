package service

import (
	"code/dao"
	"code/global"
	"code/global/constants"
	"code/model"
	"code/service/dto"
	"code/utils"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
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

func GenerateAndCacheLoginUserTokenToRedis(nUserID uint, stUserName string) (string, error) {
	token, err := utils.GenerateToken(nUserID, stUserName)
	if err == nil {
		iUserID := strconv.Itoa(int(nUserID))
		stRedisUserIDKey := strings.Replace(constants.LoginUserTokenRedisKey, "{ID}", iUserID, -1)
		err = global.RedisClient.Set(stRedisUserIDKey, token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
	}
	return token, err
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string

	iUser, err := m.Dao.GetUserByName(iUserDTO.Name)
	// 用户名或密码不正确
	if err != nil || !utils.CompareHashAndPassword(iUser.Password, iUserDTO.Password) {
		errResult = errors.New("invalid username or password")
		return model.User{}, "", errResult
	}
	// 登陆成功，生成token
	token, err = GenerateAndCacheLoginUserTokenToRedis(iUser.ID, iUser.Name)
	if err != nil {
		errResult = fmt.Errorf("generate token error: %s", err.Error())
		return model.User{}, "", errResult
	}

	return iUser, token, errResult
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
