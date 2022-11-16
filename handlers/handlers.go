package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/EtheriousKelv/go-simple-web/models"
)

var flasher Flasher

// Get
// akan menghandler halaman awal dan menampilkan seluruh isi data pada Database
func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Flasher defer
		defer flasher.Del() // otomatis akan dihapus ketika window di reload

		// Render html file
		tmpl := template.Must(template.ParseFiles("views/index.tmpl"))

		// Mendapatkan seluruh data data dari models
		datas := models.GetAllData()

		// Membuat map variabel
		// Nanti akan menjadi variabel untuk html
		// Dan penggunaannya {{ .datas }}
		data := map[string]interface{}{
			"flasher": flasher,
			"datas":   datas,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		return
	}
}

// Post
// method untuk menambahkan data baru
func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		defer flasher.Del()
		// Render html
		tmpl := template.Must(template.ParseFiles("views/tambah.tmpl"))

		data := map[string]interface{}{
			"flasher": flasher,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		return
	}
	// Jika method adalah POST
	if r.Method == "POST" {
		// Mendapatkan value dar method post
		// dengan nama `full_name` dan `age`
		age := r.PostFormValue("age")
		ageInt, _ := strconv.Atoi(age) // Render string ke int
		fullName := r.PostFormValue("full_name")

		// Deklarasikan model data
		data := models.Test{
			Age:      ageInt,
			FullName: fullName,
		}

		// Simpan ke db
		err := models.CreateNewData(&data)
		if err != nil {
			flasher.Set("error", "Error: ada kegagalan ketika menambahkan data baru.")
			http.Redirect(w, r, "/add", http.StatusFound)
			return
		}

		// Jika tidak ada error
		flasher.Set("success", "Berhasil menambahkan data baru ke database.")
		http.Redirect(w, r, "/add", http.StatusFound)

		return
	}
}

// Handler untuk edit data
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		defer flasher.Del()

		// Mencari nilai dari query `id`
		idString := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idString)

		dataDB, err := models.GetDataByID(uint(id))
		fmt.Println(dataDB)
		if err != nil {
			http.Error(w, fmt.Sprintf("Data dengan ID : `id = %s` tidak dapat ditemukan.", idString), http.StatusNotFound)
			return
		}

		var data = map[string]interface{}{
			"flahser": flasher,
			"data":    dataDB,
		}

		tmpl := template.Must(template.ParseFiles("views/edit.tmpl"))
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// If method is POST
	// adalah method untuk memperbaharui data
	if r.Method == "POST" {
		// Mendapatkan query dengan key id
		idStr := r.URL.Query().Get("id")
		idInt, _ := strconv.Atoi(idStr)

		data, err := models.GetDataByID(uint(idInt))
		if err != nil {
			http.Error(w, fmt.Sprintf("Data dengan ID : `id = %s` tidak dapat ditemukan.", idStr), http.StatusNotFound)
			return
		}

		age := r.PostFormValue("age")
		ageInt, _ := strconv.Atoi(age)

		fullName := r.PostFormValue("full_name")

		// Declare model
		data.Age = ageInt
		data.FullName = fullName

		// Save
		err = models.GetDB().Save(&data).Error
		if err != nil {
			flasher.Set("error", "Gagal ketika mengupdate data.")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		flasher.Set("success", "Berhasil mengupdate data.")
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}
}

// Delete
// handler untuk menghapus data
func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(idStr)

	data, err := models.GetDataByID(uint(idInt))
	if err != nil {
		http.Error(w, fmt.Sprintf("Data dengan ID : `id = %s` tidak dapat ditemukan.", idStr), http.StatusNotFound)
		return
	}

	err = models.GetDB().Delete(&data).Error
	if err != nil {
		flasher.Set("success", "Error: tidak bisa menghapus data dengan id: "+idStr)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	flasher.Set("success", fmt.Sprintf("Menghapus data dengan id: %s berhasil dilakukan.", idStr))
	http.Redirect(w, r, "/", http.StatusFound)
}
