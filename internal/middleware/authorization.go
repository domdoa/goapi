package middleware

import (
	"errors"
	"net/http"

	"github.com/domdoa/goapi/api"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Unauthorized")

// Modify the middleware to accept the database as an argument
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract the username and token from the request
		// var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		if token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// // Use the database instance to get login details
		// loginDetails := db.GetLoginDetails(username)
		// if loginDetails == nil || token != loginDetails.Token {
		// 	log.Error(UnAuthorizedError)
		// 	api.RequestErrorHandler(w, UnAuthorizedError)
		// 	return
		// }

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}
