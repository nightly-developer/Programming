package main

import ("fmt")

func mergeAlternately(word1 string, word2 string) {

  // var length uint = uint(min(len(word1),len(word2)));
	// var appendString uint = uint(uint(max(len(word1),len(word2))) - length);
	// var byteSlice = make([]byte,len(word1)+len(word2))
	
	word := word1 & word2
	fmt.Println(word)
	// for _,_ := range(word) {
	// 	// fmt.Println(string(char))
	// }
	result := '9' | '7'
	fmt.Println(string(result))
}

func main(){
	mergeAlternately("hell9870","w")	
	
}