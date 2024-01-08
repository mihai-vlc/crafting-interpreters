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
			name: "handles an empty program no characters",
			s:    scanner.NewScanner(""),
			want: []*scanner.Token{
				scanner.NewToken(scanner.TokenEOF, "EOF", scanner.NewPosition(1, 1)),
			},
			wantErr: false,
		},
		{
			name: "handles an empty program with spaces",
			s:    scanner.NewScanner("  \t \r\n"),
			want: []*scanner.Token{
				scanner.NewToken(scanner.TokenEOF, "EOF", scanner.NewPosition(2, 1)),
			},
			wantErr: false,
		},
		{
			name: "standard identifiers",
			s:    scanner.NewScanner("a b c"),
			want: []*scanner.Token{
				scanner.NewToken(scanner.TokenIdentifier, "a", scanner.NewPosition(1, 1)),
				scanner.NewToken(scanner.TokenIdentifier, "b", scanner.NewPosition(1, 3)),
				scanner.NewToken(scanner.TokenIdentifier, "c", scanner.NewPosition(1, 5)),
				scanner.NewToken(scanner.TokenEOF, "EOF", scanner.NewPosition(1, 5)),
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
