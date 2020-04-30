package constant

import "net/http"

//RegistrationUnsuccessfullMsg show on error time
var RegistrationUnsuccessfullMsg = "Registration Unsuccessful"

//RegistrationSuccessfullMsg show on successful registration
var RegistrationSuccessfullMsg = "Registration Successful"

//GetuserDetailSuccessfullMsg show on successful fetch user details
var GetuserDetailSuccessfullMsg = "Get User details Successful"

//SetCommonHeader for all api call function
func SetCommonHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}
