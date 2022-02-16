package usecases

import (
	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/database/mock"
)

var repo database.Repository
var fakeRepo mock.FakeRepository
