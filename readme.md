# Pdf Markup Language 


## Purpose

Main purpose of this language is to have a dedicated language to describe pdf content. 
This tool aim to convert pml to go code that build the pdf. For now it just use [jung-kurt/gofpdf](github.com/jung-kurt/gofpdf) to render a pdf

It aim to depend only a basic interface, in which you can implement the way you want. 

## Usage

```
git clone https://github.com/canadadry/pml
cd pml
go run main.go -in my-spec.pml
```

### CLI Options

`in` : input file to render
`out`: where to render the file
`param` : when using go template on top of the renderer you must specify a json parameter file
`mode` : you have several mode to play with, but most of the time the default value should be what you're looking for

 - `lexer` : only apply lexer on the pml file. No templating possible
 - `parser` : convert the file into an ast. No templating possible
 - `render` : render a file without templateing
 - `template` : no rendering done just template apply. resulting out is a pml file
 - `full` : default mode, apply template then rendre the output into a pdf file

 Why is there so many mode ? Because it is usefull when developping this tool. I will remove them when stable 

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

## Next step 

There is still important missing feature to concidere this stable

 - `Font` to allow using external font (UTF8,RTL,...)
 - `Image` to allow rendering png,jpeg,svg, ...
 - `Paragraphe` to allow styling only a part of the text
 - `Path`to draw custom form
 - more options per Item, like relative position
 - ...


