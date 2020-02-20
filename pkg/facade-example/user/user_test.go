package user

import (
	"testing"
)

var f = NewFactory()

func Test_User_Info(t *testing.T) {
	type fields struct {
		name     string
		lastname string
		age      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Ivan",
			fields: fields{
				name:     "Ivan",
				lastname: "Ivanov",
				age:      10,
			},
			want: "user: {\nname: Ivan,\nlastname: Ivanov,\nage: 10\n}",
		},
		{
			name: "Petr",
			fields: fields{
				name:     "Petr",
				lastname: "Petrov",
				age:      18,
			},
			want: "user: {\nname: Petr,\nlastname: Petrov,\nage: 18\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := f.NewUser(tt.fields.name, tt.fields.lastname, tt.fields.age)
			if got := u.Info(); got != tt.want {
				t.Errorf("Info() = %v, want %v", got, tt.want)
			}
		})
	}
}
