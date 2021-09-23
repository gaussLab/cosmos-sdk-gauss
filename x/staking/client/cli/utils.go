package cli

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr string) (commission types.CommissionRates, err error) {
	if rateStr == "" || maxRateStr == "" || maxChangeRateStr == "" {
		return commission, errors.New("must specify all validator commission parameters")
	}

	rate, err := sdk.NewDecFromStr(rateStr)
	if err != nil {
		return commission, err
	}

	maxRate, err := sdk.NewDecFromStr(maxRateStr)
	if err != nil {
		return commission, err
	}

	maxChangeRate, err := sdk.NewDecFromStr(maxChangeRateStr)
	if err != nil {
		return commission, err
	}

	commission = types.NewCommissionRates(rate, maxRate, maxChangeRate)

	return commission, nil
}

func buildCommissionReallocation(reserveRateStr, reallocatedRateStr string) (commissionReallocation types.CommissionReallocation, err error) {
	if reserveRateStr == "" || reallocatedRateStr == "" {
		return commissionReallocation, errors.New("must specify all validator commission-reallocation parameters")
	}

	reserveRate, err := sdk.NewDecFromStr(reserveRateStr)
	if err != nil {
		return commissionReallocation, err
	}

	reallocatedRate, err := sdk.NewDecFromStr(reallocatedRateStr)
	if err != nil {
		return commissionReallocation, err
	}

	commissionReallocation = types.NewCommissionReallocation(reserveRate, reallocatedRate)

	return commissionReallocation, nil
}
