package operations_test

import (
	"testing"

	omise "github.com/fesiong/omise"
	"github.com/fesiong/omise/core/testutil"
	"github.com/fesiong/omise/operations"
	. "github.com/fesiong/omise/operations"
	r "github.com/stretchr/testify/require"
)

func TestRetrieveSource(t *testing.T) {
	const (
		SourceID = "src_test_5a444nhh27tlyv81u40"
	)
	client := testutil.NewFixedClient(t)

	source := &omise.Source{}
	client.MustDo(source, &RetrieveSource{SourceID: SourceID})
	r.Equal(t, SourceID, source.ID)
}

func TestCreateSource(t *testing.T) {
	const (
		SourceID = "src_test_5mygxph6d55vvy8nn9i"
	)
	client := testutil.NewFixedClient(t)

	exampleSource, createSource := &omise.Source{}, &operations.CreateSource{
		Type:     "fpx",
		Amount:   2000,
		Currency: "myr",
		Email:    "example@omise.co",
		Bank:     "ocbc",
	}

	client.MustDo(exampleSource, createSource)
	r.Equal(t, SourceID, exampleSource.ID)
}
