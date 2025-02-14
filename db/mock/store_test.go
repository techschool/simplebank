package mockdb

import (
    "context"
    "testing"
    "github.com/golang/mock/gomock"
    "github.com/techschool/simplebank/db/sqlc"
    "fmt"
    "time"
    "github.com/google/uuid"
)


// Test generated using Keploy
func TestAddAccountBalance_ValidParams_ReturnsAccount(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.AddAccountBalanceParams{
        Amount: 100,
        ID:     1,
    }
    expectedAccount := db.Account{
        ID:      1,
        Balance: 1000,
    }

    mockStore.EXPECT().
        AddAccountBalance(ctx, params).
        Return(expectedAccount, nil)

    account, err := mockStore.AddAccountBalance(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if account != expectedAccount {
        t.Errorf("expected %v, got %v", expectedAccount, account)
    }
}

// Test generated using Keploy
func TestGetAccount_InvalidID_ReturnsError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    invalidAccountID := int64(-1)

    mockStore.EXPECT().
        GetAccount(ctx, invalidAccountID).
        Return(db.Account{}, fmt.Errorf("invalid account ID"))

    _, err := mockStore.GetAccount(ctx, invalidAccountID)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}


// Test generated using Keploy
func TestCreateUser_ValidParams_ReturnsUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.CreateUserParams{
        Username:       "testuser",
        HashedPassword: "hashedpassword",
        FullName:       "Test User",
        Email:          "test@example.com",
    }
    expectedUser := db.User{
        Username:       "testuser",
        HashedPassword: "hashedpassword",
        FullName:       "Test User",
        Email:          "test@example.com",
    }

    mockStore.EXPECT().
        CreateUser(ctx, params).
        Return(expectedUser, nil)

    user, err := mockStore.CreateUser(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user != expectedUser {
        t.Errorf("expected %v, got %v", expectedUser, user)
    }
}


// Test generated using Keploy
func TestTransferTx_InsufficientBalance_ReturnsError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.TransferTxParams{
        FromAccountID: 1,
        ToAccountID:   2,
        Amount:        10000,
    }

    mockStore.EXPECT().
        TransferTx(ctx, params).
        Return(db.TransferTxResult{}, fmt.Errorf("insufficient balance"))

    _, err := mockStore.TransferTx(ctx, params)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}


// Test generated using Keploy
func TestDeleteAccount_ValidID_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    validAccountID := int64(1)

    mockStore.EXPECT().
        DeleteAccount(ctx, validAccountID).
        Return(nil)

    err := mockStore.DeleteAccount(ctx, validAccountID)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}


// Test generated using Keploy
func TestCreateSession_ValidParams_ReturnsSession(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    sessionID := uuid.New()
    params := db.CreateSessionParams{
        ID:           sessionID,
        Username:     "testuser",
        RefreshToken: "sometoken",
        UserAgent:    "testagent",
        ClientIp:     "127.0.0.1",
        IsBlocked:    false,
        ExpiresAt:    time.Now().Add(24 * time.Hour),
    }
    expectedSession := db.Session{
        ID:           sessionID,
        Username:     "testuser",
        RefreshToken: "sometoken",
        UserAgent:    "testagent",
        ClientIp:     "127.0.0.1",
        IsBlocked:    false,
        ExpiresAt:    params.ExpiresAt,
        CreatedAt:    time.Now(),
    }

    mockStore.EXPECT().
        CreateSession(ctx, params).
        Return(expectedSession, nil)

    session, err := mockStore.CreateSession(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if session != expectedSession {
        t.Errorf("expected %v, got %v", expectedSession, session)
    }
}


// Test generated using Keploy
func TestListAccounts_ValidParams_ReturnsAccounts(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.ListAccountsParams{
        Owner:  "testowner",
        Limit:  10,
        Offset: 0,
    }
    expectedAccounts := []db.Account{
        {
            ID:       1,
            Owner:    "testowner",
            Balance:  1000,
            Currency: "USD",
        },
        {
            ID:       2,
            Owner:    "testowner",
            Balance:  2000,
            Currency: "USD",
        },
    }

    mockStore.EXPECT().
        ListAccounts(ctx, params).
        Return(expectedAccounts, nil)

    accounts, err := mockStore.ListAccounts(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(accounts) != len(expectedAccounts) {
        t.Errorf("expected %d accounts, got %d", len(expectedAccounts), len(accounts))
    }
    for i, account := range accounts {
        if account != expectedAccounts[i] {
            t.Errorf("expected %v, got %v", expectedAccounts[i], account)
        }
    }
}


// Test generated using Keploy
func TestGetEntry_InvalidID_ReturnsError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    invalidEntryID := int64(-1)

    mockStore.EXPECT().
        GetEntry(ctx, invalidEntryID).
        Return(db.Entry{}, fmt.Errorf("invalid entry ID"))

    _, err := mockStore.GetEntry(ctx, invalidEntryID)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}


// Test generated using Keploy
func TestUpdateAccount_ValidParams_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.UpdateAccountParams{
        ID:      1,
        Balance: 1500,
    }
    expectedAccount := db.Account{
        ID:       1,
        Balance:  1500,
        Currency: "USD",
    }

    mockStore.EXPECT().
        UpdateAccount(ctx, params).
        Return(expectedAccount, nil)

    account, err := mockStore.UpdateAccount(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if account != expectedAccount {
        t.Errorf("expected %v, got %v", expectedAccount, account)
    }
}


// Test generated using Keploy
func TestCreateVerifyEmail_ValidParams_ReturnsVerifyEmail(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.CreateVerifyEmailParams{
        Username:   "testuser",
        Email:      "test@example.com",
        SecretCode: "secret123",
    }
    expectedVerifyEmail := db.VerifyEmail{
        ID:         1,
        Username:   "testuser",
        Email:      "test@example.com",
        SecretCode: "secret123",
        IsUsed:     false,
        CreatedAt:  time.Now(),
        ExpiredAt:  time.Now().Add(24 * time.Hour),
    }

    mockStore.EXPECT().
        CreateVerifyEmail(ctx, params).
        Return(expectedVerifyEmail, nil)

    verifyEmail, err := mockStore.CreateVerifyEmail(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if verifyEmail != expectedVerifyEmail {
        t.Errorf("expected %v, got %v", expectedVerifyEmail, verifyEmail)
    }
}


// Test generated using Keploy
func TestVerifyEmailTx_InvalidSecretCode_ReturnsError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.VerifyEmailTxParams{
        EmailId:    1,
        SecretCode: "invalidcode",
    }

    mockStore.EXPECT().
        VerifyEmailTx(ctx, params).
        Return(db.VerifyEmailTxResult{}, fmt.Errorf("invalid secret code"))

    _, err := mockStore.VerifyEmailTx(ctx, params)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}


// Test generated using Keploy
func TestUpdateUser_InvalidParams_ReturnsError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    invalidParams := db.UpdateUserParams{
        Username: "", // Invalid username
    }

    mockStore.EXPECT().
        UpdateUser(ctx, invalidParams).
        Return(db.User{}, fmt.Errorf("invalid parameters"))

    _, err := mockStore.UpdateUser(ctx, invalidParams)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}


// Test generated using Keploy
func TestListEntries_ValidParams_ReturnsEntries(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.ListEntriesParams{
        AccountID: 1,
        Limit:     10,
        Offset:    0,
    }
    expectedEntries := []db.Entry{
        {
            ID:        1,
            AccountID: 1,
            Amount:    100,
            CreatedAt: time.Now(),
        },
        {
            ID:        2,
            AccountID: 1,
            Amount:    -50,
            CreatedAt: time.Now(),
        },
    }

    mockStore.EXPECT().
        ListEntries(ctx, params).
        Return(expectedEntries, nil)

    entries, err := mockStore.ListEntries(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(entries) != len(expectedEntries) {
        t.Errorf("expected %d entries, got %d", len(expectedEntries), len(entries))
    }
    for i, entry := range entries {
        if entry != expectedEntries[i] {
            t.Errorf("expected %v, got %v", expectedEntries[i], entry)
        }
    }
}


// Test generated using Keploy
func TestListTransfers_ValidParams_ReturnsTransfers(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockStore := NewMockStore(ctrl)
    ctx := context.Background()
    params := db.ListTransfersParams{
        FromAccountID: 1,
        ToAccountID:   2,
        Limit:         10,
        Offset:        0,
    }
    expectedTransfers := []db.Transfer{
        {
            ID:            1,
            FromAccountID: 1,
            ToAccountID:   2,
            Amount:        100,
            CreatedAt:     time.Now(),
        },
        {
            ID:            2,
            FromAccountID: 1,
            ToAccountID:   2,
            Amount:        200,
            CreatedAt:     time.Now(),
        },
    }

    mockStore.EXPECT().
        ListTransfers(ctx, params).
        Return(expectedTransfers, nil)

    transfers, err := mockStore.ListTransfers(ctx, params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(transfers) != len(expectedTransfers) {
        t.Errorf("expected %d transfers, got %d", len(expectedTransfers), len(transfers))
    }
    for i, transfer := range transfers {
        if transfer != expectedTransfers[i] {
            t.Errorf("expected %v, got %v", expectedTransfers[i], transfer)
        }
    }
}

