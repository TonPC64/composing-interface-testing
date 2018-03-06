package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type constructureFunctions struct {
	GetValue func() int
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

func GetValue() int {
	res, err := http.Get("https://api.myjson.com/bins/8hykh")
	if err != nil {
		log.Println(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	var value struct {
		Value int
	}
	json.Unmarshal(data, &value)
	return value.Value
}

func main() {
	cf := constructureFunctions{
		GetValue: GetValue,
	}
	useApi(cf)
}

func useApi(cf constructureFunctions) int {
	aaa := NewAPI(cf)
	val := aaa.GetValue()
	return val
}
