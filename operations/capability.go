package operations

import "github.com/fesiong/omise/core"

// Example: Retrieve Capability
//
//	capability, retrieve := &omise.RetrieveCapability{}, &RetrieveCapability{}
//	if e := client.Do(capability, retrieve); e != nil {
//		panic(e)
//	}
//
//	fmt.Printf("list payment methods: %#v\n", capability.PaymentMethods)
//
type RetrieveCapability struct{}

func (req *RetrieveCapability) Describe() *core.Description {
	return &core.Description{
		Endpoint:    core.API,
		Method:      "GET",
		Path:        "/capability",
		ContentType: "application/json",
	}
}
