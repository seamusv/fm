package main

import (
	"fmt"
	"github.com/seamusv/fm-integration/fm"
	"log"
	"os"
	"time"
)

type (
	AddLine struct {
		AccountCode  *string    `fm:"GLACCT"`
		MatchCode    *int       `fm:"MATCHID"`
		RequiredDate *time.Time `fm:"SCHDATE"`
	}

	PO501 struct {
		AString *string    `fm:"IDBUYR"`
		AInt64  *int64     `fm:"CNTREL"`
		ATime   *time.Time `fm:"CODESTTS"`
		ABool   *bool      `fm:"STTSPRT"`
	}
)

func main() {
	line := AddLine{
		AccountCode:  fm.String("551-101001-0201"),
		MatchCode:    fm.Int(1),
		RequiredDate: fm.Time(time.Now().Add(time.Hour * 24 * 7)),
	}

	b, err := fm.Marshal("ADD", line)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
	fmt.Print("\n\n-------------\n\n")

	res, err := fm.Parse([]byte(`<trans ok="Y"><screendata><return-fields><f n="STTSPRT" v="Y"></f><f n="IDBUYR" v="FMISPR"></f><f n="CODESTTS" v="2020/12/31"></f><f n="SWNOTEH" v="N"></f><f n="TEXTFRT" v=""></f><f n="CNTREL" v="6942"></f><f n="SWNOTEHN" v="N"></f><f n="LINEWHSE" v=""></f><f n="NAMEVEND" v=""></f><f n="TEXTOP01" v="TENDER TYPE"></f><f n="IDOP03" v=""></f><f n="IDOP02" v=""></f><f n="TAX02" v=""></f><f n="SPCLINST" v=""></f><f n="LINECARR" v=""></f><f n="TEXTOP03" v="SOA #"></f><f n="TEXTCARR" v=""></f><f n="CNTREV" v="0"></f><f n="LINEFOB" v=""></f><f n="LINESHPT" v=""></f><f n="TAX04" v=""></f><f n="IDLOCN" v="0"></f><f n="LINEBILL" v=""></f><f n="TAXENRP" v=""></f><f n="LINESHTY" v="1"></f><f n="IDORDR" v=""></f><f n="LASTPRT" v=""></f><f n="TAX01" v=""></f><f n="LINESCHD" v=""></f><f n="TAX03" v=""></f><f n="IDCURN" v=""></f><f n="NETORDR" v="0.00 "></f><f n="LINEMTCH" v="0"></f><f n="LINEROUT" v=""></f><f n="IDVEND" v=""></f><f n="SWNOTET" v="N"></f><f n="TEXTFOB" v=""></f><f n="LINEDLVY" v=""></f><f n="IDOP01" v=""></f><f n="LINETOL" v=""></f><f n="SWREL" v="N"></f><f n="TEXTSTTS" v="UNRELEASED"></f><f n="TEXTOP02" v="CONTRACT TYPE"></f><f n="PARTYPE" v="1"></f><f n="LINEFRT" v=""></f><f n="TEXTCTRC" v=""></f></return-fields></screendata><msgs><msg no="Z00007" v="Please enter key field"></msg></msgs></trans>`))
	if err != nil {
		log.Fatal(err)
	}

	po501 := &PO501{}
	err = res.Unmarshal(po501)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("String: %s\n", *po501.AString)
	fmt.Printf("Int 64: %d\n", *po501.AInt64)
	fmt.Printf("Time: %v\n", *po501.ATime)
	fmt.Printf("Bool: %v\n", *po501.ABool)
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
//
//-------------
//
//String: FMISPR
//Int 64: 6942
//Time: 2020-12-31 00:00:00 +0000 UTC
//Bool: true
