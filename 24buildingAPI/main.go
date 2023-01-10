package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Welcome to building API")
	//creating  new router
	r := mux.NewRouter()

	//Seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Web Developement", CoursePrice: 299, Author: &Author{Fullname: "Angela Yu", Website: "appbrewery.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Complete Javascript course", CoursePrice: 499, Author: &Author{Fullname: "Jonas Schmedtmann", Website: "codingheroes.io"}})

	//Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//Lisening
	log.Fatal(http.ListenAndServe(":4000", r))

}

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //author is a type of the custome made struct Author type and we are making it a pointer variable so that the reference of the variable can be sent in future
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Our Fake DataBase
var courses []Course //It is a slice of Course(custom made struct)

// Middleware, helper - file
// This method will be part of the struct so we are passing the reference of the struct
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// Controllers - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in golang </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab the courseId from the request
	params := mux.Vars(r)

	//loop through courses and find the matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}

	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return

}

// Creating a controller to add a new course to the database
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//While adding a course to the datbase we have to check the data for many possible mallicious data
	//What if: the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//What if: the data is like  this{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) //We need to destructure the data in the form of the struct we made so we arepassing the fererence of the course struct

	//As we decoded the data in the form of the struct we made so the data is stored in the course variable whose reference we passes in the Decode() method.
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in the json")
		return
	}

	//Generate unique CourseId(string)
	//Append the course into the courses slice (our fake Database)

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) //Converting the randonly generated number to string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// Update course controller function
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course!")
	w.Header().Set("Content-Type", "application/json")

	//grab the id from the request
	params := mux.Vars(r)

	// 1.Loop through the courses slice, 2.find the course with matching id, 3.Delete the found course, 4.add a new updated course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	//TODO: send a response when id is not found

}

// Controller function to delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("content-type", "application/json")

	//loop through the courses slice, find the course with the id, delete the found course
	for index, course := range courses {
		params := mux.Vars(r)
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
