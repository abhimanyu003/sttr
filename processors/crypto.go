package processors

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// MD5 encode string to md5.
type MD5 struct{}

func (p MD5) Name() string {
	return "md5"
}

func (p MD5) Alias() []string {
	return []string{"md5-sum"}
}

func (p MD5) Transform(data []byte, _ ...Flag) (string, error) {

	hasher := md5.New()
	hasher.Write(data)

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (p MD5) Flags() []Flag {
	return nil
}

func (p MD5) Title() string {
	return "MD5 Sum"
}

func (p MD5) Description() string {
	return "Get the MD5 checksum of your text"
}

func (p MD5) FilterValue() string {
	return p.Title()
}

// SHA1 encode string to sha1.
type SHA1 struct{}

func (p SHA1) Name() string {
	return "sha1"
}

func (p SHA1) Alias() []string {
	return []string{"sha1-sum"}
}

func (p SHA1) Transform(data []byte, _ ...Flag) (string, error) {
	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA1) Flags() []Flag {
	return nil
}

func (p SHA1) Title() string {
	return "SHA1 Sum"
}

func (p SHA1) Description() string {
	return "Get the SHA1 checksum of your text"
}

func (p SHA1) FilterValue() string {
	return p.Title()
}

// SHA256 encode string to sha256.
type SHA256 struct{}

func (p SHA256) Name() string {
	return "sha256"
}

func (p SHA256) Alias() []string {
	return []string{"sha256-sum"}
}

func (p SHA256) Transform(data []byte, _ ...Flag) (string, error) {
	h := sha256.New()
	h.Write(data)
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA256) Flags() []Flag {
	return nil
}

func (p SHA256) Title() string {
	return "SHA256 Sum"
}

func (p SHA256) Description() string {
	return "Get the SHA256 checksum of your text"
}

func (p SHA256) FilterValue() string {
	return p.Title()
}

// SHA512 encode string to sha256.
type SHA512 struct{}

func (p SHA512) Name() string {
	return "sha512"
}

func (p SHA512) Alias() []string {
	return []string{"sha512-sum"}
}

func (p SHA512) Transform(data []byte, _ ...Flag) (string, error) {
	h := sha512.New()
	h.Write(data)
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA512) Flags() []Flag {
	return nil
}

func (p SHA512) Title() string {
	return "SHA512 Sum"
}

func (p SHA512) Description() string {
	return "Get the SHA512 checksum of your text"
}

func (p SHA512) FilterValue() string {
	return p.Title()
}

// Bcrypt encode string to Bcrypt.
type Bcrypt struct{}

func (p Bcrypt) Name() string {
	return "bcrypt"
}

func (p Bcrypt) Alias() []string {
	return []string{"bcrypt-hash"}
}

func (p Bcrypt) Transform(data []byte, f ...Flag) (string, error) {
	var rounds int
	for _, flag := range f {
		if flag.Short == "r" {
			r, ok := flag.Value.(int)
			if ok {
				rounds = r
			}
		}
	}

	bytes, err := bcrypt.GenerateFromPassword(data, rounds)

	return string(bytes), err
}

func (p Bcrypt) Flags() []Flag {
	return []Flag{
		{
			Name:  "number-of-rounds",
			Short: "r",
			Desc:  "Number of rounds",
			Value: 10,
			Type:  FlagUint,
		},
	}
}

func (p Bcrypt) Title() string {
	return "Bcrypt Hash"
}

func (p Bcrypt) Description() string {
	return "Get the Bcrypt hash of your text"
}

func (p Bcrypt) FilterValue() string {
	return p.Title()
}
