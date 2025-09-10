// main paketi
package main

//giriş kütüphanesi
import (
	"fmt"
)

// person veri tipi belirtma
type person struct {
	name      string
	age       int
	dogumyili int
}

// fonksiyonun dışında da çalışabilecek değişken tanımlama
var packageName = "deneme package"

// yeni fonksiyon
func myMessage() {
	fmt.Println("bu bir fonksiyondur")
}

// yeni string fonksiyon
func aAdi(adi string) {
	fmt.Println("Selam", adi)

}

// yeni sayi kontrol fonksiyonu
func sayilar(SSayi1 int) {
	if SSayi1%2 == 0 {
		fmt.Println("sayi çift sayıdır")
	} else {
		fmt.Println("sayı tek sayıdır")
	}
}

// main fonksiyonunda verilen isim ve yaş bilgilerini ekrana yazdıran fonksiyon
func cumle(isim string, yas int) {
	fmt.Println("selam", isim, "yaşın", yas, "oldu")

}

// sonuç döndürmek için iki parametre alan ve bu parametrelerin toplamını döndüren fonksiyon
func islemler(c int, d int) int {
	return c + d
}

// fonksiyondan bir tane int bir tane string alır
func myFunction(e int, f string) (result1 int, txt string) {
	result1 = e * 3
	txt = f + " deneme fonksiyonudur"
	return
}

// struct ettiğimiz persone tipindeki veriyi ekrana yazdıran fonksiyon
func printperson2(per person) {
	fmt.Println("personel adı", per.name)
	fmt.Println("personel yaşı", per.age)
	fmt.Println("personel doğum yılı", per.dogumyili)
}

func bölme(bölünen, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan
}

func main() {
	adj := [2]string{"Hello", "World"}
	sehir := [3]string{"Istanbul", "Izmir", "Hatay"}
	arr1 := [4]int{1, 2, 3, 4}
	var blockvar int = 12
	const constvar = 10
	//sturct tipini kullanmak için personel tanımlama
	var person1 person
	var person2 person
	person1.name = "burcu"
	person1.age = 21
	person1.dogumyili = 2004
	person2.name = "seren"
	person2.age = 26
	person2.dogumyili = 1999
	var m, n int

	//kullanıcıdan giriş alır okur ve girişi ekrana yazar
	/* fmt.Print("lütfen birinci sayıyı giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println("girgiğiniz sayı:", value)

	fmt.Print("lütfen ikinci sayıyı giriniz:")
	reader2 := bufio.NewReader(os.Stdin)
	value2, _ := reader2.ReadString('\n')
	fmt.Print("girdiğiniz sayı:", value2)

	returnedValue := value + value2
	fmt.Print("girdiğiniz değerler toplamı:", returnedValue)*/

	var value int
	var value2 int

	fmt.Println("toplamak için iki sayı giriniz")
	fmt.Scan(&value, &value2)

	toplam2 := value + value2
	fmt.Printf("toplam= %d \n", toplam2)

	bölüm, kalan := bölme(20, 4)
	fmt.Printf("20 sayısının 4'e bölümü: bölüm = %d, kalan = %d\n", bölüm, kalan)

	isim := "defne"
	yas := 6
	fmt.Printf("isim = %s, yas = %d\n", isim, yas) //buradaki isim ve yas değişkenleri yukardaki fonksiyondan bağımsızdır.

	/* var moment time.time = time.Now()()
	fmt.Println("Şu anki zaman:", moment) çalışmadı */

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(sehir); j++ {
			fmt.Println(adj[i], sehir[j])
		}
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(sehir); j++ {
			fmt.Println(arr1[i], sehir[j])
		}
	}
	for m = 2; m <= 20; m++ {
		for n = 2; n < (m / n); n++ {
			if m%n == 0 {
				break
			}
		}
		if m%n != 0 {
			fmt.Println(m, "asal sayıdır")
		} else if m == 2 {
			fmt.Println(m, "asal sayıdır")
		} else {
			fmt.Println(m, "asal sayı değildir")
		}
	}

	if false {
		fmt.Println("bu bir if bloğudur")
	} else {
		fmt.Println("if bloğu çalışmadı")
	}

	if arr1[3] == 4 {
		fmt.Printf("arr1 dördüncü sayı = %d\n", arr1[3])
	} else {
		fmt.Println("arr1 dördüncü sayısı 4 değil")
	}

	fmt.Printf("packName = %s\n", packageName) //func main dışında tanımladığımız değişkeni yazar
	fmt.Printf("blockvar = %d \n", blockvar)   //bloktaki değişkeni yazar

	switch sehir[1] {
	case "Hatay":
		fmt.Println("hatay")
	default:
		fmt.Printf("hatay seçilmedi")
	}

	x := 11
	br := "brc"

	//xin veri tipini br nin string değerini yazar
	fmt.Printf("\nx =%T y= %s\n", x, br)

	// constvar = 20 // değerlendirilmez
	fmt.Printf("constvar =%d\n", constvar)

	//buradan
	myMessage()
	fmt.Println("fonksiyon çalıştı")

	aAdi("burcu")
	aAdi("seren")
	aAdi("duygu")

	sayilar(5)
	sayilar(10)
	sayilar(15)

	cumle("burcu", 21)
	cumle("seren", 27)
	cumle("duygu", 28)
	//buraya kadar fonksiyon değerleri belirtme

	/*func islemler(c int, d int) int {
	return c + d}*/
	result := islemler(2, 8)
	fmt.Println("sayılar 2 ve 8 toplama sonucu:", result)

	//sadece a değişkenimi yazar
	fmt.Println(myFunction(5, "bu bir"))
	a, _ := myFunction(10, "deneme")
	fmt.Println(a)

	//sadece b değişkeni yazar
	_, b := myFunction(10, "deneme")
	fmt.Println(b)

	//sturct tipi ve değişeni üstte tanımladığımız personel bilgilerini ekrana yazdırma
	fmt.Println("personel adı", person1.name)
	fmt.Println("personel yaşı", person1.age)
	fmt.Println("personel doğum yılı", person1.dogumyili)

	//person2 struct tipindeki veriyi ekrana yazdırma
	printperson2(person2)

	// First, define an outer function
	updateCounter := func() func() int {
		// define a local variable inside the function
		count := 100
		return func() int {
			count++
			return count
		}
	}

	// Now, creating a closure
	increment := updateCounter()

	// Using (calling) the closure
	fmt.Println(increment())
	fmt.Println(increment())
}
