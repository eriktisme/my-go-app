package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"net/http"
	u "passport/src/models"
)

func Login(c *gin.Context) {
	var validate *validator.Validate
	var payload u.Login

	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		return
	}

	validate = validator.New()

	if errors := validate.Struct(payload); errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": errors.Error(),
		})

		return
	}

	user, err := u.GetUserByEmail(payload.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "Invalid credentials",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Welcome, %s", user.Name),
	})
}
