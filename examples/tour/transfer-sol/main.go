package main

import (
	"context"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

var feePayer = types.AccountFromPrivateKeyBytes([]byte{178, 244, 76, 4, 247, 41, 113, 40, 111, 103, 12, 76, 195, 4, 100, 123, 88, 226, 37, 56, 209, 180, 92, 77, 39, 85, 78, 202, 121, 162, 88, 29, 125, 155, 223, 107, 139, 223, 229, 82, 89, 209, 27, 43, 108, 205, 144, 2, 74, 159, 215, 57, 198, 4, 193, 36, 161, 50, 160, 119, 89, 240, 102, 184})

var alice = types.AccountFromPrivateKeyBytes([]byte{196, 114, 86, 165, 59, 177, 63, 87, 43, 10, 176, 101, 225, 42, 129, 158, 167, 43, 81, 214, 254, 28, 196, 158, 159, 64, 55, 123, 48, 211, 78, 166, 127, 96, 107, 250, 152, 133, 208, 224, 73, 251, 113, 151, 128, 139, 86, 80, 101, 70, 138, 50, 141, 153, 218, 110, 56, 39, 122, 181, 120, 55, 86, 185})

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			sysprog.Transfer(
				alice.PublicKey, // from
				common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"), // to
				1e9, // 1 SOL
			),
		},
		Signers:         []types.Account{feePayer, alice},
		FeePayer:        feePayer.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		log.Fatalf("send raw tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)
}
