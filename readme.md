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

### Items

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
 - `scale`: a scaling factor 
 - `xAnchor` : how to x-position the item in its parent : `relative` (default), `left`, `right`, `center`, `fill`
 - `yAnchor` : how to y-position the item in its parent : `relative` (default), `top`, `bottom`, `center`, `fill`
 - `anchor` : shortcut to x and y anchor : `center`, `fill`
 - `color` : color of the rectangle must be an rgb hexavalue ex : `#ffaabb`
 - `borderColor` : color of the rectangle's border must be an rgb hexavalue ex : `#ffaabb`
 - `borderWidth` : width of the rectangle's border must be a float value in millimeter
 - `radius` : radius of the rectangle' corner must be a float value in millimeter

 `Image` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent 
 - `width` : width of the image must be a float value in millimeter
 - `height` : height of the image must be a float value in millimeter (if zero it will keep image aspect ratio)
 - `scale`: a scaling factor 
 - `xAnchor` : how to x-position the item in its parent : `relative` (default), `left`, `right`, `center`, `fill`
 - `yAnchor` : how to y-position the item in its parent : `relative` (default), `top`, `bottom`, `center`, `fill`
 - `anchor` : shortcut to x and y anchor : `center`, `fill`
 - `file` : full path from working dir to image file or base64 content to display
 - `mode`:  default `file` if property `file` contain a valid path or `b64` if property `file` contain a valid b64 string with image data

 `Vector` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the image must be a float value in millimeter
 - `height` : height of the image must be a float value in millimeter (if zero it will keep image aspect ratio)
 - `scale`: a scaling factor 
 - `xAnchor` : how to x-position the item in its parent : `relative` (default), `left`, `right`, `center`, `fill`
 - `yAnchor` : how to y-position the item in its parent : `relative` (default), `top`, `bottom`, `center`, `fill`
 - `anchor` : shortcut to x and y anchor : `center`, `fill`
 - `file` : full path from working dir to svg image file (partial support)

 `Text` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `scale`: a scaling factor 
 - `xAnchor` : how to x-position the item in its parent : `relative` (default), `left`, `right`, `center`, `fill`
 - `yAnchor` : how to y-position the item in its parent : `relative` (default), `top`, `bottom`, `center`, `fill`
 - `anchor` : shortcut to x and y anchor : `center`, `fill`
 - `color` : color of the rectangle must be an rgb hexavalue ex : `#ffaabb`
 - `text` : text to write
 - `align` : how the text is render in his box. Possible values are : `TopLeft`, `TopCenter`, `TopRight`, `MiddleLeft`, `MiddleCenter`, `MiddleRight`, `BottomLeft`, `BottomCenter`, `BottomRight`, `BaselineLeft`, `BaselineCenter`, `BaselineRight`
 - `fontSize` : define the size of the rendering font must be a float value in millimeter
 - `fontName` : select the font to draw text must be one of the registerer fonts see item `Font`

If you need your text to be diplayed on multiple line use a `Paragraph` instead

 `Paragraph` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `scale`: a scaling factor 
 - `xAnchor` : how to x-position the item in its parent : `relative` (default), `left`, `right`, `center`, `fill`
 - `yAnchor` : how to y-position the item in its parent : `relative` (default), `top`, `bottom`, `center`, `fill`
 - `anchor` : shortcut to x and y anchor : `center`, `fill`
 - `lineHeight` : height of the line must be a float value in millimeter
 - `align` : how the paragraph is render in his box. Possible values are : `Left`, `Right`, `Center`, `Justify`


Draw text child item  in a paragraph flow way ignoring their `x`,`y`,`width`,`heigh`, `align` and achors properties. 

 `Container` properties :

 - `x` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent 
 - `y` : left coordinate must be a float value in millimeter absolute position in the `Page` or Relative to the most close `Container` parent
 - `width` : width of the rectangle must be a float value in millimeter
 - `height` : height of the rectangle must be a float value in millimeter
 - `scale`: a scaling factor 
 - `xAnchor` : how to x-position the item in its parent : `relative` (default), `left`, `right`, `center`, `fill`
 - `yAnchor` : how to y-position the item in its parent : `relative` (default), `top`, `bottom`, `center`, `fill`
 - `anchor` : shortcut to x and y anchor : `center`, `fill`

### Using Templating 

Each pml file is pass throught the go template engine to be preprocessed.
It allow a lot usefull featuer : 

 - Injecting data comming from a json file
 - Splitting your pdf document accross several pml file
 - Translation  

When writing your pml file just follow the golang template language : [cheat sheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)

## Injecting Data

To Allow dynamic content you must provide a data file (json) to the `param` cli option or post it in api mode. 


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

### Splitting your pdf document accross several pml file

Instead of giving a file as an input, you can also gave a folder. Each pml file in the directory hierarchy will be parse by the go templating runtime
You will have to specify the main template file with the `-main` argument

_Example_

Suppose You have this directory structure

```
└── MyPmlProject/
    ├── Page/                     
    │   ├── Content.pml            
    │   ├── Header.pml
    │   └── Footer.pml
    └────── Main.pml
```

Run the following command to parse all file with the `Main.pml` file as the root element.

```bash
pml -in MyPmlProject/ -main MyPmlProject/Main.pml
```

To reference inject the other pml file just use the `template` command

```pml
Document{
    Page{
        {{ template "MyPmlProject/Page/Header.pml" . }}
        {{ template "MyPmlProject/Page/Content.pml" . }}
        {{ template "MyPmlProject/Page/Footer.pml" . }}
    }
}
```

### Translation

Translation is enable with the `-trans` argument. This argument expect a path to a csv file containing at least 2 columns. 
The first column should contain the translation keys and the following columns must contain traduction. 
Each column must start with the local language. 

_Example_

```csv
---,en,fr,es
hello,"Hello %name%","Bonjour %name%","Holà %name%"
thans,"Thank you","Merci","Gracias"
```

If the file only contain one language it will be used by default, orherwise you must use the `-local`argument to select the language you want to use. 

To access one translation value in your document use the `tr` pipeline as follow : 

```pml
Document{
    Page{
        Text{
            text:{{ tr "hello" "%name%" "World"}}
        }
    }
}
```

The `tr` pipe can takes extra arguments to remplace certain part of the text with orther value. See the pipe part for more.


### Pipes

For a simplier build of pml a few pipeline are added on top of the basic one provided by golang

- `tr`: to translate a text ( args are :  `tr key [search replace] [search replace] ...`)
- `data`: to build data map ( args are :  `data [key value] [key value] ...`)
- `array`: to build data array ( args are :  `array  value1 value2 ...`)
- `eval`: to evaluate a expression like `3+5*6` ( args are `eval expr`)
- `upper` : to capitalize a text ( args are `upper text`)


## Next

This is a good start of what we want to acheive. But we need to : 

- seperate the rendenring capabilites build on top of the pdf rendering library and the pml language parser
- create a static drawing and a flow drawing mode inspiration from [rml](https://github.com/zopefoundation/z3c.rml) 
