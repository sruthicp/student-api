package repositories

import (
	"log"
	"student-api/common"
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
func (sr *StudentRepo) GetStudent(id string) (std *model.Student, err *common.DBError) {
	std = &model.Student{}
	dberr := sr.con.Model(std).Where(AdmissionNoQry, id).Select()
	if dberr != nil {
		return std, common.NewDBError("failed to get student entry", dberr)
	}

	return
}

// AddStudent save student details
func (sr *StudentRepo) AddStudent(std *model.Student) (err *common.DBError) {
	if dberr := sr.con.Insert(std); dberr != nil {
		return common.NewDBError("failed to add student entry", dberr)
	}

	return
}

// UpdateStudent update student details for an Admission number
func (sr *StudentRepo) UpdateStudent(id string, std *model.Student) *common.DBError {
	log.Println("updating :", std)

	qry := sr.con.Model(std)

	res, err := qry.Where(AdmissionNoQry, id).UpdateNotNull()
	if err != nil {
		return common.NewDBError("failed to update student entry", err)
	}

	if res == nil {
		return common.NewDBError("failed to update student entry", pg.ErrNoRows)
	}

	return nil
}

// DeleteStudent delete student entry for an admission number
func (sr *StudentRepo) DeleteStudent(id string) (err *common.DBError) {
	std := &model.Student{}

	_, dberr := sr.con.Model(std).Where(AdmissionNoQry, id).Delete()
	if dberr != nil {
		return common.NewDBError("failed to delete student entry", dberr)
	}

	return
}
