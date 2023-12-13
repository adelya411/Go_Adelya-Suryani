package main

// Struktur user dengan tipe data yang tidak sesuai. Kurang deskriptif dan sulit dipahami
type user struct {
	id       int		//Penggunaan tipe data int untuk username dan password
	username int		//Penamaan variabel yang tidak deskriptif
	password int
}

// Struktur userservice dengan slice user yang tidak diinisialisasi
type userservice struct {		
	t []user   		//Penggunaan slice user yang tidak diinisialisasi
}

// Metode getallusers() yang mengembalikan seluruh data user
func (u userservice) getallusers() []user {
	return u.t
}

// Metode getuserbyid() yang mengembalikan user berdasarkan ID
func (u userservice) getuserbyid(id int) user {			//Metode getuserbyid() hanya mengembalikan user kosong jika tidak ditemukan ID
	// Iterasi melalui slice user untuk mencari user berdasarkan ID
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	// Mengembalikan user kosong jika tidak ditemukan user dengan ID yang sesuai
	return user{}		//Tidak ada kontrol akses
}


//Perbaikan code

package main

// Struct user dengan tipe data yang lebih sesuai
type user struct {
	ID       int
	Username string
	Password string
}

// Struct userservice dengan slice user yang sudah diinisialisasi
type userservice struct {
	users []user
}

// Metode GetAllUsers() yang mengembalikan seluruh data user
func (u *userservice) GetAllUsers() []user {
	return u.users
}

// Metode GetUserByID() yang mengembalikan user berdasarkan ID
func (u *userservice) GetUserByID(id int) user {
	// Iterasi melalui slice user untuk mencari user berdasarkan ID
	for _, r := range u.users {
		if id == r.ID {
			return r
		}
	}

	// Mengembalikan user kosong jika tidak ditemukan user dengan ID yang sesuai
	return user{}
}

