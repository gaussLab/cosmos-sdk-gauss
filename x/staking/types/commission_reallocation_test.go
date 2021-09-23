package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestCommissionReallocationValidate(t *testing.T) {
	testCases := []struct {
		input     types.CommissionReallocation
		expectErr bool
	}{
		// invalid commission;  < 0%
		{types.NewCommissionReallocation(sdk.ZeroDec(), sdk.MustNewDecFromStr("-1.00")), true},
		// invalid commission;  > 100%
		{types.NewCommissionReallocation(sdk.ZeroDec(), sdk.MustNewDecFromStr("2.00")), true},
		// invalid commission;  < 0%
		{types.NewCommissionReallocation(sdk.MustNewDecFromStr("-1.00"), sdk.ZeroDec()), true},
		// invalid commission;  > 100%
		{types.NewCommissionReallocation(sdk.MustNewDecFromStr("2.00"), sdk.ZeroDec()), true},
		// valid commission; 
		{types.NewCommissionReallocation(sdk.OneDec(), sdk.OneDec()), false},
		// valid commission; 
		{types.NewCommissionReallocation(sdk.ZeroDec(), sdk.ZeroDec()), false},
		// valid commission
		{types.NewCommissionReallocation(sdk.MustNewDecFromStr("0.20"), sdk.OneDec()), false},
	}

	for i, tc := range testCases {
		err := tc.input.Validate()
		require.Equal(t, tc.expectErr, err != nil, "unexpected result; tc #%d, input: %v", i, tc.input)
	}
}

func TestCommissionReallocationValidateNewRate(t *testing.T) {
	now := time.Now().UTC()
	c1 := types.NewCommissionReallocation(sdk.MustNewDecFromStr("0.40"), sdk.MustNewDecFromStr("0.80"))
	c1.UpdateTime = now

	testCases := []struct {
		input     types.CommissionReallocation
		newRate   sdk.Dec
		blockTime time.Time
		expectErr bool
	}{
		// invalid new commission rate; last update < 24h ago
		{c1, sdk.MustNewDecFromStr("0.50"), now, true},
		// invalid new commission rate; new rate < 0%
		{c1, sdk.MustNewDecFromStr("-1.00"), now.Add(48 * time.Hour), true},
		// invalid new commission rate; new rate > max rate
		{c1, sdk.MustNewDecFromStr("0.90"), now.Add(48 * time.Hour), true},
		// invalid new commission rate; new rate > max change rate
		{c1, sdk.MustNewDecFromStr("0.60"), now.Add(48 * time.Hour), true},
		// valid commission
		{c1, sdk.MustNewDecFromStr("0.50"), now.Add(48 * time.Hour), false},
		// valid commission
		{c1, sdk.MustNewDecFromStr("0.10"), now.Add(48 * time.Hour), false},
	}

	for i, tc := range testCases {
		err := tc.input.ValidateNewRate(tc.newRate, tc.blockTime)
		require.Equal(
			t, tc.expectErr, err != nil,
			"unexpected result; tc #%d, input: %v, newRate: %s, blockTime: %s",
			i, tc.input, tc.newRate, tc.blockTime,
		)
	}
}
