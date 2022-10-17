package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var port = ":5500"

type Status struct {
	Status U `json:"users"`
}

type U struct {
	Wind        int    `json:"wind"`
	Water       int    `json:"water"`
	WindStatus  string `json:"wind_status"`
	WaterStatus string `json:"water_status"`
}

func main() {
	http.HandleFunc("/", random)
	http.ListenAndServe(port, nil)
}

func random(w http.ResponseWriter, r *http.Request) {
	for {
		time.Sleep(1 * time.Second)
		jsonFile, err := os.Open("status.json")
		var tpl, errTemp = template.ParseFiles("index.html")

		if err != nil {
			fmt.Println(err)
		}

		if errTemp != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var x Status
		json.Unmarshal(byteValue, &x)

		x.Status.Water = rand.Intn(100)
		x.Status.Wind = rand.Intn(100)

		if x.Status.Water < 5 {
			x.Status.WaterStatus = "Aman"
		} else if x.Status.Water <= 8 && x.Status.Water >= 6 {
			x.Status.WaterStatus = "Siaga"
		} else if x.Status.Water > 8 {
			x.Status.WaterStatus = "Bahaya"
		}

		if x.Status.Wind < 6 {
			x.Status.WindStatus = "Aman"
		} else if x.Status.Wind <= 15 && x.Status.Water >= 7 {
			x.Status.WindStatus = "Siaga"
		} else if x.Status.Wind > 15 {
			x.Status.WindStatus = "Bahaya"
		}

		tpl.Execute(w, x.Status)
	}
}