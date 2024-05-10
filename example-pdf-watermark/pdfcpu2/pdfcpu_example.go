package pdfcpu2

import (
	"regexp"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

// 字体大小points:12, 透明度opacity:0.6,旋转30°，字体名fontname，缩放因子scalefactor,文字对齐aligntext居中
const PdfFontConfig = "points:6,opacity:0.3,rot:30,fontname:SimHei,scalefactor:1.5,aligntext:c"
const OnTop = true

func AddWaterMarkToLocal(inFile string, outFile string, waterMark string) error {
	if len(outFile) < 1 {
		outFile = "watermark.pdf"
	}
	// // Load Config from local
	// // default set to file /usr/.config/config.yml
	// // if use this line, must use cli install font before: pdfcpu font install <path to the font file>	
	// api.LoadConfiguration()

	// load font from local file in code
	// set the config file path
	// this function will automatically create a 'pdfcpu' folder
	api.EnsureDefaultConfigAt("./")
	// install fonts from path
	err := api.InstallFonts([]string{"./simhei.ttf"})
	if err != nil {
		return err
	}

	// set the watermark configuration
	wm, err := pdfcpu.ParseTextWatermarkDetails(GetWaterMarkStr(waterMark), PdfFontConfig, OnTop, 1)
	if err != nil {
		return err
	}

	//add watermark and save file to path
	err = api.AddWatermarksFile(inFile, outFile, nil, wm, nil)
	if err != nil {
		return err
	}
	return nil
}

// function for mutilple lines text
func GetWaterMarkStr(waterMark string) string {
	count := 6
	pattern := "[a-zA-Z0-9\u4e00-\u9fa5]"
	re := regexp.MustCompile(pattern)
	textLen := len(re.FindAllString(waterMark, -1))
	switch {
	case textLen > 25 && textLen <= 40:
		count = 5
	case textLen > 40 && textLen <= 60:
		count = 4
	case textLen > 60:
		count = 3
	}
	var sb1 strings.Builder
	for i := 0; i < count; i++ {
		sb1.WriteString(waterMark)
		if i < count-1 {
			sb1.WriteString(strings.Repeat(" ", 10)) //文字间隔
		}
	}
	//单行文字制作完成
	sb1Str := sb1.String()
	//拼接成多行文字
	var sb2 strings.Builder
	rows := 20
	for i := 0; i < rows; i++ {
		if i%2 == 0 {
			sb2.WriteString(sb1Str)
		} else {
			sb2.WriteString(strings.Repeat("  ", textLen) + sb1Str)
		}
		if i < rows-1 {
			sb2.WriteString("\n \n \n \n \n \n \n \n")
		}
	}
	return sb2.String()
}