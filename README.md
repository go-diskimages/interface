<p align="center"><img src="https://raw.githubusercontent.com/go-diskimages/brand/main/social/go-diskimages.png" alt="go-diskimages/interface" width="720"></p>

# go-diskimages/interface

[![ci](https://github.com/go-diskimages/interface/actions/workflows/ci.yml/badge.svg)](https://github.com/go-diskimages/interface/actions/workflows/ci.yml)

The shared interface contract for the [`go-diskimages`](https://github.com/go-diskimages) family of pure-Go disk-image format packages (`raw`, `qcow2`, `dmg`, …). It defines what every format implementation provides — create, detect, convert-to-raw, resize — so callers can work with formats polymorphically.

Package name: `diskimage_format`.

## You usually don't need to import this

Go's structural typing means a format package satisfies these interfaces **automatically**, just by having the right method set — it does not import this module. Import `interface` only when you need to hold a format behind an interface variable (e.g. a registry, or code that accepts any format).

## The contract

```go
// Format is the full interface a disk-image format implements.
type Format interface {
    Name() string                         // "raw", "qcow2", "dmg", …
    Creator                               // Create(path, sizeBytes)
    Detector                              // Detect(path) (bool, error)
    RawConverter                          // ToRaw(src, dst, w)
    Resizer                               // Resize(path, newSizeBytes)
}
```

The four capabilities are also exposed individually (`Creator`, `Detector`, `RawConverter`, `Resizer`) so callers can depend on only what they use:

| interface | method | purpose |
| --- | --- | --- |
| `Creator` | `Create(path string, sizeBytes int64) error` | create a new blank image |
| `Detector` | `Detect(path string) (bool, error)` | probe whether `path` is this format |
| `RawConverter` | `ToRaw(src, dst string, w io.Writer) error` | convert/extract to a plain raw image (progress → `w`) |
| `Resizer` | `Resize(path string, newSizeBytes int64) error` | grow (and optionally shrink) the virtual size |

## Usage

```go
import diskimage_format "github.com/go-diskimages/interface"

func convert(f diskimage_format.Format, src, dst string) error {
    return f.ToRaw(src, dst, os.Stderr)
}
```

## License

BSD-3-Clause © the go-diskimages/interface authors.
