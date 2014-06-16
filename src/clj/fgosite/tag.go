// Hiccup helpers

package tag

func Js(src) { [SCRIPT, {SRC: src}] }

func Css(href) { [LINK, {REL: "stylesheet", HREF: href}] }
