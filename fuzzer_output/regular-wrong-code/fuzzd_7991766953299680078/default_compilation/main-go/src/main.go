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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-836)), _dafny.CodePoint('b'))).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vmxs"), _dafny.UnicodeSeqOfUtf8Bytes("nrusd"))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(911), _dafny.MultiSetOf(_dafny.IntOfInt64(-611), _dafny.IntOfInt64(2)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC2_(_dafny.IntOfInt64(780), _dafny.MultiSetOf(_dafny.IntOfInt64(160)), true)).Dtor_cf6(), _dafny.MultiSetOf(_dafny.IntOfInt64(253))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(235), _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality(), _dafny.CodePoint('b'))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("bgrolje")
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	var _source0 D2 = Companion_D2_.Create_DC10_()
	_ = _source0
	if _source0.Is_DC10() {
		return !(false)
	} else if _source0.Is_DC11() {
		var _0___mcc_h0 _dafny.CodePoint = _source0.Get_().(D2_DC11).Cf17
		_ = _0___mcc_h0
		var _1___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC11).Cf18
		_ = _1___mcc_h1
		var _2_cf18 _dafny.Int = _1___mcc_h1
		_ = _2_cf18
		var _3_cf17 _dafny.CodePoint = _0___mcc_h0
		_ = _3_cf17
		return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(30)), _dafny.SeqOf(_dafny.IntOfInt64(-447)))
	} else {
		var _4___mcc_h2 _dafny.Map = _source0.Get_().(D2_DC9).Cf16
		_ = _4___mcc_h2
		var _5_cf16 _dafny.Map = _4___mcc_h2
		_ = _5_cf16
		return (!((Companion_D0_.Create_DC2_((_5_cf16).Cardinality(), _dafny.MultiSetOf(_dafny.IntOfInt64(-769)), true)).Dtor_cf8())) == (false)
	}
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(true)
		}
		return _dafny.SeqOf(true)
	})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if false {
			return _dafny.SetOf(false, true)
		}
		return _dafny.SetOf(!(false), false, false)
	})()).Intersection((func() _dafny.Set {
		if false {
			return _dafny.SetOf(false, true)
		}
		return _dafny.SetOf(false)
	})())
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('h')
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 D2, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(true, (_dafny.IntOfInt64(7)).Cmp(_dafny.IntOfInt64(211)) < 0, (func() bool {
		if true {
			return false
		}
		return true
	})())
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Map {
	var _source1 D1 = Companion_D1_.Create_DC6_(_dafny.SetOf(false))
	_ = _source1
	if _source1.Is_DC5() {
		var _6___mcc_h0 _dafny.Int = _source1.Get_().(D1_DC5).Cf11
		_ = _6___mcc_h0
		var _7_cf11 _dafny.Int = _6___mcc_h0
		_ = _7_cf11
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_cf11, _7_cf11)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dtisl")).Cardinality()), _7_cf11))
	} else if _source1.Is_DC6() {
		var _8___mcc_h1 _dafny.Set = _source1.Get_().(D1_DC6).Cf12
		_ = _8___mcc_h1
		var _9_cf12 _dafny.Set = _8___mcc_h1
		_ = _9_cf12
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wf")).Cardinality()), _dafny.IntOfInt64(589))
	} else if _source1.Is_DC7() {
		var _10___mcc_h2 _dafny.Int = _source1.Get_().(D1_DC7).Cf13
		_ = _10___mcc_h2
		var _11___mcc_h3 _dafny.Int = _source1.Get_().(D1_DC7).Cf14
		_ = _11___mcc_h3
		var _12_cf14 _dafny.Int = _11___mcc_h3
		_ = _12_cf14
		var _13_cf13 _dafny.Int = _10___mcc_h2
		_ = _13_cf13
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_12_cf14, _13_cf13)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_cf13, (_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))).Cardinality()))
	} else if _source1.Is_DC4() {
		var _14___mcc_h4 _dafny.Array = _source1.Get_().(D1_DC4).Cf10
		_ = _14___mcc_h4
		var _15_cf10 _dafny.Array = _14___mcc_h4
		_ = _15_cf10
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((Companion_D5_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(262)), false))).Dtor_cf25()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bv")).Cardinality()), _dafny.IntOfInt64(521))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-972)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-630), _dafny.IntOfInt64(558))).Cardinality())).Cardinality(), _dafny.IntOfInt64(78))).Cardinality()), _dafny.IntOfInt64(682)))
	} else {
		var _16___mcc_h5 D1 = _source1.Get_().(D1_DC8).Cf15
		_ = _16___mcc_h5
		var _17_cf15 D1 = _16___mcc_h5
		_ = _17_cf15
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), false)).Cardinality(), _dafny.IntOfInt64(592))
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, globalState *GlobalState) D1 {
	if !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ttrqiyhpy"), _dafny.UnicodeSeqOfUtf8Bytes("twandbl")) {
		return Companion_D1_.Create_DC6_(_dafny.SetOf(false))
	} else {
		return Companion_D1_.Create_DC6_(_dafny.SetOf(false, !(true), true))
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Map, p1 bool, p2 _dafny.Set, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mwuiwydjh")).Cardinality()))).Difference((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nwoc")).Cardinality()), (_dafny.SetOf(!(false))).Cardinality())).Union(_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(66), _dafny.IntOfInt64(605))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _18_v0 _dafny.Int
	_ = _18_v0
	_18_v0 = _dafny.IntOfInt64(326)
	var _19_v1 _dafny.Sequence
	_ = _19_v1
	_19_v1 = _dafny.SeqOf(_18_v0, _18_v0)
	r0 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(222), _18_v0), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_19_v1).Cardinality()), _18_v0))
	var _20_v2 D2
	_ = _20_v2
	_20_v2 = Companion_D2_.Create_DC10_()
	var _source2 D2 = _20_v2
	_ = _source2
	if _source2.Is_DC10() {
		var _21_v3 _dafny.Array
		_ = _21_v3
		var _nw0 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(10))
		_ = _nw0
		_21_v3 = _nw0
		var _22_v4 _dafny.Sequence
		_ = _22_v4
		_22_v4 = _dafny.SeqOf(_21_v3, _21_v3, _21_v3)
		var _23_v5 _dafny.Array
		_ = _23_v5
		var _nwElement0_0 _dafny.Array = _21_v3
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(27))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		(_nw1).ArraySet1(_21_v3, 1)
		(_nw1).ArraySet1(_21_v3, 2)
		(_nw1).ArraySet1(_21_v3, 3)
		(_nw1).ArraySet1(_21_v3, 4)
		(_nw1).ArraySet1(_21_v3, 5)
		(_nw1).ArraySet1(_21_v3, 6)
		(_nw1).ArraySet1(_21_v3, 7)
		(_nw1).ArraySet1(_21_v3, 8)
		(_nw1).ArraySet1(_21_v3, 9)
		(_nw1).ArraySet1(_21_v3, 10)
		(_nw1).ArraySet1(_21_v3, 11)
		(_nw1).ArraySet1(_21_v3, 12)
		(_nw1).ArraySet1(_21_v3, 13)
		(_nw1).ArraySet1(_21_v3, 14)
		(_nw1).ArraySet1(_21_v3, 15)
		(_nw1).ArraySet1(_21_v3, 16)
		(_nw1).ArraySet1(_21_v3, 17)
		(_nw1).ArraySet1(_21_v3, 18)
		(_nw1).ArraySet1(_21_v3, 19)
		(_nw1).ArraySet1(_21_v3, 20)
		(_nw1).ArraySet1(_21_v3, 21)
		(_nw1).ArraySet1(_21_v3, 22)
		(_nw1).ArraySet1((_22_v4).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_22_v4).Cardinality()))).Uint32()).(_dafny.Array), 23)
		(_nw1).ArraySet1(_21_v3, 24)
		(_nw1).ArraySet1(_21_v3, 25)
		(_nw1).ArraySet1(_21_v3, 26)
		_23_v5 = _nw1
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(430), _dafny.ArrayLen((_23_v5), 0))
		_ = _index0
		(_23_v5).ArraySet1(_21_v3, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(430), _dafny.ArrayLen((_23_v5), 0))
		_ = _index1
		(_23_v5).ArraySet1(_21_v3, (_index1).Int())
		var _24_v6 _dafny.Array
		_ = _24_v6
		var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(7))
		_ = _nw2
		_24_v6 = _nw2
		var _25_v7 _dafny.Sequence
		_ = _25_v7
		_25_v7 = _dafny.SeqOf(p1, p1, p1)
		var _26_v8 _dafny.Set
		_ = _26_v8
		_26_v8 = _dafny.SetOf(_dafny.IntOfUint32((_25_v7).Cardinality()))
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_24_v6), 0))
		_ = _index2
		(_24_v6).ArraySet1((_dafny.SetOf(_18_v0)).Union(_26_v8), (_index2).Int())
		var _27_v9 _dafny.CodePoint
		_ = _27_v9
		_27_v9 = _dafny.CodePoint('v')
		var _28_v10 _dafny.Set
		_ = _28_v10
		_28_v10 = _dafny.SetOf(_27_v9, _27_v9)
		var _29_v11 _dafny.Map
		_ = _29_v11
		_29_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_25_v7).Cardinality()))
		var _30_v12 _dafny.Map
		_ = _30_v12
		_30_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _29_v11)
		var _31_v13 _dafny.Map
		_ = _31_v13
		_31_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v9, _dafny.IntOfUint32((_19_v1).Cardinality()))
		var _32_v14 _dafny.MultiSet
		_ = _32_v14
		_32_v14 = _dafny.MultiSetOf(_dafny.IntOfUint32((p0).Cardinality()), _18_v0, (_dafny.SetOf(_27_v9)).Cardinality(), (_31_v13).Cardinality(), _18_v0)
		var _33_v15 _dafny.Array
		_ = _33_v15
		var _nwElement0_1 _dafny.Int = _dafny.IntOfInt64(56)
		_ = _nwElement0_1
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(24))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_1, 0)
		(_nw3).ArraySet1(_18_v0, 1)
		(_nw3).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("a"), _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _27_v9))).Cardinality()), 2)
		(_nw3).ArraySet1(_18_v0, 3)
		(_nw3).ArraySet1(_dafny.IntOfUint32((_25_v7).Cardinality()), 4)
		(_nw3).ArraySet1((_dafny.Zero).Minus(_18_v0), 5)
		(_nw3).ArraySet1(Companion_Default___.SafeDivisionInt(_18_v0, _18_v0), 6)
		(_nw3).ArraySet1((_26_v8).Cardinality(), 7)
		(_nw3).ArraySet1(Companion_Default___.SafeModuloInt(_18_v0, _18_v0), 8)
		(_nw3).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 9)
		(_nw3).ArraySet1(_18_v0, 10)
		(_nw3).ArraySet1(_18_v0, 11)
		(_nw3).ArraySet1(_18_v0, 12)
		(_nw3).ArraySet1((_dafny.IntOfUint32((_25_v7).Cardinality())).Minus(_18_v0), 13)
		(_nw3).ArraySet1(_18_v0, 14)
		(_nw3).ArraySet1(((func() _dafny.Set {
			if p1 {
				return _28_v10
			}
			return _dafny.SetOf(_27_v9)
		})()).Cardinality(), 15)
		(_nw3).ArraySet1((_18_v0).Times(_18_v0), 16)
		(_nw3).ArraySet1((((func() _dafny.Map {
			if (_30_v12).Contains(p1) {
				return (_30_v12).Get(p1).(_dafny.Map)
			}
			return _29_v11
		})()).Cardinality()).Plus(_18_v0), 17)
		(_nw3).ArraySet1((_32_v14).Cardinality(), 18)
		(_nw3).ArraySet1((_dafny.Zero).Minus((_18_v0).Plus(_18_v0)), 19)
		(_nw3).ArraySet1(_18_v0, 20)
		(_nw3).ArraySet1(_18_v0, 21)
		(_nw3).ArraySet1((_dafny.Zero).Minus(_18_v0), 22)
		(_nw3).ArraySet1(_18_v0, 23)
		_33_v15 = _nw3
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_33_v15), 0))
		_ = _index3
		(_33_v15).ArraySet1(_18_v0, (_index3).Int())
		var _34_v16 D0
		_ = _34_v16
		_34_v16 = Companion_D0_.Create_DC0_(_26_v8, p0)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_24_v6), 0))
		_ = _index4
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_33_v15), 0))
		_ = _index5
		var _rhs0 _dafny.Set = (func(_pat_let0_0 D0) D0 {
			return func(_35_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 _dafny.Sequence) D0 {
					return func(_36_dt__update_hcf1_h0 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_((_35_dt__update__tmp_h0).Dtor_cf0(), _36_dt__update_hcf1_h0)
					}(_pat_let1_0)
				}(_dafny.UnicodeSeqOfUtf8Bytes("dqexxyx"))
			}(_pat_let0_0)
		}(_34_v16)).Dtor_cf0()
		_ = _rhs0
		var _rhs1 _dafny.Int = (_dafny.Zero).Minus(_18_v0)
		_ = _rhs1
		var _rhs2 _dafny.CodePoint = _27_v9
		_ = _rhs2
		var _rhs3 _dafny.Int = (_dafny.Zero).Minus((_32_v14).Cardinality())
		_ = _rhs3
		var _lhs0 _dafny.Array = _24_v6
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_24_v6), 0))
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		var _lhs3 _dafny.Array = _33_v15
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_33_v15), 0))
		_ = _lhs4
		(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
		_lhs2.F16 = _rhs1
		_27_v9 = _rhs2
		(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
		(globalState).F16 = _18_v0
		var _37_v17 D0
		_ = _37_v17
		_37_v17 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqOf((_33_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_33_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(906), _18_v0, (_33_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_33_v15), 0))).Int()).(_dafny.Int), (_33_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_33_v15), 0))).Int()).(_dafny.Int))).Cardinality()), _32_v14, p1)
		(globalState).F10 = (_dafny.Zero).Minus(((_dafny.SetOf(p1, p1)).Cardinality()).Minus((_18_v0).Times((_37_v17).Dtor_cf6())))
	} else if _source2.Is_DC11() {
		var _38___mcc_h0 _dafny.CodePoint = _source2.Get_().(D2_DC11).Cf17
		_ = _38___mcc_h0
		var _39___mcc_h1 _dafny.Int = _source2.Get_().(D2_DC11).Cf18
		_ = _39___mcc_h1
		var _40_cf18 _dafny.Int = _39___mcc_h1
		_ = _40_cf18
		var _41_cf17 _dafny.CodePoint = _38___mcc_h0
		_ = _41_cf17
		var _42_v18 D1
		_ = _42_v18
		_42_v18 = Companion_D1_.Create_DC5_(_18_v0)
		var _source3 D1 = _42_v18
		_ = _source3
		if _source3.Is_DC5() {
			var _43___mcc_h3 _dafny.Int = _source3.Get_().(D1_DC5).Cf11
			_ = _43___mcc_h3
			var _44_cf11 _dafny.Int = _43___mcc_h3
			_ = _44_cf11
			var _45_v19 *C0
			_ = _45_v19
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__(_18_v0, p1)
			_45_v19 = _nw4
			var _46_v20 _dafny.Map
			_ = _46_v20
			_46_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _47_v21 D1
			_ = _47_v21
			_47_v21 = Companion_D1_.Create_DC7_(_18_v0, (_46_v20).Cardinality())
			var _48_v22 _dafny.Map
			_ = _48_v22
			_48_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_44_cf11, p1)).Update(_dafny.IntOfUint32((p0).Cardinality()), p1), Companion_Default___.Fm3(p1, _40_cf18, _dafny.IntOfUint32((p0).Cardinality()), false, globalState))
			var _49_v23 _dafny.MultiSet
			_ = _49_v23
			_49_v23 = _dafny.MultiSetOf((_48_v22).Merge(_48_v22), _48_v22)
			var _50_v24 _dafny.Map
			_ = _50_v24
			_50_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_cf18, p1)
			var _51_v25 _dafny.Set
			_ = _51_v25
			_51_v25 = _dafny.SetOf(_40_cf18, (_50_v24).Cardinality())
			var _52_v26 _dafny.Map
			_ = _52_v26
			_52_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_51_v25).Cardinality(), p1)
			var _rhs4 D1 = _47_v21
			_ = _rhs4
			var _rhs5 _dafny.Int = (func() _dafny.Int {
				if (_49_v23).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v24, p1)) {
					return (_49_v23).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v24, p1))
				}
				return (_18_v0).Times(_dafny.IntOfInt64(888))
			})()
			_ = _rhs5
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			_47_v21 = _rhs4
			_lhs5.F10 = _rhs5
			_18_v0 = (_44_cf11).Times(_18_v0)
			var _53_v27 _dafny.Sequence
			_ = _53_v27
			_53_v27 = _dafny.SeqOf(p1, p1, p1)
			var _54_v28 T0
			_ = _54_v28
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__(_40_cf18, _dafny.Companion_Sequence_.Equal(_53_v27, _dafny.Companion_Sequence_.Update(_53_v27, (Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_53_v27).Cardinality()))).Uint32(), p1)))
			_54_v28 = _nw5
			_54_v28 = _54_v28
		} else if _source3.Is_DC6() {
			var _55___mcc_h4 _dafny.Set = _source3.Get_().(D1_DC6).Cf12
			_ = _55___mcc_h4
			var _56_cf12 _dafny.Set = _55___mcc_h4
			_ = _56_cf12
			var _57_v29 D1
			_ = _57_v29
			_57_v29 = Companion_D1_.Create_DC7_(Companion_Default___.SafeDivisionInt(_18_v0, _18_v0), _40_cf18)
			var _pat_let_tv0 = p0
			_ = _pat_let_tv0
			_57_v29 = func(_pat_let2_0 D1) D1 {
				return func(_58_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let3_0 _dafny.Int) D1 {
						return func(_59_dt__update_hcf14_h0 _dafny.Int) D1 {
							return Companion_D1_.Create_DC7_((_58_dt__update__tmp_h1).Dtor_cf13(), _59_dt__update_hcf14_h0)
						}(_pat_let3_0)
					}((_dafny.Zero).Minus(_dafny.IntOfUint32((_pat_let_tv0).Cardinality())))
				}(_pat_let2_0)
			}(_57_v29)
			var _60_v30 _dafny.MultiSet
			_ = _60_v30
			_60_v30 = _dafny.MultiSetOf(_18_v0, (_56_cf12).Cardinality(), _18_v0)
			var _61_v31 _dafny.Sequence
			_ = _61_v31
			_61_v31 = _dafny.SeqOf(_60_v30)
			var _62_v32 _dafny.Map
			_ = _62_v32
			_62_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_cf18, p1)
			var _63_v33 _dafny.Sequence
			_ = _63_v33
			_63_v33 = _dafny.SeqOf(p1)
			var _64_v34 _dafny.Array
			_ = _64_v34
			var _nwElement0_2 bool = p1
			_ = _nwElement0_2
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(20))
			_ = _nw6
			(_nw6).ArraySet1(_nwElement0_2, 0)
			(_nw6).ArraySet1((_40_cf18).Cmp(_40_cf18) != 0, 1)
			(_nw6).ArraySet1(Companion_Default___.Fm3(p1, _18_v0, _40_cf18, p1, globalState), 2)
			(_nw6).ArraySet1((!(p1)) && (p1), 3)
			(_nw6).ArraySet1(p1, 4)
			(_nw6).ArraySet1(p1, 5)
			(_nw6).ArraySet1(((_61_v31).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_61_v31).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_18_v0))), 6)
			(_nw6).ArraySet1(!((func() bool {
				if (_62_v32).Contains(_18_v0) {
					return (_62_v32).Get(_18_v0).(bool)
				}
				return p1
			})()), 7)
			(_nw6).ArraySet1(p1, 8)
			(_nw6).ArraySet1(p1, 9)
			(_nw6).ArraySet1(p1, 10)
			(_nw6).ArraySet1((_63_v33).Select((Companion_Default___.SafeIndex((Companion_D1_.Create_DC7_(_18_v0, _dafny.IntOfInt64(-298))).Dtor_cf13(), _dafny.IntOfUint32((_63_v33).Cardinality()))).Uint32()).(bool), 11)
			(_nw6).ArraySet1(p1, 12)
			(_nw6).ArraySet1(Companion_Default___.Fm3(p1, _18_v0, _dafny.IntOfUint32((_19_v1).Cardinality()), p1, globalState), 13)
			(_nw6).ArraySet1(((func() bool {
				if (_62_v32).Contains(_40_cf18) {
					return (_62_v32).Get(_40_cf18).(bool)
				}
				return p1
			})()) && (p1), 14)
			(_nw6).ArraySet1(p1, 15)
			(_nw6).ArraySet1(p1, 16)
			(_nw6).ArraySet1(_dafny.Companion_Sequence_.Equal(p0, p0), 17)
			(_nw6).ArraySet1(!(false), 18)
			(_nw6).ArraySet1(p1, 19)
			_64_v34 = _nw6
			var _65_v35 D3
			_ = _65_v35
			_65_v35 = Companion_D3_.Create_DC12_(_64_v34)
			_64_v34 = (Companion_D3_.Create_DC12_((_65_v35).Dtor_cf19())).Dtor_cf19()
			var _66_v36 _dafny.Array
			_ = _66_v36
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw7
			_66_v36 = _nw7
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_64_v34), 0))
			_ = _index6
			(_64_v34).ArraySet1(p1, (_index6).Int())
			var _67_v37 _dafny.Map
			_ = _67_v37
			_67_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _68_v38 _dafny.MultiSet
			_ = _68_v38
			_68_v38 = _dafny.MultiSetOf(p1, p1)
			var _69_v39 *C0
			_ = _69_v39
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__((_68_v38).Cardinality(), p1)
			_69_v39 = _nw8
			var _70_v40 D0
			_ = _70_v40
			_70_v40 = Companion_D0_.Create_DC1_(p1, _67_v37, _69_v39, _dafny.IntOfUint32((_63_v33).Cardinality()))
			var _pat_let_tv1 = _67_v37
			_ = _pat_let_tv1
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_64_v34), 0))
			_ = _index7
			var _rhs6 _dafny.Array = _66_v36
			_ = _rhs6
			var _rhs7 _dafny.Int = (func(_pat_let4_0 D0) D0 {
				return func(_71_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let5_0 _dafny.Map) D0 {
						return func(_72_dt__update_hcf3_h0 _dafny.Map) D0 {
							return Companion_D0_.Create_DC1_((_71_dt__update__tmp_h2).Dtor_cf2(), _72_dt__update_hcf3_h0, (_71_dt__update__tmp_h2).Dtor_cf4(), (_71_dt__update__tmp_h2).Dtor_cf5())
						}(_pat_let5_0)
					}(_pat_let_tv1)
				}(_pat_let4_0)
			}(_70_v40)).Dtor_cf5()
			_ = _rhs7
			var _rhs8 bool = !((_62_v32).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_cf18, false)))
			_ = _rhs8
			var _rhs9 bool = p1
			_ = _rhs9
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			var _lhs7 _dafny.Array = _64_v34
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_64_v34), 0))
			_ = _lhs8
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			_66_v36 = _rhs6
			_lhs6.F16 = _rhs7
			(_lhs7).ArraySet1(_rhs8, (_lhs8).Int())
			_lhs9.F4 = _rhs9
			var _73_v41 *C0
			_ = _73_v41
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__(_18_v0, (p1) || (p1))
			_73_v41 = _nw9
		} else if _source3.Is_DC7() {
			var _74___mcc_h5 _dafny.Int = _source3.Get_().(D1_DC7).Cf13
			_ = _74___mcc_h5
			var _75___mcc_h6 _dafny.Int = _source3.Get_().(D1_DC7).Cf14
			_ = _75___mcc_h6
			var _76_cf14 _dafny.Int = _75___mcc_h6
			_ = _76_cf14
			var _77_cf13 _dafny.Int = _74___mcc_h5
			_ = _77_cf13
			(globalState).F16 = _77_cf13
			var _78_v42 *C0
			_ = _78_v42
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__(_18_v0, p1)
			_78_v42 = _nw10
			var _79_v43 _dafny.Map
			_ = _79_v43
			_79_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _78_v42)
			var _80_v44 _dafny.Array
			_ = _80_v44
			var _nwElement0_3 *C0 = (func() *C0 {
				if p1 {
					return _78_v42
				}
				return _78_v42
			})()
			_ = _nwElement0_3
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(17))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_3, 0)
			(_nw11).ArraySet1(_78_v42, 1)
			(_nw11).ArraySet1(_78_v42, 2)
			(_nw11).ArraySet1(_78_v42, 3)
			(_nw11).ArraySet1(_78_v42, 4)
			(_nw11).ArraySet1(_78_v42, 5)
			(_nw11).ArraySet1(_78_v42, 6)
			(_nw11).ArraySet1(_78_v42, 7)
			(_nw11).ArraySet1(_78_v42, 8)
			(_nw11).ArraySet1(_78_v42, 9)
			(_nw11).ArraySet1(_78_v42, 10)
			(_nw11).ArraySet1(_78_v42, 11)
			(_nw11).ArraySet1(_78_v42, 12)
			(_nw11).ArraySet1(_78_v42, 13)
			(_nw11).ArraySet1((func() *C0 {
				if (_79_v43).Contains(p1) {
					return (_79_v43).Get(p1).(*C0)
				}
				return _78_v42
			})(), 14)
			(_nw11).ArraySet1(_78_v42, 15)
			(_nw11).ArraySet1(_78_v42, 16)
			_80_v44 = _nw11
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_80_v44), 0))
			_ = _index8
			(_80_v44).ArraySet1(_78_v42, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_80_v44), 0))
			_ = _index9
			(_80_v44).ArraySet1(_78_v42, (_index9).Int())
			_78_v42 = _78_v42
			var _81_v45 D2
			_ = _81_v45
			_81_v45 = Companion_D2_.Create_DC11_(_41_cf17, _dafny.IntOfInt64(-243))
			var _82_v46 _dafny.Map
			_ = _82_v46
			_82_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v42, (_81_v45).Dtor_cf18())
			var _83_v47 _dafny.Sequence
			_ = _83_v47
			_83_v47 = _dafny.SeqOf(p0)
			_82_v46 = (_82_v46).Update((_80_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_80_v44), 0))).Int()).(*C0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, (_83_v47).Select((Companion_Default___.SafeIndex(_40_cf18, _dafny.IntOfUint32((_83_v47).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()))
		} else if _source3.Is_DC4() {
			var _84___mcc_h7 _dafny.Array = _source3.Get_().(D1_DC4).Cf10
			_ = _84___mcc_h7
			var _85_cf10 _dafny.Array = _84___mcc_h7
			_ = _85_cf10
			(globalState).F4 = p1
			var _86_v48 _dafny.Map
			_ = _86_v48
			_86_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _87_v49 *C0
			_ = _87_v49
			var _nw12 *C0 = New_C0_()
			_ = _nw12
			_nw12.Ctor__(_40_cf18, p1)
			_87_v49 = _nw12
			var _88_v50 _dafny.Set
			_ = _88_v50
			_88_v50 = _dafny.SetOf(_40_cf18)
			var _89_v51 D0
			_ = _89_v51
			_89_v51 = Companion_D0_.Create_DC1_(p1, _86_v48, _87_v49, (_88_v50).Cardinality())
			var _90_v52 _dafny.Sequence
			_ = _90_v52
			_90_v52 = _dafny.SeqOf(true, p1, p1)
			var _91_v53 _dafny.Array
			_ = _91_v53
			var _nwElement0_4 bool = p1
			_ = _nwElement0_4
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(27))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_4, 0)
			(_nw13).ArraySet1(p1, 1)
			(_nw13).ArraySet1((func() bool {
				if true {
					return p1
				}
				return p1
			})(), 2)
			(_nw13).ArraySet1(p1, 3)
			(_nw13).ArraySet1(true, 4)
			(_nw13).ArraySet1(p1, 5)
			(_nw13).ArraySet1(p1, 6)
			(_nw13).ArraySet1(p1, 7)
			(_nw13).ArraySet1(p1, 8)
			(_nw13).ArraySet1(!(!((_89_v51).Dtor_cf2())), 9)
			(_nw13).ArraySet1(p1, 10)
			(_nw13).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(p0, p0), 11)
			(_nw13).ArraySet1((_89_v51).Dtor_cf2(), 12)
			(_nw13).ArraySet1(!(p1), 13)
			(_nw13).ArraySet1(!(p1), 14)
			(_nw13).ArraySet1(p1, 15)
			(_nw13).ArraySet1((_90_v52).Select((Companion_Default___.SafeIndex((_19_v1).Select((Companion_Default___.SafeIndex(_40_cf18, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_90_v52).Cardinality()))).Uint32()).(bool), 16)
			(_nw13).ArraySet1(p1, 17)
			(_nw13).ArraySet1(p1, 18)
			(_nw13).ArraySet1(!((_dafny.IntOfInt64(757)).Cmp(_40_cf18) == 0), 19)
			(_nw13).ArraySet1(p1, 20)
			(_nw13).ArraySet1(true, 21)
			(_nw13).ArraySet1(p1, 22)
			(_nw13).ArraySet1(p1, 23)
			(_nw13).ArraySet1(!(p1) || (true), 24)
			(_nw13).ArraySet1(p1, 25)
			(_nw13).ArraySet1((p1) && (p1), 26)
			_91_v53 = _nw13
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_91_v53), 0))
			_ = _index10
			(_91_v53).ArraySet1(p1, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_85_cf10), 0))
			_ = _index11
			(_85_cf10).ArraySet1(_40_cf18, (_index11).Int())
			var _92_v54 _dafny.Set
			_ = _92_v54
			_92_v54 = _dafny.SetOf(p1)
			var _93_v55 D1
			_ = _93_v55
			_93_v55 = Companion_D1_.Create_DC6_(_92_v54)
			var _94_v56 D1
			_ = _94_v56
			_94_v56 = Companion_D1_.Create_DC8_(_93_v55)
			var _95_v57 D3
			_ = _95_v57
			_95_v57 = Companion_D3_.Create_DC14_(Companion_D3_.Create_DC14_(Companion_D3_.Create_DC13_(p1, p1, _94_v56)))
			var _96_v58 _dafny.Map
			_ = _96_v58
			_96_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_cf17, _95_v57)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_91_v53), 0))
			_ = _index12
			(_91_v53).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), _95_v57)).Equals(_96_v58), (_index12).Int())
			var _97_v59 D0
			_ = _97_v59
			_97_v59 = Companion_D0_.Create_DC0_(_dafny.SetOf(_40_cf18), _dafny.UnicodeSeqOfUtf8Bytes("xsto"))
			var _98_v60 _dafny.Map
			_ = _98_v60
			_98_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v59, (func() bool {
				if p1 {
					return !(p1)
				}
				return !(p1)
			})())
			var _99_v61 D1
			_ = _99_v61
			_99_v61 = Companion_D1_.Create_DC6_(_dafny.SetOf(p1))
			var _pat_let_tv2 = _90_v52
			_ = _pat_let_tv2
			var _pat_let_tv3 = _18_v0
			_ = _pat_let_tv3
			var _pat_let_tv4 = _90_v52
			_ = _pat_let_tv4
			var _pat_let_tv5 = p1
			_ = _pat_let_tv5
			var _100_v62 _dafny.MultiSet
			_ = _100_v62
			_100_v62 = _dafny.MultiSetOf(_99_v61, _99_v61, _99_v61, _99_v61, func(_pat_let6_0 D1) D1 {
				return func(_101_dt__update__tmp_h3 D1) D1 {
					return func(_pat_let7_0 _dafny.Set) D1 {
						return func(_102_dt__update_hcf12_h0 _dafny.Set) D1 {
							return Companion_D1_.Create_DC6_(_102_dt__update_hcf12_h0)
						}(_pat_let7_0)
					}(_dafny.SetOf((_pat_let_tv2).Select((Companion_Default___.SafeIndex(_pat_let_tv3, _dafny.IntOfUint32((_pat_let_tv4).Cardinality()))).Uint32()).(bool), _pat_let_tv5))
				}(_pat_let6_0)
			}(_99_v61))
			var _103_v63 _dafny.Set
			_ = _103_v63
			_103_v63 = _dafny.SetOf(p0)
			var _104_v64 _dafny.Sequence
			_ = _104_v64
			_104_v64 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v59, p1))
			var _105_v66 _dafny.Map
			_ = _105_v66
			_105_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v59, _41_cf17)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_91_v53), 0))
			_ = _index13
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_85_cf10), 0))
			_ = _index14
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_91_v53), 0))
			_ = _index15
			var _rhs10 bool = !(_100_v62).Equals(_100_v62)
			_ = _rhs10
			var _rhs11 _dafny.Int = Companion_Default___.SafeDivisionInt(_18_v0, (_103_v63).Cardinality())
			_ = _rhs11
			var _rhs12 bool = p1
			_ = _rhs12
			var _rhs13 D2 = Companion_D2_.Create_DC10_()
			_ = _rhs13
			var _rhs14 _dafny.Map = ((_104_v64).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_18_v0), _dafny.IntOfUint32((_104_v64).Cardinality()))).Uint32()).(_dafny.Map)).Merge(func() _dafny.Map {
				var _coll0 = _dafny.NewMapBuilder()
				_ = _coll0
				for _iter0 := _dafny.Iterate((_105_v66).Keys().Elements()); ; {
					_compr_0, _ok0 := _iter0()
					if !_ok0 {
						break
					}
					var _106_v65 D0
					_106_v65 = interface{}(_compr_0).(D0)
					if (_105_v66).Contains(_106_v65) {
						_coll0.Add(_106_v65, p1)
					}
				}
				return _coll0.ToMap()
			}())
			_ = _rhs14
			var _lhs10 _dafny.Array = _91_v53
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_91_v53), 0))
			_ = _lhs11
			var _lhs12 _dafny.Array = _85_cf10
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_85_cf10), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _91_v53
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_91_v53), 0))
			_ = _lhs15
			(_lhs10).ArraySet1(_rhs10, (_lhs11).Int())
			(_lhs12).ArraySet1(_rhs11, (_lhs13).Int())
			(_lhs14).ArraySet1(_rhs12, (_lhs15).Int())
			_20_v2 = _rhs13
			_98_v60 = _rhs14
			(globalState).F4 = (_89_v51).Dtor_cf2()
			r0 = _40_cf18
		} else {
			var _107___mcc_h8 D1 = _source3.Get_().(D1_DC8).Cf15
			_ = _107___mcc_h8
			var _108_cf15 D1 = _107___mcc_h8
			_ = _108_cf15
			var _109_v67 _dafny.Sequence
			_ = _109_v67
			_109_v67 = _dafny.SeqOf(Companion_Default___.Fm2(p1, p1, globalState), _dafny.Companion_Sequence_.Concatenate(p0, p0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(413))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg0 _dafny.Int) interface{} {
					return coer0(arg0)
				}
			}(func(_110_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), p0)
			_109_v67 = _109_v67
			var _111_v68 _dafny.Map
			_ = _111_v68
			_111_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm7(globalState))
			(globalState).F4 = (((_111_v68).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(242))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}((func(_112_cf17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_113_i1 _dafny.Int) _dafny.CodePoint {
					return _112_cf17
				}
			})(_41_cf17)))).Cardinality()))).Cmp((func() _dafny.Int {
				if p1 {
					return _18_v0
				}
				return _40_cf18
			})()) < 0
			var _114_v69 T0
			_ = _114_v69
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_dafny.IntOfInt64(339), p1)
			_114_v69 = _nw14
			var _115_v70 _dafny.Array
			_ = _115_v70
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw15
			_115_v70 = _nw15
			var _116_v71 D1
			_ = _116_v71
			_116_v71 = Companion_D1_.Create_DC4_(_115_v70)
			var _117_v72 D1
			_ = _117_v72
			_117_v72 = Companion_D1_.Create_DC8_(Companion_D1_.Create_DC8_(_116_v71))
			var _118_v73 _dafny.Map
			_ = _118_v73
			_118_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v69, Companion_D3_.Create_DC14_(Companion_D3_.Create_DC13_((_114_v69).F19(), p1, _117_v72)))
			var _119_v74 _dafny.Array
			_ = _119_v74
			var _nwElement0_5 bool = p1
			_ = _nwElement0_5
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(10))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_5, 0)
			(_nw16).ArraySet1((_114_v69).F19(), 1)
			(_nw16).ArraySet1(p1, 2)
			(_nw16).ArraySet1(p1, 3)
			(_nw16).ArraySet1((_114_v69).F19(), 4)
			(_nw16).ArraySet1(p1, 5)
			(_nw16).ArraySet1(p1, 6)
			(_nw16).ArraySet1((_114_v69).F19(), 7)
			(_nw16).ArraySet1(p1, 8)
			(_nw16).ArraySet1(false, 9)
			_119_v74 = _nw16
			var _120_v75 D3
			_ = _120_v75
			_120_v75 = Companion_D3_.Create_DC12_(_119_v74)
			var _121_v76 D3
			_ = _121_v76
			_121_v76 = Companion_D3_.Create_DC14_(_120_v75)
			var _122_v77 _dafny.Sequence
			_ = _122_v77
			_122_v77 = _dafny.SeqOf(_118_v73, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v69, _121_v76))
			var _123_v78 _dafny.MultiSet
			_ = _123_v78
			_123_v78 = _dafny.MultiSetOf(p1, (_114_v69).F19())
			var _124_v79 _dafny.Sequence
			_ = _124_v79
			_124_v79 = _dafny.SeqOf((_122_v77).Select((Companion_Default___.SafeIndex((_123_v78).Cardinality(), _dafny.IntOfUint32((_122_v77).Cardinality()))).Uint32()).(_dafny.Map))
			var _125_v80 _dafny.Set
			_ = _125_v80
			_125_v80 = _dafny.SetOf((_114_v69).F19(), (_114_v69).F19(), p1)
			var _126_v81 _dafny.Map
			_ = _126_v81
			_126_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v69.F18(), _125_v80)
			var _127_v82 *C0
			_ = _127_v82
			var _nw17 *C0 = New_C0_()
			_ = _nw17
			_nw17.Ctor__(_dafny.IntOfUint32((p0).Cardinality()), p1)
			_127_v82 = _nw17
			var _128_v83 _dafny.Set
			_ = _128_v83
			_128_v83 = _dafny.SetOf(_127_v82)
			var _129_v84 _dafny.Map
			_ = _129_v84
			_129_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_114_v69).F19(), (_114_v69).F19())
			var _130_v85 D2
			_ = _130_v85
			_130_v85 = Companion_D2_.Create_DC11_(_41_cf17, _40_cf18)
			var _131_v86 _dafny.Map
			_ = _131_v86
			_131_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-886), p1)
			var _132_v87 _dafny.Array
			_ = _132_v87
			var _nwElement0_6 D2 = _130_v85
			_ = _nwElement0_6
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(5))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_6, 0)
			(_nw18).ArraySet1(_130_v85, 1)
			(_nw18).ArraySet1(_130_v85, 2)
			(_nw18).ArraySet1(Companion_D2_.Create_DC11_(_dafny.CodePoint('k'), (_131_v86).Cardinality()), 3)
			(_nw18).ArraySet1(_130_v85, 4)
			_132_v87 = _nw18
			var _133_v88 _dafny.MultiSet
			_ = _133_v88
			_133_v88 = _dafny.MultiSetOf((_127_v82).Fm4(_114_v69.F18(), (_114_v69).F19(), p1, globalState), _40_cf18)
			var _134_v89 _dafny.Array
			_ = _134_v89
			var _nwElement0_7 _dafny.Int = _40_cf18
			_ = _nwElement0_7
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(28))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_7, 0)
			(_nw19).ArraySet1((_18_v0).Times(_18_v0), 1)
			(_nw19).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(503), _40_cf18), 2)
			(_nw19).ArraySet1(((_124_v79).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_124_v79).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 3)
			(_nw19).ArraySet1(_40_cf18, 4)
			(_nw19).ArraySet1((_dafny.IntOfInt64(746)).Times(Companion_Default___.Fm0(((func() _dafny.Set {
				if (_126_v81).Contains(_18_v0) {
					return (_126_v81).Get(_18_v0).(_dafny.Set)
				}
				return _125_v80
			})()).Cardinality(), true, globalState)), 5)
			(_nw19).ArraySet1(_40_cf18, 6)
			(_nw19).ArraySet1((_128_v83).Cardinality(), 7)
			(_nw19).ArraySet1(_40_cf18, 8)
			(_nw19).ArraySet1(_114_v69.F18(), 9)
			(_nw19).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 10)
			(_nw19).ArraySet1(_114_v69.F18(), 11)
			(_nw19).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality())), 12)
			(_nw19).ArraySet1(((_129_v84).Merge(_129_v84)).Cardinality(), 13)
			(_nw19).ArraySet1(_114_v69.F18(), 14)
			(_nw19).ArraySet1(_dafny.IntOfInt64(594), 15)
			(_nw19).ArraySet1((_dafny.Zero).Minus((_114_v69.F18()).Plus(_18_v0)), 16)
			(_nw19).ArraySet1(_18_v0, 17)
			(_nw19).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_114_v69.F18()), (_dafny.Zero).Minus(_114_v69.F18())), 18)
			(_nw19).ArraySet1((func() _dafny.Int {
				if Companion_Default___.Fm3(false, _40_cf18, _114_v69.F18(), (_114_v69).F19(), globalState) {
					return _dafny.IntOfUint32((_dafny.SeqOf(p1, !(p1), p1, (_114_v69).F19(), (_114_v69).F19())).Cardinality())
				}
				return _dafny.IntOfUint32((p0).Cardinality())
			})(), 19)
			(_nw19).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_114_v69.F18())), _132_v87)).Cardinality(), 20)
			(_nw19).ArraySet1((func() _dafny.Int {
				if false {
					return _18_v0
				}
				return _114_v69.F18()
			})(), 21)
			(_nw19).ArraySet1(Companion_Default___.SafeModuloInt(_18_v0, (func() _dafny.Int {
				if (_133_v88).Contains((_dafny.Zero).Minus(_18_v0)) {
					return (_133_v88).Multiplicity((_dafny.Zero).Minus(_18_v0))
				}
				return _114_v69.F18()
			})()), 22)
			(_nw19).ArraySet1((_114_v69.F18()).Times(_40_cf18), 23)
			(_nw19).ArraySet1((_114_v69.F18()).Minus(_18_v0), 24)
			(_nw19).ArraySet1(Companion_Default___.SafeModuloInt(_114_v69.F18(), _18_v0), 25)
			(_nw19).ArraySet1(Companion_Default___.SafeModuloInt(_114_v69.F18(), _114_v69.F18()), 26)
			(_nw19).ArraySet1(_dafny.IntOfInt64(713), 27)
			_134_v89 = _nw19
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_115_v70), 0))
			_ = _index16
			(_115_v70).ArraySet1(_114_v69.F18(), (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_115_v70), 0))
			_ = _index17
			(_115_v70).ArraySet1(((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_114_v69.F18(), _18_v0))).Minus(Companion_Default___.SafeDivisionInt(_114_v69.F18(), _18_v0)), (_index17).Int())
			var _135_v90 _dafny.Map
			_ = _135_v90
			_135_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v69, (_114_v69).F19())
			var _136_v91 _dafny.Map
			_ = _136_v91
			_136_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_135_v90).Cardinality())
			var _137_v92 _dafny.Set
			_ = _137_v92
			_137_v92 = _dafny.SetOf(((_136_v91).Update(true, _dafny.IntOfUint32((p0).Cardinality()))).Cardinality(), (_dafny.Zero).Minus(_114_v69.F18()))
			(globalState).F4 = ((_137_v92).Cardinality()).Cmp(_114_v69.F18()) != 0
		}
		var _138_v93 T0
		_ = _138_v93
		var _nw20 *C0 = New_C0_()
		_ = _nw20
		_nw20.Ctor__(_18_v0, !(p1))
		_138_v93 = _nw20
		var _139_v94 _dafny.Map
		_ = _139_v94
		_139_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(p1)), _138_v93)
		var _140_v95 _dafny.Sequence
		_ = _140_v95
		_140_v95 = _dafny.SeqOf(p1)
		_139_v94 = (_139_v94).Update(_140_v95, _138_v93)
		var _141_v96 _dafny.Set
		_ = _141_v96
		_141_v96 = _dafny.SetOf((_138_v93).F19())
		var _142_v97 _dafny.Map
		_ = _142_v97
		_142_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, (_19_v1).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int))).Update((_141_v96).Cardinality(), _138_v93.F18()), _dafny.UnicodeSeqOfUtf8Bytes("iyjjr"))
		var _143_v98 _dafny.Map
		_ = _143_v98
		_143_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, (_dafny.MultiSetOf(_138_v93.F18())).Cardinality())
		_142_v97 = (_142_v97).Update((_143_v98).Merge(_143_v98), _dafny.Companion_Sequence_.Concatenate(p0, p0))
		var _144_v99 _dafny.Map
		_ = _144_v99
		_144_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _138_v93.F18())
		if ((func() _dafny.Int {
			if p1 {
				return (func() _dafny.Int {
					if (_144_v99).Contains((_138_v93).F19()) {
						return (_144_v99).Get((_138_v93).F19()).(_dafny.Int)
					}
					return _138_v93.F18()
				})()
			}
			return _138_v93.F18()
		})()).Cmp(_40_cf18) < 0 {
			var _145_v100 _dafny.MultiSet
			_ = _145_v100
			_145_v100 = _dafny.MultiSetOf((_138_v93).F19())
			var _146_v101 _dafny.Array
			_ = _146_v101
			var _nwElement0_8 _dafny.MultiSet = _145_v100
			_ = _nwElement0_8
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_8, 0)
			(_nw21).ArraySet1((_145_v100).Intersection(_dafny.MultiSetOf(false, p1)), 1)
			(_nw21).ArraySet1((_145_v100).Update((_138_v93).F19(), Companion_Default___.Abs(_dafny.IntOfInt64(576))), 2)
			(_nw21).ArraySet1(_dafny.MultiSetFromSeq(_140_v95), 3)
			(_nw21).ArraySet1(_dafny.MultiSetFromSeq(_140_v95), 4)
			(_nw21).ArraySet1(_145_v100, 5)
			(_nw21).ArraySet1(_145_v100, 6)
			(_nw21).ArraySet1((_145_v100).Intersection(_145_v100), 7)
			(_nw21).ArraySet1(_145_v100, 8)
			(_nw21).ArraySet1(_dafny.MultiSetOf((_138_v93).F19(), !((_138_v93).F19())), 9)
			(_nw21).ArraySet1((_145_v100).Difference(_145_v100), 10)
			(_nw21).ArraySet1((_145_v100).Update(p1, Companion_Default___.Abs(_dafny.IntOfUint32((p0).Cardinality()))), 11)
			(_nw21).ArraySet1(_145_v100, 12)
			(_nw21).ArraySet1(_145_v100, 13)
			_146_v101 = _nw21
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_146_v101), 0))
			_ = _index18
			(_146_v101).ArraySet1(_145_v100, (_index18).Int())
			var _147_v102 _dafny.MultiSet
			_ = _147_v102
			_147_v102 = _145_v100
			var _148_v103 *C0
			_ = _148_v103
			var _nw22 *C0 = New_C0_()
			_ = _nw22
			_nw22.Ctor__(_dafny.IntOfInt64(263), (_dafny.MultiSetFromSeq(_140_v95)).IsProperSubsetOf((_147_v102)))
			_148_v103 = _nw22
			var _149_v104 _dafny.Array
			_ = _149_v104
			var _nwElement0_9 bool = (_138_v93).F19()
			_ = _nwElement0_9
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(16))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_9, 0)
			(_nw23).ArraySet1(p1, 1)
			(_nw23).ArraySet1(p1, 2)
			(_nw23).ArraySet1((_138_v93).F19(), 3)
			(_nw23).ArraySet1(!(p1), 4)
			(_nw23).ArraySet1((_138_v93).F19(), 5)
			(_nw23).ArraySet1((_138_v93).F19(), 6)
			(_nw23).ArraySet1((_138_v93).F19(), 7)
			(_nw23).ArraySet1((_138_v93).F19(), 8)
			(_nw23).ArraySet1(!(p1), 9)
			(_nw23).ArraySet1(p1, 10)
			(_nw23).ArraySet1(p1, 11)
			(_nw23).ArraySet1((_138_v93).F19(), 12)
			(_nw23).ArraySet1((_138_v93).F19(), 13)
			(_nw23).ArraySet1((_138_v93).F19(), 14)
			(_nw23).ArraySet1((_138_v93).F19(), 15)
			_149_v104 = _nw23
			var _150_v105 D3
			_ = _150_v105
			_150_v105 = Companion_D3_.Create_DC12_(_149_v104)
			var _151_v106 _dafny.Set
			_ = _151_v106
			_151_v106 = _dafny.SetOf((_141_v96).Cardinality())
			var _152_v107 D0
			_ = _152_v107
			_152_v107 = Companion_D0_.Create_DC0_(_151_v106, p0)
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_146_v101), 0))
			_ = _index19
			var _rhs15 _dafny.MultiSet = (func() _dafny.MultiSet {
				if p1 {
					return _145_v100
				}
				return _145_v100
			})()
			_ = _rhs15
			var _rhs16 *C0 = _148_v103
			_ = _rhs16
			var _rhs17 D3 = _150_v105
			_ = _rhs17
			var _rhs18 D0 = _152_v107
			_ = _rhs18
			var _lhs16 _dafny.Array = _146_v101
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_146_v101), 0))
			_ = _lhs17
			(_lhs16).ArraySet1(_rhs15, (_lhs17).Int())
			_148_v103 = _rhs16
			_150_v105 = _rhs17
			_152_v107 = _rhs18
			_147_v102 = _147_v102
			_143_v98 = (_143_v98).Update(Companion_Default___.SafeModuloInt(_138_v93.F18(), _dafny.IntOfUint32((p0).Cardinality())), _18_v0)
			(globalState).F4 = (_40_cf18).Cmp(_40_cf18) > 0
			var _153_v108 _dafny.Set
			_ = _153_v108
			_153_v108 = _dafny.SetOf(_148_v103, _148_v103)
			(globalState).F16 = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(175)).Minus(_18_v0), ((_153_v108).Intersection(_153_v108)).Cardinality())
		} else {
			(globalState).F4 = Companion_Default___.Fm3(p1, (func() _dafny.Int {
				if (_138_v93).F19() {
					return _40_cf18
				}
				return _138_v93.F18()
			})(), _40_cf18, p1, globalState)
			(globalState).F0 = (((_dafny.Zero).Minus(_138_v93.F18())).Plus(Companion_Default___.Fm0(_138_v93.F18(), (_138_v93).F19(), globalState))).Times(_dafny.IntOfInt64(-271))
			var _154_v109 _dafny.Array
			_ = _154_v109
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw24
			_154_v109 = _nw24
			var _155_v110 _dafny.Map
			_ = _155_v110
			_155_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v93, (_138_v93).F19())
			var _156_v111 _dafny.Set
			_ = _156_v111
			_156_v111 = _dafny.SetOf(_155_v110, _155_v110, _155_v110, _155_v110, _155_v110)
			var _157_v112 D0
			_ = _157_v112
			_157_v112 = Companion_D0_.Create_DC2_((_156_v111).Cardinality(), _dafny.MultiSetOf(_138_v93.F18(), _138_v93.F18(), _18_v0), (_138_v93).F19())
			var _158_v113 _dafny.Map
			_ = _158_v113
			_158_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v112, false)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_154_v109), 0))
			_ = _index20
			(_154_v109).ArraySet1(_158_v113, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_154_v109), 0))
			_ = _index21
			(_154_v109).ArraySet1(_158_v113, (_index21).Int())
			var _159_v114 _dafny.Set
			_ = _159_v114
			_159_v114 = _dafny.SetOf((_143_v98).Cardinality(), _40_cf18, _18_v0)
			var _160_v115 *C0
			_ = _160_v115
			var _nw25 *C0 = New_C0_()
			_ = _nw25
			_nw25.Ctor__((_dafny.Zero).Minus(_138_v93.F18()), (_138_v93).F19())
			_160_v115 = _nw25
			var _161_v116 _dafny.Sequence
			_ = _161_v116
			_161_v116 = _dafny.SeqOf(_160_v115, _160_v115)
			var _162_v117 _dafny.Map
			_ = _162_v117
			_162_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_159_v114).Cardinality(), p1)
			var _163_v118 _dafny.MultiSet
			_ = _163_v118
			_163_v118 = _dafny.MultiSetOf(p1, (_138_v93).F19(), p1, (_138_v93).F19(), (_138_v93).F19())
			var _164_v119 _dafny.MultiSet
			_ = _164_v119
			_164_v119 = _dafny.MultiSetOf(_138_v93)
			var _165_v120 D2
			_ = _165_v120
			_165_v120 = Companion_D2_.Create_DC11_(_41_cf17, _40_cf18)
			var _166_v121 _dafny.Array
			_ = _166_v121
			var _nwElement0_10 _dafny.Int = _dafny.IntOfUint32((p0).Cardinality())
			_ = _nwElement0_10
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(29))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_10, 0)
			(_nw26).ArraySet1((_19_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int), 1)
			(_nw26).ArraySet1(_138_v93.F18(), 2)
			(_nw26).ArraySet1(_138_v93.F18(), 3)
			(_nw26).ArraySet1((func() _dafny.Int {
				if p1 {
					return Companion_Default___.Fm0(_138_v93.F18(), (_138_v93).F19(), globalState)
				}
				return _18_v0
			})(), 4)
			(_nw26).ArraySet1(Companion_Default___.SafeModuloInt(_18_v0, _138_v93.F18()), 5)
			(_nw26).ArraySet1(_18_v0, 6)
			(_nw26).ArraySet1((_19_v1).Select((Companion_Default___.SafeIndex((_159_v114).Cardinality(), _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int), 7)
			(_nw26).ArraySet1((_40_cf18).Times(_dafny.IntOfUint32((_161_v116).Cardinality())), 8)
			(_nw26).ArraySet1(_138_v93.F18(), 9)
			(_nw26).ArraySet1(_138_v93.F18(), 10)
			(_nw26).ArraySet1((_40_cf18).Plus(_138_v93.F18()), 11)
			(_nw26).ArraySet1(_dafny.IntOfInt64(195), 12)
			(_nw26).ArraySet1(_dafny.IntOfInt64(930), 13)
			(_nw26).ArraySet1(_138_v93.F18(), 14)
			(_nw26).ArraySet1(_dafny.IntOfUint32((_140_v95).Cardinality()), 15)
			(_nw26).ArraySet1((func() _dafny.Int {
				if (_138_v93).F19() {
					return _138_v93.F18()
				}
				return (_162_v117).Cardinality()
			})(), 16)
			(_nw26).ArraySet1(_40_cf18, 17)
			(_nw26).ArraySet1(_138_v93.F18(), 18)
			(_nw26).ArraySet1((func() _dafny.Int {
				if (_163_v118).Contains(Companion_Default___.Fm3((_138_v93).F19(), _18_v0, _dafny.IntOfInt64(104), p1, globalState)) {
					return (_163_v118).Multiplicity(Companion_Default___.Fm3((_138_v93).F19(), _18_v0, _dafny.IntOfInt64(104), p1, globalState))
				}
				return (_19_v1).Select((Companion_Default___.SafeIndex(_40_cf18, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int)
			})(), 19)
			(_nw26).ArraySet1(_138_v93.F18(), 20)
			(_nw26).ArraySet1(_40_cf18, 21)
			(_nw26).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lcwsjh")).Cardinality()), 22)
			(_nw26).ArraySet1((_164_v119).Cardinality(), 23)
			(_nw26).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 24)
			(_nw26).ArraySet1((_165_v120).Dtor_cf18(), 25)
			(_nw26).ArraySet1((_160_v115).Fm4(_18_v0, (_138_v93).F19(), (_138_v93).F19(), globalState), 26)
			(_nw26).ArraySet1(_18_v0, 27)
			(_nw26).ArraySet1(_18_v0, 28)
			_166_v121 = _nw26
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_166_v121), 0))
			_ = _index22
			(_166_v121).ArraySet1(Companion_Default___.SafeModuloInt(_18_v0, _40_cf18), (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_166_v121), 0))
			_ = _index23
			(_166_v121).ArraySet1(_138_v93.F18(), (_index23).Int())
			var _167_v122 _dafny.Array
			_ = _167_v122
			var _nwElement0_11 _dafny.Set = (_141_v96).Difference(_141_v96)
			_ = _nwElement0_11
			var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(20))
			_ = _nw27
			(_nw27).ArraySet1(_nwElement0_11, 0)
			(_nw27).ArraySet1(_141_v96, 1)
			(_nw27).ArraySet1(_141_v96, 2)
			(_nw27).ArraySet1(_dafny.SetOf(p1), 3)
			(_nw27).ArraySet1(_141_v96, 4)
			(_nw27).ArraySet1(_141_v96, 5)
			(_nw27).ArraySet1(_141_v96, 6)
			(_nw27).ArraySet1((_141_v96).Difference(_141_v96), 7)
			(_nw27).ArraySet1((_141_v96).Union(_141_v96), 8)
			(_nw27).ArraySet1((Companion_Default___.Fm6((_166_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_166_v121), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_138_v93.F18(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _41_cf17), !((_138_v93).F19()), globalState)).Union(_141_v96), 9)
			(_nw27).ArraySet1(_141_v96, 10)
			(_nw27).ArraySet1((_141_v96).Intersection(_141_v96), 11)
			(_nw27).ArraySet1(_141_v96, 12)
			(_nw27).ArraySet1((_141_v96).Union(_141_v96), 13)
			(_nw27).ArraySet1((_141_v96).Intersection(_141_v96), 14)
			(_nw27).ArraySet1(_141_v96, 15)
			(_nw27).ArraySet1(_141_v96, 16)
			(_nw27).ArraySet1(_dafny.SetOf((_138_v93).F19(), (_138_v93).F19()), 17)
			(_nw27).ArraySet1((_141_v96).Difference(_141_v96), 18)
			(_nw27).ArraySet1(_141_v96, 19)
			_167_v122 = _nw27
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_167_v122), 0))
			_ = _index24
			(_167_v122).ArraySet1(_141_v96, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_167_v122), 0))
			_ = _index25
			(_167_v122).ArraySet1((_dafny.SetOf((_138_v93).F19(), !(true))).Difference(_141_v96), (_index25).Int())
		}
	} else {
		var _168___mcc_h2 _dafny.Map = _source2.Get_().(D2_DC9).Cf16
		_ = _168___mcc_h2
		var _169_cf16 _dafny.Map = _168___mcc_h2
		_ = _169_cf16
		var _170_v123 _dafny.MultiSet
		_ = _170_v123
		_170_v123 = _dafny.MultiSetOf(((_dafny.MultiSetOf(p1, p1)).Update(p1, Companion_Default___.Abs(_18_v0))).Cardinality())
		var _171_v124 T0
		_ = _171_v124
		var _nw28 *C0 = New_C0_()
		_ = _nw28
		_nw28.Ctor__(_18_v0, p1)
		_171_v124 = _nw28
		var _172_v125 _dafny.Map
		_ = _172_v125
		_172_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v124, p1)
		(globalState).F4 = Companion_Default___.Fm3(p1, (_19_v1).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_170_v123).Contains(Companion_Default___.Fm0((_172_v125).Cardinality(), p1, globalState)) {
				return (_170_v123).Multiplicity(Companion_Default___.Fm0((_172_v125).Cardinality(), p1, globalState))
			}
			return _171_v124.F18()
		})(), _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int), _18_v0, (_171_v124).F19(), globalState)
		var _173_v126 _dafny.Map
		_ = _173_v126
		_173_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(p1))
		var _174_v127 _dafny.MultiSet
		_ = _174_v127
		_174_v127 = _dafny.MultiSetOf(p1, (_171_v124).F19(), (_171_v124).F19())
		var _175_v128 _dafny.Map
		_ = _175_v128
		_175_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v127, (_173_v126).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Update(p1, true)))
		var _176_v129 _dafny.Sequence
		_ = _176_v129
		_176_v129 = _dafny.SeqOf(p1)
		var _177_v130 D2
		_ = _177_v130
		_177_v130 = Companion_D2_.Create_DC9_(_169_cf16)
		var _pat_let_tv6 = _169_cf16
		_ = _pat_let_tv6
		var _pat_let_tv7 = _169_cf16
		_ = _pat_let_tv7
		var _rhs19 _dafny.Map = (func() _dafny.Map {
			if (_175_v128).Contains((_dafny.MultiSetFromSeq(_176_v129)).Difference(Companion_Default___.Fm8(Companion_Default___.Fm0((_dafny.Zero).Minus(_171_v124.F18()), (_171_v124).F19(), globalState), func(_pat_let8_0 D2) D2 {
				return func(_178_dt__update__tmp_h4 D2) D2 {
					return func(_pat_let9_0 _dafny.Map) D2 {
						return func(_179_dt__update_hcf16_h0 _dafny.Map) D2 {
							return Companion_D2_.Create_DC9_(_179_dt__update_hcf16_h0)
						}(_pat_let9_0)
					}(_pat_let_tv6)
				}(_pat_let8_0)
			}(_177_v130), !(p1), p0, globalState))) {
				return (_175_v128).Get((_dafny.MultiSetFromSeq(_176_v129)).Difference(Companion_Default___.Fm8(Companion_Default___.Fm0((_dafny.Zero).Minus(_171_v124.F18()), (_171_v124).F19(), globalState), func(_pat_let10_0 D2) D2 {
					return func(_180_dt__update__tmp_h5 D2) D2 {
						return func(_pat_let11_0 _dafny.Map) D2 {
							return func(_181_dt__update_hcf16_h1 _dafny.Map) D2 {
								return Companion_D2_.Create_DC9_(_181_dt__update_hcf16_h1)
							}(_pat_let11_0)
						}(_pat_let_tv7)
					}(_pat_let10_0)
				}(_177_v130), !(p1), p0, globalState))).(_dafny.Map)
			}
			return _173_v126
		})()
		_ = _rhs19
		var _rhs20 bool = true
		_ = _rhs20
		var _rhs21 bool = (_171_v124).F19()
		_ = _rhs21
		var _lhs18 *GlobalState = globalState
		_ = _lhs18
		var _lhs19 *GlobalState = globalState
		_ = _lhs19
		_173_v126 = _rhs19
		_lhs18.F4 = _rhs20
		_lhs19.F4 = _rhs21
		var _182_v131 _dafny.CodePoint
		_ = _182_v131
		_182_v131 = _dafny.CodePoint('p')
		_182_v131 = _182_v131
		(globalState).F4 = !_dafny.Companion_Sequence_.Contains(p0, _182_v131)
	}
	(globalState).F16 = Companion_Default___.Fm0(_18_v0, p1, globalState)
	if !(!(p1)) {
		(globalState).F16 = (_19_v1).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int)
		var _183_v132 _dafny.Map
		_ = _183_v132
		_183_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v1, p1)).Cardinality())
		var _184_v133 *C0
		_ = _184_v133
		var _nw29 *C0 = New_C0_()
		_ = _nw29
		_nw29.Ctor__((_183_v132).Cardinality(), p1)
		_184_v133 = _nw29
		var _rhs22 *C0 = _184_v133
		_ = _rhs22
		var _rhs23 _dafny.Int = (func() _dafny.Int {
			if p1 {
				return _18_v0
			}
			return _18_v0
		})()
		_ = _rhs23
		var _lhs20 *GlobalState = globalState
		_ = _lhs20
		_184_v133 = _rhs22
		_lhs20.F0 = _rhs23
		var _185_v134 _dafny.Map
		_ = _185_v134
		_185_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_19_v1).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(_18_v0, Companion_Default___.Fm3(p1, _18_v0, _dafny.IntOfInt64(932), p1, globalState), globalState)))), p1)
		_185_v134 = (_185_v134).Update(_18_v0, p1)
		(globalState).F0 = _dafny.IntOfInt64(165)
		var _186_v135 _dafny.Array
		_ = _186_v135
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw30
		_186_v135 = _nw30
		_186_v135 = _186_v135
	} else {
		var _187_v136 _dafny.Array
		_ = _187_v136
		var _nwElement0_12 _dafny.Int = _18_v0
		_ = _nwElement0_12
		var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(24))
		_ = _nw31
		(_nw31).ArraySet1(_nwElement0_12, 0)
		(_nw31).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 1)
		(_nw31).ArraySet1(Companion_Default___.SafeModuloInt(_18_v0, _18_v0), 2)
		(_nw31).ArraySet1(_18_v0, 3)
		(_nw31).ArraySet1((_18_v0).Times(_18_v0), 4)
		(_nw31).ArraySet1(_18_v0, 5)
		(_nw31).ArraySet1(_18_v0, 6)
		(_nw31).ArraySet1(_18_v0, 7)
		(_nw31).ArraySet1(_dafny.IntOfInt64(521), 8)
		(_nw31).ArraySet1(_18_v0, 9)
		(_nw31).ArraySet1(_18_v0, 10)
		(_nw31).ArraySet1(_18_v0, 11)
		(_nw31).ArraySet1(_dafny.IntOfInt64(999), 12)
		(_nw31).ArraySet1(_18_v0, 13)
		(_nw31).ArraySet1(_18_v0, 14)
		(_nw31).ArraySet1((func() _dafny.Int {
			if p1 {
				return (_dafny.Zero).Minus(_18_v0)
			}
			return _18_v0
		})(), 15)
		(_nw31).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_18_v0).Minus(_18_v0))), 16)
		(_nw31).ArraySet1(_18_v0, 17)
		(_nw31).ArraySet1(_18_v0, 18)
		(_nw31).ArraySet1(_dafny.IntOfInt64(681), 19)
		(_nw31).ArraySet1(_18_v0, 20)
		(_nw31).ArraySet1(_18_v0, 21)
		(_nw31).ArraySet1(_18_v0, 22)
		(_nw31).ArraySet1(Companion_Default___.Fm0(_18_v0, false, globalState), 23)
		_187_v136 = _nw31
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))
		_ = _index26
		(_187_v136).ArraySet1((func() _dafny.Int {
			if p1 {
				return _18_v0
			}
			return _dafny.IntOfInt64(592)
		})(), (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))
		_ = _index27
		(_187_v136).ArraySet1(_18_v0, (_index27).Int())
		(globalState).F4 = _dafny.Companion_Sequence_.IsProperPrefixOf(p0, p0)
		var _188_v137 _dafny.MultiSet
		_ = _188_v137
		_188_v137 = _dafny.MultiSetOf(_18_v0)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))
		_ = _index28
		(_187_v136).ArraySet1((func() _dafny.Int {
			if p1 {
				return Companion_Default___.SafeModuloInt((_188_v137).Cardinality(), _18_v0)
			}
			return (_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int)
		})(), (_index28).Int())
		var _189_v138 _dafny.Set
		_ = _189_v138
		_189_v138 = _dafny.SetOf(!(p1))
		var _pat_let_tv8 = _189_v138
		_ = _pat_let_tv8
		var _source4 D1 = func(_pat_let12_0 D1) D1 {
			return func(_190_dt__update__tmp_h6 D1) D1 {
				return func(_pat_let13_0 _dafny.Set) D1 {
					return func(_191_dt__update_hcf12_h1 _dafny.Set) D1 {
						return Companion_D1_.Create_DC6_(_191_dt__update_hcf12_h1)
					}(_pat_let13_0)
				}(_pat_let_tv8)
			}(_pat_let12_0)
		}(Companion_Default___.Fm10(p1, (_dafny.Zero).Minus((_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int)), globalState))
		_ = _source4
		if _source4.Is_DC5() {
			var _192___mcc_h9 _dafny.Int = _source4.Get_().(D1_DC5).Cf11
			_ = _192___mcc_h9
			var _193_cf11 _dafny.Int = _192___mcc_h9
			_ = _193_cf11
			var _194_v139 *C0
			_ = _194_v139
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__((_188_v137).Cardinality(), p1)
			_194_v139 = _nw32
			var _195_v140 T0
			_ = _195_v140
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__((_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int), (func() bool {
				if p1 {
					return !(!(p1))
				}
				return p1
			})())
			_195_v140 = _nw33
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_19_v1).Cardinality()), (_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int)), p1)
			_195_v140 = _nw34
			var _196_v141 _dafny.Map
			_ = _196_v141
			_196_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_195_v140).F19(), _18_v0)
			(globalState).F16 = (_dafny.IntOfInt64(633)).Times((func() _dafny.Int {
				if (_196_v141).Contains(p1) {
					return (_196_v141).Get(p1).(_dafny.Int)
				}
				return (_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int)
			})())
			var _197_v142 _dafny.Map
			_ = _197_v142
			_197_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_195_v140.F18()), _195_v140.F18())
			var _198_v143 _dafny.MultiSet
			_ = _198_v143
			_198_v143 = _dafny.MultiSetOf(_188_v137, Companion_Default___.Fm11(_197_v142, p1, _189_v138, globalState), (_188_v137).Update(_18_v0, Companion_Default___.Abs(_18_v0)), _188_v137)
			var _199_v144 _dafny.Sequence
			_ = _199_v144
			_199_v144 = _dafny.SeqOf(_188_v137)
			_198_v143 = (_198_v143).Union(_dafny.MultiSetFromSeq(_199_v144))
		} else if _source4.Is_DC6() {
			var _200___mcc_h10 _dafny.Set = _source4.Get_().(D1_DC6).Cf12
			_ = _200___mcc_h10
			var _201_cf12 _dafny.Set = _200___mcc_h10
			_ = _201_cf12
			var _202_v146 _dafny.Array
			_ = _202_v146
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_0
			var _nw35 _dafny.Array
			_ = _nw35
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw35 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Set = (func(_203_v0 _dafny.Int, _204_p0 _dafny.Sequence, _205_v136 _dafny.Array, _206_p1 bool) func(_dafny.Int) _dafny.Set {
					return func(_207_i2 _dafny.Int) _dafny.Set {
						return (func() _dafny.Set {
							var _coll1 = _dafny.NewBuilder()
							_ = _coll1
							for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_204_p0).Cardinality()), (_205_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_205_v136), 0))).Int()).(_dafny.Int)), _206_p1)).Keys().Elements()); ; {
								_compr_1, _ok1 := _iter1()
								if !_ok1 {
									break
								}
								var _208_v145 _dafny.Map
								_208_v145 = interface{}(_compr_1).(_dafny.Map)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_204_p0).Cardinality()), (_205_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_205_v136), 0))).Int()).(_dafny.Int)), _206_p1)).Contains(_208_v145) {
									_coll1.Add(_208_v145)
								}
							}
							return _coll1.ToSet()
						}()).Union(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(165), _203_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_203_v0, _dafny.IntOfUint32((_dafny.SeqOf(_206_p1)).Cardinality()))))
					}
				})(_18_v0, p0, _187_v136, p1)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw35 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw35).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw35).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_202_v146 = _nw35
			var _209_v147 _dafny.Map
			_ = _209_v147
			_209_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int), _18_v0)
			var _210_v148 _dafny.Set
			_ = _210_v148
			_210_v148 = _dafny.SetOf(_209_v147, _209_v147)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_202_v146), 0))
			_ = _index29
			(_202_v146).ArraySet1(_210_v148, (_index29).Int())
			var _211_v149 _dafny.Sequence
			_ = _211_v149
			_211_v149 = _dafny.SeqOf(_209_v147, _209_v147)
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_202_v146), 0))
			_ = _index30
			(_202_v146).ArraySet1((_210_v148).Union(func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_211_v149).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _212_v150 _dafny.Map
					_212_v150 = interface{}(_compr_2).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_211_v149, _212_v150) {
						_coll2.Add(_212_v150)
					}
				}
				return _coll2.ToSet()
			}()), (_index30).Int())
			var _213_v151 _dafny.Map
			_ = _213_v151
			_213_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_18_v0), p1)
			_213_v151 = (_213_v151).Merge(_213_v151)
			var _214_v152 _dafny.Array
			_ = _214_v152
			var _nw36 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
			_ = _nw36
			_214_v152 = _nw36
			var _215_v153 _dafny.Map
			_ = _215_v153
			_215_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, _214_v152)
			_215_v153 = (_215_v153).Update(((func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate((func() _dafny.Set {
					var _coll4 = _dafny.NewBuilder()
					_ = _coll4
					for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-154), _dafny.IntOfInt64(-722))); ; {
						_compr_4, _ok4 := _iter4()
						if !_ok4 {
							break
						}
						var _216_v155 _dafny.Int
						_216_v155 = interface{}(_compr_4).(_dafny.Int)
						if ((_dafny.IntOfInt64(-154)).Cmp(_216_v155) <= 0) && ((_216_v155).Cmp(_dafny.IntOfInt64(-722)) < 0) {
							_coll4.Add((_216_v155).Times(_18_v0))
						}
					}
					return _coll4.ToSet()
				}()).Elements()); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _217_v154 _dafny.Int
					_217_v154 = interface{}(_compr_3).(_dafny.Int)
					if (func() _dafny.Set {
						var _coll5 = _dafny.NewBuilder()
						_ = _coll5
						for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-154), _dafny.IntOfInt64(-722))); ; {
							_compr_5, _ok5 := _iter5()
							if !_ok5 {
								break
							}
							var _218_v155 _dafny.Int
							_218_v155 = interface{}(_compr_5).(_dafny.Int)
							if ((_dafny.IntOfInt64(-154)).Cmp(_218_v155) <= 0) && ((_218_v155).Cmp(_dafny.IntOfInt64(-722)) < 0) {
								_coll5.Add((_218_v155).Times(_18_v0))
							}
						}
						return _coll5.ToSet()
					}()).Contains(_217_v154) {
						_coll3.Add((_217_v154).Minus(_18_v0), _dafny.CodePoint('g'))
					}
				}
				return _coll3.ToMap()
			}()).Cardinality()).Plus((_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int)), _214_v152)
			(globalState).F4 = !(p1)
		} else if _source4.Is_DC7() {
			var _219___mcc_h11 _dafny.Int = _source4.Get_().(D1_DC7).Cf13
			_ = _219___mcc_h11
			var _220___mcc_h12 _dafny.Int = _source4.Get_().(D1_DC7).Cf14
			_ = _220___mcc_h12
			var _221_cf14 _dafny.Int = _220___mcc_h12
			_ = _221_cf14
			var _222_cf13 _dafny.Int = _219___mcc_h11
			_ = _222_cf13
			var _223_v156 _dafny.Array
			_ = _223_v156
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_1
			var _nw37 _dafny.Array
			_ = _nw37
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw37 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = func(_224_i3 _dafny.Int) bool {
					return false
				}
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw37 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw37).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw37).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_223_v156 = _nw37
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))
			_ = _index31
			(_223_v156).ArraySet1(p1, (_index31).Int())
			var _225_v157 _dafny.Map
			_ = _225_v157
			_225_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !_dafny.Companion_Sequence_.Contains(_19_v1, (_19_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_221_cf14), _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int)))
			var _226_v158 *C0
			_ = _226_v158
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__(_221_cf14, p1)
			_226_v158 = _nw38
			var _227_v159 D0
			_ = _227_v159
			_227_v159 = Companion_D0_.Create_DC1_(p1, _225_v157, _226_v158, (_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int))
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))
			_ = _index32
			(_223_v156).ArraySet1(!(!((func() bool {
				if (_225_v157).Contains((_227_v159).Dtor_cf2()) {
					return (_225_v157).Get((_227_v159).Dtor_cf2()).(bool)
				}
				return p1
			})())), (_index32).Int())
			var _228_v160 D3
			_ = _228_v160
			_228_v160 = Companion_D3_.Create_DC12_(_223_v156)
			var _229_v161 _dafny.Sequence
			_ = _229_v161
			_229_v161 = _dafny.SeqOf(_228_v160)
			var _230_v162 _dafny.Map
			_ = _230_v162
			_230_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if Companion_Default___.Fm3((_223_v156).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))).Int()).(bool), _222_cf13, (_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int), false, globalState) {
					return (_223_v156).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))).Int()).(bool)
				}
				return (_223_v156).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))).Int()).(bool)
			})(), _dafny.Companion_Sequence_.Update(_229_v161, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_229_v161).Cardinality()))).Uint32(), _228_v160))
			_230_v162 = (_230_v162).Update((_223_v156).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))).Int()).(bool), _229_v161)
			var _231_v163 D1
			_ = _231_v163
			_231_v163 = Companion_D1_.Create_DC4_(_187_v136)
			var _232_v164 _dafny.Sequence
			_ = _232_v164
			_232_v164 = _dafny.SeqOf((_223_v156).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))).Int()).(bool), (_223_v156).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))).Int()).(bool))
			var _rhs24 D1 = _231_v163
			_ = _rhs24
			var _rhs25 *C0 = _226_v158
			_ = _rhs25
			var _rhs26 _dafny.Sequence = (func() _dafny.Sequence {
				if (_223_v156).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_223_v156), 0))).Int()).(bool) {
					return _232_v164
				}
				return _232_v164
			})()
			_ = _rhs26
			var _lhs21 *GlobalState = globalState
			_ = _lhs21
			_231_v163 = _rhs24
			_226_v158 = _rhs25
			_lhs21.F6 = _rhs26
			_221_cf14 = (_19_v1).Select((Companion_Default___.SafeIndex(_221_cf14, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int)
		} else if _source4.Is_DC4() {
			var _233___mcc_h13 _dafny.Array = _source4.Get_().(D1_DC4).Cf10
			_ = _233___mcc_h13
			var _234_cf10 _dafny.Array = _233___mcc_h13
			_ = _234_cf10
			(globalState).F4 = p1
			(globalState).F16 = (_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int)
			(globalState).F4 = ((_18_v0).Times(_18_v0)).Cmp((_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int)) < 0
			var _235_v165 *C0
			_ = _235_v165
			var _nw39 *C0 = New_C0_()
			_ = _nw39
			_nw39.Ctor__((_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int), p1)
			_235_v165 = _nw39
			_235_v165 = _235_v165
		} else {
			var _236___mcc_h14 D1 = _source4.Get_().(D1_DC8).Cf15
			_ = _236___mcc_h14
			var _237_cf15 D1 = _236___mcc_h14
			_ = _237_cf15
			var _238_v166 _dafny.Array
			_ = _238_v166
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.One)
			_ = _nw40
			_238_v166 = _nw40
			var _239_v167 *C0
			_ = _239_v167
			var _nw41 *C0 = New_C0_()
			_ = _nw41
			_nw41.Ctor__(_dafny.IntOfInt64(2), true)
			_239_v167 = _nw41
			var _240_v168 _dafny.Set
			_ = _240_v168
			_240_v168 = _dafny.SetOf(_239_v167, _239_v167)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_238_v166), 0))
			_ = _index33
			(_238_v166).ArraySet1(_240_v168, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_238_v166), 0))
			_ = _index34
			(_238_v166).ArraySet1((func() _dafny.Set {
				if p1 {
					return _240_v168
				}
				return _240_v168
			})(), (_index34).Int())
			_18_v0 = _dafny.IntOfInt64(99)
			var _241_v169 _dafny.Array
			_ = _241_v169
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw42
			_241_v169 = _nw42
			var _242_v170 _dafny.Map
			_ = _242_v170
			_242_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v169, p1)
			_242_v170 = (_242_v170).Update(_241_v169, p1)
			var _243_v171 _dafny.Set
			_ = _243_v171
			_243_v171 = _dafny.SetOf(_dafny.IntOfUint32((p0).Cardinality()))
			(globalState).F4 = Companion_Default___.Fm3(p1, Companion_Default___.SafeModuloInt(_18_v0, _dafny.IntOfUint32((p0).Cardinality())), (_243_v171).Cardinality(), p1, globalState)
		}
		var _244_v172 _dafny.Array
		_ = _244_v172
		var _nwElement0_13 bool = p1
		_ = _nwElement0_13
		var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.One)
		_ = _nw43
		(_nw43).ArraySet1(_nwElement0_13, 0)
		_244_v172 = _nw43
		var _245_v173 _dafny.Sequence
		_ = _245_v173
		_245_v173 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("rlhubva"))
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_244_v172), 0))
		_ = _index35
		(_244_v172).ArraySet1(_dafny.Companion_Sequence_.Contains(_245_v173, p0), (_index35).Int())
		var _246_v174 _dafny.Set
		_ = _246_v174
		_246_v174 = _dafny.SetOf((_187_v136).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_187_v136), 0))).Int()).(_dafny.Int))
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_244_v172), 0))
		_ = _index36
		(_244_v172).ArraySet1(((_246_v174).Union(_246_v174)).IsDisjointFrom(_246_v174), (_index36).Int())
	}
	(globalState).F4 = p1
	var _247_v175 _dafny.Sequence
	_ = _247_v175
	_247_v175 = _dafny.SeqOf(p1, !(p1), p1)
	var _hi0 _dafny.Int = _18_v0
	_ = _hi0
	for _248_i4 := Companion_Default___.Fm0(_dafny.IntOfUint32((_247_v175).Cardinality()), p1, globalState); _248_i4.Cmp(_hi0) < 0; _248_i4 = _248_i4.Plus(_dafny.One) {
		var _249_v176 *C0
		_ = _249_v176
		var _nw44 *C0 = New_C0_()
		_ = _nw44
		_nw44.Ctor__(_248_i4, p1)
		_249_v176 = _nw44
		_249_v176 = _249_v176
		r0 = _248_i4
		var _250_v177 _dafny.Array
		_ = _250_v177
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_2
		var _nw45 _dafny.Array
		_ = _nw45
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw45 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_251_p1 bool) func(_dafny.Int) bool {
				return func(_252_i5 _dafny.Int) bool {
					return (false) || (_251_p1)
				}
			})(p1)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw45 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw45).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw45).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_250_v177 = _nw45
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))
		_ = _index37
		(_250_v177).ArraySet1(p1, (_index37).Int())
		var _253_v178 D0
		_ = _253_v178
		_253_v178 = Companion_D0_.Create_DC1_(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), _249_v176, _248_i4)
		var _254_v179 _dafny.MultiSet
		_ = _254_v179
		_254_v179 = _dafny.MultiSetOf(p1)
		var _255_v180 _dafny.Map
		_ = _255_v180
		_255_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v178, (_254_v179).Cardinality())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))
		_ = _index38
		(_250_v177).ArraySet1(((_19_v1).Select((Companion_Default___.SafeIndex(((_255_v180).Update(_253_v178, _248_i4)).Cardinality(), _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_248_i4).Minus(_dafny.IntOfInt64(258))) > 0, (_index38).Int())
		if (_250_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))).Int()).(bool) {
			var _256_v181 *C0
			_ = _256_v181
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__(_18_v0, p1)
			_256_v181 = _nw46
			(globalState).F16 = _18_v0
			var _257_v182 _dafny.CodePoint
			_ = _257_v182
			_257_v182 = _dafny.CodePoint('f')
			var _258_v183 _dafny.Array
			_ = _258_v183
			var _nwElement0_14 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p0, p0)
			_ = _nwElement0_14
			var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(25))
			_ = _nw47
			(_nw47).ArraySet1(_nwElement0_14, 0)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(147))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_259_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))), 1)
			(_nw47).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(290))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_260_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			})), 2)
			(_nw47).ArraySet1(p0, 3)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _257_v182)), 4)
			(_nw47).ArraySet1(Companion_Default___.Fm2(true, p1, globalState), 5)
			(_nw47).ArraySet1(Companion_Default___.Fm2(p1, (_250_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))).Int()).(bool), globalState), 6)
			(_nw47).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(106))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}((func(_261_v182 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_262_i8 _dafny.Int) _dafny.CodePoint {
					return _261_v182
				}
			})(_257_v182))), 7)
			(_nw47).ArraySet1(p0, 8)
			(_nw47).ArraySet1(p0, 9)
			(_nw47).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lo"), 10)
			(_nw47).ArraySet1(p0, 11)
			(_nw47).ArraySet1(p0, 12)
			(_nw47).ArraySet1(p0, 13)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pyu"), p0), 14)
			(_nw47).ArraySet1(p0, 15)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), (Companion_Default___.SafeIndex(_248_i4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))).Uint32(), _257_v182), 16)
			(_nw47).ArraySet1(p0, 17)
			(_nw47).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ujrexjxx"), 18)
			(_nw47).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-616))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_263_v182 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_264_i9 _dafny.Int) _dafny.CodePoint {
					return _263_v182
				}
			})(_257_v182))), 19)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ekt"), (Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekt")).Cardinality()))).Uint32(), _257_v182), p0), 20)
			(_nw47).ArraySet1(p0, 21)
			(_nw47).ArraySet1(Companion_Default___.Fm2(p1, (_250_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))).Int()).(bool), globalState), 22)
			(_nw47).ArraySet1(Companion_Default___.Fm2(p1, (_250_v177).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))).Int()).(bool), globalState), 23)
			(_nw47).ArraySet1(p0, 24)
			_258_v183 = _nw47
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_258_v183), 0))
			_ = _index39
			(_258_v183).ArraySet1(p0, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_258_v183), 0))
			_ = _index40
			(_258_v183).ArraySet1(p0, (_index40).Int())
			(globalState).F4 = p1
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_258_v183), 0))
			_ = _index41
			var _rhs27 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(p0, (_258_v183).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_258_v183), 0))).Int()).(_dafny.Sequence)), _257_v182)
			_ = _rhs27
			var _rhs28 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_265_v182 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_266_i10 _dafny.Int) _dafny.CodePoint {
					return _265_v182
				}
			})(_257_v182))), _dafny.Companion_Sequence_.Concatenate((_258_v183).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_258_v183), 0))).Int()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_267_v182 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_268_i11 _dafny.Int) _dafny.CodePoint {
					return _267_v182
				}
			})(_257_v182)))))
			_ = _rhs28
			var _rhs29 _dafny.CodePoint = Companion_Default___.Fm7(globalState)
			_ = _rhs29
			var _rhs30 _dafny.Int = _248_i4
			_ = _rhs30
			var _lhs22 *GlobalState = globalState
			_ = _lhs22
			var _lhs23 _dafny.Array = _258_v183
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_258_v183), 0))
			_ = _lhs24
			var _lhs25 *GlobalState = globalState
			_ = _lhs25
			_lhs22.F4 = _rhs27
			(_lhs23).ArraySet1(_rhs28, (_lhs24).Int())
			_257_v182 = _rhs29
			_lhs25.F16 = _rhs30
		} else {
			(globalState).F4 = true
			var _269_v184 _dafny.Sequence
			_ = _269_v184
			_269_v184 = _dafny.SeqOf(_19_v1)
			var _270_v185 _dafny.Array
			_ = _270_v185
			var _nwElement0_15 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_19_v1, (Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()))
			_ = _nwElement0_15
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(12))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_15, 0)
			(_nw48).ArraySet1(_19_v1, 1)
			(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(331), _248_i4), _19_v1), 2)
			(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_19_v1, _19_v1), 3)
			(_nw48).ArraySet1(_19_v1, 4)
			(_nw48).ArraySet1(_dafny.SeqOf(_18_v0), 5)
			(_nw48).ArraySet1(_dafny.Companion_Sequence_.Update((_269_v184).Select((Companion_Default___.SafeIndex((_19_v1).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_269_v184).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_248_i4, _dafny.IntOfUint32(((_269_v184).Select((Companion_Default___.SafeIndex((_19_v1).Select((Companion_Default___.SafeIndex(_18_v0, _dafny.IntOfUint32((_19_v1).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_269_v184).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _18_v0), 6)
			(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_19_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_271_v179 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
				return func(_272_i12 _dafny.Int) _dafny.Int {
					return (_271_v179).Cardinality()
				}
			})(_254_v179)))), 7)
			(_nw48).ArraySet1(_19_v1, 8)
			(_nw48).ArraySet1(_19_v1, 9)
			(_nw48).ArraySet1(_dafny.SeqOf((_254_v179).Cardinality(), _18_v0), 10)
			(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_19_v1, _19_v1), 11)
			_270_v185 = _nw48
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_270_v185), 0))
			_ = _index42
			(_270_v185).ArraySet1(_19_v1, (_index42).Int())
			var _273_v186 _dafny.MultiSet
			_ = _273_v186
			_273_v186 = _dafny.MultiSetOf(_249_v176)
			var _274_v187 _dafny.MultiSet
			_ = _274_v187
			_274_v187 = _dafny.MultiSetOf(_273_v186, _273_v186, _273_v186)
			var _275_v188 _dafny.Sequence
			_ = _275_v188
			_275_v188 = _dafny.UnicodeSeqOfUtf8Bytes("lhknlfumr")
			var _276_v189 _dafny.Set
			_ = _276_v189
			_276_v189 = _dafny.SetOf(p1)
			var _277_v190 _dafny.Map
			_ = _277_v190
			_277_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, _248_i4)
			var _278_v191 _dafny.Map
			_ = _278_v191
			_278_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v190, _276_v189)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_270_v185), 0))
			_ = _index43
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))
			_ = _index44
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))
			_ = _index45
			var _rhs31 _dafny.Sequence = _19_v1
			_ = _rhs31
			var _rhs32 bool = p1
			_ = _rhs32
			var _rhs33 bool = (_276_v189).IsProperSubsetOf((_276_v189).Intersection((func() _dafny.Set {
				if (_278_v191).Contains(_277_v190) {
					return (_278_v191).Get(_277_v190).(_dafny.Set)
				}
				return _276_v189
			})()))
			_ = _rhs33
			var _rhs34 _dafny.MultiSet = _274_v187
			_ = _rhs34
			var _rhs35 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_275_v188, _dafny.UnicodeSeqOfUtf8Bytes("jduntgx"))
			_ = _rhs35
			var _lhs26 _dafny.Array = _270_v185
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_270_v185), 0))
			_ = _lhs27
			var _lhs28 _dafny.Array = _250_v177
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))
			_ = _lhs29
			var _lhs30 _dafny.Array = _250_v177
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_250_v177), 0))
			_ = _lhs31
			(_lhs26).ArraySet1(_rhs31, (_lhs27).Int())
			(_lhs28).ArraySet1(_rhs32, (_lhs29).Int())
			(_lhs30).ArraySet1(_rhs33, (_lhs31).Int())
			_274_v187 = _rhs34
			_275_v188 = _rhs35
			var _279_v192 *C0
			_ = _279_v192
			var _nw49 *C0 = New_C0_()
			_ = _nw49
			_nw49.Ctor__((_276_v189).Cardinality(), false)
			_279_v192 = _nw49
			(globalState).F16 = ((_276_v189).Intersection((_276_v189).Union(_276_v189))).Cardinality()
			(globalState).F16 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_18_v0, _248_i4), Companion_Default___.SafeDivisionInt((_277_v190).Cardinality(), _18_v0))
		}
	}
	r0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(859), _dafny.IntOfInt64(457))
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _280_v0 bool
	_ = _280_v0
	_280_v0 = true
	var _281_v1 _dafny.Sequence
	_ = _281_v1
	_281_v1 = _dafny.SeqOf(_280_v0, _280_v0, _280_v0)
	var _282_v2 _dafny.Array
	_ = _282_v2
	var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
	_ = _nw50
	_282_v2 = _nw50
	var _283_v3 _dafny.Int
	_ = _283_v3
	_283_v3 = _dafny.IntOfInt64(-235)
	var _284_v4 _dafny.Map
	_ = _284_v4
	_284_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _283_v3)
	var _285_v5 _dafny.Sequence
	_ = _285_v5
	_285_v5 = _dafny.SeqOf(_282_v2)
	var _286_v6 _dafny.Array
	_ = _286_v6
	var _nwElement0_16 bool = !(_280_v0)
	_ = _nwElement0_16
	var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(8))
	_ = _nw51
	(_nw51).ArraySet1(_nwElement0_16, 0)
	(_nw51).ArraySet1(!(_280_v0), 1)
	(_nw51).ArraySet1(_280_v0, 2)
	(_nw51).ArraySet1(_280_v0, 3)
	(_nw51).ArraySet1(true, 4)
	(_nw51).ArraySet1(_280_v0, 5)
	(_nw51).ArraySet1(_280_v0, 6)
	(_nw51).ArraySet1(_280_v0, 7)
	_286_v6 = _nw51
	var _287_v7 _dafny.MultiSet
	_ = _287_v7
	_287_v7 = _dafny.MultiSetOf(_286_v6)
	var _288_globalState *GlobalState
	_ = _288_globalState
	var _nw52 *GlobalState = New_GlobalState_()
	_ = _nw52
	_nw52.Ctor__(_dafny.IntOfInt64(485), _dafny.IntOfInt64(-62), _dafny.IntOfInt64(525), false, false, _dafny.IntOfInt64(646), _281_v1, _dafny.IntOfInt64(462), _282_v2, _dafny.IntOfInt64(789), _dafny.IntOfInt64(610), _284_v4, false, (_285_v5).Select((Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_285_v5).Cardinality()))).Uint32()).(_dafny.Array), _286_v6, (_287_v7).Difference(_287_v7), _dafny.IntOfInt64(789), _dafny.CodePoint('d'))
	_288_globalState = _nw52
	var _289_v8 _dafny.Map
	_ = _289_v8
	_289_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _280_v0)
	var _290_v9 _dafny.Set
	_ = _290_v9
	_290_v9 = _dafny.SetOf(_280_v0, (func() bool {
		if (_289_v8).Contains(_280_v0) {
			return (_289_v8).Get(_280_v0).(bool)
		}
		return false
	})())
	var _291_i0 _dafny.Int
	_ = _291_i0
	_291_i0 = _dafny.Zero
	{
		for (_281_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_290_v9).Difference(_290_v9)).Cardinality()), _dafny.IntOfUint32((_281_v1).Cardinality()))).Uint32()).(bool) {
			{
				if (_291_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_291_i0 = (_291_i0).Plus(_dafny.One)
				var _292_v10 _dafny.Array
				_ = _292_v10
				var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
				_ = _nw53
				_292_v10 = _nw53
				var _293_v11 _dafny.CodePoint
				_ = _293_v11
				_293_v11 = _dafny.CodePoint('n')
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_292_v10), 0))
				_ = _index46
				(_292_v10).ArraySet1CodePoint(_293_v11, (_index46).Int())
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_292_v10), 0))
				_ = _index47
				(_292_v10).ArraySet1CodePoint(_293_v11, (_index47).Int())
				var _294_v12 _dafny.Sequence
				_ = _294_v12
				_294_v12 = _dafny.UnicodeSeqOfUtf8Bytes("jx")
				var _295_v13 _dafny.Int
				_ = _295_v13
				var _out0 _dafny.Int
				_ = _out0
				_out0 = Companion_Default___.M0(_294_v12, !(!(!(_280_v0))), _288_globalState)
				_295_v13 = _out0
				var _296_v14 _dafny.Int
				_ = _296_v14
				var _out1 _dafny.Int
				_ = _out1
				_out1 = Companion_Default___.M0(_294_v12, _280_v0, _288_globalState)
				_296_v14 = _out1
				(_288_globalState).F10 = _296_v14
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
	_ = _index48
	(_286_v6).ArraySet1(!(!(_280_v0)), (_index48).Int())
	var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
	_ = _index49
	(_286_v6).ArraySet1(!((func() bool {
		if _280_v0 {
			return _280_v0
		}
		return _280_v0
	})()), (_index49).Int())
	if (func() bool {
		if (_289_v8).Contains(false) {
			return (_289_v8).Get(false).(bool)
		}
		return (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
	})() {
		var _297_v15 _dafny.Sequence
		_ = _297_v15
		_297_v15 = _dafny.UnicodeSeqOfUtf8Bytes("p")
		(_288_globalState).F16 = _dafny.IntOfUint32((_297_v15).Cardinality())
		var _298_v16 _dafny.Int
		_ = _298_v16
		var _out2 _dafny.Int
		_ = _out2
		_out2 = Companion_Default___.M0(_297_v15, !((func() bool {
			if false {
				return (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
			}
			return _280_v0
		})()), _288_globalState)
		_298_v16 = _out2
		var _299_v17 _dafny.Map
		_ = _299_v17
		_299_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _281_v1)
		var _300_v18 _dafny.Sequence
		_ = _300_v18
		_300_v18 = _dafny.SeqOf(_299_v17, _299_v17)
		var _301_v19 _dafny.Map
		_ = _301_v19
		_301_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _dafny.SeqOf(true, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))))
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _index50
		(_286_v6).ArraySet1(((_299_v17).Merge(_299_v17)).Equals((((_300_v18).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_298_v16), _dafny.IntOfUint32((_300_v18).Cardinality()))).Uint32()).(_dafny.Map)).Update(_280_v0, _281_v1)).Merge((func() _dafny.Map {
			if (_301_v19).Contains(Companion_Default___.Fm0(_298_v16, _280_v0, _288_globalState)) {
				return (_301_v19).Get(Companion_Default___.Fm0(_298_v16, _280_v0, _288_globalState)).(_dafny.Map)
			}
			return _299_v17
		})())), (_index50).Int())
		var _302_v20 _dafny.CodePoint
		_ = _302_v20
		_302_v20 = _dafny.CodePoint('b')
		var _303_v21 _dafny.Map
		_ = _303_v21
		_303_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_302_v20, _288_globalState), ((_dafny.Zero).Minus(_283_v3)).Cmp(_298_v16) < 0)
		var _304_v22 _dafny.MultiSet
		_ = _304_v22
		_304_v22 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _302_v20)).Cardinality())
		var _305_v23 _dafny.Map
		_ = _305_v23
		_305_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-695), _304_v22)
		_280_v0 = (func() bool {
			if (_303_v21).Contains(_305_v23) {
				return (_303_v21).Get(_305_v23).(bool)
			}
			return (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
		})()
		var _306_v24 _dafny.Array
		_ = _306_v24
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_3
		var _nw54 _dafny.Array
		_ = _nw54
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw54 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = (func(_307_v0 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_308_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_307_v0)
				}
			})(_280_v0)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw54 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw54).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw54).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_306_v24 = _nw54
		var _309_v27 _dafny.Array
		_ = _309_v27
		var _len0_4 _dafny.Int = _dafny.One
		_ = _len0_4
		var _nw55 _dafny.Array
		_ = _nw55
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw55 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Map = (func(_310_v3 _dafny.Int, _311_v9 _dafny.Set, _312_v16 _dafny.Int, _313_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
				return func(_314_i2 _dafny.Int) _dafny.Map {
					return (func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(175), _dafny.IntOfInt64(-248))); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _315_v25 _dafny.Int
							_315_v25 = interface{}(_compr_6).(_dafny.Int)
							if ((_dafny.IntOfInt64(175)).Cmp(_315_v25) <= 0) && ((_315_v25).Cmp(_dafny.IntOfInt64(-248)) < 0) {
								_coll6.Add((_315_v25).Plus((func() _dafny.Map {
									var _coll7 = _dafny.NewMapBuilder()
									_ = _coll7
									for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(149), _dafny.IntOfInt64(819))); ; {
										_compr_7, _ok7 := _iter7()
										if !_ok7 {
											break
										}
										var _316_v26 _dafny.Int
										_316_v26 = interface{}(_compr_7).(_dafny.Int)
										if ((_dafny.IntOfInt64(149)).Cmp(_316_v26) <= 0) && ((_316_v26).Cmp(_dafny.IntOfInt64(819)) < 0) {
											_coll7.Add(Companion_Default___.SafeDivisionInt(_316_v26, _312_v16), _313_v20)
										}
									}
									return _coll7.ToMap()
								}()).Cardinality()), _311_v9)
							}
						}
						return _coll6.ToMap()
					}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_310_v3, _311_v9))
				}
			})(_283_v3, _290_v9, _298_v16, _302_v20)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw55 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw55).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw55).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_309_v27 = _nw55
		var _317_v29 _dafny.Set
		_ = _317_v29
		_317_v29 = _dafny.SetOf(_298_v16)
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_309_v27), 0))
		_ = _index51
		(_309_v27).ArraySet1(func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_317_v29).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _318_v28 _dafny.Int
				_318_v28 = interface{}(_compr_8).(_dafny.Int)
				if (_317_v29).Contains(_318_v28) {
					_coll8.Add(Companion_Default___.SafeModuloInt(_318_v28, (_dafny.Zero).Minus(_298_v16)), _dafny.SetOf(_280_v0))
				}
			}
			return _coll8.ToMap()
		}(), (_index51).Int())
		var _319_v31 _dafny.Map
		_ = _319_v31
		_319_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_298_v16, _290_v9)
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _index52
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_309_v27), 0))
		_ = _index53
		var _rhs36 _dafny.CodePoint = _302_v20
		_ = _rhs36
		var _rhs37 bool = ((func() _dafny.Set {
			if (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool) {
				return _317_v29
			}
			return _317_v29
		})()).IsSubsetOf(func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(182), _dafny.IntOfInt64(539))); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _320_v30 _dafny.Int
				_320_v30 = interface{}(_compr_9).(_dafny.Int)
				if ((_dafny.IntOfInt64(182)).Cmp(_320_v30) <= 0) && ((_320_v30).Cmp(_dafny.IntOfInt64(539)) < 0) {
					_coll9.Add(Companion_Default___.SafeDivisionInt(_320_v30, _298_v16))
				}
			}
			return _coll9.ToSet()
		}())
		_ = _rhs37
		var _rhs38 _dafny.Array = _306_v24
		_ = _rhs38
		var _rhs39 _dafny.Map = (_319_v31).Merge(_319_v31)
		_ = _rhs39
		var _rhs40 _dafny.Array = _282_v2
		_ = _rhs40
		var _lhs32 _dafny.Array = _286_v6
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _lhs33
		var _lhs34 _dafny.Array = _309_v27
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_309_v27), 0))
		_ = _lhs35
		_302_v20 = _rhs36
		(_lhs32).ArraySet1(_rhs37, (_lhs33).Int())
		_306_v24 = _rhs38
		(_lhs34).ArraySet1(_rhs39, (_lhs35).Int())
		_282_v2 = _rhs40
	} else {
		var _321_v32 _dafny.Sequence
		_ = _321_v32
		_321_v32 = _dafny.UnicodeSeqOfUtf8Bytes("tkkvwrsmk")
		var _322_v33 _dafny.Set
		_ = _322_v33
		_322_v33 = _dafny.SetOf(_284_v4)
		var _323_v34 _dafny.Int
		_ = _323_v34
		var _out3 _dafny.Int
		_ = _out3
		_out3 = Companion_Default___.M0(_321_v32, (_322_v33).IsProperSubsetOf(_322_v33), _288_globalState)
		_323_v34 = _out3
		_280_v0 = true
		_321_v32 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(176))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_324_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), Companion_Default___.Fm2((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _280_v0, _288_globalState)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-631))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_325_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})))
		_283_v3 = (_dafny.IntOfUint32((Companion_Default___.Fm2(_280_v0, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState)).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(62))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_326_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality()))
		(_288_globalState).F4 = Companion_Default___.Fm3(_280_v0, _323_v34, _323_v34, _280_v0, _288_globalState)
	}
	var _327_v35 _dafny.Sequence
	_ = _327_v35
	_327_v35 = _dafny.UnicodeSeqOfUtf8Bytes("rgyxvjy")
	var _328_v36 _dafny.Int
	_ = _328_v36
	var _out4 _dafny.Int
	_ = _out4
	_out4 = Companion_Default___.M0(_327_v35, (_281_v1).Select((Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_281_v1).Cardinality()))).Uint32()).(bool), _288_globalState)
	_328_v36 = _out4
	var _329_v37 _dafny.Sequence
	_ = _329_v37
	_329_v37 = _dafny.SeqOf(_328_v36, Companion_Default___.Fm0(_328_v36, !(false), _288_globalState), _dafny.IntOfUint32((_327_v35).Cardinality()), _dafny.IntOfInt64(135))
	var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
	_ = _index54
	var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
	_ = _index55
	var _rhs41 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_327_v35, _dafny.UnicodeSeqOfUtf8Bytes("ryj"))
	_ = _rhs41
	var _rhs42 _dafny.Array = _286_v6
	_ = _rhs42
	var _rhs43 _dafny.Int = _328_v36
	_ = _rhs43
	var _rhs44 bool = (Companion_Default___.Fm3(_280_v0, _283_v3, (_dafny.Zero).Minus((_dafny.Zero).Minus(_283_v3)), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState)) || (_dafny.Companion_Sequence_.Contains(_329_v37, Companion_Default___.Fm0(_328_v36, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState)))
	_ = _rhs44
	var _rhs45 bool = (_281_v1).Select((Companion_Default___.SafeIndex(_328_v36, _dafny.IntOfUint32((_281_v1).Cardinality()))).Uint32()).(bool)
	_ = _rhs45
	var _lhs36 *GlobalState = _288_globalState
	_ = _lhs36
	var _lhs37 _dafny.Array = _286_v6
	_ = _lhs37
	var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
	_ = _lhs38
	var _lhs39 _dafny.Array = _286_v6
	_ = _lhs39
	var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
	_ = _lhs40
	_327_v35 = _rhs41
	_286_v6 = _rhs42
	_lhs36.F0 = _rhs43
	(_lhs37).ArraySet1(_rhs44, (_lhs38).Int())
	(_lhs39).ArraySet1(_rhs45, (_lhs40).Int())
	if _dafny.Companion_Sequence_.IsPrefixOf(_327_v35, _327_v35) {
		var _330_v38 *C0
		_ = _330_v38
		var _nw56 *C0 = New_C0_()
		_ = _nw56
		_nw56.Ctor__(_dafny.IntOfInt64(606), _280_v0)
		_330_v38 = _nw56
		var _331_v39 _dafny.Map
		_ = _331_v39
		_331_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v35, !(false))
		var _332_v40 _dafny.Map
		_ = _332_v40
		_332_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v36, (_330_v38).Fm4(_328_v36, (func() bool {
			if (_331_v39).Contains(_327_v35) {
				return (_331_v39).Get(_327_v35).(bool)
			}
			return (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
		})(), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState))
		_332_v40 = (_332_v40).Update(_328_v36, _283_v3)
		var _333_v41 _dafny.Sequence
		_ = _333_v41
		_333_v41 = _dafny.SeqOf(_281_v1)
		var _334_v42 _dafny.Map
		_ = _334_v42
		_334_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v3, _333_v41)
		var _rhs46 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_334_v42).Contains(_328_v36) {
				return (_334_v42).Get(_328_v36).(_dafny.Sequence)
			}
			return _333_v41
		})(), _333_v41), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-929), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_334_v42).Contains(_328_v36) {
				return (_334_v42).Get(_328_v36).(_dafny.Sequence)
			}
			return _333_v41
		})(), _333_v41)).Cardinality()))).Uint32(), Companion_Default___.Fm5(_288_globalState)), (Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_334_v42).Contains(_328_v36) {
				return (_334_v42).Get(_328_v36).(_dafny.Sequence)
			}
			return _333_v41
		})(), _333_v41), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-929), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_334_v42).Contains(_328_v36) {
				return (_334_v42).Get(_328_v36).(_dafny.Sequence)
			}
			return _333_v41
		})(), _333_v41)).Cardinality()))).Uint32(), Companion_Default___.Fm5(_288_globalState))).Cardinality()))).Uint32(), _dafny.SeqOf(_280_v0))
		_ = _rhs46
		var _rhs47 _dafny.Array = _286_v6
		_ = _rhs47
		var _rhs48 bool = (Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(512))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_335_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))).Cardinality()), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_327_v35, _327_v35)).Cardinality())) >= 0
		_ = _rhs48
		var _lhs41 *GlobalState = _288_globalState
		_ = _lhs41
		_333_v41 = _rhs46
		_286_v6 = _rhs47
		_lhs41.F4 = _rhs48
		var _336_v43 _dafny.Map
		_ = _336_v43
		_336_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_329_v37, (_dafny.Zero).Minus(_283_v3))
		var _337_v44 _dafny.Map
		_ = _337_v44
		_337_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_329_v37).Cardinality())), _328_v36), _336_v43)
		var _338_v45 _dafny.MultiSet
		_ = _338_v45
		_338_v45 = _dafny.MultiSetOf((_329_v37).Select((Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_329_v37).Cardinality()))).Uint32()).(_dafny.Int), _328_v36, (_329_v37).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.IntOfUint32((_329_v37).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus((_dafny.Zero).Minus(_328_v36)))
		_337_v44 = (_337_v44).Update(_328_v36, (_336_v43).Update(_329_v37, (func() _dafny.Int {
			if (_338_v45).Contains(_283_v3) {
				return (_338_v45).Multiplicity(_283_v3)
			}
			return (_290_v9).Cardinality()
		})()))
		var _source5 D0 = Companion_D0_.Create_DC1_(_280_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)), _330_v38, _328_v36)
		_ = _source5
		if _source5.Is_DC0() {
			var _339___mcc_h0 _dafny.Set = _source5.Get_().(D0_DC0).Cf0
			_ = _339___mcc_h0
			var _340___mcc_h1 _dafny.Sequence = _source5.Get_().(D0_DC0).Cf1
			_ = _340___mcc_h1
			var _341_cf1 _dafny.Sequence = _340___mcc_h1
			_ = _341_cf1
			var _342_cf0 _dafny.Set = _339___mcc_h0
			_ = _342_cf0
			_330_v38 = _330_v38
			(_288_globalState).F4 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)), _281_v1)
			var _343_v46 _dafny.Set
			_ = _343_v46
			_343_v46 = _dafny.SetOf(_286_v6, _286_v6, _286_v6, _286_v6)
			_343_v46 = _dafny.SetOf(_286_v6, _286_v6, _286_v6, _286_v6, _286_v6)
			var _344_v47 _dafny.Map
			_ = _344_v47
			_344_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _285_v5)
			var _345_v48 D1
			_ = _345_v48
			_345_v48 = Companion_D1_.Create_DC4_(_282_v2)
			_344_v47 = (_344_v47).Update((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _dafny.SeqOf((_345_v48).Dtor_cf10()))
		} else if _source5.Is_DC1() {
			var _346___mcc_h2 bool = _source5.Get_().(D0_DC1).Cf2
			_ = _346___mcc_h2
			var _347___mcc_h3 _dafny.Map = _source5.Get_().(D0_DC1).Cf3
			_ = _347___mcc_h3
			var _348___mcc_h4 *C0 = _source5.Get_().(D0_DC1).Cf4
			_ = _348___mcc_h4
			var _349___mcc_h5 _dafny.Int = _source5.Get_().(D0_DC1).Cf5
			_ = _349___mcc_h5
			var _350_cf5 _dafny.Int = _349___mcc_h5
			_ = _350_cf5
			var _351_cf4 *C0 = _348___mcc_h4
			_ = _351_cf4
			var _352_cf3 _dafny.Map = _347___mcc_h3
			_ = _352_cf3
			var _353_cf2 bool = _346___mcc_h2
			_ = _353_cf2
			_351_cf4 = _351_cf4
			var _354_v49 _dafny.Sequence
			_ = _354_v49
			_354_v49 = _dafny.SeqOf(_330_v38, _351_cf4)
			_354_v49 = _dafny.Companion_Sequence_.Concatenate(_354_v49, _dafny.SeqOf(_330_v38))
			_289_v8 = (_289_v8).Update(_353_cf2, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
			var _355_v50 _dafny.Map
			_ = _355_v50
			_355_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_329_v37).Cardinality()), _353_cf2)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
			_ = _index56
			var _rhs49 bool = ((_329_v37).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.IntOfUint32((_329_v37).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(Companion_Default___.SafeModuloInt(_328_v36, _350_cf5)) <= 0
			_ = _rhs49
			var _rhs50 bool = (func() bool {
				if (_355_v50).Contains(_283_v3) {
					return (_355_v50).Get(_283_v3).(bool)
				}
				return (_283_v3).Cmp((_329_v37).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.IntOfUint32((_329_v37).Cardinality()))).Uint32()).(_dafny.Int)) <= 0
			})()
			_ = _rhs50
			var _rhs51 bool = true
			_ = _rhs51
			var _lhs42 _dafny.Array = _286_v6
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
			_ = _lhs43
			var _lhs44 *GlobalState = _288_globalState
			_ = _lhs44
			(_lhs42).ArraySet1(_rhs49, (_lhs43).Int())
			_lhs44.F4 = _rhs50
			_353_cf2 = _rhs51
		} else if _source5.Is_DC2() {
			var _356___mcc_h6 _dafny.Int = _source5.Get_().(D0_DC2).Cf6
			_ = _356___mcc_h6
			var _357___mcc_h7 _dafny.MultiSet = _source5.Get_().(D0_DC2).Cf7
			_ = _357___mcc_h7
			var _358___mcc_h8 bool = _source5.Get_().(D0_DC2).Cf8
			_ = _358___mcc_h8
			var _359_cf8 bool = _358___mcc_h8
			_ = _359_cf8
			var _360_cf7 _dafny.MultiSet = _357___mcc_h7
			_ = _360_cf7
			var _361_cf6 _dafny.Int = _356___mcc_h6
			_ = _361_cf6
			var _362_v51 _dafny.Int
			_ = _362_v51
			var _out5 _dafny.Int
			_ = _out5
			_out5 = Companion_Default___.M0(_327_v35, _280_v0, _288_globalState)
			_362_v51 = _out5
			(_288_globalState).F4 = (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
			var _363_v52 D0
			_ = _363_v52
			_363_v52 = Companion_D0_.Create_DC1_(_280_v0, _289_v8, _330_v38, (_290_v9).Cardinality())
			(_288_globalState).F4 = (_363_v52).Dtor_cf2()
			(_288_globalState).F10 = _328_v36
		} else {
			var _364___mcc_h9 D0 = _source5.Get_().(D0_DC3).Cf9
			_ = _364___mcc_h9
			var _365_cf9 D0 = _364___mcc_h9
			_ = _365_cf9
			var _366_v53 _dafny.Int
			_ = _366_v53
			var _out6 _dafny.Int
			_ = _out6
			_out6 = Companion_Default___.M0(_dafny.UnicodeSeqOfUtf8Bytes("tiftisjra"), Companion_Default___.Fm3(true, _dafny.IntOfInt64(-562), _283_v3, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState), _288_globalState)
			_366_v53 = _out6
			var _367_v54 _dafny.Map
			_ = _367_v54
			_367_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v38, _328_v36)
			(_288_globalState).F10 = (_328_v36).Times(((_367_v54).Update(_330_v38, _366_v53)).Cardinality())
			_366_v53 = _dafny.IntOfInt64(293)
			(_288_globalState).F16 = _366_v53
		}
	} else {
		var _368_v56 _dafny.Map
		_ = _368_v56
		_368_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-75), _dafny.IntOfInt64(511))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _369_v55 _dafny.Int
				_369_v55 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(-75)).Cmp(_369_v55) <= 0) && ((_369_v55).Cmp(_dafny.IntOfInt64(511)) < 0) {
					_coll10.Add(Companion_Default___.SafeDivisionInt(_369_v55, _dafny.IntOfUint32((_329_v37).Cardinality())))
				}
			}
			return _coll10.ToSet()
		}()).Cardinality(), _328_v36)
		var _370_v57 D2
		_ = _370_v57
		_370_v57 = Companion_D2_.Create_DC9_(_368_v56)
		var _371_v58 T0
		_ = _371_v58
		var _nw57 *C0 = New_C0_()
		_ = _nw57
		_nw57.Ctor__(((_370_v57).Dtor_cf16()).Cardinality(), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
		_371_v58 = _nw57
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _index57
		var _rhs52 bool = (!((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))) == (!(_368_v56).Equals(_368_v56))
		_ = _rhs52
		var _rhs53 T0 = _371_v58
		_ = _rhs53
		var _rhs54 _dafny.Int = Companion_Default___.SafeModuloInt((_283_v3).Plus(_371_v58.F18()), _371_v58.F18())
		_ = _rhs54
		var _rhs55 bool = _280_v0
		_ = _rhs55
		var _rhs56 bool = _280_v0
		_ = _rhs56
		var _lhs45 *GlobalState = _288_globalState
		_ = _lhs45
		var _lhs46 *GlobalState = _288_globalState
		_ = _lhs46
		var _lhs47 _dafny.Array = _286_v6
		_ = _lhs47
		var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _lhs48
		_lhs45.F4 = _rhs52
		_371_v58 = _rhs53
		_283_v3 = _rhs54
		_lhs46.F4 = _rhs55
		(_lhs47).ArraySet1(_rhs56, (_lhs48).Int())
		var _372_v59 *C0
		_ = _372_v59
		var _nw58 *C0 = New_C0_()
		_ = _nw58
		_nw58.Ctor__(_dafny.IntOfUint32((_281_v1).Cardinality()), (_371_v58).F19())
		_372_v59 = _nw58
		var _373_v60 _dafny.Map
		_ = _373_v60
		_373_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _372_v59)
		_373_v60 = _373_v60
		var _374_v61 _dafny.Set
		_ = _374_v61
		_374_v61 = _dafny.SetOf(_328_v36)
		var _375_v62 D0
		_ = _375_v62
		_375_v62 = Companion_D0_.Create_DC2_((_374_v61).Cardinality(), _dafny.MultiSetOf((_dafny.Zero).Minus(_283_v3), _328_v36, _371_v58.F18(), _dafny.IntOfInt64(473), _371_v58.F18()), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
		var _rhs57 bool = (_375_v62).Dtor_cf8()
		_ = _rhs57
		var _rhs58 _dafny.Sequence = _329_v37
		_ = _rhs58
		var _lhs49 *GlobalState = _288_globalState
		_ = _lhs49
		_lhs49.F4 = _rhs57
		_329_v37 = _rhs58
		_328_v36 = Companion_Default___.Fm0(_283_v3, Companion_Default___.Fm3((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _dafny.IntOfUint32((_327_v35).Cardinality()), (_290_v9).Cardinality(), (_371_v58).F19(), _288_globalState), _288_globalState)
		var _376_v63 D1
		_ = _376_v63
		_376_v63 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(800))
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _index58
		(_286_v6).ArraySet1(Companion_Default___.Fm3(!(_280_v0), Companion_Default___.Fm0((_284_v4).Cardinality(), _280_v0, _288_globalState), (_376_v63).Dtor_cf11(), !(false), _288_globalState), (_index58).Int())
	}
	_281_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_281_v1, _281_v1), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)), (Companion_Default___.SafeIndex(_328_v36, _dafny.IntOfUint32((_dafny.SeqOf((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))).Cardinality()))).Uint32(), false))
	var _377_v64 _dafny.Array
	_ = _377_v64
	var _nw59 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(7))
	_ = _nw59
	_377_v64 = _nw59
	var _378_v65 _dafny.Map
	_ = _378_v65
	_378_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _377_v64)
	if !((_378_v65).Update((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _377_v64)).Equals((_378_v65).Merge(_378_v65)) {
		var _379_v66 _dafny.Set
		_ = _379_v66
		_379_v66 = _dafny.SetOf(Companion_Default___.Fm6(_dafny.IntOfInt64(-990), _327_v35, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState), _290_v9)
		_328_v36 = ((_379_v66).Union((_379_v66).Intersection(_379_v66))).Cardinality()
		_329_v37 = _dafny.Companion_Sequence_.Concatenate(_329_v37, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(250), _328_v36), _329_v37))
		_280_v0 = _280_v0
		if !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_329_v37, _dafny.SeqOf(_dafny.IntOfInt64(408), _328_v36)), _329_v37)) {
			(_288_globalState).F0 = (func() _dafny.Int {
				if (_284_v4).Contains(_280_v0) {
					return (_284_v4).Get(_280_v0).(_dafny.Int)
				}
				return (_283_v3).Plus(_283_v3)
			})()
			var _380_v67 D1
			_ = _380_v67
			_380_v67 = Companion_D1_.Create_DC4_(_282_v2)
			var _381_v68 _dafny.Map
			_ = _381_v68
			_381_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _380_v67)
			var _382_v69 *C0
			_ = _382_v69
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__(_328_v36, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
			_382_v69 = _nw60
			var _383_v70 _dafny.Sequence
			_ = _383_v70
			_383_v70 = _dafny.SeqOf(_382_v69)
			_381_v68 = (_381_v68).Update((_281_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_383_v70).Cardinality()), _dafny.IntOfUint32((_281_v1).Cardinality()))).Uint32()).(bool), Companion_D1_.Create_DC4_(_282_v2))
			var _384_v71 _dafny.CodePoint
			_ = _384_v71
			_384_v71 = _dafny.CodePoint('o')
			var _385_v72 _dafny.Map
			_ = _385_v72
			_385_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v3, _384_v71)
			var _386_v73 _dafny.MultiSet
			_ = _386_v73
			_386_v73 = _dafny.MultiSetOf(_283_v3, (_dafny.Zero).Minus((_dafny.Zero).Minus((_385_v72).Cardinality())), _328_v36, _328_v36, _328_v36)
			(_288_globalState).F0 = ((_386_v73).Update(_283_v3, Companion_Default___.Abs(_dafny.IntOfUint32((_327_v35).Cardinality())))).Cardinality()
			var _387_v74 D1
			_ = _387_v74
			_387_v74 = Companion_D1_.Create_DC5_(_328_v36)
			var _388_v75 *C0
			_ = _388_v75
			var _nw61 *C0 = New_C0_()
			_ = _nw61
			_nw61.Ctor__((_dafny.Zero).Minus(((_387_v74).Dtor_cf11()).Minus(_283_v3)), _280_v0)
			_388_v75 = _nw61
			var _389_v76 _dafny.Map
			_ = _389_v76
			_389_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_388_v75, _280_v0)
			_389_v76 = (_389_v76).Update((func() *C0 {
				if _280_v0 {
					return _388_v75
				}
				return _388_v75
			})(), !((func() bool {
				if (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool) {
					return (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
				}
				return (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
			})()))
		} else {
			var _390_v77 _dafny.CodePoint
			_ = _390_v77
			_390_v77 = _dafny.CodePoint('o')
			(_288_globalState).F4 = (_328_v36).Cmp(Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_327_v35, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-268), _dafny.IntOfUint32((_327_v35).Cardinality()))).Uint32(), _390_v77)).Cardinality()), _280_v0, _288_globalState)) <= 0
			var _391_v78 _dafny.MultiSet
			_ = _391_v78
			_391_v78 = _dafny.MultiSetOf((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
			var _392_v79 _dafny.Set
			_ = _392_v79
			_392_v79 = _dafny.SetOf(_283_v3, (_391_v78).Cardinality())
			var _393_v80 *C0
			_ = _393_v80
			var _nw62 *C0 = New_C0_()
			_ = _nw62
			_nw62.Ctor__((_283_v3).Minus((Companion_D1_.Create_DC7_((_392_v79).Cardinality(), _328_v36)).Dtor_cf13()), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
			_393_v80 = _nw62
			var _394_v81 _dafny.Array
			_ = _394_v81
			var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
			_ = _nw63
			_394_v81 = _nw63
			var _395_v82 _dafny.MultiSet
			_ = _395_v82
			_395_v82 = _dafny.MultiSetOf(_283_v3, _328_v36)
			var _396_v83 _dafny.Map
			_ = _396_v83
			_396_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_395_v82).Contains(_283_v3) {
					return (_395_v82).Multiplicity(_283_v3)
				}
				return _dafny.IntOfInt64(125)
			})(), _395_v82)
			var _397_v84 _dafny.Map
			_ = _397_v84
			_397_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_396_v83).Merge(_396_v83), _394_v81)
			_394_v81 = (func() _dafny.Array {
				if (_397_v84).Contains(_396_v83) {
					return (_397_v84).Get(_396_v83).(_dafny.Array)
				}
				return _394_v81
			})()
			var _398_v85 _dafny.Array
			_ = _398_v85
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
			_ = _nw64
			_398_v85 = _nw64
			_398_v85 = _398_v85
			(_288_globalState).F10 = _283_v3
		}
		var _399_v86 _dafny.Array
		_ = _399_v86
		var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
		_ = _nw65
		_399_v86 = _nw65
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_399_v86), 0))
		_ = _index59
		(_399_v86).ArraySet1(_327_v35, (_index59).Int())
		var _400_v87 *C0
		_ = _400_v87
		var _nw66 *C0 = New_C0_()
		_ = _nw66
		_nw66.Ctor__(_283_v3, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
		_400_v87 = _nw66
		var _401_v88 _dafny.Sequence
		_ = _401_v88
		_401_v88 = _dafny.SeqOf(_327_v35, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hl"), _327_v35))
		var _402_v89 _dafny.Sequence
		_ = _402_v89
		_402_v89 = _dafny.SeqOf(_400_v87, _400_v87, _400_v87, _400_v87)
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_399_v86), 0))
		_ = _index60
		var _rhs59 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-560))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_403_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))
		_ = _rhs59
		var _rhs60 _dafny.Sequence = (_401_v88).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_328_v36), _dafny.IntOfUint32((_401_v88).Cardinality()))).Uint32()).(_dafny.Sequence)
		_ = _rhs60
		var _rhs61 *C0 = (_402_v89).Select((Companion_Default___.SafeIndex((_328_v36).Plus(_dafny.IntOfUint32((_329_v37).Cardinality())), _dafny.IntOfUint32((_402_v89).Cardinality()))).Uint32()).(*C0)
		_ = _rhs61
		var _rhs62 _dafny.Int = (_dafny.Zero).Minus(_328_v36)
		_ = _rhs62
		var _lhs50 _dafny.Array = _399_v86
		_ = _lhs50
		var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_399_v86), 0))
		_ = _lhs51
		_327_v35 = _rhs59
		(_lhs50).ArraySet1(_rhs60, (_lhs51).Int())
		_400_v87 = _rhs61
		_283_v3 = _rhs62
	} else {
		var _404_v90 _dafny.CodePoint
		_ = _404_v90
		_404_v90 = _dafny.CodePoint('n')
		(_288_globalState).F4 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_327_v35, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.IntOfUint32((_327_v35).Cardinality()))).Uint32(), _404_v90), _404_v90)
		if !((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)) {
			(_288_globalState).F16 = (_328_v36).Plus(_283_v3)
			var _405_v91 _dafny.Array
			_ = _405_v91
			var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
			_ = _nw67
			_405_v91 = _nw67
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_405_v91), 0))
			_ = _index61
			(_405_v91).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_288_globalState), _281_v1), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_405_v91), 0))
			_ = _index62
			(_405_v91).ArraySet1(_dafny.SeqOf(false, _dafny.Companion_Sequence_.Equal(_327_v35, _327_v35), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)), (_index62).Int())
			var _406_v92 *C0
			_ = _406_v92
			var _nw68 *C0 = New_C0_()
			_ = _nw68
			_nw68.Ctor__(_dafny.IntOfInt64(-238), ((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)) == (Companion_Default___.Fm3(false, _328_v36, _328_v36, _280_v0, _288_globalState)))
			_406_v92 = _nw68
			(_288_globalState).F10 = (_406_v92).Fm4(_328_v36, _280_v0, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState)
			var _407_v93 D1
			_ = _407_v93
			_407_v93 = Companion_D1_.Create_DC4_(_282_v2)
			var _408_v94 _dafny.Map
			_ = _408_v94
			_408_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_407_v93, _286_v6)
			_408_v94 = (_408_v94).Update(_407_v93, _286_v6)
		} else {
			var _409_v95 _dafny.MultiSet
			_ = _409_v95
			_409_v95 = _dafny.MultiSetOf(_283_v3)
			var _410_v96 _dafny.MultiSet
			_ = _410_v96
			_410_v96 = _dafny.MultiSetOf(!(!((_409_v95).IsDisjointFrom(_409_v95))))
			_410_v96 = (_410_v96).Intersection(_410_v96)
			_404_v90 = _dafny.CodePoint('k')
			var _411_v97 _dafny.Array
			_ = _411_v97
			var _nwElement0_17 _dafny.CodePoint = _404_v90
			_ = _nwElement0_17
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(22))
			_ = _nw69
			(_nw69).ArraySet1CodePoint(_nwElement0_17, 0)
			(_nw69).ArraySet1CodePoint(_404_v90, 1)
			(_nw69).ArraySet1CodePoint(_dafny.CodePoint('t'), 2)
			(_nw69).ArraySet1CodePoint(_dafny.CodePoint('k'), 3)
			(_nw69).ArraySet1CodePoint(_404_v90, 4)
			(_nw69).ArraySet1CodePoint(_404_v90, 5)
			(_nw69).ArraySet1CodePoint(Companion_Default___.Fm7(_288_globalState), 6)
			(_nw69).ArraySet1CodePoint(_404_v90, 7)
			(_nw69).ArraySet1CodePoint(_404_v90, 8)
			(_nw69).ArraySet1CodePoint(_404_v90, 9)
			(_nw69).ArraySet1CodePoint(_404_v90, 10)
			(_nw69).ArraySet1CodePoint(_404_v90, 11)
			(_nw69).ArraySet1CodePoint(_404_v90, 12)
			(_nw69).ArraySet1CodePoint(_404_v90, 13)
			(_nw69).ArraySet1CodePoint(_404_v90, 14)
			(_nw69).ArraySet1CodePoint(_404_v90, 15)
			(_nw69).ArraySet1CodePoint(_404_v90, 16)
			(_nw69).ArraySet1CodePoint(_404_v90, 17)
			(_nw69).ArraySet1CodePoint(_dafny.CodePoint('n'), 18)
			(_nw69).ArraySet1CodePoint(_404_v90, 19)
			(_nw69).ArraySet1CodePoint(_404_v90, 20)
			(_nw69).ArraySet1CodePoint(_404_v90, 21)
			_411_v97 = _nw69
			var _412_v98 _dafny.Map
			_ = _412_v98
			_412_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _411_v97)
			_412_v98 = _412_v98
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
			_ = _index63
			(_286_v6).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-809))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_413_v90 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_414_i8 _dafny.Int) _dafny.CodePoint {
					return _413_v90
				}
			})(_404_v90))), _dafny.UnicodeSeqOfUtf8Bytes("qhsmdmwu")), _404_v90), (_index63).Int())
			var _rhs63 bool = true
			_ = _rhs63
			var _rhs64 bool = ((_409_v95).Intersection(_409_v95)).IsSubsetOf(_409_v95)
			_ = _rhs64
			var _rhs65 _dafny.Array = _286_v6
			_ = _rhs65
			var _lhs52 *GlobalState = _288_globalState
			_ = _lhs52
			_280_v0 = _rhs63
			_lhs52.F4 = _rhs64
			_286_v6 = _rhs65
		}
		var _415_v99 *C0
		_ = _415_v99
		var _nw70 *C0 = New_C0_()
		_ = _nw70
		_nw70.Ctor__(_dafny.IntOfInt64(651), _dafny.Companion_Sequence_.IsPrefixOf(_281_v1, _dafny.SeqOf((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))))
		_415_v99 = _nw70
		_415_v99 = _415_v99
		(_288_globalState).F10 = _328_v36
		var _416_v100 _dafny.Map
		_ = _416_v100
		_416_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v36, _327_v35)
		var _417_v101 _dafny.Sequence
		_ = _417_v101
		_417_v101 = _dafny.SeqOf(_327_v35, _327_v35, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(84))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_418_v90 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_419_i11 _dafny.Int) _dafny.CodePoint {
				return _418_v90
			}
		})(_404_v90))))
		var _420_v102 _dafny.Array
		_ = _420_v102
		var _nwElement0_18 _dafny.Sequence = _327_v35
		_ = _nwElement0_18
		var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(17))
		_ = _nw71
		(_nw71).ArraySet1(_nwElement0_18, 0)
		(_nw71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_327_v35, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(957))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_421_v90 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_422_i9 _dafny.Int) _dafny.CodePoint {
				return _421_v90
			}
		})(_404_v90)))), 1)
		(_nw71).ArraySet1(_327_v35, 2)
		(_nw71).ArraySet1(_327_v35, 3)
		(_nw71).ArraySet1((func() _dafny.Sequence {
			if (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool) {
				return (func() _dafny.Sequence {
					if (_416_v100).Contains(_328_v36) {
						return (_416_v100).Get(_328_v36).(_dafny.Sequence)
					}
					return _327_v35
				})()
			}
			return _dafny.Companion_Sequence_.Update(_327_v35, (Companion_Default___.SafeIndex(_328_v36, _dafny.IntOfUint32((_327_v35).Cardinality()))).Uint32(), _404_v90)
		})(), 4)
		(_nw71).ArraySet1(_327_v35, 5)
		(_nw71).ArraySet1(_327_v35, 6)
		(_nw71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pgyyvl"), _327_v35), 7)
		(_nw71).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("llc"), 8)
		(_nw71).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cqphv"), 9)
		(_nw71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_327_v35, _dafny.UnicodeSeqOfUtf8Bytes("npufrcjrg")), 10)
		(_nw71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_327_v35, _327_v35), 11)
		(_nw71).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("sa"), 12)
		(_nw71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_327_v35, _dafny.Companion_Sequence_.Update(_327_v35, (Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_327_v35).Cardinality()))).Uint32(), _404_v90)), 13)
		(_nw71).ArraySet1((func() _dafny.Sequence {
			if (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool) {
				return _327_v35
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(923))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_423_v90 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_424_i10 _dafny.Int) _dafny.CodePoint {
					return _423_v90
				}
			})(_404_v90)))
		})(), 14)
		(_nw71).ArraySet1(Companion_Default___.Fm2(!(!((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState), 15)
		(_nw71).ArraySet1((_417_v101).Select((Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_417_v101).Cardinality()))).Uint32()).(_dafny.Sequence), 16)
		_420_v102 = _nw71
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_420_v102), 0))
		_ = _index64
		(_420_v102).ArraySet1(_327_v35, (_index64).Int())
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_420_v102), 0))
		_ = _index65
		(_420_v102).ArraySet1((func() _dafny.Sequence {
			if (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool) {
				return _327_v35
			}
			return _dafny.Companion_Sequence_.Update(_327_v35, (Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_327_v35).Cardinality()))).Uint32(), _404_v90)
		})(), (_index65).Int())
	}
	var _425_v103 _dafny.MultiSet
	_ = _425_v103
	_425_v103 = _dafny.MultiSetOf((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
	var _426_v104 _dafny.Map
	_ = _426_v104
	_426_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_v103, _327_v35)
	var _427_v105 _dafny.Map
	_ = _427_v105
	_427_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v0, _327_v35)
	_426_v104 = ((_426_v104).Update(_425_v103, _327_v35)).Update(Companion_Default___.Fm8((_289_v8).Cardinality(), Companion_D2_.Create_DC9_(Companion_Default___.Fm9(_280_v0, _288_globalState)), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _327_v35, _288_globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("xifslf"), (Companion_Default___.SafeIndex(_283_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xifslf")).Cardinality()))).Uint32(), (_327_v35).Select((Companion_Default___.SafeIndex(_328_v36, _dafny.IntOfUint32((_327_v35).Cardinality()))).Uint32()).(_dafny.CodePoint)), (func() _dafny.Sequence {
		if (_427_v105).Contains(_280_v0) {
			return (_427_v105).Get(_280_v0).(_dafny.Sequence)
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("m")
	})()))
	var _428_i12 _dafny.Int
	_ = _428_i12
	_428_i12 = _dafny.Zero
	{
		for _280_v0 {
			{
				if (_428_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_428_i12 = (_428_i12).Plus(_dafny.One)
				var _429_v106 _dafny.Array
				_ = _429_v106
				var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
				_ = _nw72
				_429_v106 = _nw72
				var _430_v107 _dafny.Array
				_ = _430_v107
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_5
				var _nw73 _dafny.Array
				_ = _nw73
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw73 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Sequence = func(_431_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("tof")
					}
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw73 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw73).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw73).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_430_v107 = _nw73
				var _432_v108 _dafny.Map
				_ = _432_v108
				_432_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_429_v106, _430_v107)
				_432_v108 = (_432_v108).Update(_429_v106, _430_v107)
				var _433_v109 _dafny.CodePoint
				_ = _433_v109
				_433_v109 = _dafny.CodePoint('r')
				var _434_v110 D2
				_ = _434_v110
				_434_v110 = Companion_D2_.Create_DC11_(_433_v109, _328_v36)
				var _435_v111 _dafny.Map
				_ = _435_v111
				_435_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v110, _327_v35)
				_435_v111 = (_435_v111).Update(_434_v110, _327_v35)
				_328_v36 = _283_v3
				var _436_v112 _dafny.Set
				_ = _436_v112
				_436_v112 = _dafny.SetOf(_dafny.MultiSetFromSeq(_329_v37))
				var _437_v113 _dafny.MultiSet
				_ = _437_v113
				_437_v113 = _dafny.MultiSetOf(_328_v36)
				if (_dafny.SetOf(_437_v113, _437_v113)).IsProperSubsetOf(_436_v112) {
					_433_v109 = _433_v109
					_280_v0 = (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
					var _438_v114 _dafny.Map
					_ = _438_v114
					_438_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v36, _dafny.IntOfInt64(802))
					var _439_v115 D0
					_ = _439_v115
					_439_v115 = Companion_D0_.Create_DC2_((_dafny.Zero).Minus((func() _dafny.Int {
						if (_438_v114).Contains(_283_v3) {
							return (_438_v114).Get(_283_v3).(_dafny.Int)
						}
						return _328_v36
					})()), _437_v113, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
					var _440_v116 _dafny.Set
					_ = _440_v116
					_440_v116 = _dafny.SetOf(Companion_D0_.Create_DC2_(_328_v36, _437_v113, _280_v0))
					var _441_v117 _dafny.Int
					_ = _441_v117
					var _out7 _dafny.Int
					_ = _out7
					_out7 = Companion_Default___.M0(_327_v35, !(!((_dafny.SetOf(_439_v115)).IsDisjointFrom(_440_v116))), _288_globalState)
					_441_v117 = _out7
					var _442_v118 _dafny.Map
					_ = _442_v118
					_442_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v117, _280_v0)
					var _443_v119 *C0
					_ = _443_v119
					var _nw74 *C0 = New_C0_()
					_ = _nw74
					_nw74.Ctor__((_dafny.Zero).Minus(_328_v36), _280_v0)
					_443_v119 = _nw74
					var _444_v120 _dafny.MultiSet
					_ = _444_v120
					_444_v120 = _dafny.MultiSetOf((func() *C0 {
						if (func() bool {
							if (_442_v118).Contains(_441_v117) {
								return (_442_v118).Get(_441_v117).(bool)
							}
							return _280_v0
						})() {
							return _443_v119
						}
						return _443_v119
					})(), _443_v119)
					var _445_v121 _dafny.Sequence
					_ = _445_v121
					_445_v121 = _dafny.SeqOf(_443_v119)
					_444_v120 = (_444_v120).Intersection(_dafny.MultiSetFromSeq(_445_v121))
					(_288_globalState).F16 = _283_v3
				} else {
					var _446_v122 *C0
					_ = _446_v122
					var _nw75 *C0 = New_C0_()
					_ = _nw75
					_nw75.Ctor__(_328_v36, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
					_446_v122 = _nw75
					_446_v122 = _446_v122
					(_288_globalState).F4 = _280_v0
					(_288_globalState).F16 = _283_v3
					var _447_v123 _dafny.Array
					_ = _447_v123
					var _nw76 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
					_ = _nw76
					_447_v123 = _nw76
					var _448_v124 D2
					_ = _448_v124
					_448_v124 = Companion_D2_.Create_DC10_()
					var _449_v125 _dafny.Map
					_ = _449_v125
					_449_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_447_v123, _448_v124)
					(_288_globalState).F16 = ((_449_v125).Update(_447_v123, _448_v124)).Cardinality()
					_284_v4 = (_284_v4).Update((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _328_v36)
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	(_288_globalState).F4 = _280_v0
	for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_286_v6), 0))); ; {
		_guard_loop_0, _ok11 := _iter11()
		if !_ok11 {
			break
		}
		var _450_i14 _dafny.Int
		_450_i14 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_450_i14).Sign() != -1) && ((_450_i14).Cmp(_dafny.ArrayLen((_286_v6), 0)) < 0)) {
			(_286_v6).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v36, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v3, _280_v0))).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_327_v35, _327_v35)).Cardinality())), (_450_i14).Int())
		}
	}
	var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
	_ = _index66
	(_286_v6).ArraySet1(((_283_v3).Cmp(_dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_427_v105).Contains(false) {
			return (_427_v105).Get(false).(_dafny.Sequence)
		}
		return _327_v35
	})()).Cardinality())) == 0) && (_280_v0), (_index66).Int())
	if (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool) {
		_287_v7 = ((func() _dafny.MultiSet {
			if _280_v0 {
				return _287_v7
			}
			return _dafny.MultiSetOf(_286_v6, _286_v6, _286_v6, _286_v6, _286_v6)
		})()).Difference(_287_v7)
		if (_283_v3).Cmp(_dafny.IntOfInt64(-187)) != 0 {
			var _451_v126 _dafny.Int
			_ = _451_v126
			var _out8 _dafny.Int
			_ = _out8
			_out8 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_327_v35, _327_v35), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), _288_globalState)
			_451_v126 = _out8
			(_288_globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_281_v1, _281_v1), _dafny.SeqOf(_280_v0))
			(_288_globalState).F4 = _280_v0
			(_288_globalState).F4 = ((func() bool {
				if _280_v0 {
					return false
				}
				return (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)
			})()) && (!((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)) || (_280_v0))
			var _452_v127 *C0
			_ = _452_v127
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(_328_v36, _280_v0)
			_452_v127 = _nw77
			var _453_v128 _dafny.MultiSet
			_ = _453_v128
			_453_v128 = _dafny.MultiSetOf(_dafny.IntOfInt64(563), (Companion_D0_.Create_DC1_(_280_v0, _289_v8, _452_v127, _283_v3)).Dtor_cf5())
			(_288_globalState).F4 = !(_453_v128).Equals((_453_v128).Union(_dafny.MultiSetOf(_328_v36)))
		} else {
			var _454_v129 _dafny.Map
			_ = _454_v129
			_454_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(394), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
			_454_v129 = _454_v129
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
			_ = _index67
			(_286_v6).ArraySet1(_280_v0, (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
			_ = _index68
			(_286_v6).ArraySet1((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool), (_index68).Int())
			var _455_v130 _dafny.Array
			_ = _455_v130
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
			_ = _nw78
			_455_v130 = _nw78
			var _456_v131 _dafny.CodePoint
			_ = _456_v131
			_456_v131 = _dafny.CodePoint('b')
			var _457_v132 _dafny.Sequence
			_ = _457_v132
			_457_v132 = _dafny.SeqOf(_455_v130, _455_v130, _455_v130)
			var _458_v133 *C0
			_ = _458_v133
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__(_283_v3, _280_v0)
			_458_v133 = _nw79
			var _459_v134 _dafny.MultiSet
			_ = _459_v134
			_459_v134 = _dafny.MultiSetOf(_458_v133)
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
			_ = _index69
			var _rhs66 _dafny.Array = (_457_v132).Select((Companion_Default___.SafeIndex((_459_v134).Cardinality(), _dafny.IntOfUint32((_457_v132).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs66
			var _rhs67 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_327_v35, _327_v35)
			_ = _rhs67
			var _rhs68 bool = true
			_ = _rhs68
			var _rhs69 _dafny.Int = ((func() _dafny.Set {
				if _280_v0 {
					return _dafny.SetOf(!((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool)), _280_v0)
				}
				return _290_v9
			})()).Cardinality()
			_ = _rhs69
			var _rhs70 _dafny.CodePoint = _456_v131
			_ = _rhs70
			var _lhs53 _dafny.Array = _286_v6
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
			_ = _lhs54
			var _lhs55 *GlobalState = _288_globalState
			_ = _lhs55
			_455_v130 = _rhs66
			_327_v35 = _rhs67
			(_lhs53).ArraySet1(_rhs68, (_lhs54).Int())
			_lhs55.F0 = _rhs69
			_456_v131 = _rhs70
			(_288_globalState).F0 = (_dafny.IntOfUint32((_281_v1).Cardinality())).Times(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_327_v35).Cardinality()), _283_v3))
		}
		var _460_v135 *C0
		_ = _460_v135
		var _nw80 *C0 = New_C0_()
		_ = _nw80
		_nw80.Ctor__(Companion_Default___.SafeDivisionInt(_328_v36, _328_v36), (_283_v3).Cmp(_283_v3) > 0)
		_460_v135 = _nw80
		var _461_v136 _dafny.Array
		_ = _461_v136
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
		_ = _nw81
		_461_v136 = _nw81
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_461_v136), 0))
		_ = _index70
		(_461_v136).ArraySet1(_425_v103, (_index70).Int())
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_461_v136), 0))
		_ = _index71
		(_461_v136).ArraySet1(((_425_v103).Intersection(_425_v103)).Difference(_425_v103), (_index71).Int())
		var _462_v137 _dafny.Array
		_ = _462_v137
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
		_ = _nw82
		_462_v137 = _nw82
		_462_v137 = _462_v137
	} else {
		var _463_v138 *C0
		_ = _463_v138
		var _nw83 *C0 = New_C0_()
		_ = _nw83
		_nw83.Ctor__(_dafny.IntOfInt64(268), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))).Int()).(bool))
		_463_v138 = _nw83
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _index72
		(_286_v6).ArraySet1(_280_v0, (_index72).Int())
		var _464_v139 _dafny.MultiSet
		_ = _464_v139
		_464_v139 = _dafny.MultiSetOf(_dafny.IntOfInt64(826))
		var _465_v140 T0
		_ = _465_v140
		var _nw84 *C0 = New_C0_()
		_ = _nw84
		_nw84.Ctor__((_464_v139).Cardinality(), true)
		_465_v140 = _nw84
		var _466_v141 _dafny.CodePoint
		_ = _466_v141
		_466_v141 = _dafny.CodePoint('v')
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _index73
		var _rhs71 bool = !(!(_280_v0))
		_ = _rhs71
		var _rhs72 T0 = _465_v140
		_ = _rhs72
		var _rhs73 _dafny.CodePoint = _466_v141
		_ = _rhs73
		var _rhs74 bool = !(!(_280_v0))
		_ = _rhs74
		var _rhs75 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(Companion_Default___.Fm0(_465_v140.F18(), (_465_v140).F19(), _288_globalState), (_465_v140).F19(), _288_globalState), _dafny.IntOfUint32((_dafny.SeqOf(_466_v141, _466_v141, _466_v141, _466_v141)).Cardinality()))
		_ = _rhs75
		var _lhs56 _dafny.Array = _286_v6
		_ = _lhs56
		var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _lhs57
		var _lhs58 *GlobalState = _288_globalState
		_ = _lhs58
		var _lhs59 *GlobalState = _288_globalState
		_ = _lhs59
		(_lhs56).ArraySet1(_rhs71, (_lhs57).Int())
		_465_v140 = _rhs72
		_466_v141 = _rhs73
		_lhs58.F4 = _rhs74
		_lhs59.F16 = _rhs75
		var _467_v142 *C0
		_ = _467_v142
		var _nw85 *C0 = New_C0_()
		_ = _nw85
		_nw85.Ctor__(_328_v36, _280_v0)
		_467_v142 = _nw85
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_286_v6), 0))
		_ = _index74
		(_286_v6).ArraySet1(_dafny.Companion_Sequence_.Equal(_327_v35, _327_v35), (_index74).Int())
	}
	_327_v35 = _327_v35
	var _468_v143 _dafny.Sequence
	_ = _468_v143
	_468_v143 = _dafny.SeqOf(_327_v35)
	var _469_v144 _dafny.Int
	_ = _469_v144
	var _out9 _dafny.Int
	_ = _out9
	_out9 = Companion_Default___.M0((_468_v143).Select((Companion_Default___.SafeIndex(_328_v36, _dafny.IntOfUint32((_468_v143).Cardinality()))).Uint32()).(_dafny.Sequence), !_dafny.Companion_Sequence_.Equal(_281_v1, _281_v1), _288_globalState)
	_469_v144 = _out9
	_dafny.Print(_280_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_281_v1, _dafny.SeqOf(true, true, true, true, true, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_283_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-235)).UpdateUnsafe(false, _dafny.IntOfInt64(-2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_285_v5).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v7).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_288_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_288_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_288_globalState.F6, _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_288_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState.F11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-235))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_288_globalState).F15()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_288_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v9).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_291_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_327_v35, _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_328_v36)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_329_v37, _dafny.SeqOf(_dafny.One, _dafny.Zero, _dafny.IntOfInt64(7), _dafny.IntOfInt64(135), _dafny.IntOfInt64(250), _dafny.IntOfInt64(2), _dafny.One, _dafny.Zero, _dafny.IntOfInt64(7), _dafny.IntOfInt64(135))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_378_v65).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_425_v103).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_426_v104).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'))).UpdateUnsafe(_dafny.MultiSetOf(true, true, false), _dafny.UnicodeSeqOfUtf8Bytes("jifslfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_427_v105).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_428_i12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_468_v143, _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_469_v144)
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
	Cf0 _dafny.Set
	Cf1 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Set, Cf1 _dafny.Sequence) D0 {
	return D0{D0_DC0{Cf0, Cf1}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf2 bool
	Cf3 _dafny.Map
	Cf4 *C0
	Cf5 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf2 bool, Cf3 _dafny.Map, Cf4 *C0, Cf5 _dafny.Int) D0 {
	return D0{D0_DC1{Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf6 _dafny.Int
	Cf7 _dafny.MultiSet
	Cf8 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf6 _dafny.Int, Cf7 _dafny.MultiSet, Cf8 bool) D0 {
	return D0{D0_DC2{Cf6, Cf7, Cf8}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf9 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf9 D0) D0 {
	return D0{D0_DC3{Cf9}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.EmptySet, _dafny.EmptySeq)
}

func (_this D0) Dtor_cf0() _dafny.Set {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() *C0 {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.MultiSet {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf8() bool {
	return _this.Get_().(D0_DC2).Cf8
}

func (_this D0) Dtor_cf9() D0 {
	return _this.Get_().(D0_DC3).Cf9
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + data.Cf1.VerbatimString(true) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf0.Equals(data2.Cf0) && data1.Cf1.Equals(data2.Cf1)
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Equals(data2.Cf3) && data1.Cf4 == data2.Cf4 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7.Equals(data2.Cf7) && data1.Cf8 == data2.Cf8
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf9.Equals(data2.Cf9)
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

type D1_DC5 struct {
	Cf11 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf11 _dafny.Int) D1 {
	return D1{D1_DC5{Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf12 _dafny.Set
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf12 _dafny.Set) D1 {
	return D1{D1_DC6{Cf12}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC7 struct {
	Cf13 _dafny.Int
	Cf14 _dafny.Int
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_(Cf13 _dafny.Int, Cf14 _dafny.Int) D1 {
	return D1{D1_DC7{Cf13, Cf14}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

type D1_DC4 struct {
	Cf10 _dafny.Array
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf10 _dafny.Array) D1 {
	return D1{D1_DC4{Cf10}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC8 struct {
	Cf15 D1
}

func (D1_DC8) isD1() {}

func (CompanionStruct_D1_) Create_DC8_(Cf15 D1) D1 {
	return D1{D1_DC8{Cf15}}
}

func (_this D1) Is_DC8() bool {
	_, ok := _this.Get_().(D1_DC8)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(_dafny.Zero)
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Set {
	return _this.Get_().(D1_DC6).Cf12
}

func (_this D1) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D1_DC7).Cf13
}

func (_this D1) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D1_DC7).Cf14
}

func (_this D1) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf15() D1 {
	return _this.Get_().(D1_DC8).Cf15
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC8:
		{
			return "D1.DC8" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf12.Equals(data2.Cf12)
		}
	case D1_DC7:
		{
			data2, ok := other.Get_().(D1_DC7)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf10 == data2.Cf10
		}
	case D1_DC8:
		{
			data2, ok := other.Get_().(D1_DC8)
			return ok && data1.Cf15.Equals(data2.Cf15)
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

type D2_DC10 struct {
}

func (D2_DC10) isD2() {}

func (CompanionStruct_D2_) Create_DC10_() D2 {
	return D2{D2_DC10{}}
}

func (_this D2) Is_DC10() bool {
	_, ok := _this.Get_().(D2_DC10)
	return ok
}

type D2_DC11 struct {
	Cf17 _dafny.CodePoint
	Cf18 _dafny.Int
}

func (D2_DC11) isD2() {}

func (CompanionStruct_D2_) Create_DC11_(Cf17 _dafny.CodePoint, Cf18 _dafny.Int) D2 {
	return D2{D2_DC11{Cf17, Cf18}}
}

func (_this D2) Is_DC11() bool {
	_, ok := _this.Get_().(D2_DC11)
	return ok
}

type D2_DC9 struct {
	Cf16 _dafny.Map
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf16 _dafny.Map) D2 {
	return D2{D2_DC9{Cf16}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC10_()
}

func (_this D2) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D2_DC11).Cf17
}

func (_this D2) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D2_DC11).Cf18
}

func (_this D2) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D2_DC9).Cf16
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC10:
		{
			return "D2.DC10"
		}
	case D2_DC11:
		{
			return "D2.DC11" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC10:
		{
			_, ok := other.Get_().(D2_DC10)
			return ok
		}
	case D2_DC11:
		{
			data2, ok := other.Get_().(D2_DC11)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D3_DC13 struct {
	Cf20 bool
	Cf21 bool
	Cf22 D1
}

func (D3_DC13) isD3() {}

func (CompanionStruct_D3_) Create_DC13_(Cf20 bool, Cf21 bool, Cf22 D1) D3 {
	return D3{D3_DC13{Cf20, Cf21, Cf22}}
}

func (_this D3) Is_DC13() bool {
	_, ok := _this.Get_().(D3_DC13)
	return ok
}

type D3_DC12 struct {
	Cf19 _dafny.Array
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf19 _dafny.Array) D3 {
	return D3{D3_DC12{Cf19}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC14 struct {
	Cf23 D3
}

func (D3_DC14) isD3() {}

func (CompanionStruct_D3_) Create_DC14_(Cf23 D3) D3 {
	return D3{D3_DC14{Cf23}}
}

func (_this D3) Is_DC14() bool {
	_, ok := _this.Get_().(D3_DC14)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC13_(false, false, Companion_D1_.Default())
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC13).Cf20
}

func (_this D3) Dtor_cf21() bool {
	return _this.Get_().(D3_DC13).Cf21
}

func (_this D3) Dtor_cf22() D1 {
	return _this.Get_().(D3_DC13).Cf22
}

func (_this D3) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D3_DC12).Cf19
}

func (_this D3) Dtor_cf23() D3 {
	return _this.Get_().(D3_DC14).Cf23
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC13:
		{
			return "D3.DC13" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC14:
		{
			return "D3.DC14" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC13:
		{
			data2, ok := other.Get_().(D3_DC13)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21 && data1.Cf22.Equals(data2.Cf22)
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf19 == data2.Cf19
		}
	case D3_DC14:
		{
			data2, ok := other.Get_().(D3_DC14)
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D4_DC15 struct {
	Cf24 _dafny.MultiSet
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf24 _dafny.MultiSet) D4 {
	return D4{D4_DC15{Cf24}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D4) Dtor_cf24() _dafny.MultiSet {
	return _this.Get_().(D4_DC15).Cf24
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D5_DC17 struct {
	Cf26 D2
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf26 D2) D5 {
	return D5{D5_DC17{Cf26}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC18 struct {
	Cf27 D1
	Cf28 _dafny.Int
	Cf29 _dafny.CodePoint
	Cf30 _dafny.MultiSet
	Cf31 _dafny.Int
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf27 D1, Cf28 _dafny.Int, Cf29 _dafny.CodePoint, Cf30 _dafny.MultiSet, Cf31 _dafny.Int) D5 {
	return D5{D5_DC18{Cf27, Cf28, Cf29, Cf30, Cf31}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC16 struct {
	Cf25 _dafny.Map
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf25 _dafny.Map) D5 {
	return D5{D5_DC16{Cf25}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC19 struct {
	Cf32 D5
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf32 D5) D5 {
	return D5{D5_DC19{Cf32}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(Companion_D2_.Default())
}

func (_this D5) Dtor_cf26() D2 {
	return _this.Get_().(D5_DC17).Cf26
}

func (_this D5) Dtor_cf27() D1 {
	return _this.Get_().(D5_DC18).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf28
}

func (_this D5) Dtor_cf29() _dafny.CodePoint {
	return _this.Get_().(D5_DC18).Cf29
}

func (_this D5) Dtor_cf30() _dafny.MultiSet {
	return _this.Get_().(D5_DC18).Cf30
}

func (_this D5) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf31
}

func (_this D5) Dtor_cf25() _dafny.Map {
	return _this.Get_().(D5_DC16).Cf25
}

func (_this D5) Dtor_cf32() D5 {
	return _this.Get_().(D5_DC19).Cf32
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf32) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf26.Equals(data2.Cf26)
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf27.Equals(data2.Cf27) && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29 == data2.Cf29 && data1.Cf30.Equals(data2.Cf30) && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf32.Equals(data2.Cf32)
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
	F18() _dafny.Int
	F18_set_(value _dafny.Int)
	F19() bool
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
	F0   _dafny.Int
	F4   bool
	F6   _dafny.Sequence
	F10  _dafny.Int
	F11  _dafny.Map
	F13  _dafny.Array
	F16  _dafny.Int
	_f1  _dafny.Int
	_f2  _dafny.Int
	_f3  bool
	_f5  _dafny.Int
	_f7  _dafny.Int
	_f8  _dafny.Array
	_f9  _dafny.Int
	_f12 bool
	_f14 _dafny.Array
	_f15 _dafny.MultiSet
	_f17 _dafny.CodePoint
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F4 = false
	_this.F6 = _dafny.EmptySeq
	_this.F10 = _dafny.Zero
	_this.F11 = _dafny.EmptyMap
	_this.F13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F16 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.Zero
	_this._f3 = false
	_this._f5 = _dafny.Zero
	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f9 = _dafny.Zero
	_this._f12 = false
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = _dafny.EmptyMultiSet
	_this._f17 = _dafny.CodePoint('D')
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 _dafny.Int, f3 bool, f4 bool, f5 _dafny.Int, f6 _dafny.Sequence, f7 _dafny.Int, f8 _dafny.Array, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Map, f12 bool, f13 _dafny.Array, f14 _dafny.Array, f15 _dafny.MultiSet, f16 _dafny.Int, f17 _dafny.CodePoint) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
	}
}
func (_this *GlobalState) F1() _dafny.Int {
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
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Array {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Array {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.MultiSet {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F17() _dafny.CodePoint {
	{
		return _this._f17
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f18 _dafny.Int
	_f19 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f18 = _dafny.Zero
	_this._f19 = false
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

func (_this *C0) F18() _dafny.Int {
	return _this._f18
}
func (_this *C0) F18_set_(value _dafny.Int) {
	_this._f18 = value
}
func (_this *C0) F19() bool {
	return _this._f19
}
func (_this *C0) Ctor__(f18 _dafny.Int, f19 bool) {
	{
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *C0) Fm4(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(_this.F18())
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
