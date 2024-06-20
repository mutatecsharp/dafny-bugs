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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(315), _dafny.IntOfInt64(987))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _0_v1 _dafny.Int
				_0_v1 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(315)).Cmp(_0_v1) <= 0) && ((_0_v1).Cmp(_dafny.IntOfInt64(987)) < 0) {
					_coll1.Add((_0_v1).Plus((_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality())))
				}
			}
			return _coll1.ToSet()
		}()).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if (func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(315), _dafny.IntOfInt64(987))); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _2_v1 _dafny.Int
					_2_v1 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(315)).Cmp(_2_v1) <= 0) && ((_2_v1).Cmp(_dafny.IntOfInt64(987)) < 0) {
						_coll2.Add((_2_v1).Plus((_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality())))
					}
				}
				return _coll2.ToSet()
			}()).Contains(_1_v0) {
				_coll0.Add((_1_v0).Plus(_dafny.IntOfInt64(171)), func() _dafny.Map {
					var _coll3 = _dafny.NewMapBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(865), true)).Keys().Elements()); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _3_v2 _dafny.Int
						_3_v2 = interface{}(_compr_3).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(865), true)).Contains(_3_v2) {
							_coll3.Add((_3_v2).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eukgmrlmj")).Cardinality())), true)
						}
					}
					return _coll3.ToMap()
				}())
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Plus((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Cardinality())).Cardinality()), _dafny.IntOfInt64(286), _dafny.IntOfInt64(258), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pfru"), false)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kabkgac")).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(-446), _dafny.IntOfInt64(261), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(452), _dafny.IntOfInt64(889), (_dafny.MultiSetOf(true)).Cardinality())).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iy")).Cardinality()), _dafny.IntOfInt64(-558)))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("vhc"), _dafny.UnicodeSeqOfUtf8Bytes("ej"))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _4_v0 _dafny.Sequence
			_4_v0 = interface{}(_compr_4).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("vhc"), _dafny.UnicodeSeqOfUtf8Bytes("ej"))).Contains(_4_v0) {
				_coll4.Add(_4_v0, _dafny.UnicodeSeqOfUtf8Bytes("tlx"))
			}
		}
		return _coll4.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wtjwu"), _dafny.UnicodeSeqOfUtf8Bytes("aykia")))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer0 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_5_i0 _dafny.Int) D2 {
		return Companion_D2_.Create_DC3_(false)
	})), _dafny.SeqOf(Companion_D2_.Create_DC3_(!(true))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(23))).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aeuryux")).Cardinality())), _dafny.IntOfInt64(212))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 D2, p3 bool, globalState *GlobalState) D2 {
	if (_dafny.IntOfInt64(-998)).Cmp(_dafny.IntOfInt64(765)) <= 0 {
		return Companion_D2_.Create_DC3_(false)
	} else {
		return Companion_D2_.Create_DC3_(!(true))
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(988))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_6_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qdliekdl")).Cardinality())
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_7_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('d'))).Cardinality()
	}))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('v')
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if (_dafny.SetOf(true, true, false)).IsProperSubsetOf(_dafny.SetOf(true)) {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(634))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("k"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.UnicodeSeqOfUtf8Bytes("kkbipqqh")))
	}
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source0 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(187))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_9_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(319), _dafny.IntOfInt64(101))).Cardinality())
	})))
	_ = _source0
	var _10___mcc_h0 _dafny.MultiSet = _source0
	_ = _10___mcc_h0
	var _11_cf14 _dafny.MultiSet = _10___mcc_h0
	_ = _11_cf14
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(317))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_12_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})), _dafny.IntOfInt64(665))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.IntOfInt64(471)).Cmp(_dafny.IntOfInt64(-851)) <= 0 {
		return _dafny.SeqOf(_dafny.IntOfInt64(937), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xuck")).Cardinality()))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-132), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), true)).Cardinality())).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(229))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_13_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(613)
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(false)).Difference(_dafny.MultiSetOf(!(true)))).Difference((Companion_D6_.Create_DC13_(_dafny.MultiSetOf(true))).Dtor_cf24())
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(259))).Union(_dafny.SetOf(_dafny.IntOfInt64(-865), _dafny.IntOfInt64(677), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(658), !(!(false)))).Cardinality(), _dafny.IntOfInt64(778), _dafny.IntOfInt64(557)))).Intersection(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-44), _dafny.IntOfInt64(498))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _14_v0 _dafny.Int
				_14_v0 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(-44)).Cmp(_14_v0) <= 0) && ((_14_v0).Cmp(_dafny.IntOfInt64(498)) < 0) {
					_coll6.Add(Companion_Default___.SafeDivisionInt(_14_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(884))).Cardinality()))
				}
			}
			return _coll6.ToSet()
		}()).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _15_v1 _dafny.Int
			_15_v1 = interface{}(_compr_5).(_dafny.Int)
			if (func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-44), _dafny.IntOfInt64(498))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _16_v0 _dafny.Int
					_16_v0 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(-44)).Cmp(_16_v0) <= 0) && ((_16_v0).Cmp(_dafny.IntOfInt64(498)) < 0) {
						_coll7.Add(Companion_Default___.SafeDivisionInt(_16_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(884))).Cardinality()))
					}
				}
				return _coll7.ToSet()
			}()).Contains(_15_v1) {
				_coll5.Add(Companion_Default___.SafeDivisionInt(_15_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dtejhablw")).Cardinality())))
			}
		}
		return _coll5.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("lwb")
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), false, true), _dafny.SeqOf(true))
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((func() bool {
		if false {
			return true
		}
		return !(false)
	})(), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(198))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(598))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_18_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))), (func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(123), _dafny.IntOfInt64(-43))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(123)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(-43)) < 0) {
				_coll8.Add((_19_v0).Minus(_dafny.IntOfInt64(675)))
			}
		}
		return _coll8.ToSet()
	}()).Equals(_dafny.SetOf(_dafny.IntOfInt64(-714))))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Set, p2 bool, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _20_v0 _dafny.Sequence
	_ = _20_v0
	_20_v0 = _dafny.UnicodeSeqOfUtf8Bytes("upn")
	var _21_v1 *C4
	_ = _21_v1
	var _nw0 *C4 = New_C4_()
	_ = _nw0
	_nw0.Ctor__(_20_v0)
	_21_v1 = _nw0
	var _22_v2 _dafny.MultiSet
	_ = _22_v2
	_22_v2 = _dafny.MultiSetOf(_21_v1, _21_v1)
	(globalState).F5 = Companion_Default___.SafeModuloInt((_22_v2).Cardinality(), _dafny.IntOfInt64(73))
	var _23_v3 _dafny.Array
	_ = _23_v3
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(10)
	_ = _len0_0
	var _nw1 _dafny.Array
	_ = _nw1
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw1 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = (func(_24_p0 bool) func(_dafny.Int) bool {
			return func(_25_i0 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfInt64(222)), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_p0, _24_p0)).Cardinality())))
			}
		})(p0)
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
	_23_v3 = _nw1
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
	_ = _index0
	(_23_v3).ArraySet1((_21_v1).Fm1(p2, p0, globalState), (_index0).Int())
	var _26_v4 _dafny.Array
	_ = _26_v4
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_1
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Set = func(_27_i1 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(_dafny.IntOfInt64(122))
		}
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
	_26_v4 = _nw2
	var _28_v5 _dafny.Set
	_ = _28_v5
	_28_v5 = _dafny.SetOf(_dafny.IntOfInt64(-565))
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_26_v4), 0))
	_ = _index1
	(_26_v4).ArraySet1(_28_v5, (_index1).Int())
	var _29_v6 _dafny.Array
	_ = _29_v6
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_2
	var _nw3 _dafny.Array
	_ = _nw3
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw3 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Int = func(_30_i2 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeDivisionInt(_30_i2, _dafny.IntOfInt64(842))
		}
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
	_29_v6 = _nw3
	var _31_v7 _dafny.Int
	_ = _31_v7
	_31_v7 = _dafny.IntOfInt64(-144)
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))
	_ = _index2
	(_29_v6).ArraySet1(_31_v7, (_index2).Int())
	var _32_v8 bool
	_ = _32_v8
	_32_v8 = false
	var _33_v9 _dafny.MultiSet
	_ = _33_v9
	_33_v9 = _dafny.MultiSetOf(_32_v8, p0, p2)
	var _34_v10 _dafny.MultiSet
	_ = _34_v10
	_34_v10 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_35_i3 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(707)
	})))
	var _36_v11 D5
	_ = _36_v11
	_36_v11 = Companion_D5_.Create_DC10_(_34_v10)
	var _37_v12 _dafny.Array
	_ = _37_v12
	var _nw4 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
	_ = _nw4
	_37_v12 = _nw4
	var _38_v13 _dafny.Sequence
	_ = _38_v13
	_38_v13 = _dafny.SeqOf(_37_v12)
	var _39_v14 _dafny.Map
	_ = _39_v14
	_39_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v11, _dafny.IntOfUint32((_38_v13).Cardinality()))
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
	_ = _index3
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_26_v4), 0))
	_ = _index4
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))
	_ = _index5
	var _rhs0 bool = p2
	_ = _rhs0
	var _rhs1 _dafny.Int = Companion_Default___.SafeDivisionInt(((_dafny.Zero).Minus(_31_v7)).Times(_31_v7), _31_v7)
	_ = _rhs1
	var _rhs2 _dafny.Set = _28_v5
	_ = _rhs2
	var _rhs3 _dafny.Int = ((func() _dafny.Int {
		if (_33_v9).Contains(_32_v8) {
			return (_33_v9).Multiplicity(_32_v8)
		}
		return (_39_v14).Cardinality()
	})()).Plus(_31_v7)
	_ = _rhs3
	var _rhs4 bool = p2
	_ = _rhs4
	var _lhs0 _dafny.Array = _23_v3
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
	_ = _lhs1
	var _lhs2 *GlobalState = globalState
	_ = _lhs2
	var _lhs3 _dafny.Array = _26_v4
	_ = _lhs3
	var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_26_v4), 0))
	_ = _lhs4
	var _lhs5 _dafny.Array = _29_v6
	_ = _lhs5
	var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))
	_ = _lhs6
	(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
	_lhs2.F5 = _rhs1
	(_lhs3).ArraySet1(_rhs2, (_lhs4).Int())
	(_lhs5).ArraySet1(_rhs3, (_lhs6).Int())
	_32_v8 = _rhs4
	if (_21_v1).Fm1(p2, !((_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool)) || (_32_v8), globalState) {
		var _40_v15 _dafny.Map
		_ = _40_v15
		_40_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if p2 {
				return p0
			}
			return (_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool)
		})(), _32_v8)
		_40_v15 = (_40_v15).Update(p2, p2)
		(globalState).F5 = (_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
		_ = _index6
		(_23_v3).ArraySet1(_32_v8, (_index6).Int())
		r0 = (_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int)
		(globalState).F5 = (_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int)
	} else {
		var _41_v16 _dafny.Set
		_ = _41_v16
		_41_v16 = _dafny.SetOf(p0, _32_v8, (_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool), (_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool))
		var _42_v17 _dafny.Map
		_ = _42_v17
		_42_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool), _41_v16)
		_41_v16 = (Companion_Default___.Fm24(globalState)).Intersection((func() _dafny.Set {
			if (_42_v17).Contains(p2) {
				return (_42_v17).Get(p2).(_dafny.Set)
			}
			return _41_v16
		})())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
		_ = _index7
		(_23_v3).ArraySet1(p0, (_index7).Int())
		var _43_v18 _dafny.Array
		_ = _43_v18
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
		_ = _nw5
		_43_v18 = _nw5
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_43_v18), 0))
		_ = _index8
		(_43_v18).ArraySet1(_23_v3, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_43_v18), 0))
		_ = _index9
		(_43_v18).ArraySet1(_23_v3, (_index9).Int())
		if !((_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool)) || ((_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool)) {
			var _44_v19 _dafny.Array
			_ = _44_v19
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_3
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Map = (func(_45_v7 _dafny.Int, _46_v3 _dafny.Array, _47_p2 bool) func(_dafny.Int) _dafny.Map {
					return func(_48_i4 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v7, (_46_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_46_v3), 0))).Int()).(bool))).Cardinality(), _47_p2)
					}
				})(_31_v7, _23_v3, p2)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw6 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw6).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw6).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_44_v19 = _nw6
			var _49_v20 _dafny.Map
			_ = _49_v20
			_49_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int), p0)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_44_v19), 0))
			_ = _index10
			(_44_v19).ArraySet1(_49_v20, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_44_v19), 0))
			_ = _index11
			(_44_v19).ArraySet1((((_49_v20).Update((_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int), p2)).Update((_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int), !((_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool)))).Merge(_49_v20), (_index11).Int())
			var _50_v21 _dafny.CodePoint
			_ = _50_v21
			_50_v21 = _dafny.CodePoint('r')
			var _51_v22 *C0
			_ = _51_v22
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__((_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int), _50_v21, p0, (_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool))
			_51_v22 = _nw7
			var _52_v23 _dafny.Sequence
			_ = _52_v23
			_52_v23 = _dafny.SeqOf(_51_v22, _51_v22, _51_v22, _51_v22)
			var _53_v24 _dafny.Map
			_ = _53_v24
			_53_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool), p0)
			var _54_v25 _dafny.Sequence
			_ = _54_v25
			_54_v25 = _dafny.SeqOf((_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool), !_dafny.Companion_Sequence_.Contains(_52_v23, _51_v22), (func() bool {
				if (_53_v24).Contains(p2) {
					return (_53_v24).Get(p2).(bool)
				}
				return (_51_v22).Fm5(_32_v8, globalState)
			})())
			_54_v25 = _54_v25
			var _55_v27 D5
			_ = _55_v27
			_55_v27 = Companion_D5_.Create_DC11_()
			var _56_v28 _dafny.MultiSet
			_ = _56_v28
			_56_v28 = _dafny.MultiSetOf(_55_v27)
			var _57_v29 _dafny.Map
			_ = _57_v29
			_57_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _56_v28)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
			_ = _index12
			var _rhs5 _dafny.Int = (((func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_58_v22 *C0) func(_dafny.Int) _dafny.Int {
					return func(_59_i5 _dafny.Int) _dafny.Int {
						return _58_v22.F13
					}
				})(_51_v22))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dy")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}((func(_60_v22 *C0) func(_dafny.Int) _dafny.Int {
					return func(_61_i5 _dafny.Int) _dafny.Int {
						return _60_v22.F13
					}
				})(_51_v22)))).Cardinality()))).Uint32(), _51_v22.F13)).Elements()); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _62_v26 _dafny.Int
					_62_v26 = interface{}(_compr_9).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}((func(_63_v22 *C0) func(_dafny.Int) _dafny.Int {
						return func(_59_i5 _dafny.Int) _dafny.Int {
							return _63_v22.F13
						}
					})(_51_v22))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dy")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg13 _dafny.Int) interface{} {
							return coer13(arg13)
						}
					}((func(_64_v22 *C0) func(_dafny.Int) _dafny.Int {
						return func(_61_i5 _dafny.Int) _dafny.Int {
							return _64_v22.F13
						}
					})(_51_v22)))).Cardinality()))).Uint32(), _51_v22.F13), _62_v26) {
						_coll9.Add((_62_v26).Times(_31_v7), _dafny.MultiSetOf(Companion_D5_.Create_DC11_()))
					}
				}
				return _coll9.ToMap()
			}()).Update(_dafny.IntOfInt64(832), (func() _dafny.MultiSet {
				if (_57_v29).Contains(_32_v8) {
					return (_57_v29).Get(_32_v8).(_dafny.MultiSet)
				}
				return _56_v28
			})())).Update((_dafny.Zero).Minus((_51_v22.F13).Times((_dafny.Zero).Minus(_31_v7))), _56_v28)).Cardinality()
			_ = _rhs5
			var _rhs6 bool = _32_v8
			_ = _rhs6
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			var _lhs8 _dafny.Array = _23_v3
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
			_ = _lhs9
			_lhs7.F5 = _rhs5
			(_lhs8).ArraySet1(_rhs6, (_lhs9).Int())
			_32_v8 = p0
			var _65_v30 D1
			_ = _65_v30
			_65_v30 = Companion_D1_.Create_DC1_(_dafny.Companion_Sequence_.Update((_21_v1).F7(), (Companion_Default___.SafeIndex(_31_v7, _dafny.IntOfUint32(((_21_v1).F7()).Cardinality()))).Uint32(), _50_v21))
			(globalState).F5 = _dafny.IntOfUint32((_dafny.SeqOf(_65_v30, _65_v30)).Cardinality())
		} else {
			_20_v0 = (_21_v1).F7()
			_20_v0 = (_21_v1).F7()
			var _66_v31 _dafny.Array
			_ = _66_v31
			var _len0_4 _dafny.Int = _dafny.One
			_ = _len0_4
			var _nw8 _dafny.Array
			_ = _nw8
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw8 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_67_v1 *C4) func(_dafny.Int) _dafny.Sequence {
					return func(_68_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate((_67_v1).F7(), (_67_v1).F7())
					}
				})(_21_v1)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw8 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw8).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw8).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_66_v31 = _nw8
			var _69_v32 _dafny.Sequence
			_ = _69_v32
			_69_v32 = _dafny.SeqOf(_66_v31, _66_v31, _66_v31, _66_v31)
			_66_v31 = (_69_v32).Select((Companion_Default___.SafeIndex(_31_v7, _dafny.IntOfUint32((_69_v32).Cardinality()))).Uint32()).(_dafny.Array)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
			_ = _index13
			(_23_v3).ArraySet1(p2, (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
			_ = _index14
			(_23_v3).ArraySet1(_32_v8, (_index14).Int())
		}
		(globalState).F5 = _31_v7
	}
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))
	_ = _index15
	(_23_v3).ArraySet1((_31_v7).Cmp(_31_v7) > 0, (_index15).Int())
	var _70_i7 _dafny.Int
	_ = _70_i7
	_70_i7 = _dafny.Zero
	{
		for p2 {
			{
				if (_70_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_70_i7 = (_70_i7).Plus(_dafny.One)
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))
				_ = _index16
				(_29_v6).ArraySet1(Companion_Default___.Fm0((func() bool {
					if p0 {
						return p0
					}
					return p0
				})(), (_21_v1).F7(), globalState), (_index16).Int())
				var _71_v33 _dafny.CodePoint
				_ = _71_v33
				_71_v33 = _dafny.CodePoint('r')
				var _72_v34 _dafny.Map
				_ = _72_v34
				_72_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm23(p0, _32_v8, _31_v7, globalState), _71_v33)
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))
				_ = _index17
				(_29_v6).ArraySet1(((_72_v34).Cardinality()).Minus((_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int)), (_index17).Int())
				(globalState).F5 = Companion_Default___.SafeModuloInt((_31_v7).Plus((_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int)), (_29_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_29_v6), 0))).Int()).(_dafny.Int))
				var _73_v35 *C3
				_ = _73_v35
				var _nw9 *C3 = New_C3_()
				_ = _nw9
				_nw9.Ctor__()
				_73_v35 = _nw9
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _74_v36 _dafny.Sequence
	_ = _74_v36
	_74_v36 = _dafny.SeqOf(_23_v3, _23_v3)
	var _75_v37 _dafny.CodePoint
	_ = _75_v37
	_75_v37 = _dafny.CodePoint('e')
	var _76_v38 _dafny.Sequence
	_ = _76_v38
	_76_v38 = _dafny.SeqOf((_21_v1).F7(), _20_v0, _20_v0, _dafny.Companion_Sequence_.Update((_21_v1).F7(), (Companion_Default___.SafeIndex(_31_v7, _dafny.IntOfUint32(((_21_v1).F7()).Cardinality()))).Uint32(), _75_v37), (_21_v1).F7())
	var _77_v39 D4
	_ = _77_v39
	_77_v39 = Companion_D4_.Create_DC9_(p0, _32_v8)
	var _78_v40 _dafny.Sequence
	_ = _78_v40
	_78_v40 = _dafny.SeqOf((_74_v36).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm16(_dafny.IntOfUint32((_76_v38).Cardinality()), (_23_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_23_v3), 0))).Int()).(bool), (_77_v39).Dtor_cf21(), true, globalState)).Cardinality()), _dafny.IntOfUint32((_74_v36).Cardinality()))).Uint32()).(_dafny.Array), _23_v3, _23_v3)
	_78_v40 = _78_v40
	r0 = _31_v7
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _79_v0 _dafny.Int
	_ = _79_v0
	_79_v0 = _dafny.IntOfInt64(717)
	var _80_v1 _dafny.Sequence
	_ = _80_v1
	_80_v1 = _dafny.SeqOf(_79_v0)
	var _81_v2 _dafny.Map
	_ = _81_v2
	_81_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v0, _80_v1)
	var _82_v3 _dafny.Array
	_ = _82_v3
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
	_ = _nw10
	_82_v3 = _nw10
	var _83_globalState *GlobalState
	_ = _83_globalState
	var _nw11 *GlobalState = New_GlobalState_()
	_ = _nw11
	_nw11.Ctor__(_81_v2, false, false, _82_v3, true, _dafny.IntOfInt64(341), false)
	_83_globalState = _nw11
	var _84_v4 _dafny.Sequence
	_ = _84_v4
	_84_v4 = _dafny.UnicodeSeqOfUtf8Bytes("qhbqusrq")
	_84_v4 = _84_v4
	var _85_v5 _dafny.Map
	_ = _85_v5
	_85_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v4, _84_v4)
	var _86_v6 _dafny.Array
	_ = _86_v6
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(11)
	_ = _len0_5
	var _nw12 _dafny.Array
	_ = _nw12
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw12 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Int = (func(_87_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_88_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_88_i1, _87_v0)
			}
		})(_79_v0)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw12 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw12).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw12).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_86_v6 = _nw12
	var _89_v7 bool
	_ = _89_v7
	_89_v7 = true
	var _90_v8 _dafny.Map
	_ = _90_v8
	_90_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v7, _79_v0)
	var _91_v9 _dafny.Array
	_ = _91_v9
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(22)
	_ = _len0_6
	var _nw13 _dafny.Array
	_ = _nw13
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw13 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) _dafny.CodePoint = func(_92_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw13 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw13).ArraySet1CodePoint(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw13).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_91_v9 = _nw13
	var _93_v10 D0
	_ = _93_v10
	_93_v10 = Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v4, _84_v4), _86_v6, _90_v8, _91_v9, _86_v6)
	var _94_v11 _dafny.Map
	_ = _94_v11
	_94_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v7, (_dafny.Zero).Minus((_dafny.SetOf(Companion_D0_.Create_DC0_(_85_v5, _86_v6, _90_v8, _91_v9, _86_v6), _93_v10)).Cardinality()))
	var _95_v12 D0
	_ = _95_v12
	_95_v12 = Companion_D0_.Create_DC0_((_85_v5).Update(_84_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-417))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_96_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))), _86_v6, (func() _dafny.Map {
		if _89_v7 {
			return _90_v8
		}
		return _94_v11
	})(), _91_v9, _86_v6)
	var _source1 D0 = _95_v12
	_ = _source1
	var _97___mcc_h0 _dafny.Map = _source1.Get_().(D0_DC0).Cf0
	_ = _97___mcc_h0
	var _98___mcc_h1 _dafny.Array = _source1.Get_().(D0_DC0).Cf1
	_ = _98___mcc_h1
	var _99___mcc_h2 _dafny.Map = _source1.Get_().(D0_DC0).Cf2
	_ = _99___mcc_h2
	var _100___mcc_h3 _dafny.Array = _source1.Get_().(D0_DC0).Cf3
	_ = _100___mcc_h3
	var _101___mcc_h4 _dafny.Array = _source1.Get_().(D0_DC0).Cf4
	_ = _101___mcc_h4
	var _102_cf4 _dafny.Array = _101___mcc_h4
	_ = _102_cf4
	var _103_cf3 _dafny.Array = _100___mcc_h3
	_ = _103_cf3
	var _104_cf2 _dafny.Map = _99___mcc_h2
	_ = _104_cf2
	var _105_cf1 _dafny.Array = _98___mcc_h1
	_ = _105_cf1
	var _106_cf0 _dafny.Map = _97___mcc_h0
	_ = _106_cf0
	var _107_v13 _dafny.Set
	_ = _107_v13
	_107_v13 = _dafny.SetOf(_104_cf2, _94_v11)
	var _108_v14 _dafny.Int
	_ = _108_v14
	var _out0 _dafny.Int
	_ = _out0
	_out0 = Companion_Default___.M0(_89_v7, _107_v13, (_79_v0).Cmp(_79_v0) == 0, _83_globalState)
	_108_v14 = _out0
	_84_v4 = _84_v4
	var _109_v15 _dafny.MultiSet
	_ = _109_v15
	_109_v15 = _dafny.MultiSetOf(!(_89_v7))
	var _source2 D0 = Companion_D0_.Create_DC0_(_106_cf0, _86_v6, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_109_v15).Cardinality()), _103_cf3, _102_cf4)
	_ = _source2
	var _110___mcc_h5 _dafny.Map = _source2.Get_().(D0_DC0).Cf0
	_ = _110___mcc_h5
	var _111___mcc_h6 _dafny.Array = _source2.Get_().(D0_DC0).Cf1
	_ = _111___mcc_h6
	var _112___mcc_h7 _dafny.Map = _source2.Get_().(D0_DC0).Cf2
	_ = _112___mcc_h7
	var _113___mcc_h8 _dafny.Array = _source2.Get_().(D0_DC0).Cf3
	_ = _113___mcc_h8
	var _114___mcc_h9 _dafny.Array = _source2.Get_().(D0_DC0).Cf4
	_ = _114___mcc_h9
	var _115_cf4 _dafny.Array = _114___mcc_h9
	_ = _115_cf4
	var _116_cf3 _dafny.Array = _113___mcc_h8
	_ = _116_cf3
	var _117_cf2 _dafny.Map = _112___mcc_h7
	_ = _117_cf2
	var _118_cf1 _dafny.Array = _111___mcc_h6
	_ = _118_cf1
	var _119_cf0 _dafny.Map = _110___mcc_h5
	_ = _119_cf0
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_118_cf1), 0))
	_ = _index18
	(_118_cf1).ArraySet1(_108_v14, (_index18).Int())
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_118_cf1), 0))
	_ = _index19
	(_118_cf1).ArraySet1(((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(687))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_120_i3 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("nvgx")
	}))).Cardinality())).Minus(Companion_Default___.Fm0(_89_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-139))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_121_i4 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	})), _83_globalState))).Times(_79_v0), (_index19).Int())
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_118_cf1), 0))
	_ = _index20
	(_118_cf1).ArraySet1((_118_cf1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_118_cf1), 0))).Int()).(_dafny.Int), (_index20).Int())
	var _122_v16 _dafny.Int
	_ = _122_v16
	var _out1 _dafny.Int
	_ = _out1
	_out1 = Companion_Default___.M0(!(_89_v7), _107_v13, (_89_v7) || (_89_v7), _83_globalState)
	_122_v16 = _out1
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_118_cf1), 0))
	_ = _index21
	(_118_cf1).ArraySet1(_79_v0, (_index21).Int())
	var _123_v17 _dafny.Map
	_ = _123_v17
	_123_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_cf2, _dafny.IntOfUint32((_84_v4).Cardinality()))
	var _124_v19 _dafny.Int
	_ = _124_v19
	var _out2 _dafny.Int
	_ = _out2
	_out2 = Companion_Default___.M0(_89_v7, (func() _dafny.Set {
		if _89_v7 {
			return _107_v13
		}
		return func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_123_v17).Keys().Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _125_v18 _dafny.Map
				_125_v18 = interface{}(_compr_10).(_dafny.Map)
				if (_123_v17).Contains(_125_v18) {
					_coll10.Add(_125_v18)
				}
			}
			return _coll10.ToSet()
		}()
	})(), !(_89_v7), _83_globalState)
	_124_v19 = _out2
	_89_v7 = _89_v7
	_79_v0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_84_v4, _84_v4)).Cardinality())
	if _89_v7 {
		_89_v7 = _89_v7
		var _126_v20 _dafny.Set
		_ = _126_v20
		_126_v20 = _dafny.SetOf(_94_v11)
		var _127_v21 _dafny.Int
		_ = _127_v21
		var _out3 _dafny.Int
		_ = _out3
		_out3 = Companion_Default___.M0(!(_89_v7) || (_89_v7), _126_v20, _89_v7, _83_globalState)
		_127_v21 = _out3
		var _128_v22 _dafny.CodePoint
		_ = _128_v22
		_128_v22 = _dafny.CodePoint('l')
		var _129_v23 _dafny.Set
		_ = _129_v23
		_129_v23 = _dafny.SetOf(_128_v22, _128_v22)
		var _130_v24 _dafny.MultiSet
		_ = _130_v24
		_130_v24 = _dafny.MultiSetOf(_80_v1)
		var _131_v25 *C1
		_ = _131_v25
		var _nw14 *C1 = New_C1_()
		_ = _nw14
		_nw14.Ctor__(_130_v24, _89_v7, _128_v22, _89_v7, _89_v7)
		_131_v25 = _nw14
		var _132_v26 D4
		_ = _132_v26
		_132_v26 = Companion_D4_.Create_DC8_((_129_v23).Cardinality(), !(_89_v7), _127_v21, _131_v25)
		var _133_v27 _dafny.Map
		_ = _133_v27
		_133_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v0, _89_v7)
		var _134_v28 D1
		_ = _134_v28
		_134_v28 = Companion_D1_.Create_DC2_(_79_v0, (_133_v27).Cardinality(), _128_v22)
		var _135_v29 *C0
		_ = _135_v29
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__(Companion_Default___.Fm0((_132_v26).Dtor_cf17(), _84_v4, _83_globalState), (_134_v28).Dtor_cf8(), (_131_v25).F12(), (_131_v25).F12())
		_135_v29 = _nw15
		var _136_v30 _dafny.Array
		_ = _136_v30
		var _nw16 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(22))
		_ = _nw16
		_136_v30 = _nw16
		var _137_v31 _dafny.Sequence
		_ = _137_v31
		_137_v31 = _dafny.SeqOf((_131_v25).F12())
		var _138_v32 T1
		_ = _138_v32
		var _nw17 *C1 = New_C1_()
		_ = _nw17
		_nw17.Ctor__((_131_v25).F11(), _89_v7, _128_v22, (_137_v31).Select((Companion_Default___.SafeIndex(_79_v0, _dafny.IntOfUint32((_137_v31).Cardinality()))).Uint32()).(bool), _89_v7)
		_138_v32 = _nw17
		var _139_v33 _dafny.Sequence
		_ = _139_v33
		_139_v33 = _dafny.SeqOf(_138_v32)
		var _140_v34 _dafny.Map
		_ = _140_v34
		_140_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v33, _135_v29.F13)
		var _141_v35 _dafny.Map
		_ = _141_v35
		_141_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(441), _79_v0)
		var _142_v36 _dafny.Set
		_ = _142_v36
		_142_v36 = _dafny.SetOf(_138_v32.F9(), (_131_v25).F12())
		var _143_v37 D2
		_ = _143_v37
		_143_v37 = Companion_D2_.Create_DC4_(_84_v4, _137_v31, (func() _dafny.Int {
			if (_140_v34).Contains(_139_v33) {
				return (_140_v34).Get(_139_v33).(_dafny.Int)
			}
			return (func() _dafny.Int {
				if (_141_v35).Contains(_135_v29.F13) {
					return (_141_v35).Get(_135_v29.F13).(_dafny.Int)
				}
				return (_142_v36).Cardinality()
			})()
		})())
		var _144_v38 D2
		_ = _144_v38
		_144_v38 = Companion_D2_.Create_DC5_(_143_v37)
		var _145_v39 D2
		_ = _145_v39
		_145_v39 = Companion_D2_.Create_DC5_(_143_v37)
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_136_v30), 0))
		_ = _index22
		(_136_v30).ArraySet1(_145_v39, (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_136_v30), 0))
		_ = _index23
		(_136_v30).ArraySet1(_145_v39, (_index23).Int())
		var _146_v40 _dafny.Map
		_ = _146_v40
		_146_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v7, (_131_v25).F12())
		var _147_v41 *C2
		_ = _147_v41
		var _nw18 *C2 = New_C2_()
		_ = _nw18
		_nw18.Ctor__((_146_v40).Cardinality(), !((_131_v25).F12()), _138_v32.F9())
		_147_v41 = _nw18
		var _148_v42 _dafny.Array
		_ = _148_v42
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(22))
		_ = _nw19
		_148_v42 = _nw19
		var _rhs7 _dafny.Sequence = _84_v4
		_ = _rhs7
		var _rhs8 bool = (_131_v25).F12()
		_ = _rhs8
		var _rhs9 _dafny.Sequence = Companion_Default___.Fm21((func() _dafny.Int {
			if (_131_v25).F12() {
				return _127_v21
			}
			return _dafny.IntOfInt64(-1)
		})(), _dafny.Companion_Sequence_.Concatenate(_84_v4, _84_v4), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(109))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_149_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), (_147_v41).F14(), _83_globalState)
		_ = _rhs9
		var _rhs10 *C2 = (func() *C2 {
			if true {
				return _147_v41
			}
			return _147_v41
		})()
		_ = _rhs10
		var _rhs11 _dafny.Array = _148_v42
		_ = _rhs11
		var _lhs10 T1 = _138_v32
		_ = _lhs10
		_84_v4 = _rhs7
		_lhs10.F9_set_(_rhs8)
		_84_v4 = _rhs9
		_147_v41 = _rhs10
		_148_v42 = _rhs11
	} else {
		var _150_v43 _dafny.MultiSet
		_ = _150_v43
		_150_v43 = _dafny.MultiSetOf(_89_v7, _89_v7, _89_v7)
		var _151_v44 D2
		_ = _151_v44
		_151_v44 = Companion_D2_.Create_DC3_(false)
		var _152_v45 _dafny.Int
		_ = _152_v45
		var _out4 _dafny.Int
		_ = _out4
		_out4 = Companion_Default___.M0(!((_79_v0).Cmp(_79_v0) >= 0), _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v7, _79_v0)), Companion_Default___.Fm22((_150_v43).Cardinality(), _89_v7, !(true), (_151_v44).Dtor_cf9(), _83_globalState), _83_globalState)
		_152_v45 = _out4
		var _153_v46 _dafny.MultiSet
		_ = _153_v46
		_153_v46 = _dafny.MultiSetOf(_80_v1)
		var _154_v47 _dafny.CodePoint
		_ = _154_v47
		_154_v47 = _dafny.CodePoint('q')
		var _155_v48 *C1
		_ = _155_v48
		var _nw20 *C1 = New_C1_()
		_ = _nw20
		_nw20.Ctor__(_153_v46, _89_v7, _154_v47, _89_v7, _89_v7)
		_155_v48 = _nw20
		var _156_v49 D4
		_ = _156_v49
		_156_v49 = Companion_D4_.Create_DC8_(_dafny.IntOfInt64(984), _89_v7, _152_v45, _155_v48)
		var _157_v50 _dafny.Sequence
		_ = _157_v50
		_157_v50 = _dafny.SeqOf(Companion_Default___.Fm22((_156_v49).Dtor_cf18(), (_155_v48).F12(), _89_v7, true, _83_globalState), (_155_v48).F12(), (_155_v48).F12(), (_155_v48).F12(), (_155_v48).F12())
		var _158_v51 _dafny.Map
		_ = _158_v51
		_158_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v7, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_157_v50, _157_v50), (Companion_Default___.SafeIndex(_79_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_157_v50, _157_v50)).Cardinality()))).Uint32(), (_155_v48).F12()))
		_158_v51 = _158_v51
		_89_v7 = false
		_155_v48 = (func() *C1 {
			if _89_v7 {
				return _155_v48
			}
			return _155_v48
		})()
		_89_v7 = (_155_v48).F12()
	}
	var _159_v52 _dafny.CodePoint
	_ = _159_v52
	_159_v52 = _dafny.CodePoint('y')
	_159_v52 = Companion_Default___.Fm13(_79_v0, _83_globalState)
	_89_v7 = _89_v7
	_89_v7 = _89_v7
	var _160_v53 *C2
	_ = _160_v53
	var _nw21 *C2 = New_C2_()
	_ = _nw21
	_nw21.Ctor__(_79_v0, _89_v7, _89_v7)
	_160_v53 = _nw21
	var _161_v54 _dafny.Sequence
	_ = _161_v54
	_161_v54 = _dafny.SeqOf(Companion_D1_.Create_DC1_(_84_v4))
	var _162_v55 _dafny.Sequence
	_ = _162_v55
	_162_v55 = _dafny.SeqOf(_161_v54, _161_v54)
	var _rhs12 bool = _89_v7
	_ = _rhs12
	var _rhs13 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_161_v54, (_162_v55).Select((Companion_Default___.SafeIndex((_160_v53).F14(), _dafny.IntOfUint32((_162_v55).Cardinality()))).Uint32()).(_dafny.Sequence)), _161_v54)
	_ = _rhs13
	var _rhs14 bool = _89_v7
	_ = _rhs14
	_89_v7 = _rhs12
	_161_v54 = _rhs13
	_89_v7 = _rhs14
	var _163_v56 _dafny.MultiSet
	_ = _163_v56
	_163_v56 = _dafny.MultiSetOf(_80_v1)
	var _164_v57 T1
	_ = _164_v57
	var _nw22 *C1 = New_C1_()
	_ = _nw22
	_nw22.Ctor__(_163_v56, _89_v7, _159_v52, _89_v7, _89_v7)
	_164_v57 = _nw22
	var _165_v58 _dafny.Sequence
	_ = _165_v58
	_165_v58 = _dafny.SeqOf(_164_v57, _164_v57, _164_v57, _164_v57, _164_v57)
	var _166_v59 _dafny.Array
	_ = _166_v59
	var _nwElement0_0 T1 = (func() T1 {
		if _89_v7 {
			return _164_v57
		}
		return _164_v57
	})()
	_ = _nwElement0_0
	var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(10))
	_ = _nw23
	(_nw23).ArraySet1(_nwElement0_0, 0)
	(_nw23).ArraySet1(_164_v57, 1)
	(_nw23).ArraySet1(_164_v57, 2)
	(_nw23).ArraySet1((func() T1 {
		if _89_v7 {
			return _164_v57
		}
		return _164_v57
	})(), 3)
	(_nw23).ArraySet1(_164_v57, 4)
	(_nw23).ArraySet1((_165_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_84_v4).Cardinality()), _dafny.IntOfUint32((_165_v58).Cardinality()))).Uint32()).(T1), 5)
	(_nw23).ArraySet1(_164_v57, 6)
	(_nw23).ArraySet1(_164_v57, 7)
	(_nw23).ArraySet1(_164_v57, 8)
	(_nw23).ArraySet1(_164_v57, 9)
	_166_v59 = _nw23
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_166_v59), 0))
	_ = _index24
	(_166_v59).ArraySet1(_164_v57, (_index24).Int())
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_166_v59), 0))
	_ = _index25
	(_166_v59).ArraySet1(_164_v57, (_index25).Int())
	var _167_v60 _dafny.MultiSet
	_ = _167_v60
	_167_v60 = _dafny.MultiSetOf(_79_v0, _79_v0)
	var _168_v61 _dafny.Set
	_ = _168_v61
	_168_v61 = _dafny.SetOf((_160_v53).F14())
	if (((_167_v60).Cardinality()).Cmp(_79_v0) != 0) || ((_168_v61).IsDisjointFrom(_168_v61)) {
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_86_v6), 0))
		_ = _index26
		(_86_v6).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_79_v0, (_160_v53).F14())), (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_86_v6), 0))
		_ = _index27
		(_86_v6).ArraySet1(Companion_Default___.SafeModuloInt((_160_v53).F14(), (_79_v0).Plus(_dafny.IntOfInt64(71))), (_index27).Int())
		var _169_v62 _dafny.Array
		_ = _169_v62
		var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw24
		_169_v62 = _nw24
		_169_v62 = _169_v62
		_89_v7 = _89_v7
		(_83_globalState).F5 = (Companion_D1_.Create_DC2_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v62, _164_v57.F9())).Cardinality()), (_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int), _164_v57.F10())).Dtor_cf7()
		var _170_v63 _dafny.Sequence
		_ = _170_v63
		_170_v63 = _dafny.SeqOf(_89_v7)
		if (_170_v63).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.IntOfUint32((_170_v63).Cardinality()))).Uint32()).(bool) {
			_84_v4 = _84_v4
			_89_v7 = (_170_v63).Select((Companion_Default___.SafeIndex((_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_170_v63).Cardinality()))).Uint32()).(bool)
			var _171_v64 _dafny.Array
			_ = _171_v64
			var _nw25 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw25
			_171_v64 = _nw25
			var _172_v65 *C1
			_ = _172_v65
			var _nw26 *C1 = New_C1_()
			_ = _nw26
			_nw26.Ctor__(_163_v56, !(!(_164_v57.F9())), _164_v57.F10(), !((_164_v57).F8()), _164_v57.F9())
			_172_v65 = _nw26
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_171_v64), 0))
			_ = _index28
			(_171_v64).ArraySet1(_172_v65, (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_171_v64), 0))
			_ = _index29
			var _rhs15 _dafny.Array = _169_v62
			_ = _rhs15
			var _rhs16 *C1 = _172_v65
			_ = _rhs16
			var _lhs11 _dafny.Array = _171_v64
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_171_v64), 0))
			_ = _lhs12
			_169_v62 = _rhs15
			(_lhs11).ArraySet1(_rhs16, (_lhs12).Int())
			var _173_v66 _dafny.Map
			_ = _173_v66
			_173_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_84_v4, _84_v4), _dafny.IntOfUint32((_170_v63).Cardinality()))
			_173_v66 = (_173_v66).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_174_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_175_i6 _dafny.Int) _dafny.CodePoint {
					return _174_v52
				}
			})(_159_v52))), (Companion_Default___.SafeIndex((_160_v53).F14(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_176_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_177_i6 _dafny.Int) _dafny.CodePoint {
					return _176_v52
				}
			})(_159_v52)))).Cardinality()))).Uint32(), _164_v57.F10()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(996))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_178_v57 T1) func(_dafny.Int) _dafny.CodePoint {
				return func(_179_i7 _dafny.Int) _dafny.CodePoint {
					return _178_v57.F10()
				}
			})(_164_v57)))), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(983))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_180_v57 T1) func(_dafny.Int) _dafny.CodePoint {
				return func(_181_i8 _dafny.Int) _dafny.CodePoint {
					return _180_v57.F10()
				}
			})(_164_v57)))).Cardinality())).Times((_160_v53).F14()))
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_82_v3), 0))
			_ = _index30
			(_82_v3).ArraySet1(!((_164_v57).F8()), (_index30).Int())
			var _182_v67 _dafny.Set
			_ = _182_v67
			_182_v67 = _dafny.SetOf((_164_v57).F8(), (_172_v65).F12())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_82_v3), 0))
			_ = _index31
			(_82_v3).ArraySet1((_164_v57).Fm7((_dafny.IntOfUint32((_170_v63).Cardinality())).Cmp((_160_v53).F14()) > 0, _182_v67, _84_v4, _83_globalState), (_index31).Int())
		} else {
			var _183_v68 _dafny.Set
			_ = _183_v68
			_183_v68 = _dafny.SetOf(_94_v11)
			var _184_v69 _dafny.Int
			_ = _184_v69
			var _out5 _dafny.Int
			_ = _out5
			_out5 = Companion_Default___.M0(!((_164_v57).Fm5((_164_v57).F8(), _83_globalState)), _183_v68, true, _83_globalState)
			_184_v69 = _out5
			(_164_v57).F9_set_((func() bool {
				if !((_164_v57).F8()) {
					return (_164_v57).F8()
				}
				return !(_164_v57.F9()) || ((_164_v57).F8())
			})())
			var _185_v70 T0
			_ = _185_v70
			var _nw27 *C2 = New_C2_()
			_ = _nw27
			_nw27.Ctor__((_160_v53).F14(), _164_v57.F9(), false)
			_185_v70 = _nw27
			var _186_v71 _dafny.Sequence
			_ = _186_v71
			_186_v71 = _dafny.SeqOf(_185_v70, _185_v70)
			(_83_globalState).F5 = (func() _dafny.Int {
				if (_167_v60).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if (_167_v60).Contains(_dafny.IntOfUint32((_186_v71).Cardinality())) {
						return (_167_v60).Multiplicity(_dafny.IntOfUint32((_186_v71).Cardinality()))
					}
					return (_160_v53).F14()
				})()))) {
					return (_167_v60).Multiplicity((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
						if (_167_v60).Contains(_dafny.IntOfUint32((_186_v71).Cardinality())) {
							return (_167_v60).Multiplicity(_dafny.IntOfUint32((_186_v71).Cardinality()))
						}
						return (_160_v53).F14()
					})())))
				}
				return Companion_Default___.SafeModuloInt(_184_v69, (_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int))
			})()
			(_164_v57).F10_set_((_84_v4).Select((Companion_Default___.SafeIndex(_184_v69, _dafny.IntOfUint32((_84_v4).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_168_v61 = (Companion_Default___.Fm20((_185_v70).F8(), _83_globalState)).Union(_168_v61)
		}
	} else {
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))
		_ = _index32
		(_86_v6).ArraySet1((_160_v53).F14(), (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))
		_ = _index33
		(_86_v6).ArraySet1((_164_v57).Fm6(_164_v57.F9(), _83_globalState), (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))
		_ = _index34
		(_82_v3).ArraySet1(_89_v7, (_index34).Int())
		var _187_v72 D4
		_ = _187_v72
		_187_v72 = Companion_D4_.Create_DC9_((_164_v57).F8(), _164_v57.F9())
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))
		_ = _index35
		(_82_v3).ArraySet1(!((false) && ((_187_v72).Dtor_cf21())), (_index35).Int())
		var _188_v73 _dafny.Sequence
		_ = _188_v73
		_188_v73 = _dafny.SeqOf(Companion_Default___.Fm22((_160_v53).F14(), (_82_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))).Int()).(bool), _164_v57.F9(), _89_v7, _83_globalState))
		if (_167_v60).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfUint32((_188_v73).Cardinality()), (_160_v53).F14())) {
			var _189_v74 _dafny.Map
			_ = _189_v74
			_189_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_160_v53).F14(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_164_v57).Fm6((_164_v57).F8(), _83_globalState), (_160_v53).F14())).Update((_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int), (_160_v53).F14())).Cardinality())
			(_164_v57).F9_set_((Companion_Default___.SafeDivisionInt((_160_v53).F14(), (_189_v74).Cardinality())).Cmp(((_160_v53).F14()).Plus((_160_v53).F14())) < 0)
			var _190_v75 *C1
			_ = _190_v75
			var _nw28 *C1 = New_C1_()
			_ = _nw28
			_nw28.Ctor__((_163_v56).Difference(_163_v56), (_164_v57).F8(), _164_v57.F10(), (_164_v57).F8(), (_160_v53).Fm5((_164_v57).F8(), _83_globalState))
			_190_v75 = _nw28
			var _191_v76 D4
			_ = _191_v76
			_191_v76 = Companion_D4_.Create_DC8_((_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int), (_164_v57).F8(), (_160_v53).F14(), _190_v75)
			var _pat_let_tv0 = _86_v6
			_ = _pat_let_tv0
			var _pat_let_tv1 = _86_v6
			_ = _pat_let_tv1
			var _pat_let_tv2 = _164_v57
			_ = _pat_let_tv2
			var _nw29 *C1 = New_C1_()
			_ = _nw29
			_nw29.Ctor__((_190_v75).F11(), _89_v7, _159_v52, (func(_pat_let0_0 D4) D4 {
				return func(_192_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let1_0 _dafny.Int) D4 {
						return func(_193_dt__update_hcf16_h0 _dafny.Int) D4 {
							return func(_pat_let2_0 bool) D4 {
								return func(_194_dt__update_hcf17_h0 bool) D4 {
									return Companion_D4_.Create_DC8_(_193_dt__update_hcf16_h0, _194_dt__update_hcf17_h0, (_192_dt__update__tmp_h0).Dtor_cf18(), (_192_dt__update__tmp_h0).Dtor_cf19())
								}(_pat_let2_0)
							}(_pat_let_tv2.F9())
						}(_pat_let1_0)
					}((_pat_let_tv1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_pat_let_tv0), 0))).Int()).(_dafny.Int))
				}(_pat_let0_0)
			}(_191_v76)).Dtor_cf17(), (_164_v57).F8())
			_190_v75 = _nw29
			_167_v60 = _167_v60
			var _195_v77 _dafny.Set
			_ = _195_v77
			_195_v77 = _dafny.SetOf((_190_v75).F12(), _164_v57.F9(), (_164_v57).F8(), _164_v57.F9(), !(_164_v57.F9()))
			var _196_v78 _dafny.Array
			_ = _196_v78
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_7
			var _nw30 _dafny.Array
			_ = _nw30
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw30 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = (func(_197_v3 _dafny.Array) func(_dafny.Int) bool {
					return func(_198_i9 _dafny.Int) bool {
						return (_197_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_197_v3), 0))).Int()).(bool)
					}
				})(_82_v3)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw30 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw30).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw30).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_196_v78 = _nw30
			var _199_v79 _dafny.Map
			_ = _199_v79
			_199_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v4, (_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int))
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))
			_ = _index36
			var _rhs17 _dafny.Int = ((_160_v53).F14()).Plus((_167_v60).Cardinality())
			_ = _rhs17
			var _rhs18 _dafny.Set = _dafny.SetOf(_dafny.Companion_Sequence_.Equal(_84_v4, _dafny.SeqOf(_164_v57.F10(), _dafny.CodePoint('h'))), !((_82_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))).Int()).(bool)), !(false) || (_164_v57.F9()))
			_ = _rhs18
			var _rhs19 _dafny.Array = _196_v78
			_ = _rhs19
			var _rhs20 _dafny.Int = _dafny.IntOfUint32((_84_v4).Cardinality())
			_ = _rhs20
			var _rhs21 _dafny.Int = (func() _dafny.Int {
				if (_199_v79).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sgw"), _84_v4)) {
					return (_199_v79).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sgw"), _84_v4)).(_dafny.Int)
				}
				return (_160_v53).F14()
			})()
			_ = _rhs21
			var _lhs13 _dafny.Array = _86_v6
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))
			_ = _lhs14
			var _lhs15 *GlobalState = _83_globalState
			_ = _lhs15
			(_lhs13).ArraySet1(_rhs17, (_lhs14).Int())
			_195_v77 = _rhs18
			_196_v78 = _rhs19
			_79_v0 = _rhs20
			_lhs15.F5 = _rhs21
			_79_v0 = Companion_Default___.SafeDivisionInt((_168_v61).Cardinality(), _dafny.IntOfInt64(59))
		} else {
			var _200_v80 D2
			_ = _200_v80
			_200_v80 = Companion_D2_.Create_DC4_(_84_v4, _188_v73, _79_v0)
			var _201_v81 _dafny.Map
			_ = _201_v81
			_201_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_84_v4).Cardinality()), (_164_v57).F8())
			var _202_v82 *C1
			_ = _202_v82
			var _nw31 *C1 = New_C1_()
			_ = _nw31
			_nw31.Ctor__((Companion_D5_.Create_DC10_(_163_v56)).Dtor_cf22(), true, (_84_v4).Select((Companion_Default___.SafeIndex((_163_v56).Cardinality(), _dafny.IntOfUint32((_84_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), _89_v7, _89_v7)
			_202_v82 = _nw31
			var _203_v83 D4
			_ = _203_v83
			_203_v83 = Companion_D4_.Create_DC8_((_160_v53).Fm18((_164_v57).F8(), (_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int), (_160_v53).Fm18(_164_v57.F9(), _dafny.IntOfUint32((_188_v73).Cardinality()), (_160_v53).F14(), (_160_v53).F14(), _83_globalState), (_160_v53).F14(), _83_globalState), (_164_v57).F8(), (_201_v81).Cardinality(), _202_v82)
			var _204_v84 _dafny.Array
			_ = _204_v84
			var _nwElement0_1 _dafny.Sequence = Companion_Default___.Fm23(!(_164_v57.F9()), _164_v57.F9(), _dafny.IntOfUint32((_188_v73).Cardinality()), _83_globalState)
			_ = _nwElement0_1
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(27))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_1, 0)
			(_nw32).ArraySet1(_188_v73, 1)
			(_nw32).ArraySet1(_188_v73, 2)
			(_nw32).ArraySet1(_188_v73, 3)
			(_nw32).ArraySet1(_dafny.SeqOf((_82_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))).Int()).(bool)), 4)
			(_nw32).ArraySet1(_188_v73, 5)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v73, _188_v73), 6)
			(_nw32).ArraySet1(_188_v73, 7)
			(_nw32).ArraySet1(_188_v73, 8)
			(_nw32).ArraySet1(Companion_Default___.Fm23(!((_164_v57).F8()), !(_89_v7), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yowvpprj")).Cardinality()), _83_globalState), 9)
			(_nw32).ArraySet1((_200_v80).Dtor_cf11(), 10)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v73, _dafny.Companion_Sequence_.Update(_188_v73, (Companion_Default___.SafeIndex(_79_v0, _dafny.IntOfUint32((_188_v73).Cardinality()))).Uint32(), _89_v7)), 11)
			(_nw32).ArraySet1(_188_v73, 12)
			(_nw32).ArraySet1((func() _dafny.Sequence {
				if (_203_v83).Dtor_cf17() {
					return _188_v73
				}
				return _188_v73
			})(), 13)
			(_nw32).ArraySet1(_188_v73, 14)
			(_nw32).ArraySet1(_188_v73, 15)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_188_v73, _dafny.SeqOf(false, _89_v7)), (Companion_Default___.SafeIndex((_160_v53).F14(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_188_v73, _dafny.SeqOf(false, _89_v7))).Cardinality()))).Uint32(), _89_v7), 16)
			(_nw32).ArraySet1(_188_v73, 17)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v73, _188_v73), 18)
			(_nw32).ArraySet1(_188_v73, 19)
			(_nw32).ArraySet1(_dafny.SeqOf((_164_v57).F8(), true), 20)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v73, _188_v73), 21)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v73, _dafny.Companion_Sequence_.Update(_188_v73, (Companion_Default___.SafeIndex((_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_188_v73).Cardinality()))).Uint32(), (_82_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))).Int()).(bool))), 22)
			(_nw32).ArraySet1(_188_v73, 23)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v73, _188_v73), 24)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_188_v73, _dafny.SeqOf((_202_v82).F12())), 25)
			(_nw32).ArraySet1(_dafny.SeqOf((_164_v57).F8(), (_202_v82).F12(), (_202_v82).F12(), !((_202_v82).F12()), _164_v57.F9()), 26)
			_204_v84 = _nw32
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_204_v84), 0))
			_ = _index37
			(_204_v84).ArraySet1(Companion_Default___.Fm23((_164_v57).F8(), (_164_v57).F8(), (_dafny.SetOf(_164_v57.F10(), _dafny.CodePoint('m'))).Cardinality(), _83_globalState), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_204_v84), 0))
			_ = _index38
			(_204_v84).ArraySet1(_188_v73, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))
			_ = _index39
			(_82_v3).ArraySet1((_79_v0).Cmp((_160_v53).F14()) >= 0, (_index39).Int())
			var _205_v85 _dafny.Map
			_ = _205_v85
			_205_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_89_v7), _dafny.SeqOf((_164_v57).F8()))
			(_83_globalState).F5 = ((_205_v85).Merge(_205_v85)).Cardinality()
			var _206_v86 _dafny.Map
			_ = _206_v86
			_206_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_160_v53).F14(), _94_v11)
			var _207_v87 _dafny.Sequence
			_ = _207_v87
			_207_v87 = _dafny.SeqOf(((func() _dafny.Map {
				if (_206_v86).Contains(_dafny.IntOfInt64(566)) {
					return (_206_v86).Get(_dafny.IntOfInt64(566)).(_dafny.Map)
				}
				return _94_v11
			})()).Update((_164_v57).F8(), (_86_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_86_v6), 0))).Int()).(_dafny.Int)), _90_v8, _94_v11)
			var _208_v88 _dafny.Set
			_ = _208_v88
			_208_v88 = _dafny.SetOf(true, (_164_v57).F8())
			var _209_v89 _dafny.Map
			_ = _209_v89
			_209_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v7, _90_v8)
			var _210_v90 _dafny.Set
			_ = _210_v90
			_210_v90 = _dafny.SetOf((_207_v87).Select((Companion_Default___.SafeIndex((_208_v88).Cardinality(), _dafny.IntOfUint32((_207_v87).Cardinality()))).Uint32()).(_dafny.Map), _94_v11, (func() _dafny.Map {
				if (_209_v89).Contains(_164_v57.F9()) {
					return (_209_v89).Get(_164_v57.F9()).(_dafny.Map)
				}
				return _94_v11
			})())
			var _211_v91 _dafny.Sequence
			_ = _211_v91
			_211_v91 = _dafny.SeqOf(_80_v1, _80_v1)
			var _212_v92 _dafny.Set
			_ = _212_v92
			_212_v92 = _dafny.SetOf(_160_v53, _160_v53)
			var _213_v93 _dafny.Map
			_ = _213_v93
			_213_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_80_v1).Select((Companion_Default___.SafeIndex((_160_v53).Fm17(_164_v57.F9(), false, _211_v91, _dafny.IntOfUint32((_84_v4).Cardinality()), _83_globalState), _dafny.IntOfUint32((_80_v1).Cardinality()))).Uint32()).(_dafny.Int), _212_v92)
			var _214_v94 _dafny.Int
			_ = _214_v94
			var _out6 _dafny.Int
			_ = _out6
			_out6 = Companion_Default___.M0((_164_v57).F8(), _210_v90, ((func() _dafny.Set {
				if (_213_v93).Contains((_160_v53).F14()) {
					return (_213_v93).Get((_160_v53).F14()).(_dafny.Set)
				}
				return _212_v92
			})()).IsProperSubsetOf(_dafny.SetOf(_160_v53)), _83_globalState)
			_214_v94 = _out6
			var _215_v95 *C0
			_ = _215_v95
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__(_dafny.IntOfInt64(285), _dafny.CodePoint('y'), (_82_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))).Int()).(bool), (_164_v57).F8())
			_215_v95 = _nw33
			var _216_v96 _dafny.Sequence
			_ = _216_v96
			_216_v96 = _dafny.SeqOf(_215_v95)
			var _217_v97 _dafny.Set
			_ = _217_v97
			_217_v97 = _dafny.SetOf(_216_v96)
			_217_v97 = _dafny.SetOf(_216_v96, _dafny.Companion_Sequence_.Concatenate(_216_v96, _216_v96), _216_v96)
		}
		var _218_v98 _dafny.Set
		_ = _218_v98
		_218_v98 = _dafny.SetOf(_89_v7, _164_v57.F9())
		var _219_v99 *C0
		_ = _219_v99
		var _nw34 *C0 = New_C0_()
		_ = _nw34
		_nw34.Ctor__(((_218_v98).Cardinality()).Times((_164_v57).Fm6((_82_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))).Int()).(bool), _83_globalState)), _159_v52, (func() bool {
			if !(_164_v57.F9()) {
				return (_82_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_82_v3), 0))).Int()).(bool)
			}
			return _89_v7
		})(), _89_v7)
		_219_v99 = _nw34
		_79_v0 = _219_v99.F13
	}
	_89_v7 = !(_89_v7)
	var _220_v100 _dafny.Map
	_ = _220_v100
	_220_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(956))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}((func(_221_v57 T1) func(_dafny.Int) _dafny.CodePoint {
		return func(_222_i10 _dafny.Int) _dafny.CodePoint {
			return _221_v57.F10()
		}
	})(_164_v57))), _164_v57.F9())
	var _223_v101 *C0
	_ = _223_v101
	var _nw35 *C0 = New_C0_()
	_ = _nw35
	_nw35.Ctor__((_160_v53).F14(), _159_v52, Companion_Default___.Fm22((_160_v53).F14(), _164_v57.F9(), (_164_v57).F8(), _164_v57.F9(), _83_globalState), false)
	_223_v101 = _nw35
	var _224_v102 _dafny.Map
	_ = _224_v102
	_224_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v0, _223_v101)
	_220_v100 = (_220_v100).Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("v"), (Companion_Default___.SafeIndex((_dafny.MultiSetOf(_223_v101, _223_v101, (func() *C0 {
		if (_224_v102).Contains(_223_v101.F13) {
			return (_224_v102).Get(_223_v101.F13).(*C0)
		}
		return _223_v101
	})(), _223_v101)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality()))).Uint32(), _159_v52), _164_v57.F9())
	var _225_v103 _dafny.Sequence
	_ = _225_v103
	_225_v103 = _dafny.SeqOf((_164_v57.F9()) == (_164_v57.F9()), (_164_v57).F8())
	_89_v7 = (_225_v103).Select((Companion_Default___.SafeIndex((_223_v101.F13).Times(_223_v101.F13), _dafny.IntOfUint32((_225_v103).Cardinality()))).Uint32()).(bool)
	for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_82_v3), 0))); ; {
		_guard_loop_0, _ok11 := _iter11()
		if !_ok11 {
			break
		}
		var _226_i11 _dafny.Int
		_226_i11 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_226_i11).Sign() != -1) && ((_226_i11).Cmp(_dafny.ArrayLen((_82_v3), 0)) < 0)) {
			(_82_v3).ArraySet1((_168_v61).Equals((_168_v61).Union(_168_v61)), (_226_i11).Int())
		}
	}
	_dafny.Print(_79_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_80_v1, _dafny.SeqOf(_dafny.IntOfInt64(717))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(717), _dafny.SeqOf(_dafny.IntOfInt64(717)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_globalState.F0).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(717), _dafny.SeqOf(_dafny.IntOfInt64(717)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_83_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_84_v4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("qhbqusrq"), _dafny.UnicodeSeqOfUtf8Bytes("qhbqusrq"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_89_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(717))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("qhbqusrq"), _dafny.UnicodeSeqOfUtf8Bytes("qhbqusrq"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf2()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(717))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_v10).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_94_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("qhbqusrq"), _dafny.SeqOf(_dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf2()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(717))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf3()).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_95_v12).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_v52)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v53).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_161_v54, _dafny.SeqOf(Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lwb")), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lwb")), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lwb")))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_162_v55, _dafny.SeqOf(_dafny.SeqOf(Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lwb"))), _dafny.SeqOf(Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lwb"))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v56).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(717)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_164_v57.F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v57).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_164_v57.F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_165_v58).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.Zero).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.Zero).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.Zero).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.One).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.One).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.One).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(2)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(2)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(3)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(3)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(4)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(4)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(5)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(5)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(6)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(6)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(7)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(7)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(8)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(8)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(9)).Int())).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(Companion_T1_.CastTo_((_166_v59).ArrayGet1((_dafny.IntOfInt64(9)).Int())).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v60).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(16), _dafny.IntOfInt64(16))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v61).Equals(_dafny.SetOf(_dafny.IntOfInt64(16))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v100).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v')), true).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("v"), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_223_v101.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_225_v103, _dafny.SeqOf(true, true)))
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

type D0_DC0 struct {
	Cf0 _dafny.Map
	Cf1 _dafny.Array
	Cf2 _dafny.Map
	Cf3 _dafny.Array
	Cf4 _dafny.Array
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Map, Cf1 _dafny.Array, Cf2 _dafny.Map, Cf3 _dafny.Array, Cf4 _dafny.Array) D0 {
	return D0{D0_DC0{Cf0, Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.EmptyMap, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMap, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf0() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf4
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D0) Equals(other D0) bool {
	switch data1 := _this.Get_().(type) {
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0) && data1.Cf1 == data2.Cf1 && data1.Cf2.Equals(data2.Cf2) && data1.Cf3 == data2.Cf3 && data1.Cf4 == data2.Cf4
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
	Cf6 _dafny.Int
	Cf7 _dafny.Int
	Cf8 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf6 _dafny.Int, Cf7 _dafny.Int, Cf8 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf5 _dafny.Sequence
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf5 _dafny.Sequence) D1 {
	return D1{D1_DC1{Cf5}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, _dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf7
}

func (_this D1) Dtor_cf8() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf8
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC1).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + data.Cf5.VerbatimString(true) + ")"
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
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf5.Equals(data2.Cf5)
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
	Cf10 _dafny.Sequence
	Cf11 _dafny.Sequence
	Cf12 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf10 _dafny.Sequence, Cf11 _dafny.Sequence, Cf12 _dafny.Int) D2 {
	return D2{D2_DC4{Cf10, Cf11, Cf12}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC3 struct {
	Cf9 bool
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf9 bool) D2 {
	return D2{D2_DC3{Cf9}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC5 struct {
	Cf13 D2
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf13 D2) D2 {
	return D2{D2_DC5{Cf13}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.EmptySeq, _dafny.EmptySeq, _dafny.Zero)
}

func (_this D2) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D2_DC4).Cf10
}

func (_this D2) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D2_DC4).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf12
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC3).Cf9
}

func (_this D2) Dtor_cf13() D2 {
	return _this.Get_().(D2_DC5).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + data.Cf10.VerbatimString(true) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf10.Equals(data2.Cf10) && data1.Cf11.Equals(data2.Cf11) && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf9 == data2.Cf9
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf13.Equals(data2.Cf13)
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

type D3_DC6 struct {
	Cf14 _dafny.MultiSet
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf14 _dafny.MultiSet) D3 {
	return D3{D3_DC6{Cf14}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D3) Dtor_cf14() _dafny.MultiSet {
	return _this.Get_().(D3_DC6).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf14) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D4_DC8 struct {
	Cf16 _dafny.Int
	Cf17 bool
	Cf18 _dafny.Int
	Cf19 *C1
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf16 _dafny.Int, Cf17 bool, Cf18 _dafny.Int, Cf19 *C1) D4 {
	return D4{D4_DC8{Cf16, Cf17, Cf18, Cf19}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC9 struct {
	Cf20 bool
	Cf21 bool
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf20 bool, Cf21 bool) D4 {
	return D4{D4_DC9{Cf20, Cf21}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC7 struct {
	Cf15 _dafny.Sequence
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf15 _dafny.Sequence) D4 {
	return D4{D4_DC7{Cf15}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC8_(_dafny.Zero, false, _dafny.Zero, (*C1)(nil))
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC8).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC8).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC8).Cf18
}

func (_this D4) Dtor_cf19() *C1 {
	return _this.Get_().(D4_DC8).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC9).Cf20
}

func (_this D4) Dtor_cf21() bool {
	return _this.Get_().(D4_DC9).Cf21
}

func (_this D4) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D4_DC7).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21
		}
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_() D5 {
	return D5{D5_DC11{}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC12 struct {
	Cf23 bool
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf23 bool) D5 {
	return D5{D5_DC12{Cf23}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC10 struct {
	Cf22 _dafny.MultiSet
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf22 _dafny.MultiSet) D5 {
	return D5{D5_DC10{Cf22}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC11_()
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC12).Cf23
}

func (_this D5) Dtor_cf22() _dafny.MultiSet {
	return _this.Get_().(D5_DC10).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf22) + ")"
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
			_, ok := other.Get_().(D5_DC11)
			return ok
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf22.Equals(data2.Cf22)
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

type D6_DC14 struct {
	Cf25 _dafny.Int
	Cf26 _dafny.Int
	Cf27 _dafny.Int
	Cf28 _dafny.Int
	Cf29 bool
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf25 _dafny.Int, Cf26 _dafny.Int, Cf27 _dafny.Int, Cf28 _dafny.Int, Cf29 bool) D6 {
	return D6{D6_DC14{Cf25, Cf26, Cf27, Cf28, Cf29}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC15 struct {
	Cf30 _dafny.Int
	Cf31 _dafny.Int
	Cf32 _dafny.Set
	Cf33 _dafny.Set
	Cf34 _dafny.Int
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf30 _dafny.Int, Cf31 _dafny.Int, Cf32 _dafny.Set, Cf33 _dafny.Set, Cf34 _dafny.Int) D6 {
	return D6{D6_DC15{Cf30, Cf31, Cf32, Cf33, Cf34}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC13 struct {
	Cf24 _dafny.MultiSet
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf24 _dafny.MultiSet) D6 {
	return D6{D6_DC13{Cf24}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero, false)
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf25
}

func (_this D6) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf27
}

func (_this D6) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf28
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC14).Cf29
}

func (_this D6) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf30
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Set {
	return _this.Get_().(D6_DC15).Cf32
}

func (_this D6) Dtor_cf33() _dafny.Set {
	return _this.Get_().(D6_DC15).Cf33
}

func (_this D6) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf34
}

func (_this D6) Dtor_cf24() _dafny.MultiSet {
	return _this.Get_().(D6_DC13).Cf24
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29 == data2.Cf29
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32.Equals(data2.Cf32) && data1.Cf33.Equals(data2.Cf33) && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

// Definition of trait T0
type T0 interface {
	String() string
	F9() bool
	F9_set_(value bool)
	Fm5(p0 bool, globalState *GlobalState) bool
	F8() bool
}
type CompanionStruct_T0_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T0_ = CompanionStruct_T0_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T0_) CastTo_(x interface{}) T0 {
	var t T0
	t, _ = x.(T0)
	return t
}

// End of trait T0

// Definition of trait T1
type T1 interface {
	String() string
	F9() bool
	F9_set_(value bool)
	Fm5(p0 bool, globalState *GlobalState) bool
	F8() bool
	F10() _dafny.CodePoint
	F10_set_(value _dafny.CodePoint)
	Fm6(p0 bool, globalState *GlobalState) _dafny.Int
	Fm7(p0 bool, p1 _dafny.Set, p2 _dafny.Sequence, globalState *GlobalState) bool
}
type CompanionStruct_T1_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T1_ = CompanionStruct_T1_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T1_) CastTo_(x interface{}) T1 {
	var t T1
	t, _ = x.(T1)
	return t
}

// End of trait T1

// Definition of class GlobalState
type GlobalState struct {
	F0  _dafny.Map
	F5  _dafny.Int
	_f1 bool
	_f2 bool
	_f3 _dafny.Array
	_f4 bool
	_f6 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptyMap
	_this.F5 = _dafny.Zero
	_this._f1 = false
	_this._f2 = false
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = false
	_this._f6 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 bool, f2 bool, f3 _dafny.Array, f4 bool, f5 _dafny.Int, f6 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Array {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f10 _dafny.CodePoint
	_f9  bool
	_f8  bool
	F13  _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f10 = _dafny.CodePoint('D')
	_this._f9 = false
	_this._f8 = false
	_this.F13 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C0{}
var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F10() _dafny.CodePoint {
	return _this._f10
}
func (_this *C0) F10_set_(value _dafny.CodePoint) {
	_this._f10 = value
}
func (_this *C0) F9() bool {
	return _this._f9
}
func (_this *C0) F9_set_(value bool) {
	_this._f9 = value
}
func (_this *C0) F8() bool {
	return _this._f8
}
func (_this *C0) Ctor__(f13 _dafny.Int, f10 _dafny.CodePoint, f8 bool, f9 bool) {
	{
		(_this).F13 = f13
		(_this)._f10 = f10
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C0) Fm6(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.SetOf((_this).F8(), (_this).F8())).Cardinality()
	}
}
func (_this *C0) Fm7(p0 bool, p1 _dafny.Set, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F8()
	}
}
func (_this *C0) Fm5(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F8()
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f10 _dafny.CodePoint
	_f9  bool
	_f8  bool
	_f11 _dafny.MultiSet
	_f12 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f10 = _dafny.CodePoint('D')
	_this._f9 = false
	_this._f8 = false
	_this._f11 = _dafny.EmptyMultiSet
	_this._f12 = false
	return &_this
}

type CompanionStruct_C1_ struct {
}

var Companion_C1_ = CompanionStruct_C1_{}

func (_this *C1) Equals(other *C1) bool {
	return _this == other
}

func (_this *C1) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C1)
	return ok && _this.Equals(other)
}

func (*C1) String() string {
	return "_module.C1"
}

func Type_C1_() _dafny.TypeDescriptor {
	return type_C1_{}
}

type type_C1_ struct {
}

func (_this type_C1_) Default() interface{} {
	return (*C1)(nil)
}

func (_this type_C1_) String() string {
	return "main.C1"
}
func (_this *C1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F10() _dafny.CodePoint {
	return _this._f10
}
func (_this *C1) F10_set_(value _dafny.CodePoint) {
	_this._f10 = value
}
func (_this *C1) F9() bool {
	return _this._f9
}
func (_this *C1) F9_set_(value bool) {
	_this._f9 = value
}
func (_this *C1) F8() bool {
	return _this._f8
}
func (_this *C1) Ctor__(f11 _dafny.MultiSet, f12 bool, f10 _dafny.CodePoint, f8 bool, f9 bool) {
	{
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f10 = f10
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C1) Fm6(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_this).F12())).Cardinality(), (_this).F8())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(18), (_this).F12()))).Cardinality()
	}
}
func (_this *C1) Fm7(p0 bool, p1 _dafny.Set, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F8()
	}
}
func (_this *C1) Fm5(p0 bool, globalState *GlobalState) bool {
	{
		return !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((_dafny.MultiSetOf((_this).F12())).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(495))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}(func(_227_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(42)
		}))))
	}
}
func (_this *C1) Fm8(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(_dafny.IntOfInt64(-303))
	}
}
func (_this *C1) M4(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) (_dafny.Int, bool, bool, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		r1 = (p0).Cmp(_dafny.IntOfInt64(346)) == 0
		var _228_v0 T1
		_ = _228_v0
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__(_dafny.IntOfInt64(303), _this.F10(), p3, (_this).Fm7(false, _dafny.SetOf(true), p1, globalState))
		_228_v0 = _nw36
		var _229_v1 _dafny.Map
		_ = _229_v1
		_229_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _228_v0)
		_229_v1 = _229_v1
		var _230_v2 _dafny.Sequence
		_ = _230_v2
		_230_v2 = _dafny.SeqOf((_this).F8(), (_this).F12())
		var _231_v3 D2
		_ = _231_v3
		_231_v3 = Companion_D2_.Create_DC3_((_230_v2).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_230_v2).Cardinality()))).Uint32()).(bool))
		var _source3 D2 = _231_v3
		_ = _source3
		if _source3.Is_DC4() {
			var _232___mcc_h0 _dafny.Sequence = _source3.Get_().(D2_DC4).Cf10
			_ = _232___mcc_h0
			var _233___mcc_h1 _dafny.Sequence = _source3.Get_().(D2_DC4).Cf11
			_ = _233___mcc_h1
			var _234___mcc_h2 _dafny.Int = _source3.Get_().(D2_DC4).Cf12
			_ = _234___mcc_h2
			var _235_cf12 _dafny.Int = _234___mcc_h2
			_ = _235_cf12
			var _236_cf11 _dafny.Sequence = _233___mcc_h1
			_ = _236_cf11
			var _237_cf10 _dafny.Sequence = _232___mcc_h0
			_ = _237_cf10
			var _238_v4 _dafny.Map
			_ = _238_v4
			_238_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _235_cf12)
			r0 = (func() _dafny.Int {
				if (_238_v4).Contains(p1) {
					return (_238_v4).Get(p1).(_dafny.Int)
				}
				return Companion_Default___.Fm0(p3, _237_cf10, globalState)
			})()
			var _239_v5 _dafny.MultiSet
			_ = _239_v5
			_239_v5 = _dafny.MultiSetOf(true, !((_231_v3).Dtor_cf9()), false, true, (_228_v0).F8())
			var _240_v6 _dafny.Map
			_ = _240_v6
			_240_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm9(_235_cf12, (_228_v0).F8(), globalState), _239_v5)
			var _241_v7 _dafny.Set
			_ = _241_v7
			_241_v7 = _dafny.SetOf(_228_v0.F9())
			var _rhs22 bool = true
			_ = _rhs22
			var _rhs23 D2 = Companion_D2_.Create_DC3_((_241_v7).IsSubsetOf(_241_v7))
			_ = _rhs23
			var _rhs24 bool = (p3) == ((_241_v7).IsSubsetOf(_241_v7))
			_ = _rhs24
			var _rhs25 _dafny.Map = _240_v6
			_ = _rhs25
			var _rhs26 bool = (_239_v5).IsSubsetOf(_239_v5)
			_ = _rhs26
			var _lhs16 *C1 = _this
			_ = _lhs16
			_lhs16.F9_set_(_rhs22)
			_231_v3 = _rhs23
			r3 = _rhs24
			_240_v6 = _rhs25
			r3 = _rhs26
			var _242_v8 _dafny.Map
			_ = _242_v8
			_242_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _235_cf12)
			var _243_v9 _dafny.Map
			_ = _243_v9
			_243_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_231_v3).Dtor_cf9(), _237_cf10)
			_242_v8 = Companion_Default___.Fm10((_235_cf12).Plus(p0), _243_v9, globalState)
			var _244_v10 _dafny.Map
			_ = _244_v10
			_244_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_242_v8).Contains(_235_cf12) {
					return (_242_v8).Get(_235_cf12).(_dafny.Int)
				}
				return p0
			})(), _dafny.CodePoint('j'))
			var _245_v11 D1
			_ = _245_v11
			_245_v11 = Companion_D1_.Create_DC2_((_244_v10).Cardinality(), _dafny.IntOfUint32((_236_cf11).Cardinality()), _228_v0.F10())
			(globalState).F5 = (_245_v11).Dtor_cf7()
		} else if _source3.Is_DC3() {
			var _246___mcc_h3 bool = _source3.Get_().(D2_DC3).Cf9
			_ = _246___mcc_h3
			var _247_cf9 bool = _246___mcc_h3
			_ = _247_cf9
			var _248_v12 _dafny.Array
			_ = _248_v12
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_8
			var _nw37 _dafny.Array
			_ = _nw37
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw37 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = (func(_249_p0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_250_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(56))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg24 _dafny.Int) interface{} {
								return coer24(arg24)
							}
						}((func(_251_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_252_i1 _dafny.Int) _dafny.Int {
								return _251_p0
							}
						})(_249_p0)))
					}
				})(p0)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw37 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw37).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw37).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_248_v12 = _nw37
			var _253_v13 _dafny.Sequence
			_ = _253_v13
			_253_v13 = _dafny.SeqOf(_dafny.IntOfInt64(392))
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_248_v12), 0))
			_ = _index40
			(_248_v12).ArraySet1(_253_v13, (_index40).Int())
			var _254_v14 _dafny.Sequence
			_ = _254_v14
			_254_v14 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_248_v12), 0))
			_ = _index41
			var _rhs27 _dafny.Sequence = _253_v13
			_ = _rhs27
			var _rhs28 bool = _247_cf9
			_ = _rhs28
			var _rhs29 _dafny.Sequence = _254_v14
			_ = _rhs29
			var _rhs30 _dafny.Sequence = _254_v14
			_ = _rhs30
			var _rhs31 bool = (_this).F8()
			_ = _rhs31
			var _lhs17 _dafny.Array = _248_v12
			_ = _lhs17
			var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_248_v12), 0))
			_ = _lhs18
			var _lhs19 T1 = _228_v0
			_ = _lhs19
			(_lhs17).ArraySet1(_rhs27, (_lhs18).Int())
			_lhs19.F9_set_(_rhs28)
			_254_v14 = _rhs29
			_254_v14 = _rhs30
			_247_cf9 = _rhs31
			var _255_v15 _dafny.Set
			_ = _255_v15
			_255_v15 = _dafny.SetOf((_dafny.Zero).Minus((p0).Minus(p0)), p0, (_dafny.Zero).Minus(p0), p0)
			_255_v15 = _255_v15
			r2 = !_dafny.Companion_Sequence_.Contains(p1, (func() _dafny.CodePoint {
				if _228_v0.F9() {
					return _this.F10()
				}
				return p2
			})())
			var _256_v16 _dafny.Map
			_ = _256_v16
			_256_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), p0)
			var _257_v17 _dafny.Int
			_ = _257_v17
			var _out7 _dafny.Int
			_ = _out7
			_out7 = Companion_Default___.M0(_228_v0.F9(), _dafny.SetOf(_256_v16), (_231_v3).Dtor_cf9(), globalState)
			_257_v17 = _out7
		} else {
			var _258___mcc_h4 D2 = _source3.Get_().(D2_DC5).Cf13
			_ = _258___mcc_h4
			var _259_cf13 D2 = _258___mcc_h4
			_ = _259_cf13
			var _260_v18 _dafny.Sequence
			_ = _260_v18
			_260_v18 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("eahskv"))
			var _261_v19 _dafny.Map
			_ = _261_v19
			_261_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(740))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_262_v0 T1) func(_dafny.Int) _dafny.CodePoint {
				return func(_263_i2 _dafny.Int) _dafny.CodePoint {
					return _262_v0.F10()
				}
			})(_228_v0))), (_260_v18).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_260_v18).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _264_v20 _dafny.Array
			_ = _264_v20
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_9
			var _nw38 _dafny.Array
			_ = _nw38
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw38 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = func(_265_i3 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_265_i3, (_dafny.SetOf(!(true), false)).Cardinality())
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw38 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw38).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw38).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_264_v20 = _nw38
			var _266_v21 _dafny.Map
			_ = _266_v21
			_266_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_228_v0.F9(), p0)
			var _267_v22 _dafny.Array
			_ = _267_v22
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
			_ = _nw39
			_267_v22 = _nw39
			var _268_v23 D0
			_ = _268_v23
			_268_v23 = Companion_D0_.Create_DC0_(_261_v19, _264_v20, (_266_v21).Update((_228_v0).Fm5((_228_v0).F8(), globalState), p0), _267_v22, _264_v20)
			_268_v23 = _268_v23
			(globalState).F5 = p0
			var _269_v24 _dafny.Set
			_ = _269_v24
			_269_v24 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(p1, p1), p1)
			_269_v24 = (_269_v24).Difference(_dafny.SetOf(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _dafny.CodePoint('n'))))
			var _270_v25 D2
			_ = _270_v25
			_270_v25 = Companion_D2_.Create_DC4_(p1, _230_v2, p0)
			var _271_v26 _dafny.Sequence
			_ = _271_v26
			_271_v26 = _dafny.SeqOf(_270_v25)
			var _272_v27 D2
			_ = _272_v27
			_272_v27 = Companion_D2_.Create_DC5_((_271_v26).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_271_v26).Cardinality()))).Uint32()).(D2))
			var _273_v28 D2
			_ = _273_v28
			_273_v28 = Companion_D2_.Create_DC5_(_272_v27)
			var _274_v29 D2
			_ = _274_v29
			_274_v29 = Companion_D2_.Create_DC5_(_270_v25)
			var _275_v30 D2
			_ = _275_v30
			_275_v30 = Companion_D2_.Create_DC5_(_272_v27)
			var _276_v31 _dafny.Map
			_ = _276_v31
			_276_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf(p0, p0))
			var _277_v32 _dafny.Set
			_ = _277_v32
			_277_v32 = _dafny.SetOf(p0)
			var _rhs32 _dafny.Int = (_dafny.IntOfInt64(939)).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-269), _dafny.IntOfUint32((_230_v2).Cardinality())))
			_ = _rhs32
			var _rhs33 bool = _this.F9()
			_ = _rhs33
			var _rhs34 bool = !(((func() _dafny.Set {
				if (_276_v31).Contains(p0) {
					return (_276_v31).Get(p0).(_dafny.Set)
				}
				return _277_v32
			})()).Equals(_277_v32))
			_ = _rhs34
			var _rhs35 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_266_v21, _266_v21, _266_v21)).Cardinality())
			_ = _rhs35
			var _rhs36 D2 = _275_v30
			_ = _rhs36
			var _lhs20 *GlobalState = globalState
			_ = _lhs20
			var _lhs21 T1 = _228_v0
			_ = _lhs21
			var _lhs22 *GlobalState = globalState
			_ = _lhs22
			_lhs20.F5 = _rhs32
			_lhs21.F9_set_(_rhs33)
			r2 = _rhs34
			_lhs22.F5 = _rhs35
			_275_v30 = _rhs36
		}
		var _278_v33 _dafny.Map
		_ = _278_v33
		_278_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_228_v0).Fm5((_this).F12(), globalState), p0)
		var _279_v34 _dafny.Set
		_ = _279_v34
		_279_v34 = _dafny.SetOf(_278_v33, _278_v33)
		var _280_v35 _dafny.Int
		_ = _280_v35
		var _out8 _dafny.Int
		_ = _out8
		_out8 = Companion_Default___.M0((func() bool {
			if (_this).F8() {
				return true
			}
			return true
		})(), _279_v34, !((p0).Cmp(p0) >= 0), globalState)
		_280_v35 = _out8
		var _281_v36 _dafny.Set
		_ = _281_v36
		_281_v36 = _dafny.SetOf(_this.F10(), p2)
		var _282_v37 *C0
		_ = _282_v37
		var _nw40 *C0 = New_C0_()
		_ = _nw40
		_nw40.Ctor__(p0, _228_v0.F10(), !(!((_this).F8())) || (_this.F9()), !(_281_v36).Contains(_228_v0.F10()))
		_282_v37 = _nw40
		_282_v37 = _282_v37
		(_this).F10_set_((p1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_280_v35), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint))
		r0 = p0
		var _283_v38 _dafny.Set
		_ = _283_v38
		_283_v38 = _dafny.SetOf((_228_v0).F8())
		r1 = (_228_v0).Fm7((_this).F8(), _283_v38, p1, globalState)
		r2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(150))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_284_p1 _dafny.Sequence, _285_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
			return func(_286_i4 _dafny.Int) _dafny.CodePoint {
				return (_284_p1).Select((Companion_Default___.SafeIndex(_285_p0, _dafny.IntOfUint32((_284_p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(p1, p0))), (Companion_D1_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(197))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_287_v0 T1) func(_dafny.Int) _dafny.CodePoint {
			return func(_288_i5 _dafny.Int) _dafny.CodePoint {
				return _287_v0.F10()
			}
		})(_228_v0))))).Dtor_cf5())
		r3 = _this.F9()
		return r0, r1, r2, r3
	}
}
func (_this *C1) F11() _dafny.MultiSet {
	{
		return _this._f11
	}
}
func (_this *C1) F12() bool {
	{
		return _this._f12
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f9  bool
	_f8  bool
	_f14 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f9 = false
	_this._f8 = false
	_this._f14 = _dafny.Zero
	return &_this
}

type CompanionStruct_C2_ struct {
}

var Companion_C2_ = CompanionStruct_C2_{}

func (_this *C2) Equals(other *C2) bool {
	return _this == other
}

func (_this *C2) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C2)
	return ok && _this.Equals(other)
}

func (*C2) String() string {
	return "_module.C2"
}

func Type_C2_() _dafny.TypeDescriptor {
	return type_C2_{}
}

type type_C2_ struct {
}

func (_this type_C2_) Default() interface{} {
	return (*C2)(nil)
}

func (_this type_C2_) String() string {
	return "main.C2"
}
func (_this *C2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F9() bool {
	return _this._f9
}
func (_this *C2) F9_set_(value bool) {
	_this._f9 = value
}
func (_this *C2) F8() bool {
	return _this._f8
}
func (_this *C2) Ctor__(f14 _dafny.Int, f8 bool, f9 bool) {
	{
		(_this)._f14 = f14
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C2) Fm5(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F8()
	}
}
func (_this *C2) Fm17(p0 bool, p1 bool, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var _source4 D2 = Companion_D2_.Create_DC3_((_this).F8())
		_ = _source4
		if _source4.Is_DC4() {
			var _289___mcc_h0 _dafny.Sequence = _source4.Get_().(D2_DC4).Cf10
			_ = _289___mcc_h0
			var _290___mcc_h1 _dafny.Sequence = _source4.Get_().(D2_DC4).Cf11
			_ = _290___mcc_h1
			var _291___mcc_h2 _dafny.Int = _source4.Get_().(D2_DC4).Cf12
			_ = _291___mcc_h2
			var _292_cf12 _dafny.Int = _291___mcc_h2
			_ = _292_cf12
			var _293_cf11 _dafny.Sequence = _290___mcc_h1
			_ = _293_cf11
			var _294_cf10 _dafny.Sequence = _289___mcc_h0
			_ = _294_cf10
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_294_cf10, _294_cf10)).Cardinality())
		} else if _source4.Is_DC3() {
			var _295___mcc_h3 bool = _source4.Get_().(D2_DC3).Cf9
			_ = _295___mcc_h3
			var _296_cf9 bool = _295___mcc_h3
			_ = _296_cf9
			return (_this).F14()
		} else {
			var _297___mcc_h4 D2 = _source4.Get_().(D2_DC5).Cf13
			_ = _297___mcc_h4
			var _298_cf13 D2 = _297___mcc_h4
			_ = _298_cf13
			return ((_this).F14()).Times((_dafny.Zero).Minus((_this).F14()))
		}
	}
}
func (_this *C2) Fm18(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((_this).F14()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F14())).Cardinality()))
	}
}
func (_this *C2) F14() _dafny.Int {
	{
		return _this._f14
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	dummy byte
}

func New_C3_() *C3 {
	_this := C3{}

	return &_this
}

type CompanionStruct_C3_ struct {
}

var Companion_C3_ = CompanionStruct_C3_{}

func (_this *C3) Equals(other *C3) bool {
	return _this == other
}

func (_this *C3) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C3)
	return ok && _this.Equals(other)
}

func (*C3) String() string {
	return "_module.C3"
}

func Type_C3_() _dafny.TypeDescriptor {
	return type_C3_{}
}

type type_C3_ struct {
}

func (_this type_C3_) Default() interface{} {
	return (*C3)(nil)
}

func (_this type_C3_) String() string {
	return "main.C3"
}
func (_this *C3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__() {
	{
	}
}
func (_this *C3) Fm3(globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kwqewtt")).Cardinality()), _dafny.IntOfInt64(318))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-220))).IsDisjointFrom((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(128)))))
	}
}
func (_this *C3) Fm4(globalState *GlobalState) D2 {
	{
		return Companion_D2_.Create_DC4_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mmapr"), _dafny.UnicodeSeqOfUtf8Bytes("he")), _dafny.SeqOf(false, false), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-872)), _dafny.IntOfInt64(738)))
	}
}
func (_this *C3) M2(p0 bool, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) (_dafny.Int, _dafny.CodePoint, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var r2 bool = false
		_ = r2
		var _299_v0 _dafny.CodePoint
		_ = _299_v0
		_299_v0 = _dafny.CodePoint('t')
		var _300_v1 *C0
		_ = _300_v1
		var _nw41 *C0 = New_C0_()
		_ = _nw41
		_nw41.Ctor__(p2, _299_v0, p0, !(p0))
		_300_v1 = _nw41
		var _301_v2 _dafny.Array
		_ = _301_v2
		var _nw42 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw42
		_301_v2 = _nw42
		var _302_v3 _dafny.Sequence
		_ = _302_v3
		_302_v3 = _dafny.SeqOf(_301_v2)
		var _303_v4 _dafny.Map
		_ = _303_v4
		_303_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _301_v2)
		var _304_v5 _dafny.Array
		_ = _304_v5
		var _nwElement0_2 _dafny.Array = (_302_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_302_v3).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _nwElement0_2
		var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(8))
		_ = _nw43
		(_nw43).ArraySet1(_nwElement0_2, 0)
		(_nw43).ArraySet1(_301_v2, 1)
		(_nw43).ArraySet1(_301_v2, 2)
		(_nw43).ArraySet1(_301_v2, 3)
		(_nw43).ArraySet1(_301_v2, 4)
		(_nw43).ArraySet1(_301_v2, 5)
		(_nw43).ArraySet1(_301_v2, 6)
		(_nw43).ArraySet1((func() _dafny.Array {
			if (_303_v4).Contains(p0) {
				return (_303_v4).Get(p0).(_dafny.Array)
			}
			return _301_v2
		})(), 7)
		_304_v5 = _nw43
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_304_v5), 0))
		_ = _index42
		(_304_v5).ArraySet1(_301_v2, (_index42).Int())
		var _305_v6 _dafny.Array
		_ = _305_v6
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_10
		var _nw44 _dafny.Array
		_ = _nw44
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw44 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Int = func(_306_i0 _dafny.Int) _dafny.Int {
				return (_306_i0).Times(_dafny.IntOfInt64(968))
			}
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw44 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw44).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw44).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_305_v6 = _nw44
		var _307_v7 _dafny.Map
		_ = _307_v7
		_307_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v6, !(!(p0)))
		var _308_v8 _dafny.Set
		_ = _308_v8
		_308_v8 = _dafny.SetOf(_299_v0)
		var _309_v9 _dafny.Map
		_ = _309_v9
		_309_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-717), !(p0))
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_304_v5), 0))
		_ = _index43
		var _nwElement0_3 bool = (func() bool {
			if p0 {
				return p0
			}
			return p0
		})()
		_ = _nwElement0_3
		var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(7))
		_ = _nw45
		(_nw45).ArraySet1(_nwElement0_3, 0)
		(_nw45).ArraySet1(p0, 1)
		(_nw45).ArraySet1(p0, 2)
		(_nw45).ArraySet1(!(_307_v7).Contains(_305_v6), 3)
		(_nw45).ArraySet1((_dafny.SetOf(_299_v0)).IsSubsetOf(_308_v8), 4)
		(_nw45).ArraySet1(((Companion_D2_.Create_DC3_(!(p0))).Dtor_cf9()) == (p0), 5)
		(_nw45).ArraySet1((_309_v9).Equals(_309_v9), 6)
		(_304_v5).ArraySet1(_nw45, (_index43).Int())
		var _hi0 _dafny.Int = _300_v1.F13
		_ = _hi0
		for _310_i1 := _300_v1.F13; _310_i1.Cmp(_hi0) < 0; _310_i1 = _310_i1.Plus(_dafny.One) {
			r2 = p0
			var _311_v10 _dafny.Set
			_ = _311_v10
			_311_v10 = _dafny.SetOf(!(p0))
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_305_v6), 0))
			_ = _index44
			(_305_v6).ArraySet1((_311_v10).Cardinality(), (_index44).Int())
			var _312_v11 _dafny.Sequence
			_ = _312_v11
			_312_v11 = _dafny.UnicodeSeqOfUtf8Bytes("nurt")
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_305_v6), 0))
			_ = _index45
			(_305_v6).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_312_v11, _312_v11)).Cardinality()), _310_i1), (_index45).Int())
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_304_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_304_v5), 0))).Int()))
			_ = _arr0
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_dafny.ArrayCastTo((_304_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_304_v5), 0))).Int()))), 0))
			_ = _index46
			(_arr0).ArraySet1(p0, (_index46).Int())
			var _313_v12 _dafny.Map
			_ = _313_v12
			_313_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _299_v0)
			var _314_v13 _dafny.Sequence
			_ = _314_v13
			_314_v13 = _dafny.SeqOf(((_313_v12).Update(p0, _299_v0)).Cardinality(), _310_i1)
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_304_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_304_v5), 0))).Int()))
			_ = _arr1
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_dafny.ArrayCastTo((_304_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_304_v5), 0))).Int()))), 0))
			_ = _index47
			(_arr1).ArraySet1(!_dafny.Companion_Sequence_.Contains(_314_v13, _310_i1), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_305_v6), 0))
			_ = _index48
			(_305_v6).ArraySet1(((_dafny.IntOfInt64(355)).Times(p2)).Minus(_310_i1), (_index48).Int())
		}
		var _315_v14 _dafny.Map
		_ = _315_v14
		_315_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_302_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_302_v3).Cardinality()))).Uint32()).(_dafny.Array), _299_v0)
		_315_v14 = _315_v14
		var _316_v15 _dafny.Sequence
		_ = _316_v15
		_316_v15 = _dafny.SeqOf(p2, _300_v1.F13, _300_v1.F13, p2)
		var _317_v17 _dafny.Map
		_ = _317_v17
		_317_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Plus((_dafny.Zero).Minus(p2)), func() _dafny.Set {
			var _coll11 = _dafny.NewBuilder()
			_ = _coll11
			for _iter12 := _dafny.Iterate((_316_v15).Elements()); ; {
				_compr_11, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _318_v16 _dafny.Int
				_318_v16 = interface{}(_compr_11).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_316_v15, _318_v16) {
					_coll11.Add((_318_v16).Times(_dafny.IntOfInt64(679)))
				}
			}
			return _coll11.ToSet()
		}())
		var _319_v18 _dafny.Map
		_ = _319_v18
		_319_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(p0))
		var _320_v19 _dafny.Sequence
		_ = _320_v19
		_320_v19 = _dafny.SeqOf(p0)
		var _321_v20 _dafny.Set
		_ = _321_v20
		_321_v20 = _dafny.SetOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_319_v18).Contains(p0) {
				return (_319_v18).Get(p0).(_dafny.Sequence)
			}
			return _320_v19
		})()).Cardinality()))
		_317_v17 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_300_v1.F13, _321_v20)).Update(p2, _321_v20)
		var _322_v21 D2
		_ = _322_v21
		_322_v21 = Companion_D2_.Create_DC3_(p0)
		var _pat_let_tv3 = p0
		_ = _pat_let_tv3
		_322_v21 = (func() D2 {
			if p0 {
				return Companion_Default___.Fm11(p2, _300_v1.F13, func(_pat_let3_0 D2) D2 {
					return func(_323_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let4_0 bool) D2 {
							return func(_324_dt__update_hcf9_h0 bool) D2 {
								return Companion_D2_.Create_DC3_(_324_dt__update_hcf9_h0)
							}(_pat_let4_0)
						}(_pat_let_tv3)
					}(_pat_let3_0)
				}(_322_v21), p0, globalState)
			}
			return _322_v21
		})()
		var _325_v22 _dafny.Sequence
		_ = _325_v22
		_325_v22 = _dafny.UnicodeSeqOfUtf8Bytes("jw")
		var _326_v23 D1
		_ = _326_v23
		_326_v23 = Companion_D1_.Create_DC1_(_325_v22)
		var _327_v24 _dafny.Set
		_ = _327_v24
		_327_v24 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate((_326_v23).Dtor_cf5(), _325_v22), _dafny.UnicodeSeqOfUtf8Bytes("nujlcfog"), _325_v22)
		r0 = (_327_v24).Cardinality()
		r1 = _299_v0
		var _328_v25 _dafny.MultiSet
		_ = _328_v25
		_328_v25 = _dafny.MultiSetOf(p0)
		r2 = !((_dafny.MultiSetOf(p0, p0)).Equals(_328_v25))
		return r0, r1, r2
	}
}
func (_this *C3) M3(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		if false {
			var _329_v0 _dafny.MultiSet
			_ = _329_v0
			_329_v0 = _dafny.MultiSetOf(_dafny.SeqOf(p0), _dafny.SeqOf(p0, p0))
			var _330_v1 *C1
			_ = _330_v1
			var _nw46 *C1 = New_C1_()
			_ = _nw46
			_nw46.Ctor__(_329_v0, p3, _dafny.CodePoint('a'), p3, p3)
			_330_v1 = _nw46
			var _331_v2 _dafny.Sequence
			_ = _331_v2
			_331_v2 = _dafny.SeqOf((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_dafny.SeqOf(_330_v1, _330_v1, _330_v1, _330_v1)).Cardinality()), _dafny.IntOfInt64(444), p2)
			var _332_v3 _dafny.Map
			_ = _332_v3
			_332_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_330_v1).F12())
			var _333_v4 _dafny.Sequence
			_ = _333_v4
			_333_v4 = _dafny.SeqOf(_dafny.IntOfUint32((_331_v2).Cardinality()), (_332_v3).Cardinality(), (_dafny.Zero).Minus(p0))
			var _334_v5 _dafny.MultiSet
			_ = _334_v5
			_334_v5 = _dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_331_v2).Cardinality()), _dafny.IntOfInt64(767)), _331_v2, _333_v4, _333_v4)
			var _335_v6 _dafny.CodePoint
			_ = _335_v6
			_335_v6 = _dafny.CodePoint('q')
			var _336_v7 _dafny.Sequence
			_ = _336_v7
			_336_v7 = _dafny.SeqOf(p3)
			var _337_v8 *C1
			_ = _337_v8
			var _nw47 *C1 = New_C1_()
			_ = _nw47
			_nw47.Ctor__((Companion_Default___.Fm12(p3, globalState)).Difference(_334_v5), true, _335_v6, (func() bool {
				if false {
					return (func() bool {
						if (_332_v3).Contains(p0) {
							return (_332_v3).Get(p0).(bool)
						}
						return !((_336_v7).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_336_v7).Cardinality()))).Uint32()).(bool))
					})()
				}
				return !(p3)
			})(), (_330_v1).F12())
			_337_v8 = _nw47
			var _338_v9 _dafny.MultiSet
			_ = _338_v9
			_338_v9 = _dafny.MultiSetOf(p3)
			var _339_v10 _dafny.MultiSet
			_ = _339_v10
			_339_v10 = _dafny.MultiSetOf((func() _dafny.MultiSet {
				if false {
					return _dafny.MultiSetOf((_330_v1).F12())
				}
				return _338_v9
			})())
			var _340_v11 _dafny.Sequence
			_ = _340_v11
			_340_v11 = _dafny.SeqOf(_339_v10)
			_339_v10 = (_340_v11).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_340_v11).Cardinality()))).Uint32()).(_dafny.MultiSet)
			var _341_v12 bool
			_ = _341_v12
			_341_v12 = false
			var _342_v13 _dafny.Set
			_ = _342_v13
			_342_v13 = _dafny.SetOf(_341_v12)
			var _343_v14 _dafny.Sequence
			_ = _343_v14
			_343_v14 = _dafny.SeqOf(_342_v13)
			var _rhs37 bool = !(p3)
			_ = _rhs37
			var _rhs38 _dafny.Int = ((func() _dafny.Set {
				if (_337_v8).F12() {
					return _dafny.SetOf((_330_v1).F12())
				}
				return (_343_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.IntOfUint32((_343_v14).Cardinality()))).Uint32()).(_dafny.Set)
			})()).Cardinality()
			_ = _rhs38
			_341_v12 = _rhs37
			r0 = _rhs38
			var _344_v15 _dafny.Array
			_ = _344_v15
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw48
			_344_v15 = _nw48
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_344_v15), 0))
			_ = _index49
			(_344_v15).ArraySet1(p0, (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_344_v15), 0))
			_ = _index50
			(_344_v15).ArraySet1(p0, (_index50).Int())
			_342_v13 = (_342_v13).Difference((_342_v13).Difference(_dafny.SetOf(false)))
		} else {
			var _345_v16 _dafny.Array
			_ = _345_v16
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_11
			var _nw49 _dafny.Array
			_ = _nw49
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw49 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = func(_346_i0 _dafny.Int) bool {
					return false
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw49 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw49).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw49).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_345_v16 = _nw49
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_345_v16), 0))
			_ = _index51
			(_345_v16).ArraySet1(true, (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_345_v16), 0))
			_ = _index52
			(_345_v16).ArraySet1(!(p3), (_index52).Int())
			var _347_v17 _dafny.Set
			_ = _347_v17
			_347_v17 = _dafny.SetOf(_dafny.IntOfInt64(-311))
			var _348_v18 _dafny.MultiSet
			_ = _348_v18
			_348_v18 = _dafny.MultiSetOf(p0, (_347_v17).Cardinality())
			var _349_v19 _dafny.Sequence
			_ = _349_v19
			_349_v19 = _dafny.SeqOf((_345_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_345_v16), 0))).Int()).(bool), (p2).Cmp((_348_v18).Cardinality()) <= 0, (_345_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_345_v16), 0))).Int()).(bool))
			(globalState).F5 = _dafny.IntOfUint32((_349_v19).Cardinality())
			var _350_v20 _dafny.Map
			_ = _350_v20
			_350_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(34))
			var _351_v21 _dafny.Map
			_ = _351_v21
			_351_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _350_v20)
			var _352_v23 _dafny.MultiSet
			_ = _352_v23
			_352_v23 = _dafny.MultiSetOf((func() _dafny.Map {
				if (_351_v21).Contains(p0) {
					return (_351_v21).Get(p0).(_dafny.Map)
				}
				return func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(139), _dafny.IntOfInt64(19))); ; {
						_compr_12, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _353_v22 _dafny.Int
						_353_v22 = interface{}(_compr_12).(_dafny.Int)
						if ((_dafny.IntOfInt64(139)).Cmp(_353_v22) <= 0) && ((_353_v22).Cmp(_dafny.IntOfInt64(19)) < 0) {
							_coll12.Add((_353_v22).Plus(p2), p0)
						}
					}
					return _coll12.ToMap()
				}()
			})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), _350_v20)
			var _354_v24 _dafny.Map
			_ = _354_v24
			_354_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm13(p2, globalState), (_352_v23).Cardinality())
			var _355_v25 _dafny.CodePoint
			_ = _355_v25
			_355_v25 = _dafny.CodePoint('f')
			_354_v24 = (_354_v24).Update((func() _dafny.CodePoint {
				if p3 {
					return _355_v25
				}
				return _355_v25
			})(), p0)
			(globalState).F5 = (Companion_Default___.SafeModuloInt(p2, p0)).Plus(p0)
			var _356_v26 _dafny.Sequence
			_ = _356_v26
			_356_v26 = _dafny.UnicodeSeqOfUtf8Bytes("cb")
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_345_v16), 0))
			_ = _index53
			(_345_v16).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(726))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_357_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_358_i1 _dafny.Int) _dafny.CodePoint {
					return _357_v25
				}
			})(_355_v25))), _356_v26), (_index53).Int())
		}
		var _359_v27 bool
		_ = _359_v27
		_359_v27 = false
		_359_v27 = (func() bool {
			if p3 {
				return _359_v27
			}
			return p3
		})()
		r0 = ((p2).Minus(_dafny.IntOfInt64(792))).Times(p2)
		var _360_i2 _dafny.Int
		_ = _360_i2
		_360_i2 = _dafny.Zero
		{
			for false {
				{
					if (_360_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_360_i2 = (_360_i2).Plus(_dafny.One)
					var _361_v28 _dafny.Sequence
					_ = _361_v28
					_361_v28 = _dafny.UnicodeSeqOfUtf8Bytes("rfvu")
					if (Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_361_v28).Cardinality()))).Cmp(p0) <= 0 {
						var _362_v29 _dafny.Sequence
						_ = _362_v29
						_362_v29 = _dafny.SeqOf(_dafny.IntOfInt64(990), _dafny.IntOfUint32((_361_v28).Cardinality()), p2)
						var _363_v30 _dafny.CodePoint
						_ = _363_v30
						_363_v30 = _dafny.CodePoint('v')
						var _364_v31 _dafny.Map
						_ = _364_v31
						_364_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_359_v27, _363_v30)
						var _365_v32 _dafny.Set
						_ = _365_v32
						_365_v32 = _dafny.SetOf((_364_v31).Cardinality(), p2, (_dafny.Zero).Minus(p0))
						var _366_v33 _dafny.MultiSet
						_ = _366_v33
						_366_v33 = _dafny.MultiSetOf((func() _dafny.Set {
							if !(p3) {
								return _365_v32
							}
							return _365_v32
						})(), _365_v32, _365_v32)
						var _rhs39 _dafny.Int = p0
						_ = _rhs39
						var _rhs40 _dafny.Sequence = _362_v29
						_ = _rhs40
						var _rhs41 _dafny.MultiSet = _dafny.MultiSetOf(_365_v32)
						_ = _rhs41
						var _lhs23 *GlobalState = globalState
						_ = _lhs23
						_lhs23.F5 = _rhs39
						_362_v29 = _rhs40
						_366_v33 = _rhs41
						var _367_v34 D2
						_ = _367_v34
						_367_v34 = Companion_D2_.Create_DC4_(_361_v28, _dafny.SeqOf(_359_v27), p2)
						_367_v34 = _367_v34
						var _rhs42 bool = p3
						_ = _rhs42
						var _rhs43 bool = p3
						_ = _rhs43
						_359_v27 = _rhs42
						_359_v27 = _rhs43
						var _368_v35 _dafny.Map
						_ = _368_v35
						_368_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_359_v27) || (_359_v27)), _361_v28)
						_368_v35 = (_368_v35).Merge(Companion_Default___.Fm14(p2, _363_v30, p2, globalState))
						var _369_v36 D1
						_ = _369_v36
						_369_v36 = Companion_D1_.Create_DC1_(_361_v28)
						r0 = Companion_Default___.Fm0(p3, (_369_v36).Dtor_cf5(), globalState)
					} else {
						_359_v27 = !(false) || ((func() bool {
							if _359_v27 {
								return _359_v27
							}
							return _359_v27
						})())
						var _370_v37 _dafny.Sequence
						_ = _370_v37
						_370_v37 = _dafny.SeqOf(p2)
						var _371_v38 _dafny.Set
						_ = _371_v38
						_371_v38 = _dafny.SetOf(_370_v37)
						var _372_v39 _dafny.Sequence
						_ = _372_v39
						_372_v39 = _dafny.SeqOf(p2, (_371_v38).Cardinality(), p2)
						var _373_v40 _dafny.Sequence
						_ = _373_v40
						_373_v40 = _dafny.SeqOf(!(false), p3)
						var _374_v41 _dafny.CodePoint
						_ = _374_v41
						_374_v41 = _dafny.CodePoint('k')
						var _375_v42 D1
						_ = _375_v42
						_375_v42 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqOf(p2, (_dafny.MultiSetFromSeq(_372_v39)).Cardinality(), (_dafny.MultiSetFromSeq(_373_v40)).Cardinality())).Cardinality()), p2, _374_v41)
						(globalState).F5 = (_375_v42).Dtor_cf7()
						_359_v27 = _359_v27
						var _376_v43 *C0
						_ = _376_v43
						var _nw50 *C0 = New_C0_()
						_ = _nw50
						_nw50.Ctor__(_dafny.IntOfInt64(722), _374_v41, _359_v27, _359_v27)
						_376_v43 = _nw50
						_376_v43 = _376_v43
						var _377_v44 D2
						_ = _377_v44
						_377_v44 = Companion_D2_.Create_DC3_(!(_359_v27))
						var _378_v45 _dafny.Array
						_ = _378_v45
						var _nwElement0_4 D2 = _377_v44
						_ = _nwElement0_4
						var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(15))
						_ = _nw51
						(_nw51).ArraySet1(_nwElement0_4, 0)
						(_nw51).ArraySet1(_377_v44, 1)
						(_nw51).ArraySet1(Companion_D2_.Create_DC3_(_359_v27), 2)
						(_nw51).ArraySet1(_377_v44, 3)
						(_nw51).ArraySet1(_377_v44, 4)
						(_nw51).ArraySet1(Companion_D2_.Create_DC3_(!(_359_v27)), 5)
						(_nw51).ArraySet1(_377_v44, 6)
						(_nw51).ArraySet1(_377_v44, 7)
						(_nw51).ArraySet1(_377_v44, 8)
						(_nw51).ArraySet1(_377_v44, 9)
						(_nw51).ArraySet1(_377_v44, 10)
						(_nw51).ArraySet1(_377_v44, 11)
						(_nw51).ArraySet1(_377_v44, 12)
						(_nw51).ArraySet1(_377_v44, 13)
						(_nw51).ArraySet1(_377_v44, 14)
						_378_v45 = _nw51
						var _379_v46 _dafny.MultiSet
						_ = _379_v46
						_379_v46 = _dafny.MultiSetOf(p2, _376_v43.F13)
						var _380_v47 _dafny.Map
						_ = _380_v47
						_380_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v45, (_376_v43.F13).Cmp((_379_v46).Cardinality()) == 0)
						var _381_v48 _dafny.Sequence
						_ = _381_v48
						_381_v48 = _dafny.SeqOf(_378_v45)
						_380_v47 = (_380_v47).Update((_381_v48).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_370_v37)).Cardinality(), _dafny.IntOfUint32((_381_v48).Cardinality()))).Uint32()).(_dafny.Array), _359_v27)
					}
					var _382_v49 _dafny.Array
					_ = _382_v49
					var _len0_12 _dafny.Int = _dafny.IntOfInt64(12)
					_ = _len0_12
					var _nw52 _dafny.Array
					_ = _nw52
					if _len0_12.Cmp(_dafny.Zero) == 0 {
						_nw52 = _dafny.NewArray(_len0_12)
					} else {
						var _init12 func(_dafny.Int) _dafny.Int = (func(_383_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_384_i3 _dafny.Int) _dafny.Int {
								return (_384_i3).Plus(_383_p2)
							}
						})(p2)
						_ = _init12
						var _element0_12 = _init12(_dafny.Zero)
						_ = _element0_12
						_nw52 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
						(_nw52).ArraySet1(_element0_12, 0)
						var _nativeLen0_12 = (_len0_12).Int()
						_ = _nativeLen0_12
						for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
							(_nw52).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
						}
					}
					_382_v49 = _nw52
					var _385_v50 _dafny.CodePoint
					_ = _385_v50
					_385_v50 = _dafny.CodePoint('g')
					var _386_v51 _dafny.Sequence
					_ = _386_v51
					_386_v51 = _dafny.SeqOf(_361_v28, _361_v28, _dafny.UnicodeSeqOfUtf8Bytes("molvhg"), _dafny.Companion_Sequence_.Update(_361_v28, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_361_v28).Cardinality()))).Uint32(), _385_v50))
					var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_382_v49), 0))
					_ = _index54
					(_382_v49).ArraySet1(_dafny.IntOfUint32((_386_v51).Cardinality()), (_index54).Int())
					var _387_v52 _dafny.Array
					_ = _387_v52
					var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
					_ = _nw53
					_387_v52 = _nw53
					var _388_v53 _dafny.MultiSet
					_ = _388_v53
					_388_v53 = _dafny.MultiSetOf((func() _dafny.Array {
						if !(_359_v27) {
							return _387_v52
						}
						return _387_v52
					})(), _387_v52, _387_v52, _387_v52)
					var _389_v54 _dafny.MultiSet
					_ = _389_v54
					_389_v54 = _dafny.MultiSetOf(p3, p3)
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_382_v49), 0))
					_ = _index55
					var _rhs44 _dafny.Int = (p2).Plus(p0)
					_ = _rhs44
					var _rhs45 _dafny.Int = p2
					_ = _rhs45
					var _rhs46 _dafny.MultiSet = (_388_v53).Difference(_388_v53)
					_ = _rhs46
					var _rhs47 bool = ((func() _dafny.MultiSet {
						if p3 {
							return _389_v54
						}
						return (_dafny.MultiSetOf(_359_v27, _359_v27)).Update(_359_v27, Companion_Default___.Abs(p2))
					})()).IsDisjointFrom(_389_v54)
					_ = _rhs47
					var _lhs24 _dafny.Array = _382_v49
					_ = _lhs24
					var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_382_v49), 0))
					_ = _lhs25
					var _lhs26 *GlobalState = globalState
					_ = _lhs26
					(_lhs24).ArraySet1(_rhs44, (_lhs25).Int())
					_lhs26.F5 = _rhs45
					_388_v53 = _rhs46
					_359_v27 = _rhs47
					var _390_v55 _dafny.Sequence
					_ = _390_v55
					_390_v55 = _dafny.SeqOf(_359_v27, !((_this).Fm3(globalState)), (p2).Cmp(_dafny.IntOfInt64(779)) == 0)
					_390_v55 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_390_v55, _390_v55))
					var _391_v56 _dafny.MultiSet
					_ = _391_v56
					_391_v56 = _dafny.MultiSetOf(p0)
					var _392_v57 _dafny.Sequence
					_ = _392_v57
					_392_v57 = _dafny.SeqOf((_dafny.Zero).Minus(p0), (_dafny.Zero).Minus((_391_v56).Cardinality()), p2)
					var _393_v58 _dafny.Sequence
					_ = _393_v58
					_393_v58 = _dafny.SeqOf(_392_v57, _392_v57, _392_v57)
					var _394_v59 *C1
					_ = _394_v59
					var _nw54 *C1 = New_C1_()
					_ = _nw54
					_nw54.Ctor__(_dafny.MultiSetFromSeq(_393_v58), !(p3), _385_v50, (p3) || (true), (_390_v55).Select((Companion_Default___.SafeIndex((_382_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_382_v49), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_390_v55).Cardinality()))).Uint32()).(bool))
					_394_v59 = _nw54
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _395_v60 _dafny.Array
		_ = _395_v60
		var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
		_ = _nw55
		_395_v60 = _nw55
		var _396_v61 _dafny.Array
		_ = _396_v61
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_13
		var _nw56 _dafny.Array
		_ = _nw56
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw56 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) bool = (func(_397_v27 bool) func(_dafny.Int) bool {
				return func(_398_i4 _dafny.Int) bool {
					return _397_v27
				}
			})(_359_v27)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw56 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw56).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw56).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_396_v61 = _nw56
		var _399_v62 _dafny.MultiSet
		_ = _399_v62
		_399_v62 = _dafny.MultiSetOf(_396_v61, _396_v61)
		var _400_v63 _dafny.Int
		_ = _400_v63
		var _401_v64 _dafny.CodePoint
		_ = _401_v64
		var _402_v65 bool
		_ = _402_v65
		var _out9 _dafny.Int
		_ = _out9
		var _out10 _dafny.CodePoint
		_ = _out10
		var _out11 bool
		_ = _out11
		_out9, _out10, _out11 = (_this).M2(!(p3), _395_v60, Companion_Default___.SafeModuloInt(p0, p0), _399_v62, globalState)
		_400_v63 = _out9
		_401_v64 = _out10
		_402_v65 = _out11
		var _403_v66 _dafny.Sequence
		_ = _403_v66
		_403_v66 = _dafny.SeqOf(_402_v65)
		var _hi1 _dafny.Int = _dafny.IntOfUint32((_403_v66).Cardinality())
		_ = _hi1
		for _404_i5 := p0; _404_i5.Cmp(_hi1) < 0; _404_i5 = _404_i5.Plus(_dafny.One) {
			var _405_v67 *C0
			_ = _405_v67
			var _nw57 *C0 = New_C0_()
			_ = _nw57
			_nw57.Ctor__(p0, _401_v64, _402_v65, _402_v65)
			_405_v67 = _nw57
			var _406_v68 _dafny.Map
			_ = _406_v68
			_406_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _401_v64)
			_406_v68 = (_406_v68).Update(Companion_Default___.SafeModuloInt(_404_i5, _dafny.IntOfInt64(579)), _401_v64)
			var _407_v69 _dafny.Array
			_ = _407_v69
			var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw58
			_407_v69 = _nw58
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_407_v69), 0))
			_ = _index56
			(_407_v69).ArraySet1((_400_v63).Minus(p2), (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_407_v69), 0))
			_ = _index57
			(_407_v69).ArraySet1(_400_v63, (_index57).Int())
			if _359_v27 {
				var _408_v70 _dafny.Sequence
				_ = _408_v70
				_408_v70 = _dafny.UnicodeSeqOfUtf8Bytes("xp")
				var _409_v71 _dafny.Map
				_ = _409_v71
				_409_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_408_v70).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-240))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}((func(_410_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_411_i6 _dafny.Int) _dafny.CodePoint {
						return _410_v64
					}
				})(_401_v64))), _408_v70))
				_409_v71 = (_409_v71).Merge(_409_v71)
				var _412_v72 _dafny.Set
				_ = _412_v72
				_412_v72 = _dafny.SetOf(true, false)
				var _413_v73 _dafny.Map
				_ = _413_v73
				_413_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xas"), ((_dafny.SetOf(_359_v27)).Intersection(_412_v72)).Cardinality())
				_413_v73 = (func() _dafny.Map {
					if true {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v70, _400_v63)).Merge((_413_v73).Update(_408_v70, _405_v67.F13))
					}
					return Companion_Default___.Fm15(_401_v64, _405_v67.F13, p3, Companion_Default___.Fm0(_359_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(171))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg30 _dafny.Int) interface{} {
							return coer30(arg30)
						}
					}((func(_414_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_415_i7 _dafny.Int) _dafny.CodePoint {
							return _414_v64
						}
					})(_401_v64))), globalState), globalState)
				})()
				var _416_v74 T0
				_ = _416_v74
				var _nw59 *C2 = New_C2_()
				_ = _nw59
				_nw59.Ctor__(_400_v63, _402_v65, _359_v27)
				_416_v74 = _nw59
				var _417_v75 _dafny.Set
				_ = _417_v75
				_417_v75 = _dafny.SetOf(_416_v74)
				var _418_v76 *C1
				_ = _418_v76
				var _nw60 *C1 = New_C1_()
				_ = _nw60
				_nw60.Ctor__(_dafny.MultiSetOf(Companion_Default___.Fm16((_dafny.MultiSetOf(true)).Cardinality(), _402_v65, _359_v27, p3, globalState), _dafny.SeqOf(_405_v67.F13, _405_v67.F13)), p3, _401_v64, (_403_v66).Select((Companion_Default___.SafeIndex((_407_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_407_v69), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_403_v66).Cardinality()))).Uint32()).(bool), (_dafny.SetOf(_416_v74, _416_v74)).IsSubsetOf(_417_v75))
				_418_v76 = _nw60
				var _419_v77 _dafny.Set
				_ = _419_v77
				_419_v77 = _dafny.SetOf(_405_v67)
				_359_v27 = (_dafny.SetOf(_405_v67)).IsSubsetOf(_419_v77)
				var _420_v78 _dafny.MultiSet
				_ = _420_v78
				_420_v78 = _dafny.MultiSetOf(!(_416_v74.F9()))
				_420_v78 = (_dafny.MultiSetOf(!((_416_v74).F8()))).Intersection(_420_v78)
			} else {
				(globalState).F5 = Companion_Default___.SafeDivisionInt(p0, (func() _dafny.Int {
					if p3 {
						return (_407_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_407_v69), 0))).Int()).(_dafny.Int)
					}
					return (_dafny.Zero).Minus((_407_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_407_v69), 0))).Int()).(_dafny.Int))
				})())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_396_v61), 0))
				_ = _index58
				(_396_v61).ArraySet1(!(_402_v65) || ((_this).Fm3(globalState)), (_index58).Int())
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_396_v61), 0))
				_ = _index59
				(_396_v61).ArraySet1(p3, (_index59).Int())
				var _421_v79 _dafny.Sequence
				_ = _421_v79
				_421_v79 = _dafny.SeqOf(p2, _400_v63)
				var _422_v80 _dafny.MultiSet
				_ = _422_v80
				_422_v80 = _dafny.MultiSetOf(_421_v79)
				var _423_v81 _dafny.Sequence
				_ = _423_v81
				_423_v81 = _dafny.UnicodeSeqOfUtf8Bytes("eopmhpt")
				var _424_v82 _dafny.MultiSet
				_ = _424_v82
				_424_v82 = _dafny.MultiSetOf(_423_v81)
				var _425_v83 _dafny.Map
				_ = _425_v83
				_425_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_407_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_407_v69), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-153))
				var _426_v84 *C1
				_ = _426_v84
				var _nw61 *C1 = New_C1_()
				_ = _nw61
				_nw61.Ctor__(_422_v80, (_424_v82).IsSubsetOf(_424_v82), _401_v64, (p2).Cmp((_425_v83).Cardinality()) <= 0, p3)
				_426_v84 = _nw61
				var _427_v85 _dafny.MultiSet
				_ = _427_v85
				_427_v85 = _dafny.MultiSetOf(_402_v65)
				var _428_v86 _dafny.Map
				_ = _428_v86
				_428_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(_402_v65))).Union(_427_v85), _407_v69)
				_428_v86 = (_428_v86).Update(Companion_Default___.Fm19(_403_v66, (_396_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_396_v61), 0))).Int()).(bool), globalState), _407_v69)
				_359_v27 = !((_405_v67).Fm5(!_dafny.Companion_Sequence_.Equal(_403_v66, _403_v66), globalState))
			}
		}
		var _429_v87 _dafny.Sequence
		_ = _429_v87
		_429_v87 = _dafny.SeqOf(p0, _400_v63, Companion_Default___.Fm0(_359_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(950))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_430_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_431_i8 _dafny.Int) _dafny.CodePoint {
				return _430_v64
			}
		})(_401_v64))), globalState), _400_v63, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _400_v63)).Cardinality())
		var _432_v88 D4
		_ = _432_v88
		_432_v88 = Companion_D4_.Create_DC7_(_429_v87)
		r0 = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if p3 {
				return _dafny.IntOfUint32(((_432_v88).Dtor_cf15()).Cardinality())
			}
			return (_dafny.MultiSetOf(_401_v64)).Cardinality()
		})()))
		r1 = p0
		return r0, r1
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f7 _dafny.Sequence
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f7 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C4_ struct {
}

var Companion_C4_ = CompanionStruct_C4_{}

func (_this *C4) Equals(other *C4) bool {
	return _this == other
}

func (_this *C4) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C4)
	return ok && _this.Equals(other)
}

func (*C4) String() string {
	return "_module.C4"
}

func Type_C4_() _dafny.TypeDescriptor {
	return type_C4_{}
}

type type_C4_ struct {
}

func (_this type_C4_) Default() interface{} {
	return (*C4)(nil)
}

func (_this type_C4_) String() string {
	return "main.C4"
}
func (_this *C4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__(f7 _dafny.Sequence) {
	{
		(_this)._f7 = f7
	}
}
func (_this *C4) Fm1(p0 bool, p1 bool, globalState *GlobalState) bool {
	{
		return !(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("biuajrf"), _dafny.UnicodeSeqOfUtf8Bytes("tnja")))
	}
}
func (_this *C4) M1(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		var _433_v0 _dafny.CodePoint
		_ = _433_v0
		_433_v0 = _dafny.CodePoint('j')
		_433_v0 = _433_v0
		var _hi2 _dafny.Int = (_dafny.Zero).Minus(p2)
		_ = _hi2
		for _434_i0 := _dafny.IntOfInt64(198); _434_i0.Cmp(_hi2) < 0; _434_i0 = _434_i0.Plus(_dafny.One) {
			var _435_v1 _dafny.Array
			_ = _435_v1
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_14
			var _nw62 _dafny.Array
			_ = _nw62
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw62 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Map = (func(_436_p3 bool, _437_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_438_i1 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_436_p3), _437_p0)
					}
				})(p3, p0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw62 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw62).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw62).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_435_v1 = _nw62
			var _439_v2 _dafny.Map
			_ = _439_v2
			_439_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
			var _440_v3 _dafny.Sequence
			_ = _440_v3
			_440_v3 = _dafny.SeqOf(p3)
			var _441_v4 _dafny.Array
			_ = _441_v4
			var _nwElement0_5 bool = p3
			_ = _nwElement0_5
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(13))
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_5, 0)
			(_nw63).ArraySet1(p3, 1)
			(_nw63).ArraySet1((func() bool {
				if (_439_v2).Contains(_dafny.IntOfUint32((_440_v3).Cardinality())) {
					return (_439_v2).Get(_dafny.IntOfUint32((_440_v3).Cardinality())).(bool)
				}
				return !(p3)
			})(), 2)
			(_nw63).ArraySet1(p3, 3)
			(_nw63).ArraySet1(p3, 4)
			(_nw63).ArraySet1(p3, 5)
			(_nw63).ArraySet1(p3, 6)
			(_nw63).ArraySet1(p3, 7)
			(_nw63).ArraySet1(p3, 8)
			(_nw63).ArraySet1(p3, 9)
			(_nw63).ArraySet1(p3, 10)
			(_nw63).ArraySet1(false, 11)
			(_nw63).ArraySet1(p3, 12)
			_441_v4 = _nw63
			var _442_v5 _dafny.MultiSet
			_ = _442_v5
			_442_v5 = _dafny.MultiSetOf(_441_v4, _441_v4, _441_v4, _441_v4, _441_v4)
			var _rhs48 _dafny.Array = _435_v1
			_ = _rhs48
			var _rhs49 _dafny.MultiSet = _442_v5
			_ = _rhs49
			var _rhs50 _dafny.Int = _434_i0
			_ = _rhs50
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			_435_v1 = _rhs48
			_442_v5 = _rhs49
			_lhs27.F5 = _rhs50
			if p3 {
				var _443_v6 _dafny.Map
				_ = _443_v6
				_443_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), p0)
				var _444_v7 _dafny.Set
				_ = _444_v7
				_444_v7 = _dafny.SetOf((_dafny.Zero).Minus(p0), p2)
				_443_v6 = (_443_v6).Update(p3, (_444_v7).Cardinality())
				var _445_v8 bool
				_ = _445_v8
				_445_v8 = false
				_445_v8 = (_440_v3).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p2), p0), _dafny.IntOfUint32((_440_v3).Cardinality()))).Uint32()).(bool)
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_441_v4), 0))
				_ = _index60
				(_441_v4).ArraySet1(true, (_index60).Int())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_441_v4), 0))
				_ = _index61
				(_441_v4).ArraySet1((func() bool {
					if p3 {
						return _445_v8
					}
					return p3
				})(), (_index61).Int())
				var _446_v9 _dafny.Array
				_ = _446_v9
				var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw64
				_446_v9 = _nw64
				var _447_v10 _dafny.Map
				_ = _447_v10
				_447_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _446_v9)
				var _448_v11 D1
				_ = _448_v11
				_448_v11 = Companion_D1_.Create_DC1_(_dafny.Companion_Sequence_.Update((_this).F7(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((_this).F7()).Cardinality()))).Uint32(), ((_this).F7()).Select((Companion_Default___.SafeIndex(_434_i0, _dafny.IntOfUint32(((_this).F7()).Cardinality()))).Uint32()).(_dafny.CodePoint)))
				var _449_v12 _dafny.MultiSet
				_ = _449_v12
				_449_v12 = _dafny.MultiSetOf(_445_v8)
				var _450_v13 _dafny.Sequence
				_ = _450_v13
				_450_v13 = _dafny.SeqOf(p0)
				var _451_v14 _dafny.Sequence
				_ = _451_v14
				_451_v14 = _dafny.SeqOf(_440_v3)
				var _452_v15 _dafny.Sequence
				_ = _452_v15
				_452_v15 = _dafny.SeqOf((_451_v14).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_451_v14).Cardinality()))).Uint32()).(_dafny.Sequence), _440_v3)
				var _453_v16 _dafny.Array
				_ = _453_v16
				var _nwElement0_6 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (_447_v10).Cardinality())
				_ = _nwElement0_6
				var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(22))
				_ = _nw65
				(_nw65).ArraySet1(_nwElement0_6, 0)
				(_nw65).ArraySet1(_dafny.IntOfUint32(((_this).F7()).Cardinality()), 1)
				(_nw65).ArraySet1(Companion_Default___.SafeModuloInt(p1, _434_i0), 2)
				(_nw65).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F7(), (_448_v11).Dtor_cf5())).Cardinality()), 3)
				(_nw65).ArraySet1(p2, 4)
				(_nw65).ArraySet1((p1).Plus(Companion_Default___.Fm0(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-470))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_454_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_455_i2 _dafny.Int) _dafny.CodePoint {
						return _454_v0
					}
				})(_433_v0))), globalState)), 5)
				(_nw65).ArraySet1(_434_i0, 6)
				(_nw65).ArraySet1(p1, 7)
				(_nw65).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm0(p3, (_this).F7(), globalState)), 8)
				(_nw65).ArraySet1((_dafny.Zero).Minus((_434_i0).Plus(p0)), 9)
				(_nw65).ArraySet1(p2, 10)
				(_nw65).ArraySet1(_434_i0, 11)
				(_nw65).ArraySet1(_dafny.IntOfInt64(748), 12)
				(_nw65).ArraySet1(p0, 13)
				(_nw65).ArraySet1(p2, 14)
				(_nw65).ArraySet1((func() _dafny.Int {
					if (_449_v12).Contains(p3) {
						return (_449_v12).Multiplicity(p3)
					}
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_450_v13, (Companion_Default___.SafeIndex((_439_v2).Cardinality(), _dafny.IntOfUint32((_450_v13).Cardinality()))).Uint32(), (_dafny.Zero).Minus((_449_v12).Cardinality()))).Cardinality())
				})(), 15)
				(_nw65).ArraySet1(p2, 16)
				(_nw65).ArraySet1(_434_i0, 17)
				(_nw65).ArraySet1(_434_i0, 18)
				(_nw65).ArraySet1(_dafny.IntOfUint32(((_this).F7()).Cardinality()), 19)
				(_nw65).ArraySet1(_dafny.IntOfUint32((_452_v15).Cardinality()), 20)
				(_nw65).ArraySet1(Companion_Default___.Fm0(!(!(_445_v8)), (_this).F7(), globalState), 21)
				_453_v16 = _nw65
				_453_v16 = _446_v9
				var _456_v17 _dafny.Array
				_ = _456_v17
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
				_ = _nw66
				_456_v17 = _nw66
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_456_v17), 0))
				_ = _index62
				(_456_v17).ArraySet1((_this).F7(), (_index62).Int())
				var _457_v18 _dafny.Map
				_ = _457_v18
				_457_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate((_this).F7(), (_this).F7()), (_this).F7())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_456_v17), 0))
				_ = _index63
				var _rhs51 _dafny.Int = (_dafny.Zero).Minus(p1)
				_ = _rhs51
				var _rhs52 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_this).F7(), (Companion_Default___.SafeIndex(_434_i0, _dafny.IntOfUint32(((_this).F7()).Cardinality()))).Uint32(), ((_this).F7()).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_this).F7()).Cardinality()))).Uint32()).(_dafny.CodePoint))
				_ = _rhs52
				var _rhs53 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_440_v3, _440_v3)
				_ = _rhs53
				var _rhs54 bool = _445_v8
				_ = _rhs54
				var _rhs55 _dafny.Map = (_457_v18).Merge(Companion_Default___.Fm2(p2, (_441_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_441_v4), 0))).Int()).(bool), p3, globalState))
				_ = _rhs55
				var _lhs28 *GlobalState = globalState
				_ = _lhs28
				var _lhs29 _dafny.Array = _456_v17
				_ = _lhs29
				var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_456_v17), 0))
				_ = _lhs30
				_lhs28.F5 = _rhs51
				(_lhs29).ArraySet1(_rhs52, (_lhs30).Int())
				_440_v3 = _rhs53
				_445_v8 = _rhs54
				_457_v18 = _rhs55
			} else {
				var _458_v19 bool
				_ = _458_v19
				_458_v19 = true
				_458_v19 = !(_458_v19)
				var _459_v20 _dafny.Array
				_ = _459_v20
				var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw67
				_459_v20 = _nw67
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_459_v20), 0))
				_ = _index64
				(_459_v20).ArraySet1((_dafny.Zero).Minus(p2), (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_459_v20), 0))
				_ = _index65
				(_459_v20).ArraySet1(Companion_Default___.SafeDivisionInt(p1, (_dafny.Zero).Minus((_434_i0).Plus(_dafny.IntOfUint32(((_this).F7()).Cardinality())))), (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_459_v20), 0))
				_ = _index66
				(_459_v20).ArraySet1((_459_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_459_v20), 0))).Int()).(_dafny.Int), (_index66).Int())
				var _460_v21 _dafny.MultiSet
				_ = _460_v21
				_460_v21 = _dafny.MultiSetOf(_458_v19)
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_441_v4), 0))
				_ = _index67
				(_441_v4).ArraySet1((_460_v21).Contains(_458_v19), (_index67).Int())
				var _461_v22 D2
				_ = _461_v22
				_461_v22 = Companion_D2_.Create_DC3_(p3)
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_441_v4), 0))
				_ = _index68
				(_441_v4).ArraySet1((func(_pat_let5_0 D2) D2 {
					return func(_462_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let6_0 bool) D2 {
							return func(_463_dt__update_hcf9_h0 bool) D2 {
								return Companion_D2_.Create_DC3_(_463_dt__update_hcf9_h0)
							}(_pat_let6_0)
						}(true)
					}(_pat_let5_0)
				}(_461_v22)).Dtor_cf9(), (_index68).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_441_v4), 0))
				_ = _index69
				(_441_v4).ArraySet1((p2).Cmp((_dafny.IntOfUint32((_dafny.SeqOf(!((_this).Fm1(_458_v19, _458_v19, globalState)))).Cardinality())).Plus(_dafny.IntOfUint32(((_this).F7()).Cardinality()))) >= 0, (_index69).Int())
			}
			var _464_v23 _dafny.Array
			_ = _464_v23
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw68
			_464_v23 = _nw68
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _index70
			(_464_v23).ArraySet1(_dafny.IntOfInt64(873), (_index70).Int())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _index71
			(_464_v23).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if p3 {
					return Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32(((_this).F7()).Cardinality()))
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll13 = _dafny.NewMapBuilder()
					_ = _coll13
					for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(397), _dafny.IntOfInt64(852))); ; {
						_compr_13, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _465_v24 _dafny.Int
						_465_v24 = interface{}(_compr_13).(_dafny.Int)
						if ((_dafny.IntOfInt64(397)).Cmp(_465_v24) <= 0) && ((_465_v24).Cmp(_dafny.IntOfInt64(852)) < 0) {
							_coll13.Add((_465_v24).Minus(p1), (_440_v3).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_440_v3).Cardinality()))).Uint32()).(bool))
						}
					}
					return _coll13.ToMap()
				}(), p2)).Cardinality()
			})()), (_index71).Int())
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _index72
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _index73
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _index74
			var _rhs56 _dafny.Int = (_dafny.Zero).Minus((_464_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))).Int()).(_dafny.Int))
			_ = _rhs56
			var _rhs57 _dafny.Int = (_464_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))).Int()).(_dafny.Int)
			_ = _rhs57
			var _rhs58 _dafny.Int = _434_i0
			_ = _rhs58
			var _rhs59 _dafny.Int = p2
			_ = _rhs59
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			var _lhs32 _dafny.Array = _464_v23
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _lhs33
			var _lhs34 _dafny.Array = _464_v23
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _lhs35
			var _lhs36 _dafny.Array = _464_v23
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_464_v23), 0))
			_ = _lhs37
			_lhs31.F5 = _rhs56
			(_lhs32).ArraySet1(_rhs57, (_lhs33).Int())
			(_lhs34).ArraySet1(_rhs58, (_lhs35).Int())
			(_lhs36).ArraySet1(_rhs59, (_lhs37).Int())
		}
		var _466_v25 _dafny.Array
		_ = _466_v25
		var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
		_ = _nw69
		_466_v25 = _nw69
		for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_466_v25), 0))); ; {
			_guard_loop_1, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _467_i3 _dafny.Int
			_467_i3 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_467_i3).Sign() != -1) && ((_467_i3).Cmp(_dafny.ArrayLen((_466_v25), 0)) < 0)) {
				(_466_v25).ArraySet1((func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter16 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, p3)).Cardinality()), p0, p2)).Elements()); ; {
						_compr_14, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _468_v26 _dafny.Int
						_468_v26 = interface{}(_compr_14).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, p3)).Cardinality()), p0, p2), _468_v26) {
							_coll14.Add((_468_v26).Minus(p2), p3)
						}
					}
					return _coll14.ToMap()
				}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)), (_467_i3).Int())
			}
		}
		(globalState).F5 = p1
		if false {
			var _469_v27 *C0
			_ = _469_v27
			var _nw70 *C0 = New_C0_()
			_ = _nw70
			_nw70.Ctor__(Companion_Default___.SafeDivisionInt(p1, p1), _dafny.CodePoint('g'), ((_this).Fm1(p3, p3, globalState)) && (p3), (_this).Fm1(p3, p3, globalState))
			_469_v27 = _nw70
			if p3 {
				(globalState).F5 = (_469_v27.F13).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-130))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_470_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_471_i4 _dafny.Int) _dafny.Int {
						return _470_p1
					}
				})(p1)))).Cardinality()))
				var _472_v28 bool
				_ = _472_v28
				_472_v28 = true
				_472_v28 = _472_v28
				var _473_v29 _dafny.Sequence
				_ = _473_v29
				_473_v29 = _dafny.UnicodeSeqOfUtf8Bytes("xmevn")
				_473_v29 = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("njdcm"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("njdcm")).Cardinality()))).Uint32(), Companion_Default___.Fm13(p1, globalState))
				var _474_v30 _dafny.MultiSet
				_ = _474_v30
				_474_v30 = _dafny.MultiSetOf(_433_v0, _433_v0)
				var _475_v31 _dafny.Map
				_ = _475_v31
				_475_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
				var _476_v32 _dafny.Set
				_ = _476_v32
				_476_v32 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true), _475_v31)
				(_469_v27).F13 = ((func() _dafny.Int {
					if p3 {
						return (_474_v30).Cardinality()
					}
					return (_476_v32).Cardinality()
				})()).Plus(Companion_Default___.SafeModuloInt(_469_v27.F13, (_dafny.Zero).Minus(p0)))
				var _477_v33 _dafny.Array
				_ = _477_v33
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw71
				_477_v33 = _nw71
				var _478_v34 _dafny.Map
				_ = _478_v34
				_478_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_477_v33, _472_v28)
				var _479_v35 _dafny.Sequence
				_ = _479_v35
				_479_v35 = _dafny.SeqOf(_472_v28)
				var _480_v36 *C2
				_ = _480_v36
				var _nw72 *C2 = New_C2_()
				_ = _nw72
				_nw72.Ctor__(p2, true, (func() bool {
					if (_478_v34).Contains(_477_v33) {
						return (_478_v34).Get(_477_v33).(bool)
					}
					return (_479_v35).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_479_v35).Cardinality()))).Uint32()).(bool)
				})())
				_480_v36 = _nw72
			} else {
				var _481_v37 _dafny.Sequence
				_ = _481_v37
				_481_v37 = _dafny.SeqOf(_469_v27.F13)
				_481_v37 = _dafny.SeqOf(_469_v27.F13, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(p3)).Cardinality()), _469_v27.F13), p1)
				var _482_v38 bool
				_ = _482_v38
				_482_v38 = true
				_482_v38 = _482_v38
				_482_v38 = p3
				var _483_v39 _dafny.Array
				_ = _483_v39
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_15
				var _nw73 _dafny.Array
				_ = _nw73
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw73 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Int = (func(_484_v27 *C0) func(_dafny.Int) _dafny.Int {
						return func(_485_i5 _dafny.Int) _dafny.Int {
							return (_485_i5).Times(_484_v27.F13)
						}
					})(_469_v27)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw73 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw73).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw73).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_483_v39 = _nw73
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_483_v39), 0))
				_ = _index75
				(_483_v39).ArraySet1((_dafny.Zero).Minus(p1), (_index75).Int())
				var _486_v40 _dafny.Set
				_ = _486_v40
				_486_v40 = _dafny.SetOf(_482_v38, !(p3), _482_v38, _482_v38)
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_483_v39), 0))
				_ = _index76
				(_483_v39).ArraySet1((((func() _dafny.Set {
					if p3 {
						return _486_v40
					}
					return _486_v40
				})()).Cardinality()).Times(Companion_Default___.Fm0(p3, (_this).F7(), globalState)), (_index76).Int())
				var _487_v41 _dafny.Array
				_ = _487_v41
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
				_ = _nw74
				_487_v41 = _nw74
				_487_v41 = _487_v41
			}
			var _488_v42 *C2
			_ = _488_v42
			var _nw75 *C2 = New_C2_()
			_ = _nw75
			_nw75.Ctor__(_dafny.IntOfInt64(-178), true, p3)
			_488_v42 = _nw75
			var _489_v43 _dafny.Map
			_ = _489_v43
			_489_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _488_v42)
			var _490_v44 _dafny.Map
			_ = _490_v44
			_490_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if p3 {
					return _489_v43
				}
				return _489_v43
			})(), p3)
			var _491_v45 _dafny.Set
			_ = _491_v45
			_491_v45 = _dafny.SetOf(false, p3)
			_490_v44 = (_490_v44).Update(_489_v43, (_491_v45).IsProperSubsetOf(_491_v45))
			_433_v0 = _433_v0
			var _492_v46 _dafny.Array
			_ = _492_v46
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_16
			var _nw76 _dafny.Array
			_ = _nw76
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw76 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Int = (func(_493_v27 *C0, _494_p3 bool) func(_dafny.Int) _dafny.Int {
					return func(_495_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_495_i6, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_493_v27.F13), _dafny.SeqOf(_494_p3))).Cardinality())
					}
				})(_469_v27, p3)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw76 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw76).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw76).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_492_v46 = _nw76
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_492_v46), 0))
			_ = _index77
			(_492_v46).ArraySet1(p1, (_index77).Int())
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_492_v46), 0))
			_ = _index78
			(_492_v46).ArraySet1(p0, (_index78).Int())
		} else {
			if (p2).Cmp((_dafny.Zero).Minus(p1)) != 0 {
				var _496_v47 bool
				_ = _496_v47
				_496_v47 = false
				var _497_v48 _dafny.Sequence
				_ = _497_v48
				_497_v48 = _dafny.SeqOf(p3, p3)
				_496_v47 = (_496_v47) || ((_497_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F7()).Cardinality()), _dafny.IntOfUint32((_497_v48).Cardinality()))).Uint32()).(bool))
				_496_v47 = _496_v47
				var _498_v49 _dafny.Set
				_ = _498_v49
				_498_v49 = _dafny.SetOf((_this).Fm1(_496_v47, _496_v47, globalState))
				var _499_v50 D1
				_ = _499_v50
				_499_v50 = Companion_D1_.Create_DC2_((_498_v49).Cardinality(), p2, _dafny.CodePoint('c'))
				var _500_v51 _dafny.MultiSet
				_ = _500_v51
				_500_v51 = _dafny.MultiSetOf(Companion_D1_.Create_DC2_(p2, _dafny.IntOfInt64(646), _433_v0))
				_496_v47 = (_dafny.MultiSetOf(_499_v50, _499_v50)).IsDisjointFrom((_500_v51).Union(_dafny.MultiSetOf(_499_v50)))
				_496_v47 = _496_v47
				var _501_v52 *C3
				_ = _501_v52
				var _nw77 *C3 = New_C3_()
				_ = _nw77
				_nw77.Ctor__()
				_501_v52 = _nw77
			} else {
				(globalState).F5 = p0
				var _502_v53 _dafny.Map
				_ = _502_v53
				_502_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() _dafny.CodePoint {
					if p3 {
						return _dafny.CodePoint('o')
					}
					return _433_v0
				})())
				var _503_v54 _dafny.Sequence
				_ = _503_v54
				_503_v54 = _dafny.SeqOf(p3, p3)
				var _504_v55 _dafny.Sequence
				_ = _504_v55
				_504_v55 = _dafny.SeqOf(p2)
				var _505_v56 _dafny.Map
				_ = _505_v56
				_505_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_504_v55).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_504_v55).Cardinality()))).Uint32()).(_dafny.Int), p3)
				var _506_v57 _dafny.Array
				_ = _506_v57
				var _nwElement0_7 _dafny.Int = (_505_v56).Cardinality()
				_ = _nwElement0_7
				var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(24))
				_ = _nw78
				(_nw78).ArraySet1(_nwElement0_7, 0)
				(_nw78).ArraySet1(p0, 1)
				(_nw78).ArraySet1(p2, 2)
				(_nw78).ArraySet1(p0, 3)
				(_nw78).ArraySet1(p2, 4)
				(_nw78).ArraySet1(p1, 5)
				(_nw78).ArraySet1(p0, 6)
				(_nw78).ArraySet1(_dafny.IntOfUint32(((_this).F7()).Cardinality()), 7)
				(_nw78).ArraySet1(p0, 8)
				(_nw78).ArraySet1(p0, 9)
				(_nw78).ArraySet1(p0, 10)
				(_nw78).ArraySet1(p2, 11)
				(_nw78).ArraySet1(p0, 12)
				(_nw78).ArraySet1(p0, 13)
				(_nw78).ArraySet1(p1, 14)
				(_nw78).ArraySet1((Companion_Default___.Fm20(p3, globalState)).Cardinality(), 15)
				(_nw78).ArraySet1(p2, 16)
				(_nw78).ArraySet1(p1, 17)
				(_nw78).ArraySet1(p0, 18)
				(_nw78).ArraySet1(p2, 19)
				(_nw78).ArraySet1(p1, 20)
				(_nw78).ArraySet1(p1, 21)
				(_nw78).ArraySet1(p2, 22)
				(_nw78).ArraySet1(p2, 23)
				_506_v57 = _nw78
				var _507_v58 _dafny.Set
				_ = _507_v58
				_507_v58 = _dafny.SetOf(_506_v57, _506_v57, _506_v57)
				_502_v53 = (_502_v53).Update((Companion_D2_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(839))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}((func(_508_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_509_i7 _dafny.Int) _dafny.CodePoint {
						return _508_v0
					}
				})(_433_v0))), _503_v54, (_507_v58).Cardinality())).Dtor_cf12(), (func() _dafny.CodePoint {
					if p3 {
						return _433_v0
					}
					return _433_v0
				})())
				(globalState).F5 = (_dafny.IntOfUint32(((_this).F7()).Cardinality())).Plus((func() _dafny.Int {
					if false {
						return Companion_Default___.Fm0(p3, (_this).F7(), globalState)
					}
					return _dafny.IntOfInt64(711)
				})())
				var _510_v59 D1
				_ = _510_v59
				_510_v59 = Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("e"))
				var _511_v60 _dafny.Sequence
				_ = _511_v60
				_511_v60 = _dafny.SeqOf((_this).F7())
				_510_v59 = Companion_D1_.Create_DC1_((_511_v60).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_511_v60).Cardinality()))).Uint32()).(_dafny.Sequence))
				(globalState).F5 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(p0, p1), (_dafny.IntOfUint32(((_this).F7()).Cardinality())).Minus(p0))
			}
			(globalState).F5 = _dafny.IntOfInt64(697)
			var _512_v61 _dafny.Map
			_ = _512_v61
			_512_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xbjxfiomp"), (_this).F7())
			var _513_v62 _dafny.Set
			_ = _513_v62
			_513_v62 = _dafny.SetOf(p3, true)
			var _514_v63 _dafny.Sequence
			_ = _514_v63
			_514_v63 = _dafny.SeqOf(p2, p2)
			var _515_v64 _dafny.MultiSet
			_ = _515_v64
			_515_v64 = _dafny.MultiSetOf(_514_v63, _dafny.SeqOf(p0))
			var _516_v65 *C1
			_ = _516_v65
			var _nw79 *C1 = New_C1_()
			_ = _nw79
			_nw79.Ctor__(_515_v64, p3, _433_v0, p3, p3)
			_516_v65 = _nw79
			var _517_v66 D4
			_ = _517_v66
			_517_v66 = Companion_D4_.Create_DC8_(p2, p3, p0, _516_v65)
			var _518_v67 _dafny.Array
			_ = _518_v67
			var _nwElement0_8 _dafny.Int = (_513_v62).Cardinality()
			_ = _nwElement0_8
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_8, 0)
			(_nw80).ArraySet1(p0, 1)
			(_nw80).ArraySet1(p1, 2)
			(_nw80).ArraySet1(p2, 3)
			(_nw80).ArraySet1(p2, 4)
			(_nw80).ArraySet1((_dafny.Zero).Minus(p2), 5)
			(_nw80).ArraySet1(p0, 6)
			(_nw80).ArraySet1(p2, 7)
			(_nw80).ArraySet1((_dafny.Zero).Minus(p1), 8)
			(_nw80).ArraySet1(p0, 9)
			(_nw80).ArraySet1((_dafny.Zero).Minus((_517_v66).Dtor_cf18()), 10)
			(_nw80).ArraySet1(_dafny.IntOfInt64(508), 11)
			(_nw80).ArraySet1(p2, 12)
			(_nw80).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ll")).Cardinality()), 13)
			_518_v67 = _nw80
			var _519_v68 _dafny.Map
			_ = _519_v68
			_519_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _516_v65)
			var _520_v69 _dafny.Array
			_ = _520_v69
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_17
			var _nw81 _dafny.Array
			_ = _nw81
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw81 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.CodePoint = (func(_521_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_522_i8 _dafny.Int) _dafny.CodePoint {
						return _521_v0
					}
				})(_433_v0)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw81 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw81).ArraySet1CodePoint(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw81).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_520_v69 = _nw81
			var _523_v70 D0
			_ = _523_v70
			_523_v70 = Companion_D0_.Create_DC0_(_512_v61, _518_v67, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_516_v65).F12(), (_519_v68).Cardinality()), _520_v69, _518_v67)
			var _524_v71 _dafny.Map
			_ = _524_v71
			_524_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _523_v70)
			var _525_v72 _dafny.Array
			_ = _525_v72
			var _nwElement0_9 _dafny.Map = _524_v71
			_ = _nwElement0_9
			var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
			_ = _nw82
			(_nw82).ArraySet1(_nwElement0_9, 0)
			(_nw82).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _523_v70), 1)
			(_nw82).ArraySet1((_524_v71).Update(p3, _523_v70), 2)
			(_nw82).ArraySet1(_524_v71, 3)
			(_nw82).ArraySet1(_524_v71, 4)
			(_nw82).ArraySet1(_524_v71, 5)
			(_nw82).ArraySet1((_524_v71).Merge(_524_v71), 6)
			_525_v72 = _nw82
			_525_v72 = _525_v72
			var _526_v73 _dafny.Map
			_ = _526_v73
			_526_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), p2)
			var _527_v74 _dafny.Set
			_ = _527_v74
			_527_v74 = _dafny.SetOf(_526_v73)
			var _528_v75 _dafny.MultiSet
			_ = _528_v75
			_528_v75 = _dafny.MultiSetOf(false)
			var _529_v76 _dafny.Int
			_ = _529_v76
			var _out12 _dafny.Int
			_ = _out12
			_out12 = Companion_Default___.M0(!((_516_v65).F12()) || (p3), (_527_v74).Difference(_527_v74), !(_528_v75).Contains(p3), globalState)
			_529_v76 = _out12
			var _530_v77 bool
			_ = _530_v77
			_530_v77 = true
			_530_v77 = (_this).Fm1(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(607)), _529_v76), (_516_v65).F12(), globalState)
		}
		var _531_v78 _dafny.Map
		_ = _531_v78
		_531_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		_531_v78 = (_531_v78).Update((_dafny.Zero).Minus(p2), p1)
	}
}
func (_this *C4) F7() _dafny.Sequence {
	{
		return _this._f7
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
