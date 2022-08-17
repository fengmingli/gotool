package base64

/**
 * @Author: LFM
 * @Date: 2022/8/18 00:04
 * @Since: 1.0.0
 * @Desc: TODO
 */

import (
	b64 "encoding/base64"
)

func Encode(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}

func Decode(s string) (string, error) {
	ds, err := b64.StdEncoding.DecodeString(s)
	return string(ds), err
}
