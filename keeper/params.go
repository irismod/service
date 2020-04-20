package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/irismod/service/types"
)

// ParamKeyTable for service module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&types.Params{})
}

// MaxRequestTimeout returns the maximum request timeout
func (k Keeper) MaxRequestTimeout(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyMaxRequestTimeout, &res)
	return
}

// MinDepositMultiple returns the minimum deposit multiple
func (k Keeper) MinDepositMultiple(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyMinDepositMultiple, &res)
	return
}

// MinDeposit returns the minimum deposit
func (k Keeper) MinDeposit(ctx sdk.Context) (res sdk.Coins) {
	k.paramstore.Get(ctx, types.KeyMinDeposit, &res)
	return
}

// ServiceFeeTax returns the service fee tax
func (k Keeper) ServiceFeeTax(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyServiceFeeTax, &res)
	return
}

// SlashFraction returns the slashing fraction
func (k Keeper) SlashFraction(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeySlashFraction, &res)
	return
}

// ComplaintRetrospect returns the complaint retrospect duration
func (k Keeper) ComplaintRetrospect(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyComplaintRetrospect, &res)
	return
}

// ArbitrationTimeLimit returns the arbitration time limit
func (k Keeper) ArbitrationTimeLimit(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyArbitrationTimeLimit, &res)
	return
}

// TxSizeLimit returns the tx size limit
func (k Keeper) TxSizeLimit(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyTxSizeLimit, &res)
	return
}

// GetParams gets all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MaxRequestTimeout(ctx),
		k.MinDepositMultiple(ctx),
		k.MinDeposit(ctx),
		k.ServiceFeeTax(ctx),
		k.SlashFraction(ctx),
		k.ComplaintRetrospect(ctx),
		k.ArbitrationTimeLimit(ctx),
		k.TxSizeLimit(ctx),
	)
}

// SetParams sets the params to the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}