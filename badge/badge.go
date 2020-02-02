package badge

import (
	"io"

	svg "github.com/ajstarks/svgo"
)

// StatusBadge represent status
type StatusBadge struct {
	Text       string
	Color      string
	Status     string
	StatsColor string
	icon       bool
}

// CreateSvg return new created badge
func (b StatusBadge) CreateSvg(w io.Writer) {
	canvas := svg.New(w)

	canvas.Start(170, 20, "")

	aStyle := []svg.Offcolor{
		{Offset: 0, Opacity: 0, Color: "#000"},
		{Offset: 1, Opacity: 0.2, Color: "#000"},
	}

	canvas.LinearGradient("a", 0, 0, 0, 100, aStyle)

	canvas.ClipPath("id='c'")
	canvas.Rect(0, 0, 170, 20, "rx='2.8'")
	canvas.ClipEnd()

	canvas.Group("clip-path:url(#c)")
	canvas.Rect(0, 0, 170, 20, "fill:"+b.Color)
	canvas.Rect(111, 0, 72, 20, "fill:"+b.StatsColor)
	canvas.Rect(0, 0, 170, 20, "fill:url(#a)")
	canvas.Gend()

	if b.icon {
		canvas.Group()
		canvas.Circle(13, 9, 8, "fill:#80808094")
		canvas.Gend()
	}

	canvas.Group("fill:#fff;text-anchor:middle;font-family:DejaVu Sans,Verdana,Geneva,sans-serif;font-size:11")
	canvas.Text(67, 14, b.Text, "fill-opacity:0.3;fill:#000")
	canvas.Text(67, 14, b.Text, "fill:#fff")
	canvas.Text(135, 14, b.Status, "fill-opacity:0.3;fill:#000")
	canvas.Text(135, 14, b.Status, "fill:#fff")
	canvas.Gend()

	canvas.End()
}
