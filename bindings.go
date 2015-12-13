package gotalib

import "math"

/*
#cgo LDFLAGS: -L/home/tom/dev/ta-lib/src/.libs/ -lta_lib -lm
#include "/home/tom/dev/ta-lib/include/ta_abstract.h"
*/
import "C"

type add_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *add_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *add_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *add_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
	if index == 1 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *add_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *add_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *add_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *add_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *add_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *add_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *add_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *add_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Add() TA_Function {
	var ret add_struct
	helper_getFunctionHandle( "ADD", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type div_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *div_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *div_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *div_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
	if index == 1 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *div_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *div_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *div_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *div_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *div_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *div_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *div_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *div_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Div() TA_Function {
	var ret div_struct
	helper_getFunctionHandle( "DIV", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type max_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *max_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *max_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *max_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *max_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *max_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *max_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *max_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *max_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *max_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *max_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *max_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Max() TA_Function {
	var ret max_struct
	helper_getFunctionHandle( "MAX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type maxIndex_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *maxIndex_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *maxIndex_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *maxIndex_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *maxIndex_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *maxIndex_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *maxIndex_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *maxIndex_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *maxIndex_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *maxIndex_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *maxIndex_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *maxIndex_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MaxIndex() TA_Function {
	var ret maxIndex_struct
	helper_getFunctionHandle( "MAXINDEX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type min_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *min_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *min_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *min_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *min_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *min_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *min_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *min_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *min_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *min_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *min_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *min_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Min() TA_Function {
	var ret min_struct
	helper_getFunctionHandle( "MIN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type minIndex_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *minIndex_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *minIndex_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *minIndex_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *minIndex_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *minIndex_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *minIndex_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *minIndex_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *minIndex_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *minIndex_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *minIndex_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *minIndex_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MinIndex() TA_Function {
	var ret minIndex_struct
	helper_getFunctionHandle( "MININDEX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type minMax_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *minMax_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *minMax_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *minMax_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *minMax_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *minMax_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *minMax_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *minMax_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *minMax_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *minMax_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *minMax_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *minMax_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MinMax() TA_Function {
	var ret minMax_struct
	helper_getFunctionHandle( "MINMAX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type minMaxIndex_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *minMaxIndex_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *minMaxIndex_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *minMaxIndex_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *minMaxIndex_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *minMaxIndex_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *minMaxIndex_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *minMaxIndex_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *minMaxIndex_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *minMaxIndex_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *minMaxIndex_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )
	a.integerOutputByIndex[1] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 1, &(a.integerOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *minMaxIndex_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MinMaxIndex() TA_Function {
	var ret minMaxIndex_struct
	helper_getFunctionHandle( "MINMAXINDEX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type mult_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *mult_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *mult_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *mult_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
	if index == 1 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *mult_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *mult_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *mult_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *mult_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *mult_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *mult_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *mult_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *mult_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Mult() TA_Function {
	var ret mult_struct
	helper_getFunctionHandle( "MULT", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sub_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sub_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *sub_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *sub_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
	if index == 1 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *sub_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *sub_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *sub_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sub_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sub_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sub_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sub_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sub_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Sub() TA_Function {
	var ret sub_struct
	helper_getFunctionHandle( "SUB", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sum_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sum_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *sum_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *sum_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *sum_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *sum_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *sum_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sum_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sum_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sum_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sum_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sum_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Sum() TA_Function {
	var ret sum_struct
	helper_getFunctionHandle( "SUM", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type acos_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *acos_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *acos_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *acos_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *acos_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *acos_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *acos_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *acos_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *acos_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *acos_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *acos_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *acos_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Acos() TA_Function {
	var ret acos_struct
	helper_getFunctionHandle( "ACOS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type asin_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *asin_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *asin_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *asin_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *asin_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *asin_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *asin_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *asin_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *asin_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *asin_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *asin_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *asin_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Asin() TA_Function {
	var ret asin_struct
	helper_getFunctionHandle( "ASIN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type atan_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *atan_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *atan_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *atan_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *atan_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *atan_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *atan_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *atan_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *atan_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *atan_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *atan_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *atan_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Atan() TA_Function {
	var ret atan_struct
	helper_getFunctionHandle( "ATAN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type ceil_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *ceil_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *ceil_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *ceil_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *ceil_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *ceil_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *ceil_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *ceil_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *ceil_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *ceil_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *ceil_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *ceil_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Ceil() TA_Function {
	var ret ceil_struct
	helper_getFunctionHandle( "CEIL", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cos_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cos_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cos_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cos_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *cos_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *cos_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cos_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cos_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cos_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cos_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cos_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cos_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cos() TA_Function {
	var ret cos_struct
	helper_getFunctionHandle( "COS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cosh_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cosh_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cosh_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cosh_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *cosh_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *cosh_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cosh_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cosh_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cosh_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cosh_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cosh_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cosh_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cosh() TA_Function {
	var ret cosh_struct
	helper_getFunctionHandle( "COSH", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type exp_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *exp_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *exp_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *exp_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *exp_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *exp_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *exp_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *exp_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *exp_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *exp_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *exp_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *exp_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Exp() TA_Function {
	var ret exp_struct
	helper_getFunctionHandle( "EXP", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type floor_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *floor_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *floor_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *floor_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *floor_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *floor_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *floor_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *floor_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *floor_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *floor_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *floor_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *floor_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Floor() TA_Function {
	var ret floor_struct
	helper_getFunctionHandle( "FLOOR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type ln_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *ln_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *ln_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *ln_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *ln_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *ln_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *ln_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *ln_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *ln_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *ln_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *ln_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *ln_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Ln() TA_Function {
	var ret ln_struct
	helper_getFunctionHandle( "LN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type log10_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *log10_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *log10_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *log10_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *log10_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *log10_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *log10_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *log10_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *log10_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *log10_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *log10_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *log10_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Log10() TA_Function {
	var ret log10_struct
	helper_getFunctionHandle( "LOG10", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sin_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sin_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *sin_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *sin_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *sin_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *sin_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *sin_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sin_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sin_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sin_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sin_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sin_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Sin() TA_Function {
	var ret sin_struct
	helper_getFunctionHandle( "SIN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sinh_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sinh_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *sinh_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *sinh_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *sinh_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *sinh_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *sinh_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sinh_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sinh_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sinh_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sinh_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sinh_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Sinh() TA_Function {
	var ret sinh_struct
	helper_getFunctionHandle( "SINH", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sqrt_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sqrt_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *sqrt_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *sqrt_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *sqrt_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *sqrt_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *sqrt_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sqrt_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sqrt_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sqrt_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sqrt_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sqrt_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Sqrt() TA_Function {
	var ret sqrt_struct
	helper_getFunctionHandle( "SQRT", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type tan_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *tan_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *tan_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *tan_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *tan_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *tan_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *tan_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *tan_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *tan_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *tan_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *tan_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *tan_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Tan() TA_Function {
	var ret tan_struct
	helper_getFunctionHandle( "TAN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type tanh_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *tanh_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *tanh_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *tanh_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *tanh_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *tanh_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *tanh_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *tanh_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *tanh_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *tanh_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *tanh_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *tanh_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Tanh() TA_Function {
	var ret tanh_struct
	helper_getFunctionHandle( "TANH", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type bbands_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *bbands_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 4 )

	a.fiddleValues[0] = float64( 5.0000000000 )
	a.fiddleValues[1] = float64( 2.0000000000 )
	a.fiddleValues[2] = float64( 2.0000000000 )
	a.fiddleValues[3] = float64( 0.0000000000 )
}

func ( a *bbands_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *bbands_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *bbands_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *bbands_struct ) GetNumFiddleValues() int {
	return 4
}

func ( a *bbands_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *bbands_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 4 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *bbands_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return -2.000000 + ( ( 2.000000 - -2.000000 ) * inValue )
	}

	if fiddleValueIndex == 2 {
		return -2.000000 + ( ( 2.000000 - -2.000000 ) * inValue )
	}

	if fiddleValueIndex == 3 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *bbands_struct ) GetNumOutputValues() ( int ) {
	return 3
}

func ( a *bbands_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataReal( a.params, 1, a.fiddleValues[1] )
	helper_setOptInputDataReal( a.params, 2, a.fiddleValues[2] )
	helper_setOptInputDataInteger( a.params, 3, int( a.fiddleValues[3] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )
	a.realOutputByIndex[2] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 2, &(a.realOutputByIndex[2][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 2 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *bbands_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Bbands() TA_Function {
	var ret bbands_struct
	helper_getFunctionHandle( "BBANDS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type dema_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *dema_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *dema_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *dema_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *dema_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *dema_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *dema_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *dema_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *dema_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *dema_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *dema_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *dema_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Dema() TA_Function {
	var ret dema_struct
	helper_getFunctionHandle( "DEMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type ema_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *ema_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *ema_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *ema_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *ema_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *ema_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *ema_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *ema_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *ema_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *ema_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *ema_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *ema_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Ema() TA_Function {
	var ret ema_struct
	helper_getFunctionHandle( "EMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type htTrendline_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *htTrendline_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *htTrendline_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *htTrendline_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *htTrendline_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *htTrendline_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *htTrendline_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *htTrendline_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *htTrendline_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *htTrendline_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *htTrendline_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *htTrendline_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func HtTrendline() TA_Function {
	var ret htTrendline_struct
	helper_getFunctionHandle( "HT_TRENDLINE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type kama_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *kama_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *kama_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *kama_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *kama_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *kama_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *kama_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *kama_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *kama_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *kama_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *kama_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *kama_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Kama() TA_Function {
	var ret kama_struct
	helper_getFunctionHandle( "KAMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type movingAverage_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *movingAverage_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 2 )

	a.fiddleValues[0] = float64( 30.0000000000 )
	a.fiddleValues[1] = float64( 0.0000000000 )
}

func ( a *movingAverage_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *movingAverage_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *movingAverage_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *movingAverage_struct ) GetNumFiddleValues() int {
	return 2
}

func ( a *movingAverage_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *movingAverage_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 2 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *movingAverage_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *movingAverage_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *movingAverage_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *movingAverage_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MovingAverage() TA_Function {
	var ret movingAverage_struct
	helper_getFunctionHandle( "MA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type mama_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *mama_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 2 )

	a.fiddleValues[0] = float64( 0.5000000000 )
	a.fiddleValues[1] = float64( 0.0500000000 )
}

func ( a *mama_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *mama_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *mama_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *mama_struct ) GetNumFiddleValues() int {
	return 2
}

func ( a *mama_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *mama_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 2 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *mama_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.210000 + ( ( 0.800000 - 0.210000 ) * inValue )
	}

	if fiddleValueIndex == 1 {
		return 0.010000 + ( ( 0.600000 - 0.010000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *mama_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *mama_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )
	helper_setOptInputDataReal( a.params, 1, a.fiddleValues[1] )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *mama_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Mama() TA_Function {
	var ret mama_struct
	helper_getFunctionHandle( "MAMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type movingAverageVariablePeriod_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *movingAverageVariablePeriod_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 3 )

	a.fiddleValues[0] = float64( 2.0000000000 )
	a.fiddleValues[1] = float64( 30.0000000000 )
	a.fiddleValues[2] = float64( 0.0000000000 )
}

func ( a *movingAverageVariablePeriod_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *movingAverageVariablePeriod_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
	if index == 1 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *movingAverageVariablePeriod_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *movingAverageVariablePeriod_struct ) GetNumFiddleValues() int {
	return 3
}

func ( a *movingAverageVariablePeriod_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *movingAverageVariablePeriod_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 3 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *movingAverageVariablePeriod_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *movingAverageVariablePeriod_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *movingAverageVariablePeriod_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *movingAverageVariablePeriod_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[1] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MovingAverageVariablePeriod() TA_Function {
	var ret movingAverageVariablePeriod_struct
	helper_getFunctionHandle( "MAVP", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type midPoint_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *midPoint_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *midPoint_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *midPoint_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *midPoint_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *midPoint_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *midPoint_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *midPoint_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *midPoint_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *midPoint_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *midPoint_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *midPoint_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MidPoint() TA_Function {
	var ret midPoint_struct
	helper_getFunctionHandle( "MIDPOINT", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type midPrice_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *midPrice_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *midPrice_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *midPrice_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *midPrice_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *midPrice_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *midPrice_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *midPrice_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *midPrice_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *midPrice_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *midPrice_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *midPrice_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func MidPrice() TA_Function {
	var ret midPrice_struct
	helper_getFunctionHandle( "MIDPRICE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 2 )

	a.fiddleValues[0] = float64( 0.0200000000 )
	a.fiddleValues[1] = float64( 0.2000000000 )
}

func ( a *sar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *sar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *sar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *sar_struct ) GetNumFiddleValues() int {
	return 2
}

func ( a *sar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 2 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.010000 + ( ( 0.200000 - 0.010000 ) * inValue )
	}

	if fiddleValueIndex == 1 {
		return 0.200000 + ( ( 0.400000 - 0.200000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )
	helper_setOptInputDataReal( a.params, 1, a.fiddleValues[1] )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Sar() TA_Function {
	var ret sar_struct
	helper_getFunctionHandle( "SAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sarExt_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sarExt_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 8 )

	a.fiddleValues[0] = float64( 0.0000000000 )
	a.fiddleValues[1] = float64( 0.0000000000 )
	a.fiddleValues[2] = float64( 0.0200000000 )
	a.fiddleValues[3] = float64( 0.0200000000 )
	a.fiddleValues[4] = float64( 0.2000000000 )
	a.fiddleValues[5] = float64( 0.0200000000 )
	a.fiddleValues[6] = float64( 0.0200000000 )
	a.fiddleValues[7] = float64( 0.2000000000 )
}

func ( a *sarExt_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *sarExt_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *sarExt_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *sarExt_struct ) GetNumFiddleValues() int {
	return 8
}

func ( a *sarExt_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sarExt_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 8 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sarExt_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	if fiddleValueIndex == 1 {
		return 0.010000 + ( ( 0.150000 - 0.010000 ) * inValue )
	}

	if fiddleValueIndex == 2 {
		return 0.010000 + ( ( 0.190000 - 0.010000 ) * inValue )
	}

	if fiddleValueIndex == 3 {
		return 0.010000 + ( ( 0.200000 - 0.010000 ) * inValue )
	}

	if fiddleValueIndex == 4 {
		return 0.200000 + ( ( 0.400000 - 0.200000 ) * inValue )
	}

	if fiddleValueIndex == 5 {
		return 0.010000 + ( ( 0.190000 - 0.010000 ) * inValue )
	}

	if fiddleValueIndex == 6 {
		return 0.010000 + ( ( 0.200000 - 0.010000 ) * inValue )
	}

	if fiddleValueIndex == 7 {
		return 0.200000 + ( ( 0.400000 - 0.200000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sarExt_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sarExt_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )
	helper_setOptInputDataReal( a.params, 1, a.fiddleValues[1] )
	helper_setOptInputDataReal( a.params, 2, a.fiddleValues[2] )
	helper_setOptInputDataReal( a.params, 3, a.fiddleValues[3] )
	helper_setOptInputDataReal( a.params, 4, a.fiddleValues[4] )
	helper_setOptInputDataReal( a.params, 5, a.fiddleValues[5] )
	helper_setOptInputDataReal( a.params, 6, a.fiddleValues[6] )
	helper_setOptInputDataReal( a.params, 7, a.fiddleValues[7] )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sarExt_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func SarExt() TA_Function {
	var ret sarExt_struct
	helper_getFunctionHandle( "SAREXT", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type sma_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *sma_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *sma_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *sma_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *sma_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *sma_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *sma_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *sma_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *sma_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *sma_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *sma_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *sma_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Sma() TA_Function {
	var ret sma_struct
	helper_getFunctionHandle( "SMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type t3_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *t3_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 2 )

	a.fiddleValues[0] = float64( 5.0000000000 )
	a.fiddleValues[1] = float64( 0.7000000000 )
}

func ( a *t3_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *t3_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *t3_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *t3_struct ) GetNumFiddleValues() int {
	return 2
}

func ( a *t3_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *t3_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 2 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *t3_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return 0.010000 + ( ( 1.000000 - 0.010000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *t3_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *t3_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataReal( a.params, 1, a.fiddleValues[1] )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *t3_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func T3() TA_Function {
	var ret t3_struct
	helper_getFunctionHandle( "T3", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type tema_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *tema_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *tema_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *tema_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *tema_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *tema_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *tema_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *tema_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *tema_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *tema_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *tema_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *tema_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Tema() TA_Function {
	var ret tema_struct
	helper_getFunctionHandle( "TEMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type trima_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *trima_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *trima_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *trima_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *trima_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *trima_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *trima_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *trima_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *trima_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *trima_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *trima_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *trima_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Trima() TA_Function {
	var ret trima_struct
	helper_getFunctionHandle( "TRIMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type wma_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *wma_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *wma_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *wma_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *wma_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *wma_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *wma_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *wma_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *wma_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *wma_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *wma_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *wma_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Wma() TA_Function {
	var ret wma_struct
	helper_getFunctionHandle( "WMA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type atr_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *atr_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *atr_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *atr_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *atr_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *atr_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *atr_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *atr_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *atr_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *atr_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *atr_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *atr_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Atr() TA_Function {
	var ret atr_struct
	helper_getFunctionHandle( "ATR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type natr_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *natr_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *natr_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *natr_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *natr_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *natr_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *natr_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *natr_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *natr_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *natr_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *natr_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *natr_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Natr() TA_Function {
	var ret natr_struct
	helper_getFunctionHandle( "NATR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type trueRange_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *trueRange_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *trueRange_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *trueRange_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *trueRange_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *trueRange_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *trueRange_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *trueRange_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *trueRange_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *trueRange_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *trueRange_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *trueRange_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func TrueRange() TA_Function {
	var ret trueRange_struct
	helper_getFunctionHandle( "TRANGE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type adx_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *adx_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *adx_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *adx_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *adx_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *adx_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *adx_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *adx_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *adx_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *adx_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *adx_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *adx_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Adx() TA_Function {
	var ret adx_struct
	helper_getFunctionHandle( "ADX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type adxr_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *adxr_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *adxr_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *adxr_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *adxr_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *adxr_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *adxr_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *adxr_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *adxr_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *adxr_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *adxr_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *adxr_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Adxr() TA_Function {
	var ret adxr_struct
	helper_getFunctionHandle( "ADXR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type apo_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *apo_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 3 )

	a.fiddleValues[0] = float64( 12.0000000000 )
	a.fiddleValues[1] = float64( 26.0000000000 )
	a.fiddleValues[2] = float64( 0.0000000000 )
}

func ( a *apo_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *apo_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *apo_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *apo_struct ) GetNumFiddleValues() int {
	return 3
}

func ( a *apo_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *apo_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 3 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *apo_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *apo_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *apo_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *apo_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[1] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Apo() TA_Function {
	var ret apo_struct
	helper_getFunctionHandle( "APO", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type aroon_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *aroon_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *aroon_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *aroon_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *aroon_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *aroon_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *aroon_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *aroon_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *aroon_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *aroon_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *aroon_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *aroon_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Aroon() TA_Function {
	var ret aroon_struct
	helper_getFunctionHandle( "AROON", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type aroonOsc_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *aroonOsc_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *aroonOsc_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *aroonOsc_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *aroonOsc_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *aroonOsc_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *aroonOsc_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *aroonOsc_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *aroonOsc_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *aroonOsc_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *aroonOsc_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *aroonOsc_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func AroonOsc() TA_Function {
	var ret aroonOsc_struct
	helper_getFunctionHandle( "AROONOSC", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type bop_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *bop_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *bop_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *bop_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *bop_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *bop_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *bop_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *bop_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *bop_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *bop_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *bop_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *bop_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Bop() TA_Function {
	var ret bop_struct
	helper_getFunctionHandle( "BOP", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cci_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cci_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *cci_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cci_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cci_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cci_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cci_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cci_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cci_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cci_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cci_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cci_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cci() TA_Function {
	var ret cci_struct
	helper_getFunctionHandle( "CCI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cmo_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cmo_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *cmo_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cmo_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *cmo_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *cmo_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cmo_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cmo_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cmo_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cmo_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cmo_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cmo_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Cmo() TA_Function {
	var ret cmo_struct
	helper_getFunctionHandle( "CMO", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type dx_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *dx_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *dx_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *dx_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *dx_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *dx_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *dx_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *dx_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *dx_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *dx_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *dx_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *dx_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Dx() TA_Function {
	var ret dx_struct
	helper_getFunctionHandle( "DX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type macd_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *macd_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 3 )

	a.fiddleValues[0] = float64( 12.0000000000 )
	a.fiddleValues[1] = float64( 26.0000000000 )
	a.fiddleValues[2] = float64( 9.0000000000 )
}

func ( a *macd_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *macd_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *macd_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *macd_struct ) GetNumFiddleValues() int {
	return 3
}

func ( a *macd_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *macd_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 3 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *macd_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *macd_struct ) GetNumOutputValues() ( int ) {
	return 3
}

func ( a *macd_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )
	a.realOutputByIndex[2] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 2, &(a.realOutputByIndex[2][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 2 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *macd_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[1] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[2] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Macd() TA_Function {
	var ret macd_struct
	helper_getFunctionHandle( "MACD", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type macdExt_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *macdExt_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 6 )

	a.fiddleValues[0] = float64( 12.0000000000 )
	a.fiddleValues[1] = float64( 0.0000000000 )
	a.fiddleValues[2] = float64( 26.0000000000 )
	a.fiddleValues[3] = float64( 0.0000000000 )
	a.fiddleValues[4] = float64( 9.0000000000 )
	a.fiddleValues[5] = float64( 0.0000000000 )
}

func ( a *macdExt_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *macdExt_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *macdExt_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *macdExt_struct ) GetNumFiddleValues() int {
	return 6
}

func ( a *macdExt_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *macdExt_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 6 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *macdExt_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	if fiddleValueIndex == 2 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 3 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	if fiddleValueIndex == 4 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 5 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *macdExt_struct ) GetNumOutputValues() ( int ) {
	return 3
}

func ( a *macdExt_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )
	helper_setOptInputDataInteger( a.params, 3, int( a.fiddleValues[3] ) )
	helper_setOptInputDataInteger( a.params, 4, int( a.fiddleValues[4] ) )
	helper_setOptInputDataInteger( a.params, 5, int( a.fiddleValues[5] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )
	a.realOutputByIndex[2] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 2, &(a.realOutputByIndex[2][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 2 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *macdExt_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[2] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[4] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MacdExt() TA_Function {
	var ret macdExt_struct
	helper_getFunctionHandle( "MACDEXT", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type macdFix_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *macdFix_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 9.0000000000 )
}

func ( a *macdFix_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *macdFix_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *macdFix_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *macdFix_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *macdFix_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *macdFix_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *macdFix_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *macdFix_struct ) GetNumOutputValues() ( int ) {
	return 3
}

func ( a *macdFix_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )
	a.realOutputByIndex[2] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 2, &(a.realOutputByIndex[2][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 2 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *macdFix_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func MacdFix() TA_Function {
	var ret macdFix_struct
	helper_getFunctionHandle( "MACDFIX", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type mfi_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *mfi_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *mfi_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *mfi_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *mfi_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *mfi_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *mfi_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *mfi_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *mfi_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *mfi_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *mfi_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *mfi_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Mfi() TA_Function {
	var ret mfi_struct
	helper_getFunctionHandle( "MFI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type minusDI_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *minusDI_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *minusDI_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *minusDI_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *minusDI_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *minusDI_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *minusDI_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *minusDI_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *minusDI_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *minusDI_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *minusDI_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *minusDI_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func MinusDI() TA_Function {
	var ret minusDI_struct
	helper_getFunctionHandle( "MINUS_DI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type minusDM_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *minusDM_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *minusDM_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *minusDM_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *minusDM_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *minusDM_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *minusDM_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *minusDM_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *minusDM_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *minusDM_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *minusDM_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *minusDM_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func MinusDM() TA_Function {
	var ret minusDM_struct
	helper_getFunctionHandle( "MINUS_DM", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type mom_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *mom_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 10.0000000000 )
}

func ( a *mom_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *mom_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *mom_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *mom_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *mom_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *mom_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *mom_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *mom_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *mom_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *mom_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Mom() TA_Function {
	var ret mom_struct
	helper_getFunctionHandle( "MOM", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type plusDI_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *plusDI_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *plusDI_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *plusDI_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *plusDI_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *plusDI_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *plusDI_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *plusDI_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *plusDI_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *plusDI_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *plusDI_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *plusDI_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func PlusDI() TA_Function {
	var ret plusDI_struct
	helper_getFunctionHandle( "PLUS_DI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type plusDM_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *plusDM_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *plusDM_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *plusDM_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *plusDM_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *plusDM_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *plusDM_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *plusDM_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *plusDM_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *plusDM_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *plusDM_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *plusDM_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func PlusDM() TA_Function {
	var ret plusDM_struct
	helper_getFunctionHandle( "PLUS_DM", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type ppo_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *ppo_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 3 )

	a.fiddleValues[0] = float64( 12.0000000000 )
	a.fiddleValues[1] = float64( 26.0000000000 )
	a.fiddleValues[2] = float64( 0.0000000000 )
}

func ( a *ppo_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *ppo_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *ppo_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *ppo_struct ) GetNumFiddleValues() int {
	return 3
}

func ( a *ppo_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *ppo_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 3 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *ppo_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *ppo_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *ppo_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *ppo_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[1] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Ppo() TA_Function {
	var ret ppo_struct
	helper_getFunctionHandle( "PPO", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type roc_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *roc_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 10.0000000000 )
}

func ( a *roc_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *roc_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *roc_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *roc_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *roc_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *roc_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *roc_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *roc_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *roc_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *roc_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Roc() TA_Function {
	var ret roc_struct
	helper_getFunctionHandle( "ROC", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type rocP_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *rocP_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 10.0000000000 )
}

func ( a *rocP_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *rocP_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *rocP_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *rocP_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *rocP_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *rocP_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *rocP_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *rocP_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *rocP_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *rocP_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func RocP() TA_Function {
	var ret rocP_struct
	helper_getFunctionHandle( "ROCP", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type rocR_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *rocR_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 10.0000000000 )
}

func ( a *rocR_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *rocR_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *rocR_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *rocR_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *rocR_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *rocR_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *rocR_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *rocR_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *rocR_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *rocR_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func RocR() TA_Function {
	var ret rocR_struct
	helper_getFunctionHandle( "ROCR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type rocR100_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *rocR100_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 10.0000000000 )
}

func ( a *rocR100_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *rocR100_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *rocR100_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *rocR100_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *rocR100_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *rocR100_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *rocR100_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *rocR100_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *rocR100_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *rocR100_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func RocR100() TA_Function {
	var ret rocR100_struct
	helper_getFunctionHandle( "ROCR100", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type rsi_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *rsi_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *rsi_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *rsi_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *rsi_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *rsi_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *rsi_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *rsi_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *rsi_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *rsi_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *rsi_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *rsi_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Rsi() TA_Function {
	var ret rsi_struct
	helper_getFunctionHandle( "RSI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type stoch_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *stoch_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 5 )

	a.fiddleValues[0] = float64( 5.0000000000 )
	a.fiddleValues[1] = float64( 3.0000000000 )
	a.fiddleValues[2] = float64( 0.0000000000 )
	a.fiddleValues[3] = float64( 3.0000000000 )
	a.fiddleValues[4] = float64( 0.0000000000 )
}

func ( a *stoch_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *stoch_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *stoch_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *stoch_struct ) GetNumFiddleValues() int {
	return 5
}

func ( a *stoch_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *stoch_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 5 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *stoch_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	if fiddleValueIndex == 3 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 4 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *stoch_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *stoch_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )
	helper_setOptInputDataInteger( a.params, 3, int( a.fiddleValues[3] ) )
	helper_setOptInputDataInteger( a.params, 4, int( a.fiddleValues[4] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *stoch_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Stoch() TA_Function {
	var ret stoch_struct
	helper_getFunctionHandle( "STOCH", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type stochF_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *stochF_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 3 )

	a.fiddleValues[0] = float64( 5.0000000000 )
	a.fiddleValues[1] = float64( 3.0000000000 )
	a.fiddleValues[2] = float64( 0.0000000000 )
}

func ( a *stochF_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *stochF_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *stochF_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *stochF_struct ) GetNumFiddleValues() int {
	return 3
}

func ( a *stochF_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *stochF_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 3 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *stochF_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *stochF_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *stochF_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *stochF_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func StochF() TA_Function {
	var ret stochF_struct
	helper_getFunctionHandle( "STOCHF", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type stochRsi_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *stochRsi_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 4 )

	a.fiddleValues[0] = float64( 14.0000000000 )
	a.fiddleValues[1] = float64( 5.0000000000 )
	a.fiddleValues[2] = float64( 3.0000000000 )
	a.fiddleValues[3] = float64( 0.0000000000 )
}

func ( a *stochRsi_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *stochRsi_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *stochRsi_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *stochRsi_struct ) GetNumFiddleValues() int {
	return 4
}

func ( a *stochRsi_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *stochRsi_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 4 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *stochRsi_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 3 {
		index := int( inValue * float64( 9 ) )
		indexToRetMap := map[int]int {
			0 : 0,
			1 : 1,
			2 : 2,
			3 : 3,
			4 : 4,
			5 : 5,
			6 : 6,
			7 : 7,
			8 : 8,
		}
		return float64( indexToRetMap[index] )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *stochRsi_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *stochRsi_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )
	helper_setOptInputDataInteger( a.params, 3, int( a.fiddleValues[3] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *stochRsi_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[1] = float64( len( a.realInputByIndex[0] ) )
	a.fiddleValues[2] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func StochRsi() TA_Function {
	var ret stochRsi_struct
	helper_getFunctionHandle( "STOCHRSI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type ultOsc_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *ultOsc_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 3 )

	a.fiddleValues[0] = float64( 7.0000000000 )
	a.fiddleValues[1] = float64( 14.0000000000 )
	a.fiddleValues[2] = float64( 28.0000000000 )
}

func ( a *ultOsc_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *ultOsc_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *ultOsc_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *ultOsc_struct ) GetNumFiddleValues() int {
	return 3
}

func ( a *ultOsc_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *ultOsc_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 3 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *ultOsc_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 2 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *ultOsc_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *ultOsc_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )
	helper_setOptInputDataInteger( a.params, 2, int( a.fiddleValues[2] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *ultOsc_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func UltOsc() TA_Function {
	var ret ultOsc_struct
	helper_getFunctionHandle( "ULTOSC", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type willR_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *willR_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *willR_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *willR_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *willR_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *willR_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *willR_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *willR_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *willR_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *willR_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *willR_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *willR_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func WillR() TA_Function {
	var ret willR_struct
	helper_getFunctionHandle( "WILLR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type htDcPeriod_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *htDcPeriod_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *htDcPeriod_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *htDcPeriod_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *htDcPeriod_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *htDcPeriod_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *htDcPeriod_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *htDcPeriod_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *htDcPeriod_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *htDcPeriod_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *htDcPeriod_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *htDcPeriod_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func HtDcPeriod() TA_Function {
	var ret htDcPeriod_struct
	helper_getFunctionHandle( "HT_DCPERIOD", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type htDcPhase_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *htDcPhase_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *htDcPhase_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *htDcPhase_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *htDcPhase_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *htDcPhase_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *htDcPhase_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *htDcPhase_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *htDcPhase_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *htDcPhase_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *htDcPhase_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *htDcPhase_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func HtDcPhase() TA_Function {
	var ret htDcPhase_struct
	helper_getFunctionHandle( "HT_DCPHASE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type htPhasor_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *htPhasor_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *htPhasor_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *htPhasor_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *htPhasor_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *htPhasor_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *htPhasor_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *htPhasor_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *htPhasor_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *htPhasor_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *htPhasor_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *htPhasor_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func HtPhasor() TA_Function {
	var ret htPhasor_struct
	helper_getFunctionHandle( "HT_PHASOR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type htSine_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *htSine_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *htSine_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *htSine_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *htSine_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *htSine_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *htSine_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *htSine_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *htSine_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *htSine_struct ) GetNumOutputValues() ( int ) {
	return 2
}

func ( a *htSine_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )
	a.realOutputByIndex[1] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 1, &(a.realOutputByIndex[1][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}
	if outIndex == 1 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *htSine_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func HtSine() TA_Function {
	var ret htSine_struct
	helper_getFunctionHandle( "HT_SINE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type htTrendMode_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *htTrendMode_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *htTrendMode_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *htTrendMode_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *htTrendMode_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *htTrendMode_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *htTrendMode_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *htTrendMode_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *htTrendMode_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *htTrendMode_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *htTrendMode_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *htTrendMode_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func HtTrendMode() TA_Function {
	var ret htTrendMode_struct
	helper_getFunctionHandle( "HT_TRENDMODE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type ad_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *ad_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *ad_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *ad_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *ad_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *ad_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *ad_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *ad_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *ad_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *ad_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *ad_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *ad_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Ad() TA_Function {
	var ret ad_struct
	helper_getFunctionHandle( "AD", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type adOsc_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *adOsc_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 2 )

	a.fiddleValues[0] = float64( 3.0000000000 )
	a.fiddleValues[1] = float64( 10.0000000000 )
}

func ( a *adOsc_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *adOsc_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *adOsc_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *adOsc_struct ) GetNumFiddleValues() int {
	return 2
}

func ( a *adOsc_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *adOsc_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 2 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *adOsc_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *adOsc_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *adOsc_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataInteger( a.params, 1, int( a.fiddleValues[1] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *adOsc_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func AdOsc() TA_Function {
	var ret adOsc_struct
	helper_getFunctionHandle( "ADOSC", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type obv_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *obv_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *obv_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *obv_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *obv_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 1, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *obv_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *obv_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *obv_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *obv_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *obv_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *obv_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *obv_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Obv() TA_Function {
	var ret obv_struct
	helper_getFunctionHandle( "OBV", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdl2Crows_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdl2Crows_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdl2Crows_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdl2Crows_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdl2Crows_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdl2Crows_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdl2Crows_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdl2Crows_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdl2Crows_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdl2Crows_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdl2Crows_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdl2Crows_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cdl2Crows() TA_Function {
	var ret cdl2Crows_struct
	helper_getFunctionHandle( "CDL2CROWS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdl3BlackCrows_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdl3BlackCrows_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdl3BlackCrows_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdl3BlackCrows_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdl3BlackCrows_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdl3BlackCrows_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdl3BlackCrows_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdl3BlackCrows_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdl3BlackCrows_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdl3BlackCrows_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdl3BlackCrows_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdl3BlackCrows_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cdl3BlackCrows() TA_Function {
	var ret cdl3BlackCrows_struct
	helper_getFunctionHandle( "CDL3BLACKCROWS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdl3Inside_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdl3Inside_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdl3Inside_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdl3Inside_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdl3Inside_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdl3Inside_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdl3Inside_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdl3Inside_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdl3Inside_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdl3Inside_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdl3Inside_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdl3Inside_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cdl3Inside() TA_Function {
	var ret cdl3Inside_struct
	helper_getFunctionHandle( "CDL3INSIDE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdl3LineStrike_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdl3LineStrike_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdl3LineStrike_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdl3LineStrike_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdl3LineStrike_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdl3LineStrike_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdl3LineStrike_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdl3LineStrike_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdl3LineStrike_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdl3LineStrike_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdl3LineStrike_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdl3LineStrike_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cdl3LineStrike() TA_Function {
	var ret cdl3LineStrike_struct
	helper_getFunctionHandle( "CDL3LINESTRIKE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdl3Outside_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdl3Outside_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdl3Outside_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdl3Outside_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdl3Outside_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdl3Outside_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdl3Outside_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdl3Outside_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdl3Outside_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdl3Outside_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdl3Outside_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdl3Outside_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cdl3Outside() TA_Function {
	var ret cdl3Outside_struct
	helper_getFunctionHandle( "CDL3OUTSIDE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdl3StarsInSouth_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdl3StarsInSouth_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdl3StarsInSouth_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdl3StarsInSouth_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdl3StarsInSouth_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdl3StarsInSouth_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdl3StarsInSouth_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdl3StarsInSouth_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdl3StarsInSouth_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdl3StarsInSouth_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdl3StarsInSouth_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdl3StarsInSouth_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cdl3StarsInSouth() TA_Function {
	var ret cdl3StarsInSouth_struct
	helper_getFunctionHandle( "CDL3STARSINSOUTH", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdl3WhiteSoldiers_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdl3WhiteSoldiers_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdl3WhiteSoldiers_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdl3WhiteSoldiers_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdl3WhiteSoldiers_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdl3WhiteSoldiers_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdl3WhiteSoldiers_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdl3WhiteSoldiers_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdl3WhiteSoldiers_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdl3WhiteSoldiers_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdl3WhiteSoldiers_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdl3WhiteSoldiers_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func Cdl3WhiteSoldiers() TA_Function {
	var ret cdl3WhiteSoldiers_struct
	helper_getFunctionHandle( "CDL3WHITESOLDIERS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlAbandonedBaby_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlAbandonedBaby_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 0.3000000000 )
}

func ( a *cdlAbandonedBaby_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlAbandonedBaby_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlAbandonedBaby_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlAbandonedBaby_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cdlAbandonedBaby_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlAbandonedBaby_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlAbandonedBaby_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlAbandonedBaby_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlAbandonedBaby_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlAbandonedBaby_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlAbandonedBaby() TA_Function {
	var ret cdlAbandonedBaby_struct
	helper_getFunctionHandle( "CDLABANDONEDBABY", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlAdvanceBlock_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlAdvanceBlock_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlAdvanceBlock_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlAdvanceBlock_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlAdvanceBlock_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlAdvanceBlock_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlAdvanceBlock_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlAdvanceBlock_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlAdvanceBlock_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlAdvanceBlock_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlAdvanceBlock_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlAdvanceBlock_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlAdvanceBlock() TA_Function {
	var ret cdlAdvanceBlock_struct
	helper_getFunctionHandle( "CDLADVANCEBLOCK", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlBeltHold_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlBeltHold_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlBeltHold_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlBeltHold_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlBeltHold_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlBeltHold_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlBeltHold_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlBeltHold_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlBeltHold_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlBeltHold_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlBeltHold_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlBeltHold_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlBeltHold() TA_Function {
	var ret cdlBeltHold_struct
	helper_getFunctionHandle( "CDLBELTHOLD", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlBreakaway_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlBreakaway_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlBreakaway_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlBreakaway_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlBreakaway_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlBreakaway_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlBreakaway_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlBreakaway_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlBreakaway_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlBreakaway_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlBreakaway_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlBreakaway_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlBreakaway() TA_Function {
	var ret cdlBreakaway_struct
	helper_getFunctionHandle( "CDLBREAKAWAY", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlClosingMarubozu_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlClosingMarubozu_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlClosingMarubozu_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlClosingMarubozu_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlClosingMarubozu_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlClosingMarubozu_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlClosingMarubozu_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlClosingMarubozu_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlClosingMarubozu_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlClosingMarubozu_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlClosingMarubozu_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlClosingMarubozu_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlClosingMarubozu() TA_Function {
	var ret cdlClosingMarubozu_struct
	helper_getFunctionHandle( "CDLCLOSINGMARUBOZU", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlConcealBabysWall_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlConcealBabysWall_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlConcealBabysWall_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlConcealBabysWall_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlConcealBabysWall_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlConcealBabysWall_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlConcealBabysWall_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlConcealBabysWall_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlConcealBabysWall_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlConcealBabysWall_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlConcealBabysWall_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlConcealBabysWall_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlConcealBabysWall() TA_Function {
	var ret cdlConcealBabysWall_struct
	helper_getFunctionHandle( "CDLCONCEALBABYSWALL", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlCounterAttack_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlCounterAttack_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlCounterAttack_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlCounterAttack_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlCounterAttack_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlCounterAttack_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlCounterAttack_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlCounterAttack_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlCounterAttack_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlCounterAttack_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlCounterAttack_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlCounterAttack_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlCounterAttack() TA_Function {
	var ret cdlCounterAttack_struct
	helper_getFunctionHandle( "CDLCOUNTERATTACK", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlDarkCloudCover_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlDarkCloudCover_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 0.5000000000 )
}

func ( a *cdlDarkCloudCover_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlDarkCloudCover_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlDarkCloudCover_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlDarkCloudCover_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cdlDarkCloudCover_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlDarkCloudCover_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlDarkCloudCover_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlDarkCloudCover_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlDarkCloudCover_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlDarkCloudCover_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlDarkCloudCover() TA_Function {
	var ret cdlDarkCloudCover_struct
	helper_getFunctionHandle( "CDLDARKCLOUDCOVER", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlDoji_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlDoji_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlDoji_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlDoji_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlDoji_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlDoji_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlDoji_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlDoji_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlDoji_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlDoji_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlDoji_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlDoji_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlDoji() TA_Function {
	var ret cdlDoji_struct
	helper_getFunctionHandle( "CDLDOJI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlDojiStar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlDojiStar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlDojiStar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlDojiStar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlDojiStar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlDojiStar_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlDojiStar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlDojiStar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlDojiStar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlDojiStar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlDojiStar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlDojiStar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlDojiStar() TA_Function {
	var ret cdlDojiStar_struct
	helper_getFunctionHandle( "CDLDOJISTAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlDragonflyDoji_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlDragonflyDoji_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlDragonflyDoji_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlDragonflyDoji_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlDragonflyDoji_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlDragonflyDoji_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlDragonflyDoji_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlDragonflyDoji_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlDragonflyDoji_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlDragonflyDoji_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlDragonflyDoji_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlDragonflyDoji_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlDragonflyDoji() TA_Function {
	var ret cdlDragonflyDoji_struct
	helper_getFunctionHandle( "CDLDRAGONFLYDOJI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlEngulfing_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlEngulfing_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlEngulfing_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlEngulfing_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlEngulfing_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlEngulfing_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlEngulfing_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlEngulfing_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlEngulfing_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlEngulfing_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlEngulfing_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlEngulfing_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlEngulfing() TA_Function {
	var ret cdlEngulfing_struct
	helper_getFunctionHandle( "CDLENGULFING", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlEveningDojiStar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlEveningDojiStar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 0.3000000000 )
}

func ( a *cdlEveningDojiStar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlEveningDojiStar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlEveningDojiStar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlEveningDojiStar_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cdlEveningDojiStar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlEveningDojiStar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlEveningDojiStar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlEveningDojiStar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlEveningDojiStar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlEveningDojiStar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlEveningDojiStar() TA_Function {
	var ret cdlEveningDojiStar_struct
	helper_getFunctionHandle( "CDLEVENINGDOJISTAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlEveningStar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlEveningStar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 0.3000000000 )
}

func ( a *cdlEveningStar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlEveningStar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlEveningStar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlEveningStar_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cdlEveningStar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlEveningStar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlEveningStar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlEveningStar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlEveningStar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlEveningStar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlEveningStar() TA_Function {
	var ret cdlEveningStar_struct
	helper_getFunctionHandle( "CDLEVENINGSTAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlGapSideSideWhite_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlGapSideSideWhite_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlGapSideSideWhite_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlGapSideSideWhite_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlGapSideSideWhite_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlGapSideSideWhite_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlGapSideSideWhite_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlGapSideSideWhite_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlGapSideSideWhite_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlGapSideSideWhite_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlGapSideSideWhite_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlGapSideSideWhite_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlGapSideSideWhite() TA_Function {
	var ret cdlGapSideSideWhite_struct
	helper_getFunctionHandle( "CDLGAPSIDESIDEWHITE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlGravestoneDoji_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlGravestoneDoji_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlGravestoneDoji_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlGravestoneDoji_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlGravestoneDoji_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlGravestoneDoji_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlGravestoneDoji_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlGravestoneDoji_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlGravestoneDoji_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlGravestoneDoji_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlGravestoneDoji_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlGravestoneDoji_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlGravestoneDoji() TA_Function {
	var ret cdlGravestoneDoji_struct
	helper_getFunctionHandle( "CDLGRAVESTONEDOJI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHammer_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHammer_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHammer_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHammer_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHammer_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHammer_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHammer_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHammer_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHammer_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHammer_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHammer_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHammer_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHammer() TA_Function {
	var ret cdlHammer_struct
	helper_getFunctionHandle( "CDLHAMMER", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHangingMan_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHangingMan_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHangingMan_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHangingMan_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHangingMan_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHangingMan_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHangingMan_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHangingMan_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHangingMan_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHangingMan_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHangingMan_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHangingMan_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHangingMan() TA_Function {
	var ret cdlHangingMan_struct
	helper_getFunctionHandle( "CDLHANGINGMAN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHarami_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHarami_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHarami_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHarami_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHarami_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHarami_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHarami_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHarami_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHarami_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHarami_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHarami_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHarami_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHarami() TA_Function {
	var ret cdlHarami_struct
	helper_getFunctionHandle( "CDLHARAMI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHaramiCross_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHaramiCross_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHaramiCross_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHaramiCross_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHaramiCross_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHaramiCross_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHaramiCross_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHaramiCross_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHaramiCross_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHaramiCross_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHaramiCross_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHaramiCross_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHaramiCross() TA_Function {
	var ret cdlHaramiCross_struct
	helper_getFunctionHandle( "CDLHARAMICROSS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHignWave_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHignWave_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHignWave_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHignWave_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHignWave_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHignWave_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHignWave_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHignWave_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHignWave_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHignWave_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHignWave_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHignWave_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHignWave() TA_Function {
	var ret cdlHignWave_struct
	helper_getFunctionHandle( "CDLHIGHWAVE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHikkake_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHikkake_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHikkake_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHikkake_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHikkake_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHikkake_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHikkake_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHikkake_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHikkake_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHikkake_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHikkake_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHikkake_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHikkake() TA_Function {
	var ret cdlHikkake_struct
	helper_getFunctionHandle( "CDLHIKKAKE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHikkakeMod_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHikkakeMod_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHikkakeMod_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHikkakeMod_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHikkakeMod_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHikkakeMod_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHikkakeMod_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHikkakeMod_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHikkakeMod_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHikkakeMod_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHikkakeMod_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHikkakeMod_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHikkakeMod() TA_Function {
	var ret cdlHikkakeMod_struct
	helper_getFunctionHandle( "CDLHIKKAKEMOD", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlHomingPigeon_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlHomingPigeon_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlHomingPigeon_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlHomingPigeon_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlHomingPigeon_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlHomingPigeon_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlHomingPigeon_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlHomingPigeon_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlHomingPigeon_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlHomingPigeon_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlHomingPigeon_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlHomingPigeon_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlHomingPigeon() TA_Function {
	var ret cdlHomingPigeon_struct
	helper_getFunctionHandle( "CDLHOMINGPIGEON", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlIdentical3Crows_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlIdentical3Crows_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlIdentical3Crows_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlIdentical3Crows_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlIdentical3Crows_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlIdentical3Crows_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlIdentical3Crows_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlIdentical3Crows_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlIdentical3Crows_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlIdentical3Crows_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlIdentical3Crows_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlIdentical3Crows_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlIdentical3Crows() TA_Function {
	var ret cdlIdentical3Crows_struct
	helper_getFunctionHandle( "CDLIDENTICAL3CROWS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlInNeck_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlInNeck_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlInNeck_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlInNeck_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlInNeck_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlInNeck_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlInNeck_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlInNeck_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlInNeck_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlInNeck_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlInNeck_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlInNeck_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlInNeck() TA_Function {
	var ret cdlInNeck_struct
	helper_getFunctionHandle( "CDLINNECK", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlInvertedHammer_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlInvertedHammer_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlInvertedHammer_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlInvertedHammer_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlInvertedHammer_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlInvertedHammer_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlInvertedHammer_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlInvertedHammer_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlInvertedHammer_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlInvertedHammer_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlInvertedHammer_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlInvertedHammer_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlInvertedHammer() TA_Function {
	var ret cdlInvertedHammer_struct
	helper_getFunctionHandle( "CDLINVERTEDHAMMER", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlKicking_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlKicking_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlKicking_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlKicking_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlKicking_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlKicking_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlKicking_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlKicking_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlKicking_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlKicking_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlKicking_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlKicking_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlKicking() TA_Function {
	var ret cdlKicking_struct
	helper_getFunctionHandle( "CDLKICKING", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlKickingByLength_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlKickingByLength_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlKickingByLength_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlKickingByLength_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlKickingByLength_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlKickingByLength_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlKickingByLength_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlKickingByLength_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlKickingByLength_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlKickingByLength_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlKickingByLength_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlKickingByLength_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlKickingByLength() TA_Function {
	var ret cdlKickingByLength_struct
	helper_getFunctionHandle( "CDLKICKINGBYLENGTH", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlLadderBottom_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlLadderBottom_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlLadderBottom_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlLadderBottom_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlLadderBottom_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlLadderBottom_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlLadderBottom_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlLadderBottom_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlLadderBottom_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlLadderBottom_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlLadderBottom_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlLadderBottom_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlLadderBottom() TA_Function {
	var ret cdlLadderBottom_struct
	helper_getFunctionHandle( "CDLLADDERBOTTOM", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlLongLeggedDoji_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlLongLeggedDoji_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlLongLeggedDoji_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlLongLeggedDoji_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlLongLeggedDoji_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlLongLeggedDoji_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlLongLeggedDoji_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlLongLeggedDoji_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlLongLeggedDoji_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlLongLeggedDoji_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlLongLeggedDoji_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlLongLeggedDoji_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlLongLeggedDoji() TA_Function {
	var ret cdlLongLeggedDoji_struct
	helper_getFunctionHandle( "CDLLONGLEGGEDDOJI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlLongLine_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlLongLine_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlLongLine_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlLongLine_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlLongLine_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlLongLine_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlLongLine_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlLongLine_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlLongLine_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlLongLine_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlLongLine_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlLongLine_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlLongLine() TA_Function {
	var ret cdlLongLine_struct
	helper_getFunctionHandle( "CDLLONGLINE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlMarubozu_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlMarubozu_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlMarubozu_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlMarubozu_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlMarubozu_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlMarubozu_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlMarubozu_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlMarubozu_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlMarubozu_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlMarubozu_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlMarubozu_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlMarubozu_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlMarubozu() TA_Function {
	var ret cdlMarubozu_struct
	helper_getFunctionHandle( "CDLMARUBOZU", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlMatchingLow_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlMatchingLow_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlMatchingLow_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlMatchingLow_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlMatchingLow_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlMatchingLow_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlMatchingLow_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlMatchingLow_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlMatchingLow_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlMatchingLow_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlMatchingLow_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlMatchingLow_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlMatchingLow() TA_Function {
	var ret cdlMatchingLow_struct
	helper_getFunctionHandle( "CDLMATCHINGLOW", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlMatHold_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlMatHold_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 0.5000000000 )
}

func ( a *cdlMatHold_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlMatHold_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlMatHold_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlMatHold_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cdlMatHold_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlMatHold_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlMatHold_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlMatHold_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlMatHold_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlMatHold_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlMatHold() TA_Function {
	var ret cdlMatHold_struct
	helper_getFunctionHandle( "CDLMATHOLD", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlMorningDojiStar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlMorningDojiStar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 0.3000000000 )
}

func ( a *cdlMorningDojiStar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlMorningDojiStar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlMorningDojiStar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlMorningDojiStar_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cdlMorningDojiStar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlMorningDojiStar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlMorningDojiStar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlMorningDojiStar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlMorningDojiStar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlMorningDojiStar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlMorningDojiStar() TA_Function {
	var ret cdlMorningDojiStar_struct
	helper_getFunctionHandle( "CDLMORNINGDOJISTAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlMorningStar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlMorningStar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 0.3000000000 )
}

func ( a *cdlMorningStar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlMorningStar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlMorningStar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlMorningStar_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *cdlMorningStar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlMorningStar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlMorningStar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return 0.000000 + ( ( 0.000000 - 0.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlMorningStar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlMorningStar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataReal( a.params, 0, a.fiddleValues[0] )

	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlMorningStar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlMorningStar() TA_Function {
	var ret cdlMorningStar_struct
	helper_getFunctionHandle( "CDLMORNINGSTAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlOnNeck_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlOnNeck_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlOnNeck_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlOnNeck_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlOnNeck_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlOnNeck_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlOnNeck_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlOnNeck_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlOnNeck_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlOnNeck_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlOnNeck_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlOnNeck_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlOnNeck() TA_Function {
	var ret cdlOnNeck_struct
	helper_getFunctionHandle( "CDLONNECK", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlPiercing_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlPiercing_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlPiercing_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlPiercing_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlPiercing_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlPiercing_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlPiercing_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlPiercing_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlPiercing_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlPiercing_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlPiercing_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlPiercing_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlPiercing() TA_Function {
	var ret cdlPiercing_struct
	helper_getFunctionHandle( "CDLPIERCING", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlRickshawMan_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlRickshawMan_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlRickshawMan_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlRickshawMan_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlRickshawMan_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlRickshawMan_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlRickshawMan_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlRickshawMan_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlRickshawMan_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlRickshawMan_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlRickshawMan_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlRickshawMan_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlRickshawMan() TA_Function {
	var ret cdlRickshawMan_struct
	helper_getFunctionHandle( "CDLRICKSHAWMAN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlRiseFall3Methods_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlRiseFall3Methods_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlRiseFall3Methods_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlRiseFall3Methods_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlRiseFall3Methods_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlRiseFall3Methods_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlRiseFall3Methods_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlRiseFall3Methods_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlRiseFall3Methods_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlRiseFall3Methods_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlRiseFall3Methods_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlRiseFall3Methods_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlRiseFall3Methods() TA_Function {
	var ret cdlRiseFall3Methods_struct
	helper_getFunctionHandle( "CDLRISEFALL3METHODS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlSeperatingLines_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlSeperatingLines_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlSeperatingLines_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlSeperatingLines_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlSeperatingLines_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlSeperatingLines_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlSeperatingLines_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlSeperatingLines_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlSeperatingLines_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlSeperatingLines_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlSeperatingLines_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlSeperatingLines_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlSeperatingLines() TA_Function {
	var ret cdlSeperatingLines_struct
	helper_getFunctionHandle( "CDLSEPARATINGLINES", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlShootingStar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlShootingStar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlShootingStar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlShootingStar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlShootingStar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlShootingStar_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlShootingStar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlShootingStar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlShootingStar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlShootingStar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlShootingStar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlShootingStar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlShootingStar() TA_Function {
	var ret cdlShootingStar_struct
	helper_getFunctionHandle( "CDLSHOOTINGSTAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlShortLine_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlShortLine_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlShortLine_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlShortLine_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlShortLine_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlShortLine_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlShortLine_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlShortLine_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlShortLine_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlShortLine_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlShortLine_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlShortLine_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlShortLine() TA_Function {
	var ret cdlShortLine_struct
	helper_getFunctionHandle( "CDLSHORTLINE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlSpinningTop_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlSpinningTop_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlSpinningTop_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlSpinningTop_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlSpinningTop_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlSpinningTop_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlSpinningTop_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlSpinningTop_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlSpinningTop_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlSpinningTop_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlSpinningTop_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlSpinningTop_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlSpinningTop() TA_Function {
	var ret cdlSpinningTop_struct
	helper_getFunctionHandle( "CDLSPINNINGTOP", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlStalledPattern_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlStalledPattern_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlStalledPattern_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlStalledPattern_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlStalledPattern_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlStalledPattern_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlStalledPattern_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlStalledPattern_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlStalledPattern_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlStalledPattern_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlStalledPattern_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlStalledPattern_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlStalledPattern() TA_Function {
	var ret cdlStalledPattern_struct
	helper_getFunctionHandle( "CDLSTALLEDPATTERN", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlStickSandwhich_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlStickSandwhich_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlStickSandwhich_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlStickSandwhich_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlStickSandwhich_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlStickSandwhich_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlStickSandwhich_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlStickSandwhich_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlStickSandwhich_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlStickSandwhich_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlStickSandwhich_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlStickSandwhich_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlStickSandwhich() TA_Function {
	var ret cdlStickSandwhich_struct
	helper_getFunctionHandle( "CDLSTICKSANDWICH", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlTakuri_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlTakuri_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlTakuri_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlTakuri_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlTakuri_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlTakuri_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlTakuri_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlTakuri_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlTakuri_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlTakuri_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlTakuri_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlTakuri_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlTakuri() TA_Function {
	var ret cdlTakuri_struct
	helper_getFunctionHandle( "CDLTAKURI", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlTasukiGap_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlTasukiGap_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlTasukiGap_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlTasukiGap_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlTasukiGap_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlTasukiGap_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlTasukiGap_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlTasukiGap_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlTasukiGap_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlTasukiGap_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlTasukiGap_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlTasukiGap_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlTasukiGap() TA_Function {
	var ret cdlTasukiGap_struct
	helper_getFunctionHandle( "CDLTASUKIGAP", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlThrusting_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlThrusting_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlThrusting_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlThrusting_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlThrusting_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlThrusting_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlThrusting_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlThrusting_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlThrusting_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlThrusting_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlThrusting_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlThrusting_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlThrusting() TA_Function {
	var ret cdlThrusting_struct
	helper_getFunctionHandle( "CDLTHRUSTING", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlTristar_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlTristar_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlTristar_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlTristar_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlTristar_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlTristar_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlTristar_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlTristar_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlTristar_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlTristar_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlTristar_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlTristar_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlTristar() TA_Function {
	var ret cdlTristar_struct
	helper_getFunctionHandle( "CDLTRISTAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlUnique3River_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlUnique3River_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlUnique3River_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlUnique3River_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlUnique3River_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlUnique3River_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlUnique3River_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlUnique3River_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlUnique3River_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlUnique3River_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlUnique3River_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlUnique3River_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlUnique3River() TA_Function {
	var ret cdlUnique3River_struct
	helper_getFunctionHandle( "CDLUNIQUE3RIVER", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlUpsideGap2Crows_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlUpsideGap2Crows_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlUpsideGap2Crows_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlUpsideGap2Crows_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlUpsideGap2Crows_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlUpsideGap2Crows_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlUpsideGap2Crows_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlUpsideGap2Crows_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlUpsideGap2Crows_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlUpsideGap2Crows_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlUpsideGap2Crows_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlUpsideGap2Crows_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlUpsideGap2Crows() TA_Function {
	var ret cdlUpsideGap2Crows_struct
	helper_getFunctionHandle( "CDLUPSIDEGAP2CROWS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type cdlXSideGap3Methods_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *cdlXSideGap3Methods_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *cdlXSideGap3Methods_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *cdlXSideGap3Methods_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *cdlXSideGap3Methods_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *cdlXSideGap3Methods_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *cdlXSideGap3Methods_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *cdlXSideGap3Methods_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *cdlXSideGap3Methods_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *cdlXSideGap3Methods_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *cdlXSideGap3Methods_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.integerOutputByIndex[0] = make( []C.TA_Integer, endIndex - startIndex + 1 )
	helper_setOutputParamIntegerPtr( a.params, 0, &(a.integerOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaIntegerArrayToGoFloat64Array( a.integerOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *cdlXSideGap3Methods_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func CdlXSideGap3Methods() TA_Function {
	var ret cdlXSideGap3Methods_struct
	helper_getFunctionHandle( "CDLXSIDEGAP3METHODS", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type beta_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *beta_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 5.0000000000 )
}

func ( a *beta_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *beta_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
	if index == 1 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *beta_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *beta_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *beta_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *beta_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *beta_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *beta_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *beta_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *beta_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Beta() TA_Function {
	var ret beta_struct
	helper_getFunctionHandle( "BETA", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type correl_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *correl_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 30.0000000000 )
}

func ( a *correl_struct ) GetNumInputs() ( int ) {
	return 2
}

func ( a *correl_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
	if index == 1 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *correl_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *correl_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *correl_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *correl_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *correl_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *correl_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *correl_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 2; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *correl_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Correl() TA_Function {
	var ret correl_struct
	helper_getFunctionHandle( "CORREL", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type linearReg_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *linearReg_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *linearReg_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *linearReg_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *linearReg_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *linearReg_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *linearReg_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *linearReg_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *linearReg_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *linearReg_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *linearReg_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *linearReg_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func LinearReg() TA_Function {
	var ret linearReg_struct
	helper_getFunctionHandle( "LINEARREG", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type linearRegAngle_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *linearRegAngle_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *linearRegAngle_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *linearRegAngle_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *linearRegAngle_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *linearRegAngle_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *linearRegAngle_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *linearRegAngle_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *linearRegAngle_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *linearRegAngle_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *linearRegAngle_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *linearRegAngle_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func LinearRegAngle() TA_Function {
	var ret linearRegAngle_struct
	helper_getFunctionHandle( "LINEARREG_ANGLE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type linearRegIntercept_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *linearRegIntercept_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *linearRegIntercept_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *linearRegIntercept_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *linearRegIntercept_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *linearRegIntercept_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *linearRegIntercept_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *linearRegIntercept_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *linearRegIntercept_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *linearRegIntercept_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *linearRegIntercept_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *linearRegIntercept_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func LinearRegIntercept() TA_Function {
	var ret linearRegIntercept_struct
	helper_getFunctionHandle( "LINEARREG_INTERCEPT", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type linearRegSlope_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *linearRegSlope_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *linearRegSlope_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *linearRegSlope_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *linearRegSlope_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *linearRegSlope_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *linearRegSlope_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *linearRegSlope_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *linearRegSlope_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *linearRegSlope_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *linearRegSlope_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *linearRegSlope_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func LinearRegSlope() TA_Function {
	var ret linearRegSlope_struct
	helper_getFunctionHandle( "LINEARREG_SLOPE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type stdDev_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *stdDev_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 2 )

	a.fiddleValues[0] = float64( 5.0000000000 )
	a.fiddleValues[1] = float64( 1.0000000000 )
}

func ( a *stdDev_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *stdDev_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *stdDev_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *stdDev_struct ) GetNumFiddleValues() int {
	return 2
}

func ( a *stdDev_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *stdDev_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 2 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *stdDev_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return -2.000000 + ( ( 2.000000 - -2.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *stdDev_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *stdDev_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataReal( a.params, 1, a.fiddleValues[1] )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *stdDev_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func StdDev() TA_Function {
	var ret stdDev_struct
	helper_getFunctionHandle( "STDDEV", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type tsf_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *tsf_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 1 )

	a.fiddleValues[0] = float64( 14.0000000000 )
}

func ( a *tsf_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *tsf_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *tsf_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *tsf_struct ) GetNumFiddleValues() int {
	return 1
}

func ( a *tsf_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *tsf_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 1 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *tsf_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 4 + ( float64( 200 - 4 ) * inValue ) ) )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *tsf_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *tsf_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *tsf_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Tsf() TA_Function {
	var ret tsf_struct
	helper_getFunctionHandle( "TSF", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type variance_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *variance_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 2 )

	a.fiddleValues[0] = float64( 5.0000000000 )
	a.fiddleValues[1] = float64( 1.0000000000 )
}

func ( a *variance_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *variance_struct ) SetInputData( index int, data []float64 ) {
	if index == 0 {
		var temp []C.TA_Real
		helper_convertGoFloat64ArrayToTaRealArray( data, &temp )
		a.realInputByIndex[index] = temp
		helper_setInputDataReal( a.params, index, &(a.realInputByIndex[index][0]) )
		return
	}
}

func ( a *variance_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
}

func ( a *variance_struct ) GetNumFiddleValues() int {
	return 2
}

func ( a *variance_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *variance_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 2 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *variance_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	if fiddleValueIndex == 0 {
		return float64( int( 1 + ( float64( 200 - 1 ) * inValue ) ) )
	}

	if fiddleValueIndex == 1 {
		return -2.000000 + ( ( 2.000000 - -2.000000 ) * inValue )
	}

	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *variance_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *variance_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1

	helper_setOptInputDataInteger( a.params, 0, int( a.fiddleValues[0] ) )
	helper_setOptInputDataReal( a.params, 1, a.fiddleValues[1] )

	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *variance_struct ) GoSingle( outputIndex int ) ( float64 ) {
	a.fiddleValues[0] = float64( len( a.realInputByIndex[0] ) )
	ret := a.Go( outputIndex )
	if len( ret ) > 0 {
		ret := ret[0]
		if math.IsNaN( ret ) || math.IsInf( ret, 0 ) {
			return 0
		} else {
			return ret
		}
	} else {
		return 0
	}
}

func Variance() TA_Function {
	var ret variance_struct
	helper_getFunctionHandle( "VAR", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type avgPrice_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *avgPrice_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *avgPrice_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *avgPrice_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *avgPrice_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *avgPrice_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *avgPrice_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *avgPrice_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *avgPrice_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *avgPrice_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *avgPrice_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *avgPrice_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func AvgPrice() TA_Function {
	var ret avgPrice_struct
	helper_getFunctionHandle( "AVGPRICE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type medPrice_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *medPrice_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *medPrice_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *medPrice_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *medPrice_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *medPrice_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *medPrice_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *medPrice_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *medPrice_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *medPrice_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *medPrice_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *medPrice_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func MedPrice() TA_Function {
	var ret medPrice_struct
	helper_getFunctionHandle( "MEDPRICE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type typPrice_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *typPrice_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *typPrice_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *typPrice_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *typPrice_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *typPrice_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *typPrice_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *typPrice_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *typPrice_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *typPrice_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *typPrice_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *typPrice_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func TypPrice() TA_Function {
	var ret typPrice_struct
	helper_getFunctionHandle( "TYPPRICE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

type wclPrice_struct struct {
	params *C.TA_ParamHolder
	handle *C.TA_FuncHandle

	realInputByIndex map[int][]C.TA_Real
	integerOutputByIndex map[int][]C.TA_Integer
	realOutputByIndex map[int][]C.TA_Real

	fiddleValues []float64
}

func ( a *wclPrice_struct ) init() {
	helper_paramHolderAlloc( a.handle, &a.params )
	a.realInputByIndex = make( map[int][]C.TA_Real )
	a.realOutputByIndex = make( map[int][]C.TA_Real )
	a.integerOutputByIndex = make( map[int][]C.TA_Integer )
	a.fiddleValues = make( []float64, 0 )

}

func ( a *wclPrice_struct ) GetNumInputs() ( int ) {
	return 1
}

func ( a *wclPrice_struct ) SetInputData( index int, data []float64 ) {
}

func ( a *wclPrice_struct ) SetPriceInputData( open, high, low, close, volume, openInterest []float64 ) {
	var temp []C.TA_Real
	helper_convertGoFloat64ArrayToTaRealArray( open, &temp )
	a.realInputByIndex[0] = temp
	helper_convertGoFloat64ArrayToTaRealArray( high, &temp )
	a.realInputByIndex[1] = temp
	helper_convertGoFloat64ArrayToTaRealArray( low, &temp )
	a.realInputByIndex[2] = temp
	helper_convertGoFloat64ArrayToTaRealArray( close, &temp )
	a.realInputByIndex[3] = temp
	helper_convertGoFloat64ArrayToTaRealArray( volume, &temp )
	a.realInputByIndex[4] = temp
	helper_convertGoFloat64ArrayToTaRealArray( openInterest, &temp )
	a.realInputByIndex[5] = temp
	helper_setInputDataPrice( 
		a.params, 0, 
		&( a.realInputByIndex[0][0] ), 
		&( a.realInputByIndex[1][0] ), 
		&( a.realInputByIndex[2][0] ), 
		&( a.realInputByIndex[3][0] ), 
		&( a.realInputByIndex[4][0] ), 
		&( a.realInputByIndex[5][0] ),
	)
}

func ( a *wclPrice_struct ) GetNumFiddleValues() int {
	return 0
}

func ( a *wclPrice_struct ) GetFiddleValues() ( []float64 ) {
	return a.fiddleValues
}

func ( a *wclPrice_struct ) SetFiddleValues( v []float64 ) {
	if len( v ) != 0 {
		panic( "SetFiddleValues : bad number of fiddle values passed" )
	}
	a.fiddleValues = v
}

func (a *wclPrice_struct ) FixFiddleValue( fiddleValueIndex int, inValue float64 ) ( float64 ) {
	panic("invalid index passed to FixFiddleValue")
	return 0.0
}

func (a *wclPrice_struct ) GetNumOutputValues() ( int ) {
	return 1
}

func ( a *wclPrice_struct ) Go( outIndex int ) ( []float64 ) {
	startIndex := 0
	for i := 1; i < 1; i++ {
		if len( a.realInputByIndex[i] ) != len( a.realInputByIndex[i - 1] ) {
			panic("Input data has different lengths")
		}
	}
	endIndex := len( a.realInputByIndex[0] ) - 1


	a.realOutputByIndex[0] = make( []C.TA_Real, endIndex - startIndex + 1 )
	helper_setOutputParamRealPtr( a.params, 0, &(a.realOutputByIndex[0][0] ) )

	var temp, numElements C.TA_Integer
	helper_callFunction( a.params, startIndex, endIndex, &temp, &numElements )

	if outIndex == 0 {
		var ret []float64
		helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[outIndex], &ret )
		return ret[:numElements]
	}

	var ret []float64
	helper_convertTaRealArrayToGoFloat64Array( a.realOutputByIndex[0], &ret )
	return ret[:numElements]
}

func ( a *wclPrice_struct ) GoSingle( outputIndex int ) ( float64 ) {
	return 0
}

func WclPrice() TA_Function {
	var ret wclPrice_struct
	helper_getFunctionHandle( "WCLPRICE", &ret.handle )
	ret.init()
	return &ret
}

/*--------------------------------------------------------------------------------------------------------*/

