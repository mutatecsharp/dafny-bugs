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
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('s')
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("qcjolb"), _dafny.UnicodeSeqOfUtf8Bytes("jvvpxlotu"))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.IntOfInt64(42)).Minus(_dafny.IntOfInt64(-345))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(32))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uus"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(725))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	}))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(692), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(604))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality())
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(206)
	}))))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.SeqOf(!(!(!(!(false)))), !(false))
	} else {
		return _dafny.SeqOf(true)
	}
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) (_dafny.Sequence, _dafny.Array, _dafny.Int, bool) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var r3 bool = false
	_ = r3
	var _4_v0 _dafny.Int
	_ = _4_v0
	_4_v0 = _dafny.IntOfInt64(341)
	var _5_v1 _dafny.Sequence
	_ = _5_v1
	_5_v1 = _dafny.SeqOf(_4_v0)
	var _6_v2 _dafny.CodePoint
	_ = _6_v2
	_6_v2 = _dafny.CodePoint('n')
	var _7_v3 bool
	_ = _7_v3
	_7_v3 = false
	var _8_v4 _dafny.Array
	_ = _8_v4
	var _nwElement0_0 _dafny.Sequence = _5_v1
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(12))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(_5_v1, 1)
	(_nw0).ArraySet1(_5_v1, 2)
	(_nw0).ArraySet1(_5_v1, 3)
	(_nw0).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wuojgho")).Cardinality()), _4_v0), 4)
	(_nw0).ArraySet1(_5_v1, 5)
	(_nw0).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(439))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}((func(_9_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_10_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_9_v0)
		}
	})(_4_v0))), 6)
	(_nw0).ArraySet1(_5_v1, 7)
	(_nw0).ArraySet1(_5_v1, 8)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_5_v1, _5_v1), 9)
	(_nw0).ArraySet1(Companion_Default___.Fm8(_6_v2, _7_v3, _5_v1, globalState), 10)
	(_nw0).ArraySet1(_5_v1, 11)
	_8_v4 = _nw0
	var _11_v5 D3
	_ = _11_v5
	_11_v5 = Companion_D3_.Create_DC11_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_4_v0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(153))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_12_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("axdp")).Cardinality())
	}))))
	var _13_v6 _dafny.Set
	_ = _13_v6
	_13_v6 = _dafny.SetOf(_7_v3, _7_v3, _7_v3, _7_v3, _7_v3)
	var _14_v7 _dafny.Sequence
	_ = _14_v7
	_14_v7 = _dafny.UnicodeSeqOfUtf8Bytes("ekcwr")
	var _rhs0 bool = Companion_Default___.Fm5(_7_v3, ((_13_v6).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(648))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}((func(_15_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_16_i2 _dafny.Int) _dafny.CodePoint {
			return _15_v2
		}
	})(_6_v2)))).Cardinality())) < 0, _4_v0, _4_v0, globalState)
	_ = _rhs0
	var _rhs1 bool = ((func() _dafny.Int {
		if _7_v3 {
			return _4_v0
		}
		return _4_v0
	})()).Cmp((_dafny.Zero).Minus(Companion_Default___.Fm6(_7_v3, _7_v3, false, globalState))) == 0
	_ = _rhs1
	var _rhs2 _dafny.Array = _8_v4
	_ = _rhs2
	var _rhs3 _dafny.Int = (Companion_Default___.Fm6(true, false, _7_v3, globalState)).Plus(_dafny.IntOfUint32((_14_v7).Cardinality()))
	_ = _rhs3
	var _rhs4 D3 = _11_v5
	_ = _rhs4
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 *GlobalState = globalState
	_ = _lhs1
	_lhs0.F20 = _rhs0
	_lhs1.F20 = _rhs1
	_8_v4 = _rhs2
	r2 = _rhs3
	_11_v5 = _rhs4
	var _17_v8 _dafny.Array
	_ = _17_v8
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
	_ = _nw1
	_17_v8 = _nw1
	for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_17_v8), 0))); ; {
		_guard_loop_0, _ok0 := _iter0()
		if !_ok0 {
			break
		}
		var _18_i3 _dafny.Int
		_18_i3 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_18_i3).Sign() != -1) && ((_18_i3).Cmp(_dafny.ArrayLen((_17_v8), 0)) < 0)) {
			(_17_v8).ArraySet1(Companion_Default___.SafeDivisionInt(_18_i3, _4_v0), (_18_i3).Int())
		}
	}
	var _19_v9 _dafny.Array
	_ = _19_v9
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
	_ = _nw2
	_19_v9 = _nw2
	for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_19_v9), 0))); ; {
		_guard_loop_1, _ok1 := _iter1()
		if !_ok1 {
			break
		}
		var _20_i4 _dafny.Int
		_20_i4 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_20_i4).Sign() != -1) && ((_20_i4).Cmp(_dafny.ArrayLen((_19_v9), 0)) < 0)) {
			(_19_v9).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_4_v0, Companion_D4_.Create_DC18_()), (_20_i4).Int())
		}
	}
	(globalState).F20 = !(_7_v3)
	if (_4_v0).Cmp(Companion_Default___.SafeDivisionInt(_4_v0, _4_v0)) < 0 {
		var _21_v10 D0
		_ = _21_v10
		_21_v10 = Companion_D0_.Create_DC1_(!(_7_v3), _7_v3, _4_v0)
		var _22_v11 _dafny.Sequence
		_ = _22_v11
		_22_v11 = _dafny.SeqOf(_7_v3)
		var _23_v12 _dafny.Map
		_ = _23_v12
		_23_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_4_v0, _4_v0)
		var _24_v13 D5
		_ = _24_v13
		_24_v13 = Companion_D5_.Create_DC21_(_23_v12)
		var _25_v14 D5
		_ = _25_v14
		_25_v14 = Companion_D5_.Create_DC22_(_24_v13)
		var _26_v15 _dafny.Map
		_ = _26_v15
		_26_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-355), false)
		var _27_v16 _dafny.Array
		_ = _27_v16
		var _nwElement0_1 bool = !(true) || (_7_v3)
		_ = _nwElement0_1
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(28))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_1, 0)
		(_nw3).ArraySet1(_7_v3, 1)
		(_nw3).ArraySet1(_7_v3, 2)
		(_nw3).ArraySet1(_7_v3, 3)
		(_nw3).ArraySet1((func() bool {
			if _7_v3 {
				return _7_v3
			}
			return _7_v3
		})(), 4)
		(_nw3).ArraySet1(_7_v3, 5)
		(_nw3).ArraySet1(_7_v3, 6)
		(_nw3).ArraySet1(false, 7)
		(_nw3).ArraySet1(_7_v3, 8)
		(_nw3).ArraySet1(_7_v3, 9)
		(_nw3).ArraySet1((func() bool {
			if (_21_v10).Dtor_cf1() {
				return _7_v3
			}
			return _7_v3
		})(), 10)
		(_nw3).ArraySet1(_7_v3, 11)
		(_nw3).ArraySet1((_7_v3) == (_7_v3), 12)
		(_nw3).ArraySet1(true, 13)
		(_nw3).ArraySet1(!(!(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_22_v11).Cardinality()), _25_v14)).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(302))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_28_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_29_i5 _dafny.Int) _dafny.CodePoint {
				return _28_v2
			}
		})(_6_v2)))).Cardinality())) != 0)), 14)
		(_nw3).ArraySet1(!(_7_v3) || (_7_v3), 15)
		(_nw3).ArraySet1(_7_v3, 16)
		(_nw3).ArraySet1(_7_v3, 17)
		(_nw3).ArraySet1(_7_v3, 18)
		(_nw3).ArraySet1(_7_v3, 19)
		(_nw3).ArraySet1(_7_v3, 20)
		(_nw3).ArraySet1(_7_v3, 21)
		(_nw3).ArraySet1(true, 22)
		(_nw3).ArraySet1(!(_7_v3), 23)
		(_nw3).ArraySet1(_7_v3, 24)
		(_nw3).ArraySet1(!_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm7(_7_v3, _dafny.IntOfInt64(110), _4_v0, false, globalState), _6_v2), 25)
		(_nw3).ArraySet1(((func() bool {
			if (_26_v15).Contains(_4_v0) {
				return (_26_v15).Get(_4_v0).(bool)
			}
			return true
		})()) && (_7_v3), 26)
		(_nw3).ArraySet1(false, 27)
		_27_v16 = _nw3
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))
		_ = _index0
		(_27_v16).ArraySet1(true, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))
		_ = _index1
		var _rhs5 _dafny.Int = (_4_v0).Minus((_4_v0).Minus(_4_v0))
		_ = _rhs5
		var _rhs6 bool = _7_v3
		_ = _rhs6
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		var _lhs3 _dafny.Array = _27_v16
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))
		_ = _lhs4
		_lhs2.F5 = _rhs5
		(_lhs3).ArraySet1(_rhs6, (_lhs4).Int())
		var _30_v17 _dafny.Array
		_ = _30_v17
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw4
		_30_v17 = _nw4
		_30_v17 = (func() _dafny.Array {
			if (_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool) {
				return _30_v17
			}
			return _30_v17
		})()
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))
		_ = _index2
		(_27_v16).ArraySet1((_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool), (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))
		_ = _index3
		(_27_v16).ArraySet1(((_dafny.Zero).Minus(_4_v0)).Cmp(_4_v0) != 0, (_index3).Int())
		if (_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool) {
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))
			_ = _index4
			(_27_v16).ArraySet1(((_4_v0).Minus(_4_v0)).Cmp(_4_v0) < 0, (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))
			_ = _index5
			(_27_v16).ArraySet1(_7_v3, (_index5).Int())
			var _31_v18 _dafny.Set
			_ = _31_v18
			_31_v18 = _dafny.SetOf(_4_v0, _dafny.IntOfUint32((_5_v1).Cardinality()))
			_7_v3 = (_31_v18).IsProperSubsetOf(_31_v18)
			var _32_v19 *C0
			_ = _32_v19
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__(_4_v0, !(Companion_Default___.Fm5((_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool), _7_v3, _4_v0, _dafny.IntOfUint32((_22_v11).Cardinality()), globalState)), _dafny.Companion_Sequence_.Concatenate(_14_v7, _14_v7))
			_32_v19 = _nw5
			var _33_v20 _dafny.Map
			_ = _33_v20
			_33_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _32_v19)
			var _34_v21 D1
			_ = _34_v21
			_34_v21 = Companion_D1_.Create_DC4_((func() *C0 {
				if (_33_v20).Contains(_7_v3) {
					return (_33_v20).Get(_7_v3).(*C0)
				}
				return _32_v19
			})())
			_34_v21 = _34_v21
		} else {
			_22_v11 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool), Companion_Default___.Fm5(_7_v3, _7_v3, _4_v0, _4_v0, globalState)), Companion_Default___.Fm9(_dafny.Companion_Sequence_.Update(_22_v11, (Companion_Default___.SafeIndex(_4_v0, _dafny.IntOfUint32((_22_v11).Cardinality()))).Uint32(), (_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_35_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_36_i6 _dafny.Int) _dafny.CodePoint {
					return _35_v2
				}
			})(_6_v2)))).Cardinality()), (_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool), globalState))
			var _37_v22 *C0
			_ = _37_v22
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__((_13_v6).Cardinality(), false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-589))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_38_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_39_i7 _dafny.Int) _dafny.CodePoint {
					return _38_v2
				}
			})(_6_v2))))
			_37_v22 = _nw6
			var _40_v23 _dafny.Sequence
			_ = _40_v23
			_40_v23 = _dafny.SeqOf(_37_v22)
			var _41_v24 _dafny.Array
			_ = _41_v24
			var _nwElement0_2 *C0 = _37_v22
			_ = _nwElement0_2
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(18))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_2, 0)
			(_nw7).ArraySet1(_37_v22, 1)
			(_nw7).ArraySet1(_37_v22, 2)
			(_nw7).ArraySet1(_37_v22, 3)
			(_nw7).ArraySet1(_37_v22, 4)
			(_nw7).ArraySet1(_37_v22, 5)
			(_nw7).ArraySet1((_40_v23).Select((Companion_Default___.SafeIndex((_37_v22).F26(), _dafny.IntOfUint32((_40_v23).Cardinality()))).Uint32()).(*C0), 6)
			(_nw7).ArraySet1((func() *C0 {
				if (_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool) {
					return _37_v22
				}
				return _37_v22
			})(), 7)
			(_nw7).ArraySet1(_37_v22, 8)
			(_nw7).ArraySet1((func() *C0 {
				if _7_v3 {
					return _37_v22
				}
				return _37_v22
			})(), 9)
			(_nw7).ArraySet1(_37_v22, 10)
			(_nw7).ArraySet1(_37_v22, 11)
			(_nw7).ArraySet1(_37_v22, 12)
			(_nw7).ArraySet1(_37_v22, 13)
			(_nw7).ArraySet1(_37_v22, 14)
			(_nw7).ArraySet1(_37_v22, 15)
			(_nw7).ArraySet1(_37_v22, 16)
			(_nw7).ArraySet1(_37_v22, 17)
			_41_v24 = _nw7
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_41_v24), 0))
			_ = _index6
			(_41_v24).ArraySet1((func() *C0 {
				if (_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool) {
					return _37_v22
				}
				return _37_v22
			})(), (_index6).Int())
			var _42_v25 _dafny.Map
			_ = _42_v25
			_42_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v17, _37_v22)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_41_v24), 0))
			_ = _index7
			(_41_v24).ArraySet1((func() *C0 {
				if ((_37_v22).F26()).Cmp(_4_v0) < 0 {
					return (func() *C0 {
						if _7_v3 {
							return _37_v22
						}
						return _37_v22
					})()
				}
				return (func() *C0 {
					if (_42_v25).Contains(_30_v17) {
						return (_42_v25).Get(_30_v17).(*C0)
					}
					return _37_v22
				})()
			})(), (_index7).Int())
			var _43_v26 _dafny.MultiSet
			_ = _43_v26
			_43_v26 = _dafny.MultiSetOf((_41_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_41_v24), 0))).Int()).(*C0), _37_v22)
			var _44_v27 *C0
			_ = _44_v27
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__((_dafny.IntOfInt64(365)).Plus((_43_v26).Cardinality()), _7_v3, _14_v7)
			_44_v27 = _nw8
			var _45_v28 _dafny.MultiSet
			_ = _45_v28
			_45_v28 = _dafny.MultiSetOf(_17_v8)
			var _46_v29 T0
			_ = _46_v29
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__((_45_v28).Cardinality(), (_27_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_27_v16), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(126))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_47_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_48_i8 _dafny.Int) _dafny.CodePoint {
					return _47_v2
				}
			})(_6_v2))))
			_46_v29 = _nw9
			var _49_v30 _dafny.Map
			_ = _49_v30
			_49_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v29, (_44_v27).F26())
			_49_v30 = (_49_v30).Update(_46_v29, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-65))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_50_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_51_i9 _dafny.Int) _dafny.CodePoint {
					return _50_v2
				}
			})(_6_v2)))).Cardinality()))
			var _52_v31 *C0
			_ = _52_v31
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__((_44_v27).F26(), _7_v3, _14_v7)
			_52_v31 = _nw10
		}
	} else {
		(globalState).F5 = _4_v0
		_5_v1 = _5_v1
		var _53_v32 _dafny.Array
		_ = _53_v32
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
		_ = _nw11
		_53_v32 = _nw11
		var _54_v33 _dafny.Array
		_ = _54_v33
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_0
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = func(_55_i10 _dafny.Int) bool {
				return true
			}
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw12 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw12).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw12).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_54_v33 = _nw12
		var _56_v34 D2
		_ = _56_v34
		_56_v34 = Companion_D2_.Create_DC8_(_54_v33)
		var _57_v35 D2
		_ = _57_v35
		_57_v35 = Companion_D2_.Create_DC10_(_56_v34)
		var _58_v36 _dafny.Map
		_ = _58_v36
		_58_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v35, _4_v0)
		var _59_v37 *C0
		_ = _59_v37
		var _nw13 *C0 = New_C0_()
		_ = _nw13
		_nw13.Ctor__((_dafny.Zero).Minus((_58_v36).Cardinality()), Companion_Default___.Fm5(_7_v3, _7_v3, _4_v0, _4_v0, globalState), _14_v7)
		_59_v37 = _nw13
		var _60_v38 _dafny.Array
		_ = _60_v38
		_60_v38 = _53_v32
		var _rhs7 _dafny.Array = (_60_v38)
		_ = _rhs7
		var _rhs8 *C0 = _59_v37
		_ = _rhs8
		var _rhs9 _dafny.Int = ((_59_v37).F26()).Times(_dafny.IntOfUint32((_5_v1).Cardinality()))
		_ = _rhs9
		var _lhs5 *GlobalState = globalState
		_ = _lhs5
		_53_v32 = _rhs7
		_59_v37 = _rhs8
		_lhs5.F5 = _rhs9
		var _61_v39 _dafny.Map
		_ = _61_v39
		_61_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v3, _7_v3)
		var _62_v40 _dafny.Sequence
		_ = _62_v40
		_62_v40 = _dafny.SeqOf(_7_v3)
		_61_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v3, (_62_v40).Select((Companion_Default___.SafeIndex((_59_v37).F26(), _dafny.IntOfUint32((_62_v40).Cardinality()))).Uint32()).(bool))
		var _63_v41 _dafny.Map
		_ = _63_v41
		_63_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _6_v2)
		var _64_v42 D4
		_ = _64_v42
		_64_v42 = Companion_D4_.Create_DC17_((_63_v41).Cardinality(), _59_v37)
		var _65_v43 _dafny.Map
		_ = _65_v43
		_65_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_64_v42).Dtor_cf29(), _7_v3)
		_65_v43 = (_65_v43).Update(_4_v0, _7_v3)
	}
	var _66_v44 T0
	_ = _66_v44
	var _nw14 *C0 = New_C0_()
	_ = _nw14
	_nw14.Ctor__(_4_v0, Companion_Default___.Fm5(_7_v3, _7_v3, _dafny.IntOfInt64(-857), Companion_Default___.Fm6(_7_v3, _7_v3, _7_v3, globalState), globalState), _14_v7)
	_66_v44 = _nw14
	var _67_v45 _dafny.Map
	_ = _67_v45
	_67_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_66_v44, _dafny.IntOfInt64(-276))
	var _68_v46 _dafny.Map
	_ = _68_v46
	_68_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v45, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}((func(_69_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_70_i11 _dafny.Int) _dafny.CodePoint {
			return _69_v2
		}
	})(_6_v2))))
	(globalState).F4 = Companion_Default___.Fm7(_7_v3, (_dafny.Zero).Minus((_68_v46).Cardinality()), _4_v0, _7_v3, globalState)
	r0 = _dafny.Companion_Sequence_.Concatenate(_14_v7, _dafny.Companion_Sequence_.Concatenate(_14_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(217))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}((func(_71_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_72_i12 _dafny.Int) _dafny.CodePoint {
			return _71_v2
		}
	})(_6_v2)))))
	var _73_v47 _dafny.Array
	_ = _73_v47
	var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
	_ = _nw15
	_73_v47 = _nw15
	r1 = _73_v47
	r2 = _4_v0
	var _74_v48 *C0
	_ = _74_v48
	var _nw16 *C0 = New_C0_()
	_ = _nw16
	_nw16.Ctor__(_4_v0, _7_v3, _14_v7)
	_74_v48 = _nw16
	var _75_v49 _dafny.MultiSet
	_ = _75_v49
	_75_v49 = _dafny.MultiSetOf(_74_v48, _74_v48, _74_v48)
	var _76_v50 _dafny.Sequence
	_ = _76_v50
	_76_v50 = _dafny.SeqOf(_74_v48, _74_v48)
	r3 = ((func() _dafny.Int {
		if (_75_v49).Contains((_76_v50).Select((Companion_Default___.SafeIndex((_74_v48).F26(), _dafny.IntOfUint32((_76_v50).Cardinality()))).Uint32()).(*C0)) {
			return (_75_v49).Multiplicity((_76_v50).Select((Companion_Default___.SafeIndex((_74_v48).F26(), _dafny.IntOfUint32((_76_v50).Cardinality()))).Uint32()).(*C0))
		}
		return _4_v0
	})()).Cmp((_74_v48).F26()) > 0
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _77_v0 _dafny.Sequence
	_ = _77_v0
	_77_v0 = _dafny.UnicodeSeqOfUtf8Bytes("mkyq")
	var _78_v1 bool
	_ = _78_v1
	_78_v1 = false
	var _79_v2 _dafny.Set
	_ = _79_v2
	_79_v2 = _dafny.SetOf(_78_v1, _78_v1)
	var _80_v3 _dafny.Array
	_ = _80_v3
	var _nwElement0_3 bool = !(false)
	_ = _nwElement0_3
	var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(15))
	_ = _nw17
	(_nw17).ArraySet1(_nwElement0_3, 0)
	(_nw17).ArraySet1(_78_v1, 1)
	(_nw17).ArraySet1(_78_v1, 2)
	(_nw17).ArraySet1(_78_v1, 3)
	(_nw17).ArraySet1(_78_v1, 4)
	(_nw17).ArraySet1(_78_v1, 5)
	(_nw17).ArraySet1(true, 6)
	(_nw17).ArraySet1(_78_v1, 7)
	(_nw17).ArraySet1(!(_78_v1), 8)
	(_nw17).ArraySet1(_78_v1, 9)
	(_nw17).ArraySet1(_78_v1, 10)
	(_nw17).ArraySet1(_78_v1, 11)
	(_nw17).ArraySet1(_78_v1, 12)
	(_nw17).ArraySet1(_78_v1, 13)
	(_nw17).ArraySet1(_78_v1, 14)
	_80_v3 = _nw17
	var _81_v4 D0
	_ = _81_v4
	_81_v4 = Companion_D0_.Create_DC0_(_77_v0)
	var _82_v6 _dafny.Int
	_ = _82_v6
	_82_v6 = _dafny.IntOfInt64(462)
	var _83_v7 _dafny.Sequence
	_ = _83_v7
	_83_v7 = _dafny.SeqOf(_82_v6, _82_v6, _82_v6)
	var _84_v8 _dafny.CodePoint
	_ = _84_v8
	_84_v8 = _dafny.CodePoint('x')
	var _85_v9 _dafny.Array
	_ = _85_v9
	var _nwElement0_4 _dafny.CodePoint = _84_v8
	_ = _nwElement0_4
	var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(27))
	_ = _nw18
	(_nw18).ArraySet1CodePoint(_nwElement0_4, 0)
	(_nw18).ArraySet1CodePoint(_84_v8, 1)
	(_nw18).ArraySet1CodePoint(_84_v8, 2)
	(_nw18).ArraySet1CodePoint(_dafny.CodePoint('u'), 3)
	(_nw18).ArraySet1CodePoint(_84_v8, 4)
	(_nw18).ArraySet1CodePoint(_84_v8, 5)
	(_nw18).ArraySet1CodePoint(_84_v8, 6)
	(_nw18).ArraySet1CodePoint(_84_v8, 7)
	(_nw18).ArraySet1CodePoint(_84_v8, 8)
	(_nw18).ArraySet1CodePoint(_84_v8, 9)
	(_nw18).ArraySet1CodePoint(_84_v8, 10)
	(_nw18).ArraySet1CodePoint(_84_v8, 11)
	(_nw18).ArraySet1CodePoint(_84_v8, 12)
	(_nw18).ArraySet1CodePoint((_77_v0).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), 13)
	(_nw18).ArraySet1CodePoint(_84_v8, 14)
	(_nw18).ArraySet1CodePoint(_84_v8, 15)
	(_nw18).ArraySet1CodePoint(_84_v8, 16)
	(_nw18).ArraySet1CodePoint(_84_v8, 17)
	(_nw18).ArraySet1CodePoint(_dafny.CodePoint('e'), 18)
	(_nw18).ArraySet1CodePoint(_84_v8, 19)
	(_nw18).ArraySet1CodePoint(_84_v8, 20)
	(_nw18).ArraySet1CodePoint(_84_v8, 21)
	(_nw18).ArraySet1CodePoint(_84_v8, 22)
	(_nw18).ArraySet1CodePoint(_84_v8, 23)
	(_nw18).ArraySet1CodePoint(_84_v8, 24)
	(_nw18).ArraySet1CodePoint(_84_v8, 25)
	(_nw18).ArraySet1CodePoint(_84_v8, 26)
	_85_v9 = _nw18
	var _86_v10 _dafny.Sequence
	_ = _86_v10
	_86_v10 = _dafny.SeqOf(_85_v9, _85_v9, _85_v9, _85_v9, _85_v9)
	var _87_v11 _dafny.MultiSet
	_ = _87_v11
	_87_v11 = _dafny.MultiSetOf(_82_v6)
	var _88_v12 _dafny.Sequence
	_ = _88_v12
	_88_v12 = _dafny.SeqOf(_87_v11)
	var _89_v13 _dafny.Array
	_ = _89_v13
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(29)
	_ = _len0_1
	var _nw19 _dafny.Array
	_ = _nw19
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw19 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Map = (func(_90_v1 bool) func(_dafny.Int) _dafny.Map {
			return func(_91_i1 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v1, _dafny.IntOfInt64(-176))
			}
		})(_78_v1)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw19 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw19).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw19).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_89_v13 = _nw19
	var _92_v14 _dafny.Sequence
	_ = _92_v14
	_92_v14 = _dafny.SeqOf(_80_v3, _80_v3)
	var _93_globalState *GlobalState
	_ = _93_globalState
	var _nw20 *GlobalState = New_GlobalState_()
	_ = _nw20
	_nw20.Ctor__(_77_v0, (_79_v2).Intersection(_79_v2), _80_v3, false, (_81_v4).Dtor_cf0(), _dafny.IntOfInt64(202), (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter2 := _dafny.Iterate((_83_v7).Elements()); ; {
			_compr_0, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _94_v5 _dafny.Int
			_94_v5 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_83_v7, _94_v5) {
				_coll0.Add((_94_v5).Times(_82_v6), _dafny.CodePoint('j'))
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(668))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}((func(_95_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_96_i0 _dafny.Int) _dafny.Int {
			return _95_v6
		}
	})(_82_v6)))).Cardinality()), _84_v8)), false, _dafny.Companion_Sequence_.Concatenate(_86_v10, _86_v10), false, _dafny.IntOfInt64(525), false, false, true, (_88_v12).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_88_v12).Cardinality()))).Uint32()).(_dafny.MultiSet), _89_v13, _dafny.IntOfInt64(588), false, (_92_v14).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_92_v14).Cardinality()))).Uint32()).(_dafny.Array), _77_v0, true, true, _dafny.IntOfInt64(173), true)
	_93_globalState = _nw20
	var _source0 D0 = _81_v4
	_ = _source0
	if _source0.Is_DC1() {
		var _97___mcc_h0 bool = _source0.Get_().(D0_DC1).Cf1
		_ = _97___mcc_h0
		var _98___mcc_h1 bool = _source0.Get_().(D0_DC1).Cf2
		_ = _98___mcc_h1
		var _99___mcc_h2 _dafny.Int = _source0.Get_().(D0_DC1).Cf3
		_ = _99___mcc_h2
		var _100_cf3 _dafny.Int = _99___mcc_h2
		_ = _100_cf3
		var _101_cf2 bool = _98___mcc_h1
		_ = _101_cf2
		var _102_cf1 bool = _97___mcc_h0
		_ = _102_cf1
		var _103_v16 _dafny.MultiSet
		_ = _103_v16
		_103_v16 = _dafny.MultiSetOf(_77_v0, _77_v0)
		var _104_v17 D0
		_ = _104_v17
		_104_v17 = Companion_D0_.Create_DC1_(_78_v1, _78_v1, _100_cf3)
		var _pat_let_tv0 = _101_cf2
		_ = _pat_let_tv0
		var _pat_let_tv1 = _82_v6
		_ = _pat_let_tv1
		var _rhs10 _dafny.Int = ((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-99), _dafny.IntOfInt64(854))); ; {
				_compr_1, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _105_v15 _dafny.Int
				_105_v15 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(-99)).Cmp(_105_v15) <= 0) && ((_105_v15).Cmp(_dafny.IntOfInt64(854)) < 0) {
					_coll1.Add((_105_v15).Times(_82_v6), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-372))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg15 _dafny.Int) interface{} {
							return coer15(arg15)
						}
					}((func(_106_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_107_i2 _dafny.Int) _dafny.CodePoint {
							return _106_v8
						}
					})(_84_v8)))).Cardinality()))
				}
			}
			return _coll1.ToMap()
		}()).Cardinality()).Minus((_103_v16).Cardinality())
		_ = _rhs10
		var _rhs11 _dafny.Int = (func() _dafny.Int {
			if _101_cf2 {
				return Companion_Default___.SafeModuloInt(_82_v6, _82_v6)
			}
			return _dafny.IntOfInt64(620)
		})()
		_ = _rhs11
		var _rhs12 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if _102_cf1 {
				return _82_v6
			}
			return (_83_v7).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_83_v7).Cardinality()))).Uint32()).(_dafny.Int)
		})(), _100_cf3)
		_ = _rhs12
		var _rhs13 _dafny.Int = (func(_pat_let0_0 D0) D0 {
			return func(_110_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let3_0 _dafny.Int) D0 {
					return func(_111_dt__update_hcf3_h0 _dafny.Int) D0 {
						return func(_pat_let4_0 bool) D0 {
							return func(_112_dt__update_hcf2_h1 bool) D0 {
								return Companion_D0_.Create_DC1_((_110_dt__update__tmp_h1).Dtor_cf1(), _112_dt__update_hcf2_h1, _111_dt__update_hcf3_h0)
							}(_pat_let4_0)
						}(false)
					}(_pat_let3_0)
				}((_dafny.Zero).Minus(_pat_let_tv1))
			}(_pat_let0_0)
		}(func(_pat_let1_0 D0) D0 {
			return func(_108_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let2_0 bool) D0 {
					return func(_109_dt__update_hcf2_h0 bool) D0 {
						return Companion_D0_.Create_DC1_((_108_dt__update__tmp_h0).Dtor_cf1(), _109_dt__update_hcf2_h0, (_108_dt__update__tmp_h0).Dtor_cf3())
					}(_pat_let2_0)
				}(_pat_let_tv0)
			}(_pat_let1_0)
		}(_104_v17))).Dtor_cf3()
		_ = _rhs13
		var _rhs14 _dafny.Array = (func() _dafny.Array {
			if !(_78_v1) {
				return _80_v3
			}
			return _80_v3
		})()
		_ = _rhs14
		var _lhs6 *GlobalState = _93_globalState
		_ = _lhs6
		var _lhs7 *GlobalState = _93_globalState
		_ = _lhs7
		_lhs6.F5 = _rhs10
		_100_cf3 = _rhs11
		_lhs7.F5 = _rhs12
		_100_cf3 = _rhs13
		_80_v3 = _rhs14
		(_93_globalState).F5 = _82_v6
		var _113_v18 _dafny.Map
		_ = _113_v18
		_113_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_cf2, _dafny.IntOfInt64(-284))
		var _114_v19 *C0
		_ = _114_v19
		var _nw21 *C0 = New_C0_()
		_ = _nw21
		_nw21.Ctor__(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _dafny.IntOfInt64(-22))).Merge(_113_v18)).Cardinality(), _101_cf2, _77_v0)
		_114_v19 = _nw21
		var _115_v20 T0
		_ = _115_v20
		var _nw22 *C0 = New_C0_()
		_ = _nw22
		_nw22.Ctor__(_82_v6, _101_cf2, _dafny.UnicodeSeqOfUtf8Bytes("nlwarcww"))
		_115_v20 = _nw22
		var _116_v21 _dafny.Map
		_ = _116_v21
		_116_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v20, _100_cf3)
		(_93_globalState).F5 = (func() _dafny.Int {
			if (_116_v21).Contains(_115_v20) {
				return (_116_v21).Get(_115_v20).(_dafny.Int)
			}
			return _100_cf3
		})()
	} else if _source0.Is_DC2() {
		var _117___mcc_h3 _dafny.Array = _source0.Get_().(D0_DC2).Cf4
		_ = _117___mcc_h3
		var _118___mcc_h4 _dafny.Int = _source0.Get_().(D0_DC2).Cf5
		_ = _118___mcc_h4
		var _119_cf5 _dafny.Int = _118___mcc_h4
		_ = _119_cf5
		var _120_cf4 _dafny.Array = _117___mcc_h3
		_ = _120_cf4
		var _121_v22 *C0
		_ = _121_v22
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__(_dafny.IntOfUint32((_83_v7).Cardinality()), _78_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_122_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_123_i3 _dafny.Int) _dafny.CodePoint {
				return _122_v8
			}
		})(_84_v8))))
		_121_v22 = _nw23
		var _124_v23 _dafny.Map
		_ = _124_v23
		_124_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v22, _82_v6)
		var _125_v24 *C0
		_ = _125_v24
		var _nw24 *C0 = New_C0_()
		_ = _nw24
		_nw24.Ctor__((_124_v23).Cardinality(), _78_v1, _77_v0)
		_125_v24 = _nw24
		var _126_v25 D1
		_ = _126_v25
		_126_v25 = Companion_D1_.Create_DC4_(_121_v22)
		var _127_v26 _dafny.Map
		_ = _127_v26
		_127_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_126_v25).Dtor_cf7(), _121_v22)
		var _128_v27 _dafny.Sequence
		_ = _128_v27
		_128_v27 = _dafny.SeqOf(_127_v26, _127_v26, (_127_v26).Merge(_127_v26), (_127_v26).Merge(_127_v26), _127_v26)
		_128_v27 = _128_v27
		(_93_globalState).F20 = _78_v1
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_80_v3), 0))
		_ = _index8
		(_80_v3).ArraySet1(_78_v1, (_index8).Int())
		var _129_v28 _dafny.Sequence
		_ = _129_v28
		_129_v28 = _dafny.SeqOf(_78_v1, _78_v1)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_80_v3), 0))
		_ = _index9
		(_80_v3).ArraySet1((_129_v28).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_129_v28).Cardinality()))).Uint32()).(bool), (_index9).Int())
		var _130_v29 _dafny.Set
		_ = _130_v29
		_130_v29 = _dafny.SetOf(_82_v6, (_121_v22).F26())
		var _131_v30 _dafny.Map
		_ = _131_v30
		_131_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_cf5, !((_80_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_80_v3), 0))).Int()).(bool)))
		var _rhs15 _dafny.CodePoint = Companion_Default___.Fm4((_130_v29).Cardinality(), _82_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(102))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_132_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_133_i4 _dafny.Int) _dafny.CodePoint {
				return _132_v8
			}
		})(_84_v8))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_78_v1, _78_v1), _129_v28), _93_globalState)
		_ = _rhs15
		var _rhs16 _dafny.Int = ((_125_v24).F26()).Times((_131_v30).Cardinality())
		_ = _rhs16
		var _lhs8 *GlobalState = _93_globalState
		_ = _lhs8
		_84_v8 = _rhs15
		_lhs8.F5 = _rhs16
	} else if _source0.Is_DC0() {
		var _134___mcc_h5 _dafny.Sequence = _source0.Get_().(D0_DC0).Cf0
		_ = _134___mcc_h5
		var _135_cf0 _dafny.Sequence = _134___mcc_h5
		_ = _135_cf0
		(_93_globalState).F5 = _dafny.IntOfUint32((_77_v0).Cardinality())
		var _136_v31 *C0
		_ = _136_v31
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__(_82_v6, _78_v1, _135_cf0)
		_136_v31 = _nw25
		var _137_v32 _dafny.Sequence
		_ = _137_v32
		_137_v32 = _dafny.SeqOf(_136_v31)
		var _138_v33 _dafny.Sequence
		_ = _138_v33
		_138_v33 = _dafny.SeqOf(_137_v32)
		_138_v33 = _dafny.SeqOf(_137_v32, _dafny.Companion_Sequence_.Update(_137_v32, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.IntOfUint32((_137_v32).Cardinality()))).Uint32(), _136_v31), _137_v32)
		_80_v3 = (func() _dafny.Array {
			if !(_78_v1) {
				return _80_v3
			}
			return _80_v3
		})()
		var _139_v34 _dafny.Sequence
		_ = _139_v34
		var _140_v35 _dafny.Array
		_ = _140_v35
		var _141_v36 _dafny.Int
		_ = _141_v36
		var _142_v37 bool
		_ = _142_v37
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 _dafny.Array
		_ = _out1
		var _out2 _dafny.Int
		_ = _out2
		var _out3 bool
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M0(_93_globalState)
		_139_v34 = _out0
		_140_v35 = _out1
		_141_v36 = _out2
		_142_v37 = _out3
	} else {
		var _143___mcc_h6 D0 = _source0.Get_().(D0_DC3).Cf6
		_ = _143___mcc_h6
		var _144_cf6 D0 = _143___mcc_h6
		_ = _144_cf6
		var _source1 D0 = Companion_D0_.Create_DC0_((func() _dafny.Sequence {
			if true {
				return _77_v0
			}
			return _77_v0
		})())
		_ = _source1
		if _source1.Is_DC1() {
			var _145___mcc_h7 bool = _source1.Get_().(D0_DC1).Cf1
			_ = _145___mcc_h7
			var _146___mcc_h8 bool = _source1.Get_().(D0_DC1).Cf2
			_ = _146___mcc_h8
			var _147___mcc_h9 _dafny.Int = _source1.Get_().(D0_DC1).Cf3
			_ = _147___mcc_h9
			var _148_cf3 _dafny.Int = _147___mcc_h9
			_ = _148_cf3
			var _149_cf2 bool = _146___mcc_h8
			_ = _149_cf2
			var _150_cf1 bool = _145___mcc_h7
			_ = _150_cf1
			var _151_v38 _dafny.Sequence
			_ = _151_v38
			_151_v38 = _dafny.SeqOf(!(_149_cf2))
			_77_v0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_77_v0, _77_v0), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_82_v6, (_dafny.Zero).Minus(_82_v6)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_77_v0, _77_v0)).Cardinality()))).Uint32(), Companion_Default___.Fm4(_dafny.IntOfUint32((_77_v0).Cardinality()), _148_cf3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_152_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_153_i5 _dafny.Int) _dafny.CodePoint {
					return _152_v8
				}
			})(_84_v8))), _151_v38, _93_globalState))
			_149_cf2 = Companion_Default___.Fm5(!(_78_v1) || (_150_cf1), _150_cf1, _148_cf3, _82_v6, _93_globalState)
			_149_cf2 = _78_v1
			var _154_v39 T1
			_ = _154_v39
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__(_82_v6, _150_cf1, _77_v0)
			_154_v39 = _nw26
			var _155_v40 _dafny.Map
			_ = _155_v40
			_155_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v39, Companion_Default___.Fm5(true, _149_cf2, _148_cf3, _148_cf3, _93_globalState))
			(_93_globalState).F5 = Companion_Default___.SafeModuloInt(Companion_Default___.Fm6(!(_78_v1), _149_cf2, _149_cf2, _93_globalState), ((_155_v40).Update(_154_v39, false)).Cardinality())
		} else if _source1.Is_DC2() {
			var _156___mcc_h10 _dafny.Array = _source1.Get_().(D0_DC2).Cf4
			_ = _156___mcc_h10
			var _157___mcc_h11 _dafny.Int = _source1.Get_().(D0_DC2).Cf5
			_ = _157___mcc_h11
			var _158_cf5 _dafny.Int = _157___mcc_h11
			_ = _158_cf5
			var _159_cf4 _dafny.Array = _156___mcc_h10
			_ = _159_cf4
			var _160_v41 _dafny.MultiSet
			_ = _160_v41
			_160_v41 = _dafny.MultiSetOf(_78_v1, false)
			var _161_v42 T1
			_ = _161_v42
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__((_160_v41).Cardinality(), _78_v1, _77_v0)
			_161_v42 = _nw27
			var _162_v43 D1
			_ = _162_v43
			_162_v43 = Companion_D1_.Create_DC6_(_158_cf5, _161_v42, !(true), _78_v1)
			var _163_v44 _dafny.Map
			_ = _163_v44
			_163_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _84_v8)
			var _164_v45 _dafny.Map
			_ = _164_v45
			_164_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v6, Companion_Default___.Fm5(_78_v1, (_161_v42).F24(), _dafny.IntOfInt64(609), _158_cf5, _93_globalState))
			var _165_v46 _dafny.Map
			_ = _165_v46
			_165_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_cf5, _158_cf5)
			var _166_v48 *C0
			_ = _166_v48
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__(_158_cf5, !(!(_78_v1)), (_161_v42).F25())
			_166_v48 = _nw28
			var _167_v49 _dafny.Sequence
			_ = _167_v49
			_167_v49 = _dafny.SeqOf(_166_v48, _166_v48, _166_v48)
			var _168_v50 D0
			_ = _168_v50
			_168_v50 = Companion_D0_.Create_DC2_(_159_cf4, _82_v6)
			var _169_v51 _dafny.MultiSet
			_ = _169_v51
			_169_v51 = _dafny.MultiSetOf(_165_v46, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v6, _dafny.IntOfUint32((_83_v7).Cardinality())), func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter4 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfUint32((_167_v49).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_168_v50).Dtor_cf5())).Cardinality()), (_166_v48).F26(), (_162_v43).Dtor_cf10(), (_166_v48).F26())).Elements()); ; {
					_compr_2, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _170_v47 _dafny.Int
					_170_v47 = interface{}(_compr_2).(_dafny.Int)
					if (_dafny.SetOf(_dafny.IntOfUint32((_167_v49).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_168_v50).Dtor_cf5())).Cardinality()), (_166_v48).F26(), (_162_v43).Dtor_cf10(), (_166_v48).F26())).Contains(_170_v47) {
						_coll2.Add(Companion_Default___.SafeDivisionInt(_170_v47, _158_cf5), _82_v6)
					}
				}
				return _coll2.ToMap()
			}(), _165_v46)
			var _171_v52 _dafny.Map
			_ = _171_v52
			_171_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, (_169_v51).Cardinality())
			var _pat_let_tv2 = _161_v42
			_ = _pat_let_tv2
			var _pat_let_tv3 = _78_v1
			_ = _pat_let_tv3
			var _pat_let_tv4 = _161_v42
			_ = _pat_let_tv4
			var _pat_let_tv5 = _158_cf5
			_ = _pat_let_tv5
			(_93_globalState).F20 = (func() bool {
				if Companion_Default___.Fm5((func(_pat_let5_0 D1) D1 {
					return func(_175_dt__update__tmp_h3 D1) D1 {
						return func(_pat_let9_0 bool) D1 {
							return func(_176_dt__update_hcf13_h0 bool) D1 {
								return func(_pat_let10_0 T1) D1 {
									return func(_177_dt__update_hcf11_h1 T1) D1 {
										return func(_pat_let11_0 _dafny.Int) D1 {
											return func(_178_dt__update_hcf10_h1 _dafny.Int) D1 {
												return Companion_D1_.Create_DC6_(_178_dt__update_hcf10_h1, _177_dt__update_hcf11_h1, (_175_dt__update__tmp_h3).Dtor_cf12(), _176_dt__update_hcf13_h0)
											}(_pat_let11_0)
										}(_pat_let_tv5)
									}(_pat_let10_0)
								}(_pat_let_tv4)
							}(_pat_let9_0)
						}(!(_pat_let_tv3))
					}(_pat_let5_0)
				}(func(_pat_let6_0 D1) D1 {
					return func(_172_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let7_0 T1) D1 {
							return func(_173_dt__update_hcf11_h0 T1) D1 {
								return func(_pat_let8_0 _dafny.Int) D1 {
									return func(_174_dt__update_hcf10_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC6_(_174_dt__update_hcf10_h0, _173_dt__update_hcf11_h0, (_172_dt__update__tmp_h2).Dtor_cf12(), (_172_dt__update__tmp_h2).Dtor_cf13())
									}(_pat_let8_0)
								}(_dafny.IntOfInt64(660))
							}(_pat_let7_0)
						}(_pat_let_tv2)
					}(_pat_let6_0)
				}(_162_v43))).Dtor_cf12(), _78_v1, Companion_Default___.Fm6((_161_v42).Fm1(_163_v44, _164_v45, _83_v7, _93_globalState), _78_v1, _78_v1, _93_globalState), _158_cf5, _93_globalState) {
					return (_161_v42).F24()
				}
				return !(_171_v52).Equals(_171_v52)
			})()
			var _179_v53 _dafny.Set
			_ = _179_v53
			_179_v53 = _dafny.SetOf(_158_cf5)
			var _180_v54 _dafny.Sequence
			_ = _180_v54
			_180_v54 = _dafny.SeqOf((_161_v42).F24(), (_161_v42).F24())
			var _rhs17 _dafny.Set = (func() _dafny.Set {
				if !(_78_v1) {
					return _179_v53
				}
				return (_dafny.SetOf((_166_v48).F26())).Difference(_179_v53)
			})()
			_ = _rhs17
			var _rhs18 bool = false
			_ = _rhs18
			var _rhs19 bool = _dafny.Companion_Sequence_.Equal(_180_v54, _dafny.Companion_Sequence_.Concatenate(_180_v54, _180_v54))
			_ = _rhs19
			var _rhs20 bool = (_161_v42).F24()
			_ = _rhs20
			var _lhs9 *GlobalState = _93_globalState
			_ = _lhs9
			var _lhs10 *GlobalState = _93_globalState
			_ = _lhs10
			var _lhs11 *GlobalState = _93_globalState
			_ = _lhs11
			_179_v53 = _rhs17
			_lhs9.F20 = _rhs18
			_lhs10.F21 = _rhs19
			_lhs11.F21 = _rhs20
			var _181_v55 _dafny.Map
			_ = _181_v55
			_181_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v53, (_166_v48).F26())
			var _182_v56 *C0
			_ = _182_v56
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__((_166_v48).F26(), (_181_v55).Equals(_181_v55), _77_v0)
			_182_v56 = _nw29
			var _183_v57 _dafny.Sequence
			_ = _183_v57
			var _184_v58 _dafny.Array
			_ = _184_v58
			var _185_v59 _dafny.Int
			_ = _185_v59
			var _186_v60 bool
			_ = _186_v60
			var _out4 _dafny.Sequence
			_ = _out4
			var _out5 _dafny.Array
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			var _out7 bool
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0(_93_globalState)
			_183_v57 = _out4
			_184_v58 = _out5
			_185_v59 = _out6
			_186_v60 = _out7
		} else if _source1.Is_DC0() {
			var _187___mcc_h12 _dafny.Sequence = _source1.Get_().(D0_DC0).Cf0
			_ = _187___mcc_h12
			var _188_cf0 _dafny.Sequence = _187___mcc_h12
			_ = _188_cf0
			(_93_globalState).F20 = _78_v1
			var _189_v61 _dafny.Array
			_ = _189_v61
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw30
			_189_v61 = _nw30
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_189_v61), 0))
			_ = _index10
			(_189_v61).ArraySet1(_82_v6, (_index10).Int())
			var _190_v62 *C0
			_ = _190_v62
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(_82_v6, false, _188_cf0)
			_190_v62 = _nw31
			var _191_v63 _dafny.Sequence
			_ = _191_v63
			_191_v63 = _dafny.SeqOf(_190_v62, _190_v62)
			var _192_v64 _dafny.Sequence
			_ = _192_v64
			_192_v64 = _dafny.SeqOf(_191_v63)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_189_v61), 0))
			_ = _index11
			(_189_v61).ArraySet1((_82_v6).Plus((func() _dafny.Int {
				if false {
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_192_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-355), _dafny.IntOfUint32((_192_v64).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_190_v62).Fm3(_82_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(609))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg19 _dafny.Int) interface{} {
							return coer19(arg19)
						}
					}((func(_193_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_194_i6 _dafny.Int) _dafny.Sequence {
							return _193_v7
						}
					})(_83_v7))), _82_v6, _dafny.Companion_Sequence_.Update(_188_cf0, (Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_188_cf0).Cardinality()))).Uint32(), _84_v8), _93_globalState), _dafny.IntOfUint32(((_192_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-355), _dafny.IntOfUint32((_192_v64).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _190_v62)).Cardinality())
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_188_cf0).Cardinality()))
			})()), (_index11).Int())
			_78_v1 = (_dafny.IntOfInt64(659)).Cmp(_dafny.IntOfUint32((_188_cf0).Cardinality())) > 0
			(_93_globalState).F21 = _78_v1
		} else {
			var _195___mcc_h13 D0 = _source1.Get_().(D0_DC3).Cf6
			_ = _195___mcc_h13
			var _196_cf6 D0 = _195___mcc_h13
			_ = _196_cf6
			var _197_v65 _dafny.Sequence
			_ = _197_v65
			var _198_v66 _dafny.Array
			_ = _198_v66
			var _199_v67 _dafny.Int
			_ = _199_v67
			var _200_v68 bool
			_ = _200_v68
			var _out8 _dafny.Sequence
			_ = _out8
			var _out9 _dafny.Array
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 bool
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0(_93_globalState)
			_197_v65 = _out8
			_198_v66 = _out9
			_199_v67 = _out10
			_200_v68 = _out11
			(_93_globalState).F5 = _199_v67
			_78_v1 = !(_78_v1)
			_82_v6 = Companion_Default___.SafeModuloInt(_199_v67, _dafny.IntOfInt64(874))
		}
		_78_v1 = false
		var _201_v69 _dafny.Sequence
		_ = _201_v69
		_201_v69 = _dafny.SeqOf(_78_v1)
		var _202_v71 *C0
		_ = _202_v71
		var _nw32 *C0 = New_C0_()
		_ = _nw32
		_nw32.Ctor__(_dafny.IntOfInt64(953), _78_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-753))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_203_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_204_i7 _dafny.Int) _dafny.CodePoint {
				return _203_v8
			}
		})(_84_v8))))
		_202_v71 = _nw32
		var _205_v72 _dafny.Sequence
		_ = _205_v72
		_205_v72 = _dafny.SeqOf(_202_v71, _202_v71, _202_v71, _202_v71)
		var _206_v73 _dafny.Map
		_ = _206_v73
		_206_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_202_v71).F26(), _82_v6)
		var _207_v74 _dafny.Array
		_ = _207_v74
		var _nwElement0_5 _dafny.Int = _82_v6
		_ = _nwElement0_5
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_5, 0)
		(_nw33).ArraySet1(_82_v6, 1)
		(_nw33).ArraySet1(_82_v6, 2)
		(_nw33).ArraySet1(_82_v6, 3)
		(_nw33).ArraySet1(_dafny.IntOfUint32((_201_v69).Cardinality()), 4)
		(_nw33).ArraySet1(_82_v6, 5)
		(_nw33).ArraySet1(_dafny.IntOfInt64(225), 6)
		(_nw33).ArraySet1(_dafny.IntOfUint32((_77_v0).Cardinality()), 7)
		(_nw33).ArraySet1(_dafny.IntOfInt64(373), 8)
		(_nw33).ArraySet1((func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(970), _dafny.IntOfInt64(840))); ; {
				_compr_3, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _208_v70 _dafny.Int
				_208_v70 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(970)).Cmp(_208_v70) <= 0) && ((_208_v70).Cmp(_dafny.IntOfInt64(840)) < 0) {
					_coll3.Add((_208_v70).Times(_82_v6))
				}
			}
			return _coll3.ToSet()
		}()).Cardinality(), 9)
		(_nw33).ArraySet1(_82_v6, 10)
		(_nw33).ArraySet1((_87_v11).Cardinality(), 11)
		(_nw33).ArraySet1(_dafny.IntOfUint32((_77_v0).Cardinality()), 12)
		(_nw33).ArraySet1(_82_v6, 13)
		(_nw33).ArraySet1(_dafny.IntOfUint32((_205_v72).Cardinality()), 14)
		(_nw33).ArraySet1((_202_v71).F26(), 15)
		(_nw33).ArraySet1(_82_v6, 16)
		(_nw33).ArraySet1(_dafny.IntOfInt64(768), 17)
		(_nw33).ArraySet1(_82_v6, 18)
		(_nw33).ArraySet1((_202_v71).F26(), 19)
		(_nw33).ArraySet1(_82_v6, 20)
		(_nw33).ArraySet1((_202_v71).F26(), 21)
		(_nw33).ArraySet1((_206_v73).Cardinality(), 22)
		(_nw33).ArraySet1(_dafny.IntOfInt64(-224), 23)
		(_nw33).ArraySet1(_82_v6, 24)
		_207_v74 = _nw33
		var _209_v75 D0
		_ = _209_v75
		_209_v75 = Companion_D0_.Create_DC2_(_207_v74, _82_v6)
		var _source2 D0 = _209_v75
		_ = _source2
		if _source2.Is_DC1() {
			var _210___mcc_h14 bool = _source2.Get_().(D0_DC1).Cf1
			_ = _210___mcc_h14
			var _211___mcc_h15 bool = _source2.Get_().(D0_DC1).Cf2
			_ = _211___mcc_h15
			var _212___mcc_h16 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
			_ = _212___mcc_h16
			var _213_cf3 _dafny.Int = _212___mcc_h16
			_ = _213_cf3
			var _214_cf2 bool = _211___mcc_h15
			_ = _214_cf2
			var _215_cf1 bool = _210___mcc_h14
			_ = _215_cf1
			var _216_v76 _dafny.Map
			_ = _216_v76
			_216_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_202_v71).F26()).Times((_83_v7).Select((Companion_Default___.SafeIndex(_213_cf3, _dafny.IntOfUint32((_83_v7).Cardinality()))).Uint32()).(_dafny.Int)), _80_v3)
			var _217_v77 _dafny.Map
			_ = _217_v77
			_217_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_cf1, _dafny.IntOfInt64(-730))
			_216_v76 = (_216_v76).Update((_217_v77).Cardinality(), _80_v3)
			var _218_v78 _dafny.Map
			_ = _218_v78
			_218_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_202_v71).F26()).Minus(_213_cf3), _214_cf2)
			_218_v78 = (_218_v78).Update((_202_v71).F26(), !(_215_cf1))
			_214_cf2 = _215_cf1
			var _219_v79 _dafny.Sequence
			_ = _219_v79
			var _220_v80 _dafny.Array
			_ = _220_v80
			var _221_v81 _dafny.Int
			_ = _221_v81
			var _222_v82 bool
			_ = _222_v82
			var _out12 _dafny.Sequence
			_ = _out12
			var _out13 _dafny.Array
			_ = _out13
			var _out14 _dafny.Int
			_ = _out14
			var _out15 bool
			_ = _out15
			_out12, _out13, _out14, _out15 = Companion_Default___.M0(_93_globalState)
			_219_v79 = _out12
			_220_v80 = _out13
			_221_v81 = _out14
			_222_v82 = _out15
		} else if _source2.Is_DC2() {
			var _223___mcc_h17 _dafny.Array = _source2.Get_().(D0_DC2).Cf4
			_ = _223___mcc_h17
			var _224___mcc_h18 _dafny.Int = _source2.Get_().(D0_DC2).Cf5
			_ = _224___mcc_h18
			var _225_cf5 _dafny.Int = _224___mcc_h18
			_ = _225_cf5
			var _226_cf4 _dafny.Array = _223___mcc_h17
			_ = _226_cf4
			var _227_v83 _dafny.Map
			_ = _227_v83
			_227_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, (_79_v2).Cardinality())
			var _228_v84 _dafny.Map
			_ = _228_v84
			_228_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _82_v6)
			var _229_v85 _dafny.MultiSet
			_ = _229_v85
			_229_v85 = _dafny.MultiSetOf(_228_v84)
			var _230_v86 _dafny.Map
			_ = _230_v86
			_230_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_227_v83).Contains(_79_v2) {
					return (_227_v83).Get(_79_v2).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_83_v7).Cardinality())
			})(), (_229_v85).IsSubsetOf((_229_v85).Update(_228_v84, Companion_Default___.Abs((_202_v71).F26()))))
			_230_v86 = (_230_v86).Update((_202_v71).F26(), false)
			var _231_v87 *C0
			_ = _231_v87
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__((_dafny.Zero).Minus((_202_v71).F26()), _78_v1, _77_v0)
			_231_v87 = _nw34
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_85_v9), 0))
			_ = _index12
			(_85_v9).ArraySet1CodePoint(_84_v8, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_85_v9), 0))
			_ = _index13
			(_85_v9).ArraySet1CodePoint(_84_v8, (_index13).Int())
			_79_v2 = _79_v2
		} else if _source2.Is_DC0() {
			var _232___mcc_h19 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf0
			_ = _232___mcc_h19
			var _233_cf0 _dafny.Sequence = _232___mcc_h19
			_ = _233_cf0
			var _234_v88 _dafny.Sequence
			_ = _234_v88
			var _235_v89 _dafny.Array
			_ = _235_v89
			var _236_v90 _dafny.Int
			_ = _236_v90
			var _237_v91 bool
			_ = _237_v91
			var _out16 _dafny.Sequence
			_ = _out16
			var _out17 _dafny.Array
			_ = _out17
			var _out18 _dafny.Int
			_ = _out18
			var _out19 bool
			_ = _out19
			_out16, _out17, _out18, _out19 = Companion_Default___.M0(_93_globalState)
			_234_v88 = _out16
			_235_v89 = _out17
			_236_v90 = _out18
			_237_v91 = _out19
			var _238_v92 *C0
			_ = _238_v92
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__(_236_v90, _78_v1, Companion_Default___.Fm7(_78_v1, _236_v90, (_83_v7).Select((Companion_Default___.SafeIndex(_236_v90, _dafny.IntOfUint32((_83_v7).Cardinality()))).Uint32()).(_dafny.Int), _237_v91, _93_globalState))
			_238_v92 = _nw35
			var _239_v93 _dafny.Map
			_ = _239_v93
			_239_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v90, _237_v91)
			_239_v93 = (_239_v93).Update(((_202_v71).F26()).Plus(_82_v6), !((_82_v6).Cmp(_dafny.IntOfUint32((_201_v69).Cardinality())) > 0))
			(_93_globalState).F20 = _237_v91
		} else {
			var _240___mcc_h20 D0 = _source2.Get_().(D0_DC3).Cf6
			_ = _240___mcc_h20
			var _241_cf6 D0 = _240___mcc_h20
			_ = _241_cf6
			var _242_v94 _dafny.Array
			_ = _242_v94
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(13))
			_ = _nw36
			_242_v94 = _nw36
			var _243_v95 T1
			_ = _243_v95
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__((_202_v71).F26(), !(_78_v1), _77_v0)
			_243_v95 = _nw37
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_242_v94), 0))
			_ = _index14
			(_242_v94).ArraySet1(Companion_D1_.Create_DC6_((_202_v71).F26(), _243_v95, _78_v1, (_243_v95).F24()), (_index14).Int())
			var _244_v96 _dafny.Map
			_ = _244_v96
			_244_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_202_v71).F26(), true)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_242_v94), 0))
			_ = _index15
			(_242_v94).ArraySet1(Companion_D1_.Create_DC6_((_202_v71).F26(), _243_v95, (_243_v95).F24(), (_243_v95).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_243_v95).F24(), _84_v8), _244_v96, _83_v7, _93_globalState)), (_index15).Int())
			var _245_v98 _dafny.Map
			_ = _245_v98
			_245_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-822), _dafny.IntOfInt64(224))); ; {
					_compr_4, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _246_v97 _dafny.Int
					_246_v97 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(-822)).Cmp(_246_v97) <= 0) && ((_246_v97).Cmp(_dafny.IntOfInt64(224)) < 0) {
						_coll4.Add(Companion_Default___.SafeDivisionInt(_246_v97, (_202_v71).F26()), (_243_v95).F24())
					}
				}
				return _coll4.ToMap()
			}(), _dafny.MultiSetOf((_243_v95).F24(), _78_v1, _78_v1))
			var _247_v99 _dafny.MultiSet
			_ = _247_v99
			_247_v99 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_77_v0, (Companion_Default___.SafeIndex((_245_v98).Cardinality(), _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32(), _dafny.CodePoint('k')))
			_247_v99 = (_247_v99).Intersection(_247_v99)
			_78_v1 = (_dafny.IntOfInt64(551)).Cmp(((_202_v71).F26()).Times(_dafny.IntOfInt64(564))) <= 0
			var _248_v100 _dafny.Array
			_ = _248_v100
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw38
			_248_v100 = _nw38
			var _249_v101 D2
			_ = _249_v101
			_249_v101 = Companion_D2_.Create_DC8_(_80_v3)
			var _250_v102 _dafny.Array
			_ = _250_v102
			var _nwElement0_6 _dafny.Array = _80_v3
			_ = _nwElement0_6
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(14))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_6, 0)
			(_nw39).ArraySet1(_80_v3, 1)
			(_nw39).ArraySet1(_80_v3, 2)
			(_nw39).ArraySet1(_80_v3, 3)
			(_nw39).ArraySet1(_80_v3, 4)
			(_nw39).ArraySet1((_249_v101).Dtor_cf15(), 5)
			(_nw39).ArraySet1(_80_v3, 6)
			(_nw39).ArraySet1(_80_v3, 7)
			(_nw39).ArraySet1(_80_v3, 8)
			(_nw39).ArraySet1(_80_v3, 9)
			(_nw39).ArraySet1(_80_v3, 10)
			(_nw39).ArraySet1(_80_v3, 11)
			(_nw39).ArraySet1(_80_v3, 12)
			(_nw39).ArraySet1(_80_v3, 13)
			_250_v102 = _nw39
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_248_v100), 0))
			_ = _index16
			(_248_v100).ArraySet1(_250_v102, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_248_v100), 0))
			_ = _index17
			(_248_v100).ArraySet1((func() _dafny.Array {
				if (_243_v95).F24() {
					return _250_v102
				}
				return _250_v102
			})(), (_index17).Int())
		}
		var _251_v103 _dafny.Map
		_ = _251_v103
		_251_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _84_v8)
		var _252_v104 _dafny.Map
		_ = _252_v104
		_252_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.Fm6(_78_v1, _78_v1, _78_v1, _93_globalState)), _201_v69)
		(_93_globalState).F21 = ((func() bool {
			if (_202_v71).Fm1(_251_v103, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_252_v104).Cardinality(), _78_v1), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_253_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_254_i8 _dafny.Int) _dafny.Int {
					return _253_v6
				}
			})(_82_v6))), _93_globalState) {
				return _78_v1
			}
			return !(false)
		})()) && (_78_v1)
	}
	var _255_v105 _dafny.Array
	_ = _255_v105
	var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
	_ = _nw40
	_255_v105 = _nw40
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_255_v105), 0))); ; {
		_guard_loop_2, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _256_i9 _dafny.Int
		_256_i9 = interface{}(_guard_loop_2).(_dafny.Int)
		if (true) && (((_256_i9).Sign() != -1) && ((_256_i9).Cmp(_dafny.ArrayLen((_255_v105), 0)) < 0)) {
			(_255_v105).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _84_v8)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _dafny.CodePoint('t'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _84_v8))), (_256_i9).Int())
		}
	}
	var _257_v106 _dafny.Sequence
	_ = _257_v106
	_257_v106 = _dafny.SeqOf(_78_v1, _78_v1)
	(_93_globalState).F21 = (!((_257_v106).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_77_v0).Cardinality()), _dafny.IntOfUint32((_257_v106).Cardinality()))).Uint32()).(bool))) == (_78_v1)
	var _258_v107 _dafny.Map
	_ = _258_v107
	_258_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(49), _82_v6)
	var _259_v108 D1
	_ = _259_v108
	_259_v108 = Companion_D1_.Create_DC5_(_78_v1, _82_v6)
	var _pat_let_tv6 = _82_v6
	_ = _pat_let_tv6
	var _pat_let_tv7 = _82_v6
	_ = _pat_let_tv7
	var _pat_let_tv8 = _82_v6
	_ = _pat_let_tv8
	var _pat_let_tv9 = _82_v6
	_ = _pat_let_tv9
	var _pat_let_tv10 = _82_v6
	_ = _pat_let_tv10
	var _pat_let_tv11 = _82_v6
	_ = _pat_let_tv11
	var _rhs21 bool = (_78_v1) == (_78_v1)
	_ = _rhs21
	var _rhs22 _dafny.Int = (func() _dafny.Int {
		if (_258_v107).Contains(_82_v6) {
			return (_258_v107).Get(_82_v6).(_dafny.Int)
		}
		return _82_v6
	})()
	_ = _rhs22
	var _rhs23 _dafny.Int = func(_source3 D1) _dafny.Int {
		if _source3.Is_DC5() {
			var _260___mcc_h21 bool = _source3.Get_().(D1_DC5).Cf8
			_ = _260___mcc_h21
			var _261___mcc_h22 _dafny.Int = _source3.Get_().(D1_DC5).Cf9
			_ = _261___mcc_h22
			var _262_cf9 _dafny.Int = _261___mcc_h22
			_ = _262_cf9
			var _263_cf8 bool = _260___mcc_h21
			_ = _263_cf8
			return (_pat_let_tv6).Times(_pat_let_tv7)
		} else if _source3.Is_DC6() {
			var _264___mcc_h23 _dafny.Int = _source3.Get_().(D1_DC6).Cf10
			_ = _264___mcc_h23
			var _265___mcc_h24 T1 = _source3.Get_().(D1_DC6).Cf11
			_ = _265___mcc_h24
			var _266___mcc_h25 bool = _source3.Get_().(D1_DC6).Cf12
			_ = _266___mcc_h25
			var _267___mcc_h26 bool = _source3.Get_().(D1_DC6).Cf13
			_ = _267___mcc_h26
			var _268_cf13 bool = _267___mcc_h26
			_ = _268_cf13
			var _269_cf12 bool = _266___mcc_h25
			_ = _269_cf12
			var _270_cf11 T1 = _265___mcc_h24
			_ = _270_cf11
			var _271_cf10 _dafny.Int = _264___mcc_h23
			_ = _271_cf10
			return (_pat_let_tv8).Times(_pat_let_tv9)
		} else if _source3.Is_DC4() {
			var _272___mcc_h27 *C0 = _source3.Get_().(D1_DC4).Cf7
			_ = _272___mcc_h27
			var _273_cf7 *C0 = _272___mcc_h27
			_ = _273_cf7
			return (_273_cf7).F26()
		} else {
			var _274___mcc_h28 D1 = _source3.Get_().(D1_DC7).Cf14
			_ = _274___mcc_h28
			var _275_cf14 D1 = _274___mcc_h28
			_ = _275_cf14
			return (func() _dafny.Set {
				var _coll5 = _dafny.NewBuilder()
				_ = _coll5
				for _iter8 := _dafny.Iterate((_dafny.SetOf(_pat_let_tv10)).Elements()); ; {
					_compr_5, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _276_v109 _dafny.Int
					_276_v109 = interface{}(_compr_5).(_dafny.Int)
					if (_dafny.SetOf(_pat_let_tv11)).Contains(_276_v109) {
						_coll5.Add(Companion_Default___.SafeDivisionInt(_276_v109, _dafny.IntOfInt64(53)))
					}
				}
				return _coll5.ToSet()
			}()).Cardinality()
		}
	}(_259_v108)
	_ = _rhs23
	var _lhs12 *GlobalState = _93_globalState
	_ = _lhs12
	_lhs12.F20 = _rhs21
	_82_v6 = _rhs22
	_82_v6 = _rhs23
	var _277_v110 T1
	_ = _277_v110
	var _nw41 *C0 = New_C0_()
	_ = _nw41
	_nw41.Ctor__(_82_v6, _78_v1, _77_v0)
	_277_v110 = _nw41
	var _278_v111 D1
	_ = _278_v111
	_278_v111 = Companion_D1_.Create_DC6_(_82_v6, _277_v110, true, (_277_v110).F24())
	var _279_v112 _dafny.Map
	_ = _279_v112
	_279_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, ((_278_v111).Dtor_cf10()).Cmp(_82_v6) > 0)
	_279_v112 = (_279_v112).Update((_277_v110).F24(), (_82_v6).Cmp(_82_v6) != 0)
	var _hi0 _dafny.Int = _82_v6
	_ = _hi0
	for _280_i10 := _dafny.IntOfInt64(658); _280_i10.Cmp(_hi0) < 0; _280_i10 = _280_i10.Plus(_dafny.One) {
		_82_v6 = _280_i10
		var _281_v113 _dafny.Array
		_ = _281_v113
		var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw42
		_281_v113 = _nw42
		_281_v113 = _281_v113
		var _282_v114 _dafny.Map
		_ = _282_v114
		_282_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v6, _277_v110)
		_282_v114 = (_282_v114).Update((_dafny.Zero).Minus(_dafny.IntOfUint32((_83_v7).Cardinality())), _277_v110)
		_279_v112 = (_279_v112).Update(_78_v1, (_257_v106).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_257_v106).Cardinality()))).Uint32()).(bool))
	}
	var _283_v115 _dafny.Set
	_ = _283_v115
	_283_v115 = _dafny.SetOf(_84_v8)
	var _284_v116 _dafny.Set
	_ = _284_v116
	_284_v116 = _dafny.SetOf(_283_v115)
	var _285_i11 _dafny.Int
	_ = _285_i11
	_285_i11 = _dafny.Zero
	{
		for ((_284_v116).Union(_dafny.SetOf(_283_v115, _283_v115))).IsProperSubsetOf(_284_v116) {
			{
				if (_285_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_285_i11 = (_285_i11).Plus(_dafny.One)
				var _286_v117 _dafny.Array
				_ = _286_v117
				var _nwElement0_7 _dafny.Int = _82_v6
				_ = _nwElement0_7
				var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.One)
				_ = _nw43
				(_nw43).ArraySet1(_nwElement0_7, 0)
				_286_v117 = _nw43
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_286_v117), 0))
				_ = _index18
				(_286_v117).ArraySet1(_82_v6, (_index18).Int())
				var _287_v118 _dafny.Map
				_ = _287_v118
				_287_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate((_277_v110).F25(), _77_v0), _82_v6)
				var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_286_v117), 0))
				_ = _index19
				var _rhs24 bool = (_257_v106).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_82_v6, _82_v6), _dafny.IntOfUint32((_257_v106).Cardinality()))).Uint32()).(bool)
				_ = _rhs24
				var _rhs25 _dafny.Int = (func() _dafny.Int {
					if (_287_v118).Contains((_277_v110).F25()) {
						return (_287_v118).Get((_277_v110).F25()).(_dafny.Int)
					}
					return _82_v6
				})()
				_ = _rhs25
				var _rhs26 bool = (_dafny.IntOfInt64(674)).Cmp(Companion_Default___.Fm6(_78_v1, (_277_v110).F24(), true, _93_globalState)) == 0
				_ = _rhs26
				var _lhs13 *GlobalState = _93_globalState
				_ = _lhs13
				var _lhs14 _dafny.Array = _286_v117
				_ = _lhs14
				var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_286_v117), 0))
				_ = _lhs15
				var _lhs16 *GlobalState = _93_globalState
				_ = _lhs16
				_lhs13.F21 = _rhs24
				(_lhs14).ArraySet1(_rhs25, (_lhs15).Int())
				_lhs16.F21 = _rhs26
				var _288_v119 _dafny.Sequence
				_ = _288_v119
				var _289_v120 _dafny.Array
				_ = _289_v120
				var _290_v121 _dafny.Int
				_ = _290_v121
				var _291_v122 bool
				_ = _291_v122
				var _out20 _dafny.Sequence
				_ = _out20
				var _out21 _dafny.Array
				_ = _out21
				var _out22 _dafny.Int
				_ = _out22
				var _out23 bool
				_ = _out23
				_out20, _out21, _out22, _out23 = Companion_Default___.M0(_93_globalState)
				_288_v119 = _out20
				_289_v120 = _out21
				_290_v121 = _out22
				_291_v122 = _out23
				_78_v1 = ((_dafny.IntOfInt64(54)).Plus(_dafny.IntOfUint32((_83_v7).Cardinality()))).Cmp(_290_v121) == 0
				var _292_v123 *C0
				_ = _292_v123
				var _nw44 *C0 = New_C0_()
				_ = _nw44
				_nw44.Ctor__((_dafny.Zero).Minus(_82_v6), _291_v122, _288_v119)
				_292_v123 = _nw44
				var _293_v124 _dafny.Array
				_ = _293_v124
				var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
				_ = _nw45
				_293_v124 = _nw45
				var _294_v125 _dafny.Map
				_ = _294_v125
				_294_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v11, !((_277_v110).F24()))
				var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_293_v124), 0))
				_ = _index20
				(_293_v124).ArraySet1(_294_v125, (_index20).Int())
				var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_293_v124), 0))
				_ = _index21
				var _rhs27 *C0 = _292_v123
				_ = _rhs27
				var _rhs28 _dafny.Map = (func() _dafny.Map {
					var _coll6 = _dafny.NewMapBuilder()
					_ = _coll6
					for _iter9 := _dafny.Iterate((_88_v12).Elements()); ; {
						_compr_6, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _295_v126 _dafny.MultiSet
						_295_v126 = interface{}(_compr_6).(_dafny.MultiSet)
						if _dafny.Companion_Sequence_.Contains(_88_v12, _295_v126) {
							_coll6.Add(_295_v126, (_277_v110).F24())
						}
					}
					return _coll6.ToMap()
				}()).Merge((_294_v125).Merge(_294_v125))
				_ = _rhs28
				var _rhs29 _dafny.Array = _80_v3
				_ = _rhs29
				var _rhs30 _dafny.Int = (_286_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_286_v117), 0))).Int()).(_dafny.Int)
				_ = _rhs30
				var _lhs17 _dafny.Array = _293_v124
				_ = _lhs17
				var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_293_v124), 0))
				_ = _lhs18
				_292_v123 = _rhs27
				(_lhs17).ArraySet1(_rhs28, (_lhs18).Int())
				_80_v3 = _rhs29
				_290_v121 = _rhs30
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _296_v127 _dafny.Array
	_ = _296_v127
	var _nwElement0_8 _dafny.Sequence = (_277_v110).F25()
	_ = _nwElement0_8
	var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(16))
	_ = _nw46
	(_nw46).ArraySet1(_nwElement0_8, 0)
	(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_77_v0, _dafny.Companion_Sequence_.Update((_277_v110).F25(), (Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32(((_277_v110).F25()).Cardinality()))).Uint32(), _84_v8)), 1)
	(_nw46).ArraySet1((_277_v110).F25(), 2)
	(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vhh"), _77_v0), 3)
	(_nw46).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_277_v110).F25(), _77_v0), (Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_277_v110).F25(), _77_v0)).Cardinality()))).Uint32(), _84_v8), 4)
	(_nw46).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_277_v110).F25(), (_277_v110).F25()), 5)
	(_nw46).ArraySet1((_277_v110).F25(), 6)
	(_nw46).ArraySet1((_277_v110).F25(), 7)
	(_nw46).ArraySet1((_277_v110).F25(), 8)
	(_nw46).ArraySet1((_277_v110).F25(), 9)
	(_nw46).ArraySet1(Companion_Default___.Fm7((_277_v110).F24(), _82_v6, _82_v6, (_277_v110).F24(), _93_globalState), 10)
	(_nw46).ArraySet1(_77_v0, 11)
	(_nw46).ArraySet1(_77_v0, 12)
	(_nw46).ArraySet1(_77_v0, 13)
	(_nw46).ArraySet1((_277_v110).F25(), 14)
	(_nw46).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}((func(_297_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_298_i12 _dafny.Int) _dafny.CodePoint {
			return _297_v8
		}
	})(_84_v8))), 15)
	_296_v127 = _nw46
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))
	_ = _index22
	(_296_v127).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("chxreot"), (Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("chxreot")).Cardinality()))).Uint32(), _84_v8), (_index22).Int())
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))
	_ = _index23
	(_296_v127).ArraySet1((_277_v110).F25(), (_index23).Int())
	if _78_v1 {
		_80_v3 = _80_v3
		var _299_v128 _dafny.Sequence
		_ = _299_v128
		var _300_v129 _dafny.Array
		_ = _300_v129
		var _301_v130 _dafny.Int
		_ = _301_v130
		var _302_v131 bool
		_ = _302_v131
		var _out24 _dafny.Sequence
		_ = _out24
		var _out25 _dafny.Array
		_ = _out25
		var _out26 _dafny.Int
		_ = _out26
		var _out27 bool
		_ = _out27
		_out24, _out25, _out26, _out27 = Companion_Default___.M0(_93_globalState)
		_299_v128 = _out24
		_300_v129 = _out25
		_301_v130 = _out26
		_302_v131 = _out27
		(_93_globalState).F21 = !(!((_82_v6).Cmp(_dafny.IntOfInt64(991)) > 0)) || (_78_v1)
		var _303_v132 _dafny.Sequence
		_ = _303_v132
		var _304_v133 _dafny.Array
		_ = _304_v133
		var _305_v134 _dafny.Int
		_ = _305_v134
		var _306_v135 bool
		_ = _306_v135
		var _out28 _dafny.Sequence
		_ = _out28
		var _out29 _dafny.Array
		_ = _out29
		var _out30 _dafny.Int
		_ = _out30
		var _out31 bool
		_ = _out31
		_out28, _out29, _out30, _out31 = Companion_Default___.M0(_93_globalState)
		_303_v132 = _out28
		_304_v133 = _out29
		_305_v134 = _out30
		_306_v135 = _out31
		var _307_v136 T0
		_ = _307_v136
		var _nw47 *C0 = New_C0_()
		_ = _nw47
		_nw47.Ctor__(_dafny.IntOfUint32((_303_v132).Cardinality()), true, (_277_v110).F25())
		_307_v136 = _nw47
		var _308_v137 _dafny.Map
		_ = _308_v137
		_308_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _307_v136)
		var _309_v138 _dafny.MultiSet
		_ = _309_v138
		_309_v138 = _dafny.MultiSetOf((_308_v137).Update(_302_v131, _307_v136))
		var _310_v139 _dafny.Sequence
		_ = _310_v139
		_310_v139 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("dcoudebbi"), _303_v132)
		var _311_v140 _dafny.Sequence
		_ = _311_v140
		_311_v140 = _dafny.SeqOf((_310_v139).Select((Companion_Default___.SafeIndex(_305_v134, _dafny.IntOfUint32((_310_v139).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(524))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_312_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_313_i13 _dafny.Int) _dafny.CodePoint {
				return _312_v8
			}
		})(_84_v8))))
		var _314_v141 _dafny.Map
		_ = _314_v141
		_314_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v3, _dafny.UnicodeSeqOfUtf8Bytes("yckgtwu"))
		var _315_v142 _dafny.Array
		_ = _315_v142
		var _nwElement0_9 _dafny.Int = _82_v6
		_ = _nwElement0_9
		var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(17))
		_ = _nw48
		(_nw48).ArraySet1(_nwElement0_9, 0)
		(_nw48).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_77_v0, (_277_v110).F25())).Cardinality()), 1)
		(_nw48).ArraySet1((_305_v134).Plus(_305_v134), 2)
		(_nw48).ArraySet1((_309_v138).Cardinality(), 3)
		(_nw48).ArraySet1(_305_v134, 4)
		(_nw48).ArraySet1(_dafny.IntOfUint32(((_277_v110).F25()).Cardinality()), 5)
		(_nw48).ArraySet1((func() _dafny.Int {
			if _78_v1 {
				return _dafny.IntOfUint32(((_310_v139).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_310_v139).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
			}
			return _dafny.IntOfInt64(995)
		})(), 6)
		(_nw48).ArraySet1(Companion_Default___.SafeDivisionInt(_301_v130, _301_v130), 7)
		(_nw48).ArraySet1((_dafny.IntOfInt64(782)).Plus(_301_v130), 8)
		(_nw48).ArraySet1((_82_v6).Minus(_301_v130), 9)
		(_nw48).ArraySet1(_dafny.IntOfInt64(929), 10)
		(_nw48).ArraySet1(_dafny.IntOfUint32(((_277_v110).F25()).Cardinality()), 11)
		(_nw48).ArraySet1(_dafny.IntOfInt64(507), 12)
		(_nw48).ArraySet1(Companion_Default___.SafeDivisionInt((_314_v141).Cardinality(), _305_v134), 13)
		(_nw48).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(555))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_316_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_317_i14 _dafny.Int) _dafny.CodePoint {
				return _316_v8
			}
		})(_84_v8)))).Cardinality()), _305_v134), 14)
		(_nw48).ArraySet1(Companion_Default___.Fm6((_277_v110).F24(), false, _306_v135, _93_globalState), 15)
		(_nw48).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_83_v7, _dafny.SeqOf(_301_v130))).Cardinality()), 16)
		_315_v142 = _nw48
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_315_v142), 0))
		_ = _index24
		(_315_v142).ArraySet1(_305_v134, (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_315_v142), 0))
		_ = _index25
		(_315_v142).ArraySet1((_82_v6).Times(_82_v6), (_index25).Int())
	} else {
		var _318_v143 D2
		_ = _318_v143
		_318_v143 = Companion_D2_.Create_DC9_(_78_v1, _82_v6)
		var _319_v144 *C0
		_ = _319_v144
		var _nw49 *C0 = New_C0_()
		_ = _nw49
		_nw49.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v4, (_277_v110).F24())).Cardinality(), (_277_v110).F24(), (_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence))
		_319_v144 = _nw49
		var _320_v145 _dafny.Set
		_ = _320_v145
		_320_v145 = _dafny.SetOf(_319_v144)
		var _321_v146 _dafny.Map
		_ = _321_v146
		_321_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _320_v145)
		var _322_v147 _dafny.Map
		_ = _322_v147
		_322_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_318_v143, (_321_v146).Cardinality()))
		var _323_v149 _dafny.Map
		_ = _323_v149
		_323_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F24(), _dafny.IntOfInt64(428)))
		_322_v147 = func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter10 := _dafny.Iterate((_323_v149).Keys().Elements()); ; {
				_compr_7, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _324_v148 _dafny.Sequence
				_324_v148 = interface{}(_compr_7).(_dafny.Sequence)
				if (_323_v149).Contains(_324_v148) {
					_coll7.Add(_324_v148, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_318_v143, _dafny.IntOfUint32(((_277_v110).F25()).Cardinality())))
				}
			}
			return _coll7.ToMap()
		}()
		var _325_v150 _dafny.Sequence
		_ = _325_v150
		var _326_v151 _dafny.Array
		_ = _326_v151
		var _327_v152 _dafny.Int
		_ = _327_v152
		var _328_v153 bool
		_ = _328_v153
		var _out32 _dafny.Sequence
		_ = _out32
		var _out33 _dafny.Array
		_ = _out33
		var _out34 _dafny.Int
		_ = _out34
		var _out35 bool
		_ = _out35
		_out32, _out33, _out34, _out35 = Companion_Default___.M0(_93_globalState)
		_325_v150 = _out32
		_326_v151 = _out33
		_327_v152 = _out34
		_328_v153 = _out35
		var _329_v154 _dafny.Map
		_ = _329_v154
		_329_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v152, _84_v8)
		var _330_v155 D0
		_ = _330_v155
		_330_v155 = Companion_D0_.Create_DC1_(!(_78_v1), (_277_v110).F24(), _327_v152)
		var _331_v156 _dafny.Sequence
		_ = _331_v156
		_331_v156 = _dafny.SeqOf(Companion_D1_.Create_DC6_((_330_v155).Dtor_cf3(), _277_v110, true, (_277_v110).F24()))
		_329_v154 = (_329_v154).Update(_dafny.IntOfUint32((_331_v156).Cardinality()), Companion_Default___.Fm4((_329_v154).Cardinality(), (_319_v144).F26(), _dafny.UnicodeSeqOfUtf8Bytes("g"), _257_v106, _93_globalState))
		if (_277_v110).F24() {
			var _332_v157 _dafny.Sequence
			_ = _332_v157
			var _333_v158 _dafny.Array
			_ = _333_v158
			var _334_v159 _dafny.Int
			_ = _334_v159
			var _335_v160 bool
			_ = _335_v160
			var _out36 _dafny.Sequence
			_ = _out36
			var _out37 _dafny.Array
			_ = _out37
			var _out38 _dafny.Int
			_ = _out38
			var _out39 bool
			_ = _out39
			_out36, _out37, _out38, _out39 = Companion_Default___.M0(_93_globalState)
			_332_v157 = _out36
			_333_v158 = _out37
			_334_v159 = _out38
			_335_v160 = _out39
			var _336_v161 *C0
			_ = _336_v161
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__(_334_v159, _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("v"), (_277_v110).F25()), _dafny.Companion_Sequence_.Concatenate(_332_v157, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(728))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_337_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_338_i15 _dafny.Int) _dafny.CodePoint {
					return _337_v8
				}
			})(_84_v8)))))
			_336_v161 = _nw50
			var _339_v162 D0
			_ = _339_v162
			_339_v162 = Companion_D0_.Create_DC1_(_78_v1, (_257_v106).Select((Companion_Default___.SafeIndex(_327_v152, _dafny.IntOfUint32((_257_v106).Cardinality()))).Uint32()).(bool), (_279_v112).Cardinality())
			var _340_v163 D0
			_ = _340_v163
			_340_v163 = Companion_D0_.Create_DC3_(_339_v162)
			var _341_v164 D0
			_ = _341_v164
			_341_v164 = Companion_D0_.Create_DC3_(_339_v162)
			var _342_v165 D0
			_ = _342_v165
			_342_v165 = Companion_D0_.Create_DC3_(_339_v162)
			var _pat_let_tv12 = _340_v163
			_ = _pat_let_tv12
			var _rhs31 bool = _335_v160
			_ = _rhs31
			var _rhs32 bool = (_277_v110).F24()
			_ = _rhs32
			var _rhs33 _dafny.Int = (func() _dafny.Int {
				if !(_78_v1) {
					return _327_v152
				}
				return _334_v159
			})()
			_ = _rhs33
			var _rhs34 _dafny.Int = _dafny.IntOfInt64(-766)
			_ = _rhs34
			var _rhs35 D0 = func(_pat_let12_0 D0) D0 {
				return func(_343_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let13_0 D0) D0 {
						return func(_344_dt__update_hcf6_h0 D0) D0 {
							return Companion_D0_.Create_DC3_(_344_dt__update_hcf6_h0)
						}(_pat_let13_0)
					}(Companion_D0_.Create_DC3_(_pat_let_tv12))
				}(_pat_let12_0)
			}(_342_v165)
			_ = _rhs35
			var _lhs19 *GlobalState = _93_globalState
			_ = _lhs19
			var _lhs20 *GlobalState = _93_globalState
			_ = _lhs20
			var _lhs21 *GlobalState = _93_globalState
			_ = _lhs21
			_335_v160 = _rhs31
			_lhs19.F20 = _rhs32
			_lhs20.F5 = _rhs33
			_lhs21.F5 = _rhs34
			_342_v165 = _rhs35
			(_93_globalState).F20 = _78_v1
			var _345_v166 D3
			_ = _345_v166
			_345_v166 = Companion_D3_.Create_DC11_(_83_v7)
			_335_v160 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate((_345_v166).Dtor_cf19(), _83_v7), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(796))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_346_v161 *C0, _347_v110 T1) func(_dafny.Int) _dafny.Int {
				return func(_348_i16 _dafny.Int) _dafny.Int {
					return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_346_v161).F26(), (_347_v110).F24()), func() _dafny.Map {
						var _coll8 = _dafny.NewMapBuilder()
						_ = _coll8
						for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(636), _dafny.IntOfInt64(445))); ; {
							_compr_8, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _349_v167 _dafny.Int
							_349_v167 = interface{}(_compr_8).(_dafny.Int)
							if ((_dafny.IntOfInt64(636)).Cmp(_349_v167) <= 0) && ((_349_v167).Cmp(_dafny.IntOfInt64(445)) < 0) {
								_coll8.Add((_349_v167).Times((_346_v161).F26()), (_347_v110).F24())
							}
						}
						return _coll8.ToMap()
					}())).Cardinality()
				}
			})(_336_v161, _277_v110))))
		} else {
			var _350_v168 _dafny.Sequence
			_ = _350_v168
			var _351_v169 _dafny.Array
			_ = _351_v169
			var _352_v170 _dafny.Int
			_ = _352_v170
			var _353_v171 bool
			_ = _353_v171
			var _out40 _dafny.Sequence
			_ = _out40
			var _out41 _dafny.Array
			_ = _out41
			var _out42 _dafny.Int
			_ = _out42
			var _out43 bool
			_ = _out43
			_out40, _out41, _out42, _out43 = Companion_Default___.M0(_93_globalState)
			_350_v168 = _out40
			_351_v169 = _out41
			_352_v170 = _out42
			_353_v171 = _out43
			var _354_v172 _dafny.Sequence
			_ = _354_v172
			var _355_v173 _dafny.Array
			_ = _355_v173
			var _356_v174 _dafny.Int
			_ = _356_v174
			var _357_v175 bool
			_ = _357_v175
			var _out44 _dafny.Sequence
			_ = _out44
			var _out45 _dafny.Array
			_ = _out45
			var _out46 _dafny.Int
			_ = _out46
			var _out47 bool
			_ = _out47
			_out44, _out45, _out46, _out47 = Companion_Default___.M0(_93_globalState)
			_354_v172 = _out44
			_355_v173 = _out45
			_356_v174 = _out46
			_357_v175 = _out47
			var _358_v176 _dafny.Sequence
			_ = _358_v176
			var _359_v177 _dafny.Array
			_ = _359_v177
			var _360_v178 _dafny.Int
			_ = _360_v178
			var _361_v179 bool
			_ = _361_v179
			var _out48 _dafny.Sequence
			_ = _out48
			var _out49 _dafny.Array
			_ = _out49
			var _out50 _dafny.Int
			_ = _out50
			var _out51 bool
			_ = _out51
			_out48, _out49, _out50, _out51 = Companion_Default___.M0(_93_globalState)
			_358_v176 = _out48
			_359_v177 = _out49
			_360_v178 = _out50
			_361_v179 = _out51
			(_93_globalState).F20 = (_318_v143).Dtor_cf16()
			(_93_globalState).F21 = !((_277_v110).F24())
		}
		var _362_v180 T0
		_ = _362_v180
		var _nw51 *C0 = New_C0_()
		_ = _nw51
		_nw51.Ctor__(_327_v152, (_278_v111).Dtor_cf13(), _dafny.UnicodeSeqOfUtf8Bytes("iwdsggla"))
		_362_v180 = _nw51
		_362_v180 = _362_v180
	}
	var _363_v181 _dafny.Map
	_ = _363_v181
	_363_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v8, _82_v6)
	var _364_v182 _dafny.MultiSet
	_ = _364_v182
	_364_v182 = _dafny.MultiSetOf((_277_v110).F24())
	var _365_v183 _dafny.Sequence
	_ = _365_v183
	_365_v183 = _dafny.SeqOf(_364_v182)
	var _366_v185 _dafny.Array
	_ = _366_v185
	var _nwElement0_10 _dafny.Int = _82_v6
	_ = _nwElement0_10
	var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(18))
	_ = _nw52
	(_nw52).ArraySet1(_nwElement0_10, 0)
	(_nw52).ArraySet1((_363_v181).Cardinality(), 1)
	(_nw52).ArraySet1(_dafny.IntOfInt64(37), 2)
	(_nw52).ArraySet1(_82_v6, 3)
	(_nw52).ArraySet1(_82_v6, 4)
	(_nw52).ArraySet1(((func() _dafny.Int {
		if (_258_v107).Contains(_dafny.IntOfUint32(((_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence)).Cardinality())) {
			return (_258_v107).Get(_dafny.IntOfUint32(((_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence)).Cardinality())).(_dafny.Int)
		}
		return _82_v6
	})()).Times(_82_v6), 5)
	(_nw52).ArraySet1(_82_v6, 6)
	(_nw52).ArraySet1(Companion_Default___.SafeDivisionInt(_82_v6, _dafny.IntOfInt64(655)), 7)
	(_nw52).ArraySet1(_82_v6, 8)
	(_nw52).ArraySet1(Companion_Default___.SafeModuloInt(_82_v6, _82_v6), 9)
	(_nw52).ArraySet1(_82_v6, 10)
	(_nw52).ArraySet1(_dafny.IntOfUint32(((_277_v110).F25()).Cardinality()), 11)
	(_nw52).ArraySet1(_dafny.IntOfInt64(201), 12)
	(_nw52).ArraySet1((func() _dafny.Int {
		if (_277_v110).F24() {
			return _82_v6
		}
		return (func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter12 := _dafny.Iterate((_365_v183).Elements()); ; {
				_compr_9, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _367_v184 _dafny.MultiSet
				_367_v184 = interface{}(_compr_9).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_365_v183, _367_v184) {
					_coll9.Add(_367_v184)
				}
			}
			return _coll9.ToSet()
		}()).Cardinality()
	})(), 13)
	(_nw52).ArraySet1(_dafny.IntOfUint32((_83_v7).Cardinality()), 14)
	(_nw52).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
		if _78_v1 {
			return _83_v7
		}
		return _83_v7
	})()).Cardinality()), 15)
	(_nw52).ArraySet1(_82_v6, 16)
	(_nw52).ArraySet1(_82_v6, 17)
	_366_v185 = _nw52
	var _rhs36 _dafny.Int = _82_v6
	_ = _rhs36
	var _rhs37 bool = (_277_v110).F24()
	_ = _rhs37
	var _rhs38 _dafny.Int = (_82_v6).Minus((func() _dafny.Int {
		if (_364_v182).Contains(_78_v1) {
			return (_364_v182).Multiplicity(_78_v1)
		}
		return _82_v6
	})())
	_ = _rhs38
	var _rhs39 _dafny.Array = _366_v185
	_ = _rhs39
	var _lhs22 *GlobalState = _93_globalState
	_ = _lhs22
	_82_v6 = _rhs36
	_78_v1 = _rhs37
	_lhs22.F5 = _rhs38
	_366_v185 = _rhs39
	var _368_i17 _dafny.Int
	_ = _368_i17
	_368_i17 = _dafny.Zero
	{
		for _78_v1 {
			{
				if (_368_i17).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_368_i17 = (_368_i17).Plus(_dafny.One)
				(_93_globalState).F21 = !(((_277_v110).F24()) || ((_277_v110).F24()))
				(_93_globalState).F5 = _82_v6
				(_93_globalState).F21 = (_277_v110).F24()
				(_93_globalState).F5 = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_82_v6, _82_v6, _dafny.IntOfInt64(529), _82_v6, _82_v6)).Cardinality()))).Plus(_82_v6)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _369_v186 *C0
	_ = _369_v186
	var _nw53 *C0 = New_C0_()
	_ = _nw53
	_nw53.Ctor__(_82_v6, _78_v1, (_277_v110).F25())
	_369_v186 = _nw53
	var _370_v187 _dafny.Sequence
	_ = _370_v187
	_370_v187 = _dafny.SeqOf(_369_v186)
	_82_v6 = (Companion_Default___.SafeModuloInt(_82_v6, _82_v6)).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm7((_277_v110).F24(), _82_v6, _dafny.IntOfInt64(25), _78_v1, _93_globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_370_v187).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm7((_277_v110).F24(), _82_v6, _dafny.IntOfInt64(25), _78_v1, _93_globalState)).Cardinality()))).Uint32(), _84_v8)).Cardinality())))
	if _78_v1 {
		(_93_globalState).F5 = (_dafny.IntOfUint32(((_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence)).Cardinality())).Times(((func() _dafny.MultiSet {
			if _78_v1 {
				return _dafny.MultiSetOf((_83_v7).Select((Companion_Default___.SafeIndex((_369_v186).F26(), _dafny.IntOfUint32((_83_v7).Cardinality()))).Uint32()).(_dafny.Int), (_369_v186).F26(), (_369_v186).F26())
			}
			return _87_v11
		})()).Cardinality())
		var _371_v188 _dafny.Map
		_ = _371_v188
		_371_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_366_v185, _369_v186)
		_371_v188 = (_371_v188).Merge(_371_v188)
		var _372_v189 _dafny.Map
		_ = _372_v189
		_372_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F24(), _279_v112)
		var _373_v190 _dafny.Set
		_ = _373_v190
		_373_v190 = _dafny.SetOf((_369_v186).F26())
		var _rhs40 _dafny.Int = _dafny.IntOfUint32(((_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence)).Cardinality())
		_ = _rhs40
		var _rhs41 _dafny.Map = (func() _dafny.Map {
			if (_372_v189).Contains((_79_v2).IsProperSubsetOf(_79_v2)) {
				return (_372_v189).Get((_79_v2).IsProperSubsetOf(_79_v2)).(_dafny.Map)
			}
			return _279_v112
		})()
		_ = _rhs41
		var _rhs42 _dafny.Int = Companion_Default___.Fm6((_277_v110).F24(), (_277_v110).F24(), !((_277_v110).F24()), _93_globalState)
		_ = _rhs42
		var _rhs43 _dafny.Int = ((_369_v186).F26()).Minus(((_373_v190).Cardinality()).Minus(_82_v6))
		_ = _rhs43
		var _lhs23 *GlobalState = _93_globalState
		_ = _lhs23
		var _lhs24 *GlobalState = _93_globalState
		_ = _lhs24
		var _lhs25 *GlobalState = _93_globalState
		_ = _lhs25
		_lhs23.F5 = _rhs40
		_279_v112 = _rhs41
		_lhs24.F5 = _rhs42
		_lhs25.F5 = _rhs43
		(_93_globalState).F5 = (func() _dafny.Int {
			if (func() bool {
				if (_277_v110).F24() {
					return (_277_v110).F24()
				}
				return (_277_v110).F24()
			})() {
				return _82_v6
			}
			return (_369_v186).F26()
		})()
		var _374_v191 _dafny.Sequence
		_ = _374_v191
		_374_v191 = _dafny.SeqOf(_277_v110, _277_v110, _277_v110, _277_v110, _277_v110)
		var _375_v192 _dafny.Set
		_ = _375_v192
		_375_v192 = _dafny.SetOf(_79_v2)
		var _376_v193 T0
		_ = _376_v193
		var _nw54 *C0 = New_C0_()
		_ = _nw54
		_nw54.Ctor__(_dafny.IntOfUint32((_257_v106).Cardinality()), _78_v1, _dafny.UnicodeSeqOfUtf8Bytes("bacmjqi"))
		_376_v193 = _nw54
		var _377_v194 _dafny.Sequence
		_ = _377_v194
		_377_v194 = _dafny.SeqOf(_376_v193)
		var _378_v195 _dafny.Map
		_ = _378_v195
		_378_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_375_v192).Cardinality()), (_377_v194).Select((Companion_Default___.SafeIndex((_369_v186).F26(), _dafny.IntOfUint32((_377_v194).Cardinality()))).Uint32()).(T0))
		var _rhs44 _dafny.Int = Companion_Default___.SafeDivisionInt((_369_v186).F26(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_277_v110), _374_v191)).Cardinality()))
		_ = _rhs44
		var _rhs45 _dafny.Int = Companion_Default___.SafeDivisionInt(((_369_v186).F26()).Minus((_369_v186).F26()), ((_378_v195).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v6, _376_v193))).Cardinality())
		_ = _rhs45
		var _rhs46 bool = (_277_v110).F24()
		_ = _rhs46
		var _lhs26 *GlobalState = _93_globalState
		_ = _lhs26
		var _lhs27 *GlobalState = _93_globalState
		_ = _lhs27
		var _lhs28 *GlobalState = _93_globalState
		_ = _lhs28
		_lhs26.F5 = _rhs44
		_lhs27.F5 = _rhs45
		_lhs28.F21 = _rhs46
	} else {
		var _379_v196 _dafny.Map
		_ = _379_v196
		_379_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v0, _78_v1)
		(_93_globalState).F20 = (func() bool {
			if (_379_v196).Contains(_dafny.UnicodeSeqOfUtf8Bytes("nxd")) {
				return (_379_v196).Get(_dafny.UnicodeSeqOfUtf8Bytes("nxd")).(bool)
			}
			return _78_v1
		})()
		var _380_v197 D3
		_ = _380_v197
		_380_v197 = Companion_D3_.Create_DC11_(_83_v7)
		var _381_v198 D3
		_ = _381_v198
		_381_v198 = Companion_D3_.Create_DC11_((_380_v197).Dtor_cf19())
		var _source4 D3 = _380_v197
		_ = _source4
		if _source4.Is_DC12() {
			var _382___mcc_h29 bool = _source4.Get_().(D3_DC12).Cf20
			_ = _382___mcc_h29
			var _383___mcc_h30 _dafny.Int = _source4.Get_().(D3_DC12).Cf21
			_ = _383___mcc_h30
			var _384_cf21 _dafny.Int = _383___mcc_h30
			_ = _384_cf21
			var _385_cf20 bool = _382___mcc_h29
			_ = _385_cf20
			(_93_globalState).F5 = _82_v6
			_82_v6 = _dafny.IntOfUint32((_92_v14).Cardinality())
			_80_v3 = (func() _dafny.Array {
				if _78_v1 {
					return _80_v3
				}
				return _80_v3
			})()
			_366_v185 = _366_v185
		} else if _source4.Is_DC13() {
			var _386___mcc_h31 *C0 = _source4.Get_().(D3_DC13).Cf22
			_ = _386___mcc_h31
			var _387___mcc_h32 _dafny.Int = _source4.Get_().(D3_DC13).Cf23
			_ = _387___mcc_h32
			var _388_cf23 _dafny.Int = _387___mcc_h32
			_ = _388_cf23
			var _389_cf22 *C0 = _386___mcc_h31
			_ = _389_cf22
			var _390_v199 _dafny.Array
			_ = _390_v199
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
			_ = _nw55
			_390_v199 = _nw55
			var _391_v200 D4
			_ = _391_v200
			_391_v200 = Companion_D4_.Create_DC15_(_257_v106)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_390_v199), 0))
			_ = _index26
			(_390_v199).ArraySet1((_391_v200).Dtor_cf25(), (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_390_v199), 0))
			_ = _index27
			(_390_v199).ArraySet1(_257_v106, (_index27).Int())
			var _392_v201 _dafny.Sequence
			_ = _392_v201
			var _393_v202 _dafny.Array
			_ = _393_v202
			var _394_v203 _dafny.Int
			_ = _394_v203
			var _395_v204 bool
			_ = _395_v204
			var _out52 _dafny.Sequence
			_ = _out52
			var _out53 _dafny.Array
			_ = _out53
			var _out54 _dafny.Int
			_ = _out54
			var _out55 bool
			_ = _out55
			_out52, _out53, _out54, _out55 = Companion_Default___.M0(_93_globalState)
			_392_v201 = _out52
			_393_v202 = _out53
			_394_v203 = _out54
			_395_v204 = _out55
			(_93_globalState).F5 = _dafny.IntOfUint32(((_277_v110).F25()).Cardinality())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_80_v3), 0))
			_ = _index28
			(_80_v3).ArraySet1((_277_v110).F24(), (_index28).Int())
			var _396_v205 _dafny.Map
			_ = _396_v205
			_396_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F24(), _84_v8)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_80_v3), 0))
			_ = _index29
			(_80_v3).ArraySet1((_389_cf22).Fm1(_396_v205, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_394_v203, (_277_v110).F24()), _dafny.Companion_Sequence_.Concatenate(_83_v7, _83_v7), _93_globalState), (_index29).Int())
		} else if _source4.Is_DC11() {
			var _397___mcc_h33 _dafny.Sequence = _source4.Get_().(D3_DC11).Cf19
			_ = _397___mcc_h33
			var _398_cf19 _dafny.Sequence = _397___mcc_h33
			_ = _398_cf19
			var _399_v206 _dafny.Map
			_ = _399_v206
			_399_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F24(), (_369_v186).F26())
			(_93_globalState).F5 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F24(), (_369_v186).F26())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F24(), _82_v6)).Merge(_399_v206))).Cardinality()
			var _400_v207 _dafny.Sequence
			_ = _400_v207
			var _401_v208 _dafny.Array
			_ = _401_v208
			var _402_v209 _dafny.Int
			_ = _402_v209
			var _403_v210 bool
			_ = _403_v210
			var _out56 _dafny.Sequence
			_ = _out56
			var _out57 _dafny.Array
			_ = _out57
			var _out58 _dafny.Int
			_ = _out58
			var _out59 bool
			_ = _out59
			_out56, _out57, _out58, _out59 = Companion_Default___.M0(_93_globalState)
			_400_v207 = _out56
			_401_v208 = _out57
			_402_v209 = _out58
			_403_v210 = _out59
			var _404_v211 _dafny.Sequence
			_ = _404_v211
			var _405_v212 _dafny.Array
			_ = _405_v212
			var _406_v213 _dafny.Int
			_ = _406_v213
			var _407_v214 bool
			_ = _407_v214
			var _out60 _dafny.Sequence
			_ = _out60
			var _out61 _dafny.Array
			_ = _out61
			var _out62 _dafny.Int
			_ = _out62
			var _out63 bool
			_ = _out63
			_out60, _out61, _out62, _out63 = Companion_Default___.M0(_93_globalState)
			_404_v211 = _out60
			_405_v212 = _out61
			_406_v213 = _out62
			_407_v214 = _out63
			var _408_v215 _dafny.Sequence
			_ = _408_v215
			_408_v215 = _dafny.SeqOf(_83_v7)
			var _409_v216 _dafny.Map
			_ = _409_v216
			_409_v216 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v110).F24(), _84_v8)
			var _410_v217 _dafny.Map
			_ = _410_v217
			_410_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v6, _83_v7)
			(_93_globalState).F5 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_408_v215).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.IntOfUint32((_408_v215).Cardinality()))).Uint32()).(_dafny.Sequence), (_369_v186).Fm1((_409_v216).Update(_403_v210, _84_v8), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_369_v186).F26(), _403_v210), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_410_v217).Contains((_79_v2).Cardinality()) {
					return (_410_v217).Get((_79_v2).Cardinality()).(_dafny.Sequence)
				}
				return _398_cf19
			})(), (Companion_Default___.SafeIndex((_369_v186).F26(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_410_v217).Contains((_79_v2).Cardinality()) {
					return (_410_v217).Get((_79_v2).Cardinality()).(_dafny.Sequence)
				}
				return _398_cf19
			})()).Cardinality()))).Uint32(), Companion_Default___.Fm6(_78_v1, !(_403_v210), (_277_v110).F24(), _93_globalState)), _93_globalState))).Cardinality()
		} else {
			var _411___mcc_h34 D3 = _source4.Get_().(D3_DC14).Cf24
			_ = _411___mcc_h34
			var _412_cf24 D3 = _411___mcc_h34
			_ = _412_cf24
			var _413_v218 _dafny.Map
			_ = _413_v218
			_413_v218 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _84_v8)
			var _414_v219 _dafny.Map
			_ = _414_v219
			_414_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v6, (_277_v110).F24())
			var _415_v220 T0
			_ = _415_v220
			var _nw56 *C0 = New_C0_()
			_ = _nw56
			_nw56.Ctor__(_82_v6, (_277_v110).Fm1(_413_v218, _414_v219, _dafny.Companion_Sequence_.Update(_83_v7, (Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_83_v7).Cardinality()))).Uint32(), (_369_v186).F26()), _93_globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(940))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_416_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_417_i18 _dafny.Int) _dafny.CodePoint {
					return _416_v8
				}
			})(_84_v8))))
			_415_v220 = _nw56
			var _418_v221 D4
			_ = _418_v221
			_418_v221 = Companion_D4_.Create_DC16_((_369_v186).F26(), (_257_v106).Select((Companion_Default___.SafeIndex(_82_v6, _dafny.IntOfUint32((_257_v106).Cardinality()))).Uint32()).(bool), _415_v220)
			var _419_v222 _dafny.Map
			_ = _419_v222
			_419_v222 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_418_v221, ((_279_v112).Cardinality()).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_277_v110).F25()).Cardinality()))))
			_419_v222 = (_419_v222).Update(_418_v221, _dafny.IntOfUint32((_257_v106).Cardinality()))
			var _nw57 *C0 = New_C0_()
			_ = _nw57
			_nw57.Ctor__(((_369_v186).F26()).Minus((_369_v186).Fm3((_369_v186).F26(), _dafny.SeqOf(_dafny.SeqOf((func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter13 := _dafny.Iterate((_83_v7).Elements()); ; {
					_compr_10, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _420_v223 _dafny.Int
					_420_v223 = interface{}(_compr_10).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_83_v7, _420_v223) {
						_coll10.Add(Companion_Default___.SafeDivisionInt(_420_v223, (_369_v186).F26()), (_277_v110).F24())
					}
				}
				return _coll10.ToMap()
			}()).Cardinality())), _dafny.IntOfInt64(638), (_277_v110).F25(), _93_globalState)), (func() bool {
				if (_277_v110).F24() {
					return _78_v1
				}
				return false
			})(), (_277_v110).F25())
			_369_v186 = _nw57
			var _421_v224 _dafny.Set
			_ = _421_v224
			_421_v224 = _dafny.SetOf(_380_v197)
			(_93_globalState).F5 = ((_279_v112).Cardinality()).Times((_421_v224).Cardinality())
			var _422_v225 _dafny.Map
			_ = _422_v225
			_422_v225 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_v107, (_277_v110).F24())
			var _423_v226 D5
			_ = _423_v226
			_423_v226 = Companion_D5_.Create_DC20_(_422_v225)
			var _rhs47 bool = ((_423_v226).Dtor_cf32()).Contains(_258_v107)
			_ = _rhs47
			var _rhs48 _dafny.Int = (_369_v186).F26()
			_ = _rhs48
			var _lhs29 *GlobalState = _93_globalState
			_ = _lhs29
			var _lhs30 *GlobalState = _93_globalState
			_ = _lhs30
			_lhs29.F20 = _rhs47
			_lhs30.F5 = _rhs48
		}
		var _424_v227 _dafny.Map
		_ = _424_v227
		_424_v227 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v3, (_369_v186).F26())
		var _425_v228 _dafny.Map
		_ = _425_v228
		_425_v228 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_277_v110).F24() {
				return _82_v6
			}
			return (_424_v227).Cardinality()
		})(), (_82_v6).Cmp(_dafny.IntOfUint32((_92_v14).Cardinality())) != 0)
		_425_v228 = (_425_v228).Update(_dafny.IntOfInt64(488), _78_v1)
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_80_v3), 0))
		_ = _index30
		(_80_v3).ArraySet1((_277_v110).F24(), (_index30).Int())
		var _426_v229 T0
		_ = _426_v229
		var _nw58 *C0 = New_C0_()
		_ = _nw58
		_nw58.Ctor__((_369_v186).F26(), _78_v1, (_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence))
		_426_v229 = _nw58
		var _427_v230 _dafny.Sequence
		_ = _427_v230
		_427_v230 = _dafny.SeqOf(_426_v229, _426_v229, _426_v229)
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_80_v3), 0))
		_ = _index31
		var _rhs49 _dafny.Int = (_369_v186).F26()
		_ = _rhs49
		var _rhs50 bool = (_dafny.IntOfInt64(-804)).Cmp((_369_v186).F26()) <= 0
		_ = _rhs50
		var _rhs51 bool = !_dafny.Companion_Sequence_.Equal(_83_v7, _dafny.Companion_Sequence_.Update(_83_v7, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_427_v230).Cardinality()), _dafny.IntOfUint32((_83_v7).Cardinality()))).Uint32(), _82_v6))
		_ = _rhs51
		var _rhs52 bool = (_257_v106).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.IntOfUint32((_257_v106).Cardinality()))).Uint32()).(bool)
		_ = _rhs52
		var _lhs31 _dafny.Array = _80_v3
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_80_v3), 0))
		_ = _lhs32
		_82_v6 = _rhs49
		_78_v1 = _rhs50
		(_lhs31).ArraySet1(_rhs51, (_lhs32).Int())
		_78_v1 = _rhs52
		(_93_globalState).F5 = (_369_v186).F26()
	}
	var _428_v231 _dafny.Array
	_ = _428_v231
	var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
	_ = _nw59
	_428_v231 = _nw59
	var _429_v232 _dafny.Map
	_ = _429_v232
	_429_v232 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_428_v231, _dafny.MultiSetFromSeq(_dafny.SeqOf(false, (_277_v110).F24())))
	_429_v232 = (_429_v232).Update(_428_v231, _364_v182)
	var _430_v233 *C0
	_ = _430_v233
	var _nw60 *C0 = New_C0_()
	_ = _nw60
	_nw60.Ctor__(_82_v6, (_277_v110).F24(), (_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence))
	_430_v233 = _nw60
	var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_80_v3), 0))
	_ = _index32
	(_80_v3).ArraySet1((_277_v110).F24(), (_index32).Int())
	var _431_v234 _dafny.Set
	_ = _431_v234
	_431_v234 = _dafny.SetOf(_80_v3, _80_v3, _80_v3)
	var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_80_v3), 0))
	_ = _index33
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))
	_ = _index34
	var _rhs53 bool = (_431_v234).IsDisjointFrom(_431_v234)
	_ = _rhs53
	var _rhs54 _dafny.Sequence = _77_v0
	_ = _rhs54
	var _rhs55 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_296_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))).Int()).(_dafny.Sequence), _77_v0)
	_ = _rhs55
	var _lhs33 _dafny.Array = _80_v3
	_ = _lhs33
	var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_80_v3), 0))
	_ = _lhs34
	var _lhs35 _dafny.Array = _296_v127
	_ = _lhs35
	var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_296_v127), 0))
	_ = _lhs36
	var _lhs37 *GlobalState = _93_globalState
	_ = _lhs37
	(_lhs33).ArraySet1(_rhs53, (_lhs34).Int())
	(_lhs35).ArraySet1(_rhs54, (_lhs36).Int())
	_lhs37.F19 = _rhs55
	_dafny.Print(_77_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_78_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v2).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v3).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v4).Dtor_cf0().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_82_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_83_v7, _dafny.SeqOf(_dafny.IntOfInt64(462), _dafny.IntOfInt64(462), _dafny.IntOfInt64(462))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_84_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(22)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(23)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(24)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(25)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v9).ArrayGet1CodePoint((_dafny.IntOfInt64(26)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_86_v10).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_87_v11).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(462))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_88_v12, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(462)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_89_v13).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_92_v14).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F0().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState.F1).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_93_globalState.F4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_93_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F6()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(213444), _dafny.CodePoint('j')).UpdateUnsafe(_dafny.IntOfInt64(668), _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_93_globalState.F8).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState.F14).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(462))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_93_globalState).F15()).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_93_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_93_globalState.F19.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_93_globalState.F20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_93_globalState.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v105).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_257_v106, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v107).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(49), _dafny.IntOfInt64(462))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v108).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v108).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v110).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v110).F25().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_278_v111).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_278_v111).Dtor_cf11()).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_278_v111).Dtor_cf11()).F25().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_278_v111).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_278_v111).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v112).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_283_v115).Equals(_dafny.SetOf(_dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v116).Equals(_dafny.SetOf(_dafny.SetOf(_dafny.CodePoint('x')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_285_i11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v127).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_296_v127).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_363_v181).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(213443))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v182).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_365_v183, _dafny.SeqOf(_dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v185).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_368_i17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_369_v186).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_370_v187).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_429_v232).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_430_v233).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_431_v234).Cardinality())
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
	Cf1 bool
	Cf2 bool
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 bool, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 _dafny.Array
	Cf5 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 _dafny.Array, Cf5 _dafny.Int) D0 {
	return D0{D0_DC2{Cf4, Cf5}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Sequence) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC3 struct {
	Cf6 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf6 D0) D0 {
	return D0{D0_DC3{Cf6}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf6() D0 {
	return _this.Get_().(D0_DC3).Cf6
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + data.Cf0.VerbatimString(true) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf8 bool
	Cf9 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf8 bool, Cf9 _dafny.Int) D1 {
	return D1{D1_DC5{Cf8, Cf9}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf10 _dafny.Int
	Cf11 T1
	Cf12 bool
	Cf13 bool
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf10 _dafny.Int, Cf11 T1, Cf12 bool, Cf13 bool) D1 {
	return D1{D1_DC6{Cf10, Cf11, Cf12, Cf13}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC4 struct {
	Cf7 *C0
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 *C0) D1 {
	return D1{D1_DC4{Cf7}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC7 struct {
	Cf14 D1
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_(Cf14 D1) D1 {
	return D1{D1_DC7{Cf14}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false, _dafny.Zero)
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf10
}

func (_this D1) Dtor_cf11() T1 {
	return _this.Get_().(D1_DC6).Cf11
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC6).Cf12
}

func (_this D1) Dtor_cf13() bool {
	return _this.Get_().(D1_DC6).Cf13
}

func (_this D1) Dtor_cf7() *C0 {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf14() D1 {
	return _this.Get_().(D1_DC7).Cf14
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && _dafny.AreEqual(data1.Cf11, data2.Cf11) && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf7 == data2.Cf7
		}
	case D1_DC7:
		{
			data2, ok := other.Get_().(D1_DC7)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D2_DC9 struct {
	Cf16 bool
	Cf17 _dafny.Int
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf16 bool, Cf17 _dafny.Int) D2 {
	return D2{D2_DC9{Cf16, Cf17}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC8 struct {
	Cf15 _dafny.Array
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf15 _dafny.Array) D2 {
	return D2{D2_DC8{Cf15}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC10 struct {
	Cf18 D2
}

func (D2_DC10) isD2() {}

func (CompanionStruct_D2_) Create_DC10_(Cf18 D2) D2 {
	return D2{D2_DC10{Cf18}}
}

func (_this D2) Is_DC10() bool {
	_, ok := _this.Get_().(D2_DC10)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC9_(false, _dafny.Zero)
}

func (_this D2) Dtor_cf16() bool {
	return _this.Get_().(D2_DC9).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf17
}

func (_this D2) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) Dtor_cf18() D2 {
	return _this.Get_().(D2_DC10).Cf18
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC10:
		{
			return "D2.DC10" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D2_DC10:
		{
			data2, ok := other.Get_().(D2_DC10)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D3_DC12 struct {
	Cf20 bool
	Cf21 _dafny.Int
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf20 bool, Cf21 _dafny.Int) D3 {
	return D3{D3_DC12{Cf20, Cf21}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC13 struct {
	Cf22 *C0
	Cf23 _dafny.Int
}

func (D3_DC13) isD3() {}

func (CompanionStruct_D3_) Create_DC13_(Cf22 *C0, Cf23 _dafny.Int) D3 {
	return D3{D3_DC13{Cf22, Cf23}}
}

func (_this D3) Is_DC13() bool {
	_, ok := _this.Get_().(D3_DC13)
	return ok
}

type D3_DC11 struct {
	Cf19 _dafny.Sequence
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf19 _dafny.Sequence) D3 {
	return D3{D3_DC11{Cf19}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC14 struct {
	Cf24 D3
}

func (D3_DC14) isD3() {}

func (CompanionStruct_D3_) Create_DC14_(Cf24 D3) D3 {
	return D3{D3_DC14{Cf24}}
}

func (_this D3) Is_DC14() bool {
	_, ok := _this.Get_().(D3_DC14)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC12_(false, _dafny.Zero)
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC12).Cf20
}

func (_this D3) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D3_DC12).Cf21
}

func (_this D3) Dtor_cf22() *C0 {
	return _this.Get_().(D3_DC13).Cf22
}

func (_this D3) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D3_DC13).Cf23
}

func (_this D3) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf24() D3 {
	return _this.Get_().(D3_DC14).Cf24
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D3_DC13:
		{
			return "D3.DC13" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC14:
		{
			return "D3.DC14" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Cmp(data2.Cf21) == 0
		}
	case D3_DC13:
		{
			data2, ok := other.Get_().(D3_DC13)
			return ok && data1.Cf22 == data2.Cf22 && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf19.Equals(data2.Cf19)
		}
	case D3_DC14:
		{
			data2, ok := other.Get_().(D3_DC14)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D4_DC16 struct {
	Cf26 _dafny.Int
	Cf27 bool
	Cf28 T0
}

func (D4_DC16) isD4() {}

func (CompanionStruct_D4_) Create_DC16_(Cf26 _dafny.Int, Cf27 bool, Cf28 T0) D4 {
	return D4{D4_DC16{Cf26, Cf27, Cf28}}
}

func (_this D4) Is_DC16() bool {
	_, ok := _this.Get_().(D4_DC16)
	return ok
}

type D4_DC17 struct {
	Cf29 _dafny.Int
	Cf30 *C0
}

func (D4_DC17) isD4() {}

func (CompanionStruct_D4_) Create_DC17_(Cf29 _dafny.Int, Cf30 *C0) D4 {
	return D4{D4_DC17{Cf29, Cf30}}
}

func (_this D4) Is_DC17() bool {
	_, ok := _this.Get_().(D4_DC17)
	return ok
}

type D4_DC18 struct {
}

func (D4_DC18) isD4() {}

func (CompanionStruct_D4_) Create_DC18_() D4 {
	return D4{D4_DC18{}}
}

func (_this D4) Is_DC18() bool {
	_, ok := _this.Get_().(D4_DC18)
	return ok
}

type D4_DC15 struct {
	Cf25 _dafny.Sequence
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf25 _dafny.Sequence) D4 {
	return D4{D4_DC15{Cf25}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC19 struct {
	Cf31 D4
}

func (D4_DC19) isD4() {}

func (CompanionStruct_D4_) Create_DC19_(Cf31 D4) D4 {
	return D4{D4_DC19{Cf31}}
}

func (_this D4) Is_DC19() bool {
	_, ok := _this.Get_().(D4_DC19)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC16_(_dafny.Zero, false, (T0)(nil))
}

func (_this D4) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D4_DC16).Cf26
}

func (_this D4) Dtor_cf27() bool {
	return _this.Get_().(D4_DC16).Cf27
}

func (_this D4) Dtor_cf28() T0 {
	return _this.Get_().(D4_DC16).Cf28
}

func (_this D4) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D4_DC17).Cf29
}

func (_this D4) Dtor_cf30() *C0 {
	return _this.Get_().(D4_DC17).Cf30
}

func (_this D4) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D4_DC15).Cf25
}

func (_this D4) Dtor_cf31() D4 {
	return _this.Get_().(D4_DC19).Cf31
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC16:
		{
			return "D4.DC16" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D4_DC17:
		{
			return "D4.DC17" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D4_DC18:
		{
			return "D4.DC18"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC19:
		{
			return "D4.DC19" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC16:
		{
			data2, ok := other.Get_().(D4_DC16)
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27 == data2.Cf27 && _dafny.AreEqual(data1.Cf28, data2.Cf28)
		}
	case D4_DC17:
		{
			data2, ok := other.Get_().(D4_DC17)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30 == data2.Cf30
		}
	case D4_DC18:
		{
			_, ok := other.Get_().(D4_DC18)
			return ok
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D4_DC19:
		{
			data2, ok := other.Get_().(D4_DC19)
			return ok && data1.Cf31.Equals(data2.Cf31)
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

type D5_DC21 struct {
	Cf33 _dafny.Map
}

func (D5_DC21) isD5() {}

func (CompanionStruct_D5_) Create_DC21_(Cf33 _dafny.Map) D5 {
	return D5{D5_DC21{Cf33}}
}

func (_this D5) Is_DC21() bool {
	_, ok := _this.Get_().(D5_DC21)
	return ok
}

type D5_DC20 struct {
	Cf32 _dafny.Map
}

func (D5_DC20) isD5() {}

func (CompanionStruct_D5_) Create_DC20_(Cf32 _dafny.Map) D5 {
	return D5{D5_DC20{Cf32}}
}

func (_this D5) Is_DC20() bool {
	_, ok := _this.Get_().(D5_DC20)
	return ok
}

type D5_DC22 struct {
	Cf34 D5
}

func (D5_DC22) isD5() {}

func (CompanionStruct_D5_) Create_DC22_(Cf34 D5) D5 {
	return D5{D5_DC22{Cf34}}
}

func (_this D5) Is_DC22() bool {
	_, ok := _this.Get_().(D5_DC22)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC21_(_dafny.EmptyMap)
}

func (_this D5) Dtor_cf33() _dafny.Map {
	return _this.Get_().(D5_DC21).Cf33
}

func (_this D5) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D5_DC20).Cf32
}

func (_this D5) Dtor_cf34() D5 {
	return _this.Get_().(D5_DC22).Cf34
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC21:
		{
			return "D5.DC21" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D5_DC20:
		{
			return "D5.DC20" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D5_DC22:
		{
			return "D5.DC22" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC21:
		{
			data2, ok := other.Get_().(D5_DC21)
			return ok && data1.Cf33.Equals(data2.Cf33)
		}
	case D5_DC20:
		{
			data2, ok := other.Get_().(D5_DC20)
			return ok && data1.Cf32.Equals(data2.Cf32)
		}
	case D5_DC22:
		{
			data2, ok := other.Get_().(D5_DC22)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D6_DC23 struct {
	Cf35 _dafny.Array
}

func (D6_DC23) isD6() {}

func (CompanionStruct_D6_) Create_DC23_(Cf35 _dafny.Array) D6 {
	return D6{D6_DC23{Cf35}}
}

func (_this D6) Is_DC23() bool {
	_, ok := _this.Get_().(D6_DC23)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D6) Dtor_cf35() _dafny.Array {
	return _this.Get_().(D6_DC23).Cf35
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC23:
		{
			return "D6.DC23" + "(" + _dafny.String(data.Cf35) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC23:
		{
			data2, ok := other.Get_().(D6_DC23)
			return ok && data1.Cf35 == data2.Cf35
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
	Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map
	Fm1(p0 _dafny.Map, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) bool
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
	Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map
	Fm1(p0 _dafny.Map, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) bool
	F24() bool
	F25() _dafny.Sequence
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
	F1   _dafny.Set
	F4   _dafny.Sequence
	F5   _dafny.Int
	F8   _dafny.Sequence
	F14  _dafny.MultiSet
	F19  _dafny.Sequence
	F20  bool
	F21  bool
	_f0  _dafny.Sequence
	_f2  _dafny.Array
	_f3  bool
	_f6  _dafny.Map
	_f7  bool
	_f9  bool
	_f10 _dafny.Int
	_f11 bool
	_f12 bool
	_f13 bool
	_f15 _dafny.Array
	_f16 _dafny.Int
	_f17 bool
	_f18 _dafny.Array
	_f22 _dafny.Int
	_f23 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.EmptySet
	_this.F4 = _dafny.EmptySeq
	_this.F5 = _dafny.Zero
	_this.F8 = _dafny.EmptySeq
	_this.F14 = _dafny.EmptyMultiSet
	_this.F19 = _dafny.EmptySeq
	_this.F20 = false
	_this.F21 = false
	_this._f0 = _dafny.EmptySeq
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = false
	_this._f6 = _dafny.EmptyMap
	_this._f7 = false
	_this._f9 = false
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f12 = false
	_this._f13 = false
	_this._f15 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f16 = _dafny.Zero
	_this._f17 = false
	_this._f18 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f22 = _dafny.Zero
	_this._f23 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 _dafny.Set, f2 _dafny.Array, f3 bool, f4 _dafny.Sequence, f5 _dafny.Int, f6 _dafny.Map, f7 bool, f8 _dafny.Sequence, f9 bool, f10 _dafny.Int, f11 bool, f12 bool, f13 bool, f14 _dafny.MultiSet, f15 _dafny.Array, f16 _dafny.Int, f17 bool, f18 _dafny.Array, f19 _dafny.Sequence, f20 bool, f21 bool, f22 _dafny.Int, f23 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this).F20 = f20
		(_this).F21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
	}
}
func (_this *GlobalState) F0() _dafny.Sequence {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() _dafny.Array {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F6() _dafny.Map {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() _dafny.Array {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Array {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F22() _dafny.Int {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() bool {
	{
		return _this._f23
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f24 bool
	_f25 _dafny.Sequence
	_f26 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f24 = false
	_this._f25 = _dafny.EmptySeq
	_this._f26 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C0{}
var _ T1 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F24() bool {
	return _this._f24
}
func (_this *C0) F25() _dafny.Sequence {
	return _this._f25
}
func (_this *C0) Ctor__(f26 _dafny.Int, f24 bool, f25 _dafny.Sequence) {
	{
		(_this)._f26 = f26
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C0) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-128), (_this).F24())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F24()))
	}
}
func (_this *C0) Fm1(p0 _dafny.Map, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return ((_this).F24()) && ((_this).F24())
	}
}
func (_this *C0) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC1_((_this).F24(), (_this).F24(), (_this).F26())
	}
}
func (_this *C0) Fm3(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F26()
	}
}
func (_this *C0) F26() _dafny.Int {
	{
		return _this._f26
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
