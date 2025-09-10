package main

import "fmt"

// ===========================================
// 1. POINTER NEDİR?
// ===========================================

func basicPointerExample() {
	fmt.Println("=== 1. TEMEL POINTER KULLANIMI ===")

	// Normal değişken
	x := 42
	fmt.Printf("x değeri: %d\n", x)
	fmt.Printf("x'in bellek adresi: %p\n", &x) // & operatörü adres verir

	// Pointer tanımlama
	var p *int // int tipinde pointer
	p = &x     // x'in adresini p'ye ata

	fmt.Printf("p pointer'ının değeri (x'in adresi): %p\n", p)
	fmt.Printf("p'nin gösterdiği değer: %d\n", *p) // * operatörü değeri verir

	// Pointer ile değer değiştirme
	*p = 100
	fmt.Printf("Pointer ile değiştirdikten sonra x: %d\n", x)
	fmt.Println()
}

// ===========================================
// 2. ZERO VALUE VE NIL POINTER
// ===========================================

func nilPointerExample() {
	fmt.Println("=== 2. NIL POINTER ===")

	var p *int
	fmt.Printf("Başlangıçta p: %v\n", p) // nil

	// Nil pointer kontrolü
	if p == nil {
		fmt.Println("p nil pointer!")
	}

	// Nil pointer'a değer atamaya çalışmak PANIC yaratır!
	// *p = 42  // Bu satır çalıştırılırsa program çöker

	// Güvenli kullanım
	x := 42
	p = &x
	if p != nil {
		fmt.Printf("Artık p güvenli: %d\n", *p)
	}
	fmt.Println()
}

// ===========================================
// 3. POINTER VE FUNCTION
// ===========================================

// Value ile fonksiyon (kopya oluşur)
func changeByValue(x int) {
	x = 999
	fmt.Printf("Fonksiyon içinde x: %d\n", x)
}

// Pointer ile fonksiyon (orijinal değişken değişir)
func changeByPointer(x *int) {
	*x = 999
	fmt.Printf("Fonksiyon içinde *x: %d\n", *x)
}

func functionPointerExample() {
	fmt.Println("=== 3. FONKSİYONLARDA POINTER ===")

	a := 42
	b := 42

	fmt.Printf("Başlangıçta a: %d, b: %d\n", a, b)

	// Value ile çağırma
	changeByValue(a)
	fmt.Printf("changeByValue sonrası a: %d\n", a) // Değişmez!

	// Pointer ile çağırma
	changeByPointer(&b)
	fmt.Printf("changeByPointer sonrası b: %d\n", b) // Değişir!
	fmt.Println()
}

// ===========================================
// 4. STRUCT VE POINTER
// ===========================================

type Person struct {
	Name string
	Age  int
}

// Value receiver (kopya ile çalışır)
func (p Person) introduce() {
	fmt.Printf("Merhaba, ben %s ve %d yaşındayım\n", p.Name, p.Age)
}

// Pointer receiver (orijinal ile çalışır)
func (p *Person) haveBirthday() {
	p.Age++
	fmt.Printf("%s artık %d yaşında!\n", p.Name, p.Age)
}

// Pointer receiver ile değer güncelleme
func (p *Person) changeName(newName string) {
	p.Name = newName
}

func structPointerExample() {
	fmt.Println("=== 4. STRUCT VE POINTER ===")

	// Struct oluşturma
	person1 := Person{Name: "Ali", Age: 25}
	person1.introduce()

	// Pointer ile struct oluşturma
	person2 := &Person{Name: "Ayşe", Age: 30}
	person2.introduce() // Go otomatik olarak (*person2).introduce() yapar

	// Value receiver - orjinal değişmez
	person1.introduce()

	// Pointer receiver - orijinal değişir
	person1.haveBirthday()
	person1.introduce()

	// Pointer ile name değiştirme
	person2.changeName("Fatma")
	person2.introduce()
	fmt.Println()
}

// ===========================================
// 5. SLICE VE POINTER
// ===========================================

func modifySlice(s []int) {
	s[0] = 999 // Slice'lar reference type, değişir!
}

func modifySlicePointer(s *[]int) {
	*s = append(*s, 100) // Slice'ın kendisini değiştirme
}

func slicePointerExample() {
	fmt.Println("=== 5. SLICE VE POINTER ===")

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Başlangıçta slice1: %v\n", slice1)

	// Slice zaten reference type
	modifySlice(slice1)
	fmt.Printf("modifySlice sonrası: %v\n", slice1)

	// Slice pointer kullanımı
	modifySlicePointer(&slice1)
	fmt.Printf("modifySlicePointer sonrası: %v\n", slice1)
	fmt.Println()
}

// ===========================================
// 6. NEW VE MAKE FONKSİYONLARI
// ===========================================

func newMakeExample() {
	fmt.Println("=== 6. NEW VE MAKE ===")

	// new() fonksiyonu - pointer döndürür
	p1 := new(int)
	fmt.Printf("new(int) ile oluşturulan: %p, değer: %d\n", p1, *p1)
	*p1 = 42
	fmt.Printf("Değer atandıktan sonra: %d\n", *p1)

	// struct için new
	p2 := new(Person)
	fmt.Printf("new(Person): %+v\n", *p2) // Zero values
	p2.Name = "Mehmet"
	p2.Age = 35
	fmt.Printf("Değerler atandıktan sonra: %+v\n", *p2)

	// make() fonksiyonu - slice, map, channel için
	slice := make([]int, 5)
	fmt.Printf("make ile slice: %v\n", slice)

	mp := make(map[string]int)
	mp["anahtar"] = 42
	fmt.Printf("make ile map: %v\n", mp)
	fmt.Println()
}

// ===========================================
// 7. POINTER ARİTMETİĞİ (GO'DA YOK!)
// ===========================================

func pointerArithmeticExample() {
	fmt.Println("=== 7. POINTER ARİTMETİĞİ ===")
	fmt.Println("Go'da C/C++ gibi pointer aritmetiği YOKTUR!")
	fmt.Println("p++, p--, p+1 gibi işlemler YASAKTIR!")
	fmt.Println("Bu Go'nun güvenlik özelliğidir.")

	// Bu kod hata verir:
	// arr := [5]int{1, 2, 3, 4, 5}
	// p := &arr[0]
	// p++  // COMPILE ERROR!
	fmt.Println()
}

// ===========================================
// 8. PRATİK POINTER KULLANIM ÖRNEKLERİ
// ===========================================

// Linked List node örneği
type Node struct {
	Data int
	Next *Node
}

// Binary Tree node örneği
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func practicalExamples() {
	fmt.Println("=== 8. PRATİK ÖRNEKLER ===")

	// Linked List
	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 3}

	node1.Next = node2
	node2.Next = node3

	fmt.Println("Linked List:")
	current := node1
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")

	// Binary Tree
	root := &TreeNode{
		Value: 10,
		Left:  &TreeNode{Value: 5},
		Right: &TreeNode{Value: 15},
	}

	fmt.Printf("Tree root: %d\n", root.Value)
	fmt.Printf("Left child: %d\n", root.Left.Value)
	fmt.Printf("Right child: %d\n", root.Right.Value)
	fmt.Println()
}

// ===========================================
// 9. POINTER İLE İLGİLİ İPUÇLARI
// ===========================================

func tips() {
	fmt.Println("=== 9. POINTER İPUÇLARI ===")
	fmt.Println("✅ Pointer ne zaman kullanılır:")
	fmt.Println("  - Büyük struct'ları kopyalamaktan kaçınmak için")
	fmt.Println("  - Fonksiyonda orijinal değeri değiştirmek için")
	fmt.Println("  - Bellek tasarrufu için")
	fmt.Println("  - Linked list, tree gibi veri yapılarında")
	fmt.Println()

	fmt.Println("❌ Pointer ne zaman kullanılMAZ:")
	fmt.Println("  - Küçük değerler için (int, bool vs.)")
	fmt.Println("  - Sadece okuma işlemleri için")
	fmt.Println("  - Slice, map, channel zaten reference type")
	fmt.Println()

	fmt.Println("🔍 Performans karşılaştırması:")

	// Büyük struct örneği
	type BigStruct struct {
		data [1000]int
	}

	big := BigStruct{}

	// Value ile (kopyalama maliyeti yüksek)
	_ = big

	// Pointer ile (sadece adres kopyalanır)
	bigPtr := &big
	_ = bigPtr

	fmt.Println("  - Büyük struct: Pointer kullan!")
	fmt.Println("  - Küçük değerler: Value kullan!")
}

// ===========================================
// MAIN FONKSİYON - TÜM ÖRNEKLERİ ÇALIŞTIR
// ===========================================

func main() {
	fmt.Println("🚀 GO POINTER REHBERİ 🚀")
	fmt.Println("========================")
	fmt.Println()

	basicPointerExample()
	nilPointerExample()
	functionPointerExample()
	structPointerExample()
	slicePointerExample()
	newMakeExample()
	pointerArithmeticExample()
	practicalExamples()
	tips()

	fmt.Println("🎉 Pointer rehberi tamamlandı!")
}
