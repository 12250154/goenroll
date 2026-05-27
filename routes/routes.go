// package routes

// import (
// 	"fmt"
// 	"log"
// 	"myapp/controller"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// // func InitializeRoutes() {

// // 	//creating a new router
// // 	router := mux.NewRouter()

// // 	//register handler function with the mux router
// // 	//student routes
// // 	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
// // 	router.HandleFunc("/student/all", controller.GetAllStudents)
// // 	router.HandleFunc("/student/{sid}", controller.GetStudent).Methods("GET") //by default the method is GET
// // 	router.HandleFunc("/student/{sid}", controller.UpdateStudent).Methods("PUT")
// // 	router.HandleFunc("/student/{sid}", controller.DeleteStudent).Methods("DELETE")
// // 	//router.HandleFunc("/students", controller.GetAllStudents)

// // 	// signup and login
// // 	router.HandleFunc("/signup", controller.Signup).Methods("POST")
// // 	router.HandleFunc("/login", controller.Login).Methods("POST")
// // 	router.HandleFunc("/logout", controller.Logout)

// // 	//enroll APIs
// // 	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")

// // 	//load static files
// // 	fhandler := http.FileServer(http.Dir("./view"))
// // 	//serve static files as a route by registering all static files on the mux router
// // 	router.PathPrefix("/").Handler(fhandler)
// // 	//log.Println("Application running on port", port)

// // 	//course routes
// // 	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
// // 	router.HandleFunc("/course/all", controller.GetAllStudents).Methods("GET")
// // 	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET") //by default the method is GET
// // 	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
// // 	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")
// // 	//router.HandleFunc("/courses", controller.GetAllCourses)

// // 	fmt.Println("Server started successfully...")

// // 	//start the http server
// // 	log.Fatal(http.ListenAndServe(":8080", router))
// // }

// func InitializeRoutes() {

// 	router := mux.NewRouter()

// 	// ================= STUDENT ROUTES =================
// 	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
// 	router.HandleFunc("/student/all", controller.GetAllStudents).Methods("GET")
// 	router.HandleFunc("/student/{sid}", controller.GetStudent).Methods("GET")
// 	router.HandleFunc("/student/{sid}", controller.UpdateStudent).Methods("PUT")
// 	router.HandleFunc("/student/{sid}", controller.DeleteStudent).Methods("DELETE")

// 	// ================= AUTH ROUTES =================
// 	router.HandleFunc("/signup", controller.Signup).Methods("POST")
// 	router.HandleFunc("/login", controller.Login).Methods("POST")
// 	router.HandleFunc("/logout", controller.Logout)

// 	// ================= ENROLL =================
// 	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")
// 	router.HandleFunc("/enroll/all", controller.GetAllEnrollments).Methods("GET")

// 	// ================= COURSE ROUTES =================
// 	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
// 	router.HandleFunc("/course/all", controller.GetAllCourses).Methods("GET")
// 	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
// 	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
// 	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")

// 	// ================= STATIC FILES (ALWAYS LAST) =================
// 	fhandler := http.FileServer(http.Dir("./view"))
// 	router.PathPrefix("/").Handler(fhandler)

// 	fmt.Println("Server started successfully...")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// 	//http.ListenAndServe("0.0.0.0:8080", router)
// }

package routes

import (
	"fmt"
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes() {

	router := mux.NewRouter()

	// ================= STUDENT ROUTES =================
	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/all", controller.GetAllStudents).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.GetStudent).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStudent).Methods("DELETE")

	// ================= AUTH ROUTES =================
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout)

	// ================= ENROLL =================
	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")
	router.HandleFunc("/enroll/all", controller.GetAllEnrollments).Methods("GET")

	// ================= COURSE ROUTES =================
	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/all", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")

	// ================= STATIC FILES =================
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	// ================= CORS =================
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"https://goenroll-4-2ulc.onrender.com",
		}),
		handlers.AllowedMethods([]string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		}),
		handlers.AllowedHeaders([]string{
			"Content-Type", "Authorization",
		}),
	)(router)

	fmt.Println("Server started successfully...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
