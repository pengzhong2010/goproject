package util

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密,CBC
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

//base64编码
func Base64Encrypt(input []byte) (output string, err error) {
	if len(input) == 0 {
		return
	}
	output = base64.StdEncoding.EncodeToString(input)
	return
}

//base64解码
func Base64Decrypt(input string) (output []byte, err error) {
	if len(input) == 0 {
		return
	}
	output, err = base64.StdEncoding.DecodeString(input)
	return
}

//gzip压缩
func MarshalToJsonWithGzip(in []byte) (out []byte, err error) {
	var buf bytes.Buffer
	gzipWiter := gzip.NewWriter(&buf)
	_, err = gzipWiter.Write(in)
	if err != nil {
		return
	}
	err = gzipWiter.Close()
	if err != nil {
		return
	}
	out = buf.Bytes()
	return
}

//gzip解压
func UnmarshalDataFromJsonWithGzip(in []byte) (out []byte, err error) {
	bytesReader := bytes.NewReader(in)
	gzipReader, err := gzip.NewReader(bytesReader)
	if err != nil {
		return
	}
	defer gzipReader.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(gzipReader)
	if err != nil {
		return
	}
	out = buf.Bytes()
	return
}
