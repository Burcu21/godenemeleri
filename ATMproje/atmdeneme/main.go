package main

import (
	"fmt"
	"time"
)

type Account struct {
	count int
}

func (a *Account) paracekme(amount int) bool {
	if a.count >= amount {
		a.count -= amount
		fmt.Printf("%d TL para çekildi.\n", amount)
		return true
	} else {
		fmt.Println("Yetersiz bakiye!")
		return false
	}
}
func (a *Account) parayatirma(amount int) {
	a.count += amount
	fmt.Printf("%d TL para yatırıldı.\n", amount)

}
func (a *Account) bakiye() int {
	return a.count
}

func main() {
	var password string = "123"
	var epassword string
	var process string
	var amount int
	account := &Account{count: 1000}

	fmt.Println("HOŞGELDİNİZ")
	fmt.Println("NOT: 3 kere şifre girme hakkınız var!")

	for i := 1; i <= 3; i++ {
		fmt.Println("Lütfen şifreyi giriniz")
		fmt.Scan(&epassword)
		time.Sleep(1 * time.Second)

		if password == epassword {
			for {
				fmt.Println("Yapmak istediğiniz işlemi giriniz")
				fmt.Println("1. Hesaptan para çekme ")
				fmt.Println("2. Hesaba para yatırma ")
				fmt.Println("3. Hesap bakiyesi öğrenme ")
				fmt.Println("4. Çıkış")
				fmt.Scan(&process)
				fmt.Println("Yönlendiriliyor...")
				time.Sleep(1 * time.Second)
				fmt.Println()

				switch process {
				case "1":
					fmt.Println("Hesaptan ne kadar çekmek istiyorsunuz?")
					fmt.Scan(&amount)
					fmt.Println()

					if amount <= 0 {
						fmt.Println("Geçersiz tutar")
						continue
					}
					account.paracekme(amount)
					fmt.Println()

					time.Sleep(1 * time.Second)

				case "2":
					fmt.Println("Hesaptan ne kadar yatırmak istiyorsunuz?")
					fmt.Scan(&amount)
					fmt.Println()

					if amount <= 0 {
						fmt.Println("Geçersiz tutar")
						continue
					}
					account.parayatirma(amount)
					fmt.Println()

					time.Sleep(1 * time.Second)

				case "3":
					fmt.Printf("Hesap bakiyesi: %d\n", account.bakiye())
					time.Sleep(1 * time.Second)
					fmt.Println()

				case "4":
					fmt.Println("Çıkış yapılıyor...")
					time.Sleep(1 * time.Second)
					fmt.Println()

					return
				default:
					fmt.Println("Geçersiz seçim!")
					fmt.Println()

					fmt.Println("Lütfen 1-4 arasında iseçim yapınız")
					time.Sleep(1 * time.Second)

				}

			}

		} else {
			fmt.Println("Yanlış şifre girdiniz!")
			fmt.Println()

			if i == 3 {
				fmt.Println("Karınız bloke edildi...")
				break
			}
		}
	}

}
