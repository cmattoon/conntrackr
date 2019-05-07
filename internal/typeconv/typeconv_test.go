package typeconv

func TestHex2Int(t *testing.T) {
	for _, tc := range []struct {
		hexval string
		intval uint32
	}{
		{
			hexval: "FFFFFFFF",
			intval: 4294967295,
		},
		{
			hexval: "FFFF0000",
			intval: 4294901760,
		},
		{
			hexval: "0000FFFF",
			intval: 65535,
		},
		{
			hexval: "FFFF",
			intval: 65535,
		},
	} {
		if tc.intval != hex2int(tc.hexval) {
			t.Fail("conversion error")
		}
	}
}
