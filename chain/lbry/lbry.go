package lbry

import (
	"github.com/btcsuite/btcd/chaincfg"
)

func init() {
	if err := chaincfg.Register(&MainNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&TestNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&RegressionNetParams); err != nil {
		panic(err)
	}
}

// MainNetParams returns the chain configuration for mainnet.
var MainNetParams = chaincfg.Params{
	Name:        "mainnet",
	Net:         0xfae4aaf1,
	DefaultPort: "9246",

	// Address encoding magics
	PubKeyHashAddrID: 85,
	ScriptHashAddrID: 122,
	PrivateKeyID:     28,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "lbc",
}

// TestNetParams returns the chain configuration for testnet.
var TestNetParams = chaincfg.Params{
	Name:        "testnet",
	Net:         0xfae4aae1,
	DefaultPort: "19246",

	// Address encoding magics
	PubKeyHashAddrID: 111,
	ScriptHashAddrID: 196,
	PrivateKeyID:     239,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with xpub

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "tlbc",
}

// RegressionNetParams returns the chain configuration for regression net.
var RegressionNetParams = chaincfg.Params{
	Name: "regtest",

	// Litecoin has 0xdab5bffa as RegTest (same as Bitcoin's RegTest).
	// Setting it to an arbitrary value (leet_hex(litecoin)), so that we can
	// register the regtest network.
	Net: 0xfae4aad1,

	// Address encoding magics
	PubKeyHashAddrID: 111,
	ScriptHashAddrID: 196,
	PrivateKeyID:     239,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with xpub

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "rlbc",
}
