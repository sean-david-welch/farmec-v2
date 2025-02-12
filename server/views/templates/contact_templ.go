// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sean-david-welch/farmec-v2/server/views/forms"
import "github.com/sean-david-welch/farmec-v2/server/views/components"

func Contact() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<section id=\"contact\"><h1 class=\"sectionHeading\">Contact Us:</h1><div class=\"contactSection\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Map("api key").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<div class=\"infoSection\"><h1 class=\"subHeading\">Business Information:</h1><div class=\"info\"><div class=\"infoItem\">Opening Hours:<br><span class=\"infoItemText\">Monday - Friday: 9am - 5:30pm</span></div><div class=\"infoItem\">Telephone:<br><span class=\"infoItemText\"><a href=\"tel:01 825 9289\">01 825 9289</a></span></div><div class=\"infoItem\">International:<br><span class=\"infoItemText\"><a href=\"tel:+353 1 825 9289\">+353 1 825 9289</a></span></div><div class=\"infoItem\">Email:<br><span class=\"infoItemText\">Info@farmec.ie</span></div><div class=\"infoItem\">Address:<br><span class=\"infoItemText\">Clonross, Drumree, Co. Meath, A85PK30</span></div><div class=\"infoItem\"><div class=\"socialLinks\"><a class=\"socials\" href=\"https://www.facebook.com/FarmecIreland/\" target=\"_blank\" rel=\"noopener noreferrer\" aria-label=\"Visit our Facebook page\"><img src=\"/public/icons/facebook.svg\" alt=\"Facebook\" class=\"icon\"></a> <a class=\"socials\" href=\"https://twitter.com/farmec1?lang=en\" target=\"_blank\" rel=\"noopener noreferrer\" aria-label=\"Visit our Twitter page\"><img src=\"/public/icons/twitter.svg\" alt=\"Twitter\" class=\"icon\"></a></div></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
