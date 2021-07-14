package utils

import "testing"

func TestGenerateToken(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{"zhang"},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
