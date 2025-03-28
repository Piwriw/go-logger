package logger

import (
	"bytes"
	"flag"
	"fmt"
	"k8s.io/klog/v2"
	"testing"
)

func TestKlog(t *testing.T) {
	klog.InitFlags(nil)
	// By default klog writes to stderr. Setting logtostderr to false makes klog
	// write to a log file.
	flag.Set("logtostderr", "false")
	flag.Set("log_file", "myfile.log")
	flag.Parse()
	klog.Info("nice to meet you")
	klog.Flush()
}

func TestKlogSetOutput(t *testing.T) {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("alsologtostderr", "false")
	flag.Parse()
	buf := new(bytes.Buffer)
	klog.SetOutput(buf)
	klog.Info("nice to meet you")
	klog.Flush()

	fmt.Printf("LOGGED: %s", buf.String())
}
