package gutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
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

// MD5 加密
func MD5(tag string) string {
	h := md5.New()
	h.Write([]byte(tag))
	return hex.EncodeToString(h.Sum(nil))
}

var lock sync.Mutex

// NewID 随机生成32位ID
func NewID() string {
	lock.Lock()
	defer lock.Unlock()
	b := make([]byte, 48)
	str := fmt.Sprintf("%d", time.Now().UnixNano())
	if _, err := io.ReadFull(rand.Reader, b); err == nil {
		str += base64.URLEncoding.EncodeToString(b)
	}
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// RegGetMapByName 正则表达是提取命名组结果
func RegGetMapByName(regRule, tagStr string) map[string]string {
	result := make(map[string]string)
	reg, err := regexp.Compile(regRule)
	if err == nil {
		matchs := reg.FindStringSubmatch(tagStr)
		if matchs != nil {
			for i, name := range reg.SubexpNames() {
				if !(name == "") {
					result[name] = matchs[i]
				}
			}
		}
	}
	return result
}

// If 三元运算符
//eg a,b:=0,1 If(a>b,a,b)
func If(condition bool, yesVal, noVal interface{}) interface{} {
	if condition {
		return yesVal
	}
	return noVal
}

// GetFileContent 获取文件内容
func GetFileContent(filePath string) string {
	if data, err := ioutil.ReadFile(filePath); err != nil {
		return ""
	} else {
		return string(data)
	}
}

// 检查文件是否存在
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

//创建文件夹
func makeDirAll(filePath string) error {
	return os.MkdirAll(filepath.Dir(filePath), 0660)
}

// WriteFileByOverwrite 以覆盖的方式写入文件 如果不存在 则创建
func WriteFileByOverwrite(filePath string, content []byte) error {
	if err := makeDirAll(filePath); err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, content, 0660)
}

// WriteFileByAppend 以追加的方式写入文件 如果不存在 则创建
func WriteFileByAppend(filePath, content string) error {
	if err := makeDirAll(filePath); err != nil {
		return err
	}
	var f *os.File
	var err error
	if checkFileIsExist(filePath) {
		f, err = os.OpenFile(filePath, os.O_APPEND, 0660)
		if err != nil {
			return err
		}
	} else {
		f, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	_, err = io.WriteString(f, content)
	return err
}
