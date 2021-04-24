package haversine

import (
	"github.com/hojabri/suss/pkg/entities"
	"testing"
)

func TestDistance(t *testing.T) {
	type args struct {
		p entities.Geo
		q entities.Geo
	}
	tests := []struct {
		name   string
		args   args
		wantMi float64
		wantKm float64
	}{
		{
			name:   "test1",
			args:   args{
				p: entities.Geo{
					Lat:    0,
					Lon:    0,
					Radius: 0,
				},
				q: entities.Geo{
					Lat:    0,
					Lon:    0,
					Radius: 0,
				},
			},
			wantMi: 0,
			wantKm: 0,
		},
		{
			name:   "test2",
			args:   args{
				p: entities.Geo{
					Lat:    0,
					Lon:    0,
					Radius: 0,
				},
				q: entities.Geo{
					Lat:    52.464469,
					Lon:    10.116863,
					Radius: 0,
				},
			},
			wantMi: 3671.3211467204637,
			wantKm: 5909.546999938372,
		},
		{
			name:   "test3",
			args:   args{
				p: entities.Geo{
					Lat:    52.507473,
					Lon:    13.418760,
					Radius: 0,
				},
				q: entities.Geo{
					Lat:    52.464469,
					Lon:    10.116863,
					Radius: 0,
				},
			},
			wantMi: 138.91968767950044,
			wantKm: 223.61226129512315,
		},
		{
			name:   "test4",
			args:   args{
				p: entities.Geo{
					Lat:    52.507473,
					Lon:    13.418760,
					Radius: 0,
				},
				q: entities.Geo{
					Lat:    51.077037,
					Lon:    13.723125,
					Radius: 0,
				},
			},
			wantMi: 99.66651949252388,
			wantKm: 160.42834656060376,
		},
		{
			name:   "test5",
			args:   args{
				p: entities.Geo{
					Lat:    59.434917,
					Lon:    24.752821,
					Radius: 0,
				},
				q: entities.Geo{
					Lat:    51.077037,
					Lon:    13.723125,
					Radius: 0,
				},
			},
			wantMi: 720.4816963476937,
			wantKm: 1159.7243272943801,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMi, gotKm := Distance(tt.args.p, tt.args.q)
			if gotMi != tt.wantMi {
				t.Errorf("Distance() gotMi = %v, want %v", gotMi, tt.wantMi)
			}
			if gotKm != tt.wantKm {
				t.Errorf("Distance() gotKm = %v, want %v", gotKm, tt.wantKm)
			}
		})
	}
}

func Test_degreesToRadians(t *testing.T) {
	type args struct {
		d float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				34.7732,
			},
			want: 0.6069068314544922,
		},
		{
			name: "test2",
			args: args{
				113.722,
			},
			want: 1.9848233319529913,
		},
		{
			name: "test3",
			args: args{
				24.752821,
			},
			want: 0.4320182256067953,
		},
		{
			name: "test4",
			args: args{
				0,
			},
			want: 0,
		},
		{
			name: "test5",
			args: args{
				1000000000000,
			},
			want: 1.7453292519943295e+10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := degreesToRadians(tt.args.d); got != tt.want {
				t.Errorf("degreesToRadians() = %v, want %v", got, tt.want)
			}
		})
	}
}
