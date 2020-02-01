package badge

import (
	"io"

	svg "github.com/ajstarks/svgo"
)

// Create return necanvas badge
func Create(w io.Writer) {
	canvas := svg.New(w)

	canvas.Start(184, 20, "")

	aStyle := []svg.Offcolor{
		{Offset: 0, Opacity: 0, Color: "#000"},
		{Offset: 1, Opacity: 0.2, Color: "#000"},
	}

	canvas.LinearGradient("a", 0, 0, 0, 100, aStyle)

	canvas.ClipPath("id='c'")
	canvas.Rect(0, 0, 184, 20)
	canvas.ClipEnd()

	canvas.Group("clip-path:url(#c)")
	canvas.Rect(0, 0, 184, 20, "fill:#555555")
	canvas.Rect(111, 0, 72, 20, "fill:#007ACC")
	canvas.Rect(0, 0, 184, 20, "fill:url(#a)")
	canvas.Gend()

	canvas.Group("fill:#fff;text-anchor:middle;font-family:DejaVu Sans,Verdana,Geneva,sans-serif;font-size:11")
	canvas.Text(64, 15, "Maid chan", "fill-opacity:0.3;fill:#000")
	canvas.Text(64, 15, "Maid chan", "fill:#fff")
	canvas.Text(147, 15, "Online", "fill-opacity:0.3;fill:#000")
	canvas.Text(147, 15, "Online", "fill:#fff")
	canvas.Gend()

	canvas.End()
}
