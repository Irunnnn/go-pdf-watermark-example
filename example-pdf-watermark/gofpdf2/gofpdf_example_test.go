package gofpdf2

import "testing"

func TestAddWaterMark(t *testing.T) {
	type args struct {
		inFile    string
		outFile   string
		watermark string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Chinese watermark test",
			args:    args{
				inFile:    "./example.pdf",
				outFile:   "",
				watermark: "中文水印测试专用",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddWaterMark(tt.args.inFile, tt.args.outFile, tt.args.watermark); (err != nil) != tt.wantErr {
				t.Errorf("AddWaterMark() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
