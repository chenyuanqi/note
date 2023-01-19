package demo

import (
	"errors"
	"mlj/app/model/demo"

	"github.com/asaskevich/govalidator"
)

type UserCreate struct {
	Request  UserCreateParams
	Response map[string]interface{}
}

type UserCreateParams struct {
	Name     string `valid:"required" form:"name"`
	Email    string `valid:"required" form:"email"`
	Password string `valid:"required" form:"password"`
}

func (u *UserCreate) Create() error {
	if _, err := govalidator.ValidateStruct(u.Request); err != nil {
		return err
	}

	user := &demo.Users{
		Name:     u.Request.Name,
		Email:    u.Request.Email,
		Password: u.Request.Password,
	}
	user.Create()

	return nil
}

type UserUpdate struct {
	Request  UserUpdateParams
	Response map[string]interface{}
}

type UserUpdateParams struct {
	ID    int    `valid:"required" form:"id"`
	Name  string `valid:"required" form:"name"`
	Email string `valid:"required" form:"email"`
}

func (u *UserUpdate) Update() error {
	if _, err := govalidator.ValidateStruct(u.Request); err != nil {
		return err
	}

	user := &demo.Users{}
	user.FindByID(int64(u.Request.ID))
	if user.ID == 0 {
		return errors.New("user not found")
	}

	user.Name = u.Request.Name
	user.Email = u.Request.Email
	if user.Save() == 0 {
		return errors.New("user update failed")
	}

	return nil
}

type UserDelete struct {
	Request  UserDeleteParams
	Response map[string]interface{}
}

type UserDeleteParams struct {
	ID int `valid:"required" form:"id"`
}

func (u *UserDelete) Delete() error {
	if _, err := govalidator.ValidateStruct(u.Request); err != nil {
		return err
	}

	user := &demo.Users{}
	user.FindByID(int64(u.Request.ID))
	if user.ID == 0 {
		return errors.New("user not found")
	}

	if user.Delete() == 0 {
		return errors.New("user delete failed")
	}

	return nil
}

type UserDetail struct {
	Request  UserDetailParams
	Response struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
}

type UserDetailParams struct {
	ID int `valid:"required" form:"id"`
}

func (u *UserDetail) Detail() error {
	if _, err := govalidator.ValidateStruct(u.Request); err != nil {
		return err
	}

	user := &demo.Users{}
	user.FindByID(int64(u.Request.ID))
	if user.ID == 0 {
		return errors.New("user not found")
	}

	u.Response.ID = int(user.ID)
	u.Response.Name = user.Name
	u.Response.Email = user.Email
	u.Response.CreatedAt = user.CreatedAtDateTime()
	u.Response.UpdatedAt = user.UpdatedAtDateTime()

	return nil
}
