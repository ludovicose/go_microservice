package dbclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/callistaenterprise/goblog/accountservice/model"
)

type MockBoltClient struct {
	mock.Mock
}


// From here, we'll declare three functions that makes our MockBoltClient fulfill the interface IBoltClient that we declared in part 3.
func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
	args := m.Mock.Called(accountId)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoldDB() {
	// Does nothing
}

func (m *MockBoltClient) Seed() {
	// Does nothing
}