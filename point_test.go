package atarimodels

import (
	"reflect"
	"testing"
)

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
