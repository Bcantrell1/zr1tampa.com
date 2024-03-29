// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package about

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AboutInfo() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section><div class=\"mx-auto max-w-screen-xl px-4 py-8 sm:px-6 sm:py-12 lg:px-8 lg:py-16\"><div class=\"grid grid-cols-1 gap-8 lg:grid-cols-2 lg:gap-16\"><div class=\"relative h-64 overflow-hidden sm:h-80 lg:order-last lg:h-full\"><img alt=\"Cody\" src=\"/static/images/content/cody_owner.jpg\" class=\"absolute inset-0 h-full w-full object-cover\"></div><div class=\"lg:py-24\"><h2 class=\"text-3xl font-bold sm:text-4xl\">Lorem ipsum dolor, sit amet consectetur adipisicing elit.\r</h2><p class=\"mt-4 text-gray-200\">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Aut qui hic atque tenetur quis\r eius quos ea neque sunt, accusantium soluta minus veniam tempora deserunt? Molestiae eius\r quidem quam repellat.\r</p><a href=\"#\" class=\"mt-8 btn btn-outline\">Get Started Today\r</a></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
