package handlers

type OTPRequest struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Location string `json:"location,omitempty"`
}

