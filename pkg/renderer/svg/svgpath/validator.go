package svgpath

func validateCommand(cmd Command) bool {

	switch cmd.Kind {
	case 'z':
		return len(cmd.Points) == 0
	case 'Z':
		return len(cmd.Points) == 0
	case 'm':
		return len(cmd.Points) == 1
	case 'M':
		return len(cmd.Points) == 1
	case 'l':
		return len(cmd.Points) == 1
	case 'L':
		return len(cmd.Points) == 1
	case 'c':
		return len(cmd.Points)%3 == 0 && len(cmd.Points) > 0
	case 'C':
		return len(cmd.Points)%3 == 0 && len(cmd.Points) > 0
	}
	return false
}
