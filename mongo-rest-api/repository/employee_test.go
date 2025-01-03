package repository

import (
	"context"
	"log"
	"main/model"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error while connecting mongodb", err)
	}

	log.Println("Connected to MongoDB!")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	log.Println("Ping to MongoDB!")

	return mongoTestClient

}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	coll := mongoTestClient.Database("test").Collection("testEmployee")

	empRepo := EmployeeRepo{MongoCollection: coll}

	// inserting employee
	t.Run("InsertEmployee", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Tony Stark",
			Department: "Engineering",
			EmployeeID: emp1,
		}
		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("insert employee failed", err)
		}
		t.Log("inserted employee id: ", result)
	})

	t.Run("InsertEmployee", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Bob Ross",
			Department: "Art",
			EmployeeID: emp2,
		}
		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("insert employee failed", err)
		}
		t.Log("inserted employee id: ", result)
	})

	// getting data

	t.Run("GetEmployeeByID", func(t *testing.T) {
		emp, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal("get employee by id failed", err)
		}
		t.Log("employee data: ", emp)
	})

	// getting all data
	t.Run("GetAllEmployee", func(t *testing.T) {
		emps, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("get all employee failed", err)
		}
		t.Log("all employee data: ", emps)
	})

	// updating data

	t.Run("UpdateEmployeeByID", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Tony Stark AKA Iron Man",
			Department: "Engineering",
			EmployeeID: emp1,
		}

		result, err := empRepo.UpdateEmployeeByID(emp1, &emp)

		if err != nil {
			log.Fatal("update employee failed", err)
		}

		t.Log("update count:", result)
	})

	// checking the update
	t.Run("GetEmployeeByID", func(t *testing.T) {
		emp, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal("get employee by id failed", err)
		}
		t.Log("employee data: ", emp)
	})

	// deleting data
	t.Run("DeleteEmployeeByID", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp2)

		if err != nil {
			t.Fatal("delete employee failed", err)
		}
		t.Log("delete count:", result)
	})

}
