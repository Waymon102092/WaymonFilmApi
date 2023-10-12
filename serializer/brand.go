package serializer

import "Waymon_api/model"

type Brand struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

func BuildBrand(brand model.Brand) Brand {
	return Brand{
		Value: brand.Title,
		Text:  brand.Title,
	}
}

func BuildBrands(items []model.Brand) (brands []Brand) {
	for _, item := range items {
		brand := BuildBrand(item)
		brands = append(brands, brand)
	}
	return
}
