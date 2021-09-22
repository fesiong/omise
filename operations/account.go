package operations

import (
	"github.com/fesiong/omise/core"
)

// Example:
//
//	account := &omise.Account{}
//	if e := client.Do(account, &RetrieveAccount{}); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("my account!: %#v\n", account)
//
type RetrieveAccount struct{}

func (req *RetrieveAccount) Describe() *core.Description {
	return &core.Description{
		Endpoint:    core.API,
		Method:      "GET",
		Path:        "/account",
		ContentType: "application/json",
	}
}
