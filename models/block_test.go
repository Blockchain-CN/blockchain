package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGenerateBlock(t *testing.T) {
	b := GenerateBlock("0", "TEST_0", 0)
	fmt.Println(b)
}

func TestFormatBlock(t *testing.T) {
	b := GenerateBlock("0", "TEST_0", 0)
	jb, _ := json.Marshal(b)
	bf, err := FormatBlock(jb)
	if err != nil {
		fmt.Println(err)
	}
	if *bf != *b {
		t.Fail()
	}
	fmt.Println("bf == b")
}

func TestBlock_IsValid(t *testing.T) {
	b := GenerateBlock("0", "TEST_0", 0)
	b1 := GenerateBlock(b.Hash, "TEST_1", 1)
	b2 := GenerateBlock(b.Hash, "TEST_2", 2)
	fmt.Println("b1 is behind b:", b1.IsValid(b))
	fmt.Println("b2 is behind b1:", b2.IsValid(b1))
	fmt.Println("b2 is behind b:", b2.IsValid(b))
}
