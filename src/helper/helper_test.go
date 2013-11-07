package helper

import (
	"net/url"
	"testing"
)

func TestParseUrl(t *testing.T) {
	tests := []struct {
		url           string
		expectedError error
		expectedKey   []string
		expectedValue float32
		expectedTime  int64
	}{
		{
			"/a/valid/key/1383844036/2",
			nil,
			[]string{"a", "valid", "key"},
			2,
			1383844036,
		},
	}

	for i, tt := range tests {
		u, e := url.ParseRequestURI(tt.url)
		if e != nil {
			t.Fatalf("Bad test url in test %d", i)
		}

		key, time, value, err := ParseUrl(u)
		if err != tt.expectedError {
			t.Fatalf("Expected err == %v got %v instead", tt.expectedError, err)
		}

		if time != tt.expectedTime {
			t.Fatalf("Expected time == %v got %v instead", tt.expectedTime, time)
		}

		if value != tt.expectedValue {
			t.Fatalf("Expected value == %v got %v instead", tt.expectedValue, value)
		}

		for j, v := range key {
			if v != tt.expectedKey[j] {
				t.Fatalf("Expected value == %v got %v instead", tt.expectedKey, key)
			}
		}
	}
}
