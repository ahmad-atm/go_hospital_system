package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func requestAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var appointment Appointment
	var patient Patient
	var doctor Doctor
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		log.Fatalln(err)
	}
	db.First(&patient, "user_id == ?", currentUser.ID)
	db.First(&doctor, "id == ?", appointment.DoctorId)
	var newAppointment = Appointment{
		ID:               uuid.NewString(),
		PatientID:        patient.ID,
		PatientFirstName: patient.FirstName,
		PatientLastName:  patient.LastName,
		PatientProfile:   patient.ProfileImage,
		DoctorId:         appointment.DoctorId,
		DoctorProfile:    appointment.DoctorProfile,
		DoctorFirstName:	doctor.FirstName,
		DoctorLastName:		doctor.LastName,
		Purpose:          appointment.Purpose,
		Status:           "Pending"}
	db.Create(&newAppointment)
	json.NewEncoder(w).Encode(newAppointment)
}

func getAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var appointment Appointment

	id := strings.TrimPrefix(r.URL.Path, "/get_appointment/")
	db.First(&appointment, "id == ?", id)
	json.NewEncoder(w).Encode(&appointment)
}

func limitedDoctorAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var appointements []Appointment
	var doctor Doctor

	db.First(&doctor, "user_id == ?", currentUser.ID)
	db.Find(&appointements, "doctor_id == ?", doctor.ID).Limit(3)
	json.NewEncoder(w).Encode(appointements)
}

func doctorAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var appointments []Appointment
	var doctor Doctor

	db.First(&doctor, "user_id == ?", currentUser.ID)
	db.Find(&appointments, "doctor_id == ?", doctor.ID)
	json.NewEncoder(w).Encode(appointments)
}

func handleAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var patient Patient
	var appointment Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		log.Fatalln(err)
	}

	db.First(&patient, "user_id == ?", currentUser.ID)
	db.Model(&appointment).Updates(&Appointment{
		Purpose:   appointment.Purpose,
		Status:    "Pending",
		DoctorId:  appointment.DoctorId,
		Day:       appointment.Day,
		Date:      appointment.Date,
		StartTime: appointment.StartTime,
		StopTime:  appointment.StopTime,
	})
	json.NewEncoder(w).Encode(appointment)
}

// Patient Appointments
func patientAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var appointments []Appointment
	var patient Patient
	db.First(&patient, "user_id == ?", currentUser.ID)
	db.Find(&appointments, "patient_id == ?", patient.ID)
	json.NewEncoder(w).Encode(appointments)
}


func declineAppointment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	var appointment Appointment

	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		log.Fatalln(err)
	}

	db.First(&appointment, "id == ?", appointment.ID)
		db.Model(&appointment).Select("Status").Updates(&Appointment{
		Status:       "Declined",
	})
	json.NewEncoder(w).Encode(&appointment)
}

func acceptAppointment(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	var appointment Appointment

	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		log.Fatalln(err)
	}

	db.First(&appointment, "id == ?", appointment.ID)
		db.Model(&appointment).Select("Status").Updates(&Appointment{
		Status:       "Accepted",
	})
	json.NewEncoder(w).Encode(&appointment)
}