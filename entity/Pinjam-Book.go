package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type PinjamBuku struct {
	gorm.Model
	IDBook         int
	IDuser         int
	TanggalPinjam  time.Time
	TanggalKembali time.Time
}

type Pinjam struct {
	DB *gorm.DB
}

func (p *Pinjam) PinjamanBuku(BorrowBook PinjamBuku) PinjamBuku {
	//var BorrowBook PinjamBuku
	if err := p.DB.Create(&BorrowBook).Error; err != nil {
		log.Print(err)
		return PinjamBuku{}
	}
	return BorrowBook
}

func (p *Pinjam) KembalikanBuku(IDBuku uint) bool {
	postExc := p.DB.Where("id_book = ?", IDBuku).Delete(&PinjamBuku{})
	if err := postExc.Error; err != nil {
		log.Print(err)
		return false
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada buku yang dikembalikan")
		return false
	}
	return true
}
