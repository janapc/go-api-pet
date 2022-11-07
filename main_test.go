package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-pet/controllers"
	"go-api-pet/database"
	"go-api-pet/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routers := gin.Default()
	return routers
}

func AddPetMock() {
	pet := models.Pet{Name: "Gamora", Observation: "", Breed: "Husky Siberiano", Size: "MEDIUM"}
	database.DB.Create(&pet)
	ID = int(pet.ID)
}

func RemovePetMock() {
	var pet models.Pet
	database.DB.Delete(&pet, ID)
}

func TestVerifyHealthCheckApi(t *testing.T) {
	r := SetupRouters()
	r.GET("/api/v1/pets/health", controllers.HealthCheck)
	req, _ := http.NewRequest("GET", "/api/v1/pets/health", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	messageError := fmt.Sprintf("Status Code receive: %d; expected: %d", response.Code, http.StatusOK)
	assert.Equal(t, http.StatusOK, response.Code, messageError)
}

func TestListAllPets(t *testing.T) {
	database.Connection()
	AddPetMock()
	defer RemovePetMock()
	r := SetupRouters()
	r.GET("/api/v1/pets", controllers.ListAll)
	request, _ := http.NewRequest("GET", "/api/v1/pets", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	messageError := fmt.Sprintf("Status Code receive: %d; expected: %d", response.Code, http.StatusOK)
	assert.Equal(t, http.StatusOK, response.Code, messageError)
	var petsResponse []models.Pet
	json.Unmarshal(response.Body.Bytes(), &petsResponse)
	assert.Len(t, petsResponse, 1, "Length receive: %d; expected: %d", len(petsResponse), 1)
}

func TestFindOnePetById(t *testing.T) {
	database.Connection()
	AddPetMock()
	defer RemovePetMock()
	r := SetupRouters()
	r.GET("/api/v1/pets/:id", controllers.FindOneById)
	path := "/api/v1/pets/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	messageError := fmt.Sprintf("Status Code receive: %d; expected: %d", response.Code, http.StatusOK)
	assert.Equal(t, http.StatusOK, response.Code, messageError)
	var petResponse models.Pet
	json.Unmarshal(response.Body.Bytes(), &petResponse)
	assert.Equal(t, "Gamora", petResponse.Name, "Name should be equals")
	assert.Equal(t, "", petResponse.Observation, "Observation should be equals")
	assert.Equal(t, "Husky Siberiano", petResponse.Breed, "Breed should be equals")
	assert.Equal(t, "MEDIUM", petResponse.Size, "Size should be equals")
}

func TestNotFindOnePetById(t *testing.T) {
	database.Connection()
	AddPetMock()
	defer RemovePetMock()
	r := SetupRouters()
	r.GET("/api/v1/pets/:id", controllers.FindOneById)
	path := "/api/v1/pets/0"
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	messageError := fmt.Sprintf("Status Code receive: %d; expected: %d", response.Code, http.StatusOK)
	assert.Equal(t, http.StatusNotFound, response.Code, messageError)
	var messageResponse controllers.Message
	json.Unmarshal(response.Body.Bytes(), &messageResponse)
	assert.Equal(t, "Pet is not registered", messageResponse.Message)
}

func TestDeleteStudent(t *testing.T) {
	database.Connection()
	AddPetMock()
	r := SetupRouters()
	r.DELETE("/api/v1/pets/:id", controllers.Remove)
	path := "/api/v1/pets/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestUpdateOnePet(t *testing.T) {
	database.Connection()
	AddPetMock()
	defer RemovePetMock()
	r := SetupRouters()
	pet := models.Pet{Name: "Gamora", Observation: "não gosta de secador", Breed: "Husky Siberiano", Size: "MEDIUM"}
	petFormatted, _ := json.Marshal(pet)
	r.PATCH("/api/v1/pets/:id", controllers.Update)
	path := "/api/v1/pets/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(petFormatted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestNotUpdateOnePet(t *testing.T) {
	database.Connection()
	AddPetMock()
	defer RemovePetMock()
	r := SetupRouters()
	pet := models.Pet{Name: "Gamora", Observation: "não gosta de secador", Breed: "Husky Siberiano", Size: "MEDIUM"}
	petFormatted, _ := json.Marshal(pet)
	r.PATCH("/api/v1/pets/:id", controllers.Update)
	path := "/api/v1/pets/0"
	request, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(petFormatted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestAddPet(t *testing.T) {
	database.Connection()
	defer RemovePetMock()
	r := SetupRouters()
	r.POST("/api/v1/pets", controllers.Add)
	pet := models.Pet{Name: "Morcega", Observation: "", Breed: "Pinscher", Size: "SMALL"}
	petFormatted, _ := json.Marshal(pet)
	request, _ := http.NewRequest("POST", "/api/v1/pets", bytes.NewBuffer(petFormatted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	var petResponse models.Pet
	json.Unmarshal(response.Body.Bytes(), &petResponse)
	assert.Equal(t, "Morcega", petResponse.Name)
	ID = int(petResponse.ID)
}
