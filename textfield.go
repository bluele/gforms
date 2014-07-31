package gforms

type TextField struct {
	BaseField
}

func (self *TextField) html(vs ...string) string {
	return renderTemplate("TextTypeField", newTemplateContext(self, vs...))
}

func NewTextField(name string, vs Validators, ws ...Widget) *TextField {
	self := new(TextField)
	self.name = name
	self.validators = vs
	if len(ws) > 0 {
		self.Widget = ws[0]
	}
	return self
}
