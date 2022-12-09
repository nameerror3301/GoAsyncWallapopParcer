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
				Original       string `json:"original"`
				Xsmall         string `json:"xsmall"`
				Small          string `json:"small"`
				Large          string `json:"large"`
				Medium         string `json:"medium"`
				Xlarge         string `json:"xlarge"`
				OriginalWidth  int    `json:"original_width"`
				OriginalHeight int    `json:"original_height"`
			} `json:"image"`
			Online bool   `json:"online"`
			Kind   string `json:"kind"`
		} `json:"user"`
		Flags struct {
			Pending  bool `json:"pending"`
			Sold     bool `json:"sold"`
			Reserved bool `json:"reserved"`
			Banned   bool `json:"banned"`
			Expired  bool `json:"expired"`
			Onhold   bool `json:"onhold"`
		} `json:"flags"`
		VisibilityFlags struct {
			Bumped        bool `json:"bumped"`
			Highlighted   bool `json:"highlighted"`
			Urgent        bool `json:"urgent"`
			CountryBumped bool `json:"country_bumped"`
			Boosted       bool `json:"boosted"`
		} `json:"visibility_flags"`
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
		SupportsShipping bool   `json:"supports_shipping"`
		ShippingAllowed  bool   `json:"shipping_allowed"`
		SellerID         string `json:"seller_id"`
		Favorited        bool   `json:"favorited"`
		CreationDate     int64  `json:"creation_date"`
		ModificationDate int64  `json:"modification_date"`
		Location         struct {
			City        string `json:"city"`
			PostalCode  string `json:"postal_code"`
			CountryCode string `json:"country_code"`
		} `json:"location"`
		TypeAttributes interface{} `json:"type_attributes"`
	} `json:"search_objects"`
}

// Структуры для формирования данных и дальнейшей их отправке (Из за паралельного подхода в сборе всех товаров со всех категорий пришлось прибегнуть к такому решения с множеством структур)
type Cars struct {
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
			Original       string `json:"original"`
			Xsmall         string `json:"xsmall"`
			Small          string `json:"small"`
			Large          string `json:"large"`
			Medium         string `json:"medium"`
			Xlarge         string `json:"xlarge"`
			OriginalWidth  int    `json:"original_width"`
			OriginalHeight int    `json:"original_height"`
		} `json:"image"`
		Online bool   `json:"online"`
		Kind   string `json:"kind"`
	} `json:"user"`
	Flags struct {
		Pending  bool `json:"pending"`
		Sold     bool `json:"sold"`
		Reserved bool `json:"reserved"`
		Banned   bool `json:"banned"`
		Expired  bool `json:"expired"`
		Onhold   bool `json:"onhold"`
	} `json:"flags"`
	VisibilityFlags struct {
		Bumped        bool `json:"bumped"`
		Highlighted   bool `json:"highlighted"`
		Urgent        bool `json:"urgent"`
		CountryBumped bool `json:"country_bumped"`
		Boosted       bool `json:"boosted"`
	} `json:"visibility_flags"`
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
	SupportsShipping bool   `json:"supports_shipping"`
	ShippingAllowed  bool   `json:"shipping_allowed"`
	SellerID         string `json:"seller_id"`
	Favorited        bool   `json:"favorited"`
	CreationDate     int64  `json:"creation_date"`
	ModificationDate int64  `json:"modification_date"`
	Location         struct {
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		CountryCode string `json:"country_code"`
	} `json:"location"`
}

type Electronic struct {
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
			Original       string `json:"original"`
			Xsmall         string `json:"xsmall"`
			Small          string `json:"small"`
			Large          string `json:"large"`
			Medium         string `json:"medium"`
			Xlarge         string `json:"xlarge"`
			OriginalWidth  int    `json:"original_width"`
			OriginalHeight int    `json:"original_height"`
		} `json:"image"`
		Online bool   `json:"online"`
		Kind   string `json:"kind"`
	} `json:"user"`
	Flags struct {
		Pending  bool `json:"pending"`
		Sold     bool `json:"sold"`
		Reserved bool `json:"reserved"`
		Banned   bool `json:"banned"`
		Expired  bool `json:"expired"`
		Onhold   bool `json:"onhold"`
	} `json:"flags"`
	VisibilityFlags struct {
		Bumped        bool `json:"bumped"`
		Highlighted   bool `json:"highlighted"`
		Urgent        bool `json:"urgent"`
		CountryBumped bool `json:"country_bumped"`
		Boosted       bool `json:"boosted"`
	} `json:"visibility_flags"`
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
	SupportsShipping bool   `json:"supports_shipping"`
	ShippingAllowed  bool   `json:"shipping_allowed"`
	SellerID         string `json:"seller_id"`
	Favorited        bool   `json:"favorited"`
	CreationDate     int64  `json:"creation_date"`
	ModificationDate int64  `json:"modification_date"`
	Location         struct {
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		CountryCode string `json:"country_code"`
	} `json:"location"`
}

type Home struct {
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
			Original       string `json:"original"`
			Xsmall         string `json:"xsmall"`
			Small          string `json:"small"`
			Large          string `json:"large"`
			Medium         string `json:"medium"`
			Xlarge         string `json:"xlarge"`
			OriginalWidth  int    `json:"original_width"`
			OriginalHeight int    `json:"original_height"`
		} `json:"image"`
		Online bool   `json:"online"`
		Kind   string `json:"kind"`
	} `json:"user"`
	Flags struct {
		Pending  bool `json:"pending"`
		Sold     bool `json:"sold"`
		Reserved bool `json:"reserved"`
		Banned   bool `json:"banned"`
		Expired  bool `json:"expired"`
		Onhold   bool `json:"onhold"`
	} `json:"flags"`
	VisibilityFlags struct {
		Bumped        bool `json:"bumped"`
		Highlighted   bool `json:"highlighted"`
		Urgent        bool `json:"urgent"`
		CountryBumped bool `json:"country_bumped"`
		Boosted       bool `json:"boosted"`
	} `json:"visibility_flags"`
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
	SupportsShipping bool   `json:"supports_shipping"`
	ShippingAllowed  bool   `json:"shipping_allowed"`
	SellerID         string `json:"seller_id"`
	Favorited        bool   `json:"favorited"`
	CreationDate     int64  `json:"creation_date"`
	ModificationDate int64  `json:"modification_date"`
	Location         struct {
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		CountryCode string `json:"country_code"`
	} `json:"location"`
}

type Motos struct {
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
			Original       string `json:"original"`
			Xsmall         string `json:"xsmall"`
			Small          string `json:"small"`
			Large          string `json:"large"`
			Medium         string `json:"medium"`
			Xlarge         string `json:"xlarge"`
			OriginalWidth  int    `json:"original_width"`
			OriginalHeight int    `json:"original_height"`
		} `json:"image"`
		Online bool   `json:"online"`
		Kind   string `json:"kind"`
	} `json:"user"`
	Flags struct {
		Pending  bool `json:"pending"`
		Sold     bool `json:"sold"`
		Reserved bool `json:"reserved"`
		Banned   bool `json:"banned"`
		Expired  bool `json:"expired"`
		Onhold   bool `json:"onhold"`
	} `json:"flags"`
	VisibilityFlags struct {
		Bumped        bool `json:"bumped"`
		Highlighted   bool `json:"highlighted"`
		Urgent        bool `json:"urgent"`
		CountryBumped bool `json:"country_bumped"`
		Boosted       bool `json:"boosted"`
	} `json:"visibility_flags"`
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
	SupportsShipping bool   `json:"supports_shipping"`
	ShippingAllowed  bool   `json:"shipping_allowed"`
	SellerID         string `json:"seller_id"`
	Favorited        bool   `json:"favorited"`
	CreationDate     int64  `json:"creation_date"`
	ModificationDate int64  `json:"modification_date"`
	Location         struct {
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		CountryCode string `json:"country_code"`
	} `json:"location"`
}

type PhotoTV struct {
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
			Original       string `json:"original"`
			Xsmall         string `json:"xsmall"`
			Small          string `json:"small"`
			Large          string `json:"large"`
			Medium         string `json:"medium"`
			Xlarge         string `json:"xlarge"`
			OriginalWidth  int    `json:"original_width"`
			OriginalHeight int    `json:"original_height"`
		} `json:"image"`
		Online bool   `json:"online"`
		Kind   string `json:"kind"`
	} `json:"user"`
	Flags struct {
		Pending  bool `json:"pending"`
		Sold     bool `json:"sold"`
		Reserved bool `json:"reserved"`
		Banned   bool `json:"banned"`
		Expired  bool `json:"expired"`
		Onhold   bool `json:"onhold"`
	} `json:"flags"`
	VisibilityFlags struct {
		Bumped        bool `json:"bumped"`
		Highlighted   bool `json:"highlighted"`
		Urgent        bool `json:"urgent"`
		CountryBumped bool `json:"country_bumped"`
		Boosted       bool `json:"boosted"`
	} `json:"visibility_flags"`
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
	SupportsShipping bool   `json:"supports_shipping"`
	ShippingAllowed  bool   `json:"shipping_allowed"`
	SellerID         string `json:"seller_id"`
	Favorited        bool   `json:"favorited"`
	CreationDate     int64  `json:"creation_date"`
	ModificationDate int64  `json:"modification_date"`
	Location         struct {
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		CountryCode string `json:"country_code"`
	} `json:"location"`
}

// Объявление структур
var (
	CarsCat    []Cars
	ElecCat    []Electronic
	HomeCat    []Home
	MotosCat   []Motos
	PhotoTVCat []PhotoTV
)

// Основная универсальная функция для работы со всеми категориями
func FindAllInCategory(url string, urlSlug string, category string) error {
	var all AllData

	resp, err := http.Get(url)
	if err != nil {
		logrus.Warnf("Error request to wallapop data - %s", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Warnf("Problems with receiving data from a resource - %d", resp.StatusCode)
		time.Sleep(time.Second * 10)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Warnf("Err parce resp body - %s", err)
		return err
	}

	if err := gojson.Unmarshal(body, &all); err != nil {
		logrus.Errorf("Err unmarshal json to struct - %s", err)
		return err
	}

	for _, val := range all.SearchObjects {
		switch category {
		case "Cars":
			CarsCat = append(CarsCat, Cars{
				ID:               val.ID,
				Title:            val.Title,
				Description:      val.Description,
				Distance:         val.Distance,
				Images:           val.Images,
				User:             val.User,
				Flags:            val.Flags,
				VisibilityFlags:  val.VisibilityFlags,
				Price:            val.Price,
				Currency:         val.Currency,
				FreeShipping:     val.FreeShipping,
				WebSlug:          fmt.Sprintf("%s/%s", urlSlug, val.WebSlug),
				CategoryID:       val.CategoryID,
				Shipping:         val.Shipping,
				SupportsShipping: val.SupportsShipping,
				ShippingAllowed:  val.ShippingAllowed,
				SellerID:         val.SellerID,
				Favorited:        val.Favorited,
				CreationDate:     val.CreationDate,
				ModificationDate: val.ModificationDate,
				Location:         val.Location,
			})
		case "Electronic":
			ElecCat = append(ElecCat, Electronic{
				ID:               val.ID,
				Title:            val.Title,
				Description:      val.Description,
				Distance:         val.Distance,
				Images:           val.Images,
				User:             val.User,
				Flags:            val.Flags,
				VisibilityFlags:  val.VisibilityFlags,
				Price:            val.Price,
				Currency:         val.Currency,
				FreeShipping:     val.FreeShipping,
				WebSlug:          fmt.Sprintf("%s/%s", urlSlug, val.WebSlug),
				CategoryID:       val.CategoryID,
				Shipping:         val.Shipping,
				SupportsShipping: val.SupportsShipping,
				ShippingAllowed:  val.ShippingAllowed,
				SellerID:         val.SellerID,
				Favorited:        val.Favorited,
				CreationDate:     val.CreationDate,
				ModificationDate: val.ModificationDate,
				Location:         val.Location,
			})
		case "Home":
			HomeCat = append(HomeCat, Home{
				ID:               val.ID,
				Title:            val.Title,
				Description:      val.Description,
				Distance:         val.Distance,
				Images:           val.Images,
				User:             val.User,
				Flags:            val.Flags,
				VisibilityFlags:  val.VisibilityFlags,
				Price:            val.Price,
				Currency:         val.Currency,
				FreeShipping:     val.FreeShipping,
				WebSlug:          fmt.Sprintf("%s/%s", urlSlug, val.WebSlug),
				CategoryID:       val.CategoryID,
				Shipping:         val.Shipping,
				SupportsShipping: val.SupportsShipping,
				ShippingAllowed:  val.ShippingAllowed,
				SellerID:         val.SellerID,
				Favorited:        val.Favorited,
				CreationDate:     val.CreationDate,
				ModificationDate: val.ModificationDate,
				Location:         val.Location,
			})
		case "Motos":
			MotosCat = append(MotosCat, Motos{
				ID:               val.ID,
				Title:            val.Title,
				Description:      val.Description,
				Distance:         val.Distance,
				Images:           val.Images,
				User:             val.User,
				Flags:            val.Flags,
				VisibilityFlags:  val.VisibilityFlags,
				Price:            val.Price,
				Currency:         val.Currency,
				FreeShipping:     val.FreeShipping,
				WebSlug:          fmt.Sprintf("%s/%s", urlSlug, val.WebSlug),
				CategoryID:       val.CategoryID,
				Shipping:         val.Shipping,
				SupportsShipping: val.SupportsShipping,
				ShippingAllowed:  val.ShippingAllowed,
				SellerID:         val.SellerID,
				Favorited:        val.Favorited,
				CreationDate:     val.CreationDate,
				ModificationDate: val.ModificationDate,
				Location:         val.Location,
			})
		case "PhotoTV":
			PhotoTVCat = append(PhotoTVCat, PhotoTV{
				ID:               val.ID,
				Title:            val.Title,
				Description:      val.Description,
				Distance:         val.Distance,
				Images:           val.Images,
				User:             val.User,
				Flags:            val.Flags,
				VisibilityFlags:  val.VisibilityFlags,
				Price:            val.Price,
				Currency:         val.Currency,
				FreeShipping:     val.FreeShipping,
				WebSlug:          fmt.Sprintf("%s/%s", urlSlug, val.WebSlug),
				CategoryID:       val.CategoryID,
				Shipping:         val.Shipping,
				SupportsShipping: val.SupportsShipping,
				ShippingAllowed:  val.ShippingAllowed,
				SellerID:         val.SellerID,
				Favorited:        val.Favorited,
				CreationDate:     val.CreationDate,
				ModificationDate: val.ModificationDate,
				Location:         val.Location,
			})
		}
	}
	return nil
}
