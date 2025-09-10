package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var tpl *template.Template

func main() {
	var err error
	connStr := "host=localhost user=postgres password=postgres dbname=chatdeneme sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız: %v", err)
	}
	fmt.Println("Veritabanı bağlantısı başarılı")

	tpl = template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/submit", registerHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server 8080 portunda çalışıyor...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Ana sayfa
func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Ana sayfa istendi: %s", r.URL.Path)
	tpl.Execute(w, nil)
}

// Giriş yapma işlemi
func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Login handler çağrıldı: Method=%s", r.Method)

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	log.Printf("Login veriler: email=%s, password_length=%d", email, len(password))

	// Eğer Postman'dan gelmişse JSON döndür
	if isFromPostman(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"email":    email,
			"password": password,
		})
		return
	}

	if email == "" || password == "" {
		renderMessage(w, "E-posta ve şifre gerekli", "error")
		return
	}

	var hashedPassword, name string
	err := db.QueryRow("SELECT password, name FROM users WHERE email=$1", email).Scan(&hashedPassword, &name)
	if err != nil {
		renderMessage(w, "E-posta veya şifre hatalı", "error")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		renderMessage(w, "E-posta veya şifre hatalı", "error")
		return
	}

	data := map[string]interface{}{
		"Message":     fmt.Sprintf("Hoş geldiniz, %s! Başarıyla giriş yaptınız.", name),
		"MessageType": "success",
		"UserName":    name,
		"UserEmail":   email,
		"IsLoggedIn":  true,
	}
	tpl.Execute(w, data)
}

// Kayıt olma işlemi
func registerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Register handler çağrıldı: Method=%s", r.Method)

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	log.Printf("Register veriler: name=%s, email=%s, password_length=%d", name, email, len(password))

	// Eğer Postman'dan gelmişse JSON döndür
	if isFromPostman(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"name":     name,
			"email":    email,
			"password": password,
		})
		return
	}

	if name == "" || email == "" || password == "" {
		renderMessage(w, "Tüm alanları doldurun", "error")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		renderMessage(w, "Şifre işlenemedi", "error")
		return
	}

	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", name, email, string(hashedPassword))
	if err != nil {
		renderMessage(w, "E-posta zaten kayıtlı veya veritabanı hatası", "error")
		return
	}

	renderMessage(w, "Kayıt başarılı", "success")
}

// Mesaj gösterme
func renderMessage(w http.ResponseWriter, message string, messageType string) {
	data := map[string]interface{}{
		"Message":     message,
		"MessageType": messageType,
	}
	tpl.Execute(w, data)
}

// Postman isteğini tespit eden yardımcı fonksiyon
func isFromPostman(r *http.Request) bool {
	ua := strings.ToLower(r.Header.Get("User-Agent"))
	if strings.Contains(ua, "postman") || r.Header.Get("Postman-Token") != "" {
		return true
	}
	return false
}
