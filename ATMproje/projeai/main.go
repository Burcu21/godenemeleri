package main

import "fmt"

// Hesap bilgilerini tutan struct
type Account struct {
	balance int // Hesap bakiyesi
}

// Hesaptan para çekme fonksiyonu (pointer receiver)
func (a *Account) withdraw(amount int) bool {
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("%.d TL çekildi. Kalan bakiye: %d TL\n", amount, a.balance)
		return true
	} else {
		fmt.Println("Yetersiz bakiye!")
		return false
	}
}

// Hesaba para yatırma fonksiyonu (pointer receiver)
func (a *Account) deposit(amount int) {
	a.balance += amount
	fmt.Printf("%d TL yatırıldı. Güncel bakiye: %d TL\n", amount, a.balance)
}

// Bakiye sorgulama fonksiyonu (value receiver - sadece okuma)
func (a Account) getBalance() int {
	return a.balance
}

func main() {
	// Hesap oluştur (başlangıç bakiyesi 1000 TL)
	account := &Account{balance: 1000}

	var enteredPassword string
	var choice string
	password := "123"
	maxAttempts := 3

	fmt.Println("=== ATM'ye Hoşgeldiniz ===")

	// Şifre kontrolü
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Print("Lütfen şifrenizi giriniz: ")
		fmt.Scan(&enteredPassword)

		if password == enteredPassword {
			fmt.Println("Giriş başarılı!")

			// Ana menü döngüsü
			for {
				fmt.Println("\n=== ATM İşlemleri ===")
				fmt.Println("1. Hesaptan para çekme")
				fmt.Println("2. Hesaba para yatırma")
				fmt.Println("3. Hesap bakiyesi öğrenme")
				fmt.Println("4. Çıkış")
				fmt.Print("Yapmak istediğiniz işlemi seçiniz: ")
				fmt.Scan(&choice)

				switch choice {
				case "1":
					fmt.Println("\n--- Para Çekme ---")
					var amount int
					fmt.Print("Çekmek istediğiniz tutarı giriniz: ")
					fmt.Scan(&amount)

					if amount <= 0 {
						fmt.Println("Geçersiz tutar!")
						continue
					}

					// Pointer kullanarak para çekme
					account.withdraw(amount)

				case "2":
					fmt.Println("\n--- Para Yatırma ---")
					var amount int
					fmt.Print("Yatırmak istediğiniz tutarı giriniz: ")
					fmt.Scan(&amount)

					if amount <= 0 {
						fmt.Println("Geçersiz tutar!")
						continue
					}

					// Pointer kullanarak para yatırma
					account.deposit(amount)

				case "3":
					fmt.Println("\n--- Bakiye Sorgulama ---")
					// Value receiver kullanarak bakiye öğrenme
					fmt.Printf("Güncel bakiyeniz: %d TL\n", account.getBalance())

				case "4":
					fmt.Println("Çıkış yapılıyor... İyi günler!")
					return

				default:
					fmt.Println("Geçersiz seçim! Lütfen 1-4 arası bir değer giriniz.")
				}
			}
		} else {
			remainingAttempts := maxAttempts - attempt
			if remainingAttempts > 0 {
				fmt.Printf("Hatalı şifre! Kalan deneme hakkınız: %d\n", remainingAttempts)
			} else {
				fmt.Println("Kartınız bloke edildi! Lütfen bankanızla iletişime geçiniz.")
				return
			}
		}
	}
}
