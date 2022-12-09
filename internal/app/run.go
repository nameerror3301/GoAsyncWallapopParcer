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
	var outCars []byte

	for i := 40; i <= 400; {
		wg.Add(1)
		go func(urlSlug string) {
			outCars = FindAllCars(urlSlug, i)
			wg.Done()
		}(urlSlug)
		i = i + 40
		wg.Wait()
	}

	fmt.Println(string(outCars))
}

func FindAllCars(urlSlug string, i int) []byte {

	urlCars := fmt.Sprintf("https://api.wallapop.com/api/v3/general/search?user_province=Madrid&latitude=40.41956&start=%s&user_region=Comunidad+de+Madrid&user_city=Madrid&search_id=b5492bdb-1e92-4a18-bc2d-e06fe6b04501&country_code=ES&user_postal_code=28014&experiment=freshness_factor_variation_D&items_count=120&density_type=20&filters_source=seo_landing&order_by=closest&step=0&category_ids=12800&longitude=-3.69196", strconv.Itoa(i))
	err := models.FindAllInCategory(urlCars, urlSlug)

	if err != nil {
		logrus.Infof("Err parce data - %s", err)
	}

	out, err := gojson.Marshal(models.Last)
	if err != nil {
		logrus.Debugf("Err marshal struct to json - %s", &err)
	}
	return out
}
