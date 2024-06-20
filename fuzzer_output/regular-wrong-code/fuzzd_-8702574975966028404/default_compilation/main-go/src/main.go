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
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) bool {
	return ((_dafny.SetOf(_dafny.CodePoint('h'))).Intersection(_dafny.SetOf(_dafny.CodePoint('n'), _dafny.CodePoint('i')))).IsSubsetOf((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('l'))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.CodePoint
			_0_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('l')), _0_v0) {
				_coll0.Add(_0_v0)
			}
		}
		return _coll0.ToSet()
	}()).Difference(_dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('i'))))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.SetOf(Companion_D1_.Create_DC4_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(128), _dafny.IntOfInt64(6))).Cardinality()), Companion_D1_.Create_DC4_(false, _dafny.IntOfInt64(14)), Companion_D1_.Create_DC4_(false, _dafny.IntOfInt64(-690)), Companion_D1_.Create_DC4_(true, (_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true)))).Cardinality()), Companion_D1_.Create_DC4_(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(321))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality())))).IsSubsetOf(_dafny.SetOf(Companion_D1_.Create_DC4_(true, _dafny.IntOfInt64(997)), Companion_D1_.Create_DC4_(false, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfInt64(195))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf(false))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v0 _dafny.Set
			_2_v0 = interface{}(_compr_1).(_dafny.Set)
			if (_dafny.MultiSetOf(_dafny.SetOf(false))).Contains(_2_v0) {
				_coll1.Add(_2_v0, true)
			}
		}
		return _coll1.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, true, false), false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, true, true), false))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.CodePoint, p2 D1, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("wmqadti")
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if !(true) || (false) {
		return _dafny.CodePoint('p')
	} else if !(!(!(false))) {
		return _dafny.CodePoint('l')
	} else {
		return _dafny.CodePoint('b')
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	var _source0 D4 = Companion_D4_.Create_DC9_(_dafny.SetOf(_dafny.SeqOf(false)))
	_ = _source0
	if _source0.Is_DC10() {
		var _3___mcc_h0 D3 = _source0.Get_().(D4_DC10).Cf16
		_ = _3___mcc_h0
		var _4___mcc_h1 _dafny.Int = _source0.Get_().(D4_DC10).Cf17
		_ = _4___mcc_h1
		var _5___mcc_h2 _dafny.CodePoint = _source0.Get_().(D4_DC10).Cf18
		_ = _5___mcc_h2
		var _6_cf18 _dafny.CodePoint = _5___mcc_h2
		_ = _6_cf18
		var _7_cf17 _dafny.Int = _4___mcc_h1
		_ = _7_cf17
		var _8_cf16 D3 = _3___mcc_h0
		_ = _8_cf16
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(true))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_cf17, !(!(true)))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _7_cf17))
	} else if _source0.Is_DC11() {
		var _9___mcc_h3 _dafny.Int = _source0.Get_().(D4_DC11).Cf19
		_ = _9___mcc_h3
		var _10___mcc_h4 bool = _source0.Get_().(D4_DC11).Cf20
		_ = _10___mcc_h4
		var _11_cf20 bool = _10___mcc_h4
		_ = _11_cf20
		var _12_cf19 _dafny.Int = _9___mcc_h3
		_ = _12_cf19
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_cf20, _12_cf19)
	} else {
		var _13___mcc_h5 _dafny.Set = _source0.Get_().(D4_DC9).Cf15
		_ = _13___mcc_h5
		var _14_cf15 _dafny.Set = _13___mcc_h5
		_ = _14_cf15
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality())).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(49))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-950)))).Cardinality(), (_dafny.IntOfInt64(132)).Times((_dafny.MultiSetOf(true, true)).Cardinality()), (_dafny.IntOfInt64(811)).Minus((_dafny.MultiSetOf(true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_15_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(true)
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_16_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(true)
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(512))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_17_i2 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(!(!(true)))
	})), _dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(true, !(true)))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(761))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ttdvaiasm")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, false, false)).Cardinality()), _dafny.IntOfInt64(-830), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(817)))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(75)))), (Companion_D6_.Create_DC13_(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(458), (_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality()))), _dafny.MultiSetOf(_dafny.IntOfInt64(898)), _dafny.MultiSetOf(_dafny.IntOfInt64(-691), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(560), (_dafny.SetOf(false)).Cardinality())))).Dtor_cf22()))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(((_dafny.MultiSetOf(false)).Cardinality()).Times(_dafny.IntOfInt64(27)), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(171), _dafny.IntOfInt64(389)))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	r0 = p0
	var _18_v0 bool
	_ = _18_v0
	_18_v0 = true
	_18_v0 = _18_v0
	var _19_v1 *C0
	_ = _19_v1
	var _nw0 *C0 = New_C0_()
	_ = _nw0
	_nw0.Ctor__()
	_19_v1 = _nw0
	_19_v1 = _19_v1
	var _20_v2 _dafny.Array
	_ = _20_v2
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
	_ = _nw1
	_20_v2 = _nw1
	var _21_v3 _dafny.Sequence
	_ = _21_v3
	_21_v3 = _dafny.UnicodeSeqOfUtf8Bytes("eo")
	var _rhs0 _dafny.Int = p0
	_ = _rhs0
	var _rhs1 _dafny.Array = _20_v2
	_ = _rhs1
	var _rhs2 _dafny.Int = p0
	_ = _rhs2
	var _rhs3 _dafny.Int = (func() _dafny.Int {
		if _18_v0 {
			return p0
		}
		return _dafny.IntOfUint32((_21_v3).Cardinality())
	})()
	_ = _rhs3
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 *GlobalState = globalState
	_ = _lhs1
	_lhs0.F4 = _rhs0
	_20_v2 = _rhs1
	_lhs1.F4 = _rhs2
	r0 = _rhs3
	if _dafny.Companion_Sequence_.IsProperPrefixOf(_21_v3, _21_v3) {
		_19_v1 = _19_v1
		(globalState).F4 = p0
		var _22_v4 _dafny.Array
		_ = _22_v4
		var _nw2 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw2
		_22_v4 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_22_v4), 0))
		_ = _index0
		(_22_v4).ArraySet1(_18_v0, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_22_v4), 0))
		_ = _index1
		(_22_v4).ArraySet1((func() bool {
			if _18_v0 {
				return _18_v0
			}
			return Companion_Default___.Fm2(globalState)
		})(), (_index1).Int())
		var _23_v5 _dafny.MultiSet
		_ = _23_v5
		_23_v5 = _dafny.MultiSetOf(p0, p0)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_22_v4), 0))
		_ = _index2
		var _rhs4 bool = (Companion_Default___.SafeDivisionInt(p0, p0)).Cmp(p0) != 0
		_ = _rhs4
		var _rhs5 bool = !(!((_23_v5).IsSubsetOf(_23_v5)))
		_ = _rhs5
		var _rhs6 _dafny.Int = p0
		_ = _rhs6
		var _rhs7 bool = ((_dafny.MultiSetOf((_22_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_22_v4), 0))).Int()).(bool), (_22_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_22_v4), 0))).Int()).(bool))).Cardinality()).Cmp(p0) >= 0
		_ = _rhs7
		var _lhs2 _dafny.Array = _22_v4
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_22_v4), 0))
		_ = _lhs3
		var _lhs4 *GlobalState = globalState
		_ = _lhs4
		(_lhs2).ArraySet1(_rhs4, (_lhs3).Int())
		_18_v0 = _rhs5
		_lhs4.F4 = _rhs6
		_18_v0 = _rhs7
		var _24_v6 _dafny.Array
		_ = _24_v6
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_0
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.CodePoint = func(_25_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw3 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw3).ArraySet1CodePoint(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw3).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_24_v6 = _nw3
		var _26_v7 _dafny.CodePoint
		_ = _26_v7
		_26_v7 = _dafny.CodePoint('h')
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_24_v6), 0))
		_ = _index3
		(_24_v6).ArraySet1CodePoint(_26_v7, (_index3).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_24_v6), 0))
		_ = _index4
		(_24_v6).ArraySet1CodePoint(_26_v7, (_index4).Int())
	} else {
		var _27_v8 _dafny.Set
		_ = _27_v8
		_27_v8 = _dafny.SetOf(_18_v0)
		var _28_v9 _dafny.Map
		_ = _28_v9
		_28_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _27_v8)
		(globalState).F4 = (_dafny.Zero).Minus(((_dafny.Zero).Minus(((_28_v9).Cardinality()).Minus(p0))).Minus((p0).Minus(p0)))
		var _29_v10 _dafny.Map
		_ = _29_v10
		_29_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_19_v1).Fm1(false, _dafny.IntOfInt64(605), p0, globalState))
		var _30_v11 _dafny.Sequence
		_ = _30_v11
		_30_v11 = _dafny.SeqOf((_29_v10).Cardinality(), p0, p0)
		var _31_v12 _dafny.Sequence
		_ = _31_v12
		_31_v12 = _dafny.SeqOf(_dafny.IntOfUint32((_30_v11).Cardinality()), p0, p0, p0, p0)
		var _32_v13 _dafny.MultiSet
		_ = _32_v13
		_32_v13 = _dafny.MultiSetOf((Companion_Default___.Fm13((_30_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_30_v11).Cardinality()))).Uint32()).(_dafny.Int), _18_v0, _18_v0, _18_v0, globalState)).Cardinality())
		var _33_v14 _dafny.MultiSet
		_ = _33_v14
		_33_v14 = _dafny.MultiSetOf(_18_v0, _18_v0, true)
		var _34_v15 _dafny.Map
		_ = _34_v15
		_34_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _18_v0)
		r0 = (func() _dafny.Int {
			if (_32_v13).Contains(Companion_Default___.SafeModuloInt((_33_v14).Cardinality(), p0)) {
				return (_32_v13).Multiplicity(Companion_Default___.SafeModuloInt((_33_v14).Cardinality(), p0))
			}
			return (func() _dafny.Int {
				if (_29_v10).Contains(p0) {
					return (_29_v10).Get(p0).(_dafny.Int)
				}
				return (_34_v15).Cardinality()
			})()
		})()
		var _35_v16 *C0
		_ = _35_v16
		var _nw4 *C0 = New_C0_()
		_ = _nw4
		_nw4.Ctor__()
		_35_v16 = _nw4
		var _36_v17 _dafny.Array
		_ = _36_v17
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw5
		_36_v17 = _nw5
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_36_v17), 0))
		_ = _index5
		(_36_v17).ArraySet1(false, (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_36_v17), 0))
		_ = _index6
		(_36_v17).ArraySet1(_18_v0, (_index6).Int())
		(globalState).F4 = p0
	}
	var _hi0 _dafny.Int = (p0).Times(p0)
	_ = _hi0
	for _37_i1 := p0; _37_i1.Cmp(_hi0) < 0; _37_i1 = _37_i1.Plus(_dafny.One) {
		var _38_v18 D1
		_ = _38_v18
		_38_v18 = Companion_D1_.Create_DC4_(_18_v0, _37_i1)
		if (func() bool {
			if (_38_v18).Dtor_cf8() {
				return _18_v0
			}
			return _18_v0
		})() {
			_21_v3 = _21_v3
			var _39_v19 _dafny.Array
			_ = _39_v19
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw6
			_39_v19 = _nw6
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_39_v19), 0))
			_ = _index7
			(_39_v19).ArraySet1(_18_v0, (_index7).Int())
			var _40_v20 _dafny.Map
			_ = _40_v20
			_40_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_20_v2, _21_v3)
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_39_v19), 0))
			_ = _index8
			(_39_v19).ArraySet1(!(_40_v20).Contains(_20_v2), (_index8).Int())
			var _41_v21 _dafny.Array
			_ = _41_v21
			var _nwElement0_0 _dafny.Array = _20_v2
			_ = _nwElement0_0
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(10))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_0, 0)
			(_nw7).ArraySet1(_20_v2, 1)
			(_nw7).ArraySet1(_20_v2, 2)
			(_nw7).ArraySet1(_20_v2, 3)
			(_nw7).ArraySet1(_20_v2, 4)
			(_nw7).ArraySet1(_20_v2, 5)
			(_nw7).ArraySet1(_20_v2, 6)
			(_nw7).ArraySet1(_20_v2, 7)
			(_nw7).ArraySet1(_20_v2, 8)
			(_nw7).ArraySet1(_20_v2, 9)
			_41_v21 = _nw7
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_41_v21), 0))
			_ = _index9
			(_41_v21).ArraySet1(_20_v2, (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_41_v21), 0))
			_ = _index10
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_39_v19), 0))
			_ = _index11
			var _rhs8 _dafny.Array = _20_v2
			_ = _rhs8
			var _rhs9 bool = _18_v0
			_ = _rhs9
			var _lhs5 _dafny.Array = _41_v21
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_41_v21), 0))
			_ = _lhs6
			var _lhs7 _dafny.Array = _39_v19
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_39_v19), 0))
			_ = _lhs8
			(_lhs5).ArraySet1(_rhs8, (_lhs6).Int())
			(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
			var _42_v22 _dafny.Map
			_ = _42_v22
			_42_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _37_i1)
			var _43_v23 _dafny.Map
			_ = _43_v23
			_43_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v1, (_42_v22).Merge(_42_v22))
			var _44_v24 _dafny.Map
			_ = _44_v24
			_44_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm2(globalState)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v1, _42_v22))
			_43_v23 = (_43_v23).Merge((func() _dafny.Map {
				if (_44_v24).Contains(_18_v0) {
					return (_44_v24).Get(_18_v0).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v1, _42_v22)
			})())
			var _45_v25 _dafny.Map
			_ = _45_v25
			_45_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, _19_v1)
			var _46_v26 _dafny.CodePoint
			_ = _46_v26
			_46_v26 = _dafny.CodePoint('p')
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_39_v19), 0))
			_ = _index12
			(_39_v19).ArraySet1((Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(815), (_45_v25).Cardinality())).Cmp((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_21_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.IntOfUint32((_21_v3).Cardinality()))).Uint32(), _46_v26)).Cardinality())).Times(p0)) <= 0, (_index12).Int())
		} else {
			var _47_v27 _dafny.Array
			_ = _47_v27
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_1
			var _nw8 _dafny.Array
			_ = _nw8
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw8 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_48_v0 bool, _49_p0 _dafny.Int) func(_dafny.Int) bool {
					return func(_50_i2 _dafny.Int) bool {
						return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_48_v0))).Cardinality()).Cmp(_49_p0) < 0
					}
				})(_18_v0, p0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw8 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw8).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw8).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_47_v27 = _nw8
			var _51_v28 _dafny.Array
			_ = _51_v28
			_51_v28 = _47_v27
			_47_v27 = (_51_v28)
			(globalState).F4 = _dafny.IntOfInt64(617)
			(globalState).F4 = (_dafny.Zero).Minus(p0)
			var _52_v29 _dafny.CodePoint
			_ = _52_v29
			_52_v29 = _dafny.CodePoint('q')
			_52_v29 = _52_v29
			var _53_v30 _dafny.Map
			_ = _53_v30
			_53_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_v3, Companion_D0_.Create_DC0_(_20_v2, p0))
			var _54_v31 _dafny.MultiSet
			_ = _54_v31
			_54_v31 = _dafny.MultiSetOf(_18_v0, _18_v0)
			var _55_v32 _dafny.MultiSet
			_ = _55_v32
			_55_v32 = _dafny.MultiSetOf((_19_v1).Fm1(true, p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_53_v30).Cardinality())).Cardinality(), globalState), p0, p0, (_54_v31).Cardinality(), p0)
			var _56_v33 _dafny.Array
			_ = _56_v33
			var _nwElement0_1 _dafny.MultiSet = _55_v32
			_ = _nwElement0_1
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(3))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_1, 0)
			(_nw9).ArraySet1(_55_v32, 1)
			(_nw9).ArraySet1(_55_v32, 2)
			_56_v33 = _nw9
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_2
			var _nw10 _dafny.Array
			_ = _nw10
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw10 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.MultiSet = func(_57_i3 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_dafny.IntOfInt64(418))
				}
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw10 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw10).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw10).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_56_v33 = _nw10
		}
		var _58_v34 _dafny.Map
		_ = _58_v34
		_58_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _59_v35 _dafny.Sequence
		_ = _59_v35
		_59_v35 = _dafny.SeqOf((_58_v34).Merge(_58_v34), _58_v34, _58_v34, _58_v34, _58_v34)
		var _rhs10 *C0 = _19_v1
		_ = _rhs10
		var _rhs11 _dafny.Map = (_59_v35).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_59_v35).Cardinality()))).Uint32()).(_dafny.Map)
		_ = _rhs11
		var _rhs12 _dafny.Int = (_dafny.IntOfInt64(618)).Plus(_37_i1)
		_ = _rhs12
		var _lhs9 *GlobalState = globalState
		_ = _lhs9
		_19_v1 = _rhs10
		_58_v34 = _rhs11
		_lhs9.F4 = _rhs12
		var _60_v36 *C0
		_ = _60_v36
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__()
		_60_v36 = _nw11
		var _61_v37 *C0
		_ = _61_v37
		var _nw12 *C0 = New_C0_()
		_ = _nw12
		_nw12.Ctor__()
		_61_v37 = _nw12
	}
	r0 = (func() _dafny.Int {
		if (false) && (_18_v0) {
			return _dafny.IntOfUint32((_21_v3).Cardinality())
		}
		return Companion_Default___.SafeModuloInt(p0, p0)
	})()
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _62_v0 _dafny.Array
	_ = _62_v0
	var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
	_ = _nw13
	_62_v0 = _nw13
	var _63_globalState *GlobalState
	_ = _63_globalState
	var _nw14 *GlobalState = New_GlobalState_()
	_ = _nw14
	_nw14.Ctor__(true, _dafny.IntOfInt64(-521), _dafny.IntOfInt64(537), _62_v0, _dafny.IntOfInt64(681), _dafny.CodePoint('e'), false)
	_63_globalState = _nw14
	var _64_v1 _dafny.Int
	_ = _64_v1
	_64_v1 = _dafny.IntOfInt64(176)
	var _65_v2 bool
	_ = _65_v2
	_65_v2 = true
	var _hi1 _dafny.Int = (func() _dafny.Int {
		if _65_v2 {
			return _64_v1
		}
		return _dafny.IntOfInt64(-845)
	})()
	_ = _hi1
	for _66_i0 := _64_v1; _66_i0.Cmp(_hi1) < 0; _66_i0 = _66_i0.Plus(_dafny.One) {
		(_63_globalState).F4 = _66_i0
		var _67_v3 _dafny.Sequence
		_ = _67_v3
		_67_v3 = _dafny.UnicodeSeqOfUtf8Bytes("mjypph")
		_67_v3 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(163))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_68_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})), _67_v3)
		_64_v1 = _dafny.IntOfInt64(999)
		var _69_v4 *C0
		_ = _69_v4
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__()
		_69_v4 = _nw15
	}
	_65_v2 = _65_v2
	var _70_v5 *C0
	_ = _70_v5
	var _nw16 *C0 = New_C0_()
	_ = _nw16
	_nw16.Ctor__()
	_70_v5 = _nw16
	var _71_v6 _dafny.Map
	_ = _71_v6
	_71_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v5, _64_v1)
	var _72_v7 _dafny.Map
	_ = _72_v7
	_72_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_64_v1).Cmp((_71_v6).Cardinality()) <= 0, _65_v2)
	var _73_v8 _dafny.Map
	_ = _73_v8
	_73_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, _dafny.IntOfInt64(188))
	_72_v7 = (_72_v7).Update(_65_v2, (_73_v8).Contains((_73_v8).Cardinality()))
	var _74_v9 _dafny.Sequence
	_ = _74_v9
	_74_v9 = _dafny.SeqOf(_62_v0)
	(_63_globalState).F4 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_74_v9, _dafny.SeqOf(_62_v0, _62_v0, _62_v0))).Cardinality())
	var _75_i2 _dafny.Int
	_ = _75_i2
	_75_i2 = _dafny.Zero
	{
		for (_65_v2) || (Companion_Default___.Fm2(_63_globalState)) {
			{
				if (_75_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_75_i2 = (_75_i2).Plus(_dafny.One)
				var _76_v10 _dafny.Array
				_ = _76_v10
				var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw17
				_76_v10 = _nw17
				var _77_v11 _dafny.Sequence
				_ = _77_v11
				_77_v11 = _dafny.SeqOf(_dafny.IntOfInt64(65))
				var _78_v12 _dafny.CodePoint
				_ = _78_v12
				_78_v12 = _dafny.CodePoint('w')
				var _79_v13 _dafny.Map
				_ = _79_v13
				_79_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v12, _64_v1)
				var _80_v14 _dafny.Sequence
				_ = _80_v14
				_80_v14 = _dafny.SeqOf((_77_v11).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_79_v13).Contains(_78_v12) {
						return (_79_v13).Get(_78_v12).(_dafny.Int)
					}
					return _64_v1
				})(), _dafny.IntOfUint32((_77_v11).Cardinality()))).Uint32()).(_dafny.Int))
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_76_v10), 0))
				_ = _index13
				(_76_v10).ArraySet1((_dafny.Zero).Minus((_80_v14).Select((Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_80_v14).Cardinality()))).Uint32()).(_dafny.Int)), (_index13).Int())
				var _81_v15 _dafny.Set
				_ = _81_v15
				_81_v15 = _dafny.SetOf((_dafny.Zero).Minus(_64_v1), _64_v1, _64_v1)
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_76_v10), 0))
				_ = _index14
				(_76_v10).ArraySet1(((_81_v15).Difference((_81_v15).Intersection(_81_v15))).Cardinality(), (_index14).Int())
				var _82_v16 _dafny.Int
				_ = _82_v16
				var _out0 _dafny.Int
				_ = _out0
				_out0 = Companion_Default___.M0(Companion_Default___.SafeDivisionInt((_76_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_76_v10), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_dafny.IntOfUint32((_80_v14).Cardinality()))), _63_globalState)
				_82_v16 = _out0
				var _83_v18 _dafny.Array
				_ = _83_v18
				var _len0_3 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_3
				var _nw18 _dafny.Array
				_ = _nw18
				if _len0_3.Cmp(_dafny.Zero) == 0 {
					_nw18 = _dafny.NewArray(_len0_3)
				} else {
					var _init3 func(_dafny.Int) _dafny.Map = (func(_84_v7 _dafny.Map, _85_v15 _dafny.Set, _86_v2 bool) func(_dafny.Int) _dafny.Map {
						return func(_87_i3 _dafny.Int) _dafny.Map {
							return (func() _dafny.Map {
								var _coll2 = _dafny.NewMapBuilder()
								_ = _coll2
								for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v7, (_85_v15).Cardinality())).Keys().Elements()); ; {
									_compr_2, _ok2 := _iter2()
									if !_ok2 {
										break
									}
									var _88_v17 _dafny.Map
									_88_v17 = interface{}(_compr_2).(_dafny.Map)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v7, (_85_v15).Cardinality())).Contains(_88_v17) {
										_coll2.Add(_88_v17, _86_v2)
									}
								}
								return _coll2.ToMap()
							}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_84_v7).Update(_86_v2, _86_v2), false))
						}
					})(_72_v7, _81_v15, _65_v2)
					_ = _init3
					var _element0_3 = _init3(_dafny.Zero)
					_ = _element0_3
					_nw18 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
					(_nw18).ArraySet1(_element0_3, 0)
					var _nativeLen0_3 = (_len0_3).Int()
					_ = _nativeLen0_3
					for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
						(_nw18).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
					}
				}
				_83_v18 = _nw18
				var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_83_v18), 0))
				_ = _index15
				(_83_v18).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v7, _65_v2)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v7, _65_v2)), (_index15).Int())
				var _89_v19 _dafny.Array
				_ = _89_v19
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_4
				var _nw19 _dafny.Array
				_ = _nw19
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw19 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.CodePoint = (func(_90_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_91_i4 _dafny.Int) _dafny.CodePoint {
							return _90_v12
						}
					})(_78_v12)
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw19 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw19).ArraySet1CodePoint(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw19).ArraySet1CodePoint(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_89_v19 = _nw19
				var _92_v20 _dafny.Array
				_ = _92_v20
				var _nwElement0_2 _dafny.Array = _89_v19
				_ = _nwElement0_2
				var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(3))
				_ = _nw20
				(_nw20).ArraySet1(_nwElement0_2, 0)
				(_nw20).ArraySet1(_89_v19, 1)
				(_nw20).ArraySet1(_89_v19, 2)
				_92_v20 = _nw20
				var _93_v21 _dafny.Sequence
				_ = _93_v21
				_93_v21 = _dafny.SeqOf(_92_v20)
				var _94_v22 _dafny.Array
				_ = _94_v22
				var _nwElement0_3 _dafny.Array = (_93_v21).Select((Companion_Default___.SafeIndex(_82_v16, _dafny.IntOfUint32((_93_v21).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _nwElement0_3
				var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(6))
				_ = _nw21
				(_nw21).ArraySet1(_nwElement0_3, 0)
				(_nw21).ArraySet1(_92_v20, 1)
				(_nw21).ArraySet1(_92_v20, 2)
				(_nw21).ArraySet1(_92_v20, 3)
				(_nw21).ArraySet1(_92_v20, 4)
				(_nw21).ArraySet1(_92_v20, 5)
				_94_v22 = _nw21
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_94_v22), 0))
				_ = _index16
				(_94_v22).ArraySet1(_92_v20, (_index16).Int())
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_83_v18), 0))
				_ = _index17
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_94_v22), 0))
				_ = _index18
				var _rhs13 *C0 = _70_v5
				_ = _rhs13
				var _rhs14 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v7, _65_v2)
				_ = _rhs14
				var _rhs15 bool = _65_v2
				_ = _rhs15
				var _rhs16 _dafny.Array = _92_v20
				_ = _rhs16
				var _lhs10 _dafny.Array = _83_v18
				_ = _lhs10
				var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_83_v18), 0))
				_ = _lhs11
				var _lhs12 _dafny.Array = _94_v22
				_ = _lhs12
				var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_94_v22), 0))
				_ = _lhs13
				_70_v5 = _rhs13
				(_lhs10).ArraySet1(_rhs14, (_lhs11).Int())
				_65_v2 = _rhs15
				(_lhs12).ArraySet1(_rhs16, (_lhs13).Int())
				_65_v2 = Companion_Default___.Fm2(_63_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_65_v2 = (_65_v2) && (_65_v2)
	var _95_v23 _dafny.Sequence
	_ = _95_v23
	_95_v23 = _dafny.UnicodeSeqOfUtf8Bytes("grpdnh")
	(_63_globalState).F4 = (_64_v1).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_95_v23).Cardinality())))
	var _96_v24 D0
	_ = _96_v24
	_96_v24 = Companion_D0_.Create_DC1_(_65_v2, _64_v1, _64_v1, _64_v1)
	var _97_v25 _dafny.Map
	_ = _97_v25
	_97_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_65_v2)), _96_v24)
	var _98_v26 _dafny.MultiSet
	_ = _98_v26
	_98_v26 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v2, _96_v24)).Update(_65_v2, _96_v24), _97_v25)
	var _99_v27 D0
	_ = _99_v27
	_99_v27 = Companion_D0_.Create_DC1_(_65_v2, _64_v1, _64_v1, (func() _dafny.Int {
		if (_98_v26).Contains(_97_v25) {
			return (_98_v26).Multiplicity(_97_v25)
		}
		return (_70_v5).Fm1(!(_65_v2), _64_v1, _64_v1, _63_globalState)
	})())
	var _source1 D0 = _99_v27
	_ = _source1
	if _source1.Is_DC0() {
		var _100___mcc_h0 _dafny.Array = _source1.Get_().(D0_DC0).Cf0
		_ = _100___mcc_h0
		var _101___mcc_h1 _dafny.Int = _source1.Get_().(D0_DC0).Cf1
		_ = _101___mcc_h1
		var _102_cf1 _dafny.Int = _101___mcc_h1
		_ = _102_cf1
		var _103_cf0 _dafny.Array = _100___mcc_h0
		_ = _103_cf0
		var _104_v28 *C0
		_ = _104_v28
		var _nw22 *C0 = New_C0_()
		_ = _nw22
		_nw22.Ctor__()
		_104_v28 = _nw22
		var _105_v29 _dafny.CodePoint
		_ = _105_v29
		_105_v29 = _dafny.CodePoint('m')
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_103_cf0), 0))
		_ = _index19
		(_103_cf0).ArraySet1(_64_v1, (_index19).Int())
		var _106_v30 _dafny.Map
		_ = _106_v30
		_106_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v29, _65_v2)
		var _107_v31 _dafny.Sequence
		_ = _107_v31
		_107_v31 = _dafny.SeqOf(_dafny.MultiSetOf((_106_v30).Cardinality()))
		var _108_v32 _dafny.MultiSet
		_ = _108_v32
		_108_v32 = _dafny.MultiSetOf(_64_v1)
		var _109_v33 _dafny.Sequence
		_ = _109_v33
		_109_v33 = _dafny.SeqOf(_65_v2)
		var _110_v34 _dafny.Map
		_ = _110_v34
		_110_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, _109_v33)
		var _111_v35 _dafny.Array
		_ = _111_v35
		var _nwElement0_4 _dafny.Sequence = _107_v31
		_ = _nwElement0_4
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(16))
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_4, 0)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v31, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_108_v32, _108_v32), (Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_dafny.SeqOf(_108_v32, _108_v32)).Cardinality()))).Uint32(), _108_v32)), 1)
		(_nw23).ArraySet1(_107_v31, 2)
		(_nw23).ArraySet1(_107_v31, 3)
		(_nw23).ArraySet1(_107_v31, 4)
		(_nw23).ArraySet1(_dafny.SeqOf(_108_v32), 5)
		(_nw23).ArraySet1(_107_v31, 6)
		(_nw23).ArraySet1(_107_v31, 7)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Update(_107_v31, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_110_v34).Contains(_64_v1) {
				return (_110_v34).Get(_64_v1).(_dafny.Sequence)
			}
			return _109_v33
		})()).Cardinality()), _dafny.IntOfUint32((_107_v31).Cardinality()))).Uint32(), _108_v32), 8)
		(_nw23).ArraySet1(_107_v31, 9)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v31, _107_v31), 10)
		(_nw23).ArraySet1(_107_v31, 11)
		(_nw23).ArraySet1(_107_v31, 12)
		(_nw23).ArraySet1(_dafny.SeqOf(_108_v32, (_107_v31).Select((Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_107_v31).Cardinality()))).Uint32()).(_dafny.MultiSet), _108_v32, _108_v32), 13)
		(_nw23).ArraySet1(_107_v31, 14)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v31, _107_v31), 15)
		_111_v35 = _nw23
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_111_v35), 0))
		_ = _index20
		(_111_v35).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_107_v31, _107_v31), (_index20).Int())
		var _112_v36 _dafny.Map
		_ = _112_v36
		_112_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_cf0, (_dafny.Zero).Minus(_64_v1))
		var _113_v37 _dafny.Map
		_ = _113_v37
		_113_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-229))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_114_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_115_i5 _dafny.Int) _dafny.CodePoint {
				return _114_v29
			}
		})(_105_v29))))
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_103_cf0), 0))
		_ = _index21
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_111_v35), 0))
		_ = _index22
		var _rhs17 _dafny.CodePoint = _105_v29
		_ = _rhs17
		var _rhs18 _dafny.Int = (func() _dafny.Int {
			if (_112_v36).Contains(_103_cf0) {
				return (_112_v36).Get(_103_cf0).(_dafny.Int)
			}
			return (_102_cf1).Times(_102_cf1)
		})()
		_ = _rhs18
		var _rhs19 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_107_v31, (Companion_Default___.SafeIndex(_102_cf1, _dafny.IntOfUint32((_107_v31).Cardinality()))).Uint32(), _108_v32), _dafny.Companion_Sequence_.Concatenate(_107_v31, _dafny.SeqOf(_108_v32, _108_v32)))
		_ = _rhs19
		var _rhs20 *C0 = _104_v28
		_ = _rhs20
		var _rhs21 _dafny.Int = Companion_Default___.SafeDivisionInt(((_113_v37).Update(_65_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}((func(_116_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_117_i6 _dafny.Int) _dafny.CodePoint {
				return _116_v29
			}
		})(_105_v29))))).Cardinality(), Companion_Default___.SafeDivisionInt(_64_v1, (_104_v28).Fm1(true, _102_cf1, _102_cf1, _63_globalState)))
		_ = _rhs21
		var _lhs14 _dafny.Array = _103_cf0
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_103_cf0), 0))
		_ = _lhs15
		var _lhs16 _dafny.Array = _111_v35
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_111_v35), 0))
		_ = _lhs17
		var _lhs18 *GlobalState = _63_globalState
		_ = _lhs18
		_105_v29 = _rhs17
		(_lhs14).ArraySet1(_rhs18, (_lhs15).Int())
		(_lhs16).ArraySet1(_rhs19, (_lhs17).Int())
		_104_v28 = _rhs20
		_lhs18.F4 = _rhs21
		_105_v29 = _105_v29
		var _118_v38 _dafny.Set
		_ = _118_v38
		_118_v38 = _dafny.SetOf(_65_v2)
		var _119_v39 _dafny.Map
		_ = _119_v39
		_119_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_108_v32).Intersection(_108_v32), (_118_v38).Cardinality())
		_102_cf1 = (func() _dafny.Int {
			if (_119_v39).Contains(_108_v32) {
				return (_119_v39).Get(_108_v32).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_95_v23).Cardinality())
		})()
	} else if _source1.Is_DC1() {
		var _120___mcc_h2 bool = _source1.Get_().(D0_DC1).Cf2
		_ = _120___mcc_h2
		var _121___mcc_h3 _dafny.Int = _source1.Get_().(D0_DC1).Cf3
		_ = _121___mcc_h3
		var _122___mcc_h4 _dafny.Int = _source1.Get_().(D0_DC1).Cf4
		_ = _122___mcc_h4
		var _123___mcc_h5 _dafny.Int = _source1.Get_().(D0_DC1).Cf5
		_ = _123___mcc_h5
		var _124_cf5 _dafny.Int = _123___mcc_h5
		_ = _124_cf5
		var _125_cf4 _dafny.Int = _122___mcc_h4
		_ = _125_cf4
		var _126_cf3 _dafny.Int = _121___mcc_h3
		_ = _126_cf3
		var _127_cf2 bool = _120___mcc_h2
		_ = _127_cf2
		if _65_v2 {
			_72_v7 = _72_v7
			var _128_v40 _dafny.Sequence
			_ = _128_v40
			_128_v40 = _dafny.SeqOf(_dafny.IntOfInt64(227))
			_64_v1 = (func() _dafny.Int {
				if _65_v2 {
					return _dafny.IntOfInt64(623)
				}
				return ((_128_v40).Select((Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_128_v40).Cardinality()))).Uint32()).(_dafny.Int)).Times((_dafny.Zero).Minus(_125_cf4))
			})()
			var _129_v41 _dafny.Sequence
			_ = _129_v41
			_129_v41 = _dafny.SeqOf(_70_v5, _70_v5)
			var _130_v42 _dafny.Sequence
			_ = _130_v42
			_130_v42 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_129_v41, _dafny.SeqOf(_70_v5)))
			_130_v42 = _130_v42
			_73_v8 = ((_73_v8).Merge(_73_v8)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_cf3, _126_cf3))
			var _131_v43 _dafny.Map
			_ = _131_v43
			_131_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, _dafny.SetOf(_70_v5))
			var _132_v44 D1
			_ = _132_v44
			_132_v44 = Companion_D1_.Create_DC3_(_131_v43)
			_131_v43 = (_132_v44).Dtor_cf7()
		} else {
			_70_v5 = _70_v5
			var _133_v45 _dafny.Array
			_ = _133_v45
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw24
			_133_v45 = _nw24
			var _134_v46 _dafny.Sequence
			_ = _134_v46
			_134_v46 = _dafny.SeqOf(_65_v2)
			var _135_v47 _dafny.Sequence
			_ = _135_v47
			_135_v47 = _dafny.SeqOf(_134_v46, _dafny.SeqOf(_127_cf2), (_70_v5).Fm0(_124_cf5, _dafny.IntOfUint32((_134_v46).Cardinality()), _125_cf4, _65_v2, _63_globalState))
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_133_v45), 0))
			_ = _index23
			(_133_v45).ArraySet1((_135_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_95_v23).Cardinality()), _dafny.IntOfUint32((_135_v47).Cardinality()))).Uint32()).(_dafny.Sequence), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_133_v45), 0))
			_ = _index24
			(_133_v45).ArraySet1(_134_v46, (_index24).Int())
			var _136_v48 _dafny.Array
			_ = _136_v48
			var _nwElement0_5 _dafny.Array = _62_v0
			_ = _nwElement0_5
			var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(10))
			_ = _nw25
			(_nw25).ArraySet1(_nwElement0_5, 0)
			(_nw25).ArraySet1(_62_v0, 1)
			(_nw25).ArraySet1(_62_v0, 2)
			(_nw25).ArraySet1(_62_v0, 3)
			(_nw25).ArraySet1(_62_v0, 4)
			(_nw25).ArraySet1(_62_v0, 5)
			(_nw25).ArraySet1(_62_v0, 6)
			(_nw25).ArraySet1(_62_v0, 7)
			(_nw25).ArraySet1(_62_v0, 8)
			(_nw25).ArraySet1(_62_v0, 9)
			_136_v48 = _nw25
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_136_v48), 0))
			_ = _index25
			(_136_v48).ArraySet1(_62_v0, (_index25).Int())
			var _137_v49 _dafny.Array
			_ = _137_v49
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw26
			_137_v49 = _nw26
			var _138_v50 _dafny.Map
			_ = _138_v50
			_138_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v49, _127_cf2)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_136_v48), 0))
			_ = _index26
			(_136_v48).ArraySet1((func() _dafny.Array {
				if !(_138_v50).Equals(_138_v50) {
					return _62_v0
				}
				return _62_v0
			})(), (_index26).Int())
			var _139_v51 _dafny.Array
			_ = _139_v51
			var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
			_ = _nw27
			_139_v51 = _nw27
			var _140_v52 _dafny.CodePoint
			_ = _140_v52
			_140_v52 = _dafny.CodePoint('u')
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_139_v51), 0))
			_ = _index27
			(_139_v51).ArraySet1CodePoint(_140_v52, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_139_v51), 0))
			_ = _index28
			(_139_v51).ArraySet1CodePoint(_140_v52, (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_137_v49), 0))
			_ = _index29
			(_137_v49).ArraySet1(_124_cf5, (_index29).Int())
			var _141_v53 _dafny.MultiSet
			_ = _141_v53
			_141_v53 = _dafny.MultiSetOf(_127_cf2, _127_cf2)
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_137_v49), 0))
			_ = _index30
			var _rhs22 _dafny.Int = (func() _dafny.Int {
				if (_73_v8).Contains(Companion_Default___.SafeModuloInt(_126_cf3, _125_cf4)) {
					return (_73_v8).Get(Companion_Default___.SafeModuloInt(_126_cf3, _125_cf4)).(_dafny.Int)
				}
				return _64_v1
			})()
			_ = _rhs22
			var _rhs23 *C0 = _70_v5
			_ = _rhs23
			var _rhs24 _dafny.Array = (func() _dafny.Array {
				if _65_v2 {
					return _137_v49
				}
				return _137_v49
			})()
			_ = _rhs24
			var _rhs25 _dafny.Int = (((Companion_Default___.Fm3(_63_globalState)).Intersection(_141_v53)).Union((_141_v53).Update(_65_v2, Companion_Default___.Abs(_125_cf4)))).Cardinality()
			_ = _rhs25
			var _rhs26 _dafny.Array = (func() _dafny.Array {
				if (func() bool {
					if _65_v2 {
						return _127_cf2
					}
					return _65_v2
				})() {
					return _dafny.ArrayCastTo((_136_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_136_v48), 0))).Int()))
				}
				return _dafny.ArrayCastTo((_136_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_136_v48), 0))).Int()))
			})()
			_ = _rhs26
			var _lhs19 _dafny.Array = _137_v49
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_137_v49), 0))
			_ = _lhs20
			_64_v1 = _rhs22
			_70_v5 = _rhs23
			_137_v49 = _rhs24
			(_lhs19).ArraySet1(_rhs25, (_lhs20).Int())
			_62_v0 = _rhs26
		}
		_127_cf2 = _65_v2
		if ((_126_cf3).Plus(_64_v1)).Cmp(_dafny.IntOfInt64(290)) > 0 {
			_127_cf2 = _127_cf2
			var _142_v54 _dafny.Int
			_ = _142_v54
			var _out1 _dafny.Int
			_ = _out1
			_out1 = Companion_Default___.M0((_dafny.Zero).Minus(_124_cf5), _63_globalState)
			_142_v54 = _out1
			var _143_v55 _dafny.Array
			_ = _143_v55
			var _nwElement0_6 *C0 = _70_v5
			_ = _nwElement0_6
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_6, 0)
			(_nw28).ArraySet1(_70_v5, 1)
			(_nw28).ArraySet1(_70_v5, 2)
			(_nw28).ArraySet1(_70_v5, 3)
			(_nw28).ArraySet1(_70_v5, 4)
			(_nw28).ArraySet1(_70_v5, 5)
			(_nw28).ArraySet1(_70_v5, 6)
			(_nw28).ArraySet1(_70_v5, 7)
			(_nw28).ArraySet1(_70_v5, 8)
			(_nw28).ArraySet1(_70_v5, 9)
			(_nw28).ArraySet1(_70_v5, 10)
			(_nw28).ArraySet1((func() *C0 {
				if _65_v2 {
					return _70_v5
				}
				return _70_v5
			})(), 11)
			(_nw28).ArraySet1(_70_v5, 12)
			_143_v55 = _nw28
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_143_v55), 0))
			_ = _index31
			(_143_v55).ArraySet1(_70_v5, (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_143_v55), 0))
			_ = _index32
			(_143_v55).ArraySet1(_70_v5, (_index32).Int())
			var _144_v56 _dafny.Array
			_ = _144_v56
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(15))
			_ = _nw29
			_144_v56 = _nw29
			var _145_v57 D1
			_ = _145_v57
			_145_v57 = Companion_D1_.Create_DC4_(_65_v2, (_dafny.Zero).Minus(_125_cf4))
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_144_v56), 0))
			_ = _index33
			(_144_v56).ArraySet1(_145_v57, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_144_v56), 0))
			_ = _index34
			(_144_v56).ArraySet1(_145_v57, (_index34).Int())
			var _146_v58 _dafny.Int
			_ = _146_v58
			var _out2 _dafny.Int
			_ = _out2
			_out2 = Companion_Default___.M0(_126_cf3, _63_globalState)
			_146_v58 = _out2
		} else {
			var _147_v59 _dafny.Set
			_ = _147_v59
			_147_v59 = _dafny.SetOf(_70_v5, _70_v5)
			var _148_v60 _dafny.Map
			_ = _148_v60
			_148_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_cf3, _147_v59)
			var _149_v61 D1
			_ = _149_v61
			_149_v61 = Companion_D1_.Create_DC3_(_148_v60)
			_149_v61 = _149_v61
			_65_v2 = _127_cf2
			_70_v5 = _70_v5
			var _150_v62 _dafny.Array
			_ = _150_v62
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw30
			_150_v62 = _nw30
			var _151_v63 _dafny.Map
			_ = _151_v63
			_151_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v2, _150_v62)
			var _152_v64 _dafny.Map
			_ = _152_v64
			_152_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_65_v2) == (_65_v2), _151_v63)
			var _153_v65 _dafny.Sequence
			_ = _153_v65
			_153_v65 = _dafny.SeqOf(_127_cf2)
			var _154_v66 _dafny.Map
			_ = _154_v66
			_154_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_cf3, _153_v65)
			var _155_v67 _dafny.Sequence
			_ = _155_v67
			_155_v67 = _dafny.SeqOf(_125_cf4)
			_152_v64 = (_152_v64).Update(_dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
				if (_154_v66).Contains(_dafny.IntOfUint32((_155_v67).Cardinality())) {
					return (_154_v66).Get(_dafny.IntOfUint32((_155_v67).Cardinality())).(_dafny.Sequence)
				}
				return _153_v65
			})(), _dafny.SeqOf(!(_65_v2))), _151_v63)
			_153_v65 = _dafny.Companion_Sequence_.Concatenate(_153_v65, _dafny.SeqOf(!(true), _127_cf2, _65_v2, !(true), _127_cf2))
		}
		var _156_v68 *C0
		_ = _156_v68
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__()
		_156_v68 = _nw31
	} else {
		var _157___mcc_h6 D0 = _source1.Get_().(D0_DC2).Cf6
		_ = _157___mcc_h6
		var _158_cf6 D0 = _157___mcc_h6
		_ = _158_cf6
		var _159_v69 _dafny.Map
		_ = _159_v69
		_159_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v2, _70_v5)
		var _160_v70 _dafny.Map
		_ = _160_v70
		_160_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v2, _70_v5)).Merge(_159_v69), _62_v0)
		_62_v0 = (func() _dafny.Array {
			if (_160_v70).Contains((_159_v69).Update(_65_v2, _70_v5)) {
				return (_160_v70).Get((_159_v69).Update(_65_v2, _70_v5)).(_dafny.Array)
			}
			return (func() _dafny.Array {
				if _65_v2 {
					return _62_v0
				}
				return _62_v0
			})()
		})()
		if Companion_Default___.Fm2(_63_globalState) {
			_64_v1 = (_64_v1).Plus(_64_v1)
			(_63_globalState).F4 = _dafny.IntOfUint32((_95_v23).Cardinality())
			var _161_v71 _dafny.Array
			_ = _161_v71
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw32
			_161_v71 = _nw32
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_161_v71), 0))
			_ = _index35
			(_161_v71).ArraySet1(Companion_Default___.SafeModuloInt(_64_v1, _64_v1), (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_161_v71), 0))
			_ = _index36
			(_161_v71).ArraySet1(_64_v1, (_index36).Int())
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__()
			_70_v5 = _nw33
			var _162_v72 _dafny.MultiSet
			_ = _162_v72
			_162_v72 = _dafny.MultiSetOf(_dafny.IntOfInt64(39), (_dafny.IntOfUint32((_95_v23).Cardinality())).Plus(_64_v1))
			_162_v72 = (_162_v72).Intersection(Companion_Default___.Fm4(_dafny.IntOfInt64(877), _63_globalState))
		} else {
			var _163_v73 _dafny.Map
			_ = _163_v73
			_163_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, !(_65_v2))
			var _164_v74 _dafny.Sequence
			_ = _164_v74
			_164_v74 = _dafny.SeqOf(_163_v73, (_163_v73).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, (_72_v7).Cardinality())).Cardinality(), _65_v2), _163_v73, _163_v73, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, true))
			var _165_v75 _dafny.Map
			_ = _165_v75
			_165_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v74, _64_v1)
			_165_v75 = (_165_v75).Update(_164_v74, _dafny.IntOfInt64(-41))
			var _166_v76 _dafny.CodePoint
			_ = _166_v76
			_166_v76 = _dafny.CodePoint('n')
			_166_v76 = _166_v76
			var _167_v77 _dafny.Array
			_ = _167_v77
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_5
			var _nw34 _dafny.Array
			_ = _nw34
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw34 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Sequence = (func(_168_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_169_i7 _dafny.Int) _dafny.Sequence {
						return _168_v23
					}
				})(_95_v23)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw34 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw34).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw34).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_167_v77 = _nw34
			var _170_v78 _dafny.Array
			_ = _170_v78
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw35
			_170_v78 = _nw35
			var _rhs27 _dafny.Array = _167_v77
			_ = _rhs27
			var _rhs28 _dafny.Array = _170_v78
			_ = _rhs28
			_167_v77 = _rhs27
			_170_v78 = _rhs28
			(_63_globalState).F4 = _dafny.IntOfInt64(382)
			var _171_v79 _dafny.Map
			_ = _171_v79
			_171_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(303), _dafny.SetOf(_70_v5))
			var _172_v80 D1
			_ = _172_v80
			_172_v80 = Companion_D1_.Create_DC3_(_171_v79)
			var _173_v81 _dafny.Int
			_ = _173_v81
			var _out3 _dafny.Int
			_ = _out3
			_out3 = Companion_Default___.M0((_dafny.IntOfInt64(754)).Plus((_dafny.MultiSetOf(_172_v80)).Cardinality()), _63_globalState)
			_173_v81 = _out3
		}
		var _174_v82 _dafny.MultiSet
		_ = _174_v82
		_174_v82 = _dafny.MultiSetOf(_64_v1)
		_174_v82 = _174_v82
		var _175_v83 _dafny.Array
		_ = _175_v83
		var _nwElement0_7 _dafny.Int = _64_v1
		_ = _nwElement0_7
		var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(4))
		_ = _nw36
		(_nw36).ArraySet1(_nwElement0_7, 0)
		(_nw36).ArraySet1(_64_v1, 1)
		(_nw36).ArraySet1(_64_v1, 2)
		(_nw36).ArraySet1(_64_v1, 3)
		_175_v83 = _nw36
		var _176_v84 D0
		_ = _176_v84
		_176_v84 = Companion_D0_.Create_DC0_(_175_v83, _64_v1)
		_176_v84 = _176_v84
	}
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
	_ = _index37
	(_62_v0).ArraySet1(_65_v2, (_index37).Int())
	var _177_v85 _dafny.Sequence
	_ = _177_v85
	_177_v85 = _dafny.SeqOf(Companion_Default___.Fm2(_63_globalState), Companion_Default___.Fm2(_63_globalState))
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
	_ = _index38
	(_62_v0).ArraySet1((_177_v85).Select((Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_177_v85).Cardinality()))).Uint32()).(bool), (_index38).Int())
	if _65_v2 {
		var _178_v86 *C0
		_ = _178_v86
		var _nw37 *C0 = New_C0_()
		_ = _nw37
		_nw37.Ctor__()
		_178_v86 = _nw37
		_177_v85 = (func() _dafny.Sequence {
			if (_64_v1).Cmp(_64_v1) >= 0 {
				return _177_v85
			}
			return _177_v85
		})()
		(_63_globalState).F4 = _dafny.IntOfInt64(145)
		_70_v5 = _178_v86
		var _179_v87 _dafny.Sequence
		_ = _179_v87
		_179_v87 = _dafny.SeqOf(_64_v1, (func() _dafny.Int {
			if _65_v2 {
				return _64_v1
			}
			return _64_v1
		})(), (_64_v1).Plus(_dafny.IntOfInt64(995)))
		_179_v87 = _179_v87
	} else {
		var _180_v88 _dafny.Map
		_ = _180_v88
		_180_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v23, _64_v1)
		_180_v88 = (_180_v88).Update(_95_v23, _64_v1)
		var _181_v89 _dafny.Sequence
		_ = _181_v89
		_181_v89 = _dafny.SeqOf(_64_v1)
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
		_ = _index39
		(_62_v0).ArraySet1((!_dafny.Companion_Sequence_.Equal(_181_v89, _181_v89)) == (_65_v2), (_index39).Int())
		var _182_v90 _dafny.Array
		_ = _182_v90
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_6
		var _nw38 _dafny.Array
		_ = _nw38
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw38 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = func(_183_i8 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_183_i8, _dafny.IntOfInt64(-554))
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw38 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw38).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw38).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_182_v90 = _nw38
		var _184_v91 _dafny.Sequence
		_ = _184_v91
		_184_v91 = _dafny.SeqOf(_182_v90, _182_v90)
		var _185_v92 _dafny.MultiSet
		_ = _185_v92
		_185_v92 = _dafny.MultiSetOf(_64_v1)
		var _186_v93 _dafny.Map
		_ = _186_v93
		_186_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, _70_v5)
		var _187_v94 D0
		_ = _187_v94
		_187_v94 = Companion_D0_.Create_DC0_((_184_v91).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.IntOfUint32((_184_v91).Cardinality()))).Uint32()).(_dafny.Array), (func() _dafny.Int {
			if (_185_v92).Contains((_186_v93).Cardinality()) {
				return (_185_v92).Multiplicity((_186_v93).Cardinality())
			}
			return _64_v1
		})())
		var _188_v95 D0
		_ = _188_v95
		_188_v95 = Companion_D0_.Create_DC2_(_187_v94)
		var _source2 D0 = _188_v95
		_ = _source2
		if _source2.Is_DC0() {
			var _189___mcc_h7 _dafny.Array = _source2.Get_().(D0_DC0).Cf0
			_ = _189___mcc_h7
			var _190___mcc_h8 _dafny.Int = _source2.Get_().(D0_DC0).Cf1
			_ = _190___mcc_h8
			var _191_cf1 _dafny.Int = _190___mcc_h8
			_ = _191_cf1
			var _192_cf0 _dafny.Array = _189___mcc_h7
			_ = _192_cf0
			var _193_v96 _dafny.CodePoint
			_ = _193_v96
			_193_v96 = _dafny.CodePoint('u')
			_193_v96 = _193_v96
			_73_v8 = (_73_v8).Update(_dafny.IntOfInt64(207), _dafny.IntOfInt64(274))
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_192_cf0), 0))
			_ = _index40
			(_192_cf0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_95_v23, (Companion_Default___.SafeIndex((Companion_Default___.Fm5(_191_cf1, _63_globalState)).Cardinality(), _dafny.IntOfUint32((_95_v23).Cardinality()))).Uint32(), _193_v96)).Cardinality()), (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_192_cf0), 0))
			_ = _index41
			(_192_cf0).ArraySet1(_64_v1, (_index41).Int())
			_62_v0 = _62_v0
		} else if _source2.Is_DC1() {
			var _194___mcc_h9 bool = _source2.Get_().(D0_DC1).Cf2
			_ = _194___mcc_h9
			var _195___mcc_h10 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
			_ = _195___mcc_h10
			var _196___mcc_h11 _dafny.Int = _source2.Get_().(D0_DC1).Cf4
			_ = _196___mcc_h11
			var _197___mcc_h12 _dafny.Int = _source2.Get_().(D0_DC1).Cf5
			_ = _197___mcc_h12
			var _198_cf5 _dafny.Int = _197___mcc_h12
			_ = _198_cf5
			var _199_cf4 _dafny.Int = _196___mcc_h11
			_ = _199_cf4
			var _200_cf3 _dafny.Int = _195___mcc_h10
			_ = _200_cf3
			var _201_cf2 bool = _194___mcc_h9
			_ = _201_cf2
			_64_v1 = _64_v1
			(_63_globalState).F4 = _dafny.IntOfInt64(320)
			_201_cf2 = _201_cf2
			var _202_v97 _dafny.Sequence
			_ = _202_v97
			_202_v97 = _dafny.SeqOf(_70_v5)
			var _203_v98 _dafny.MultiSet
			_ = _203_v98
			_203_v98 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool)))
			_70_v5 = (func() *C0 {
				if true {
					return _70_v5
				}
				return (_202_v97).Select((Companion_Default___.SafeIndex((_203_v98).Cardinality(), _dafny.IntOfUint32((_202_v97).Cardinality()))).Uint32()).(*C0)
			})()
		} else {
			var _204___mcc_h13 D0 = _source2.Get_().(D0_DC2).Cf6
			_ = _204___mcc_h13
			var _205_cf6 D0 = _204___mcc_h13
			_ = _205_cf6
			var _206_v99 _dafny.Array
			_ = _206_v99
			var _nwElement0_8 *C0 = _70_v5
			_ = _nwElement0_8
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(5))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_8, 0)
			(_nw39).ArraySet1(_70_v5, 1)
			(_nw39).ArraySet1(_70_v5, 2)
			(_nw39).ArraySet1(_70_v5, 3)
			(_nw39).ArraySet1(_70_v5, 4)
			_206_v99 = _nw39
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_206_v99), 0))
			_ = _index42
			(_206_v99).ArraySet1(_70_v5, (_index42).Int())
			var _207_v100 _dafny.Map
			_ = _207_v100
			_207_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-461), _95_v23)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _index43
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_206_v99), 0))
			_ = _index44
			var _rhs29 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_208_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), _dafny.UnicodeSeqOfUtf8Bytes("rrodo")), (func() _dafny.Sequence {
				if (_207_v100).Contains(_64_v1) {
					return (_207_v100).Get(_64_v1).(_dafny.Sequence)
				}
				return _95_v23
			})())
			_ = _rhs29
			var _rhs30 *C0 = _70_v5
			_ = _rhs30
			var _lhs21 _dafny.Array = _62_v0
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _lhs22
			var _lhs23 _dafny.Array = _206_v99
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_206_v99), 0))
			_ = _lhs24
			(_lhs21).ArraySet1(_rhs29, (_lhs22).Int())
			(_lhs23).ArraySet1(_rhs30, (_lhs24).Int())
			var _rhs31 _dafny.Int = (_64_v1).Plus((_dafny.Zero).Minus(_64_v1))
			_ = _rhs31
			var _rhs32 _dafny.Int = _64_v1
			_ = _rhs32
			var _rhs33 _dafny.Int = _64_v1
			_ = _rhs33
			var _rhs34 bool = _65_v2
			_ = _rhs34
			var _lhs25 *GlobalState = _63_globalState
			_ = _lhs25
			var _lhs26 *GlobalState = _63_globalState
			_ = _lhs26
			var _lhs27 *GlobalState = _63_globalState
			_ = _lhs27
			_lhs25.F4 = _rhs31
			_lhs26.F4 = _rhs32
			_lhs27.F4 = _rhs33
			_65_v2 = _rhs34
			var _209_v101 *C0
			_ = _209_v101
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__()
			_209_v101 = _nw40
			(_63_globalState).F4 = _64_v1
		}
		var _210_v102 D0
		_ = _210_v102
		_210_v102 = Companion_D0_.Create_DC0_(_182_v90, _64_v1)
		var _source3 D0 = _210_v102
		_ = _source3
		if _source3.Is_DC0() {
			var _211___mcc_h14 _dafny.Array = _source3.Get_().(D0_DC0).Cf0
			_ = _211___mcc_h14
			var _212___mcc_h15 _dafny.Int = _source3.Get_().(D0_DC0).Cf1
			_ = _212___mcc_h15
			var _213_cf1 _dafny.Int = _212___mcc_h15
			_ = _213_cf1
			var _214_cf0 _dafny.Array = _211___mcc_h14
			_ = _214_cf0
			var _215_v103 _dafny.MultiSet
			_ = _215_v103
			_215_v103 = _dafny.MultiSetOf(_65_v2)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _index45
			(_62_v0).ArraySet1((_215_v103).IsSubsetOf(_215_v103), (_index45).Int())
			var _216_v104 D1
			_ = _216_v104
			_216_v104 = Companion_D1_.Create_DC4_((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _dafny.IntOfUint32((_177_v85).Cardinality()))
			(_63_globalState).F4 = (_216_v104).Dtor_cf9()
			var _217_v105 _dafny.Set
			_ = _217_v105
			_217_v105 = _dafny.SetOf(_70_v5)
			_65_v2 = (_217_v105).IsProperSubsetOf((_217_v105).Union(_217_v105))
			var _218_v106 *C0
			_ = _218_v106
			var _nw41 *C0 = New_C0_()
			_ = _nw41
			_nw41.Ctor__()
			_218_v106 = _nw41
		} else if _source3.Is_DC1() {
			var _219___mcc_h16 bool = _source3.Get_().(D0_DC1).Cf2
			_ = _219___mcc_h16
			var _220___mcc_h17 _dafny.Int = _source3.Get_().(D0_DC1).Cf3
			_ = _220___mcc_h17
			var _221___mcc_h18 _dafny.Int = _source3.Get_().(D0_DC1).Cf4
			_ = _221___mcc_h18
			var _222___mcc_h19 _dafny.Int = _source3.Get_().(D0_DC1).Cf5
			_ = _222___mcc_h19
			var _223_cf5 _dafny.Int = _222___mcc_h19
			_ = _223_cf5
			var _224_cf4 _dafny.Int = _221___mcc_h18
			_ = _224_cf4
			var _225_cf3 _dafny.Int = _220___mcc_h17
			_ = _225_cf3
			var _226_cf2 bool = _219___mcc_h16
			_ = _226_cf2
			var _227_v107 *C0
			_ = _227_v107
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__()
			_227_v107 = _nw42
			var _228_v108 _dafny.CodePoint
			_ = _228_v108
			_228_v108 = _dafny.CodePoint('t')
			var _229_v109 D1
			_ = _229_v109
			_229_v109 = Companion_D1_.Create_DC4_(true, _dafny.IntOfInt64(355))
			var _230_v110 _dafny.Int
			_ = _230_v110
			var _out4 _dafny.Int
			_ = _out4
			_out4 = Companion_Default___.M0(((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm6(false, _228_v108, _229_v109, _226_cf2, _63_globalState)).Cardinality()))).Times((_dafny.Zero).Minus(_223_cf5)), _63_globalState)
			_230_v110 = _out4
			var _231_v111 _dafny.Array
			_ = _231_v111
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw43
			_231_v111 = _nw43
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _index46
			var _rhs35 _dafny.Array = _231_v111
			_ = _rhs35
			var _rhs36 bool = (_96_v24).Dtor_cf2()
			_ = _rhs36
			var _lhs28 _dafny.Array = _62_v0
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _lhs29
			_231_v111 = _rhs35
			(_lhs28).ArraySet1(_rhs36, (_lhs29).Int())
			var _rhs37 _dafny.Int = (_dafny.Zero).Minus(_223_cf5)
			_ = _rhs37
			var _rhs38 bool = _65_v2
			_ = _rhs38
			var _rhs39 *C0 = _227_v107
			_ = _rhs39
			var _lhs30 *GlobalState = _63_globalState
			_ = _lhs30
			_lhs30.F4 = _rhs37
			_226_cf2 = _rhs38
			_70_v5 = _rhs39
		} else {
			var _232___mcc_h20 D0 = _source3.Get_().(D0_DC2).Cf6
			_ = _232___mcc_h20
			var _233_cf6 D0 = _232___mcc_h20
			_ = _233_cf6
			var _234_v112 *C0
			_ = _234_v112
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__()
			_234_v112 = _nw44
			var _235_v113 D1
			_ = _235_v113
			_235_v113 = Companion_D1_.Create_DC4_(_65_v2, _64_v1)
			var _236_v114 _dafny.CodePoint
			_ = _236_v114
			_236_v114 = _dafny.CodePoint('y')
			var _237_v115 _dafny.MultiSet
			_ = _237_v115
			_237_v115 = _dafny.MultiSetOf(_65_v2)
			var _238_v116 _dafny.Map
			_ = _238_v116
			_238_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v114, _237_v115)
			var _239_v117 _dafny.Map
			_ = _239_v117
			_239_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_234_v112, _95_v23)
			var _240_v118 _dafny.Map
			_ = _240_v118
			_240_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(_65_v2, false, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_239_v117).Contains(_70_v5) {
					return (_239_v117).Get(_70_v5).(_dafny.Sequence)
				}
				return _95_v23
			})(), (Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_239_v117).Contains(_70_v5) {
					return (_239_v117).Get(_70_v5).(_dafny.Sequence)
				}
				return _95_v23
			})()).Cardinality()))).Uint32(), _236_v114)).Cardinality()), _63_globalState), _182_v90)
			var _241_v119 _dafny.Array
			_ = _241_v119
			var _nwElement0_9 _dafny.Array = _182_v90
			_ = _nwElement0_9
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(14))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_9, 0)
			(_nw45).ArraySet1((func() _dafny.Array {
				if (_235_v113).Dtor_cf8() {
					return _182_v90
				}
				return _182_v90
			})(), 1)
			(_nw45).ArraySet1(_182_v90, 2)
			(_nw45).ArraySet1((_184_v91).Select((Companion_Default___.SafeIndex(((func() _dafny.MultiSet {
				if (_238_v116).Contains(_236_v114) {
					return (_238_v116).Get(_236_v114).(_dafny.MultiSet)
				}
				return _237_v115
			})()).Cardinality(), _dafny.IntOfUint32((_184_v91).Cardinality()))).Uint32()).(_dafny.Array), 3)
			(_nw45).ArraySet1(_182_v90, 4)
			(_nw45).ArraySet1(_182_v90, 5)
			(_nw45).ArraySet1(_182_v90, 6)
			(_nw45).ArraySet1(_182_v90, 7)
			(_nw45).ArraySet1(_182_v90, 8)
			(_nw45).ArraySet1((func() _dafny.Array {
				if (_240_v118).Contains(_236_v114) {
					return (_240_v118).Get(_236_v114).(_dafny.Array)
				}
				return _182_v90
			})(), 9)
			(_nw45).ArraySet1(_182_v90, 10)
			(_nw45).ArraySet1(_182_v90, 11)
			(_nw45).ArraySet1(_182_v90, 12)
			(_nw45).ArraySet1(_182_v90, 13)
			_241_v119 = _nw45
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_241_v119), 0))
			_ = _index47
			(_241_v119).ArraySet1(_182_v90, (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_241_v119), 0))
			_ = _index48
			(_241_v119).ArraySet1(_182_v90, (_index48).Int())
			var _242_v120 _dafny.Int
			_ = _242_v120
			var _out5 _dafny.Int
			_ = _out5
			_out5 = Companion_Default___.M0(_64_v1, _63_globalState)
			_242_v120 = _out5
			var _243_v121 *C0
			_ = _243_v121
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__()
			_243_v121 = _nw46
		}
		_181_v89 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_181_v89, _181_v89), (Companion_Default___.SafeIndex((_64_v1).Minus(_64_v1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_181_v89, _181_v89)).Cardinality()))).Uint32(), (_70_v5).Fm1((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _64_v1, _64_v1, _63_globalState))
	}
	var _244_v122 _dafny.Sequence
	_ = _244_v122
	_244_v122 = _dafny.SeqOf(_177_v85)
	var _245_v123 _dafny.Sequence
	_ = _245_v123
	_245_v123 = _dafny.SeqOf(_64_v1)
	var _246_v124 _dafny.Array
	_ = _246_v124
	var _nwElement0_10 _dafny.Int = _64_v1
	_ = _nwElement0_10
	var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(14))
	_ = _nw47
	(_nw47).ArraySet1(_nwElement0_10, 0)
	(_nw47).ArraySet1(_64_v1, 1)
	(_nw47).ArraySet1(Companion_Default___.SafeModuloInt(_64_v1, _64_v1), 2)
	(_nw47).ArraySet1(_64_v1, 3)
	(_nw47).ArraySet1((_64_v1).Plus(_64_v1), 4)
	(_nw47).ArraySet1(Companion_Default___.SafeDivisionInt(_64_v1, _64_v1), 5)
	(_nw47).ArraySet1((_70_v5).Fm1(Companion_Default___.Fm2(_63_globalState), _64_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lyoe")).Cardinality()), _63_globalState), 6)
	(_nw47).ArraySet1(_64_v1, 7)
	(_nw47).ArraySet1(_64_v1, 8)
	(_nw47).ArraySet1((Companion_Default___.Fm8(_65_v2, _65_v2, _63_globalState)).Cardinality(), 9)
	(_nw47).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_177_v85, (_244_v122).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_245_v123).Cardinality()), _dafny.IntOfUint32((_244_v122).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()), 10)
	(_nw47).ArraySet1((_64_v1).Plus(_64_v1), 11)
	(_nw47).ArraySet1((_dafny.IntOfInt64(-480)).Minus(_64_v1), 12)
	(_nw47).ArraySet1(_64_v1, 13)
	_246_v124 = _nw47
	var _247_v125 _dafny.Set
	_ = _247_v125
	_247_v125 = _dafny.SetOf(_dafny.IntOfInt64(617))
	_246_v124 = (func() _dafny.Array {
		if (_247_v125).IsSubsetOf(_247_v125) {
			return _246_v124
		}
		return (func() _dafny.Array {
			if _65_v2 {
				return _246_v124
			}
			return _246_v124
		})()
	})()
	var _248_v128 _dafny.Array
	_ = _248_v128
	var _nwElement0_11 _dafny.Map = _73_v8
	_ = _nwElement0_11
	var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(8))
	_ = _nw48
	(_nw48).ArraySet1(_nwElement0_11, 0)
	(_nw48).ArraySet1(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_245_v123).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _249_v127 _dafny.Int
				_249_v127 = interface{}(_compr_4).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_245_v123, _249_v127) {
					_coll4.Add(Companion_Default___.SafeDivisionInt(_249_v127, (_dafny.SetOf(_64_v1, _64_v1)).Cardinality()), (Companion_D2_.Create_DC5_(_dafny.MultiSetOf(_64_v1))).Dtor_cf10())
				}
			}
			return _coll4.ToMap()
		}()).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _250_v126 _dafny.Int
			_250_v126 = interface{}(_compr_3).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_245_v123).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _249_v127 _dafny.Int
					_249_v127 = interface{}(_compr_5).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_245_v123, _249_v127) {
						_coll5.Add(Companion_Default___.SafeDivisionInt(_249_v127, (_dafny.SetOf(_64_v1, _64_v1)).Cardinality()), (Companion_D2_.Create_DC5_(_dafny.MultiSetOf(_64_v1))).Dtor_cf10())
					}
				}
				return _coll5.ToMap()
			}()).Contains(_250_v126) {
				_coll3.Add(Companion_Default___.SafeDivisionInt(_250_v126, (_73_v8).Cardinality()), _dafny.IntOfUint32((_177_v85).Cardinality()))
			}
		}
		return _coll3.ToMap()
	}(), 1)
	(_nw48).ArraySet1(_73_v8, 2)
	(_nw48).ArraySet1(_73_v8, 3)
	(_nw48).ArraySet1(_73_v8, 4)
	(_nw48).ArraySet1(_73_v8, 5)
	(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _64_v1), 6)
	(_nw48).ArraySet1(_73_v8, 7)
	_248_v128 = _nw48
	_248_v128 = _248_v128
	_247_v125 = _247_v125
	if _65_v2 {
		var _251_v129 D2
		_ = _251_v129
		_251_v129 = Companion_D2_.Create_DC6_(_64_v1, _64_v1, _64_v1)
		var _252_v130 _dafny.Set
		_ = _252_v130
		_252_v130 = _dafny.SetOf(_70_v5, _70_v5)
		var _253_v131 _dafny.MultiSet
		_ = _253_v131
		_253_v131 = _dafny.MultiSetOf(_64_v1)
		var _pat_let_tv0 = _252_v130
		_ = _pat_let_tv0
		var _pat_let_tv1 = _64_v1
		_ = _pat_let_tv1
		var _rhs40 bool = (_65_v2) && ((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool))
		_ = _rhs40
		var _rhs41 D2 = func(_pat_let0_0 D2) D2 {
			return func(_254_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let1_0 _dafny.Int) D2 {
					return func(_255_dt__update_hcf11_h0 _dafny.Int) D2 {
						return func(_pat_let2_0 _dafny.Int) D2 {
							return func(_256_dt__update_hcf12_h0 _dafny.Int) D2 {
								return Companion_D2_.Create_DC6_(_255_dt__update_hcf11_h0, _256_dt__update_hcf12_h0, (_254_dt__update__tmp_h0).Dtor_cf13())
							}(_pat_let2_0)
						}(_pat_let_tv1)
					}(_pat_let1_0)
				}((func() _dafny.Int {
					if false {
						return (_pat_let_tv0).Cardinality()
					}
					return _dafny.IntOfInt64(264)
				})())
			}(_pat_let0_0)
		}(Companion_Default___.Fm9(_64_v1, (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _63_globalState))
		_ = _rhs41
		var _rhs42 bool = (_dafny.MultiSetOf(_64_v1)).IsProperSubsetOf(_253_v131)
		_ = _rhs42
		_65_v2 = _rhs40
		_251_v129 = _rhs41
		_65_v2 = _rhs42
		var _257_v132 _dafny.Int
		_ = _257_v132
		var _out6 _dafny.Int
		_ = _out6
		_out6 = Companion_Default___.M0(_64_v1, _63_globalState)
		_257_v132 = _out6
		if (func() bool {
			if (Companion_Default___.Fm2(_63_globalState)) || ((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool)) {
				return (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool)
			}
			return (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool)
		})() {
			var _258_v133 _dafny.Array
			_ = _258_v133
			var _nw49 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
			_ = _nw49
			_258_v133 = _nw49
			var _259_v134 _dafny.MultiSet
			_ = _259_v134
			_259_v134 = _dafny.MultiSetOf(_258_v133)
			var _260_v135 D3
			_ = _260_v135
			_260_v135 = Companion_D3_.Create_DC7_(_259_v134)
			_259_v134 = ((_259_v134).Difference(_259_v134)).Intersection((_260_v135).Dtor_cf14())
			(_63_globalState).F4 = (Companion_D0_.Create_DC0_(_246_v124, _dafny.IntOfUint32((_177_v85).Cardinality()))).Dtor_cf1()
			_65_v2 = _65_v2
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_246_v124), 0))
			_ = _index49
			(_246_v124).ArraySet1((func() _dafny.Int {
				if (_253_v131).Contains((_70_v5).Fm1((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _dafny.IntOfInt64(720), _64_v1, _63_globalState)) {
					return (_253_v131).Multiplicity((_70_v5).Fm1((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _dafny.IntOfInt64(720), _64_v1, _63_globalState))
				}
				return _257_v132
			})(), (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_246_v124), 0))
			_ = _index50
			(_246_v124).ArraySet1((_64_v1).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_177_v85).Cardinality()))), (_index50).Int())
			_95_v23 = _95_v23
		} else {
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _index51
			(_62_v0).ArraySet1(_65_v2, (_index51).Int())
			var _261_v137 _dafny.CodePoint
			_ = _261_v137
			_261_v137 = _dafny.CodePoint('u')
			var _262_v138 _dafny.MultiSet
			_ = _262_v138
			_262_v138 = _dafny.MultiSetOf(_261_v137)
			_65_v2 = (Companion_D0_.Create_DC1_((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _64_v1, _dafny.IntOfUint32((_95_v23).Cardinality()), (func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_262_v138).Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _263_v136 _dafny.CodePoint
					_263_v136 = interface{}(_compr_6).(_dafny.CodePoint)
					if (_262_v138).Contains(_263_v136) {
						_coll6.Add(_263_v136, (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool))
					}
				}
				return _coll6.ToMap()
			}()).Cardinality())).Dtor_cf2()
			var _264_v139 _dafny.Set
			_ = _264_v139
			_264_v139 = _dafny.SetOf((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool))
			var _265_v140 _dafny.Map
			_ = _265_v140
			_265_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_264_v139, _70_v5)
			var _266_v141 _dafny.MultiSet
			_ = _266_v141
			_266_v141 = _dafny.MultiSetOf(_70_v5, (func() *C0 {
				if (_265_v140).Contains(_264_v139) {
					return (_265_v140).Get(_264_v139).(*C0)
				}
				return _70_v5
			})(), _70_v5)
			var _267_v142 _dafny.Map
			_ = _267_v142
			_267_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_63_globalState), (_266_v141).Cardinality())
			_65_v2 = (((func() _dafny.Map {
				if _65_v2 {
					return _267_v142
				}
				return _267_v142
			})()).Cardinality()).Cmp(_257_v132) == 0
			var _268_v143 _dafny.Array
			_ = _268_v143
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_7
			var _nw50 _dafny.Array
			_ = _nw50
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw50 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Sequence = (func(_269_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_270_i10 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_269_v23, _269_v23)
					}
				})(_95_v23)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw50 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw50).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw50).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_268_v143 = _nw50
			_268_v143 = (func() _dafny.Array {
				if _65_v2 {
					return _268_v143
				}
				return _268_v143
			})()
			_65_v2 = !((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool)) || ((false) == ((_177_v85).Select((Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_177_v85).Cardinality()))).Uint32()).(bool)))
		}
		_64_v1 = _64_v1
		_70_v5 = _70_v5
	} else {
		(_63_globalState).F4 = _64_v1
		var _271_v144 _dafny.Array
		_ = _271_v144
		var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw51
		_271_v144 = _nw51
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_271_v144), 0))
		_ = _index52
		(_271_v144).ArraySet1(_95_v23, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_271_v144), 0))
		_ = _index53
		(_271_v144).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("aqmgw"), (_index53).Int())
		var _272_v145 _dafny.MultiSet
		_ = _272_v145
		_272_v145 = _dafny.MultiSetOf(_70_v5, _70_v5)
		var _273_v146 D0
		_ = _273_v146
		_273_v146 = Companion_D0_.Create_DC1_((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _64_v1, _dafny.IntOfInt64(188), (_272_v145).Cardinality())
		var _pat_let_tv2 = _64_v1
		_ = _pat_let_tv2
		var _274_v147 D0
		_ = _274_v147
		_274_v147 = Companion_D0_.Create_DC2_((func() D0 {
			if (func(_pat_let3_0 D0) D0 {
				return func(_275_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let4_0 _dafny.Int) D0 {
						return func(_276_dt__update_hcf5_h0 _dafny.Int) D0 {
							return func(_pat_let5_0 bool) D0 {
								return func(_277_dt__update_hcf2_h0 bool) D0 {
									return Companion_D0_.Create_DC1_(_277_dt__update_hcf2_h0, (_275_dt__update__tmp_h1).Dtor_cf3(), (_275_dt__update__tmp_h1).Dtor_cf4(), _276_dt__update_hcf5_h0)
								}(_pat_let5_0)
							}(true)
						}(_pat_let4_0)
					}(_pat_let_tv2)
				}(_pat_let3_0)
			}(_96_v24)).Dtor_cf2() {
				return _273_v146
			}
			return _273_v146
		})())
		var _source4 D0 = _274_v147
		_ = _source4
		if _source4.Is_DC0() {
			var _278___mcc_h21 _dafny.Array = _source4.Get_().(D0_DC0).Cf0
			_ = _278___mcc_h21
			var _279___mcc_h22 _dafny.Int = _source4.Get_().(D0_DC0).Cf1
			_ = _279___mcc_h22
			var _280_cf1 _dafny.Int = _279___mcc_h22
			_ = _280_cf1
			var _281_cf0 _dafny.Array = _278___mcc_h21
			_ = _281_cf0
			(_63_globalState).F4 = (_dafny.Zero).Minus(_280_cf1)
			var _282_v148 _dafny.Int
			_ = _282_v148
			var _out7 _dafny.Int
			_ = _out7
			_out7 = Companion_Default___.M0(_280_cf1, _63_globalState)
			_282_v148 = _out7
			var _283_v149 _dafny.Int
			_ = _283_v149
			var _out8 _dafny.Int
			_ = _out8
			_out8 = Companion_Default___.M0(_282_v148, _63_globalState)
			_283_v149 = _out8
			var _284_v150 _dafny.Array
			_ = _284_v150
			var _nw52 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
			_ = _nw52
			_284_v150 = _nw52
			_284_v150 = _284_v150
		} else if _source4.Is_DC1() {
			var _285___mcc_h23 bool = _source4.Get_().(D0_DC1).Cf2
			_ = _285___mcc_h23
			var _286___mcc_h24 _dafny.Int = _source4.Get_().(D0_DC1).Cf3
			_ = _286___mcc_h24
			var _287___mcc_h25 _dafny.Int = _source4.Get_().(D0_DC1).Cf4
			_ = _287___mcc_h25
			var _288___mcc_h26 _dafny.Int = _source4.Get_().(D0_DC1).Cf5
			_ = _288___mcc_h26
			var _289_cf5 _dafny.Int = _288___mcc_h26
			_ = _289_cf5
			var _290_cf4 _dafny.Int = _287___mcc_h25
			_ = _290_cf4
			var _291_cf3 _dafny.Int = _286___mcc_h24
			_ = _291_cf3
			var _292_cf2 bool = _285___mcc_h23
			_ = _292_cf2
			var _293_v151 *C0
			_ = _293_v151
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__()
			_293_v151 = _nw53
			(_63_globalState).F4 = (_dafny.Zero).Minus(_289_cf5)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_271_v144), 0))
			_ = _index54
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_271_v144), 0))
			_ = _index55
			var _rhs43 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_74_v9, _74_v9), _74_v9)
			_ = _rhs43
			var _rhs44 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("bsuwye")
			_ = _rhs44
			var _rhs45 *C0 = _70_v5
			_ = _rhs45
			var _rhs46 _dafny.Array = _62_v0
			_ = _rhs46
			var _rhs47 _dafny.Sequence = _95_v23
			_ = _rhs47
			var _lhs31 _dafny.Array = _271_v144
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_271_v144), 0))
			_ = _lhs32
			var _lhs33 _dafny.Array = _271_v144
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_271_v144), 0))
			_ = _lhs34
			_74_v9 = _rhs43
			(_lhs31).ArraySet1(_rhs44, (_lhs32).Int())
			_70_v5 = _rhs45
			_62_v0 = _rhs46
			(_lhs33).ArraySet1(_rhs47, (_lhs34).Int())
			var _294_v152 _dafny.Map
			_ = _294_v152
			_294_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_289_cf5, _292_cf2)
			_294_v152 = func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(3), _dafny.IntOfInt64(694))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _295_v153 _dafny.Int
					_295_v153 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(3)).Cmp(_295_v153) <= 0) && ((_295_v153).Cmp(_dafny.IntOfInt64(694)) < 0) {
						_coll7.Add((_295_v153).Times(_289_cf5), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool))
					}
				}
				return _coll7.ToMap()
			}()
		} else {
			var _296___mcc_h27 D0 = _source4.Get_().(D0_DC2).Cf6
			_ = _296___mcc_h27
			var _297_cf6 D0 = _296___mcc_h27
			_ = _297_cf6
			var _298_v155 _dafny.Sequence
			_ = _298_v155
			_298_v155 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, _65_v2)).Merge(func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-746), _dafny.IntOfInt64(796))); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _299_v154 _dafny.Int
					_299_v154 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(-746)).Cmp(_299_v154) <= 0) && ((_299_v154).Cmp(_dafny.IntOfInt64(796)) < 0) {
						_coll8.Add(Companion_Default___.SafeModuloInt(_299_v154, _64_v1), _65_v2)
					}
				}
				return _coll8.ToMap()
			}()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v1, !(_65_v2)))
			_298_v155 = _298_v155
			var _300_v156 *C0
			_ = _300_v156
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__()
			_300_v156 = _nw54
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _index56
			(_62_v0).ArraySet1(_65_v2, (_index56).Int())
			(_63_globalState).F4 = (_300_v156).Fm1(_65_v2, _dafny.IntOfInt64(182), _64_v1, _63_globalState)
		}
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
		_ = _index57
		var _rhs48 bool = (_96_v24).Dtor_cf2()
		_ = _rhs48
		var _rhs49 bool = _65_v2
		_ = _rhs49
		var _lhs35 _dafny.Array = _62_v0
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
		_ = _lhs36
		(_lhs35).ArraySet1(_rhs48, (_lhs36).Int())
		_65_v2 = _rhs49
		var _rhs50 _dafny.Sequence = _95_v23
		_ = _rhs50
		var _rhs51 _dafny.Int = (_64_v1).Plus(_64_v1)
		_ = _rhs51
		var _lhs37 *GlobalState = _63_globalState
		_ = _lhs37
		_95_v23 = _rhs50
		_lhs37.F4 = _rhs51
	}
	var _301_v157 _dafny.Set
	_ = _301_v157
	_301_v157 = _dafny.SetOf(_177_v85)
	var _302_v158 _dafny.Sequence
	_ = _302_v158
	_302_v158 = _dafny.SeqOf(Companion_Default___.Fm10(_64_v1, _64_v1, _63_globalState), _244_v122)
	var _303_v160 _dafny.Array
	_ = _303_v160
	var _nwElement0_12 _dafny.Set = _301_v157
	_ = _nwElement0_12
	var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(11))
	_ = _nw55
	(_nw55).ArraySet1(_nwElement0_12, 0)
	(_nw55).ArraySet1((_301_v157).Difference(_dafny.SetOf(_dafny.Companion_Sequence_.Update(_177_v85, (Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_177_v85).Cardinality()))).Uint32(), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool)), _177_v85, _177_v85)), 1)
	(_nw55).ArraySet1(_301_v157, 2)
	(_nw55).ArraySet1((_301_v157).Union(_301_v157), 3)
	(_nw55).ArraySet1(_301_v157, 4)
	(_nw55).ArraySet1(func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(((_302_v158).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_95_v23).Cardinality()), _dafny.IntOfUint32((_302_v158).Cardinality()))).Uint32()).(_dafny.Sequence)).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _304_v159 _dafny.Sequence
			_304_v159 = interface{}(_compr_9).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains((_302_v158).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_95_v23).Cardinality()), _dafny.IntOfUint32((_302_v158).Cardinality()))).Uint32()).(_dafny.Sequence), _304_v159) {
				_coll9.Add(_304_v159)
			}
		}
		return _coll9.ToSet()
	}(), 5)
	(_nw55).ArraySet1(_dafny.SetOf(_177_v85), 6)
	(_nw55).ArraySet1(_301_v157, 7)
	(_nw55).ArraySet1((_dafny.SetOf(_177_v85)).Intersection(_301_v157), 8)
	(_nw55).ArraySet1(_301_v157, 9)
	(_nw55).ArraySet1(_301_v157, 10)
	_303_v160 = _nw55
	var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_303_v160), 0))
	_ = _index58
	(_303_v160).ArraySet1(_301_v157, (_index58).Int())
	var _305_v161 _dafny.Array
	_ = _305_v161
	var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
	_ = _nw56
	_305_v161 = _nw56
	var _306_v162 _dafny.Map
	_ = _306_v162
	_306_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _65_v2), false)
	var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_305_v161), 0))
	_ = _index59
	(_305_v161).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v2, _65_v2), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool))).Merge(_306_v162), (_index59).Int())
	var _307_v163 D4
	_ = _307_v163
	_307_v163 = Companion_D4_.Create_DC9_(_dafny.SetOf(_177_v85))
	var _308_v165 D2
	_ = _308_v165
	_308_v165 = Companion_D2_.Create_DC6_(_64_v1, _64_v1, _64_v1)
	var _309_v166 _dafny.Map
	_ = _309_v166
	_309_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm11(_65_v2, _63_globalState), _308_v165)
	var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_303_v160), 0))
	_ = _index60
	var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_305_v161), 0))
	_ = _index61
	var _rhs52 _dafny.Sequence = (func() _dafny.Sequence {
		if _65_v2 {
			return _245_v123
		}
		return _245_v123
	})()
	_ = _rhs52
	var _rhs53 _dafny.Set = ((_301_v157).Difference((_307_v163).Dtor_cf15())).Union((_301_v157).Intersection(_301_v157))
	_ = _rhs53
	var _rhs54 bool = (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool)
	_ = _rhs54
	var _rhs55 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v7, _65_v2)).Merge((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_309_v166).Keys().Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _310_v164 _dafny.Map
			_310_v164 = interface{}(_compr_10).(_dafny.Map)
			if (_309_v166).Contains(_310_v164) {
				_coll10.Add(_310_v164, (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool))
			}
		}
		return _coll10.ToMap()
	}()).Merge(_306_v162))
	_ = _rhs55
	var _lhs38 _dafny.Array = _303_v160
	_ = _lhs38
	var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_303_v160), 0))
	_ = _lhs39
	var _lhs40 _dafny.Array = _305_v161
	_ = _lhs40
	var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_305_v161), 0))
	_ = _lhs41
	_245_v123 = _rhs52
	(_lhs38).ArraySet1(_rhs53, (_lhs39).Int())
	_65_v2 = _rhs54
	(_lhs40).ArraySet1(_rhs55, (_lhs41).Int())
	if _65_v2 {
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))
		_ = _index62
		(_246_v124).ArraySet1(_64_v1, (_index62).Int())
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))
		_ = _index63
		(_246_v124).ArraySet1(_64_v1, (_index63).Int())
		var _311_v167 *C0
		_ = _311_v167
		var _nw57 *C0 = New_C0_()
		_ = _nw57
		_nw57.Ctor__()
		_311_v167 = _nw57
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))
		_ = _index64
		(_246_v124).ArraySet1(Companion_Default___.SafeDivisionInt((_246_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
			if (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool) {
				return (_246_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))).Int()).(_dafny.Int)
			}
			return (_246_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))).Int()).(_dafny.Int)
		})()), (_index64).Int())
		(_63_globalState).F4 = (_246_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))).Int()).(_dafny.Int)
		var _312_v168 _dafny.MultiSet
		_ = _312_v168
		_312_v168 = _dafny.MultiSetOf((_246_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))).Int()).(_dafny.Int))
		var _313_v169 _dafny.Sequence
		_ = _313_v169
		_313_v169 = _dafny.SeqOf(_dafny.MultiSetOf((_70_v5).Fm1((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _64_v1, _64_v1, _63_globalState)), _312_v168)
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
		_ = _index65
		(_62_v0).ArraySet1(!(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm12((Companion_D4_.Create_DC11_((_246_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_246_v124), 0))).Int()).(_dafny.Int), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool))).Dtor_cf19(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _63_globalState), _313_v169)), (_index65).Int())
	} else {
		var _314_v170 _dafny.Array
		_ = _314_v170
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_8
		var _nw58 _dafny.Array
		_ = _nw58
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw58 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Sequence = (func(_315_v23 _dafny.Sequence, _316_v1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_317_i11 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Update(_315_v23, (Companion_Default___.SafeIndex(_316_v1, _dafny.IntOfUint32((_315_v23).Cardinality()))).Uint32(), _dafny.CodePoint('a'))
				}
			})(_95_v23, _64_v1)
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
		_314_v170 = _nw58
		var _318_v171 _dafny.Set
		_ = _318_v171
		_318_v171 = _dafny.SetOf(Companion_Default___.Fm2(_63_globalState), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), _65_v2)
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
		_ = _index66
		var _rhs56 bool = (_318_v171).IsDisjointFrom((_318_v171).Difference(_318_v171))
		_ = _rhs56
		var _rhs57 _dafny.Array = (func() _dafny.Array {
			if (_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool) {
				return _314_v170
			}
			return _314_v170
		})()
		_ = _rhs57
		var _lhs42 _dafny.Array = _62_v0
		_ = _lhs42
		var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
		_ = _lhs43
		(_lhs42).ArraySet1(_rhs56, (_lhs43).Int())
		_314_v170 = _rhs57
		var _319_v172 _dafny.Int
		_ = _319_v172
		var _out9 _dafny.Int
		_ = _out9
		_out9 = Companion_Default___.M0(_64_v1, _63_globalState)
		_319_v172 = _out9
		var _nw59 *C0 = New_C0_()
		_ = _nw59
		_nw59.Ctor__()
		_70_v5 = _nw59
		var _rhs58 _dafny.Array = _62_v0
		_ = _rhs58
		var _rhs59 _dafny.Int = _64_v1
		_ = _rhs59
		_62_v0 = _rhs58
		_319_v172 = _rhs59
		var _320_v173 _dafny.MultiSet
		_ = _320_v173
		_320_v173 = _dafny.MultiSetOf((_62_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))).Int()).(bool), true)
		if !(_320_v173).Equals(_dafny.MultiSetFromSeq(_177_v85)) {
			_95_v23 = _dafny.Companion_Sequence_.Concatenate(_95_v23, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(323))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_321_i12 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			})))
			var _322_v174 *C0
			_ = _322_v174
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__()
			_322_v174 = _nw60
			var _rhs60 _dafny.Int = _dafny.IntOfInt64(-685)
			_ = _rhs60
			var _rhs61 _dafny.Int = _319_v172
			_ = _rhs61
			var _rhs62 _dafny.Int = _319_v172
			_ = _rhs62
			var _lhs44 *GlobalState = _63_globalState
			_ = _lhs44
			var _lhs45 *GlobalState = _63_globalState
			_ = _lhs45
			_64_v1 = _rhs60
			_lhs44.F4 = _rhs61
			_lhs45.F4 = _rhs62
			var _323_v175 _dafny.Map
			_ = _323_v175
			_323_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_177_v85).Cardinality()), !(_65_v2))
			var _324_v176 _dafny.Map
			_ = _324_v176
			_324_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_323_v175, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_62_v0, _62_v0), (Companion_Default___.SafeIndex(_64_v1, _dafny.IntOfUint32((_dafny.SeqOf(_62_v0, _62_v0)).Cardinality()))).Uint32(), _62_v0)).Cardinality()))
			_65_v2 = (true) == (!(_324_v176).Contains(_323_v175))
			var _325_v177 _dafny.Sequence
			_ = _325_v177
			_325_v177 = _dafny.SeqOf(_245_v123)
			_325_v177 = _dafny.Companion_Sequence_.Concatenate(_325_v177, _325_v177)
		} else {
			var _326_v178 _dafny.MultiSet
			_ = _326_v178
			_326_v178 = _dafny.MultiSetOf(_62_v0, _62_v0, _62_v0, _62_v0, _62_v0)
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _index67
			(_62_v0).ArraySet1((_326_v178).IsDisjointFrom(_326_v178), (_index67).Int())
			var _327_v179 _dafny.MultiSet
			_ = _327_v179
			_327_v179 = _dafny.MultiSetOf(_177_v85, _177_v85, _177_v85)
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_62_v0), 0))
			_ = _index68
			(_62_v0).ArraySet1(((_327_v179).Difference(_327_v179)).IsProperSubsetOf((_327_v179).Union(_327_v179)), (_index68).Int())
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_246_v124), 0))
			_ = _index69
			(_246_v124).ArraySet1(_319_v172, (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_246_v124), 0))
			_ = _index70
			(_246_v124).ArraySet1((_70_v5).Fm1((_64_v1).Cmp(_319_v172) >= 0, _64_v1, (_245_v123).Select((Companion_Default___.SafeIndex(_319_v172, _dafny.IntOfUint32((_245_v123).Cardinality()))).Uint32()).(_dafny.Int), _63_globalState), (_index70).Int())
			var _328_v180 _dafny.Int
			_ = _328_v180
			var _out10 _dafny.Int
			_ = _out10
			_out10 = Companion_Default___.M0(_64_v1, _63_globalState)
			_328_v180 = _out10
			var _329_v181 D0
			_ = _329_v181
			_329_v181 = Companion_D0_.Create_DC1_(false, _64_v1, _319_v172, _64_v1)
			var _330_v182 D0
			_ = _330_v182
			_330_v182 = Companion_D0_.Create_DC2_(_329_v181)
			var _331_v183 _dafny.Sequence
			_ = _331_v183
			_331_v183 = _dafny.SeqOf(_329_v181)
			var _332_v184 D0
			_ = _332_v184
			_332_v184 = Companion_D0_.Create_DC2_((_331_v183).Select((Companion_Default___.SafeIndex(_328_v180, _dafny.IntOfUint32((_331_v183).Cardinality()))).Uint32()).(D0))
			var _333_v185 D0
			_ = _333_v185
			_333_v185 = Companion_D0_.Create_DC2_(_332_v184)
			_333_v185 = _333_v185
		}
	}
	_dafny.Print((_62_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_63_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_63_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_63_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_63_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_63_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_63_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_64_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_65_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_72_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true).UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_74_v9).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_75_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_95_v23.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v24).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v24).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v24).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v24).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v25).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(176), _dafny.IntOfInt64(176), _dafny.IntOfInt64(176)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v26).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(176), _dafny.IntOfInt64(176), _dafny.IntOfInt64(176))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(176), _dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_99_v27).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_99_v27).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_99_v27).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_99_v27).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_177_v85, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_244_v122, _dafny.SeqOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_245_v123, _dafny.SeqOf(_dafny.IntOfInt64(623))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v124).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_247_v125).Equals(_dafny.SetOf(_dafny.IntOfInt64(617))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(623), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(623))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_248_v128).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), _dafny.IntOfInt64(176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_301_v157).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_302_v158, _dafny.SeqOf(_dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(true), _dafny.SeqOf(false), _dafny.SeqOf(true, false)), _dafny.SeqOf(_dafny.SeqOf(true, true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true), _dafny.SeqOf(false), _dafny.SeqOf(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v160).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_305_v161).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true).UpdateUnsafe(true, false), false).UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true).UpdateUnsafe(true, true), true).UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v162).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_307_v163).Dtor_cf15()).Equals(_dafny.SetOf(_dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v165).Dtor_cf11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v165).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v165).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_v166).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true).UpdateUnsafe(true, true), Companion_D2_.Create_DC6_(_dafny.IntOfInt64(623), _dafny.IntOfInt64(623), _dafny.IntOfInt64(623)))))
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
	Cf0 _dafny.Array
	Cf1 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Array, Cf1 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0, Cf1}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf2 bool
	Cf3 _dafny.Int
	Cf4 _dafny.Int
	Cf5 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf2 bool, Cf3 _dafny.Int, Cf4 _dafny.Int, Cf5 _dafny.Int) D0 {
	return D0{D0_DC1{Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf6 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf6 D0) D0 {
	return D0{D0_DC2{Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D0) Dtor_cf0() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf5
}

func (_this D0) Dtor_cf6() D0 {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf0 == data2.Cf0 && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
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

type D1_DC4 struct {
	Cf8 bool
	Cf9 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 bool, Cf9 _dafny.Int) D1 {
	return D1{D1_DC4{Cf8, Cf9}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf7 _dafny.Map
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 _dafny.Map) D1 {
	return D1{D1_DC3{Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, _dafny.Zero)
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf11 _dafny.Int
	Cf12 _dafny.Int
	Cf13 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Int, Cf12 _dafny.Int, Cf13 _dafny.Int) D2 {
	return D2{D2_DC6{Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf10 _dafny.MultiSet
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf10 _dafny.MultiSet) D2 {
	return D2{D2_DC5{Cf10}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) Dtor_cf10() _dafny.MultiSet {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
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

type D3_DC7 struct {
	Cf14 _dafny.MultiSet
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf14 _dafny.MultiSet) D3 {
	return D3{D3_DC7{Cf14}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_()
}

func (_this D3) Dtor_cf14() _dafny.MultiSet {
	return _this.Get_().(D3_DC7).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf14) + ")"
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
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf14.Equals(data2.Cf14)
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
	Cf16 D3
	Cf17 _dafny.Int
	Cf18 _dafny.CodePoint
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf16 D3, Cf17 _dafny.Int, Cf18 _dafny.CodePoint) D4 {
	return D4{D4_DC10{Cf16, Cf17, Cf18}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC11 struct {
	Cf19 _dafny.Int
	Cf20 bool
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf19 _dafny.Int, Cf20 bool) D4 {
	return D4{D4_DC11{Cf19, Cf20}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC9 struct {
	Cf15 _dafny.Set
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf15 _dafny.Set) D4 {
	return D4{D4_DC9{Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(Companion_D3_.Default(), _dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D4) Dtor_cf16() D3 {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) Dtor_cf18() _dafny.CodePoint {
	return _this.Get_().(D4_DC10).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC11).Cf20
}

func (_this D4) Dtor_cf15() _dafny.Set {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf15) + ")"
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
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16) && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18 == data2.Cf18
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf21 _dafny.Array
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf21 _dafny.Array) D5 {
	return D5{D5_DC12{Cf21}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf21() _dafny.Array {
	return _this.Get_().(D5_DC12).Cf21
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf21 == data2.Cf21
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
	Cf23 _dafny.CodePoint
	Cf24 _dafny.Int
	Cf25 _dafny.Map
	Cf26 _dafny.Int
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf23 _dafny.CodePoint, Cf24 _dafny.Int, Cf25 _dafny.Map, Cf26 _dafny.Int) D6 {
	return D6{D6_DC14{Cf23, Cf24, Cf25, Cf26}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC13 struct {
	Cf22 _dafny.Sequence
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf22 _dafny.Sequence) D6 {
	return D6{D6_DC13{Cf22}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC15 struct {
	Cf27 D6
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf27 D6) D6 {
	return D6{D6_DC15{Cf27}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(_dafny.CodePoint('D'), _dafny.Zero, _dafny.EmptyMap, _dafny.Zero)
}

func (_this D6) Dtor_cf23() _dafny.CodePoint {
	return _this.Get_().(D6_DC14).Cf23
}

func (_this D6) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Map {
	return _this.Get_().(D6_DC14).Cf25
}

func (_this D6) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf26
}

func (_this D6) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D6_DC13).Cf22
}

func (_this D6) Dtor_cf27() D6 {
	return _this.Get_().(D6_DC15).Cf27
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf27) + ")"
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
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Equals(data2.Cf25) && data1.Cf26.Cmp(data2.Cf26) == 0
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf22.Equals(data2.Cf22)
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf27.Equals(data2.Cf27)
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

// Definition of class GlobalState
type GlobalState struct {
	F4  _dafny.Int
	_f0 bool
	_f1 _dafny.Int
	_f2 _dafny.Int
	_f3 _dafny.Array
	_f5 _dafny.CodePoint
	_f6 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = _dafny.Zero
	_this._f0 = false
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.Zero
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Array, f4 _dafny.Int, f5 _dafny.CodePoint, f6 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
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
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Array {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() _dafny.CodePoint {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
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
func (_this *C0) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true, true, true))
	}
}
func (_this *C0) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-214)
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
