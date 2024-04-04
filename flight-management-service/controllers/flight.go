package controllers

import (
	"net/http"

	"github.com/my-crazy-lab/airline-group-services/flight-management-service/forms"
	"github.com/my-crazy-lab/airline-group-services/flight-management-service/models"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
}

var planeModel = new(models.Plane)

func (ctrl FlightController) Insert(c *gin.Context) {
	var form forms.InsertPlaneForm

	id, err := planeModel.Insert(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Plane could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plane inserted", "id": id})
}
