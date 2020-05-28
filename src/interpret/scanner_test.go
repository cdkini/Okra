package interpret

import (
	"testing"
)

func TestScanWhitespace(t *testing.T) {
	table := []struct {
		in   string
		out  []TokenType
		line int
	}{
		{} // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanSingleCharTokens(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanDoubleCharTokens(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanComments(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanString(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanNumeric(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanIdentifierAndKeyword(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}
