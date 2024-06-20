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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true, true)).Union(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_()
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	var _source0 D0 = Companion_D0_.Create_DC0_(true)
	_ = _source0
	if _source0.Is_DC1() {
		return _dafny.UnicodeSeqOfUtf8Bytes("ttioiues")
	} else {
		var _0___mcc_h0 bool = _source0.Get_().(D0_DC0).Cf0
		_ = _0___mcc_h0
		var _1_cf0 bool = _0___mcc_h0
		_ = _1_cf0
		return _dafny.UnicodeSeqOfUtf8Bytes("p")
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(-413)
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("meafobh")).Cardinality())
	}))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, globalState *GlobalState) D0 {
	var _source1 D0 = Companion_D0_.Create_DC1_()
	_ = _source1
	if _source1.Is_DC1() {
		return Companion_D0_.Create_DC0_(true)
	} else {
		var _3___mcc_h0 bool = _source1.Get_().(D0_DC0).Cf0
		_ = _3___mcc_h0
		var _4_cf0 bool = _3___mcc_h0
		_ = _4_cf0
		return Companion_D0_.Create_DC0_(_4_cf0)
	}
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var r2 bool = false
	_ = r2
	var _5_v0 _dafny.Array
	_ = _5_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
	_ = _nw0
	_5_v0 = _nw0
	for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_5_v0), 0))); ; {
		_guard_loop_0, _ok0 := _iter0()
		if !_ok0 {
			break
		}
		var _6_i0 _dafny.Int
		_6_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_6_i0).Sign() != -1) && ((_6_i0).Cmp(_dafny.ArrayLen((_5_v0), 0)) < 0)) {
			(_5_v0).ArraySet1((func() bool {
				if (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("icxytex")).Cardinality())).Cmp(_dafny.IntOfInt64(410)) != 0 {
					return p0
				}
				return p0
			})(), (_6_i0).Int())
		}
	}
	var _7_v1 _dafny.Int
	_ = _7_v1
	_7_v1 = _dafny.IntOfInt64(14)
	var _8_v2 _dafny.Sequence
	_ = _8_v2
	_8_v2 = _dafny.SeqOf(_7_v1)
	var _9_v3 *C0
	_ = _9_v3
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__(_8_v2)
	_9_v3 = _nw1
	var _10_v4 _dafny.Map
	_ = _10_v4
	_10_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v3, p1)
	var _11_v5 _dafny.Map
	_ = _11_v5
	_11_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v1, _10_v4)
	var _12_v6 _dafny.Sequence
	_ = _12_v6
	_12_v6 = _dafny.SeqOf(_11_v5)
	var _13_v7 D1
	_ = _13_v7
	_13_v7 = Companion_D1_.Create_DC4_((_12_v6).Select((Companion_Default___.SafeIndex(_7_v1, _dafny.IntOfUint32((_12_v6).Cardinality()))).Uint32()).(_dafny.Map), _7_v1)
	var _source2 D0 = Companion_Default___.Fm2((_13_v7).Dtor_cf5(), globalState)
	_ = _source2
	if _source2.Is_DC1() {
		r2 = false
		var _14_v8 _dafny.Sequence
		_ = _14_v8
		_14_v8 = _dafny.SeqOf(p0, p0)
		var _15_v9 _dafny.Map
		_ = _15_v9
		_15_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_14_v8).Select((Companion_Default___.SafeIndex(_7_v1, _dafny.IntOfUint32((_14_v8).Cardinality()))).Uint32()).(bool), _7_v1)
		var _16_v10 _dafny.Array
		_ = _16_v10
		var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw2
		_16_v10 = _nw2
		var _17_v11 _dafny.MultiSet
		_ = _17_v11
		_17_v11 = _dafny.MultiSetOf(_7_v1, Companion_Default___.Fm4(_7_v1, _7_v1, p0, p0, globalState), (_15_v9).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_16_v10, _16_v10, _16_v10)).Cardinality()))
		var _18_v12 D0
		_ = _18_v12
		_18_v12 = Companion_D0_.Create_DC0_((_dafny.MultiSetOf(_7_v1)).Equals(_17_v11))
		var _19_v13 _dafny.Sequence
		_ = _19_v13
		_19_v13 = _dafny.UnicodeSeqOfUtf8Bytes("avymisae")
		var _20_v14 _dafny.Map
		_ = _20_v14
		_20_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _19_v13)
		var _21_v15 _dafny.Map
		_ = _21_v15
		_21_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), (func() _dafny.Sequence {
			if (_20_v14).Contains(p0) {
				return (_20_v14).Get(p0).(_dafny.Sequence)
			}
			return _19_v13
		})())
		var _22_v16 _dafny.Array
		_ = _22_v16
		var _nwElement0_0 _dafny.Sequence = _19_v13
		_ = _nwElement0_0
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(16))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_0, 0)
		(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_19_v13, _19_v13), 1)
		(_nw3).ArraySet1(_19_v13, 2)
		(_nw3).ArraySet1(_19_v13, 3)
		(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_19_v13, _19_v13), 4)
		(_nw3).ArraySet1((func() _dafny.Sequence {
			if (_21_v15).Contains(p1) {
				return (_21_v15).Get(p1).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Update(_19_v13, (Companion_Default___.SafeIndex(_7_v1, _dafny.IntOfUint32((_19_v13).Cardinality()))).Uint32(), _dafny.CodePoint('w'))
		})(), 5)
		(_nw3).ArraySet1(_19_v13, 6)
		(_nw3).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(490))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_23_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})), 7)
		(_nw3).ArraySet1(_19_v13, 8)
		(_nw3).ArraySet1((func() _dafny.Sequence {
			if p0 {
				return _19_v13
			}
			return _19_v13
		})(), 9)
		(_nw3).ArraySet1(_19_v13, 10)
		(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_19_v13, _19_v13), 11)
		(_nw3).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rwsgvn"), 12)
		(_nw3).ArraySet1(Companion_Default___.Fm3(globalState), 13)
		(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_19_v13, _19_v13), 14)
		(_nw3).ArraySet1(_19_v13, 15)
		_22_v16 = _nw3
		var _24_v17 _dafny.CodePoint
		_ = _24_v17
		_24_v17 = _dafny.CodePoint('v')
		var _25_v18 _dafny.Array
		_ = _25_v18
		var _nwElement0_1 _dafny.CodePoint = _24_v17
		_ = _nwElement0_1
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(4))
		_ = _nw4
		(_nw4).ArraySet1CodePoint(_nwElement0_1, 0)
		(_nw4).ArraySet1CodePoint(_24_v17, 1)
		(_nw4).ArraySet1CodePoint(_24_v17, 2)
		(_nw4).ArraySet1CodePoint(_24_v17, 3)
		_25_v18 = _nw4
		var _26_v19 _dafny.Sequence
		_ = _26_v19
		_26_v19 = _dafny.SeqOf(_9_v3)
		var _27_v20 _dafny.Array
		_ = _27_v20
		var _nwElement0_2 *C0 = (_26_v19).Select((Companion_Default___.SafeIndex(_7_v1, _dafny.IntOfUint32((_26_v19).Cardinality()))).Uint32()).(*C0)
		_ = _nwElement0_2
		var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(10))
		_ = _nw5
		(_nw5).ArraySet1(_nwElement0_2, 0)
		(_nw5).ArraySet1(_9_v3, 1)
		(_nw5).ArraySet1(_9_v3, 2)
		(_nw5).ArraySet1(_9_v3, 3)
		(_nw5).ArraySet1(_9_v3, 4)
		(_nw5).ArraySet1(_9_v3, 5)
		(_nw5).ArraySet1(_9_v3, 6)
		(_nw5).ArraySet1(_9_v3, 7)
		(_nw5).ArraySet1(_9_v3, 8)
		(_nw5).ArraySet1(_9_v3, 9)
		_27_v20 = _nw5
		var _28_v21 _dafny.Map
		_ = _28_v21
		_28_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_14_v8).Cardinality()), _27_v20)
		var _29_v22 _dafny.Map
		_ = _29_v22
		_29_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, (func() _dafny.Array {
			if (_28_v21).Contains(_7_v1) {
				return (_28_v21).Get(_7_v1).(_dafny.Array)
			}
			return _27_v20
		})())
		var _30_v23 _dafny.Map
		_ = _30_v23
		_30_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v18, _29_v22)
		var _31_v24 _dafny.Sequence
		_ = _31_v24
		_31_v24 = _dafny.SeqOf(_29_v22)
		var _32_v25 _dafny.Array
		_ = _32_v25
		var _nwElement0_3 _dafny.Map = (func() _dafny.Map {
			if (_30_v23).Contains(_25_v18) {
				return (_30_v23).Get(_25_v18).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, _27_v20)
		})()
		_ = _nwElement0_3
		var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(23))
		_ = _nw6
		(_nw6).ArraySet1(_nwElement0_3, 0)
		(_nw6).ArraySet1(_29_v22, 1)
		(_nw6).ArraySet1((_29_v22).Merge(_29_v22), 2)
		(_nw6).ArraySet1(_29_v22, 3)
		(_nw6).ArraySet1(_29_v22, 4)
		(_nw6).ArraySet1(_29_v22, 5)
		(_nw6).ArraySet1(_29_v22, 6)
		(_nw6).ArraySet1(_29_v22, 7)
		(_nw6).ArraySet1((_29_v22).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, _27_v20)), 8)
		(_nw6).ArraySet1((Companion_D3_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, _27_v20))).Dtor_cf7(), 9)
		(_nw6).ArraySet1(_29_v22, 10)
		(_nw6).ArraySet1(_29_v22, 11)
		(_nw6).ArraySet1(_29_v22, 12)
		(_nw6).ArraySet1(_29_v22, 13)
		(_nw6).ArraySet1(_29_v22, 14)
		(_nw6).ArraySet1((_29_v22).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, _27_v20)), 15)
		(_nw6).ArraySet1(_29_v22, 16)
		(_nw6).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, _27_v20), 17)
		(_nw6).ArraySet1((_31_v24).Select((Companion_Default___.SafeIndex(_7_v1, _dafny.IntOfUint32((_31_v24).Cardinality()))).Uint32()).(_dafny.Map), 18)
		(_nw6).ArraySet1((_29_v22).Merge(_29_v22), 19)
		(_nw6).ArraySet1((_29_v22).Merge(_29_v22), 20)
		(_nw6).ArraySet1(_29_v22, 21)
		(_nw6).ArraySet1((_29_v22).Update(_5_v0, _27_v20), 22)
		_32_v25 = _nw6
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_32_v25), 0))
		_ = _index0
		(_32_v25).ArraySet1(_29_v22, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_32_v25), 0))
		_ = _index1
		var _rhs0 bool = p0
		_ = _rhs0
		var _rhs1 D0 = _18_v12
		_ = _rhs1
		var _rhs2 _dafny.Array = _22_v16
		_ = _rhs2
		var _rhs3 _dafny.Map = (_29_v22).Merge(_29_v22)
		_ = _rhs3
		var _lhs0 _dafny.Array = _32_v25
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_32_v25), 0))
		_ = _lhs1
		r2 = _rhs0
		_18_v12 = _rhs1
		_22_v16 = _rhs2
		(_lhs0).ArraySet1(_rhs3, (_lhs1).Int())
		var _33_v26 _dafny.Array
		_ = _33_v26
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
		_ = _nw7
		_33_v26 = _nw7
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_33_v26), 0))
		_ = _index2
		(_33_v26).ArraySet1(_16_v10, (_index2).Int())
		var _34_v27 _dafny.Map
		_ = _34_v27
		_34_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _17_v11)
		var _35_v28 _dafny.MultiSet
		_ = _35_v28
		_35_v28 = _dafny.MultiSetOf(p1, p0)
		var _36_v29 _dafny.Sequence
		_ = _36_v29
		_36_v29 = _dafny.SeqOf(_16_v10)
		var _37_v30 _dafny.Map
		_ = _37_v30
		_37_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v11).IsProperSubsetOf((func() _dafny.MultiSet {
			if (_34_v27).Contains(p0) {
				return (_34_v27).Get(p0).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf(_7_v1, (_35_v28).Cardinality(), _dafny.IntOfInt64(57), _7_v1)
		})()), (_36_v29).Select((Companion_Default___.SafeIndex(_7_v1, _dafny.IntOfUint32((_36_v29).Cardinality()))).Uint32()).(_dafny.Array))
		var _38_v31 _dafny.Array
		_ = _38_v31
		_38_v31 = _16_v10
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_33_v26), 0))
		_ = _index3
		(_33_v26).ArraySet1((func() _dafny.Array {
			if (_37_v30).Contains(p1) {
				return (_37_v30).Get(p1).(_dafny.Array)
			}
			return (_38_v31)
		})(), (_index3).Int())
		r2 = !(!(!(p0)))
	} else {
		var _39___mcc_h0 bool = _source2.Get_().(D0_DC0).Cf0
		_ = _39___mcc_h0
		var _40_cf0 bool = _39___mcc_h0
		_ = _40_cf0
		var _41_v32 _dafny.Sequence
		_ = _41_v32
		_41_v32 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-50))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}((func(_42_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_43_i2 _dafny.Int) _dafny.Int {
				return _42_v1
			}
		})(_7_v1)))
		_8_v2 = _dafny.Companion_Sequence_.Concatenate((_9_v3).F0(), (_41_v32))
		var _44_v33 _dafny.CodePoint
		_ = _44_v33
		_44_v33 = _dafny.CodePoint('y')
		_44_v33 = _44_v33
		r1 = (_7_v1).Minus(_dafny.IntOfInt64(-38))
		var _45_v34 _dafny.Set
		_ = _45_v34
		_45_v34 = _dafny.SetOf(_13_v7)
		var _46_v35 _dafny.Map
		_ = _46_v35
		_46_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v3, (_45_v34).Intersection(_45_v34))
		var _47_v36 _dafny.Sequence
		_ = _47_v36
		_47_v36 = _dafny.SeqOf(_45_v34, _45_v34, _45_v34, _45_v34, _45_v34)
		_46_v35 = (_46_v35).Update(_9_v3, (_47_v36).Select((Companion_Default___.SafeIndex(_7_v1, _dafny.IntOfUint32((_47_v36).Cardinality()))).Uint32()).(_dafny.Set))
	}
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_5_v0), 0))
	_ = _index4
	(_5_v0).ArraySet1(p0, (_index4).Int())
	var _48_v37 _dafny.Sequence
	_ = _48_v37
	_48_v37 = _dafny.UnicodeSeqOfUtf8Bytes("mexshnevu")
	var _49_v38 _dafny.MultiSet
	_ = _49_v38
	_49_v38 = _dafny.MultiSetOf(Companion_Default___.Fm1(p1, p0, _dafny.IntOfUint32((_dafny.SeqOf(_8_v2)).Cardinality()), globalState), p0)
	var _50_v39 _dafny.Sequence
	_ = _50_v39
	_50_v39 = _dafny.SeqOf(p0, p0)
	var _51_v40 _dafny.MultiSet
	_ = _51_v40
	_51_v40 = _dafny.MultiSetOf(_9_v3)
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_5_v0), 0))
	_ = _index5
	var _rhs4 bool = p0
	_ = _rhs4
	var _rhs5 bool = !_dafny.Companion_Sequence_.Contains(_50_v39, (_49_v38).IsSubsetOf(_49_v38))
	_ = _rhs5
	var _rhs6 bool = (_51_v40).IsProperSubsetOf(_51_v40)
	_ = _rhs6
	var _rhs7 _dafny.Sequence = _48_v37
	_ = _rhs7
	var _lhs2 _dafny.Array = _5_v0
	_ = _lhs2
	var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_5_v0), 0))
	_ = _lhs3
	r2 = _rhs4
	r2 = _rhs5
	(_lhs2).ArraySet1(_rhs6, (_lhs3).Int())
	_48_v37 = _rhs7
	var _52_v41 *C0
	_ = _52_v41
	var _nw8 *C0 = New_C0_()
	_ = _nw8
	_nw8.Ctor__(_dafny.Companion_Sequence_.Concatenate(_8_v2, (_9_v3).F0()))
	_52_v41 = _nw8
	var _53_v42 _dafny.Array
	_ = _53_v42
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(9)
	_ = _len0_0
	var _nw9 _dafny.Array
	_ = _nw9
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw9 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Map = (func(_54_v0 _dafny.Array) func(_dafny.Int) _dafny.Map {
			return func(_55_i3 _dafny.Int) _dafny.Map {
				return (func() _dafny.Map {
					if false {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_54_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_54_v0), 0))).Int()).(bool), Companion_D0_.Create_DC1_())
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_54_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_54_v0), 0))).Int()).(bool), Companion_D0_.Create_DC1_())
				})()
			}
		})(_5_v0)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw9 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw9).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw9).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_53_v42 = _nw9
	var _56_v43 D0
	_ = _56_v43
	_56_v43 = Companion_D0_.Create_DC1_()
	var _57_v44 _dafny.Map
	_ = _57_v44
	_57_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _56_v43)
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_53_v42), 0))
	_ = _index6
	(_53_v42).ArraySet1(_57_v44, (_index6).Int())
	var _pat_let_tv0 = _57_v44
	_ = _pat_let_tv0
	var _pat_let_tv1 = _57_v44
	_ = _pat_let_tv1
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_53_v42), 0))
	_ = _index7
	(_53_v42).ArraySet1(func(_source3 D0) _dafny.Map {
		if _source3.Is_DC1() {
			return _pat_let_tv0
		} else {
			var _58___mcc_h1 bool = _source3.Get_().(D0_DC0).Cf0
			_ = _58___mcc_h1
			var _59_cf0 bool = _58___mcc_h1
			_ = _59_cf0
			return _pat_let_tv1
		}
	}(Companion_Default___.Fm7((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_5_v0), 0))).Int()).(bool), globalState)), (_index7).Int())
	var _60_v45 _dafny.Array
	_ = _60_v45
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
	_ = _nw10
	_60_v45 = _nw10
	var _61_v46 _dafny.Array
	_ = _61_v46
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(26)
	_ = _len0_1
	var _nw11 _dafny.Array
	_ = _nw11
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw11 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Map = (func(_62_v1 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_63_i4 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_62_v1, (_dafny.MultiSetOf(_62_v1)).Cardinality())).Cardinality(), _62_v1)
			}
		})(_7_v1)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw11 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw11).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw11).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_61_v46 = _nw11
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_60_v45), 0))
	_ = _index8
	(_60_v45).ArraySet1(_61_v46, (_index8).Int())
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_60_v45), 0))
	_ = _index9
	(_60_v45).ArraySet1(_61_v46, (_index9).Int())
	r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(139), _7_v1, _7_v1), _8_v2)
	r1 = _7_v1
	r2 = p0
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _64_globalState *GlobalState
	_ = _64_globalState
	var _nw12 *GlobalState = New_GlobalState_()
	_ = _nw12
	_nw12.Ctor__()
	_64_globalState = _nw12
	var _65_v0 _dafny.Array
	_ = _65_v0
	var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
	_ = _nw13
	_65_v0 = _nw13
	var _66_v1 _dafny.Set
	_ = _66_v1
	_66_v1 = _dafny.SetOf(_65_v0)
	_66_v1 = _66_v1
	var _67_v2 _dafny.Int
	_ = _67_v2
	_67_v2 = _dafny.IntOfInt64(98)
	var _68_v3 bool
	_ = _68_v3
	_68_v3 = true
	var _69_v4 _dafny.Map
	_ = _69_v4
	_69_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_68_v3, _68_v3, _68_v3)).Cardinality(), _67_v2)
	var _70_v6 D0
	_ = _70_v6
	_70_v6 = Companion_D0_.Create_DC0_(false)
	var _rhs8 _dafny.Int = _67_v2
	_ = _rhs8
	var _rhs9 _dafny.Map = (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-725), _dafny.IntOfInt64(-518))); ; {
			_compr_0, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _71_v5 _dafny.Int
			_71_v5 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-725)).Cmp(_71_v5) <= 0) && ((_71_v5).Cmp(_dafny.IntOfInt64(-518)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_71_v5, _67_v2), (_dafny.Zero).Minus((_dafny.Zero).Minus(_67_v2)))
			}
		}
		return _coll0.ToMap()
	}()).Merge(_69_v4)
	_ = _rhs9
	var _rhs10 bool = (_68_v3) == ((_70_v6).Dtor_cf0())
	_ = _rhs10
	_67_v2 = _rhs8
	_69_v4 = _rhs9
	_68_v3 = _rhs10
	for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_65_v0), 0))); ; {
		_guard_loop_1, _ok2 := _iter2()
		if !_ok2 {
			break
		}
		var _72_i0 _dafny.Int
		_72_i0 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_72_i0).Sign() != -1) && ((_72_i0).Cmp(_dafny.ArrayLen((_65_v0), 0)) < 0)) {
			(_65_v0).ArraySet1(_68_v3, (_72_i0).Int())
		}
	}
	var _73_v7 _dafny.CodePoint
	_ = _73_v7
	_73_v7 = _dafny.CodePoint('y')
	var _74_v8 _dafny.Map
	_ = _74_v8
	_74_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v7, _68_v3)
	var _hi0 _dafny.Int = _67_v2
	_ = _hi0
	for _75_i1 := (_74_v8).Cardinality(); _75_i1.Cmp(_hi0) < 0; _75_i1 = _75_i1.Plus(_dafny.One) {
		var _76_v9 D0
		_ = _76_v9
		_76_v9 = Companion_D0_.Create_DC1_()
		_76_v9 = _76_v9
		var _77_v10 _dafny.Set
		_ = _77_v10
		_77_v10 = _dafny.SetOf(_68_v3, _68_v3)
		var _78_v11 _dafny.Sequence
		_ = _78_v11
		var _79_v12 _dafny.Int
		_ = _79_v12
		var _80_v13 bool
		_ = _80_v13
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		var _out2 bool
		_ = _out2
		_out0, _out1, _out2 = Companion_Default___.M0(!((_77_v10).IsSubsetOf(Companion_Default___.Fm0(_64_globalState))), !(true) || (true), _64_globalState)
		_78_v11 = _out0
		_79_v12 = _out1
		_80_v13 = _out2
		var _rhs11 bool = Companion_Default___.Fm1(_80_v13, _80_v13, _79_v12, _64_globalState)
		_ = _rhs11
		var _rhs12 _dafny.CodePoint = _73_v7
		_ = _rhs12
		var _rhs13 bool = _80_v13
		_ = _rhs13
		_68_v3 = _rhs11
		_73_v7 = _rhs12
		_68_v3 = _rhs13
		_77_v10 = _77_v10
	}
	var _81_v14 _dafny.Sequence
	_ = _81_v14
	_81_v14 = _dafny.UnicodeSeqOfUtf8Bytes("braqfma")
	var _82_v15 D0
	_ = _82_v15
	_82_v15 = Companion_D0_.Create_DC1_()
	var _83_v16 _dafny.Array
	_ = _83_v16
	var _nwElement0_4 D0 = _82_v15
	_ = _nwElement0_4
	var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
	_ = _nw14
	(_nw14).ArraySet1(_nwElement0_4, 0)
	(_nw14).ArraySet1(_82_v15, 1)
	(_nw14).ArraySet1(_82_v15, 2)
	(_nw14).ArraySet1(_82_v15, 3)
	(_nw14).ArraySet1(_82_v15, 4)
	(_nw14).ArraySet1(Companion_Default___.Fm2(_67_v2, _64_globalState), 5)
	(_nw14).ArraySet1(_82_v15, 6)
	(_nw14).ArraySet1(Companion_D0_.Create_DC1_(), 7)
	(_nw14).ArraySet1(_82_v15, 8)
	(_nw14).ArraySet1(_82_v15, 9)
	(_nw14).ArraySet1(_82_v15, 10)
	(_nw14).ArraySet1(_82_v15, 11)
	_83_v16 = _nw14
	var _84_v17 _dafny.Map
	_ = _84_v17
	_84_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v16, _68_v3)
	var _rhs14 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_81_v14, _81_v14), _dafny.UnicodeSeqOfUtf8Bytes("dvtinys"))
	_ = _rhs14
	var _rhs15 bool = _68_v3
	_ = _rhs15
	var _rhs16 bool = _68_v3
	_ = _rhs16
	var _rhs17 _dafny.Map = (_84_v17).Merge(_84_v17)
	_ = _rhs17
	var _rhs18 D0 = _70_v6
	_ = _rhs18
	_81_v14 = _rhs14
	_68_v3 = _rhs15
	_68_v3 = _rhs16
	_84_v17 = _rhs17
	_70_v6 = _rhs18
	var _85_v18 _dafny.MultiSet
	_ = _85_v18
	_85_v18 = _dafny.MultiSetOf(_81_v14, _81_v14, _81_v14)
	var _rhs19 bool = true
	_ = _rhs19
	var _rhs20 _dafny.Int = (_67_v2).Plus(((_85_v18).Cardinality()).Plus(_67_v2))
	_ = _rhs20
	_68_v3 = _rhs19
	_67_v2 = _rhs20
	var _86_i2 _dafny.Int
	_ = _86_i2
	_86_i2 = _dafny.Zero
	{
		for (_68_v3) == (!(_68_v3)) {
			{
				if (_86_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_86_i2 = (_86_i2).Plus(_dafny.One)
				_67_v2 = (_dafny.IntOfInt64(175)).Plus(_67_v2)
				var _87_v19 _dafny.Sequence
				_ = _87_v19
				var _88_v20 _dafny.Int
				_ = _88_v20
				var _89_v21 bool
				_ = _89_v21
				var _out3 _dafny.Sequence
				_ = _out3
				var _out4 _dafny.Int
				_ = _out4
				var _out5 bool
				_ = _out5
				_out3, _out4, _out5 = Companion_Default___.M0(false, _68_v3, _64_globalState)
				_87_v19 = _out3
				_88_v20 = _out4
				_89_v21 = _out5
				var _90_v22 _dafny.MultiSet
				_ = _90_v22
				_90_v22 = _dafny.MultiSetOf(_89_v21)
				_89_v21 = (_67_v2).Cmp(Companion_Default___.SafeModuloInt(_88_v20, (_90_v22).Cardinality())) != 0
				_89_v21 = _89_v21
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_81_v14 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm3(_64_globalState), _dafny.UnicodeSeqOfUtf8Bytes("ttjivi"))
	if !(_68_v3) {
		_68_v3 = _68_v3
		var _91_v23 _dafny.Map
		_ = _91_v23
		_91_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v2, _68_v3)
		var _92_v24 _dafny.Map
		_ = _92_v24
		_92_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v3, _67_v2)
		var _93_v25 _dafny.Sequence
		_ = _93_v25
		_93_v25 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v3, (_91_v23).Cardinality()), _92_v24)
		var _94_v26 _dafny.Map
		_ = _94_v26
		_94_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_Default___.Fm4(_67_v2, _dafny.IntOfUint32((_93_v25).Cardinality()), _68_v3, _68_v3, _64_globalState)).Minus(_dafny.IntOfInt64(-707)))
		_92_v24 = (_92_v24).Update(_68_v3, (_67_v2).Minus(_67_v2))
		var _95_v27 _dafny.Sequence
		_ = _95_v27
		_95_v27 = _dafny.SeqOf(_68_v3)
		_68_v3 = (_95_v27).Select((Companion_Default___.SafeIndex(_67_v2, _dafny.IntOfUint32((_95_v27).Cardinality()))).Uint32()).(bool)
		_67_v2 = _dafny.IntOfInt64(640)
		_67_v2 = _67_v2
	} else {
		if _68_v3 {
			var _96_v28 _dafny.Array
			_ = _96_v28
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw15
			_96_v28 = _nw15
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_96_v28), 0))
			_ = _index10
			(_96_v28).ArraySet1(_67_v2, (_index10).Int())
			var _97_v29 _dafny.MultiSet
			_ = _97_v29
			_97_v29 = _dafny.MultiSetOf(_73_v7, _dafny.CodePoint('k'), _73_v7)
			var _98_v30 _dafny.Map
			_ = _98_v30
			_98_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v3, _dafny.IntOfInt64(-34))
			var _99_v31 _dafny.MultiSet
			_ = _99_v31
			_99_v31 = _dafny.MultiSetOf(_68_v3, _68_v3, _68_v3)
			var _100_v32 _dafny.MultiSet
			_ = _100_v32
			_100_v32 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_97_v29).Contains(_73_v7) {
					return (_97_v29).Multiplicity(_73_v7)
				}
				return Companion_Default___.Fm4(_67_v2, _67_v2, false, _68_v3, _64_globalState)
			})(), (_dafny.MultiSetOf(_68_v3)).Cardinality(), _67_v2, (((_98_v30).Update(Companion_Default___.Fm1(_68_v3, _68_v3, _67_v2, _64_globalState), _67_v2)).Cardinality()).Plus((_99_v31).Cardinality()), _67_v2)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_96_v28), 0))
			_ = _index11
			(_96_v28).ArraySet1((func() _dafny.Int {
				if (_100_v32).Contains(_67_v2) {
					return (_100_v32).Multiplicity(_67_v2)
				}
				return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(850), (_dafny.Zero).Minus(_67_v2))
			})(), (_index11).Int())
			var _101_v33 _dafny.Sequence
			_ = _101_v33
			_101_v33 = _dafny.SeqOf(false)
			var _102_v34 *C0
			_ = _102_v34
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__(_dafny.SeqOf((_dafny.MultiSetFromSeq(_101_v33)).Cardinality()))
			_102_v34 = _nw16
			_81_v14 = _81_v14
			_73_v7 = (_81_v14).Select((Companion_Default___.SafeIndex(_67_v2, _dafny.IntOfUint32((_81_v14).Cardinality()))).Uint32()).(_dafny.CodePoint)
			var _103_v35 _dafny.Array
			_ = _103_v35
			var _nw17 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
			_ = _nw17
			_103_v35 = _nw17
			var _104_v36 _dafny.Map
			_ = _104_v36
			_104_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v35, _68_v3)
			var _105_v37 _dafny.Sequence
			_ = _105_v37
			var _106_v38 _dafny.Int
			_ = _106_v38
			var _107_v39 bool
			_ = _107_v39
			var _out6 _dafny.Sequence
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			var _out8 bool
			_ = _out8
			_out6, _out7, _out8 = Companion_Default___.M0(_68_v3, !(((_dafny.Zero).Minus((_104_v36).Cardinality())).Cmp((_102_v34).Fm5(_68_v3, _68_v3, _68_v3, (_96_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_96_v28), 0))).Int()).(_dafny.Int), _64_globalState)) < 0), _64_globalState)
			_105_v37 = _out6
			_106_v38 = _out7
			_107_v39 = _out8
		} else {
			var _108_v40 _dafny.Set
			_ = _108_v40
			_108_v40 = _dafny.SetOf(!(_68_v3))
			var _109_v41 _dafny.Map
			_ = _109_v41
			_109_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v6, _108_v40)
			_68_v3 = (func() bool {
				if _68_v3 {
					return _68_v3
				}
				return (_109_v41).Equals(_109_v41)
			})()
			_67_v2 = (_67_v2).Plus((_dafny.Zero).Minus(_67_v2))
			_67_v2 = _67_v2
			var _110_v42 _dafny.Map
			_ = _110_v42
			_110_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v3, _68_v3)
			_110_v42 = (_110_v42).Merge(_110_v42)
			var _111_v43 _dafny.Sequence
			_ = _111_v43
			_111_v43 = _dafny.SeqOf(true, _68_v3, false, _68_v3, _68_v3)
			var _112_v44 _dafny.MultiSet
			_ = _112_v44
			_112_v44 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_())
			var _113_v45 _dafny.Map
			_ = _113_v45
			_113_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_111_v43).Select((Companion_Default___.SafeIndex(_67_v2, _dafny.IntOfUint32((_111_v43).Cardinality()))).Uint32()).(bool), (func() _dafny.Int {
				if (_112_v44).Contains(Companion_D0_.Create_DC1_()) {
					return (_112_v44).Multiplicity(Companion_D0_.Create_DC1_())
				}
				return _67_v2
			})())
			_67_v2 = (Companion_Default___.SafeModuloInt((_113_v45).Cardinality(), _67_v2)).Plus(_67_v2)
		}
		var _114_v46 _dafny.Map
		_ = _114_v46
		_114_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v3, _68_v3)
		_114_v46 = (_114_v46).Update(_68_v3, (_70_v6).Dtor_cf0())
		var _115_v47 _dafny.MultiSet
		_ = _115_v47
		_115_v47 = _dafny.MultiSetOf(_68_v3, _68_v3, _68_v3)
		var _116_v48 *C0
		_ = _116_v48
		var _nw18 *C0 = New_C0_()
		_ = _nw18
		_nw18.Ctor__(_dafny.SeqOf(_dafny.IntOfUint32((_81_v14).Cardinality()), _dafny.IntOfInt64(706), (_115_v47).Cardinality()))
		_116_v48 = _nw18
		var _117_v49 _dafny.Map
		_ = _117_v49
		_117_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _116_v48)
		var _118_v50 D1
		_ = _118_v50
		_118_v50 = Companion_D1_.Create_DC2_((func() *C0 {
			if (_117_v49).Contains(_68_v3) {
				return (_117_v49).Get(_68_v3).(*C0)
			}
			return _116_v48
		})())
		var _119_v51 _dafny.Array
		_ = _119_v51
		var _nwElement0_5 *C0 = _116_v48
		_ = _nwElement0_5
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(16))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_5, 0)
		(_nw19).ArraySet1(_116_v48, 1)
		(_nw19).ArraySet1(_116_v48, 2)
		(_nw19).ArraySet1(_116_v48, 3)
		(_nw19).ArraySet1((func() *C0 {
			if true {
				return _116_v48
			}
			return _116_v48
		})(), 4)
		(_nw19).ArraySet1(_116_v48, 5)
		(_nw19).ArraySet1(_116_v48, 6)
		(_nw19).ArraySet1(_116_v48, 7)
		(_nw19).ArraySet1(_116_v48, 8)
		(_nw19).ArraySet1(_116_v48, 9)
		(_nw19).ArraySet1(_116_v48, 10)
		(_nw19).ArraySet1(_116_v48, 11)
		(_nw19).ArraySet1((_118_v50).Dtor_cf1(), 12)
		(_nw19).ArraySet1(_116_v48, 13)
		(_nw19).ArraySet1(_116_v48, 14)
		(_nw19).ArraySet1(_116_v48, 15)
		_119_v51 = _nw19
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_119_v51), 0))
		_ = _index12
		(_119_v51).ArraySet1(_116_v48, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_119_v51), 0))
		_ = _index13
		(_119_v51).ArraySet1(_116_v48, (_index13).Int())
		_68_v3 = _68_v3
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_65_v0), 0))
		_ = _index14
		(_65_v0).ArraySet1((_68_v3) == (_68_v3), (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_65_v0), 0))
		_ = _index15
		(_65_v0).ArraySet1(_68_v3, (_index15).Int())
	}
	if (_70_v6).Dtor_cf0() {
		var _120_v52 _dafny.Sequence
		_ = _120_v52
		_120_v52 = _dafny.SeqOf(_68_v3)
		var _121_v53 _dafny.MultiSet
		_ = _121_v53
		_121_v53 = _dafny.MultiSetOf(_68_v3, _68_v3)
		_67_v2 = (func() _dafny.Int {
			if (_121_v53).IsProperSubsetOf(_dafny.MultiSetFromSeq(_120_v52)) {
				return _67_v2
			}
			return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_120_v52).Cardinality()))).Times(_67_v2)
		})()
		var _122_v54 _dafny.Set
		_ = _122_v54
		_122_v54 = _dafny.SetOf(true)
		_122_v54 = (func() _dafny.Set {
			if _68_v3 {
				return (_dafny.SetOf(false, _68_v3)).Difference(_dafny.SetOf(!(_68_v3)))
			}
			return _dafny.SetOf(_68_v3, _68_v3, _68_v3)
		})()
		var _123_v55 _dafny.Array
		_ = _123_v55
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_2
		var _nw20 _dafny.Array
		_ = _nw20
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw20 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_124_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_125_i3 _dafny.Int) _dafny.Int {
					return (_125_i3).Minus(_124_v2)
				}
			})(_67_v2)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw20 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw20).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw20).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_123_v55 = _nw20
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_123_v55), 0))
		_ = _index16
		(_123_v55).ArraySet1(_67_v2, (_index16).Int())
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_123_v55), 0))
		_ = _index17
		(_123_v55).ArraySet1(_dafny.IntOfInt64(-266), (_index17).Int())
		var _126_v56 _dafny.Sequence
		_ = _126_v56
		_126_v56 = _dafny.SeqOf(_67_v2)
		var _127_v57 _dafny.Sequence
		_ = _127_v57
		_127_v57 = _dafny.SeqOf(_dafny.IntOfUint32((_126_v56).Cardinality()), _dafny.IntOfUint32((_126_v56).Cardinality()))
		var _128_v58 *C0
		_ = _128_v58
		var _nw21 *C0 = New_C0_()
		_ = _nw21
		_nw21.Ctor__(_127_v57)
		_128_v58 = _nw21
		var _129_v59 _dafny.Map
		_ = _129_v59
		_129_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v58, true)
		var _130_v60 _dafny.Map
		_ = _130_v60
		_130_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_81_v14).Cardinality()), _129_v59)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_123_v55), 0))
		_ = _index18
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_123_v55), 0))
		_ = _index19
		var _rhs21 _dafny.Int = (_dafny.Zero).Minus(_67_v2)
		_ = _rhs21
		var _rhs22 bool = (_70_v6).Dtor_cf0()
		_ = _rhs22
		var _rhs23 _dafny.Int = _67_v2
		_ = _rhs23
		var _rhs24 _dafny.Int = _dafny.IntOfInt64(246)
		_ = _rhs24
		var _rhs25 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_81_v14).Cardinality()), (Companion_D1_.Create_DC4_(_130_v60, _67_v2)).Dtor_cf5()))).Minus((_128_v58).Fm5(false, _68_v3, _68_v3, Companion_Default___.Fm4(_67_v2, _67_v2, _68_v3, _68_v3, _64_globalState), _64_globalState))
		_ = _rhs25
		var _lhs4 _dafny.Array = _123_v55
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_123_v55), 0))
		_ = _lhs5
		var _lhs6 _dafny.Array = _123_v55
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_123_v55), 0))
		_ = _lhs7
		_67_v2 = _rhs21
		_68_v3 = _rhs22
		(_lhs4).ArraySet1(_rhs23, (_lhs5).Int())
		_67_v2 = _rhs24
		(_lhs6).ArraySet1(_rhs25, (_lhs7).Int())
		_68_v3 = (_121_v53).IsSubsetOf(_dafny.MultiSetFromSeq(_120_v52))
		_68_v3 = true
	} else {
		_69_v4 = (_69_v4).Update(_67_v2, (_dafny.Zero).Minus(_67_v2))
		_67_v2 = _dafny.IntOfInt64(935)
		_67_v2 = (_dafny.Zero).Minus(_67_v2)
		_68_v3 = !(_68_v3)
		var _131_v61 _dafny.Map
		_ = _131_v61
		_131_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_68_v3), _67_v2)
		var _132_v62 _dafny.MultiSet
		_ = _132_v62
		_132_v62 = _dafny.MultiSetOf(_68_v3)
		_68_v3 = ((func() _dafny.Int {
			if (_131_v61).Contains(_132_v62) {
				return (_131_v61).Get(_132_v62).(_dafny.Int)
			}
			return _67_v2
		})()).Cmp(_67_v2) <= 0
	}
	var _133_i4 _dafny.Int
	_ = _133_i4
	_133_i4 = _dafny.Zero
	{
		for _68_v3 {
			{
				if (_133_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_133_i4 = (_133_i4).Plus(_dafny.One)
				_67_v2 = _67_v2
				var _134_v63 _dafny.Map
				_ = _134_v63
				_134_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_67_v2)), _68_v3)
				var _135_v64 *C0
				_ = _135_v64
				var _nw22 *C0 = New_C0_()
				_ = _nw22
				_nw22.Ctor__((Companion_Default___.Fm6(_67_v2, _68_v3, (func() bool {
					if (_134_v63).Contains(_dafny.IntOfInt64(834)) {
						return (_134_v63).Get(_dafny.IntOfInt64(834)).(bool)
					}
					return _68_v3
				})(), _64_globalState)))
				_135_v64 = _nw22
				_81_v14 = _81_v14
				_67_v2 = (func() _dafny.Int {
					if _68_v3 {
						return _67_v2
					}
					return _67_v2
				})()
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_68_v3 = _68_v3
	_68_v3 = (Companion_Default___.Fm4(_67_v2, (_dafny.Zero).Minus(_67_v2), false, true, _64_globalState)).Cmp(_67_v2) <= 0
	var _136_v65 _dafny.Set
	_ = _136_v65
	_136_v65 = _dafny.SetOf(_dafny.IntOfUint32((_81_v14).Cardinality()))
	var _137_v66 _dafny.Sequence
	_ = _137_v66
	_137_v66 = _dafny.SeqOf(_67_v2)
	var _138_v67 *C0
	_ = _138_v67
	var _nw23 *C0 = New_C0_()
	_ = _nw23
	_nw23.Ctor__(_137_v66)
	_138_v67 = _nw23
	var _139_v68 _dafny.Map
	_ = _139_v68
	_139_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v67, _68_v3)
	var _140_v69 _dafny.Map
	_ = _140_v69
	_140_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(232), _139_v68)
	var _141_v70 D1
	_ = _141_v70
	_141_v70 = Companion_D1_.Create_DC4_(_140_v69, (_dafny.MultiSetOf(!(_68_v3), _68_v3)).Cardinality())
	var _142_v71 _dafny.Map
	_ = _142_v71
	_142_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_136_v65, _141_v70)
	var _143_v72 _dafny.Sequence
	_ = _143_v72
	_143_v72 = _dafny.SeqOf((_142_v71).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-911)), _dafny.IntOfUint32((_81_v14).Cardinality()))
	var _144_v73 *C0
	_ = _144_v73
	var _nw24 *C0 = New_C0_()
	_ = _nw24
	_nw24.Ctor__(_137_v66)
	_144_v73 = _nw24
	var _145_v74 _dafny.Sequence
	_ = _145_v74
	var _146_v75 _dafny.Int
	_ = _146_v75
	var _147_v76 bool
	_ = _147_v76
	var _out9 _dafny.Sequence
	_ = _out9
	var _out10 _dafny.Int
	_ = _out10
	var _out11 bool
	_ = _out11
	_out9, _out10, _out11 = Companion_Default___.M0(_68_v3, (_70_v6).Dtor_cf0(), _64_globalState)
	_145_v74 = _out9
	_146_v75 = _out10
	_147_v76 = _out11
	_146_v75 = (func() _dafny.Int {
		if _68_v3 {
			return _dafny.IntOfInt64(-550)
		}
		return Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_67_v2), (_dafny.Zero).Minus(_146_v75))
	})()
	_dafny.Print((_65_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_65_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_65_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_65_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_65_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_65_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_65_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v1).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_67_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_68_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(59), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(60), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(61), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(62), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(63), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(64), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(65), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(66), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(68), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(69), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(70), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(71), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(72), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(73), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(74), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(75), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(76), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(77), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(78), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(79), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(80), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(81), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(82), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(83), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(84), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(85), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(86), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(87), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(88), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(89), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(90), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(91), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(92), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(93), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(94), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(95), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(96), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(97), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(3), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(4), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(5), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(6), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(7), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(8), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(10), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(11), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(12), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(13), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(14), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(15), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(16), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(17), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(18), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(19), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(20), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(21), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(22), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(23), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(24), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(25), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(26), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(27), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(28), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(29), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(30), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(31), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(32), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(33), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(34), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(35), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(36), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(37), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(38), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(39), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(40), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(41), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(42), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(43), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(44), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(45), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(46), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(47), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(48), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(49), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(50), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(51), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(52), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(53), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(54), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(55), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(56), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(57), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(58), _dafny.IntOfInt64(98)).UpdateUnsafe(_dafny.IntOfInt64(199), _dafny.IntOfInt64(-199))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_70_v6).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_73_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_81_v14.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_84_v17).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v18).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("braqfmabraqfmadvtinys"), _dafny.UnicodeSeqOfUtf8Bytes("braqfmabraqfmadvtinys"), _dafny.UnicodeSeqOfUtf8Bytes("braqfmabraqfmadvtinys"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_86_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v65).Equals(_dafny.SetOf(_dafny.IntOfInt64(7))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_137_v66, _dafny.SeqOf(_dafny.IntOfInt64(-935))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_138_v67).F0(), _dafny.SeqOf(_dafny.IntOfInt64(-935))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v68).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_140_v69).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_141_v70).Dtor_cf4()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v70).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_142_v71).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_143_v72, _dafny.SeqOf(_dafny.One, _dafny.IntOfInt64(911), _dafny.IntOfInt64(7))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_144_v73).F0(), _dafny.SeqOf(_dafny.IntOfInt64(-935))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_145_v74, _dafny.SeqOf(_dafny.IntOfInt64(139), _dafny.IntOfInt64(14), _dafny.IntOfInt64(14), _dafny.IntOfInt64(14))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_v75)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_v76)
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
	Cf2 _dafny.CodePoint
	Cf3 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf2 _dafny.CodePoint, Cf3 _dafny.Int) D1 {
	return D1{D1_DC3{Cf2, Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf4 _dafny.Map
	Cf5 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.Map, Cf5 _dafny.Int) D1 {
	return D1{D1_DC4{Cf4, Cf5}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf1 *C0
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf1 *C0) D1 {
	return D1{D1_DC2{Cf1}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D1) Dtor_cf2() _dafny.CodePoint {
	return _this.Get_().(D1_DC3).Cf2
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf1() *C0 {
	return _this.Get_().(D1_DC2).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf4.Equals(data2.Cf4) && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf1 == data2.Cf1
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
	Cf6 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf6 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf6}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D2) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf8 bool
	Cf9 bool
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf8 bool, Cf9 bool) D3 {
	return D3{D3_DC7{Cf8, Cf9}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC6 struct {
	Cf7 _dafny.Map
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf7 _dafny.Map) D3 {
	return D3{D3_DC6{Cf7}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_(false, false)
}

func (_this D3) Dtor_cf8() bool {
	return _this.Get_().(D3_DC7).Cf8
}

func (_this D3) Dtor_cf9() bool {
	return _this.Get_().(D3_DC7).Cf9
}

func (_this D3) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D3_DC6).Cf7
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf10 _dafny.Array
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf10 _dafny.Array) D4 {
	return D4{D4_DC8{Cf10}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D4) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D4_DC8).Cf10
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf10 == data2.Cf10
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
	dummy byte
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

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

func (_this *GlobalState) Ctor__() {
	{
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f0 _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f0 = _dafny.EmptySeq
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

func (_this *C0) Ctor__(f0 _dafny.Sequence) {
	{
		(_this)._f0 = f0
	}
}
func (_this *C0) Fm5(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-344)
	}
}
func (_this *C0) F0() _dafny.Sequence {
	{
		return _this._f0
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
