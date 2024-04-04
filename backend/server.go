package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("site.db"), &gorm.Config{})
var jwtSecret = []byte("mysecretkey")
var currentUser LoggedUser

func main() {

	dbMigration()
	if err != nil {
		panic("failed to connect database")
	}

	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	handler := cors.Handler(mux)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/", isAuthenticated(home)) // Done
	mux.HandleFunc("/register", register)  //Done
	mux.HandleFunc("/login", login)        //Done

	//Lab Routes
	mux.Handle("/lab_home", isAuthenticated(labHome))                        //Done
	mux.Handle("/lab_profile_register", isAuthenticated(labProfileRegister)) //Done
	mux.Handle("/lab_profile_update", isAuthenticated(labProfileUpdate))
	mux.Handle("/lab_profile", isAuthenticated(labProfile))

	//Doctor Routes
	mux.Handle("/doctor_profile_register", isAuthenticated(doctorProfileRegister))
	mux.Handle("/doctor_home", isAuthenticated(doctorHome)) //
	mux.Handle("/doctor_profile_update", isAuthenticated(doctorProfileUpdate))
	mux.Handle("/doctor_list", isAuthenticated(doctorList))
	mux.Handle("/doctor_tests", isAuthenticated(doctorTests))
	mux.Handle("/get_doctor/", isAuthenticated(getDoctor))
	mux.Handle("/doctor_profile", isAuthenticated(doctorProfile))

	//Patient Routes
	mux.Handle("/patient_home", isAuthenticated(patientHome))                        //Done
	mux.Handle("/get_patient/", isAuthenticated(getPatient))                         //Done
	mux.Handle("/patient_profile_register", isAuthenticated(patientProfileRegister)) //Done
	mux.Handle("/patient_profile_update", isAuthenticated(patientProfileUpdate))
	mux.Handle("/patient_list", isAuthenticated(patientList)) //Done
	mux.Handle("/patient_profile", isAuthenticated(patientProfile))

	//Patient Tests Route
	mux.Handle("/get_patient_test/", isAuthenticated(getPatientTest))
	mux.Handle("/update_patient_test", isAuthenticated(updatePatientTest))
	mux.Handle("/create_patient_test", isAuthenticated(createPatientTest))
	mux.Handle("/patient_tests", isAuthenticated(patientTest))

	// Test Route
	mux.Handle("/get_test/", isAuthenticated(getTest))
	mux.Handle("/all_tests", isAuthenticated(allTest))
	mux.Handle("/create_test", isAuthenticated(createTest))
	mux.Handle("/update_test", isAuthenticated(updateTest))
	mux.Handle("/delete_test", isAuthenticated(deleteTest))
	mux.Handle("/patients_test_history", isAuthenticated(patientTestHistory))

	//Appointments Route
	mux.Handle("/request_appointment", isAuthenticated(requestAppointment))
	mux.Handle("/handle_appointment", isAuthenticated(handleAppointment))
	mux.Handle("/get_appointment/", isAuthenticated(getAppointment))
	mux.Handle("/doctor_appointments", isAuthenticated(doctorAppointments))
	mux.Handle("/limited_doctor_appointments", isAuthenticated(limitedDoctorAppointments))
	mux.Handle("/patient_appointments", isAuthenticated(patientAppointments))
	mux.Handle("/decline_appointment", isAuthenticated(declineAppointment))
	mux.Handle("/accept_appointment", isAuthenticated(acceptAppointment))

	mux.Handle("/count", isAuthenticated(labModelCount))
	mux.Handle("/doctor_model_count", isAuthenticated(doctorModelCount))
	mux.Handle("/patient_model_count", isAuthenticated(patientModelCount))

	fmt.Println("Server Running on port :8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func dbMigration() {
	db.AutoMigrate(User{})
	db.AutoMigrate(Doctor{})
	db.AutoMigrate(Patient{})
	db.AutoMigrate(&PatientTest{})
	db.AutoMigrate(Test{})
	db.AutoMigrate(LabAdmin{})
	db.AutoMigrate(Schedule{})
	db.AutoMigrate(&Appointment{})
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	db.First(&user, "id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(user)
}

// Auth
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.First(&user, "email = ?", user.Email)
	if result.RowsAffected != 0 {
		json.NewEncoder(w).Encode("User With This Email Already Exists")
		return
	}
	hw_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return
	}
	u_result := db.Create(&User{
		ID:       uuid.NewString(),
		Email:    user.Email,
		Password: string(hw_password),
		Role:     user.Role,
		Approved: false,
	})
	if u_result.Error != nil {
		json.NewEncoder(w).Encode("Failed To Create User")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func login(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Email    string
		Password string
	}
	var body Body
	var user User

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	result := db.Where("email = ?", body.Email).First(&user)
	if result.RowsAffected == 0 {
		http.Error(w, "User Does Not Exist", http.StatusNotFound)
		return
	}
	n_err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if n_err != nil {
		// fmt.Println(n_err.Error())
		// http.Error(w, "Incorrect Password", http.StatusNotFound)
		return
	} else {
		fmt.Println("Hash Success")
	}
	// //Generate Jwt Token7
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		http.Error(w, "Failed to create token", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&tokenString)
}

func isAuthenticated(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Checking Auth Token")

		reqToken := r.Header.Get("Authorization")

		splitToken := strings.Split(reqToken, " ")
		if len(splitToken) != 2 {
			fmt.Printf("Error Reading Token")
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])

		if reqToken != "" {
			// fmt.Println("Header is set! We can serve content!")

			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if float64(time.Now().Unix()) > claims["exp"].(float64) {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}

				var user User
				db.First(&user, "id = ?", claims["sub"])

				if user.ID == "" {
					http.Error(w, err.Error(), http.StatusUnauthorized)
				}
				currentUser = LoggedUser{ID: user.ID,
					Email:    user.Email,
					Role:     user.Role,
					Approved: user.Approved}
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			endpoint(w, r)
		} else {
			fmt.Println("Not Authorized")
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
