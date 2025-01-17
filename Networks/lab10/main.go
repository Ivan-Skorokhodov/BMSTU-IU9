package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Transaction struct {
	ChainID  *big.Int `json:"chainId"`
	Hash     string   `json:"hash"`
	Value    string   `json:"value"`
	Cost     string   `json:"cost"`
	To       string   `json:"to"`
	Gas      uint64   `json:"gas"`
	GasPrice string   `json:"gasPrice"`
}

type Block struct {
	Number        uint64        `json:"number"`
	Transactions []Transaction `json:"transactions"`
}

func sendToFirebase(block Block) error {
	firebaseURL := "https://bmstu-lab10-default-rtdb.europe-west1.firebasedatabase.app/blocks.json"

	blockJSON, err := json.Marshal(block)
	if err != nil {
		return fmt.Errorf("error marshaling block to JSON: %v", err)
	}

	resp, err := http.Post(firebaseURL, "application/json", strings.NewReader(string(blockJSON)))
	if err != nil {
		return fmt.Errorf("error sending POST request to Firebase: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected response from Firebase: %s", resp.Status)
	}

	return nil
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/01061036a71549b3933b31f171f64051")
	if err != nil {
		log.Fatalln(err)
	}

	previousBlockNumber := big.NewInt(0)

	for {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}

		blockNumber := big.NewInt(header.Number.Int64())

		if blockNumber.Cmp(previousBlockNumber) == 0 {
			time.Sleep(time.Second)
			continue
		}

		previousBlockNumber = blockNumber
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		blockData := Block{
			Number:        block.Number().Uint64(),
			Transactions: []Transaction{},
		}

		for _, tx := range block.Transactions() {
			toAddress := ""
			if tx.To() != nil {
				toAddress = tx.To().Hex()
			}

			transaction := Transaction{
				ChainID:  tx.ChainId(),
				Hash:     tx.Hash().Hex(),
				Value:    tx.Value().String(),
				Cost:     tx.Cost().String(),
				To:       toAddress,
				Gas:      tx.Gas(),
				GasPrice: tx.GasPrice().String(),
			}
			blockData.Transactions = append(blockData.Transactions, transaction)
		}

		if err := sendToFirebase(blockData); err != nil {
			log.Fatalf("Failed to send block data to Firebase: %v", err)
		}

		fmt.Printf("Block %s data successfully sent to Firebase\n", blockNumber.String())

	}


}