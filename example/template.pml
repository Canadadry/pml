{{define "page"}}	
Page{
	Rectangle{
		x:0.0
		y:0.0
		width:210.0
		height:297.0
		color:{{ .bgcolor }}
	}
	Rectangle{
		x:0.0
		y:0.0
		width:100.0
		height:100.0
		color:#ffaa00
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
{{end}}

Document{
{{range .page}}
    {{template "page" .}}
{{end}}
}