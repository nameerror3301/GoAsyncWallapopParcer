package models

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	gojson "github.com/goccy/go-json"
)

type CarsCategory struct {
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

type LastData struct {
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

var Last []LastData

func FindAllInCategory(url string, urlSlug string) error {
	var all CarsCategory

	resp, err := http.Get(url)
	if err != nil {
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
		Last = append(Last, LastData{
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
	return nil
}
