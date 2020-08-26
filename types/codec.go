package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDefineService{}, "irismod/service/MsgDefineService", nil)
	cdc.RegisterConcrete(&MsgBindService{}, "irismod/service/MsgBindService", nil)
	cdc.RegisterConcrete(&MsgUpdateServiceBinding{}, "irismod/service/MsgUpdateServiceBinding", nil)
	cdc.RegisterConcrete(&MsgSetWithdrawAddress{}, "irismod/service/MsgSetWithdrawAddress", nil)
	cdc.RegisterConcrete(&MsgDisableServiceBinding{}, "irismod/service/MsgDisableServiceBinding", nil)
	cdc.RegisterConcrete(&MsgEnableServiceBinding{}, "irismod/service/MsgEnableServiceBinding", nil)
	cdc.RegisterConcrete(&MsgRefundServiceDeposit{}, "irismod/service/MsgRefundServiceDeposit", nil)

	cdc.RegisterConcrete(&MsgCallService{}, "irismod/service/MsgCallService", nil)
	cdc.RegisterConcrete(&MsgRespondService{}, "irismod/service/MsgRespondService", nil)
	cdc.RegisterConcrete(&MsgPauseRequestContext{}, "irismod/service/MsgPauseRequestContext", nil)
	cdc.RegisterConcrete(&MsgStartRequestContext{}, "irismod/service/MsgStartRequestContext", nil)
	cdc.RegisterConcrete(&MsgKillRequestContext{}, "irismod/service/MsgKillRequestContext", nil)
	cdc.RegisterConcrete(&MsgUpdateRequestContext{}, "irismod/service/MsgUpdateRequestContext", nil)
	cdc.RegisterConcrete(&MsgWithdrawEarnedFees{}, "irismod/service/MsgWithdrawEarnedFees", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDefineService{},
		&MsgBindService{},
		&MsgUpdateServiceBinding{},
		&MsgSetWithdrawAddress{},
		&MsgDisableServiceBinding{},
		&MsgEnableServiceBinding{},
		&MsgRefundServiceDeposit{},
		&MsgCallService{},
		&MsgRespondService{},
		&MsgPauseRequestContext{},
		&MsgStartRequestContext{},
		&MsgKillRequestContext{},
		&MsgUpdateRequestContext{},
		&MsgWithdrawEarnedFees{},
	)
}

var (
	amino = codec.New()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
