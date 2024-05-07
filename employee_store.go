package main

import (
	"sync"
)

// EmployeeStore represents an in-memory store for managing employees
type EmployeeStore struct {
	sync.Mutex
	employees map[int]Employee
}

// NewEmployeeStore creates a new instance of EmployeeStore
func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]Employee),
	}
}

// CreateEmployee adds a new employee to the store
func (es *EmployeeStore) CreateEmployee(employee Employee) {
	es.Lock()
	defer es.Unlock()
	es.employees[employee.ID] = employee
}

// GetEmployeeByID retrieves an employee from the store by ID
func (es *EmployeeStore) GetEmployeeByID(id int) (Employee, bool) {
	es.Lock()
	defer es.Unlock()
	employee, ok := es.employees[id]
	return employee, ok
}

// UpdateEmployee updates the details of an existing employee
func (es *EmployeeStore) UpdateEmployee(employee Employee) {
	es.Lock()
	defer es.Unlock()
	es.employees[employee.ID] = employee
}

// DeleteEmployee deletes an employee from the store by ID
func (es *EmployeeStore) DeleteEmployee(id int) {
	es.Lock()
	defer es.Unlock()
	delete(es.employees, id)
}
