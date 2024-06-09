package designparkingsystem

import "testing"

func TestParkingSystem_AddCar(t *testing.T) {
	type args struct {
		carType int
	}
	tests := []struct {
		name string
		lots [3]int
		args args
		want bool
	}{
		{
			lots: [3]int{1, 0, 0},
			args: args{
				carType: 1,
			},
			want: true,
		},
		{
			lots: [3]int{0, 0, 0},
			args: args{
				carType: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := Constructor(tt.lots[0], tt.lots[1], tt.lots[2])
			if got := sut.AddCar(tt.args.carType); got != tt.want {
				t.Errorf("ParkingSystem.AddCar() = %v, want %v", got, tt.want)
			}
		})
	}
}
