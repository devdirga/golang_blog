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
	Signup(ctx context.Context, post model.User) (model.User, error)
	Signin(ctx context.Context, post model.User) (model.User, error)
	Google(ctx context.Context, post model.User) (model.User, error)
	Me(ctx context.Context, post model.User) (model.User, error)
	UpdateProfile(ctx context.Context, user model.User) (model.User, error)
}

type PostgresUserRespository struct {
	Db *gorm.DB
}

func (r *PostgresUserRespository) Signup(ctx context.Context, user model.User) (model.User, error) {
	if err := r.Db.WithContext(ctx).Where("email=?", user.Username); err == nil {
		return model.User{}, errors.New("user already exist")
	}
	if bytes, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password+config.GetConf().Secret),
		bcrypt.DefaultCost,
	); err != nil {
		return model.User{}, err
	} else {
		user.Password = string(bytes)
	}
	if err := r.Db.WithContext(ctx).Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *PostgresUserRespository) Signin(ctx context.Context, user model.User) (model.User, error) {
	var u model.User
	if err := r.Db.WithContext(ctx).Where("username=?", user.Username).First(&u).Error; err != nil {
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
		"email": u.Username,
		"exp":   util.GetNow().Add(time.Hour * 8640).Unix(), // 1 year
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

func (r *PostgresUserRespository) UpdateProfile(ctx context.Context, user model.User) (model.User, error) {
	var u model.User
	if err := r.Db.WithContext(ctx).Where("id=?", user.ID).Find(&u).Error; err != nil {
		return model.User{}, err
	}
	if u.ID == 0 {
		return model.User{}, errors.New("post not fount")
	}
	u.Name = user.Name
	if err := r.Db.WithContext(ctx).Save(&u).Error; err != nil {
		return model.User{}, err
	}
	return u, nil
}
