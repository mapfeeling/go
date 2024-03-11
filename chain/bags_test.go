package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

const (
	binanceChainPrivateKeyFile = "/Users/hcl/work/xiaomi/myGo/mapfeeling/go/chain/private_key.txt"
	binanceChainAddressFile    = "/Users/hcl/work/xiaomi/myGo/mapfeeling/go/chain/address.txt"
)

func TestBags(t *testing.T) {
	// 生成私钥的数量
	numOfWallets := 10

	for i := 0; i < numOfWallets; i++ {
		privateKey, err := generatePrivateKey()
		if err != nil {
			log.Fatalf("Failed to generate private key: %v", err)
		}
		fmt.Printf("Private Key %d: %s\n", i, privateKey)
		writePrivateKeyToFile(privateKey)
	}
}

func generatePrivateKey() (string, error) {
	// 生成随机私钥，这里简化为SHA256哈希值作为私钥（实际中应使用更安全的随机数生成算法）
	privateKeyBytes := sha256.Sum256([]byte(fmt.Sprintf("%d", 12345))) // 仅为示例，实际应使用更安全的随机数生成方法
	privateKey := hex.EncodeToString(privateKeyBytes[:])
	return privateKey, nil
}

func writePrivateKeyToFile(privateKey string) {
	err := ioutil.WriteFile(binanceChainPrivateKeyFile, []byte(privateKey), 0644)
	if err != nil {
		log.Fatalf("Failed to write private key to file: %v", err)
	}
	fmt.Println("Private key written to file:", binanceChainPrivateKeyFile)
}
