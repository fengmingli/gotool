package copy

/**
 * @Author: LFM
 * @Date: 2022/8/7 23:29
 * @Since: 1.0.0
 * @Desc: TODO
 */

import (
	"bytes"
	"encoding/gob"
)

//DeepCopy 基于序列化和反序列化的深度拷贝
func DeepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
