package main

//go:generate sh -c "mkdir -p gen/slstr && sed 's,interface{},string,g' ./slide.go > gen/slstr/slide.go"
//go:generate sh -c "mkdir -p gen/slpstr && sed 's,interface{},*string,g' ./slide.go > gen/slpstr/slide.go"

//go:generate sh -c "mkdir -p gen/slint && sed 's,interface{},int,g' ./slide.go > gen/slint/slide.go"
//go:generate sh -c "mkdir -p gen/slpint && sed 's,interface{},*int,g' ./slide.go > gen/slpint/slide.go"
//go:generate sh -c "mkdir -p gen/slint8 && sed 's,interface{},int8,g' ./slide.go > gen/slint8/slide.go"
//go:generate sh -c "mkdir -p gen/slpint8 && sed 's,interface{},*int8,g' ./slide.go > gen/slpint8/slide.go"
//go:generate sh -c "mkdir -p gen/slint16 && sed 's,interface{},int16,g' ./slide.go > gen/slint16/slide.go"
//go:generate sh -c "mkdir -p gen/slpint16 && sed 's,interface{},*int16,g' ./slide.go > gen/slpint16/slide.go"
//go:generate sh -c "mkdir -p gen/slint32 && sed 's,interface{},int32,g' ./slide.go > gen/slint32/slide.go"
//go:generate sh -c "mkdir -p gen/slpint32 && sed 's,interface{},*int32,g' ./slide.go > gen/slpint32/slide.go"
//go:generate sh -c "mkdir -p gen/slint64 && sed 's,interface{},int64,g' ./slide.go > gen/slint64/slide.go"
//go:generate sh -c "mkdir -p gen/slpint64 && sed 's,interface{},*int64,g' ./slide.go > gen/slpint64/slide.go"
//go:generate sh -c "mkdir -p gen/slrune && sed 's,interface{},rune,g' ./slide.go > gen/slrune/slide.go"
//go:generate sh -c "mkdir -p gen/slprune && sed 's,interface{},*rune,g' ./slide.go > gen/slprune/slide.go"

//go:generate sh -c "mkdir -p gen/sluint && sed 's,interface{},uint,g' ./slide.go > gen/sluint/slide.go"
//go:generate sh -c "mkdir -p gen/slpuint && sed 's,interface{},*uint,g' ./slide.go > gen/slpuint/slide.go"
//go:generate sh -c "mkdir -p gen/sluint8 && sed 's,interface{},uint8,g' ./slide.go > gen/sluint8/slide.go"
//go:generate sh -c "mkdir -p gen/slpuint8 && sed 's,interface{},*uint8,g' ./slide.go > gen/slpuint8/slide.go"
//go:generate sh -c "mkdir -p gen/sluint16 && sed 's,interface{},uint16,g' ./slide.go > gen/sluint16/slide.go"
//go:generate sh -c "mkdir -p gen/slpuint16 && sed 's,interface{},*uint16,g' ./slide.go > gen/slpuint16/slide.go"
//go:generate sh -c "mkdir -p gen/sluint32 && sed 's,interface{},uint32,g' ./slide.go > gen/sluint32/slide.go"
//go:generate sh -c "mkdir -p gen/slpuint32 && sed 's,interface{},*uint32,g' ./slide.go > gen/slpuint32/slide.go"
//go:generate sh -c "mkdir -p gen/sluint64 && sed 's,interface{},uint64,g' ./slide.go > gen/sluint64/slide.go"
//go:generate sh -c "mkdir -p gen/slpuint64 && sed 's,interface{},*uint64,g' ./slide.go > gen/slpuint64/slide.go"

//go:generate sh -c "mkdir -p gen/slbyte && sed 's,interface{},byte,g' ./slide.go > gen/slbyte/slide.go"
//go:generate sh -c "mkdir -p gen/slpbyte && sed 's,interface{},*byte,g' ./slide.go > gen/slpbyte/slide.go"

//go:generate sh -c "mkdir -p gen/slfloat32 && sed 's,interface{},float32,g' ./slide.go > gen/slfloat32/slide.go"
//go:generate sh -c "mkdir -p gen/slpfloat32 && sed 's,interface{},*float32,g' ./slide.go > gen/slpfloat32/slide.go"
//go:generate sh -c "mkdir -p gen/slfloat64 && sed 's,interface{},float64,g' ./slide.go > gen/slfloat64/slide.go"
//go:generate sh -c "mkdir -p gen/slpfloat64 && sed 's,interface{},*float64,g' ./slide.go > gen/slpfloat64/slide.go"

//go:generate sh -c "mkdir -p gen/slcomplex64 && sed 's,interface{},complex64,g' ./slide.go > gen/slcomplex64/slide.go"
//go:generate sh -c "mkdir -p gen/slpcomplex64 && sed 's,interface{},*complex64,g' ./slide.go > gen/slpcomplex64/slide.go"
//go:generate sh -c "mkdir -p gen/slcomplex128 && sed 's,interface{},complex128,g' ./slide.go > gen/slcomplex128/slide.go"
//go:generate sh -c "mkdir -p gen/slpcomplex128 && sed 's,interface{},*complex128,g' ./slide.go > gen/slpcomplex128/slide.go"

//go:generate sh -c "mkdir -p gen/slbool && sed 's,interface{},bool,g' ./slide.go > gen/slbool/slide.go"
//go:generate sh -c "mkdir -p gen/slpbool && sed 's,interface{},*bool,g' ./slide.go > gen/slpbool/slide.go"
