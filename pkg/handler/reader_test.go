package handler

import (
	"bytes"
	nameenrich "name-enrich"
	"name-enrich/pkg/service"
	mock_service "name-enrich/pkg/service/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_read(t *testing.T) {
	type mockBehavior func(s *mock_service.MockEnrich, p nameenrich.Person)

	testTable := []struct {
		name                string
		inputBody           string
		inputPerson         nameenrich.Person
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: "{\"name\":\"Dmitry\", \"surname\":\"Ushakov\"}",
			inputPerson: nameenrich.Person{
				Surname: "Ushakov",
				Name:    "Dmitry",
			},
			mockBehavior: func(s *mock_service.MockEnrich, p nameenrich.Person) {
				s.EXPECT().EnrichAge(p).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "{\"id\":\"0\", \"name\":\"Dmitry\", \"surname\":\"Ushakov\"}",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			enrich := mock_service.NewMockEnrich(c)
			testCase.mockBehavior(enrich, testCase.inputPerson)

			services := &service.Service{Enrich: enrich}
			handler := NewHandler(services)

			r := gin.New()

			r.POST("/", handler.takePerson)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedStatusCode, w.Code.String())
		})
	}
}
