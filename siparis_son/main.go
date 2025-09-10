package main

import (
	"fmt"
	"time"
)

type Product struct {
	Name  string
	Price int
	Cap   int
	Stok  int
	Total int
}

func (a *Product) urun(num int) bool {
	if a.Cap >= num {
		a.Cap -= num
		price := a.Price * num

		fmt.Println(num, "ürün sepete eklenmiştir")
		fmt.Println("Toplam tutar : ", price)
		return true

	} else {
		fmt.Println("Stok yetersiz!")
		return false

	}

}

func (a *Product) clear(num2 int) bool {
	fmt.Println(a.Name, "ne kadar sepetten çıkarılsın?")
	fmt.Println("Sepetinizde", a.Stok-a.Cap, "tane ürün vardır.")
	fmt.Scan(&num2)
	sepet := a.Stok - a.Cap

	if sepet >= num2 {
		a.Cap += num2
		fmt.Println(a.Name, "ürününden", num2, "miktar kadar sepetten çıkarılmıştır")
		return true
	} else {
		fmt.Println(a.Name, "üründen sepette o kadar miktar yoktur.")
		return false
	}

}

func (a *Product) order() {
	//var total int = 0
	num3 := a.Price * (a.Stok - a.Cap) // tutar
	num4 := a.Stok - a.Cap
	//total = +num3

	if num4 == 0 {
	} else {
		fmt.Println("--", a.Name, "--", "\nMiktarı:", num4, "\nFiyatı: ", a.Price, "\nTutar: ", num3, "\n----------")

	}

}

func main() {
	var choice string
	var num int

	p1 := &Product{Cap: 100, Price: 1000, Stok: 100, Name: "Elbise", Total: 100000}
	p2 := &Product{Cap: 40, Price: 300, Stok: 40, Name: "Çanta", Total: 12000}
	p3 := &Product{Cap: 80, Price: 500, Stok: 80, Name: "Kazak", Total: 40000}
	p4 := &Product{Cap: 20, Price: 2000, Stok: 20, Name: "Ayakkabı", Total: 40000}
	p5 := &Product{Cap: 90, Price: 750, Stok: 90, Name: "Takı", Total: 67500}

	fmt.Println("HOŞGELDİNİZ")

	for {
		fmt.Println("Yapmak istediğiniz işlemi seçiniz")
		fmt.Println("1-     Ürün seçmek istiyorum")
		fmt.Println("2-     Sepetteki ürünleri görmek istiyorum")
		fmt.Println("3-     Kalan stoğu görmek istiyorum")
		fmt.Println("4-     Sepetimden ürün silmek istiyorum")
		fmt.Println("5-     Sipariş vermek istiyorum")
		fmt.Println("6-     Çıkış")

		fmt.Scan(&choice)

	randloop:

		switch choice {
		case "1":
			for {
				fmt.Println("Hangi üründen almak istiyorsunuz")
				fmt.Println("1-    Ürün:", p1.Name, "- Stok:", (p1.Cap), "- Fiyat:", p1.Price, "TL")
				fmt.Println("2-    Ürün:", p2.Name, "- Stok:", (p2.Cap), "- Fiyat:", p2.Price, "TL")
				fmt.Println("3-    Ürün:", p3.Name, "- Stok:", (p3.Cap), "- Fiyat:", p3.Price, "TL")
				fmt.Println("4-    Ürün:", p4.Name, "- Stok:", (p4.Cap), "- Fiyat:", p4.Price, "TL")
				fmt.Println("5-    Ürün:", p5.Name, "- Stok:", (p5.Cap), "- Fiyat:", p5.Price, "TL")
				fmt.Println("6-    Menüye dön")

				fmt.Scan(&choice)
				fmt.Println()
				time.Sleep(1 * time.Second)

				switch choice {
				case "1":
					fmt.Println("Elbiseden kaç tane almak istiyorsunuz")
					fmt.Scan(&num)
					fmt.Println()
					time.Sleep(1 * time.Second)

					p1.urun(num)
					fmt.Println()

				case "2":
					fmt.Println("Çantadan kaç tane almak istiyorsunuz")
					fmt.Scan(&num)
					fmt.Println()
					time.Sleep(1 * time.Second)

					p2.urun(num)
					fmt.Println()

				case "3":
					fmt.Println("Kazaktan kaç tane almak istiyorsunuz")
					fmt.Scan(&num)
					fmt.Println()
					time.Sleep(1 * time.Second)

					p3.urun(num)
					fmt.Println()

				case "4":
					fmt.Println("Ayakkabıdan kaç tane almak istiyorsunuz")
					fmt.Scan(&num)
					fmt.Println()
					time.Sleep(1 * time.Second)

					p4.urun(num)
					fmt.Println()

				case "5":
					fmt.Println("Takıdan kaç tane almak istiyorsunuz")
					fmt.Scan(&num)
					fmt.Println()
					time.Sleep(1 * time.Second)

					p5.urun(num)
					fmt.Println()

				case "6":
					fmt.Println("Menüye dönülüyor...")
					fmt.Println()
					time.Sleep(1 * time.Second)
					break randloop

				default:
					fmt.Println("Yanlış seçim yaptınız.")
					fmt.Println()
					time.Sleep(1 * time.Second)

				}
			}

		case "2":
			fmt.Println("Sepetinizdeki ürünler:")
			fmt.Println()
			time.Sleep(1 * time.Second)

			Total1 := p1.Total - (p1.Price * p1.Cap)
			Total2 := p2.Total - (p2.Price * p2.Cap)
			Total3 := p3.Total - (p3.Price * p3.Cap)
			Total4 := p4.Total - (p4.Price * p4.Cap)
			Total5 := p5.Total - (p5.Price * p5.Cap)

			Total := Total1 + Total2 + Total3 + Total4 + Total5

			p1.order()
			p2.order()
			p3.order()
			p4.order()
			p5.order()

			fmt.Println("Sepet Tutarınız: ", Total)
			fmt.Println()
			time.Sleep(1 * time.Second)

		case "3":
			fmt.Println("Kalan Stok")
			fmt.Println()
			time.Sleep(1 * time.Second)

			fmt.Println("Elbise  : ", p1.Cap)
			fmt.Println("Çanta   : ", p2.Cap)
			fmt.Println("Kazak   : ", p3.Cap)
			fmt.Println("Ayakkabı: ", p4.Cap)
			fmt.Println("Takı    : ", p5.Cap)

		case "4":
		randloop2:

			for {
				fmt.Println("Hangi ürünü sepetten silmek istiyorsunuz")
				fmt.Println("1- ", p1.Name, "--Sepetinizdeki ürün miktarı: ", p1.Stok-p1.Cap)
				fmt.Println("2- ", p2.Name, "--Sepetinizdeki ürün miktarı: ", p2.Stok-p2.Cap)
				fmt.Println("3- ", p3.Name, "--Sepetinizdeki ürün miktarı: ", p3.Stok-p3.Cap)
				fmt.Println("4- ", p4.Name, "--Sepetinizdeki ürün miktarı: ", p4.Stok-p4.Cap)
				fmt.Println("5- ", p5.Name, "--Sepetinizdeki ürün miktarı: ", p5.Stok-p5.Cap)
				fmt.Println("6-  Ana menüye dön")

				fmt.Scan(&choice)

				switch choice {
				case "1":
					p1.clear(num)
					fmt.Println()
					time.Sleep(1 * time.Second)
				case "2":
					p2.clear(num)
					fmt.Println()
					time.Sleep(1 * time.Second)
				case "3":
					p3.clear(num)
					fmt.Println()
					time.Sleep(1 * time.Second)
				case "4":
					p4.clear(num)
					fmt.Println()
					time.Sleep(1 * time.Second)
				case "5":
					p5.clear(num)
					fmt.Println()
					time.Sleep(1 * time.Second)
				case "6":
					fmt.Println("Dönüş yapılıyor...")
					fmt.Println()
					time.Sleep(1 * time.Second)
					break randloop2
				default:
					fmt.Println("Geçersiz işlem...")
					fmt.Println()

				}
			}
		case "5":
			fmt.Println("Sipariş veriliyor...")
			fmt.Println()

			Total1 := p1.Total - (p1.Price * p1.Cap)
			Total2 := p2.Total - (p2.Price * p2.Cap)
			Total3 := p3.Total - (p3.Price * p3.Cap)
			Total4 := p4.Total - (p4.Price * p4.Cap)
			Total5 := p5.Total - (p5.Price * p5.Cap)

			Total := Total1 + Total2 + Total3 + Total4 + Total5

			p1.order()
			p2.order()
			p3.order()
			p4.order()
			p5.order()

			fmt.Println("Sepet Tutarınız: ", Total)
			fmt.Println()
			time.Sleep(1 * time.Second)

			time.Sleep(1 * time.Second)
			fmt.Println("Siparişiniz verilmiştir.")
			fmt.Println()
			time.Sleep(1 * time.Second)

		case "6":
			fmt.Println("Çıkış yapılıyor...")
			fmt.Println()
			time.Sleep(1 * time.Second)
			fmt.Println("Çıkış yapıldı.")
			fmt.Println()
			time.Sleep(1 * time.Second)

			return

		default:
			fmt.Println("Geçersiz işlem...")
			fmt.Println()

		}

	}
}
