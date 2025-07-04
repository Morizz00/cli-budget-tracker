package data

import (
	"encoding/json"
	"os"

	"github.com/Morizz00/budget-tracker/model"
)

const filename = "data.json"

func saveEntry(e model.Entry) error {
	var entries []model.Entry
	data, err := os.ReadFile(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if err := json.Unmarshal(data, &entries); err != nil {
			return err
		}
	}
	entries = append(entries, e)
	jsonData, err := json.MarshalIndent(entries, "", "  ")

	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

func loadEntries() ([]model.Entry, error) {
	var entries []model.Entry
	data, err := os.ReadFile(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return []model.Entry{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
