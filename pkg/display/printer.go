package display

import (
	"fmt"
	"time"
	"weather-cli/pkg/weather"

	"github.com/fatih/color"
)

const (
	currentHeader = "=== Current Conditions ==="
	hourlyHeader  = "=== Next 24 Hours ==="
	dailyHeader   = "=== 7-Day Forecast ==="
)

func getWindDirection(degrees int) string {
	directions := []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	index := int((float64(degrees) + 11.25) / 22.5)
	return directions[index%16]
}

func formatTime(timeStr string, timezone string) (string, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}

	t, err := time.ParseInLocation("2006-01-02T15:04", timeStr, loc)
	if err != nil {
		return "", err
	}

	return t.Format("3:04 PM"), nil
}

func PrintCurrentWeather(w *weather.WeatherResponse) {
	bold := color.New(color.Bold).SprintFunc()
	temp := color.New(color.FgHiYellow).SprintfFunc()
	humid := color.New(color.FgBlue).SprintfFunc()
	wind := color.New(color.FgGreen).SprintfFunc()
	sun := color.New(color.FgHiWhite).SprintfFunc()

	tempUnit := w.CurrentUnits.Temperature2m

	fmt.Printf("\n%s\n", bold(currentHeader))
	fmt.Printf("🌡️  %s | 💧 %s | 💨 %s | ☔ %.2f in\n",
		temp("%.1f%s", w.Current.Temperature2m, tempUnit),
		humid("%d%%", w.Current.RelativeHumidity),
		wind("%.1f mph %s", w.Current.WindSpeed10m, getWindDirection(w.Current.WindDirection10m)),
		w.Current.Precipitation)

	fmt.Println("sunrise/sunset")
	sunrise, _ := formatTime(w.Daily.Sunrise[0], w.Timezone)
	sunset, _ := formatTime(w.Daily.Sunset[0], w.Timezone)
	fmt.Printf("🌅 Sunrise: %s | 🌇 Sunset: %s\n",
		sun(sunrise),
		sun(sunset))
}

func PrintHourlyForecast(w *weather.WeatherResponse) {
	bold := color.New(color.Bold).SprintFunc()
	temp := color.New(color.FgHiYellow).SprintfFunc()
	precip := color.New(color.FgBlue).SprintfFunc()
	wind := color.New(color.FgGreen).SprintfFunc()

	tempUnit := w.CurrentUnits.Temperature2m

	fmt.Printf("\n%s\n", bold(hourlyHeader))
	fmt.Println("┌──────────┬──────────┬────────┬─────────┐")
	fmt.Println("│   TIME   │   TEMP   │  RAIN  │  WIND   │")
	fmt.Println("├──────────┼──────────┼────────┼─────────┤")

	for i := 0; i < 24; i += 3 {
		localTime, _ := formatTime(w.Hourly.Time[i], w.Timezone)

		isLast := i >= 21
		lineEnd := "\n"
		if isLast {
			lineEnd = ""
		}

		fmt.Printf("│ %-8s │ %s │ %s │ %s │%s",
			localTime,
			temp("%6.1f%s", w.Hourly.Temperature2m[i], tempUnit),
			precip("%3d%%  ", w.Hourly.PrecipitationProbability[i]),
			wind("%3.0f mph", w.Hourly.WindSpeed10m[i]),
			lineEnd)
	}
	fmt.Printf("\n└──────────┴──────────┴────────┴─────────┘\n")
}

func PrintDailyForecast(w *weather.WeatherResponse) {
	bold := color.New(color.Bold).SprintFunc()
	temp := color.New(color.FgHiYellow).SprintfFunc()
	precip := color.New(color.FgBlue).SprintfFunc()
	wind := color.New(color.FgGreen).SprintfFunc()

	tempUnit := w.CurrentUnits.Temperature2m

	fmt.Printf("\n%s\n", bold(dailyHeader))
	fmt.Println("┌───────────┬───────────┬──────────┬─────────┬──────────┐")
	fmt.Println("│   DATE    │ HIGH/LOW  │ PRECIP   │  WIND   │  CHANCE  │")
	fmt.Println("├───────────┼───────────┼──────────┼─────────┼──────────┤")

	for i := 0; i < len(w.Daily.Time); i++ {
		date, _ := time.Parse("2006-01-02", w.Daily.Time[i])

		isLast := i == len(w.Daily.Time)-1
		lineEnd := "\n"
		if isLast {
			lineEnd = ""
		}

		fmt.Printf("│ %-9s │ %s │ %s │ %s │ %s │%s",
			date.Format("Mon 1/2"),
			temp("%3.0f/%-3.0f%s", w.Daily.Temperature2mMax[i], w.Daily.Temperature2mMin[i], tempUnit),
			precip("%5.2f\"  ", w.Daily.PrecipitationSum[i]),
			wind("%3.0f mph", w.Daily.WindSpeed10mMax[i]),
			precip("%3d%%    ", w.Daily.PrecipitationProbabilityMax[i]),
			lineEnd)
	}
	fmt.Printf("\n└───────────┴───────────┴──────────┴─────────┴──────────┘\n")
}
