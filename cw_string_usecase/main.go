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
	// startTime := time.Now()
	// jsonFile, err := os.Open("data.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer jsonFile.Close()

	// allPartnerData := []Partner{}
	// dataFromJSON, _ := io.ReadAll(jsonFile)
	// dataFromJSON := jsonBlob

	// _ = json.Unmarshal([]byte(jsonBlob), &allPartnerData)
	// fmt.Println(allPartnerData)

	// unique := findAllUniqueSiteID(allPartnerData)
	// fmt.Println(unique, len(unique))

	// fmt.Println("Time taken : ", time.Since(startTime))

	// c := make(chan string)
	// js.Global().Set("findAllUniqueSiteID", methodWrapper())
	// <-c

	wait := make(chan struct{}, 0)

	js.Global().Set("findAllUniqueSiteID", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		// go func() {
		// 	// fmt.Println(this, p[0])
		startTime := time.Now()

		stringOfP := p[0].String()
		dataToBeExported := findAllUniqueSiteID(&stringOfP)
		log.Println(time.Since(startTime))

		// Convert the slice to a JavaScript array
		jsArray := js.Global().Get("Array").New(len(dataToBeExported))
		for i, str := range dataToBeExported {
			jsArray.SetIndex(i, js.ValueOf(str))
		}

		// Simulate asynchronous operation
		// time.Sleep(3 * time.Second)

		// Signal completion to JavaScript
		// close(c)
		// 	fmt.Println(time.Since(startTime))
		// }()
		return jsArray
		// return nil
	}))

	<-wait
}

// func findAllUniqueSiteID(allData []Partner) []string {
// 	solutionMap := make(map[string]bool, 0)
// 	finalData := []string{}
// 	for _, data := range allData {
// 		if !solutionMap[data.SiteID] {
// 			solutionMap[data.SiteID] = true
// 			finalData = append(finalData, data.SiteID)
// 		}
// 	}
// 	return finalData
// }

func findAllUniqueSiteID(allData *string) []string {
	// fmt.Println("allData", allData)
	allPartnerData := []Partner{}
	_ = json.Unmarshal([]byte(*allData), &allPartnerData)
	// fmt.Println("allPartnerData", allPartnerData)

	solutionMap := make(map[string]bool, 0)
	finalData := []string{}
	for _, data := range allPartnerData {
		if !solutionMap[data.SiteID] {
			solutionMap[data.SiteID] = true
			finalData = append(finalData, data.SiteID)
		}
	}
	// fmt.Println(finalData)
	return finalData
}

// func methodWrapper() js.Func {
// 	uniqueFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
// 		if len(args) != 1 {
// 			return "Invalid no of arguments passed"
// 		}
// 		inputJSON := args[0].String()
// 		fmt.Printf("input %s\n", inputJSON)
// 		pretty := findAllUniqueSiteID(inputJSON)
// 		return pretty
// 	})
// 	return uniqueFunc
// }
