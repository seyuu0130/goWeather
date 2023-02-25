package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

const port = 8080

func main() {
	// 天気
	wetherHandler := func(w http.ResponseWriter, _ *http.Request) {
		resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=35.69&longitude=139.69&daily=weathercode,temperature_2m_max,temperature_2m_min&timezone=Asia%2FTokyo")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var dto WeatherDto

		// json.Unmarshal(body, &dto)
		var decoder = json.NewDecoder(bytes.NewReader(body))
		decoder.UseNumber()

		if err := decoder.Decode(&dto); err != nil {
			log.Fatal(err)
			return
		}

		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, dto); err != nil {
			log.Fatal(err)
		}

		// for i := 0; i < 7; i++ {
		// 	io.WriteString(w, dto.Daily.Time[i]+":"+dto.Daily.Weathercode[i].string())
		// 	io.WriteString(w, "\n")
		// }
		fmt.Println(dto)
	}
	http.HandleFunc("/wether/", wetherHandler)

	// ポート番号 8080 で待ち受けを開始
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
