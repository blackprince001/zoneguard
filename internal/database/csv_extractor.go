package database

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"

	"gorm.io/gorm"
)

func CSVtoSqlite(db *gorm.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, row := range data[1:] {
		newIp := Iplocator{}

		for i, field := range row {
			switch i {
			case 0:
				continue
			case 1:
				newIp.RelatedID, err = strconv.ParseUint(field, 10, 64)
				if err != nil {
					return fmt.Errorf("error parsing RelatedID: %w", err)
				}

				if newIp.RelatedID > math.MaxUint {
					return fmt.Errorf("RelatedID value exceeds uint range")
				}
			case 2:
				newIp.CountryCode = field
			case 3:
				newIp.CountryName = field
			case 4:
				newIp.StateProvince = field
			case 5:
				newIp.City = field
			case 6:
				newIp.Latitude, err = strconv.ParseFloat(field, 64)
				if err != nil {
					return fmt.Errorf("error parsing Latitude: %w", err)
				}
			case 7:
				newIp.Longitude, err = strconv.ParseFloat(field, 64)
				if err != nil {
					return fmt.Errorf("error parsing Longitude: %w", err)
				}
			default:
				fmt.Printf("Ignoring unexpected column: %s\n", field)
			}
		}

		err = db.Create(&newIp).Error
		if err != nil {
			return fmt.Errorf("error saving record: %w", err)
		}
	}

	return nil
}
