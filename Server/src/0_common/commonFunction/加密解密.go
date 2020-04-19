package commonFunction

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Encode和Decode要配合使用！

// 对str进行AES加密，密钥为key，返回加密后的Base64字符串
func Encode(str string, key string) (string, error) {
	// 对str进行填充，使其满足AES加密的要求
	paddingBytes := padding([]byte(str))

	// 加密后的长度和加密前的长度(已填充)是一样的
	resultBytes := make([]byte, len(paddingBytes))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 创建加密模式
	blockMode := cipher.NewCBCEncrypter(block, []byte(key))

	// 加密
	blockMode.CryptBlocks(resultBytes, paddingBytes)
	return base64.StdEncoding.EncodeToString(resultBytes), nil // 对加密后的结果进行Base64编码
}

// 对base64字符串进行AES解密,密钥为key
func Decode(str string, key string) (string, error) {
	paddingBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	// 解密后的长度和解密前的长度(已填充)是一样的
	resultBytes := make([]byte, len(paddingBytes))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 创建解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(key))

	// 解密
	blockMode.CryptBlocks(resultBytes, paddingBytes)
	return string(unPadding(resultBytes)), nil
}

// 以下是为了保证所需加密字符串符合AES加密的要求

// 将origin长度补全为 16字节 的倍数
func padding(origin []byte) []byte {
	paddingLength := 16 - len(origin)%16
	paddingBytes := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(origin, paddingBytes...)
}

// 还原原始字节数组
func unPadding(dst []byte) []byte {
	if len(dst) == 0 {
		return dst
	}
	unPaddingLength := int(dst[len(dst)-1])
	return dst[:len(dst)-unPaddingLength]
}
