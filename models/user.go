package models

type User struct {
	ID           string `json:"id"`
	Nama         string `json:"nama"`
	Umur         string `json:"umur"`
	Alamat       string `json:"alamat"`
	Agama        string `json:"agama"`
	JenisKelamin string `json:"jenis_kelamin"`
}

type UpdateUser struct {
	ID           string `json:"id"`
	Nama         string `json:"nama"`
	Umur         string `json:"umur"`
	Alamat       string `json:"alamat"`
	Agama        string `json:"agama"`
	JenisKelamin string `json:"jenis_kelamin"`
}

type InsertUser struct {
	Nama         string `json:"nama"`
	Umur         string `json:"umur"`
	Alamat       string `json:"alamat"`
	Agama        string `json:"agama"`
	JenisKelamin string `json:"jenis_kelamin"`
}

type Response struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
