package user_http

import (
	user_model "amartha-technical-interview/internal/modules/user/model"
	mock_service "amartha-technical-interview/mock/service"
	mock_util "amartha-technical-interview/mock/util"
	util_error "amartha-technical-interview/util/error"
	util_http "amartha-technical-interview/util/http"
	util_http_middleware "amartha-technical-interview/util/http/middleware"
	util_logger "amartha-technical-interview/util/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserHttpHandlerGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_service.NewMockUserService(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	handler := NewUserHttpHandler(serviceMock, jwtMock)

	url := "/user/:id"

	util_logger.InitLogger(gin.TestMode, "test", "./user_http_handler.log")

	app := util_http.NewHTTPServer(gin.TestMode)
	app.Use(
		util_http_middleware.TraceIdAssignmentMiddleware(),
		util_http_middleware.LogHandlerMiddleware(),
		util_http_middleware.ErrorHandlerMiddleware(),
	)

	app.POST(url, handler.GetUserByID)

	userID := uuid.NewString()
	urlWithUserID := fmt.Sprintf("/user/%s", userID)

	expRes := user_model.GetUserResponse{
		ID:      userID,
		Name:    "John Smith",
		Email:   "john@example.com",
		Address: "Sesame Street 123",
	}

	t.Run("should successfully get user", func(t *testing.T) {
		serviceMock.EXPECT().GetUserByID(gomock.Any(), userID).Return(expRes, nil)

		req := httptest.NewRequest(http.MethodPost, urlWithUserID, nil)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.NotEmpty(t, resBody)
	})

	t.Run("should return error when service error", func(t *testing.T) {
		serviceMock.EXPECT().GetUserByID(gomock.Any(), userID).
			Return(user_model.GetUserResponse{}, util_error.NewNotFound(sql.ErrNoRows, "not found"))

		req := httptest.NewRequest(http.MethodPost, urlWithUserID, nil)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})
}
