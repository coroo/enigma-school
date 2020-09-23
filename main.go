package main

import (
	"database/sql"
	"fmt"
	"github/coroo/enigma-school/config"
	"github/coroo/enigma-school/deliveries"
	"github/coroo/enigma-school/models"
	"github/coroo/enigma-school/repositories"
	"github/coroo/enigma-school/usecases"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Empty struct {
}
type app struct {
	db *sql.DB
}
type menuChoosed string

func newApp() app {
	c := config.NewConfig()
	err := c.InitDb()
	if err != nil {
		panic(err)
	}
	myapp := app{
		db: c.Db,
	}
	return myapp
}

func MainMenuForm() {
	var appMenu = map[string]string{
		"01": "Master Pelajaran",
		"02": "Data Mahasiswa",
		"03": "Assign Mata Kuliah",
		"q":  "Exit",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Main Menu Application")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range MenuChoiceMahasiswaed(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}
func SubMainMenuForm(main menuChoosed) {
	var menuName string
	var subAppMenu map[string]string
	subAppMenu = map[string]string{}

	if main == "01" {
		menuName = "Master Pelajaran"
		subAppMenu["A"] = "Tambah pelajaran"
		subAppMenu["B"] = "Lihat semua pelajaran"
		subAppMenu["C"] = "Detail pelajaran"
		subAppMenu["D"] = "Ubah data pelajaran"
		subAppMenu["E"] = "Hapus pelajaran"
		subAppMenu["q"] = "Back to Main Menu"
	} else if main == "02" {
		menuName = "Data Mahasiswa"
		subAppMenu["A"] = "Tambah mahasiswa"
		subAppMenu["B"] = "Detail mahasiswa"
		subAppMenu["C"] = "Ubah data mahasiswa"
		subAppMenu["D"] = "Hapus mahasiswa"
		subAppMenu["q"] = "Back to Main Menu"
	} else if main == "03" {
		menuName = "Penempatan Mahasiswa"
		subAppMenu["A"] = "Penempatan mahasiswa"
		subAppMenu["B"] = "Lihat penempatan mahasiswa"
		subAppMenu["C"] = "Hapus penempatan mahasiswa"
		subAppMenu["q"] = "Back to Main Menu"
	} else {
		menuName = "Unknown Menu"
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", menuName)
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range MenuChoiceMahasiswaed(subAppMenu) {
		fmt.Printf("%s. %s\n", menuCode, subAppMenu[menuCode])
	}
}
func MenuChoiceMahasiswaed(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
func (a app) Run() {
	var isExist = false
	var userChoice string

	MainMenuForm()

	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			var isBack = false
			var userChoice string

			SubMainMenuForm("01")
			repo := repositories.NewPelajaranRepository(a.db)
			usecase := usecases.NewPelajaranUseCase(repo)
			pelajaranDelivery := deliveries.NewPelajaranDelivery(usecase)

			for isBack == false {
				fmt.Printf("\n%s", "Your Next Choice: ")
				fmt.Scan(&userChoice)
				switch {
				case userChoice == "A":
					var kodePelajaran string
					var namaPelajaran string
					var jumlahSks int
					fmt.Printf("\n%s", "Kode Pelajaran: ")
					fmt.Scan(&kodePelajaran)
					fmt.Printf("%s", "Nama Pelajaran: ")
					fmt.Scan(&namaPelajaran)
					fmt.Printf("%s", "Jumlah SKS: ")
					fmt.Scan(&jumlahSks)
					_, err := pelajaranDelivery.RegisterNewPelajaran(models.Pelajaran{
						KodePelajaran: kodePelajaran,
						NamaPelajaran: namaPelajaran,
						JumlahSks:     jumlahSks,
					})
					if err != nil {
						panic(err)
					}
					isBack = true
					fmt.Printf("%s", "Data pelajaran telah tersimpan")
					isExist = true
				case userChoice == "B":
					fmt.Printf("\n%s", "Lihat Semua Pelajaran")
					result, _ := pelajaranDelivery.GetPelajaranAll()
					pelajaranDelivery.PrintPelajaran(result)
					isBack = true
					isExist = true
				case userChoice == "C":
					fmt.Printf("\n%s", "Detail Pelajaran")
					var kodePelajaran string
					fmt.Printf("\n%s", "ID Pelajaran: ")
					fmt.Scan(&kodePelajaran)
					result, _ := pelajaranDelivery.GetPelajaranById(kodePelajaran)
					pelajaranDelivery.PrintOnePelajaran(result)
					isBack = true
					isExist = true
				case userChoice == "D":
					var kodePelajaran string
					var namaPelajaran string
					var jumlahSks int
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&kodePelajaran)
					_, checkErr := pelajaranDelivery.GetPelajaranById(kodePelajaran)
					if checkErr != nil {
						fmt.Printf("Pelajaran not found")
						fmt.Printf("\n%s\n\n", strings.Repeat("-", 30))
						isBack = true
						MainMenuForm()
					}
					fmt.Printf("%s", "Nama Pelajaran: ")
					fmt.Scan(&namaPelajaran)
					fmt.Printf("%s", "Jumlah SKS: ")
					fmt.Scan(&jumlahSks)
					_, err := pelajaranDelivery.UpdateExistingPelajaran(models.Pelajaran{
						KodePelajaran: kodePelajaran,
						NamaPelajaran: namaPelajaran,
						JumlahSks:     jumlahSks,
					})
					if err != nil {
						panic(err)
					}
					isBack = true
					fmt.Printf("%s", "Perubahan data pelajaran telah tersimpan")
					isExist = true
				case userChoice == "E":
					fmt.Printf("\n%s", "Delete Pelajaran")
					var kodePelajaran string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&kodePelajaran)
					_, _ = pelajaranDelivery.DeletePelajaranById(kodePelajaran)
					isBack = true
					fmt.Printf("%s", "Data pelajaran telah dihapus")
					isExist = true
				case userChoice == "q":
					isBack = true
					MainMenuForm()
				default:
					fmt.Println("Unknown Menu Code")
				}
			}
		case userChoice == "02":
			var isBack = false
			var userChoice string

			SubMainMenuForm("02")
			repo := repositories.NewMahasiswaRepository(a.db)
			usecase := usecases.NewMahasiswaUseCase(repo)
			mahasiswaDelivery := deliveries.NewMahasiswaDelivery(usecase)

			for isBack == false {
				fmt.Printf("\n%s", "Your Next Choice: ")
				fmt.Scan(&userChoice)
				switch {
				case userChoice == "A":
					var nim string
					var namaMahasiswa string
					var jurusanMahasiswa string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&nim)
					fmt.Printf("%s", "Nama Mahasiswa: ")
					fmt.Scan(&namaMahasiswa)
					fmt.Printf("%s", "Jurusan Mahasiswa: ")
					fmt.Scan(&jurusanMahasiswa)
					_, err := mahasiswaDelivery.RegisterNewMahasiswa(models.MahasiswaWithPelajaran{
						Mahasiswa: models.Mahasiswa{
							Nim:              nim,
							NamaMahasiswa:    namaMahasiswa,
							JurusanMahasiswa: jurusanMahasiswa,
						},
					})
					if err != nil {
						panic(err)
					}
					isBack = true
					fmt.Printf("%s", "Data mahasiswa telah tersimpan")
					isExist = true
				case userChoice == "B":
					fmt.Printf("\n%s", "Detail Mahasiswa")
					var nim string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&nim)
					result, _ := mahasiswaDelivery.GetMahasiswaById(nim)
					mahasiswaDelivery.PrintOneMahasiswa(result)
					isBack = true
					isExist = true
				case userChoice == "C":
					var nim string
					var namaMahasiswa string
					var jurusanMahasiswa string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&nim)
					_, checkErr := mahasiswaDelivery.GetMahasiswaById(nim)
					if checkErr != nil {
						fmt.Printf("Mahasiswa not found")
						fmt.Printf("\n%s\n\n", strings.Repeat("-", 30))
						isBack = true
						MainMenuForm()
						break
					}
					fmt.Printf("%s", "Nama Mahasiswa: ")
					fmt.Scan(&namaMahasiswa)
					fmt.Printf("%s", "Jurusan Mahasiswa: ")
					fmt.Scan(&jurusanMahasiswa)
					_, err := mahasiswaDelivery.UpdateExistingMahasiswa(models.MahasiswaWithPelajaran{
						Mahasiswa: models.Mahasiswa{
							Nim:              nim,
							NamaMahasiswa:    namaMahasiswa,
							JurusanMahasiswa: jurusanMahasiswa,
						},
					})
					if err != nil {
						panic(err)
					}
					isBack = true
					fmt.Printf("%s", "Perubahan data mahasiswa telah tersimpan")
					isExist = true
				case userChoice == "D":
					fmt.Printf("\n%s", "Delete Mahasiswa")
					var nim string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&nim)
					_, _ = mahasiswaDelivery.DeleteMahasiswaById(nim)
					isBack = true
					fmt.Printf("%s", "Data mahasiswa telah dihapus")
					isExist = true
				case userChoice == "q":
					isBack = true
					MainMenuForm()
				default:
					fmt.Println("Unknown Menu Code")
				}
			}
		case userChoice == "03":
			var isBack = false
			var userChoice string

			SubMainMenuForm("03")
			repo := repositories.NewPenempatanRepository(a.db)
			usecase := usecases.NewPenempatanUseCase(repo)
			penempatanDelivery := deliveries.NewPenempatanDelivery(usecase)

			mRepo := repositories.NewMahasiswaRepository(a.db)
			mUsecase := usecases.NewMahasiswaUseCase(mRepo)
			mahasiswaDelivery := deliveries.NewMahasiswaDelivery(mUsecase)

			pRepo := repositories.NewPelajaranRepository(a.db)
			pUsecase := usecases.NewPelajaranUseCase(pRepo)
			pelajaranDelivery := deliveries.NewPelajaranDelivery(pUsecase)

			for isBack == false {
				fmt.Printf("\n%s", "Your Next Choice: ")
				fmt.Scan(&userChoice)
				switch {
				case userChoice == "A":
					var nim string
					var kodePelajaran string
					var tahunAjaran string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&nim)
					_, checkMahasiswa := mahasiswaDelivery.GetMahasiswaById(nim)
					if checkMahasiswa != nil {
						fmt.Printf("Mahasiswa not found")
						fmt.Printf("\n%s\n\n", strings.Repeat("-", 30))
						isBack = true
						MainMenuForm()
						break
					}

					_, checkErr := penempatanDelivery.GetPenempatanByMahasiswaNim(nim)
					if checkErr != nil {
						fmt.Printf("Mahasiswa not found")
						fmt.Printf("\n%s\n\n", strings.Repeat("-", 30))
						isBack = true
						MainMenuForm()
						break
					}
					fmt.Printf("%s", "Kode Pelajaran: ")
					fmt.Scan(&kodePelajaran)
					_, pErr := pelajaranDelivery.GetPelajaranById(kodePelajaran)
					if pErr != nil {
						fmt.Printf("Kode pelajar not found")
						fmt.Printf("\nPlease find the correct kode pelajaran at: 01. Master Pelajaran")
						fmt.Printf("\n%s\n\n", strings.Repeat("-", 30))
						isBack = true
						MainMenuForm()
						break
					}
					fmt.Printf("%s", "Tahun Ajaran: ")
					fmt.Scan(&tahunAjaran)
					_, err := penempatanDelivery.RegisterNewPenempatan(models.Penempatan{
						Nim:           nim,
						KodePelajaran: kodePelajaran,
						TahunAjaran:   tahunAjaran,
					})
					if err != nil {
						panic(err)
					}
					isBack = true
					fmt.Printf("%s", "Data penempatan telah tersimpan")
					isExist = true
				case userChoice == "B":
					var nim string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&nim)

					mList, checkErr := penempatanDelivery.GetPenempatanByMahasiswaNim(nim)
					if checkErr != nil {
						fmt.Printf("Mahasiswa not found")
						fmt.Printf("\n%s\n\n", strings.Repeat("-", 30))
						isBack = true
						MainMenuForm()
						break
					}
					penempatanDelivery.PrintPenempatan(mList)
					isBack = true
					isExist = true
				case userChoice == "C":
					fmt.Printf("\n%s", "Delete Penempatan")
					var nim string
					fmt.Printf("\n%s", "NIM: ")
					fmt.Scan(&nim)
					_, _ = penempatanDelivery.DeletePenempatanById(nim)
					isBack = true
					fmt.Printf("%s", "Data penempatan telah dihapus")
					isExist = true
				case userChoice == "q":
					isBack = true
					MainMenuForm()
				default:
					fmt.Println("Unknown Menu Code")
				}
			}
		case userChoice == "q":
			isExist = true
		default:
			fmt.Println("Unknown Menu Code")

		}
	}
}

func main() {
	newApp().Run()
}
