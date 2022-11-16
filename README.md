# go-simple-web

Membuat web CRUD menggunakan Golang

## Dokumentasi

Untuk url dan pathnya sebagai berikut:

| **Path**       | **Methods** | **Description**                           |
|----------------|-------------|-------------------------------------------|
| /              | GET         | Menampilkan seluruh data pada database.   |
| /add           | POST        | Membuat data baru ke dalam database.      |
| /edit?id=123   | GET \| POST | Mengedit data data lama dengan yang baru. |
| /delete?id=123 | GET         | Menghapus data                            |