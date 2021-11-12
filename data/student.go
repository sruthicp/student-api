package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"student-api/db"
)

const table = "school.student"

type Student struct {
	AdmNo   int64  `json:"admission_no"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Class   string `json:"class"`
	Age     int64  `json:"age"`
}

func (sd *Student) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(sd)
}

// GetStudent returns student details
func GetStudent(con *db.DBConnection, id int64) (std Student, err error) {
	var sqlRows *sql.Rows

	sqlStr := fmt.Sprintf("select admission_no, std_name, std_address, std_class, std_age from school.student where admission_no = '%d'", id)
	log.Printf("query : %s", sqlStr)

	if sqlRows, err = con.DB.Query(sqlStr); err != nil {
		return
	}

	defer sqlRows.Close()
	for sqlRows.Next() {
		if err = sqlRows.Scan(&std.AdmNo, &std.Name, &std.Address, &std.Class, &std.Age); err != nil {
			return
		}
	}
	if std == (Student{}) {
		return std, sql.ErrNoRows
	}
	return
}

// AddStudent save student details
func AddStudent(con *db.DBConnection, std Student) (err error) {
	var sqlRows *sql.Rows

	ib := con.Builder.NewInsertBuilder()
	ib.InsertInto(table).Cols("admission_no", "std_name", "std_address", "std_class", "std_age").Values(std.AdmNo, std.Name, std.Address, std.Class, std.Age)

	sqlStr, args := ib.Build()
	log.Println(sqlStr, args)
	if sqlRows, err = con.DB.Query(sqlStr, args...); err != nil {
		return
	}

	defer sqlRows.Close()

	return
}

// UpdateStudent update student details for an Admission number
func UpdateStudent(con *db.DBConnection, id int64, address, class string, age int64) (err error) {
	var sqlRows *sql.Rows
	var setAssigns []string

	ub := con.Builder.NewUpdateBuilder()
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
	if sqlRows, err = con.DB.Query(sqlStr, args...); err != nil {
		return
	}

	defer sqlRows.Close()

	return
}

// DeleteStudent delete student entry for an admission number
func DeleteStudent(con *db.DBConnection, id int64) (err error) {
	delb := con.Builder.NewDeleteBuilder()
	delb.DeleteFrom(table).Where(delb.Equal("admission_no", id))
	sqlStr, args := delb.Build()
	log.Println(sqlStr, args)
	if _, err = con.DB.Exec(sqlStr, args...); err != nil {
		return
	}

	return
}
