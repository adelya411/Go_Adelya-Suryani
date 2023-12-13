// Kode acak

package main

type kendaraan struct {
	totalroda       int
	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()
}

//Perbaikan
package main

// Struct kendaraan dengan informasi total roda dan kecepatan per jam
type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

// Struct mobil yang mempunyai kendaraan sebagai bagian dari komposisi
type mobil struct {
	kendaraan // Embedding kendaraan pada struct mobil
}

// Metode berjalan untuk mobil yang menambah kecepatan sebesar nilai tertentu
func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

// Metode tambahKecepatan untuk menambah kecepatan mobil
func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam += kecepatanBaru
}

func main() {
	// Membuat mobilcepat sebagai instance mobil
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	// Membuat mobillamban sebagai instance mobil lainnya
	mobilLamban := mobil{}
	mobilLamban.berjalan()
}
