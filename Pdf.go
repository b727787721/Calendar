package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"math"
	"unicode/utf8"
)

type Pdf struct {
}

const (
	PageWidth  = 110 //mm
	PageHeight = 170

	FirstUpcomingPrefixYPos     = 22
	FirstUpcomingPrefixFontSize = 16
	HighlightAreaYPos           = 32
	HighlightAreaMaxFontSize    = 99
	HighlightAreaMaxFontLen     = 3
	HighlightAreaHeight         = 50
	FirstUpcomingDiffYPos       = 57
	FirstUpcomingDiffFontSize   = 16
	DateYPos                    = 84
	DateFontSize                = 16
	SecondThirdUpcomingYPos     = 102
	SecondThirdUpcomingFontSize = 16
	OneLineHeight               = 16

	// Line para
	SplitLineX1 = 20
	SplitLineY1 = 103
	SplitLineX2 = 90
	SplitLineY2 = 103

	ImageYPos   = 123
	ImageXPos   = 45
	ImageWidth  = 49.7
	ImageHeight = 41.7

	Utf8Font      = "utf8"
	ZmxFont       = "zmx"
	ZmtFontFactor = 0.15
)

func (p *Pdf) GeneratedCalendarPdf(texts []Text) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font(Utf8Font, "", "./HYRuiYiSongW.ttf")
	pdf.AddUTF8Font(ZmxFont, "", "./HYShangWeiShouShuW.ttf")

	p.AddCoverPage(pdf)

	for _, text := range texts {
		pdf.AddPageFormat("P", gofpdf.SizeType{Wd: PageWidth, Ht: PageHeight})

		pdf.ImageOptions(
			"background.png",
			0, 0,
			PageWidth, PageHeight,
			false,
			gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
			0,
			"",
		)

		// Set first upcoming message prefix
		pdf.SetFont(Utf8Font, "", FirstUpcomingPrefixFontSize)
		firstUpcomingPrefixWidth := pdf.GetStringWidth(text.FirstUpcomingPrefix)
		pdf.SetXY((PageWidth-firstUpcomingPrefixWidth)/2, FirstUpcomingPrefixYPos)
		pdf.CellFormat(firstUpcomingPrefixWidth, OneLineHeight, text.FirstUpcomingPrefix, "0", 0, "CM", false, 0, "")

		// Set highlight area
		highlightAreaLen := float64(utf8.RuneCountInString(text.HighLightArea))
		highlightAreaFontSize := math.Min(HighlightAreaMaxFontSize, (HighlightAreaMaxFontSize*HighlightAreaMaxFontLen)/highlightAreaLen)
		pdf.SetFont(ZmxFont, "", highlightAreaFontSize)
		highLightAreaWidth := pdf.GetStringWidth(text.HighLightArea)
		pdf.SetXY((PageWidth-highLightAreaWidth)/2, HighlightAreaYPos)
		pdf.CellFormat(highLightAreaWidth, HighlightAreaHeight, text.HighLightArea, "0", 0, "CM", false, 0, "")

		// Set "day"
		pdf.SetFont(Utf8Font, "", FirstUpcomingDiffFontSize)
		firstUpcomingDiffWidth := pdf.GetStringWidth(text.FirstUpcomingDiff)
		pdf.SetXY((highlightAreaLen*highlightAreaFontSize*ZmtFontFactor+PageWidth)/2, FirstUpcomingDiffYPos)
		pdf.CellFormat(firstUpcomingDiffWidth, OneLineHeight, text.FirstUpcomingDiff, "0", 0, "CM", false, 0, "")

		// Set date
		pdf.SetFont(Utf8Font, "", DateFontSize)
		dateWidth := pdf.GetStringWidth(text.Date)
		pdf.SetXY((PageWidth-dateWidth)/2, DateYPos)
		pdf.CellFormat(dateWidth, OneLineHeight, text.Date, "0", 0, "CM", false, 0, "")

		// draw line
		pdf.SetDrawColor(217, 217, 217)
		pdf.Line(SplitLineX1, SplitLineY1, SplitLineX2, SplitLineY2)

		// reset color
		pdf.SetDrawColor(255, 255, 255)

		// Set second and third upcoming memorial days message
		pdf.SetFont(Utf8Font, "", SecondThirdUpcomingFontSize)
		secondUpcomingWidth := pdf.GetStringWidth(text.SecondUpcoming)
		pdf.SetXY((PageWidth-secondUpcomingWidth)/2, SecondThirdUpcomingYPos)
		pdf.CellFormat(secondUpcomingWidth, OneLineHeight, text.SecondUpcoming, "0", 1, "CM", false, 0, "")

		thirdUpcomingWidth := pdf.GetStringWidth(text.ThirdUpcoming)
		pdf.SetXY((PageWidth-thirdUpcomingWidth)/2, SecondThirdUpcomingYPos+6)
		pdf.CellFormat(thirdUpcomingWidth, OneLineHeight, text.ThirdUpcoming, "0", 0, "CM", false, 0, "")

		pdf.ImageOptions(
			"pic.png",
			(PageWidth-ImageWidth)/2, ImageYPos,
			ImageWidth, ImageHeight,
			false,
			gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
			0,
			"",
		)
	}

	err := pdf.OutputFileAndClose("./calendar.pdf")
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Pdf) AddCoverPage(pdf *gofpdf.Fpdf) {
	// Add front page
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: PageWidth, Ht: PageHeight})
	pdf.ImageOptions(
		"background.png",
		0, 0,
		PageWidth, PageHeight,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	pdf.SetFont(ZmxFont, "", 44)
	pdf.SetXY(65, 90)
	pdf.CellFormat(40, OneLineHeight, "山今历", "0", 0, "CM", false, 0, "")
	pdf.SetXY(65, 105)
	pdf.CellFormat(40, OneLineHeight, "2023", "0", 0, "CM", false, 0, "")

	// Add front page
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: PageWidth, Ht: PageHeight})
	pdf.ImageOptions(
		"background.png",
		0, 0,
		PageWidth, PageHeight,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	pdf.SetFont(ZmxFont, "", 44)
	pdf.SetXY(35, 60)
	pdf.CellFormat(40, OneLineHeight, "致我们的", "0", 0, "CM", false, 0, "")
	pdf.SetXY(35, 75)
	pdf.CellFormat(40, OneLineHeight, "第五年", "0", 0, "CM", false, 0, "")
}
