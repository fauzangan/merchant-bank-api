package repositories

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "merchant-bank-api/models"
    "merchant-bank-api/config"
)

func GetCustomers() ([]models.Customer, error) {
    file, err := os.Open(config.CustomerDataPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    byteValue, _ := ioutil.ReadAll(file)
    var customers []models.Customer
    json.Unmarshal(byteValue, &customers)

    return customers, nil
}
