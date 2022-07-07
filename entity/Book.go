package entity

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	IDuser  int
	Judul   string
	Genre   string
	Penulis string
}

type Buku struct {
	DB *gorm.DB
}

func (b *Buku) TambahBuku(BukuBaru Book) Book {
	err := b.DB.Create(&BukuBaru).Error
	if err != nil {
		log.Print(err)
		return Book{}
	}
	return BukuBaru
}

func (b *Buku) DaftarBuku() []Book {
	var ListBuku = []Book{}

	if err := b.DB.Find(&ListBuku).Error; err != nil {
		log.Print(err)
		return nil
	}
	return ListBuku
}

func (b *Buku) UpdateBuku(IDuser int, newBuku Book) Book {
	err := b.DB.Model(Book{}).Where("IDuser =? ", IDuser).Updates(newBuku)
	if err.Error != nil {
		log.Print(err)
		return Book{}
	}
	return newBuku
}

// func (b *Buku) UpdateBuku(ID_Buku uint) bool {
// 	Exc := b.DB.Where("ID = ?", ID_Buku).Updates(&Book{})
// 	if err := Exc.Error; err != nil {
// 		log.Print(err)
// 		return false
// 	}
// 	// if aff := Exc.RowsAffected; aff < 1 {
// 	// 	log.Println("Tidak ada data yang berubah")
// 	// 	return false
// 	// }
// 	return true
// }

func (b *Buku) HapusBuku(ID_Buku uint) bool {
	postExc := b.DB.Where("ID = ?", ID_Buku).Delete(&Book{})
	if err := postExc.Error; err != nil {
		log.Print(err)
		return false
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}
	return true
}
