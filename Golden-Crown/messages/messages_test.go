package messages

import (
	"testing"

	"github.com/ContinuumLLC/GO/Golden-Crown/kingdom"
)

func TestProcessMessages(t *testing.T) {
	kingdom.Setup()
	type args struct {
		inputs [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"RULER",
			args{
				inputs: [][]string{{"AIR", "ROZO"}, {"LAND", "FAIJWJSOOFAMAU"}, {"ICE", "STHSTSTVSASOS"}},
			},
			"SPACE AIR LAND ICE",
		},
		{
			"NOT RULER",
			args{
				inputs: [][]string{{"AIR", "OWLAOWLBOWLC"}, {"LAND", "OFBBMUFDICCSO"}, {"ICE", "VTBTBHTBBBOBAB"}, {"WATER", "SUMMERISCOMING"}},
			},
			"NONE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessMessages(tt.args.inputs); got != tt.want {
				t.Errorf("ProcessMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compare(t *testing.T) {
	type args struct {
		emblem  map[int]int
		message map[int]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Length is not same",
			args{
				map[int]int{1: 1},
				map[int]int{1: 1, 2: 2},
			},
			false,
		},
		{
			"Key is not present",
			args{
				map[int]int{1: 1},
				map[int]int{2: 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.args.emblem, tt.args.message); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
