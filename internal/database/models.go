package database

type Iplocator struct {
	ID            uint    `gorm:"primary_key" json:"id"`
	RelatedID     uint64  `json:"related_id"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	StateProvince string  `json:"state_province"`
	City          string  `json:"city"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}
