/**
 * Created by zhouwenzhe on 2023/5/8
 */

package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"自定义",
			args{
				s: "PAYPALISHIRINGAJLKDLAKDAFJLKDB",
				/*
					P0				H8				L16				F
					A1			S7	I9			J15	K17			A	J
					Y2		I6		R10		A14		D18		L		L
					P3	L5			I11	G13			L19	D			K	B
					A4				N12				K20				D
				*/
				numRows: 5,
			},
			"PHLFASIJKAJYIRADDLPLIGLKKBANAD"},
		{"示例 1",
			args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			"PAHNAPLSIIGYIR"},
		{"示例 2",
			args{
				s: "PAYPALISHIRING",
				/*
					P0			I6			N12
					A1		L5	S7		I11 G13
					Y2	A4   	H8	R10
					P3     		I9
				*/
				numRows: 4,
			},
			"PINALSIGYAHRPI"}, {"示例 3",
			args{
				s:       "A",
				numRows: 1,
			},
			"A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
