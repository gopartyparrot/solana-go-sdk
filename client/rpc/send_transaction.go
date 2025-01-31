package rpc

import (
	"context"
	"encoding/base64"
	"errors"
)

type SendTransactionConfig struct {
	SkipPreflight       bool       `json:"skipPreflight"`       // default: false
	PreflightCommitment Commitment `json:"preflightCommitment"` // default: max
	Encoding            string     `json:"encoding"`            // base58 or base64
}

// SendRawTransaction is a quick way to send the serialize tx
func (s *RpcClient) SendRawTransaction(ctx context.Context, tx []byte) (string, error) {
	res := struct {
		GeneralResponse
		Result string `json:"result"`
	}{}
	err := s.request(
		ctx,
		"sendTransaction",
		[]interface{}{
			base64.StdEncoding.EncodeToString([]byte(tx)),
			SendTransactionConfig{
				SkipPreflight:       false,
				PreflightCommitment: CommitmentFinalized,
				Encoding:            "base64",
			},
		},
		&res,
	)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", errors.New(res.Error.Message)
	}
	return res.Result, nil
}

// SendTransaction submits a signed transaction to the cluster for processing
func (s *RpcClient) SendTransaction(ctx context.Context, tx string, cfg SendTransactionConfig) (string, error) {
	res := struct {
		GeneralResponse
		Result string `json:"result"`
	}{}
	err := s.request(ctx, "sendTransaction", []interface{}{tx, cfg}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", errors.New(res.Error.Message)
	}
	return res.Result, nil
}
