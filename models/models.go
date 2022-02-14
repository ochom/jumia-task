package models

// Customer ...
type Customer struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}
