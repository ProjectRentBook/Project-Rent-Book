package main

import (
	"fmt"

	config "Project_Rent_Book/Config"
	"Project_Rent_Book/entity"
)

func main() {
	conn := config.InitDB()
	SignUp := entity.SignUp{conn}
	Buku := entity.Buku{conn}
	var input int = 0
	for input != 11 {
		fmt.Println("--Menu--")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Update Profile")
		fmt.Println("4. Delete Account")
		fmt.Println("5. Mendaftarkan Buku")
		fmt.Println("6. List Buku")
		fmt.Println("7. Update Buku")
		fmt.Println("8. Hapus Buku")
		fmt.Println("9. Pinjam Buku")
		fmt.Println("10. Kembalikan Buku")
		fmt.Println("11. Exit")
		fmt.Print("Pilih menu :")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var UserBaru entity.User
			fmt.Print("Nama:")
			fmt.Scanln(&UserBaru.Nama)
			fmt.Print("Email:")
			fmt.Scanln(&UserBaru.Email)
			fmt.Print("Password:")
			fmt.Scanln(&UserBaru.Password)
			fmt.Print("Usia :")
			fmt.Scanln(&UserBaru.Umur)
			res := SignUp.RegisterUser(UserBaru)
			if res.ID == 0 {
				fmt.Println("tidak bisa input User")
				break
			}
			fmt.Println("berhasil input user")

		// case 2:
		// 	fmt.Println("Daftar semua User")
		// 	for _, v := range SignUp.LoginUser() {
		// 		fmt.Println(v)
		// 	}
		case 3:

		case 4:
			var ID uint
			fmt.Print("Masukkan ID yang akan dihapus: ")
			fmt.Scanln(&ID)
			SignUp.DeleteUser(ID)

		case 5:
			var BukuBaru entity.Book
			fmt.Print("IDuser")
			fmt.Scanln(&BukuBaru.IDuser)
			fmt.Print("Judul")
			fmt.Scanln(&BukuBaru.Judul)
			fmt.Print("Genre")
			fmt.Scanln(&BukuBaru.Genre)
			fmt.Print("Penulis")
			fmt.Scanln(&BukuBaru.Penulis)
			res := Buku.TambahBuku(BukuBaru)
			if res.ID == 0 {
				fmt.Println("tidak bisa input Buku")
				break
			}
			fmt.Println("berhasil input Buku")

		case 6:
			fmt.Println("Daftar semua Buku")
			for _, v := range Buku.DaftarBuku() {
				fmt.Println(v)
			}
		case 7:
		default:
			continue
		}

	}
	fmt.Println("Thank you for using our program, see next time")
}
