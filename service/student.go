package service

import (
	"student-api/model"
	"student-api/repositories"

	"log"

	"github.com/google/uuid"
)

type StudentService struct {
	repo *repositories.StudentRepo
}

func NewStudentService(studrepo *repositories.StudentRepo) *StudentService {
	return &StudentService{
		repo: studrepo,
	}
}

func (ss *StudentService) AddStudent(req *model.Student) (admnno string, err error) {
	req.AdmNo = uuid.New().String()

	if err = ss.repo.AddStudent(req); err != nil {
		log.Println("error creating student: ", err)
		return "", err
	}

	return req.AdmNo, err
}

func (ss *StudentService) GetStudent(id string) (res *model.Student, err error) {
	res, err = ss.repo.GetStudent(id)
	if err != nil {
		log.Println("error getting student: ", err)
		return &model.Student{}, err
	}

	return res, err
}

func (ss *StudentService) UpdateStudent(id, address, class string, age int32) (err error) {
	err = ss.repo.UpdateStudent(id, address, class, age)
	if err != nil {
		log.Println("error updating student: ", err)
		return err
	}

	return err
}

func (ss *StudentService) DeleteStudent(id string) (err error) {
	err = ss.repo.DeleteStudent(id)
	if err != nil {
		log.Println("error deleting student: ", err)
		return err
	}

	return err
}
