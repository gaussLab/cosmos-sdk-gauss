package types

import (
	"time"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewCommissionReallocation returns an initialized validator commission-reallocation.
func NewCommissionReallocation(reserveRate, reallocatedRate sdk.Dec) CommissionReallocation {
	return CommissionReallocation{
		ReserveRate:     reserveRate,
		ReallocatedRate:  reallocatedRate,
		UpdateTime:      time.Unix(0, 0).UTC(),
	}
}

// NewCommissionReallocationWithTime returns an initialized validator commission-reallocation with a specified
// update time which should be the current block BFT time.
func NewCommissionReallocationWithTime(reserveRate, reallocatedRate sdk.Dec, updatedAt time.Time) CommissionReallocation {
	return CommissionReallocation{
		ReserveRate:     reserveRate,
		ReallocatedRate:  reallocatedRate,
		UpdateTime:      updatedAt,
	}
}

// String implements the Stringer interface for a CommissionReallocation object.
func (c CommissionReallocation) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}

// Validate performs basic sanity validation checks of initial commission-reallocation
// parameters. If validation fails, an SDK error is returned.
func (c CommissionReallocation) Validate() error {
	switch {
	case c.ReserveRate.IsNegative():
		// rate cannot be negative
		return ErrCommissionReallocationNegative

	case c.ReallocatedRate.IsNegative():
		// rate cannot be negative
		return ErrCommissionReallocationNegative

	case c.ReserveRate.GT(sdk.OneDec()):
		// rate cannot be greater than 1
		return ErrCommissionReallocationHuge

	case c.ReallocatedRate.GT(sdk.OneDec()):
		// rate cannot be greater than 1
		return ErrCommissionReallocationHuge
	}

	return nil
}

// ValidateNewRate performs basic sanity validation checks of a new commission-reallocation
// rate. If validation fails, an SDK error is returned.
func (c CommissionReallocation) ValidateNewRate(reserveRate sdk.Dec, blockTime time.Time) error {
	switch {
	case blockTime.Sub(c.UpdateTime).Hours() < 24:
		// new rate cannot be changed more than once within 24 hours
		return ErrCommissionReallocationUpdateTime

	case reserveRate.IsNegative():
		// new rate cannot be negative
		return ErrCommissionReallocationNegative

	case reserveRate.GT(sdk.OneDec()):
		// new rate cannot be greater than 1
		return ErrCommissionReallocationHuge
	}

	return nil
}
