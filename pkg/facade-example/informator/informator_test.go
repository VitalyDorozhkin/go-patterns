package informator

import (
	"testing"
)

var i = &informator{
	maxNumberPattern: "%d is too big! max %d",
	minNumberPattern: "%d is too small! min %d",
	maxLengthPattern: "'%s' is too long! max %d",
	minLengthPattern: "'%s' is too short! min %d",
}

func Test_informator_InformAge(t *testing.T) {
	type args struct {
		age int
		min int
		max int
	}
	tests := []struct {
		name     string
		args     args
		wantInfo string
	}{
		{
			name: "middle",
			args: args{4, 2, 6},
			wantInfo: "Age: ok",
		},
		{
			name: "bigger",
			args: args{7, 2, 6},
			wantInfo: "Age: 7 is too big! max 6",
		},
		{
			name: "smaller",
			args: args{1, 2, 6},
			wantInfo: "Age: 1 is too small! min 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInfo := i.InformAge(tt.args.age, tt.args.min, tt.args.max); gotInfo != tt.wantInfo {
				t.Errorf("InformAge() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func Test_informator_InformLastNameLength(t *testing.T) {
	type args struct {
		lastname string
		min int
		max int
	}
	tests := []struct {
		name     string
		args     args
		wantInfo string
	}{
		{
			name: "middle",
			args: args{"ab", 2, 6},
			wantInfo: "LastName Length: ok",
		},
		{
			name: "longer",
			args: args{"abcdefghi", 2, 6},
			wantInfo: "LastName Length: 'abcdefghi' is too long! max 6",
		},
		{
			name: "shorter",
			args: args{"A", 2, 6},
			wantInfo: "LastName Length: 'A' is too short! min 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInfo := i.InformLastNameLength(tt.args.lastname, tt.args.min, tt.args.max); gotInfo != tt.wantInfo {
				t.Errorf("InformLastNameLength() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func Test_informator_InformNameLength(t *testing.T) {
	type args struct {
		name string
		min int
		max int
	}
	tests := []struct {
		name     string
		args     args
		wantInfo string
	}{
		{
			name: "middle",
			args: args{"ab", 2, 6},
			wantInfo: "Name Length: ok",
		},
		{
			name: "longer",
			args: args{"abcdefghi", 2, 6},
			wantInfo: "Name Length: 'abcdefghi' is too long! max 6",
		},
		{
			name: "shorter",
			args: args{"A", 2, 6},
			wantInfo: "Name Length: 'A' is too short! min 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInfo := i.InformNameLength(tt.args.name, tt.args.min, tt.args.max); gotInfo != tt.wantInfo {
				t.Errorf("InformNameLength() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
