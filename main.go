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
	Pinjam := entity.Pinjam{conn}
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
			fmt.Print("Nama: ")
			fmt.Scanln(&UserBaru.Nama)
			fmt.Print("Email: ")
			fmt.Scanln(&UserBaru.Email)
			fmt.Print("Password: ")
			fmt.Scanln(&UserBaru.Password)
			fmt.Print("Usia : ")
			fmt.Scanln(&UserBaru.Umur)
			res := SignUp.RegisterUser(UserBaru)
			if res.ID == 0 {
				fmt.Println("tidak bisa input User")
				break
			}
			fmt.Println("berhasil input user")

		case 2:
			fmt.Println("Login User")
			for _, v := range SignUp.LoginUser() {
				fmt.Println(v)
			}
		case 3:
			var Update entity.User
			fmt.Print("Nama: ")
			fmt.Scanln(&Update.Nama)
			fmt.Print("Email: ")
			fmt.Scanln(&Update.Email)
			fmt.Print("Password: ")
			fmt.Scanln(&Update.Password)
			fmt.Print("Usia: ")
			fmt.Scanln(&Update.Umur)
			res := SignUp.UpdateUser()
			if res.ID == 0 {
				fmt.Println("Tidak ada yang diupdate")
				break
			}
			fmt.Println("Berhasil Update Profil")

		case 4:
			var ID uint
			fmt.Print("Masukkan ID yang akan dihapus: ")
			fmt.Scanln(&ID)
			SignUp.DeleteUser(ID)

		case 5:
			var BukuBaru entity.Book
			fmt.Print("IDuser: ")
			fmt.Scanln(&BukuBaru.IDuser)
			fmt.Print("Judul: ")
			fmt.Scanln(&BukuBaru.Judul)
			fmt.Print("Genre: ")
			fmt.Scanln(&BukuBaru.Genre)
			fmt.Print("Penulis: ")
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
			var ID_Buku entity.Book
			fmt.Print("IDuser: ")
			fmt.Scanln(&ID_Buku.IDuser)
			fmt.Print("Judul: ")
			fmt.Scanln(&ID_Buku.Judul)
			fmt.Print("Genre: ")
			fmt.Scanln(&ID_Buku.Genre)
			fmt.Print("Penulis: ")
			fmt.Scanln(&ID_Buku.Penulis)
			res := Buku.UpdateBuku()
			if res.ID == 1 {
				fmt.Println("Tidak ada buku yang di Update")
				break
			}
			fmt.Println("Berhasil Update Buku")
		case 8:
			var ID uint
			fmt.Print("Masukkan ID Buku yang akan Dihapus: ")
			fmt.Scanln(&ID)
			Buku.HapusBuku(ID)
		case 9:
			var BorrowBook entity.PinjamBuku
			fmt.Print("ID Buku: ")
			fmt.Scanln(&BorrowBook.IDBook)
			fmt.Print("ID User: ")
			fmt.Scanln(&BorrowBook.IDuser)
			fmt.Print("Tanggal Pinjam: ")
			fmt.Scanln(&BorrowBook.TanggalPinjam)
			v := Pinjam.PinjamanBuku()
			if v.ID < 1 {
				fmt.Println("Gagal meminjam buku")
				break
			}
			fmt.Println("Berhasil meminjam buku")
		case 10:
			var ReturnBook entity.PinjamBuku
			fmt.Print("ID Buku: ")
			fmt.Scanln(&ReturnBook.IDBook)
			fmt.Print("ID User: ")
			fmt.Scanln(&ReturnBook.IDuser)
			fmt.Print("Tanggal Kembali: ")
			fmt.Scanln(&ReturnBook.TanggalKembali)
			v := Pinjam.KembalikanBuku()
			if v.ID < 1 {
				fmt.Println("Gagal mengembalikan Buku")
				break
			}
			fmt.Println("Berhasil kembalikan Buku")
		default:
			continue
		}

	}
	fmt.Println("Thank you for using our program, see next time")
}