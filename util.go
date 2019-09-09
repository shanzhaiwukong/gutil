package gutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
)

// AesEncrypt 加密
// data：待加密的数据 key：密钥
// return string (is a base64 stdencoding)
func AesEncrypt(data, key []byte) (res string, err error) {
	defer func() {
		if ok := recover(); ok != nil {
			res = ""
			err = errors.New("异常")
		}
	}()
	var pkcs5Padding = func(ciphertext []byte, blockSize int) []byte {
		padding := blockSize - len(ciphertext)%blockSize
		padtext := bytes.Repeat([]byte{byte(padding)}, padding)
		return append(ciphertext, padtext...)
	}
	if block, err := aes.NewCipher(key); err != nil {
		return "", err
	} else {
		size := block.BlockSize()
		data = pkcs5Padding(data, size)
		model := cipher.NewCBCEncrypter(block, key[:size])
		crypted := make([]byte, len(data))
		model.CryptBlocks(crypted, data)
		result := base64.StdEncoding.EncodeToString(crypted)
		return result, nil
	}
}

// AesDecrypt 解密
// data：带解密的字符串（加密后的字符串） key：密钥
// return []byte (is orgin data)
func AesDecrypt(data string, key []byte) (res []byte, err error) {
	defer func() {
		if ok := recover(); ok != nil {
			res = nil
			err = errors.New("异常")
		}
	}()
	var pkcs5UnPadding = func(origData []byte) []byte {
		length := len(origData)
		unpadding := int(origData[length-1])
		return origData[:(length - unpadding)]
	}
	if block, err := aes.NewCipher(key); err != nil {
		return nil, err
	} else {
		srcData, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return nil, err
		}
		size := block.BlockSize()
		model := cipher.NewCBCDecrypter(block, key[:size])
		result := make([]byte, len(srcData))
		model.CryptBlocks(result, srcData)
		result = pkcs5UnPadding(result)
		return result, nil
	}
}

// LoadJSON2Struct 从json文件转转换对象
// filePath 文件路径
// v (need a pointer interface)
func LoadJSON2Struct(filePath string, v interface{}) error {
	if bs, err := ioutil.ReadFile(filePath); err != nil {
		return err
	} else {
		if err = json.Unmarshal(bs, v); err != nil {
			return err
		}
		return nil
	}
}
