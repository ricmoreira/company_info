package controllers

import (
	"company_info/models/request"
	"company_info/models/response"
	"company_info/util/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// stub HeaderService behaviour
type MockHeaderService struct{}

// mocked behaviour for CreateOne
func (ps *MockHeaderService) CreateOne(pReq *mrequest.HeaderCreate) (*mresponse.HeaderCreate, *mresponse.ErrorResponse) {
	// validate request
	err := errors.ValidateRequest(pReq)
	if err != nil {
		return nil, err
	}

	pRes := mresponse.HeaderCreate{}
	pRes.ID = "some-unique-id"

	return &pRes, nil
}

// mocked behaviour for ReadOne
func (ps *MockHeaderService) ReadOne(p *mrequest.HeaderRead) (*mresponse.HeaderRead, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for UpdateOne
func (ps *MockHeaderService) UpdateOne(p *mrequest.HeaderUpdate) (*mresponse.Header, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for DeleteOne
func (ps *MockHeaderService) DeleteOne(p *mrequest.HeaderDelete) (*mresponse.Header, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func (ps *MockHeaderService) CreateMany(*[]*mrequest.HeaderCreate) (*[]*mresponse.HeaderCreate, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func (ps *MockHeaderService) List(*mrequest.ListRequest) (*mresponse.HeaderList, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}
func TestCreateHeaderAction(t *testing.T) {

	// Mock the server

	// Switch to test mode in order to don't get such noisy output
	gin.SetMode(gin.TestMode)

	mhs := &MockHeaderService{}

	hc := HeaderController{
		HeaderService: mhs,
	}

	r := gin.Default()

	r.POST("/api/v1/header", hc.CreateAction)

	// TEST SUCCESS

	// Mock a request
	body := mrequest.HeaderCreate{}
	body.TaxRegistrationNumber = 500200500
	body.BusinessName = "SomeBusinnessName"
	body.FiscalYear = 2018

	jsonValue, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/header", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder in order to inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Do asssertions
	if w.Code != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(w.Body)
		bodyString := string(bodyBytes)

		t.Fatalf("Expected to get status %d but instead got %d\nResponse body:\n%s", http.StatusOK, w.Code, bodyString)
	}
}
