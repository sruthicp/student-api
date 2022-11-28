package controller

import (
	"context"
	"errors"
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

func (sc *StudentController) CreateStudent(ctx context.Context, req *protos.CreateStudentRequest) (*protos.CreateStudentResponse, error) {
	res := &protos.CreateStudentResponse{
		Message:    "Successfully added student details",
		StatusCode: http.StatusOK,
	}

	reqModel := &model.Student{
		Name:    req.Name,
		Address: req.Address,
		Class:   req.Class,
		Age:     req.Age,
	}

	id, err := sc.service.AddStudent(reqModel)
	if err != nil {
		res.Message = err.Msg
		res.StatusCode = int32(err.Status)

		return res, errors.New(err.Msg)
	}

	res.AdmNo = id

	return res, nil
}

func (sc *StudentController) GetStudent(ctx context.Context, req *protos.BaseStudentRequest) (*protos.GetStudentResponse, error) {
	res := &protos.GetStudentResponse{
		Message:    "Successfully fetched student details",
		StatusCode: http.StatusOK,
	}
	var details *model.Student

	details, err := sc.service.GetStudent(req.AdmNo)
	if err != nil {
		res.Message = err.Msg
		res.StatusCode = int32(err.Status)

		return res, errors.New(err.Msg)
	}

	res.Details = modelToProto(details)

	return res, nil
}

func (sc *StudentController) UpdateStudent(ctx context.Context, req *protos.UpdateStudentRequest) (*protos.BaseStudentResponse, error) {
	res := &protos.BaseStudentResponse{
		Message:    "Successfully updated student details",
		StatusCode: http.StatusOK,
	}
	std := &model.Student{
		Name:    req.Name,
		Address: req.Address,
		Class:   req.Class,
		Age:     req.Age,
	}

	if err := sc.service.UpdateStudent(req.AdmNo, std); err != nil {
		res.Message = err.Msg
		res.StatusCode = int32(err.Status)

		return res, errors.New(err.Msg)
	}

	return res, nil
}

func (sc *StudentController) DeleteStudent(ctx context.Context, req *protos.BaseStudentRequest) (*protos.BaseStudentResponse, error) {
	res := &protos.BaseStudentResponse{
		Message:    "Successfully deleting student details",
		StatusCode: http.StatusOK,
	}

	if err := sc.service.DeleteStudent(req.AdmNo); err != nil {
		res.Message = err.Msg
		res.StatusCode = int32(err.Status)
		return res, errors.New(err.Msg)
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
