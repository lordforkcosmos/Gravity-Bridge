package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.NewLegacyAmino()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterInterfaces registers the interfaces for the proto stuff
// nolint: exhaustruct
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValsetConfirm{},
		&MsgSendToEth{},
		&MsgRequestBatch{},
		&MsgConfirmBatch{},
		&MsgConfirmLogicCall{},
		&MsgSendToCosmosClaim{},
		&MsgBatchSendToEthClaim{},
		&MsgERC20DeployedClaim{},
		&MsgSetOrchestratorAddress{},
		&MsgLogicCallExecutedClaim{},
		&MsgValsetUpdatedClaim{},
		&MsgCancelSendToEth{},
		&MsgSubmitBadSignatureEvidence{},
	)

	registry.RegisterInterface(
		"polygonbridge.v1beta1.EthereumClaim",
		(*EthereumClaim)(nil),
		&MsgSendToCosmosClaim{},
		&MsgBatchSendToEthClaim{},
		&MsgERC20DeployedClaim{},
		&MsgLogicCallExecutedClaim{},
		&MsgValsetUpdatedClaim{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil), &UnhaltBridgeProposal{}, &AirdropProposal{}, &IBCMetadataProposal{})

	registry.RegisterInterface("polygonbridge.v1beta1.EthereumSigned", (*EthereumSigned)(nil), &Valset{}, &OutgoingTxBatch{}, &OutgoingLogicCall{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterCodec registers concrete types on the Amino codec
// nolint: exhaustruct
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*EthereumClaim)(nil), nil)
	cdc.RegisterConcrete(&MsgSetOrchestratorAddress{}, "polygonbridge/MsgSetOrchestratorAddress", nil)
	cdc.RegisterConcrete(&MsgValsetConfirm{}, "polygonbridge/MsgValsetConfirm", nil)
	cdc.RegisterConcrete(&MsgSendToEth{}, "polygonbridge/MsgSendToEth", nil)
	cdc.RegisterConcrete(&MsgRequestBatch{}, "polygonbridge/MsgRequestBatch", nil)
	cdc.RegisterConcrete(&MsgConfirmBatch{}, "polygonbridge/MsgConfirmBatch", nil)
	cdc.RegisterConcrete(&MsgConfirmLogicCall{}, "polygonbridge/MsgConfirmLogicCall", nil)
	cdc.RegisterConcrete(&Valset{}, "polygonbridge/Valset", nil)
	cdc.RegisterConcrete(&MsgSendToCosmosClaim{}, "polygonbridge/MsgSendToCosmosClaim", nil)
	cdc.RegisterConcrete(&MsgBatchSendToEthClaim{}, "polygonbridge/MsgBatchSendToEthClaim", nil)
	cdc.RegisterConcrete(&MsgERC20DeployedClaim{}, "polygonbridge/MsgERC20DeployedClaim", nil)
	cdc.RegisterConcrete(&MsgLogicCallExecutedClaim{}, "polygonbridge/MsgLogicCallExecutedClaim", nil)
	cdc.RegisterConcrete(&MsgValsetUpdatedClaim{}, "polygonbridge/MsgValsetUpdatedClaim", nil)
	cdc.RegisterConcrete(&OutgoingTxBatch{}, "polygonbridge/OutgoingTxBatch", nil)
	cdc.RegisterConcrete(&MsgCancelSendToEth{}, "polygonbridge/MsgCancelSendToEth", nil)
	cdc.RegisterConcrete(&OutgoingTransferTx{}, "polygonbridge/OutgoingTransferTx", nil)
	cdc.RegisterConcrete(&ERC20Token{}, "polygonbridge/ERC20Token", nil)
	cdc.RegisterConcrete(&IDSet{}, "polygonbridge/IDSet", nil)
	cdc.RegisterConcrete(&Attestation{}, "polygonbridge/Attestation", nil)
	cdc.RegisterConcrete(&MsgSubmitBadSignatureEvidence{}, "polygonbridge/MsgSubmitBadSignatureEvidence", nil)
}
