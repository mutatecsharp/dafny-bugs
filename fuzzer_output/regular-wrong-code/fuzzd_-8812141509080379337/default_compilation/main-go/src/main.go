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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(74)
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	return ((_dafny.IntOfInt64(151)).Plus((_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality()))).Cmp(_dafny.IntOfInt64(-621)) > 0
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("awbiybj"), _dafny.UnicodeSeqOfUtf8Bytes("d"))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('f')
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(true))).Intersection(_dafny.SetOf(true))).Union(_dafny.SetOf(true, !(true), !(true)))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.CodePoint('d'))).Difference((_dafny.SetOf(_dafny.CodePoint('o'))).Difference(_dafny.SetOf(_dafny.CodePoint('s'))))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (bool, _dafny.Set) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Set = _dafny.EmptySet
	_ = r1
	var _0_v0 bool
	_ = _0_v0
	_0_v0 = false
	var _1_i0 _dafny.Int
	_ = _1_i0
	_1_i0 = _dafny.Zero
	{
		for _0_v0 {
			{
				if (_1_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_1_i0 = (_1_i0).Plus(_dafny.One)
				var _2_v1 _dafny.CodePoint
				_ = _2_v1
				_2_v1 = _dafny.CodePoint('q')
				var _3_v2 _dafny.Sequence
				_ = _3_v2
				_3_v2 = _dafny.SeqOf(_dafny.CodePoint('k'), _2_v1, _2_v1)
				var _4_v3 _dafny.Map
				_ = _4_v3
				_4_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_0_v0, p0)
				var _5_v4 _dafny.Set
				_ = _5_v4
				_5_v4 = _dafny.SetOf(p1, p0, p0, (_4_v3).Cardinality(), p0)
				var _6_v5 *C0
				_ = _6_v5
				var _nw0 *C0 = New_C0_()
				_ = _nw0
				_nw0.Ctor__(_3_v2, (p0).Plus(p1), _3_v2, (_5_v4).Contains(p1))
				_6_v5 = _nw0
				_6_v5 = _6_v5
				var _7_v6 _dafny.Map
				_ = _7_v6
				_7_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hpidlcdvp"), _3_v2))
				var _8_v7 _dafny.Set
				_ = _8_v7
				_8_v7 = _dafny.SetOf(_0_v0)
				_3_v2 = (func() _dafny.Sequence {
					if (_7_v6).Contains(((func() _dafny.Set {
						if _0_v0 {
							return _8_v7
						}
						return Companion_Default___.Fm7(p0, globalState)
					})()).Cardinality()) {
						return (_7_v6).Get(((func() _dafny.Set {
							if _0_v0 {
								return _8_v7
							}
							return Companion_Default___.Fm7(p0, globalState)
						})()).Cardinality()).(_dafny.Sequence)
					}
					return _dafny.Companion_Sequence_.Concatenate((_6_v5).F7(), _3_v2)
				})()
				var _9_v8 *C0
				_ = _9_v8
				var _nw1 *C0 = New_C0_()
				_ = _nw1
				_nw1.Ctor__(_3_v2, p1, (_6_v5).F7(), _0_v0)
				_9_v8 = _nw1
				(globalState).F1 = p0
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _10_v9 _dafny.CodePoint
	_ = _10_v9
	_10_v9 = _dafny.CodePoint('p')
	var _11_v10 _dafny.Sequence
	_ = _11_v10
	_11_v10 = _dafny.UnicodeSeqOfUtf8Bytes("u")
	var _12_v11 *C0
	_ = _12_v11
	var _nw2 *C0 = New_C0_()
	_ = _nw2
	_nw2.Ctor__(_dafny.SeqOf(_10_v9), _dafny.IntOfUint32((_11_v10).Cardinality()), _11_v10, _0_v0)
	_12_v11 = _nw2
	var _13_v12 _dafny.Map
	_ = _13_v12
	_13_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(795))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_14_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})), (_12_v11).F8())
	var _15_v13 D6
	_ = _15_v13
	_15_v13 = Companion_D6_.Create_DC15_(p0, _12_v11, _13_v12, _0_v0)
	if (_15_v13).Dtor_cf33() {
		var _16_v14 _dafny.MultiSet
		_ = _16_v14
		_16_v14 = _dafny.MultiSetOf(_0_v0, _0_v0, _0_v0)
		(globalState).F1 = ((_16_v14).Difference((_16_v14).Union(_16_v14))).Cardinality()
		(globalState).F1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-28), p1)
		var _17_v15 _dafny.Map
		_ = _17_v15
		_17_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_0_v0, _10_v9)
		var _pat_let_tv0 = _13_v12
		_ = _pat_let_tv0
		var _pat_let_tv1 = _0_v0
		_ = _pat_let_tv1
		var _pat_let_tv2 = _12_v11
		_ = _pat_let_tv2
		var _18_v16 _dafny.Map
		_ = _18_v16
		_18_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let0_0 D6) D6 {
			return func(_19_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let1_0 _dafny.Map) D6 {
					return func(_20_dt__update_hcf32_h0 _dafny.Map) D6 {
						return func(_pat_let2_0 bool) D6 {
							return func(_21_dt__update_hcf33_h0 bool) D6 {
								return func(_pat_let3_0 *C0) D6 {
									return func(_22_dt__update_hcf31_h0 *C0) D6 {
										return Companion_D6_.Create_DC15_((_19_dt__update__tmp_h0).Dtor_cf30(), _22_dt__update_hcf31_h0, _20_dt__update_hcf32_h0, _21_dt__update_hcf33_h0)
									}(_pat_let3_0)
								}(_pat_let_tv2)
							}(_pat_let2_0)
						}(_pat_let_tv1)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_15_v13)).Dtor_cf33(), p1)
		var _23_v17 _dafny.MultiSet
		_ = _23_v17
		_23_v17 = _dafny.MultiSetOf(_10_v9, _10_v9, _10_v9)
		var _24_v18 _dafny.Map
		_ = _24_v18
		_24_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_12_v11, _0_v0)
		var _25_v19 _dafny.Sequence
		_ = _25_v19
		_25_v19 = _dafny.SeqOf(_0_v0)
		var _26_v20 _dafny.Array
		_ = _26_v20
		var _nwElement0_0 _dafny.Int = ((_12_v11).F8()).Times((_12_v11).F8())
		_ = _nwElement0_0
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(28))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_0, 0)
		(_nw3).ArraySet1(p0, 1)
		(_nw3).ArraySet1((p0).Times((_dafny.Zero).Minus(p1)), 2)
		(_nw3).ArraySet1((p0).Plus(p1), 3)
		(_nw3).ArraySet1(_dafny.IntOfInt64(-691), 4)
		(_nw3).ArraySet1((_12_v11).F8(), 5)
		(_nw3).ArraySet1(p1, 6)
		(_nw3).ArraySet1((_12_v11).F8(), 7)
		(_nw3).ArraySet1((_17_v15).Cardinality(), 8)
		(_nw3).ArraySet1((func() _dafny.Int {
			if (_18_v16).Contains(_0_v0) {
				return (_18_v16).Get(_0_v0).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_11_v10).Cardinality())
		})(), 9)
		(_nw3).ArraySet1((_12_v11).F8(), 10)
		(_nw3).ArraySet1((_12_v11).F8(), 11)
		(_nw3).ArraySet1(p0, 12)
		(_nw3).ArraySet1(p1, 13)
		(_nw3).ArraySet1(p1, 14)
		(_nw3).ArraySet1((_12_v11).F8(), 15)
		(_nw3).ArraySet1(p1, 16)
		(_nw3).ArraySet1(((_12_v11).F8()).Times((_15_v13).Dtor_cf30()), 17)
		(_nw3).ArraySet1(p0, 18)
		(_nw3).ArraySet1((_12_v11).F8(), 19)
		(_nw3).ArraySet1((_12_v11).F8(), 20)
		(_nw3).ArraySet1((func() _dafny.Int {
			if (_23_v17).Contains(_10_v9) {
				return (_23_v17).Multiplicity(_10_v9)
			}
			return (_12_v11).F8()
		})(), 21)
		(_nw3).ArraySet1((_12_v11).F8(), 22)
		(_nw3).ArraySet1((_12_v11).F8(), 23)
		(_nw3).ArraySet1(p0, 24)
		(_nw3).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, (_24_v18).Cardinality())), 25)
		(_nw3).ArraySet1(_dafny.IntOfUint32((_25_v19).Cardinality()), 26)
		(_nw3).ArraySet1((_12_v11).F8(), 27)
		_26_v20 = _nw3
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_26_v20), 0))
		_ = _index0
		(_26_v20).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_16_v14).Contains(_0_v0) {
				return (_16_v14).Multiplicity(_0_v0)
			}
			return (_12_v11).F8()
		})(), p1), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_26_v20), 0))
		_ = _index1
		(_26_v20).ArraySet1((_12_v11).F8(), (_index1).Int())
		_25_v19 = _25_v19
		(globalState).F0 = Companion_Default___.Fm1(_dafny.IntOfInt64(740), _0_v0, globalState)
	} else {
		var _27_v21 bool
		_ = _27_v21
		_27_v21 = _0_v0
		var _28_v22 _dafny.Array
		_ = _28_v22
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw4
		_28_v22 = _nw4
		var _29_v23 _dafny.Map
		_ = _29_v23
		_29_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v21, _28_v22)
		_29_v23 = ((_29_v23).Merge(_29_v23)).Merge(_29_v23)
		var _30_v24 *C0
		_ = _30_v24
		var _nw5 *C0 = New_C0_()
		_ = _nw5
		_nw5.Ctor__(_11_v10, (_12_v11).F8(), (_12_v11).F7(), (_0_v0) && (_0_v0))
		_30_v24 = _nw5
		var _31_v25 _dafny.Sequence
		_ = _31_v25
		_31_v25 = _dafny.SeqOf(true)
		var _32_v26 *C0
		_ = _32_v26
		var _nw6 *C0 = New_C0_()
		_ = _nw6
		_nw6.Ctor__(_dafny.Companion_Sequence_.Concatenate((_30_v24).F7(), _dafny.SeqOf(_10_v9, Companion_Default___.Fm6(globalState), _10_v9)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qf")).Cardinality()), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}((func(_33_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_34_i2 _dafny.Int) _dafny.CodePoint {
				return _33_v9
			}
		})(_10_v9))), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}((func(_35_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_36_i2 _dafny.Int) _dafny.CodePoint {
				return _35_v9
			}
		})(_10_v9)))).Cardinality()))).Uint32(), _10_v9), !((_31_v25).Select((Companion_Default___.SafeIndex((_30_v24).F8(), _dafny.IntOfUint32((_31_v25).Cardinality()))).Uint32()).(bool)) || (_0_v0))
		_32_v26 = _nw6
		var _37_v27 _dafny.Map
		_ = _37_v27
		_37_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_10_v9, (_30_v24).F8())
		var _38_v28 _dafny.Set
		_ = _38_v28
		_38_v28 = _dafny.SetOf(_10_v9)
		var _39_v29 _dafny.Map
		_ = _39_v29
		_39_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_37_v27).Equals(_37_v27), (func() _dafny.Set {
			if _0_v0 {
				return _38_v28
			}
			return Companion_Default___.Fm8(globalState)
		})())
		_39_v29 = (_39_v29).Update(_0_v0, _38_v28)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_28_v22), 0))
		_ = _index2
		(_28_v22).ArraySet1((_32_v26).F8(), (_index2).Int())
		var _40_v30 _dafny.Map
		_ = _40_v30
		_40_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v22, _0_v0)
		var _41_v31 _dafny.Map
		_ = _41_v31
		_41_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_40_v30).Cardinality(), _dafny.IntOfInt64(754))
		var _42_v32 _dafny.Map
		_ = _42_v32
		_42_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_0_v0, _0_v0)
		var _43_v33 _dafny.Sequence
		_ = _43_v33
		_43_v33 = _dafny.SeqOf((_dafny.Zero).Minus((_41_v31).Cardinality()), (_42_v32).Cardinality())
		var _44_v34 _dafny.Map
		_ = _44_v34
		_44_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_0_v0), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(415))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_45_i3 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(537)
		}))).Cardinality())).Plus(_dafny.IntOfUint32((_43_v33).Cardinality())))
		var _46_v35 _dafny.Set
		_ = _46_v35
		_46_v35 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_32_v26).F8())))
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_28_v22), 0))
		_ = _index3
		(_28_v22).ArraySet1((func() _dafny.Int {
			if (_44_v34).Contains(!_dafny.Companion_Sequence_.Equal(_31_v25, _31_v25)) {
				return (_44_v34).Get(!_dafny.Companion_Sequence_.Equal(_31_v25, _31_v25)).(_dafny.Int)
			}
			return ((_dafny.SetOf(p0)).Union(_46_v35)).Cardinality()
		})(), (_index3).Int())
	}
	var _hi0 _dafny.Int = p1
	_ = _hi0
	for _47_i4 := (_12_v11).F8(); _47_i4.Cmp(_hi0) < 0; _47_i4 = _47_i4.Plus(_dafny.One) {
		var _48_v36 _dafny.Sequence
		_ = _48_v36
		_48_v36 = _dafny.SeqOf(p1, p1)
		if Companion_Default___.Fm1((func() _dafny.Int {
			if _0_v0 {
				return p1
			}
			return (_48_v36).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_48_v36).Cardinality()))).Uint32()).(_dafny.Int)
		})(), _dafny.Companion_Sequence_.IsProperPrefixOf(_11_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(324))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}((func(_49_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_50_i5 _dafny.Int) _dafny.CodePoint {
				return _49_v9
			}
		})(_10_v9)))), globalState) {
			_11_v10 = (_12_v11).F7()
			(globalState).F0 = !((_0_v0) == (_0_v0))
			(globalState).F1 = (func() _dafny.Int {
				if (func() bool {
					if false {
						return _0_v0
					}
					return _0_v0
				})() {
					return (_12_v11).F8()
				}
				return (_dafny.Zero).Minus(p1)
			})()
			var _51_v37 T1
			_ = _51_v37
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__(_11_v10, (_12_v11).F8(), (_12_v11).F7(), true)
			_51_v37 = _nw7
			var _52_v38 _dafny.Set
			_ = _52_v38
			_52_v38 = _dafny.SetOf(_51_v37, _51_v37)
			var _53_v39 _dafny.MultiSet
			_ = _53_v39
			_53_v39 = _dafny.MultiSetOf(_0_v0, _0_v0)
			var _54_v40 _dafny.Map
			_ = _54_v40
			_54_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_52_v38, _53_v39)
			var _rhs0 _dafny.Int = _dafny.IntOfUint32(((_12_v11).F7()).Cardinality())
			_ = _rhs0
			var _rhs1 _dafny.CodePoint = Companion_Default___.Fm6(globalState)
			_ = _rhs1
			var _rhs2 _dafny.Int = Companion_Default___.SafeDivisionInt(((_54_v40).Merge(_54_v40)).Cardinality(), (_dafny.Zero).Minus((_12_v11).F8()))
			_ = _rhs2
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			var _lhs1 *GlobalState = globalState
			_ = _lhs1
			_lhs0.F1 = _rhs0
			_10_v9 = _rhs1
			_lhs1.F1 = _rhs2
			var _55_v41 _dafny.Map
			_ = _55_v41
			_55_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_0_v0, (_12_v11).F8())
			_55_v41 = (_55_v41).Update(_0_v0, Companion_Default___.SafeModuloInt(p0, p0))
		} else {
			_10_v9 = _dafny.CodePoint('i')
			var _56_v42 _dafny.Sequence
			_ = _56_v42
			_56_v42 = _dafny.SeqOf(_12_v11, _12_v11, _12_v11, _12_v11, _12_v11)
			_56_v42 = _dafny.Companion_Sequence_.Concatenate(_56_v42, _56_v42)
			var _57_v43 _dafny.Sequence
			_ = _57_v43
			_57_v43 = _dafny.SeqOf(_0_v0, false, _0_v0)
			var _58_v44 _dafny.Array
			_ = _58_v44
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_0
			var _nw8 _dafny.Array
			_ = _nw8
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw8 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) bool = (func(_59_v0 bool) func(_dafny.Int) bool {
					return func(_60_i6 _dafny.Int) bool {
						return _59_v0
					}
				})(_0_v0)
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
			_58_v44 = _nw8
			var _61_v45 _dafny.Sequence
			_ = _61_v45
			_61_v45 = _dafny.SeqOf(_58_v44, _58_v44)
			var _62_v46 _dafny.Map
			_ = _62_v46
			_62_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_61_v45).Cardinality()), _dafny.Companion_Sequence_.Contains((_12_v11).F7(), _10_v9))
			var _rhs3 _dafny.Int = (_12_v11).Fm4((_12_v11).F8(), _dafny.IntOfInt64(843), _dafny.Companion_Sequence_.Concatenate(_57_v43, Companion_Default___.Fm9(_dafny.UnicodeSeqOfUtf8Bytes("klg"), p0, globalState)), globalState)
			_ = _rhs3
			var _rhs4 _dafny.Int = (_62_v46).Cardinality()
			_ = _rhs4
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			_lhs2.F1 = _rhs3
			_lhs3.F1 = _rhs4
			_12_v11 = _12_v11
			_12_v11 = (_15_v13).Dtor_cf31()
		}
		_13_v12 = (_13_v12).Update(_11_v10, (_12_v11).F8())
		var _63_v47 T1
		_ = _63_v47
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__((_12_v11).F7(), (_12_v11).F8(), _dafny.UnicodeSeqOfUtf8Bytes("cwoxfwq"), _0_v0)
		_63_v47 = _nw9
		var _64_v48 bool
		_ = _64_v48
		_64_v48 = _0_v0
		var _65_v49 _dafny.Map
		_ = _65_v49
		_65_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_0_v0, _0_v0)
		var _66_v50 _dafny.Array
		_ = _66_v50
		var _nwElement0_1 bool = _0_v0
		_ = _nwElement0_1
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(29))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_1, 0)
		(_nw10).ArraySet1(true, 1)
		(_nw10).ArraySet1(false, 2)
		(_nw10).ArraySet1(_0_v0, 3)
		(_nw10).ArraySet1(_0_v0, 4)
		(_nw10).ArraySet1(_0_v0, 5)
		(_nw10).ArraySet1(_0_v0, 6)
		(_nw10).ArraySet1((_64_v48), 7)
		(_nw10).ArraySet1(_0_v0, 8)
		(_nw10).ArraySet1(true, 9)
		(_nw10).ArraySet1(_63_v47.F6(), 10)
		(_nw10).ArraySet1(_0_v0, 11)
		(_nw10).ArraySet1(true, 12)
		(_nw10).ArraySet1(_63_v47.F6(), 13)
		(_nw10).ArraySet1(_0_v0, 14)
		(_nw10).ArraySet1(_0_v0, 15)
		(_nw10).ArraySet1(_63_v47.F6(), 16)
		(_nw10).ArraySet1(_63_v47.F6(), 17)
		(_nw10).ArraySet1(_63_v47.F6(), 18)
		(_nw10).ArraySet1(!(_0_v0), 19)
		(_nw10).ArraySet1(_63_v47.F6(), 20)
		(_nw10).ArraySet1(_63_v47.F6(), 21)
		(_nw10).ArraySet1(_0_v0, 22)
		(_nw10).ArraySet1(_63_v47.F6(), 23)
		(_nw10).ArraySet1(_63_v47.F6(), 24)
		(_nw10).ArraySet1(false, 25)
		(_nw10).ArraySet1(_63_v47.F6(), 26)
		(_nw10).ArraySet1((func() bool {
			if (_65_v49).Contains(_63_v47.F6()) {
				return (_65_v49).Get(_63_v47.F6()).(bool)
			}
			return _0_v0
		})(), 27)
		(_nw10).ArraySet1(_63_v47.F6(), 28)
		_66_v50 = _nw10
		var _67_v51 _dafny.Map
		_ = _67_v51
		_67_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v47, _66_v50)
		_67_v51 = (_67_v51).Update(_63_v47, _66_v50)
		if (_47_i4).Cmp(p0) > 0 {
			var _68_v52 _dafny.Array
			_ = _68_v52
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw11
			_68_v52 = _nw11
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_68_v52), 0))
			_ = _index4
			(_68_v52).ArraySet1((_48_v36).Select((Companion_Default___.SafeIndex((_12_v11).F8(), _dafny.IntOfUint32((_48_v36).Cardinality()))).Uint32()).(_dafny.Int), (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_68_v52), 0))
			_ = _index5
			(_68_v52).ArraySet1((_12_v11).F8(), (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_68_v52), 0))
			_ = _index6
			(_68_v52).ArraySet1(_dafny.IntOfInt64(749), (_index6).Int())
			var _69_v53 *C0
			_ = _69_v53
			var _nw12 *C0 = New_C0_()
			_ = _nw12
			_nw12.Ctor__((_12_v11).F7(), _dafny.IntOfInt64(617), (_63_v47).F5(), _63_v47.F6())
			_69_v53 = _nw12
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_66_v50), 0))
			_ = _index7
			(_66_v50).ArraySet1(_63_v47.F6(), (_index7).Int())
			var _70_v54 D6
			_ = _70_v54
			_70_v54 = Companion_D6_.Create_DC13_(_63_v47.F6(), !(_0_v0))
			var _71_v55 _dafny.Array
			_ = _71_v55
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_1
			var _nw13 _dafny.Array
			_ = _nw13
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw13 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Set = (func(_72_p0 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_73_i7 _dafny.Int) _dafny.Set {
						return _dafny.SetOf(_72_p0)
					}
				})(p0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw13 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw13).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw13).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_71_v55 = _nw13
			var _74_v56 _dafny.Set
			_ = _74_v56
			_74_v56 = _dafny.SetOf((_12_v11).F8(), (_68_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_68_v52), 0))).Int()).(_dafny.Int), p1, (_69_v53).F8())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_71_v55), 0))
			_ = _index8
			(_71_v55).ArraySet1((_74_v56).Union(_74_v56), (_index8).Int())
			var _75_v57 _dafny.Array
			_ = _75_v57
			var _nwElement0_2 *C0 = (func() *C0 {
				if _63_v47.F6() {
					return _12_v11
				}
				return _12_v11
			})()
			_ = _nwElement0_2
			var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
			_ = _nw14
			(_nw14).ArraySet1(_nwElement0_2, 0)
			(_nw14).ArraySet1(_12_v11, 1)
			_75_v57 = _nw14
			var _76_v58 _dafny.Map
			_ = _76_v58
			_76_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_12_v11, _65_v49)
			var _77_v59 _dafny.Sequence
			_ = _77_v59
			_77_v59 = _dafny.SeqOf(_74_v56)
			var _pat_let_tv3 = _0_v0
			_ = _pat_let_tv3
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_66_v50), 0))
			_ = _index9
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_71_v55), 0))
			_ = _index10
			var _rhs5 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_0_v0, _63_v47.F6())).Merge(((func() _dafny.Map {
				if (_76_v58).Contains(_12_v11) {
					return (_76_v58).Get(_12_v11).(_dafny.Map)
				}
				return _65_v49
			})()).Update(_0_v0, _0_v0))
			_ = _rhs5
			var _rhs6 bool = _63_v47.F6()
			_ = _rhs6
			var _rhs7 D6 = func(_pat_let4_0 D6) D6 {
				return func(_78_dt__update__tmp_h1 D6) D6 {
					return func(_pat_let5_0 bool) D6 {
						return func(_79_dt__update_hcf26_h0 bool) D6 {
							return Companion_D6_.Create_DC13_(_79_dt__update_hcf26_h0, (_78_dt__update__tmp_h1).Dtor_cf27())
						}(_pat_let5_0)
					}(_pat_let_tv3)
				}(_pat_let4_0)
			}(_70_v54)
			_ = _rhs7
			var _rhs8 _dafny.Set = (_77_v59).Select((Companion_Default___.SafeIndex((_69_v53).F8(), _dafny.IntOfUint32((_77_v59).Cardinality()))).Uint32()).(_dafny.Set)
			_ = _rhs8
			var _rhs9 _dafny.Array = _75_v57
			_ = _rhs9
			var _lhs4 _dafny.Array = _66_v50
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_66_v50), 0))
			_ = _lhs5
			var _lhs6 _dafny.Array = _71_v55
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_71_v55), 0))
			_ = _lhs7
			_65_v49 = _rhs5
			(_lhs4).ArraySet1(_rhs6, (_lhs5).Int())
			_70_v54 = _rhs7
			(_lhs6).ArraySet1(_rhs8, (_lhs7).Int())
			_75_v57 = _rhs9
			(globalState).F1 = ((_dafny.Zero).Minus((_dafny.Zero).Minus(p0))).Plus((_12_v11).F8())
		} else {
			var _80_v60 _dafny.Array
			_ = _80_v60
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw15
			_80_v60 = _nw15
			var _81_v61 T0
			_ = _81_v61
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__(_dafny.SeqOf(Companion_Default___.Fm6(globalState)), (_dafny.Zero).Minus((_12_v11).F8()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(264))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_82_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_83_i8 _dafny.Int) _dafny.CodePoint {
					return _82_v9
				}
			})(_10_v9))), _63_v47.F6())
			_81_v61 = _nw16
			var _84_v62 _dafny.Sequence
			_ = _84_v62
			_84_v62 = _dafny.SeqOf(_81_v61)
			var _85_v63 _dafny.Sequence
			_ = _85_v63
			_85_v63 = _dafny.SeqOf(_84_v62, _84_v62, _84_v62, _84_v62, _84_v62)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_80_v60), 0))
			_ = _index11
			(_80_v60).ArraySet1((_85_v63).Select((Companion_Default___.SafeIndex(_47_i4, _dafny.IntOfUint32((_85_v63).Cardinality()))).Uint32()).(_dafny.Sequence), (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_80_v60), 0))
			_ = _index12
			(_80_v60).ArraySet1(_84_v62, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_66_v50), 0))
			_ = _index13
			(_66_v50).ArraySet1(_63_v47.F6(), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_66_v50), 0))
			_ = _index14
			(_66_v50).ArraySet1(_63_v47.F6(), (_index14).Int())
			var _86_v64 _dafny.Map
			_ = _86_v64
			_86_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v47, (_12_v11).F8())
			_86_v64 = (_86_v64).Update(_63_v47, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_48_v36, _48_v36)).Cardinality()))
			(_63_v47).F6_set_((_66_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_66_v50), 0))).Int()).(bool))
			var _87_v65 *C0
			_ = _87_v65
			var _nw17 *C0 = New_C0_()
			_ = _nw17
			_nw17.Ctor__((_12_v11).F7(), (_12_v11).F8(), (_81_v61).F5(), !(_63_v47.F6()) || (true))
			_87_v65 = _nw17
		}
	}
	_11_v10 = _dafny.Companion_Sequence_.Concatenate(_11_v10, (_12_v11).F7())
	var _88_v66 _dafny.Set
	_ = _88_v66
	_88_v66 = _dafny.SetOf(_0_v0)
	var _89_v67 _dafny.Sequence
	_ = _89_v67
	_89_v67 = _dafny.SeqOf(_0_v0, !(_0_v0))
	var _hi1 _dafny.Int = (_dafny.MultiSetFromSeq(_89_v67)).Cardinality()
	_ = _hi1
	for _90_i9 := (_88_v66).Cardinality(); _90_i9.Cmp(_hi1) < 0; _90_i9 = _90_i9.Plus(_dafny.One) {
		var _91_v68 _dafny.Array
		_ = _91_v68
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
		_ = _nw18
		_91_v68 = _nw18
		var _92_v69 _dafny.Set
		_ = _92_v69
		_92_v69 = _dafny.SetOf(p1)
		var _93_v70 _dafny.Map
		_ = _93_v70
		_93_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), false)
		var _94_v71 _dafny.MultiSet
		_ = _94_v71
		_94_v71 = _dafny.MultiSetOf(p0)
		var _95_v72 _dafny.Sequence
		_ = _95_v72
		_95_v72 = _dafny.SeqOf((_94_v71).Cardinality(), p1)
		var _96_v73 _dafny.Array
		_ = _96_v73
		var _nwElement0_3 _dafny.Int = (_92_v69).Cardinality()
		_ = _nwElement0_3
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(11))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_3, 0)
		(_nw19).ArraySet1(_90_i9, 1)
		(_nw19).ArraySet1((_dafny.Zero).Minus((_93_v70).Cardinality()), 2)
		(_nw19).ArraySet1(_dafny.IntOfInt64(-43), 3)
		(_nw19).ArraySet1(_dafny.IntOfUint32((_95_v72).Cardinality()), 4)
		(_nw19).ArraySet1(_90_i9, 5)
		(_nw19).ArraySet1((_94_v71).Cardinality(), 6)
		(_nw19).ArraySet1(_90_i9, 7)
		(_nw19).ArraySet1((_12_v11).F8(), 8)
		(_nw19).ArraySet1(_dafny.IntOfInt64(-71), 9)
		(_nw19).ArraySet1((_95_v72).Select((Companion_Default___.SafeIndex((_12_v11).F8(), _dafny.IntOfUint32((_95_v72).Cardinality()))).Uint32()).(_dafny.Int), 10)
		_96_v73 = _nw19
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_91_v68), 0))
		_ = _index15
		(_91_v68).ArraySet1(_96_v73, (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_91_v68), 0))
		_ = _index16
		(_91_v68).ArraySet1((func() _dafny.Array {
			if _0_v0 {
				return (_96_v73)
			}
			return _96_v73
		})(), (_index16).Int())
		var _97_v74 *C0
		_ = _97_v74
		var _nw20 *C0 = New_C0_()
		_ = _nw20
		_nw20.Ctor__((_12_v11).F7(), p0, (_12_v11).F7(), _0_v0)
		_97_v74 = _nw20
		var _98_v75 _dafny.Array
		_ = _98_v75
		var _nw21 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw21
		_98_v75 = _nw21
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_98_v75), 0))
		_ = _index17
		(_98_v75).ArraySet1(_12_v11, (_index17).Int())
		var _99_v76 _dafny.Array
		_ = _99_v76
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw22
		_99_v76 = _nw22
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_99_v76), 0))
		_ = _index18
		(_99_v76).ArraySet1((_0_v0) && (_0_v0), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_99_v76), 0))
		_ = _index19
		(_99_v76).ArraySet1(!(_88_v66).Contains(!(_0_v0)), (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_96_v73), 0))
		_ = _index20
		(_96_v73).ArraySet1((_12_v11).F8(), (_index20).Int())
		var _100_v77 _dafny.Sequence
		_ = _100_v77
		_100_v77 = _dafny.SeqOf(_97_v74, _12_v11, _12_v11)
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_98_v75), 0))
		_ = _index21
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_99_v76), 0))
		_ = _index22
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_99_v76), 0))
		_ = _index23
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_96_v73), 0))
		_ = _index24
		var _rhs10 *C0 = (_100_v77).Select((Companion_Default___.SafeIndex((_12_v11).F8(), _dafny.IntOfUint32((_100_v77).Cardinality()))).Uint32()).(*C0)
		_ = _rhs10
		var _rhs11 bool = !(!(_0_v0))
		_ = _rhs11
		var _rhs12 _dafny.Sequence = (_97_v74).F7()
		_ = _rhs12
		var _rhs13 bool = _0_v0
		_ = _rhs13
		var _rhs14 _dafny.Int = (_12_v11).F8()
		_ = _rhs14
		var _lhs8 _dafny.Array = _98_v75
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_98_v75), 0))
		_ = _lhs9
		var _lhs10 _dafny.Array = _99_v76
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_99_v76), 0))
		_ = _lhs11
		var _lhs12 _dafny.Array = _99_v76
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_99_v76), 0))
		_ = _lhs13
		var _lhs14 _dafny.Array = _96_v73
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_96_v73), 0))
		_ = _lhs15
		(_lhs8).ArraySet1(_rhs10, (_lhs9).Int())
		(_lhs10).ArraySet1(_rhs11, (_lhs11).Int())
		_11_v10 = _rhs12
		(_lhs12).ArraySet1(_rhs13, (_lhs13).Int())
		(_lhs14).ArraySet1(_rhs14, (_lhs15).Int())
		var _101_v78 _dafny.Map
		_ = _101_v78
		_101_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_99_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_99_v76), 0))).Int()).(bool) {
				return p0
			}
			return (_12_v11).F8()
		})(), (_97_v74).F8())
		var _102_v79 D6
		_ = _102_v79
		_102_v79 = Companion_D6_.Create_DC13_(_0_v0, _0_v0)
		var _103_v80 D6
		_ = _103_v80
		_103_v80 = Companion_D6_.Create_DC16_(_102_v79)
		var _104_v81 _dafny.Map
		_ = _104_v81
		_104_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v80, true)
		var _105_v82 _dafny.Map
		_ = _105_v82
		_105_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_0_v0, _104_v81)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_96_v73), 0))
		_ = _index25
		(_96_v73).ArraySet1((func() _dafny.Int {
			if (_101_v78).Contains((_97_v74).F8()) {
				return (_101_v78).Get((_97_v74).F8()).(_dafny.Int)
			}
			return (_dafny.Zero).Minus((_105_v82).Cardinality())
		})(), (_index25).Int())
	}
	_13_v12 = (_13_v12).Update(_11_v10, ((_12_v11).F8()).Times(p1))
	r0 = _0_v0
	var _106_v83 _dafny.Set
	_ = _106_v83
	_106_v83 = _dafny.SetOf((_12_v11).F8())
	r1 = _106_v83
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _107_globalState *GlobalState
	_ = _107_globalState
	var _nw23 *GlobalState = New_GlobalState_()
	_ = _nw23
	_nw23.Ctor__(false, _dafny.IntOfInt64(475), _dafny.CodePoint('r'), _dafny.IntOfInt64(449), false)
	_107_globalState = _nw23
	var _108_v0 _dafny.Int
	_ = _108_v0
	_108_v0 = _dafny.IntOfInt64(-316)
	(_107_globalState).F1 = _108_v0
	var _109_v1 _dafny.Array
	_ = _109_v1
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_2
	var _nw24 _dafny.Array
	_ = _nw24
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw24 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Sequence = func(_110_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eas"), _dafny.UnicodeSeqOfUtf8Bytes("opxmqoxt"))
		}
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw24 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw24).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw24).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_109_v1 = _nw24
	var _111_v2 _dafny.Sequence
	_ = _111_v2
	_111_v2 = _dafny.UnicodeSeqOfUtf8Bytes("eedslel")
	var _112_v3 _dafny.Map
	_ = _112_v3
	_112_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _111_v2)
	var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
	_ = _index26
	(_109_v1).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if (_112_v3).Contains(false) {
			return (_112_v3).Get(false).(_dafny.Sequence)
		}
		return _111_v2
	})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-891))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_113_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))), (_index26).Int())
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
	_ = _index27
	(_109_v1).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("osndtkw"), (_index27).Int())
	var _114_v4 bool
	_ = _114_v4
	var _115_v5 _dafny.Set
	_ = _115_v5
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Set
	_ = _out1
	_out0, _out1 = Companion_Default___.M0(_108_v0, (_dafny.Zero).Minus(_108_v0), _107_globalState)
	_114_v4 = _out0
	_115_v5 = _out1
	var _116_v6 _dafny.Array
	_ = _116_v6
	var _nw25 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
	_ = _nw25
	_116_v6 = _nw25
	var _117_v7 _dafny.Set
	_ = _117_v7
	_117_v7 = _dafny.SetOf(_116_v6, _116_v6)
	var _rhs15 _dafny.Set = (_dafny.SetOf(_116_v6)).Difference(_117_v7)
	_ = _rhs15
	var _rhs16 _dafny.Int = _108_v0
	_ = _rhs16
	var _rhs17 bool = _114_v4
	_ = _rhs17
	var _lhs16 *GlobalState = _107_globalState
	_ = _lhs16
	_117_v7 = _rhs15
	_lhs16.F1 = _rhs16
	_114_v4 = _rhs17
	if ((_117_v7).Difference(_117_v7)).IsProperSubsetOf(_dafny.SetOf(_116_v6)) {
		var _118_v8 _dafny.Sequence
		_ = _118_v8
		_118_v8 = _dafny.SeqOf(_114_v4)
		var _119_v9 _dafny.MultiSet
		_ = _119_v9
		_119_v9 = _dafny.MultiSetOf(_108_v0, _108_v0)
		var _120_v10 _dafny.MultiSet
		_ = _120_v10
		_120_v10 = _dafny.MultiSetOf(_114_v4)
		var _121_v11 _dafny.CodePoint
		_ = _121_v11
		_121_v11 = _dafny.CodePoint('p')
		var _rhs18 bool = _dafny.Companion_Sequence_.Contains(_118_v8, _114_v4)
		_ = _rhs18
		var _rhs19 bool = _114_v4
		_ = _rhs19
		var _rhs20 _dafny.Int = Companion_Default___.Fm0((_dafny.MultiSetOf((_120_v10).Cardinality())).IsProperSubsetOf(_119_v9), _111_v2, _114_v4, _107_globalState)
		_ = _rhs20
		var _rhs21 bool = !(_114_v4) || (_dafny.Companion_Sequence_.Contains((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), _121_v11))
		_ = _rhs21
		var _lhs17 *GlobalState = _107_globalState
		_ = _lhs17
		var _lhs18 *GlobalState = _107_globalState
		_ = _lhs18
		var _lhs19 *GlobalState = _107_globalState
		_ = _lhs19
		_lhs17.F0 = _rhs18
		_114_v4 = _rhs19
		_lhs18.F1 = _rhs20
		_lhs19.F0 = _rhs21
		if !(_114_v4) {
			(_107_globalState).F1 = _dafny.IntOfUint32((_118_v8).Cardinality())
			var _122_v12 _dafny.Map
			_ = _122_v12
			_122_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((_dafny.Zero).Minus(_108_v0), _114_v4, _107_globalState), _114_v4)
			var _123_v13 _dafny.Set
			_ = _123_v13
			_123_v13 = _dafny.SetOf(_114_v4)
			var _rhs22 bool = (func() bool {
				if (_122_v12).Contains(_114_v4) {
					return (_122_v12).Get(_114_v4).(bool)
				}
				return (_dafny.SetOf(_114_v4, !(_114_v4))).IsProperSubsetOf(_123_v13)
			})()
			_ = _rhs22
			var _rhs23 bool = Companion_Default___.Fm1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_111_v2).Cardinality()), _108_v0), _114_v4, _107_globalState)
			_ = _rhs23
			_114_v4 = _rhs22
			_114_v4 = _rhs23
			var _124_v14 bool
			_ = _124_v14
			var _125_v15 _dafny.Set
			_ = _125_v15
			var _out2 bool
			_ = _out2
			var _out3 _dafny.Set
			_ = _out3
			_out2, _out3 = Companion_Default___.M0(((_115_v5).Difference(_115_v5)).Cardinality(), _dafny.IntOfInt64(416), _107_globalState)
			_124_v14 = _out2
			_125_v15 = _out3
			(_107_globalState).F0 = _124_v14
			var _126_v16 _dafny.Array
			_ = _126_v16
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
			_ = _nw26
			_126_v16 = _nw26
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_126_v16), 0))
			_ = _index28
			(_126_v16).ArraySet1(_118_v8, (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_126_v16), 0))
			_ = _index29
			(_126_v16).ArraySet1(_dafny.SeqOf(_114_v4), (_index29).Int())
		} else {
			_121_v11 = _dafny.CodePoint('h')
			var _127_v17 bool
			_ = _127_v17
			_127_v17 = _114_v4
			var _128_v18 _dafny.Map
			_ = _128_v18
			_128_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_127_v17), (_dafny.Zero).Minus(_108_v0))
			_128_v18 = (_128_v18).Update(_114_v4, (_108_v0).Plus(_108_v0))
			(_107_globalState).F0 = _114_v4
			_108_v0 = _108_v0
			var _129_v19 _dafny.MultiSet
			_ = _129_v19
			_129_v19 = _dafny.MultiSetOf(_119_v9)
			(_107_globalState).F0 = !(_129_v19).Contains(_dafny.MultiSetOf(_dafny.IntOfInt64(765)))
		}
		var _130_v21 D1
		_ = _130_v21
		_130_v21 = Companion_D1_.Create_DC3_(false, _dafny.IntOfUint32(((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence)).Cardinality()), _114_v4, _121_v11)
		var _rhs24 _dafny.Int = (Companion_D1_.Create_DC3_(_114_v4, (func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_119_v9).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _131_v20 _dafny.Int
				_131_v20 = interface{}(_compr_0).(_dafny.Int)
				if (_119_v9).Contains(_131_v20) {
					_coll0.Add(Companion_Default___.SafeDivisionInt(_131_v20, _dafny.IntOfUint32(((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())), _114_v4)
				}
			}
			return _coll0.ToMap()
		}()).Cardinality(), (_130_v21).Dtor_cf9(), _121_v11)).Dtor_cf8()
		_ = _rhs24
		var _rhs25 bool = _114_v4
		_ = _rhs25
		var _lhs20 *GlobalState = _107_globalState
		_ = _lhs20
		_108_v0 = _rhs24
		_lhs20.F0 = _rhs25
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
		_ = _index30
		(_109_v1).ArraySet1(_111_v2, (_index30).Int())
		var _132_v22 bool
		_ = _132_v22
		var _133_v23 _dafny.Set
		_ = _133_v23
		var _out4 bool
		_ = _out4
		var _out5 _dafny.Set
		_ = _out5
		_out4, _out5 = Companion_Default___.M0((_dafny.Zero).Minus(_108_v0), _108_v0, _107_globalState)
		_132_v22 = _out4
		_133_v23 = _out5
	} else {
		var _134_v24 _dafny.Array
		_ = _134_v24
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
		_ = _nw27
		_134_v24 = _nw27
		var _135_v25 _dafny.Map
		_ = _135_v25
		_135_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v0, true)
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_134_v24), 0))
		_ = _index31
		(_134_v24).ArraySet1(_135_v25, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_134_v24), 0))
		_ = _index32
		(_134_v24).ArraySet1((func() _dafny.Map {
			if Companion_Default___.Fm1(_108_v0, _114_v4, _107_globalState) {
				return _135_v25
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v0, _114_v4)
		})(), (_index32).Int())
		var _136_v26 _dafny.Sequence
		_ = _136_v26
		_136_v26 = _dafny.SeqOf(_108_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32(((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())))
		var _137_v27 _dafny.MultiSet
		_ = _137_v27
		_137_v27 = _dafny.MultiSetOf(_114_v4)
		if ((func() _dafny.Int {
			if Companion_Default___.Fm1((_dafny.Zero).Minus(_108_v0), _114_v4, _107_globalState) {
				return _108_v0
			}
			return _dafny.IntOfUint32((_136_v26).Cardinality())
		})()).Cmp((_dafny.Zero).Minus(((func() _dafny.Int {
			if (_137_v27).Contains(_114_v4) {
				return (_137_v27).Multiplicity(_114_v4)
			}
			return _108_v0
		})()).Minus(_108_v0))) > 0 {
			var _138_v28 _dafny.Array
			_ = _138_v28
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
			_ = _nw28
			_138_v28 = _nw28
			var _139_v29 _dafny.Sequence
			_ = _139_v29
			_139_v29 = _dafny.SeqOf(_114_v4, _114_v4)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_138_v28), 0))
			_ = _index33
			(_138_v28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_139_v29, _139_v29), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_138_v28), 0))
			_ = _index34
			(_138_v28).ArraySet1(_139_v29, (_index34).Int())
			(_107_globalState).F1 = _dafny.IntOfUint32(((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())
			var _140_v30 D1
			_ = _140_v30
			_140_v30 = Companion_D1_.Create_DC1_(_108_v0)
			_140_v30 = _140_v30
			var _141_v31 bool
			_ = _141_v31
			var _142_v32 _dafny.Set
			_ = _142_v32
			var _out6 bool
			_ = _out6
			var _out7 _dafny.Set
			_ = _out7
			_out6, _out7 = Companion_Default___.M0(_108_v0, Companion_Default___.SafeModuloInt(_108_v0, _108_v0), _107_globalState)
			_141_v31 = _out6
			_142_v32 = _out7
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw29
			_116_v6 = _nw29
		} else {
			var _143_v33 _dafny.CodePoint
			_ = _143_v33
			_143_v33 = _dafny.CodePoint('s')
			_143_v33 = _143_v33
			var _144_v34 D1
			_ = _144_v34
			_144_v34 = Companion_D1_.Create_DC2_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_108_v0), _136_v26), _108_v0, _108_v0, (func() _dafny.Sequence {
				if _114_v4 {
					return _111_v2
				}
				return Companion_Default___.Fm2(true, _114_v4, _108_v0, _114_v4, _107_globalState)
			})(), _115_v5)
			_144_v34 = _144_v34
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_116_v6), 0))
			_ = _index35
			(_116_v6).ArraySet1(_114_v4, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_116_v6), 0))
			_ = _index36
			(_116_v6).ArraySet1(!((func() bool {
				if _114_v4 {
					return _114_v4
				}
				return !(_114_v4)
			})()) || (_114_v4), (_index36).Int())
			var _145_v35 _dafny.Sequence
			_ = _145_v35
			_145_v35 = _dafny.SeqOf(_114_v4, _114_v4)
			var _146_v36 bool
			_ = _146_v36
			var _147_v37 _dafny.Set
			_ = _147_v37
			var _out8 bool
			_ = _out8
			var _out9 _dafny.Set
			_ = _out9
			_out8, _out9 = Companion_Default___.M0(_dafny.IntOfUint32((_145_v35).Cardinality()), (_dafny.IntOfUint32((_111_v2).Cardinality())).Minus(_108_v0), _107_globalState)
			_146_v36 = _out8
			_147_v37 = _out9
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_116_v6), 0))
			_ = _index37
			var _rhs26 bool = (_108_v0).Cmp(_108_v0) != 0
			_ = _rhs26
			var _rhs27 bool = _114_v4
			_ = _rhs27
			var _lhs21 _dafny.Array = _116_v6
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_116_v6), 0))
			_ = _lhs22
			_114_v4 = _rhs26
			(_lhs21).ArraySet1(_rhs27, (_lhs22).Int())
		}
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
		_ = _index38
		(_109_v1).ArraySet1((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), (_index38).Int())
		var _148_v38 _dafny.Array
		_ = _148_v38
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
		_ = _nw30
		_148_v38 = _nw30
		var _149_v39 _dafny.Map
		_ = _149_v39
		_149_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v4, _148_v38)
		_149_v39 = (_149_v39).Update(false, _148_v38)
		var _150_v40 _dafny.Sequence
		_ = _150_v40
		_150_v40 = _dafny.SeqOf(Companion_Default___.Fm1((_136_v26).Select((Companion_Default___.SafeIndex(_108_v0, _dafny.IntOfUint32((_136_v26).Cardinality()))).Uint32()).(_dafny.Int), false, _107_globalState))
		(_107_globalState).F0 = !((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_114_v4, true), _150_v40))).IsDisjointFrom(_137_v27))
	}
	var _151_v41 _dafny.Array
	_ = _151_v41
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(9)
	_ = _len0_3
	var _nw31 _dafny.Array
	_ = _nw31
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw31 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) D1 = (func(_152_v0 _dafny.Int) func(_dafny.Int) D1 {
			return func(_153_i2 _dafny.Int) D1 {
				return Companion_D1_.Create_DC4_(Companion_D1_.Create_DC1_(_152_v0))
			}
		})(_108_v0)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw31 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw31).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw31).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_151_v41 = _nw31
	var _154_v42 _dafny.MultiSet
	_ = _154_v42
	_154_v42 = _dafny.MultiSetOf(_108_v0)
	var _155_v43 _dafny.Map
	_ = _155_v43
	_155_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v4, _154_v42)
	var _156_v44 _dafny.Array
	_ = _156_v44
	var _nwElement0_4 _dafny.Int = _dafny.IntOfUint32((_111_v2).Cardinality())
	_ = _nwElement0_4
	var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(14))
	_ = _nw32
	(_nw32).ArraySet1(_nwElement0_4, 0)
	(_nw32).ArraySet1(_108_v0, 1)
	(_nw32).ArraySet1(_108_v0, 2)
	(_nw32).ArraySet1((_155_v43).Cardinality(), 3)
	(_nw32).ArraySet1(_108_v0, 4)
	(_nw32).ArraySet1(_108_v0, 5)
	(_nw32).ArraySet1(_dafny.IntOfInt64(99), 6)
	(_nw32).ArraySet1(_108_v0, 7)
	(_nw32).ArraySet1(_108_v0, 8)
	(_nw32).ArraySet1(_dafny.IntOfInt64(861), 9)
	(_nw32).ArraySet1(_108_v0, 10)
	(_nw32).ArraySet1(_108_v0, 11)
	(_nw32).ArraySet1(_108_v0, 12)
	(_nw32).ArraySet1(_dafny.IntOfInt64(623), 13)
	_156_v44 = _nw32
	var _157_v45 _dafny.Sequence
	_ = _157_v45
	_157_v45 = _dafny.SeqOf(_156_v44, _156_v44)
	var _158_v46 _dafny.Map
	_ = _158_v46
	_158_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_157_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.IntOfUint32((_157_v45).Cardinality()))).Uint32()).(_dafny.Array), _114_v4)
	var _159_v47 D1
	_ = _159_v47
	_159_v47 = Companion_D1_.Create_DC1_((_158_v46).Cardinality())
	var _160_v48 D1
	_ = _160_v48
	_160_v48 = Companion_D1_.Create_DC4_(_159_v47)
	var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_151_v41), 0))
	_ = _index39
	(_151_v41).ArraySet1(_160_v48, (_index39).Int())
	var _161_v49 _dafny.Array
	_ = _161_v49
	var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
	_ = _nw33
	_161_v49 = _nw33
	var _162_v50 _dafny.Map
	_ = _162_v50
	_162_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(524))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_163_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))).Cardinality()))
	var _164_v51 _dafny.Set
	_ = _164_v51
	_164_v51 = _dafny.SetOf(_114_v4)
	var _pat_let_tv4 = _159_v47
	_ = _pat_let_tv4
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_151_v41), 0))
	_ = _index40
	var _rhs28 _dafny.Int = _108_v0
	_ = _rhs28
	var _rhs29 D1 = func(_pat_let6_0 D1) D1 {
		return func(_165_dt__update__tmp_h0 D1) D1 {
			return func(_pat_let7_0 D1) D1 {
				return func(_166_dt__update_hcf11_h0 D1) D1 {
					return Companion_D1_.Create_DC4_(_166_dt__update_hcf11_h0)
				}(_pat_let7_0)
			}(_pat_let_tv4)
		}(_pat_let6_0)
	}(_160_v48)
	_ = _rhs29
	var _rhs30 bool = Companion_Default___.Fm1(((func() _dafny.Int {
		if (_162_v50).Contains(_114_v4) {
			return (_162_v50).Get(_114_v4).(_dafny.Int)
		}
		return _108_v0
	})()).Times(_108_v0), (func() bool {
		if _114_v4 {
			return _114_v4
		}
		return true
	})(), _107_globalState)
	_ = _rhs30
	var _rhs31 _dafny.Array = _161_v49
	_ = _rhs31
	var _rhs32 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
		if (_154_v42).Contains((_164_v51).Cardinality()) {
			return (_154_v42).Multiplicity((_164_v51).Cardinality())
		}
		return (_dafny.Zero).Minus(_108_v0)
	})())
	_ = _rhs32
	var _lhs23 _dafny.Array = _151_v41
	_ = _lhs23
	var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_151_v41), 0))
	_ = _lhs24
	var _lhs25 *GlobalState = _107_globalState
	_ = _lhs25
	var _lhs26 *GlobalState = _107_globalState
	_ = _lhs26
	_108_v0 = _rhs28
	(_lhs23).ArraySet1(_rhs29, (_lhs24).Int())
	_lhs25.F0 = _rhs30
	_161_v49 = _rhs31
	_lhs26.F1 = _rhs32
	_108_v0 = _108_v0
	for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_116_v6), 0))); ; {
		_guard_loop_0, _ok1 := _iter1()
		if !_ok1 {
			break
		}
		var _167_i4 _dafny.Int
		_167_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_167_i4).Sign() != -1) && ((_167_i4).Cmp(_dafny.ArrayLen((_116_v6), 0)) < 0)) {
			(_116_v6).ArraySet1(_114_v4, (_167_i4).Int())
		}
	}
	var _168_v52 _dafny.CodePoint
	_ = _168_v52
	_168_v52 = _dafny.CodePoint('f')
	var _169_v53 D1
	_ = _169_v53
	_169_v53 = Companion_D1_.Create_DC3_(false, _108_v0, true, _168_v52)
	var _170_v54 _dafny.Set
	_ = _170_v54
	_170_v54 = _dafny.SetOf(_169_v53, _169_v53)
	_170_v54 = _170_v54
	var _171_v55 _dafny.Sequence
	_ = _171_v55
	_171_v55 = _dafny.SeqOf(_108_v0)
	_111_v2 = _dafny.Companion_Sequence_.Update(_111_v2, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((((_162_v50).Update(_114_v4, _dafny.IntOfUint32((_171_v55).Cardinality()))).Merge(_162_v50)).Cardinality()), _dafny.IntOfUint32((_111_v2).Cardinality()))).Uint32(), _168_v52)
	var _source0 D1 = _169_v53
	_ = _source0
	if _source0.Is_DC2() {
		var _172___mcc_h0 _dafny.Sequence = _source0.Get_().(D1_DC2).Cf2
		_ = _172___mcc_h0
		var _173___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC2).Cf3
		_ = _173___mcc_h1
		var _174___mcc_h2 _dafny.Int = _source0.Get_().(D1_DC2).Cf4
		_ = _174___mcc_h2
		var _175___mcc_h3 _dafny.Sequence = _source0.Get_().(D1_DC2).Cf5
		_ = _175___mcc_h3
		var _176___mcc_h4 _dafny.Set = _source0.Get_().(D1_DC2).Cf6
		_ = _176___mcc_h4
		var _177_cf6 _dafny.Set = _176___mcc_h4
		_ = _177_cf6
		var _178_cf5 _dafny.Sequence = _175___mcc_h3
		_ = _178_cf5
		var _179_cf4 _dafny.Int = _174___mcc_h2
		_ = _179_cf4
		var _180_cf3 _dafny.Int = _173___mcc_h1
		_ = _180_cf3
		var _181_cf2 _dafny.Sequence = _172___mcc_h0
		_ = _181_cf2
		_114_v4 = !(false)
		var _182_v56 _dafny.Map
		_ = _182_v56
		_182_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(595), (_dafny.Zero).Minus(_180_cf3))
		_182_v56 = (_182_v56).Update(_dafny.IntOfInt64(526), _108_v0)
		_154_v42 = ((_dafny.MultiSetOf(_180_cf3, _108_v0)).Intersection(_154_v42)).Difference(_154_v42)
		var _183_v57 bool
		_ = _183_v57
		_183_v57 = _114_v4
		var _source1 bool = _183_v57
		_ = _source1
		var _184___mcc_h11 bool = _source1
		_ = _184___mcc_h11
		var _185_cf0 bool = _184___mcc_h11
		_ = _185_cf0
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_116_v6), 0))
		_ = _index41
		(_116_v6).ArraySet1(_114_v4, (_index41).Int())
		var _186_v58 _dafny.Map
		_ = _186_v58
		_186_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v4, _185_cf0)
		var _187_v59 _dafny.Map
		_ = _187_v59
		_187_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_169_v53).Dtor_cf8(), _168_v52)
		var _188_v60 _dafny.MultiSet
		_ = _188_v60
		_188_v60 = _dafny.MultiSetOf((func() _dafny.CodePoint {
			if (_187_v59).Contains(_180_cf3) {
				return (_187_v59).Get(_180_cf3).(_dafny.CodePoint)
			}
			return _168_v52
		})())
		var _189_v61 _dafny.Map
		_ = _189_v61
		_189_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_cf4, _188_v60)
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_116_v6), 0))
		_ = _index42
		(_116_v6).ArraySet1((func() bool {
			if (_186_v58).Contains((_dafny.MultiSetFromSeq(_111_v2)).IsDisjointFrom((func() _dafny.MultiSet {
				if (_189_v61).Contains(_108_v0) {
					return (_189_v61).Get(_108_v0).(_dafny.MultiSet)
				}
				return _188_v60
			})())) {
				return (_186_v58).Get((_dafny.MultiSetFromSeq(_111_v2)).IsDisjointFrom((func() _dafny.MultiSet {
					if (_189_v61).Contains(_108_v0) {
						return (_189_v61).Get(_108_v0).(_dafny.MultiSet)
					}
					return _188_v60
				})())).(bool)
			}
			return _114_v4
		})(), (_index42).Int())
		_178_cf5 = _178_cf5
		(_107_globalState).F0 = _114_v4
		_164_v51 = _164_v51
	} else if _source0.Is_DC3() {
		var _190___mcc_h5 bool = _source0.Get_().(D1_DC3).Cf7
		_ = _190___mcc_h5
		var _191___mcc_h6 _dafny.Int = _source0.Get_().(D1_DC3).Cf8
		_ = _191___mcc_h6
		var _192___mcc_h7 bool = _source0.Get_().(D1_DC3).Cf9
		_ = _192___mcc_h7
		var _193___mcc_h8 _dafny.CodePoint = _source0.Get_().(D1_DC3).Cf10
		_ = _193___mcc_h8
		var _194_cf10 _dafny.CodePoint = _193___mcc_h8
		_ = _194_cf10
		var _195_cf9 bool = _192___mcc_h7
		_ = _195_cf9
		var _196_cf8 _dafny.Int = _191___mcc_h6
		_ = _196_cf8
		var _197_cf7 bool = _190___mcc_h5
		_ = _197_cf7
		_115_v5 = (_115_v5).Difference(_115_v5)
		var _198_v62 bool
		_ = _198_v62
		_198_v62 = _197_cf7
		var _199_v63 _dafny.Map
		_ = _199_v63
		_199_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v0, _198_v62)
		var _200_v64 _dafny.Map
		_ = _200_v64
		_200_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v6, _108_v0)
		var _201_v65 bool
		_ = _201_v65
		var _202_v66 _dafny.Set
		_ = _202_v66
		var _out10 bool
		_ = _out10
		var _out11 _dafny.Set
		_ = _out11
		_out10, _out11 = Companion_Default___.M0(Companion_Default___.SafeModuloInt((_199_v63).Cardinality(), _196_cf8), ((_200_v64).Merge(_200_v64)).Cardinality(), _107_globalState)
		_201_v65 = _out10
		_202_v66 = _out11
		if !(_197_cf7) {
			_156_v44 = _156_v44
			_154_v42 = _154_v42
			(_107_globalState).F1 = _dafny.IntOfUint32((_111_v2).Cardinality())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_156_v44), 0))
			_ = _index43
			(_156_v44).ArraySet1((func() _dafny.Int {
				if _197_cf7 {
					return (_dafny.Zero).Minus(_196_cf8)
				}
				return _dafny.IntOfInt64(-94)
			})(), (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_156_v44), 0))
			_ = _index44
			(_156_v44).ArraySet1(_dafny.IntOfInt64(532), (_index44).Int())
			var _203_v67 _dafny.Map
			_ = _203_v67
			_203_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm3((_156_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_156_v44), 0))).Int()).(_dafny.Int), _108_v0, _107_globalState)), _201_v65)
			_203_v67 = _203_v67
		} else {
			var _204_v68 _dafny.Sequence
			_ = _204_v68
			_204_v68 = _dafny.SeqOf(_201_v65)
			var _205_v69 bool
			_ = _205_v69
			var _206_v70 _dafny.Set
			_ = _206_v70
			var _out12 bool
			_ = _out12
			var _out13 _dafny.Set
			_ = _out13
			_out12, _out13 = Companion_Default___.M0((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_204_v68, _204_v68)).Cardinality())), (_dafny.Zero).Minus(_108_v0), _107_globalState)
			_205_v69 = _out12
			_206_v70 = _out13
			var _207_v71 bool
			_ = _207_v71
			var _208_v72 _dafny.Set
			_ = _208_v72
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Set
			_ = _out15
			_out14, _out15 = Companion_Default___.M0(_dafny.IntOfInt64(320), (_196_cf8).Times(_196_cf8), _107_globalState)
			_207_v71 = _out14
			_208_v72 = _out15
			var _209_v73 _dafny.Array
			_ = _209_v73
			_209_v73 = _156_v44
			var _210_v74 _dafny.MultiSet
			_ = _210_v74
			_210_v74 = _dafny.MultiSetOf(_156_v44, _156_v44, (_209_v73))
			var _rhs33 bool = _197_cf7
			_ = _rhs33
			var _rhs34 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_111_v2, _111_v2)).Cardinality())
			_ = _rhs34
			var _rhs35 bool = (_210_v74).IsSubsetOf((_210_v74).Difference(_210_v74))
			_ = _rhs35
			_195_cf9 = _rhs33
			_196_cf8 = _rhs34
			_197_cf7 = _rhs35
			var _211_v75 *C0
			_ = _211_v75
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__(_dafny.Companion_Sequence_.Update(_111_v2, (Companion_Default___.SafeIndex(_196_cf8, _dafny.IntOfUint32((_111_v2).Cardinality()))).Uint32(), _168_v52), (_108_v0).Times(_108_v0), _dafny.UnicodeSeqOfUtf8Bytes("homss"), _197_cf7)
			_211_v75 = _nw34
			var _212_v76 bool
			_ = _212_v76
			var _213_v77 _dafny.Set
			_ = _213_v77
			var _out16 bool
			_ = _out16
			var _out17 _dafny.Set
			_ = _out17
			_out16, _out17 = Companion_Default___.M0(_108_v0, _dafny.IntOfInt64(90), _107_globalState)
			_212_v76 = _out16
			_213_v77 = _out17
		}
		_195_cf9 = _197_cf7
	} else if _source0.Is_DC1() {
		var _214___mcc_h9 _dafny.Int = _source0.Get_().(D1_DC1).Cf1
		_ = _214___mcc_h9
		var _215_cf1 _dafny.Int = _214___mcc_h9
		_ = _215_cf1
		var _216_v78 _dafny.Map
		_ = _216_v78
		_216_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v6, _114_v4)
		var _217_v79 _dafny.MultiSet
		_ = _217_v79
		_217_v79 = _dafny.MultiSetOf(_114_v4, !((func() bool {
			if (_216_v78).Contains(_116_v6) {
				return (_216_v78).Get(_116_v6).(bool)
			}
			return _114_v4
		})()), _114_v4, _114_v4, _114_v4)
		(_107_globalState).F0 = ((_114_v4) == (_114_v4)) == ((_217_v79).Equals(_217_v79))
		var _218_v81 _dafny.Map
		_ = _218_v81
		_218_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_cf1, _215_cf1)
		var _219_v82 _dafny.Sequence
		_ = _219_v82
		_219_v82 = _dafny.SeqOf(_218_v81)
		(_107_globalState).F1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter2 := _dafny.Iterate((_219_v82).Elements()); ; {
				_compr_1, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _220_v80 _dafny.Map
				_220_v80 = interface{}(_compr_1).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_219_v82, _220_v80) {
					_coll1.Add(_220_v80, _114_v4)
				}
			}
			return _coll1.ToMap()
		}()).Cardinality(), _215_cf1), (_dafny.IntOfInt64(105)).Plus(_108_v0))
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_156_v44), 0))
		_ = _index45
		(_156_v44).ArraySet1(_108_v0, (_index45).Int())
		var _221_v83 _dafny.Map
		_ = _221_v83
		_221_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v2, _114_v4)
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
		_ = _index46
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_156_v44), 0))
		_ = _index47
		var _rhs36 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("s")
		_ = _rhs36
		var _rhs37 _dafny.Int = _108_v0
		_ = _rhs37
		var _rhs38 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_221_v83).Cardinality(), _108_v0), (_108_v0).Minus(_215_cf1))
		_ = _rhs38
		var _rhs39 _dafny.Int = (func() _dafny.Int {
			if _114_v4 {
				return ((_164_v51).Union(_164_v51)).Cardinality()
			}
			return _215_cf1
		})()
		_ = _rhs39
		var _rhs40 _dafny.Int = _108_v0
		_ = _rhs40
		var _lhs27 _dafny.Array = _109_v1
		_ = _lhs27
		var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
		_ = _lhs28
		var _lhs29 _dafny.Array = _156_v44
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_156_v44), 0))
		_ = _lhs30
		var _lhs31 *GlobalState = _107_globalState
		_ = _lhs31
		(_lhs27).ArraySet1(_rhs36, (_lhs28).Int())
		(_lhs29).ArraySet1(_rhs37, (_lhs30).Int())
		_lhs31.F1 = _rhs38
		_108_v0 = _rhs39
		_215_cf1 = _rhs40
		if (_114_v4) == ((func() bool {
			if _114_v4 {
				return _114_v4
			}
			return _114_v4
		})()) {
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
			_ = _index48
			(_109_v1).ArraySet1((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), (_index48).Int())
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
			_ = _index49
			(_109_v1).ArraySet1(_111_v2, (_index49).Int())
			(_107_globalState).F0 = false
			(_107_globalState).F0 = _114_v4
			var _222_v84 _dafny.Map
			_ = _222_v84
			_222_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), _168_v52)
			_222_v84 = (_222_v84).Update((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), _168_v52)
		} else {
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))
			_ = _index50
			(_116_v6).ArraySet1(_114_v4, (_index50).Int())
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))
			_ = _index51
			(_116_v6).ArraySet1(_114_v4, (_index51).Int())
			(_107_globalState).F0 = _114_v4
			var _223_v85 _dafny.Array
			_ = _223_v85
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw35
			_223_v85 = _nw35
			var _224_v86 _dafny.Sequence
			_ = _224_v86
			_224_v86 = _dafny.SeqOf(true, (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
			var _225_v87 _dafny.Map
			_ = _225_v87
			_225_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _224_v86)
			var _226_v88 _dafny.Map
			_ = _226_v88
			_226_v88 = _225_v87
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_223_v85), 0))
			_ = _index52
			(_223_v85).ArraySet1(_226_v88, (_index52).Int())
			var _227_v89 bool
			_ = _227_v89
			_227_v89 = (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
			var _228_v90 T0
			_ = _228_v90
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__(_dafny.SeqOf(_168_v52), Companion_Default___.Fm0(_114_v4, (_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), (_227_v89), _107_globalState), (_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
			_228_v90 = _nw36
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_223_v85), 0))
			_ = _index53
			var _rhs41 _dafny.Map = _226_v88
			_ = _rhs41
			var _rhs42 T0 = _228_v90
			_ = _rhs42
			var _lhs32 _dafny.Array = _223_v85
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_223_v85), 0))
			_ = _lhs33
			(_lhs32).ArraySet1(_rhs41, (_lhs33).Int())
			_228_v90 = _rhs42
			_168_v52 = Companion_Default___.Fm6(_107_globalState)
			var _229_v91 *C0
			_ = _229_v91
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), (_156_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_156_v44), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("lf"), (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
			_229_v91 = _nw37
			var _230_v92 _dafny.Map
			_ = _230_v92
			_230_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v4, _229_v91)
			var _231_v93 _dafny.Array
			_ = _231_v93
			var _nwElement0_5 *C0 = _229_v91
			_ = _nwElement0_5
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(26))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_5, 0)
			(_nw38).ArraySet1(_229_v91, 1)
			(_nw38).ArraySet1(_229_v91, 2)
			(_nw38).ArraySet1(_229_v91, 3)
			(_nw38).ArraySet1(_229_v91, 4)
			(_nw38).ArraySet1(_229_v91, 5)
			(_nw38).ArraySet1(_229_v91, 6)
			(_nw38).ArraySet1(_229_v91, 7)
			(_nw38).ArraySet1(_229_v91, 8)
			(_nw38).ArraySet1(_229_v91, 9)
			(_nw38).ArraySet1(_229_v91, 10)
			(_nw38).ArraySet1(_229_v91, 11)
			(_nw38).ArraySet1(_229_v91, 12)
			(_nw38).ArraySet1(_229_v91, 13)
			(_nw38).ArraySet1((func() *C0 {
				if (_230_v92).Contains((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)) {
					return (_230_v92).Get((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)).(*C0)
				}
				return _229_v91
			})(), 14)
			(_nw38).ArraySet1(_229_v91, 15)
			(_nw38).ArraySet1(_229_v91, 16)
			(_nw38).ArraySet1(_229_v91, 17)
			(_nw38).ArraySet1(_229_v91, 18)
			(_nw38).ArraySet1(_229_v91, 19)
			(_nw38).ArraySet1(_229_v91, 20)
			(_nw38).ArraySet1(_229_v91, 21)
			(_nw38).ArraySet1(_229_v91, 22)
			(_nw38).ArraySet1(_229_v91, 23)
			(_nw38).ArraySet1(_229_v91, 24)
			(_nw38).ArraySet1(_229_v91, 25)
			_231_v93 = _nw38
			var _232_v94 _dafny.Array
			_ = _232_v94
			var _nwElement0_6 _dafny.Array = _231_v93
			_ = _nwElement0_6
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(7))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_6, 0)
			(_nw39).ArraySet1(_231_v93, 1)
			(_nw39).ArraySet1(_231_v93, 2)
			(_nw39).ArraySet1(_231_v93, 3)
			(_nw39).ArraySet1(_231_v93, 4)
			(_nw39).ArraySet1((func() _dafny.Array {
				if _114_v4 {
					return _231_v93
				}
				return _231_v93
			})(), 5)
			(_nw39).ArraySet1(_231_v93, 6)
			_232_v94 = _nw39
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_232_v94), 0))
			_ = _index54
			(_232_v94).ArraySet1(_231_v93, (_index54).Int())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_232_v94), 0))
			_ = _index55
			(_232_v94).ArraySet1(_231_v93, (_index55).Int())
		}
	} else {
		var _233___mcc_h10 D1 = _source0.Get_().(D1_DC4).Cf11
		_ = _233___mcc_h10
		var _234_cf11 D1 = _233___mcc_h10
		_ = _234_cf11
		var _235_v95 *C0
		_ = _235_v95
		var _nw40 *C0 = New_C0_()
		_ = _nw40
		_nw40.Ctor__((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), _108_v0, _111_v2, !(!(_114_v4)))
		_235_v95 = _nw40
		var _236_v96 _dafny.Map
		_ = _236_v96
		_236_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v95, _116_v6)
		_236_v96 = (_236_v96).Update(_235_v95, _116_v6)
		var _237_v97 bool
		_ = _237_v97
		_237_v97 = _114_v4
		(_107_globalState).F0 = (func() bool {
			if _237_v97 {
				return _114_v4
			}
			return _114_v4
		})()
		var _238_v98 _dafny.Sequence
		_ = _238_v98
		_238_v98 = _dafny.SeqOf(!(_114_v4), _114_v4)
		var _239_v99 T0
		_ = _239_v99
		var _nw41 *C0 = New_C0_()
		_ = _nw41
		_nw41.Ctor__((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), (_235_v95).Fm4((_235_v95).F8(), _108_v0, _238_v98, _107_globalState), _dafny.UnicodeSeqOfUtf8Bytes("xalwfr"), _114_v4)
		_239_v99 = _nw41
		_239_v99 = _239_v99
		var _240_v100 _dafny.Map
		_ = _240_v100
		_240_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_239_v99).F5(), _dafny.UnicodeSeqOfUtf8Bytes("le"))).Cardinality()), _114_v4)
		_240_v100 = (_240_v100).Update(Companion_Default___.SafeModuloInt((_235_v95).F8(), (_171_v55).Select((Companion_Default___.SafeIndex((_235_v95).F8(), _dafny.IntOfUint32((_171_v55).Cardinality()))).Uint32()).(_dafny.Int)), _114_v4)
	}
	var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))
	_ = _index56
	(_116_v6).ArraySet1(!(_114_v4), (_index56).Int())
	var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))
	_ = _index57
	(_116_v6).ArraySet1(!(false), (_index57).Int())
	_111_v2 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm2((func() bool {
		if _114_v4 {
			return (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
		}
		return (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
	})(), _114_v4, _108_v0, _114_v4, _107_globalState), (Companion_Default___.SafeIndex((_108_v0).Plus(Companion_Default___.Fm0((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool), _111_v2, Companion_Default___.Fm1(_108_v0, false, _107_globalState), _107_globalState)), _dafny.IntOfUint32((Companion_Default___.Fm2((func() bool {
		if _114_v4 {
			return (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
		}
		return (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
	})(), _114_v4, _108_v0, _114_v4, _107_globalState)).Cardinality()))).Uint32(), _168_v52)
	var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))
	_ = _index58
	(_156_v44).ArraySet1(_dafny.IntOfUint32((_111_v2).Cardinality()), (_index58).Int())
	var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))
	_ = _index59
	(_156_v44).ArraySet1(_dafny.IntOfInt64(82), (_index59).Int())
	var _241_v101 _dafny.Set
	_ = _241_v101
	_241_v101 = _dafny.SetOf(_156_v44, _156_v44)
	var _242_v102 _dafny.Map
	_ = _242_v102
	_242_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v101, (true) || ((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)))
	_242_v102 = (_242_v102).Update(_dafny.SetOf(_156_v44), (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
	var _hi2 _dafny.Int = (_dafny.Zero).Minus(_108_v0)
	_ = _hi2
	for _243_i5 := _108_v0; _243_i5.Cmp(_hi2) < 0; _243_i5 = _243_i5.Plus(_dafny.One) {
		var _244_v103 *C0
		_ = _244_v103
		var _nw42 *C0 = New_C0_()
		_ = _nw42
		_nw42.Ctor__(_111_v2, Companion_Default___.SafeDivisionInt(_108_v0, (_156_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.Concatenate((_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence), (func() _dafny.Sequence {
			if (_112_v3).Contains(!((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))) {
				return (_112_v3).Get(!((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(902))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_245_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_246_i6 _dafny.Int) _dafny.CodePoint {
					return _245_v52
				}
			})(_168_v52)))
		})()), (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
		_244_v103 = _nw42
		if (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool) {
			var _247_v105 _dafny.MultiSet
			_ = _247_v105
			_247_v105 = _dafny.MultiSetOf((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool), _114_v4)
			var _248_v106 _dafny.Map
			_ = _248_v106
			_248_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(79), _dafny.IntOfInt64(-6))); ; {
					_compr_2, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _249_v104 _dafny.Int
					_249_v104 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(79)).Cmp(_249_v104) <= 0) && ((_249_v104).Cmp(_dafny.IntOfInt64(-6)) < 0) {
						_coll2.Add((_249_v104).Plus(_dafny.IntOfUint32(((_244_v103).F7()).Cardinality())), (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
					}
				}
				return _coll2.ToMap()
			}()).Cardinality(), _247_v105)
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
			_ = _index60
			var _rhs43 bool = (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
			_ = _rhs43
			var _rhs44 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("tkbtcsqb")
			_ = _rhs44
			var _rhs45 _dafny.Int = _108_v0
			_ = _rhs45
			var _rhs46 bool = !(!((func() _dafny.Map {
				if (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool) {
					return _248_v106
				}
				return _248_v106
			})()).Contains(_dafny.IntOfInt64(2)))
			_ = _rhs46
			var _rhs47 bool = !((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
			_ = _rhs47
			var _lhs34 _dafny.Array = _109_v1
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
			_ = _lhs35
			var _lhs36 *GlobalState = _107_globalState
			_ = _lhs36
			var _lhs37 *GlobalState = _107_globalState
			_ = _lhs37
			_114_v4 = _rhs43
			(_lhs34).ArraySet1(_rhs44, (_lhs35).Int())
			_lhs36.F1 = _rhs45
			_lhs37.F0 = _rhs46
			_114_v4 = _rhs47
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
			_ = _index61
			var _rhs48 _dafny.Sequence = (_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence)
			_ = _rhs48
			var _rhs49 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ilf"), (func() _dafny.Sequence {
				if true {
					return _111_v2
				}
				return (_109_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))).Int()).(_dafny.Sequence)
			})())
			_ = _rhs49
			var _rhs50 bool = (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
			_ = _rhs50
			var _lhs38 _dafny.Array = _109_v1
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_109_v1), 0))
			_ = _lhs39
			var _lhs40 *GlobalState = _107_globalState
			_ = _lhs40
			_111_v2 = _rhs48
			(_lhs38).ArraySet1(_rhs49, (_lhs39).Int())
			_lhs40.F0 = _rhs50
			var _250_v107 _dafny.Map
			_ = _250_v107
			_250_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v4, _244_v103)
			var _251_v108 _dafny.Sequence
			_ = _251_v108
			_251_v108 = _dafny.SeqOf(_244_v103)
			var _252_v110 _dafny.Sequence
			_ = _252_v110
			_252_v110 = _dafny.SeqOf(_115_v5, _115_v5, func() _dafny.Set {
				var _coll3 = _dafny.NewBuilder()
				_ = _coll3
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(172), _dafny.IntOfInt64(-314))); ; {
					_compr_3, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _253_v109 _dafny.Int
					_253_v109 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(172)).Cmp(_253_v109) <= 0) && ((_253_v109).Cmp(_dafny.IntOfInt64(-314)) < 0) {
						_coll3.Add(Companion_Default___.SafeDivisionInt(_253_v109, _108_v0))
					}
				}
				return _coll3.ToSet()
			}(), _115_v5, _115_v5)
			var _254_v111 _dafny.Map
			_ = _254_v111
			_254_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_i5, (_115_v5).IsProperSubsetOf((_252_v110).Select((Companion_Default___.SafeIndex(_108_v0, _dafny.IntOfUint32((_252_v110).Cardinality()))).Uint32()).(_dafny.Set)))
			var _255_v112 _dafny.Sequence
			_ = _255_v112
			_255_v112 = _dafny.SeqOf(_114_v4, _114_v4, (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool), !(_114_v4))
			var _rhs51 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_114_v4) || (_114_v4), _244_v103)
			_ = _rhs51
			var _rhs52 _dafny.CodePoint = ((_244_v103).F7()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_171_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_251_v108).Cardinality()), _dafny.IntOfUint32((_171_v55).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32(((_244_v103).F7()).Cardinality()), (func() _dafny.Int {
				if (_247_v105).Contains((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)) {
					return (_247_v105).Multiplicity((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
				}
				return (Companion_Default___.Fm7((_244_v103).F8(), _107_globalState)).Cardinality()
			})(), _243_i5)).Cardinality()), _dafny.IntOfUint32(((_244_v103).F7()).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs52
			var _rhs53 bool = (func() bool {
				if (_254_v111).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_255_v112, _255_v112)).Cardinality())) {
					return (_254_v111).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_255_v112, _255_v112)).Cardinality())).(bool)
				}
				return _114_v4
			})()
			_ = _rhs53
			var _lhs41 *GlobalState = _107_globalState
			_ = _lhs41
			_250_v107 = _rhs51
			_168_v52 = _rhs52
			_lhs41.F0 = _rhs53
			var _256_v113 _dafny.Array
			_ = _256_v113
			var _nwElement0_7 _dafny.Array = _116_v6
			_ = _nwElement0_7
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(19))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_7, 0)
			(_nw43).ArraySet1(_116_v6, 1)
			(_nw43).ArraySet1(_116_v6, 2)
			(_nw43).ArraySet1(_116_v6, 3)
			(_nw43).ArraySet1(_116_v6, 4)
			(_nw43).ArraySet1(_116_v6, 5)
			(_nw43).ArraySet1(_116_v6, 6)
			(_nw43).ArraySet1(_116_v6, 7)
			(_nw43).ArraySet1((Companion_D4_.Create_DC7_(_116_v6)).Dtor_cf14(), 8)
			(_nw43).ArraySet1(_116_v6, 9)
			(_nw43).ArraySet1(_116_v6, 10)
			(_nw43).ArraySet1(_116_v6, 11)
			(_nw43).ArraySet1(_116_v6, 12)
			(_nw43).ArraySet1(_116_v6, 13)
			(_nw43).ArraySet1(_116_v6, 14)
			(_nw43).ArraySet1(_116_v6, 15)
			(_nw43).ArraySet1(_116_v6, 16)
			(_nw43).ArraySet1(_116_v6, 17)
			(_nw43).ArraySet1(_116_v6, 18)
			_256_v113 = _nw43
			_256_v113 = _256_v113
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))
			_ = _index62
			(_156_v44).ArraySet1((_244_v103).F8(), (_index62).Int())
		} else {
			var _257_v114 _dafny.Sequence
			_ = _257_v114
			_257_v114 = _157_v45
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))
			_ = _index63
			var _rhs54 _dafny.Array = _156_v44
			_ = _rhs54
			var _rhs55 *C0 = _244_v103
			_ = _rhs55
			var _rhs56 bool = !((_115_v5).Equals(_115_v5))
			_ = _rhs56
			var _rhs57 bool = _114_v4
			_ = _rhs57
			var _rhs58 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mxtsyoj")).Cardinality())).Minus((_dafny.IntOfUint32((_171_v55).Cardinality())).Plus(_dafny.IntOfUint32((_257_v114).Cardinality())))
			_ = _rhs58
			var _lhs42 *GlobalState = _107_globalState
			_ = _lhs42
			var _lhs43 *GlobalState = _107_globalState
			_ = _lhs43
			var _lhs44 _dafny.Array = _156_v44
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))
			_ = _lhs45
			_156_v44 = _rhs54
			_244_v103 = _rhs55
			_lhs42.F0 = _rhs56
			_lhs43.F0 = _rhs57
			(_lhs44).ArraySet1(_rhs58, (_lhs45).Int())
			var _258_v115 *C0
			_ = _258_v115
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__((_244_v103).F7(), _243_i5, _111_v2, (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))
			_258_v115 = _nw44
			var _259_v116 D6
			_ = _259_v116
			_259_v116 = Companion_D6_.Create_DC12_(_154_v42)
			var _260_v117 _dafny.MultiSet
			_ = _260_v117
			_260_v117 = _dafny.MultiSetOf(_114_v4, ((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)) && (_114_v4), (func() bool {
				if (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool) {
					return (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)
				}
				return _114_v4
			})(), ((_259_v116).Dtor_cf25()).IsDisjointFrom(_154_v42), !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool), (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool))).Contains((_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool)))
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))
			_ = _index64
			(_156_v44).ArraySet1((_260_v117).Cardinality(), (_index64).Int())
			var _261_v118 _dafny.Array
			_ = _261_v118
			var _nwElement0_8 _dafny.Array = _116_v6
			_ = _nwElement0_8
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(8))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_8, 0)
			(_nw45).ArraySet1(_116_v6, 1)
			(_nw45).ArraySet1(_116_v6, 2)
			(_nw45).ArraySet1(_116_v6, 3)
			(_nw45).ArraySet1(_116_v6, 4)
			(_nw45).ArraySet1(_116_v6, 5)
			(_nw45).ArraySet1(_116_v6, 6)
			(_nw45).ArraySet1((func() _dafny.Array {
				if _114_v4 {
					return _116_v6
				}
				return _116_v6
			})(), 7)
			_261_v118 = _nw45
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_261_v118), 0))
			_ = _index65
			(_261_v118).ArraySet1(_116_v6, (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_261_v118), 0))
			_ = _index66
			(_261_v118).ArraySet1(_116_v6, (_index66).Int())
			var _262_v119 _dafny.Array
			_ = _262_v119
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
			_ = _nw46
			_262_v119 = _nw46
			var _263_v121 _dafny.MultiSet
			_ = _263_v121
			_263_v121 = _dafny.MultiSetOf(_111_v2, (_244_v103).F7())
			var _264_v122 D6
			_ = _264_v122
			_264_v122 = Companion_D6_.Create_DC15_((_244_v103).F8(), _258_v115, func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter5 := _dafny.Iterate((_263_v121).Elements()); ; {
					_compr_4, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _265_v120 _dafny.Sequence
					_265_v120 = interface{}(_compr_4).(_dafny.Sequence)
					if (_263_v121).Contains(_265_v120) {
						_coll4.Add(_265_v120, _108_v0)
					}
				}
				return _coll4.ToMap()
			}(), Companion_Default___.Fm1((_260_v117).Cardinality(), (_116_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_116_v6), 0))).Int()).(bool), _107_globalState))
			var _266_v123 D6
			_ = _266_v123
			_266_v123 = Companion_D6_.Create_DC16_(_264_v122)
			var _267_v124 _dafny.Map
			_ = _267_v124
			_267_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v123, _dafny.IntOfUint32(((_244_v103).F7()).Cardinality()))
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_262_v119), 0))
			_ = _index67
			(_262_v119).ArraySet1((_267_v124).Merge(_267_v124), (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_262_v119), 0))
			_ = _index68
			(_262_v119).ArraySet1((_267_v124).Merge((_267_v124).Merge(_267_v124)), (_index68).Int())
		}
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_156_v44), 0))
		_ = _index69
		(_156_v44).ArraySet1(_243_i5, (_index69).Int())
		(_107_globalState).F0 = _114_v4
	}
	_dafny.Print(_107_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_107_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_107_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_107_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_107_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_108_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v1).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_111_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_112_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("eedslel"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_115_v5).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v6).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_117_v7).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_151_v41).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D1)).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v42).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v43).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v44).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_157_v45).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v46).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v47).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_160_v48).Dtor_cf11()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_v50).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(524))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v51).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_168_v52)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v53).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v53).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v53).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v53).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v54).Equals(_dafny.SetOf(Companion_D1_.Create_DC3_(false, _dafny.One, true, _dafny.CodePoint('f')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_171_v55, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v101).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v102).Cardinality())
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
	Cf2 _dafny.Sequence
	Cf3 _dafny.Int
	Cf4 _dafny.Int
	Cf5 _dafny.Sequence
	Cf6 _dafny.Set
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Sequence, Cf3 _dafny.Int, Cf4 _dafny.Int, Cf5 _dafny.Sequence, Cf6 _dafny.Set) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4, Cf5, Cf6}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf7  bool
	Cf8  _dafny.Int
	Cf9  bool
	Cf10 _dafny.CodePoint
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 bool, Cf8 _dafny.Int, Cf9 bool, Cf10 _dafny.CodePoint) D1 {
	return D1{D1_DC3{Cf7, Cf8, Cf9, Cf10}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Int
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Int) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

type D1_DC4 struct {
	Cf11 D1
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf11 D1) D1 {
	return D1{D1_DC4{Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero, _dafny.EmptySeq, _dafny.EmptySet)
}

func (_this D1) Dtor_cf2() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Set {
	return _this.Get_().(D1_DC2).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC3).Cf9
}

func (_this D1) Dtor_cf10() _dafny.CodePoint {
	return _this.Get_().(D1_DC3).Cf10
}

func (_this D1) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf11() D1 {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + data.Cf5.VerbatimString(true) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf2.Equals(data2.Cf2) && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Equals(data2.Cf5) && data1.Cf6.Equals(data2.Cf6)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf12 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf12 _dafny.Array) D2 {
	return D2{D2_DC5{Cf12}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf12() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf12 == data2.Cf12
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
	Cf13 _dafny.Map
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf13 _dafny.Map) D3 {
	return D3{D3_DC6{Cf13}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf13() _dafny.Map {
	return _this.Get_().(D3_DC6).Cf13
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf15 _dafny.Int
	Cf16 T0
	Cf17 bool
	Cf18 bool
	Cf19 *C0
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf15 _dafny.Int, Cf16 T0, Cf17 bool, Cf18 bool, Cf19 *C0) D4 {
	return D4{D4_DC8{Cf15, Cf16, Cf17, Cf18, Cf19}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC9 struct {
	Cf20 bool
	Cf21 _dafny.Int
	Cf22 T0
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf20 bool, Cf21 _dafny.Int, Cf22 T0) D4 {
	return D4{D4_DC9{Cf20, Cf21, Cf22}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC7 struct {
	Cf14 _dafny.Array
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf14 _dafny.Array) D4 {
	return D4{D4_DC7{Cf14}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

type D4_DC10 struct {
	Cf23 D4
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf23 D4) D4 {
	return D4{D4_DC10{Cf23}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC8_(_dafny.Zero, (T0)(nil), false, false, (*C0)(nil))
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC8).Cf15
}

func (_this D4) Dtor_cf16() T0 {
	return _this.Get_().(D4_DC8).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC8).Cf17
}

func (_this D4) Dtor_cf18() bool {
	return _this.Get_().(D4_DC8).Cf18
}

func (_this D4) Dtor_cf19() *C0 {
	return _this.Get_().(D4_DC8).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC9).Cf20
}

func (_this D4) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf21
}

func (_this D4) Dtor_cf22() T0 {
	return _this.Get_().(D4_DC9).Cf22
}

func (_this D4) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D4_DC7).Cf14
}

func (_this D4) Dtor_cf23() D4 {
	return _this.Get_().(D4_DC10).Cf23
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && _dafny.AreEqual(data1.Cf16, data2.Cf16) && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Cmp(data2.Cf21) == 0 && _dafny.AreEqual(data1.Cf22, data2.Cf22)
		}
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf14 == data2.Cf14
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf23.Equals(data2.Cf23)
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
	Cf24 _dafny.Sequence
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf24 _dafny.Sequence) D5 {
	return D5{D5_DC11{Cf24}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D5) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D5_DC11).Cf24
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf26 bool
	Cf27 bool
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf26 bool, Cf27 bool) D6 {
	return D6{D6_DC13{Cf26, Cf27}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC14 struct {
	Cf28 bool
	Cf29 bool
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf28 bool, Cf29 bool) D6 {
	return D6{D6_DC14{Cf28, Cf29}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC15 struct {
	Cf30 _dafny.Int
	Cf31 *C0
	Cf32 _dafny.Map
	Cf33 bool
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf30 _dafny.Int, Cf31 *C0, Cf32 _dafny.Map, Cf33 bool) D6 {
	return D6{D6_DC15{Cf30, Cf31, Cf32, Cf33}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC12 struct {
	Cf25 _dafny.MultiSet
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf25 _dafny.MultiSet) D6 {
	return D6{D6_DC12{Cf25}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC16 struct {
	Cf34 D6
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf34 D6) D6 {
	return D6{D6_DC16{Cf34}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC13_(false, false)
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC13).Cf26
}

func (_this D6) Dtor_cf27() bool {
	return _this.Get_().(D6_DC13).Cf27
}

func (_this D6) Dtor_cf28() bool {
	return _this.Get_().(D6_DC14).Cf28
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC14).Cf29
}

func (_this D6) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf30
}

func (_this D6) Dtor_cf31() *C0 {
	return _this.Get_().(D6_DC15).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D6_DC15).Cf32
}

func (_this D6) Dtor_cf33() bool {
	return _this.Get_().(D6_DC15).Cf33
}

func (_this D6) Dtor_cf25() _dafny.MultiSet {
	return _this.Get_().(D6_DC12).Cf25
}

func (_this D6) Dtor_cf34() D6 {
	return _this.Get_().(D6_DC16).Cf34
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf34) + ")"
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
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31 == data2.Cf31 && data1.Cf32.Equals(data2.Cf32) && data1.Cf33 == data2.Cf33
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf34.Equals(data2.Cf34)
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
	F5() _dafny.Sequence
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
	F5() _dafny.Sequence
	F6() bool
	F6_set_(value bool)
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
	F0  bool
	F1  _dafny.Int
	_f2 _dafny.CodePoint
	_f3 _dafny.Int
	_f4 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F1 = _dafny.Zero
	_this._f2 = _dafny.CodePoint('D')
	_this._f3 = _dafny.Zero
	_this._f4 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.CodePoint, f3 _dafny.Int, f4 bool) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *GlobalState) F2() _dafny.CodePoint {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f6 bool
	_f5 _dafny.Sequence
	_f7 _dafny.Sequence
	_f8 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f6 = false
	_this._f5 = _dafny.EmptySeq
	_this._f7 = _dafny.EmptySeq
	_this._f8 = _dafny.Zero
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

func (_this *C0) F6() bool {
	return _this._f6
}
func (_this *C0) F6_set_(value bool) {
	_this._f6 = value
}
func (_this *C0) F5() _dafny.Sequence {
	return _this._f5
}
func (_this *C0) Ctor__(f7 _dafny.Sequence, f8 _dafny.Int, f5 _dafny.Sequence, f6 bool) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C0) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F8(), (func() _dafny.Int {
			if _this.F6() {
				return (_this).F8()
			}
			return (_this).F8()
		})())
	}
}
func (_this *C0) Fm5(p0 bool, p1 D1, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _dafny.SeqOf(_this.F6())))
	}
}
func (_this *C0) F7() _dafny.Sequence {
	{
		return _this._f7
	}
}
func (_this *C0) F8() _dafny.Int {
	{
		return _this._f8
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
