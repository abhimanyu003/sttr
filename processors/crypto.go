package processors

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// MD5Encode encode string to md5.
type MD5Encode struct{}

func (p MD5Encode) Name() string {
	return "md5-sum"
}

func (p MD5Encode) Alias() []string {
	return []string{"md5"}
}

func (p MD5Encode) Transform(data string, _ ...Flag) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(data))

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (p MD5Encode) Flags() []Flag {
	return nil
}

func (p MD5Encode) Title() string {
	return "MD5 Sum"
}

func (p MD5Encode) Description() string {
	return "Get the MD5 hash of your text"
}

func (p MD5Encode) FilterValue() string {
	return p.Title()
}

// SHA1Encode encode string to sha1.
type SHA1Encode struct{}

func (p SHA1Encode) Name() string {
	return "sha1-sum"
}

func (p SHA1Encode) Alias() []string {
	return []string{"sha1"}
}

func (p SHA1Encode) Transform(data string, _ ...Flag) (string, error) {
	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA1Encode) Flags() []Flag {
	return nil
}

func (p SHA1Encode) Title() string {
	return "SHA1 Sum"
}

func (p SHA1Encode) Description() string {
	return "Get the SHA1 hash of your text"
}

func (p SHA1Encode) FilterValue() string {
	return p.Title()
}

// SHA256Encode encode string to sha256.
type SHA256Encode struct{}

func (p SHA256Encode) Name() string {
	return "sha256-sum"
}

func (p SHA256Encode) Alias() []string {
	return []string{"sha256"}
}

func (p SHA256Encode) Transform(data string, _ ...Flag) (string, error) {
	h := sha256.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA256Encode) Flags() []Flag {
	return nil
}

func (p SHA256Encode) Title() string {
	return "SHA256 Sum"
}

func (p SHA256Encode) Description() string {
	return "Get the SHA256 hash of your text"
}

func (p SHA256Encode) FilterValue() string {
	return p.Title()
}

// SHA512Encode encode string to sha256.
type SHA512Encode struct{}

func (p SHA512Encode) Name() string {
	return "sha512-sum"
}

func (p SHA512Encode) Alias() []string {
	return []string{"sha256"}
}

func (p SHA512Encode) Transform(data string, _ ...Flag) (string, error) {
	h := sha512.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA512Encode) Flags() []Flag {
	return nil
}

func (p SHA512Encode) Title() string {
	return "SHA512 Sum"
}

func (p SHA512Encode) Description() string {
	return "Get the SHA512 checksum of your text"
}

func (p SHA512Encode) FilterValue() string {
	return p.Title()
}
