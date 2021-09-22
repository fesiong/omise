package operations

import (
	"github.com/fesiong/omise/core"
)

// Example:
//
//	balance := &omise.Balance{}
//	if e := client.Do(balance, &RetrieveBalance{}); e != nil {
//		panic(e)
//	}
//
//	fmt.Println("we have $", balance.Available)
//
type RetrieveBalance struct{}

func (req *RetrieveBalance) Describe() *core.Description {
	return &core.Description{
		Endpoint:    core.API,
		Method:      "GET",
		Path:        "/balance",
		ContentType: "application/json",
	}
}
