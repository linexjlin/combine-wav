package combine

import (
	"github.com/linexjlin/simple-log"
	"github.com/youpy/go-wav"
	"io"
	"os"
)

func formatEqual(f1, f2 *wav.WavFormat) bool {
	return true
}

func Combine(files []string) (format *wav.WavFormat, samplesDat [][]wav.Sample, sampleCnt uint32) {
	for _, fn := range files {
		if file, err := os.Open(fn); err != nil {
			log.Error(err)
		} else {
			reader := wav.NewReader(file)
			if newFormat, err := reader.Format(); err != nil {
				log.Error(err)
			} else {
				if formatEqual(newFormat, format) || format == nil {
					format = newFormat
					for {
						if newSamples, err := reader.ReadSamples(); err != nil {
							if err == io.EOF {
								break
							}
						} else {
							samplesDat = append(samplesDat, newSamples)
							sampleCnt += uint32(len(newSamples))
						}
					}
				}

			}
			defer file.Close()
		}
	}
	return
}

func WriteToFile(fileName string, format *wav.WavFormat, samplesDat [][]wav.Sample, sampelCnt uint32) {
	outfile, err := os.Create(fileName)
	if err != nil {
		log.Error(err)
	}

	defer outfile.Close()

	var numSamples uint32 = sampelCnt
	var numChannels uint16 = format.NumChannels
	var sampleRate uint32 = format.SampleRate
	var bitsPerSample uint16 = format.BitsPerSample

	writer := wav.NewWriter(outfile, numSamples, numChannels, sampleRate, bitsPerSample)
	for _, samples := range samplesDat {
		err = writer.WriteSamples(samples)
		if err != nil {
			log.Error(err)
		}
	}
	outfile.Close()
}
