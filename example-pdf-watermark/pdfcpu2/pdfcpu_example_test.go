package pdfcpu2

import "testing"

func TestAddWaterMarkToLocal(t *testing.T) {
	type args struct {
		inFile    string
		outFile   string
		waterMark string
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
				waterMark: "中文水印测试专用",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddWaterMarkToLocal(tt.args.inFile, tt.args.outFile, tt.args.waterMark); (err != nil) != tt.wantErr {
				t.Errorf("AddWaterMarkToLocal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
