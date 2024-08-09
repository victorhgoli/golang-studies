package service

import (
	logger "estudo-test/infra/logger"
	"estudo-test/integration"
	"estudo-test/internal/repository"
	"estudo-test/pkg/models"
)

type userService struct {
	Repo        repository.UserRepository
	Log         logger.Logger
	InfoTestInt integration.InfoTestIntegration
}

type UserService interface {
	CreateUser(nome, email string) (int64, error)
	GetUsersById(id int64) (*models.User, error)
}

func NewUserService(repo repository.UserRepository, log logger.Logger, infoTesteIntegration integration.InfoTestIntegration) UserService {
	return &userService{Repo: repo, Log: log, InfoTestInt: infoTesteIntegration}
}

func (s *userService) CreateUser(nome, email string) (int64, error) {
	user := &models.User{Nome: nome, Email: email}
	id, er := s.Repo.InsertUser(user)

	s.Log.Infof("Inseriu usuario %v", user)

	data, err := s.InfoTestInt.GetInfo()
	if err != nil {
		return 0, nil
	}

	s.Log.Infof("Dados da API: %v", data)

	return id, er
}

func (s *userService) GetUsersById(id int64) (*models.User, error) {

	return s.Repo.GetUserById(id)
}
