package logger

import (
	"bytes"
	"flag"
	"fmt"
	"testing"

	"k8s.io/klog/v2"
)

func TestKlog(t *testing.T) {
	klog.InitFlags(nil)
	// By default klog writes to stderr. Setting logtostderr to false makes klog
	// write to a log file.
	if err := flag.Set("logtostderr", "false"); err != nil {
		t.Fatal(err)
	}
	if err := flag.Set("log_file", "myfile.log"); err != nil {
		t.Fatal(err)
	}
	flag.Parse()
	klog.Info("nice to meet you")
	klog.Flush()
}

func TestKlogSetOutput(t *testing.T) {
	klog.InitFlags(nil)
	if err := flag.Set("logtostderr", "false"); err != nil {
		t.Fatal(err)
	}
	if err := flag.Set("alsologtostderr", "false"); err != nil {
		t.Fatal(err)
	}
	flag.Parse()
	buf := new(bytes.Buffer)
	klog.SetOutput(buf)
	klog.Info("nice to meet you")
	klog.Flush()

	fmt.Printf("LOGGED: %s", buf.String())
}
