package wallet

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewMnemonic(t *testing.T) {

	expected := "guilt follow soon broken crawl radar pony refuse soft cement south course over galaxy impulse trouble erase mechanic raccoon tone action result erupt fork"

	actual, err := NewMnemonic()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(actual)
	assert.NotEqual(t, expected, actual)

}

func TestNewSeedFromMnemonic(t *testing.T) {

	mnemonic := "guilt follow soon broken crawl radar pony refuse soft cement south course over galaxy impulse trouble erase mechanic raccoon tone action result erupt fork"
	expect := []byte{188, 204, 171, 175, 188, 116, 164, 99, 55, 11, 89, 215, 55, 42, 158, 138, 254, 141, 1, 35, 56,
		199, 112, 251, 126, 215, 251, 229, 124, 141, 141, 216, 26, 172, 182, 67, 16, 184, 146, 91, 63, 47, 18, 91, 202,
		8, 182, 17, 244, 37, 204, 192, 108, 108, 107, 210, 60, 117, 177, 103, 192, 95, 41, 219}

	actual, err := NewSeedFromMnemonic(mnemonic, "")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expect, actual)

}
