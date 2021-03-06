package operations_test

import (
	"testing"

	omise "github.com/fesiong/omise"
	"github.com/fesiong/omise/core/testutil"
	. "github.com/fesiong/omise/operations"
	r "github.com/stretchr/testify/assert"
)

func TestCapability(t *testing.T) {
	client := testutil.NewFixedClient(t)

	capability := &omise.Capability{}
	client.MustDo(capability, &RetrieveCapability{})
	r.Equal(t, "capability", capability.Object)

	examplePaymentMethod := capability.PaymentMethods[0]
	r.Equal(t, "card", examplePaymentMethod.Name)
	r.Equal(t, []string{"THB", "JPY", "MYR"}, examplePaymentMethod.Currencies)
}
