package kubectl

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// func TestInspectSubcommandInfo(args []string) (*SubcommandInfo, bool) {
func TestInspectSubcommandInfo(t *testing.T) {
	tests := []struct {
		args       string
		expected   *SubcommandInfo
		expectedOK bool
	}{
		{"get pods", &SubcommandInfo{Subcommand: Get}, true},
		{"get pod", &SubcommandInfo{Subcommand: Get}, true},
		{"get po", &SubcommandInfo{Subcommand: Get}, true},

		{"get pod -o wide", &SubcommandInfo{Subcommand: Get, FormatOption: Wide}, true},
		{"get pod -o=wide", &SubcommandInfo{Subcommand: Get, FormatOption: Wide}, true},
		{"get pod -owide", &SubcommandInfo{Subcommand: Get, FormatOption: Wide}, true},

		{"get pod -o json", &SubcommandInfo{Subcommand: Get, FormatOption: Json}, true},
		{"get pod -o=json", &SubcommandInfo{Subcommand: Get, FormatOption: Json}, true},
		{"get pod -ojson", &SubcommandInfo{Subcommand: Get, FormatOption: Json}, true},

		{"get pod -o yaml", &SubcommandInfo{Subcommand: Get, FormatOption: Yaml}, true},
		{"get pod -o=yaml", &SubcommandInfo{Subcommand: Get, FormatOption: Yaml}, true},
		{"get pod -oyaml", &SubcommandInfo{Subcommand: Get, FormatOption: Yaml}, true},

		{"get pod --no-headers", &SubcommandInfo{Subcommand: Get, NoHeader: true}, true},
		{"get pod -w", &SubcommandInfo{Subcommand: Get, Watch: true}, true},
		{"get pod --watch", &SubcommandInfo{Subcommand: Get, Watch: true}, true},
		{"get pod -h", &SubcommandInfo{Subcommand: Get, Help: true}, true},
		{"get pod --help", &SubcommandInfo{Subcommand: Get, Help: true}, true},

		{"describe pod pod-aaa", &SubcommandInfo{Subcommand: Describe}, true},
		{"top pod", &SubcommandInfo{Subcommand: Top}, true},
		{"top pods", &SubcommandInfo{Subcommand: Top}, true},

		{"", &SubcommandInfo{}, false},
		{"apply", &SubcommandInfo{}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.args, func(t *testing.T) {
			t.Parallel()
			s, ok := InspectSubcommandInfo(strings.Split(tt.args, " "))
			if tt.expectedOK != ok {
				t.Error("failed")
			}

			if diff := cmp.Diff(s, tt.expected); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
