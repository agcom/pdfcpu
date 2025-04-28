package types

import "testing"

func TestArray_PDFString(t *testing.T) {
	tests := []struct {
		name string
		a    Array
		want string
	}{
		{
			name: "Names",
			a:    NewNameArray("a", "b", "c"),
			want: "[/a /b /c]",
		},
		{
			name: "Name,Integer",
			a:    Array{Integer(0), Name("a")},
			want: "[0 /a]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.PDFString(); got != tt.want {
				t.Errorf("PDFString() = %v, want %v", got, tt.want)
			}
		})
	}
}
