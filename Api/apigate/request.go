package apigate

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		log.Println(err)
		requestorDetail := GetRequestorDetail(r)
		requestorDetail.Body = string(body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		LogRequest("", requestorDetail, "")
		next.ServeHTTP(w, r)

	})
}

//  In case You need To set Read Time out  use Below methods
// func RequestAndTimeoutMiddleware(next http.Handler, getReadTimeout func(*http.Request) time.Duration) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Read the entire request body
// 		body, err := ioutil.ReadAll(r.Body)
// 		log.Println(err)

// 		// Get request details and update the body
// 		requestorDetail := GetRequestorDetail(r)
// 		requestorDetail.Body = string(body)
// 		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

// 		// Log the request
// 		LogRequest("", requestorDetail)

// 		// Use the timeoutMiddleware for handling timeouts
// 		timeoutMiddleware(next, getReadTimeout).ServeHTTP(w, r)
// 	})
// }

// // Custom WriteTimeout function
// func CustomReadTimeout(r *http.Request) time.Duration {
// 	// Example: Set no timeout for a particular endpoint set -1
// 	if r.URL.Path == "/getCurVersion" {
// 		return 20 * time.Second
// 	}
// 	// Default write timeout
// 	return 15 * time.Second
// }
// func timeoutMiddleware(next http.Handler, getReadTimeout func(*http.Request) time.Duration) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Get the read timeout duration from the custom function
// 		readTimeout := getReadTimeout(r)

// 		// Check if a specific timeout is set for this endpoint
// 		if readTimeout <= 0 {
// 			// No timeout for this endpoint
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		// Create a new context with the specified read timeout
// 		ctx, cancel := context.WithTimeout(r.Context(), readTimeout)
// 		defer cancel()
// 		r = r.WithContext(ctx)

// 		// Create a channel to signal when the request handling is finished
// 		done := make(chan struct{})

// 		// Start a goroutine to execute the actual HTTP request handling
// 		go func() {
// 			// Based on Done Return Server will close
// 			defer close(done) // Signal that handling is finished
// 			next.ServeHTTP(w, r)
// 		}()

// 		// Use a select statement to wait for either the request handling to complete or the timeout to occur
// 		select {
// 		case <-done:
// 			// Request completed within the specified timeout
// 		case <-ctx.Done():
// 			// request timed out
// 			w.WriteHeader(http.StatusGatewayTimeout)
// 			// handle timeout, like logging or sending a specific response
// 		}
// 	})
// }
