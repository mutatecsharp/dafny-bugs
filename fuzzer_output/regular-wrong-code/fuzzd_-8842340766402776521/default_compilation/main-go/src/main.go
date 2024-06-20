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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Int {
	return (Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(898))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mcftfnqg")).Cardinality()), _dafny.IntOfInt64(-391), (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(194), _dafny.IntOfInt64(317))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(194)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(317)) < 0) {
				_coll0.Add((_0_v0).Minus(_dafny.IntOfInt64(-949)), _dafny.SetOf(false))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())).Dtor_cf3()
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) bool {
	return (_dafny.SetOf(false)).IsDisjointFrom((_dafny.SetOf(true)).Difference(_dafny.SetOf(true, false, false)))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false, true, !(!(false)), !(false))).Cardinality(), false))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(807))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality())).Cmp((_dafny.SetOf(!(!(!(false))), !(!(!(!(true)))), !(true), false)).Cardinality()) == 0 {
		return _dafny.CodePoint('d')
	} else {
		return _dafny.CodePoint('g')
	}
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(759)), (_dafny.SetOf(true, false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-685))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_2_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-994)
		}))).Cardinality()), true)).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Int
			_3_v0 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-685))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_2_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-994)
			}))).Cardinality()), true)).Contains(_3_v0) {
				_coll1.Add((_3_v0).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg3 _dafny.Int) interface{} {
						return coer3(arg3)
					}
				}(func(_4_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('b'), _dafny.CodePoint('g'))
				}))).Cardinality())))
			}
		}
		return _coll1.ToSet()
	}(), _dafny.IntOfInt64(930)))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("j"), _dafny.UnicodeSeqOfUtf8Bytes("iptpuc")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(838))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 D1 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(243), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("adlori"))).Cardinality(), Companion_D3_.Create_DC9_(true))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("omdgnphda")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gtoykysed")).Cardinality()))
	_ = _source0
	if _source0.Is_DC2() {
		var _6___mcc_h0 _dafny.Int = _source0.Get_().(D1_DC2).Cf2
		_ = _6___mcc_h0
		var _7___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC2).Cf3
		_ = _7___mcc_h1
		var _8___mcc_h2 _dafny.Int = _source0.Get_().(D1_DC2).Cf4
		_ = _8___mcc_h2
		var _9___mcc_h3 _dafny.Int = _source0.Get_().(D1_DC2).Cf5
		_ = _9___mcc_h3
		var _10_cf5 _dafny.Int = _9___mcc_h3
		_ = _10_cf5
		var _11_cf4 _dafny.Int = _8___mcc_h2
		_ = _11_cf4
		var _12_cf3 _dafny.Int = _7___mcc_h1
		_ = _12_cf3
		var _13_cf2 _dafny.Int = _6___mcc_h0
		_ = _13_cf2
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false), true), _dafny.SeqOf(!(true)))
	} else if _source0.Is_DC1() {
		var _14___mcc_h4 _dafny.Map = _source0.Get_().(D1_DC1).Cf1
		_ = _14___mcc_h4
		var _15_cf1 _dafny.Map = _14___mcc_h4
		_ = _15_cf1
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.SeqOf(!(true), true))
	} else {
		var _16___mcc_h5 D1 = _source0.Get_().(D1_DC3).Cf6
		_ = _16___mcc_h5
		var _17_cf6 D1 = _16___mcc_h5
		_ = _17_cf6
		return _dafny.SeqOf(true, false, false)
	}
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, globalState *GlobalState) {
	(globalState).F10 = (p1).Times((p1).Times(p1))
	var _18_v0 *C0
	_ = _18_v0
	var _nw0 *C0 = New_C0_()
	_ = _nw0
	_nw0.Ctor__(p1)
	_18_v0 = _nw0
	var _19_v1 _dafny.Map
	_ = _19_v1
	_19_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, (_18_v0).F14())
	var _20_v2 _dafny.Sequence
	_ = _20_v2
	_20_v2 = _dafny.SeqOf(p0)
	var _21_v3 _dafny.Sequence
	_ = _21_v3
	_21_v3 = _dafny.UnicodeSeqOfUtf8Bytes("iaqmfjp")
	var _22_v4 _dafny.MultiSet
	_ = _22_v4
	_22_v4 = _dafny.MultiSetOf((_20_v2).Select((Companion_Default___.SafeIndex((_18_v0).F14(), _dafny.IntOfUint32((_20_v2).Cardinality()))).Uint32()).(bool), p0, p0)
	var _23_v5 _dafny.Array
	_ = _23_v5
	var _nwElement0_0 bool = (_19_v1).Contains(_18_v0)
	_ = _nwElement0_0
	var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(12))
	_ = _nw1
	(_nw1).ArraySet1(_nwElement0_0, 0)
	(_nw1).ArraySet1((_20_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_21_v3).Cardinality()), _dafny.IntOfUint32((_20_v2).Cardinality()))).Uint32()).(bool), 1)
	(_nw1).ArraySet1(p0, 2)
	(_nw1).ArraySet1(!((p0) || (p0)), 3)
	(_nw1).ArraySet1(p0, 4)
	(_nw1).ArraySet1(p0, 5)
	(_nw1).ArraySet1(p0, 6)
	(_nw1).ArraySet1(p0, 7)
	(_nw1).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((_21_v3).Cardinality()))).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rpjgtvwk")).Cardinality())) != 0, 8)
	(_nw1).ArraySet1(p0, 9)
	(_nw1).ArraySet1((_22_v4).IsProperSubsetOf(_22_v4), 10)
	(_nw1).ArraySet1((_18_v0).Fm4(p0, globalState), 11)
	_23_v5 = _nw1
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
	_ = _index0
	(_23_v5).ArraySet1(true, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
	_ = _index1
	(_23_v5).ArraySet1(p0, (_index1).Int())
	var _24_v6 *C0
	_ = _24_v6
	var _nw2 *C0 = New_C0_()
	_ = _nw2
	_nw2.Ctor__(Companion_Default___.SafeDivisionInt(p1, p1))
	_24_v6 = _nw2
	var _25_v7 _dafny.Map
	_ = _25_v7
	_25_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_20_v2).Cardinality()))
	(globalState).F10 = ((func() _dafny.Int {
		if (_25_v7).Contains(p0) {
			return (_25_v7).Get(p0).(_dafny.Int)
		}
		return p1
	})()).Minus((_18_v0).F14())
	var _26_v8 _dafny.Sequence
	_ = _26_v8
	_26_v8 = _dafny.SeqOf(_24_v6, _18_v0)
	_26_v8 = _26_v8
	if !((_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool)) {
		(globalState).F6 = (_24_v6).F14()
		(globalState).F10 = p1
		_25_v7 = (_25_v7).Update(p0, _dafny.IntOfInt64(800))
		var _27_v9 _dafny.Array
		_ = _27_v9
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw3
		_27_v9 = _nw3
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_27_v9), 0))
		_ = _index2
		(_27_v9).ArraySet1(p1, (_index2).Int())
		var _28_v10 _dafny.Map
		_ = _28_v10
		_28_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _23_v5)
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_27_v9), 0))
		_ = _index3
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
		_ = _index4
		var _rhs0 _dafny.Sequence = _21_v3
		_ = _rhs0
		var _rhs1 _dafny.Int = (_24_v6).F14()
		_ = _rhs1
		var _rhs2 _dafny.Int = (_28_v10).Cardinality()
		_ = _rhs2
		var _rhs3 bool = !(p0)
		_ = _rhs3
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 _dafny.Array = _27_v9
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_27_v9), 0))
		_ = _lhs2
		var _lhs3 _dafny.Array = _23_v5
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
		_ = _lhs4
		_21_v3 = _rhs0
		_lhs0.F10 = _rhs1
		(_lhs1).ArraySet1(_rhs2, (_lhs2).Int())
		(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
		var _29_v11 _dafny.Map
		_ = _29_v11
		_29_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(823))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_30_v5 _dafny.Array) func(_dafny.Int) _dafny.Map {
			return func(_31_i0 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_30_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_30_v5), 0))).Int()).(bool), (_30_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_30_v5), 0))).Int()).(bool))
			}
		})(_23_v5))))
		var _32_v12 _dafny.Map
		_ = _32_v12
		_32_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool), p0)
		var _33_v13 _dafny.Sequence
		_ = _33_v13
		_33_v13 = _dafny.SeqOf(_32_v12, _32_v12)
		_29_v11 = (_29_v11).Update((_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool), _33_v13)
	} else {
		var _34_v14 _dafny.Map
		_ = _34_v14
		_34_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(713), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool), (_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool))).Cardinality())
		var _35_v15 D1
		_ = _35_v15
		_35_v15 = Companion_D1_.Create_DC1_(_34_v14)
		var _pat_let_tv0 = _34_v14
		_ = _pat_let_tv0
		var _source1 D1 = func(_pat_let0_0 D1) D1 {
			return func(_36_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let1_0 _dafny.Map) D1 {
					return func(_37_dt__update_hcf1_h0 _dafny.Map) D1 {
						return Companion_D1_.Create_DC1_(_37_dt__update_hcf1_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_35_v15)
		_ = _source1
		if _source1.Is_DC2() {
			var _38___mcc_h0 _dafny.Int = _source1.Get_().(D1_DC2).Cf2
			_ = _38___mcc_h0
			var _39___mcc_h1 _dafny.Int = _source1.Get_().(D1_DC2).Cf3
			_ = _39___mcc_h1
			var _40___mcc_h2 _dafny.Int = _source1.Get_().(D1_DC2).Cf4
			_ = _40___mcc_h2
			var _41___mcc_h3 _dafny.Int = _source1.Get_().(D1_DC2).Cf5
			_ = _41___mcc_h3
			var _42_cf5 _dafny.Int = _41___mcc_h3
			_ = _42_cf5
			var _43_cf4 _dafny.Int = _40___mcc_h2
			_ = _43_cf4
			var _44_cf3 _dafny.Int = _39___mcc_h1
			_ = _44_cf3
			var _45_cf2 _dafny.Int = _38___mcc_h0
			_ = _45_cf2
			var _46_v16 _dafny.MultiSet
			_ = _46_v16
			_46_v16 = _dafny.MultiSetOf(_21_v3)
			var _47_v17 _dafny.Map
			_ = _47_v17
			_47_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_46_v16).Cardinality(), _21_v3)
			_21_v3 = (func() _dafny.Sequence {
				if (_47_v17).Contains((_18_v0).F14()) {
					return (_47_v17).Get((_18_v0).F14()).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(166))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_48_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				}))
			})()
			var _49_v18 *C0
			_ = _49_v18
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__((_18_v0).F14())
			_49_v18 = _nw4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index5
			(_23_v5).ArraySet1((_43_cf4).Cmp(_43_cf4) > 0, (_index5).Int())
			_44_cf3 = (_24_v6).F14()
		} else if _source1.Is_DC1() {
			var _50___mcc_h4 _dafny.Map = _source1.Get_().(D1_DC1).Cf1
			_ = _50___mcc_h4
			var _51_cf1 _dafny.Map = _50___mcc_h4
			_ = _51_cf1
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index6
			(_23_v5).ArraySet1((((_dafny.Zero).Minus((_24_v6).F14())).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_21_v3).Cardinality())))).Cmp((func() _dafny.Int {
				if (_25_v7).Contains(p0) {
					return (_25_v7).Get(p0).(_dafny.Int)
				}
				return (_24_v6).F14()
			})()) <= 0, (_index6).Int())
			var _52_v19 _dafny.Array
			_ = _52_v19
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw5
			_52_v19 = _nw5
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_52_v19), 0))
			_ = _index7
			(_52_v19).ArraySet1((_18_v0).F14(), (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_52_v19), 0))
			_ = _index8
			var _rhs4 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_21_v3).Cardinality())), (_24_v6).F14())
			_ = _rhs4
			var _rhs5 _dafny.Int = (_dafny.IntOfInt64(74)).Minus(p1)
			_ = _rhs5
			var _rhs6 _dafny.Int = Companion_Default___.SafeModuloInt((_18_v0).F14(), p1)
			_ = _rhs6
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			var _lhs6 _dafny.Array = _52_v19
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_52_v19), 0))
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			_lhs5.F10 = _rhs4
			(_lhs6).ArraySet1(_rhs5, (_lhs7).Int())
			_lhs8.F10 = _rhs6
			var _53_v20 _dafny.CodePoint
			_ = _53_v20
			_53_v20 = _dafny.CodePoint('c')
			_53_v20 = _53_v20
			_52_v19 = _52_v19
		} else {
			var _54___mcc_h5 D1 = _source1.Get_().(D1_DC3).Cf6
			_ = _54___mcc_h5
			var _55_cf6 D1 = _54___mcc_h5
			_ = _55_cf6
			(globalState).F10 = _dafny.IntOfInt64(568)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index9
			(_23_v5).ArraySet1((_24_v6).Fm5(globalState), (_index9).Int())
			var _56_v21 *C0
			_ = _56_v21
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__((_18_v0).F14())
			_56_v21 = _nw6
			var _57_v22 _dafny.Array
			_ = _57_v22
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
			_ = _nw7
			_57_v22 = _nw7
			var _58_v23 _dafny.Map
			_ = _58_v23
			_58_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _dafny.Companion_Sequence_.Update(_20_v2, (Companion_Default___.SafeIndex((_56_v21).F14(), _dafny.IntOfUint32((_20_v2).Cardinality()))).Uint32(), (_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool)))
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_57_v22), 0))
			_ = _index10
			(_57_v22).ArraySet1(_58_v23, (_index10).Int())
			var _59_v24 _dafny.CodePoint
			_ = _59_v24
			_59_v24 = _dafny.CodePoint('d')
			var _60_v25 _dafny.Map
			_ = _60_v25
			_60_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality())), _56_v21)
			var _61_v26 _dafny.Map
			_ = _61_v26
			_61_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v24, (func() *C0 {
				if (_60_v25).Contains(_25_v7) {
					return (_60_v25).Get(_25_v7).(*C0)
				}
				return _56_v21
			})())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_57_v22), 0))
			_ = _index11
			var _rhs7 _dafny.Int = (_24_v6).F14()
			_ = _rhs7
			var _rhs8 _dafny.Map = _58_v23
			_ = _rhs8
			var _rhs9 _dafny.Map = (_61_v26).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v24, _18_v0))
			_ = _rhs9
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			var _lhs10 _dafny.Array = _57_v22
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_57_v22), 0))
			_ = _lhs11
			_lhs9.F10 = _rhs7
			(_lhs10).ArraySet1(_rhs8, (_lhs11).Int())
			_61_v26 = _rhs9
		}
		var _62_v27 _dafny.Map
		_ = _62_v27
		_62_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool))
		(globalState).F10 = ((_62_v27).Cardinality()).Minus((_dafny.Zero).Minus(p1))
		if !(_dafny.Companion_Sequence_.Equal(_21_v3, _21_v3)) {
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index12
			(_23_v5).ArraySet1(false, (_index12).Int())
			var _63_v28 _dafny.Sequence
			_ = _63_v28
			_63_v28 = _dafny.SeqOf(p1, (_18_v0).F14())
			var _64_v29 D2
			_ = _64_v29
			_64_v29 = Companion_D2_.Create_DC5_(p0, p0)
			var _65_v30 _dafny.MultiSet
			_ = _65_v30
			_65_v30 = _dafny.MultiSetOf(_64_v29, _64_v29)
			var _66_v31 _dafny.MultiSet
			_ = _66_v31
			_66_v31 = _dafny.MultiSetOf(_dafny.IntOfInt64(568), (_65_v30).Cardinality(), (_24_v6).F14())
			var _67_v32 _dafny.Map
			_ = _67_v32
			_67_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.Companion_Sequence_.Update(_63_v28, (Companion_Default___.SafeIndex((_18_v0).F14(), _dafny.IntOfUint32((_63_v28).Cardinality()))).Uint32(), (_66_v31).Cardinality()))
			var _68_v33 _dafny.Set
			_ = _68_v33
			_68_v33 = _dafny.SetOf((_18_v0).F14(), _dafny.IntOfInt64(831))
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index13
			(_23_v5).ArraySet1(((func() _dafny.Set {
				if p0 {
					return _68_v33
				}
				return _68_v33
			})()).Contains((_67_v32).Cardinality()), (_index13).Int())
			var _69_v34 _dafny.CodePoint
			_ = _69_v34
			_69_v34 = _dafny.CodePoint('l')
			var _70_v35 _dafny.Map
			_ = _70_v35
			_70_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool)), _69_v34)
			(globalState).F10 = (_63_v28).Select((Companion_Default___.SafeIndex(((_70_v35).Cardinality()).Minus((_24_v6).F14()), _dafny.IntOfUint32((_63_v28).Cardinality()))).Uint32()).(_dafny.Int)
			var _71_v36 D1
			_ = _71_v36
			_71_v36 = Companion_D1_.Create_DC1_(_34_v14)
			var _72_v37 D1
			_ = _72_v37
			_72_v37 = Companion_D1_.Create_DC3_(_71_v36)
			_72_v37 = (func() D1 {
				if p0 {
					return Companion_D1_.Create_DC3_(_71_v36)
				}
				return _72_v37
			})()
			(globalState).F10 = p1
		} else {
			(globalState).F6 = (_18_v0).F14()
			var _73_v38 _dafny.CodePoint
			_ = _73_v38
			_73_v38 = _dafny.CodePoint('g')
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index14
			(_23_v5).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_21_v3, _21_v3), (Companion_Default___.SafeIndex((_18_v0).F14(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_21_v3, _21_v3)).Cardinality()))).Uint32(), _73_v38)).Cardinality()))).Cmp((_18_v0).F14()) != 0, (_index14).Int())
			var _74_v39 D2
			_ = _74_v39
			_74_v39 = Companion_D2_.Create_DC6_(!((_24_v6).Fm5(globalState)), ((_24_v6).F14()).Cmp(_dafny.IntOfUint32((_21_v3).Cardinality())) < 0, Companion_Default___.Fm3((_18_v0).F14(), (_24_v6).F14(), p0, globalState))
			var _pat_let_tv1 = _20_v2
			_ = _pat_let_tv1
			var _pat_let_tv2 = p0
			_ = _pat_let_tv2
			_74_v39 = func(_pat_let2_0 D2) D2 {
				return func(_75_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let3_0 bool) D2 {
						return func(_76_dt__update_hcf11_h0 bool) D2 {
							return Companion_D2_.Create_DC6_((_75_dt__update__tmp_h1).Dtor_cf10(), _76_dt__update_hcf11_h0, (_75_dt__update__tmp_h1).Dtor_cf12())
						}(_pat_let3_0)
					}(!_dafny.Companion_Sequence_.Contains(_pat_let_tv1, _pat_let_tv2))
				}(_pat_let2_0)
			}(_74_v39)
			var _77_v40 D1
			_ = _77_v40
			_77_v40 = Companion_D1_.Create_DC2_((_18_v0).F14(), (_18_v0).F14(), (_24_v6).F14(), Companion_Default___.Fm0(globalState))
			var _78_v41 D1
			_ = _78_v41
			_78_v41 = Companion_D1_.Create_DC3_(_77_v40)
			var _79_v42 _dafny.Array
			_ = _79_v42
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw8
			_79_v42 = _nw8
			var _80_v43 _dafny.Sequence
			_ = _80_v43
			_80_v43 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v7, (_24_v6).F14())).Cardinality(), (_24_v6).F14())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_79_v42), 0))
			_ = _index15
			(_79_v42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_80_v43, _80_v43), (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_79_v42), 0))
			_ = _index17
			var _rhs10 D1 = _78_v41
			_ = _rhs10
			var _rhs11 bool = (p0) == ((_23_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))).Int()).(bool))
			_ = _rhs11
			var _rhs12 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_80_v43, _80_v43)
			_ = _rhs12
			var _lhs12 _dafny.Array = _23_v5
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _79_v42
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_79_v42), 0))
			_ = _lhs15
			_78_v41 = _rhs10
			(_lhs12).ArraySet1(_rhs11, (_lhs13).Int())
			(_lhs14).ArraySet1(_rhs12, (_lhs15).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
			_ = _index18
			(_23_v5).ArraySet1(false, (_index18).Int())
		}
		var _81_v44 bool
		_ = _81_v44
		_81_v44 = false
		_81_v44 = _81_v44
		var _82_v45 D1
		_ = _82_v45
		_82_v45 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-506), (_34_v14).Cardinality(), (_24_v6).F14(), (_24_v6).F14())
		var _83_v46 _dafny.Sequence
		_ = _83_v46
		_83_v46 = _dafny.SeqOf((_dafny.Zero).Minus((_24_v6).F14()), (_82_v45).Dtor_cf4())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_23_v5), 0))
		_ = _index19
		(_23_v5).ArraySet1((((_24_v6).F14()).Minus((_dafny.Zero).Minus((_83_v46).Select((Companion_Default___.SafeIndex((_18_v0).F14(), _dafny.IntOfUint32((_83_v46).Cardinality()))).Uint32()).(_dafny.Int)))).Cmp(_dafny.IntOfUint32((_83_v46).Cardinality())) > 0, (_index19).Int())
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _84_v0 _dafny.Int
	_ = _84_v0
	_84_v0 = _dafny.IntOfInt64(457)
	var _85_v1 bool
	_ = _85_v1
	_85_v1 = false
	var _86_v2 _dafny.Map
	_ = _86_v2
	_86_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, _85_v1)
	var _87_v3 _dafny.Sequence
	_ = _87_v3
	_87_v3 = _dafny.SeqOf((func() bool {
		if (_86_v2).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("koq")).Cardinality())) {
			return (_86_v2).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("koq")).Cardinality())).(bool)
		}
		return _85_v1
	})())
	var _88_v4 _dafny.Map
	_ = _88_v4
	_88_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(924), _87_v3)
	var _89_v5 _dafny.Sequence
	_ = _89_v5
	_89_v5 = _dafny.UnicodeSeqOfUtf8Bytes("rlnyx")
	var _90_v6 _dafny.Array
	_ = _90_v6
	var _nw9 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
	_ = _nw9
	_90_v6 = _nw9
	var _91_globalState *GlobalState
	_ = _91_globalState
	var _nw10 *GlobalState = New_GlobalState_()
	_ = _nw10
	_nw10.Ctor__(true, true, true, true, (_88_v4).Merge((_88_v4).Update(_84_v0, _87_v3)), _dafny.IntOfInt64(-8), _dafny.IntOfInt64(256), false, _89_v5, _89_v5, _dafny.IntOfInt64(133), _90_v6, true, _dafny.IntOfInt64(191))
	_91_globalState = _nw10
	if !(((_dafny.Zero).Minus((_dafny.Zero).Minus(_84_v0))).Cmp(Companion_Default___.Fm0(_91_globalState)) > 0) || (_85_v1) {
		var _92_v7 _dafny.Array
		_ = _92_v7
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
		_ = _nw11
		_92_v7 = _nw11
		var _93_v8 _dafny.Map
		_ = _93_v8
		_93_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v7, _85_v1)
		_93_v8 = (_93_v8).Update(_92_v7, _85_v1)
		_84_v0 = _84_v0
		(_91_globalState).F6 = _dafny.IntOfInt64(-930)
		var _94_v9 _dafny.Array
		_ = _94_v9
		var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw12
		_94_v9 = _nw12
		_94_v9 = _94_v9
		var _95_v10 _dafny.Sequence
		_ = _95_v10
		_95_v10 = _dafny.SeqOf(_84_v0)
		var _96_v11 _dafny.MultiSet
		_ = _96_v11
		_96_v11 = _dafny.MultiSetOf(_84_v0)
		var _97_v12 _dafny.Set
		_ = _97_v12
		_97_v12 = _dafny.SetOf(_96_v11, _96_v11)
		var _98_v13 _dafny.Map
		_ = _98_v13
		_98_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_97_v12).Cardinality(), _84_v0)
		var _99_v14 _dafny.Set
		_ = _99_v14
		_99_v14 = _dafny.SetOf(_dafny.IntOfInt64(320), _dafny.IntOfUint32((_89_v5).Cardinality()), _dafny.IntOfUint32((_95_v10).Cardinality()), _84_v0, (_98_v13).Cardinality())
		_85_v1 = ((_99_v14).Union(_dafny.SetOf(_84_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ryggvk")).Cardinality())))).IsProperSubsetOf(_dafny.SetOf(_84_v0, _dafny.IntOfInt64(-755), _84_v0, _84_v0))
	} else {
		(_91_globalState).F10 = (_dafny.Zero).Minus(_84_v0)
		var _100_v15 _dafny.Map
		_ = _100_v15
		_100_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v1, _84_v0)
		var _101_v16 _dafny.Sequence
		_ = _101_v16
		_101_v16 = _dafny.SeqOf((_86_v2).Cardinality(), (_100_v15).Cardinality())
		_85_v1 = Companion_Default___.Fm1((_101_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.IntOfUint32((_101_v16).Cardinality()))).Uint32()).(_dafny.Int), false, _85_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_86_v2, Companion_Default___.Fm2(_84_v0, _85_v1, _85_v1, _91_globalState)), _dafny.SeqOf(_86_v2)), _91_globalState)
		var _102_v17 _dafny.Map
		_ = _102_v17
		_102_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, Companion_Default___.Fm0(_91_globalState))
		_102_v17 = (_102_v17).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_89_v5, (Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_89_v5).Cardinality()))).Uint32(), Companion_Default___.Fm3(_84_v0, _84_v0, _85_v1, _91_globalState)), _89_v5)).Cardinality()), _84_v0)
		(_91_globalState).F6 = Companion_Default___.SafeDivisionInt(_84_v0, _84_v0)
		_85_v1 = _85_v1
	}
	var _103_v18 _dafny.Map
	_ = _103_v18
	_103_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, _84_v0)
	var _104_v20 _dafny.Sequence
	_ = _104_v20
	_104_v20 = _dafny.SeqOf(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(920), _dafny.IntOfInt64(-802))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _105_v19 _dafny.Int
			_105_v19 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(920)).Cmp(_105_v19) <= 0) && ((_105_v19).Cmp(_dafny.IntOfInt64(-802)) < 0) {
				_coll2.Add((_105_v19).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v1, _85_v1)).Cardinality()), _85_v1)
			}
		}
		return _coll2.ToMap()
	}())
	var _106_v21 _dafny.Map
	_ = _106_v21
	_106_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_103_v18).Cardinality(), _104_v20)
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
	_ = _index20
	(_90_v6).ArraySet1(Companion_Default___.Fm1(_84_v0, true, false, (func() _dafny.Sequence {
		if (_106_v21).Contains(_84_v0) {
			return (_106_v21).Get(_84_v0).(_dafny.Sequence)
		}
		return _104_v20
	})(), _91_globalState), (_index20).Int())
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
	_ = _index21
	(_90_v6).ArraySet1(_85_v1, (_index21).Int())
	var _107_v22 _dafny.CodePoint
	_ = _107_v22
	_107_v22 = _dafny.CodePoint('i')
	var _108_v23 _dafny.Map
	_ = _108_v23
	_108_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("c"), (Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()))).Uint32(), _107_v22), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
	_85_v1 = !((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)) || ((func() bool {
		if (_108_v23).Contains(_89_v5) {
			return (_108_v23).Get(_89_v5).(bool)
		}
		return true
	})())
	var _109_v24 _dafny.Array
	_ = _109_v24
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_0
	var _nw13 _dafny.Array
	_ = _nw13
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw13 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = (func(_110_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_111_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_111_i0, _110_v0)
			}
		})(_84_v0)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw13 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw13).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw13).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_109_v24 = _nw13
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
	_ = _index22
	(_109_v24).ArraySet1(_84_v0, (_index22).Int())
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_109_v24), 0))
	_ = _index23
	(_109_v24).ArraySet1((_dafny.MultiSetOf((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _85_v1)).Cardinality(), (_index23).Int())
	var _112_v25 _dafny.Sequence
	_ = _112_v25
	_112_v25 = _dafny.SeqOf(_84_v0, _84_v0, _84_v0)
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
	_ = _index24
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_109_v24), 0))
	_ = _index25
	var _rhs13 _dafny.Int = Companion_Default___.SafeModuloInt(_84_v0, (_dafny.Zero).Minus(Companion_Default___.Fm0(_91_globalState)))
	_ = _rhs13
	var _rhs14 _dafny.Int = ((_84_v0).Plus((_103_v18).Cardinality())).Times(_84_v0)
	_ = _rhs14
	var _rhs15 _dafny.CodePoint = (_89_v5).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_89_v5).Cardinality()))).Uint32()).(_dafny.CodePoint)
	_ = _rhs15
	var _rhs16 _dafny.Int = (_dafny.Zero).Minus(_84_v0)
	_ = _rhs16
	var _rhs17 _dafny.Sequence = _112_v25
	_ = _rhs17
	var _lhs16 _dafny.Array = _109_v24
	_ = _lhs16
	var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
	_ = _lhs17
	var _lhs18 *GlobalState = _91_globalState
	_ = _lhs18
	var _lhs19 _dafny.Array = _109_v24
	_ = _lhs19
	var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_109_v24), 0))
	_ = _lhs20
	(_lhs16).ArraySet1(_rhs13, (_lhs17).Int())
	_lhs18.F6 = _rhs14
	_107_v22 = _rhs15
	(_lhs19).ArraySet1(_rhs16, (_lhs20).Int())
	_112_v25 = _rhs17
	var _hi0 _dafny.Int = (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)
	_ = _hi0
	for _113_i1 := (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int); _113_i1.Cmp(_hi0) < 0; _113_i1 = _113_i1.Plus(_dafny.One) {
		var _114_v26 _dafny.Set
		_ = _114_v26
		_114_v26 = _dafny.SetOf(_85_v1)
		var _115_v27 _dafny.Map
		_ = _115_v27
		_115_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v1, (_114_v26).Cardinality())
		_115_v27 = (_115_v27).Update(true, _dafny.IntOfInt64(482))
		Companion_Default___.M0((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_112_v25, _112_v25)).Cardinality()))), _91_globalState)
		_87_v3 = _dafny.Companion_Sequence_.Concatenate(_87_v3, _87_v3)
		var _116_v28 _dafny.Map
		_ = _116_v28
		_116_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v3, _109_v24)
		var _117_v29 _dafny.Sequence
		_ = _117_v29
		_117_v29 = _dafny.SeqOf(_109_v24, _109_v24, (func() _dafny.Array {
			if (_116_v28).Contains(_dafny.SeqOf(_85_v1)) {
				return (_116_v28).Get(_dafny.SeqOf(_85_v1)).(_dafny.Array)
			}
			return _109_v24
		})())
		_109_v24 = (_117_v29).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_117_v29).Cardinality()))).Uint32()).(_dafny.Array)
	}
	for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_90_v6), 0))); ; {
		_guard_loop_0, _ok3 := _iter3()
		if !_ok3 {
			break
		}
		var _118_i2 _dafny.Int
		_118_i2 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_118_i2).Sign() != -1) && ((_118_i2).Cmp(_dafny.ArrayLen((_90_v6), 0)) < 0)) {
			(_90_v6).ArraySet1(_85_v1, (_118_i2).Int())
		}
	}
	var _119_v30 _dafny.Map
	_ = _119_v30
	_119_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), _107_v22)
	_119_v30 = (_119_v30).Update(_84_v0, _107_v22)
	if _85_v1 {
		(_91_globalState).F6 = ((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)).Times((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))
		if (!(false)) || (Companion_Default___.Fm1(_dafny.IntOfUint32((_89_v5).Cardinality()), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _104_v20, _91_globalState)) {
			_85_v1 = (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)
			_90_v6 = _90_v6
			_89_v5 = _89_v5
			var _120_v31 *C0
			_ = _120_v31
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))
			_120_v31 = _nw14
			var _121_v32 _dafny.Map
			_ = _121_v32
			_121_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_87_v3, _87_v3), !((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)) || ((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)))
			_121_v32 = (_121_v32).Update(_dafny.Companion_Sequence_.Update(_87_v3, (Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_87_v3).Cardinality()))).Uint32(), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)), !(((_87_v3).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_87_v3).Cardinality()))).Uint32()).(bool)) == ((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))))
		} else {
			var _122_v33 _dafny.Sequence
			_ = _122_v33
			_122_v33 = _dafny.SeqOf(_90_v6)
			Companion_Default___.M0((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _dafny.IntOfUint32((_122_v33).Cardinality()), _91_globalState)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index26
			(_109_v24).ArraySet1(Companion_Default___.Fm0(_91_globalState), (_index26).Int())
			_85_v1 = _85_v1
			var _123_v34 _dafny.Sequence
			_ = _123_v34
			_123_v34 = _dafny.SeqOf(_89_v5, _89_v5)
			var _124_v35 _dafny.Map
			_ = _124_v35
			_124_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), _dafny.SeqOf(_84_v0, (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)))
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index27
			var _rhs18 bool = (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)
			_ = _rhs18
			var _rhs19 _dafny.Int = _dafny.IntOfInt64(602)
			_ = _rhs19
			var _rhs20 bool = !(!((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)))
			_ = _rhs20
			var _rhs21 _dafny.Int = Companion_Default___.SafeDivisionInt(_84_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_123_v34, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.IntOfUint32((_123_v34).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("c"))).Cardinality()))
			_ = _rhs21
			var _rhs22 bool = (_124_v35).Contains(((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_87_v3).Cardinality())))
			_ = _rhs22
			var _lhs21 *GlobalState = _91_globalState
			_ = _lhs21
			var _lhs22 *GlobalState = _91_globalState
			_ = _lhs22
			var _lhs23 _dafny.Array = _90_v6
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _lhs24
			_85_v1 = _rhs18
			_lhs21.F10 = _rhs19
			_85_v1 = _rhs20
			_lhs22.F6 = _rhs21
			(_lhs23).ArraySet1(_rhs22, (_lhs24).Int())
			var _125_v36 *C0
			_ = _125_v36
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(_dafny.IntOfInt64(478))
			_125_v36 = _nw15
		}
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
		_ = _index28
		(_90_v6).ArraySet1(_85_v1, (_index28).Int())
		var _126_v37 _dafny.Map
		_ = _126_v37
		_126_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v5, _90_v6)
		_90_v6 = (func() _dafny.Array {
			if (_126_v37).Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_127_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_128_i3 _dafny.Int) _dafny.CodePoint {
					return _127_v22
				}
			})(_107_v22))), _89_v5), (Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_129_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_130_i3 _dafny.Int) _dafny.CodePoint {
					return _129_v22
				}
			})(_107_v22))), _89_v5)).Cardinality()))).Uint32(), _dafny.CodePoint('c'))) {
				return (_126_v37).Get(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}((func(_131_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_132_i3 _dafny.Int) _dafny.CodePoint {
						return _131_v22
					}
				})(_107_v22))), _89_v5), (Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_133_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_134_i3 _dafny.Int) _dafny.CodePoint {
						return _133_v22
					}
				})(_107_v22))), _89_v5)).Cardinality()))).Uint32(), _dafny.CodePoint('c'))).(_dafny.Array)
			}
			return _90_v6
		})()
		var _135_v38 _dafny.Set
		_ = _135_v38
		_135_v38 = _dafny.SetOf(_112_v25)
		Companion_Default___.M0((_135_v38).IsProperSubsetOf(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-144)))), (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), _91_globalState)
	} else {
		_107_v22 = _dafny.CodePoint('y')
		_89_v5 = _89_v5
		if (_dafny.IntOfUint32((_89_v5).Cardinality())).Cmp(_84_v0) == 0 {
			var _136_v39 _dafny.Array
			_ = _136_v39
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
			_ = _nw16
			_136_v39 = _nw16
			var _137_v40 *C0
			_ = _137_v40
			var _nw17 *C0 = New_C0_()
			_ = _nw17
			_nw17.Ctor__((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))
			_137_v40 = _nw17
			var _138_v41 _dafny.Sequence
			_ = _138_v41
			_138_v41 = _dafny.SeqOf(_137_v40)
			var _139_v42 _dafny.MultiSet
			_ = _139_v42
			_139_v42 = _dafny.MultiSetOf(_138_v41, _138_v41)
			var _140_v43 _dafny.Map
			_ = _140_v43
			_140_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v24, _139_v42)
			var _141_v44 _dafny.Sequence
			_ = _141_v44
			_141_v44 = _dafny.SeqOf(_138_v41)
			var _142_v45 _dafny.Map
			_ = _142_v45
			_142_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _85_v1)
			var _143_v46 _dafny.Map
			_ = _143_v46
			_143_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_84_v0), _142_v45)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_136_v39), 0))
			_ = _index29
			(_136_v39).ArraySet1(((func() _dafny.MultiSet {
				if (_140_v43).Contains(_109_v24) {
					return (_140_v43).Get(_109_v24).(_dafny.MultiSet)
				}
				return _dafny.MultiSetFromSeq(_141_v44)
			})()).Update(_138_v41, Companion_Default___.Abs((_143_v46).Cardinality())), (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_136_v39), 0))
			_ = _index30
			(_136_v39).ArraySet1(_139_v42, (_index30).Int())
			Companion_Default___.M0((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), _91_globalState)
			_112_v25 = _dafny.Companion_Sequence_.Concatenate(_112_v25, _dafny.Companion_Sequence_.Concatenate(_112_v25, _112_v25))
			Companion_Default___.M0(_85_v1, _84_v0, _91_globalState)
			_103_v18 = (Companion_D1_.Create_DC1_(_103_v18)).Dtor_cf1()
		} else {
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index31
			(_90_v6).ArraySet1(false, (_index31).Int())
			var _144_v47 _dafny.Array
			_ = _144_v47
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
			_ = _nw18
			_144_v47 = _nw18
			var _145_v48 _dafny.Map
			_ = _145_v48
			_145_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v47, _dafny.IntOfUint32((_112_v25).Cardinality()))
			_145_v48 = (_145_v48).Update(_144_v47, (_dafny.IntOfUint32((_87_v3).Cardinality())).Minus((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)))
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index32
			(_109_v24).ArraySet1(_dafny.IntOfInt64(878), (_index32).Int())
			_88_v4 = _88_v4
			var _146_v49 _dafny.Set
			_ = _146_v49
			_146_v49 = _dafny.SetOf((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))
			var _147_v50 _dafny.Map
			_ = _147_v50
			_147_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
			var _148_v51 _dafny.Map
			_ = _148_v51
			_148_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v49, ((_147_v50).Update((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))).Cardinality())
			var _149_v52 *C0
			_ = _149_v52
			var _nw19 *C0 = New_C0_()
			_ = _nw19
			_nw19.Ctor__((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))
			_149_v52 = _nw19
			var _150_v53 _dafny.Sequence
			_ = _150_v53
			_150_v53 = _dafny.SeqOf(_149_v52, _149_v52)
			var _151_v54 _dafny.Sequence
			_ = _151_v54
			_151_v54 = _dafny.SeqOf(_148_v51, _148_v51)
			var _152_v55 _dafny.Array
			_ = _152_v55
			var _nwElement0_1 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_84_v0), _84_v0)
			_ = _nwElement0_1
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(19))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_1, 0)
			(_nw20).ArraySet1((_148_v51).Update(_146_v49, _dafny.IntOfUint32((_150_v53).Cardinality())), 1)
			(_nw20).ArraySet1((_148_v51).Merge(_148_v51), 2)
			(_nw20).ArraySet1(_148_v51, 3)
			(_nw20).ArraySet1((_151_v54).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_151_v54).Cardinality()))).Uint32()).(_dafny.Map), 4)
			(_nw20).ArraySet1((_148_v51).Merge(_148_v51), 5)
			(_nw20).ArraySet1(_148_v51, 6)
			(_nw20).ArraySet1(_148_v51, 7)
			(_nw20).ArraySet1(Companion_Default___.Fm6((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _85_v1, (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _85_v1, _91_globalState), 8)
			(_nw20).ArraySet1(_148_v51, 9)
			(_nw20).ArraySet1((_148_v51).Merge(_148_v51), 10)
			(_nw20).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v49, _dafny.IntOfInt64(735)), 11)
			(_nw20).ArraySet1(_148_v51, 12)
			(_nw20).ArraySet1(_148_v51, 13)
			(_nw20).ArraySet1((_148_v51).Merge(_148_v51), 14)
			(_nw20).ArraySet1(_148_v51, 15)
			(_nw20).ArraySet1((_148_v51).Update(_146_v49, (_dafny.Zero).Minus((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))), 16)
			(_nw20).ArraySet1(_148_v51, 17)
			(_nw20).ArraySet1(_148_v51, 18)
			_152_v55 = _nw20
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_152_v55), 0))
			_ = _index33
			(_152_v55).ArraySet1(_148_v51, (_index33).Int())
			var _153_v56 _dafny.Sequence
			_ = _153_v56
			_153_v56 = _dafny.SeqOf(_89_v5)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index34
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_152_v55), 0))
			_ = _index35
			var _rhs23 _dafny.Int = _dafny.IntOfUint32((_153_v56).Cardinality())
			_ = _rhs23
			var _rhs24 bool = (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)
			_ = _rhs24
			var _rhs25 _dafny.Int = Companion_Default___.Fm0(_91_globalState)
			_ = _rhs25
			var _rhs26 _dafny.Map = (_148_v51).Merge(_148_v51)
			_ = _rhs26
			var _lhs25 _dafny.Array = _109_v24
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _lhs26
			var _lhs27 *GlobalState = _91_globalState
			_ = _lhs27
			var _lhs28 _dafny.Array = _152_v55
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_152_v55), 0))
			_ = _lhs29
			(_lhs25).ArraySet1(_rhs23, (_lhs26).Int())
			_85_v1 = _rhs24
			_lhs27.F6 = _rhs25
			(_lhs28).ArraySet1(_rhs26, (_lhs29).Int())
		}
		var _154_v57 _dafny.Map
		_ = _154_v57
		_154_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v1, _84_v0)
		_154_v57 = _154_v57
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
		_ = _index36
		(_90_v6).ArraySet1((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_index36).Int())
	}
	var _hi1 _dafny.Int = _dafny.IntOfInt64(998)
	_ = _hi1
	for _155_i4 := (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int); _155_i4.Cmp(_hi1) < 0; _155_i4 = _155_i4.Plus(_dafny.One) {
		var _156_v58 _dafny.Array
		_ = _156_v58
		var _nw21 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
		_ = _nw21
		_156_v58 = _nw21
		var _157_v59 *C0
		_ = _157_v59
		var _nw22 *C0 = New_C0_()
		_ = _nw22
		_nw22.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ook")).Cardinality()))
		_157_v59 = _nw22
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_156_v58), 0))
		_ = _index37
		(_156_v58).ArraySet1(_157_v59, (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_156_v58), 0))
		_ = _index38
		(_156_v58).ArraySet1(_157_v59, (_index38).Int())
		var _158_v60 _dafny.MultiSet
		_ = _158_v60
		_158_v60 = _dafny.MultiSetOf(_85_v1)
		var _159_v61 D1
		_ = _159_v61
		_159_v61 = Companion_D1_.Create_DC2_((_157_v59).F14(), (_158_v60).Cardinality(), _dafny.IntOfInt64(250), _155_i4)
		var _160_v62 D1
		_ = _160_v62
		_160_v62 = Companion_D1_.Create_DC3_(_159_v61)
		var _161_v63 D1
		_ = _161_v63
		_161_v63 = Companion_D1_.Create_DC3_(_159_v61)
		var _pat_let_tv3 = _159_v61
		_ = _pat_let_tv3
		_161_v63 = func(_pat_let4_0 D1) D1 {
			return func(_162_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let5_0 D1) D1 {
					return func(_163_dt__update_hcf6_h0 D1) D1 {
						return Companion_D1_.Create_DC3_(_163_dt__update_hcf6_h0)
					}(_pat_let5_0)
				}(_pat_let_tv3)
			}(_pat_let4_0)
		}(_161_v63)
		var _164_v64 _dafny.Map
		_ = _164_v64
		_164_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_156_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_156_v58), 0))).Int()).(*C0), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
		if (func() bool {
			if (_164_v64).Contains((_156_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_156_v58), 0))).Int()).(*C0)) {
				return (_164_v64).Get((_156_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_156_v58), 0))).Int()).(*C0)).(bool)
			}
			return (_85_v1) && ((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
		})() {
			var _165_v65 _dafny.Set
			_ = _165_v65
			_165_v65 = _dafny.SetOf((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
			var _166_v66 _dafny.MultiSet
			_ = _166_v66
			_166_v66 = _dafny.MultiSetOf(_90_v6)
			var _nwElement0_2 _dafny.Int = _155_i4
			_ = _nwElement0_2
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(14))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_2, 0)
			(_nw23).ArraySet1((_157_v59).F14(), 1)
			(_nw23).ArraySet1(_155_i4, 2)
			(_nw23).ArraySet1(Companion_Default___.SafeDivisionInt((_158_v60).Cardinality(), _155_i4), 3)
			(_nw23).ArraySet1(_dafny.IntOfInt64(-490), 4)
			(_nw23).ArraySet1(Companion_Default___.SafeModuloInt((_157_v59).F14(), _dafny.IntOfInt64(461)), 5)
			(_nw23).ArraySet1(_84_v0, 6)
			(_nw23).ArraySet1((_157_v59).F14(), 7)
			(_nw23).ArraySet1((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), 8)
			(_nw23).ArraySet1((_165_v65).Cardinality(), 9)
			(_nw23).ArraySet1((_157_v59).F14(), 10)
			(_nw23).ArraySet1((_157_v59).F14(), 11)
			(_nw23).ArraySet1((_dafny.Zero).Minus((_157_v59).F14()), 12)
			(_nw23).ArraySet1(((_166_v66).Cardinality()).Times(_84_v0), 13)
			_109_v24 = _nw23
			Companion_Default___.M0(!(_85_v1), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(556))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_167_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_168_i5 _dafny.Int) _dafny.CodePoint {
					return _167_v22
				}
			})(_107_v22)))).Cardinality()), _91_globalState)
			Companion_Default___.M0((!(_85_v1)) && (_85_v1), (_157_v59).F14(), _91_globalState)
			_157_v59 = _157_v59
			var _169_v67 _dafny.MultiSet
			_ = _169_v67
			_169_v67 = _dafny.MultiSetOf(_107_v22, _107_v22)
			var _170_v68 *C0
			_ = _170_v68
			var _nw24 *C0 = New_C0_()
			_ = _nw24
			_nw24.Ctor__(((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)).Times((func() _dafny.Int {
				if (_169_v67).Contains(_107_v22) {
					return (_169_v67).Multiplicity(_107_v22)
				}
				return _155_i4
			})()))
			_170_v68 = _nw24
		} else {
			var _171_v69 D1
			_ = _171_v69
			_171_v69 = Companion_D1_.Create_DC1_((_103_v18).Merge(_103_v18))
			var _172_v70 _dafny.Set
			_ = _172_v70
			_172_v70 = _dafny.SetOf((_156_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_156_v58), 0))).Int()).(*C0))
			var _173_v71 D2
			_ = _173_v71
			_173_v71 = Companion_D2_.Create_DC4_(_172_v70)
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index39
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index40
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index41
			var _rhs27 bool = !(Companion_Default___.Fm1((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _dafny.SeqOf(_86_v2), _91_globalState))
			_ = _rhs27
			var _rhs28 bool = ((func() _dafny.Set {
				if (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool) {
					return _172_v70
				}
				return (_173_v71).Dtor_cf7()
			})()).IsDisjointFrom(_172_v70)
			_ = _rhs28
			var _rhs29 bool = _85_v1
			_ = _rhs29
			var _rhs30 D1 = _171_v69
			_ = _rhs30
			var _rhs31 bool = (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)
			_ = _rhs31
			var _lhs30 _dafny.Array = _90_v6
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _lhs31
			var _lhs32 _dafny.Array = _90_v6
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _lhs33
			var _lhs34 _dafny.Array = _90_v6
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _lhs35
			(_lhs30).ArraySet1(_rhs27, (_lhs31).Int())
			_85_v1 = _rhs28
			(_lhs32).ArraySet1(_rhs29, (_lhs33).Int())
			_171_v69 = _rhs30
			(_lhs34).ArraySet1(_rhs31, (_lhs35).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index42
			(_109_v24).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_84_v0, _dafny.IntOfInt64(371))), (_index42).Int())
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_1
			var _nw25 _dafny.Array
			_ = _nw25
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw25 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_174_v6 _dafny.Array, _175_v1 bool) func(_dafny.Int) _dafny.Int {
					return func(_176_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_176_i6, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_174_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_174_v6), 0))).Int()).(bool), _175_v1, (_174_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_174_v6), 0))).Int()).(bool), true)).Cardinality())))
					}
				})(_90_v6, _85_v1)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw25 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw25).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw25).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_109_v24 = _nw25
			var _177_v72 *C0
			_ = _177_v72
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__((_dafny.Zero).Minus((_84_v0).Minus((_157_v59).F14())))
			_177_v72 = _nw26
			Companion_Default___.M0(_85_v1, Companion_Default___.SafeDivisionInt((_157_v59).F14(), (_157_v59).F14()), _91_globalState)
		}
		_85_v1 = (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)
	}
	var _178_v73 _dafny.Array
	_ = _178_v73
	var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
	_ = _nw27
	_178_v73 = _nw27
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
	_ = _index43
	var _rhs32 _dafny.Array = _178_v73
	_ = _rhs32
	var _rhs33 bool = _85_v1
	_ = _rhs33
	var _lhs36 _dafny.Array = _90_v6
	_ = _lhs36
	var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
	_ = _lhs37
	_178_v73 = _rhs32
	(_lhs36).ArraySet1(_rhs33, (_lhs37).Int())
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
	_ = _index44
	(_109_v24).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))), (_index44).Int())
	var _179_v74 *C0
	_ = _179_v74
	var _nw28 *C0 = New_C0_()
	_ = _nw28
	_nw28.Ctor__(_84_v0)
	_179_v74 = _nw28
	var _180_v75 _dafny.Sequence
	_ = _180_v75
	_180_v75 = _dafny.SeqOf(_179_v74)
	var _181_v76 _dafny.MultiSet
	_ = _181_v76
	_181_v76 = _dafny.MultiSetOf(_179_v74)
	var _182_v77 _dafny.Set
	_ = _182_v77
	_182_v77 = _dafny.SetOf(_dafny.IntOfUint32((_180_v75).Cardinality()), _84_v0, (_dafny.Zero).Minus((_179_v74).F14()), _dafny.IntOfUint32((_87_v3).Cardinality()), (_181_v76).Cardinality())
	_85_v1 = ((_182_v77).Union(_182_v77)).IsDisjointFrom((_182_v77).Union(_182_v77))
	if _85_v1 {
		var _183_v78 _dafny.Map
		_ = _183_v78
		_183_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v1, (_179_v74).F14())
		var _184_v79 _dafny.Sequence
		_ = _184_v79
		_184_v79 = _dafny.SeqOf(_183_v78)
		var _185_v80 _dafny.Array
		_ = _185_v80
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_2
		var _nw29 _dafny.Array
		_ = _nw29
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw29 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Sequence = (func(_186_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_187_i7 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(700))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}((func(_188_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_189_i8 _dafny.Int) _dafny.Sequence {
							return _188_v3
						}
					})(_186_v3)))
				}
			})(_87_v3)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw29 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw29).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw29).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_185_v80 = _nw29
		var _190_v81 _dafny.Sequence
		_ = _190_v81
		_190_v81 = _dafny.SeqOf(_dafny.SeqOf(_85_v1, _85_v1))
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_185_v80), 0))
		_ = _index45
		(_185_v80).ArraySet1(_190_v81, (_index45).Int())
		var _191_v82 _dafny.Map
		_ = _191_v82
		_191_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v22, _107_v22)
		var _192_v83 _dafny.Set
		_ = _192_v83
		_192_v83 = _dafny.SetOf(_191_v82)
		var _193_v84 D2
		_ = _193_v84
		_193_v84 = Companion_D2_.Create_DC6_(!((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)), _85_v1, _107_v22)
		var _194_v86 _dafny.Set
		_ = _194_v86
		_194_v86 = _dafny.SetOf(_107_v22)
		var _195_v87 D3
		_ = _195_v87
		_195_v87 = Companion_D3_.Create_DC8_(_194_v86)
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_185_v80), 0))
		_ = _index46
		var _rhs34 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_184_v79, (Companion_Default___.SafeIndex((_179_v74).F14(), _dafny.IntOfUint32((_184_v79).Cardinality()))).Uint32(), _183_v78)
		_ = _rhs34
		var _rhs35 _dafny.Sequence = _190_v81
		_ = _rhs35
		var _rhs36 _dafny.Set = (_192_v83).Intersection(func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter4 := _dafny.Iterate((_dafny.MultiSetOf(_191_v82, func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter5 := _dafny.Iterate(((_195_v87).Dtor_cf14()).Elements()); ; {
					_compr_4, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _196_v85 _dafny.CodePoint
					_196_v85 = interface{}(_compr_4).(_dafny.CodePoint)
					if ((_195_v87).Dtor_cf14()).Contains(_196_v85) {
						_coll4.Add(_196_v85, _107_v22)
					}
				}
				return _coll4.ToMap()
			}())).Elements()); ; {
				_compr_3, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _197_v88 _dafny.Map
				_197_v88 = interface{}(_compr_3).(_dafny.Map)
				if (_dafny.MultiSetOf(_191_v82, func() _dafny.Map {
					var _coll5 = _dafny.NewMapBuilder()
					_ = _coll5
					for _iter6 := _dafny.Iterate(((_195_v87).Dtor_cf14()).Elements()); ; {
						_compr_5, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _196_v85 _dafny.CodePoint
						_196_v85 = interface{}(_compr_5).(_dafny.CodePoint)
						if ((_195_v87).Dtor_cf14()).Contains(_196_v85) {
							_coll5.Add(_196_v85, _107_v22)
						}
					}
					return _coll5.ToMap()
				}())).Contains(_197_v88) {
					_coll3.Add(_197_v88)
				}
			}
			return _coll3.ToSet()
		}())
		_ = _rhs36
		var _rhs37 D2 = _193_v84
		_ = _rhs37
		var _lhs38 _dafny.Array = _185_v80
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_185_v80), 0))
		_ = _lhs39
		_184_v79 = _rhs34
		(_lhs38).ArraySet1(_rhs35, (_lhs39).Int())
		_192_v83 = _rhs36
		_193_v84 = _rhs37
		Companion_Default___.M0((Companion_D2_.Create_DC5_(_85_v1, (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))).Dtor_cf9(), Companion_Default___.SafeDivisionInt((_179_v74).F14(), (_179_v74).F14()), _91_globalState)
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
		_ = _index47
		(_109_v24).ArraySet1((_179_v74).F14(), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
		_ = _index48
		var _rhs38 bool = (_dafny.IntOfUint32((_89_v5).Cardinality())).Cmp((_103_v18).Cardinality()) != 0
		_ = _rhs38
		var _rhs39 _dafny.Map = func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-729), _dafny.IntOfInt64(42))); ; {
				_compr_6, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _198_v89 _dafny.Int
				_198_v89 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(-729)).Cmp(_198_v89) <= 0) && ((_198_v89).Cmp(_dafny.IntOfInt64(42)) < 0) {
					_coll6.Add(Companion_Default___.SafeModuloInt(_198_v89, (_179_v74).F14()), _85_v1)
				}
			}
			return _coll6.ToMap()
		}()
		_ = _rhs39
		var _rhs40 _dafny.Int = _dafny.IntOfInt64(516)
		_ = _rhs40
		var _rhs41 _dafny.Int = (_dafny.IntOfUint32((_89_v5).Cardinality())).Times((_179_v74).F14())
		_ = _rhs41
		var _rhs42 _dafny.Int = (_dafny.Zero).Minus(_84_v0)
		_ = _rhs42
		var _lhs40 _dafny.Array = _109_v24
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
		_ = _lhs41
		var _lhs42 *GlobalState = _91_globalState
		_ = _lhs42
		var _lhs43 *GlobalState = _91_globalState
		_ = _lhs43
		_85_v1 = _rhs38
		_86_v2 = _rhs39
		(_lhs40).ArraySet1(_rhs40, (_lhs41).Int())
		_lhs42.F10 = _rhs41
		_lhs43.F6 = _rhs42
		_86_v2 = (_86_v2).Update((_179_v74).F14(), _85_v1)
	} else {
		if _85_v1 {
			var _199_v90 _dafny.Sequence
			_ = _199_v90
			_199_v90 = _dafny.SeqOf(_182_v77)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index49
			var _rhs43 bool = (_179_v74).Fm4(_85_v1, _91_globalState)
			_ = _rhs43
			var _rhs44 bool = _85_v1
			_ = _rhs44
			var _rhs45 _dafny.Set = (_199_v90).Select((Companion_Default___.SafeIndex(((_179_v74).F14()).Plus((_179_v74).F14()), _dafny.IntOfUint32((_199_v90).Cardinality()))).Uint32()).(_dafny.Set)
			_ = _rhs45
			var _rhs46 _dafny.CodePoint = _107_v22
			_ = _rhs46
			var _lhs44 _dafny.Array = _90_v6
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _lhs45
			(_lhs44).ArraySet1(_rhs43, (_lhs45).Int())
			_85_v1 = _rhs44
			_182_v77 = _rhs45
			_107_v22 = _rhs46
			var _200_v91 _dafny.Sequence
			_ = _200_v91
			_200_v91 = _dafny.SeqOf(_dafny.SetOf(_107_v22))
			var _201_v92 _dafny.Map
			_ = _201_v92
			_201_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v74, (_200_v91).Select((Companion_Default___.SafeIndex((_179_v74).F14(), _dafny.IntOfUint32((_200_v91).Cardinality()))).Uint32()).(_dafny.Set))
			var _202_v93 D3
			_ = _202_v93
			_202_v93 = Companion_D3_.Create_DC8_((func() _dafny.Set {
				if (_201_v92).Contains(_179_v74) {
					return (_201_v92).Get(_179_v74).(_dafny.Set)
				}
				return _dafny.SetOf(_107_v22)
			})())
			var _203_v94 _dafny.MultiSet
			_ = _203_v94
			_203_v94 = _dafny.MultiSetOf(_202_v93)
			var _204_v95 _dafny.MultiSet
			_ = _204_v95
			_204_v95 = _dafny.MultiSetOf(_85_v1, (_85_v1) || (_85_v1), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _85_v1, (_203_v94).Equals(_203_v94))
			_204_v95 = ((_204_v95).Intersection(_204_v95)).Difference(_204_v95)
			Companion_Default___.M0(_85_v1, (_179_v74).F14(), _91_globalState)
			var _205_v96 _dafny.Array
			_ = _205_v96
			var _nw30 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
			_ = _nw30
			_205_v96 = _nw30
			_205_v96 = _205_v96
			var _206_v97 _dafny.MultiSet
			_ = _206_v97
			_206_v97 = _dafny.MultiSetOf(_dafny.IntOfInt64(235))
			(_91_globalState).F10 = (_112_v25).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_206_v97).Contains((_179_v74).F14()) {
					return (_206_v97).Multiplicity((_179_v74).F14())
				}
				return (_179_v74).F14()
			})(), _dafny.IntOfUint32((_112_v25).Cardinality()))).Uint32()).(_dafny.Int)
		} else {
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw31
			_109_v24 = _nw31
			var _207_v98 _dafny.MultiSet
			_ = _207_v98
			_207_v98 = _dafny.MultiSetOf(_107_v22)
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index50
			(_109_v24).ArraySet1((((_207_v98).Union(_dafny.MultiSetFromSeq(_89_v5))).Cardinality()).Minus((_179_v74).F14()), (_index50).Int())
			var _208_v99 _dafny.Set
			_ = _208_v99
			_208_v99 = _dafny.SetOf((_180_v75).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_180_v75).Cardinality()))).Uint32()).(*C0))
			var _209_v100 D2
			_ = _209_v100
			_209_v100 = Companion_D2_.Create_DC4_((_208_v99).Intersection(_208_v99))
			_209_v100 = _209_v100
			var _210_v101 D2
			_ = _210_v101
			_210_v101 = Companion_D2_.Create_DC5_(!((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)), _85_v1)
			var _211_v102 _dafny.Map
			_ = _211_v102
			_211_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_109_v24, _109_v24, _109_v24), (func() D2 {
				if (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool) {
					return _210_v101
				}
				return _210_v101
			})())
			_211_v102 = (_211_v102).Update(_dafny.SetOf(_109_v24, _109_v24), _210_v101)
			Companion_Default___.M0((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_179_v74).F14(), _91_globalState)
		}
		_109_v24 = _109_v24
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
		_ = _index51
		(_109_v24).ArraySet1((_dafny.Zero).Minus((_179_v74).F14()), (_index51).Int())
		var _212_v103 _dafny.MultiSet
		_ = _212_v103
		_212_v103 = _dafny.MultiSetOf(_85_v1)
		var _213_v104 _dafny.Map
		_ = _213_v104
		_213_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v2, !((_87_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_212_v103).Contains((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)) {
				return (_212_v103).Multiplicity((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
			}
			return (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)
		})(), _dafny.IntOfUint32((_87_v3).Cardinality()))).Uint32()).(bool)))
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
		_ = _index52
		var _rhs47 bool = (_212_v103).IsProperSubsetOf(_212_v103)
		_ = _rhs47
		var _rhs48 _dafny.Int = Companion_Default___.SafeModuloInt(_84_v0, (_179_v74).F14())
		_ = _rhs48
		var _rhs49 _dafny.Map = (_213_v104).Merge(_213_v104)
		_ = _rhs49
		var _lhs46 _dafny.Array = _90_v6
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
		_ = _lhs47
		var _lhs48 *GlobalState = _91_globalState
		_ = _lhs48
		(_lhs46).ArraySet1(_rhs47, (_lhs47).Int())
		_lhs48.F6 = _rhs48
		_213_v104 = _rhs49
		if _85_v1 {
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index53
			(_90_v6).ArraySet1((_84_v0).Cmp((_179_v74).F14()) != 0, (_index53).Int())
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index54
			(_109_v24).ArraySet1((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), (_index54).Int())
			_85_v1 = (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)
			var _214_v105 _dafny.Sequence
			_ = _214_v105
			_214_v105 = _dafny.SeqOf(_112_v25)
			_214_v105 = _214_v105
			_85_v1 = (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)
		} else {
			_85_v1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(611), _84_v0), _112_v25)
			_85_v1 = !((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)) || (_85_v1)
			var _215_v106 _dafny.Array
			_ = _215_v106
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
			_ = _nw32
			_215_v106 = _nw32
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_215_v106), 0))
			_ = _index55
			(_215_v106).ArraySet1(_181_v76, (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_215_v106), 0))
			_ = _index56
			(_215_v106).ArraySet1(((_dafny.MultiSetOf(_179_v74)).Union(_dafny.MultiSetOf(_179_v74))).Intersection((_181_v76).Update(_179_v74, Companion_Default___.Abs((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)))), (_index56).Int())
			(_91_globalState).F10 = ((_179_v74).F14()).Minus((_179_v74).F14())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index57
			(_90_v6).ArraySet1(false, (_index57).Int())
		}
	}
	var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
	_ = _index58
	(_90_v6).ArraySet1((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), (_index58).Int())
	var _216_v107 D2
	_ = _216_v107
	_216_v107 = Companion_D2_.Create_DC5_(_85_v1, (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
	var _217_v108 D2
	_ = _217_v108
	_217_v108 = Companion_D2_.Create_DC7_(_216_v107)
	var _source2 D2 = (func() D2 {
		if _85_v1 {
			return _217_v108
		}
		return _217_v108
	})()
	_ = _source2
	if _source2.Is_DC5() {
		var _218___mcc_h0 bool = _source2.Get_().(D2_DC5).Cf8
		_ = _218___mcc_h0
		var _219___mcc_h1 bool = _source2.Get_().(D2_DC5).Cf9
		_ = _219___mcc_h1
		var _220_cf9 bool = _219___mcc_h1
		_ = _220_cf9
		var _221_cf8 bool = _218___mcc_h0
		_ = _221_cf8
		var _222_v109 _dafny.Map
		_ = _222_v109
		_222_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v22, _179_v74)
		_222_v109 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), _179_v74)).Update(_107_v22, _179_v74)
		var _223_v110 _dafny.Map
		_ = _223_v110
		_223_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, Companion_D3_.Create_DC10_(Companion_D2_.Create_DC5_(_220_cf9, _221_cf8), _107_v22, _dafny.IntOfUint32((_89_v5).Cardinality())))
		var _224_v111 D2
		_ = _224_v111
		_224_v111 = Companion_D2_.Create_DC5_(_220_cf9, _221_cf8)
		var _225_v112 D3
		_ = _225_v112
		_225_v112 = Companion_D3_.Create_DC10_(_224_v111, _107_v22, (_179_v74).F14())
		_223_v110 = (_223_v110).Update((_179_v74).F14(), _225_v112)
		_220_cf9 = !(!_dafny.Companion_Sequence_.Equal(_89_v5, _dafny.UnicodeSeqOfUtf8Bytes("fglj")))
		_85_v1 = (_179_v74).Fm5(_91_globalState)
	} else if _source2.Is_DC6() {
		var _226___mcc_h2 bool = _source2.Get_().(D2_DC6).Cf10
		_ = _226___mcc_h2
		var _227___mcc_h3 bool = _source2.Get_().(D2_DC6).Cf11
		_ = _227___mcc_h3
		var _228___mcc_h4 _dafny.CodePoint = _source2.Get_().(D2_DC6).Cf12
		_ = _228___mcc_h4
		var _229_cf12 _dafny.CodePoint = _228___mcc_h4
		_ = _229_cf12
		var _230_cf11 bool = _227___mcc_h3
		_ = _230_cf11
		var _231_cf10 bool = _226___mcc_h2
		_ = _231_cf10
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_3
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Int = (func(_232_v25 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_233_i9 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_233_i9, _dafny.IntOfUint32((_232_v25).Cardinality()))
				}
			})(_112_v25)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw33 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw33).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw33).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_109_v24 = _nw33
		var _234_v113 *C0
		_ = _234_v113
		var _nw34 *C0 = New_C0_()
		_ = _nw34
		_nw34.Ctor__((_179_v74).F14())
		_234_v113 = _nw34
		var _235_v114 _dafny.Map
		_ = _235_v114
		_235_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v1, _230_cf11)
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
		_ = _index59
		(_90_v6).ArraySet1((func() bool {
			if (_235_v114).Contains((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)) {
				return (_235_v114).Get((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool)).(bool)
			}
			return false
		})(), (_index59).Int())
		Companion_Default___.M0(_85_v1, _dafny.IntOfInt64(644), _91_globalState)
	} else if _source2.Is_DC4() {
		var _236___mcc_h5 _dafny.Set = _source2.Get_().(D2_DC4).Cf7
		_ = _236___mcc_h5
		var _237_cf7 _dafny.Set = _236___mcc_h5
		_ = _237_cf7
		if _85_v1 {
			var _238_v115 D3
			_ = _238_v115
			_238_v115 = Companion_D3_.Create_DC9_((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))
			var _239_v116 _dafny.Map
			_ = _239_v116
			_239_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_179_v74).F14()).Plus(_84_v0)), _238_v115)
			var _pat_let_tv4 = _85_v1
			_ = _pat_let_tv4
			_239_v116 = (_239_v116).Update((_179_v74).F14(), func(_pat_let6_0 D3) D3 {
				return func(_240_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let7_0 bool) D3 {
						return func(_241_dt__update_hcf15_h0 bool) D3 {
							return Companion_D3_.Create_DC9_(_241_dt__update_hcf15_h0)
						}(_pat_let7_0)
					}(!(_pat_let_tv4))
				}(_pat_let6_0)
			}(_238_v115))
			var _242_v117 *C0
			_ = _242_v117
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__((_179_v74).F14())
			_242_v117 = _nw35
			_85_v1 = _85_v1
			var _243_v118 _dafny.Map
			_ = _243_v118
			_243_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, _109_v24)
			_243_v118 = ((_243_v118).Update((_242_v117).F14(), _109_v24)).Merge((_243_v118).Merge(_243_v118))
			_179_v74 = _179_v74
		} else {
			_103_v18 = (_103_v18).Update(_dafny.IntOfUint32((_dafny.SeqOf(_85_v1, (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool))).Cardinality()), (_179_v74).F14())
			var _244_v119 _dafny.Array
			_ = _244_v119
			var _nwElement0_3 _dafny.Array = _178_v73
			_ = _nwElement0_3
			var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(2))
			_ = _nw36
			(_nw36).ArraySet1(_nwElement0_3, 0)
			(_nw36).ArraySet1(_178_v73, 1)
			_244_v119 = _nw36
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_244_v119), 0))
			_ = _index60
			(_244_v119).ArraySet1((func() _dafny.Array {
				if (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool) {
					return _178_v73
				}
				return _178_v73
			})(), (_index60).Int())
			var _245_v120 D2
			_ = _245_v120
			_245_v120 = Companion_D2_.Create_DC4_(_237_cf7)
			var _246_v121 _dafny.Set
			_ = _246_v121
			_246_v121 = _dafny.SetOf(_245_v120, _245_v120, _245_v120, _245_v120, Companion_D2_.Create_DC4_(_237_cf7))
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_244_v119), 0))
			_ = _index61
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index62
			var _rhs50 _dafny.Array = _178_v73
			_ = _rhs50
			var _rhs51 _dafny.Int = (_179_v74).F14()
			_ = _rhs51
			var _rhs52 bool = ((_179_v74).F14()).Cmp(_84_v0) == 0
			_ = _rhs52
			var _rhs53 _dafny.Int = ((_246_v121).Union(_246_v121)).Cardinality()
			_ = _rhs53
			var _rhs54 _dafny.Int = Companion_Default___.SafeDivisionInt((_179_v74).F14(), _dafny.IntOfUint32((_89_v5).Cardinality()))
			_ = _rhs54
			var _lhs49 _dafny.Array = _244_v119
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_244_v119), 0))
			_ = _lhs50
			var _lhs51 _dafny.Array = _109_v24
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _lhs52
			var _lhs53 *GlobalState = _91_globalState
			_ = _lhs53
			var _lhs54 *GlobalState = _91_globalState
			_ = _lhs54
			(_lhs49).ArraySet1(_rhs50, (_lhs50).Int())
			(_lhs51).ArraySet1(_rhs51, (_lhs52).Int())
			_85_v1 = _rhs52
			_lhs53.F6 = _rhs53
			_lhs54.F6 = _rhs54
			_84_v0 = Companion_Default___.Fm0(_91_globalState)
			_107_v22 = _107_v22
			var _247_v122 *C0
			_ = _247_v122
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__((_179_v74).F14())
			_247_v122 = _nw37
		}
		var _248_v123 *C0
		_ = _248_v123
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(806))))
		_248_v123 = _nw38
		if ((_dafny.MultiSetOf(_179_v74)).Intersection(_181_v76)).IsProperSubsetOf(_181_v76) {
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index63
			(_109_v24).ArraySet1((_dafny.IntOfUint32((_89_v5).Cardinality())).Minus((_179_v74).F14()), (_index63).Int())
			(_91_globalState).F6 = (_179_v74).F14()
			(_91_globalState).F6 = (_179_v74).F14()
			var _249_v124 _dafny.Array
			_ = _249_v124
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(28))
			_ = _nw39
			_249_v124 = _nw39
			var _250_v125 _dafny.MultiSet
			_ = _250_v125
			_250_v125 = _dafny.MultiSetOf(_249_v124, _249_v124, _249_v124, _249_v124, _249_v124)
			var _rhs55 _dafny.Array = _90_v6
			_ = _rhs55
			var _rhs56 _dafny.Int = (func() _dafny.Int {
				if (_250_v125).Contains((Companion_D4_.Create_DC11_(_249_v124)).Dtor_cf19()) {
					return (_250_v125).Multiplicity((Companion_D4_.Create_DC11_(_249_v124)).Dtor_cf19())
				}
				return ((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)).Plus((_179_v74).F14())
			})()
			_ = _rhs56
			var _rhs57 _dafny.Int = (_248_v123).F14()
			_ = _rhs57
			var _rhs58 _dafny.Int = (_179_v74).F14()
			_ = _rhs58
			var _lhs55 *GlobalState = _91_globalState
			_ = _lhs55
			var _lhs56 *GlobalState = _91_globalState
			_ = _lhs56
			var _lhs57 *GlobalState = _91_globalState
			_ = _lhs57
			_90_v6 = _rhs55
			_lhs55.F6 = _rhs56
			_lhs56.F6 = _rhs57
			_lhs57.F6 = _rhs58
			var _251_v126 _dafny.MultiSet
			_ = _251_v126
			_251_v126 = _dafny.MultiSetOf(_90_v6)
			(_91_globalState).F6 = Companion_Default___.SafeDivisionInt(((_179_v74).F14()).Plus((_251_v126).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)).Times(_84_v0))))
		} else {
			Companion_Default___.M0(_85_v1, (_dafny.Zero).Minus(((_248_v123).F14()).Plus(_dafny.IntOfInt64(254))), _91_globalState)
			_89_v5 = _89_v5
			_85_v1 = (_dafny.IntOfUint32((_89_v5).Cardinality())).Cmp((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)) <= 0
			var _252_v127 _dafny.MultiSet
			_ = _252_v127
			_252_v127 = _dafny.MultiSetOf((_179_v74).F14())
			var _rhs59 _dafny.Int = ((func() _dafny.Int {
				if (_252_v127).Contains((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)) {
					return (_252_v127).Multiplicity((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))
				}
				return (_248_v123).F14()
			})()).Plus(((_179_v74).F14()).Times(_dafny.IntOfUint32((Companion_Default___.Fm7((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _107_v22, _91_globalState)).Cardinality())))
			_ = _rhs59
			var _rhs60 _dafny.Int = (_dafny.Zero).Minus((_248_v123).F14())
			_ = _rhs60
			var _lhs58 *GlobalState = _91_globalState
			_ = _lhs58
			var _lhs59 *GlobalState = _91_globalState
			_ = _lhs59
			_lhs58.F6 = _rhs59
			_lhs59.F6 = _rhs60
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw40
			_109_v24 = _nw40
		}
		var _253_v128 _dafny.Map
		_ = _253_v128
		_253_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v74, _112_v25)
		var _254_v129 _dafny.Sequence
		_ = _254_v129
		_254_v129 = _dafny.SeqOf((_253_v128).Merge(_253_v128))
		_253_v128 = (_254_v129).Select((Companion_Default___.SafeIndex((_179_v74).F14(), _dafny.IntOfUint32((_254_v129).Cardinality()))).Uint32()).(_dafny.Map)
	} else {
		var _255___mcc_h6 D2 = _source2.Get_().(D2_DC7).Cf13
		_ = _255___mcc_h6
		var _256_cf13 D2 = _255___mcc_h6
		_ = _256_cf13
		var _nw41 *C0 = New_C0_()
		_ = _nw41
		_nw41.Ctor__(Companion_Default___.SafeDivisionInt((_179_v74).F14(), _84_v0))
		_179_v74 = _nw41
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
		_ = _index64
		var _rhs61 bool = ((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)).Cmp((_179_v74).F14()) >= 0
		_ = _rhs61
		var _rhs62 _dafny.Int = (_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int)
		_ = _rhs62
		var _lhs60 _dafny.Array = _90_v6
		_ = _lhs60
		var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
		_ = _lhs61
		var _lhs62 *GlobalState = _91_globalState
		_ = _lhs62
		(_lhs60).ArraySet1(_rhs61, (_lhs61).Int())
		_lhs62.F6 = _rhs62
		_85_v1 = _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_85_v1), Companion_Default___.Fm8((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _91_globalState))
		if (_87_v3).Select((Companion_Default___.SafeIndex((_179_v74).F14(), _dafny.IntOfUint32((_87_v3).Cardinality()))).Uint32()).(bool) {
			var _257_v130 D3
			_ = _257_v130
			_257_v130 = Companion_D3_.Create_DC8_(_dafny.SetOf(_107_v22, _107_v22, _107_v22, _107_v22, _dafny.CodePoint('v')))
			var _258_v131 _dafny.Set
			_ = _258_v131
			_258_v131 = _dafny.SetOf(_107_v22, _107_v22)
			var _pat_let_tv5 = _258_v131
			_ = _pat_let_tv5
			_257_v130 = func(_pat_let8_0 D3) D3 {
				return func(_259_dt__update__tmp_h2 D3) D3 {
					return func(_pat_let9_0 _dafny.Set) D3 {
						return func(_260_dt__update_hcf14_h0 _dafny.Set) D3 {
							return Companion_D3_.Create_DC8_(_260_dt__update_hcf14_h0)
						}(_pat_let9_0)
					}(_pat_let_tv5)
				}(_pat_let8_0)
			}(_257_v130)
			_257_v130 = (func() D3 {
				if (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool) {
					return _257_v130
				}
				return _257_v130
			})()
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_178_v73), 0))
			_ = _index65
			(_178_v73).ArraySet1(_89_v5, (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_178_v73), 0))
			_ = _index66
			(_178_v73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_89_v5, _89_v5), (_index66).Int())
			var _261_v132 _dafny.Array
			_ = _261_v132
			var _nw42 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
			_ = _nw42
			_261_v132 = _nw42
			_261_v132 = _261_v132
			(_91_globalState).F6 = (_179_v74).F14()
		} else {
			(_91_globalState).F6 = (_dafny.Zero).Minus((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int))
			var _262_v133 _dafny.Sequence
			_ = _262_v133
			_262_v133 = _dafny.SeqOf(_109_v24, _109_v24, _109_v24, _109_v24)
			var _263_v134 _dafny.MultiSet
			_ = _263_v134
			_263_v134 = _dafny.MultiSetOf((_262_v133).Select((Companion_Default___.SafeIndex((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_262_v133).Cardinality()))).Uint32()).(_dafny.Array))
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))
			_ = _index67
			(_90_v6).ArraySet1((_dafny.MultiSetOf(_109_v24)).IsProperSubsetOf(_263_v134), (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))
			_ = _index68
			(_109_v24).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_89_v5, Companion_Default___.Fm7((_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), Companion_Default___.Fm3((_109_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_109_v24), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-30), (_90_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_90_v6), 0))).Int()).(bool), _91_globalState), _91_globalState))).Cardinality()), (_index68).Int())
			(_91_globalState).F10 = (_dafny.Zero).Minus(_84_v0)
			var _264_v135 _dafny.Map
			_ = _264_v135
			_264_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v77, Companion_Default___.Fm0(_91_globalState))
			_264_v135 = (_264_v135).Update(_182_v77, (_179_v74).F14())
		}
	}
	Companion_Default___.M0(false, (_179_v74).F14(), _91_globalState)
	_dafny.Print(_84_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(457), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_87_v3, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(924), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_89_v5.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v6).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F4()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(924), _dafny.SeqOf(false)).UpdateUnsafe(_dafny.IntOfInt64(457), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_91_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F9().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_91_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_v18).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(457), _dafny.IntOfInt64(457))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_104_v20, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap())))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v21).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap()))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_107_v22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_108_v23).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("i"), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v24).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_112_v25, _dafny.SeqOf(_dafny.IntOfInt64(457), _dafny.IntOfInt64(457), _dafny.IntOfInt64(457))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v30).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.CodePoint('n')).UpdateUnsafe(_dafny.IntOfInt64(457), _dafny.CodePoint('n'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v74).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_180_v75).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_v76).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v77).Equals(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(457), _dafny.IntOfInt64(-457))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_216_v107).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_216_v107).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_217_v108).Dtor_cf13()).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_217_v108).Dtor_cf13()).Dtor_cf9())
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

func (CompanionStruct_D0_) Default() bool {
	return false
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	Cf2 _dafny.Int
	Cf3 _dafny.Int
	Cf4 _dafny.Int
	Cf5 _dafny.Int
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 _dafny.Int, Cf5 _dafny.Int) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4, Cf5}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Map
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Map) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

type D1_DC3 struct {
	Cf6 D1
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 D1) D1 {
	return D1{D1_DC3{Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf1() _dafny.Map {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf6() D1 {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6)
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

type D2_DC5 struct {
	Cf8 bool
	Cf9 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf8 bool, Cf9 bool) D2 {
	return D2{D2_DC5{Cf8, Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC6 struct {
	Cf10 bool
	Cf11 bool
	Cf12 _dafny.CodePoint
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 bool, Cf11 bool, Cf12 _dafny.CodePoint) D2 {
	return D2{D2_DC6{Cf10, Cf11, Cf12}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC4 struct {
	Cf7 _dafny.Set
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf7 _dafny.Set) D2 {
	return D2{D2_DC4{Cf7}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC7 struct {
	Cf13 D2
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf13 D2) D2 {
	return D2{D2_DC7{Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(false, false)
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) Dtor_cf12() _dafny.CodePoint {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf7() _dafny.Set {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) Dtor_cf13() D2 {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf7.Equals(data2.Cf7)
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
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

type D3_DC9 struct {
	Cf15 bool
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf15 bool) D3 {
	return D3{D3_DC9{Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf16 D2
	Cf17 _dafny.CodePoint
	Cf18 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 D2, Cf17 _dafny.CodePoint, Cf18 _dafny.Int) D3 {
	return D3{D3_DC10{Cf16, Cf17, Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf14 _dafny.Set
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 _dafny.Set) D3 {
	return D3{D3_DC8{Cf14}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(false)
}

func (_this D3) Dtor_cf15() bool {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) Dtor_cf16() D2 {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf14() _dafny.Set {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf15 == data2.Cf15
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16) && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
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

type D4_DC12 struct {
	Cf20 bool
	Cf21 bool
	Cf22 bool
	Cf23 _dafny.Sequence
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf20 bool, Cf21 bool, Cf22 bool, Cf23 _dafny.Sequence) D4 {
	return D4{D4_DC12{Cf20, Cf21, Cf22, Cf23}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC11 struct {
	Cf19 _dafny.Array
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf19 _dafny.Array) D4 {
	return D4{D4_DC11{Cf19}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(false, false, false, _dafny.EmptySeq)
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) Dtor_cf21() bool {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf22() bool {
	return _this.Get_().(D4_DC12).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D4_DC12).Cf23
}

func (_this D4) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf19
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23.Equals(data2.Cf23)
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf19 == data2.Cf19
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

// Definition of class GlobalState
type GlobalState struct {
	F6   _dafny.Int
	F10  _dafny.Int
	_f0  bool
	_f1  bool
	_f2  bool
	_f3  bool
	_f4  _dafny.Map
	_f5  _dafny.Int
	_f7  bool
	_f8  _dafny.Sequence
	_f9  _dafny.Sequence
	_f11 _dafny.Array
	_f12 bool
	_f13 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F6 = _dafny.Zero
	_this.F10 = _dafny.Zero
	_this._f0 = false
	_this._f1 = false
	_this._f2 = false
	_this._f3 = false
	_this._f4 = _dafny.EmptyMap
	_this._f5 = _dafny.Zero
	_this._f7 = false
	_this._f8 = _dafny.EmptySeq
	_this._f9 = _dafny.EmptySeq
	_this._f11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f12 = false
	_this._f13 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 bool, f3 bool, f4 _dafny.Map, f5 _dafny.Int, f6 _dafny.Int, f7 bool, f8 _dafny.Sequence, f9 _dafny.Sequence, f10 _dafny.Int, f11 _dafny.Array, f12 bool, f13 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
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
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Map {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Sequence {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F11() _dafny.Array {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f14 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f14 = _dafny.Zero
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

func (_this *C0) Ctor__(f14 _dafny.Int) {
	{
		(_this)._f14 = f14
	}
}
func (_this *C0) Fm4(p0 bool, globalState *GlobalState) bool {
	{
		return (true)
	}
}
func (_this *C0) Fm5(globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf(!(true), false, !(false), true)).IsDisjointFrom((_dafny.MultiSetOf(true, false)).Union(_dafny.MultiSetOf(true)))
	}
}
func (_this *C0) F14() _dafny.Int {
	{
		return _this._f14
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
