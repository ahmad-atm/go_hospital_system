package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func labHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lab LabAdmin
	db.First(&lab, "user_id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(&lab)
}

func labProfileRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 10024)
	var labAdmin LabAdmin
	var user User

	labAdmin.FirstName = r.FormValue("firstName")
	labAdmin.LastName = r.FormValue("lastName")
	labAdmin.Surname = r.FormValue("surname")
	labAdmin.Address = r.FormValue("address")
	labAdmin.Contact = r.FormValue("contact")
	file, _, err := r.FormFile("profileImage")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.First(&user, "id == ?", currentUser.ID)

	defer file.Close()
	temp, err2 := os.CreateTemp("static", "image-*.jpg")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer temp.Close()

	fileBytes, err3 := io.ReadAll(file)
	if err3 != nil {
		log.Fatal(err3)
		return
	}
	temp.Write(fileBytes)

	result := db.Where(LabAdmin{UserID: user.ID}).FirstOrCreate(&LabAdmin{
		ID:           uuid.NewString(),
		FirstName:    labAdmin.FirstName,
		LastName:     labAdmin.LastName,
		Surname:      labAdmin.Surname,
		Address:      labAdmin.Address,
		Contact:      labAdmin.Contact,
		UserID:       currentUser.ID,
		Email:        currentUser.Email,
		ProfileImage: temp.Name(),
	})
	if result.Error != nil {
		http.Error(w, "Lab Admin Registration failed", http.StatusBadRequest)
		return
	}
	db.Model(&user).Update("approved", true)
	json.NewEncoder(w).Encode(&labAdmin)
}

func labProfileUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateLab LabAdmin
	var lab LabAdmin
	err := json.NewDecoder(r.Body).Decode(&updateLab)
	if err != nil {
		log.Fatal(err)
		return
	}

	db.First(&lab, "id == ?", currentUser.ID)
	db.Model(&lab).Updates(&LabAdmin{
		FirstName: updateLab.FirstName,
		LastName:  updateLab.LastName,
		Surname:   updateLab.Surname,
		Contact:   updateLab.Contact,
		Address:   updateLab.Address,
	})
	json.NewEncoder(w).Encode(&lab)
}

func labModelCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Count struct {
		DoctorCount   int64
		LabAdminCount int64
		PatientCount  int64
	}
	var count Count

	db.Find(&Doctor{}).Count(&count.DoctorCount)
	db.Find(&LabAdmin{}).Count(&count.LabAdminCount)
	db.Find(&Patient{}).Count(&count.PatientCount)

	json.NewEncoder(w).Encode(count)
}

func labProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var labAdmin LabAdmin
	db.Find(&labAdmin, "user_id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(&labAdmin)
}
