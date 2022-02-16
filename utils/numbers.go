package utils

import (
	"fmt"
	"regexp"

	"github.com/ochom/jumia-interview-task/models"
	"github.com/ochom/jumia-interview-task/models/dto"
)

// FilterInCountry gets numbers that match country code only
func FilterInCountry(code string, customers []*models.Customer) (resp []*models.Customer) {
	for _, v := range customers {
		customer := *v
		countryCodeRegex := fmt.Sprintf(`^\(%s\)`, code)
		match, err := regexp.Match(countryCodeRegex, []byte(customer.Phone))
		if err == nil && match {
			resp = append(resp, &customer)
		}
	}
	return
}

// GetCountry finds a country that matches regexp
func GetCountry(phone string) *models.Country {
	for _, country := range models.AllCountries {

		countryCodeRegex := fmt.Sprintf(`^\(%s\)`, country.Code)

		match, err := regexp.Match(countryCodeRegex, []byte(phone))
		if err != nil {
			return nil
		}

		if match {
			return &country
		}

	}
	return nil
}

// FormatNumbers returns a list of numbers formated according to country code
func FormatNumbers(customers []*models.Customer) (resp []*dto.FormattedNumber) {
	for _, v := range customers {
		customer := *v

		country := GetCountry(customer.Phone)
		if country == nil {
			return nil
		}

		regex := regexp.MustCompile(`\S+$`)
		phone := regex.FindStringSubmatch(customer.Phone)[0]

		formatted := &dto.FormattedNumber{
			Country: country.Name,
			Code:    fmt.Sprintf("+%s", country.Code),
			Phone:   phone,
		}

		// validate the whole phone number string
		match, err := regexp.Match(country.Regex, []byte(customer.Phone))
		if err != nil || !match {
			formatted.State = "NOK"
		} else {
			formatted.State = "OK"
		}

		resp = append(resp, formatted)
	}

	return
}
