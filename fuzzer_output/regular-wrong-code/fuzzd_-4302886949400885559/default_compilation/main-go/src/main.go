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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('y')
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.Int {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()), (_dafny.MultiSetOf(!(!(!(false))))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-825))), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wirhrrwqa")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-305)), _dafny.IntOfInt64(846), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))).Intersection(_dafny.MultiSetOf((_dafny.MultiSetOf(!(true))).Cardinality())))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	return (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-177)), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, true, true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(288), !(false))).Cardinality()))).IsSubsetOf(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(423)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kxsyhfc")).Cardinality())), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.CodePoint('y'), _dafny.CodePoint('b'), _dafny.CodePoint('e'))
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(800))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(186))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(331))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(206)
		}))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(331))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_0_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(206)
			})), _1_v0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_1_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return (Companion_D5_.Create_DC5_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('d'))))).Dtor_cf5()
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(934))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() bool {
		if true {
			return true
		}
		return !(!(false))
	})(), true, false, !(((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality(), _dafny.IntOfInt64(274), _dafny.IntOfInt64(377))).Cardinality()).Cmp(_dafny.IntOfInt64(-936)) != 0), !(false) || (false))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	var r0 _dafny.Map = _dafny.EmptyMap
	_ = r0
	var _2_v0 _dafny.Int
	_ = _2_v0
	_2_v0 = _dafny.IntOfInt64(896)
	(globalState).F12 = (_2_v0).Cmp(_2_v0) >= 0
	var _3_v1 _dafny.Set
	_ = _3_v1
	_3_v1 = _dafny.SetOf(p2)
	var _4_v2 _dafny.Map
	_ = _4_v2
	_4_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2_v0)
	var _5_v3 _dafny.Map
	_ = _5_v3
	_5_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(p1, _4_v2, globalState), p2)
	if ((_3_v1).Intersection(_dafny.SetOf((func() bool {
		if (_5_v3).Contains((_dafny.Zero).Minus(_2_v0)) {
			return (_5_v3).Get((_dafny.Zero).Minus(_2_v0)).(bool)
		}
		return p2
	})()))).IsSubsetOf(_3_v1) {
		var _6_v4 _dafny.Sequence
		_ = _6_v4
		_6_v4 = _dafny.SeqOf(p0)
		_6_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _6_v4), _6_v4)
		var _7_v5 _dafny.Array
		_ = _7_v5
		var _nwElement0_0 bool = p2
		_ = _nwElement0_0
		var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(12))
		_ = _nw0
		(_nw0).ArraySet1(_nwElement0_0, 0)
		(_nw0).ArraySet1(p2, 1)
		(_nw0).ArraySet1(p1, 2)
		(_nw0).ArraySet1(p2, 3)
		(_nw0).ArraySet1(p2, 4)
		(_nw0).ArraySet1(!(!(p0)), 5)
		(_nw0).ArraySet1(!(p2), 6)
		(_nw0).ArraySet1(p1, 7)
		(_nw0).ArraySet1(p1, 8)
		(_nw0).ArraySet1(p0, 9)
		(_nw0).ArraySet1(!(p0), 10)
		(_nw0).ArraySet1(p1, 11)
		_7_v5 = _nw0
		var _8_v6 *C0
		_ = _8_v6
		var _nw1 *C0 = New_C0_()
		_ = _nw1
		_nw1.Ctor__()
		_8_v6 = _nw1
		var _9_v7 _dafny.Map
		_ = _9_v7
		_9_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v5, _8_v6)
		_9_v7 = (_9_v7).Update(_7_v5, _8_v6)
		_6_v4 = _dafny.Companion_Sequence_.Concatenate(_6_v4, _6_v4)
		var _10_v8 _dafny.Map
		_ = _10_v8
		_10_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		var _source0 _dafny.Map = _10_v8
		_ = _source0
		var _11___mcc_h0 _dafny.Map = _source0
		_ = _11___mcc_h0
		var _12_cf1 _dafny.Map = _11___mcc_h0
		_ = _12_cf1
		var _13_v9 *C0
		_ = _13_v9
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__()
		_13_v9 = _nw2
		var _14_v10 _dafny.Map
		_ = _14_v10
		_14_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v6, p2)
		var _15_v11 _dafny.MultiSet
		_ = _15_v11
		_15_v11 = _dafny.MultiSetOf((_4_v2).Merge(Companion_Default___.Fm6((func() bool {
			if (_14_v10).Contains(_8_v6) {
				return (_14_v10).Get(_8_v6).(bool)
			}
			return p1
		})(), p0, globalState)))
		(globalState).F5 = (func() _dafny.Int {
			if (_15_v11).Contains(_4_v2) {
				return (_15_v11).Multiplicity(_4_v2)
			}
			return _2_v0
		})()
		var _16_v12 _dafny.Array
		_ = _16_v12
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_0
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = func(_17_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("mm")
			}
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw3 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw3).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw3).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_16_v12 = _nw3
		var _18_v13 _dafny.Sequence
		_ = _18_v13
		_18_v13 = _dafny.UnicodeSeqOfUtf8Bytes("igqyhvkw")
		var _19_v14 _dafny.Sequence
		_ = _19_v14
		_19_v14 = _18_v13
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_16_v12), 0))
		_ = _index0
		(_16_v12).ArraySet1(_19_v14, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_16_v12), 0))
		_ = _index1
		(_16_v12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_18_v13, _18_v13), (_index1).Int())
		var _20_v15 _dafny.CodePoint
		_ = _20_v15
		_20_v15 = _dafny.CodePoint('j')
		var _21_v16 _dafny.Int
		_ = _21_v16
		_21_v16 = _2_v0
		var _22_v17 _dafny.MultiSet
		_ = _22_v17
		_22_v17 = _dafny.MultiSetOf(_13_v9)
		var _23_v18 _dafny.Array
		_ = _23_v18
		var _nwElement0_1 _dafny.Int = (_22_v17).Cardinality()
		_ = _nwElement0_1
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(10))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_1, 0)
		(_nw4).ArraySet1(_2_v0, 1)
		(_nw4).ArraySet1((_4_v2).Cardinality(), 2)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_18_v13).Cardinality()), 3)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_2_v0)).Cardinality()), 4)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_6_v4).Cardinality()), 5)
		(_nw4).ArraySet1(_2_v0, 6)
		(_nw4).ArraySet1(_2_v0, 7)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_18_v13).Cardinality()), 8)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_18_v13, (Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_18_v13).Cardinality()))).Uint32(), _20_v15), (Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_18_v13, (Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_18_v13).Cardinality()))).Uint32(), _20_v15)).Cardinality()))).Uint32(), _20_v15)).Cardinality()), 9)
		_23_v18 = _nw4
		var _24_v19 _dafny.MultiSet
		_ = _24_v19
		_24_v19 = _dafny.MultiSetOf((_21_v16), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _23_v18)).Cardinality(), (_2_v0).Times(Companion_Default___.Fm1(p0, _4_v2, globalState)))
		var _rhs0 bool = (Companion_Default___.Fm1(p0, _4_v2, globalState)).Cmp(_2_v0) > 0
		_ = _rhs0
		var _rhs1 _dafny.CodePoint = _20_v15
		_ = _rhs1
		var _rhs2 _dafny.MultiSet = _24_v19
		_ = _rhs2
		var _rhs3 bool = p2
		_ = _rhs3
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		_lhs0.F7 = _rhs0
		_20_v15 = _rhs1
		_24_v19 = _rhs2
		_lhs1.F7 = _rhs3
		if !((_dafny.IntOfUint32((_6_v4).Cardinality())).Cmp(_2_v0) > 0) {
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_7_v5), 0))
			_ = _index2
			(_7_v5).ArraySet1(p1, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_7_v5), 0))
			_ = _index3
			(_7_v5).ArraySet1(true, (_index3).Int())
			(globalState).F7 = (_2_v0).Cmp(_2_v0) >= 0
			(globalState).F12 = p0
			var _25_v20 _dafny.Array
			_ = _25_v20
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_1
			var _nw5 _dafny.Array
			_ = _nw5
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw5 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_26_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_27_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_27_i1, _26_v0)
					}
				})(_2_v0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw5 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw5).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw5).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_25_v20 = _nw5
			var _28_v21 _dafny.Sequence
			_ = _28_v21
			_28_v21 = _dafny.SeqOf(_25_v20, _25_v20)
			_28_v21 = _28_v21
			(globalState).F3 = _2_v0
		} else {
			(globalState).F12 = (_2_v0).Cmp((_2_v0).Minus((_dafny.Zero).Minus(_2_v0))) < 0
			var _29_v22 _dafny.Sequence
			_ = _29_v22
			_29_v22 = _dafny.UnicodeSeqOfUtf8Bytes("aov")
			var _30_v23 _dafny.Map
			_ = _30_v23
			_30_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v6, _29_v22)
			var _31_v24 _dafny.Int
			_ = _31_v24
			_31_v24 = _dafny.IntOfInt64(-98)
			var _32_v25 _dafny.Map
			_ = _32_v25
			_32_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_31_v24), _29_v22)
			var _33_v26 _dafny.MultiSet
			_ = _33_v26
			_33_v26 = _dafny.MultiSetOf(_8_v6)
			var _34_v27 _dafny.Array
			_ = _34_v27
			var _nwElement0_2 _dafny.Int = Companion_Default___.Fm1(p2, _4_v2, globalState)
			_ = _nwElement0_2
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(21))
			_ = _nw6
			(_nw6).ArraySet1(_nwElement0_2, 0)
			(_nw6).ArraySet1(((_30_v23).Merge(_30_v23)).Cardinality(), 1)
			(_nw6).ArraySet1(_2_v0, 2)
			(_nw6).ArraySet1((_dafny.Zero).Minus(_2_v0), 3)
			(_nw6).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm1(p1, _4_v2, globalState)), 4)
			(_nw6).ArraySet1(_2_v0, 5)
			(_nw6).ArraySet1((_31_v24), 6)
			(_nw6).ArraySet1(_2_v0, 7)
			(_nw6).ArraySet1(_2_v0, 8)
			(_nw6).ArraySet1((_32_v25).Cardinality(), 9)
			(_nw6).ArraySet1(_2_v0, 10)
			(_nw6).ArraySet1(((_4_v2).Merge(_4_v2)).Cardinality(), 11)
			(_nw6).ArraySet1(_2_v0, 12)
			(_nw6).ArraySet1(_2_v0, 13)
			(_nw6).ArraySet1(_dafny.IntOfInt64(597), 14)
			(_nw6).ArraySet1(_2_v0, 15)
			(_nw6).ArraySet1(_dafny.IntOfInt64(353), 16)
			(_nw6).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_2_v0, (_dafny.Zero).Minus((_3_v1).Cardinality()))), 17)
			(_nw6).ArraySet1((_2_v0).Minus((_33_v26).Cardinality()), 18)
			(_nw6).ArraySet1(_2_v0, 19)
			(_nw6).ArraySet1(Companion_Default___.Fm1(p0, _4_v2, globalState), 20)
			_34_v27 = _nw6
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_34_v27), 0))
			_ = _index4
			(_34_v27).ArraySet1(_2_v0, (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_34_v27), 0))
			_ = _index5
			(_34_v27).ArraySet1((_2_v0).Times(_2_v0), (_index5).Int())
			var _35_v28 _dafny.CodePoint
			_ = _35_v28
			_35_v28 = _dafny.CodePoint('x')
			_35_v28 = _35_v28
			var _36_v29 *C0
			_ = _36_v29
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__()
			_36_v29 = _nw7
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_7_v5), 0))
			_ = _index6
			(_7_v5).ArraySet1(p2, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_7_v5), 0))
			_ = _index7
			(_7_v5).ArraySet1(p0, (_index7).Int())
		}
	} else {
		var _37_v30 *C0
		_ = _37_v30
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__()
		_37_v30 = _nw8
		var _38_v31 _dafny.Sequence
		_ = _38_v31
		_38_v31 = _dafny.UnicodeSeqOfUtf8Bytes("corwclx")
		var _39_v32 _dafny.Map
		_ = _39_v32
		_39_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2_v0, Companion_Default___.Fm8(true, _2_v0, _38_v31, _2_v0, globalState))
		var _40_v33 _dafny.MultiSet
		_ = _40_v33
		_40_v33 = _dafny.MultiSetOf(_2_v0)
		_39_v32 = (_39_v32).Update(_2_v0, _40_v33)
		var _41_v34 _dafny.Map
		_ = _41_v34
		_41_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(296), _5_v3)
		_41_v34 = (_41_v34).Update(Companion_Default___.SafeDivisionInt(_2_v0, _2_v0), _5_v3)
		(globalState).F12 = p2
		var _42_v35 _dafny.Array
		_ = _42_v35
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_2
		var _nw9 _dafny.Array
		_ = _nw9
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw9 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.CodePoint = func(_43_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw9 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw9).ArraySet1CodePoint(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw9).ArraySet1CodePoint(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_42_v35 = _nw9
		var _44_v36 _dafny.CodePoint
		_ = _44_v36
		_44_v36 = _dafny.CodePoint('b')
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_42_v35), 0))
		_ = _index8
		(_42_v35).ArraySet1CodePoint(_44_v36, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_42_v35), 0))
		_ = _index9
		(_42_v35).ArraySet1CodePoint(_44_v36, (_index9).Int())
	}
	var _45_i3 _dafny.Int
	_ = _45_i3
	_45_i3 = _dafny.Zero
	{
		for p0 {
			{
				if (_45_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_45_i3 = (_45_i3).Plus(_dafny.One)
				if p1 {
					(globalState).F3 = _2_v0
					var _46_v37 _dafny.Array
					_ = _46_v37
					var _len0_3 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_3
					var _nw10 _dafny.Array
					_ = _nw10
					if _len0_3.Cmp(_dafny.Zero) == 0 {
						_nw10 = _dafny.NewArray(_len0_3)
					} else {
						var _init3 func(_dafny.Int) _dafny.Sequence = func(_47_i4 _dafny.Int) _dafny.Sequence {
							return _dafny.UnicodeSeqOfUtf8Bytes("woh")
						}
						_ = _init3
						var _element0_3 = _init3(_dafny.Zero)
						_ = _element0_3
						_nw10 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
						(_nw10).ArraySet1(_element0_3, 0)
						var _nativeLen0_3 = (_len0_3).Int()
						_ = _nativeLen0_3
						for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
							(_nw10).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
						}
					}
					_46_v37 = _nw10
					var _48_v38 _dafny.Sequence
					_ = _48_v38
					_48_v38 = Companion_Default___.Fm3(p0, _2_v0, globalState)
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_46_v37), 0))
					_ = _index10
					(_46_v37).ArraySet1(_48_v38, (_index10).Int())
					var _49_v39 _dafny.Sequence
					_ = _49_v39
					_49_v39 = _dafny.UnicodeSeqOfUtf8Bytes("siujnxa")
					var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_46_v37), 0))
					_ = _index11
					(_46_v37).ArraySet1(_49_v39, (_index11).Int())
					(globalState).F7 = p1
					var _50_v40 _dafny.Array
					_ = _50_v40
					var _nwElement0_3 bool = p1
					_ = _nwElement0_3
					var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(13))
					_ = _nw11
					(_nw11).ArraySet1(_nwElement0_3, 0)
					(_nw11).ArraySet1(p1, 1)
					(_nw11).ArraySet1(p2, 2)
					(_nw11).ArraySet1(p1, 3)
					(_nw11).ArraySet1((func() bool {
						if p0 {
							return p2
						}
						return p2
					})(), 4)
					(_nw11).ArraySet1(p2, 5)
					(_nw11).ArraySet1(p1, 6)
					(_nw11).ArraySet1(p0, 7)
					(_nw11).ArraySet1(p0, 8)
					(_nw11).ArraySet1(!(_3_v1).Contains(Companion_Default___.Fm2((_dafny.Zero).Minus(_2_v0), p0, globalState)), 9)
					(_nw11).ArraySet1(p2, 10)
					(_nw11).ArraySet1(p0, 11)
					(_nw11).ArraySet1((p0) && (p1), 12)
					_50_v40 = _nw11
					_50_v40 = _50_v40
					var _51_v41 *C0
					_ = _51_v41
					var _nw12 *C0 = New_C0_()
					_ = _nw12
					_nw12.Ctor__()
					_51_v41 = _nw12
					var _52_v42 _dafny.Map
					_ = _52_v42
					_52_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
					var _53_v43 _dafny.MultiSet
					_ = _53_v43
					_53_v43 = _dafny.MultiSetOf(false, true, (func() bool {
						if (_52_v42).Contains(true) {
							return (_52_v42).Get(true).(bool)
						}
						return p0
					})())
					var _rhs4 bool = p0
					_ = _rhs4
					var _rhs5 _dafny.Int = _2_v0
					_ = _rhs5
					var _rhs6 bool = false
					_ = _rhs6
					var _rhs7 bool = (_53_v43).IsProperSubsetOf(_53_v43)
					_ = _rhs7
					var _rhs8 *C0 = _51_v41
					_ = _rhs8
					var _lhs2 *GlobalState = globalState
					_ = _lhs2
					var _lhs3 *GlobalState = globalState
					_ = _lhs3
					var _lhs4 *GlobalState = globalState
					_ = _lhs4
					var _lhs5 *GlobalState = globalState
					_ = _lhs5
					_lhs2.F12 = _rhs4
					_lhs3.F3 = _rhs5
					_lhs4.F12 = _rhs6
					_lhs5.F12 = _rhs7
					_51_v41 = _rhs8
				} else {
					var _54_v44 *C0
					_ = _54_v44
					var _nw13 *C0 = New_C0_()
					_ = _nw13
					_nw13.Ctor__()
					_54_v44 = _nw13
					var _55_v45 _dafny.Sequence
					_ = _55_v45
					_55_v45 = _dafny.UnicodeSeqOfUtf8Bytes("sb")
					(globalState).F12 = (_2_v0).Cmp(_dafny.IntOfUint32((_55_v45).Cardinality())) > 0
					var _56_v46 _dafny.MultiSet
					_ = _56_v46
					_56_v46 = _dafny.MultiSetOf(_2_v0)
					(globalState).F7 = ((_56_v46).Difference(_56_v46)).IsProperSubsetOf(Companion_Default___.Fm8(p0, _2_v0, _dafny.UnicodeSeqOfUtf8Bytes("eqibxy"), Companion_Default___.Fm1(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _2_v0), globalState), globalState))
					var _57_v47 _dafny.Array
					_ = _57_v47
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(14)
					_ = _len0_4
					var _nw14 _dafny.Array
					_ = _nw14
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw14 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) bool = (func(_58_p1 bool) func(_dafny.Int) bool {
							return func(_59_i5 _dafny.Int) bool {
								return _58_p1
							}
						})(p1)
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw14 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw14).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw14).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_57_v47 = _nw14
					var _60_v48 _dafny.Map
					_ = _60_v48
					_60_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
					var _61_v49 _dafny.Sequence
					_ = _61_v49
					_61_v49 = _dafny.SeqOf(_60_v48, _60_v48)
					var _62_v50 _dafny.Map
					_ = _62_v50
					_62_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v47, (_61_v49).Select((Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_61_v49).Cardinality()))).Uint32()).(_dafny.Map))
					var _63_v51 _dafny.Sequence
					_ = _63_v51
					_63_v51 = _dafny.SeqOf(p2)
					var _64_v52 _dafny.Map
					_ = _64_v52
					_64_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_63_v51).Cardinality()), _60_v48)
					var _rhs9 _dafny.Int = (_dafny.Zero).Minus(_2_v0)
					_ = _rhs9
					var _rhs10 _dafny.Map = (func() _dafny.Map {
						if (_62_v50).Contains(_57_v47) {
							return (_62_v50).Get(_57_v47).(_dafny.Map)
						}
						return (func() _dafny.Map {
							if (_64_v52).Contains(_dafny.IntOfInt64(768)) {
								return (_64_v52).Get(_dafny.IntOfInt64(768)).(_dafny.Map)
							}
							return _60_v48
						})()
					})()
					_ = _rhs10
					var _lhs6 *GlobalState = globalState
					_ = _lhs6
					_lhs6.F5 = _rhs9
					r0 = _rhs10
					var _65_v53 _dafny.Array
					_ = _65_v53
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_5
					var _nw15 _dafny.Array
					_ = _nw15
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw15 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) _dafny.Sequence = (func(_66_v45 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_67_i6 _dafny.Int) _dafny.Sequence {
								return _66_v45
							}
						})(_55_v45)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw15 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw15).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw15).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_65_v53 = _nw15
					_65_v53 = _65_v53
				}
				(globalState).F12 = !(true)
				var _68_v54 _dafny.Array
				_ = _68_v54
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_6
				var _nw16 _dafny.Array
				_ = _nw16
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw16 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) bool = (func(_69_p2 bool) func(_dafny.Int) bool {
						return func(_70_i7 _dafny.Int) bool {
							return _69_p2
						}
					})(p2)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw16 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw16).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw16).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_68_v54 = _nw16
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_7
				var _nw17 _dafny.Array
				_ = _nw17
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw17 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) bool = (func(_71_p2 bool, _72_p0 bool, _73_v3 _dafny.Map, _74_p1 bool) func(_dafny.Int) bool {
						return func(_75_i8 _dafny.Int) bool {
							return (func() bool {
								if _71_p2 {
									return (func() bool {
										if (_73_v3).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_72_p0)).Cardinality())) {
											return (_73_v3).Get(_dafny.IntOfUint32((_dafny.SeqOf(_72_p0)).Cardinality())).(bool)
										}
										return _74_p1
									})()
								}
								return _74_p1
							})()
						}
					})(p2, p0, _5_v3, p1)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw17 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw17).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw17).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_68_v54 = _nw17
				var _76_v55 _dafny.Map
				_ = _76_v55
				_76_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v54, p0)
				(globalState).F5 = (_2_v0).Times(((_76_v55).Merge(_76_v55)).Cardinality())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _77_v56 _dafny.CodePoint
	_ = _77_v56
	_77_v56 = _dafny.CodePoint('w')
	var _78_v57 _dafny.Sequence
	_ = _78_v57
	_78_v57 = _dafny.UnicodeSeqOfUtf8Bytes("rwuaf")
	if !(!((func() bool {
		if p0 {
			return !_dafny.Companion_Sequence_.Contains(_78_v57, _77_v56)
		}
		return p2
	})())) {
		var _79_v58 _dafny.Array
		_ = _79_v58
		var _nw18 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
		_ = _nw18
		_79_v58 = _nw18
		_79_v58 = _79_v58
		var _80_v59 _dafny.Set
		_ = _80_v59
		_80_v59 = _dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_2_v0))).Cardinality())
		var _81_v60 _dafny.Map
		_ = _81_v60
		_81_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _80_v59)
		_81_v60 = (_81_v60).Update(p2, _80_v59)
		var _82_v61 *C0
		_ = _82_v61
		var _nw19 *C0 = New_C0_()
		_ = _nw19
		_nw19.Ctor__()
		_82_v61 = _nw19
		(globalState).F12 = p0
		var _83_v62 _dafny.Sequence
		_ = _83_v62
		_83_v62 = _dafny.SeqOf(_2_v0, _2_v0, _dafny.IntOfInt64(179), (_dafny.Zero).Minus(_2_v0))
		var _84_v63 _dafny.Map
		_ = _84_v63
		_84_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v62, !(!(!(p0)) || (p2)))
		_84_v63 = (_84_v63).Update(_dafny.Companion_Sequence_.Concatenate(_83_v62, _83_v62), p2)
	} else {
		_78_v57 = _78_v57
		var _85_v64 _dafny.MultiSet
		_ = _85_v64
		_85_v64 = _dafny.MultiSetOf(p1)
		var _86_v65 _dafny.Set
		_ = _86_v65
		_86_v65 = _dafny.SetOf(_2_v0)
		var _87_v66 _dafny.MultiSet
		_ = _87_v66
		_87_v66 = _dafny.MultiSetOf(_dafny.IntOfUint32((_78_v57).Cardinality()))
		var _88_v67 _dafny.MultiSet
		_ = _88_v67
		_88_v67 = _dafny.MultiSetOf(_87_v66, _87_v66)
		var _89_v68 _dafny.Map
		_ = _89_v68
		_89_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v67, _4_v2)
		var _90_v69 _dafny.Sequence
		_ = _90_v69
		_90_v69 = _dafny.SeqOf(_2_v0)
		var _91_v70 _dafny.Array
		_ = _91_v70
		var _nwElement0_4 _dafny.Int = _2_v0
		_ = _nwElement0_4
		var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(19))
		_ = _nw20
		(_nw20).ArraySet1(_nwElement0_4, 0)
		(_nw20).ArraySet1(_2_v0, 1)
		(_nw20).ArraySet1(_2_v0, 2)
		(_nw20).ArraySet1(_2_v0, 3)
		(_nw20).ArraySet1(_2_v0, 4)
		(_nw20).ArraySet1(_2_v0, 5)
		(_nw20).ArraySet1(_2_v0, 6)
		(_nw20).ArraySet1(_2_v0, 7)
		(_nw20).ArraySet1((_85_v64).Cardinality(), 8)
		(_nw20).ArraySet1(_dafny.IntOfUint32((_78_v57).Cardinality()), 9)
		(_nw20).ArraySet1(_2_v0, 10)
		(_nw20).ArraySet1(_2_v0, 11)
		(_nw20).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_90_v69).Cardinality())), 12)
		(_nw20).ArraySet1(_2_v0, 13)
		(_nw20).ArraySet1(_2_v0, 14)
		(_nw20).ArraySet1(_2_v0, 15)
		(_nw20).ArraySet1(_2_v0, 16)
		(_nw20).ArraySet1(_2_v0, 17)
		(_nw20).ArraySet1(_2_v0, 18)
		_91_v70 = _nw20
		var _92_v71 _dafny.Sequence
		_ = _92_v71
		_92_v71 = _dafny.SeqOf(_91_v70, _91_v70)
		var _93_v72 _dafny.Sequence
		_ = _93_v72
		_93_v72 = _dafny.SeqOf((_92_v71).Select((Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_92_v71).Cardinality()))).Uint32()).(_dafny.Array))
		var _94_v73 bool
		_ = _94_v73
		_94_v73 = false
		var _95_v74 bool
		_ = _95_v74
		_95_v74 = (_94_v73)
		var _96_v75 *C0
		_ = _96_v75
		var _nw21 *C0 = New_C0_()
		_ = _nw21
		_nw21.Ctor__()
		_96_v75 = _nw21
		var _97_v76 _dafny.Map
		_ = _97_v76
		_97_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v75, _78_v57)
		var _98_v77 _dafny.Array
		_ = _98_v77
		var _nwElement0_5 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality())
		_ = _nwElement0_5
		var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
		_ = _nw22
		(_nw22).ArraySet1(_nwElement0_5, 0)
		(_nw22).ArraySet1((_85_v64).Cardinality(), 1)
		(_nw22).ArraySet1(_2_v0, 2)
		(_nw22).ArraySet1(_2_v0, 3)
		(_nw22).ArraySet1(_dafny.IntOfInt64(-713), 4)
		(_nw22).ArraySet1(Companion_Default___.Fm1(p0, _4_v2, globalState), 5)
		(_nw22).ArraySet1((func() _dafny.Int {
			if p2 {
				return _2_v0
			}
			return _dafny.IntOfInt64(144)
		})(), 6)
		(_nw22).ArraySet1(((_86_v65).Intersection(_86_v65)).Cardinality(), 7)
		(_nw22).ArraySet1(_2_v0, 8)
		(_nw22).ArraySet1(_2_v0, 9)
		(_nw22).ArraySet1(Companion_Default___.Fm1(Companion_Default___.Fm2(_dafny.IntOfInt64(661), p0, globalState), (func() _dafny.Map {
			if (_89_v68).Contains(_88_v67) {
				return (_89_v68).Get(_88_v67).(_dafny.Map)
			}
			return _4_v2
		})(), globalState), 10)
		(_nw22).ArraySet1(_2_v0, 11)
		(_nw22).ArraySet1(_dafny.IntOfUint32((_93_v72).Cardinality()), 12)
		(_nw22).ArraySet1(_2_v0, 13)
		(_nw22).ArraySet1((func() _dafny.Int {
			if _94_v73 {
				return _2_v0
			}
			return _2_v0
		})(), 14)
		(_nw22).ArraySet1(_2_v0, 15)
		(_nw22).ArraySet1(_2_v0, 16)
		(_nw22).ArraySet1(_dafny.IntOfUint32((_90_v69).Cardinality()), 17)
		(_nw22).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_90_v69, _90_v69)).Cardinality()), 18)
		(_nw22).ArraySet1(_2_v0, 19)
		(_nw22).ArraySet1((_97_v76).Cardinality(), 20)
		(_nw22).ArraySet1(_2_v0, 21)
		(_nw22).ArraySet1(_2_v0, 22)
		(_nw22).ArraySet1((_3_v1).Cardinality(), 23)
		(_nw22).ArraySet1(_2_v0, 24)
		_98_v77 = _nw22
		var _rhs11 _dafny.Array = _98_v77
		_ = _rhs11
		var _rhs12 _dafny.CodePoint = _77_v56
		_ = _rhs12
		_91_v70 = _rhs11
		_77_v56 = _rhs12
		var _99_v78 _dafny.Map
		_ = _99_v78
		_99_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _100_v79 _dafny.Map
		_ = _100_v79
		_100_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_2_v0, _2_v0, _2_v0), Companion_Default___.Fm0((func() _dafny.Int {
			if (_85_v64).Contains(p2) {
				return (_85_v64).Multiplicity(p2)
			}
			return _2_v0
		})(), p1, _99_v78, globalState))
		_100_v79 = (_100_v79).Update((_dafny.SetOf(_2_v0, _2_v0)).Intersection(_dafny.SetOf(_2_v0, _2_v0)), (_78_v57).Select((Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_78_v57).Cardinality()))).Uint32()).(_dafny.CodePoint))
		_85_v64 = _85_v64
		var _101_v80 *C0
		_ = _101_v80
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__()
		_101_v80 = _nw23
	}
	var _102_v81 _dafny.Map
	_ = _102_v81
	_102_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
	var _103_v82 _dafny.Map
	_ = _103_v82
	_103_v82 = _102_v81
	var _pat_let_tv0 = p2
	_ = _pat_let_tv0
	var _rhs13 bool = func(_source1 _dafny.Map) bool {
		var _104___mcc_h1 _dafny.Map = _source1
		_ = _104___mcc_h1
		var _105_cf1 _dafny.Map = _104___mcc_h1
		_ = _105_cf1
		return _pat_let_tv0
	}(_103_v82)
	_ = _rhs13
	var _rhs14 bool = p0
	_ = _rhs14
	var _lhs7 *GlobalState = globalState
	_ = _lhs7
	var _lhs8 *GlobalState = globalState
	_ = _lhs8
	_lhs7.F7 = _rhs13
	_lhs8.F12 = _rhs14
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_78_v57, _78_v57)).Cardinality())
	_ = _hi0
	for _106_i9 := (_3_v1).Cardinality(); _106_i9.Cmp(_hi0) < 0; _106_i9 = _106_i9.Plus(_dafny.One) {
		(globalState).F3 = _2_v0
		var _107_v83 _dafny.Sequence
		_ = _107_v83
		_107_v83 = _dafny.SeqOf(p2)
		var _108_v84 bool
		_ = _108_v84
		_108_v84 = p1
		var _109_v85 _dafny.Map
		_ = _109_v85
		_109_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _107_v83)
		var _110_v86 _dafny.Array
		_ = _110_v86
		var _nwElement0_6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_107_v83, _dafny.SeqOf(p0, p2))
		_ = _nwElement0_6
		var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(25))
		_ = _nw24
		(_nw24).ArraySet1(_nwElement0_6, 0)
		(_nw24).ArraySet1(_107_v83, 1)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83), 2)
		(_nw24).ArraySet1(_107_v83, 3)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83), 4)
		(_nw24).ArraySet1(_107_v83, 5)
		(_nw24).ArraySet1(_dafny.SeqOf(p1), 6)
		(_nw24).ArraySet1(_107_v83, 7)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v83, _dafny.SeqOf(p0, Companion_Default___.Fm2(_106_i9, p0, globalState))), 8)
		(_nw24).ArraySet1(_107_v83, 9)
		(_nw24).ArraySet1(_107_v83, 10)
		(_nw24).ArraySet1(_107_v83, 11)
		(_nw24).ArraySet1(_107_v83, 12)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Update(_107_v83, (Companion_Default___.SafeIndex(_106_i9, _dafny.IntOfUint32((_107_v83).Cardinality()))).Uint32(), p0), 13)
		(_nw24).ArraySet1(_107_v83, 14)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83), (Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83)).Cardinality()))).Uint32(), !(p0)), (Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83), (Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83)).Cardinality()))).Uint32(), !(p0))).Cardinality()))).Uint32(), p1), 15)
		(_nw24).ArraySet1(Companion_Default___.Fm9((_108_v84), globalState), 16)
		(_nw24).ArraySet1(_107_v83, 17)
		(_nw24).ArraySet1(_107_v83, 18)
		(_nw24).ArraySet1(_dafny.SeqOf(p0, p1, !(!(p0)), p2), 19)
		(_nw24).ArraySet1(_107_v83, 20)
		(_nw24).ArraySet1(_dafny.SeqOf(p1), 21)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83), (Companion_Default___.SafeIndex(_2_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83)).Cardinality()))).Uint32(), p2), 22)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83), 23)
		(_nw24).ArraySet1((func() _dafny.Sequence {
			if (_109_v85).Contains((func() bool {
				if (_5_v3).Contains(_106_i9) {
					return (_5_v3).Get(_106_i9).(bool)
				}
				return p1
			})()) {
				return (_109_v85).Get((func() bool {
					if (_5_v3).Contains(_106_i9) {
						return (_5_v3).Get(_106_i9).(bool)
					}
					return p1
				})()).(_dafny.Sequence)
			}
			return _dafny.SeqOf(p0)
		})(), 24)
		_110_v86 = _nw24
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_110_v86), 0))
		_ = _index12
		(_110_v86).ArraySet1(_107_v83, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_110_v86), 0))
		_ = _index13
		var _rhs15 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_107_v83, _107_v83)
		_ = _rhs15
		var _rhs16 bool = true
		_ = _rhs16
		var _rhs17 bool = p1
		_ = _rhs17
		var _lhs9 _dafny.Array = _110_v86
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_110_v86), 0))
		_ = _lhs10
		var _lhs11 *GlobalState = globalState
		_ = _lhs11
		var _lhs12 *GlobalState = globalState
		_ = _lhs12
		(_lhs9).ArraySet1(_rhs15, (_lhs10).Int())
		_lhs11.F12 = _rhs16
		_lhs12.F7 = _rhs17
		(globalState).F3 = _106_i9
		(globalState).F5 = _106_i9
	}
	r0 = _102_v81
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _111_v0 _dafny.Sequence
	_ = _111_v0
	_111_v0 = _dafny.UnicodeSeqOfUtf8Bytes("olqswci")
	var _112_v1 bool
	_ = _112_v1
	_112_v1 = true
	var _113_v2 _dafny.Int
	_ = _113_v2
	_113_v2 = _dafny.IntOfInt64(-891)
	var _114_globalState *GlobalState
	_ = _114_globalState
	var _nw25 *GlobalState = New_GlobalState_()
	_ = _nw25
	_nw25.Ctor__(false, _dafny.IntOfInt64(115), true, _dafny.IntOfInt64(-895), _111_v0, _dafny.IntOfInt64(403), _dafny.IntOfInt64(751), false, _111_v0, _dafny.IntOfInt64(958), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _113_v2)).Update(_112_v1, _113_v2), true, false, false, _111_v0)
	_114_globalState = _nw25
	if _112_v1 {
		var _115_v3 _dafny.Set
		_ = _115_v3
		_115_v3 = _dafny.SetOf((_112_v1), _112_v1, _112_v1, _112_v1, _112_v1)
		(_114_globalState).F5 = (func() _dafny.Int {
			if true {
				return (_113_v2).Minus((_115_v3).Cardinality())
			}
			return _113_v2
		})()
		var _116_v4 _dafny.MultiSet
		_ = _116_v4
		_116_v4 = _dafny.MultiSetOf(_dafny.IntOfInt64(-807))
		_112_v1 = (_116_v4).IsProperSubsetOf((_116_v4).Update(_113_v2, Companion_Default___.Abs(_dafny.IntOfInt64(837))))
		var _117_v5 _dafny.Array
		_ = _117_v5
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw26
		_117_v5 = _nw26
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_117_v5), 0))
		_ = _index14
		(_117_v5).ArraySet1(_111_v0, (_index14).Int())
		var _118_v6 bool
		_ = _118_v6
		_118_v6 = _112_v1
		var _119_v7 _dafny.Sequence
		_ = _119_v7
		_119_v7 = _dafny.SeqOf(_112_v1, _112_v1)
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_117_v5), 0))
		_ = _index15
		var _rhs18 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_119_v7, _119_v7)
		_ = _rhs18
		var _rhs19 _dafny.Int = _113_v2
		_ = _rhs19
		var _rhs20 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_111_v0, _111_v0)
		_ = _rhs20
		var _rhs21 bool = _118_v6
		_ = _rhs21
		var _rhs22 _dafny.Int = Companion_Default___.SafeModuloInt((_113_v2).Plus(_113_v2), Companion_Default___.SafeDivisionInt(_113_v2, _113_v2))
		_ = _rhs22
		var _lhs13 *GlobalState = _114_globalState
		_ = _lhs13
		var _lhs14 *GlobalState = _114_globalState
		_ = _lhs14
		var _lhs15 _dafny.Array = _117_v5
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_117_v5), 0))
		_ = _lhs16
		var _lhs17 *GlobalState = _114_globalState
		_ = _lhs17
		_lhs13.F12 = _rhs18
		_lhs14.F3 = _rhs19
		(_lhs15).ArraySet1(_rhs20, (_lhs16).Int())
		_118_v6 = _rhs21
		_lhs17.F5 = _rhs22
		var _120_v8 _dafny.Array
		_ = _120_v8
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_8
		var _nw27 _dafny.Array
		_ = _nw27
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw27 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = (func(_121_v1 bool) func(_dafny.Int) bool {
				return func(_122_i0 _dafny.Int) bool {
					return (_dafny.MultiSetOf(_121_v1)).IsSubsetOf(_dafny.MultiSetOf(_121_v1))
				}
			})(_112_v1)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw27 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw27).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw27).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_120_v8 = _nw27
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_120_v8), 0))
		_ = _index16
		(_120_v8).ArraySet1(_112_v1, (_index16).Int())
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_120_v8), 0))
		_ = _index17
		(_120_v8).ArraySet1(_112_v1, (_index17).Int())
		var _123_v9 _dafny.Array
		_ = _123_v9
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw28
		_123_v9 = _nw28
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_123_v9), 0))
		_ = _index18
		(_123_v9).ArraySet1(_113_v2, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_123_v9), 0))
		_ = _index19
		var _rhs23 _dafny.Int = _113_v2
		_ = _rhs23
		var _rhs24 _dafny.Int = _dafny.IntOfInt64(670)
		_ = _rhs24
		var _lhs18 _dafny.Array = _123_v9
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_123_v9), 0))
		_ = _lhs19
		var _lhs20 *GlobalState = _114_globalState
		_ = _lhs20
		(_lhs18).ArraySet1(_rhs23, (_lhs19).Int())
		_lhs20.F3 = _rhs24
	} else {
		var _124_v10 _dafny.Map
		_ = _124_v10
		var _out0 _dafny.Map
		_ = _out0
		_out0 = Companion_Default___.M0(true, _112_v1, _112_v1, _114_globalState)
		_124_v10 = _out0
		var _125_v11 _dafny.Array
		_ = _125_v11
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_9
		var _nw29 _dafny.Array
		_ = _nw29
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw29 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) bool = (func(_126_v1 bool) func(_dafny.Int) bool {
				return func(_127_i1 _dafny.Int) bool {
					return _126_v1
				}
			})(_112_v1)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw29 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw29).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw29).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_125_v11 = _nw29
		var _128_v12 bool
		_ = _128_v12
		_128_v12 = (_112_v1)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.ArrayLen((_125_v11), 0))
		_ = _index20
		(_125_v11).ArraySet1(_128_v12, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.ArrayLen((_125_v11), 0))
		_ = _index21
		(_125_v11).ArraySet1((func() bool {
			if _112_v1 {
				return _128_v12
			}
			return _128_v12
		})(), (_index21).Int())
		var _129_v13 _dafny.CodePoint
		_ = _129_v13
		_129_v13 = _dafny.CodePoint('e')
		var _130_v14 _dafny.Map
		_ = _130_v14
		_130_v14 = _124_v10
		_129_v13 = Companion_Default___.Fm0(_113_v2, _112_v1, (_130_v14), _114_globalState)
		if true {
			var _131_v15 _dafny.Sequence
			_ = _131_v15
			_131_v15 = _dafny.SeqOf(true)
			var _132_v16 _dafny.Set
			_ = _132_v16
			_132_v16 = _dafny.SetOf(_112_v1)
			var _133_v17 _dafny.Set
			_ = _133_v17
			_133_v17 = _dafny.SetOf(Companion_Default___.Fm0((_132_v16).Cardinality(), _112_v1, _124_v10, _114_globalState), _129_v13)
			var _134_v18 _dafny.Map
			_ = _134_v18
			_134_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_111_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.IntOfUint32((_111_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), _113_v2)
			var _135_v20 _dafny.Map
			_ = _135_v20
			_135_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_112_v1), _113_v2)
			var _136_v21 _dafny.Map
			_ = _136_v21
			_136_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _112_v1), _113_v2), _113_v2)
			var _137_v22 _dafny.Map
			_ = _137_v22
			_137_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v10, _dafny.IntOfUint32((_131_v15).Cardinality()))
			var _138_v23 _dafny.Sequence
			_ = _138_v23
			_138_v23 = _dafny.SeqOf(_113_v2)
			var _139_v24 _dafny.Array
			_ = _139_v24
			var _nwElement0_7 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_131_v15, _131_v15)).Cardinality())
			_ = _nwElement0_7
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(26))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_7, 0)
			(_nw30).ArraySet1(_113_v2, 1)
			(_nw30).ArraySet1(Companion_Default___.SafeModuloInt(_113_v2, _113_v2), 2)
			(_nw30).ArraySet1(_113_v2, 3)
			(_nw30).ArraySet1(_dafny.IntOfInt64(737), 4)
			(_nw30).ArraySet1(((func() _dafny.Set {
				if _112_v1 {
					return _133_v17
				}
				return func() _dafny.Set {
					var _coll1 = _dafny.NewBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate((_134_v18).Keys().Elements()); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _140_v19 _dafny.CodePoint
						_140_v19 = interface{}(_compr_1).(_dafny.CodePoint)
						if (_134_v18).Contains(_140_v19) {
							_coll1.Add(_140_v19)
						}
					}
					return _coll1.ToSet()
				}()
			})()).Cardinality(), 5)
			(_nw30).ArraySet1(_113_v2, 6)
			(_nw30).ArraySet1(_113_v2, 7)
			(_nw30).ArraySet1(_113_v2, 8)
			(_nw30).ArraySet1(_113_v2, 9)
			(_nw30).ArraySet1(Companion_Default___.Fm1(!(Companion_Default___.Fm2(_113_v2, _112_v1, _114_globalState)), _135_v20, _114_globalState), 10)
			(_nw30).ArraySet1(_113_v2, 11)
			(_nw30).ArraySet1(_113_v2, 12)
			(_nw30).ArraySet1(_113_v2, 13)
			(_nw30).ArraySet1((_113_v2).Minus(_113_v2), 14)
			(_nw30).ArraySet1(_113_v2, 15)
			(_nw30).ArraySet1(Companion_Default___.Fm1(_112_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _113_v2), _114_globalState), 16)
			(_nw30).ArraySet1((func() _dafny.Int {
				if (_136_v21).Contains(_137_v22) {
					return (_136_v21).Get(_137_v22).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(847))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_141_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				}))).Cardinality())
			})(), 17)
			(_nw30).ArraySet1((_dafny.Zero).Minus(_113_v2), 18)
			(_nw30).ArraySet1(_dafny.IntOfInt64(-851), 19)
			(_nw30).ArraySet1(_113_v2, 20)
			(_nw30).ArraySet1(_113_v2, 21)
			(_nw30).ArraySet1((_113_v2).Times(_113_v2), 22)
			(_nw30).ArraySet1(_113_v2, 23)
			(_nw30).ArraySet1((func() _dafny.Int {
				if _112_v1 {
					return (_138_v23).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_112_v1, _135_v20, _114_globalState), _dafny.IntOfUint32((_138_v23).Cardinality()))).Uint32()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_113_v2)
			})(), 24)
			(_nw30).ArraySet1((_113_v2).Plus(_113_v2), 25)
			_139_v24 = _nw30
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))
			_ = _index22
			(_139_v24).ArraySet1(_113_v2, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))
			_ = _index23
			(_139_v24).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(115))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}((func(_142_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_143_i3 _dafny.Int) _dafny.CodePoint {
					return _142_v13
				}
			})(_129_v13)))).Cardinality()), (_index23).Int())
			var _144_v25 _dafny.Array
			_ = _144_v25
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw31
			_144_v25 = _nw31
			var _145_v26 _dafny.Map
			_ = _145_v26
			_145_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_139_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))).Int()).(_dafny.Int), !(true))
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_144_v25), 0))
			_ = _index24
			(_144_v25).ArraySet1((func() bool {
				if (_145_v26).Contains(_113_v2) {
					return (_145_v26).Get(_113_v2).(bool)
				}
				return _112_v1
			})(), (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_144_v25), 0))
			_ = _index25
			(_144_v25).ArraySet1(_112_v1, (_index25).Int())
			var _rhs25 _dafny.Sequence = (func() _dafny.Sequence {
				if (_144_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_144_v25), 0))).Int()).(bool) {
					return Companion_Default___.Fm3(_112_v1, (_139_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))).Int()).(_dafny.Int), _114_globalState)
				}
				return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_111_v0, (Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_111_v0).Cardinality()))).Uint32(), _129_v13), _111_v0)
			})()
			_ = _rhs25
			var _rhs26 _dafny.CodePoint = (_111_v0).Select((Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_111_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs26
			var _rhs27 bool = (_144_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_144_v25), 0))).Int()).(bool)
			_ = _rhs27
			var _rhs28 _dafny.Int = _113_v2
			_ = _rhs28
			var _rhs29 bool = _112_v1
			_ = _rhs29
			var _lhs21 *GlobalState = _114_globalState
			_ = _lhs21
			var _lhs22 *GlobalState = _114_globalState
			_ = _lhs22
			_111_v0 = _rhs25
			_129_v13 = _rhs26
			_lhs21.F12 = _rhs27
			_lhs22.F5 = _rhs28
			_112_v1 = _rhs29
			var _146_v27 _dafny.Map
			_ = _146_v27
			_146_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_144_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_144_v25), 0))).Int()).(bool)) == (Companion_Default___.Fm2(_dafny.IntOfUint32((_111_v0).Cardinality()), _112_v1, _114_globalState)), _111_v0)
			_146_v27 = _146_v27
			var _147_v28 _dafny.Sequence
			_ = _147_v28
			_147_v28 = _dafny.SeqOf(_dafny.SeqOf((_139_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))).Int()).(_dafny.Int)), _138_v23, _dafny.SeqOf((_139_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.Update(_138_v23, (Companion_Default___.SafeIndex((_139_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_138_v23).Cardinality()))).Uint32(), _113_v2), _138_v23)
			var _148_v29 _dafny.MultiSet
			_ = _148_v29
			_148_v29 = _dafny.MultiSetOf((_147_v28).Select((Companion_Default___.SafeIndex((_139_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_139_v24), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_147_v28).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _149_v30 _dafny.Map
			_ = _149_v30
			var _out1 _dafny.Map
			_ = _out1
			_out1 = Companion_Default___.M0((_144_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_144_v25), 0))).Int()).(bool), (_144_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_144_v25), 0))).Int()).(bool), !((_148_v29).Equals(_148_v29)), _114_globalState)
			_149_v30 = _out1
		} else {
			var _150_v31 *C0
			_ = _150_v31
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__()
			_150_v31 = _nw32
			_130_v14 = _130_v14
			var _151_v32 _dafny.Array
			_ = _151_v32
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_10
			var _nw33 _dafny.Array
			_ = _nw33
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw33 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Int = func(_152_i4 _dafny.Int) _dafny.Int {
					return (_152_i4).Times(_dafny.IntOfInt64(-925))
				}
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw33 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw33).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw33).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_151_v32 = _nw33
			var _153_v33 _dafny.Sequence
			_ = _153_v33
			_153_v33 = _dafny.SeqOf(_151_v32, _151_v32, _151_v32)
			var _154_v34 _dafny.Map
			_ = _154_v34
			_154_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v10, (func() _dafny.Array {
				if _112_v1 {
					return (_153_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_111_v0, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _113_v2)).Cardinality(), _dafny.IntOfUint32((_111_v0).Cardinality()))).Uint32(), _129_v13)).Cardinality()), _dafny.IntOfUint32((_153_v33).Cardinality()))).Uint32()).(_dafny.Array)
				}
				return _151_v32
			})())
			_154_v34 = ((_154_v34).Merge(_154_v34)).Merge(_154_v34)
			var _155_v35 _dafny.MultiSet
			_ = _155_v35
			_155_v35 = _dafny.MultiSetOf(_112_v1)
			var _156_v36 _dafny.Map
			_ = _156_v36
			_156_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v2, Companion_Default___.SafeModuloInt(_113_v2, (func() _dafny.Int {
				if (_155_v35).Contains(false) {
					return (_155_v35).Multiplicity(false)
				}
				return _113_v2
			})()))
			var _157_v37 _dafny.Map
			_ = _157_v37
			_157_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _113_v2)
			_156_v36 = (_156_v36).Update(Companion_Default___.Fm1(_112_v1, _157_v37, _114_globalState), _113_v2)
			var _158_v38 _dafny.Array
			_ = _158_v38
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_11
			var _nw34 _dafny.Array
			_ = _nw34
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw34 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Sequence = (func(_159_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_160_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_159_v0, _159_v0)
					}
				})(_111_v0)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw34 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw34).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw34).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_158_v38 = _nw34
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_158_v38), 0))
			_ = _index26
			(_158_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_111_v0, _dafny.UnicodeSeqOfUtf8Bytes("hydejw")), (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_158_v38), 0))
			_ = _index27
			(_158_v38).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(255))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_161_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})), (_index27).Int())
		}
		(_114_globalState).F7 = _112_v1
	}
	var _162_i7 _dafny.Int
	_ = _162_i7
	_162_i7 = _dafny.Zero
	{
		for true {
			{
				if (_162_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_162_i7 = (_162_i7).Plus(_dafny.One)
				var _163_v39 _dafny.Set
				_ = _163_v39
				_163_v39 = _dafny.SetOf(_112_v1, false, _112_v1, _112_v1)
				var _164_v40 _dafny.MultiSet
				_ = _164_v40
				_164_v40 = _dafny.MultiSetOf(_dafny.IntOfUint32((_111_v0).Cardinality()))
				var _165_v41 _dafny.Sequence
				_ = _165_v41
				_165_v41 = _dafny.SeqOf(_113_v2)
				var _166_v42 _dafny.Array
				_ = _166_v42
				var _nwElement0_8 bool = _112_v1
				_ = _nwElement0_8
				var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(16))
				_ = _nw35
				(_nw35).ArraySet1(_nwElement0_8, 0)
				(_nw35).ArraySet1((func() bool {
					if _112_v1 {
						return _112_v1
					}
					return _112_v1
				})(), 1)
				(_nw35).ArraySet1(false, 2)
				(_nw35).ArraySet1(_112_v1, 3)
				(_nw35).ArraySet1(_112_v1, 4)
				(_nw35).ArraySet1(_112_v1, 5)
				(_nw35).ArraySet1(((_dafny.Zero).Minus((_163_v39).Cardinality())).Cmp(_113_v2) > 0, 6)
				(_nw35).ArraySet1((_164_v40).IsDisjointFrom((_dafny.MultiSetOf(_113_v2, _113_v2)).Update(_113_v2, Companion_Default___.Abs(_113_v2))), 7)
				(_nw35).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_165_v41, _165_v41)), 8)
				(_nw35).ArraySet1(_112_v1, 9)
				(_nw35).ArraySet1(_112_v1, 10)
				(_nw35).ArraySet1(_112_v1, 11)
				(_nw35).ArraySet1(_112_v1, 12)
				(_nw35).ArraySet1(_112_v1, 13)
				(_nw35).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _112_v1)).Contains(_112_v1), 14)
				(_nw35).ArraySet1(_112_v1, 15)
				_166_v42 = _nw35
				var _167_v43 _dafny.Sequence
				_ = _167_v43
				_167_v43 = _dafny.SeqOf(_112_v1)
				var _168_v44 _dafny.Array
				_ = _168_v44
				var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw36
				_168_v44 = _nw36
				var _rhs30 _dafny.Array = _166_v42
				_ = _rhs30
				var _rhs31 _dafny.Sequence = _167_v43
				_ = _rhs31
				var _rhs32 _dafny.Array = _168_v44
				_ = _rhs32
				var _rhs33 _dafny.Int = (_113_v2).Minus((_113_v2).Minus(_113_v2))
				_ = _rhs33
				_166_v42 = _rhs30
				_167_v43 = _rhs31
				_168_v44 = _rhs32
				_113_v2 = _rhs33
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_166_v42), 0))
				_ = _index28
				(_166_v42).ArraySet1(_112_v1, (_index28).Int())
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_166_v42), 0))
				_ = _index29
				(_166_v42).ArraySet1((_164_v40).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_111_v0).Cardinality()))), (_index29).Int())
				(_114_globalState).F12 = (_166_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_166_v42), 0))).Int()).(bool)
				var _169_v45 _dafny.Int
				_ = _169_v45
				_169_v45 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_166_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_166_v42), 0))).Int()).(bool), Companion_Default___.Fm2(_113_v2, _112_v1, _114_globalState))).Cardinality()
				var _170_v46 _dafny.Set
				_ = _170_v46
				_170_v46 = _dafny.SetOf((_169_v45), (_dafny.Zero).Minus(_113_v2), _113_v2)
				var _171_v47 _dafny.Map
				_ = _171_v47
				_171_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_113_v2), _113_v2)
				_170_v46 = (_170_v46).Difference((_dafny.SetOf((_171_v47).Cardinality(), _113_v2)).Difference(_170_v46))
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _172_v48 _dafny.Array
	_ = _172_v48
	var _len0_12 _dafny.Int = _dafny.IntOfInt64(11)
	_ = _len0_12
	var _nw37 _dafny.Array
	_ = _nw37
	if _len0_12.Cmp(_dafny.Zero) == 0 {
		_nw37 = _dafny.NewArray(_len0_12)
	} else {
		var _init12 func(_dafny.Int) bool = (func(_173_v1 bool) func(_dafny.Int) bool {
			return func(_174_i8 _dafny.Int) bool {
				return !(_173_v1)
			}
		})(_112_v1)
		_ = _init12
		var _element0_12 = _init12(_dafny.Zero)
		_ = _element0_12
		_nw37 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
		(_nw37).ArraySet1(_element0_12, 0)
		var _nativeLen0_12 = (_len0_12).Int()
		_ = _nativeLen0_12
		for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
			(_nw37).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
		}
	}
	_172_v48 = _nw37
	var _175_v49 _dafny.Map
	_ = _175_v49
	_175_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_v48, _112_v1)
	_175_v49 = (_175_v49).Update(_172_v48, _112_v1)
	var _hi1 _dafny.Int = _113_v2
	_ = _hi1
	for _176_i9 := _113_v2; _176_i9.Cmp(_hi1) < 0; _176_i9 = _176_i9.Plus(_dafny.One) {
		var _177_v50 *C0
		_ = _177_v50
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__()
		_177_v50 = _nw38
		var _178_v51 bool
		_ = _178_v51
		_178_v51 = _112_v1
		var _179_v52 _dafny.Set
		_ = _179_v52
		_179_v52 = _dafny.SetOf(_112_v1)
		var _rhs34 _dafny.Set = _179_v52
		_ = _rhs34
		var _rhs35 bool = _112_v1
		_ = _rhs35
		var _lhs23 *GlobalState = _114_globalState
		_ = _lhs23
		_179_v52 = _rhs34
		_lhs23.F7 = _rhs35
		var _180_v53 _dafny.Map
		_ = _180_v53
		var _out2 _dafny.Map
		_ = _out2
		_out2 = Companion_Default___.M0(false, _112_v1, _112_v1, _114_globalState)
		_180_v53 = _out2
		var _181_v54 _dafny.MultiSet
		_ = _181_v54
		_181_v54 = _dafny.MultiSetOf(_dafny.IntOfInt64(791))
		var _182_v55 _dafny.Sequence
		_ = _182_v55
		_182_v55 = _dafny.SeqOf(_181_v54, _dafny.MultiSetOf(_113_v2, _113_v2), _181_v54)
		var _183_v56 _dafny.Sequence
		_ = _183_v56
		_183_v56 = _dafny.SeqOf(_176_i9, _dafny.IntOfInt64(736))
		var _rhs36 _dafny.Int = (_dafny.IntOfInt64(858)).Minus(_113_v2)
		_ = _rhs36
		var _rhs37 _dafny.Int = _176_i9
		_ = _rhs37
		var _rhs38 bool = _112_v1
		_ = _rhs38
		var _rhs39 bool = _112_v1
		_ = _rhs39
		var _rhs40 _dafny.MultiSet = ((_182_v55).Select((Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_182_v55).Cardinality()))).Uint32()).(_dafny.MultiSet)).Update((func() _dafny.Int {
			if _112_v1 {
				return _113_v2
			}
			return (_183_v56).Select((Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_183_v56).Cardinality()))).Uint32()).(_dafny.Int)
		})(), Companion_Default___.Abs(_113_v2))
		_ = _rhs40
		var _lhs24 *GlobalState = _114_globalState
		_ = _lhs24
		var _lhs25 *GlobalState = _114_globalState
		_ = _lhs25
		_lhs24.F5 = _rhs36
		_lhs25.F5 = _rhs37
		_112_v1 = _rhs38
		_112_v1 = _rhs39
		_181_v54 = _rhs40
	}
	var _184_v57 _dafny.Map
	_ = _184_v57
	_184_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _113_v2)
	var _hi2 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm1(_112_v1, _184_v57, _114_globalState))
	_ = _hi2
	for _185_i10 := _113_v2; _185_i10.Cmp(_hi2) < 0; _185_i10 = _185_i10.Plus(_dafny.One) {
		if _112_v1 {
			var _186_v58 _dafny.Sequence
			_ = _186_v58
			_186_v58 = _dafny.SeqOf(_111_v0)
			var _187_v59 *C0
			_ = _187_v59
			var _nw39 *C0 = New_C0_()
			_ = _nw39
			_nw39.Ctor__()
			_187_v59 = _nw39
			var _188_v60 _dafny.Set
			_ = _188_v60
			_188_v60 = _dafny.SetOf(_187_v59)
			var _189_v61 _dafny.Sequence
			_ = _189_v61
			_189_v61 = _dafny.SeqOf(_188_v60, _188_v60)
			var _190_v62 _dafny.Sequence
			_ = _190_v62
			_190_v62 = _111_v0
			var _191_v63 bool
			_ = _191_v63
			_191_v63 = _112_v1
			var _192_v64 _dafny.CodePoint
			_ = _192_v64
			_192_v64 = _dafny.CodePoint('b')
			var _193_v65 _dafny.Array
			_ = _193_v65
			var _nwElement0_9 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("apenmlek")
			_ = _nwElement0_9
			var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(20))
			_ = _nw40
			(_nw40).ArraySet1(_nwElement0_9, 0)
			(_nw40).ArraySet1(Companion_Default___.Fm3(_112_v1, (_dafny.Zero).Minus(_185_i10), _114_globalState), 1)
			(_nw40).ArraySet1(_111_v0, 2)
			(_nw40).ArraySet1(_111_v0, 3)
			(_nw40).ArraySet1(_111_v0, 4)
			(_nw40).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jxraijmpf"), 5)
			(_nw40).ArraySet1((_186_v58).Select((Companion_Default___.SafeIndex(((_189_v61).Select((Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_189_v61).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _dafny.IntOfUint32((_186_v58).Cardinality()))).Uint32()).(_dafny.Sequence), 6)
			(_nw40).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_190_v62), _111_v0), 7)
			(_nw40).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vqhbcmux"), 8)
			(_nw40).ArraySet1(_111_v0, 9)
			(_nw40).ArraySet1(_dafny.Companion_Sequence_.Update((_190_v62), (Companion_Default___.SafeIndex(_185_i10, _dafny.IntOfUint32((_190_v62).Cardinality()))).Uint32(), _dafny.CodePoint('o')), 10)
			(_nw40).ArraySet1(_111_v0, 11)
			(_nw40).ArraySet1(_111_v0, 12)
			(_nw40).ArraySet1(_111_v0, 13)
			(_nw40).ArraySet1(_111_v0, 14)
			(_nw40).ArraySet1(Companion_Default___.Fm3((_187_v59).Fm4(_112_v1, _114_globalState), _185_i10, _114_globalState), 15)
			(_nw40).ArraySet1(_111_v0, 16)
			(_nw40).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm3(!(_191_v63), (_dafny.Zero).Minus(_113_v2), _114_globalState), (Companion_Default___.SafeIndex(_185_i10, _dafny.IntOfUint32((Companion_Default___.Fm3(!(_191_v63), (_dafny.Zero).Minus(_113_v2), _114_globalState)).Cardinality()))).Uint32(), _192_v64), 17)
			(_nw40).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(247))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_194_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_195_i11 _dafny.Int) _dafny.CodePoint {
					return _194_v64
				}
			})(_192_v64))), 18)
			(_nw40).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_111_v0, _111_v0), 19)
			_193_v65 = _nw40
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_193_v65), 0))
			_ = _index30
			(_193_v65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("evqq"), (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_193_v65), 0))
			_ = _index31
			(_193_v65).ArraySet1(_111_v0, (_index31).Int())
			(_114_globalState).F7 = _112_v1
			var _196_v66 _dafny.Map
			_ = _196_v66
			_196_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v2, _187_v59)
			var _197_v67 _dafny.Sequence
			_ = _197_v67
			_197_v67 = _dafny.SeqOf((func() *C0 {
				if _112_v1 {
					return _187_v59
				}
				return (func() *C0 {
					if (_196_v66).Contains(_113_v2) {
						return (_196_v66).Get(_113_v2).(*C0)
					}
					return _187_v59
				})()
			})(), _187_v59)
			_187_v59 = (_197_v67).Select((Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_197_v67).Cardinality()))).Uint32()).(*C0)
			var _198_v68 *C0
			_ = _198_v68
			var _nw41 *C0 = New_C0_()
			_ = _nw41
			_nw41.Ctor__()
			_198_v68 = _nw41
			var _199_v69 _dafny.Array
			_ = _199_v69
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw42
			_199_v69 = _nw42
			var _200_v70 _dafny.Array
			_ = _200_v70
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw43
			_200_v70 = _nw43
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_199_v69), 0))
			_ = _index32
			(_199_v69).ArraySet1(_200_v70, (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_199_v69), 0))
			_ = _index33
			(_199_v69).ArraySet1(_200_v70, (_index33).Int())
		} else {
			var _201_v71 _dafny.Map
			_ = _201_v71
			var _out3 _dafny.Map
			_ = _out3
			_out3 = Companion_Default___.M0(_112_v1, _112_v1, _112_v1, _114_globalState)
			_201_v71 = _out3
			(_114_globalState).F7 = _112_v1
			(_114_globalState).F7 = _dafny.Companion_Sequence_.IsProperPrefixOf(_111_v0, _111_v0)
			var _202_v72 _dafny.Array
			_ = _202_v72
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw44
			_202_v72 = _nw44
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_202_v72), 0))
			_ = _index34
			(_202_v72).ArraySet1(_dafny.One, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_202_v72), 0))
			_ = _index35
			(_202_v72).ArraySet1((_185_i10).Minus(Companion_Default___.SafeDivisionInt(_185_i10, _113_v2)), (_index35).Int())
			(_114_globalState).F5 = Companion_Default___.Fm1(_dafny.Companion_Sequence_.Equal(_111_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(170))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_203_i12 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			}))), _184_v57, _114_globalState)
		}
		var _204_v73 _dafny.Sequence
		_ = _204_v73
		_204_v73 = _dafny.SeqOf(_185_i10, _113_v2)
		var _205_v74 _dafny.Sequence
		_ = _205_v74
		_205_v74 = _dafny.SeqOf(_204_v73, _204_v73)
		var _206_v75 _dafny.Map
		_ = _206_v75
		var _out4 _dafny.Map
		_ = _out4
		_out4 = Companion_Default___.M0(false, _dafny.Companion_Sequence_.Equal(_205_v74, _dafny.SeqOf(_204_v73, _dafny.SeqOf(_185_i10, _185_i10, _113_v2))), _112_v1, _114_globalState)
		_206_v75 = _out4
		(_114_globalState).F12 = _112_v1
		var _207_v76 _dafny.Array
		_ = _207_v76
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
		_ = _nw45
		_207_v76 = _nw45
		var _208_v77 _dafny.Array
		_ = _208_v77
		var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw46
		_208_v77 = _nw46
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_208_v77), 0))
		_ = _index36
		(_208_v77).ArraySet1((_185_i10).Times(_185_i10), (_index36).Int())
		var _209_v78 _dafny.MultiSet
		_ = _209_v78
		_209_v78 = _dafny.MultiSetOf(_112_v1)
		var _210_v79 _dafny.MultiSet
		_ = _210_v79
		_210_v79 = _dafny.MultiSetOf((func() _dafny.Int {
			if (_209_v78).Contains(_112_v1) {
				return (_209_v78).Multiplicity(_112_v1)
			}
			return _113_v2
		})())
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_208_v77), 0))
		_ = _index37
		(_208_v77).ArraySet1((_210_v79).Cardinality(), (_index37).Int())
		var _211_v80 *C0
		_ = _211_v80
		var _nw47 *C0 = New_C0_()
		_ = _nw47
		_nw47.Ctor__()
		_211_v80 = _nw47
		var _212_v81 _dafny.Map
		_ = _212_v81
		_212_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v80, _dafny.UnicodeSeqOfUtf8Bytes("pqq"))
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_208_v77), 0))
		_ = _index38
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_208_v77), 0))
		_ = _index39
		var _rhs41 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_212_v81).Contains(_211_v80) {
				return (_212_v81).Get(_211_v80).(_dafny.Sequence)
			}
			return _111_v0
		})()).Cardinality()))
		_ = _rhs41
		var _rhs42 _dafny.Array = _207_v76
		_ = _rhs42
		var _rhs43 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_111_v0).Cardinality()))).Plus(_113_v2)
		_ = _rhs43
		var _rhs44 _dafny.Int = _dafny.IntOfUint32((_111_v0).Cardinality())
		_ = _rhs44
		var _lhs26 *GlobalState = _114_globalState
		_ = _lhs26
		var _lhs27 _dafny.Array = _208_v77
		_ = _lhs27
		var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_208_v77), 0))
		_ = _lhs28
		var _lhs29 _dafny.Array = _208_v77
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_208_v77), 0))
		_ = _lhs30
		_lhs26.F5 = _rhs41
		_207_v76 = _rhs42
		(_lhs27).ArraySet1(_rhs43, (_lhs28).Int())
		(_lhs29).ArraySet1(_rhs44, (_lhs30).Int())
	}
	if true {
		var _213_v82 *C0
		_ = _213_v82
		var _nw48 *C0 = New_C0_()
		_ = _nw48
		_nw48.Ctor__()
		_213_v82 = _nw48
		_184_v57 = (_184_v57).Update(_112_v1, Companion_Default___.SafeModuloInt(_113_v2, _113_v2))
		var _214_v83 _dafny.MultiSet
		_ = _214_v83
		_214_v83 = _dafny.MultiSetOf(_172_v48)
		var _215_v84 _dafny.Sequence
		_ = _215_v84
		_215_v84 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uknqeb")).Cardinality()), _113_v2)
		var _216_v85 _dafny.Map
		_ = _216_v85
		_216_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v84, (_214_v83).Update(_172_v48, Companion_Default___.Abs(_113_v2)))
		_214_v83 = (func() _dafny.MultiSet {
			if (_216_v85).Contains(_215_v84) {
				return (_216_v85).Get(_215_v84).(_dafny.MultiSet)
			}
			return (_214_v83).Union(_214_v83)
		})()
		(_114_globalState).F7 = (false) || (_112_v1)
		var _217_v86 _dafny.Map
		_ = _217_v86
		_217_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v2, _113_v2)
		var _218_v87 _dafny.MultiSet
		_ = _218_v87
		_218_v87 = _dafny.MultiSetOf(_112_v1)
		_217_v86 = (_217_v86).Update(_dafny.IntOfInt64(467), ((func() _dafny.MultiSet {
			if !(_112_v1) {
				return _218_v87
			}
			return _218_v87
		})()).Cardinality())
	} else {
		var _219_v88 _dafny.MultiSet
		_ = _219_v88
		_219_v88 = _dafny.MultiSetOf(_112_v1)
		var _220_v89 _dafny.Sequence
		_ = _220_v89
		_220_v89 = _dafny.SeqOf(_113_v2)
		var _221_v90 _dafny.Set
		_ = _221_v90
		_221_v90 = _dafny.SetOf(_172_v48)
		var _222_v91 _dafny.CodePoint
		_ = _222_v91
		_222_v91 = _dafny.CodePoint('d')
		var _223_v92 _dafny.Sequence
		_ = _223_v92
		_223_v92 = _dafny.SeqOf(_dafny.SetOf(_222_v91))
		var _224_v93 _dafny.Map
		_ = _224_v93
		_224_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v92, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-475))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_225_i14 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		})))
		var _226_v94 _dafny.Sequence
		_ = _226_v94
		_226_v94 = _dafny.SeqOf(_112_v1)
		var _227_v95 _dafny.Array
		_ = _227_v95
		var _nwElement0_10 _dafny.Int = _113_v2
		_ = _nwElement0_10
		var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(4))
		_ = _nw49
		(_nw49).ArraySet1(_nwElement0_10, 0)
		(_nw49).ArraySet1(_113_v2, 1)
		(_nw49).ArraySet1(_dafny.IntOfUint32((_226_v94).Cardinality()), 2)
		(_nw49).ArraySet1(_113_v2, 3)
		_227_v95 = _nw49
		var _228_v96 _dafny.Map
		_ = _228_v96
		_228_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v95, _112_v1)
		var _229_v97 _dafny.Array
		_ = _229_v97
		var _nwElement0_11 _dafny.Int = _113_v2
		_ = _nwElement0_11
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(26))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_11, 0)
		(_nw50).ArraySet1(_113_v2, 1)
		(_nw50).ArraySet1(_113_v2, 2)
		(_nw50).ArraySet1(_113_v2, 3)
		(_nw50).ArraySet1(_113_v2, 4)
		(_nw50).ArraySet1(_dafny.IntOfInt64(-621), 5)
		(_nw50).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-603))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_230_i13 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))).Cardinality()), (_219_v88).Cardinality()), 6)
		(_nw50).ArraySet1(Companion_Default___.Fm1(_112_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(742)), _114_globalState), 7)
		(_nw50).ArraySet1(_113_v2, 8)
		(_nw50).ArraySet1(_dafny.IntOfInt64(-310), 9)
		(_nw50).ArraySet1(_113_v2, 10)
		(_nw50).ArraySet1((_dafny.SetOf(_112_v1)).Cardinality(), 11)
		(_nw50).ArraySet1((_184_v57).Cardinality(), 12)
		(_nw50).ArraySet1(_113_v2, 13)
		(_nw50).ArraySet1((func() _dafny.Int {
			if !(_112_v1) {
				return _dafny.IntOfUint32((_220_v89).Cardinality())
			}
			return (_220_v89).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oe")).Cardinality())), _dafny.IntOfUint32((_220_v89).Cardinality()))).Uint32()).(_dafny.Int)
		})(), 14)
		(_nw50).ArraySet1(Companion_Default___.Fm1(_112_v1, _184_v57, _114_globalState), 15)
		(_nw50).ArraySet1(_113_v2, 16)
		(_nw50).ArraySet1(Companion_Default___.Fm1(_112_v1, _184_v57, _114_globalState), 17)
		(_nw50).ArraySet1(Companion_Default___.SafeDivisionInt(_113_v2, _113_v2), 18)
		(_nw50).ArraySet1(_113_v2, 19)
		(_nw50).ArraySet1(Companion_Default___.SafeModuloInt(_113_v2, (_221_v90).Cardinality()), 20)
		(_nw50).ArraySet1(_113_v2, 21)
		(_nw50).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_224_v93).Contains(_223_v92) {
				return (_224_v93).Get(_223_v92).(_dafny.Sequence)
			}
			return _111_v0
		})()).Cardinality()), 22)
		(_nw50).ArraySet1(_113_v2, 23)
		(_nw50).ArraySet1(_113_v2, 24)
		(_nw50).ArraySet1((_228_v96).Cardinality(), 25)
		_229_v97 = _nw50
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_227_v95), 0))
		_ = _index40
		(_227_v95).ArraySet1(_113_v2, (_index40).Int())
		var _231_v98 _dafny.Int
		_ = _231_v98
		_231_v98 = _113_v2
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_227_v95), 0))
		_ = _index41
		(_227_v95).ArraySet1((_231_v98), (_index41).Int())
		var _232_v99 _dafny.Set
		_ = _232_v99
		_232_v99 = _dafny.SetOf((func() _dafny.Int {
			if _112_v1 {
				return (_227_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_227_v95), 0))).Int()).(_dafny.Int)
			}
			return (_227_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_227_v95), 0))).Int()).(_dafny.Int)
		})())
		var _233_v100 bool
		_ = _233_v100
		_233_v100 = _112_v1
		var _234_v101 _dafny.Array
		_ = _234_v101
		var _nwElement0_12 bool = Companion_Default___.Fm2(_113_v2, _112_v1, _114_globalState)
		_ = _nwElement0_12
		var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(7))
		_ = _nw51
		(_nw51).ArraySet1(_nwElement0_12, 0)
		(_nw51).ArraySet1(_233_v100, 1)
		(_nw51).ArraySet1(_233_v100, 2)
		(_nw51).ArraySet1(_233_v100, 3)
		(_nw51).ArraySet1(_112_v1, 4)
		(_nw51).ArraySet1(Companion_Default___.Fm5(false, _dafny.IntOfInt64(549), _112_v1, _113_v2, _114_globalState), 5)
		(_nw51).ArraySet1(_112_v1, 6)
		_234_v101 = _nw51
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_234_v101), 0))
		_ = _index42
		(_234_v101).ArraySet1(_233_v100, (_index42).Int())
		var _235_v102 *C0
		_ = _235_v102
		var _nw52 *C0 = New_C0_()
		_ = _nw52
		_nw52.Ctor__()
		_235_v102 = _nw52
		var _236_v103 _dafny.Sequence
		_ = _236_v103
		_236_v103 = _dafny.SeqOf(_235_v102)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_234_v101), 0))
		_ = _index43
		var _rhs45 _dafny.Set = _232_v99
		_ = _rhs45
		var _rhs46 bool = !(_112_v1) || (_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_235_v102), _236_v103))
		_ = _rhs46
		var _rhs47 bool = !(Companion_Default___.Fm6(_112_v1, _112_v1, _114_globalState)).Equals(_184_v57)
		_ = _rhs47
		var _rhs48 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(_112_v1)), _226_v94), _226_v94)
		_ = _rhs48
		var _lhs31 *GlobalState = _114_globalState
		_ = _lhs31
		var _lhs32 _dafny.Array = _234_v101
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_234_v101), 0))
		_ = _lhs33
		_232_v99 = _rhs45
		_lhs31.F12 = _rhs46
		(_lhs32).ArraySet1(_rhs47, (_lhs33).Int())
		_112_v1 = _rhs48
		var _237_v104 _dafny.Map
		_ = _237_v104
		var _out5 _dafny.Map
		_ = _out5
		_out5 = Companion_Default___.M0(_dafny.Companion_Sequence_.Contains(_220_v89, (_227_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_227_v95), 0))).Int()).(_dafny.Int)), _112_v1, _112_v1, _114_globalState)
		_237_v104 = _out5
		var _238_v105 _dafny.Array
		_ = _238_v105
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
		_ = _nw53
		_238_v105 = _nw53
		var _239_v106 _dafny.Set
		_ = _239_v106
		_239_v106 = _dafny.SetOf(_112_v1)
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_238_v105), 0))
		_ = _index44
		(_238_v105).ArraySet1(_239_v106, (_index44).Int())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_238_v105), 0))
		_ = _index45
		var _rhs49 _dafny.MultiSet = _dafny.MultiSetOf(true)
		_ = _rhs49
		var _rhs50 _dafny.Set = (_239_v106).Difference((func() _dafny.Set {
			if _112_v1 {
				return _239_v106
			}
			return _239_v106
		})())
		_ = _rhs50
		var _rhs51 _dafny.Int = _113_v2
		_ = _rhs51
		var _lhs34 _dafny.Array = _238_v105
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_238_v105), 0))
		_ = _lhs35
		var _lhs36 *GlobalState = _114_globalState
		_ = _lhs36
		_219_v88 = _rhs49
		(_lhs34).ArraySet1(_rhs50, (_lhs35).Int())
		_lhs36.F3 = _rhs51
		var _240_v107 *C0
		_ = _240_v107
		var _nw54 *C0 = New_C0_()
		_ = _nw54
		_nw54.Ctor__()
		_240_v107 = _nw54
	}
	var _241_v108 _dafny.Map
	_ = _241_v108
	_241_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _112_v1)
	var _242_v109 _dafny.Set
	_ = _242_v109
	_242_v109 = _dafny.SetOf(_241_v108, _241_v108)
	var _243_v110 _dafny.Int
	_ = _243_v110
	_243_v110 = (_242_v109).Cardinality()
	var _244_v111 _dafny.Array
	_ = _244_v111
	var _nwElement0_13 _dafny.Int = (_243_v110)
	_ = _nwElement0_13
	var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(5))
	_ = _nw55
	(_nw55).ArraySet1(_nwElement0_13, 0)
	(_nw55).ArraySet1(Companion_Default___.Fm1(_112_v1, _184_v57, _114_globalState), 1)
	(_nw55).ArraySet1(_113_v2, 2)
	(_nw55).ArraySet1(_113_v2, 3)
	(_nw55).ArraySet1(_113_v2, 4)
	_244_v111 = _nw55
	var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
	_ = _nw56
	_244_v111 = _nw56
	var _245_v112 _dafny.CodePoint
	_ = _245_v112
	_245_v112 = _dafny.CodePoint('j')
	var _246_v113 _dafny.Sequence
	_ = _246_v113
	_246_v113 = _111_v0
	var _247_v114 _dafny.Map
	_ = _247_v114
	_247_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_111_v0, _245_v112), _dafny.Companion_Sequence_.Concatenate((_246_v113), _111_v0))
	var _248_v115 _dafny.Map
	_ = _248_v115
	_248_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _112_v1)
	_247_v114 = Companion_Default___.Fm7(_248_v115, _114_globalState)
	_112_v1 = _112_v1
	var _249_v116 _dafny.Map
	_ = _249_v116
	var _out6 _dafny.Map
	_ = _out6
	_out6 = Companion_Default___.M0(_112_v1, !(_112_v1), _112_v1, _114_globalState)
	_249_v116 = _out6
	var _pat_let_tv1 = _113_v2
	_ = _pat_let_tv1
	var _pat_let_tv2 = _112_v1
	_ = _pat_let_tv2
	var _pat_let_tv3 = _112_v1
	_ = _pat_let_tv3
	_113_v2 = func(_source2 _dafny.Sequence) _dafny.Int {
		var _250___mcc_h0 _dafny.Sequence = _source2
		_ = _250___mcc_h0
		var _251_cf3 _dafny.Sequence = _250___mcc_h0
		_ = _251_cf3
		return (_pat_let_tv1).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv2, _dafny.SeqOf(_pat_let_tv3))).Cardinality())
	}(_246_v113)
	if !(_112_v1) {
		(_114_globalState).F7 = _112_v1
		var _252_v117 _dafny.Sequence
		_ = _252_v117
		_252_v117 = _dafny.SeqOf(_244_v111)
		_244_v111 = (func() _dafny.Array {
			if (_112_v1) && (_112_v1) {
				return _244_v111
			}
			return (_252_v117).Select((Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_252_v117).Cardinality()))).Uint32()).(_dafny.Array)
		})()
		var _253_v118 _dafny.Map
		_ = _253_v118
		_253_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_113_v2).Cmp((func() _dafny.Int {
			if (_184_v57).Contains(_112_v1) {
				return (_184_v57).Get(_112_v1).(_dafny.Int)
			}
			return _113_v2
		})()) != 0, _244_v111)
		_253_v118 = (_253_v118).Update(!(false), _244_v111)
		(_114_globalState).F3 = _113_v2
		(_114_globalState).F8 = (func() _dafny.Sequence {
			if (func() bool {
				if (_248_v115).Contains(_112_v1) {
					return (_248_v115).Get(_112_v1).(bool)
				}
				return _112_v1
			})() {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_111_v0, (Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_111_v0).Cardinality()))).Uint32(), _245_v112), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}((func(_254_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_255_i15 _dafny.Int) _dafny.CodePoint {
						return _254_v112
					}
				})(_245_v112))), (Companion_Default___.SafeIndex(_113_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_256_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_257_i15 _dafny.Int) _dafny.CodePoint {
						return _256_v112
					}
				})(_245_v112)))).Cardinality()))).Uint32(), _245_v112))
			}
			return _111_v0
		})()
	} else {
		var _258_v119 *C0
		_ = _258_v119
		var _nw57 *C0 = New_C0_()
		_ = _nw57
		_nw57.Ctor__()
		_258_v119 = _nw57
		var _259_v120 _dafny.Map
		_ = _259_v120
		_259_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_v119, !(_112_v1))
		var _rhs52 bool = Companion_Default___.Fm2((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, _258_v119)).Cardinality(), false, _114_globalState)
		_ = _rhs52
		var _rhs53 *C0 = _258_v119
		_ = _rhs53
		var _rhs54 bool = (func() bool {
			if _112_v1 {
				return _112_v1
			}
			return (func() bool {
				if (_259_v120).Contains(_258_v119) {
					return (_259_v120).Get(_258_v119).(bool)
				}
				return _112_v1
			})()
		})()
		_ = _rhs54
		var _lhs37 *GlobalState = _114_globalState
		_ = _lhs37
		_112_v1 = _rhs52
		_258_v119 = _rhs53
		_lhs37.F7 = _rhs54
		var _260_v121 bool
		_ = _260_v121
		_260_v121 = _112_v1
		_113_v2 = (((_249_v116).Merge(_248_v115)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v1, (_260_v121)))).Cardinality()
		var _261_v122 _dafny.MultiSet
		_ = _261_v122
		_261_v122 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_258_v119)).Cardinality()))
		_245_v112 = Companion_Default___.Fm0((_261_v122).Cardinality(), _112_v1, (_249_v116).Merge((_249_v116).Update(_112_v1, _112_v1)), _114_globalState)
		(_114_globalState).F7 = _112_v1
		(_114_globalState).F7 = true
	}
	var _262_i16 _dafny.Int
	_ = _262_i16
	_262_i16 = _dafny.Zero
	{
		for _112_v1 {
			{
				if (_262_i16).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_262_i16 = (_262_i16).Plus(_dafny.One)
				_112_v1 = _112_v1
				var _263_v123 *C0
				_ = _263_v123
				var _nw58 *C0 = New_C0_()
				_ = _nw58
				_nw58.Ctor__()
				_263_v123 = _nw58
				_263_v123 = _263_v123
				var _264_v124 *C0
				_ = _264_v124
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__()
				_264_v124 = _nw59
				var _265_v125 _dafny.Map
				_ = _265_v125
				var _out7 _dafny.Map
				_ = _out7
				_out7 = Companion_Default___.M0(_112_v1, _112_v1, (_112_v1) || (_112_v1), _114_globalState)
				_265_v125 = _out7
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _hi3 _dafny.Int = (_113_v2).Minus(_113_v2)
	_ = _hi3
	for _266_i17 := _113_v2; _266_i17.Cmp(_hi3) < 0; _266_i17 = _266_i17.Plus(_dafny.One) {
		var _source3 _dafny.Sequence = _246_v113
		_ = _source3
		var _267___mcc_h1 _dafny.Sequence = _source3
		_ = _267___mcc_h1
		var _268_cf3 _dafny.Sequence = _267___mcc_h1
		_ = _268_cf3
		(_114_globalState).F5 = _266_i17
		_184_v57 = (_184_v57).Update(!(!(_112_v1) || (_112_v1)), _266_i17)
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_172_v48), 0))
		_ = _index46
		(_172_v48).ArraySet1(Companion_Default___.Fm2(_113_v2, _112_v1, _114_globalState), (_index46).Int())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_172_v48), 0))
		_ = _index47
		(_172_v48).ArraySet1(_112_v1, (_index47).Int())
		(_114_globalState).F7 = Companion_Default___.Fm2((_113_v2).Plus(_266_i17), !(_112_v1), _114_globalState)
		var _269_v126 _dafny.Set
		_ = _269_v126
		_269_v126 = _dafny.SetOf((_dafny.Zero).Minus(_266_i17))
		(_114_globalState).F12 = !((_269_v126).IsProperSubsetOf(_dafny.SetOf(_113_v2)))
		var _270_v127 *C0
		_ = _270_v127
		var _nw60 *C0 = New_C0_()
		_ = _nw60
		_nw60.Ctor__()
		_270_v127 = _nw60
		var _271_v128 _dafny.Map
		_ = _271_v128
		_271_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_272_i18 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), _112_v1)
		var _273_v129 _dafny.Sequence
		_ = _273_v129
		_273_v129 = _dafny.SeqOf(_112_v1)
		var _rhs55 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (func() bool {
				if (_271_v128).Contains(_111_v0) {
					return (_271_v128).Get(_111_v0).(bool)
				}
				return false
			})() {
				return _111_v0
			}
			return _111_v0
		})(), _111_v0)
		_ = _rhs55
		var _rhs56 _dafny.Int = (func() _dafny.Int {
			if (_184_v57).Contains(!(true)) {
				return (_184_v57).Get(!(true)).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_273_v129).Cardinality())
		})()
		_ = _rhs56
		var _rhs57 *C0 = _270_v127
		_ = _rhs57
		var _rhs58 bool = _dafny.Companion_Sequence_.IsPrefixOf(_111_v0, _dafny.Companion_Sequence_.Concatenate(_111_v0, _111_v0))
		_ = _rhs58
		var _lhs38 *GlobalState = _114_globalState
		_ = _lhs38
		_111_v0 = _rhs55
		_lhs38.F5 = _rhs56
		_270_v127 = _rhs57
		_112_v1 = _rhs58
		var _274_v130 _dafny.Array
		_ = _274_v130
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_13
		var _nw61 _dafny.Array
		_ = _nw61
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw61 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.CodePoint = (func(_275_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_276_i20 _dafny.Int) _dafny.CodePoint {
					return _275_v112
				}
			})(_245_v112)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw61 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw61).ArraySet1CodePoint(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw61).ArraySet1CodePoint(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_274_v130 = _nw61
		var _277_v131 _dafny.Map
		_ = _277_v131
		_277_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_274_v130, _112_v1)
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_172_v48), 0))
		_ = _index48
		(_172_v48).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-295))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_278_i17 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_279_i19 _dafny.Int) _dafny.Int {
				return _278_i17
			}
		})(_266_i17)))).Cardinality())).Cmp((_277_v131).Cardinality()) > 0, (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_172_v48), 0))
		_ = _index49
		(_172_v48).ArraySet1(!((func() bool {
			if _112_v1 {
				return (_113_v2).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dvpiwucvt")).Cardinality())) <= 0
			}
			return false
		})()), (_index49).Int())
	}
	_112_v1 = !(true) || (_112_v1)
	(_114_globalState).F12 = !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_111_v0, _111_v0), _111_v0))
	_dafny.Print(_111_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_112_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_113_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F4().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_globalState.F8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_114_globalState).F10()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-891))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_globalState.F12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_globalState).F14().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_i7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v48).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_v49).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v57).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v108).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v109).Equals(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_243_v110))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_245_v112)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v113).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_247_v114).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_248_v115).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_249_v116).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_262_i16)
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

func (CompanionStruct_D1_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D1) Dtor_cf1() _dafny.Map {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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

type D2_DC2 struct {
	Cf2 _dafny.Int
}

func (D2_DC2) isD2() {}

func (CompanionStruct_D2_) Create_DC2_(Cf2 _dafny.Int) D2 {
	return D2{D2_DC2{Cf2}}
}

func (_this D2) Is_DC2() bool {
	_, ok := _this.Get_().(D2_DC2)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Int {
	return _dafny.Zero
}

func (_this D2) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D2_DC2).Cf2
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC2:
		{
			return "D2.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC2:
		{
			data2, ok := other.Get_().(D2_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0
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

type D3_DC3 struct {
	Cf3 _dafny.Sequence
}

func (D3_DC3) isD3() {}

func (CompanionStruct_D3_) Create_DC3_(Cf3 _dafny.Sequence) D3 {
	return D3{D3_DC3{Cf3}}
}

func (_this D3) Is_DC3() bool {
	_, ok := _this.Get_().(D3_DC3)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D3) Dtor_cf3() _dafny.Sequence {
	return _this.Get_().(D3_DC3).Cf3
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC3:
		{
			return "D3.DC3" + "(" + data.Cf3.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC3:
		{
			data2, ok := other.Get_().(D3_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
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

type D4_DC4 struct {
	Cf4 _dafny.Map
}

func (D4_DC4) isD4() {}

func (CompanionStruct_D4_) Create_DC4_(Cf4 _dafny.Map) D4 {
	return D4{D4_DC4{Cf4}}
}

func (_this D4) Is_DC4() bool {
	_, ok := _this.Get_().(D4_DC4)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D4) Dtor_cf4() _dafny.Map {
	return _this.Get_().(D4_DC4).Cf4
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC4:
		{
			return "D4.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC4:
		{
			data2, ok := other.Get_().(D4_DC4)
			return ok && data1.Cf4.Equals(data2.Cf4)
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

type D5_DC6 struct {
}

func (D5_DC6) isD5() {}

func (CompanionStruct_D5_) Create_DC6_() D5 {
	return D5{D5_DC6{}}
}

func (_this D5) Is_DC6() bool {
	_, ok := _this.Get_().(D5_DC6)
	return ok
}

type D5_DC5 struct {
	Cf5 _dafny.Map
}

func (D5_DC5) isD5() {}

func (CompanionStruct_D5_) Create_DC5_(Cf5 _dafny.Map) D5 {
	return D5{D5_DC5{Cf5}}
}

func (_this D5) Is_DC5() bool {
	_, ok := _this.Get_().(D5_DC5)
	return ok
}

type D5_DC7 struct {
	Cf6 D5
}

func (D5_DC7) isD5() {}

func (CompanionStruct_D5_) Create_DC7_(Cf6 D5) D5 {
	return D5{D5_DC7{Cf6}}
}

func (_this D5) Is_DC7() bool {
	_, ok := _this.Get_().(D5_DC7)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC6_()
}

func (_this D5) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D5_DC5).Cf5
}

func (_this D5) Dtor_cf6() D5 {
	return _this.Get_().(D5_DC7).Cf6
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC6:
		{
			return "D5.DC6"
		}
	case D5_DC5:
		{
			return "D5.DC5" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D5_DC7:
		{
			return "D5.DC7" + "(" + _dafny.String(data.Cf6) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC6:
		{
			_, ok := other.Get_().(D5_DC6)
			return ok
		}
	case D5_DC5:
		{
			data2, ok := other.Get_().(D5_DC5)
			return ok && data1.Cf5.Equals(data2.Cf5)
		}
	case D5_DC7:
		{
			data2, ok := other.Get_().(D5_DC7)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	F3   _dafny.Int
	F5   _dafny.Int
	F7   bool
	F8   _dafny.Sequence
	F12  bool
	_f0  bool
	_f1  _dafny.Int
	_f2  bool
	_f4  _dafny.Sequence
	_f6  _dafny.Int
	_f9  _dafny.Int
	_f10 _dafny.Map
	_f11 bool
	_f13 bool
	_f14 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = _dafny.Zero
	_this.F5 = _dafny.Zero
	_this.F7 = false
	_this.F8 = _dafny.EmptySeq
	_this.F12 = false
	_this._f0 = false
	_this._f1 = _dafny.Zero
	_this._f2 = false
	_this._f4 = _dafny.EmptySeq
	_this._f6 = _dafny.Zero
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.EmptyMap
	_this._f11 = false
	_this._f13 = false
	_this._f14 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 bool, f3 _dafny.Int, f4 _dafny.Sequence, f5 _dafny.Int, f6 _dafny.Int, f7 bool, f8 _dafny.Sequence, f9 _dafny.Int, f10 _dafny.Map, f11 bool, f12 bool, f13 bool, f14 _dafny.Sequence) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Map {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.Sequence {
	{
		return _this._f14
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
func (_this *C0) Fm4(p0 bool, globalState *GlobalState) bool {
	{
		return true
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
