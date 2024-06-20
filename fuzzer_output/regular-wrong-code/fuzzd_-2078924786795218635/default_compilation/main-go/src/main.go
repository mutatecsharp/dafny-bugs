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
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.IntOfInt64(910)).Plus(_dafny.IntOfInt64(260)))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.CodePoint, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uirae"), _dafny.UnicodeSeqOfUtf8Bytes("r"))).Cardinality())).Times(_dafny.IntOfInt64(21))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Map, p1 _dafny.CodePoint, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true, true)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-93)))
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_D9_.Create_DC26_(_dafny.MultiSetOf(_dafny.IntOfInt64(-657), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality())))).Dtor_cf38())
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(56)).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())), (func() _dafny.Int {
		if true {
			return ((Companion_D10_.Create_DC28_(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(384))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			}))).Cardinality())))).Dtor_cf39()).Cardinality()
		}
		return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("y"))).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("qaoo"))).Intersection((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _3_v0 _dafny.Sequence
			_3_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})))).Contains(_3_v0) {
				_coll0.Add(_3_v0)
			}
		}
		return _coll0.ToSet()
	}()).Intersection(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-523))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_4_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	})), _dafny.UnicodeSeqOfUtf8Bytes("flq"))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	return !(((_dafny.Zero).Minus((_dafny.IntOfInt64(478)).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality())))).Cmp((_dafny.MultiSetOf(_dafny.IntOfInt64(128), _dafny.IntOfInt64(805))).Cardinality()) != 0)
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_(_dafny.CodePoint('i'), _dafny.IntOfInt64(216), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(989), (_dafny.MultiSetOf(_dafny.CodePoint('s'))).Cardinality())).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(989), (_dafny.MultiSetOf(_dafny.CodePoint('s'))).Cardinality()), _5_v0) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_5_v0, _dafny.IntOfInt64(-206)))
			}
		}
		return _coll1.ToSet()
	}()).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), Companion_D10_.Create_DC29_((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))), (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(832), _dafny.IntOfInt64(101))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _6_v1 _dafny.Int
			_6_v1 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(832)).Cmp(_6_v1) <= 0) && ((_6_v1).Cmp(_dafny.IntOfInt64(101)) < 0) {
				_coll2.Add(Companion_Default___.SafeDivisionInt(_6_v1, _dafny.IntOfInt64(228)), !(true))
			}
		}
		return _coll2.ToMap()
	}()).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sh"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pcvumw"), _dafny.UnicodeSeqOfUtf8Bytes("yl")))
}
func (_static *CompanionStruct_Default___) M1(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) (_dafny.Sequence, *C0, _dafny.Sequence, bool) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 *C0 = (*C0)(nil)
	_ = r1
	var r2 _dafny.Sequence = _dafny.EmptySeq
	_ = r2
	var r3 bool = false
	_ = r3
	if p1 {
		var _7_v0 _dafny.Int
		_ = _7_v0
		_7_v0 = _dafny.IntOfInt64(845)
		_7_v0 = _7_v0
		var _8_v1 *C1
		_ = _8_v1
		var _nw0 *C1 = New_C1_()
		_ = _nw0
		_nw0.Ctor__()
		_8_v1 = _nw0
		var _9_v3 _dafny.Sequence
		_ = _9_v3
		_9_v3 = _dafny.SeqOf(p1)
		var _10_v4 _dafny.Map
		_ = _10_v4
		_10_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v0, _9_v3)
		var _11_v5 _dafny.Map
		_ = _11_v5
		_11_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v0, Companion_Default___.Fm8(_7_v0, func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_10_v4).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _12_v2 _dafny.Int
				_12_v2 = interface{}(_compr_3).(_dafny.Int)
				if (_10_v4).Contains(_12_v2) {
					_coll3.Add((_12_v2).Minus(_7_v0), _7_v0)
				}
			}
			return _coll3.ToMap()
		}(), _7_v0, true, globalState))
		var _13_v6 _dafny.Map
		_ = _13_v6
		_13_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v1, _11_v5)
		var _14_v7 D8
		_ = _14_v7
		_14_v7 = Companion_D8_.Create_DC22_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v0, p1)))
		var _15_v8 D8
		_ = _15_v8
		_15_v8 = Companion_D8_.Create_DC22_((_14_v7).Dtor_cf29())
		_13_v6 = ((_13_v6).Merge((_14_v7).Dtor_cf29())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v1, _11_v5)).Update(_8_v1, _11_v5))
		var _16_v9 _dafny.Array
		_ = _16_v9
		var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw1
		_16_v9 = _nw1
		var _17_v10 _dafny.Sequence
		_ = _17_v10
		_17_v10 = _dafny.SeqOf(_dafny.MultiSetOf(_7_v0))
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_16_v9), 0))
		_ = _index0
		(_16_v9).ArraySet1(_17_v10, (_index0).Int())
		var _18_v11 _dafny.Map
		_ = _18_v11
		_18_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v0, _7_v0)
		var _19_v12 _dafny.Set
		_ = _19_v12
		_19_v12 = _dafny.SetOf(p1, p1)
		var _20_v13 _dafny.Map
		_ = _20_v13
		_20_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm2(globalState))
		var _21_v14 bool
		_ = _21_v14
		_21_v14 = false
		var _22_v15 *C0
		_ = _22_v15
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__(_20_v13, _7_v0, _21_v14)
		_22_v15 = _nw2
		var _23_v16 _dafny.Map
		_ = _23_v16
		_23_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(303))
		var _24_v17 _dafny.Map
		_ = _24_v17
		_24_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _23_v16)
		var _25_v18 _dafny.MultiSet
		_ = _25_v18
		_25_v18 = _dafny.MultiSetOf(_7_v0)
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_16_v9), 0))
		_ = _index1
		var _rhs0 bool = Companion_Default___.Fm8(_7_v0, (_18_v11).Update((_19_v12).Cardinality(), _7_v0), _7_v0, p1, globalState)
		_ = _rhs0
		var _rhs1 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if (p1) && (true) {
				return _7_v0
			}
			return _7_v0
		})())
		_ = _rhs1
		var _rhs2 *C0 = _22_v15
		_ = _rhs2
		var _rhs3 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm3(p0, p1, p1, globalState), (_dafny.Zero).Minus((_7_v0).Times(_dafny.IntOfInt64(547))))
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_17_v10, _dafny.Companion_Sequence_.Concatenate(_17_v10, _dafny.Companion_Sequence_.Update(_17_v10, (Companion_Default___.SafeIndex((_24_v17).Cardinality(), _dafny.IntOfUint32((_17_v10).Cardinality()))).Uint32(), (_25_v18).Update(_7_v0, Companion_Default___.Abs(_7_v0)))))
		_ = _rhs4
		var _lhs0 _dafny.Array = _16_v9
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_16_v9), 0))
		_ = _lhs1
		r3 = _rhs0
		_7_v0 = _rhs1
		r1 = _rhs2
		_7_v0 = _rhs3
		(_lhs0).ArraySet1(_rhs4, (_lhs1).Int())
		var _26_v19 _dafny.Sequence
		_ = _26_v19
		_26_v19 = _dafny.UnicodeSeqOfUtf8Bytes("pbwt")
		var _rhs5 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if p1 {
				return _7_v0
			}
			return _7_v0
		})())
		_ = _rhs5
		var _rhs6 _dafny.Int = _7_v0
		_ = _rhs6
		var _rhs7 _dafny.Int = (func() _dafny.Int {
			if p1 {
				return _7_v0
			}
			return (_22_v15).Fm0(_7_v0, p0, p1, _7_v0, globalState)
		})()
		_ = _rhs7
		var _rhs8 _dafny.Sequence = _26_v19
		_ = _rhs8
		_7_v0 = _rhs5
		_7_v0 = _rhs6
		_7_v0 = _rhs7
		r0 = _rhs8
		_23_v16 = (_23_v16).Update(true, _dafny.IntOfUint32((_26_v19).Cardinality()))
	} else {
		var _27_v20 _dafny.Array
		_ = _27_v20
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw3
		_27_v20 = _nw3
		var _28_v21 _dafny.Sequence
		_ = _28_v21
		_28_v21 = _dafny.SeqOf(_27_v20, _27_v20, _27_v20)
		_28_v21 = _28_v21
		var _29_v22 _dafny.Int
		_ = _29_v22
		_29_v22 = _dafny.IntOfInt64(736)
		_29_v22 = (_29_v22).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-887), _29_v22))
		var _30_v23 _dafny.MultiSet
		_ = _30_v23
		_30_v23 = _dafny.MultiSetOf(_29_v22)
		var _31_v24 _dafny.Map
		_ = _31_v24
		_31_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _30_v23)
		var _32_v25 bool
		_ = _32_v25
		_32_v25 = p1
		var _33_v26 T0
		_ = _33_v26
		var _nw4 *C0 = New_C0_()
		_ = _nw4
		_nw4.Ctor__(_31_v24, _dafny.IntOfUint32((_dafny.SeqOf(!(p1))).Cardinality()), _32_v25)
		_33_v26 = _nw4
		var _34_v27 _dafny.MultiSet
		_ = _34_v27
		_34_v27 = _dafny.MultiSetOf(_33_v26)
		var _35_v28 _dafny.MultiSet
		_ = _35_v28
		_35_v28 = _dafny.MultiSetOf(_34_v27, _34_v27)
		_35_v28 = (_dafny.MultiSetOf(_34_v27)).Difference(_35_v28)
		var _36_v29 _dafny.Map
		_ = _36_v29
		_36_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_33_v26).F6(), (_33_v26).F6())
		if !(Companion_Default___.Fm8((_dafny.Zero).Minus(_29_v22), _36_v29, (func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(454), _dafny.IntOfInt64(-189))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _37_v30 _dafny.Int
				_37_v30 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(454)).Cmp(_37_v30) <= 0) && ((_37_v30).Cmp(_dafny.IntOfInt64(-189)) < 0) {
					_coll4.Add((_37_v30).Plus((_dafny.MultiSetOf(_dafny.SetOf(p1))).Cardinality()), (_dafny.Zero).Minus(_29_v22))
				}
			}
			return _coll4.ToMap()
		}()).Cardinality(), p1, globalState)) || (p1) {
			var _38_v31 *C0
			_ = _38_v31
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _30_v23), _29_v22, p1)
			_38_v31 = _nw5
			_29_v22 = _29_v22
			var _39_v32 _dafny.Set
			_ = _39_v32
			_39_v32 = _dafny.SetOf(p1)
			var _40_v33 _dafny.Map
			_ = _40_v33
			_40_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v32, false)
			_29_v22 = (Companion_Default___.SafeModuloInt((_33_v26).F6(), (_33_v26).F6())).Minus((_40_v33).Cardinality())
			var _41_v34 D6
			_ = _41_v34
			_41_v34 = Companion_D6_.Create_DC15_(_38_v31)
			var _42_v35 _dafny.Sequence
			_ = _42_v35
			_42_v35 = _dafny.SeqOf(_38_v31, _38_v31, _38_v31, _38_v31)
			var _43_v36 _dafny.Sequence
			_ = _43_v36
			_43_v36 = _dafny.SeqOf(p1)
			var _44_v37 _dafny.Map
			_ = _44_v37
			_44_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_43_v36).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-721), _dafny.IntOfUint32((_43_v36).Cardinality()))).Uint32()).(bool), _29_v22)
			var _45_v38 _dafny.Sequence
			_ = _45_v38
			_45_v38 = _dafny.UnicodeSeqOfUtf8Bytes("njpmb")
			var _46_v39 _dafny.Array
			_ = _46_v39
			var _nwElement0_0 *C0 = _38_v31
			_ = _nwElement0_0
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(10))
			_ = _nw6
			(_nw6).ArraySet1(_nwElement0_0, 0)
			(_nw6).ArraySet1(_38_v31, 1)
			(_nw6).ArraySet1(_38_v31, 2)
			(_nw6).ArraySet1(_38_v31, 3)
			(_nw6).ArraySet1((_41_v34).Dtor_cf21(), 4)
			(_nw6).ArraySet1(_38_v31, 5)
			(_nw6).ArraySet1(_38_v31, 6)
			(_nw6).ArraySet1((_42_v35).Select((Companion_Default___.SafeIndex(((_44_v37).Update(p1, _dafny.IntOfUint32((_45_v38).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_42_v35).Cardinality()))).Uint32()).(*C0), 7)
			(_nw6).ArraySet1(_38_v31, 8)
			(_nw6).ArraySet1(_38_v31, 9)
			_46_v39 = _nw6
			_46_v39 = _46_v39
			var _47_v40 _dafny.Map
			_ = _47_v40
			_47_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			_47_v40 = (_47_v40).Update(p1, !(_dafny.Companion_Sequence_.IsProperPrefixOf(_45_v38, _45_v38)))
		} else {
			var _48_v41 _dafny.Array
			_ = _48_v41
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw7
			_48_v41 = _nw7
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_48_v41), 0))
			_ = _index2
			(_48_v41).ArraySet1(p1, (_index2).Int())
			var _49_v42 _dafny.Map
			_ = _49_v42
			_49_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_29_v22, p1)
			var _50_v43 _dafny.Sequence
			_ = _50_v43
			_50_v43 = _dafny.SeqOf((func() bool {
				if (_49_v42).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus(_29_v22))) {
					return (_49_v42).Get((_dafny.Zero).Minus((_dafny.Zero).Minus(_29_v22))).(bool)
				}
				return p1
			})())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_48_v41), 0))
			_ = _index3
			(_48_v41).ArraySet1(((_33_v26).F6()).Cmp(_dafny.IntOfUint32((_50_v43).Cardinality())) < 0, (_index3).Int())
			_29_v22 = (_dafny.Zero).Minus(_29_v22)
			var _51_v44 *C1
			_ = _51_v44
			var _nw8 *C1 = New_C1_()
			_ = _nw8
			_nw8.Ctor__()
			_51_v44 = _nw8
			var _52_v45 _dafny.Map
			_ = _52_v45
			_52_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_29_v22, _51_v44)
			var _53_v46 D1
			_ = _53_v46
			_53_v46 = Companion_D1_.Create_DC2_(_dafny.CodePoint('u'), _29_v22, _dafny.IntOfInt64(-318))
			var _54_v47 _dafny.Set
			_ = _54_v47
			_54_v47 = _dafny.SetOf(_dafny.IntOfUint32((_50_v43).Cardinality()))
			var _55_v48 _dafny.Set
			_ = _55_v48
			_55_v48 = _dafny.SetOf((Companion_D1_.Create_DC2_((_53_v46).Dtor_cf2(), (_33_v26).F6(), (_dafny.Zero).Minus((_54_v47).Cardinality()))).Dtor_cf4())
			_52_v45 = (_52_v45).Update((_33_v26).F6(), (func() *C1 {
				if !(Companion_Default___.Fm8((_33_v26).F6(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_33_v26).F6(), (_33_v26).F6()), (_54_v47).Cardinality(), p1, globalState)) {
					return _51_v44
				}
				return _51_v44
			})())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_48_v41), 0))
			_ = _index4
			(_48_v41).ArraySet1(p1, (_index4).Int())
			var _56_v49 *C1
			_ = _56_v49
			var _nw9 *C1 = New_C1_()
			_ = _nw9
			_nw9.Ctor__()
			_56_v49 = _nw9
		}
		var _57_v50 *C1
		_ = _57_v50
		var _nw10 *C1 = New_C1_()
		_ = _nw10
		_nw10.Ctor__()
		_57_v50 = _nw10
	}
	var _58_v51 _dafny.Array
	_ = _58_v51
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(19)
	_ = _len0_0
	var _nw11 _dafny.Array
	_ = _nw11
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw11 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = (func(_59_p1 bool) func(_dafny.Int) _dafny.Int {
			return func(_60_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_60_i0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ovbwu"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_59_p1)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ovbwu")).Cardinality()))).Uint32(), _dafny.CodePoint('n'))).Cardinality()))
			}
		})(p1)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw11 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw11).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw11).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_58_v51 = _nw11
	var _61_v52 _dafny.Int
	_ = _61_v52
	_61_v52 = _dafny.IntOfInt64(926)
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))
	_ = _index5
	(_58_v51).ArraySet1((_dafny.Zero).Minus(_61_v52), (_index5).Int())
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))
	_ = _index6
	(_58_v51).ArraySet1(_61_v52, (_index6).Int())
	var _62_v53 _dafny.Map
	_ = _62_v53
	_62_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.MultiSetOf(_61_v52))
	var _63_v54 bool
	_ = _63_v54
	_63_v54 = p1
	var _64_v55 T0
	_ = _64_v55
	var _nw12 *C0 = New_C0_()
	_ = _nw12
	_nw12.Ctor__(_62_v53, (_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int), _63_v54)
	_64_v55 = _nw12
	var _65_v56 D7
	_ = _65_v56
	_65_v56 = Companion_D7_.Create_DC21_(p1, _64_v55)
	if (_65_v56).Dtor_cf27() {
		_61_v52 = (_dafny.Zero).Minus(((_64_v55).F6()).Minus((_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)))
		var _66_v57 _dafny.Sequence
		_ = _66_v57
		_66_v57 = _dafny.UnicodeSeqOfUtf8Bytes("fqqxnh")
		var _67_v58 _dafny.MultiSet
		_ = _67_v58
		_67_v58 = _dafny.MultiSetOf((_64_v55).F6())
		var _68_v59 _dafny.MultiSet
		_ = _68_v59
		_68_v59 = _dafny.MultiSetOf(_58_v51, _58_v51)
		var _69_v60 D3
		_ = _69_v60
		_69_v60 = Companion_D3_.Create_DC6_(_68_v59)
		var _70_v61 *C1
		_ = _70_v61
		var _nw13 *C1 = New_C1_()
		_ = _nw13
		_nw13.Ctor__()
		_70_v61 = _nw13
		var _71_v62 _dafny.Map
		_ = _71_v62
		_71_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v60, _70_v61)
		var _72_v63 _dafny.Sequence
		_ = _72_v63
		_72_v63 = _dafny.SeqOf((_64_v55).F6(), ((_64_v55).F6()).Times((_64_v55).F6()), Companion_Default___.SafeDivisionInt((_64_v55).F6(), _dafny.IntOfUint32((_66_v57).Cardinality())), ((_67_v58).Difference(_67_v58)).Cardinality(), ((_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)).Times((_71_v62).Cardinality()))
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))
		_ = _index7
		(_58_v51).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_72_v63).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_61_v52)), _dafny.IntOfUint32((_72_v63).Cardinality()))).Uint32()).(_dafny.Int))), (_index7).Int())
		r3 = p1
		var _73_v64 *C0
		_ = _73_v64
		var _nw14 *C0 = New_C0_()
		_ = _nw14
		_nw14.Ctor__(Companion_Default___.Fm5(globalState), _61_v52, _63_v54)
		_73_v64 = _nw14
		var _74_v65 _dafny.Map
		_ = _74_v65
		_74_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v64, !(false))
		_74_v65 = (_74_v65).Update(_73_v64, p1)
		var _75_v66 _dafny.Set
		_ = _75_v66
		_75_v66 = _dafny.SetOf(_72_v63)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))
		_ = _index8
		var _rhs9 _dafny.Int = _dafny.IntOfInt64(-252)
		_ = _rhs9
		var _rhs10 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_66_v57, _dafny.Companion_Sequence_.Concatenate(_66_v57, _dafny.UnicodeSeqOfUtf8Bytes("jnqvr")))
		_ = _rhs10
		var _rhs11 _dafny.Int = _61_v52
		_ = _rhs11
		var _rhs12 _dafny.Set = _dafny.SetOf(_72_v63, _72_v63)
		_ = _rhs12
		var _lhs2 _dafny.Array = _58_v51
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))
		_ = _lhs3
		(_lhs2).ArraySet1(_rhs9, (_lhs3).Int())
		_66_v57 = _rhs10
		_61_v52 = _rhs11
		_75_v66 = _rhs12
	} else {
		var _76_v67 D5
		_ = _76_v67
		_76_v67 = Companion_D5_.Create_DC12_((_61_v52).Cmp((_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)) >= 0)
		_76_v67 = _76_v67
		var _77_v68 _dafny.Sequence
		_ = _77_v68
		_77_v68 = _dafny.UnicodeSeqOfUtf8Bytes("ibjngh")
		r3 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_77_v68, Companion_Default___.Fm10(globalState)), _77_v68)
		var _78_v69 _dafny.Sequence
		_ = _78_v69
		_78_v69 = _dafny.SeqOf(_dafny.IntOfUint32((_77_v68).Cardinality()), _dafny.IntOfUint32((_77_v68).Cardinality()), _dafny.IntOfInt64(751), _dafny.IntOfInt64(427))
		var _79_v70 _dafny.Map
		_ = _79_v70
		_79_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _64_v55)
		var _80_v71 _dafny.Map
		_ = _80_v71
		_80_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v69, _79_v70)
		var _81_v72 _dafny.Map
		_ = _81_v72
		_81_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v52, (_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int))
		r3 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_79_v70), ((func() _dafny.Map {
			if (_80_v71).Contains(_78_v69) {
				return (_80_v71).Get(_78_v69).(_dafny.Map)
			}
			return _79_v70
		})()).Update(Companion_Default___.Fm8((_64_v55).F6(), _81_v72, _dafny.IntOfInt64(902), p1, globalState), _64_v55))
		r3 = (((_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)).Minus((_64_v55).F6())).Cmp((_64_v55).F6()) == 0
		var _82_v73 D1
		_ = _82_v73
		_82_v73 = Companion_D1_.Create_DC2_(p0, _61_v52, (_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int))
		r3 = (func() bool {
			if !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D1_.Create_DC2_(p0, _61_v52, (_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)), _82_v73), _82_v73) {
				return p1
			}
			return (_dafny.IntOfInt64(557)).Cmp((_64_v55).F6()) == 0
		})()
	}
	var _83_v74 *C0
	_ = _83_v74
	var _nw15 *C0 = New_C0_()
	_ = _nw15
	_nw15.Ctor__((_62_v53).Merge(_62_v53), _61_v52, _63_v54)
	_83_v74 = _nw15
	if p1 {
		var _84_v75 _dafny.Sequence
		_ = _84_v75
		_84_v75 = _dafny.UnicodeSeqOfUtf8Bytes("b")
		r0 = _84_v75
		var _85_v76 *C1
		_ = _85_v76
		var _nw16 *C1 = New_C1_()
		_ = _nw16
		_nw16.Ctor__()
		_85_v76 = _nw16
		var _86_v77 _dafny.Int
		_ = _86_v77
		var _out0 _dafny.Int
		_ = _out0
		_out0 = (_85_v76).M0(p1, p0, p1, globalState)
		_86_v77 = _out0
		var _87_v78 _dafny.Sequence
		_ = _87_v78
		_87_v78 = _dafny.SeqOf(_84_v75)
		var _rhs13 _dafny.Sequence = _87_v78
		_ = _rhs13
		var _rhs14 *C0 = _83_v74
		_ = _rhs14
		var _rhs15 _dafny.Int = ((_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)).Plus(((_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)).Times(_86_v77))
		_ = _rhs15
		var _rhs16 bool = ((_dafny.IntOfInt64(405)).Plus(_dafny.IntOfInt64(375))).Cmp((_64_v55).F6()) <= 0
		_ = _rhs16
		_87_v78 = _rhs13
		_83_v74 = _rhs14
		_86_v77 = _rhs15
		r3 = _rhs16
		var _88_v79 _dafny.Map
		_ = _88_v79
		_88_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v55, _64_v55)
		_64_v55 = (func() T0 {
			if p1 {
				return (func() T0 {
					if (_88_v79).Contains(_64_v55) {
						return (_88_v79).Get(_64_v55).(T0)
					}
					return _64_v55
				})()
			}
			return _64_v55
		})()
	} else {
		var _89_v80 _dafny.Array
		_ = _89_v80
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_1
		var _nw17 _dafny.Array
		_ = _nw17
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw17 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_90_p1 bool) func(_dafny.Int) bool {
				return func(_91_i1 _dafny.Int) bool {
					return _90_p1
				}
			})(p1)
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
		_89_v80 = _nw17
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
		_ = _index9
		(_89_v80).ArraySet1(p1, (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
		_ = _index10
		(_89_v80).ArraySet1(!(p1), (_index10).Int())
		var _92_v81 _dafny.Set
		_ = _92_v81
		_92_v81 = _dafny.SetOf(_89_v80)
		_92_v81 = (_92_v81).Intersection(_dafny.SetOf(_89_v80, _89_v80, _89_v80))
		var _93_v82 _dafny.CodePoint
		_ = _93_v82
		_93_v82 = _dafny.CodePoint('a')
		_93_v82 = p0
		_61_v52 = (_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int)
		if (_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool) {
			var _94_v83 _dafny.Set
			_ = _94_v83
			_94_v83 = _dafny.SetOf(_93_v82)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
			_ = _index11
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
			_ = _index12
			var _rhs17 bool = (_94_v83).IsSubsetOf(_94_v83)
			_ = _rhs17
			var _rhs18 bool = (func() bool {
				if (_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool) {
					return false
				}
				return (_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool)
			})()
			_ = _rhs18
			var _lhs4 _dafny.Array = _89_v80
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
			_ = _lhs5
			var _lhs6 _dafny.Array = _89_v80
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
			_ = _lhs7
			(_lhs4).ArraySet1(_rhs17, (_lhs5).Int())
			(_lhs6).ArraySet1(_rhs18, (_lhs7).Int())
			var _95_v84 _dafny.Sequence
			_ = _95_v84
			_95_v84 = _dafny.SeqOf(p1, true, p1, p1)
			var _96_v85 _dafny.MultiSet
			_ = _96_v85
			_96_v85 = _dafny.MultiSetOf((_64_v55).F6())
			var _97_v86 _dafny.Sequence
			_ = _97_v86
			_97_v86 = _dafny.SeqOf((_96_v85).Cardinality())
			var _98_v87 _dafny.Set
			_ = _98_v87
			_98_v87 = _dafny.SetOf(!(p1), _dafny.Companion_Sequence_.Contains(_97_v86, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool), (_95_v84).Select((Companion_Default___.SafeIndex((_58_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_58_v51), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_95_v84).Cardinality()))).Uint32()).(bool))).Cardinality()), p1)
			(globalState).F5 = _98_v87
			var _99_v88 *C1
			_ = _99_v88
			var _nw18 *C1 = New_C1_()
			_ = _nw18
			_nw18.Ctor__()
			_99_v88 = _nw18
			var _100_v89 _dafny.Array
			_ = _100_v89
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_2
			var _nw19 _dafny.Array
			_ = _nw19
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw19 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_101_v86 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_102_i2 _dafny.Int) _dafny.Sequence {
						return _101_v86
					}
				})(_97_v86)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw19 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw19).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw19).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_100_v89 = _nw19
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_100_v89), 0))
			_ = _index13
			(_100_v89).ArraySet1(_97_v86, (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_100_v89), 0))
			_ = _index14
			(_100_v89).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(937))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_103_v52 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_104_i3 _dafny.Int) _dafny.Int {
					return (_dafny.SetOf(_103_v52)).Cardinality()
				}
			})(_61_v52))), (_index14).Int())
			var _105_v90 D4
			_ = _105_v90
			_105_v90 = Companion_D4_.Create_DC9_(_89_v80)
			var _106_v91 _dafny.Map
			_ = _106_v91
			_106_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v90, _93_v82)
			_93_v82 = (func() _dafny.CodePoint {
				if (_106_v91).Contains(Companion_D4_.Create_DC9_(_89_v80)) {
					return (_106_v91).Get(Companion_D4_.Create_DC9_(_89_v80)).(_dafny.CodePoint)
				}
				return p0
			})()
		} else {
			var _107_v92 *C0
			_ = _107_v92
			var _nw20 *C0 = New_C0_()
			_ = _nw20
			_nw20.Ctor__((_83_v74.F8).Merge(_83_v74.F8), _61_v52, p1)
			_107_v92 = _nw20
			var _108_v93 *C1
			_ = _108_v93
			var _nw21 *C1 = New_C1_()
			_ = _nw21
			_nw21.Ctor__()
			_108_v93 = _nw21
			var _109_v94 _dafny.Map
			_ = _109_v94
			_109_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool), _108_v93)
			var _110_v95 _dafny.MultiSet
			_ = _110_v95
			_110_v95 = _dafny.MultiSetOf(p1, (_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool), p1)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
			_ = _index15
			(_89_v80).ArraySet1(!((_83_v74).Fm1((_64_v55).F6(), false, (_109_v94).Cardinality(), _110_v95, globalState)) || (p1), (_index15).Int())
			var _111_v96 _dafny.Sequence
			_ = _111_v96
			_111_v96 = _dafny.SeqOf((_dafny.Zero).Minus((_64_v55).F6()))
			var _112_v97 _dafny.Sequence
			_ = _112_v97
			_112_v97 = _dafny.SeqOf((_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool), _dafny.Companion_Sequence_.Equal(_111_v96, _111_v96), p1)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
			_ = _index16
			var _rhs19 *C1 = _108_v93
			_ = _rhs19
			var _rhs20 bool = !((_112_v97).Select((Companion_Default___.SafeIndex((_64_v55).F6(), _dafny.IntOfUint32((_112_v97).Cardinality()))).Uint32()).(bool))
			_ = _rhs20
			var _rhs21 bool = (_89_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))).Int()).(bool)
			_ = _rhs21
			var _lhs8 _dafny.Array = _89_v80
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_89_v80), 0))
			_ = _lhs9
			_108_v93 = _rhs19
			(_lhs8).ArraySet1(_rhs20, (_lhs9).Int())
			r3 = _rhs21
			var _113_v98 _dafny.Int
			_ = _113_v98
			var _out1 _dafny.Int
			_ = _out1
			_out1 = (_108_v93).M0(!(true), _93_v82, (_64_v55).F7(), globalState)
			_113_v98 = _out1
			var _114_v99 _dafny.Sequence
			_ = _114_v99
			_114_v99 = _dafny.UnicodeSeqOfUtf8Bytes("aqwo")
			r0 = _114_v99
		}
	}
	var _115_v100 _dafny.Array
	_ = _115_v100
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_3
	var _nw22 _dafny.Array
	_ = _nw22
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw22 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Set = (func(_116_p1 bool) func(_dafny.Int) _dafny.Set {
			return func(_117_i4 _dafny.Int) _dafny.Set {
				return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_p1, _116_p1), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_p1, !(_116_p1)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_p1, _116_p1))).Intersection(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_116_p1), _116_p1)))
			}
		})(p1)
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
	_115_v100 = _nw22
	var _118_v101 _dafny.Set
	_ = _118_v101
	_118_v101 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false))
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_115_v100), 0))
	_ = _index17
	(_115_v100).ArraySet1((func() _dafny.Set {
		if p1 {
			return _118_v101
		}
		return _118_v101
	})(), (_index17).Int())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_115_v100), 0))
	_ = _index18
	(_115_v100).ArraySet1(_118_v101, (_index18).Int())
	var _119_v102 _dafny.Sequence
	_ = _119_v102
	_119_v102 = _dafny.UnicodeSeqOfUtf8Bytes("pkeaema")
	r0 = _119_v102
	r1 = _83_v74
	r2 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-574))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}((func(_120_p1 bool, _121_v52 _dafny.Int) func(_dafny.Int) _dafny.Map {
		return func(_122_i5 _dafny.Int) _dafny.Map {
			return func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(146), _dafny.IntOfInt64(220))); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _123_v103 _dafny.Int
					_123_v103 = interface{}(_compr_5).(_dafny.Int)
					if ((_dafny.IntOfInt64(146)).Cmp(_123_v103) <= 0) && ((_123_v103).Cmp(_dafny.IntOfInt64(220)) < 0) {
						_coll5.Add((_123_v103).Minus((_dafny.SetOf(_120_p1)).Cardinality()), _121_v52)
					}
				}
				return _coll5.ToMap()
			}()
		}
	})(p1, _61_v52)))
	r3 = p1
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _124_v0 _dafny.Int
	_ = _124_v0
	_124_v0 = _dafny.IntOfInt64(197)
	var _125_v1 _dafny.MultiSet
	_ = _125_v1
	_125_v1 = _dafny.MultiSetOf(_124_v0)
	var _126_v2 _dafny.Array
	_ = _126_v2
	var _nwElement0_1 _dafny.MultiSet = _dafny.MultiSetOf(_124_v0)
	_ = _nwElement0_1
	var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(13))
	_ = _nw23
	(_nw23).ArraySet1(_nwElement0_1, 0)
	(_nw23).ArraySet1(_125_v1, 1)
	(_nw23).ArraySet1(_125_v1, 2)
	(_nw23).ArraySet1(_125_v1, 3)
	(_nw23).ArraySet1(_125_v1, 4)
	(_nw23).ArraySet1(_dafny.MultiSetOf(_124_v0, _124_v0, _dafny.IntOfInt64(265)), 5)
	(_nw23).ArraySet1(_125_v1, 6)
	(_nw23).ArraySet1(_125_v1, 7)
	(_nw23).ArraySet1(_125_v1, 8)
	(_nw23).ArraySet1(_125_v1, 9)
	(_nw23).ArraySet1(_125_v1, 10)
	(_nw23).ArraySet1(_125_v1, 11)
	(_nw23).ArraySet1(_125_v1, 12)
	_126_v2 = _nw23
	var _127_v3 _dafny.Array
	_ = _127_v3
	var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
	_ = _nw24
	_127_v3 = _nw24
	var _128_v4 _dafny.Map
	_ = _128_v4
	_128_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v0, _127_v3)
	var _129_v5 bool
	_ = _129_v5
	_129_v5 = false
	var _130_v6 _dafny.Set
	_ = _130_v6
	_130_v6 = _dafny.SetOf(_129_v5, _129_v5, _129_v5, _129_v5)
	var _131_globalState *GlobalState
	_ = _131_globalState
	var _nw25 *GlobalState = New_GlobalState_()
	_ = _nw25
	_nw25.Ctor__(_126_v2, _126_v2, true, _128_v4, _dafny.CodePoint('x'), (_130_v6).Union(_130_v6))
	_131_globalState = _nw25
	var _hi0 _dafny.Int = _124_v0
	_ = _hi0
	for _132_i0 := (_124_v0).Plus(_dafny.IntOfInt64(302)); _132_i0.Cmp(_hi0) < 0; _132_i0 = _132_i0.Plus(_dafny.One) {
		var _133_v7 _dafny.Sequence
		_ = _133_v7
		_133_v7 = _dafny.UnicodeSeqOfUtf8Bytes("a")
		var _134_v8 _dafny.Sequence
		_ = _134_v8
		_134_v8 = _dafny.SeqOf(_129_v5, _129_v5)
		var _135_v9 _dafny.CodePoint
		_ = _135_v9
		_135_v9 = _dafny.CodePoint('b')
		var _136_v10 _dafny.Set
		_ = _136_v10
		_136_v10 = _dafny.SetOf(_135_v9)
		var _137_v11 bool
		_ = _137_v11
		_137_v11 = _129_v5
		var _138_v12 _dafny.Map
		_ = _138_v12
		_138_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_v9, false)
		var _139_v14 _dafny.Array
		_ = _139_v14
		var _nwElement0_2 bool = _129_v5
		_ = _nwElement0_2
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(21))
		_ = _nw26
		(_nw26).ArraySet1(_nwElement0_2, 0)
		(_nw26).ArraySet1(_129_v5, 1)
		(_nw26).ArraySet1(_129_v5, 2)
		(_nw26).ArraySet1(!_dafny.Companion_Sequence_.Equal(_133_v7, _dafny.Companion_Sequence_.Update(_133_v7, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_134_v8).Cardinality()), _dafny.IntOfUint32((_133_v7).Cardinality()))).Uint32(), _135_v9)), 3)
		(_nw26).ArraySet1(_129_v5, 4)
		(_nw26).ArraySet1(_129_v5, 5)
		(_nw26).ArraySet1(_129_v5, 6)
		(_nw26).ArraySet1((_124_v0).Cmp(_132_i0) == 0, 7)
		(_nw26).ArraySet1(_129_v5, 8)
		(_nw26).ArraySet1(((_136_v10).Cardinality()).Cmp(_132_i0) >= 0, 9)
		(_nw26).ArraySet1(_129_v5, 10)
		(_nw26).ArraySet1(_129_v5, 11)
		(_nw26).ArraySet1(_129_v5, 12)
		(_nw26).ArraySet1((_137_v11), 13)
		(_nw26).ArraySet1(_129_v5, 14)
		(_nw26).ArraySet1(_129_v5, 15)
		(_nw26).ArraySet1(!(_dafny.MultiSetOf(_129_v5)).Contains(_129_v5), 16)
		(_nw26).ArraySet1(_129_v5, 17)
		(_nw26).ArraySet1((_136_v10).IsProperSubsetOf(func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_138_v12).Keys().Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _140_v13 _dafny.CodePoint
				_140_v13 = interface{}(_compr_6).(_dafny.CodePoint)
				if (_138_v12).Contains(_140_v13) {
					_coll6.Add(_140_v13)
				}
			}
			return _coll6.ToSet()
		}()), 18)
		(_nw26).ArraySet1(_129_v5, 19)
		(_nw26).ArraySet1(_129_v5, 20)
		_139_v14 = _nw26
		var _141_v15 _dafny.Set
		_ = _141_v15
		_141_v15 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("b"))
		var _142_v16 _dafny.Sequence
		_ = _142_v16
		_142_v16 = _dafny.SeqOf(_dafny.IntOfInt64(-894), (_dafny.Zero).Minus(_124_v0))
		var _143_v17 _dafny.Set
		_ = _143_v17
		_143_v17 = _dafny.SetOf((func() _dafny.Int {
			if _137_v11 {
				return _132_i0
			}
			return _dafny.IntOfUint32((_134_v8).Cardinality())
		})(), Companion_Default___.SafeDivisionInt(_124_v0, ((_125_v1).Update(_dafny.IntOfInt64(320), Companion_Default___.Abs((_142_v16).Select((Companion_Default___.SafeIndex(_124_v0, _dafny.IntOfUint32((_142_v16).Cardinality()))).Uint32()).(_dafny.Int)))).Cardinality()))
		var _144_v18 _dafny.Map
		_ = _144_v18
		_144_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v16, _139_v14)
		var _145_v19 _dafny.Sequence
		_ = _145_v19
		_145_v19 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_130_v6).Cardinality(), _124_v0))
		var _146_v20 _dafny.Sequence
		_ = _146_v20
		_146_v20 = _dafny.SeqOf(_143_v17)
		var _rhs22 bool = _129_v5
		_ = _rhs22
		var _rhs23 _dafny.Array = (func() _dafny.Array {
			if (_144_v18).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_142_v16, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-122), _dafny.IntOfUint32((_142_v16).Cardinality()))).Uint32(), ((_145_v19).Select((Companion_Default___.SafeIndex(_124_v0, _dafny.IntOfUint32((_145_v19).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()), _142_v16)) {
				return (_144_v18).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_142_v16, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-122), _dafny.IntOfUint32((_142_v16).Cardinality()))).Uint32(), ((_145_v19).Select((Companion_Default___.SafeIndex(_124_v0, _dafny.IntOfUint32((_145_v19).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()), _142_v16)).(_dafny.Array)
			}
			return _139_v14
		})()
		_ = _rhs23
		var _rhs24 _dafny.Set = _141_v15
		_ = _rhs24
		var _rhs25 _dafny.Set = (_143_v17).Intersection((_146_v20).Select((Companion_Default___.SafeIndex(_132_i0, _dafny.IntOfUint32((_146_v20).Cardinality()))).Uint32()).(_dafny.Set))
		_ = _rhs25
		_129_v5 = _rhs22
		_139_v14 = _rhs23
		_141_v15 = _rhs24
		_143_v17 = _rhs25
		var _147_v21 *C1
		_ = _147_v21
		var _nw27 *C1 = New_C1_()
		_ = _nw27
		_nw27.Ctor__()
		_147_v21 = _nw27
		_129_v5 = _129_v5
		_124_v0 = _124_v0
	}
	var _hi1 _dafny.Int = _124_v0
	_ = _hi1
	for _148_i1 := _dafny.IntOfInt64(71); _148_i1.Cmp(_hi1) < 0; _148_i1 = _148_i1.Plus(_dafny.One) {
		var _149_v22 _dafny.CodePoint
		_ = _149_v22
		_149_v22 = _dafny.CodePoint('r')
		var _150_v23 _dafny.Sequence
		_ = _150_v23
		_150_v23 = _dafny.SeqOf(_148_i1)
		var _151_v24 D1
		_ = _151_v24
		_151_v24 = Companion_D1_.Create_DC2_(_149_v22, _124_v0, _dafny.IntOfUint32((_150_v23).Cardinality()))
		var _152_v25 D3
		_ = _152_v25
		_152_v25 = Companion_D3_.Create_DC7_(_151_v24, _129_v5)
		_129_v5 = (_152_v25).Dtor_cf15()
		var _153_v26 _dafny.Sequence
		_ = _153_v26
		_153_v26 = _dafny.UnicodeSeqOfUtf8Bytes("dbh")
		var _154_v27 _dafny.MultiSet
		_ = _154_v27
		_154_v27 = _dafny.MultiSetOf(_153_v26, _dafny.Companion_Sequence_.Update(_153_v26, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_150_v23).Cardinality()), _dafny.IntOfUint32((_153_v26).Cardinality()))).Uint32(), _149_v22), _153_v26)
		_154_v27 = _154_v27
		var _155_v28 *C1
		_ = _155_v28
		var _nw28 *C1 = New_C1_()
		_ = _nw28
		_nw28.Ctor__()
		_155_v28 = _nw28
		_155_v28 = _155_v28
		var _156_v29 bool
		_ = _156_v29
		_156_v29 = !(_129_v5)
		var _157_v30 *C0
		_ = _157_v30
		var _nw29 *C0 = New_C0_()
		_ = _nw29
		_nw29.Ctor__(Companion_Default___.Fm5(_131_globalState), (_148_i1).Plus(_124_v0), _156_v29)
		_157_v30 = _nw29
	}
	var _158_v31 _dafny.Sequence
	_ = _158_v31
	_158_v31 = _dafny.UnicodeSeqOfUtf8Bytes("nt")
	var _hi2 _dafny.Int = _124_v0
	_ = _hi2
	for _159_i2 := _dafny.IntOfUint32((_158_v31).Cardinality()); _159_i2.Cmp(_hi2) < 0; _159_i2 = _159_i2.Plus(_dafny.One) {
		var _160_v32 D2
		_ = _160_v32
		_160_v32 = Companion_D2_.Create_DC4_(_129_v5, _124_v0)
		var _161_v33 _dafny.Array
		_ = _161_v33
		var _nwElement0_3 bool = _129_v5
		_ = _nwElement0_3
		var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(8))
		_ = _nw30
		(_nw30).ArraySet1(_nwElement0_3, 0)
		(_nw30).ArraySet1(_129_v5, 1)
		(_nw30).ArraySet1(_129_v5, 2)
		(_nw30).ArraySet1(_129_v5, 3)
		(_nw30).ArraySet1(_129_v5, 4)
		(_nw30).ArraySet1(_129_v5, 5)
		(_nw30).ArraySet1((_160_v32).Dtor_cf6(), 6)
		(_nw30).ArraySet1(_129_v5, 7)
		_161_v33 = _nw30
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))
		_ = _index19
		(_161_v33).ArraySet1((_159_i2).Cmp(_159_i2) > 0, (_index19).Int())
		var _162_v34 _dafny.CodePoint
		_ = _162_v34
		_162_v34 = _dafny.CodePoint('u')
		var _163_v35 _dafny.Map
		_ = _163_v35
		_163_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_i2, _159_i2)
		var _164_v36 _dafny.Set
		_ = _164_v36
		_164_v36 = _dafny.SetOf(_158_v31)
		var _165_v37 D2
		_ = _165_v37
		_165_v37 = Companion_D2_.Create_DC5_((func() _dafny.Map {
			if _129_v5 {
				return Companion_Default___.Fm6(Companion_Default___.Fm3(_162_v34, _129_v5, _129_v5, _131_globalState), _129_v5, _159_i2, _131_globalState)
			}
			return _163_v35
		})(), (Companion_Default___.Fm7(_129_v5, _129_v5, _131_globalState)).IsDisjointFrom(_164_v36), _129_v5, _129_v5, _129_v5)
		var _166_v38 _dafny.Sequence
		_ = _166_v38
		_166_v38 = _dafny.SeqOf(_129_v5)
		var _167_v39 _dafny.Map
		_ = _167_v39
		_167_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_166_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_166_v38).Cardinality()), _dafny.IntOfUint32((_166_v38).Cardinality()))).Uint32()).(bool), _129_v5)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))
		_ = _index20
		var _rhs26 bool = !(_129_v5)
		_ = _rhs26
		var _rhs27 bool = ((func() bool {
			if (_167_v39).Contains(false) {
				return (_167_v39).Get(false).(bool)
			}
			return !(_129_v5)
		})()) == ((_159_i2).Cmp(Companion_Default___.Fm3(_162_v34, _129_v5, _129_v5, _131_globalState)) >= 0)
		_ = _rhs27
		var _rhs28 D2 = Companion_D2_.Create_DC5_(_163_v35, _129_v5, _129_v5, _129_v5, _129_v5)
		_ = _rhs28
		var _lhs10 _dafny.Array = _161_v33
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))
		_ = _lhs11
		_129_v5 = _rhs26
		(_lhs10).ArraySet1(_rhs27, (_lhs11).Int())
		_165_v37 = _rhs28
		var _168_v40 _dafny.Map
		_ = _168_v40
		_168_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _127_v3)
		var _pat_let_tv0 = _162_v34
		_ = _pat_let_tv0
		var _source0 D1 = func(_pat_let0_0 D1) D1 {
			return func(_169_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let1_0 _dafny.CodePoint) D1 {
					return func(_170_dt__update_hcf2_h0 _dafny.CodePoint) D1 {
						return Companion_D1_.Create_DC2_(_170_dt__update_hcf2_h0, (_169_dt__update__tmp_h0).Dtor_cf3(), (_169_dt__update__tmp_h0).Dtor_cf4())
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}((func() D1 {
			if (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool) {
				return Companion_D1_.Create_DC2_(_162_v34, _124_v0, (_168_v40).Cardinality())
			}
			return Companion_D1_.Create_DC2_(_162_v34, _dafny.IntOfUint32((_158_v31).Cardinality()), (_dafny.Zero).Minus(_124_v0))
		})())
		_ = _source0
		if _source0.Is_DC2() {
			var _171___mcc_h0 _dafny.CodePoint = _source0.Get_().(D1_DC2).Cf2
			_ = _171___mcc_h0
			var _172___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC2).Cf3
			_ = _172___mcc_h1
			var _173___mcc_h2 _dafny.Int = _source0.Get_().(D1_DC2).Cf4
			_ = _173___mcc_h2
			var _174_cf4 _dafny.Int = _173___mcc_h2
			_ = _174_cf4
			var _175_cf3 _dafny.Int = _172___mcc_h1
			_ = _175_cf3
			var _176_cf2 _dafny.CodePoint = _171___mcc_h0
			_ = _176_cf2
			var _177_v41 _dafny.Sequence
			_ = _177_v41
			_177_v41 = _dafny.SeqOf(_124_v0, _159_i2)
			var _178_v42 _dafny.Map
			_ = _178_v42
			_178_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_166_v38).Select((Companion_Default___.SafeIndex(_124_v0, _dafny.IntOfUint32((_166_v38).Cardinality()))).Uint32()).(bool), _dafny.MultiSetFromSeq(_177_v41))
			var _179_v43 bool
			_ = _179_v43
			_179_v43 = _129_v5
			var _180_v44 *C0
			_ = _180_v44
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(_178_v42, _dafny.IntOfUint32((_158_v31).Cardinality()), _179_v43)
			_180_v44 = _nw31
			_180_v44 = _180_v44
			_127_v3 = _127_v3
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))
			_ = _index21
			(_161_v33).ArraySet1(_dafny.Companion_Sequence_.Equal(_158_v31, _dafny.UnicodeSeqOfUtf8Bytes("nwf")), (_index21).Int())
			var _181_v45 _dafny.Map
			_ = _181_v45
			_181_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v34, _129_v5)
			_175_cf3 = ((_181_v45).Merge(_181_v45)).Cardinality()
		} else {
			var _182___mcc_h3 _dafny.Map = _source0.Get_().(D1_DC1).Cf1
			_ = _182___mcc_h3
			var _183_cf1 _dafny.Map = _182___mcc_h3
			_ = _183_cf1
			_129_v5 = (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool)
			var _184_v46 *C1
			_ = _184_v46
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__()
			_184_v46 = _nw32
			var _185_v47 _dafny.Array
			_ = _185_v47
			var _nwElement0_4 *C1 = _184_v46
			_ = _nwElement0_4
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(15))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_4, 0)
			(_nw33).ArraySet1(_184_v46, 1)
			(_nw33).ArraySet1(_184_v46, 2)
			(_nw33).ArraySet1(_184_v46, 3)
			(_nw33).ArraySet1(_184_v46, 4)
			(_nw33).ArraySet1((func() *C1 {
				if (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool) {
					return _184_v46
				}
				return _184_v46
			})(), 5)
			(_nw33).ArraySet1(_184_v46, 6)
			(_nw33).ArraySet1(_184_v46, 7)
			(_nw33).ArraySet1(_184_v46, 8)
			(_nw33).ArraySet1(_184_v46, 9)
			(_nw33).ArraySet1(_184_v46, 10)
			(_nw33).ArraySet1(_184_v46, 11)
			(_nw33).ArraySet1(_184_v46, 12)
			(_nw33).ArraySet1((func() *C1 {
				if _129_v5 {
					return _184_v46
				}
				return _184_v46
			})(), 13)
			(_nw33).ArraySet1(_184_v46, 14)
			_185_v47 = _nw33
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_185_v47), 0))
			_ = _index22
			(_185_v47).ArraySet1(_184_v46, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_185_v47), 0))
			_ = _index23
			(_185_v47).ArraySet1(_184_v46, (_index23).Int())
			var _186_v48 D4
			_ = _186_v48
			_186_v48 = Companion_D4_.Create_DC9_(_161_v33)
			_161_v33 = (_186_v48).Dtor_cf17()
			_127_v3 = _127_v3
		}
		if !(_129_v5) {
			var _187_v49 _dafny.Array
			_ = _187_v49
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
			_ = _nw34
			_187_v49 = _nw34
			var _188_v50 _dafny.Sequence
			_ = _188_v50
			_188_v50 = _dafny.SeqOf(_159_i2, _124_v0, _159_i2)
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_187_v49), 0))
			_ = _index24
			(_187_v49).ArraySet1(_dafny.SeqOf((_188_v50).Select((Companion_Default___.SafeIndex(_124_v0, _dafny.IntOfUint32((_188_v50).Cardinality()))).Uint32()).(_dafny.Int), (_163_v35).Cardinality(), _159_i2, (_125_v1).Cardinality()), (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_187_v49), 0))
			_ = _index25
			(_187_v49).ArraySet1(_188_v50, (_index25).Int())
			_129_v5 = !(!(Companion_Default___.Fm8(_124_v0, _163_v35, (func() _dafny.Int {
				if (_125_v1).Contains(_dafny.IntOfInt64(254)) {
					return (_125_v1).Multiplicity(_dafny.IntOfInt64(254))
				}
				return _dafny.IntOfUint32((_166_v38).Cardinality())
			})(), _129_v5, _131_globalState)) || (false))
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))
			_ = _index26
			(_161_v33).ArraySet1((func() bool {
				if (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool) {
					return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool)), _dafny.SeqOf(_129_v5))
				}
				return (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool)
			})(), (_index26).Int())
			var _189_v51 _dafny.Map
			_ = _189_v51
			_189_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool), (_125_v1).Difference(_dafny.MultiSetOf(_124_v0)))
			_189_v51 = (_189_v51).Update((_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool), Companion_Default___.Fm2(_131_globalState))
			var _190_v52 _dafny.Map
			_ = _190_v52
			_190_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("lpuwc"), Companion_Default___.Fm3(_162_v34, (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool), _129_v5, _131_globalState))
			var _191_v53 bool
			_ = _191_v53
			_191_v53 = (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool)
			var _192_v54 *C0
			_ = _192_v54
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__(_189_v51, (_125_v1).Cardinality(), _191_v53)
			_192_v54 = _nw35
			var _193_v55 _dafny.Map
			_ = _193_v55
			_193_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v54, _129_v5)
			_190_v52 = (_190_v52).Update(_dafny.UnicodeSeqOfUtf8Bytes("bo"), ((_193_v55).Update(_192_v54, (_161_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))).Int()).(bool))).Cardinality())
		} else {
			_124_v0 = _124_v0
			var _194_v56 _dafny.Sequence
			_ = _194_v56
			_194_v56 = _dafny.SeqOf(_127_v3)
			var _195_v57 D5
			_ = _195_v57
			_195_v57 = Companion_D5_.Create_DC11_(_194_v56)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_161_v33), 0))
			_ = _index27
			(_161_v33).ArraySet1(Companion_Default___.Fm8(_159_i2, _163_v35, _124_v0, Companion_Default___.Fm8(_dafny.IntOfUint32(((_195_v57).Dtor_cf18()).Cardinality()), (_163_v35).Update(_159_i2, _124_v0), _159_i2, _129_v5, _131_globalState), _131_globalState), (_index27).Int())
			_194_v56 = _dafny.SeqOf(_127_v3, _127_v3, _127_v3, _127_v3, _127_v3)
			_124_v0 = _159_i2
			var _196_v58 *C1
			_ = _196_v58
			var _nw36 *C1 = New_C1_()
			_ = _nw36
			_nw36.Ctor__()
			_196_v58 = _nw36
		}
		_129_v5 = Companion_Default___.Fm8(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(700))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_197_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
			return func(_198_i3 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}((func(_199_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_200_i4 _dafny.Int) _dafny.CodePoint {
						return _199_v34
					}
				})(_197_v34)))).Cardinality())
			}
		})(_162_v34)))).Cardinality()), _163_v35, _159_i2, _129_v5, _131_globalState)
	}
	if _129_v5 {
		var _201_v59 _dafny.Array
		_ = _201_v59
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_4
		var _nw37 _dafny.Array
		_ = _nw37
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw37 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) bool = (func(_202_v5 bool) func(_dafny.Int) bool {
				return func(_203_i5 _dafny.Int) bool {
					return _202_v5
				}
			})(_129_v5)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw37 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw37).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw37).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_201_v59 = _nw37
		var _204_v60 _dafny.CodePoint
		_ = _204_v60
		_204_v60 = _dafny.CodePoint('m')
		var _205_v61 _dafny.Map
		_ = _205_v61
		_205_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_204_v60, _124_v0)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_201_v59), 0))
		_ = _index28
		(_201_v59).ArraySet1(!((_125_v1).IsDisjointFrom(_dafny.MultiSetOf(_124_v0, _124_v0, _124_v0, (_205_v61).Cardinality()))), (_index28).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_201_v59), 0))
		_ = _index29
		(_201_v59).ArraySet1(_129_v5, (_index29).Int())
		var _206_v62 _dafny.Set
		_ = _206_v62
		_206_v62 = _dafny.SetOf(_201_v59)
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_201_v59), 0))
		_ = _index30
		(_201_v59).ArraySet1((_dafny.SetOf(_201_v59, _201_v59)).IsSubsetOf((_206_v62).Intersection(_206_v62)), (_index30).Int())
		_124_v0 = _dafny.IntOfInt64(705)
		var _207_v63 _dafny.Map
		_ = _207_v63
		_207_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _125_v1)
		var _208_v64 bool
		_ = _208_v64
		_208_v64 = (_201_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_201_v59), 0))).Int()).(bool)
		var _209_v65 *C0
		_ = _209_v65
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__(_207_v63, _124_v0, _208_v64)
		_209_v65 = _nw38
		var _210_v66 _dafny.Map
		_ = _210_v66
		_210_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_201_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_201_v59), 0))).Int()).(bool), _124_v0)
		var _211_v67 _dafny.Set
		_ = _211_v67
		_211_v67 = _dafny.SetOf((_130_v6).Cardinality(), _dafny.IntOfInt64(238), (_210_v66).Cardinality())
		var _212_v68 _dafny.Sequence
		_ = _212_v68
		_212_v68 = _dafny.SeqOf(_211_v67)
		var _213_v69 D6
		_ = _213_v69
		_213_v69 = Companion_D6_.Create_DC15_(_209_v65)
		var _214_v70 D7
		_ = _214_v70
		_214_v70 = Companion_D7_.Create_DC19_(_212_v68)
		var _rhs29 *C0 = (_213_v69).Dtor_cf21()
		_ = _rhs29
		var _rhs30 _dafny.Int = _124_v0
		_ = _rhs30
		var _rhs31 _dafny.Sequence = (_214_v70).Dtor_cf25()
		_ = _rhs31
		_209_v65 = _rhs29
		_124_v0 = _rhs30
		_212_v68 = _rhs31
		var _215_v71 *C1
		_ = _215_v71
		var _nw39 *C1 = New_C1_()
		_ = _nw39
		_nw39.Ctor__()
		_215_v71 = _nw39
		_215_v71 = _215_v71
	} else {
		var _216_v72 _dafny.Map
		_ = _216_v72
		_216_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_124_v0), _124_v0)
		var _217_v73 _dafny.MultiSet
		_ = _217_v73
		_217_v73 = _dafny.MultiSetOf(_129_v5)
		_216_v72 = (_216_v72).Update(((_217_v73).Cardinality()).Plus(_124_v0), (_dafny.Zero).Minus(_124_v0))
		_124_v0 = (_124_v0).Minus(Companion_Default___.SafeModuloInt(_124_v0, _124_v0))
		var _218_v74 _dafny.CodePoint
		_ = _218_v74
		_218_v74 = _dafny.CodePoint('n')
		_158_v31 = _dafny.Companion_Sequence_.Update(_158_v31, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.IntOfUint32((_158_v31).Cardinality()))).Uint32(), _218_v74)
		var _219_v75 _dafny.Array
		_ = _219_v75
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_5
		var _nw40 _dafny.Array
		_ = _nw40
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw40 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Sequence = (func(_220_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_221_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_220_v0, _dafny.IntOfInt64(535))
				}
			})(_124_v0)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw40 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw40).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw40).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_219_v75 = _nw40
		var _222_v76 _dafny.Sequence
		_ = _222_v76
		_222_v76 = _dafny.SeqOf(_dafny.IntOfInt64(117))
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_219_v75), 0))
		_ = _index31
		(_219_v75).ArraySet1(_dafny.Companion_Sequence_.Update(_222_v76, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.IntOfUint32((_222_v76).Cardinality()))).Uint32(), _dafny.IntOfInt64(462)), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_219_v75), 0))
		_ = _index32
		(_219_v75).ArraySet1(_222_v76, (_index32).Int())
		var _223_v77 _dafny.Sequence
		_ = _223_v77
		var _224_v78 *C0
		_ = _224_v78
		var _225_v79 _dafny.Sequence
		_ = _225_v79
		var _226_v80 bool
		_ = _226_v80
		var _out2 _dafny.Sequence
		_ = _out2
		var _out3 *C0
		_ = _out3
		var _out4 _dafny.Sequence
		_ = _out4
		var _out5 bool
		_ = _out5
		_out2, _out3, _out4, _out5 = Companion_Default___.M1(_218_v74, _129_v5, _131_globalState)
		_223_v77 = _out2
		_224_v78 = _out3
		_225_v79 = _out4
		_226_v80 = _out5
	}
	var _227_v81 _dafny.Map
	_ = _227_v81
	_227_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v0, _129_v5)
	_227_v81 = (_227_v81).Update(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_124_v0)).Cardinality()), _dafny.IntOfInt64(623)), _129_v5)
	if true {
		var _228_v82 _dafny.CodePoint
		_ = _228_v82
		_228_v82 = _dafny.CodePoint('c')
		var _229_v83 _dafny.Sequence
		_ = _229_v83
		var _230_v84 *C0
		_ = _230_v84
		var _231_v85 _dafny.Sequence
		_ = _231_v85
		var _232_v86 bool
		_ = _232_v86
		var _out6 _dafny.Sequence
		_ = _out6
		var _out7 *C0
		_ = _out7
		var _out8 _dafny.Sequence
		_ = _out8
		var _out9 bool
		_ = _out9
		_out6, _out7, _out8, _out9 = Companion_Default___.M1((func() _dafny.CodePoint {
			if true {
				return _228_v82
			}
			return _228_v82
		})(), _129_v5, _131_globalState)
		_229_v83 = _out6
		_230_v84 = _out7
		_231_v85 = _out8
		_232_v86 = _out9
		_129_v5 = _129_v5
		_124_v0 = _124_v0
		var _233_v87 _dafny.Sequence
		_ = _233_v87
		var _234_v88 *C0
		_ = _234_v88
		var _235_v89 _dafny.Sequence
		_ = _235_v89
		var _236_v90 bool
		_ = _236_v90
		var _out10 _dafny.Sequence
		_ = _out10
		var _out11 *C0
		_ = _out11
		var _out12 _dafny.Sequence
		_ = _out12
		var _out13 bool
		_ = _out13
		_out10, _out11, _out12, _out13 = Companion_Default___.M1(_228_v82, _129_v5, _131_globalState)
		_233_v87 = _out10
		_234_v88 = _out11
		_235_v89 = _out12
		_236_v90 = _out13
		_236_v90 = true
	} else {
		var _237_v91 _dafny.Sequence
		_ = _237_v91
		_237_v91 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(162))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_238_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))).Cardinality()))
		var _239_v92 _dafny.Map
		_ = _239_v92
		_239_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v5, _dafny.MultiSetFromSeq(_237_v91))
		var _240_v93 _dafny.CodePoint
		_ = _240_v93
		_240_v93 = _dafny.CodePoint('x')
		var _241_v94 *C0
		_ = _241_v94
		var _nw41 *C0 = New_C0_()
		_ = _nw41
		_nw41.Ctor__(_239_v92, Companion_Default___.Fm3(_240_v93, _129_v5, _129_v5, _131_globalState), false)
		_241_v94 = _nw41
		var _242_v95 _dafny.Sequence
		_ = _242_v95
		_242_v95 = _dafny.SeqOf(_241_v94)
		var _243_v96 D6
		_ = _243_v96
		_243_v96 = Companion_D6_.Create_DC15_(_241_v94)
		var _244_v97 _dafny.Sequence
		_ = _244_v97
		_244_v97 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_242_v95, _242_v95), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_158_v31).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_242_v95, _242_v95)).Cardinality()))).Uint32(), (_243_v96).Dtor_cf21()), _242_v95, _242_v95)
		_244_v97 = _dafny.SeqOf(_242_v95)
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_127_v3), 0))
		_ = _index33
		(_127_v3).ArraySet1((_124_v0).Minus(_124_v0), (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_127_v3), 0))
		_ = _index34
		(_127_v3).ArraySet1((_124_v0).Times(_124_v0), (_index34).Int())
		var _245_v98 _dafny.Sequence
		_ = _245_v98
		_245_v98 = _dafny.SeqOf(false, _129_v5, _129_v5)
		var _246_v99 _dafny.Sequence
		_ = _246_v99
		_246_v99 = _dafny.SeqOf(_245_v98, _245_v98, _dafny.Companion_Sequence_.Update(_245_v98, (Companion_Default___.SafeIndex((_127_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_127_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_245_v98).Cardinality()))).Uint32(), _129_v5), _245_v98, _245_v98)
		var _247_v100 bool
		_ = _247_v100
		_247_v100 = false
		var _248_v101 *C0
		_ = _248_v101
		var _nw42 *C0 = New_C0_()
		_ = _nw42
		_nw42.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v5, _dafny.MultiSetOf((_127_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_127_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_127_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_127_v3), 0))).Int()).(_dafny.Int), _124_v0)).Cardinality()))), _dafny.IntOfUint32((_246_v99).Cardinality()), !(_129_v5))
		_248_v101 = _nw42
		var _249_v103 _dafny.Set
		_ = _249_v103
		_249_v103 = _dafny.SetOf(_227_v81, func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(930), _dafny.IntOfInt64(440))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _250_v102 _dafny.Int
				_250_v102 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(930)).Cmp(_250_v102) <= 0) && ((_250_v102).Cmp(_dafny.IntOfInt64(440)) < 0) {
					_coll7.Add((_250_v102).Times((_127_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_127_v3), 0))).Int()).(_dafny.Int)), true)
				}
			}
			return _coll7.ToMap()
		}())
		var _251_v104 _dafny.Map
		_ = _251_v104
		_251_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_245_v98).Cardinality()), _249_v103)
		_251_v104 = (_251_v104).Update(_dafny.IntOfInt64(-48), _249_v103)
		var _252_v105 _dafny.Array
		_ = _252_v105
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_6
		var _nw43 _dafny.Array
		_ = _nw43
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw43 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Map = (func(_253_v5 bool) func(_dafny.Int) _dafny.Map {
				return func(_254_i8 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_253_v5), _253_v5)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(_253_v5)))
				}
			})(_129_v5)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw43 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw43).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw43).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_252_v105 = _nw43
		var _255_v106 _dafny.Map
		_ = _255_v106
		_255_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_129_v5), _129_v5)
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_252_v105), 0))
		_ = _index35
		(_252_v105).ArraySet1((_255_v106).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v5, _129_v5)), (_index35).Int())
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_252_v105), 0))
		_ = _index36
		(_252_v105).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _129_v5), (_index36).Int())
	}
	var _hi3 _dafny.Int = _124_v0
	_ = _hi3
	for _256_i9 := _124_v0; _256_i9.Cmp(_hi3) < 0; _256_i9 = _256_i9.Plus(_dafny.One) {
		var _257_v107 _dafny.Map
		_ = _257_v107
		_257_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _125_v1)
		var _258_v108 bool
		_ = _258_v108
		_258_v108 = _129_v5
		var _259_v109 *C0
		_ = _259_v109
		var _nw44 *C0 = New_C0_()
		_ = _nw44
		_nw44.Ctor__(_257_v107, (_124_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_158_v31).Cardinality()))), _258_v108)
		_259_v109 = _nw44
		_259_v109 = _259_v109
		var _260_v110 _dafny.Array
		_ = _260_v110
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw45
		_260_v110 = _nw45
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_260_v110), 0))
		_ = _index37
		(_260_v110).ArraySet1(_129_v5, (_index37).Int())
		var _261_v111 _dafny.MultiSet
		_ = _261_v111
		_261_v111 = _dafny.MultiSetOf(!(false) || (_129_v5), _129_v5)
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_260_v110), 0))
		_ = _index38
		var _rhs32 _dafny.Int = _124_v0
		_ = _rhs32
		var _rhs33 _dafny.MultiSet = (_125_v1).Intersection(_125_v1)
		_ = _rhs33
		var _rhs34 bool = _129_v5
		_ = _rhs34
		var _rhs35 _dafny.Int = (func() _dafny.Int {
			if (_261_v111).Contains(false) {
				return (_261_v111).Multiplicity(false)
			}
			return _124_v0
		})()
		_ = _rhs35
		var _lhs12 _dafny.Array = _260_v110
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_260_v110), 0))
		_ = _lhs13
		_124_v0 = _rhs32
		_125_v1 = _rhs33
		(_lhs12).ArraySet1(_rhs34, (_lhs13).Int())
		_124_v0 = _rhs35
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_260_v110), 0))
		_ = _index39
		(_260_v110).ArraySet1((((_261_v111).Intersection(_261_v111)).Cardinality()).Cmp(_256_i9) != 0, (_index39).Int())
		var _262_v112 *C1
		_ = _262_v112
		var _nw46 *C1 = New_C1_()
		_ = _nw46
		_nw46.Ctor__()
		_262_v112 = _nw46
	}
	var _263_v113 _dafny.Array
	_ = _263_v113
	var _nw47 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
	_ = _nw47
	_263_v113 = _nw47
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))
	_ = _index40
	(_263_v113).ArraySet1(true, (_index40).Int())
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))
	_ = _index41
	(_263_v113).ArraySet1(_129_v5, (_index41).Int())
	var _264_v114 _dafny.Map
	_ = _264_v114
	_264_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), (_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool))
	var _265_v115 _dafny.Sequence
	_ = _265_v115
	_265_v115 = _dafny.SeqOf(_124_v0)
	_264_v114 = (_264_v114).Update((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), _dafny.Companion_Sequence_.Contains(_265_v115, _124_v0))
	var _266_v116 _dafny.CodePoint
	_ = _266_v116
	_266_v116 = _dafny.CodePoint('i')
	_266_v116 = _266_v116
	_266_v116 = _dafny.CodePoint('h')
	var _267_v117 _dafny.Map
	_ = _267_v117
	_267_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v0, _124_v0)
	var _268_v118 _dafny.Map
	_ = _268_v118
	_268_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_267_v117).Cardinality(), _124_v0)
	var _269_v119 _dafny.Sequence
	_ = _269_v119
	_269_v119 = _dafny.SeqOf(_129_v5)
	var _hi4 _dafny.Int = ((func() _dafny.Int {
		if (_268_v118).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(329))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}((func(_270_v116 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_271_i11 _dafny.Int) _dafny.CodePoint {
				return _270_v116
			}
		})(_266_v116)))).Cardinality())) {
			return (_268_v118).Get(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(329))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_272_v116 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_273_i11 _dafny.Int) _dafny.CodePoint {
					return _272_v116
				}
			})(_266_v116)))).Cardinality())).(_dafny.Int)
		}
		return _124_v0
	})()).Minus(_dafny.IntOfUint32((_269_v119).Cardinality()))
	_ = _hi4
	for _274_i10 := (_124_v0).Times(_124_v0); _274_i10.Cmp(_hi4) < 0; _274_i10 = _274_i10.Plus(_dafny.One) {
		_129_v5 = !(_129_v5)
		if _129_v5 {
			var _275_v120 _dafny.Set
			_ = _275_v120
			_275_v120 = _dafny.SetOf(_266_v116)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))
			_ = _index42
			(_263_v113).ArraySet1(((_dafny.IntOfInt64(-528)).Minus(_274_i10)).Cmp((_275_v120).Cardinality()) >= 0, (_index42).Int())
			var _276_v121 _dafny.Sequence
			_ = _276_v121
			var _277_v122 *C0
			_ = _277_v122
			var _278_v123 _dafny.Sequence
			_ = _278_v123
			var _279_v124 bool
			_ = _279_v124
			var _out14 _dafny.Sequence
			_ = _out14
			var _out15 *C0
			_ = _out15
			var _out16 _dafny.Sequence
			_ = _out16
			var _out17 bool
			_ = _out17
			_out14, _out15, _out16, _out17 = Companion_Default___.M1((_158_v31).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_124_v0), _dafny.IntOfUint32((_158_v31).Cardinality()))).Uint32()).(_dafny.CodePoint), (_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), _131_globalState)
			_276_v121 = _out14
			_277_v122 = _out15
			_278_v123 = _out16
			_279_v124 = _out17
			var _280_v125 bool
			_ = _280_v125
			_280_v125 = _129_v5
			var _281_v126 T0
			_ = _281_v126
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__(Companion_Default___.Fm5(_131_globalState), Companion_Default___.SafeDivisionInt(_274_i10, _dafny.IntOfInt64(-74)), _280_v125)
			_281_v126 = _nw48
			_281_v126 = _281_v126
			_124_v0 = _dafny.IntOfInt64(201)
			_124_v0 = (Companion_Default___.Fm9(_274_i10, _279_v124, (_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), _dafny.IntOfUint32((_276_v121).Cardinality()), _131_globalState)).Dtor_cf4()
		} else {
			var _282_v127 _dafny.Map
			_ = _282_v127
			_282_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), Companion_Default___.Fm3(_266_v116, _129_v5, (_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), _131_globalState))
			_282_v127 = (_282_v127).Update(!((_124_v0).Cmp(_274_i10) != 0), _dafny.IntOfUint32((_269_v119).Cardinality()))
			var _283_v128 _dafny.Map
			_ = _283_v128
			_283_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v6, _129_v5)
			_283_v128 = (_283_v128).Update(_130_v6, _129_v5)
			_265_v115 = _dafny.Companion_Sequence_.Concatenate(_265_v115, _265_v115)
			_129_v5 = ((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool)) || ((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool))
			var _284_v129 _dafny.Sequence
			_ = _284_v129
			var _285_v130 *C0
			_ = _285_v130
			var _286_v131 _dafny.Sequence
			_ = _286_v131
			var _287_v132 bool
			_ = _287_v132
			var _out18 _dafny.Sequence
			_ = _out18
			var _out19 *C0
			_ = _out19
			var _out20 _dafny.Sequence
			_ = _out20
			var _out21 bool
			_ = _out21
			_out18, _out19, _out20, _out21 = Companion_Default___.M1((func() _dafny.CodePoint {
				if _129_v5 {
					return _266_v116
				}
				return _266_v116
			})(), (_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), _131_globalState)
			_284_v129 = _out18
			_285_v130 = _out19
			_286_v131 = _out20
			_287_v132 = _out21
		}
		var _288_v133 _dafny.Map
		_ = _288_v133
		_288_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool), _124_v0)
		var _289_v134 _dafny.Map
		_ = _289_v134
		_289_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v5, _125_v1)
		var _290_v135 bool
		_ = _290_v135
		_290_v135 = (_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool)
		var _291_v136 T0
		_ = _291_v136
		var _nw49 *C0 = New_C0_()
		_ = _nw49
		_nw49.Ctor__(_289_v134, _dafny.IntOfUint32((_158_v31).Cardinality()), _290_v135)
		_291_v136 = _nw49
		var _292_v137 _dafny.Map
		_ = _292_v137
		_292_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v136, _158_v31)
		_268_v118 = (_268_v118).Update(_124_v0, (func() _dafny.Int {
			if (_288_v133).Contains((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool)) {
				return (_288_v133).Get((_263_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_263_v113), 0))).Int()).(bool)).(_dafny.Int)
			}
			return (_292_v137).Cardinality()
		})())
		_124_v0 = _274_i10
	}
	_129_v5 = _129_v5
	var _293_v138 _dafny.Sequence
	_ = _293_v138
	_293_v138 = _dafny.SeqOf(_127_v3, _127_v3, _127_v3)
	var _294_v139 D5
	_ = _294_v139
	_294_v139 = Companion_D5_.Create_DC11_(_293_v138)
	var _rhs36 _dafny.Array = (func() _dafny.Array {
		if _129_v5 {
			return _263_v113
		}
		return _263_v113
	})()
	_ = _rhs36
	var _rhs37 D5 = _294_v139
	_ = _rhs37
	_263_v113 = _rhs36
	_294_v139 = _rhs37
	_129_v5 = _129_v5
	var _295_v140 _dafny.Array
	_ = _295_v140
	var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
	_ = _nw50
	_295_v140 = _nw50
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_295_v140), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _296_i12 _dafny.Int
		_296_i12 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_296_i12).Sign() != -1) && ((_296_i12).Cmp(_dafny.ArrayLen((_295_v140), 0)) < 0)) {
			(_295_v140).ArraySet1(_265_v115, (_296_i12).Int())
		}
	}
	_dafny.Print(_124_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v1).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197), _dafny.IntOfInt64(197), _dafny.IntOfInt64(265))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v2).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v4).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_129_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v6).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197), _dafny.IntOfInt64(197), _dafny.IntOfInt64(265))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197), _dafny.IntOfInt64(197), _dafny.IntOfInt64(265))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_131_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(197))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F3()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState.F5).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_158_v31.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_227_v81).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(196), false).UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_263_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_264_v114).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_265_v115, _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_266_v116)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v117).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(196), _dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v118).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_269_v119, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_293_v138).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_294_v139).Dtor_cf18()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_295_v140).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(196))))
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
	Cf2 _dafny.CodePoint
	Cf3 _dafny.Int
	Cf4 _dafny.Int
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.CodePoint, Cf3 _dafny.Int, Cf4 _dafny.Int) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4}}
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

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.CodePoint('D'), _dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf2() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf1() _dafny.Map {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
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

type D2_DC4 struct {
	Cf6 bool
	Cf7 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf6 bool, Cf7 _dafny.Int) D2 {
	return D2{D2_DC4{Cf6, Cf7}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC5 struct {
	Cf8  _dafny.Map
	Cf9  bool
	Cf10 bool
	Cf11 bool
	Cf12 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf8 _dafny.Map, Cf9 bool, Cf10 bool, Cf11 bool, Cf12 bool) D2 {
	return D2{D2_DC5{Cf8, Cf9, Cf10, Cf11, Cf12}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC3 struct {
	Cf5 _dafny.Map
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf5 _dafny.Map) D2 {
	return D2{D2_DC3{Cf5}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(false, _dafny.Zero)
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Map {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC5).Cf11
}

func (_this D2) Dtor_cf12() bool {
	return _this.Get_().(D2_DC5).Cf12
}

func (_this D2) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D2_DC3).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf5) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf8.Equals(data2.Cf8) && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
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
	Cf14 D1
	Cf15 bool
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf14 D1, Cf15 bool) D3 {
	return D3{D3_DC7{Cf14, Cf15}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC6 struct {
	Cf13 _dafny.MultiSet
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf13 _dafny.MultiSet) D3 {
	return D3{D3_DC6{Cf13}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC8 struct {
	Cf16 D3
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf16 D3) D3 {
	return D3{D3_DC8{Cf16}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_(Companion_D1_.Default(), false)
}

func (_this D3) Dtor_cf14() D1 {
	return _this.Get_().(D3_DC7).Cf14
}

func (_this D3) Dtor_cf15() bool {
	return _this.Get_().(D3_DC7).Cf15
}

func (_this D3) Dtor_cf13() _dafny.MultiSet {
	return _this.Get_().(D3_DC6).Cf13
}

func (_this D3) Dtor_cf16() D3 {
	return _this.Get_().(D3_DC8).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf14.Equals(data2.Cf14) && data1.Cf15 == data2.Cf15
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf13.Equals(data2.Cf13)
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D4_DC10 struct {
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_() D4 {
	return D4{D4_DC10{}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC9 struct {
	Cf17 _dafny.Array
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf17 _dafny.Array) D4 {
	return D4{D4_DC9{Cf17}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_()
}

func (_this D4) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf17) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC10:
		{
			_, ok := other.Get_().(D4_DC10)
			return ok
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
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
	Cf19 bool
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf19 bool) D5 {
	return D5{D5_DC12{Cf19}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC13 struct {
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_() D5 {
	return D5{D5_DC13{}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC11 struct {
	Cf18 _dafny.Sequence
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf18 _dafny.Sequence) D5 {
	return D5{D5_DC11{Cf18}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC14 struct {
	Cf20 D5
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf20 D5) D5 {
	return D5{D5_DC14{Cf20}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC12_(false)
}

func (_this D5) Dtor_cf19() bool {
	return _this.Get_().(D5_DC12).Cf19
}

func (_this D5) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D5_DC11).Cf18
}

func (_this D5) Dtor_cf20() D5 {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf20) + ")"
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
			return ok && data1.Cf19 == data2.Cf19
		}
	case D5_DC13:
		{
			_, ok := other.Get_().(D5_DC13)
			return ok
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf20.Equals(data2.Cf20)
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

type D6_DC16 struct {
	Cf22 bool
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 bool) D6 {
	return D6{D6_DC16{Cf22}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC17 struct {
	Cf23 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf23 bool) D6 {
	return D6{D6_DC17{Cf23}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf24 bool
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf24 bool) D6 {
	return D6{D6_DC18{Cf24}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC15 struct {
	Cf21 *C0
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf21 *C0) D6 {
	return D6{D6_DC15{Cf21}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(false)
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) Dtor_cf23() bool {
	return _this.Get_().(D6_DC17).Cf23
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC18).Cf24
}

func (_this D6) Dtor_cf21() *C0 {
	return _this.Get_().(D6_DC15).Cf21
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf21) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf22 == data2.Cf22
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf21 == data2.Cf21
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

type D7_DC20 struct {
	Cf26 _dafny.Array
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf26 _dafny.Array) D7 {
	return D7{D7_DC20{Cf26}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC21 struct {
	Cf27 bool
	Cf28 T0
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf27 bool, Cf28 T0) D7 {
	return D7{D7_DC21{Cf27, Cf28}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC19 struct {
	Cf25 _dafny.Sequence
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf25 _dafny.Sequence) D7 {
	return D7{D7_DC19{Cf25}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D7) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D7_DC20).Cf26
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC21).Cf27
}

func (_this D7) Dtor_cf28() T0 {
	return _this.Get_().(D7_DC21).Cf28
}

func (_this D7) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D7_DC19).Cf25
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf26 == data2.Cf26
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf27 == data2.Cf27 && _dafny.AreEqual(data1.Cf28, data2.Cf28)
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
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

type D8_DC23 struct {
	Cf30 bool
	Cf31 _dafny.CodePoint
	Cf32 bool
	Cf33 _dafny.Sequence
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf30 bool, Cf31 _dafny.CodePoint, Cf32 bool, Cf33 _dafny.Sequence) D8 {
	return D8{D8_DC23{Cf30, Cf31, Cf32, Cf33}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf34 bool
	Cf35 _dafny.Int
	Cf36 _dafny.Int
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf34 bool, Cf35 _dafny.Int, Cf36 _dafny.Int) D8 {
	return D8{D8_DC24{Cf34, Cf35, Cf36}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC22 struct {
	Cf29 _dafny.Map
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf29 _dafny.Map) D8 {
	return D8{D8_DC22{Cf29}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC25 struct {
	Cf37 D8
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf37 D8) D8 {
	return D8{D8_DC25{Cf37}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(false, _dafny.CodePoint('D'), false, _dafny.EmptySeq)
}

func (_this D8) Dtor_cf30() bool {
	return _this.Get_().(D8_DC23).Cf30
}

func (_this D8) Dtor_cf31() _dafny.CodePoint {
	return _this.Get_().(D8_DC23).Cf31
}

func (_this D8) Dtor_cf32() bool {
	return _this.Get_().(D8_DC23).Cf32
}

func (_this D8) Dtor_cf33() _dafny.Sequence {
	return _this.Get_().(D8_DC23).Cf33
}

func (_this D8) Dtor_cf34() bool {
	return _this.Get_().(D8_DC24).Cf34
}

func (_this D8) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf35
}

func (_this D8) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf36
}

func (_this D8) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D8_DC22).Cf29
}

func (_this D8) Dtor_cf37() D8 {
	return _this.Get_().(D8_DC25).Cf37
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf37) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33.Equals(data2.Cf33)
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Cmp(data2.Cf36) == 0
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf37.Equals(data2.Cf37)
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

type D9_DC27 struct {
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_() D9 {
	return D9{D9_DC27{}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC26 struct {
	Cf38 _dafny.MultiSet
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf38 _dafny.MultiSet) D9 {
	return D9{D9_DC26{Cf38}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC27_()
}

func (_this D9) Dtor_cf38() _dafny.MultiSet {
	return _this.Get_().(D9_DC26).Cf38
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC27:
		{
			return "D9.DC27"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC27:
		{
			_, ok := other.Get_().(D9_DC27)
			return ok
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf38.Equals(data2.Cf38)
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

// Definition of datatype D10
type D10 struct {
	Data_D10_
}

func (_this D10) Get_() Data_D10_ {
	return _this.Data_D10_
}

type Data_D10_ interface {
	isD10()
}

type CompanionStruct_D10_ struct {
}

var Companion_D10_ = CompanionStruct_D10_{}

type D10_DC29 struct {
	Cf40 _dafny.Int
	Cf41 _dafny.Int
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf40 _dafny.Int, Cf41 _dafny.Int) D10 {
	return D10{D10_DC29{Cf40, Cf41}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC28 struct {
	Cf39 _dafny.Set
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf39 _dafny.Set) D10 {
	return D10{D10_DC28{Cf39}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(_dafny.Zero, _dafny.Zero)
}

func (_this D10) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf40
}

func (_this D10) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf41
}

func (_this D10) Dtor_cf39() _dafny.Set {
	return _this.Get_().(D10_DC28).Cf39
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Cmp(data2.Cf41) == 0
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf39.Equals(data2.Cf39)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D10) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D10)
	return ok && _this.Equals(typed)
}

func Type_D10_() _dafny.TypeDescriptor {
	return type_D10_{}
}

type type_D10_ struct {
}

func (_this type_D10_) Default() interface{} {
	return Companion_D10_.Default()
}

func (_this type_D10_) String() string {
	return "main.D10"
}
func (_this D10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D10{}

// End of datatype D10

// Definition of trait T0
type T0 interface {
	String() string
	Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int
	F6() _dafny.Int
	F7() bool
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
	F5  _dafny.Set
	_f0 _dafny.Array
	_f1 _dafny.Array
	_f2 bool
	_f3 _dafny.Map
	_f4 _dafny.CodePoint
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F5 = _dafny.EmptySet
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = false
	_this._f3 = _dafny.EmptyMap
	_this._f4 = _dafny.CodePoint('D')
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Array, f2 bool, f3 _dafny.Map, f4 _dafny.CodePoint, f5 _dafny.Set) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Map {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.CodePoint {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f6 _dafny.Int
	_f7 bool
	F8  _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f6 = _dafny.Zero
	_this._f7 = false
	_this.F8 = _dafny.EmptyMap
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

func (_this *C0) F6() _dafny.Int {
	return _this._f6
}
func (_this *C0) F7() bool {
	return _this._f7
}
func (_this *C0) Ctor__(f8 _dafny.Map, f6 _dafny.Int, f7 bool) {
	{
		(_this).F8 = f8
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C0) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F6()), (_this).F6()), (_this).F6())
	}
}
func (_this *C0) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return !(((_this).F6()).Cmp((func() _dafny.Int {
			if true {
				return (_this).F6()
			}
			return (_this).F6()
		})()) > 0)
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	dummy byte
}

func New_C1_() *C1 {
	_this := C1{}

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

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) M0(p0 bool, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _297_v0 _dafny.Sequence
		_ = _297_v0
		_297_v0 = _dafny.UnicodeSeqOfUtf8Bytes("mdcxtl")
		var _298_v1 _dafny.Int
		_ = _298_v1
		_298_v1 = _dafny.IntOfInt64(362)
		var _299_v2 _dafny.Set
		_ = _299_v2
		_299_v2 = _dafny.SetOf(_dafny.CodePoint('t'), p1, (_297_v0).Select((Companion_Default___.SafeIndex(_298_v1, _dafny.IntOfUint32((_297_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
		_299_v2 = ((_299_v2).Difference(_dafny.SetOf(p1))).Union(_299_v2)
		var _300_i0 _dafny.Int
		_ = _300_i0
		_300_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_300_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_300_i0 = (_300_i0).Plus(_dafny.One)
					var _301_v3 bool
					_ = _301_v3
					_301_v3 = false
					_301_v3 = (p2)
					var _302_v4 _dafny.Map
					_ = _302_v4
					_302_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v3, Companion_Default___.Fm2(globalState))
					var _303_v5 D1
					_ = _303_v5
					_303_v5 = Companion_D1_.Create_DC1_(_302_v4)
					var _304_v6 *C0
					_ = _304_v6
					var _nw51 *C0 = New_C0_()
					_ = _nw51
					_nw51.Ctor__((_303_v5).Dtor_cf1(), _298_v1, p2)
					_304_v6 = _nw51
					var _305_v7 _dafny.MultiSet
					_ = _305_v7
					_305_v7 = _dafny.MultiSetOf(p0)
					var _306_v8 _dafny.MultiSet
					_ = _306_v8
					_306_v8 = _dafny.MultiSetOf(_304_v6, _304_v6)
					if (_304_v6).Fm1(_298_v1, !((_304_v6).Fm1(_298_v1, _301_v3, _298_v1, _305_v7, globalState)), Companion_Default___.SafeModuloInt((_306_v8).Cardinality(), (_305_v7).Cardinality()), (_305_v7).Intersection(_305_v7), globalState) {
						var _307_v9 _dafny.Sequence
						_ = _307_v9
						_307_v9 = _dafny.SeqOf((_305_v7).IsProperSubsetOf(_305_v7))
						_307_v9 = _307_v9
						r0 = _298_v1
						var _308_v10 _dafny.Set
						_ = _308_v10
						_308_v10 = _dafny.SetOf(p0)
						var _309_v11 *C0
						_ = _309_v11
						var _nw52 *C0 = New_C0_()
						_ = _nw52
						_nw52.Ctor__(_304_v6.F8, (_308_v10).Cardinality(), p2)
						_309_v11 = _nw52
						var _310_v12 _dafny.Array
						_ = _310_v12
						var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
						_ = _nw53
						_310_v12 = _nw53
						var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_310_v12), 0))
						_ = _index43
						(_310_v12).ArraySet1(Companion_Default___.SafeDivisionInt(_298_v1, _298_v1), (_index43).Int())
						var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_310_v12), 0))
						_ = _index44
						(_310_v12).ArraySet1(_298_v1, (_index44).Int())
						var _311_v13 _dafny.Array
						_ = _311_v13
						var _nwElement0_5 bool = false
						_ = _nwElement0_5
						var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(5))
						_ = _nw54
						(_nw54).ArraySet1(_nwElement0_5, 0)
						(_nw54).ArraySet1(_301_v3, 1)
						(_nw54).ArraySet1(false, 2)
						(_nw54).ArraySet1(_301_v3, 3)
						(_nw54).ArraySet1((_dafny.SetOf(!(_301_v3))).IsDisjointFrom(_308_v10), 4)
						_311_v13 = _nw54
						var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_311_v13), 0))
						_ = _index45
						(_311_v13).ArraySet1((_dafny.IntOfInt64(341)).Cmp(_298_v1) <= 0, (_index45).Int())
						var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_311_v13), 0))
						_ = _index46
						(_311_v13).ArraySet1(p0, (_index46).Int())
					} else {
						r0 = _298_v1
						var _312_v14 _dafny.Array
						_ = _312_v14
						var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
						_ = _nw55
						_312_v14 = _nw55
						var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_312_v14), 0))
						_ = _index47
						(_312_v14).ArraySet1(Companion_Default___.SafeModuloInt(_298_v1, _298_v1), (_index47).Int())
						var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_312_v14), 0))
						_ = _index48
						(_312_v14).ArraySet1(_dafny.IntOfInt64(47), (_index48).Int())
						_298_v1 = (_312_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_312_v14), 0))).Int()).(_dafny.Int)
						var _313_v15 *C0
						_ = _313_v15
						var _nw56 *C0 = New_C0_()
						_ = _nw56
						_nw56.Ctor__(_304_v6.F8, _298_v1, true)
						_313_v15 = _nw56
						var _314_v16 _dafny.Sequence
						_ = _314_v16
						_314_v16 = _dafny.SeqOf((_312_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_312_v14), 0))).Int()).(_dafny.Int))
						var _315_v17 _dafny.MultiSet
						_ = _315_v17
						_315_v17 = _dafny.MultiSetOf(_dafny.IntOfUint32((_314_v16).Cardinality()))
						var _316_v18 _dafny.Map
						_ = _316_v18
						_316_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v17, (_312_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_312_v14), 0))).Int()).(_dafny.Int))
						var _317_v19 _dafny.Sequence
						_ = _317_v19
						_317_v19 = _dafny.SeqOf(_301_v3)
						_316_v18 = (_316_v18).Update(_dafny.MultiSetOf(_298_v1, _dafny.IntOfUint32((_317_v19).Cardinality())), _298_v1)
					}
					var _318_v20 _dafny.MultiSet
					_ = _318_v20
					_318_v20 = _dafny.MultiSetOf(_298_v1)
					var _319_v21 *C0
					_ = _319_v21
					var _nw57 *C0 = New_C0_()
					_ = _nw57
					_nw57.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _318_v20)).Merge(_302_v4), Companion_Default___.SafeModuloInt(_298_v1, _298_v1), p2)
					_319_v21 = _nw57
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _320_v22 bool
		_ = _320_v22
		_320_v22 = false
		_320_v22 = !((Companion_Default___.Fm3(p1, _320_v22, _320_v22, globalState)).Cmp(_298_v1) != 0)
		var _321_i1 _dafny.Int
		_ = _321_i1
		_321_i1 = _dafny.Zero
		{
			for p0 {
				{
					if (_321_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_321_i1 = (_321_i1).Plus(_dafny.One)
					var _322_v23 _dafny.Array
					_ = _322_v23
					var _len0_7 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_7
					var _nw58 _dafny.Array
					_ = _nw58
					if _len0_7.Cmp(_dafny.Zero) == 0 {
						_nw58 = _dafny.NewArray(_len0_7)
					} else {
						var _init7 func(_dafny.Int) bool = (func(_323_p0 bool) func(_dafny.Int) bool {
							return func(_324_i2 _dafny.Int) bool {
								return _323_p0
							}
						})(p0)
						_ = _init7
						var _element0_7 = _init7(_dafny.Zero)
						_ = _element0_7
						_nw58 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
						(_nw58).ArraySet1(_element0_7, 0)
						var _nativeLen0_7 = (_len0_7).Int()
						_ = _nativeLen0_7
						for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
							(_nw58).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
						}
					}
					_322_v23 = _nw58
					var _325_v24 _dafny.Set
					_ = _325_v24
					_325_v24 = _dafny.SetOf(_322_v23)
					_325_v24 = _325_v24
					var _326_v25 _dafny.Array
					_ = _326_v25
					var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
					_ = _nw59
					_326_v25 = _nw59
					var _327_v26 _dafny.Map
					_ = _327_v26
					_327_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _326_v25)
					var _328_v27 _dafny.Array
					_ = _328_v27
					var _nwElement0_6 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rymhitw")).Cardinality())
					_ = _nwElement0_6
					var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(21))
					_ = _nw60
					(_nw60).ArraySet1(_nwElement0_6, 0)
					(_nw60).ArraySet1(_298_v1, 1)
					(_nw60).ArraySet1(_dafny.IntOfInt64(687), 2)
					(_nw60).ArraySet1(_298_v1, 3)
					(_nw60).ArraySet1(_dafny.IntOfInt64(77), 4)
					(_nw60).ArraySet1(_dafny.IntOfUint32((_297_v0).Cardinality()), 5)
					(_nw60).ArraySet1((_298_v1).Plus(_298_v1), 6)
					(_nw60).ArraySet1((_298_v1).Plus(_298_v1), 7)
					(_nw60).ArraySet1(_298_v1, 8)
					(_nw60).ArraySet1(_dafny.IntOfInt64(899), 9)
					(_nw60).ArraySet1(_dafny.IntOfInt64(397), 10)
					(_nw60).ArraySet1((_298_v1).Plus(_298_v1), 11)
					(_nw60).ArraySet1(_298_v1, 12)
					(_nw60).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-807))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}((func(_329_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_330_i3 _dafny.Int) _dafny.CodePoint {
							return _329_p1
						}
					})(p1))), _297_v0)).Cardinality()), 13)
					(_nw60).ArraySet1(((_327_v26).Update(_320_v22, _326_v25)).Cardinality(), 14)
					(_nw60).ArraySet1(_298_v1, 15)
					(_nw60).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_297_v0).Cardinality()), _298_v1)), 16)
					(_nw60).ArraySet1(_dafny.IntOfInt64(140), 17)
					(_nw60).ArraySet1((_298_v1).Plus(_298_v1), 18)
					(_nw60).ArraySet1((_298_v1).Minus(_298_v1), 19)
					(_nw60).ArraySet1((_dafny.Zero).Minus(_298_v1), 20)
					_328_v27 = _nw60
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))
					_ = _index49
					(_326_v25).ArraySet1((_298_v1).Times(_dafny.IntOfInt64(56)), (_index49).Int())
					var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))
					_ = _index50
					(_326_v25).ArraySet1((_298_v1).Times((_dafny.IntOfInt64(-528)).Plus(_298_v1)), (_index50).Int())
					if (_298_v1).Cmp(((_dafny.SetOf((_326_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))).Int()).(_dafny.Int), _298_v1)).Cardinality()).Times(((Companion_D2_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v22, _298_v1))).Dtor_cf5()).Cardinality())) < 0 {
						_320_v22 = p0
						var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))
						_ = _index51
						(_326_v25).ArraySet1((_326_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))).Int()).(_dafny.Int), (_index51).Int())
						var _331_v28 D1
						_ = _331_v28
						_331_v28 = Companion_D1_.Create_DC2_(p1, _dafny.IntOfInt64(-281), _298_v1)
						var _332_v29 _dafny.MultiSet
						_ = _332_v29
						_332_v29 = _dafny.MultiSetOf(_331_v28, _331_v28, _331_v28, _331_v28)
						_320_v22 = (_332_v29).IsSubsetOf(_dafny.MultiSetOf(_331_v28))
						var _333_v30 _dafny.MultiSet
						_ = _333_v30
						_333_v30 = _dafny.MultiSetOf((_326_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))).Int()).(_dafny.Int))
						var _334_v31 _dafny.Sequence
						_ = _334_v31
						_334_v31 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v22, _333_v30))
						var _335_v32 _dafny.Sequence
						_ = _335_v32
						_335_v32 = _dafny.SeqOf(_320_v22)
						var _336_v33 _dafny.Sequence
						_ = _336_v33
						_336_v33 = _dafny.SeqOf(_335_v32)
						var _337_v34 *C0
						_ = _337_v34
						var _nw61 *C0 = New_C0_()
						_ = _nw61
						_nw61.Ctor__((_334_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_336_v33).Cardinality()), _dafny.IntOfUint32((_334_v31).Cardinality()))).Uint32()).(_dafny.Map), (_326_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))).Int()).(_dafny.Int), p2)
						_337_v34 = _nw61
						var _338_v35 *C0
						_ = _338_v35
						var _nw62 *C0 = New_C0_()
						_ = _nw62
						_nw62.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _333_v30), (_326_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))).Int()).(_dafny.Int), p2)
						_338_v35 = _nw62
					} else {
						var _339_v36 _dafny.Map
						_ = _339_v36
						_339_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_298_v1)), (_dafny.Zero).Minus(_dafny.IntOfInt64(-582)))
						var _340_v37 _dafny.Map
						_ = _340_v37
						_340_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (Companion_Default___.Fm4(_339_v36, p1, p0, _320_v22, globalState)).Cardinality())
						var _341_v38 D2
						_ = _341_v38
						_341_v38 = Companion_D2_.Create_DC3_(_340_v37)
						_341_v38 = _341_v38
						var _342_v39 _dafny.Sequence
						_ = _342_v39
						_342_v39 = _dafny.SeqOf(_320_v22, p0, _320_v22, _320_v22, p0)
						_340_v37 = (_340_v37).Update(p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_342_v39).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_297_v0).Cardinality()))))))
						_320_v22 = _320_v22
						var _343_v40 _dafny.Set
						_ = _343_v40
						_343_v40 = _dafny.SetOf(p0, p0)
						var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_322_v23), 0))
						_ = _index52
						(_322_v23).ArraySet1((_343_v40).IsProperSubsetOf(_343_v40), (_index52).Int())
						var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_322_v23), 0))
						_ = _index53
						(_322_v23).ArraySet1((_320_v22) || (p0), (_index53).Int())
						var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))
						_ = _index54
						(_326_v25).ArraySet1((_326_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_326_v25), 0))).Int()).(_dafny.Int), (_index54).Int())
					}
					r0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_298_v1))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		_298_v1 = (Companion_D2_.Create_DC4_(p0, _298_v1)).Dtor_cf7()
		var _344_v41 _dafny.Set
		_ = _344_v41
		_344_v41 = _dafny.SetOf(_dafny.IntOfInt64(980))
		var _345_v42 _dafny.Sequence
		_ = _345_v42
		_345_v42 = _dafny.SeqOf(_298_v1, _298_v1)
		var _346_v43 _dafny.Array
		_ = _346_v43
		var _nwElement0_7 _dafny.Int = _298_v1
		_ = _nwElement0_7
		var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(23))
		_ = _nw63
		(_nw63).ArraySet1(_nwElement0_7, 0)
		(_nw63).ArraySet1(_298_v1, 1)
		(_nw63).ArraySet1(_298_v1, 2)
		(_nw63).ArraySet1(_298_v1, 3)
		(_nw63).ArraySet1(_298_v1, 4)
		(_nw63).ArraySet1(_dafny.IntOfInt64(107), 5)
		(_nw63).ArraySet1(_298_v1, 6)
		(_nw63).ArraySet1(_298_v1, 7)
		(_nw63).ArraySet1(_298_v1, 8)
		(_nw63).ArraySet1(_298_v1, 9)
		(_nw63).ArraySet1(_298_v1, 10)
		(_nw63).ArraySet1(_298_v1, 11)
		(_nw63).ArraySet1((_344_v41).Cardinality(), 12)
		(_nw63).ArraySet1((_345_v42).Select((Companion_Default___.SafeIndex(_298_v1, _dafny.IntOfUint32((_345_v42).Cardinality()))).Uint32()).(_dafny.Int), 13)
		(_nw63).ArraySet1(_298_v1, 14)
		(_nw63).ArraySet1(_dafny.IntOfInt64(696), 15)
		(_nw63).ArraySet1(_298_v1, 16)
		(_nw63).ArraySet1(_298_v1, 17)
		(_nw63).ArraySet1(_298_v1, 18)
		(_nw63).ArraySet1(_298_v1, 19)
		(_nw63).ArraySet1(_298_v1, 20)
		(_nw63).ArraySet1(_298_v1, 21)
		(_nw63).ArraySet1(_298_v1, 22)
		_346_v43 = _nw63
		var _347_v44 _dafny.MultiSet
		_ = _347_v44
		_347_v44 = _dafny.MultiSetOf(_346_v43)
		var _rhs38 _dafny.Sequence = _297_v0
		_ = _rhs38
		var _rhs39 _dafny.MultiSet = ((Companion_D3_.Create_DC6_(_347_v44)).Dtor_cf13()).Intersection(_347_v44)
		_ = _rhs39
		var _rhs40 bool = !(_320_v22)
		_ = _rhs40
		_297_v0 = _rhs38
		_347_v44 = _rhs39
		_320_v22 = _rhs40
		r0 = Companion_Default___.SafeDivisionInt(_298_v1, (_298_v1).Minus(_298_v1))
		return r0
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
