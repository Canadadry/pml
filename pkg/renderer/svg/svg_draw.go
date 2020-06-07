package svg

import (
	"pml/pkg/renderer/svg/svgpath"
)

func (sn *svgNode) draw(d Drawer) error {

	position := svgpath.Point{}

	for _, cmd := range sn.commands {
		switch cmd.Kind {
		case 'M':
			position.X = cmd.Points[0].X
			position.Y = cmd.Points[0].Y
			d.MoveTo(position.X, position.Y)
		case 'm':
			position.X += cmd.Points[0].X
			position.Y += cmd.Points[0].Y
			d.MoveTo(position.X, position.Y)
		case 'L':
			position.X = cmd.Points[0].X
			position.Y = cmd.Points[0].Y
			d.LineTo(position.X, position.Y)
		case 'H':
			position.X = cmd.Points[0].X
			d.LineTo(position.X, position.Y)
		case 'h':
			position.X += cmd.Points[0].X
			d.LineTo(position.X, position.Y)
		case 'V':
			position.Y = cmd.Points[0].X
			d.LineTo(position.X, position.Y)
		case 'v':
			position.Y += cmd.Points[0].X
			d.LineTo(position.X, position.Y)
		case 'l':
			position.X += cmd.Points[0].X
			position.Y += cmd.Points[0].Y
			d.LineTo(position.X, position.Y)
		case 'C':
			position.X = cmd.Points[2].X
			position.Y = cmd.Points[2].Y
			d.BezierTo(cmd.Points[0].X, cmd.Points[0].Y, cmd.Points[1].X, cmd.Points[1].Y, cmd.Points[2].X, cmd.Points[2].Y)
		case 'c':
			d.BezierTo(
				cmd.Points[0].X+position.X, cmd.Points[0].Y+position.Y,
				cmd.Points[1].X+position.X, cmd.Points[1].Y+position.Y,
				cmd.Points[2].X+position.X, cmd.Points[2].Y+position.Y,
			)
			position.X += cmd.Points[2].X
			position.Y += cmd.Points[2].Y
		case 'Z':
			d.CloseAndDraw()
		}
	}

	for _, child := range sn.children {
		err := child.draw(d)
		if err != nil {
			return err
		}
	}
	return nil
}
