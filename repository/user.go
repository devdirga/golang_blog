package repository

import (
	"context"
	"errors"
	"goblog/config"
	"goblog/model"
	"goblog/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Signup(ctx context.Context, user model.User) error
	Signin(ctx context.Context, user model.User) (model.User, error)
	Google(ctx context.Context, user model.User) (model.User, error)
	Me(ctx context.Context, user model.User) (model.User, error)
	UpdateProfile(ctx context.Context, user model.User) error
}

type PostgresUserRespository struct {
	Db *gorm.DB
}

func (r *PostgresUserRespository) Signup(ctx context.Context, user model.User) error {
	if err := r.Db.WithContext(ctx).Where("email=?", user.Email); err == nil {
		return errors.New("already exist")
	}
	if bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password+config.GetConf().Secret), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		user.Password = string(bytes)
	}
	if err := r.Db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRespository) Signin(ctx context.Context, user model.User) (model.User, error) {
	var u model.User
	if err := r.Db.WithContext(ctx).Where("email=?", user.Email).First(&u).Error; err != nil {
		return model.User{}, err
	}
	if ok, _ := util.CompareHash(
		u.Password, user.Password, config.GetConf().Secret,
	); !ok {
		return model.User{}, errors.New("password incorrect")
	}
	claims := jwt.MapClaims{
		"id":    strconv.Itoa(u.ID),
		"name":  u.Name,
		"email": u.Email,
		"exp":   util.GetNow().Add(time.Hour * 8640).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GetConf().Secret))
	if err != nil {
		return model.User{}, err
	}
	u.Token = t
	return u, nil
}

func (r *PostgresUserRespository) Google(ctx context.Context, user model.User) (model.User, error) {
	if err := r.Db.WithContext(ctx).Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *PostgresUserRespository) Me(ctx context.Context, user model.User) (model.User, error) {
	var u model.User
	if err := r.Db.WithContext(ctx).Where("id=?", user.ID).Find(&u).Error; err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *PostgresUserRespository) UpdateProfile(ctx context.Context, user model.User) error {
	var u model.User
	if err := r.Db.WithContext(ctx).Where("id=?", user.ID).Find(&u).Error; err != nil {
		return err
	}
	if u.ID == 0 {
		return errors.New("post not fount")
	}
	u.Name = user.Name
	if err := r.Db.WithContext(ctx).Save(&u).Error; err != nil {
		return err
	}
	return nil
}
