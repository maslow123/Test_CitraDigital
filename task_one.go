package main
import (
	"fmt" 
	"sort" 
	"strings"
)

func main() {
	sortWord(strings.ToLower("Aman"))
}

func sortWord(word string) {
	var arrVowel, arrConsonant []string
	
	param := strings.Replace(word," ","",-1)
	// Pecahkan nilai string parameter menjadi array
	arrWord := strings.Split(param,"")

	vowels := [5]string {"a","i","u","e","o"}// [ a i u e o ]
	// Lakukan perulangan untuk mendapatkan nilai setiap index arrWord.
	for _,value := range arrWord {
		// Lakukan perulangan lagi untuk mendapatkan nilai setiap index dari vokal
		for index,v := range vowels {
			// Cek jika ada value yang sama dengan v, maka masukkan  value kedalam array arrVowel untuk mendapatkan huruf hidup
			if value == v {
				arrVowel = append(arrVowel, value)
				break
			// Lanjut kekondisi kedua untuk mendapatkan huruf mati yang ada pada nilai parameter
			} else if index == len(vowels) - 1{
				arrConsonant = append(arrConsonant, value)
				// fmt.Println(arrConsonant)
			}
		}
	}
	sort.Strings(arrVowel) 
	sort.Strings(arrConsonant)

	result := strings.Join(arrVowel,"") + strings.Join(arrConsonant,"")
	
	fmt.Printf(result)
}