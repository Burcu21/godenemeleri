package main

import (
	"fmt"
	"siparis/models"
	"time"
)

func main() {
	var choice string

	fmt.Println("HOŞGELDİNİZ")

	for {
		fmt.Println("Yapmak istediğiniz işlemi seçiniz")
		fmt.Println("1-     Ürün seçmek istiyorum")
		fmt.Println("2-     Sepetteki ürünleri görmek istiyorum")
		fmt.Println("3-     Sepetimden ürün silmek istiyorum")
		fmt.Println("4-     Sipariş vermek istiyorum")
		fmt.Println("5-     Çıkış")

		fmt.Scan(&choice)

		switch choice {
		case "1":
			models.Urunler()
			fmt.Println()
			time.Sleep(1 * time.Second)
			models.SepetEkle()

		case "2":
			models.Sepet()
			fmt.Println()
			time.Sleep(1 * time.Second)
		case "3":
			models.SepetCikar()
			fmt.Println()
			time.Sleep(1 * time.Second)

		case "4":
			models.Sepet()
			fmt.Println()
			fmt.Println()
			time.Sleep(1 * time.Second)

			fmt.Println("Sipariş verilmiştir")
			fmt.Println()
			time.Sleep(1 * time.Second)

			return

		case "5":
			fmt.Println("Çıkış yapılıyor...")
			fmt.Println()
			time.Sleep(1 * time.Second)
			return
		default:
			fmt.Println("Yanlış seçim yaptınız!")

		}
	}
}
