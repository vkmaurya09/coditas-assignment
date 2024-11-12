package handlers

import (
	"fmt"
	"log"
	"net/http"
	"pan-service/internal/models"
	validation "pan-service/internal/validator"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("pan", validation.ValidatePAN)
}

func SubmitHandler(data *[]models.Payload, mu *sync.Mutex, errorLog *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			errorLog.Printf("Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}

		if err := validate.Struct(payload); err != nil {
			errs := err.(validator.ValidationErrors)
			errorMessages := make(map[string]string)
			for _, e := range errs {
				switch e.Tag() {
				case "required":
					errorMessages[e.Field()] = fmt.Sprintf("%s is required", e.Field())
				case "len":
					errorMessages[e.Field()] = fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
				case "numeric":
					errorMessages[e.Field()] = fmt.Sprintf("%s must be a numeric value", e.Field())
				case "email":
					errorMessages[e.Field()] = fmt.Sprintf("%s must be a valid email address", e.Field())
				case "pan":
					errorMessages[e.Field()] = fmt.Sprintf("%s must be a valid PAN number", e.Field())
				default:
					errorMessages[e.Field()] = fmt.Sprintf("%s is not valid", e.Field())
				}
			}
			errorLog.Printf("Validation errors: %v", errorMessages)
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "errors": errorMessages})
			return
		}

		mu.Lock()
		*data = append(*data, payload)
		mu.Unlock()

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Details received successfully"})
	}
}
