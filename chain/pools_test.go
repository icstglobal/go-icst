package chain

import "testing"

func TestPools(t *testing.T) {
	var blc Chain
	if _, err := Get(Eth); err == nil {
		t.Fatal("should report error when not block chain registered")
	}
	Set(Eth, blc)

	if _, err := Get(Eth); err != nil {
		t.Fatal("should not report error as Eth has been registered")
	}

	if _, err := Get(EOS); err == nil {
		t.Fatal("should report error as EOS has not been registered")
	}
}
