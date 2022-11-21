package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"student-api/db"
	"student-api/model"
)

const table = "school.student"

type StudentRepo struct {
	con *db.DBConnection
}

func NewStudentRepo() *StudentRepo {
	connection, _ := db.NewDBConnection()

	return &StudentRepo{
		con: connection,
	}
}

// GetStudent returns student details
func (sr *StudentRepo) GetStudent(id string) (std *model.Student, err error) {
	var sqlRows *sql.Rows

	sqlStr := fmt.Sprintf("select admission_no, std_name, std_address, std_class, std_age from school.student where admission_no = '%s'", id)
	log.Printf("query : %s", sqlStr)

	if sqlRows, err = sr.con.DB.Query(sqlStr); err != nil {
		return
	}

	defer sqlRows.Close()
	for sqlRows.Next() {
		if err = sqlRows.Scan(&std.AdmNo, &std.Name, &std.Address, &std.Class, &std.Age); err != nil {
			return
		}
	}
	if std == (&model.Student{}) {
		return std, sql.ErrNoRows
	}
	return
}

// AddStudent save student details
func (sr *StudentRepo) AddStudent(std *model.Student) (err error) {
	var sqlRows *sql.Rows

	ib := sr.con.Builder.NewInsertBuilder()
	ib.InsertInto(table).Cols("admission_no", "std_name", "std_address", "std_class", "std_age").Values(std.AdmNo, std.Name, std.Address, std.Class, std.Age)

	sqlStr, args := ib.Build()
	log.Println(sqlStr, args)
	if sqlRows, err = sr.con.DB.Query(sqlStr, args...); err != nil {
		return
	}

	defer sqlRows.Close()

	return
}

// UpdateStudent update student details for an Admission number
func (sr *StudentRepo) UpdateStudent(id, address, class string, age int32) (err error) {
	var sqlRows *sql.Rows
	var setAssigns []string

	ub := sr.con.Builder.NewUpdateBuilder()
	ub.Update(table)
	if address != "" {
		setAssigns = append(setAssigns, ub.Assign("std_address", address))
	}
	if class != "" {
		setAssigns = append(setAssigns, ub.Assign("std_class", class))
	}
	if age != 0 {
		setAssigns = append(setAssigns, ub.Assign("std_age", age))
	}

	ub.Set(setAssigns...).Where(ub.Equal("admission_no", id))

	sqlStr, args := ub.Build()
	log.Println(sqlStr, args)
	if sqlRows, err = sr.con.DB.Query(sqlStr, args...); err != nil {
		return
	}

	defer sqlRows.Close()

	return
}

// DeleteStudent delete student entry for an admission number
func (sr *StudentRepo) DeleteStudent(id string) (err error) {
	delb := sr.con.Builder.NewDeleteBuilder()
	delb.DeleteFrom(table).Where(delb.Equal("admission_no", id))
	sqlStr, args := delb.Build()
	log.Println(sqlStr, args)
	if _, err = sr.con.DB.Exec(sqlStr, args...); err != nil {
		return
	}

	return
}
