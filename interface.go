// Package diskimage_format defines the common interfaces for disk image
// format packages (raw, qcow2, dmg, …).
//
// Format packages do not need to import this module: Go's structural typing
// means any type whose method set is a superset of an interface automatically
// satisfies that interface. Callers that need to use formats polymorphically
// should import this module and declare the appropriate interface variable.
package diskimage_format

import "io"

// Creator can create a new blank disk image.
type Creator interface {
	// Create creates a new image at path with the given size in bytes.
	Create(path string, sizeBytes int64) error
}

// Detector probes whether a file (or directory) at path is in this format.
type Detector interface {
	// Detect returns (true, nil) when path is a valid image of this format.
	// It returns (false, nil) when the file exists but is not this format.
	// It returns (false, err) when the path cannot be examined.
	Detect(path string) (bool, error)
}

// RawConverter converts or extracts an image to a plain raw disk image.
// src may be a file path or a directory (e.g. an OCI layout cache).
// Progress messages are written to w.
type RawConverter interface {
	ToRaw(src, dst string, w io.Writer) error
}

// Resizer adjusts the virtual size of an existing disk image.
type Resizer interface {
        // Resize changes the virtual size of the image at path to newSizeBytes.
        // Implementations must support growing; shrink support is optional and
        // may return an error when newSizeBytes is smaller than the current size.
        Resize(path string, newSizeBytes int64) error
}

// Format is the full interface for a disk image format: it can create
// blank images, detect existing images, convert them to raw, and resize them.
type Format interface {
        // Name returns the format identifier (e.g. "raw", "qcow2", "dmg").
        Name() string
        Creator
        Detector
        RawConverter
        Resizer
}
