package scanner_test

import (
	"lox/internal/scanner"
	"reflect"
	"testing"
)

func TestScanner_ScanTokens(t *testing.T) {
	tests := []struct {
		name    string
		s       *scanner.Scanner
		want    []*scanner.Token
		wantErr bool
	}{
		{
			name: "handles an empty program",
			s:    scanner.NewScanner(""),
			want: []*scanner.Token{
				scanner.NewToken(scanner.TokenEOF, "EOF", scanner.NewPosition(1, 1)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ScanTokens()
			if (err != nil) != tt.wantErr {
				t.Errorf("Scanner.ScanTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.ScanTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
