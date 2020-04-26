package models

//User is a repreasention of user data
type User struct {
	UserID     int    `json:"userid"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	EmailID    string `json:"emailid"`
	Password   string `json:"password"`
	IsRemember bool   `json:"isremember"`
}

//ResponseResult represention of error and result type
type ResponseResult struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Result  []User `json:"result"`
}
