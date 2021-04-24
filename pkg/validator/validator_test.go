package validator

import (
	"testing"
)

func TestPattern(t *testing.T) {
	type args struct {
		data    string
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1 - valid email",
			args:    args{
				data:    "o.hojabri@gmail.com",
				pattern: "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$",
			},
			wantErr: false,
		},
		{
			name:    "test2 - invalid email",
			args:    args{
				data:    "o.hojabri@gmailcom",
				pattern: "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$",
			},
			wantErr: true,
		},
		{
			name:    "test3 - invalid pattern",
			args:    args{
				data:    "o.hojabri@gmail.com",
				pattern: "^[\\w-\\.]+@([\\w-+\\.)+[\\w-]{2,4}$",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Pattern(tt.args.data, tt.args.pattern); (err != nil) != tt.wantErr {
				t.Errorf("Pattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


