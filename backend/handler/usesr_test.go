package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"u-fes-2021-team-c/model"
	"u-fes-2021-team-c/testutils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name           string
		userReq        RegisteruserReq
		fakeCreateUser func(user *model.User) (int, error)
		want           gin.H
		code           int
	}{
		{
			name: "success",
			userReq: RegisteruserReq{
				Name:     "name",
				Password: "password",
			},
			fakeCreateUser: func(user *model.User) (int, error) {
				return 1, nil
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
		{
			name: "failed register new user 1",
			userReq: RegisteruserReq{
				Name:     "name",
				Password: "password",
			},
			fakeCreateUser: func(user *model.User) (int, error) {
				return -1, errors.New("failed register new user")
			},
			want: gin.H{"err": errors.New("failed register new user")},
			code: 500,
		},
		{
			name: "failed register new user 2",
			userReq: RegisteruserReq{
				Name:     "name",
				Password: "password",
			},
			fakeCreateUser: func(user *model.User) (int, error) {
				return -1, nil
			},
			want: gin.H{"err": errors.New("regisster usesr failed")},
			code: 500,
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeCreateUser: tt.fakeCreateUser,
		}

		userHandler := NewUserHandler(userRepo)

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
		name            string
		fakeGetAllUsers func() ([]*model.User, error)
		want            []*model.User
		code            int
		isError         bool
		wantError       error
	}{
		{
			name: "success",
			fakeGetAllUsers: func() ([]*model.User, error) {
				return []*model.User{
					{
						Id:       1,
						Name:     "name",
						Password: "pass",
					},
				}, nil
			},
			want: []*model.User{
				{
					Id:       1,
					Name:     "name",
					Password: "pass",
				},
			},
			code:    200,
			isError: false,
		},
		{
			name: "failed get all users 1",
			fakeGetAllUsers: func() ([]*model.User, error) {
				return nil, errors.New("get all user error")
			},
			code:      500,
			isError:   true,
			wantError: errors.New("get all user error"),
		},
		{
			name: "failed get all users 2",
			fakeGetAllUsers: func() ([]*model.User, error) {
				return []*model.User{}, nil
			},
			code:      500,
			isError:   true,
			wantError: errors.New("users not found"),
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeGetAllUsers: tt.fakeGetAllUsers,
		}

		userHandler := NewUserHandler(userRepo)

		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/users",
				nil,
			)
			userHandler.GetAllUsers(c)

			assert.Equal(t, tt.code, response.Code)
			if !tt.isError {
				var responseBody []*model.User
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.want, responseBody)
			} else {
				var responseBody map[string]interface{}
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.wantError.Error(), responseBody["err"])
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name        string
		fakeGetUser func(userId int) (*model.User, error)
		want        *model.User
		code        int
		isError     bool
		wantError   error
	}{
		{
			name: "success",
			fakeGetUser: func(userId int) (*model.User, error) {
				return &model.User{
					Id:       1,
					Name:     "name",
					Password: "pass",
				}, nil
			},
			want: &model.User{
				Id:       1,
				Name:     "name",
				Password: "pass",
			},
			code:    200,
			isError: false,
		},
		{
			name: "failed get user by id",
			fakeGetUser: func(userId int) (*model.User, error) {
				return nil, errors.New("get user error")
			},
			code:      500,
			isError:   true,
			wantError: errors.New("get user error"),
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeGetUserById: tt.fakeGetUser,
		}

		userHandler := NewUserHandler(userRepo)

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

			assert.Equal(t, tt.code, response.Code)
			if !tt.isError {
				var responseBody *model.User
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.want, responseBody)
			} else {
				var responseBody map[string]interface{}
				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.wantError.Error(), responseBody["err"])
			}
		})
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name                         string
		userReq                      LoginReq
		fakeGetUSerByNameAndPassword func(name string, password string) (*model.User, error)
		want                         gin.H
		code                         int
	}{
		{
			name: "success",
			userReq: LoginReq{
				Name:     "name",
				Password: "password",
			},
			fakeGetUSerByNameAndPassword: func(name string, password string) (*model.User, error) {
				return &model.User{
					Id:       1,
					Name:     "name",
					Password: "pass",
				}, nil
			},
			want: gin.H{"userId": 1},
			code: http.StatusOK,
		},
		{
			name: "failed username is null",
			userReq: LoginReq{
				Password: "password",
			},
			want: gin.H{"err": errors.New("username or password field not null")},
			code: 500,
		},
		{
			name: "failed password is null",
			userReq: LoginReq{
				Name: "name",
			},
			want: gin.H{"err": errors.New("username or password field not null")},
			code: 500,
		},
		{
			name: "failed register new user 1",
			userReq: LoginReq{
				Name:     "name",
				Password: "password",
			},
			fakeGetUSerByNameAndPassword: func(name string, password string) (*model.User, error) {
				return nil, errors.New("record not found")
			},
			want: gin.H{"err": errors.New("record not found")},
			code: 500,
		},
	}

	for _, tt := range tests {
		userRepo := testutils.FakeUserRepository{
			FakeGetUserByNameAndPasssword: tt.fakeGetUSerByNameAndPassword,
		}

		userHandler := NewUserHandler(userRepo)

		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.userReq)
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/login",
				bytes.NewBuffer(body),
			)
			userHandler.Login(c)

			var responseBody map[string]interface{}
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.code, response.Code)
			assert.Equal(t, tt.want["usesrId"], responseBody["usesrId"])
		})
	}
}
