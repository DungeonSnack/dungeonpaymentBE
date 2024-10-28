package slug

import "strings"

func GenerateSlug(Namaorder string) string {
    return strings.ToLower(strings.ReplaceAll(Namaorder, " ", "-"))
}