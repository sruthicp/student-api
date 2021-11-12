package data

import (
	"fmt"
	"os"
	"student-api/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStudent(t *testing.T) {
	var postgres *db.DBConnection
	os.Setenv("DB_HOST", "localhost")

	postgres, _ = db.NewDBConnection()

	type args struct {
		con *db.DBConnection
		id  int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid case",
			args: args{
				con: postgres,
				id:  2,
			},
		},
	}
	for _, tt := range tests {
		sd, err := GetStudent(tt.args.con, tt.args.id)
		assert.NotNil(t, sd)
		fmt.Println("sudent:", sd)
		assert.NoError(t, err)
	}
}

func TestAddStudent(t *testing.T) {
	var postgres *db.DBConnection

	postgres, _ = db.NewDBConnection()

	type args struct {
		con *db.DBConnection
		std Student
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid case",
			args: args{
				con: postgres,
				std: Student{
					AdmNo:   1,
					Name:    "Sruthi",
					Address: "Kerala",
					Class:   "CSE",
					Age:     28,
				},
			},
		},
	}
	for _, tt := range tests {
		err := AddStudent(tt.args.con, tt.args.std)
		assert.NoError(t, err)
	}
}
