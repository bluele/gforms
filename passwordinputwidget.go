package gforms

// Generate password input field: <input type="password" ...>
func PasswordInputWidget(attrs map[string]string) Widget {
	w := new(textInputWidget)
	w.Type = "password"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
