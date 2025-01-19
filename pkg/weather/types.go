package weather

type WeatherResponse struct {
	Timezone     string `json:"timezone"`
	CurrentUnits struct {
		Temperature2m string `json:"temperature_2m"`
	} `json:"current_units"`
	Current struct {
		Time             string  `json:"time"`
		Temperature2m    float64 `json:"temperature_2m"`
		RelativeHumidity int     `json:"relative_humidity_2m"`
		WindSpeed10m     float64 `json:"wind_speed_10m"`
		WindDirection10m int     `json:"wind_direction_10m"`
		WindGusts10m     float64 `json:"wind_gusts_10m"`
		Precipitation    float64 `json:"precipitation"`
	} `json:"current"`
	Daily struct {
		Time                        []string  `json:"time"`
		Temperature2mMax            []float64 `json:"temperature_2m_max"`
		Temperature2mMin            []float64 `json:"temperature_2m_min"`
		Sunrise                     []string  `json:"sunrise"`
		Sunset                      []string  `json:"sunset"`
		PrecipitationSum            []float64 `json:"precipitation_sum"`
		PrecipitationProbabilityMax []int     `json:"precipitation_probability_max"`
		WindSpeed10mMax             []float64 `json:"wind_speed_10m_max"`
	} `json:"daily"`
	Hourly struct {
		Time                     []string  `json:"time"`
		Temperature2m            []float64 `json:"temperature_2m"`
		RelativeHumidity2m       []int     `json:"relative_humidity_2m"`
		PrecipitationProbability []int     `json:"precipitation_probability"`
		WindSpeed10m             []float64 `json:"wind_speed_10m"`
	} `json:"hourly"`
}
