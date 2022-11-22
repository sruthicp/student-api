package repositories

import (
	"log"
	"student-api/model"

	"github.com/go-pg/pg"
)

const (
	AdmissionNoQry = "admission_no = ?"
	AddressQry     = "std_address = ?"
	ClassQry       = "std_class = ?"
	AgeQry         = "std_age = ?"
)

type StudentRepo struct {
	con *pg.DB
}

func NewStudentRepo(connection *pg.DB) *StudentRepo {
	return &StudentRepo{
		con: connection,
	}
}

// GetStudent returns student details
func (sr *StudentRepo) GetStudent(id string) (std *model.Student, err error) {
	std = &model.Student{}
	err = sr.con.Model(std).Where(AdmissionNoQry, id).Select()
	if err != nil {
		return
	}

	return
}

// AddStudent save student details
func (sr *StudentRepo) AddStudent(std *model.Student) (err error) {
	if err = sr.con.Insert(std); err != nil {
		return
	}

	return
}

// UpdateStudent update student details for an Admission number
func (sr *StudentRepo) UpdateStudent(id string, std *model.Student) (err error) {
	log.Println("updating :", std)

	qry := sr.con.Model(std)

	if _, err = qry.Where(AdmissionNoQry, id).UpdateNotNull(); err != nil {
		return
	}

	return
}

// DeleteStudent delete student entry for an admission number
func (sr *StudentRepo) DeleteStudent(id string) (err error) {
	std := &model.Student{}

	_, err = sr.con.Model(std).Where(AdmissionNoQry, id).Delete()
	if err != nil {
		return
	}

	return
}
