// Dafny program main.dfy compiled into Go
package main

import (
	_System "System_"
	_dafny "dafny"
	os "os"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "_module.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Abs(x _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return (_dafny.IntOfInt64(-1)).Times(x)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeIndex(x _dafny.Int, length _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return _dafny.Zero
	} else if (x).Cmp(length) >= 0 {
		return (x).Modulo(length)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeDivisionInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).DivBy(x2)
	}
}
func (_static *CompanionStruct_Default___) SafeModuloInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).Modulo(x2)
	}
}
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) bool {
	return ((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(true, true, false))
		}
		return _dafny.MultiSetOf(true)
	})()).IsProperSubsetOf((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(true, false, !(true))
		}
		return _dafny.MultiSetOf(!(!(true)), false)
	})())
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-606)), _dafny.Zero)
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eh"), _dafny.UnicodeSeqOfUtf8Bytes("rm")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qlxxw"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(992))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))))
}
func (_static *CompanionStruct_Default___) Fm4(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("sh"), (func() _dafny.Sequence {
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("siqr")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("ovpwyj")
	})(), _dafny.UnicodeSeqOfUtf8Bytes("amlqk"), _dafny.UnicodeSeqOfUtf8Bytes("reultjxs"))
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 bool, p2 D0, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(119), true)).Cardinality(), _dafny.IntOfInt64(921), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(836)))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(836)))).Contains(_1_v0) {
				_coll0.Add((_1_v0).Times(_dafny.IntOfInt64(248)))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality(), false)).Cardinality())).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(839))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(42))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality()), (_dafny.SetOf(false, false, !(false), false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("iffq"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(609))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})), _dafny.UnicodeSeqOfUtf8Bytes("rxl")))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Map {
	var _source0 D5 = Companion_D5_.Create_DC10_(_dafny.MultiSetOf(false))
	_ = _source0
	if _source0.Is_DC11() {
		var _4___mcc_h0 *C0 = _source0.Get_().(D5_DC11).Cf17
		_ = _4___mcc_h0
		var _5_cf17 *C0 = _4___mcc_h0
		_ = _5_cf17
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("csjiecvhl"), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), (func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-39), _dafny.IntOfInt64(780))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _6_v0 _dafny.Int
				_6_v0 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(-39)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(780)) < 0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_6_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nshammh")).Cardinality())), true)
				}
			}
			return _coll1.ToMap()
		}()).Cardinality())).Cardinality())).Cardinality())).Cardinality()))).Cardinality())).Cardinality(), Companion_D4_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(564))))
	} else if _source0.Is_DC12() {
		var _7___mcc_h1 _dafny.Int = _source0.Get_().(D5_DC12).Cf18
		_ = _7___mcc_h1
		var _8___mcc_h2 _dafny.CodePoint = _source0.Get_().(D5_DC12).Cf19
		_ = _8___mcc_h2
		var _9___mcc_h3 bool = _source0.Get_().(D5_DC12).Cf20
		_ = _9___mcc_h3
		var _10___mcc_h4 _dafny.Sequence = _source0.Get_().(D5_DC12).Cf21
		_ = _10___mcc_h4
		var _11___mcc_h5 _dafny.Int = _source0.Get_().(D5_DC12).Cf22
		_ = _11___mcc_h5
		var _12_cf22 _dafny.Int = _11___mcc_h5
		_ = _12_cf22
		var _13_cf21 _dafny.Sequence = _10___mcc_h4
		_ = _13_cf21
		var _14_cf20 bool = _9___mcc_h3
		_ = _14_cf20
		var _15_cf19 _dafny.CodePoint = _8___mcc_h2
		_ = _15_cf19
		var _16_cf18 _dafny.Int = _7___mcc_h1
		_ = _16_cf18
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_12_cf22, Companion_D4_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_cf20, _16_cf18)))
	} else {
		var _17___mcc_h6 _dafny.MultiSet = _source0.Get_().(D5_DC10).Cf16
		_ = _17___mcc_h6
		var _18_cf16 _dafny.MultiSet = _17___mcc_h6
		_ = _18_cf16
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-694), Companion_D4_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(false)).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality(), Companion_D4_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-749)))))
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(755))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_19_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_dafny.IntOfInt64(-863))
			}))
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-749))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_20_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uhedjju")).Cardinality())
		}))
	})(), (func() _dafny.Sequence {
		if !(true) {
			return _dafny.SeqOf(_dafny.IntOfInt64(227))
		}
		return _dafny.SeqOf((_dafny.SetOf(false, false, false, !(false))).Cardinality(), _dafny.IntOfInt64(215))
	})())
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(221)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(6))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_21_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.MultiSetOf(true))).Cardinality()
	})))).Cardinality()), !((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dwneg")).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sjycrdgy")).Cardinality())) >= 0))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), !(true)), _dafny.SeqOf(!(false))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(!(false))))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) _dafny.CodePoint {
	if false {
		if true {
			return _dafny.CodePoint('k')
		} else {
			return _dafny.CodePoint('t')
		}
	} else if false {
		return _dafny.CodePoint('l')
	} else {
		return _dafny.CodePoint('h')
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (Companion_D7_.Create_DC14_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("qsiewudbc"), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(926), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-434))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality()), _dafny.IntOfInt64(-827), _dafny.IntOfInt64(851), _dafny.IntOfInt64(812))).Cardinality()), _dafny.IntOfInt64(380))))).Dtor_cf24()
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.MultiSet, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var r2 bool = false
	_ = r2
	var _23_v0 bool
	_ = _23_v0
	_23_v0 = true
	var _24_v1 _dafny.Sequence
	_ = _24_v1
	_24_v1 = _dafny.UnicodeSeqOfUtf8Bytes("wtfyrctyq")
	var _25_v2 _dafny.Sequence
	_ = _25_v2
	_25_v2 = _dafny.UnicodeSeqOfUtf8Bytes("plyeo")
	var _source1 _dafny.Sequence = (func() _dafny.Sequence {
		if _23_v0 {
			return _24_v1
		}
		return _25_v2
	})()
	_ = _source1
	var _26___mcc_h0 _dafny.Sequence = _source1
	_ = _26___mcc_h0
	var _27_cf4 _dafny.Sequence = _26___mcc_h0
	_ = _27_cf4
	var _28_v3 _dafny.CodePoint
	_ = _28_v3
	_28_v3 = _dafny.CodePoint('m')
	var _29_v4 _dafny.MultiSet
	_ = _29_v4
	_29_v4 = _dafny.MultiSetOf(_28_v3, _28_v3, _28_v3)
	var _30_v5 _dafny.Map
	_ = _30_v5
	_30_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_29_v4).Cardinality(), _23_v0)
	var _31_v6 _dafny.Int
	_ = _31_v6
	_31_v6 = _dafny.IntOfInt64(847)
	_30_v5 = (_30_v5).Update(_31_v6, (_31_v6).Cmp(_dafny.IntOfInt64(-326)) >= 0)
	if _23_v0 {
		var _32_v7 _dafny.Sequence
		_ = _32_v7
		_32_v7 = _dafny.SeqOf(_27_cf4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(180))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_33_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_34_i0 _dafny.Int) _dafny.CodePoint {
				return _33_v3
			}
		})(_28_v3))))
		var _35_v8 _dafny.Array
		_ = _35_v8
		var _nwElement0_0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_32_v7, _32_v7)
		_ = _nwElement0_0
		var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(6))
		_ = _nw0
		(_nw0).ArraySet1(_nwElement0_0, 0)
		(_nw0).ArraySet1(_32_v7, 1)
		(_nw0).ArraySet1(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("tsajg")), 2)
		(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_32_v7, _32_v7), 3)
		(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_32_v7, _dafny.SeqOf(_24_v1, _27_cf4)), 4)
		(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_32_v7, _32_v7), 5)
		_35_v8 = _nw0
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_35_v8), 0))
		_ = _index0
		(_35_v8).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-648))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_36_cf4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_37_i1 _dafny.Int) _dafny.Sequence {
				return _36_cf4
			}
		})(_27_cf4))), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_35_v8), 0))
		_ = _index1
		(_35_v8).ArraySet1(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_38_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_39_i2 _dafny.Int) _dafny.CodePoint {
				return _38_v3
			}
		})(_28_v3)))), (_index1).Int())
		var _40_v9 _dafny.Array
		_ = _40_v9
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = (func(_41_v6 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_42_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_41_v6, _41_v6), _dafny.SeqOf(_41_v6))
				}
			})(_31_v6)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw1 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw1).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw1).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_40_v9 = _nw1
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_40_v9), 0))
		_ = _index2
		(_40_v9).ArraySet1(Companion_Default___.Fm8(true, _29_v4, globalState), (_index2).Int())
		var _43_v10 _dafny.Sequence
		_ = _43_v10
		_43_v10 = _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_31_v6)))
		var _44_v11 _dafny.Sequence
		_ = _44_v11
		_44_v11 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_43_v10, Companion_Default___.Fm8(_23_v0, _29_v4, globalState)))
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_40_v9), 0))
		_ = _index3
		(_40_v9).ArraySet1((_44_v11).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_31_v6, _31_v6), _dafny.IntOfUint32((_44_v11).Cardinality()))).Uint32()).(_dafny.Sequence), (_index3).Int())
		var _45_v12 _dafny.Map
		_ = _45_v12
		_45_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(842), _30_v5)
		var _46_v13 _dafny.Sequence
		_ = _46_v13
		_46_v13 = _dafny.SeqOf((func() _dafny.Map {
			if (_45_v12).Contains((_dafny.Zero).Minus(_31_v6)) {
				return (_45_v12).Get((_dafny.Zero).Minus(_31_v6)).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v6, _23_v0)
		})(), _30_v5)
		_46_v13 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_30_v5).Update(_31_v6, _23_v0), Companion_Default___.Fm9(_23_v0, _23_v0, _23_v0, _23_v0, globalState), _30_v5, _30_v5), _46_v13)
		var _47_v14 _dafny.Array
		_ = _47_v14
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_1
		var _nw2 _dafny.Array
		_ = _nw2
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw2 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_48_v0 bool) func(_dafny.Int) bool {
				return func(_49_i4 _dafny.Int) bool {
					return _48_v0
				}
			})(_23_v0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw2 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw2).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw2).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_47_v14 = _nw2
		var _50_v15 _dafny.Map
		_ = _50_v15
		_50_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v0, _47_v14)
		var _51_v16 _dafny.Map
		_ = _51_v16
		_51_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_31_v6).Minus(_31_v6), _31_v6)
		var _52_v17 _dafny.Array
		_ = _52_v17
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_2
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_53_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_54_i5 _dafny.Int) _dafny.Int {
					return (_54_i5).Times(_53_v6)
				}
			})(_31_v6)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw3 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw3).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw3).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_52_v17 = _nw3
		var _rhs0 _dafny.Map = (_50_v15).Merge((_50_v15).Merge(_50_v15))
		_ = _rhs0
		var _rhs1 _dafny.Int = (_31_v6).Plus(_31_v6)
		_ = _rhs1
		var _rhs2 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v6, _31_v6)).Merge(_51_v16)
		_ = _rhs2
		var _rhs3 bool = (_31_v6).Cmp(_31_v6) != 0
		_ = _rhs3
		var _rhs4 _dafny.Array = _52_v17
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		_50_v15 = _rhs0
		_31_v6 = _rhs1
		_51_v16 = _rhs2
		_lhs0.F1 = _rhs3
		_52_v17 = _rhs4
		var _55_v18 _dafny.Sequence
		_ = _55_v18
		_55_v18 = _dafny.SeqOf(_23_v0)
		var _56_v19 _dafny.MultiSet
		_ = _56_v19
		_56_v19 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_55_v18).Cardinality())))
		var _57_v20 _dafny.Map
		_ = _57_v20
		_57_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v0, _43_v10)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_47_v14), 0))
		_ = _index4
		(_47_v14).ArraySet1(Companion_Default___.Fm0((func() _dafny.Int {
			if (_56_v19).Contains((func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v3, (_57_v20).Cardinality())).Keys().Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _58_v21 _dafny.CodePoint
					_58_v21 = interface{}(_compr_2).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v3, (_57_v20).Cardinality())).Contains(_58_v21) {
						_coll2.Add(_58_v21)
					}
				}
				return _coll2.ToSet()
			}()).Cardinality()) {
				return (_56_v19).Multiplicity((func() _dafny.Set {
					var _coll3 = _dafny.NewBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v3, (_57_v20).Cardinality())).Keys().Elements()); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _59_v21 _dafny.CodePoint
						_59_v21 = interface{}(_compr_3).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v3, (_57_v20).Cardinality())).Contains(_59_v21) {
							_coll3.Add(_59_v21)
						}
					}
					return _coll3.ToSet()
				}()).Cardinality())
			}
			return _dafny.IntOfInt64(589)
		})(), globalState), (_index4).Int())
		var _60_v22 _dafny.Map
		_ = _60_v22
		_60_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v3, _52_v17)
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_47_v14), 0))
		_ = _index5
		(_47_v14).ArraySet1(!((func() _dafny.Map {
			if _23_v0 {
				return _60_v22
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v3, _52_v17)
		})()).Equals((_60_v22).Merge(_60_v22)), (_index5).Int())
	} else {
		var _61_v23 _dafny.Map
		_ = _61_v23
		_61_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v0, _27_cf4)
		_61_v23 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(269))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}((func(_62_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_63_i6 _dafny.Int) _dafny.CodePoint {
				return _62_v3
			}
		})(_28_v3))))).Merge(_61_v23)).Merge(_61_v23)
		var _64_v24 *C0
		_ = _64_v24
		var _nw4 *C0 = New_C0_()
		_ = _nw4
		_nw4.Ctor__()
		_64_v24 = _nw4
		var _65_v25 _dafny.Array
		_ = _65_v25
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
		_ = _nw5
		_65_v25 = _nw5
		var _66_v26 _dafny.Set
		_ = _66_v26
		_66_v26 = _dafny.SetOf(_31_v6)
		var _67_v27 _dafny.MultiSet
		_ = _67_v27
		_67_v27 = _dafny.MultiSetOf(Companion_Default___.Fm11(_23_v0, _23_v0, (_66_v26).Cardinality(), globalState))
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_65_v25), 0))
		_ = _index6
		(_65_v25).ArraySet1(Companion_Default___.Fm10((_67_v27).Cardinality(), globalState), (_index6).Int())
		var _68_v28 _dafny.Array
		_ = _68_v28
		var _nwElement0_1 _dafny.Sequence = _27_cf4
		_ = _nwElement0_1
		var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(24))
		_ = _nw6
		(_nw6).ArraySet1(_nwElement0_1, 0)
		(_nw6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_24_v1, _dafny.UnicodeSeqOfUtf8Bytes("knesulqj")), 1)
		(_nw6).ArraySet1(_27_cf4, 2)
		(_nw6).ArraySet1(_27_cf4, 3)
		(_nw6).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yxthifwc"), 4)
		(_nw6).ArraySet1((func() _dafny.Sequence {
			if (_61_v23).Contains(true) {
				return (_61_v23).Get(true).(_dafny.Sequence)
			}
			return _24_v1
		})(), 5)
		(_nw6).ArraySet1(_27_cf4, 6)
		(_nw6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_24_v1, _27_cf4), 7)
		(_nw6).ArraySet1(_27_cf4, 8)
		(_nw6).ArraySet1(_24_v1, 9)
		(_nw6).ArraySet1(_27_cf4, 10)
		(_nw6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_24_v1, _dafny.Companion_Sequence_.Update(_24_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_24_v1).Cardinality()), _dafny.IntOfUint32((_24_v1).Cardinality()))).Uint32(), _28_v3)), 11)
		(_nw6).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dqmctbicw"), 12)
		(_nw6).ArraySet1(_24_v1, 13)
		(_nw6).ArraySet1(_24_v1, 14)
		(_nw6).ArraySet1(_27_cf4, 15)
		(_nw6).ArraySet1(_27_cf4, 16)
		(_nw6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_24_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-182))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_69_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_70_i7 _dafny.Int) _dafny.CodePoint {
				return _69_v3
			}
		})(_28_v3)))), 17)
		(_nw6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_27_cf4, _dafny.UnicodeSeqOfUtf8Bytes("gkykgdr")), 18)
		(_nw6).ArraySet1(_24_v1, 19)
		(_nw6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_24_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(18))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_71_v1 _dafny.Sequence, _72_v6 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
			return func(_73_i8 _dafny.Int) _dafny.CodePoint {
				return (_71_v1).Select((Companion_Default___.SafeIndex(_72_v6, _dafny.IntOfUint32((_71_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(_24_v1, _31_v6)))), 20)
		(_nw6).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_27_cf4, _dafny.UnicodeSeqOfUtf8Bytes("hyohbsaxk")), (Companion_Default___.SafeIndex(_31_v6, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_27_cf4, _dafny.UnicodeSeqOfUtf8Bytes("hyohbsaxk"))).Cardinality()))).Uint32(), _28_v3), 21)
		(_nw6).ArraySet1(_24_v1, 22)
		(_nw6).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("krh"), 23)
		_68_v28 = _nw6
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_68_v28), 0))
		_ = _index7
		(_68_v28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_27_cf4, _27_cf4), (_index7).Int())
		var _74_v29 _dafny.Set
		_ = _74_v29
		_74_v29 = _dafny.SetOf(!(true), _23_v0, _23_v0, _23_v0)
		var _75_v30 _dafny.Sequence
		_ = _75_v30
		_75_v30 = _dafny.SeqOf(_23_v0)
		var _76_v31 _dafny.Sequence
		_ = _76_v31
		_76_v31 = _dafny.SeqOf(_75_v30)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_65_v25), 0))
		_ = _index8
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_68_v28), 0))
		_ = _index9
		var _rhs5 bool = Companion_Default___.Fm0((_74_v29).Cardinality(), globalState)
		_ = _rhs5
		var _rhs6 _dafny.Sequence = _76_v31
		_ = _rhs6
		var _rhs7 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_24_v1, _dafny.Companion_Sequence_.Update(_27_cf4, (Companion_Default___.SafeIndex(_31_v6, _dafny.IntOfUint32((_27_cf4).Cardinality()))).Uint32(), (_24_v1).Select((Companion_Default___.SafeIndex(_31_v6, _dafny.IntOfUint32((_24_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)))
		_ = _rhs7
		var _rhs8 _dafny.Int = (_dafny.Zero).Minus(_31_v6)
		_ = _rhs8
		var _rhs9 *C0 = _64_v24
		_ = _rhs9
		var _lhs1 _dafny.Array = _65_v25
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_65_v25), 0))
		_ = _lhs2
		var _lhs3 _dafny.Array = _68_v28
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_68_v28), 0))
		_ = _lhs4
		_23_v0 = _rhs5
		(_lhs1).ArraySet1(_rhs6, (_lhs2).Int())
		(_lhs3).ArraySet1(_rhs7, (_lhs4).Int())
		r0 = _rhs8
		_64_v24 = _rhs9
		_25_v2 = _25_v2
		_24_v1 = _27_cf4
	}
	r1 = Companion_Default___.SafeDivisionInt(_31_v6, _dafny.IntOfUint32((_27_cf4).Cardinality()))
	_31_v6 = _31_v6
	var _77_v32 _dafny.Sequence
	_ = _77_v32
	_77_v32 = _dafny.SeqOf(_23_v0, true, _23_v0)
	var _78_v33 _dafny.Int
	_ = _78_v33
	_78_v33 = _dafny.IntOfInt64(514)
	var _79_i9 _dafny.Int
	_ = _79_i9
	_79_i9 = _dafny.Zero
	{
		for (_77_v32).Select((Companion_Default___.SafeIndex(_78_v33, _dafny.IntOfUint32((_77_v32).Cardinality()))).Uint32()).(bool) {
			{
				if (_79_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_79_i9 = (_79_i9).Plus(_dafny.One)
				var _80_v34 _dafny.CodePoint
				_ = _80_v34
				_80_v34 = _dafny.CodePoint('b')
				var _81_v35 _dafny.Map
				_ = _81_v35
				_81_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v0) == (_23_v0), _25_v2)
				var _82_v36 *C0
				_ = _82_v36
				var _nw7 *C0 = New_C0_()
				_ = _nw7
				_nw7.Ctor__()
				_82_v36 = _nw7
				var _83_v37 _dafny.Sequence
				_ = _83_v37
				_83_v37 = _dafny.SeqOf(_82_v36, _82_v36)
				var _84_v38 _dafny.Map
				_ = _84_v38
				_84_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_83_v37).Select((Companion_Default___.SafeIndex(_78_v33, _dafny.IntOfUint32((_83_v37).Cardinality()))).Uint32()).(*C0), _78_v33)
				var _85_v39 _dafny.MultiSet
				_ = _85_v39
				_85_v39 = _dafny.MultiSetOf(_78_v33, _78_v33, _78_v33)
				var _rhs10 _dafny.Int = (func() _dafny.Int {
					if (_84_v38).Contains(_82_v36) {
						return (_84_v38).Get(_82_v36).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_85_v39).Contains(_dafny.IntOfInt64(925)) {
							return (_85_v39).Multiplicity(_dafny.IntOfInt64(925))
						}
						return _78_v33
					})()
				})()
				_ = _rhs10
				var _rhs11 _dafny.CodePoint = _80_v34
				_ = _rhs11
				var _rhs12 _dafny.Map = _81_v35
				_ = _rhs12
				r1 = _rhs10
				_80_v34 = _rhs11
				_81_v35 = _rhs12
				var _86_v40 _dafny.Map
				_ = _86_v40
				_86_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(867))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg13 _dafny.Int) interface{} {
						return coer13(arg13)
					}
				}((func(_87_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_88_i10 _dafny.Int) _dafny.CodePoint {
						return _87_v34
					}
				})(_80_v34)))).Cardinality()), Companion_Default___.Fm0(_78_v33, globalState))
				var _89_v41 _dafny.Array
				_ = _89_v41
				var _nwElement0_2 bool = _23_v0
				_ = _nwElement0_2
				var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(11))
				_ = _nw8
				(_nw8).ArraySet1(_nwElement0_2, 0)
				(_nw8).ArraySet1(_23_v0, 1)
				(_nw8).ArraySet1(_23_v0, 2)
				(_nw8).ArraySet1(_23_v0, 3)
				(_nw8).ArraySet1(_23_v0, 4)
				(_nw8).ArraySet1((func() bool {
					if (_86_v40).Contains(_78_v33) {
						return (_86_v40).Get(_78_v33).(bool)
					}
					return _23_v0
				})(), 5)
				(_nw8).ArraySet1(!(_23_v0), 6)
				(_nw8).ArraySet1(_23_v0, 7)
				(_nw8).ArraySet1(_23_v0, 8)
				(_nw8).ArraySet1(_23_v0, 9)
				(_nw8).ArraySet1(_23_v0, 10)
				_89_v41 = _nw8
				var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_89_v41), 0))
				_ = _index10
				(_89_v41).ArraySet1((func() bool {
					if _23_v0 {
						return false
					}
					return _23_v0
				})(), (_index10).Int())
				var _90_v42 D5
				_ = _90_v42
				_90_v42 = Companion_D5_.Create_DC12_(_78_v33, Companion_Default___.Fm12(globalState), _23_v0, _24_v1, _78_v33)
				var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_89_v41), 0))
				_ = _index11
				(_89_v41).ArraySet1(!((_23_v0) || ((_90_v42).Dtor_cf20())), (_index11).Int())
				var _rhs13 _dafny.Int = _78_v33
				_ = _rhs13
				var _rhs14 bool = (_89_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_89_v41), 0))).Int()).(bool)
				_ = _rhs14
				var _rhs15 bool = (_89_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_89_v41), 0))).Int()).(bool)
				_ = _rhs15
				var _rhs16 _dafny.Int = (_78_v33).Plus(_78_v33)
				_ = _rhs16
				var _lhs5 *GlobalState = globalState
				_ = _lhs5
				r0 = _rhs13
				_23_v0 = _rhs14
				_lhs5.F1 = _rhs15
				r1 = _rhs16
				var _91_v43 _dafny.MultiSet
				_ = _91_v43
				_91_v43 = _dafny.MultiSetOf(_23_v0)
				_91_v43 = _91_v43
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _92_v44 _dafny.MultiSet
	_ = _92_v44
	_92_v44 = _dafny.MultiSetOf(_23_v0)
	var _93_v45 _dafny.Set
	_ = _93_v45
	_93_v45 = _dafny.SetOf((func() _dafny.Int {
		if (_92_v44).Contains(_23_v0) {
			return (_92_v44).Multiplicity(_23_v0)
		}
		return _78_v33
	})(), _dafny.IntOfUint32((_77_v32).Cardinality()))
	var _94_v46 _dafny.Sequence
	_ = _94_v46
	_94_v46 = _dafny.SeqOf(_93_v45)
	var _95_v47 _dafny.Map
	_ = _95_v47
	_95_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v1, _93_v45)
	var _96_v49 _dafny.Sequence
	_ = _96_v49
	_96_v49 = _dafny.SeqOf(_24_v1, _24_v1, _24_v1)
	var _97_v50 _dafny.CodePoint
	_ = _97_v50
	_97_v50 = _dafny.CodePoint('l')
	var _98_v51 D7
	_ = _98_v51
	_98_v51 = Companion_D7_.Create_DC14_(_95_v47)
	var _99_v54 _dafny.MultiSet
	_ = _99_v54
	_99_v54 = _dafny.MultiSetOf(_24_v1)
	var _100_v55 _dafny.Array
	_ = _100_v55
	var _nwElement0_3 _dafny.Map = _95_v47
	_ = _nwElement0_3
	var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(18))
	_ = _nw9
	(_nw9).ArraySet1(_nwElement0_3, 0)
	(_nw9).ArraySet1(_95_v47, 1)
	(_nw9).ArraySet1(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_96_v49).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _101_v48 _dafny.Sequence
			_101_v48 = interface{}(_compr_4).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_96_v49, _101_v48) {
				_coll4.Add(_101_v48, _93_v45)
			}
		}
		return _coll4.ToMap()
	}(), 2)
	(_nw9).ArraySet1(Companion_Default___.Fm13(_23_v0, false, _24_v1, globalState), 3)
	(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v1, _93_v45), 4)
	(_nw9).ArraySet1(_95_v47, 5)
	(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v1, _93_v45), 6)
	(_nw9).ArraySet1((_95_v47).Merge(_95_v47), 7)
	(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_25_v2), (Companion_Default___.SafeIndex(_78_v33, _dafny.IntOfUint32((_25_v2).Cardinality()))).Uint32(), _97_v50), _93_v45), 8)
	(_nw9).ArraySet1((_95_v47).Merge((_98_v51).Dtor_cf24()), 9)
	(_nw9).ArraySet1((func() _dafny.Map {
		if _23_v0 {
			return _95_v47
		}
		return func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_96_v49).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _102_v52 _dafny.Sequence
				_102_v52 = interface{}(_compr_5).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_96_v49, _102_v52) {
					_coll5.Add(_102_v52, _93_v45)
				}
			}
			return _coll5.ToMap()
		}()
	})(), 10)
	(_nw9).ArraySet1(_95_v47, 11)
	(_nw9).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xrfel"), _dafny.SetOf(_78_v33, _78_v33))).Update(_24_v1, _93_v45), 12)
	(_nw9).ArraySet1(_95_v47, 13)
	(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v1, _93_v45), 14)
	(_nw9).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v1, _93_v45)).Merge(_95_v47), 15)
	(_nw9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v1, _93_v45), 16)
	(_nw9).ArraySet1(func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_99_v54).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _103_v53 _dafny.Sequence
			_103_v53 = interface{}(_compr_6).(_dafny.Sequence)
			if (_99_v54).Contains(_103_v53) {
				_coll6.Add(_103_v53, _93_v45)
			}
		}
		return _coll6.ToMap()
	}(), 17)
	_100_v55 = _nw9
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_100_v55), 0))
	_ = _index12
	(_100_v55).ArraySet1(_95_v47, (_index12).Int())
	var _104_v56 _dafny.Array
	_ = _104_v56
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
	_ = _nw10
	_104_v56 = _nw10
	var _105_v57 _dafny.Array
	_ = _105_v57
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(23)
	_ = _len0_3
	var _nw11 _dafny.Array
	_ = _nw11
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw11 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) D2 = (func(_106_v33 _dafny.Int) func(_dafny.Int) D2 {
			return func(_107_i11 _dafny.Int) D2 {
				return Companion_D2_.Create_DC3_(_dafny.MultiSetOf(_106_v33))
			}
		})(_78_v33)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw11 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw11).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw11).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_105_v57 = _nw11
	var _108_v58 D2
	_ = _108_v58
	_108_v58 = Companion_D2_.Create_DC3_(_dafny.MultiSetOf(_78_v33, _78_v33))
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_105_v57), 0))
	_ = _index13
	(_105_v57).ArraySet1(_108_v58, (_index13).Int())
	var _109_v59 D4
	_ = _109_v59
	_109_v59 = Companion_D4_.Create_DC9_(_104_v56, _dafny.IntOfUint32((_96_v49).Cardinality()), _104_v56)
	var _110_v60 _dafny.MultiSet
	_ = _110_v60
	_110_v60 = _dafny.MultiSetOf(_78_v33, _78_v33)
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_100_v55), 0))
	_ = _index14
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_105_v57), 0))
	_ = _index15
	var _rhs17 _dafny.Sequence = _94_v46
	_ = _rhs17
	var _rhs18 _dafny.Map = _95_v47
	_ = _rhs18
	var _rhs19 _dafny.Array = (func() _dafny.Array {
		if (func() bool {
			if _23_v0 {
				return Companion_Default___.Fm0(_78_v33, globalState)
			}
			return _23_v0
		})() {
			return (_109_v59).Dtor_cf13()
		}
		return _104_v56
	})()
	_ = _rhs19
	var _rhs20 bool = ((_dafny.Zero).Minus((func() _dafny.Int {
		if (_110_v60).Contains(_dafny.IntOfInt64(353)) {
			return (_110_v60).Multiplicity(_dafny.IntOfInt64(353))
		}
		return _78_v33
	})())).Cmp(_78_v33) >= 0
	_ = _rhs20
	var _rhs21 D2 = _108_v58
	_ = _rhs21
	var _lhs6 _dafny.Array = _100_v55
	_ = _lhs6
	var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_100_v55), 0))
	_ = _lhs7
	var _lhs8 _dafny.Array = _105_v57
	_ = _lhs8
	var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_105_v57), 0))
	_ = _lhs9
	_94_v46 = _rhs17
	(_lhs6).ArraySet1(_rhs18, (_lhs7).Int())
	_104_v56 = _rhs19
	r2 = _rhs20
	(_lhs8).ArraySet1(_rhs21, (_lhs9).Int())
	var _111_v61 *C0
	_ = _111_v61
	var _nw12 *C0 = New_C0_()
	_ = _nw12
	_nw12.Ctor__()
	_111_v61 = _nw12
	var _hi0 _dafny.Int = _78_v33
	_ = _hi0
	for _112_i12 := _78_v33; _112_i12.Cmp(_hi0) < 0; _112_i12 = _112_i12.Plus(_dafny.One) {
		var _113_v62 _dafny.Array
		_ = _113_v62
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw13
		_113_v62 = _nw13
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_113_v62), 0))
		_ = _index16
		(_113_v62).ArraySet1(true, (_index16).Int())
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_113_v62), 0))
		_ = _index17
		(_113_v62).ArraySet1(_23_v0, (_index17).Int())
		var _114_v63 _dafny.Array
		_ = _114_v63
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(20))
		_ = _nw14
		_114_v63 = _nw14
		_114_v63 = _114_v63
		var _115_v64 *C0
		_ = _115_v64
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__()
		_115_v64 = _nw15
		(_111_v61).M1(_112_i12, !(_23_v0), globalState)
	}
	var _116_v65 D0
	_ = _116_v65
	_116_v65 = Companion_D0_.Create_DC0_((_78_v33).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(89))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}((func(_117_v33 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_118_i13 _dafny.Int) _dafny.Int {
			return _117_v33
		}
	})(_78_v33)))).Cardinality())) != 0)
	var _pat_let_tv0 = _110_v60
	_ = _pat_let_tv0
	var _pat_let_tv1 = _110_v60
	_ = _pat_let_tv1
	_116_v65 = func(_pat_let0_0 D0) D0 {
		return func(_119_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 bool) D0 {
				return func(_120_dt__update_hcf0_h0 bool) D0 {
					return Companion_D0_.Create_DC0_(_120_dt__update_hcf0_h0)
				}(_pat_let1_0)
			}((_pat_let_tv0).IsSubsetOf(_pat_let_tv1))
		}(_pat_let0_0)
	}(_116_v65)
	r0 = _78_v33
	r1 = _dafny.IntOfInt64(-928)
	r2 = _23_v0
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _121_globalState *GlobalState
	_ = _121_globalState
	var _nw16 *GlobalState = New_GlobalState_()
	_ = _nw16
	_nw16.Ctor__(_dafny.SeqOf(_dafny.IntOfInt64(879)), true)
	_121_globalState = _nw16
	var _122_v0 _dafny.Array
	_ = _122_v0
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_4
	var _nw17 _dafny.Array
	_ = _nw17
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw17 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) bool = func(_123_i1 _dafny.Int) bool {
			return !(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(645), _dafny.IntOfInt64(936), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(901), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_124_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))).Cardinality()), _dafny.IntOfInt64(989), _dafny.IntOfInt64(532), _dafny.IntOfInt64(440))).Cardinality()))).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}(func(_125_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))).Cardinality()))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(872))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}(func(_126_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(467)
			}))).Cardinality())))).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(874)))
		}
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw17 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw17).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw17).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_122_v0 = _nw17
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_122_v0), 0))); ; {
		_guard_loop_0, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _127_i0 _dafny.Int
		_127_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_127_i0).Sign() != -1) && ((_127_i0).Cmp(_dafny.ArrayLen((_122_v0), 0)) < 0)) {
			(_122_v0).ArraySet1((Companion_D0_.Create_DC0_(true)).Dtor_cf0(), (_127_i0).Int())
		}
	}
	var _128_v1 D0
	_ = _128_v1
	_128_v1 = Companion_D0_.Create_DC0_(!(true))
	var _source2 D0 = _128_v1
	_ = _source2
	if _source2.Is_DC1() {
		var _129___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
		_ = _129___mcc_h0
		var _130___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
		_ = _130___mcc_h1
		var _131___mcc_h2 bool = _source2.Get_().(D0_DC1).Cf3
		_ = _131___mcc_h2
		var _132_cf3 bool = _131___mcc_h2
		_ = _132_cf3
		var _133_cf2 _dafny.Int = _130___mcc_h1
		_ = _133_cf2
		var _134_cf1 _dafny.Int = _129___mcc_h0
		_ = _134_cf1
		var _135_v2 _dafny.Sequence
		_ = _135_v2
		_135_v2 = _dafny.UnicodeSeqOfUtf8Bytes("hkd")
		var _136_v3 _dafny.Map
		_ = _136_v3
		_136_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_cf1, _133_cf2)
		var _137_v4 _dafny.Sequence
		_ = _137_v4
		_137_v4 = _dafny.SeqOf(_dafny.IntOfInt64(892), (_136_v3).Cardinality())
		var _138_v5 _dafny.Map
		_ = _138_v5
		_138_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_cf3, _137_v4)
		if !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(308))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}(func(_139_i5 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(488))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_140_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))).Cardinality())
		})), (func() _dafny.Sequence {
			if (_138_v5).Contains(_132_cf3) {
				return (_138_v5).Get(_132_cf3).(_dafny.Sequence)
			}
			return _137_v4
		})()), _dafny.IntOfUint32((_135_v2).Cardinality())) {
			_134_cf1 = Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_cf3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(669))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_141_cf2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_142_i7 _dafny.Int) _dafny.Int {
					return _141_cf2
				}
			})(_133_cf2))))).Cardinality(), _133_cf2)
			var _143_v6 _dafny.MultiSet
			_ = _143_v6
			_143_v6 = _dafny.MultiSetOf((_dafny.IntOfUint32((_135_v2).Cardinality())).Cmp(_134_cf1) <= 0, _132_cf3, _132_cf3, !(_132_cf3) || (_132_cf3), !(_132_cf3) || (Companion_Default___.Fm0(_134_cf1, _121_globalState)))
			_143_v6 = (_143_v6).Intersection(_143_v6)
			var _144_v7 _dafny.Int
			_ = _144_v7
			var _145_v8 _dafny.Int
			_ = _145_v8
			var _146_v9 bool
			_ = _146_v9
			var _out0 _dafny.Int
			_ = _out0
			var _out1 _dafny.Int
			_ = _out1
			var _out2 bool
			_ = _out2
			_out0, _out1, _out2 = Companion_Default___.M0(_dafny.MultiSetOf(_137_v4, _137_v4), _121_globalState)
			_144_v7 = _out0
			_145_v8 = _out1
			_146_v9 = _out2
			var _147_v10 _dafny.Map
			_ = _147_v10
			_147_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v9, !(!(_146_v9) || (!((_128_v1).Dtor_cf0()))))
			var _148_v11 _dafny.Sequence
			_ = _148_v11
			_148_v11 = _dafny.SeqOf(_132_cf3)
			_147_v10 = (_147_v10).Update(_146_v9, (_148_v11).Select((Companion_Default___.SafeIndex(_144_v7, _dafny.IntOfUint32((_148_v11).Cardinality()))).Uint32()).(bool))
			(_121_globalState).F1 = _146_v9
		} else {
			var _149_v12 _dafny.Int
			_ = _149_v12
			var _150_v13 _dafny.Int
			_ = _150_v13
			var _151_v14 bool
			_ = _151_v14
			var _out3 _dafny.Int
			_ = _out3
			var _out4 _dafny.Int
			_ = _out4
			var _out5 bool
			_ = _out5
			_out3, _out4, _out5 = Companion_Default___.M0(_dafny.MultiSetOf(_137_v4, _137_v4), _121_globalState)
			_149_v12 = _out3
			_150_v13 = _out4
			_151_v14 = _out5
			(_121_globalState).F1 = (Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_134_cf1), Companion_Default___.Fm1(_133_cf2, _121_globalState))).Cmp(_149_v12) <= 0
			_135_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ne"), _135_v2), _dafny.Companion_Sequence_.Concatenate(_135_v2, _135_v2))
			(_121_globalState).F1 = false
			var _rhs22 bool = _151_v14
			_ = _rhs22
			var _rhs23 bool = ((_dafny.IntOfInt64(395)).Minus(_134_cf1)).Cmp(_133_cf2) == 0
			_ = _rhs23
			_151_v14 = _rhs22
			_132_cf3 = _rhs23
		}
		if !(!(_132_cf3)) || ((_128_v1).Dtor_cf0()) {
			var _152_v15 _dafny.Array
			_ = _152_v15
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw18
			_152_v15 = _nw18
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
			_ = _nw19
			_152_v15 = _nw19
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_122_v0), 0))
			_ = _index18
			(_122_v0).ArraySet1(Companion_Default___.Fm0(_134_cf1, _121_globalState), (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_122_v0), 0))
			_ = _index19
			(_122_v0).ArraySet1(!(_132_cf3) || (_132_cf3), (_index19).Int())
			(_121_globalState).F0 = _137_v4
			_134_cf1 = Companion_Default___.Fm1((_dafny.IntOfUint32((_137_v4).Cardinality())).Plus(_134_cf1), _121_globalState)
			_133_cf2 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(332))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_153_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()), _133_cf2)
		} else {
			var _154_v16 _dafny.Array
			_ = _154_v16
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_5
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_155_cf1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_156_i9 _dafny.Int) _dafny.Int {
						return (_156_i9).Plus(_155_cf1)
					}
				})(_134_cf1)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw20 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw20).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw20).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_154_v16 = _nw20
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_6
			var _nw21 _dafny.Array
			_ = _nw21
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw21 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = (func(_157_cf1 _dafny.Int, _158_v3 _dafny.Map, _159_cf2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_160_i10 _dafny.Int) _dafny.Int {
						return (_160_i10).Plus((func() _dafny.Int {
							if (_158_v3).Contains(_157_cf1) {
								return (_158_v3).Get(_157_cf1).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_159_cf2)
						})())
					}
				})(_134_cf1, _136_v3, _133_cf2)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw21 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw21).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw21).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_154_v16 = _nw21
			var _161_v17 _dafny.Map
			_ = _161_v17
			_161_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_cf2, _132_cf3)
			var _162_v18 _dafny.Sequence
			_ = _162_v18
			_162_v18 = _dafny.SeqOf(_161_v17, _161_v17)
			var _rhs24 bool = _132_cf3
			_ = _rhs24
			var _rhs25 _dafny.Sequence = _162_v18
			_ = _rhs25
			var _lhs10 *GlobalState = _121_globalState
			_ = _lhs10
			_lhs10.F1 = _rhs24
			_162_v18 = _rhs25
			(_121_globalState).F1 = _132_cf3
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_122_v0), 0))
			_ = _index20
			(_122_v0).ArraySet1(_132_cf3, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_122_v0), 0))
			_ = _index21
			var _rhs26 bool = _132_cf3
			_ = _rhs26
			var _rhs27 _dafny.Int = (_134_cf1).Minus(_dafny.IntOfInt64(653))
			_ = _rhs27
			var _lhs11 _dafny.Array = _122_v0
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_122_v0), 0))
			_ = _lhs12
			(_lhs11).ArraySet1(_rhs26, (_lhs12).Int())
			_133_cf2 = _rhs27
			var _163_v19 _dafny.Map
			_ = _163_v19
			_163_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), _134_cf1)
			var _164_v20 _dafny.MultiSet
			_ = _164_v20
			_164_v20 = _dafny.MultiSetOf(_137_v4)
			var _165_v21 _dafny.Int
			_ = _165_v21
			var _166_v22 _dafny.Int
			_ = _166_v22
			var _167_v23 bool
			_ = _167_v23
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			var _out8 bool
			_ = _out8
			_out6, _out7, _out8 = Companion_Default___.M0((_dafny.MultiSetOf(_dafny.SeqOf(_133_cf2, _134_cf1), _137_v4, _dafny.SeqOf(_133_cf2, (_163_v19).Cardinality()), _137_v4, _137_v4)).Intersection(_164_v20), _121_globalState)
			_165_v21 = _out6
			_166_v22 = _out7
			_167_v23 = _out8
		}
		_133_cf2 = (_dafny.Zero).Minus(_133_cf2)
		var _168_v24 _dafny.Map
		_ = _168_v24
		_168_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_cf3, true)
		_168_v24 = (_168_v24).Update((_134_cf1).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_134_cf1))) > 0, _132_cf3)
	} else {
		var _169___mcc_h3 bool = _source2.Get_().(D0_DC0).Cf0
		_ = _169___mcc_h3
		var _170_cf0 bool = _169___mcc_h3
		_ = _170_cf0
		var _171_v25 _dafny.Int
		_ = _171_v25
		_171_v25 = _dafny.IntOfInt64(136)
		var _172_v26 _dafny.Array
		_ = _172_v26
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_7
		var _nw22 _dafny.Array
		_ = _nw22
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw22 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Int = (func(_173_v25 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_174_i11 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_174_i11, _173_v25)
				}
			})(_171_v25)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw22 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw22).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw22).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_172_v26 = _nw22
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_172_v26), 0))
		_ = _index22
		(_172_v26).ArraySet1(Companion_Default___.SafeDivisionInt(_171_v25, _171_v25), (_index22).Int())
		var _175_v27 _dafny.MultiSet
		_ = _175_v27
		_175_v27 = _dafny.MultiSetOf(_170_cf0)
		var _176_v28 _dafny.Sequence
		_ = _176_v28
		_176_v28 = _dafny.SeqOf(_175_v27)
		var _177_v29 D0
		_ = _177_v29
		_177_v29 = Companion_D0_.Create_DC1_(_171_v25, _171_v25, _170_cf0)
		var _178_v30 _dafny.Map
		_ = _178_v30
		_178_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_170_cf0), _dafny.IntOfInt64(-24))
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_172_v26), 0))
		_ = _index23
		var _rhs28 _dafny.Int = ((func() _dafny.MultiSet {
			if _170_cf0 {
				return (_176_v28).Select((Companion_Default___.SafeIndex(_171_v25, _dafny.IntOfUint32((_176_v28).Cardinality()))).Uint32()).(_dafny.MultiSet)
			}
			return (_dafny.MultiSetOf((_177_v29).Dtor_cf3(), _170_cf0)).Update(Companion_Default___.Fm0((_178_v30).Cardinality(), _121_globalState), Companion_Default___.Abs(_171_v25))
		})()).Cardinality()
		_ = _rhs28
		var _rhs29 _dafny.Int = _171_v25
		_ = _rhs29
		var _rhs30 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_cf0, _170_cf0)).Cardinality()).Minus(_171_v25)
		_ = _rhs30
		var _rhs31 _dafny.Int = (func() _dafny.Int {
			if (_178_v30).Contains(true) {
				return (_178_v30).Get(true).(_dafny.Int)
			}
			return _171_v25
		})()
		_ = _rhs31
		var _rhs32 bool = Companion_Default___.Fm0(_171_v25, _121_globalState)
		_ = _rhs32
		var _lhs13 _dafny.Array = _172_v26
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_172_v26), 0))
		_ = _lhs14
		_171_v25 = _rhs28
		_171_v25 = _rhs29
		_171_v25 = _rhs30
		(_lhs13).ArraySet1(_rhs31, (_lhs14).Int())
		_170_cf0 = _rhs32
		var _179_v31 _dafny.Set
		_ = _179_v31
		_179_v31 = _dafny.SetOf(_178_v30, _178_v30)
		var _180_v32 _dafny.Set
		_ = _180_v32
		_180_v32 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("wop"))
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_122_v0), 0))
		_ = _index24
		(_122_v0).ArraySet1(!(_180_v32).Equals(_180_v32), (_index24).Int())
		var _181_v33 _dafny.Sequence
		_ = _181_v33
		_181_v33 = _dafny.UnicodeSeqOfUtf8Bytes("iqinws")
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_122_v0), 0))
		_ = _index25
		var _rhs33 _dafny.Set = _179_v31
		_ = _rhs33
		var _rhs34 bool = !_dafny.Companion_Sequence_.Equal(_181_v33, _dafny.Companion_Sequence_.Update(_181_v33, (Companion_Default___.SafeIndex(_171_v25, _dafny.IntOfUint32((_181_v33).Cardinality()))).Uint32(), _dafny.CodePoint('p')))
		_ = _rhs34
		var _rhs35 bool = !(_170_cf0) || (((_dafny.Zero).Minus(_171_v25)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_182_i12 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))).Cardinality())) <= 0)
		_ = _rhs35
		var _rhs36 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm1((_dafny.Zero).Minus(_171_v25), _121_globalState))
		_ = _rhs36
		var _rhs37 bool = _170_cf0
		_ = _rhs37
		var _lhs15 *GlobalState = _121_globalState
		_ = _lhs15
		var _lhs16 _dafny.Array = _122_v0
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_122_v0), 0))
		_ = _lhs17
		_179_v31 = _rhs33
		_lhs15.F1 = _rhs34
		_170_cf0 = _rhs35
		_171_v25 = _rhs36
		(_lhs16).ArraySet1(_rhs37, (_lhs17).Int())
		var _source3 D0 = _177_v29
		_ = _source3
		if _source3.Is_DC1() {
			var _183___mcc_h4 _dafny.Int = _source3.Get_().(D0_DC1).Cf1
			_ = _183___mcc_h4
			var _184___mcc_h5 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
			_ = _184___mcc_h5
			var _185___mcc_h6 bool = _source3.Get_().(D0_DC1).Cf3
			_ = _185___mcc_h6
			var _186_cf3 bool = _185___mcc_h6
			_ = _186_cf3
			var _187_cf2 _dafny.Int = _184___mcc_h5
			_ = _187_cf2
			var _188_cf1 _dafny.Int = _183___mcc_h4
			_ = _188_cf1
			var _189_v34 *C0
			_ = _189_v34
			var _nw23 *C0 = New_C0_()
			_ = _nw23
			_nw23.Ctor__()
			_189_v34 = _nw23
			var _190_v35 _dafny.Sequence
			_ = _190_v35
			_190_v35 = _dafny.SeqOf((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool))
			var _191_v36 _dafny.Sequence
			_ = _191_v36
			_191_v36 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), _171_v25)).Cardinality())
			var _192_v37 _dafny.Map
			_ = _192_v37
			_192_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(496), _189_v34)
			var _rhs38 D0 = Companion_D0_.Create_DC0_((_175_v27).Equals(_175_v27))
			_ = _rhs38
			var _rhs39 _dafny.Sequence = _dafny.SeqOf(false, (_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), _170_cf0)
			_ = _rhs39
			var _rhs40 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_191_v36, _191_v36), _dafny.SeqOf(_dafny.IntOfUint32((_190_v35).Cardinality()), _187_cf2, _187_cf2))
			_ = _rhs40
			var _rhs41 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_181_v33).Cardinality()))).Plus(((_192_v37).Merge(_192_v37)).Cardinality())
			_ = _rhs41
			var _rhs42 *C0 = _189_v34
			_ = _rhs42
			var _lhs18 *GlobalState = _121_globalState
			_ = _lhs18
			_128_v1 = _rhs38
			_190_v35 = _rhs39
			_lhs18.F0 = _rhs40
			_187_cf2 = _rhs41
			_189_v34 = _rhs42
			var _193_v38 *C0
			_ = _193_v38
			var _nw24 *C0 = New_C0_()
			_ = _nw24
			_nw24.Ctor__()
			_193_v38 = _nw24
			(_193_v38).M1(_188_cf1, _186_cf3, _121_globalState)
		} else {
			var _194___mcc_h7 bool = _source3.Get_().(D0_DC0).Cf0
			_ = _194___mcc_h7
			var _195_cf0 bool = _194___mcc_h7
			_ = _195_cf0
			var _196_v39 *C0
			_ = _196_v39
			var _nw25 *C0 = New_C0_()
			_ = _nw25
			_nw25.Ctor__()
			_196_v39 = _nw25
			var _197_v40 _dafny.Set
			_ = _197_v40
			_197_v40 = _dafny.SetOf(_196_v39, _196_v39, _196_v39, _196_v39)
			_197_v40 = ((_197_v40).Intersection(_197_v40)).Difference(_dafny.SetOf(_196_v39))
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_172_v26), 0))
			_ = _index26
			(_172_v26).ArraySet1(_171_v25, (_index26).Int())
			var _198_v41 _dafny.Sequence
			_ = _198_v41
			_198_v41 = _dafny.SeqOf((_172_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_172_v26), 0))).Int()).(_dafny.Int))
			var _199_v42 _dafny.MultiSet
			_ = _199_v42
			_199_v42 = _dafny.MultiSetOf(_198_v41)
			var _200_v43 _dafny.Int
			_ = _200_v43
			var _201_v44 _dafny.Int
			_ = _201_v44
			var _202_v45 bool
			_ = _202_v45
			var _out9 _dafny.Int
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 bool
			_ = _out11
			_out9, _out10, _out11 = Companion_Default___.M0(_199_v42, _121_globalState)
			_200_v43 = _out9
			_201_v44 = _out10
			_202_v45 = _out11
			_181_v33 = _181_v33
		}
		var _203_v46 *C0
		_ = _203_v46
		var _nw26 *C0 = New_C0_()
		_ = _nw26
		_nw26.Ctor__()
		_203_v46 = _nw26
		var _204_v47 _dafny.CodePoint
		_ = _204_v47
		_204_v47 = _dafny.CodePoint('j')
		var _205_v48 _dafny.Map
		_ = _205_v48
		_205_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_204_v47, _203_v46)
		_203_v46 = (func() *C0 {
			if (_205_v48).Contains(_204_v47) {
				return (_205_v48).Get(_204_v47).(*C0)
			}
			return _203_v46
		})()
	}
	var _206_v49 *C0
	_ = _206_v49
	var _nw27 *C0 = New_C0_()
	_ = _nw27
	_nw27.Ctor__()
	_206_v49 = _nw27
	var _207_v50 bool
	_ = _207_v50
	_207_v50 = false
	if _207_v50 {
		var _208_v51 _dafny.CodePoint
		_ = _208_v51
		_208_v51 = _dafny.CodePoint('i')
		_208_v51 = _208_v51
		var _209_v52 _dafny.Int
		_ = _209_v52
		_209_v52 = _dafny.IntOfInt64(296)
		var _210_v53 _dafny.Set
		_ = _210_v53
		_210_v53 = _dafny.SetOf((_dafny.Zero).Minus(_209_v52), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_207_v50, _207_v50)).Cardinality())))
		_209_v52 = Companion_Default___.Fm1(((func() _dafny.Set {
			if _207_v50 {
				return _210_v53
			}
			return _210_v53
		})()).Cardinality(), _121_globalState)
		_208_v51 = _208_v51
		var _211_v54 *C0
		_ = _211_v54
		var _nw28 *C0 = New_C0_()
		_ = _nw28
		_nw28.Ctor__()
		_211_v54 = _nw28
		if (Companion_Default___.SafeModuloInt(_209_v52, _209_v52)).Cmp(_209_v52) <= 0 {
			(_206_v49).M1((_209_v52).Plus(_209_v52), _207_v50, _121_globalState)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))
			_ = _index27
			(_122_v0).ArraySet1(_207_v50, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))
			_ = _index28
			(_122_v0).ArraySet1(_207_v50, (_index28).Int())
			(_121_globalState).F1 = !(!((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool)))
			var _212_v55 _dafny.Array
			_ = _212_v55
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_8
			var _nw29 _dafny.Array
			_ = _nw29
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw29 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = func(_213_i13 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("hvo")
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw29 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw29).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw29).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_212_v55 = _nw29
			var _214_v56 _dafny.Sequence
			_ = _214_v56
			_214_v56 = _dafny.UnicodeSeqOfUtf8Bytes("sdpbjuu")
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_212_v55), 0))
			_ = _index29
			(_212_v55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_214_v56, _214_v56), (_index29).Int())
			var _215_v57 D0
			_ = _215_v57
			_215_v57 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(692), _209_v52, (_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool))
			var _216_v58 _dafny.MultiSet
			_ = _216_v58
			_216_v58 = _dafny.MultiSetOf((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), (_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), _207_v50, (_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), !((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool)))
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_212_v55), 0))
			_ = _index30
			(_212_v55).ArraySet1((func() _dafny.Sequence {
				if !((Companion_Default___.Fm5((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), (_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), _215_v57, _121_globalState)).IsDisjointFrom((Companion_D2_.Create_DC3_(_dafny.MultiSetOf(_209_v52, ((_216_v58).Update((_122_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_122_v0), 0))).Int()).(bool), Companion_Default___.Abs(_209_v52))).Cardinality()))).Dtor_cf5())) {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(196))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg23 _dafny.Int) interface{} {
							return coer23(arg23)
						}
					}((func(_217_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_218_i14 _dafny.Int) _dafny.CodePoint {
							return _217_v51
						}
					})(_208_v51)))
				}
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-556))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}((func(_219_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_220_i15 _dafny.Int) _dafny.CodePoint {
						return _219_v51
					}
				})(_208_v51))), _214_v56)
			})(), (_index30).Int())
			var _221_v59 _dafny.Sequence
			_ = _221_v59
			_221_v59 = _dafny.SeqOf(_209_v52, (_216_v58).Cardinality(), (_dafny.Zero).Minus(_209_v52))
			_209_v52 = ((_221_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.IntOfUint32((_221_v59).Cardinality()))).Uint32()).(_dafny.Int)).Times((_209_v52).Times(_209_v52))
		} else {
			var _222_v60 _dafny.Array
			_ = _222_v60
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_9
			var _nw30 _dafny.Array
			_ = _nw30
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw30 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = func(_223_i16 _dafny.Int) _dafny.Int {
					return (_223_i16).Plus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw30 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw30).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw30).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_222_v60 = _nw30
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_222_v60), 0))
			_ = _index31
			(_222_v60).ArraySet1((_dafny.Zero).Minus(_209_v52), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_222_v60), 0))
			_ = _index32
			(_222_v60).ArraySet1(_209_v52, (_index32).Int())
			var _224_v61 _dafny.Map
			_ = _224_v61
			_224_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_208_v51, _209_v52)
			var _225_v62 _dafny.Map
			_ = _225_v62
			_225_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_222_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_222_v60), 0))).Int()).(_dafny.Int), (_222_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_222_v60), 0))).Int()).(_dafny.Int))
			_224_v61 = (_224_v61).Update(_208_v51, Companion_Default___.SafeDivisionInt((_225_v62).Cardinality(), _dafny.IntOfInt64(196)))
			var _226_v63 _dafny.MultiSet
			_ = _226_v63
			_226_v63 = _dafny.MultiSetOf(_208_v51, _dafny.CodePoint('w'), _208_v51, _208_v51, _208_v51)
			var _227_v64 _dafny.Sequence
			_ = _227_v64
			_227_v64 = _dafny.SeqOf(_208_v51, _208_v51, _dafny.CodePoint('a'), _208_v51, _208_v51)
			var _228_v65 _dafny.Sequence
			_ = _228_v65
			_228_v65 = _dafny.SeqOf(_226_v63, _226_v63, (_226_v63).Update(_208_v51, Companion_Default___.Abs(_209_v52)), _dafny.MultiSetFromSeq(_227_v64), _226_v63)
			_226_v63 = (_228_v65).Select((Companion_Default___.SafeIndex(_209_v52, _dafny.IntOfUint32((_228_v65).Cardinality()))).Uint32()).(_dafny.MultiSet)
			_209_v52 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_209_v52, _209_v52), _209_v52)
			(_206_v49).M1((_222_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_222_v60), 0))).Int()).(_dafny.Int), (func() bool {
				if _207_v50 {
					return !(_207_v50)
				}
				return _207_v50
			})(), _121_globalState)
		}
	} else {
		var _229_v66 _dafny.MultiSet
		_ = _229_v66
		_229_v66 = _dafny.MultiSetOf(!(true))
		var _230_v67 _dafny.Int
		_ = _230_v67
		_230_v67 = _dafny.IntOfInt64(399)
		var _rhs43 _dafny.MultiSet = _229_v66
		_ = _rhs43
		var _rhs44 _dafny.Int = _230_v67
		_ = _rhs44
		_229_v66 = _rhs43
		_230_v67 = _rhs44
		_230_v67 = _dafny.IntOfInt64(644)
		_206_v49 = _206_v49
		var _231_v68 _dafny.Array
		_ = _231_v68
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
		_ = _nw31
		_231_v68 = _nw31
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_231_v68), 0))
		_ = _index33
		(_231_v68).ArraySet1(Companion_Default___.Fm1(_230_v67, _121_globalState), (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_231_v68), 0))
		_ = _index34
		(_231_v68).ArraySet1(_230_v67, (_index34).Int())
		var _232_v69 _dafny.Sequence
		_ = _232_v69
		_232_v69 = _dafny.SeqOf(_207_v50)
		var _233_v70 _dafny.MultiSet
		_ = _233_v70
		_233_v70 = _dafny.MultiSetOf(_232_v69)
		var _234_v71 _dafny.Map
		_ = _234_v71
		_234_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_231_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_231_v68), 0))).Int()).(_dafny.Int), _233_v70)
		_207_v50 = (func() bool {
			if ((func() _dafny.MultiSet {
				if (_234_v71).Contains(_230_v67) {
					return (_234_v71).Get(_230_v67).(_dafny.MultiSet)
				}
				return _233_v70
			})()).IsDisjointFrom(_233_v70) {
				return true
			}
			return _207_v50
		})()
	}
	var _235_v72 _dafny.Int
	_ = _235_v72
	_235_v72 = _dafny.IntOfInt64(694)
	_235_v72 = _235_v72
	var _236_v73 _dafny.Array
	_ = _236_v73
	var _len0_10 _dafny.Int = _dafny.IntOfInt64(3)
	_ = _len0_10
	var _nw32 _dafny.Array
	_ = _nw32
	if _len0_10.Cmp(_dafny.Zero) == 0 {
		_nw32 = _dafny.NewArray(_len0_10)
	} else {
		var _init10 func(_dafny.Int) _dafny.Sequence = (func(_237_v50 bool) func(_dafny.Int) _dafny.Sequence {
			return func(_238_i18 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(true, _237_v50)
			}
		})(_207_v50)
		_ = _init10
		var _element0_10 = _init10(_dafny.Zero)
		_ = _element0_10
		_nw32 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
		(_nw32).ArraySet1(_element0_10, 0)
		var _nativeLen0_10 = (_len0_10).Int()
		_ = _nativeLen0_10
		for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
			(_nw32).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
		}
	}
	_236_v73 = _nw32
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_236_v73), 0))); ; {
		_guard_loop_1, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _239_i17 _dafny.Int
		_239_i17 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_239_i17).Sign() != -1) && ((_239_i17).Cmp(_dafny.ArrayLen((_236_v73), 0)) < 0)) {
			(_236_v73).ArraySet1(_dafny.SeqOf(_207_v50, _207_v50, _207_v50), (_239_i17).Int())
		}
	}
	var _240_v74 *C0
	_ = _240_v74
	var _nw33 *C0 = New_C0_()
	_ = _nw33
	_nw33.Ctor__()
	_240_v74 = _nw33
	if Companion_Default___.Fm0((_235_v72).Minus(_235_v72), _121_globalState) {
		var _241_v75 _dafny.CodePoint
		_ = _241_v75
		_241_v75 = _dafny.CodePoint('o')
		var _242_v76 _dafny.Map
		_ = _242_v76
		_242_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, _241_v75)
		var _243_v77 _dafny.Map
		_ = _243_v77
		_243_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, _242_v76)
		var _244_v78 _dafny.Map
		_ = _244_v78
		_244_v78 = (func() _dafny.Map {
			if (_243_v77).Contains(_207_v50) {
				return (_243_v77).Get(_207_v50).(_dafny.Map)
			}
			return _242_v76
		})()
		var _245_v79 _dafny.Sequence
		_ = _245_v79
		_245_v79 = _dafny.SeqOf(((func() _dafny.Map {
			if _207_v50 {
				return _242_v76
			}
			return (_244_v78)
		})()).Cardinality())
		var _rhs45 _dafny.Int = _235_v72
		_ = _rhs45
		var _rhs46 _dafny.Sequence = _245_v79
		_ = _rhs46
		var _lhs19 *GlobalState = _121_globalState
		_ = _lhs19
		_235_v72 = _rhs45
		_lhs19.F0 = _rhs46
		var _246_v80 _dafny.MultiSet
		_ = _246_v80
		_246_v80 = _dafny.MultiSetOf(_235_v72, _235_v72)
		var _247_v81 _dafny.Map
		_ = _247_v81
		_247_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v72, _207_v50)
		var _248_v82 D2
		_ = _248_v82
		_248_v82 = Companion_D2_.Create_DC5_(_235_v72, _235_v72)
		var _249_v84 _dafny.Array
		_ = _249_v84
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw34
		_249_v84 = _nw34
		var _250_v85 _dafny.MultiSet
		_ = _250_v85
		_250_v85 = _dafny.MultiSetOf(_249_v84)
		var _251_v86 _dafny.Array
		_ = _251_v86
		var _nwElement0_4 _dafny.Int = (_246_v80).Cardinality()
		_ = _nwElement0_4
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_4, 0)
		(_nw35).ArraySet1(_235_v72, 1)
		(_nw35).ArraySet1((Companion_Default___.Fm6(_207_v50, _247_v81, (Companion_D0_.Create_DC0_(_207_v50)).Dtor_cf0(), _235_v72, _121_globalState)).Cardinality(), 2)
		(_nw35).ArraySet1((_dafny.Zero).Minus(_235_v72), 3)
		(_nw35).ArraySet1((_248_v82).Dtor_cf8(), 4)
		(_nw35).ArraySet1(Companion_Default___.Fm1((func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(144), _dafny.IntOfInt64(803))); ; {
				_compr_7, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _252_v83 _dafny.Int
				_252_v83 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(144)).Cmp(_252_v83) <= 0) && ((_252_v83).Cmp(_dafny.IntOfInt64(803)) < 0) {
					_coll7.Add((_252_v83).Plus(_235_v72), _dafny.IntOfInt64(785))
				}
			}
			return _coll7.ToMap()
		}()).Cardinality(), _121_globalState), 5)
		(_nw35).ArraySet1((_250_v85).Cardinality(), 6)
		(_nw35).ArraySet1(_235_v72, 7)
		(_nw35).ArraySet1(_235_v72, 8)
		_251_v86 = _nw35
		var _253_v87 _dafny.Sequence
		_ = _253_v87
		_253_v87 = _dafny.SeqOf(_249_v84, _249_v84)
		var _254_v88 _dafny.Set
		_ = _254_v88
		_254_v88 = _dafny.SetOf(_dafny.IntOfUint32((_253_v87).Cardinality()))
		_254_v88 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_245_v79, (Companion_Default___.SafeIndex(_235_v72, _dafny.IntOfUint32((_245_v79).Cardinality()))).Uint32(), _235_v72)).Cardinality())), _235_v72)
		var _255_v89 _dafny.Sequence
		_ = _255_v89
		_255_v89 = _dafny.UnicodeSeqOfUtf8Bytes("hlrf")
		var _256_v90 _dafny.Map
		_ = _256_v90
		_256_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_235_v72), (func() _dafny.Sequence {
			if !(!(_207_v50)) {
				return _dafny.UnicodeSeqOfUtf8Bytes("iff")
			}
			return _255_v89
		})())
		_255_v89 = (func() _dafny.Sequence {
			if (_256_v90).Contains(_235_v72) {
				return (_256_v90).Get(_235_v72).(_dafny.Sequence)
			}
			return _255_v89
		})()
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_251_v86), 0))
		_ = _index35
		(_251_v86).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, _235_v72)).Cardinality(), (_index35).Int())
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_251_v86), 0))
		_ = _index36
		(_251_v86).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_255_v89, _255_v89)).Cardinality()), (_245_v79).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.IntOfUint32((_245_v79).Cardinality()))).Uint32()).(_dafny.Int)), (_index36).Int())
		_206_v49 = _240_v74
	} else {
		var _257_v91 *C0
		_ = _257_v91
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__()
		_257_v91 = _nw36
		_206_v49 = _240_v74
		(_240_v74).M1(_235_v72, true, _121_globalState)
		var _258_v92 _dafny.Sequence
		_ = _258_v92
		_258_v92 = _dafny.SeqOf((_dafny.IntOfInt64(256)).Cmp(_235_v72) < 0)
		_207_v50 = (_258_v92).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-242), _dafny.IntOfUint32((_258_v92).Cardinality()))).Uint32()).(bool)
		var _259_v93 _dafny.Sequence
		_ = _259_v93
		_259_v93 = _dafny.SeqOf((_dafny.Zero).Minus(_235_v72), (func() _dafny.Int {
			if _207_v50 {
				return _235_v72
			}
			return _235_v72
		})(), _235_v72)
		_235_v72 = (_259_v93).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-252), _dafny.IntOfUint32((_259_v93).Cardinality()))).Uint32()).(_dafny.Int)
	}
	var _260_v94 _dafny.MultiSet
	_ = _260_v94
	_260_v94 = _dafny.MultiSetOf(_240_v74, _240_v74, _240_v74)
	_260_v94 = _260_v94
	var _hi1 _dafny.Int = _235_v72
	_ = _hi1
	for _261_i19 := _dafny.IntOfInt64(762); _261_i19.Cmp(_hi1) < 0; _261_i19 = _261_i19.Plus(_dafny.One) {
		(_121_globalState).F1 = _207_v50
		(_121_globalState).F1 = _207_v50
		_235_v72 = _235_v72
		var _262_v95 _dafny.Array
		_ = _262_v95
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_11
		var _nw37 _dafny.Array
		_ = _nw37
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw37 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = (func(_263_v72 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_264_i20 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_264_i20, _263_v72)
				}
			})(_235_v72)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw37 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw37).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw37).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_262_v95 = _nw37
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_262_v95), 0))
		_ = _index37
		(_262_v95).ArraySet1((_dafny.Zero).Minus(_261_i19), (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_262_v95), 0))
		_ = _index38
		(_262_v95).ArraySet1(_261_i19, (_index38).Int())
	}
	var _hi2 _dafny.Int = _235_v72
	_ = _hi2
	for _265_i21 := _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality()); _265_i21.Cmp(_hi2) < 0; _265_i21 = _265_i21.Plus(_dafny.One) {
		var _266_v96 _dafny.MultiSet
		_ = _266_v96
		_266_v96 = _dafny.MultiSetOf(_265_i21)
		_235_v72 = ((func() _dafny.Int {
			if (_266_v96).Contains(_dafny.IntOfInt64(247)) {
				return (_266_v96).Multiplicity(_dafny.IntOfInt64(247))
			}
			return _265_i21
		})()).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(92))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_267_i22 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality())))
		var _268_v97 _dafny.CodePoint
		_ = _268_v97
		_268_v97 = _dafny.CodePoint('f')
		_268_v97 = _268_v97
		if false {
			var _269_v98 _dafny.Array
			_ = _269_v98
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw38
			_269_v98 = _nw38
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_269_v98), 0))
			_ = _index39
			(_269_v98).ArraySet1((Companion_Default___.Fm1(_dafny.IntOfInt64(-790), _121_globalState)).Minus(_235_v72), (_index39).Int())
			var _270_v99 _dafny.Sequence
			_ = _270_v99
			_270_v99 = _dafny.UnicodeSeqOfUtf8Bytes("ensayrry")
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_269_v98), 0))
			_ = _index40
			(_269_v98).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm3(_121_globalState), _270_v99), _270_v99)).Cardinality()), (_index40).Int())
			var _271_v100 D4
			_ = _271_v100
			_271_v100 = Companion_D4_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_235_v72, _121_globalState), (_269_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_269_v98), 0))).Int()).(_dafny.Int)))
			var _272_v101 _dafny.Map
			_ = _272_v101
			_272_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_269_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_269_v98), 0))).Int()).(_dafny.Int))
			var _273_v102 _dafny.Map
			_ = _273_v102
			_273_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_271_v100).Dtor_cf12()).Merge(_272_v101), _268_v97)
			_273_v102 = ((_273_v102).Update(_272_v101, _268_v97)).Merge(_273_v102)
			_268_v97 = _dafny.CodePoint('k')
			var _274_v103 *C0
			_ = _274_v103
			var _nw39 *C0 = New_C0_()
			_ = _nw39
			_nw39.Ctor__()
			_274_v103 = _nw39
			var _275_v104 _dafny.Map
			_ = _275_v104
			_275_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _235_v72)
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_269_v98), 0))
			_ = _index41
			(_269_v98).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfInt64(-723))).Minus(Companion_Default___.SafeDivisionInt((_275_v104).Cardinality(), (_269_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_269_v98), 0))).Int()).(_dafny.Int))), (_index41).Int())
		} else {
			var _276_v105 *C0
			_ = _276_v105
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__()
			_276_v105 = _nw40
			_207_v50 = true
			var _277_v106 _dafny.Sequence
			_ = _277_v106
			_277_v106 = _dafny.UnicodeSeqOfUtf8Bytes("sxn")
			_277_v106 = _277_v106
			var _278_v107 _dafny.MultiSet
			_ = _278_v107
			_278_v107 = _dafny.MultiSetOf(_207_v50)
			var _279_v108 D5
			_ = _279_v108
			_279_v108 = Companion_D5_.Create_DC10_(_278_v107)
			(_121_globalState).F1 = ((_278_v107).Union((_279_v108).Dtor_cf16())).IsProperSubsetOf((_278_v107).Difference(_278_v107))
			_277_v106 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_277_v106, (Companion_Default___.SafeIndex(_265_i21, _dafny.IntOfUint32((_277_v106).Cardinality()))).Uint32(), _268_v97), Companion_Default___.Fm3(_121_globalState)), _277_v106)
		}
		var _280_v109 _dafny.Set
		_ = _280_v109
		_280_v109 = _dafny.SetOf(_235_v72)
		var _281_v110 _dafny.Sequence
		_ = _281_v110
		_281_v110 = _dafny.UnicodeSeqOfUtf8Bytes("ybjqinu")
		_235_v72 = ((func() _dafny.Int {
			if _207_v50 {
				return (_280_v109).Cardinality()
			}
			return _dafny.IntOfUint32((_281_v110).Cardinality())
		})()).Plus((_235_v72).Minus(_235_v72))
	}
	var _282_v111 _dafny.MultiSet
	_ = _282_v111
	_282_v111 = _dafny.MultiSetOf(_207_v50)
	(_206_v49).M1((_282_v111).Cardinality(), _207_v50, _121_globalState)
	(_240_v74).M1(_235_v72, _207_v50, _121_globalState)
	var _hi3 _dafny.Int = _dafny.IntOfInt64(285)
	_ = _hi3
	for _283_i23 := _dafny.IntOfInt64(188); _283_i23.Cmp(_hi3) < 0; _283_i23 = _283_i23.Plus(_dafny.One) {
		var _284_v112 _dafny.Sequence
		_ = _284_v112
		_284_v112 = _dafny.UnicodeSeqOfUtf8Bytes("imxv")
		var _285_v113 _dafny.Sequence
		_ = _285_v113
		_285_v113 = _dafny.SeqOf(_283_i23, (_dafny.Zero).Minus(_283_i23), (_235_v72).Times(Companion_Default___.Fm1(_dafny.IntOfUint32((_284_v112).Cardinality()), _121_globalState)))
		_235_v72 = (_285_v113).Select((Companion_Default___.SafeIndex(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(612))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_286_i24 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality()))).Plus(Companion_Default___.Fm1(_283_i23, _121_globalState)), _dafny.IntOfUint32((_285_v113).Cardinality()))).Uint32()).(_dafny.Int)
		_236_v73 = _236_v73
		if ((_235_v72).Times(_dafny.IntOfInt64(-552))).Cmp((Companion_Default___.Fm7(_121_globalState)).Cardinality()) != 0 {
			var _287_v114 D2
			_ = _287_v114
			_287_v114 = Companion_D2_.Create_DC5_((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-94))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_288_v72 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_289_i25 _dafny.Int) _dafny.Int {
					return _288_v72
				}
			})(_235_v72))))).Cardinality(), _dafny.IntOfInt64(902))
			var _290_v115 _dafny.MultiSet
			_ = _290_v115
			_290_v115 = _dafny.MultiSetOf(_287_v114, _287_v114)
			var _291_v116 _dafny.Set
			_ = _291_v116
			_291_v116 = _dafny.SetOf(_283_i23, _dafny.IntOfInt64(914))
			_235_v72 = (func() _dafny.Int {
				if (_290_v115).Contains(_287_v114) {
					return (_290_v115).Multiplicity(_287_v114)
				}
				return (_235_v72).Times((_291_v116).Cardinality())
			})()
			var _292_v117 _dafny.CodePoint
			_ = _292_v117
			_292_v117 = _dafny.CodePoint('b')
			var _293_v118 _dafny.MultiSet
			_ = _293_v118
			_293_v118 = _dafny.MultiSetOf(_292_v117, _dafny.CodePoint('p'))
			var _294_v119 _dafny.Sequence
			_ = _294_v119
			_294_v119 = _dafny.SeqOf(_207_v50, Companion_Default___.Fm0(_283_i23, _121_globalState))
			var _295_v120 _dafny.Array
			_ = _295_v120
			var _nwElement0_5 _dafny.Int = _235_v72
			_ = _nwElement0_5
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(5))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_5, 0)
			(_nw41).ArraySet1((_285_v113).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_293_v118).Contains(_292_v117) {
					return (_293_v118).Multiplicity(_292_v117)
				}
				return _dafny.IntOfUint32((_294_v119).Cardinality())
			})(), _dafny.IntOfUint32((_285_v113).Cardinality()))).Uint32()).(_dafny.Int), 1)
			(_nw41).ArraySet1(_235_v72, 2)
			(_nw41).ArraySet1(Companion_Default___.SafeModuloInt(_235_v72, _dafny.IntOfUint32((_294_v119).Cardinality())), 3)
			(_nw41).ArraySet1((func() _dafny.Int {
				if _207_v50 {
					return _dafny.IntOfInt64(154)
				}
				return _283_i23
			})(), 4)
			_295_v120 = _nw41
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_295_v120), 0))
			_ = _index42
			(_295_v120).ArraySet1(_235_v72, (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_295_v120), 0))
			_ = _index43
			(_295_v120).ArraySet1(_283_i23, (_index43).Int())
			(_121_globalState).F1 = _207_v50
			var _296_v121 _dafny.Map
			_ = _296_v121
			_296_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, true)
			var _297_v122 _dafny.Array
			_ = _297_v122
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_12
			var _nw42 _dafny.Array
			_ = _nw42
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw42 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.CodePoint = func(_298_i26 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw42 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw42).ArraySet1CodePoint(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw42).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_297_v122 = _nw42
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_297_v122), 0))
			_ = _index44
			(_297_v122).ArraySet1CodePoint(_292_v117, (_index44).Int())
			var _299_v123 _dafny.Array
			_ = _299_v123
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_13
			var _nw43 _dafny.Array
			_ = _nw43
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw43 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.MultiSet = (func(_300_v111 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_301_i27 _dafny.Int) _dafny.MultiSet {
						return _300_v111
					}
				})(_282_v111)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw43 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw43).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw43).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_299_v123 = _nw43
			var _302_v124 D5
			_ = _302_v124
			_302_v124 = Companion_D5_.Create_DC10_(_282_v111)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_295_v120), 0))
			_ = _index45
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_297_v122), 0))
			_ = _index46
			var _rhs47 _dafny.MultiSet = (_282_v111).Union((_302_v124).Dtor_cf16())
			_ = _rhs47
			var _rhs48 _dafny.Map = _296_v121
			_ = _rhs48
			var _rhs49 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_283_i23))
			_ = _rhs49
			var _rhs50 _dafny.CodePoint = _292_v117
			_ = _rhs50
			var _rhs51 _dafny.Array = _299_v123
			_ = _rhs51
			var _lhs20 _dafny.Array = _295_v120
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_295_v120), 0))
			_ = _lhs21
			var _lhs22 _dafny.Array = _297_v122
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_297_v122), 0))
			_ = _lhs23
			_282_v111 = _rhs47
			_296_v121 = _rhs48
			(_lhs20).ArraySet1(_rhs49, (_lhs21).Int())
			(_lhs22).ArraySet1CodePoint(_rhs50, (_lhs23).Int())
			_299_v123 = _rhs51
			(_121_globalState).F1 = Companion_Default___.Fm0(Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ehbbvaf")).Cardinality()), _121_globalState), _121_globalState)
		} else {
			_235_v72 = _235_v72
			var _303_v125 D0
			_ = _303_v125
			_303_v125 = Companion_D0_.Create_DC1_(_283_i23, _235_v72, !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_i23, _207_v50)).Contains(_235_v72))
			_303_v125 = _303_v125
			var _304_v127 _dafny.Array
			_ = _304_v127
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_14
			var _nw44 _dafny.Array
			_ = _nw44
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw44 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = (func(_305_v72 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_306_i28 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_306_i28, (func() _dafny.Map {
							var _coll8 = _dafny.NewMapBuilder()
							_ = _coll8
							for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(374), _dafny.IntOfInt64(471))); ; {
								_compr_8, _ok10 := _iter10()
								if !_ok10 {
									break
								}
								var _307_v126 _dafny.Int
								_307_v126 = interface{}(_compr_8).(_dafny.Int)
								if ((_dafny.IntOfInt64(374)).Cmp(_307_v126) <= 0) && ((_307_v126).Cmp(_dafny.IntOfInt64(471)) < 0) {
									_coll8.Add((_307_v126).Times(_dafny.IntOfInt64(-80)), _305_v72)
								}
							}
							return _coll8.ToMap()
						}()).Cardinality())
					}
				})(_235_v72)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw44 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw44).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw44).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_304_v127 = _nw44
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_304_v127), 0))
			_ = _index47
			(_304_v127).ArraySet1(_283_i23, (_index47).Int())
			var _308_v128 _dafny.CodePoint
			_ = _308_v128
			_308_v128 = _dafny.CodePoint('o')
			var _309_v129 _dafny.Sequence
			_ = _309_v129
			_309_v129 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_284_v112, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.IntOfUint32((_284_v112).Cardinality()))).Uint32(), _308_v128))
			var _310_v130 _dafny.Array
			_ = _310_v130
			var _nwElement0_6 _dafny.Sequence = _284_v112
			_ = _nwElement0_6
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(18))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_6, 0)
			(_nw45).ArraySet1(_284_v112, 1)
			(_nw45).ArraySet1(_284_v112, 2)
			(_nw45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_284_v112, _284_v112), 3)
			(_nw45).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cscllebh"), 4)
			(_nw45).ArraySet1(_284_v112, 5)
			(_nw45).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_284_v112, (Companion_Default___.SafeIndex(_283_i23, _dafny.IntOfUint32((_284_v112).Cardinality()))).Uint32(), _308_v128), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_311_i29 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}))), (Companion_Default___.SafeIndex(_283_i23, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_284_v112, (Companion_Default___.SafeIndex(_283_i23, _dafny.IntOfUint32((_284_v112).Cardinality()))).Uint32(), _308_v128), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}(func(_312_i29 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})))).Cardinality()))).Uint32(), _308_v128), 6)
			(_nw45).ArraySet1((_309_v129).Select((Companion_Default___.SafeIndex((_285_v113).Select((Companion_Default___.SafeIndex(_283_i23, _dafny.IntOfUint32((_285_v113).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_309_v129).Cardinality()))).Uint32()).(_dafny.Sequence), 7)
			(_nw45).ArraySet1(_284_v112, 8)
			(_nw45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_284_v112, _284_v112), 9)
			(_nw45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_284_v112, _284_v112), 10)
			(_nw45).ArraySet1(_284_v112, 11)
			(_nw45).ArraySet1(_284_v112, 12)
			(_nw45).ArraySet1(_284_v112, 13)
			(_nw45).ArraySet1(_284_v112, 14)
			(_nw45).ArraySet1(_284_v112, 15)
			(_nw45).ArraySet1((func() _dafny.Sequence {
				if false {
					return _284_v112
				}
				return _284_v112
			})(), 16)
			(_nw45).ArraySet1(_284_v112, 17)
			_310_v130 = _nw45
			var _313_v131 _dafny.Map
			_ = _313_v131
			_313_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_i23, _235_v72)
			var _314_v132 _dafny.Map
			_ = _314_v132
			_314_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, _240_v74)
			var _315_v133 _dafny.MultiSet
			_ = _315_v133
			_315_v133 = _dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Int {
				if (_313_v131).Contains(_283_i23) {
					return (_313_v131).Get(_283_i23).(_dafny.Int)
				}
				return (_314_v132).Cardinality()
			})()))
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_304_v127), 0))
			_ = _index48
			var _rhs52 bool = !(_207_v50)
			_ = _rhs52
			var _rhs53 _dafny.Int = (_235_v72).Times(((_315_v133).Union(_315_v133)).Cardinality())
			_ = _rhs53
			var _rhs54 _dafny.Array = _310_v130
			_ = _rhs54
			var _rhs55 _dafny.Int = (_dafny.Zero).Minus((_283_i23).Plus(_283_i23))
			_ = _rhs55
			var _lhs24 _dafny.Array = _304_v127
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_304_v127), 0))
			_ = _lhs25
			_207_v50 = _rhs52
			(_lhs24).ArraySet1(_rhs53, (_lhs25).Int())
			_310_v130 = _rhs54
			_235_v72 = _rhs55
			_304_v127 = _304_v127
			_207_v50 = !(_207_v50) || (Companion_Default___.Fm0(_dafny.IntOfInt64(340), _121_globalState))
		}
		(_121_globalState).F1 = _207_v50
	}
	if !(Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_235_v72, _235_v72), _121_globalState)) {
		var _316_v134 _dafny.Sequence
		_ = _316_v134
		_316_v134 = _dafny.UnicodeSeqOfUtf8Bytes("hceiynyhv")
		_316_v134 = _316_v134
		var _317_v135 *C0
		_ = _317_v135
		var _nw46 *C0 = New_C0_()
		_ = _nw46
		_nw46.Ctor__()
		_317_v135 = _nw46
		var _318_v136 _dafny.Array
		_ = _318_v136
		var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
		_ = _nw47
		_318_v136 = _nw47
		var _319_v137 _dafny.MultiSet
		_ = _319_v137
		_319_v137 = _dafny.MultiSetOf(_235_v72, (_dafny.Zero).Minus(_235_v72), _dafny.IntOfUint32((_316_v134).Cardinality()))
		var _320_v138 _dafny.Map
		_ = _320_v138
		_320_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, _235_v72)
		var _321_v139 _dafny.Map
		_ = _321_v139
		_321_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if true {
				return _318_v136
			}
			return _318_v136
		})(), ((_319_v137).Union((_319_v137).Update(_235_v72, Companion_Default___.Abs((_320_v138).Cardinality())))).Cardinality())
		_321_v139 = (_321_v139).Update(_318_v136, _235_v72)
		(_121_globalState).F1 = _207_v50
		var _322_v140 _dafny.Map
		_ = _322_v140
		_322_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v72, _dafny.IntOfInt64(573))
		var _323_v141 _dafny.MultiSet
		_ = _323_v141
		_323_v141 = _dafny.MultiSetOf(_322_v140, _322_v140)
		var _324_v142 _dafny.Map
		_ = _324_v142
		_324_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, _207_v50)
		if !((func() bool {
			if (_324_v142).Contains(false) {
				return (_324_v142).Get(false).(bool)
			}
			return _207_v50
		})()) || ((_323_v141).IsSubsetOf(_323_v141)) {
			var _325_v143 *C0
			_ = _325_v143
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__()
			_325_v143 = _nw48
			(_121_globalState).F1 = (func() bool {
				if _207_v50 {
					return _207_v50
				}
				return _207_v50
			})()
			var _326_v144 _dafny.Sequence
			_ = _326_v144
			_326_v144 = _dafny.SeqOf(_235_v72)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_122_v0), 0))
			_ = _index49
			(_122_v0).ArraySet1(_dafny.Companion_Sequence_.Equal(_326_v144, _326_v144), (_index49).Int())
			var _327_v145 D0
			_ = _327_v145
			_327_v145 = Companion_D0_.Create_DC1_(_235_v72, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-55), _235_v72)).Cardinality(), _207_v50)
			var _328_v146 _dafny.Map
			_ = _328_v146
			_328_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.IntOfInt64(-268), _121_globalState), _326_v144)
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_122_v0), 0))
			_ = _index50
			(_122_v0).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_235_v72, (_dafny.Zero).Minus((Companion_Default___.Fm5(_207_v50, _207_v50, _327_v145, _121_globalState)).Cardinality())), (func() _dafny.Sequence {
				if (_328_v146).Contains(_207_v50) {
					return (_328_v146).Get(_207_v50).(_dafny.Sequence)
				}
				return _326_v144
			})()), Companion_Default___.Fm1(_235_v72, _121_globalState)), (_index50).Int())
			_207_v50 = (_207_v50) && (true)
			_235_v72 = Companion_Default___.SafeModuloInt((_235_v72).Plus(_235_v72), _235_v72)
		} else {
			_235_v72 = ((_dafny.IntOfUint32((_316_v134).Cardinality())).Minus(_235_v72)).Times(_235_v72)
			var _329_v147 _dafny.Array
			_ = _329_v147
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
			_ = _nw49
			_329_v147 = _nw49
			var _330_v148 _dafny.Array
			_ = _330_v148
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw50
			_330_v148 = _nw50
			var _331_v149 _dafny.Map
			_ = _331_v149
			_331_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v0, _329_v147)
			var _332_v150 _dafny.Sequence
			_ = _332_v150
			_332_v150 = _dafny.SeqOf((func() _dafny.Array {
				if (_331_v149).Contains(_122_v0) {
					return (_331_v149).Get(_122_v0).(_dafny.Array)
				}
				return _329_v147
			})(), _329_v147, _329_v147, _329_v147)
			var _rhs56 _dafny.Array = (_332_v150).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_235_v72, _121_globalState), _dafny.IntOfUint32((_332_v150).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs56
			var _rhs57 _dafny.Int = _235_v72
			_ = _rhs57
			var _rhs58 bool = !((_207_v50) && (_207_v50)) || (_207_v50)
			_ = _rhs58
			var _rhs59 _dafny.Array = _330_v148
			_ = _rhs59
			var _lhs26 *GlobalState = _121_globalState
			_ = _lhs26
			_329_v147 = _rhs56
			_235_v72 = _rhs57
			_lhs26.F1 = _rhs58
			_330_v148 = _rhs59
			var _333_v151 _dafny.Map
			_ = _333_v151
			_333_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_207_v50)).Cardinality(), _207_v50)
			(_121_globalState).F1 = Companion_Default___.Fm0((_235_v72).Minus((_333_v151).Cardinality()), _121_globalState)
			var _334_v152 _dafny.CodePoint
			_ = _334_v152
			_334_v152 = _dafny.CodePoint('h')
			var _335_v153 _dafny.Map
			_ = _335_v153
			_335_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(_207_v50, _dafny.MultiSetOf(_334_v152), _121_globalState), _dafny.SeqOf(_235_v72, _235_v72)), _317_v135)
			var _336_v154 _dafny.Sequence
			_ = _336_v154
			_336_v154 = _dafny.SeqOf(_235_v72, _235_v72)
			_317_v135 = (func() *C0 {
				if (_335_v153).Contains(_336_v154) {
					return (_335_v153).Get(_336_v154).(*C0)
				}
				return (func() *C0 {
					if _207_v50 {
						return _317_v135
					}
					return _317_v135
				})()
			})()
			var _337_v155 _dafny.Array
			_ = _337_v155
			var _nw51 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
			_ = _nw51
			_337_v155 = _nw51
			_337_v155 = _337_v155
		}
	} else {
		(_121_globalState).F1 = _207_v50
		var _338_v156 _dafny.Sequence
		_ = _338_v156
		_338_v156 = _dafny.SeqOf(_235_v72)
		var _339_v157 _dafny.Map
		_ = _339_v157
		_339_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_338_v156, (Companion_Default___.SafeIndex((_282_v111).Cardinality(), _dafny.IntOfUint32((_338_v156).Cardinality()))).Uint32(), _235_v72)).Cardinality())).Minus(_235_v72), _207_v50)
		(_121_globalState).F1 = (func() bool {
			if (_339_v157).Contains(_235_v72) {
				return (_339_v157).Get(_235_v72).(bool)
			}
			return _207_v50
		})()
		var _340_v158 _dafny.Sequence
		_ = _340_v158
		_340_v158 = _dafny.SeqOf(_207_v50, _207_v50)
		_340_v158 = _dafny.SeqOf(_207_v50, _207_v50)
		(_121_globalState).F1 = ((_282_v111).Intersection(_282_v111)).IsProperSubsetOf(_282_v111)
		var _341_v159 _dafny.Array
		_ = _341_v159
		var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw52
		_341_v159 = _nw52
		var _342_v160 _dafny.MultiSet
		_ = _342_v160
		_342_v160 = _dafny.MultiSetOf(_341_v159)
		var _343_v161 _dafny.MultiSet
		_ = _343_v161
		_343_v161 = _342_v160
		var _344_v162 _dafny.Map
		_ = _344_v162
		_344_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v72, ((_343_v161).Intersection(_342_v160)).Cardinality())
		var _345_v163 _dafny.CodePoint
		_ = _345_v163
		_345_v163 = _dafny.CodePoint('p')
		var _346_v164 _dafny.Map
		_ = _346_v164
		_346_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v50, _345_v163)
		var _347_v165 _dafny.Map
		_ = _347_v165
		_347_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v164, _dafny.IntOfInt64(133))
		_344_v162 = (_344_v162).Update(_235_v72, (func() _dafny.Int {
			if (_347_v165).Contains(_346_v164) {
				return (_347_v165).Get(_346_v164).(_dafny.Int)
			}
			return _235_v72
		})())
	}
	var _348_v166 _dafny.Sequence
	_ = _348_v166
	_348_v166 = _dafny.SeqOf(true)
	_207_v50 = _dafny.Companion_Sequence_.IsPrefixOf(_348_v166, _348_v166)
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_121_globalState.F0, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_121_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v1).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_207_v50)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_235_v72)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_236_v73).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_236_v73).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_236_v73).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v94).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_282_v111).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_348_v166, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
}

// End of class Default__

// Definition of datatype D0
type D0 struct {
	Data_D0_
}

func (_this D0) Get_() Data_D0_ {
	return _this.Data_D0_
}

type Data_D0_ interface {
	isD0()
}

type CompanionStruct_D0_ struct {
}

var Companion_D0_ = CompanionStruct_D0_{}

type D0_DC1 struct {
	Cf1 _dafny.Int
	Cf2 _dafny.Int
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 bool
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 bool) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D0) Equals(other D0) bool {
	switch data1 := _this.Get_().(type) {
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D0) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D0)
	return ok && _this.Equals(typed)
}

func Type_D0_() _dafny.TypeDescriptor {
	return type_D0_{}
}

type type_D0_ struct {
}

func (_this type_D0_) Default() interface{} {
	return Companion_D0_.Default()
}

func (_this type_D0_) String() string {
	return "main.D0"
}
func (_this D0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D0{}

// End of datatype D0

// Definition of datatype D1
type D1 struct {
	Data_D1_
}

func (_this D1) Get_() Data_D1_ {
	return _this.Data_D1_
}

type Data_D1_ interface {
	isD1()
}

type CompanionStruct_D1_ struct {
}

var Companion_D1_ = CompanionStruct_D1_{}

type D1_DC2 struct {
	Cf4 _dafny.Sequence
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + data.Cf4.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D1)
	return ok && _this.Equals(typed)
}

func Type_D1_() _dafny.TypeDescriptor {
	return type_D1_{}
}

type type_D1_ struct {
}

func (_this type_D1_) Default() interface{} {
	return Companion_D1_.Default()
}

func (_this type_D1_) String() string {
	return "main.D1"
}
func (_this D1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D1{}

// End of datatype D1

// Definition of datatype D2
type D2 struct {
	Data_D2_
}

func (_this D2) Get_() Data_D2_ {
	return _this.Data_D2_
}

type Data_D2_ interface {
	isD2()
}

type CompanionStruct_D2_ struct {
}

var Companion_D2_ = CompanionStruct_D2_{}

type D2_DC4 struct {
	Cf6 bool
	Cf7 bool
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf6 bool, Cf7 bool) D2 {
	return D2{D2_DC4{Cf6, Cf7}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC5 struct {
	Cf8 _dafny.Int
	Cf9 _dafny.Int
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf8 _dafny.Int, Cf9 _dafny.Int) D2 {
	return D2{D2_DC5{Cf8, Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC3 struct {
	Cf5 _dafny.MultiSet
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf5 _dafny.MultiSet) D2 {
	return D2{D2_DC3{Cf5}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC6 struct {
	Cf10 D2
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 D2) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(false, false)
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf5() _dafny.MultiSet {
	return _this.Get_().(D2_DC3).Cf5
}

func (_this D2) Dtor_cf10() D2 {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf5.Equals(data2.Cf5)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D2) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D2)
	return ok && _this.Equals(typed)
}

func Type_D2_() _dafny.TypeDescriptor {
	return type_D2_{}
}

type type_D2_ struct {
}

func (_this type_D2_) Default() interface{} {
	return Companion_D2_.Default()
}

func (_this type_D2_) String() string {
	return "main.D2"
}
func (_this D2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D2{}

// End of datatype D2

// Definition of datatype D3
type D3 struct {
	Data_D3_
}

func (_this D3) Get_() Data_D3_ {
	return _this.Data_D3_
}

type Data_D3_ interface {
	isD3()
}

type CompanionStruct_D3_ struct {
}

var Companion_D3_ = CompanionStruct_D3_{}

type D3_DC7 struct {
	Cf11 _dafny.Map
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 _dafny.Map) D3 {
	return D3{D3_DC7{Cf11}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D3) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D3)
	return ok && _this.Equals(typed)
}

func Type_D3_() _dafny.TypeDescriptor {
	return type_D3_{}
}

type type_D3_ struct {
}

func (_this type_D3_) Default() interface{} {
	return Companion_D3_.Default()
}

func (_this type_D3_) String() string {
	return "main.D3"
}
func (_this D3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D3{}

// End of datatype D3

// Definition of datatype D4
type D4 struct {
	Data_D4_
}

func (_this D4) Get_() Data_D4_ {
	return _this.Data_D4_
}

type Data_D4_ interface {
	isD4()
}

type CompanionStruct_D4_ struct {
}

var Companion_D4_ = CompanionStruct_D4_{}

type D4_DC9 struct {
	Cf13 _dafny.Array
	Cf14 _dafny.Int
	Cf15 _dafny.Array
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf13 _dafny.Array, Cf14 _dafny.Int, Cf15 _dafny.Array) D4 {
	return D4{D4_DC9{Cf13, Cf14, Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC8 struct {
	Cf12 _dafny.Map
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf12 _dafny.Map) D4 {
	return D4{D4_DC8{Cf12}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC9_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D4) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf13
}

func (_this D4) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D4_DC8).Cf12
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf12.Equals(data2.Cf12)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D4) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D4)
	return ok && _this.Equals(typed)
}

func Type_D4_() _dafny.TypeDescriptor {
	return type_D4_{}
}

type type_D4_ struct {
}

func (_this type_D4_) Default() interface{} {
	return Companion_D4_.Default()
}

func (_this type_D4_) String() string {
	return "main.D4"
}
func (_this D4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D4{}

// End of datatype D4

// Definition of datatype D5
type D5 struct {
	Data_D5_
}

func (_this D5) Get_() Data_D5_ {
	return _this.Data_D5_
}

type Data_D5_ interface {
	isD5()
}

type CompanionStruct_D5_ struct {
}

var Companion_D5_ = CompanionStruct_D5_{}

type D5_DC11 struct {
	Cf17 *C0
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf17 *C0) D5 {
	return D5{D5_DC11{Cf17}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC12 struct {
	Cf18 _dafny.Int
	Cf19 _dafny.CodePoint
	Cf20 bool
	Cf21 _dafny.Sequence
	Cf22 _dafny.Int
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf18 _dafny.Int, Cf19 _dafny.CodePoint, Cf20 bool, Cf21 _dafny.Sequence, Cf22 _dafny.Int) D5 {
	return D5{D5_DC12{Cf18, Cf19, Cf20, Cf21, Cf22}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC10 struct {
	Cf16 _dafny.MultiSet
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf16 _dafny.MultiSet) D5 {
	return D5{D5_DC10{Cf16}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC11_((*C0)(nil))
}

func (_this D5) Dtor_cf17() *C0 {
	return _this.Get_().(D5_DC11).Cf17
}

func (_this D5) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) Dtor_cf19() _dafny.CodePoint {
	return _this.Get_().(D5_DC12).Cf19
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC12).Cf20
}

func (_this D5) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf22
}

func (_this D5) Dtor_cf16() _dafny.MultiSet {
	return _this.Get_().(D5_DC10).Cf16
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + data.Cf21.VerbatimString(true) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf17 == data2.Cf17
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20 && data1.Cf21.Equals(data2.Cf21) && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D5) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D5)
	return ok && _this.Equals(typed)
}

func Type_D5_() _dafny.TypeDescriptor {
	return type_D5_{}
}

type type_D5_ struct {
}

func (_this type_D5_) Default() interface{} {
	return Companion_D5_.Default()
}

func (_this type_D5_) String() string {
	return "main.D5"
}
func (_this D5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D5{}

// End of datatype D5

// Definition of datatype D6
type D6 struct {
	Data_D6_
}

func (_this D6) Get_() Data_D6_ {
	return _this.Data_D6_
}

type Data_D6_ interface {
	isD6()
}

type CompanionStruct_D6_ struct {
}

var Companion_D6_ = CompanionStruct_D6_{}

type D6_DC13 struct {
	Cf23 _dafny.MultiSet
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf23 _dafny.MultiSet) D6 {
	return D6{D6_DC13{Cf23}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D6) Dtor_cf23() _dafny.MultiSet {
	return _this.Get_().(D6_DC13).Cf23
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf23.Equals(data2.Cf23)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D6) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D6)
	return ok && _this.Equals(typed)
}

func Type_D6_() _dafny.TypeDescriptor {
	return type_D6_{}
}

type type_D6_ struct {
}

func (_this type_D6_) Default() interface{} {
	return Companion_D6_.Default()
}

func (_this type_D6_) String() string {
	return "main.D6"
}
func (_this D6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D6{}

// End of datatype D6

// Definition of datatype D7
type D7 struct {
	Data_D7_
}

func (_this D7) Get_() Data_D7_ {
	return _this.Data_D7_
}

type Data_D7_ interface {
	isD7()
}

type CompanionStruct_D7_ struct {
}

var Companion_D7_ = CompanionStruct_D7_{}

type D7_DC15 struct {
	Cf25 _dafny.Sequence
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf25 _dafny.Sequence) D7 {
	return D7{D7_DC15{Cf25}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC14 struct {
	Cf24 _dafny.Map
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf24 _dafny.Map) D7 {
	return D7{D7_DC14{Cf24}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC15_(_dafny.EmptySeq)
}

func (_this D7) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D7_DC15).Cf25
}

func (_this D7) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D7_DC14).Cf24
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + data.Cf25.VerbatimString(true) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf24.Equals(data2.Cf24)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D7) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D7)
	return ok && _this.Equals(typed)
}

func Type_D7_() _dafny.TypeDescriptor {
	return type_D7_{}
}

type type_D7_ struct {
}

func (_this type_D7_) Default() interface{} {
	return Companion_D7_.Default()
}

func (_this type_D7_) String() string {
	return "main.D7"
}
func (_this D7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D7{}

// End of datatype D7

// Definition of class GlobalState
type GlobalState struct {
	F0 _dafny.Sequence
	F1 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptySeq
	_this.F1 = false
	return &_this
}

type CompanionStruct_GlobalState_ struct {
}

var Companion_GlobalState_ = CompanionStruct_GlobalState_{}

func (_this *GlobalState) Equals(other *GlobalState) bool {
	return _this == other
}

func (_this *GlobalState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*GlobalState)
	return ok && _this.Equals(other)
}

func (*GlobalState) String() string {
	return "_module.GlobalState"
}

func Type_GlobalState_() _dafny.TypeDescriptor {
	return type_GlobalState_{}
}

type type_GlobalState_ struct {
}

func (_this type_GlobalState_) Default() interface{} {
	return (*GlobalState)(nil)
}

func (_this type_GlobalState_) String() string {
	return "main.GlobalState"
}
func (_this *GlobalState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &GlobalState{}

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 bool) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	dummy byte
}

func New_C0_() *C0 {
	_this := C0{}

	return &_this
}

type CompanionStruct_C0_ struct {
}

var Companion_C0_ = CompanionStruct_C0_{}

func (_this *C0) Equals(other *C0) bool {
	return _this == other
}

func (_this *C0) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C0)
	return ok && _this.Equals(other)
}

func (*C0) String() string {
	return "_module.C0"
}

func Type_C0_() _dafny.TypeDescriptor {
	return type_C0_{}
}

type type_C0_ struct {
}

func (_this type_C0_) Default() interface{} {
	return (*C0)(nil)
}

func (_this type_C0_) String() string {
	return "main.C0"
}
func (_this *C0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__() {
	{
	}
}
func (_this *C0) M1(p0 _dafny.Int, p1 bool, globalState *GlobalState) {
	{
		var _hi4 _dafny.Int = p0
		_ = _hi4
		for _349_i0 := p0; _349_i0.Cmp(_hi4) < 0; _349_i0 = _349_i0.Plus(_dafny.One) {
			var _350_v0 _dafny.Int
			_ = _350_v0
			_350_v0 = _dafny.IntOfInt64(-15)
			var _351_v2 _dafny.CodePoint
			_ = _351_v2
			_351_v2 = _dafny.CodePoint('e')
			var _352_v3 _dafny.Map
			_ = _352_v3
			_352_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_351_v2)).Cardinality(), p0)
			_350_v0 = (_350_v0).Times(((func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(666), _dafny.IntOfInt64(-533))); ; {
					_compr_9, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _353_v1 _dafny.Int
					_353_v1 = interface{}(_compr_9).(_dafny.Int)
					if ((_dafny.IntOfInt64(666)).Cmp(_353_v1) <= 0) && ((_353_v1).Cmp(_dafny.IntOfInt64(-533)) < 0) {
						_coll9.Add(Companion_Default___.SafeModuloInt(_353_v1, _350_v0), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dchcbikbk")).Cardinality()))
					}
				}
				return _coll9.ToMap()
			}()).Merge(_352_v3)).Cardinality())
			var _354_v4 _dafny.Sequence
			_ = _354_v4
			_354_v4 = _dafny.SeqOf(_349_i0)
			var _355_v5 _dafny.Sequence
			_ = _355_v5
			_355_v5 = _dafny.SeqOf(_354_v4)
			var _356_v6 _dafny.Map
			_ = _356_v6
			_356_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_350_v0, !(p1))
			(globalState).F0 = (_355_v5).Select((Companion_Default___.SafeIndex(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)).Merge(_356_v6)).Cardinality(), _dafny.IntOfUint32((_355_v5).Cardinality()))).Uint32()).(_dafny.Sequence)
			(globalState).F1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p1), (Companion_Default___.SafeIndex(_349_i0, _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))).Uint32(), p1)).Cardinality())).Cmp((Companion_D0_.Create_DC1_(_350_v0, Companion_Default___.Fm1(_349_i0, globalState), p1)).Dtor_cf1()) > 0
			(globalState).F1 = p1
		}
		(globalState).F1 = p1
		var _357_v7 _dafny.Int
		_ = _357_v7
		_357_v7 = _dafny.IntOfInt64(436)
		var _358_v8 _dafny.MultiSet
		_ = _358_v8
		_358_v8 = _dafny.MultiSetOf(p1, p1)
		_357_v7 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm1(p0, globalState), (_dafny.IntOfInt64(803)).Plus((_358_v8).Cardinality()))
		var _359_v9 _dafny.Sequence
		_ = _359_v9
		_359_v9 = _dafny.UnicodeSeqOfUtf8Bytes("hb")
		var _360_v10 _dafny.Sequence
		_ = _360_v10
		_360_v10 = _dafny.SeqOf(p1)
		var _361_v11 _dafny.Map
		_ = _361_v11
		_361_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm0(_357_v7, globalState))).Cardinality()), p1)
		var _362_v12 _dafny.Array
		_ = _362_v12
		var _nwElement0_7 _dafny.Sequence = _dafny.SeqOf(true, p1)
		_ = _nwElement0_7
		var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(20))
		_ = _nw53
		(_nw53).ArraySet1(_nwElement0_7, 0)
		(_nw53).ArraySet1(_360_v10, 1)
		(_nw53).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_360_v10, (Companion_Default___.SafeIndex(_357_v7, _dafny.IntOfUint32((_360_v10).Cardinality()))).Uint32(), p1), _360_v10), 2)
		(_nw53).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_360_v10, _360_v10), 3)
		(_nw53).ArraySet1(_dafny.SeqOf(p1, p1), 4)
		(_nw53).ArraySet1(_360_v10, 5)
		(_nw53).ArraySet1(_dafny.SeqOf(p1, p1), 6)
		(_nw53).ArraySet1(_dafny.Companion_Sequence_.Update(_360_v10, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.IntOfUint32((_360_v10).Cardinality()))).Uint32(), p1), 7)
		(_nw53).ArraySet1(_360_v10, 8)
		(_nw53).ArraySet1(_360_v10, 9)
		(_nw53).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_360_v10, _360_v10), 10)
		(_nw53).ArraySet1(_360_v10, 11)
		(_nw53).ArraySet1(_dafny.SeqOf(p1, p1, p1), 12)
		(_nw53).ArraySet1(_360_v10, 13)
		(_nw53).ArraySet1(_dafny.SeqOf(p1, p1, p1), 14)
		(_nw53).ArraySet1(Companion_Default___.Fm2(p1, p1, _361_v11, p1, globalState), 15)
		(_nw53).ArraySet1(_360_v10, 16)
		(_nw53).ArraySet1((func() _dafny.Sequence {
			if p1 {
				return _360_v10
			}
			return _360_v10
		})(), 17)
		(_nw53).ArraySet1(_360_v10, 18)
		(_nw53).ArraySet1(_dafny.SeqOf(true, p1), 19)
		_362_v12 = _nw53
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_362_v12), 0))
		_ = _index51
		(_362_v12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_360_v10, _360_v10), (_index51).Int())
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_362_v12), 0))
		_ = _index52
		var _rhs60 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm3(globalState), _359_v9)
		_ = _rhs60
		var _rhs61 _dafny.Sequence = _360_v10
		_ = _rhs61
		var _rhs62 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm4(globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_363_v9 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_364_i1 _dafny.Int) _dafny.Sequence {
				return (_363_v9)
			}
		})(_359_v9))))
		_ = _rhs62
		var _rhs63 bool = p1
		_ = _rhs63
		var _lhs27 _dafny.Array = _362_v12
		_ = _lhs27
		var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_362_v12), 0))
		_ = _lhs28
		var _lhs29 *GlobalState = globalState
		_ = _lhs29
		var _lhs30 *GlobalState = globalState
		_ = _lhs30
		_359_v9 = _rhs60
		(_lhs27).ArraySet1(_rhs61, (_lhs28).Int())
		_lhs29.F1 = _rhs62
		_lhs30.F1 = _rhs63
		_357_v7 = _357_v7
		var _365_v13 _dafny.Map
		_ = _365_v13
		_365_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_357_v7, _357_v7)
		var _366_v14 _dafny.Map
		_ = _366_v14
		_366_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cmp((func() _dafny.Int {
			if (_365_v13).Contains(_357_v7) {
				return (_365_v13).Get(_357_v7).(_dafny.Int)
			}
			return p0
		})()) == 0, _dafny.IntOfInt64(-600))
		var _rhs64 _dafny.Int = ((p0).Plus((_dafny.Zero).Minus((_365_v13).Cardinality()))).Times(p0)
		_ = _rhs64
		var _rhs65 _dafny.Int = (func() _dafny.Int {
			if (_366_v14).Contains(p1) {
				return (_366_v14).Get(p1).(_dafny.Int)
			}
			return (_dafny.Zero).Minus(_357_v7)
		})()
		_ = _rhs65
		_357_v7 = _rhs64
		_357_v7 = _rhs65
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
