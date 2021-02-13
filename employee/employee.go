package employee

import (
	"context"
	"log"
	"time"
)

type Employee struct {
	Id                   int       `json:"id"`
	Name                 string    `json:"name"`
	Age                  int       `json:"age"`
	Department           string    `json:"department"`
	AnnualSalary         int       `json:"annual_salary"`
	SocialSecurityNumber string    `json:"social_security_number"`
	Address              string    `json:"address, omitempty"`
	CreatedAt            time.Time `json:"created_at"`
	LastUpdatedAt        time.Time `json:"last_updated_at"`
}

type SearchFilter struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type EmpService struct {
}

type EmployeeService interface {
	FindEmployeeById(ctx context.Context, empId int) (*Employee, error)
	FindEmployees(ctx context.Context, filter SearchFilter) ([]*Employee, error)
	CreateEmployee(ctx context.Context, employee Employee) (int, error)
	UpdateEmployee(ctx context.Context, employee *Employee) error
	DeleteEmployee(ctx context.Context, empId int) error
}

var d Database

func initializeDBConn() error {
	d = Database{}
	if err := d.CreateConnection(); err != nil {
		log.Fatalf("Error: %s", err.Error())
		return err
	}
	return nil
}
func releaseConn() {
	d.closeConnection()
}

func (s *EmpService) CreateEmployee(ctx context.Context, employee Employee) (int, error) {
	if err := initializeDBConn(); err != nil {
		return 0, err
	}
	defer releaseConn()
	return d.InsertEntry(employee)
}

func (s *EmpService) FindEmployeeById(ctx context.Context, empId int) (*Employee, error) {
	if err := initializeDBConn(); err != nil {
		return nil, err
	}
	defer releaseConn()
	return d.QueryById(empId)
}

func (s *EmpService) FindEmployees(ctx context.Context, filter SearchFilter) ([]*Employee, error) {
	if err := initializeDBConn(); err != nil {
		return nil, err
	}
	defer releaseConn()
	return d.QueryBySearchName(filter)
}

func (s *EmpService) UpdateEmployee(ctx context.Context, employee *Employee) error {
	if err := initializeDBConn(); err != nil {
		return err
	}
	defer releaseConn()
	return d.UpdateRecord(employee)
}

func (s *EmpService) DeleteEmployee(ctx context.Context, empId int) error {
	if err := initializeDBConn(); err != nil {
		return err
	}
	defer releaseConn()
	return d.DeleteEmployee(empId)
}
