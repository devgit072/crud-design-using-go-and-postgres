<h2>Introduction    </h2>
A simple project demonstrating the best practices around designing CRUD api in Go. We will be creating crud functions for 
doing CRUD operations of Employees in Postgres DB.

<hr>
First of all we will define all crud functions in interface and then <b>Employee</b> struct will implement those.

A standard signature for all typical crud functions will look like this:
```go
type EmployeeService interface {
	FindEmployeeById(ctx context.Context, empId int) (*Employee, error)
	FindEmployees(ctx context.Context, filter SearchFilter) ([]*Employee, error)
	CreateEmployee(ctx context.Context, employee Employee) (int, error)
	UpdateEmployee(ctx context.Context, employee *Employee) error
	DeleteEmployee(ctx context.Context, empId int) error
}
```

There are two library out their for Go to work with postgres db:
https://github.com/jackc/pgx and https://github.com/lib/pq

https://github.com/lib/pq is more used but https://github.com/jackc/pgx has become very popular due to its performance.

However, we will just https://github.com/lib/pq for our projects here.

This projects just demonstrate how to do crud operations in postgres DB.