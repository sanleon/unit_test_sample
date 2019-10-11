package ops

import "testing"



func Test_keyOp_Degenerate(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		wantX int
		wantY int
		wantErr bool
	} {
		{
			name: "success",
			args: args{
				"40_49",
			},
			wantX: 40,
			wantY: 49,
		},
		{
			name: "failure",
			args: args {
				"4099",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			kp := GetKeyOperator()
			gotX, gotY, gotErr := kp.Degenerate(tt.args.s)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("keyOp.Degenerate() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if gotX != tt.wantX {
				t.Errorf("keyOp.Degenerate() error = %v, wantX %v", gotErr, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("keyOp.Degenerate() error = %v, wantY %v", gotErr, tt.wantY)
			}
		})
	}
}

func Test_keyOp_Generate(t *testing.T) {
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				x:5,
				y:50,
			},
			want: "5_50",
		},
		{
			name: "success large integer",
			args: args {
				x: 50000,
				y: 999999,
			},
			want: "50000_999999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp := GetKeyOperator()
			if got := kp.Generate(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("keyOp.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
