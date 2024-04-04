package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func createTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var test Test
	err := json.NewDecoder(r.Body).Decode(&test)

	if err != nil {
		log.Fatal(err)
		return
	}

	reslt := db.Create(&Test{
		ID:          uuid.NewString(),
		Title:       test.Title,
		Description: test.Description,
		Price:       test.Price,
	})
	if reslt.Error != nil {
		http.Error(w, "Test Registration failed", http.StatusBadRequest)
		log.Fatal(reslt.Error)
		return
	}
	json.NewEncoder(w).Encode(&test)
}

func allTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var test []Test

	db.Find(&test)
	json.NewEncoder(w).Encode(&test)

}

func getTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var test Test

	id := strings.TrimPrefix(r.URL.Path, "/get_test/")
	fmt.Println(id)
	db.First(&test, "id == ?", id)
	json.NewEncoder(w).Encode(&test)
}

func updateTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var test Test
	var updateTest Test

	json.NewDecoder(r.Body).Decode(&updateTest)
	db.First(&test, "id == ?", updateTest.ID)
	fmt.Println(updateTest.Price)

	db.Model(&test).Select("Title", "Description", "Price").Updates(&Test{
		Title:       updateTest.Title,
		Description: updateTest.Description,
		Price:       updateTest.Price,
	})

	json.NewEncoder(w).Encode(&test)
}

func deleteTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var test Test

	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		http.Error(w, "Failed to delete test", http.StatusNotImplemented)
		return
	}
	db.Where("id == ?", test.ID).Unscoped().Delete(Test{})

	json.NewEncoder(w).Encode(test)
}
