Document{
	Font{
		file:"example/fonts/OpenSans-Bold.ttf"
		name:"OpenSans-Bold"
	}
	Font{
		file:"example/fonts/OpenSans-Regular.ttf"
		name:"OpenSans-Regular"
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
			width:110.0
			height:110.0
			color:#ffaa00
		}
		Paragraph{
			x:0.0
			y:0.0
			width:110.0
			height:100.0
			Text{
				text:"First hello World from a Pml File"
				fontName:"OpenSans-Bold"
				fontSize: 4.0
			}
			Text{
				text:"Second hello World from a Pml File"
				fontName:"OpenSans-Regular"
				fontSize: 4.0
			}
			Text{
				text:"Third hello World from a Pml File"
				fontName:"OpenSans-Regular"
				fontSize: 5.0
			}
		}
	}
}