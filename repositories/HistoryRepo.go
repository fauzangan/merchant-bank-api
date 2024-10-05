package repositories

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "merchant-bank-api/models"
)

func GetHistory() ([]models.History, error) {
    file, err := os.Open("data/history.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    byteValue, _ := ioutil.ReadAll(file)
    var history []models.History
    json.Unmarshal(byteValue, &history)

    return history, nil
}

func AddHistory(entry models.History) error {
    history, err := GetHistory()
    if err != nil {
        return err
    }

    history = append(history, entry)

    file, err := json.MarshalIndent(history, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile("data/history.json", file, 0644)
}
