package main

import "fmt"

type data struct {
	nim, jumlahBayar, kurangBayar, lebihBayar int
	nama, status                              string
}

type tabData [999]data

func main() {
	var jumMurid, bulan int
	var kas tabData

	fmt.Println("============================")
	fmt.Println("   APLIKASI CEK UANG KAS   ")
	fmt.Println("============================")
	menu(&jumMurid, &bulan, &kas)
}

// ============================================================
// HELPER INPUT
// ============================================================

func bacaInt(prompt string) int {
	var hasil int
	var err error
	var n int

	for {
		fmt.Print(prompt)
		n, err = fmt.Scan(&hasil)

		if err == nil && n == 1 {
			return hasil
		}

		fmt.Println("  [!] Input tidak valid. Masukkan angka.")
		fmt.Scanln()
	}
}

func bacaIntPositif(prompt string) int {
	for {
		angka := bacaInt(prompt)
		if angka > 0 {
			return angka
		}
		fmt.Println("  [!] Harus lebih dari 0.")
	}
}

func bacaIntMinNol(prompt string) int {
	for {
		angka := bacaInt(prompt)
		if angka >= 0 {
			return angka
		}
		fmt.Println("  [!] Tidak boleh negatif.")
	}
}

func bacaString(prompt string) string {
	var hasil string
	var err error
	var n int

	for {
		fmt.Print(prompt)
		n, err = fmt.Scan(&hasil)

		if err == nil && n == 1 && hasil != "" {
			return hasil
		}

		fmt.Println("  [!] Input tidak boleh kosong.")
		fmt.Scanln()
	}
}

func bacaKonfirmasi(prompt string) bool {
	for {
		var input string
		fmt.Print(prompt)
		fmt.Scan(&input)

		switch input {
		case "y", "Y":
			return true
		case "n", "N":
			return false
		default:
			fmt.Println("  [!] Masukkan 'y' atau 'n'.")
		}
	}
}

// ============================================================
// HELPER DATA
// ============================================================

func cariIndexNIM(kas *tabData, jumMurid int, target int) int {
	for i := 0; i < jumMurid; i++ {
		if kas[i].nim == target {
			return i
		}
	}
	return -1
}

func bacaDataMurid(kas *tabData, idx int) {
	var nim, jumlahBayar int
	var nama string
	var err error
	var n int

	for {
		n, err = fmt.Scan(&nim, &nama, &jumlahBayar)

		if err == nil && n == 3 {
			kas[idx].nim = nim
			kas[idx].nama = nama
			kas[idx].jumlahBayar = jumlahBayar
			return
		}

		fmt.Println("  [!] Format salah. Masukkan: NIM NAMA JUMLAH_BAYAR")
		fmt.Scanln()
		fmt.Printf("Data ke-%d : ", idx+1)
	}
}

func hitungStatus(kasPerBulan int, jumMurid *int, bulan *int, kas *tabData) {
	totalBayarWajib := *bulan * kasPerBulan

	for i := 0; i < *jumMurid; i++ {
		selisih := kas[i].jumlahBayar - totalBayarWajib

		if selisih >= 0 {
			kas[i].status = "LUNAS"
			kas[i].kurangBayar = 0
			kas[i].lebihBayar = selisih
		} else {
			kas[i].status = "BELUM LUNAS"
			kas[i].kurangBayar = -selisih
			kas[i].lebihBayar = 0
		}
	}
}

// ============================================================
// MENU UTAMA
// ============================================================

func menu(jumMurid *int, bulan *int, kas *tabData) {
	var kasPerBulan int

	for {
		fmt.Println()
		fmt.Println("========== MENU UTAMA ==========")
		fmt.Println("1. Input Uang Kas")
		fmt.Println("2. Tampilkan Semua")
		fmt.Println("3. Cek Status by NIM")
		fmt.Println("4. Cek Detail by NIM")
		fmt.Println("5. Urutkan Tampilan")
		fmt.Println("6. Edit Data Kas")
		fmt.Println("0. Exit")
		fmt.Println("================================")

		pilihan := bacaInt("Pilihan : ")

		if pilihan >= 2 && pilihan <= 6 && *jumMurid == 0 {
			fmt.Println("\n[!] Data masih kosong. Silakan input data terlebih dahulu.")
			pilihan = 1
		}

		switch pilihan {
		case 1:
			inputUangKas(&kasPerBulan, jumMurid, bulan, kas)
		case 2:
			tampilkanSemua(jumMurid, bulan, kas, &kasPerBulan)
		case 3:
			cekLunas(jumMurid, kas)
		case 4:
			cekNim(jumMurid, kas)
		case 5:
			urutkanTampilan(jumMurid, bulan, kas)
		case 6:
			editData(&kasPerBulan, jumMurid, bulan, kas)
		case 0:
			fmt.Println("Terima kasih. Program selesai.")
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid. Pilih 0-6.")
		}
	}
}

// ============================================================
// INPUT & HITUNG
// ============================================================

func inputUangKas(kasPerBulan, jumMurid *int, bulan *int, kas *tabData) {
	fmt.Println()
	fmt.Println("======== INPUT UANG KAS ========")

	*jumMurid = bacaIntPositif("Jumlah Murid : ")
	*bulan = bacaIntPositif("Jumlah Bulan : ")
	*kasPerBulan = bacaIntPositif("Kas per Bulan (Rp) : ")

	fmt.Println()
	fmt.Println("Masukkan data (NIM - NAMA - JUMLAH_BAYAR) :")
	fmt.Println("Contoh: 123456 BudiSantoso 150000")

	for i := 0; i < *jumMurid; i++ {
		fmt.Printf("Data ke-%d : ", i+1)
		bacaDataMurid(kas, i)
	}

	hitungStatus(*kasPerBulan, jumMurid, bulan, kas)
	fmt.Println()
	fmt.Println("Data berhasil diinput!")
}

// ============================================================
// TAMPILKAN
// ============================================================

func tampilkanSemua(jumMurid *int, bulan *int, kas *tabData, kasPerBulan *int) {
	fmt.Println()
	fmt.Println("=========== TAMPILKAN SEMUA ===========")

	sep := "+------+----------+---------------+-----------------+---------------+--------------+--------------+"
	fmt.Println(sep)
	fmt.Printf("| %-4s | %-8s | %-13s | %-15s | %-13s | %-12s | %-12s |\n",
		"NO", "NIM", "NAMA", "JUMLAH BAYAR", "STATUS", "KURANG BAYAR", "LEBIH BAYAR")
	fmt.Println(sep)

	for i := 0; i < *jumMurid; i++ {
		fmt.Printf("| %-4d | %-8d | %-13s | %-15d | %-13s | %-12d | %-12d |\n",
			i+1,
			kas[i].nim,
			kas[i].nama,
			kas[i].jumlahBayar,
			kas[i].status,
			kas[i].kurangBayar,
			kas[i].lebihBayar,
		)
	}

	fmt.Println(sep)
	fmt.Println()
	ringkasanData(jumMurid, kas)
}

func ringkasanData(jumMurid *int, kas *tabData) {
	var totalBayar, totalLunas, totalBelumLunas, totalLebihBayar, totalKurangBayar int

	for i := 0; i < *jumMurid; i++ {
		totalBayar += kas[i].jumlahBayar

		if kas[i].status == "LUNAS" {
			totalLunas++
			totalLebihBayar += kas[i].lebihBayar
		} else {
			totalBelumLunas++
			totalKurangBayar += kas[i].kurangBayar
		}
	}

	fmt.Println("------- RINGKASAN -------")
	fmt.Printf("Total Murid        : %d\n", *jumMurid)
	fmt.Printf("Lunas              : %d murid\n", totalLunas)
	fmt.Printf("Belum Lunas        : %d murid\n", totalBelumLunas)
	fmt.Printf("Total Bayar        : Rp %d\n", totalBayar)
	fmt.Printf("Total Lebih Bayar  : Rp %d\n", totalLebihBayar)
	fmt.Printf("Total Kurang Bayar : Rp %d\n", totalKurangBayar)
	fmt.Println("-------------------------")
}

// ============================================================
// CEK
// ============================================================

func cekLunas(jumMurid *int, kas *tabData) {
	for {
		fmt.Println()
		fmt.Println("======== CEK STATUS BY NIM ========")

		target := bacaIntMinNol("NIM Dicari (0 untuk kembali) : ")
		if target == 0 {
			return
		}

		idx := cariIndexNIM(kas, *jumMurid, target)

		if idx == -1 {
			fmt.Println("  [!] NIM tidak ditemukan.")
		} else {
			sep := "+----------+---------------+---------------+"
			fmt.Println(sep)
			fmt.Printf("| %-8s | %-13s | %-13s |\n", "NIM", "NAMA", "STATUS")
			fmt.Println(sep)
			fmt.Printf("| %-8d | %-13s | %-13s |\n",
				kas[idx].nim,
				kas[idx].nama,
				kas[idx].status)
			fmt.Println(sep)
		}
	}
}

func cekNim(jumMurid *int, kas *tabData) {
	for {
		fmt.Println()
		fmt.Println("======== CEK DETAIL BY NIM ========")

		target := bacaIntMinNol("NIM Dicari (0 untuk kembali) : ")
		if target == 0 {
			return
		}

		left := 0
		right := *jumMurid - 1
		foundIdx := -1

		for left <= right && foundIdx == -1 {
			mid := left + (right-left)/2

			if kas[mid].nim == target {
				foundIdx = mid
			} else if kas[mid].nim < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if foundIdx == -1 {
			fmt.Println("  [!] NIM tidak ditemukan.")
		} else {
			sep := "+----------+---------------+-----------------+---------------+--------------+--------------+"
			fmt.Println(sep)
			fmt.Printf("| %-8s | %-13s | %-15s | %-13s | %-12s | %-12s |\n",
				"NIM", "NAMA", "JUMLAH BAYAR", "STATUS", "KURANG BAYAR", "LEBIH BAYAR")
			fmt.Println(sep)
			fmt.Printf("| %-8d | %-13s | %-15d | %-13s | %-12d | %-12d |\n",
				kas[foundIdx].nim,
				kas[foundIdx].nama,
				kas[foundIdx].jumlahBayar,
				kas[foundIdx].status,
				kas[foundIdx].kurangBayar,
				kas[foundIdx].lebihBayar,
			)
			fmt.Println(sep)
		}
	}
}

// ============================================================
// URUTKAN
// ============================================================

func urutkanTampilan(jumMurid *int, bulan *int, kas *tabData) {
	fmt.Println()
	fmt.Println("======== URUTKAN TAMPILAN ========")
	fmt.Println("1. Urutkan by NIM (ascending)")
	fmt.Println("2. Belum Lunas dahulu")
	fmt.Println("0. Batal")

	pilihan := bacaInt("Pilihan : ")

	switch pilihan {
	case 1:
		for i := 1; i < *jumMurid; i++ {
			key := kas[i]
			j := i - 1

			for j >= 0 && kas[j].nim > key.nim {
				kas[j+1] = kas[j]
				j--
			}
			kas[j+1] = key
		}

		fmt.Println("Data berhasil diurutkan by NIM.")
		var kasPerBulan int
		tampilkanSemua(jumMurid, bulan, kas, &kasPerBulan)

	case 2:
		for i := 0; i < *jumMurid-1; i++ {
			idxMin := i

			for j := i + 1; j < *jumMurid; j++ {
				if kas[idxMin].status == "LUNAS" &&
					kas[j].status == "BELUM LUNAS" {
					idxMin = j
				}
			}

			temp := kas[i]
			kas[i] = kas[idxMin]
			kas[idxMin] = temp
		}

		fmt.Println("Data berhasil diurutkan: Belum Lunas dahulu.")
		var kasPerBulan int
		tampilkanSemua(jumMurid, bulan, kas, &kasPerBulan)

	case 0:
		return
	default:
		fmt.Println("  [!] Pilihan tidak valid.")
	}
}

// ============================================================
// EDIT DATA
// ============================================================

func editData(kasPerBulan *int, jumMurid *int, bulan *int, kas *tabData) {
	fmt.Println()
	fmt.Println("======== EDIT DATA KAS ========")
	fmt.Println("1. Reset Semua Data")
	fmt.Println("2. Tambah Data Murid")
	fmt.Println("3. Bayar Uang Kas")
	fmt.Println("4. Hapus Data Murid")
	fmt.Println("0. Batal")

	pilihan := bacaInt("Pilihan : ")

	switch pilihan {
	case 1:
		editReset(jumMurid, bulan, kas)
	case 2:
		editTambah(kasPerBulan, jumMurid, bulan, kas)
	case 3:
		editBayar(kasPerBulan, jumMurid, bulan, kas)
	case 4:
		editHapus(kasPerBulan, jumMurid, bulan, kas)
	case 0:
		return
	default:
		fmt.Println("  [!] Pilihan tidak valid.")
	}
}

func editReset(jumMurid *int, bulan *int, kas *tabData) {
	fmt.Println()
	fmt.Println("======== RESET DATA KAS ========")
	fmt.Println("PERINGATAN: Semua data akan dihapus!")

	if bacaKonfirmasi("Yakin reset? (y/n) : ") {
		*jumMurid = 0
		*bulan = 0
		*kas = tabData{}
		fmt.Println("Data berhasil direset.")
		return
	}

	fmt.Println("Reset dibatalkan.")
}

func editTambah(kasPerBulan *int, jumMurid *int, bulan *int, kas *tabData) {
	fmt.Println()
	fmt.Println("======== TAMBAH DATA MURID ========")

	if *jumMurid >= len(kas) {
		fmt.Println("  [!] Data penuh.")
		return
	}

	kas[*jumMurid].nim = bacaIntPositif("NIM          : ")
	kas[*jumMurid].nama = bacaString("Nama         : ")
	kas[*jumMurid].jumlahBayar = bacaIntMinNol("Jumlah Bayar : Rp ")
	*jumMurid = *jumMurid + 1

	hitungStatus(*kasPerBulan, jumMurid, bulan, kas)
	fmt.Println("Data berhasil ditambahkan.")
}

func editBayar(kasPerBulan *int, jumMurid *int, bulan *int, kas *tabData) {
	fmt.Println()
	fmt.Println("======== BAYAR UANG KAS ========")

	target := bacaIntPositif("Masukkan NIM : ")
	idx := cariIndexNIM(kas, *jumMurid, target)

	if idx == -1 {
		fmt.Println("  [!] NIM tidak ditemukan.")
		return
	}

	fmt.Printf("Ditemukan: %s (bayar saat ini: Rp %d)\n", kas[idx].nama, kas[idx].jumlahBayar)
	tambahan := bacaIntPositif("Jumlah Bayar Tambahan : Rp ")
	kas[idx].jumlahBayar += tambahan

	hitungStatus(*kasPerBulan, jumMurid, bulan, kas)
	fmt.Println("Pembayaran berhasil diperbarui.")
}

func editHapus(kasPerBulan *int, jumMurid *int, bulan *int, kas *tabData) {
	fmt.Println()
	fmt.Println("======== HAPUS DATA MURID ========")

	target := bacaIntPositif("Masukkan NIM : ")
	idx := cariIndexNIM(kas, *jumMurid, target)

	if idx == -1 {
		fmt.Println("  [!] NIM tidak ditemukan.")
		return
	}

	fmt.Printf("Data ditemukan: %s. ", kas[idx].nama)
	if !bacaKonfirmasi("Yakin hapus? (y/n) : ") {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	for i := idx; i < *jumMurid-1; i++ {
		kas[i] = kas[i+1]
	}
	kas[*jumMurid-1] = data{}
	*jumMurid = *jumMurid - 1

	hitungStatus(*kasPerBulan, jumMurid, bulan, kas)
	fmt.Println("Data berhasil dihapus.")
}
