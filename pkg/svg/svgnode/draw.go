package svgnode

import (
	"github.com/canadadry/pml/pkg/svg/svgdrawer"
	"github.com/canadadry/pml/pkg/svg/svgpath"
)

func (sn *SvgNode) Draw(d svgdrawer.Drawer) error {

	position := svgpath.Point{}

	cmdCount := 0

	for _, cmd := range sn.commands {
		cmdCount = cmdCount + 1
		switch cmd.Kind {
		case 'M':
			position.X = cmd.Points[0].X
			position.Y = cmd.Points[0].Y
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.MoveTo(x, y)
		case 'm':
			position.X += cmd.Points[0].X
			position.Y += cmd.Points[0].Y
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.MoveTo(x, y)
		case 'L':
			position.X = cmd.Points[0].X
			position.Y = cmd.Points[0].Y
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.LineTo(x, y)
		case 'l':
			position.X += cmd.Points[0].X
			position.Y += cmd.Points[0].Y
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.LineTo(x, y)
		case 'H':
			position.X = cmd.Points[0].X
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.LineTo(x, y)
		case 'h':
			position.X += cmd.Points[0].X
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.LineTo(x, y)
		case 'V':
			position.Y = cmd.Points[0].X
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.LineTo(x, y)
		case 'v':
			position.Y += cmd.Points[0].X
			x, y := sn.worldToLocal.ProjectPoint(position.X, position.Y)
			d.LineTo(x, y)
		case 'Q':
			for i := 0; i < len(cmd.Points); i += 2 {
				cx, cy := sn.worldToLocal.ProjectPoint(cmd.Points[i+0].X, cmd.Points[i+0].Y)
				x, y := sn.worldToLocal.ProjectPoint(cmd.Points[i+1].X, cmd.Points[i+1].Y)
				d.BezierTo(cx, cy, cx, cy, x, y)
				position.X = cmd.Points[i+1].X
				position.Y = cmd.Points[i+1].Y
			}
		case 'q':
			for i := 0; i < len(cmd.Points); i += 2 {
				cx, cy := sn.worldToLocal.ProjectPoint(cmd.Points[i+0].X+position.X, cmd.Points[i+0].Y+position.Y)
				x, y := sn.worldToLocal.ProjectPoint(cmd.Points[i+1].X+position.X, cmd.Points[i+1].Y+position.Y)
				d.BezierTo(cx, cy, cx, cy, x, y)
				position.X += cmd.Points[i+1].X
				position.Y += cmd.Points[i+1].Y
			}
		case 'C':
			for i := 0; i < len(cmd.Points); i += 3 {
				c1x, c1y := sn.worldToLocal.ProjectPoint(cmd.Points[i+0].X, cmd.Points[i+0].Y)
				c2x, c2y := sn.worldToLocal.ProjectPoint(cmd.Points[i+1].X, cmd.Points[i+1].Y)
				x, y := sn.worldToLocal.ProjectPoint(cmd.Points[i+2].X, cmd.Points[i+2].Y)
				d.BezierTo(c1x, c1y, c2x, c2y, x, y)
				position.X = cmd.Points[i+2].X
				position.Y = cmd.Points[i+2].Y
			}
		case 'c':
			for i := 0; i < len(cmd.Points); i += 3 {
				c1x, c1y := sn.worldToLocal.ProjectPoint(cmd.Points[i+0].X+position.X, cmd.Points[i+0].Y+position.Y)
				c2x, c2y := sn.worldToLocal.ProjectPoint(cmd.Points[i+1].X+position.X, cmd.Points[i+1].Y+position.Y)
				x, y := sn.worldToLocal.ProjectPoint(cmd.Points[i+2].X+position.X, cmd.Points[i+2].Y+position.Y)
				d.BezierTo(c1x, c1y, c2x, c2y, x, y)
				position.X += cmd.Points[i+2].X
				position.Y += cmd.Points[i+2].Y
			}
		case 'Z':
			d.CloseAndDraw(sn.style)
			cmdCount = 0
		case 'z':
			d.CloseAndDraw(sn.style)
			cmdCount = 0
		}
	}

	if cmdCount > 0 {
		d.CloseAndDraw(sn.style)
	}

	for _, child := range sn.children {
		err := child.Draw(d)
		if err != nil {
			return err
		}
	}
	return nil
}
