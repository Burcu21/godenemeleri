package product

import "fmt"

type Product struct {
	Cap   int
	Name  string
	Price int
}

var num int

func Urun1(b *Product) (num int) {
	var Product1 = Product{
		Cap:   100,
		Name:  "Elbise",
		Price: 1000,
	}

	if Product1.Cap >= num {
		dress := num * Product1.Price
		Product1.Cap -= num
		fmt.Println(Product1.Name, "Toplam Tutarı:  ", dress, "\nKalan Stok: ", Product1.Cap)
		return Product1.Cap
	} else {
		fmt.Println("Stokta o kadar ürün yoktur")
		return
	}
}

func Urun2(Urun22 Product) {
	Urun22.Cap = 40
	Urun22.Name = "Çanta"
	Urun22.Price = 300

	fmt.Println("Çantadan kaç tane almak istiyorsunuz")
	fmt.Scan(&num)

	if Urun22.Cap >= num {
		purse := num * Urun22.Price
		Urun22.Cap -= num
		fmt.Println(Urun22.Name, "Toplam Tutarı:  ", purse, "\nKalan Stok: ", Urun22.Cap)
	} else {
		fmt.Println("Stokta o kadar ürün yoktur")
	}
}
func Urun3(Urun33 Product) {
	Urun33.Cap = 80
	Urun33.Name = "Kazak"
	Urun33.Price = 500

	fmt.Println("Kazaktan kaç tane almak istiyorsunuz")
	fmt.Scan(&num)

	if Urun33.Cap >= num {
		sweater := num * Urun33.Price
		Urun33.Cap -= num
		fmt.Println(Urun33.Name, "Toplam Tutarı:  ", sweater, "\nKalan Stok: ", Urun33.Cap)
	} else {
		fmt.Println("Stokta o kadar ürün yoktur")
	}
}
func Urun4(Urun44 Product) {
	Urun44.Cap = 40
	Urun44.Name = "Ayakkabı"
	Urun44.Price = 2000

	fmt.Println("Ayakkabıdan kaç tane almak istiyorsunuz")
	fmt.Scan(&num)

	if Urun44.Cap >= num {
		shoe := num * Urun44.Price
		Urun44.Cap -= num
		fmt.Println(Urun44.Name, "Toplam Tutarı:  ", shoe, "\nKalan Stok: ", Urun44.Cap)
	} else {
		fmt.Println("Stokta o kadar ürün yoktur")
	}
}
func Urun5(Urun55 Product) {
	Urun55.Cap = 90
	Urun55.Name = "Takı"
	Urun55.Price = 750

	fmt.Println("Takıdan kaç tane almak istiyorsunuz")
	fmt.Scan(&num)

	if Urun55.Cap >= num {
		jewelry := num * Urun55.Price
		Urun55.Cap -= num
		fmt.Println(Urun55.Name, "Toplam Tutarı:  ", jewelry, "\nKalan Stok: ", Urun55.Cap)
	} else {
		fmt.Println("Stokta o kadar ürün yoktur")
	}
}
func Order() {
	fmt.Println("")
}
