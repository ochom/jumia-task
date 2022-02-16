package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/models"
	"github.com/ochom/jumia-interview-task/usecases"
)

// Handler abstracts methods implemented for rest requests
type Handler interface {
	// GetCountries get a list of all countries
	GetCountries() gin.HandlerFunc

	// GetPhonenumbers gets all phone numbers
	GetPhonenumbers() gin.HandlerFunc
}

type impl struct {
	uc usecases.CustomersUsecase
}

// New creates a new instance of the Handler interface
func New(repo database.Repository) Handler {
	uc := usecases.New(repo)
	return &impl{
		uc: uc,
	}
}

func (h *impl) GetCountries() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.AllCountries)
	}
}

func (h *impl) GetPhonenumbers() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		state := c.Param("state")

		res, err := h.uc.GetNumbers(c, code, state)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
