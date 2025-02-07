package model

type APIResponse struct {
	Result   int        `json:"result"`
	Status   int        `json:"status"`
	Message  string     `json:"message"`
	MockData []MockData `json:"mock_data"`
}

type MockData struct {
	Id        int16  `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Gender    string `json:"gender,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
}
