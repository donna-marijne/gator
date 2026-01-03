package main

import (
	"testing"
	"time"
)

func Test_parseTime(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s       string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "hnrss.org",
			s:       "Fri, 02 Jan 2026 18:57:24 +0000",
			want:    time.Date(2026, time.January, 2, 18, 57, 24, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "techcrunch.com",
			s:       "Fri, 02 Jan 2026 20:13:31 +0000",
			want:    time.Date(2026, time.January, 2, 20, 13, 31, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "ycombinator.com",
			s:       "Sat, 03 Jan 2026 04:01:03 +0000",
			want:    time.Date(2026, time.January, 3, 4, 1, 3, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "boot.dev",
			s:       "Thu, 30 Oct 2025 00:00:00 +0000",
			want:    time.Date(2025, time.October, 30, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "empty string",
			s:       "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := parseTime(tt.s)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parseTime() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parseTime() succeeded unexpectedly")
			}
			if !got.Equal(tt.want) {
				t.Errorf("parseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
