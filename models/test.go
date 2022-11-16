package models

// Ini akan dirender ke dalam database
// Misalkan `Data`. nah ini dibuat sebuah table dengan nama `data`
// Untuk mengimplementasikan models lihat dokumentasi dari https://gorm.io
type Test struct {
	ID       uint   `gorm:"primaryKey" json:"id"` // id tipe data PRIMARY KEY
	Age      int    `gorm:"size:5"`               // tipe data INT(2)
	FullName string `gorm:"size:150"`             // full_name tipe data VARCHAR(150)
}

// CreateNewData
// adalah sebuah fungsi untuk membuat atau INSERT data ke dalam database `Data`
func CreateNewData(test *Test) error {
	return db.Create(test).Error
}

// GetDataByID
// mendapatkan data dengan memberikan `id`
func GetDataByID(id uint) (Test, error) {
	var data Test
	err := db.Model(&Test{}).Where("id = ?", id).First(&data).Error
	return data, err
}

// GetAllData
// mendapatkan seluruh data pada database
func GetAllData() []Test {
	var datas []Test
	db.Model(&Test{}).Find(&datas)
	return datas
}
