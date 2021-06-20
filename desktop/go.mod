module github.com/jmillerv/shortcast/desktop

go 1.16

require (
	fyne.io/fyne/v2 v2.0.3
	github.com/jmillerv/shortcast/backend v0.0.0
	github.com/stretchr/testify v1.7.0
)

replace github.com/jmillerv/shortcast/backend v0.0.0 => ../backend // TODO move backend, desktop, and cli into separate repos
