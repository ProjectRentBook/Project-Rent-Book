package main

import (
	config "Project_Rent_Book/Config"
	"Project_Rent_Book/entity"
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	conn := config.InitDB()
	SignUp := entity.SignUp{DB: conn}
	Buku := entity.Buku{DB: conn}
	Pinjam := entity.Pinjam{DB: conn}
	var input int = 0
	for input != 11 {
		fmt.Println("|________________________________________|")
		fmt.Println("|         WELCOME TO OUR LIBRARY         |")
		fmt.Println("|________________________________________|")
		fmt.Println("|                                        |")
		fmt.Println("|===============--MENU--=================|")
		fmt.Println("|========================================|")
		fmt.Println("|1. Sign Up                              |")
		fmt.Println("|2. Login                                |")
		fmt.Println("|3. List Buku                            |")
		fmt.Println("|11.Exit                                 |")
		fmt.Println("|________________________________________|")
		fmt.Print("Pilih menu :")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var UserBaru entity.User
			fmt.Print("Nama\t: ")
			in := bufio.NewReader(os.Stdin)
			UserBaru.Nama, _ = in.ReadString('\n')
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
			var ListUser entity.User
			fmt.Print("Email\t: ")
			fmt.Scanln(&ListUser.Email)
			fmt.Println(ListUser.Email)
			fmt.Print("Password: ")
			fmt.Scanln(&ListUser.Password)
			fmt.Println(ListUser.Password)
			res := SignUp.LoginUser(ListUser.Email, ListUser.Password)
			Email := res.Email
			fmt.Println(Email)
			fmt.Println(res.Email)

			// if res[0].ID < 1 {
			// 	fmt.Println("Tidak bisa Login")
			// 	break
			// }
			// fmt.Println("Berhasil Login")

			var input2 int
			for input2 != 8 {
				fmt.Println("                          ")
				fmt.Println("|________________________|")
				fmt.Println("|1. Update Profile       |")
				fmt.Println("|2. Delete Account       |")
				fmt.Println("|3. Mendaftarkan Buku    |")
				fmt.Println("|4. Update Buku          |")
				fmt.Println("|5. Hapus Buku           |")
				fmt.Println("|6. Pinjam Buku          |")
				fmt.Println("|7. Kembalikan Buku      |")
				fmt.Println("|8. Exit                 |")
				fmt.Println("|________________________|")
				fmt.Print("Pilih menu :")
				fmt.Scanln(&input2)
				switch input2 {
				case 1:
					// var Update entity.User
					// // fmt.Println("ID: ")
					// // fmt.Scanln(&Update.ID)
					// fmt.Print("Nama\t: ")
					// in := bufio.NewReader(os.Stdin)
					// Update.Nama, _ = in.ReadString('\n')
					// fmt.Print("Email: ")
					// fmt.Scanln(&Update.email)
					// fmt.Print("Password: ")
					// fmt.Scanln(&Update.Password)
					// fmt.Print("Usia: ")
					// fmt.Scanln(&Update.Umur)
					// res2 := SignUp.UpdateUser(Email, Update)
					// if res2.ID == 0 {
					// 	fmt.Println("Tidak ada yang diupdate")
					// 	break
					// }
					// fmt.Println("Berhasil Update Profil")
					// ----------------------------------------------
					// var Update entity.User
					// fmt.Println("ID: ")
					// fmt.Scanln(&Update.ID)
					var namaUpdate, emailUpdate, passwordUpdate string
					var umurUpdate string
					fmt.Print("Nama\t: ")
					in := bufio.NewReader(os.Stdin)
					namaUpdate, _ = in.ReadString('\n')
					fmt.Print("Email: ")
					fmt.Scanln(&emailUpdate)
					fmt.Print("Password: ")
					fmt.Scanln(&passwordUpdate)
					fmt.Print("Usia: ")
					fmt.Scanln(&umurUpdate)
					res2 := SignUp.UpdateUser(Email, namaUpdate, emailUpdate, passwordUpdate, umurUpdate)
					if res2 {
						fmt.Print("berhasil update profil")
						break
					}
					fmt.Println("Gagal Update Profil")

				case 2:
					var ID uint
					fmt.Print("Masukkan ID yang akan dihapus: ")
					fmt.Scanln(&ID)
					SignUp.DeleteUser(ID)

				case 3:
					var BukuBaru entity.Book
					fmt.Print("IDuser: ")
					fmt.Scanln(&BukuBaru.IDuser)
					fmt.Print("Judul\t: ")
					in := bufio.NewReader(os.Stdin)
					BukuBaru.Judul, _ = in.ReadString('\n')
					fmt.Print("Genre\t: ")
					BukuBaru.Genre, _ = in.ReadString('\n')
					fmt.Print("Penulis\t: ")
					BukuBaru.Penulis, _ = in.ReadString('\n')
					res := Buku.TambahBuku(BukuBaru)
					if res.ID == 0 {
						fmt.Println("tidak bisa input Buku")
						break
					}
					fmt.Println("berhasil input Buku")
				case 4:
					var ID_Buku entity.Book
					fmt.Print("ID: ")
					fmt.Scanln(&ID_Buku.ID)
					fmt.Print("Judul\t: ")
					in := bufio.NewReader(os.Stdin)
					ID_Buku.Judul, _ = in.ReadString('\n')
					fmt.Print("Genre\t: ")
					ID_Buku.Genre, _ = in.ReadString('\n')
					fmt.Print("Penulis\t: ")
					ID_Buku.Penulis, _ = in.ReadString('\n')
					res2 := Buku.UpdateBuku(int(ID_Buku.ID), ID_Buku)
					if res2.ID == 0 {
						fmt.Println("Tidak ada buku yang diupdate")
						break
					}
					fmt.Println("Berhasil Update Buku")
				case 5:
					var ID uint
					fmt.Print("Masukkan ID Buku yang akan Dihapus: ")
					fmt.Scanln(&ID)
					Buku.HapusBuku(ID)
				case 6:
					var BorrowBook entity.PinjamBuku
					fmt.Print("ID Buku: ")
					fmt.Scanln(&BorrowBook.IDBook)
					fmt.Print("ID User: ")
					fmt.Scanln(&BorrowBook.IDuser)
					fmt.Print("Tanggal Pinjam: ")
					fmt.Scanln(&BorrowBook.TanggalPinjam)
					v := Pinjam.PinjamanBuku(BorrowBook)
					if v.ID < 1 {
						fmt.Println("Gagal meminjam buku")
						break
					}
					fmt.Println("Berhasil meminjam buku")
				case 7:
					var IDBuku uint
					fmt.Print("Masukkan ID Buku yang akan Dikembalikan: ")
					fmt.Scanln(&IDBuku)
					Pinjam.KembalikanBuku(IDBuku)
					fmt.Println("Semoga Harimu Menyenangkan")
				}
			}
		case 3:
			fmt.Println("Daftar semua Buku")
			for _, v := range Buku.DaftarBuku() {
				fmt.Println(v)
			}
		default:
			continue
		}
	}
	fmt.Println("Thank you for using our program, see you next time")
}
