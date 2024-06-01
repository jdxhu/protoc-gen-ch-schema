package converter

import "testing"

func TestCheckFill(t *testing.T) {
	type args struct {
		tplString string
		kwargs    map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				tplString: "[a:{a}][|b:{b}]",
				kwargs:    map[string]string{"a": "test1", "b": "test2"},
			},
			wantRet: "a:test1|b:test2",
		},
		{
			name: "Test 2",
			args: args{
				tplString: "[a:{a}][|b:{b}]",
				kwargs:    map[string]string{"a": "test1"},
			},
			wantRet: "a:test1",
		},
		{
			name: "Test 3",
			args: args{
				tplString: "[a:{a}[|b: {b}]]",
				kwargs:    map[string]string{"a": "test1", "b": "test2"},
			},
			wantRet: "a:test1|b: test2",
		},
		{
			name: "Test 4",
			args: args{
				tplString: "[a:{a}[|b: {b}]]",
				kwargs:    map[string]string{"a": "test1"},
			},
			wantRet: "a:test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := CheckFill(tt.args.tplString, tt.args.kwargs); gotRet != tt.wantRet {
				t.Errorf("CheckFill() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
