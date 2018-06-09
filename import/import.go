package core

import (
	"errors"
	"io"

	"github.com/bakape/hydron/common"
	"github.com/bakape/hydron/fetch"
	"github.com/bakape/thumbnailer"
)

// Common errors
var (
	ErrUnsupportedFile = errors.New("usuported file type")
	ErrImported        = errors.New("file already imported")
)

var thumbnailerOpts = thumbnailer.Options{
	JPEGQuality: 90,
	ThumbDims: thumbnailer.Dims{
		Width:  200,
		Height: 200,
	},
}

// Recursively import list of files and/or directories
func ImportPaths(
	paths []string,
	del, fetchTags bool,
	tags [][]byte,
	iLog common.Logger,
	fLog fetch.FetchLogger,
) (err error) {
	panic("TODO")
	return nil
	// files, err := files.Traverse(paths)
	// if err != nil {
	// 	return
	// }

	// // Parallelize across all cores and report progress in real time
	// type response struct {
	// 	KeyValue
	// 	path string
	// 	err  error
	// }
	// src := make(chan string)
	// res := make(chan response)
	// go func() {
	// 	for _, f := range files {
	// 		src <- f
	// 	}
	// 	close(src)
	// }()

	// n := runtime.NumCPU()
	// for i := 0; i < n; i++ {
	// 	go func() {
	// 		for path := range src {
	// 			var (
	// 				f         io.ReadCloser
	// 				err       error
	// 				fetchable = IsFetchable(path)
	// 			)
	// 			if fetchable {
	// 				f, err = FetchFile(path)
	// 			} else {
	// 				f, err = os.Open(path)
	// 			}
	// 			if err != nil {
	// 				res <- response{
	// 					path: path,
	// 					err:  err,
	// 				}
	// 				continue
	// 			}

	// 			kv, err := ImportFile(f, tags)
	// 			f.Close()
	// 			if del && !fetchable {
	// 				switch err {
	// 				case nil, ErrImported:
	// 					os.Remove(path)
	// 				}
	// 			}
	// 			res <- response{
	// 				KeyValue: kv,
	// 				path:     path,
	// 				err:      err,
	// 			}
	// 		}
	// 	}()
	// }

	// iLog.SetTotal(len(files))
	// fLog.SetTotal(len(files))
	// var (
	// 	toFetch  chan KeyValue
	// 	fecthErr chan error
	// )
	// if fetchTags {
	// 	toFetch = make(chan KeyValue, n)
	// 	fecthErr = make(chan error)
	// 	go func() {
	// 		fecthErr <- FetchTags(toFetch, fLog)
	// 	}()
	// }
	// for range files {
	// 	switch r := <-res; r.err {
	// 	case ErrUnsupportedFile, ErrImported:
	// 		iLog.Done(r.KeyValue)
	// 	case nil:
	// 		// Fetch tags of any newly imported files
	// 		if fetchTags && CanFetchTags(r.Record) {
	// 			toFetch <- r.KeyValue
	// 		}
	// 		iLog.Done(r.KeyValue)
	// 	default:
	// 		iLog.Err(fmt.Errorf("%s: %s", r.err, r.path))
	// 	}
	// }

	// if fetchTags {
	// 	close(toFetch)
	// 	return <-fecthErr
	// }

	// return nil
}

// Attempt to import any readable stream
func ImportFile(f io.Reader, tags [][]byte) (r common.Record, err error) {
	panic("TODO")
	return common.Record{}, nil
	// buf, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	return
	// }
	// hash := sha1.Sum(buf)

	// // Check, if not already in the database
	// isImported, err := isImported(hash)
	// if err != nil {
	// 	return
	// } else if isImported {
	// 	err = ErrImported
	// 	return
	// }

	// var noImage bool
	// src, thumb, err := thumbnailer.ProcessBuffer(buf, thumbnailerOpts)
	// switch err {
	// case nil:
	// case thumbnailer.ErrNoCoverArt:
	// 	noImage = true
	// case thumbnailer.ErrNoStreams:
	// 	err = ErrUnsupportedFile
	// 	return
	// default:
	// 	if _, ok := err.(thumbnailer.UnsupportedMIMEError); ok {
	// 		err = ErrUnsupportedFile
	// 	}
	// 	return
	// }

	// // Add title and artists metadata as tags, if any
	// if src.Artist != "" || src.Title != "" {
	// 	// Perform a copy, as not to modify the underlying array of the original
	// 	l := len(tags)
	// 	c := make([][]byte, l, l+2)
	// 	copy(c, tags)
	// 	tags = c

	// 	for _, s := range [...]string{src.Title, src.Artist} {
	// 		if s != "" {
	// 			tags = append(tags, normalizeTag([]byte(s)))
	// 		}
	// 	}
	// }

	// var (
	// 	typ   = MimeTypes[src.Mime]
	// 	errCh chan error
	// )
	// if !noImage {
	// 	errCh = make(chan error)
	// 	id := hex.EncodeToString(hash[:])
	// 	go func() {
	// 		errCh <- writeFile(SourcePath(id, typ), buf)
	// 	}()
	// 	go func() {
	// 		errCh <- writeFile(ThumbPath(id, thumb.IsPNG), thumb.Data)
	// 	}()
	// }

	// // Create database key-value pair
	// kv = KeyValue{
	// 	SHA1:   hash,
	// 	Record: make(Record, offsetTags),
	// }
	// kv.setType(typ)
	// kv.setStats(
	// 	uint64(time.Now().Unix()),
	// 	uint64(len(buf)),
	// 	uint64(src.Width),
	// 	uint64(src.Height),
	// 	uint64(src.Length/time.Second),
	// )
	// if thumb.IsPNG {
	// 	kv.setPNGThumb()
	// }
	// kv.setMD5(md5.Sum(buf))

	// // Receive any disk write errors
	// if !noImage {
	// 	for i := 0; i < 2; i++ {
	// 		err = <-errCh
	// 		if err != nil {
	// 			return
	// 		}
	// 	}
	// } else {
	// 	kv.setNoThumb()
	// }

	// err = writeRecord(kv, tags)
	// return
}