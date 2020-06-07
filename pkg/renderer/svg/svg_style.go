package svg

import (
	"image/color"
	"pml/pkg/renderer/svg/svgparser"
)

func parseStyleAttribute(element *svgparser.Element) Style {
	s := Style{
		Fill:        true,
		FillColor:   color.RGBA{255, 0, 0, 0},
		BorderSize:  0.1,
		BorderColor: color.RGBA{0, 255, 0, 0},
	}

	return s
}

// parseStyleFromAttributes(svg:string)
// {
// 	let fill:boolean = false;
// 	let stroke:boolean = false;

// 	let styles = svg.split(';');
// 	for(let style of styles)
// 	{
// 		let arg = style.split(':');
// 		switch (arg[0])
// 		{
// 			case "fill":
// 			{
// 				fill =  true;
// 				if(arg[1].substring(0,4) == 'rgb(')
// 				{
// 					this.style.fillColor = arg[1].substring(4,arg[1].length-1).split(',').map(c=>parseFloat(c));
// 				}
// 				else
// 				{
// 					let dict = this.getColorDictionnary();
// 					if(dict.hasOwnProperty(arg[1]))
// 					{
// 						this.style.fillColor = dict[arg[1]];
// 					}
// 				}
// 				break;
// 			}
// 			case "stroke":
// 			{
// 				stroke =  true;
// 				if(arg[1].substring(0,4) == 'rgb(')
// 				{
// 					this.style.stokeColor = arg[1].substring(4,arg[1].length-1).split(',').map(c=>parseFloat(c));
// 				}
// 				else
// 				{
// 					let dict = this.getColorDictionnary();
// 					if(dict.hasOwnProperty(arg[1]))
// 					{
// 						this.style.fillColor = dict[arg[1]];
// 					}
// 				}
// 				break;
// 			}
// 			case "stroke-width":
// 			{
// 				stroke =  true;
// 				if(arg[1].substring(arg[1].length-2,arg[1].length) == 'px')
// 				{
// 					let width = parseFloat(arg[1].substring(0,arg[1].length-2));
// 					let newWidth = this.worldToLocal.multiplyPoint(width,0,0);
// 					this.style.strokeWidth = newWidth.x;
// 				}
// 				break;
// 			}

// 		}
// 	}

// 	if( fill && stroke)
// 	{
// 		this.style.mode = 'FD';
// 	}
// 	else if( fill )
// 	{
// 		this.style.mode = 'F';
// 	}
// 	else
// 	{
// 		this.style.mode = 'S';
// 	}
// }
