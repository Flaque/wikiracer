package wikimedia

import (
	"strings"
)

func IsValidLink(link string) bool {
	isLocalLink := strings.HasPrefix(link, "./")
	hasNoPrefix := !strings.ContainsRune(TrimLinkPrefix(link), ':') // This could break if there's items that have other prefixes.
	return isLocalLink && hasNoPrefix
}

func IgnoreHashes(link string) string {
	return strings.Split(link, "#")[0]
}

func TrimLinkPrefix(link string) string {
	return strings.TrimPrefix(link, "./")
}

func LinkToTitle(link string) string {
	str := TrimLinkPrefix(link)
	return IgnoreHashes(str)
}
