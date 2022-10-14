package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMobile(t *testing.T) {
	tests := []struct {
		name    string
		mobile  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid mobile",
			mobile:  "254712345678",
			want:    "254712345678",
			wantErr: false,
		},
		{
			name:    "valid mobile with +",
			mobile:  "+254712345678",
			want:    "254712345678",
			wantErr: false,
		},
		{
			name:    "valid mobile with spaces",
			mobile:  "254 712 345 678",
			want:    "254712345678",
			wantErr: false,
		},
		{
			name:    "valid mobile with spaces and +",
			mobile:  "+254 712 345 678",
			want:    "254712345678",
			wantErr: false,
		},
		{
			name:    "valid mobile with -",
			mobile:  "254-712-345-678",
			want:    "254712345678",
			wantErr: false,
		},
		{
			name:    "valid mobile starting 0",
			mobile:  "0712345678",
			want:    "254712345678",
			wantErr: false,
		},
		{
			name:    "valid mobile starting 7",
			mobile:  "712345678",
			want:    "254712345678",
			wantErr: false,
		},
		{
			name:    "invalid mobile",
			mobile:  "25471234567",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMobile(tt.mobile)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, got)
		})
	}
}
