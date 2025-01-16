package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RedColdHearted/go-demo-exam/app"
	"github.com/RedColdHearted/go-demo-exam/models"
	"github.com/RedColdHearted/go-demo-exam/routers/api"
	"github.com/RedColdHearted/go-demo-exam/tests"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ApplicationSuite struct {
	suite.Suite
	App *app.Application
}

func (as *ApplicationSuite) SetupSuite() {
	database, err := tests.SetUpDB()
	if err != nil {
		panic(err)
	}
	as.App = &app.Application{
		R:  gin.Default(),
		DB: database,
	}
	api.SetupV1Api(as.App)

	var input models.CreateReservation
	gofakeit.Struct(&input)
	reservation := models.Reservation{
		BaseReservation: input.BaseReservation,
	}
	as.App.DB.Create(&reservation)
}

func (as *ApplicationSuite) SetupTest() {
}

func (as *ApplicationSuite) TearDownTest() {
}

func (as *ApplicationSuite) TearDownSuite() {
	tests.TearDownDB(as.App.DB)
}

func (as *ApplicationSuite) TestCreateReservation() {
	assert := assert.New(as.T())
	var input models.CreateReservation
	gofakeit.Struct(&input)
	jsonData, err := json.Marshal(input)
	if err != nil {
		as.T().Log(err)
	}

	request, err := http.NewRequest(
		"POST",
		"/api/v1/reservation",
		bytes.NewBuffer(jsonData),
	)
	assert.NoError(err)
	recorder := httptest.NewRecorder()
	as.App.R.ServeHTTP(recorder, request)
    var response map[string]any
    err = json.NewDecoder(recorder.Body).Decode(&response)
	assert.NoError(err)
	assert.Equal("reservation created", response["message"])
	assert.NoError(err)
	assert.Equal(http.StatusCreated, recorder.Code)
    as.T().Log(response["object"])
}

func TestApiV1(t *testing.T) {
	suite.Run(t, new(ApplicationSuite))
}
