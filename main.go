package main

import "fmt"

type data struct {
	jumlahBayar, nim, kurangBayar, lebihBayar int
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

func menu(jumMurid *int, bulan *int, kas *tabData) {
	var pilihan int
	for {
		fmt.Println()
		fmt.Println("========== MENU UTAMA ==========")
		fmt.Println("1. Input Uang Kas")
		fmt.Println("2. Tampilkan Semua")
		fmt.Println("3. Cek Status by NIM")
		fmt.Println("4. Cek Detail by NIM")
		fmt.Println("5. Urutkan Tampilan")
		fmt.Println("6. Reset Data Kas")
		fmt.Println("0. Exit")
		fmt.Println("================================")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		flushBuffer()

		switch pilihan {
		case 1:
			inputUangKas(jumMurid, bulan, kas)
		case 2:
			tampilkanSemua(jumMurid, bulan, kas)
		case 3:
			cekLunas(jumMurid, kas)
		case 4:
			cekNim(jumMurid, kas)
		case 5:
			urutkanTampilan(jumMurid, bulan, kas)
		case 6:
			resetData(jumMurid, bulan, kas)
		case 0:
			fmt.Println("Terima kasih. Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func inputUangKas(jumMurid *int, bulan *int, kas *tabData) {
	var kasPerBulan int

	fmt.Println()
	fmt.Println("======== INPUT UANG KAS ========")
	fmt.Print("Jumlah Murid : ")
	fmt.Scan(jumMurid)
	fmt.Print("Jumlah Bulan : ")
	fmt.Scan(bulan)
	fmt.Print("Kas per Bulan (Rp) : ")
	fmt.Scan(&kasPerBulan)

	fmt.Println()
	fmt.Println("Masukkan data (NIM - NAMA - JUMLAH_BAYAR) :")
	fmt.Println("Contoh: 123456 BudiSantoso 150000")
	for i := 0; i < *jumMurid; i++ {
		fmt.Printf("Data ke-%d : ", i+1)
		fmt.Scan(&kas[i].nim, &kas[i].nama, &kas[i].jumlahBayar)
	}

	hitungStatus(kasPerBulan, jumMurid, bulan, kas)
	fmt.Println()
	fmt.Println("Data berhasil diinput!")
	flushBuffer()
}

func flushBuffer() {
	var dummy string
	fmt.Scanln(&dummy)
}

func hitungStatus(kasPerBulan int, jumMurid *int, bulan *int, kas *tabData) {
	total := *bulan * kasPerBulan
	for i := 0; i < *jumMurid; i++ {
		selisih := kas[i].jumlahBayar - total
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

func tampilkanSemua(jumMurid *int, bulan *int, kas *tabData) {
	if *jumMurid == 0 {
		fmt.Println("\nBelum ada data. Masukkan data terlebih dahulu.")
		inputUangKas(jumMurid, bulan, kas)
		tampilkanSemua(jumMurid, bulan, kas)
		return
	}

	fmt.Println()
	fmt.Println("=========== TAMPILKAN SEMUA ===========")
	sep := "+------+----------+---------------+-----------------+-------------+--------------+--------------+"
	fmt.Println(sep)
	fmt.Printf("| %-4s | %-8s | %-13s | %-15s | %-11s | %-12s | %-12s |\n",
		"NO", "NIM", "NAMA", "JUMLAH BAYAR", "STATUS", "KURANG BAYAR", "LEBIH BAYAR")
	fmt.Println(sep)
	for i := 0; i < *jumMurid; i++ {
		fmt.Printf("| %-4d | %-8d | %-13s | %-15d | %-11s | %-12d | %-12d |\n",
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
	var totalLunas, totalBelumLunas, totalLebihBayar, totalKurangBayar int
	for i := 0; i < *jumMurid; i++ {
		if kas[i].status == "LUNAS" {
			totalLunas++
			totalLebihBayar += kas[i].lebihBayar
		} else {
			totalBelumLunas++
			totalKurangBayar += kas[i].kurangBayar
		}
	}
	fmt.Println("------- RINGKASAN -------")
	fmt.Printf("Total Murid   : %d\n", *jumMurid)
	fmt.Printf("Lunas         : %d murid\n", totalLunas)
	fmt.Printf("Belum Lunas   : %d murid\n", totalBelumLunas)
	fmt.Printf("Total Lebih Bayar  : Rp %d\n", totalLebihBayar)
	fmt.Printf("Total Kurang Bayar : Rp %d\n", totalKurangBayar)
	fmt.Println("-------------------------")
}

func cekLunas(jumMurid *int, kas *tabData) {
	if *jumMurid == 0 {
		fmt.Println("\nBelum ada data. Masukkan data terlebih dahulu.")
		var bulan int
		inputUangKas(jumMurid, &bulan, kas)
		cekLunas(jumMurid, kas)
		return
	}

	var target int
	for {
		fmt.Println()
		fmt.Println("======== CEK STATUS BY NIM ========")
		fmt.Print("NIM Dicari (0 untuk kembali) : ")
		fmt.Scan(&target)
		if target == 0 {
			return
		}

		found := false
		for i := 0; i < *jumMurid; i++ {
			if target == kas[i].nim {
				sep := "+----------+---------------+-------------+"
				fmt.Println(sep)
				fmt.Printf("| %-8s | %-13s | %-11s |\n", "NIM", "NAMA", "STATUS")
				fmt.Println(sep)
				fmt.Printf("| %-8d | %-13s | %-11s |\n", kas[i].nim, kas[i].nama, kas[i].status)
				fmt.Println(sep)
				found = true
			}
		}
		if !found {
			fmt.Println("NIM tidak ditemukan.")
		}
	}
}

func cekNim(jumMurid *int, kas *tabData) {
	if *jumMurid == 0 {
		fmt.Println("\nBelum ada data. Masukkan data terlebih dahulu.")
		var bulan int
		inputUangKas(jumMurid, &bulan, kas)
		cekNim(jumMurid, kas)
		return
	}

	var target int
	for {
		fmt.Println()
		fmt.Println("======== CEK DETAIL BY NIM ========")
		fmt.Print("NIM Dicari (0 untuk kembali) : ")
		fmt.Scan(&target)
		if target == 0 {
			return
		}

		left := 0
		right := *jumMurid - 1
		foundIdx := -1

		for left <= right {
			mid := left + (right-left)/2

			if kas[mid].nim == target {
				foundIdx = mid
				break
			} else if kas[mid].nim < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if foundIdx != -1 {
			sep := "+----------+---------------+-----------------+-------------+--------------+--------------+"
			fmt.Println(sep)
			fmt.Printf("| %-8s | %-13s | %-15s | %-11s | %-12s | %-12s |\n",
				"NIM", "NAMA", "JUMLAH BAYAR", "STATUS", "KURANG BAYAR", "LEBIH BAYAR")
			fmt.Println(sep)
			fmt.Printf("| %-8d | %-13s | %-15d | %-11s | %-12d | %-12d |\n",
				kas[foundIdx].nim,
				kas[foundIdx].nama,
				kas[foundIdx].jumlahBayar,
				kas[foundIdx].status,
				kas[foundIdx].kurangBayar,
				kas[foundIdx].lebihBayar,
			)
			fmt.Println(sep)
		} else {
			fmt.Println("NIM tidak ditemukan.")
		}
	}
}

func urutkanTampilan(jumMurid *int, bulan *int, kas *tabData) {
	if *jumMurid == 0 {
		fmt.Println("\nBelum ada data. Masukkan data terlebih dahulu.")
		inputUangKas(jumMurid, bulan, kas)
		urutkanTampilan(jumMurid, bulan, kas)
		return
	}

	var pilihan int
	fmt.Println()
	fmt.Println("======== URUTKAN TAMPILAN ========")
	fmt.Println("1. Urutkan by NIM (ascending)")
	fmt.Println("2. Belum Lunas dahulu")
	fmt.Println("0. Batal")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)
	flushBuffer()

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
		tampilkanSemua(jumMurid, bulan, kas)
	case 2:
		for i := 0; i < *jumMurid-1; i++ {
			for j := 0; j < *jumMurid-1-i; j++ {
				if kas[j].status == "LUNAS" && kas[j+1].status == "BELUM LUNAS" {
					kas[j], kas[j+1] = kas[j+1], kas[j]
				}
			}
		}
		fmt.Println("Data berhasil diurutkan: Belum Lunas dahulu.")
		tampilkanSemua(jumMurid, bulan, kas)
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func resetData(jumMurid *int, bulan *int, kas *tabData) {
	var konfirmasi string
	fmt.Println()
	fmt.Println("======== RESET DATA KAS ========")
	fmt.Println("PERINGATAN: Semua data akan dihapus!")
	fmt.Print("Yakin reset? (y/n) : ")
	fmt.Scan(&konfirmasi)
	flushBuffer()

	if konfirmasi == "y" || konfirmasi == "Y" {
		*jumMurid = 0
		*bulan = 0
		*kas = tabData{}
		fmt.Println("Data berhasil direset. Silakan input ulang.")
	} else {
		fmt.Println("Reset dibatalkan.")
	}
}
