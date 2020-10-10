// Code generated by qtc from "basepage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line templates/basepage.qtpl:3
package templates

//line templates/basepage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/basepage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/basepage.qtpl:4
type Page interface {
//line templates/basepage.qtpl:4
	Title() string
//line templates/basepage.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line templates/basepage.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line templates/basepage.qtpl:4
	Body() string
//line templates/basepage.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line templates/basepage.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line templates/basepage.qtpl:4
}

// Page prints a page implementing Page interface.

//line templates/basepage.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line templates/basepage.qtpl:12
	qw422016.N().S(`
<html>
	<head>
		<title>`)
//line templates/basepage.qtpl:15
	p.StreamTitle(qw422016)
//line templates/basepage.qtpl:15
	qw422016.N().S(`</title>
		<link href="/static/style.css" rel="stylesheet">
	</head>
	<body>
	    <div class="min-h-screen flex items-center flex-col justify-center bg-gray-100 sm:px-6 lg:px-8 rounded-md">
		    `)
//line templates/basepage.qtpl:20
	p.StreamBody(qw422016)
//line templates/basepage.qtpl:20
	qw422016.N().S(`
		</div>
	</body>
</html>
`)
//line templates/basepage.qtpl:24
}

//line templates/basepage.qtpl:24
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line templates/basepage.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/basepage.qtpl:24
	StreamPageTemplate(qw422016, p)
//line templates/basepage.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line templates/basepage.qtpl:24
}

//line templates/basepage.qtpl:24
func PageTemplate(p Page) string {
//line templates/basepage.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/basepage.qtpl:24
	WritePageTemplate(qb422016, p)
//line templates/basepage.qtpl:24
	qs422016 := string(qb422016.B)
//line templates/basepage.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/basepage.qtpl:24
	return qs422016
//line templates/basepage.qtpl:24
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line templates/basepage.qtpl:29
type BasePage struct{}

//line templates/basepage.qtpl:30
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/basepage.qtpl:30
	qw422016.N().S(`Bird Looking Glass`)
//line templates/basepage.qtpl:30
}

//line templates/basepage.qtpl:30
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/basepage.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/basepage.qtpl:30
	p.StreamTitle(qw422016)
//line templates/basepage.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line templates/basepage.qtpl:30
}

//line templates/basepage.qtpl:30
func (p *BasePage) Title() string {
//line templates/basepage.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/basepage.qtpl:30
	p.WriteTitle(qb422016)
//line templates/basepage.qtpl:30
	qs422016 := string(qb422016.B)
//line templates/basepage.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/basepage.qtpl:30
	return qs422016
//line templates/basepage.qtpl:30
}
