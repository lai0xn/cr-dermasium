package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func ComputeHMAC(data []byte) string {
	h := hmac.New(sha256.New, []byte(CHARGILIY_SECRET_KEY))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
