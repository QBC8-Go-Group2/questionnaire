package storage

import (
	"context"

	"github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserDbID, error) {
	userStorage := mapper.UserDomain2Storage(userDomain)
	return domain.UserDbID(userStorage.ID), u.db.Table("users").WithContext(ctx).Create(&userStorage).Error
}

func (u *userRepo) Update(ctx context.Context, user domain.User) (domain.UserDbID, error) {
	userStorage := mapper.UserDomain2Storage(user)
	return domain.UserDbID(userStorage.ID), u.db.Table("users").WithContext(ctx).Updates(&userStorage).Error
}

func (u *userRepo) FindWithUserID(ctx context.Context, userId domain.UserID) (domain.User, error) {
	var user types.User
	err := u.db.Table("users").WithContext(ctx).Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return mapper.UserStorage2Domain(user), nil
}

func (u *userRepo) FindWithUserDbID(ctx context.Context, userDbId domain.UserDbID) (domain.User, error) {
	var user types.User
	err := u.db.Table("users").WithContext(ctx).Where("id = ?", userDbId).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return mapper.UserStorage2Domain(user), nil
}

func (u *userRepo) FindWithEmail(ctx context.Context, email string) (domain.User, error) {
	var user types.User
	err := u.db.Table("users").WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return mapper.UserStorage2Domain(user), nil
}

func (u *userRepo) DeleteWithUserID(ctx context.Context, user domain.UserID) error {
	return u.db.Table("users").WithContext(ctx).Delete(&types.User{}, "user_id = ?", user).Error
}

func (u *userRepo) DeleteWithUserDbId(ctx context.Context, user domain.UserDbID) error {
	return u.db.Table("users").WithContext(ctx).Where("id = ?", user).Delete(&types.User{}).Error
}
