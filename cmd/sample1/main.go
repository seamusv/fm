package main

import (
	"github.com/seamusv/fm-integration/screendata"
	"log"
	"os"
	"time"
)

type AddLine struct {
	AccountCode  string    `fm:"GLACCT"`
	MatchCode    int       `fm:"MATCHID"`
	RequiredDate time.Time `fm:"SCHDATE"`
}

func main() {
	line := AddLine{
		AccountCode:  "551-101001-0201",
		MatchCode:    1,
		RequiredDate: time.Now().Add(time.Hour * 24 * 7),
	}

	b, err := screendata.Marshal("ADD", line)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
}

// OUTPUT:
// <trans gui="N">
//  <command cmd="ADD" proc="FM">
//    <screendata>
//      <put-fields>
//        <f n="GLACCT" v="551-101001-0201"></f>
//        <f n="MATCHID" v="1"></f>
//        <f n="SCHDATE" v="2020/12/05"></f>
//      </put-fields>
//    </screendata>
//  </command>
//</trans>