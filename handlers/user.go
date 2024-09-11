package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tes-rssa/database"
	"tes-rssa/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]



	var user models.User
	result := database.DB.Debug().Raw("SELECT * FROM user WHERE ID = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	jenisKelamin := "Laki-Laki"
	if user.JenisKelamin == "2" {
		jenisKelamin = "Perempuan"
	}

	user.JenisKelamin = jenisKelamin
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := database.DB.Debug().Raw("SELECT * FROM user").Find(&users)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Users not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	for i, v := range users {
		jenisKelamin := "Laki-Laki"
		if v.JenisKelamin == "2" {
			jenisKelamin = "Perempuan"
		}

		v.JenisKelamin = jenisKelamin
		users[i] = v
	}
	fmt.Println(users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.InsertUser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	insertData := &models.InsertUser{
		Nama:         u.Nama,
		Umur:         u.Umur,
		Alamat:       u.Alamat,
		Agama:        u.Agama,
		JenisKelamin: u.JenisKelamin,
	}
	queryInsert := "INSERT INTO user (nama, umur, alamat, agama, jenis_kelamin) VALUES (?, ?, ?, ?, ?)"
	result := database.DB.Debug().Exec(queryInsert, insertData.Nama, insertData.Umur, insertData.Alamat, insertData.Agama, insertData.JenisKelamin)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	res := &models.Response{
		ResponseCode:    "00",
		ResponseMessage: "Berhasil Insert Data",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var u models.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	update := &models.UpdateUser{
		Nama:         u.Nama,
		Umur:         u.Umur,
		Alamat:       u.Alamat,
		Agama:        u.Agama,
		JenisKelamin: u.JenisKelamin,
	}
	result := database.DB.Debug().Exec("UPDATE user SET nama = ?, umur = ?, alamat = ?, agama = ?, jenis_kelamin = ? WHERE ID = ?", update.Nama, update.Umur, update.Alamat, update.Agama, update.JenisKelamin, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	res := &models.Response{
		ResponseCode:    "00",
		ResponseMessage: "Berhasil Update Data",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result := database.DB.Debug().Exec("DELETE from user where id = ?", id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	res := &models.Response{
		ResponseCode:    "00",
		ResponseMessage: "Berhasil Update Data",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}