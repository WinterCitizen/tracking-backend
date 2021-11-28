package accounts

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

type Connection interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type AccountRepository struct {
	Connection Connection
	Now        func() time.Time
}

func (accountRepository AccountRepository) Add(ctx context.Context, accountValue AccountValue) (AccountEntity, error) {
	createdAccountEntity := AccountEntity{}
	row := accountRepository.Connection.QueryRow(
		ctx, "INSERT INTO accounts VALUES ($1, $2, $3) RETURNING *", accountValue.Username, accountValue.Password, accountRepository.Now(),
	)
	err := row.Scan(&createdAccountEntity.Username, &createdAccountEntity.Password, &createdAccountEntity.CreatedAt)

	if err != nil {
		return AccountEntity{}, err
	}

	createdAccountEntity.CreatedAt = createdAccountEntity.CreatedAt.UTC()
	return createdAccountEntity, nil
}

func (accountRepository AccountRepository) Get(ctx context.Context, spec Spec) ([]AccountEntity, error) {
	retreivedAccountEntities := []AccountEntity{}

	rows, err := accountRepository.Connection.Query(ctx, spec.query, spec.parameters...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		accountEntity := AccountEntity{}
		err = rows.Scan(&accountEntity.Username, &accountEntity.Password, &accountEntity.CreatedAt)
		if err != nil {
			return nil, err
		}
		retreivedAccountEntities = append(retreivedAccountEntities, accountEntity)
	}

	return retreivedAccountEntities, nil
}
