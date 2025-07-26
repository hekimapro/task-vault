package models

type GridColumn struct {
	SM         int
	MD         int
	LG         int
	MDOffset   int
	LGOffset   int
	IsCenter   bool
	IsRight    bool
	ClassNames []string
}

type Input struct {
	Type        string
	Name        string
	Label       string
	Error       string
	Placeholder string
}

type Button struct {
	Type       string
	Title      string
	ClassNames []string
	IsDisabled bool
}
