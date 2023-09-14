package handlers

import (
	"errors"
	"github/meshachdamilare/hng_stage_two/database"
	"github/meshachdamilare/hng_stage_two/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePerson(c *gin.Context) {
	var person models.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := database.CreatePerson(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Person created",
		"id":      person.ID,
		"name":    person.Name,
		"track":   person.Track,
	})
}

func GetPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := database.GetPerson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func UpdatePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingPerson, err := database.GetPerson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedFields models.Person
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only the provided fields
	if updatedFields.Name != "" {
		existingPerson.Name = updatedFields.Name
	}
	if updatedFields.Track != "" {
		existingPerson.Track = updatedFields.Track
	}

	// Ensure that the ID remains the same as the existing person
	existingPerson.ID = uint(id)

	updatedPerson := database.UpdatePerson(*existingPerson)
	c.JSON(http.StatusOK, updatedPerson)
}

func DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DeletePerson(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}
