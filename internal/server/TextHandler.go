package server

import (
	"Mangtas_/internal/helpers"
	"Mangtas_/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server)  TextHandler () gin.HandlerFunc{
	return func(context *gin.Context) {
		extractor := &models.Extractor{}

		// Binds the extractor struct
		err := context.BindJSON(extractor)
		if err != nil{
			context.JSON(http.StatusBadRequest,"unable to bind json")
			return
		}

		// Gets the result from the counter helper and returns the result
		result := helpers.Counter(extractor.Content)

		var resultExtracted = " "

		for _, v := range result{
			resultExtracted += fmt.Sprintf("%s -> : %d,  ",v.Key,v.Value )
		}

		// returns the result in JSON format
		context.JSON(http.StatusCreated,gin.H{
			"result" : resultExtracted,
		})
	}
}