package utils

import (
	"reflect"
	"testing"

	lettercombinations "github.com/mekitmedia/leetcode/17_letterCombinations"
)

func Test_dfs(t *testing.T) {
	type args struct {
		nodes   []string
		adjList map[string][]string
		perf    string
		pos     int
		res     *[]string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantRes *[]string
	}{
		{
			name: "letterCombinations",
			args: args{
				nodes:   []string{"2"},
				adjList: lettercombinations.KeyboardLetters(),
				perf:    "",
				pos:     0,
				res:     &[]string{},
			},
			want:    false,
			wantRes: &[]string{"a", "b", "c"},
		}, {
			name: "letterCombinations",
			args: args{
				nodes:   []string{"2", "3"},
				adjList: lettercombinations.KeyboardLetters(),
				perf:    "",
				pos:     0,
				res:     &[]string{},
			},
			want:    false,
			wantRes: &[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.nodes, tt.args.adjList, tt.args.perf, tt.args.pos, tt.args.res); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.args.res, tt.wantRes) {
				t.Errorf("dfs() = %v, want %v", tt.args.res, tt.wantRes)
			}
		})
	}
}

func Test_eddgeListToAdjList(t *testing.T) {
	tests := []struct {
		name string
		edge [][]string
		adj  map[string][]string
	}{
		{
			name: "simple test: 1-2",
			edge: [][]string{{"1", "2"}},
			adj: map[string][]string{
				"1": {"2"},
			},
		}, {
			name: "1-23 2-3",
			edge: [][]string{{"1", "2"}, {"1", "3"}, {"2", "3"}},
			adj: map[string][]string{
				"1": {"2", "3"},
				"2": {"3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := edgeListToAdjList(tt.edge)

			if !reflect.DeepEqual(got, tt.adj) {
				t.Errorf("eddgeListToAdjList() = %v, want %v", got, tt.adj)
			}
		})
	}
}
