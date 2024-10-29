package slug

import "strings"

func GenerateSlug(Namaorderan string) string {
    return strings.ToLower(strings.ReplaceAll(Namaorderan, " ", "-"))
}