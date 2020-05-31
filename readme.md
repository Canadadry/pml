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
	Font{
		id:title
		familly:"helevetica"
		size:12
		weight:3
	}
	Page{
		id:page
		Row{
			Rectangle{
				height: 20
				color: #ffeeaa
				Text{
					text:"Hello world"
					font:title
				}
			}	

		}
	}
}
```

