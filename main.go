package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/alecthomas/kingpin"
	"golang.org/x/crypto/ssh/terminal"

	crypt "github.com/evantbyrne/crypt/lib"
)

var (
	app         = kingpin.New("crypt", "Utility for encrypting and decrypting files with AES-256 GCM and Scrypt.")
	encryptFlag = app.Flag("encrypt", "Encryption mode.").Short('e').Bool()
	decryptFlag = app.Flag("decrypt", "Decryption mode.").Short('d').Bool()
	inFlag      = app.Flag("in", "Input file.").Short('i').Required().String()
	outFlag     = app.Flag("out", "Output file.").Short('o').String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if !*encryptFlag && !*decryptFlag {
		app.Usage(nil)
		os.Exit(1)
	}

	stdin, err := ioutil.ReadFile(*inFlag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	if *encryptFlag {
		salt := crypt.NewNonce()
		key := crypt.NewKey(salt, password)
		nonce := crypt.NewNonce()
		ciphertext, err := crypt.Encrypt(key, nonce, stdin)
		if err != nil {
			log.Fatal(err)
		}

		out := append(salt, nonce...)
		out = append(out, ciphertext...)

		if *outFlag == "" {
			fmt.Printf("%s\n", out)
		} else {
			err = ioutil.WriteFile(*outFlag, out, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if *decryptFlag {
		salt := stdin[:12]
		nonce := stdin[12:24]
		ciphertext := stdin[24:]
		key := crypt.NewKey(salt, password)
		out, err := crypt.Decrypt(key, nonce, ciphertext)
		if err != nil {
			log.Fatal(err)
		}

		if *outFlag == "" {
			fmt.Printf("%s\n", out)
		} else {
			err = ioutil.WriteFile(*outFlag, out, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
