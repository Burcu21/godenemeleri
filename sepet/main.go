package main

import (
	"fmt"
	"spet/basket"
	"time"
)

func main() {
	var efruit string
	var choice string

	fmt.Println("HOŞGELDİNİZ")

	for {
		fmt.Println("Lütfen seçim yapınız")
		fmt.Println("1. Mevcut meyveleri görmek ")
		fmt.Println("2. Meyve eklemek")
		fmt.Println("3. Meyve cikarma")
		fmt.Println("4. Cikis ")
		fmt.Scan(&choice)

		switch choice {
		case "1":

			fmt.Println("Mevcut meyveler:", basket.Basket)
			time.Sleep(1 * time.Second)
			fmt.Println()

		case "2":
			fmt.Println("Lütfen eklemek istediğiniz meyveyi yaziniz.")
			fmt.Scan(&efruit)
			time.Sleep(1 * time.Second)
			fmt.Println()

			basket.Basket = append(basket.Basket, efruit)
			fmt.Println()
			fmt.Printf("%s sepete eklenmiştir\n", efruit)

		case "3":
			fmt.Println("Lütfen cikarmak istediğiniz meyveyi yaziniz.")
			fmt.Scan(&efruit)

			for i, v := range basket.Basket {
				if v == efruit {
					basket.Basket = append(basket.Basket[:i], basket.Basket[i+1:]...)
				}
			}

		case "4":
			fmt.Println("Cikis yapılıyor ")
			fmt.Println()
			return

		default:
			fmt.Println("Yanlis secim yaptınız")
			fmt.Println()

		}
	}
}
