package atarimodels

import (
	"reflect"
	"testing"
)

func TestPointIsNil(test *testing.T) {
	type fields struct {
		column int
		row    int
	}
	type data struct {
		fields fields
		want   bool
	}

	for _, data := range []data{
		data{
			fields: fields{
				column: 2,
				row:    3,
			},
			want: false,
		},
		data{
			fields: fields{
				column: -1,
				row:    -1,
			},
			want: true,
		},
	} {
		point := Point{
			Column: data.fields.column,
			Row:    data.fields.row,
		}
		got := point.IsNil()

		if got != data.want {
			test.Fail()
		}
	}
}

func TestPointTranslate(test *testing.T) {
	type fields struct {
		column int
		row    int
	}
	type args struct {
		translation Point
	}
	type data struct {
		fields fields
		args   args
		want   Point
	}

	for _, data := range []data{
		data{
			fields: fields{
				column: 2,
				row:    3,
			},
			args: args{
				translation: Point{
					Column: 4,
					Row:    2,
				},
			},
			want: Point{
				Column: 6,
				Row:    5,
			},
		},
		data{
			fields: fields{
				column: 2,
				row:    3,
			},
			args: args{
				translation: Point{
					Column: -4,
					Row:    -2,
				},
			},
			want: Point{
				Column: -2,
				Row:    1,
			},
		},
	} {
		point := Point{
			Column: data.fields.column,
			Row:    data.fields.row,
		}
		got :=
			point.Translate(data.args.translation)

		if !reflect.DeepEqual(
			got,
			data.want,
		) {
			test.Fail()
		}
	}
}
