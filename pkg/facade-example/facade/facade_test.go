package facade

import (
	"fmt"
	informator2 "github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/informator"
	user2 "github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/user"
	validator2 "github.com/VitalyDorozhkin/go-patterns/pkg/facade-example/validator"
	"reflect"
	"testing"
)

const (
	testName                   = "Test facade"
	methodCheckNameLength      = "CheckNameLength"
	methodCheckLastNameLength  = "CheckLastNameLength"
	methodCheckAge             = "CheckAge"
	methodInformNameLength     = "InformNameLength"
	methodInformLastNameLength = "InformLastNameLength"
	methodInformAge            = "InformAge"
	successName                = "Vanyaa"
	successLastname            = "Ivanov"
	successAge                 = 8
	failedName                 = "Van"
	failedLastname             = "Ivan"
	failedAge                  = 2
	min                        = 5
	max                        = 10
	badNameMessage = "bad name"
	badLastNameMessage = "bad lastname"
	badAgeMessage = "bad age"
)

func Test_userCreator_NewUser(t *testing.T) {

	type args struct {
		name     string
		lastname string
		age      int
	}

	tests := []struct {
		name     string
		args     args
		wantUser user
		wantErr  []error
	}{
		{
			name: "all valid",
			args: args{
				name:     successName,
				lastname: successLastname,
				age:      successAge,
			},
			wantUser: user2.NewFactory().NewUser(successName, successLastname, successAge),
			wantErr:  nil,
		},
		{
			name: "all invalid",
			args: args{
				name:     failedName,
				lastname: failedLastname,
				age:      failedAge,
			},
			wantUser: nil,
			wantErr: []error{
				fmt.Errorf(badNameMessage),
				fmt.Errorf(badLastNameMessage),
				fmt.Errorf(badAgeMessage),
			},
		},
		{
			name: "name invalid",
			args: args{
				name:     failedName,
				lastname: successLastname,
				age:      successAge,
			},
			wantUser: nil,
			wantErr: []error{
				fmt.Errorf(badNameMessage),
			},
		},
		{
			name: "lastname invalid",
			args: args{
				name:     successName,
				lastname: failedLastname,
				age:      successAge,
			},
			wantUser: nil,
			wantErr: []error{
				fmt.Errorf(badLastNameMessage),
			},
		},
		{
			name: "age invalid",
			args: args{
				name:     successName,
				lastname: successLastname,
				age:      failedAge,
			},
			wantUser: nil,
			wantErr: []error{
				fmt.Errorf(badAgeMessage),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validationMock := new(validator2.Mock)

			validationMock.On(methodCheckNameLength, successName).Return(true, min, max).Once()
			validationMock.On(methodCheckLastNameLength, successLastname).Return(true, min, max).Once()
			validationMock.On(methodCheckAge, successAge).Return(true, min, max).Once()

			validationMock.On(methodCheckNameLength, failedName).Return(false, min, max).Once()
			validationMock.On(methodCheckLastNameLength, failedLastname).Return(false, min, max).Once()
			validationMock.On(methodCheckAge, failedAge).Return(false, min, max).Once()

			informationMock := new(informator2.Mock)

			informationMock.On(methodInformNameLength, failedName, min, max).Return(badNameMessage).Once()
			informationMock.On(methodInformLastNameLength, failedLastname, min, max).Return(badLastNameMessage).Once()
			informationMock.On(methodInformAge, failedAge, min, max).Return(badAgeMessage).Once()

			u := NewUserCreator(user2.NewFactory(), validationMock, informationMock)

			gotUser, gotErr := u.NewUser(tt.args.name, tt.args.lastname, tt.args.age)
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("NewUser() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("NewUser() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
