package services

import (
    "errors"
    "merchant-bank-api/models"
    "merchant-bank-api/repositories"
    "merchant-bank-api/utils"
    "time"
)

var loggedInCustomer *models.Customer

func Login(username, password string) (*models.Customer, error) {
    customers, err := repositories.GetCustomers()
    if err != nil {
        return nil, err
    }

    for _, customer := range customers {
        if customer.Name == username && customer.Password == password {
            loggedInCustomer = &customer
            
            // Log aktivitas login
            logEntry := models.History{
                ID:        "log-" + time.Now().Format("20060102150405"),
                CustomerID: customer.ID,
                Action:    "login",
                Timestamp: time.Now(),
            }
            repositories.AddHistory(logEntry)
            utils.LogInfo("Customer " + customer.Name + " logged in.")

            return &customer, nil
        }
    }

    return nil, errors.New("customer not found or incorrect password")
}

func Logout() {
    if loggedInCustomer != nil {
        // Log aktivitas logout
        logEntry := models.History{
            ID:        "log-" + time.Now().Format("20060102150405"),
            CustomerID: loggedInCustomer.ID,
            Action:    "logout",
            Timestamp: time.Now(),
        }
        repositories.AddHistory(logEntry)
        utils.LogInfo("Customer " + loggedInCustomer.Name + " logged out.")
        
        loggedInCustomer = nil
    }
}
