package main

import (
	"encoding/json" // Package to work with JSON data
	"fmt"           // Package to print output to console
	"math/rand"     // Package to generate random numbers
	"net/http"      // Package to create HTTP servers and handle requests
	"strconv"       // Package to convert data types

	"github.com/gorilla/mux"
)

// Course struct represents the structure of a course object with details.
type Course struct {
	CourseId    string  `json:"courseid"`   // Unique identifier for the course
	CourseName  string  `json:"coursename"` // Name of the course
	CoursePrice int     `json:"price"`      // Price of the course
	Author      *Author `json:"author"`     // Author of the course
}

// Author struct represents the structure of an author object.
type Author struct {
	Fullname string `json:"fullname"` // Author's full name
	Website  string `json:"website"`  // Author's website
}

// Slice to store list of courses
var courses []Course

// Middleware function to check if a course is empty
func isEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

// main function is the entry point of the application.
func main() {
	fmt.Println("API in Golang")

	// Preload some course data for testing
	courses = append(courses, Course{CourseId: "1", CourseName: "Go Programming", CoursePrice: 299, Author: &Author{Fullname: "Tanmay Kumar", Website: "https://tanmaykumar.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python Programming", CoursePrice: 399, Author: &Author{Fullname: "John Doe", Website: "https://johndoe.com"}})

	// Set up routes for the HTTP server
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/courses", getAllCourses)
	http.HandleFunc("/course", createCourse)
	http.HandleFunc("/course/", getCourseById)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}

// serveHome handles requests to the home page.
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Home page - API made in Golang</h1>"))
}

// getAllCourses handles GET requests to fetch all available courses.
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) // Encode the slice of courses as JSON and send it as a response
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching course by ID")

	params := mux.Vars(r)
	courseId, exists := params["id"]

	if !exists {
		http.Error(w, "Course ID not provided", http.StatusBadRequest)
		return
	}

	// Search for the course with the matching ID
	for _, course := range courses {
		if course.CourseId == courseId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(course) // Send the matching course as a JSON response
			return
		}
	}

	// If not found, respond with 404 error
	http.Error(w, "Course not found", http.StatusNotFound)
}

// createCourse handles POST requests to add a new course.
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding a new course")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body into a Course struct
	var newCourse Course
	json.NewDecoder(r.Body).Decode(&newCourse)

	// Check if the course details are valid
	if isEmpty(&newCourse) {
		http.Error(w, "Invalid course details", http.StatusBadRequest)
		return
	}

	// Generate a random course ID for the new course
	newCourse.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, newCourse) // Add the new course to the list

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCourse) // Send the new course as a JSON response
}
