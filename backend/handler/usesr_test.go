package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"u-fes-2021-team-c/config"
	"u-fes-2021-team-c/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name                string
		userReq             RegisteruserReq
		fakeRegisterNewUser func(string, string) (int, error)
		want                gin.H
	}{
		{
			name: "success",
			userReq: RegisteruserReq{
				Name:     "name",
				Password: "password",
			},
			fakeRegisterNewUser: func(string, string) (int, error) {
				return 1, nil
			},
			want: gin.H{"userId": 1},
		},
		{
			name: "failed username is null",
			userReq: RegisteruserReq{
				Password: "password",
			},
			fakeRegisterNewUser: nil,
			want:                gin.H{"err": errors.New("username or password field not null")},
		},
	}

	sqlHandler, _ := database.NewSqlClient(&config.Config{})
	userHandler := NewUserHandler(*sqlHandler)

	for _, tt := range tests {
		body, _ := json.Marshal(tt.userReq)
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/user",
			bytes.NewBuffer(body),
		)
		userHandler.RegisterUser(c)

		// err := json.Unmarshal(response.Body.Bytes(), &product)
		// assert.EqualValues(t, http.StatusOK, response.Code)
		fmt.Println("CODE TYPE", reflect.TypeOf(http.StatusOK))
		fmt.Println("RESPONSE BODY", response.Body)
		assert.Equal(t, tt.want, response.Body)
	}
}
