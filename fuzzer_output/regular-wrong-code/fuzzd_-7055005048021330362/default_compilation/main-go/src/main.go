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
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.IntOfInt64(104)).Cmp((_dafny.SetOf(_dafny.IntOfInt64(-395))).Cardinality()) == 0 {
		return _dafny.CodePoint('m')
	} else {
		return _dafny.CodePoint('p')
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(33))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i1 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
		})), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))))).Cardinality())
	}))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nasbdl")).Cardinality())).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('h'), _dafny.CodePoint('k'))).Cardinality()))), ((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-139), _dafny.IntOfInt64(404))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-139)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(404)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_2_v0, _dafny.IntOfInt64(153)), func() _dafny.Map {
					var _coll1 = _dafny.NewMapBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(741))).Elements()); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _3_v1 _dafny.Int
						_3_v1 = interface{}(_compr_1).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(741)), _3_v1) {
							_coll1.Add(Companion_Default___.SafeDivisionInt(_3_v1, _dafny.IntOfInt64(293)), !(true))
						}
					}
					return _coll1.ToMap()
				}())
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("u")).Cardinality())))).Cardinality(), func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(22), _dafny.IntOfInt64(203))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v2 _dafny.Int
			_4_v2 = interface{}(_compr_2).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(22), _dafny.IntOfInt64(203))).Contains(_4_v2) {
				_coll2.Add((_4_v2).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), false)
			}
		}
		return _coll2.ToMap()
	}()))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-545))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_5_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(true, false, false)
		})), _dafny.SeqOf(_dafny.SeqOf(true, true, false)))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _6_v0 _dafny.Sequence
			_6_v0 = interface{}(_compr_3).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-545))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_5_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(true, false, false)
			})), _dafny.SeqOf(_dafny.SeqOf(true, true, false))), _6_v0) {
				_coll3.Add(_6_v0, !((!(false)) == (true)))
			}
		}
		return _coll3.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_(_dafny.SetOf(_dafny.SetOf(_dafny.CodePoint('x'), _dafny.CodePoint('e'))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(555))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 _dafny.MultiSet, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('h')), _dafny.SeqOf(_dafny.CodePoint('k')))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jlwcc"))).Difference((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("dulq"), _dafny.UnicodeSeqOfUtf8Bytes("ouob"), _dafny.UnicodeSeqOfUtf8Bytes("wntsnxr"))).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("oikjsmub"))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-141), _dafny.IntOfInt64(493))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _8_v0 _dafny.Int
			_8_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(-141)).Cmp(_8_v0) <= 0) && ((_8_v0).Cmp(_dafny.IntOfInt64(493)) < 0) {
				_coll4.Add((_8_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false, true), _dafny.SeqOf(true))).Cardinality())))
			}
		}
		return _coll4.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.CodePoint('a'), _dafny.CodePoint('a'), _dafny.CodePoint('r'), _dafny.CodePoint('a'))).Union((_dafny.SetOf(_dafny.CodePoint('j'))).Difference(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('r'))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _9_v0 _dafny.CodePoint
			_9_v0 = interface{}(_compr_5).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('r'))).Contains(_9_v0) {
				_coll5.Add(_9_v0)
			}
		}
		return _coll5.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Set, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()), (_dafny.MultiSetOf(_dafny.SetOf(!(false)))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(!((Companion_D9_.Create_DC20_(_dafny.UnicodeSeqOfUtf8Bytes("lt"), false, _dafny.IntOfInt64(130))).Dtor_cf38()), false, true, false)).Difference((_dafny.MultiSetOf(!(true))).Union((Companion_D3_.Create_DC7_(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true), true)))).Dtor_cf12()))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rkakqfc"), _dafny.UnicodeSeqOfUtf8Bytes("uhejkofsg")), _dafny.IntOfUint32(((func() _dafny.Sequence {
		if false {
			return _dafny.UnicodeSeqOfUtf8Bytes("olxnmwt")
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(461))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_10_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))
	})()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-997)), Companion_D3_.Create_DC7_(_dafny.MultiSetOf(!(true), true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-629)), Companion_D3_.Create_DC7_(_dafny.MultiSetOf(true))))).Keys().Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _11_v0 _dafny.MultiSet
			_11_v0 = interface{}(_compr_6).(_dafny.MultiSet)
			if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-997)), Companion_D3_.Create_DC7_(_dafny.MultiSetOf(!(true), true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-629)), Companion_D3_.Create_DC7_(_dafny.MultiSetOf(true))))).Contains(_11_v0) {
				_coll6.Add(_11_v0, ((_dafny.Zero).Minus(_dafny.IntOfInt64(-516))).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bsjog")).Cardinality())) <= 0)
			}
		}
		return _coll6.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Array, p1 bool, globalState *GlobalState) bool {
	var r0 bool = false
	_ = r0
	var _12_v0 _dafny.Map
	_ = _12_v0
	_12_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-251), p1)
	var _13_v1 _dafny.Int
	_ = _13_v1
	_13_v1 = _dafny.IntOfInt64(-477)
	(globalState).F3 = ((_12_v0).Cardinality()).Minus(_13_v1)
	var _14_v2 _dafny.Sequence
	_ = _14_v2
	_14_v2 = _dafny.UnicodeSeqOfUtf8Bytes("wqmiaea")
	var _15_v3 _dafny.Map
	_ = _15_v3
	_15_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v2, _13_v1)
	var _16_v4 _dafny.Sequence
	_ = _16_v4
	_16_v4 = _dafny.SeqOf(_dafny.IntOfInt64(382), _13_v1)
	_15_v3 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v2, _13_v1)).Merge(_15_v3)).Merge((Companion_Default___.Fm18(_dafny.IntOfInt64(878), _16_v4, _13_v1, _13_v1, globalState)).Merge(_15_v3))
	var _17_v5 _dafny.MultiSet
	_ = _17_v5
	_17_v5 = _dafny.MultiSetOf((_16_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-426), _dafny.IntOfUint32((_16_v4).Cardinality()))).Uint32()).(_dafny.Int))
	var _18_v6 _dafny.CodePoint
	_ = _18_v6
	_18_v6 = _dafny.CodePoint('t')
	var _hi0 _dafny.Int = (func() _dafny.Int {
		if p1 {
			return _dafny.IntOfUint32((_14_v2).Cardinality())
		}
		return _13_v1
	})()
	_ = _hi0
	for _19_i0 := (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v5).Cardinality(), _18_v6)).Cardinality()).Times(_dafny.IntOfInt64(582))); _19_i0.Cmp(_hi0) < 0; _19_i0 = _19_i0.Plus(_dafny.One) {
		if p1 {
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))
			_ = _index0
			(p0).ArraySet1(!((p1) == (p1)), (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))
			_ = _index1
			(p0).ArraySet1((p1) || (!(((Companion_Default___.Fm19(globalState)).Cardinality()).Cmp(_19_i0) != 0)), (_index1).Int())
			var _20_v7 _dafny.Set
			_ = _20_v7
			_20_v7 = _dafny.SetOf(_14_v2, _14_v2, _14_v2, _14_v2)
			var _21_v8 *C0
			_ = _21_v8
			var _nw0 *C0 = New_C0_()
			_ = _nw0
			_nw0.Ctor__(_20_v7)
			_21_v8 = _nw0
			var _22_v9 _dafny.Sequence
			_ = _22_v9
			_22_v9 = _dafny.SeqOf(_21_v8)
			var _23_v10 _dafny.MultiSet
			_ = _23_v10
			_23_v10 = _dafny.MultiSetOf(_21_v8, _21_v8)
			var _24_v11 *C3
			_ = _24_v11
			var _nw1 *C3 = New_C3_()
			_ = _nw1
			_nw1.Ctor__((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))).Int()).(bool), p1, _14_v2, p1)
			_24_v11 = _nw1
			var _25_v12 D7
			_ = _25_v12
			_25_v12 = Companion_D7_.Create_DC16_(p1, false, _24_v11, (_24_v11).F27())
			var _26_v13 _dafny.Map
			_ = _26_v13
			_26_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v12, _18_v6)
			var _27_v14 _dafny.MultiSet
			_ = _27_v14
			_27_v14 = _dafny.MultiSetOf((_24_v11).F27())
			var _28_v15 _dafny.Array
			_ = _28_v15
			var _nw2 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw2
			_28_v15 = _nw2
			var _29_v16 _dafny.Map
			_ = _29_v16
			_29_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_16_v4).Select((Companion_Default___.SafeIndex(_13_v1, _dafny.IntOfUint32((_16_v4).Cardinality()))).Uint32()).(_dafny.Int), _28_v15)
			var _30_v17 _dafny.Map
			_ = _30_v17
			_30_v17 = _29_v16
			var _31_v18 _dafny.Map
			_ = _31_v18
			_31_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v1, _13_v1)
			var _32_v19 _dafny.Array
			_ = _32_v19
			var _nwElement0_0 _dafny.Int = (_13_v1).Plus((_dafny.Zero).Minus(_19_i0))
			_ = _nwElement0_0
			var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(20))
			_ = _nw3
			(_nw3).ArraySet1(_nwElement0_0, 0)
			(_nw3).ArraySet1(_19_i0, 1)
			(_nw3).ArraySet1((_13_v1).Plus(_19_i0), 2)
			(_nw3).ArraySet1(_19_i0, 3)
			(_nw3).ArraySet1(Companion_Default___.Fm6(globalState), 4)
			(_nw3).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_22_v9, _22_v9)).Cardinality()), 5)
			(_nw3).ArraySet1((func() _dafny.Int {
				if (_23_v10).Contains(_21_v8) {
					return (_23_v10).Multiplicity(_21_v8)
				}
				return _13_v1
			})(), 6)
			(_nw3).ArraySet1(_19_i0, 7)
			(_nw3).ArraySet1(_13_v1, 8)
			(_nw3).ArraySet1((_13_v1).Minus((_26_v13).Cardinality()), 9)
			(_nw3).ArraySet1(_dafny.IntOfInt64(628), 10)
			(_nw3).ArraySet1((func() _dafny.Int {
				if (_17_v5).Contains(_13_v1) {
					return (_17_v5).Multiplicity(_13_v1)
				}
				return (_27_v14).Cardinality()
			})(), 11)
			(_nw3).ArraySet1(((_29_v16).Merge((_30_v17))).Cardinality(), 12)
			(_nw3).ArraySet1(_19_i0, 13)
			(_nw3).ArraySet1(_13_v1, 14)
			(_nw3).ArraySet1(_13_v1, 15)
			(_nw3).ArraySet1(Companion_Default___.SafeDivisionInt(_13_v1, _19_i0), 16)
			(_nw3).ArraySet1((func() _dafny.Int {
				if (_31_v18).Contains(_13_v1) {
					return (_31_v18).Get(_13_v1).(_dafny.Int)
				}
				return _13_v1
			})(), 17)
			(_nw3).ArraySet1((_13_v1).Times((_21_v8.F26).Cardinality()), 18)
			(_nw3).ArraySet1(_19_i0, 19)
			_32_v19 = _nw3
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_32_v19), 0))
			_ = _index2
			(_32_v19).ArraySet1(_19_i0, (_index2).Int())
			var _33_v20 _dafny.Array
			_ = _33_v20
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(9))
			_ = _nw4
			_33_v20 = _nw4
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_33_v20), 0))
			_ = _index3
			(_33_v20).ArraySet1CodePoint(_18_v6, (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_32_v19), 0))
			_ = _index4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_33_v20), 0))
			_ = _index5
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))
			_ = _index6
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))
			_ = _index7
			var _rhs0 _dafny.Int = (_19_i0).Plus((func() _dafny.Int {
				if _24_v11.F28 {
					return _13_v1
				}
				return _dafny.IntOfInt64(468)
			})())
			_ = _rhs0
			var _rhs1 _dafny.Int = Companion_Default___.SafeModuloInt(_19_i0, _dafny.IntOfUint32((_14_v2).Cardinality()))
			_ = _rhs1
			var _rhs2 _dafny.CodePoint = _18_v6
			_ = _rhs2
			var _rhs3 bool = !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_14_v2, _14_v2), (Companion_Default___.SafeIndex(_13_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_14_v2, _14_v2)).Cardinality()))).Uint32(), _18_v6), _14_v2))
			_ = _rhs3
			var _rhs4 bool = (_24_v11).F27()
			_ = _rhs4
			var _lhs0 _dafny.Array = _32_v19
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_32_v19), 0))
			_ = _lhs1
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			var _lhs3 _dafny.Array = _33_v20
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_33_v20), 0))
			_ = _lhs4
			var _lhs5 _dafny.Array = p0
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))
			_ = _lhs6
			var _lhs7 _dafny.Array = p0
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))
			_ = _lhs8
			(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
			_lhs2.F20 = _rhs1
			(_lhs3).ArraySet1CodePoint(_rhs2, (_lhs4).Int())
			(_lhs5).ArraySet1(_rhs3, (_lhs6).Int())
			(_lhs7).ArraySet1(_rhs4, (_lhs8).Int())
			var _34_v21 T0
			_ = _34_v21
			var _nw5 *C3 = New_C3_()
			_ = _nw5
			_nw5.Ctor__(p1, true, _14_v2, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p0), 0))).Int()).(bool))
			_34_v21 = _nw5
			var _35_v22 _dafny.Map
			_ = _35_v22
			_35_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_34_v21, p1)
			var _36_v23 _dafny.Sequence
			_ = _36_v23
			_36_v23 = _dafny.SeqOf(_34_v21.F22())
			var _37_v24 _dafny.Map
			_ = _37_v24
			_37_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v11.F28, _13_v1)
			var _38_v25 D1
			_ = _38_v25
			_38_v25 = Companion_D1_.Create_DC5_(_34_v21.F22(), _37_v24, _13_v1, (_dafny.MultiSetOf(_dafny.IntOfInt64(379))).Update((_32_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_32_v19), 0))).Int()).(_dafny.Int), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(627))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_39_v20 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_40_i1 _dafny.Int) _dafny.CodePoint {
					return (_39_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_39_v20), 0))).Int())
				}
			})(_33_v20)))).Cardinality()))))
			var _41_v26 D3
			_ = _41_v26
			_41_v26 = Companion_D3_.Create_DC8_(_dafny.IntOfUint32((_36_v23).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v14, _dafny.IntOfUint32((_36_v23).Cardinality())), _38_v25, p1)
			var _42_v27 D1
			_ = _42_v27
			_42_v27 = Companion_D1_.Create_DC5_((_41_v26).Dtor_cf16(), _37_v24, (_dafny.Zero).Minus(_13_v1), _dafny.MultiSetOf(_19_i0, (_32_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_32_v19), 0))).Int()).(_dafny.Int)))
			var _43_v28 _dafny.Sequence
			_ = _43_v28
			_43_v28 = _dafny.SeqOf(_16_v4)
			var _44_v29 T2
			_ = _44_v29
			var _nw6 *C2 = New_C2_()
			_ = _nw6
			_nw6.Ctor__(_43_v28, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(604))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_45_v20 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_46_i2 _dafny.Int) _dafny.CodePoint {
					return (_45_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_45_v20), 0))).Int())
				}
			})(_33_v20))), _24_v11.F28)
			_44_v29 = _nw6
			var _47_v30 _dafny.Sequence
			_ = _47_v30
			_47_v30 = _dafny.SeqOf(_44_v29, _44_v29)
			var _48_v31 _dafny.Map
			_ = _48_v31
			_48_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_35_v22).Equals(_35_v22), (func() _dafny.Sequence {
				if (_42_v27).Dtor_cf7() {
					return _47_v30
				}
				return _47_v30
			})())
			_48_v31 = (_48_v31).Update(!(_34_v21.F22()), _dafny.Companion_Sequence_.Concatenate(_47_v30, _47_v30))
			_44_v29 = _44_v29
			var _49_v32 _dafny.Set
			_ = _49_v32
			_49_v32 = _dafny.SetOf(_44_v29)
			_49_v32 = (_49_v32).Union(_49_v32)
		} else {
			var _50_v33 _dafny.MultiSet
			_ = _50_v33
			_50_v33 = _dafny.MultiSetOf(p1, false)
			var _51_v34 _dafny.MultiSet
			_ = _51_v34
			_51_v34 = _dafny.MultiSetOf(_50_v33, _50_v33, _50_v33, _50_v33, _50_v33)
			var _52_v35 _dafny.Sequence
			_ = _52_v35
			_52_v35 = _dafny.SeqOf(Companion_Default___.Fm12(_13_v1, Companion_Default___.Fm17(_13_v1, Companion_Default___.Fm12(_13_v1, _50_v33, p1, _51_v34, globalState), globalState), p1, _51_v34, globalState))
			r0 = (_52_v35).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm6(globalState), _dafny.IntOfUint32((_52_v35).Cardinality()))).Uint32()).(bool)
			r0 = _dafny.Companion_Sequence_.Equal(_14_v2, _14_v2)
			var _53_v36 _dafny.Array
			_ = _53_v36
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_0
			var _nw7 _dafny.Array
			_ = _nw7
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw7 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.MultiSet = (func(_54_v33 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_55_i3 _dafny.Int) _dafny.MultiSet {
						return _54_v33
					}
				})(_50_v33)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw7 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw7).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw7).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_53_v36 = _nw7
			var _56_v37 *C1
			_ = _56_v37
			var _nw8 *C1 = New_C1_()
			_ = _nw8
			_nw8.Ctor__(_53_v36, p1)
			_56_v37 = _nw8
			_56_v37 = _56_v37
			var _57_v38 _dafny.Set
			_ = _57_v38
			_57_v38 = _dafny.SetOf((_56_v37).F25())
			var _58_v39 _dafny.Set
			_ = _58_v39
			_58_v39 = _dafny.SetOf(_57_v38, _dafny.SetOf(true, (_56_v37).F25()))
			var _59_v40 _dafny.Map
			_ = _59_v40
			_59_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D9_.Create_DC18_(_58_v39)).Dtor_cf31(), p1)
			_59_v40 = (_59_v40).Update(_58_v39, p1)
			var _60_v41 _dafny.Map
			_ = _60_v41
			_60_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_i0, _18_v6)
			var _61_v42 _dafny.Sequence
			_ = _61_v42
			_61_v42 = _dafny.SeqOf(_50_v33, _50_v33)
			r0 = _dafny.Companion_Sequence_.Contains(_14_v2, (func() _dafny.CodePoint {
				if (_60_v41).Contains(((_61_v42).Select((Companion_Default___.SafeIndex(_13_v1, _dafny.IntOfUint32((_61_v42).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()) {
					return (_60_v41).Get(((_61_v42).Select((Companion_Default___.SafeIndex(_13_v1, _dafny.IntOfUint32((_61_v42).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()).(_dafny.CodePoint)
				}
				return _18_v6
			})())
		}
		var _62_v43 _dafny.Map
		_ = _62_v43
		_62_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _13_v1)
		_13_v1 = (_62_v43).Cardinality()
		_13_v1 = _dafny.IntOfInt64(-286)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((p0), 0))
		_ = _index8
		(p0).ArraySet1(p1, (_index8).Int())
		var _63_v44 _dafny.Array
		_ = _63_v44
		var _nw9 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw9
		_63_v44 = _nw9
		var _64_v45 _dafny.Sequence
		_ = _64_v45
		_64_v45 = _dafny.SeqOf(p0)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((p0), 0))
		_ = _index9
		var _rhs5 bool = _dafny.Companion_Sequence_.Contains(_16_v4, (_13_v1).Minus(_19_i0))
		_ = _rhs5
		var _rhs6 _dafny.Array = (func() _dafny.Array {
			if !(true) {
				return p0
			}
			return (_64_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.IntOfUint32((_64_v45).Cardinality()))).Uint32()).(_dafny.Array)
		})()
		_ = _rhs6
		var _lhs9 _dafny.Array = p0
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((p0), 0))
		_ = _lhs10
		(_lhs9).ArraySet1(_rhs5, (_lhs10).Int())
		_63_v44 = _rhs6
	}
	var _65_v46 _dafny.MultiSet
	_ = _65_v46
	_65_v46 = _dafny.MultiSetOf(p1)
	var _66_v47 _dafny.Map
	_ = _66_v47
	_66_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v46, Companion_Default___.Fm6(globalState))
	var _67_v48 D1
	_ = _67_v48
	_67_v48 = Companion_D1_.Create_DC5_(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus(_13_v1)), _13_v1, _17_v5)
	var _pat_let_tv0 = _16_v4
	_ = _pat_let_tv0
	var _pat_let_tv1 = _16_v4
	_ = _pat_let_tv1
	var _pat_let_tv2 = _16_v4
	_ = _pat_let_tv2
	var _rhs7 _dafny.Sequence = func(_source0 D3) _dafny.Sequence {
		if _source0.Is_DC8() {
			var _68___mcc_h0 _dafny.Int = _source0.Get_().(D3_DC8).Cf13
			_ = _68___mcc_h0
			var _69___mcc_h1 _dafny.Map = _source0.Get_().(D3_DC8).Cf14
			_ = _69___mcc_h1
			var _70___mcc_h2 D1 = _source0.Get_().(D3_DC8).Cf15
			_ = _70___mcc_h2
			var _71___mcc_h3 bool = _source0.Get_().(D3_DC8).Cf16
			_ = _71___mcc_h3
			var _72_cf16 bool = _71___mcc_h3
			_ = _72_cf16
			var _73_cf15 D1 = _70___mcc_h2
			_ = _73_cf15
			var _74_cf14 _dafny.Map = _69___mcc_h1
			_ = _74_cf14
			var _75_cf13 _dafny.Int = _68___mcc_h0
			_ = _75_cf13
			return _pat_let_tv0
		} else if _source0.Is_DC9() {
			var _76___mcc_h4 _dafny.Int = _source0.Get_().(D3_DC9).Cf17
			_ = _76___mcc_h4
			var _77___mcc_h5 _dafny.Int = _source0.Get_().(D3_DC9).Cf18
			_ = _77___mcc_h5
			var _78___mcc_h6 _dafny.Int = _source0.Get_().(D3_DC9).Cf19
			_ = _78___mcc_h6
			var _79_cf19 _dafny.Int = _78___mcc_h6
			_ = _79_cf19
			var _80_cf18 _dafny.Int = _77___mcc_h5
			_ = _80_cf18
			var _81_cf17 _dafny.Int = _76___mcc_h4
			_ = _81_cf17
			return _pat_let_tv1
		} else {
			var _82___mcc_h7 _dafny.MultiSet = _source0.Get_().(D3_DC7).Cf12
			_ = _82___mcc_h7
			var _83_cf12 _dafny.MultiSet = _82___mcc_h7
			_ = _83_cf12
			return _pat_let_tv2
		}
	}(Companion_D3_.Create_DC8_(_13_v1, _66_v47, _67_v48, !(p1)))
	_ = _rhs7
	var _rhs8 bool = p1
	_ = _rhs8
	var _rhs9 bool = p1
	_ = _rhs9
	_16_v4 = _rhs7
	r0 = _rhs8
	r0 = _rhs9
	var _84_v49 _dafny.Set
	_ = _84_v49
	_84_v49 = _dafny.SetOf((_12_v0).Cardinality())
	var _85_v50 _dafny.Set
	_ = _85_v50
	_85_v50 = _84_v49
	var _86_v51 _dafny.Sequence
	_ = _86_v51
	_86_v51 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_16_v4, (Companion_Default___.SafeIndex(_13_v1, _dafny.IntOfUint32((_16_v4).Cardinality()))).Uint32(), _13_v1), _16_v4, _16_v4)
	var _87_v52 T1
	_ = _87_v52
	var _nw10 *C4 = New_C4_()
	_ = _nw10
	_nw10.Ctor__(_85_v50, p1, _86_v51, _14_v2, ((_84_v49).Cardinality()).Cmp(_13_v1) == 0)
	_87_v52 = _nw10
	var _88_v53 _dafny.Set
	_ = _88_v53
	_88_v53 = _dafny.SetOf(false)
	var _89_v54 _dafny.Map
	_ = _89_v54
	_89_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v53, p1)
	var _90_v56 D9
	_ = _90_v56
	_90_v56 = Companion_D9_.Create_DC18_(func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_89_v54).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _91_v55 _dafny.Set
			_91_v55 = interface{}(_compr_7).(_dafny.Set)
			if (_89_v54).Contains(_91_v55) {
				_coll7.Add(_91_v55)
			}
		}
		return _coll7.ToSet()
	}())
	var _92_v57 _dafny.Map
	_ = _92_v57
	_92_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
	var _93_v58 D9
	_ = _93_v58
	_93_v58 = Companion_D9_.Create_DC20_(_87_v52.F21(), p1, (_92_v57).Cardinality())
	var _94_v59 D9
	_ = _94_v59
	_94_v59 = Companion_D9_.Create_DC20_((_93_v58).Dtor_cf37(), true, _13_v1)
	var _pat_let_tv3 = _13_v1
	_ = _pat_let_tv3
	var _pat_let_tv4 = _67_v48
	_ = _pat_let_tv4
	var _pat_let_tv5 = _13_v1
	_ = _pat_let_tv5
	var _pat_let_tv6 = _13_v1
	_ = _pat_let_tv6
	var _pat_let_tv7 = _65_v46
	_ = _pat_let_tv7
	var _pat_let_tv8 = _65_v46
	_ = _pat_let_tv8
	var _pat_let_tv9 = p1
	_ = _pat_let_tv9
	var _pat_let_tv10 = p1
	_ = _pat_let_tv10
	var _pat_let_tv11 = p1
	_ = _pat_let_tv11
	var _pat_let_tv12 = _65_v46
	_ = _pat_let_tv12
	var _pat_let_tv13 = _65_v46
	_ = _pat_let_tv13
	var _rhs10 _dafny.Int = func(_source1 D9) _dafny.Int {
		if _source1.Is_DC19() {
			var _95___mcc_h8 _dafny.Int = _source1.Get_().(D9_DC19).Cf32
			_ = _95___mcc_h8
			var _96___mcc_h9 _dafny.Int = _source1.Get_().(D9_DC19).Cf33
			_ = _96___mcc_h9
			var _97___mcc_h10 bool = _source1.Get_().(D9_DC19).Cf34
			_ = _97___mcc_h10
			var _98___mcc_h11 _dafny.Int = _source1.Get_().(D9_DC19).Cf35
			_ = _98___mcc_h11
			var _99___mcc_h12 _dafny.Set = _source1.Get_().(D9_DC19).Cf36
			_ = _99___mcc_h12
			var _100_cf36 _dafny.Set = _99___mcc_h12
			_ = _100_cf36
			var _101_cf35 _dafny.Int = _98___mcc_h11
			_ = _101_cf35
			var _102_cf34 bool = _97___mcc_h10
			_ = _102_cf34
			var _103_cf33 _dafny.Int = _96___mcc_h9
			_ = _103_cf33
			var _104_cf32 _dafny.Int = _95___mcc_h8
			_ = _104_cf32
			return _pat_let_tv3
		} else if _source1.Is_DC20() {
			var _105___mcc_h13 _dafny.Sequence = _source1.Get_().(D9_DC20).Cf37
			_ = _105___mcc_h13
			var _106___mcc_h14 bool = _source1.Get_().(D9_DC20).Cf38
			_ = _106___mcc_h14
			var _107___mcc_h15 _dafny.Int = _source1.Get_().(D9_DC20).Cf39
			_ = _107___mcc_h15
			var _108_cf39 _dafny.Int = _107___mcc_h15
			_ = _108_cf39
			var _109_cf38 bool = _106___mcc_h14
			_ = _109_cf38
			var _110_cf37 _dafny.Sequence = _105___mcc_h13
			_ = _110_cf37
			return (((_pat_let_tv4).Dtor_cf8()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_cf38, _pat_let_tv5))).Cardinality()
		} else if _source1.Is_DC21() {
			var _111___mcc_h16 _dafny.Array = _source1.Get_().(D9_DC21).Cf40
			_ = _111___mcc_h16
			var _112___mcc_h17 bool = _source1.Get_().(D9_DC21).Cf41
			_ = _112___mcc_h17
			var _113_cf41 bool = _112___mcc_h17
			_ = _113_cf41
			var _114_cf40 _dafny.Array = _111___mcc_h16
			_ = _114_cf40
			return _dafny.IntOfInt64(995)
		} else {
			var _115___mcc_h18 _dafny.Set = _source1.Get_().(D9_DC18).Cf31
			_ = _115___mcc_h18
			var _116_cf31 _dafny.Set = _115___mcc_h18
			_ = _116_cf31
			return _pat_let_tv6
		}
	}(_90_v56)
	_ = _rhs10
	var _rhs11 T1 = _87_v52
	_ = _rhs11
	var _rhs12 _dafny.MultiSet = func(_source2 D9) _dafny.MultiSet {
		if _source2.Is_DC19() {
			var _117___mcc_h19 _dafny.Int = _source2.Get_().(D9_DC19).Cf32
			_ = _117___mcc_h19
			var _118___mcc_h20 _dafny.Int = _source2.Get_().(D9_DC19).Cf33
			_ = _118___mcc_h20
			var _119___mcc_h21 bool = _source2.Get_().(D9_DC19).Cf34
			_ = _119___mcc_h21
			var _120___mcc_h22 _dafny.Int = _source2.Get_().(D9_DC19).Cf35
			_ = _120___mcc_h22
			var _121___mcc_h23 _dafny.Set = _source2.Get_().(D9_DC19).Cf36
			_ = _121___mcc_h23
			var _122_cf36 _dafny.Set = _121___mcc_h23
			_ = _122_cf36
			var _123_cf35 _dafny.Int = _120___mcc_h22
			_ = _123_cf35
			var _124_cf34 bool = _119___mcc_h21
			_ = _124_cf34
			var _125_cf33 _dafny.Int = _118___mcc_h20
			_ = _125_cf33
			var _126_cf32 _dafny.Int = _117___mcc_h19
			_ = _126_cf32
			return _pat_let_tv7
		} else if _source2.Is_DC20() {
			var _127___mcc_h24 _dafny.Sequence = _source2.Get_().(D9_DC20).Cf37
			_ = _127___mcc_h24
			var _128___mcc_h25 bool = _source2.Get_().(D9_DC20).Cf38
			_ = _128___mcc_h25
			var _129___mcc_h26 _dafny.Int = _source2.Get_().(D9_DC20).Cf39
			_ = _129___mcc_h26
			var _130_cf39 _dafny.Int = _129___mcc_h26
			_ = _130_cf39
			var _131_cf38 bool = _128___mcc_h25
			_ = _131_cf38
			var _132_cf37 _dafny.Sequence = _127___mcc_h24
			_ = _132_cf37
			return _pat_let_tv8
		} else if _source2.Is_DC21() {
			var _133___mcc_h27 _dafny.Array = _source2.Get_().(D9_DC21).Cf40
			_ = _133___mcc_h27
			var _134___mcc_h28 bool = _source2.Get_().(D9_DC21).Cf41
			_ = _134___mcc_h28
			var _135_cf41 bool = _134___mcc_h28
			_ = _135_cf41
			var _136_cf40 _dafny.Array = _133___mcc_h27
			_ = _136_cf40
			return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv9, _pat_let_tv10), _dafny.SeqOf(_135_cf41, !(_pat_let_tv11), false, true, _135_cf41)))
		} else {
			var _137___mcc_h29 _dafny.Set = _source2.Get_().(D9_DC18).Cf31
			_ = _137___mcc_h29
			var _138_cf31 _dafny.Set = _137___mcc_h29
			_ = _138_cf31
			return (_pat_let_tv12).Union(_pat_let_tv13)
		}
	}(_94_v59)
	_ = _rhs12
	var _lhs11 *GlobalState = globalState
	_ = _lhs11
	_lhs11.F3 = _rhs10
	_87_v52 = _rhs11
	_65_v46 = _rhs12
	(_87_v52).F22_set_(p1)
	var _139_v60 _dafny.Sequence
	_ = _139_v60
	_139_v60 = _dafny.SeqOf(_87_v52.F22(), _87_v52.F22(), _87_v52.F22(), _87_v52.F22(), true)
	r0 = (!(!(!_dafny.Companion_Sequence_.Contains(_139_v60, _87_v52.F22())))) && (_87_v52.F22())
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _140_v1 _dafny.Sequence
	_ = _140_v1
	_140_v1 = _dafny.UnicodeSeqOfUtf8Bytes("truxk")
	var _141_v2 _dafny.Set
	_ = _141_v2
	_141_v2 = _dafny.SetOf(_140_v1, _140_v1)
	var _142_v3 _dafny.Int
	_ = _142_v3
	_142_v3 = _dafny.IntOfInt64(-610)
	var _143_v4 _dafny.CodePoint
	_ = _143_v4
	_143_v4 = _dafny.CodePoint('g')
	var _144_v5 _dafny.Array
	_ = _144_v5
	var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
	_ = _nw11
	_144_v5 = _nw11
	var _145_v6 _dafny.Array
	_ = _145_v6
	var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
	_ = _nw12
	_145_v6 = _nw12
	var _146_v7 _dafny.MultiSet
	_ = _146_v7
	_146_v7 = _dafny.MultiSetOf(_142_v3)
	var _147_globalState *GlobalState
	_ = _147_globalState
	var _nw13 *GlobalState = New_GlobalState_()
	_ = _nw13
	_nw13.Ctor__(false, func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_141_v2).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _148_v0 _dafny.Sequence
			_148_v0 = interface{}(_compr_8).(_dafny.Sequence)
			if (_141_v2).Contains(_148_v0) {
				_coll8.Add(_148_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(281))).Cardinality()))
			}
		}
		return _coll8.ToMap()
	}(), _dafny.IntOfInt64(-658), _dafny.IntOfInt64(2), _dafny.IntOfInt64(-712), true, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("v"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_142_v3), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality()))).Uint32(), _143_v4), (Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("v"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_142_v3), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality()))).Uint32(), _143_v4)).Cardinality()))).Uint32(), _143_v4), true, _140_v1, _140_v1, _dafny.IntOfInt64(896), true, _dafny.CodePoint('c'), _dafny.IntOfInt64(-158), _dafny.IntOfInt64(316), _dafny.IntOfInt64(-654), _144_v5, _145_v6, false, _146_v7, _dafny.IntOfInt64(308))
	_147_globalState = _nw13
	var _149_v8 bool
	_ = _149_v8
	_149_v8 = false
	var _150_i0 _dafny.Int
	_ = _150_i0
	_150_i0 = _dafny.Zero
	{
		for !(!(_149_v8)) {
			{
				if (_150_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_150_i0 = (_150_i0).Plus(_dafny.One)
				var _151_v9 _dafny.Array
				_ = _151_v9
				var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw14
				_151_v9 = _nw14
				var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_151_v9), 0))
				_ = _index10
				(_151_v9).ArraySet1(_144_v5, (_index10).Int())
				var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_151_v9), 0))
				_ = _index11
				var _rhs13 _dafny.Int = (_142_v3).Minus((func() _dafny.Int {
					if _149_v8 {
						return _142_v3
					}
					return _142_v3
				})())
				_ = _rhs13
				var _rhs14 _dafny.Array = _144_v5
				_ = _rhs14
				var _rhs15 _dafny.Int = _142_v3
				_ = _rhs15
				var _lhs12 _dafny.Array = _151_v9
				_ = _lhs12
				var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_151_v9), 0))
				_ = _lhs13
				var _lhs14 *GlobalState = _147_globalState
				_ = _lhs14
				_142_v3 = _rhs13
				(_lhs12).ArraySet1(_rhs14, (_lhs13).Int())
				_lhs14.F20 = _rhs15
				(_147_globalState).F20 = _142_v3
				_142_v3 = _142_v3
				var _152_v10 bool
				_ = _152_v10
				var _out0 bool
				_ = _out0
				_out0 = Companion_Default___.M0(_144_v5, _149_v8, _147_globalState)
				_152_v10 = _out0
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	if _149_v8 {
		var _153_v11 _dafny.Array
		_ = _153_v11
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_1
		var _nw15 _dafny.Array
		_ = _nw15
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw15 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.MultiSet = func(_154_i1 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(true)
			}
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw15 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw15).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw15).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_153_v11 = _nw15
		var _155_v12 *C1
		_ = _155_v12
		var _nw16 *C1 = New_C1_()
		_ = _nw16
		_nw16.Ctor__(_153_v11, !(false) || (_149_v8))
		_155_v12 = _nw16
		var _156_v13 _dafny.MultiSet
		_ = _156_v13
		_156_v13 = _dafny.MultiSetOf(true)
		var _157_v14 _dafny.MultiSet
		_ = _157_v14
		_157_v14 = _dafny.MultiSetOf(_156_v13)
		var _158_v15 _dafny.Set
		_ = _158_v15
		_158_v15 = _dafny.SetOf((_155_v12).F25(), false)
		var _159_v16 *C2
		_ = _159_v16
		var _nw17 *C2 = New_C2_()
		_ = _nw17
		_nw17.Ctor__(Companion_Default___.Fm16(_158_v15, (_155_v12).F25(), _149_v8, _147_globalState), _140_v1, _149_v8)
		_159_v16 = _nw17
		var _160_v17 _dafny.Map
		_ = _160_v17
		_160_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12(_142_v3, _156_v13, (_155_v12).F25(), (_157_v14).Update(_156_v13, Companion_Default___.Abs(_dafny.IntOfUint32((_140_v1).Cardinality()))), _147_globalState), _159_v16)
		var _161_v18 _dafny.Map
		_ = _161_v18
		_161_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v1, _160_v17)
		var _162_v19 _dafny.Map
		_ = _162_v19
		_162_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v8, false)
		var _163_v20 _dafny.Sequence
		_ = _163_v20
		_163_v20 = _dafny.SeqOf(_162_v19, _162_v19, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v8, false))
		var _rhs16 _dafny.Int = _142_v3
		_ = _rhs16
		var _rhs17 _dafny.Map = _161_v18
		_ = _rhs17
		var _rhs18 _dafny.Int = Companion_Default___.SafeDivisionInt(_142_v3, _142_v3)
		_ = _rhs18
		var _rhs19 bool = !_dafny.Companion_Sequence_.Contains(_163_v20, _162_v19)
		_ = _rhs19
		var _lhs15 *GlobalState = _147_globalState
		_ = _lhs15
		var _lhs16 *GlobalState = _147_globalState
		_ = _lhs16
		_lhs15.F3 = _rhs16
		_161_v18 = _rhs17
		_lhs16.F3 = _rhs18
		_149_v8 = _rhs19
		var _164_v21 bool
		_ = _164_v21
		var _out1 bool
		_ = _out1
		_out1 = Companion_Default___.M0(_144_v5, !((_155_v12).F25()), _147_globalState)
		_164_v21 = _out1
		_149_v8 = ((_155_v12).F25()) || ((_155_v12).F25())
		_142_v3 = (func() _dafny.Int {
			if (_159_v16).Fm3(_147_globalState) {
				return _142_v3
			}
			return _142_v3
		})()
	} else {
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_145_v6), 0))
		_ = _index12
		(_145_v6).ArraySet1(_dafny.IntOfInt64(-385), (_index12).Int())
		var _165_v22 _dafny.Sequence
		_ = _165_v22
		_165_v22 = _dafny.SeqOf(_142_v3)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_145_v6), 0))
		_ = _index13
		(_145_v6).ArraySet1((_dafny.Zero).Minus(((_165_v22).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_165_v22).Cardinality()))).Uint32()).(_dafny.Int)).Times(_142_v3)), (_index13).Int())
		var _166_v23 _dafny.MultiSet
		_ = _166_v23
		_166_v23 = _dafny.MultiSetOf(_149_v8, _149_v8)
		var _167_v24 _dafny.Map
		_ = _167_v24
		_167_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_145_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_145_v6), 0))).Int()).(_dafny.Int), _149_v8)
		var _168_v25 _dafny.MultiSet
		_ = _168_v25
		_168_v25 = _dafny.MultiSetOf(_166_v23)
		var _169_v26 _dafny.Array
		_ = _169_v26
		var _nwElement0_1 _dafny.MultiSet = _166_v23
		_ = _nwElement0_1
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(27))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_1, 0)
		(_nw18).ArraySet1(_dafny.MultiSetOf(true), 1)
		(_nw18).ArraySet1(_166_v23, 2)
		(_nw18).ArraySet1(_166_v23, 3)
		(_nw18).ArraySet1(_166_v23, 4)
		(_nw18).ArraySet1(_166_v23, 5)
		(_nw18).ArraySet1(_166_v23, 6)
		(_nw18).ArraySet1(_166_v23, 7)
		(_nw18).ArraySet1(_166_v23, 8)
		(_nw18).ArraySet1(_166_v23, 9)
		(_nw18).ArraySet1(_166_v23, 10)
		(_nw18).ArraySet1(_166_v23, 11)
		(_nw18).ArraySet1(_166_v23, 12)
		(_nw18).ArraySet1(_dafny.MultiSetOf(_149_v8, _149_v8, _149_v8), 13)
		(_nw18).ArraySet1(_166_v23, 14)
		(_nw18).ArraySet1(_166_v23, 15)
		(_nw18).ArraySet1(_dafny.MultiSetOf(_149_v8), 16)
		(_nw18).ArraySet1(_dafny.MultiSetOf(_149_v8, _149_v8), 17)
		(_nw18).ArraySet1(_dafny.MultiSetOf(_149_v8, _149_v8), 18)
		(_nw18).ArraySet1(_dafny.MultiSetOf(_149_v8, _149_v8, Companion_Default___.Fm12(_142_v3, _dafny.MultiSetOf(_149_v8, true, (func() bool {
			if (_167_v24).Contains(_142_v3) {
				return (_167_v24).Get(_142_v3).(bool)
			}
			return _149_v8
		})()), !(Companion_Default___.Fm12(_142_v3, _dafny.MultiSetOf(_149_v8, _149_v8), _149_v8, _168_v25, _147_globalState)), _168_v25, _147_globalState)), 19)
		(_nw18).ArraySet1(_166_v23, 20)
		(_nw18).ArraySet1(Companion_Default___.Fm17((_145_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_145_v6), 0))).Int()).(_dafny.Int), _149_v8, _147_globalState), 21)
		(_nw18).ArraySet1(_dafny.MultiSetOf(_149_v8, _149_v8, _149_v8), 22)
		(_nw18).ArraySet1(_166_v23, 23)
		(_nw18).ArraySet1(_166_v23, 24)
		(_nw18).ArraySet1(_dafny.MultiSetOf(_149_v8), 25)
		(_nw18).ArraySet1(_166_v23, 26)
		_169_v26 = _nw18
		var _170_v27 _dafny.Map
		_ = _170_v27
		_170_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v8, _142_v3)).Cardinality(), _166_v23, _149_v8, _168_v25, _147_globalState), _142_v3)
		var _171_v28 *C1
		_ = _171_v28
		var _nw19 *C1 = New_C1_()
		_ = _nw19
		_nw19.Ctor__(_169_v26, ((_145_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_145_v6), 0))).Int()).(_dafny.Int)).Cmp((_170_v27).Cardinality()) >= 0)
		_171_v28 = _nw19
		var _172_v29 _dafny.Sequence
		_ = _172_v29
		_172_v29 = _dafny.SeqOf(true, _149_v8)
		var _173_v30 D3
		_ = _173_v30
		_173_v30 = Companion_D3_.Create_DC7_(_dafny.MultiSetFromSeq(_172_v29))
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_145_v6), 0))
		_ = _index14
		var _rhs20 *C1 = _171_v28
		_ = _rhs20
		var _rhs21 _dafny.Int = (func() _dafny.Int {
			if (_142_v3).Cmp(_142_v3) >= 0 {
				return ((_141_v2).Union(_141_v2)).Cardinality()
			}
			return _dafny.IntOfInt64(-794)
		})()
		_ = _rhs21
		var _rhs22 D3 = _173_v30
		_ = _rhs22
		var _rhs23 bool = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_165_v22).Cardinality()), _142_v3)).Cmp(_142_v3) >= 0
		_ = _rhs23
		var _lhs17 _dafny.Array = _145_v6
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_145_v6), 0))
		_ = _lhs18
		_171_v28 = _rhs20
		(_lhs17).ArraySet1(_rhs21, (_lhs18).Int())
		_173_v30 = _rhs22
		_149_v8 = _rhs23
		var _174_v31 _dafny.Array
		_ = _174_v31
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_2
		var _nw20 _dafny.Array
		_ = _nw20
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw20 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Set = func(_175_i2 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(true)
			}
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
		_174_v31 = _nw20
		var _176_v32 _dafny.Array
		_ = _176_v32
		var _nwElement0_2 _dafny.Array = _174_v31
		_ = _nwElement0_2
		var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
		_ = _nw21
		(_nw21).ArraySet1(_nwElement0_2, 0)
		(_nw21).ArraySet1(_174_v31, 1)
		(_nw21).ArraySet1(_174_v31, 2)
		(_nw21).ArraySet1(_174_v31, 3)
		_176_v32 = _nw21
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_176_v32), 0))
		_ = _index15
		(_176_v32).ArraySet1(_174_v31, (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_176_v32), 0))
		_ = _index16
		(_176_v32).ArraySet1(_174_v31, (_index16).Int())
		var _177_v33 _dafny.Sequence
		_ = _177_v33
		_177_v33 = _dafny.SeqOf(_165_v22, Companion_Default___.Fm5(_142_v3, _142_v3, _142_v3, _147_globalState))
		var _178_v34 *C2
		_ = _178_v34
		var _nw22 *C2 = New_C2_()
		_ = _nw22
		_nw22.Ctor__(_177_v33, _dafny.UnicodeSeqOfUtf8Bytes("m"), _149_v8)
		_178_v34 = _nw22
		var _179_v35 _dafny.Set
		_ = _179_v35
		_179_v35 = _dafny.SetOf(_178_v34)
		_179_v35 = _dafny.SetOf(_178_v34, _178_v34, _178_v34, _178_v34, _178_v34)
		var _180_v36 D0
		_ = _180_v36
		_180_v36 = Companion_D0_.Create_DC1_((_171_v28).F25())
		var _181_v37 *C1
		_ = _181_v37
		var _nw23 *C1 = New_C1_()
		_ = _nw23
		_nw23.Ctor__(_171_v28.F24, (_180_v36).Dtor_cf1())
		_181_v37 = _nw23
	}
	_149_v8 = _149_v8
	var _182_v38 _dafny.MultiSet
	_ = _182_v38
	_182_v38 = _dafny.MultiSetOf(false)
	var _183_v39 _dafny.MultiSet
	_ = _183_v39
	_183_v39 = _dafny.MultiSetOf(_182_v38, _182_v38)
	var _184_v40 D0
	_ = _184_v40
	_184_v40 = Companion_D0_.Create_DC0_(_144_v5)
	var _185_v41 D0
	_ = _185_v41
	_185_v41 = Companion_D0_.Create_DC2_(_184_v40)
	var _source3 D0 = (func() D0 {
		if Companion_Default___.Fm12(_142_v3, _182_v38, _149_v8, _183_v39, _147_globalState) {
			return _185_v41
		}
		return _185_v41
	})()
	_ = _source3
	if _source3.Is_DC1() {
		var _186___mcc_h0 bool = _source3.Get_().(D0_DC1).Cf1
		_ = _186___mcc_h0
		var _187_cf1 bool = _186___mcc_h0
		_ = _187_cf1
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_145_v6), 0))
		_ = _index17
		(_145_v6).ArraySet1(_dafny.IntOfInt64(203), (_index17).Int())
		var _188_v42 _dafny.Sequence
		_ = _188_v42
		_188_v42 = _dafny.SeqOf(_142_v3, _dafny.IntOfInt64(555), _142_v3, _142_v3)
		var _189_v43 _dafny.Sequence
		_ = _189_v43
		_189_v43 = _dafny.SeqOf(_149_v8, _149_v8)
		var _190_v44 _dafny.Array
		_ = _190_v44
		var _nwElement0_3 _dafny.MultiSet = _182_v38
		_ = _nwElement0_3
		var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(19))
		_ = _nw24
		(_nw24).ArraySet1(_nwElement0_3, 0)
		(_nw24).ArraySet1(_182_v38, 1)
		(_nw24).ArraySet1(_182_v38, 2)
		(_nw24).ArraySet1(_182_v38, 3)
		(_nw24).ArraySet1(_182_v38, 4)
		(_nw24).ArraySet1(_182_v38, 5)
		(_nw24).ArraySet1(_182_v38, 6)
		(_nw24).ArraySet1(_dafny.MultiSetOf(_187_cf1, _187_cf1, _187_cf1), 7)
		(_nw24).ArraySet1(_182_v38, 8)
		(_nw24).ArraySet1(Companion_Default___.Fm17(_dafny.IntOfUint32((_188_v42).Cardinality()), _149_v8, _147_globalState), 9)
		(_nw24).ArraySet1(_182_v38, 10)
		(_nw24).ArraySet1(_182_v38, 11)
		(_nw24).ArraySet1(_182_v38, 12)
		(_nw24).ArraySet1(_182_v38, 13)
		(_nw24).ArraySet1(_dafny.MultiSetFromSeq(_189_v43), 14)
		(_nw24).ArraySet1(_182_v38, 15)
		(_nw24).ArraySet1(_182_v38, 16)
		(_nw24).ArraySet1(_182_v38, 17)
		(_nw24).ArraySet1(Companion_Default___.Fm17(_142_v3, true, _147_globalState), 18)
		_190_v44 = _nw24
		var _191_v45 _dafny.Sequence
		_ = _191_v45
		_191_v45 = _dafny.SeqOf(_190_v44)
		var _192_v46 *C1
		_ = _192_v46
		var _nw25 *C1 = New_C1_()
		_ = _nw25
		_nw25.Ctor__((_191_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.IntOfUint32((_191_v45).Cardinality()))).Uint32()).(_dafny.Array), _149_v8)
		_192_v46 = _nw25
		var _193_v47 _dafny.Sequence
		_ = _193_v47
		_193_v47 = _dafny.SeqOf(_192_v46, _192_v46, _192_v46)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_145_v6), 0))
		_ = _index18
		var _rhs24 _dafny.Int = (_142_v3).Plus(_dafny.IntOfUint32((_193_v47).Cardinality()))
		_ = _rhs24
		var _rhs25 _dafny.Int = (_142_v3).Minus(_dafny.IntOfInt64(-786))
		_ = _rhs25
		var _lhs19 _dafny.Array = _145_v6
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_145_v6), 0))
		_ = _lhs20
		var _lhs21 *GlobalState = _147_globalState
		_ = _lhs21
		(_lhs19).ArraySet1(_rhs24, (_lhs20).Int())
		_lhs21.F20 = _rhs25
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_145_v6), 0))
		_ = _index19
		(_145_v6).ArraySet1(_dafny.IntOfInt64(723), (_index19).Int())
		var _194_v48 _dafny.Array
		_ = _194_v48
		var _nw26 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
		_ = _nw26
		_194_v48 = _nw26
		var _195_v49 T0
		_ = _195_v49
		var _nw27 *C3 = New_C3_()
		_ = _nw27
		_nw27.Ctor__((_192_v46).F25(), _149_v8, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("r"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_142_v3), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()))).Uint32(), _143_v4), false)
		_195_v49 = _nw27
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_194_v48), 0))
		_ = _index20
		(_194_v48).ArraySet1(_195_v49, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_194_v48), 0))
		_ = _index21
		(_194_v48).ArraySet1(_195_v49, (_index21).Int())
		var _196_v50 _dafny.Array
		_ = _196_v50
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw28
		_196_v50 = _nw28
		_196_v50 = _196_v50
	} else if _source3.Is_DC0() {
		var _197___mcc_h1 _dafny.Array = _source3.Get_().(D0_DC0).Cf0
		_ = _197___mcc_h1
		var _198_cf0 _dafny.Array = _197___mcc_h1
		_ = _198_cf0
		var _199_v51 _dafny.Sequence
		_ = _199_v51
		_199_v51 = _dafny.SeqOf(_142_v3)
		var _200_v52 _dafny.Sequence
		_ = _200_v52
		_200_v52 = _dafny.SeqOf(_199_v51)
		var _201_v53 T2
		_ = _201_v53
		var _nw29 *C2 = New_C2_()
		_ = _nw29
		_nw29.Ctor__(_200_v52, _140_v1, (func() bool {
			if _149_v8 {
				return _149_v8
			}
			return _149_v8
		})())
		_201_v53 = _nw29
		_201_v53 = _201_v53
		var _202_v54 _dafny.Map
		_ = _202_v54
		_202_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v1, _142_v3)
		var _203_v56 *C0
		_ = _203_v56
		var _nw30 *C0 = New_C0_()
		_ = _nw30
		_nw30.Ctor__(func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_202_v54).Keys().Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _204_v55 _dafny.Sequence
				_204_v55 = interface{}(_compr_9).(_dafny.Sequence)
				if (_202_v54).Contains(_204_v55) {
					_coll9.Add(_204_v55)
				}
			}
			return _coll9.ToSet()
		}())
		_203_v56 = _nw30
		(_147_globalState).F8 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_205_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_206_i3 _dafny.Int) _dafny.CodePoint {
				return _205_v4
			}
		})(_143_v4)))
		var _207_v57 _dafny.Sequence
		_ = _207_v57
		_207_v57 = _dafny.SeqOf(!(_201_v53.F22()), true, _201_v53.F22())
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_207_v57, _207_v57) {
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_145_v6), 0))
			_ = _index22
			(_145_v6).ArraySet1(_dafny.IntOfInt64(842), (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_144_v5), 0))
			_ = _index23
			(_144_v5).ArraySet1(!(_149_v8), (_index23).Int())
			var _208_v58 _dafny.Sequence
			_ = _208_v58
			_208_v58 = _dafny.SeqOf(_201_v53)
			var _209_v59 _dafny.Set
			_ = _209_v59
			_209_v59 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('y')))
			var _210_v60 _dafny.Map
			_ = _210_v60
			_210_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_209_v59, _201_v53)
			var _211_v61 *C3
			_ = _211_v61
			var _nw31 *C3 = New_C3_()
			_ = _nw31
			_nw31.Ctor__(_201_v53.F22(), _149_v8, _201_v53.F21(), false)
			_211_v61 = _nw31
			var _212_v62 _dafny.MultiSet
			_ = _212_v62
			_212_v62 = _dafny.MultiSetOf(_211_v61)
			var _213_v63 *C3
			_ = _213_v63
			_213_v63 = _211_v61
			var _214_v64 _dafny.Array
			_ = _214_v64
			var _nwElement0_4 T2 = _201_v53
			_ = _nwElement0_4
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_4, 0)
			(_nw32).ArraySet1(_201_v53, 1)
			(_nw32).ArraySet1(_201_v53, 2)
			(_nw32).ArraySet1(_201_v53, 3)
			(_nw32).ArraySet1(_201_v53, 4)
			(_nw32).ArraySet1((_208_v58).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_142_v3), _dafny.IntOfUint32((_208_v58).Cardinality()))).Uint32()).(T2), 5)
			(_nw32).ArraySet1(_201_v53, 6)
			(_nw32).ArraySet1(_201_v53, 7)
			(_nw32).ArraySet1(_201_v53, 8)
			(_nw32).ArraySet1(_201_v53, 9)
			(_nw32).ArraySet1((func() T2 {
				if (_210_v60).Contains(_209_v59) {
					return (_210_v60).Get(_209_v59).(T2)
				}
				return _201_v53
			})(), 10)
			(_nw32).ArraySet1(_201_v53, 11)
			(_nw32).ArraySet1(_201_v53, 12)
			(_nw32).ArraySet1((func() T2 {
				if _201_v53.F22() {
					return _201_v53
				}
				return _201_v53
			})(), 13)
			(_nw32).ArraySet1(_201_v53, 14)
			(_nw32).ArraySet1(_201_v53, 15)
			(_nw32).ArraySet1(_201_v53, 16)
			(_nw32).ArraySet1(_201_v53, 17)
			(_nw32).ArraySet1((_208_v58).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_212_v62).Contains((_213_v63)) {
					return (_212_v62).Multiplicity((_213_v63))
				}
				return _142_v3
			})(), _dafny.IntOfUint32((_208_v58).Cardinality()))).Uint32()).(T2), 18)
			(_nw32).ArraySet1((_208_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_201_v53.F21()).Cardinality()), _dafny.IntOfUint32((_208_v58).Cardinality()))).Uint32()).(T2), 19)
			(_nw32).ArraySet1(_201_v53, 20)
			(_nw32).ArraySet1(_201_v53, 21)
			(_nw32).ArraySet1(_201_v53, 22)
			(_nw32).ArraySet1(_201_v53, 23)
			(_nw32).ArraySet1(_201_v53, 24)
			(_nw32).ArraySet1(_201_v53, 25)
			(_nw32).ArraySet1(_201_v53, 26)
			(_nw32).ArraySet1(_201_v53, 27)
			_214_v64 = _nw32
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_214_v64), 0))
			_ = _index24
			(_214_v64).ArraySet1(_201_v53, (_index24).Int())
			var _215_v65 _dafny.Sequence
			_ = _215_v65
			_215_v65 = _dafny.SeqOf(_201_v53.F21(), _140_v1)
			var _216_v66 _dafny.Map
			_ = _216_v66
			_216_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v61.F28, _142_v3)
			var _217_v67 _dafny.Map
			_ = _217_v67
			_217_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v66, _142_v3)
			var _218_v68 _dafny.Array
			_ = _218_v68
			var _nwElement0_5 _dafny.Sequence = (_215_v65).Select((Companion_Default___.SafeIndex((_217_v67).Cardinality(), _dafny.IntOfUint32((_215_v65).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _nwElement0_5
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(14))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_5, 0)
			(_nw33).ArraySet1(_140_v1, 1)
			(_nw33).ArraySet1(_201_v53.F21(), 2)
			(_nw33).ArraySet1(_140_v1, 3)
			(_nw33).ArraySet1(_201_v53.F21(), 4)
			(_nw33).ArraySet1(Companion_Default___.Fm11((_211_v61).F27(), _147_globalState), 5)
			(_nw33).ArraySet1(_201_v53.F21(), 6)
			(_nw33).ArraySet1(_201_v53.F21(), 7)
			(_nw33).ArraySet1(_201_v53.F21(), 8)
			(_nw33).ArraySet1(_140_v1, 9)
			(_nw33).ArraySet1(Companion_Default___.Fm11(_211_v61.F28, _147_globalState), 10)
			(_nw33).ArraySet1(_201_v53.F21(), 11)
			(_nw33).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(434))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_219_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_220_i4 _dafny.Int) _dafny.CodePoint {
					return _219_v4
				}
			})(_143_v4))), 12)
			(_nw33).ArraySet1(_201_v53.F21(), 13)
			_218_v68 = _nw33
			var _221_v69 _dafny.MultiSet
			_ = _221_v69
			_221_v69 = _dafny.MultiSetOf(_218_v68)
			var _222_v70 _dafny.Set
			_ = _222_v70
			_222_v70 = _dafny.SetOf(_145_v6, _145_v6, _145_v6, _145_v6, _145_v6)
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_145_v6), 0))
			_ = _index25
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_144_v5), 0))
			_ = _index26
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_214_v64), 0))
			_ = _index27
			var _rhs26 _dafny.Int = (func() _dafny.Int {
				if (_221_v69).Contains(_218_v68) {
					return (_221_v69).Multiplicity(_218_v68)
				}
				return _142_v3
			})()
			_ = _rhs26
			var _rhs27 bool = (_222_v70).IsDisjointFrom((_222_v70).Intersection(_dafny.SetOf(_145_v6)))
			_ = _rhs27
			var _rhs28 T2 = _201_v53
			_ = _rhs28
			var _lhs22 _dafny.Array = _145_v6
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_145_v6), 0))
			_ = _lhs23
			var _lhs24 _dafny.Array = _144_v5
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_144_v5), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _214_v64
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_214_v64), 0))
			_ = _lhs27
			(_lhs22).ArraySet1(_rhs26, (_lhs23).Int())
			(_lhs24).ArraySet1(_rhs27, (_lhs25).Int())
			(_lhs26).ArraySet1(_rhs28, (_lhs27).Int())
			var _223_v71 _dafny.Array
			_ = _223_v71
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw34
			_223_v71 = _nw34
			_223_v71 = _223_v71
			var _224_v72 D1
			_ = _224_v72
			_224_v72 = Companion_D1_.Create_DC5_(_149_v8, _216_v66, _142_v3, _146_v7)
			var _225_v73 T0
			_ = _225_v73
			var _nw35 *C3 = New_C3_()
			_ = _nw35
			_nw35.Ctor__((_142_v3).Cmp(_142_v3) == 0, _211_v61.F28, _201_v53.F21(), (_142_v3).Cmp((_224_v72).Dtor_cf9()) <= 0)
			_225_v73 = _nw35
			var _226_v74 _dafny.Set
			_ = _226_v74
			_226_v74 = _dafny.SetOf(_142_v3)
			var _227_v75 _dafny.Set
			_ = _227_v75
			_227_v75 = _226_v74
			var _228_v76 T1
			_ = _228_v76
			var _nw36 *C4 = New_C4_()
			_ = _nw36
			_nw36.Ctor__(_227_v75, (_211_v61).F27(), _200_v52, _140_v1, (_142_v3).Cmp(_dafny.IntOfUint32((_207_v57).Cardinality())) < 0)
			_228_v76 = _nw36
			var _rhs29 _dafny.Sequence = Companion_Default___.Fm11((_225_v73).Fm1(_140_v1, (_144_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_144_v5), 0))).Int()).(bool), (_211_v61).F27(), _147_globalState), _147_globalState)
			_ = _rhs29
			var _rhs30 T0 = _225_v73
			_ = _rhs30
			var _rhs31 T1 = _228_v76
			_ = _rhs31
			var _lhs28 T2 = _201_v53
			_ = _lhs28
			_lhs28.F21_set_(_rhs29)
			_225_v73 = _rhs30
			_228_v76 = _rhs31
			var _229_v77 _dafny.Map
			_ = _229_v77
			_229_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_145_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_145_v6), 0))).Int()).(_dafny.Int), _225_v73.F22())
			var _230_v78 _dafny.Map
			_ = _230_v78
			_230_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v77, _198_cf0)
			var _231_v79 _dafny.Int
			_ = _231_v79
			var _out2 _dafny.Int
			_ = _out2
			_out2 = (_201_v53).M1(_225_v73.F22(), _230_v78, _147_globalState)
			_231_v79 = _out2
			var _232_v80 _dafny.Array
			_ = _232_v80
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
			_ = _nw37
			_232_v80 = _nw37
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_232_v80), 0))
			_ = _index28
			(_232_v80).ArraySet1CodePoint(_143_v4, (_index28).Int())
			var _233_v81 _dafny.Set
			_ = _233_v81
			_233_v81 = _dafny.SetOf(!((_211_v61).F27()), _149_v8)
			var _234_v82 *C4
			_ = _234_v82
			var _nw38 *C4 = New_C4_()
			_ = _nw38
			_nw38.Ctor__(_227_v75, !(true) || (_149_v8), Companion_Default___.Fm16(_233_v81, _211_v61.F28, _149_v8, _147_globalState), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("lpt"), (Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lpt")).Cardinality()))).Uint32(), _143_v4), _201_v53.F22())
			_234_v82 = _nw38
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_144_v5), 0))
			_ = _index29
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_232_v80), 0))
			_ = _index30
			var _rhs32 bool = (_234_v82).Fm1(_140_v1, !((_234_v82).F30()), _211_v61.F28, _147_globalState)
			_ = _rhs32
			var _rhs33 _dafny.CodePoint = _143_v4
			_ = _rhs33
			var _rhs34 *C4 = _234_v82
			_ = _rhs34
			var _lhs29 _dafny.Array = _144_v5
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_144_v5), 0))
			_ = _lhs30
			var _lhs31 _dafny.Array = _232_v80
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_232_v80), 0))
			_ = _lhs32
			(_lhs29).ArraySet1(_rhs32, (_lhs30).Int())
			(_lhs31).ArraySet1CodePoint(_rhs33, (_lhs32).Int())
			_234_v82 = _rhs34
		} else {
			(_201_v53).F22_set_((_207_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if _201_v53.F22() {
					return _199_v51
				}
				return Companion_Default___.Fm5((_dafny.Zero).Minus(Companion_Default___.Fm6(_147_globalState)), _142_v3, _dafny.IntOfInt64(109), _147_globalState)
			})()).Cardinality()), _dafny.IntOfUint32((_207_v57).Cardinality()))).Uint32()).(bool))
			(_147_globalState).F20 = _142_v3
			_149_v8 = _149_v8
			var _235_v83 _dafny.Array
			_ = _235_v83
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
			_ = _nw39
			_235_v83 = _nw39
			var _236_v84 *C1
			_ = _236_v84
			var _nw40 *C1 = New_C1_()
			_ = _nw40
			_nw40.Ctor__(_235_v83, false)
			_236_v84 = _nw40
			var _237_v85 _dafny.Map
			_ = _237_v85
			_237_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v53.F22(), _dafny.MultiSetFromSeq(_dafny.SeqOf(false)))
			var _238_v86 _dafny.Map
			_ = _238_v86
			_238_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v84, Companion_Default___.Fm12(_142_v3, (func() _dafny.MultiSet {
				if (_237_v85).Contains(false) {
					return (_237_v85).Get(false).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(_201_v53.F22())
			})(), _201_v53.F22(), _183_v39, _147_globalState))
			_238_v86 = (_238_v86).Update(_236_v84, (_207_v57).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm6(_147_globalState), _dafny.IntOfUint32((_207_v57).Cardinality()))).Uint32()).(bool))
			_182_v38 = _dafny.MultiSetOf(!(_201_v53.F22()))
		}
	} else {
		var _239___mcc_h2 D0 = _source3.Get_().(D0_DC2).Cf2
		_ = _239___mcc_h2
		var _240_cf2 D0 = _239___mcc_h2
		_ = _240_cf2
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_145_v6), 0))
		_ = _index31
		(_145_v6).ArraySet1(_142_v3, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_145_v6), 0))
		_ = _index32
		(_145_v6).ArraySet1(_142_v3, (_index32).Int())
		(_147_globalState).F20 = Companion_Default___.Fm6(_147_globalState)
		var _241_v87 _dafny.Sequence
		_ = _241_v87
		_241_v87 = _dafny.SeqOf((_145_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_145_v6), 0))).Int()).(_dafny.Int), _142_v3)
		var _rhs35 _dafny.Int = (_142_v3).Plus(_142_v3)
		_ = _rhs35
		var _rhs36 _dafny.Int = (_241_v87).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_241_v87).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs36
		var _lhs33 *GlobalState = _147_globalState
		_ = _lhs33
		var _lhs34 *GlobalState = _147_globalState
		_ = _lhs34
		_lhs33.F20 = _rhs35
		_lhs34.F3 = _rhs36
		var _242_v88 _dafny.Set
		_ = _242_v88
		_242_v88 = _dafny.SetOf((_145_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_145_v6), 0))).Int()).(_dafny.Int))
		var _243_v89 _dafny.Sequence
		_ = _243_v89
		_243_v89 = _dafny.SeqOf(_242_v88)
		var _244_v90 _dafny.Set
		_ = _244_v90
		_244_v90 = (_243_v89).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_243_v89).Cardinality()))).Uint32()).(_dafny.Set)
		var _245_v91 _dafny.Sequence
		_ = _245_v91
		_245_v91 = _dafny.SeqOf(_241_v87, _241_v87)
		var _246_v92 T1
		_ = _246_v92
		var _nw41 *C4 = New_C4_()
		_ = _nw41
		_nw41.Ctor__(_dafny.SetOf((_145_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_145_v6), 0))).Int()).(_dafny.Int)), _149_v8, _245_v91, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-37))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}((func(_247_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_248_i5 _dafny.Int) _dafny.CodePoint {
				return _247_v4
			}
		})(_143_v4))), _dafny.Companion_Sequence_.Contains(_241_v87, Companion_Default___.Fm6(_147_globalState)))
		_246_v92 = _nw41
		var _rhs37 bool = _246_v92.F22()
		_ = _rhs37
		var _rhs38 T1 = (func() T1 {
			if !(_149_v8) || (_149_v8) {
				return _246_v92
			}
			return _246_v92
		})()
		_ = _rhs38
		var _rhs39 _dafny.Int = _142_v3
		_ = _rhs39
		var _lhs35 *GlobalState = _147_globalState
		_ = _lhs35
		_149_v8 = _rhs37
		_246_v92 = _rhs38
		_lhs35.F3 = _rhs39
	}
	var _249_v93 _dafny.MultiSet
	_ = _249_v93
	_249_v93 = _dafny.MultiSetOf(_144_v5, _144_v5)
	var _250_v94 _dafny.Map
	_ = _250_v94
	_250_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _149_v8)
	var _251_v95 _dafny.Sequence
	_ = _251_v95
	_251_v95 = _dafny.SeqOf(_149_v8, _149_v8)
	var _252_v96 _dafny.Array
	_ = _252_v96
	var _nwElement0_6 _dafny.MultiSet = _249_v93
	_ = _nwElement0_6
	var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(25))
	_ = _nw42
	(_nw42).ArraySet1(_nwElement0_6, 0)
	(_nw42).ArraySet1(_249_v93, 1)
	(_nw42).ArraySet1(_249_v93, 2)
	(_nw42).ArraySet1((_249_v93).Intersection(((Companion_D7_.Create_DC15_(_249_v93)).Dtor_cf25()).Update(_144_v5, Companion_Default___.Abs(_142_v3))), 3)
	(_nw42).ArraySet1(_249_v93, 4)
	(_nw42).ArraySet1(_dafny.MultiSetOf(_144_v5, _144_v5), 5)
	(_nw42).ArraySet1((_249_v93).Difference(_249_v93), 6)
	(_nw42).ArraySet1(_249_v93, 7)
	(_nw42).ArraySet1(_249_v93, 8)
	(_nw42).ArraySet1(_249_v93, 9)
	(_nw42).ArraySet1(_249_v93, 10)
	(_nw42).ArraySet1((_249_v93).Union(_249_v93), 11)
	(_nw42).ArraySet1(_249_v93, 12)
	(_nw42).ArraySet1((func() _dafny.MultiSet {
		if (func() bool {
			if (_250_v94).Contains(_142_v3) {
				return (_250_v94).Get(_142_v3).(bool)
			}
			return !(_149_v8)
		})() {
			return _249_v93
		}
		return _249_v93
	})(), 13)
	(_nw42).ArraySet1(_249_v93, 14)
	(_nw42).ArraySet1(_dafny.MultiSetOf(_144_v5), 15)
	(_nw42).ArraySet1(_249_v93, 16)
	(_nw42).ArraySet1((func() _dafny.MultiSet {
		if Companion_Default___.Fm12(_dafny.IntOfInt64(572), _dafny.MultiSetOf(_149_v8, _149_v8, _149_v8, _149_v8, _149_v8), (_251_v95).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_140_v1).Cardinality()), _dafny.IntOfUint32((_251_v95).Cardinality()))).Uint32()).(bool), _dafny.MultiSetOf(_182_v38), _147_globalState) {
			return _249_v93
		}
		return _249_v93
	})(), 17)
	(_nw42).ArraySet1((Companion_D7_.Create_DC15_(_249_v93)).Dtor_cf25(), 18)
	(_nw42).ArraySet1(_249_v93, 19)
	(_nw42).ArraySet1(_249_v93, 20)
	(_nw42).ArraySet1((_249_v93).Update(_144_v5, Companion_Default___.Abs(_142_v3)), 21)
	(_nw42).ArraySet1(_249_v93, 22)
	(_nw42).ArraySet1((_249_v93).Union(_249_v93), 23)
	(_nw42).ArraySet1((_dafny.MultiSetOf(_144_v5)).Union(_249_v93), 24)
	_252_v96 = _nw42
	var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_252_v96), 0))
	_ = _index33
	(_252_v96).ArraySet1((_dafny.MultiSetOf(_144_v5, _144_v5)).Union(_249_v93), (_index33).Int())
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_252_v96), 0))
	_ = _index34
	(_252_v96).ArraySet1((((_249_v93).Update(_144_v5, Companion_Default___.Abs((_250_v94).Cardinality()))).Union(_249_v93)).Difference(_249_v93), (_index34).Int())
	var _253_v97 _dafny.Set
	_ = _253_v97
	_253_v97 = _dafny.SetOf(_142_v3, _142_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ly")).Cardinality()))
	var _254_v98 _dafny.Set
	_ = _254_v98
	_254_v98 = _253_v97
	var _255_v99 _dafny.Sequence
	_ = _255_v99
	_255_v99 = _dafny.SeqOf(_142_v3)
	var _256_v100 *C4
	_ = _256_v100
	var _nw43 *C4 = New_C4_()
	_ = _nw43
	_nw43.Ctor__(_254_v98, _149_v8, _dafny.SeqOf(_dafny.SeqOf(_142_v3), _255_v99, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_142_v3), (Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_dafny.SeqOf(_142_v3)).Cardinality()))).Uint32(), _142_v3)), _140_v1, true)
	_256_v100 = _nw43
	var _257_v101 _dafny.Set
	_ = _257_v101
	_257_v101 = _dafny.SetOf(_256_v100, _256_v100, _256_v100)
	var _258_v102 _dafny.Map
	_ = _258_v102
	_258_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12(_dafny.IntOfUint32((_140_v1).Cardinality()), _dafny.MultiSetOf(_149_v8), _149_v8, _183_v39, _147_globalState), _257_v101)
	var _hi1 _dafny.Int = (_dafny.Zero).Minus((_258_v102).Cardinality())
	_ = _hi1
	for _259_i6 := _142_v3; _259_i6.Cmp(_hi1) < 0; _259_i6 = _259_i6.Plus(_dafny.One) {
		var _260_v103 _dafny.Sequence
		_ = _260_v103
		_260_v103 = _dafny.SeqOf(_255_v99)
		var _261_v104 _dafny.Sequence
		_ = _261_v104
		_261_v104 = _dafny.SeqOf((_260_v103).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_260_v103).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _262_v105 *C2
		_ = _262_v105
		var _nw44 *C2 = New_C2_()
		_ = _nw44
		_nw44.Ctor__(_260_v103, _dafny.Companion_Sequence_.Concatenate(_140_v1, _dafny.Companion_Sequence_.Update(_140_v1, (Companion_Default___.SafeIndex(_259_i6, _dafny.IntOfUint32((_140_v1).Cardinality()))).Uint32(), _dafny.CodePoint('u'))), (_142_v3).Cmp(_142_v3) >= 0)
		_262_v105 = _nw44
		var _263_v106 _dafny.Sequence
		_ = _263_v106
		_263_v106 = _dafny.SeqOf(_250_v94)
		var _264_v107 _dafny.Int
		_ = _264_v107
		var _out3 _dafny.Int
		_ = _out3
		_out3 = (_262_v105).M1((_256_v100).F30(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v106).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.IntOfUint32((_263_v106).Cardinality()))).Uint32()).(_dafny.Map), _144_v5), _147_globalState)
		_264_v107 = _out3
		_149_v8 = (_256_v100).F30()
		_255_v99 = _255_v99
	}
	_149_v8 = (_256_v100).F30()
	var _265_v108 _dafny.Array
	_ = _265_v108
	var _nwElement0_7 _dafny.CodePoint = _143_v4
	_ = _nwElement0_7
	var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(8))
	_ = _nw45
	(_nw45).ArraySet1CodePoint(_nwElement0_7, 0)
	(_nw45).ArraySet1CodePoint((func() _dafny.CodePoint {
		if _149_v8 {
			return _143_v4
		}
		return _143_v4
	})(), 1)
	(_nw45).ArraySet1CodePoint(_143_v4, 2)
	(_nw45).ArraySet1CodePoint(_143_v4, 3)
	(_nw45).ArraySet1CodePoint(_143_v4, 4)
	(_nw45).ArraySet1CodePoint(_143_v4, 5)
	(_nw45).ArraySet1CodePoint(_143_v4, 6)
	(_nw45).ArraySet1CodePoint(_143_v4, 7)
	_265_v108 = _nw45
	_265_v108 = _265_v108
	_255_v99 = _255_v99
	var _266_v109 D3
	_ = _266_v109
	_266_v109 = Companion_D3_.Create_DC7_(_182_v38)
	var _pat_let_tv14 = _182_v38
	_ = _pat_let_tv14
	var _267_v110 _dafny.Array
	_ = _267_v110
	var _nwElement0_8 _dafny.MultiSet = _dafny.MultiSetOf(false)
	_ = _nwElement0_8
	var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
	_ = _nw46
	(_nw46).ArraySet1(_nwElement0_8, 0)
	(_nw46).ArraySet1(Companion_Default___.Fm17(_142_v3, (_256_v100).F30(), _147_globalState), 1)
	(_nw46).ArraySet1(_182_v38, 2)
	(_nw46).ArraySet1(_182_v38, 3)
	(_nw46).ArraySet1(_182_v38, 4)
	(_nw46).ArraySet1((func(_pat_let0_0 D3) D3 {
		return func(_268_dt__update__tmp_h1 D3) D3 {
			return func(_pat_let1_0 _dafny.MultiSet) D3 {
				return func(_269_dt__update_hcf12_h0 _dafny.MultiSet) D3 {
					return Companion_D3_.Create_DC7_(_269_dt__update_hcf12_h0)
				}(_pat_let1_0)
			}(_pat_let_tv14)
		}(_pat_let0_0)
	}(_266_v109)).Dtor_cf12(), 5)
	(_nw46).ArraySet1(_182_v38, 6)
	(_nw46).ArraySet1(_182_v38, 7)
	(_nw46).ArraySet1(_182_v38, 8)
	(_nw46).ArraySet1(_dafny.MultiSetOf(_149_v8, (_256_v100).F30()), 9)
	(_nw46).ArraySet1(_182_v38, 10)
	(_nw46).ArraySet1(_182_v38, 11)
	(_nw46).ArraySet1(_182_v38, 12)
	(_nw46).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_251_v95, (Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_251_v95).Cardinality()))).Uint32(), (_256_v100).Fm1(_140_v1, _149_v8, (_256_v100).F30(), _147_globalState))), 13)
	(_nw46).ArraySet1(_dafny.MultiSetOf(!((_256_v100).F30())), 14)
	(_nw46).ArraySet1(_182_v38, 15)
	(_nw46).ArraySet1(_182_v38, 16)
	(_nw46).ArraySet1(_182_v38, 17)
	(_nw46).ArraySet1(_dafny.MultiSetOf((_256_v100).F30()), 18)
	(_nw46).ArraySet1(_182_v38, 19)
	(_nw46).ArraySet1((_182_v38).Update(_149_v8, Companion_Default___.Abs((_dafny.Zero).Minus(_142_v3))), 20)
	(_nw46).ArraySet1(_182_v38, 21)
	(_nw46).ArraySet1(_182_v38, 22)
	(_nw46).ArraySet1(Companion_Default___.Fm17(_dafny.IntOfInt64(889), _149_v8, _147_globalState), 23)
	(_nw46).ArraySet1(_dafny.MultiSetFromSeq(_251_v95), 24)
	_267_v110 = _nw46
	var _270_v111 *C1
	_ = _270_v111
	var _nw47 *C1 = New_C1_()
	_ = _nw47
	_nw47.Ctor__(_267_v110, (_256_v100).F30())
	_270_v111 = _nw47
	var _271_v112 _dafny.Map
	_ = _271_v112
	_271_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _142_v3)
	var _272_v113 _dafny.Map
	_ = _272_v113
	_272_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v38, _142_v3)
	var _273_v114 D1
	_ = _273_v114
	_273_v114 = Companion_D1_.Create_DC5_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_270_v111).F25(), _142_v3), _142_v3, _146_v7)
	_142_v3 = Companion_Default___.SafeModuloInt((_271_v112).Cardinality(), (Companion_D3_.Create_DC8_(_142_v3, _272_v113, _273_v114, _149_v8)).Dtor_cf13())
	var _hi2 _dafny.Int = _142_v3
	_ = _hi2
	for _274_i7 := Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(86), _142_v3); _274_i7.Cmp(_hi2) < 0; _274_i7 = _274_i7.Plus(_dafny.One) {
		var _275_v115 _dafny.Sequence
		_ = _275_v115
		_275_v115 = _dafny.SeqOf(Companion_Default___.Fm11((_270_v111).F25(), _147_globalState), _140_v1, Companion_Default___.Fm11(false, _147_globalState))
		_140_v1 = _dafny.Companion_Sequence_.Concatenate((_275_v115).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_275_v115).Cardinality()))).Uint32()).(_dafny.Sequence), _140_v1)
		(_147_globalState).F20 = _dafny.IntOfInt64(761)
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_265_v108), 0))
		_ = _index35
		(_265_v108).ArraySet1CodePoint(_143_v4, (_index35).Int())
		var _276_v116 _dafny.MultiSet
		_ = _276_v116
		_276_v116 = _dafny.MultiSetOf(_140_v1)
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_265_v108), 0))
		_ = _index36
		(_265_v108).ArraySet1CodePoint((_140_v1).Select((Companion_Default___.SafeIndex(((_dafny.MultiSetOf(_140_v1)).Intersection(_276_v116)).Cardinality(), _dafny.IntOfUint32((_140_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index36).Int())
		_142_v3 = ((_182_v38).Update((_270_v111).F25(), Companion_Default___.Abs((_dafny.Zero).Minus((_142_v3).Plus(_dafny.IntOfInt64(984)))))).Cardinality()
	}
	var _277_v117 *C3
	_ = _277_v117
	var _nw48 *C3 = New_C3_()
	_ = _nw48
	_nw48.Ctor__((_270_v111).F25(), _149_v8, _140_v1, (_270_v111).F25())
	_277_v117 = _nw48
	_142_v3 = _142_v3
	(_147_globalState).F12 = (func() _dafny.CodePoint {
		if (_277_v117).F27() {
			return Companion_Default___.Fm4(_277_v117.F28, _182_v38, _147_globalState)
		}
		return _dafny.CodePoint('q')
	})()
	(_277_v117).F28 = (_270_v111).F25()
	_dafny.Print(_140_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v2).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("truxk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_142_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_143_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_v6).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_v7).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-610))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F1()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("truxk"), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_globalState.F6.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_globalState.F8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F9().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_globalState.F12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F17()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_147_globalState).F17()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_globalState.F19).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-610))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_globalState.F20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_149_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_150_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v38).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_183_v39).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_249_v93).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v94).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-610), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_251_v95, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v96).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.MultiSet)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v97).Equals(_dafny.SetOf(_dafny.IntOfInt64(-610), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v98).Equals(_dafny.SetOf(_dafny.IntOfInt64(-610), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_255_v99, _dafny.SeqOf(_dafny.IntOfInt64(-610))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v100.F29).Equals(_dafny.SetOf(_dafny.IntOfInt64(-610), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v100).F30())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v101).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v102).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v108).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_266_v109).Dtor_cf12()).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v110).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v111.F24).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v111).F25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_271_v112).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-610), _dafny.IntOfInt64(-610))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v113).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _dafny.IntOfInt64(-610))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v114).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_273_v114).Dtor_cf8()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-610))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v114).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_273_v114).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-610))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v117).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_277_v117.F28)
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Array
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Array) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf2 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 D0) D0 {
	return D0{D0_DC2{Cf2}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf0() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf2() D0 {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf1 == data2.Cf1
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2.Equals(data2.Cf2)
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
	Cf4 _dafny.CodePoint
	Cf5 _dafny.Map
	Cf6 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.CodePoint, Cf5 _dafny.Map, Cf6 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf4, Cf5, Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf7  bool
	Cf8  _dafny.Map
	Cf9  _dafny.Int
	Cf10 _dafny.MultiSet
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 bool, Cf8 _dafny.Map, Cf9 _dafny.Int, Cf10 _dafny.MultiSet) D1 {
	return D1{D1_DC5{Cf7, Cf8, Cf9, Cf10}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf3 _dafny.Set
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Set) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.CodePoint('D'), _dafny.EmptyMap, _dafny.EmptySeq)
}

func (_this D1) Dtor_cf4() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Map {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.MultiSet {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf3() _dafny.Set {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5.Equals(data2.Cf5) && data1.Cf6.Equals(data2.Cf6)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Equals(data2.Cf8) && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Equals(data2.Cf10)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf11 _dafny.Set
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Set) D2 {
	return D2{D2_DC6{Cf11}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D2) Dtor_cf11() _dafny.Set {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf13 _dafny.Int
	Cf14 _dafny.Map
	Cf15 D1
	Cf16 bool
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf13 _dafny.Int, Cf14 _dafny.Map, Cf15 D1, Cf16 bool) D3 {
	return D3{D3_DC8{Cf13, Cf14, Cf15, Cf16}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf17 _dafny.Int
	Cf18 _dafny.Int
	Cf19 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf17 _dafny.Int, Cf18 _dafny.Int, Cf19 _dafny.Int) D3 {
	return D3{D3_DC9{Cf17, Cf18, Cf19}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC7 struct {
	Cf12 _dafny.MultiSet
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf12 _dafny.MultiSet) D3 {
	return D3{D3_DC7{Cf12}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.Zero, _dafny.EmptyMap, Companion_D1_.Default(), false)
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Map {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf15() D1 {
	return _this.Get_().(D3_DC8).Cf15
}

func (_this D3) Dtor_cf16() bool {
	return _this.Get_().(D3_DC8).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf19
}

func (_this D3) Dtor_cf12() _dafny.MultiSet {
	return _this.Get_().(D3_DC7).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf12) + ")"
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
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14.Equals(data2.Cf14) && data1.Cf15.Equals(data2.Cf15) && data1.Cf16 == data2.Cf16
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Cmp(data2.Cf19) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf21 _dafny.Set
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf21 _dafny.Set) D4 {
	return D4{D4_DC11{Cf21}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf20 *C0
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf20 *C0) D4 {
	return D4{D4_DC10{Cf20}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC12 struct {
	Cf22 D4
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf22 D4) D4 {
	return D4{D4_DC12{Cf22}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.EmptySet)
}

func (_this D4) Dtor_cf21() _dafny.Set {
	return _this.Get_().(D4_DC11).Cf21
}

func (_this D4) Dtor_cf20() *C0 {
	return _this.Get_().(D4_DC10).Cf20
}

func (_this D4) Dtor_cf22() D4 {
	return _this.Get_().(D4_DC12).Cf22
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf20) + ")"
		}
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
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf20 == data2.Cf20
		}
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

type D5_DC13 struct {
	Cf23 *C1
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf23 *C1) D5 {
	return D5{D5_DC13{Cf23}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() *C1 {
	return (*C1)(nil)
}

func (_this D5) Dtor_cf23() *C1 {
	return _this.Get_().(D5_DC13).Cf23
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf23 == data2.Cf23
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
	Cf24 *C3
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf24 *C3) D6 {
	return D6{D6_DC14{Cf24}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() *C3 {
	return (*C3)(nil)
}

func (_this D6) Dtor_cf24() *C3 {
	return _this.Get_().(D6_DC14).Cf24
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf24 == data2.Cf24
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
	Cf26 bool
	Cf27 bool
	Cf28 *C3
	Cf29 bool
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf26 bool, Cf27 bool, Cf28 *C3, Cf29 bool) D7 {
	return D7{D7_DC16{Cf26, Cf27, Cf28, Cf29}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC15 struct {
	Cf25 _dafny.MultiSet
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf25 _dafny.MultiSet) D7 {
	return D7{D7_DC15{Cf25}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC16_(false, false, (*C3)(nil), false)
}

func (_this D7) Dtor_cf26() bool {
	return _this.Get_().(D7_DC16).Cf26
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC16).Cf27
}

func (_this D7) Dtor_cf28() *C3 {
	return _this.Get_().(D7_DC16).Cf28
}

func (_this D7) Dtor_cf29() bool {
	return _this.Get_().(D7_DC16).Cf29
}

func (_this D7) Dtor_cf25() _dafny.MultiSet {
	return _this.Get_().(D7_DC15).Cf25
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29
		}
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf30 _dafny.Map
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf30 _dafny.Map) D8 {
	return D8{D8_DC17{Cf30}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D8) Dtor_cf30() _dafny.Map {
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

// Definition of datatype D9
type D9 struct {
	Data_D9_
}

func (_this D9) Get_() Data_D9_ {
	return _this.Data_D9_
}

type Data_D9_ interface {
	isD9()
}

type CompanionStruct_D9_ struct {
}

var Companion_D9_ = CompanionStruct_D9_{}

type D9_DC19 struct {
	Cf32 _dafny.Int
	Cf33 _dafny.Int
	Cf34 bool
	Cf35 _dafny.Int
	Cf36 _dafny.Set
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf32 _dafny.Int, Cf33 _dafny.Int, Cf34 bool, Cf35 _dafny.Int, Cf36 _dafny.Set) D9 {
	return D9{D9_DC19{Cf32, Cf33, Cf34, Cf35, Cf36}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC20 struct {
	Cf37 _dafny.Sequence
	Cf38 bool
	Cf39 _dafny.Int
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf37 _dafny.Sequence, Cf38 bool, Cf39 _dafny.Int) D9 {
	return D9{D9_DC20{Cf37, Cf38, Cf39}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC21 struct {
	Cf40 _dafny.Array
	Cf41 bool
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf40 _dafny.Array, Cf41 bool) D9 {
	return D9{D9_DC21{Cf40, Cf41}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

type D9_DC18 struct {
	Cf31 _dafny.Set
}

func (D9_DC18) isD9() {}

func (CompanionStruct_D9_) Create_DC18_(Cf31 _dafny.Set) D9 {
	return D9{D9_DC18{Cf31}}
}

func (_this D9) Is_DC18() bool {
	_, ok := _this.Get_().(D9_DC18)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC19_(_dafny.Zero, _dafny.Zero, false, _dafny.Zero, _dafny.EmptySet)
}

func (_this D9) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D9_DC19).Cf32
}

func (_this D9) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D9_DC19).Cf33
}

func (_this D9) Dtor_cf34() bool {
	return _this.Get_().(D9_DC19).Cf34
}

func (_this D9) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D9_DC19).Cf35
}

func (_this D9) Dtor_cf36() _dafny.Set {
	return _this.Get_().(D9_DC19).Cf36
}

func (_this D9) Dtor_cf37() _dafny.Sequence {
	return _this.Get_().(D9_DC20).Cf37
}

func (_this D9) Dtor_cf38() bool {
	return _this.Get_().(D9_DC20).Cf38
}

func (_this D9) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D9_DC20).Cf39
}

func (_this D9) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D9_DC21).Cf40
}

func (_this D9) Dtor_cf41() bool {
	return _this.Get_().(D9_DC21).Cf41
}

func (_this D9) Dtor_cf31() _dafny.Set {
	return _this.Get_().(D9_DC18).Cf31
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + data.Cf37.VerbatimString(true) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC18:
		{
			return "D9.DC18" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34 == data2.Cf34 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Equals(data2.Cf36)
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf37.Equals(data2.Cf37) && data1.Cf38 == data2.Cf38 && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41 == data2.Cf41
		}
	case D9_DC18:
		{
			data2, ok := other.Get_().(D9_DC18)
			return ok && data1.Cf31.Equals(data2.Cf31)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D9) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D9)
	return ok && _this.Equals(typed)
}

func Type_D9_() _dafny.TypeDescriptor {
	return type_D9_{}
}

type type_D9_ struct {
}

func (_this type_D9_) Default() interface{} {
	return Companion_D9_.Default()
}

func (_this type_D9_) String() string {
	return "main.D9"
}
func (_this D9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D9{}

// End of datatype D9

// Definition of trait T0
type T0 interface {
	String() string
	F21() _dafny.Sequence
	F21_set_(value _dafny.Sequence)
	F22() bool
	F22_set_(value bool)
	Fm0(globalState *GlobalState) _dafny.Set
	Fm1(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) bool
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
	F21() _dafny.Sequence
	F21_set_(value _dafny.Sequence)
	F22() bool
	F22_set_(value bool)
	Fm0(globalState *GlobalState) _dafny.Set
	Fm1(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) bool
	F23() _dafny.Sequence
	F23_set_(value _dafny.Sequence)
	Fm2(p0 bool, globalState *GlobalState) bool
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

// Definition of trait T2
type T2 interface {
	String() string
	F21() _dafny.Sequence
	F21_set_(value _dafny.Sequence)
	F22() bool
	F22_set_(value bool)
	F23() _dafny.Sequence
	F23_set_(value _dafny.Sequence)
	Fm0(globalState *GlobalState) _dafny.Set
	Fm1(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) bool
	Fm2(p0 bool, globalState *GlobalState) bool
	M1(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.Int
}
type CompanionStruct_T2_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T2_ = CompanionStruct_T2_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T2_) CastTo_(x interface{}) T2 {
	var t T2
	t, _ = x.(T2)
	return t
}

// End of trait T2

// Definition of class GlobalState
type GlobalState struct {
	F3   _dafny.Int
	F6   _dafny.Sequence
	F8   _dafny.Sequence
	F12  _dafny.CodePoint
	F19  _dafny.MultiSet
	F20  _dafny.Int
	_f0  bool
	_f1  _dafny.Map
	_f2  _dafny.Int
	_f4  _dafny.Int
	_f5  bool
	_f7  bool
	_f9  _dafny.Sequence
	_f10 _dafny.Int
	_f11 bool
	_f13 _dafny.Int
	_f14 _dafny.Int
	_f15 _dafny.Int
	_f16 _dafny.Array
	_f17 _dafny.Array
	_f18 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = _dafny.Zero
	_this.F6 = _dafny.EmptySeq
	_this.F8 = _dafny.EmptySeq
	_this.F12 = _dafny.CodePoint('D')
	_this.F19 = _dafny.EmptyMultiSet
	_this.F20 = _dafny.Zero
	_this._f0 = false
	_this._f1 = _dafny.EmptyMap
	_this._f2 = _dafny.Zero
	_this._f4 = _dafny.Zero
	_this._f5 = false
	_this._f7 = false
	_this._f9 = _dafny.EmptySeq
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f13 = _dafny.Zero
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f17 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Map, f2 _dafny.Int, f3 _dafny.Int, f4 _dafny.Int, f5 bool, f6 _dafny.Sequence, f7 bool, f8 _dafny.Sequence, f9 _dafny.Sequence, f10 _dafny.Int, f11 bool, f12 _dafny.CodePoint, f13 _dafny.Int, f14 _dafny.Int, f15 _dafny.Int, f16 _dafny.Array, f17 _dafny.Array, f18 bool, f19 _dafny.MultiSet, f20 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this).F20 = f20
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Map {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
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
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() _dafny.Sequence {
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
func (_this *GlobalState) F13() _dafny.Int {
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
func (_this *GlobalState) F16() _dafny.Array {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Array {
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
	F26 _dafny.Set
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F26 = _dafny.EmptySet
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

func (_this *C0) Ctor__(f26 _dafny.Set) {
	{
		(_this).F26 = f26
	}
}
func (_this *C0) Fm7(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(false) || ((false) && (true))
	}
}
func (_this *C0) Fm8(globalState *GlobalState) _dafny.Set {
	{
		return (Companion_D1_.Create_DC3_(_dafny.SetOf(_dafny.SetOf(_dafny.CodePoint('r'))))).Dtor_cf3()
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F24  _dafny.Array
	_f25 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F24 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f25 = false
	return &_this
}

type CompanionStruct_C1_ struct {
}

var Companion_C1_ = CompanionStruct_C1_{}

func (_this *C1) Equals(other *C1) bool {
	return _this == other
}

func (_this *C1) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C1)
	return ok && _this.Equals(other)
}

func (*C1) String() string {
	return "_module.C1"
}

func Type_C1_() _dafny.TypeDescriptor {
	return type_C1_{}
}

type type_C1_ struct {
}

func (_this type_C1_) Default() interface{} {
	return (*C1)(nil)
}

func (_this type_C1_) String() string {
	return "main.C1"
}
func (_this *C1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__(f24 _dafny.Array, f25 bool) {
	{
		(_this).F24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C1) M3(p0 bool, p1 bool, p2 bool, globalState *GlobalState) {
	{
		var _278_v0 bool
		_ = _278_v0
		_278_v0 = true
		var _279_v1 _dafny.Int
		_ = _279_v1
		_279_v1 = _dafny.IntOfInt64(976)
		var _280_v2 _dafny.Set
		_ = _280_v2
		_280_v2 = _dafny.SetOf(_279_v1, _279_v1, _279_v1)
		_278_v0 = (p1) || (!(_280_v2).Equals(_280_v2))
		var _281_v3 _dafny.Array
		_ = _281_v3
		var _nw49 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
		_ = _nw49
		_281_v3 = _nw49
		var _282_v4 D0
		_ = _282_v4
		_282_v4 = Companion_D0_.Create_DC0_(_281_v3)
		var _pat_let_tv15 = _281_v3
		_ = _pat_let_tv15
		_282_v4 = func(_pat_let2_0 D0) D0 {
			return func(_283_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let3_0 _dafny.Array) D0 {
					return func(_284_dt__update_hcf0_h0 _dafny.Array) D0 {
						return Companion_D0_.Create_DC0_(_284_dt__update_hcf0_h0)
					}(_pat_let3_0)
				}(_pat_let_tv15)
			}(_pat_let2_0)
		}(_282_v4)
		var _285_v5 _dafny.Sequence
		_ = _285_v5
		_285_v5 = _dafny.SeqOf(p0, (_this).F25())
		var _286_v6 _dafny.MultiSet
		_ = _286_v6
		_286_v6 = _dafny.MultiSetOf(!(!(_278_v0)), (_this).F25(), p0, (_285_v5).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm6(globalState), _dafny.IntOfUint32((_285_v5).Cardinality()))).Uint32()).(bool), (_this).F25())
		var _287_v7 _dafny.Sequence
		_ = _287_v7
		_287_v7 = _dafny.UnicodeSeqOfUtf8Bytes("ekkjfqs")
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))
		_ = _index37
		(_281_v3).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_287_v7, _287_v7), (_index37).Int())
		var _288_v8 _dafny.Sequence
		_ = _288_v8
		_288_v8 = _dafny.SeqOf(_285_v5, _285_v5, _285_v5)
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))
		_ = _index38
		var _rhs40 _dafny.MultiSet = _286_v6
		_ = _rhs40
		var _rhs41 bool = _dafny.Companion_Sequence_.Equal(_285_v5, (_288_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_287_v7).Cardinality()), _dafny.IntOfUint32((_288_v8).Cardinality()))).Uint32()).(_dafny.Sequence))
		_ = _rhs41
		var _rhs42 bool = p0
		_ = _rhs42
		var _lhs36 _dafny.Array = _281_v3
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))
		_ = _lhs37
		_286_v6 = _rhs40
		_278_v0 = _rhs41
		(_lhs36).ArraySet1(_rhs42, (_lhs37).Int())
		var _289_i0 _dafny.Int
		_ = _289_i0
		_289_i0 = _dafny.Zero
		{
			for (_281_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))).Int()).(bool) {
				{
					if (_289_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_289_i0 = (_289_i0).Plus(_dafny.One)
					var _290_v9 _dafny.Array
					_ = _290_v9
					var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
					_ = _nw50
					_290_v9 = _nw50
					_290_v9 = _290_v9
					_281_v3 = _281_v3
					var _291_v10 D0
					_ = _291_v10
					_291_v10 = Companion_D0_.Create_DC1_((_this).F25())
					if (_291_v10).Dtor_cf1() {
						(globalState).F3 = ((_dafny.SetOf((_dafny.Zero).Minus(_279_v1))).Cardinality()).Plus((_dafny.MultiSetOf(_279_v1)).Cardinality())
						_278_v0 = p0
						var _292_v11 _dafny.Set
						_ = _292_v11
						_292_v11 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("u"), _dafny.UnicodeSeqOfUtf8Bytes("px"))
						var _293_v12 *C0
						_ = _293_v12
						var _nw51 *C0 = New_C0_()
						_ = _nw51
						_nw51.Ctor__(_292_v11)
						_293_v12 = _nw51
						(globalState).F3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_287_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(687))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg11 _dafny.Int) interface{} {
								return coer11(arg11)
							}
						}(func(_294_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('c')
						})))).Cardinality())
						var _295_v13 _dafny.Map
						_ = _295_v13
						_295_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _279_v1)
						var _296_v14 _dafny.Sequence
						_ = _296_v14
						_296_v14 = _dafny.SeqOf(_279_v1, (Companion_Default___.Fm9(_278_v0, p2, globalState)).Cardinality())
						var _297_v15 _dafny.Map
						_ = _297_v15
						_297_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v1, (_dafny.MultiSetFromSeq(_296_v14)).Update(_279_v1, Companion_Default___.Abs(_279_v1)))
						var _298_v16 _dafny.MultiSet
						_ = _298_v16
						_298_v16 = _dafny.MultiSetOf(_dafny.IntOfUint32((_285_v5).Cardinality()))
						var _299_v17 D1
						_ = _299_v17
						_299_v17 = Companion_D1_.Create_DC5_((_291_v10).Dtor_cf1(), _295_v13, (_dafny.Zero).Minus(_dafny.IntOfUint32((_287_v7).Cardinality())), (func() _dafny.MultiSet {
							if (_297_v15).Contains(_279_v1) {
								return (_297_v15).Get(_279_v1).(_dafny.MultiSet)
							}
							return _298_v16
						})())
						_299_v17 = _299_v17
					} else {
						var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))
						_ = _index39
						(_281_v3).ArraySet1(!(p0) || (p0), (_index39).Int())
						_279_v1 = (func() _dafny.Int {
							if p1 {
								return _dafny.IntOfInt64(515)
							}
							return _279_v1
						})()
						var _300_v18 _dafny.Set
						_ = _300_v18
						_300_v18 = _dafny.SetOf(_287_v7)
						var _301_v19 *C0
						_ = _301_v19
						var _nw52 *C0 = New_C0_()
						_ = _nw52
						_nw52.Ctor__(_300_v18)
						_301_v19 = _nw52
						(globalState).F20 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(219), _279_v1)
						_278_v0 = _278_v0
					}
					var _source4 D1 = Companion_Default___.Fm10(globalState)
					_ = _source4
					if _source4.Is_DC4() {
						var _302___mcc_h0 _dafny.CodePoint = _source4.Get_().(D1_DC4).Cf4
						_ = _302___mcc_h0
						var _303___mcc_h1 _dafny.Map = _source4.Get_().(D1_DC4).Cf5
						_ = _303___mcc_h1
						var _304___mcc_h2 _dafny.Sequence = _source4.Get_().(D1_DC4).Cf6
						_ = _304___mcc_h2
						var _305_cf6 _dafny.Sequence = _304___mcc_h2
						_ = _305_cf6
						var _306_cf5 _dafny.Map = _303___mcc_h1
						_ = _306_cf5
						var _307_cf4 _dafny.CodePoint = _302___mcc_h0
						_ = _307_cf4
						var _308_v20 _dafny.Array
						_ = _308_v20
						var _len0_3 _dafny.Int = _dafny.IntOfInt64(20)
						_ = _len0_3
						var _nw53 _dafny.Array
						_ = _nw53
						if _len0_3.Cmp(_dafny.Zero) == 0 {
							_nw53 = _dafny.NewArray(_len0_3)
						} else {
							var _init3 func(_dafny.Int) D0 = (func(_309_v10 D0, _310_p1 bool) func(_dafny.Int) D0 {
								return func(_311_i2 _dafny.Int) D0 {
									return func(_pat_let4_0 D0) D0 {
										return func(_312_dt__update__tmp_h1 D0) D0 {
											return func(_pat_let5_0 bool) D0 {
												return func(_313_dt__update_hcf1_h0 bool) D0 {
													return Companion_D0_.Create_DC1_(_313_dt__update_hcf1_h0)
												}(_pat_let5_0)
											}(_310_p1)
										}(_pat_let4_0)
									}(_309_v10)
								}
							})(_291_v10, p1)
							_ = _init3
							var _element0_3 = _init3(_dafny.Zero)
							_ = _element0_3
							_nw53 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
							(_nw53).ArraySet1(_element0_3, 0)
							var _nativeLen0_3 = (_len0_3).Int()
							_ = _nativeLen0_3
							for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
								(_nw53).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
							}
						}
						_308_v20 = _nw53
						var _314_v21 _dafny.Map
						_ = _314_v21
						_314_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v20, _279_v1)
						var _315_v22 _dafny.Map
						_ = _315_v22
						_315_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_281_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))).Int()).(bool), _314_v21)
						_315_v22 = (_315_v22).Update((_this).F25(), _314_v21)
						var _316_v23 _dafny.Array
						_ = _316_v23
						var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
						_ = _nw54
						_316_v23 = _nw54
						var _317_v24 _dafny.Map
						_ = _317_v24
						_317_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _316_v23)
						_317_v24 = (_317_v24).Update(p0, _316_v23)
						var _318_v25 _dafny.Array
						_ = _318_v25
						var _len0_4 _dafny.Int = _dafny.IntOfInt64(8)
						_ = _len0_4
						var _nw55 _dafny.Array
						_ = _nw55
						if _len0_4.Cmp(_dafny.Zero) == 0 {
							_nw55 = _dafny.NewArray(_len0_4)
						} else {
							var _init4 func(_dafny.Int) _dafny.Sequence = (func(_319_cf4 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
								return func(_320_i3 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-865))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg12 _dafny.Int) interface{} {
											return coer12(arg12)
										}
									}((func(_321_cf4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
										return func(_322_i4 _dafny.Int) _dafny.CodePoint {
											return _321_cf4
										}
									})(_319_cf4)))
								}
							})(_307_cf4)
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
						_318_v25 = _nw55
						var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_318_v25), 0))
						_ = _index40
						(_318_v25).ArraySet1(Companion_Default___.Fm11(false, globalState), (_index40).Int())
						var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_318_v25), 0))
						_ = _index41
						(_318_v25).ArraySet1(_287_v7, (_index41).Int())
						_279_v1 = (_dafny.Zero).Minus(((_279_v1).Minus(_279_v1)).Minus((_305_cf6).Select((Companion_Default___.SafeIndex(_279_v1, _dafny.IntOfUint32((_305_cf6).Cardinality()))).Uint32()).(_dafny.Int)))
					} else if _source4.Is_DC5() {
						var _323___mcc_h3 bool = _source4.Get_().(D1_DC5).Cf7
						_ = _323___mcc_h3
						var _324___mcc_h4 _dafny.Map = _source4.Get_().(D1_DC5).Cf8
						_ = _324___mcc_h4
						var _325___mcc_h5 _dafny.Int = _source4.Get_().(D1_DC5).Cf9
						_ = _325___mcc_h5
						var _326___mcc_h6 _dafny.MultiSet = _source4.Get_().(D1_DC5).Cf10
						_ = _326___mcc_h6
						var _327_cf10 _dafny.MultiSet = _326___mcc_h6
						_ = _327_cf10
						var _328_cf9 _dafny.Int = _325___mcc_h5
						_ = _328_cf9
						var _329_cf8 _dafny.Map = _324___mcc_h4
						_ = _329_cf8
						var _330_cf7 bool = _323___mcc_h3
						_ = _330_cf7
						var _331_v26 _dafny.CodePoint
						_ = _331_v26
						_331_v26 = _dafny.CodePoint('t')
						_278_v0 = (func() bool {
							if p1 {
								return !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("bdncgjj"), _331_v26)
							}
							return (_281_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))).Int()).(bool)
						})()
						var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))
						_ = _index42
						(_281_v3).ArraySet1(((_280_v2).Difference(_280_v2)).IsProperSubsetOf(_280_v2), (_index42).Int())
						var _332_v27 _dafny.Sequence
						_ = _332_v27
						_332_v27 = _dafny.SeqOf(_328_cf9, _328_cf9)
						var _333_v28 _dafny.MultiSet
						_ = _333_v28
						_333_v28 = _dafny.MultiSetOf(_286_v6)
						var _334_v29 _dafny.Map
						_ = _334_v29
						_334_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v1, _328_cf9)
						var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))
						_ = _index43
						(_281_v3).ArraySet1(Companion_Default___.Fm12(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_332_v27, _332_v27)).Cardinality()), ((_286_v6).Update(_278_v0, Companion_Default___.Abs(_328_cf9))).Intersection(_286_v6), (_281_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))).Int()).(bool), _dafny.MultiSetOf(_dafny.MultiSetOf(Companion_Default___.Fm12(_328_cf9, _286_v6, p2, _333_v28, globalState), p2), _dafny.MultiSetFromSeq(_285_v5), _dafny.MultiSetOf((_285_v5).Select((Companion_Default___.SafeIndex((_332_v27).Select((Companion_Default___.SafeIndex((_334_v29).Cardinality(), _dafny.IntOfUint32((_332_v27).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_285_v5).Cardinality()))).Uint32()).(bool), _278_v0), _dafny.MultiSetOf(p0, (_this).F25())), globalState), (_index43).Int())
						var _335_v30 _dafny.Map
						_ = _335_v30
						_335_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_cf9, p2)
						var _336_v31 _dafny.Map
						_ = _336_v31
						_336_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_cf10, Companion_Default___.Fm11((func() bool {
							if (_335_v30).Contains(_279_v1) {
								return (_335_v30).Get(_279_v1).(bool)
							}
							return _278_v0
						})(), globalState))
						_336_v31 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_cf10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg13 _dafny.Int) interface{} {
								return coer13(arg13)
							}
						}((func(_337_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_338_i5 _dafny.Int) _dafny.CodePoint {
								return _337_v26
							}
						})(_331_v26))))).Merge(_336_v31)
					} else {
						var _339___mcc_h7 _dafny.Set = _source4.Get_().(D1_DC3).Cf3
						_ = _339___mcc_h7
						var _340_cf3 _dafny.Set = _339___mcc_h7
						_ = _340_cf3
						_278_v0 = ((_279_v1).Times(_279_v1)).Cmp(_279_v1) <= 0
						(globalState).F20 = Companion_Default___.Fm6(globalState)
						(globalState).F3 = _279_v1
						var _341_v32 _dafny.Set
						_ = _341_v32
						_341_v32 = _dafny.SetOf(_287_v7)
						var _342_v33 *C0
						_ = _342_v33
						var _nw56 *C0 = New_C0_()
						_ = _nw56
						_nw56.Ctor__((_341_v32).Intersection(_341_v32))
						_342_v33 = _nw56
					}
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _343_v34 _dafny.Map
		_ = _343_v34
		_343_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_285_v5).Cardinality()), _dafny.IntOfUint32((_287_v7).Cardinality()))
		var _344_v35 _dafny.Sequence
		_ = _344_v35
		_344_v35 = _dafny.SeqOf(_343_v34)
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))
		_ = _index44
		(_281_v3).ArraySet1(!(func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(((_344_v35).Select((Companion_Default___.SafeIndex(_279_v1, _dafny.IntOfUint32((_344_v35).Cardinality()))).Uint32()).(_dafny.Map)).Keys().Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _345_v36 _dafny.Int
				_345_v36 = interface{}(_compr_10).(_dafny.Int)
				if ((_344_v35).Select((Companion_Default___.SafeIndex(_279_v1, _dafny.IntOfUint32((_344_v35).Cardinality()))).Uint32()).(_dafny.Map)).Contains(_345_v36) {
					_coll10.Add((_345_v36).Minus((_dafny.SetOf((Companion_D0_.Create_DC1_(true)).Dtor_cf1())).Cardinality()))
				}
			}
			return _coll10.ToSet()
		}()).Equals((_280_v2).Union(_280_v2)), (_index44).Int())
		var _346_v37 _dafny.Map
		_ = _346_v37
		_346_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _279_v1)
		var _347_v38 _dafny.MultiSet
		_ = _347_v38
		_347_v38 = _dafny.MultiSetOf((_280_v2).Cardinality())
		var _348_v39 D1
		_ = _348_v39
		_348_v39 = Companion_D1_.Create_DC5_((_281_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_281_v3), 0))).Int()).(bool), _346_v37, _279_v1, _347_v38)
		var _pat_let_tv16 = _281_v3
		_ = _pat_let_tv16
		var _pat_let_tv17 = _281_v3
		_ = _pat_let_tv17
		var _pat_let_tv18 = _279_v1
		_ = _pat_let_tv18
		var _pat_let_tv19 = _279_v1
		_ = _pat_let_tv19
		(globalState).F3 = func(_source5 D1) _dafny.Int {
			if _source5.Is_DC4() {
				var _349___mcc_h8 _dafny.CodePoint = _source5.Get_().(D1_DC4).Cf4
				_ = _349___mcc_h8
				var _350___mcc_h9 _dafny.Map = _source5.Get_().(D1_DC4).Cf5
				_ = _350___mcc_h9
				var _351___mcc_h10 _dafny.Sequence = _source5.Get_().(D1_DC4).Cf6
				_ = _351___mcc_h10
				var _352_cf6 _dafny.Sequence = _351___mcc_h10
				_ = _352_cf6
				var _353_cf5 _dafny.Map = _350___mcc_h9
				_ = _353_cf5
				var _354_cf4 _dafny.CodePoint = _349___mcc_h8
				_ = _354_cf4
				return _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_pat_let_tv17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_pat_let_tv16), 0))).Int()).(bool) {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(478))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg14 _dafny.Int) interface{} {
								return coer14(arg14)
							}
						}((func(_355_cf6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_356_i6 _dafny.Int) _dafny.Sequence {
								return _355_cf6
							}
						})(_352_cf6)))
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(437))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg15 _dafny.Int) interface{} {
							return coer15(arg15)
						}
					}((func(_357_cf4 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_358_i7 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf((_dafny.SetOf(_357_cf4)).Cardinality())
						}
					})(_354_cf4)))
				})()).Cardinality())
			} else if _source5.Is_DC5() {
				var _359___mcc_h11 bool = _source5.Get_().(D1_DC5).Cf7
				_ = _359___mcc_h11
				var _360___mcc_h12 _dafny.Map = _source5.Get_().(D1_DC5).Cf8
				_ = _360___mcc_h12
				var _361___mcc_h13 _dafny.Int = _source5.Get_().(D1_DC5).Cf9
				_ = _361___mcc_h13
				var _362___mcc_h14 _dafny.MultiSet = _source5.Get_().(D1_DC5).Cf10
				_ = _362___mcc_h14
				var _363_cf10 _dafny.MultiSet = _362___mcc_h14
				_ = _363_cf10
				var _364_cf9 _dafny.Int = _361___mcc_h13
				_ = _364_cf9
				var _365_cf8 _dafny.Map = _360___mcc_h12
				_ = _365_cf8
				var _366_cf7 bool = _359___mcc_h11
				_ = _366_cf7
				return _pat_let_tv18
			} else {
				var _367___mcc_h15 _dafny.Set = _source5.Get_().(D1_DC3).Cf3
				_ = _367___mcc_h15
				var _368_cf3 _dafny.Set = _367___mcc_h15
				_ = _368_cf3
				return (_pat_let_tv19).Minus(_dafny.IntOfInt64(585))
			}
		}(_348_v39)
	}
}
func (_this *C1) F25() bool {
	{
		return _this._f25
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f21 _dafny.Sequence
	_f22 bool
	_f23 _dafny.Sequence
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f21 = _dafny.EmptySeq
	_this._f22 = false
	_this._f23 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C2_ struct {
}

var Companion_C2_ = CompanionStruct_C2_{}

func (_this *C2) Equals(other *C2) bool {
	return _this == other
}

func (_this *C2) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C2)
	return ok && _this.Equals(other)
}

func (*C2) String() string {
	return "_module.C2"
}

func Type_C2_() _dafny.TypeDescriptor {
	return type_C2_{}
}

type type_C2_ struct {
}

func (_this type_C2_) Default() interface{} {
	return (*C2)(nil)
}

func (_this type_C2_) String() string {
	return "main.C2"
}
func (_this *C2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C2{}
var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F21() _dafny.Sequence {
	return _this._f21
}
func (_this *C2) F21_set_(value _dafny.Sequence) {
	_this._f21 = value
}
func (_this *C2) F22() bool {
	return _this._f22
}
func (_this *C2) F22_set_(value bool) {
	_this._f22 = value
}
func (_this *C2) F23() _dafny.Sequence {
	return _this._f23
}
func (_this *C2) F23_set_(value _dafny.Sequence) {
	_this._f23 = value
}
func (_this *C2) Ctor__(f23 _dafny.Sequence, f21 _dafny.Sequence, f22 bool) {
	{
		(_this)._f23 = f23
		(_this)._f21 = f21
		(_this)._f22 = f22
	}
}
func (_this *C2) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return !((func() bool {
			if false {
				return _this.F22()
			}
			return _this.F22()
		})()) || (!(false))
	}
}
func (_this *C2) Fm0(globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf(_this.F22(), _this.F22(), false, _this.F22(), _this.F22())
	}
}
func (_this *C2) Fm1(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.IntOfInt64(99)).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_this.F22())).Cardinality())).Cardinality()))).Cmp(((_dafny.SetOf(_this.F22(), _this.F22())).Union(_dafny.SetOf(true))).Cardinality()) == 0
	}
}
func (_this *C2) Fm3(globalState *GlobalState) bool {
	{
		return !(((func() _dafny.Int {
			if _this.F22() {
				return _dafny.IntOfInt64(-141)
			}
			return _dafny.IntOfInt64(419)
		})()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_this.F21())).Cardinality())) <= 0)
	}
}
func (_this *C2) M1(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _369_v0 _dafny.Int
		_ = _369_v0
		_369_v0 = _dafny.IntOfInt64(191)
		var _370_v1 _dafny.Sequence
		_ = _370_v1
		_370_v1 = _dafny.SeqOf(_369_v0, _369_v0)
		var _371_v2 _dafny.MultiSet
		_ = _371_v2
		_371_v2 = _dafny.MultiSetOf((_this.F23()).Select((Companion_Default___.SafeIndex(_369_v0, _dafny.IntOfUint32((_this.F23()).Cardinality()))).Uint32()).(_dafny.Sequence), _370_v1)
		var _372_v3 _dafny.Set
		_ = _372_v3
		_372_v3 = _dafny.SetOf((func() _dafny.Int {
			if (_371_v2).Contains(_370_v1) {
				return (_371_v2).Multiplicity(_370_v1)
			}
			return _369_v0
		})())
		var _373_v5 _dafny.Set
		_ = _373_v5
		_373_v5 = _dafny.SetOf(_372_v3, _372_v3, _372_v3, (func() _dafny.Set {
			var _coll11 = _dafny.NewBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(132), _dafny.IntOfInt64(989))); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _374_v4 _dafny.Int
				_374_v4 = interface{}(_compr_11).(_dafny.Int)
				if ((_dafny.IntOfInt64(132)).Cmp(_374_v4) <= 0) && ((_374_v4).Cmp(_dafny.IntOfInt64(989)) < 0) {
					_coll11.Add((_374_v4).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_369_v0, p0)).Cardinality()))
				}
			}
			return _coll11.ToSet()
		}()).Difference(_372_v3))
		_373_v5 = ((func() _dafny.Set {
			if p0 {
				return _dafny.SetOf(_372_v3)
			}
			return _373_v5
		})()).Union(_373_v5)
		if true {
			(globalState).F20 = _369_v0
			var _375_v6 _dafny.Map
			_ = _375_v6
			_375_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(_369_v0, _dafny.IntOfUint32((_dafny.SeqOf(_this.F22())).Cardinality()), _369_v0))).Cardinality(), _this.F21())
			var _376_v7 _dafny.CodePoint
			_ = _376_v7
			_376_v7 = _dafny.CodePoint('u')
			var _377_v8 _dafny.Sequence
			_ = _377_v8
			_377_v8 = _dafny.SeqOf(_this.F21(), (func() _dafny.Sequence {
				if (_375_v6).Contains(_369_v0) {
					return (_375_v6).Get(_369_v0).(_dafny.Sequence)
				}
				return _this.F21()
			})(), _dafny.SeqOf(_376_v7))
			(_this).F21_set_(_dafny.Companion_Sequence_.Update((_377_v8).Select((Companion_Default___.SafeIndex(_369_v0, _dafny.IntOfUint32((_377_v8).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_369_v0, _dafny.IntOfUint32(((_377_v8).Select((Companion_Default___.SafeIndex(_369_v0, _dafny.IntOfUint32((_377_v8).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _376_v7))
			var _378_v9 _dafny.Array
			_ = _378_v9
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_5
			var _nw57 _dafny.Array
			_ = _nw57
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw57 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = func(_379_i0 _dafny.Int) bool {
					return true
				}
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw57 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw57).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw57).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_378_v9 = _nw57
			var _380_v10 D0
			_ = _380_v10
			_380_v10 = Companion_D0_.Create_DC0_(_378_v9)
			var _381_v11 _dafny.Sequence
			_ = _381_v11
			_381_v11 = _dafny.SeqOf(_378_v9)
			var _382_v12 _dafny.Array
			_ = _382_v12
			var _nwElement0_9 _dafny.Array = _378_v9
			_ = _nwElement0_9
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(6))
			_ = _nw58
			(_nw58).ArraySet1(_nwElement0_9, 0)
			(_nw58).ArraySet1(_378_v9, 1)
			(_nw58).ArraySet1(_378_v9, 2)
			(_nw58).ArraySet1((_380_v10).Dtor_cf0(), 3)
			(_nw58).ArraySet1(_378_v9, 4)
			(_nw58).ArraySet1((_381_v11).Select((Companion_Default___.SafeIndex(_369_v0, _dafny.IntOfUint32((_381_v11).Cardinality()))).Uint32()).(_dafny.Array), 5)
			_382_v12 = _nw58
			var _383_v13 D0
			_ = _383_v13
			_383_v13 = Companion_D0_.Create_DC1_(_this.F22())
			var _384_v14 _dafny.MultiSet
			_ = _384_v14
			_384_v14 = _dafny.MultiSetOf(_this.F22())
			var _385_v15 _dafny.Map
			_ = _385_v15
			_385_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_this.F22(), _384_v14, globalState), _382_v12)
			_382_v12 = (func() _dafny.Array {
				if (func(_pat_let6_0 D0) D0 {
					return func(_386_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let7_0 bool) D0 {
							return func(_387_dt__update_hcf1_h0 bool) D0 {
								return Companion_D0_.Create_DC1_(_387_dt__update_hcf1_h0)
							}(_pat_let7_0)
						}(false)
					}(_pat_let6_0)
				}(_383_v13)).Dtor_cf1() {
					return (func() _dafny.Array {
						if (_385_v15).Contains(_dafny.CodePoint('x')) {
							return (_385_v15).Get(_dafny.CodePoint('x')).(_dafny.Array)
						}
						return _382_v12
					})()
				}
				return _382_v12
			})()
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_382_v12), 0))
			_ = _index45
			(_382_v12).ArraySet1(_378_v9, (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_382_v12), 0))
			_ = _index46
			(_382_v12).ArraySet1(_378_v9, (_index46).Int())
			var _388_v16 _dafny.MultiSet
			_ = _388_v16
			_388_v16 = _dafny.MultiSetOf(_369_v0)
			r0 = ((_388_v16).Cardinality()).Times((_370_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.IntOfUint32((_370_v1).Cardinality()))).Uint32()).(_dafny.Int))
		} else {
			var _389_v17 _dafny.Map
			_ = _389_v17
			_389_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _369_v0)
			var _390_v18 _dafny.Map
			_ = _390_v18
			_390_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_389_v17, !(_this.F22()) || (_this.F22()))
			_390_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_389_v17, true)
			(_this).F22_set_(_this.F22())
			var _391_v19 _dafny.Array
			_ = _391_v19
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw59
			_391_v19 = _nw59
			var _392_v20 bool
			_ = _392_v20
			var _out4 bool
			_ = _out4
			_out4 = Companion_Default___.M0(_391_v19, (true) && (p0), globalState)
			_392_v20 = _out4
			(_this).F22_set_(_392_v20)
			var _393_v21 _dafny.Sequence
			_ = _393_v21
			_393_v21 = _dafny.SeqOf(_372_v3, _372_v3)
			(globalState).F20 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_393_v21, _dafny.Companion_Sequence_.Update(_393_v21, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.IntOfUint32((_393_v21).Cardinality()))).Uint32(), _dafny.SetOf(_369_v0, (_dafny.Zero).Minus(_369_v0))))).Cardinality())
		}
		var _394_v22 D0
		_ = _394_v22
		_394_v22 = Companion_D0_.Create_DC1_(true)
		var _395_v23 D0
		_ = _395_v23
		_395_v23 = Companion_D0_.Create_DC2_(_394_v22)
		var _396_v24 _dafny.Sequence
		_ = _396_v24
		_396_v24 = _dafny.SeqOf(_this.F21())
		var _397_v25 _dafny.Sequence
		_ = _397_v25
		_397_v25 = _dafny.SeqOf(_this.F22(), _this.F22(), !(p0))
		var _398_v26 _dafny.Array
		_ = _398_v26
		var _nwElement0_10 _dafny.Int = _dafny.IntOfUint32((_396_v24).Cardinality())
		_ = _nwElement0_10
		var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(15))
		_ = _nw60
		(_nw60).ArraySet1(_nwElement0_10, 0)
		(_nw60).ArraySet1(_369_v0, 1)
		(_nw60).ArraySet1(_369_v0, 2)
		(_nw60).ArraySet1(_369_v0, 3)
		(_nw60).ArraySet1(_dafny.IntOfInt64(426), 4)
		(_nw60).ArraySet1(Companion_Default___.SafeModuloInt(_369_v0, _369_v0), 5)
		(_nw60).ArraySet1((_dafny.Zero).Minus(_369_v0), 6)
		(_nw60).ArraySet1(_dafny.IntOfInt64(834), 7)
		(_nw60).ArraySet1(_369_v0, 8)
		(_nw60).ArraySet1(_dafny.IntOfUint32((_396_v24).Cardinality()), 9)
		(_nw60).ArraySet1(_369_v0, 10)
		(_nw60).ArraySet1(_dafny.IntOfInt64(-412), 11)
		(_nw60).ArraySet1(_369_v0, 12)
		(_nw60).ArraySet1(_369_v0, 13)
		(_nw60).ArraySet1(_dafny.IntOfUint32((_397_v25).Cardinality()), 14)
		_398_v26 = _nw60
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))
		_ = _index47
		(_398_v26).ArraySet1((_dafny.IntOfUint32((_370_v1).Cardinality())).Times(_369_v0), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))
		_ = _index48
		var _rhs43 _dafny.Int = _dafny.IntOfUint32((_this.F23()).Cardinality())
		_ = _rhs43
		var _rhs44 _dafny.Int = _369_v0
		_ = _rhs44
		var _rhs45 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_369_v0))
		_ = _rhs45
		var _rhs46 D0 = _395_v23
		_ = _rhs46
		var _rhs47 _dafny.Int = Companion_Default___.SafeDivisionInt(_369_v0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_397_v25, _397_v25)).Cardinality()))))
		_ = _rhs47
		var _lhs38 *GlobalState = globalState
		_ = _lhs38
		var _lhs39 *GlobalState = globalState
		_ = _lhs39
		var _lhs40 _dafny.Array = _398_v26
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))
		_ = _lhs41
		_lhs38.F20 = _rhs43
		_369_v0 = _rhs44
		_lhs39.F20 = _rhs45
		_395_v23 = _rhs46
		(_lhs40).ArraySet1(_rhs47, (_lhs41).Int())
		r0 = (_398_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))).Int()).(_dafny.Int)
		var _399_i1 _dafny.Int
		_ = _399_i1
		_399_i1 = _dafny.Zero
		{
			for false {
				{
					if (_399_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_399_i1 = (_399_i1).Plus(_dafny.One)
					var _400_v27 _dafny.CodePoint
					_ = _400_v27
					_400_v27 = _dafny.CodePoint('i')
					var _401_v28 _dafny.Map
					_ = _401_v28
					_401_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_398_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))).Int()).(_dafny.Int), _400_v27)
					var _402_v29 _dafny.Map
					_ = _402_v29
					_402_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22(), (_401_v28).Cardinality())
					var _403_v30 _dafny.Set
					_ = _403_v30
					_403_v30 = _dafny.SetOf(Companion_Default___.Fm5((_398_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _369_v0, globalState))
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))
					_ = _index49
					(_398_v26).ArraySet1((((_402_v29).Cardinality()).Minus(_dafny.IntOfInt64(914))).Minus(((func() _dafny.Set {
						if false {
							return _403_v30
						}
						return _dafny.SetOf(_370_v1, _dafny.SeqOf(_369_v0, (_398_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))).Int()).(_dafny.Int), _369_v0))
					})()).Cardinality()), (_index49).Int())
					var _404_v31 _dafny.Array
					_ = _404_v31
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(25)
					_ = _len0_6
					var _nw61 _dafny.Array
					_ = _nw61
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw61 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.MultiSet = (func(_405_p0 bool) func(_dafny.Int) _dafny.MultiSet {
							return func(_406_i2 _dafny.Int) _dafny.MultiSet {
								return _dafny.MultiSetOf(!(_this.F22()), _405_p0)
							}
						})(p0)
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw61 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw61).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw61).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_404_v31 = _nw61
					var _407_v32 *C1
					_ = _407_v32
					var _nw62 *C1 = New_C1_()
					_ = _nw62
					_nw62.Ctor__(_404_v31, true)
					_407_v32 = _nw62
					var _408_v33 _dafny.Set
					_ = _408_v33
					_408_v33 = _dafny.SetOf(_400_v27)
					var _409_v34 D1
					_ = _409_v34
					_409_v34 = Companion_D1_.Create_DC3_(_dafny.SetOf(_408_v33))
					var _410_v35 _dafny.Set
					_ = _410_v35
					_410_v35 = _dafny.SetOf(_409_v34, _409_v34)
					var _411_v36 _dafny.Set
					_ = _411_v36
					_411_v36 = _dafny.SetOf((_410_v35).Difference(_410_v35), _dafny.SetOf(Companion_Default___.Fm10(globalState), _409_v34, Companion_Default___.Fm10(globalState), _409_v34))
					_411_v36 = _411_v36
					var _412_v37 _dafny.Array
					_ = _412_v37
					var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(20))
					_ = _nw63
					_412_v37 = _nw63
					var _413_v38 _dafny.MultiSet
					_ = _413_v38
					_413_v38 = _dafny.MultiSetOf(_this.F21(), _dafny.Companion_Sequence_.Update(_this.F21(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_397_v25).Cardinality()), _dafny.IntOfUint32((_this.F21()).Cardinality()))).Uint32(), _400_v27), _this.F21(), _this.F21())
					var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_412_v37), 0))
					_ = _index50
					(_412_v37).ArraySet1(_413_v38, (_index50).Int())
					var _414_v39 D0
					_ = _414_v39
					_414_v39 = Companion_D0_.Create_DC1_((_407_v32).F25())
					var _415_v40 _dafny.Set
					_ = _415_v40
					_415_v40 = _dafny.SetOf((_414_v39).Dtor_cf1(), false, !(_this.F22()), p0)
					var _416_v41 _dafny.Map
					_ = _416_v41
					_416_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_407_v32, _415_v40)
					var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_412_v37), 0))
					_ = _index51
					(_412_v37).ArraySet1(Companion_Default___.Fm13(_this.F22(), _415_v40, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_this.F21()).Cardinality()), (_dafny.MultiSetOf((func() _dafny.Set {
						if (_416_v41).Contains(_407_v32) {
							return (_416_v41).Get(_407_v32).(_dafny.Set)
						}
						return _415_v40
					})(), _415_v40, _dafny.SetOf(false, p0, (_407_v32).F25(), (_407_v32).F25(), (_407_v32).F25()))).Cardinality()), p0, globalState), (_index51).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _417_v43 D1
		_ = _417_v43
		_417_v43 = Companion_D1_.Create_DC3_(_dafny.SetOf(func() _dafny.Set {
			var _coll12 = _dafny.NewBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate((_dafny.MultiSetFromSeq(_this.F21())).Elements()); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _418_v42 _dafny.CodePoint
				_418_v42 = interface{}(_compr_12).(_dafny.CodePoint)
				if (_dafny.MultiSetFromSeq(_this.F21())).Contains(_418_v42) {
					_coll12.Add(_418_v42)
				}
			}
			return _coll12.ToSet()
		}()))
		var _419_v44 _dafny.Map
		_ = _419_v44
		_419_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_369_v0, p0)
		var _pat_let_tv20 = _398_v26
		_ = _pat_let_tv20
		var _pat_let_tv21 = _398_v26
		_ = _pat_let_tv21
		var _pat_let_tv22 = _369_v0
		_ = _pat_let_tv22
		var _pat_let_tv23 = _369_v0
		_ = _pat_let_tv23
		var _pat_let_tv24 = _398_v26
		_ = _pat_let_tv24
		var _pat_let_tv25 = _398_v26
		_ = _pat_let_tv25
		var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F21(), _this.F21())
		_ = _rhs48
		var _rhs49 _dafny.Int = func(_source6 D1) _dafny.Int {
			if _source6.Is_DC4() {
				var _420___mcc_h0 _dafny.CodePoint = _source6.Get_().(D1_DC4).Cf4
				_ = _420___mcc_h0
				var _421___mcc_h1 _dafny.Map = _source6.Get_().(D1_DC4).Cf5
				_ = _421___mcc_h1
				var _422___mcc_h2 _dafny.Sequence = _source6.Get_().(D1_DC4).Cf6
				_ = _422___mcc_h2
				var _423_cf6 _dafny.Sequence = _422___mcc_h2
				_ = _423_cf6
				var _424_cf5 _dafny.Map = _421___mcc_h1
				_ = _424_cf5
				var _425_cf4 _dafny.CodePoint = _420___mcc_h0
				_ = _425_cf4
				return ((_pat_let_tv21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_pat_let_tv20), 0))).Int()).(_dafny.Int)).Minus((Companion_D1_.Create_DC5_(_this.F22(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F22()), _dafny.IntOfUint32((_423_cf6).Cardinality())), (_dafny.SetOf(_this.F22(), false)).Cardinality(), _dafny.MultiSetOf(_pat_let_tv22, _dafny.IntOfUint32((_this.F21()).Cardinality())))).Dtor_cf9())
			} else if _source6.Is_DC5() {
				var _426___mcc_h3 bool = _source6.Get_().(D1_DC5).Cf7
				_ = _426___mcc_h3
				var _427___mcc_h4 _dafny.Map = _source6.Get_().(D1_DC5).Cf8
				_ = _427___mcc_h4
				var _428___mcc_h5 _dafny.Int = _source6.Get_().(D1_DC5).Cf9
				_ = _428___mcc_h5
				var _429___mcc_h6 _dafny.MultiSet = _source6.Get_().(D1_DC5).Cf10
				_ = _429___mcc_h6
				var _430_cf10 _dafny.MultiSet = _429___mcc_h6
				_ = _430_cf10
				var _431_cf9 _dafny.Int = _428___mcc_h5
				_ = _431_cf9
				var _432_cf8 _dafny.Map = _427___mcc_h4
				_ = _432_cf8
				var _433_cf7 bool = _426___mcc_h3
				_ = _433_cf7
				return _pat_let_tv23
			} else {
				var _434___mcc_h7 _dafny.Set = _source6.Get_().(D1_DC3).Cf3
				_ = _434___mcc_h7
				var _435_cf3 _dafny.Set = _434___mcc_h7
				_ = _435_cf3
				return ((_pat_let_tv25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_pat_let_tv24), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(443))
			}
		}(_417_v43)
		_ = _rhs49
		var _rhs50 _dafny.Array = _398_v26
		_ = _rhs50
		var _rhs51 _dafny.Int = ((_398_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_398_v26), 0))).Int()).(_dafny.Int)).Minus(_369_v0)
		_ = _rhs51
		var _rhs52 _dafny.Int = ((_419_v44).Merge(_419_v44)).Cardinality()
		_ = _rhs52
		var _lhs42 *C2 = _this
		_ = _lhs42
		var _lhs43 *GlobalState = globalState
		_ = _lhs43
		var _lhs44 *GlobalState = globalState
		_ = _lhs44
		_lhs42.F21_set_(_rhs48)
		r0 = _rhs49
		_398_v26 = _rhs50
		_lhs43.F20 = _rhs51
		_lhs44.F20 = _rhs52
		var _436_v45 _dafny.CodePoint
		_ = _436_v45
		_436_v45 = _dafny.CodePoint('s')
		var _437_v46 _dafny.Set
		_ = _437_v46
		_437_v46 = _dafny.SetOf(_this.F22(), _this.F22(), p0)
		var _438_v47 _dafny.Map
		_ = _438_v47
		_438_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_437_v46).Cardinality(), Companion_Default___.Fm6(globalState))
		var _pat_let_tv26 = p0
		_ = _pat_let_tv26
		var _pat_let_tv27 = p0
		_ = _pat_let_tv27
		var _pat_let_tv28 = p0
		_ = _pat_let_tv28
		var _pat_let_tv29 = p0
		_ = _pat_let_tv29
		var _pat_let_tv30 = p0
		_ = _pat_let_tv30
		r0 = (func(_source7 D1) _dafny.Map {
			if _source7.Is_DC4() {
				var _439___mcc_h8 _dafny.CodePoint = _source7.Get_().(D1_DC4).Cf4
				_ = _439___mcc_h8
				var _440___mcc_h9 _dafny.Map = _source7.Get_().(D1_DC4).Cf5
				_ = _440___mcc_h9
				var _441___mcc_h10 _dafny.Sequence = _source7.Get_().(D1_DC4).Cf6
				_ = _441___mcc_h10
				var _442_cf6 _dafny.Sequence = _441___mcc_h10
				_ = _442_cf6
				var _443_cf5 _dafny.Map = _440___mcc_h9
				_ = _443_cf5
				var _444_cf4 _dafny.CodePoint = _439___mcc_h8
				_ = _444_cf4
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv26, _this.F22())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv27, false))
			} else if _source7.Is_DC5() {
				var _445___mcc_h11 bool = _source7.Get_().(D1_DC5).Cf7
				_ = _445___mcc_h11
				var _446___mcc_h12 _dafny.Map = _source7.Get_().(D1_DC5).Cf8
				_ = _446___mcc_h12
				var _447___mcc_h13 _dafny.Int = _source7.Get_().(D1_DC5).Cf9
				_ = _447___mcc_h13
				var _448___mcc_h14 _dafny.MultiSet = _source7.Get_().(D1_DC5).Cf10
				_ = _448___mcc_h14
				var _449_cf10 _dafny.MultiSet = _448___mcc_h14
				_ = _449_cf10
				var _450_cf9 _dafny.Int = _447___mcc_h13
				_ = _450_cf9
				var _451_cf8 _dafny.Map = _446___mcc_h12
				_ = _451_cf8
				var _452_cf7 bool = _445___mcc_h11
				_ = _452_cf7
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv28, _452_cf7)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F22()))
			} else {
				var _453___mcc_h15 _dafny.Set = _source7.Get_().(D1_DC3).Cf3
				_ = _453___mcc_h15
				var _454_cf3 _dafny.Set = _453___mcc_h15
				_ = _454_cf3
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv29, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22(), _pat_let_tv30))
			}
		}(Companion_D1_.Create_DC4_(_436_v45, _438_v47, _370_v1))).Cardinality()
		return r0
	}
}
func (_this *C2) M2(p0 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _455_v0 _dafny.Array
		_ = _455_v0
		var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw64
		_455_v0 = _nw64
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))
		_ = _index52
		(_455_v0).ArraySet1(_dafny.IntOfUint32((_this.F21()).Cardinality()), (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))
		_ = _index53
		(_455_v0).ArraySet1((p0).Minus(p0), (_index53).Int())
		var _456_v1 _dafny.Sequence
		_ = _456_v1
		_456_v1 = _dafny.SeqOf(_this.F22())
		var _457_v2 _dafny.Set
		_ = _457_v2
		_457_v2 = _dafny.SetOf(true)
		_456_v1 = _dafny.SeqOf(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(350), (_457_v2).Cardinality())).Contains((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)), _this.F22(), (_this.F22()) && (_this.F22()))
		if _this.F22() {
			var _458_v3 _dafny.MultiSet
			_ = _458_v3
			_458_v3 = _dafny.MultiSetOf(_this.F22())
			(globalState).F3 = (func() _dafny.Int {
				if !((func() bool {
					if !(_this.F22()) {
						return _this.F22()
					}
					return _this.F22()
				})()) {
					return Companion_Default___.SafeModuloInt(Companion_Default___.Fm6(globalState), (_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int))
				}
				return (_458_v3).Cardinality()
			})()
			var _459_v4 _dafny.CodePoint
			_ = _459_v4
			_459_v4 = _dafny.CodePoint('u')
			var _460_v5 _dafny.Map
			_ = _460_v5
			_460_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_458_v3).Cardinality(), (_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int))
			var _461_v6 _dafny.Sequence
			_ = _461_v6
			_461_v6 = _dafny.SeqOf((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int))
			var _462_v7 D1
			_ = _462_v7
			_462_v7 = Companion_D1_.Create_DC4_(_459_v4, _460_v5, _461_v6)
			_462_v7 = _462_v7
			var _463_v8 _dafny.Map
			_ = _463_v8
			_463_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int), _dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F22())), _this.F22(), _dafny.MultiSetOf(_dafny.MultiSetOf(false, !(true)), _458_v3, _458_v3, _dafny.MultiSetOf(_this.F22(), _this.F22(), _this.F22()), _458_v3), globalState), (_dafny.Zero).Minus((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)))
			var _464_v9 _dafny.Map
			_ = _464_v9
			_464_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22(), false)
			var _465_v10 D0
			_ = _465_v10
			_465_v10 = Companion_D0_.Create_DC1_((func() bool {
				if (_464_v9).Contains(true) {
					return (_464_v9).Get(true).(bool)
				}
				return _this.F22()
			})())
			_463_v8 = (_463_v8).Update((_465_v10).Dtor_cf1(), (_dafny.Zero).Minus(p0))
			var _466_v11 _dafny.Array
			_ = _466_v11
			var _nwElement0_11 bool = _this.F22()
			_ = _nwElement0_11
			var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(27))
			_ = _nw65
			(_nw65).ArraySet1(_nwElement0_11, 0)
			(_nw65).ArraySet1(_this.F22(), 1)
			(_nw65).ArraySet1(_this.F22(), 2)
			(_nw65).ArraySet1(_this.F22(), 3)
			(_nw65).ArraySet1(_this.F22(), 4)
			(_nw65).ArraySet1(_this.F22(), 5)
			(_nw65).ArraySet1(_this.F22(), 6)
			(_nw65).ArraySet1(_this.F22(), 7)
			(_nw65).ArraySet1(false, 8)
			(_nw65).ArraySet1(_this.F22(), 9)
			(_nw65).ArraySet1(_this.F22(), 10)
			(_nw65).ArraySet1(true, 11)
			(_nw65).ArraySet1(_this.F22(), 12)
			(_nw65).ArraySet1(_this.F22(), 13)
			(_nw65).ArraySet1(_this.F22(), 14)
			(_nw65).ArraySet1(_this.F22(), 15)
			(_nw65).ArraySet1(_this.F22(), 16)
			(_nw65).ArraySet1(true, 17)
			(_nw65).ArraySet1(_this.F22(), 18)
			(_nw65).ArraySet1(false, 19)
			(_nw65).ArraySet1(_this.F22(), 20)
			(_nw65).ArraySet1(_this.F22(), 21)
			(_nw65).ArraySet1(!(_this.F22()), 22)
			(_nw65).ArraySet1(_this.F22(), 23)
			(_nw65).ArraySet1(_this.F22(), 24)
			(_nw65).ArraySet1(false, 25)
			(_nw65).ArraySet1(_this.F22(), 26)
			_466_v11 = _nw65
			var _467_v12 D0
			_ = _467_v12
			_467_v12 = Companion_D0_.Create_DC0_(_466_v11)
			var _468_v13 D0
			_ = _468_v13
			_468_v13 = Companion_D0_.Create_DC2_(_467_v12)
			var _469_v14 _dafny.Map
			_ = _469_v14
			_469_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v13, _dafny.IntOfUint32((_this.F21()).Cardinality()))
			_469_v14 = (_469_v14).Update(_468_v13, p0)
			r1 = _this.F22()
		} else {
			(globalState).F20 = ((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("etxvn")).Cardinality()))
			(globalState).F12 = _dafny.CodePoint('d')
			var _470_v15 _dafny.Array
			_ = _470_v15
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_7
			var _nw66 _dafny.Array
			_ = _nw66
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw66 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Set = (func(_471_v0 _dafny.Array) func(_dafny.Int) _dafny.Set {
					return func(_472_i0 _dafny.Int) _dafny.Set {
						return _dafny.SetOf((_471_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_471_v0), 0))).Int()).(_dafny.Int))
					}
				})(_455_v0)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw66 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw66).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw66).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_470_v15 = _nw66
			var _473_v16 _dafny.Set
			_ = _473_v16
			_473_v16 = _dafny.SetOf((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int), p0)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_470_v15), 0))
			_ = _index54
			(_470_v15).ArraySet1((_473_v16), (_index54).Int())
			var _474_v17 _dafny.Map
			_ = _474_v17
			_474_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm1(_this.F21(), true, false, globalState), _this.F22())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_470_v15), 0))
			_ = _index55
			(_470_v15).ArraySet1(Companion_Default___.Fm14(_474_v17, p0, globalState), (_index55).Int())
			var _source8 _dafny.Set = _473_v16
			_ = _source8
			var _475___mcc_h0 _dafny.Set = _source8
			_ = _475___mcc_h0
			var _476_cf11 _dafny.Set = _475___mcc_h0
			_ = _476_cf11
			(globalState).F20 = (Companion_Default___.Fm6(globalState)).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_this.F21()).Cardinality()), (_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)))
			_473_v16 = _473_v16
			_474_v17 = _474_v17
			(_this).F22_set_(((func() bool {
				if _this.F22() {
					return _this.F22()
				}
				return _this.F22()
			})()) && ((_this).Fm2(_this.F22(), globalState)))
			var _477_v18 _dafny.Array
			_ = _477_v18
			var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(4))
			_ = _nw67
			_477_v18 = _nw67
			var _478_v19 _dafny.Map
			_ = _478_v19
			_478_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_477_v18, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_this.F21()).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_this.F21()).Cardinality()), p0)))
			var _479_v20 _dafny.Map
			_ = _479_v20
			_479_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int), p0)
			_478_v19 = (_478_v19).Update(_477_v18, _479_v20)
		}
		if true {
			var _480_v21 _dafny.Set
			_ = _480_v21
			_480_v21 = _dafny.SetOf((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int))
			var _481_v22 _dafny.Set
			_ = _481_v22
			_481_v22 = (_480_v21).Intersection(_480_v21)
			_481_v22 = _481_v22
			var _482_v23 _dafny.Set
			_ = _482_v23
			_482_v23 = _dafny.SetOf(_this.F21(), _this.F21())
			var _483_v24 *C0
			_ = _483_v24
			var _nw68 *C0 = New_C0_()
			_ = _nw68
			_nw68.Ctor__(_482_v23)
			_483_v24 = _nw68
			_483_v24 = _483_v24
			var _484_v25 _dafny.Array
			_ = _484_v25
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_8
			var _nw69 _dafny.Array
			_ = _nw69
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw69 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) bool = func(_485_i1 _dafny.Int) bool {
					return (func() bool {
						if _this.F22() {
							return _this.F22()
						}
						return _this.F22()
					})()
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw69 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw69).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw69).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_484_v25 = _nw69
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_484_v25), 0))
			_ = _index56
			(_484_v25).ArraySet1(_this.F22(), (_index56).Int())
			var _486_v26 _dafny.MultiSet
			_ = _486_v26
			_486_v26 = _dafny.MultiSetOf(_this.F22(), (_this).Fm1(_this.F21(), _this.F22(), _this.F22(), globalState))
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_484_v25), 0))
			_ = _index57
			(_484_v25).ArraySet1(!(((_dafny.IntOfInt64(211)).Cmp((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)) == 0) == ((_486_v26).IsDisjointFrom(_dafny.MultiSetFromSeq(_456_v1)))), (_index57).Int())
			r1 = _this.F22()
			var _487_v27 _dafny.Array
			_ = _487_v27
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
			_ = _nw70
			_487_v27 = _nw70
			var _488_v28 _dafny.MultiSet
			_ = _488_v28
			_488_v28 = _dafny.MultiSetOf(_this.F21(), Companion_Default___.Fm11(false, globalState), _dafny.UnicodeSeqOfUtf8Bytes("bdbdchmi"))
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_484_v25), 0))
			_ = _index58
			var _rhs53 _dafny.Int = (func() _dafny.Int {
				if (_488_v28).Contains(_this.F21()) {
					return (_488_v28).Multiplicity(_this.F21())
				}
				return (_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)
			})()
			_ = _rhs53
			var _rhs54 bool = (_484_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_484_v25), 0))).Int()).(bool)
			_ = _rhs54
			var _rhs55 bool = (_484_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_484_v25), 0))).Int()).(bool)
			_ = _rhs55
			var _rhs56 _dafny.Array = _487_v27
			_ = _rhs56
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			var _lhs46 _dafny.Array = _484_v25
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_484_v25), 0))
			_ = _lhs47
			var _lhs48 *C2 = _this
			_ = _lhs48
			_lhs45.F20 = _rhs53
			(_lhs46).ArraySet1(_rhs54, (_lhs47).Int())
			_lhs48.F22_set_(_rhs55)
			_487_v27 = _rhs56
		} else {
			var _489_v29 _dafny.Array
			_ = _489_v29
			var _nw71 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw71
			_489_v29 = _nw71
			_489_v29 = _489_v29
			var _490_v30 _dafny.MultiSet
			_ = _490_v30
			_490_v30 = _dafny.MultiSetOf(_this.F22(), false)
			var _491_v32 _dafny.Set
			_ = _491_v32
			_491_v32 = _dafny.SetOf(_490_v30)
			var _492_v33 _dafny.Sequence
			_ = _492_v33
			_492_v33 = _dafny.SeqOf(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_491_v32).Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _493_v31 _dafny.MultiSet
					_493_v31 = interface{}(_compr_13).(_dafny.MultiSet)
					if (_491_v32).Contains(_493_v31) {
						_coll13.Add(_493_v31, (_dafny.Zero).Minus(p0))
					}
				}
				return _coll13.ToMap()
			}())
			var _494_v34 _dafny.Map
			_ = _494_v34
			_494_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_490_v30, (_492_v33).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_492_v33).Cardinality()))).Uint32()).(_dafny.Map))
			_494_v34 = (_494_v34).Update((_dafny.MultiSetOf(_this.F22())).Intersection(_490_v30), (_492_v33).Select((Companion_Default___.SafeIndex((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_492_v33).Cardinality()))).Uint32()).(_dafny.Map))
			_456_v1 = _dafny.Companion_Sequence_.Concatenate(_456_v1, _456_v1)
			if !(_this.F22()) {
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))
				_ = _index59
				var _rhs57 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if _this.F22() {
						return ((func() _dafny.Int {
							if (_490_v30).Contains(_this.F22()) {
								return (_490_v30).Multiplicity(_this.F22())
							}
							return _dafny.IntOfInt64(-936)
						})()).Plus((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int))
					}
					return p0
				})())
				_ = _rhs57
				var _rhs58 bool = _this.F22()
				_ = _rhs58
				var _rhs59 _dafny.Array = _455_v0
				_ = _rhs59
				var _lhs49 _dafny.Array = _455_v0
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))
				_ = _lhs50
				var _lhs51 *C2 = _this
				_ = _lhs51
				(_lhs49).ArraySet1(_rhs57, (_lhs50).Int())
				_lhs51.F22_set_(_rhs58)
				_455_v0 = _rhs59
				var _495_v35 _dafny.Map
				_ = _495_v35
				_495_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22(), _489_v29)
				var _496_v36 D0
				_ = _496_v36
				_496_v36 = Companion_D0_.Create_DC1_(_this.F22())
				var _497_v37 _dafny.Map
				_ = _497_v37
				_497_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if (_495_v35).Contains(_this.F22()) {
						return (_495_v35).Get(_this.F22()).(_dafny.Array)
					}
					return _489_v29
				})(), _496_v36)
				(globalState).F20 = ((_497_v37).Merge(_497_v37)).Cardinality()
				r1 = (_456_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_456_v1).Cardinality()))).Uint32()).(bool)
				var _498_v38 _dafny.CodePoint
				_ = _498_v38
				_498_v38 = _dafny.CodePoint('x')
				var _499_v39 _dafny.MultiSet
				_ = _499_v39
				_499_v39 = _dafny.MultiSetOf(_498_v38, _498_v38)
				var _rhs60 _dafny.MultiSet = (_499_v39).Difference(_499_v39)
				_ = _rhs60
				var _rhs61 bool = _this.F22()
				_ = _rhs61
				var _lhs52 *C2 = _this
				_ = _lhs52
				_499_v39 = _rhs60
				_lhs52.F22_set_(_rhs61)
				var _500_v40 _dafny.Map
				_ = _500_v40
				_500_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(globalState), _this.F22())
				var _501_v41 _dafny.Sequence
				_ = _501_v41
				_501_v41 = _dafny.SeqOf(p0)
				var _502_v42 _dafny.Set
				_ = _502_v42
				_502_v42 = _dafny.SetOf(_dafny.IntOfUint32((_501_v41).Cardinality()))
				_500_v40 = (_500_v40).Update((_dafny.SetOf(_dafny.IntOfInt64(28), p0, (_490_v30).Cardinality())).IsDisjointFrom(_502_v42), (_this).Fm3(globalState))
			} else {
				var _503_v43 _dafny.Map
				_ = _503_v43
				_503_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int), _this.F22())
				(_this).F22_set_((_dafny.IntOfUint32((_this.F21()).Cardinality())).Cmp(Companion_Default___.SafeModuloInt((_503_v43).Cardinality(), (_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int))) <= 0)
				r1 = !((func() bool {
					if _this.F22() {
						return !(_this.F22())
					}
					return (func() bool {
						if (func() bool {
							if (_503_v43).Contains(_dafny.IntOfInt64(643)) {
								return (_503_v43).Get(_dafny.IntOfInt64(643)).(bool)
							}
							return _this.F22()
						})() {
							return _this.F22()
						}
						return _this.F22()
					})()
				})())
				var _504_v44 _dafny.CodePoint
				_ = _504_v44
				_504_v44 = _dafny.CodePoint('g')
				var _505_v45 _dafny.Set
				_ = _505_v45
				_505_v45 = _dafny.SetOf(_504_v44, _504_v44)
				var _506_v46 D1
				_ = _506_v46
				_506_v46 = Companion_D1_.Create_DC3_(_dafny.SetOf(_505_v45, _dafny.SetOf(_504_v44)))
				var _507_v47 _dafny.Sequence
				_ = _507_v47
				_507_v47 = _dafny.SeqOf(_506_v46, Companion_Default___.Fm10(globalState))
				(globalState).F20 = _dafny.IntOfUint32((_507_v47).Cardinality())
				var _508_v48 _dafny.Array
				_ = _508_v48
				var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw72
				_508_v48 = _nw72
				var _509_v49 _dafny.Array
				_ = _509_v49
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_9
				var _nw73 _dafny.Array
				_ = _nw73
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw73 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.MultiSet = (func(_510_v30 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_511_i2 _dafny.Int) _dafny.MultiSet {
							return _510_v30
						}
					})(_490_v30)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw73 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw73).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw73).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_509_v49 = _nw73
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_508_v48), 0))
				_ = _index60
				(_508_v48).ArraySet1(_509_v49, (_index60).Int())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_508_v48), 0))
				_ = _index61
				(_508_v48).ArraySet1(_509_v49, (_index61).Int())
				_489_v29 = _489_v29
			}
			var _512_v50 _dafny.Array
			_ = _512_v50
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_10
			var _nw74 _dafny.Array
			_ = _nw74
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw74 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.MultiSet = (func(_513_v30 _dafny.MultiSet, _514_p0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
					return func(_515_i3 _dafny.Int) _dafny.MultiSet {
						return (_513_v30).Update(_this.F22(), Companion_Default___.Abs(_514_p0))
					}
				})(_490_v30, p0)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw74 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw74).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw74).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_512_v50 = _nw74
			var _516_v51 *C1
			_ = _516_v51
			var _nw75 *C1 = New_C1_()
			_ = _nw75
			_nw75.Ctor__(_512_v50, _this.F22())
			_516_v51 = _nw75
			var _517_v52 _dafny.Sequence
			_ = _517_v52
			_517_v52 = _dafny.SeqOf(_516_v51, _516_v51)
			r1 = (!(_this.F22())) == ((_dafny.SetOf(_517_v52, _517_v52)).IsProperSubsetOf(_dafny.SetOf(_517_v52, _517_v52, _517_v52)))
		}
		var _518_v53 _dafny.Set
		_ = _518_v53
		_518_v53 = _dafny.SetOf(_this.F21(), _this.F21())
		var _519_v54 *C0
		_ = _519_v54
		var _nw76 *C0 = New_C0_()
		_ = _nw76
		_nw76.Ctor__(_518_v53)
		_519_v54 = _nw76
		var _520_i4 _dafny.Int
		_ = _520_i4
		_520_i4 = _dafny.Zero
		{
			for (p0).Cmp(p0) >= 0 {
				{
					if (_520_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_520_i4 = (_520_i4).Plus(_dafny.One)
					var _521_v55 _dafny.MultiSet
					_ = _521_v55
					_521_v55 = _dafny.MultiSetOf(_this.F22())
					var _522_v56 D3
					_ = _522_v56
					_522_v56 = Companion_D3_.Create_DC7_(_dafny.MultiSetOf(_this.F22()))
					_521_v55 = (_522_v56).Dtor_cf12()
					var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))
					_ = _index62
					(_455_v0).ArraySet1(p0, (_index62).Int())
					(globalState).F20 = Companion_Default___.Fm6(globalState)
					var _523_v57 _dafny.Set
					_ = _523_v57
					_523_v57 = _dafny.SetOf(p0, (_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int), p0)
					var _524_v58 _dafny.Map
					_ = _524_v58
					_524_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if _this.F22() {
							return _this.F22()
						}
						return _this.F22()
					})(), _this.F22())
					var _525_v59 D0
					_ = _525_v59
					_525_v59 = Companion_D0_.Create_DC1_(true)
					var _526_v60 D0
					_ = _526_v60
					_526_v60 = Companion_D0_.Create_DC2_(_525_v59)
					var _rhs62 _dafny.Set = (Companion_Default___.Fm14(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if (_524_v58).Contains(!(_this.F22())) {
							return (_524_v58).Get(!(_this.F22())).(bool)
						}
						return _this.F22()
					})(), _this.F22()), _dafny.IntOfUint32((_this.F21()).Cardinality()), globalState)).Difference(_523_v57)
					_ = _rhs62
					var _rhs63 bool = (_456_v1).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm15(globalState)).Cardinality(), _dafny.IntOfUint32((_456_v1).Cardinality()))).Uint32()).(bool)
					_ = _rhs63
					var _rhs64 _dafny.Map = (func() _dafny.Map {
						if _this.F22() {
							return _524_v58
						}
						return _524_v58
					})()
					_ = _rhs64
					var _rhs65 D0 = _526_v60
					_ = _rhs65
					_523_v57 = _rhs62
					r1 = _rhs63
					_524_v58 = _rhs64
					_526_v60 = _rhs65
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _527_v62 _dafny.Set
		_ = _527_v62
		_527_v62 = func() _dafny.Set {
			var _coll14 = _dafny.NewBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-25), _dafny.IntOfInt64(428))); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _528_v61 _dafny.Int
				_528_v61 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(-25)).Cmp(_528_v61) <= 0) && ((_528_v61).Cmp(_dafny.IntOfInt64(428)) < 0) {
					_coll14.Add(Companion_Default___.SafeModuloInt(_528_v61, (_455_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_455_v0), 0))).Int()).(_dafny.Int)))
				}
			}
			return _coll14.ToSet()
		}()
		r0 = func(_source9 _dafny.Set) _dafny.Sequence {
			var _529___mcc_h1 _dafny.Set = _source9
			_ = _529___mcc_h1
			var _530_cf11 _dafny.Set = _529___mcc_h1
			_ = _530_cf11
			return _this.F21()
		}(_527_v62)
		var _531_v63 _dafny.MultiSet
		_ = _531_v63
		_531_v63 = _dafny.MultiSetOf(_this.F22(), _this.F22(), true, _this.F22(), _this.F22())
		var _532_v64 _dafny.Map
		_ = _532_v64
		_532_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v1, p0)
		var _533_v65 _dafny.Array
		_ = _533_v65
		var _nw77 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw77
		_533_v65 = _nw77
		var _534_v66 _dafny.Sequence
		_ = _534_v66
		_534_v66 = _dafny.SeqOf(_533_v65, _533_v65)
		var _535_v67 _dafny.Map
		_ = _535_v67
		_535_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_532_v64).Cardinality(), (_534_v66).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_534_v66).Cardinality()))).Uint32()).(_dafny.Array))
		var _536_v68 _dafny.MultiSet
		_ = _536_v68
		_536_v68 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_456_v1))
		r1 = Companion_Default___.Fm12(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hsldowxme")).Cardinality()), (func() _dafny.MultiSet {
			if !(_this.F22()) {
				return _531_v63
			}
			return _531_v63
		})(), (_535_v67).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(88), _533_v65)), _536_v68, globalState)
		return r0, r1
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f21 _dafny.Sequence
	_f22 bool
	F28  bool
	_f27 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f21 = _dafny.EmptySeq
	_this._f22 = false
	_this.F28 = false
	_this._f27 = false
	return &_this
}

type CompanionStruct_C3_ struct {
}

var Companion_C3_ = CompanionStruct_C3_{}

func (_this *C3) Equals(other *C3) bool {
	return _this == other
}

func (_this *C3) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C3)
	return ok && _this.Equals(other)
}

func (*C3) String() string {
	return "_module.C3"
}

func Type_C3_() _dafny.TypeDescriptor {
	return type_C3_{}
}

type type_C3_ struct {
}

func (_this type_C3_) Default() interface{} {
	return (*C3)(nil)
}

func (_this type_C3_) String() string {
	return "main.C3"
}
func (_this *C3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F21() _dafny.Sequence {
	return _this._f21
}
func (_this *C3) F21_set_(value _dafny.Sequence) {
	_this._f21 = value
}
func (_this *C3) F22() bool {
	return _this._f22
}
func (_this *C3) F22_set_(value bool) {
	_this._f22 = value
}
func (_this *C3) Ctor__(f27 bool, f28 bool, f21 _dafny.Sequence, f22 bool) {
	{
		(_this)._f27 = f27
		(_this).F28 = f28
		(_this)._f21 = f21
		(_this)._f22 = f22
	}
}
func (_this *C3) Fm0(globalState *GlobalState) _dafny.Set {
	{
		return ((func() _dafny.Set {
			if false {
				return _dafny.SetOf(_this.F22(), true)
			}
			return _dafny.SetOf((_this).F27(), _this.F28)
		})()).Union(_dafny.SetOf((_this).F27()))
	}
}
func (_this *C3) Fm1(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		return (_this).F27()
	}
}
func (_this *C3) M4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		(_this).F22_set_(false)
		var _hi3 _dafny.Int = p1
		_ = _hi3
		for _537_i0 := _dafny.IntOfUint32((_this.F21()).Cardinality()); _537_i0.Cmp(_hi3) < 0; _537_i0 = _537_i0.Plus(_dafny.One) {
			var _538_v0 _dafny.Array
			_ = _538_v0
			var _nw78 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw78
			_538_v0 = _nw78
			var _539_v1 _dafny.MultiSet
			_ = _539_v1
			_539_v1 = _dafny.MultiSetOf(_538_v0, _538_v0)
			_539_v1 = ((_539_v1).Intersection(_539_v1)).Union((_539_v1).Intersection(_539_v1))
			(globalState).F6 = _this.F21()
			if (_this).F27() {
				var _540_v2 _dafny.Map
				_ = _540_v2
				_540_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22(), p1)
				var _541_v3 _dafny.Set
				_ = _541_v3
				_541_v3 = _dafny.SetOf(_540_v2, _540_v2)
				(_this).F22_set_(((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28, p3), _540_v2, _540_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28, p2))).Union(_541_v3)).IsDisjointFrom(_541_v3))
				var _542_v4 _dafny.Set
				_ = _542_v4
				_542_v4 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("cafmfqmaf"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-524))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}(func(_543_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				})))
				var _544_v5 *C0
				_ = _544_v5
				var _nw79 *C0 = New_C0_()
				_ = _nw79
				_nw79.Ctor__(_542_v4)
				_544_v5 = _nw79
				var _545_v6 D4
				_ = _545_v6
				_545_v6 = Companion_D4_.Create_DC10_(_544_v5)
				var _rhs66 _dafny.Int = p3
				_ = _rhs66
				var _rhs67 *C0 = (_545_v6).Dtor_cf20()
				_ = _rhs67
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				_lhs53.F20 = _rhs66
				_544_v5 = _rhs67
				var _546_v7 _dafny.Sequence
				_ = _546_v7
				_546_v7 = _dafny.SeqOf(p1, Companion_Default___.Fm6(globalState), p2)
				_546_v7 = _546_v7
				var _547_v8 _dafny.Array
				_ = _547_v8
				var _nw80 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw80
				_547_v8 = _nw80
				var _rhs68 _dafny.Array = _547_v8
				_ = _rhs68
				var _rhs69 _dafny.Int = (func() _dafny.Int {
					if true {
						return Companion_Default___.SafeDivisionInt(_537_i0, _537_i0)
					}
					return Companion_Default___.SafeDivisionInt(_537_i0, p3)
				})()
				_ = _rhs69
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				_547_v8 = _rhs68
				_lhs54.F3 = _rhs69
				(_this).F22_set_(false)
			} else {
				var _548_v9 _dafny.MultiSet
				_ = _548_v9
				_548_v9 = _dafny.MultiSetOf(p3, p0)
				var _549_v10 _dafny.Map
				_ = _549_v10
				_549_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22(), (_548_v9).Cardinality())
				var _550_v11 _dafny.Sequence
				_ = _550_v11
				_550_v11 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p0)).Merge(_549_v10))
				_550_v11 = _550_v11
				var _551_v12 _dafny.Set
				_ = _551_v12
				_551_v12 = _dafny.SetOf(p1, _dafny.IntOfInt64(-257), _dafny.IntOfInt64(555), _dafny.IntOfInt64(879), _537_i0)
				var _552_v14 _dafny.Map
				_ = _552_v14
				_552_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mkm")).Cardinality()), func() _dafny.Set {
					var _coll15 = _dafny.NewBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(502))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}(func(_553_i2 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(711)
					}))).Elements()); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _554_v13 _dafny.Int
						_554_v13 = interface{}(_compr_15).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(502))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg18 _dafny.Int) interface{} {
								return coer18(arg18)
							}
						}(func(_553_i2 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(711)
						})), _554_v13) {
							_coll15.Add((_554_v13).Plus(_dafny.IntOfInt64(-762)))
						}
					}
					return _coll15.ToSet()
				}())
				var _555_v15 _dafny.Set
				_ = _555_v15
				_555_v15 = _dafny.SetOf(_551_v12, (func() _dafny.Set {
					if (_552_v14).Contains(p1) {
						return (_552_v14).Get(p1).(_dafny.Set)
					}
					return _551_v12
				})())
				(globalState).F3 = ((_555_v15).Union(_dafny.SetOf(_dafny.SetOf(p0, (_551_v12).Cardinality(), p2)))).Cardinality()
				var _556_v16 D3
				_ = _556_v16
				_556_v16 = Companion_D3_.Create_DC7_(_dafny.MultiSetOf((_this).F27(), _this.F28))
				var _557_v17 _dafny.Map
				_ = _557_v17
				_557_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _556_v16)
				_557_v17 = (_557_v17).Merge(_557_v17)
				(_this).F22_set_(((_this).F27()) || ((_this).F27()))
				var _558_v18 _dafny.Array
				_ = _558_v18
				var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw81
				_558_v18 = _nw81
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_558_v18), 0))
				_ = _index63
				(_558_v18).ArraySet1(_this.F21(), (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_558_v18), 0))
				_ = _index64
				(_558_v18).ArraySet1(_this.F21(), (_index64).Int())
			}
			var _559_v19 _dafny.Array
			_ = _559_v19
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_11
			var _nw82 _dafny.Array
			_ = _nw82
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw82 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = func(_560_i3 _dafny.Int) bool {
					return _dafny.Companion_Sequence_.Equal(_this.F21(), _this.F21())
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw82 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw82).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw82).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_559_v19 = _nw82
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_559_v19), 0))
			_ = _index65
			(_559_v19).ArraySet1((_this).F27(), (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_559_v19), 0))
			_ = _index66
			(_559_v19).ArraySet1((_this.F22()) || ((_this).F27()), (_index66).Int())
		}
		var _561_v20 _dafny.Array
		_ = _561_v20
		var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw83
		_561_v20 = _nw83
		var _562_v21 _dafny.Array
		_ = _562_v21
		var _nw84 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw84
		_562_v21 = _nw84
		var _563_v22 _dafny.Sequence
		_ = _563_v22
		_563_v22 = _dafny.SeqOf(_this.F28, false, true, _this.F22(), _this.F22())
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
		_ = _index67
		(_562_v21).ArraySet1((_563_v22).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_563_v22).Cardinality()))).Uint32()).(bool), (_index67).Int())
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((_562_v21), 0))
		_ = _index68
		(_562_v21).ArraySet1((_this).F27(), (_index68).Int())
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
		_ = _index69
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((_562_v21), 0))
		_ = _index70
		var _rhs70 _dafny.Array = _561_v20
		_ = _rhs70
		var _rhs71 bool = _this.F22()
		_ = _rhs71
		var _rhs72 bool = (_563_v22).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-531), _dafny.IntOfUint32((_563_v22).Cardinality()))).Uint32()).(bool)
		_ = _rhs72
		var _rhs73 bool = (func() bool {
			if _this.F22() {
				return (_this.F28) || (_this.F22())
			}
			return true
		})()
		_ = _rhs73
		var _lhs55 _dafny.Array = _562_v21
		_ = _lhs55
		var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
		_ = _lhs56
		var _lhs57 _dafny.Array = _562_v21
		_ = _lhs57
		var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((_562_v21), 0))
		_ = _lhs58
		var _lhs59 *C3 = _this
		_ = _lhs59
		_561_v20 = _rhs70
		(_lhs55).ArraySet1(_rhs71, (_lhs56).Int())
		(_lhs57).ArraySet1(_rhs72, (_lhs58).Int())
		_lhs59.F28 = _rhs73
		var _564_v23 _dafny.Set
		_ = _564_v23
		_564_v23 = _dafny.SetOf(((_dafny.MultiSetFromSeq(_dafny.SeqOf(p2, p3))).Update(p2, Companion_Default___.Abs(p0))).Cardinality(), p1)
		var _565_v24 _dafny.Set
		_ = _565_v24
		_565_v24 = _564_v23
		var _source10 _dafny.Set = _565_v24
		_ = _source10
		var _566___mcc_h0 _dafny.Set = _source10
		_ = _566___mcc_h0
		var _567_cf11 _dafny.Set = _566___mcc_h0
		_ = _567_cf11
		(globalState).F3 = p3
		var _568_v25 _dafny.Sequence
		_ = _568_v25
		_568_v25 = _dafny.SeqOf(_561_v20, _561_v20, (func() _dafny.Array {
			if _this.F28 {
				return _561_v20
			}
			return _561_v20
		})(), _561_v20, _561_v20)
		_568_v25 = _568_v25
		(_this).F28 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(336))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_569_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), _this.F21())
		var _570_v26 _dafny.CodePoint
		_ = _570_v26
		_570_v26 = _dafny.CodePoint('t')
		var _571_v27 _dafny.Set
		_ = _571_v27
		_571_v27 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_572_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_573_i5 _dafny.Int) _dafny.CodePoint {
				return _572_v26
			}
		})(_570_v26))))
		var _574_v28 _dafny.Map
		_ = _574_v28
		_574_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_570_v26, _571_v27)
		var _575_v29 *C0
		_ = _575_v29
		var _nw85 *C0 = New_C0_()
		_ = _nw85
		_nw85.Ctor__((func() _dafny.Set {
			if (_574_v28).Contains(_570_v26) {
				return (_574_v28).Get(_570_v26).(_dafny.Set)
			}
			return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ce"))
		})())
		_575_v29 = _nw85
		var _576_v31 _dafny.Map
		_ = _576_v31
		_576_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(526))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_577_p2 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_578_i6 _dafny.Int) _dafny.Map {
				return func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(680), _dafny.IntOfInt64(613))); ; {
						_compr_16, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _579_v30 _dafny.Int
						_579_v30 = interface{}(_compr_16).(_dafny.Int)
						if ((_dafny.IntOfInt64(680)).Cmp(_579_v30) <= 0) && ((_579_v30).Cmp(_dafny.IntOfInt64(613)) < 0) {
							_coll16.Add((_579_v30).Plus(_577_p2), (_this).F27())
						}
					}
					return _coll16.ToMap()
				}()
			}
		})(p2))), _dafny.SeqOf(_576_v31, _576_v31, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F27()), _576_v31)) {
			var _580_v32 _dafny.Array
			_ = _580_v32
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw86
			_580_v32 = _nw86
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_580_v32), 0))
			_ = _index71
			(_580_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_563_v22, _dafny.SeqOf(_this.F22(), (_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool), (_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool), (_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool))), (_index71).Int())
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_580_v32), 0))
			_ = _index72
			(_580_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if !(_this.F22()) {
					return _563_v22
				}
				return _563_v22
			})(), _563_v22), (_index72).Int())
			var _581_v33 _dafny.Array
			_ = _581_v33
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_12
			var _nw87 _dafny.Array
			_ = _nw87
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw87 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.MultiSet = func(_582_i7 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(!(_this.F22()), _this.F22())
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw87 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw87).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw87).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_581_v33 = _nw87
			var _583_v34 *C1
			_ = _583_v34
			var _nw88 *C1 = New_C1_()
			_ = _nw88
			_nw88.Ctor__(_581_v33, (_this).F27())
			_583_v34 = _nw88
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_561_v20), 0))
			_ = _index73
			(_561_v20).ArraySet1(p2, (_index73).Int())
			var _584_v35 _dafny.Sequence
			_ = _584_v35
			_584_v35 = _dafny.SeqOf(_dafny.IntOfUint32((_563_v22).Cardinality()), _dafny.IntOfUint32((_this.F21()).Cardinality()))
			var _585_v36 _dafny.Sequence
			_ = _585_v36
			_585_v36 = _dafny.SeqOf(_dafny.IntOfInt64(244), p3, (_dafny.MultiSetFromSeq(_584_v35)).Cardinality(), _dafny.IntOfUint32((_563_v22).Cardinality()), p2)
			var _586_v37 _dafny.MultiSet
			_ = _586_v37
			_586_v37 = _dafny.MultiSetOf(_dafny.IntOfInt64(886))
			var _587_v38 _dafny.Map
			_ = _587_v38
			_587_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_586_v37).Cardinality(), p0)
			var _588_v39 _dafny.MultiSet
			_ = _588_v39
			_588_v39 = _dafny.MultiSetOf((_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool))
			var _589_v40 _dafny.MultiSet
			_ = _589_v40
			_589_v40 = _dafny.MultiSetOf(_588_v39, _588_v39, _588_v39, _588_v39, _588_v39)
			var _590_v41 _dafny.Map
			_ = _590_v41
			_590_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), false)
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_561_v20), 0))
			_ = _index74
			var _rhs74 *C1 = _583_v34
			_ = _rhs74
			var _rhs75 _dafny.CodePoint = (func() _dafny.CodePoint {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_585_v36, Companion_Default___.Fm5(p3, (_587_v38).Cardinality(), p3, globalState)) {
					return (_this.F21()).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_this.F21()).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
				return _dafny.CodePoint('k')
			})()
			_ = _rhs75
			var _rhs76 _dafny.Int = (Companion_Default___.Fm6(globalState)).Minus(_dafny.IntOfUint32(((_580_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_580_v32), 0))).Int()).(_dafny.Sequence)).Cardinality()))
			_ = _rhs76
			var _rhs77 _dafny.Int = (func() _dafny.Int {
				if Companion_Default___.Fm12(p1, _588_v39, _this.F28, _589_v40, globalState) {
					return Companion_Default___.Fm6(globalState)
				}
				return ((_585_v36).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_588_v39).Contains((_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool)) {
						return (_588_v39).Multiplicity((_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool))
					}
					return (_dafny.Zero).Minus((_590_v41).Cardinality())
				})(), _dafny.IntOfUint32((_585_v36).Cardinality()))).Uint32()).(_dafny.Int)).Plus(p3)
			})()
			_ = _rhs77
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			var _lhs61 _dafny.Array = _561_v20
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_561_v20), 0))
			_ = _lhs62
			var _lhs63 *GlobalState = globalState
			_ = _lhs63
			_583_v34 = _rhs74
			_lhs60.F12 = _rhs75
			(_lhs61).ArraySet1(_rhs76, (_lhs62).Int())
			_lhs63.F3 = _rhs77
			var _591_v42 _dafny.CodePoint
			_ = _591_v42
			_591_v42 = _dafny.CodePoint('e')
			var _592_v43 _dafny.Set
			_ = _592_v43
			_592_v43 = _dafny.SetOf(_dafny.CodePoint('e'), _591_v42, _591_v42, _591_v42)
			var _593_v44 _dafny.Set
			_ = _593_v44
			_593_v44 = _dafny.SetOf(_592_v43)
			var _594_v45 D1
			_ = _594_v45
			_594_v45 = Companion_D1_.Create_DC3_(_593_v44)
			var _source11 D1 = _594_v45
			_ = _source11
			if _source11.Is_DC4() {
				var _595___mcc_h1 _dafny.CodePoint = _source11.Get_().(D1_DC4).Cf4
				_ = _595___mcc_h1
				var _596___mcc_h2 _dafny.Map = _source11.Get_().(D1_DC4).Cf5
				_ = _596___mcc_h2
				var _597___mcc_h3 _dafny.Sequence = _source11.Get_().(D1_DC4).Cf6
				_ = _597___mcc_h3
				var _598_cf6 _dafny.Sequence = _597___mcc_h3
				_ = _598_cf6
				var _599_cf5 _dafny.Map = _596___mcc_h2
				_ = _599_cf5
				var _600_cf4 _dafny.CodePoint = _595___mcc_h1
				_ = _600_cf4
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
				_ = _index75
				(_562_v21).ArraySet1((_583_v34).F25(), (_index75).Int())
				_600_cf4 = _591_v42
				var _pat_let_tv31 = _593_v44
				_ = _pat_let_tv31
				_594_v45 = func(_pat_let8_0 D1) D1 {
					return func(_601_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let9_0 _dafny.Set) D1 {
							return func(_602_dt__update_hcf3_h0 _dafny.Set) D1 {
								return Companion_D1_.Create_DC3_(_602_dt__update_hcf3_h0)
							}(_pat_let9_0)
						}(_pat_let_tv31)
					}(_pat_let8_0)
				}(_594_v45)
				var _603_v46 *C1
				_ = _603_v46
				_603_v46 = _583_v34
				var _rhs78 *C1 = (_603_v46)
				_ = _rhs78
				var _rhs79 _dafny.Int = p3
				_ = _rhs79
				var _rhs80 bool = (func() bool {
					if _this.F28 {
						return (_this).F27()
					}
					return _dafny.Companion_Sequence_.IsProperPrefixOf(_this.F21(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-738))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}((func(_604_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_605_i8 _dafny.Int) _dafny.CodePoint {
							return _604_v42
						}
					})(_591_v42))))
				})()
				_ = _rhs80
				var _lhs64 *GlobalState = globalState
				_ = _lhs64
				var _lhs65 *C3 = _this
				_ = _lhs65
				_583_v34 = _rhs78
				_lhs64.F3 = _rhs79
				_lhs65.F22_set_(_rhs80)
			} else if _source11.Is_DC5() {
				var _606___mcc_h4 bool = _source11.Get_().(D1_DC5).Cf7
				_ = _606___mcc_h4
				var _607___mcc_h5 _dafny.Map = _source11.Get_().(D1_DC5).Cf8
				_ = _607___mcc_h5
				var _608___mcc_h6 _dafny.Int = _source11.Get_().(D1_DC5).Cf9
				_ = _608___mcc_h6
				var _609___mcc_h7 _dafny.MultiSet = _source11.Get_().(D1_DC5).Cf10
				_ = _609___mcc_h7
				var _610_cf10 _dafny.MultiSet = _609___mcc_h7
				_ = _610_cf10
				var _611_cf9 _dafny.Int = _608___mcc_h6
				_ = _611_cf9
				var _612_cf8 _dafny.Map = _607___mcc_h5
				_ = _612_cf8
				var _613_cf7 bool = _606___mcc_h4
				_ = _613_cf7
				(_this).F22_set_((_588_v39).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate((_580_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_580_v32), 0))).Int()).(_dafny.Sequence), (_580_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_580_v32), 0))).Int()).(_dafny.Sequence)))))
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
				_ = _index76
				(_562_v21).ArraySet1(_this.F22(), (_index76).Int())
				(_this).F28 = _613_cf7
				var _614_v47 _dafny.Sequence
				_ = _614_v47
				_614_v47 = _dafny.SeqOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_v22, _this.F22())).Cardinality()), _585_v36, _585_v36)
				var _615_v48 *C2
				_ = _615_v48
				var _nw89 *C2 = New_C2_()
				_ = _nw89
				_nw89.Ctor__(_614_v47, _this.F21(), ((_585_v36).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_580_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_580_v32), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_585_v36).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(645)) >= 0)
				_615_v48 = _nw89
			} else {
				var _616___mcc_h8 _dafny.Set = _source11.Get_().(D1_DC3).Cf3
				_ = _616___mcc_h8
				var _617_cf3 _dafny.Set = _616___mcc_h8
				_ = _617_cf3
				var _618_v49 _dafny.Set
				_ = _618_v49
				_618_v49 = _dafny.SetOf(!((_this).F27()), false, !((_583_v34).F25()))
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_561_v20), 0))
				_ = _index77
				(_561_v20).ArraySet1((((_618_v49).Cardinality()).Plus(_dafny.IntOfInt64(-765))).Minus(p1), (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
				_ = _index78
				var _rhs81 _dafny.Int = (_dafny.Zero).Minus((((_561_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_561_v20), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfInt64(876))).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_563_v22).Cardinality()))).Cardinality(), Companion_Default___.Fm6(globalState))))
				_ = _rhs81
				var _rhs82 bool = (_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool)
				_ = _rhs82
				var _rhs83 _dafny.CodePoint = Companion_Default___.Fm4(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p1), (_dafny.Zero).Minus((_561_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_561_v20), 0))).Int()).(_dafny.Int))), _dafny.MultiSetOf((_583_v34).F25(), !(_this.F28), (_this).F27(), (_this).F27(), (_583_v34).F25()), globalState)
				_ = _rhs83
				var _lhs66 *GlobalState = globalState
				_ = _lhs66
				var _lhs67 _dafny.Array = _562_v21
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
				_ = _lhs68
				_lhs66.F20 = _rhs81
				(_lhs67).ArraySet1(_rhs82, (_lhs68).Int())
				_591_v42 = _rhs83
				(_this).F22_set_(!((_583_v34).F25()))
				(globalState).F3 = p0
			}
			var _619_v50 _dafny.Map
			_ = _619_v50
			_619_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_591_v42, _dafny.IntOfUint32((_this.F21()).Cardinality()))
			_619_v50 = (_619_v50).Update(_591_v42, Companion_Default___.SafeModuloInt(p0, p0))
			(_this).F22_set_(_this.F28)
		} else {
			var _620_v51 _dafny.CodePoint
			_ = _620_v51
			_620_v51 = _dafny.CodePoint('r')
			var _621_v52 _dafny.Map
			_ = _621_v52
			_621_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _620_v51)
			var _622_v53 _dafny.Array
			_ = _622_v53
			var _nwElement0_12 _dafny.CodePoint = _620_v51
			_ = _nwElement0_12
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(19))
			_ = _nw90
			(_nw90).ArraySet1CodePoint(_nwElement0_12, 0)
			(_nw90).ArraySet1CodePoint(_620_v51, 1)
			(_nw90).ArraySet1CodePoint(_620_v51, 2)
			(_nw90).ArraySet1CodePoint(_620_v51, 3)
			(_nw90).ArraySet1CodePoint(_620_v51, 4)
			(_nw90).ArraySet1CodePoint(_620_v51, 5)
			(_nw90).ArraySet1CodePoint(_620_v51, 6)
			(_nw90).ArraySet1CodePoint(_620_v51, 7)
			(_nw90).ArraySet1CodePoint(_620_v51, 8)
			(_nw90).ArraySet1CodePoint(_620_v51, 9)
			(_nw90).ArraySet1CodePoint(_620_v51, 10)
			(_nw90).ArraySet1CodePoint(_620_v51, 11)
			(_nw90).ArraySet1CodePoint(_620_v51, 12)
			(_nw90).ArraySet1CodePoint(_620_v51, 13)
			(_nw90).ArraySet1CodePoint(_620_v51, 14)
			(_nw90).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _this.F28 {
					return _dafny.CodePoint('v')
				}
				return (func() _dafny.CodePoint {
					if (_621_v52).Contains(Companion_Default___.Fm6(globalState)) {
						return (_621_v52).Get(Companion_Default___.Fm6(globalState)).(_dafny.CodePoint)
					}
					return _620_v51
				})()
			})(), 15)
			(_nw90).ArraySet1CodePoint(_620_v51, 16)
			(_nw90).ArraySet1CodePoint(_620_v51, 17)
			(_nw90).ArraySet1CodePoint(_dafny.CodePoint('b'), 18)
			_622_v53 = _nw90
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_622_v53), 0))
			_ = _index79
			(_622_v53).ArraySet1CodePoint(_620_v51, (_index79).Int())
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_622_v53), 0))
			_ = _index80
			(_622_v53).ArraySet1CodePoint(_620_v51, (_index80).Int())
			var _623_v54 _dafny.Sequence
			_ = _623_v54
			_623_v54 = _dafny.SeqOf((_dafny.Zero).Minus(p1))
			var _624_v55 _dafny.Sequence
			_ = _624_v55
			_624_v55 = _dafny.SeqOf((_623_v54).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_623_v54).Cardinality()))).Uint32()).(_dafny.Int))
			var _625_v56 _dafny.Map
			_ = _625_v56
			_625_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p2)
			var _626_v57 D1
			_ = _626_v57
			_626_v57 = Companion_D1_.Create_DC5_(_this.F22(), _625_v56, p1, _dafny.MultiSetFromSeq(_623_v54))
			var _627_v58 _dafny.Map
			_ = _627_v58
			_627_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_620_v51, p2)
			var _628_v59 _dafny.Sequence
			_ = _628_v59
			_628_v59 = _dafny.SeqOf(Companion_Default___.Fm5(p1, (func() _dafny.Int {
				if (_627_v58).Contains((_622_v53).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_622_v53), 0))).Int())) {
					return (_627_v58).Get((_622_v53).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_622_v53), 0))).Int())).(_dafny.Int)
				}
				return p0
			})(), p2, globalState), _dafny.SeqOf(p1), _624_v55)
			var _629_v60 _dafny.Array
			_ = _629_v60
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(17))
			_ = _nw91
			_629_v60 = _nw91
			var _630_v61 *C1
			_ = _630_v61
			var _nw92 *C1 = New_C1_()
			_ = _nw92
			_nw92.Ctor__(_629_v60, _this.F28)
			_630_v61 = _nw92
			var _631_v62 _dafny.Map
			_ = _631_v62
			_631_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool), _630_v61)
			var _632_v63 _dafny.MultiSet
			_ = _632_v63
			_632_v63 = _dafny.MultiSetOf((_630_v61).F25())
			var _633_v64 _dafny.MultiSet
			_ = _633_v64
			_633_v64 = _dafny.MultiSetOf(_632_v63)
			var _634_v65 _dafny.Map
			_ = _634_v65
			_634_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_v22, (func() bool {
				if (_576_v31).Contains(p1) {
					return (_576_v31).Get(p1).(bool)
				}
				return Companion_Default___.Fm12((_631_v62).Cardinality(), _632_v63, false, _633_v64, globalState)
			})())
			var _635_v66 _dafny.Array
			_ = _635_v66
			var _nwElement0_13 _dafny.Sequence = _623_v54
			_ = _nwElement0_13
			var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(27))
			_ = _nw93
			(_nw93).ArraySet1(_nwElement0_13, 0)
			(_nw93).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-664))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_636_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_637_i9 _dafny.Int) _dafny.Int {
					return _636_p2
				}
			})(p2))), 1)
			(_nw93).ArraySet1(_624_v55, 2)
			(_nw93).ArraySet1(_dafny.SeqOf((_626_v57).Dtor_cf9(), p2, p0, _dafny.IntOfUint32((_623_v54).Cardinality())), 3)
			(_nw93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_623_v54, (_628_v59).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_628_v59).Cardinality()))).Uint32()).(_dafny.Sequence)), 4)
			(_nw93).ArraySet1(_624_v55, 5)
			(_nw93).ArraySet1(_623_v54, 6)
			(_nw93).ArraySet1(Companion_Default___.Fm5(p3, p0, _dafny.IntOfInt64(-369), globalState), 7)
			(_nw93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_623_v54, _623_v54), 8)
			(_nw93).ArraySet1(_624_v55, 9)
			(_nw93).ArraySet1(Companion_Default___.Fm5(p3, p3, _dafny.IntOfInt64(-379), globalState), 10)
			(_nw93).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(919)), 11)
			(_nw93).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_638_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_639_i10 _dafny.Int) _dafny.Int {
					return _638_p1
				}
			})(p1))), 12)
			(_nw93).ArraySet1(_624_v55, 13)
			(_nw93).ArraySet1(_623_v54, 14)
			(_nw93).ArraySet1(_623_v54, 15)
			(_nw93).ArraySet1(_624_v55, 16)
			(_nw93).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_640_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_641_i11 _dafny.Int) _dafny.Int {
					return _640_p2
				}
			})(p2))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_642_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_643_i11 _dafny.Int) _dafny.Int {
					return _642_p2
				}
			})(p2)))).Cardinality()))).Uint32(), (_634_v65).Cardinality()), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_644_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_645_i11 _dafny.Int) _dafny.Int {
					return _644_p2
				}
			})(p2))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_646_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_647_i11 _dafny.Int) _dafny.Int {
					return _646_p2
				}
			})(p2)))).Cardinality()))).Uint32(), (_634_v65).Cardinality())).Cardinality()))).Uint32(), p1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F21()).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_648_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_649_i11 _dafny.Int) _dafny.Int {
					return _648_p2
				}
			})(p2))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_650_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_651_i11 _dafny.Int) _dafny.Int {
					return _650_p2
				}
			})(p2)))).Cardinality()))).Uint32(), (_634_v65).Cardinality()), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_652_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_653_i11 _dafny.Int) _dafny.Int {
					return _652_p2
				}
			})(p2))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}((func(_654_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_655_i11 _dafny.Int) _dafny.Int {
					return _654_p2
				}
			})(p2)))).Cardinality()))).Uint32(), (_634_v65).Cardinality())).Cardinality()))).Uint32(), p1)).Cardinality()))).Uint32(), p0), 17)
			(_nw93).ArraySet1(_623_v54, 18)
			(_nw93).ArraySet1(_623_v54, 19)
			(_nw93).ArraySet1(_624_v55, 20)
			(_nw93).ArraySet1(_623_v54, 21)
			(_nw93).ArraySet1(_dafny.SeqOf(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p1)).Cardinality(), p3, p2), 22)
			(_nw93).ArraySet1(_dafny.SeqOf((_625_v56).Cardinality(), p1), 23)
			(_nw93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_623_v54, _624_v55), 24)
			(_nw93).ArraySet1(_623_v54, 25)
			(_nw93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_623_v54, _624_v55), 26)
			_635_v66 = _nw93
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_635_v66), 0))
			_ = _index81
			(_635_v66).ArraySet1(_dafny.SeqOf(p0), (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_635_v66), 0))
			_ = _index82
			(_635_v66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_624_v55, _624_v55), (_index82).Int())
			var _656_v67 D3
			_ = _656_v67
			_656_v67 = Companion_D3_.Create_DC9_((p3).Plus(p1), p0, p0)
			var _source12 D3 = _656_v67
			_ = _source12
			if _source12.Is_DC8() {
				var _657___mcc_h9 _dafny.Int = _source12.Get_().(D3_DC8).Cf13
				_ = _657___mcc_h9
				var _658___mcc_h10 _dafny.Map = _source12.Get_().(D3_DC8).Cf14
				_ = _658___mcc_h10
				var _659___mcc_h11 D1 = _source12.Get_().(D3_DC8).Cf15
				_ = _659___mcc_h11
				var _660___mcc_h12 bool = _source12.Get_().(D3_DC8).Cf16
				_ = _660___mcc_h12
				var _661_cf16 bool = _660___mcc_h12
				_ = _661_cf16
				var _662_cf15 D1 = _659___mcc_h11
				_ = _662_cf15
				var _663_cf14 _dafny.Map = _658___mcc_h10
				_ = _663_cf14
				var _664_cf13 _dafny.Int = _657___mcc_h9
				_ = _664_cf13
				(globalState).F3 = (p3).Times(Companion_Default___.Fm6(globalState))
				var _665_v68 _dafny.Map
				_ = _665_v68
				_665_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _664_cf13)
				var _666_v69 _dafny.Set
				_ = _666_v69
				_666_v69 = _dafny.SetOf(_this.F21())
				var _667_v70 *C0
				_ = _667_v70
				var _nw94 *C0 = New_C0_()
				_ = _nw94
				_nw94.Ctor__(_666_v69)
				_667_v70 = _nw94
				var _668_v71 _dafny.Sequence
				_ = _668_v71
				_668_v71 = _dafny.SeqOf(_667_v70, _667_v70)
				var _669_v72 _dafny.Sequence
				_ = _669_v72
				_669_v72 = _dafny.SeqOf(_668_v71, _668_v71, _dafny.SeqOf(_667_v70, _667_v70))
				var _670_v73 _dafny.MultiSet
				_ = _670_v73
				_670_v73 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_669_v72)).Cardinality(), _664_cf13, _664_cf13)
				var _pat_let_tv32 = _625_v56
				_ = _pat_let_tv32
				var _pat_let_tv33 = _625_v56
				_ = _pat_let_tv33
				var _pat_let_tv34 = _625_v56
				_ = _pat_let_tv34
				var _pat_let_tv35 = p1
				_ = _pat_let_tv35
				var _pat_let_tv36 = _670_v73
				_ = _pat_let_tv36
				var _pat_let_tv37 = p2
				_ = _pat_let_tv37
				var _rhs84 _dafny.Int = (func() _dafny.Int {
					if (_665_v68).Contains((_dafny.IntOfUint32((_this.F21()).Cardinality())).Times(p1)) {
						return (_665_v68).Get((_dafny.IntOfUint32((_this.F21()).Cardinality())).Times(p1)).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(p3)
				})()
				_ = _rhs84
				var _rhs85 bool = _this.F22()
				_ = _rhs85
				var _rhs86 D1 = func(_pat_let10_0 D1) D1 {
					return func(_671_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let11_0 _dafny.Map) D1 {
							return func(_672_dt__update_hcf8_h0 _dafny.Map) D1 {
								return func(_pat_let12_0 _dafny.MultiSet) D1 {
									return func(_673_dt__update_hcf10_h0 _dafny.MultiSet) D1 {
										return func(_pat_let13_0 _dafny.Int) D1 {
											return func(_674_dt__update_hcf9_h0 _dafny.Int) D1 {
												return Companion_D1_.Create_DC5_((_671_dt__update__tmp_h1).Dtor_cf7(), _672_dt__update_hcf8_h0, _674_dt__update_hcf9_h0, _673_dt__update_hcf10_h0)
											}(_pat_let13_0)
										}(_pat_let_tv37)
									}(_pat_let12_0)
								}(_pat_let_tv36)
							}(_pat_let11_0)
						}((_pat_let_tv32).Update(_this.F28, (func() _dafny.Int {
							if (_pat_let_tv33).Contains(false) {
								return (_pat_let_tv34).Get(false).(_dafny.Int)
							}
							return _pat_let_tv35
						})()))
					}(_pat_let10_0)
				}(_662_cf15)
				_ = _rhs86
				var _rhs87 bool = (_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool)
				_ = _rhs87
				var _rhs88 bool = _this.F22()
				_ = _rhs88
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				var _lhs70 *C3 = _this
				_ = _lhs70
				_lhs69.F20 = _rhs84
				_661_cf16 = _rhs85
				_626_v57 = _rhs86
				_661_cf16 = _rhs87
				_lhs70.F22_set_(_rhs88)
				var _675_v74 _dafny.Map
				_ = _675_v74
				_675_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_661_cf16), ((_665_v68).Cardinality()).Cmp(p1) != 0)
				var _rhs89 _dafny.Map = _675_v74
				_ = _rhs89
				var _rhs90 *C0 = _667_v70
				_ = _rhs90
				var _rhs91 bool = _661_cf16
				_ = _rhs91
				var _rhs92 _dafny.Map = (_675_v74).Merge(_675_v74)
				_ = _rhs92
				var _lhs71 *C3 = _this
				_ = _lhs71
				_675_v74 = _rhs89
				_667_v70 = _rhs90
				_lhs71.F28 = _rhs91
				_675_v74 = _rhs92
				var _rhs93 bool = ((_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool)) && ((func() bool {
					if _this.F28 {
						return (_630_v61).F25()
					}
					return _this.F28
				})())
				_ = _rhs93
				var _rhs94 bool = (_this).F27()
				_ = _rhs94
				var _lhs72 *C3 = _this
				_ = _lhs72
				var _lhs73 *C3 = _this
				_ = _lhs73
				_lhs72.F28 = _rhs93
				_lhs73.F22_set_(_rhs94)
			} else if _source12.Is_DC9() {
				var _676___mcc_h13 _dafny.Int = _source12.Get_().(D3_DC9).Cf17
				_ = _676___mcc_h13
				var _677___mcc_h14 _dafny.Int = _source12.Get_().(D3_DC9).Cf18
				_ = _677___mcc_h14
				var _678___mcc_h15 _dafny.Int = _source12.Get_().(D3_DC9).Cf19
				_ = _678___mcc_h15
				var _679_cf19 _dafny.Int = _678___mcc_h15
				_ = _679_cf19
				var _680_cf18 _dafny.Int = _677___mcc_h14
				_ = _680_cf18
				var _681_cf17 _dafny.Int = _676___mcc_h13
				_ = _681_cf17
				var _682_v75 _dafny.Sequence
				_ = _682_v75
				_682_v75 = _dafny.SeqOf(_629_v60)
				var _683_v76 *C1
				_ = _683_v76
				var _nw95 *C1 = New_C1_()
				_ = _nw95
				_nw95.Ctor__((_682_v75).Select((Companion_Default___.SafeIndex(_680_cf18, _dafny.IntOfUint32((_682_v75).Cardinality()))).Uint32()).(_dafny.Array), ((_562_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))).Int()).(bool)) || ((_this).F27()))
				_683_v76 = _nw95
				var _684_v77 bool
				_ = _684_v77
				var _out5 bool
				_ = _out5
				_out5 = Companion_Default___.M0(_562_v21, _this.F22(), globalState)
				_684_v77 = _out5
				var _685_v78 bool
				_ = _685_v78
				var _out6 bool
				_ = _out6
				_out6 = Companion_Default___.M0(_562_v21, (_630_v61).F25(), globalState)
				_685_v78 = _out6
				(_this).F22_set_(_684_v77)
			} else {
				var _686___mcc_h16 _dafny.MultiSet = _source12.Get_().(D3_DC7).Cf12
				_ = _686___mcc_h16
				var _687_cf12 _dafny.MultiSet = _686___mcc_h16
				_ = _687_cf12
				(globalState).F20 = p3
				var _688_v79 _dafny.Map
				_ = _688_v79
				_688_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (func() bool {
					if _this.F22() {
						return true
					}
					return (_630_v61).F25()
				})())
				_688_v79 = (_688_v79).Update((_this).F27(), (_630_v61).F25())
				var _689_v80 _dafny.Array
				_ = _689_v80
				var _nwElement0_14 _dafny.Sequence = _this.F21()
				_ = _nwElement0_14
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(11))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_14, 0)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm11(_this.F22(), globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.IntOfUint32((Companion_Default___.Fm11(_this.F22(), globalState)).Cardinality()))).Uint32(), _620_v51), 1)
				(_nw96).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jnsquovwx"), 2)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F21(), _dafny.UnicodeSeqOfUtf8Bytes("ittgfr")), 3)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cqonwvi"), _this.F21()), 4)
				(_nw96).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xwfxaqtqq"), 5)
				(_nw96).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mt"), 6)
				(_nw96).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(242))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}(func(_690_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				})), 7)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-696))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}((func(_691_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_692_i13 _dafny.Int) _dafny.CodePoint {
						return _691_v51
					}
				})(_620_v51))), _this.F21()), 8)
				(_nw96).ArraySet1(_this.F21(), 9)
				(_nw96).ArraySet1(_this.F21(), 10)
				_689_v80 = _nw96
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_689_v80), 0))
				_ = _index83
				(_689_v80).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-122))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}((func(_693_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_694_i14 _dafny.Int) _dafny.CodePoint {
						return _693_v51
					}
				})(_620_v51))), _dafny.UnicodeSeqOfUtf8Bytes("kmvw")), (_index83).Int())
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_689_v80), 0))
				_ = _index84
				(_689_v80).ArraySet1(_this.F21(), (_index84).Int())
				_624_v55 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(301))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}(func(_695_i15 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(222)
				}))
			}
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_562_v21), 0))
			_ = _index85
			(_562_v21).ArraySet1(true, (_index85).Int())
			var _696_v81 _dafny.Sequence
			_ = _696_v81
			_696_v81 = _dafny.SeqOf(_dafny.MultiSetOf(_632_v63))
			_631_v62 = (_631_v62).Update(Companion_Default___.Fm12(p2, _dafny.MultiSetOf((_this).F27(), _this.F22(), (_630_v61).F25(), Companion_Default___.Fm12(p1, _632_v63, (_this).F27(), _dafny.MultiSetOf(_632_v63), globalState), Companion_Default___.Fm12(_dafny.IntOfUint32((_this.F21()).Cardinality()), _dafny.MultiSetOf(_this.F22(), _this.F28), _this.F22(), _633_v64, globalState)), _this.F22(), (_696_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-979), _dafny.IntOfUint32((_696_v81).Cardinality()))).Uint32()).(_dafny.MultiSet), globalState), _630_v61)
		}
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_561_v20), 0))); ; {
			_guard_loop_0, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _697_i16 _dafny.Int
			_697_i16 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_697_i16).Sign() != -1) && ((_697_i16).Cmp(_dafny.ArrayLen((_561_v20), 0)) < 0)) {
				(_561_v20).ArraySet1((_697_i16).Times(p3), (_697_i16).Int())
			}
		}
		var _698_v82 _dafny.Array
		_ = _698_v82
		var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
		_ = _nw97
		_698_v82 = _nw97
		r0 = _698_v82
		return r0
	}
}
func (_this *C3) F27() bool {
	{
		return _this._f27
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f21 _dafny.Sequence
	_f22 bool
	_f23 _dafny.Sequence
	F29  _dafny.Set
	_f30 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f21 = _dafny.EmptySeq
	_this._f22 = false
	_this._f23 = _dafny.EmptySeq
	_this.F29 = _dafny.EmptySet
	_this._f30 = false
	return &_this
}

type CompanionStruct_C4_ struct {
}

var Companion_C4_ = CompanionStruct_C4_{}

func (_this *C4) Equals(other *C4) bool {
	return _this == other
}

func (_this *C4) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C4)
	return ok && _this.Equals(other)
}

func (*C4) String() string {
	return "_module.C4"
}

func Type_C4_() _dafny.TypeDescriptor {
	return type_C4_{}
}

type type_C4_ struct {
}

func (_this type_C4_) Default() interface{} {
	return (*C4)(nil)
}

func (_this type_C4_) String() string {
	return "main.C4"
}
func (_this *C4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F21() _dafny.Sequence {
	return _this._f21
}
func (_this *C4) F21_set_(value _dafny.Sequence) {
	_this._f21 = value
}
func (_this *C4) F22() bool {
	return _this._f22
}
func (_this *C4) F22_set_(value bool) {
	_this._f22 = value
}
func (_this *C4) F23() _dafny.Sequence {
	return _this._f23
}
func (_this *C4) F23_set_(value _dafny.Sequence) {
	_this._f23 = value
}
func (_this *C4) Ctor__(f29 _dafny.Set, f30 bool, f23 _dafny.Sequence, f21 _dafny.Sequence, f22 bool) {
	{
		(_this).F29 = f29
		(_this)._f30 = f30
		(_this)._f23 = f23
		(_this)._f21 = f21
		(_this)._f22 = f22
	}
}
func (_this *C4) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		if (_this).F30() {
			return false
		} else {
			return _this.F22()
		}
	}
}
func (_this *C4) Fm0(globalState *GlobalState) _dafny.Set {
	{
		return (_dafny.SetOf(false)).Intersection((_dafny.SetOf(_this.F22(), (_this).F30())).Intersection(_dafny.SetOf((_this).F30())))
	}
}
func (_this *C4) Fm1(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		return _this.F22()
	}
}
func (_this *C4) F30() bool {
	{
		return _this._f30
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
