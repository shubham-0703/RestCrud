package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCRUDOperations(t *testing.T) {
	store := NewEmployeeStore()

	// Create an employee
	employee := Employee{ID: 1, Name: "John Doe", Position: "Software Engineer", Salary: 50000}
	store.CreateEmployee(employee)

	// Retrieve the employee
	retrievedEmployee, ok := store.GetEmployeeByID(1)
	if !ok {
		t.Errorf("Failed to retrieve employee by ID")
	}

	// Update the employee
	retrievedEmployee.Position = "Senior Software Engineer"
	store.UpdateEmployee(retrievedEmployee)

	// Verify the update
	updatedEmployee, _ := store.GetEmployeeByID(1)
	if updatedEmployee.Position != "Senior Software Engineer" {
		t.Errorf("Failed to update employee details")
	}

	// Delete the employee
	store.DeleteEmployee(1)

	// Verify deletion
	_, ok = store.GetEmployeeByID(1)
	if ok {
		t.Errorf("Failed to delete employee")
	}
}

func TestListEmployeesHandler(t *testing.T) {
	store := NewEmployeeStore()
	api := NewAPI(store)

	// Create some sample employees
	employees := []Employee{
		{ID: 1, Name: "John Doe", Position: "Software Engineer", Salary: 50000},
		{ID: 2, Name: "Jane Smith", Position: "Project Manager", Salary: 60000},
		{ID: 3, Name: "Alice Johnson", Position: "Data Scientist", Salary: 70000},
		{ID: 4, Name: "Bob Brown", Position: "DevOps Engineer", Salary: 55000},
	}

	for _, employee := range employees {
		store.CreateEmployee(employee)
	}

	// Test pagination
	req, err := http.NewRequest("GET", "/employees?page=1&pageSize=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.ListEmployeesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"ID":1,"Name":"John Doe","Position":"Software Engineer","Salary":50000},{"ID":2,"Name":"Jane Smith","Position":"Project Manager","Salary":60000}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
