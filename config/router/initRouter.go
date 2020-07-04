package router

import (
	"database/sql"
	"fmt"
	"golang_clean_architecture_gerin/domains/students"
	"golang_clean_architecture_gerin/domains/subjects"
	"golang_clean_architecture_gerin/domains/teachers"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	STUDENTS_MAIN_ROUTE = "/students"
	TEACHERS_MAIN_ROUTE = "/teachers"
	SUBJECTS_MAIN_ROUTE = "/subjects"
)

type ConfigRouter struct {
	DB     *sql.DB
	Router *mux.Router
}

func (ar *ConfigRouter) InitRouter() {
	students.InitStudentRoute(STUDENTS_MAIN_ROUTE, ar.DB, ar.Router)
	teachers.InitTeacherRoute(TEACHERS_MAIN_ROUTE, ar.DB, ar.Router)
	subjects.InitSubjectRoute(SUBJECTS_MAIN_ROUTE, ar.DB, ar.Router)
	ar.Router.NotFoundHandler = http.HandlerFunc(notFound)
	ar.Router.HandleFunc("/", homePage)
}

// NewAppRouter for creating new Route
func NewAppRouter(db *sql.DB, r *mux.Router) *ConfigRouter {
	return &ConfigRouter{
		DB:     db,
		Router: r,
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,
		`<!DOCTYPE html>
		<html lang="en">
		  <head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>Enigma | School</title>
		  </head>
		  <body>
			<h1>Selamat Datang Di Enigma School</h1>
			<ul>
			  <li>
				<a href="http://localhost:8080/students"><h4>Data Students</h4></a>
			  </li>
			  <li>
				<a href="http://localhost:8080/teachers"><h4>Data Seachers</h4></a>
			  </li>
			  <li>
				<a href="http://localhost:8080/subjects"><h4>Data Subjects</h4></a>
			  </li>
			</ul>
		  </body>
		</html>`)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>404 Status Not Found</h1>`)
}
