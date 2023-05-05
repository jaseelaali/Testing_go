package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockUsecase "github.com/jaseelali/go-gin-clean-arch/pkg/usecase/mockUseCase"
	"github.com/stretchr/testify/assert"
)

func TestFindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
	UserHandler := NewUserHandler(userUseCase)

	testData := []struct {
		name             string
		email            string
		expectedStatus   int
		expectedResponce Response
		expectedErr      error
		buildStub        func(userUsecase mockUsecase.MockUserUseCase)
	}{
		{
			name:           "success",
			email:          "jaseela@gmail.com",
			expectedStatus: 200,
			expectedResponce: Response{
				ID:      1,
				Email:   "jaseela@gmail.com",
				Name:    "jaseela",
				Surname: "ali",
			},
			expectedErr: nil,
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().FindByEmail(gomock.Any(), "jaseela@gmail.com")
			},
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)
			engine := gin.Default()            //create an engin instance
			recorder := httptest.NewRecorder() //creeate a responce recorder to capture the responce from the request
			engine.GET("users/:email", UserHandler.FindByEmail)
			var body []byte
			body, err := json.Marshal(tt.email) //marshal the user data field into json
			assert.NoError(t, err)
			url := "users/:email"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body)) //create a new http request
			engine.ServeHTTP(recorder, req)                                         //execute the http req
			var actual Response
			err = json.Unmarshal(recorder.Body.Bytes(), &actual) //unmarshal the op
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, recorder.Code)
			assert.Equal(t, tt.expectedResponce.Email, actual.Email)

		})
	}
}
