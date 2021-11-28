package accounts

type Spec struct {
	query      string
	parameters []interface{}
}

func GetAccountByUsernameSpec(username string) Spec {
	return Spec{
		query:      "SELECT * FROM accounts WHERE username = $1;",
		parameters: []interface{}{username},
	}
}
