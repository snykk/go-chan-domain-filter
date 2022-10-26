package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	}

	if !website.Valid {
		chErr <- errors.New("domain not valid")
		return
	}

	if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	}

	substractDom := strings.Split(website.Domain, ".")
	var tld, idn_tld string

	switch string(substractDom[len(substractDom)-1]) {
	case "com":
		tld = ".com"
		idn_tld = ".co.id"
	case "gov":
		tld = ".gov"
		idn_tld = ".go.id"
	case "org":
		tld = ".org"
		idn_tld = ".org.id"
	default:
		tld = "." + string(substractDom[len(substractDom)-1])
		idn_tld = "TLD"
	}

	website.TLD = tld
	website.IDN_TLD = idn_tld

	ch <- website
	time.Sleep(100 * time.Millisecond)
}

var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	var result = []RowData{}

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}

	for i := 0; i < len(data); i++ {
		select {
		case msg := <-errCh:
			return []RowData{}, msg
		case msg := <-ch:
			if TLD == msg.TLD {
				result = append(result, msg)
			}
		}

	}

	return result, nil
}

func main() {
	fmt.Println(FilterAndFillData(".org", []RowData{
		{
			RankWebsite: 1,
			Domain:      "google.com",
			Valid:       true,
			RefIPs:      1,
			TLD:         "",
			IDN_TLD:     "",
		}, {
			RankWebsite: 2,
			Domain:      "bukanruang.org",
			Valid:       true,
			RefIPs:      1,
			TLD:         "",
			IDN_TLD:     "",
		}, {
			RankWebsite: 3,
			Domain:      "bukanjudi.xyz",
			Valid:       false,
			RefIPs:      1,
			TLD:         "",
			IDN_TLD:     "",
		},
	}))
}
