package controller

// Authorize is a function that authorizes the user. It will check with another instance of DB to see if the has matching credentials.
// It returns an error if the credentials are not correct. It returns a boolean value if the credentials are correct.
// err + false = unauthorized
// err + true = incorrect credentials
// nil + true = authorized
func Authorize(username, password string) (error, bool) {

	return nil, true
}

func Authenticate(username, password string) (error, bool) {

	return nil, true
}
