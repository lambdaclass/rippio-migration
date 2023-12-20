package main

import (
	"context"
	"fmt"
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

type MyTransaction struct {
    *types.Transaction
}

func (tx *MyTransaction) WithSignature2() (types.Transaction) {
	r := new(big.Int)
	r.SetString("106287454196759066847211792031325706661401235312908304949829649037033979701705", 10)
	s := new(big.Int)
	s.SetString("7989732764615526430734405065785953595367951189355047462118874195239141735945", 10)
	v := big.NewInt(int64(872))
	chainId := big.NewInt(int64(418))

	cpy := tx.inner.Copy()
	cpy.SetSignatureValues(chainId, v, r, s)

	return types.Transaction{inner: cpy, time: tx.time}
}

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

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Block number: ", blockNumber)

	// {"Nonce":0,"GasPrice":1000,"Gas":21000,"To":"0x08f5f9a336Aae6A72c795Ddf307864B13d13F0Aa","Value":1000000000000000000,"Input":null,"V":872,"R":106287454196759066847211792031325706661401235312908304949829649037033979701705,"S":7989732764615526430734405065785953595367951189355047462118874195239141735945,"Hash":"0x0785cb804d5e6db08d336922d68785f66de28d918da907ba1710b4c066eef477","From":"0x0000000000000000000000000000000000000000"}

	address := new(common.Address)
	*address = common.HexToAddress("0x08f5f9a336Aae6A72c795Ddf307864B13d13F0Aa")



	transaction := MyTransaction{types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   big.NewInt(int64(418)), // (*a.signer).Domain().ChainId,
			Nonce:     0,                      // preparedTx.Nonce.Uint64(),
			GasTipCap: big.NewInt(int64(0)),   // preparedTx.GasTipCap,
			GasFeeCap: big.NewInt(int64(0)),   // preparedTx.GasFeeCap,
			Gas:       21000,                  // preparedTx.Gas.Uint64(),
			To:        address,                // preparedTx.To,
			Value:     nil,                    // preparedTx.Value,
		})}

	
	transactionWithSignature := transaction.WithSignature2()
	if err != nil {
		log.Println(err)
	}

	err = client.SendTransaction(context.Background(), transactionWithSignature)
}
