//go:build !darwin && arm64 && ios

package sherpa_onnx

// #cgo LDFLAGS: -L ${SRCDIR}/lib/arm64-apple-ios -lsherpa-onnx-c-api -lonnxruntime -Wl,-rpath,${SRCDIR}/lib/arm64-apple-ios
import "C"
