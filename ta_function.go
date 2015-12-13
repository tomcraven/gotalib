package gotalib

type TA_Function interface {
	init()

	GetNumInputs() ( int )
	SetInputData( int, []float64 )
	SetPriceInputData( []float64, []float64, []float64, []float64, []float64, []float64 )

	GetNumFiddleValues() ( int )
	GetFiddleValues() ( []float64 )
	FixFiddleValue( int, float64 ) ( float64 )
	SetFiddleValues( []float64 )

	GetNumOutputValues() ( int )
	Go( int ) ( []float64 )
	GoSingle( int ) ( float64 )
}
