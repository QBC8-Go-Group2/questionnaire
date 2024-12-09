package storage

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	optionDomain "github.com/QBC8-Go-Group2/questionnaire/internal/option/domain"
	domain2 "github.com/QBC8-Go-Group2/questionnaire/internal/question/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/domain"
	"github.com/QBC8-Go-Group2/questionnaire/internal/response/port"
	userDomian "github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/mapper"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type ResponseRepoTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo port.Repo
}

func (suite *ResponseRepoTestSuite) SetupTest() {
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
	suite.repo = NewResponseRepo(gormDB)
}

func (suite *ResponseRepoTestSuite) TearDownTest() {
	sqlDB, err := suite.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func (suite *ResponseRepoTestSuite) TestCreateResponse() {
	response := domain.Response{
		Type:       domain2.QuestionTypeAnatomical,
		UserID:     userDomian.UserDbID(3),
		QuestionID: domain2.QuestionDbID(1),
		Data:       "template data",
		OptionID:   optionDomain.OptionID(4),
		CreatedAt:  time.Now(),
	}
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT INTO `responses` \\(`type`,`user_id`,`question_id`,`data`,`option_id`,`created_at`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?\\)").
		WithArgs(response.Type, response.UserID, response.QuestionID, response.Data, response.OptionID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	responseDbId, err := suite.repo.Create(context.Background(), response)
	suite.Require().NoError(err)
	assert.Equal(suite.T(), responseDbId, responseDbId, "User ID should match")

	err = suite.mock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *ResponseRepoTestSuite) TestFindByIdResponse() {
	responseID := domain.ResponseID(1)
	storedResponse := types.Response{
		ID:         uint(responseID),
		Type:       2,
		UserID:     3,
		QuestionID: 4,
		Data:       "retrieved data",
		OptionID:   5,
	}

	suite.mock.ExpectQuery("SELECT .* FROM `responses` WHERE id = \\?").
		WithArgs(responseID, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "type", "user_id", "question_id", "data", "option_id", "created_at",
		}).AddRow(
			storedResponse.ID, storedResponse.Type, storedResponse.UserID,
			storedResponse.QuestionID, storedResponse.Data, storedResponse.OptionID,
			storedResponse.CreatedAt,
		))

	// Act
	response, err := suite.repo.FindById(context.Background(), responseID)

	// Assert
	suite.Require().NoError(err)
	assert.Equal(suite.T(), mapper.ResponseStorage2Domain(storedResponse), response)
	err = suite.mock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func (suite *ResponseRepoTestSuite) TestDeleteResponse() {
	// Arrange
	response := domain.Response{
		ID: 1,
	}

	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("DELETE FROM `responses` WHERE id = \\?").
		WithArgs(response.ID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	// Act
	err := suite.repo.Delete(context.Background(), response)

	// Assert
	suite.Require().NoError(err)
	err = suite.mock.ExpectationsWereMet()
	suite.Require().NoError(err)
}

func TestResponseRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ResponseRepoTestSuite))
}
