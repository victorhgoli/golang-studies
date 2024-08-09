package service

import (
	"estudo-test/pkg/models"
	mock_integration "estudo-test/tests/mocks/integration"
	mock_logger "estudo-test/tests/mocks/logger"
	mock_repository "estudo-test/tests/mocks/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupMocks(ctrl *gomock.Controller) (*mock_logger.MockLogger, *mock_repository.MockUserRepository, *mock_integration.MockInfoTestIntegration) {
	mockLogger := mock_logger.NewMockLogger(ctrl)
	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockIntegration := mock_integration.NewMockInfoTestIntegration(ctrl)
	return mockLogger, mockUserRepo, mockIntegration
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger, mockUserRepo, mockIntegration := setupMocks(ctrl)
	defer ctrl.Finish()

	mockLogger.EXPECT().Infof(gomock.Any(), gomock.Any()).Return().AnyTimes()

	user := &models.User{Nome: "Test User", Email: "test@example.com"}
	mockUserRepo.EXPECT().InsertUser(user).Return(int64(1), nil)

	apiResponse := map[string]interface{}{
		"userId":    1,
		"id":        1,
		"title":     "delectus aut autem",
		"completed": false,
	}
	/*apiResponseBody, _ := json.Marshal(apiResponse)
	mockHTTP.EXPECT().Get("https://jsonplaceholder.typicode.com/todos/1").Return(
		&http.Response{StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewReader(apiResponseBody)),
		}, nil)

	*/

	mockIntegration.EXPECT().GetInfo().Return(apiResponse, nil)
	userService := NewUserService(mockUserRepo, mockLogger, mockIntegration)
	id, err := userService.CreateUser("Test User", "test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)

}
