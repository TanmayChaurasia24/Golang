package main

import (
	"fmt"
	"net/http"
)

// model for course -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// middleware
func isEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("api in golang")
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home page api made in golang</h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("content-type", "application/json")
}
