Document{
	Page{
		Rectangle{
			width:210.0
			height:297.0
			color:#00aaff
		}
		Container{
			{{template "block"}}
		}

		Container{
			y:30
			x:10
			{{template "block"}}
		}
	}
}


{{define "block"}}
Text{
	text:"Hello World from a Pml File"
	width:25.0
	height:25.0
	color:#000000
}
Image{
	x:30.0
	width:25.0
	file:"example/image/github_logo.png"
}
{{end}}