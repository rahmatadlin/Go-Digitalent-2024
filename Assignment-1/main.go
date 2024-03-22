package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"github.com/rahmatadlin/Go-Digitalent-2024/Assignment-1/db"
)

func main() {
	studentDb := db.StudentDb{}
	initDbData(&studentDb)

	id := os.Args

	if len(id) < 2 {
		err := errors.New("error: user Id not provided, please insert user Id")
		fmt.Println(err.Error())
	} else {
		userId, _ := strconv.Atoi(id[1])
		printStudentById(&studentDb, userId)
	}
}

func initDbData(studentDb *db.StudentDb) {
	rawDatas := [][]string{
		{"Bambang Waluyo", "Semarang", "Software Engineer", "Belajar hal yang baru"},
		{"Joko Samsudin", "Bandung", "Data Analyst", "Iseng"},
		{"Nana", "Jakarta", "UI/UX Designer", "Switch Career"},
		{"Asep Seapudin", "Pontianak", "Product Manager", "Ngisi Waktu"},
		{"Budi Santoso", "Madura", "DevOps Engineer", "Cari pengalaman aja"},
	}

	for _, data := range rawDatas {
		student := db.Student{Name: data[0], Address: data[1], Work: data[2], JoinReason: data[3]}
		studentDb.AddStudent(student)
	}
}

func printStudentById(studentDb *db.StudentDb, id int) {
	student, err := studentDb.GetStudentById(id)

	if err == nil {
		student.PrintStudent()
	} else {
		fmt.Println(err.Error())
	}
}