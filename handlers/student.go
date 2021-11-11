package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"student-api/data"
	"student-api/db"

	"github.com/gorilla/mux"
)

type Student struct {
	conn *db.DBConnection
}

func NewStudent(c *db.DBConnection) *Student {
	return &Student{c}
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
	var id int64
	var err error
	vars := mux.Vars(r)
	if id, err = strconv.ParseInt(vars["id"], 10, 64); err != nil {
		http.Error(rw, "Invalid parameter type", http.StatusBadRequest)
		return
	}
	log.Println("id:", id)
	sd, err := data.GetStudent(s.conn, id)
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
	var std data.Student
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
	if std.AdmNo == 0 || std.Name == "" {
		http.Error(rw, "Should provide Admission number and Name", http.StatusBadRequest)
		return
	}
	err = data.AddStudent(s.conn, std)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (s *Student) updateStudent(rw http.ResponseWriter, r *http.Request) {
	var std data.Student
	var err error
	var id int64
	vars := mux.Vars(r)
	if id, err = strconv.ParseInt(vars["id"], 10, 64); err != nil {
		http.Error(rw, "Invalid parameter type", http.StatusBadRequest)
		return
	}
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
	if std.AdmNo != 0 {
		http.Error(rw, "Admission number cannot be updated", http.StatusBadRequest)
		return
	}
	if std.Name != "" {
		http.Error(rw, "Name cannot be updated", http.StatusBadRequest)
		return
	}
	err = data.UpdateStudent(s.conn, id, std.Address, std.Class, std.Age)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (s *Student) deleteStudent(rw http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	vars := mux.Vars(r)
	if id, err = strconv.ParseInt(vars["id"], 10, 64); err != nil {
		http.Error(rw, "Invalid parameter type", http.StatusBadRequest)
		return
	}
	log.Println("id:", id)
	err = data.DeleteStudent(s.conn, id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
