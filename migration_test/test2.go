package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/clients"
)

var (
	ZkSyncProvider   = "http://0.0.0.0:3050/"            // zkSync Era testnet
	EthereumProvider = "https://rpc.ankr.com/eth_goerli" // goerli testnet
)

func main() {
	// Connect to zkSync network
	client, err := clients.Dial(ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	// Connect to Ethereum network
	ethClient, err := ethclient.Dial(EthereumProvider)
	if err != nil {
		log.Panic(err)
	}
	defer ethClient.Close()

	// privateKey1 := "8844ab4ded805b0cae9d030295005582227d6edd3c64b2ce475e2d3cb72ca1e9"
	// address1 := "0x4AA86c556ca7a5a571F4Cc9979747C442FABa7C6"
	// address2 := "0x08f5f9a336Aae6A72c795Ddf307864B13d13F0Aa"

	// // Create wallet
	// wallet, err := accounts.NewWallet(common.Hex2Bytes(privateKey1), &client, nil)
	// if err != nil {
	// 	log.Println(err)
	// }

	// balance, err := wallet.Balance(context.Background(), utils.EthAddress, nil)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println("Balance antes del transfer: ", balance)

	// // Perform transfer
	// tx, err := wallet.Transfer(nil, accounts.TransferTransaction{
	// 	To:     common.HexToAddress(address2),
	// 	Amount: big.NewInt(1_000_000_000_000_000_000),
	// 	Token:  utils.EthAddress,
	// })
	// if err != nil {
	// 	log.Panic(err)
	// // }
	// log.Println("Transaction: ", tx.Hash())

	// transaction_info
	// "hash": "0x45bac28c64b986182a57e9f8128ba97e4c8e0bbfbc100e4f3f9149663d8ad743"
	// "nonce":"0x0"
	// "blockHash":"0x0c58c1d7fda87ca0d5dc963e3a81301fa60489d5f283a1fdeb03d2a59dd28a29"
	// "blockNumber":"0xe"
	// "transactionIndex":"0x0"
	// "from":"0x4aa86c556ca7a5a571f4cc9979747c442faba7c6"
	// "to":"0x08f5f9a336aae6a72c795ddf307864b13d13f0aa"
	// "value":"0xde0b6b3a7640000"
	// "gasPrice":"0xee6b280"
	// "gas":"0x29cd9"
	// "input":"0x"
	// "v":"0x1"
	// "r":"0x7b83ef6f81a8eb124c84b7d80b133600c89c9bd67aa565d4024a849d759da56b"
	// "s":"0x762462c7880aa0b713583c9096b556a7789a668273a72c1477cfc47fd0df814a"
	// "type":"0x2"
	// "maxFeePerGas":"0xee6b280"
	// "maxPriorityFeePerGas":"0x5f5e100"
	// "chainId":"0x10e"
	// "l1BatchNumber":"0x7"
	// "l1BatchTxIndex":"0x0"

	// trasaccion "manual"

	address2 := new(common.Address)
	*address2 = common.HexToAddress("0x08f5f9a336Aae6A72c795Ddf307864B13d13F0Aa")

	r := new(big.Int)
	r.SetString("7b83ef6f81a8eb124c84b7d80b133600c89c9bd67aa565d4024a849d759da56b", 16)
	s := new(big.Int)
	s.SetString("762462c7880aa0b713583c9096b556a7789a668273a72c1477cfc47fd0df814a", 16)

	transaction := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   big.NewInt(int64(270)),                 // (*a.signer).Domain().ChainId,
			Nonce:     1,                                      // preparedTx.Nonce.Uint64(),
			GasTipCap: big.NewInt(int64(0)),                   // preparedTx.GasTipCap,
			GasFeeCap: big.NewInt(int64(1000000000)),          // preparedTx.GasFeeCap,
			Gas:       171225,                                 // preparedTx.Gas.Uint64(),
			To:        address2,                               // preparedTx.To,
			Value:     big.NewInt(int64(1000000000000000000)), // preparedTx.Value,
			V:         big.NewInt(int64(1)),
			R:         r,
			S:         s,
		})

	err = client.SendTransaction(context.Background(), transaction)
	if err != nil {
		log.Println(err)
	}

	signer := types.Signer()

	log.Println(transaction.WithSignature())
}
