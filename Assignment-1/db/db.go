package db

import (
	"errors"
	"fmt"
)

type StudentDb struct {
	students []Student
}

func (db *StudentDb) AddStudent(student Student) {
	id := len(db.students) + 1
	student.ID = id
	db.students = append(db.students, student)
}

func (db *StudentDb) GetStudents() ([]Student, error) {
	return db.students, nil
}

func (db *StudentDb) GetStudentById(id int) (Student, error) {
	for _, student := range db.students {
		if student.ID == id {
			return student, nil
		}
	}
	return Student{}, errors.New("error: data not found")
}

type Student struct {
	ID         int
	Name       string
	Address    string
	Work       string
	JoinReason string
}

func (studentData *Student) PrintStudent() {
	fmt.Println("Name : ", studentData.Name)
	fmt.Println("Address : ", studentData.Address)
	fmt.Println("Work : ", studentData.Work)
	fmt.Println("Join Reason : ", studentData.JoinReason)
}