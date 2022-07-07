# HelloReaders-Peminjaman Buku
Aplikasi sederhana untuk peminjaman Buku dengan Go, GORM, dan MySQL.

## Daftar Tabel
- User (akun pengguna)
- Book (Daftar Buku)
- Pinjam Buku (untuk melihat daftar buku yang dipinjam)

Dengan menggunakan Program ini, para pengguna bisa menggunakannya untuk mendaftarkan buku untuk dipinjamkan maupun meminjam buku yang sudah ada. 
Tantangan yang dihadapi adalah program ini masih sangat sederhana sehingga masih banyak yang harus diperbaiki, dan berharap kedepannya untuk membuat fitur program ini semakin lengkap dan mudah untuk digunakan para pengguna.
### HOW TO INSTALL and RUN THE PROGRAM
pertama : install gorm terlebih dahulu di VSCode [](go get -U gorm.io/gorm)
kedua : menginstall [](go get -u gorm.io/driver/mysql) untuk menghubungkan go dengan database.
lalu membuat kode menggunakan fitur gorm diantaranya : gorm.Model <untuk membuat daftar ID secara otomatis>, lalu DB.Where <untuk menentukan kolom atau parameter yang dituju>, .Create<untuk membuat isi dari tabel>, .Find<untuk mengquery atau menampilkan data yang ingin ditampilkan>, .Updates<untuk mengupdate data>, .Delete<untuk menghapus data>.
 #### membuat struktur
  type <nama tabel> struct{
  [] daftar kolom dan tipe data}
  type <nama tabel> struct{
  gorm.Model}
##### create data
func <nama struct> <nama func>(parameter tipe_data) output{
  DB.Create(parameter) }
#### Login user
func <nama struct> <nama func>(parameter tipe_data) output{
  var <nama variable> tipe_data
  DB.Where(parameter).first(&nama variable) }
#### Update user
  func <nama struct> <nama func>(parameter tipe_data) output{
  DB.Model(nama struct).Where(parameter yang ditunjuk).Updates(parameter untuk menampung data baru)}
#### Delete user
  func <nama struc> <nama func>(parameter tipe_data) output{
  DB.Where(parameter).Delete(&nama struct {})}
  
catatan : untuk melihat daftar data bisa menggunakan Login user tanpa harus menggunakan parameter dan outputnya menggunakan [] array.

### Features
1. Menambahkan Akun Pengguna
2. Login Pengguna
3. Update Profile Pengguna
4. Hapus Akun Pengguna
5. Menambahkan Buku
6. Mengupdate buku
7. Menghapus buku
8. Menginput data peminjaman dan pengembalian buku

### Team Members
1. Jerry Young as Mentor (JerryBE1709)
2. Ghalib Assaidy as team (Ghalib12)
3. cindy shintia as team (cindy05-shintia)
  
  ### Contributing
  Pull requests are Welcome. For changes, Please open an issue first to discuss what would you like to change. Happy Coding
