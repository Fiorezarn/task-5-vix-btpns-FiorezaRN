
# Deskripsi

Ketentuan API :
Pada bagian User Endpoint :

POST : /users/register, dan gunakan atribut berikut ini
ID (primary key, required)
Username (required)
Email (unique & required) 
Password (required & minlength 6)
Relasi dengan model Photo (Gunakan constraint cascade)
Created At (timestamp)
Updated At (timestamp)
GET: /users/login
Using email & password (required)
PUT : /users/:userId (Update User)
DELETE : /users/:userId (Delete User)
Photos Endpoint

POST : /photos 
ID
Title
Caption
PhotoUrl
UserID
Relasi dengan model User
GET : /photos
PUT : /photoId
DELETE : /:photoId


Requirement :
Authorization dapat menggunakan tool Go JWT (https://github.com/dgrijalva/jwt-go) 
Pastikan hanya user yang membuat foto yang dapat menghapus / mengubah foto


Environment
Struktur dokumen / environment dari GoLang yang akan dibentuk kurang lebih sebagai berikut :

app :Menampung pembuatan struct dalam kasus ini menggunakan struct user untuk keperluan data dan authentication
controllers : Berisi antara logic database yaitu models dan query
database: Berisi konfigurasi database serta digunakan untuk menjalankan koneksi database dan migration
helpers : Berisi fungsi-fungsi yang dapat digunakan di setiap tempat dalam hal ini jwt, bcrypt, headerValue
middlewares :Berisi fungsi yang digunakan untuk proses otentikasi jwt yang digunakan untuk proteksi api
models : Berisi models yang digunakan untuk relasi database 
router : Berisi konfigurasi routing / endpoint yang akan digunakan untuk mengakses api
go mod : Yang digunakan untuk manajemen package / dependency berupa library


Tools :
Tools yang dapat kalian gunakan : 

Gin Gonic Framework : https://github.com/gin-gonic/gin 
Gorm : https://gorm.io/index.html 
JWT Go : https://github.com/dgrijalva/jwt-go 
Go Validator : http://github.com/asaskevich/govalidator 


Untuk database, gunakanlah server SQL open source seperti MySQL, PostgreSQL, atau Microsoft SQL Server.
