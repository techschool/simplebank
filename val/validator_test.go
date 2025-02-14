package val

import (
    "testing"



)


// Test generated using Keploy
func TestValidateString_ShortInput_Error(t *testing.T) {
    err := ValidateString("ab", 3, 10)
    if err == nil {
        t.Errorf("Expected an error for input shorter than minimum length, but got nil")
    }
}

// Test generated using Keploy
func TestValidateUsername_InvalidCharacters_Error(t *testing.T) {
    err := ValidateUsername("Invalid@Username")
    if err == nil {
        t.Errorf("Expected an error for username with invalid characters, but got nil")
    }
}


// Test generated using Keploy
func TestValidateFullName_InvalidCharacters_Error(t *testing.T) {
    err := ValidateFullName("John123")
    if err == nil {
        t.Errorf("Expected an error for full name with invalid characters, but got nil")
    }
}


// Test generated using Keploy
func TestValidatePassword_ShortPassword_Error(t *testing.T) {
    err := ValidatePassword("12345")
    if err == nil {
        t.Errorf("Expected an error for password shorter than minimum length, but got nil")
    }
}


// Test generated using Keploy
func TestValidateEmail_InvalidEmail_Error(t *testing.T) {
    err := ValidateEmail("invalid-email")
    if err == nil {
        t.Errorf("Expected an error for invalid email address, but got nil")
    }
}


// Test generated using Keploy
func TestValidateEmailId_NonPositiveInteger_Error(t *testing.T) {
    err := ValidateEmailId(-1)
    if err == nil {
        t.Errorf("Expected an error for non-positive email ID, but got nil")
    }
}


// Test generated using Keploy
func TestValidateFullName_ValidFullName_NoError(t *testing.T) {
    err := ValidateFullName("John Doe")
    if err != nil {
        t.Errorf("Expected no error for a valid full name, but got %v", err)
    }
}


// Test generated using Keploy
func TestValidateEmail_ValidEmail_NoError(t *testing.T) {
    err := ValidateEmail("example@example.com")
    if err != nil {
        t.Errorf("Expected no error for a valid email address, but got %v", err)
    }
}


// Test generated using Keploy
func TestValidateEmailId_PositiveInteger_NoError(t *testing.T) {
    err := ValidateEmailId(123)
    if err != nil {
        t.Errorf("Expected no error for a positive integer email ID, but got %v", err)
    }
}

