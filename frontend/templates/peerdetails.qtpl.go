// Code generated by qtc from "peerdetails.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/peerdetails.qtpl:2
package templates

//line templates/peerdetails.qtpl:2
import (
	"github.com/NULLx76/bird-lg-go/api/proxy"
	"github.com/valyala/fasthttp"
)

//line templates/peerdetails.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/peerdetails.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/peerdetails.qtpl:9
type PeerPage struct {
	CTX  *fasthttp.RequestCtx
	Peer *proxy.PeerDetails
}

//line templates/peerdetails.qtpl:15
func (p *PeerPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/peerdetails.qtpl:15
	qw422016.N().S(`
	Bird Looking Glass - `)
//line templates/peerdetails.qtpl:16
	qw422016.E().S(p.Peer.Info.Name)
//line templates/peerdetails.qtpl:16
	qw422016.N().S(`
`)
//line templates/peerdetails.qtpl:17
}

//line templates/peerdetails.qtpl:17
func (p *PeerPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/peerdetails.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/peerdetails.qtpl:17
	p.StreamTitle(qw422016)
//line templates/peerdetails.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line templates/peerdetails.qtpl:17
}

//line templates/peerdetails.qtpl:17
func (p *PeerPage) Title() string {
//line templates/peerdetails.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/peerdetails.qtpl:17
	p.WriteTitle(qb422016)
//line templates/peerdetails.qtpl:17
	qs422016 := string(qb422016.B)
//line templates/peerdetails.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/peerdetails.qtpl:17
	return qs422016
//line templates/peerdetails.qtpl:17
}

//line templates/peerdetails.qtpl:19
func (p *PeerPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/peerdetails.qtpl:19
	qw422016.N().S(`
	<h2 class="p-3">Peer Details</h2>
	<table>
        <tr>
            <th>Name</th>
            <th>Proto</th>
            <th>Table</th>
            <th>State</th>
            <th>Since</th>
            <th>Info</th>
        </tr>
        <tr>
            <td>`)
//line templates/peerdetails.qtpl:31
	qw422016.E().S(p.Peer.Info.Name)
//line templates/peerdetails.qtpl:31
	qw422016.N().S(`</td>
            <td>`)
//line templates/peerdetails.qtpl:32
	qw422016.E().S(p.Peer.Info.Proto)
//line templates/peerdetails.qtpl:32
	qw422016.N().S(`</td>
            <td>`)
//line templates/peerdetails.qtpl:33
	qw422016.E().S(p.Peer.Info.Table)
//line templates/peerdetails.qtpl:33
	qw422016.N().S(`</td>
            <td>`)
//line templates/peerdetails.qtpl:34
	qw422016.E().S(p.Peer.Info.State)
//line templates/peerdetails.qtpl:34
	qw422016.N().S(`</td>
            <td>`)
//line templates/peerdetails.qtpl:35
	qw422016.E().S(p.Peer.Info.Since)
//line templates/peerdetails.qtpl:35
	qw422016.N().S(`</td>
            <td>`)
//line templates/peerdetails.qtpl:36
	qw422016.E().S(p.Peer.Info.Info)
//line templates/peerdetails.qtpl:36
	qw422016.N().S(`</td>
        </tr>
	</table>
	`)
//line templates/peerdetails.qtpl:39
	if p.Peer.Details != "" {
//line templates/peerdetails.qtpl:39
		qw422016.N().S(`
	<pre>
`)
//line templates/peerdetails.qtpl:41
		qw422016.E().S(p.Peer.Details)
//line templates/peerdetails.qtpl:41
		qw422016.N().S(`
	</pre>
	`)
//line templates/peerdetails.qtpl:43
	}
//line templates/peerdetails.qtpl:43
	qw422016.N().S(`
`)
//line templates/peerdetails.qtpl:44
}

//line templates/peerdetails.qtpl:44
func (p *PeerPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/peerdetails.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/peerdetails.qtpl:44
	p.StreamBody(qw422016)
//line templates/peerdetails.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line templates/peerdetails.qtpl:44
}

//line templates/peerdetails.qtpl:44
func (p *PeerPage) Body() string {
//line templates/peerdetails.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/peerdetails.qtpl:44
	p.WriteBody(qb422016)
//line templates/peerdetails.qtpl:44
	qs422016 := string(qb422016.B)
//line templates/peerdetails.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/peerdetails.qtpl:44
	return qs422016
//line templates/peerdetails.qtpl:44
}
