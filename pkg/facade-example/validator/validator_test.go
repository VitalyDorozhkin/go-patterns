package validator

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/models"
	"testing"
)

var v = &validator{
	nameLength:     models.Interval{Min: 2, Max: 6},
	lastnameLength: models.Interval{Min: 2, Max: 6},
	age:            models.Interval{Min: 2, Max: 6},
}

func Test_validator_CheckAge(t *testing.T) {
	type args struct {
		age int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 int
	}{
		{
			name:  "long",
			args:  args{age: 7},
			want:  false,
			want1: 2,
			want2: 6,
		},
		{
			name:  "short",
			args:  args{age: 1},
			want:  false,
			want1: 2,
			want2: 6,
		},
		{
			name:  "middle",
			args:  args{age: 4},
			want:  true,
			want1: 2,
			want2: 6,
		},
		{
			name:  "max",
			args:  args{age: 6},
			want:  true,
			want1: 2,
			want2: 6,
		},
		{
			name:  "min",
			args:  args{age: 2},
			want:  true,
			want1: 2,
			want2: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := v.CheckAge(tt.args.age)
			if got != tt.want {
				t.Errorf("CheckAge() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckAge() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CheckAge() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_validator_CheckLastNameLength(t *testing.T) {
	type args struct {
		lastname string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 int
	}{
		{
			name:  "long",
			args:  args{lastname: "abcdefg"},
			want:  false,
			want1: 2,
			want2: 6,
		},
		{
			name:  "short",
			args:  args{lastname: "a"},
			want:  false,
			want1: 2,
			want2: 6,
		},
		{
			name:  "middle",
			args:  args{lastname: "abcd"},
			want:  true,
			want1: 2,
			want2: 6,
		},
		{
			name:  "max",
			args:  args{lastname: "abcdef"},
			want:  true,
			want1: 2,
			want2: 6,
		},
		{
			name:  "min",
			args:  args{lastname: "ab"},
			want:  true,
			want1: 2,
			want2: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := v.CheckLastNameLength(tt.args.lastname)
			if got != tt.want {
				t.Errorf("CheckLastNameLength() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckLastNameLength() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CheckLastNameLength() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_validator_CheckNameLength(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 int
	}{
		{
			name:  "long",
			args:  args{name: "abcdefg"},
			want:  false,
			want1: 2,
			want2: 6,
		},
		{
			name:  "short",
			args:  args{name: "a"},
			want:  false,
			want1: 2,
			want2: 6,
		},
		{
			name:  "middle",
			args:  args{name: "abcd"},
			want:  true,
			want1: 2,
			want2: 6,
		},
		{
			name:  "max",
			args:  args{name: "abcdef"},
			want:  true,
			want1: 2,
			want2: 6,
		},
		{
			name:  "min",
			args:  args{name: "ab"},
			want:  true,
			want1: 2,
			want2: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := v.CheckNameLength(tt.args.name)
			if got != tt.want {
				t.Errorf("CheckNameLength() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckNameLength() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CheckNameLength() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
