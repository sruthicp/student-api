package controller

import (
	"net/http"
	"student-api/model"
	protos "student-api/proto/student"
	"student-api/service"
)

type StudentController struct {
	service *service.StudentService
}

func NewStudentController(svc *service.StudentService) *StudentController {
	return &StudentController{
		service: svc,
	}
}

func (sc *StudentController) CreateStudent(req *protos.CreateStudentRequest) (res *protos.CreateStudentResponse, err error) {
	res = &protos.CreateStudentResponse{
		Message:    "Successfully added student details",
		Statuscode: http.StatusOK,
	}

	reqModel := &model.Student{
		Name:    req.Name,
		Address: req.Address,
		Class:   req.Class,
		Age:     req.Age,
	}

	id, err := sc.service.AddStudent(reqModel)
	if err != nil {
		res.Message = "Failed to add student details"
		res.Statuscode = http.StatusInternalServerError

		return res, err
	}

	res.AdmNo = id

	return res, nil
}

func (sc *StudentController) GetStudent(req *protos.BaseStudentRequest) (res *protos.GetStudentResponse, err error) {
	res = &protos.GetStudentResponse{
		Message:    "Successfully fetched student details",
		Statuscode: http.StatusOK,
	}
	var details *model.Student

	details, err = sc.service.GetStudent(req.AdmNo)
	if err != nil {
		res.Message = "Failed to get student details"
		res.Statuscode = http.StatusInternalServerError

		return res, err
	}

	res.Details = modelToProto(details)

	return res, nil
}

func (sc *StudentController) UpdateStudent(req *protos.UpdateStudentRequest) (res *protos.BaseStudentResponse, err error) {
	res = &protos.BaseStudentResponse{
		Message:    "Successfully updated student details",
		Statuscode: http.StatusOK,
	}

	if err = sc.service.UpdateStudent(req.AdmNo, req.Address, req.Class, req.Age); err != nil {
		res.Message = "Failed to update student details"
		res.Statuscode = http.StatusInternalServerError

		return res, err
	}

	return res, nil
}

func (sc *StudentController) DeleteStudent(req *protos.BaseStudentRequest) (res *protos.BaseStudentResponse, err error) {
	res = &protos.BaseStudentResponse{
		Message:    "Successfully deleting student details",
		Statuscode: http.StatusOK,
	}

	if err = sc.service.DeleteStudent(req.AdmNo); err != nil {
		res.Message = "Failed to delete student details"
		res.Statuscode = http.StatusInternalServerError

		return res, err
	}

	return res, nil
}

func modelToProto(in *model.Student) *protos.StudentDetails {
	return &protos.StudentDetails{
		AdmNo:   in.AdmNo,
		Name:    in.Name,
		Class:   in.Class,
		Address: in.Address,
		Age:     in.Age,
	}
}
