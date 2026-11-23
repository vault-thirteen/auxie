package mcci

const (
	ChannelName_Red   = "red"
	ChannelName_Green = "green"
	ChannelName_Blue  = "blue"
	ChannelName_Alpha = "alpha"

	ChannelName_Orange = "orange"
	ChannelName_Yellow = "yellow"
	ChannelName_Cyan   = "cyan"
)

// Wave length were selected according to the online tool:
// https://colorsontheweb.com/color-tools/wavelength-tool
const (
	ChannelWaveLengthNm_Red   = 645
	ChannelWaveLengthNm_Green = 510
	ChannelWaveLengthNm_Blue  = 440

	ChannelWaveLengthNm_Orange = 615
	ChannelWaveLengthNm_Yellow = 580
	ChannelWaveLengthNm_Cyan   = 490

	WaveLengthNm_LowMargin  = 780 // Red - Infrared.
	WaveLengthNm_HighMargin = 380 // Violet - Ultraviolet.
	WaveLengthNm_Tolerance  = 10
	WaveLengthNm_None       = 0
)

type Channel struct {
	Name          string
	WaveLength_Nm float32
}
