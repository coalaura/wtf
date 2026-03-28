package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/coalaura/plain"
	"github.com/coalaura/wtf/types"
)

var (
	Version = "dev"
	log     = plain.New()
)

const MaxReadSize = 128 * 1024

func main() {
	var (
		targetFile string
		porcelain  bool
		timing     bool
	)

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help":
			fmt.Println("wtf - hardware-accelerated file sniffer")
			fmt.Println("\nUsage: wtf [flags] <file>")
			fmt.Println("\nFlags:")
			fmt.Println("  -p, --porcelain  Print easily parseable output (tab-separated: Kind\\tType)")
			fmt.Println("  -t, --time       Print time taken (read / sniff; disabled by -p)")
			fmt.Println("  -v, --version    Print version information")
			fmt.Println("  -h, --help       Print this help message")

			os.Exit(0)
		case "-v", "--version":
			fmt.Printf("wtf %s\n", Version)

			os.Exit(0)
		case "-p", "--porcelain":
			porcelain = true
		case "-t", "--time":
			timing = true
		default:
			if targetFile == "" && (len(arg) == 0 || arg[0] != '-') {
				targetFile = arg
			} else {
				log.Errorln("Unknown argument or multiple files specified: ", arg)

				os.Exit(1)
			}
		}
	}

	if targetFile == "" {
		log.Errorln("Usage: wtf [flags] <file>")

		os.Exit(1)
	}

	path, err := filepath.Abs(targetFile)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	name := filepath.Base(path)

	t0 := time.Now()

	info, err := os.Lstat(path)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	meta := detectPath(name, info)
	if meta != nil {
		d0 := time.Since(t0)

		printMeta(meta, formatTimings(d0, 0, timing), porcelain)

		return
	}

	readSize := min(int(info.Size()), MaxReadSize)

	if readSize == 0 {
		log.Errorln("empty file")

		os.Exit(2)
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Errorln(err)

		os.Exit(1)
	}

	defer file.Close()

	buf := make([]byte, readSize)

	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		log.Errorln(err)

		os.Exit(1)
	}

	d0 := time.Since(t0)
	t1 := time.Now()

	meta, err = types.Detect(name, buf[:n])
	if err != nil {
		if porcelain {
			fmt.Println("Unknown\t")
		} else {
			log.Errorln(err)
		}

		os.Exit(1)
	}

	d1 := time.Since(t1)

	printMeta(meta, formatTimings(d0, d1, timing), porcelain)
}

func detectPath(name string, info os.FileInfo) *types.Metadata {
	mode := info.Mode()

	if mode&os.ModeSymlink != 0 {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeSymbolicLink,
		}
	}

	if mode.IsDir() {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeDirectory,
		}
	}

	if mode&os.ModeNamedPipe != 0 {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeNamedPipe,
		}
	}

	if mode&os.ModeSocket != 0 {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeSocket,
		}
	}

	if mode&os.ModeDevice != 0 {
		if mode&os.ModeCharDevice != 0 {
			return &types.Metadata{
				File: name,
				Kind: types.KindFilesystemEntry,
				Type: types.TypeCharacterDevice,
			}
		}

		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeBlockDevice,
		}
	}

	if !mode.IsRegular() {
		return &types.Metadata{
			File: name,
			Kind: types.KindFilesystemEntry,
			Type: types.TypeSpecial,
		}
	}

	return nil
}

func printMeta(meta *types.Metadata, timings string, porcelain bool) {
	if porcelain {
		fmt.Printf("%s\t%s\t%s\n", meta.Kind.String(), meta.Type.String(), meta.Confidence.String())
	} else {
		log.Println(meta.Format(timings))
	}
}

func formatTimings(d0, d1 time.Duration, enabled bool) string {
	if !enabled {
		return ""
	}

	return fmt.Sprintf("%s / %s", formatTime(d0), formatTime(d1))
}

func formatTime(d time.Duration) string {
	if d < 0 {
		d = 0
	}

	if d == 0 {
		return "0s"
	}

	type unit struct {
		value  time.Duration
		suffix string
	}

	units := []unit{
		{time.Hour, "h"},
		{time.Minute, "m"},
		{time.Second, "s"},
		{time.Millisecond, "ms"},
		{time.Microsecond, "µs"},
	}

	parts := make([]string, 0, len(units))

	for _, unit := range units {
		if d < unit.value {
			continue
		}

		count := d / unit.value
		d -= count * unit.value

		parts = append(parts, fmt.Sprintf("%d%s", count, unit.suffix))
	}

	if len(parts) == 0 {
		return fmt.Sprintf("%dns", d/time.Nanosecond)
	}

	return strings.Join(parts, " ")
}
