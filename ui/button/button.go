package button

import (
	"github.com/a-h/templ"
	"github.com/xxii22w/fullstackgo/ui"
)

const (
	buttonBaseClass          = "inline-flex items-center justify-center px-4 py-2 font-medium text-sm tracking-wide transition-colors duration-200 rounded-md focus:ring focus:shadow-outline focus:outline-none"
	buttonVariantPrimary     = "text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
	buttonVariantOutline     = "text-blue-700 hover:text-white border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-blue-500 dark:text-blue-500 dark:hover:text-white dark:hover:bg-blue-500 dark:focus:ring-blue-800"
	buttonVariantSecondary   = "text-gray-900 bg-white hover:bg-gray-100 focus:outline-none"
	buttonVariantDestructive = "text-white bg-red-500 hover:bg-red-600 focus:outline-none"
)

func New(opt ...func(*templ.Attributes)) templ.Attributes {
	return ui.CreateAttrs(buttonBaseClass, buttonVariantPrimary, opt...)
}

func Outline(opts ...func(*templ.Attributes)) templ.Attributes {
	return appendVariant("outline", opts...)
}

func Primary(opts ...func(*templ.Attributes)) templ.Attributes {
	return appendVariant("primary", opts...)
}

func Secondary(opts ...func(*templ.Attributes)) templ.Attributes {
	return appendVariant("secondary", opts...)
}

func Destructive(opts ...func(*templ.Attributes)) templ.Attributes {
	return appendVariant("destructive", opts...)
}

func Variant(variant string) func(*templ.Attributes) {
	return func(attrs *templ.Attributes) {
		att := *attrs
		switch variant {
		case "primary":
			att["class"] = ui.Merge(buttonBaseClass, buttonVariantPrimary)
		case "outline":
			att["class"] = ui.Merge(buttonBaseClass, buttonVariantOutline)
		case "secondary":
			att["class"] = ui.Merge(buttonBaseClass, buttonVariantSecondary)
		case "destructive":
			att["class"] = ui.Merge(buttonBaseClass, buttonVariantDestructive)
		}
	}
}

func appendVariant(variant string, opts ...func(*templ.Attributes)) templ.Attributes {
	opt := []func(*templ.Attributes){
		Variant(variant),
	}
	opt = append(opt, opts...)
	return New(opt...)
}
