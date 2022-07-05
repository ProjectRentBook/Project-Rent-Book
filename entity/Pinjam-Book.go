package entity

import (
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

//func (p *Pinjam) PinjamanBuku()
