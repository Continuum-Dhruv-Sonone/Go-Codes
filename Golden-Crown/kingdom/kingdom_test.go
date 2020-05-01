package kingdom

import "testing"

func TestGetEmblem(t *testing.T) {
	Setup()
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Success GetEmblem",
			args{"LAND"},
			"Panda",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEmblem(tt.args.name); got != tt.want {
				t.Errorf("GetEmblem() = %v, want %v", got, tt.want)
			}
		})
	}
}
