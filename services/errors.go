package services

type SoftError struct {
	Err error
}

func (e *SoftError) Error() string {
	return e.Err.Error()
}

type HardError struct {
	Err error
}

func (e *HardError) Error() string {
	return e.Err.Error()
}
