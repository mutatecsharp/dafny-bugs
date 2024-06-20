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
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 D0, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(869)
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Set, globalState *GlobalState) D0 {
	var _source0 D4 = Companion_D4_.Create_DC14_(true, _dafny.IntOfInt64(574), true, !(false), _dafny.IntOfInt64(674))
	_ = _source0
	if _source0.Is_DC12() {
		return Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("envdq")).Cardinality()))
	} else if _source0.Is_DC13() {
		var _0___mcc_h0 _dafny.Int = _source0.Get_().(D4_DC13).Cf15
		_ = _0___mcc_h0
		var _1___mcc_h1 bool = _source0.Get_().(D4_DC13).Cf16
		_ = _1___mcc_h1
		var _2_cf16 bool = _1___mcc_h1
		_ = _2_cf16
		var _3_cf15 _dafny.Int = _0___mcc_h0
		_ = _3_cf15
		return Companion_D0_.Create_DC0_(_3_cf15)
	} else if _source0.Is_DC14() {
		var _4___mcc_h2 bool = _source0.Get_().(D4_DC14).Cf17
		_ = _4___mcc_h2
		var _5___mcc_h3 _dafny.Int = _source0.Get_().(D4_DC14).Cf18
		_ = _5___mcc_h3
		var _6___mcc_h4 bool = _source0.Get_().(D4_DC14).Cf19
		_ = _6___mcc_h4
		var _7___mcc_h5 bool = _source0.Get_().(D4_DC14).Cf20
		_ = _7___mcc_h5
		var _8___mcc_h6 _dafny.Int = _source0.Get_().(D4_DC14).Cf21
		_ = _8___mcc_h6
		var _9_cf21 _dafny.Int = _8___mcc_h6
		_ = _9_cf21
		var _10_cf20 bool = _7___mcc_h5
		_ = _10_cf20
		var _11_cf19 bool = _6___mcc_h4
		_ = _11_cf19
		var _12_cf18 _dafny.Int = _5___mcc_h3
		_ = _12_cf18
		var _13_cf17 bool = _4___mcc_h2
		_ = _13_cf17
		return Companion_D0_.Create_DC0_((_dafny.MultiSetOf(_12_cf18)).Cardinality())
	} else {
		var _14___mcc_h7 _dafny.Array = _source0.Get_().(D4_DC11).Cf14
		_ = _14___mcc_h7
		var _15_cf14 _dafny.Array = _14___mcc_h7
		_ = _15_cf14
		return Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-39))
	}
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) bool {
	return !((false) && (!(((_dafny.MultiSetOf(true, !(false))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hxyxgi")).Cardinality())) != 0)))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(634)
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(!(true), true, true, !(false), true)).IsDisjointFrom(_dafny.MultiSetOf(false)), true)
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(817))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(485), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(-780)))
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("hymakl")
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_16_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(519), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eedrs")).Cardinality()))
	}))).Cardinality()), (_dafny.Zero).Minus((_dafny.IntOfInt64(-321)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_17_i1 _dafny.Int) _dafny.CodePoint {
		return (Companion_D9_.Create_DC24_(_dafny.CodePoint('y'), true, false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(835), _dafny.IntOfInt64(310))).Cardinality())), false)).Dtor_cf31()
	}))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('b')
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true, true, true, false, false)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)))).Intersection(_dafny.MultiSetOf(false, !(false)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _19_v0 _dafny.CodePoint
			_19_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), _19_v0) {
				_coll0.Add(_19_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()), _dafny.IntOfInt64(834))).Cardinality())
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())).Difference(func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(882), _dafny.IntOfInt64(-520))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _20_v1 _dafny.Int
			_20_v1 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(882)).Cmp(_20_v1) <= 0) && ((_20_v1).Cmp(_dafny.IntOfInt64(-520)) < 0) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_20_v1, _dafny.IntOfInt64(246)))
			}
		}
		return _coll1.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	if false {
		return (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(823))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_21_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(true, false, false, true)
		})))).Difference(_dafny.MultiSetOf(_dafny.SetOf(false)))
	} else {
		return _dafny.MultiSetOf(_dafny.SetOf(true, true))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.MultiSetOf(false, !(false), false), _dafny.MultiSetOf(false, false, false), _dafny.MultiSetFromSeq(_dafny.SeqOf(!(true))), (Companion_D11_.Create_DC28_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, false)))).Dtor_cf42(), _dafny.MultiSetOf(true, false, true))).Union(func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.MultiSetOf(false), (Companion_D11_.Create_DC28_(_dafny.MultiSetOf(true, true))).Dtor_cf42())).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _22_v0 _dafny.MultiSet
			_22_v0 = interface{}(_compr_2).(_dafny.MultiSet)
			if (_dafny.MultiSetOf(_dafny.MultiSetOf(false), (Companion_D11_.Create_DC28_(_dafny.MultiSetOf(true, true))).Dtor_cf42())).Contains(_22_v0) {
				_coll2.Add(_22_v0)
			}
		}
		return _coll2.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Union(_dafny.SetOf(!(true)))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Set, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_((_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())).Cmp(_dafny.IntOfInt64(546)) < 0)
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.CodePoint, p1 D3, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D2_.Create_DC6_(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.SetOf(true, !(true), true)).Cardinality())))))).Dtor_cf6()
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, globalState *GlobalState) D3 {
	if false {
		return Companion_D3_.Create_DC9_(false)
	} else {
		return Companion_D3_.Create_DC9_(false)
	}
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false)).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _23_v0 _dafny.CodePoint
			_23_v0 = interface{}(_compr_3).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false)).Contains(_23_v0) {
				_coll3.Add(_23_v0)
			}
		}
		return _coll3.ToSet()
	}()).Intersection(_dafny.SetOf(_dafny.CodePoint('s')))).Intersection(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfInt64(56))).Keys().Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _24_v1 _dafny.CodePoint
			_24_v1 = interface{}(_compr_4).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfInt64(56))).Contains(_24_v1) {
				_coll4.Add(_24_v1)
			}
		}
		return _coll4.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC14_(true, (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-283), (Companion_D4_.Create_DC14_(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("grjefhcse")).Cardinality())), false, false, _dafny.IntOfInt64(413))).Dtor_cf21())).Cardinality()).Times((_dafny.SetOf(false)).Cardinality())), (_dafny.IntOfInt64(598)).Cmp(_dafny.IntOfInt64(889)) <= 0, (true) && (true), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D3_.Create_DC8_(!(!(true)), _dafny.IntOfInt64(355), _dafny.IntOfInt64(700))), _dafny.SeqOf(Companion_D3_.Create_DC8_(false, _dafny.IntOfUint32((_dafny.SeqOf(!(!(true)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), Companion_D3_.Create_DC8_(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_25_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vt")).Cardinality()))))).Cardinality()))
}
func (_static *CompanionStruct_Default___) M4(p0 bool, p1 _dafny.Set, globalState *GlobalState) (_dafny.Array, _dafny.Int, bool) {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var r2 bool = false
	_ = r2
	if (p0) && (!(p0)) {
		var _26_v0 _dafny.Array
		_ = _26_v0
		var _nwElement0_0 _dafny.Int = _dafny.IntOfInt64(210)
		_ = _nwElement0_0
		var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.One)
		_ = _nw0
		(_nw0).ArraySet1(_nwElement0_0, 0)
		_26_v0 = _nw0
		var _27_v1 _dafny.Array
		_ = _27_v1
		var _nwElement0_1 _dafny.Array = _26_v0
		_ = _nwElement0_1
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(11))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_1, 0)
		(_nw1).ArraySet1(_26_v0, 1)
		(_nw1).ArraySet1(_26_v0, 2)
		(_nw1).ArraySet1(_26_v0, 3)
		(_nw1).ArraySet1(_26_v0, 4)
		(_nw1).ArraySet1(_26_v0, 5)
		(_nw1).ArraySet1(_26_v0, 6)
		(_nw1).ArraySet1(_26_v0, 7)
		(_nw1).ArraySet1(_26_v0, 8)
		(_nw1).ArraySet1(_26_v0, 9)
		(_nw1).ArraySet1(_26_v0, 10)
		_27_v1 = _nw1
		var _28_v2 _dafny.CodePoint
		_ = _28_v2
		_28_v2 = _dafny.CodePoint('l')
		var _29_v3 *C0
		_ = _29_v3
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__((func() bool {
			if p0 {
				return p0
			}
			return p0
		})(), _27_v1, _28_v2)
		_29_v3 = _nw2
		var _30_v4 _dafny.Sequence
		_ = _30_v4
		_30_v4 = _dafny.SeqOf(p0, p0)
		var _31_v5 _dafny.Int
		_ = _31_v5
		_31_v5 = _dafny.IntOfInt64(-101)
		var _32_v6 _dafny.Array
		_ = _32_v6
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw3
		_32_v6 = _nw3
		var _33_v7 _dafny.Map
		_ = _33_v7
		_33_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Array {
			if (_30_v4).Select((Companion_Default___.SafeIndex(_31_v5, _dafny.IntOfUint32((_30_v4).Cardinality()))).Uint32()).(bool) {
				return _32_v6
			}
			return _32_v6
		})())
		_33_v7 = (_33_v7).Update(p0, _32_v6)
		(globalState).F6 = _31_v5
		var _34_v8 T0
		_ = _34_v8
		var _nw4 *C0 = New_C0_()
		_ = _nw4
		_nw4.Ctor__(_29_v3.F27, _27_v1, _28_v2)
		_34_v8 = _nw4
		var _35_v9 _dafny.Array
		_ = _35_v9
		var _nwElement0_2 T0 = _34_v8
		_ = _nwElement0_2
		var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(6))
		_ = _nw5
		(_nw5).ArraySet1(_nwElement0_2, 0)
		(_nw5).ArraySet1(_34_v8, 1)
		(_nw5).ArraySet1(_34_v8, 2)
		(_nw5).ArraySet1(_34_v8, 3)
		(_nw5).ArraySet1(_34_v8, 4)
		(_nw5).ArraySet1(_34_v8, 5)
		_35_v9 = _nw5
		var _36_v10 D4
		_ = _36_v10
		_36_v10 = Companion_D4_.Create_DC11_(_35_v9)
		_36_v10 = _36_v10
		r1 = _31_v5
	} else {
		var _37_v11 _dafny.Int
		_ = _37_v11
		_37_v11 = _dafny.IntOfInt64(-559)
		var _38_v12 _dafny.MultiSet
		_ = _38_v12
		_38_v12 = _dafny.MultiSetOf(_37_v11)
		var _39_v13 D4
		_ = _39_v13
		_39_v13 = Companion_D4_.Create_DC13_(_37_v11, false)
		var _source1 D4 = (func() D4 {
			if ((_38_v12).Cardinality()).Cmp(_37_v11) < 0 {
				return _39_v13
			}
			return _39_v13
		})()
		_ = _source1
		if _source1.Is_DC12() {
			var _40_v14 _dafny.Sequence
			_ = _40_v14
			_40_v14 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_37_v11)).Cardinality()), _37_v11)
			var _41_v15 _dafny.MultiSet
			_ = _41_v15
			_41_v15 = _dafny.MultiSetOf(p0, p0)
			var _42_v16 _dafny.Sequence
			_ = _42_v16
			_42_v16 = _dafny.SeqOf(false, p0, false)
			var _43_v17 _dafny.Map
			_ = _43_v17
			_43_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_41_v15).Cardinality(), _42_v16)
			var _44_v18 D0
			_ = _44_v18
			_44_v18 = Companion_D0_.Create_DC1_(_43_v17)
			(globalState).F3 = (_40_v14).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2(_dafny.IntOfInt64(32), _44_v18, p0, p0, globalState), _dafny.IntOfUint32((_40_v14).Cardinality()))).Uint32()).(_dafny.Int)
			var _45_v19 D9
			_ = _45_v19
			_45_v19 = Companion_D9_.Create_DC25_(Companion_D9_.Create_DC23_(!(Companion_Default___.Fm6(globalState)), _37_v11))
			_45_v19 = _45_v19
			var _rhs0 bool = !(p0)
			_ = _rhs0
			var _rhs1 _dafny.Int = _37_v11
			_ = _rhs1
			r2 = _rhs0
			r1 = _rhs1
			_38_v12 = _38_v12
		} else if _source1.Is_DC13() {
			var _46___mcc_h0 _dafny.Int = _source1.Get_().(D4_DC13).Cf15
			_ = _46___mcc_h0
			var _47___mcc_h1 bool = _source1.Get_().(D4_DC13).Cf16
			_ = _47___mcc_h1
			var _48_cf16 bool = _47___mcc_h1
			_ = _48_cf16
			var _49_cf15 _dafny.Int = _46___mcc_h0
			_ = _49_cf15
			var _50_v20 _dafny.Sequence
			_ = _50_v20
			_50_v20 = _dafny.SeqOf(_37_v11)
			var _51_v21 _dafny.MultiSet
			_ = _51_v21
			_51_v21 = _dafny.MultiSetOf(_50_v20)
			var _52_v22 _dafny.Array
			_ = _52_v22
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw6
			_52_v22 = _nw6
			var _53_v23 _dafny.Map
			_ = _53_v23
			_53_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_51_v21).Union(_dafny.MultiSetOf(_dafny.SeqOf(_37_v11, _37_v11), _50_v20, _50_v20)), _52_v22)
			var _54_v24 _dafny.Map
			_ = _54_v24
			_54_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_cf15, _52_v22)
			var _55_v25 _dafny.Sequence
			_ = _55_v25
			_55_v25 = _dafny.UnicodeSeqOfUtf8Bytes("jnaow")
			_53_v23 = (_53_v23).Update(_51_v21, (func() _dafny.Array {
				if (_54_v24).Contains(_dafny.IntOfUint32((_55_v25).Cardinality())) {
					return (_54_v24).Get(_dafny.IntOfUint32((_55_v25).Cardinality())).(_dafny.Array)
				}
				return _52_v22
			})())
			var _56_v26 _dafny.Array
			_ = _56_v26
			var _nw7 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
			_ = _nw7
			_56_v26 = _nw7
			var _57_v27 D4
			_ = _57_v27
			_57_v27 = Companion_D4_.Create_DC11_(_56_v26)
			var _58_v28 _dafny.MultiSet
			_ = _58_v28
			_58_v28 = _dafny.MultiSetOf(_48_cf16)
			var _59_v29 _dafny.Map
			_ = _59_v29
			_59_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v27, _58_v28)
			var _60_v30 _dafny.CodePoint
			_ = _60_v30
			_60_v30 = _dafny.CodePoint('l')
			var _61_v31 _dafny.Map
			_ = _61_v31
			_61_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_cf16, _60_v30)
			var _62_v32 _dafny.Map
			_ = _62_v32
			_62_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v29, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(503), (_61_v31).Cardinality()))
			_62_v32 = (_62_v32).Update(_59_v29, (_37_v11).Times(_49_cf15))
			var _rhs2 _dafny.Int = _dafny.IntOfInt64(-196)
			_ = _rhs2
			var _rhs3 bool = (_49_cf15).Cmp(_37_v11) < 0
			_ = _rhs3
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			_lhs0.F0 = _rhs2
			_48_cf16 = _rhs3
			var _63_v33 _dafny.Sequence
			_ = _63_v33
			_63_v33 = _dafny.SeqOf(p0)
			var _64_v34 _dafny.Map
			_ = _64_v34
			_64_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_50_v20).Cardinality()), _63_v33)
			var _65_v35 D0
			_ = _65_v35
			_65_v35 = Companion_D0_.Create_DC1_(_64_v34)
			var _66_v36 _dafny.Map
			_ = _66_v36
			_66_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_cf16, _48_cf16)
			var _67_v37 _dafny.Map
			_ = _67_v37
			_67_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_cf15, _48_cf16)
			var _rhs4 _dafny.Int = Companion_Default___.Fm2((_dafny.Zero).Minus(Companion_Default___.Fm2(_37_v11, _65_v35, p0, !(p0), globalState)), _65_v35, (func() bool {
				if (_66_v36).Contains(!(p0)) {
					return (_66_v36).Get(!(p0)).(bool)
				}
				return _48_cf16
			})(), _48_cf16, globalState)
			_ = _rhs4
			var _rhs5 bool = (_48_cf16) == (p0)
			_ = _rhs5
			var _rhs6 bool = !((func() bool {
				if (_67_v37).Contains((_37_v11).Times(_49_cf15)) {
					return (_67_v37).Get((_37_v11).Times(_49_cf15)).(bool)
				}
				return !((_61_v31).Contains(_48_cf16))
			})())
			_ = _rhs6
			var _lhs1 *GlobalState = globalState
			_ = _lhs1
			_lhs1.F6 = _rhs4
			r2 = _rhs5
			r2 = _rhs6
		} else if _source1.Is_DC14() {
			var _68___mcc_h2 bool = _source1.Get_().(D4_DC14).Cf17
			_ = _68___mcc_h2
			var _69___mcc_h3 _dafny.Int = _source1.Get_().(D4_DC14).Cf18
			_ = _69___mcc_h3
			var _70___mcc_h4 bool = _source1.Get_().(D4_DC14).Cf19
			_ = _70___mcc_h4
			var _71___mcc_h5 bool = _source1.Get_().(D4_DC14).Cf20
			_ = _71___mcc_h5
			var _72___mcc_h6 _dafny.Int = _source1.Get_().(D4_DC14).Cf21
			_ = _72___mcc_h6
			var _73_cf21 _dafny.Int = _72___mcc_h6
			_ = _73_cf21
			var _74_cf20 bool = _71___mcc_h5
			_ = _74_cf20
			var _75_cf19 bool = _70___mcc_h4
			_ = _75_cf19
			var _76_cf18 _dafny.Int = _69___mcc_h3
			_ = _76_cf18
			var _77_cf17 bool = _68___mcc_h2
			_ = _77_cf17
			var _78_v38 _dafny.Array
			_ = _78_v38
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw8
			_78_v38 = _nw8
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_78_v38), 0))
			_ = _index0
			(_78_v38).ArraySet1(_73_cf21, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_78_v38), 0))
			_ = _index1
			(_78_v38).ArraySet1(_73_cf21, (_index1).Int())
			var _79_v39 _dafny.Set
			_ = _79_v39
			_79_v39 = _dafny.SetOf(_75_cf19, true)
			_79_v39 = _79_v39
			var _80_v40 _dafny.Sequence
			_ = _80_v40
			_80_v40 = _dafny.SeqOf(_78_v38)
			var _81_v41 _dafny.Set
			_ = _81_v41
			_81_v41 = _dafny.SetOf(_37_v11, _76_cf18)
			var _82_v42 _dafny.Array
			_ = _82_v42
			var _nwElement0_3 _dafny.Array = _78_v38
			_ = _nwElement0_3
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(13))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_3, 0)
			(_nw9).ArraySet1((_80_v40).Select((Companion_Default___.SafeIndex((_81_v41).Cardinality(), _dafny.IntOfUint32((_80_v40).Cardinality()))).Uint32()).(_dafny.Array), 1)
			(_nw9).ArraySet1(_78_v38, 2)
			(_nw9).ArraySet1(_78_v38, 3)
			(_nw9).ArraySet1(_78_v38, 4)
			(_nw9).ArraySet1(_78_v38, 5)
			(_nw9).ArraySet1(_78_v38, 6)
			(_nw9).ArraySet1(_78_v38, 7)
			(_nw9).ArraySet1(_78_v38, 8)
			(_nw9).ArraySet1(_78_v38, 9)
			(_nw9).ArraySet1(_78_v38, 10)
			(_nw9).ArraySet1(_78_v38, 11)
			(_nw9).ArraySet1(_78_v38, 12)
			_82_v42 = _nw9
			_82_v42 = _82_v42
			var _rhs7 bool = _75_cf19
			_ = _rhs7
			var _rhs8 _dafny.Int = (_76_cf18).Plus(_73_cf21)
			_ = _rhs8
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			r2 = _rhs7
			_lhs2.F3 = _rhs8
		} else {
			var _83___mcc_h7 _dafny.Array = _source1.Get_().(D4_DC11).Cf14
			_ = _83___mcc_h7
			var _84_cf14 _dafny.Array = _83___mcc_h7
			_ = _84_cf14
			var _85_v43 _dafny.Array
			_ = _85_v43
			var _len0_0 _dafny.Int = _dafny.One
			_ = _len0_0
			var _nw10 _dafny.Array
			_ = _nw10
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw10 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_86_v11 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_87_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_86_v11)
					}
				})(_37_v11)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw10 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw10).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw10).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_85_v43 = _nw10
			var _88_v44 _dafny.Sequence
			_ = _88_v44
			_88_v44 = _dafny.SeqOf(_37_v11)
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_85_v43), 0))
			_ = _index2
			(_85_v43).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_37_v11, _37_v11), _88_v44), (_index2).Int())
			var _89_v45 _dafny.Map
			_ = _89_v45
			_89_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v13, Companion_Default___.Fm18(_38_v12, globalState))
			var _90_v46 _dafny.Array
			_ = _90_v46
			var _nwElement0_4 bool = p0
			_ = _nwElement0_4
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(10))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_4, 0)
			(_nw11).ArraySet1(p0, 1)
			(_nw11).ArraySet1(!(_89_v45).Contains(Companion_D4_.Create_DC13_(_37_v11, p0)), 2)
			(_nw11).ArraySet1(true, 3)
			(_nw11).ArraySet1(p0, 4)
			(_nw11).ArraySet1((p0) == (p0), 5)
			(_nw11).ArraySet1(p0, 6)
			(_nw11).ArraySet1(!((_37_v11).Cmp(_37_v11) <= 0), 7)
			(_nw11).ArraySet1(p0, 8)
			(_nw11).ArraySet1(p0, 9)
			_90_v46 = _nw11
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_90_v46), 0))
			_ = _index3
			(_90_v46).ArraySet1(false, (_index3).Int())
			var _91_v47 _dafny.Sequence
			_ = _91_v47
			_91_v47 = _dafny.SeqOf(p0)
			var _92_v48 _dafny.Map
			_ = _92_v48
			_92_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_37_v11, _91_v47)
			var _93_v49 D0
			_ = _93_v49
			_93_v49 = Companion_D0_.Create_DC1_(_92_v48)
			var _94_v50 _dafny.Set
			_ = _94_v50
			_94_v50 = _dafny.SetOf((_dafny.Zero).Minus(_37_v11))
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_85_v43), 0))
			_ = _index4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_90_v46), 0))
			_ = _index5
			var _rhs9 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(524))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_95_v44 _dafny.Sequence, _96_v11 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_97_i1 _dafny.Int) _dafny.Int {
					return (_95_v44).Select((Companion_Default___.SafeIndex(_96_v11, _dafny.IntOfUint32((_95_v44).Cardinality()))).Uint32()).(_dafny.Int)
				}
			})(_88_v44, _37_v11))), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_88_v44, (Companion_Default___.SafeIndex(Companion_Default___.Fm2(_37_v11, _93_v49, false, true, globalState), _dafny.IntOfUint32((_88_v44).Cardinality()))).Uint32(), (_94_v50).Cardinality()), _88_v44))
			_ = _rhs9
			var _rhs10 _dafny.Int = (func() _dafny.Int {
				if Companion_Default___.Fm6(globalState) {
					return _37_v11
				}
				return (_dafny.Zero).Minus((_37_v11).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gnpsvlbkw")).Cardinality())))
			})()
			_ = _rhs10
			var _rhs11 bool = p0
			_ = _rhs11
			var _rhs12 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_88_v44, (Companion_Default___.SafeIndex((_37_v11).Times(_37_v11), _dafny.IntOfUint32((_88_v44).Cardinality()))).Uint32(), _37_v11)
			_ = _rhs12
			var _lhs3 _dafny.Array = _85_v43
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_85_v43), 0))
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			var _lhs6 _dafny.Array = _90_v46
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_90_v46), 0))
			_ = _lhs7
			(_lhs3).ArraySet1(_rhs9, (_lhs4).Int())
			_lhs5.F0 = _rhs10
			(_lhs6).ArraySet1(_rhs11, (_lhs7).Int())
			_88_v44 = _rhs12
			var _98_v51 _dafny.Array
			_ = _98_v51
			var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw12
			_98_v51 = _nw12
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_98_v51), 0))
			_ = _index6
			(_98_v51).ArraySet1(_37_v11, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_98_v51), 0))
			_ = _index7
			(_98_v51).ArraySet1(_37_v11, (_index7).Int())
			var _99_v52 _dafny.MultiSet
			_ = _99_v52
			_99_v52 = _dafny.MultiSetOf(_90_v46, _90_v46)
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_98_v51), 0))
			_ = _index8
			(_98_v51).ArraySet1((_dafny.Zero).Minus(((_99_v52).Difference(_99_v52)).Cardinality()), (_index8).Int())
			var _100_v53 *C2
			_ = _100_v53
			var _nw13 *C2 = New_C2_()
			_ = _nw13
			_nw13.Ctor__()
			_100_v53 = _nw13
			var _101_v54 _dafny.Map
			_ = _101_v54
			_101_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(29))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_102_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			})), _100_v53)
			var _103_v55 _dafny.Sequence
			_ = _103_v55
			_103_v55 = _dafny.UnicodeSeqOfUtf8Bytes("wuqfpyidr")
			var _104_v56 D10
			_ = _104_v56
			_104_v56 = Companion_D10_.Create_DC26_(_100_v53)
			_101_v54 = (_101_v54).Update(_dafny.Companion_Sequence_.Concatenate(_103_v55, _103_v55), (_104_v56).Dtor_cf37())
		}
		(globalState).F3 = (_dafny.Zero).Minus((_37_v11).Minus(_37_v11))
		var _105_v57 _dafny.Sequence
		_ = _105_v57
		_105_v57 = _dafny.SeqOf(p0)
		var _106_v58 _dafny.Sequence
		_ = _106_v58
		_106_v58 = _dafny.SeqOf((_37_v11).Minus(_37_v11), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality()), (_dafny.Zero).Minus(_37_v11)), (func() _dafny.Int {
			if p0 {
				return _dafny.IntOfInt64(98)
			}
			return _37_v11
		})())
		var _107_v59 _dafny.Map
		_ = _107_v59
		_107_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p0))
		var _108_v60 _dafny.Sequence
		_ = _108_v60
		_108_v60 = _dafny.UnicodeSeqOfUtf8Bytes("kpuyvgkg")
		var _rhs13 _dafny.Sequence = _105_v57
		_ = _rhs13
		var _rhs14 _dafny.Int = (((_107_v59).Update(p0, p0)).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_108_v60, _108_v60), true)).Cardinality()
		_ = _rhs14
		var _rhs15 _dafny.Sequence = _106_v58
		_ = _rhs15
		var _rhs16 _dafny.Int = _37_v11
		_ = _rhs16
		var _lhs8 *GlobalState = globalState
		_ = _lhs8
		_105_v57 = _rhs13
		_lhs8.F3 = _rhs14
		_106_v58 = _rhs15
		_37_v11 = _rhs16
		(globalState).F0 = _37_v11
		var _109_v61 _dafny.CodePoint
		_ = _109_v61
		_109_v61 = _dafny.CodePoint('l')
		var _110_v62 *C1
		_ = _110_v62
		var _nw14 *C1 = New_C1_()
		_ = _nw14
		_nw14.Ctor__(false, p0, _109_v61)
		_110_v62 = _nw14
		var _111_v63 _dafny.MultiSet
		_ = _111_v63
		_111_v63 = _dafny.MultiSetOf(_110_v62, _110_v62, _110_v62, _110_v62, _110_v62)
		var _112_v64 *C1
		_ = _112_v64
		var _nw15 *C1 = New_C1_()
		_ = _nw15
		_nw15.Ctor__(Companion_Default___.Fm6(globalState), (_111_v63).Contains(_110_v62), (func() _dafny.CodePoint {
			if p0 {
				return _109_v61
			}
			return _109_v61
		})())
		_112_v64 = _nw15
	}
	var _113_v65 _dafny.Int
	_ = _113_v65
	_113_v65 = _dafny.IntOfInt64(94)
	var _114_v66 _dafny.Map
	_ = _114_v66
	_114_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _113_v65)
	var _115_v67 _dafny.Sequence
	_ = _115_v67
	_115_v67 = _dafny.UnicodeSeqOfUtf8Bytes("wonmmp")
	var _116_v68 D3
	_ = _116_v68
	_116_v68 = Companion_D3_.Create_DC9_(true)
	var _117_v69 _dafny.CodePoint
	_ = _117_v69
	_117_v69 = _dafny.CodePoint('p')
	var _pat_let_tv0 = _113_v65
	_ = _pat_let_tv0
	var _pat_let_tv1 = _113_v65
	_ = _pat_let_tv1
	var _rhs17 _dafny.Int = _113_v65
	_ = _rhs17
	var _rhs18 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_115_v67, _dafny.Companion_Sequence_.Concatenate(_115_v67, _115_v67))
	_ = _rhs18
	var _rhs19 _dafny.Map = _114_v66
	_ = _rhs19
	var _rhs20 _dafny.Int = func(_source2 D3) _dafny.Int {
		if _source2.Is_DC8() {
			var _118___mcc_h8 bool = _source2.Get_().(D3_DC8).Cf8
			_ = _118___mcc_h8
			var _119___mcc_h9 _dafny.Int = _source2.Get_().(D3_DC8).Cf9
			_ = _119___mcc_h9
			var _120___mcc_h10 _dafny.Int = _source2.Get_().(D3_DC8).Cf10
			_ = _120___mcc_h10
			var _121_cf10 _dafny.Int = _120___mcc_h10
			_ = _121_cf10
			var _122_cf9 _dafny.Int = _119___mcc_h9
			_ = _122_cf9
			var _123_cf8 bool = _118___mcc_h8
			_ = _123_cf8
			return _122_cf9
		} else if _source2.Is_DC9() {
			var _124___mcc_h11 bool = _source2.Get_().(D3_DC9).Cf11
			_ = _124___mcc_h11
			var _125_cf11 bool = _124___mcc_h11
			_ = _125_cf11
			return _pat_let_tv0
		} else if _source2.Is_DC10() {
			var _126___mcc_h12 T0 = _source2.Get_().(D3_DC10).Cf12
			_ = _126___mcc_h12
			var _127___mcc_h13 _dafny.Int = _source2.Get_().(D3_DC10).Cf13
			_ = _127___mcc_h13
			var _128_cf13 _dafny.Int = _127___mcc_h13
			_ = _128_cf13
			var _129_cf12 T0 = _126___mcc_h12
			_ = _129_cf12
			return _128_cf13
		} else {
			var _130___mcc_h14 _dafny.Array = _source2.Get_().(D3_DC7).Cf7
			_ = _130___mcc_h14
			var _131_cf7 _dafny.Array = _130___mcc_h14
			_ = _131_cf7
			return _pat_let_tv1
		}
	}(_116_v68)
	_ = _rhs20
	var _rhs21 _dafny.CodePoint = _117_v69
	_ = _rhs21
	var _lhs9 *GlobalState = globalState
	_ = _lhs9
	var _lhs10 *GlobalState = globalState
	_ = _lhs10
	var _lhs11 *GlobalState = globalState
	_ = _lhs11
	_lhs9.F3 = _rhs17
	_lhs10.F1 = _rhs18
	_114_v66 = _rhs19
	r1 = _rhs20
	_lhs11.F9 = _rhs21
	var _132_v70 D9
	_ = _132_v70
	_132_v70 = Companion_D9_.Create_DC24_(_117_v69, p0, true, _114_v66, p0)
	var _pat_let_tv2 = _113_v65
	_ = _pat_let_tv2
	var _pat_let_tv3 = _113_v65
	_ = _pat_let_tv3
	var _pat_let_tv4 = _115_v67
	_ = _pat_let_tv4
	if func(_source3 D9) bool {
		if _source3.Is_DC23() {
			var _133___mcc_h15 bool = _source3.Get_().(D9_DC23).Cf29
			_ = _133___mcc_h15
			var _134___mcc_h16 _dafny.Int = _source3.Get_().(D9_DC23).Cf30
			_ = _134___mcc_h16
			var _135_cf30 _dafny.Int = _134___mcc_h16
			_ = _135_cf30
			var _136_cf29 bool = _133___mcc_h15
			_ = _136_cf29
			return !(true) || (_136_cf29)
		} else if _source3.Is_DC24() {
			var _137___mcc_h17 _dafny.CodePoint = _source3.Get_().(D9_DC24).Cf31
			_ = _137___mcc_h17
			var _138___mcc_h18 bool = _source3.Get_().(D9_DC24).Cf32
			_ = _138___mcc_h18
			var _139___mcc_h19 bool = _source3.Get_().(D9_DC24).Cf33
			_ = _139___mcc_h19
			var _140___mcc_h20 _dafny.Map = _source3.Get_().(D9_DC24).Cf34
			_ = _140___mcc_h20
			var _141___mcc_h21 bool = _source3.Get_().(D9_DC24).Cf35
			_ = _141___mcc_h21
			var _142_cf35 bool = _141___mcc_h21
			_ = _142_cf35
			var _143_cf34 _dafny.Map = _140___mcc_h20
			_ = _143_cf34
			var _144_cf33 bool = _139___mcc_h19
			_ = _144_cf33
			var _145_cf32 bool = _138___mcc_h18
			_ = _145_cf32
			var _146_cf31 _dafny.CodePoint = _137___mcc_h17
			_ = _146_cf31
			return (true) || (_145_cf32)
		} else if _source3.Is_DC22() {
			var _147___mcc_h22 _dafny.Sequence = _source3.Get_().(D9_DC22).Cf28
			_ = _147___mcc_h22
			var _148_cf28 _dafny.Sequence = _147___mcc_h22
			_ = _148_cf28
			return true
		} else {
			var _149___mcc_h23 D9 = _source3.Get_().(D9_DC25).Cf36
			_ = _149___mcc_h23
			var _150_cf36 D9 = _149___mcc_h23
			_ = _150_cf36
			return (_pat_let_tv2).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv3, _dafny.IntOfInt64(943), _dafny.IntOfUint32((_pat_let_tv4).Cardinality()))).Cardinality())) > 0
		}
	}(_132_v70) {
		r2 = p0
		var _151_v71 _dafny.Array
		_ = _151_v71
		var _nw16 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
		_ = _nw16
		_151_v71 = _nw16
		var _152_v72 _dafny.Array
		_ = _152_v72
		var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
		_ = _nw17
		_152_v72 = _nw17
		var _153_v73 T0
		_ = _153_v73
		var _nw18 *C0 = New_C0_()
		_ = _nw18
		_nw18.Ctor__(p0, _152_v72, _117_v69)
		_153_v73 = _nw18
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_151_v71), 0))
		_ = _index9
		(_151_v71).ArraySet1(_153_v73, (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_151_v71), 0))
		_ = _index10
		var _rhs22 bool = !(p0)
		_ = _rhs22
		var _rhs23 T0 = _153_v73
		_ = _rhs23
		var _lhs12 _dafny.Array = _151_v71
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_151_v71), 0))
		_ = _lhs13
		r2 = _rhs22
		(_lhs12).ArraySet1(_rhs23, (_lhs13).Int())
		var _154_v74 *C2
		_ = _154_v74
		var _nw19 *C2 = New_C2_()
		_ = _nw19
		_nw19.Ctor__()
		_154_v74 = _nw19
		var _155_v75 _dafny.Array
		_ = _155_v75
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_1
		var _nw20 _dafny.Array
		_ = _nw20
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw20 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_156_p0 bool) func(_dafny.Int) bool {
				return func(_157_i3 _dafny.Int) bool {
					return _156_p0
				}
			})(p0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw20 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw20).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw20).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_155_v75 = _nw20
		var _158_v76 _dafny.Map
		_ = _158_v76
		_158_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v72, _115_v67)
		var _159_v77 _dafny.Map
		_ = _159_v77
		_159_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _115_v67)
		var _160_v78 _dafny.Sequence
		_ = _160_v78
		_160_v78 = _dafny.SeqOf((_159_v77).Cardinality())
		var _161_v79 _dafny.Map
		_ = _161_v79
		_161_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_158_v76).Contains(_152_v72) {
				return (_158_v76).Get(_152_v72).(_dafny.Sequence)
			}
			return _115_v67
		})(), (Companion_Default___.SafeIndex(_113_v65, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_158_v76).Contains(_152_v72) {
				return (_158_v76).Get(_152_v72).(_dafny.Sequence)
			}
			return _115_v67
		})()).Cardinality()))).Uint32(), _117_v69)).Cardinality()), _160_v78)
		var _162_v80 _dafny.Map
		_ = _162_v80
		_162_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v75, _161_v79)
		_162_v80 = (_162_v80).Update(_155_v75, _161_v79)
		var _163_v81 _dafny.MultiSet
		_ = _163_v81
		_163_v81 = _dafny.MultiSetOf(_113_v65, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(428), _113_v65), _113_v65)
		_163_v81 = _163_v81
	} else {
		var _164_v82 _dafny.Map
		_ = _164_v82
		_164_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('l'))
		var _165_v83 _dafny.Map
		_ = _165_v83
		_165_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
			if (_164_v82).Contains(p0) {
				return (_164_v82).Get(p0).(_dafny.CodePoint)
			}
			return _117_v69
		})(), _113_v65)
		var _166_v84 _dafny.Map
		_ = _166_v84
		_166_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v66, _165_v83)
		var _167_v87 _dafny.Sequence
		_ = _167_v87
		_167_v87 = _dafny.SeqOf(_166_v84, ((_166_v84).Update(_114_v66, func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_165_v83).Keys().Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _168_v86 _dafny.CodePoint
					_168_v86 = interface{}(_compr_6).(_dafny.CodePoint)
					if (_165_v83).Contains(_168_v86) {
						_coll6.Add(_168_v86)
					}
				}
				return _coll6.ToSet()
			}()).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _169_v85 _dafny.CodePoint
				_169_v85 = interface{}(_compr_5).(_dafny.CodePoint)
				if (func() _dafny.Set {
					var _coll7 = _dafny.NewBuilder()
					_ = _coll7
					for _iter7 := _dafny.Iterate((_165_v83).Keys().Elements()); ; {
						_compr_7, _ok7 := _iter7()
						if !_ok7 {
							break
						}
						var _170_v86 _dafny.CodePoint
						_170_v86 = interface{}(_compr_7).(_dafny.CodePoint)
						if (_165_v83).Contains(_170_v86) {
							_coll7.Add(_170_v86)
						}
					}
					return _coll7.ToSet()
				}()).Contains(_169_v85) {
					_coll5.Add(_169_v85, _113_v65)
				}
			}
			return _coll5.ToMap()
		}())).Merge(_166_v84), _166_v84)
		var _171_v88 _dafny.Map
		_ = _171_v88
		_171_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v69, p0)
		_166_v84 = (_167_v87).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm23(_113_v65, _171_v88, _dafny.SeqOf(_113_v65, _113_v65, _113_v65, _113_v65, _113_v65), globalState)).Dtor_cf21(), _dafny.IntOfUint32((_167_v87).Cardinality()))).Uint32()).(_dafny.Map)
		if !(p0) || (p0) {
			var _172_v89 _dafny.Array
			_ = _172_v89
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_2
			var _nw21 _dafny.Array
			_ = _nw21
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw21 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) bool = func(_173_i4 _dafny.Int) bool {
					return false
				}
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw21 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw21).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw21).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_172_v89 = _nw21
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_172_v89), 0))
			_ = _index11
			(_172_v89).ArraySet1(p0, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_172_v89), 0))
			_ = _index12
			(_172_v89).ArraySet1(p0, (_index12).Int())
			var _174_v90 _dafny.MultiSet
			_ = _174_v90
			_174_v90 = _dafny.MultiSetOf(_113_v65)
			_174_v90 = _174_v90
			r2 = _dafny.Companion_Sequence_.IsPrefixOf(_115_v67, _115_v67)
			var _175_v91 _dafny.Array
			_ = _175_v91
			var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw22
			_175_v91 = _nw22
			var _176_v92 _dafny.Sequence
			_ = _176_v92
			_176_v92 = _dafny.SeqOf(_175_v91, _175_v91)
			var _177_v93 T0
			_ = _177_v93
			var _nw23 *C0 = New_C0_()
			_ = _nw23
			_nw23.Ctor__((_172_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_172_v89), 0))).Int()).(bool), (_176_v92).Select((Companion_Default___.SafeIndex(_113_v65, _dafny.IntOfUint32((_176_v92).Cardinality()))).Uint32()).(_dafny.Array), _117_v69)
			_177_v93 = _nw23
			_177_v93 = _177_v93
			var _178_v94 _dafny.MultiSet
			_ = _178_v94
			_178_v94 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_115_v67, _115_v67), _115_v67, _115_v67)
			_178_v94 = (_178_v94).Intersection(_178_v94)
		} else {
			var _179_v95 _dafny.Array
			_ = _179_v95
			var _nw24 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
			_ = _nw24
			_179_v95 = _nw24
			var _180_v96 *C2
			_ = _180_v96
			var _nw25 *C2 = New_C2_()
			_ = _nw25
			_nw25.Ctor__()
			_180_v96 = _nw25
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_179_v95), 0))
			_ = _index13
			(_179_v95).ArraySet1(_180_v96, (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_179_v95), 0))
			_ = _index14
			(_179_v95).ArraySet1(_180_v96, (_index14).Int())
			r2 = !(p0)
			(globalState).F3 = _113_v65
			_117_v69 = _117_v69
			(globalState).F0 = _113_v65
		}
		if p0 {
			var _181_v97 _dafny.Array
			_ = _181_v97
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw26
			_181_v97 = _nw26
			var _182_v98 *C0
			_ = _182_v98
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__(p0, _181_v97, _117_v69)
			_182_v98 = _nw27
			var _183_v99 _dafny.Map
			_ = _183_v99
			_183_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v98, p0)
			_183_v99 = (_183_v99).Update(_182_v98, Companion_Default___.Fm6(globalState))
			var _184_v100 _dafny.Array
			_ = _184_v100
			var _len0_3 _dafny.Int = _dafny.One
			_ = _len0_3
			var _nw28 _dafny.Array
			_ = _nw28
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw28 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = func(_185_i5 _dafny.Int) _dafny.Int {
					return (_185_i5).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(231))).Cardinality()))
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw28 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw28).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw28).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_184_v100 = _nw28
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_184_v100), 0))
			_ = _index15
			(_184_v100).ArraySet1(_113_v65, (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_184_v100), 0))
			_ = _index16
			(_184_v100).ArraySet1(_dafny.IntOfUint32((_115_v67).Cardinality()), (_index16).Int())
			(globalState).F3 = ((_dafny.IntOfUint32((_115_v67).Cardinality())).Times((_184_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_184_v100), 0))).Int()).(_dafny.Int))).Minus(_113_v65)
			var _186_v102 _dafny.Map
			_ = _186_v102
			_186_v102 = func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(745))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}((func(_187_v65 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_188_i6 _dafny.Int) _dafny.Int {
						return (_dafny.SetOf(_dafny.SeqOf(_187_v65))).Cardinality()
					}
				})(_113_v65)))).Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _189_v101 _dafny.Int
					_189_v101 = interface{}(_compr_8).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(745))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}((func(_190_v65 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_188_i6 _dafny.Int) _dafny.Int {
							return (_dafny.SetOf(_dafny.SeqOf(_190_v65))).Cardinality()
						}
					})(_113_v65))), _189_v101) {
						_coll8.Add(Companion_Default___.SafeDivisionInt(_189_v101, _dafny.IntOfInt64(907)), _182_v98.F27)
					}
				}
				return _coll8.ToMap()
			}()
			_186_v102 = _186_v102
			var _191_v103 D0
			_ = _191_v103
			_191_v103 = Companion_D0_.Create_DC2_(_182_v98.F27)
			var _192_v104 _dafny.Map
			_ = _192_v104
			_192_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_191_v103), _182_v98.F27)
			var _193_v105 _dafny.Array
			_ = _193_v105
			var _nwElement0_5 bool = _182_v98.F27
			_ = _nwElement0_5
			var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(29))
			_ = _nw29
			(_nw29).ArraySet1(_nwElement0_5, 0)
			(_nw29).ArraySet1(p0, 1)
			(_nw29).ArraySet1(p0, 2)
			(_nw29).ArraySet1(p0, 3)
			(_nw29).ArraySet1(p0, 4)
			(_nw29).ArraySet1((_191_v103).Dtor_cf2(), 5)
			(_nw29).ArraySet1(!(_182_v98.F27), 6)
			(_nw29).ArraySet1(_182_v98.F27, 7)
			(_nw29).ArraySet1(_182_v98.F27, 8)
			(_nw29).ArraySet1(p0, 9)
			(_nw29).ArraySet1(_182_v98.F27, 10)
			(_nw29).ArraySet1(_182_v98.F27, 11)
			(_nw29).ArraySet1((_182_v98).Fm7(_192_v104, Companion_Default___.Fm6(globalState), globalState), 12)
			(_nw29).ArraySet1(_182_v98.F27, 13)
			(_nw29).ArraySet1(p0, 14)
			(_nw29).ArraySet1(p0, 15)
			(_nw29).ArraySet1(p0, 16)
			(_nw29).ArraySet1(p0, 17)
			(_nw29).ArraySet1(false, 18)
			(_nw29).ArraySet1(_182_v98.F27, 19)
			(_nw29).ArraySet1(Companion_Default___.Fm6(globalState), 20)
			(_nw29).ArraySet1(p0, 21)
			(_nw29).ArraySet1(p0, 22)
			(_nw29).ArraySet1(false, 23)
			(_nw29).ArraySet1(_182_v98.F27, 24)
			(_nw29).ArraySet1(p0, 25)
			(_nw29).ArraySet1(p0, 26)
			(_nw29).ArraySet1(_182_v98.F27, 27)
			(_nw29).ArraySet1((_182_v98).Fm7(_192_v104, Companion_Default___.Fm6(globalState), globalState), 28)
			_193_v105 = _nw29
			var _194_v106 _dafny.Map
			_ = _194_v106
			_194_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v105, _182_v98.F27)
			var _195_v107 _dafny.Map
			_ = _195_v107
			_195_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v69, _115_v67)
			_194_v106 = (_194_v106).Update(_193_v105, (_195_v107).Equals(_195_v107))
		} else {
			var _196_v108 *C1
			_ = _196_v108
			var _nw30 *C1 = New_C1_()
			_ = _nw30
			_nw30.Ctor__(p0, !(p0) || (p0), _117_v69)
			_196_v108 = _nw30
			var _197_v109 _dafny.Array
			_ = _197_v109
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw31
			_197_v109 = _nw31
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_197_v109), 0))
			_ = _index17
			(_197_v109).ArraySet1((_196_v108).F26(), (_index17).Int())
			var _198_v110 _dafny.Set
			_ = _198_v110
			_198_v110 = _dafny.SetOf((_196_v108).Fm3(_dafny.SetOf(_113_v65, _113_v65), (_196_v108).F26(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_199_v69 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_200_i7 _dafny.Int) _dafny.CodePoint {
					return _199_v69
				}
			})(_117_v69))), (_196_v108).F25(), globalState))
			var _201_v112 _dafny.MultiSet
			_ = _201_v112
			_201_v112 = _dafny.MultiSetOf((func() _dafny.Set {
				var _coll9 = _dafny.NewBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(154))).Elements()); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _202_v111 _dafny.Int
					_202_v111 = interface{}(_compr_9).(_dafny.Int)
					if (_dafny.SetOf(_dafny.IntOfInt64(154))).Contains(_202_v111) {
						_coll9.Add(Companion_Default___.SafeDivisionInt(_202_v111, _dafny.IntOfInt64(752)))
					}
				}
				return _coll9.ToSet()
			}()).IsProperSubsetOf(_198_v110))
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_197_v109), 0))
			_ = _index18
			var _rhs24 _dafny.Int = (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(458))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_203_v108 *C1, _204_v65 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_205_i8 _dafny.Int) _dafny.Int {
					return (func() _dafny.Int {
						if (_203_v108).F25() {
							return _204_v65
						}
						return _204_v65
					})()
				}
			})(_196_v108, _113_v65))))).Cardinality()
			_ = _rhs24
			var _rhs25 bool = Companion_Default___.Fm6(globalState)
			_ = _rhs25
			var _rhs26 _dafny.MultiSet = ((_201_v112).Difference(_201_v112)).Difference((_201_v112).Intersection(_201_v112))
			_ = _rhs26
			var _lhs14 _dafny.Array = _197_v109
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_197_v109), 0))
			_ = _lhs15
			_113_v65 = _rhs24
			(_lhs14).ArraySet1(_rhs25, (_lhs15).Int())
			_201_v112 = _rhs26
			var _206_v113 _dafny.Sequence
			_ = _206_v113
			_206_v113 = _dafny.SeqOf(false)
			var _207_v114 _dafny.Map
			_ = _207_v114
			_207_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_196_v108).F26(), _115_v67)
			var _208_v115 _dafny.Map
			_ = _208_v115
			_208_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_206_v113).Select((Companion_Default___.SafeIndex(_113_v65, _dafny.IntOfUint32((_206_v113).Cardinality()))).Uint32()).(bool) {
					return (func() _dafny.Sequence {
						if (_207_v114).Contains(false) {
							return (_207_v114).Get(false).(_dafny.Sequence)
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(63))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg12 _dafny.Int) interface{} {
								return coer12(arg12)
							}
						}((func(_209_v69 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_210_i9 _dafny.Int) _dafny.CodePoint {
								return _209_v69
							}
						})(_117_v69)))
					})()
				}
				return _115_v67
			})(), (_196_v108).F26())
			var _211_v116 _dafny.MultiSet
			_ = _211_v116
			_211_v116 = _dafny.MultiSetOf(_113_v65, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("amooda")).Cardinality()))
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_197_v109), 0))
			_ = _index19
			(_197_v109).ArraySet1((func() bool {
				if (_208_v115).Contains(_115_v67) {
					return (_208_v115).Get(_115_v67).(bool)
				}
				return !(_211_v116).Contains((_dafny.Zero).Minus(_113_v65))
			})(), (_index19).Int())
			(globalState).F6 = (_113_v65).Minus(_113_v65)
			var _212_v117 T0
			_ = _212_v117
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__(Companion_Default___.Fm6(globalState), (_196_v108).F25(), _117_v69)
			_212_v117 = _nw32
			var _213_v118 _dafny.Sequence
			_ = _213_v118
			_213_v118 = _dafny.SeqOf(_212_v117, _212_v117)
			var _214_v119 _dafny.Map
			_ = _214_v119
			_214_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v65, _206_v113)
			var _215_v120 D0
			_ = _215_v120
			_215_v120 = Companion_D0_.Create_DC1_(_214_v119)
			var _216_v121 _dafny.Array
			_ = _216_v121
			var _nwElement0_6 _dafny.Int = (_dafny.Zero).Minus(_113_v65)
			_ = _nwElement0_6
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(18))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_6, 0)
			(_nw33).ArraySet1(_113_v65, 1)
			(_nw33).ArraySet1(_113_v65, 2)
			(_nw33).ArraySet1(_113_v65, 3)
			(_nw33).ArraySet1((_113_v65).Times(_113_v65), 4)
			(_nw33).ArraySet1(_113_v65, 5)
			(_nw33).ArraySet1((_dafny.MultiSetFromSeq(_213_v118)).Cardinality(), 6)
			(_nw33).ArraySet1(_113_v65, 7)
			(_nw33).ArraySet1(_dafny.IntOfInt64(480), 8)
			(_nw33).ArraySet1(_113_v65, 9)
			(_nw33).ArraySet1(_113_v65, 10)
			(_nw33).ArraySet1(_dafny.IntOfInt64(-630), 11)
			(_nw33).ArraySet1(_113_v65, 12)
			(_nw33).ArraySet1((_113_v65).Times(_dafny.IntOfUint32((_206_v113).Cardinality())), 13)
			(_nw33).ArraySet1(_113_v65, 14)
			(_nw33).ArraySet1(_113_v65, 15)
			(_nw33).ArraySet1(Companion_Default___.Fm2(_113_v65, _215_v120, (_196_v108).F26(), (_197_v109).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_197_v109), 0))).Int()).(bool), globalState), 16)
			(_nw33).ArraySet1(_dafny.IntOfInt64(404), 17)
			_216_v121 = _nw33
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_216_v121), 0))
			_ = _index20
			(_216_v121).ArraySet1(_dafny.IntOfUint32((_115_v67).Cardinality()), (_index20).Int())
			var _217_v122 _dafny.Map
			_ = _217_v122
			_217_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_216_v121), 0))
			_ = _index21
			(_216_v121).ArraySet1((Companion_Default___.Fm8(_113_v65, _113_v65, (func() bool {
				if (_217_v122).Contains((_197_v109).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_197_v109), 0))).Int()).(bool)) {
					return (_217_v122).Get((_197_v109).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_197_v109), 0))).Int()).(bool)).(bool)
				}
				return p0
			})(), globalState)).Plus(Companion_Default___.SafeModuloInt(_113_v65, _dafny.IntOfUint32((_213_v118).Cardinality()))), (_index21).Int())
		}
		var _218_v123 _dafny.Sequence
		_ = _218_v123
		_218_v123 = _dafny.SeqOf(p0)
		_218_v123 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_218_v123, _218_v123), _218_v123)
		var _219_v126 D0
		_ = _219_v126
		_219_v126 = Companion_D0_.Create_DC1_(func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(61), _dafny.IntOfInt64(555))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _220_v124 _dafny.Int
				_220_v124 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(61)).Cmp(_220_v124) <= 0) && ((_220_v124).Cmp(_dafny.IntOfInt64(555)) < 0) {
					_coll10.Add((_220_v124).Minus((func() _dafny.Map {
						var _coll11 = _dafny.NewMapBuilder()
						_ = _coll11
						for _iter11 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg13 _dafny.Int) interface{} {
								return coer13(arg13)
							}
						}(func(_221_i10 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wne")).Cardinality())
						}))).Elements()); ; {
							_compr_11, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _222_v125 _dafny.Int
							_222_v125 = interface{}(_compr_11).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg14 _dafny.Int) interface{} {
									return coer14(arg14)
								}
							}(func(_221_i10 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wne")).Cardinality())
							})), _222_v125) {
								_coll11.Add(Companion_Default___.SafeModuloInt(_222_v125, _113_v65), p0)
							}
						}
						return _coll11.ToMap()
					}()).Cardinality()), _218_v123)
				}
			}
			return _coll10.ToMap()
		}())
		(globalState).F3 = Companion_Default___.Fm2(_dafny.IntOfUint32((_115_v67).Cardinality()), _219_v126, p0, p0, globalState)
	}
	var _223_i11 _dafny.Int
	_ = _223_i11
	_223_i11 = _dafny.Zero
	{
		for (p0) && (p0) {
			{
				if (_223_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_223_i11 = (_223_i11).Plus(_dafny.One)
				var _224_v127 *C2
				_ = _224_v127
				var _nw34 *C2 = New_C2_()
				_ = _nw34
				_nw34.Ctor__()
				_224_v127 = _nw34
				_224_v127 = _224_v127
				r2 = p0
				if !(p0) {
					var _225_v128 _dafny.Map
					_ = _225_v128
					_225_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v69, (_113_v65).Cmp(_113_v65) >= 0)
					_225_v128 = (_225_v128).Update(_117_v69, false)
					var _226_v129 _dafny.Array
					_ = _226_v129
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_4
					var _nw35 _dafny.Array
					_ = _nw35
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw35 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) _dafny.Int = func(_227_i12 _dafny.Int) _dafny.Int {
							return (_227_i12).Minus(_dafny.IntOfInt64(916))
						}
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw35 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw35).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw35).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_226_v129 = _nw35
					var _228_v130 _dafny.Set
					_ = _228_v130
					_228_v130 = _dafny.SetOf(_226_v129, _226_v129, _226_v129)
					_228_v130 = (_228_v130).Union(_228_v130)
					var _229_v131 _dafny.Sequence
					_ = _229_v131
					_229_v131 = _dafny.SeqOf(false, p0)
					var _230_v132 _dafny.Map
					_ = _230_v132
					_230_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v65, _229_v131)
					var _231_v133 D0
					_ = _231_v133
					_231_v133 = Companion_D0_.Create_DC1_(_230_v132)
					var _232_v134 _dafny.Array
					_ = _232_v134
					var _len0_5 _dafny.Int = _dafny.One
					_ = _len0_5
					var _nw36 _dafny.Array
					_ = _nw36
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw36 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_233_p0 bool) func(_dafny.Int) _dafny.MultiSet {
							return func(_234_i13 _dafny.Int) _dafny.MultiSet {
								return _dafny.MultiSetOf(_233_p0, false)
							}
						})(p0)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw36 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw36).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw36).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_232_v134 = _nw36
					var _235_v135 _dafny.Map
					_ = _235_v135
					_235_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v133, _232_v134)
					_235_v135 = (_235_v135).Update(_231_v133, _232_v134)
					r2 = true
					var _236_v136 _dafny.Array
					_ = _236_v136
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_6
					var _nw37 _dafny.Array
					_ = _nw37
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw37 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.Sequence = (func(_237_v67 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_238_i14 _dafny.Int) _dafny.Sequence {
								return _237_v67
							}
						})(_115_v67)
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw37 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw37).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw37).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_236_v136 = _nw37
					_236_v136 = _236_v136
				} else {
					var _239_v137 _dafny.Array
					_ = _239_v137
					var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
					_ = _nw38
					_239_v137 = _nw38
					var _240_v138 *C0
					_ = _240_v138
					var _nw39 *C0 = New_C0_()
					_ = _nw39
					_nw39.Ctor__(p0, _239_v137, (_224_v127).Fm0(_113_v65, globalState))
					_240_v138 = _nw39
					var _241_v139 _dafny.Sequence
					_ = _241_v139
					_241_v139 = _dafny.SeqOf(_240_v138, _240_v138)
					var _242_v140 T0
					_ = _242_v140
					var _nw40 *C0 = New_C0_()
					_ = _nw40
					_nw40.Ctor__(_240_v138.F27, (_240_v138).F28(), _dafny.CodePoint('m'))
					_242_v140 = _nw40
					var _243_v141 _dafny.MultiSet
					_ = _243_v141
					_243_v141 = _dafny.MultiSetOf(_113_v65)
					var _244_v142 D3
					_ = _244_v142
					_244_v142 = Companion_D3_.Create_DC10_(_242_v140, (_243_v141).Cardinality())
					var _245_v143 _dafny.MultiSet
					_ = _245_v143
					_245_v143 = _dafny.MultiSetOf(p0, p0)
					var _246_v144 _dafny.Map
					_ = _246_v144
					_246_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_241_v139).Select((Companion_Default___.SafeIndex((_244_v142).Dtor_cf13(), _dafny.IntOfUint32((_241_v139).Cardinality()))).Uint32()).(*C0), _245_v143)
					_246_v144 = _246_v144
					var _247_v145 _dafny.Array
					_ = _247_v145
					var _nwElement0_7 bool = p0
					_ = _nwElement0_7
					var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(2))
					_ = _nw41
					(_nw41).ArraySet1(_nwElement0_7, 0)
					(_nw41).ArraySet1(_240_v138.F27, 1)
					_247_v145 = _nw41
					var _248_v146 D0
					_ = _248_v146
					_248_v146 = Companion_D0_.Create_DC2_(p0)
					var _249_v147 D10
					_ = _249_v147
					_249_v147 = Companion_D10_.Create_DC27_(_113_v65, _247_v145, _113_v65, (_248_v146).Dtor_cf2())
					var _pat_let_tv5 = _113_v65
					_ = _pat_let_tv5
					var _pat_let_tv6 = _115_v67
					_ = _pat_let_tv6
					var _pat_let_tv7 = _115_v67
					_ = _pat_let_tv7
					_249_v147 = func(_pat_let0_0 D10) D10 {
						return func(_250_dt__update__tmp_h0 D10) D10 {
							return func(_pat_let1_0 _dafny.Int) D10 {
								return func(_251_dt__update_hcf40_h0 _dafny.Int) D10 {
									return func(_pat_let2_0 _dafny.Int) D10 {
										return func(_252_dt__update_hcf38_h0 _dafny.Int) D10 {
											return Companion_D10_.Create_DC27_(_252_dt__update_hcf38_h0, (_250_dt__update__tmp_h0).Dtor_cf39(), _251_dt__update_hcf40_h0, (_250_dt__update__tmp_h0).Dtor_cf41())
										}(_pat_let2_0)
									}(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_pat_let_tv6, _pat_let_tv7)).Cardinality()))
								}(_pat_let1_0)
							}(_pat_let_tv5)
						}(_pat_let0_0)
					}(_249_v147)
					(_240_v138).F27 = _240_v138.F27
					var _253_v148 *C2
					_ = _253_v148
					var _nw42 *C2 = New_C2_()
					_ = _nw42
					_nw42.Ctor__()
					_253_v148 = _nw42
					(_240_v138).F27 = _240_v138.F27
				}
				(globalState).F6 = _113_v65
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _254_v149 _dafny.Map
	_ = _254_v149
	_254_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v69, p0)
	r2 = (_254_v149).Contains(_117_v69)
	var _255_v150 _dafny.Sequence
	_ = _255_v150
	_255_v150 = _dafny.SeqOf(_dafny.IntOfUint32((_115_v67).Cardinality()))
	var _256_v151 _dafny.Map
	_ = _256_v151
	_256_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_113_v65).Cmp(_113_v65) >= 0, _dafny.Companion_Sequence_.Concatenate(_255_v150, _255_v150))
	_256_v151 = (_256_v151).Update(!(p0), _255_v150)
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(2)
	_ = _len0_7
	var _nw43 _dafny.Array
	_ = _nw43
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw43 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) _dafny.Int = (func(_257_v65 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_258_i15 _dafny.Int) _dafny.Int {
				return (_258_i15).Minus(_257_v65)
			}
		})(_113_v65)
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw43 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw43).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw43).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	r0 = _nw43
	r1 = Companion_Default___.SafeModuloInt(_113_v65, _113_v65)
	r2 = !(Companion_Default___.Fm6(globalState))
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _259_v0 _dafny.Sequence
	_ = _259_v0
	_259_v0 = _dafny.UnicodeSeqOfUtf8Bytes("cliqmvcxh")
	var _260_v1 _dafny.Int
	_ = _260_v1
	_260_v1 = _dafny.IntOfInt64(584)
	var _261_v2 _dafny.Map
	_ = _261_v2
	_261_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(979))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_262_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	})))
	var _263_v3 _dafny.Sequence
	_ = _263_v3
	_263_v3 = _dafny.SeqOf(_260_v1)
	var _264_v4 _dafny.Set
	_ = _264_v4
	_264_v4 = _dafny.SetOf(_263_v3, _263_v3, _263_v3, _263_v3, _263_v3)
	var _265_v5 bool
	_ = _265_v5
	_265_v5 = true
	var _266_v6 _dafny.Map
	_ = _266_v6
	_266_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v5, _265_v5)
	var _267_v7 _dafny.Set
	_ = _267_v7
	_267_v7 = _dafny.SetOf(_265_v5)
	var _268_v8 _dafny.Array
	_ = _268_v8
	var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
	_ = _nw44
	_268_v8 = _nw44
	var _269_v9 _dafny.Array
	_ = _269_v9
	var _nwElement0_8 bool = _265_v5
	_ = _nwElement0_8
	var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(28))
	_ = _nw45
	(_nw45).ArraySet1(_nwElement0_8, 0)
	(_nw45).ArraySet1(_265_v5, 1)
	(_nw45).ArraySet1(!(_265_v5), 2)
	(_nw45).ArraySet1(_265_v5, 3)
	(_nw45).ArraySet1(_265_v5, 4)
	(_nw45).ArraySet1(_265_v5, 5)
	(_nw45).ArraySet1(_265_v5, 6)
	(_nw45).ArraySet1(_265_v5, 7)
	(_nw45).ArraySet1(_265_v5, 8)
	(_nw45).ArraySet1(_265_v5, 9)
	(_nw45).ArraySet1(_265_v5, 10)
	(_nw45).ArraySet1(_265_v5, 11)
	(_nw45).ArraySet1(_265_v5, 12)
	(_nw45).ArraySet1(_265_v5, 13)
	(_nw45).ArraySet1(_265_v5, 14)
	(_nw45).ArraySet1(_265_v5, 15)
	(_nw45).ArraySet1(_265_v5, 16)
	(_nw45).ArraySet1(_265_v5, 17)
	(_nw45).ArraySet1(true, 18)
	(_nw45).ArraySet1(_265_v5, 19)
	(_nw45).ArraySet1(_265_v5, 20)
	(_nw45).ArraySet1(_265_v5, 21)
	(_nw45).ArraySet1(_265_v5, 22)
	(_nw45).ArraySet1(true, 23)
	(_nw45).ArraySet1(!(_265_v5), 24)
	(_nw45).ArraySet1(_265_v5, 25)
	(_nw45).ArraySet1(_265_v5, 26)
	(_nw45).ArraySet1(!(_265_v5), 27)
	_269_v9 = _nw45
	var _270_v10 _dafny.Array
	_ = _270_v10
	var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
	_ = _nw46
	_270_v10 = _nw46
	var _271_v11 _dafny.Sequence
	_ = _271_v11
	_271_v11 = _dafny.SeqOf(_270_v10, _270_v10, _270_v10)
	var _272_globalState *GlobalState
	_ = _272_globalState
	var _nw47 *GlobalState = New_GlobalState_()
	_ = _nw47
	_nw47.Ctor__(_dafny.IntOfInt64(800), _dafny.Companion_Sequence_.Concatenate(_259_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(205))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_273_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))), (func() _dafny.Sequence {
		if (_261_v2).Contains((_dafny.Zero).Minus(_260_v1)) {
			return (_261_v2).Get((_dafny.Zero).Minus(_260_v1)).(_dafny.Sequence)
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("iqu")
	})(), _dafny.IntOfInt64(-913), _dafny.IntOfInt64(482), _dafny.IntOfInt64(-153), _dafny.IntOfInt64(499), _264_v4, _266_v6, _dafny.CodePoint('k'), _dafny.IntOfInt64(340), false, _267_v7, _268_v8, false, true, false, _dafny.UnicodeSeqOfUtf8Bytes("nkfdofn"), false, false, false, false, _269_v9, (_271_v11).Select((Companion_Default___.SafeIndex(_260_v1, _dafny.IntOfUint32((_271_v11).Cardinality()))).Uint32()).(_dafny.Array))
	_272_globalState = _nw47
	var _274_v12 _dafny.MultiSet
	_ = _274_v12
	_274_v12 = _dafny.MultiSetOf(_260_v1, _260_v1, _dafny.IntOfInt64(346))
	var _275_v13 _dafny.Map
	_ = _275_v13
	_275_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v5, _dafny.IntOfInt64(-87))
	_265_v5 = ((_274_v12).Cardinality()).Cmp(((_275_v13).Update(_265_v5, _dafny.IntOfUint32((_259_v0).Cardinality()))).Cardinality()) <= 0
	var _hi0 _dafny.Int = _dafny.IntOfInt64(619)
	_ = _hi0
	for _276_i2 := (_260_v1).Times(_260_v1); _276_i2.Cmp(_hi0) < 0; _276_i2 = _276_i2.Plus(_dafny.One) {
		var _277_v14 D0
		_ = _277_v14
		_277_v14 = Companion_D0_.Create_DC0_(_260_v1)
		var _source4 D0 = _277_v14
		_ = _source4
		if _source4.Is_DC0() {
			var _278___mcc_h0 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
			_ = _278___mcc_h0
			var _279_cf0 _dafny.Int = _278___mcc_h0
			_ = _279_cf0
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_269_v9), 0))
			_ = _index22
			(_269_v9).ArraySet1(_265_v5, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_269_v9), 0))
			_ = _index23
			(_269_v9).ArraySet1((_267_v7).IsProperSubsetOf(_267_v7), (_index23).Int())
			_277_v14 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(_260_v1))
			var _280_v15 *C2
			_ = _280_v15
			var _nw48 *C2 = New_C2_()
			_ = _nw48
			_nw48.Ctor__()
			_280_v15 = _nw48
			var _281_v16 _dafny.Map
			_ = _281_v16
			_281_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(401), _265_v5)
			var _282_v17 _dafny.Map
			_ = _282_v17
			_282_v17 = _281_v16
			var _283_v18 _dafny.Sequence
			_ = _283_v18
			_283_v18 = _dafny.SeqOf((_282_v17), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v1, (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool)), _281_v16)
			_283_v18 = (func() _dafny.Sequence {
				if _265_v5 {
					return _dafny.Companion_Sequence_.Concatenate(_283_v18, _283_v18)
				}
				return _283_v18
			})()
		} else if _source4.Is_DC1() {
			var _284___mcc_h1 _dafny.Map = _source4.Get_().(D0_DC1).Cf1
			_ = _284___mcc_h1
			var _285_cf1 _dafny.Map = _284___mcc_h1
			_ = _285_cf1
			_267_v7 = Companion_Default___.Fm18(_274_v12, _272_globalState)
			var _286_v19 _dafny.CodePoint
			_ = _286_v19
			_286_v19 = _dafny.CodePoint('g')
			var _287_v20 _dafny.Set
			_ = _287_v20
			_287_v20 = _dafny.SetOf(_286_v19, _286_v19)
			var _288_v21 _dafny.Array
			_ = _288_v21
			var _289_v22 _dafny.Int
			_ = _289_v22
			var _290_v23 bool
			_ = _290_v23
			var _out0 _dafny.Array
			_ = _out0
			var _out1 _dafny.Int
			_ = _out1
			var _out2 bool
			_ = _out2
			_out0, _out1, _out2 = Companion_Default___.M4(_265_v5, _287_v20, _272_globalState)
			_288_v21 = _out0
			_289_v22 = _out1
			_290_v23 = _out2
			var _291_v24 _dafny.Array
			_ = _291_v24
			var _nw49 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
			_ = _nw49
			_291_v24 = _nw49
			var _292_v25 T0
			_ = _292_v25
			var _nw50 *C1 = New_C1_()
			_ = _nw50
			_nw50.Ctor__(_290_v23, _290_v23, _286_v19)
			_292_v25 = _nw50
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_291_v24), 0))
			_ = _index24
			(_291_v24).ArraySet1(_292_v25, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_291_v24), 0))
			_ = _index25
			(_291_v24).ArraySet1(_292_v25, (_index25).Int())
			var _293_v26 _dafny.Map
			_ = _293_v26
			_293_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_289_v22, (_dafny.Zero).Minus(_289_v22))
			_293_v26 = _293_v26
		} else {
			var _294___mcc_h2 bool = _source4.Get_().(D0_DC2).Cf2
			_ = _294___mcc_h2
			var _295_cf2 bool = _294___mcc_h2
			_ = _295_cf2
			var _296_v27 *C2
			_ = _296_v27
			var _nw51 *C2 = New_C2_()
			_ = _nw51
			_nw51.Ctor__()
			_296_v27 = _nw51
			var _297_v28 D2
			_ = _297_v28
			_297_v28 = Companion_D2_.Create_DC6_(_263_v3)
			var _298_v29 _dafny.CodePoint
			_ = _298_v29
			_298_v29 = _dafny.CodePoint('y')
			var _299_v30 *C0
			_ = _299_v30
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__(!_dafny.Companion_Sequence_.Equal(_263_v3, (_297_v28).Dtor_cf6()), _268_v8, _298_v29)
			_299_v30 = _nw52
			var _300_v31 _dafny.Array
			_ = _300_v31
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
			_ = _nw53
			_300_v31 = _nw53
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_300_v31), 0))
			_ = _index26
			(_300_v31).ArraySet1(_276_i2, (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_300_v31), 0))
			_ = _index27
			(_300_v31).ArraySet1(_dafny.IntOfUint32((_259_v0).Cardinality()), (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_269_v9), 0))
			_ = _index28
			(_269_v9).ArraySet1(_299_v30.F27, (_index28).Int())
			var _301_v32 _dafny.Map
			_ = _301_v32
			_301_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_295_cf2) || (_299_v30.F27), _299_v30)
			var _302_v33 _dafny.Map
			_ = _302_v33
			_302_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(_dafny.IntOfUint32((_259_v0).Cardinality()))).Plus(_276_i2), _295_cf2)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_300_v31), 0))
			_ = _index29
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_300_v31), 0))
			_ = _index30
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_269_v9), 0))
			_ = _index31
			var _rhs27 *C0 = (func() *C0 {
				if (_301_v32).Contains(!(!(_295_cf2) || (_295_cf2))) {
					return (_301_v32).Get(!(!(_295_cf2) || (_295_cf2))).(*C0)
				}
				return _299_v30
			})()
			_ = _rhs27
			var _rhs28 _dafny.Int = Companion_Default___.SafeModuloInt(_276_i2, ((func() _dafny.Int {
				if (_275_v13).Contains(_295_cf2) {
					return (_275_v13).Get(_295_cf2).(_dafny.Int)
				}
				return _260_v1
			})()).Minus(_260_v1))
			_ = _rhs28
			var _rhs29 _dafny.Int = Companion_Default___.Fm8(_276_i2, _dafny.IntOfInt64(427), _265_v5, _272_globalState)
			_ = _rhs29
			var _rhs30 _dafny.Int = (_302_v33).Cardinality()
			_ = _rhs30
			var _rhs31 bool = (_265_v5) && (_265_v5)
			_ = _rhs31
			var _lhs16 _dafny.Array = _300_v31
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_300_v31), 0))
			_ = _lhs17
			var _lhs18 _dafny.Array = _300_v31
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_300_v31), 0))
			_ = _lhs19
			var _lhs20 *GlobalState = _272_globalState
			_ = _lhs20
			var _lhs21 _dafny.Array = _269_v9
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_269_v9), 0))
			_ = _lhs22
			_299_v30 = _rhs27
			(_lhs16).ArraySet1(_rhs28, (_lhs17).Int())
			(_lhs18).ArraySet1(_rhs29, (_lhs19).Int())
			_lhs20.F6 = _rhs30
			(_lhs21).ArraySet1(_rhs31, (_lhs22).Int())
			var _303_v34 *C2
			_ = _303_v34
			var _nw54 *C2 = New_C2_()
			_ = _nw54
			_nw54.Ctor__()
			_303_v34 = _nw54
			var _304_v35 D3
			_ = _304_v35
			_304_v35 = Companion_D3_.Create_DC9_(_299_v30.F27)
			var _305_v36 _dafny.Map
			_ = _305_v36
			_305_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_299_v30.F27), _304_v35)
			_305_v36 = (_305_v36).Update(_295_cf2, _304_v35)
		}
		var _306_v37 _dafny.CodePoint
		_ = _306_v37
		_306_v37 = _dafny.CodePoint('p')
		var _307_v38 T0
		_ = _307_v38
		var _nw55 *C1 = New_C1_()
		_ = _nw55
		_nw55.Ctor__(_265_v5, _265_v5, _306_v37)
		_307_v38 = _nw55
		var _308_v39 _dafny.Map
		_ = _308_v39
		_308_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v9, _dafny.IntOfInt64(-647))
		var _309_v40 *C2
		_ = _309_v40
		var _nw56 *C2 = New_C2_()
		_ = _nw56
		_nw56.Ctor__()
		_309_v40 = _nw56
		var _310_v41 _dafny.Map
		_ = _310_v41
		_310_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_309_v40, _308_v39)
		var _311_v42 _dafny.Map
		_ = _311_v42
		_311_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_307_v38, (func() _dafny.Map {
			if _265_v5 {
				return _308_v39
			}
			return (func() _dafny.Map {
				if (_310_v41).Contains(_309_v40) {
					return (_310_v41).Get(_309_v40).(_dafny.Map)
				}
				return _308_v39
			})()
		})())
		_311_v42 = (_311_v42).Update(_307_v38, _308_v39)
		var _312_v43 _dafny.MultiSet
		_ = _312_v43
		_312_v43 = _dafny.MultiSetOf(_260_v1)
		var _313_v44 _dafny.Sequence
		_ = _313_v44
		_313_v44 = _dafny.SeqOf(true)
		var _314_v45 _dafny.Map
		_ = _314_v45
		_314_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_259_v0).Cardinality()), _313_v44)
		var _315_v46 D0
		_ = _315_v46
		_315_v46 = Companion_D0_.Create_DC1_(_314_v45)
		var _316_v47 _dafny.Sequence
		_ = _316_v47
		_316_v47 = _dafny.SeqOf(_dafny.MultiSetOf(_260_v1), _274_v12)
		var _317_v48 _dafny.Array
		_ = _317_v48
		var _nwElement0_9 _dafny.MultiSet = (_274_v12).Difference(Companion_Default___.Fm10(_272_globalState))
		_ = _nwElement0_9
		var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(29))
		_ = _nw57
		(_nw57).ArraySet1(_nwElement0_9, 0)
		(_nw57).ArraySet1((_274_v12).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_276_i2))), 1)
		(_nw57).ArraySet1((func() _dafny.MultiSet {
			if _265_v5 {
				return _274_v12
			}
			return (_312_v43)
		})(), 2)
		(_nw57).ArraySet1(_274_v12, 3)
		(_nw57).ArraySet1(_274_v12, 4)
		(_nw57).ArraySet1(_dafny.MultiSetFromSeq(_263_v3), 5)
		(_nw57).ArraySet1(_274_v12, 6)
		(_nw57).ArraySet1(_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_260_v1)), _276_i2), 7)
		(_nw57).ArraySet1(_274_v12, 8)
		(_nw57).ArraySet1(_dafny.MultiSetOf(Companion_Default___.Fm2(_dafny.IntOfUint32((_313_v44).Cardinality()), _315_v46, !(Companion_Default___.Fm6(_272_globalState)), _265_v5, _272_globalState)), 9)
		(_nw57).ArraySet1((_274_v12).Difference(_dafny.MultiSetFromSeq(_263_v3)), 10)
		(_nw57).ArraySet1(_dafny.MultiSetOf(_260_v1, _dafny.IntOfInt64(91), _276_i2), 11)
		(_nw57).ArraySet1(_274_v12, 12)
		(_nw57).ArraySet1(_274_v12, 13)
		(_nw57).ArraySet1(_274_v12, 14)
		(_nw57).ArraySet1((_274_v12).Intersection(_274_v12), 15)
		(_nw57).ArraySet1(Companion_Default___.Fm10(_272_globalState), 16)
		(_nw57).ArraySet1(Companion_Default___.Fm10(_272_globalState), 17)
		(_nw57).ArraySet1(_274_v12, 18)
		(_nw57).ArraySet1((_274_v12).Intersection(_274_v12), 19)
		(_nw57).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_260_v1, _dafny.IntOfUint32((_dafny.SeqOf(_265_v5)).Cardinality()))), 20)
		(_nw57).ArraySet1((_274_v12).Difference(_274_v12), 21)
		(_nw57).ArraySet1(_274_v12, 22)
		(_nw57).ArraySet1(_274_v12, 23)
		(_nw57).ArraySet1(_274_v12, 24)
		(_nw57).ArraySet1((_274_v12).Difference((_316_v47).Select((Companion_Default___.SafeIndex(_276_i2, _dafny.IntOfUint32((_316_v47).Cardinality()))).Uint32()).(_dafny.MultiSet)), 25)
		(_nw57).ArraySet1((_274_v12).Update(_dafny.IntOfInt64(838), Companion_Default___.Abs(_276_i2)), 26)
		(_nw57).ArraySet1(_dafny.MultiSetOf(_260_v1, _260_v1, _276_i2, _260_v1), 27)
		(_nw57).ArraySet1(_274_v12, 28)
		_317_v48 = _nw57
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_317_v48), 0))
		_ = _index32
		(_317_v48).ArraySet1((_dafny.MultiSetFromSeq(_263_v3)).Union(_274_v12), (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_317_v48), 0))
		_ = _index33
		(_317_v48).ArraySet1(((_274_v12).Intersection(_274_v12)).Difference(_274_v12), (_index33).Int())
		if _265_v5 {
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))
			_ = _index34
			(_269_v9).ArraySet1(_265_v5, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))
			_ = _index35
			var _rhs32 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_263_v3).Cardinality()), (_260_v1).Plus(_276_i2))
			_ = _rhs32
			var _rhs33 bool = (_265_v5) || (_265_v5)
			_ = _rhs33
			var _lhs23 _dafny.Array = _269_v9
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))
			_ = _lhs24
			_260_v1 = _rhs32
			(_lhs23).ArraySet1(_rhs33, (_lhs24).Int())
			var _318_v49 _dafny.Array
			_ = _318_v49
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_8
			var _nw58 _dafny.Array
			_ = _nw58
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw58 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = (func(_319_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_320_i3 _dafny.Int) _dafny.Sequence {
						return _319_v0
					}
				})(_259_v0)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw58 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw58).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw58).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_318_v49 = _nw58
			var _321_v50 bool
			_ = _321_v50
			var _322_v51 bool
			_ = _322_v51
			var _out3 bool
			_ = _out3
			var _out4 bool
			_ = _out4
			_out3, _out4 = (_309_v40).M0(_318_v49, _260_v1, (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool), _272_globalState)
			_321_v50 = _out3
			_322_v51 = _out4
			var _323_v52 *C1
			_ = _323_v52
			var _nw59 *C1 = New_C1_()
			_ = _nw59
			_nw59.Ctor__((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool), _265_v5, _307_v38.F24())
			_323_v52 = _nw59
			var _324_v53 _dafny.Map
			_ = _324_v53
			_324_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), _323_v52)
			var _325_v55 _dafny.Map
			_ = _325_v55
			_325_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v7, _322_v51)
			var _326_v56 _dafny.Map
			_ = _326_v56
			_326_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_276_i2, _265_v5)
			var _327_v57 _dafny.Sequence
			_ = _327_v57
			_327_v57 = _dafny.SeqOf(_259_v0, _259_v0, _259_v0)
			var _328_v58 *C0
			_ = _328_v58
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__((_323_v52).F26(), _268_v8, _307_v38.F24())
			_328_v58 = _nw60
			var _329_v59 _dafny.MultiSet
			_ = _329_v59
			_329_v59 = _dafny.MultiSetOf(_328_v58, _328_v58, _328_v58, _328_v58)
			var _330_v60 _dafny.Set
			_ = _330_v60
			_330_v60 = _dafny.SetOf(_260_v1)
			var _331_v61 _dafny.Array
			_ = _331_v61
			var _nwElement0_10 _dafny.Int = _276_i2
			_ = _nwElement0_10
			var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(23))
			_ = _nw61
			(_nw61).ArraySet1(_nwElement0_10, 0)
			(_nw61).ArraySet1(_276_i2, 1)
			(_nw61).ArraySet1((_dafny.Zero).Minus(_276_i2), 2)
			(_nw61).ArraySet1((_260_v1).Minus(_260_v1), 3)
			(_nw61).ArraySet1(((_324_v53).Update(_307_v38.F24(), _323_v52)).Cardinality(), 4)
			(_nw61).ArraySet1(_260_v1, 5)
			(_nw61).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_263_v3, _263_v3)).Cardinality()), 6)
			(_nw61).ArraySet1(_260_v1, 7)
			(_nw61).ArraySet1(_260_v1, 8)
			(_nw61).ArraySet1(_276_i2, 9)
			(_nw61).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality()), 10)
			(_nw61).ArraySet1(((func() _dafny.Set {
				var _coll12 = _dafny.NewBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_263_v3).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _332_v54 _dafny.Int
					_332_v54 = interface{}(_compr_12).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_263_v3, _332_v54) {
						_coll12.Add(Companion_Default___.SafeDivisionInt(_332_v54, _dafny.IntOfInt64(28)))
					}
				}
				return _coll12.ToSet()
			}()).Cardinality()).Times((_dafny.Zero).Minus(_260_v1)), 11)
			(_nw61).ArraySet1(_276_i2, 12)
			(_nw61).ArraySet1((_276_i2).Times(_276_i2), 13)
			(_nw61).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (func() bool {
					if (_325_v55).Contains(_dafny.SetOf(false, !((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool)), (_323_v52).F26(), true, (_323_v52).F25())) {
						return (_325_v55).Get(_dafny.SetOf(false, !((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool)), (_323_v52).F26(), true, (_323_v52).F25())).(bool)
					}
					return (func() bool {
						if (_326_v56).Contains((_dafny.MultiSetFromSeq(_327_v57)).Cardinality()) {
							return (_326_v56).Get((_dafny.MultiSetFromSeq(_327_v57)).Cardinality()).(bool)
						}
						return true
					})()
				})() {
					return _260_v1
				}
				return _dafny.IntOfInt64(14)
			})()), 14)
			(_nw61).ArraySet1(_260_v1, 15)
			(_nw61).ArraySet1(_260_v1, 16)
			(_nw61).ArraySet1(_260_v1, 17)
			(_nw61).ArraySet1(_260_v1, 18)
			(_nw61).ArraySet1(_260_v1, 19)
			(_nw61).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_329_v59).Cardinality(), _276_i2)), 20)
			(_nw61).ArraySet1(Companion_Default___.SafeModuloInt(_276_i2, (_330_v60).Cardinality()), 21)
			(_nw61).ArraySet1(_dafny.IntOfUint32((_327_v57).Cardinality()), 22)
			_331_v61 = _nw61
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_331_v61), 0))
			_ = _index36
			(_331_v61).ArraySet1(_276_i2, (_index36).Int())
			var _333_v62 _dafny.MultiSet
			_ = _333_v62
			_333_v62 = _dafny.MultiSetOf(!((_323_v52).F25()))
			var _334_v63 _dafny.Map
			_ = _334_v63
			_334_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
				if (_323_v52).F26() {
					return _333_v62
				}
				return _333_v62
			})(), _328_v58)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_331_v61), 0))
			_ = _index37
			var _rhs34 _dafny.Int = Companion_Default___.SafeModuloInt(_260_v1, ((_333_v62).Update((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool), Companion_Default___.Abs(Companion_Default___.Fm2((_330_v60).Cardinality(), _315_v46, (_323_v52).F26(), _328_v58.F27, _272_globalState)))).Cardinality())
			_ = _rhs34
			var _rhs35 bool = _322_v51
			_ = _rhs35
			var _rhs36 bool = _328_v58.F27
			_ = _rhs36
			var _rhs37 *C0 = (func() *C0 {
				if (_334_v63).Contains(_333_v62) {
					return (_334_v63).Get(_333_v62).(*C0)
				}
				return _328_v58
			})()
			_ = _rhs37
			var _lhs25 _dafny.Array = _331_v61
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_331_v61), 0))
			_ = _lhs26
			(_lhs25).ArraySet1(_rhs34, (_lhs26).Int())
			_321_v50 = _rhs35
			_321_v50 = _rhs36
			_328_v58 = _rhs37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))
			_ = _index38
			(_269_v9).ArraySet1((_276_i2).Cmp(Companion_Default___.Fm8((_263_v3).Select((Companion_Default___.SafeIndex((_266_v6).Cardinality(), _dafny.IntOfUint32((_263_v3).Cardinality()))).Uint32()).(_dafny.Int), (_331_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_331_v61), 0))).Int()).(_dafny.Int), (_323_v52).F25(), _272_globalState)) < 0, (_index38).Int())
			var _335_v64 D0
			_ = _335_v64
			_335_v64 = Companion_D0_.Create_DC2_((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool))
			var _336_v65 _dafny.Array
			_ = _336_v65
			var _nwElement0_11 bool = _328_v58.F27
			_ = _nwElement0_11
			var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(5))
			_ = _nw62
			(_nw62).ArraySet1(_nwElement0_11, 0)
			(_nw62).ArraySet1((_323_v52).F26(), 1)
			(_nw62).ArraySet1(false, 2)
			(_nw62).ArraySet1((_335_v64).Dtor_cf2(), 3)
			(_nw62).ArraySet1(Companion_Default___.Fm6(_272_globalState), 4)
			_336_v65 = _nw62
			var _337_v66 bool
			_ = _337_v66
			var _out5 bool
			_ = _out5
			_out5 = (_309_v40).M1(_336_v65, _dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool)), _313_v44), (_260_v1).Plus((_331_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_331_v61), 0))).Int()).(_dafny.Int)), _272_globalState)
			_337_v66 = _out5
		} else {
			(_272_globalState).F6 = _dafny.IntOfInt64(136)
			(_272_globalState).F6 = ((_dafny.IntOfInt64(125)).Times(_260_v1)).Plus(_276_i2)
			(_272_globalState).F1 = _dafny.Companion_Sequence_.Concatenate(_259_v0, _259_v0)
			_265_v5 = false
			var _338_v67 _dafny.MultiSet
			_ = _338_v67
			_338_v67 = _dafny.MultiSetOf(_307_v38.F24(), _307_v38.F24(), _307_v38.F24())
			var _339_v69 _dafny.Array
			_ = _339_v69
			var _340_v70 _dafny.Int
			_ = _340_v70
			var _341_v71 bool
			_ = _341_v71
			var _out6 _dafny.Array
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			var _out8 bool
			_ = _out8
			_out6, _out7, _out8 = Companion_Default___.M4((_309_v40).Fm1(_272_globalState), func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_338_v67).Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _342_v68 _dafny.CodePoint
					_342_v68 = interface{}(_compr_13).(_dafny.CodePoint)
					if (_338_v67).Contains(_342_v68) {
						_coll13.Add(_342_v68)
					}
				}
				return _coll13.ToSet()
			}(), _272_globalState)
			_339_v69 = _out6
			_340_v70 = _out7
			_341_v71 = _out8
		}
	}
	var _343_i4 _dafny.Int
	_ = _343_i4
	_343_i4 = _dafny.Zero
	{
		for (_265_v5) && (_265_v5) {
			{
				if (_343_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_343_i4 = (_343_i4).Plus(_dafny.One)
				_265_v5 = _265_v5
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_269_v9), 0))
				_ = _index39
				(_269_v9).ArraySet1(_265_v5, (_index39).Int())
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_269_v9), 0))
				_ = _index40
				(_269_v9).ArraySet1((!(_265_v5)) == (true), (_index40).Int())
				var _344_v72 _dafny.CodePoint
				_ = _344_v72
				_344_v72 = _dafny.CodePoint('p')
				var _345_v73 *C0
				_ = _345_v73
				var _nw63 *C0 = New_C0_()
				_ = _nw63
				_nw63.Ctor__(_265_v5, _268_v8, _344_v72)
				_345_v73 = _nw63
				(_272_globalState).F3 = ((_260_v1).Plus(_260_v1)).Times(_260_v1)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _346_v74 _dafny.Sequence
	_ = _346_v74
	_346_v74 = _dafny.SeqOf(_265_v5)
	var _347_v75 _dafny.Set
	_ = _347_v75
	_347_v75 = _dafny.SetOf(_346_v74)
	var _hi1 _dafny.Int = ((_347_v75).Difference(_347_v75)).Cardinality()
	_ = _hi1
	for _348_i5 := _260_v1; _348_i5.Cmp(_hi1) < 0; _348_i5 = _348_i5.Plus(_dafny.One) {
		var _349_v76 D4
		_ = _349_v76
		_349_v76 = Companion_D4_.Create_DC13_(_348_i5, _265_v5)
		_265_v5 = (_349_v76).Dtor_cf16()
		var _350_v77 _dafny.Sequence
		_ = _350_v77
		_350_v77 = _dafny.SeqOf(_268_v8)
		var _351_v78 _dafny.CodePoint
		_ = _351_v78
		_351_v78 = _dafny.CodePoint('o')
		var _352_v79 *C0
		_ = _352_v79
		var _nw64 *C0 = New_C0_()
		_ = _nw64
		_nw64.Ctor__((_348_i5).Cmp(_260_v1) >= 0, (_350_v77).Select((Companion_Default___.SafeIndex(_348_i5, _dafny.IntOfUint32((_350_v77).Cardinality()))).Uint32()).(_dafny.Array), _351_v78)
		_352_v79 = _nw64
		var _353_v80 _dafny.Array
		_ = _353_v80
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_9
		var _nw65 _dafny.Array
		_ = _nw65
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw65 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Set = (func(_354_v1 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_355_i6 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_354_v1)
				}
			})(_260_v1)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw65 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw65).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw65).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_353_v80 = _nw65
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_353_v80), 0))
		_ = _index41
		(_353_v80).ArraySet1(_dafny.SetOf(_348_i5, _260_v1), (_index41).Int())
		var _356_v81 _dafny.Set
		_ = _356_v81
		_356_v81 = _dafny.SetOf(_348_i5, _260_v1)
		var _357_v82 D7
		_ = _357_v82
		_357_v82 = Companion_D7_.Create_DC17_(_356_v81)
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_353_v80), 0))
		_ = _index42
		(_353_v80).ArraySet1((_357_v82).Dtor_cf24(), (_index42).Int())
		var _358_v83 _dafny.Array
		_ = _358_v83
		var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
		_ = _nw66
		_358_v83 = _nw66
		var _rhs38 bool = (func() bool {
			if _352_v79.F27 {
				return _352_v79.F27
			}
			return _265_v5
		})()
		_ = _rhs38
		var _rhs39 _dafny.Array = (func() _dafny.Array {
			if _265_v5 {
				return _358_v83
			}
			return _358_v83
		})()
		_ = _rhs39
		_265_v5 = _rhs38
		_358_v83 = _rhs39
	}
	var _359_v84 _dafny.Array
	_ = _359_v84
	var _nwElement0_12 _dafny.Sequence = _259_v0
	_ = _nwElement0_12
	var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(3))
	_ = _nw67
	(_nw67).ArraySet1(_nwElement0_12, 0)
	(_nw67).ArraySet1(Companion_Default___.Fm11(_272_globalState), 1)
	(_nw67).ArraySet1(_259_v0, 2)
	_359_v84 = _nw67
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_359_v84), 0))
	_ = _index43
	(_359_v84).ArraySet1(_259_v0, (_index43).Int())
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_359_v84), 0))
	_ = _index44
	(_359_v84).ArraySet1(_259_v0, (_index44).Int())
	var _360_v85 _dafny.Map
	_ = _360_v85
	_360_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v5, _dafny.SetOf(_265_v5))
	var _361_v86 _dafny.Map
	_ = _361_v86
	_361_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v1, _346_v74)
	var _362_v87 D0
	_ = _362_v87
	_362_v87 = Companion_D0_.Create_DC1_(_361_v86)
	var _363_v88 D0
	_ = _363_v88
	_363_v88 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(893))
	var _364_v89 _dafny.MultiSet
	_ = _364_v89
	_364_v89 = _dafny.MultiSetOf(_363_v88, _363_v88)
	var _365_v90 _dafny.Array
	_ = _365_v90
	var _nwElement0_13 _dafny.Int = _260_v1
	_ = _nwElement0_13
	var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(20))
	_ = _nw68
	(_nw68).ArraySet1(_nwElement0_13, 0)
	(_nw68).ArraySet1(_260_v1, 1)
	(_nw68).ArraySet1(_260_v1, 2)
	(_nw68).ArraySet1(_dafny.IntOfInt64(293), 3)
	(_nw68).ArraySet1(_260_v1, 4)
	(_nw68).ArraySet1(Companion_Default___.Fm2((_360_v85).Cardinality(), _362_v87, _265_v5, _265_v5, _272_globalState), 5)
	(_nw68).ArraySet1(_260_v1, 6)
	(_nw68).ArraySet1(_260_v1, 7)
	(_nw68).ArraySet1(_260_v1, 8)
	(_nw68).ArraySet1((func() _dafny.Int {
		if (_275_v13).Contains(false) {
			return (_275_v13).Get(false).(_dafny.Int)
		}
		return _260_v1
	})(), 9)
	(_nw68).ArraySet1(_260_v1, 10)
	(_nw68).ArraySet1(_dafny.IntOfInt64(898), 11)
	(_nw68).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_260_v1).Plus(_dafny.IntOfUint32((_346_v74).Cardinality())))), 12)
	(_nw68).ArraySet1((func() _dafny.Int {
		if _265_v5 {
			return _260_v1
		}
		return _260_v1
	})(), 13)
	(_nw68).ArraySet1(_260_v1, 14)
	(_nw68).ArraySet1(_260_v1, 15)
	(_nw68).ArraySet1((_260_v1).Times(_260_v1), 16)
	(_nw68).ArraySet1((_364_v89).Cardinality(), 17)
	(_nw68).ArraySet1(_260_v1, 18)
	(_nw68).ArraySet1((_260_v1).Times(_260_v1), 19)
	_365_v90 = _nw68
	for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_365_v90), 0))); ; {
		_guard_loop_0, _ok14 := _iter14()
		if !_ok14 {
			break
		}
		var _366_i7 _dafny.Int
		_366_i7 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_366_i7).Sign() != -1) && ((_366_i7).Cmp(_dafny.ArrayLen((_365_v90), 0)) < 0)) {
			(_365_v90).ArraySet1(Companion_Default___.SafeModuloInt(_366_i7, _260_v1), (_366_i7).Int())
		}
	}
	var _367_v91 _dafny.MultiSet
	_ = _367_v91
	_367_v91 = _dafny.MultiSetOf(false)
	var _368_v92 _dafny.Map
	_ = _368_v92
	_368_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mi"), _265_v5)
	if (((func() _dafny.Int {
		if (_275_v13).Contains(_265_v5) {
			return (_275_v13).Get(_265_v5).(_dafny.Int)
		}
		return _260_v1
	})()).Plus(_260_v1)).Cmp((func() _dafny.Int {
		if (_367_v91).Contains((func() bool {
			if (_368_v92).Contains(_259_v0) {
				return (_368_v92).Get(_259_v0).(bool)
			}
			return _265_v5
		})()) {
			return (_367_v91).Multiplicity((func() bool {
				if (_368_v92).Contains(_259_v0) {
					return (_368_v92).Get(_259_v0).(bool)
				}
				return _265_v5
			})())
		}
		return _260_v1
	})()) == 0 {
		var _369_v93 _dafny.MultiSet
		_ = _369_v93
		_369_v93 = _274_v12
		var _370_v94 _dafny.Map
		_ = _370_v94
		_370_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_369_v93, !(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(115))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_371_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_372_i8 _dafny.Int) _dafny.Sequence {
				return _371_v3
			}
		})(_263_v3))), _263_v3)))
		_370_v94 = (_370_v94).Update((func() _dafny.MultiSet {
			if _265_v5 {
				return _369_v93
			}
			return _369_v93
		})(), _265_v5)
		_265_v5 = _265_v5
		(_272_globalState).F6 = _260_v1
		var _373_v95 *C2
		_ = _373_v95
		var _nw69 *C2 = New_C2_()
		_ = _nw69
		_nw69.Ctor__()
		_373_v95 = _nw69
		var _374_v96 _dafny.CodePoint
		_ = _374_v96
		_374_v96 = _dafny.CodePoint('j')
		var _375_v97 T0
		_ = _375_v97
		var _nw70 *C1 = New_C1_()
		_ = _nw70
		_nw70.Ctor__(!(_265_v5), _265_v5, _374_v96)
		_375_v97 = _nw70
		var _376_v98 _dafny.Map
		_ = _376_v98
		_376_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_373_v95, _375_v97)
		var _377_v99 _dafny.Map
		_ = _377_v99
		_377_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_375_v97.F24(), _265_v5)
		var _378_v100 _dafny.Map
		_ = _378_v100
		_378_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v1, _260_v1)
		var _379_v101 _dafny.MultiSet
		_ = _379_v101
		_379_v101 = _dafny.MultiSetOf(_375_v97.F24(), _375_v97.F24(), _375_v97.F24())
		var _380_v102 _dafny.Set
		_ = _380_v102
		_380_v102 = _dafny.SetOf(Companion_Default___.Fm2((_376_v98).Cardinality(), _362_v87, (func() bool {
			if (_377_v99).Contains(_374_v96) {
				return (_377_v99).Get(_374_v96).(bool)
			}
			return _265_v5
		})(), _265_v5, _272_globalState), (func() _dafny.Int {
			if (_378_v100).Contains(_260_v1) {
				return (_378_v100).Get(_260_v1).(_dafny.Int)
			}
			return (_379_v101).Cardinality()
		})(), _260_v1)
		_380_v102 = _380_v102
		_373_v95 = _373_v95
	} else {
		(_272_globalState).F3 = _260_v1
		_365_v90 = _365_v90
		var _381_v103 _dafny.Array
		_ = _381_v103
		var _nw71 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
		_ = _nw71
		_381_v103 = _nw71
		var _382_v104 *C2
		_ = _382_v104
		var _nw72 *C2 = New_C2_()
		_ = _nw72
		_nw72.Ctor__()
		_382_v104 = _nw72
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_381_v103), 0))
		_ = _index45
		(_381_v103).ArraySet1(_382_v104, (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_381_v103), 0))
		_ = _index46
		var _nw73 *C2 = New_C2_()
		_ = _nw73
		_nw73.Ctor__()
		(_381_v103).ArraySet1(_nw73, (_index46).Int())
		var _383_v105 D0
		_ = _383_v105
		_383_v105 = Companion_D0_.Create_DC2_(_265_v5)
		_265_v5 = (_383_v105).Dtor_cf2()
		(_272_globalState).F0 = (_368_v92).Cardinality()
	}
	_265_v5 = ((_260_v1).Plus(_260_v1)).Cmp((_dafny.IntOfInt64(336)).Times(_dafny.IntOfInt64(881))) != 0
	var _384_i9 _dafny.Int
	_ = _384_i9
	_384_i9 = _dafny.Zero
	{
		for _265_v5 {
			{
				if (_384_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_384_i9 = (_384_i9).Plus(_dafny.One)
				var _385_v106 _dafny.CodePoint
				_ = _385_v106
				_385_v106 = _dafny.CodePoint('d')
				var _386_v107 *C0
				_ = _386_v107
				var _nw74 *C0 = New_C0_()
				_ = _nw74
				_nw74.Ctor__(_265_v5, _268_v8, _385_v106)
				_386_v107 = _nw74
				var _387_v108 _dafny.MultiSet
				_ = _387_v108
				_387_v108 = _dafny.MultiSetOf(_386_v107, _386_v107)
				var _388_v109 _dafny.Set
				_ = _388_v109
				_388_v109 = _dafny.SetOf((_259_v0).Select((Companion_Default___.SafeIndex(_260_v1, _dafny.IntOfUint32((_259_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _389_v110 _dafny.Array
				_ = _389_v110
				var _390_v111 _dafny.Int
				_ = _390_v111
				var _391_v112 bool
				_ = _391_v112
				var _out9 _dafny.Array
				_ = _out9
				var _out10 _dafny.Int
				_ = _out10
				var _out11 bool
				_ = _out11
				_out9, _out10, _out11 = Companion_Default___.M4((_387_v108).IsDisjointFrom(_387_v108), (_388_v109).Difference(_388_v109), _272_globalState)
				_389_v110 = _out9
				_390_v111 = _out10
				_391_v112 = _out11
				var _392_v113 D7
				_ = _392_v113
				_392_v113 = Companion_D7_.Create_DC19_(_346_v74)
				var _pat_let_tv8 = _346_v74
				_ = _pat_let_tv8
				var _source5 D7 = func(_pat_let3_0 D7) D7 {
					return func(_393_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let4_0 _dafny.Sequence) D7 {
							return func(_394_dt__update_hcf25_h0 _dafny.Sequence) D7 {
								return Companion_D7_.Create_DC19_(_394_dt__update_hcf25_h0)
							}(_pat_let4_0)
						}(_pat_let_tv8)
					}(_pat_let3_0)
				}(_392_v113)
				_ = _source5
				if _source5.Is_DC18() {
					var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_389_v110), 0))
					_ = _index47
					(_389_v110).ArraySet1(_390_v111, (_index47).Int())
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_389_v110), 0))
					_ = _index48
					(_389_v110).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_390_v111)), (_index48).Int())
					_259_v0 = _dafny.UnicodeSeqOfUtf8Bytes("bcjgwf")
					var _395_v114 _dafny.Map
					_ = _395_v114
					_395_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v7, (_389_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_389_v110), 0))).Int()).(_dafny.Int))
					var _396_v115 D3
					_ = _396_v115
					_396_v115 = Companion_D3_.Create_DC8_(_386_v107.F27, (_389_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_389_v110), 0))).Int()).(_dafny.Int), _260_v1)
					var _397_v116 _dafny.Array
					_ = _397_v116
					var _nwElement0_14 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_395_v114).Cardinality()), (Companion_Default___.SafeIndex(_390_v111, _dafny.IntOfUint32((_dafny.SeqOf((_395_v114).Cardinality())).Cardinality()))).Uint32(), (Companion_Default___.Fm14(_272_globalState)).Cardinality())
					_ = _nwElement0_14
					var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(5))
					_ = _nw75
					(_nw75).ArraySet1(_nwElement0_14, 0)
					(_nw75).ArraySet1(_263_v3, 1)
					(_nw75).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_263_v3, _dafny.SeqOf((_389_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_389_v110), 0))).Int()).(_dafny.Int), (_389_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_389_v110), 0))).Int()).(_dafny.Int), Companion_Default___.Fm2(_390_v111, _362_v87, (_396_v115).Dtor_cf8(), _391_v112, _272_globalState), (func() _dafny.Int {
						if (_367_v91).Contains(_386_v107.F27) {
							return (_367_v91).Multiplicity(_386_v107.F27)
						}
						return _260_v1
					})())), 2)
					(_nw75).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_263_v3, _263_v3), 3)
					(_nw75).ArraySet1(_263_v3, 4)
					_397_v116 = _nw75
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_397_v116), 0))
					_ = _index49
					(_397_v116).ArraySet1(_dafny.SeqOf((_389_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_389_v110), 0))).Int()).(_dafny.Int), _260_v1), (_index49).Int())
					var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_397_v116), 0))
					_ = _index50
					(_397_v116).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(764))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}((func(_398_v110 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_399_i10 _dafny.Int) _dafny.Int {
							return (_398_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_398_v110), 0))).Int()).(_dafny.Int)
						}
					})(_389_v110))), (_index50).Int())
					var _400_v117 _dafny.Array
					_ = _400_v117
					var _len0_10 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_10
					var _nw76 _dafny.Array
					_ = _nw76
					if _len0_10.Cmp(_dafny.Zero) == 0 {
						_nw76 = _dafny.NewArray(_len0_10)
					} else {
						var _init10 func(_dafny.Int) _dafny.Map = (func(_401_v86 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_402_i11 _dafny.Int) _dafny.Map {
								return _401_v86
							}
						})(_361_v86)
						_ = _init10
						var _element0_10 = _init10(_dafny.Zero)
						_ = _element0_10
						_nw76 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
						(_nw76).ArraySet1(_element0_10, 0)
						var _nativeLen0_10 = (_len0_10).Int()
						_ = _nativeLen0_10
						for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
							(_nw76).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
						}
					}
					_400_v117 = _nw76
					var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_400_v117), 0))
					_ = _index51
					(_400_v117).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_397_v116).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_397_v116), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.IntOfUint32(((_397_v116).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_397_v116), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int), _346_v74), (_index51).Int())
					var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_400_v117), 0))
					_ = _index52
					(_400_v117).ArraySet1(_361_v86, (_index52).Int())
				} else if _source5.Is_DC19() {
					var _403___mcc_h3 _dafny.Sequence = _source5.Get_().(D7_DC19).Cf25
					_ = _403___mcc_h3
					var _404_cf25 _dafny.Sequence = _403___mcc_h3
					_ = _404_cf25
					var _405_v118 _dafny.Array
					_ = _405_v118
					var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
					_ = _nw77
					_405_v118 = _nw77
					var _406_v119 T0
					_ = _406_v119
					var _nw78 *C1 = New_C1_()
					_ = _nw78
					_nw78.Ctor__(_386_v107.F27, _391_v112, _385_v106)
					_406_v119 = _nw78
					var _407_v120 _dafny.Array
					_ = _407_v120
					var _nwElement0_15 T0 = _406_v119
					_ = _nwElement0_15
					var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(4))
					_ = _nw79
					(_nw79).ArraySet1(_nwElement0_15, 0)
					(_nw79).ArraySet1(_406_v119, 1)
					(_nw79).ArraySet1(_406_v119, 2)
					(_nw79).ArraySet1(_406_v119, 3)
					_407_v120 = _nw79
					var _408_v121 D4
					_ = _408_v121
					_408_v121 = Companion_D4_.Create_DC11_(_407_v120)
					var _409_v122 _dafny.Map
					_ = _409_v122
					_409_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(570), _265_v5)
					var _410_v123 _dafny.Map
					_ = _410_v123
					_410_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v121, (func() bool {
						if (_409_v122).Contains(_260_v1) {
							return (_409_v122).Get(_260_v1).(bool)
						}
						return _386_v107.F27
					})())
					var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_405_v118), 0))
					_ = _index53
					(_405_v118).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_263_v3, _410_v123), (_index53).Int())
					var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v120), 0))
					_ = _index54
					(_407_v120).ArraySet1(_406_v119, (_index54).Int())
					var _411_v124 _dafny.Map
					_ = _411_v124
					_411_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_263_v3, _410_v123)
					var _412_v125 D8
					_ = _412_v125
					_412_v125 = Companion_D8_.Create_DC20_(_411_v124)
					var _413_v126 D3
					_ = _413_v126
					_413_v126 = Companion_D3_.Create_DC10_(_406_v119, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("weesf")).Cardinality()))
					var _414_v127 _dafny.Sequence
					_ = _414_v127
					_414_v127 = _dafny.SeqOf(_406_v119)
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_405_v118), 0))
					_ = _index55
					var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v120), 0))
					_ = _index56
					var _rhs40 _dafny.Map = (_412_v125).Dtor_cf26()
					_ = _rhs40
					var _rhs41 T0 = (func() T0 {
						if _391_v112 {
							return (_413_v126).Dtor_cf12()
						}
						return (_414_v127).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.IntOfUint32((_414_v127).Cardinality()))).Uint32()).(T0)
					})()
					_ = _rhs41
					var _lhs27 _dafny.Array = _405_v118
					_ = _lhs27
					var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_405_v118), 0))
					_ = _lhs28
					var _lhs29 _dafny.Array = _407_v120
					_ = _lhs29
					var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v120), 0))
					_ = _lhs30
					(_lhs27).ArraySet1(_rhs40, (_lhs28).Int())
					(_lhs29).ArraySet1(_rhs41, (_lhs30).Int())
					var _415_v128 _dafny.Array
					_ = _415_v128
					var _len0_11 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_11
					var _nw80 _dafny.Array
					_ = _nw80
					if _len0_11.Cmp(_dafny.Zero) == 0 {
						_nw80 = _dafny.NewArray(_len0_11)
					} else {
						var _init11 func(_dafny.Int) _dafny.Sequence = (func(_416_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_417_i12 _dafny.Int) _dafny.Sequence {
								return _416_v3
							}
						})(_263_v3)
						_ = _init11
						var _element0_11 = _init11(_dafny.Zero)
						_ = _element0_11
						_nw80 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
						(_nw80).ArraySet1(_element0_11, 0)
						var _nativeLen0_11 = (_len0_11).Int()
						_ = _nativeLen0_11
						for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
							(_nw80).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
						}
					}
					_415_v128 = _nw80
					var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_415_v128), 0))
					_ = _index57
					(_415_v128).ArraySet1(_263_v3, (_index57).Int())
					var _418_v129 _dafny.Array
					_ = _418_v129
					var _nw81 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
					_ = _nw81
					_418_v129 = _nw81
					var _419_v130 _dafny.MultiSet
					_ = _419_v130
					_419_v130 = _dafny.MultiSetOf(_418_v129, _418_v129)
					var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_415_v128), 0))
					_ = _index58
					(_415_v128).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_390_v111), _dafny.SeqOf((_419_v130).Cardinality())), (_index58).Int())
					var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_270_v10), 0))
					_ = _index59
					(_270_v10).ArraySet1(_269_v9, (_index59).Int())
					var _420_v131 _dafny.Sequence
					_ = _420_v131
					_420_v131 = _dafny.SeqOf(_259_v0)
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_270_v10), 0))
					_ = _index60
					var _rhs42 _dafny.Array = _269_v9
					_ = _rhs42
					var _rhs43 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.Fm2(_260_v1, _362_v87, _386_v107.F27, _386_v107.F27, _272_globalState))).Minus(_dafny.IntOfUint32(((_420_v131).Select((Companion_Default___.SafeIndex(_260_v1, _dafny.IntOfUint32((_420_v131).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
					_ = _rhs43
					var _lhs31 _dafny.Array = _270_v10
					_ = _lhs31
					var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_270_v10), 0))
					_ = _lhs32
					var _lhs33 *GlobalState = _272_globalState
					_ = _lhs33
					(_lhs31).ArraySet1(_rhs42, (_lhs32).Int())
					_lhs33.F0 = _rhs43
					var _421_v132 *C0
					_ = _421_v132
					var _nw82 *C0 = New_C0_()
					_ = _nw82
					_nw82.Ctor__(_391_v112, _268_v8, _406_v119.F24())
					_421_v132 = _nw82
				} else {
					var _422___mcc_h4 _dafny.Set = _source5.Get_().(D7_DC17).Cf24
					_ = _422___mcc_h4
					var _423_cf24 _dafny.Set = _422___mcc_h4
					_ = _423_cf24
					var _424_v133 T0
					_ = _424_v133
					var _nw83 *C1 = New_C1_()
					_ = _nw83
					_nw83.Ctor__(_265_v5, true, _385_v106)
					_424_v133 = _nw83
					var _425_v134 D3
					_ = _425_v134
					_425_v134 = Companion_D3_.Create_DC10_(_424_v133, _260_v1)
					var _426_v135 _dafny.Map
					_ = _426_v135
					_426_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
						if (_274_v12).Contains((_425_v134).Dtor_cf13()) {
							return (_274_v12).Multiplicity((_425_v134).Dtor_cf13())
						}
						return _260_v1
					})()).Times(_260_v1), _423_cf24)
					_426_v135 = (_426_v135).Update((_dafny.IntOfInt64(-188)).Plus(_260_v1), _423_cf24)
					var _427_v136 D8
					_ = _427_v136
					_427_v136 = Companion_D8_.Create_DC21_(_260_v1)
					(_272_globalState).F0 = (_427_v136).Dtor_cf27()
					var _428_v137 D0
					_ = _428_v137
					_428_v137 = Companion_D0_.Create_DC2_(_386_v107.F27)
					var _429_v138 _dafny.Sequence
					_ = _429_v138
					_429_v138 = _dafny.SeqOf(_428_v137)
					var _430_v139 _dafny.Map
					_ = _430_v139
					_430_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_429_v138, _265_v5)
					var _431_v140 *C1
					_ = _431_v140
					var _nw84 *C1 = New_C1_()
					_ = _nw84
					_nw84.Ctor__(false, _386_v107.F27, _424_v133.F24())
					_431_v140 = _nw84
					var _432_v141 _dafny.Map
					_ = _432_v141
					_432_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_386_v107).Fm7(_430_v139, _386_v107.F27, _272_globalState), _431_v140)
					_432_v141 = (_432_v141).Update((func() bool {
						if false {
							return _386_v107.F27
						}
						return _386_v107.F27
					})(), _431_v140)
					_391_v112 = (func() bool {
						if (_431_v140).F25() {
							return (_431_v140).F26()
						}
						return _391_v112
					})()
				}
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_269_v9), 0))
				_ = _index61
				(_269_v9).ArraySet1(_265_v5, (_index61).Int())
				var _433_v142 _dafny.Sequence
				_ = _433_v142
				_433_v142 = _dafny.SeqOf((func() *C0 {
					if _391_v112 {
						return _386_v107
					}
					return _386_v107
				})())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_269_v9), 0))
				_ = _index62
				var _rhs44 _dafny.Sequence = _259_v0
				_ = _rhs44
				var _rhs45 *C0 = (_433_v142).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(-876)).Plus((_dafny.Zero).Minus(_260_v1)), _dafny.IntOfUint32((_433_v142).Cardinality()))).Uint32()).(*C0)
				_ = _rhs45
				var _rhs46 bool = !((Companion_D0_.Create_DC2_(_391_v112)).Dtor_cf2())
				_ = _rhs46
				var _lhs34 *GlobalState = _272_globalState
				_ = _lhs34
				var _lhs35 _dafny.Array = _269_v9
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_269_v9), 0))
				_ = _lhs36
				_lhs34.F1 = _rhs44
				_386_v107 = _rhs45
				(_lhs35).ArraySet1(_rhs46, (_lhs36).Int())
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_263_v3, _265_v5)).Contains(_263_v3) {
					var _434_v143 _dafny.Set
					_ = _434_v143
					_434_v143 = _dafny.SetOf(_390_v111)
					var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_389_v110), 0))
					_ = _index63
					(_389_v110).ArraySet1(((_434_v143).Cardinality()).Minus(_260_v1), (_index63).Int())
					var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_389_v110), 0))
					_ = _index64
					(_389_v110).ArraySet1(_390_v111, (_index64).Int())
					var _435_v144 _dafny.Map
					_ = _435_v144
					_435_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_389_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_389_v110), 0))).Int()).(_dafny.Int), _386_v107.F27)
					var _436_v145 _dafny.Set
					_ = _436_v145
					_436_v145 = _dafny.SetOf(_435_v144)
					var _437_v146 _dafny.Map
					_ = _437_v146
					_437_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_436_v145, Companion_Default___.SafeDivisionInt(_390_v111, _390_v111))
					_437_v146 = (_437_v146).Update(_436_v145, Companion_Default___.SafeDivisionInt(_390_v111, (_389_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_389_v110), 0))).Int()).(_dafny.Int)))
					var _438_v147 *C0
					_ = _438_v147
					var _nw85 *C0 = New_C0_()
					_ = _nw85
					_nw85.Ctor__(_386_v107.F27, (_386_v107).F28(), _385_v106)
					_438_v147 = _nw85
					var _439_v148 _dafny.CodePoint
					_ = _439_v148
					var _440_v149 _dafny.Int
					_ = _440_v149
					var _441_v150 _dafny.Int
					_ = _441_v150
					var _out12 _dafny.CodePoint
					_ = _out12
					var _out13 _dafny.Int
					_ = _out13
					var _out14 _dafny.Int
					_ = _out14
					_out12, _out13, _out14 = (_386_v107).M3(_dafny.IntOfInt64(943), _385_v106, _265_v5, _272_globalState)
					_439_v148 = _out12
					_440_v149 = _out13
					_441_v150 = _out14
					var _442_v151 _dafny.Array
					_ = _442_v151
					var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
					_ = _nw86
					_442_v151 = _nw86
					var _443_v152 *C1
					_ = _443_v152
					var _nw87 *C1 = New_C1_()
					_ = _nw87
					_nw87.Ctor__((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool), (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool), _385_v106)
					_443_v152 = _nw87
					var _444_v153 _dafny.Map
					_ = _444_v153
					_444_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_440_v149, _443_v152)
					var _445_v154 _dafny.Sequence
					_ = _445_v154
					_445_v154 = _dafny.SeqOf(_444_v153)
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_442_v151), 0))
					_ = _index65
					(_442_v151).ArraySet1(_445_v154, (_index65).Int())
					var _446_v155 D9
					_ = _446_v155
					_446_v155 = Companion_D9_.Create_DC22_(_445_v154)
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_442_v151), 0))
					_ = _index66
					(_442_v151).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_445_v154, (_446_v155).Dtor_cf28()), (_index66).Int())
				} else {
					var _447_v156 _dafny.Sequence
					_ = _447_v156
					_447_v156 = _dafny.SeqOf(_274_v12)
					_391_v112 = ((_447_v156).Select((Companion_Default___.SafeIndex(_260_v1, _dafny.IntOfUint32((_447_v156).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsSubsetOf(_dafny.MultiSetOf(_260_v1, _dafny.IntOfUint32((_263_v3).Cardinality()), _dafny.IntOfInt64(153), _390_v111))
					_391_v112 = (_390_v111).Cmp(_390_v111) >= 0
					var _448_v157 _dafny.Map
					_ = _448_v157
					_448_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_390_v111, _390_v111)
					(_272_globalState).F6 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm8((func() _dafny.Int {
						if (_274_v12).Contains(_dafny.IntOfInt64(972)) {
							return (_274_v12).Multiplicity(_dafny.IntOfInt64(972))
						}
						return _260_v1
					})(), _dafny.IntOfInt64(444), _386_v107.F27, _272_globalState), (func() _dafny.Int {
						if (_448_v157).Contains(_260_v1) {
							return (_448_v157).Get(_260_v1).(_dafny.Int)
						}
						return _390_v111
					})())
					var _449_v158 _dafny.Map
					_ = _449_v158
					_449_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_260_v1, Companion_D0_.Create_DC1_(_361_v86), true, _391_v112, _272_globalState), _265_v5)
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_389_v110), 0))
					_ = _index67
					(_389_v110).ArraySet1((_390_v111).Minus((func() _dafny.Set {
						var _coll14 = _dafny.NewBuilder()
						_ = _coll14
						for _iter15 := _dafny.Iterate((_449_v158).Keys().Elements()); ; {
							_compr_14, _ok15 := _iter15()
							if !_ok15 {
								break
							}
							var _450_v159 _dafny.Int
							_450_v159 = interface{}(_compr_14).(_dafny.Int)
							if (_449_v158).Contains(_450_v159) {
								_coll14.Add(Companion_Default___.SafeDivisionInt(_450_v159, _dafny.IntOfInt64(618)))
							}
						}
						return _coll14.ToSet()
					}()).Cardinality()), (_index67).Int())
					var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_389_v110), 0))
					_ = _index68
					(_389_v110).ArraySet1((_dafny.Zero).Minus(_260_v1), (_index68).Int())
					var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_269_v9), 0))
					_ = _index69
					(_269_v9).ArraySet1((_386_v107.F27) || ((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool)), (_index69).Int())
				}
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _451_v160 _dafny.CodePoint
	_ = _451_v160
	_451_v160 = _dafny.CodePoint('s')
	var _452_v161 _dafny.Array
	_ = _452_v161
	var _453_v162 _dafny.Int
	_ = _453_v162
	var _454_v163 bool
	_ = _454_v163
	var _out15 _dafny.Array
	_ = _out15
	var _out16 _dafny.Int
	_ = _out16
	var _out17 bool
	_ = _out17
	_out15, _out16, _out17 = Companion_Default___.M4(Companion_Default___.Fm6(_272_globalState), _dafny.SetOf(_451_v160, _451_v160), _272_globalState)
	_452_v161 = _out15
	_453_v162 = _out16
	_454_v163 = _out17
	var _455_v164 *C0
	_ = _455_v164
	var _nw88 *C0 = New_C0_()
	_ = _nw88
	_nw88.Ctor__((_dafny.SetOf(_265_v5, _454_v163)).IsProperSubsetOf(_267_v7), _268_v8, _451_v160)
	_455_v164 = _nw88
	var _hi2 _dafny.Int = _260_v1
	_ = _hi2
	for _456_i13 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("havjj"), (Companion_Default___.SafeIndex(_453_v162, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("havjj")).Cardinality()))).Uint32(), (_259_v0).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_265_v5)).Cardinality(), _dafny.IntOfUint32((_259_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()); _456_i13.Cmp(_hi2) < 0; _456_i13 = _456_i13.Plus(_dafny.One) {
		var _457_v165 _dafny.Set
		_ = _457_v165
		_457_v165 = _dafny.SetOf(_453_v162)
		var _458_v166 _dafny.Map
		_ = _458_v166
		_458_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_457_v165).Union(_457_v165), _266_v6)
		_458_v166 = (_458_v166).Update((_dafny.SetOf(_456_i13)).Difference(func() _dafny.Set {
			var _coll15 = _dafny.NewBuilder()
			_ = _coll15
			for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(199), _dafny.IntOfInt64(506))); ; {
				_compr_15, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _459_v167 _dafny.Int
				_459_v167 = interface{}(_compr_15).(_dafny.Int)
				if ((_dafny.IntOfInt64(199)).Cmp(_459_v167) <= 0) && ((_459_v167).Cmp(_dafny.IntOfInt64(506)) < 0) {
					_coll15.Add(Companion_Default___.SafeModuloInt(_459_v167, _456_i13))
				}
			}
			return _coll15.ToSet()
		}()), _266_v6)
		var _460_v168 _dafny.Sequence
		_ = _460_v168
		_460_v168 = _dafny.SeqOf((_455_v164).F28(), (_455_v164).F28())
		var _461_v169 _dafny.Map
		_ = _461_v169
		_461_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_453_v162, (_dafny.Zero).Minus(_456_i13))
		var _462_v170 *C0
		_ = _462_v170
		var _nw89 *C0 = New_C0_()
		_ = _nw89
		_nw89.Ctor__((_367_v91).IsProperSubsetOf(_dafny.MultiSetOf(_455_v164.F27)), (_460_v168).Select((Companion_Default___.SafeIndex((_461_v169).Cardinality(), _dafny.IntOfUint32((_460_v168).Cardinality()))).Uint32()).(_dafny.Array), _451_v160)
		_462_v170 = _nw89
		_461_v169 = (_461_v169).Update(_453_v162, _453_v162)
		_267_v7 = ((_dafny.SetOf(_455_v164.F27, _455_v164.F27, _455_v164.F27, _455_v164.F27)).Intersection(_267_v7)).Intersection(_267_v7)
	}
	(_272_globalState).F1 = _259_v0
	var _463_v171 D3
	_ = _463_v171
	_463_v171 = Companion_D3_.Create_DC9_(false)
	var _source6 D3 = (func() D3 {
		if _454_v163 {
			return _463_v171
		}
		return _463_v171
	})()
	_ = _source6
	if _source6.Is_DC8() {
		var _464___mcc_h5 bool = _source6.Get_().(D3_DC8).Cf8
		_ = _464___mcc_h5
		var _465___mcc_h6 _dafny.Int = _source6.Get_().(D3_DC8).Cf9
		_ = _465___mcc_h6
		var _466___mcc_h7 _dafny.Int = _source6.Get_().(D3_DC8).Cf10
		_ = _466___mcc_h7
		var _467_cf10 _dafny.Int = _466___mcc_h7
		_ = _467_cf10
		var _468_cf9 _dafny.Int = _465___mcc_h6
		_ = _468_cf9
		var _469_cf8 bool = _464___mcc_h5
		_ = _469_cf8
		(_272_globalState).F22 = _269_v9
		var _470_v172 D0
		_ = _470_v172
		_470_v172 = Companion_D0_.Create_DC2_(_454_v163)
		var _471_v173 *C2
		_ = _471_v173
		var _nw90 *C2 = New_C2_()
		_ = _nw90
		_nw90.Ctor__()
		_471_v173 = _nw90
		var _472_v174 _dafny.Map
		_ = _472_v174
		_472_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v164.F27, _471_v173)
		var _473_v175 _dafny.Set
		_ = _473_v175
		_473_v175 = _dafny.SetOf((_472_v174).Cardinality(), _453_v162)
		var _474_v176 _dafny.Sequence
		_ = _474_v176
		_474_v176 = _dafny.SeqOf(_470_v172, Companion_Default___.Fm19(_455_v164.F27, _473_v175, _272_globalState), _470_v172, _470_v172)
		var _475_v177 _dafny.Map
		_ = _475_v177
		_475_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_474_v176, _265_v5)
		var _476_v178 T0
		_ = _476_v178
		var _nw91 *C1 = New_C1_()
		_ = _nw91
		_nw91.Ctor__(_455_v164.F27, (_455_v164).Fm7(_475_v177, _455_v164.F27, _272_globalState), _451_v160)
		_476_v178 = _nw91
		_476_v178 = _476_v178
		var _477_v179 _dafny.Map
		_ = _477_v179
		_477_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_471_v173, _471_v173, _471_v173, _471_v173)).Contains(_471_v173), _dafny.SeqOf(_454_v163, _455_v164.F27, _469_cf8))
		_477_v179 = (_477_v179).Update(_455_v164.F27, _346_v74)
		var _478_v180 _dafny.Map
		_ = _478_v180
		_478_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-172))).Cardinality()))
		(_476_v178).F24_set_((_259_v0).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_478_v180).Contains(_453_v162) {
				return (_478_v180).Get(_453_v162).(_dafny.Int)
			}
			return _467_cf10
		})(), _dafny.IntOfUint32((_259_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
	} else if _source6.Is_DC9() {
		var _479___mcc_h8 bool = _source6.Get_().(D3_DC9).Cf11
		_ = _479___mcc_h8
		var _480_cf11 bool = _479___mcc_h8
		_ = _480_cf11
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_12
		var _nw92 _dafny.Array
		_ = _nw92
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw92 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) bool = (func(_481_v5 bool) func(_dafny.Int) bool {
				return func(_482_i14 _dafny.Int) bool {
					return _481_v5
				}
			})(_265_v5)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw92 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw92).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw92).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		(_272_globalState).F22 = _nw92
		var _483_v182 *C1
		_ = _483_v182
		var _nw93 *C1 = New_C1_()
		_ = _nw93
		_nw93.Ctor__(!(_480_cf11), false, _451_v160)
		_483_v182 = _nw93
		var _484_v183 _dafny.Map
		_ = _484_v183
		_484_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(955), _dafny.IntOfInt64(499))); ; {
				_compr_16, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _485_v181 _dafny.Int
				_485_v181 = interface{}(_compr_16).(_dafny.Int)
				if ((_dafny.IntOfInt64(955)).Cmp(_485_v181) <= 0) && ((_485_v181).Cmp(_dafny.IntOfInt64(499)) < 0) {
					_coll16.Add((_485_v181).Minus(_260_v1), _453_v162)
				}
			}
			return _coll16.ToMap()
		}()).Cardinality(), _483_v182)
		var _486_v184 _dafny.Sequence
		_ = _486_v184
		_486_v184 = _dafny.SeqOf(_484_v183)
		var _487_v185 D9
		_ = _487_v185
		_487_v185 = Companion_D9_.Create_DC22_((func() _dafny.Sequence {
			if _454_v163 {
				return _486_v184
			}
			return _486_v184
		})())
		var _rhs47 D9 = _487_v185
		_ = _rhs47
		var _rhs48 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_263_v3, (Companion_Default___.SafeIndex(_453_v162, _dafny.IntOfUint32((_263_v3).Cardinality()))).Uint32(), _260_v1), _263_v3)).Cardinality())
		_ = _rhs48
		var _lhs37 *GlobalState = _272_globalState
		_ = _lhs37
		_487_v185 = _rhs47
		_lhs37.F3 = _rhs48
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen(((_455_v164).F28()), 0))
		_ = _index70
		((_455_v164).F28()).ArraySet1(_452_v161, (_index70).Int())
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen(((_455_v164).F28()), 0))
		_ = _index71
		((_455_v164).F28()).ArraySet1(_452_v161, (_index71).Int())
		(_272_globalState).F1 = _259_v0
	} else if _source6.Is_DC10() {
		var _488___mcc_h9 T0 = _source6.Get_().(D3_DC10).Cf12
		_ = _488___mcc_h9
		var _489___mcc_h10 _dafny.Int = _source6.Get_().(D3_DC10).Cf13
		_ = _489___mcc_h10
		var _490_cf13 _dafny.Int = _489___mcc_h10
		_ = _490_cf13
		var _491_cf12 T0 = _488___mcc_h9
		_ = _491_cf12
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))
		_ = _index72
		(_269_v9).ArraySet1(_454_v163, (_index72).Int())
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))
		_ = _index73
		(_269_v9).ArraySet1(_455_v164.F27, (_index73).Int())
		(_272_globalState).F3 = _260_v1
		(_272_globalState).F0 = _260_v1
		if (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool) {
			var _492_v186 _dafny.Array
			_ = _492_v186
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_13
			var _nw94 _dafny.Array
			_ = _nw94
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw94 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Map = (func(_493_v6 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_494_i15 _dafny.Int) _dafny.Map {
						return _493_v6
					}
				})(_266_v6)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw94 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw94).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw94).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_492_v186 = _nw94
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_492_v186), 0))
			_ = _index74
			(_492_v186).ArraySet1(_266_v6, (_index74).Int())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))
			_ = _index75
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_492_v186), 0))
			_ = _index76
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))
			_ = _index77
			var _rhs49 bool = _455_v164.F27
			_ = _rhs49
			var _rhs50 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_260_v1, (_dafny.Zero).Minus((_490_cf13).Plus(_260_v1))))
			_ = _rhs50
			var _rhs51 _dafny.Map = _266_v6
			_ = _rhs51
			var _rhs52 bool = !((_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool))
			_ = _rhs52
			var _lhs38 _dafny.Array = _269_v9
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))
			_ = _lhs39
			var _lhs40 *GlobalState = _272_globalState
			_ = _lhs40
			var _lhs41 _dafny.Array = _492_v186
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_492_v186), 0))
			_ = _lhs42
			var _lhs43 _dafny.Array = _269_v9
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))
			_ = _lhs44
			(_lhs38).ArraySet1(_rhs49, (_lhs39).Int())
			_lhs40.F0 = _rhs50
			(_lhs41).ArraySet1(_rhs51, (_lhs42).Int())
			(_lhs43).ArraySet1(_rhs52, (_lhs44).Int())
			var _495_v187 _dafny.Array
			_ = _495_v187
			var _nwElement0_16 _dafny.CodePoint = _491_cf12.F24()
			_ = _nwElement0_16
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(19))
			_ = _nw95
			(_nw95).ArraySet1CodePoint(_nwElement0_16, 0)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 1)
			(_nw95).ArraySet1CodePoint(_dafny.CodePoint('w'), 2)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 3)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 4)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 5)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 6)
			(_nw95).ArraySet1CodePoint(_dafny.CodePoint('y'), 7)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 8)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 9)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 10)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 11)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 12)
			(_nw95).ArraySet1CodePoint(_451_v160, 13)
			(_nw95).ArraySet1CodePoint(_451_v160, 14)
			(_nw95).ArraySet1CodePoint(_451_v160, 15)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 16)
			(_nw95).ArraySet1CodePoint(_491_cf12.F24(), 17)
			(_nw95).ArraySet1CodePoint(_451_v160, 18)
			_495_v187 = _nw95
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_495_v187), 0))
			_ = _index78
			(_495_v187).ArraySet1CodePoint(_451_v160, (_index78).Int())
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_495_v187), 0))
			_ = _index79
			(_495_v187).ArraySet1CodePoint(Companion_Default___.Fm13(_272_globalState), (_index79).Int())
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_495_v187), 0))
			_ = _index80
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_495_v187), 0))
			_ = _index81
			var _rhs53 _dafny.CodePoint = Companion_Default___.Fm13(_272_globalState)
			_ = _rhs53
			var _rhs54 _dafny.CodePoint = (func() _dafny.CodePoint {
				if _455_v164.F27 {
					return _dafny.CodePoint('q')
				}
				return _491_cf12.F24()
			})()
			_ = _rhs54
			var _rhs55 _dafny.CodePoint = _491_cf12.F24()
			_ = _rhs55
			var _lhs45 _dafny.Array = _495_v187
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_495_v187), 0))
			_ = _lhs46
			var _lhs47 T0 = _491_cf12
			_ = _lhs47
			var _lhs48 _dafny.Array = _495_v187
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_495_v187), 0))
			_ = _lhs49
			(_lhs45).ArraySet1CodePoint(_rhs53, (_lhs46).Int())
			_lhs47.F24_set_(_rhs54)
			(_lhs48).ArraySet1CodePoint(_rhs55, (_lhs49).Int())
			var _496_v188 D9
			_ = _496_v188
			_496_v188 = Companion_D9_.Create_DC23_(_455_v164.F27, _dafny.IntOfUint32((_346_v74).Cardinality()))
			var _497_v189 _dafny.Array
			_ = _497_v189
			var _nwElement0_17 bool = (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool)
			_ = _nwElement0_17
			var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(4))
			_ = _nw96
			(_nw96).ArraySet1(_nwElement0_17, 0)
			(_nw96).ArraySet1(true, 1)
			(_nw96).ArraySet1((_496_v188).Dtor_cf29(), 2)
			(_nw96).ArraySet1(_265_v5, 3)
			_497_v189 = _nw96
			var _498_v190 _dafny.Map
			_ = _498_v190
			_498_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_497_v189, (_346_v74).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_454_v163, _454_v163)).Cardinality(), _dafny.IntOfUint32((_346_v74).Cardinality()))).Uint32()).(bool))
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))
			_ = _index82
			(_269_v9).ArraySet1((func() bool {
				if (_498_v190).Contains(_497_v189) {
					return (_498_v190).Get(_497_v189).(bool)
				}
				return _454_v163
			})(), (_index82).Int())
			var _499_v191 *C1
			_ = _499_v191
			var _nw97 *C1 = New_C1_()
			_ = _nw97
			_nw97.Ctor__(_265_v5, (func() bool {
				if (_368_v92).Contains(_dafny.UnicodeSeqOfUtf8Bytes("mvs")) {
					return (_368_v92).Get(_dafny.UnicodeSeqOfUtf8Bytes("mvs")).(bool)
				}
				return (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool)
			})(), _491_cf12.F24())
			_499_v191 = _nw97
			_453_v162 = _453_v162
		} else {
			(_272_globalState).F6 = _490_cf13
			(_272_globalState).F1 = _dafny.Companion_Sequence_.Update(_259_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.IntOfUint32((_259_v0).Cardinality()))).Uint32(), _491_cf12.F24())
			var _500_v192 _dafny.CodePoint
			_ = _500_v192
			var _501_v193 _dafny.Int
			_ = _501_v193
			var _502_v194 _dafny.Int
			_ = _502_v194
			var _out18 _dafny.CodePoint
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			var _out20 _dafny.Int
			_ = _out20
			_out18, _out19, _out20 = (_455_v164).M3(_260_v1, _491_cf12.F24(), (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool), _272_globalState)
			_500_v192 = _out18
			_501_v193 = _out19
			_502_v194 = _out20
			(_272_globalState).F6 = ((_dafny.Zero).Minus(_502_v194)).Times(_490_cf13)
			var _503_v195 _dafny.Array
			_ = _503_v195
			var _nw98 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
			_ = _nw98
			_503_v195 = _nw98
			var _504_v196 D4
			_ = _504_v196
			_504_v196 = Companion_D4_.Create_DC11_(_503_v195)
			var _pat_let_tv9 = _503_v195
			_ = _pat_let_tv9
			var _505_v197 _dafny.Array
			_ = _505_v197
			var _nwElement0_18 D4 = _504_v196
			_ = _nwElement0_18
			var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(21))
			_ = _nw99
			(_nw99).ArraySet1(_nwElement0_18, 0)
			(_nw99).ArraySet1(_504_v196, 1)
			(_nw99).ArraySet1(_504_v196, 2)
			(_nw99).ArraySet1(_504_v196, 3)
			(_nw99).ArraySet1(_504_v196, 4)
			(_nw99).ArraySet1(_504_v196, 5)
			(_nw99).ArraySet1(_504_v196, 6)
			(_nw99).ArraySet1(_504_v196, 7)
			(_nw99).ArraySet1(Companion_D4_.Create_DC11_(_503_v195), 8)
			(_nw99).ArraySet1(func(_pat_let5_0 D4) D4 {
				return func(_506_dt__update__tmp_h1 D4) D4 {
					return func(_pat_let6_0 _dafny.Array) D4 {
						return func(_507_dt__update_hcf14_h0 _dafny.Array) D4 {
							return Companion_D4_.Create_DC11_(_507_dt__update_hcf14_h0)
						}(_pat_let6_0)
					}(_pat_let_tv9)
				}(_pat_let5_0)
			}(_504_v196), 9)
			(_nw99).ArraySet1((func() D4 {
				if (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool) {
					return _504_v196
				}
				return Companion_D4_.Create_DC11_(_503_v195)
			})(), 10)
			(_nw99).ArraySet1(_504_v196, 11)
			(_nw99).ArraySet1(_504_v196, 12)
			(_nw99).ArraySet1(_504_v196, 13)
			(_nw99).ArraySet1(_504_v196, 14)
			(_nw99).ArraySet1(_504_v196, 15)
			(_nw99).ArraySet1(Companion_D4_.Create_DC11_(_503_v195), 16)
			(_nw99).ArraySet1(_504_v196, 17)
			(_nw99).ArraySet1((func() D4 {
				if (_269_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_269_v9), 0))).Int()).(bool) {
					return _504_v196
				}
				return _504_v196
			})(), 18)
			(_nw99).ArraySet1(_504_v196, 19)
			(_nw99).ArraySet1(_504_v196, 20)
			_505_v197 = _nw99
			var _508_v198 _dafny.Map
			_ = _508_v198
			_508_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC21_(Companion_Default___.Fm8(_dafny.IntOfUint32(((_359_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_359_v84), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-929))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_509_v192 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_510_i16 _dafny.Int) _dafny.CodePoint {
					return _509_v192
				}
			})(_500_v192)))).Cardinality()), _265_v5, _272_globalState)), _505_v197)
			var _511_v199 D8
			_ = _511_v199
			_511_v199 = Companion_D8_.Create_DC21_(_dafny.IntOfUint32((_259_v0).Cardinality()))
			_505_v197 = (func() _dafny.Array {
				if (_508_v198).Contains(_511_v199) {
					return (_508_v198).Get(_511_v199).(_dafny.Array)
				}
				return _505_v197
			})()
		}
	} else {
		var _512___mcc_h11 _dafny.Array = _source6.Get_().(D3_DC7).Cf7
		_ = _512___mcc_h11
		var _513_cf7 _dafny.Array = _512___mcc_h11
		_ = _513_cf7
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))
		_ = _index83
		(_452_v161).ArraySet1(_260_v1, (_index83).Int())
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))
		_ = _index84
		(_452_v161).ArraySet1((_260_v1).Minus(_453_v162), (_index84).Int())
		var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_359_v84), 0))
		_ = _index85
		(_359_v84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_259_v0, _259_v0), (_index85).Int())
		var _514_v200 T0
		_ = _514_v200
		var _nw100 *C0 = New_C0_()
		_ = _nw100
		_nw100.Ctor__(_265_v5, (_455_v164).F28(), _451_v160)
		_514_v200 = _nw100
		var _515_v201 _dafny.Sequence
		_ = _515_v201
		_515_v201 = _dafny.SeqOf(_514_v200)
		var _516_v202 _dafny.Set
		_ = _516_v202
		_516_v202 = _dafny.SetOf(_453_v162)
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))
		_ = _index86
		(_452_v161).ArraySet1((_263_v3).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_515_v201, (Companion_Default___.SafeIndex((_516_v202).Cardinality(), _dafny.IntOfUint32((_515_v201).Cardinality()))).Uint32(), _514_v200)).Cardinality()), _453_v162), _dafny.IntOfUint32((_263_v3).Cardinality()))).Uint32()).(_dafny.Int), (_index86).Int())
		var _517_v203 *C1
		_ = _517_v203
		var _nw101 *C1 = New_C1_()
		_ = _nw101
		_nw101.Ctor__(_455_v164.F27, _454_v163, _dafny.CodePoint('y'))
		_517_v203 = _nw101
		var _518_v204 _dafny.Map
		_ = _518_v204
		_518_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_453_v162, _517_v203)
		var _519_v205 _dafny.Sequence
		_ = _519_v205
		_519_v205 = _dafny.SeqOf(_518_v204, _518_v204, _518_v204, _518_v204)
		var _520_v206 D9
		_ = _520_v206
		_520_v206 = Companion_D9_.Create_DC22_(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_519_v205, (Companion_Default___.SafeIndex(_260_v1, _dafny.IntOfUint32((_519_v205).Cardinality()))).Uint32(), _518_v204), _519_v205), (Companion_Default___.SafeIndex(_453_v162, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_519_v205, (Companion_Default___.SafeIndex(_260_v1, _dafny.IntOfUint32((_519_v205).Cardinality()))).Uint32(), _518_v204), _519_v205)).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_267_v7).Cardinality(), _517_v203)))
		var _source7 D9 = _520_v206
		_ = _source7
		if _source7.Is_DC23() {
			var _521___mcc_h12 bool = _source7.Get_().(D9_DC23).Cf29
			_ = _521___mcc_h12
			var _522___mcc_h13 _dafny.Int = _source7.Get_().(D9_DC23).Cf30
			_ = _522___mcc_h13
			var _523_cf30 _dafny.Int = _522___mcc_h13
			_ = _523_cf30
			var _524_cf29 bool = _521___mcc_h12
			_ = _524_cf29
			(_455_v164).F27 = (((_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int)).Minus((_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int))).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_453_v162))) > 0
			(_272_globalState).F9 = _dafny.CodePoint('w')
			(_272_globalState).F0 = (_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int)
			(_455_v164).F27 = _524_cf29
		} else if _source7.Is_DC24() {
			var _525___mcc_h14 _dafny.CodePoint = _source7.Get_().(D9_DC24).Cf31
			_ = _525___mcc_h14
			var _526___mcc_h15 bool = _source7.Get_().(D9_DC24).Cf32
			_ = _526___mcc_h15
			var _527___mcc_h16 bool = _source7.Get_().(D9_DC24).Cf33
			_ = _527___mcc_h16
			var _528___mcc_h17 _dafny.Map = _source7.Get_().(D9_DC24).Cf34
			_ = _528___mcc_h17
			var _529___mcc_h18 bool = _source7.Get_().(D9_DC24).Cf35
			_ = _529___mcc_h18
			var _530_cf35 bool = _529___mcc_h18
			_ = _530_cf35
			var _531_cf34 _dafny.Map = _528___mcc_h17
			_ = _531_cf34
			var _532_cf33 bool = _527___mcc_h16
			_ = _532_cf33
			var _533_cf32 bool = _526___mcc_h15
			_ = _533_cf32
			var _534_cf31 _dafny.CodePoint = _525___mcc_h14
			_ = _534_cf31
			(_455_v164).F27 = _454_v163
			var _535_v207 _dafny.Map
			_ = _535_v207
			_535_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v9, _260_v1)
			var _rhs56 _dafny.Int = Companion_Default___.SafeDivisionInt(_260_v1, (func() _dafny.Int {
				if (_535_v207).Contains(_269_v9) {
					return (_535_v207).Get(_269_v9).(_dafny.Int)
				}
				return _260_v1
			})())
			_ = _rhs56
			var _rhs57 bool = _265_v5
			_ = _rhs57
			var _rhs58 bool = (_260_v1).Cmp((_dafny.Zero).Minus((_260_v1).Minus(_260_v1))) <= 0
			_ = _rhs58
			var _lhs50 *GlobalState = _272_globalState
			_ = _lhs50
			_lhs50.F0 = _rhs56
			_533_cf32 = _rhs57
			_454_v163 = _rhs58
			var _536_v208 _dafny.Set
			_ = _536_v208
			_536_v208 = _dafny.SetOf(_514_v200.F24(), _534_cf31)
			var _537_v209 _dafny.Array
			_ = _537_v209
			var _538_v210 _dafny.Int
			_ = _538_v210
			var _539_v211 bool
			_ = _539_v211
			var _out21 _dafny.Array
			_ = _out21
			var _out22 _dafny.Int
			_ = _out22
			var _out23 bool
			_ = _out23
			_out21, _out22, _out23 = Companion_Default___.M4(((_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int)).Cmp(_453_v162) == 0, (_536_v208).Union(_536_v208), _272_globalState)
			_537_v209 = _out21
			_538_v210 = _out22
			_539_v211 = _out23
			var _540_v214 D7
			_ = _540_v214
			_540_v214 = Companion_D7_.Create_DC17_((func() _dafny.Set {
				var _coll17 = _dafny.NewBuilder()
				_ = _coll17
				for _iter18 := _dafny.Iterate((_274_v12).Elements()); ; {
					_compr_17, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _541_v212 _dafny.Int
					_541_v212 = interface{}(_compr_17).(_dafny.Int)
					if (_274_v12).Contains(_541_v212) {
						_coll17.Add((_541_v212).Minus((Companion_D2_.Create_DC5_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(129))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg20 _dafny.Int) interface{} {
								return coer20(arg20)
							}
						}(func(_542_i17 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-170))).Cardinality()
						}))).Cardinality()))).Dtor_cf5()))
					}
				}
				return _coll17.ToSet()
			}()).Difference(func() _dafny.Set {
				var _coll18 = _dafny.NewBuilder()
				_ = _coll18
				for _iter19 := _dafny.Iterate((_263_v3).Elements()); ; {
					_compr_18, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _543_v213 _dafny.Int
					_543_v213 = interface{}(_compr_18).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_263_v3, _543_v213) {
						_coll18.Add((_543_v213).Times(_dafny.IntOfInt64(969)))
					}
				}
				return _coll18.ToSet()
			}()))
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_513_cf7), 0))
			_ = _index87
			(_513_cf7).ArraySet1(false, (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_513_cf7), 0))
			_ = _index88
			var _rhs59 D7 = _540_v214
			_ = _rhs59
			var _rhs60 _dafny.Set = ((_dafny.SetOf(_539_v211)).Difference(_267_v7)).Difference(_267_v7)
			_ = _rhs60
			var _rhs61 bool = !((_516_v202).Intersection(_516_v202)).Equals(_516_v202)
			_ = _rhs61
			var _lhs51 _dafny.Array = _513_cf7
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_513_cf7), 0))
			_ = _lhs52
			_540_v214 = _rhs59
			_267_v7 = _rhs60
			(_lhs51).ArraySet1(_rhs61, (_lhs52).Int())
		} else if _source7.Is_DC22() {
			var _544___mcc_h19 _dafny.Sequence = _source7.Get_().(D9_DC22).Cf28
			_ = _544___mcc_h19
			var _545_cf28 _dafny.Sequence = _544___mcc_h19
			_ = _545_cf28
			_515_v201 = _dafny.SeqOf(_514_v200)
			var _546_v215 _dafny.Map
			_ = _546_v215
			_546_v215 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v1, _260_v1)
			var _547_v216 _dafny.CodePoint
			_ = _547_v216
			var _548_v217 _dafny.Int
			_ = _548_v217
			var _549_v218 _dafny.Int
			_ = _549_v218
			var _out24 _dafny.CodePoint
			_ = _out24
			var _out25 _dafny.Int
			_ = _out25
			var _out26 _dafny.Int
			_ = _out26
			_out24, _out25, _out26 = (_455_v164).M3((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_263_v3, Companion_Default___.Fm20(_514_v200.F24(), Companion_Default___.Fm21((_546_v215).Cardinality(), _272_globalState), _272_globalState))).Cardinality())), _514_v200.F24(), true, _272_globalState)
			_547_v216 = _out24
			_548_v217 = _out25
			_549_v218 = _out26
			_548_v217 = _453_v162
			var _550_v220 D0
			_ = _550_v220
			_550_v220 = Companion_D0_.Create_DC2_(_454_v163)
			var _551_v221 _dafny.Sequence
			_ = _551_v221
			_551_v221 = _dafny.SeqOf(_550_v220)
			var _552_v222 _dafny.Sequence
			_ = _552_v222
			_552_v222 = _dafny.SeqOf(_551_v221, _551_v221)
			var _rhs62 bool = _265_v5
			_ = _rhs62
			var _rhs63 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_346_v74, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_455_v164).Fm7(func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter20 := _dafny.Iterate((_552_v222).Elements()); ; {
						_compr_19, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _553_v219 _dafny.Sequence
						_553_v219 = interface{}(_compr_19).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_552_v222, _553_v219) {
							_coll19.Add(_553_v219, _265_v5)
						}
					}
					return _coll19.ToMap()
				}(), (_517_v203).F25(), _272_globalState) {
					return _346_v74
				}
				return _346_v74
			})(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_263_v3, (Companion_Default___.SafeIndex((_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_263_v3).Cardinality()))).Uint32(), (_dafny.Zero).Minus((_546_v215).Cardinality()))).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_455_v164).Fm7(func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter21 := _dafny.Iterate((_552_v222).Elements()); ; {
						_compr_20, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _554_v219 _dafny.Sequence
						_554_v219 = interface{}(_compr_20).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_552_v222, _554_v219) {
							_coll20.Add(_554_v219, _265_v5)
						}
					}
					return _coll20.ToMap()
				}(), (_517_v203).F25(), _272_globalState) {
					return _346_v74
				}
				return _346_v74
			})()).Cardinality()))).Uint32(), (_517_v203).F26()))
			_ = _rhs63
			var _lhs53 *C0 = _455_v164
			_ = _lhs53
			_265_v5 = _rhs62
			_lhs53.F27 = _rhs63
		} else {
			var _555___mcc_h20 D9 = _source7.Get_().(D9_DC25).Cf36
			_ = _555___mcc_h20
			var _556_cf36 D9 = _555___mcc_h20
			_ = _556_cf36
			(_272_globalState).F0 = ((_dafny.Zero).Minus((_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int))).Minus((_453_v162).Minus(_260_v1))
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_452_v161), 0))
			_ = _index89
			(_452_v161).ArraySet1(_dafny.IntOfInt64(143), (_index89).Int())
			var _557_v223 bool
			_ = _557_v223
			var _558_v224 _dafny.Array
			_ = _558_v224
			var _559_v225 bool
			_ = _559_v225
			var _560_v226 _dafny.Int
			_ = _560_v226
			var _out27 bool
			_ = _out27
			var _out28 _dafny.Array
			_ = _out28
			var _out29 bool
			_ = _out29
			var _out30 _dafny.Int
			_ = _out30
			_out27, _out28, _out29, _out30 = (_517_v203).M2((_359_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_359_v84), 0))).Int()).(_dafny.Sequence), (_455_v164).F28(), _514_v200, _272_globalState)
			_557_v223 = _out27
			_558_v224 = _out28
			_559_v225 = _out29
			_560_v226 = _out30
			_559_v225 = !(_dafny.SetOf(_559_v225, _455_v164.F27, _557_v223)).Contains(!(_455_v164.F27))
		}
	}
	if _454_v163 {
		var _561_v227 _dafny.Array
		_ = _561_v227
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_14
		var _nw102 _dafny.Array
		_ = _nw102
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw102 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Sequence = (func(_562_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_563_i18 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_562_v3, _562_v3)
				}
			})(_263_v3)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw102 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw102).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw102).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_561_v227 = _nw102
		var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_561_v227), 0))
		_ = _index90
		(_561_v227).ArraySet1(_263_v3, (_index90).Int())
		var _564_v228 _dafny.Sequence
		_ = _564_v228
		_564_v228 = _dafny.SeqOf(_269_v9)
		var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_561_v227), 0))
		_ = _index91
		var _rhs64 _dafny.Array = (_564_v228).Select((Companion_Default___.SafeIndex((func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(659), _dafny.IntOfInt64(151))); ; {
				_compr_21, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _565_v229 _dafny.Int
				_565_v229 = interface{}(_compr_21).(_dafny.Int)
				if ((_dafny.IntOfInt64(659)).Cmp(_565_v229) <= 0) && ((_565_v229).Cmp(_dafny.IntOfInt64(151)) < 0) {
					_coll21.Add(Companion_Default___.SafeDivisionInt(_565_v229, _260_v1))
				}
			}
			return _coll21.ToSet()
		}()).Cardinality(), _dafny.IntOfUint32((_564_v228).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs64
		var _rhs65 _dafny.Sequence = Companion_Default___.Fm20(_451_v160, Companion_Default___.Fm21(_453_v162, _272_globalState), _272_globalState)
		_ = _rhs65
		var _lhs54 *GlobalState = _272_globalState
		_ = _lhs54
		var _lhs55 _dafny.Array = _561_v227
		_ = _lhs55
		var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_561_v227), 0))
		_ = _lhs56
		_lhs54.F22 = _rhs64
		(_lhs55).ArraySet1(_rhs65, (_lhs56).Int())
		_260_v1 = (func() _dafny.Int {
			if _265_v5 {
				return _453_v162
			}
			return (_453_v162).Times(_453_v162)
		})()
		var _566_v230 *C1
		_ = _566_v230
		var _nw103 *C1 = New_C1_()
		_ = _nw103
		_nw103.Ctor__(_455_v164.F27, (_346_v74).Select((Companion_Default___.SafeIndex(_260_v1, _dafny.IntOfUint32((_346_v74).Cardinality()))).Uint32()).(bool), _451_v160)
		_566_v230 = _nw103
		var _567_v231 _dafny.Map
		_ = _567_v231
		_567_v231 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_453_v162, _566_v230)
		var _568_v232 _dafny.Sequence
		_ = _568_v232
		_568_v232 = _dafny.SeqOf(_567_v231, _567_v231)
		var _569_v233 D9
		_ = _569_v233
		_569_v233 = Companion_D9_.Create_DC22_(_568_v232)
		var _pat_let_tv10 = _568_v232
		_ = _pat_let_tv10
		var _source8 D9 = func(_pat_let7_0 D9) D9 {
			return func(_570_dt__update__tmp_h2 D9) D9 {
				return func(_pat_let8_0 _dafny.Sequence) D9 {
					return func(_571_dt__update_hcf28_h0 _dafny.Sequence) D9 {
						return Companion_D9_.Create_DC22_(_571_dt__update_hcf28_h0)
					}(_pat_let8_0)
				}(_pat_let_tv10)
			}(_pat_let7_0)
		}(_569_v233)
		_ = _source8
		if _source8.Is_DC23() {
			var _572___mcc_h21 bool = _source8.Get_().(D9_DC23).Cf29
			_ = _572___mcc_h21
			var _573___mcc_h22 _dafny.Int = _source8.Get_().(D9_DC23).Cf30
			_ = _573___mcc_h22
			var _574_cf30 _dafny.Int = _573___mcc_h22
			_ = _574_cf30
			var _575_cf29 bool = _572___mcc_h21
			_ = _575_cf29
			(_272_globalState).F3 = _260_v1
			(_272_globalState).F0 = _574_cf30
			var _576_v234 T0
			_ = _576_v234
			var _nw104 *C0 = New_C0_()
			_ = _nw104
			_nw104.Ctor__(_455_v164.F27, _268_v8, _451_v160)
			_576_v234 = _nw104
			var _nw105 *C0 = New_C0_()
			_ = _nw105
			_nw105.Ctor__(!((_566_v230).F25()), _268_v8, _451_v160)
			_576_v234 = _nw105
			(_272_globalState).F0 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(630))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_577_v84 _dafny.Array, _578_v1 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_579_i19 _dafny.Int) _dafny.CodePoint {
					return ((_577_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_577_v84), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_578_v1, _dafny.IntOfUint32(((_577_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_577_v84), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_359_v84, _260_v1)))).Cardinality())
		} else if _source8.Is_DC24() {
			var _580___mcc_h23 _dafny.CodePoint = _source8.Get_().(D9_DC24).Cf31
			_ = _580___mcc_h23
			var _581___mcc_h24 bool = _source8.Get_().(D9_DC24).Cf32
			_ = _581___mcc_h24
			var _582___mcc_h25 bool = _source8.Get_().(D9_DC24).Cf33
			_ = _582___mcc_h25
			var _583___mcc_h26 _dafny.Map = _source8.Get_().(D9_DC24).Cf34
			_ = _583___mcc_h26
			var _584___mcc_h27 bool = _source8.Get_().(D9_DC24).Cf35
			_ = _584___mcc_h27
			var _585_cf35 bool = _584___mcc_h27
			_ = _585_cf35
			var _586_cf34 _dafny.Map = _583___mcc_h26
			_ = _586_cf34
			var _587_cf33 bool = _582___mcc_h25
			_ = _587_cf33
			var _588_cf32 bool = _581___mcc_h24
			_ = _588_cf32
			var _589_cf31 _dafny.CodePoint = _580___mcc_h23
			_ = _589_cf31
			var _590_v235 *C2
			_ = _590_v235
			var _nw106 *C2 = New_C2_()
			_ = _nw106
			_nw106.Ctor__()
			_590_v235 = _nw106
			var _591_v236 _dafny.Sequence
			_ = _591_v236
			_591_v236 = _dafny.SeqOf(_452_v161, _452_v161)
			_591_v236 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_566_v230).F25() {
					return _591_v236
				}
				return _dafny.SeqOf(_365_v90, _365_v90, _365_v90)
			})(), _591_v236)
			(_272_globalState).F0 = Companion_Default___.SafeModuloInt(_453_v162, Companion_Default___.SafeModuloInt(Companion_Default___.Fm2(_453_v162, _362_v87, (_566_v230).F25(), _455_v164.F27, _272_globalState), _453_v162))
			var _592_v237 _dafny.Map
			_ = _592_v237
			_592_v237 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_587_cf33, _455_v164)
			_592_v237 = ((_592_v237).Merge(_592_v237)).Merge(_592_v237)
		} else if _source8.Is_DC22() {
			var _593___mcc_h28 _dafny.Sequence = _source8.Get_().(D9_DC22).Cf28
			_ = _593___mcc_h28
			var _594_cf28 _dafny.Sequence = _593___mcc_h28
			_ = _594_cf28
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_452_v161), 0))
			_ = _index92
			(_452_v161).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_566_v230).F25(), (_566_v230).F26())).Cardinality(), (_index92).Int())
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_452_v161), 0))
			_ = _index93
			var _rhs66 _dafny.Int = Companion_Default___.SafeModuloInt(_453_v162, _453_v162)
			_ = _rhs66
			var _rhs67 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(623))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}(func(_595_i20 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})), _259_v0)
			_ = _rhs67
			var _lhs57 _dafny.Array = _452_v161
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_452_v161), 0))
			_ = _lhs58
			var _lhs59 *GlobalState = _272_globalState
			_ = _lhs59
			(_lhs57).ArraySet1(_rhs66, (_lhs58).Int())
			_lhs59.F1 = _rhs67
			var _596_v238 D4
			_ = _596_v238
			_596_v238 = Companion_D4_.Create_DC13_((_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int), _455_v164.F27)
			(_455_v164).F27 = (_596_v238).Dtor_cf16()
			var _597_v239 *C1
			_ = _597_v239
			var _nw107 *C1 = New_C1_()
			_ = _nw107
			_nw107.Ctor__((_566_v230).F26(), _265_v5, _dafny.CodePoint('w'))
			_597_v239 = _nw107
			var _598_v240 T0
			_ = _598_v240
			var _nw108 *C1 = New_C1_()
			_ = _nw108
			_nw108.Ctor__(false, _265_v5, _451_v160)
			_598_v240 = _nw108
			var _599_v241 _dafny.Sequence
			_ = _599_v241
			_599_v241 = _dafny.SeqOf(_598_v240)
			_598_v240 = (_599_v241).Select((Companion_Default___.SafeIndex((_452_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_452_v161), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_599_v241).Cardinality()))).Uint32()).(T0)
		} else {
			var _600___mcc_h29 D9 = _source8.Get_().(D9_DC25).Cf36
			_ = _600___mcc_h29
			var _601_cf36 D9 = _600___mcc_h29
			_ = _601_cf36
			var _602_v242 _dafny.Map
			_ = _602_v242
			_602_v242 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v1, _454_v163)
			_602_v242 = (_602_v242).Update(_453_v162, _455_v164.F27)
			var _603_v243 _dafny.Array
			_ = _603_v243
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_15
			var _nw109 _dafny.Array
			_ = _nw109
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw109 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Map = (func(_604_v6 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_605_i21 _dafny.Int) _dafny.Map {
						return _604_v6
					}
				})(_266_v6)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw109 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw109).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw109).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_603_v243 = _nw109
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_603_v243), 0))
			_ = _index94
			(_603_v243).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v164.F27, _454_v163), (_index94).Int())
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_603_v243), 0))
			_ = _index95
			(_603_v243).ArraySet1(((_266_v6).Merge(_266_v6)).Merge(Companion_Default___.Fm9(_453_v162, _453_v162, _451_v160, _272_globalState)), (_index95).Int())
			(_272_globalState).F3 = _260_v1
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_365_v90), 0))
			_ = _index96
			(_365_v90).ArraySet1(_453_v162, (_index96).Int())
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_365_v90), 0))
			_ = _index97
			(_365_v90).ArraySet1(_453_v162, (_index97).Int())
		}
		(_455_v164).F27 = (func() bool {
			if ((_dafny.Zero).Minus(_260_v1)).Cmp(((_561_v227).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_561_v227), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_453_v162, _dafny.IntOfUint32(((_561_v227).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_561_v227), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int)) < 0 {
				return (_367_v91).IsSubsetOf(_367_v91)
			}
			return _455_v164.F27
		})()
		(_455_v164).F27 = (_566_v230).F26()
	} else {
		var _606_v244 *C2
		_ = _606_v244
		var _nw110 *C2 = New_C2_()
		_ = _nw110
		_nw110.Ctor__()
		_606_v244 = _nw110
		var _607_v245 _dafny.Map
		_ = _607_v245
		_607_v245 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_453_v162, _455_v164.F27)
		var _608_v246 _dafny.Map
		_ = _608_v246
		_608_v246 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_274_v12).Difference(_274_v12), (_607_v245).Merge(_607_v245))
		var _609_v247 _dafny.MultiSet
		_ = _609_v247
		_609_v247 = _274_v12
		_608_v246 = (((_608_v246).Update(_274_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_453_v162, _454_v163))).Update((_609_v247), _607_v245)).Update((_274_v12).Intersection(_274_v12), _607_v245)
		var _610_v248 *C2
		_ = _610_v248
		var _nw111 *C2 = New_C2_()
		_ = _nw111
		_nw111.Ctor__()
		_610_v248 = _nw111
		var _nw112 *C2 = New_C2_()
		_ = _nw112
		_nw112.Ctor__()
		_610_v248 = _nw112
		var _611_v249 D4
		_ = _611_v249
		_611_v249 = Companion_D4_.Create_DC12_()
		_611_v249 = _611_v249
		var _612_v250 _dafny.Set
		_ = _612_v250
		_612_v250 = _dafny.SetOf(_451_v160, _451_v160)
		var _613_v251 _dafny.Array
		_ = _613_v251
		var _614_v252 _dafny.Int
		_ = _614_v252
		var _615_v253 bool
		_ = _615_v253
		var _out31 _dafny.Array
		_ = _out31
		var _out32 _dafny.Int
		_ = _out32
		var _out33 bool
		_ = _out33
		_out31, _out32, _out33 = Companion_Default___.M4(_455_v164.F27, (Companion_Default___.Fm22(_272_globalState)).Intersection(_612_v250), _272_globalState)
		_613_v251 = _out31
		_614_v252 = _out32
		_615_v253 = _out33
	}
	var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_359_v84), 0))
	_ = _index98
	(_359_v84).ArraySet1((func() _dafny.Sequence {
		if (_261_v2).Contains(_260_v1) {
			return (_261_v2).Get(_260_v1).(_dafny.Sequence)
		}
		return _259_v0
	})(), (_index98).Int())
	_dafny.Print(_259_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_260_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_261_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(584), _dafny.SeqOf(_dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_263_v3, _dafny.SeqOf(_dafny.IntOfInt64(584))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_264_v4).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(584)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_265_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_266_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v7).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_268_v8).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_268_v8).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v9).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_270_v10).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_271_v11).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_272_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_272_globalState.F1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F2().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_272_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_272_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F7).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(584)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_272_globalState).F8()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_272_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_272_globalState).F12()).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F17().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_globalState.F22).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_272_globalState).F23()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v12).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(584), _dafny.IntOfInt64(584), _dafny.IntOfInt64(346))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-87))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_343_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_346_v74, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_347_v75).Equals(_dafny.SetOf(_dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_359_v84).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_359_v84).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_359_v84).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_360_v85).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_361_v86).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(584), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_362_v87).Dtor_cf1()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(584), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_363_v88).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v89).Equals(_dafny.MultiSetOf(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(893)), Companion_D0_.Create_DC0_(_dafny.IntOfInt64(893)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_365_v90).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_367_v91).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_368_v92).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mi"), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_384_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_451_v160)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_452_v161).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_452_v161).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_453_v162)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_454_v163)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_455_v164.F27)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_455_v164).F28()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_455_v164).F28()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_463_v171).Dtor_cf11())
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
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf1 _dafny.Map
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Map) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool) D0 {
	return D0{D0_DC2{Cf2}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Map {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2
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
	Cf3 _dafny.Map
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Map) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D1) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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

type D2_DC5 struct {
	Cf5 _dafny.Int
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 _dafny.Int) D2 {
	return D2{D2_DC5{Cf5}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC6 struct {
	Cf6 _dafny.Sequence
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf6 _dafny.Sequence) D2 {
	return D2{D2_DC6{Cf6}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC4 struct {
	Cf4 _dafny.Sequence
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.Sequence) D2 {
	return D2{D2_DC4{Cf4}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(_dafny.Zero)
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + data.Cf4.VerbatimString(true) + ")"
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
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf6.Equals(data2.Cf6)
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf4.Equals(data2.Cf4)
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
	Cf8  bool
	Cf9  _dafny.Int
	Cf10 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf8 bool, Cf9 _dafny.Int, Cf10 _dafny.Int) D3 {
	return D3{D3_DC8{Cf8, Cf9, Cf10}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf11 bool
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf11 bool) D3 {
	return D3{D3_DC9{Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf12 T0
	Cf13 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf12 T0, Cf13 _dafny.Int) D3 {
	return D3{D3_DC10{Cf12, Cf13}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC7 struct {
	Cf7 _dafny.Array
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf7 _dafny.Array) D3 {
	return D3{D3_DC7{Cf7}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D3) Dtor_cf8() bool {
	return _this.Get_().(D3_DC8).Cf8
}

func (_this D3) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf9
}

func (_this D3) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf10
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) Dtor_cf12() T0 {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D3_DC7).Cf7
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Cmp(data2.Cf10) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf11 == data2.Cf11
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && _dafny.AreEqual(data1.Cf12, data2.Cf12) && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf7 == data2.Cf7
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
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_() D4 {
	return D4{D4_DC12{}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf15 _dafny.Int
	Cf16 bool
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf15 _dafny.Int, Cf16 bool) D4 {
	return D4{D4_DC13{Cf15, Cf16}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf17 bool
	Cf18 _dafny.Int
	Cf19 bool
	Cf20 bool
	Cf21 _dafny.Int
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf17 bool, Cf18 _dafny.Int, Cf19 bool, Cf20 bool, Cf21 _dafny.Int) D4 {
	return D4{D4_DC14{Cf17, Cf18, Cf19, Cf20, Cf21}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC11 struct {
	Cf14 _dafny.Array
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf14 _dafny.Array) D4 {
	return D4{D4_DC11{Cf14}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_()
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf15
}

func (_this D4) Dtor_cf16() bool {
	return _this.Get_().(D4_DC13).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC14).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf18
}

func (_this D4) Dtor_cf19() bool {
	return _this.Get_().(D4_DC14).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC14).Cf20
}

func (_this D4) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf14
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf14) + ")"
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
			_, ok := other.Get_().(D4_DC12)
			return ok
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20 && data1.Cf21.Cmp(data2.Cf21) == 0
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf14 == data2.Cf14
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
	Cf22 _dafny.Map
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf22 _dafny.Map) D5 {
	return D5{D5_DC15{Cf22}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D5) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D5_DC15).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf22.Equals(data2.Cf22)
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
	Cf23 _dafny.MultiSet
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf23 _dafny.MultiSet) D6 {
	return D6{D6_DC16{Cf23}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D6) Dtor_cf23() _dafny.MultiSet {
	return _this.Get_().(D6_DC16).Cf23
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && data1.Cf23.Equals(data2.Cf23)
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
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_() D7 {
	return D7{D7_DC18{}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
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

type D7_DC17 struct {
	Cf24 _dafny.Set
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf24 _dafny.Set) D7 {
	return D7{D7_DC17{Cf24}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_()
}

func (_this D7) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D7_DC19).Cf25
}

func (_this D7) Dtor_cf24() _dafny.Set {
	return _this.Get_().(D7_DC17).Cf24
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf24) + ")"
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
			_, ok := other.Get_().(D7_DC18)
			return ok
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf27 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf27 _dafny.Int) D8 {
	return D8{D8_DC21{Cf27}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC20 struct {
	Cf26 _dafny.Map
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf26 _dafny.Map) D8 {
	return D8{D8_DC20{Cf26}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(_dafny.Zero)
}

func (_this D8) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf27
}

func (_this D8) Dtor_cf26() _dafny.Map {
	return _this.Get_().(D8_DC20).Cf26
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D9_DC23 struct {
	Cf29 bool
	Cf30 _dafny.Int
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf29 bool, Cf30 _dafny.Int) D9 {
	return D9{D9_DC23{Cf29, Cf30}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC24 struct {
	Cf31 _dafny.CodePoint
	Cf32 bool
	Cf33 bool
	Cf34 _dafny.Map
	Cf35 bool
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf31 _dafny.CodePoint, Cf32 bool, Cf33 bool, Cf34 _dafny.Map, Cf35 bool) D9 {
	return D9{D9_DC24{Cf31, Cf32, Cf33, Cf34, Cf35}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC22 struct {
	Cf28 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf28 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf28}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC25 struct {
	Cf36 D9
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf36 D9) D9 {
	return D9{D9_DC25{Cf36}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC23_(false, _dafny.Zero)
}

func (_this D9) Dtor_cf29() bool {
	return _this.Get_().(D9_DC23).Cf29
}

func (_this D9) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf30
}

func (_this D9) Dtor_cf31() _dafny.CodePoint {
	return _this.Get_().(D9_DC24).Cf31
}

func (_this D9) Dtor_cf32() bool {
	return _this.Get_().(D9_DC24).Cf32
}

func (_this D9) Dtor_cf33() bool {
	return _this.Get_().(D9_DC24).Cf33
}

func (_this D9) Dtor_cf34() _dafny.Map {
	return _this.Get_().(D9_DC24).Cf34
}

func (_this D9) Dtor_cf35() bool {
	return _this.Get_().(D9_DC24).Cf35
}

func (_this D9) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf28
}

func (_this D9) Dtor_cf36() D9 {
	return _this.Get_().(D9_DC25).Cf36
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf36) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33 && data1.Cf34.Equals(data2.Cf34) && data1.Cf35 == data2.Cf35
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf36.Equals(data2.Cf36)
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

type D10_DC27 struct {
	Cf38 _dafny.Int
	Cf39 _dafny.Array
	Cf40 _dafny.Int
	Cf41 bool
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf38 _dafny.Int, Cf39 _dafny.Array, Cf40 _dafny.Int, Cf41 bool) D10 {
	return D10{D10_DC27{Cf38, Cf39, Cf40, Cf41}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

type D10_DC26 struct {
	Cf37 *C2
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf37 *C2) D10 {
	return D10{D10_DC26{Cf37}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC27_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, false)
}

func (_this D10) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D10_DC27).Cf38
}

func (_this D10) Dtor_cf39() _dafny.Array {
	return _this.Get_().(D10_DC27).Cf39
}

func (_this D10) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D10_DC27).Cf40
}

func (_this D10) Dtor_cf41() bool {
	return _this.Get_().(D10_DC27).Cf41
}

func (_this D10) Dtor_cf37() *C2 {
	return _this.Get_().(D10_DC26).Cf37
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf37) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39 == data2.Cf39 && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41 == data2.Cf41
		}
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf37 == data2.Cf37
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

// Definition of datatype D11
type D11 struct {
	Data_D11_
}

func (_this D11) Get_() Data_D11_ {
	return _this.Data_D11_
}

type Data_D11_ interface {
	isD11()
}

type CompanionStruct_D11_ struct {
}

var Companion_D11_ = CompanionStruct_D11_{}

type D11_DC29 struct {
	Cf43 _dafny.Int
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf43 _dafny.Int) D11 {
	return D11{D11_DC29{Cf43}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC30 struct {
	Cf44 bool
	Cf45 bool
	Cf46 _dafny.Int
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf44 bool, Cf45 bool, Cf46 _dafny.Int) D11 {
	return D11{D11_DC30{Cf44, Cf45, Cf46}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC28 struct {
	Cf42 _dafny.MultiSet
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf42 _dafny.MultiSet) D11 {
	return D11{D11_DC28{Cf42}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC31 struct {
	Cf47 D11
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf47 D11) D11 {
	return D11{D11_DC31{Cf47}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC29_(_dafny.Zero)
}

func (_this D11) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D11_DC29).Cf43
}

func (_this D11) Dtor_cf44() bool {
	return _this.Get_().(D11_DC30).Cf44
}

func (_this D11) Dtor_cf45() bool {
	return _this.Get_().(D11_DC30).Cf45
}

func (_this D11) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D11_DC30).Cf46
}

func (_this D11) Dtor_cf42() _dafny.MultiSet {
	return _this.Get_().(D11_DC28).Cf42
}

func (_this D11) Dtor_cf47() D11 {
	return _this.Get_().(D11_DC31).Cf47
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf43.Cmp(data2.Cf43) == 0
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf44 == data2.Cf44 && data1.Cf45 == data2.Cf45 && data1.Cf46.Cmp(data2.Cf46) == 0
		}
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf47.Equals(data2.Cf47)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D11) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D11)
	return ok && _this.Equals(typed)
}

func Type_D11_() _dafny.TypeDescriptor {
	return type_D11_{}
}

type type_D11_ struct {
}

func (_this type_D11_) Default() interface{} {
	return Companion_D11_.Default()
}

func (_this type_D11_) String() string {
	return "main.D11"
}
func (_this D11) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D11{}

// End of datatype D11

// Definition of trait T0
type T0 interface {
	String() string
	F24() _dafny.CodePoint
	F24_set_(value _dafny.CodePoint)
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
	F0   _dafny.Int
	F1   _dafny.Sequence
	F3   _dafny.Int
	F6   _dafny.Int
	F7   _dafny.Set
	F9   _dafny.CodePoint
	F22  _dafny.Array
	_f2  _dafny.Sequence
	_f4  _dafny.Int
	_f5  _dafny.Int
	_f8  _dafny.Map
	_f10 _dafny.Int
	_f11 bool
	_f12 _dafny.Set
	_f13 _dafny.Array
	_f14 bool
	_f15 bool
	_f16 bool
	_f17 _dafny.Sequence
	_f18 bool
	_f19 bool
	_f20 bool
	_f21 bool
	_f23 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F1 = _dafny.EmptySeq
	_this.F3 = _dafny.Zero
	_this.F6 = _dafny.Zero
	_this.F7 = _dafny.EmptySet
	_this.F9 = _dafny.CodePoint('D')
	_this.F22 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = _dafny.EmptySeq
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f8 = _dafny.EmptyMap
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f12 = _dafny.EmptySet
	_this._f13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f14 = false
	_this._f15 = false
	_this._f16 = false
	_this._f17 = _dafny.EmptySeq
	_this._f18 = false
	_this._f19 = false
	_this._f20 = false
	_this._f21 = false
	_this._f23 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Sequence, f2 _dafny.Sequence, f3 _dafny.Int, f4 _dafny.Int, f5 _dafny.Int, f6 _dafny.Int, f7 _dafny.Set, f8 _dafny.Map, f9 _dafny.CodePoint, f10 _dafny.Int, f11 bool, f12 _dafny.Set, f13 _dafny.Array, f14 bool, f15 bool, f16 bool, f17 _dafny.Sequence, f18 bool, f19 bool, f20 bool, f21 bool, f22 _dafny.Array, f23 _dafny.Array) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
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
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this).F22 = f22
		(_this)._f23 = f23
	}
}
func (_this *GlobalState) F2() _dafny.Sequence {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F8() _dafny.Map {
	{
		return _this._f8
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
func (_this *GlobalState) F12() _dafny.Set {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Array {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() bool {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() bool {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() bool {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() bool {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F23() _dafny.Array {
	{
		return _this._f23
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f24 _dafny.CodePoint
	F27  bool
	_f28 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f24 = _dafny.CodePoint('D')
	_this.F27 = false
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) F24() _dafny.CodePoint {
	return _this._f24
}
func (_this *C0) F24_set_(value _dafny.CodePoint) {
	_this._f24 = value
}
func (_this *C0) Ctor__(f27 bool, f28 _dafny.Array, f24 _dafny.CodePoint) {
	{
		(_this).F27 = f27
		(_this)._f28 = f28
		(_this)._f24 = f24
	}
}
func (_this *C0) Fm7(p0 _dafny.Map, p1 bool, globalState *GlobalState) bool {
	{
		return ((func() bool {
			if false {
				return _this.F27
			}
			return _this.F27
		})()) || (!(_this.F27))
	}
}
func (_this *C0) M3(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _616_v0 _dafny.Sequence
		_ = _616_v0
		_616_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dlew")
		var _617_v1 _dafny.MultiSet
		_ = _617_v1
		_617_v1 = _dafny.MultiSetOf(_616_v0, _616_v0, _dafny.UnicodeSeqOfUtf8Bytes("hedtlengs"))
		_617_v1 = _617_v1
		var _618_i0 _dafny.Int
		_ = _618_i0
		_618_i0 = _dafny.Zero
		{
			for true {
				{
					if (_618_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_618_i0 = (_618_i0).Plus(_dafny.One)
					(globalState).F0 = p0
					if p2 {
						var _619_v2 _dafny.Array
						_ = _619_v2
						var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
						_ = _nw113
						_619_v2 = _nw113
						var _620_v3 _dafny.MultiSet
						_ = _620_v3
						_620_v3 = _dafny.MultiSetOf(_619_v2, _619_v2, _619_v2, _619_v2)
						var _621_v4 D2
						_ = _621_v4
						_621_v4 = Companion_D2_.Create_DC5_(p0)
						var _622_v5 _dafny.Set
						_ = _622_v5
						_622_v5 = _dafny.SetOf(_621_v4)
						var _623_v6 _dafny.Array
						_ = _623_v6
						var _nwElement0_19 _dafny.Int = _dafny.IntOfInt64(286)
						_ = _nwElement0_19
						var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(5))
						_ = _nw114
						(_nw114).ArraySet1(_nwElement0_19, 0)
						(_nw114).ArraySet1((Companion_Default___.Fm8(Companion_Default___.Fm8(p0, p0, false, globalState), p0, _this.F27, globalState)).Minus((_620_v3).Cardinality()), 1)
						(_nw114).ArraySet1((_622_v5).Cardinality(), 2)
						(_nw114).ArraySet1((_dafny.Zero).Minus(p0), 3)
						(_nw114).ArraySet1(p0, 4)
						_623_v6 = _nw114
						_619_v2 = _623_v6
						var _624_v7 _dafny.Set
						_ = _624_v7
						_624_v7 = _dafny.SetOf(p0)
						var _625_v9 _dafny.Sequence
						_ = _625_v9
						_625_v9 = _dafny.SeqOf(_624_v7, func() _dafny.Set {
							var _coll22 = _dafny.NewBuilder()
							_ = _coll22
							for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(571), _dafny.IntOfInt64(906))); ; {
								_compr_22, _ok23 := _iter23()
								if !_ok23 {
									break
								}
								var _626_v8 _dafny.Int
								_626_v8 = interface{}(_compr_22).(_dafny.Int)
								if ((_dafny.IntOfInt64(571)).Cmp(_626_v8) <= 0) && ((_626_v8).Cmp(_dafny.IntOfInt64(906)) < 0) {
									_coll22.Add((_626_v8).Times(_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality())))
								}
							}
							return _coll22.ToSet()
						}())
						var _627_v10 _dafny.Sequence
						_ = _627_v10
						_627_v10 = _dafny.SeqOf(_dafny.IntOfUint32((_625_v9).Cardinality()))
						var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_623_v6), 0))
						_ = _index99
						(_623_v6).ArraySet1(_dafny.IntOfUint32((_627_v10).Cardinality()), (_index99).Int())
						var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_623_v6), 0))
						_ = _index100
						(_623_v6).ArraySet1(p0, (_index100).Int())
						var _628_v11 _dafny.Map
						_ = _628_v11
						_628_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
						_628_v11 = (_628_v11).Update(_dafny.CodePoint('t'), (_dafny.IntOfInt64(367)).Plus(p0))
						var _629_v12 _dafny.Set
						_ = _629_v12
						_629_v12 = _dafny.SetOf(p2)
						(globalState).F3 = ((_629_v12).Union(_629_v12)).Cardinality()
						var _630_v13 _dafny.Array
						_ = _630_v13
						var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
						_ = _nw115
						_630_v13 = _nw115
						var _631_v14 _dafny.Array
						_ = _631_v14
						var _len0_16 _dafny.Int = _dafny.IntOfInt64(8)
						_ = _len0_16
						var _nw116 _dafny.Array
						_ = _nw116
						if _len0_16.Cmp(_dafny.Zero) == 0 {
							_nw116 = _dafny.NewArray(_len0_16)
						} else {
							var _init16 func(_dafny.Int) bool = func(_632_i1 _dafny.Int) bool {
								return _this.F27
							}
							_ = _init16
							var _element0_16 = _init16(_dafny.Zero)
							_ = _element0_16
							_nw116 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
							(_nw116).ArraySet1(_element0_16, 0)
							var _nativeLen0_16 = (_len0_16).Int()
							_ = _nativeLen0_16
							for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
								(_nw116).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
							}
						}
						_631_v14 = _nw116
						var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_630_v13), 0))
						_ = _index101
						(_630_v13).ArraySet1(_631_v14, (_index101).Int())
						var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_630_v13), 0))
						_ = _index102
						(_630_v13).ArraySet1((Companion_D3_.Create_DC7_(_631_v14)).Dtor_cf7(), (_index102).Int())
					} else {
						var _633_v15 _dafny.Sequence
						_ = _633_v15
						_633_v15 = _dafny.SeqOf(p2, !(_this.F27))
						var _634_v16 _dafny.Map
						_ = _634_v16
						_634_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() bool {
							if p2 {
								return p2
							}
							return !((_633_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_633_v15).Cardinality()))).Uint32()).(bool))
						})())
						(globalState).F3 = (_634_v16).Cardinality()
						(_this).F27 = _this.F27
						(globalState).F3 = p0
						(globalState).F9 = (_616_v0).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_616_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
						(_this).F27 = !(!(p2))
					}
					var _635_v17 _dafny.Map
					_ = _635_v17
					_635_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, _this.F27)
					_635_v17 = (_635_v17).Update(false, false)
					(globalState).F1 = _616_v0
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		if (_this.F27) && (p2) {
			(_this).F27 = (func() bool {
				if p2 {
					return (func() bool {
						if p2 {
							return _this.F27
						}
						return p2
					})()
				}
				return _this.F27
			})()
			var _636_v18 _dafny.Map
			_ = _636_v18
			_636_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F27)
			(globalState).F6 = (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(252), _dafny.IntOfInt64(-199))).Minus((func() _dafny.Int {
				if true {
					return (_636_v18).Cardinality()
				}
				return p0
			})())
			(globalState).F3 = p0
			var _637_v19 _dafny.MultiSet
			_ = _637_v19
			_637_v19 = _dafny.MultiSetOf(p2, _this.F27, p2)
			(_this).F27 = !((p0).Cmp((func() _dafny.Int {
				if (_637_v19).Contains(false) {
					return (_637_v19).Multiplicity(false)
				}
				return _dafny.IntOfInt64(420)
			})()) <= 0)
			(_this).F27 = !_dafny.Companion_Sequence_.Equal(_616_v0, _616_v0)
		} else {
			var _638_v20 _dafny.Array
			_ = _638_v20
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_17
			var _nw117 _dafny.Array
			_ = _nw117
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw117 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Int = (func(_639_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_640_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_640_i2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_639_v0).Cardinality())))
					}
				})(_616_v0)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw117 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw117).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw117).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_638_v20 = _nw117
			_638_v20 = _638_v20
			(_this).F27 = _this.F27
			(_this).F27 = !((!(p2)) && (p2))
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index103
			((_this).F28()).ArraySet1(_638_v20, (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index104
			((_this).F28()).ArraySet1(_638_v20, (_index104).Int())
			var _641_v21 _dafny.Array
			_ = _641_v21
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(11))
			_ = _nw118
			_641_v21 = _nw118
			var _642_v22 _dafny.Map
			_ = _642_v22
			_642_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, _641_v21)
			var _643_v23 _dafny.Array
			_ = _643_v23
			var _nwElement0_20 _dafny.Array = _641_v21
			_ = _nwElement0_20
			var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(21))
			_ = _nw119
			(_nw119).ArraySet1(_nwElement0_20, 0)
			(_nw119).ArraySet1((func() _dafny.Array {
				if (_642_v22).Contains(p2) {
					return (_642_v22).Get(p2).(_dafny.Array)
				}
				return _641_v21
			})(), 1)
			(_nw119).ArraySet1(_641_v21, 2)
			(_nw119).ArraySet1(_641_v21, 3)
			(_nw119).ArraySet1(_641_v21, 4)
			(_nw119).ArraySet1(_641_v21, 5)
			(_nw119).ArraySet1(_641_v21, 6)
			(_nw119).ArraySet1(_641_v21, 7)
			(_nw119).ArraySet1(_641_v21, 8)
			(_nw119).ArraySet1(_641_v21, 9)
			(_nw119).ArraySet1(_641_v21, 10)
			(_nw119).ArraySet1(_641_v21, 11)
			(_nw119).ArraySet1(_641_v21, 12)
			(_nw119).ArraySet1(_641_v21, 13)
			(_nw119).ArraySet1(_641_v21, 14)
			(_nw119).ArraySet1(_641_v21, 15)
			(_nw119).ArraySet1(_641_v21, 16)
			(_nw119).ArraySet1((func() _dafny.Array {
				if true {
					return _641_v21
				}
				return _641_v21
			})(), 17)
			(_nw119).ArraySet1(_641_v21, 18)
			(_nw119).ArraySet1((func() _dafny.Array {
				if p2 {
					return _641_v21
				}
				return _641_v21
			})(), 19)
			(_nw119).ArraySet1(_641_v21, 20)
			_643_v23 = _nw119
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_643_v23), 0))
			_ = _index105
			(_643_v23).ArraySet1(_641_v21, (_index105).Int())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_643_v23), 0))
			_ = _index106
			(_643_v23).ArraySet1(_641_v21, (_index106).Int())
		}
		var _644_v24 _dafny.Sequence
		_ = _644_v24
		_644_v24 = _dafny.SeqOf((Companion_Default___.Fm9(p0, p0, _this.F24(), globalState)).Cardinality(), ((_dafny.MultiSetFromSeq(_dafny.SeqOf(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vofekyvc")).Cardinality())))).Difference(Companion_Default___.Fm10(globalState))).Cardinality(), p0)
		(globalState).F3 = _dafny.IntOfUint32((_644_v24).Cardinality())
		var _645_v25 _dafny.Set
		_ = _645_v25
		_645_v25 = _dafny.SetOf((p0).Plus(_dafny.IntOfInt64(-221)))
		_645_v25 = (func() _dafny.Set {
			if _this.F27 {
				return _645_v25
			}
			return (_645_v25).Difference(func() _dafny.Set {
				var _coll23 = _dafny.NewBuilder()
				_ = _coll23
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(191), _dafny.IntOfInt64(846))); ; {
					_compr_23, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _646_v26 _dafny.Int
					_646_v26 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(191)).Cmp(_646_v26) <= 0) && ((_646_v26).Cmp(_dafny.IntOfInt64(846)) < 0) {
						_coll23.Add((_646_v26).Minus(p0))
					}
				}
				return _coll23.ToSet()
			}())
		})()
		var _647_v27 _dafny.Map
		_ = _647_v27
		_647_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)
		var _648_v28 _dafny.Map
		_ = _648_v28
		_648_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if p2 {
				return (_647_v27).Cardinality()
			}
			return p0
		})(), _dafny.Companion_Sequence_.Concatenate(_616_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(867))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_649_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_650_i3 _dafny.Int) _dafny.CodePoint {
				return _649_p1
			}
		})(p1)))))
		_648_v28 = (_648_v28).Update(p0, _616_v0)
		var _651_v29 _dafny.Sequence
		_ = _651_v29
		_651_v29 = _dafny.SeqOf(_644_v24)
		var _652_v30 _dafny.Sequence
		_ = _652_v30
		_652_v30 = _dafny.SeqOf(p2)
		r0 = (_616_v0).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_651_v29).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_651_v29).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_651_v29).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_651_v29).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), p0)).Cardinality())).Times(_dafny.IntOfUint32((_652_v30).Cardinality())), _dafny.IntOfUint32((_616_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
		r1 = p0
		r2 = _dafny.IntOfInt64(446)
		return r0, r1, r2
	}
}
func (_this *C0) F28() _dafny.Array {
	{
		return _this._f28
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f24 _dafny.CodePoint
	_f25 bool
	_f26 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f24 = _dafny.CodePoint('D')
	_this._f25 = false
	_this._f26 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F24() _dafny.CodePoint {
	return _this._f24
}
func (_this *C1) F24_set_(value _dafny.CodePoint) {
	_this._f24 = value
}
func (_this *C1) Ctor__(f25 bool, f26 bool, f24 _dafny.CodePoint) {
	{
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this)._f24 = f24
	}
}
func (_this *C1) Fm3(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gcfpseboa"), _dafny.UnicodeSeqOfUtf8Bytes("hljhlevh")), _dafny.UnicodeSeqOfUtf8Bytes("mmc"))).Cardinality())
	}
}
func (_this *C1) Fm4(p0 _dafny.Map, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(52)
	}
}
func (_this *C1) M2(p0 _dafny.Sequence, p1 _dafny.Array, p2 T0, globalState *GlobalState) (bool, _dafny.Array, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _653_i0 _dafny.Int
		_ = _653_i0
		_653_i0 = _dafny.Zero
		{
			for (_this).F25() {
				{
					if (_653_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_653_i0 = (_653_i0).Plus(_dafny.One)
					var _654_v0 _dafny.Int
					_ = _654_v0
					_654_v0 = _dafny.IntOfInt64(-553)
					var _655_v1 _dafny.Map
					_ = _655_v1
					_655_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), !((_this).F25()))
					if !((_this).F26()) || (!((_654_v0).Cmp((_655_v1).Cardinality()) > 0)) {
						var _656_v2 _dafny.Map
						_ = _656_v2
						_656_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _654_v0)
						var _657_v3 D0
						_ = _657_v3
						_657_v3 = Companion_D0_.Create_DC0_((func() _dafny.Int {
							if (_656_v2).Contains((_this).F26()) {
								return (_656_v2).Get((_this).F26()).(_dafny.Int)
							}
							return _dafny.IntOfInt64(699)
						})())
						var _658_v4 _dafny.Sequence
						_ = _658_v4
						_658_v4 = _dafny.SeqOf((_this).F25())
						var _659_v5 _dafny.Map
						_ = _659_v5
						_659_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v0, _654_v0)
						var _660_v6 _dafny.Map
						_ = _660_v6
						_660_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_658_v4).Cardinality()), _659_v5)
						var _661_v8 _dafny.Set
						_ = _661_v8
						_661_v8 = _dafny.SetOf(func() _dafny.Set {
							var _coll24 = _dafny.NewBuilder()
							_ = _coll24
							for _iter25 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC2_((_this).F26()), (_this).F26())).Keys().Elements()); ; {
								_compr_24, _ok25 := _iter25()
								if !_ok25 {
									break
								}
								var _662_v7 D0
								_662_v7 = interface{}(_compr_24).(D0)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC2_((_this).F26()), (_this).F26())).Contains(_662_v7) {
									_coll24.Add(_662_v7)
								}
							}
							return _coll24.ToSet()
						}())
						var _pat_let_tv11 = _656_v2
						_ = _pat_let_tv11
						var _pat_let_tv12 = _656_v2
						_ = _pat_let_tv12
						var _pat_let_tv13 = _657_v3
						_ = _pat_let_tv13
						var _pat_let_tv14 = _654_v0
						_ = _pat_let_tv14
						var _663_v9 _dafny.Array
						_ = _663_v9
						var _nwElement0_21 D0 = _657_v3
						_ = _nwElement0_21
						var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(23))
						_ = _nw120
						(_nw120).ArraySet1(_nwElement0_21, 0)
						(_nw120).ArraySet1(func(_pat_let9_0 D0) D0 {
							return func(_666_dt__update__tmp_h1 D0) D0 {
								return func(_pat_let12_0 _dafny.Int) D0 {
									return func(_667_dt__update_hcf0_h1 _dafny.Int) D0 {
										return Companion_D0_.Create_DC0_(_667_dt__update_hcf0_h1)
									}(_pat_let12_0)
								}(_pat_let_tv14)
							}(_pat_let9_0)
						}(func(_pat_let10_0 D0) D0 {
							return func(_664_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let11_0 _dafny.Int) D0 {
									return func(_665_dt__update_hcf0_h0 _dafny.Int) D0 {
										return Companion_D0_.Create_DC0_(_665_dt__update_hcf0_h0)
									}(_pat_let11_0)
								}((func() _dafny.Int {
									if (_pat_let_tv11).Contains((_this).F26()) {
										return (_pat_let_tv12).Get((_this).F26()).(_dafny.Int)
									}
									return (_pat_let_tv13).Dtor_cf0()
								})())
							}(_pat_let10_0)
						}(_657_v3)), 1)
						(_nw120).ArraySet1(Companion_Default___.Fm5(_660_v6, _654_v0, _661_v8, globalState), 2)
						(_nw120).ArraySet1(_657_v3, 3)
						(_nw120).ArraySet1(Companion_D0_.Create_DC0_(_654_v0), 4)
						(_nw120).ArraySet1(Companion_D0_.Create_DC0_(_654_v0), 5)
						(_nw120).ArraySet1(func(_pat_let13_0 D0) D0 {
							return func(_668_dt__update__tmp_h2 D0) D0 {
								return func(_pat_let14_0 _dafny.Int) D0 {
									return func(_669_dt__update_hcf0_h2 _dafny.Int) D0 {
										return Companion_D0_.Create_DC0_(_669_dt__update_hcf0_h2)
									}(_pat_let14_0)
								}(_dafny.IntOfInt64(-255))
							}(_pat_let13_0)
						}(_657_v3), 6)
						(_nw120).ArraySet1(_657_v3, 7)
						(_nw120).ArraySet1(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-420)), 8)
						(_nw120).ArraySet1(_657_v3, 9)
						(_nw120).ArraySet1(_657_v3, 10)
						(_nw120).ArraySet1(_657_v3, 11)
						(_nw120).ArraySet1(_657_v3, 12)
						(_nw120).ArraySet1(_657_v3, 13)
						(_nw120).ArraySet1(_657_v3, 14)
						(_nw120).ArraySet1(_657_v3, 15)
						(_nw120).ArraySet1(_657_v3, 16)
						(_nw120).ArraySet1(_657_v3, 17)
						(_nw120).ArraySet1(Companion_D0_.Create_DC0_(_654_v0), 18)
						(_nw120).ArraySet1(Companion_D0_.Create_DC0_(_654_v0), 19)
						(_nw120).ArraySet1(_657_v3, 20)
						(_nw120).ArraySet1(Companion_D0_.Create_DC0_(_654_v0), 21)
						(_nw120).ArraySet1(_657_v3, 22)
						_663_v9 = _nw120
						var _670_v11 _dafny.Sequence
						_ = _670_v11
						_670_v11 = _dafny.SeqOf(_654_v0, _654_v0)
						var _671_v12 _dafny.Array
						_ = _671_v12
						var _nwElement0_22 bool = (_this).F25()
						_ = _nwElement0_22
						var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(14))
						_ = _nw121
						(_nw121).ArraySet1(_nwElement0_22, 0)
						(_nw121).ArraySet1((_this).F26(), 1)
						(_nw121).ArraySet1((_this).F26(), 2)
						(_nw121).ArraySet1((_this).F26(), 3)
						(_nw121).ArraySet1((_this).F25(), 4)
						(_nw121).ArraySet1((_this).F25(), 5)
						(_nw121).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_this).Fm4(func() _dafny.Map {
							var _coll25 = _dafny.NewMapBuilder()
							_ = _coll25
							for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(198), _dafny.IntOfInt64(631))); ; {
								_compr_25, _ok26 := _iter26()
								if !_ok26 {
									break
								}
								var _672_v10 _dafny.Int
								_672_v10 = interface{}(_compr_25).(_dafny.Int)
								if ((_dafny.IntOfInt64(198)).Cmp(_672_v10) <= 0) && ((_672_v10).Cmp(_dafny.IntOfInt64(631)) < 0) {
									_coll25.Add((_672_v10).Times((_dafny.MultiSetFromSeq(_658_v4)).Cardinality()), _dafny.SetOf(_654_v0, _654_v0))
								}
							}
							return _coll25.ToMap()
						}(), (_this).F25(), globalState), _654_v0), _670_v11), 6)
						(_nw121).ArraySet1((_this).F25(), 7)
						(_nw121).ArraySet1((func() bool {
							if false {
								return Companion_Default___.Fm6(globalState)
							}
							return (_this).F26()
						})(), 8)
						(_nw121).ArraySet1((_this).F25(), 9)
						(_nw121).ArraySet1(((_this).F26()) && (true), 10)
						(_nw121).ArraySet1((_this).F25(), 11)
						(_nw121).ArraySet1((_654_v0).Cmp(_654_v0) != 0, 12)
						(_nw121).ArraySet1(false, 13)
						_671_v12 = _nw121
						var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_671_v12), 0))
						_ = _index107
						(_671_v12).ArraySet1(_dafny.Companion_Sequence_.Contains(_658_v4, (_this).F25()), (_index107).Int())
						var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_671_v12), 0))
						_ = _index108
						var _rhs68 _dafny.Array = _663_v9
						_ = _rhs68
						var _rhs69 bool = Companion_Default___.Fm6(globalState)
						_ = _rhs69
						var _rhs70 bool = false
						_ = _rhs70
						var _lhs60 _dafny.Array = _671_v12
						_ = _lhs60
						var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_671_v12), 0))
						_ = _lhs61
						_663_v9 = _rhs68
						(_lhs60).ArraySet1(_rhs69, (_lhs61).Int())
						r0 = _rhs70
						var _673_v13 *C0
						_ = _673_v13
						var _nw122 *C0 = New_C0_()
						_ = _nw122
						_nw122.Ctor__((_this).F25(), p1, _dafny.CodePoint('j'))
						_673_v13 = _nw122
						(globalState).F3 = _654_v0
						var _674_v14 D2
						_ = _674_v14
						_674_v14 = Companion_D2_.Create_DC6_(_670_v11)
						_674_v14 = _674_v14
						(globalState).F6 = _654_v0
					} else {
						var _675_v15 _dafny.Map
						_ = _675_v15
						_675_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_654_v0).Cmp(_654_v0) < 0, _654_v0)
						(globalState).F0 = (func() _dafny.Int {
							if (_675_v15).Contains((_this).F26()) {
								return (_675_v15).Get((_this).F26()).(_dafny.Int)
							}
							return _654_v0
						})()
						var _676_v16 _dafny.Sequence
						_ = _676_v16
						_676_v16 = _dafny.SeqOf((_dafny.Zero).Minus(_654_v0))
						var _677_v17 _dafny.Set
						_ = _677_v17
						_677_v17 = _dafny.SetOf(p2.F24())
						var _678_v18 _dafny.Set
						_ = _678_v18
						_678_v18 = _dafny.SetOf(_677_v17, _677_v17)
						var _679_v19 _dafny.Set
						_ = _679_v19
						_679_v19 = _dafny.SetOf(_654_v0, _654_v0, (_678_v18).Cardinality())
						var _680_v20 _dafny.Set
						_ = _680_v20
						_680_v20 = _dafny.SetOf(!((_this).F26()))
						var _681_v21 _dafny.Map
						_ = _681_v21
						_681_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2.F24(), _654_v0)
						var _682_v22 _dafny.MultiSet
						_ = _682_v22
						_682_v22 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _654_v0), _681_v21, _681_v21)
						var _683_v23 _dafny.MultiSet
						_ = _683_v23
						_683_v23 = _dafny.MultiSetOf(_654_v0)
						var _684_v24 _dafny.MultiSet
						_ = _684_v24
						_684_v24 = _dafny.MultiSetOf((_this).F25())
						var _685_v25 _dafny.Array
						_ = _685_v25
						var _nwElement0_23 _dafny.Int = Companion_Default___.SafeDivisionInt(_654_v0, _654_v0)
						_ = _nwElement0_23
						var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(25))
						_ = _nw123
						(_nw123).ArraySet1(_nwElement0_23, 0)
						(_nw123).ArraySet1((_676_v16).Select((Companion_Default___.SafeIndex(_654_v0, _dafny.IntOfUint32((_676_v16).Cardinality()))).Uint32()).(_dafny.Int), 1)
						(_nw123).ArraySet1(_654_v0, 2)
						(_nw123).ArraySet1(_dafny.IntOfInt64(75), 3)
						(_nw123).ArraySet1(_654_v0, 4)
						(_nw123).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_this).F25() {
								return _676_v16
							}
							return _dafny.SeqOf(_654_v0)
						})()).Cardinality()), 5)
						(_nw123).ArraySet1(_dafny.IntOfInt64(700), 6)
						(_nw123).ArraySet1(_654_v0, 7)
						(_nw123).ArraySet1(_dafny.IntOfUint32((_676_v16).Cardinality()), 8)
						(_nw123).ArraySet1((_dafny.Zero).Minus(_654_v0), 9)
						(_nw123).ArraySet1((_654_v0).Times(_654_v0), 10)
						(_nw123).ArraySet1(Companion_Default___.SafeModuloInt(_654_v0, (_this).Fm3(_679_v19, (_this).F25(), p0, (_this).F25(), globalState)), 11)
						(_nw123).ArraySet1(_dafny.IntOfInt64(344), 12)
						(_nw123).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_654_v0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), p2.F24())).Cardinality())), 13)
						(_nw123).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_654_v0, (_680_v20).Cardinality())), 14)
						(_nw123).ArraySet1(_654_v0, 15)
						(_nw123).ArraySet1(_654_v0, 16)
						(_nw123).ArraySet1((func() _dafny.Int {
							if (_682_v22).Contains(_681_v21) {
								return (_682_v22).Multiplicity(_681_v21)
							}
							return _654_v0
						})(), 17)
						(_nw123).ArraySet1((_dafny.Zero).Minus(_654_v0), 18)
						(_nw123).ArraySet1((_654_v0).Times(_dafny.IntOfUint32((p0).Cardinality())), 19)
						(_nw123).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
							if (_683_v23).Contains(_654_v0) {
								return (_683_v23).Multiplicity(_654_v0)
							}
							return _654_v0
						})()), 20)
						(_nw123).ArraySet1(_654_v0, 21)
						(_nw123).ArraySet1(_dafny.IntOfInt64(360), 22)
						(_nw123).ArraySet1((func() _dafny.Int {
							if (_684_v24).Contains((_this).F25()) {
								return (_684_v24).Multiplicity((_this).F25())
							}
							return _654_v0
						})(), 23)
						(_nw123).ArraySet1(_654_v0, 24)
						_685_v25 = _nw123
						var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_685_v25), 0))
						_ = _index109
						(_685_v25).ArraySet1(((_679_v19).Intersection(_679_v19)).Cardinality(), (_index109).Int())
						var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_685_v25), 0))
						_ = _index110
						(_685_v25).ArraySet1(_654_v0, (_index110).Int())
						_675_v15 = (_675_v15).Update((_this).F25(), (_dafny.IntOfInt64(552)).Minus((_685_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_685_v25), 0))).Int()).(_dafny.Int)))
						var _686_v26 _dafny.Sequence
						_ = _686_v26
						_686_v26 = _dafny.SeqOf((_this).F25(), (_this).F26(), false, (_this).F26())
						var _rhs71 bool = (_this).F26()
						_ = _rhs71
						var _rhs72 bool = (_686_v26).Select((Companion_Default___.SafeIndex((_this).Fm3(_679_v19, (_this).F25(), p0, (_this).F26(), globalState), _dafny.IntOfUint32((_686_v26).Cardinality()))).Uint32()).(bool)
						_ = _rhs72
						var _rhs73 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm11(globalState), p0)).Cardinality())
						_ = _rhs73
						var _lhs62 *GlobalState = globalState
						_ = _lhs62
						r0 = _rhs71
						r2 = _rhs72
						_lhs62.F6 = _rhs73
						var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_685_v25), 0))
						_ = _index111
						(_685_v25).ArraySet1((_685_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_685_v25), 0))).Int()).(_dafny.Int), (_index111).Int())
					}
					var _687_v27 D3
					_ = _687_v27
					_687_v27 = Companion_D3_.Create_DC10_(p2, _654_v0)
					var _688_v28 _dafny.MultiSet
					_ = _688_v28
					_688_v28 = _dafny.MultiSetOf((_this).F25(), (_this).F25(), (_this).F26())
					var _689_v29 _dafny.Map
					_ = _689_v29
					_689_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(72), (_this).F25())
					var _rhs74 _dafny.Int = (func() _dafny.Int {
						if (_688_v28).Contains((_this).F25()) {
							return (_688_v28).Multiplicity((_this).F25())
						}
						return (_689_v29).Cardinality()
					})()
					_ = _rhs74
					var _rhs75 D3 = _687_v27
					_ = _rhs75
					var _lhs63 *GlobalState = globalState
					_ = _lhs63
					_lhs63.F3 = _rhs74
					_687_v27 = _rhs75
					_654_v0 = (_dafny.IntOfUint32((p0).Cardinality())).Plus(_654_v0)
					(globalState).F3 = _dafny.IntOfInt64(-34)
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _690_v30 _dafny.Int
		_ = _690_v30
		_690_v30 = _dafny.IntOfInt64(687)
		var _691_v31 _dafny.Map
		_ = _691_v31
		_691_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_690_v30, !((_this).F25()))
		var _692_v32 _dafny.Sequence
		_ = _692_v32
		_692_v32 = _dafny.SeqOf((_this).F26(), (_this).F25(), false)
		var _693_v33 _dafny.Sequence
		_ = _693_v33
		_693_v33 = _dafny.SeqOf(_dafny.IntOfInt64(-878))
		var _694_v35 _dafny.Set
		_ = _694_v35
		_694_v35 = _dafny.SetOf((_this).F26(), (_this).F25(), (_this).F25(), (_this).F25())
		var _695_v36 _dafny.Array
		_ = _695_v36
		var _nwElement0_24 bool = !_dafny.Companion_Sequence_.Contains(_692_v32, (func() bool {
			if (_691_v31).Contains(_690_v30) {
				return (_691_v31).Get(_690_v30).(bool)
			}
			return (_692_v32).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_690_v30, (_693_v33).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter27 := _dafny.Iterate((_691_v31).Keys().Elements()); ; {
					_compr_26, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _696_v34 _dafny.Int
					_696_v34 = interface{}(_compr_26).(_dafny.Int)
					if (_691_v31).Contains(_696_v34) {
						_coll26.Add((_696_v34).Minus(_dafny.IntOfInt64(-759)), !((_this).F26()))
					}
				}
				return _coll26.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_693_v33).Cardinality()))).Uint32()).(_dafny.Int))).Cardinality(), _dafny.IntOfUint32((_692_v32).Cardinality()))).Uint32()).(bool)
		})())
		_ = _nwElement0_24
		var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(18))
		_ = _nw124
		(_nw124).ArraySet1(_nwElement0_24, 0)
		(_nw124).ArraySet1((_694_v35).IsDisjointFrom(_dafny.SetOf((_this).F26())), 1)
		(_nw124).ArraySet1((_this).F26(), 2)
		(_nw124).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(p0, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(2))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_697_p2 T0) func(_dafny.Int) _dafny.CodePoint {
			return func(_698_i1 _dafny.Int) _dafny.CodePoint {
				return _697_p2.F24()
			}
		})(p2))), (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(2))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_699_p2 T0) func(_dafny.Int) _dafny.CodePoint {
			return func(_700_i1 _dafny.Int) _dafny.CodePoint {
				return _699_p2.F24()
			}
		})(p2)))).Cardinality()))).Uint32(), (p0).Select((Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint))), 3)
		(_nw124).ArraySet1(!(Companion_Default___.Fm6(globalState)) || (Companion_Default___.Fm6(globalState)), 4)
		(_nw124).ArraySet1((_this).F25(), 5)
		(_nw124).ArraySet1((_this).F25(), 6)
		(_nw124).ArraySet1((_this).F25(), 7)
		(_nw124).ArraySet1((_this).F25(), 8)
		(_nw124).ArraySet1((_this).F26(), 9)
		(_nw124).ArraySet1((_this).F25(), 10)
		(_nw124).ArraySet1(((_this).F26()) == ((_this).F25()), 11)
		(_nw124).ArraySet1((Companion_Default___.Fm12((_this).F26(), false, _690_v30, (_this).F25(), globalState)).Dtor_cf8(), 12)
		(_nw124).ArraySet1((_this).F25(), 13)
		(_nw124).ArraySet1(Companion_Default___.Fm6(globalState), 14)
		(_nw124).ArraySet1((_this).F26(), 15)
		(_nw124).ArraySet1((_690_v30).Cmp(_690_v30) >= 0, 16)
		(_nw124).ArraySet1((_this).F26(), 17)
		_695_v36 = _nw124
		var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_695_v36), 0))
		_ = _index112
		(_695_v36).ArraySet1((_690_v30).Cmp(_690_v30) == 0, (_index112).Int())
		var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_695_v36), 0))
		_ = _index113
		(_695_v36).ArraySet1((_this).F25(), (_index113).Int())
		var _701_i2 _dafny.Int
		_ = _701_i2
		_701_i2 = _dafny.Zero
		{
			for !((_this).F25()) {
				{
					if (_701_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_701_i2 = (_701_i2).Plus(_dafny.One)
					var _702_v37 _dafny.Array
					_ = _702_v37
					var _nw125 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
					_ = _nw125
					_702_v37 = _nw125
					var _703_v38 *C0
					_ = _703_v38
					var _nw126 *C0 = New_C0_()
					_ = _nw126
					_nw126.Ctor__((_this).F26(), p1, p2.F24())
					_703_v38 = _nw126
					var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_702_v37), 0))
					_ = _index114
					(_702_v37).ArraySet1(_703_v38, (_index114).Int())
					var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_702_v37), 0))
					_ = _index115
					(_702_v37).ArraySet1(_703_v38, (_index115).Int())
					var _704_v39 _dafny.Set
					_ = _704_v39
					_704_v39 = _dafny.SetOf(_690_v30)
					var _705_v40 _dafny.Map
					_ = _705_v40
					_705_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_704_v39).Cardinality()), _dafny.Companion_Sequence_.Concatenate(p0, p0))
					var _706_v41 _dafny.MultiSet
					_ = _706_v41
					_706_v41 = _dafny.MultiSetOf(true)
					(globalState).F1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if (_705_v40).Contains(((_706_v41).Intersection(_706_v41)).Cardinality()) {
							return (_705_v40).Get(((_706_v41).Intersection(_706_v41)).Cardinality()).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _this.F24())
					})(), (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_705_v40).Contains(((_706_v41).Intersection(_706_v41)).Cardinality()) {
							return (_705_v40).Get(((_706_v41).Intersection(_706_v41)).Cardinality()).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _this.F24())
					})()).Cardinality()))).Uint32(), _this.F24()), (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if (_705_v40).Contains(((_706_v41).Intersection(_706_v41)).Cardinality()) {
							return (_705_v40).Get(((_706_v41).Intersection(_706_v41)).Cardinality()).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _this.F24())
					})(), (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_705_v40).Contains(((_706_v41).Intersection(_706_v41)).Cardinality()) {
							return (_705_v40).Get(((_706_v41).Intersection(_706_v41)).Cardinality()).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_690_v30, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _this.F24())
					})()).Cardinality()))).Uint32(), _this.F24())).Cardinality()))).Uint32(), Companion_Default___.Fm13(globalState))
					(globalState).F6 = ((_dafny.Zero).Minus(_690_v30)).Plus(_690_v30)
					var _707_v42 _dafny.Array
					_ = _707_v42
					var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
					_ = _nw127
					_707_v42 = _nw127
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_707_v42), 0))
					_ = _index116
					(_707_v42).ArraySet1(_693_v33, (_index116).Int())
					var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_707_v42), 0))
					_ = _index117
					(_707_v42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_693_v33, _693_v33), (_index117).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _708_v43 _dafny.MultiSet
		_ = _708_v43
		_708_v43 = _dafny.MultiSetOf(_690_v30, _690_v30, _690_v30)
		var _709_v44 _dafny.Map
		_ = _709_v44
		_709_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
			if (_695_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_695_v36), 0))).Int()).(bool) {
				return _708_v43
			}
			return _708_v43
		})(), (_this).F26())
		_709_v44 = (_709_v44).Update(Companion_Default___.Fm10(globalState), (_695_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_695_v36), 0))).Int()).(bool))
		var _710_v45 _dafny.Array
		_ = _710_v45
		var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw128
		_710_v45 = _nw128
		for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_710_v45), 0))); ; {
			_guard_loop_1, _ok28 := _iter28()
			if !_ok28 {
				break
			}
			var _711_i3 _dafny.Int
			_711_i3 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_711_i3).Sign() != -1) && ((_711_i3).Cmp(_dafny.ArrayLen((_710_v45), 0)) < 0)) {
				(_710_v45).ArraySet1(Companion_Default___.SafeDivisionInt(_711_i3, _690_v30), (_711_i3).Int())
			}
		}
		var _712_v46 _dafny.Map
		_ = _712_v46
		_712_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_690_v30, _dafny.IntOfInt64(700))
		(globalState).F6 = ((_712_v46).Cardinality()).Plus(_690_v30)
		r0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(296))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_713_p2 T0) func(_dafny.Int) _dafny.CodePoint {
			return func(_714_i4 _dafny.Int) _dafny.CodePoint {
				return _713_p2.F24()
			}
		})(p2))), p0)
		var _715_v47 _dafny.MultiSet
		_ = _715_v47
		_715_v47 = _dafny.MultiSetOf((_this).F26())
		var _716_v48 _dafny.Sequence
		_ = _716_v48
		_716_v48 = _dafny.SeqOf(_715_v47, _715_v47, _715_v47, _715_v47)
		var _717_v49 _dafny.Map
		_ = _717_v49
		_717_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_716_v48, _695_v36)
		r1 = (func() _dafny.Array {
			if (_717_v49).Contains(_716_v48) {
				return (_717_v49).Get(_716_v48).(_dafny.Array)
			}
			return _695_v36
		})()
		var _718_v50 D3
		_ = _718_v50
		_718_v50 = Companion_D3_.Create_DC8_((_this).F26(), _690_v30, _690_v30)
		r2 = (_718_v50).Dtor_cf8()
		r3 = _dafny.IntOfInt64(8)
		return r0, r1, r2, r3
	}
}
func (_this *C1) F25() bool {
	{
		return _this._f25
	}
}
func (_this *C1) F26() bool {
	{
		return _this._f26
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	dummy byte
}

func New_C2_() *C2 {
	_this := C2{}

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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__() {
	{
	}
}
func (_this *C2) Fm0(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('h')
	}
}
func (_this *C2) Fm1(globalState *GlobalState) bool {
	{
		var _source9 D0 = Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())
		_ = _source9
		if _source9.Is_DC0() {
			var _719___mcc_h0 _dafny.Int = _source9.Get_().(D0_DC0).Cf0
			_ = _719___mcc_h0
			var _720_cf0 _dafny.Int = _719___mcc_h0
			_ = _720_cf0
			return !((_720_cf0).Cmp(_720_cf0) >= 0)
		} else if _source9.Is_DC1() {
			var _721___mcc_h1 _dafny.Map = _source9.Get_().(D0_DC1).Cf1
			_ = _721___mcc_h1
			var _722_cf1 _dafny.Map = _721___mcc_h1
			_ = _722_cf1
			return false
		} else {
			var _723___mcc_h2 bool = _source9.Get_().(D0_DC2).Cf2
			_ = _723___mcc_h2
			var _724_cf2 bool = _723___mcc_h2
			_ = _724_cf2
			return _724_cf2
		}
	}
}
func (_this *C2) M0(p0 _dafny.Array, p1 _dafny.Int, p2 bool, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		r1 = p2
		(globalState).F6 = p1
		var _725_v0 D0
		_ = _725_v0
		_725_v0 = Companion_D0_.Create_DC0_(p1)
		var _726_v1 _dafny.CodePoint
		_ = _726_v1
		_726_v1 = _dafny.CodePoint('d')
		var _727_v2 _dafny.Map
		_ = _727_v2
		_727_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SetOf(_726_v1))
		var _728_v3 _dafny.Set
		_ = _728_v3
		_728_v3 = _dafny.SetOf(p2, p2, p2)
		var _729_v4 _dafny.Set
		_ = _729_v4
		_729_v4 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_727_v2).Cardinality())), (_728_v3).Cardinality())
		var _730_v5 _dafny.Array
		_ = _730_v5
		var _nwElement0_25 bool = !((_729_v4).IsSubsetOf(_729_v4))
		_ = _nwElement0_25
		var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(2))
		_ = _nw129
		(_nw129).ArraySet1(_nwElement0_25, 0)
		(_nw129).ArraySet1((p2) && (false), 1)
		_730_v5 = _nw129
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
		_ = _index118
		(_730_v5).ArraySet1(true, (_index118).Int())
		var _pat_let_tv15 = p1
		_ = _pat_let_tv15
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
		_ = _index119
		var _rhs76 D0 = func(_pat_let15_0 D0) D0 {
			return func(_731_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let16_0 _dafny.Int) D0 {
					return func(_732_dt__update_hcf0_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_732_dt__update_hcf0_h0)
					}(_pat_let16_0)
				}(_pat_let_tv15)
			}(_pat_let15_0)
		}(_725_v0)
		_ = _rhs76
		var _rhs77 bool = (func() bool {
			if (p1).Cmp(p1) != 0 {
				return !(p2)
			}
			return false
		})()
		_ = _rhs77
		var _rhs78 _dafny.Int = p1
		_ = _rhs78
		var _lhs64 _dafny.Array = _730_v5
		_ = _lhs64
		var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
		_ = _lhs65
		var _lhs66 *GlobalState = globalState
		_ = _lhs66
		_725_v0 = _rhs76
		(_lhs64).ArraySet1(_rhs77, (_lhs65).Int())
		_lhs66.F6 = _rhs78
		var _733_i0 _dafny.Int
		_ = _733_i0
		_733_i0 = _dafny.Zero
		{
			for (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool) {
				{
					if (_733_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_733_i0 = (_733_i0).Plus(_dafny.One)
					var _734_v6 _dafny.Sequence
					_ = _734_v6
					_734_v6 = _dafny.SeqOf((_729_v4).Cardinality())
					_734_v6 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(556))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg27 _dafny.Int) interface{} {
							return coer27(arg27)
						}
					}(func(_735_i1 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(317)
					})), _734_v6)
					var _736_v7 D0
					_ = _736_v7
					_736_v7 = Companion_D0_.Create_DC2_(true)
					var _pat_let_tv16 = p2
					_ = _pat_let_tv16
					var _source10 D0 = func(_pat_let17_0 D0) D0 {
						return func(_737_dt__update__tmp_h1 D0) D0 {
							return func(_pat_let18_0 bool) D0 {
								return func(_738_dt__update_hcf2_h0 bool) D0 {
									return Companion_D0_.Create_DC2_(_738_dt__update_hcf2_h0)
								}(_pat_let18_0)
							}(_pat_let_tv16)
						}(_pat_let17_0)
					}(_736_v7)
					_ = _source10
					if _source10.Is_DC0() {
						var _739___mcc_h0 _dafny.Int = _source10.Get_().(D0_DC0).Cf0
						_ = _739___mcc_h0
						var _740_cf0 _dafny.Int = _739___mcc_h0
						_ = _740_cf0
						(globalState).F6 = _740_cf0
						var _741_v8 _dafny.Sequence
						_ = _741_v8
						_741_v8 = _dafny.SeqOf(true)
						var _742_v9 D0
						_ = _742_v9
						_742_v9 = Companion_D0_.Create_DC1_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_740_cf0, _741_v8))
						var _743_v10 D0
						_ = _743_v10
						_743_v10 = Companion_D0_.Create_DC1_((_742_v9).Dtor_cf1())
						var _rhs79 bool = p2
						_ = _rhs79
						var _rhs80 D0 = (func() D0 {
							if p2 {
								return _725_v0
							}
							return _725_v0
						})()
						_ = _rhs80
						var _rhs81 _dafny.Int = Companion_Default___.Fm2(_dafny.IntOfUint32((_741_v8).Cardinality()), _742_v9, p2, (func() bool {
							if (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool) {
								return !(!(!(p2)))
							}
							return !(p2)
						})(), globalState)
						_ = _rhs81
						var _rhs82 _dafny.Int = _740_cf0
						_ = _rhs82
						var _lhs67 *GlobalState = globalState
						_ = _lhs67
						r0 = _rhs79
						_725_v0 = _rhs80
						_740_cf0 = _rhs81
						_lhs67.F0 = _rhs82
						(globalState).F0 = p1
						_740_cf0 = _740_cf0
					} else if _source10.Is_DC1() {
						var _744___mcc_h1 _dafny.Map = _source10.Get_().(D0_DC1).Cf1
						_ = _744___mcc_h1
						var _745_cf1 _dafny.Map = _744___mcc_h1
						_ = _745_cf1
						var _746_v11 _dafny.Array
						_ = _746_v11
						var _len0_18 _dafny.Int = _dafny.IntOfInt64(21)
						_ = _len0_18
						var _nw130 _dafny.Array
						_ = _nw130
						if _len0_18.Cmp(_dafny.Zero) == 0 {
							_nw130 = _dafny.NewArray(_len0_18)
						} else {
							var _init18 func(_dafny.Int) _dafny.Int = (func(_747_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_748_i2 _dafny.Int) _dafny.Int {
									return (_748_i2).Minus(_747_p1)
								}
							})(p1)
							_ = _init18
							var _element0_18 = _init18(_dafny.Zero)
							_ = _element0_18
							_nw130 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
							(_nw130).ArraySet1(_element0_18, 0)
							var _nativeLen0_18 = (_len0_18).Int()
							_ = _nativeLen0_18
							for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
								(_nw130).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
							}
						}
						_746_v11 = _nw130
						var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_746_v11), 0))
						_ = _index120
						(_746_v11).ArraySet1(p1, (_index120).Int())
						var _749_v12 _dafny.MultiSet
						_ = _749_v12
						_749_v12 = _dafny.MultiSetOf(p2)
						var _750_v13 _dafny.Map
						_ = _750_v13
						_750_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_749_v12).Cardinality())
						var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_746_v11), 0))
						_ = _index121
						(_746_v11).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
							if (_750_v13).Contains((_this).Fm1(globalState)) {
								return (_750_v13).Get((_this).Fm1(globalState)).(_dafny.Int)
							}
							return p1
						})()), (_index121).Int())
						var _751_v14 _dafny.Array
						_ = _751_v14
						var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.One)
						_ = _nw131
						_751_v14 = _nw131
						var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_751_v14), 0))
						_ = _index122
						(_751_v14).ArraySet1(_729_v4, (_index122).Int())
						var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_751_v14), 0))
						_ = _index123
						(_751_v14).ArraySet1(_dafny.SetOf((_746_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_746_v11), 0))).Int()).(_dafny.Int)), (_index123).Int())
						var _rhs83 bool = ((_749_v12).Union(_749_v12)).IsProperSubsetOf(_749_v12)
						_ = _rhs83
						var _rhs84 _dafny.Int = p1
						_ = _rhs84
						var _lhs68 *GlobalState = globalState
						_ = _lhs68
						r1 = _rhs83
						_lhs68.F0 = _rhs84
						var _752_v15 _dafny.Sequence
						_ = _752_v15
						_752_v15 = _dafny.SeqOf(p2, p2)
						r0 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_752_v15, _752_v15), _752_v15)
					} else {
						var _753___mcc_h2 bool = _source10.Get_().(D0_DC2).Cf2
						_ = _753___mcc_h2
						var _754_cf2 bool = _753___mcc_h2
						_ = _754_cf2
						r1 = (func() bool {
							if p2 {
								return (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool)
							}
							return _754_cf2
						})()
						var _755_v16 _dafny.Array
						_ = _755_v16
						var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
						_ = _nw132
						_755_v16 = _nw132
						var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_755_v16), 0))
						_ = _index124
						(_755_v16).ArraySet1(_dafny.IntOfInt64(229), (_index124).Int())
						var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_755_v16), 0))
						_ = _index125
						(_755_v16).ArraySet1(p1, (_index125).Int())
						var _756_v17 _dafny.Sequence
						_ = _756_v17
						_756_v17 = _dafny.UnicodeSeqOfUtf8Bytes("xjmfx")
						var _757_v18 bool
						_ = _757_v18
						var _out34 bool
						_ = _out34
						_out34 = (_this).M1(_730_v5, !(_dafny.Companion_Sequence_.IsPrefixOf(_756_v17, _756_v17)), p1, globalState)
						_757_v18 = _out34
						var _758_v19 _dafny.Array
						_ = _758_v19
						var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(9))
						_ = _nw133
						_758_v19 = _nw133
						var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_758_v19), 0))
						_ = _index126
						(_758_v19).ArraySet1CodePoint(_726_v1, (_index126).Int())
						var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_758_v19), 0))
						_ = _index127
						var _rhs85 _dafny.CodePoint = _726_v1
						_ = _rhs85
						var _rhs86 _dafny.Int = p1
						_ = _rhs86
						var _rhs87 _dafny.Int = p1
						_ = _rhs87
						var _rhs88 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_756_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg28 _dafny.Int) interface{} {
								return coer28(arg28)
							}
						}(func(_759_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('x')
						}))), (Companion_D2_.Create_DC4_(_756_v17)).Dtor_cf4())).Cardinality())
						_ = _rhs88
						var _lhs69 _dafny.Array = _758_v19
						_ = _lhs69
						var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_758_v19), 0))
						_ = _lhs70
						var _lhs71 *GlobalState = globalState
						_ = _lhs71
						var _lhs72 *GlobalState = globalState
						_ = _lhs72
						var _lhs73 *GlobalState = globalState
						_ = _lhs73
						(_lhs69).ArraySet1CodePoint(_rhs85, (_lhs70).Int())
						_lhs71.F0 = _rhs86
						_lhs72.F6 = _rhs87
						_lhs73.F0 = _rhs88
					}
					var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
					_ = _index128
					(_730_v5).ArraySet1((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), (_index128).Int())
					r0 = (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool)
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _760_v20 D0
		_ = _760_v20
		_760_v20 = Companion_D0_.Create_DC2_(p2)
		var _761_v21 _dafny.MultiSet
		_ = _761_v21
		_761_v21 = _dafny.MultiSetOf(_760_v20)
		_761_v21 = (_761_v21).Union(_dafny.MultiSetOf(_760_v20, _760_v20, _760_v20, func(_pat_let19_0 D0) D0 {
			return func(_762_dt__update__tmp_h2 D0) D0 {
				return func(_pat_let20_0 bool) D0 {
					return func(_763_dt__update_hcf2_h1 bool) D0 {
						return Companion_D0_.Create_DC2_(_763_dt__update_hcf2_h1)
					}(_pat_let20_0)
				}(true)
			}(_pat_let19_0)
		}(_760_v20), _760_v20))
		if p2 {
			var _764_v22 _dafny.Sequence
			_ = _764_v22
			_764_v22 = _dafny.SeqOf((p1).Cmp(p1) >= 0, (_729_v4).IsProperSubsetOf(_729_v4), (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))
			_764_v22 = _764_v22
			var _765_v23 _dafny.Map
			_ = _765_v23
			_765_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
			var _766_v24 _dafny.Sequence
			_ = _766_v24
			_766_v24 = _dafny.SeqOf(_765_v23, _765_v23)
			var _767_v25 _dafny.Sequence
			_ = _767_v25
			_767_v25 = _dafny.SeqOf(_766_v24, _dafny.SeqOf(_765_v23))
			if _dafny.Companion_Sequence_.IsPrefixOf((_767_v25).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_767_v25).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_765_v23, _765_v23), _766_v24)) {
				var _768_v26 _dafny.Sequence
				_ = _768_v26
				_768_v26 = _dafny.SeqOf(_726_v1, _726_v1)
				var _769_v27 _dafny.Map
				_ = _769_v27
				_769_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), _dafny.CodePoint('s')))
				var _770_v28 _dafny.Map
				_ = _770_v28
				_770_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), _726_v1)
				var _771_v29 _dafny.Map
				_ = _771_v29
				_771_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() _dafny.Map {
					if (_769_v27).Contains(p1) {
						return (_769_v27).Get(p1).(_dafny.Map)
					}
					return _770_v28
				})())
				var _772_v30 *C1
				_ = _772_v30
				var _nw134 *C1 = New_C1_()
				_ = _nw134
				_nw134.Ctor__((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), Companion_Default___.Fm6(globalState), (_768_v26).Select((Companion_Default___.SafeIndex((_769_v27).Cardinality(), _dafny.IntOfUint32((_768_v26).Cardinality()))).Uint32()).(_dafny.CodePoint))
				_772_v30 = _nw134
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
				_ = _index129
				(_730_v5).ArraySet1(!((_772_v30).F26()), (_index129).Int())
				(globalState).F9 = _726_v1
				(globalState).F0 = (_dafny.Zero).Minus((_dafny.IntOfInt64(267)).Plus(p1))
				var _773_v31 _dafny.Map
				_ = _773_v31
				_773_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p1, p1), !((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool)))
				_773_v31 = (_773_v31).Update(p1, (_772_v30).F25())
			} else {
				var _774_v32 _dafny.Sequence
				_ = _774_v32
				_774_v32 = _dafny.SeqOf(p1)
				var _775_v33 _dafny.Map
				_ = _775_v33
				_775_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v32, p2)
				var _776_v34 _dafny.MultiSet
				_ = _776_v34
				_776_v34 = _dafny.MultiSetOf(p1, _dafny.IntOfInt64(718), _dafny.IntOfInt64(-937), p1, p1)
				var _777_v35 _dafny.Map
				_ = _777_v35
				_777_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_v1, _776_v34)
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
				_ = _index130
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
				_ = _index131
				var _rhs89 _dafny.Int = (_725_v0).Dtor_cf0()
				_ = _rhs89
				var _rhs90 bool = p2
				_ = _rhs90
				var _rhs91 bool = (func() bool {
					if (_775_v33).Contains(_dafny.Companion_Sequence_.Concatenate(_774_v32, _774_v32)) {
						return (_775_v33).Get(_dafny.Companion_Sequence_.Concatenate(_774_v32, _774_v32)).(bool)
					}
					return !((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))
				})()
				_ = _rhs91
				var _rhs92 bool = (_dafny.MultiSetOf(p1)).IsProperSubsetOf((func() _dafny.MultiSet {
					if (_777_v35).Contains(_726_v1) {
						return (_777_v35).Get(_726_v1).(_dafny.MultiSet)
					}
					return (_776_v34).Update(_dafny.IntOfInt64(157), Companion_Default___.Abs(_dafny.IntOfInt64(310)))
				})())
				_ = _rhs92
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				var _lhs75 _dafny.Array = _730_v5
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
				_ = _lhs76
				var _lhs77 _dafny.Array = _730_v5
				_ = _lhs77
				var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
				_ = _lhs78
				_lhs74.F6 = _rhs89
				(_lhs75).ArraySet1(_rhs90, (_lhs76).Int())
				(_lhs77).ArraySet1(_rhs91, (_lhs78).Int())
				r1 = _rhs92
				var _778_v36 _dafny.Sequence
				_ = _778_v36
				_778_v36 = _dafny.UnicodeSeqOfUtf8Bytes("gg")
				var _779_v37 _dafny.Array
				_ = _779_v37
				var _nwElement0_26 _dafny.Int = p1
				_ = _nwElement0_26
				var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(18))
				_ = _nw135
				(_nw135).ArraySet1(_nwElement0_26, 0)
				(_nw135).ArraySet1(p1, 1)
				(_nw135).ArraySet1(p1, 2)
				(_nw135).ArraySet1(p1, 3)
				(_nw135).ArraySet1((_dafny.Zero).Minus(p1), 4)
				(_nw135).ArraySet1(_dafny.IntOfInt64(239), 5)
				(_nw135).ArraySet1(p1, 6)
				(_nw135).ArraySet1(p1, 7)
				(_nw135).ArraySet1((_765_v23).Cardinality(), 8)
				(_nw135).ArraySet1(p1, 9)
				(_nw135).ArraySet1(p1, 10)
				(_nw135).ArraySet1(p1, 11)
				(_nw135).ArraySet1(p1, 12)
				(_nw135).ArraySet1(p1, 13)
				(_nw135).ArraySet1(p1, 14)
				(_nw135).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_778_v36).Cardinality()))).Cardinality()), 15)
				(_nw135).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p1)), 16)
				(_nw135).ArraySet1(p1, 17)
				_779_v37 = _nw135
				var _780_v38 _dafny.Map
				_ = _780_v38
				_780_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_v1, _779_v37)
				var _781_v39 _dafny.Map
				_ = _781_v39
				_781_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_764_v22, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_v1, _779_v37))
				_780_v38 = (((func() _dafny.Map {
					if (_781_v39).Contains(_764_v22) {
						return (_781_v39).Get(_764_v22).(_dafny.Map)
					}
					return _780_v38
				})()).Merge(_780_v38)).Update(_726_v1, _779_v37)
				var _782_v40 T0
				_ = _782_v40
				var _nw136 *C1 = New_C1_()
				_ = _nw136
				_nw136.Ctor__(p2, p2, _726_v1)
				_782_v40 = _nw136
				_782_v40 = _782_v40
				var _783_v41 D3
				_ = _783_v41
				_783_v41 = Companion_D3_.Create_DC9_((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))
				_783_v41 = (func() D3 {
					if (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool) {
						return (func() D3 {
							if true {
								return _783_v41
							}
							return _783_v41
						})()
					}
					return Companion_D3_.Create_DC9_((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))
				})()
				var _784_v42 _dafny.Map
				_ = _784_v42
				_784_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))
				_784_v42 = (_784_v42).Update(p1, !(p2))
			}
			if p2 {
				var _785_v43 _dafny.MultiSet
				_ = _785_v43
				_785_v43 = _dafny.MultiSetOf(p1)
				var _786_v44 _dafny.Array
				_ = _786_v44
				var _nw137 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw137
				_786_v44 = _nw137
				var _787_v45 *C0
				_ = _787_v45
				var _nw138 *C0 = New_C0_()
				_ = _nw138
				_nw138.Ctor__(!((_785_v43).IsDisjointFrom((_785_v43).Update(p1, Companion_Default___.Abs(_dafny.IntOfUint32((_764_v22).Cardinality()))))), _786_v44, _726_v1)
				_787_v45 = _nw138
				_785_v43 = _785_v43
				var _788_v46 _dafny.Sequence
				_ = _788_v46
				_788_v46 = _dafny.UnicodeSeqOfUtf8Bytes("iytixmgk")
				var _789_v47 _dafny.MultiSet
				_ = _789_v47
				_789_v47 = _dafny.MultiSetOf(_788_v46, _dafny.UnicodeSeqOfUtf8Bytes("ehdle"), _788_v46)
				_789_v47 = (_789_v47).Intersection(_789_v47)
				r0 = (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool)
				var _790_v48 _dafny.Array
				_ = _790_v48
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(8))
				_ = _nw139
				_790_v48 = _nw139
				var _791_v49 _dafny.Sequence
				_ = _791_v49
				_791_v49 = _dafny.SeqOf(p1)
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_790_v48), 0))
				_ = _index132
				(_790_v48).ArraySet1(Companion_D2_.Create_DC6_(_791_v49), (_index132).Int())
				var _792_v50 D2
				_ = _792_v50
				_792_v50 = Companion_D2_.Create_DC6_(_dafny.SeqOf(p1, _dafny.IntOfInt64(604)))
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_790_v48), 0))
				_ = _index133
				var _rhs93 _dafny.Int = p1
				_ = _rhs93
				var _rhs94 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm11(globalState), _788_v46)
				_ = _rhs94
				var _rhs95 D2 = _792_v50
				_ = _rhs95
				var _rhs96 _dafny.Int = p1
				_ = _rhs96
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				var _lhs81 _dafny.Array = _790_v48
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_790_v48), 0))
				_ = _lhs82
				var _lhs83 *GlobalState = globalState
				_ = _lhs83
				_lhs79.F6 = _rhs93
				_lhs80.F1 = _rhs94
				(_lhs81).ArraySet1(_rhs95, (_lhs82).Int())
				_lhs83.F0 = _rhs96
			} else {
				r0 = false
				r1 = p2
				r1 = !((_this).Fm1(globalState))
				var _793_v51 _dafny.Sequence
				_ = _793_v51
				_793_v51 = _dafny.SeqOf(p1, p1)
				var _794_v52 D2
				_ = _794_v52
				_794_v52 = Companion_D2_.Create_DC6_(_dafny.Companion_Sequence_.Concatenate(_793_v51, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}((func(_795_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_796_i4 _dafny.Int) _dafny.Int {
						return _795_p1
					}
				})(p1)))))
				_794_v52 = _794_v52
				var _797_v53 _dafny.Sequence
				_ = _797_v53
				_797_v53 = _dafny.UnicodeSeqOfUtf8Bytes("bmi")
				var _798_v54 D2
				_ = _798_v54
				_798_v54 = Companion_D2_.Create_DC4_(_797_v53)
				var _799_v55 _dafny.Map
				_ = _799_v55
				_799_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _764_v22)
				var _800_v56 D0
				_ = _800_v56
				_800_v56 = Companion_D0_.Create_DC1_(_799_v55)
				var _801_v57 _dafny.Array
				_ = _801_v57
				var _nwElement0_27 _dafny.Int = Companion_Default___.Fm2(_dafny.IntOfInt64(896), _800_v56, (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), p2, globalState)
				_ = _nwElement0_27
				var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(5))
				_ = _nw140
				(_nw140).ArraySet1(_nwElement0_27, 0)
				(_nw140).ArraySet1(p1, 1)
				(_nw140).ArraySet1(p1, 2)
				(_nw140).ArraySet1(p1, 3)
				(_nw140).ArraySet1(_dafny.IntOfUint32((_793_v51).Cardinality()), 4)
				_801_v57 = _nw140
				var _802_v58 _dafny.Array
				_ = _802_v58
				var _nwElement0_28 _dafny.Array = _801_v57
				_ = _nwElement0_28
				var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(7))
				_ = _nw141
				(_nw141).ArraySet1(_nwElement0_28, 0)
				(_nw141).ArraySet1(_801_v57, 1)
				(_nw141).ArraySet1(_801_v57, 2)
				(_nw141).ArraySet1(_801_v57, 3)
				(_nw141).ArraySet1(_801_v57, 4)
				(_nw141).ArraySet1(_801_v57, 5)
				(_nw141).ArraySet1(_801_v57, 6)
				_802_v58 = _nw141
				var _803_v59 *C0
				_ = _803_v59
				var _nw142 *C0 = New_C0_()
				_ = _nw142
				_nw142.Ctor__(p2, _802_v58, _726_v1)
				_803_v59 = _nw142
				var _804_v60 _dafny.Map
				_ = _804_v60
				_804_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D2 {
					if !((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool)) {
						return _798_v54
					}
					return Companion_D2_.Create_DC4_(_797_v53)
				})(), _803_v59)
				_804_v60 = (_804_v60).Update(_798_v54, (func() *C0 {
					if _803_v59.F27 {
						return _803_v59
					}
					return _803_v59
				})())
			}
			var _805_v61 _dafny.Sequence
			_ = _805_v61
			_805_v61 = _dafny.UnicodeSeqOfUtf8Bytes("skq")
			var _806_v62 _dafny.Map
			_ = _806_v62
			_806_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_805_v61).Cardinality()), _dafny.IntOfInt64(809))
			var _807_v63 _dafny.Map
			_ = _807_v63
			_807_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool) {
					return p1
				}
				return p1
			})(), (_806_v62).Cardinality())
			_806_v62 = (_806_v62).Update((_729_v4).Cardinality(), _dafny.IntOfUint32((_805_v61).Cardinality()))
			var _808_v64 _dafny.Array
			_ = _808_v64
			var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw143
			_808_v64 = _nw143
			var _809_v65 *C0
			_ = _809_v65
			var _nw144 *C0 = New_C0_()
			_ = _nw144
			_nw144.Ctor__((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), _808_v64, _726_v1)
			_809_v65 = _nw144
		} else {
			var _810_v66 _dafny.MultiSet
			_ = _810_v66
			_810_v66 = _dafny.MultiSetOf(p2, (_728_v3).IsProperSubsetOf(_dafny.SetOf(p2, (_this).Fm1(globalState), (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))))
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
			_ = _index134
			var _rhs97 bool = (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool)
			_ = _rhs97
			var _rhs98 _dafny.MultiSet = Companion_Default___.Fm14(globalState)
			_ = _rhs98
			var _lhs84 _dafny.Array = _730_v5
			_ = _lhs84
			var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
			_ = _lhs85
			(_lhs84).ArraySet1(_rhs97, (_lhs85).Int())
			_810_v66 = _rhs98
			var _811_v67 _dafny.Array
			_ = _811_v67
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw145
			_811_v67 = _nw145
			_811_v67 = _811_v67
			r0 = (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool)
			if (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool) {
				var _812_v68 _dafny.Sequence
				_ = _812_v68
				_812_v68 = _dafny.SeqOf(p2, (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))
				_ = _index135
				(_811_v67).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_812_v68).Cardinality())), (_index135).Int())
				var _813_v69 _dafny.MultiSet
				_ = _813_v69
				_813_v69 = _dafny.MultiSetOf(p1)
				var _814_v70 _dafny.Sequence
				_ = _814_v70
				_814_v70 = _dafny.UnicodeSeqOfUtf8Bytes("qfprsk")
				var _815_v71 D2
				_ = _815_v71
				_815_v71 = Companion_D2_.Create_DC5_(_dafny.IntOfUint32((_814_v70).Cardinality()))
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))
				_ = _index136
				var _rhs99 _dafny.Int = ((_813_v69).Cardinality()).Minus((_815_v71).Dtor_cf5())
				_ = _rhs99
				var _rhs100 bool = !((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool))
				_ = _rhs100
				var _lhs86 _dafny.Array = _811_v67
				_ = _lhs86
				var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))
				_ = _lhs87
				(_lhs86).ArraySet1(_rhs99, (_lhs87).Int())
				r0 = _rhs100
				var _816_v72 _dafny.Array
				_ = _816_v72
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw146
				_816_v72 = _nw146
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_816_v72), 0))
				_ = _index137
				(_816_v72).ArraySet1(_730_v5, (_index137).Int())
				var _817_v73 _dafny.Map
				_ = _817_v73
				_817_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_811_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))).Int()).(_dafny.Int)), _812_v68)
				var _818_v74 _dafny.Array
				_ = _818_v74
				var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw147
				_818_v74 = _nw147
				var _819_v75 _dafny.Array
				_ = _819_v75
				var _nwElement0_29 _dafny.Array = _818_v74
				_ = _nwElement0_29
				var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(13))
				_ = _nw148
				(_nw148).ArraySet1(_nwElement0_29, 0)
				(_nw148).ArraySet1(_818_v74, 1)
				(_nw148).ArraySet1(_818_v74, 2)
				(_nw148).ArraySet1(_818_v74, 3)
				(_nw148).ArraySet1(_818_v74, 4)
				(_nw148).ArraySet1(_818_v74, 5)
				(_nw148).ArraySet1(_818_v74, 6)
				(_nw148).ArraySet1(_818_v74, 7)
				(_nw148).ArraySet1(_818_v74, 8)
				(_nw148).ArraySet1(_818_v74, 9)
				(_nw148).ArraySet1(_818_v74, 10)
				(_nw148).ArraySet1(_818_v74, 11)
				(_nw148).ArraySet1(_818_v74, 12)
				_819_v75 = _nw148
				var _820_v76 T0
				_ = _820_v76
				var _nw149 *C0 = New_C0_()
				_ = _nw149
				_nw149.Ctor__((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), _819_v75, _726_v1)
				_820_v76 = _nw149
				var _821_v77 _dafny.Set
				_ = _821_v77
				_821_v77 = _dafny.SetOf(_820_v76, _820_v76, _820_v76, _820_v76)
				var _822_v78 _dafny.Map
				_ = _822_v78
				_822_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _821_v77)
				var _823_v79 D0
				_ = _823_v79
				_823_v79 = Companion_D0_.Create_DC1_(_817_v73)
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_816_v72), 0))
				_ = _index138
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))
				_ = _index139
				var _rhs101 _dafny.Array = _730_v5
				_ = _rhs101
				var _rhs102 _dafny.Int = (p1).Times((_811_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))).Int()).(_dafny.Int))
				_ = _rhs102
				var _rhs103 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm2((_729_v4).Cardinality(), Companion_D0_.Create_DC1_(_817_v73), p2, (_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), globalState), Companion_Default___.Fm8((_811_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_814_v70).Cardinality()), !(Companion_Default___.Fm6(globalState)), globalState))
				_ = _rhs103
				var _rhs104 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(293)).Plus(Companion_Default___.Fm2(((func() _dafny.Set {
					if (_822_v78).Contains(p2) {
						return (_822_v78).Get(p2).(_dafny.Set)
					}
					return _821_v77
				})()).Cardinality(), _823_v79, false, p2, globalState)), (_dafny.Zero).Minus(_dafny.IntOfInt64(-493)))
				_ = _rhs104
				var _lhs88 _dafny.Array = _816_v72
				_ = _lhs88
				var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_816_v72), 0))
				_ = _lhs89
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				var _lhs92 _dafny.Array = _811_v67
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))
				_ = _lhs93
				(_lhs88).ArraySet1(_rhs101, (_lhs89).Int())
				_lhs90.F3 = _rhs102
				_lhs91.F6 = _rhs103
				(_lhs92).ArraySet1(_rhs104, (_lhs93).Int())
				var _824_v80 bool
				_ = _824_v80
				var _out35 bool
				_ = _out35
				_out35 = (_this).M1(_dafny.ArrayCastTo((_816_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_816_v72), 0))).Int())), (_729_v4).IsDisjointFrom(Companion_Default___.Fm15((_811_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))).Int()).(_dafny.Int), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_814_v70).Cardinality())), (_811_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))).Int()).(_dafny.Int)), p1, globalState)), (p1).Times((_729_v4).Cardinality()), globalState)
				_824_v80 = _out35
				(_820_v76).F24_set_(Companion_Default___.Fm13(globalState))
				var _825_v81 _dafny.Map
				_ = _825_v81
				_825_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_811_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_811_v67), 0))).Int()).(_dafny.Int), _824_v80)
				var _826_v82 _dafny.Map
				_ = _826_v82
				_826_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_814_v70, Companion_Default___.SafeDivisionInt((_825_v81).Cardinality(), (_729_v4).Cardinality()))
				_826_v82 = (_826_v82).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-428))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_827_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_828_i5 _dafny.Int) _dafny.CodePoint {
						return _827_v1
					}
				})(_726_v1))), p1)
			} else {
				var _829_v83 T0
				_ = _829_v83
				var _nw150 *C1 = New_C1_()
				_ = _nw150
				_nw150.Ctor__((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), false, _726_v1)
				_829_v83 = _nw150
				var _830_v84 _dafny.MultiSet
				_ = _830_v84
				_830_v84 = _dafny.MultiSetOf(_829_v83)
				var _831_v85 _dafny.Sequence
				_ = _831_v85
				_831_v85 = _dafny.UnicodeSeqOfUtf8Bytes("mm")
				r1 = (p1).Cmp((func() _dafny.Int {
					if (_830_v84).Contains(_829_v83) {
						return (_830_v84).Multiplicity(_829_v83)
					}
					return _dafny.IntOfUint32((_831_v85).Cardinality())
				})()) >= 0
				var _832_v86 _dafny.Array
				_ = _832_v86
				var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw151
				_832_v86 = _nw151
				var _833_v87 *C0
				_ = _833_v87
				var _nw152 *C0 = New_C0_()
				_ = _nw152
				_nw152.Ctor__((_730_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))).Int()).(bool), _832_v86, _829_v83.F24())
				_833_v87 = _nw152
				(globalState).F1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("q"), _831_v85), (Companion_Default___.SafeIndex((_725_v0).Dtor_cf0(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("q"), _831_v85)).Cardinality()))).Uint32(), _726_v1)
				var _834_v89 _dafny.Sequence
				_ = _834_v89
				_834_v89 = _dafny.SeqOf(p1)
				var _835_v90 _dafny.Set
				_ = _835_v90
				_835_v90 = _dafny.SetOf(func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter29 := _dafny.Iterate((_834_v89).Elements()); ; {
						_compr_27, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _836_v88 _dafny.Int
						_836_v88 = interface{}(_compr_27).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_834_v89, _836_v88) {
							_coll27.Add(Companion_Default___.SafeModuloInt(_836_v88, p1), p1)
						}
					}
					return _coll27.ToMap()
				}())
				var _837_v91 _dafny.Map
				_ = _837_v91
				_837_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				_835_v90 = _dafny.SetOf(_837_v91, _837_v91)
				var _838_v92 _dafny.MultiSet
				_ = _838_v92
				_838_v92 = _dafny.MultiSetOf(p1, p1, p1, p1)
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
				_ = _index140
				var _rhs105 bool = (_838_v92).IsProperSubsetOf(_838_v92)
				_ = _rhs105
				var _rhs106 _dafny.Int = p1
				_ = _rhs106
				var _lhs94 _dafny.Array = _730_v5
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_730_v5), 0))
				_ = _lhs95
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				(_lhs94).ArraySet1(_rhs105, (_lhs95).Int())
				_lhs96.F0 = _rhs106
			}
			var _839_v93 _dafny.Sequence
			_ = _839_v93
			_839_v93 = _dafny.SeqOf(p1)
			var _840_v94 _dafny.Map
			_ = _840_v94
			_840_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_839_v93, _730_v5)
			_840_v94 = (_840_v94).Update(_839_v93, (func() _dafny.Array {
				if !(true) {
					return _730_v5
				}
				return _730_v5
			})())
		}
		r0 = p2
		var _841_v95 _dafny.Map
		_ = _841_v95
		_841_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (Companion_D0_.Create_DC2_(p2)).Dtor_cf2())
		var _pat_let_tv17 = _730_v5
		_ = _pat_let_tv17
		var _pat_let_tv18 = _730_v5
		_ = _pat_let_tv18
		r1 = func(_source11 _dafny.Map) bool {
			var _842___mcc_h3 _dafny.Map = _source11
			_ = _842___mcc_h3
			var _843_cf3 _dafny.Map = _842___mcc_h3
			_ = _843_cf3
			return (_pat_let_tv18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_pat_let_tv17), 0))).Int()).(bool)
		}(_841_v95)
		return r0, r1
	}
}
func (_this *C2) M1(p0 _dafny.Array, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _hi3 _dafny.Int = (p2).Minus(p2)
		_ = _hi3
		for _844_i0 := p2; _844_i0.Cmp(_hi3) < 0; _844_i0 = _844_i0.Plus(_dafny.One) {
			var _845_v0 _dafny.Sequence
			_ = _845_v0
			_845_v0 = _dafny.SeqOf(p2, (_dafny.Zero).Minus(p2))
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))
			_ = _index141
			(p0).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_845_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(142))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_846_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_847_i1 _dafny.Int) _dafny.Int {
					return _846_p2
				}
			})(p2)))), (_index141).Int())
			var _848_v1 _dafny.Sequence
			_ = _848_v1
			_848_v1 = _dafny.SeqOf(p1)
			var _849_v2 _dafny.Sequence
			_ = _849_v2
			_849_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ofoujvwnv")
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))
			_ = _index142
			(p0).ArraySet1(((_848_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_849_v2).Cardinality()), _dafny.IntOfUint32((_848_v1).Cardinality()))).Uint32()).(bool)) && (p1), (_index142).Int())
			(globalState).F6 = p2
			var _850_v3 _dafny.Array
			_ = _850_v3
			var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
			_ = _nw153
			_850_v3 = _nw153
			var _851_v4 _dafny.CodePoint
			_ = _851_v4
			_851_v4 = _dafny.CodePoint('j')
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_850_v3), 0))
			_ = _index143
			(_850_v3).ArraySet1CodePoint(_851_v4, (_index143).Int())
			var _852_v5 _dafny.Set
			_ = _852_v5
			_852_v5 = _dafny.SetOf(p1, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool), p1)
			var _853_v6 _dafny.MultiSet
			_ = _853_v6
			_853_v6 = _dafny.MultiSetOf(_dafny.SetOf(p1), _dafny.SetOf(false, p1))
			var _854_v7 _dafny.Map
			_ = _854_v7
			_854_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-641), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool))
			var _855_v8 _dafny.Sequence
			_ = _855_v8
			_855_v8 = _dafny.SeqOf(_852_v5, _852_v5, _852_v5, _852_v5, _852_v5)
			var _856_v9 _dafny.Map
			_ = _856_v9
			_856_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool), _dafny.MultiSetFromSeq(_855_v8))
			var _857_v10 _dafny.Sequence
			_ = _857_v10
			_857_v10 = _dafny.SeqOf(_853_v6)
			var _858_v11 _dafny.Array
			_ = _858_v11
			var _nwElement0_30 _dafny.MultiSet = (_853_v6).Intersection(_853_v6)
			_ = _nwElement0_30
			var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(20))
			_ = _nw154
			(_nw154).ArraySet1(_nwElement0_30, 0)
			(_nw154).ArraySet1(_853_v6, 1)
			(_nw154).ArraySet1(_853_v6, 2)
			(_nw154).ArraySet1(_853_v6, 3)
			(_nw154).ArraySet1(_853_v6, 4)
			(_nw154).ArraySet1(_dafny.MultiSetOf(_852_v5), 5)
			(_nw154).ArraySet1((_853_v6).Update(_dafny.SetOf(p1, p1), Companion_Default___.Abs((_854_v7).Cardinality())), 6)
			(_nw154).ArraySet1(_853_v6, 7)
			(_nw154).ArraySet1(((func() _dafny.MultiSet {
				if (_856_v9).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
					return (_856_v9).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(_dafny.MultiSet)
				}
				return _853_v6
			})()).Intersection(_853_v6), 8)
			(_nw154).ArraySet1(_853_v6, 9)
			(_nw154).ArraySet1(_853_v6, 10)
			(_nw154).ArraySet1((_857_v10).Select((Companion_Default___.SafeIndex((_852_v5).Cardinality(), _dafny.IntOfUint32((_857_v10).Cardinality()))).Uint32()).(_dafny.MultiSet), 11)
			(_nw154).ArraySet1(_dafny.MultiSetOf(_dafny.SetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool), p1)), 12)
			(_nw154).ArraySet1((_853_v6).Update(_852_v5, Companion_Default___.Abs(p2)), 13)
			(_nw154).ArraySet1((_853_v6).Intersection(_853_v6), 14)
			(_nw154).ArraySet1((_dafny.MultiSetOf(_852_v5, _852_v5)).Intersection(_853_v6), 15)
			(_nw154).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool)), _dafny.SetOf(p1)), _dafny.Companion_Sequence_.Update(_855_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_849_v2).Cardinality()), _dafny.IntOfUint32((_855_v8).Cardinality()))).Uint32(), _852_v5))), 16)
			(_nw154).ArraySet1(_853_v6, 17)
			(_nw154).ArraySet1(_853_v6, 18)
			(_nw154).ArraySet1(_853_v6, 19)
			_858_v11 = _nw154
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_858_v11), 0))
			_ = _index144
			(_858_v11).ArraySet1(_dafny.MultiSetFromSeq(_855_v8), (_index144).Int())
			var _859_v12 D3
			_ = _859_v12
			_859_v12 = Companion_D3_.Create_DC8_((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool), _dafny.IntOfInt64(54), _844_i0)
			var _860_v13 _dafny.Map
			_ = _860_v13
			_860_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _849_v2)
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_850_v3), 0))
			_ = _index145
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))
			_ = _index146
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_858_v11), 0))
			_ = _index147
			var _rhs107 _dafny.CodePoint = (_849_v2).Select((Companion_Default___.SafeIndex(_844_i0, _dafny.IntOfUint32((_849_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs107
			var _rhs108 bool = p1
			_ = _rhs108
			var _rhs109 _dafny.Set = (_852_v5).Difference(_852_v5)
			_ = _rhs109
			var _rhs110 bool = (_859_v12).Dtor_cf8()
			_ = _rhs110
			var _rhs111 _dafny.MultiSet = Companion_Default___.Fm16((func() _dafny.Sequence {
				if p1 {
					return (func() _dafny.Sequence {
						if (_860_v13).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
							return (_860_v13).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(_dafny.Sequence)
						}
						return Companion_Default___.Fm11(globalState)
					})()
				}
				return _849_v2
			})(), globalState)
			_ = _rhs111
			var _lhs97 _dafny.Array = _850_v3
			_ = _lhs97
			var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_850_v3), 0))
			_ = _lhs98
			var _lhs99 _dafny.Array = p0
			_ = _lhs99
			var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))
			_ = _lhs100
			var _lhs101 _dafny.Array = _858_v11
			_ = _lhs101
			var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_858_v11), 0))
			_ = _lhs102
			(_lhs97).ArraySet1CodePoint(_rhs107, (_lhs98).Int())
			r0 = _rhs108
			_852_v5 = _rhs109
			(_lhs99).ArraySet1(_rhs110, (_lhs100).Int())
			(_lhs101).ArraySet1(_rhs111, (_lhs102).Int())
			var _861_v14 _dafny.Map
			_ = _861_v14
			_861_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool), p2)
			var _rhs112 bool = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((p0), 0))).Int()).(bool)
			_ = _rhs112
			var _rhs113 _dafny.Map = (_861_v14).Merge(_861_v14)
			_ = _rhs113
			r0 = _rhs112
			_861_v14 = _rhs113
		}
		if (p1) == (!(!(false))) {
			var _862_v15 _dafny.CodePoint
			_ = _862_v15
			_862_v15 = _dafny.CodePoint('f')
			var _863_v16 *C1
			_ = _863_v16
			var _nw155 *C1 = New_C1_()
			_ = _nw155
			_nw155.Ctor__(p1, p1, _862_v15)
			_863_v16 = _nw155
			r0 = (_this).Fm1(globalState)
			var _864_v17 _dafny.MultiSet
			_ = _864_v17
			_864_v17 = _dafny.MultiSetOf((_863_v16).F26())
			r0 = !((_864_v17).IsSubsetOf(_864_v17))
			var _865_v18 _dafny.Sequence
			_ = _865_v18
			_865_v18 = _dafny.SeqOf(_862_v15, _dafny.CodePoint('p'))
			var _866_v19 D2
			_ = _866_v19
			_866_v19 = Companion_D2_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("be"))
			var _867_v20 _dafny.Sequence
			_ = _867_v20
			_867_v20 = _dafny.SeqOf(p1, !((_863_v16).F25()))
			var _868_v21 _dafny.Map
			_ = _868_v21
			_868_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _867_v20)
			var _869_v22 D0
			_ = _869_v22
			_869_v22 = Companion_D0_.Create_DC1_(_868_v21)
			var _870_v23 _dafny.Map
			_ = _870_v23
			_870_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_865_v18, (_866_v19).Dtor_cf4()), Companion_Default___.Fm2(p2, _869_v22, (_863_v16).F25(), (_863_v16).F25(), globalState))
			_870_v23 = (_870_v23).Merge(_870_v23)
			(globalState).F1 = _dafny.Companion_Sequence_.Concatenate(_865_v18, _865_v18)
		} else {
			if (p1) == ((p2).Cmp(p2) == 0) {
				var _871_v24 _dafny.CodePoint
				_ = _871_v24
				_871_v24 = _dafny.CodePoint('d')
				var _872_v25 *C1
				_ = _872_v25
				var _nw156 *C1 = New_C1_()
				_ = _nw156
				_nw156.Ctor__(p1, p1, _871_v24)
				_872_v25 = _nw156
				var _873_v26 _dafny.Sequence
				_ = _873_v26
				_873_v26 = _dafny.SeqOf((_872_v25).F25())
				var _874_v27 _dafny.MultiSet
				_ = _874_v27
				_874_v27 = _dafny.MultiSetOf((p1) || ((_873_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.IntOfUint32((_873_v26).Cardinality()))).Uint32()).(bool)), (p2).Cmp(p2) == 0)
				_874_v27 = _dafny.MultiSetOf((_872_v25).F25())
				r0 = (_872_v25).F25()
				var _875_v28 _dafny.Map
				_ = _875_v28
				_875_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				_875_v28 = _875_v28
				var _876_v29 _dafny.Map
				_ = _876_v29
				_876_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _877_v30 _dafny.Set
				_ = _877_v30
				_877_v30 = _dafny.SetOf(!((_872_v25).F26()))
				var _878_v31 _dafny.Map
				_ = _878_v31
				_878_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_876_v29).Merge(_876_v29), _877_v30)
				_878_v31 = _878_v31
			} else {
				r0 = (p2).Cmp(p2) < 0
				var _879_v32 _dafny.CodePoint
				_ = _879_v32
				_879_v32 = _dafny.CodePoint('o')
				(globalState).F9 = _879_v32
				var _880_v33 _dafny.MultiSet
				_ = _880_v33
				_880_v33 = _dafny.MultiSetOf(p2, _dafny.IntOfInt64(407), p2)
				(globalState).F0 = ((_dafny.Zero).Minus(p2)).Minus(Companion_Default___.Fm8((func() _dafny.Int {
					if (_880_v33).Contains(p2) {
						return (_880_v33).Multiplicity(p2)
					}
					return _dafny.IntOfInt64(428)
				})(), p2, p1, globalState))
				var _881_v34 T0
				_ = _881_v34
				var _nw157 *C1 = New_C1_()
				_ = _nw157
				_nw157.Ctor__(p1, !(p1), Companion_Default___.Fm13(globalState))
				_881_v34 = _nw157
				_881_v34 = _881_v34
				var _882_v35 _dafny.Array
				_ = _882_v35
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
				_ = _nw158
				_882_v35 = _nw158
				var _883_v36 _dafny.Sequence
				_ = _883_v36
				_883_v36 = _dafny.SeqOf(p0)
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_882_v35), 0))
				_ = _index148
				(_882_v35).ArraySet1((_883_v36).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_883_v36).Cardinality()))).Uint32()).(_dafny.Array), (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_882_v35), 0))
				_ = _index149
				(_882_v35).ArraySet1(p0, (_index149).Int())
			}
			var _884_v37 _dafny.Sequence
			_ = _884_v37
			_884_v37 = _dafny.SeqOf(p1, (_this).Fm1(globalState), p1)
			(globalState).F0 = (Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vtgyalopl"), _dafny.IntOfUint32((_884_v37).Cardinality()))).Cardinality(), p2)).Minus(p2)
			var _885_v38 _dafny.CodePoint
			_ = _885_v38
			_885_v38 = _dafny.CodePoint('l')
			(globalState).F9 = _885_v38
			var _886_v39 _dafny.Sequence
			_ = _886_v39
			_886_v39 = _dafny.UnicodeSeqOfUtf8Bytes("uwp")
			var _887_v40 D0
			_ = _887_v40
			_887_v40 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(p2))
			var _888_v41 _dafny.Set
			_ = _888_v41
			_888_v41 = _dafny.SetOf(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_886_v39).Cardinality())), _887_v40)
			var _889_v43 _dafny.Map
			_ = _889_v43
			_889_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll28 = _dafny.NewBuilder()
				_ = _coll28
				for _iter30 := _dafny.Iterate((_888_v41).Elements()); ; {
					_compr_28, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _890_v42 D0
					_890_v42 = interface{}(_compr_28).(D0)
					if (_888_v41).Contains(_890_v42) {
						_coll28.Add(_890_v42)
					}
				}
				return _coll28.ToSet()
			}(), _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))
			(globalState).F3 = (_dafny.Zero).Minus((_889_v43).Cardinality())
			var _891_v44 _dafny.Array
			_ = _891_v44
			var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw159
			_891_v44 = _nw159
			var _892_v45 *C0
			_ = _892_v45
			var _nw160 *C0 = New_C0_()
			_ = _nw160
			_nw160.Ctor__(p1, _891_v44, _885_v38)
			_892_v45 = _nw160
		}
		var _893_v46 _dafny.Set
		_ = _893_v46
		_893_v46 = _dafny.SetOf(_dafny.MultiSetOf(true, p1))
		var _894_v48 D0
		_ = _894_v48
		_894_v48 = Companion_D0_.Create_DC0_(p2)
		if (Companion_Default___.Fm17((_894_v48).Dtor_cf0(), globalState)).IsProperSubsetOf(func() _dafny.Set {
			var _coll29 = _dafny.NewBuilder()
			_ = _coll29
			for _iter31 := _dafny.Iterate((_893_v46).Elements()); ; {
				_compr_29, _ok31 := _iter31()
				if !_ok31 {
					break
				}
				var _895_v47 _dafny.MultiSet
				_895_v47 = interface{}(_compr_29).(_dafny.MultiSet)
				if (_893_v46).Contains(_895_v47) {
					_coll29.Add(_895_v47)
				}
			}
			return _coll29.ToSet()
		}()) {
			var _896_v49 _dafny.Set
			_ = _896_v49
			_896_v49 = _dafny.SetOf(p1)
			var _897_v50 _dafny.Sequence
			_ = _897_v50
			_897_v50 = _dafny.SeqOf(_896_v49, _dafny.SetOf(p1))
			var _898_v51 _dafny.MultiSet
			_ = _898_v51
			_898_v51 = _dafny.MultiSetOf(p1)
			var _899_v52 _dafny.Sequence
			_ = _899_v52
			_899_v52 = _dafny.SeqOf((((_897_v50).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_897_v50).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()).Plus(p2), p2, ((_898_v51).Intersection(_898_v51)).Cardinality())
			_899_v52 = _899_v52
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((p0), 0))
			_ = _index150
			(p0).ArraySet1(p1, (_index150).Int())
			var _900_v53 _dafny.CodePoint
			_ = _900_v53
			_900_v53 = _dafny.CodePoint('c')
			var _901_v54 T0
			_ = _901_v54
			var _nw161 *C1 = New_C1_()
			_ = _nw161
			_nw161.Ctor__(p1, p1, _900_v53)
			_901_v54 = _nw161
			var _902_v55 *C1
			_ = _902_v55
			var _nw162 *C1 = New_C1_()
			_ = _nw162
			_nw162.Ctor__(p1, p1, _901_v54.F24())
			_902_v55 = _nw162
			var _903_v56 _dafny.Map
			_ = _903_v56
			_903_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v54, _902_v55)
			var _904_v57 _dafny.Array
			_ = _904_v57
			var _nwElement0_31 *C1 = (func() *C1 {
				if (_903_v56).Contains(_901_v54) {
					return (_903_v56).Get(_901_v54).(*C1)
				}
				return _902_v55
			})()
			_ = _nwElement0_31
			var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(6))
			_ = _nw163
			(_nw163).ArraySet1(_nwElement0_31, 0)
			(_nw163).ArraySet1(_902_v55, 1)
			(_nw163).ArraySet1(_902_v55, 2)
			(_nw163).ArraySet1(_902_v55, 3)
			(_nw163).ArraySet1(_902_v55, 4)
			(_nw163).ArraySet1(_902_v55, 5)
			_904_v57 = _nw163
			var _905_v58 _dafny.Map
			_ = _905_v58
			_905_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_904_v57, (_902_v55).F25())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((p0), 0))
			_ = _index151
			var _rhs114 bool = (func() bool {
				if (_905_v58).Contains(_904_v57) {
					return (_905_v58).Get(_904_v57).(bool)
				}
				return (_902_v55).F25()
			})()
			_ = _rhs114
			var _rhs115 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-66), p2)
			_ = _rhs115
			var _lhs103 _dafny.Array = p0
			_ = _lhs103
			var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((p0), 0))
			_ = _lhs104
			var _lhs105 *GlobalState = globalState
			_ = _lhs105
			(_lhs103).ArraySet1(_rhs114, (_lhs104).Int())
			_lhs105.F3 = _rhs115
			var _906_v59 _dafny.Sequence
			_ = _906_v59
			_906_v59 = _dafny.UnicodeSeqOfUtf8Bytes("tucrynwj")
			var _907_v60 _dafny.Sequence
			_ = _907_v60
			_907_v60 = _dafny.SeqOf((_902_v55).F26())
			var _908_v61 D2
			_ = _908_v61
			_908_v61 = Companion_D2_.Create_DC5_(p2)
			var _909_v62 D3
			_ = _909_v62
			_909_v62 = Companion_D3_.Create_DC10_(_901_v54, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("n")).Cardinality()))
			var _910_v63 _dafny.Map
			_ = _910_v63
			_910_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((p0), 0))).Int()).(bool), _dafny.IntOfInt64(348))
			var _911_v64 _dafny.Array
			_ = _911_v64
			var _nwElement0_32 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xyugw"), _906_v59)).Cardinality())
			_ = _nwElement0_32
			var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(13))
			_ = _nw164
			(_nw164).ArraySet1(_nwElement0_32, 0)
			(_nw164).ArraySet1((p2).Times((_896_v49).Cardinality()), 1)
			(_nw164).ArraySet1(p2, 2)
			(_nw164).ArraySet1((_dafny.IntOfUint32((_907_v60).Cardinality())).Times((_908_v61).Dtor_cf5()), 3)
			(_nw164).ArraySet1(Companion_Default___.Fm8(p2, p2, true, globalState), 4)
			(_nw164).ArraySet1((p2).Minus(p2), 5)
			(_nw164).ArraySet1((_909_v62).Dtor_cf13(), 6)
			(_nw164).ArraySet1(p2, 7)
			(_nw164).ArraySet1(p2, 8)
			(_nw164).ArraySet1(_dafny.IntOfInt64(156), 9)
			(_nw164).ArraySet1(p2, 10)
			(_nw164).ArraySet1(p2, 11)
			(_nw164).ArraySet1((func() _dafny.Int {
				if (_910_v63).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
					return (_910_v63).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(_dafny.Int)
				}
				return p2
			})(), 12)
			_911_v64 = _nw164
			_911_v64 = _911_v64
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((p0), 0))
			_ = _index152
			(p0).ArraySet1((_dafny.IntOfUint32((_906_v59).Cardinality())).Cmp((p2).Times(p2)) < 0, (_index152).Int())
			var _912_v65 _dafny.Array
			_ = _912_v65
			var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw165
			_912_v65 = _nw165
			var _913_v66 *C0
			_ = _913_v66
			var _nw166 *C0 = New_C0_()
			_ = _nw166
			_nw166.Ctor__(Companion_Default___.Fm6(globalState), _912_v65, _901_v54.F24())
			_913_v66 = _nw166
		} else {
			var _914_v67 _dafny.CodePoint
			_ = _914_v67
			_914_v67 = _dafny.CodePoint('e')
			var _915_v68 T0
			_ = _915_v68
			var _nw167 *C1 = New_C1_()
			_ = _nw167
			_nw167.Ctor__(true, p1, _914_v67)
			_915_v68 = _nw167
			var _916_v69 D3
			_ = _916_v69
			_916_v69 = Companion_D3_.Create_DC10_(_915_v68, (p2).Plus(p2))
			var _source12 D3 = _916_v69
			_ = _source12
			if _source12.Is_DC8() {
				var _917___mcc_h0 bool = _source12.Get_().(D3_DC8).Cf8
				_ = _917___mcc_h0
				var _918___mcc_h1 _dafny.Int = _source12.Get_().(D3_DC8).Cf9
				_ = _918___mcc_h1
				var _919___mcc_h2 _dafny.Int = _source12.Get_().(D3_DC8).Cf10
				_ = _919___mcc_h2
				var _920_cf10 _dafny.Int = _919___mcc_h2
				_ = _920_cf10
				var _921_cf9 _dafny.Int = _918___mcc_h1
				_ = _921_cf9
				var _922_cf8 bool = _917___mcc_h0
				_ = _922_cf8
				var _923_v70 _dafny.Array
				_ = _923_v70
				var _nwElement0_33 _dafny.Int = p2
				_ = _nwElement0_33
				var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(7))
				_ = _nw168
				(_nw168).ArraySet1(_nwElement0_33, 0)
				(_nw168).ArraySet1(_dafny.IntOfInt64(344), 1)
				(_nw168).ArraySet1(_920_cf10, 2)
				(_nw168).ArraySet1(_921_cf9, 3)
				(_nw168).ArraySet1(p2, 4)
				(_nw168).ArraySet1(_920_cf10, 5)
				(_nw168).ArraySet1(_921_cf9, 6)
				_923_v70 = _nw168
				var _924_v71 _dafny.Map
				_ = _924_v71
				_924_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_922_cf8), _923_v70)
				var _925_v72 _dafny.Map
				_ = _925_v72
				_925_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_920_cf10, _924_v71)
				var _926_v73 _dafny.Sequence
				_ = _926_v73
				_926_v73 = _dafny.UnicodeSeqOfUtf8Bytes("lejg")
				var _927_v74 _dafny.Map
				_ = _927_v74
				_927_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_926_v73, p2)
				var _928_v75 _dafny.Map
				_ = _928_v75
				_928_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_923_v70, (func() _dafny.Map {
					if (_925_v72).Contains((func() _dafny.Int {
						if (_927_v74).Contains(_926_v73) {
							return (_927_v74).Get(_926_v73).(_dafny.Int)
						}
						return _920_cf10
					})()) {
						return (_925_v72).Get((func() _dafny.Int {
							if (_927_v74).Contains(_926_v73) {
								return (_927_v74).Get(_926_v73).(_dafny.Int)
							}
							return _920_cf10
						})()).(_dafny.Map)
					}
					return _924_v71
				})())
				var _929_v76 _dafny.Map
				_ = _929_v76
				_929_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _923_v70))
				var _930_v77 _dafny.Array
				_ = _930_v77
				var _nwElement0_34 _dafny.Map = _924_v71
				_ = _nwElement0_34
				var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(11))
				_ = _nw169
				(_nw169).ArraySet1(_nwElement0_34, 0)
				(_nw169).ArraySet1((_924_v71).Merge(_924_v71), 1)
				(_nw169).ArraySet1(_924_v71, 2)
				(_nw169).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _923_v70), 3)
				(_nw169).ArraySet1(_924_v71, 4)
				(_nw169).ArraySet1((_924_v71).Merge(_924_v71), 5)
				(_nw169).ArraySet1(_924_v71, 6)
				(_nw169).ArraySet1((func() _dafny.Map {
					if (_928_v75).Contains(_923_v70) {
						return (_928_v75).Get(_923_v70).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_922_cf8, _923_v70)
				})(), 7)
				(_nw169).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _923_v70), 8)
				(_nw169).ArraySet1(((func() _dafny.Map {
					if (_929_v76).Contains(_922_cf8) {
						return (_929_v76).Get(_922_cf8).(_dafny.Map)
					}
					return _924_v71
				})()).Merge(_924_v71), 9)
				(_nw169).ArraySet1(_924_v71, 10)
				_930_v77 = _nw169
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_930_v77), 0))
				_ = _index153
				(_930_v77).ArraySet1((_924_v71).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _923_v70)), (_index153).Int())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_930_v77), 0))
				_ = _index154
				(_930_v77).ArraySet1((_924_v71).Merge(_924_v71), (_index154).Int())
				var _931_v78 _dafny.Map
				_ = _931_v78
				_931_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_920_cf10, !(p1))
				var _932_v79 _dafny.Set
				_ = _932_v79
				_932_v79 = _dafny.SetOf(p1, p1, p1, _922_cf8)
				var _933_v80 _dafny.Map
				_ = _933_v80
				_933_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_932_v79, _926_v73)
				_931_v78 = (_931_v78).Update((_dafny.Zero).Minus(p2), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("cxhndncj"), (func() _dafny.Sequence {
					if (_933_v80).Contains(_932_v79) {
						return (_933_v80).Get(_932_v79).(_dafny.Sequence)
					}
					return _926_v73
				})()))
				(globalState).F9 = (func() _dafny.CodePoint {
					if p1 {
						return _915_v68.F24()
					}
					return _915_v68.F24()
				})()
				var _934_v81 _dafny.Array
				_ = _934_v81
				var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw170
				_934_v81 = _nw170
				var _935_v82 *C0
				_ = _935_v82
				var _nw171 *C0 = New_C0_()
				_ = _nw171
				_nw171.Ctor__(p1, _934_v81, _dafny.CodePoint('a'))
				_935_v82 = _nw171
				var _936_v83 _dafny.Map
				_ = _936_v83
				_936_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_935_v82, _921_cf9)
				var _937_v84 D3
				_ = _937_v84
				_937_v84 = Companion_D3_.Create_DC9_(false)
				var _938_v85 _dafny.MultiSet
				_ = _938_v85
				_938_v85 = _dafny.MultiSetOf(_937_v84)
				var _939_v86 _dafny.Map
				_ = _939_v86
				_939_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_936_v83).Merge(_936_v83), _938_v85)
				var _940_v87 _dafny.Sequence
				_ = _940_v87
				_940_v87 = _dafny.SeqOf(Companion_D3_.Create_DC9_(true), Companion_D3_.Create_DC9_(_935_v82.F27), _937_v84, _937_v84)
				_939_v86 = (_939_v86).Update(_936_v83, _dafny.MultiSetFromSeq(_940_v87))
			} else if _source12.Is_DC9() {
				var _941___mcc_h3 bool = _source12.Get_().(D3_DC9).Cf11
				_ = _941___mcc_h3
				var _942_cf11 bool = _941___mcc_h3
				_ = _942_cf11
				(globalState).F0 = p2
				var _943_v88 _dafny.Array
				_ = _943_v88
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw172
				_943_v88 = _nw172
				var _944_v89 *C0
				_ = _944_v89
				var _nw173 *C0 = New_C0_()
				_ = _nw173
				_nw173.Ctor__(p1, _943_v88, _914_v67)
				_944_v89 = _nw173
				var _945_v90 _dafny.Map
				_ = _945_v90
				_945_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _946_v91 _dafny.MultiSet
				_ = _946_v91
				_946_v91 = _dafny.MultiSetOf(true, _942_cf11, _944_v89.F27, _944_v89.F27)
				_945_v90 = (_945_v90).Update(((_946_v91).Intersection(_946_v91)).Cardinality(), Companion_Default___.Fm8(p2, p2, _944_v89.F27, globalState))
				var _947_v92 *C0
				_ = _947_v92
				var _nw174 *C0 = New_C0_()
				_ = _nw174
				_nw174.Ctor__(_942_cf11, (_944_v89).F28(), _dafny.CodePoint('k'))
				_947_v92 = _nw174
			} else if _source12.Is_DC10() {
				var _948___mcc_h4 T0 = _source12.Get_().(D3_DC10).Cf12
				_ = _948___mcc_h4
				var _949___mcc_h5 _dafny.Int = _source12.Get_().(D3_DC10).Cf13
				_ = _949___mcc_h5
				var _950_cf13 _dafny.Int = _949___mcc_h5
				_ = _950_cf13
				var _951_cf12 T0 = _948___mcc_h4
				_ = _951_cf12
				var _952_v93 _dafny.Map
				_ = _952_v93
				_952_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _915_v68.F24())
				_952_v93 = (_952_v93).Update(!(!(p1) || (true)), _951_cf12.F24())
				(globalState).F6 = _950_cf13
				var _953_v94 _dafny.Array
				_ = _953_v94
				var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw175
				_953_v94 = _nw175
				var _954_v95 _dafny.Sequence
				_ = _954_v95
				_954_v95 = _dafny.SeqOf(_953_v94)
				var _955_v96 _dafny.Array
				_ = _955_v96
				var _nwElement0_35 _dafny.Array = _953_v94
				_ = _nwElement0_35
				var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.One)
				_ = _nw176
				(_nw176).ArraySet1(_nwElement0_35, 0)
				_955_v96 = _nw176
				var _956_v97 _dafny.Map
				_ = _956_v97
				_956_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_954_v95, _955_v96)
				_956_v97 = (_956_v97).Update(_dafny.Companion_Sequence_.Concatenate(_954_v95, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_953_v94, _953_v94, _953_v94, _953_v94, _953_v94), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf(_953_v94, _953_v94, _953_v94, _953_v94, _953_v94)).Cardinality()))).Uint32(), _953_v94)), _955_v96)
				var _957_v98 _dafny.Array
				_ = _957_v98
				var _nwElement0_36 T0 = _951_cf12
				_ = _nwElement0_36
				var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(16))
				_ = _nw177
				(_nw177).ArraySet1(_nwElement0_36, 0)
				(_nw177).ArraySet1(_951_cf12, 1)
				(_nw177).ArraySet1(_915_v68, 2)
				(_nw177).ArraySet1(_915_v68, 3)
				(_nw177).ArraySet1(_951_cf12, 4)
				(_nw177).ArraySet1(_915_v68, 5)
				(_nw177).ArraySet1(_951_cf12, 6)
				(_nw177).ArraySet1(_915_v68, 7)
				(_nw177).ArraySet1(_915_v68, 8)
				(_nw177).ArraySet1(_915_v68, 9)
				(_nw177).ArraySet1(_915_v68, 10)
				(_nw177).ArraySet1(_915_v68, 11)
				(_nw177).ArraySet1(_951_cf12, 12)
				(_nw177).ArraySet1(_951_cf12, 13)
				(_nw177).ArraySet1(_951_cf12, 14)
				(_nw177).ArraySet1(_915_v68, 15)
				_957_v98 = _nw177
				var _958_v99 D4
				_ = _958_v99
				_958_v99 = Companion_D4_.Create_DC11_(_957_v98)
				var _959_v100 _dafny.Array
				_ = _959_v100
				var _nwElement0_37 _dafny.Array = _957_v98
				_ = _nwElement0_37
				var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(29))
				_ = _nw178
				(_nw178).ArraySet1(_nwElement0_37, 0)
				(_nw178).ArraySet1(_957_v98, 1)
				(_nw178).ArraySet1(_957_v98, 2)
				(_nw178).ArraySet1(_957_v98, 3)
				(_nw178).ArraySet1(_957_v98, 4)
				(_nw178).ArraySet1(_957_v98, 5)
				(_nw178).ArraySet1(_957_v98, 6)
				(_nw178).ArraySet1(_957_v98, 7)
				(_nw178).ArraySet1(_957_v98, 8)
				(_nw178).ArraySet1(_957_v98, 9)
				(_nw178).ArraySet1(_957_v98, 10)
				(_nw178).ArraySet1(_957_v98, 11)
				(_nw178).ArraySet1(_957_v98, 12)
				(_nw178).ArraySet1(_957_v98, 13)
				(_nw178).ArraySet1(_957_v98, 14)
				(_nw178).ArraySet1(_957_v98, 15)
				(_nw178).ArraySet1(_957_v98, 16)
				(_nw178).ArraySet1(_957_v98, 17)
				(_nw178).ArraySet1(_957_v98, 18)
				(_nw178).ArraySet1(_957_v98, 19)
				(_nw178).ArraySet1(_957_v98, 20)
				(_nw178).ArraySet1(_957_v98, 21)
				(_nw178).ArraySet1(_957_v98, 22)
				(_nw178).ArraySet1((func() _dafny.Array {
					if p1 {
						return _957_v98
					}
					return _957_v98
				})(), 23)
				(_nw178).ArraySet1(_957_v98, 24)
				(_nw178).ArraySet1(_957_v98, 25)
				(_nw178).ArraySet1((_958_v99).Dtor_cf14(), 26)
				(_nw178).ArraySet1(_957_v98, 27)
				(_nw178).ArraySet1(_957_v98, 28)
				_959_v100 = _nw178
				var _nwElement0_38 _dafny.Array = _957_v98
				_ = _nwElement0_38
				var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(10))
				_ = _nw179
				(_nw179).ArraySet1(_nwElement0_38, 0)
				(_nw179).ArraySet1(_957_v98, 1)
				(_nw179).ArraySet1((_958_v99).Dtor_cf14(), 2)
				(_nw179).ArraySet1(_957_v98, 3)
				(_nw179).ArraySet1(_957_v98, 4)
				(_nw179).ArraySet1(_957_v98, 5)
				(_nw179).ArraySet1((func() _dafny.Array {
					if p1 {
						return _957_v98
					}
					return _957_v98
				})(), 6)
				(_nw179).ArraySet1(_957_v98, 7)
				(_nw179).ArraySet1(_957_v98, 8)
				(_nw179).ArraySet1(_957_v98, 9)
				_959_v100 = _nw179
			} else {
				var _960___mcc_h6 _dafny.Array = _source12.Get_().(D3_DC7).Cf7
				_ = _960___mcc_h6
				var _961_cf7 _dafny.Array = _960___mcc_h6
				_ = _961_cf7
				var _962_v101 _dafny.Array
				_ = _962_v101
				var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw180
				_962_v101 = _nw180
				_962_v101 = _962_v101
				r0 = p1
				r0 = Companion_Default___.Fm6(globalState)
				var _963_v102 _dafny.Array
				_ = _963_v102
				var _nw181 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
				_ = _nw181
				_963_v102 = _nw181
				var _964_v103 *C1
				_ = _964_v103
				var _nw182 *C1 = New_C1_()
				_ = _nw182
				_nw182.Ctor__(p1, !(p1), _915_v68.F24())
				_964_v103 = _nw182
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_963_v102), 0))
				_ = _index155
				(_963_v102).ArraySet1(_964_v103, (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_963_v102), 0))
				_ = _index156
				(_963_v102).ArraySet1(_964_v103, (_index156).Int())
			}
			var _965_v104 *C1
			_ = _965_v104
			var _nw183 *C1 = New_C1_()
			_ = _nw183
			_nw183.Ctor__(false, p1, _dafny.CodePoint('b'))
			_965_v104 = _nw183
			var _966_v105 _dafny.Sequence
			_ = _966_v105
			_966_v105 = _dafny.UnicodeSeqOfUtf8Bytes("bbeo")
			(globalState).F1 = _966_v105
			var _967_v106 *C1
			_ = _967_v106
			var _nw184 *C1 = New_C1_()
			_ = _nw184
			_nw184.Ctor__((_this).Fm1(globalState), p1, _915_v68.F24())
			_967_v106 = _nw184
			var _968_v107 _dafny.Sequence
			_ = _968_v107
			_968_v107 = _dafny.SeqOf(!((_967_v106).F25()) || ((_967_v106).F26()), (_965_v104).F25())
			_968_v107 = _dafny.Companion_Sequence_.Concatenate(_968_v107, _968_v107)
		}
		var _969_v108 _dafny.CodePoint
		_ = _969_v108
		_969_v108 = _dafny.CodePoint('u')
		var _970_v109 _dafny.Map
		_ = _970_v109
		_970_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _969_v108)
		var _971_v110 _dafny.Map
		_ = _971_v110
		_971_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _969_v108)
		var _972_v111 _dafny.Sequence
		_ = _972_v111
		_972_v111 = _dafny.UnicodeSeqOfUtf8Bytes("trbu")
		var _973_v112 _dafny.Map
		_ = _973_v112
		_973_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-81))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}((func(_974_v108 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_975_i2 _dafny.Int) _dafny.CodePoint {
				return _974_v108
			}
		})(_969_v108)))).Cardinality()), (func() _dafny.CodePoint {
			if (_970_v109).Contains(p1) {
				return (_970_v109).Get(p1).(_dafny.CodePoint)
			}
			return (func() _dafny.CodePoint {
				if (_971_v110).Contains(_dafny.IntOfUint32((_972_v111).Cardinality())) {
					return (_971_v110).Get(_dafny.IntOfUint32((_972_v111).Cardinality())).(_dafny.CodePoint)
				}
				return _969_v108
			})()
		})())
		var _976_v113 _dafny.Array
		_ = _976_v113
		var _nwElement0_39 _dafny.CodePoint = _969_v108
		_ = _nwElement0_39
		var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(19))
		_ = _nw185
		(_nw185).ArraySet1CodePoint(_nwElement0_39, 0)
		(_nw185).ArraySet1CodePoint(_dafny.CodePoint('h'), 1)
		(_nw185).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_973_v112).Contains(_dafny.IntOfInt64(90)) {
				return (_973_v112).Get(_dafny.IntOfInt64(90)).(_dafny.CodePoint)
			}
			return _969_v108
		})(), 2)
		(_nw185).ArraySet1CodePoint(_969_v108, 3)
		(_nw185).ArraySet1CodePoint(_dafny.CodePoint('b'), 4)
		(_nw185).ArraySet1CodePoint(_969_v108, 5)
		(_nw185).ArraySet1CodePoint(_969_v108, 6)
		(_nw185).ArraySet1CodePoint(_969_v108, 7)
		(_nw185).ArraySet1CodePoint(Companion_Default___.Fm13(globalState), 8)
		(_nw185).ArraySet1CodePoint(_969_v108, 9)
		(_nw185).ArraySet1CodePoint(_969_v108, 10)
		(_nw185).ArraySet1CodePoint(_969_v108, 11)
		(_nw185).ArraySet1CodePoint(_dafny.CodePoint('m'), 12)
		(_nw185).ArraySet1CodePoint(_969_v108, 13)
		(_nw185).ArraySet1CodePoint(_969_v108, 14)
		(_nw185).ArraySet1CodePoint(_969_v108, 15)
		(_nw185).ArraySet1CodePoint(_dafny.CodePoint('a'), 16)
		(_nw185).ArraySet1CodePoint(_969_v108, 17)
		(_nw185).ArraySet1CodePoint(_969_v108, 18)
		_976_v113 = _nw185
		var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_976_v113), 0))
		_ = _index157
		(_976_v113).ArraySet1CodePoint(_969_v108, (_index157).Int())
		var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_976_v113), 0))
		_ = _index158
		(_976_v113).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_973_v112).Contains(_dafny.IntOfUint32((_972_v111).Cardinality())) {
				return (_973_v112).Get(_dafny.IntOfUint32((_972_v111).Cardinality())).(_dafny.CodePoint)
			}
			return _969_v108
		})(), (_index158).Int())
		var _977_v114 *C1
		_ = _977_v114
		var _nw186 *C1 = New_C1_()
		_ = _nw186
		_nw186.Ctor__(p1, p1, _969_v108)
		_977_v114 = _nw186
		var _978_v115 _dafny.Map
		_ = _978_v115
		_978_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(530), p2), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_972_v111).Cardinality()), _dafny.IntOfInt64(528)))
		_978_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(58), p2)
		r0 = true
		return r0
	}
}

// End of class C2
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
