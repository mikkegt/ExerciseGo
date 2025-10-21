package main

import "fmt"

func emptyMap() {
	var m map[string]int
	fmt.Println(m["errors"]) // 読み取り専用だから動く
}

func nilMap() {
	var m map[string]int // 初期化していないから
	m["errors"]++        // nilマップへの書き込みは、panic: assignment to entry in nil map
	fmt.Println(m["errors"])
}

func initializedMap() {
	m := make(map[string]int)
	m["errors"]++
	fmt.Println(m["errors"])
}

func wordCount() {
	m := make(map[string]int)
	words := []string{"hello", "world", "GoLang", "hello", "world", "gopher", "world", "GoLang"}
	for _, word := range words {
		m[word]++
	}
	fmt.Println("カウント結果:", m)
}

func main() {
	emptyMap()
	// nilMap()
	initializedMap()
	wordCount()
}

/**
 *  // 他の言語だとキーの存在チェックが必要
 *  if _, exists := m[word]; exists {
 *      m[word] = m[word] + 1
 *  } else {
 *      m[word] = 1
 *  }
 */
