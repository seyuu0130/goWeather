package main

type Daily struct {
	Time             []string  `json:"time"`
	Weathercode      []Weather `json:"weathercode"`
	Temperature2mMax []float64 `json:"temperature_2m_max"`
	Temperature2mMin []float64 `json:"temperature_2m_min"`
}

type WeatherDto struct {
	Timezone string  `json:"timezone"`
	Latitude float64 `json:"latitude"`
	Daily    Daily   `json:"daily"`
}
