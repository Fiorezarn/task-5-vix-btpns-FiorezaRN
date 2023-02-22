
# Deskripsi

Ketentuan API :
Pada bagian User Endpoint :

<ol><li>POST : /users/register, dan gunakan atribut berikut ini</li>
<li>ID (primary key, required)</li>
<li>Username (required)</li>
<li>Email (unique & required) </li>
<li>Password (required & minlength 6)</li>
<li>Relasi dengan model Photo (Gunakan constraint cascade)</li>
<li>Created At (timestamp)</li>
<li>Updated At (timestamp)</li>
<li>GET: /users/login</li>
<li>Using email & password (required)</li>
<li>PUT : /users/:userId (Update User)</li>
<li>DELETE : /users/:userId (Delete User)</li>
<li>Photos Endpoint</li></ol>

POST : /photos 
<ol><li>ID</li>
<li>Title</li>
<li>Caption</li>
<li>PhotoUrl</li>
<li>UserID</li>
<li>Relasi dengan model User</li>
<li>GET : /photos</li>
<li>PUT : /photoId</li>
<li>DELETE : /:photoId</li></ol>


Requirement :
Authorization dapat menggunakan tool Go JWT (https://github.com/dgrijalva/jwt-go) 
Pastikan hanya user yang membuat foto yang dapat menghapus / mengubah foto


Environment
Struktur dokumen / environment dari GoLang yang akan dibentuk kurang lebih sebagai berikut :

<ol><li>app :Menampung pembuatan struct dalam kasus ini menggunakan struct user untuk keperluan data dan authentication</li>
<li>controllers : Berisi antara logic database yaitu models dan query</li>
<li>database: Berisi konfigurasi database serta digunakan untuk menjalankan koneksi database dan migration</li>
<li>helpers : Berisi fungsi-fungsi yang dapat digunakan di setiap tempat dalam hal ini jwt, bcrypt, headerValue</li>
<li>middlewares :Berisi fungsi yang digunakan untuk proses otentikasi jwt yang digunakan untuk proteksi api</li>
<li>models : Berisi models yang digunakan untuk relasi database </li>
<li>router : Berisi konfigurasi routing / endpoint yang akan digunakan untuk mengakses api</li>
<li>go mod : Yang digunakan untuk manajemen package / dependency berupa library</li></ol>


Tools :
Tools yang dapat kalian gunakan : 

<ol><li>Gin Gonic Framework : https://github.com/gin-gonic/gin </li>
<li>Gorm : https://gorm.io/index.html </li>
<li>JWT Go : https://github.com/dgrijalva/jwt-go </li>
<li>Go Validator : http://github.com/asaskevich/govalidator</li> </ol>


Untuk database, gunakanlah server SQL open source seperti MySQL, PostgreSQL, atau Microsoft SQL Server.
