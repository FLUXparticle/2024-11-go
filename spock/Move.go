package spock

// Move Interface
type Move interface {
	Verb(other Move) string
	String() string
}

// SteinMove
type SteinMove struct{}

func (m *SteinMove) Verb(other Move) string {
	switch other.(type) {
	case *SchereMove:
		return "schleift"
	}
	return ""
}

func (m *SteinMove) String() string {
	return "Stein"
}

// PapierMove
type PapierMove struct{}

func (m *PapierMove) Verb(other Move) string {
	switch other.(type) {
	case *SteinMove:
		return "umhüllt"
	}
	return ""
}

func (m *PapierMove) String() string {
	return "Papier"
}

// SchereMove
type SchereMove struct{}

func (m *SchereMove) Verb(other Move) string {
	switch other.(type) {
	case *PapierMove:
		return "zerschneidet"
	}
	return ""
}

func (m *SchereMove) String() string {
	return "Schere"
}
