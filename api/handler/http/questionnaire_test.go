package http

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type QuestionnaireTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (suite *QuestionnaireTestSuite) SetupTest() {

}

func (suite *QuestionnaireTestSuite) TearDownTest() {

}
