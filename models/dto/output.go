package dto

// FormattedNumber are phone numbers formatted into respective countries
type FormattedNumber struct {
	Country string
	State   string
	Code    string
	Phone   string
}
