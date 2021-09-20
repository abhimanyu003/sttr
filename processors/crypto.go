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
func MD5Encode(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))

	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA1Encode encode string to sha1.
func SHA1Encode(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

// SHA256Encode encode string to sha256.
func SHA256Encode(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

// SHA512Encode encode string to sha256.
func SHA512Encode(data string) string {
	h := sha512.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
