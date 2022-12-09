package models

/*
	https://api.wallapop.com/api/v3/general/search?
		user_province=Madrid
		&latitude=40.41956
		&start=40
		&user_region=Comunidad+de+Madrid
		&user_city=Madrid
		&search_id=b5492bdb-1e92-4a18-bc2d-e06fe6b04501
		&country_code=ES
		&user_postal_code=28014
		&experiment=freshness_factor_variation_D
		&items_count=120&density_type=20
		&filters_source=seo_landing
		&order_by=closest&step=0
		&category_ids=12800
		&longitude=-3.69196
*/

type CarsCategory struct {
	SearchObjects []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Content struct {
			ID           string  `json:"id"`
			Title        string  `json:"title"`
			Storytelling string  `json:"storytelling"`
			Distance     float64 `json:"distance"`
			Images       []struct {
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
			Price            float64 `json:"price"`
			Currency         string  `json:"currency"`
			WebSlug          string  `json:"web_slug"` // https://es.wallapop.com/item/{web-slug}
			CategoryID       int     `json:"category_id"`
			Brand            string  `json:"brand"`
			Model            string  `json:"model"`
			Year             int     `json:"year"`
			Version          string  `json:"version"`
			Km               int     `json:"km"`
			Engine           string  `json:"engine"`
			Gearbox          string  `json:"gearbox"`
			Horsepower       float64 `json:"horsepower"`
			Favorited        bool    `json:"favorited"`
			CreationDate     int64   `json:"creation_date"`
			ModificationDate int64   `json:"modification_date"`
			Location         struct {
				City        string `json:"city"`
				PostalCode  string `json:"postal_code"`
				CountryCode string `json:"country_code"`
			} `json:"location"`
			Shipping struct {
				ItemIsShippable     bool        `json:"item_is_shippable"`
				UserAllowsShipping  bool        `json:"user_allows_shipping"`
				CostConfigurationID interface{} `json:"cost_configuration_id"`
			} `json:"shipping"`
			SupportsShipping bool `json:"supports_shipping"`
		} `json:"content,omitempty"`
	}
}
