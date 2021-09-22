package operations

import (
	"github.com/fesiong/omise"
	"github.com/fesiong/omise/core"
)

// Example:
//
//	result, search := &omise.ChargeSearchResult{}, &Search{
//		Scope: omise.ChargeScope,
//		Query: "chrg_test_12345",
//	}
//	if e := client.Do(result, search); e != nil {
//		panic(e)
//	}
//
type Search struct {
	Scope   omise.SearchScope `json:"scope"`
	Query   string            `json:"query,omitempty"`
	Filters map[string]string `json:"filters,omitempty"`
	Order   omise.Ordering    `json:"order,omitempty"`
	Page    int               `json:"page,omitempty"`
	PerPage int               `json:"per_page,omitempty"`
}

func (req *Search) Describe() *core.Description {
	return &core.Description{
		Endpoint:    core.API,
		Method:      "GET",
		Path:        "/search",
		ContentType: "application/json",
	}
}
