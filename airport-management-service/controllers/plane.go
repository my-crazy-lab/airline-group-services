package controllers

import (
	"net/http"

	"github.com/my-crazy-lab/airline-group-services/airport-management-service/forms"
	"github.com/my-crazy-lab/airline-group-services/airport-management-service/models"

	"github.com/gin-gonic/gin"
)

type PlaneController struct {
}

var planeModel = new(models.Plane)

func (ctrl PlaneController) Insert(c *gin.Context) {
	var form forms.InsertPlaneForm

	id, err := planeModel.Insert(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Plane could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plane inserted", "id": id})
}
