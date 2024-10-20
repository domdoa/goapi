package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"dom": {
		Username: "dom",
		Token: "123",
	},
	"john": {
		Username: "john",
		Token: "456",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"dom": {
		Coins: 100,
		Username: "dom",
	},
	"john": {
		Coins: 200,
		Username: "john",
	},
}

func (d *mockDB) GetLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}