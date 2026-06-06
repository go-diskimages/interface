package diskimage_format

import (
	"io"
	"testing"
)

// fakeFormat is a minimal implementation of Format used as a compile-time
// interface check only.
type fakeFormat struct{}

// Compile-time assertion: fakeFormat must satisfy all four interfaces.
var _ Creator = fakeFormat{}
var _ Detector = fakeFormat{}
var _ RawConverter = fakeFormat{}
var _ Resizer = fakeFormat{}
var _ Format = fakeFormat{}

func (fakeFormat) Name() string                                 { return "fake" }
func (fakeFormat) Create(path string, sizeBytes int64) error    { return nil }
func (fakeFormat) Detect(path string) (bool, error)             { return true, nil }
func (fakeFormat) ToRaw(src, dst string, w io.Writer) error     { return nil }
func (fakeFormat) Resize(path string, newSizeBytes int64) error { return nil }

// TestFormatInterface verifies that a value implementing Format is accepted
// wherever each constituent interface is required.
func TestFormatInterface(t *testing.T) {
	var f Format = fakeFormat{}

	if f.Name() != "fake" {
		t.Fatalf("Name() = %q, want %q", f.Name(), "fake")
	}
	if err := f.Create("/dev/null", 1); err != nil {
		t.Fatalf("Create: %v", err)
	}
	if ok, err := f.Detect("/dev/null"); !ok || err != nil {
		t.Fatalf("Detect: ok=%v err=%v", ok, err)
	}
	if err := f.ToRaw("/dev/null", "/dev/null", io.Discard); err != nil {
		t.Fatalf("ToRaw: %v", err)
	}
	if err := f.Resize("/dev/null", 1); err != nil {
		t.Fatalf("Resize: %v", err)
	}
}

// TestFormatInterface checks that each sub-interface can be used independently.
func TestCreatorInterface(t *testing.T) {
	var c Creator = fakeFormat{}
	if err := c.Create("/dev/null", 1); err != nil {
		t.Fatalf("Creator.Create: %v", err)
	}
}

func TestDetectorInterface(t *testing.T) {
	var d Detector = fakeFormat{}
	if ok, err := d.Detect("/dev/null"); !ok || err != nil {
		t.Fatalf("Detector.Detect: ok=%v err=%v", ok, err)
	}
}

func TestRawConverterInterface(t *testing.T) {
	var r RawConverter = fakeFormat{}
	if err := r.ToRaw("/dev/null", "/dev/null", io.Discard); err != nil {
		t.Fatalf("RawConverter.ToRaw: %v", err)
	}
}
func TestResizerInterface(t *testing.T) {
	var r Resizer = fakeFormat{}
	if err := r.Resize("/dev/null", 1); err != nil {
		t.Fatalf("Resizer.Resize: %v", err)
	}
}
