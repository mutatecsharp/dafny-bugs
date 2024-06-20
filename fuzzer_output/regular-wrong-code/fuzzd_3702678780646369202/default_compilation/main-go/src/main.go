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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) bool {
	return !((_dafny.IntOfInt64(567)).Cmp((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('q'))).Cardinality())).Times(_dafny.IntOfUint32(((Companion_D6_.Create_DC14_(_dafny.SeqOf(true, false))).Dtor_cf20()).Cardinality()))) >= 0)
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.Int {
	return (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(151))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _0_v1 _dafny.Int
				_0_v1 = interface{}(_compr_1).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(151)), _0_v1) {
					_coll1.Add((_0_v1).Minus(_dafny.IntOfInt64(829)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(80))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg0 _dafny.Int) interface{} {
							return coer0(arg0)
						}
					}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('y')
					}))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(523))).Cardinality()))
				}
			}
			return _coll1.ToMap()
		}()).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_0).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(151))).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _0_v1 _dafny.Int
					_0_v1 = interface{}(_compr_2).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(151)), _0_v1) {
						_coll2.Add((_0_v1).Minus(_dafny.IntOfInt64(829)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(80))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg1 _dafny.Int) interface{} {
								return coer1(arg1)
							}
						}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('y')
						}))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(523))).Cardinality()))
					}
				}
				return _coll2.ToMap()
			}()).Contains(_2_v0) {
				_coll0.Add((_2_v0).Minus(_dafny.IntOfInt64(542)), _dafny.IntOfInt64(921))
			}
		}
		return _coll0.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(627), _dafny.IntOfInt64(464)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(638)))).Cardinality()), _dafny.IntOfInt64(126)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-220), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(487))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))).Cardinality()
	}))).Cardinality())))).Cardinality()), _dafny.IntOfInt64(-511))).Plus(_dafny.IntOfInt64(535))
}
func (_static *CompanionStruct_Default___) Fm4(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(462))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rnvmueh")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(614), _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	}))))).Cardinality())).Cardinality()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(239))).Cardinality()), _dafny.IntOfInt64(-650), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality()))).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-585)), _dafny.SeqOf(_dafny.IntOfInt64(913), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("f")).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(301), false)).Cardinality()))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-620), (_dafny.MultiSetOf(false)).Cardinality())).Cardinality())).Cardinality()), true))).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(963), _dafny.IntOfInt64(663))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(963)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(663)) < 0) {
				_coll3.Add(Companion_Default___.SafeModuloInt(_5_v0, _dafny.IntOfInt64(441)), false)
			}
		}
		return _coll3.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(false, false, true)
		}
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(true, false, false, !(false), false))
	})()).Union(_dafny.MultiSetOf(!(true)))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.IntOfInt64(90)).Cmp((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qjl")).Cardinality()), _dafny.IntOfInt64(-53))).Cardinality())) >= 0)
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("usw")
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(335), _dafny.IntOfInt64(539))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _6_v0 _dafny.Int
			_6_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(335)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(539)) < 0) {
				_coll4.Add((_6_v0).Times((_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(129))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))).Cardinality())))
			}
		}
		return _coll4.ToMap()
	}()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(186), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())), !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(true, true, !(true), true, true), true))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xqmuuiom"), _dafny.UnicodeSeqOfUtf8Bytes("sjc"), _dafny.UnicodeSeqOfUtf8Bytes("f"))).Intersection(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ofdbvnl")))).Union((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("bc"))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(379))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_8_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(42))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_9_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))
	})))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 D7, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('v')))).Difference(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('c')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('s')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('s'))))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) {
	var _10_v0 bool
	_ = _10_v0
	_10_v0 = false
	var _11_v1 _dafny.Int
	_ = _11_v1
	_11_v1 = _dafny.IntOfInt64(-798)
	var _12_v2 _dafny.Map
	_ = _12_v2
	_12_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_10_v0, _10_v0)
	var _13_v3 _dafny.Sequence
	_ = _13_v3
	_13_v3 = _dafny.SeqOf(!(_10_v0))
	var _14_v4 _dafny.Set
	_ = _14_v4
	_14_v4 = _dafny.SetOf(_11_v1, (_dafny.Zero).Minus(_11_v1))
	var _15_v5 _dafny.Set
	_ = _15_v5
	_15_v5 = _dafny.SetOf(_10_v0)
	var _16_v6 _dafny.Array
	_ = _16_v6
	var _nwElement0_0 _dafny.Int = _11_v1
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(22))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1((_12_v2).Cardinality(), 1)
	(_nw0).ArraySet1((_dafny.Zero).Minus(_11_v1), 2)
	(_nw0).ArraySet1(_dafny.IntOfUint32((_13_v3).Cardinality()), 3)
	(_nw0).ArraySet1((_14_v4).Cardinality(), 4)
	(_nw0).ArraySet1(_11_v1, 5)
	(_nw0).ArraySet1(_11_v1, 6)
	(_nw0).ArraySet1(_11_v1, 7)
	(_nw0).ArraySet1(_11_v1, 8)
	(_nw0).ArraySet1(_11_v1, 9)
	(_nw0).ArraySet1(_11_v1, 10)
	(_nw0).ArraySet1(_11_v1, 11)
	(_nw0).ArraySet1(_11_v1, 12)
	(_nw0).ArraySet1(_11_v1, 13)
	(_nw0).ArraySet1(_11_v1, 14)
	(_nw0).ArraySet1((_15_v5).Cardinality(), 15)
	(_nw0).ArraySet1(_dafny.IntOfUint32((_13_v3).Cardinality()), 16)
	(_nw0).ArraySet1(_dafny.IntOfInt64(-292), 17)
	(_nw0).ArraySet1(_11_v1, 18)
	(_nw0).ArraySet1(_11_v1, 19)
	(_nw0).ArraySet1(_11_v1, 20)
	(_nw0).ArraySet1(_11_v1, 21)
	_16_v6 = _nw0
	var _17_v7 *C0
	_ = _17_v7
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__(_10_v0, (func() _dafny.Array {
		if _10_v0 {
			return _16_v6
		}
		return _16_v6
	})())
	_17_v7 = _nw1
	var _18_i0 _dafny.Int
	_ = _18_i0
	_18_i0 = _dafny.Zero
	{
		for (_17_v7).F28() {
			{
				if (_18_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_18_i0 = (_18_i0).Plus(_dafny.One)
				if Companion_Default___.Fm0(globalState) {
					var _19_v8 _dafny.CodePoint
					_ = _19_v8
					_19_v8 = _dafny.CodePoint('d')
					_19_v8 = (func() _dafny.CodePoint {
						if _10_v0 {
							return _19_v8
						}
						return _19_v8
					})()
					var _20_v9 _dafny.Array
					_ = _20_v9
					var _nwElement0_1 bool = _10_v0
					_ = _nwElement0_1
					var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(19))
					_ = _nw2
					(_nw2).ArraySet1(_nwElement0_1, 0)
					(_nw2).ArraySet1((_17_v7).F28(), 1)
					(_nw2).ArraySet1(_10_v0, 2)
					(_nw2).ArraySet1((_17_v7).F28(), 3)
					(_nw2).ArraySet1((_17_v7).F28(), 4)
					(_nw2).ArraySet1((_17_v7).Fm2((_17_v7).F28(), globalState), 5)
					(_nw2).ArraySet1((_17_v7).F28(), 6)
					(_nw2).ArraySet1((_17_v7).F28(), 7)
					(_nw2).ArraySet1(!((_17_v7).F28()), 8)
					(_nw2).ArraySet1(_10_v0, 9)
					(_nw2).ArraySet1((_17_v7).F28(), 10)
					(_nw2).ArraySet1(_10_v0, 11)
					(_nw2).ArraySet1((_17_v7).F28(), 12)
					(_nw2).ArraySet1((_17_v7).F28(), 13)
					(_nw2).ArraySet1((_17_v7).F28(), 14)
					(_nw2).ArraySet1((_17_v7).F28(), 15)
					(_nw2).ArraySet1(!((_17_v7).F28()), 16)
					(_nw2).ArraySet1(_10_v0, 17)
					(_nw2).ArraySet1((_17_v7).F28(), 18)
					_20_v9 = _nw2
					var _21_v10 _dafny.Array
					_ = _21_v10
					_21_v10 = _20_v9
					var _22_v11 _dafny.Sequence
					_ = _22_v11
					_22_v11 = _dafny.SeqOf(_20_v9, _20_v9, _20_v9, _20_v9, _20_v9)
					var _23_v12 _dafny.Sequence
					_ = _23_v12
					_23_v12 = _dafny.UnicodeSeqOfUtf8Bytes("btlenkt")
					var _24_v13 _dafny.Set
					_ = _24_v13
					_24_v13 = _dafny.SetOf(_20_v9, (_21_v10), _20_v9, (_22_v11).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_23_v12).Cardinality())), _dafny.IntOfUint32((_22_v11).Cardinality()))).Uint32()).(_dafny.Array))
					var _25_v14 D0
					_ = _25_v14
					_25_v14 = Companion_D0_.Create_DC1_((_17_v7).F28(), _11_v1)
					var _26_v15 D0
					_ = _26_v15
					_26_v15 = Companion_D0_.Create_DC2_(_25_v14)
					var _27_v16 _dafny.MultiSet
					_ = _27_v16
					_27_v16 = _dafny.MultiSetOf(_10_v0, !((_17_v7).F28()))
					var _28_v17 _dafny.Map
					_ = _28_v17
					_28_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _20_v9)
					var _29_v18 _dafny.Sequence
					_ = _29_v18
					_29_v18 = _dafny.SeqOf(_dafny.SetOf((func() _dafny.Array {
						if (_28_v17).Contains((_dafny.SetOf(_11_v1, _11_v1, (_15_v5).Cardinality(), _11_v1)).Cardinality()) {
							return (_28_v17).Get((_dafny.SetOf(_11_v1, _11_v1, (_15_v5).Cardinality(), _11_v1)).Cardinality()).(_dafny.Array)
						}
						return _20_v9
					})()), _24_v13, (_24_v13).Difference(_24_v13), _dafny.SetOf(_20_v9), (_24_v13).Union(_24_v13))
					var _30_v19 D1
					_ = _30_v19
					_30_v19 = Companion_D1_.Create_DC4_((_17_v7).F28())
					var _rhs0 _dafny.Set = (_29_v18).Select((Companion_Default___.SafeIndex(_11_v1, _dafny.IntOfUint32((_29_v18).Cardinality()))).Uint32()).(_dafny.Set)
					_ = _rhs0
					var _rhs1 D0 = _26_v15
					_ = _rhs1
					var _rhs2 _dafny.Array = _20_v9
					_ = _rhs2
					var _rhs3 _dafny.MultiSet = ((_27_v16).Union(_27_v16)).Intersection(Companion_Default___.Fm6((_30_v19).Dtor_cf5(), _11_v1, globalState))
					_ = _rhs3
					_24_v13 = _rhs0
					_26_v15 = _rhs1
					_20_v9 = _rhs2
					_27_v16 = _rhs3
					var _31_v20 *C0
					_ = _31_v20
					var _nw3 *C0 = New_C0_()
					_ = _nw3
					_nw3.Ctor__((_17_v7).F28(), (_17_v7).F29())
					_31_v20 = _nw3
					var _32_v21 _dafny.Map
					_ = _32_v21
					_32_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v7).F28(), _13_v3)
					var _33_v22 _dafny.Sequence
					_ = _33_v22
					_33_v22 = _dafny.SeqOf(_dafny.SeqOf((_17_v7).F28()), _13_v3, _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
						if (_32_v21).Contains(_10_v0) {
							return (_32_v21).Get(_10_v0).(_dafny.Sequence)
						}
						return _13_v3
					})(), _13_v3), _dafny.SeqOf((_31_v20).F28()))
					var _34_v23 _dafny.Map
					_ = _34_v23
					_34_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _11_v1)
					var _35_v24 D3
					_ = _35_v24
					_35_v24 = Companion_D3_.Create_DC10_((_17_v7).Fm2((_17_v7).F28(), globalState), (_34_v23).Cardinality(), !((_17_v7).F28()), _dafny.IntOfInt64(282), (_17_v7).F29())
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_20_v9), 0))
					_ = _index0
					(_20_v9).ArraySet1((_35_v24).Dtor_cf14(), (_index0).Int())
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_20_v9), 0))
					_ = _index1
					var _rhs4 bool = false
					_ = _rhs4
					var _rhs5 bool = (_11_v1).Cmp((_11_v1).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _dafny.IntOfInt64(-924))).Cardinality())) > 0
					_ = _rhs5
					var _rhs6 _dafny.Sequence = _33_v22
					_ = _rhs6
					var _rhs7 bool = !((_17_v7).F28())
					_ = _rhs7
					var _lhs0 *GlobalState = globalState
					_ = _lhs0
					var _lhs1 *GlobalState = globalState
					_ = _lhs1
					var _lhs2 _dafny.Array = _20_v9
					_ = _lhs2
					var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_20_v9), 0))
					_ = _lhs3
					_lhs0.F0 = _rhs4
					_lhs1.F0 = _rhs5
					_33_v22 = _rhs6
					(_lhs2).ArraySet1(_rhs7, (_lhs3).Int())
					var _36_v25 D2
					_ = _36_v25
					_36_v25 = Companion_D2_.Create_DC6_(_27_v16)
					var _37_v26 _dafny.Map
					_ = _37_v26
					_37_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D2 {
						if (_31_v20).Fm2((_17_v7).F28(), globalState) {
							return _36_v25
						}
						return _36_v25
					})(), _dafny.IntOfInt64(841))
					_37_v26 = (_37_v26).Update(Companion_D2_.Create_DC6_(_27_v16), _11_v1)
				} else {
					var _38_v27 _dafny.MultiSet
					_ = _38_v27
					_38_v27 = _dafny.MultiSetOf((_11_v1).Cmp(_11_v1) <= 0, (_11_v1).Cmp(_11_v1) == 0)
					var _rhs8 _dafny.Int = _11_v1
					_ = _rhs8
					var _rhs9 _dafny.Int = (func() _dafny.Int {
						if (_38_v27).Contains((_17_v7).F28()) {
							return (_38_v27).Multiplicity((_17_v7).F28())
						}
						return _dafny.IntOfInt64(591)
					})()
					_ = _rhs9
					var _rhs10 _dafny.Array = (_17_v7).F29()
					_ = _rhs10
					_11_v1 = _rhs8
					_11_v1 = _rhs9
					_16_v6 = _rhs10
					var _39_v28 _dafny.Array
					_ = _39_v28
					var _nw4 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
					_ = _nw4
					_39_v28 = _nw4
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_39_v28), 0))
					_ = _index2
					(_39_v28).ArraySet1((_17_v7).F28(), (_index2).Int())
					var _40_v29 _dafny.CodePoint
					_ = _40_v29
					_40_v29 = _dafny.CodePoint('o')
					var _41_v30 _dafny.Sequence
					_ = _41_v30
					_41_v30 = _dafny.SeqOf(_40_v29, _40_v29, _40_v29)
					var _42_v31 D0
					_ = _42_v31
					_42_v31 = Companion_D0_.Create_DC1_(_10_v0, _dafny.IntOfInt64(-832))
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_39_v28), 0))
					_ = _index3
					var _rhs11 _dafny.Int = (_dafny.IntOfUint32((_41_v30).Cardinality())).Plus((_42_v31).Dtor_cf2())
					_ = _rhs11
					var _rhs12 bool = (_17_v7).F28()
					_ = _rhs12
					var _rhs13 bool = (func() bool {
						if (_17_v7).F28() {
							return (func() _dafny.Set {
								var _coll5 = _dafny.NewBuilder()
								_ = _coll5
								for _iter5 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(944))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg7 _dafny.Int) interface{} {
										return coer7(arg7)
									}
								}((func(_43_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
									return func(_44_i1 _dafny.Int) _dafny.Int {
										return _43_v1
									}
								})(_11_v1)))).Elements()); ; {
									_compr_5, _ok5 := _iter5()
									if !_ok5 {
										break
									}
									var _45_v32 _dafny.Int
									_45_v32 = interface{}(_compr_5).(_dafny.Int)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(944))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
										return func(arg8 _dafny.Int) interface{} {
											return coer8(arg8)
										}
									}((func(_46_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
										return func(_44_i1 _dafny.Int) _dafny.Int {
											return _46_v1
										}
									})(_11_v1))), _45_v32) {
										_coll5.Add(Companion_Default___.SafeModuloInt(_45_v32, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sbrnrtwv")).Cardinality())))
									}
								}
								return _coll5.ToSet()
							}()).Contains(_11_v1)
						}
						return (_11_v1).Cmp(_11_v1) < 0
					})()
					_ = _rhs13
					var _lhs4 _dafny.Array = _39_v28
					_ = _lhs4
					var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_39_v28), 0))
					_ = _lhs5
					_11_v1 = _rhs11
					(_lhs4).ArraySet1(_rhs12, (_lhs5).Int())
					_10_v0 = _rhs13
					_40_v29 = _40_v29
					var _47_v33 _dafny.MultiSet
					_ = _47_v33
					_47_v33 = _dafny.MultiSetOf(_17_v7, _17_v7, _17_v7)
					_11_v1 = (func() _dafny.Int {
						if (_47_v33).Contains(_17_v7) {
							return (_47_v33).Multiplicity(_17_v7)
						}
						return (_11_v1).Plus(_11_v1)
					})()
					var _48_v34 _dafny.Map
					_ = _48_v34
					_48_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v30, _10_v0)
					var _49_v35 _dafny.Array
					_ = _49_v35
					var _nwElement0_2 _dafny.Map = (_12_v2).Update(_10_v0, false)
					_ = _nwElement0_2
					var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(7))
					_ = _nw5
					(_nw5).ArraySet1(_nwElement0_2, 0)
					(_nw5).ArraySet1(_12_v2, 1)
					(_nw5).ArraySet1(_12_v2, 2)
					(_nw5).ArraySet1((_12_v2).Merge(Companion_Default___.Fm7((_17_v7).F28(), (_17_v7).F28(), (_39_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_39_v28), 0))).Int()).(bool), _11_v1, globalState)), 3)
					(_nw5).ArraySet1(_12_v2, 4)
					(_nw5).ArraySet1((Companion_Default___.Fm7(_10_v0, Companion_Default___.Fm0(globalState), (_17_v7).F28(), (_48_v34).Cardinality(), globalState)).Update(_10_v0, _10_v0), 5)
					(_nw5).ArraySet1(_12_v2, 6)
					_49_v35 = _nw5
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_49_v35), 0))
					_ = _index4
					(_49_v35).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v7).F28(), (_17_v7).Fm2((_17_v7).F28(), globalState))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_39_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_39_v28), 0))).Int()).(bool), (_17_v7).F28())), (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_49_v35), 0))
					_ = _index5
					(_49_v35).ArraySet1(_12_v2, (_index5).Int())
				}
				if !(_10_v0) || (!((_17_v7).F28()) || (!((_17_v7).F28()))) {
					var _50_v36 _dafny.Sequence
					_ = _50_v36
					_50_v36 = _dafny.UnicodeSeqOfUtf8Bytes("vithblii")
					(globalState).F0 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_50_v36, _50_v36), _dafny.CodePoint('i'))
					var _51_v37 _dafny.Array
					_ = _51_v37
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(7)
					_ = _len0_0
					var _nw6 _dafny.Array
					_ = _nw6
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw6 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) _dafny.CodePoint = func(_52_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('y')
						}
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw6 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw6).ArraySet1CodePoint(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw6).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_51_v37 = _nw6
					var _53_v38 _dafny.Sequence
					_ = _53_v38
					_53_v38 = _50_v36
					var _rhs14 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}(func(_54_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					})), _dafny.Companion_Sequence_.Concatenate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(870))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg10 _dafny.Int) interface{} {
							return coer10(arg10)
						}
					}(func(_55_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('x')
					}))), _50_v36))
					_ = _rhs14
					var _rhs15 _dafny.Array = (func() _dafny.Array {
						if (_17_v7).F28() {
							return _51_v37
						}
						return (func() _dafny.Array {
							if (_17_v7).F28() {
								return _51_v37
							}
							return _51_v37
						})()
					})()
					_ = _rhs15
					var _lhs6 *GlobalState = globalState
					_ = _lhs6
					_lhs6.F0 = _rhs14
					_51_v37 = _rhs15
					_10_v0 = _10_v0
					_11_v1 = (Companion_Default___.SafeDivisionInt(_11_v1, _11_v1)).Minus(Companion_Default___.Fm3(_11_v1, (_17_v7).F28(), _11_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, (_17_v7).F28()), globalState))
					var _56_v39 _dafny.Array
					_ = _56_v39
					var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(23))
					_ = _nw7
					_56_v39 = _nw7
					var _57_v40 _dafny.Map
					_ = _57_v40
					_57_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _17_v7)
					var _58_v41 _dafny.Map
					_ = _58_v41
					_58_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v40).Cardinality(), Companion_Default___.Fm8(_11_v1, globalState))
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_56_v39), 0))
					_ = _index6
					(_56_v39).ArraySet1((_dafny.SetOf((_17_v7).F28(), _10_v0, _10_v0, (_17_v7).F28(), (_17_v7).F28())).Difference((func() _dafny.Set {
						if (_58_v41).Contains(_dafny.IntOfInt64(330)) {
							return (_58_v41).Get(_dafny.IntOfInt64(330)).(_dafny.Set)
						}
						return Companion_Default___.Fm8(_11_v1, globalState)
					})()), (_index6).Int())
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_56_v39), 0))
					_ = _index7
					(_56_v39).ArraySet1((_15_v5).Difference((_15_v5).Union(_15_v5)), (_index7).Int())
				} else {
					var _59_v42 _dafny.CodePoint
					_ = _59_v42
					_59_v42 = _dafny.CodePoint('l')
					var _60_v43 D3
					_ = _60_v43
					_60_v43 = Companion_D3_.Create_DC10_(_10_v0, _11_v1, (_17_v7).F28(), _11_v1, (_17_v7).F29())
					var _61_v44 _dafny.Map
					_ = _61_v44
					_61_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v43).Dtor_cf14(), _17_v7)
					var _62_v45 _dafny.Sequence
					_ = _62_v45
					_62_v45 = _dafny.SeqOf((func() *C0 {
						if (_61_v44).Contains(!((_17_v7).Fm2((_17_v7).F28(), globalState))) {
							return (_61_v44).Get(!((_17_v7).Fm2((_17_v7).F28(), globalState))).(*C0)
						}
						return _17_v7
					})(), _17_v7, _17_v7)
					var _63_v46 _dafny.Sequence
					_ = _63_v46
					_63_v46 = _dafny.UnicodeSeqOfUtf8Bytes("squrxyb")
					var _64_v47 _dafny.Sequence
					_ = _64_v47
					_64_v47 = _63_v46
					var _65_v48 _dafny.Map
					_ = _65_v48
					_65_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v47, false)
					var _66_v49 D6
					_ = _66_v49
					_66_v49 = Companion_D6_.Create_DC13_(_59_v42)
					var _pat_let_tv0 = _59_v42
					_ = _pat_let_tv0
					var _67_v50 D6
					_ = _67_v50
					_67_v50 = Companion_D6_.Create_DC13_((func(_pat_let0_0 D6) D6 {
						return func(_68_dt__update__tmp_h1 D6) D6 {
							return func(_pat_let1_0 _dafny.CodePoint) D6 {
								return func(_69_dt__update_hcf19_h0 _dafny.CodePoint) D6 {
									return Companion_D6_.Create_DC13_(_69_dt__update_hcf19_h0)
								}(_pat_let1_0)
							}(_pat_let_tv0)
						}(_pat_let0_0)
					}(_66_v49)).Dtor_cf19())
					var _rhs16 _dafny.Int = _11_v1
					_ = _rhs16
					var _rhs17 *C0 = (_62_v45).Select((Companion_Default___.SafeIndex(((_65_v48).Merge(_65_v48)).Cardinality(), _dafny.IntOfUint32((_62_v45).Cardinality()))).Uint32()).(*C0)
					_ = _rhs17
					var _rhs18 _dafny.CodePoint = (_66_v49).Dtor_cf19()
					_ = _rhs18
					_11_v1 = _rhs16
					_17_v7 = _rhs17
					_59_v42 = _rhs18
					var _70_v51 _dafny.Array
					_ = _70_v51
					var _nwElement0_3 bool = (_17_v7).F28()
					_ = _nwElement0_3
					var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(5))
					_ = _nw8
					(_nw8).ArraySet1(_nwElement0_3, 0)
					(_nw8).ArraySet1((_17_v7).F28(), 1)
					(_nw8).ArraySet1(_10_v0, 2)
					(_nw8).ArraySet1((_17_v7).F28(), 3)
					(_nw8).ArraySet1(_10_v0, 4)
					_70_v51 = _nw8
					var _71_v52 _dafny.Array
					_ = _71_v52
					var _nwElement0_4 _dafny.Array = (_17_v7).F29()
					_ = _nwElement0_4
					var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(18))
					_ = _nw9
					(_nw9).ArraySet1(_nwElement0_4, 0)
					(_nw9).ArraySet1(_16_v6, 1)
					(_nw9).ArraySet1((_17_v7).F29(), 2)
					(_nw9).ArraySet1((_17_v7).F29(), 3)
					(_nw9).ArraySet1((_17_v7).F29(), 4)
					(_nw9).ArraySet1((_17_v7).F29(), 5)
					(_nw9).ArraySet1((_17_v7).F29(), 6)
					(_nw9).ArraySet1((_17_v7).F29(), 7)
					(_nw9).ArraySet1((_17_v7).F29(), 8)
					(_nw9).ArraySet1((_17_v7).F29(), 9)
					(_nw9).ArraySet1((_17_v7).F29(), 10)
					(_nw9).ArraySet1((_17_v7).F29(), 11)
					(_nw9).ArraySet1((_17_v7).F29(), 12)
					(_nw9).ArraySet1((_17_v7).F29(), 13)
					(_nw9).ArraySet1(_16_v6, 14)
					(_nw9).ArraySet1((_17_v7).F29(), 15)
					(_nw9).ArraySet1((_17_v7).F29(), 16)
					(_nw9).ArraySet1(_16_v6, 17)
					_71_v52 = _nw9
					var _72_v53 _dafny.Map
					_ = _72_v53
					_72_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v51, _71_v52)
					_72_v53 = (_72_v53).Update(_70_v51, _71_v52)
					var _73_v54 *C0
					_ = _73_v54
					var _nw10 *C0 = New_C0_()
					_ = _nw10
					_nw10.Ctor__(_10_v0, (_17_v7).F29())
					_73_v54 = _nw10
					var _74_v55 _dafny.MultiSet
					_ = _74_v55
					_74_v55 = _dafny.MultiSetOf(_11_v1)
					var _75_v56 _dafny.Map
					_ = _75_v56
					_75_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-735))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}((func(_76_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_77_i5 _dafny.Int) _dafny.CodePoint {
							return _76_v42
						}
					})(_59_v42))), (_74_v55).Cardinality())
					_75_v56 = (_75_v56).Update(Companion_Default___.Fm9(globalState), _11_v1)
					var _78_v57 _dafny.Map
					_ = _78_v57
					_78_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(275), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("v"), _63_v46)).Cardinality()))
					var _79_v58 _dafny.Map
					_ = _79_v58
					_79_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, (_17_v7).F28())
					var _80_v59 _dafny.Sequence
					_ = _80_v59
					_80_v59 = _dafny.SeqOf(_79_v58)
					_78_v57 = (_78_v57).Update((func() _dafny.Int {
						if (_13_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.IntOfUint32((_13_v3).Cardinality()))).Uint32()).(bool) {
							return _11_v1
						}
						return Companion_Default___.Fm3(_11_v1, (_17_v7).F28(), _11_v1, (_80_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_63_v46).Cardinality()), _dafny.IntOfUint32((_80_v59).Cardinality()))).Uint32()).(_dafny.Map), globalState)
					})(), _11_v1)
				}
				_11_v1 = (func() _dafny.Int {
					if (_17_v7).Fm2(_10_v0, globalState) {
						return _11_v1
					}
					return _11_v1
				})()
				(globalState).F0 = (_17_v7).F28()
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _81_v60 _dafny.MultiSet
	_ = _81_v60
	_81_v60 = _dafny.MultiSetOf(_10_v0, (_17_v7).F28(), _10_v0, (_17_v7).F28())
	var _82_v61 _dafny.Array
	_ = _82_v61
	var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
	_ = _nw11
	_82_v61 = _nw11
	var _83_v62 _dafny.Map
	_ = _83_v62
	_83_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(618), _82_v61)
	var _84_v63 _dafny.Sequence
	_ = _84_v63
	_84_v63 = _dafny.SeqOf(_83_v62, _83_v62, _83_v62)
	var _rhs19 bool = (_81_v60).IsDisjointFrom(_81_v60)
	_ = _rhs19
	var _rhs20 bool = true
	_ = _rhs20
	var _rhs21 bool = (_17_v7).F28()
	_ = _rhs21
	var _rhs22 _dafny.Int = ((((_84_v63).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_11_v1, _11_v1)).Cardinality()), _dafny.IntOfUint32((_84_v63).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_83_v62)).Cardinality()).Minus(_11_v1)
	_ = _rhs22
	var _lhs7 *GlobalState = globalState
	_ = _lhs7
	var _lhs8 *GlobalState = globalState
	_ = _lhs8
	_lhs7.F0 = _rhs19
	_10_v0 = _rhs20
	_lhs8.F0 = _rhs21
	_11_v1 = _rhs22
	if !((_17_v7).F28()) {
		_11_v1 = _11_v1
		(globalState).F0 = (_17_v7).F28()
		var _85_v64 _dafny.Sequence
		_ = _85_v64
		_85_v64 = _dafny.SeqOf(_dafny.SeqOf(_10_v0))
		var _86_v65 D0
		_ = _86_v65
		_86_v65 = Companion_D0_.Create_DC1_(_10_v0, _dafny.IntOfUint32((_85_v64).Cardinality()))
		if (_86_v65).Dtor_cf1() {
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_82_v61), 0))
			_ = _index8
			(_82_v61).ArraySet1((_17_v7).F28(), (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_82_v61), 0))
			_ = _index9
			(_82_v61).ArraySet1(((_17_v7).F28()) == ((_17_v7).F28()), (_index9).Int())
			var _87_v66 _dafny.Array
			_ = _87_v66
			var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(26))
			_ = _nw12
			_87_v66 = _nw12
			var _88_v67 _dafny.Set
			_ = _88_v67
			_88_v67 = _dafny.SetOf(_87_v66)
			(globalState).F0 = (_dafny.SetOf(_87_v66)).IsSubsetOf(_88_v67)
			var _89_v68 D6
			_ = _89_v68
			_89_v68 = Companion_D6_.Create_DC14_(Companion_Default___.Fm10(_11_v1, _dafny.IntOfInt64(673), _11_v1, globalState))
			var _90_v69 D1
			_ = _90_v69
			_90_v69 = Companion_D1_.Create_DC3_((_17_v7).F29())
			var _91_v70 _dafny.Map
			_ = _91_v70
			_91_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_89_v68).Dtor_cf20()).Cardinality()), _90_v69)
			var _92_v71 D1
			_ = _92_v71
			_92_v71 = Companion_D1_.Create_DC5_((func() D1 {
				if (_91_v70).Contains(_11_v1) {
					return (_91_v70).Get(_11_v1).(D1)
				}
				return _90_v69
			})())
			var _93_v72 D6
			_ = _93_v72
			_93_v72 = Companion_D6_.Create_DC16_(_dafny.SeqOf(_92_v71), _dafny.UnicodeSeqOfUtf8Bytes("am"))
			var _94_v73 _dafny.Sequence
			_ = _94_v73
			_94_v73 = _dafny.UnicodeSeqOfUtf8Bytes("jrtiwitvl")
			var _pat_let_tv1 = _90_v69
			_ = _pat_let_tv1
			var _95_v74 _dafny.Sequence
			_ = _95_v74
			_95_v74 = _dafny.SeqOf(func(_pat_let2_0 D1) D1 {
				return func(_96_dt__update__tmp_h2 D1) D1 {
					return func(_pat_let3_0 D1) D1 {
						return func(_97_dt__update_hcf6_h0 D1) D1 {
							return Companion_D1_.Create_DC5_(_97_dt__update_hcf6_h0)
						}(_pat_let3_0)
					}(_pat_let_tv1)
				}(_pat_let2_0)
			}(_92_v71), _92_v71)
			var _98_v75 _dafny.CodePoint
			_ = _98_v75
			_98_v75 = _dafny.CodePoint('i')
			var _pat_let_tv2 = _94_v73
			_ = _pat_let_tv2
			var _pat_let_tv3 = _95_v74
			_ = _pat_let_tv3
			var _pat_let_tv4 = _94_v73
			_ = _pat_let_tv4
			var _pat_let_tv5 = _12_v2
			_ = _pat_let_tv5
			var _pat_let_tv6 = _94_v73
			_ = _pat_let_tv6
			var _pat_let_tv7 = _98_v75
			_ = _pat_let_tv7
			var _99_v76 _dafny.Array
			_ = _99_v76
			var _nwElement0_5 D6 = _93_v72
			_ = _nwElement0_5
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(10))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_5, 0)
			(_nw13).ArraySet1(func(_pat_let4_0 D6) D6 {
				return func(_102_dt__update__tmp_h4 D6) D6 {
					return func(_pat_let7_0 _dafny.Sequence) D6 {
						return func(_103_dt__update_hcf24_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC16_(_103_dt__update_hcf24_h0, (_102_dt__update__tmp_h4).Dtor_cf25())
						}(_pat_let7_0)
					}(_pat_let_tv3)
				}(_pat_let4_0)
			}(func(_pat_let5_0 D6) D6 {
				return func(_100_dt__update__tmp_h3 D6) D6 {
					return func(_pat_let6_0 _dafny.Sequence) D6 {
						return func(_101_dt__update_hcf25_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC16_((_100_dt__update__tmp_h3).Dtor_cf24(), _101_dt__update_hcf25_h0)
						}(_pat_let6_0)
					}(_pat_let_tv2)
				}(_pat_let5_0)
			}(_93_v72)), 1)
			(_nw13).ArraySet1(_93_v72, 2)
			(_nw13).ArraySet1(_93_v72, 3)
			(_nw13).ArraySet1(_93_v72, 4)
			(_nw13).ArraySet1(_93_v72, 5)
			(_nw13).ArraySet1(func(_pat_let8_0 D6) D6 {
				return func(_104_dt__update__tmp_h5 D6) D6 {
					return func(_pat_let9_0 _dafny.Sequence) D6 {
						return func(_105_dt__update_hcf25_h1 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC16_((_104_dt__update__tmp_h5).Dtor_cf24(), _105_dt__update_hcf25_h1)
						}(_pat_let9_0)
					}(_dafny.Companion_Sequence_.Update(_pat_let_tv4, (Companion_Default___.SafeIndex((_pat_let_tv5).Cardinality(), _dafny.IntOfUint32((_pat_let_tv6).Cardinality()))).Uint32(), _pat_let_tv7))
				}(_pat_let8_0)
			}(_93_v72), 6)
			(_nw13).ArraySet1(_93_v72, 7)
			(_nw13).ArraySet1(_93_v72, 8)
			(_nw13).ArraySet1(_93_v72, 9)
			_99_v76 = _nw13
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_99_v76), 0))
			_ = _index10
			(_99_v76).ArraySet1(Companion_D6_.Create_DC16_(_95_v74, _94_v73), (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_99_v76), 0))
			_ = _index11
			(_99_v76).ArraySet1((func() D6 {
				if (_17_v7).F28() {
					return _93_v72
				}
				return _93_v72
			})(), (_index11).Int())
			(globalState).F0 = (_17_v7).F28()
			_93_v72 = _93_v72
		} else {
			var _106_v77 _dafny.Set
			_ = _106_v77
			_106_v77 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("xsxaj"))
			(globalState).F0 = (_106_v77).IsSubsetOf(_106_v77)
			var _107_v78 _dafny.MultiSet
			_ = _107_v78
			_107_v78 = _dafny.MultiSetOf(_dafny.IntOfInt64(777))
			var _108_v79 _dafny.Map
			_ = _108_v79
			_108_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_11_v1).Times((_14_v4).Cardinality()), _107_v78)
			_107_v78 = (func() _dafny.MultiSet {
				if (_108_v79).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ysstjicth"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-758))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}(func(_109_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})))).Cardinality())) {
					return (_108_v79).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ysstjicth"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-758))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg13 _dafny.Int) interface{} {
							return coer13(arg13)
						}
					}(func(_110_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					})))).Cardinality())).(_dafny.MultiSet)
				}
				return _107_v78
			})()
			var _111_v80 _dafny.Sequence
			_ = _111_v80
			_111_v80 = _dafny.UnicodeSeqOfUtf8Bytes("ipw")
			var _112_v81 _dafny.Map
			_ = _112_v81
			_112_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _111_v80)
			var _rhs23 _dafny.Sequence = Companion_Default___.Fm10((_dafny.IntOfInt64(-680)).Plus(_11_v1), _11_v1, _11_v1, globalState)
			_ = _rhs23
			var _rhs24 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_111_v80, (func() _dafny.Sequence {
				if (_112_v81).Contains(_11_v1) {
					return (_112_v81).Get(_11_v1).(_dafny.Sequence)
				}
				return _111_v80
			})()), _111_v80)
			_ = _rhs24
			var _rhs25 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(491), _11_v1)
			_ = _rhs25
			var _rhs26 bool = (_17_v7).F28()
			_ = _rhs26
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			var _lhs10 *GlobalState = globalState
			_ = _lhs10
			_13_v3 = _rhs23
			_lhs9.F2 = _rhs24
			_11_v1 = _rhs25
			_lhs10.F0 = _rhs26
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index12
			((_17_v7).F29()).ArraySet1(_11_v1, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index13
			((_17_v7).F29()).ArraySet1(_11_v1, (_index13).Int())
			var _113_v82 _dafny.Map
			_ = _113_v82
			_113_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_v7, (_17_v7).F29())
			_113_v82 = (_113_v82).Update(_17_v7, (_17_v7).F29())
		}
		var _114_v83 _dafny.CodePoint
		_ = _114_v83
		_114_v83 = _dafny.CodePoint('p')
		_114_v83 = _114_v83
		var _115_v84 _dafny.Array
		_ = _115_v84
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
		_ = _nw14
		_115_v84 = _nw14
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_115_v84), 0))
		_ = _index14
		(_115_v84).ArraySet1(func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-913), _dafny.IntOfInt64(328))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _116_v85 _dafny.Int
				_116_v85 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(-913)).Cmp(_116_v85) <= 0) && ((_116_v85).Cmp(_dafny.IntOfInt64(328)) < 0) {
					_coll6.Add((_116_v85).Minus((_dafny.MultiSetOf(_11_v1, _11_v1)).Cardinality()))
				}
			}
			return _coll6.ToSet()
		}(), (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_115_v84), 0))
		_ = _index15
		(_115_v84).ArraySet1(_14_v4, (_index15).Int())
	} else {
		var _117_v86 _dafny.Map
		_ = _117_v86
		_117_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_10_v0, _11_v1)
		_11_v1 = ((func() _dafny.Int {
			if (_117_v86).Contains(!(!((_17_v7).F28()))) {
				return (_117_v86).Get(!(!((_17_v7).F28()))).(_dafny.Int)
			}
			return _11_v1
		})()).Plus(_11_v1)
		(globalState).F0 = true
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))
		_ = _index16
		(_82_v61).ArraySet1(!((_17_v7).F28()) || ((_17_v7).F28()), (_index16).Int())
		var _118_v87 _dafny.Sequence
		_ = _118_v87
		_118_v87 = _dafny.UnicodeSeqOfUtf8Bytes("dicp")
		var _119_v88 _dafny.Sequence
		_ = _119_v88
		_119_v88 = _dafny.SeqOf(_11_v1)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))
		_ = _index17
		var _rhs27 _dafny.Sequence = _13_v3
		_ = _rhs27
		var _rhs28 _dafny.Int = (func() _dafny.Int {
			if _10_v0 {
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_118_v87, _dafny.UnicodeSeqOfUtf8Bytes("dqnhuuat"))).Cardinality())
			}
			return (_dafny.Zero).Minus((_119_v88).Select((Companion_Default___.SafeIndex(_11_v1, _dafny.IntOfUint32((_119_v88).Cardinality()))).Uint32()).(_dafny.Int))
		})()
		_ = _rhs28
		var _rhs29 bool = (_17_v7).F28()
		_ = _rhs29
		var _lhs11 _dafny.Array = _82_v61
		_ = _lhs11
		var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))
		_ = _lhs12
		_13_v3 = _rhs27
		_11_v1 = _rhs28
		(_lhs11).ArraySet1(_rhs29, (_lhs12).Int())
		var _120_v89 _dafny.Array
		_ = _120_v89
		var _nwElement0_6 bool = (_82_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))).Int()).(bool)
		_ = _nwElement0_6
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_6, 0)
		(_nw15).ArraySet1((_81_v60).IsDisjointFrom(_dafny.MultiSetOf(_10_v0, _10_v0, (_17_v7).F28(), _10_v0)), 1)
		(_nw15).ArraySet1(_10_v0, 2)
		(_nw15).ArraySet1(false, 3)
		(_nw15).ArraySet1((_17_v7).F28(), 4)
		(_nw15).ArraySet1((_82_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))).Int()).(bool), 5)
		(_nw15).ArraySet1(true, 6)
		(_nw15).ArraySet1((_17_v7).F28(), 7)
		(_nw15).ArraySet1((_82_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))).Int()).(bool), 8)
		(_nw15).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfInt64(672))).Equals(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(752))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_121_v88 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_122_i7 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_121_v88).Cardinality())
			}
		})(_119_v88))))), 9)
		(_nw15).ArraySet1((_82_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))).Int()).(bool), 10)
		(_nw15).ArraySet1((_17_v7).F28(), 11)
		(_nw15).ArraySet1((_17_v7).F28(), 12)
		_120_v89 = _nw15
		_120_v89 = _120_v89
		var _123_v90 _dafny.Sequence
		_ = _123_v90
		_123_v90 = _dafny.SeqOf(_17_v7, _17_v7, _17_v7)
		var _124_v91 _dafny.Map
		_ = _124_v91
		_124_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, (_82_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_82_v61), 0))).Int()).(bool))
		_17_v7 = (_123_v90).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm3(_dafny.IntOfInt64(502), false, _11_v1, _124_v91, globalState), _dafny.IntOfUint32((_123_v90).Cardinality()))).Uint32()).(*C0)
	}
	var _125_v92 D0
	_ = _125_v92
	_125_v92 = Companion_D0_.Create_DC1_((func() bool {
		if !((_17_v7).F28()) {
			return (_17_v7).F28()
		}
		return _10_v0
	})(), _11_v1)
	var _source0 D0 = _125_v92
	_ = _source0
	if _source0.Is_DC1() {
		var _126___mcc_h0 bool = _source0.Get_().(D0_DC1).Cf1
		_ = _126___mcc_h0
		var _127___mcc_h1 _dafny.Int = _source0.Get_().(D0_DC1).Cf2
		_ = _127___mcc_h1
		var _128_cf2 _dafny.Int = _127___mcc_h1
		_ = _128_cf2
		var _129_cf1 bool = _126___mcc_h0
		_ = _129_cf1
		_11_v1 = _11_v1
		var _130_v93 _dafny.Map
		_ = _130_v93
		_130_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_cf2, !(_129_cf1))
		_128_cf2 = (_dafny.Zero).Minus(Companion_Default___.Fm3((_128_cf2).Minus(_128_cf2), (_17_v7).F28(), _11_v1, _130_v93, globalState))
		var _pat_let_tv8 = _11_v1
		_ = _pat_let_tv8
		var _131_v94 _dafny.Map
		_ = _131_v94
		_131_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let10_0 D0) D0 {
			return func(_132_dt__update__tmp_h6 D0) D0 {
				return func(_pat_let11_0 _dafny.Int) D0 {
					return func(_133_dt__update_hcf2_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_((_132_dt__update__tmp_h6).Dtor_cf1(), _133_dt__update_hcf2_h0)
					}(_pat_let11_0)
				}(_pat_let_tv8)
			}(_pat_let10_0)
		}(_125_v92), _130_v93)
		var _134_v95 _dafny.Sequence
		_ = _134_v95
		_134_v95 = _dafny.SeqOf(_14_v4, _14_v4)
		var _135_v96 _dafny.Sequence
		_ = _135_v96
		_135_v96 = _dafny.SeqOf(_11_v1, _128_cf2, _11_v1, _11_v1)
		var _rhs30 _dafny.Map = (_131_v94).Merge((_131_v94).Merge(_131_v94))
		_ = _rhs30
		var _rhs31 bool = !(((_134_v95).Select((Companion_Default___.SafeIndex(_128_cf2, _dafny.IntOfUint32((_134_v95).Cardinality()))).Uint32()).(_dafny.Set)).Equals(_14_v4)) || (!((_17_v7).F28()))
		_ = _rhs31
		var _rhs32 _dafny.Int = Companion_Default___.SafeDivisionInt(_128_cf2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_135_v96).Cardinality())))
		_ = _rhs32
		var _rhs33 _dafny.Int = _128_cf2
		_ = _rhs33
		var _lhs13 *GlobalState = globalState
		_ = _lhs13
		_131_v94 = _rhs30
		_lhs13.F0 = _rhs31
		_128_cf2 = _rhs32
		_11_v1 = _rhs33
		if (_11_v1).Cmp(_11_v1) != 0 {
			var _136_v97 *C0
			_ = _136_v97
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__((_17_v7).F28(), (_17_v7).F29())
			_136_v97 = _nw16
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index18
			((_17_v7).F29()).ArraySet1((_dafny.IntOfInt64(507)).Times(_128_cf2), (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index19
			((_17_v7).F29()).ArraySet1(((_12_v2).Merge(_12_v2)).Cardinality(), (_index19).Int())
			var _137_v98 _dafny.Sequence
			_ = _137_v98
			_137_v98 = _dafny.SeqOf(_82_v61)
			var _138_v99 _dafny.MultiSet
			_ = _138_v99
			_138_v99 = _dafny.MultiSetOf(_11_v1)
			var _139_v100 _dafny.Sequence
			_ = _139_v100
			_139_v100 = _dafny.UnicodeSeqOfUtf8Bytes("ckgrfoe")
			var _140_v101 _dafny.Map
			_ = _140_v101
			_140_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _139_v100)
			var _141_v102 D0
			_ = _141_v102
			_141_v102 = Companion_D0_.Create_DC0_(_138_v99)
			var _142_v103 _dafny.Map
			_ = _142_v103
			_142_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int), _138_v99)
			var _143_v104 _dafny.Array
			_ = _143_v104
			var _nwElement0_7 _dafny.MultiSet = (_138_v99).Union(_138_v99)
			_ = _nwElement0_7
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(18))
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_7, 0)
			(_nw17).ArraySet1(((_138_v99).Update(_dafny.IntOfInt64(187), Companion_Default___.Abs(((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int)))).Difference(_138_v99), 1)
			(_nw17).ArraySet1(_138_v99, 2)
			(_nw17).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfUint32((_139_v100).Cardinality()), _128_cf2, _11_v1, _dafny.IntOfUint32((_139_v100).Cardinality()), _11_v1)).Intersection(_138_v99), 3)
			(_nw17).ArraySet1(_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_140_v101).Cardinality())))), 4)
			(_nw17).ArraySet1(_138_v99, 5)
			(_nw17).ArraySet1((_dafny.MultiSetFromSeq(_135_v96)).Update(_11_v1, Companion_Default___.Abs(_11_v1)), 6)
			(_nw17).ArraySet1(_138_v99, 7)
			(_nw17).ArraySet1(_138_v99, 8)
			(_nw17).ArraySet1(_dafny.MultiSetFromSeq(_135_v96), 9)
			(_nw17).ArraySet1((_138_v99).Intersection(_138_v99), 10)
			(_nw17).ArraySet1((_141_v102).Dtor_cf0(), 11)
			(_nw17).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_11_v1)), 12)
			(_nw17).ArraySet1(_138_v99, 13)
			(_nw17).ArraySet1(_138_v99, 14)
			(_nw17).ArraySet1(_138_v99, 15)
			(_nw17).ArraySet1(((func() _dafny.MultiSet {
				if (_142_v103).Contains(_dafny.IntOfInt64(705)) {
					return (_142_v103).Get(_dafny.IntOfInt64(705)).(_dafny.MultiSet)
				}
				return _138_v99
			})()).Update(_128_cf2, Companion_Default___.Abs(_11_v1)), 16)
			(_nw17).ArraySet1(_dafny.MultiSetOf(_128_cf2, _dafny.IntOfInt64(-663)), 17)
			_143_v104 = _nw17
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_143_v104), 0))
			_ = _index20
			(_143_v104).ArraySet1(_dafny.MultiSetFromSeq(_135_v96), (_index20).Int())
			var _144_v105 _dafny.CodePoint
			_ = _144_v105
			_144_v105 = _dafny.CodePoint('b')
			var _145_v106 _dafny.Map
			_ = _145_v106
			_145_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _144_v105)
			var _146_v107 _dafny.MultiSet
			_ = _146_v107
			_146_v107 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_cf1, _144_v105), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('k')), (_145_v106).Update((_17_v7).F28(), _144_v105), _145_v106, (func() _dafny.Map {
				if (_136_v97).F28() {
					return _145_v106
				}
				return _145_v106
			})())
			var _147_v108 _dafny.Map
			_ = _147_v108
			_147_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_139_v100).Cardinality()), _136_v97)
			var _148_v109 D7
			_ = _148_v109
			_148_v109 = Companion_D7_.Create_DC17_((func() *C0 {
				if (_147_v108).Contains(_dafny.IntOfUint32((_13_v3).Cardinality())) {
					return (_147_v108).Get(_dafny.IntOfUint32((_13_v3).Cardinality())).(*C0)
				}
				return _17_v7
			})())
			var _149_v110 _dafny.Sequence
			_ = _149_v110
			_149_v110 = _dafny.SeqOf((_148_v109).Dtor_cf26())
			var _150_v111 _dafny.Sequence
			_ = _150_v111
			_150_v111 = _dafny.SeqOf(_139_v100, _139_v100)
			var _151_v112 _dafny.MultiSet
			_ = _151_v112
			_151_v112 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("f"), _139_v100, _139_v100, _dafny.UnicodeSeqOfUtf8Bytes("a"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(55))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_152_v105 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_153_i8 _dafny.Int) _dafny.CodePoint {
					return _152_v105
				}
			})(_144_v105))))
			var _154_v113 D7
			_ = _154_v113
			_154_v113 = Companion_D7_.Create_DC18_((_17_v7).F28())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index21
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_143_v104), 0))
			_ = _index22
			var _rhs34 _dafny.Sequence = _137_v98
			_ = _rhs34
			var _rhs35 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_149_v110, _dafny.SeqOf(_136_v97, _136_v97, _136_v97)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_149_v110, _dafny.SeqOf(_136_v97, _136_v97, _136_v97))).Cardinality()))).Uint32(), _136_v97)).Cardinality())
			_ = _rhs35
			var _rhs36 _dafny.MultiSet = _138_v99
			_ = _rhs36
			var _rhs37 bool = ((func() _dafny.MultiSet {
				if (_136_v97).F28() {
					return Companion_Default___.Fm11(_dafny.IntOfInt64(-642), _dafny.CodePoint('a'), globalState)
				}
				return _151_v112
			})()).IsProperSubsetOf((_dafny.MultiSetFromSeq(_150_v111)).Difference(_151_v112))
			_ = _rhs37
			var _rhs38 _dafny.MultiSet = Companion_Default___.Fm12(_11_v1, ((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if (_81_v60).Contains((_136_v97).F28()) {
					return (_81_v60).Multiplicity((_136_v97).F28())
				}
				return ((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int)
			})(), _154_v113, globalState)
			_ = _rhs38
			var _lhs14 _dafny.Array = (_17_v7).F29()
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _lhs15
			var _lhs16 _dafny.Array = _143_v104
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_143_v104), 0))
			_ = _lhs17
			var _lhs18 *GlobalState = globalState
			_ = _lhs18
			_137_v98 = _rhs34
			(_lhs14).ArraySet1(_rhs35, (_lhs15).Int())
			(_lhs16).ArraySet1(_rhs36, (_lhs17).Int())
			_lhs18.F0 = _rhs37
			_146_v107 = _rhs38
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_82_v61), 0))
			_ = _index23
			(_82_v61).ArraySet1((_17_v7).F28(), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_82_v61), 0))
			_ = _index24
			(_82_v61).ArraySet1((_17_v7).F28(), (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index25
			((_17_v7).F29()).ArraySet1((_128_cf2).Plus((_11_v1).Times(((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int))), (_index25).Int())
		} else {
			_130_v93 = (_130_v93).Update(_11_v1, (_17_v7).F28())
			_17_v7 = _17_v7
			var _155_v114 _dafny.Array
			_ = _155_v114
			_155_v114 = _82_v61
			var _156_v115 _dafny.Array
			_ = _156_v115
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
			_ = _nw18
			_156_v115 = _nw18
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_156_v115), 0))
			_ = _index26
			(_156_v115).ArraySet1CodePoint(_dafny.CodePoint('k'), (_index26).Int())
			var _157_v116 _dafny.CodePoint
			_ = _157_v116
			_157_v116 = _dafny.CodePoint('p')
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_156_v115), 0))
			_ = _index27
			var _rhs39 _dafny.Array = _82_v61
			_ = _rhs39
			var _rhs40 _dafny.CodePoint = _157_v116
			_ = _rhs40
			var _lhs19 _dafny.Array = _156_v115
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_156_v115), 0))
			_ = _lhs20
			_155_v114 = _rhs39
			(_lhs19).ArraySet1CodePoint(_rhs40, (_lhs20).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index28
			((_17_v7).F29()).ArraySet1(_dafny.IntOfInt64(21), (_index28).Int())
			var _158_v117 _dafny.MultiSet
			_ = _158_v117
			_158_v117 = _dafny.MultiSetOf(_dafny.IntOfInt64(-324))
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index29
			((_17_v7).F29()).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_158_v117).Contains(_11_v1) {
					return (_158_v117).Multiplicity(_11_v1)
				}
				return _11_v1
			})()), (_index29).Int())
			var _159_v118 _dafny.Sequence
			_ = _159_v118
			_159_v118 = _dafny.SeqOf(_82_v61, (func() _dafny.Array {
				if (_17_v7).F28() {
					return _82_v61
				}
				return _82_v61
			})(), _82_v61)
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index30
			var _rhs41 _dafny.Int = _11_v1
			_ = _rhs41
			var _rhs42 _dafny.Array = (_159_v118).Select((Companion_Default___.SafeIndex(((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_159_v118).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs42
			var _rhs43 _dafny.Int = _11_v1
			_ = _rhs43
			var _lhs21 _dafny.Array = (_17_v7).F29()
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _lhs22
			(_lhs21).ArraySet1(_rhs41, (_lhs22).Int())
			_82_v61 = _rhs42
			_128_cf2 = _rhs43
		}
	} else if _source0.Is_DC0() {
		var _160___mcc_h2 _dafny.MultiSet = _source0.Get_().(D0_DC0).Cf0
		_ = _160___mcc_h2
		var _161_cf0 _dafny.MultiSet = _160___mcc_h2
		_ = _161_cf0
		_11_v1 = Companion_Default___.Fm3((_11_v1).Plus(_11_v1), (_17_v7).F28(), _11_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, (_17_v7).F28()), globalState)
		var _162_v119 D6
		_ = _162_v119
		_162_v119 = Companion_D6_.Create_DC15_((_17_v7).F28(), (_17_v7).F28(), _11_v1)
		_11_v1 = (_162_v119).Dtor_cf23()
		_11_v1 = Companion_Default___.SafeDivisionInt(_11_v1, _11_v1)
		var _163_v120 _dafny.Sequence
		_ = _163_v120
		_163_v120 = _dafny.UnicodeSeqOfUtf8Bytes("sudyiqf")
		(globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_163_v120, _dafny.Companion_Sequence_.Concatenate(_163_v120, _163_v120))
	} else {
		var _164___mcc_h3 D0 = _source0.Get_().(D0_DC2).Cf3
		_ = _164___mcc_h3
		var _165_cf3 D0 = _164___mcc_h3
		_ = _165_cf3
		var _166_v121 D1
		_ = _166_v121
		_166_v121 = Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_((_17_v7).F29()))
		var _167_v122 D1
		_ = _167_v122
		_167_v122 = Companion_D1_.Create_DC5_(_166_v121)
		var _168_v123 D1
		_ = _168_v123
		_168_v123 = Companion_D1_.Create_DC5_(_167_v122)
		var _169_v124 _dafny.Sequence
		_ = _169_v124
		_169_v124 = _dafny.SeqOf(_168_v123)
		var _170_v125 _dafny.Sequence
		_ = _170_v125
		_170_v125 = _dafny.UnicodeSeqOfUtf8Bytes("yxedcgbw")
		var _171_v126 D6
		_ = _171_v126
		_171_v126 = Companion_D6_.Create_DC16_(_169_v124, _170_v125)
		var _source1 D6 = _171_v126
		_ = _source1
		if _source1.Is_DC14() {
			var _172___mcc_h4 _dafny.Sequence = _source1.Get_().(D6_DC14).Cf20
			_ = _172___mcc_h4
			var _173_cf20 _dafny.Sequence = _172___mcc_h4
			_ = _173_cf20
			var _174_v127 _dafny.Map
			_ = _174_v127
			_174_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v7).F28(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kmsp")).Cardinality()))
			_11_v1 = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v7).F28(), _11_v1)).Merge(_174_v127)).Merge(_174_v127)).Cardinality()
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index31
			((_17_v7).F29()).ArraySet1(_11_v1, (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index32
			((_17_v7).F29()).ArraySet1(_11_v1, (_index32).Int())
			_11_v1 = (((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int)).Minus(((_17_v7).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen(((_17_v7).F29()), 0))).Int()).(_dafny.Int))
			var _175_v128 _dafny.Sequence
			_ = _175_v128
			_175_v128 = _dafny.SeqOf(_170_v125)
			(globalState).F2 = (_175_v128).Select((Companion_Default___.SafeIndex(_11_v1, _dafny.IntOfUint32((_175_v128).Cardinality()))).Uint32()).(_dafny.Sequence)
		} else if _source1.Is_DC15() {
			var _176___mcc_h5 bool = _source1.Get_().(D6_DC15).Cf21
			_ = _176___mcc_h5
			var _177___mcc_h6 bool = _source1.Get_().(D6_DC15).Cf22
			_ = _177___mcc_h6
			var _178___mcc_h7 _dafny.Int = _source1.Get_().(D6_DC15).Cf23
			_ = _178___mcc_h7
			var _179_cf23 _dafny.Int = _178___mcc_h7
			_ = _179_cf23
			var _180_cf22 bool = _177___mcc_h6
			_ = _180_cf22
			var _181_cf21 bool = _176___mcc_h5
			_ = _181_cf21
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index33
			((_17_v7).F29()).ArraySet1(_dafny.IntOfInt64(-981), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen(((_17_v7).F29()), 0))
			_ = _index34
			((_17_v7).F29()).ArraySet1(_11_v1, (_index34).Int())
			(globalState).F2 = _170_v125
			(globalState).F0 = true
			(globalState).F0 = !((_17_v7).F28())
		} else if _source1.Is_DC16() {
			var _182___mcc_h8 _dafny.Sequence = _source1.Get_().(D6_DC16).Cf24
			_ = _182___mcc_h8
			var _183___mcc_h9 _dafny.Sequence = _source1.Get_().(D6_DC16).Cf25
			_ = _183___mcc_h9
			var _184_cf25 _dafny.Sequence = _183___mcc_h9
			_ = _184_cf25
			var _185_cf24 _dafny.Sequence = _182___mcc_h8
			_ = _185_cf24
			var _186_v129 _dafny.Set
			_ = _186_v129
			_186_v129 = _dafny.SetOf(_17_v7)
			var _187_v130 _dafny.Map
			_ = _187_v130
			_187_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v129, (_17_v7).F28())
			var _188_v131 _dafny.Array
			_ = _188_v131
			var _nwElement0_8 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v129, (_17_v7).F28())
			_ = _nwElement0_8
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(13))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_8, 0)
			(_nw19).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_17_v7), !((_17_v7).F28()))).Merge(_187_v130), 1)
			(_nw19).ArraySet1((_187_v130).Merge(_187_v130), 2)
			(_nw19).ArraySet1(_187_v130, 3)
			(_nw19).ArraySet1(_187_v130, 4)
			(_nw19).ArraySet1(_187_v130, 5)
			(_nw19).ArraySet1((_187_v130).Merge(_187_v130), 6)
			(_nw19).ArraySet1(_187_v130, 7)
			(_nw19).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v129, (_17_v7).F28())).Update(_dafny.SetOf(_17_v7), (_17_v7).F28()), 8)
			(_nw19).ArraySet1((_187_v130).Update(_186_v129, (_17_v7).F28()), 9)
			(_nw19).ArraySet1(_187_v130, 10)
			(_nw19).ArraySet1(_187_v130, 11)
			(_nw19).ArraySet1(_187_v130, 12)
			_188_v131 = _nw19
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_188_v131), 0))
			_ = _index35
			(_188_v131).ArraySet1((_187_v130).Merge(_187_v130), (_index35).Int())
			var _189_v132 D7
			_ = _189_v132
			_189_v132 = Companion_D7_.Create_DC18_((_17_v7).F28())
			var _190_v133 D7
			_ = _190_v133
			_190_v133 = Companion_D7_.Create_DC19_(_189_v132)
			var _191_v134 D7
			_ = _191_v134
			_191_v134 = Companion_D7_.Create_DC19_(_189_v132)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_188_v131), 0))
			_ = _index36
			var _rhs44 _dafny.Map = _187_v130
			_ = _rhs44
			var _rhs45 *C0 = (func() *C0 {
				if _10_v0 {
					return _17_v7
				}
				return _17_v7
			})()
			_ = _rhs45
			var _rhs46 _dafny.Sequence = Companion_Default___.Fm9(globalState)
			_ = _rhs46
			var _rhs47 D7 = _191_v134
			_ = _rhs47
			var _lhs23 _dafny.Array = _188_v131
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_188_v131), 0))
			_ = _lhs24
			var _lhs25 *GlobalState = globalState
			_ = _lhs25
			(_lhs23).ArraySet1(_rhs44, (_lhs24).Int())
			_17_v7 = _rhs45
			_lhs25.F2 = _rhs46
			_191_v134 = _rhs47
			_14_v4 = _14_v4
			_11_v1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_184_cf25).Cardinality()), _11_v1)
			var _192_v135 _dafny.Map
			_ = _192_v135
			_192_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_11_v1, _11_v1)), (_17_v7).F28())
			var _193_v136 _dafny.Sequence
			_ = _193_v136
			_193_v136 = _dafny.SeqOf(_11_v1)
			_192_v135 = (_192_v135).Update((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(294)), _193_v136, _193_v136)).Cardinality(), (_17_v7).F28())
		} else {
			var _194___mcc_h10 _dafny.CodePoint = _source1.Get_().(D6_DC13).Cf19
			_ = _194___mcc_h10
			var _195_cf19 _dafny.CodePoint = _194___mcc_h10
			_ = _195_cf19
			(globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_170_v125, (_171_v126).Dtor_cf25())
			var _196_v137 *C0
			_ = _196_v137
			var _nw20 *C0 = New_C0_()
			_ = _nw20
			_nw20.Ctor__((_17_v7).F28(), _16_v6)
			_196_v137 = _nw20
			var _197_v138 D1
			_ = _197_v138
			_197_v138 = Companion_D1_.Create_DC4_((_11_v1).Cmp(_11_v1) <= 0)
			var _rhs48 _dafny.Int = Companion_Default___.SafeModuloInt(_11_v1, _11_v1)
			_ = _rhs48
			var _rhs49 D1 = _197_v138
			_ = _rhs49
			var _rhs50 bool = (_17_v7).Fm2((_17_v7).F28(), globalState)
			_ = _rhs50
			var _rhs51 bool = (!((_17_v7).F28())) || ((_196_v137).F28())
			_ = _rhs51
			var _rhs52 _dafny.Array = (_196_v137).F29()
			_ = _rhs52
			var _lhs26 *GlobalState = globalState
			_ = _lhs26
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			_11_v1 = _rhs48
			_197_v138 = _rhs49
			_lhs26.F0 = _rhs50
			_lhs27.F0 = _rhs51
			_16_v6 = _rhs52
			_11_v1 = (_11_v1).Minus((_dafny.Zero).Minus(_11_v1))
		}
		var _198_v139 _dafny.CodePoint
		_ = _198_v139
		_198_v139 = _dafny.CodePoint('g')
		var _199_v140 _dafny.Array
		_ = _199_v140
		var _nwElement0_9 _dafny.CodePoint = _198_v139
		_ = _nwElement0_9
		var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(2))
		_ = _nw21
		(_nw21).ArraySet1CodePoint(_nwElement0_9, 0)
		(_nw21).ArraySet1CodePoint(_198_v139, 1)
		_199_v140 = _nw21
		var _200_v141 _dafny.Map
		_ = _200_v141
		_200_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v139, _199_v140)
		var _201_v142 _dafny.Sequence
		_ = _201_v142
		_201_v142 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v139, _199_v140))
		var _202_v143 _dafny.Sequence
		_ = _202_v143
		_202_v143 = _dafny.SeqOf(_17_v7)
		var _203_v144 _dafny.MultiSet
		_ = _203_v144
		_203_v144 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(597))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_204_i9 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()), _dafny.IntOfUint32((_202_v143).Cardinality()))
		var _205_v145 D8
		_ = _205_v145
		_205_v145 = Companion_D8_.Create_DC20_(_15_v5)
		var _206_v146 _dafny.Array
		_ = _206_v146
		var _nwElement0_10 _dafny.Map = _200_v141
		_ = _nwElement0_10
		var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(24))
		_ = _nw22
		(_nw22).ArraySet1(_nwElement0_10, 0)
		(_nw22).ArraySet1(_200_v141, 1)
		(_nw22).ArraySet1((_201_v142).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_203_v144).Contains(_11_v1) {
				return (_203_v144).Multiplicity(_11_v1)
			}
			return ((_205_v145).Dtor_cf29()).Cardinality()
		})(), _dafny.IntOfUint32((_201_v142).Cardinality()))).Uint32()).(_dafny.Map), 2)
		(_nw22).ArraySet1(_200_v141, 3)
		(_nw22).ArraySet1(_200_v141, 4)
		(_nw22).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v139, _199_v140), 5)
		(_nw22).ArraySet1(_200_v141, 6)
		(_nw22).ArraySet1(_200_v141, 7)
		(_nw22).ArraySet1((_200_v141).Merge(_200_v141), 8)
		(_nw22).ArraySet1((_201_v142).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-425), _dafny.IntOfUint32((_201_v142).Cardinality()))).Uint32()).(_dafny.Map), 9)
		(_nw22).ArraySet1(_200_v141, 10)
		(_nw22).ArraySet1(_200_v141, 11)
		(_nw22).ArraySet1((_201_v142).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.IntOfUint32((_201_v142).Cardinality()))).Uint32()).(_dafny.Map), 12)
		(_nw22).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v139, _199_v140)).Merge(_200_v141), 13)
		(_nw22).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v139, _199_v140), 14)
		(_nw22).ArraySet1(_200_v141, 15)
		(_nw22).ArraySet1(_200_v141, 16)
		(_nw22).ArraySet1(_200_v141, 17)
		(_nw22).ArraySet1((_200_v141).Update(_198_v139, _199_v140), 18)
		(_nw22).ArraySet1(_200_v141, 19)
		(_nw22).ArraySet1(_200_v141, 20)
		(_nw22).ArraySet1(_200_v141, 21)
		(_nw22).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v139, _199_v140)).Merge((_200_v141).Update(_198_v139, _199_v140)), 22)
		(_nw22).ArraySet1(((_200_v141).Update(_198_v139, _199_v140)).Merge(_200_v141), 23)
		_206_v146 = _nw22
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_206_v146), 0))
		_ = _index37
		(_206_v146).ArraySet1(_200_v141, (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_206_v146), 0))
		_ = _index38
		(_206_v146).ArraySet1((_200_v141).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v139, _199_v140)), (_index38).Int())
		var _207_v147 _dafny.Map
		_ = _207_v147
		_207_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, (_17_v7).F28())
		_11_v1 = Companion_Default___.SafeModuloInt(_11_v1, (_dafny.Zero).Minus((_dafny.Zero).Minus(((_207_v147).Cardinality()).Minus(_11_v1))))
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen(((_17_v7).F29()), 0))
		_ = _index39
		((_17_v7).F29()).ArraySet1(_dafny.IntOfInt64(346), (_index39).Int())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen(((_17_v7).F29()), 0))
		_ = _index40
		((_17_v7).F29()).ArraySet1((_11_v1).Minus((_125_v92).Dtor_cf2()), (_index40).Int())
	}
	var _208_v148 _dafny.Map
	_ = _208_v148
	_208_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _11_v1)
	var _209_v149 _dafny.Map
	_ = _209_v149
	_209_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_208_v148, _208_v148), (_17_v7).F28())
	var _210_v150 _dafny.CodePoint
	_ = _210_v150
	_210_v150 = _dafny.CodePoint('q')
	var _211_v151 _dafny.Map
	_ = _211_v151
	_211_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v7).F28(), _210_v150)
	_209_v149 = (_209_v149).Update(_dafny.SeqOf((_208_v148).Update(_11_v1, (_211_v151).Cardinality())), (_17_v7).F28())
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _212_v0 bool
	_ = _212_v0
	_212_v0 = false
	var _213_v1 _dafny.Map
	_ = _213_v1
	_213_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v0, _212_v0)
	var _214_v2 _dafny.Set
	_ = _214_v2
	_214_v2 = _dafny.SetOf(_212_v0)
	var _215_v3 _dafny.Sequence
	_ = _215_v3
	_215_v3 = _dafny.UnicodeSeqOfUtf8Bytes("sk")
	var _216_v4 _dafny.Sequence
	_ = _216_v4
	_216_v4 = _dafny.SeqOf(_212_v0, _212_v0, _212_v0, _212_v0)
	var _217_v5 _dafny.Int
	_ = _217_v5
	_217_v5 = _dafny.IntOfInt64(228)
	var _218_v6 _dafny.Sequence
	_ = _218_v6
	_218_v6 = _dafny.SeqOf(_217_v5)
	var _219_v7 _dafny.Set
	_ = _219_v7
	_219_v7 = _dafny.SetOf(_dafny.IntOfUint32((_218_v6).Cardinality()))
	var _220_v8 _dafny.Sequence
	_ = _220_v8
	_220_v8 = _dafny.SeqOf((_dafny.Zero).Minus(_217_v5), (_219_v7).Cardinality())
	var _221_v9 _dafny.CodePoint
	_ = _221_v9
	_221_v9 = _dafny.CodePoint('v')
	var _222_v10 _dafny.Map
	_ = _222_v10
	_222_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_217_v5), _221_v9)
	var _223_v11 _dafny.MultiSet
	_ = _223_v11
	_223_v11 = _dafny.MultiSetOf(_222_v10)
	var _224_v12 _dafny.Array
	_ = _224_v12
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(13)
	_ = _len0_1
	var _nw23 _dafny.Array
	_ = _nw23
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw23 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = (func(_225_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_226_i0 _dafny.Int) _dafny.Int {
				return (_226_i0).Plus(_225_v5)
			}
		})(_217_v5)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw23 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw23).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw23).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_224_v12 = _nw23
	var _227_v13 _dafny.Map
	_ = _227_v13
	_227_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ybllip"), _dafny.MultiSetOf(_217_v5, _217_v5))
	var _228_v15 _dafny.Sequence
	_ = _228_v15
	_228_v15 = _dafny.SeqOf(_218_v6)
	var _229_v16 D0
	_ = _229_v16
	_229_v16 = Companion_D0_.Create_DC1_(_212_v0, (_213_v1).Cardinality())
	var _230_v17 D0
	_ = _230_v17
	_230_v17 = Companion_D0_.Create_DC2_(_229_v16)
	var _231_v18 _dafny.MultiSet
	_ = _231_v18
	_231_v18 = _dafny.MultiSetOf(_230_v17, _230_v17)
	var _232_v19 _dafny.MultiSet
	_ = _232_v19
	_232_v19 = _dafny.MultiSetOf(_231_v18)
	var _233_v20 _dafny.Map
	_ = _233_v20
	_233_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if (_232_v19).Contains(_231_v18) {
			return (_232_v19).Multiplicity(_231_v18)
		}
		return _217_v5
	})(), _212_v0)
	var _234_globalState *GlobalState
	_ = _234_globalState
	var _nw24 *GlobalState = New_GlobalState_()
	_ = _nw24
	_nw24.Ctor__(true, (_dafny.SetOf((func() bool {
		if (_213_v1).Contains(_212_v0) {
			return (_213_v1).Get(_212_v0).(bool)
		}
		return _212_v0
	})())).Difference(_214_v2), _215_v3, _216_v4, _220_v8, _dafny.IntOfInt64(-929), _dafny.IntOfInt64(851), true, _223_v11, _dafny.IntOfInt64(90), _dafny.IntOfInt64(-909), _dafny.IntOfInt64(917), _dafny.IntOfInt64(336), _224_v12, _dafny.IntOfInt64(206), _dafny.IntOfInt64(664), _dafny.IntOfInt64(-432), (_227_v13).Merge(func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-191))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_235_i1 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("incqt")
		})), (Companion_Default___.SafeIndex(_217_v5, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-191))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}(func(_236_i1 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("incqt")
		}))).Cardinality()))).Uint32(), _215_v3)).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _237_v14 _dafny.Sequence
			_237_v14 = interface{}(_compr_7).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-191))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_235_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("incqt")
			})), (Companion_Default___.SafeIndex(_217_v5, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-191))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_236_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("incqt")
			}))).Cardinality()))).Uint32(), _215_v3), _237_v14) {
				_coll7.Add(_237_v14, (Companion_D0_.Create_DC0_(_dafny.MultiSetOf(_217_v5))).Dtor_cf0())
			}
		}
		return _coll7.ToMap()
	}()), true, _228_v15, _dafny.Companion_Sequence_.Update(_215_v3, (Companion_Default___.SafeIndex(_217_v5, _dafny.IntOfUint32((_215_v3).Cardinality()))).Uint32(), _221_v9), _dafny.IntOfInt64(-661), _dafny.IntOfInt64(-89), _dafny.SeqOf(_217_v5), true, _dafny.IntOfInt64(849), false, _233_v20)
	_234_globalState = _nw24
	Companion_Default___.M0(_234_globalState)
	if (_dafny.IntOfUint32((_216_v4).Cardinality())).Cmp((_233_v20).Cardinality()) == 0 {
		_217_v5 = _dafny.IntOfUint32((_215_v3).Cardinality())
		(_234_globalState).F0 = !(!(Companion_Default___.Fm0(_234_globalState)))
		var _238_v21 *C0
		_ = _238_v21
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__(_212_v0, _224_v12)
		_238_v21 = _nw25
		var _239_v22 D0
		_ = _239_v22
		_239_v22 = Companion_D0_.Create_DC1_(_212_v0, Companion_Default___.Fm3(_217_v5, _212_v0, Companion_Default___.Fm3(_217_v5, (_238_v21).F28(), _dafny.IntOfInt64(544), _233_v20, _234_globalState), _233_v20, _234_globalState))
		var _240_v23 D1
		_ = _240_v23
		_240_v23 = Companion_D1_.Create_DC3_(_224_v12)
		var _241_v24 *C0
		_ = _241_v24
		var _nw26 *C0 = New_C0_()
		_ = _nw26
		_nw26.Ctor__((_239_v22).Dtor_cf1(), (_240_v23).Dtor_cf4())
		_241_v24 = _nw26
		_241_v24 = _241_v24
		_221_v9 = (_241_v24).Fm1(Companion_D0_.Create_DC1_((_238_v21).F28(), _217_v5), _217_v5, _217_v5, _dafny.IntOfUint32((_215_v3).Cardinality()), _234_globalState)
	} else {
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_224_v12), 0))
		_ = _index41
		(_224_v12).ArraySet1((_219_v7).Cardinality(), (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_224_v12), 0))
		_ = _index42
		(_224_v12).ArraySet1((_dafny.Zero).Minus(((_213_v1).Merge((_213_v1).Update(_212_v0, _212_v0))).Cardinality()), (_index42).Int())
		var _242_v25 _dafny.Array
		_ = _242_v25
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_2
		var _nw27 _dafny.Array
		_ = _nw27
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw27 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_243_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_244_i2 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_244_i2, _dafny.IntOfUint32((_243_v8).Cardinality()))
				}
			})(_220_v8)
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
		_242_v25 = _nw27
		var _245_v26 *C0
		_ = _245_v26
		var _nw28 *C0 = New_C0_()
		_ = _nw28
		_nw28.Ctor__(_212_v0, _242_v25)
		_245_v26 = _nw28
		(_234_globalState).F0 = (_245_v26).F28()
		_217_v5 = (_224_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_224_v12), 0))).Int()).(_dafny.Int)
		var _246_v27 _dafny.Array
		_ = _246_v27
		var _nw29 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw29
		_246_v27 = _nw29
		var _247_v29 _dafny.MultiSet
		_ = _247_v29
		_247_v29 = _dafny.MultiSetOf(_233_v20, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_224_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_224_v12), 0))).Int()).(_dafny.Int), (_245_v26).F28()), func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-392))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_248_v26 *C0) func(_dafny.Int) _dafny.Int {
				return func(_249_i3 _dafny.Int) _dafny.Int {
					return ((Companion_D2_.Create_DC6_(_dafny.MultiSetOf((_248_v26).F28()))).Dtor_cf7()).Cardinality()
				}
			})(_245_v26)))).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _250_v28 _dafny.Int
				_250_v28 = interface{}(_compr_8).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-392))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}((func(_251_v26 *C0) func(_dafny.Int) _dafny.Int {
					return func(_249_i3 _dafny.Int) _dafny.Int {
						return ((Companion_D2_.Create_DC6_(_dafny.MultiSetOf((_251_v26).F28()))).Dtor_cf7()).Cardinality()
					}
				})(_245_v26))), _250_v28) {
					_coll8.Add(Companion_Default___.SafeDivisionInt(_250_v28, _217_v5), (_245_v26).F28())
				}
			}
			return _coll8.ToMap()
		}(), _233_v20, _233_v20)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_246_v27), 0))
		_ = _index43
		(_246_v27).ArraySet1(((_247_v29).Cardinality()).Cmp((_224_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_224_v12), 0))).Int()).(_dafny.Int)) < 0, (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_246_v27), 0))
		_ = _index44
		(_246_v27).ArraySet1((_245_v26).F28(), (_index44).Int())
	}
	_212_v0 = _212_v0
	_217_v5 = _217_v5
	var _252_v30 _dafny.Sequence
	_ = _252_v30
	_252_v30 = _dafny.SeqOf(_215_v3, _215_v3)
	_252_v30 = _dafny.Companion_Sequence_.Concatenate(_252_v30, _dafny.Companion_Sequence_.Concatenate(_252_v30, _252_v30))
	var _253_v31 _dafny.MultiSet
	_ = _253_v31
	_253_v31 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_217_v5), (_dafny.Zero).Minus(_217_v5)))
	_253_v31 = (Companion_Default___.Fm4(_234_globalState)).Difference(_253_v31)
	var _254_v32 D3
	_ = _254_v32
	_254_v32 = Companion_D3_.Create_DC9_(_233_v20)
	(_234_globalState).F0 = (func() bool {
		if (_217_v5).Cmp(Companion_Default___.Fm3((_dafny.Zero).Minus(_217_v5), _212_v0, _217_v5, (_254_v32).Dtor_cf11(), _234_globalState)) <= 0 {
			return !((_dafny.MultiSetFromSeq(_dafny.SeqOf(_217_v5))).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v0, true)).Cardinality()))
		}
		return !(_212_v0) || (!(_212_v0))
	})()
	var _255_v33 _dafny.Array
	_ = _255_v33
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(20)
	_ = _len0_3
	var _nw30 _dafny.Array
	_ = _nw30
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw30 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Sequence = func(_256_i4 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("skjjyean")
		}
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw30 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw30).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw30).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_255_v33 = _nw30
	var _257_v34 _dafny.Map
	_ = _257_v34
	_257_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v33, _217_v5)
	_257_v34 = (_257_v34).Update(_255_v33, _217_v5)
	_217_v5 = (_217_v5).Plus(_dafny.IntOfUint32((_215_v3).Cardinality()))
	(_234_globalState).F0 = _212_v0
	(_234_globalState).F0 = !((func() bool {
		if (_213_v1).Contains(_212_v0) {
			return (_213_v1).Get(_212_v0).(bool)
		}
		return (!(_212_v0)) == (_212_v0)
	})())
	var _258_v35 _dafny.Map
	_ = _258_v35
	_258_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v0, _224_v12)
	var _259_v36 *C0
	_ = _259_v36
	var _nw31 *C0 = New_C0_()
	_ = _nw31
	_nw31.Ctor__(_212_v0, _224_v12)
	_259_v36 = _nw31
	var _260_v37 _dafny.Set
	_ = _260_v37
	_260_v37 = _dafny.SetOf(_259_v36, _259_v36, _259_v36)
	var _261_v38 _dafny.Map
	_ = _261_v38
	_261_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_v35, (_260_v37).IsSubsetOf(_260_v37))
	var _262_v39 _dafny.Sequence
	_ = _262_v39
	_262_v39 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_259_v36).F28(), _224_v12))
	var _263_v40 _dafny.Sequence
	_ = _263_v40
	_263_v40 = _dafny.SeqOf((_262_v39).Select((Companion_Default___.SafeIndex(_217_v5, _dafny.IntOfUint32((_262_v39).Cardinality()))).Uint32()).(_dafny.Map))
	_261_v38 = (_261_v38).Update(((_263_v40).Select((Companion_Default___.SafeIndex(_217_v5, _dafny.IntOfUint32((_263_v40).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_258_v35), true)
	var _264_v41 _dafny.Map
	_ = _264_v41
	_264_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v0, (_259_v36).F28())).Update(true, _212_v0), _259_v36)
	var _hi0 _dafny.Int = (_264_v41).Cardinality()
	_ = _hi0
	for _265_i5 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_252_v30, _252_v30)).Cardinality()); _265_i5.Cmp(_hi0) < 0; _265_i5 = _265_i5.Plus(_dafny.One) {
		(_234_globalState).F0 = !(!(!(_219_v7).Equals(_219_v7)))
		(_234_globalState).F0 = (_259_v36).F28()
		var _266_v42 _dafny.MultiSet
		_ = _266_v42
		_266_v42 = _dafny.MultiSetOf(_259_v36)
		var _267_v43 *C0
		_ = _267_v43
		var _nw32 *C0 = New_C0_()
		_ = _nw32
		_nw32.Ctor__((_dafny.MultiSetOf(_259_v36)).IsSubsetOf(_266_v42), _224_v12)
		_267_v43 = _nw32
		var _268_v44 _dafny.Map
		_ = _268_v44
		_268_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_i5, _dafny.SeqOf((_259_v36).F28(), _212_v0, !((_259_v36).F28())))
		_216_v4 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_268_v44).Contains(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v43, _217_v5)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_259_v36, _dafny.IntOfInt64(-199))).Update(_267_v43, _265_i5))).Cardinality()) {
				return (_268_v44).Get(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v43, _217_v5)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_259_v36, _dafny.IntOfInt64(-199))).Update(_267_v43, _265_i5))).Cardinality()).(_dafny.Sequence)
			}
			return _dafny.SeqOf(!((_259_v36).F28()))
		})(), (Companion_Default___.SafeIndex((_265_i5).Plus(_217_v5), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_268_v44).Contains(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v43, _217_v5)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_259_v36, _dafny.IntOfInt64(-199))).Update(_267_v43, _265_i5))).Cardinality()) {
				return (_268_v44).Get(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v43, _217_v5)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_259_v36, _dafny.IntOfInt64(-199))).Update(_267_v43, _265_i5))).Cardinality()).(_dafny.Sequence)
			}
			return _dafny.SeqOf(!((_259_v36).F28()))
		})()).Cardinality()))).Uint32(), (_259_v36).F28())
	}
	_254_v32 = _254_v32
	var _269_v45 D3
	_ = _269_v45
	_269_v45 = Companion_D3_.Create_DC10_((_259_v36).F28(), _dafny.IntOfInt64(281), _212_v0, _217_v5, (_259_v36).F29())
	_217_v5 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iy")).Cardinality()), Companion_Default___.Fm3(_217_v5, (_259_v36).F28(), _dafny.IntOfUint32((_dafny.SeqOf(!((_269_v45).Dtor_cf14()))).Cardinality()), Companion_Default___.Fm5(_221_v9, _234_globalState), _234_globalState))
	for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_255_v33), 0))); ; {
		_guard_loop_0, _ok9 := _iter9()
		if !_ok9 {
			break
		}
		var _270_i6 _dafny.Int
		_270_i6 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_270_i6).Sign() != -1) && ((_270_i6).Cmp(_dafny.ArrayLen((_255_v33), 0)) < 0)) {
			(_255_v33).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lnqf"), (_270_i6).Int())
		}
	}
	_dafny.Print(_212_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_213_v1).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v2).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_215_v3.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_216_v4, _dafny.SeqOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_217_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_218_v6, _dafny.SeqOf(_dafny.IntOfInt64(228))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v7).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_220_v8, _dafny.SeqOf(_dafny.IntOfInt64(-228), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_221_v9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-228), _dafny.CodePoint('v'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v11).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-228), _dafny.CodePoint('v')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v12).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_227_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ybllip"), _dafny.MultiSetOf(_dafny.IntOfInt64(228), _dafny.IntOfInt64(228)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_228_v15, _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(228)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v16).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v16).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_v17).Dtor_cf3()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_v17).Dtor_cf3()).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_231_v18).Equals(_dafny.MultiSetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false, _dafny.One)), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false, _dafny.One)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_232_v19).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false, _dafny.One)), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false, _dafny.One))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_233_v20).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_234_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F1()).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_234_globalState.F2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_234_globalState).F3(), _dafny.SeqOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_234_globalState.F4, _dafny.SeqOf(_dafny.IntOfInt64(-228), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F8()).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-228), _dafny.CodePoint('v')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState.F17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ybllip"), _dafny.MultiSetOf(_dafny.IntOfInt64(228), _dafny.IntOfInt64(228))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("incqt"), _dafny.MultiSetOf(_dafny.IntOfInt64(228))).UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("sk"), _dafny.MultiSetOf(_dafny.IntOfInt64(228)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_234_globalState).F19(), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(228)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F20().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_234_globalState).F23(), _dafny.SeqOf(_dafny.IntOfInt64(228))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_globalState).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_globalState).F27()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_252_v30, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("sk"), _dafny.UnicodeSeqOfUtf8Bytes("sk"), _dafny.UnicodeSeqOfUtf8Bytes("sk"), _dafny.UnicodeSeqOfUtf8Bytes("sk"), _dafny.UnicodeSeqOfUtf8Bytes("sk"), _dafny.UnicodeSeqOfUtf8Bytes("sk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v31).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(462), _dafny.IntOfInt64(-1), _dafny.IntOfInt64(-650))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_254_v32).Dtor_cf11()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v33).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v34).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v35).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v36).F28())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v36).F29()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v37).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_261_v38).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_262_v39).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_263_v40).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_264_v41).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v45).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v45).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v45).Dtor_cf14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v45).Dtor_cf15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v45).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
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
	Cf2 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.MultiSet
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.MultiSet) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf3 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf3 D0) D0 {
	return D0{D0_DC2{Cf3}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf0() _dafny.MultiSet {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf5 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 bool) D1 {
	return D1{D1_DC4{Cf5}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf4 _dafny.Array
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.Array) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC5 struct {
	Cf6 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf6 D1) D1 {
	return D1{D1_DC5{Cf6}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false)
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf6() D1 {
	return _this.Get_().(D1_DC5).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf5 == data2.Cf5
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4 == data2.Cf4
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
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

type D2_DC7 struct {
	Cf8 _dafny.Int
	Cf9 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf8 _dafny.Int, Cf9 _dafny.Int) D2 {
	return D2{D2_DC7{Cf8, Cf9}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf7 _dafny.MultiSet
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf7 _dafny.MultiSet) D2 {
	return D2{D2_DC6{Cf7}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC8 struct {
	Cf10 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf10 D2) D2 {
	return D2{D2_DC8{Cf10}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) Dtor_cf7() _dafny.MultiSet {
	return _this.Get_().(D2_DC6).Cf7
}

func (_this D2) Dtor_cf10() D2 {
	return _this.Get_().(D2_DC8).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf7.Equals(data2.Cf7)
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
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

type D3_DC10 struct {
	Cf12 bool
	Cf13 _dafny.Int
	Cf14 bool
	Cf15 _dafny.Int
	Cf16 _dafny.Array
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf12 bool, Cf13 _dafny.Int, Cf14 bool, Cf15 _dafny.Int, Cf16 _dafny.Array) D3 {
	return D3{D3_DC10{Cf12, Cf13, Cf14, Cf15, Cf16}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf11 _dafny.Map
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf11 _dafny.Map) D3 {
	return D3{D3_DC9{Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, _dafny.Zero, false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf14() bool {
	return _this.Get_().(D3_DC10).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
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
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16
		}
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
	Cf17 _dafny.Array
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 _dafny.Array) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D4) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf17) + ")"
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
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf17 == data2.Cf17
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

type D5_DC12 struct {
	Cf18 _dafny.Sequence
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf18 _dafny.Sequence) D5 {
	return D5{D5_DC12{Cf18}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D5) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + data.Cf18.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
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
	Cf20 _dafny.Sequence
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf20 _dafny.Sequence) D6 {
	return D6{D6_DC14{Cf20}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC15 struct {
	Cf21 bool
	Cf22 bool
	Cf23 _dafny.Int
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf21 bool, Cf22 bool, Cf23 _dafny.Int) D6 {
	return D6{D6_DC15{Cf21, Cf22, Cf23}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC16 struct {
	Cf24 _dafny.Sequence
	Cf25 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf24 _dafny.Sequence, Cf25 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf24, Cf25}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC13 struct {
	Cf19 _dafny.CodePoint
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf19 _dafny.CodePoint) D6 {
	return D6{D6_DC13{Cf19}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(_dafny.EmptySeq)
}

func (_this D6) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D6_DC14).Cf20
}

func (_this D6) Dtor_cf21() bool {
	return _this.Get_().(D6_DC15).Cf21
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC15).Cf22
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf23
}

func (_this D6) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf25
}

func (_this D6) Dtor_cf19() _dafny.CodePoint {
	return _this.Get_().(D6_DC13).Cf19
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf24) + ", " + data.Cf25.VerbatimString(true) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf20.Equals(data2.Cf20)
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf24.Equals(data2.Cf24) && data1.Cf25.Equals(data2.Cf25)
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf19 == data2.Cf19
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

type D7_DC18 struct {
	Cf27 bool
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf27 bool) D7 {
	return D7{D7_DC18{Cf27}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC17 struct {
	Cf26 *C0
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf26 *C0) D7 {
	return D7{D7_DC17{Cf26}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC19 struct {
	Cf28 D7
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf28 D7) D7 {
	return D7{D7_DC19{Cf28}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_(false)
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC18).Cf27
}

func (_this D7) Dtor_cf26() *C0 {
	return _this.Get_().(D7_DC17).Cf26
}

func (_this D7) Dtor_cf28() D7 {
	return _this.Get_().(D7_DC19).Cf28
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf27 == data2.Cf27
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf26 == data2.Cf26
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf28.Equals(data2.Cf28)
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

type D8_DC21 struct {
	Cf30 _dafny.Int
	Cf31 bool
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf30 _dafny.Int, Cf31 bool) D8 {
	return D8{D8_DC21{Cf30, Cf31}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC20 struct {
	Cf29 _dafny.Set
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf29 _dafny.Set) D8 {
	return D8{D8_DC20{Cf29}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(_dafny.Zero, false)
}

func (_this D8) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf30
}

func (_this D8) Dtor_cf31() bool {
	return _this.Get_().(D8_DC21).Cf31
}

func (_this D8) Dtor_cf29() _dafny.Set {
	return _this.Get_().(D8_DC20).Cf29
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31 == data2.Cf31
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf29.Equals(data2.Cf29)
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
	F0   bool
	F2   _dafny.Sequence
	F4   _dafny.Sequence
	F17  _dafny.Map
	_f1  _dafny.Set
	_f3  _dafny.Sequence
	_f5  _dafny.Int
	_f6  _dafny.Int
	_f7  bool
	_f8  _dafny.MultiSet
	_f9  _dafny.Int
	_f10 _dafny.Int
	_f11 _dafny.Int
	_f12 _dafny.Int
	_f13 _dafny.Array
	_f14 _dafny.Int
	_f15 _dafny.Int
	_f16 _dafny.Int
	_f18 bool
	_f19 _dafny.Sequence
	_f20 _dafny.Sequence
	_f21 _dafny.Int
	_f22 _dafny.Int
	_f23 _dafny.Sequence
	_f24 bool
	_f25 _dafny.Int
	_f26 bool
	_f27 _dafny.Map
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F2 = _dafny.EmptySeq
	_this.F4 = _dafny.EmptySeq
	_this.F17 = _dafny.EmptyMap
	_this._f1 = _dafny.EmptySet
	_this._f3 = _dafny.EmptySeq
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f7 = false
	_this._f8 = _dafny.EmptyMultiSet
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.Zero
	_this._f11 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.Zero
	_this._f18 = false
	_this._f19 = _dafny.EmptySeq
	_this._f20 = _dafny.EmptySeq
	_this._f21 = _dafny.Zero
	_this._f22 = _dafny.Zero
	_this._f23 = _dafny.EmptySeq
	_this._f24 = false
	_this._f25 = _dafny.Zero
	_this._f26 = false
	_this._f27 = _dafny.EmptyMap
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Set, f2 _dafny.Sequence, f3 _dafny.Sequence, f4 _dafny.Sequence, f5 _dafny.Int, f6 _dafny.Int, f7 bool, f8 _dafny.MultiSet, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Int, f12 _dafny.Int, f13 _dafny.Array, f14 _dafny.Int, f15 _dafny.Int, f16 _dafny.Int, f17 _dafny.Map, f18 bool, f19 _dafny.Sequence, f20 _dafny.Sequence, f21 _dafny.Int, f22 _dafny.Int, f23 _dafny.Sequence, f24 bool, f25 _dafny.Int, f26 bool, f27 _dafny.Map) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
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
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this).F17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this)._f27 = f27
	}
}
func (_this *GlobalState) F1() _dafny.Set {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Sequence {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.MultiSet {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Int {
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
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() _dafny.Sequence {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() _dafny.Sequence {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() _dafny.Int {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F22() _dafny.Int {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() _dafny.Sequence {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() bool {
	{
		return _this._f24
	}
}
func (_this *GlobalState) F25() _dafny.Int {
	{
		return _this._f25
	}
}
func (_this *GlobalState) F26() bool {
	{
		return _this._f26
	}
}
func (_this *GlobalState) F27() _dafny.Map {
	{
		return _this._f27
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f28 bool
	_f29 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f28 = false
	_this._f29 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) Ctor__(f28 bool, f29 _dafny.Array) {
	{
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C0) Fm1(p0 D0, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		if !(true) {
			return _dafny.CodePoint('u')
		} else {
			return _dafny.CodePoint('o')
		}
	}
}
func (_this *C0) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F28()
	}
}
func (_this *C0) F28() bool {
	{
		return _this._f28
	}
}
func (_this *C0) F29() _dafny.Array {
	{
		return _this._f29
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
