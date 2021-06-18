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

func TestRegisterStudentInfo(t *testing.T) {
	tests := []struct {
		name                  string
		req                   RegisterStudentInfoReq
		fakeCreateStudentInfo func(studentInfo *model.StudentInfo) (int, error)
		want                  gin.H
		wantErr               error
		code                  int
		isErr                 bool
	}{
		{
			name: "success",
			req: RegisterStudentInfoReq{
				UserId:        1,
				Name:          "name",
				StudentNumber: 11111111,
			},
			fakeCreateStudentInfo: func(studentInfo *model.StudentInfo) (int, error) {
				return 1, nil
			},
			want:  gin.H{"id": 1},
			code:  http.StatusOK,
			isErr: false,
		},
		{
			name: "failed student name is null",
			req: RegisterStudentInfoReq{
				UserId:        1,
				StudentNumber: 11111111,
			},
			wantErr: errors.New("student name field not null"),
			code:    500,
			isErr:   true,
		},
		{
			name: "failed register new student info 1",
			req: RegisterStudentInfoReq{
				UserId:        1,
				Name:          "name",
				StudentNumber: 11111111,
			},
			fakeCreateStudentInfo: func(studentInfo *model.StudentInfo) (int, error) {
				return -1, errors.New("")
			},
			wantErr: errors.New("register new student info failed"),
			code:    500,
			isErr:   true,
		},
		{
			name: "failed register new student info 2",
			req: RegisterStudentInfoReq{
				UserId:        1,
				Name:          "name",
				StudentNumber: 11111111,
			},
			fakeCreateStudentInfo: func(studentInfo *model.StudentInfo) (int, error) {
				return -1, nil
			},
			wantErr: errors.New("register new student info failed"),
			code:    500,
			isErr:   true,
		},
	}

	for _, tt := range tests {
		repo := testutils.FakeStudentInfoRepository{
			FakeCreateStudentInfo: tt.fakeCreateStudentInfo,
		}

		studentInfoHandler := NewStudentinfoHandler(&repo)

		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.req)
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/student_info",
				bytes.NewBuffer(body),
			)
			studentInfoHandler.RegisterStudentInfo(c)

			var responseBody map[string]interface{}
			_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.code, response.Code)

			if tt.isErr {
				assert.Equal(t, tt.wantErr.Error(), responseBody["err"])
			} else {
				val, _ := responseBody["id"]
				id := int(val.(float64))
				assert.Equal(t, tt.want["id"], id)
			}
		})
	}
}

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

func TestGetStudentInfoByUserID(t *testing.T) {
	tests := []struct {
		name               string
		fakeGetStudentInfo func(usesrId int) (*model.StudentInfo, error)
		want               *model.StudentInfo
		code               int
		isError            bool
		wantError          error
	}{
		{
			name: "success",
			fakeGetStudentInfo: func(userId int) (*model.StudentInfo, error) {
				return &model.StudentInfo{
					Id:            1,
					UseId:         1,
					Name:          "name",
					StudentNumber: 11111111,
				}, nil
			},
			want: &model.StudentInfo{
				Id:            1,
				UseId:         1,
				Name:          "name",
				StudentNumber: 11111111,
			},
			code:    200,
			isError: false,
		},
		{
			name: "failed get user by id",
			fakeGetStudentInfo: func(userId int) (*model.StudentInfo, error) {
				return nil, errors.New("get user error")
			},
			code:      500,
			isError:   true,
			wantError: errors.New("get user error"),
		},
	}

	for _, tt := range tests {
		repo := testutils.FakeStudentInfoRepository{
			FakeGetStudentInfoByUSesrID: tt.fakeGetStudentInfo,
		}

		studentInfoHandler := NewStudentinfoHandler(&repo)

		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/student_info",
				nil,
			)

			params := c.Request.URL.Query()
			params.Add("id", "1")
			c.Request.URL.RawQuery = params.Encode()

			studentInfoHandler.GetStudentInfoByUserId(c)

			assert.Equal(t, tt.code, response.Code)
			if !tt.isError {
				var responseBody *model.StudentInfo
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
