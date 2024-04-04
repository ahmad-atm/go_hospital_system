package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Approved bool   `json:"approved"`
}

type LoggedUser struct {
	ID       string
	Email    string
	Role     string
	Approved bool
}

type Doctor struct {
	gorm.Model
	ID                   string `gorm:"primaryKey"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Surname              string `json:"surname"`
	Email                string `json:"email"`
	DateOfBirth          string `json:"dob"`
	Gender               string `json:"gender"`
	Address              string `json:"address"`
	BloodType            string `json:"bloodType"`
	Department           string `json:"department"`
	Specialty            string `json:"specialty"`
	MedicalLicenceNumber string `json:"medicalLicenceNumber"`
	Contact              string `json:"contact"`
	ProfileImage         string `json:"tou"`
	UserID               string `json:"userId"`
}

type Patient struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
	DateOfBirth  string `json:"dob"`
	BloodType    string `json:"bloodType"`
	Address      string `json:"address"`
	Contact      string `json:"contact"`
	ProfileImage string `json:"profileImage"`
	UserID       string `json:"userId"`
}

type MedicalHistory struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	DoctorId     string `json:"doctorId"`
	PatientId    string `json:"patientId"`
	PatientName  string `json:"patientName"`
	PatientImage string `json:"patientImage"`
	Title        string `json:"title"`
	Medication   string `json:"medication"`
}

type LabAdmin struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Surname      string `json:"surname"`
	Contact      string `json:"contact"`
	Address      string `json:"address"`
	UserID       string `json:"userId"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileImage"`
}

type PatientTest struct {
	gorm.Model
	ID                string `gorm:"primaryKey"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	PatientID         string `json:"patientId"`
	PatientFirstName  string `json:"patientFirstName"`
	PatientLastName   string `json:"patientLastName"`
	PatientSurname    string `json:"patientSurname"`
	PatientImage      string `json:"patientImage"`
	Price             string `json:"price"`
	LabAdminID        string `json:"labAdminId"`
	LabAdminFirstName string `json:"labAdminFirstName"`
	LabAdminLastName  string `json:"labAdminLastName"`
	LabAdminSurname   string `json:"labAdminSurname"`
	DoctorId          string `json:"doctorId"`
	DoctorFirstName   string `json:"doctorFirstName"`
	DoctorLastName    string `json:"doctorLastName"`
	DoctorSurname     string `json:"doctorSurname"`
	Status            string `json:"status"`
}

type Test struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type Schedule struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Day      string `json:"date"`
	From     string `json:"from"`
	To       string `json:"to"`
	DoctorId string `json:"doctorId"`
}

type Appointment struct {
	gorm.Model
	ID               string `gorm:"primaryKey"`
	Day              string `json:"day"`
	Date             string `json:"date"`
	StartTime        string `json:"startTime"`
	StopTime         string `json:"stopTime"`
	PatientID        string `json:"patientId"`
	PatientFirstName string `json:"patientFirstName"`
	PatientLastName  string `json:"patientLastName"`
	PatientProfile   string `json:"patientProfile"`
	DoctorId         string `json:"doctorId"`
	DoctorProfile    string `json:"doctorProfile"`
	DoctorFirstName string `json:"doctorFirstName"`
	DoctorLastName  string `json:"doctorLastName"`
	Purpose          string `json:"purpose"`
	Status           string `json:"status"`
}
