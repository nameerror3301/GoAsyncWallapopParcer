package app

import (
	"GoAsyncWallapopParcer/internal/config"
	"GoAsyncWallapopParcer/internal/models"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

const (
	Cars       = "Cars"
	Electronic = "Electronic"
	Motos      = "Motos"
	Home       = "Home"
	PhotoTV    = "PhotoTV"
)

const (
	QueryElectronic = "electronics"
	QueryCars       = "cars"
	QueryMotos      = "motos"
	QueryHome       = "home"
	QueryPhotoTV    = "phototv"
)

func Run() {
	var wg sync.WaitGroup
	var urlSlug string = "https://es.wallapop.com/item"
	conf := config.ReadConfig()

	for i := 40; i <= 400; {
		wg.Add(5)
		var (
			urlCar        = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=b5492bdb-1e92-4a18-bc2d-e06fe6b04501&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=120&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=12800&longitude=-3.69196", strconv.Itoa(i))
			urlElectronic = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=c32b03a2-07a3-40cc-bd3c-c2a2a05c5abb&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=80&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=15000&longitude=-3.69196", strconv.Itoa(i))
			urlMotos      = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=723dc55e-d048-4332-948b-7c84f8d1f1de&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=40&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=14000&longitude=-3.69196", strconv.Itoa(i))
			urlHome       = fmt.Sprintf("https://api.wallapop.com/api/v3/real_estate/search?filters_source=seo_landing&latitude=40.41956&start=%s&order_by=most_relevance&step=0&category_ids=200&longitude=-3.69196&search_id=70de3634-8711-4d63-b4e6-4b5c7e828383", strconv.Itoa(i))
			urlPhotoTV    = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=1e5fb024-9140-495d-92b5-47710871c000&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=200&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=12545&longitude=-3.69196", strconv.Itoa(i))
		)

		// ???????? ?????????????? ???? ?????????????????? "?????? ?????? ??????????"
		go func(urlSlug string, urlCar string) {
			models.FindAllInCategory(urlCar, urlSlug, Cars)
			defer wg.Done()
		}(urlSlug, urlCar)

		// ???????? ?????????????? ???? ?????????????????? "??????????????????????"
		go func(urlSlug string, urlElectronic string) {
			models.FindAllInCategory(urlElectronic, urlSlug, Electronic)
			defer wg.Done()
		}(urlSlug, urlElectronic)

		// ???????? ?????????????? ???? ?????????????????? "?????? ?????? ????????"
		go func(urlSlug string, urlHome string) {
			models.FindAllInCategory(urlHome, urlSlug, Home)
			defer wg.Done()
		}(urlSlug, urlHome)

		// ???????? ?????????????? ???? ?????????????????? "??????????????????"
		go func(urlSlug string, urlMotos string) {
			models.FindAllInCategory(urlMotos, urlSlug, Motos)
			defer wg.Done()
		}(urlSlug, urlMotos)

		// ???????? ?????????????? ???? ?????????????????? "?????? ?????? ???????? ?? ????"
		go func(urlSlug string, urlPhotoTV string) {
			models.FindAllInCategory(urlPhotoTV, urlSlug, PhotoTV)
			defer wg.Done()
		}(urlSlug, urlPhotoTV)

		i = i + 40
		wg.Wait()
	}

	/*
		???????????????? ?????????????? ?? ???????????? ??????????????????
	*/
	SendData(MarshalData(models.Elec), QueryElectronic, conf.Data.JwtToken)

	SendData(MarshalData(models.Cars), QueryCars, conf.Data.JwtToken)

	SendData(MarshalData(models.Home), QueryHome, conf.Data.JwtToken)

	SendData(MarshalData(models.Motos), QueryMotos, conf.Data.JwtToken)

	SendData(MarshalData(models.PhotoTV), QueryPhotoTV, conf.Data.JwtToken)
}

func MarshalData(data interface{}) []byte {
	out, err := gojson.Marshal(data)
	if err != nil {
		logrus.Errorf("Err marshal struct to json - %s", &err)
	}

	return out
}

func SendData(data []byte, category string, token string) {
	conf := config.ReadConfig()

	url := fmt.Sprintf("%sadd?category=%s&market=wallapop", conf.Data.OutStorageAddr, category)

	reader := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		logrus.Error("Err request generation - %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		logrus.Error("Err send data - %s", err)
		time.Sleep(5 * time.Second)
		SendData(MarshalData(models.Elec), category, conf.Data.JwtToken)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		logrus.Warnf("Check jwt token - %d", http.StatusUnauthorized)
	}

	if res.StatusCode != http.StatusOK {
		logrus.Errorf("Err sending data - %s", err)
	} else {
		logrus.Info("Success sending data")
	}
}
