# hydron
Media tagger and organizer backend and GUI frontend.
Hydron aims to be a much faster alternative to the
[Hydrus Network](https://github.com/hydrusnetwork/hydrus) and is currently in
early development.

Platforms: Linux, OSX, Win64

## Running

1. Start the `hydron` server. See `hydron -h` for more options.
2. Navigate to "http://localhost:8010" in a web browser

### Runtime dependecies
* ffmpeg >= 3.0 libraries (libswscale, libavcodec, libavutil, libavformat)
* GraphicsMagick
* Qt5 >= 5.10

## Building

`go get github.com/bakape/hydron`

### Build dependencies
* [Go](https://golang.org/doc/install) >= 1.10
* GCC or Clang
* pkg-config
* pthread
* ffmpeg >= 3.0 libraries (libswscale, libavcodec, libavutil, libavformat)
* GraphicsMagick
