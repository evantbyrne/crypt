# crypt

Utility for encrypting and decrypting files with AES-256 GCM and Scrypt.

## Install

Make sure that [Go](https://golang.org/) is installed and your PATH includes GOBIN. Then run the following:

	bash$ go get github.com/evantbyrne/crypt
	bash$ go get gopkg.in/alecthomas/kingpin.v2
	bash$ go install github.com/evantbyrne/crypt/crypt

## Usage

	bash$ crypt --help
	usage: crypt --in=IN [<flags>]

	Utility for encrypting and decrypting files with AES-256 GCM and Scrypt.

	Flags:
	  --help         Show context-sensitive help (also try --help-long and --help-man).
	  -e, --encrypt  Encryption mode.
	  -d, --decrypt  Decryption mode.
	  -i, --in=IN    Input file.
	  -o, --out=OUT  Output file.

Encrypt file and store in file:

	bash$ crypt -i foo.txt -e -o foo.txt.crypt
	Password: 

Decrypt file and display in terminal:

	bash$ crypt -i foo.txt.crypt -d
	Password: 
	The quick brown fox jumps over the lazy dog.

Decrypt file and store in file:

	bash$ crypt -i foo.txt.crypt -d -o bar.txt
	Password: 

## Encrypted data format

	+---------------------+-------------------+-------------------+
	| 12 byte scrypt salt | 12 byte gcm nonce | encrypted data... |
	+---------------------+-------------------+-------------------+
