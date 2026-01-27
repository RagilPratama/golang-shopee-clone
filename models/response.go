package models

type FeaturesResponse struct {
	Data []Feature `json:"data"`
}

type FeatureResponse struct {
	Data Feature `json:"data"`
}

type MenusResponse struct {
	Data []Menu `json:"data"`
}

type MenuResponse struct {
	Data Menu `json:"data"`
}

type ProductsResponse struct {
	Data []Product `json:"data"`
}

type ProductResponse struct {
	Data Product `json:"data"`
}
