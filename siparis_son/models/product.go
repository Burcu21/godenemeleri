package models

import "fmt"

type Product struct {
	ID    string
	Name  string
	Stok  int
	Price int
	Cap   int
}

var Urun = &[...]Product{
	{ID: "1", Name: "Elbise", Stok: 100, Price: 1000, Cap: 100},
	{ID: "2", Name: "Çanta", Stok: 40, Price: 300, Cap: 40},
	{ID: "3", Name: "Kazak", Stok: 80, Price: 500, Cap: 80},
	{ID: "4", Name: "Ayakkabı", Stok: 20, Price: 2000, Cap: 20},
	{ID: "5", Name: "Takı", Stok: 90, Price: 750, Cap: 90},
}

func Urunler() {

	fmt.Println("Ürünlerimiz:")
	for _, i := range Urun {
		fmt.Println(i.ID, "--Ürün Adı:--", i.Name, "--Ürün Fiyatı:--", i.Price, "--Ürün Stoğu:--", i.Cap)
	}
}

func SepetEkle() (number string, add int) {
	fmt.Println("İstediğiniz ürünün ID numarasını giriniz(1-5)")
	fmt.Scan(&number)

	var index = -1

	for i, urun := range Urun {
		if urun.ID == number {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Geçersiz seçim")
		return
	}

	fmt.Println("Üründen ne kadar istediğinizi yazınız.")
	fmt.Scan(&add)

	if Urun[index].Cap < add {
		fmt.Println("Yetersiz stok")
		return
	}
	Urun[index].Cap -= add
	fmt.Println(add, "Ürün sepete eklendi.")

	return number, add
}

func SepetCikar() (number string, dis int) {
	Sepet()
	fmt.Println("İstediğiniz ürünün ID numarasını giriniz(1-5)")
	fmt.Scan(&number)

	var index = -1

	for i, urun := range Urun {
		if urun.ID == number {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Geçersiz seçim")
		return
	}

	fmt.Println("Üründen ne kadar çıkarmak istediğinizi yazınız.")
	fmt.Scan(&dis)

	if Urun[index].Stok-Urun[index].Cap < dis {
		fmt.Println("Sepette bu kadar ürün yok")
		return
	}
	Urun[index].Cap += dis
	fmt.Println(dis, "Ürün sepeten çıkarıldı.")

	return number, dis
}

func Sepet() {
	var toplam int
	for _, i := range Urun {
		if i.Stok-i.Cap == 0 {
			continue
		}
		basket := i.Stok - i.Cap
		price := basket * i.Price
		fmt.Println("SEPETİNİZ")
		fmt.Println(i.ID, i.Name, "--", "Adet:", basket, "Tutar:", price, "\n------------")
		fmt.Println()

		toplam += (i.Stok - i.Cap) * i.Price
		fmt.Println("Sepet Toplamı:", toplam)

	}

}
