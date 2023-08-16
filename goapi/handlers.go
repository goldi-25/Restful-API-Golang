package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	if emp.EmpSalary == 0 || emp.EmpName == "" || emp.Email == "" {
		http.Error(w, "Salary, name, and email must be provided", http.StatusBadRequest)
		return
	}
	Database.Create(&emp)
	json.NewEncoder(w).Encode(&emp)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["id"])

	json.NewEncoder(w).Encode(employee)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&employee)
	if employee.EmpSalary == 0 || employee.EmpName == "" || employee.Email == "" {
		http.Error(w, "Salary, name, and email must be provided", http.StatusBadRequest)
		return
	}
	Database.Save(&employee)
	json.NewEncoder(w).Encode(" Employee Updated successfully")
	json.NewEncoder(w).Encode(employee)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	Database.Delete(&emp, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(" Employee deleted successfully")
	json.NewEncoder(w).Encode(&emp)
}
