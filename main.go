package main

import (
	"context"
	"github.com/devgit072/crud-design-in-go/employee"
	"log"
	"time"
)

func main() {
	//create()
	find()
	//findByName()
	//update()
	//deleteEmp()
}

func create() {
	e := employee.Employee{
		Name:                 "Nitin",
		Age:                  30,
		Department:           "CS",
		AnnualSalary:         1000,
		SocialSecurityNumber: "hhggs-32899-shh",
		Address:              "India",
		CreatedAt:            time.Time{},
		LastUpdatedAt:        time.Time{},
	}
	var service employee.EmployeeService = &employee.EmpService{}
	id, err := service.CreateEmployee(context.Background(), e)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("Employee created succesfully with id:", id)
}

func find() {
	var service employee.EmployeeService = &employee.EmpService{}
	e, err := service.FindEmployeeById(context.Background(), 2)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("Employee with id: %d, \nEmployee: %+v\n", 2, *e)
}

func findByName() {
	filter := employee.SearchFilter{
		Name: "Devraj",
	}
	var service employee.EmployeeService = &employee.EmpService{}
	employees, err := service.FindEmployees(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	for _, e := range employees {
		log.Printf("Employee: %+v\n", *e)
	}
}

func update() {
	e := employee.Employee{
		Id:           1,
		AnnualSalary: 2000,
	}
	var service employee.EmployeeService = &employee.EmpService{}
	if err := service.UpdateEmployee(context.Background(), &e); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func deleteEmp() {
	var service employee.EmployeeService = &employee.EmpService{}
	if err := service.DeleteEmployee(context.Background(), 2); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
