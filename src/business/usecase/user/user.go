package user

import (
	userDom "agrokan-backend/src/business/domain/user"
	"agrokan-backend/src/business/entity"
	"agrokan-backend/src/lib/auth"

	"golang.org/x/crypto/bcrypt"
)

type Interface interface {
	Create(user entity.CreateUserParam) (entity.User, error)
	Login(userParam entity.LoginUserParam) (string, error)
	Me(id int) (entity.User, error)
}

type user struct {
	user userDom.Interface
}

func Init(ud userDom.Interface) Interface {
	u := &user{
		user: ud,
	}

	return u
}

func (u *user) Create(userParam entity.CreateUserParam) (entity.User, error) {
	user := entity.User{}

	hash, err := bcrypt.GenerateFromPassword([]byte(userParam.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user = entity.User{
		Email:    userParam.Email,
		Password: string(hash),
	}

	newUser, err := u.user.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *user) Login(userParam entity.LoginUserParam) (string, error) {
	user, err := u.user.Get(entity.SelectUserParam{Email: userParam.Email})
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userParam.Password))
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *user) Me(id int) (entity.User, error) {
	user, err := u.user.Get(entity.SelectUserParam{Id: id})
	if err != nil {
		return user, err
	}

	return user, nil
}
