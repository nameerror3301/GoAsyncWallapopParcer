package app

import (
	"GoAsyncWallapopParcer/internal/models"
	"fmt"
	"strconv"
	"sync"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

func Run() {
	var wg sync.WaitGroup
	var urlSlug string = "https://es.wallapop.com/item"
	var outCars, outElectronic, outMotos, outHome, outPhotoTV []byte

	for i := 40; i <= 400; {
		wg.Add(5)
		var (
			urlCarCat        = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=b5492bdb-1e92-4a18-bc2d-e06fe6b04501&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=120&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=12800&longitude=-3.69196", strconv.Itoa(i))
			urlElectronicCat = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=c32b03a2-07a3-40cc-bd3c-c2a2a05c5abb&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=80&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=15000&longitude=-3.69196", strconv.Itoa(i))
			urlMotosCat      = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=723dc55e-d048-4332-948b-7c84f8d1f1de&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=40&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=14000&longitude=-3.69196", strconv.Itoa(i))
			urlHomeCat       = fmt.Sprintf("https://api.wallapop.com/api/v3/real_estate/search?filters_source=seo_landing&latitude=40.41956&start=%s&order_by=most_relevance&step=0&category_ids=200&longitude=-3.69196&search_id=70de3634-8711-4d63-b4e6-4b5c7e828383", strconv.Itoa(i))
			urlPhotoTVCat    = fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=1e5fb024-9140-495d-92b5-47710871c000&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=200&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=12545&longitude=-3.69196", strconv.Itoa(i))
		)

		// Сбор товаров из категории "Все для машин"
		go func(urlSlug string, urlCarCat string) {
			outCars = FindAll(urlSlug, urlCarCat, "Cars")
			defer wg.Done()
		}(urlSlug, urlCarCat)

		// Сбор товаров из категории "Электроника"
		go func(urlSlug string, urlElectronicCat string) {
			outElectronic = FindAll(urlSlug, urlElectronicCat, "Electronic")
			defer wg.Done()
		}(urlSlug, urlElectronicCat)

		// Сбор товаров из категории "Все для дома"
		go func(urlSlug string, urlHomeCat string) {
			outHome = FindAll(urlSlug, urlHomeCat, "Home")
			defer wg.Done()
		}(urlSlug, urlHomeCat)

		// Сбор товаров из категории "Мотоциклы"
		go func(urlSlug string, urlMotosCat string) {
			outMotos = FindAll(urlSlug, urlMotosCat, "Motos")
			defer wg.Done()
		}(urlSlug, urlMotosCat)

		// Сбор товаров из категории "Все для Фото и Тв"
		go func(urlSlug string, urlPhotoTVCat string) {
			outPhotoTV = FindAll(urlSlug, urlPhotoTVCat, "PhotoTV")
			defer wg.Done()
		}(urlSlug, urlPhotoTVCat)

		i = i + 40
		wg.Wait()
	}

	/*
		Отправка запроса в сервис хранилища
	*/
	fmt.Println("Машины и все что с ними связанно", string(outCars))
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Электроника вся", string(outElectronic))
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Все для дома", string(outHome))
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Мотоциклы и все что с ними связанно", string(outMotos))
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Все для фото и тв", string(outPhotoTV))
}

func FindAll(urlSlug string, url string, category string) []byte {
	var out []byte
	var err error

	err = models.FindAllInCategory(url, urlSlug, category)

	if err != nil {
		logrus.Infof("Err parce data - %s", err)
	}

	switch category {
	case "Cars":
		out, err = gojson.Marshal(models.CarsCat)
		if err != nil {
			logrus.Debugf("Err marshal struct to json Cars - %s", &err)
		}
	case "Electronic":
		out, err = gojson.Marshal(models.ElecCat)
		if err != nil {
			logrus.Debugf("Err marshal struct to json Electronic - %s", &err)
		}
	case "Home":
		out, err = gojson.Marshal(models.HomeCat)
		if err != nil {
			logrus.Debugf("Err marshal struct to json Home - %s", &err)
		}
	case "Motos":
		out, err = gojson.Marshal(models.MotosCat)
		if err != nil {
			logrus.Debugf("Err marshal struct to json Motos - %s", &err)
		}
	case "PhotoTV":
		out, err = gojson.Marshal(models.PhotoTVCat)
		if err != nil {
			logrus.Debugf("Err marshal struct to json PhoneTV - %s", &err)
		}
	}
	return out
}
