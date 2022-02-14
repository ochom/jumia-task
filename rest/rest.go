package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/usecases"
)

// Handler abstracts methods implemented for rest requests
type Handler interface {
	// GetPhonenumbers gets all phone numbers
	GetPhonenumbers() gin.HandlerFunc

	// GetCountryPhoneNumbers gets all valid phone numbers in a country
	GetCountryPhoneNumbers() gin.HandlerFunc
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

func (h *impl) GetPhonenumbers() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := h.uc.GetAll(c)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (h *impl) GetCountryPhoneNumbers() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		if code == "" {
			c.String(http.StatusBadRequest, "country `code` is required")
			return
		}

		res, err := h.uc.GetByCountry(c, code)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, res)

	}
}
