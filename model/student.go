package model

import (
	"encoding/json"
	"io"
)

type Student struct {
	AdmNo   string `json:"admission_no"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Class   string `json:"class"`
	Age     int32  `json:"age"`
}

func (sd *Student) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(sd)
}
