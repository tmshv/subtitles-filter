# subtitles-filter

- Remove short text
- Remove `Dima Torzhok`

## Usage

1. Get binary

```sh
go build ./cmd/subfilter

```

2. Run 

Read from file:

```sh
./subfilter /path/to/webvtt-file.vtt >> clean.vtt
```

Read stdin:

```sh
./subfilter < /path/to/webvtt-file.vtt >> clean.vtt
```
```sh
cat /path/to/webvtt-file.vtt | ./subfilter >> clean.vtt
```

