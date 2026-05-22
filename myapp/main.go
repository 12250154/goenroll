package main

import "myapp/routes"

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	//extract URL parameter from request
// 	myMap := mux.Vars(r)
// 	value := myMap["course"]
// 	//write the response to client
// 	_, err := w.Write([]byte("Hello world \nThe value send in url request is " + value))
// 	if err != nil {
// 		fmt.Println("error: ", err.Error())
// 	}
// }

func main() {
	// //creating a new router
	// router := mux.NewRouter()

	// //register handler function with the mux router
	// router.HandleFunc("/home/{course}", homeHandler)
	// fmt.Println("Server started successfully...")

	// //start the http server
	// log.Fatal(http.ListenAndServe(":8080", router))
	routes.InitializeRoutes()

}

// Define handler function to handle incoming request
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	//write the response to client

// 	_, err := w.Write([]byte("hello world"))
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 	}
// }
