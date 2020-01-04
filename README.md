# Belajar CRUD di Golang

### Requirement 
1. [GO Dep](https://github.com/golang/dep)
2. Mysql database (jika menggunakan docker, sudah tersedia)

### Cara Install
Clone repo ini, ketik perintah `dep ensure` kemudian jalan kan dengan `go run main.go`. 
gocrud membutuhkan mysql sebagai database, file sql sudah tersedia di repo. 

Jika kamu menggunakan docker, cukup clone repo ini kemudian `docker-compose up` tapi docker disini hanya menjalankan 
container mysql saja, jadi untuk bisa unjalankan nya harus dengan perintah `go run main.go`

