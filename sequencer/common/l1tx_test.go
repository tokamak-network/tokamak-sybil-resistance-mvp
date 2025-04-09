package common

import "testing"

func TestTypeFromByte(t *testing.T) {
	tests := []struct {
		name string
		in   byte
		out  TxType
	}{
		{"create account deposit", 0, TxTypeCreateAccountDeposit},
		{"deposit", 1, TxTypeDeposit},
		{"withdraw", 2, TxTypeWithdraw},
		{"create vouch", 3, TxTypeCreateVouch},
		{"delete vouch", 4, TxTypeDeleteVouch},
		{"explode", 5, TxTypeExplode},
		{"unknown", 6, TxTypeUnknown},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := TypeFromByte(test.in); got != test.out {
				t.Errorf("TypeFromByte(%d) = %v, want %v", test.in, got, test.out)
			}
		})
	}
}
