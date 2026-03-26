package types

import "bytes"

const (
	ansiReset  = "\x1b[0m"
	ansiFile   = "\x1b[1;92m"
	ansiName   = "\x1b[1;97m"
	ansiLabel  = "\x1b[90m"
	ansiValue  = "\x1b[3m"
	ansiHigh   = "\x1b[32m"
	ansiMedium = "\x1b[33m"
	ansiLow    = "\x1b[31m"
)

func (m *Metadata) Format() string {
	if m == nil {
		return "Unknown"
	}

	var b bytes.Buffer

	kind := m.Kind.String()
	typ := m.Type.String()

	b.Grow(len(m.File) + len(kind) + len(typ) + 64)

	if m.File != "" {
		b.WriteString(ansiFile)
		b.WriteString(m.File)
		b.WriteString(ansiReset)
		b.WriteByte('\n')
	}

	if kind != "" {
		b.WriteString("  ")
		b.WriteString(ansiFile)
		b.WriteString("└─ ")
		b.WriteString(ansiReset)
		b.WriteString(ansiName)
		b.WriteString(kind)
		b.WriteString(ansiReset)

		b.WriteString(" ")
		b.WriteString(ansiLabel)
		b.WriteString("[")
		b.WriteString(ansiReset)

		switch m.Confidence {
		case ConfidenceHigh:
			b.WriteString(ansiHigh)
		case ConfidenceMedium:
			b.WriteString(ansiMedium)
		case ConfidenceLow:
			b.WriteString(ansiLow)
		}

		b.WriteString(m.Confidence.String())
		b.WriteString(ansiReset)

		b.WriteString(ansiLabel)
		b.WriteString("]")
		b.WriteString(ansiReset)
	}

	if typ != "" {
		b.WriteByte('\n')
		b.WriteString("     ")
		b.WriteString(ansiLabel)
		b.WriteString("Type: ")
		b.WriteString(ansiReset)
		b.WriteString(ansiValue)
		b.WriteString(typ)
		b.WriteString(ansiReset)
	}

	return b.String()
}
