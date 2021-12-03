package accounts_test

import (
	"context"
	"log"
	"testing"
	"tracking/app/accounts"
	"tracking/app/config"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func getTestSettings() config.Settings {
	return config.Settings{PostgresURI: "postgres://postgres@localhost:5432/test_tracking"}
}

func getTestConnection(ctx context.Context, settings config.Settings, t *testing.T) pgx.Tx {
	conn, err := pgx.Connect(ctx, settings.PostgresURI)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}

	return tx
}

func TestAccountsRepositoryAddGet(t *testing.T) {
	ctx := context.Background()
	settings := getTestSettings()
	tx := getTestConnection(ctx, settings, t)

	defer tx.Rollback(ctx)

	// Given: taльб9л9g repository & tag entity to create
	accountRepository := accounts.AccountRepository{Connection: tx, Now: testNow}
	accountValue := accounts.AccountValue{Username: "Daniechka", Password: "password hash"}

	// When: tag repository add called with tag entity
	_, err := accountRepository.Add(ctx, accountValue)
	if err != nil {
		log.Fatal(err)
	}

	// Then: tag repository get returns added tag entity
	retreivedAccountEntities, err := accountRepository.Get(ctx, accounts.GetAccountByUsernameSpec(accountValue.Username))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(retreivedAccountEntities))
}
