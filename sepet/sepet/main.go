package main

import (
	"fmt"
	"spet/basket"
	"time"
)

func main() {
	var choice string
	var efruit string

	/* for _, v := range basket.Basket {
		fmt.Println(v)

		if v != input {
			fmt.Println("Sepette böyle bir yok ")
		}
	}
	*/
	fmt.Println("HOŞGELDİNİZ")

	for {
		fmt.Println("Lütfen yapmak istediğiniz işlemi seçin")
		fmt.Println("1. Sepetteki meyveleri görmek")
		fmt.Println("2. Sepete meyve eklemek")
		fmt.Println("3. Sepetten meyve çıkarmak")
		fmt.Println("4. Çıkış")
		fmt.Println()
		time.Sleep(1 * time.Second)
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case "4":
			fmt.Println("Çıkış yapılıyor...")
			fmt.Println()
			time.Sleep(1 * time.Second)
			return

		case "3":
			fmt.Println("Çıkarmak istediğiniz meyveyi yazınız")
			fmt.Scan(&efruit)
			fmt.Println()
			time.Sleep(1 * time.Second)

			/* for _, v := range basket.Basket {
			fmt.Println(v) */
			
			var v string 
			if v != efruit {
				fmt.Println("Sepette böyle bir meyve yok ")
				continue
			} else {

				for i, v := range basket.Basket {
					if v == efruit {
						basket.Basket = append(basket.Basket[:i], basket.Basket[i+1:]...)

						fmt.Printf("%s sepete çıkarılmıştır.", efruit)

						fmt.Println()
						break

					}
				}
			}

		case "2":
			fmt.Println("Eklemek istediğiniz meyveyi yazınız")
			fmt.Scan(&efruit)
			fmt.Println()
			time.Sleep(1 * time.Second)
			for _, v := range basket.Basket {
				fmt.Println(v)

				if v == efruit {
					fmt.Println("Sepette zaten bu meyve var ")
					continue
				} else {
					basket.Basket = append(basket.Basket, efruit)
					fmt.Printf("%s sepete eklenmiştir.", efruit)
					fmt.Println()
				}
			}

		case "1":
			fmt.Println("Sepetteki meyveler : ", basket.Basket)
			fmt.Println()
			time.Sleep(1 * time.Second)

		default:
			fmt.Println("Yanlış seçim yaptınız!")
			fmt.Println()

		}
	}
}
