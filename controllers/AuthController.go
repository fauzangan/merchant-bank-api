package controllers

import (
    "fmt"
    "merchant-bank-api/services"
)

func Login(username, password string) (*services.Customer, error) {
    customer, err := services.Login(username, password)
    if err != nil {
        return nil, err
    }
    fmt.Println("Login successful for customer:", customer.Name)
    return customer, nil
}

func Logout() {
    services.Logout()
    fmt.Println("Customer has been logged out.")
}
