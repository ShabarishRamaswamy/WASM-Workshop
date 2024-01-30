//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"log"
	"syscall/js"
	"time"
)

type Partner struct {
	ID          string `json:"_id"`
	SiteID      string `json:"siteID"`
	CompanyID   string `json:"companyID"`
	SiteName    string `json:"siteName"`
	CompanyName string `json:"companyName"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

func main() {
	wait := make(chan struct{})

	js.Global().Set("findAllUniqueSiteID", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		startTime := time.Now()

		stringOfP := p[0].String()
		dataToBeExported := findAllUniqueSiteID(&stringOfP)
		log.Println(time.Since(startTime))

		// Convert the slice to a JavaScript array
		jsArray := js.Global().Get("Array").New(len(dataToBeExported))
		for i, str := range dataToBeExported {
			jsArray.SetIndex(i, js.ValueOf(str))
		}

		return jsArray
	}))

	<-wait
}

func findAllUniqueSiteID(allData *string) []string {
	allPartnerData := []Partner{}
	_ = json.Unmarshal([]byte(*allData), &allPartnerData)

	solutionMap := make(map[string]bool, 0)
	finalData := []string{}
	for _, data := range allPartnerData {
		if !solutionMap[data.SiteID] {
			solutionMap[data.SiteID] = true
			finalData = append(finalData, data.SiteID)
		}
	}
	return finalData
}
