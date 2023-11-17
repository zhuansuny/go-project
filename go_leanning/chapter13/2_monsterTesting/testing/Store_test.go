package testing

import (
	"go_code/chapter13/2_monsterTesting/monster"
	"testing"
)

func TestStore(t *testing.T) {
	monster := model.Monster{
		Name:  "牛魔王",
		Age:   199,
		Skill: "牛魔拳",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("储存 执行错误，期望值=%v 实际值=%v\n", true, false)
	}

	t.Logf("储存 执行正确...")
	name := monster.Name
	monster.Name = ""
	monster.ReStore()
	if name != monster.Name {
		t.Fatalf("储存 执行错误，期望值=%v 实际值=%v\n", name, monster.Name)
	}

	t.Logf("储存 执行正确...")
}

func TestReStore(t *testing.T) {
	var monster model.Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("读取 执行错误，期望值=%v 实际值=%v\n", true, false)
	}

	t.Logf("储存 执行正确...")
	monster.ReStore()
	if monster.Name != "牛魔王1" {
		t.Fatalf("储存 执行结果错误，期望值=%v 实际值=%v\n", "牛魔王1", monster.Name)
	}

	t.Logf("储存读取 执行结果正确...")
}
