package storage

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/user/port"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type UserRepoTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo port.Repo
}

func (suite *UserRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	suite.Require().NoError(err)
	suite.Require().NotNil(db)
	suite.Require().NotNil(mock)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	suite.Require().NoError(err)

	suite.db = gormDB
	suite.mock = mock
	suite.repo = NewUserRepo(gormDB)
}

func (suite *UserRepoTestSuite) TearDownTest() {
	sqlDB, err := suite.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func (suite *UserRepoTestSuite) TestCreateUser() {
	user := domain.User{
		CreatedAT: time.Now(),
		Email:     "mojfajd@fjaf.com",
		Password:  "fjahlfda",
		NatId:     "4586633669",
		Role:      domain.AdminRole,
	}
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT INTO `users` .*").
		WithArgs(sqlmock.AnyArg(), user.Email, user.Password, user.NatId, user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	userDbID, err := suite.repo.Create(context.Background(), user)
	suite.Require().NoError(err)
	assert.Equal(suite.T(), domain.UserDbID(1), userDbID, "User ID should match")

	err = suite.mock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *UserRepoTestSuite) TestUpdateUser() {
	updatedUser := domain.User{
		Email:    "updated@gmail.com",
		Password: "newpassword",
		NatId:    "1234567890",
		Role:     domain.UserRole,
	}
	userStorage := mapper.UserDomain2Storage(updatedUser)

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("UPDATE `users` SET .* WHERE id = \\?").
		WithArgs(updatedUser.Email, updatedUser.Password, updatedUser.NatId, updatedUser.Role, userStorage.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	userDbID, err := suite.repo.Update(context.Background(), updatedUser)

	suite.Require().NoError(err)
	assert.Equal(suite.T(), domain.UserDbID(userStorage.ID), userDbID)
	err = suite.mock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *UserRepoTestSuite) TestFindWithUserDbID() {
	userDbID := domain.UserDbID(1)
	storedUser := types.User{
		ID:        uint(userDbID),
		Email:     "test@domain.com",
		Password:  "password",
		NatID:     "1234567890",
		Role:      uint8(domain.UserRole),
		CreatedAt: time.Now(),
	}

	suite.mock.ExpectQuery("SELECT .* FROM `users` WHERE id = ?").
		WithArgs(userDbID, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "nat_id", "role", "created_at"}).
			AddRow(storedUser.ID, storedUser.Email, storedUser.Password, storedUser.NatID, storedUser.Role, storedUser.CreatedAt))

	user, err := suite.repo.FindWithUserDbID(context.Background(), userDbID)

	suite.Require().NoError(err)
	assert.Equal(suite.T(), mapper.UserStorage2Domain(storedUser), user)
	err = suite.mock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *UserRepoTestSuite) TestDeleteWithUserDbID() {
	userDbID := domain.UserDbID(1)

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("DELETE FROM `users` WHERE id = ?").
		WithArgs(userDbID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mock.ExpectCommit()

	err := suite.repo.DeleteWithUserDbId(context.Background(), userDbID)

	suite.Require().NoError(err)
	err = suite.mock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
