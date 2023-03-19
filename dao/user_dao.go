package dao

import (
	"code/model"
	"code/service/dto"
	"github.com/jinzhu/copier"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", stUserName, stPassword).Find(&iUser)
	return iUser
}

func (m *UserDao) CheckUserNameExist(stUserName string) bool {
	var nTotal int64
	m.Orm.Model(&model.User{}).Where("name=?", stUserName).Count(&nTotal)
	return nTotal > 0
}

func (m *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	err := copier.Copy(&iUser, iUserAddDTO)
	if err != nil {
		return err
	}

	err = m.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}

func (m *UserDao) GetUserByID(ID uint) (model.User, error) {
	var iUser model.User
	err := m.Orm.First(&iUser, ID).Error
	return iUser, err
}

func (m *UserDao) GetUserList(iUserListDTO dto.UserListDTO) ([]model.User, int64, error) {
	var giUserList []model.User
	var nTotal int64
	err := m.Orm.Model(&model.User{}).
		Scopes(Paginate(iUserListDTO.Paginate)).
		Find(&giUserList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error
	return giUserList, nTotal, err
}

func (m *UserDao) UpdateUser(iUserUpdateDTO dto.UserUpdateDTO) error {
	var iUser model.User
	m.Orm.First(&iUser, iUserUpdateDTO.ID)
	err := copier.Copy(&iUser, &iUserUpdateDTO)
	if err != nil {
		return err
	}
	return m.Orm.Save(&iUser).Error
}

func (m *UserDao) DeleteUserByID(ID uint) error {
	return m.Orm.Delete(&model.User{}, ID).Error
}
