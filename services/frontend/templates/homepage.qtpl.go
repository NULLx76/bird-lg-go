// Code generated by qtc from "homepage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/homepage.qtpl:2
package templates

//line templates/homepage.qtpl:2
import (
	"github.com/NULLx76/bird-lg-go/pkg/proxy"
)

//line templates/homepage.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/homepage.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/homepage.qtpl:8
type MainPage struct {
	Summaries map[string]proxy.SummaryTable
}

//line templates/homepage.qtpl:13
func (p *MainPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/homepage.qtpl:13
	qw422016.N().S(`
	Bird Looking Glass
`)
//line templates/homepage.qtpl:15
}

//line templates/homepage.qtpl:15
func (p *MainPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/homepage.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/homepage.qtpl:15
	p.StreamTitle(qw422016)
//line templates/homepage.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line templates/homepage.qtpl:15
}

//line templates/homepage.qtpl:15
func (p *MainPage) Title() string {
//line templates/homepage.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/homepage.qtpl:15
	p.WriteTitle(qb422016)
//line templates/homepage.qtpl:15
	qs422016 := string(qb422016.B)
//line templates/homepage.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/homepage.qtpl:15
	return qs422016
//line templates/homepage.qtpl:15
}

//line templates/homepage.qtpl:17
func (p *MainPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/homepage.qtpl:17
	qw422016.N().S(`
	<h1>dn42 Looking Glass</h1>
    `)
//line templates/homepage.qtpl:19
	for name, server := range p.Summaries {
//line templates/homepage.qtpl:19
		qw422016.N().S(`
    <div class="text-center">
        <h2 class="p-3">Server: `)
//line templates/homepage.qtpl:21
		qw422016.E().S(name)
//line templates/homepage.qtpl:21
		qw422016.N().S(`</h2>
        <table>
            <tr>
                <th>Name</th>
                <th>Proto</th>
                <th>Table</th>
                <th>State</th>
                <th>Since</th>
                <th>Info</th>
            </tr>
            `)
//line templates/homepage.qtpl:31
		for row := range server {
//line templates/homepage.qtpl:31
			qw422016.N().S(`
            <tr>
                <td><a href="/`)
//line templates/homepage.qtpl:33
			qw422016.E().S(name)
//line templates/homepage.qtpl:33
			qw422016.N().S(`/details/`)
//line templates/homepage.qtpl:33
			qw422016.E().S(server[row].Name)
//line templates/homepage.qtpl:33
			qw422016.N().S(`"> `)
//line templates/homepage.qtpl:33
			qw422016.E().S(server[row].Name)
//line templates/homepage.qtpl:33
			qw422016.N().S(` </a></td>
                <td>`)
//line templates/homepage.qtpl:34
			qw422016.E().S(server[row].Proto)
//line templates/homepage.qtpl:34
			qw422016.N().S(`</td>
                <td>`)
//line templates/homepage.qtpl:35
			qw422016.E().S(server[row].Table)
//line templates/homepage.qtpl:35
			qw422016.N().S(`</td>
                <td>`)
//line templates/homepage.qtpl:36
			qw422016.E().S(server[row].State)
//line templates/homepage.qtpl:36
			qw422016.N().S(`</td>
                <td>`)
//line templates/homepage.qtpl:37
			qw422016.E().S(server[row].Since)
//line templates/homepage.qtpl:37
			qw422016.N().S(`</td>
                <td>`)
//line templates/homepage.qtpl:38
			qw422016.E().S(server[row].Info)
//line templates/homepage.qtpl:38
			qw422016.N().S(`</td>
            </tr>
            `)
//line templates/homepage.qtpl:40
		}
//line templates/homepage.qtpl:40
		qw422016.N().S(`
        </table>
    </div>
    `)
//line templates/homepage.qtpl:43
	}
//line templates/homepage.qtpl:43
	qw422016.N().S(`
`)
//line templates/homepage.qtpl:44
}

//line templates/homepage.qtpl:44
func (p *MainPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/homepage.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/homepage.qtpl:44
	p.StreamBody(qw422016)
//line templates/homepage.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line templates/homepage.qtpl:44
}

//line templates/homepage.qtpl:44
func (p *MainPage) Body() string {
//line templates/homepage.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/homepage.qtpl:44
	p.WriteBody(qb422016)
//line templates/homepage.qtpl:44
	qs422016 := string(qb422016.B)
//line templates/homepage.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/homepage.qtpl:44
	return qs422016
//line templates/homepage.qtpl:44
}
