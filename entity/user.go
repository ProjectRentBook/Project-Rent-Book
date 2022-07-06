package entity

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string
	Password string
	Umur     int
	//Books    []Book `gorm:"many2many:User_Books;"`
}

type SignUp struct {
	DB *gorm.DB
}

func (su *SignUp) RegisterUser(UserBaru User) User {
	err := su.DB.Create(&UserBaru).Error
	if err != nil {
		log.Fatal(err)
		return User{}
	}
	return UserBaru
}

func (su *SignUp) LoginUser(Nama, Password string) []User {
	var ListUser []User

	if err := su.DB.Where("Nama = ? AND Password = ?", Nama, Password).Find(&ListUser).Error; err != nil {
		log.Fatal(err)
		return nil
	}
	return ListUser
}

func (su *SignUp) UpdateUser(ID_User uint) User {
	//var Update User
	err := su.DB.Model(&User{}).Where("ID = ?", ID_User).Updates(User{})
	if err != nil {
		log.Print(err)
		return User{}
	}
	return User{}
}

func (su *SignUp) DeleteUser(ID_User uint) bool {
	//var HapusUser user
	postExc := su.DB.Where("ID = ?", ID_User).Delete(&User{})
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
