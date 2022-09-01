package byte64

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: args{
				s: "----",
			},
			want: []byte{
				0x00, 0x00, 0x00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeAll(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(1024))
		if err != nil {
			t.Errorf("got err: %v", err)
		}

		raw := make([]byte, n.Int64())
		if _, err := rand.Read(raw); err != nil {
			t.Errorf("got err: %v", err)
		}
		fmt.Println(len(raw), raw)
		s := Encode(raw)
		fmt.Println(s)
		target := Decode(s)
		fmt.Println(len(target), target)
		if !reflect.DeepEqual(target, raw) {
			t.Errorf("raw = %v, target = %v", raw, target)
		}
	}
}

func BenchmarkDecodeAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw := make([]byte, 1024)
		_, _ = rand.Read(raw)
		s := Encode(raw)
		target := Decode(s)
		if !reflect.DeepEqual(target, raw) {
			b.Errorf("raw = %v, target = %v", raw, target)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("EFojH2hWzezqkUwmYw_rEz7V5ViwtCKJcq2xvUVRWnN20INEl_1sd7J2ZojCsDbdZcQhi4Xre52GqRbKkNPdhVJBh6bW-CtgTK01x72bVWI03qQPAlO0ZMp40sHKtXaZl-7UiuicVTZRgch_0hinOKwP9_fx2hRUw2cHGxZaq9T7w6C7QlkA8nnCfsPMGRDkrkYSFMas3r4ls0Zqfu557HDIDOl4sv_D7AkZCtJovr9ArwYojYnllUwW3XSPTpsjcK1d7JR85J60TXNBhWZMtyrSiVJ2r4HPnBskL8plN10Phi2WsGlRKCxTjEkhfzLjUNMrbqWRShXT62A8AvYKT3fHTc7PuL3Q_hsAAG2YwoldV5Vkz8p3LqURhDmvLVndd_yNBYP25SAI9GPbtMAt3hR6xHIr48vG-Xc-TjH1OlvG9ULmRRSTtlQ-lAKqmo_36Nhb3PDDmF5CIMD8F-P5OZNQAb5L-ORcTRPyuhSHNKyrw0ww1bmnnjE7IGj-nmFYLigHePf_VxDryX5lY-UwGkkeTWrRUJ4brd0IVvuegUpZ1lctlQqzoWkH9auvDInoxhe_mYf3m4EJhS6kx_tmcgjIxt0uTh5iXxefilVW_t-6A-lFWBirl23dtUQUAcJ04T3WkEC2MVzEyFuyrO0XFeRNv-mijGDAviIDnHq-GUEALviL4xbay7DBttwcj1RxTuuKs0WSaTFzha48hA34Yh3Zff3hv8hUYCc_hYFy64nymR50faacD8TFjcjCrIb0Lp7-u17ff4cOIuuQuTBV0Ic8U7cCl3cLiVMfvzAzwe51Hkui0PdN7M8TfK8dwO2dMGPJ7zMETELfw7bJGO8Sgtx86z5Mu6X0J74UfwHzDu2kLeddQd7l6O8YMjno8T2P-qsmbbKxq3MnaatwdLYUCIOgsjDjYcG7Q5fbJ2lBplGHbop5tsrveX7KnVAViGG5MmESNI3OcYs6kjSPO_iDQDW75uX8Z0FxA3PvF8wbDk6Vcbv11Hlw78cXj2DpBe_9sSeKlR7jcm9myqwQW9kTMn0GxJRbiZE-Z8SFnRVdo8JCezjiThmu-rznZr1-MccYrrN99RPQwZUC8MY6bK0dccm3jLcKLgnb-GTMCNPQ-BlS05ISOXklTrJo2enVn3jD7KQXZCVEX-hwakVkTzxf7xmJwBKP6-OPp5pPh_xE6O-sz5-thop4NsH_OK6cOhVLEIuNVk7s-ehAcq0Jf9JahA5gW6N4HIqN0xd5EOviCq-6fT72Hjxa4RWCwxcA1wm4O2RPOhtpzrl6f4Go8TAxHL-EcqVQU990-udpGhoSnDH3x6FyKZ8xhKQktXhysYqq-EzQ9lfhlhx-OeloRTvgn1")
	}
}
