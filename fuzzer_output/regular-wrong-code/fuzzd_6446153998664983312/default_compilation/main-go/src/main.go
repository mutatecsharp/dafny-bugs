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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) _dafny.Int {
	return ((_dafny.SetOf(!(false))).Union(_dafny.SetOf(false))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(!(false))).Intersection((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(false))
		}
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(!(true)))
	})())
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) bool {
	return !(true)
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true, false, false)).Difference((_dafny.SetOf(true)).Intersection(_dafny.SetOf(false, !(false))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return ((func() D5 {
		if true {
			return Companion_D5_.Create_DC14_(_dafny.SeqOf((Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(642))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(336))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg0 _dafny.Int) interface{} {
					return coer0(arg0)
				}
			}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			}))).Cardinality()))).Cardinality()))).Dtor_cf17()))
		}
		return Companion_D5_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(675), _dafny.IntOfInt64(827), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(718), _dafny.IntOfInt64(-405))).Cardinality()))
	})()).Dtor_cf18()
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	var _source0 D2 = Companion_D2_.Create_DC8_(true, false)
	_ = _source0
	if _source0.Is_DC7() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ammm"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))
	} else if _source0.Is_DC8() {
		var _2___mcc_h0 bool = _source0.Get_().(D2_DC8).Cf9
		_ = _2___mcc_h0
		var _3___mcc_h1 bool = _source0.Get_().(D2_DC8).Cf10
		_ = _3___mcc_h1
		var _4_cf10 bool = _3___mcc_h1
		_ = _4_cf10
		var _5_cf9 bool = _2___mcc_h0
		_ = _5_cf9
		return _dafny.UnicodeSeqOfUtf8Bytes("widm")
	} else {
		var _6___mcc_h2 _dafny.Set = _source0.Get_().(D2_DC6).Cf8
		_ = _6___mcc_h2
		var _7_cf8 _dafny.Set = _6___mcc_h2
		_ = _7_cf8
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("ult")
		} else {
			return _dafny.UnicodeSeqOfUtf8Bytes("krthaq")
		}
	}
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.MultiSet {
	var _source1 D5 = (func() D5 {
		if !(false) {
			return Companion_D5_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(268)))
		}
		return Companion_D5_.Create_DC14_(_dafny.SeqOf(_dafny.IntOfInt64(-267)))
	})()
	_ = _source1
	if _source1.Is_DC15() {
		var _8___mcc_h0 _dafny.Int = _source1.Get_().(D5_DC15).Cf19
		_ = _8___mcc_h0
		var _9___mcc_h1 bool = _source1.Get_().(D5_DC15).Cf20
		_ = _9___mcc_h1
		var _10_cf20 bool = _9___mcc_h1
		_ = _10_cf20
		var _11_cf19 _dafny.Int = _8___mcc_h0
		_ = _11_cf19
		if _10_cf20 {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_10_cf20, _10_cf20)).Cardinality())))
		} else {
			return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _dafny.UnicodeSeqOfUtf8Bytes("isbarghd"))).Cardinality())
		}
	} else {
		var _12___mcc_h2 _dafny.Sequence = _source1.Get_().(D5_DC14).Cf18
		_ = _12___mcc_h2
		var _13_cf18 _dafny.Sequence = _12___mcc_h2
		_ = _13_cf18
		return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _15_v0 _dafny.CodePoint
			_15_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			})), _15_v0) {
				_coll0.Add(_15_v0)
			}
		}
		return _coll0.ToSet()
	}()).Intersection(_dafny.SetOf(_dafny.CodePoint('f'), _dafny.CodePoint('m'), _dafny.CodePoint('i'), _dafny.CodePoint('f'), _dafny.CodePoint('x')))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false, (_dafny.IntOfInt64(-386)).Cmp(_dafny.IntOfInt64(486)) > 0, !((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('v'))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _16_v0 _dafny.CodePoint
			_16_v0 = interface{}(_compr_1).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('v')), _16_v0) {
				_coll1.Add(_16_v0)
			}
		}
		return _coll1.ToSet()
	}()).IsSubsetOf(_dafny.SetOf(_dafny.CodePoint('s'), _dafny.CodePoint('q')))), !((func() bool {
		if false {
			return false
		}
		return !(!(true))
	})()))
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Map {
	if (_dafny.SetOf(true, !(false))).IsSubsetOf(_dafny.SetOf(true, !(false))) {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), !(false))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, true), true))
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Map, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-464), _dafny.UnicodeSeqOfUtf8Bytes("creunkjsj"))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("ft"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lvvvveiw")).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})))))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var _18_v0 _dafny.Array
	_ = _18_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
	_ = _nw0
	_18_v0 = _nw0
	for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_18_v0), 0))); ; {
		_guard_loop_0, _ok2 := _iter2()
		if !_ok2 {
			break
		}
		var _19_i0 _dafny.Int
		_19_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_19_i0).Sign() != -1) && ((_19_i0).Cmp(_dafny.ArrayLen((_18_v0), 0)) < 0)) {
			(_18_v0).ArraySet1(false, (_19_i0).Int())
		}
	}
	var _20_v1 _dafny.Int
	_ = _20_v1
	_20_v1 = _dafny.IntOfInt64(555)
	var _21_v2 D0
	_ = _21_v2
	_21_v2 = Companion_D0_.Create_DC0_(p0)
	var _22_v3 *C0
	_ = _22_v3
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__(_20_v1, p0, _21_v2)
	_22_v3 = _nw1
	var _23_v4 _dafny.CodePoint
	_ = _23_v4
	_23_v4 = _dafny.CodePoint('p')
	var _24_v5 _dafny.Set
	_ = _24_v5
	_24_v5 = _dafny.SetOf(_23_v4)
	var _25_v6 _dafny.Set
	_ = _25_v6
	_25_v6 = _dafny.SetOf(_24_v5, _24_v5, _24_v5, _24_v5, Companion_Default___.Fm9((_22_v3).F12(), _20_v1, _dafny.IntOfUint32((p1).Cardinality()), globalState))
	var _hi0 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_22_v3, (_25_v6).Cardinality())).Cardinality()
	_ = _hi0
	for _26_i1 := (func() _dafny.Int {
		if p0 {
			return _20_v1
		}
		return _20_v1
	})(); _26_i1.Cmp(_hi0) < 0; _26_i1 = _26_i1.Plus(_dafny.One) {
		var _27_v7 _dafny.Sequence
		_ = _27_v7
		_27_v7 = _dafny.UnicodeSeqOfUtf8Bytes("lotbhuwwl")
		var _28_v8 _dafny.Set
		_ = _28_v8
		_28_v8 = _dafny.SetOf(_dafny.IntOfUint32((_27_v7).Cardinality()))
		var _29_v9 _dafny.Sequence
		_ = _29_v9
		_29_v9 = _dafny.SeqOf(p1)
		var _30_v10 _dafny.Map
		_ = _30_v10
		_30_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_22_v3).F12(), _26_i1)
		var _31_v11 _dafny.MultiSet
		_ = _31_v11
		_31_v11 = _dafny.MultiSetOf(_18_v0)
		var _32_v12 _dafny.MultiSet
		_ = _32_v12
		_32_v12 = _31_v11
		var _33_v13 _dafny.Array
		_ = _33_v13
		var _nwElement0_0 _dafny.Sequence = p1
		_ = _nwElement0_0
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(22))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_0, 0)
		(_nw2).ArraySet1(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_26_i1, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), !(!(p0))), 1)
		(_nw2).ArraySet1(p1, 2)
		(_nw2).ArraySet1(p1, 3)
		(_nw2).ArraySet1(p1, 4)
		(_nw2).ArraySet1(p1, 5)
		(_nw2).ArraySet1((_29_v9).Select((Companion_Default___.SafeIndex(_20_v1, _dafny.IntOfUint32((_29_v9).Cardinality()))).Uint32()).(_dafny.Sequence), 6)
		(_nw2).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(!(true)), (Companion_Default___.SafeIndex((_22_v3).F12(), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))).Uint32(), p0)), 7)
		(_nw2).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 8)
		(_nw2).ArraySet1(Companion_Default___.Fm10(_dafny.IntOfInt64(880), p0, (_22_v3).F12(), _26_i1, globalState), 9)
		(_nw2).ArraySet1(p1, 10)
		(_nw2).ArraySet1(p1, 11)
		(_nw2).ArraySet1(p1, 12)
		(_nw2).ArraySet1((func() _dafny.Sequence {
			if p0 {
				return p1
			}
			return _dafny.SeqOf(p0)
		})(), 13)
		(_nw2).ArraySet1(Companion_Default___.Fm10((_22_v3).F12(), p0, (func() _dafny.Int {
			if (_30_v10).Contains((_32_v12).Cardinality()) {
				return (_30_v10).Get((_32_v12).Cardinality()).(_dafny.Int)
			}
			return _26_i1
		})(), _26_i1, globalState), 14)
		(_nw2).ArraySet1(p1, 15)
		(_nw2).ArraySet1((func() _dafny.Sequence {
			if p0 {
				return p1
			}
			return _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_20_v1, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), p0)
		})(), 16)
		(_nw2).ArraySet1(_dafny.SeqOf(p0, p0, p0, p0), 17)
		(_nw2).ArraySet1((func() _dafny.Sequence {
			if p0 {
				return p1
			}
			return p1
		})(), 18)
		(_nw2).ArraySet1(_dafny.SeqOf(p0, p0, p0, p0, !(p0)), 19)
		(_nw2).ArraySet1(_dafny.SeqOf(Companion_Default___.Fm2(p0, p1, globalState), p0), 20)
		(_nw2).ArraySet1(p1, 21)
		_33_v13 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_33_v13), 0))
		_ = _index0
		(_33_v13).ArraySet1(p1, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_33_v13), 0))
		_ = _index1
		var _rhs0 _dafny.Set = _28_v8
		_ = _rhs0
		var _rhs1 bool = p0
		_ = _rhs1
		var _rhs2 _dafny.Sequence = p1
		_ = _rhs2
		var _rhs3 *C0 = _22_v3
		_ = _rhs3
		var _rhs4 bool = p0
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 _dafny.Array = _33_v13
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_33_v13), 0))
		_ = _lhs2
		var _lhs3 *GlobalState = globalState
		_ = _lhs3
		_28_v8 = _rhs0
		_lhs0.F5 = _rhs1
		(_lhs1).ArraySet1(_rhs2, (_lhs2).Int())
		_22_v3 = _rhs3
		_lhs3.F5 = _rhs4
		var _34_v14 D2
		_ = _34_v14
		_34_v14 = Companion_D2_.Create_DC7_()
		var _35_v15 _dafny.Set
		_ = _35_v15
		_35_v15 = _dafny.SetOf(_34_v14)
		if (((_22_v3).F12()).Times((_35_v15).Cardinality())).Cmp(_26_i1) > 0 {
			var _36_v16 _dafny.Array
			_ = _36_v16
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw3
			_36_v16 = _nw3
			_36_v16 = _36_v16
			var _37_v17 *C0
			_ = _37_v17
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__((_dafny.IntOfUint32((_27_v7).Cardinality())).Times(_26_i1), (func() bool {
				if p0 {
					return p0
				}
				return p0
			})(), _21_v2)
			_37_v17 = _nw4
			var _pat_let_tv0 = p0
			_ = _pat_let_tv0
			var _pat_let_tv1 = p0
			_ = _pat_let_tv1
			var _pat_let_tv2 = p0
			_ = _pat_let_tv2
			var _rhs5 bool = !(!(p0) || (true))
			_ = _rhs5
			var _rhs6 _dafny.Int = (_26_i1).Minus(Companion_Default___.SafeDivisionInt((_37_v17).F12(), (_dafny.Zero).Minus(_26_i1)))
			_ = _rhs6
			var _rhs7 D0 = func(_pat_let0_0 D0) D0 {
				return func(_38_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let1_0 bool) D0 {
						return func(_39_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_39_dt__update_hcf0_h0)
						}(_pat_let1_0)
					}((func() bool {
						if !(_pat_let_tv0) {
							return _pat_let_tv1
						}
						return _pat_let_tv2
					})())
				}(_pat_let0_0)
			}(_21_v2)
			_ = _rhs7
			var _rhs8 bool = !(!(p0) || (p0))
			_ = _rhs8
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			_lhs4.F8 = _rhs5
			_20_v1 = _rhs6
			_21_v2 = _rhs7
			_lhs5.F8 = _rhs8
			var _40_v18 _dafny.Sequence
			_ = _40_v18
			_40_v18 = _dafny.SeqOf((_22_v3).F12())
			var _41_v19 _dafny.Map
			_ = _41_v19
			_41_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v18, _37_v17)
			var _42_v20 _dafny.Map
			_ = _42_v20
			_42_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v18, _37_v17)).Merge(_41_v19), (_30_v10).Merge(_30_v10))
			_42_v20 = (_42_v20).Update((_41_v19).Update(_40_v18, _37_v17), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(485), (_22_v3).F12()))
			var _43_v21 _dafny.Sequence
			_ = _43_v21
			_43_v21 = _dafny.SeqOf(_37_v17, _22_v3, _37_v17, _22_v3)
			_43_v21 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_37_v17, _22_v3), _43_v21), _43_v21)
		} else {
			var _44_v22 _dafny.Array
			_ = _44_v22
			var _nwElement0_1 _dafny.Int = _dafny.IntOfInt64(-460)
			_ = _nwElement0_1
			var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(4))
			_ = _nw5
			(_nw5).ArraySet1(_nwElement0_1, 0)
			(_nw5).ArraySet1((_22_v3).F12(), 1)
			(_nw5).ArraySet1(Companion_Default___.SafeModuloInt((_22_v3).F12(), (_22_v3).F12()), 2)
			(_nw5).ArraySet1((_22_v3).F12(), 3)
			_44_v22 = _nw5
			_44_v22 = _44_v22
			var _45_v23 _dafny.MultiSet
			_ = _45_v23
			_45_v23 = _dafny.MultiSetOf(Companion_Default___.Fm0((_dafny.Zero).Minus(_20_v1), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0), globalState))
			var _46_v24 D0
			_ = _46_v24
			_46_v24 = Companion_D0_.Create_DC2_(p0, (_45_v23).Cardinality())
			var _47_v25 _dafny.Map
			_ = _47_v25
			_47_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_33_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_33_v13), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.IntOfUint32(((_33_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_33_v13), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), p0), p0)
			var _48_v26 _dafny.Sequence
			_ = _48_v26
			_48_v26 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0), _47_v25, _47_v25)
			var _rhs9 _dafny.Int = (_dafny.Zero).Minus((_22_v3).F12())
			_ = _rhs9
			var _rhs10 _dafny.Int = (func() _dafny.Int {
				if (_45_v23).Contains((_22_v3).F12()) {
					return (_45_v23).Multiplicity((_22_v3).F12())
				}
				return Companion_Default___.Fm0((_22_v3).F12(), (_48_v26).Select((Companion_Default___.SafeIndex(_26_i1, _dafny.IntOfUint32((_48_v26).Cardinality()))).Uint32()).(_dafny.Map), globalState)
			})()
			_ = _rhs10
			var _rhs11 _dafny.CodePoint = _23_v4
			_ = _rhs11
			var _rhs12 D0 = _46_v24
			_ = _rhs12
			_20_v1 = _rhs9
			_20_v1 = _rhs10
			_23_v4 = _rhs11
			_46_v24 = _rhs12
			(globalState).F5 = true
			var _49_v27 _dafny.Sequence
			_ = _49_v27
			_49_v27 = _dafny.SeqOf(Companion_Default___.Fm0(_dafny.IntOfUint32((_27_v7).Cardinality()), Companion_Default___.Fm11(globalState), globalState), _dafny.IntOfUint32(((_33_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_33_v13), 0))).Int()).(_dafny.Sequence)).Cardinality()))
			var _50_v28 _dafny.Map
			_ = _50_v28
			_50_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v4, (_22_v3).F12())
			(globalState).F5 = ((Companion_Default___.Fm12(_30_v10, _28_v8, (_50_v28).Cardinality(), p0, globalState)).Update((_22_v3).F12(), _27_v7)).Contains((_49_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.IntOfUint32((_49_v27).Cardinality()))).Uint32()).(_dafny.Int))
			var _51_v29 T0
			_ = _51_v29
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__(_20_v1, p0, _21_v2)
			_51_v29 = _nw6
			var _52_v30 _dafny.Map
			_ = _52_v30
			_52_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), _51_v29)
			var _53_v31 _dafny.Set
			_ = _53_v31
			_53_v31 = _dafny.SetOf(_52_v30)
			_53_v31 = _53_v31
		}
		var _54_v32 _dafny.Set
		_ = _54_v32
		_54_v32 = _dafny.SetOf(p0)
		var _55_v33 _dafny.Map
		_ = _55_v33
		_55_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(459), _54_v32)
		var _56_v34 _dafny.MultiSet
		_ = _56_v34
		_56_v34 = _dafny.MultiSetOf(((_55_v33).Merge(_55_v33)).Cardinality(), ((_22_v3).F12()).Plus(_20_v1), _20_v1, _dafny.IntOfInt64(865), (_26_i1).Plus((_22_v3).F12()))
		_56_v34 = _56_v34
		_22_v3 = _22_v3
	}
	var _57_v35 _dafny.Map
	_ = _57_v35
	_57_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.MultiSetOf(p0))
	var _58_v36 _dafny.Map
	_ = _58_v36
	_58_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
	(globalState).F8 = (_22_v3).Fm7(_21_v2, _57_v35, (func() _dafny.Int {
		if (func() bool {
			if (_58_v36).Contains(p0) {
				return (_58_v36).Get(p0).(bool)
			}
			return p0
		})() {
			return _20_v1
		}
		return _dafny.IntOfInt64(609)
	})(), globalState)
	var _59_v37 T0
	_ = _59_v37
	var _nw7 *C0 = New_C0_()
	_ = _nw7
	_nw7.Ctor__((_dafny.Zero).Minus(_dafny.IntOfInt64(-566)), p0, _21_v2)
	_59_v37 = _nw7
	var _60_v38 _dafny.Map
	_ = _60_v38
	_60_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_22_v3).F12(), _59_v37)
	var _61_v39 _dafny.MultiSet
	_ = _61_v39
	_61_v39 = _dafny.MultiSetOf(_59_v37, (func() T0 {
		if (_60_v38).Contains((_22_v3).F12()) {
			return (_60_v38).Get((_22_v3).F12()).(T0)
		}
		return _59_v37
	})())
	var _62_v40 _dafny.Map
	_ = _62_v40
	_62_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v39, !((_20_v1).Cmp(_dafny.IntOfInt64(922)) != 0))
	(globalState).F8 = (func() bool {
		if (_62_v40).Contains(_61_v39) {
			return (_62_v40).Get(_61_v39).(bool)
		}
		return (_59_v37).F10()
	})()
	var _63_i2 _dafny.Int
	_ = _63_i2
	_63_i2 = _dafny.Zero
	{
		for (Companion_D2_.Create_DC8_((_59_v37).F10(), p0)).Dtor_cf10() {
			{
				if (_63_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_63_i2 = (_63_i2).Plus(_dafny.One)
				var _rhs13 _dafny.Int = _dafny.IntOfInt64(107)
				_ = _rhs13
				var _rhs14 _dafny.Int = _20_v1
				_ = _rhs14
				var _rhs15 bool = false
				_ = _rhs15
				var _lhs6 *GlobalState = globalState
				_ = _lhs6
				_20_v1 = _rhs13
				_20_v1 = _rhs14
				_lhs6.F8 = _rhs15
				if ((_22_v3).F12()).Cmp((_22_v3).F12()) == 0 {
					_20_v1 = (_dafny.Zero).Minus(_20_v1)
					var _64_v41 _dafny.Sequence
					_ = _64_v41
					_64_v41 = _dafny.SeqOf((_22_v3).F12(), _20_v1)
					var _65_v42 D2
					_ = _65_v42
					_65_v42 = Companion_D2_.Create_DC8_((_59_v37).F10(), (_59_v37).F10())
					var _66_v43 _dafny.MultiSet
					_ = _66_v43
					_66_v43 = _dafny.MultiSetOf((_59_v37).F10())
					var _67_v44 _dafny.Sequence
					_ = _67_v44
					_67_v44 = _dafny.SeqOf(_66_v43)
					var _68_v45 _dafny.Array
					_ = _68_v45
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_0
					var _nw8 _dafny.Array
					_ = _nw8
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw8 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) _dafny.Int = (func(_69_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_70_i3 _dafny.Int) _dafny.Int {
								return (_70_i3).Times(_69_v1)
							}
						})(_20_v1)
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw8 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw8).ArraySet1(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw8).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_68_v45 = _nw8
					var _71_v46 _dafny.Map
					_ = _71_v46
					_71_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Select((Companion_Default___.SafeIndex((_22_v3).F12(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), _68_v45)
					var _rhs16 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_64_v41, (func() _dafny.Sequence {
						if (_59_v37).F10() {
							return _64_v41
						}
						return _dafny.SeqOf((_22_v3).F12())
					})())
					_ = _rhs16
					var _rhs17 bool = !(!_dafny.Companion_Sequence_.Contains(_67_v44, _66_v43)) || (p0)
					_ = _rhs17
					var _rhs18 _dafny.Int = (_dafny.IntOfInt64(854)).Times(((_22_v3).F12()).Times((_71_v46).Cardinality()))
					_ = _rhs18
					var _rhs19 D2 = _65_v42
					_ = _rhs19
					var _rhs20 T0 = (func() T0 {
						if (_60_v38).Contains((_22_v3).F12()) {
							return (_60_v38).Get((_22_v3).F12()).(T0)
						}
						return _59_v37
					})()
					_ = _rhs20
					var _lhs7 *GlobalState = globalState
					_ = _lhs7
					_64_v41 = _rhs16
					_lhs7.F8 = _rhs17
					_20_v1 = _rhs18
					_65_v42 = _rhs19
					_59_v37 = _rhs20
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_68_v45), 0))
					_ = _index2
					(_68_v45).ArraySet1((_22_v3).F12(), (_index2).Int())
					var _72_v47 _dafny.MultiSet
					_ = _72_v47
					_72_v47 = _dafny.MultiSetOf(_18_v0)
					var _73_v48 _dafny.Set
					_ = _73_v48
					_73_v48 = _dafny.SetOf(_64_v41)
					var _74_v49 _dafny.Sequence
					_ = _74_v49
					_74_v49 = _dafny.SeqOf(_22_v3)
					var _75_v50 _dafny.Map
					_ = _75_v50
					_75_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_73_v48).IsSubsetOf(_73_v48), _dafny.Companion_Sequence_.Concatenate(_74_v49, _74_v49))
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_68_v45), 0))
					_ = _index3
					var _rhs21 _dafny.Int = (((_72_v47).Union(_72_v47)).Cardinality()).Times((_22_v3).F12())
					_ = _rhs21
					var _rhs22 _dafny.Int = (_22_v3).F12()
					_ = _rhs22
					var _rhs23 _dafny.Int = (_22_v3).F12()
					_ = _rhs23
					var _rhs24 _dafny.Int = (_75_v50).Cardinality()
					_ = _rhs24
					var _rhs25 bool = ((_59_v37).F10()) || ((_59_v37).F10())
					_ = _rhs25
					var _lhs8 _dafny.Array = _68_v45
					_ = _lhs8
					var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_68_v45), 0))
					_ = _lhs9
					var _lhs10 *GlobalState = globalState
					_ = _lhs10
					(_lhs8).ArraySet1(_rhs21, (_lhs9).Int())
					_20_v1 = _rhs22
					_20_v1 = _rhs23
					_20_v1 = _rhs24
					_lhs10.F8 = _rhs25
					_59_v37 = _59_v37
					var _76_v51 _dafny.Array
					_ = _76_v51
					var _len0_1 _dafny.Int = _dafny.One
					_ = _len0_1
					var _nw9 _dafny.Array
					_ = _nw9
					if _len0_1.Cmp(_dafny.Zero) == 0 {
						_nw9 = _dafny.NewArray(_len0_1)
					} else {
						var _init1 func(_dafny.Int) D1 = (func(_77_v37 T0, _78_p0 bool, _79_v3 *C0) func(_dafny.Int) D1 {
							return func(_80_i4 _dafny.Int) D1 {
								return Companion_D1_.Create_DC4_((_77_v37).F10(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_p0, (_79_v3).F12()), !((_77_v37).F10()))
							}
						})(_59_v37, p0, _22_v3)
						_ = _init1
						var _element0_1 = _init1(_dafny.Zero)
						_ = _element0_1
						_nw9 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
						(_nw9).ArraySet1(_element0_1, 0)
						var _nativeLen0_1 = (_len0_1).Int()
						_ = _nativeLen0_1
						for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
							(_nw9).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
						}
					}
					_76_v51 = _nw9
					var _81_v52 _dafny.Map
					_ = _81_v52
					_81_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_59_v37).F10(), (_22_v3).F12())
					var _82_v53 D1
					_ = _82_v53
					_82_v53 = Companion_D1_.Create_DC4_((_59_v37).F10(), _81_v52, true)
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_76_v51), 0))
					_ = _index4
					(_76_v51).ArraySet1(_82_v53, (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_76_v51), 0))
					_ = _index5
					(_76_v51).ArraySet1(_82_v53, (_index5).Int())
				} else {
					var _83_v54 D0
					_ = _83_v54
					_83_v54 = Companion_D0_.Create_DC1_()
					_83_v54 = _83_v54
					var _84_v55 D2
					_ = _84_v55
					_84_v55 = Companion_D2_.Create_DC8_((_59_v37).F10(), (_59_v37).F10())
					var _85_v56 _dafny.MultiSet
					_ = _85_v56
					_85_v56 = _dafny.MultiSetOf((_84_v55).Dtor_cf9(), (_59_v37).F10(), (_59_v37).F10())
					(globalState).F8 = (func() bool {
						if (_59_v37).F10() {
							return (_59_v37).F10()
						}
						return (_22_v3).Fm7(Companion_D0_.Create_DC0_(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_59_v37).F10(), _85_v56), (_22_v3).F12(), globalState)
					})()
					_59_v37 = _59_v37
					_59_v37 = _59_v37
					var _86_v57 _dafny.Array
					_ = _86_v57
					var _nw10 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
					_ = _nw10
					_86_v57 = _nw10
					var _87_v58 _dafny.Map
					_ = _87_v58
					_87_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_22_v3).F12(), _22_v3)
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_86_v57), 0))
					_ = _index6
					(_86_v57).ArraySet1((func() *C0 {
						if (_59_v37).F10() {
							return _22_v3
						}
						return (func() *C0 {
							if (_87_v58).Contains((_22_v3).F12()) {
								return (_87_v58).Get((_22_v3).F12()).(*C0)
							}
							return _22_v3
						})()
					})(), (_index6).Int())
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_86_v57), 0))
					_ = _index7
					(_86_v57).ArraySet1(_22_v3, (_index7).Int())
				}
				var _88_v59 D0
				_ = _88_v59
				_88_v59 = Companion_D0_.Create_DC2_((_59_v37).F10(), (_22_v3).F12())
				var _89_v60 _dafny.Map
				_ = _89_v60
				_89_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_59_v37).F10())
				var _90_v61 D4
				_ = _90_v61
				_90_v61 = Companion_D4_.Create_DC12_((_59_v37).F10(), _59_v37, true, (_59_v37).F10())
				var _91_v62 _dafny.Map
				_ = _91_v62
				_91_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_22_v3, _90_v61)
				var _92_v63 _dafny.Sequence
				_ = _92_v63
				_92_v63 = _dafny.SeqOf(_dafny.IntOfInt64(-463))
				var _93_v64 _dafny.Array
				_ = _93_v64
				var _nwElement0_2 _dafny.Int = _dafny.IntOfInt64(-181)
				_ = _nwElement0_2
				var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(19))
				_ = _nw11
				(_nw11).ArraySet1(_nwElement0_2, 0)
				(_nw11).ArraySet1((_22_v3).F12(), 1)
				(_nw11).ArraySet1(_dafny.IntOfInt64(658), 2)
				(_nw11).ArraySet1((_88_v59).Dtor_cf2(), 3)
				(_nw11).ArraySet1(_20_v1, 4)
				(_nw11).ArraySet1((_22_v3).F12(), 5)
				(_nw11).ArraySet1(Companion_Default___.Fm0((_22_v3).F12(), _89_v60, globalState), 6)
				(_nw11).ArraySet1((_22_v3).F12(), 7)
				(_nw11).ArraySet1((_91_v62).Cardinality(), 8)
				(_nw11).ArraySet1(_dafny.IntOfInt64(-413), 9)
				(_nw11).ArraySet1(_20_v1, 10)
				(_nw11).ArraySet1(_20_v1, 11)
				(_nw11).ArraySet1(_20_v1, 12)
				(_nw11).ArraySet1((_22_v3).F12(), 13)
				(_nw11).ArraySet1(_dafny.IntOfInt64(528), 14)
				(_nw11).ArraySet1(_20_v1, 15)
				(_nw11).ArraySet1((_22_v3).F12(), 16)
				(_nw11).ArraySet1(_20_v1, 17)
				(_nw11).ArraySet1(_dafny.IntOfUint32((_92_v63).Cardinality()), 18)
				_93_v64 = _nw11
				var _94_v65 D4
				_ = _94_v65
				_94_v65 = Companion_D4_.Create_DC10_(_93_v64)
				var _95_v66 _dafny.Map
				_ = _95_v66
				_95_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(17), (_94_v65).Dtor_cf12())
				_95_v66 = (_95_v66).Update((func() _dafny.Int {
					if p0 {
						return _20_v1
					}
					return _dafny.IntOfInt64(-648)
				})(), _93_v64)
				(globalState).F8 = ((_59_v37).F10()) == ((Companion_Default___.Fm2((_59_v37).F10(), p1, globalState)) && ((_59_v37).F10()))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _nw12 *C0 = New_C0_()
	_ = _nw12
	_nw12.Ctor__((_22_v3).F12(), p0, (_59_v37).F11())
	_59_v37 = _nw12
	var _96_v67 _dafny.Sequence
	_ = _96_v67
	_96_v67 = _dafny.UnicodeSeqOfUtf8Bytes("qf")
	r0 = _dafny.Companion_Sequence_.Concatenate(_96_v67, _96_v67)
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _97_v0 bool
	_ = _97_v0
	_97_v0 = false
	var _98_v1 _dafny.Sequence
	_ = _98_v1
	_98_v1 = _dafny.SeqOf(_97_v0, _97_v0)
	var _99_v2 _dafny.Int
	_ = _99_v2
	_99_v2 = _dafny.IntOfInt64(147)
	var _100_v3 _dafny.MultiSet
	_ = _100_v3
	_100_v3 = _dafny.MultiSetOf(_dafny.SeqOf(true), _98_v1, _dafny.Companion_Sequence_.Update(_98_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-758), _dafny.IntOfUint32((_98_v1).Cardinality()))).Uint32(), false), _dafny.Companion_Sequence_.Update(_98_v1, (Companion_Default___.SafeIndex(_99_v2, _dafny.IntOfUint32((_98_v1).Cardinality()))).Uint32(), !(_97_v0)))
	var _101_v4 D0
	_ = _101_v4
	_101_v4 = Companion_D0_.Create_DC0_(_97_v0)
	var _102_v5 _dafny.Array
	_ = _102_v5
	var _nwElement0_3 bool = _97_v0
	_ = _nwElement0_3
	var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(11))
	_ = _nw13
	(_nw13).ArraySet1(_nwElement0_3, 0)
	(_nw13).ArraySet1((_101_v4).Dtor_cf0(), 1)
	(_nw13).ArraySet1(_97_v0, 2)
	(_nw13).ArraySet1(_97_v0, 3)
	(_nw13).ArraySet1(_97_v0, 4)
	(_nw13).ArraySet1(_97_v0, 5)
	(_nw13).ArraySet1(_97_v0, 6)
	(_nw13).ArraySet1(_97_v0, 7)
	(_nw13).ArraySet1(_97_v0, 8)
	(_nw13).ArraySet1(_97_v0, 9)
	(_nw13).ArraySet1(_97_v0, 10)
	_102_v5 = _nw13
	var _103_v6 _dafny.Map
	_ = _103_v6
	_103_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v0, _99_v2)
	var _104_v7 D1
	_ = _104_v7
	_104_v7 = Companion_D1_.Create_DC4_(_97_v0, _103_v6, _97_v0)
	var _105_v8 _dafny.Map
	_ = _105_v8
	_105_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(_97_v0))).Cardinality()), _97_v0)
	var _106_globalState *GlobalState
	_ = _106_globalState
	var _nw14 *GlobalState = New_GlobalState_()
	_ = _nw14
	_nw14.Ctor__(_100_v3, _102_v5, _dafny.IntOfInt64(-299), true, _dafny.IntOfInt64(-378), true, (_104_v7).Dtor_cf5(), _105_v8, false, true)
	_106_globalState = _nw14
	var _107_v9 _dafny.MultiSet
	_ = _107_v9
	_107_v9 = _dafny.MultiSetOf(!(_97_v0), _97_v0, _97_v0)
	if !(((_107_v9).Intersection(_107_v9)).IsSubsetOf(_107_v9)) {
		var _108_v10 _dafny.Map
		_ = _108_v10
		_108_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v1, _97_v0)
		var _109_v11 _dafny.Sequence
		_ = _109_v11
		_109_v11 = _dafny.UnicodeSeqOfUtf8Bytes("dsjek")
		_99_v2 = (Companion_Default___.Fm0(_99_v2, _108_v10, _106_globalState)).Plus((_dafny.IntOfUint32((_109_v11).Cardinality())).Plus(_99_v2))
		var _110_v12 _dafny.Array
		_ = _110_v12
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_2
		var _nw15 _dafny.Array
		_ = _nw15
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw15 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_111_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_112_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_112_i0, _111_v2)
				}
			})(_99_v2)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw15 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw15).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw15).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_110_v12 = _nw15
		_110_v12 = _110_v12
		var _113_v13 _dafny.Sequence
		_ = _113_v13
		var _out0 _dafny.Sequence
		_ = _out0
		_out0 = Companion_Default___.M0(_97_v0, _98_v1, _106_globalState)
		_113_v13 = _out0
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_110_v12), 0))
		_ = _index8
		(_110_v12).ArraySet1(_99_v2, (_index8).Int())
		var _114_v14 _dafny.Set
		_ = _114_v14
		_114_v14 = _dafny.SetOf(_97_v0, _97_v0, false, _97_v0, _97_v0)
		var _115_v15 _dafny.Set
		_ = _115_v15
		_115_v15 = _dafny.SetOf((_114_v14).Cardinality())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_110_v12), 0))
		_ = _index9
		(_110_v12).ArraySet1(((_115_v15).Cardinality()).Plus(_99_v2), (_index9).Int())
		var _116_v16 _dafny.Map
		_ = _116_v16
		_116_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v0, _97_v0)
		var _117_v17 _dafny.Sequence
		_ = _117_v17
		_117_v17 = _dafny.SeqOf((func() _dafny.Int {
			if !((func() bool {
				if (_116_v16).Contains(true) {
					return (_116_v16).Get(true).(bool)
				}
				return _97_v0
			})()) {
				return _dafny.IntOfInt64(446)
			}
			return (_110_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_110_v12), 0))).Int()).(_dafny.Int)
		})())
		var _118_v18 _dafny.Map
		_ = _118_v18
		_118_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v0, _117_v17)
		_117_v17 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_117_v17, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_118_v18).Contains(_97_v0) {
				return (_118_v18).Get(_97_v0).(_dafny.Sequence)
			}
			return _117_v17
		})(), (Companion_Default___.SafeIndex(_99_v2, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_118_v18).Contains(_97_v0) {
				return (_118_v18).Get(_97_v0).(_dafny.Sequence)
			}
			return _117_v17
		})()).Cardinality()))).Uint32(), (_110_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_110_v12), 0))).Int()).(_dafny.Int))), _dafny.SeqOf((_110_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_110_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(584)))
	} else {
		_107_v9 = (_107_v9).Union(Companion_Default___.Fm1(_99_v2, _97_v0, _99_v2, _97_v0, _106_globalState))
		(_106_globalState).F8 = (Companion_D1_.Create_DC4_(_97_v0, _103_v6, _97_v0)).Dtor_cf6()
		var _119_v19 D1
		_ = _119_v19
		_119_v19 = Companion_D1_.Create_DC3_(_103_v6)
		var _120_v20 D1
		_ = _120_v20
		_120_v20 = Companion_D1_.Create_DC5_(_119_v19)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))
		_ = _index10
		(_102_v5).ArraySet1(false, (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))
		_ = _index11
		var _rhs26 _dafny.Int = _99_v2
		_ = _rhs26
		var _rhs27 _dafny.Map = _105_v8
		_ = _rhs27
		var _rhs28 D1 = _120_v20
		_ = _rhs28
		var _rhs29 _dafny.Int = _99_v2
		_ = _rhs29
		var _rhs30 bool = _97_v0
		_ = _rhs30
		var _lhs11 _dafny.Array = _102_v5
		_ = _lhs11
		var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))
		_ = _lhs12
		_99_v2 = _rhs26
		_105_v8 = _rhs27
		_120_v20 = _rhs28
		_99_v2 = _rhs29
		(_lhs11).ArraySet1(_rhs30, (_lhs12).Int())
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))
		_ = _index12
		(_102_v5).ArraySet1(Companion_Default___.Fm2((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate(_98_v1, _98_v1), _106_globalState), (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))
		_ = _index13
		var _rhs31 bool = !(_97_v0) || (_97_v0)
		_ = _rhs31
		var _rhs32 bool = (_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool)
		_ = _rhs32
		var _lhs13 *GlobalState = _106_globalState
		_ = _lhs13
		var _lhs14 _dafny.Array = _102_v5
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_102_v5), 0))
		_ = _lhs15
		_lhs13.F8 = _rhs31
		(_lhs14).ArraySet1(_rhs32, (_lhs15).Int())
	}
	(_106_globalState).F8 = _97_v0
	_99_v2 = _99_v2
	var _hi1 _dafny.Int = _99_v2
	_ = _hi1
	for _121_i1 := _99_v2; _121_i1.Cmp(_hi1) < 0; _121_i1 = _121_i1.Plus(_dafny.One) {
		var _122_v21 _dafny.Set
		_ = _122_v21
		_122_v21 = _dafny.SetOf((_dafny.Zero).Minus(_99_v2), _121_i1, _121_i1)
		var _123_v22 _dafny.Sequence
		_ = _123_v22
		_123_v22 = _dafny.UnicodeSeqOfUtf8Bytes("wgy")
		var _124_v23 D0
		_ = _124_v23
		_124_v23 = Companion_D0_.Create_DC1_()
		var _125_v24 _dafny.Map
		_ = _125_v24
		_125_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_123_v22, _124_v23)
		_99_v2 = (func() _dafny.Int {
			if !(_107_v9).Equals(_107_v9) {
				return (_122_v21).Cardinality()
			}
			return (_121_i1).Times((_125_v24).Cardinality())
		})()
		var _126_v25 _dafny.Map
		_ = _126_v25
		_126_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_i1, _99_v2)
		_107_v9 = (_dafny.MultiSetOf(_97_v0, _97_v0, _97_v0, _97_v0)).Update(((func() _dafny.Int {
			if (_126_v25).Contains(_99_v2) {
				return (_126_v25).Get(_99_v2).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_123_v22).Cardinality())
		})()).Cmp(_121_i1) <= 0, Companion_Default___.Abs(_121_i1))
		(_106_globalState).F8 = _97_v0
		if !(!((_97_v0) == (!(_97_v0)))) || (_97_v0) {
			var _127_v26 _dafny.CodePoint
			_ = _127_v26
			_127_v26 = _dafny.CodePoint('f')
			var _128_v27 _dafny.Array
			_ = _128_v27
			var _nwElement0_4 _dafny.CodePoint = _127_v26
			_ = _nwElement0_4
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(3))
			_ = _nw16
			(_nw16).ArraySet1CodePoint(_nwElement0_4, 0)
			(_nw16).ArraySet1CodePoint((func() _dafny.CodePoint {
				if !(_97_v0) {
					return _dafny.CodePoint('y')
				}
				return _127_v26
			})(), 1)
			(_nw16).ArraySet1CodePoint(_127_v26, 2)
			_128_v27 = _nw16
			_128_v27 = _128_v27
			(_106_globalState).F8 = _97_v0
			_98_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_98_v1, _98_v1), _98_v1)
			var _129_v28 _dafny.Sequence
			_ = _129_v28
			var _out1 _dafny.Sequence
			_ = _out1
			_out1 = Companion_Default___.M0(_97_v0, _dafny.Companion_Sequence_.Concatenate(_98_v1, _dafny.SeqOf(!(_97_v0))), _106_globalState)
			_129_v28 = _out1
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_102_v5), 0))
			_ = _index14
			(_102_v5).ArraySet1(true, (_index14).Int())
			var _130_v29 _dafny.Array
			_ = _130_v29
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_3
			var _nw17 _dafny.Array
			_ = _nw17
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw17 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_131_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_132_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_132_i2, _131_v2)
					}
				})(_99_v2)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw17 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw17).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw17).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_130_v29 = _nw17
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_130_v29), 0))
			_ = _index15
			(_130_v29).ArraySet1((_dafny.Zero).Minus(_99_v2), (_index15).Int())
			var _133_v30 D2
			_ = _133_v30
			_133_v30 = Companion_D2_.Create_DC6_(Companion_Default___.Fm3(_106_globalState))
			var _134_v31 _dafny.Set
			_ = _134_v31
			_134_v31 = _dafny.SetOf(_97_v0)
			var _135_v32 _dafny.Sequence
			_ = _135_v32
			_135_v32 = _dafny.SeqOf(_99_v2, _dafny.IntOfInt64(-223))
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_102_v5), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_130_v29), 0))
			_ = _index17
			var _rhs33 bool = (((_133_v30).Dtor_cf8()).Difference(_134_v31)).IsDisjointFrom((func() _dafny.Set {
				if _97_v0 {
					return _dafny.SetOf(_97_v0, _97_v0)
				}
				return _134_v31
			})())
			_ = _rhs33
			var _rhs34 _dafny.Int = Companion_Default___.SafeDivisionInt(_99_v2, _99_v2)
			_ = _rhs34
			var _rhs35 _dafny.Int = (_135_v32).Select((Companion_Default___.SafeIndex((_121_i1).Times(_121_i1), _dafny.IntOfUint32((_135_v32).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs35
			var _rhs36 _dafny.Int = (_dafny.Zero).Minus((_121_i1).Plus(_dafny.IntOfInt64(325)))
			_ = _rhs36
			var _lhs16 _dafny.Array = _102_v5
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_102_v5), 0))
			_ = _lhs17
			var _lhs18 _dafny.Array = _130_v29
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_130_v29), 0))
			_ = _lhs19
			(_lhs16).ArraySet1(_rhs33, (_lhs17).Int())
			_99_v2 = _rhs34
			_99_v2 = _rhs35
			(_lhs18).ArraySet1(_rhs36, (_lhs19).Int())
		} else {
			_123_v22 = _123_v22
			_99_v2 = _121_i1
			var _136_v33 _dafny.Map
			_ = _136_v33
			_136_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _97_v0)
			(_106_globalState).F5 = Companion_Default___.Fm2((Companion_Default___.Fm0(_121_i1, _136_v33, _106_globalState)).Cmp(_99_v2) != 0, _98_v1, _106_globalState)
			_101_v4 = _101_v4
			var _137_v34 _dafny.CodePoint
			_ = _137_v34
			_137_v34 = _dafny.CodePoint('i')
			_137_v34 = _137_v34
		}
	}
	var _138_v36 _dafny.Map
	_ = _138_v36
	_138_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_99_v2, func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v1, _97_v0)).Keys().Elements()); ; {
			_compr_2, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _139_v35 _dafny.Sequence
			_139_v35 = interface{}(_compr_2).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v1, _97_v0)).Contains(_139_v35) {
				_coll2.Add(_139_v35, _97_v0)
			}
		}
		return _coll2.ToMap()
	}(), _106_globalState), _99_v2)
	var _140_v37 _dafny.MultiSet
	_ = _140_v37
	_140_v37 = _dafny.MultiSetOf(_dafny.IntOfInt64(717))
	var _141_v38 _dafny.Sequence
	_ = _141_v38
	_141_v38 = _dafny.UnicodeSeqOfUtf8Bytes("toc")
	var _142_v39 _dafny.CodePoint
	_ = _142_v39
	_142_v39 = _dafny.CodePoint('y')
	var _143_v40 _dafny.Array
	_ = _143_v40
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(11)
	_ = _len0_4
	var _nw18 _dafny.Array
	_ = _nw18
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw18 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Int = (func(_144_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_145_i3 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_145_i3, _144_v2)
			}
		})(_99_v2)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw18 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw18).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw18).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_143_v40 = _nw18
	var _146_v41 _dafny.Sequence
	_ = _146_v41
	_146_v41 = _dafny.SeqOf(_143_v40)
	var _rhs37 bool = (_138_v36).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(249), (func() _dafny.Int {
		if (_140_v37).Contains(_99_v2) {
			return (_140_v37).Multiplicity(_99_v2)
		}
		return _dafny.IntOfUint32((_141_v38).Cardinality())
	})())).Merge(_138_v36))
	_ = _rhs37
	var _rhs38 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm4(_142_v39, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_142_v39, _106_globalState), _141_v38), _97_v0, _106_globalState)).Cardinality())
	_ = _rhs38
	var _rhs39 bool = _97_v0
	_ = _rhs39
	var _rhs40 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_143_v40), _146_v41), _dafny.Companion_Sequence_.Concatenate(_146_v41, _dafny.SeqOf(_143_v40, _143_v40, _143_v40, _143_v40)))).Cardinality()))
	_ = _rhs40
	var _lhs20 *GlobalState = _106_globalState
	_ = _lhs20
	_97_v0 = _rhs37
	_99_v2 = _rhs38
	_lhs20.F8 = _rhs39
	_99_v2 = _rhs40
	for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_102_v5), 0))); ; {
		_guard_loop_1, _ok4 := _iter4()
		if !_ok4 {
			break
		}
		var _147_i4 _dafny.Int
		_147_i4 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_147_i4).Sign() != -1) && ((_147_i4).Cmp(_dafny.ArrayLen((_102_v5), 0)) < 0)) {
			(_102_v5).ArraySet1((_99_v2).Cmp((_99_v2).Times(_99_v2)) > 0, (_147_i4).Int())
		}
	}
	_141_v38 = _141_v38
	_99_v2 = _99_v2
	var _148_v43 _dafny.Map
	_ = _148_v43
	_148_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v2, _99_v2)).Cardinality(), func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(728), _dafny.IntOfInt64(918))); ; {
			_compr_3, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _149_v42 _dafny.Int
			_149_v42 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(728)).Cmp(_149_v42) <= 0) && ((_149_v42).Cmp(_dafny.IntOfInt64(918)) < 0) {
				_coll3.Add(Companion_Default___.SafeDivisionInt(_149_v42, _99_v2), _97_v0)
			}
		}
		return _coll3.ToMap()
	}())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))
	_ = _index18
	(_102_v5).ArraySet1(!(_148_v43).Contains(_dafny.IntOfInt64(927)), (_index18).Int())
	var _150_v45 _dafny.Set
	_ = _150_v45
	_150_v45 = _dafny.SetOf(_98_v1)
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))
	_ = _index19
	var _rhs41 bool = _97_v0
	_ = _rhs41
	var _rhs42 _dafny.CodePoint = (func() _dafny.CodePoint {
		if true {
			return (func() _dafny.CodePoint {
				if _97_v0 {
					return _142_v39
				}
				return _142_v39
			})()
		}
		return _142_v39
	})()
	_ = _rhs42
	var _rhs43 _dafny.Int = Companion_Default___.SafeModuloInt(_99_v2, Companion_Default___.SafeModuloInt(_99_v2, Companion_Default___.Fm0(_99_v2, func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter6 := _dafny.Iterate((_150_v45).Elements()); ; {
			_compr_4, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _151_v44 _dafny.Sequence
			_151_v44 = interface{}(_compr_4).(_dafny.Sequence)
			if (_150_v45).Contains(_151_v44) {
				_coll4.Add(_151_v44, _97_v0)
			}
		}
		return _coll4.ToMap()
	}(), _106_globalState)))
	_ = _rhs43
	var _rhs44 _dafny.Int = (_dafny.IntOfInt64(601)).Plus(_99_v2)
	_ = _rhs44
	var _lhs21 _dafny.Array = _102_v5
	_ = _lhs21
	var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))
	_ = _lhs22
	(_lhs21).ArraySet1(_rhs41, (_lhs22).Int())
	_142_v39 = _rhs42
	_99_v2 = _rhs43
	_99_v2 = _rhs44
	var _152_v46 _dafny.Map
	_ = _152_v46
	_152_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool), (_97_v0) == ((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool)))
	var _153_v48 _dafny.Sequence
	_ = _153_v48
	_153_v48 = _dafny.SeqOf(_99_v2)
	_152_v46 = (_152_v46).Update((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter7 := _dafny.Iterate((_153_v48).Elements()); ; {
			_compr_5, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _154_v47 _dafny.Int
			_154_v47 = interface{}(_compr_5).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_153_v48, _154_v47) {
				_coll5.Add((_154_v47).Minus(_99_v2), _140_v37)
			}
		}
		return _coll5.ToMap()
	}()).Contains((_140_v37).Cardinality()), (_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool))
	var _155_v49 *C0
	_ = _155_v49
	var _nw19 *C0 = New_C0_()
	_ = _nw19
	_nw19.Ctor__(_99_v2, _97_v0, _101_v4)
	_155_v49 = _nw19
	_99_v2 = ((_155_v49).F12()).Times(((_155_v49).F12()).Minus((_153_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.IntOfUint32((_153_v48).Cardinality()))).Uint32()).(_dafny.Int)))
	var _hi2 _dafny.Int = _99_v2
	_ = _hi2
	for _156_i5 := ((_105_v8).Cardinality()).Plus((_155_v49).F12()); _156_i5.Cmp(_hi2) < 0; _156_i5 = _156_i5.Plus(_dafny.One) {
		(_106_globalState).F8 = _97_v0
		var _157_v50 *C0
		_ = _157_v50
		var _nw20 *C0 = New_C0_()
		_ = _nw20
		_nw20.Ctor__((_dafny.Zero).Minus(((_155_v49).F12()).Times((_103_v6).Cardinality())), (func() bool {
			if (_105_v8).Contains((_155_v49).F12()) {
				return (_105_v8).Get((_155_v49).F12()).(bool)
			}
			return true
		})(), _101_v4)
		_157_v50 = _nw20
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))
		_ = _index20
		(_102_v5).ArraySet1(!(((Companion_Default___.Fm8(_106_globalState)).Intersection(_dafny.MultiSetFromSeq(_153_v48))).IsProperSubsetOf(_140_v37)), (_index20).Int())
		var _158_v51 _dafny.Sequence
		_ = _158_v51
		_158_v51 = _dafny.SeqOf(_157_v50, _157_v50, _157_v50)
		if (_dafny.IntOfUint32((_158_v51).Cardinality())).Cmp(_99_v2) == 0 {
			var _159_v52 _dafny.Sequence
			_ = _159_v52
			var _out2 _dafny.Sequence
			_ = _out2
			_out2 = Companion_Default___.M0((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool), _98_v1, _106_globalState)
			_159_v52 = _out2
			var _160_v53 _dafny.Sequence
			_ = _160_v53
			var _out3 _dafny.Sequence
			_ = _out3
			_out3 = Companion_Default___.M0((_155_v49).Fm6(_97_v0, _99_v2, true, _106_globalState), _98_v1, _106_globalState)
			_160_v53 = _out3
			(_106_globalState).F8 = _97_v0
			var _rhs45 bool = _97_v0
			_ = _rhs45
			var _rhs46 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_155_v49).F12()), ((_155_v49).F12()).Minus(_dafny.IntOfInt64(-61)))
			_ = _rhs46
			var _lhs23 *GlobalState = _106_globalState
			_ = _lhs23
			_lhs23.F5 = _rhs45
			_99_v2 = _rhs46
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))
			_ = _index21
			(_102_v5).ArraySet1(_97_v0, (_index21).Int())
		} else {
			_99_v2 = (_157_v50).F12()
			_99_v2 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_99_v2).Minus(_dafny.IntOfInt64(458)), (_dafny.Zero).Minus((_155_v49).F12())))
			_142_v39 = _142_v39
			var _161_v54 D2
			_ = _161_v54
			_161_v54 = Companion_D2_.Create_DC6_(_dafny.SetOf((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool)))
			var _162_v55 _dafny.Map
			_ = _162_v55
			_162_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v54, (_157_v50).F12())
			_99_v2 = (func() _dafny.Int {
				if (_162_v55).Contains(_161_v54) {
					return (_162_v55).Get(_161_v54).(_dafny.Int)
				}
				return ((_155_v49).F12()).Plus((_155_v49).F12())
			})()
			var _163_v56 _dafny.Sequence
			_ = _163_v56
			var _out4 _dafny.Sequence
			_ = _out4
			_out4 = Companion_Default___.M0((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool), _98_v1, _106_globalState)
			_163_v56 = _out4
		}
	}
	var _pat_let_tv3 = _102_v5
	_ = _pat_let_tv3
	var _pat_let_tv4 = _102_v5
	_ = _pat_let_tv4
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))
	_ = _index22
	(_102_v5).ArraySet1((func(_pat_let2_0 D2) D2 {
		return func(_164_dt__update__tmp_h0 D2) D2 {
			return func(_pat_let3_0 bool) D2 {
				return func(_165_dt__update_hcf10_h0 bool) D2 {
					return Companion_D2_.Create_DC8_((_164_dt__update__tmp_h0).Dtor_cf9(), _165_dt__update_hcf10_h0)
				}(_pat_let3_0)
			}((_pat_let_tv4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_pat_let_tv3), 0))).Int()).(bool))
		}(_pat_let2_0)
	}(Companion_D2_.Create_DC8_(_97_v0, (_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool)))).Dtor_cf9(), (_index22).Int())
	var _166_v57 *C0
	_ = _166_v57
	var _nw21 *C0 = New_C0_()
	_ = _nw21
	_nw21.Ctor__((_155_v49).F12(), _97_v0, Companion_D0_.Create_DC0_((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool)))
	_166_v57 = _nw21
	var _167_v58 _dafny.Sequence
	_ = _167_v58
	var _out5 _dafny.Sequence
	_ = _out5
	_out5 = Companion_Default___.M0((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool), _dafny.SeqOf((_102_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_102_v5), 0))).Int()).(bool)), _106_globalState)
	_167_v58 = _out5
	_dafny.Print(_97_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_98_v1, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_99_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v3).Equals(_dafny.MultiSetOf(_dafny.SeqOf(true), _dafny.SeqOf(false, false), _dafny.SeqOf(false, false), _dafny.SeqOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v4).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(147))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v7).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_104_v7).Dtor_cf5()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(147))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v7).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F0()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(true), _dafny.SeqOf(false, false), _dafny.SeqOf(false, false), _dafny.SeqOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_106_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F6()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(147))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_globalState).F7()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_106_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_107_v9).Equals(_dafny.MultiSetOf(true, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v36).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.IntOfInt64(147))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_140_v37).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(717))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_141_v38.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_142_v39)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v40).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_146_v41).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_v43).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-104), false).UpdateUnsafe(_dafny.IntOfInt64(-105), false).UpdateUnsafe(_dafny.IntOfInt64(-106), false).UpdateUnsafe(_dafny.IntOfInt64(-107), false).UpdateUnsafe(_dafny.IntOfInt64(-108), false).UpdateUnsafe(_dafny.IntOfInt64(-109), false).UpdateUnsafe(_dafny.IntOfInt64(-110), false).UpdateUnsafe(_dafny.IntOfInt64(-111), false).UpdateUnsafe(_dafny.IntOfInt64(-112), false).UpdateUnsafe(_dafny.IntOfInt64(-113), false).UpdateUnsafe(_dafny.IntOfInt64(-114), false).UpdateUnsafe(_dafny.IntOfInt64(-115), false).UpdateUnsafe(_dafny.IntOfInt64(-116), false).UpdateUnsafe(_dafny.IntOfInt64(-117), false).UpdateUnsafe(_dafny.IntOfInt64(-118), false).UpdateUnsafe(_dafny.IntOfInt64(-119), false).UpdateUnsafe(_dafny.IntOfInt64(-120), false).UpdateUnsafe(_dafny.IntOfInt64(-121), false).UpdateUnsafe(_dafny.IntOfInt64(-122), false).UpdateUnsafe(_dafny.IntOfInt64(-123), false).UpdateUnsafe(_dafny.IntOfInt64(-124), false).UpdateUnsafe(_dafny.IntOfInt64(-125), false).UpdateUnsafe(_dafny.IntOfInt64(-126), false).UpdateUnsafe(_dafny.IntOfInt64(-127), false).UpdateUnsafe(_dafny.IntOfInt64(-128), false).UpdateUnsafe(_dafny.IntOfInt64(-129), false).UpdateUnsafe(_dafny.IntOfInt64(-130), false).UpdateUnsafe(_dafny.IntOfInt64(-131), false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_150_v45).Equals(_dafny.SetOf(_dafny.SeqOf(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v46).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_153_v48, _dafny.SeqOf(_dafny.IntOfInt64(594))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v49).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v57).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_167_v58.VerbatimString(false))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_() D0 {
	return D0{D0_DC1{}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf1 bool
	Cf2 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf1 bool, Cf2 _dafny.Int) D0 {
	return D0{D0_DC2{Cf1, Cf2}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
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
	return Companion_D0_.Create_DC1_()
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC2).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf2
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
			return "D0.DC1"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
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
			_, ok := other.Get_().(D0_DC1)
			return ok
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0
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

type D1_DC4 struct {
	Cf4 bool
	Cf5 _dafny.Map
	Cf6 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 bool, Cf5 _dafny.Map, Cf6 bool) D1 {
	return D1{D1_DC4{Cf4, Cf5, Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf3 _dafny.Map
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Map) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC5 struct {
	Cf7 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 D1) D1 {
	return D1{D1_DC5{Cf7}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, _dafny.EmptyMap, false)
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf7() D1 {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5.Equals(data2.Cf5) && data1.Cf6 == data2.Cf6
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_() D2 {
	return D2{D2_DC7{}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf9  bool
	Cf10 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf9 bool, Cf10 bool) D2 {
	return D2{D2_DC8{Cf9, Cf10}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC6 struct {
	Cf8 _dafny.Set
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf8 _dafny.Set) D2 {
	return D2{D2_DC6{Cf8}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_()
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC8).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC8).Cf10
}

func (_this D2) Dtor_cf8() _dafny.Set {
	return _this.Get_().(D2_DC6).Cf8
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf8) + ")"
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
			_, ok := other.Get_().(D2_DC7)
			return ok
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf8.Equals(data2.Cf8)
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

type D3_DC9 struct {
	Cf11 _dafny.MultiSet
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf11 _dafny.MultiSet) D3 {
	return D3{D3_DC9{Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D3) Dtor_cf11() _dafny.MultiSet {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
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

type D4_DC11 struct {
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_() D4 {
	return D4{D4_DC11{}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf13 bool
	Cf14 T0
	Cf15 bool
	Cf16 bool
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf13 bool, Cf14 T0, Cf15 bool, Cf16 bool) D4 {
	return D4{D4_DC12{Cf13, Cf14, Cf15, Cf16}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf17 _dafny.Int
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf17 _dafny.Int) D4 {
	return D4{D4_DC13{Cf17}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC10 struct {
	Cf12 _dafny.Array
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf12 _dafny.Array) D4 {
	return D4{D4_DC10{Cf12}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_()
}

func (_this D4) Dtor_cf13() bool {
	return _this.Get_().(D4_DC12).Cf13
}

func (_this D4) Dtor_cf14() T0 {
	return _this.Get_().(D4_DC12).Cf14
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC12).Cf15
}

func (_this D4) Dtor_cf16() bool {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf17
}

func (_this D4) Dtor_cf12() _dafny.Array {
	return _this.Get_().(D4_DC10).Cf12
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC11:
		{
			_, ok := other.Get_().(D4_DC11)
			return ok
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf13 == data2.Cf13 && _dafny.AreEqual(data1.Cf14, data2.Cf14) && data1.Cf15 == data2.Cf15 && data1.Cf16 == data2.Cf16
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf12 == data2.Cf12
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

type D5_DC15 struct {
	Cf19 _dafny.Int
	Cf20 bool
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf19 _dafny.Int, Cf20 bool) D5 {
	return D5{D5_DC15{Cf19, Cf20}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC14 struct {
	Cf18 _dafny.Sequence
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf18 _dafny.Sequence) D5 {
	return D5{D5_DC14{Cf18}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(_dafny.Zero, false)
}

func (_this D5) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D5_DC15).Cf19
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC15).Cf20
}

func (_this D5) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D5_DC14).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm6(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool
	Fm7(p0 D0, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool
	F10() bool
	F11() D0
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

// Definition of class GlobalState
type GlobalState struct {
	F5  bool
	F8  bool
	_f0 _dafny.MultiSet
	_f1 _dafny.Array
	_f2 _dafny.Int
	_f3 bool
	_f4 _dafny.Int
	_f6 _dafny.Map
	_f7 _dafny.Map
	_f9 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F5 = false
	_this.F8 = false
	_this._f0 = _dafny.EmptyMultiSet
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = _dafny.Zero
	_this._f3 = false
	_this._f4 = _dafny.Zero
	_this._f6 = _dafny.EmptyMap
	_this._f7 = _dafny.EmptyMap
	_this._f9 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.MultiSet, f1 _dafny.Array, f2 _dafny.Int, f3 bool, f4 _dafny.Int, f5 bool, f6 _dafny.Map, f7 _dafny.Map, f8 bool, f9 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
	}
}
func (_this *GlobalState) F0() _dafny.MultiSet {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.Map {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Map {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f10 bool
	_f11 D0
	_f12 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f10 = false
	_this._f11 = Companion_D0_.Default()
	_this._f12 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F10() bool {
	return _this._f10
}
func (_this *C0) F11() D0 {
	return _this._f11
}
func (_this *C0) Ctor__(f12 _dafny.Int, f10 bool, f11 D0) {
	{
		(_this)._f12 = f12
		(_this)._f10 = f10
		(_this)._f11 = f11
	}
}
func (_this *C0) Fm6(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return ((_this).F12()).Cmp((_dafny.Zero).Minus((_this).F12())) > 0
	}
}
func (_this *C0) Fm7(p0 D0, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(false) || (false)
	}
}
func (_this *C0) F12() _dafny.Int {
	{
		return _this._f12
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
