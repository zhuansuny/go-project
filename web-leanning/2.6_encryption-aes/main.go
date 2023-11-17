package main

//crypto/aes包：AES(Advanced Encryption Standard)，又称Rijndael加密法，
//是美国联邦政府采用的一种区块加密标准。

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	plaintext := []byte("my name is zsyy") //需要去加密的字符串
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}
	//aes的加密字符串
	key_text := "zsyyddr12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}
	fmt.Println("key_text len:", len(key_text))

	//创建加密算法aes
	//上面通过调用函数aes.NewCipher(参数key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法),返回了一个cipher.Block接口
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	//解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)

}
