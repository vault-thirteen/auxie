package bom

// Report stores the result of making probes for all possible encodings.
type Report struct {
	probes map[Encoding]*Probe
}

// GetEncodingsReport tries to get the report about the specified probes.
// Please note that some encodings have similar BOMs and this fact can make
// probe results inaccurate.
func GetEncodingsReport(data []byte, encodingsToProbe map[Encoding]bool) (report *Report, err error) {
	report = &Report{
		probes: make(map[Encoding]*Probe),
	}

	var probe *Probe
	for enc, _ := range encodingsToProbe {
		probe, err = ProbeForEncoding(data, enc)
		if err != nil {
			return nil, err
		}

		report.probes[enc] = probe
	}

	return report, nil
}

// IsAccurate tells whether all the probes of the report are accurate or not.
func (r *Report) IsAccurate() bool {
	for _, p := range r.probes {
		if !p.IsAccurate() {
			return false
		}
	}

	return true
}

// GetAccurateProbes returns accurate probes of the report.
func (r *Report) GetAccurateProbes() (accurateProbes []*Probe) {
	accurateProbes = make([]*Probe, 0, len(r.probes))

	for _, p := range r.probes {
		if p.IsAccurate() {
			accurateProbes = append(accurateProbes, p)
		}
	}

	return accurateProbes
}
