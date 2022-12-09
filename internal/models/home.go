package models

/*
	https://api.wallapop.com/api/v3/real_estate/search?
		filters_source=seo_landing
		&latitude=40.41956
		&start=40
		&order_by=most_relevance
		&step=0
		&category_ids=200
		&longitude=-3.69196
		&search_id=70de3634-8711-4d63-b4e6-4b5c7e828383
*/

type HomeCategory struct {
	SearchObjects []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Content struct {
			CategoryID int     `json:"category_id"`
			ID         string  `json:"id"`
			Title      string  `json:"title"`
			Distance   float64 `json:"distance"`
			Images     []struct {
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
			Price     float64 `json:"price"`
			Currency  string  `json:"currency"`
			WebSlug   string  `json:"web_slug"` // https://es.wallapop.com/item/{web-slug}
			Favorited bool    `json:"favorited"`
			Location  struct {
				City        string `json:"city"`
				PostalCode  string `json:"postal_code"`
				CountryCode string `json:"country_code"`
			} `json:"location"`
			Operation        string `json:"operation"`
			Type             string `json:"type"`
			Rooms            int    `json:"rooms"`
			Bathrooms        int    `json:"bathrooms"`
			Garage           bool   `json:"garage"`
			Terrace          bool   `json:"terrace"`
			Elevator         bool   `json:"elevator"`
			Pool             bool   `json:"pool"`
			Garden           bool   `json:"garden"`
			Condition        string `json:"condition"`
			Storytelling     string `json:"storytelling"`
			CreationDate     int64  `json:"creation_date"`
			ModificationDate int64  `json:"modification_date"`
			Shipping         struct {
				ItemIsShippable     bool        `json:"item_is_shippable"`
				UserAllowsShipping  bool        `json:"user_allows_shipping"`
				CostConfigurationID interface{} `json:"cost_configuration_id"`
			} `json:"shipping"`
			SupportsShipping bool `json:"supports_shipping"`
		} `json:"content"`
	} `json:"search_objects"`
}
