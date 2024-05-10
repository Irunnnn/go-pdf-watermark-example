package gofpdf2

import (
	"os"
	"regexp"
	"strings"

	"github.com/signintech/gopdf"
)

func AddWaterMark(inFile string, outFile string, watermark string) error {
	if len(outFile) < 1 {
		outFile = "watermark.pdf"
	}
	
	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	pdf.AddPage()

	err := pdf.AddTTFFont("SimHei", "./simhei.ttf")
	if err != nil {
		return err
	}

	err = pdf.SetFont("SimHei", "", 14)

	if err != nil {
		return err
	}

	f, err := os.Open(inFile)
	if err != nil{
		return err
	}

	defer f.Close()

	for i := 0; i < 8; i++{
		// Import page 1
		tpl1 := pdf.ImportPage("example.pdf", i+1, "/MediaBox")

		// Draw pdf onto page
		pdf.UseImportedTemplate(tpl1, 0, 0, 595.28, 841.89)

		pdf.Rotate(-30.0, 595.0, 0.0)
		for k := 0; k < 30; k ++{
			pdf.SetGrayFill(0.4)
			pdf.SetFontSize(10)
			option := gopdf.CellOption{Align:gopdf.Center ,Float:gopdf.Bottom,Transparency: &gopdf.Transparency{Alpha: 0.3}}
			wts := GetWaterMarkStr(watermark)
			pdf.CellWithOption(nil, strings.Repeat(" ", 8) + wts, option)
			pdf.SetXY(0, pdf.GetY() + 50)
			option = gopdf.CellOption{Align:gopdf.Center ,Float:gopdf.Bottom,Transparency: &gopdf.Transparency{Alpha: 0.3}}
			pdf.CellWithOption(nil, wts, option)
			pdf.SetXY(0, pdf.GetY() + 50)
		}
		if i < 7{
			pdf.AddPage()
		}
	}

	pdf.WritePdf(outFile)
	return nil

}

func GetWaterMarkStr(waterMark string) string {
	count := 8
	pattern := "[a-zA-Z0-9\u4e00-\u9fa5]"
	re := regexp.MustCompile(pattern)
	textLen := len(re.FindAllString(waterMark, -1))
	switch {
	case textLen > 25 && textLen <= 40:
		count = 7
	case textLen > 40 && textLen <= 60:
		count = 6
	case textLen > 60:
		count = 5
	}
	var sb1 strings.Builder
	for i := 0; i < count; i++ {
		sb1.WriteString(waterMark)
		if i < count-1 {
			sb1.WriteString(strings.Repeat(" ", 15)) //文字间隔
		}
	}
	//单行文字制作完成
	sb1Str := sb1.String()
	return sb1Str
}