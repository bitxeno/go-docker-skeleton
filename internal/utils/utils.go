package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Ext(path string) string {
	return strings.ToLower(filepath.Ext(path))
}

func FileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func IsCJK(s string) bool {
	for _, r := range []rune(s) {
		c := int(r)
		if (c >= 0x4E00 && c <= 0x9FFF) ||
			(c >= 0x3400 && c <= 0x4DBF) ||
			(c >= 0x20000 && c <= 0x2A6DF) ||
			(c >= 0x2A700 && c <= 0x2B73F) ||
			(c >= 0x2B740 && c <= 0x2B81F) ||
			(c >= 0x2B820 && c <= 0x2CEAF) ||
			(c >= 0x2CEB0 && c <= 0x2EBEF) ||
			(c >= 0x30000 && c <= 0x3134F) ||
			(c >= 0xF900 && c <= 0xFAFF) ||
			(c >= 0x2F800 && c <= 0x2FA1F) ||
			(c >= 0x3000 && c <= 0x303F) ||
			(c >= 0x3040 && c <= 0x309F) ||
			(c >= 0x30A0 && c <= 0x30FF) ||
			(c >= 0xAC00 && c <= 0xD7AF) ||
			(c >= 0xFF00 && c <= 0xFFEF) {
			return true
		}
	}
	return false
}
