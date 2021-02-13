package employee

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

//import  "database/sql"
import _ "github.com/lib/pq"

func (database *Database) CreateConnection() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostname, port, username, password, dbName)
	db, err := sql.Open(POSTGRES_DRIVER_NAME, connStr)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	dbConn = db
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Println("Connected with db successfully")
	return nil
}

func (database *Database) closeConnection() error {
	log.Println("Closing DB connection.")
	return dbConn.Close()
}

func (database *Database) InsertEntry(emp Employee) (int, error) {
	insertQuery := `
    INSERT INTO employee (name, age, department, annual_salary, social_security_number, address, created_at, last_updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id`
	id := 0
	err := dbConn.QueryRow(insertQuery, emp.Name, emp.Age, emp.Department, emp.AnnualSalary, emp.SocialSecurityNumber, emp.Address, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (database *Database) QueryById(id int) (*Employee, error) {
	query := `
	SELECT id,name, age, department, annual_salary, social_security_number, address FROM employee where id=$1;
`
	var empId int
	var name string
	var age int
	var department string
	var annualSalary int
	var socialSecurityNumber string
	var address string
	row := dbConn.QueryRow(query, id)
	switch err := row.Scan(&empId, &name, &age, &department, &annualSalary, &socialSecurityNumber, &address); err {
	case sql.ErrNoRows:
		log.Println("No records found!")
		return nil, fmt.Errorf("No record found woth id: %d", id)
	case nil:
		employee := Employee{
			Name:                 name,
			Age:                  age,
			Department:           department,
			AnnualSalary:         annualSalary,
			SocialSecurityNumber: socialSecurityNumber,
			Address:              address,
		}
		log.Printf("Employee: %+v", employee)
		return &employee, nil
	default:
		return nil, fmt.Errorf("Error: %s", err.Error())
	}
}

func (database *Database) QueryBySearchName(filter SearchFilter) ([]*Employee, error) {
	query := `
	SELECT id,name, age, department, annual_salary, social_security_number, address FROM employee where name=$1;
`
	rows, err := dbConn.Query(query, filter.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []*Employee
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.Id, &e.Name, &e.Age, &e.Department, &e.AnnualSalary, &e.SocialSecurityNumber, &e.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &e)
	}
	return employees, nil
}

func (database *Database) UpdateRecord(employee *Employee) error {
	updateQuery := `
	UPDATE employee 
	SET annual_salary = $1, last_updated_at=$2
    WHERE id=$3;
`
	_, err := dbConn.Exec(updateQuery, employee.AnnualSalary, time.Now(), employee.Id)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) DeleteEmployee(id int) error {
	deleteQuery := `
	DELETE from employee
	WHERE id=$1;
`
	_, err := dbConn.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	log.Printf("Employee with id: %d deleted\n", id)
	return nil
}
