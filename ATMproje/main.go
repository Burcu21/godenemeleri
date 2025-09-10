package main

import (
	"fmt"
)

type users struct {
	name     string
	id       int
	password int
	count    int
}
type counti interface {
	topla() int
	cikar() int
}
type hesap struct {
	t int
	c int
}

func (e hesap) topla() int {
	return e.t

}
func (r hesap) cikar() int {
	return r.c
}


func main() {
	var user1 users

	user1.name = "Burcu"
	user1.id = 123456
	user1.password = 1234
	user1.count = 1000

	var giris int
	var sifre int
	var islem int
	var ekle int
	var cek int

	fmt.Println("Hosgeldiniz ")
	fmt.Println("Lütfen ID numaranzı yazınız")
	fmt.Scan(&giris)
	fmt.Println("Lütfen şifrenizi giriniz")
	fmt.Scan(&sifre)

	if giris == user1.id && sifre == user1.password {
		fmt.Println("Hoşgeldiniz ", user1.name)
		fmt.Println("Hesap bilgileriniz")
		fmt.Printf("%v\n", user1)

		fmt.Println("Yapmak istediğiniz işlemi giriniz")
		fmt.Println("1. Hesap bakiyesi öğrenme")
		fmt.Println("2. Hesaba para yatırma")
		fmt.Println("3. Hesaptan para çekme")
		fmt.Println("4. Çıkış")

		fmt.Scan(&islem)

		for i := 0; i < 10; i++ {
			if islem == 1 {
				fmt.Printf("Hesap bakiyeniz: %d", user1.count)
				break

			} else if islem == 2 {
				fmt.Println("Yatırmak istediğiniz bakiyeyi giriniz: ")
				fmt.Scan(&ekle)
				fmt.Println("Hesabınıza para yatırılmıştır")
				break
			} else if islem == 3 {
				fmt.Println("Çekmek istediğiniz bakiyeyi giriniz")
				fmt.Scan(&cek)
				fmt.Println("Hesabınıza para çekilmiştir")
				break
			} else if islem == 4 {
				fmt.Println("Hesabınızdan çıkış yapılıyor...")
				break
			} else {
				fmt.Println("Yanlış giriş yaptınız, tekrar deneyiniz")
				break
			}
		}

	} else {
		fmt.Println("Yanlış ID veya şifre girdiniz")

		fmt.Scan(&giris)
		for attempts := 0; attempts < 3; attempts++ {
			if giris != user1.id || sifre != user1.password {
				fmt.Println(3-attempts, "hakkınız kaldı")
				fmt.Print("lütfen şifrenizi tekrar giriniz: ")
			} else {
				fmt.Println("Şifre ya da ID doğru")
			}
		}
	}
}
