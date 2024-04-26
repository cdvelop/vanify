package page

func (p *Page) SetDefault() error {
	p.StyleSheet = "static/style.css"
	p.AppName = "appTest"
	p.appVersion = "v0.0.0"
	p.UserName = ""
	p.UserArea = ""
	p.Message = ""
	p.Script = "static/main.js"

	return nil

}
