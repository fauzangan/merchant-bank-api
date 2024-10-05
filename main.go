package main

import (
    "fmt"
    "merchant-bank-api/controllers"
)

func main() {
    fmt.Println("Merchant Bank API")
    
    // Simulate user login
    customer, err := controllers.Login("user1", "password123")
    if err != nil {
        fmt.Println("Login failed:", err)
        return
    }
    fmt.Println("Login successful:", customer)

    // Simulate payment process
    err = controllers.Payment(customer.ID, 1000)
    if err != nil {
        fmt.Println("Payment failed:", err)
        return
    }
    fmt.Println("Payment successful")

    // Logout the customer
    controllers.Logout()
    fmt.Println("Customer logged out.")
}
