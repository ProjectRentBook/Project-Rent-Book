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

func (p *Pinjam) PinjamanBuku() PinjamBuku {
	var BorrowBook PinjamBuku
	if err := p.DB.Create(&BorrowBook).Error; err != nil {
		log.Print(err)
		return PinjamBuku{}
	}
	return BorrowBook
}

func (p *Pinjam) KembalikanBuku() PinjamBuku {
	var ReturnBook PinjamBuku
	if err := p.DB.Create(&ReturnBook).Error; err != nil {
		log.Print(err)
		return PinjamBuku{}
	}
	return ReturnBook
}
