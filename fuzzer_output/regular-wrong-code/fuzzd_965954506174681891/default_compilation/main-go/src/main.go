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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool {
	return ((_dafny.IntOfUint32((_dafny.SeqOf(true, !(false))).Cardinality())).Times(_dafny.IntOfInt64(840))).Cmp((_dafny.IntOfInt64(-513)).Plus(_dafny.IntOfInt64(882))) < 0
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(736)
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("euwd"), _dafny.UnicodeSeqOfUtf8Bytes("d"))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-211))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})), _dafny.UnicodeSeqOfUtf8Bytes("bmpbyeb"), _dafny.UnicodeSeqOfUtf8Bytes("jeeg"))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Sequence
			_2_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-211))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})), _dafny.UnicodeSeqOfUtf8Bytes("bmpbyeb"), _dafny.UnicodeSeqOfUtf8Bytes("jeeg"))).Contains(_2_v0) {
				_coll0.Add(_2_v0)
			}
		}
		return _coll0.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(240))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_3_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_4_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))
		}))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v0 _dafny.Sequence
			_5_v0 = interface{}(_compr_1).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(240))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_3_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_4_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}))
			})), _5_v0) {
				_coll1.Add(_5_v0)
			}
		}
		return _coll1.ToSet()
	}()).Intersection(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("khuk"), _dafny.UnicodeSeqOfUtf8Bytes("gb"), _dafny.UnicodeSeqOfUtf8Bytes("gbthcxlck"))))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(582), (Companion_D7_.Create_DC17_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(121), _dafny.IntOfInt64(901))).Cardinality(), false, true)).Dtor_cf29(), _dafny.IntOfInt64(-450)), _dafny.SeqOf(_dafny.IntOfInt64(630), _dafny.IntOfInt64(714))), _dafny.SeqOf(_dafny.IntOfInt64(549)))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.IntOfInt64(528)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (Companion_D7_.Create_DC17_(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(706), _dafny.IntOfInt64(932))).Cardinality()), true, !(true))).Dtor_cf30())).Cardinality()) != 0, ((_dafny.MultiSetOf((func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(469), _dafny.IntOfInt64(-47))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _6_v0 _dafny.Int
			_6_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(469)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(-47)) < 0) {
				_coll2.Add((_6_v0).Minus(_dafny.IntOfInt64(437)))
			}
		}
		return _coll2.ToSet()
	}()).Cardinality(), (_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))).Cardinality(), _dafny.IntOfInt64(581), (_dafny.Zero).Minus(_dafny.IntOfInt64(-48)))).Cardinality()).Cmp(_dafny.IntOfInt64(561)) >= 0, false, !((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(!(true), !(true)), _dafny.SeqOf(true, true, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(!(true), true, true), _dafny.SeqOf(true, false, true))).Cardinality())).Cmp((_dafny.MultiSetOf(false, false)).Cardinality()) == 0), false)
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	if (true) || (!(true)) {
		return _dafny.CodePoint('j')
	} else {
		return _dafny.CodePoint('u')
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(65))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})), _dafny.UnicodeSeqOfUtf8Bytes("xm"))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Difference(_dafny.SetOf(true, false))).Union((_dafny.SetOf(true, false)).Difference(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(986), (_dafny.MultiSetOf(true)).Cardinality())).Cardinality())).Cardinality()), (_dafny.SetOf(_dafny.CodePoint('k'))).Cardinality()))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-994), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-991))).Cardinality())).Cardinality()))), Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-577), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(660), _dafny.IntOfInt64(407)))).Cardinality(), _dafny.IntOfInt64(584)))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var _8_v0 _dafny.Int
	_ = _8_v0
	_8_v0 = _dafny.IntOfInt64(301)
	var _9_v1 bool
	_ = _9_v1
	_9_v1 = false
	var _10_v2 _dafny.Map
	_ = _10_v2
	_10_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, _dafny.CodePoint('l'))
	var _hi0 _dafny.Int = (_dafny.Zero).Minus(((_10_v2).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, _dafny.CodePoint('k')))).Cardinality())
	_ = _hi0
	for _11_i0 := _8_v0; _11_i0.Cmp(_hi0) < 0; _11_i0 = _11_i0.Plus(_dafny.One) {
		r1 = _8_v0
		_8_v0 = _8_v0
		if _9_v1 {
			var _12_v3 _dafny.Sequence
			_ = _12_v3
			_12_v3 = _dafny.UnicodeSeqOfUtf8Bytes("r")
			var _13_v4 _dafny.CodePoint
			_ = _13_v4
			_13_v4 = _dafny.CodePoint('l')
			var _rhs0 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_9_v1, _dafny.Companion_Sequence_.IsProperPrefixOf(_12_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(196))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_14_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))), _9_v1, _9_v1, (_9_v1) && (Companion_Default___.Fm0(_9_v1, _9_v1, _12_v3, globalState))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(267))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_15_i0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_16_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-835))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}((func(_17_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_18_i3 _dafny.Int) _dafny.Int {
							return _17_i0
						}
					})(_15_i0)))
				}
			})(_11_i0)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_9_v1, _dafny.Companion_Sequence_.IsProperPrefixOf(_12_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(196))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_19_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))), _9_v1, _9_v1, (_9_v1) && (Companion_Default___.Fm0(_9_v1, _9_v1, _12_v3, globalState)))).Cardinality()))).Uint32(), (_9_v1) && (_9_v1))
			_ = _rhs0
			var _rhs1 _dafny.Int = _11_i0
			_ = _rhs1
			var _rhs2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("quhelj"), _dafny.Companion_Sequence_.Update(_12_v3, (Companion_Default___.SafeIndex(_8_v0, _dafny.IntOfUint32((_12_v3).Cardinality()))).Uint32(), _13_v4)), _12_v3)
			_ = _rhs2
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			var _lhs1 *GlobalState = globalState
			_ = _lhs1
			_lhs0.F9 = _rhs0
			_lhs1.F7 = _rhs1
			_12_v3 = _rhs2
			var _20_v5 _dafny.Array
			_ = _20_v5
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_0
			var _nw0 _dafny.Array
			_ = _nw0
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw0 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_21_i0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_22_i4 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_21_i0), _dafny.SeqOf(_21_i0))
					}
				})(_11_i0)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw0).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw0).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_20_v5 = _nw0
			var _23_v6 _dafny.Array
			_ = _23_v6
			var _nw1 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw1
			_23_v6 = _nw1
			var _24_v7 _dafny.Array
			_ = _24_v7
			var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw2
			_24_v7 = _nw2
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_24_v7), 0))
			_ = _index0
			(_24_v7).ArraySet1(_23_v6, (_index0).Int())
			var _25_v8 _dafny.Map
			_ = _25_v8
			_25_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, _23_v6)
			var _26_v9 D5
			_ = _26_v9
			_26_v9 = Companion_D5_.Create_DC11_(_9_v1, _11_i0, true, _23_v6, _dafny.IntOfInt64(319))
			var _27_v10 _dafny.Map
			_ = _27_v10
			_27_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_12_v3).Select((Companion_Default___.SafeIndex(_8_v0, _dafny.IntOfUint32((_12_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), _12_v3)
			var _28_v11 _dafny.Set
			_ = _28_v11
			_28_v11 = _dafny.SetOf(Companion_Default___.Fm1(false, _8_v0, (_8_v0), globalState))
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_24_v7), 0))
			_ = _index1
			var _rhs3 _dafny.Array = _20_v5
			_ = _rhs3
			var _rhs4 _dafny.Array = (func() _dafny.Array {
				if (_25_v8).Contains(!(_9_v1)) {
					return (_25_v8).Get(!(_9_v1)).(_dafny.Array)
				}
				return _23_v6
			})()
			_ = _rhs4
			var _rhs5 _dafny.Array = (_26_v9).Dtor_cf14()
			_ = _rhs5
			var _rhs6 bool = ((_27_v10).Cardinality()).Cmp(Companion_Default___.SafeModuloInt(_8_v0, (_28_v11).Cardinality())) > 0
			_ = _rhs6
			var _lhs2 _dafny.Array = _24_v7
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_24_v7), 0))
			_ = _lhs3
			_20_v5 = _rhs3
			_23_v6 = _rhs4
			(_lhs2).ArraySet1(_rhs5, (_lhs3).Int())
			_9_v1 = _rhs6
			_8_v0 = _11_i0
			var _29_v12 _dafny.Array
			_ = _29_v12
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw3
			_29_v12 = _nw3
			var _30_v13 _dafny.Sequence
			_ = _30_v13
			_30_v13 = _dafny.SeqOf(_8_v0)
			var _31_v14 _dafny.Map
			_ = _31_v14
			_31_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, _9_v1)
			var _32_v15 *C0
			_ = _32_v15
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__(_9_v1, _31_v14)
			_32_v15 = _nw4
			var _33_v16 _dafny.Sequence
			_ = _33_v16
			_33_v16 = _dafny.SeqOf(_dafny.IntOfUint32((_30_v13).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v15, _8_v0)).Cardinality())
			var _34_v17 _dafny.Map
			_ = _34_v17
			_34_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_30_v13).Cardinality()), _9_v1)
			var _35_v18 _dafny.MultiSet
			_ = _35_v18
			_35_v18 = _dafny.MultiSetOf(_11_i0, _8_v0)
			var _36_v19 _dafny.Sequence
			_ = _36_v19
			_36_v19 = _dafny.SeqOf(_9_v1, (func() bool {
				if (_34_v17).Contains((func() _dafny.Int {
					if (_35_v18).Contains((_dafny.Zero).Minus(_8_v0)) {
						return (_35_v18).Multiplicity((_dafny.Zero).Minus(_8_v0))
					}
					return _8_v0
				})()) {
					return (_34_v17).Get((func() _dafny.Int {
						if (_35_v18).Contains((_dafny.Zero).Minus(_8_v0)) {
							return (_35_v18).Multiplicity((_dafny.Zero).Minus(_8_v0))
						}
						return _8_v0
					})()).(bool)
				}
				return _32_v15.F19
			})(), _32_v15.F19)
			var _37_v20 _dafny.Map
			_ = _37_v20
			_37_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_29_v12, _36_v19)
			_37_v20 = _37_v20
			_13_v4 = _13_v4
		} else {
			(globalState).F8 = _8_v0
			var _rhs7 _dafny.Int = (_8_v0).Plus(_8_v0)
			_ = _rhs7
			var _rhs8 bool = _9_v1
			_ = _rhs8
			r1 = _rhs7
			_9_v1 = _rhs8
			var _38_v21 _dafny.Array
			_ = _38_v21
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw5
			_38_v21 = _nw5
			_38_v21 = _38_v21
			var _39_v22 _dafny.Sequence
			_ = _39_v22
			_39_v22 = _dafny.UnicodeSeqOfUtf8Bytes("xbhg")
			var _40_v23 *C0
			_ = _40_v23
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__(_9_v1, Companion_Default___.Fm9(_8_v0, _dafny.IntOfUint32((_39_v22).Cardinality()), Companion_Default___.Fm1(_9_v1, _11_i0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("q"), _8_v0)).Cardinality(), globalState), globalState))
			_40_v23 = _nw6
			var _41_v24 _dafny.CodePoint
			_ = _41_v24
			_41_v24 = _dafny.CodePoint('y')
			var _42_v25 _dafny.Map
			_ = _42_v25
			_42_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, _dafny.IntOfInt64(770))
			var _rhs9 bool = Companion_Default___.Fm0(_40_v23.F19, (_8_v0).Cmp(_8_v0) >= 0, _dafny.Companion_Sequence_.Update(_39_v22, (Companion_Default___.SafeIndex(_11_i0, _dafny.IntOfUint32((_39_v22).Cardinality()))).Uint32(), _41_v24), globalState)
			_ = _rhs9
			var _rhs10 *C0 = _40_v23
			_ = _rhs10
			var _rhs11 _dafny.Int = (_dafny.Zero).Minus((_11_i0).Plus((_42_v25).Cardinality()))
			_ = _rhs11
			var _rhs12 _dafny.Sequence = _39_v22
			_ = _rhs12
			var _rhs13 bool = (func() bool {
				if _9_v1 {
					return _40_v23.F19
				}
				return _9_v1
			})()
			_ = _rhs13
			_9_v1 = _rhs9
			_40_v23 = _rhs10
			r0 = _rhs11
			_39_v22 = _rhs12
			_9_v1 = _rhs13
			_8_v0 = _11_i0
		}
		var _43_v26 _dafny.Array
		_ = _43_v26
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
		_ = _nw7
		_43_v26 = _nw7
		var _44_v27 _dafny.Sequence
		_ = _44_v27
		_44_v27 = _dafny.SeqOf(_43_v26, _43_v26, (func() _dafny.Array {
			if _9_v1 {
				return _43_v26
			}
			return _43_v26
		})())
		_43_v26 = (_44_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.IntOfUint32((_44_v27).Cardinality()))).Uint32()).(_dafny.Array)
	}
	r0 = _8_v0
	var _45_v28 _dafny.CodePoint
	_ = _45_v28
	_45_v28 = _dafny.CodePoint('f')
	var _46_v29 _dafny.Sequence
	_ = _46_v29
	_46_v29 = _dafny.SeqOf(_45_v28, _45_v28, _45_v28)
	var _47_v30 _dafny.Sequence
	_ = _47_v30
	_47_v30 = _dafny.SeqOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_46_v29, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(373))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}((func(_48_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_49_i5 _dafny.Int) _dafny.CodePoint {
			return _48_v28
		}
	})(_45_v28)))), _9_v1)
	var _50_v31 _dafny.Map
	_ = _50_v31
	_50_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}((func(_51_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_52_i6 _dafny.Int) _dafny.Int {
			return _51_v0
		}
	})(_8_v0)))).Cardinality()))
	var _rhs14 _dafny.Int = ((func() _dafny.Int {
		if !(_9_v1) {
			return _dafny.IntOfInt64(12)
		}
		return _8_v0
	})()).Times(_8_v0)
	_ = _rhs14
	var _rhs15 bool = (_47_v30).Select((Companion_Default___.SafeIndex((_8_v0).Times((func() _dafny.Int {
		if (_50_v31).Contains(_dafny.IntOfUint32((_46_v29).Cardinality())) {
			return (_50_v31).Get(_dafny.IntOfUint32((_46_v29).Cardinality())).(_dafny.Int)
		}
		return Companion_Default___.Fm1(_9_v1, _8_v0, _8_v0, globalState)
	})()), _dafny.IntOfUint32((_47_v30).Cardinality()))).Uint32()).(bool)
	_ = _rhs15
	r1 = _rhs14
	_9_v1 = _rhs15
	var _53_v32 _dafny.Array
	_ = _53_v32
	var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
	_ = _nw8
	_53_v32 = _nw8
	var _54_v33 _dafny.Sequence
	_ = _54_v33
	_54_v33 = _46_v29
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_53_v32), 0))
	_ = _index2
	(_53_v32).ArraySet1((_54_v33), (_index2).Int())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_53_v32), 0))
	_ = _index3
	(_53_v32).ArraySet1(_46_v29, (_index3).Int())
	var _55_v34 _dafny.Array
	_ = _55_v34
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_1
	var _nw9 _dafny.Array
	_ = _nw9
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw9 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) bool = (func(_56_v0 _dafny.Int) func(_dafny.Int) bool {
			return func(_57_i7 _dafny.Int) bool {
				return (_56_v0).Cmp(_56_v0) < 0
			}
		})(_8_v0)
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
	_55_v34 = _nw9
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
	_ = _index4
	(_55_v34).ArraySet1((_dafny.IntOfInt64(483)).Cmp(_8_v0) > 0, (_index4).Int())
	var _58_v35 _dafny.Array
	_ = _58_v35
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
	_ = _nw10
	_58_v35 = _nw10
	var _59_v36 _dafny.Map
	_ = _59_v36
	_59_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v30, _58_v35)
	var _60_v37 D0
	_ = _60_v37
	_60_v37 = Companion_D0_.Create_DC0_((_59_v36).Merge(_59_v36))
	var _pat_let_tv0 = _59_v36
	_ = _pat_let_tv0
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
	_ = _index5
	var _rhs16 bool = _9_v1
	_ = _rhs16
	var _rhs17 _dafny.Int = _8_v0
	_ = _rhs17
	var _rhs18 bool = _9_v1
	_ = _rhs18
	var _rhs19 D0 = func(_pat_let0_0 D0) D0 {
		return func(_61_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 _dafny.Map) D0 {
				return func(_62_dt__update_hcf0_h0 _dafny.Map) D0 {
					return Companion_D0_.Create_DC0_(_62_dt__update_hcf0_h0)
				}(_pat_let1_0)
			}(_pat_let_tv0)
		}(_pat_let0_0)
	}(_60_v37)
	_ = _rhs19
	var _lhs4 _dafny.Array = _55_v34
	_ = _lhs4
	var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
	_ = _lhs5
	_9_v1 = _rhs16
	r1 = _rhs17
	(_lhs4).ArraySet1(_rhs18, (_lhs5).Int())
	_60_v37 = _rhs19
	if (_8_v0).Cmp(_8_v0) < 0 {
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
		_ = _index6
		(_55_v34).ArraySet1(!(_9_v1), (_index6).Int())
		r0 = _8_v0
		(globalState).F7 = _8_v0
		if _9_v1 {
			var _63_v38 _dafny.Map
			_ = _63_v38
			_63_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, !(_9_v1))
			var _64_v39 _dafny.Sequence
			_ = _64_v39
			_64_v39 = _dafny.SeqOf(_63_v38, _63_v38, _63_v38, _63_v38, (_63_v38).Update((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), _9_v1))
			var _65_v40 _dafny.Sequence
			_ = _65_v40
			_65_v40 = _dafny.SeqOf(Companion_Default___.Fm1(true, (_50_v31).Cardinality(), _8_v0, globalState), _8_v0)
			var _66_v41 _dafny.Array
			_ = _66_v41
			var _nwElement0_0 _dafny.Map = _63_v38
			_ = _nwElement0_0
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(21))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_0, 0)
			(_nw11).ArraySet1(_63_v38, 1)
			(_nw11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, (_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)), 2)
			(_nw11).ArraySet1(_63_v38, 3)
			(_nw11).ArraySet1(_63_v38, 4)
			(_nw11).ArraySet1(_63_v38, 5)
			(_nw11).ArraySet1(_63_v38, 6)
			(_nw11).ArraySet1(_63_v38, 7)
			(_nw11).ArraySet1(_63_v38, 8)
			(_nw11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), (_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)), 9)
			(_nw11).ArraySet1(_63_v38, 10)
			(_nw11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, (_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)), 11)
			(_nw11).ArraySet1(_63_v38, 12)
			(_nw11).ArraySet1((_64_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_65_v40).Cardinality()), _dafny.IntOfUint32((_64_v39).Cardinality()))).Uint32()).(_dafny.Map), 13)
			(_nw11).ArraySet1(_63_v38, 14)
			(_nw11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, (_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)), 15)
			(_nw11).ArraySet1(_63_v38, 16)
			(_nw11).ArraySet1(_63_v38, 17)
			(_nw11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), _9_v1), 18)
			(_nw11).ArraySet1(_63_v38, 19)
			(_nw11).ArraySet1(_63_v38, 20)
			_66_v41 = _nw11
			var _67_v42 _dafny.Map
			_ = _67_v42
			_67_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _66_v41)
			var _68_v43 _dafny.Sequence
			_ = _68_v43
			_68_v43 = _dafny.SeqOf(_66_v41, _66_v41, _66_v41)
			var _69_v44 _dafny.Set
			_ = _69_v44
			_69_v44 = _dafny.SetOf(_8_v0, _dafny.IntOfUint32((_65_v40).Cardinality()))
			var _70_v45 bool
			_ = _70_v45
			_70_v45 = _9_v1
			var _71_v46 _dafny.Array
			_ = _71_v46
			var _nwElement0_1 _dafny.Array = _66_v41
			_ = _nwElement0_1
			var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(27))
			_ = _nw12
			(_nw12).ArraySet1(_nwElement0_1, 0)
			(_nw12).ArraySet1(_66_v41, 1)
			(_nw12).ArraySet1(_66_v41, 2)
			(_nw12).ArraySet1(_66_v41, 3)
			(_nw12).ArraySet1(_66_v41, 4)
			(_nw12).ArraySet1(_66_v41, 5)
			(_nw12).ArraySet1(_66_v41, 6)
			(_nw12).ArraySet1(_66_v41, 7)
			(_nw12).ArraySet1(_66_v41, 8)
			(_nw12).ArraySet1(_66_v41, 9)
			(_nw12).ArraySet1(_66_v41, 10)
			(_nw12).ArraySet1(_66_v41, 11)
			(_nw12).ArraySet1(_66_v41, 12)
			(_nw12).ArraySet1(_66_v41, 13)
			(_nw12).ArraySet1(_66_v41, 14)
			(_nw12).ArraySet1((func() _dafny.Array {
				if (_67_v42).Contains((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)) {
					return (_67_v42).Get((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _66_v41
			})(), 15)
			(_nw12).ArraySet1(_66_v41, 16)
			(_nw12).ArraySet1(_66_v41, 17)
			(_nw12).ArraySet1((_68_v43).Select((Companion_Default___.SafeIndex((_69_v44).Cardinality(), _dafny.IntOfUint32((_68_v43).Cardinality()))).Uint32()).(_dafny.Array), 18)
			(_nw12).ArraySet1((func() _dafny.Array {
				if _9_v1 {
					return _66_v41
				}
				return _66_v41
			})(), 19)
			(_nw12).ArraySet1(_66_v41, 20)
			(_nw12).ArraySet1(_66_v41, 21)
			(_nw12).ArraySet1(_66_v41, 22)
			(_nw12).ArraySet1(_66_v41, 23)
			(_nw12).ArraySet1(_66_v41, 24)
			(_nw12).ArraySet1(_66_v41, 25)
			(_nw12).ArraySet1((func() _dafny.Array {
				if _70_v45 {
					return _66_v41
				}
				return _66_v41
			})(), 26)
			_71_v46 = _nw12
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_71_v46), 0))
			_ = _index7
			(_71_v46).ArraySet1(_66_v41, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_71_v46), 0))
			_ = _index8
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_2
			var _nw13 _dafny.Array
			_ = _nw13
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw13 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Map = (func(_72_v38 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_73_i8 _dafny.Int) _dafny.Map {
						return _72_v38
					}
				})(_63_v38)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw13 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw13).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw13).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			(_71_v46).ArraySet1(_nw13, (_index8).Int())
			var _74_v47 *C0
			_ = _74_v47
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_9_v1, _63_v38)
			_74_v47 = _nw14
			_74_v47 = _74_v47
			r1 = _8_v0
			_9_v1 = ((_8_v0).Minus((_63_v38).Cardinality())).Cmp(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_46_v29, (Companion_Default___.SafeIndex(_8_v0, _dafny.IntOfUint32((_46_v29).Cardinality()))).Uint32(), _45_v28)).Cardinality()))).Minus((_dafny.Zero).Minus(_8_v0))) != 0
			(globalState).F8 = _8_v0
		} else {
			_8_v0 = (_dafny.Zero).Minus(_8_v0)
			var _75_v48 _dafny.Map
			_ = _75_v48
			_75_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, _9_v1)
			var _76_v49 *C0
			_ = _76_v49
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(_9_v1, _75_v48)
			_76_v49 = _nw15
			var _77_v50 _dafny.Sequence
			_ = _77_v50
			_77_v50 = _dafny.SeqOf(_58_v35)
			var _78_v51 _dafny.MultiSet
			_ = _78_v51
			_78_v51 = _dafny.MultiSetOf(_58_v35, (_77_v50).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_46_v29).Cardinality()), _dafny.IntOfUint32((_77_v50).Cardinality()))).Uint32()).(_dafny.Array), _58_v35, _58_v35)
			var _79_v52 _dafny.Map
			_ = _79_v52
			_79_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_8_v0).Minus(_8_v0), _76_v49)
			var _80_v53 _dafny.Array
			_ = _80_v53
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(5))
			_ = _nw16
			_80_v53 = _nw16
			var _81_v54 D1
			_ = _81_v54
			_81_v54 = Companion_D1_.Create_DC4_(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("qvldbtq"), _dafny.UnicodeSeqOfUtf8Bytes("aiws")))
			var _82_v55 D1
			_ = _82_v55
			_82_v55 = Companion_D1_.Create_DC5_(_81_v54)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_80_v53), 0))
			_ = _index9
			(_80_v53).ArraySet1(_82_v55, (_index9).Int())
			var _83_v56 _dafny.MultiSet
			_ = _83_v56
			_83_v56 = _dafny.MultiSetOf(_47_v30)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_80_v53), 0))
			_ = _index10
			var _rhs20 *C0 = (func() *C0 {
				if !(_9_v1) {
					return _76_v49
				}
				return _76_v49
			})()
			_ = _rhs20
			var _rhs21 _dafny.MultiSet = _78_v51
			_ = _rhs21
			var _rhs22 _dafny.Int = ((_dafny.Zero).Minus(((_dafny.MultiSetFromSeq(_dafny.SeqOf(_47_v30))).Intersection(_83_v56)).Cardinality())).Plus(_dafny.IntOfInt64(79))
			_ = _rhs22
			var _rhs23 _dafny.Map = (_79_v52).Merge(_79_v52)
			_ = _rhs23
			var _rhs24 D1 = _82_v55
			_ = _rhs24
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			var _lhs7 _dafny.Array = _80_v53
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_80_v53), 0))
			_ = _lhs8
			_76_v49 = _rhs20
			_78_v51 = _rhs21
			_lhs6.F7 = _rhs22
			_79_v52 = _rhs23
			(_lhs7).ArraySet1(_rhs24, (_lhs8).Int())
			_10_v2 = (_10_v2).Update(_9_v1, _45_v28)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
			_ = _index11
			(_55_v34).ArraySet1(_dafny.Companion_Sequence_.Contains(_47_v30, _9_v1), (_index11).Int())
			var _84_v57 _dafny.Map
			_ = _84_v57
			_84_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v28, _76_v49.F19)
			_84_v57 = (_84_v57).Merge(_84_v57)
		}
		r0 = (_dafny.Zero).Minus(_8_v0)
	} else {
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))
		_ = _index12
		(_58_v35).ArraySet1(_8_v0, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))
		_ = _index13
		(_58_v35).ArraySet1(_8_v0, (_index13).Int())
		var _85_v58 _dafny.Set
		_ = _85_v58
		_85_v58 = _dafny.SetOf(_dafny.SeqOf((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)), Companion_Default___.Fm8(_dafny.IntOfInt64(-7), _9_v1, globalState))
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
		_ = _index14
		(_55_v34).ArraySet1((_dafny.SetOf(_47_v30)).IsProperSubsetOf(_85_v58), (_index14).Int())
		_8_v0 = _8_v0
		var _86_v59 _dafny.Set
		_ = _86_v59
		_86_v59 = _dafny.SetOf(_8_v0, (_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int))
		var _87_v60 _dafny.MultiSet
		_ = _87_v60
		_87_v60 = _dafny.MultiSetOf((_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int), (_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int))
		var _88_v61 _dafny.Map
		_ = _88_v61
		_88_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC14_(_87_v60)).Dtor_cf18(), (_53_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_53_v32), 0))).Int()).(_dafny.Sequence))
		var _89_v62 D5
		_ = _89_v62
		_89_v62 = Companion_D5_.Create_DC10_(_86_v59, _88_v61)
		var _source0 D5 = _89_v62
		_ = _source0
		if _source0.Is_DC10() {
			var _90___mcc_h0 _dafny.Set = _source0.Get_().(D5_DC10).Cf9
			_ = _90___mcc_h0
			var _91___mcc_h1 _dafny.Map = _source0.Get_().(D5_DC10).Cf10
			_ = _91___mcc_h1
			var _92_cf10 _dafny.Map = _91___mcc_h1
			_ = _92_cf10
			var _93_cf9 _dafny.Set = _90___mcc_h0
			_ = _93_cf9
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))
			_ = _index15
			(_58_v35).ArraySet1(_8_v0, (_index15).Int())
			var _94_v63 _dafny.Set
			_ = _94_v63
			_94_v63 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(318))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_95_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_96_i9 _dafny.Int) _dafny.CodePoint {
					return _95_v28
				}
			})(_45_v28))))
			_9_v1 = (_94_v63).IsProperSubsetOf((_94_v63).Union(_94_v63))
			var _97_v64 _dafny.Array
			_ = _97_v64
			var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(6))
			_ = _nw17
			_97_v64 = _nw17
			var _98_v65 D1
			_ = _98_v65
			_98_v65 = Companion_D1_.Create_DC3_(_97_v64)
			var _pat_let_tv1 = _97_v64
			_ = _pat_let_tv1
			var _99_v66 _dafny.Array
			_ = _99_v66
			var _nwElement0_2 D1 = _98_v65
			_ = _nwElement0_2
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(28))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_2, 0)
			(_nw18).ArraySet1(_98_v65, 1)
			(_nw18).ArraySet1(_98_v65, 2)
			(_nw18).ArraySet1(_98_v65, 3)
			(_nw18).ArraySet1(_98_v65, 4)
			(_nw18).ArraySet1(_98_v65, 5)
			(_nw18).ArraySet1(_98_v65, 6)
			(_nw18).ArraySet1(Companion_D1_.Create_DC3_(_97_v64), 7)
			(_nw18).ArraySet1(_98_v65, 8)
			(_nw18).ArraySet1(_98_v65, 9)
			(_nw18).ArraySet1(_98_v65, 10)
			(_nw18).ArraySet1((func() D1 {
				if _9_v1 {
					return _98_v65
				}
				return _98_v65
			})(), 11)
			(_nw18).ArraySet1(_98_v65, 12)
			(_nw18).ArraySet1(_98_v65, 13)
			(_nw18).ArraySet1(_98_v65, 14)
			(_nw18).ArraySet1(_98_v65, 15)
			(_nw18).ArraySet1(Companion_D1_.Create_DC3_((Companion_D1_.Create_DC3_(_97_v64)).Dtor_cf2()), 16)
			(_nw18).ArraySet1((func() D1 {
				if _9_v1 {
					return _98_v65
				}
				return _98_v65
			})(), 17)
			(_nw18).ArraySet1(_98_v65, 18)
			(_nw18).ArraySet1(Companion_D1_.Create_DC3_(_97_v64), 19)
			(_nw18).ArraySet1(_98_v65, 20)
			(_nw18).ArraySet1(Companion_D1_.Create_DC3_(_97_v64), 21)
			(_nw18).ArraySet1(_98_v65, 22)
			(_nw18).ArraySet1(func(_pat_let2_0 D1) D1 {
				return func(_100_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let3_0 _dafny.Array) D1 {
						return func(_101_dt__update_hcf2_h0 _dafny.Array) D1 {
							return Companion_D1_.Create_DC3_(_101_dt__update_hcf2_h0)
						}(_pat_let3_0)
					}(_pat_let_tv1)
				}(_pat_let2_0)
			}(_98_v65), 23)
			(_nw18).ArraySet1(_98_v65, 24)
			(_nw18).ArraySet1(_98_v65, 25)
			(_nw18).ArraySet1(_98_v65, 26)
			(_nw18).ArraySet1(_98_v65, 27)
			_99_v66 = _nw18
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_99_v66), 0))
			_ = _index16
			(_99_v66).ArraySet1(_98_v65, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_99_v66), 0))
			_ = _index17
			(_99_v66).ArraySet1(_98_v65, (_index17).Int())
			var _102_v67 D7
			_ = _102_v67
			_102_v67 = Companion_D7_.Create_DC14_(_87_v60)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))
			_ = _index18
			var _rhs25 D7 = (func() D7 {
				if Companion_Default___.Fm0(!(!(_9_v1)), _9_v1, _dafny.UnicodeSeqOfUtf8Bytes("i"), globalState) {
					return _102_v67
				}
				return _102_v67
			})()
			_ = _rhs25
			var _rhs26 _dafny.Int = (_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int)
			_ = _rhs26
			var _lhs9 _dafny.Array = _58_v35
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))
			_ = _lhs10
			_102_v67 = _rhs25
			(_lhs9).ArraySet1(_rhs26, (_lhs10).Int())
		} else if _source0.Is_DC11() {
			var _103___mcc_h2 bool = _source0.Get_().(D5_DC11).Cf11
			_ = _103___mcc_h2
			var _104___mcc_h3 _dafny.Int = _source0.Get_().(D5_DC11).Cf12
			_ = _104___mcc_h3
			var _105___mcc_h4 bool = _source0.Get_().(D5_DC11).Cf13
			_ = _105___mcc_h4
			var _106___mcc_h5 _dafny.Array = _source0.Get_().(D5_DC11).Cf14
			_ = _106___mcc_h5
			var _107___mcc_h6 _dafny.Int = _source0.Get_().(D5_DC11).Cf15
			_ = _107___mcc_h6
			var _108_cf15 _dafny.Int = _107___mcc_h6
			_ = _108_cf15
			var _109_cf14 _dafny.Array = _106___mcc_h5
			_ = _109_cf14
			var _110_cf13 bool = _105___mcc_h4
			_ = _110_cf13
			var _111_cf12 _dafny.Int = _104___mcc_h3
			_ = _111_cf12
			var _112_cf11 bool = _103___mcc_h2
			_ = _112_cf11
			var _113_v68 D5
			_ = _113_v68
			_113_v68 = Companion_D5_.Create_DC11_((_47_v30).Select((Companion_Default___.SafeIndex(_111_cf12, _dafny.IntOfUint32((_47_v30).Cardinality()))).Uint32()).(bool), _8_v0, _9_v1, _109_cf14, (_dafny.Zero).Minus(_111_cf12))
			_113_v68 = _113_v68
			_8_v0 = (func() _dafny.Int {
				if _112_cf11 {
					return _8_v0
				}
				return _8_v0
			})()
			_112_cf11 = _112_cf11
			var _114_v69 *C0
			_ = _114_v69
			var _nw19 *C0 = New_C0_()
			_ = _nw19
			_nw19.Ctor__(((_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int)).Cmp(_8_v0) < 0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), _9_v1))
			_114_v69 = _nw19
			_114_v69 = _114_v69
		} else if _source0.Is_DC9() {
			var _115___mcc_h7 *C0 = _source0.Get_().(D5_DC9).Cf8
			_ = _115___mcc_h7
			var _116_cf8 *C0 = _115___mcc_h7
			_ = _116_cf8
			var _117_v70 _dafny.Int
			_ = _117_v70
			_117_v70 = _dafny.IntOfInt64(96)
			var _118_v71 _dafny.Map
			_ = _118_v71
			_118_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_9_v1), (_117_v70))
			var _119_v72 _dafny.Array
			_ = _119_v72
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(Companion_D5_.Default(), _dafny.IntOfInt64(20))
			_ = _nw20
			_119_v72 = _nw20
			var _120_v73 _dafny.MultiSet
			_ = _120_v73
			_120_v73 = _dafny.MultiSetOf(_119_v72)
			_118_v71 = (_118_v71).Update((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), Companion_Default___.SafeDivisionInt((_120_v73).Cardinality(), _dafny.IntOfUint32(((_53_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_53_v32), 0))).Int()).(_dafny.Sequence)).Cardinality())))
			var _121_v74 D5
			_ = _121_v74
			_121_v74 = Companion_D5_.Create_DC11_(_116_cf8.F19, (_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int), (_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), _55_v34, (_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int))
			var _122_v75 _dafny.MultiSet
			_ = _122_v75
			_122_v75 = _dafny.MultiSetOf((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool))
			var _123_v76 _dafny.Sequence
			_ = _123_v76
			_123_v76 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_116_cf8.F19, _9_v1)), _122_v75, _122_v75)
			var _124_v77 _dafny.Array
			_ = _124_v77
			var _nwElement0_3 _dafny.Int = _8_v0
			_ = _nwElement0_3
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(11))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_3, 0)
			(_nw21).ArraySet1(((_121_v74).Dtor_cf15()).Plus(((_116_cf8).F20()).Cardinality()), 1)
			(_nw21).ArraySet1(_8_v0, 2)
			(_nw21).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ivkit")).Cardinality()), 3)
			(_nw21).ArraySet1((_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int), 4)
			(_nw21).ArraySet1((_dafny.Zero).Minus((_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int)), 5)
			(_nw21).ArraySet1((_dafny.Zero).Minus((_8_v0).Plus((_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int))), 6)
			(_nw21).ArraySet1(_8_v0, 7)
			(_nw21).ArraySet1(_dafny.IntOfUint32((_123_v76).Cardinality()), 8)
			(_nw21).ArraySet1(_8_v0, 9)
			(_nw21).ArraySet1(_8_v0, 10)
			_124_v77 = _nw21
			_124_v77 = _124_v77
			var _125_v78 _dafny.Set
			_ = _125_v78
			_125_v78 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("pbduko"))
			_125_v78 = (_125_v78).Union(_125_v78)
			var _126_v79 _dafny.Array
			_ = _126_v79
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_3
			var _nw22 _dafny.Array
			_ = _nw22
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw22 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_127_v30 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_128_i10 _dafny.Int) _dafny.Sequence {
						return _127_v30
					}
				})(_47_v30)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw22 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw22).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw22).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_126_v79 = _nw22
			var _129_v80 _dafny.Map
			_ = _129_v80
			_129_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v79, (_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int))
			_129_v80 = (_129_v80).Update(_126_v79, (_dafny.MultiSetOf((_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int))).Cardinality())
		} else {
			var _130___mcc_h8 D5 = _source0.Get_().(D5_DC12).Cf16
			_ = _130___mcc_h8
			var _131_cf16 D5 = _130___mcc_h8
			_ = _131_cf16
			var _132_v81 _dafny.Map
			_ = _132_v81
			_132_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_9_v1, !((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)))
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
			_ = _index19
			var _rhs27 bool = (_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)
			_ = _rhs27
			var _rhs28 _dafny.Int = (_132_v81).Cardinality()
			_ = _rhs28
			var _rhs29 bool = true
			_ = _rhs29
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			var _lhs12 _dafny.Array = _55_v34
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))
			_ = _lhs13
			_9_v1 = _rhs27
			_lhs11.F8 = _rhs28
			(_lhs12).ArraySet1(_rhs29, (_lhs13).Int())
			var _133_v82 _dafny.Array
			_ = _133_v82
			var _nw23 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
			_ = _nw23
			_133_v82 = _nw23
			_133_v82 = _133_v82
			(globalState).F8 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-109), (_dafny.Zero).Minus((_58_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_58_v35), 0))).Int()).(_dafny.Int)))
			var _134_v83 *C0
			_ = _134_v83
			var _nw24 *C0 = New_C0_()
			_ = _nw24
			_nw24.Ctor__((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), (_132_v81).Update(false, (_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)))
			_134_v83 = _nw24
		}
		var _135_v84 _dafny.Set
		_ = _135_v84
		_135_v84 = _dafny.SetOf(_9_v1)
		var _136_v85 _dafny.Map
		_ = _136_v85
		_136_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool), _135_v84)
		var _137_v86 _dafny.Sequence
		_ = _137_v86
		_137_v86 = _dafny.SeqOf((func() _dafny.Set {
			if (_136_v85).Contains((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)) {
				return (_136_v85).Get((_55_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_55_v34), 0))).Int()).(bool)).(_dafny.Set)
			}
			return _135_v84
		})(), _135_v84)
		_137_v86 = _137_v86
	}
	r0 = _8_v0
	r1 = Companion_Default___.SafeModuloInt(_8_v0, _dafny.IntOfInt64(554))
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _138_v0 _dafny.CodePoint
	_ = _138_v0
	_138_v0 = _dafny.CodePoint('u')
	var _139_v1 _dafny.Array
	_ = _139_v1
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(13)
	_ = _len0_4
	var _nw25 _dafny.Array
	_ = _nw25
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw25 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Int = func(_140_i0 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeDivisionInt(_140_i0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality()))
		}
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw25 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw25).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw25).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_139_v1 = _nw25
	var _141_v2 _dafny.Map
	_ = _141_v2
	_141_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v0, _139_v1)
	var _142_v3 bool
	_ = _142_v3
	_142_v3 = false
	var _143_v4 _dafny.Set
	_ = _143_v4
	_143_v4 = _dafny.SetOf(_142_v3, _142_v3)
	var _144_v5 _dafny.Int
	_ = _144_v5
	_144_v5 = _dafny.IntOfInt64(172)
	var _145_v6 _dafny.Sequence
	_ = _145_v6
	_145_v6 = _dafny.SeqOf(_142_v3, _142_v3, _142_v3)
	var _146_v7 _dafny.Map
	_ = _146_v7
	_146_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, false)
	var _147_globalState *GlobalState
	_ = _147_globalState
	var _nw26 *GlobalState = New_GlobalState_()
	_ = _nw26
	_nw26.Ctor__(true, _dafny.UnicodeSeqOfUtf8Bytes("pvfmfbioj"), _141_v2, _143_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(492))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}((func(_148_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_149_i1 _dafny.Int) _dafny.CodePoint {
			return _148_v0
		}
	})(_138_v0))), _144_v5), _dafny.IntOfInt64(-984), false, _dafny.IntOfInt64(148), _dafny.IntOfInt64(-376), _dafny.Companion_Sequence_.Concatenate(_145_v6, _145_v6), false, _139_v1, true, _146_v7, _dafny.IntOfInt64(363), _dafny.IntOfInt64(680), true, false, false)
	_147_globalState = _nw26
	var _150_v8 _dafny.MultiSet
	_ = _150_v8
	_150_v8 = _dafny.MultiSetOf(_142_v3)
	_142_v3 = !((_dafny.MultiSetOf(_142_v3, _142_v3, _142_v3)).Intersection(_150_v8)).Contains(_142_v3)
	var _151_v9 _dafny.Sequence
	_ = _151_v9
	_151_v9 = _dafny.UnicodeSeqOfUtf8Bytes("qu")
	var _152_i2 _dafny.Int
	_ = _152_i2
	_152_i2 = _dafny.Zero
	{
		for Companion_Default___.Fm0(_142_v3, _142_v3, _151_v9, _147_globalState) {
			{
				if (_152_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_152_i2 = (_152_i2).Plus(_dafny.One)
				var _153_v10 _dafny.Array
				_ = _153_v10
				var _nwElement0_4 bool = _142_v3
				_ = _nwElement0_4
				var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
				_ = _nw27
				(_nw27).ArraySet1(_nwElement0_4, 0)
				(_nw27).ArraySet1(_142_v3, 1)
				(_nw27).ArraySet1(_142_v3, 2)
				(_nw27).ArraySet1(_142_v3, 3)
				(_nw27).ArraySet1(_142_v3, 4)
				(_nw27).ArraySet1(_142_v3, 5)
				(_nw27).ArraySet1(!(_142_v3), 6)
				(_nw27).ArraySet1(Companion_Default___.Fm0(_142_v3, _142_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(558))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}((func(_154_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_155_i3 _dafny.Int) _dafny.CodePoint {
						return _154_v0
					}
				})(_138_v0))), _147_globalState), 7)
				(_nw27).ArraySet1(_142_v3, 8)
				(_nw27).ArraySet1(false, 9)
				(_nw27).ArraySet1(_142_v3, 10)
				(_nw27).ArraySet1(false, 11)
				(_nw27).ArraySet1(false, 12)
				(_nw27).ArraySet1(_142_v3, 13)
				(_nw27).ArraySet1(_142_v3, 14)
				(_nw27).ArraySet1(_142_v3, 15)
				(_nw27).ArraySet1(_142_v3, 16)
				(_nw27).ArraySet1(_142_v3, 17)
				(_nw27).ArraySet1(_142_v3, 18)
				(_nw27).ArraySet1(_142_v3, 19)
				(_nw27).ArraySet1(_142_v3, 20)
				(_nw27).ArraySet1(_142_v3, 21)
				(_nw27).ArraySet1(_142_v3, 22)
				(_nw27).ArraySet1(_142_v3, 23)
				(_nw27).ArraySet1(_142_v3, 24)
				(_nw27).ArraySet1(_142_v3, 25)
				(_nw27).ArraySet1(_142_v3, 26)
				(_nw27).ArraySet1(Companion_Default___.Fm0(_142_v3, true, _dafny.UnicodeSeqOfUtf8Bytes("w"), _147_globalState), 27)
				_153_v10 = _nw27
				var _156_v11 _dafny.MultiSet
				_ = _156_v11
				_156_v11 = _dafny.MultiSetOf(_153_v10, _153_v10)
				_156_v11 = (((_156_v11).Update(_153_v10, Companion_Default___.Abs(_144_v5))).Difference(_dafny.MultiSetOf(_153_v10))).Difference(_156_v11)
				var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_153_v10), 0))
				_ = _index20
				(_153_v10).ArraySet1(true, (_index20).Int())
				var _157_v12 _dafny.MultiSet
				_ = _157_v12
				_157_v12 = _dafny.MultiSetOf((_151_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(46))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_158_v7 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_159_i4 _dafny.Int) _dafny.Int {
						return (_158_v7).Cardinality()
					}
				})(_146_v7)))).Cardinality()), _dafny.IntOfUint32((_151_v9).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _160_v13 _dafny.Map
				_ = _160_v13
				_160_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_157_v12).Union(_157_v12)).Cardinality(), _142_v3)
				var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_153_v10), 0))
				_ = _index21
				var _rhs30 bool = _142_v3
				_ = _rhs30
				var _rhs31 _dafny.Map = _160_v13
				_ = _rhs31
				var _lhs14 _dafny.Array = _153_v10
				_ = _lhs14
				var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_153_v10), 0))
				_ = _lhs15
				(_lhs14).ArraySet1(_rhs30, (_lhs15).Int())
				_160_v13 = _rhs31
				var _161_v14 _dafny.MultiSet
				_ = _161_v14
				_161_v14 = _dafny.MultiSetOf(_145_v6)
				var _162_v15 _dafny.Array
				_ = _162_v15
				var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw28
				_162_v15 = _nw28
				var _163_v16 _dafny.Map
				_ = _163_v16
				_163_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v14, _162_v15)
				_163_v16 = (_163_v16).Update(_161_v14, _162_v15)
				_142_v3 = ((_153_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_153_v10), 0))).Int()).(bool)) && ((_153_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_153_v10), 0))).Int()).(bool))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _164_v17 _dafny.Int
	_ = _164_v17
	var _165_v18 _dafny.Int
	_ = _165_v18
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	_out0, _out1 = Companion_Default___.M0(_147_globalState)
	_164_v17 = _out0
	_165_v18 = _out1
	var _hi1 _dafny.Int = _164_v17
	_ = _hi1
	for _166_i5 := Companion_Default___.SafeDivisionInt(_144_v5, _164_v17); _166_i5.Cmp(_hi1) < 0; _166_i5 = _166_i5.Plus(_dafny.One) {
		_142_v3 = true
		var _167_v19 _dafny.Map
		_ = _167_v19
		_167_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v1, _dafny.IntOfUint32((_145_v6).Cardinality()))
		var _168_v20 _dafny.Map
		_ = _168_v20
		_168_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v0, _167_v19)
		_168_v20 = (_168_v20).Update(_138_v0, _167_v19)
		_146_v7 = _146_v7
		var _169_v21 _dafny.Int
		_ = _169_v21
		var _170_v22 _dafny.Int
		_ = _170_v22
		var _out2 _dafny.Int
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out2, _out3 = Companion_Default___.M0(_147_globalState)
		_169_v21 = _out2
		_170_v22 = _out3
	}
	(_147_globalState).F8 = _165_v18
	_142_v3 = _142_v3
	var _171_v23 _dafny.Array
	_ = _171_v23
	var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
	_ = _nw29
	_171_v23 = _nw29
	var _172_v24 _dafny.Array
	_ = _172_v24
	var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
	_ = _nw30
	_172_v24 = _nw30
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))
	_ = _index22
	(_171_v23).ArraySet1((func() _dafny.Array {
		if _142_v3 {
			return _172_v24
		}
		return _172_v24
	})(), (_index22).Int())
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))
	_ = _index23
	(_171_v23).ArraySet1(_172_v24, (_index23).Int())
	var _173_v25 _dafny.Array
	_ = _173_v25
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(15)
	_ = _len0_5
	var _nw31 _dafny.Array
	_ = _nw31
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw31 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Sequence = (func(_174_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
			return func(_175_i6 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-562))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_176_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_177_i7 _dafny.Int) _dafny.CodePoint {
						return _176_v0
					}
				})(_174_v0))), _dafny.UnicodeSeqOfUtf8Bytes("uueqdwt"))
			}
		})(_138_v0)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw31 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw31).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw31).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_173_v25 = _nw31
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))
	_ = _index24
	(_173_v25).ArraySet1(_151_v9, (_index24).Int())
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))
	_ = _index25
	(_173_v25).ArraySet1(_151_v9, (_index25).Int())
	(_147_globalState).F9 = _145_v6
	var _178_i8 _dafny.Int
	_ = _178_i8
	_178_i8 = _dafny.Zero
	{
		for !(!(_142_v3) || ((_164_v17).Cmp(_165_v18) == 0)) {
			{
				if (_178_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_178_i8 = (_178_i8).Plus(_dafny.One)
				_138_v0 = _dafny.CodePoint('h')
				var _179_v26 _dafny.MultiSet
				_ = _179_v26
				_179_v26 = _dafny.MultiSetOf(_143_v4)
				var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_139_v1), 0))
				_ = _index26
				(_139_v1).ArraySet1(Companion_Default___.Fm1(_142_v3, (_179_v26).Cardinality(), _165_v18, _147_globalState), (_index26).Int())
				var _180_v27 _dafny.MultiSet
				_ = _180_v27
				_180_v27 = _dafny.MultiSetOf(_138_v0, _138_v0)
				var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_139_v1), 0))
				_ = _index27
				(_139_v1).ArraySet1(Companion_Default___.Fm1((_180_v27).IsSubsetOf(_180_v27), ((_dafny.Zero).Minus(_164_v17)).Times(_165_v18), _164_v17, _147_globalState), (_index27).Int())
				var _181_v28 _dafny.Array
				_ = _181_v28
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_6
				var _nw32 _dafny.Array
				_ = _nw32
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw32 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Int = (func(_182_v18 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_183_i9 _dafny.Int) _dafny.Int {
							return (_183_i9).Times(_182_v18)
						}
					})(_165_v18)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw32 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw32).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw32).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_181_v28 = _nw32
				var _184_v29 _dafny.Map
				_ = _184_v29
				_184_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _181_v28)
				var _185_v30 _dafny.Map
				_ = _185_v30
				_185_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_142_v3), _181_v28)
				var _186_v31 D0
				_ = _186_v31
				_186_v31 = Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v6, _181_v28))
				var _187_v32 _dafny.Array
				_ = _187_v32
				var _nwElement0_5 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v6, (func() _dafny.Array {
					if (_184_v29).Contains(_142_v3) {
						return (_184_v29).Get(_142_v3).(_dafny.Array)
					}
					return _181_v28
				})())).Merge(_185_v30)
				_ = _nwElement0_5
				var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(22))
				_ = _nw33
				(_nw33).ArraySet1(_nwElement0_5, 0)
				(_nw33).ArraySet1((_185_v30).Merge(_185_v30), 1)
				(_nw33).ArraySet1((func() _dafny.Map {
					if _142_v3 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v6, _181_v28)
					}
					return _185_v30
				})(), 2)
				(_nw33).ArraySet1(_185_v30, 3)
				(_nw33).ArraySet1(_185_v30, 4)
				(_nw33).ArraySet1(_185_v30, 5)
				(_nw33).ArraySet1(_185_v30, 6)
				(_nw33).ArraySet1(_185_v30, 7)
				(_nw33).ArraySet1(_185_v30, 8)
				(_nw33).ArraySet1(_185_v30, 9)
				(_nw33).ArraySet1(_185_v30, 10)
				(_nw33).ArraySet1((_186_v31).Dtor_cf0(), 11)
				(_nw33).ArraySet1(_185_v30, 12)
				(_nw33).ArraySet1(_185_v30, 13)
				(_nw33).ArraySet1((_185_v30).Merge(_185_v30), 14)
				(_nw33).ArraySet1(_185_v30, 15)
				(_nw33).ArraySet1(_185_v30, 16)
				(_nw33).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v6, _181_v28), 17)
				(_nw33).ArraySet1(((_186_v31).Dtor_cf0()).Merge(_185_v30), 18)
				(_nw33).ArraySet1((_186_v31).Dtor_cf0(), 19)
				(_nw33).ArraySet1((_185_v30).Update(_145_v6, _181_v28), 20)
				(_nw33).ArraySet1(_185_v30, 21)
				_187_v32 = _nw33
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_187_v32), 0))
				_ = _index28
				(_187_v32).ArraySet1(_185_v30, (_index28).Int())
				var _188_v33 _dafny.Map
				_ = _188_v33
				_188_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _185_v30)
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_187_v32), 0))
				_ = _index29
				(_187_v32).ArraySet1((func() _dafny.Map {
					if _142_v3 {
						return _185_v30
					}
					return (func() _dafny.Map {
						if (_188_v33).Contains(_142_v3) {
							return (_188_v33).Get(_142_v3).(_dafny.Map)
						}
						return _185_v30
					})()
				})(), (_index29).Int())
				(_147_globalState).F8 = _165_v18
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _189_i10 _dafny.Int
	_ = _189_i10
	_189_i10 = _dafny.Zero
	{
		for _dafny.Companion_Sequence_.IsPrefixOf((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), (_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)) {
			{
				if (_189_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_189_i10 = (_189_i10).Plus(_dafny.One)
				var _190_v34 *C0
				_ = _190_v34
				var _nw34 *C0 = New_C0_()
				_ = _nw34
				_nw34.Ctor__(false, _146_v7)
				_190_v34 = _nw34
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_139_v1), 0))
				_ = _index30
				(_139_v1).ArraySet1(_144_v5, (_index30).Int())
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_139_v1), 0))
				_ = _index31
				(_139_v1).ArraySet1(_dafny.IntOfInt64(-698), (_index31).Int())
				var _191_v35 _dafny.Array
				_ = _191_v35
				var _nwElement0_6 _dafny.MultiSet = _150_v8
				_ = _nwElement0_6
				var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(29))
				_ = _nw35
				(_nw35).ArraySet1(_nwElement0_6, 0)
				(_nw35).ArraySet1((_dafny.MultiSetOf(_190_v34.F19, _142_v3, _190_v34.F19, (func() bool {
					if (_146_v7).Contains(_142_v3) {
						return (_146_v7).Get(_142_v3).(bool)
					}
					return _142_v3
				})())).Difference(_dafny.MultiSetOf(_190_v34.F19, _142_v3)), 1)
				(_nw35).ArraySet1(_dafny.MultiSetOf(Companion_Default___.Fm0(_142_v3, _142_v3, Companion_Default___.Fm4(_190_v34.F19, _142_v3, !(_142_v3), (_143_v4).Cardinality(), _147_globalState), _147_globalState), _190_v34.F19), 2)
				(_nw35).ArraySet1(_dafny.MultiSetFromSeq(_145_v6), 3)
				(_nw35).ArraySet1(_150_v8, 4)
				(_nw35).ArraySet1((_150_v8).Update(_142_v3, Companion_Default___.Abs((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))), 5)
				(_nw35).ArraySet1((_150_v8).Difference(_dafny.MultiSetOf(_142_v3)), 6)
				(_nw35).ArraySet1(_150_v8, 7)
				(_nw35).ArraySet1((_150_v8).Update(_190_v34.F19, Companion_Default___.Abs(_165_v18)), 8)
				(_nw35).ArraySet1(_150_v8, 9)
				(_nw35).ArraySet1(_150_v8, 10)
				(_nw35).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_190_v34.F19)), 11)
				(_nw35).ArraySet1(_150_v8, 12)
				(_nw35).ArraySet1((_150_v8).Update(false, Companion_Default___.Abs((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))), 13)
				(_nw35).ArraySet1(_150_v8, 14)
				(_nw35).ArraySet1((_150_v8).Difference(_dafny.MultiSetOf(_190_v34.F19, _190_v34.F19)), 15)
				(_nw35).ArraySet1(_150_v8, 16)
				(_nw35).ArraySet1(_dafny.MultiSetOf(_190_v34.F19, _190_v34.F19, _190_v34.F19), 17)
				(_nw35).ArraySet1(_150_v8, 18)
				(_nw35).ArraySet1((_150_v8).Union(_dafny.MultiSetOf(_142_v3)), 19)
				(_nw35).ArraySet1(_dafny.MultiSetOf(_142_v3), 20)
				(_nw35).ArraySet1(_150_v8, 21)
				(_nw35).ArraySet1((_150_v8).Update(_142_v3, Companion_Default___.Abs(_164_v17)), 22)
				(_nw35).ArraySet1(_150_v8, 23)
				(_nw35).ArraySet1(_150_v8, 24)
				(_nw35).ArraySet1(_dafny.MultiSetOf(_190_v34.F19, _142_v3, _142_v3), 25)
				(_nw35).ArraySet1(_150_v8, 26)
				(_nw35).ArraySet1((_150_v8).Update(Companion_Default___.Fm0(_142_v3, _142_v3, (_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), _147_globalState), Companion_Default___.Abs((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))), 27)
				(_nw35).ArraySet1(_150_v8, 28)
				_191_v35 = _nw35
				_191_v35 = (Companion_D1_.Create_DC3_(_191_v35)).Dtor_cf2()
				(_190_v34).F19 = _190_v34.F19
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _source1 D0 = Companion_D0_.Create_DC1_()
	_ = _source1
	if _source1.Is_DC1() {
		var _192_v36 _dafny.Map
		_ = _192_v36
		_192_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(!(_142_v3), _142_v3, (_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), _147_globalState), Companion_Default___.Fm5(_150_v8, _147_globalState))
		var _193_v37 _dafny.Set
		_ = _193_v37
		_193_v37 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("odbnb"))
		var _194_v38 D1
		_ = _194_v38
		_194_v38 = Companion_D1_.Create_DC4_((func() _dafny.Set {
			if (_192_v36).Contains(!(_142_v3)) {
				return (_192_v36).Get(!(_142_v3)).(_dafny.Set)
			}
			return _193_v37
		})())
		var _195_v39 _dafny.Sequence
		_ = _195_v39
		_195_v39 = _dafny.SeqOf(_dafny.IntOfInt64(40))
		var _196_v40 _dafny.Sequence
		_ = _196_v40
		_196_v40 = _dafny.SeqOf(_195_v39, Companion_Default___.Fm7(_147_globalState))
		var _pat_let_tv2 = _193_v37
		_ = _pat_let_tv2
		_194_v38 = func(_pat_let4_0 D1) D1 {
			return func(_197_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let5_0 _dafny.Set) D1 {
					return func(_198_dt__update_hcf3_h0 _dafny.Set) D1 {
						return Companion_D1_.Create_DC4_(_198_dt__update_hcf3_h0)
					}(_pat_let5_0)
				}(_pat_let_tv2)
			}(_pat_let4_0)
		}(Companion_Default___.Fm6(_196_v40, _142_v3, _142_v3, _147_globalState))
		_142_v3 = (_142_v3) && (_142_v3)
		var _199_v41 _dafny.Map
		_ = _199_v41
		_199_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(_147_globalState), _142_v3)
		_199_v41 = (_199_v41).Update(_195_v39, (_142_v3) == (_142_v3))
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_172_v24), 0))
		_ = _index32
		(_172_v24).ArraySet1(_142_v3, (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_172_v24), 0))
		_ = _index33
		(_172_v24).ArraySet1((_dafny.IntOfUint32(((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)).Cardinality())).Cmp((_143_v4).Cardinality()) >= 0, (_index33).Int())
	} else if _source1.Is_DC0() {
		var _200___mcc_h0 _dafny.Map = _source1.Get_().(D0_DC0).Cf0
		_ = _200___mcc_h0
		var _201_cf0 _dafny.Map = _200___mcc_h0
		_ = _201_cf0
		(_147_globalState).F8 = _dafny.IntOfInt64(551)
		var _202_v42 *C0
		_ = _202_v42
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__(!(_142_v3) || (_142_v3), (_146_v7).Merge(_146_v7))
		_202_v42 = _nw36
		_202_v42 = _202_v42
		(_147_globalState).F8 = _144_v5
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))
		_ = _arr0
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))), 0))
		_ = _index34
		(_arr0).ArraySet1(_142_v3, (_index34).Int())
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))
		_ = _arr1
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))), 0))
		_ = _index35
		(_arr1).ArraySet1((_165_v18).Cmp(_dafny.IntOfInt64(-638)) < 0, (_index35).Int())
	} else {
		var _203___mcc_h1 D0 = _source1.Get_().(D0_DC2).Cf1
		_ = _203___mcc_h1
		var _204_cf1 D0 = _203___mcc_h1
		_ = _204_cf1
		var _205_v43 _dafny.Array
		_ = _205_v43
		var _nw37 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
		_ = _nw37
		_205_v43 = _nw37
		_205_v43 = _205_v43
		if _142_v3 {
			(_147_globalState).F7 = _144_v5
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_172_v24), 0))
			_ = _index36
			(_172_v24).ArraySet1(_142_v3, (_index36).Int())
			var _206_v44 _dafny.Map
			_ = _206_v44
			_206_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _dafny.CodePoint('g'))
			var _207_v45 _dafny.Set
			_ = _207_v45
			_207_v45 = _dafny.SetOf(_144_v5, _144_v5, Companion_Default___.Fm1(_142_v3, _144_v5, (_206_v44).Cardinality(), _147_globalState), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-178), _164_v17), _165_v18)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_172_v24), 0))
			_ = _index37
			var _rhs32 _dafny.Int = (func() _dafny.Int {
				if Companion_Default___.Fm0(_142_v3, _142_v3, _151_v9, _147_globalState) {
					return _144_v5
				}
				return _165_v18
			})()
			_ = _rhs32
			var _rhs33 _dafny.Sequence = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_164_v17).Cmp(_165_v18) <= 0 {
					return _145_v6
				}
				return Companion_Default___.Fm8(_165_v18, _142_v3, _147_globalState)
			})(), (Companion_Default___.SafeIndex(_144_v5, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_164_v17).Cmp(_165_v18) <= 0 {
					return _145_v6
				}
				return Companion_Default___.Fm8(_165_v18, _142_v3, _147_globalState)
			})()).Cardinality()))).Uint32(), !(_142_v3) || (!(false)))
			_ = _rhs33
			var _rhs34 bool = (func() bool {
				if _142_v3 {
					return !(_dafny.Companion_Sequence_.IsPrefixOf((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), (_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)))
				}
				return (_dafny.IntOfInt64(-337)).Cmp((_dafny.MultiSetOf(_dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int())), _172_v24, _172_v24)).Cardinality()) > 0
			})()
			_ = _rhs34
			var _rhs35 _dafny.Set = _207_v45
			_ = _rhs35
			var _lhs16 *GlobalState = _147_globalState
			_ = _lhs16
			var _lhs17 *GlobalState = _147_globalState
			_ = _lhs17
			var _lhs18 _dafny.Array = _172_v24
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_172_v24), 0))
			_ = _lhs19
			_lhs16.F8 = _rhs32
			_lhs17.F9 = _rhs33
			(_lhs18).ArraySet1(_rhs34, (_lhs19).Int())
			_207_v45 = _rhs35
			var _208_v46 _dafny.Int
			_ = _208_v46
			var _209_v47 _dafny.Int
			_ = _209_v47
			var _out4 _dafny.Int
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out4, _out5 = Companion_Default___.M0(_147_globalState)
			_208_v46 = _out4
			_209_v47 = _out5
			var _210_v48 *C0
			_ = _210_v48
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__(!(true), _146_v7)
			_210_v48 = _nw38
			_210_v48 = _210_v48
			_142_v3 = ((_210_v48).Fm2(_138_v0, Companion_Default___.Fm0(_210_v48.F19, (_172_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_172_v24), 0))).Int()).(bool), _151_v9, _147_globalState), _147_globalState)).Cmp(_144_v5) >= 0
		} else {
			var _211_v49 _dafny.Int
			_ = _211_v49
			var _212_v50 _dafny.Int
			_ = _212_v50
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out6, _out7 = Companion_Default___.M0(_147_globalState)
			_211_v49 = _out6
			_212_v50 = _out7
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_172_v24), 0))
			_ = _index38
			(_172_v24).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm8(_164_v17, _142_v3, _147_globalState), Companion_Default___.Fm8(_211_v49, _142_v3, _147_globalState)), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_172_v24), 0))
			_ = _index39
			(_172_v24).ArraySet1((!(_142_v3)) && (_142_v3), (_index39).Int())
			var _213_v51 _dafny.Map
			_ = _213_v51
			_213_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mnvbm"), _212_v50)
			(_147_globalState).F7 = (func() _dafny.Int {
				if (_213_v51).Contains(_151_v9) {
					return (_213_v51).Get(_151_v9).(_dafny.Int)
				}
				return _212_v50
			})()
			var _214_v52 _dafny.Map
			_ = _214_v52
			_214_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v49, _211_v49)
			var _215_v53 *C0
			_ = _215_v53
			var _nw39 *C0 = New_C0_()
			_ = _nw39
			_nw39.Ctor__(_142_v3, Companion_Default___.Fm9((func() _dafny.Int {
				if (_214_v52).Contains((_dafny.IntOfUint32(((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)).Cardinality()))) {
					return (_214_v52).Get((_dafny.IntOfUint32(((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)).Cardinality()))).(_dafny.Int)
				}
				return _211_v49
			})(), _144_v5, _dafny.IntOfUint32((_151_v9).Cardinality()), _147_globalState))
			_215_v53 = _nw39
			var _216_v54 _dafny.Map
			_ = _216_v54
			_216_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v1, (_212_v50).Minus(_144_v5))
			_216_v54 = (_216_v54).Update(_139_v1, (_dafny.Zero).Minus(_211_v49))
		}
		_138_v0 = _138_v0
		_142_v3 = !(true)
	}
	var _217_i11 _dafny.Int
	_ = _217_i11
	_217_i11 = _dafny.Zero
	{
		for _142_v3 {
			{
				if (_217_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_217_i11 = (_217_i11).Plus(_dafny.One)
				var _rhs36 bool = !(_142_v3) || (_142_v3)
				_ = _rhs36
				var _rhs37 _dafny.Int = _164_v17
				_ = _rhs37
				var _rhs38 _dafny.Int = _dafny.IntOfInt64(744)
				_ = _rhs38
				var _lhs20 *GlobalState = _147_globalState
				_ = _lhs20
				var _lhs21 *GlobalState = _147_globalState
				_ = _lhs21
				_142_v3 = _rhs36
				_lhs20.F8 = _rhs37
				_lhs21.F8 = _rhs38
				if Companion_Default___.Fm0(_142_v3, _142_v3, _dafny.UnicodeSeqOfUtf8Bytes("xahmwdx"), _147_globalState) {
					var _218_v55 _dafny.Map
					_ = _218_v55
					_218_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_151_v9).Cardinality())).Times(_165_v18), _dafny.CodePoint('p'))
					_218_v55 = (_218_v55).Update(_144_v5, _138_v0)
					var _219_v56 *C0
					_ = _219_v56
					var _nw40 *C0 = New_C0_()
					_ = _nw40
					_nw40.Ctor__(_142_v3, _146_v7)
					_219_v56 = _nw40
					var _220_v57 _dafny.Int
					_ = _220_v57
					var _221_v58 _dafny.Int
					_ = _221_v58
					var _out8 _dafny.Int
					_ = _out8
					var _out9 _dafny.Int
					_ = _out9
					_out8, _out9 = Companion_Default___.M0(_147_globalState)
					_220_v57 = _out8
					_221_v58 = _out9
					var _222_v59 bool
					_ = _222_v59
					_222_v59 = !(_142_v3)
					(_219_v56).F19 = (_222_v59)
					var _223_v60 _dafny.Sequence
					_ = _223_v60
					_223_v60 = _dafny.SeqOf((func() _dafny.Int {
						if (_150_v8).Contains(_219_v56.F19) {
							return (_150_v8).Multiplicity(_219_v56.F19)
						}
						return _165_v18
					})(), _144_v5)
					var _224_v61 _dafny.Map
					_ = _224_v61
					_224_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v56, _223_v60)
					_224_v61 = (_224_v61).Update(_219_v56, _223_v60)
				} else {
					var _225_v62 *C0
					_ = _225_v62
					var _nw41 *C0 = New_C0_()
					_ = _nw41
					_nw41.Ctor__(_142_v3, (_146_v7).Merge(_146_v7))
					_225_v62 = _nw41
					_146_v7 = (_146_v7).Update(Companion_Default___.Fm0(_142_v3, _142_v3, (_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), _147_globalState), !(_142_v3))
					var _226_v63 _dafny.Map
					_ = _226_v63
					_226_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm1(_142_v3, _165_v18, (_dafny.Zero).Minus(_144_v5), _147_globalState), _164_v17), Companion_Default___.SafeDivisionInt(_164_v17, _165_v18))
					_226_v63 = (_226_v63).Update((_165_v18).Times(_165_v18), (_dafny.MultiSetOf((_dafny.Zero).Minus(_164_v17), (_225_v62).Fm2(_138_v0, _225_v62.F19, _147_globalState), _144_v5)).Cardinality())
					_146_v7 = (_146_v7).Update(false, _142_v3)
					var _227_v64 *C0
					_ = _227_v64
					var _nw42 *C0 = New_C0_()
					_ = _nw42
					_nw42.Ctor__((_225_v62.F19) == ((_145_v6).Select((Companion_Default___.SafeIndex(_165_v18, _dafny.IntOfUint32((_145_v6).Cardinality()))).Uint32()).(bool)), _146_v7)
					_227_v64 = _nw42
				}
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_139_v1), 0))
				_ = _index40
				(_139_v1).ArraySet1(_165_v18, (_index40).Int())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_139_v1), 0))
				_ = _index41
				var _rhs39 _dafny.Int = _dafny.IntOfUint32(((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)).Cardinality())
				_ = _rhs39
				var _rhs40 _dafny.Int = Companion_Default___.SafeDivisionInt(_164_v17, _164_v17)
				_ = _rhs40
				var _rhs41 _dafny.Int = _144_v5
				_ = _rhs41
				var _lhs22 *GlobalState = _147_globalState
				_ = _lhs22
				var _lhs23 *GlobalState = _147_globalState
				_ = _lhs23
				var _lhs24 _dafny.Array = _139_v1
				_ = _lhs24
				var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_139_v1), 0))
				_ = _lhs25
				_lhs22.F8 = _rhs39
				_lhs23.F7 = _rhs40
				(_lhs24).ArraySet1(_rhs41, (_lhs25).Int())
				var _228_v65 _dafny.MultiSet
				_ = _228_v65
				_228_v65 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_164_v17, _dafny.IntOfUint32(((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _138_v0), _151_v9)
				var _229_v66 _dafny.Sequence
				_ = _229_v66
				_229_v66 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_144_v5, _dafny.IntOfUint32(((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _138_v0))
				_142_v3 = (_dafny.MultiSetFromSeq(_229_v66)).IsProperSubsetOf((_228_v65).Union(_228_v65))
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	_142_v3 = _142_v3
	if _142_v3 {
		if (false) && (_142_v3) {
			var _230_v67 *C0
			_ = _230_v67
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__(_142_v3, _146_v7)
			_230_v67 = _nw43
			_230_v67 = _230_v67
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_7
			var _nw44 _dafny.Array
			_ = _nw44
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw44 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = (func(_231_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_232_i12 _dafny.Int) _dafny.Int {
						return (_232_i12).Times(_231_v5)
					}
				})(_144_v5)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw44 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw44).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw44).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_139_v1 = _nw44
			(_147_globalState).F8 = _165_v18
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))
			_ = _arr2
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))), 0))
			_ = _index42
			(_arr2).ArraySet1((_144_v5).Cmp(_164_v17) != 0, (_index42).Int())
			var _233_v68 _dafny.Set
			_ = _233_v68
			_233_v68 = _dafny.SetOf(_164_v17, _dafny.IntOfInt64(151))
			var _234_v70 _dafny.Map
			_ = _234_v70
			_234_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll3 = _dafny.NewBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(168), _dafny.IntOfInt64(106))); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _235_v69 _dafny.Int
					_235_v69 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(168)).Cmp(_235_v69) <= 0) && ((_235_v69).Cmp(_dafny.IntOfInt64(106)) < 0) {
						_coll3.Add((_235_v69).Times(_144_v5))
					}
				}
				return _coll3.ToSet()
			}(), _145_v6)
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))
			_ = _arr3
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int()))), 0))
			_ = _index43
			(_arr3).ArraySet1(!(_234_v70).Contains(_233_v68), (_index43).Int())
			_230_v67 = _230_v67
		} else {
			var _236_v71 *C0
			_ = _236_v71
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__((func() bool {
				if (_146_v7).Contains(_142_v3) {
					return (_146_v7).Get(_142_v3).(bool)
				}
				return _142_v3
			})(), _146_v7)
			_236_v71 = _nw45
			var _237_v72 D5
			_ = _237_v72
			_237_v72 = Companion_D5_.Create_DC9_(_236_v71)
			var _pat_let_tv3 = _236_v71
			_ = _pat_let_tv3
			_236_v71 = (func(_pat_let6_0 D5) D5 {
				return func(_238_dt__update__tmp_h1 D5) D5 {
					return func(_pat_let7_0 *C0) D5 {
						return func(_239_dt__update_hcf8_h0 *C0) D5 {
							return Companion_D5_.Create_DC9_(_239_dt__update_hcf8_h0)
						}(_pat_let7_0)
					}(_pat_let_tv3)
				}(_pat_let6_0)
			}(_237_v72)).Dtor_cf8()
			var _240_v73 _dafny.Array
			_ = _240_v73
			var _nwElement0_7 _dafny.MultiSet = _150_v8
			_ = _nwElement0_7
			var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(8))
			_ = _nw46
			(_nw46).ArraySet1(_nwElement0_7, 0)
			(_nw46).ArraySet1(_150_v8, 1)
			(_nw46).ArraySet1(_150_v8, 2)
			(_nw46).ArraySet1(_150_v8, 3)
			(_nw46).ArraySet1(_150_v8, 4)
			(_nw46).ArraySet1(_150_v8, 5)
			(_nw46).ArraySet1(_150_v8, 6)
			(_nw46).ArraySet1(_dafny.MultiSetOf(false), 7)
			_240_v73 = _nw46
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(19))
			_ = _nw47
			_240_v73 = _nw47
			var _241_v74 _dafny.Sequence
			_ = _241_v74
			_241_v74 = _dafny.SeqOf((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), _151_v9)
			_151_v9 = _dafny.Companion_Sequence_.Concatenate((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate((_241_v74).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.IntOfUint32((_241_v74).Cardinality()))).Uint32()).(_dafny.Sequence), _151_v9))
			_236_v71 = _236_v71
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_139_v1), 0))
			_ = _index44
			(_139_v1).ArraySet1(_dafny.IntOfInt64(-834), (_index44).Int())
			var _242_v75 _dafny.Set
			_ = _242_v75
			_242_v75 = _dafny.SetOf((_236_v71).Fm2(_138_v0, !(!(_236_v71.F19)), _147_globalState), _165_v18)
			var _243_v76 _dafny.MultiSet
			_ = _243_v76
			_243_v76 = _dafny.MultiSetOf((_150_v8).Cardinality())
			var _244_v77 _dafny.Map
			_ = _244_v77
			_244_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_243_v76).Update(_dafny.IntOfUint32((_151_v9).Cardinality()), Companion_Default___.Abs(_144_v5)), _165_v18)
			var _245_v78 _dafny.Sequence
			_ = _245_v78
			_245_v78 = _dafny.SeqOf(_164_v17)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_139_v1), 0))
			_ = _index45
			var _rhs42 _dafny.Int = (_245_v78).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_165_v18), _dafny.IntOfUint32((_245_v78).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs42
			var _rhs43 _dafny.Set = (func() _dafny.Set {
				if _142_v3 {
					return _242_v75
				}
				return _242_v75
			})()
			_ = _rhs43
			var _rhs44 _dafny.CodePoint = Companion_Default___.Fm10(_245_v78, Companion_Default___.Fm11(_138_v0, _236_v71.F19, _164_v17, _165_v18, _147_globalState), _147_globalState)
			_ = _rhs44
			var _rhs45 _dafny.Map = (_244_v77).Merge(_244_v77)
			_ = _rhs45
			var _rhs46 bool = _142_v3
			_ = _rhs46
			var _lhs26 _dafny.Array = _139_v1
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_139_v1), 0))
			_ = _lhs27
			(_lhs26).ArraySet1(_rhs42, (_lhs27).Int())
			_242_v75 = _rhs43
			_138_v0 = _rhs44
			_244_v77 = _rhs45
			_142_v3 = _rhs46
		}
		var _246_v79 _dafny.Sequence
		_ = _246_v79
		_246_v79 = _dafny.SeqOf(_164_v17)
		var _247_v80 _dafny.Sequence
		_ = _247_v80
		_247_v80 = _dafny.SeqOf(_246_v79, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-204))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_248_v17 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_249_i13 _dafny.Int) _dafny.Int {
				return (_dafny.MultiSetOf(_248_v17, _dafny.IntOfInt64(362))).Cardinality()
			}
		})(_164_v17))))
		_143_v4 = Companion_Default___.Fm12(_247_v80, (_144_v5).Cmp(_144_v5) >= 0, (func() bool {
			if (_146_v7).Contains(!(false)) {
				return (_146_v7).Get(!(false)).(bool)
			}
			return _142_v3
		})(), _147_globalState)
		if !(_142_v3) || ((func() bool {
			if (_145_v6).Select((Companion_Default___.SafeIndex(_165_v18, _dafny.IntOfUint32((_145_v6).Cardinality()))).Uint32()).(bool) {
				return _142_v3
			}
			return _142_v3
		})()) {
			var _250_v81 _dafny.Array
			_ = _250_v81
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(24))
			_ = _nw48
			_250_v81 = _nw48
			_250_v81 = _250_v81
			_171_v23 = _171_v23
			var _251_v82 _dafny.Sequence
			_ = _251_v82
			_251_v82 = _dafny.SeqOf(_172_v24, _172_v24, _dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int())))
			_251_v82 = _dafny.Companion_Sequence_.Concatenate(_251_v82, _251_v82)
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_172_v24), 0))
			_ = _index46
			(_172_v24).ArraySet1(!(true), (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_172_v24), 0))
			_ = _index47
			(_172_v24).ArraySet1(((func() bool {
				if (_146_v7).Contains(_142_v3) {
					return (_146_v7).Get(_142_v3).(bool)
				}
				return _142_v3
			})()) && (_142_v3), (_index47).Int())
			(_147_globalState).F8 = _164_v17
		} else {
			var _252_v83 *C0
			_ = _252_v83
			var _nw49 *C0 = New_C0_()
			_ = _nw49
			_nw49.Ctor__(_142_v3, _146_v7)
			_252_v83 = _nw49
			_252_v83 = _252_v83
			(_147_globalState).F7 = Companion_Default___.SafeModuloInt(_144_v5, _165_v18)
			var _253_v84 _dafny.Sequence
			_ = _253_v84
			_253_v84 = _dafny.SeqOf(_252_v83)
			_252_v83 = (_253_v84).Select((Companion_Default___.SafeIndex((_144_v5).Minus(_144_v5), _dafny.IntOfUint32((_253_v84).Cardinality()))).Uint32()).(*C0)
			(_147_globalState).F8 = _dafny.IntOfUint32((_dafny.SeqOf(_252_v83.F19, !(_142_v3), _142_v3, (func() bool {
				if _142_v3 {
					return _252_v83.F19
				}
				return _142_v3
			})())).Cardinality())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_172_v24), 0))
			_ = _index48
			(_172_v24).ArraySet1(_252_v83.F19, (_index48).Int())
			var _254_v85 _dafny.Map
			_ = _254_v85
			_254_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(442), _142_v3)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_172_v24), 0))
			_ = _index49
			var _rhs47 bool = _252_v83.F19
			_ = _rhs47
			var _rhs48 _dafny.Map = _254_v85
			_ = _rhs48
			var _lhs28 _dafny.Array = _172_v24
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_172_v24), 0))
			_ = _lhs29
			(_lhs28).ArraySet1(_rhs47, (_lhs29).Int())
			_254_v85 = _rhs48
		}
		var _255_v87 _dafny.MultiSet
		_ = _255_v87
		_255_v87 = _dafny.MultiSetOf((_dafny.Zero).Minus((_146_v7).Cardinality()), _dafny.IntOfInt64(693), _164_v17)
		var _256_v88 _dafny.Set
		_ = _256_v88
		_256_v88 = _dafny.SetOf(_dafny.MultiSetOf(_164_v17), _255_v87, _255_v87, _255_v87, _255_v87)
		var _257_v89 D5
		_ = _257_v89
		_257_v89 = Companion_D5_.Create_DC10_(_dafny.SetOf(_144_v5, _144_v5, _165_v18), func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_256_v88).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _258_v86 _dafny.MultiSet
				_258_v86 = interface{}(_compr_4).(_dafny.MultiSet)
				if (_256_v88).Contains(_258_v86) {
					_coll4.Add(_258_v86, _151_v9)
				}
			}
			return _coll4.ToMap()
		}())
		var _259_v90 _dafny.Sequence
		_ = _259_v90
		_259_v90 = _dafny.SeqOf((func() D5 {
			if _142_v3 {
				return _257_v89
			}
			return _257_v89
		})(), _257_v89, _257_v89)
		var _260_v91 _dafny.Set
		_ = _260_v91
		_260_v91 = _dafny.SetOf(_144_v5)
		var _261_v92 _dafny.Map
		_ = _261_v92
		_261_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v87, (_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence))
		var _rhs49 _dafny.Array = _139_v1
		_ = _rhs49
		var _rhs50 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_257_v89, Companion_D5_.Create_DC10_(_260_v91, _261_v92)), (Companion_Default___.SafeIndex(_164_v17, _dafny.IntOfUint32((_dafny.SeqOf(_257_v89, Companion_D5_.Create_DC10_(_260_v91, _261_v92))).Cardinality()))).Uint32(), _257_v89)
		_ = _rhs50
		var _rhs51 _dafny.CodePoint = _138_v0
		_ = _rhs51
		_139_v1 = _rhs49
		_259_v90 = _rhs50
		_138_v0 = _rhs51
		var _262_v93 _dafny.Sequence
		_ = _262_v93
		_262_v93 = _151_v9
		var _263_v94 *C0
		_ = _263_v94
		var _nw50 *C0 = New_C0_()
		_ = _nw50
		_nw50.Ctor__(Companion_Default___.Fm0(_142_v3, _142_v3, (_262_v93), _147_globalState), _146_v7)
		_263_v94 = _nw50
	} else {
		if ((_164_v17).Times(_dafny.IntOfInt64(821))).Cmp(_144_v5) > 0 {
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))
			_ = _index50
			(_139_v1).ArraySet1((_165_v18).Minus(Companion_Default___.Fm1(_142_v3, _165_v18, (_dafny.Zero).Minus(_164_v17), _147_globalState)), (_index50).Int())
			var _264_v95 _dafny.Sequence
			_ = _264_v95
			_264_v95 = _dafny.SeqOf(_139_v1)
			var _265_v96 _dafny.Map
			_ = _265_v96
			_265_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v18, _165_v18)
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))
			_ = _index51
			var _rhs52 _dafny.Int = (func() _dafny.Int {
				if _142_v3 {
					return _dafny.IntOfUint32((_151_v9).Cardinality())
				}
				return ((_dafny.Zero).Minus(_164_v17)).Minus(_164_v17)
			})()
			_ = _rhs52
			var _rhs53 _dafny.Array = (_264_v95).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_142_v3, _142_v3)).Cardinality(), _dafny.IntOfUint32((_264_v95).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs53
			var _rhs54 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_165_v18), Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_265_v96).Contains(_144_v5) {
					return (_265_v96).Get(_144_v5).(_dafny.Int)
				}
				return _164_v17
			})(), _144_v5))
			_ = _rhs54
			var _lhs30 *GlobalState = _147_globalState
			_ = _lhs30
			var _lhs31 _dafny.Array = _139_v1
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))
			_ = _lhs32
			_lhs30.F7 = _rhs52
			_139_v1 = _rhs53
			(_lhs31).ArraySet1(_rhs54, (_lhs32).Int())
			var _266_v97 _dafny.MultiSet
			_ = _266_v97
			_266_v97 = _dafny.MultiSetOf(_165_v18, (_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))
			var _267_v98 _dafny.Map
			_ = _267_v98
			_267_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v97, _dafny.IntOfUint32(((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence)).Cardinality()))
			var _268_v99 _dafny.Map
			_ = _268_v99
			_268_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int), _267_v98)
			var _269_v100 _dafny.Sequence
			_ = _269_v100
			_269_v100 = _dafny.SeqOf(_146_v7, _146_v7, _146_v7)
			var _270_v102 *C0
			_ = _270_v102
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__(_142_v3, _146_v7)
			_270_v102 = _nw51
			var _271_v103 _dafny.Map
			_ = _271_v103
			_271_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_v102, _142_v3)
			var _272_v104 _dafny.Array
			_ = _272_v104
			var _nwElement0_8 _dafny.Map = (func() _dafny.Map {
				if _142_v3 {
					return (_267_v98).Update((_266_v97).Update((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int), Companion_Default___.Abs(_165_v18)), (_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))
				}
				return (func() _dafny.Map {
					if (_268_v99).Contains(((_269_v100).Select((Companion_Default___.SafeIndex(_165_v18, _dafny.IntOfUint32((_269_v100).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()) {
						return (_268_v99).Get(((_269_v100).Select((Companion_Default___.SafeIndex(_165_v18, _dafny.IntOfUint32((_269_v100).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).(_dafny.Map)
					}
					return _267_v98
				})()
			})()
			_ = _nwElement0_8
			var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(8))
			_ = _nw52
			(_nw52).ArraySet1(_nwElement0_8, 0)
			(_nw52).ArraySet1((func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_dafny.SeqOf(_266_v97)).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _273_v101 _dafny.MultiSet
					_273_v101 = interface{}(_compr_5).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_266_v97), _273_v101) {
						_coll5.Add(_273_v101, _144_v5)
					}
				}
				return _coll5.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v97, _144_v5)), 1)
			(_nw52).ArraySet1(_267_v98, 2)
			(_nw52).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v97, _144_v5)).Update(_266_v97, (_271_v103).Cardinality()), 3)
			(_nw52).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_164_v17, _dafny.IntOfInt64(274), (_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int), _164_v17), _164_v17), 4)
			(_nw52).ArraySet1(Companion_Default___.Fm13(_142_v3, (_dafny.Zero).Minus(_144_v5), _147_globalState), 5)
			(_nw52).ArraySet1(_267_v98, 6)
			(_nw52).ArraySet1((_267_v98).Merge(_267_v98), 7)
			_272_v104 = _nw52
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_272_v104), 0))
			_ = _index52
			(_272_v104).ArraySet1(_267_v98, (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_272_v104), 0))
			_ = _index53
			var _rhs55 _dafny.Int = (_dafny.Zero).Minus((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))
			_ = _rhs55
			var _rhs56 _dafny.Int = Companion_Default___.Fm1(true, _164_v17, (_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int), _147_globalState)
			_ = _rhs56
			var _rhs57 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_173_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))).Int()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(905))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_274_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_275_i14 _dafny.Int) _dafny.CodePoint {
					return _274_v0
				}
			})(_138_v0)))), _dafny.Companion_Sequence_.Concatenate(_151_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(934))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_276_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_277_i15 _dafny.Int) _dafny.CodePoint {
					return _276_v0
				}
			})(_138_v0)))))
			_ = _rhs57
			var _rhs58 _dafny.Map = _267_v98
			_ = _rhs58
			var _lhs33 *GlobalState = _147_globalState
			_ = _lhs33
			var _lhs34 _dafny.Array = _272_v104
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_272_v104), 0))
			_ = _lhs35
			_lhs33.F8 = _rhs55
			_165_v18 = _rhs56
			_151_v9 = _rhs57
			(_lhs34).ArraySet1(_rhs58, (_lhs35).Int())
			var _278_v105 _dafny.Sequence
			_ = _278_v105
			_278_v105 = _dafny.SeqOf(_172_v24, _172_v24, _172_v24, _172_v24, _dafny.ArrayCastTo((_171_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_171_v23), 0))).Int())))
			_278_v105 = _278_v105
			var _279_v106 _dafny.Sequence
			_ = _279_v106
			_279_v106 = _dafny.SeqOf((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_164_v17))
			var _280_v107 _dafny.Sequence
			_ = _280_v107
			_280_v107 = _dafny.SeqOf(_279_v106, _279_v106, _279_v106, _dafny.SeqOf(_164_v17, _dafny.IntOfUint32((_dafny.SeqOf((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))).Cardinality()), _144_v5))
			var _281_v108 _dafny.Set
			_ = _281_v108
			_281_v108 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("rjjbn"))
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))
			_ = _index54
			var _rhs59 bool = _142_v3
			_ = _rhs59
			var _rhs60 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_280_v107).Select((Companion_Default___.SafeIndex(_144_v5, _dafny.IntOfUint32((_280_v107).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf((_281_v108).Cardinality()))
			_ = _rhs60
			var _rhs61 bool = _142_v3
			_ = _rhs61
			var _rhs62 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mwwcq"), _151_v9)
			_ = _rhs62
			var _lhs36 *C0 = _270_v102
			_ = _lhs36
			var _lhs37 _dafny.Array = _173_v25
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_173_v25), 0))
			_ = _lhs38
			_142_v3 = _rhs59
			_279_v106 = _rhs60
			_lhs36.F19 = _rhs61
			(_lhs37).ArraySet1(_rhs62, (_lhs38).Int())
			(_270_v102).F19 = (false) == (_270_v102.F19)
		} else {
			_142_v3 = _142_v3
			var _282_v109 *C0
			_ = _282_v109
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__(_142_v3, _146_v7)
			_282_v109 = _nw53
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_139_v1), 0))
			_ = _index55
			(_139_v1).ArraySet1((_165_v18).Times(_164_v17), (_index55).Int())
			var _283_v110 _dafny.Map
			_ = _283_v110
			_283_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v8, _164_v17)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_139_v1), 0))
			_ = _index56
			(_139_v1).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_283_v110).Cardinality(), Companion_Default___.SafeModuloInt(_144_v5, _165_v18))), (_index56).Int())
			(_282_v109).F19 = _282_v109.F19
			(_147_globalState).F7 = ((_dafny.Zero).Minus((_139_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_139_v1), 0))).Int()).(_dafny.Int))).Minus(_164_v17)
		}
		var _284_v111 *C0
		_ = _284_v111
		var _nw54 *C0 = New_C0_()
		_ = _nw54
		_nw54.Ctor__((_145_v6).Select((Companion_Default___.SafeIndex(_164_v17, _dafny.IntOfUint32((_145_v6).Cardinality()))).Uint32()).(bool), _146_v7)
		_284_v111 = _nw54
		_146_v7 = (_146_v7).Update(_142_v3, _284_v111.F19)
		var _285_v112 _dafny.Map
		_ = _285_v112
		_285_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v0, _284_v111)
		_285_v112 = (func() _dafny.Map {
			if _142_v3 {
				return ((_285_v112).Update(_138_v0, _284_v111)).Update(_dafny.CodePoint('w'), _284_v111)
			}
			return _285_v112
		})()
		_142_v3 = _284_v111.F19
	}
	_146_v7 = (_146_v7).Update(_142_v3, _142_v3)
	_dafny.Print(_138_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v2).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_142_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v4).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_144_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_145_v6, _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F2()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F3()).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState.F4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u')), _dafny.IntOfInt64(172))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_147_globalState.F9, _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F11()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F13()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_150_v8).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_v9.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_152_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_164_v17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_165_v18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_171_v23).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v24).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v25).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_178_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_189_i10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_217_i11)
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
	Cf0 _dafny.Map
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Map) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf1 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf1 D0) D0 {
	return D0{D0_DC2{Cf1}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_()
}

func (_this D0) Dtor_cf0() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() D0 {
	return _this.Get_().(D0_DC2).Cf1
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
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf1.Equals(data2.Cf1)
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
	Cf3 _dafny.Set
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf3 _dafny.Set) D1 {
	return D1{D1_DC4{Cf3}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf2 _dafny.Array
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf2 _dafny.Array) D1 {
	return D1{D1_DC3{Cf2}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC5 struct {
	Cf4 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf4 D1) D1 {
	return D1{D1_DC5{Cf4}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.EmptySet)
}

func (_this D1) Dtor_cf3() _dafny.Set {
	return _this.Get_().(D1_DC4).Cf3
}

func (_this D1) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D1_DC3).Cf2
}

func (_this D1) Dtor_cf4() D1 {
	return _this.Get_().(D1_DC5).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf2) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf3.Equals(data2.Cf3)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf2 == data2.Cf2
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
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

type D2_DC6 struct {
	Cf5 _dafny.Map
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf5 _dafny.Map) D2 {
	return D2{D2_DC6{Cf5}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D2) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf5) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf5.Equals(data2.Cf5)
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
	Cf6 _dafny.Int
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf6 _dafny.Int) D3 {
	return D3{D3_DC7{Cf6}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Int {
	return _dafny.Zero
}

func (_this D3) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D3_DC7).Cf6
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0
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
	Cf7 bool
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf7 bool) D4 {
	return D4{D4_DC8{Cf7}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

func (CompanionStruct_D4_) Default() bool {
	return false
}

func (_this D4) Dtor_cf7() bool {
	return _this.Get_().(D4_DC8).Cf7
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf7 == data2.Cf7
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

type D5_DC10 struct {
	Cf9  _dafny.Set
	Cf10 _dafny.Map
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf9 _dafny.Set, Cf10 _dafny.Map) D5 {
	return D5{D5_DC10{Cf9, Cf10}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

type D5_DC11 struct {
	Cf11 bool
	Cf12 _dafny.Int
	Cf13 bool
	Cf14 _dafny.Array
	Cf15 _dafny.Int
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf11 bool, Cf12 _dafny.Int, Cf13 bool, Cf14 _dafny.Array, Cf15 _dafny.Int) D5 {
	return D5{D5_DC11{Cf11, Cf12, Cf13, Cf14, Cf15}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC9 struct {
	Cf8 *C0
}

func (D5_DC9) isD5() {}

func (CompanionStruct_D5_) Create_DC9_(Cf8 *C0) D5 {
	return D5{D5_DC9{Cf8}}
}

func (_this D5) Is_DC9() bool {
	_, ok := _this.Get_().(D5_DC9)
	return ok
}

type D5_DC12 struct {
	Cf16 D5
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf16 D5) D5 {
	return D5{D5_DC12{Cf16}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC10_(_dafny.EmptySet, _dafny.EmptyMap)
}

func (_this D5) Dtor_cf9() _dafny.Set {
	return _this.Get_().(D5_DC10).Cf9
}

func (_this D5) Dtor_cf10() _dafny.Map {
	return _this.Get_().(D5_DC10).Cf10
}

func (_this D5) Dtor_cf11() bool {
	return _this.Get_().(D5_DC11).Cf11
}

func (_this D5) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D5_DC11).Cf12
}

func (_this D5) Dtor_cf13() bool {
	return _this.Get_().(D5_DC11).Cf13
}

func (_this D5) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D5_DC11).Cf14
}

func (_this D5) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D5_DC11).Cf15
}

func (_this D5) Dtor_cf8() *C0 {
	return _this.Get_().(D5_DC9).Cf8
}

func (_this D5) Dtor_cf16() D5 {
	return _this.Get_().(D5_DC12).Cf16
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D5_DC9:
		{
			return "D5.DC9" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf9.Equals(data2.Cf9) && data1.Cf10.Equals(data2.Cf10)
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D5_DC9:
		{
			data2, ok := other.Get_().(D5_DC9)
			return ok && data1.Cf8 == data2.Cf8
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
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
	Cf17 _dafny.Sequence
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf17 _dafny.Sequence) D6 {
	return D6{D6_DC13{Cf17}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D6) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D6_DC13).Cf17
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + data.Cf17.VerbatimString(true) + ")"
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
			return ok && data1.Cf17.Equals(data2.Cf17)
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
	Cf19 *C0
	Cf20 _dafny.Int
	Cf21 _dafny.Int
	Cf22 _dafny.CodePoint
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf19 *C0, Cf20 _dafny.Int, Cf21 _dafny.Int, Cf22 _dafny.CodePoint) D7 {
	return D7{D7_DC15{Cf19, Cf20, Cf21, Cf22}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC16 struct {
	Cf23 _dafny.CodePoint
	Cf24 _dafny.Array
	Cf25 _dafny.Int
	Cf26 _dafny.Int
	Cf27 bool
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf23 _dafny.CodePoint, Cf24 _dafny.Array, Cf25 _dafny.Int, Cf26 _dafny.Int, Cf27 bool) D7 {
	return D7{D7_DC16{Cf23, Cf24, Cf25, Cf26, Cf27}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC17 struct {
	Cf28 bool
	Cf29 _dafny.Int
	Cf30 bool
	Cf31 bool
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf28 bool, Cf29 _dafny.Int, Cf30 bool, Cf31 bool) D7 {
	return D7{D7_DC17{Cf28, Cf29, Cf30, Cf31}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC14 struct {
	Cf18 _dafny.MultiSet
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf18 _dafny.MultiSet) D7 {
	return D7{D7_DC14{Cf18}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC15_((*C0)(nil), _dafny.Zero, _dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D7) Dtor_cf19() *C0 {
	return _this.Get_().(D7_DC15).Cf19
}

func (_this D7) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf20
}

func (_this D7) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf21
}

func (_this D7) Dtor_cf22() _dafny.CodePoint {
	return _this.Get_().(D7_DC15).Cf22
}

func (_this D7) Dtor_cf23() _dafny.CodePoint {
	return _this.Get_().(D7_DC16).Cf23
}

func (_this D7) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D7_DC16).Cf24
}

func (_this D7) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D7_DC16).Cf25
}

func (_this D7) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D7_DC16).Cf26
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC16).Cf27
}

func (_this D7) Dtor_cf28() bool {
	return _this.Get_().(D7_DC17).Cf28
}

func (_this D7) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D7_DC17).Cf29
}

func (_this D7) Dtor_cf30() bool {
	return _this.Get_().(D7_DC17).Cf30
}

func (_this D7) Dtor_cf31() bool {
	return _this.Get_().(D7_DC17).Cf31
}

func (_this D7) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D7_DC14).Cf18
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22 == data2.Cf22
		}
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27 == data2.Cf27
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	F4   _dafny.Map
	F7   _dafny.Int
	F8   _dafny.Int
	F9   _dafny.Sequence
	_f0  bool
	_f1  _dafny.Sequence
	_f2  _dafny.Map
	_f3  _dafny.Set
	_f5  _dafny.Int
	_f6  bool
	_f10 bool
	_f11 _dafny.Array
	_f12 bool
	_f13 _dafny.Map
	_f14 _dafny.Int
	_f15 _dafny.Int
	_f16 bool
	_f17 bool
	_f18 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = _dafny.EmptyMap
	_this.F7 = _dafny.Zero
	_this.F8 = _dafny.Zero
	_this.F9 = _dafny.EmptySeq
	_this._f0 = false
	_this._f1 = _dafny.EmptySeq
	_this._f2 = _dafny.EmptyMap
	_this._f3 = _dafny.EmptySet
	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this._f10 = false
	_this._f11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f12 = false
	_this._f13 = _dafny.EmptyMap
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.Zero
	_this._f16 = false
	_this._f17 = false
	_this._f18 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Sequence, f2 _dafny.Map, f3 _dafny.Set, f4 _dafny.Map, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 _dafny.Int, f9 _dafny.Sequence, f10 bool, f11 _dafny.Array, f12 bool, f13 _dafny.Map, f14 _dafny.Int, f15 _dafny.Int, f16 bool, f17 bool, f18 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this).F8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Map {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Set {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
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
func (_this *GlobalState) F13() _dafny.Map {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F19  bool
	_f20 _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F19 = false
	_this._f20 = _dafny.EmptyMap
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

func (_this *C0) Ctor__(f19 bool, f20 _dafny.Map) {
	{
		(_this).F19 = f19
		(_this)._f20 = f20
	}
}
func (_this *C0) Fm2(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()), ((_this).F20()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()))).Cardinality())).Cardinality()))).Intersection((_dafny.SetOf(_dafny.IntOfInt64(689))).Union(_dafny.SetOf(_dafny.IntOfInt64(468), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("doisvt")).Cardinality()))))).Cardinality())
	}
}
func (_this *C0) Fm3(p0 bool, p1 _dafny.Int, globalState *GlobalState) D0 {
	{
		var _source2 D0 = Companion_D0_.Create_DC1_()
		_ = _source2
		if _source2.Is_DC1() {
			return Companion_D0_.Create_DC1_()
		} else if _source2.Is_DC0() {
			var _286___mcc_h0 _dafny.Map = _source2.Get_().(D0_DC0).Cf0
			_ = _286___mcc_h0
			var _287_cf0 _dafny.Map = _286___mcc_h0
			_ = _287_cf0
			return Companion_D0_.Create_DC1_()
		} else {
			var _288___mcc_h1 D0 = _source2.Get_().(D0_DC2).Cf1
			_ = _288___mcc_h1
			var _289_cf1 D0 = _288___mcc_h1
			_ = _289_cf1
			return Companion_D0_.Create_DC1_()
		}
	}
}
func (_this *C0) F20() _dafny.Map {
	{
		return _this._f20
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
