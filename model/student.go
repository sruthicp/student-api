package model

import (
	"encoding/json"
	"io"
)

type Student struct {
	tableName struct{} `sql:"student" pg:",discard_unknown_columns"` //nolint
	AdmNo     string   `sql:"admission_no,pk" json:"admission_no"`
	Name      string   `sql:"std_name" json:"name"`
	Address   string   `sql:"std_address" json:"address"`
	Class     string   `sql:"std_class" json:"class"`
	Age       int32    `sql:"std_age" json:"age"`
}

func (sd *Student) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(sd)
}
