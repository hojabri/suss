package suspiciousDetector

import (
	"github.com/hojabri/suss/pkg/entities"
	"testing"
)

func TestIsMovementSuspicious(t *testing.T) {
	type args struct {
		fromEvent *entities.Event
		toEvent   *entities.Event
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		wantSpeed float64
	}{
		{
			name:  "test1 - long distance",
			args:  args{
				fromEvent: &entities.Event{
					UnixTimestamp: 1619268188,
					Lat:           -34.8576,
					Lon:           -56.1702,
					Radius:        10,
				},
				toEvent:   &entities.Event{
					UnixTimestamp: 1619268288,
					Lat:           37.751,
					Lon:           -97.822,
					Radius:        100,
				},
			},
			want:      true,
			wantSpeed: 202081.97048861248,
		},
		{
			name:  "test2 - same point",
			args:  args{
				fromEvent: &entities.Event{
					UnixTimestamp: 1619268188,
					Lat:           0,
					Lon:           0,
					Radius:        0,
				},
				toEvent:   &entities.Event{
					UnixTimestamp: 0,
					Lat:           0,
					Lon:           0,
					Radius:        0,
				},
			},
			want:      false,
			wantSpeed: 0,
		},
		{
			name:  "test3 - near distance",
			args:  args{
				fromEvent: &entities.Event{
					UnixTimestamp: 1619268188,
					Lat:           59.434917,
					Lon:           24.752821,
					Radius:        0,
				},
				toEvent:   &entities.Event{
					UnixTimestamp: 1619268288,
					Lat:           59.439438,
					Lon:           24.753675,
					Radius:        0,
				},
			},
			want:      false,
			wantSpeed: 11.294950074341553,
		},
		{
			name:  "test4 - near distance but in a very short duration",
			args:  args{
				fromEvent: &entities.Event{
					UnixTimestamp: 1619268188,
					Lat:           59.434917,
					Lon:           24.752821,
					Radius:        0,
				},
				toEvent:   &entities.Event{
					UnixTimestamp: 1619268190,
					Lat:           59.439438,
					Lon:           24.753675,
					Radius:        0,
				},
			},
			want:      true,
			wantSpeed: 564.7475037170777,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsMovementSuspicious(tt.args.fromEvent, tt.args.toEvent)
			if got != tt.want {
				t.Errorf("IsMovementSuspicious() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantSpeed {
				t.Errorf("IsMovementSuspicious() got1 = %v, want %v", got1, tt.wantSpeed)
			}
		})
	}
}

func Test_kilometersToMiles(t *testing.T) {
	type args struct {
		km float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kilometersToMiles(tt.args.km); got != tt.want {
				t.Errorf("kilometersToMiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_milesToKilometers(t *testing.T) {
	type args struct {
		miles float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := milesToKilometers(tt.args.miles); got != tt.want {
				t.Errorf("milesToKilometers() = %v, want %v", got, tt.want)
			}
		})
	}
}
