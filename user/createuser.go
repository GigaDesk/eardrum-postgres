package user

import (
    
    "github.com/GigaDesk/eardrum-interfaces/user"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// creates an unverified user record
func CreateUser(s user.NewUser, Db *gorm.DB) (user.User, error) {
    // 1. Check if the phone number already exists
    phonenumberexists, err := CheckUserPhoneNumber(Db, s.GetPhoneNumber())

    if err != nil {
        // Database connection/query failure on lookup -> 500 Internal
        return nil, ErrDBLookupFailure("Failed to check user existence due to a database connection or query issue.",err)
    }

    if phonenumberexists.Unverified {
        // Business logic error: User exists but needs verification -> 409 Conflict
        return nil, ErrUserConflict("User phone number already exists but is unverified. Please complete verification.")
    }

    if phonenumberexists.Verified {
        // Business logic error: User fully exists -> 409 Conflict
        return nil, ErrUserConflict("User phone number already exists and is verified.")
    }
    
    // 2. Create unverified user data
    // (Omitted fields for brevity, assuming UnverifiedUser struct exists)
    unverifieduser := &UnverifiedUser{
        UserName: s.GetUserName(),
        PhoneNumber: s.GetPhoneNumber(),
        Password: s.GetPassword(),
        MpesaNumber: s.GetMpesaNumber(),
    }

    // Generate a new UUID and assign it to the QrCode field
    unverifieduser.QrCode = uuid.New()

    // 3. Create record in the database
    if err := Db.Create(unverifieduser).Error; err != nil {
        
        // A. Check for specific, recoverable DB errors (like Unique Constraint Violation)
        if isUniqueConstraintViolation(err) {
            // This is a known conflict (e.g., QR code or phone number violation) -> 409 Conflict
            return nil, ErrUserConflict("Phone Number is already in use.")
        }

        // B. All other unexpected DB errors -> 500 Internal
        // The helper wraps the raw 'err' to preserve it for logging.
        return nil, ErrDBPersistenceFailure(err)
    }

    return unverifieduser, nil
}