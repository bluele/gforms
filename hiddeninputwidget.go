package gforms

// Generate hidden input field: <input type="hidden" ...>
func HiddenInputWidget(attrs map[string]string) Widget {
	w := new(textInputWidget)
	w.Type = "hidden"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
