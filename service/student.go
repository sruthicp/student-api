package service

import (
	"net/http"
	"student-api/common"
	"student-api/model"
	"student-api/repositories"

	"log"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

const (
	addErr    = "Failed to add student details"
	getErr    = "Failed to get student details"
	updateErr = "Failed to update student details"
	deleteErr = "Failed to delete student details"
)

type StudentService struct {
	repo *repositories.StudentRepo
}

func NewStudentService(studrepo *repositories.StudentRepo) *StudentService {
	return &StudentService{
		repo: studrepo,
	}
}

func (ss *StudentService) AddStudent(req *model.Student) (string, *common.ServerError) {
	req.AdmNo = uuid.New().String()

	if err := ss.repo.AddStudent(req); err != nil {
		log.Println("error creating student: ", err)
		return "", common.NewServerError(addErr, http.StatusInternalServerError)
	}

	return req.AdmNo, nil
}

func (ss *StudentService) GetStudent(id string) (*model.Student, *common.ServerError) {
	res, err := ss.repo.GetStudent(id)
	if err != nil && err.Err == pg.ErrNoRows {
		log.Println("error is pg.ErrNoRows")
		return &model.Student{}, common.NewServerError(getErr, http.StatusNotFound)
	} else if err != nil {
		log.Println("error getting student: ", err)
		return &model.Student{}, common.NewServerError(getErr, http.StatusInternalServerError)
	}

	return res, nil
}

func (ss *StudentService) UpdateStudent(id string, std *model.Student) *common.ServerError {
	err := ss.repo.UpdateStudent(id, std)
	if err != nil && err.Err == pg.ErrNoRows {
		return common.NewServerError(updateErr, http.StatusNotFound)
	} else if err != nil {
		log.Println("error updating student: ", err)
		return common.NewServerError(updateErr, http.StatusInternalServerError)
	}

	return nil
}

func (ss *StudentService) DeleteStudent(id string) *common.ServerError {
	err := ss.repo.DeleteStudent(id)
	if err != nil {
		log.Println("error deleting student: ", err)
		return common.NewServerError(deleteErr, http.StatusInternalServerError)
	}

	return nil
}
