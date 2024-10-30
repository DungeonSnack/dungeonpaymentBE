package slug

import "strings"

func GenerateSlug(NamaProduk string) string {
    return strings.ToLower(strings.ReplaceAll(NamaProduk, " ", "-"))
}