package main

import (
	"fmt"
	"siparis/product"
)

/* type Product struct {
	Name  string
	Price int
	Cap   int
} */

func main() {
	var choice string
	/* var num int

	var Product1 Product
	var Product2 Product
	var Product3 Product
	var Product4 Product
	var Product5 Product

	Product1.Name = "Elbise"
	Product1.Price = 1000
	Product1.Cap = 100
	Product2.Name = "Çanta"
	Product2.Price = 300
	Product2.Cap = 40
	Product3.Name = "Kazak"
	Product3.Price = 500
	Product3.Cap = 80
	Product4.Name = "Ayakkabı"
	Product4.Price = 2000
	Product4.Cap = 20
	Product5.Name = "Takı"
	Product5.Price = 750
	Product5.Cap = 90  */

	fmt.Println("HOŞGELDİNİZ")
	for {
		fmt.Println("Almak istediğiniz ürünüzü seçiniz")
		fmt.Println("1-     Elbise(100) -> 1000 TL")
		fmt.Println("2-     Çanta(40)  -> 300 TL")
		fmt.Println("3-     Kazak(80)  -> 500 TL")
		fmt.Println("4-     Ayakkabı(20)  -> 2000 TL")
		fmt.Println("5-     Takı(90)  -> 750 TL")
		fmt.Println("6-     Sepetimi görmek istiyorum")
		fmt.Println("7-     Sepetimi onaylamak istiyorum")
		fmt.Println("8-     Çıkış")

		fmt.Scan(&choice)

		switch choice {
		case "1":
			/* fmt.Println("Elbiseden kaç tane almak istiyorsunuz")
			fmt.Scan(&num)
			if Product1.Cap >= num {
				dress := num * Product1.Price
				Product1.Cap -= num
				fmt.Println(Product1.Name, "Toplam Tutarı:  ", dress, "\nKalan Stok: ", Product1.Cap)
			} else {
				fmt.Println("Stokta o kadar ürün yoktur")
			}
			*/
			

		case "2":
			/* fmt.Println("Çantadan kaç tane almak istiyorsunuz")
			fmt.Scan(&num)

			if Product2.Cap >= num {
				purse := num * Product2.Price
				Product2.Cap -= num
				fmt.Println(Product2.Name, "Toplam Tutarı:  ", purse, "\nKalan Stok: ", Product2.Cap)
			} else {
				fmt.Println("Stokta o kadar ürün yoktur")
			} */
			product.Urun2(product.Product{})

		case "3":
			/* fmt.Println("Kazaktan kaç tane almak istiyorsunuz")
			fmt.Scan(&num)

			if Product3.Cap >= num {
				sweater := num * Product3.Price
				Product3.Cap -= num
				fmt.Println(Product3.Name, "Toplam Tutarı:  ", sweater, "\nKalan Stok: ", Product3.Cap)
			} else {
				fmt.Println("Stokta o kadar ürün yoktur")
			} */
			product.Urun3(product.Product{})

		case "4":
			/* fmt.Println("Ayakkabıdan kaç tane almak istiyorsunuz")
			fmt.Scan(&num)

			if Product4.Cap >= num {
				shoe := num * Product4.Price
				Product4.Cap -= num
				fmt.Println(Product4.Name, "Toplam Tutarı:  ", shoe, "\nKalan Stok: ", Product4.Cap)
			} else {
				fmt.Println("Stokta o kadar ürün yoktur")
			} */
			product.Urun4(product.Product{})

		case "5":
		/* 	fmt.Println("Takıdan kaç tane almak istiyorsunuz")
		fmt.Scan(&num)

		if Product5.Cap >= num {
			jewelry := num * Product5.Price
			Product5.Cap -= num
			fmt.Println(Product5.Name, "Toplam Tutarı:  ", jewelry, "\nKalan Stok: ", Product5.Cap)
		} else {
			fmt.Println("Stokta o kadar ürün yoktur")
		} */
		 			product.Urun5(product.Product{})


		case "6":
			fmt.Println("Sepetinizde bu ürünler vardır.")
			fmt.Println("Elbiselerin toplam tutarı: ")

		case "7":
			fmt.Println("Sipariş veriliyor...")

		case "8":
			fmt.Println("Çıkış yapılıyor...")
			return

		default:
			fmt.Println("Geçersiz işlem...")

		}

	}
}
