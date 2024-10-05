package tests

import (
    "merchant-bank-api/services"
    "testing"
)

func TestPaymentNegativeAmount(t *testing.T) {
    services.Login("user1", "password123")
    err := services.Payment("1", -500)
    if err == nil {
        t.Errorf("Expected error for negative amount, but payment succeeded")
    }
}

func TestPaymentWithoutLogin(t *testing.T) {
    services.Logout() // Pastikan tidak ada customer yang login
    err := services.Payment("1", 1000)
    if err == nil {
        t.Errorf("Expected error for no logged-in customer, but payment succeeded")
    }
}
