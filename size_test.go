package atarimodels

import (
	"testing"
)

func TestSizeHasPoint(test *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	type args struct {
		point Point
	}
	type data struct {
		fields fields
		args   args
		want   bool
	}

	for _, data := range []data{
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				point: Point{4, 1},
			},
			want: true,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				point: Point{-1, 1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				point: Point{4, -1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				point: Point{-1, -1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				point: Point{10, 1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				point: Point{4, 10},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				point: Point{10, 10},
			},
			want: false,
		},
	} {
		size := Size{
			Width:  data.fields.Width,
			Height: data.fields.Height,
		}
		got := size.HasPoint(data.args.point)

		if got != data.want {
			test.Fail()
		}
	}
}
