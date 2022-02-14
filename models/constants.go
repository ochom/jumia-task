package models

// Country ...
type Country struct {
	Name  string
	Code  string
	Regex string
}

// AllCountries returns a slice of all countries
var AllCountries = []Country{
	{Name: "Cameron", Code: "237", Regex: `\(237\)\ ?[2368]\d{7,8}$`},
	{Name: "Ethiopia", Code: "251", Regex: `\(251\)\ ?[1-59]\d{8}$`},
	{Name: "Moroco", Code: "212", Regex: `\(212\)\ ?[5-9]\d{8}$`},
	{Name: "Mozambique", Code: "258", Regex: `\(258\)\ ?[28]\d{7,8}$`},
	{Name: "Uganda", Code: "256", Regex: `\(256\)\ ?\d{9}$`},
}
