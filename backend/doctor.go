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

func doctorHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var doctor Doctor
	db.First(&doctor, "user_id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(&doctor)
}

func getDoctor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/get_doctor/")
	var doctor Doctor
	db.First(&doctor, "id == ?", id)
	json.NewEncoder(w).Encode(&doctor)
}

func doctorProfileRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024)
	var doctor Doctor
	var user User

	doctor.FirstName = r.FormValue("firstName")
	doctor.LastName = r.FormValue("lastName")
	doctor.Surname = r.FormValue("surname")
	doctor.Address = r.FormValue("address")
	doctor.Contact = r.FormValue("contact")
	doctor.Gender = r.FormValue("gender")
	doctor.DateOfBirth = r.FormValue("dob")
	doctor.BloodType = r.FormValue("bloodType")
	doctor.Department = r.FormValue("department")
	doctor.Specialty = r.FormValue("specialty")
	doctor.MedicalLicenceNumber = r.FormValue("medicalLicenceNumber")
	file, _, err := r.FormFile("profileImage")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.First(&user, "id == ?", currentUser.ID)

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

	result := db.Where(Doctor{UserID: user.ID}).FirstOrCreate(&Doctor{
		ID:                   uuid.NewString(),
		FirstName:            doctor.FirstName,
		LastName:             doctor.LastName,
		Surname:              doctor.Surname,
		Address:              doctor.Address,
		Gender:               doctor.Address,
		DateOfBirth:          doctor.DateOfBirth,
		BloodType:            doctor.BloodType,
		Department:           doctor.Department,
		Specialty:            doctor.Specialty,
		MedicalLicenceNumber: doctor.MedicalLicenceNumber,
		Contact:              doctor.Contact,
		UserID:               currentUser.ID,
		Email:                currentUser.Email,
		ProfileImage:         temp.Name(),
	})
	if result.Error != nil {
		http.Error(w, "doctor Registration failed", http.StatusBadRequest)
		return
	}
	db.Model(&user).Update("approved", true)
	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode("A Doctor With This ID Is Already Registered")
	} else {
		json.NewEncoder(w).Encode(&doctor)
	}
}

func doctorProfileUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateDoctor Doctor
	var doctor Doctor

	err := json.NewDecoder(r.Body).Decode(&updateDoctor)
	if err != nil {
		log.Fatal(err)
	}
	db.First(&doctor, "id == ?", updateDoctor.ID)

	if doctor.UserID != currentUser.ID {
		http.Error(w, "Unauthorized Access", http.StatusUnauthorized)
		return
	}

	db.Model(&doctor).Updates(&Doctor{
		FirstName:            updateDoctor.FirstName,
		LastName:             updateDoctor.LastName,
		Surname:              updateDoctor.Surname,
		Email:                updateDoctor.Email,
		DateOfBirth:          updateDoctor.DateOfBirth,
		Gender:               updateDoctor.Gender,
		Address:              updateDoctor.Address,
		BloodType:            updateDoctor.BloodType,
		Department:           updateDoctor.Department,
		Specialty:            updateDoctor.Specialty,
		MedicalLicenceNumber: updateDoctor.MedicalLicenceNumber,
		Contact:              updateDoctor.Contact,
	})
	json.NewEncoder(w).Encode(&doctor)
}

func doctorList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var doctors []Doctor
	db.Find(&doctors)
	json.NewEncoder(w).Encode(doctors)
}

func doctorTests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var test []PatientTest
	var doctor Doctor
	db.First(&doctor, "user_id == ?", currentUser.ID)
	db.Find(&test, "doctor_id == ?", doctor.ID)
	json.NewEncoder(w).Encode(test)
}

func doctorModelCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	type Count struct {
		AppointmentCount int64
		doctorCount      int64
		PatientCount     int64
	}
	var doctor Doctor
	var count Count
	db.First(&doctor, "user_id == ?", currentUser.ID)
	db.Find(&Appointment{}, "doctor_id == ?", doctor.ID).Count(&count.AppointmentCount)
	db.Find(&LabAdmin{}).Count(&count.doctorCount)
	db.Find(&PatientTest{}, "doctor_id == ?", doctor.ID).Count(&count.PatientCount)

	json.NewEncoder(w).Encode(count)
}

func doctorProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var doctor Doctor
	db.Find(&doctor, "user_id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(&doctor)
}
