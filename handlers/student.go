package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"student-api/config"
	"student-api/db"
	"student-api/model"
	"student-api/repositories"

	"github.com/gorilla/mux"
)

//to be deprecated...

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"student-api/db"
// 	"student-api/model"
// 	"student-api/repositories"

// 	"github.com/gorilla/mux"
// )

type Student struct {
	repo *repositories.StudentRepo
}

func NewStudent() *Student {
	config.NewServiceConfig()
	connection, _ := db.NewDBConnection(config.SvcConf)

	return &Student{
		repo: repositories.NewStudentRepo(connection),
	}
}

func (s *Student) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Service started..")
	switch r.Method {
	case http.MethodGet:
		s.getStudent(rw, r)
	case http.MethodPost:
		s.addStudent(rw, r)
	case http.MethodPut:
		s.updateStudent(rw, r)
	case http.MethodDelete:
		s.deleteStudent(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (s *Student) getStudent(rw http.ResponseWriter, r *http.Request) {
	var id string
	var err error
	vars := mux.Vars(r)
	id = vars["id"]
	log.Println("id:", id)
	sd, err := s.repo.GetStudent(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(rw, "Admission number not found", http.StatusNotFound)
			return
		}
		http.Error(rw, "Unable to get student details", http.StatusInternalServerError)
		return
	}
	// serialize the list to JSON
	err = sd.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (s *Student) addStudent(rw http.ResponseWriter, r *http.Request) {
	std := &model.Student{}
	var err error

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if string(b) == "{}" {
		http.Error(rw, "Empty request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(b, &std); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if std.AdmNo == "" || std.Name == "" {
		http.Error(rw, "Should provide Admission number and Name", http.StatusBadRequest)
		return
	}
	err = s.repo.AddStudent(std)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (s *Student) updateStudent(rw http.ResponseWriter, r *http.Request) {
	var std *model.Student
	var err error
	var id string
	vars := mux.Vars(r)

	id = vars["id"]
	log.Println("id:", id)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if string(b) == "{}" {
		http.Error(rw, "Empty request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(b, &std); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if std.AdmNo != "" {
		http.Error(rw, "Admission number cannot be updated", http.StatusBadRequest)
		return
	}
	if std.Name != "" {
		http.Error(rw, "Name cannot be updated", http.StatusBadRequest)
		return
	}
	err = s.repo.UpdateStudent(id, std)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (s *Student) deleteStudent(rw http.ResponseWriter, r *http.Request) {
	var id string
	var err error
	vars := mux.Vars(r)
	id = vars["id"]

	log.Println("id:", id)
	err = s.repo.DeleteStudent(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
