package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func patientProfileRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024)

	var patient Patient
	var user User

	patient.FirstName = r.FormValue("firstName")
	patient.LastName = r.FormValue("lastName")
	patient.Surname = r.FormValue("surname")
	patient.Address = r.FormValue("address")
	patient.Gender = r.FormValue("gender")
	patient.BloodType = r.FormValue("bloodType")
	patient.DateOfBirth = r.FormValue("dob")
	patient.Contact = r.FormValue("contact")
	file, _, err := r.FormFile("profileImage")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()
	temp, err2 := os.CreateTemp("static", "image-*.jpg")
	if err2 != nil {
		log.Fatal(err2)
		return
	}
	defer temp.Close()

	fileBytes, err3 := io.ReadAll(file)
	if err3 != nil {
		log.Fatal(err3)
		return
	}
	temp.Write(fileBytes)

	db.First(&user, "id == ?", currentUser.ID)
	result := db.Where(Patient{UserID: user.ID}).FirstOrCreate(&Patient{
		ID:           uuid.NewString(),
		FirstName:    patient.FirstName,
		LastName:     patient.LastName,
		Surname:      patient.Surname,
		Address:      patient.Address,
		Gender:       patient.Gender,
		DateOfBirth:  patient.DateOfBirth,
		BloodType:    patient.BloodType,
		Contact:      patient.Contact,
		UserID:       currentUser.ID,
		Email:        currentUser.Email,
		ProfileImage: temp.Name(),
	})
	if result.Error != nil {
		http.Error(w, "Patient Registration failed", http.StatusBadRequest)
		return
	}

	db.Model(&user).Update("approved", true)
	json.NewEncoder(w).Encode(&patient)
}

func getPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/get_patient/")
	var patient Patient
	db.First(&patient, "id == ?", id)
	json.NewEncoder(w).Encode(&patient)
}

func patientProfileUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatePatient Patient
	var patient Patient

	err := json.NewDecoder(r.Body).Decode(&updatePatient)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&patient, "id == ?", updatePatient.ID)

	if patient.UserID != currentUser.ID {
		http.Error(w, "Unauthorized Access", http.StatusUnauthorized)
		return
	}
	db.Model(&patient).Updates(Patient{
		FirstName:   updatePatient.FirstName,
		LastName:    updatePatient.LastName,
		Surname:     updatePatient.Surname,
		Email:       updatePatient.Email,
		Gender:      updatePatient.Gender,
		DateOfBirth: updatePatient.DateOfBirth,
		BloodType:   updatePatient.BloodType,
		Address:     updatePatient.Address,
		Contact:     updatePatient.Contact,
	})
	json.NewEncoder(w).Encode(&patient)
}

func patientList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var patient []Patient
	db.Find(&patient)
	json.NewEncoder(w).Encode(&patient)
}

func patientHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var patient Patient
	db.First(&patient, "user_id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(&patient)
}

func patientModelCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	type Count struct {
		AppointmentCount int64
		PatientTestCount int64
		// PatientCount     int64
	}
	var patient Patient
	var count Count
	db.First(&patient, "user_id == ?", currentUser.ID)
	db.Find(&Appointment{}, "patient_id == ?", patient.ID).Count(&count.AppointmentCount)
	db.Find(&PatientTest{}, "patient_id == ?", patient.ID).Count(&count.PatientTestCount)
	// db.Find(&Patient{}).Count(&count.PatientCount)
	json.NewEncoder(w).Encode(count)
}

func patientProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var patient Patient
	db.Find(&patient, "user_id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(&patient)
}
