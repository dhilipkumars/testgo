package main

import (
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("VCAP_APP_PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	list_num := strings.Split(req.URL.Path, "/")
	for _, tstr := range list_num {
		if tstr != "/" {
			fval, err := strconv.Atoi(tstr)
			if err == nil {
				fmt.Fprintln(res, PFact(fval, 2))
			}
		}
	}
}

func FactRange(st, en int, res chan *big.Int) {

	//fmt.Printf("Fact Range st=%d en=%d\n", st, en)
	fact := big.NewInt(1)
	i1 := big.NewInt(int64(st))
	i2 := big.NewInt(1)
	for i := st; i <= en; i = i + 1 {

		fact.Mul(fact, i1)
		i1.Add(i1, i2)

	}
	//fmt.Printf("res=%v\n", fact)
	if res != nil {
		res <- fact
	}
	//fmt.Printf("Fact Range st=%d en=%d\n", st, en)
}

func PFact(x, agent int) *big.Int {
	if x <= 0 {
		return big.NewInt(1) 
	}
	if agent <= 4 {
		agent = 4
	}
	rch := make(chan *big.Int)
	quad := x / agent
	st := 1
	en := quad
	for i := 0; i < (agent - 1); i++ {
		go FactRange(st, en, rch)
		st = en + 1
		en = en + quad
	}
	go FactRange(st, x, rch)
	fact := big.NewInt(1)
	for i := 0; i < agent; i++ {

		fact.Mul(fact, <-rch)

	}
	return fact

}
