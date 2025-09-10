package main

func main() {
	//var text string = "İstanbul'un kalabalık sokaklarında yürürken, şehrin benzersiz ritmini hissediyorum."

	//fmt.Println(strings.ReplaceAll(text, "kalabalık", "gürültülü"))
	//Replace

	//fmt.Printf("%s\n", strings.Split(text, ","))
	//Split

	//fmt.Println(strings.ToUpper(text))
	//ToUpper

	//fmt.Println(strings.Trim(text, "İ"))

	/* replacer := strings.NewReplacer("'", "-", ",", ";")
	text = replacer.Replace(text)
	fmt.Println(text) */

	//fmt.Println("ba" + strings.Repeat("na", 2))

	/* s := []string{"aa", "bb", "cc"}
	fmt.Println(strings.Join(s, " ")) */

	//fmt.Println(strings.HasSuffix("kalabalık", "k"))//true

	//fmt.Printf("Fields are: %q", strings.Fields("  aa bb  cc   "))

	/* fmt.Println(strings.EqualFold("Go", "go"))//true
	fmt.Println(strings.EqualFold("AB", "ab"))//true
	fmt.Println(strings.EqualFold("ß", "ss"))//false */

	/* fmt.Println(strings.Count("peynir", "e")) // 1
	fmt.Println(strings.Count("beş", ""))     // 4 */

	/* fmt.Println(strings.Contains("deniz", "en")) //true
	fmt.Println(strings.Contains("deniz", "b"))  //false
	fmt.Println(strings.Contains("deniz", ""))   //true
	fmt.Println(strings.Contains("", ""))        //true */

	/* fmt.Println(strings.Compare("b", "d")) //-1
	fmt.Println(strings.Compare("a", "a")) //0
	fmt.Println(strings.Compare("g", "a")) //1 */

}
