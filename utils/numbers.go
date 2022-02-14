package utils

import (
	"fmt"
	"regexp"

	"github.com/ochom/jumia-interview-task/models"
	"github.com/ochom/jumia-interview-task/models/dto"
)

// GetCountry finds a country that matches regexp
func GetCountry(phone string) *models.Country {
	for _, country := range models.AllCountries {

		match, err := regexp.Match(country.Regex, []byte(phone))
		if err != nil {
			return nil
		}

		if match {
			return &country
		}

	}
	return nil
}

// FormatNumber formats a number according to country regex
func FormatNumber(customer models.Customer) *dto.FormattedNumber {
	country := GetCountry(customer.Phone)
	if country == nil {
		return nil
	}

	regex := regexp.MustCompile(`\S+$`)
	parts := regex.FindStringSubmatch(customer.Phone)

	formatted := &dto.FormattedNumber{
		Country: country.Name,
		Code:    fmt.Sprintf("+%s", country.Code),
		State:   "OK",
		Phone:   parts[0],
	}
	return formatted
}
