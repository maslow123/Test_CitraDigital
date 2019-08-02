package main

import(
	"fmt"
	"strings"
)

func main() {
	vowelAndConsonant(strings.ToLower("omama"))
}

func vowelAndConsonant(w string) {
	var arrVowel, arrConsonant []string
	param := strings.Replace(w," ","",-1)

	// Pecahkan nilai string parameter menjadi array
	arrWord := strings.Split(param,"")
	
	vowel := [5]string {"a","i","u","e","o"}
	for _, value := range arrWord {
		// fmt.Println(value)
		for index,v := range vowel {
			// cek apakah value ada yang sama dengan v atau tidak, jika ada maka lakukan sortir dan kekondisi selanjutnya
			if value == v {
				// cek lagi jika ada index array yang tidak nol, maka lakukan perulangan lagi untuk mendapatkan huruf hidup
				if len(arrVowel) != 0 {
					for u,y := range arrVowel {
						// cek apakah ada huruf hidup yang sama, jika ada maka kondisi akan selesai dan berlanjut ke kondisi berikutnya
						if value == y {
							break
						} else if u == len(arrVowel) - 1 {
							// Jika u sama dengan panjang arrVowel, maka masukkan nilai array yang ada pada variable value kedalam variable arrVowel
							arrVowel = append(arrVowel, value) // [ o , a ]
						}
					}
				// jika panjang arrVowel nol, maka langsung masukkan nilai array yang ada pada variable value kedalam variable arrVowel
				}else {
					arrVowel = append(arrVowel, value)
				}
				break
			// ambil banyaknya jumlah huruf mati yang ada pada nilai parameter
			}else if index == len(vowel) -1 {
				arrConsonant = append(arrConsonant, value)
			}
		}
	}
	
	fmt.Printf("Huruf mati : %d\n",len(arrConsonant))
	fmt.Printf("Huruf hidup: %d",len(arrVowel))
}