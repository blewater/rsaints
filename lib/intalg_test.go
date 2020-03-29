package lib

import (
	"testing"
)

func TestIsPrimeTrialByError(t *testing.T) {
	type args struct {
		posInteger int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimeTrialByError(tt.args.posInteger); got != tt.want {
				t.Errorf("IsPrimeTrialByError() = %v, want %v", got, tt.want)
			}
		})
	}
}
