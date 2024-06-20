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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(421), _dafny.IntOfInt64(316))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(421)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(316)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_0_v0, _dafny.IntOfInt64(-169)), false)
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(296), !(false)))
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('l')
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) bool {
	return !(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(573), _dafny.IntOfUint32((_dafny.SeqOf(!(!(true)), !(true), false)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-447)), _dafny.IntOfInt64(235)))).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), true)).Cardinality(), _dafny.IntOfInt64(271))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-452))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(321)
	}))).Cardinality()), _dafny.IntOfInt64(4)))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(369))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(369)), _2_v0) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_2_v0, _dafny.IntOfInt64(526)))
			}
		}
		return _coll1.ToSet()
	}()).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(261))).Cardinality()))), false)
}
func (_static *CompanionStruct_Default___) Fm4(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(((_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))).Cardinality()), _dafny.IntOfInt64(247))
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, false, true, true, !(false)), _dafny.IntOfInt64(376))).Cardinality()).Cmp(_dafny.IntOfInt64(-423)) < 0, true, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(117), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sebrnl")).Cardinality())), true)
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), true))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("j"), _dafny.UnicodeSeqOfUtf8Bytes("wfhne"))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(((Companion_D3_.Create_DC9_(_dafny.IntOfInt64(80))).Dtor_cf17()).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gy")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(true))).Union((_dafny.SetOf(true)).Union(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Map, p1 bool, p2 bool, p3 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(972), (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(127))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(211))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xom")).Cardinality()), _dafny.IntOfInt64(419))).Cardinality())).Cardinality())).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v0 _dafny.Sequence
			_4_v0 = interface{}(_compr_2).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(127))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(211))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xom")).Cardinality()), _dafny.IntOfInt64(419))).Cardinality())).Cardinality())).Contains(_4_v0) {
				_coll2.Add(_4_v0, true)
			}
		}
		return _coll2.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(785))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_5_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality()), _dafny.IntOfInt64(-884)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(849))).Cardinality())))).Cardinality()), _dafny.CodePoint('r'), (_dafny.IntOfInt64(-240)).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(534))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_6_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(593))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_7_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}))).Cardinality())
	}))).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_8_i4 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_9_i5 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(352)
			}))).Cardinality())
		}))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _10_v1 _dafny.Int
			_10_v1 = interface{}(_compr_3).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_8_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}(func(_9_i5 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(352)
				}))).Cardinality())
			})), _10_v1) {
				_coll3.Add(Companion_Default___.SafeDivisionInt(_10_v1, _dafny.IntOfInt64(967)), true)
			}
		}
		return _coll3.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(-539)), (Companion_D1_.Create_DC2_(_dafny.IntOfInt64(70), _dafny.CodePoint('v'), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(515))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_11_i6 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-71), _dafny.IntOfInt64(213))).Cardinality())
	}))).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(66)), _dafny.CodePoint('d'))).Dtor_cf7()), _dafny.CodePoint('m'))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 _dafny.Int, p2 D3, p3 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-935), _dafny.IntOfInt64(330))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _12_v0 _dafny.Int
			_12_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(-935)).Cmp(_12_v0) <= 0) && ((_12_v0).Cmp(_dafny.IntOfInt64(330)) < 0) {
				_coll4.Add((_12_v0).Times((_dafny.SetOf(!(false))).Cardinality()), (_dafny.MultiSetOf(false)).Difference(_dafny.MultiSetOf(!(!(true)))))
			}
		}
		return _coll4.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool, bool, _dafny.Set) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 bool = false
	_ = r2
	var r3 _dafny.Set = _dafny.EmptySet
	_ = r3
	var _13_v0 bool
	_ = _13_v0
	_13_v0 = false
	var _14_v1 *C0
	_ = _14_v1
	var _nw0 *C0 = New_C0_()
	_ = _nw0
	_nw0.Ctor__()
	_14_v1 = _nw0
	var _15_v2 _dafny.Map
	_ = _15_v2
	_15_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _13_v0)
	var _16_v3 _dafny.Map
	_ = _16_v3
	_16_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v1, _15_v2)
	r1 = (p0).Cmp((_dafny.IntOfUint32((_dafny.SeqOf(_13_v0)).Cardinality())).Times((_16_v3).Cardinality())) >= 0
	var _17_v4 _dafny.Sequence
	_ = _17_v4
	_17_v4 = _dafny.UnicodeSeqOfUtf8Bytes("d")
	var _18_v5 D3
	_ = _18_v5
	_18_v5 = Companion_D3_.Create_DC10_(!(_13_v0), _17_v4, _13_v0)
	var _19_v6 D3
	_ = _19_v6
	_19_v6 = Companion_D3_.Create_DC11_(_18_v5)
	_19_v6 = _19_v6
	var _20_v7 _dafny.Set
	_ = _20_v7
	_20_v7 = _dafny.SetOf(_dafny.IntOfUint32((_17_v4).Cardinality()))
	var _21_i0 _dafny.Int
	_ = _21_i0
	_21_i0 = _dafny.Zero
	{
		for (_20_v7).IsProperSubsetOf(_20_v7) {
			{
				if (_21_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_21_i0 = (_21_i0).Plus(_dafny.One)
				var _22_v8 _dafny.CodePoint
				_ = _22_v8
				_22_v8 = _dafny.CodePoint('n')
				_22_v8 = _22_v8
				if !(!(Companion_Default___.Fm2(globalState))) {
					var _23_v9 _dafny.Set
					_ = _23_v9
					_23_v9 = _dafny.SetOf(_13_v0)
					var _24_v10 _dafny.Sequence
					_ = _24_v10
					_24_v10 = _dafny.SeqOf(_23_v9)
					_24_v10 = _24_v10
					var _25_v11 D3
					_ = _25_v11
					_25_v11 = Companion_D3_.Create_DC10_(_13_v0, _dafny.UnicodeSeqOfUtf8Bytes("yfugb"), _13_v0)
					var _26_v12 _dafny.Array
					_ = _26_v12
					var _nwElement0_0 bool = _13_v0
					_ = _nwElement0_0
					var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(19))
					_ = _nw1
					(_nw1).ArraySet1(_nwElement0_0, 0)
					(_nw1).ArraySet1(_13_v0, 1)
					(_nw1).ArraySet1(_13_v0, 2)
					(_nw1).ArraySet1(true, 3)
					(_nw1).ArraySet1(_13_v0, 4)
					(_nw1).ArraySet1(Companion_Default___.Fm2(globalState), 5)
					(_nw1).ArraySet1(_13_v0, 6)
					(_nw1).ArraySet1(_13_v0, 7)
					(_nw1).ArraySet1(_13_v0, 8)
					(_nw1).ArraySet1(_13_v0, 9)
					(_nw1).ArraySet1(_13_v0, 10)
					(_nw1).ArraySet1(_13_v0, 11)
					(_nw1).ArraySet1(_13_v0, 12)
					(_nw1).ArraySet1(_13_v0, 13)
					(_nw1).ArraySet1(_13_v0, 14)
					(_nw1).ArraySet1(_13_v0, 15)
					(_nw1).ArraySet1(_13_v0, 16)
					(_nw1).ArraySet1(_13_v0, 17)
					(_nw1).ArraySet1(_13_v0, 18)
					_26_v12 = _nw1
					var _27_v13 _dafny.Map
					_ = _27_v13
					_27_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v12, _13_v0)
					var _pat_let_tv0 = _27_v13
					_ = _pat_let_tv0
					var _pat_let_tv1 = _26_v12
					_ = _pat_let_tv1
					var _pat_let_tv2 = _27_v13
					_ = _pat_let_tv2
					var _pat_let_tv3 = _26_v12
					_ = _pat_let_tv3
					var _pat_let_tv4 = _13_v0
					_ = _pat_let_tv4
					var _pat_let_tv5 = _20_v7
					_ = _pat_let_tv5
					var _pat_let_tv6 = _20_v7
					_ = _pat_let_tv6
					_25_v11 = func(_pat_let0_0 D3) D3 {
						return func(_30_dt__update__tmp_h1 D3) D3 {
							return func(_pat_let3_0 bool) D3 {
								return func(_31_dt__update_hcf20_h1 bool) D3 {
									return Companion_D3_.Create_DC10_((_30_dt__update__tmp_h1).Dtor_cf18(), (_30_dt__update__tmp_h1).Dtor_cf19(), _31_dt__update_hcf20_h1)
								}(_pat_let3_0)
							}((_pat_let_tv5).IsProperSubsetOf(_pat_let_tv6))
						}(_pat_let0_0)
					}(func(_pat_let1_0 D3) D3 {
						return func(_28_dt__update__tmp_h0 D3) D3 {
							return func(_pat_let2_0 bool) D3 {
								return func(_29_dt__update_hcf20_h0 bool) D3 {
									return Companion_D3_.Create_DC10_((_28_dt__update__tmp_h0).Dtor_cf18(), (_28_dt__update__tmp_h0).Dtor_cf19(), _29_dt__update_hcf20_h0)
								}(_pat_let2_0)
							}(!((func() bool {
								if (_pat_let_tv0).Contains(_pat_let_tv1) {
									return (_pat_let_tv2).Get(_pat_let_tv3).(bool)
								}
								return _pat_let_tv4
							})()))
						}(_pat_let1_0)
					}(_25_v11))
					_13_v0 = _13_v0
					_14_v1 = _14_v1
					var _32_v14 D5
					_ = _32_v14
					_32_v14 = Companion_D5_.Create_DC14_(_13_v0, _13_v0, _14_v1, _dafny.IntOfInt64(151))
					var _pat_let_tv7 = _15_v2
					_ = _pat_let_tv7
					var _pat_let_tv8 = _13_v0
					_ = _pat_let_tv8
					var _pat_let_tv9 = _15_v2
					_ = _pat_let_tv9
					var _pat_let_tv10 = _13_v0
					_ = _pat_let_tv10
					var _pat_let_tv11 = _13_v0
					_ = _pat_let_tv11
					var _33_v15 _dafny.Sequence
					_ = _33_v15
					_33_v15 = _dafny.SeqOf((func(_pat_let4_0 D5) D5 {
						return func(_34_dt__update__tmp_h2 D5) D5 {
							return func(_pat_let5_0 bool) D5 {
								return func(_35_dt__update_hcf25_h0 bool) D5 {
									return Companion_D5_.Create_DC14_((_34_dt__update__tmp_h2).Dtor_cf24(), _35_dt__update_hcf25_h0, (_34_dt__update__tmp_h2).Dtor_cf26(), (_34_dt__update__tmp_h2).Dtor_cf27())
								}(_pat_let5_0)
							}((func() bool {
								if (_pat_let_tv7).Contains(_pat_let_tv8) {
									return (_pat_let_tv9).Get(_pat_let_tv10).(bool)
								}
								return _pat_let_tv11
							})())
						}(_pat_let4_0)
					}(_32_v14)).Dtor_cf26())
					_14_v1 = (_33_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_33_v15).Cardinality()))).Uint32()).(*C0)
				} else {
					var _36_v16 _dafny.Array
					_ = _36_v16
					var _nw2 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
					_ = _nw2
					_36_v16 = _nw2
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))
					_ = _index0
					(_36_v16).ArraySet1((p0).Cmp(p0) > 0, (_index0).Int())
					var _37_v17 _dafny.MultiSet
					_ = _37_v17
					_37_v17 = _dafny.MultiSetOf(!(_13_v0), _13_v0)
					var _38_v18 _dafny.Set
					_ = _38_v18
					_38_v18 = _dafny.SetOf(_13_v0, _13_v0)
					var _39_v19 _dafny.Array
					_ = _39_v19
					var _nwElement0_1 _dafny.Int = (p0).Minus(p0)
					_ = _nwElement0_1
					var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(5))
					_ = _nw3
					(_nw3).ArraySet1(_nwElement0_1, 0)
					(_nw3).ArraySet1((_37_v17).Cardinality(), 1)
					(_nw3).ArraySet1(p0, 2)
					(_nw3).ArraySet1(p0, 3)
					(_nw3).ArraySet1((func() _dafny.Int {
						if false {
							return (_dafny.Zero).Minus((_38_v18).Cardinality())
						}
						return p0
					})(), 4)
					_39_v19 = _nw3
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))
					_ = _index1
					(_39_v19).ArraySet1(p0, (_index1).Int())
					var _40_v20 _dafny.Array
					_ = _40_v20
					var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
					_ = _nw4
					_40_v20 = _nw4
					var _41_v21 _dafny.Map
					_ = _41_v21
					_41_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _40_v20)
					var _42_v22 _dafny.Map
					_ = _42_v22
					_42_v22 = _41_v21
					var _43_v23 _dafny.MultiSet
					_ = _43_v23
					_43_v23 = _dafny.MultiSetOf(p0)
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_39_v19), 0))
					_ = _index2
					(_39_v19).ArraySet1(((_42_v22).Cardinality()).Times((_43_v23).Cardinality()), (_index2).Int())
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))
					_ = _index3
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))
					_ = _index4
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_39_v19), 0))
					_ = _index5
					var _rhs0 bool = _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm7(p0, globalState), _17_v4)
					_ = _rhs0
					var _rhs1 _dafny.Int = p0
					_ = _rhs1
					var _rhs2 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-370), p0)
					_ = _rhs2
					var _rhs3 _dafny.Int = (_dafny.Zero).Minus(p0)
					_ = _rhs3
					var _lhs0 _dafny.Array = _36_v16
					_ = _lhs0
					var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))
					_ = _lhs1
					var _lhs2 _dafny.Array = _39_v19
					_ = _lhs2
					var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))
					_ = _lhs3
					var _lhs4 *GlobalState = globalState
					_ = _lhs4
					var _lhs5 _dafny.Array = _39_v19
					_ = _lhs5
					var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_39_v19), 0))
					_ = _lhs6
					(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
					(_lhs2).ArraySet1(_rhs1, (_lhs3).Int())
					_lhs4.F2 = _rhs2
					(_lhs5).ArraySet1(_rhs3, (_lhs6).Int())
					var _44_v24 _dafny.Array
					_ = _44_v24
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_0
					var _nw5 _dafny.Array
					_ = _nw5
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw5 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) _dafny.Sequence = (func(_45_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
							return func(_46_i1 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-298))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg11 _dafny.Int) interface{} {
										return coer11(arg11)
									}
								}((func(_47_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_48_i2 _dafny.Int) _dafny.CodePoint {
										return _47_v8
									}
								})(_45_v8)))
							}
						})(_22_v8)
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw5 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw5).ArraySet1(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw5).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_44_v24 = _nw5
					var _rhs4 _dafny.Array = _44_v24
					_ = _rhs4
					var _rhs5 bool = (_36_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))).Int()).(bool)
					_ = _rhs5
					var _rhs6 _dafny.MultiSet = (_43_v23).Difference(_43_v23)
					_ = _rhs6
					var _rhs7 _dafny.Sequence = Companion_Default___.Fm7((func() _dafny.Int {
						if (_36_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))).Int()).(bool) {
							return (_39_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))).Int()).(_dafny.Int)
						}
						return (_39_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))).Int()).(_dafny.Int)
					})(), globalState)
					_ = _rhs7
					var _lhs7 *GlobalState = globalState
					_ = _lhs7
					_44_v24 = _rhs4
					r1 = _rhs5
					_43_v23 = _rhs6
					_lhs7.F11 = _rhs7
					var _49_v25 _dafny.Sequence
					_ = _49_v25
					_49_v25 = _dafny.SeqOf((_39_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))).Int()).(_dafny.Int))
					var _50_v26 _dafny.Sequence
					_ = _50_v26
					_50_v26 = _dafny.SeqOf(p0, (_39_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_49_v25).Cardinality()))
					var _51_v27 _dafny.Sequence
					_ = _51_v27
					_51_v27 = _dafny.SeqOf(_50_v26, _49_v25)
					var _52_v28 _dafny.MultiSet
					_ = _52_v28
					_52_v28 = _dafny.MultiSetOf(_50_v26, (_51_v27).Select((Companion_Default___.SafeIndex((_39_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_39_v19), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_51_v27).Cardinality()))).Uint32()).(_dafny.Sequence))
					var _53_v29 _dafny.Map
					_ = _53_v29
					_53_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_v4, _52_v28)
					_53_v29 = (_53_v29).Update(_17_v4, _52_v28)
					var _rhs8 *C0 = _14_v1
					_ = _rhs8
					var _rhs9 bool = (_36_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))).Int()).(bool)
					_ = _rhs9
					var _rhs10 bool = (func() bool {
						if (_36_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))).Int()).(bool) {
							return !(_13_v0)
						}
						return (_36_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_36_v16), 0))).Int()).(bool)
					})()
					_ = _rhs10
					_14_v1 = _rhs8
					r1 = _rhs9
					_13_v0 = _rhs10
					var _54_v30 *C0
					_ = _54_v30
					var _nw6 *C0 = New_C0_()
					_ = _nw6
					_nw6.Ctor__()
					_54_v30 = _nw6
				}
				var _55_v32 _dafny.Sequence
				_ = _55_v32
				_55_v32 = _dafny.SeqOf(p0, (_dafny.Zero).Minus(p0))
				var _56_v33 D1
				_ = _56_v33
				_56_v33 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(214), _22_v8, (func() _dafny.Map {
					var _coll5 = _dafny.NewMapBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_20_v7).Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _57_v31 _dafny.Int
						_57_v31 = interface{}(_compr_5).(_dafny.Int)
						if (_20_v7).Contains(_57_v31) {
							_coll5.Add((_57_v31).Plus(p0), (func() bool {
								if (_15_v2).Contains(true) {
									return (_15_v2).Get(true).(bool)
								}
								return _13_v0
							})())
						}
					}
					return _coll5.ToMap()
				}()).Cardinality(), _55_v32, _22_v8)
				var _source0 D1 = Companion_D1_.Create_DC4_(_56_v33)
				_ = _source0
				if _source0.Is_DC2() {
					var _58___mcc_h0 _dafny.Int = _source0.Get_().(D1_DC2).Cf4
					_ = _58___mcc_h0
					var _59___mcc_h1 _dafny.CodePoint = _source0.Get_().(D1_DC2).Cf5
					_ = _59___mcc_h1
					var _60___mcc_h2 _dafny.Int = _source0.Get_().(D1_DC2).Cf6
					_ = _60___mcc_h2
					var _61___mcc_h3 _dafny.Sequence = _source0.Get_().(D1_DC2).Cf7
					_ = _61___mcc_h3
					var _62___mcc_h4 _dafny.CodePoint = _source0.Get_().(D1_DC2).Cf8
					_ = _62___mcc_h4
					var _63_cf8 _dafny.CodePoint = _62___mcc_h4
					_ = _63_cf8
					var _64_cf7 _dafny.Sequence = _61___mcc_h3
					_ = _64_cf7
					var _65_cf6 _dafny.Int = _60___mcc_h2
					_ = _65_cf6
					var _66_cf5 _dafny.CodePoint = _59___mcc_h1
					_ = _66_cf5
					var _67_cf4 _dafny.Int = _58___mcc_h0
					_ = _67_cf4
					var _68_v34 _dafny.Array
					_ = _68_v34
					var _len0_1 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_1
					var _nw7 _dafny.Array
					_ = _nw7
					if _len0_1.Cmp(_dafny.Zero) == 0 {
						_nw7 = _dafny.NewArray(_len0_1)
					} else {
						var _init1 func(_dafny.Int) _dafny.Int = (func(_69_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_70_i3 _dafny.Int) _dafny.Int {
								return (_70_i3).Plus(_69_cf4)
							}
						})(_67_cf4)
						_ = _init1
						var _element0_1 = _init1(_dafny.Zero)
						_ = _element0_1
						_nw7 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
						(_nw7).ArraySet1(_element0_1, 0)
						var _nativeLen0_1 = (_len0_1).Int()
						_ = _nativeLen0_1
						for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
							(_nw7).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
						}
					}
					_68_v34 = _nw7
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_68_v34), 0))
					_ = _index6
					(_68_v34).ArraySet1(_67_cf4, (_index6).Int())
					var _71_v35 D3
					_ = _71_v35
					_71_v35 = Companion_D3_.Create_DC9_(_67_cf4)
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_68_v34), 0))
					_ = _index7
					(_68_v34).ArraySet1((_71_v35).Dtor_cf17(), (_index7).Int())
					var _72_v36 *C0
					_ = _72_v36
					var _nw8 *C0 = New_C0_()
					_ = _nw8
					_nw8.Ctor__()
					_72_v36 = _nw8
					var _73_v37 _dafny.Map
					_ = _73_v37
					_73_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_cf4, _65_cf6)
					_73_v37 = (_73_v37).Update(_65_cf6, (_65_cf6).Times(_67_cf4))
					var _74_v38 _dafny.Array
					_ = _74_v38
					var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
					_ = _nw9
					_74_v38 = _nw9
					var _75_v39 _dafny.Array
					_ = _75_v39
					var _nwElement0_2 bool = Companion_Default___.Fm2(globalState)
					_ = _nwElement0_2
					var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(6))
					_ = _nw10
					(_nw10).ArraySet1(_nwElement0_2, 0)
					(_nw10).ArraySet1(_13_v0, 1)
					(_nw10).ArraySet1(_13_v0, 2)
					(_nw10).ArraySet1(_13_v0, 3)
					(_nw10).ArraySet1(_13_v0, 4)
					(_nw10).ArraySet1(_13_v0, 5)
					_75_v39 = _nw10
					var _76_v40 _dafny.Map
					_ = _76_v40
					_76_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v39, _72_v36)
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_74_v38), 0))
					_ = _index8
					(_74_v38).ArraySet1(_76_v40, (_index8).Int())
					var _77_v41 _dafny.Sequence
					_ = _77_v41
					_77_v41 = _dafny.SeqOf(_13_v0, _13_v0)
					var _78_v42 D2
					_ = _78_v42
					_78_v42 = Companion_D2_.Create_DC5_(_77_v41)
					var _79_v43 _dafny.Map
					_ = _79_v43
					_79_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v42, _64_cf7)
					var _80_v44 _dafny.Map
					_ = _80_v44
					_80_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v36, (_65_cf6).Times(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_79_v43).Contains(_78_v42) {
							return (_79_v43).Get(_78_v42).(_dafny.Sequence)
						}
						return _64_cf7
					})()).Cardinality())))
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_74_v38), 0))
					_ = _index9
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_68_v34), 0))
					_ = _index10
					var _rhs11 *C0 = _14_v1
					_ = _rhs11
					var _rhs12 _dafny.Map = _76_v40
					_ = _rhs12
					var _rhs13 _dafny.Int = (((_20_v7).Union(_20_v7)).Cardinality()).Minus((_dafny.Zero).Minus(p0))
					_ = _rhs13
					var _rhs14 _dafny.Int = _67_cf4
					_ = _rhs14
					var _rhs15 _dafny.Map = _80_v44
					_ = _rhs15
					var _lhs8 _dafny.Array = _74_v38
					_ = _lhs8
					var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_74_v38), 0))
					_ = _lhs9
					var _lhs10 *GlobalState = globalState
					_ = _lhs10
					var _lhs11 _dafny.Array = _68_v34
					_ = _lhs11
					var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_68_v34), 0))
					_ = _lhs12
					_72_v36 = _rhs11
					(_lhs8).ArraySet1(_rhs12, (_lhs9).Int())
					_lhs10.F2 = _rhs13
					(_lhs11).ArraySet1(_rhs14, (_lhs12).Int())
					_80_v44 = _rhs15
				} else if _source0.Is_DC3() {
					var _81___mcc_h5 bool = _source0.Get_().(D1_DC3).Cf9
					_ = _81___mcc_h5
					var _82___mcc_h6 bool = _source0.Get_().(D1_DC3).Cf10
					_ = _82___mcc_h6
					var _83___mcc_h7 _dafny.Int = _source0.Get_().(D1_DC3).Cf11
					_ = _83___mcc_h7
					var _84___mcc_h8 bool = _source0.Get_().(D1_DC3).Cf12
					_ = _84___mcc_h8
					var _85_cf12 bool = _84___mcc_h8
					_ = _85_cf12
					var _86_cf11 _dafny.Int = _83___mcc_h7
					_ = _86_cf11
					var _87_cf10 bool = _82___mcc_h6
					_ = _87_cf10
					var _88_cf9 bool = _81___mcc_h5
					_ = _88_cf9
					r1 = _88_cf9
					var _89_v45 *C0
					_ = _89_v45
					var _nw11 *C0 = New_C0_()
					_ = _nw11
					_nw11.Ctor__()
					_89_v45 = _nw11
					var _90_v47 _dafny.Array
					_ = _90_v47
					var _len0_2 _dafny.Int = _dafny.IntOfInt64(19)
					_ = _len0_2
					var _nw12 _dafny.Array
					_ = _nw12
					if _len0_2.Cmp(_dafny.Zero) == 0 {
						_nw12 = _dafny.NewArray(_len0_2)
					} else {
						var _init2 func(_dafny.Int) _dafny.Map = (func(_91_cf10 bool, _92_p0 _dafny.Int, _93_cf9 bool) func(_dafny.Int) _dafny.Map {
							return func(_94_i4 _dafny.Int) _dafny.Map {
								return (func() _dafny.Map {
									if !(_91_cf10) {
										return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_p0, _dafny.MultiSetOf(_93_cf9, _91_cf10))
									}
									return func() _dafny.Map {
										var _coll6 = _dafny.NewMapBuilder()
										_ = _coll6
										for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-475), _dafny.IntOfInt64(301))); ; {
											_compr_6, _ok6 := _iter6()
											if !_ok6 {
												break
											}
											var _95_v46 _dafny.Int
											_95_v46 = interface{}(_compr_6).(_dafny.Int)
											if ((_dafny.IntOfInt64(-475)).Cmp(_95_v46) <= 0) && ((_95_v46).Cmp(_dafny.IntOfInt64(301)) < 0) {
												_coll6.Add((_95_v46).Times(_92_p0), _dafny.MultiSetOf(_91_cf10))
											}
										}
										return _coll6.ToMap()
									}()
								})()
							}
						})(_87_cf10, p0, _88_cf9)
						_ = _init2
						var _element0_2 = _init2(_dafny.Zero)
						_ = _element0_2
						_nw12 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
						(_nw12).ArraySet1(_element0_2, 0)
						var _nativeLen0_2 = (_len0_2).Int()
						_ = _nativeLen0_2
						for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
							(_nw12).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
						}
					}
					_90_v47 = _nw12
					var _96_v49 _dafny.Set
					_ = _96_v49
					_96_v49 = _dafny.SetOf(_85_cf12, _87_cf10, _13_v0, _88_cf9, _87_cf10)
					var _97_v50 _dafny.Map
					_ = _97_v50
					_97_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_17_v4).Cardinality()), _96_v49)
					var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_90_v47), 0))
					_ = _index11
					(_90_v47).ArraySet1(func() _dafny.Map {
						var _coll7 = _dafny.NewMapBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate((_97_v50).Keys().Elements()); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _98_v48 _dafny.Int
							_98_v48 = interface{}(_compr_7).(_dafny.Int)
							if (_97_v50).Contains(_98_v48) {
								_coll7.Add((_98_v48).Times(_dafny.IntOfUint32((_dafny.SeqOf(_87_cf10)).Cardinality())), _dafny.MultiSetOf(_87_cf10, !(_85_cf12), _88_cf9))
							}
						}
						return _coll7.ToMap()
					}(), (_index11).Int())
					var _99_v51 _dafny.Sequence
					_ = _99_v51
					_99_v51 = _dafny.SeqOf(_85_cf12)
					var _100_v52 _dafny.MultiSet
					_ = _100_v52
					_100_v52 = _dafny.MultiSetOf(_86_cf11)
					var _101_v53 _dafny.Array
					_ = _101_v53
					var _nwElement0_3 _dafny.Int = p0
					_ = _nwElement0_3
					var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(11))
					_ = _nw13
					(_nw13).ArraySet1(_nwElement0_3, 0)
					(_nw13).ArraySet1(p0, 1)
					(_nw13).ArraySet1(_86_cf11, 2)
					(_nw13).ArraySet1(p0, 3)
					(_nw13).ArraySet1(_dafny.IntOfInt64(221), 4)
					(_nw13).ArraySet1(_86_cf11, 5)
					(_nw13).ArraySet1(_86_cf11, 6)
					(_nw13).ArraySet1(_dafny.IntOfUint32((_17_v4).Cardinality()), 7)
					(_nw13).ArraySet1(_86_cf11, 8)
					(_nw13).ArraySet1(_86_cf11, 9)
					(_nw13).ArraySet1(_86_cf11, 10)
					_101_v53 = _nw13
					var _102_v54 _dafny.Array
					_ = _102_v54
					_102_v54 = _101_v53
					var _103_v55 _dafny.Map
					_ = _103_v55
					_103_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v54, _13_v0)
					var _104_v56 D2
					_ = _104_v56
					_104_v56 = Companion_D2_.Create_DC7_(Companion_D2_.Create_DC7_(Companion_D2_.Create_DC6_()))
					var _105_v57 _dafny.Map
					_ = _105_v57
					_105_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(globalState), _104_v56)
					var _106_v58 _dafny.Array
					_ = _106_v58
					var _nwElement0_4 _dafny.Int = _dafny.IntOfInt64(858)
					_ = _nwElement0_4
					var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(23))
					_ = _nw14
					(_nw14).ArraySet1(_nwElement0_4, 0)
					(_nw14).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dnsnwbotg"), _dafny.UnicodeSeqOfUtf8Bytes("xvnoa"))).Cardinality()), 1)
					(_nw14).ArraySet1((func() _dafny.Int {
						if _85_cf12 {
							return _86_cf11
						}
						return _86_cf11
					})(), 2)
					(_nw14).ArraySet1(p0, 3)
					(_nw14).ArraySet1((_dafny.IntOfInt64(975)).Plus(_dafny.IntOfInt64(655)), 4)
					(_nw14).ArraySet1(p0, 5)
					(_nw14).ArraySet1(_86_cf11, 6)
					(_nw14).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_17_v4).Cardinality()), p0), 7)
					(_nw14).ArraySet1((_100_v52).Cardinality(), 8)
					(_nw14).ArraySet1(_86_cf11, 9)
					(_nw14).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality())).Times(p0), 10)
					(_nw14).ArraySet1((_20_v7).Cardinality(), 11)
					(_nw14).ArraySet1(p0, 12)
					(_nw14).ArraySet1((_103_v55).Cardinality(), 13)
					(_nw14).ArraySet1(p0, 14)
					(_nw14).ArraySet1((_86_cf11).Times(p0), 15)
					(_nw14).ArraySet1((func() _dafny.Int {
						if Companion_Default___.Fm2(globalState) {
							return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-769))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg12 _dafny.Int) interface{} {
									return coer12(arg12)
								}
							}((func(_107_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_108_i5 _dafny.Int) _dafny.CodePoint {
									return _107_v8
								}
							})(_22_v8)))).Cardinality())
						}
						return (_105_v57).Cardinality()
					})(), 16)
					(_nw14).ArraySet1(_86_cf11, 17)
					(_nw14).ArraySet1(_86_cf11, 18)
					(_nw14).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), _dafny.IntOfUint32((_17_v4).Cardinality())), 19)
					(_nw14).ArraySet1(_86_cf11, 20)
					(_nw14).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pjqydcgsg")).Cardinality()), 21)
					(_nw14).ArraySet1(_86_cf11, 22)
					_106_v58 = _nw14
					var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_106_v58), 0))
					_ = _index12
					(_106_v58).ArraySet1(p0, (_index12).Int())
					var _109_v59 D3
					_ = _109_v59
					_109_v59 = Companion_D3_.Create_DC9_(p0)
					var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_90_v47), 0))
					_ = _index13
					var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_106_v58), 0))
					_ = _index14
					var _rhs16 _dafny.Map = Companion_Default___.Fm11(!(_87_cf10) || (_13_v0), _dafny.IntOfInt64(749), _109_v59, !_dafny.Companion_Sequence_.Contains(_17_v4, _22_v8), globalState)
					_ = _rhs16
					var _rhs17 _dafny.Sequence = _99_v51
					_ = _rhs17
					var _rhs18 _dafny.Int = Companion_Default___.SafeModuloInt(_86_cf11, _86_cf11)
					_ = _rhs18
					var _lhs13 _dafny.Array = _90_v47
					_ = _lhs13
					var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_90_v47), 0))
					_ = _lhs14
					var _lhs15 _dafny.Array = _106_v58
					_ = _lhs15
					var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_106_v58), 0))
					_ = _lhs16
					(_lhs13).ArraySet1(_rhs16, (_lhs14).Int())
					_99_v51 = _rhs17
					(_lhs15).ArraySet1(_rhs18, (_lhs16).Int())
					_88_cf9 = _87_cf10
				} else if _source0.Is_DC1() {
					var _110___mcc_h9 bool = _source0.Get_().(D1_DC1).Cf3
					_ = _110___mcc_h9
					var _111_cf3 bool = _110___mcc_h9
					_ = _111_cf3
					_15_v2 = (_15_v2).Update(Companion_Default___.Fm2(globalState), _111_cf3)
					var _112_v60 *C0
					_ = _112_v60
					var _nw15 *C0 = New_C0_()
					_ = _nw15
					_nw15.Ctor__()
					_112_v60 = _nw15
					var _113_v61 _dafny.Array
					_ = _113_v61
					var _nwElement0_5 _dafny.Map = (_15_v2).Merge(_15_v2)
					_ = _nwElement0_5
					var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(4))
					_ = _nw16
					(_nw16).ArraySet1(_nwElement0_5, 0)
					(_nw16).ArraySet1(_15_v2, 1)
					(_nw16).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _111_cf3), 2)
					(_nw16).ArraySet1(_15_v2, 3)
					_113_v61 = _nw16
					_113_v61 = _113_v61
					var _114_v62 _dafny.Array
					_ = _114_v62
					var _len0_3 _dafny.Int = _dafny.IntOfInt64(6)
					_ = _len0_3
					var _nw17 _dafny.Array
					_ = _nw17
					if _len0_3.Cmp(_dafny.Zero) == 0 {
						_nw17 = _dafny.NewArray(_len0_3)
					} else {
						var _init3 func(_dafny.Int) _dafny.Int = (func(_115_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_116_i6 _dafny.Int) _dafny.Int {
								return (_116_i6).Plus(_115_p0)
							}
						})(p0)
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
					_114_v62 = _nw17
					_114_v62 = _114_v62
				} else {
					var _117___mcc_h10 D1 = _source0.Get_().(D1_DC4).Cf13
					_ = _117___mcc_h10
					var _118_cf13 D1 = _117___mcc_h10
					_ = _118_cf13
					(globalState).F2 = p0
					var _119_v63 _dafny.Array
					_ = _119_v63
					var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
					_ = _nw18
					_119_v63 = _nw18
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_119_v63), 0))
					_ = _index15
					(_119_v63).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lbrfdtt"), (_index15).Int())
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_119_v63), 0))
					_ = _index16
					(_119_v63).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rkrne"), (_index16).Int())
					var _120_v64 *C0
					_ = _120_v64
					var _nw19 *C0 = New_C0_()
					_ = _nw19
					_nw19.Ctor__()
					_120_v64 = _nw19
					var _rhs19 _dafny.Sequence = _55_v32
					_ = _rhs19
					var _rhs20 bool = _13_v0
					_ = _rhs20
					_55_v32 = _rhs19
					_13_v0 = _rhs20
				}
				_22_v8 = _22_v8
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _121_v65 _dafny.Array
	_ = _121_v65
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(17)
	_ = _len0_4
	var _nw20 _dafny.Array
	_ = _nw20
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw20 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) bool = (func(_122_v4 _dafny.Sequence) func(_dafny.Int) bool {
			return func(_123_i8 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.IsProperPrefixOf((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(100))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg13 _dafny.Int) interface{} {
						return coer13(arg13)
					}
				}(func(_124_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				})), _122_v4)), _dafny.SeqOf(_122_v4))
			}
		})(_17_v4)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw20 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw20).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw20).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_121_v65 = _nw20
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_121_v65), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _125_i7 _dafny.Int
		_125_i7 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_125_i7).Sign() != -1) && ((_125_i7).Cmp(_dafny.ArrayLen((_121_v65), 0)) < 0)) {
			(_121_v65).ArraySet1((_dafny.IntOfInt64(613)).Cmp(((_15_v2).Cardinality()).Minus(p0)) < 0, (_125_i7).Int())
		}
	}
	var _126_v66 _dafny.Array
	_ = _126_v66
	var _nw21 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
	_ = _nw21
	_126_v66 = _nw21
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_126_v66), 0))
	_ = _index17
	(_126_v66).ArraySet1(_14_v1, (_index17).Int())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_126_v66), 0))
	_ = _index18
	(_126_v66).ArraySet1(_14_v1, (_index18).Int())
	var _127_v67 _dafny.Array
	_ = _127_v67
	var _nwElement0_6 _dafny.Array = _121_v65
	_ = _nwElement0_6
	var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(4))
	_ = _nw22
	(_nw22).ArraySet1(_nwElement0_6, 0)
	(_nw22).ArraySet1(_121_v65, 1)
	(_nw22).ArraySet1(_121_v65, 2)
	(_nw22).ArraySet1(_121_v65, 3)
	_127_v67 = _nw22
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_127_v67), 0))
	_ = _index19
	(_127_v67).ArraySet1(_121_v65, (_index19).Int())
	var _128_v68 D3
	_ = _128_v68
	_128_v68 = Companion_D3_.Create_DC8_(_121_v65)
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_127_v67), 0))
	_ = _index20
	(_127_v67).ArraySet1((_128_v68).Dtor_cf16(), (_index20).Int())
	var _129_v70 D1
	_ = _129_v70
	_129_v70 = Companion_D1_.Create_DC3_(_13_v0, _13_v0, p0, _13_v0)
	var _130_v71 _dafny.MultiSet
	_ = _130_v71
	_130_v71 = _dafny.MultiSetOf(_129_v70)
	var _131_v72 _dafny.Sequence
	_ = _131_v72
	_131_v72 = _dafny.SeqOf(func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter9 := _dafny.Iterate((_130_v71).Elements()); ; {
			_compr_8, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _132_v69 D1
			_132_v69 = interface{}(_compr_8).(D1)
			if (_130_v71).Contains(_132_v69) {
				_coll8.Add(_132_v69, _13_v0)
			}
		}
		return _coll8.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v70, true))
	r0 = _131_v72
	r1 = !(_13_v0)
	r2 = _13_v0
	var _133_v73 _dafny.Set
	_ = _133_v73
	_133_v73 = _dafny.SetOf((func() bool {
		if _13_v0 {
			return _13_v0
		}
		return _13_v0
	})())
	r3 = _133_v73
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _134_v0 _dafny.Array
	_ = _134_v0
	var _nw23 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
	_ = _nw23
	_134_v0 = _nw23
	var _135_v1 _dafny.Array
	_ = _135_v1
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(21)
	_ = _len0_5
	var _nw24 _dafny.Array
	_ = _nw24
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw24 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Set = func(_136_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(_dafny.IntOfInt64(-34))
		}
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw24 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw24).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw24).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_135_v1 = _nw24
	var _137_globalState *GlobalState
	_ = _137_globalState
	var _nw25 *GlobalState = New_GlobalState_()
	_ = _nw25
	_nw25.Ctor__(false, _134_v0, _dafny.IntOfInt64(212), false, _dafny.IntOfInt64(3), true, false, _dafny.IntOfInt64(743), true, _135_v1, _dafny.IntOfInt64(700), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(199))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_138_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	})))
	_137_globalState = _nw25
	var _139_v2 bool
	_ = _139_v2
	_139_v2 = true
	if _139_v2 {
		var _140_v3 _dafny.Int
		_ = _140_v3
		_140_v3 = _dafny.IntOfInt64(249)
		var _141_v4 _dafny.Map
		_ = _141_v4
		_141_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v3, _139_v2)
		var _142_v5 _dafny.Sequence
		_ = _142_v5
		_142_v5 = _dafny.SeqOf((Companion_Default___.Fm0(_137_globalState)).Merge(_141_v4))
		_142_v5 = (func() _dafny.Sequence {
			if _139_v2 {
				return _142_v5
			}
			return _142_v5
		})()
		var _143_v6 _dafny.Map
		_ = _143_v6
		_143_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_140_v3))
		var _144_v7 _dafny.Sequence
		_ = _144_v7
		_144_v7 = _dafny.UnicodeSeqOfUtf8Bytes("xbl")
		var _145_v8 _dafny.MultiSet
		_ = _145_v8
		_145_v8 = _dafny.MultiSetOf(_dafny.IntOfInt64(-795), _dafny.IntOfUint32((_144_v7).Cardinality()), _140_v3, _140_v3, _140_v3)
		var _146_v9 D0
		_ = _146_v9
		_146_v9 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus((_143_v6).Cardinality()), _145_v8, _140_v3)
		var _source1 D0 = _146_v9
		_ = _source1
		var _147___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC0).Cf0
		_ = _147___mcc_h0
		var _148___mcc_h1 _dafny.MultiSet = _source1.Get_().(D0_DC0).Cf1
		_ = _148___mcc_h1
		var _149___mcc_h2 _dafny.Int = _source1.Get_().(D0_DC0).Cf2
		_ = _149___mcc_h2
		var _150_cf2 _dafny.Int = _149___mcc_h2
		_ = _150_cf2
		var _151_cf1 _dafny.MultiSet = _148___mcc_h1
		_ = _151_cf1
		var _152_cf0 _dafny.Int = _147___mcc_h0
		_ = _152_cf0
		var _153_v10 _dafny.CodePoint
		_ = _153_v10
		_153_v10 = _dafny.CodePoint('g')
		var _154_v11 _dafny.Array
		_ = _154_v11
		var _nwElement0_7 _dafny.CodePoint = _153_v10
		_ = _nwElement0_7
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(9))
		_ = _nw26
		(_nw26).ArraySet1CodePoint(_nwElement0_7, 0)
		(_nw26).ArraySet1CodePoint(_153_v10, 1)
		(_nw26).ArraySet1CodePoint(Companion_Default___.Fm1(_137_globalState), 2)
		(_nw26).ArraySet1CodePoint(_153_v10, 3)
		(_nw26).ArraySet1CodePoint(_153_v10, 4)
		(_nw26).ArraySet1CodePoint(_153_v10, 5)
		(_nw26).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _139_v2 {
				return _153_v10
			}
			return _153_v10
		})(), 6)
		(_nw26).ArraySet1CodePoint(Companion_Default___.Fm1(_137_globalState), 7)
		(_nw26).ArraySet1CodePoint(_153_v10, 8)
		_154_v11 = _nw26
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_154_v11), 0))
		_ = _index21
		(_154_v11).ArraySet1CodePoint(_153_v10, (_index21).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_154_v11), 0))
		_ = _index22
		(_154_v11).ArraySet1CodePoint(_153_v10, (_index22).Int())
		var _155_v12 _dafny.Map
		_ = _155_v12
		_155_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
		var _156_v13 _dafny.Sequence
		_ = _156_v13
		_156_v13 = _dafny.SeqOf((_dafny.Zero).Minus((_155_v12).Cardinality()))
		var _157_v14 _dafny.Array
		_ = _157_v14
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
		_ = _nw27
		_157_v14 = _nw27
		var _158_v15 _dafny.Set
		_ = _158_v15
		_158_v15 = _dafny.SetOf(_139_v2, _139_v2)
		var _159_v16 _dafny.Array
		_ = _159_v16
		var _nwElement0_8 _dafny.Set = _158_v15
		_ = _nwElement0_8
		var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.One)
		_ = _nw28
		(_nw28).ArraySet1(_nwElement0_8, 0)
		_159_v16 = _nw28
		var _160_v17 _dafny.Map
		_ = _160_v17
		_160_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_v16, _157_v14)
		var _161_v18 _dafny.Sequence
		_ = _161_v18
		_161_v18 = _dafny.SeqOf(_159_v16)
		var _162_v19 D1
		_ = _162_v19
		_162_v19 = Companion_D1_.Create_DC1_(_139_v2)
		var _rhs21 _dafny.Sequence = _156_v13
		_ = _rhs21
		var _rhs22 _dafny.Array = (func() _dafny.Array {
			if (_160_v17).Contains((_161_v18).Select((Companion_Default___.SafeIndex(_150_cf2, _dafny.IntOfUint32((_161_v18).Cardinality()))).Uint32()).(_dafny.Array)) {
				return (_160_v17).Get((_161_v18).Select((Companion_Default___.SafeIndex(_150_cf2, _dafny.IntOfUint32((_161_v18).Cardinality()))).Uint32()).(_dafny.Array)).(_dafny.Array)
			}
			return _157_v14
		})()
		_ = _rhs22
		var _rhs23 bool = _139_v2
		_ = _rhs23
		var _rhs24 bool = (func() bool {
			if (_150_cf2).Cmp(_150_cf2) == 0 {
				return _139_v2
			}
			return (_162_v19).Dtor_cf3()
		})()
		_ = _rhs24
		_156_v13 = _rhs21
		_157_v14 = _rhs22
		_139_v2 = _rhs23
		_139_v2 = _rhs24
		_140_v3 = _150_cf2
		var _163_v20 _dafny.Sequence
		_ = _163_v20
		_163_v20 = _dafny.SeqOf(_139_v2)
		var _164_v21 _dafny.Map
		_ = _164_v21
		_164_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v10, _139_v2)
		var _165_v22 _dafny.Sequence
		_ = _165_v22
		_165_v22 = _dafny.SeqOf((_163_v20).Select((Companion_Default___.SafeIndex(_140_v3, _dafny.IntOfUint32((_163_v20).Cardinality()))).Uint32()).(bool), false, (func() bool {
			if (_164_v21).Contains((_154_v11).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_154_v11), 0))).Int())) {
				return (_164_v21).Get((_154_v11).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_154_v11), 0))).Int())).(bool)
			}
			return _139_v2
		})(), !(_139_v2), !(_139_v2))
		var _166_v23 _dafny.Array
		_ = _166_v23
		var _nwElement0_9 _dafny.MultiSet = _151_cf1
		_ = _nwElement0_9
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(21))
		_ = _nw29
		(_nw29).ArraySet1(_nwElement0_9, 0)
		(_nw29).ArraySet1(_151_cf1, 1)
		(_nw29).ArraySet1(_151_cf1, 2)
		(_nw29).ArraySet1(_151_cf1, 3)
		(_nw29).ArraySet1(_dafny.MultiSetFromSeq(_156_v13), 4)
		(_nw29).ArraySet1(_151_cf1, 5)
		(_nw29).ArraySet1(_145_v8, 6)
		(_nw29).ArraySet1(_145_v8, 7)
		(_nw29).ArraySet1(_dafny.MultiSetFromSeq(_156_v13), 8)
		(_nw29).ArraySet1(_dafny.MultiSetOf(_150_cf2, _152_cf0), 9)
		(_nw29).ArraySet1(_145_v8, 10)
		(_nw29).ArraySet1(_145_v8, 11)
		(_nw29).ArraySet1(_151_cf1, 12)
		(_nw29).ArraySet1(_145_v8, 13)
		(_nw29).ArraySet1(_145_v8, 14)
		(_nw29).ArraySet1(_145_v8, 15)
		(_nw29).ArraySet1(_dafny.MultiSetOf(_150_cf2), 16)
		(_nw29).ArraySet1(_145_v8, 17)
		(_nw29).ArraySet1(_145_v8, 18)
		(_nw29).ArraySet1(_151_cf1, 19)
		(_nw29).ArraySet1(_145_v8, 20)
		_166_v23 = _nw29
		var _167_v24 _dafny.Map
		_ = _167_v24
		_167_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v22, _166_v23)
		_167_v24 = (_167_v24).Update(_163_v20, _166_v23)
		_141_v4 = (_141_v4).Update((_143_v6).Cardinality(), (!(_139_v2)) && (_139_v2))
		var _168_v25 _dafny.Array
		_ = _168_v25
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
		_ = _nw30
		_168_v25 = _nw30
		var _169_v26 _dafny.Map
		_ = _169_v26
		_169_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v25, (func() bool {
			if _139_v2 {
				return _139_v2
			}
			return _139_v2
		})())
		var _170_v29 _dafny.Map
		_ = _170_v29
		_170_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_144_v7).Cardinality()), _140_v3)
		var _171_v30 _dafny.MultiSet
		_ = _171_v30
		_171_v30 = _dafny.MultiSetOf(func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter10 := _dafny.Iterate((_145_v8).Elements()); ; {
				_compr_9, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _172_v27 _dafny.Int
				_172_v27 = interface{}(_compr_9).(_dafny.Int)
				if (_145_v8).Contains(_172_v27) {
					_coll9.Add((_172_v27).Minus(_140_v3), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32(((Companion_D2_.Create_DC5_(_dafny.SeqOf(_139_v2, _139_v2))).Dtor_cf14()).Cardinality()))).Cardinality()))).Cardinality())
				}
			}
			return _coll9.ToMap()
		}(), func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter11 := _dafny.Iterate((_141_v4).Keys().Elements()); ; {
				_compr_10, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _173_v28 _dafny.Int
				_173_v28 = interface{}(_compr_10).(_dafny.Int)
				if (_141_v4).Contains(_173_v28) {
					_coll10.Add(Companion_Default___.SafeDivisionInt(_173_v28, _140_v3), _140_v3)
				}
			}
			return _coll10.ToMap()
		}(), _170_v29)
		_169_v26 = (_169_v26).Update(_168_v25, (_171_v30).IsDisjointFrom(_171_v30))
		var _174_v31 _dafny.Map
		_ = _174_v31
		_174_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfInt64(725)), _dafny.SeqOf(_140_v3, _140_v3)), (_170_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v3, _140_v3)))
		_170_v29 = (func() _dafny.Map {
			if (_174_v31).Contains(_139_v2) {
				return (_174_v31).Get(_139_v2).(_dafny.Map)
			}
			return _170_v29
		})()
	} else {
		var _175_v32 _dafny.CodePoint
		_ = _175_v32
		_175_v32 = _dafny.CodePoint('i')
		_175_v32 = _175_v32
		_139_v2 = _139_v2
		var _176_v33 _dafny.Int
		_ = _176_v33
		_176_v33 = _dafny.IntOfInt64(880)
		var _177_v34 _dafny.Map
		_ = _177_v34
		_177_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_176_v33).Cmp(_176_v33) != 0, _dafny.IntOfInt64(636))
		_177_v34 = (_177_v34).Update(_139_v2, _176_v33)
		_175_v32 = _dafny.CodePoint('n')
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(535))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_178_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_179_i2 _dafny.Int) _dafny.CodePoint {
				return _178_v32
			}
		})(_175_v32))), _dafny.UnicodeSeqOfUtf8Bytes("jrvkbowy")) {
			var _180_v35 _dafny.MultiSet
			_ = _180_v35
			_180_v35 = _dafny.MultiSetOf(_175_v32, _175_v32)
			var _181_v36 _dafny.MultiSet
			_ = _181_v36
			_181_v36 = _dafny.MultiSetOf(_139_v2, Companion_Default___.Fm2(_137_globalState))
			var _182_v37 _dafny.Map
			_ = _182_v37
			_182_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v33, _176_v33)
			var _183_v38 _dafny.Array
			_ = _183_v38
			var _nwElement0_10 _dafny.Int = (_180_v35).Cardinality()
			_ = _nwElement0_10
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(10))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_10, 0)
			(_nw31).ArraySet1(_176_v33, 1)
			(_nw31).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(903))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_184_v33 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_185_i3 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_184_v33)
				}
			})(_176_v33)))).Cardinality()), 2)
			(_nw31).ArraySet1(_176_v33, 3)
			(_nw31).ArraySet1(_176_v33, 4)
			(_nw31).ArraySet1(_176_v33, 5)
			(_nw31).ArraySet1(_176_v33, 6)
			(_nw31).ArraySet1(_176_v33, 7)
			(_nw31).ArraySet1((_181_v36).Cardinality(), 8)
			(_nw31).ArraySet1((_182_v37).Cardinality(), 9)
			_183_v38 = _nw31
			var _186_v39 _dafny.Set
			_ = _186_v39
			_186_v39 = _dafny.SetOf(_183_v38, _183_v38)
			var _rhs25 _dafny.Int = (_186_v39).Cardinality()
			_ = _rhs25
			var _rhs26 bool = Companion_Default___.Fm2(_137_globalState)
			_ = _rhs26
			var _rhs27 bool = Companion_Default___.Fm2(_137_globalState)
			_ = _rhs27
			var _rhs28 bool = _139_v2
			_ = _rhs28
			var _rhs29 _dafny.Int = _176_v33
			_ = _rhs29
			var _lhs17 *GlobalState = _137_globalState
			_ = _lhs17
			var _lhs18 *GlobalState = _137_globalState
			_ = _lhs18
			_lhs17.F2 = _rhs25
			_139_v2 = _rhs26
			_139_v2 = _rhs27
			_139_v2 = _rhs28
			_lhs18.F2 = _rhs29
			var _187_v40 _dafny.Array
			_ = _187_v40
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw32
			_187_v40 = _nw32
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_187_v40), 0))
			_ = _index23
			(_187_v40).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(266))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_188_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_189_i4 _dafny.Int) _dafny.CodePoint {
					return _188_v32
				}
			})(_175_v32))), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_187_v40), 0))
			_ = _index24
			(_187_v40).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yjcvetkn"), (_index24).Int())
			var _190_v41 _dafny.Array
			_ = _190_v41
			var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw33
			_190_v41 = _nw33
			var _191_v42 _dafny.Sequence
			_ = _191_v42
			_191_v42 = _dafny.SeqOf(_dafny.SeqOf(_176_v33, _176_v33))
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_190_v41), 0))
			_ = _index25
			(_190_v41).ArraySet1(_191_v42, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_190_v41), 0))
			_ = _index26
			(_190_v41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_191_v42, _191_v42), _dafny.Companion_Sequence_.Concatenate(_191_v42, _191_v42)), (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_183_v38), 0))
			_ = _index27
			(_183_v38).ArraySet1(_176_v33, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_183_v38), 0))
			_ = _index28
			(_183_v38).ArraySet1(_dafny.IntOfInt64(59), (_index28).Int())
			var _192_v43 _dafny.Sequence
			_ = _192_v43
			_192_v43 = _dafny.SeqOf(_dafny.IntOfInt64(-970))
			(_137_globalState).F2 = (func() _dafny.Int {
				if (_181_v36).Contains(_dafny.Companion_Sequence_.IsProperPrefixOf(_192_v43, _192_v43)) {
					return (_181_v36).Multiplicity(_dafny.Companion_Sequence_.IsProperPrefixOf(_192_v43, _192_v43))
				}
				return (_dafny.Zero).Minus(_176_v33)
			})()
		} else {
			var _193_v44 _dafny.MultiSet
			_ = _193_v44
			_193_v44 = _dafny.MultiSetOf(!(Companion_Default___.Fm2(_137_globalState)))
			(_137_globalState).F2 = (_193_v44).Cardinality()
			var _194_v45 _dafny.Map
			_ = _194_v45
			_194_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v2, _139_v2)
			_139_v2 = ((_194_v45).Cardinality()).Cmp(_176_v33) == 0
			_194_v45 = (_194_v45).Update(_139_v2, !(_139_v2) || (_139_v2))
			var _195_v46 *C0
			_ = _195_v46
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__()
			_195_v46 = _nw34
			var _rhs30 bool = _139_v2
			_ = _rhs30
			var _rhs31 _dafny.Int = (_176_v33).Times(_176_v33)
			_ = _rhs31
			var _rhs32 _dafny.Map = Companion_Default___.Fm3(_139_v2, _dafny.MultiSetOf(false, _139_v2, _139_v2), _139_v2, _176_v33, _137_globalState)
			_ = _rhs32
			var _rhs33 _dafny.Int = (((_194_v45).Merge(_194_v45)).Merge(_194_v45)).Cardinality()
			_ = _rhs33
			var _lhs19 *GlobalState = _137_globalState
			_ = _lhs19
			_139_v2 = _rhs30
			_lhs19.F2 = _rhs31
			_194_v45 = _rhs32
			_176_v33 = _rhs33
		}
	}
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
	_ = _index29
	(_134_v0).ArraySet1(_139_v2, (_index29).Int())
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
	_ = _index30
	(_134_v0).ArraySet1(_139_v2, (_index30).Int())
	var _196_v47 _dafny.Array
	_ = _196_v47
	var _nw35 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(2))
	_ = _nw35
	_196_v47 = _nw35
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_196_v47), 0))); ; {
		_guard_loop_1, _ok12 := _iter12()
		if !_ok12 {
			break
		}
		var _197_i5 _dafny.Int
		_197_i5 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_197_i5).Sign() != -1) && ((_197_i5).Cmp(_dafny.ArrayLen((_196_v47), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_196_v47, (_197_i5).Int(), (func() D1 {
				if _139_v2 {
					return Companion_D1_.Create_DC4_(Companion_D1_.Create_DC1_((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)))
				}
				return Companion_D1_.Create_DC4_(Companion_D1_.Create_DC1_(_139_v2))
			})()))
		}
	}
	for _iter13 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok13 := _iter13()
		if !_ok13 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	if !((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)) {
		var _198_v48 _dafny.Int
		_ = _198_v48
		_198_v48 = _dafny.IntOfInt64(890)
		var _199_v49 _dafny.Map
		_ = _199_v49
		_199_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v48, Companion_Default___.Fm2(_137_globalState))
		var _200_v50 _dafny.Sequence
		_ = _200_v50
		var _201_v51 bool
		_ = _201_v51
		var _202_v52 bool
		_ = _202_v52
		var _203_v53 _dafny.Set
		_ = _203_v53
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 bool
		_ = _out2
		var _out3 _dafny.Set
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M0((_199_v49).Cardinality(), _137_globalState)
		_200_v50 = _out0
		_201_v51 = _out1
		_202_v52 = _out2
		_203_v53 = _out3
		(_137_globalState).F2 = _198_v48
		var _204_v54 _dafny.Array
		_ = _204_v54
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw36
		_204_v54 = _nw36
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_204_v54), 0))
		_ = _index31
		(_204_v54).ArraySet1(_198_v48, (_index31).Int())
		var _205_v55 _dafny.Set
		_ = _205_v55
		_205_v55 = _dafny.SetOf(_198_v48, _198_v48)
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_204_v54), 0))
		_ = _index32
		(_204_v54).ArraySet1(((_205_v55).Cardinality()).Plus(_198_v48), (_index32).Int())
		_201_v51 = ((_dafny.SetOf(_201_v51, true)).Union(_203_v53)).IsProperSubsetOf((_203_v53).Difference(_203_v53))
		var _206_v56 *C0
		_ = _206_v56
		var _nw37 *C0 = New_C0_()
		_ = _nw37
		_nw37.Ctor__()
		_206_v56 = _nw37
	} else {
		var _207_v57 _dafny.Int
		_ = _207_v57
		_207_v57 = _dafny.IntOfInt64(970)
		var _208_v58 _dafny.MultiSet
		_ = _208_v58
		_208_v58 = _dafny.MultiSetOf(_207_v57, _207_v57)
		_208_v58 = (Companion_Default___.Fm4(_137_globalState)).Intersection(_208_v58)
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _index33
		(_134_v0).ArraySet1((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool), (_index33).Int())
		var _209_v59 _dafny.Sequence
		_ = _209_v59
		var _210_v60 bool
		_ = _210_v60
		var _211_v61 bool
		_ = _211_v61
		var _212_v62 _dafny.Set
		_ = _212_v62
		var _out4 _dafny.Sequence
		_ = _out4
		var _out5 bool
		_ = _out5
		var _out6 bool
		_ = _out6
		var _out7 _dafny.Set
		_ = _out7
		_out4, _out5, _out6, _out7 = Companion_Default___.M0(_207_v57, _137_globalState)
		_209_v59 = _out4
		_210_v60 = _out5
		_211_v61 = _out6
		_212_v62 = _out7
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _index34
		(_134_v0).ArraySet1(!(_139_v2), (_index34).Int())
		var _213_v63 _dafny.Map
		_ = _213_v63
		_213_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v57, false)
		var _214_v64 _dafny.Sequence
		_ = _214_v64
		_214_v64 = _dafny.UnicodeSeqOfUtf8Bytes("ljigfrq")
		var _215_v65 _dafny.Map
		_ = _215_v65
		_215_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v57, _214_v64)
		if (_207_v57).Cmp((((_213_v63).Update(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_215_v65).Contains(_207_v57) {
				return (_215_v65).Get(_207_v57).(_dafny.Sequence)
			}
			return _214_v64
		})()).Cardinality()), (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))).Merge(_213_v63)).Cardinality()) >= 0 {
			_210_v60 = _139_v2
			(_137_globalState).F2 = _dafny.IntOfInt64(387)
			var _216_v66 _dafny.Array
			_ = _216_v66
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw38
			_216_v66 = _nw38
			_216_v66 = _216_v66
			var _217_v67 _dafny.Array
			_ = _217_v67
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw39
			_217_v67 = _nw39
			var _218_v68 _dafny.Array
			_ = _218_v68
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_6
			var _nw40 _dafny.Array
			_ = _nw40
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw40 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.CodePoint = func(_219_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw40 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw40).ArraySet1CodePoint(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw40).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_218_v68 = _nw40
			var _220_v69 _dafny.CodePoint
			_ = _220_v69
			_220_v69 = _dafny.CodePoint('j')
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_218_v68), 0))
			_ = _index35
			(_218_v68).ArraySet1CodePoint(_220_v69, (_index35).Int())
			var _221_v70 _dafny.Array
			_ = _221_v70
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(8))
			_ = _nw41
			_221_v70 = _nw41
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_221_v70), 0))
			_ = _index36
			(_221_v70).ArraySet1(Companion_Default___.Fm5(_137_globalState), (_index36).Int())
			var _222_v71 *C0
			_ = _222_v71
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__()
			_222_v71 = _nw42
			var _223_v72 _dafny.Map
			_ = _223_v72
			_223_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v71, _207_v57)
			var _224_v73 _dafny.Sequence
			_ = _224_v73
			_224_v73 = _dafny.SeqOf((_207_v57).Plus((func() _dafny.Int {
				if (_223_v72).Contains(_222_v71) {
					return (_223_v72).Get(_222_v71).(_dafny.Int)
				}
				return _dafny.IntOfInt64(777)
			})()))
			var _225_v74 D1
			_ = _225_v74
			_225_v74 = Companion_D1_.Create_DC3_(_211_v61, _211_v61, (_dafny.SetOf(_210_v60)).Cardinality(), _139_v2)
			var _pat_let_tv12 = _207_v57
			_ = _pat_let_tv12
			var _pat_let_tv13 = _134_v0
			_ = _pat_let_tv13
			var _pat_let_tv14 = _134_v0
			_ = _pat_let_tv14
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
			_ = _index37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_218_v68), 0))
			_ = _index38
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_221_v70), 0))
			_ = _index39
			var _rhs34 bool = (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)
			_ = _rhs34
			var _rhs35 _dafny.Array = _217_v67
			_ = _rhs35
			var _rhs36 _dafny.CodePoint = _220_v69
			_ = _rhs36
			var _rhs37 D1 = func(_pat_let6_0 D1) D1 {
				return func(_226_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let7_0 bool) D1 {
						return func(_227_dt__update_hcf9_h0 bool) D1 {
							return func(_pat_let8_0 bool) D1 {
								return func(_228_dt__update_hcf12_h0 bool) D1 {
									return Companion_D1_.Create_DC3_(_227_dt__update_hcf9_h0, (_226_dt__update__tmp_h0).Dtor_cf10(), (_226_dt__update__tmp_h0).Dtor_cf11(), _228_dt__update_hcf12_h0)
								}(_pat_let8_0)
							}((_pat_let_tv14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_pat_let_tv13), 0))).Int()).(bool))
						}(_pat_let7_0)
					}((_pat_let_tv12).Cmp(_dafny.IntOfInt64(574)) > 0)
				}(_pat_let6_0)
			}(_225_v74)
			_ = _rhs37
			var _rhs38 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-651))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_229_v57 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_230_i7 _dafny.Int) _dafny.Int {
					return _229_v57
				}
			})(_207_v57)))
			_ = _rhs38
			var _lhs20 _dafny.Array = _134_v0
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
			_ = _lhs21
			var _lhs22 _dafny.Array = _218_v68
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_218_v68), 0))
			_ = _lhs23
			var _lhs24 _dafny.Array = _221_v70
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_221_v70), 0))
			_ = _lhs25
			(_lhs20).ArraySet1(_rhs34, (_lhs21).Int())
			_217_v67 = _rhs35
			(_lhs22).ArraySet1CodePoint(_rhs36, (_lhs23).Int())
			(_lhs24).ArraySet1(_rhs37, (_lhs25).Int())
			_224_v73 = _rhs38
			var _231_v75 D0
			_ = _231_v75
			_231_v75 = Companion_D0_.Create_DC0_(_207_v57, _dafny.MultiSetOf(_207_v57), _207_v57)
			var _232_v76 _dafny.Sequence
			_ = _232_v76
			_232_v76 = _dafny.SeqOf(_134_v0)
			var _233_v77 D1
			_ = _233_v77
			_233_v77 = Companion_D1_.Create_DC2_(Companion_Default___.SafeModuloInt(_207_v57, _207_v57), Companion_Default___.Fm1(_137_globalState), ((_231_v75).Dtor_cf0()).Times(_dafny.IntOfUint32((_232_v76).Cardinality())), _224_v73, (func() _dafny.CodePoint {
				if (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool) {
					return _220_v69
				}
				return (_218_v68).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_218_v68), 0))).Int())
			})())
			_233_v77 = _233_v77
		} else {
			var _234_v78 _dafny.Map
			_ = _234_v78
			_234_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(351), _207_v57)
			var _235_v79 _dafny.MultiSet
			_ = _235_v79
			_235_v79 = _dafny.MultiSetOf(!(_139_v2), _211_v61, (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
			var _236_v80 _dafny.Sequence
			_ = _236_v80
			_236_v80 = _dafny.SeqOf(_139_v2)
			var _237_v81 _dafny.Array
			_ = _237_v81
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw43
			_237_v81 = _nw43
			var _238_v82 _dafny.Set
			_ = _238_v82
			_238_v82 = _dafny.SetOf(_237_v81, _237_v81, _237_v81)
			var _239_v83 _dafny.Array
			_ = _239_v83
			var _nwElement0_11 _dafny.Int = _207_v57
			_ = _nwElement0_11
			var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(28))
			_ = _nw44
			(_nw44).ArraySet1(_nwElement0_11, 0)
			(_nw44).ArraySet1((_dafny.Zero).Minus(_207_v57), 1)
			(_nw44).ArraySet1((_207_v57).Times((_208_v58).Cardinality()), 2)
			(_nw44).ArraySet1(_207_v57, 3)
			(_nw44).ArraySet1(_207_v57, 4)
			(_nw44).ArraySet1((_dafny.IntOfUint32((_214_v64).Cardinality())).Plus(_dafny.IntOfInt64(352)), 5)
			(_nw44).ArraySet1(_207_v57, 6)
			(_nw44).ArraySet1(_207_v57, 7)
			(_nw44).ArraySet1(_207_v57, 8)
			(_nw44).ArraySet1(_207_v57, 9)
			(_nw44).ArraySet1(_207_v57, 10)
			(_nw44).ArraySet1(_207_v57, 11)
			(_nw44).ArraySet1(_207_v57, 12)
			(_nw44).ArraySet1(_207_v57, 13)
			(_nw44).ArraySet1(_207_v57, 14)
			(_nw44).ArraySet1((_207_v57).Plus((func() _dafny.Int {
				if (_208_v58).Contains(_207_v57) {
					return (_208_v58).Multiplicity(_207_v57)
				}
				return _207_v57
			})()), 15)
			(_nw44).ArraySet1(Companion_Default___.SafeDivisionInt(_207_v57, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-382))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_240_i8 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(53)
			}))).Cardinality())), 16)
			(_nw44).ArraySet1(_207_v57, 17)
			(_nw44).ArraySet1(_207_v57, 18)
			(_nw44).ArraySet1((_207_v57).Minus(_207_v57), 19)
			(_nw44).ArraySet1(Companion_Default___.SafeDivisionInt((_234_v78).Cardinality(), _207_v57), 20)
			(_nw44).ArraySet1(_dafny.IntOfInt64(94), 21)
			(_nw44).ArraySet1(_207_v57, 22)
			(_nw44).ArraySet1((func() _dafny.Int {
				if (_235_v79).Contains(_211_v61) {
					return (_235_v79).Multiplicity(_211_v61)
				}
				return _207_v57
			})(), 23)
			(_nw44).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if _210_v60 {
					return _236_v80
				}
				return _236_v80
			})()).Cardinality()), 24)
			(_nw44).ArraySet1((_dafny.IntOfInt64(599)).Minus(_207_v57), 25)
			(_nw44).ArraySet1((func() _dafny.Int {
				if (_235_v79).Contains((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)) {
					return (_235_v79).Multiplicity((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
				}
				return _207_v57
			})(), 26)
			(_nw44).ArraySet1((_238_v82).Cardinality(), 27)
			_239_v83 = _nw44
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_237_v81), 0))
			_ = _index40
			(_237_v81).ArraySet1(_207_v57, (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
			_ = _index41
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_237_v81), 0))
			_ = _index42
			var _rhs39 bool = Companion_Default___.Fm2(_137_globalState)
			_ = _rhs39
			var _rhs40 bool = true
			_ = _rhs40
			var _rhs41 _dafny.Int = _dafny.IntOfInt64(821)
			_ = _rhs41
			var _rhs42 bool = _210_v60
			_ = _rhs42
			var _lhs26 _dafny.Array = _134_v0
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
			_ = _lhs27
			var _lhs28 _dafny.Array = _237_v81
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_237_v81), 0))
			_ = _lhs29
			_211_v61 = _rhs39
			(_lhs26).ArraySet1(_rhs40, (_lhs27).Int())
			(_lhs28).ArraySet1(_rhs41, (_lhs29).Int())
			_211_v61 = _rhs42
			(_137_globalState).F2 = (_dafny.Zero).Minus((_237_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_237_v81), 0))).Int()).(_dafny.Int))
			var _241_v84 D1
			_ = _241_v84
			_241_v84 = Companion_D1_.Create_DC3_(!(!((_236_v80).Select((Companion_Default___.SafeIndex(_207_v57, _dafny.IntOfUint32((_236_v80).Cardinality()))).Uint32()).(bool))), _211_v61, (_237_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_237_v81), 0))).Int()).(_dafny.Int), _211_v61)
			_211_v61 = (_241_v84).Dtor_cf12()
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_237_v81), 0))
			_ = _index43
			(_237_v81).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if _210_v60 {
					return _207_v57
				}
				return (_237_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_237_v81), 0))).Int()).(_dafny.Int)
			})()), (_index43).Int())
			_207_v57 = _dafny.IntOfInt64(-825)
		}
	}
	var _242_v85 _dafny.Int
	_ = _242_v85
	_242_v85 = _dafny.IntOfInt64(-845)
	var _hi0 _dafny.Int = Companion_Default___.SafeDivisionInt(_242_v85, Companion_Default___.Fm6(_242_v85, _137_globalState))
	_ = _hi0
	for _243_i9 := _242_v85; _243_i9.Cmp(_hi0) < 0; _243_i9 = _243_i9.Plus(_dafny.One) {
		var _244_v86 _dafny.Sequence
		_ = _244_v86
		_244_v86 = _dafny.UnicodeSeqOfUtf8Bytes("fjukyxth")
		var _245_v87 _dafny.Map
		_ = _245_v87
		_245_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_i9, _244_v86)
		_245_v87 = (_245_v87).Update((func() _dafny.Int {
			if (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool) {
				return (_dafny.SetOf(_242_v85, _242_v85, _dafny.IntOfUint32((Companion_Default___.Fm7(_243_i9, _137_globalState)).Cardinality()))).Cardinality()
			}
			return _242_v85
		})(), _244_v86)
		var _246_v88 *C0
		_ = _246_v88
		var _nw45 *C0 = New_C0_()
		_ = _nw45
		_nw45.Ctor__()
		_246_v88 = _nw45
		var _247_v89 D1
		_ = _247_v89
		_247_v89 = Companion_D1_.Create_DC1_(_139_v2)
		var _248_v90 _dafny.Sequence
		_ = _248_v90
		_248_v90 = _dafny.SeqOf((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool), (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
		var _pat_let_tv15 = _139_v2
		_ = _pat_let_tv15
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _index44
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _index45
		var _rhs43 bool = false
		_ = _rhs43
		var _rhs44 bool = (_248_v90).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if _139_v2 {
				return _242_v85
			}
			return _242_v85
		})(), _dafny.IntOfUint32((_248_v90).Cardinality()))).Uint32()).(bool)
		_ = _rhs44
		var _rhs45 D1 = func(_pat_let9_0 D1) D1 {
			return func(_249_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let10_0 bool) D1 {
					return func(_250_dt__update_hcf3_h0 bool) D1 {
						return Companion_D1_.Create_DC1_(_250_dt__update_hcf3_h0)
					}(_pat_let10_0)
				}(_pat_let_tv15)
			}(_pat_let9_0)
		}(Companion_D1_.Create_DC1_((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)))
		_ = _rhs45
		var _rhs46 bool = (_139_v2) == (((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)) && (_139_v2))
		_ = _rhs46
		var _rhs47 *C0 = _246_v88
		_ = _rhs47
		var _lhs30 _dafny.Array = _134_v0
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _lhs31
		var _lhs32 _dafny.Array = _134_v0
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _lhs33
		_139_v2 = _rhs43
		(_lhs30).ArraySet1(_rhs44, (_lhs31).Int())
		_247_v89 = _rhs45
		(_lhs32).ArraySet1(_rhs46, (_lhs33).Int())
		_246_v88 = _rhs47
		var _pat_let_tv16 = _134_v0
		_ = _pat_let_tv16
		var _pat_let_tv17 = _134_v0
		_ = _pat_let_tv17
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _index46
		(_134_v0).ArraySet1((func(_pat_let11_0 D1) D1 {
			return func(_251_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let12_0 bool) D1 {
					return func(_252_dt__update_hcf3_h1 bool) D1 {
						return Companion_D1_.Create_DC1_(_252_dt__update_hcf3_h1)
					}(_pat_let12_0)
				}((_pat_let_tv17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_pat_let_tv16), 0))).Int()).(bool))
			}(_pat_let11_0)
		}(_247_v89)).Dtor_cf3(), (_index46).Int())
	}
	var _253_v91 _dafny.Sequence
	_ = _253_v91
	_253_v91 = _dafny.UnicodeSeqOfUtf8Bytes("dbygq")
	var _254_v92 _dafny.Sequence
	_ = _254_v92
	_254_v92 = _dafny.SeqOf(_242_v85, _242_v85, _242_v85, _dafny.IntOfUint32((_253_v91).Cardinality()))
	var _255_v93 _dafny.Sequence
	_ = _255_v93
	var _256_v94 bool
	_ = _256_v94
	var _257_v95 bool
	_ = _257_v95
	var _258_v96 _dafny.Set
	_ = _258_v96
	var _out8 _dafny.Sequence
	_ = _out8
	var _out9 bool
	_ = _out9
	var _out10 bool
	_ = _out10
	var _out11 _dafny.Set
	_ = _out11
	_out8, _out9, _out10, _out11 = Companion_Default___.M0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_254_v92, _254_v92)).Cardinality()), _137_globalState)
	_255_v93 = _out8
	_256_v94 = _out9
	_257_v95 = _out10
	_258_v96 = _out11
	var _259_v97 D3
	_ = _259_v97
	_259_v97 = Companion_D3_.Create_DC8_(_134_v0)
	var _260_v98 _dafny.Map
	_ = _260_v98
	_260_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v85, (_259_v97).Dtor_cf16())
	_260_v98 = (_260_v98).Update(_242_v85, _134_v0)
	var _261_v99 *C0
	_ = _261_v99
	var _nw46 *C0 = New_C0_()
	_ = _nw46
	_nw46.Ctor__()
	_261_v99 = _nw46
	var _262_v100 _dafny.Map
	_ = _262_v100
	_262_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool), _261_v99)
	_261_v99 = (func() *C0 {
		if (_262_v100).Contains(_139_v2) {
			return (_262_v100).Get(_139_v2).(*C0)
		}
		return _261_v99
	})()
	var _263_v101 _dafny.MultiSet
	_ = _263_v101
	_263_v101 = _dafny.MultiSetOf(_242_v85)
	var _264_v102 _dafny.Sequence
	_ = _264_v102
	var _265_v103 bool
	_ = _265_v103
	var _266_v104 bool
	_ = _266_v104
	var _267_v105 _dafny.Set
	_ = _267_v105
	var _out12 _dafny.Sequence
	_ = _out12
	var _out13 bool
	_ = _out13
	var _out14 bool
	_ = _out14
	var _out15 _dafny.Set
	_ = _out15
	_out12, _out13, _out14, _out15 = Companion_Default___.M0(((_263_v101).Difference(_263_v101)).Cardinality(), _137_globalState)
	_264_v102 = _out12
	_265_v103 = _out13
	_266_v104 = _out14
	_267_v105 = _out15
	_254_v92 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_242_v85), _254_v92), _254_v92)
	for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_134_v0), 0))); ; {
		_guard_loop_2, _ok14 := _iter14()
		if !_ok14 {
			break
		}
		var _268_i10 _dafny.Int
		_268_i10 = interface{}(_guard_loop_2).(_dafny.Int)
		if (true) && (((_268_i10).Sign() != -1) && ((_268_i10).Cmp(_dafny.ArrayLen((_134_v0), 0)) < 0)) {
			(_134_v0).ArraySet1(_257_v95, (_268_i10).Int())
		}
	}
	var _ingredients1 = _dafny.NewBuilder()
	_ = _ingredients1
	for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_134_v0), 0))); ; {
		_guard_loop_3, _ok15 := _iter15()
		if !_ok15 {
			break
		}
		var _269_i11 _dafny.Int
		_269_i11 = interface{}(_guard_loop_3).(_dafny.Int)
		if (true) && (((_269_i11).Sign() != -1) && ((_269_i11).Cmp(_dafny.ArrayLen((_134_v0), 0)) < 0)) {
			_ingredients1.Add(_dafny.TupleOf(_134_v0, (_269_i11).Int(), (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)))
		}
	}
	for _iter16 := _dafny.Iterate(_ingredients1); ; {
		_tup1, _ok16 := _iter16()
		if !_ok16 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup1.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup1.(_dafny.Tuple)).IndexInt(2)), (*(_tup1.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	var _hi1 _dafny.Int = _242_v85
	_ = _hi1
	for _270_i12 := _242_v85; _270_i12.Cmp(_hi1) < 0; _270_i12 = _270_i12.Plus(_dafny.One) {
		var _271_v106 _dafny.CodePoint
		_ = _271_v106
		_271_v106 = _dafny.CodePoint('n')
		var _272_v107 _dafny.Sequence
		_ = _272_v107
		_272_v107 = _dafny.SeqOf(_254_v92, _254_v92)
		var _273_v108 D1
		_ = _273_v108
		_273_v108 = Companion_D1_.Create_DC2_(_242_v85, _271_v106, _270_i12, _dafny.Companion_Sequence_.Concatenate((_272_v107).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_253_v91).Cardinality()), _dafny.IntOfUint32((_272_v107).Cardinality()))).Uint32()).(_dafny.Sequence), _254_v92), _271_v106)
		var _274_v109 D3
		_ = _274_v109
		_274_v109 = Companion_D3_.Create_DC10_(Companion_Default___.Fm2(_137_globalState), _253_v91, _265_v103)
		var _275_v110 _dafny.Sequence
		_ = _275_v110
		_275_v110 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("l"), _253_v91), _dafny.Companion_Sequence_.Concatenate(_253_v91, (_274_v109).Dtor_cf19()))
		var _276_v111 _dafny.Sequence
		_ = _276_v111
		_276_v111 = _dafny.SeqOf(_256_v94)
		var _277_v112 _dafny.Sequence
		_ = _277_v112
		_277_v112 = _dafny.SeqOf(_276_v111)
		var _rhs48 _dafny.Sequence = (_275_v110).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_263_v101).Contains(_dafny.IntOfUint32(((_277_v112).Select((Companion_Default___.SafeIndex(_242_v85, _dafny.IntOfUint32((_277_v112).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())) {
				return (_263_v101).Multiplicity(_dafny.IntOfUint32(((_277_v112).Select((Companion_Default___.SafeIndex(_242_v85, _dafny.IntOfUint32((_277_v112).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			}
			return _270_i12
		})(), _dafny.IntOfUint32((_275_v110).Cardinality()))).Uint32()).(_dafny.Sequence)
		_ = _rhs48
		var _rhs49 D1 = _273_v108
		_ = _rhs49
		var _rhs50 bool = Companion_Default___.Fm2(_137_globalState)
		_ = _rhs50
		var _rhs51 _dafny.Int = _270_i12
		_ = _rhs51
		var _lhs34 *GlobalState = _137_globalState
		_ = _lhs34
		_lhs34.F11 = _rhs48
		_273_v108 = _rhs49
		_265_v103 = _rhs50
		_242_v85 = _rhs51
		var _278_v113 _dafny.Map
		_ = _278_v113
		_278_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm6(_dafny.IntOfInt64(54), _137_globalState), (_242_v85).Cmp(_dafny.IntOfInt64(148)) < 0)
		var _279_v114 _dafny.Map
		_ = _279_v114
		_279_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_i12, (_dafny.Zero).Minus(_242_v85))
		var _280_v115 _dafny.Map
		_ = _280_v115
		_280_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
		var _281_v116 _dafny.Map
		_ = _281_v116
		_281_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, (_dafny.Zero).Minus((_280_v115).Cardinality()))
		var _282_v117 _dafny.Array
		_ = _282_v117
		var _nwElement0_12 _dafny.Int = _242_v85
		_ = _nwElement0_12
		var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(17))
		_ = _nw47
		(_nw47).ArraySet1(_nwElement0_12, 0)
		(_nw47).ArraySet1(_242_v85, 1)
		(_nw47).ArraySet1(_242_v85, 2)
		(_nw47).ArraySet1(((_279_v114).Merge(_279_v114)).Cardinality(), 3)
		(_nw47).ArraySet1(_270_i12, 4)
		(_nw47).ArraySet1(_270_i12, 5)
		(_nw47).ArraySet1(_270_i12, 6)
		(_nw47).ArraySet1(_242_v85, 7)
		(_nw47).ArraySet1((func() _dafny.Int {
			if (_281_v116).Contains(_134_v0) {
				return (_281_v116).Get(_134_v0).(_dafny.Int)
			}
			return _270_i12
		})(), 8)
		(_nw47).ArraySet1(Companion_Default___.SafeDivisionInt(_270_i12, _242_v85), 9)
		(_nw47).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(878), (_dafny.Zero).Minus(_242_v85)), 10)
		(_nw47).ArraySet1(_270_i12, 11)
		(_nw47).ArraySet1(_dafny.IntOfUint32((_253_v91).Cardinality()), 12)
		(_nw47).ArraySet1((_242_v85).Times(_dafny.IntOfUint32(((_274_v109).Dtor_cf19()).Cardinality())), 13)
		(_nw47).ArraySet1(_242_v85, 14)
		(_nw47).ArraySet1(_242_v85, 15)
		(_nw47).ArraySet1(_270_i12, 16)
		_282_v117 = _nw47
		var _283_v118 _dafny.Map
		_ = _283_v118
		_283_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm8(_242_v85, _266_v104, _139_v2, _137_globalState), _278_v113)
		var _284_v120 _dafny.Sequence
		_ = _284_v120
		_284_v120 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_i12, _266_v104), (func() _dafny.Map {
			if (_283_v118).Contains(_dafny.SetOf((_dafny.Zero).Minus(_242_v85))) {
				return (_283_v118).Get(_dafny.SetOf((_dafny.Zero).Minus(_242_v85))).(_dafny.Map)
			}
			return _278_v113
		})(), _278_v113, func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter17 := _dafny.Iterate((_263_v101).Elements()); ; {
				_compr_11, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _285_v119 _dafny.Int
				_285_v119 = interface{}(_compr_11).(_dafny.Int)
				if (_263_v101).Contains(_285_v119) {
					_coll11.Add((_285_v119).Times(_242_v85), true)
				}
			}
			return _coll11.ToMap()
		}(), (_278_v113).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-579), (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))))
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _index47
		var _rhs52 bool = (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)
		_ = _rhs52
		var _rhs53 _dafny.Map = (_284_v120).Select((Companion_Default___.SafeIndex(_270_i12, _dafny.IntOfUint32((_284_v120).Cardinality()))).Uint32()).(_dafny.Map)
		_ = _rhs53
		var _rhs54 _dafny.Array = _282_v117
		_ = _rhs54
		var _lhs35 _dafny.Array = _134_v0
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _lhs36
		(_lhs35).ArraySet1(_rhs52, (_lhs36).Int())
		_278_v113 = _rhs53
		_282_v117 = _rhs54
		var _286_v121 _dafny.Array
		_ = _286_v121
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
		_ = _nw48
		_286_v121 = _nw48
		var _287_v122 _dafny.Set
		_ = _287_v122
		_287_v122 = _258_v96
		var _288_v123 _dafny.Map
		_ = _288_v123
		_288_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)), _267_v105)
		var _289_v124 _dafny.Array
		_ = _289_v124
		var _nwElement0_13 _dafny.Set = _258_v96
		_ = _nwElement0_13
		var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(26))
		_ = _nw49
		(_nw49).ArraySet1(_nwElement0_13, 0)
		(_nw49).ArraySet1(_258_v96, 1)
		(_nw49).ArraySet1(_267_v105, 2)
		(_nw49).ArraySet1(_267_v105, 3)
		(_nw49).ArraySet1(_258_v96, 4)
		(_nw49).ArraySet1(_dafny.SetOf(false, false), 5)
		(_nw49).ArraySet1(_258_v96, 6)
		(_nw49).ArraySet1(_258_v96, 7)
		(_nw49).ArraySet1(_258_v96, 8)
		(_nw49).ArraySet1(_258_v96, 9)
		(_nw49).ArraySet1(_258_v96, 10)
		(_nw49).ArraySet1(_267_v105, 11)
		(_nw49).ArraySet1((_287_v122), 12)
		(_nw49).ArraySet1(_267_v105, 13)
		(_nw49).ArraySet1(_267_v105, 14)
		(_nw49).ArraySet1(_258_v96, 15)
		(_nw49).ArraySet1(_258_v96, 16)
		(_nw49).ArraySet1(_258_v96, 17)
		(_nw49).ArraySet1(_267_v105, 18)
		(_nw49).ArraySet1(_258_v96, 19)
		(_nw49).ArraySet1((func() _dafny.Set {
			if (_288_v123).Contains(true) {
				return (_288_v123).Get(true).(_dafny.Set)
			}
			return _258_v96
		})(), 20)
		(_nw49).ArraySet1(_dafny.SetOf(_257_v95, _256_v94), 21)
		(_nw49).ArraySet1(_267_v105, 22)
		(_nw49).ArraySet1(_267_v105, 23)
		(_nw49).ArraySet1(_267_v105, 24)
		(_nw49).ArraySet1(_267_v105, 25)
		_289_v124 = _nw49
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_286_v121), 0))
		_ = _index48
		(_286_v121).ArraySet1(_289_v124, (_index48).Int())
		var _290_v125 _dafny.Map
		_ = _290_v125
		_290_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(683))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_291_v106 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_292_i13 _dafny.Int) _dafny.CodePoint {
				return _291_v106
			}
		})(_271_v106))), _289_v124)
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_286_v121), 0))
		_ = _index49
		(_286_v121).ArraySet1((func() _dafny.Array {
			if (_290_v125).Contains(_253_v91) {
				return (_290_v125).Get(_253_v91).(_dafny.Array)
			}
			return _289_v124
		})(), (_index49).Int())
		(_137_globalState).F2 = _270_i12
	}
	if (_dafny.IntOfInt64(745)).Cmp((_dafny.Zero).Minus(_242_v85)) >= 0 {
		var _293_v126 _dafny.Array
		_ = _293_v126
		var _nwElement0_14 _dafny.Sequence = _254_v92
		_ = _nwElement0_14
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(4))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_14, 0)
		(_nw50).ArraySet1(_254_v92, 1)
		(_nw50).ArraySet1(_254_v92, 2)
		(_nw50).ArraySet1((func() _dafny.Sequence {
			if Companion_Default___.Fm2(_137_globalState) {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}((func(_294_v85 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_295_i14 _dafny.Int) _dafny.Int {
						return _294_v85
					}
				})(_242_v85)))
			}
			return _254_v92
		})(), 3)
		_293_v126 = _nw50
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_293_v126), 0))
		_ = _index50
		(_293_v126).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(688))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_296_v85 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_297_i15 _dafny.Int) _dafny.Int {
				return _296_v85
			}
		})(_242_v85))), (_index50).Int())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_293_v126), 0))
		_ = _index51
		(_293_v126).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus(_242_v85), (_254_v92).Select((Companion_Default___.SafeIndex(_242_v85, _dafny.IntOfUint32((_254_v92).Cardinality()))).Uint32()).(_dafny.Int)), (_index51).Int())
		var _298_v127 _dafny.CodePoint
		_ = _298_v127
		_298_v127 = _dafny.CodePoint('u')
		var _299_v128 _dafny.Set
		_ = _299_v128
		_299_v128 = _dafny.SetOf(_298_v127, _298_v127)
		var _300_v129 _dafny.Map
		_ = _300_v129
		_300_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v85, _299_v128)
		var _301_v130 D5
		_ = _301_v130
		_301_v130 = Companion_D5_.Create_DC13_(_299_v128)
		_265_v103 = ((func() _dafny.Set {
			if (_300_v129).Contains(_242_v85) {
				return (_300_v129).Get(_242_v85).(_dafny.Set)
			}
			return (_301_v130).Dtor_cf23()
		})()).IsDisjointFrom((_299_v128).Intersection(_299_v128))
		var _302_v131 _dafny.Sequence
		_ = _302_v131
		_302_v131 = _dafny.SeqOf(_139_v2)
		var _303_v132 _dafny.Array
		_ = _303_v132
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_7
		var _nw51 _dafny.Array
		_ = _nw51
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw51 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Set = (func(_304_v96 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_305_i16 _dafny.Int) _dafny.Set {
					return _304_v96
				}
			})(_258_v96)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw51 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw51).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw51).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_303_v132 = _nw51
		var _306_v133 _dafny.Set
		_ = _306_v133
		_306_v133 = _258_v96
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_303_v132), 0))
		_ = _index52
		(_303_v132).ArraySet1(_306_v133, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _index53
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_303_v132), 0))
		_ = _index54
		var _rhs55 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_302_v131, _302_v131)
		_ = _rhs55
		var _rhs56 bool = _139_v2
		_ = _rhs56
		var _rhs57 _dafny.Set = _267_v105
		_ = _rhs57
		var _rhs58 _dafny.Int = _dafny.IntOfUint32((_253_v91).Cardinality())
		_ = _rhs58
		var _lhs37 _dafny.Array = _134_v0
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
		_ = _lhs38
		var _lhs39 _dafny.Array = _303_v132
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_303_v132), 0))
		_ = _lhs40
		var _lhs41 *GlobalState = _137_globalState
		_ = _lhs41
		_302_v131 = _rhs55
		(_lhs37).ArraySet1(_rhs56, (_lhs38).Int())
		(_lhs39).ArraySet1(_rhs57, (_lhs40).Int())
		_lhs41.F2 = _rhs58
		_257_v95 = _266_v104
		(_137_globalState).F2 = _242_v85
	} else {
		if !(!(_139_v2)) {
			var _307_v134 _dafny.MultiSet
			_ = _307_v134
			_307_v134 = _dafny.MultiSetOf(_253_v91)
			var _308_v135 _dafny.Sequence
			_ = _308_v135
			var _309_v136 bool
			_ = _309_v136
			var _310_v137 bool
			_ = _310_v137
			var _311_v138 _dafny.Set
			_ = _311_v138
			var _out16 _dafny.Sequence
			_ = _out16
			var _out17 bool
			_ = _out17
			var _out18 bool
			_ = _out18
			var _out19 _dafny.Set
			_ = _out19
			_out16, _out17, _out18, _out19 = Companion_Default___.M0((_307_v134).Cardinality(), _137_globalState)
			_308_v135 = _out16
			_309_v136 = _out17
			_310_v137 = _out18
			_311_v138 = _out19
			var _312_v139 _dafny.Array
			_ = _312_v139
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_8
			var _nw52 _dafny.Array
			_ = _nw52
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw52 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = (func(_313_v85 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_314_i17 _dafny.Int) _dafny.Int {
						return (_314_i17).Plus((_dafny.SetOf(_dafny.IntOfInt64(948), _313_v85, _313_v85)).Cardinality())
					}
				})(_242_v85)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw52 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw52).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw52).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_312_v139 = _nw52
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_9
			var _nw53 _dafny.Array
			_ = _nw53
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw53 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = func(_315_i18 _dafny.Int) _dafny.Int {
					return (_315_i18).Times(_dafny.IntOfInt64(336))
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw53 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw53).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw53).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_312_v139 = _nw53
			var _316_v140 _dafny.CodePoint
			_ = _316_v140
			_316_v140 = _dafny.CodePoint('h')
			_316_v140 = _316_v140
			_139_v2 = _266_v104
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
			_ = _index55
			(_134_v0).ArraySet1(_256_v94, (_index55).Int())
		} else {
			var _317_v141 _dafny.Sequence
			_ = _317_v141
			var _318_v142 bool
			_ = _318_v142
			var _319_v143 bool
			_ = _319_v143
			var _320_v144 _dafny.Set
			_ = _320_v144
			var _out20 _dafny.Sequence
			_ = _out20
			var _out21 bool
			_ = _out21
			var _out22 bool
			_ = _out22
			var _out23 _dafny.Set
			_ = _out23
			_out20, _out21, _out22, _out23 = Companion_Default___.M0(_242_v85, _137_globalState)
			_317_v141 = _out20
			_318_v142 = _out21
			_319_v143 = _out22
			_320_v144 = _out23
			_261_v99 = _261_v99
			_261_v99 = (func() *C0 {
				if _266_v104 {
					return _261_v99
				}
				return _261_v99
			})()
			var _321_v145 _dafny.Map
			_ = _321_v145
			_321_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v85, (_dafny.IntOfInt64(-784)).Plus(_242_v85))
			_321_v145 = (_321_v145).Update(_242_v85, _242_v85)
			var _322_v146 _dafny.Set
			_ = _322_v146
			_322_v146 = _dafny.SetOf(_263_v101)
			var _323_v147 _dafny.Set
			_ = _323_v147
			_323_v147 = _dafny.SetOf(_261_v99, _261_v99)
			var _324_v148 _dafny.Map
			_ = _324_v148
			_324_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v104, _139_v2)
			var _325_v149 _dafny.Array
			_ = _325_v149
			var _nwElement0_15 _dafny.Int = _242_v85
			_ = _nwElement0_15
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(20))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_15, 0)
			(_nw54).ArraySet1(_242_v85, 1)
			(_nw54).ArraySet1(_242_v85, 2)
			(_nw54).ArraySet1(_242_v85, 3)
			(_nw54).ArraySet1(_242_v85, 4)
			(_nw54).ArraySet1(Companion_Default___.SafeDivisionInt(_242_v85, Companion_Default___.Fm6(_242_v85, _137_globalState)), 5)
			(_nw54).ArraySet1(_242_v85, 6)
			(_nw54).ArraySet1(_242_v85, 7)
			(_nw54).ArraySet1(Companion_Default___.Fm6(_242_v85, _137_globalState), 8)
			(_nw54).ArraySet1(_242_v85, 9)
			(_nw54).ArraySet1(((_dafny.Zero).Minus((_322_v146).Cardinality())).Plus(_242_v85), 10)
			(_nw54).ArraySet1(Companion_Default___.Fm6(_242_v85, _137_globalState), 11)
			(_nw54).ArraySet1(Companion_Default___.SafeModuloInt(_242_v85, Companion_Default___.Fm6(_dafny.IntOfUint32((_253_v91).Cardinality()), _137_globalState)), 12)
			(_nw54).ArraySet1((_323_v147).Cardinality(), 13)
			(_nw54).ArraySet1(_242_v85, 14)
			(_nw54).ArraySet1(_242_v85, 15)
			(_nw54).ArraySet1((_dafny.Zero).Minus(_242_v85), 16)
			(_nw54).ArraySet1(_dafny.IntOfInt64(-986), 17)
			(_nw54).ArraySet1((_324_v148).Cardinality(), 18)
			(_nw54).ArraySet1(Companion_Default___.SafeModuloInt(_242_v85, _242_v85), 19)
			_325_v149 = _nw54
			var _326_v150 _dafny.Sequence
			_ = _326_v150
			_326_v150 = _dafny.SeqOf(_325_v149, _325_v149, _325_v149, _325_v149, _325_v149)
			_325_v149 = (_326_v150).Select((Companion_Default___.SafeIndex((_242_v85).Times(_242_v85), _dafny.IntOfUint32((_326_v150).Cardinality()))).Uint32()).(_dafny.Array)
		}
		var _327_v151 _dafny.Map
		_ = _327_v151
		_327_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v85, _265_v103)
		var _328_v152 _dafny.MultiSet
		_ = _328_v152
		_328_v152 = _dafny.MultiSetOf(true, _266_v104)
		var _329_v153 _dafny.Map
		_ = _329_v153
		_329_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_261_v99, _328_v152)
		var _rhs59 bool = !((func() bool {
			if (_327_v151).Contains(Companion_Default___.SafeModuloInt(_242_v85, ((func() _dafny.MultiSet {
				if (_329_v153).Contains(_261_v99) {
					return (_329_v153).Get(_261_v99).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
			})()).Cardinality())) {
				return (_327_v151).Get(Companion_Default___.SafeModuloInt(_242_v85, ((func() _dafny.MultiSet {
					if (_329_v153).Contains(_261_v99) {
						return (_329_v153).Get(_261_v99).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
				})()).Cardinality())).(bool)
			}
			return _139_v2
		})())
		_ = _rhs59
		var _rhs60 _dafny.Int = Companion_Default___.SafeModuloInt(_242_v85, _242_v85)
		_ = _rhs60
		var _lhs42 *GlobalState = _137_globalState
		_ = _lhs42
		_266_v104 = _rhs59
		_lhs42.F2 = _rhs60
		_242_v85 = _242_v85
		(_137_globalState).F2 = _dafny.IntOfUint32((_253_v91).Cardinality())
		(_137_globalState).F2 = (func() _dafny.Int {
			if _256_v94 {
				return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_242_v85)))
			}
			return (func() _dafny.Int {
				if _266_v104 {
					return Companion_Default___.Fm6(_242_v85, _137_globalState)
				}
				return _242_v85
			})()
		})()
	}
	for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_134_v0), 0))); ; {
		_guard_loop_4, _ok18 := _iter18()
		if !_ok18 {
			break
		}
		var _330_i19 _dafny.Int
		_330_i19 = interface{}(_guard_loop_4).(_dafny.Int)
		if (true) && (((_330_i19).Sign() != -1) && ((_330_i19).Cmp(_dafny.ArrayLen((_134_v0), 0)) < 0)) {
			(_134_v0).ArraySet1(_139_v2, (_330_i19).Int())
		}
	}
	var _331_i20 _dafny.Int
	_ = _331_i20
	_331_i20 = _dafny.Zero
	{
		for _265_v103 {
			{
				if (_331_i20).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_331_i20 = (_331_i20).Plus(_dafny.One)
				var _332_v154 _dafny.Array
				_ = _332_v154
				var _nwElement0_16 _dafny.Sequence = _253_v91
				_ = _nwElement0_16
				var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(5))
				_ = _nw55
				(_nw55).ArraySet1(_nwElement0_16, 0)
				(_nw55).ArraySet1(_253_v91, 1)
				(_nw55).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cugllwi"), 2)
				(_nw55).ArraySet1(_253_v91, 3)
				(_nw55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_253_v91, _253_v91), 4)
				_332_v154 = _nw55
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_332_v154), 0))
				_ = _index56
				(_332_v154).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}(func(_333_i21 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), (_index56).Int())
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_332_v154), 0))
				_ = _index57
				(_332_v154).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-999))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}((func(_334_v104 bool) func(_dafny.Int) _dafny.CodePoint {
					return func(_335_i22 _dafny.Int) _dafny.CodePoint {
						return (func() _dafny.CodePoint {
							if _334_v104 {
								return _dafny.CodePoint('m')
							}
							return _dafny.CodePoint('u')
						})()
					}
				})(_266_v104))), (_index57).Int())
				if (_dafny.IntOfUint32((_254_v92).Cardinality())).Cmp(_dafny.IntOfInt64(911)) >= 0 {
					var _336_v155 _dafny.Sequence
					_ = _336_v155
					var _337_v156 bool
					_ = _337_v156
					var _338_v157 bool
					_ = _338_v157
					var _339_v158 _dafny.Set
					_ = _339_v158
					var _out24 _dafny.Sequence
					_ = _out24
					var _out25 bool
					_ = _out25
					var _out26 bool
					_ = _out26
					var _out27 _dafny.Set
					_ = _out27
					_out24, _out25, _out26, _out27 = Companion_Default___.M0(Companion_Default___.Fm6(_242_v85, _137_globalState), _137_globalState)
					_336_v155 = _out24
					_337_v156 = _out25
					_338_v157 = _out26
					_339_v158 = _out27
					var _340_v159 _dafny.Map
					_ = _340_v159
					_340_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Int {
						if _257_v95 {
							return _242_v85
						}
						return _242_v85
					})()), _337_v156)
					_340_v159 = (_340_v159).Update((_dafny.Zero).Minus((_242_v85).Minus(_242_v85)), !((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)) || (_257_v95))
					(_137_globalState).F2 = _242_v85
					var _341_v160 _dafny.MultiSet
					_ = _341_v160
					_341_v160 = _dafny.MultiSetOf((_332_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_332_v154), 0))).Int()).(_dafny.Sequence))
					var _342_v161 _dafny.CodePoint
					_ = _342_v161
					_342_v161 = _dafny.CodePoint('q')
					var _rhs61 _dafny.MultiSet = _dafny.MultiSetOf(_253_v91, _253_v91)
					_ = _rhs61
					var _rhs62 _dafny.CodePoint = _342_v161
					_ = _rhs62
					_341_v160 = _rhs61
					_342_v161 = _rhs62
					var _343_v162 _dafny.Array
					_ = _343_v162
					var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
					_ = _nw56
					_343_v162 = _nw56
					var _344_v163 _dafny.Set
					_ = _344_v163
					_344_v163 = _dafny.SetOf(_343_v162)
					var _345_v164 _dafny.Map
					_ = _345_v164
					_345_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_344_v163).Cardinality(), _242_v85)
					(_137_globalState).F2 = Companion_Default___.Fm6((_242_v85).Minus((_345_v164).Cardinality()), _137_globalState)
				} else {
					var _346_v165 _dafny.Sequence
					_ = _346_v165
					var _347_v166 bool
					_ = _347_v166
					var _348_v167 bool
					_ = _348_v167
					var _349_v168 _dafny.Set
					_ = _349_v168
					var _out28 _dafny.Sequence
					_ = _out28
					var _out29 bool
					_ = _out29
					var _out30 bool
					_ = _out30
					var _out31 _dafny.Set
					_ = _out31
					_out28, _out29, _out30, _out31 = Companion_Default___.M0(_242_v85, _137_globalState)
					_346_v165 = _out28
					_347_v166 = _out29
					_348_v167 = _out30
					_349_v168 = _out31
					_347_v166 = _266_v104
					(_137_globalState).F2 = _242_v85
					var _350_v169 _dafny.Array
					_ = _350_v169
					var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
					_ = _nw57
					_350_v169 = _nw57
					var _351_v170 _dafny.Map
					_ = _351_v170
					_351_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)), _350_v169)
					var _352_v171 D2
					_ = _352_v171
					_352_v171 = Companion_D2_.Create_DC5_(_dafny.SeqOf(_139_v2, _139_v2, (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool)))
					var _353_v172 D2
					_ = _353_v172
					_353_v172 = Companion_D2_.Create_DC7_(_352_v171)
					var _354_v173 _dafny.Sequence
					_ = _354_v173
					_354_v173 = _dafny.SeqOf(_353_v172, Companion_D2_.Create_DC7_(Companion_D2_.Create_DC6_()))
					var _355_v174 _dafny.MultiSet
					_ = _355_v174
					_355_v174 = _dafny.MultiSetOf(_354_v173)
					var _356_v175 _dafny.Map
					_ = _356_v175
					_356_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _242_v85)
					var _357_v176 _dafny.Map
					_ = _357_v176
					_357_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, _347_v166)
					var _358_v177 _dafny.Array
					_ = _358_v177
					var _nwElement0_17 _dafny.Int = (_254_v92).Select((Companion_Default___.SafeIndex((_351_v170).Cardinality(), _dafny.IntOfUint32((_254_v92).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _nwElement0_17
					var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(21))
					_ = _nw58
					(_nw58).ArraySet1(_nwElement0_17, 0)
					(_nw58).ArraySet1(_242_v85, 1)
					(_nw58).ArraySet1(_242_v85, 2)
					(_nw58).ArraySet1((_355_v174).Cardinality(), 3)
					(_nw58).ArraySet1((_356_v175).Cardinality(), 4)
					(_nw58).ArraySet1(_242_v85, 5)
					(_nw58).ArraySet1((_dafny.Zero).Minus(_242_v85), 6)
					(_nw58).ArraySet1(_242_v85, 7)
					(_nw58).ArraySet1(_242_v85, 8)
					(_nw58).ArraySet1(_242_v85, 9)
					(_nw58).ArraySet1(_242_v85, 10)
					(_nw58).ArraySet1(_242_v85, 11)
					(_nw58).ArraySet1(_dafny.IntOfInt64(-668), 12)
					(_nw58).ArraySet1(_242_v85, 13)
					(_nw58).ArraySet1(_242_v85, 14)
					(_nw58).ArraySet1(_242_v85, 15)
					(_nw58).ArraySet1(_242_v85, 16)
					(_nw58).ArraySet1(_242_v85, 17)
					(_nw58).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_348_v167, _dafny.IntOfInt64(288))).Cardinality(), 18)
					(_nw58).ArraySet1((_357_v176).Cardinality(), 19)
					(_nw58).ArraySet1(_242_v85, 20)
					_358_v177 = _nw58
					var _359_v178 _dafny.Map
					_ = _359_v178
					_359_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-755), (_350_v169))
					var _360_v179 _dafny.Array
					_ = _360_v179
					var _nwElement0_18 _dafny.Array = _358_v177
					_ = _nwElement0_18
					var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(15))
					_ = _nw59
					(_nw59).ArraySet1(_nwElement0_18, 0)
					(_nw59).ArraySet1(_358_v177, 1)
					(_nw59).ArraySet1((func() _dafny.Array {
						if (_359_v178).Contains((_dafny.Zero).Minus((_dafny.SetOf(_242_v85)).Cardinality())) {
							return (_359_v178).Get((_dafny.Zero).Minus((_dafny.SetOf(_242_v85)).Cardinality())).(_dafny.Array)
						}
						return _350_v169
					})(), 2)
					(_nw59).ArraySet1(_358_v177, 3)
					(_nw59).ArraySet1(_350_v169, 4)
					(_nw59).ArraySet1(_358_v177, 5)
					(_nw59).ArraySet1(_358_v177, 6)
					(_nw59).ArraySet1(_358_v177, 7)
					(_nw59).ArraySet1(_358_v177, 8)
					(_nw59).ArraySet1(_358_v177, 9)
					(_nw59).ArraySet1(_358_v177, 10)
					(_nw59).ArraySet1(_350_v169, 11)
					(_nw59).ArraySet1(_350_v169, 12)
					(_nw59).ArraySet1(_350_v169, 13)
					(_nw59).ArraySet1(_350_v169, 14)
					_360_v179 = _nw59
					var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_360_v179), 0))
					_ = _index58
					(_360_v179).ArraySet1(_350_v169, (_index58).Int())
					var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_360_v179), 0))
					_ = _index59
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _index60
					var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _index61
					var _rhs63 bool = _256_v94
					_ = _rhs63
					var _rhs64 _dafny.Array = _350_v169
					_ = _rhs64
					var _rhs65 bool = (func() bool {
						if (_242_v85).Cmp(_242_v85) > 0 {
							return _347_v166
						}
						return _257_v95
					})()
					_ = _rhs65
					var _rhs66 bool = _139_v2
					_ = _rhs66
					var _lhs43 _dafny.Array = _360_v179
					_ = _lhs43
					var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_360_v179), 0))
					_ = _lhs44
					var _lhs45 _dafny.Array = _134_v0
					_ = _lhs45
					var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _lhs46
					var _lhs47 _dafny.Array = _134_v0
					_ = _lhs47
					var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _lhs48
					_256_v94 = _rhs63
					(_lhs43).ArraySet1(_rhs64, (_lhs44).Int())
					(_lhs45).ArraySet1(_rhs65, (_lhs46).Int())
					(_lhs47).ArraySet1(_rhs66, (_lhs48).Int())
					var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _index62
					(_134_v0).ArraySet1(_265_v103, (_index62).Int())
				}
				if _265_v103 {
					var _361_v180 _dafny.MultiSet
					_ = _361_v180
					_361_v180 = _dafny.MultiSetOf(_256_v94, !(_265_v103), false, _266_v104)
					var _362_v181 _dafny.Sequence
					_ = _362_v181
					_362_v181 = _dafny.SeqOf(_361_v180, _361_v180, _361_v180, _361_v180)
					_362_v181 = _dafny.Companion_Sequence_.Concatenate(_362_v181, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(729))).Uint32(), func(coer25 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}((func(_363_v95 bool, _364_v104 bool) func(_dafny.Int) _dafny.MultiSet {
						return func(_365_i23 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf(_363_v95, _364_v104)
						}
					})(_257_v95, _266_v104))))
					_361_v180 = (_362_v181).Select((Companion_Default___.SafeIndex(_242_v85, _dafny.IntOfUint32((_362_v181).Cardinality()))).Uint32()).(_dafny.MultiSet)
					var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _index63
					(_134_v0).ArraySet1(!(_256_v94) || (!((_242_v85).Cmp((_dafny.Zero).Minus(_242_v85)) >= 0)), (_index63).Int())
					var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _index64
					(_134_v0).ArraySet1(false, (_index64).Int())
					_257_v95 = _266_v104
				} else {
					var _366_v182 _dafny.Sequence
					_ = _366_v182
					_366_v182 = _dafny.SeqOf(_267_v105, _258_v96, _267_v105)
					var _367_v183 _dafny.Set
					_ = _367_v183
					_367_v183 = Companion_Default___.Fm9(_263_v101, _266_v104, _139_v2, _137_globalState)
					var _368_v184 _dafny.CodePoint
					_ = _368_v184
					_368_v184 = _dafny.CodePoint('w')
					var _369_v185 _dafny.Map
					_ = _369_v185
					_369_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_v184, _dafny.SetOf(_265_v103, _256_v94, _265_v103, _139_v2, _257_v95))
					var _370_v186 _dafny.Map
					_ = _370_v186
					_370_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_v184, _368_v184)
					var _371_v187 _dafny.Array
					_ = _371_v187
					var _nwElement0_19 _dafny.Set = (_267_v105).Union(_258_v96)
					_ = _nwElement0_19
					var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(22))
					_ = _nw60
					(_nw60).ArraySet1(_nwElement0_19, 0)
					(_nw60).ArraySet1((_dafny.SetOf(false, _265_v103)).Union(_267_v105), 1)
					(_nw60).ArraySet1(Companion_Default___.Fm9(_dafny.MultiSetOf(_242_v85, _242_v85), true, _266_v104, _137_globalState), 2)
					(_nw60).ArraySet1(_dafny.SetOf(false), 3)
					(_nw60).ArraySet1((_366_v182).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm6(_242_v85, _137_globalState), _dafny.IntOfUint32((_366_v182).Cardinality()))).Uint32()).(_dafny.Set), 4)
					(_nw60).ArraySet1((_dafny.SetOf(_265_v103, _266_v104)).Intersection(_dafny.SetOf(true, _256_v94, _139_v2, _266_v104)), 5)
					(_nw60).ArraySet1(_267_v105, 6)
					(_nw60).ArraySet1((_367_v183), 7)
					(_nw60).ArraySet1((_258_v96).Intersection(_dafny.SetOf(_139_v2, _266_v104)), 8)
					(_nw60).ArraySet1(_258_v96, 9)
					(_nw60).ArraySet1(_258_v96, 10)
					(_nw60).ArraySet1(_258_v96, 11)
					(_nw60).ArraySet1((func() _dafny.Set {
						if (_369_v185).Contains((func() _dafny.CodePoint {
							if (_370_v186).Contains(_368_v184) {
								return (_370_v186).Get(_368_v184).(_dafny.CodePoint)
							}
							return _368_v184
						})()) {
							return (_369_v185).Get((func() _dafny.CodePoint {
								if (_370_v186).Contains(_368_v184) {
									return (_370_v186).Get(_368_v184).(_dafny.CodePoint)
								}
								return _368_v184
							})()).(_dafny.Set)
						}
						return _258_v96
					})(), 12)
					(_nw60).ArraySet1(_258_v96, 13)
					(_nw60).ArraySet1(_258_v96, 14)
					(_nw60).ArraySet1(_267_v105, 15)
					(_nw60).ArraySet1(_258_v96, 16)
					(_nw60).ArraySet1((func() _dafny.Set {
						if _256_v94 {
							return _258_v96
						}
						return (_367_v183)
					})(), 17)
					(_nw60).ArraySet1(Companion_Default___.Fm9(_263_v101, !(_266_v104), _256_v94, _137_globalState), 18)
					(_nw60).ArraySet1(_258_v96, 19)
					(_nw60).ArraySet1(_258_v96, 20)
					(_nw60).ArraySet1(_267_v105, 21)
					_371_v187 = _nw60
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_371_v187), 0))
					_ = _index65
					(_371_v187).ArraySet1(_267_v105, (_index65).Int())
					var _372_v188 _dafny.Map
					_ = _372_v188
					_372_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_v184, _263_v101)
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_371_v187), 0))
					_ = _index66
					(_371_v187).ArraySet1((func() _dafny.Set {
						if (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool) {
							return _258_v96
						}
						return Companion_Default___.Fm9((func() _dafny.MultiSet {
							if (_372_v188).Contains(_368_v184) {
								return (_372_v188).Get(_368_v184).(_dafny.MultiSet)
							}
							return _263_v101
						})(), (_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool), _257_v95, _137_globalState)
					})(), (_index66).Int())
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))
					_ = _index67
					(_134_v0).ArraySet1(false, (_index67).Int())
					var _373_v189 _dafny.Sequence
					_ = _373_v189
					_373_v189 = _dafny.SeqOf((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
					(_137_globalState).F2 = Companion_Default___.SafeDivisionInt((_242_v85).Minus(_dafny.IntOfInt64(254)), (_242_v85).Plus((func() _dafny.Int {
						if (_263_v101).Contains(_242_v85) {
							return (_263_v101).Multiplicity(_242_v85)
						}
						return _dafny.IntOfUint32((_373_v189).Cardinality())
					})()))
					_139_v2 = (_256_v94) || ((_134_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_134_v0), 0))).Int()).(bool))
					(_137_globalState).F2 = (_242_v85).Times(Companion_Default___.Fm6(_242_v85, _137_globalState))
				}
				var _374_v190 _dafny.Array
				_ = _374_v190
				var _nwElement0_20 _dafny.Int = _242_v85
				_ = _nwElement0_20
				var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(3))
				_ = _nw61
				(_nw61).ArraySet1(_nwElement0_20, 0)
				(_nw61).ArraySet1((_dafny.IntOfInt64(-198)).Times(_242_v85), 1)
				(_nw61).ArraySet1(Companion_Default___.Fm6(_dafny.IntOfUint32((_253_v91).Cardinality()), _137_globalState), 2)
				_374_v190 = _nw61
				var _375_v191 _dafny.Map
				_ = _375_v191
				_375_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v2, _242_v85)
				var _376_v192 _dafny.Map
				_ = _376_v192
				_376_v192 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_375_v191, _266_v104)
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_374_v190), 0))
				_ = _index68
				(_374_v190).ArraySet1((_242_v85).Times((_dafny.Zero).Minus((_376_v192).Cardinality())), (_index68).Int())
				var _377_v193 _dafny.Map
				_ = _377_v193
				_377_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(196), _242_v85)
				var _378_v194 _dafny.CodePoint
				_ = _378_v194
				_378_v194 = _dafny.CodePoint('t')
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_374_v190), 0))
				_ = _index69
				(_374_v190).ArraySet1(((Companion_Default___.Fm10(_377_v193, _265_v103, _266_v104, false, _137_globalState)).Dtor_cf4()).Minus(Companion_Default___.SafeDivisionInt(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v194, _261_v99)).Update(_378_v194, _261_v99)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dwaa")).Cardinality()))), (_index69).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_dafny.Print((_134_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v0).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v1).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_137_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(-34))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_137_globalState.F11, _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_139_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_196_v47).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf13()).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_196_v47).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf13()).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_242_v85)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_253_v91.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_254_v92, _dafny.SeqOf(_dafny.IntOfInt64(-845), _dafny.IntOfInt64(-845), _dafny.IntOfInt64(-845), _dafny.IntOfInt64(-845), _dafny.IntOfInt64(5), _dafny.IntOfInt64(-845), _dafny.IntOfInt64(-845), _dafny.IntOfInt64(-845), _dafny.IntOfInt64(5))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_255_v93, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(false, false, _dafny.IntOfInt64(8), false), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(false, false, _dafny.IntOfInt64(8), false), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_256_v94)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_257_v95)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v96).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v97).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v98).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v100).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_263_v101).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-845))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_264_v102, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(false, false, _dafny.Zero, false), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(false, false, _dafny.Zero, false), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_265_v103)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_266_v104)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v105).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_331_i20)
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
	Cf0 _dafny.Int
	Cf1 _dafny.MultiSet
	Cf2 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int, Cf1 _dafny.MultiSet, Cf2 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0, Cf1, Cf2}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero, _dafny.EmptyMultiSet, _dafny.Zero)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.MultiSet {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf2
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1.Equals(data2.Cf1) && data1.Cf2.Cmp(data2.Cf2) == 0
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
	Cf4 _dafny.Int
	Cf5 _dafny.CodePoint
	Cf6 _dafny.Int
	Cf7 _dafny.Sequence
	Cf8 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Int, Cf5 _dafny.CodePoint, Cf6 _dafny.Int, Cf7 _dafny.Sequence, Cf8 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf4, Cf5, Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf9  bool
	Cf10 bool
	Cf11 _dafny.Int
	Cf12 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf9 bool, Cf10 bool, Cf11 _dafny.Int, Cf12 bool) D1 {
	return D1{D1_DC3{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC1 struct {
	Cf3 bool
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf3 bool) D1 {
	return D1{D1_DC1{Cf3}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

type D1_DC4 struct {
	Cf13 D1
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf13 D1) D1 {
	return D1{D1_DC4{Cf13}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, _dafny.CodePoint('D'), _dafny.Zero, _dafny.EmptySeq, _dafny.CodePoint('D'))
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf5() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf7
}

func (_this D1) Dtor_cf8() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC3).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC3).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf11
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC3).Cf12
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC1).Cf3
}

func (_this D1) Dtor_cf13() D1 {
	return _this.Get_().(D1_DC4).Cf13
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7.Equals(data2.Cf7) && data1.Cf8 == data2.Cf8
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12 == data2.Cf12
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf3 == data2.Cf3
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_() D2 {
	return D2{D2_DC6{}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf14 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf14 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf14}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC7 struct {
	Cf15 D2
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf15 D2) D2 {
	return D2{D2_DC7{Cf15}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_()
}

func (_this D2) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf14
}

func (_this D2) Dtor_cf15() D2 {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf15) + ")"
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
			_, ok := other.Get_().(D2_DC6)
			return ok
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf14.Equals(data2.Cf14)
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf17 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf17 _dafny.Int) D3 {
	return D3{D3_DC9{Cf17}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf18 bool
	Cf19 _dafny.Sequence
	Cf20 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf18 bool, Cf19 _dafny.Sequence, Cf20 bool) D3 {
	return D3{D3_DC10{Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf16 _dafny.Array
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf16 _dafny.Array) D3 {
	return D3{D3_DC8{Cf16}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC11 struct {
	Cf21 D3
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf21 D3) D3 {
	return D3{D3_DC11{Cf21}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.Zero)
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf19
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC10).Cf20
}

func (_this D3) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D3_DC8).Cf16
}

func (_this D3) Dtor_cf21() D3 {
	return _this.Get_().(D3_DC11).Cf21
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf18) + ", " + data.Cf19.VerbatimString(true) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19.Equals(data2.Cf19) && data1.Cf20 == data2.Cf20
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf16 == data2.Cf16
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf22 _dafny.Set
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf22 _dafny.Set) D4 {
	return D4{D4_DC12{Cf22}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D4) Dtor_cf22() _dafny.Set {
	return _this.Get_().(D4_DC12).Cf22
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf22.Equals(data2.Cf22)
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

type D5_DC14 struct {
	Cf24 bool
	Cf25 bool
	Cf26 *C0
	Cf27 _dafny.Int
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf24 bool, Cf25 bool, Cf26 *C0, Cf27 _dafny.Int) D5 {
	return D5{D5_DC14{Cf24, Cf25, Cf26, Cf27}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf23 _dafny.Set
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf23 _dafny.Set) D5 {
	return D5{D5_DC13{Cf23}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(false, false, (*C0)(nil), _dafny.Zero)
}

func (_this D5) Dtor_cf24() bool {
	return _this.Get_().(D5_DC14).Cf24
}

func (_this D5) Dtor_cf25() bool {
	return _this.Get_().(D5_DC14).Cf25
}

func (_this D5) Dtor_cf26() *C0 {
	return _this.Get_().(D5_DC14).Cf26
}

func (_this D5) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf27
}

func (_this D5) Dtor_cf23() _dafny.Set {
	return _this.Get_().(D5_DC13).Cf23
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf24 == data2.Cf24 && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D6_DC15 struct {
	Cf28 _dafny.Array
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf28 _dafny.Array) D6 {
	return D6{D6_DC15{Cf28}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D6) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D6_DC15).Cf28
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf28 == data2.Cf28
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

type D7_DC16 struct {
	Cf29 _dafny.Map
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf29 _dafny.Map) D7 {
	return D7{D7_DC16{Cf29}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D7) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D7_DC16).Cf29
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf29.Equals(data2.Cf29)
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

type D8_DC17 struct {
	Cf30 _dafny.Sequence
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf30 _dafny.Sequence) D8 {
	return D8{D8_DC17{Cf30}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf30() _dafny.Sequence {
	return _this.Get_().(D8_DC17).Cf30
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	F2   _dafny.Int
	F11  _dafny.Sequence
	_f0  bool
	_f1  _dafny.Array
	_f3  bool
	_f4  _dafny.Int
	_f5  bool
	_f6  bool
	_f7  _dafny.Int
	_f8  bool
	_f9  _dafny.Array
	_f10 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.Zero
	_this.F11 = _dafny.EmptySeq
	_this._f0 = false
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = false
	_this._f4 = _dafny.Zero
	_this._f5 = false
	_this._f6 = false
	_this._f7 = _dafny.Zero
	_this._f8 = false
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f10 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Array, f2 _dafny.Int, f3 bool, f4 _dafny.Int, f5 bool, f6 bool, f7 _dafny.Int, f8 bool, f9 _dafny.Array, f10 _dafny.Int, f11 _dafny.Sequence) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
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
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Array {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	dummy byte
}

func New_C0_() *C0 {
	_this := C0{}

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

func (_this *C0) Ctor__() {
	{
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
