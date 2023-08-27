package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

/*
	This hash function will generate a sum on sha256 and return it as a byte array
	to then be encoded.
*/
func hash(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))

	return algo.Sum(nil)
}


/*
	This function will generate a base58 encoded string from a byte array.
*/
func encode(bytes []byte) string {
	encoder := base58.BitcoinEncoding
	encoded, err := encoder.Encode(bytes)
	if err != nil {
		panic(err)
	}

	return string(encoded)
}

func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := hash(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	
	return finalString[:8]
}