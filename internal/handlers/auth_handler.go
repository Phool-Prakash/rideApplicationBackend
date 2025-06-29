package handlers

import (
	"encoding/json"
	"net/http"
	"rideApplicationBackend/internal/services"
	"rideApplicationBackend/internal/utils"
)

type OTPRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Location string `json:"location,omitempty"`
}

func sendOTP(w http.ResponseWriter, r *http.Request) {
	var req OTPRequest
	json.NewDecoder(r.Body).Decode(&req)

	if len(req.Phone) < 10{
		utils.ResponseJSON(w, http.StatusBadRequest, "Phone number must be at least 10 digits",nil)
		return
	}

	otp := services.GenerateOTP(req.Phone)
	utils.ResponseJSON(w,http.StatusOK,"OTP sent", map[string]string{"OTP":otp})
}