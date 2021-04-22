package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZWBgaGPo8TydselPrO55kuZl/+dg8jUCz6blzNLwspn1WWM7cWcU0s5ASK3bTf0PRbyt5t8Q99AwXf1ESQurpEuXOGHwQdtvntUlNo2sTPAhRzS9d5dv7cfrduhR/iFh8vKYgAxbYxn2A4INTACKIZJBkYJjCeYHig1wAkIXQEE4JmYGhgAcmDaJA8iGYA0g08QCYDRD/IHBBfAUSrMjCDaLg61QZGsDoGkDwDSB5Mg/SB+TxQPhBEMEHkQTRIHszngfIZQO6EyINoiLsh8mA+xDVgebBtklA+D8x2CADJM0DlGaDyCAAAbBwSwYgBAAA="

	slab, err := slabdecoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabBase64, err := slabencoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
