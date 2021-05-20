package handler

import (
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

func TestGetAllStudentInfo(t *testing.T) {
	tests := []struct {
		name                  string
		fakeGetAllStudentInfo func() ([]*model.StudentInfo, error)
		want                  []*model.StudentInfo
		code                  int
		isError               bool
		wantError             error
	}{
		{
			name: "success",
			fakeGetAllStudentInfo: func() ([]*model.StudentInfo, error) {
				return []*model.StudentInfo{
					{
						Id:            1,
						UseId:         1,
						Name:          "name",
						StudentNumber: 111111111,
					},
				}, nil
			},
			want: []*model.StudentInfo{
				{
					Id:            1,
					UseId:         1,
					Name:          "name",
					StudentNumber: 111111111,
				},
			},
			code:    200,
			isError: false,
		},
		{
			name: "failed get all StudentInfos 1",
			fakeGetAllStudentInfo: func() ([]*model.StudentInfo, error) {
				return nil, errors.New("get all student info error")
			},
			code:      500,
			isError:   true,
			wantError: errors.New("get all student info error"),
		},
		{
			name: "failed get all StudentInfos 2",
			fakeGetAllStudentInfo: func() ([]*model.StudentInfo, error) {
				return []*model.StudentInfo{}, nil
			},
			code:      500,
			isError:   true,
			wantError: errors.New("studentinfo not found"),
		},
	}

	for _, tt := range tests {
		studentInfoRepo := testutils.FakeStudentInfoRepository{
			FakeGetAllStudentInfo: tt.fakeGetAllStudentInfo,
		}

		studentInfoHandler := NewStudentinfoHandler(&studentInfoRepo)

		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/student_infos",
				nil,
			)
			studentInfoHandler.GetAllStudentInfo(c)

			assert.Equal(t, tt.code, response.Code)
			if !tt.isError {
				var responseBody []*model.StudentInfo
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
