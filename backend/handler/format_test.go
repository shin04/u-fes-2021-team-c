package handler

// func TestConvertImageToPdf(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		filename string
// 		code     int
// 		isError  bool
// 	}{
// 		{
// 			name:     "success",
// 			filename: "filename",
// 			code:     http.StatusOK,
// 			isError:  false,
// 		},
// 		{
// 			name:     "failed",
// 			filename: "",
// 			code:     500,
// 			isError:  true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		formatHandler := NewFormatHandler()

// 		t.Run(tt.name, func(t *testing.T) {
// 			response := httptest.NewRecorder()
// 			c, _ := gin.CreateTestContext(response)
// 			c.Request, _ = http.NewRequest(
// 				http.MethodGet,
// 				"/convert",
// 				nil,
// 			)
// 			formatHandler.ConvertImageToPdf(c)

// 			assert.Equal(t, tt.code, response.Code)
// 			if !tt.isError {
// 				var responseBody map[string]interface{}
// 				_ = json.Unmarshal(response.Body.Bytes(), &responseBody)
// 				assert.Equal(t, tt.filename, responseBody["filename"])
// 			}
// 		})
// 	}
// }
