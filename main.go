package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)
	fmt.Printf("現在日時 (JST): %s\n", now.Format("2006-01-02 15:04:05"))

	// 1-10 のランダムな数値を出力
	n := rand.Intn(10) + 1
	fmt.Printf("ランダム数値 (1-10): %d\n", n)
}
