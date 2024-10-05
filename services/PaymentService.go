package services

import (
    "errors"
    "merchant-bank-api/models"
    "merchant-bank-api/repositories"
    "time"
)

func Payment(customerID string, amount float64) error {
    if loggedInCustomer == nil {
        return errors.New("no customer is logged in")
    }
    
    if loggedInCustomer.ID != customerID {
        return errors.New("unauthorized: customer ID mismatch")
    }

    if amount <= 0 {
        return errors.New("invalid amount, must be greater than zero")
    }

    // Simulasi melakukan pembayaran tanpa batasan
    history := models.History{
        ID:        "txn-" + time.Now().Format("20060102150405"),
        CustomerID: customerID,
        Action:    "payment",
        Amount:    amount,
        Timestamp: time.Now(),
    }

    err := repositories.AddHistory(history)
    if err != nil {
        return errors.New("failed to log payment")
    }

    return nil
}
