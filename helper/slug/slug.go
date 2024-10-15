package slug

import "strings"

func GenerateSlug(NamaKategori string) string {
    return strings.ToLower(strings.ReplaceAll(NamaKategori, " ", "-"))
}