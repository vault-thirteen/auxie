package bom

import (
	"testing"

	tsb "github.com/vault-thirteen/auxie/TSB"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetEncodingsReport(t *testing.T) {
	tst := tester.New(t)
	var data []byte
	var encodingsToProbe map[Encoding]bool
	var report *Report
	var err error

	// Test #1.
	data = []byte{0xFE, 0xFF, 'C'}
	encodingsToProbe = map[Encoding]bool{
		EncodingUTF16BE: true,
		EncodingUTF8:    true,
	}
	report, err = GetEncodingsReport(data, encodingsToProbe)
	tst.MustBeNoError(err)
	tst.MustBeEqual(report, &Report{
		probes: map[Encoding]*Probe{
			EncodingUTF16BE: {
				Encoding:       EncodingUTF16BE,
				Probability:    tsb.Yes,
				ReadBytesCount: 2,
			},
			EncodingUTF8: {
				Encoding:       EncodingUTF8,
				Probability:    tsb.No,
				ReadBytesCount: 3,
			},
		},
	})

	// Test #2.
	data = []byte{'A'}
	encodingsToProbe = map[Encoding]bool{
		Encoding(99): true,
	}
	report, err = GetEncodingsReport(data, encodingsToProbe)
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), "unknown encoding: 99")
	tst.MustBeEqual(report, (*Report)(nil))
}

func Test_IsAccurate_Report(t *testing.T) {
	tst := tester.New(t)
	var report *Report

	prA := &Probe{Encoding: Encoding(33), Probability: tsb.Yes}
	prB := &Probe{Encoding: Encoding(44), Probability: tsb.No}
	prC := &Probe{Encoding: Encoding(55), Probability: tsb.Maybe}

	// Test #1.
	report = &Report{probes: map[Encoding]*Probe{}}
	tst.MustBeEqual(report.IsAccurate(), true)

	// Test #2.
	report = &Report{
		probes: map[Encoding]*Probe{
			Encoding(55): prC,
		},
	}
	tst.MustBeEqual(report.IsAccurate(), false)

	// Test #3.
	report = &Report{
		probes: map[Encoding]*Probe{
			Encoding(33): prA,
		},
	}
	tst.MustBeEqual(report.IsAccurate(), true)

	// Test #4.
	report = &Report{
		probes: map[Encoding]*Probe{
			Encoding(33): prA,
			Encoding(44): prB,
			Encoding(55): prC,
		},
	}
	tst.MustBeEqual(report.IsAccurate(), false)
}

func Test_GetAccurateProbes(t *testing.T) {
	tst := tester.New(t)
	var report *Report
	var expectedProbes []*Probe

	prA := &Probe{Encoding: Encoding(33), Probability: tsb.Yes}
	prC := &Probe{Encoding: Encoding(55), Probability: tsb.Maybe}

	// Test #1.
	report = &Report{probes: map[Encoding]*Probe{}}
	expectedProbes = []*Probe{}
	tst.MustBeEqual(report.GetAccurateProbes(), expectedProbes)

	// Test #2.
	report = &Report{
		probes: map[Encoding]*Probe{
			Encoding(33): prA,
			Encoding(55): prC,
		},
	}
	expectedProbes = []*Probe{prA}
	tst.MustBeEqual(report.GetAccurateProbes(), expectedProbes)

	// Test #3.
	report = &Report{
		probes: map[Encoding]*Probe{
			Encoding(55): prC,
		},
	}
	expectedProbes = []*Probe{}
	tst.MustBeEqual(report.GetAccurateProbes(), expectedProbes)
}
