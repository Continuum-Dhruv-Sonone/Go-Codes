package cipher

import "testing"

func TestDecrypt(t *testing.T) {
	type args struct {
		key     int
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"ALL CAPS",
			args{5, "FAIJWJSOOFAMAU"},
			"AVDERENJJAVHVP",
		},
		{
			"Key len greater than 26",
			args{31, "FAIJWJSOOFAMAU"},
			"AVDERENJJAVHVP",
		},
		{
			"Small alphabets",
			args{3, "Rozo"},
			"Olwl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.key, tt.args.message); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
