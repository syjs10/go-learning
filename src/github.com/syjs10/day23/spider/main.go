package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func getData() {
	rand.Seed(time.Now().UnixNano())
	u := strconv.Itoa(rand.Int() / 100000000000)
	rand.Seed(time.Now().UnixNano())
	p := strconv.Itoa(rand.Int() / 100000000000)
	resp, err := http.PostForm("http://lol-xw.rrkiso.cn/rufen.php", url.Values{"u": {u}, "p": {p}, "bianhao": {"1"}})
	fmt.Println(resp, err)
}

func main() {
	for {
		for i := 0; i < 10; i++ {
			go getData()
		}
		time.Sleep(time.Microsecond * 1000)
	}
}
