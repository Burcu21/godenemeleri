package main

import "fmt"

type Count struct {
	hesap int
}

func (a *Count) withdraw(amount int) bool {
	if a.hesap >= amount {
		a.hesap -= amount
		fmt.Printf("%.d TL çekildi. Kalan bakide: %d Tl \n", amount, a.hesap)
		return true
	} else {
		fmt.Println("Yetersiz bakiye")
		return false
	}
}
func (a *Count) deposit(amount int) {
	a.hesap += amount
	fmt.Printf("%d TL yatırlıdı. Güncel bakiye: %d TL\n", amount, a.hesap)
}
func (a *Count) getBalance() int {
	return a.hesap
}

func main() {
	account := &Count{hesap: 1000}

	var epassword string
	var process string
	var amount int
	password := "123"
	var count int = 3
	// pointerla dene
	fmt.Println("3 kere denme hakknız var!")

	for i := 1; i <= count; i++ {
		fmt.Println("Lütfen Şifre Giriniz :")

		fmt.Scan(&epassword)

		if password == epassword {
			for {
				fmt.Println("Yapmak istediğiniz işlemi giriniz")
				fmt.Println("1. Hesaptan para çekme ")
				fmt.Println("2. Hesaba para yatırma ")
				fmt.Println("3. Hesap bakiyesi öğrenme ")
				fmt.Println("4. Çıkış")
				fmt.Scan(&process)

				switch process {
				case "1":
					fmt.Println("Para çekme ekranı")
					fmt.Println("Hesaptan ne kadar para çekmek istiyorsunuz?:")
					fmt.Scan(&amount)

					if amount <= 0 {
						fmt.Println("Geçersiz tutar")
						continue
					}
					account.withdraw(amount)
				case "2":
					fmt.Println("Para yatırma ekranı")
					fmt.Println("Hesaba ne kadar para yatırmak istiyorsunuz?:")
					fmt.Scan(&amount)

					if amount <= 0 {
						fmt.Println("Geçersiz tutar")
						continue
					}
					account.deposit(amount)

				case "3":
					fmt.Println("Hesap bakiyesini öğrenme ekranı")
					fmt.Printf("Güncel bakiyeniz: %d TL \n", account.getBalance())

				case "4":
					fmt.Println("Çıkış Yapılıyor...")
					return
				default:
					fmt.Println("Böyle Bir İşlem Yok")
				}

			}
		} else {
			if i == 3 {
				fmt.Println("Kartınız Bloke Edildi...")
				break
			}

		}

	}
}
