package models

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	gojson "github.com/goccy/go-json"
)

// Основная структура в которую помещаются данные ответа на запрос от Wallapop
type AllData struct {
	SearchObjects []struct {
		ID          string  `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Distance    float64 `json:"distance"`
		Images      []struct {
			Original       string `json:"original"`
			Xsmall         string `json:"xsmall"`
			Small          string `json:"small"`
			Large          string `json:"large"`
			Medium         string `json:"medium"`
			Xlarge         string `json:"xlarge"`
			OriginalWidth  int    `json:"original_width"`
			OriginalHeight int    `json:"original_height"`
		} `json:"images"`
		User struct {
			ID        string `json:"id"`
			MicroName string `json:"micro_name"`
			Image     struct {
				Original string `json:"original"`
			} `json:"image"`
			Online bool   `json:"online"`
			Kind   string `json:"kind"`
		} `json:"user"`
		Price        float64 `json:"price"`
		Currency     string  `json:"currency"`
		FreeShipping bool    `json:"free_shipping"`
		WebSlug      string  `json:"web_slug"`
		CategoryID   int     `json:"category_id"`
		Shipping     struct {
			ItemIsShippable     bool        `json:"item_is_shippable"`
			UserAllowsShipping  bool        `json:"user_allows_shipping"`
			CostConfigurationID interface{} `json:"cost_configuration_id"`
		} `json:"shipping"`
		Location struct {
			City        string `json:"city"`
			PostalCode  string `json:"postal_code"`
			CountryCode string `json:"country_code"`
		} `json:"location"`
		TypeAttributes interface{} `json:"type_attributes"`
	} `json:"search_objects"`
}

type RequestLast struct {
	User struct {
		Name        string `json:"user_name"`
		DateRegistr string `json:"user_date_reg"`
		PhoneNumber string `json:"user_phone"`
	} `json:"user_data"`
	Products struct {
		ProdName    string `json:"product_name"`
		PhotoUrl    string `json:"photo_url"`
		Price       string `json:"price"`
		Description string `json:"description"`
		Url         string `json:"url"`
	} `json:"product_data"`
}

// Объявление структур
var (
	Cars    []RequestLast
	Elec    []RequestLast
	Home    []RequestLast
	Motos   []RequestLast
	PhotoTV []RequestLast
)

// Основная универсальная функция для работы со всеми категориями
func FindAllInCategory(url string, urlSlug string, category string) {
	var all AllData

	resp, err := http.Get(url)
	if err != nil {
		logrus.Warnf("Error request to wallapop data - %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Warnf("Problems with receiving data from a resource - %d", resp.StatusCode)
		time.Sleep(time.Second * 10)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Warnf("Err parce resp body - %s", err)
	}

	if err := gojson.Unmarshal(body, &all); err != nil {
		logrus.Errorf("Err unmarshal json to struct - %s", err)
	}

	switch category {
	case "Cars":
		AppendData(&all, &Cars, urlSlug)
	case "Electronic":
		AppendData(&all, &Elec, urlSlug)
	case "Home":
		AppendData(&all, &Home, urlSlug)
	case "Motos":
		AppendData(&all, &Motos, urlSlug)
	case "PhotoTV":
		AppendData(&all, &PhotoTV, urlSlug)
	}
}

func AppendData(all *AllData, data *[]RequestLast, urlSlug string) {
	for _, val := range all.SearchObjects {
		*data = append(*data, RequestLast{
			User: struct {
				Name        string "json:\"user_name\""
				DateRegistr string "json:\"user_date_reg\""
				PhoneNumber string "json:\"user_phone\""
			}{
				Name:        val.User.MicroName,
				DateRegistr: "No Data",
				PhoneNumber: "No Data",
			},
			Products: struct {
				ProdName    string "json:\"product_name\""
				PhotoUrl    string "json:\"photo_url\""
				Price       string "json:\"price\""
				Description string "json:\"description\""
				Url         string "json:\"url\""
			}{
				ProdName:    val.Title,
				PhotoUrl:    val.Images[0].Original,
				Price:       fmt.Sprintf("%f", val.Price),
				Description: val.Description,
				Url:         val.WebSlug,
			},
		})
	}
}
