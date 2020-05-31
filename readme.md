# Pdf Markup Language 


## Purpose

Main purpose of this language is to have a dedicated language to describe pdf content. 
This tool convert pml to go code that build the pdf. 

It aim to depend only a basic interface, in which you can implement the way you want. 

## Usage

```
go get ????
go run main.go -file my-spec.pml
```


## Example 

Here basic pdf document

```pml
Document{
	Page{
		Rectangle{
			x:0.0
			y:0.0
			width:100.0
			height:100.0
			color:#ffaa00
		}
		Text{
			text:"Hello World from a Pml File"
			x:0.0
			y:0.0
			width:100.0
			height:100.0
			color:#000000
		}
	}
}
```

Which produce the following  ![GitHub Logo](/example/helloworld.pdf)

more in [example folder](/example)

## Documentation 

### File Structure

Every file is written follow this syntax 


```pml

Item{
	property:value
	Child{
		...
	}
}

```

### Item

Where `Item` is one of the following : 

 - `Document` : Must be the root item
 - `Page` : to add page, every child will be drawn inside this one
 - `Rectangle` : to draw a rectangle
 - `Text` : to write text


 ### Properties

 `Document` has no property

 `Page` has no property

 `Rectangle` properties :

 - `x` : left coordinate must be a float value in millimeter
 - `y` : left coordinate must be a float value in millimeter
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `color` : color of the rectangle must be an rgb hexavalue ex : `#ffaabb`


 `Text` properties :

 - `x` : left coordinate must be a float value in millimeter
 - `y` : left coordinate must be a float value in millimeter
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `color` : color of the rectangle must be an rgb hexavalue ex : `#ffaabb`
 - `text` : text to write





