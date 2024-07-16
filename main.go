package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	StatusUser string `json:"status_user"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type InsertUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	StatusUser string `json:"status_user"`
	TipeUser   string `json:"tipe_user"`
}

type Response struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

var db *gorm.DB
var err error

func main() {
	dsn := "root:P@ssw0rd@tcp(127.0.0.1:3306)/tes_rssa?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/user", getAllUser).Methods("GET")
	r.HandleFunc("/create", createUser).Methods("POST")
	r.HandleFunc("/update", updateUser).Methods("POST")

	log.Println("Server berjalan di port 8002")
	log.Fatal(http.ListenAndServe(":8002", r))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User
	result := db.Debug().Raw("SELECT * FROM user WHERE ID = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	var user []User
	result := db.Debug().Raw("SELECT id, username, password, status_user, fk_tipeuser_id FROM user").Find(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u InsertUser
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	insertData := &InsertUser{
		Username:   u.Username,
		Password:   u.Password,
		StatusUser: u.StatusUser,
		TipeUser:   u.TipeUser,
	}
	queryInsert := "insert into user " +
		"(username, password, status_user, fk_tipeuser_id) " +
		"values (?, ?, ?, ?)"
	result := db.Debug().Exec(queryInsert,
		insertData.Username,
		insertData.Password,
		insertData.StatusUser,
		insertData.TipeUser,
	)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	res := &Response{
		ResponseCode:    "00",
		ResponseMessage: "Berhasil Insert Data",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var u UpdateUser
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	update := &UpdateUser{
		ID:       u.ID,
		Password: u.Password,
	}
	result := db.Debug().Exec("UPDATE user SET PASSWORD = ? WHERE ID = ?", update.Password, update.ID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	res := &Response{
		ResponseCode:    "00",
		ResponseMessage: "Berhasil Update Data",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
