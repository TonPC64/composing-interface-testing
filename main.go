package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type constructureFunctions struct {
	GetValue func(api string) int
}

type API struct {
	value int
	constructureFunctions
}

func NewAPI(functions constructureFunctions) *API {
	intFunc := &API{}
	intFunc.constructureFunctions = functions
	return intFunc
}

func GetValue(api string) int {
	res, err := http.Get(api)
	if err != nil {
		log.Println(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		var value map[string]int
		json.Unmarshal(data, &value)
		return value["value"]
	}
	return 0
}

func main() {
	cf := constructureFunctions{
		GetValue: GetValue,
	}
	log.Println(useApi(cf))
}

func useApi(cf constructureFunctions) int {
	aaa := NewAPI(cf)
	val := aaa.GetValue("https://api.myjson.com/bins/8hykh")
	return val
}
