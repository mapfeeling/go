package _024_01

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	sentinel2 "git.n.xiaomi.com/mitv-media/dragonlib.go/sentinel"
)

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func Decrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func Encrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func Test20240109(t *testing.T) {
	key := []byte("2fa6c1e9")
	s, _ := Encrypt("hello", key)
	fmt.Println(s)
	//fmt.Println(Decrypt("pWdWjMl628Gy5oVyNaQdQBAwuZwqTrKt", []byte("constant")))
	//data := []byte("hello world")
	//hash := sha256.Sum256(data)
	//fmt.Println(hash)
	//fmt.Printf("SHA256哈希值：%x", hash)
	//
	//// 务必先进行初始化
	//err := sentinel.InitDefault()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 配置一条限流规则
	//_, err = flow.LoadRules([]*flow.Rule{
	//	{
	//		Resource:               "some-test", //熔断器规则生效的埋点资源的名称
	//		TokenCalculateStrategy: flow.Direct, //Direct表示直接使用字段
	//		ControlBehavior:        flow.Reject, //Reject表示超过阈值直接拒绝
	//		Threshold:              10,          //QPS流控阈值10个
	//		StatIntervalInMs:       1000,        //QPS时间长度1s
	//	},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	//log.Fatalf("配置规则出错：", err.Error())
	//}
	//b()
	////模拟流量
	////for i := 0; i < 12; i++ {
	////	e, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
	////	if b != nil {
	////		fmt.Println("限流了")
	////	} else {
	////		fmt.Println("检查通过")
	////		e.Exit()
	////	}
	////}
}

func b() {
	//模拟流量
	for i := 0; i < 12; i++ {
		var restySentinelClient = new(sentinel2.RestySentinelClient)
		r := restySentinelClient.SetSentinelParams("some-test", "some-test", "some-test").HttpSentinelCheck().CheckAndRelease()
		if r == nil {
			fmt.Println("restySentinelClient == nil")
		} else {
			fmt.Println(r.Err, r.SentinelBlockError)
		}

	}
}
