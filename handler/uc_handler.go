package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"enigmacamp.com/be-lms-university/model"
	"github.com/gin-gonic/gin"
)

var UCDb []model.UserCredential

func CreateUCHandler(ctx *gin.Context) {
	// TODO:
	/**
	1. Kita siapkan sebuah payload (Struct UserCredential)
	2. Kemudian kita validasi menggunakan :
		- ctx.MustBindWith
		- ctx.ShouldBind
		- ctx.ShouldBindJson
		- dll..
	3. Setelah berhasil divalidasi, kita bisa sampan hasilnya, mekanismenya adalah:
		- Simpan di slice saja -> []UserCredential
	4. Kita kembalikan response hasil yang diinput/request body nya
	*/

	var payload model.UserCredential

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	UCDb = append(UCDb, payload)

	model.SendSingleResponse(ctx, "Successfully created user", payload)
}

func CreateUCWithPhotoHandler(ctx *gin.Context) {
	// TODO:
	/*
		1. Kita lakukan insert data sekaligus upload photo
		2. Body payloadnya akan berbentuk form-data
		3. Kemudian data email dan photo akan tersimpan di payload
		4. Photo hanya akan mengambil pathnya saja
	*/

	user := ctx.PostForm("user")
	file, header, err := ctx.Request.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	defer file.Close()

	// Kita siapkan sebuah path dimana photo itu disimpan
	// uploads/gambar.jpg -> save di database
	// GO => file => kita membuat dan menulis file
	// Kita juga bisa menggabungkan nama path + naa file nya
	// Nama file bisa di ubah ya
	// bisa menggunakan random string atau pattern => username_photo.jpg / random number_photo.jg
	// ambil ekstension nya

	newFileName := fmt.Sprintf("%v_photo%s", rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Int(), filepath.Ext(header.Filename))

	fileLocation := filepath.Join("uploads", newFileName)

	// Setelah itu kita harus buat foldernya untuk menyimpan gambar salinannya
	// Buat folder secara otomatis
	os.MkdirAll("uploads", os.ModePerm) // buat foldernya sekaligus kasih permission 0666

	// save file yg diupload
	if err := ctx.SaveUploadedFile(header, fileLocation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Kita perlu ekstrack dari key "user"
	// Ekstrak menggunakan json.Unmarshal
	// Ekstrak ke variable userCredential
	var userCredential model.UserCredential
	// Masukkan datanya ke dalam struct
	json.Unmarshal([]byte(user), &userCredential)

	// Simpan ke slice
	// Ambil nama filenya
	userCredential.Photo = fileLocation
	UCDb = append(UCDb, userCredential)

	model.SendSingleResponse(ctx, "Successfully created user", userCredential)
}
