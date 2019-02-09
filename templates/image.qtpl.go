// This file is automatically generated by qtc from "image.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line image.qtpl:1
package templates

//line image.qtpl:1
import "github.com/bakape/hydron/common"

//line image.qtpl:2
import "github.com/bakape/hydron/files"

//line image.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line image.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line image.qtpl:4
func StreamThumbnail(qw422016 *qt422016.Writer, img common.CompactImage, page common.Page, highlight bool) {
	//line image.qtpl:4
	qw422016.N().S(`<figure data-href="`)
	//line image.qtpl:5
	qw422016.N().S(files.NetSourcePath(img.SHA1, img.Type))
	//line image.qtpl:5
	qw422016.N().S(`"`)
	//line image.qtpl:5
	if highlight {
		//line image.qtpl:5
		qw422016.N().S(` `)
		//line image.qtpl:5
		qw422016.N().S(`class="highlight"`)
		//line image.qtpl:5
	}
	//line image.qtpl:5
	qw422016.N().S(`><input type="checkbox" name="img:`)
	//line image.qtpl:6
	qw422016.N().S(img.SHA1)
	//line image.qtpl:6
	qw422016.N().S(`"><div class="background"></div><a href="/image/`)
	//line image.qtpl:8
	qw422016.N().S(img.SHA1)
	//line image.qtpl:8
	qw422016.N().S(`?`)
	//line image.qtpl:8
	qw422016.N().S(page.Query())
	//line image.qtpl:8
	qw422016.N().S(`"><img width="`)
	//line image.qtpl:9
	qw422016.N().D(int(img.Thumb.Width))
	//line image.qtpl:9
	qw422016.N().S(`" height="`)
	//line image.qtpl:9
	qw422016.N().D(int(img.Thumb.Height))
	//line image.qtpl:9
	qw422016.N().S(`" src="`)
	//line image.qtpl:9
	qw422016.N().S(files.NetThumbPath(img.SHA1))
	//line image.qtpl:9
	qw422016.N().S(`"></a></figure>`)
//line image.qtpl:12
}

//line image.qtpl:12
func WriteThumbnail(qq422016 qtio422016.Writer, img common.CompactImage, page common.Page, highlight bool) {
	//line image.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line image.qtpl:12
	StreamThumbnail(qw422016, img, page, highlight)
	//line image.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line image.qtpl:12
}

//line image.qtpl:12
func Thumbnail(img common.CompactImage, page common.Page, highlight bool) string {
	//line image.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	//line image.qtpl:12
	WriteThumbnail(qb422016, img, page, highlight)
	//line image.qtpl:12
	qs422016 := string(qb422016.B)
	//line image.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	//line image.qtpl:12
	return qs422016
//line image.qtpl:12
}

//line image.qtpl:14
func StreamImagePage(qw422016 *qt422016.Writer, img common.Image, page common.Page) {
	//line image.qtpl:15
	title := img.Name

	//line image.qtpl:16
	if title == "" {
		//line image.qtpl:17
		title = "hydron"

		//line image.qtpl:18
	}
	//line image.qtpl:19
	streamhead(qw422016, title)
	//line image.qtpl:19
	qw422016.N().S(`<body><div id="image-view"><section id="tags">`)
	//line image.qtpl:23
	if img.Name != "" {
		//line image.qtpl:23
		qw422016.N().S(`<span class="truncate image-name"><a href="/search?q=`)
		//line image.qtpl:25
		qw422016.E().S(img.Name)
		//line image.qtpl:25
		qw422016.N().S(`" title="Search for name">Name:`)
		//line image.qtpl:26
		qw422016.N().S(` `)
		//line image.qtpl:26
		qw422016.E().S(img.Name)
		//line image.qtpl:26
		qw422016.N().S(`</a></span>`)
		//line image.qtpl:29
	}
	//line image.qtpl:30
	org := organizeTags(img.Tags)

	//line image.qtpl:31
	streamrenderTags(qw422016, org[common.Character], page)
	//line image.qtpl:32
	streamrenderTags(qw422016, org[common.Series], page)
	//line image.qtpl:33
	streamrenderTags(qw422016, org[common.Author], page)
	//line image.qtpl:34
	streamrenderTags(qw422016, org[common.Rating], page)
	//line image.qtpl:35
	streamrenderTags(qw422016, org[common.Meta], page)
	//line image.qtpl:36
	streamrenderTags(qw422016, org[common.Undefined], page)
	//line image.qtpl:36
	qw422016.N().S(`</section>`)
	//line image.qtpl:38
	src := files.NetSourcePath(img.SHA1, img.Type)

	//line image.qtpl:39
	switch common.GetMediaType(img.Type) {
	//line image.qtpl:40
	case common.MediaImage:
		//line image.qtpl:40
		qw422016.N().S(`<img src="`)
		//line image.qtpl:41
		qw422016.N().S(src)
		//line image.qtpl:41
		qw422016.N().S(`">`)
	//line image.qtpl:42
	case common.MediaVideo:
		//line image.qtpl:42
		qw422016.N().S(`<video src="`)
		//line image.qtpl:43
		qw422016.N().S(src)
		//line image.qtpl:43
		qw422016.N().S(`" autoplay loop controls>`)
	//line image.qtpl:44
	default:
		//line image.qtpl:44
		qw422016.N().S(`<b>Display not supported for this file format</b>`)
		//line image.qtpl:46
	}
	//line image.qtpl:46
	qw422016.N().S(`</div></body>`)
//line image.qtpl:49
}

//line image.qtpl:49
func WriteImagePage(qq422016 qtio422016.Writer, img common.Image, page common.Page) {
	//line image.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line image.qtpl:49
	StreamImagePage(qw422016, img, page)
	//line image.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line image.qtpl:49
}

//line image.qtpl:49
func ImagePage(img common.Image, page common.Page) string {
	//line image.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
	//line image.qtpl:49
	WriteImagePage(qb422016, img, page)
	//line image.qtpl:49
	qs422016 := string(qb422016.B)
	//line image.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
	//line image.qtpl:49
	return qs422016
//line image.qtpl:49
}

// Render tag adition and direct tag query links

//line image.qtpl:52
func streamrenderTags(qw422016 *qt422016.Writer, tags []common.Tag, page common.Page) {
	//line image.qtpl:53
	page.Page = 0

	//line image.qtpl:54
	init := page.Filters

	//line image.qtpl:55
	for _, t := range tags {
		//line image.qtpl:56
		page.Filters = init

		//line image.qtpl:57
		filter := common.TagFilter{TagBase: t.TagBase}

		//line image.qtpl:58
		page.Filters.Tag = append(page.Filters.Tag, filter)

		//line image.qtpl:58
		qw422016.N().S(`<span class="spaced truncate tag-`)
		//line image.qtpl:59
		qw422016.N().Z(common.BufferWriter(t.Type))
		//line image.qtpl:59
		qw422016.N().S(`"><a href="`)
		//line image.qtpl:60
		qw422016.N().S(page.URL())
		//line image.qtpl:60
		qw422016.N().S(`" class="char-button" title="Add to search">+</a>`)
		//line image.qtpl:63
		page.Filters.Tag[len(page.Filters.Tag)-1].Negative = true

		//line image.qtpl:63
		qw422016.N().S(`<a href="`)
		//line image.qtpl:64
		qw422016.N().S(page.URL())
		//line image.qtpl:64
		qw422016.N().S(`" class="char-button" title="Remove from search">-</a>`)
		//line image.qtpl:67
		page.Filters = common.FilterSet{
			Tag: []common.TagFilter{filter},
		}

		//line image.qtpl:69
		qw422016.N().S(`<a href="`)
		//line image.qtpl:70
		qw422016.N().S(page.URL())
		//line image.qtpl:70
		qw422016.N().S(`" title="Search for`)
		//line image.qtpl:70
		qw422016.N().S(` `)
		//line image.qtpl:70
		qw422016.E().S(t.Tag)
		//line image.qtpl:70
		qw422016.N().S(`">`)
		//line image.qtpl:71
		if t.Type == common.Rating {
			//line image.qtpl:71
			qw422016.N().S(`rating:`)
			//line image.qtpl:72
			qw422016.N().S(` `)
			//line image.qtpl:73
		}
		//line image.qtpl:74
		qw422016.E().S(t.Tag)
		//line image.qtpl:74
		qw422016.N().S(`</a></span>`)
		//line image.qtpl:77
	}
//line image.qtpl:78
}

//line image.qtpl:78
func writerenderTags(qq422016 qtio422016.Writer, tags []common.Tag, page common.Page) {
	//line image.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line image.qtpl:78
	streamrenderTags(qw422016, tags, page)
	//line image.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line image.qtpl:78
}

//line image.qtpl:78
func renderTags(tags []common.Tag, page common.Page) string {
	//line image.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
	//line image.qtpl:78
	writerenderTags(qb422016, tags, page)
	//line image.qtpl:78
	qs422016 := string(qb422016.B)
	//line image.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
	//line image.qtpl:78
	return qs422016
//line image.qtpl:78
}
