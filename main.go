package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")

	// Basic Grid
	drawGrid(pdf)

	err := pdf.OutputFileAndClose("second.pdf")
	if err != nil {
		panic(err)
	}
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	_, lineHeight := pdf.GetFontSize()
	for x := 0.0; x < w; x += (w / 20.0) {
		pdf.Line(x, 0, x, h)
		pdf.Text(x, lineHeight, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y += (w / 20.0) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}

// w, h := pdf.GetPageSize()
// fmt.Println(w, h)
// pdf.AddPage()

// // Basic Text
// pdf.MoveTo(0, 0)
// pdf.SetFont("arial", "B", 38)
// pdf.SetTextColor(180, 250, 180)
// _, lineHeight := pdf.GetFontSize()
// pdf.Text(0, lineHeight, "Hello World")
// pdf.MoveTo(0, lineHeight*2.0)

// pdf.SetFont("times", "", 16)
// pdf.SetTextColor(100, 100, 100)
// _, lineHeight = pdf.GetFontSize()
// pdf.MultiCell(0, lineHeight*1.5, "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", gofpdf.BorderNone, gofpdf.AlignRight, false)

// // Basic Shapes
// pdf.SetFillColor(180, 180, 250)
// pdf.SetDrawColor(250, 180, 180)
// pdf.Rect(10, 100, 100, 100, "FD")
// pdf.Circle(100, 200, 20.0, "F")
// pdf.Polygon([]gofpdf.PointType{
// 	{200, 250},
// 	{250, 300},
// 	{200, 350},
// 	{150, 300},
// }, "FD")

// // Basic Image
// pdf.ImageOptions("../transformingImages/img/in_199203691.jpg", 275, 275, 92, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
