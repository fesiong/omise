package operations

import (
	"github.com/fesiong/omise/core"
)

type ListLinks struct {
	List
}

func (req *ListLinks) Describe() *core.Description {
	return &core.Description{
		Endpoint:    core.API,
		Method:      "GET",
		Path:        "/links",
		ContentType: "application/json",
	}
}

type CreateLink struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Multiple    bool   `json:"multiple"`
}

func (req *CreateLink) Describe() *core.Description {
	return &core.Description{
		Endpoint:    core.API,
		Method:      "POST",
		Path:        "/links",
		ContentType: "application/json",
	}
}

type RetrieveLink struct {
	LinkID string `json:"-"`
}

func (req *RetrieveLink) Describe() *core.Description {
	return &core.Description{
		Endpoint:    core.API,
		Method:      "GET",
		Path:        "/links/" + req.LinkID,
		ContentType: "application/json",
	}
}
