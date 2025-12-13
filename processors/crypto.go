package processors

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// MD5 encodes string to MD5.
type MD5 struct{}

// Implement StreamingProcessor interface
func (p MD5) CanStream() bool {
	return true
}

func (p MD5) PreferStream() bool {
	return true // Hash functions benefit greatly from streaming
}

func (p MD5) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	hasher := md5.New()
	_, err := io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := hex.EncodeToString(hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

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
	title := "MD5 Sum"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p MD5) Description() string {
	return "Get the MD5 checksum of your text"
}

func (p MD5) FilterValue() string {
	return p.Title()
}

// SHA1 encodes string to SHA-1.
type SHA1 struct{}

// Implement StreamingProcessor interface
func (p SHA1) CanStream() bool {
	return true
}

func (p SHA1) PreferStream() bool {
	return true
}

func (p SHA1) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	hasher := sha1.New()
	_, err := io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := fmt.Sprintf("%x", hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

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
	title := "SHA-1 Sum"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p SHA1) Description() string {
	return "Get the SHA-1 checksum of your text"
}

func (p SHA1) FilterValue() string {
	return p.Title()
}

// SHA256 encodes string to SHA-256.
type SHA256 struct{}

// Implement StreamingProcessor interface
func (p SHA256) CanStream() bool {
	return true
}

func (p SHA256) PreferStream() bool {
	return true
}

func (p SHA256) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	hasher := sha256.New()
	_, err := io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := fmt.Sprintf("%x", hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

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
	title := "SHA-256 Sum"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p SHA256) Description() string {
	return "Get the SHA-256 checksum of your text"
}

func (p SHA256) FilterValue() string {
	return p.Title()
}

// SHA512 encodes string to SHA-512.
type SHA512 struct{}

// Implement StreamingProcessor interface
func (p SHA512) CanStream() bool {
	return true
}

func (p SHA512) PreferStream() bool {
	return true
}

func (p SHA512) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	hasher := sha512.New()
	_, err := io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := fmt.Sprintf("%x", hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

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
	title := "SHA-512 Sum"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p SHA512) Description() string {
	return "Get the SHA-512 checksum of your text"
}

func (p SHA512) FilterValue() string {
	return p.Title()
}

// SHA224 encode string to SHA-224.
type SHA224 struct{}

// Implement StreamingProcessor interface
func (p SHA224) CanStream() bool {
	return true
}

func (p SHA224) PreferStream() bool {
	return true
}

func (p SHA224) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	hasher := sha256.New224()
	_, err := io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := fmt.Sprintf("%x", hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

func (p SHA224) Name() string {
	return "sha224"
}

func (p SHA224) Alias() []string {
	return []string{"sha224-sum"}
}

func (p SHA224) Transform(data []byte, _ ...Flag) (string, error) {
	bs := sha256.Sum224(data)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA224) Flags() []Flag {
	return nil
}

func (p SHA224) Title() string {
	title := "SHA-224 Sum"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p SHA224) Description() string {
	return "Get the SHA-224 checksum of your text"
}

func (p SHA224) FilterValue() string {
	return p.Title()
}

// SHA384 encodes string to SHA-384.
type SHA384 struct{}

// Implement StreamingProcessor interface
func (p SHA384) CanStream() bool {
	return true
}

func (p SHA384) PreferStream() bool {
	return true
}

func (p SHA384) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	hasher := sha512.New384()
	_, err := io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := fmt.Sprintf("%x", hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

func (p SHA384) Name() string {
	return "sha384"
}

func (p SHA384) Alias() []string {
	return []string{"sha384-sum"}
}

func (p SHA384) Transform(data []byte, _ ...Flag) (string, error) {
	bs := sha512.Sum384(data)

	return fmt.Sprintf("%x", bs), nil
}

func (p SHA384) Flags() []Flag {
	return nil
}

func (p SHA384) Title() string {
	title := "SHA-384 Sum"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p SHA384) Description() string {
	return "Get the SHA-384 checksum of your text"
}

func (p SHA384) FilterValue() string {
	return p.Title()
}

// Bcrypt encodes string to bcrypt.
type Bcrypt struct{}

func (p Bcrypt) Name() string {
	return "bcrypt"
}

func (p Bcrypt) Alias() []string {
	return []string{"bcrypt-hash"}
}

func (p Bcrypt) Transform(data []byte, f ...Flag) (string, error) {
	var rounds uint
	for _, flag := range f {
		if flag.Short == "r" {
			r, ok := flag.Value.(uint)
			if ok {
				rounds = r
			}
		}
	}

	bytes, err := bcrypt.GenerateFromPassword(data, int(rounds))

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
	title := "bcrypt Hash"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Bcrypt) Description() string {
	return "Get the bcrypt hash of your text"
}

func (p Bcrypt) FilterValue() string {
	return p.Title()
}
