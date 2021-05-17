package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"u-fes-2021-team-c/config"
	"u-fes-2021-team-c/database"
	"u-fes-2021-team-c/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name    string
		userReq RegisteruserReq
		want    gin.H
		code    int
	}{
		{
			name: "success",
			userReq: RegisteruserReq{
				Name:     "name",
				Password: "password",
			},
			want: gin.H{"userId": 1},
			code: http.StatusOK,
		},
		{
			name: "failed username is null",
			userReq: RegisteruserReq{
				Password: "password",
			},
			want: gin.H{"err": errors.New("username or password field not null")},
			code: 500,
		},
		{
			name: "failed password is null",
			userReq: RegisteruserReq{
				Name: "name",
			},
			want: gin.H{"err": errors.New("username or password field not null")},
			code: 500,
		},
	}

	sqlHandler, _ := database.NewSqlClient(&config.Config{})
	userRepo := database.UserRepository{
		SqlHandler: *sqlHandler,
	}

	userHandler := NewUserHandler(userRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.userReq)
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/user",
				bytes.NewBuffer(body),
			)
			userHandler.RegisterUser(c)

			var responseBody map[string]interface{}
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.code, response.Code)
			assert.Equal(t, tt.want["usesrId"], responseBody["usesrId"])
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name string
		want []*model.User
		code int
	}{
		{
			name: "success",
			want: []*model.User{
				{
					Id:       1,
					Name:     "name",
					Password: "pass",
				},
			},
			code: 200,
		},
	}

	sqlHandler, _ := database.NewSqlClient(&config.Config{})
	userRepo := database.UserRepository{
		SqlHandler: *sqlHandler,
	}

	userHandler := NewUserHandler(userRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/users",
				nil,
			)
			userHandler.GetAllUsers(c)

			var responseBody []*model.User
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.code, response.Code)
			assert.Equal(t, tt.want, responseBody)
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name string
		want *model.User
		code int
	}{
		{
			name: "success",
			want: &model.User{
				Id:       1,
				Name:     "name",
				Password: "pass",
			},
			code: 200,
		},
	}

	sqlHandler, _ := database.NewSqlClient(&config.Config{})
	userRepo := database.UserRepository{
		SqlHandler: *sqlHandler,
	}

	userHandler := NewUserHandler(userRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/user",
				nil,
			)

			params := c.Request.URL.Query()
			params.Add("id", "1")
			c.Request.URL.RawQuery = params.Encode()

			userHandler.GetUser(c)

			var responseBody *model.User
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.code, response.Code)
			assert.Equal(t, tt.want, responseBody)
		})
	}
}
