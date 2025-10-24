package user

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

// User represents the customer model for the system.
type User struct {
    gorm.Model
    UserName              string `gorm:"not null"` // UserName of the user
    PhoneNumber           string `gorm:"uniqueIndex;not null"` // Phone number of the user
    Password              string `gorm:"not null"` // The user's password. It should be stored as a secure hash.
    AccountBalanceInCents uint   `gorm:"not null;default:0"`  // The user's account balance in cents
    PinCode               *string   `gorm:""` // The user's security PIN code (can be NULL). Must be stored as a secure hash.
    MpesaNumber           *string   `gorm:""` // The user's M-Pesa number used for withdrawals (can be NULL)
	QrCode                uuid.UUID `gorm:"uniqueIndex;type:uuid"` // The user's unique Qr Code 
}

// Returns the unique ID of the user
func (u User) GetID() int64 {
	return int64(u.ID)
}

// Returns the creation timestamp of the user
func (u User) GetCreatedAt() time.Time {
	return u.CreatedAt.UTC()
}

// Returns the update timestamp of the user
func (u User) GetUpdatedAt() time.Time {
	return u.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the user
func (u User) GetDeletedAt() time.Time {
	return u.DeletedAt.Time.UTC()
}

// Returns the username of the user
func (u User) GetUserName() string {
	return u.UserName
}

// Returns the phone number of the user
func (u User) GetPhoneNumber() string {
	return u.PhoneNumber
}

// Returns the m-pesa number of the user used for withdrawals
func (u User) GetMpesaNumber() *string {
	return u.MpesaNumber
}

// Returns the account balance in cents of the user
func (u User) GetAccountBalanceInCents() int64 {
	return int64(u.AccountBalanceInCents)
}

// Returns the security password of the user
func (u User) GetPassword() string {
	return u.Password
}

// Returns the security PIN code of the user
func (u User) GetPinCode() *string {
	return u.PinCode
}

// GetQrCodeUUID returns the UUID stored in the QrCode field.
func (u User) GetQrCodeUUID() uuid.UUID {
    return u.QrCode
}

// GetQrCodeBase64 returns the Base64 string of the QR code image.
// It generates a QR code image from the user's UUID and encodes it.
func (u User) GetQrCodeBase64() string {
    // Convert the UUID to a string
    uuidStr := u.QrCode.String()

    // Generate the QR code image as a PNG byte slice
    // The size is set to 256 for a clear image.
    png, err := qrcode.Encode(uuidStr, qrcode.Medium, 256)
    if err != nil {
        log.Printf("Error generating QR code for UUID %s: %v", uuidStr, err)
        return "" // Return an empty string or handle the error as needed
    }

    // Encode the PNG byte slice to a Base64 string
    base64Str := base64.StdEncoding.EncodeToString(png)
    return base64Str
}


// UnverifiedUser represents the unverified customer model for the system.
type UnverifiedUser struct {
    gorm.Model
    UserName              string `gorm:"not null"` // UserName of the user
    PhoneNumber           string `gorm:"uniqueIndex;not null"` // Phone number of the user
    Password              string `gorm:"not null"` // The user's password. It should be stored as a secure hash.
    AccountBalanceInCents uint   `gorm:"not null;default:0"`  // The user's account balance in cents
    PinCode               *string   `gorm:""` // The user's security PIN code (can be NULL). Must be stored as a secure hash.
    MpesaNumber           *string   `gorm:""` // The user's M-Pesa number used for withdrawals (can be NULL)
	QrCode                uuid.UUID `gorm:"uniqueIndex;type:uuid"` // The user's unique Qr Code 
}

// Returns the unique ID of the unverified user
func (u UnverifiedUser) GetID() int64 {
	return int64(u.ID)
}

// Returns the creation timestamp of the unverified user
func (u UnverifiedUser) GetCreatedAt() time.Time {
	return u.CreatedAt.UTC()
}

// Returns the update timestamp of the unverified user
func (u UnverifiedUser) GetUpdatedAt() time.Time {
	return u.UpdatedAt.UTC()
}

// Returns the deletion timestamp of the unverified user
func (u UnverifiedUser) GetDeletedAt() time.Time {
	return u.DeletedAt.Time.UTC()
}

// Returns the username of the unverified user
func (u UnverifiedUser) GetUserName() string {
	return u.UserName
}

// Returns the phone number of the unverified user
func (u UnverifiedUser) GetPhoneNumber() string {
	return u.PhoneNumber
}

// Returns the m-pesa number of the unverified user used for withdrawals
func (u UnverifiedUser) GetMpesaNumber() *string {
	return u.MpesaNumber
}

// Returns the account balance in cents of the unverified user
func (u UnverifiedUser) GetAccountBalanceInCents() int64 {
	return int64(u.AccountBalanceInCents)
}

// Returns the security password of the unverified user
func (u UnverifiedUser) GetPassword() string {
	return u.Password
}

// Returns the security PIN code of the unverified user
func (u UnverifiedUser) GetPinCode() *string {
	return u.PinCode
}

// GetQrCodeUUID returns the UUID stored in the QrCode field.
func (u UnverifiedUser) GetQrCodeUUID() uuid.UUID {
    return u.QrCode
}

// GetQrCodeBase64 returns the Base64 string of the QR code image.
// It generates a QR code image from the user's UUID and encodes it.
func (u UnverifiedUser) GetQrCodeBase64() string {
    // Convert the UUID to a string
    uuidStr := u.QrCode.String()

    // Generate the QR code image as a PNG byte slice
    // The size is set to 256 for a clear image.
    png, err := qrcode.Encode(uuidStr, qrcode.Medium, 256)
    if err != nil {
        log.Printf("Error generating QR code for UUID %s: %v", uuidStr, err)
        return "" // Return an empty string or handle the error as needed
    }

    // Encode the PNG byte slice to a Base64 string
    base64Str := base64.StdEncoding.EncodeToString(png)
    return base64Str
}

type PhoneNumberExists struct {
	Verified   bool
	Unverified bool
}
