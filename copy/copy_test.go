package copy

import "testing"

/**
 * @Author: LFM
 * @Date: 2022/8/7 23:29
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestDeepCopy(t *testing.T) {
	type User struct {
		Name string
	}

	user1 := &User{Name: "pibigstar"}
	var user2 User
	err := DeepCopy(user1, &user2)
	if err != nil {
		t.Error(err)
	}
	t.Log(user2)
}
