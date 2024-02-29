package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Name       string
	Address    string
	Employment string
	Reason     string
}

func GetDataByAbsen(absen int) Friend {
	dataFriend := map[int]Friend{
		1: {"Brian", "Jalan jambu air 1", "Web Developer", "Karena memiliki kemampuan untuk berkomunikasi dengan aman antar saluran"},
		2: {"Rifqi", "Jalan anggur 2", "IOS Developer", "Karena jumlah pengguna golang sangat banyak"},
		3: {"Ali", "Jalan pisang 3", "Backend Developer", "Karena cocok untuk pengembangan backend yang menanggani permintaan berat"},
		4: {"Ikhsan", "Jalan apel 4", "Data Science", "Karena golang memudahkan pengguna dalam menangani I/O"},
		5: {"Brian", "Jalan jambu air 1", "Web Developer", "Karena mirip dengan bahasa C"},
	}

	friend, found := dataFriend[absen]
	if !found {
		fmt.Println("Tidak ditemukan data teman dengan nomer absen", absen)
		os.Exit(1)
	}

	return friend
}

func PrintDataFriend(friend Friend) {
	fmt.Printf("Name: %s\n", friend.Name)
	fmt.Printf("Address: %s\n", friend.Address)
	fmt.Printf("Employment: %s\n", friend.Employment)
	fmt.Printf("Reason: %s\n", friend.Reason)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Gunakan: go run biodata. <nomer_absen>")
		os.Exit(1)
	}

	absen := os.Args[1]

	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomer absen harus berupa angka")
		os.Exit(1)
	}

	friend := GetDataByAbsen(absenInt)

	PrintDataFriend(friend)
}