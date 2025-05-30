package helper

import (
	"encoding/csv"
	"fmt"
	"os"
	"scrap/model"
)

func WriteProductsToCSV(products []model.Product, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"Url",
		"Image",
		"Name",
		"Price",
	}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}

	for _, product := range products {
		record := []string{
			product.Url,
			product.Image,
			product.Name,
			product.Price,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write product record to CSV: %w", err)
		}
	}

	return nil
}
