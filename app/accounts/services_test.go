package accounts_test

import (
	"context"
	"testing"
	"time"
	"tracking/app/accounts"

	"github.com/stretchr/testify/assert"
)

func testHasher(password string) (string, error) {
	if len(password) > 60 {
		password = password[:61]
	}

	return password, nil
}

func testNow() time.Time {
	return time.Unix(1638134678, 0).UTC()
}

func TestGetAccountValue(t *testing.T) {
	registrationCredentials := accounts.RegistrationCredentials{
		Username:  "Daniechka",
		Password1: "test password",
		Password2: "test password",
	}

	accountValue, err := accounts.GetAccountValue(testHasher, registrationCredentials)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, accounts.AccountValue{Username: registrationCredentials.Username, Password: registrationCredentials.Password1}, accountValue)
}

func TestGetBcryptHasher(t *testing.T) {
	testPassword := "test password"
	bcryptHasher := accounts.GetBcryptHasher(10)

	passwordHash, err := bcryptHasher(testPassword)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, passwordHash, 60)
}

func TestRegisterAccount(t *testing.T) {
	ctx := context.Background()
	tx := getTestConnection(ctx, getTestSettings(), t)
	defer tx.Rollback(ctx)

	accountRepository := accounts.AccountRepository{Connection: tx, Now: testNow}

	accountValue := accounts.AccountValue{Username: "Daniechka", Password: "password hash"}

	accountEntity, err := accounts.RegisterAccount(ctx, accountRepository, accountValue)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, accounts.AccountEntity{Username: accountValue.Username, Password: accountValue.Password, CreatedAt: testNow()}, accountEntity)
}
