package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func createPatientTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lab LabAdmin
	var test PatientTest
	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&lab, "user_id == ?", currentUser.ID)

	result := db.Create(&PatientTest{
		ID:                uuid.NewString(),
		Title:             test.Title,
		Description:       test.Description,
		PatientID:         test.PatientID,
		PatientFirstName:  test.PatientFirstName,
		PatientLastName:   test.PatientLastName,
		PatientSurname:    test.PatientSurname,
		PatientImage:      test.PatientImage,
		Price:             test.Price,
		LabAdminID:        lab.ID,
		LabAdminFirstName: lab.FirstName,
		LabAdminLastName:  lab.LastName,
		LabAdminSurname:   lab.Surname,
		DoctorId:          test.DoctorId,
		DoctorFirstName:   test.DoctorFirstName,
		DoctorLastName:    test.DoctorLastName,
		DoctorSurname:     test.DoctorSurname,
		Status:            "Pending",
	})
	if result.Error != nil {
		http.Error(w, "Test Registration failed", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(test)
}

func getPatientTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var patientTest PatientTest
	id := strings.TrimPrefix(r.URL.Path, "/get_patient_test/")
	db.First(&patientTest, "id == ?", id)
	json.NewEncoder(w).Encode(patientTest)
}

func updatePatientTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateTest PatientTest
	var test PatientTest

	err := json.NewDecoder(r.Body).Decode(&updateTest)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&test, "id == ?", updateTest.ID)

	db.Model(&test).Updates(&PatientTest{
		Title:            updateTest.Title,
		Description:      updateTest.Description,
		PatientID:        updateTest.PatientID,
		PatientFirstName: updateTest.PatientFirstName,
		PatientLastName:  updateTest.PatientLastName,
		PatientSurname:   updateTest.PatientSurname,
		Price:            updateTest.Price,
		DoctorId:         updateTest.DoctorId,
		DoctorFirstName:  updateTest.DoctorFirstName,
		DoctorLastName:   updateTest.DoctorLastName,
		DoctorSurname:    updateTest.DoctorLastName,
		Status:           updateTest.Status,
	})
	json.NewEncoder(w).Encode(&test)
}

func patientTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var test []PatientTest
	var patient Patient
	db.First(&patient, "user_id == ?", currentUser.ID)
	db.Find(&test, "patient_id == ?", patient.ID)
	json.NewEncoder(w).Encode(test)
}
func patientTestHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var test []PatientTest

	db.Find(&test).Limit(3)
	json.NewEncoder(w).Encode(test)
}
