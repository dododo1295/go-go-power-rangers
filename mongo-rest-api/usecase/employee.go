package usecase

import (
	"encoding/json"
	"log"
	"main/model"
	"main/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error decoding request body: ", err)
		res.Error = err.Error()
		return
	}
	//employee assignment
	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// insert employee

	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error inserting employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("employee inserted with id: ", insertID)
}
func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get employee ID
	empID := mux.Vars(r)["id"]

	log.Println("employee ID: ", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeByID(empID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error finding employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)

}
func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEmployee()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error finding employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get employee ID
	empID := mux.Vars(r)["id"]
	log.Println("employee ID: ", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee ID")
		res.Error = "invalid employee ID"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error decoding request body: ", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = empID

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error updating employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)

}
func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get employee ID
	empID := mux.Vars(r)["id"]
	log.Println("employee ID: ", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteEmployeeByID(empID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error deleting employee: ", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteAllEmployee()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error deleting employee: ", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
