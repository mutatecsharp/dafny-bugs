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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D3_.Create_DC8_()), _dafny.SeqOf(Companion_D3_.Create_DC8_(), Companion_D3_.Create_DC8_())))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gyruwxl")).Cardinality()), (_dafny.IntOfInt64(123)).Times(_dafny.IntOfInt64(758)))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))).Merge(func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg0 _dafny.Int) interface{} {
						return coer0(arg0)
					}
				}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}))).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _1_v1 _dafny.CodePoint
					_1_v1 = interface{}(_compr_2).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg1 _dafny.Int) interface{} {
							return coer1(arg1)
						}
					}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('e')
					})), _1_v1) {
						_coll2.Add(_1_v1, _dafny.IntOfInt64(904))
					}
				}
				return _coll2.ToMap()
			}()).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _2_v0 _dafny.CodePoint
				_2_v0 = interface{}(_compr_1).(_dafny.CodePoint)
				if (func() _dafny.Map {
					var _coll3 = _dafny.NewMapBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg2 _dafny.Int) interface{} {
							return coer2(arg2)
						}
					}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('e')
					}))).Elements()); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _1_v1 _dafny.CodePoint
						_1_v1 = interface{}(_compr_3).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg3 _dafny.Int) interface{} {
								return coer3(arg3)
							}
						}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('e')
						})), _1_v1) {
							_coll3.Add(_1_v1, _dafny.IntOfInt64(904))
						}
					}
					return _coll3.ToMap()
				}()).Contains(_2_v0) {
					_coll1.Add(_2_v0, _dafny.SeqOf(_dafny.IntOfInt64(587)))
				}
			}
			return _coll1.ToMap()
		}())).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _3_v2 _dafny.CodePoint
			_3_v2 = interface{}(_compr_0).(_dafny.CodePoint)
			if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))).Merge(func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((func() _dafny.Map {
					var _coll5 = _dafny.NewMapBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg4 _dafny.Int) interface{} {
							return coer4(arg4)
						}
					}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('e')
					}))).Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _1_v1 _dafny.CodePoint
						_1_v1 = interface{}(_compr_5).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg5 _dafny.Int) interface{} {
								return coer5(arg5)
							}
						}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('e')
						})), _1_v1) {
							_coll5.Add(_1_v1, _dafny.IntOfInt64(904))
						}
					}
					return _coll5.ToMap()
				}()).Keys().Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _2_v0 _dafny.CodePoint
					_2_v0 = interface{}(_compr_4).(_dafny.CodePoint)
					if (func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg6 _dafny.Int) interface{} {
								return coer6(arg6)
							}
						}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('e')
						}))).Elements()); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _1_v1 _dafny.CodePoint
							_1_v1 = interface{}(_compr_6).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg7 _dafny.Int) interface{} {
									return coer7(arg7)
								}
							}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('e')
							})), _1_v1) {
								_coll6.Add(_1_v1, _dafny.IntOfInt64(904))
							}
						}
						return _coll6.ToMap()
					}()).Contains(_2_v0) {
						_coll4.Add(_2_v0, _dafny.SeqOf(_dafny.IntOfInt64(587)))
					}
				}
				return _coll4.ToMap()
			}())).Contains(_3_v2) {
				_coll0.Add(_3_v2)
			}
		}
		return _coll0.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(!(false), false)).Cardinality()))).Difference((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(844))).Cardinality()))
		}
		return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(118))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('q')
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), false, true), _dafny.SeqOf(true))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("b"))).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("bl"), _dafny.UnicodeSeqOfUtf8Bytes("oj")))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _5_v0 _dafny.Array
	_ = _5_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
	_ = _nw0
	_5_v0 = _nw0
	var _6_v1 bool
	_ = _6_v1
	_6_v1 = true
	var _7_v2 _dafny.Sequence
	_ = _7_v2
	_7_v2 = _dafny.SeqOf(!(_6_v1), _6_v1)
	var _8_v3 _dafny.Int
	_ = _8_v3
	_8_v3 = _dafny.IntOfInt64(861)
	var _9_v4 _dafny.Set
	_ = _9_v4
	_9_v4 = _dafny.SetOf(_8_v3, _8_v3)
	var _10_v5 bool
	_ = _10_v5
	_10_v5 = false
	var _11_v6 _dafny.Array
	_ = _11_v6
	var _nwElement0_0 bool = true
	_ = _nwElement0_0
	var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(12))
	_ = _nw1
	(_nw1).ArraySet1(_nwElement0_0, 0)
	(_nw1).ArraySet1((_10_v5), 1)
	(_nw1).ArraySet1(false, 2)
	(_nw1).ArraySet1((_7_v2).Select((Companion_Default___.SafeIndex(_8_v3, _dafny.IntOfUint32((_7_v2).Cardinality()))).Uint32()).(bool), 3)
	(_nw1).ArraySet1(_6_v1, 4)
	(_nw1).ArraySet1(_6_v1, 5)
	(_nw1).ArraySet1(!(_6_v1), 6)
	(_nw1).ArraySet1(_6_v1, 7)
	(_nw1).ArraySet1(_6_v1, 8)
	(_nw1).ArraySet1(_6_v1, 9)
	(_nw1).ArraySet1(!(_6_v1), 10)
	(_nw1).ArraySet1(_6_v1, 11)
	_11_v6 = _nw1
	var _12_v7 _dafny.CodePoint
	_ = _12_v7
	_12_v7 = _dafny.CodePoint('l')
	var _13_v8 _dafny.Map
	_ = _13_v8
	_13_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v3, _12_v7)
	var _14_v9 _dafny.Array
	_ = _14_v9
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
	_ = _nw2
	_14_v9 = _nw2
	var _15_v10 _dafny.Map
	_ = _15_v10
	_15_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_v1, _12_v7)
	var _16_v11 _dafny.MultiSet
	_ = _16_v11
	_16_v11 = _dafny.MultiSetOf(_15_v10, _15_v10)
	var _17_v12 D1
	_ = _17_v12
	_17_v12 = Companion_D1_.Create_DC1_(_5_v0)
	var _18_globalState *GlobalState
	_ = _18_globalState
	var _nw3 *GlobalState = New_GlobalState_()
	_ = _nw3
	_nw3.Ctor__(_5_v0, _dafny.Companion_Sequence_.Concatenate(_7_v2, _7_v2), false, _9_v4, _11_v6, _5_v0, _13_v8, true, false, _14_v9, false, _dafny.CodePoint('s'), _11_v6, _dafny.IntOfInt64(242), true, _dafny.IntOfInt64(442), false, _dafny.IntOfInt64(-101), _dafny.IntOfInt64(92), _dafny.IntOfInt64(999), (_16_v11).Union(_16_v11), false, (_17_v12).Dtor_cf1(), (_17_v12).Dtor_cf1(), _dafny.IntOfInt64(279), _dafny.IntOfInt64(491))
	_18_globalState = _nw3
	if (Companion_Default___.SafeDivisionInt(_8_v3, _8_v3)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_6_v1, _6_v1, _6_v1)).Cardinality())) >= 0 {
		_11_v6 = _11_v6
		var _19_v13 _dafny.Sequence
		_ = _19_v13
		_19_v13 = _dafny.UnicodeSeqOfUtf8Bytes("shmbaf")
		var _20_v14 _dafny.MultiSet
		_ = _20_v14
		_20_v14 = _dafny.MultiSetOf(_19_v13, _19_v13, _19_v13)
		var _21_v15 _dafny.Map
		_ = _21_v15
		_21_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.SeqOf(_8_v3), _8_v3, (_dafny.Zero).Minus(_dafny.IntOfUint32((_19_v13).Cardinality())), _18_globalState), (_dafny.Zero).Minus((_9_v4).Cardinality()))
		_8_v3 = (func() _dafny.Int {
			if (_20_v14).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-845))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_22_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_23_i0 _dafny.Int) _dafny.CodePoint {
					return _22_v7
				}
			})(_12_v7)))) {
				return (_20_v14).Multiplicity(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-845))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_24_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_25_i0 _dafny.Int) _dafny.CodePoint {
						return _24_v7
					}
				})(_12_v7))))
			}
			return (func() _dafny.Int {
				if _6_v1 {
					return (_21_v15).Cardinality()
				}
				return _dafny.IntOfInt64(-872)
			})()
		})()
		_6_v1 = !_dafny.Companion_Sequence_.Contains(_19_v13, _12_v7)
		if (_9_v4).IsProperSubsetOf(_9_v4) {
			(_18_globalState).F15 = _8_v3
			_6_v1 = _6_v1
			var _26_v16 *C0
			_ = _26_v16
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__(_6_v1)
			_26_v16 = _nw4
			(_26_v16).F26 = !((_8_v3).Cmp(Companion_Default___.SafeDivisionInt(_8_v3, _dafny.IntOfInt64(-365))) > 0)
			var _27_v17 bool
			_ = _27_v17
			var _28_v18 _dafny.Sequence
			_ = _28_v18
			var _out0 bool
			_ = _out0
			var _out1 _dafny.Sequence
			_ = _out1
			_out0, _out1 = (_26_v16).M0(_26_v16.F26, (_dafny.Zero).Minus(_dafny.IntOfUint32((_19_v13).Cardinality())), _18_globalState)
			_27_v17 = _out0
			_28_v18 = _out1
		} else {
			var _29_v19 *C0
			_ = _29_v19
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__((_8_v3).Cmp(_8_v3) != 0)
			_29_v19 = _nw5
			var _30_v20 _dafny.Map
			_ = _30_v20
			_30_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v3, (_dafny.Zero).Minus(_8_v3))
			var _31_v21 D1
			_ = _31_v21
			_31_v21 = Companion_D1_.Create_DC2_(_8_v3, (_dafny.Zero).Minus(_8_v3), _8_v3, (_30_v20).Merge(_30_v20))
			_31_v21 = _31_v21
			var _32_v22 _dafny.Array
			_ = _32_v22
			var _nwElement0_1 _dafny.Array = _5_v0
			_ = _nwElement0_1
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(12))
			_ = _nw6
			(_nw6).ArraySet1(_nwElement0_1, 0)
			(_nw6).ArraySet1(_5_v0, 1)
			(_nw6).ArraySet1(_5_v0, 2)
			(_nw6).ArraySet1(_5_v0, 3)
			(_nw6).ArraySet1(_5_v0, 4)
			(_nw6).ArraySet1(_5_v0, 5)
			(_nw6).ArraySet1(_5_v0, 6)
			(_nw6).ArraySet1(_5_v0, 7)
			(_nw6).ArraySet1(_5_v0, 8)
			(_nw6).ArraySet1(_5_v0, 9)
			(_nw6).ArraySet1(_5_v0, 10)
			(_nw6).ArraySet1(_5_v0, 11)
			_32_v22 = _nw6
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_32_v22), 0))
			_ = _index0
			(_32_v22).ArraySet1(_5_v0, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_32_v22), 0))
			_ = _index1
			var _rhs0 _dafny.Sequence = _19_v13
			_ = _rhs0
			var _rhs1 _dafny.Int = _8_v3
			_ = _rhs1
			var _rhs2 _dafny.Array = _5_v0
			_ = _rhs2
			var _rhs3 bool = _29_v19.F26
			_ = _rhs3
			var _rhs4 _dafny.Int = _dafny.IntOfInt64(850)
			_ = _rhs4
			var _lhs0 *GlobalState = _18_globalState
			_ = _lhs0
			var _lhs1 _dafny.Array = _32_v22
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_32_v22), 0))
			_ = _lhs2
			var _lhs3 *C0 = _29_v19
			_ = _lhs3
			var _lhs4 *GlobalState = _18_globalState
			_ = _lhs4
			_19_v13 = _rhs0
			_lhs0.F19 = _rhs1
			(_lhs1).ArraySet1(_rhs2, (_lhs2).Int())
			_lhs3.F26 = _rhs3
			_lhs4.F19 = _rhs4
			_19_v13 = _19_v13
			var _33_v23 D1
			_ = _33_v23
			_33_v23 = Companion_D1_.Create_DC3_(_8_v3)
			_33_v23 = _33_v23
		}
		var _34_v24 *C0
		_ = _34_v24
		var _nw7 *C0 = New_C0_()
		_ = _nw7
		_nw7.Ctor__(_6_v1)
		_34_v24 = _nw7
		var _35_v25 _dafny.Set
		_ = _35_v25
		_35_v25 = _dafny.SetOf(_34_v24)
		var _36_v26 _dafny.MultiSet
		_ = _36_v26
		_36_v26 = _dafny.MultiSetOf((_35_v25).Cardinality(), _8_v3, _dafny.IntOfInt64(721), _8_v3)
		_6_v1 = ((_36_v26).Cardinality()).Cmp(_8_v3) >= 0
	} else {
		_5_v0 = _5_v0
		_12_v7 = _12_v7
		var _37_v27 _dafny.Sequence
		_ = _37_v27
		_37_v27 = _dafny.UnicodeSeqOfUtf8Bytes("u")
		var _38_v28 _dafny.Set
		_ = _38_v28
		_38_v28 = _dafny.SetOf(_6_v1)
		(_18_globalState).F15 = ((_dafny.IntOfUint32((_37_v27).Cardinality())).Times(_8_v3)).Minus((func() _dafny.Int {
			if _6_v1 {
				return _8_v3
			}
			return (_38_v28).Cardinality()
		})())
		var _39_v29 _dafny.Array
		_ = _39_v29
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_0
		var _nw8 _dafny.Array
		_ = _nw8
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw8 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = (func(_40_v3 _dafny.Int, _41_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_42_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(123), _40_v3, _40_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v3, (_dafny.MultiSetFromSeq(_41_v2)).Cardinality())))
				}
			})(_8_v3, _7_v2)
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
		_39_v29 = _nw8
		var _43_v30 D2
		_ = _43_v30
		_43_v30 = Companion_D2_.Create_DC4_(_39_v29)
		_39_v29 = (_43_v30).Dtor_cf7()
		var _44_v31 _dafny.Array
		_ = _44_v31
		var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
		_ = _nw9
		_44_v31 = _nw9
		var _45_v32 _dafny.Map
		_ = _45_v32
		_45_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_v1, _6_v1)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_44_v31), 0))
		_ = _index2
		(_44_v31).ArraySet1(_45_v32, (_index2).Int())
		var _46_v33 _dafny.MultiSet
		_ = _46_v33
		_46_v33 = _dafny.MultiSetOf(_37_v27)
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_44_v31), 0))
		_ = _index3
		var _rhs5 _dafny.Int = ((_46_v33).Difference(_46_v33)).Cardinality()
		_ = _rhs5
		var _rhs6 _dafny.Sequence = _37_v27
		_ = _rhs6
		var _rhs7 _dafny.Map = _45_v32
		_ = _rhs7
		var _lhs5 *GlobalState = _18_globalState
		_ = _lhs5
		var _lhs6 _dafny.Array = _44_v31
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_44_v31), 0))
		_ = _lhs7
		_lhs5.F19 = _rhs5
		_37_v27 = _rhs6
		(_lhs6).ArraySet1(_rhs7, (_lhs7).Int())
	}
	var _47_v34 *C0
	_ = _47_v34
	var _nw10 *C0 = New_C0_()
	_ = _nw10
	_nw10.Ctor__(_6_v1)
	_47_v34 = _nw10
	var _48_i2 _dafny.Int
	_ = _48_i2
	_48_i2 = _dafny.Zero
	{
		for _47_v34.F26 {
			{
				if (_48_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_48_i2 = (_48_i2).Plus(_dafny.One)
				var _49_v35 _dafny.Map
				_ = _49_v35
				_49_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_v1, false)
				var _50_v36 _dafny.Map
				_ = _50_v36
				_50_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_v1, _49_v35)
				var _51_v37 _dafny.Map
				_ = _51_v37
				_51_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v36, _6_v1)
				var _52_v38 _dafny.Sequence
				_ = _52_v38
				_52_v38 = _dafny.SeqOf(_8_v3)
				(_47_v34).F26 = (func() bool {
					if (_51_v37).Contains(_50_v36) {
						return (_51_v37).Get(_50_v36).(bool)
					}
					return Companion_Default___.Fm0(_52_v38, _8_v3, _8_v3, _18_globalState)
				})()
				var _53_v39 _dafny.Array
				_ = _53_v39
				var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
				_ = _nw11
				_53_v39 = _nw11
				var _54_v40 D2
				_ = _54_v40
				_54_v40 = Companion_D2_.Create_DC4_((func() _dafny.Array {
					if _6_v1 {
						return _53_v39
					}
					return _53_v39
				})())
				_54_v40 = _54_v40
				var _55_v41 _dafny.Set
				_ = _55_v41
				_55_v41 = _dafny.SetOf(_6_v1)
				var _56_v42 _dafny.Array
				_ = _56_v42
				var _nwElement0_2 _dafny.Set = (_55_v41).Intersection(_55_v41)
				_ = _nwElement0_2
				var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
				_ = _nw12
				(_nw12).ArraySet1(_nwElement0_2, 0)
				(_nw12).ArraySet1((_dafny.SetOf(_6_v1)).Intersection(_55_v41), 1)
				(_nw12).ArraySet1(_55_v41, 2)
				(_nw12).ArraySet1(_55_v41, 3)
				_56_v42 = _nw12
				var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_56_v42), 0))
				_ = _index4
				(_56_v42).ArraySet1(_55_v41, (_index4).Int())
				var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_56_v42), 0))
				_ = _index5
				(_56_v42).ArraySet1((_55_v41).Intersection(_dafny.SetOf(_47_v34.F26, _47_v34.F26)), (_index5).Int())
				if _6_v1 {
					_49_v35 = (_49_v35).Update(_47_v34.F26, _6_v1)
					var _57_v43 _dafny.Map
					_ = _57_v43
					_57_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v6, _8_v3)
					var _58_v44 _dafny.Map
					_ = _58_v44
					_58_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v3, _dafny.IntOfUint32((_52_v38).Cardinality()))
					_57_v43 = (_57_v43).Update(_11_v6, Companion_Default___.SafeModuloInt((func() _dafny.Int {
						if (_58_v44).Contains(_dafny.IntOfUint32((_52_v38).Cardinality())) {
							return (_58_v44).Get(_dafny.IntOfUint32((_52_v38).Cardinality())).(_dafny.Int)
						}
						return _dafny.IntOfInt64(999)
					})(), _8_v3))
					var _59_v45 bool
					_ = _59_v45
					var _60_v46 _dafny.Sequence
					_ = _60_v46
					var _out2 bool
					_ = _out2
					var _out3 _dafny.Sequence
					_ = _out3
					_out2, _out3 = (_47_v34).M0(_47_v34.F26, _8_v3, _18_globalState)
					_59_v45 = _out2
					_60_v46 = _out3
					var _61_v47 *C0
					_ = _61_v47
					var _nw13 *C0 = New_C0_()
					_ = _nw13
					_nw13.Ctor__(true)
					_61_v47 = _nw13
					var _62_v48 D1
					_ = _62_v48
					_62_v48 = Companion_D1_.Create_DC3_(((_9_v4).Intersection(_9_v4)).Cardinality())
					var _63_v49 _dafny.Array
					_ = _63_v49
					var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
					_ = _nw14
					_63_v49 = _nw14
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_63_v49), 0))
					_ = _index6
					(_63_v49).ArraySet1(_5_v0, (_index6).Int())
					var _64_v50 _dafny.Sequence
					_ = _64_v50
					_64_v50 = _dafny.SeqOf(_11_v6)
					var _65_v51 _dafny.Sequence
					_ = _65_v51
					_65_v51 = _dafny.UnicodeSeqOfUtf8Bytes("qfhvebs")
					var _66_v52 _dafny.Map
					_ = _66_v52
					_66_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v34.F26, _65_v51)
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_63_v49), 0))
					_ = _index7
					var _rhs8 *C0 = _61_v47
					_ = _rhs8
					var _rhs9 _dafny.Int = (Companion_D2_.Create_DC5_(_5_v0, _64_v50, (_66_v52).Cardinality())).Dtor_cf10()
					_ = _rhs9
					var _rhs10 D1 = Companion_D1_.Create_DC3_(_8_v3)
					_ = _rhs10
					var _rhs11 bool = _47_v34.F26
					_ = _rhs11
					var _rhs12 _dafny.Array = _5_v0
					_ = _rhs12
					var _lhs8 *C0 = _47_v34
					_ = _lhs8
					var _lhs9 _dafny.Array = _63_v49
					_ = _lhs9
					var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_63_v49), 0))
					_ = _lhs10
					_61_v47 = _rhs8
					_8_v3 = _rhs9
					_62_v48 = _rhs10
					_lhs8.F26 = _rhs11
					(_lhs9).ArraySet1(_rhs12, (_lhs10).Int())
					var _67_v53 _dafny.Array
					_ = _67_v53
					var _nwElement0_3 _dafny.CodePoint = Companion_Default___.Fm6(_18_globalState)
					_ = _nwElement0_3
					var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(14))
					_ = _nw15
					(_nw15).ArraySet1CodePoint(_nwElement0_3, 0)
					(_nw15).ArraySet1CodePoint(_12_v7, 1)
					(_nw15).ArraySet1CodePoint(_12_v7, 2)
					(_nw15).ArraySet1CodePoint(_12_v7, 3)
					(_nw15).ArraySet1CodePoint(_12_v7, 4)
					(_nw15).ArraySet1CodePoint(_12_v7, 5)
					(_nw15).ArraySet1CodePoint(_12_v7, 6)
					(_nw15).ArraySet1CodePoint(_12_v7, 7)
					(_nw15).ArraySet1CodePoint(_12_v7, 8)
					(_nw15).ArraySet1CodePoint(_12_v7, 9)
					(_nw15).ArraySet1CodePoint(_12_v7, 10)
					(_nw15).ArraySet1CodePoint(_12_v7, 11)
					(_nw15).ArraySet1CodePoint(_12_v7, 12)
					(_nw15).ArraySet1CodePoint(_12_v7, 13)
					_67_v53 = _nw15
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_67_v53), 0))
					_ = _index8
					(_67_v53).ArraySet1CodePoint(_12_v7, (_index8).Int())
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_67_v53), 0))
					_ = _index9
					(_67_v53).ArraySet1CodePoint(_12_v7, (_index9).Int())
				} else {
					var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
					_ = _nw16
					(_18_globalState).F5 = _nw16
					(_18_globalState).F15 = _8_v3
					var _68_v54 _dafny.Map
					_ = _68_v54
					_68_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(), _12_v7)
					var _69_v55 _dafny.Map
					_ = _69_v55
					_69_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v3, _68_v54)
					(_18_globalState).F19 = (((_69_v55).Cardinality()).Minus(_8_v3)).Times(_8_v3)
					_7_v2 = _dafny.Companion_Sequence_.Update(_7_v2, (Companion_Default___.SafeIndex(_8_v3, _dafny.IntOfUint32((_7_v2).Cardinality()))).Uint32(), _47_v34.F26)
					_12_v7 = _12_v7
				}
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _70_v56 _dafny.Array
	_ = _70_v56
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(7)
	_ = _len0_1
	var _nw17 _dafny.Array
	_ = _nw17
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw17 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) D2 = func(_71_i3 _dafny.Int) D2 {
			return Companion_D2_.Create_DC6_()
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw17 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw17).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw17).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_70_v56 = _nw17
	var _72_v57 _dafny.Array
	_ = _72_v57
	var _nwElement0_4 _dafny.Array = _70_v56
	_ = _nwElement0_4
	var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(25))
	_ = _nw18
	(_nw18).ArraySet1(_nwElement0_4, 0)
	(_nw18).ArraySet1(_70_v56, 1)
	(_nw18).ArraySet1(_70_v56, 2)
	(_nw18).ArraySet1((func() _dafny.Array {
		if _47_v34.F26 {
			return _70_v56
		}
		return _70_v56
	})(), 3)
	(_nw18).ArraySet1(_70_v56, 4)
	(_nw18).ArraySet1((func() _dafny.Array {
		if !(_47_v34.F26) {
			return _70_v56
		}
		return _70_v56
	})(), 5)
	(_nw18).ArraySet1(_70_v56, 6)
	(_nw18).ArraySet1(_70_v56, 7)
	(_nw18).ArraySet1(_70_v56, 8)
	(_nw18).ArraySet1(_70_v56, 9)
	(_nw18).ArraySet1(_70_v56, 10)
	(_nw18).ArraySet1(_70_v56, 11)
	(_nw18).ArraySet1(_70_v56, 12)
	(_nw18).ArraySet1(_70_v56, 13)
	(_nw18).ArraySet1((func() _dafny.Array {
		if _6_v1 {
			return _70_v56
		}
		return _70_v56
	})(), 14)
	(_nw18).ArraySet1(_70_v56, 15)
	(_nw18).ArraySet1(_70_v56, 16)
	(_nw18).ArraySet1(_70_v56, 17)
	(_nw18).ArraySet1(_70_v56, 18)
	(_nw18).ArraySet1(_70_v56, 19)
	(_nw18).ArraySet1(_70_v56, 20)
	(_nw18).ArraySet1(_70_v56, 21)
	(_nw18).ArraySet1(_70_v56, 22)
	(_nw18).ArraySet1(_70_v56, 23)
	(_nw18).ArraySet1(_70_v56, 24)
	_72_v57 = _nw18
	_72_v57 = (func() _dafny.Array {
		if _47_v34.F26 {
			return _72_v57
		}
		return _72_v57
	})()
	var _73_v58 _dafny.Sequence
	_ = _73_v58
	_73_v58 = _dafny.UnicodeSeqOfUtf8Bytes("xseuq")
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
	_ = _index10
	(_5_v0).ArraySet1(_dafny.IntOfUint32((_73_v58).Cardinality()), (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
	_ = _index11
	(_5_v0).ArraySet1((_dafny.Zero).Minus(_8_v3), (_index11).Int())
	var _hi0 _dafny.Int = _8_v3
	_ = _hi0
	for _74_i4 := _dafny.IntOfInt64(210); _74_i4.Cmp(_hi0) < 0; _74_i4 = _74_i4.Plus(_dafny.One) {
		var _75_v59 bool
		_ = _75_v59
		var _76_v60 _dafny.Sequence
		_ = _76_v60
		var _out4 bool
		_ = _out4
		var _out5 _dafny.Sequence
		_ = _out5
		_out4, _out5 = (_47_v34).M0(_47_v34.F26, ((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)).Times(_8_v3), _18_globalState)
		_75_v59 = _out4
		_76_v60 = _out5
		var _77_v61 _dafny.Map
		_ = _77_v61
		_77_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _9_v4)
		_77_v61 = (_77_v61).Update(_12_v7, _9_v4)
		var _78_v62 _dafny.Sequence
		_ = _78_v62
		_78_v62 = _dafny.SeqOf(_73_v58)
		if !_dafny.Companion_Sequence_.Equal(_73_v58, _dafny.Companion_Sequence_.Concatenate((_78_v62).Select((Companion_Default___.SafeIndex((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_78_v62).Cardinality()))).Uint32()).(_dafny.Sequence), _73_v58)) {
			(_18_globalState).F5 = _5_v0
			var _79_v63 _dafny.Array
			_ = _79_v63
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw19
			_79_v63 = _nw19
			var _80_v64 _dafny.Map
			_ = _80_v64
			_80_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v63, _5_v0)
			_80_v64 = (_80_v64).Update(_79_v63, _5_v0)
			var _81_v65 _dafny.Sequence
			_ = _81_v65
			_81_v65 = _dafny.SeqOf(_11_v6)
			var _82_v66 D2
			_ = _82_v66
			_82_v66 = Companion_D2_.Create_DC5_(_5_v0, _81_v65, (_dafny.Zero).Minus((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)))
			var _pat_let_tv0 = _5_v0
			_ = _pat_let_tv0
			var _83_v67 _dafny.Map
			_ = _83_v67
			_83_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v34.F26, func(_pat_let0_0 D2) D2 {
				return func(_84_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let1_0 _dafny.Int) D2 {
						return func(_85_dt__update_hcf10_h0 _dafny.Int) D2 {
							return func(_pat_let2_0 _dafny.Array) D2 {
								return func(_86_dt__update_hcf8_h0 _dafny.Array) D2 {
									return Companion_D2_.Create_DC5_(_86_dt__update_hcf8_h0, (_84_dt__update__tmp_h0).Dtor_cf9(), _85_dt__update_hcf10_h0)
								}(_pat_let2_0)
							}(_pat_let_tv0)
						}(_pat_let1_0)
					}(_dafny.IntOfInt64(762))
				}(_pat_let0_0)
			}(_82_v66))
			_83_v67 = _83_v67
			var _87_v68 bool
			_ = _87_v68
			var _88_v69 _dafny.Sequence
			_ = _88_v69
			var _out6 bool
			_ = _out6
			var _out7 _dafny.Sequence
			_ = _out7
			_out6, _out7 = (_47_v34).M0((_dafny.IntOfInt64(826)).Cmp(_74_i4) < 0, _dafny.IntOfInt64(143), _18_globalState)
			_87_v68 = _out6
			_88_v69 = _out7
			(_47_v34).F26 = _47_v34.F26
		} else {
			var _89_v70 _dafny.Map
			_ = _89_v70
			_89_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_47_v34.F26) || (_6_v1), _dafny.Companion_Sequence_.IsProperPrefixOf(_73_v58, _73_v58))
			_89_v70 = (_89_v70).Update(_47_v34.F26, !(_75_v59) || (_47_v34.F26))
			_6_v1 = !(_6_v1) || (_47_v34.F26)
			var _90_v71 bool
			_ = _90_v71
			var _91_v72 _dafny.Sequence
			_ = _91_v72
			var _out8 bool
			_ = _out8
			var _out9 _dafny.Sequence
			_ = _out9
			_out8, _out9 = (_47_v34).M0(_75_v59, _dafny.IntOfUint32((Companion_Default___.Fm7(_18_globalState)).Cardinality()), _18_globalState)
			_90_v71 = _out8
			_91_v72 = _out9
			var _92_v73 *C0
			_ = _92_v73
			var _nw20 *C0 = New_C0_()
			_ = _nw20
			_nw20.Ctor__(_90_v71)
			_92_v73 = _nw20
			_92_v73 = _47_v34
			var _93_v74 bool
			_ = _93_v74
			var _94_v75 _dafny.Sequence
			_ = _94_v75
			var _out10 bool
			_ = _out10
			var _out11 _dafny.Sequence
			_ = _out11
			_out10, _out11 = (_92_v73).M0(true, _74_i4, _18_globalState)
			_93_v74 = _out10
			_94_v75 = _out11
		}
		var _95_v76 _dafny.Sequence
		_ = _95_v76
		_95_v76 = _dafny.SeqOf(_7_v2, _7_v2)
		var _96_v77 D3
		_ = _96_v77
		_96_v77 = Companion_D3_.Create_DC7_(_7_v2)
		var _97_v78 _dafny.MultiSet
		_ = _97_v78
		_97_v78 = _dafny.MultiSetOf((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
		var _98_v79 _dafny.Map
		_ = _98_v79
		_98_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v78, _dafny.SeqOf(_6_v1))
		var _99_v80 _dafny.Array
		_ = _99_v80
		var _nwElement0_5 _dafny.Sequence = _7_v2
		_ = _nwElement0_5
		var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(28))
		_ = _nw21
		(_nw21).ArraySet1(_nwElement0_5, 0)
		(_nw21).ArraySet1(_7_v2, 1)
		(_nw21).ArraySet1(_7_v2, 2)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_7_v2, _7_v2), 3)
		(_nw21).ArraySet1(_dafny.SeqOf(_47_v34.F26, _6_v1, true, _75_v59, !(_47_v34.F26)), 4)
		(_nw21).ArraySet1(_7_v2, 5)
		(_nw21).ArraySet1((func() _dafny.Sequence {
			if true {
				return _7_v2
			}
			return _7_v2
		})(), 6)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_7_v2, (_95_v76).Select((Companion_Default___.SafeIndex(_8_v3, _dafny.IntOfUint32((_95_v76).Cardinality()))).Uint32()).(_dafny.Sequence)), 7)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_7_v2, _7_v2), 8)
		(_nw21).ArraySet1(_7_v2, 9)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_47_v34.F26), _dafny.SeqOf(_47_v34.F26)), 10)
		(_nw21).ArraySet1(_7_v2, 11)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_6_v1, _47_v34.F26), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(574))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_100_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
			return func(_101_i5 _dafny.Int) _dafny.Int {
				return (_100_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_100_v0), 0))).Int()).(_dafny.Int)
			}
		})(_5_v0)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_6_v1, _47_v34.F26)).Cardinality()))).Uint32(), _47_v34.F26), 12)
		(_nw21).ArraySet1((_96_v77).Dtor_cf11(), 13)
		(_nw21).ArraySet1(_7_v2, 14)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Update(_7_v2, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32(((_47_v34).Fm1(_75_v59, _18_globalState)).Cardinality())), _dafny.IntOfUint32((_7_v2).Cardinality()))).Uint32(), _47_v34.F26), 15)
		(_nw21).ArraySet1(_7_v2, 16)
		(_nw21).ArraySet1(_7_v2, 17)
		(_nw21).ArraySet1((func() _dafny.Sequence {
			if _6_v1 {
				return _7_v2
			}
			return _7_v2
		})(), 18)
		(_nw21).ArraySet1(_7_v2, 19)
		(_nw21).ArraySet1(_7_v2, 20)
		(_nw21).ArraySet1((_95_v76).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-260), _dafny.IntOfUint32((_95_v76).Cardinality()))).Uint32()).(_dafny.Sequence), 21)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_7_v2, _7_v2), 22)
		(_nw21).ArraySet1(_7_v2, 23)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_7_v2, _7_v2), 24)
		(_nw21).ArraySet1((func() _dafny.Sequence {
			if (_98_v79).Contains(_dafny.MultiSetOf(_8_v3)) {
				return (_98_v79).Get(_dafny.MultiSetOf(_8_v3)).(_dafny.Sequence)
			}
			return _dafny.SeqOf(_47_v34.F26, _6_v1)
		})(), 25)
		(_nw21).ArraySet1(_7_v2, 26)
		(_nw21).ArraySet1(_7_v2, 27)
		_99_v80 = _nw21
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_99_v80), 0))
		_ = _index12
		(_99_v80).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_7_v2, (Companion_Default___.SafeIndex(_8_v3, _dafny.IntOfUint32((_7_v2).Cardinality()))).Uint32(), _47_v34.F26), _7_v2), (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_99_v80), 0))
		_ = _index13
		(_99_v80).ArraySet1(_dafny.SeqOf(!(false), _6_v1, _47_v34.F26), (_index13).Int())
	}
	if _6_v1 {
		if _10_v5 {
			(_47_v34).F26 = _47_v34.F26
			var _102_v81 bool
			_ = _102_v81
			var _103_v82 _dafny.Sequence
			_ = _103_v82
			var _out12 bool
			_ = _out12
			var _out13 _dafny.Sequence
			_ = _out13
			_out12, _out13 = (_47_v34).M0(_47_v34.F26, (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), _18_globalState)
			_102_v81 = _out12
			_103_v82 = _out13
			(_18_globalState).F15 = (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)
			(_18_globalState).F15 = _dafny.IntOfInt64(398)
			(_47_v34).F26 = _102_v81
		} else {
			var _104_v83 bool
			_ = _104_v83
			var _105_v84 _dafny.Sequence
			_ = _105_v84
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Sequence
			_ = _out15
			_out14, _out15 = (_47_v34).M0(_47_v34.F26, _8_v3, _18_globalState)
			_104_v83 = _out14
			_105_v84 = _out15
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))
			_ = _index14
			(_11_v6).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("aqoumw"), _12_v7), (_index14).Int())
			var _106_v85 _dafny.Array
			_ = _106_v85
			var _nwElement0_6 _dafny.CodePoint = _12_v7
			_ = _nwElement0_6
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(5))
			_ = _nw22
			(_nw22).ArraySet1CodePoint(_nwElement0_6, 0)
			(_nw22).ArraySet1CodePoint(_12_v7, 1)
			(_nw22).ArraySet1CodePoint(_12_v7, 2)
			(_nw22).ArraySet1CodePoint(_12_v7, 3)
			(_nw22).ArraySet1CodePoint(_12_v7, 4)
			_106_v85 = _nw22
			var _107_v86 _dafny.Map
			_ = _107_v86
			_107_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_12_v7, _106_v85)
			var _108_v87 _dafny.Array
			_ = _108_v87
			var _nwElement0_7 _dafny.Array = (func() _dafny.Array {
				if _47_v34.F26 {
					return _106_v85
				}
				return _106_v85
			})()
			_ = _nwElement0_7
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(8))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_7, 0)
			(_nw23).ArraySet1(_106_v85, 1)
			(_nw23).ArraySet1(_106_v85, 2)
			(_nw23).ArraySet1(_106_v85, 3)
			(_nw23).ArraySet1(_106_v85, 4)
			(_nw23).ArraySet1((func() _dafny.Array {
				if (_107_v86).Contains(_12_v7) {
					return (_107_v86).Get(_12_v7).(_dafny.Array)
				}
				return _106_v85
			})(), 5)
			(_nw23).ArraySet1(_106_v85, 6)
			(_nw23).ArraySet1(_106_v85, 7)
			_108_v87 = _nw23
			var _109_v88 D3
			_ = _109_v88
			_109_v88 = Companion_D3_.Create_DC9_((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
			var _110_v89 D3
			_ = _110_v89
			_110_v89 = Companion_D3_.Create_DC10_(_109_v88)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))
			_ = _index15
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
			_ = _index16
			var _rhs13 bool = (_10_v5)
			_ = _rhs13
			var _rhs14 _dafny.Array = _108_v87
			_ = _rhs14
			var _rhs15 D3 = (func() D3 {
				if !(_47_v34.F26) {
					return _110_v89
				}
				return _110_v89
			})()
			_ = _rhs15
			var _rhs16 _dafny.Int = (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs16
			var _lhs11 _dafny.Array = _11_v6
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))
			_ = _lhs12
			var _lhs13 _dafny.Array = _5_v0
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
			_ = _lhs14
			(_lhs11).ArraySet1(_rhs13, (_lhs12).Int())
			_108_v87 = _rhs14
			_110_v89 = _rhs15
			(_lhs13).ArraySet1(_rhs16, (_lhs14).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))
			_ = _index17
			var _rhs17 _dafny.Int = (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs17
			var _rhs18 bool = Companion_Default___.Fm0(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-879))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_111_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_112_i6 _dafny.Int) _dafny.Int {
					return _111_v3
				}
			})(_8_v3))), Companion_Default___.SafeDivisionInt((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)), (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), _18_globalState)
			_ = _rhs18
			var _lhs15 *GlobalState = _18_globalState
			_ = _lhs15
			var _lhs16 _dafny.Array = _11_v6
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))
			_ = _lhs17
			_lhs15.F19 = _rhs17
			(_lhs16).ArraySet1(_rhs18, (_lhs17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
			_ = _index18
			(_5_v0).ArraySet1((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), (_index18).Int())
			var _113_v90 _dafny.Sequence
			_ = _113_v90
			_113_v90 = _dafny.SeqOf(_47_v34, _47_v34)
			var _114_v91 _dafny.Sequence
			_ = _114_v91
			_114_v91 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_113_v90).Cardinality())))
			var _115_v92 _dafny.Map
			_ = _115_v92
			_115_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v34.F26, (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
			var _116_v93 _dafny.Map
			_ = _116_v93
			_116_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v3, _115_v92)
			var _117_v94 _dafny.MultiSet
			_ = _117_v94
			_117_v94 = _dafny.MultiSetOf(Companion_Default___.Fm0(_114_v91, (_dafny.Zero).Minus(_8_v3), _8_v3, _18_globalState), !((func() _dafny.Map {
				if (_116_v93).Contains((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)) {
					return (_116_v93).Get((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)).(_dafny.Map)
				}
				return _115_v92
			})()).Contains(_47_v34.F26))
			var _118_v95 _dafny.Map
			_ = _118_v95
			_118_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), (_11_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))).Int()).(bool))
			var _119_v96 _dafny.Map
			_ = _119_v96
			_119_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_12_v7, _dafny.MultiSetOf((_11_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))).Int()).(bool), (_11_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))).Int()).(bool)))
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
			_ = _index19
			var _rhs19 _dafny.Int = (Companion_Default___.Fm3(_dafny.IntOfInt64(82), _118_v95, (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), _18_globalState)).Plus(Companion_Default___.SafeDivisionInt(_8_v3, (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)))
			_ = _rhs19
			var _rhs20 _dafny.CodePoint = _12_v7
			_ = _rhs20
			var _rhs21 bool = (_11_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_11_v6), 0))).Int()).(bool)
			_ = _rhs21
			var _rhs22 _dafny.MultiSet = (func() _dafny.MultiSet {
				if (_119_v96).Contains(_12_v7) {
					return (_119_v96).Get(_12_v7).(_dafny.MultiSet)
				}
				return _117_v94
			})()
			_ = _rhs22
			var _lhs18 _dafny.Array = _5_v0
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
			_ = _lhs19
			var _lhs20 *C0 = _47_v34
			_ = _lhs20
			(_lhs18).ArraySet1(_rhs19, (_lhs19).Int())
			_12_v7 = _rhs20
			_lhs20.F26 = _rhs21
			_117_v94 = _rhs22
		}
		if _6_v1 {
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
			_ = _index20
			(_5_v0).ArraySet1(_8_v3, (_index20).Int())
			var _120_v97 _dafny.Sequence
			_ = _120_v97
			_120_v97 = _dafny.SeqOf((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
			var _121_v98 _dafny.Map
			_ = _121_v98
			_121_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(_120_v97, _8_v3), _47_v34.F26)
			_121_v98 = (_121_v98).Update(_47_v34.F26, !(_47_v34.F26))
			var _122_v99 _dafny.Set
			_ = _122_v99
			_122_v99 = _dafny.SetOf(_47_v34.F26)
			_122_v99 = _122_v99
			var _123_v100 *C0
			_ = _123_v100
			var _nw24 *C0 = New_C0_()
			_ = _nw24
			_nw24.Ctor__(false)
			_123_v100 = _nw24
			var _124_v101 _dafny.Sequence
			_ = _124_v101
			_124_v101 = _dafny.SeqOf(_123_v100, _123_v100)
			var _125_v102 D4
			_ = _125_v102
			_125_v102 = Companion_D4_.Create_DC11_((_124_v101).Select((Companion_Default___.SafeIndex((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_124_v101).Cardinality()))).Uint32()).(*C0))
			_123_v100 = (_125_v102).Dtor_cf14()
			var _126_v103 _dafny.Array
			_ = _126_v103
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
			_ = _nw25
			_126_v103 = _nw25
			var _127_v104 _dafny.Map
			_ = _127_v104
			_127_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, _126_v103)
			var _128_v105 _dafny.Array
			_ = _128_v105
			var _nwElement0_8 _dafny.Array = _126_v103
			_ = _nwElement0_8
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_8, 0)
			(_nw26).ArraySet1(_126_v103, 1)
			(_nw26).ArraySet1(_126_v103, 2)
			(_nw26).ArraySet1(_126_v103, 3)
			(_nw26).ArraySet1(_126_v103, 4)
			(_nw26).ArraySet1(_126_v103, 5)
			(_nw26).ArraySet1(_126_v103, 6)
			(_nw26).ArraySet1(_126_v103, 7)
			(_nw26).ArraySet1(_126_v103, 8)
			(_nw26).ArraySet1((func() _dafny.Array {
				if (_127_v104).Contains(_5_v0) {
					return (_127_v104).Get(_5_v0).(_dafny.Array)
				}
				return _126_v103
			})(), 9)
			(_nw26).ArraySet1(_126_v103, 10)
			(_nw26).ArraySet1(_126_v103, 11)
			(_nw26).ArraySet1(_126_v103, 12)
			(_nw26).ArraySet1(_126_v103, 13)
			_128_v105 = _nw26
			var _129_v106 _dafny.Map
			_ = _129_v106
			_129_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v2, _128_v105)
			_129_v106 = (_129_v106).Update(_dafny.Companion_Sequence_.Concatenate(_7_v2, _7_v2), _128_v105)
		} else {
			var _130_v107 *C0
			_ = _130_v107
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__(true)
			_130_v107 = _nw27
			(_47_v34).F26 = true
			_12_v7 = _12_v7
			var _131_v108 _dafny.Map
			_ = _131_v108
			_131_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v107, _12_v7)
			var _132_v109 _dafny.Sequence
			_ = _132_v109
			_132_v109 = _dafny.SeqOf((_131_v108).Cardinality(), _dafny.IntOfInt64(254))
			var _133_v110 _dafny.Map
			_ = _133_v110
			_133_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v34.F26, _47_v34.F26)).Cardinality()), _47_v34.F26)
			var _134_v111 _dafny.Map
			_ = _134_v111
			_134_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v110, _dafny.IntOfUint32((_73_v58).Cardinality()))
			var _rhs23 _dafny.Sequence = _132_v109
			_ = _rhs23
			var _rhs24 _dafny.Int = ((_134_v111).Cardinality()).Plus(((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)).Minus((_dafny.MultiSetOf(_130_v107.F26)).Cardinality()))
			_ = _rhs24
			var _lhs21 *GlobalState = _18_globalState
			_ = _lhs21
			_132_v109 = _rhs23
			_lhs21.F19 = _rhs24
			var _135_v112 _dafny.Sequence
			_ = _135_v112
			_135_v112 = _dafny.SeqOf(_73_v58)
			var _136_v113 _dafny.MultiSet
			_ = _136_v113
			_136_v113 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("d"))
			var _137_v114 _dafny.Array
			_ = _137_v114
			var _nwElement0_9 _dafny.MultiSet = (_dafny.MultiSetOf(_73_v58, _73_v58, _dafny.UnicodeSeqOfUtf8Bytes("cbpcgr"))).Intersection(_dafny.MultiSetFromSeq(_135_v112))
			_ = _nwElement0_9
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_9, 0)
			(_nw28).ArraySet1(_136_v113, 1)
			(_nw28).ArraySet1((_136_v113).Union(Companion_Default___.Fm8(_18_globalState)), 2)
			(_nw28).ArraySet1((_dafny.MultiSetOf(_73_v58)).Update(_dafny.UnicodeSeqOfUtf8Bytes("ihs"), Companion_Default___.Abs(_8_v3)), 3)
			(_nw28).ArraySet1(Companion_Default___.Fm8(_18_globalState), 4)
			(_nw28).ArraySet1(_136_v113, 5)
			(_nw28).ArraySet1((_dafny.MultiSetOf(_73_v58, _73_v58)).Union(_136_v113), 6)
			_137_v114 = _nw28
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_137_v114), 0))
			_ = _index21
			(_137_v114).ArraySet1(_136_v113, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_137_v114), 0))
			_ = _index22
			(_137_v114).ArraySet1(_136_v113, (_index22).Int())
		}
		_11_v6 = _11_v6
		(_18_globalState).F15 = (_dafny.IntOfInt64(906)).Plus(_dafny.IntOfInt64(758))
		var _138_v115 _dafny.Map
		_ = _138_v115
		_138_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC5_(_5_v0, _dafny.SeqOf(_11_v6, _11_v6, _11_v6, _11_v6), _dafny.IntOfUint32((_7_v2).Cardinality())), _73_v58)
		_138_v115 = _138_v115
	} else {
		var _139_v116 _dafny.Sequence
		_ = _139_v116
		_139_v116 = _dafny.SeqOf(_11_v6, _11_v6, _11_v6, _11_v6)
		var _140_v117 D2
		_ = _140_v117
		_140_v117 = Companion_D2_.Create_DC5_(_5_v0, _139_v116, (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
		var _141_v118 _dafny.Sequence
		_ = _141_v118
		_141_v118 = _dafny.SeqOf(_139_v116)
		_140_v117 = Companion_D2_.Create_DC5_(_5_v0, (_141_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_5_v0, _5_v0, _5_v0)).Cardinality()), _dafny.IntOfUint32((_141_v118).Cardinality()))).Uint32()).(_dafny.Sequence), (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
		var _142_v119 _dafny.Sequence
		_ = _142_v119
		_142_v119 = _dafny.SeqOf(_dafny.IntOfInt64(195))
		var _143_v120 _dafny.Map
		_ = _143_v120
		_143_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_7_v2).Select((Companion_Default___.SafeIndex((_142_v119).Select((Companion_Default___.SafeIndex(_8_v3, _dafny.IntOfUint32((_142_v119).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_7_v2).Cardinality()))).Uint32()).(bool), _8_v3)
		var _144_v121 _dafny.MultiSet
		_ = _144_v121
		_144_v121 = _dafny.MultiSetOf((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
		var _145_v122 _dafny.Map
		_ = _145_v122
		_145_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), (_144_v121).Cardinality())
		var _146_v123 _dafny.MultiSet
		_ = _146_v123
		_146_v123 = _dafny.MultiSetOf(_145_v122)
		_143_v120 = (_143_v120).Update(_47_v34.F26, (_146_v123).Cardinality())
		var _147_v124 _dafny.Map
		_ = _147_v124
		_147_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_6_v1), _47_v34.F26)
		_147_v124 = (_147_v124).Update(_47_v34.F26, (_7_v2).Select((Companion_Default___.SafeIndex(_8_v3, _dafny.IntOfUint32((_7_v2).Cardinality()))).Uint32()).(bool))
		var _148_v125 *C0
		_ = _148_v125
		var _nw29 *C0 = New_C0_()
		_ = _nw29
		_nw29.Ctor__((_47_v34.F26) == (_6_v1))
		_148_v125 = _nw29
		var _149_v126 _dafny.Set
		_ = _149_v126
		_149_v126 = _dafny.SetOf(true)
		_149_v126 = (_149_v126).Union(_149_v126)
	}
	var _150_v127 _dafny.Map
	_ = _150_v127
	_150_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_5_v0, _dafny.IntOfUint32((_7_v2).Cardinality()))
	(_47_v34).F26 = (Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _6_v1)).Cardinality(), (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))).Cmp(((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus((_150_v127).Cardinality()))) != 0
	(_18_globalState).F19 = (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)
	var _151_v128 _dafny.MultiSet
	_ = _151_v128
	_151_v128 = _dafny.MultiSetOf(_47_v34.F26, _6_v1)
	var _152_v129 _dafny.Map
	_ = _152_v129
	_152_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("stmah"), _73_v58)).Cardinality()), !((_151_v128).IsProperSubsetOf(_151_v128)))
	_152_v129 = (_152_v129).Update(_8_v3, _6_v1)
	var _153_v130 _dafny.Array
	_ = _153_v130
	var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
	_ = _nw30
	_153_v130 = _nw30
	var _154_v131 D4
	_ = _154_v131
	_154_v131 = Companion_D4_.Create_DC12_(_153_v130, _11_v6, _6_v1)
	var _155_v132 D4
	_ = _155_v132
	_155_v132 = Companion_D4_.Create_DC13_(_154_v131)
	_155_v132 = Companion_D4_.Create_DC13_(_154_v131)
	var _156_v133 *C0
	_ = _156_v133
	var _nw31 *C0 = New_C0_()
	_ = _nw31
	_nw31.Ctor__(_6_v1)
	_156_v133 = _nw31
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
	_ = _index23
	var _rhs25 _dafny.Int = _dafny.IntOfInt64(872)
	_ = _rhs25
	var _rhs26 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_7_v2, _7_v2)
	_ = _rhs26
	var _rhs27 _dafny.Int = ((_8_v3).Times(_dafny.IntOfInt64(165))).Times((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))
	_ = _rhs27
	var _rhs28 _dafny.Sequence = _dafny.SeqOf(_6_v1, _156_v133.F26, _47_v34.F26)
	_ = _rhs28
	var _lhs22 _dafny.Array = _5_v0
	_ = _lhs22
	var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
	_ = _lhs23
	var _lhs24 *GlobalState = _18_globalState
	_ = _lhs24
	(_lhs22).ArraySet1(_rhs25, (_lhs23).Int())
	_7_v2 = _rhs26
	_lhs24.F19 = _rhs27
	_7_v2 = _rhs28
	if _47_v34.F26 {
		(_47_v34).F26 = _47_v34.F26
		var _157_v134 *C0
		_ = _157_v134
		var _nw32 *C0 = New_C0_()
		_ = _nw32
		_nw32.Ctor__(_47_v34.F26)
		_157_v134 = _nw32
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
		_ = _index24
		(_5_v0).ArraySet1(Companion_Default___.SafeModuloInt((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), _8_v3), (_index24).Int())
		var _158_v135 _dafny.Map
		_ = _158_v135
		_158_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v34.F26, _157_v134.F26)
		var _159_v136 bool
		_ = _159_v136
		var _160_v137 _dafny.Sequence
		_ = _160_v137
		var _out16 bool
		_ = _out16
		var _out17 _dafny.Sequence
		_ = _out17
		_out16, _out17 = (_157_v134).M0((_6_v1) == ((_7_v2).Select((Companion_Default___.SafeIndex(_8_v3, _dafny.IntOfUint32((_7_v2).Cardinality()))).Uint32()).(bool)), (_158_v135).Cardinality(), _18_globalState)
		_159_v136 = _out16
		_160_v137 = _out17
		_8_v3 = _8_v3
	} else {
		var _161_v138 bool
		_ = _161_v138
		var _162_v139 _dafny.Sequence
		_ = _162_v139
		var _out18 bool
		_ = _out18
		var _out19 _dafny.Sequence
		_ = _out19
		_out18, _out19 = (_156_v133).M0(_47_v34.F26, _8_v3, _18_globalState)
		_161_v138 = _out18
		_162_v139 = _out19
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
		_ = _index25
		(_5_v0).ArraySet1(_8_v3, (_index25).Int())
		_156_v133 = _156_v133
		var _163_v140 D5
		_ = _163_v140
		_163_v140 = Companion_D5_.Create_DC14_(_152_v129)
		var _164_v141 _dafny.MultiSet
		_ = _164_v141
		_164_v141 = _dafny.MultiSetOf(_73_v58)
		var _165_v142 _dafny.Map
		_ = _165_v142
		_165_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_73_v58, (Companion_Default___.SafeIndex(Companion_Default___.Fm3(_8_v3, (_163_v140).Dtor_cf19(), _8_v3, _18_globalState), _dafny.IntOfUint32((_73_v58).Cardinality()))).Uint32(), _12_v7), (_164_v141).Cardinality())
		_165_v142 = (_165_v142).Update(_73_v58, (_8_v3).Plus(_8_v3))
		if (!(_47_v34.F26)) == (!(_156_v133.F26)) {
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))
			_ = _index26
			(_5_v0).ArraySet1((_dafny.Zero).Minus(_8_v3), (_index26).Int())
			(_18_globalState).F5 = _5_v0
			_12_v7 = (func() _dafny.CodePoint {
				if (_15_v10).Contains(_47_v34.F26) {
					return (_15_v10).Get(_47_v34.F26).(_dafny.CodePoint)
				}
				return _12_v7
			})()
			var _166_v143 _dafny.Map
			_ = _166_v143
			_166_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_12_v7, _5_v0)
			var _167_v144 _dafny.Array
			_ = _167_v144
			var _nwElement0_10 _dafny.Array = (func() _dafny.Array {
				if (_166_v143).Contains(_12_v7) {
					return (_166_v143).Get(_12_v7).(_dafny.Array)
				}
				return _5_v0
			})()
			_ = _nwElement0_10
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_10, 0)
			(_nw33).ArraySet1(_5_v0, 1)
			(_nw33).ArraySet1(_5_v0, 2)
			(_nw33).ArraySet1(_5_v0, 3)
			(_nw33).ArraySet1(_5_v0, 4)
			(_nw33).ArraySet1(_5_v0, 5)
			(_nw33).ArraySet1(_5_v0, 6)
			(_nw33).ArraySet1(_5_v0, 7)
			(_nw33).ArraySet1(_5_v0, 8)
			(_nw33).ArraySet1(_5_v0, 9)
			(_nw33).ArraySet1(_5_v0, 10)
			(_nw33).ArraySet1(_5_v0, 11)
			(_nw33).ArraySet1(_5_v0, 12)
			_167_v144 = _nw33
			_167_v144 = _167_v144
			var _168_v145 bool
			_ = _168_v145
			var _169_v146 _dafny.Sequence
			_ = _169_v146
			var _out20 bool
			_ = _out20
			var _out21 _dafny.Sequence
			_ = _out21
			_out20, _out21 = (_47_v34).M0(_47_v34.F26, Companion_Default___.SafeModuloInt(_8_v3, _8_v3), _18_globalState)
			_168_v145 = _out20
			_169_v146 = _out21
		} else {
			_161_v138 = ((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)).Cmp((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)) > 0
			var _170_v147 D2
			_ = _170_v147
			_170_v147 = Companion_D2_.Create_DC6_()
			_170_v147 = _170_v147
			(_156_v133).F26 = _47_v34.F26
			var _171_v148 bool
			_ = _171_v148
			var _172_v149 _dafny.Sequence
			_ = _172_v149
			var _out22 bool
			_ = _out22
			var _out23 _dafny.Sequence
			_ = _out23
			_out22, _out23 = (_47_v34).M0(_47_v34.F26, _8_v3, _18_globalState)
			_171_v148 = _out22
			_172_v149 = _out23
			var _173_v150 _dafny.Array
			_ = _173_v150
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_2
			var _nw34 _dafny.Array
			_ = _nw34
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw34 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_174_v58 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_175_i7 _dafny.Int) _dafny.Sequence {
						return _174_v58
					}
				})(_73_v58)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw34 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw34).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw34).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_173_v150 = _nw34
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_173_v150), 0))
			_ = _index27
			(_173_v150).ArraySet1(_73_v58, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_173_v150), 0))
			_ = _index28
			(_173_v150).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _47_v34.F26 {
					return _73_v58
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("faxxuf")
			})(), _73_v58), (_index28).Int())
		}
	}
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_14_v9), 0))); ; {
		_guard_loop_0, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _176_i8 _dafny.Int
		_176_i8 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_176_i8).Sign() != -1) && ((_176_i8).Cmp(_dafny.ArrayLen((_14_v9), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_14_v9, (_176_i8).Int(), _dafny.SetOf(((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)).Minus((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-379)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-899))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_177_i9 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(212)
			}))), _dafny.SeqOf(_dafny.SeqOf((_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int), _8_v3, (_5_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_5_v0), 0))).Int()).(_dafny.Int))))).Cardinality()), _dafny.IntOfInt64(934), (_dafny.Zero).Minus(_8_v3))))
		}
	}
	for _iter8 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	_6_v1 = !(_47_v34.F26)
	_dafny.Print((_5_v0).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_6_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_7_v2, _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_8_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_9_v4).Equals(_dafny.SetOf(_dafny.IntOfInt64(861))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_10_v5))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_11_v6).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_12_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_13_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(861), _dafny.CodePoint('l'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_14_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_14_v9).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_14_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_14_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_14_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_15_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_16_v11).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_17_v12).Dtor_cf1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_18_globalState).F1(), _dafny.SeqOf(false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F3()).Equals(_dafny.SetOf(_dafny.IntOfInt64(861))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState.F5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F6()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(861), _dafny.CodePoint('l'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_18_globalState).F9()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_18_globalState).F9()).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_18_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_18_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_18_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(3), _dafny.IntOfInt64(934), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_18_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_18_globalState.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState.F20).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F22()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_18_globalState).F23()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_globalState).F25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_v34.F26)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_48_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_73_v58.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_150_v127).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v128).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v129).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(10), true).UpdateUnsafe(_dafny.One, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v131).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v131).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_155_v132).Dtor_cf18()).Dtor_cf16()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_155_v132).Dtor_cf18()).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_156_v133.F26)
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
	Cf5 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 _dafny.Int, Cf5 _dafny.Map) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4, Cf5}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf6 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.Int) D1 {
	return D1{D1_DC3{Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Array
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Array) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.EmptyMap)
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

func (_this D1) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Equals(data2.Cf5)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
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
	Cf8  _dafny.Array
	Cf9  _dafny.Sequence
	Cf10 _dafny.Int
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf8 _dafny.Array, Cf9 _dafny.Sequence, Cf10 _dafny.Int) D2 {
	return D2{D2_DC5{Cf8, Cf9, Cf10}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

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

type D2_DC4 struct {
	Cf7 _dafny.Array
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf7 _dafny.Array) D2 {
	return D2{D2_DC4{Cf7}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq, _dafny.Zero)
}

func (_this D2) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Equals(data2.Cf9) && data1.Cf10.Cmp(data2.Cf10) == 0
		}
	case D2_DC6:
		{
			_, ok := other.Get_().(D2_DC6)
			return ok
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf7 == data2.Cf7
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

type D3_DC8 struct {
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_() D3 {
	return D3{D3_DC8{}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf12 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf12 _dafny.Int) D3 {
	return D3{D3_DC9{Cf12}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC7 struct {
	Cf11 _dafny.Sequence
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 _dafny.Sequence) D3 {
	return D3{D3_DC7{Cf11}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC10 struct {
	Cf13 D3
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf13 D3) D3 {
	return D3{D3_DC10{Cf13}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_()
}

func (_this D3) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf12
}

func (_this D3) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) Dtor_cf13() D3 {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC8:
		{
			_, ok := other.Get_().(D3_DC8)
			return ok
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11)
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
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

type D4_DC12 struct {
	Cf15 _dafny.Array
	Cf16 _dafny.Array
	Cf17 bool
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf15 _dafny.Array, Cf16 _dafny.Array, Cf17 bool) D4 {
	return D4{D4_DC12{Cf15, Cf16, Cf17}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC11 struct {
	Cf14 *C0
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf14 *C0) D4 {
	return D4{D4_DC11{Cf14}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC13 struct {
	Cf18 D4
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf18 D4) D4 {
	return D4{D4_DC13{Cf18}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D4) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC12).Cf17
}

func (_this D4) Dtor_cf14() *C0 {
	return _this.Get_().(D4_DC11).Cf14
}

func (_this D4) Dtor_cf18() D4 {
	return _this.Get_().(D4_DC13).Cf18
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf15 == data2.Cf15 && data1.Cf16 == data2.Cf16 && data1.Cf17 == data2.Cf17
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf14 == data2.Cf14
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 *C0
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf20 *C0) D5 {
	return D5{D5_DC15{Cf20}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC14 struct {
	Cf19 _dafny.Map
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf19 _dafny.Map) D5 {
	return D5{D5_DC14{Cf19}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC16 struct {
	Cf21 D5
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf21 D5) D5 {
	return D5{D5_DC16{Cf21}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_((*C0)(nil))
}

func (_this D5) Dtor_cf20() *C0 {
	return _this.Get_().(D5_DC15).Cf20
}

func (_this D5) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D5_DC14).Cf19
}

func (_this D5) Dtor_cf21() D5 {
	return _this.Get_().(D5_DC16).Cf21
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf20 == data2.Cf20
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf19.Equals(data2.Cf19)
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf21.Equals(data2.Cf21)
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

// Definition of class GlobalState
type GlobalState struct {
	F5   _dafny.Array
	F15  _dafny.Int
	F19  _dafny.Int
	F20  _dafny.MultiSet
	_f0  _dafny.Array
	_f1  _dafny.Sequence
	_f2  bool
	_f3  _dafny.Set
	_f4  _dafny.Array
	_f6  _dafny.Map
	_f7  bool
	_f8  bool
	_f9  _dafny.Array
	_f10 bool
	_f11 _dafny.CodePoint
	_f12 _dafny.Array
	_f13 _dafny.Int
	_f14 bool
	_f16 bool
	_f17 _dafny.Int
	_f18 _dafny.Int
	_f21 bool
	_f22 _dafny.Array
	_f23 _dafny.Array
	_f24 _dafny.Int
	_f25 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F15 = _dafny.Zero
	_this.F19 = _dafny.Zero
	_this.F20 = _dafny.EmptyMultiSet
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.EmptySeq
	_this._f2 = false
	_this._f3 = _dafny.EmptySet
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.EmptyMap
	_this._f7 = false
	_this._f8 = false
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f10 = false
	_this._f11 = _dafny.CodePoint('D')
	_this._f12 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f13 = _dafny.Zero
	_this._f14 = false
	_this._f16 = false
	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.Zero
	_this._f21 = false
	_this._f22 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f23 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f24 = _dafny.Zero
	_this._f25 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Sequence, f2 bool, f3 _dafny.Set, f4 _dafny.Array, f5 _dafny.Array, f6 _dafny.Map, f7 bool, f8 bool, f9 _dafny.Array, f10 bool, f11 _dafny.CodePoint, f12 _dafny.Array, f13 _dafny.Int, f14 bool, f15 _dafny.Int, f16 bool, f17 _dafny.Int, f18 _dafny.Int, f19 _dafny.Int, f20 _dafny.MultiSet, f21 bool, f22 _dafny.Array, f23 _dafny.Array, f24 _dafny.Int, f25 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
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
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this).F20 = f20
		(_this)._f21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Set {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Array {
	{
		return _this._f4
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
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.CodePoint {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Array {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Int {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F21() bool {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F22() _dafny.Array {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() _dafny.Array {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() _dafny.Int {
	{
		return _this._f24
	}
}
func (_this *GlobalState) F25() _dafny.Int {
	{
		return _this._f25
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F26 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F26 = false
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

func (_this *C0) Ctor__(f26 bool) {
	{
		(_this).F26 = f26
	}
}
func (_this *C0) Fm1(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("yim")
	}
}
func (_this *C0) M0(p0 bool, p1 _dafny.Int, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _178_v0 _dafny.Sequence
		_ = _178_v0
		_178_v0 = _dafny.SeqOf(p1, p1)
		var _179_v1 bool
		_ = _179_v1
		_179_v1 = p0
		var _180_v2 _dafny.Sequence
		_ = _180_v2
		_180_v2 = _dafny.SeqOf(_178_v0, _178_v0, Companion_Default___.Fm2((_179_v1), globalState))
		var _181_v3 _dafny.Map
		_ = _181_v3
		_181_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cmp(p1) < 0, (_180_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(532))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_182_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))).Cardinality())), _dafny.IntOfUint32((_180_v2).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _183_v4 _dafny.Array
		_ = _183_v4
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
		_ = _nw35
		_183_v4 = _nw35
		var _184_v5 _dafny.Sequence
		_ = _184_v5
		_184_v5 = _dafny.UnicodeSeqOfUtf8Bytes("ylwodeou")
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_183_v4), 0))
		_ = _index29
		(_183_v4).ArraySet1(_184_v5, (_index29).Int())
		var _185_v6 _dafny.Sequence
		_ = _185_v6
		_185_v6 = _dafny.SeqOf(_this.F26, !(_this.F26), _this.F26)
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_183_v4), 0))
		_ = _index30
		var _rhs29 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_178_v0, _dafny.Companion_Sequence_.Concatenate(_178_v0, _178_v0))).Cardinality())
		_ = _rhs29
		var _rhs30 bool = (_185_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_185_v6).Cardinality()))).Uint32()).(bool)
		_ = _rhs30
		var _rhs31 _dafny.Map = (_181_v3).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F26, _178_v0))
		_ = _rhs31
		var _rhs32 _dafny.Sequence = _184_v5
		_ = _rhs32
		var _lhs25 *GlobalState = globalState
		_ = _lhs25
		var _lhs26 _dafny.Array = _183_v4
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_183_v4), 0))
		_ = _lhs27
		_lhs25.F15 = _rhs29
		r0 = _rhs30
		_181_v3 = _rhs31
		(_lhs26).ArraySet1(_rhs32, (_lhs27).Int())
		r0 = (_179_v1)
		var _hi1 _dafny.Int = p1
		_ = _hi1
		for _186_i1 := (_dafny.Zero).Minus(p1); _186_i1.Cmp(_hi1) < 0; _186_i1 = _186_i1.Plus(_dafny.One) {
			var _187_v7 _dafny.Array
			_ = _187_v7
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw36
			_187_v7 = _nw36
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_187_v7), 0))
			_ = _index31
			(_187_v7).ArraySet1(_186_i1, (_index31).Int())
			var _188_v8 _dafny.Map
			_ = _188_v8
			_188_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _189_v9 _dafny.Map
			_ = _189_v9
			_189_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_i1, _this.F26)
			var _190_v10 _dafny.MultiSet
			_ = _190_v10
			_190_v10 = _dafny.MultiSetOf(_186_i1)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_187_v7), 0))
			_ = _index32
			var _rhs33 _dafny.Array = _187_v7
			_ = _rhs33
			var _rhs34 bool = !(false)
			_ = _rhs34
			var _rhs35 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, Companion_Default___.SafeDivisionInt(p1, ((_188_v8).Update(true, p1)).Cardinality()))
			_ = _rhs35
			var _rhs36 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(848))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_191_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			}))
			_ = _rhs36
			var _rhs37 _dafny.Int = (func() _dafny.Int {
				if (_188_v8).Contains((p1).Cmp(Companion_Default___.Fm3(_186_i1, _189_v9, _186_i1, globalState)) < 0) {
					return (_188_v8).Get((p1).Cmp(Companion_Default___.Fm3(_186_i1, _189_v9, _186_i1, globalState)) < 0).(_dafny.Int)
				}
				return (_190_v10).Cardinality()
			})()
			_ = _rhs37
			var _lhs28 *GlobalState = globalState
			_ = _lhs28
			var _lhs29 _dafny.Array = _187_v7
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_187_v7), 0))
			_ = _lhs30
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			_lhs28.F5 = _rhs33
			r0 = _rhs34
			(_lhs29).ArraySet1(_rhs35, (_lhs30).Int())
			_184_v5 = _rhs36
			_lhs31.F19 = _rhs37
			(_this).F26 = _this.F26
			var _192_v11 _dafny.Map
			_ = _192_v11
			_192_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if p0 {
					return p0
				}
				return p0
			})(), _this.F26)
			_192_v11 = (_192_v11).Update((func() bool {
				if p0 {
					return p0
				}
				return !(_this.F26)
			})(), _this.F26)
			var _193_v12 _dafny.Map
			_ = _193_v12
			_193_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_i1, _187_v7)
			var _rhs38 bool = _this.F26
			_ = _rhs38
			var _rhs39 bool = (_185_v6).Select((Companion_Default___.SafeIndex((p1).Times((_189_v9).Cardinality()), _dafny.IntOfUint32((_185_v6).Cardinality()))).Uint32()).(bool)
			_ = _rhs39
			var _rhs40 _dafny.Int = (_186_i1).Times((_193_v12).Cardinality())
			_ = _rhs40
			var _rhs41 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_178_v0).Cardinality()))
			_ = _rhs41
			var _rhs42 bool = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F26, _this.F26)).Cardinality()).Times((_187_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_187_v7), 0))).Int()).(_dafny.Int))).Cmp((_187_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_187_v7), 0))).Int()).(_dafny.Int)) >= 0
			_ = _rhs42
			var _lhs32 *GlobalState = globalState
			_ = _lhs32
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			var _lhs34 *C0 = _this
			_ = _lhs34
			r0 = _rhs38
			r0 = _rhs39
			_lhs32.F15 = _rhs40
			_lhs33.F15 = _rhs41
			_lhs34.F26 = _rhs42
		}
		var _194_v13 _dafny.Map
		_ = _194_v13
		_194_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F26, p1)
		var _195_v14 _dafny.Set
		_ = _195_v14
		_195_v14 = _dafny.SetOf(p0, _this.F26, p0)
		var _196_v15 _dafny.Map
		_ = _196_v15
		_196_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(456))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_197_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()), _this.F26)
		_194_v13 = ((_194_v13).Merge(_194_v13)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_178_v0, _dafny.IntOfInt64(693), (_195_v14).Cardinality(), globalState), Companion_Default___.Fm3(p1, _196_v15, p1, globalState)))
		(globalState).F19 = Companion_Default___.Fm3((Companion_Default___.Fm4((_dafny.Zero).Minus(p1), _this.F26, _this.F26, (_dafny.MultiSetOf(p1, p1)).Cardinality(), globalState)).Cardinality(), _196_v15, (_dafny.Zero).Minus((p1).Plus(p1)), globalState)
		var _198_v16 _dafny.MultiSet
		_ = _198_v16
		_198_v16 = _dafny.MultiSetOf(_this.F26, false, _this.F26, _this.F26, _this.F26)
		var _199_v17 _dafny.Map
		_ = _199_v17
		_199_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _200_v18 D1
		_ = _200_v18
		_200_v18 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32(((_183_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_183_v4), 0))).Int()).(_dafny.Sequence)).Cardinality()), p1, _dafny.IntOfInt64(-594), _199_v17)
		var _201_v19 _dafny.Array
		_ = _201_v19
		var _nwElement0_11 bool = _this.F26
		_ = _nwElement0_11
		var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(19))
		_ = _nw37
		(_nw37).ArraySet1(_nwElement0_11, 0)
		(_nw37).ArraySet1(false, 1)
		(_nw37).ArraySet1(_this.F26, 2)
		(_nw37).ArraySet1(p0, 3)
		(_nw37).ArraySet1(false, 4)
		(_nw37).ArraySet1(true, 5)
		(_nw37).ArraySet1(p0, 6)
		(_nw37).ArraySet1(true, 7)
		(_nw37).ArraySet1(_this.F26, 8)
		(_nw37).ArraySet1(p0, 9)
		(_nw37).ArraySet1(_this.F26, 10)
		(_nw37).ArraySet1(true, 11)
		(_nw37).ArraySet1(p0, 12)
		(_nw37).ArraySet1((p0), 13)
		(_nw37).ArraySet1(p0, 14)
		(_nw37).ArraySet1(_this.F26, 15)
		(_nw37).ArraySet1(_this.F26, 16)
		(_nw37).ArraySet1(p0, 17)
		(_nw37).ArraySet1(true, 18)
		_201_v19 = _nw37
		var _202_v20 _dafny.Map
		_ = _202_v20
		_202_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v18, _201_v19)
		var _203_v21 _dafny.Map
		_ = _203_v21
		_203_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(p1, _196_v15, (_198_v16).Cardinality(), globalState), (func() _dafny.Array {
			if (_202_v20).Contains(_200_v18) {
				return (_202_v20).Get(_200_v18).(_dafny.Array)
			}
			return _201_v19
		})())
		_203_v21 = _203_v21
		var _204_v22 _dafny.MultiSet
		_ = _204_v22
		_204_v22 = _dafny.MultiSetOf(p1)
		var _205_v23 _dafny.CodePoint
		_ = _205_v23
		_205_v23 = _dafny.CodePoint('c')
		var _206_v24 _dafny.Map
		_ = _206_v24
		_206_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_204_v22, _205_v23)
		r0 = !((_206_v24).Merge(_206_v24)).Equals(_206_v24)
		var _207_v25 _dafny.Sequence
		_ = _207_v25
		_207_v25 = _dafny.SeqOf(_204_v22, (_204_v22).Difference(Companion_Default___.Fm5(Companion_Default___.Fm0(_178_v0, p1, p1, globalState), _this.F26, _this.F26, globalState)))
		r1 = _207_v25
		return r0, r1
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
