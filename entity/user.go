package entity

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string
	Password string
	Umur     string
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

func (su *SignUp) LoginUser(Email, Password string) User {
	var ListUser = User{}
	if err := su.DB.Where("email = ? AND Password = ?", Email, Password).Find(&ListUser).Error; err != nil {
		log.Print(err)
		fmt.Println("tidak bisa login")
		return User{}
	}
	fmt.Println("berhasil login")
	return ListUser
}

// func (su *SignUp)  gettAlluser(email string) User {
// 	var Alluser = User{}
// 	if err := su.DB.Where()
// }

// func (su *SignUp) UpdateUser(Email string, newData User) User {
// 	// err := su.DB.Model(User{}).Where("email = ?", Email).Updates(newData)
// 	// if err.Error != nil {
// 	// 	log.Print(err)
// 	// 	return User{}
// 	// }
// 	// return newData
// 	}
func (su *SignUp) UpdateUser(Email string, namaUpdate string, emailUpdate string, passwordUpdate string, umurUpdate string) bool {
	updateExc := su.DB.Model(User{}).Where("email = ?", Email).Updates(User{Nama: namaUpdate, Email: emailUpdate, Password: passwordUpdate, Umur: umurUpdate})
	if err := updateExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := updateExc.RowsAffected; aff < 1 {
		return false
	}
	return true
}

func (su *SignUp) DeleteUser(ID_User uint) bool {
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
