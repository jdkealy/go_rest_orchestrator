package templates

var JsFieldsTemplate = `
export default {
	listFields: {{.JsFieldsConfig}},
	newFormFields: {{.JsFieldsConfig}}	
}
`
