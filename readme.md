# Pdf Markup Language 


## Purpose

Main purpose of this language is to have a dedicated language to describe pdf content. 
Rendering use [jung-kurt/gofpdf](github.com/jung-kurt/gofpdf). Made to be easy to fork and change implementation

## Usage

```bash
go get github.com/canadadry/pml
pml -in example/template.pml -param example/template.json -out example/template.pdf
```
### CLI Options

 - `in` : input file to render
 - `out`: where to render the file
 - `param` : when using go template on top of the renderer you must specify a json parameter file (default : out.pdf)

## Example 

Here basic pml document

```pml
Document{
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

Which produce the following  ![GitHub Logo](/example/helloworld.png)

More in [example folder](/example)

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

Where `Item` and `Child` are one of the following : 

 - `Document` : Must be the root item
 - `Page` : to add page, every child will be drawn inside this one
 - `Rectangle` : to draw a rectangle
 - `Text` : to write text
 - `Image` : to draw an image
 - `Vector` : to draw a svg image
 - `Paragraph` : to draw a pragraph of text. Which feature multiple style text (changing color and font) and cariage return
 - `Container` : draw content relatively to his position

 ### Properties

 `Document` has no property

 `Font` proterties : 
 - `file` : full path from working dir to ttf file. (must have `cp12__.map` file along side to work) go [here](https://github.com/jung-kurt/gofpdf/tree/master/font) to find them 
 - `name` : registered name, will be use by `Text` item

 `Page` has no property

 `Rectangle` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `color` : color of the rectangle must be an rgb hexavalue ex : `#ffaabb`

 `Image` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent 
 - `width` : width of the image must be a float value in millimeter
 - `height` : height of the image must be a float value in millimeter (if zero it will keep image aspect ratio)
 - `file` : full path from working dir to image file or base64 content to display
 - `mode`:  default `file` if property `file` contain a valid path or `b64` if property `file` contain a valid b64 string with image data

 `Vector` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the image must be a float value in millimeter
 - `height` : height of the image must be a float value in millimeter (if zero it will keep image aspect ratio)
 - `file` : full path from working dir to svg image file (partial support)

 `Text` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `color` : color of the rectangle must be an rgb hexavalue ex : `#ffaabb`
 - `text` : text to write
 - `align` : how the text is render in his box. Possible values are : `TopLeft`, `TopCenter`, `TopRight`, `MiddleLeft`, `MiddleCenter`, `MiddleRight`, `BottomLeft`, `BottomCenter`, `BottomRight`, `BaselineLeft`, `BaselineCenter`, `BaselineRight`
 - `fontSize` : define the size of the rendering font must be a float value in millimeter
 - `fontName` : select the font to draw text must be one of the registerer fonts see item `Font`

 `Paragraph` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `lineHeight` : height of the line must be a float value in millimeter

Draw text child item  in a paragraph flow way ignoring their `x`,`y`,`width`,`heigh` and `align` properties. 

 `Container` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent 
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent

### Using Variable 

To Allow dynamic content you must provide a data file (json) to the `param` cli option or post it in api mode. 
When writing your pml file just follow the golang template language : [cheat sheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)

Example : 
```pml
Document{
    Page{
        Rectangle{
            x:0.0
            y:0.0
            width:210.0
            height:297.0
            color:{{ .bgcolor }}
        }
        Text{
            text:"{{ .title }}"
            x:{{ .x }}
            y:{{ .y }}
            width:100.0
            height:100.0
            color:#000000
            align: {{ .align }}
        }
    }
}
```

Can be customised with this json file : 

```json
{   
    "title":"titre",
    "x":"10.0",
    "y":"10.0",
    "bgcolor":"#ff0000",
    "align":"TopLeft"
}
```
