package main

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module = "github.com/diamondburned/gotk4/pkg"
	thisModule  = "github.com/nothub/goflatpak/pkg"
)

var packages = []gendata.Package{
	{PkgName: "gdk-pixbuf-2.0", Namespaces: []string{"Flatpak-1"}},
}

var pkgGenerated = []string{
	"flatpak",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
}

var preprocessors []types.Preprocessor

var filters []types.FilterMatcher
