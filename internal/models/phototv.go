package models

/*
	https://api.wallapop.com/api/v3/general/search?
		user_province=Madrid
		&latitude=40.41956
		&start=200
		&user_region=Comunidad+de+Madrid
		&user_city=Madrid
		&search_id=1e5fb024-9140-495d-92b5-47710871c000
		&country_code=ES
		&user_postal_code=28014
		&experiment=freshness_factor_variation_D
		&items_count=200
		&density_type=20
		&filters_source=seo_landing
		&order_by=closest&step=0
		&category_ids=12545
		&longitude=-3.69196
*/

type PhotoTvCategory struct {
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
		WebSlug      string  `json:"web_slug"` // https://es.wallapop.com/item/{web-slug}
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
