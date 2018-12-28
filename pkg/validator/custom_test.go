package validator

import "testing"

func TestValidateURL(t *testing.T) {
	type args struct {
		urls string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test 1 [Success]",
			args: args{
				urls: "http://10.164.6.10",
			},
			want: true,
		}, {
			name: "Test 2 [Fail]",
			args: args{
				urls: "http//10.164.6.10",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateURL(tt.args.urls)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
