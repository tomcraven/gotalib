package gotalib

/*
#cgo LDFLAGS: -L/home/tom/dev/ta-lib/src/.libs/ -lta_lib -lm
#include "/home/tom/dev/ta-lib/include/ta_abstract.h"
*/
import "C"

func helper_paramHolderAlloc(handle *C.TA_FuncHandle, params **C.TA_ParamHolder) {
	if C.TA_ParamHolderAlloc(handle, params) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_setOutputParamRealPtr(params *C.TA_ParamHolder, index int, value *C.TA_Real) {
	if C.TA_SetOutputParamRealPtr(params, 0, value) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_setInputDataInteger(params *C.TA_ParamHolder, paramIndex int, value *C.TA_Integer) {
	if C.TA_SetInputParamIntegerPtr(params, C.uint(paramIndex), value) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_setOutputParamIntegerPtr(params *C.TA_ParamHolder, paramIndex int, value *C.TA_Integer) {
	if C.TA_SetOutputParamIntegerPtr(params, C.uint(paramIndex), value) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_setInputDataReal(params *C.TA_ParamHolder, paramIndex int, value *C.TA_Real) {
	if C.TA_SetInputParamRealPtr(params, C.uint(paramIndex), value) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_setOptInputDataInteger(params *C.TA_ParamHolder, paramIndex int, value int) {
	if ret := C.TA_SetOptInputParamInteger(params, C.uint(paramIndex), C.TA_Integer(value)); ret != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_setOptInputDataReal(params *C.TA_ParamHolder, paramIndex int, value float64) {
	if C.TA_SetOptInputParamReal(params, C.uint(paramIndex), C.TA_Real(value)) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_setInputDataPrice(params *C.TA_ParamHolder, paramIndex int, open, high, close, low, volume, openInterest *C.TA_Real) {
	if C.TA_SetInputParamPricePtr(params, C.uint(paramIndex), open, high, low, close, volume, openInterest) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_callFunction(params *C.TA_ParamHolder, startIndex, endIndex int, out1, out2 *C.TA_Integer) {
	if C.TA_CallFunc(
		params,
		C.TA_Integer(startIndex),
		C.TA_Integer(endIndex),
		out1, out2) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_getFunctionHandle(name string, handle **C.TA_FuncHandle) {
	if C.TA_GetFuncHandle(C.CString(name), handle) != C.TA_SUCCESS {
		panic("doh")
	}
}

func helper_convertGoIntArrayToTaIntegerArray(inData []int, outData *[]C.TA_Integer) {
	dataLength := len(inData)
	*outData = make([]C.TA_Integer, dataLength)
	for i := 0; i < dataLength; i++ {
		(*outData)[i] = C.TA_Integer(inData[i])
	}
}

func helper_convertGoFloat64ArrayToTaRealArray(inData []float64, outData *[]C.TA_Real) {
	dataLength := len(inData)
	*outData = make([]C.TA_Real, dataLength)
	for i := 0; i < dataLength; i++ {
		(*outData)[i] = C.TA_Real(inData[i])
	}
}

func helper_convertTaIntegerArrayToGoFloat64Array(inData []C.TA_Integer, outData *[]float64) {
	dataLength := len(inData)
	*outData = make([]float64, dataLength)
	for i := 0; i < dataLength; i++ {
		(*outData)[i] = float64(inData[i])
	}
}

func helper_convertTaRealArrayToGoFloat64Array(inData []C.TA_Real, outData *[]float64) {
	dataLength := len(inData)
	*outData = make([]float64, dataLength)
	for i := 0; i < dataLength; i++ {
		(*outData)[i] = float64(inData[i])
	}
}
