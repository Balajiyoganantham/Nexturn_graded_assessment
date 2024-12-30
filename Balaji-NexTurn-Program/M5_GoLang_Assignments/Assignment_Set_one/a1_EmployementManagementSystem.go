// // Exercise 1: Employee Management System
// // Topics Covered: Go Conditions, Go Loops, Go Constants, Go Functions, Go Arrays, Go
// // Strings, Go Errors
// // Case Study:
// // A company wants to manage its employees' data in memory. Each employee has an ID,
// // name, age, and department. You need to build a small application that performs the
// // following:
// // 1. Add Employee: Accept input for employee details and store them in an array of
// // structs. Validate the input:
// // o ID must be unique.
// // o Age should be greater than 18. If validation fails, return custom error
// // messages.
// // 2. Search Employee: Search for an employee by ID or name using conditions.
// // Return the details if found, or return an error if not found.
// // 3. List Employees by Department: Use loops to filter and display all employees in
// // a given department.
// // 4. Count Employees: Use constants to define a department (e.g., "HR", "IT"), and
// // display the count of employees in that department.

package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

func addEmployee(id int, name string, age int, department string) error {
	// Validate ID uniqueness
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	// Validate age
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	return nil
}

func searchEmployee(query string) (Employee, error) {
	for _, emp := range employees {
		if fmt.Sprintf("%d", emp.ID) == query || emp.Name == query {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

func listEmployeesByDepartment(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if emp.Department == department {
			result = append(result, emp)
		}
	}
	return result
}

func countEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if emp.Department == department {
			count++
		}
	}
	return count
}

func main() {
	// Constants for departments
	const (
		HR      = "HR"
		IT      = "IT"
		Finance = "Finance"
	)

	// Adding employees
	if err := addEmployee(1, "Alice", 30, HR); err != nil {
		fmt.Println("Error:", err)
	}
	if err := addEmployee(2, "Bob", 25, IT); err != nil {
		fmt.Println("Error:", err)
	}
	if err := addEmployee(3, "Charlie", 22, IT); err != nil {
		fmt.Println("Error:", err)
	}

	// Search for an employee
	emp, err := searchEmployee("Alice")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found Employee: %+v\n", emp)
	}

	// List employees by department
	fmt.Println("Employees in IT Department:")
	for _, emp := range listEmployeesByDepartment(IT) {
		fmt.Printf("%+v\n", emp)
	}

	// Count employees in HR
	fmt.Printf("Number of employees in HR: %d\n", countEmployees(HR))
}
