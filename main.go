package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	unsei := omikujiExec()
	fmt.Fprint(w, "今日の運勢は"+unsei+"です")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func omikujiExec() string {
	var result string
	rand.Seed(time.Now().UnixNano())
	sai := rand.Intn(6) + 1
	switch sai {
	case 1:
		result = "凶"
	case 2:
		fallthrough
	case 3:
		result = "吉"
	case 4:
		fallthrough
	case 5:
		result = "中吉"
	default:
		result = "大吉"
	}
	return result
}
