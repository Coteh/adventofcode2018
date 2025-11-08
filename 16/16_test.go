package main

import "testing"

func Test_addr(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,5,4}

	addr(&reg, 0, 3, 2)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_addi(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,3,8}

	addi(&reg, 0, 7, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_mulr(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{6,2,3,4}

	mulr(&reg, 1, 2, 0)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_muli(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,3,14}

	muli(&reg, 1, 7, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_banr(t *testing.T) {
	reg := Register{6,7,8,9}
	expectedReg := Register{1,7,8,9}

	banr(&reg, 1, 3, 0)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_bani(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{4,2,3,4}

	bani(&reg, 3, 15, 0)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_borr(t *testing.T) {
	reg := Register{1,2,3,12}
	expectedReg := Register{15,2,3,12}

	borr(&reg, 2, 3, 0)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_bori(t *testing.T) {
	reg := Register{2,3,4,5}
	expectedReg := Register{15,3,4,5}

	bori(&reg, 3, 10, 0)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_setr(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{4,2,3,4}

	setr(&reg, 3, 0)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_seti(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{99,2,3,4}

	seti(&reg, 99, 0)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_gtir(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,3,1}

	gtir(&reg, 4, 2, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}

	reg = Register{1,2,3,4}
	expectedReg = Register{1,2,3,0}

	gtir(&reg, 2, 2, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_gtri(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,1,3,4}

	gtri(&reg, 3, 3, 1)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}

	reg = Register{1,2,3,4}
	expectedReg = Register{1,0,3,4}

	gtri(&reg, 1, 3, 1)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_gtrr(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,3,0}

	gtrr(&reg, 1, 2, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}

	reg = Register{1,2,3,4}
	expectedReg = Register{1,2,3,1}

	gtrr(&reg, 3, 2, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_eqir(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,3,1}

	eqir(&reg, 2, 1, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}

	reg = Register{1,2,3,4}
	expectedReg = Register{1,2,3,0}

	eqir(&reg, 4, 1, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_eqri(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,3,1}

	eqri(&reg, 2, 3, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}

	reg = Register{1,2,3,4}
	expectedReg = Register{1,2,3,0}

	eqri(&reg, 2, 1, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}

func Test_eqrr(t *testing.T) {
	reg := Register{1,2,3,4}
	expectedReg := Register{1,2,3,1}

	eqrr(&reg, 1, 1, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}

	reg = Register{1,2,3,4}
	expectedReg = Register{1,2,3,0}

	eqrr(&reg, 1, 2, 3)

	if !compareRegisters(reg, expectedReg) {
		t.Fatalf("case does not match, expected: '%v', actual: '%v'", expectedReg, reg)
	}
}
