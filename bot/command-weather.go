package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

const URL string = "https://api.openweathermap.org/data/2.5/weather?"

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`

	Name string `json:"name"`
}

func getCurrentWeather(message string) *discordgo.MessageSend {
	// 匹配五位数的 US 邮编
	r, _ := regexp.Compile(`\d{5}`)
	zip := r.FindString(message)

	if zip == "" {
		return &discordgo.MessageSend{
			Content: "没有找到该邮编",
		}
	}

	weatherURL := fmt.Sprintf("%szip=%s&units=imperial&appid=%s", URL, zip, OpenWeatherToken)
	client := http.Client{Timeout: 5 * time.Second}

	// Query OpenWeather API
	response, err := client.Get(weatherURL)

	if err != nil {
		return &discordgo.MessageSend{
			Content: "找天气的过程中产生了错误",
		}
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var data WeatherData
	json.Unmarshal([]byte(body), &data)

	city := data.Name
	conditions := data.Weather[0].Description
	temperature := strconv.FormatFloat(data.Main.Temp, 'f', 2, 64)
	humidity := strconv.Itoa(data.Main.Humidity)
	wind := strconv.FormatFloat(data.Wind.Speed, 'f', 2, 64)

	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "当前天气",
			Description: city + " 的天气",
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Conditions", Value: conditions, Inline: true},
				{Name: "Temperature", Value: temperature + "°F", Inline: true},
				{Name: "Humidity", Value: humidity + "%", Inline: true},
				{Name: "Wind", Value: wind + " mph", Inline: true},
				{Name: "Conditions", Value: conditions, Inline: true},
				{Name: "Conditions", Value: conditions, Inline: true},
				{Name: "Conditions", Value: conditions, Inline: true},
				{Name: "Conditions", Value: conditions, Inline: true},
				{Name: "Conditions", Value: conditions, Inline: true},
			},
		}},
	}

	return embed
}
