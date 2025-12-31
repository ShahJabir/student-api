package storage

import "github.com/ShahJabir/student-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(Id int64) (types.Student, error)
}
