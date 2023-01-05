package main

import (
	"fmt"
	"os"
	"strconv"
)

const T = 20
const N = 20
const M = 20

var harga int = 5000

type kendaraan struct {
	nk         string
	jenis      string
	jam_masuk  string
	jam_keluar string
}
type arrMotor struct {
	area1               [T]kendaraan
	jumlah, jumlahMotor int
}
type arrMobil struct {
	area2                            [N][M]kendaraan
	jumlah, jumlahMobil, jumlahTruck int
}

var area1 arrMotor
var area2 arrMobil

func cekArea1() {
	for i := 0; i < T; i++ {
		fmt.Println(area1.area1[i])
	}
}

func cekArea2() {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Print(area2.area2[i][j])
		}
	}
}

func penuh() {
	var status1, status2 bool = false, false
	for i := 0; i < T && !status1; i++ {
		if area1.area1[i].nk == "" {
			status1 = true
		}
	}
	if status1 {
		fmt.Println("Area1 Kosong")
	} else {
		fmt.Println("Area1 Penuh")
	}
	for i := 0; i < N && !status2; i++ {
		for j := 0; j < M && !status2; j++ {
			if area2.area2[i][j].nk == "" {
				status2 = true
			}
		}
	}
	if status2 {
		fmt.Println("Area2 Kosong")
	} else {
		fmt.Println("Area2 Penuh")
	}
}

func bayarArea1(i int) int {
	area1masuk := area1.area1[i].jam_masuk
	area1keluar := area1.area1[i].jam_keluar

	jamMasuk := area1masuk[0:2]
	menitMasuk := area1masuk[3:5]

	jamKeluar := area1keluar[0:2]
	menitKeluar := area1keluar[3:5]

	jamIntMasuk, err := strconv.Atoi(jamMasuk)
	menitIntMasuk, err := strconv.Atoi(menitMasuk)

	jamIntKeluar, err := strconv.Atoi(jamKeluar)
	menitIntKeluar, err := strconv.Atoi(menitKeluar)

	durasijam := jamIntKeluar - jamIntMasuk
	durasimenit := menitIntKeluar - menitIntMasuk
	if err == nil {
		if (durasijam == 0) && (durasimenit >= 10) {
			harga = 5000
		} else {
			if jamIntMasuk > jamIntKeluar {
				durasijam = (jamIntKeluar + 24) - jamIntMasuk
				if durasimenit >= 10 || durasimenit <= -10 {
					harga = (durasijam + 1) * harga
				} else {
					harga = durasijam * harga
				}
			} else {
				if durasimenit >= 10 || durasimenit <= -10 {
					harga = (durasijam + 1) * harga
				} else {
					harga = durasijam * harga
				}
			}
		}
	}
	return harga
}

func bayarArea2(i, j int) int {
	area1masuk := area2.area2[i][j].jam_masuk
	area1keluar := area2.area2[i][j].jam_keluar

	jamMasuk := area1masuk[0:2]
	menitMasuk := area1masuk[3:5]

	jamKeluar := area1keluar[0:2]
	menitKeluar := area1keluar[3:5]

	jamIntMasuk, err := strconv.Atoi(jamMasuk)
	menitIntMasuk, err := strconv.Atoi(menitMasuk)

	jamIntKeluar, err := strconv.Atoi(jamKeluar)
	menitIntKeluar, err := strconv.Atoi(menitKeluar)

	durasijam := jamIntKeluar - jamIntMasuk
	durasimenit := menitIntKeluar - menitIntMasuk
	if err == nil {
		if (durasijam == 0) && (durasimenit >= 10) {
			harga = 5000
		} else {
			if jamIntMasuk > jamIntKeluar {
				durasijam = (jamIntKeluar + 24) - jamIntMasuk
				if durasimenit >= 10 || durasimenit <= -10 {
					harga = (durasijam + 1) * harga
				} else {
					harga = durasijam * harga
				}
			} else {
				if durasimenit >= 10 || durasimenit <= -10 {
					harga = (durasijam + 1) * harga
				} else {
					harga = durasijam * harga
				}
			}
		}
	}
	return harga
}

func cetakstrukArea1(i int) {
	fmt.Println("Jenis kendaraan :", area1.area1[i].jenis)
	fmt.Println("Nomor kendaraan :", area1.area1[i].nk)
	fmt.Println("Jam masuk kendaraan :", area1.area1[i].jam_masuk)
	fmt.Println("lokasi berada pada :", i)
	fmt.Println("Jam keluar kendaraan :", area1.area1[i].jam_keluar)
	fmt.Println("Biaya parkir: ", bayarArea1(i))
}

func cetakstrukArea2(i, j int) {
	fmt.Println("Jenis kendaraan :", area2.area2[i][j].jenis)
	fmt.Println("Nomor kendaraan :", area2.area2[i][j].nk)
	fmt.Println("Jam masuk kendaraan :", area2.area2[i][j].jam_masuk)
	fmt.Println("lokasi berada pada :", i, j)
	fmt.Println("Jam keluar kendaraan :", area2.area2[i][j].jam_keluar)
	fmt.Println("Biaya parkir: ", bayarArea2(i, j))
}

func isiYangKosong(kendaraan kendaraan) {
	var isi bool = false
	for i := 0; i < T && !isi; i++ {
		if area1.area1[i].nk == "" {
			area1.area1[i] = kendaraan
			area1.jumlah++
			area1.jumlahMotor++
			isi = true
			fmt.Println("kendaraan berhasil dimasukan")
		}
	}
}

func cari() {
	var kendaraan kendaraan
	fmt.Print("Silahkan masukkan jenis kendaraan : ")
	fmt.Scan(&kendaraan.jenis)
	fmt.Print("Silahkan masukkan nomor kendaraan (Format AA1234XXX): ")
	fmt.Scan(&kendaraan.nk)
	if kendaraan.jenis == "motor" {
		var isi bool = false
		for i := 0; i < T && !isi; i++ {
			if area1.area1[i].nk == kendaraan.nk {
				isi = true
			}
		}
		if !isi {
			fmt.Println("Tidak ada kendaraan dengan nomor kendaraan ", kendaraan.nk)
		} else {
			fmt.Println("Ada kendaraan dengan nomor kendaraan ", kendaraan.nk)
		}
	} else {
		if kendaraan.jenis == "mobil" {
			var isi bool = false
			for i := 0; i < N && !isi; i++ {
				for j := 0; j < M && !isi; j++ {
					if area2.area2[i][j].nk == kendaraan.nk {
						isi = true
					}
				}

			}
			if !isi {
				fmt.Println("Tidak ada kendaraan dengan nomor kendaraan ", kendaraan.nk)
			} else {
				fmt.Println("Ada kendaraan dengan nomor kendaraan ", kendaraan.nk)
			}
		} else {
			var isi bool = false
			for i := 0; i < N && !isi; i++ {
				for j := 0; j < M && !isi; j++ {
					if area2.area2[i][j].nk == kendaraan.nk {
						isi = true
					}
				}
			}
			if !isi {
				fmt.Println("Tidak ada kendaraan dengan nomor kendaraan ", kendaraan.nk)
			} else {
				fmt.Println("Ada kendaraan dengan nomor kendaraan ", kendaraan.nk)
			}

		}
	}
}

func kendaraanMasuk() {
	var kendaraan kendaraan
	fmt.Print("Silahkan masukkan jenis kendaraan : ")
	fmt.Scan(&kendaraan.jenis)
	fmt.Print("Silahkan masukkan nomor kendaraan (Format AA1234XXX) : ")
	fmt.Scan(&kendaraan.nk)
	fmt.Print("Silahkan masukkan jam masuk kendaraan (format HH:MM): ")
	fmt.Scan(&kendaraan.jam_masuk)
	if kendaraan.jenis == "motor" {
		isiYangKosong(kendaraan)
		cekArea1()
	} else {
		if kendaraan.jenis == "mobil" {
			var isi bool = false
			for i := 0; i < N && !isi; i++ {
				for j := 0; j < M && !isi; j++ {
					if area2.area2[i][j].nk == "" {
						area2.area2[i][j] = kendaraan
						area2.jumlah++
						area2.jumlahMobil++
						isi = true
						fmt.Println("kendaraan berhasil dimasukan")
					}
				}
			}
			cekArea2()
		} else {
			var isi bool = false
			for i := 0; i < N && !isi; i++ {
				if area2.area2[i][M-2].nk == "" {
					for j := 0; j < M && !isi; j++ {
						if area2.area2[i][j].nk == "" {
							area2.area2[i][j] = kendaraan
							area2.area2[i][j+1] = kendaraan
							area2.jumlah += 2
							area2.jumlahTruck++
							isi = true
							fmt.Println("kendaraan berhasil dimasukan")
						}
					}
				}
			}
			cekArea2()
		}
	}
}

func kendaraanKeluar() {

	var kendaraan kendaraan
	fmt.Print("Masukkan jenis kendaraan : ")
	fmt.Scan(&kendaraan.jenis)
	fmt.Print("Masukkan nomor kendaraan (Format AA1234XXX): ")
	fmt.Scan(&kendaraan.nk)
	if kendaraan.jenis == "motor" {
		var isi bool = false
		for i := 0; i < T && !isi; i++ {
			if area1.area1[i].nk == kendaraan.nk {
				fmt.Print("Jam keluar kendaraan (format HH:MM): ")
				fmt.Scan(&area1.area1[i].jam_keluar)
				cetakstrukArea1(i)
				catatMotor(i)
				area1.area1[i].jenis = ""
				area1.area1[i].nk = ""
				area1.area1[i].jam_masuk = ""
				area1.area1[i].jam_keluar = ""
				area1.jumlah--
				isi = true
				fmt.Println("berhasil keluar dari parkiran")
			}
		}
		cekArea1()
	} else {
		if kendaraan.jenis == "mobil" {
			var isi bool = false
			for i := 0; i < N && !isi; i++ {
				for j := 0; j < M && !isi; j++ {
					if area2.area2[i][j].nk == kendaraan.nk {
						fmt.Print("Jam keluar kendaraan (format HH:MM): ")
						fmt.Scan(&area2.area2[i][j].jam_keluar)
						cetakstrukArea2(i, j)
						catatMobil(i, j)
						area2.area2[i][j].jenis = ""
						area2.area2[i][j].nk = ""
						area2.area2[i][j].jam_masuk = ""
						area2.area2[i][j].jam_keluar = ""
						isi = true
						area2.jumlah--
						fmt.Println("berhasil keluar dari parkiran")
					}
				}
			}
			cekArea2()
		} else {
			var isi bool = false
			for i := 0; i < N && !isi; i++ {
				for j := 0; j < M && !isi; j++ {
					if area2.area2[i][j].nk == kendaraan.nk {
						fmt.Print("Jam keluar kendaraan (format HH:MM): ")
						fmt.Scan(&area2.area2[i][j].jam_keluar)
						cetakstrukArea2(i, j)
						catatTruk(i, j)
						area2.area2[i][j].jenis = ""
						area2.area2[i][j].nk = ""
						area2.area2[i][j].jam_masuk = ""
						area2.area2[i][j].jam_keluar = ""
						area2.area2[i][j+1].jenis = ""
						area2.area2[i][j+1].nk = ""
						area2.area2[i][j+1].jam_masuk = ""
						area2.area2[i][j+1].jam_keluar = ""
						isi = true
						area2.jumlah -= 2
						fmt.Println("berhasil keluar dari parkiran")
					}
				}
			}
			cekArea2()
		}
	}
}

func nomorKendaraan() {
	var cari string
	var i, j int
	fmt.Println("pilih jenis kendaraan yang mau diurutkan (motor/mobil)")
	fmt.Scan(&cari)
	if cari == "motor" {
		i = 1
		for i < area1.jumlah {
			j = i - 1
			a, err := strconv.Atoi(area1.area1[j].nk[2:6])
			b, err := strconv.Atoi(area1.area1[j+1].nk[2:6])
			if err == nil {
				for j >= 0 && a > b {
					temp := area1.area1[j]
					area1.area1[j] = area1.area1[j+1]
					area1.area1[j+1] = temp
					j--
				}
			}
			i++
		}
		i = 0
		n := area1.jumlah
		for i < n {
			if area1.area1[i].nk != "" {
				fmt.Print(area1.area1[i], " ")
			}
			i++
		}

	}
}

func hitungSehari() {
	Motor := area1.jumlahMotor * harga
	Mobil := area2.jumlahMobil * harga
	Truck := area2.jumlahTruck * harga
	Jumlah := Motor + Mobil + Truck
	fmt.Println("Pendapatan Hari ini")
	fmt.Println("Pendapatan Parkir Motor: ", Motor)
	fmt.Println("Pendapatan Parkir Mobil: ", Mobil)
	fmt.Println("Pendapatan Parkir Kendaraan Besar: ", Truck)
	fmt.Println("Pendapatan total : ", Jumlah)
}

func statistik() {
	fmt.Println("Total kendaraan yang parkir pada hari ini: ", area1.jumlahMotor+area2.jumlahMobil+area2.jumlahTruck)
	fmt.Println("Dengan detail: ")
	fmt.Println("Kendaraan Motor:", area1.jumlahMotor)
	fmt.Println("kendaraan kecil:", area2.jumlahMobil)
	fmt.Println("kendaraan besar:", area2.jumlahTruck)
}

func presentase() {
	area1 := area1.jumlahMotor / T * 100
	area2 := (area2.jumlahMobil + (area2.jumlahMobil * 2)) / (N * M) * 100
	fmt.Println("presentase okupansi area 1 adalah", area1, "%")
	fmt.Println("presentase okupansi area 2 adalah", area2, "%")
}

func catatMobil(i, j int) {
	f, err := os.OpenFile("kendaraan kecil.txt", os.O_APPEND|os.O_CREATE, 0600)

	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("jenis kendaraan = %s\nplat kendaraan = %s\njam masuk kendaraan = %s\njam keluar kendaraan = %s \n\n", area2.area2[i][j].jenis, area2.area2[i][j].nk, area2.area2[i][j].jam_masuk, area2.area2[i][j].jam_keluar))
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}

}
func catatTruk(i, j int) {
	f, err := os.OpenFile("kendaraan besar.txt", os.O_APPEND|os.O_CREATE, 0600)

	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("jenis kendaraan = %s\nplat kendaraan = %s\njam masuk kendaraan = %s\njam keluar kendaraan = %s \n\n", area2.area2[i][j].jenis, area2.area2[i][j].nk, area2.area2[i][j].jam_masuk, area2.area2[i][j].jam_keluar))
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}

}
func catatMotor(i int) {
	f, err := os.OpenFile("motor.txt", os.O_APPEND|os.O_CREATE, 0600)

	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("jenis kendaraan = %s\nplat kendaraan = %s\njam masuk kendaraan = %s\njam keluar kendaraan = %s \n\n", area1.area1[i].jenis, area1.area1[i].nk, area1.area1[i].jam_masuk, area1.area1[i].jam_keluar))
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}

}

func main() {
	var run bool = true
	var pilih int
	for run {
		fmt.Println("SELAMAT DATANG DI PARKIRAN")
		fmt.Println()
		fmt.Println("Silakan pilih opsi berikut untuk mengakses fitur-fitur berikut ini:")
		fmt.Println()
		fmt.Println("Pilih 1 untuk masuk parkiran")
		fmt.Println("Pilih 2 untuk keluar parkiran")
		fmt.Println("Pilih 3 untuk mencari kendaraan tertentu")
		fmt.Println("Pilih 4 untuk mengecek penuh atau tidak")
		fmt.Println("Pilih 5 untuk menampilkan pendapatan sehari")
		fmt.Println("Pilih 6 untuk mengurutkan kendaraan berdasarkan nomor kendaraan")
		fmt.Println("Pilih 7 untuk menampilkan presentase isi parkiran")
		fmt.Println("Pilih 8 untuk menampilkan statistik kendaraan")
		fmt.Println("Pilih 0 untuk keluar")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			kendaraanMasuk()
		case 2:
			kendaraanKeluar()
		case 3:
			cari()
		case 4:
			penuh()
		case 5:
			hitungSehari()
		case 6:
			nomorKendaraan()
		case 7:
			presentase()
		case 8:
			statistik()
		default:
			fmt.Println("keluar")
			run = false
		}
	}
}
