package controllers

import (
    "fmt"
    "merchant-bank-api/services"
)

func Payment(customerID string, amount float64) error {
    err := services.Payment(customerID, amount)
    if err != nil {
        return err
    }
    fmt.Println("Payment of", amount, "was successful for customer:", customerID)
    return nil
}
