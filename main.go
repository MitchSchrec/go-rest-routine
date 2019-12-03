// TODO - Fix error where request isn't sending
/* Notes:
*	Estimated Time Worked On: 4 hours
*	Knowledge of go before: none
* 	Knowledge of go after: little
*	Getting the server up and running was very simple using mux. Documentation for this was great
*	Working with goroutines and channels is where I struggled the most. With the little time I spent on it, I was unable
*	to grasp the concept.
* 	However, I know with more time and classes, I will begin to understand the Go Language and be able to write better code
* 	I am more used to working with data vs working with bytes
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// Used for server creation
	"github.com/gorilla/mux"
)

type Count struct {
	Count int
}

type HTTPError struct {
	Code    string
	Message string
}

// get the body request which looks like {"count": xx}
// Using postman to test
func parseCount(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var count Count
	err := decoder.Decode(&count)
	if err != nil {
		panic(err)
	}

	// Declare channel
	c := make(chan int)

	// Loop from 1 to the count passed in, wait a second, then add to the channel
	for i := 1; i <= count.Count; i++ {
		go func(i int) {
			time.Sleep(time.Duration(1)*time.Second)
			c <- + i
		}(i)
	}

	// print count
	for val := range c {
		fmt.Println(val)
	}

	// close channel
	close(c)

	// success message
	fmt.Fprint(rw, "Count is complete")
}

// return error code and message when route is not found
func Handle404() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		json.NewEncoder(rw).Encode(HTTPError{
			Code:    string(http.StatusNotFound),
			Message: http.StatusText(http.StatusNotFound),
		})
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Only allow post methods to the following path
	router.HandleFunc("/print", parseCount).Methods("POST")

	// Handle when route is not found
	router.NotFoundHandler = Handle404()

	port := "8080"
	fmt.Println("Running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
