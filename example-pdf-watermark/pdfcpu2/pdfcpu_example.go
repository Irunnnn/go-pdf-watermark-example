package pdfcpu2

import (
	"regexp"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

// fontsize-points:6, opacity:0.3,rotate:30°，fontname:simhei，scalefactor: 1.5 abd,text alignment-aligntext:center
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
	textLen := len(waterMark)
	// automatically calculate the line count
	count :=  int(math.Round(250.0 / (float64(textLen) + 10.0)))
	var sb1 strings.Builder
	for i := 0; i < count; i++ {
		sb1.WriteString(waterMark)
		if i < count-1 {
			sb1.WriteString(strings.Repeat(" ", 10)) //文字间隔
		}
	}
	//single line
	sb1Str := sb1.String()
	//multiple lines
	var sb2 strings.Builder
	lineSpace := "\n \n \n \n \n \n \n"
	height := float64(len(lineSpace) / 2.0)
	rows := int(math.Round(100.0 / height))
	for i := 0; i < rows; i++ {
		if i%2 == 0 {
			sb2.WriteString(sb1Str)
		} else {
			sb2.WriteString(strings.Repeat("  ", 5) + sb1Str[:len(sb1Str)-5])
		}
		if i < rows-1 {
			sb2.WriteString(lineSpace)
		}
	}
	return sb2.String()
}
