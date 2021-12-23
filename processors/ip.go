package processors

import (
	"net"
	"regexp"
	"strings"
)

// ExtractIPs encode string to base64.
type ExtractIPs struct{}

func (p ExtractIPs) FilterValue() string {
	return p.Title()
}

func (p ExtractIPs) Name() string {
	return "extract-ip"
}

func (p ExtractIPs) Alias() []string {
	return []string{"find-ips", "find-ip", "extract-ips"}
}

func (p ExtractIPs) Transform(data []byte, _ ...Flag) (string, error) {
	var tempIps []string
	var validIps []string

	// Find all IPv4
	ipv4RegexString := regexp.MustCompile(`([0-9]{0,3}\.){3}[0-9]{0,3}`)
	ipv4 := ipv4RegexString.FindAllString(string(data), -1)

	// Find all IPv6
	ipv6RegexString := regexp.MustCompile(`(fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])\.{3,3})(25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])\.{3,3})(25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])|:((:[0-9a-fA-F]{1,4}){1,7}|:))`)
	ipv6 := ipv6RegexString.FindAllString(string(data), -1)

	tempIps = append(tempIps, ipv4...)
	tempIps = append(tempIps, ipv6...)

	for _, v := range tempIps {
		if ip := net.ParseIP(v); ip != nil {
			validIps = append(validIps, ip.String())
		}
	}

	return strings.Join(validIps, "\n"), nil
}

func (p ExtractIPs) Flags() []Flag {
	return nil
}

func (p ExtractIPs) Title() string {
	return "Extract IPs"
}

func (p ExtractIPs) Description() string {
	return "Extract IPv4 and IPv6 from your text"
}
