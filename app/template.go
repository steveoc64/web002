package main

// This package has been automatically generated with temple.
// Do not edit manually!

import (
	"github.com/go-humble/temple/temple"
)

var (
	GetTemplate     func(name string) (*temple.Template, error)
	GetPartial      func(name string) (*temple.Partial, error)
	GetLayout       func(name string) (*temple.Layout, error)
	MustGetTemplate func(name string) *temple.Template
	MustGetPartial  func(name string) *temple.Partial
	MustGetLayout   func(name string) *temple.Layout
)

func init() {
	var err error
	g := temple.NewGroup()

	if err = g.AddTemplate("admin-actions", `<div class="action-grid">
	<div class="action__item" url="/stock">
		<div class="action__title">Stock Items</div>
		<div class="action__icon"><i class="fa fa-pause-circle-o fa-lg"></i></div>
		<div class="action__text">
			Stock Items
		</div>
	</div>
</div>`); err != nil {
		panic(err)
	}

	GetTemplate = g.GetTemplate
	GetPartial = g.GetPartial
	GetLayout = g.GetLayout
	MustGetTemplate = g.MustGetTemplate
	MustGetPartial = g.MustGetPartial
	MustGetLayout = g.MustGetLayout
}
