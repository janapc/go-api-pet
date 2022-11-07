package controllers

import (
	"go-api-pet/database"
	"go-api-pet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

// HealthCheck godoc
// @Summary Check the API
// @Description Check the API is on
// @Tags pets
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/pets/health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

// ListAll godoc
// @Summary List all pets
// @Description List all pets
// @Tags pets
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Pet
// @Router /api/v1/pets [get]
func ListAll(c *gin.Context) {
	var pets []models.Pet
	database.DB.Find(&pets)
	c.JSON(http.StatusOK, pets)
}

// FindOneById godoc
// @Summary List a pet
// @Description List a pet by Id
// @Tags pets
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "id of pet"
// @Success 200 {object} models.Pet
// @Failure 404 {object} Message
// @Router /api/v1/pets/{id} [get]
func FindOneById(c *gin.Context) {
	var pet models.Pet
	id := c.Params.ByName("id")
	database.DB.First(&pet, id)
	if pet.ID == 0 {
		messageError := Message{"Pet is not registered"}
		c.JSON(http.StatusNotFound, messageError)
		return
	}
	c.JSON(http.StatusOK, pet)
}

// Add godoc
// @Summary Add a new pet
// @Description Add a new pet
// @Tags pets
// @Accept  json
// @Produce  json
// @Param   pet     body    models.Pet     true        "Model of pet"
// @Success 201 {object} models.Pet
// @Failure 400 {object} Message
// @Router /api/v1/pets [post]
func Add(c *gin.Context) {
	var pet models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		messageError := Message{err.Error()}
		c.JSON(http.StatusBadRequest, messageError)
		return
	}
	if err := models.ValidatorPet(&pet); err != nil {
		messageError := Message{err.Error()}
		c.JSON(http.StatusBadRequest, messageError)
		return
	}
	database.DB.Create(&pet)
	c.JSON(http.StatusCreated, pet)
}

// Update godoc
// @Summary Update a new pet
// @Description Update a new pet by Id
// @Tags pets
// @Accept  json
// @Produce  json
// @Param   pet     body    models.Pet     true        "Model of pet"
// @Param   id    path    int     true        "Id of pet"
// @Success 200 {object} models.Pet
// @Failure 400 {object} message
// @Failure 404 {object} message
// @Router /api/v1/pets/{id} [patch]
func Update(c *gin.Context) {
	var pet models.Pet
	id := c.Params.ByName("id")
	database.DB.First(&pet, id)
	if pet.ID == 0 {
		messageError := Message{"Pet is not registered"}
		c.JSON(http.StatusNotFound, messageError)
		return
	}
	if err := c.ShouldBindJSON(&pet); err != nil {
		messageError := Message{err.Error()}
		c.JSON(http.StatusBadRequest, messageError)
		return
	}
	if err := models.ValidatorPet(&pet); err != nil {
		messageError := Message{err.Error()}
		c.JSON(http.StatusBadRequest, messageError)
		return
	}
	database.DB.Model(&pet).UpdateColumns(pet)
	c.JSON(http.StatusOK, pet)
}

// Remove godoc
// @Summary Remove the registration of a pet
// @Description Remove the registration of a pet by Id
// @Tags pets
// @Accept  json
// @Produce  json
// @Param   id    path    int     true        "Id of pet"
// @Success 204
// @Router /api/v1/pets/{id} [delete]
func Remove(c *gin.Context) {
	var pet models.Pet
	id := c.Params.ByName("id")
	database.DB.Delete(&pet, id)
	c.JSON(http.StatusNoContent, nil)
}
