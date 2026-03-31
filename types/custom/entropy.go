package custom

import (
	"math"
	"sort"

	"github.com/coalaura/wtf/types"
)

const (
	entropyThresholdEncrypted = 7.9
	entropyThresholdHigh      = 7.5
	entropyThresholdMedium    = 6.0
)

type ByteFrequency [256]int

type EntropyResult struct {
	ShannonEntropy float64
	IsEncrypted    bool
	IsHighEntropy  bool
	ByteFreq       ByteFrequency
	ZeroBytes      int
	PrintableBytes int
	HighBytes      int
}

type byteCount struct {
	byte  byte
	count int
}

func CalculateEntropy(b types.Buffer) *EntropyResult {
	if b.Len() == 0 {
		return nil
	}

	limit := min(b.Len(), 65536)
	data := b[:limit]

	var (
		freq      ByteFrequency
		zeros     int
		printable int
		high      int
	)

	for _, c := range data {
		freq[c]++

		if c == 0 {
			zeros++
		}

		if c >= 0x20 && c <= 0x7e || c == '\n' || c == '\r' || c == '\t' {
			printable++
		}

		if c >= 0x80 {
			high++
		}
	}

	entropy := 0.0
	total := float64(len(data))

	for _, count := range freq {
		if count == 0 {
			continue
		}

		p := float64(count) / total
		entropy -= p * math.Log2(p)
	}

	return &EntropyResult{
		ShannonEntropy: entropy,
		IsEncrypted:    entropy >= entropyThresholdEncrypted && zeros == 0,
		IsHighEntropy:  entropy >= entropyThresholdHigh,
		ByteFreq:       freq,
		ZeroBytes:      zeros,
		PrintableBytes: printable,
		HighBytes:      high,
	}
}

func DetectHighEntropy(b types.Buffer) *types.Metadata {
	result := CalculateEntropy(b)
	if result == nil {
		return nil
	}

	// Only flag as high entropy if we have enough data to be meaningful
	if b.Len() < 256 {
		return nil
	}

	if result.IsEncrypted {
		return &types.Metadata{
			Kind:       types.KindHighEntropyData,
			Type:       types.TypeEncrypted,
			Confidence: types.ConfidenceMedium,
		}
	}

	if result.IsHighEntropy {
		return &types.Metadata{
			Kind:       types.KindHighEntropyData,
			Type:       types.TypeHighEntropy,
			Confidence: types.ConfidenceLow,
		}
	}

	return nil
}

func GetByteDistributionSummary(b types.Buffer) map[string]float64 {
	result := CalculateEntropy(b)
	if result == nil {
		return nil
	}

	total := float64(min(b.Len(), 65536))

	summary := map[string]float64{
		"entropy":         math.Round(result.ShannonEntropy*1000) / 1000,
		"zero_ratio":      math.Round(float64(result.ZeroBytes)/total*10000) / 10000,
		"printable_ratio": math.Round(float64(result.PrintableBytes)/total*10000) / 10000,
		"high_byte_ratio": math.Round(float64(result.HighBytes)/total*10000) / 10000,
	}

	var uniqueBytes int

	for _, count := range result.ByteFreq {
		if count > 0 {
			uniqueBytes++
		}
	}

	summary["unique_byte_ratio"] = math.Round(float64(uniqueBytes)/256*10000) / 10000

	var topBytes []byteCount

	for i, count := range result.ByteFreq {
		if count > 0 {
			topBytes = append(topBytes, byteCount{byte(i), count})
		}
	}

	sort.Slice(topBytes, func(i, j int) bool {
		return topBytes[i].count > topBytes[j].count
	})

	for i, bc := range topBytes {
		if i >= 5 {
			break
		}

		key := string(rune('a'+i)) + "_most_freq"
		summary[key] = math.Round(float64(bc.count)/total*10000) / 10000
	}

	return summary
}
