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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool {
	if !(false) {
		return !(true)
	} else {
		return (_dafny.IntOfInt64(761)).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality())) == 0
	}
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vvsslon"), _dafny.UnicodeSeqOfUtf8Bytes("iphflypt")), _dafny.UnicodeSeqOfUtf8Bytes("qc"))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.SetOf(false, true)).Cardinality()), _dafny.IntOfInt64(-55), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("suxyudnjw")).Cardinality()))).Union(_dafny.MultiSetOf((_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()))).Cardinality())).Cardinality()))).Difference((_dafny.MultiSetOf(_dafny.IntOfInt64(367))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(265)))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(676))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("byditj")))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D1_.Create_DC4_(_dafny.SeqOf(true, true), false)).Dtor_cf7(), _dafny.UnicodeSeqOfUtf8Bytes("s")))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(619))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xnuinsd")).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
	}))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true)), (_dafny.SetOf(_dafny.IntOfInt64(418), (_dafny.MultiSetOf(_dafny.IntOfInt64(-592))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, false, true)).Cardinality()))).IsSubsetOf(_dafny.SetOf((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(604), _dafny.IntOfInt64(391))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(604)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(391)) < 0) {
				_coll0.Add((_2_v0).Plus(_dafny.IntOfInt64(262)))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(893))))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-274), _dafny.IntOfInt64(697))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Int
			_3_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(-274)).Cmp(_3_v0) <= 0) && ((_3_v0).Cmp(_dafny.IntOfInt64(697)) < 0) {
				_coll1.Add(Companion_Default___.SafeModuloInt(_3_v0, _dafny.IntOfInt64(59)), false)
			}
		}
		return _coll1.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(464), false))).Intersection(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("x")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aqb")).Cardinality()))).Cardinality(), !(true))))).Difference((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(965), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true, false))).Cardinality()), true)).Cardinality(), true), func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(237), _dafny.IntOfInt64(872))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(237)).Cmp(_4_v0) <= 0) && ((_4_v0).Cmp(_dafny.IntOfInt64(872)) < 0) {
				_coll2.Add((_4_v0).Minus(_dafny.IntOfInt64(-833)), true)
			}
		}
		return _coll2.ToMap()
	}())).Union(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-32), false))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (Companion_D8_.Create_DC22_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('j')))).Dtor_cf31()
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.CodePoint('m'))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('b')
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("buhcp"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("lqxqtv")))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-81))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_5_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rqow")).Cardinality())
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(496))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_6_i1 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(617), _dafny.IntOfInt64(128)))).Cardinality()
	})))).Cardinality()), _dafny.IntOfInt64(-39))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var _7_v0 _dafny.Sequence
	_ = _7_v0
	_7_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dfhmvgywh")
	var _hi0 _dafny.Int = ((_dafny.MultiSetOf(_7_v0, _dafny.UnicodeSeqOfUtf8Bytes("yet"), _7_v0)).Cardinality()).Times(_dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality()))
	_ = _hi0
	for _8_i0 := (_dafny.Zero).Minus(p2); _8_i0.Cmp(_hi0) < 0; _8_i0 = _8_i0.Plus(_dafny.One) {
		if false {
			var _9_v1 *C0
			_ = _9_v1
			var _nw0 *C0 = New_C0_()
			_ = _nw0
			_nw0.Ctor__(p0)
			_9_v1 = _nw0
			var _10_v2 _dafny.Sequence
			_ = _10_v2
			_10_v2 = _dafny.SeqOf(p1)
			var _11_v3 _dafny.MultiSet
			_ = _11_v3
			_11_v3 = _dafny.MultiSetOf((Companion_Default___.Fm11((_9_v1).F24(), (_dafny.Zero).Minus(_8_i0), globalState)).Cardinality(), _8_i0, _8_i0)
			var _12_v4 _dafny.Map
			_ = _12_v4
			_12_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_i0, _11_v3)
			r1 = (_9_v1).Fm4((Companion_D1_.Create_DC4_(_10_v2, p0)).Dtor_cf7(), p0, (func() _dafny.MultiSet {
				if Companion_Default___.Fm0(p2, (_9_v1).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(337))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}((func(_13_i0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_14_i1 _dafny.Int) _dafny.CodePoint {
						return (Companion_D3_.Create_DC10_(_13_i0, _dafny.CodePoint('m'))).Dtor_cf17()
					}
				})(_8_i0))), globalState) {
					return _11_v3
				}
				return (func() _dafny.MultiSet {
					if (_12_v4).Contains(_8_i0) {
						return (_12_v4).Get(_8_i0).(_dafny.MultiSet)
					}
					return _11_v3
				})()
			})(), Companion_Default___.Fm0(p2, p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(532))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_15_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			})), globalState), globalState)
			var _16_v5 _dafny.Sequence
			_ = _16_v5
			_16_v5 = _dafny.SeqOf(_9_v1)
			var _17_v6 _dafny.Sequence
			_ = _17_v6
			_17_v6 = _dafny.SeqOf(_16_v5)
			var _18_v7 _dafny.Set
			_ = _18_v7
			_18_v7 = _dafny.SetOf(_dafny.SeqOf(_9_v1), (_17_v6).Select((Companion_Default___.SafeIndex(_8_i0, _dafny.IntOfUint32((_17_v6).Cardinality()))).Uint32()).(_dafny.Sequence), _16_v5, _16_v5, _16_v5)
			var _19_v8 _dafny.Array
			_ = _19_v8
			var _nwElement0_0 _dafny.Set = _18_v7
			_ = _nwElement0_0
			var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(17))
			_ = _nw1
			(_nw1).ArraySet1(_nwElement0_0, 0)
			(_nw1).ArraySet1(_18_v7, 1)
			(_nw1).ArraySet1(_18_v7, 2)
			(_nw1).ArraySet1((_dafny.SetOf(_16_v5)).Difference(_18_v7), 3)
			(_nw1).ArraySet1(_18_v7, 4)
			(_nw1).ArraySet1(_18_v7, 5)
			(_nw1).ArraySet1((_18_v7).Difference(_18_v7), 6)
			(_nw1).ArraySet1(_dafny.SetOf(_16_v5), 7)
			(_nw1).ArraySet1(_dafny.SetOf(_16_v5), 8)
			(_nw1).ArraySet1((_18_v7).Union(_18_v7), 9)
			(_nw1).ArraySet1(_18_v7, 10)
			(_nw1).ArraySet1((_18_v7).Union(_18_v7), 11)
			(_nw1).ArraySet1(_18_v7, 12)
			(_nw1).ArraySet1(_18_v7, 13)
			(_nw1).ArraySet1((_dafny.SetOf(_16_v5)).Intersection(_dafny.SetOf(_dafny.SeqOf(_9_v1))), 14)
			(_nw1).ArraySet1((_18_v7).Union(_18_v7), 15)
			(_nw1).ArraySet1((_18_v7).Intersection(_18_v7), 16)
			_19_v8 = _nw1
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_19_v8), 0))
			_ = _index0
			(_19_v8).ArraySet1((_18_v7).Union(_dafny.SetOf(_16_v5)), (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_19_v8), 0))
			_ = _index1
			(_19_v8).ArraySet1(_18_v7, (_index1).Int())
			(globalState).F23 = _dafny.IntOfInt64(-419)
			_9_v1 = _9_v1
		} else {
			var _20_v9 _dafny.Sequence
			_ = _20_v9
			_20_v9 = _dafny.SeqOf(p2, p2)
			r1 = (Companion_Default___.SafeModuloInt((_20_v9).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_20_v9).Cardinality()))).Uint32()).(_dafny.Int), _8_i0)).Times(_dafny.IntOfInt64(-624))
			var _21_v10 _dafny.Map
			_ = _21_v10
			_21_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !(p0) || (false))
			_21_v10 = _21_v10
			(globalState).F2 = p0
			var _22_v11 _dafny.CodePoint
			_ = _22_v11
			_22_v11 = _dafny.CodePoint('q')
			var _23_v12 D0
			_ = _23_v12
			_23_v12 = Companion_D0_.Create_DC0_(_22_v11)
			_23_v12 = (func() D0 {
				if p1 {
					return _23_v12
				}
				return Companion_Default___.Fm12(p1, true, globalState)
			})()
			var _24_v13 _dafny.Sequence
			_ = _24_v13
			_24_v13 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(76)))
			var _25_v15 _dafny.Map
			_ = _25_v15
			_25_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _8_i0)
			var _26_v16 _dafny.MultiSet
			_ = _26_v16
			_26_v16 = _dafny.MultiSetOf(Companion_Default___.Fm13((_25_v15).Cardinality(), globalState), _22_v11)
			var _27_v17 _dafny.Map
			_ = _27_v17
			_27_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate((_26_v16).Elements()); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _28_v14 _dafny.CodePoint
					_28_v14 = interface{}(_compr_3).(_dafny.CodePoint)
					if (_26_v16).Contains(_28_v14) {
						_coll3.Add(_28_v14, _8_i0)
					}
				}
				return _coll3.ToMap()
			}(), p1)
			var _rhs0 bool = true
			_ = _rhs0
			var _rhs1 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_24_v13).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_24_v13).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_8_i0, _dafny.IntOfInt64(701), (_27_v17).Cardinality()), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf(_8_i0, _dafny.IntOfInt64(701), (_27_v17).Cardinality())).Cardinality()))).Uint32(), p2))
			_ = _rhs1
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			_lhs0.F18 = _rhs0
			_20_v9 = _rhs1
		}
		var _29_v18 *C0
		_ = _29_v18
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__(p0)
		_29_v18 = _nw2
		if p0 {
			(globalState).F18 = p1
			_7_v0 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(globalState), _7_v0)
			var _nw3 *C0 = New_C0_()
			_ = _nw3
			_nw3.Ctor__((p1) || (p0))
			_29_v18 = _nw3
			var _30_v19 _dafny.Array
			_ = _30_v19
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw4
			_30_v19 = _nw4
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_30_v19), 0))
			_ = _index2
			(_30_v19).ArraySet1((_dafny.Zero).Minus((_8_i0).Plus(_8_i0)), (_index2).Int())
			var _31_v20 _dafny.Map
			_ = _31_v20
			_31_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_8_i0), p1)
			var _32_v21 _dafny.Sequence
			_ = _32_v21
			_32_v21 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_7_v0).Cardinality()))).Cardinality()), (_31_v20).Cardinality())
			var _33_v22 D4
			_ = _33_v22
			_33_v22 = Companion_D4_.Create_DC13_(_32_v21)
			var _34_v23 _dafny.Array
			_ = _34_v23
			var _nwElement0_1 _dafny.Sequence = _32_v21
			_ = _nwElement0_1
			var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(5))
			_ = _nw5
			(_nw5).ArraySet1(_nwElement0_1, 0)
			(_nw5).ArraySet1(_32_v21, 1)
			(_nw5).ArraySet1((_33_v22).Dtor_cf20(), 2)
			(_nw5).ArraySet1(_32_v21, 3)
			(_nw5).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(799))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_35_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_36_i3 _dafny.Int) _dafny.Int {
					return _35_p2
				}
			})(p2))), 4)
			_34_v23 = _nw5
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_34_v23), 0))
			_ = _index3
			(_34_v23).ArraySet1(_dafny.SeqOf(_8_i0), (_index3).Int())
			var _37_v24 _dafny.Sequence
			_ = _37_v24
			_37_v24 = _dafny.SeqOf((_29_v18).F24())
			var _38_v25 _dafny.Set
			_ = _38_v25
			_38_v25 = _dafny.SetOf(_37_v24)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_30_v19), 0))
			_ = _index4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_34_v23), 0))
			_ = _index5
			var _rhs2 _dafny.Int = ((_38_v25).Intersection(_38_v25)).Cardinality()
			_ = _rhs2
			var _rhs3 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_32_v21, _32_v21)
			_ = _rhs3
			var _lhs1 _dafny.Array = _30_v19
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_30_v19), 0))
			_ = _lhs2
			var _lhs3 _dafny.Array = _34_v23
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_34_v23), 0))
			_ = _lhs4
			(_lhs1).ArraySet1(_rhs2, (_lhs2).Int())
			(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_30_v19), 0))
			_ = _index6
			(_30_v19).ArraySet1((_30_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_30_v19), 0))).Int()).(_dafny.Int), (_index6).Int())
		} else {
			var _39_v26 _dafny.MultiSet
			_ = _39_v26
			_39_v26 = _dafny.MultiSetOf(p1, p1)
			r0 = (p2).Times(((_39_v26).Cardinality()).Times(_8_i0))
			var _40_v27 _dafny.Array
			_ = _40_v27
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw6
			_40_v27 = _nw6
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_40_v27), 0))
			_ = _index7
			(_40_v27).ArraySet1(p2, (_index7).Int())
			var _41_v28 _dafny.Map
			_ = _41_v28
			_41_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_8_i0).Cmp(_8_i0) != 0, _7_v0)
			var _42_v29 _dafny.Map
			_ = _42_v29
			_42_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(648))
			var _43_v30 _dafny.Sequence
			_ = _43_v30
			_43_v30 = _dafny.SeqOf((_42_v29).Cardinality(), p2)
			var _44_v31 _dafny.CodePoint
			_ = _44_v31
			_44_v31 = _dafny.CodePoint('t')
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_40_v27), 0))
			_ = _index8
			var _rhs4 _dafny.Int = (_29_v18).Fm4((_29_v18).F24(), !(_dafny.Companion_Sequence_.IsPrefixOf(_43_v30, _43_v30)), _dafny.MultiSetOf(_8_i0, p2, _8_i0, _8_i0), (_29_v18).F24(), globalState)
			_ = _rhs4
			var _rhs5 bool = p0
			_ = _rhs5
			var _rhs6 bool = p0
			_ = _rhs6
			var _rhs7 _dafny.Map = Companion_Default___.Fm14(_44_v31, globalState)
			_ = _rhs7
			var _lhs5 _dafny.Array = _40_v27
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_40_v27), 0))
			_ = _lhs6
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			(_lhs5).ArraySet1(_rhs4, (_lhs6).Int())
			_lhs7.F2 = _rhs5
			_lhs8.F2 = _rhs6
			_41_v28 = _rhs7
			var _45_v32 _dafny.Map
			_ = _45_v32
			_45_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-384))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_46_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_47_i4 _dafny.Int) _dafny.Int {
					return _46_i0
				}
			})(_8_i0)))).Cardinality()), p2)
			var _48_v33 _dafny.Map
			_ = _48_v33
			_48_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_29_v18).F24(), p0), !(true))
			var _49_v34 _dafny.Map
			_ = _49_v34
			_49_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_45_v32).Cardinality(), !((func() bool {
				if (_48_v33).Contains(_dafny.SeqOf(p1)) {
					return (_48_v33).Get(_dafny.SeqOf(p1)).(bool)
				}
				return (_29_v18).F24()
			})()))
			var _50_v35 _dafny.MultiSet
			_ = _50_v35
			_50_v35 = _dafny.MultiSetOf((_49_v34).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)))
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_40_v27), 0))
			_ = _index9
			(_40_v27).ArraySet1((func() _dafny.Int {
				if (_50_v35).Contains(Companion_Default___.Fm9((_dafny.Zero).Minus((_39_v26).Cardinality()), (_29_v18).F24(), p2, p1, globalState)) {
					return (_50_v35).Multiplicity(Companion_Default___.Fm9((_dafny.Zero).Minus((_39_v26).Cardinality()), (_29_v18).F24(), p2, p1, globalState))
				}
				return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_8_i0), _8_i0)
			})(), (_index9).Int())
			var _51_v36 D3
			_ = _51_v36
			_51_v36 = Companion_D3_.Create_DC10_(p2, _44_v31)
			var _52_v37 _dafny.Map
			_ = _52_v37
			_52_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_29_v18, _51_v36)
			var _53_v38 _dafny.Array
			_ = _53_v38
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw7
			_53_v38 = _nw7
			var _54_v39 _dafny.Array
			_ = _54_v39
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw8
			_54_v39 = _nw8
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_53_v38), 0))
			_ = _index10
			(_53_v38).ArraySet1(_54_v39, (_index10).Int())
			var _55_v40 D5
			_ = _55_v40
			_55_v40 = Companion_D5_.Create_DC15_(_54_v39)
			var _56_v41 _dafny.Sequence
			_ = _56_v41
			_56_v41 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(422))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_57_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_58_i5 _dafny.Int) _dafny.CodePoint {
					return _57_v31
				}
			})(_44_v31))))
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_53_v38), 0))
			_ = _index11
			var _rhs8 _dafny.Map = ((_52_v37).Merge(_52_v37)).Update(_29_v18, _51_v36)
			_ = _rhs8
			var _rhs9 _dafny.Array = (_55_v40).Dtor_cf22()
			_ = _rhs9
			var _rhs10 _dafny.Sequence = (_56_v41).Select((Companion_Default___.SafeIndex((_40_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_40_v27), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_56_v41).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs10
			var _lhs9 _dafny.Array = _53_v38
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_53_v38), 0))
			_ = _lhs10
			_52_v37 = _rhs8
			(_lhs9).ArraySet1(_rhs9, (_lhs10).Int())
			_7_v0 = _rhs10
			_29_v18 = _29_v18
		}
		var _59_v42 D3
		_ = _59_v42
		_59_v42 = Companion_D3_.Create_DC11_((_29_v18).F24())
		var _60_v43 *C0
		_ = _60_v43
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__((_dafny.IntOfInt64(422)).Cmp((_dafny.MultiSetOf(_dafny.SetOf((_59_v42).Dtor_cf18(), p0))).Cardinality()) < 0)
		_60_v43 = _nw9
	}
	(globalState).F18 = Companion_Default___.Fm0(p2, p1, _dafny.Companion_Sequence_.Concatenate(_7_v0, _7_v0), globalState)
	var _hi1 _dafny.Int = p2
	_ = _hi1
	for _61_i6 := p2; _61_i6.Cmp(_hi1) < 0; _61_i6 = _61_i6.Plus(_dafny.One) {
		var _62_v44 _dafny.CodePoint
		_ = _62_v44
		_62_v44 = _dafny.CodePoint('h')
		var _63_v45 *C0
		_ = _63_v45
		var _nw10 *C0 = New_C0_()
		_ = _nw10
		_nw10.Ctor__(p1)
		_63_v45 = _nw10
		var _64_v46 _dafny.Map
		_ = _64_v46
		_64_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v45, (_63_v45).F24())
		var _65_v47 _dafny.Map
		_ = _65_v47
		_65_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_i6, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_64_v46).Cardinality(), (_63_v45).F24())).Update(_61_i6, p0))
		var _66_v48 D6
		_ = _66_v48
		_66_v48 = Companion_D6_.Create_DC18_(_65_v47)
		var _67_v49 _dafny.Map
		_ = _67_v49
		_67_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_64_v46).Merge(_64_v46), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-676))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_68_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_69_i7 _dafny.Int) _dafny.CodePoint {
				return _68_v44
			}
		})(_62_v44))))
		var _rhs11 _dafny.CodePoint = _62_v44
		_ = _rhs11
		var _rhs12 _dafny.Map = (_66_v48).Dtor_cf26()
		_ = _rhs12
		var _rhs13 _dafny.Sequence = (func() _dafny.Sequence {
			if (_67_v49).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v45, (_63_v45).F24())) {
				return (_67_v49).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v45, (_63_v45).F24())).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Concatenate(_7_v0, _7_v0)
		})()
		_ = _rhs13
		var _rhs14 *C0 = _63_v45
		_ = _rhs14
		var _rhs15 bool = (_61_i6).Cmp(_dafny.IntOfInt64(153)) != 0
		_ = _rhs15
		var _lhs11 *GlobalState = globalState
		_ = _lhs11
		_62_v44 = _rhs11
		_65_v47 = _rhs12
		_7_v0 = _rhs13
		_63_v45 = _rhs14
		_lhs11.F18 = _rhs15
		r1 = _61_i6
		var _70_v50 _dafny.Sequence
		_ = _70_v50
		_70_v50 = _dafny.SeqOf(p1)
		var _71_v51 D1
		_ = _71_v51
		_71_v51 = Companion_D1_.Create_DC4_(_70_v50, !(p0))
		var _72_v52 _dafny.Sequence
		_ = _72_v52
		_72_v52 = _dafny.SeqOf(_7_v0)
		var _73_v54 _dafny.Sequence
		_ = _73_v54
		_73_v54 = _dafny.SeqOf(_dafny.IntOfUint32((_70_v50).Cardinality()), p2)
		var _74_v55 _dafny.Map
		_ = _74_v55
		_74_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_63_v45).Fm4(p0, (_63_v45).F24(), _dafny.MultiSetFromSeq(_73_v54), p1, globalState), _61_i6)
		var _rhs16 bool = p0
		_ = _rhs16
		var _rhs17 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_72_v52).Cardinality()), _61_i6), ((func() _dafny.Map {
			if p0 {
				return func() _dafny.Map {
					var _coll4 = _dafny.NewMapBuilder()
					_ = _coll4
					for _iter4 := _dafny.Iterate((Companion_Default___.Fm15(_61_i6, p0, p1, globalState)).Keys().Elements()); ; {
						_compr_4, _ok4 := _iter4()
						if !_ok4 {
							break
						}
						var _75_v53 _dafny.Int
						_75_v53 = interface{}(_compr_4).(_dafny.Int)
						if (Companion_Default___.Fm15(_61_i6, p0, p1, globalState)).Contains(_75_v53) {
							_coll4.Add((_75_v53).Times(_61_i6), p2)
						}
					}
					return _coll4.ToMap()
				}()
			}
			return _74_v55
		})()).Cardinality())
		_ = _rhs17
		var _rhs18 D1 = _71_v51
		_ = _rhs18
		var _lhs12 *GlobalState = globalState
		_ = _lhs12
		var _lhs13 *GlobalState = globalState
		_ = _lhs13
		_lhs12.F18 = _rhs16
		_lhs13.F23 = _rhs17
		_71_v51 = _rhs18
		r1 = (_dafny.Zero).Minus((_61_i6).Minus(p2))
	}
	var _76_i8 _dafny.Int
	_ = _76_i8
	_76_i8 = _dafny.Zero
	{
		for Companion_Default___.Fm0((p2).Plus(p2), p1, _7_v0, globalState) {
			{
				if (_76_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_76_i8 = (_76_i8).Plus(_dafny.One)
				(globalState).F18 = p1
				var _77_v56 D6
				_ = _77_v56
				_77_v56 = Companion_D6_.Create_DC20_()
				_77_v56 = _77_v56
				var _78_v57 _dafny.Array
				_ = _78_v57
				var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw11
				_78_v57 = _nw11
				var _79_v58 _dafny.Array
				_ = _79_v58
				var _nw12 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw12
				_79_v58 = _nw12
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_78_v57), 0))
				_ = _index12
				(_78_v57).ArraySet1(_79_v58, (_index12).Int())
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_78_v57), 0))
				_ = _index13
				var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw13
				(_78_v57).ArraySet1(_nw13, (_index13).Int())
				(globalState).F2 = p0
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _80_v59 _dafny.Array
	_ = _80_v59
	var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
	_ = _nw14
	_80_v59 = _nw14
	var _81_v60 _dafny.CodePoint
	_ = _81_v60
	_81_v60 = _dafny.CodePoint('v')
	var _82_v61 _dafny.Array
	_ = _82_v61
	var _nwElement0_2 _dafny.Sequence = _7_v0
	_ = _nwElement0_2
	var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(9))
	_ = _nw15
	(_nw15).ArraySet1(_nwElement0_2, 0)
	(_nw15).ArraySet1(_7_v0, 1)
	(_nw15).ArraySet1(_7_v0, 2)
	(_nw15).ArraySet1(_7_v0, 3)
	(_nw15).ArraySet1(_7_v0, 4)
	(_nw15).ArraySet1(_7_v0, 5)
	(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("c"), 6)
	(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("krdlpcyb"), 7)
	(_nw15).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("sj"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sj")).Cardinality()))).Uint32(), _81_v60), 8)
	_82_v61 = _nw15
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_80_v59), 0))
	_ = _index14
	(_80_v59).ArraySet1(_82_v61, (_index14).Int())
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_80_v59), 0))
	_ = _index15
	(_80_v59).ArraySet1(_82_v61, (_index15).Int())
	(globalState).F2 = p1
	var _83_v62 _dafny.Array
	_ = _83_v62
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(9)
	_ = _len0_0
	var _nw16 _dafny.Array
	_ = _nw16
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw16 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = (func(_84_p0 bool) func(_dafny.Int) bool {
			return func(_85_i9 _dafny.Int) bool {
				return _84_p0
			}
		})(p0)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw16 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw16).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw16).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_83_v62 = _nw16
	var _86_v63 _dafny.MultiSet
	_ = _86_v63
	_86_v63 = _dafny.MultiSetOf(_83_v62, _83_v62)
	r0 = ((_86_v63).Difference((_86_v63).Update(_83_v62, Companion_Default___.Abs(_dafny.IntOfInt64(177))))).Cardinality()
	var _87_v64 _dafny.Set
	_ = _87_v64
	_87_v64 = _dafny.SetOf(_81_v60, _81_v60)
	r1 = (p2).Plus((_dafny.Zero).Minus(((_87_v64).Difference(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_87_v64).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _88_v65 _dafny.CodePoint
			_88_v65 = interface{}(_compr_5).(_dafny.CodePoint)
			if (_87_v64).Contains(_88_v65) {
				_coll5.Add(_88_v65)
			}
		}
		return _coll5.ToSet()
	}())).Cardinality()))
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _89_v0 _dafny.Sequence
	_ = _89_v0
	_89_v0 = _dafny.UnicodeSeqOfUtf8Bytes("spiw")
	var _90_v1 _dafny.MultiSet
	_ = _90_v1
	_90_v1 = _dafny.MultiSetOf(_89_v0, _89_v0, _89_v0, _89_v0, _89_v0)
	var _91_v2 bool
	_ = _91_v2
	_91_v2 = true
	var _92_v3 _dafny.Int
	_ = _92_v3
	_92_v3 = _dafny.IntOfInt64(56)
	var _93_v4 _dafny.Map
	_ = _93_v4
	_93_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v2, _92_v3)
	var _94_v5 _dafny.Map
	_ = _94_v5
	_94_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v4, _92_v3)
	var _95_v6 _dafny.Array
	_ = _95_v6
	var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
	_ = _nw17
	_95_v6 = _nw17
	var _96_v7 _dafny.Map
	_ = _96_v7
	_96_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v2, _95_v6)
	var _97_v8 _dafny.Array
	_ = _97_v8
	var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
	_ = _nw18
	_97_v8 = _nw18
	var _98_v9 _dafny.CodePoint
	_ = _98_v9
	_98_v9 = _dafny.CodePoint('n')
	var _99_v10 _dafny.Map
	_ = _99_v10
	_99_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v2, _98_v9)
	var _100_v11 _dafny.Map
	_ = _100_v11
	_100_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v3, _91_v2)
	var _101_v12 _dafny.Array
	_ = _101_v12
	var _nwElement0_3 _dafny.CodePoint = _98_v9
	_ = _nwElement0_3
	var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(16))
	_ = _nw19
	(_nw19).ArraySet1CodePoint(_nwElement0_3, 0)
	(_nw19).ArraySet1CodePoint(_98_v9, 1)
	(_nw19).ArraySet1CodePoint(_98_v9, 2)
	(_nw19).ArraySet1CodePoint(_98_v9, 3)
	(_nw19).ArraySet1CodePoint(_dafny.CodePoint('w'), 4)
	(_nw19).ArraySet1CodePoint(_98_v9, 5)
	(_nw19).ArraySet1CodePoint(_98_v9, 6)
	(_nw19).ArraySet1CodePoint(_98_v9, 7)
	(_nw19).ArraySet1CodePoint(_98_v9, 8)
	(_nw19).ArraySet1CodePoint(_98_v9, 9)
	(_nw19).ArraySet1CodePoint(_98_v9, 10)
	(_nw19).ArraySet1CodePoint(_98_v9, 11)
	(_nw19).ArraySet1CodePoint((func() _dafny.CodePoint {
		if (_99_v10).Contains((func() bool {
			if (_100_v11).Contains(_92_v3) {
				return (_100_v11).Get(_92_v3).(bool)
			}
			return _91_v2
		})()) {
			return (_99_v10).Get((func() bool {
				if (_100_v11).Contains(_92_v3) {
					return (_100_v11).Get(_92_v3).(bool)
				}
				return _91_v2
			})()).(_dafny.CodePoint)
		}
		return _98_v9
	})(), 12)
	(_nw19).ArraySet1CodePoint(_98_v9, 13)
	(_nw19).ArraySet1CodePoint(_98_v9, 14)
	(_nw19).ArraySet1CodePoint(_dafny.CodePoint('g'), 15)
	_101_v12 = _nw19
	var _102_v13 _dafny.Sequence
	_ = _102_v13
	_102_v13 = _dafny.SeqOf(_92_v3, _92_v3)
	var _103_globalState *GlobalState
	_ = _103_globalState
	var _nw20 *GlobalState = New_GlobalState_()
	_ = _nw20
	_nw20.Ctor__(_dafny.IntOfInt64(580), _90_v1, false, _94_v5, _dafny.IntOfInt64(548), true, true, _dafny.IntOfInt64(376), _dafny.IntOfInt64(87), true, (func() _dafny.Array {
		if (_96_v7).Contains(false) {
			return (_96_v7).Get(false).(_dafny.Array)
		}
		return _95_v6
	})(), _dafny.CodePoint('w'), _dafny.IntOfInt64(-257), _97_v8, _dafny.Companion_Sequence_.Concatenate(_89_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-26))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_104_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))), _101_v12, _dafny.IntOfInt64(-272), _101_v12, false, true, _dafny.UnicodeSeqOfUtf8Bytes("dblg"), _102_v13, _102_v13, _dafny.IntOfInt64(-128))
	_103_globalState = _nw20
	var _105_v14 _dafny.Int
	_ = _105_v14
	var _106_v15 _dafny.Int
	_ = _106_v15
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	_out0, _out1 = Companion_Default___.M0(_91_v2, _91_v2, _92_v3, _103_globalState)
	_105_v14 = _out0
	_106_v15 = _out1
	var _107_v16 _dafny.MultiSet
	_ = _107_v16
	_107_v16 = _dafny.MultiSetOf(_91_v2, _91_v2)
	if (_107_v16).IsProperSubsetOf(_107_v16) {
		var _108_v17 _dafny.Int
		_ = _108_v17
		var _109_v18 _dafny.Int
		_ = _109_v18
		var _out2 _dafny.Int
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out2, _out3 = Companion_Default___.M0(_91_v2, (_105_v14).Cmp(_92_v3) <= 0, (_106_v15).Minus(_92_v3), _103_globalState)
		_108_v17 = _out2
		_109_v18 = _out3
		if _91_v2 {
			var _110_v19 _dafny.Array
			_ = _110_v19
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_1
			var _nw21 _dafny.Array
			_ = _nw21
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw21 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = func(_111_i1 _dafny.Int) bool {
					return true
				}
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw21 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw21).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw21).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_110_v19 = _nw21
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))
			_ = _index16
			(_110_v19).ArraySet1(_91_v2, (_index16).Int())
			var _112_v20 _dafny.Sequence
			_ = _112_v20
			_112_v20 = _dafny.SeqOf(_91_v2, _91_v2)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))
			_ = _index17
			(_110_v19).ArraySet1((_112_v20).Select((Companion_Default___.SafeIndex(_109_v18, _dafny.IntOfUint32((_112_v20).Cardinality()))).Uint32()).(bool), (_index17).Int())
			var _113_v21 _dafny.Map
			_ = _113_v21
			_113_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v2, (_110_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))).Int()).(bool))
			var _114_v22 _dafny.Array
			_ = _114_v22
			var _nwElement0_4 _dafny.Map = _113_v21
			_ = _nwElement0_4
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_4, 0)
			(_nw22).ArraySet1((_113_v21).Update(_91_v2, (_110_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))).Int()).(bool)), 1)
			(_nw22).ArraySet1(_113_v21, 2)
			(_nw22).ArraySet1(_113_v21, 3)
			(_nw22).ArraySet1((_113_v21).Update((_110_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))).Int()).(bool), _91_v2), 4)
			(_nw22).ArraySet1(_113_v21, 5)
			(_nw22).ArraySet1(_113_v21, 6)
			(_nw22).ArraySet1(_113_v21, 7)
			(_nw22).ArraySet1((_113_v21).Merge(_113_v21), 8)
			(_nw22).ArraySet1(_113_v21, 9)
			(_nw22).ArraySet1((_113_v21).Merge(_113_v21), 10)
			(_nw22).ArraySet1(_113_v21, 11)
			_114_v22 = _nw22
			_114_v22 = (func() _dafny.Array {
				if (_110_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))).Int()).(bool) {
					return _114_v22
				}
				return (func() _dafny.Array {
					if _91_v2 {
						return _114_v22
					}
					return _114_v22
				})()
			})()
			_100_v11 = (_100_v11).Update(_105_v14, _91_v2)
			var _115_v23 D0
			_ = _115_v23
			_115_v23 = Companion_D0_.Create_DC0_(_98_v9)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))
			_ = _index18
			(_110_v19).ArraySet1(Companion_Default___.Fm0((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_92_v3, _92_v3)), !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(_103_globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_89_v0).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm1(_103_globalState)).Cardinality()))).Uint32(), (_115_v23).Dtor_cf0()), _dafny.UnicodeSeqOfUtf8Bytes("d"))), _89_v0, _103_globalState), (_index18).Int())
			var _116_v24 _dafny.Int
			_ = _116_v24
			var _117_v25 _dafny.Int
			_ = _117_v25
			var _out4 _dafny.Int
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out4, _out5 = Companion_Default___.M0((_110_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_110_v19), 0))).Int()).(bool), _91_v2, _109_v18, _103_globalState)
			_116_v24 = _out4
			_117_v25 = _out5
		} else {
			_105_v14 = _92_v3
			var _118_v26 D1
			_ = _118_v26
			_118_v26 = Companion_D1_.Create_DC2_(_91_v2)
			(_103_globalState).F18 = Companion_Default___.Fm0(_108_v17, (_118_v26).Dtor_cf2(), _89_v0, _103_globalState)
			(_103_globalState).F23 = _109_v18
			var _119_v27 _dafny.Sequence
			_ = _119_v27
			_119_v27 = _dafny.SeqOf(_91_v2, _91_v2)
			var _120_v28 _dafny.Map
			_ = _120_v28
			_120_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v2, !(false))
			var _121_v29 _dafny.Array
			_ = _121_v29
			var _nwElement0_5 bool = !(_91_v2) || (_91_v2)
			_ = _nwElement0_5
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(17))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_5, 0)
			(_nw23).ArraySet1(false, 1)
			(_nw23).ArraySet1(_91_v2, 2)
			(_nw23).ArraySet1(_91_v2, 3)
			(_nw23).ArraySet1(_91_v2, 4)
			(_nw23).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_89_v0, _89_v0)), 5)
			(_nw23).ArraySet1((_105_v14).Cmp(_dafny.IntOfInt64(-375)) == 0, 6)
			(_nw23).ArraySet1(_91_v2, 7)
			(_nw23).ArraySet1((_105_v14).Cmp(_105_v14) != 0, 8)
			(_nw23).ArraySet1((_91_v2) && (_91_v2), 9)
			(_nw23).ArraySet1(true, 10)
			(_nw23).ArraySet1(Companion_Default___.Fm0(_105_v14, _91_v2, _dafny.UnicodeSeqOfUtf8Bytes("n"), _103_globalState), 11)
			(_nw23).ArraySet1(_91_v2, 12)
			(_nw23).ArraySet1(_91_v2, 13)
			(_nw23).ArraySet1(_91_v2, 14)
			(_nw23).ArraySet1(true, 15)
			(_nw23).ArraySet1((_119_v27).Select((Companion_Default___.SafeIndex((_120_v28).Cardinality(), _dafny.IntOfUint32((_119_v27).Cardinality()))).Uint32()).(bool), 16)
			_121_v29 = _nw23
			var _rhs19 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(952))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_122_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_123_i2 _dafny.Int) _dafny.CodePoint {
					return _122_v9
				}
			})(_98_v9)))
			_ = _rhs19
			var _rhs20 _dafny.Sequence = _89_v0
			_ = _rhs20
			var _rhs21 _dafny.Array = _121_v29
			_ = _rhs21
			var _rhs22 bool = ((_dafny.Zero).Minus(_92_v3)).Cmp(_109_v18) != 0
			_ = _rhs22
			var _rhs23 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_89_v0, _dafny.UnicodeSeqOfUtf8Bytes("nboyosu"))).Cardinality())).Plus((_dafny.Zero).Minus(_109_v18))
			_ = _rhs23
			var _lhs14 *GlobalState = _103_globalState
			_ = _lhs14
			var _lhs15 *GlobalState = _103_globalState
			_ = _lhs15
			_89_v0 = _rhs19
			_lhs14.F20 = _rhs20
			_121_v29 = _rhs21
			_lhs15.F2 = _rhs22
			_105_v14 = _rhs23
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_121_v29), 0))
			_ = _index19
			(_121_v29).ArraySet1(!(_91_v2), (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_121_v29), 0))
			_ = _index20
			(_121_v29).ArraySet1(_91_v2, (_index20).Int())
		}
		var _124_v30 _dafny.Set
		_ = _124_v30
		_124_v30 = _dafny.SetOf(_108_v17)
		var _125_v31 _dafny.MultiSet
		_ = _125_v31
		_125_v31 = _dafny.MultiSetOf(_109_v18, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nfrvenm")).Cardinality()), _92_v3)
		var _126_v32 _dafny.Int
		_ = _126_v32
		var _127_v33 _dafny.Int
		_ = _127_v33
		var _out6 _dafny.Int
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		_out6, _out7 = Companion_Default___.M0(_91_v2, (Companion_Default___.Fm2(_124_v30, _106_v15, _105_v14, _dafny.IntOfInt64(916), _103_globalState)).IsDisjointFrom(_125_v31), Companion_Default___.Fm3(_106_v15, _103_globalState), _103_globalState)
		_126_v32 = _out6
		_127_v33 = _out7
		_109_v18 = _106_v15
		if _91_v2 {
			var _128_v34 _dafny.Int
			_ = _128_v34
			var _129_v35 _dafny.Int
			_ = _129_v35
			var _out8 _dafny.Int
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			_out8, _out9 = Companion_Default___.M0(false, _91_v2, Companion_Default___.SafeModuloInt(_109_v18, _105_v14), _103_globalState)
			_128_v34 = _out8
			_129_v35 = _out9
			var _130_v36 D1
			_ = _130_v36
			_130_v36 = Companion_D1_.Create_DC2_(_91_v2)
			var _131_v37 _dafny.Array
			_ = _131_v37
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw24
			_131_v37 = _nw24
			var _132_v38 _dafny.Map
			_ = _132_v38
			_132_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v36, _131_v37)
			var _rhs24 bool = _91_v2
			_ = _rhs24
			var _rhs25 _dafny.Map = _132_v38
			_ = _rhs25
			var _lhs16 *GlobalState = _103_globalState
			_ = _lhs16
			_lhs16.F18 = _rhs24
			_132_v38 = _rhs25
			var _133_v39 *C0
			_ = _133_v39
			var _nw25 *C0 = New_C0_()
			_ = _nw25
			_nw25.Ctor__((func() bool {
				if _91_v2 {
					return !(_91_v2)
				}
				return !(_91_v2)
			})())
			_133_v39 = _nw25
			var _134_v40 *C0
			_ = _134_v40
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__((_133_v39).F24())
			_134_v40 = _nw26
			_134_v40 = _133_v39
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_131_v37), 0))
			_ = _index21
			(_131_v37).ArraySet1((_133_v39).F24(), (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_131_v37), 0))
			_ = _index22
			(_131_v37).ArraySet1(((_107_v16).Difference(_107_v16)).IsDisjointFrom((_107_v16).Difference(_107_v16)), (_index22).Int())
		} else {
			var _135_v41 _dafny.Set
			_ = _135_v41
			_135_v41 = _dafny.SetOf(_98_v9, _98_v9)
			var _136_v42 _dafny.Map
			_ = _136_v42
			_136_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v0, _135_v41)
			var _137_v43 _dafny.Sequence
			_ = _137_v43
			_137_v43 = _dafny.SeqOf((func() _dafny.Set {
				if (_136_v42).Contains(_89_v0) {
					return (_136_v42).Get(_89_v0).(_dafny.Set)
				}
				return _135_v41
			})())
			var _138_v44 _dafny.Int
			_ = _138_v44
			var _139_v45 _dafny.Int
			_ = _139_v45
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out10, _out11 = Companion_Default___.M0(Companion_Default___.Fm0(_92_v3, _91_v2, _89_v0, _103_globalState), _dafny.Companion_Sequence_.IsPrefixOf(_137_v43, _137_v43), _dafny.IntOfInt64(495), _103_globalState)
			_138_v44 = _out10
			_139_v45 = _out11
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_97_v8), 0))
			_ = _index23
			(_97_v8).ArraySet1((_105_v14).Times(_126_v32), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_97_v8), 0))
			_ = _index24
			(_97_v8).ArraySet1(_139_v45, (_index24).Int())
			var _140_v46 _dafny.Array
			_ = _140_v46
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_2
			var _nw27 _dafny.Array
			_ = _nw27
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw27 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) bool = (func(_141_v2 bool) func(_dafny.Int) bool {
					return func(_142_i3 _dafny.Int) bool {
						return _141_v2
					}
				})(_91_v2)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw27 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw27).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw27).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_140_v46 = _nw27
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_140_v46), 0))
			_ = _index25
			(_140_v46).ArraySet1(_91_v2, (_index25).Int())
			var _143_v47 D2
			_ = _143_v47
			_143_v47 = Companion_D2_.Create_DC6_(_89_v0)
			var _144_v48 _dafny.Array
			_ = _144_v48
			var _nwElement0_6 _dafny.Sequence = _89_v0
			_ = _nwElement0_6
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(18))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_6, 0)
			(_nw28).ArraySet1(_89_v0, 1)
			(_nw28).ArraySet1(_89_v0, 2)
			(_nw28).ArraySet1(_89_v0, 3)
			(_nw28).ArraySet1(_89_v0, 4)
			(_nw28).ArraySet1(_89_v0, 5)
			(_nw28).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-932))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_145_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_146_i4 _dafny.Int) _dafny.CodePoint {
					return _145_v9
				}
			})(_98_v9))), 6)
			(_nw28).ArraySet1(_89_v0, 7)
			(_nw28).ArraySet1(_89_v0, 8)
			(_nw28).ArraySet1((_143_v47).Dtor_cf9(), 9)
			(_nw28).ArraySet1(_89_v0, 10)
			(_nw28).ArraySet1(_89_v0, 11)
			(_nw28).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-501))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_147_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})), 12)
			(_nw28).ArraySet1(_89_v0, 13)
			(_nw28).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("awa"), 14)
			(_nw28).ArraySet1(_89_v0, 15)
			(_nw28).ArraySet1(_89_v0, 16)
			(_nw28).ArraySet1(_89_v0, 17)
			_144_v48 = _nw28
			var _148_v49 D1
			_ = _148_v49
			_148_v49 = Companion_D1_.Create_DC3_((Companion_D1_.Create_DC3_(_144_v48, _dafny.IntOfInt64(779), _91_v2)).Dtor_cf3(), _105_v14, _91_v2)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_140_v46), 0))
			_ = _index26
			(_140_v46).ArraySet1(!((_148_v49).Dtor_cf5()), (_index26).Int())
			var _149_v50 *C0
			_ = _149_v50
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__(Companion_Default___.Fm0((_124_v30).Cardinality(), (_140_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_140_v46), 0))).Int()).(bool), _dafny.Companion_Sequence_.Update(_89_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_127_v33), _dafny.IntOfUint32((_89_v0).Cardinality()))).Uint32(), _98_v9), _103_globalState))
			_149_v50 = _nw29
			_126_v32 = (_97_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_97_v8), 0))).Int()).(_dafny.Int)
		}
	} else {
		_91_v2 = _91_v2
		var _150_v51 _dafny.Array
		_ = _150_v51
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
		_ = _nw30
		_150_v51 = _nw30
		var _151_v52 _dafny.Sequence
		_ = _151_v52
		_151_v52 = _dafny.SeqOf(_91_v2)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))
		_ = _index27
		(_150_v51).ArraySet1(!_dafny.Companion_Sequence_.Equal(_151_v52, _151_v52), (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))
		_ = _index28
		(_150_v51).ArraySet1((_105_v14).Cmp(Companion_Default___.Fm3(_106_v15, _103_globalState)) <= 0, (_index28).Int())
		var _152_v53 _dafny.Array
		_ = _152_v53
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
		_ = _nw31
		_152_v53 = _nw31
		var _153_v54 _dafny.MultiSet
		_ = _153_v54
		_153_v54 = _dafny.MultiSetOf(_105_v14, _105_v14)
		var _154_v55 D1
		_ = _154_v55
		_154_v55 = Companion_D1_.Create_DC3_(_152_v53, (_153_v54).Cardinality(), Companion_Default___.Fm0(_92_v3, (_150_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))).Int()).(bool), _89_v0, _103_globalState))
		var _155_v56 _dafny.Int
		_ = _155_v56
		var _156_v57 _dafny.Int
		_ = _156_v57
		var _out12 _dafny.Int
		_ = _out12
		var _out13 _dafny.Int
		_ = _out13
		_out12, _out13 = Companion_Default___.M0(((_107_v16).Cardinality()).Cmp((_154_v55).Dtor_cf4()) >= 0, (_105_v14).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("efngv")).Cardinality())) == 0, (_102_v13).Select((Companion_Default___.SafeIndex(_106_v15, _dafny.IntOfUint32((_102_v13).Cardinality()))).Uint32()).(_dafny.Int), _103_globalState)
		_155_v56 = _out12
		_156_v57 = _out13
		var _157_v58 _dafny.Int
		_ = _157_v58
		var _158_v59 _dafny.Int
		_ = _158_v59
		var _out14 _dafny.Int
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		_out14, _out15 = Companion_Default___.M0(_91_v2, (_150_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))).Int()).(bool), _dafny.IntOfInt64(210), _103_globalState)
		_157_v58 = _out14
		_158_v59 = _out15
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))
		_ = _index29
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))
		_ = _index30
		var _rhs26 _dafny.Array = _150_v51
		_ = _rhs26
		var _rhs27 bool = ((_106_v15).Cmp(_158_v59) >= 0) && ((_dafny.MultiSetOf(_156_v57, _92_v3)).IsSubsetOf(_153_v54))
		_ = _rhs27
		var _rhs28 bool = (_150_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))).Int()).(bool)
		_ = _rhs28
		var _rhs29 bool = (_155_v56).Cmp((_dafny.Zero).Minus(_156_v57)) >= 0
		_ = _rhs29
		var _rhs30 _dafny.Int = (_dafny.Zero).Minus((_158_v59).Minus(_155_v56))
		_ = _rhs30
		var _lhs17 _dafny.Array = _150_v51
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))
		_ = _lhs18
		var _lhs19 _dafny.Array = _150_v51
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_150_v51), 0))
		_ = _lhs20
		var _lhs21 *GlobalState = _103_globalState
		_ = _lhs21
		var _lhs22 *GlobalState = _103_globalState
		_ = _lhs22
		_150_v51 = _rhs26
		(_lhs17).ArraySet1(_rhs27, (_lhs18).Int())
		(_lhs19).ArraySet1(_rhs28, (_lhs20).Int())
		_lhs21.F2 = _rhs29
		_lhs22.F23 = _rhs30
	}
	var _159_v60 _dafny.Array
	_ = _159_v60
	var _nw32 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
	_ = _nw32
	_159_v60 = _nw32
	var _160_v61 _dafny.Map
	_ = _160_v61
	_160_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_v60, _89_v0)
	(_103_globalState).F2 = Companion_Default___.Fm0(_105_v14, _91_v2, (func() _dafny.Sequence {
		if (_160_v61).Contains(_159_v60) {
			return (_160_v61).Get(_159_v60).(_dafny.Sequence)
		}
		return _89_v0
	})(), _103_globalState)
	var _161_v62 *C0
	_ = _161_v62
	var _nw33 *C0 = New_C0_()
	_ = _nw33
	_nw33.Ctor__(!(_91_v2))
	_161_v62 = _nw33
	var _162_v63 _dafny.Map
	_ = _162_v63
	_162_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v2, _161_v62)
	var _163_v64 _dafny.Map
	_ = _163_v64
	_163_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v63, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(950))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}((func(_164_v13 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
		return func(_165_i6 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_164_v13).Cardinality())
		}
	})(_102_v13))))
	_163_v64 = (_163_v64).Update(_162_v63, _102_v13)
	var _166_v65 _dafny.Map
	_ = _166_v65
	_166_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_161_v62).F24(), true)
	if (_166_v65).Contains((_161_v62).F24()) {
		var _167_v66 _dafny.Map
		_ = _167_v66
		_167_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_106_v15, _92_v3), _97_v8)
		_167_v66 = (_167_v66).Update(_dafny.IntOfInt64(271), _97_v8)
		var _168_v67 _dafny.Array
		_ = _168_v67
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_3
		var _nw34 _dafny.Array
		_ = _nw34
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw34 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = (func(_169_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_170_i7 _dafny.Int) _dafny.Sequence {
					return _169_v0
				}
			})(_89_v0)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw34 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw34).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw34).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_168_v67 = _nw34
		var _source0 D1 = Companion_D1_.Create_DC3_(_168_v67, _dafny.IntOfInt64(712), !((_dafny.IntOfInt64(-154)).Cmp((func() _dafny.Int {
			if (_107_v16).Contains((_161_v62).F24()) {
				return (_107_v16).Multiplicity((_161_v62).F24())
			}
			return _dafny.IntOfInt64(-409)
		})()) >= 0))
		_ = _source0
		if _source0.Is_DC3() {
			var _171___mcc_h0 _dafny.Array = _source0.Get_().(D1_DC3).Cf3
			_ = _171___mcc_h0
			var _172___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC3).Cf4
			_ = _172___mcc_h1
			var _173___mcc_h2 bool = _source0.Get_().(D1_DC3).Cf5
			_ = _173___mcc_h2
			var _174_cf5 bool = _173___mcc_h2
			_ = _174_cf5
			var _175_cf4 _dafny.Int = _172___mcc_h1
			_ = _175_cf4
			var _176_cf3 _dafny.Array = _171___mcc_h0
			_ = _176_cf3
			var _177_v68 D0
			_ = _177_v68
			_177_v68 = Companion_D0_.Create_DC1_(_101_v12)
			var _178_v69 _dafny.Map
			_ = _178_v69
			_178_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v14, _177_v68)
			var _179_v70 _dafny.Int
			_ = _179_v70
			var _180_v71 _dafny.Int
			_ = _180_v71
			var _out16 _dafny.Int
			_ = _out16
			var _out17 _dafny.Int
			_ = _out17
			_out16, _out17 = Companion_Default___.M0((_161_v62).F24(), Companion_Default___.Fm0((_178_v69).Cardinality(), (_161_v62).F24(), _dafny.Companion_Sequence_.Update(_89_v0, (Companion_Default___.SafeIndex(_106_v15, _dafny.IntOfUint32((_89_v0).Cardinality()))).Uint32(), _98_v9), _103_globalState), _106_v15, _103_globalState)
			_179_v70 = _out16
			_180_v71 = _out17
			_91_v2 = _174_cf5
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_159_v60), 0))
			_ = _index31
			(_159_v60).ArraySet1((false) && (_91_v2), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_159_v60), 0))
			_ = _index32
			var _rhs31 bool = ((_100_v11).Cardinality()).Cmp((_dafny.Zero).Minus(_175_cf4)) <= 0
			_ = _rhs31
			var _rhs32 bool = (_161_v62).F24()
			_ = _rhs32
			var _rhs33 bool = (_161_v62).F24()
			_ = _rhs33
			var _lhs23 _dafny.Array = _159_v60
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_159_v60), 0))
			_ = _lhs24
			var _lhs25 *GlobalState = _103_globalState
			_ = _lhs25
			(_lhs23).ArraySet1(_rhs31, (_lhs24).Int())
			_lhs25.F2 = _rhs32
			_174_cf5 = _rhs33
			var _181_v72 _dafny.Int
			_ = _181_v72
			var _182_v73 _dafny.Int
			_ = _182_v73
			var _out18 _dafny.Int
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			_out18, _out19 = Companion_Default___.M0((_dafny.SetOf((_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool))).IsSubsetOf(_dafny.SetOf((_161_v62).F24(), true, _174_cf5)), (_161_v62).F24(), _106_v15, _103_globalState)
			_181_v72 = _out18
			_182_v73 = _out19
		} else if _source0.Is_DC4() {
			var _183___mcc_h3 _dafny.Sequence = _source0.Get_().(D1_DC4).Cf6
			_ = _183___mcc_h3
			var _184___mcc_h4 bool = _source0.Get_().(D1_DC4).Cf7
			_ = _184___mcc_h4
			var _185_cf7 bool = _184___mcc_h4
			_ = _185_cf7
			var _186_cf6 _dafny.Sequence = _183___mcc_h3
			_ = _186_cf6
			var _187_v74 D2
			_ = _187_v74
			_187_v74 = Companion_D2_.Create_DC7_(_106_v15, _dafny.CodePoint('a'), _161_v62, (_161_v62).F24())
			var _188_v75 _dafny.MultiSet
			_ = _188_v75
			_188_v75 = _dafny.MultiSetOf(_161_v62, (_187_v74).Dtor_cf12())
			var _189_v76 _dafny.MultiSet
			_ = _189_v76
			_189_v76 = _dafny.MultiSetOf((_188_v75).Cardinality())
			_106_v15 = (func() _dafny.Int {
				if (_107_v16).Contains((_161_v62).F24()) {
					return (_107_v16).Multiplicity((_161_v62).F24())
				}
				return Companion_Default___.SafeDivisionInt(_105_v14, (_189_v76).Cardinality())
			})()
			var _pat_let_tv0 = _105_v14
			_ = _pat_let_tv0
			var _pat_let_tv1 = _161_v62
			_ = _pat_let_tv1
			var _190_v77 _dafny.Array
			_ = _190_v77
			var _nwElement0_7 D2 = _187_v74
			_ = _nwElement0_7
			var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(16))
			_ = _nw35
			(_nw35).ArraySet1(_nwElement0_7, 0)
			(_nw35).ArraySet1(Companion_D2_.Create_DC7_(_106_v15, _98_v9, _161_v62, (_161_v62).F24()), 1)
			(_nw35).ArraySet1(_187_v74, 2)
			(_nw35).ArraySet1(_187_v74, 3)
			(_nw35).ArraySet1((func() D2 {
				if !(_91_v2) {
					return _187_v74
				}
				return _187_v74
			})(), 4)
			(_nw35).ArraySet1(Companion_D2_.Create_DC7_(_106_v15, _98_v9, _161_v62, _185_cf7), 5)
			(_nw35).ArraySet1(_187_v74, 6)
			(_nw35).ArraySet1(_187_v74, 7)
			(_nw35).ArraySet1(_187_v74, 8)
			(_nw35).ArraySet1(_187_v74, 9)
			(_nw35).ArraySet1(_187_v74, 10)
			(_nw35).ArraySet1(_187_v74, 11)
			(_nw35).ArraySet1((func() D2 {
				if (_161_v62).F24() {
					return _187_v74
				}
				return _187_v74
			})(), 12)
			(_nw35).ArraySet1(_187_v74, 13)
			(_nw35).ArraySet1(func(_pat_let0_0 D2) D2 {
				return func(_191_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let1_0 _dafny.Int) D2 {
						return func(_192_dt__update_hcf10_h0 _dafny.Int) D2 {
							return func(_pat_let2_0 *C0) D2 {
								return func(_193_dt__update_hcf12_h0 *C0) D2 {
									return Companion_D2_.Create_DC7_(_192_dt__update_hcf10_h0, (_191_dt__update__tmp_h0).Dtor_cf11(), _193_dt__update_hcf12_h0, (_191_dt__update__tmp_h0).Dtor_cf13())
								}(_pat_let2_0)
							}(_pat_let_tv1)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_187_v74), 14)
			(_nw35).ArraySet1(_187_v74, 15)
			_190_v77 = _nw35
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_190_v77), 0))
			_ = _index33
			(_190_v77).ArraySet1(_187_v74, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_97_v8), 0))
			_ = _index34
			(_97_v8).ArraySet1(_92_v3, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_190_v77), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_97_v8), 0))
			_ = _index36
			var _rhs34 D2 = (func() D2 {
				if (_105_v14).Cmp(_dafny.IntOfInt64(309)) > 0 {
					return _187_v74
				}
				return Companion_D2_.Create_DC7_((Companion_Default___.Fm6((func() _dafny.Int {
					if (_189_v76).Contains(_106_v15) {
						return (_189_v76).Multiplicity(_106_v15)
					}
					return _92_v3
				})(), (_161_v62).F24(), (_161_v62).F24(), _103_globalState)).Cardinality(), _98_v9, _161_v62, _91_v2)
			})()
			_ = _rhs34
			var _rhs35 _dafny.Int = _106_v15
			_ = _rhs35
			var _rhs36 _dafny.Int = _106_v15
			_ = _rhs36
			var _rhs37 _dafny.Int = _92_v3
			_ = _rhs37
			var _lhs26 _dafny.Array = _190_v77
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_190_v77), 0))
			_ = _lhs27
			var _lhs28 *GlobalState = _103_globalState
			_ = _lhs28
			var _lhs29 *GlobalState = _103_globalState
			_ = _lhs29
			var _lhs30 _dafny.Array = _97_v8
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_97_v8), 0))
			_ = _lhs31
			(_lhs26).ArraySet1(_rhs34, (_lhs27).Int())
			_lhs28.F23 = _rhs35
			_lhs29.F23 = _rhs36
			(_lhs30).ArraySet1(_rhs37, (_lhs31).Int())
			_161_v62 = _161_v62
			_92_v3 = (_dafny.Zero).Minus((_102_v13).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_89_v0).Cardinality()), _dafny.IntOfUint32((_89_v0).Cardinality())), _dafny.IntOfUint32((_102_v13).Cardinality()))).Uint32()).(_dafny.Int))
		} else if _source0.Is_DC2() {
			var _194___mcc_h5 bool = _source0.Get_().(D1_DC2).Cf2
			_ = _194___mcc_h5
			var _195_cf2 bool = _194___mcc_h5
			_ = _195_cf2
			var _196_v78 *C0
			_ = _196_v78
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__((_161_v62).F24())
			_196_v78 = _nw36
			var _197_v79 _dafny.Int
			_ = _197_v79
			var _198_v80 _dafny.Int
			_ = _198_v80
			var _out20 _dafny.Int
			_ = _out20
			var _out21 _dafny.Int
			_ = _out21
			_out20, _out21 = Companion_Default___.M0((_161_v62).F24(), (_196_v78).F24(), _106_v15, _103_globalState)
			_197_v79 = _out20
			_198_v80 = _out21
			(_103_globalState).F18 = _91_v2
			(_103_globalState).F20 = _89_v0
		} else {
			var _199___mcc_h6 D1 = _source0.Get_().(D1_DC5).Cf8
			_ = _199___mcc_h6
			var _200_cf8 D1 = _199___mcc_h6
			_ = _200_cf8
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_159_v60), 0))
			_ = _index37
			(_159_v60).ArraySet1((_161_v62).F24(), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_159_v60), 0))
			_ = _index38
			(_159_v60).ArraySet1(Companion_Default___.Fm0(Companion_Default___.SafeDivisionInt(_106_v15, _106_v15), (_161_v62).F24(), _dafny.Companion_Sequence_.Concatenate(_89_v0, _89_v0), _103_globalState), (_index38).Int())
			_92_v3 = _dafny.IntOfUint32((_89_v0).Cardinality())
			var _201_v81 _dafny.Array
			_ = _201_v81
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_4
			var _nw37 _dafny.Array
			_ = _nw37
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw37 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = (func(_202_v62 *C0) func(_dafny.Int) bool {
					return func(_203_i8 _dafny.Int) bool {
						return (func() bool {
							if (_202_v62).F24() {
								return false
							}
							return (_202_v62).F24()
						})()
					}
				})(_161_v62)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw37 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw37).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw37).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_201_v81 = _nw37
			_201_v81 = _201_v81
			var _204_v82 D2
			_ = _204_v82
			_204_v82 = Companion_D2_.Create_DC7_(_105_v14, _98_v9, _161_v62, (_161_v62).F24())
			var _205_v83 _dafny.Map
			_ = _205_v83
			_205_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v62, !((_161_v62).F24()))
			var _206_v84 _dafny.Int
			_ = _206_v84
			var _207_v85 _dafny.Int
			_ = _207_v85
			var _out22 _dafny.Int
			_ = _out22
			var _out23 _dafny.Int
			_ = _out23
			_out22, _out23 = Companion_Default___.M0(Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-210))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_208_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))).Cardinality()), (_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool), _89_v0, _103_globalState), Companion_Default___.Fm0(_dafny.IntOfInt64(65), (_204_v82).Dtor_cf13(), _89_v0, _103_globalState), (_205_v83).Cardinality(), _103_globalState)
			_206_v84 = _out22
			_207_v85 = _out23
		}
		_107_v16 = _dafny.MultiSetOf(_91_v2, (_dafny.IntOfInt64(741)).Cmp(_106_v15) == 0, (_161_v62).F24())
		var _209_v86 _dafny.Map
		_ = _209_v86
		_209_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(11), _dafny.IntOfInt64(619))
		_92_v3 = ((func() _dafny.Int {
			if (_209_v86).Contains((_dafny.Zero).Minus(_105_v14)) {
				return (_209_v86).Get((_dafny.Zero).Minus(_105_v14)).(_dafny.Int)
			}
			return _106_v15
		})()).Plus(_dafny.IntOfUint32((_89_v0).Cardinality()))
		var _210_v87 _dafny.Int
		_ = _210_v87
		var _211_v88 _dafny.Int
		_ = _211_v88
		var _out24 _dafny.Int
		_ = _out24
		var _out25 _dafny.Int
		_ = _out25
		_out24, _out25 = Companion_Default___.M0((_161_v62).F24(), (_161_v62).F24(), _106_v15, _103_globalState)
		_210_v87 = _out24
		_211_v88 = _out25
	} else {
		var _212_v89 _dafny.Array
		_ = _212_v89
		var _nw38 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(9))
		_ = _nw38
		_212_v89 = _nw38
		var _213_v90 _dafny.Array
		_ = _213_v90
		var _nwElement0_8 _dafny.Sequence = _89_v0
		_ = _nwElement0_8
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(12))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_8, 0)
		(_nw39).ArraySet1(_89_v0, 1)
		(_nw39).ArraySet1(_89_v0, 2)
		(_nw39).ArraySet1(_89_v0, 3)
		(_nw39).ArraySet1(_89_v0, 4)
		(_nw39).ArraySet1(_89_v0, 5)
		(_nw39).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_214_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_215_i10 _dafny.Int) _dafny.CodePoint {
				return _214_v9
			}
		})(_98_v9))), 6)
		(_nw39).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pqgonfj"), 7)
		(_nw39).ArraySet1(_89_v0, 8)
		(_nw39).ArraySet1(_89_v0, 9)
		(_nw39).ArraySet1(_89_v0, 10)
		(_nw39).ArraySet1(_89_v0, 11)
		_213_v90 = _nw39
		var _216_v91 D1
		_ = _216_v91
		_216_v91 = Companion_D1_.Create_DC3_(_213_v90, _106_v15, true)
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_212_v89), 0))
		_ = _index39
		(_212_v89).ArraySet1(_216_v91, (_index39).Int())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_212_v89), 0))
		_ = _index40
		(_212_v89).ArraySet1(_216_v91, (_index40).Int())
		_98_v9 = _dafny.CodePoint('c')
		var _217_v92 _dafny.Map
		_ = _217_v92
		_217_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_105_v14).Minus(_92_v3), _98_v9)
		_217_v92 = (_217_v92).Update(_92_v3, _98_v9)
		(_103_globalState).F18 = (_161_v62).F24()
		var _218_v93 D1
		_ = _218_v93
		_218_v93 = Companion_D1_.Create_DC2_((_105_v14).Cmp(_dafny.IntOfUint32((_102_v13).Cardinality())) >= 0)
		var _pat_let_tv2 = _91_v2
		_ = _pat_let_tv2
		var _rhs38 _dafny.Int = _105_v14
		_ = _rhs38
		var _rhs39 D1 = func(_pat_let3_0 D1) D1 {
			return func(_219_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let4_0 bool) D1 {
					return func(_220_dt__update_hcf2_h0 bool) D1 {
						return Companion_D1_.Create_DC2_(_220_dt__update_hcf2_h0)
					}(_pat_let4_0)
				}(_pat_let_tv2)
			}(_pat_let3_0)
		}((func() D1 {
			if (_161_v62).F24() {
				return _218_v93
			}
			return _218_v93
		})())
		_ = _rhs39
		var _rhs40 _dafny.CodePoint = _98_v9
		_ = _rhs40
		_105_v14 = _rhs38
		_218_v93 = _rhs39
		_98_v9 = _rhs40
	}
	_93_v4 = (_93_v4).Update(_91_v2, (_dafny.Zero).Minus(((_dafny.Zero).Minus(_106_v15)).Minus(_92_v3)))
	var _hi2 _dafny.Int = _106_v15
	_ = _hi2
	for _221_i11 := (_105_v14).Plus(_106_v15); _221_i11.Cmp(_hi2) < 0; _221_i11 = _221_i11.Plus(_dafny.One) {
		var _222_v94 _dafny.MultiSet
		_ = _222_v94
		_222_v94 = _dafny.MultiSetOf(_92_v3)
		if ((_dafny.Zero).Minus((_106_v15).Times(_106_v15))).Cmp((func() _dafny.Int {
			if (_222_v94).Contains(_92_v3) {
				return (_222_v94).Multiplicity(_92_v3)
			}
			return _106_v15
		})()) > 0 {
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_159_v60), 0))
			_ = _index41
			(_159_v60).ArraySet1(_91_v2, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_159_v60), 0))
			_ = _index42
			(_159_v60).ArraySet1((_161_v62).F24(), (_index42).Int())
			var _223_v95 _dafny.Sequence
			_ = _223_v95
			_223_v95 = _dafny.SeqOf(_101_v12, _101_v12, _101_v12, _101_v12)
			(_103_globalState).F15 = (_223_v95).Select((Companion_Default___.SafeIndex(_92_v3, _dafny.IntOfUint32((_223_v95).Cardinality()))).Uint32()).(_dafny.Array)
			_98_v9 = _98_v9
			var _224_v96 _dafny.Sequence
			_ = _224_v96
			_224_v96 = _dafny.SeqOf(true, (_161_v62).F24(), _91_v2)
			var _225_v97 _dafny.Array
			_ = _225_v97
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw40
			_225_v97 = _nw40
			var _226_v98 _dafny.Array
			_ = _226_v98
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(24))
			_ = _nw41
			_226_v98 = _nw41
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_225_v97), 0))
			_ = _index43
			(_225_v97).ArraySet1(_226_v98, (_index43).Int())
			var _227_v99 _dafny.Sequence
			_ = _227_v99
			_227_v99 = _dafny.SeqOf(_226_v98, _226_v98, _226_v98)
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_225_v97), 0))
			_ = _index44
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_159_v60), 0))
			_ = _index45
			var _rhs41 _dafny.Int = (_dafny.Zero).Minus(_221_i11)
			_ = _rhs41
			var _rhs42 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7(_103_globalState), _102_v13)
			_ = _rhs42
			var _rhs43 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_224_v96, _dafny.SeqOf((_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool), (_161_v62).F24()))
			_ = _rhs43
			var _rhs44 _dafny.Array = (_227_v99).Select((Companion_Default___.SafeIndex(_106_v15, _dafny.IntOfUint32((_227_v99).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs44
			var _rhs45 bool = (_92_v3).Cmp(_221_i11) < 0
			_ = _rhs45
			var _lhs32 _dafny.Array = _225_v97
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_225_v97), 0))
			_ = _lhs33
			var _lhs34 _dafny.Array = _159_v60
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_159_v60), 0))
			_ = _lhs35
			_92_v3 = _rhs41
			_102_v13 = _rhs42
			_224_v96 = _rhs43
			(_lhs32).ArraySet1(_rhs44, (_lhs33).Int())
			(_lhs34).ArraySet1(_rhs45, (_lhs35).Int())
			var _228_v100 _dafny.Sequence
			_ = _228_v100
			_228_v100 = _dafny.SeqOf(_107_v16, _107_v16)
			var _229_v101 _dafny.Map
			_ = _229_v101
			_229_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_161_v62).F24()), _dafny.MultiSetFromSeq(_228_v100))
			var _230_v102 _dafny.MultiSet
			_ = _230_v102
			_230_v102 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_224_v96), _107_v16)
			var _231_v103 _dafny.Int
			_ = _231_v103
			var _232_v104 _dafny.Int
			_ = _232_v104
			var _out26 _dafny.Int
			_ = _out26
			var _out27 _dafny.Int
			_ = _out27
			_out26, _out27 = Companion_Default___.M0((_161_v62).F24(), _91_v2, ((func() _dafny.MultiSet {
				if (_229_v101).Contains((_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool)) {
					return (_229_v101).Get((_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool)).(_dafny.MultiSet)
				}
				return _230_v102
			})()).Cardinality(), _103_globalState)
			_231_v103 = _out26
			_232_v104 = _out27
		} else {
			_100_v11 = (_100_v11).Update(_105_v14, (_161_v62).F24())
			var _233_v105 _dafny.Set
			_ = _233_v105
			_233_v105 = _dafny.SetOf(_91_v2)
			var _234_v106 _dafny.Set
			_ = _234_v106
			_234_v106 = _dafny.SetOf((_dafny.SetOf((_161_v62).F24())).Difference(_233_v105), _233_v105, _233_v105, _233_v105, _233_v105)
			_234_v106 = _dafny.SetOf((_233_v105).Difference(_233_v105), _233_v105)
			(_103_globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_89_v0, (Companion_Default___.SafeIndex(_92_v3, _dafny.IntOfUint32((_89_v0).Cardinality()))).Uint32(), _98_v9), _89_v0)
			var _235_v107 D2
			_ = _235_v107
			_235_v107 = Companion_D2_.Create_DC6_(_89_v0)
			_235_v107 = _235_v107
			_92_v3 = _221_i11
		}
		var _236_v108 _dafny.Sequence
		_ = _236_v108
		_236_v108 = _dafny.SeqOf((func() *C0 {
			if (_161_v62).F24() {
				return _161_v62
			}
			return _161_v62
		})(), _161_v62, _161_v62)
		var _237_v109 _dafny.Array
		_ = _237_v109
		var _nwElement0_9 _dafny.Sequence = _89_v0
		_ = _nwElement0_9
		var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
		_ = _nw42
		(_nw42).ArraySet1(_nwElement0_9, 0)
		(_nw42).ArraySet1(_89_v0, 1)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_89_v0, _89_v0), 2)
		(_nw42).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hoeogo"), 3)
		(_nw42).ArraySet1(_89_v0, 4)
		(_nw42).ArraySet1(_89_v0, 5)
		(_nw42).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hsmngmhxs"), 6)
		_237_v109 = _nw42
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_237_v109), 0))
		_ = _index46
		(_237_v109).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ite"), (_index46).Int())
		var _238_v110 _dafny.Map
		_ = _238_v110
		_238_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v0, _89_v0)
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_237_v109), 0))
		_ = _index47
		var _rhs46 _dafny.Sequence = (func() _dafny.Sequence {
			if (_238_v110).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rvomcrdu"), _dafny.UnicodeSeqOfUtf8Bytes("pk"))) {
				return (_238_v110).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rvomcrdu"), _dafny.UnicodeSeqOfUtf8Bytes("pk"))).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Concatenate(_89_v0, _89_v0)
		})()
		_ = _rhs46
		var _rhs47 _dafny.Sequence = _236_v108
		_ = _rhs47
		var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_89_v0, _89_v0), _89_v0)
		_ = _rhs48
		var _lhs36 *GlobalState = _103_globalState
		_ = _lhs36
		var _lhs37 _dafny.Array = _237_v109
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_237_v109), 0))
		_ = _lhs38
		_lhs36.F20 = _rhs46
		_236_v108 = _rhs47
		(_lhs37).ArraySet1(_rhs48, (_lhs38).Int())
		var _239_v111 _dafny.Array
		_ = _239_v111
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_5
		var _nw43 _dafny.Array
		_ = _nw43
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw43 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_240_v16 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_241_i12 _dafny.Int) _dafny.MultiSet {
					return _240_v16
				}
			})(_107_v16)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw43 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw43).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw43).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_239_v111 = _nw43
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_239_v111), 0))
		_ = _index48
		(_239_v111).ArraySet1(_107_v16, (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_239_v111), 0))
		_ = _index49
		(_239_v111).ArraySet1(((_107_v16).Union(_107_v16)).Difference((_107_v16).Update((_161_v62).F24(), Companion_Default___.Abs(_221_i11))), (_index49).Int())
		var _242_v112 _dafny.Set
		_ = _242_v112
		_242_v112 = _dafny.SetOf((_161_v62).F24(), _91_v2)
		var _source1 D1 = Companion_D1_.Create_DC2_((_dafny.SetOf(_91_v2)).IsSubsetOf(_242_v112))
		_ = _source1
		if _source1.Is_DC3() {
			var _243___mcc_h7 _dafny.Array = _source1.Get_().(D1_DC3).Cf3
			_ = _243___mcc_h7
			var _244___mcc_h8 _dafny.Int = _source1.Get_().(D1_DC3).Cf4
			_ = _244___mcc_h8
			var _245___mcc_h9 bool = _source1.Get_().(D1_DC3).Cf5
			_ = _245___mcc_h9
			var _246_cf5 bool = _245___mcc_h9
			_ = _246_cf5
			var _247_cf4 _dafny.Int = _244___mcc_h8
			_ = _247_cf4
			var _248_cf3 _dafny.Array = _243___mcc_h7
			_ = _248_cf3
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_159_v60), 0))
			_ = _index50
			(_159_v60).ArraySet1(false, (_index50).Int())
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_159_v60), 0))
			_ = _index51
			(_159_v60).ArraySet1((_221_i11).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_92_v3, _106_v15)))) == 0, (_index51).Int())
			var _249_v113 _dafny.Map
			_ = _249_v113
			_249_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v14, _247_cf4)
			_249_v113 = (_249_v113).Update(_dafny.IntOfInt64(103), (_105_v14).Times(_106_v15))
			var _250_v114 *C0
			_ = _250_v114
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__((_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool))
			_250_v114 = _nw44
			(_103_globalState).F18 = (_246_cf5) == ((_161_v62).F24())
		} else if _source1.Is_DC4() {
			var _251___mcc_h10 _dafny.Sequence = _source1.Get_().(D1_DC4).Cf6
			_ = _251___mcc_h10
			var _252___mcc_h11 bool = _source1.Get_().(D1_DC4).Cf7
			_ = _252___mcc_h11
			var _253_cf7 bool = _252___mcc_h11
			_ = _253_cf7
			var _254_cf6 _dafny.Sequence = _251___mcc_h10
			_ = _254_cf6
			var _255_v115 _dafny.Map
			_ = _255_v115
			_255_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v13, (_161_v62).F24())
			(_103_globalState).F23 = ((_255_v115).Cardinality()).Times((_dafny.IntOfInt64(-803)).Times(_dafny.IntOfInt64(-924)))
			(_103_globalState).F23 = _105_v14
			var _256_v116 _dafny.Int
			_ = _256_v116
			var _257_v117 _dafny.Int
			_ = _257_v117
			var _out28 _dafny.Int
			_ = _out28
			var _out29 _dafny.Int
			_ = _out29
			_out28, _out29 = Companion_Default___.M0(_91_v2, true, (_dafny.Zero).Minus(_221_i11), _103_globalState)
			_256_v116 = _out28
			_257_v117 = _out29
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_159_v60), 0))
			_ = _index52
			(_159_v60).ArraySet1((_161_v62).F24(), (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_159_v60), 0))
			_ = _index53
			(_159_v60).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_89_v0, _89_v0), (_index53).Int())
			var _258_v119 D1
			_ = _258_v119
			_258_v119 = Companion_D1_.Create_DC4_(_254_cf6, _91_v2)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_159_v60), 0))
			_ = _index54
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_159_v60), 0))
			_ = _index55
			var _rhs49 bool = !(func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(248), _dafny.IntOfInt64(-932))); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _259_v118 _dafny.Int
					_259_v118 = interface{}(_compr_6).(_dafny.Int)
					if ((_dafny.IntOfInt64(248)).Cmp(_259_v118) <= 0) && ((_259_v118).Cmp(_dafny.IntOfInt64(-932)) < 0) {
						_coll6.Add((_259_v118).Plus(_106_v15))
					}
				}
				return _coll6.ToSet()
			}()).Contains(_dafny.IntOfInt64(728))
			_ = _rhs49
			var _rhs50 *C0 = _161_v62
			_ = _rhs50
			var _rhs51 bool = (_161_v62).F24()
			_ = _rhs51
			var _rhs52 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_254_cf6, (_258_v119).Dtor_cf6())
			_ = _rhs52
			var _rhs53 bool = (_161_v62).F24()
			_ = _rhs53
			var _lhs39 _dafny.Array = _159_v60
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_159_v60), 0))
			_ = _lhs40
			var _lhs41 *GlobalState = _103_globalState
			_ = _lhs41
			var _lhs42 _dafny.Array = _159_v60
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_159_v60), 0))
			_ = _lhs43
			(_lhs39).ArraySet1(_rhs49, (_lhs40).Int())
			_161_v62 = _rhs50
			_lhs41.F2 = _rhs51
			_254_cf6 = _rhs52
			(_lhs42).ArraySet1(_rhs53, (_lhs43).Int())
		} else if _source1.Is_DC2() {
			var _260___mcc_h12 bool = _source1.Get_().(D1_DC2).Cf2
			_ = _260___mcc_h12
			var _261_cf2 bool = _260___mcc_h12
			_ = _261_cf2
			_105_v14 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(978), _92_v3)).Minus((_221_i11).Plus(_221_i11))
			var _262_v120 _dafny.Array
			_ = _262_v120
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_6
			var _nw45 _dafny.Array
			_ = _nw45
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw45 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Map = (func(_263_i11 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_264_i13 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_263_i11))
					}
				})(_221_i11)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw45 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw45).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw45).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_262_v120 = _nw45
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_262_v120), 0))
			_ = _index56
			(_262_v120).ArraySet1(_93_v4, (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_262_v120), 0))
			_ = _index57
			(_262_v120).ArraySet1(_93_v4, (_index57).Int())
			var _265_v121 _dafny.Array
			_ = _265_v121
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw46
			_265_v121 = _nw46
			var _266_v122 _dafny.Map
			_ = _266_v122
			_266_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_239_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_239_v111), 0))).Int()).(_dafny.MultiSet)).Difference((_239_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_239_v111), 0))).Int()).(_dafny.MultiSet)), _161_v62)
			var _267_v123 D3
			_ = _267_v123
			_267_v123 = Companion_D3_.Create_DC9_(_159_v60)
			var _268_v124 _dafny.Map
			_ = _268_v124
			_268_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pwflyn"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(_261_cf2, (_161_v62).F24())), _161_v62))
			var _rhs54 _dafny.Array = _265_v121
			_ = _rhs54
			var _rhs55 *C0 = _161_v62
			_ = _rhs55
			var _rhs56 _dafny.Array = (_267_v123).Dtor_cf15()
			_ = _rhs56
			var _rhs57 bool = _91_v2
			_ = _rhs57
			var _rhs58 _dafny.Map = (func() _dafny.Map {
				if (_268_v124).Contains(_89_v0) {
					return (_268_v124).Get(_89_v0).(_dafny.Map)
				}
				return (_266_v122).Merge(_266_v122)
			})()
			_ = _rhs58
			_265_v121 = _rhs54
			_161_v62 = _rhs55
			_159_v60 = _rhs56
			_261_cf2 = _rhs57
			_266_v122 = _rhs58
			(_103_globalState).F23 = (_dafny.IntOfInt64(810)).Plus(_106_v15)
		} else {
			var _269___mcc_h13 D1 = _source1.Get_().(D1_DC5).Cf8
			_ = _269___mcc_h13
			var _270_cf8 D1 = _269___mcc_h13
			_ = _270_cf8
			var _271_v125 _dafny.Map
			_ = _271_v125
			_271_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_221_i11, _105_v14)
			var _272_v126 _dafny.Map
			_ = _272_v126
			_272_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_166_v65).Contains((_161_v62).F24()) {
					return (_166_v65).Get((_161_v62).F24()).(bool)
				}
				return (_161_v62).F24()
			})(), _271_v125)
			var _273_v127 _dafny.Map
			_ = _273_v127
			_273_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v3, (func() _dafny.Int {
				if (_93_v4).Contains(false) {
					return (_93_v4).Get(false).(_dafny.Int)
				}
				return (_272_v126).Cardinality()
			})())
			(_103_globalState).F23 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_161_v62).F24() {
					return _dafny.IntOfInt64(422)
				}
				return _dafny.IntOfUint32((_89_v0).Cardinality())
			})(), (_dafny.Zero).Minus((_dafny.Zero).Minus(((_dafny.Zero).Minus(_105_v14)).Times((_dafny.Zero).Minus((_273_v127).Cardinality())))))
			(_103_globalState).F20 = (func() _dafny.Sequence {
				if _91_v2 {
					return (_161_v62).Fm5(_103_globalState)
				}
				return _89_v0
			})()
			var _274_v128 _dafny.Sequence
			_ = _274_v128
			_274_v128 = _dafny.SeqOf((_161_v62).F24())
			var _275_v129 _dafny.Map
			_ = _275_v129
			_275_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.Companion_Sequence_.Update(_274_v128, (Companion_Default___.SafeIndex(_106_v15, _dafny.IntOfUint32((_274_v128).Cardinality()))).Uint32(), !((_161_v62).F24())), (_161_v62).F24()), _106_v15)
			_275_v129 = (_275_v129).Update(Companion_Default___.Fm8(_106_v15, _dafny.IntOfUint32((_102_v13).Cardinality()), !((_161_v62).F24()), _103_globalState), (_dafny.Zero).Minus(_221_i11))
			(_103_globalState).F23 = _221_i11
		}
	}
	var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))
	_ = _index58
	(_159_v60).ArraySet1((_161_v62).F24(), (_index58).Int())
	var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))
	_ = _index59
	(_159_v60).ArraySet1(((_92_v3).Times(_105_v14)).Cmp(_105_v14) < 0, (_index59).Int())
	var _276_i14 _dafny.Int
	_ = _276_i14
	_276_i14 = _dafny.Zero
	{
		for (func() bool {
			if (_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool) {
				return (_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool)
			}
			return (_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool)
		})() {
			{
				if (_276_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_276_i14 = (_276_i14).Plus(_dafny.One)
				(_103_globalState).F2 = (_91_v2) || ((_161_v62).F24())
				(_103_globalState).F23 = _dafny.IntOfInt64(-84)
				var _277_v130 _dafny.Sequence
				_ = _277_v130
				_277_v130 = _dafny.SeqOf((_161_v62).F24())
				var _278_v131 _dafny.Int
				_ = _278_v131
				var _279_v132 _dafny.Int
				_ = _279_v132
				var _out30 _dafny.Int
				_ = _out30
				var _out31 _dafny.Int
				_ = _out31
				_out30, _out31 = Companion_Default___.M0((_277_v130).Select((Companion_Default___.SafeIndex(_92_v3, _dafny.IntOfUint32((_277_v130).Cardinality()))).Uint32()).(bool), _91_v2, _105_v14, _103_globalState)
				_278_v131 = _out30
				_279_v132 = _out31
				var _280_v133 _dafny.Int
				_ = _280_v133
				var _281_v134 _dafny.Int
				_ = _281_v134
				var _out32 _dafny.Int
				_ = _out32
				var _out33 _dafny.Int
				_ = _out33
				_out32, _out33 = Companion_Default___.M0((_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool), (Companion_Default___.Fm3(_dafny.IntOfInt64(984), _103_globalState)).Cmp((_161_v62).Fm4(false, _91_v2, _dafny.MultiSetFromSeq(Companion_Default___.Fm7(_103_globalState)), _91_v2, _103_globalState)) == 0, _106_v15, _103_globalState)
				_280_v133 = _out32
				_281_v134 = _out33
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _rhs59 _dafny.Array = _97_v8
	_ = _rhs59
	var _rhs60 _dafny.Int = _dafny.IntOfInt64(-668)
	_ = _rhs60
	_97_v8 = _rhs59
	_92_v3 = _rhs60
	var _282_v135 _dafny.Set
	_ = _282_v135
	_282_v135 = _dafny.SetOf(_91_v2, _91_v2)
	var _283_v136 _dafny.Array
	_ = _283_v136
	var _nwElement0_10 *C0 = _161_v62
	_ = _nwElement0_10
	var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(4))
	_ = _nw47
	(_nw47).ArraySet1(_nwElement0_10, 0)
	(_nw47).ArraySet1(_161_v62, 1)
	(_nw47).ArraySet1(_161_v62, 2)
	(_nw47).ArraySet1(_161_v62, 3)
	_283_v136 = _nw47
	var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_283_v136), 0))
	_ = _index60
	(_283_v136).ArraySet1(_161_v62, (_index60).Int())
	var _284_v137 D0
	_ = _284_v137
	_284_v137 = Companion_D0_.Create_DC0_(_98_v9)
	var _285_v138 _dafny.Map
	_ = _285_v138
	_285_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_284_v137, _161_v62)
	var _pat_let_tv3 = _98_v9
	_ = _pat_let_tv3
	var _pat_let_tv4 = _98_v9
	_ = _pat_let_tv4
	var _286_v139 _dafny.Sequence
	_ = _286_v139
	_286_v139 = _dafny.SeqOf((func() *C0 {
		if (_285_v138).Contains(func(_pat_let5_0 D0) D0 {
			return func(_287_dt__update__tmp_h2 D0) D0 {
				return func(_pat_let6_0 _dafny.CodePoint) D0 {
					return func(_288_dt__update_hcf0_h0 _dafny.CodePoint) D0 {
						return Companion_D0_.Create_DC0_(_288_dt__update_hcf0_h0)
					}(_pat_let6_0)
				}(_pat_let_tv3)
			}(_pat_let5_0)
		}(_284_v137)) {
			return (_285_v138).Get(func(_pat_let7_0 D0) D0 {
				return func(_289_dt__update__tmp_h3 D0) D0 {
					return func(_pat_let8_0 _dafny.CodePoint) D0 {
						return func(_290_dt__update_hcf0_h1 _dafny.CodePoint) D0 {
							return Companion_D0_.Create_DC0_(_290_dt__update_hcf0_h1)
						}(_pat_let8_0)
					}(_pat_let_tv4)
				}(_pat_let7_0)
			}(_284_v137)).(*C0)
		}
		return _161_v62
	})(), _161_v62)
	var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_283_v136), 0))
	_ = _index61
	var _rhs61 _dafny.Array = _159_v60
	_ = _rhs61
	var _rhs62 _dafny.Set = _282_v135
	_ = _rhs62
	var _rhs63 bool = !((func() bool {
		if (_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool) {
			return (_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool)
		}
		return (_161_v62).F24()
	})())
	_ = _rhs63
	var _rhs64 *C0 = (_286_v139).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.IntOfUint32((_286_v139).Cardinality()))).Uint32()).(*C0)
	_ = _rhs64
	var _lhs44 *GlobalState = _103_globalState
	_ = _lhs44
	var _lhs45 _dafny.Array = _283_v136
	_ = _lhs45
	var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_283_v136), 0))
	_ = _lhs46
	_159_v60 = _rhs61
	_282_v135 = _rhs62
	_lhs44.F2 = _rhs63
	(_lhs45).ArraySet1(_rhs64, (_lhs46).Int())
	var _nw48 *C0 = New_C0_()
	_ = _nw48
	_nw48.Ctor__(!((_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool)))
	_161_v62 = _nw48
	_106_v15 = (_105_v14).Minus((_105_v14).Plus(_dafny.IntOfUint32((_102_v13).Cardinality())))
	var _291_v140 _dafny.MultiSet
	_ = _291_v140
	_291_v140 = _dafny.MultiSetOf(Companion_Default___.Fm9(_106_v15, (_159_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_159_v60), 0))).Int()).(bool), _92_v3, _91_v2, _103_globalState), _100_v11)
	var _292_v141 _dafny.Map
	_ = _292_v141
	_292_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v140, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
		if (_161_v62).F24() {
			return _dafny.UnicodeSeqOfUtf8Bytes("balyols")
		}
		return _89_v0
	})(), (Companion_Default___.SafeIndex(_106_v15, _dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_161_v62).F24() {
			return _dafny.UnicodeSeqOfUtf8Bytes("balyols")
		}
		return _89_v0
	})()).Cardinality()))).Uint32(), _98_v9))
	var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_283_v136), 0))
	_ = _index62
	var _rhs65 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_102_v13, _dafny.SeqOf(_92_v3))
	_ = _rhs65
	var _rhs66 *C0 = (_283_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_283_v136), 0))).Int()).(*C0)
	_ = _rhs66
	var _rhs67 _dafny.Array = _159_v60
	_ = _rhs67
	var _rhs68 _dafny.Sequence = (func() _dafny.Sequence {
		if (_292_v141).Contains((_291_v140).Difference(Companion_Default___.Fm10(_91_v2, _91_v2, (_161_v62).F24(), _103_globalState))) {
			return (_292_v141).Get((_291_v140).Difference(Companion_Default___.Fm10(_91_v2, _91_v2, (_161_v62).F24(), _103_globalState))).(_dafny.Sequence)
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("nhrm")
	})()
	_ = _rhs68
	var _lhs47 _dafny.Array = _283_v136
	_ = _lhs47
	var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_283_v136), 0))
	_ = _lhs48
	var _lhs49 *GlobalState = _103_globalState
	_ = _lhs49
	_102_v13 = _rhs65
	(_lhs47).ArraySet1(_rhs66, (_lhs48).Int())
	_159_v60 = _rhs67
	_lhs49.F20 = _rhs68
	var _293_v142 *C0
	_ = _293_v142
	var _nw49 *C0 = New_C0_()
	_ = _nw49
	_nw49.Ctor__(true)
	_293_v142 = _nw49
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_97_v8), 0))); ; {
		_guard_loop_0, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _294_i15 _dafny.Int
		_294_i15 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_294_i15).Sign() != -1) && ((_294_i15).Cmp(_dafny.ArrayLen((_97_v8), 0)) < 0)) {
			(_97_v8).ArraySet1(Companion_Default___.SafeModuloInt(_294_i15, ((func() _dafny.Map {
				if (_293_v142).F24() {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v14, _92_v3)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v3, _92_v3)
			})()).Cardinality()), (_294_i15).Int())
		}
	}
	_dafny.Print(_89_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v1).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_91_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_92_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(56)).UpdateUnsafe(false, _dafny.IntOfInt64(116))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_94_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(56)), _dafny.IntOfInt64(56))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v7).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v8).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_v9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_99_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('n'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(56), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v12).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_102_v13, _dafny.SeqOf(_dafny.IntOfInt64(56), _dafny.IntOfInt64(56), _dafny.IntOfInt64(-668))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F1).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"), _dafny.UnicodeSeqOfUtf8Bytes("spiw"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_103_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F3()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(56)), _dafny.IntOfInt64(56))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F14().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState.F15).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_103_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_103_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_103_globalState.F20.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_103_globalState).F21(), _dafny.SeqOf(_dafny.IntOfInt64(56), _dafny.IntOfInt64(56))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_103_globalState).F22(), _dafny.SeqOf(_dafny.IntOfInt64(56), _dafny.IntOfInt64(56))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_103_globalState.F23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_v14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_106_v15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_107_v16).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v60).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v60).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v61).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v62).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_v63).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v64).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v65).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_276_i14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_282_v135).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v136).ArrayGet1((_dafny.Zero).Int()).(*C0)).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v136).ArrayGet1((_dafny.One).Int()).(*C0)).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v136).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C0)).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v136).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C0)).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v137).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_285_v138).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_286_v139).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_291_v140).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(21), false).UpdateUnsafe(_dafny.IntOfInt64(22), false).UpdateUnsafe(_dafny.IntOfInt64(23), false).UpdateUnsafe(_dafny.IntOfInt64(24), false).UpdateUnsafe(_dafny.IntOfInt64(25), false).UpdateUnsafe(_dafny.IntOfInt64(26), false).UpdateUnsafe(_dafny.IntOfInt64(27), false).UpdateUnsafe(_dafny.IntOfInt64(28), false).UpdateUnsafe(_dafny.IntOfInt64(29), false).UpdateUnsafe(_dafny.IntOfInt64(30), false).UpdateUnsafe(_dafny.IntOfInt64(31), false).UpdateUnsafe(_dafny.IntOfInt64(32), false).UpdateUnsafe(_dafny.IntOfInt64(33), false).UpdateUnsafe(_dafny.IntOfInt64(34), false).UpdateUnsafe(_dafny.IntOfInt64(35), false).UpdateUnsafe(_dafny.IntOfInt64(36), false).UpdateUnsafe(_dafny.IntOfInt64(37), false).UpdateUnsafe(_dafny.IntOfInt64(38), false).UpdateUnsafe(_dafny.IntOfInt64(39), false).UpdateUnsafe(_dafny.IntOfInt64(40), false).UpdateUnsafe(_dafny.IntOfInt64(41), false).UpdateUnsafe(_dafny.IntOfInt64(42), false).UpdateUnsafe(_dafny.IntOfInt64(43), false).UpdateUnsafe(_dafny.IntOfInt64(44), false).UpdateUnsafe(_dafny.IntOfInt64(45), false).UpdateUnsafe(_dafny.IntOfInt64(46), false).UpdateUnsafe(_dafny.IntOfInt64(47), false).UpdateUnsafe(_dafny.IntOfInt64(48), false).UpdateUnsafe(_dafny.IntOfInt64(49), false).UpdateUnsafe(_dafny.IntOfInt64(50), false).UpdateUnsafe(_dafny.IntOfInt64(51), false).UpdateUnsafe(_dafny.IntOfInt64(52), false).UpdateUnsafe(_dafny.IntOfInt64(53), false).UpdateUnsafe(_dafny.IntOfInt64(54), false).UpdateUnsafe(_dafny.IntOfInt64(55), false).UpdateUnsafe(_dafny.IntOfInt64(56), false).UpdateUnsafe(_dafny.IntOfInt64(57), false).UpdateUnsafe(_dafny.IntOfInt64(58), false).UpdateUnsafe(_dafny.Zero, false).UpdateUnsafe(_dafny.One, false).UpdateUnsafe(_dafny.IntOfInt64(2), false).UpdateUnsafe(_dafny.IntOfInt64(3), false).UpdateUnsafe(_dafny.IntOfInt64(4), false).UpdateUnsafe(_dafny.IntOfInt64(5), false).UpdateUnsafe(_dafny.IntOfInt64(6), false).UpdateUnsafe(_dafny.IntOfInt64(7), false).UpdateUnsafe(_dafny.IntOfInt64(8), false).UpdateUnsafe(_dafny.IntOfInt64(9), false).UpdateUnsafe(_dafny.IntOfInt64(10), false).UpdateUnsafe(_dafny.IntOfInt64(11), false).UpdateUnsafe(_dafny.IntOfInt64(12), false).UpdateUnsafe(_dafny.IntOfInt64(13), false).UpdateUnsafe(_dafny.IntOfInt64(14), false).UpdateUnsafe(_dafny.IntOfInt64(15), false).UpdateUnsafe(_dafny.IntOfInt64(16), false).UpdateUnsafe(_dafny.IntOfInt64(17), false).UpdateUnsafe(_dafny.IntOfInt64(18), false).UpdateUnsafe(_dafny.IntOfInt64(19), false).UpdateUnsafe(_dafny.IntOfInt64(20), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(56), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_292_v141).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(21), false).UpdateUnsafe(_dafny.IntOfInt64(22), false).UpdateUnsafe(_dafny.IntOfInt64(23), false).UpdateUnsafe(_dafny.IntOfInt64(24), false).UpdateUnsafe(_dafny.IntOfInt64(25), false).UpdateUnsafe(_dafny.IntOfInt64(26), false).UpdateUnsafe(_dafny.IntOfInt64(27), false).UpdateUnsafe(_dafny.IntOfInt64(28), false).UpdateUnsafe(_dafny.IntOfInt64(29), false).UpdateUnsafe(_dafny.IntOfInt64(30), false).UpdateUnsafe(_dafny.IntOfInt64(31), false).UpdateUnsafe(_dafny.IntOfInt64(32), false).UpdateUnsafe(_dafny.IntOfInt64(33), false).UpdateUnsafe(_dafny.IntOfInt64(34), false).UpdateUnsafe(_dafny.IntOfInt64(35), false).UpdateUnsafe(_dafny.IntOfInt64(36), false).UpdateUnsafe(_dafny.IntOfInt64(37), false).UpdateUnsafe(_dafny.IntOfInt64(38), false).UpdateUnsafe(_dafny.IntOfInt64(39), false).UpdateUnsafe(_dafny.IntOfInt64(40), false).UpdateUnsafe(_dafny.IntOfInt64(41), false).UpdateUnsafe(_dafny.IntOfInt64(42), false).UpdateUnsafe(_dafny.IntOfInt64(43), false).UpdateUnsafe(_dafny.IntOfInt64(44), false).UpdateUnsafe(_dafny.IntOfInt64(45), false).UpdateUnsafe(_dafny.IntOfInt64(46), false).UpdateUnsafe(_dafny.IntOfInt64(47), false).UpdateUnsafe(_dafny.IntOfInt64(48), false).UpdateUnsafe(_dafny.IntOfInt64(49), false).UpdateUnsafe(_dafny.IntOfInt64(50), false).UpdateUnsafe(_dafny.IntOfInt64(51), false).UpdateUnsafe(_dafny.IntOfInt64(52), false).UpdateUnsafe(_dafny.IntOfInt64(53), false).UpdateUnsafe(_dafny.IntOfInt64(54), false).UpdateUnsafe(_dafny.IntOfInt64(55), false).UpdateUnsafe(_dafny.IntOfInt64(56), false).UpdateUnsafe(_dafny.IntOfInt64(57), false).UpdateUnsafe(_dafny.IntOfInt64(58), false).UpdateUnsafe(_dafny.Zero, false).UpdateUnsafe(_dafny.One, false).UpdateUnsafe(_dafny.IntOfInt64(2), false).UpdateUnsafe(_dafny.IntOfInt64(3), false).UpdateUnsafe(_dafny.IntOfInt64(4), false).UpdateUnsafe(_dafny.IntOfInt64(5), false).UpdateUnsafe(_dafny.IntOfInt64(6), false).UpdateUnsafe(_dafny.IntOfInt64(7), false).UpdateUnsafe(_dafny.IntOfInt64(8), false).UpdateUnsafe(_dafny.IntOfInt64(9), false).UpdateUnsafe(_dafny.IntOfInt64(10), false).UpdateUnsafe(_dafny.IntOfInt64(11), false).UpdateUnsafe(_dafny.IntOfInt64(12), false).UpdateUnsafe(_dafny.IntOfInt64(13), false).UpdateUnsafe(_dafny.IntOfInt64(14), false).UpdateUnsafe(_dafny.IntOfInt64(15), false).UpdateUnsafe(_dafny.IntOfInt64(16), false).UpdateUnsafe(_dafny.IntOfInt64(17), false).UpdateUnsafe(_dafny.IntOfInt64(18), false).UpdateUnsafe(_dafny.IntOfInt64(19), false).UpdateUnsafe(_dafny.IntOfInt64(20), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(56), true)), _dafny.UnicodeSeqOfUtf8Bytes("nalyols"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_293_v142).F24())
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
	Cf1 _dafny.Array
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.CodePoint
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.CodePoint) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf0() _dafny.CodePoint {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf1 == data2.Cf1
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

type D1_DC3 struct {
	Cf3 _dafny.Array
	Cf4 _dafny.Int
	Cf5 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Array, Cf4 _dafny.Int, Cf5 bool) D1 {
	return D1{D1_DC3{Cf3, Cf4, Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf6 _dafny.Sequence
	Cf7 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.Sequence, Cf7 bool) D1 {
	return D1{D1_DC4{Cf6, Cf7}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf2 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 bool) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC5 struct {
	Cf8 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf8 D1) D1 {
	return D1{D1_DC5{Cf8}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, false)
}

func (_this D1) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf8() D1 {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf8) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7 == data2.Cf7
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2 == data2.Cf2
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf8.Equals(data2.Cf8)
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

type D2_DC7 struct {
	Cf10 _dafny.Int
	Cf11 _dafny.CodePoint
	Cf12 *C0
	Cf13 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf10 _dafny.Int, Cf11 _dafny.CodePoint, Cf12 *C0, Cf13 bool) D2 {
	return D2{D2_DC7{Cf10, Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf9 _dafny.Sequence
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf9 _dafny.Sequence) D2 {
	return D2{D2_DC6{Cf9}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC8 struct {
	Cf14 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf14 D2) D2 {
	return D2{D2_DC8{Cf14}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.Zero, _dafny.CodePoint('D'), (*C0)(nil), false)
}

func (_this D2) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf10
}

func (_this D2) Dtor_cf11() _dafny.CodePoint {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() *C0 {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf9
}

func (_this D2) Dtor_cf14() D2 {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + data.Cf9.VerbatimString(true) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf14) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf9.Equals(data2.Cf9)
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D3_DC10 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.CodePoint
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 _dafny.Int, Cf17 _dafny.CodePoint) D3 {
	return D3{D3_DC10{Cf16, Cf17}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf18 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf18 bool) D3 {
	return D3{D3_DC11{Cf18}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf15 _dafny.Array
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf15 _dafny.Array) D3 {
	return D3{D3_DC9{Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC12 struct {
	Cf19 D3
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf19 D3) D3 {
	return D3{D3_DC12{Cf19}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC11).Cf18
}

func (_this D3) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) Dtor_cf19() D3 {
	return _this.Get_().(D3_DC12).Cf19
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf18 == data2.Cf18
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf19.Equals(data2.Cf19)
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

type D4_DC14 struct {
	Cf21 bool
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf21 bool) D4 {
	return D4{D4_DC14{Cf21}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC13 struct {
	Cf20 _dafny.Sequence
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf20 _dafny.Sequence) D4 {
	return D4{D4_DC13{Cf20}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_(false)
}

func (_this D4) Dtor_cf21() bool {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf20) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf21 == data2.Cf21
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf20.Equals(data2.Cf20)
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

type D5_DC16 struct {
	Cf23 bool
	Cf24 _dafny.Sequence
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf23 bool, Cf24 _dafny.Sequence) D5 {
	return D5{D5_DC16{Cf23, Cf24}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC15 struct {
	Cf22 _dafny.Array
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf22 _dafny.Array) D5 {
	return D5{D5_DC15{Cf22}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC17 struct {
	Cf25 D5
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf25 D5) D5 {
	return D5{D5_DC17{Cf25}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC16_(false, _dafny.EmptySeq)
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC16).Cf23
}

func (_this D5) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D5_DC16).Cf24
}

func (_this D5) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D5_DC15).Cf22
}

func (_this D5) Dtor_cf25() D5 {
	return _this.Get_().(D5_DC17).Cf25
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Equals(data2.Cf24)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf22 == data2.Cf22
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

type D6_DC19 struct {
	Cf27 _dafny.Map
	Cf28 _dafny.Set
	Cf29 bool
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf27 _dafny.Map, Cf28 _dafny.Set, Cf29 bool) D6 {
	return D6{D6_DC19{Cf27, Cf28, Cf29}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC20 struct {
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_() D6 {
	return D6{D6_DC20{}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

type D6_DC18 struct {
	Cf26 _dafny.Map
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf26 _dafny.Map) D6 {
	return D6{D6_DC18{Cf26}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC19_(_dafny.EmptyMap, _dafny.EmptySet, false)
}

func (_this D6) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D6_DC19).Cf27
}

func (_this D6) Dtor_cf28() _dafny.Set {
	return _this.Get_().(D6_DC19).Cf28
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC19).Cf29
}

func (_this D6) Dtor_cf26() _dafny.Map {
	return _this.Get_().(D6_DC18).Cf26
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D6_DC20:
		{
			return "D6.DC20"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf27.Equals(data2.Cf27) && data1.Cf28.Equals(data2.Cf28) && data1.Cf29 == data2.Cf29
		}
	case D6_DC20:
		{
			_, ok := other.Get_().(D6_DC20)
			return ok
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D7_DC21 struct {
	Cf30 _dafny.Map
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf30 _dafny.Map) D7 {
	return D7{D7_DC21{Cf30}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D7) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D7_DC21).Cf30
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

// Definition of datatype D8
type D8 struct {
	Data_D8_
}

func (_this D8) Get_() Data_D8_ {
	return _this.Data_D8_
}

type Data_D8_ interface {
	isD8()
}

type CompanionStruct_D8_ struct {
}

var Companion_D8_ = CompanionStruct_D8_{}

type D8_DC23 struct {
	Cf32 _dafny.Int
	Cf33 _dafny.Sequence
	Cf34 _dafny.CodePoint
	Cf35 _dafny.Int
	Cf36 _dafny.Int
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf32 _dafny.Int, Cf33 _dafny.Sequence, Cf34 _dafny.CodePoint, Cf35 _dafny.Int, Cf36 _dafny.Int) D8 {
	return D8{D8_DC23{Cf32, Cf33, Cf34, Cf35, Cf36}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf37 _dafny.Int
	Cf38 bool
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf37 _dafny.Int, Cf38 bool) D8 {
	return D8{D8_DC24{Cf37, Cf38}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC25 struct {
	Cf39 bool
	Cf40 _dafny.Int
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf39 bool, Cf40 _dafny.Int) D8 {
	return D8{D8_DC25{Cf39, Cf40}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC22 struct {
	Cf31 _dafny.Map
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf31 _dafny.Map) D8 {
	return D8{D8_DC22{Cf31}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(_dafny.Zero, _dafny.EmptySeq, _dafny.CodePoint('D'), _dafny.Zero, _dafny.Zero)
}

func (_this D8) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf32
}

func (_this D8) Dtor_cf33() _dafny.Sequence {
	return _this.Get_().(D8_DC23).Cf33
}

func (_this D8) Dtor_cf34() _dafny.CodePoint {
	return _this.Get_().(D8_DC23).Cf34
}

func (_this D8) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf35
}

func (_this D8) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf37
}

func (_this D8) Dtor_cf38() bool {
	return _this.Get_().(D8_DC24).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC25).Cf39
}

func (_this D8) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D8_DC25).Cf40
}

func (_this D8) Dtor_cf31() _dafny.Map {
	return _this.Get_().(D8_DC22).Cf31
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf32) + ", " + data.Cf33.VerbatimString(true) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Equals(data2.Cf33) && data1.Cf34 == data2.Cf34 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Cmp(data2.Cf36) == 0
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38 == data2.Cf38
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf39 == data2.Cf39 && data1.Cf40.Cmp(data2.Cf40) == 0
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf31.Equals(data2.Cf31)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D8) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D8)
	return ok && _this.Equals(typed)
}

func Type_D8_() _dafny.TypeDescriptor {
	return type_D8_{}
}

type type_D8_ struct {
}

func (_this type_D8_) Default() interface{} {
	return Companion_D8_.Default()
}

func (_this type_D8_) String() string {
	return "main.D8"
}
func (_this D8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D8{}

// End of datatype D8

// Definition of class GlobalState
type GlobalState struct {
	F1   _dafny.MultiSet
	F2   bool
	F15  _dafny.Array
	F18  bool
	F20  _dafny.Sequence
	F23  _dafny.Int
	_f0  _dafny.Int
	_f3  _dafny.Map
	_f4  _dafny.Int
	_f5  bool
	_f6  bool
	_f7  _dafny.Int
	_f8  _dafny.Int
	_f9  bool
	_f10 _dafny.Array
	_f11 _dafny.CodePoint
	_f12 _dafny.Int
	_f13 _dafny.Array
	_f14 _dafny.Sequence
	_f16 _dafny.Int
	_f17 _dafny.Array
	_f19 bool
	_f21 _dafny.Sequence
	_f22 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.EmptyMultiSet
	_this.F2 = false
	_this.F15 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F18 = false
	_this.F20 = _dafny.EmptySeq
	_this.F23 = _dafny.Zero
	_this._f0 = _dafny.Zero
	_this._f3 = _dafny.EmptyMap
	_this._f4 = _dafny.Zero
	_this._f5 = false
	_this._f6 = false
	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f9 = false
	_this._f10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f11 = _dafny.CodePoint('D')
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f14 = _dafny.EmptySeq
	_this._f16 = _dafny.Zero
	_this._f17 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f19 = false
	_this._f21 = _dafny.EmptySeq
	_this._f22 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.MultiSet, f2 bool, f3 _dafny.Map, f4 _dafny.Int, f5 bool, f6 bool, f7 _dafny.Int, f8 _dafny.Int, f9 bool, f10 _dafny.Array, f11 _dafny.CodePoint, f12 _dafny.Int, f13 _dafny.Array, f14 _dafny.Sequence, f15 _dafny.Array, f16 _dafny.Int, f17 _dafny.Array, f18 bool, f19 bool, f20 _dafny.Sequence, f21 _dafny.Sequence, f22 _dafny.Sequence, f23 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this).F18 = f18
		(_this)._f19 = f19
		(_this).F20 = f20
		(_this)._f21 = f21
		(_this)._f22 = f22
		(_this).F23 = f23
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F3() _dafny.Map {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Array {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.CodePoint {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Array {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.Sequence {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Array {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F19() bool {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F21() _dafny.Sequence {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F22() _dafny.Sequence {
	{
		return _this._f22
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f24 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f24 = false
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

func (_this *C0) Ctor__(f24 bool) {
	{
		(_this)._f24 = f24
	}
}
func (_this *C0) Fm4(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((func() _dafny.Int {
			if true {
				return _dafny.IntOfInt64(845)
			}
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vwwnotdl")).Cardinality())
		})()).Times((_dafny.Zero).Minus(((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_295_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-399), !((_this).F24()))).Cardinality())
		})))).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(!((_this).F24()))).Cardinality()), _dafny.IntOfInt64(980)))).Cardinality()))
	}
}
func (_this *C0) Fm5(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("lffyht")
	}
}
func (_this *C0) F24() bool {
	{
		return _this._f24
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
