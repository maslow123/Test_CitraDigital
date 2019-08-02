package main

import (
	"log"
	"net/http"
	"encoding/json"
	"strings"
	"sort"
	"github.com/gorilla/mux"
)
/*=================================================================================*/

// Model for sort
type Sort struct {
	Word string `json: "SortWord"`
}

func createWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var sortWords Sort
	// _ = json.NewDecoder(r.Body).Decode(&sortWords)
	// sorts = append(sorts, sortWords) // {{ inputan dari body }]
	
	// json.NewEncoder(w).Encode(sortWords)

	input := r.FormValue("Word")
	if input == "" {
		log.Printf("Please input cant be empty")
		w.WriteHeader(500)
		return
	}
	inputResult := sortWords(strings.ToLower(input))

	result := Sort{ inputResult }
	json.NewEncoder(w).Encode(result)
}

/*====================================================================*/
func sortWords(w string)(result string) {
	var arrVowel, arrConsonant []string
	
	param := strings.Replace(w," ","",-1)
	// Pecahkan nilai string parameter menjadi array
	arrWord := strings.Split(param,"")

	vowels := [5]string {"a","i","u","e","o"}
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

	result = strings.Join(arrVowel,"") + strings.Join(arrConsonant,"")
	return result
	
}
/*=================================================================================*/

func main() {
	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/sortWord",createWords).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000",router))

}