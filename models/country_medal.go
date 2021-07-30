package models

// CountryMedal struct
type CountryMedal struct {
	Position     int64  `json:"position"`
	Country      string `json:"country"`
	GoldMedalds  int64  `json:"goldmedalds"`
	SilverMedals int64  `json:"silvermedalds"`
	BronzeMedals int64  `json:"bronzemedalds"`
}

type CountriesMedals []CountryMedal

var Countries = CountriesMedals{
	{
		Position:     1,
		Country:      "China",
		GoldMedalds:  13,
		SilverMedals: 4,
		BronzeMedals: 5,
	}, {
		Position:     2,
		Country:      "USA",
		GoldMedalds:  12,
		SilverMedals: 6,
		BronzeMedals: 9,
	}, {
		Position:     3,
		Country:      "Japon",
		GoldMedalds:  11,
		SilverMedals: 11,
		BronzeMedals: 9,
	}, {
		Position:     4,
		Country:      "Colombia",
		GoldMedalds:  7,
		SilverMedals: 10,
		BronzeMedals: 6,
	}, {
		Position:     5,
		Country:      "Australia",
		GoldMedalds:  6,
		SilverMedals: 1,
		BronzeMedals: 9,
	},
}
