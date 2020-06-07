package svg

import ()

func (sn *svgNode) draw(d Drawer) error {

	for _, cmd := range sn.commands {
		switch cmd.Kind {
		case 'M':
			d.MoveTo(cmd.Points[0].X, cmd.Points[0].Y)
		case 'L':
			d.LineTo(cmd.Points[0].X, cmd.Points[0].Y)
		case 'C':
			d.BezierTo(cmd.Points[0].X, cmd.Points[0].Y, cmd.Points[1].X, cmd.Points[1].Y, cmd.Points[2].X, cmd.Points[2].Y)
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
