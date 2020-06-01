Document{
	Font{
		file:"example/fonts/{{ .font}}.ttf"
		name:"{{ .font}}"
	}
	Page{
		Rectangle{
			x:0.0
			y:0.0
			width:210.0
			height:297.0
			color:#00aaff
		}
		Rectangle{
			x:0.0
			y:0.0
			width:100.0
			height:100.0
			color:#ffaa00
		}
		Text{
			text:"{{ .font}} éöî"
			x:0.0
			y:0.0
			width:100.0
			height:100.0
			color:#000000
			fontName:"{{ .font}}"
		}
	}
}