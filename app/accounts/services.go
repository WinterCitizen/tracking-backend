package accounts

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type hasher func(string) (string, error)

func GetBcryptHasher(cost int) hasher {
	return func(password string) (string, error) {
		passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			return "", err
		}

		return string(passwordHashBytes), nil
	}
}

func GetAccountValue(hashPassword hasher, registrationCredentials RegistrationCredentials) (AccountValue, error) {
	if registrationCredentials.Password1 != registrationCredentials.Password2 {
		return AccountValue{}, errors.New("passwords don't match")
	}

	passwordHashBytes, err := hashPassword(registrationCredentials.Password1)
	if err != nil {
		return AccountValue{}, err
	}

	return AccountValue{Username: registrationCredentials.Username, Password: passwordHashBytes}, nil
}

func RegisterAccount(ctx context.Context, accountRepository AccountRepository, accountValue AccountValue) (AccountEntity, error) {
	return accountRepository.Add(ctx, accountValue)
}
