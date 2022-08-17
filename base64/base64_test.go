package base64

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author: LFM
 * @Date: 2022/8/18 00:04
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestBase64(t *testing.T) {
	a := assert.New(t)
	str := "5203344abcdAbcd"
	encode := Encode(str)

	decode, _ := Decode(encode)
	a.Equal(str, decode)
}
