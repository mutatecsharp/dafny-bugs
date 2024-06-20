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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	return (Companion_D1_.Create_DC2_(_dafny.IntOfInt64(121))).Dtor_cf2()
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) bool {
	return (!(!(!((_dafny.MultiSetOf(_dafny.IntOfInt64(-776))).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(448))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(func() _dafny.Set {
			var _coll0 = _dafny.NewBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(686), _dafny.IntOfInt64(805))); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _1_v0 _dafny.Int
				_1_v0 = interface{}(_compr_0).(_dafny.Int)
				if ((_dafny.IntOfInt64(686)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(805)) < 0) {
					_coll0.Add((_1_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(323), true))).Cardinality())))
				}
			}
			return _coll0.ToSet()
		}()))).Cardinality()
	})))))))) && (!(!((func() bool {
		if false {
			return false
		}
		return false
	})())))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ifnkwt"), _dafny.UnicodeSeqOfUtf8Bytes("hammgsnta")), ((_dafny.Zero).Minus(_dafny.IntOfInt64(-352))).Cmp((_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality())) <= 0, (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(240))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(277)
	})))), false, true)
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	var _source0 D4 = Companion_D4_.Create_DC11_(_dafny.IntOfInt64(823))
	_ = _source0
	if _source0.Is_DC9() {
		var _3___mcc_h0 _dafny.Array = _source0.Get_().(D4_DC9).Cf10
		_ = _3___mcc_h0
		var _4___mcc_h1 *C0 = _source0.Get_().(D4_DC9).Cf11
		_ = _4___mcc_h1
		var _5___mcc_h2 bool = _source0.Get_().(D4_DC9).Cf12
		_ = _5___mcc_h2
		var _6_cf12 bool = _5___mcc_h2
		_ = _6_cf12
		var _7_cf11 *C0 = _4___mcc_h1
		_ = _7_cf11
		var _8_cf10 _dafny.Array = _3___mcc_h0
		_ = _8_cf10
		return _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aqlkl")).Cardinality()))
	} else if _source0.Is_DC10() {
		var _9___mcc_h3 _dafny.Int = _source0.Get_().(D4_DC10).Cf13
		_ = _9___mcc_h3
		var _10___mcc_h4 _dafny.CodePoint = _source0.Get_().(D4_DC10).Cf14
		_ = _10___mcc_h4
		var _11___mcc_h5 _dafny.Int = _source0.Get_().(D4_DC10).Cf15
		_ = _11___mcc_h5
		var _12___mcc_h6 bool = _source0.Get_().(D4_DC10).Cf16
		_ = _12___mcc_h6
		var _13___mcc_h7 _dafny.Sequence = _source0.Get_().(D4_DC10).Cf17
		_ = _13___mcc_h7
		var _14_cf17 _dafny.Sequence = _13___mcc_h7
		_ = _14_cf17
		var _15_cf16 bool = _12___mcc_h6
		_ = _15_cf16
		var _16_cf15 _dafny.Int = _11___mcc_h5
		_ = _16_cf15
		var _17_cf14 _dafny.CodePoint = _10___mcc_h4
		_ = _17_cf14
		var _18_cf13 _dafny.Int = _9___mcc_h3
		_ = _18_cf13
		return (_dafny.SetOf(_16_cf15)).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(611))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}((func(_19_cf14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_20_i0 _dafny.Int) _dafny.CodePoint {
				return _19_cf14
			}
		})(_17_cf14)))).Cardinality()), _16_cf15))
	} else if _source0.Is_DC11() {
		var _21___mcc_h8 _dafny.Int = _source0.Get_().(D4_DC11).Cf18
		_ = _21___mcc_h8
		var _22_cf18 _dafny.Int = _21___mcc_h8
		_ = _22_cf18
		return _dafny.SetOf(_22_cf18, _22_cf18, _22_cf18, (_dafny.Zero).Minus(_22_cf18), (_dafny.Zero).Minus((_dafny.Zero).Minus(_22_cf18)))
	} else if _source0.Is_DC8() {
		var _23___mcc_h9 bool = _source0.Get_().(D4_DC8).Cf9
		_ = _23___mcc_h9
		var _24_cf9 bool = _23___mcc_h9
		_ = _24_cf9
		return _dafny.SetOf(_dafny.IntOfInt64(-944), _dafny.IntOfInt64(396), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_cf9, _dafny.IntOfInt64(228))).Cardinality())
	} else {
		var _25___mcc_h10 D4 = _source0.Get_().(D4_DC12).Cf19
		_ = _25___mcc_h10
		var _26_cf19 D4 = _25___mcc_h10
		_ = _26_cf19
		return func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(73), _dafny.IntOfInt64(-142))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _27_v0 _dafny.Int
				_27_v0 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(73)).Cmp(_27_v0) <= 0) && ((_27_v0).Cmp(_dafny.IntOfInt64(-142)) < 0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_27_v0, _dafny.IntOfInt64(972)))
				}
			}
			return _coll1.ToSet()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Map {
	if !((_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, true, true, true, true)).Cardinality())), _dafny.IntOfInt64(207))).Cardinality())).Cmp(_dafny.IntOfInt64(350)) >= 0) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(75), _dafny.IntOfInt64(606))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _28_v0 _dafny.Int
				_28_v0 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(75)).Cmp(_28_v0) <= 0) && ((_28_v0).Cmp(_dafny.IntOfInt64(606)) < 0) {
					_coll2.Add(Companion_Default___.SafeDivisionInt(_28_v0, _dafny.IntOfInt64(436)), true)
				}
			}
			return _coll2.ToMap()
		}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bkytbt")).Cardinality()), false)))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-230), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-856))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(514), false)).Cardinality())).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _29_v1 _dafny.Int
				_29_v1 = interface{}(_compr_3).(_dafny.Int)
				if (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-230), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-856))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(514), false)).Cardinality())).Contains(_29_v1) {
					_coll3.Add(Companion_Default___.SafeModuloInt(_29_v1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(340))).Cardinality()))).Cardinality()))), false)
				}
			}
			return _coll3.ToMap()
		}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qevan")).Cardinality()), false)))
	}
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 bool, p2 D1, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(165))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_30_i0 _dafny.Int) _dafny.CodePoint {
		return (func() _dafny.CodePoint {
			if !(false) {
				return _dafny.CodePoint('y')
			}
			return _dafny.CodePoint('g')
		})()
	}))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(914), (_dafny.Zero).Minus(_dafny.IntOfInt64(-417))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(400)), _dafny.SeqOf(_dafny.IntOfInt64(553))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC6_(_dafny.IntOfInt64(522))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if false {
		return _dafny.CodePoint('u')
	} else {
		return _dafny.CodePoint('q')
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("gywvd"))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(817), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(583))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_31_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(147)
	}))).Cardinality()))).Cardinality())).Cardinality()))).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(-387), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('x'))).Cardinality(), _dafny.IntOfInt64(868), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oendfkfg")).Cardinality()))).Cardinality()))).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(559))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_32_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("htm")).Cardinality()))).Cardinality())
	})), _dafny.SeqOf(_dafny.IntOfInt64(766))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true)), _dafny.SeqOf(!(!(true))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (((Companion_D8_.Create_DC19_(_dafny.SetOf(_dafny.SeqOf(true, false, false, true), _dafny.SeqOf(true)))).Dtor_cf26()).Union(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), true)).Keys().Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _33_v0 _dafny.Sequence
			_33_v0 = interface{}(_compr_4).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), true)).Contains(_33_v0) {
				_coll4.Add(_33_v0)
			}
		}
		return _coll4.ToSet()
	}())).Intersection(_dafny.SetOf(_dafny.SeqOf(false, true, true, true, !(true)), _dafny.SeqOf(true, false), _dafny.SeqOf(true, true, false)))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false)))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) {
	var _34_v0 bool
	_ = _34_v0
	_34_v0 = false
	var _35_v1 *C0
	_ = _35_v1
	var _nw0 *C0 = New_C0_()
	_ = _nw0
	_nw0.Ctor__(_34_v0)
	_35_v1 = _nw0
	var _36_v2 _dafny.Array
	_ = _36_v2
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_0
	var _nw1 _dafny.Array
	_ = _nw1
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw1 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.MultiSet = func(_37_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(_dafny.IntOfInt64(497), (_dafny.SetOf(_dafny.IntOfInt64(-320), _dafny.IntOfInt64(-68), _dafny.IntOfInt64(69))).Cardinality(), _dafny.IntOfInt64(-332))
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw1 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw1).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw1).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_36_v2 = _nw1
	var _38_v3 _dafny.Map
	_ = _38_v3
	_38_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v1, _36_v2)
	_38_v3 = (_38_v3).Update(_35_v1, _36_v2)
	var _39_v4 _dafny.Int
	_ = _39_v4
	_39_v4 = _dafny.IntOfInt64(570)
	var _40_v5 _dafny.CodePoint
	_ = _40_v5
	_40_v5 = _dafny.CodePoint('h')
	var _pat_let_tv0 = _35_v1
	_ = _pat_let_tv0
	var _pat_let_tv1 = _35_v1
	_ = _pat_let_tv1
	if func(_source1 D2) bool {
		if _source1.Is_DC4() {
			var _41___mcc_h0 _dafny.Int = _source1.Get_().(D2_DC4).Cf4
			_ = _41___mcc_h0
			var _42___mcc_h1 _dafny.Int = _source1.Get_().(D2_DC4).Cf5
			_ = _42___mcc_h1
			var _43_cf5 _dafny.Int = _42___mcc_h1
			_ = _43_cf5
			var _44_cf4 _dafny.Int = _41___mcc_h0
			_ = _44_cf4
			return _pat_let_tv0.F15
		} else {
			var _45___mcc_h2 _dafny.Sequence = _source1.Get_().(D2_DC3).Cf3
			_ = _45___mcc_h2
			var _46_cf3 _dafny.Sequence = _45___mcc_h2
			_ = _46_cf3
			return !(_pat_let_tv1.F15)
		}
	}(Companion_Default___.Fm10(_39_v4, _40_v5, _34_v0, globalState)) {
		var _47_v6 _dafny.Array
		_ = _47_v6
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_1
		var _nw2 _dafny.Array
		_ = _nw2
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw2 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Sequence = func(_48_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("i")
			}
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw2 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw2).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw2).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_47_v6 = _nw2
		_47_v6 = _47_v6
		var _49_v7 _dafny.Sequence
		_ = _49_v7
		_49_v7 = _dafny.UnicodeSeqOfUtf8Bytes("rxcmd")
		var _50_v8 _dafny.Array
		_ = _50_v8
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_2
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = func(_51_i3 _dafny.Int) bool {
				return true
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw3 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw3).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw3).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_50_v8 = _nw3
		var _52_v9 D4
		_ = _52_v9
		_52_v9 = Companion_D4_.Create_DC9_(_50_v8, _35_v1, _34_v0)
		var _53_v10 _dafny.Sequence
		_ = _53_v10
		_53_v10 = _dafny.SeqOf(_35_v1.F15, _35_v1.F15, _35_v1.F15)
		var _54_v11 _dafny.MultiSet
		_ = _54_v11
		_54_v11 = _dafny.MultiSetOf(_35_v1.F15)
		var _55_v12 _dafny.Set
		_ = _55_v12
		_55_v12 = _dafny.SetOf(_35_v1, _35_v1)
		var _56_v13 _dafny.Set
		_ = _56_v13
		_56_v13 = _dafny.SetOf(_35_v1.F15)
		var _pat_let_tv2 = _35_v1
		_ = _pat_let_tv2
		var _pat_let_tv3 = _50_v8
		_ = _pat_let_tv3
		var _57_v14 _dafny.Array
		_ = _57_v14
		var _nwElement0_0 bool = _dafny.Companion_Sequence_.Equal(_49_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(391))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_58_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})))
		_ = _nwElement0_0
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(28))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_0, 0)
		(_nw4).ArraySet1(_35_v1.F15, 1)
		(_nw4).ArraySet1(_35_v1.F15, 2)
		(_nw4).ArraySet1((func(_pat_let0_0 D4) D4 {
			return func(_59_dt__update__tmp_h0 D4) D4 {
				return func(_pat_let1_0 *C0) D4 {
					return func(_60_dt__update_hcf11_h0 *C0) D4 {
						return func(_pat_let2_0 _dafny.Array) D4 {
							return func(_61_dt__update_hcf10_h0 _dafny.Array) D4 {
								return Companion_D4_.Create_DC9_(_61_dt__update_hcf10_h0, _60_dt__update_hcf11_h0, (_59_dt__update__tmp_h0).Dtor_cf12())
							}(_pat_let2_0)
						}(_pat_let_tv3)
					}(_pat_let1_0)
				}(_pat_let_tv2)
			}(_pat_let0_0)
		}(_52_v9)).Dtor_cf12(), 3)
		(_nw4).ArraySet1(((_dafny.Zero).Minus((_dafny.SetOf(_34_v0, _35_v1.F15, _35_v1.F15, _35_v1.F15, _35_v1.F15)).Cardinality())).Cmp(_dafny.IntOfUint32((_49_v7).Cardinality())) == 0, 4)
		(_nw4).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("difl"), _40_v5), 5)
		(_nw4).ArraySet1(true, 6)
		(_nw4).ArraySet1(_35_v1.F15, 7)
		(_nw4).ArraySet1(_35_v1.F15, 8)
		(_nw4).ArraySet1((_35_v1.F15) && ((_53_v10).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_34_v0, _39_v4, _35_v1.F15, globalState), _dafny.IntOfUint32((_53_v10).Cardinality()))).Uint32()).(bool)), 9)
		(_nw4).ArraySet1((_54_v11).IsSubsetOf((_54_v11).Update(_34_v0, Companion_Default___.Abs(_39_v4))), 10)
		(_nw4).ArraySet1(_34_v0, 11)
		(_nw4).ArraySet1(_35_v1.F15, 12)
		(_nw4).ArraySet1(_35_v1.F15, 13)
		(_nw4).ArraySet1((!(_34_v0)) || (true), 14)
		(_nw4).ArraySet1(Companion_Default___.Fm2(globalState), 15)
		(_nw4).ArraySet1(_35_v1.F15, 16)
		(_nw4).ArraySet1(_34_v0, 17)
		(_nw4).ArraySet1((_55_v12).IsDisjointFrom(_55_v12), 18)
		(_nw4).ArraySet1(_34_v0, 19)
		(_nw4).ArraySet1((_56_v13).Equals(_56_v13), 20)
		(_nw4).ArraySet1(!(_35_v1.F15) || (_35_v1.F15), 21)
		(_nw4).ArraySet1(_35_v1.F15, 22)
		(_nw4).ArraySet1(_35_v1.F15, 23)
		(_nw4).ArraySet1(_35_v1.F15, 24)
		(_nw4).ArraySet1(_34_v0, 25)
		(_nw4).ArraySet1(_35_v1.F15, 26)
		(_nw4).ArraySet1(_35_v1.F15, 27)
		_57_v14 = _nw4
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_50_v8), 0))
		_ = _index0
		(_50_v8).ArraySet1(_34_v0, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_50_v8), 0))
		_ = _index1
		(_50_v8).ArraySet1((func() bool {
			if !((_dafny.IntOfInt64(332)).Cmp(_39_v4) == 0) {
				return true
			}
			return (_35_v1.F15) || (_34_v0)
		})(), (_index1).Int())
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_50_v8), 0))
		_ = _index2
		(_50_v8).ArraySet1(Companion_Default___.Fm2(globalState), (_index2).Int())
		_39_v4 = _39_v4
		var _62_v15 _dafny.Array
		_ = _62_v15
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw5
		_62_v15 = _nw5
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_62_v15), 0))
		_ = _index3
		(_62_v15).ArraySet1(_39_v4, (_index3).Int())
		var _63_v16 _dafny.Sequence
		_ = _63_v16
		_63_v16 = _dafny.SeqOf(_35_v1)
		var _64_v17 _dafny.MultiSet
		_ = _64_v17
		_64_v17 = _dafny.MultiSetOf(_63_v16, _63_v16, _63_v16)
		var _65_v18 _dafny.Sequence
		_ = _65_v18
		_65_v18 = _dafny.SeqOf(_63_v16)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_62_v15), 0))
		_ = _index4
		var _rhs0 _dafny.Array = _62_v15
		_ = _rhs0
		var _rhs1 _dafny.Int = _39_v4
		_ = _rhs1
		var _rhs2 bool = ((_64_v17).Intersection(_dafny.MultiSetFromSeq(_65_v18))).IsSubsetOf((_64_v17).Intersection(_64_v17))
		_ = _rhs2
		var _rhs3 _dafny.Int = _39_v4
		_ = _rhs3
		var _lhs0 _dafny.Array = _62_v15
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_62_v15), 0))
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		var _lhs3 *GlobalState = globalState
		_ = _lhs3
		_62_v15 = _rhs0
		(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
		_lhs2.F8 = _rhs2
		_lhs3.F0 = _rhs3
	} else {
		var _66_v19 _dafny.Map
		_ = _66_v19
		_66_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_34_v0, Companion_Default___.Fm2(globalState))
		var _67_v20 _dafny.Sequence
		_ = _67_v20
		_67_v20 = _dafny.SeqOf(_66_v19)
		var _68_v22 D1
		_ = _68_v22
		_68_v22 = Companion_D1_.Create_DC2_(_39_v4)
		(globalState).F1 = ((func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_67_v20).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _69_v21 _dafny.Map
				_69_v21 = interface{}(_compr_5).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_67_v20, _69_v21) {
					_coll5.Add(_69_v21)
				}
			}
			return _coll5.ToSet()
		}()).Cardinality()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm6(_39_v4, _35_v1.F15, _68_v22, false, globalState)).Cardinality()))) <= 0
		var _70_v23 _dafny.MultiSet
		_ = _70_v23
		_70_v23 = _dafny.MultiSetOf(false)
		var _71_v24 _dafny.Sequence
		_ = _71_v24
		_71_v24 = _dafny.UnicodeSeqOfUtf8Bytes("kv")
		(globalState).F10 = ((func() _dafny.Int {
			if (_70_v23).Contains(_34_v0) {
				return (_70_v23).Multiplicity(_34_v0)
			}
			return _dafny.IntOfUint32((_71_v24).Cardinality())
		})()).Cmp(_dafny.IntOfInt64(717)) == 0
		var _72_v25 _dafny.Map
		_ = _72_v25
		_72_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v5, _35_v1)
		var _73_v26 _dafny.Map
		_ = _73_v26
		_73_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v25, _34_v0)
		var _74_v27 D5
		_ = _74_v27
		_74_v27 = Companion_D5_.Create_DC13_(_73_v26)
		_73_v26 = (_74_v27).Dtor_cf20()
		_35_v1 = _35_v1
		_66_v19 = (_66_v19).Update(true, _35_v1.F15)
	}
	var _75_i4 _dafny.Int
	_ = _75_i4
	_75_i4 = _dafny.Zero
	{
		for !(_34_v0) {
			{
				if (_75_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_75_i4 = (_75_i4).Plus(_dafny.One)
				var _76_v28 _dafny.Array
				_ = _76_v28
				var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw6
				_76_v28 = _nw6
				var _77_v29 _dafny.Sequence
				_ = _77_v29
				_77_v29 = _dafny.UnicodeSeqOfUtf8Bytes("brhkdmqxe")
				var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_76_v28), 0))
				_ = _index5
				(_76_v28).ArraySet1(_77_v29, (_index5).Int())
				var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_76_v28), 0))
				_ = _index6
				(_76_v28).ArraySet1(_77_v29, (_index6).Int())
				var _78_v30 _dafny.Map
				_ = _78_v30
				_78_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_34_v0, _40_v5)
				var _79_v31 _dafny.Array
				_ = _79_v31
				var _nwElement0_1 _dafny.CodePoint = _40_v5
				_ = _nwElement0_1
				var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(22))
				_ = _nw7
				(_nw7).ArraySet1CodePoint(_nwElement0_1, 0)
				(_nw7).ArraySet1CodePoint(_40_v5, 1)
				(_nw7).ArraySet1CodePoint(_40_v5, 2)
				(_nw7).ArraySet1CodePoint(_40_v5, 3)
				(_nw7).ArraySet1CodePoint(Companion_Default___.Fm9(_dafny.IntOfInt64(398), (_76_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_76_v28), 0))).Int()).(_dafny.Sequence), _39_v4, globalState), 4)
				(_nw7).ArraySet1CodePoint(_40_v5, 5)
				(_nw7).ArraySet1CodePoint(_40_v5, 6)
				(_nw7).ArraySet1CodePoint(_40_v5, 7)
				(_nw7).ArraySet1CodePoint(_40_v5, 8)
				(_nw7).ArraySet1CodePoint(_40_v5, 9)
				(_nw7).ArraySet1CodePoint(_40_v5, 10)
				(_nw7).ArraySet1CodePoint(_40_v5, 11)
				(_nw7).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _34_v0 {
						return _dafny.CodePoint('s')
					}
					return _dafny.CodePoint('l')
				})(), 12)
				(_nw7).ArraySet1CodePoint(_40_v5, 13)
				(_nw7).ArraySet1CodePoint(_40_v5, 14)
				(_nw7).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _34_v0 {
						return _dafny.CodePoint('v')
					}
					return _40_v5
				})(), 15)
				(_nw7).ArraySet1CodePoint(_dafny.CodePoint('r'), 16)
				(_nw7).ArraySet1CodePoint(_dafny.CodePoint('o'), 17)
				(_nw7).ArraySet1CodePoint(_40_v5, 18)
				(_nw7).ArraySet1CodePoint((_77_v29).Select((Companion_Default___.SafeIndex(_39_v4, _dafny.IntOfUint32((_77_v29).Cardinality()))).Uint32()).(_dafny.CodePoint), 19)
				(_nw7).ArraySet1CodePoint(_40_v5, 20)
				(_nw7).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_78_v30).Contains(_34_v0) {
						return (_78_v30).Get(_34_v0).(_dafny.CodePoint)
					}
					return _40_v5
				})(), 21)
				_79_v31 = _nw7
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_79_v31), 0))
				_ = _index7
				(_79_v31).ArraySet1CodePoint(_40_v5, (_index7).Int())
				var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_79_v31), 0))
				_ = _index8
				(_79_v31).ArraySet1CodePoint(_40_v5, (_index8).Int())
				var _80_v32 *C0
				_ = _80_v32
				var _nw8 *C0 = New_C0_()
				_ = _nw8
				_nw8.Ctor__(_34_v0)
				_80_v32 = _nw8
				var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_79_v31), 0))
				_ = _index9
				(_79_v31).ArraySet1CodePoint(_40_v5, (_index9).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	if _34_v0 {
		var _81_v33 _dafny.Array
		_ = _81_v33
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_3
		var _nw9 _dafny.Array
		_ = _nw9
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw9 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_82_v1 *C0) func(_dafny.Int) bool {
				return func(_83_i5 _dafny.Int) bool {
					return _82_v1.F15
				}
			})(_35_v1)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw9 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw9).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw9).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_81_v33 = _nw9
		var _84_v34 _dafny.Map
		_ = _84_v34
		_84_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v4, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_35_v1.F15, _34_v0, _35_v1.F15), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_39_v4), _dafny.IntOfUint32((_dafny.SeqOf(_35_v1.F15, _34_v0, _35_v1.F15)).Cardinality()))).Uint32(), _34_v0))
		var _85_v35 _dafny.Sequence
		_ = _85_v35
		_85_v35 = _dafny.SeqOf(_34_v0, _35_v1.F15)
		var _86_v36 _dafny.Set
		_ = _86_v36
		_86_v36 = _dafny.SetOf(_85_v35)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_81_v33), 0))
		_ = _index10
		(_81_v33).ArraySet1(!(_86_v36).Contains((func() _dafny.Sequence {
			if (_84_v34).Contains((_dafny.MultiSetOf(!(_35_v1.F15), _34_v0)).Cardinality()) {
				return (_84_v34).Get((_dafny.MultiSetOf(!(_35_v1.F15), _34_v0)).Cardinality()).(_dafny.Sequence)
			}
			return _85_v35
		})()), (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_81_v33), 0))
		_ = _index11
		(_81_v33).ArraySet1(_34_v0, (_index11).Int())
		var _87_v37 _dafny.Sequence
		_ = _87_v37
		_87_v37 = _dafny.SeqOf(_39_v4)
		var _88_v38 _dafny.MultiSet
		_ = _88_v38
		_88_v38 = _dafny.MultiSetOf(_87_v37, _87_v37)
		var _89_v39 _dafny.Sequence
		_ = _89_v39
		_89_v39 = _dafny.SeqOf(_87_v37)
		var _90_v40 _dafny.Map
		_ = _90_v40
		_90_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v1.F15, _39_v4)
		var _91_v41 _dafny.Map
		_ = _91_v41
		_91_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v4, _39_v4)
		var _92_v42 _dafny.Array
		_ = _92_v42
		var _nwElement0_2 _dafny.MultiSet = _88_v38
		_ = _nwElement0_2
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(13))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_2, 0)
		(_nw10).ArraySet1(_88_v38, 1)
		(_nw10).ArraySet1(_dafny.MultiSetFromSeq(_89_v39), 2)
		(_nw10).ArraySet1(_88_v38, 3)
		(_nw10).ArraySet1(_88_v38, 4)
		(_nw10).ArraySet1(_88_v38, 5)
		(_nw10).ArraySet1(_dafny.MultiSetFromSeq(_89_v39), 6)
		(_nw10).ArraySet1(Companion_Default___.Fm11(_39_v4, _39_v4, _34_v0, _34_v0, globalState), 7)
		(_nw10).ArraySet1(_dafny.MultiSetOf(_dafny.SeqOf((_90_v40).Cardinality(), (_91_v41).Cardinality()), _87_v37, _dafny.SeqOf(_39_v4), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_39_v4), (Companion_Default___.SafeIndex(_39_v4, _dafny.IntOfUint32((_dafny.SeqOf(_39_v4)).Cardinality()))).Uint32(), Companion_Default___.Fm0((_81_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_81_v33), 0))).Int()).(bool), _39_v4, _35_v1.F15, globalState)), _87_v37), 8)
		(_nw10).ArraySet1((_88_v38).Difference(_88_v38), 9)
		(_nw10).ArraySet1(_88_v38, 10)
		(_nw10).ArraySet1(_dafny.MultiSetOf(_87_v37), 11)
		(_nw10).ArraySet1((func() _dafny.MultiSet {
			if _35_v1.F15 {
				return _88_v38
			}
			return _88_v38
		})(), 12)
		_92_v42 = _nw10
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_92_v42), 0))
		_ = _index12
		(_92_v42).ArraySet1(Companion_Default___.Fm11(_39_v4, _39_v4, !(_35_v1.F15), !(false), globalState), (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_92_v42), 0))
		_ = _index13
		(_92_v42).ArraySet1((_88_v38).Difference((_88_v38).Union(_88_v38)), (_index13).Int())
		(globalState).F4 = _39_v4
		(globalState).F1 = (_81_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_81_v33), 0))).Int()).(bool)
		(globalState).F0 = _39_v4
	} else {
		var _93_v43 _dafny.Sequence
		_ = _93_v43
		_93_v43 = _dafny.SeqOf(_39_v4, _39_v4)
		var _94_v44 _dafny.Map
		_ = _94_v44
		_94_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v5, _93_v43))
		var _95_v45 _dafny.Map
		_ = _95_v45
		_95_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_93_v43).Cardinality()), _39_v4)
		var _96_v46 _dafny.Sequence
		_ = _96_v46
		_96_v46 = _dafny.SeqOf(_35_v1.F15)
		_94_v44 = (_94_v44).Update(_39_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v5, _dafny.SeqOf((_95_v45).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_96_v46).Cardinality()))), _39_v4, _39_v4, _39_v4)))
		var _97_v47 _dafny.Array
		_ = _97_v47
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw11
		_97_v47 = _nw11
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))
		_ = _index14
		(_97_v47).ArraySet1(false, (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))
		_ = _index15
		(_97_v47).ArraySet1(!(true) || (false), (_index15).Int())
		var _98_v48 _dafny.Array
		_ = _98_v48
		var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw12
		_98_v48 = _nw12
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_98_v48), 0))
		_ = _index16
		(_98_v48).ArraySet1(Companion_Default___.Fm0((_97_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))).Int()).(bool), _dafny.IntOfInt64(604), (_97_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))).Int()).(bool), globalState), (_index16).Int())
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_98_v48), 0))
		_ = _index17
		(_98_v48).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_39_v4), _39_v4), (_index17).Int())
		var _99_v49 _dafny.Map
		_ = _99_v49
		_99_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_97_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))).Int()).(bool), _35_v1.F15)
		(globalState).F0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_96_v46, (Companion_Default___.SafeIndex((func() _dafny.Int {
			if (func() bool {
				if (_99_v49).Contains(_34_v0) {
					return (_99_v49).Get(_34_v0).(bool)
				}
				return _35_v1.F15
			})() {
				return _39_v4
			}
			return (_dafny.Zero).Minus(_39_v4)
		})(), _dafny.IntOfUint32((_96_v46).Cardinality()))).Uint32(), _35_v1.F15)).Cardinality())
		if _34_v0 {
			var _rhs4 _dafny.CodePoint = _40_v5
			_ = _rhs4
			var _rhs5 _dafny.Map = _95_v45
			_ = _rhs5
			_40_v5 = _rhs4
			_95_v45 = _rhs5
			var _100_v50 _dafny.Set
			_ = _100_v50
			_100_v50 = _dafny.SetOf(_39_v4)
			var _101_v51 _dafny.Map
			_ = _101_v51
			_101_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12((_100_v50).Cardinality(), globalState), _dafny.SetOf((_97_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))).Int()).(bool)))
			var _102_v52 _dafny.Set
			_ = _102_v52
			_102_v52 = _dafny.SetOf(true)
			_101_v51 = (_101_v51).Update(_96_v46, _102_v52)
			var _103_v53 *C0
			_ = _103_v53
			var _nw13 *C0 = New_C0_()
			_ = _nw13
			_nw13.Ctor__((_dafny.IntOfInt64(367)).Cmp(_dafny.IntOfInt64(643)) == 0)
			_103_v53 = _nw13
			_96_v46 = _dafny.Companion_Sequence_.Concatenate(_96_v46, _96_v46)
			(globalState).F11 = _97_v47
		} else {
			var _104_v54 _dafny.Map
			_ = _104_v54
			_104_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_98_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_98_v48), 0))).Int()).(_dafny.Int), _35_v1)
			var _105_v55 _dafny.Set
			_ = _105_v55
			_105_v55 = _dafny.SetOf(_35_v1.F15, !(_35_v1.F15))
			var _106_v56 _dafny.Map
			_ = _106_v56
			_106_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v55, (_98_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_98_v48), 0))).Int()).(_dafny.Int))
			var _107_v57 _dafny.Set
			_ = _107_v57
			_107_v57 = _dafny.SetOf(_35_v1, (func() *C0 {
				if (_104_v54).Contains((_106_v56).Cardinality()) {
					return (_104_v54).Get((_106_v56).Cardinality()).(*C0)
				}
				return _35_v1
			})(), _35_v1)
			var _108_v58 _dafny.MultiSet
			_ = _108_v58
			_108_v58 = _dafny.MultiSetOf((_97_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))).Int()).(bool), (_97_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_97_v47), 0))).Int()).(bool))
			var _rhs6 _dafny.Int = (_dafny.Zero).Minus((_98_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_98_v48), 0))).Int()).(_dafny.Int))
			_ = _rhs6
			var _rhs7 *C0 = (func() *C0 {
				if (_104_v54).Contains(Companion_Default___.SafeDivisionInt(_39_v4, (_108_v58).Cardinality())) {
					return (_104_v54).Get(Companion_Default___.SafeDivisionInt(_39_v4, (_108_v58).Cardinality())).(*C0)
				}
				return _35_v1
			})()
			_ = _rhs7
			var _rhs8 _dafny.Int = _39_v4
			_ = _rhs8
			var _rhs9 _dafny.Set = ((_107_v57).Union(_107_v57)).Union(_107_v57)
			_ = _rhs9
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			_lhs4.F4 = _rhs6
			_35_v1 = _rhs7
			_lhs5.F6 = _rhs8
			_107_v57 = _rhs9
			var _109_v59 _dafny.Map
			_ = _109_v59
			_109_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v47, (_98_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_98_v48), 0))).Int()).(_dafny.Int))
			_109_v59 = (_109_v59).Update(_97_v47, _dafny.IntOfInt64(282))
			var _110_v60 _dafny.Map
			_ = _110_v60
			_110_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_39_v4).Cmp(_39_v4) <= 0, _97_v47)
			_110_v60 = (_110_v60).Update(Companion_Default___.Fm2(globalState), _97_v47)
			(globalState).F6 = Companion_Default___.Fm0(_35_v1.F15, (_93_v43).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-243), _dafny.IntOfUint32((_93_v43).Cardinality()))).Uint32()).(_dafny.Int), _35_v1.F15, globalState)
			var _111_v61 _dafny.Sequence
			_ = _111_v61
			_111_v61 = _dafny.UnicodeSeqOfUtf8Bytes("mkynuw")
			_95_v45 = (_95_v45).Update((_98_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_98_v48), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_111_v61, _111_v61)).Cardinality()))
		}
	}
	if _35_v1.F15 {
		var _112_v62 _dafny.Array
		_ = _112_v62
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_4
		var _nw14 _dafny.Array
		_ = _nw14
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw14 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Int = (func(_113_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_114_i6 _dafny.Int) _dafny.Int {
					return (_114_i6).Minus(_113_v4)
				}
			})(_39_v4)
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
		_112_v62 = _nw14
		var _115_v63 _dafny.MultiSet
		_ = _115_v63
		_115_v63 = _dafny.MultiSetOf(_40_v5)
		var _116_v64 _dafny.Set
		_ = _116_v64
		_116_v64 = _dafny.SetOf(_39_v4)
		var _117_v65 _dafny.Sequence
		_ = _117_v65
		_117_v65 = _dafny.UnicodeSeqOfUtf8Bytes("f")
		var _118_v66 D1
		_ = _118_v66
		_118_v66 = Companion_D1_.Create_DC2_(_39_v4)
		var _119_v67 _dafny.Map
		_ = _119_v67
		_119_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_34_v0, _35_v1.F15)
		var _120_v68 _dafny.Array
		_ = _120_v68
		var _nwElement0_3 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_117_v65, _117_v65)
		_ = _nwElement0_3
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(25))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_3, 0)
		(_nw15).ArraySet1(Companion_Default___.Fm6(_39_v4, _35_v1.F15, _118_v66, true, globalState), 1)
		(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kov"), 2)
		(_nw15).ArraySet1(_117_v65, 3)
		(_nw15).ArraySet1(_117_v65, 4)
		(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ndxuxbq"), 5)
		(_nw15).ArraySet1(_117_v65, 6)
		(_nw15).ArraySet1(_117_v65, 7)
		(_nw15).ArraySet1(_117_v65, 8)
		(_nw15).ArraySet1(_117_v65, 9)
		(_nw15).ArraySet1(_117_v65, 10)
		(_nw15).ArraySet1(Companion_Default___.Fm6(_39_v4, _35_v1.F15, Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_117_v65).Cardinality())), _35_v1.F15, globalState), 11)
		(_nw15).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(372))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_121_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_122_i7 _dafny.Int) _dafny.CodePoint {
				return _121_v5
			}
		})(_40_v5))), 12)
		(_nw15).ArraySet1(_117_v65, 13)
		(_nw15).ArraySet1(_117_v65, 14)
		(_nw15).ArraySet1(_117_v65, 15)
		(_nw15).ArraySet1((func() _dafny.Sequence {
			if !(_35_v1.F15) {
				return _117_v65
			}
			return _117_v65
		})(), 16)
		(_nw15).ArraySet1(_117_v65, 17)
		(_nw15).ArraySet1(_117_v65, 18)
		(_nw15).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(123))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_123_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_124_i8 _dafny.Int) _dafny.CodePoint {
				return _123_v5
			}
		})(_40_v5))), (Companion_Default___.SafeIndex(_39_v4, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(123))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_125_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_126_i8 _dafny.Int) _dafny.CodePoint {
				return _125_v5
			}
		})(_40_v5)))).Cardinality()))).Uint32(), _40_v5), 19)
		(_nw15).ArraySet1(_117_v65, 20)
		(_nw15).ArraySet1(_117_v65, 21)
		(_nw15).ArraySet1(Companion_Default___.Fm6(_39_v4, Companion_Default___.Fm2(globalState), Companion_D1_.Create_DC2_((_119_v67).Cardinality()), _35_v1.F15, globalState), 22)
		(_nw15).ArraySet1(Companion_Default___.Fm6(_39_v4, true, _118_v66, _35_v1.F15, globalState), 23)
		(_nw15).ArraySet1(_117_v65, 24)
		_120_v68 = _nw15
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_120_v68), 0))
		_ = _index18
		(_120_v68).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ju"), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))
		_ = _index19
		(_112_v62).ArraySet1(_39_v4, (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_120_v68), 0))
		_ = _index20
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))
		_ = _index21
		var _rhs10 _dafny.Array = _112_v62
		_ = _rhs10
		var _rhs11 _dafny.MultiSet = _115_v63
		_ = _rhs11
		var _rhs12 _dafny.Set = (_dafny.SetOf(_39_v4)).Union(_116_v64)
		_ = _rhs12
		var _rhs13 _dafny.Sequence = _117_v65
		_ = _rhs13
		var _rhs14 _dafny.Int = (_39_v4).Plus(_39_v4)
		_ = _rhs14
		var _lhs6 _dafny.Array = _120_v68
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_120_v68), 0))
		_ = _lhs7
		var _lhs8 _dafny.Array = _112_v62
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))
		_ = _lhs9
		_112_v62 = _rhs10
		_115_v63 = _rhs11
		_116_v64 = _rhs12
		(_lhs6).ArraySet1(_rhs13, (_lhs7).Int())
		(_lhs8).ArraySet1(_rhs14, (_lhs9).Int())
		var _127_v69 _dafny.Array
		_ = _127_v69
		var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw16
		_127_v69 = _nw16
		var _128_v70 _dafny.Sequence
		_ = _128_v70
		_128_v70 = _dafny.SeqOf((_112_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0(_35_v1.F15, _39_v4, !(_34_v0), globalState), _39_v4, (_112_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))).Int()).(_dafny.Int), (_112_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))).Int()).(_dafny.Int))
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_127_v69), 0))
		_ = _index22
		(_127_v69).ArraySet1(_128_v70, (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_127_v69), 0))
		_ = _index23
		(_127_v69).ArraySet1(_128_v70, (_index23).Int())
		var _129_v71 _dafny.Array
		_ = _129_v71
		var _nw17 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw17
		_129_v71 = _nw17
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))
		_ = _index24
		(_129_v71).ArraySet1(!(_35_v1.F15), (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))
		_ = _index25
		(_129_v71).ArraySet1(_34_v0, (_index25).Int())
		var _130_v72 _dafny.Map
		_ = _130_v72
		_130_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_120_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_120_v68), 0))).Int()).(_dafny.Sequence)).Cardinality()), _35_v1.F15)
		var _131_v73 _dafny.Sequence
		_ = _131_v73
		_131_v73 = _dafny.SeqOf(_112_v62, _112_v62)
		var _132_v74 D6
		_ = _132_v74
		_132_v74 = Companion_D6_.Create_DC15_(_112_v62)
		var _133_v75 _dafny.Array
		_ = _133_v75
		var _nwElement0_4 _dafny.Array = _112_v62
		_ = _nwElement0_4
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(16))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_4, 0)
		(_nw18).ArraySet1(_112_v62, 1)
		(_nw18).ArraySet1(_112_v62, 2)
		(_nw18).ArraySet1((func() _dafny.Array {
			if (func() bool {
				if (_130_v72).Contains(_39_v4) {
					return (_130_v72).Get(_39_v4).(bool)
				}
				return _35_v1.F15
			})() {
				return _112_v62
			}
			return (_131_v73).Select((Companion_Default___.SafeIndex((_112_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_131_v73).Cardinality()))).Uint32()).(_dafny.Array)
		})(), 3)
		(_nw18).ArraySet1(_112_v62, 4)
		(_nw18).ArraySet1(_112_v62, 5)
		(_nw18).ArraySet1(_112_v62, 6)
		(_nw18).ArraySet1(_112_v62, 7)
		(_nw18).ArraySet1(_112_v62, 8)
		(_nw18).ArraySet1((func() _dafny.Array {
			if (_129_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))).Int()).(bool) {
				return _112_v62
			}
			return (_132_v74).Dtor_cf21()
		})(), 9)
		(_nw18).ArraySet1(_112_v62, 10)
		(_nw18).ArraySet1(_112_v62, 11)
		(_nw18).ArraySet1(_112_v62, 12)
		(_nw18).ArraySet1(_112_v62, 13)
		(_nw18).ArraySet1(_112_v62, 14)
		(_nw18).ArraySet1(_112_v62, 15)
		_133_v75 = _nw18
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_133_v75), 0))
		_ = _index26
		(_133_v75).ArraySet1(_112_v62, (_index26).Int())
		var _134_v77 _dafny.Array
		_ = _134_v77
		var _len0_5 _dafny.Int = _dafny.One
		_ = _len0_5
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Set = (func(_135_v1 *C0, _136_v62 _dafny.Array, _137_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.Set {
				return func(_138_i9 _dafny.Int) _dafny.Set {
					return func() _dafny.Set {
						var _coll6 = _dafny.NewBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(_135_v1.F15, (Companion_D4_.Create_DC10_((_136_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_136_v62), 0))).Int()).(_dafny.Int), _137_v5, (_136_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_136_v62), 0))).Int()).(_dafny.Int), _135_v1.F15, _dafny.SeqOf(Companion_D3_.Create_DC6_((_136_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_136_v62), 0))).Int()).(_dafny.Int))))).Dtor_cf16()))).Elements()); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _139_v76 _dafny.Sequence
							_139_v76 = interface{}(_compr_6).(_dafny.Sequence)
							if (_dafny.MultiSetOf(_dafny.SeqOf(_135_v1.F15, (Companion_D4_.Create_DC10_((_136_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_136_v62), 0))).Int()).(_dafny.Int), _137_v5, (_136_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_136_v62), 0))).Int()).(_dafny.Int), _135_v1.F15, _dafny.SeqOf(Companion_D3_.Create_DC6_((_136_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_136_v62), 0))).Int()).(_dafny.Int))))).Dtor_cf16()))).Contains(_139_v76) {
								_coll6.Add(_139_v76)
							}
						}
						return _coll6.ToSet()
					}()
				}
			})(_35_v1, _112_v62, _40_v5)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw19 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw19).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw19).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_134_v77 = _nw19
		var _140_v78 _dafny.Sequence
		_ = _140_v78
		_140_v78 = _dafny.SeqOf(true, _35_v1.F15, (_129_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))).Int()).(bool))
		var _141_v79 _dafny.Set
		_ = _141_v79
		_141_v79 = _dafny.SetOf(_140_v78, _140_v78, _140_v78)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_134_v77), 0))
		_ = _index27
		(_134_v77).ArraySet1(_141_v79, (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_133_v75), 0))
		_ = _index28
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_134_v77), 0))
		_ = _index29
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))
		_ = _index30
		var _rhs15 _dafny.Array = (_132_v74).Dtor_cf21()
		_ = _rhs15
		var _rhs16 bool = _35_v1.F15
		_ = _rhs16
		var _rhs17 _dafny.Int = (func() _dafny.Int {
			if _35_v1.F15 {
				return (_112_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))).Int()).(_dafny.Int)
			}
			return (_112_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))).Int()).(_dafny.Int)
		})()
		_ = _rhs17
		var _rhs18 _dafny.Set = (Companion_Default___.Fm13((_112_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_112_v62), 0))).Int()).(_dafny.Int), globalState)).Difference(_141_v79)
		_ = _rhs18
		var _rhs19 bool = (_129_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))).Int()).(bool)
		_ = _rhs19
		var _lhs10 _dafny.Array = _133_v75
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_133_v75), 0))
		_ = _lhs11
		var _lhs12 *GlobalState = globalState
		_ = _lhs12
		var _lhs13 *GlobalState = globalState
		_ = _lhs13
		var _lhs14 _dafny.Array = _134_v77
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_134_v77), 0))
		_ = _lhs15
		var _lhs16 _dafny.Array = _129_v71
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))
		_ = _lhs17
		(_lhs10).ArraySet1(_rhs15, (_lhs11).Int())
		_lhs12.F10 = _rhs16
		_lhs13.F0 = _rhs17
		(_lhs14).ArraySet1(_rhs18, (_lhs15).Int())
		(_lhs16).ArraySet1(_rhs19, (_lhs17).Int())
		(globalState).F10 = (_129_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_129_v71), 0))).Int()).(bool)
	} else {
		(_35_v1).F15 = _35_v1.F15
		if _35_v1.F15 {
			var _142_v80 _dafny.Array
			_ = _142_v80
			var _nwElement0_5 _dafny.Int = _39_v4
			_ = _nwElement0_5
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.One)
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_5, 0)
			_142_v80 = _nw20
			var _143_v81 _dafny.Map
			_ = _143_v81
			_143_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v4, _142_v80)
			var _144_v82 _dafny.Array
			_ = _144_v82
			var _nwElement0_6 bool = _35_v1.F15
			_ = _nwElement0_6
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(25))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_6, 0)
			(_nw21).ArraySet1(_35_v1.F15, 1)
			(_nw21).ArraySet1(_35_v1.F15, 2)
			(_nw21).ArraySet1(false, 3)
			(_nw21).ArraySet1(_35_v1.F15, 4)
			(_nw21).ArraySet1(_35_v1.F15, 5)
			(_nw21).ArraySet1(_35_v1.F15, 6)
			(_nw21).ArraySet1(_35_v1.F15, 7)
			(_nw21).ArraySet1(_35_v1.F15, 8)
			(_nw21).ArraySet1(_34_v0, 9)
			(_nw21).ArraySet1(_35_v1.F15, 10)
			(_nw21).ArraySet1(_34_v0, 11)
			(_nw21).ArraySet1(_35_v1.F15, 12)
			(_nw21).ArraySet1(_35_v1.F15, 13)
			(_nw21).ArraySet1(true, 14)
			(_nw21).ArraySet1(_34_v0, 15)
			(_nw21).ArraySet1(_35_v1.F15, 16)
			(_nw21).ArraySet1(_34_v0, 17)
			(_nw21).ArraySet1(_35_v1.F15, 18)
			(_nw21).ArraySet1(_34_v0, 19)
			(_nw21).ArraySet1(_35_v1.F15, 20)
			(_nw21).ArraySet1(_35_v1.F15, 21)
			(_nw21).ArraySet1(_34_v0, 22)
			(_nw21).ArraySet1(_34_v0, 23)
			(_nw21).ArraySet1(false, 24)
			_144_v82 = _nw21
			var _145_v83 _dafny.Array
			_ = _145_v83
			_145_v83 = _144_v82
			var _146_v84 _dafny.Array
			_ = _146_v84
			var _nwElement0_7 _dafny.Array = _142_v80
			_ = _nwElement0_7
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(18))
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_7, 0)
			(_nw22).ArraySet1(_142_v80, 1)
			(_nw22).ArraySet1((func() _dafny.Array {
				if (_143_v81).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v83, _35_v1.F15)).Cardinality()) {
					return (_143_v81).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v83, _35_v1.F15)).Cardinality()).(_dafny.Array)
				}
				return _142_v80
			})(), 2)
			(_nw22).ArraySet1(_142_v80, 3)
			(_nw22).ArraySet1(_142_v80, 4)
			(_nw22).ArraySet1(_142_v80, 5)
			(_nw22).ArraySet1(_142_v80, 6)
			(_nw22).ArraySet1(_142_v80, 7)
			(_nw22).ArraySet1(_142_v80, 8)
			(_nw22).ArraySet1(_142_v80, 9)
			(_nw22).ArraySet1(_142_v80, 10)
			(_nw22).ArraySet1(_142_v80, 11)
			(_nw22).ArraySet1(_142_v80, 12)
			(_nw22).ArraySet1(_142_v80, 13)
			(_nw22).ArraySet1(_142_v80, 14)
			(_nw22).ArraySet1(_142_v80, 15)
			(_nw22).ArraySet1(_142_v80, 16)
			(_nw22).ArraySet1(_142_v80, 17)
			_146_v84 = _nw22
			var _147_v85 _dafny.Sequence
			_ = _147_v85
			_147_v85 = _dafny.SeqOf(_146_v84)
			var _148_v86 _dafny.Sequence
			_ = _148_v86
			_148_v86 = _dafny.SeqOf(_35_v1.F15, _35_v1.F15)
			var _149_v87 _dafny.Map
			_ = _149_v87
			_149_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_147_v85).Select((Companion_Default___.SafeIndex(_39_v4, _dafny.IntOfUint32((_147_v85).Cardinality()))).Uint32()).(_dafny.Array), _dafny.Companion_Sequence_.Concatenate(_148_v86, _148_v86))
			_149_v87 = (_149_v87).Update(_146_v84, _148_v86)
			_34_v0 = (_39_v4).Cmp(_39_v4) == 0
			var _150_v88 _dafny.Map
			_ = _150_v88
			_150_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v1.F15, _35_v1.F15)
			_150_v88 = (_150_v88).Update(_35_v1.F15, _35_v1.F15)
			var _151_v89 D3
			_ = _151_v89
			_151_v89 = Companion_D3_.Create_DC5_(_148_v86)
			(globalState).F0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_148_v86, Companion_D3_.Create_DC7_(_151_v89))).Cardinality()
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_144_v82), 0))
			_ = _index31
			(_144_v82).ArraySet1(_35_v1.F15, (_index31).Int())
			var _152_v90 D1
			_ = _152_v90
			_152_v90 = Companion_D1_.Create_DC2_(_39_v4)
			var _153_v91 _dafny.Map
			_ = _153_v91
			_153_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(249), _39_v4)
			var _154_v92 _dafny.Array
			_ = _154_v92
			var _nwElement0_8 D1 = _152_v90
			_ = _nwElement0_8
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(2))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_8, 0)
			(_nw23).ArraySet1(Companion_D1_.Create_DC2_((_153_v91).Cardinality()), 1)
			_154_v92 = _nw23
			var _155_v93 _dafny.MultiSet
			_ = _155_v93
			_155_v93 = _dafny.MultiSetOf(_154_v92)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_144_v82), 0))
			_ = _index32
			var _rhs20 _dafny.Int = _39_v4
			_ = _rhs20
			var _rhs21 bool = _35_v1.F15
			_ = _rhs21
			var _rhs22 bool = ((func() _dafny.MultiSet {
				if true {
					return _dafny.MultiSetOf(_154_v92)
				}
				return _155_v93
			})()).IsDisjointFrom(_155_v93)
			_ = _rhs22
			var _rhs23 _dafny.Array = _142_v80
			_ = _rhs23
			var _lhs18 *GlobalState = globalState
			_ = _lhs18
			var _lhs19 _dafny.Array = _144_v82
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_144_v82), 0))
			_ = _lhs20
			var _lhs21 *C0 = _35_v1
			_ = _lhs21
			_lhs18.F4 = _rhs20
			(_lhs19).ArraySet1(_rhs21, (_lhs20).Int())
			_lhs21.F15 = _rhs22
			_142_v80 = _rhs23
		} else {
			var _156_v94 _dafny.Map
			_ = _156_v94
			_156_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v1.F15, _35_v1)
			var _157_v95 _dafny.Array
			_ = _157_v95
			var _nwElement0_9 *C0 = (func() *C0 {
				if (_156_v94).Contains(!(true)) {
					return (_156_v94).Get(!(true)).(*C0)
				}
				return _35_v1
			})()
			_ = _nwElement0_9
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(14))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_9, 0)
			(_nw24).ArraySet1(_35_v1, 1)
			(_nw24).ArraySet1(_35_v1, 2)
			(_nw24).ArraySet1(_35_v1, 3)
			(_nw24).ArraySet1(_35_v1, 4)
			(_nw24).ArraySet1(_35_v1, 5)
			(_nw24).ArraySet1(_35_v1, 6)
			(_nw24).ArraySet1(_35_v1, 7)
			(_nw24).ArraySet1(_35_v1, 8)
			(_nw24).ArraySet1((func() *C0 {
				if (_156_v94).Contains(_35_v1.F15) {
					return (_156_v94).Get(_35_v1.F15).(*C0)
				}
				return _35_v1
			})(), 9)
			(_nw24).ArraySet1(_35_v1, 10)
			(_nw24).ArraySet1(_35_v1, 11)
			(_nw24).ArraySet1(_35_v1, 12)
			(_nw24).ArraySet1(_35_v1, 13)
			_157_v95 = _nw24
			var _158_v96 _dafny.Sequence
			_ = _158_v96
			_158_v96 = _dafny.SeqOf(_35_v1, _35_v1)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_157_v95), 0))
			_ = _index33
			(_157_v95).ArraySet1((_158_v96).Select((Companion_Default___.SafeIndex(_39_v4, _dafny.IntOfUint32((_158_v96).Cardinality()))).Uint32()).(*C0), (_index33).Int())
			var _159_v97 _dafny.Sequence
			_ = _159_v97
			_159_v97 = _dafny.UnicodeSeqOfUtf8Bytes("iaufjsuky")
			var _160_v98 _dafny.Sequence
			_ = _160_v98
			_160_v98 = _dafny.SeqOf(_159_v97, _159_v97)
			var _161_v99 D2
			_ = _161_v99
			_161_v99 = Companion_D2_.Create_DC3_(_159_v97)
			var _162_v100 _dafny.Array
			_ = _162_v100
			var _nwElement0_10 _dafny.Sequence = (_160_v98).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_159_v97).Cardinality()), _dafny.IntOfUint32((_160_v98).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _nwElement0_10
			var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(5))
			_ = _nw25
			(_nw25).ArraySet1(_nwElement0_10, 0)
			(_nw25).ArraySet1(_159_v97, 1)
			(_nw25).ArraySet1(_159_v97, 2)
			(_nw25).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_163_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_164_i10 _dafny.Int) _dafny.CodePoint {
					return _163_v5
				}
			})(_40_v5))), _159_v97), (Companion_Default___.SafeIndex(_39_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_165_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_166_i10 _dafny.Int) _dafny.CodePoint {
					return _165_v5
				}
			})(_40_v5))), _159_v97)).Cardinality()))).Uint32(), _40_v5), 3)
			(_nw25).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_161_v99).Dtor_cf3(), _dafny.UnicodeSeqOfUtf8Bytes("rtb")), 4)
			_162_v100 = _nw25
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_162_v100), 0))
			_ = _index34
			(_162_v100).ArraySet1(_159_v97, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_157_v95), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_162_v100), 0))
			_ = _index36
			var _rhs24 bool = _35_v1.F15
			_ = _rhs24
			var _rhs25 _dafny.Int = (Companion_Default___.Fm0(_35_v1.F15, _39_v4, _35_v1.F15, globalState)).Minus((_dafny.Zero).Minus(_39_v4))
			_ = _rhs25
			var _rhs26 *C0 = _35_v1
			_ = _rhs26
			var _rhs27 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mjfq"), _159_v97)
			_ = _rhs27
			var _lhs22 *C0 = _35_v1
			_ = _lhs22
			var _lhs23 *GlobalState = globalState
			_ = _lhs23
			var _lhs24 _dafny.Array = _157_v95
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_157_v95), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _162_v100
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_162_v100), 0))
			_ = _lhs27
			_lhs22.F15 = _rhs24
			_lhs23.F4 = _rhs25
			(_lhs24).ArraySet1(_rhs26, (_lhs25).Int())
			(_lhs26).ArraySet1(_rhs27, (_lhs27).Int())
			var _167_v101 _dafny.Sequence
			_ = _167_v101
			_167_v101 = _dafny.SeqOf(_35_v1.F15, _35_v1.F15, Companion_Default___.Fm2(globalState))
			_167_v101 = _167_v101
			var _rhs28 bool = ((_35_v1.F15) || (_34_v0)) || (_35_v1.F15)
			_ = _rhs28
			var _rhs29 _dafny.Int = (func() _dafny.Int {
				if _35_v1.F15 {
					return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_162_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_162_v100), 0))).Int()).(_dafny.Sequence)).Cardinality()), _39_v4)
				}
				return _39_v4
			})()
			_ = _rhs29
			var _lhs28 *GlobalState = globalState
			_ = _lhs28
			_lhs28.F1 = _rhs28
			_39_v4 = _rhs29
			var _168_v102 _dafny.Sequence
			_ = _168_v102
			_168_v102 = _dafny.SeqOf(_39_v4)
			var _169_v103 _dafny.Set
			_ = _169_v103
			_169_v103 = _dafny.SetOf(!(false) || (false), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-681))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_170_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_171_i11 _dafny.Int) _dafny.Int {
					return _170_v4
				}
			})(_39_v4))), _168_v102))
			var _172_v104 _dafny.Set
			_ = _172_v104
			_172_v104 = _dafny.SetOf(_39_v4, _dafny.IntOfInt64(741), Companion_Default___.SafeDivisionInt(_39_v4, _dafny.IntOfUint32(((_162_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_162_v100), 0))).Int()).(_dafny.Sequence)).Cardinality())))
			var _rhs30 _dafny.Int = (_dafny.Zero).Minus((_169_v103).Cardinality())
			_ = _rhs30
			var _rhs31 _dafny.Int = (_dafny.Zero).Minus(_39_v4)
			_ = _rhs31
			var _rhs32 _dafny.Int = (_172_v104).Cardinality()
			_ = _rhs32
			var _lhs29 *GlobalState = globalState
			_ = _lhs29
			var _lhs30 *GlobalState = globalState
			_ = _lhs30
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			_lhs29.F6 = _rhs30
			_lhs30.F0 = _rhs31
			_lhs31.F0 = _rhs32
			var _173_v105 _dafny.Array
			_ = _173_v105
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_6
			var _nw26 _dafny.Array
			_ = _nw26
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw26 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Map = (func(_174_v4 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_175_i12 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v4, _174_v4)
					}
				})(_39_v4)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw26 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw26).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw26).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_173_v105 = _nw26
			var _176_v106 _dafny.Map
			_ = _176_v106
			_176_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v1.F15, _35_v1.F15)
			var _177_v107 D1
			_ = _177_v107
			_177_v107 = Companion_D1_.Create_DC2_(_39_v4)
			var _rhs33 _dafny.Array = _173_v105
			_ = _rhs33
			var _rhs34 _dafny.Int = (Companion_Default___.Fm0(_35_v1.F15, (_dafny.Zero).Minus((_dafny.Zero).Minus(_39_v4)), _35_v1.F15, globalState)).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm6((_176_v106).Cardinality(), _35_v1.F15, _177_v107, _35_v1.F15, globalState)).Cardinality())))
			_ = _rhs34
			var _rhs35 _dafny.Int = _39_v4
			_ = _rhs35
			var _rhs36 _dafny.Int = _39_v4
			_ = _rhs36
			var _lhs32 *GlobalState = globalState
			_ = _lhs32
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			_173_v105 = _rhs33
			_lhs32.F0 = _rhs34
			_lhs33.F6 = _rhs35
			_lhs34.F4 = _rhs36
		}
		var _178_v108 _dafny.Array
		_ = _178_v108
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw27
		_178_v108 = _nw27
		var _179_v109 _dafny.Array
		_ = _179_v109
		_179_v109 = _178_v108
		var _source2 _dafny.Array = _179_v109
		_ = _source2
		var _180___mcc_h3 _dafny.Array = _source2
		_ = _180___mcc_h3
		var _181_cf0 _dafny.Array = _180___mcc_h3
		_ = _181_cf0
		var _182_v110 _dafny.Array
		_ = _182_v110
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
		_ = _nw28
		_182_v110 = _nw28
		var _183_v111 _dafny.Sequence
		_ = _183_v111
		_183_v111 = _dafny.UnicodeSeqOfUtf8Bytes("mftvud")
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_182_v110), 0))
		_ = _index37
		(_182_v110).ArraySet1((func() _dafny.Sequence {
			if _35_v1.F15 {
				return _183_v111
			}
			return _183_v111
		})(), (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_181_cf0), 0))
		_ = _index38
		(_181_cf0).ArraySet1(_35_v1.F15, (_index38).Int())
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_182_v110), 0))
		_ = _index39
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_181_cf0), 0))
		_ = _index40
		var _rhs37 _dafny.Sequence = _183_v111
		_ = _rhs37
		var _rhs38 _dafny.Int = _39_v4
		_ = _rhs38
		var _rhs39 _dafny.Int = Companion_Default___.SafeDivisionInt(_39_v4, Companion_Default___.SafeDivisionInt(_39_v4, _39_v4))
		_ = _rhs39
		var _rhs40 bool = !(!(_35_v1.F15))
		_ = _rhs40
		var _lhs35 _dafny.Array = _182_v110
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_182_v110), 0))
		_ = _lhs36
		var _lhs37 *GlobalState = globalState
		_ = _lhs37
		var _lhs38 _dafny.Array = _181_cf0
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_181_cf0), 0))
		_ = _lhs39
		(_lhs35).ArraySet1(_rhs37, (_lhs36).Int())
		_39_v4 = _rhs38
		_lhs37.F0 = _rhs39
		(_lhs38).ArraySet1(_rhs40, (_lhs39).Int())
		var _184_v112 _dafny.Map
		_ = _184_v112
		_184_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if !(_34_v0) {
				return _182_v110
			}
			return _182_v110
		})(), (_dafny.Zero).Minus(_39_v4))
		_184_v112 = (_184_v112).Update(_182_v110, Companion_Default___.Fm0(false, _39_v4, _34_v0, globalState))
		var _185_v113 *C0
		_ = _185_v113
		var _nw29 *C0 = New_C0_()
		_ = _nw29
		_nw29.Ctor__(!(!(_35_v1.F15) || (_35_v1.F15)))
		_185_v113 = _nw29
		var _186_v114 _dafny.Array
		_ = _186_v114
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
		_ = _nw30
		_186_v114 = _nw30
		var _187_v115 _dafny.Array
		_ = _187_v115
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw31
		_187_v115 = _nw31
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_186_v114), 0))
		_ = _index41
		(_186_v114).ArraySet1(_187_v115, (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_186_v114), 0))
		_ = _index42
		(_186_v114).ArraySet1(_187_v115, (_index42).Int())
		if _35_v1.F15 {
			_34_v0 = !(_35_v1.F15)
			(globalState).F4 = _39_v4
			var _188_v116 _dafny.Map
			_ = _188_v116
			_188_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v1.F15, _39_v4)
			var _189_v117 _dafny.Sequence
			_ = _189_v117
			_189_v117 = _dafny.UnicodeSeqOfUtf8Bytes("shhs")
			var _190_v118 D1
			_ = _190_v118
			_190_v118 = Companion_D1_.Create_DC1_(_39_v4)
			var _191_v119 _dafny.Set
			_ = _191_v119
			_191_v119 = _dafny.SetOf(_35_v1.F15, _35_v1.F15)
			var _192_v120 _dafny.Array
			_ = _192_v120
			var _nwElement0_11 _dafny.Int = _dafny.IntOfInt64(115)
			_ = _nwElement0_11
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(23))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_11, 0)
			(_nw32).ArraySet1((_188_v116).Cardinality(), 1)
			(_nw32).ArraySet1(_dafny.IntOfUint32((_189_v117).Cardinality()), 2)
			(_nw32).ArraySet1(_39_v4, 3)
			(_nw32).ArraySet1(_dafny.IntOfUint32((_189_v117).Cardinality()), 4)
			(_nw32).ArraySet1((_190_v118).Dtor_cf1(), 5)
			(_nw32).ArraySet1(_dafny.IntOfUint32((_189_v117).Cardinality()), 6)
			(_nw32).ArraySet1((_dafny.Zero).Minus(_39_v4), 7)
			(_nw32).ArraySet1(_39_v4, 8)
			(_nw32).ArraySet1(_39_v4, 9)
			(_nw32).ArraySet1(_39_v4, 10)
			(_nw32).ArraySet1(_39_v4, 11)
			(_nw32).ArraySet1((Companion_Default___.Fm4(_35_v1.F15, _39_v4, globalState)).Cardinality(), 12)
			(_nw32).ArraySet1((_191_v119).Cardinality(), 13)
			(_nw32).ArraySet1(_39_v4, 14)
			(_nw32).ArraySet1(_39_v4, 15)
			(_nw32).ArraySet1(Companion_Default___.Fm0(_35_v1.F15, _39_v4, _35_v1.F15, globalState), 16)
			(_nw32).ArraySet1(_39_v4, 17)
			(_nw32).ArraySet1(_39_v4, 18)
			(_nw32).ArraySet1((_dafny.Zero).Minus(_39_v4), 19)
			(_nw32).ArraySet1(_39_v4, 20)
			(_nw32).ArraySet1(_39_v4, 21)
			(_nw32).ArraySet1(_39_v4, 22)
			_192_v120 = _nw32
			var _193_v121 _dafny.Array
			_ = _193_v121
			var _nwElement0_12 _dafny.Array = _192_v120
			_ = _nwElement0_12
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(17))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_12, 0)
			(_nw33).ArraySet1(_192_v120, 1)
			(_nw33).ArraySet1(_192_v120, 2)
			(_nw33).ArraySet1(_192_v120, 3)
			(_nw33).ArraySet1(_192_v120, 4)
			(_nw33).ArraySet1(_192_v120, 5)
			(_nw33).ArraySet1(_192_v120, 6)
			(_nw33).ArraySet1(_192_v120, 7)
			(_nw33).ArraySet1(_192_v120, 8)
			(_nw33).ArraySet1(_192_v120, 9)
			(_nw33).ArraySet1(_192_v120, 10)
			(_nw33).ArraySet1(_192_v120, 11)
			(_nw33).ArraySet1(_192_v120, 12)
			(_nw33).ArraySet1(_192_v120, 13)
			(_nw33).ArraySet1(_192_v120, 14)
			(_nw33).ArraySet1(_192_v120, 15)
			(_nw33).ArraySet1(_192_v120, 16)
			_193_v121 = _nw33
			var _194_v122 _dafny.Sequence
			_ = _194_v122
			_194_v122 = _dafny.SeqOf(_193_v121, _193_v121, _193_v121)
			var _195_v123 D7
			_ = _195_v123
			_195_v123 = Companion_D7_.Create_DC17_(_193_v121)
			var _196_v124 _dafny.Array
			_ = _196_v124
			var _nwElement0_13 _dafny.Array = _193_v121
			_ = _nwElement0_13
			var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(16))
			_ = _nw34
			(_nw34).ArraySet1(_nwElement0_13, 0)
			(_nw34).ArraySet1(_193_v121, 1)
			(_nw34).ArraySet1((_194_v122).Select((Companion_Default___.SafeIndex(_39_v4, _dafny.IntOfUint32((_194_v122).Cardinality()))).Uint32()).(_dafny.Array), 2)
			(_nw34).ArraySet1(_193_v121, 3)
			(_nw34).ArraySet1(_193_v121, 4)
			(_nw34).ArraySet1(_193_v121, 5)
			(_nw34).ArraySet1(_193_v121, 6)
			(_nw34).ArraySet1(_193_v121, 7)
			(_nw34).ArraySet1(_193_v121, 8)
			(_nw34).ArraySet1(_193_v121, 9)
			(_nw34).ArraySet1(_193_v121, 10)
			(_nw34).ArraySet1(_193_v121, 11)
			(_nw34).ArraySet1(_193_v121, 12)
			(_nw34).ArraySet1((_195_v123).Dtor_cf25(), 13)
			(_nw34).ArraySet1(_193_v121, 14)
			(_nw34).ArraySet1(_193_v121, 15)
			_196_v124 = _nw34
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_196_v124), 0))
			_ = _index43
			(_196_v124).ArraySet1(_193_v121, (_index43).Int())
			var _197_v125 _dafny.Array
			_ = _197_v125
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw35
			_197_v125 = _nw35
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_196_v124), 0))
			_ = _index44
			var _rhs41 _dafny.CodePoint = _40_v5
			_ = _rhs41
			var _rhs42 _dafny.Array = _193_v121
			_ = _rhs42
			var _rhs43 _dafny.Array = _197_v125
			_ = _rhs43
			var _rhs44 _dafny.Int = _39_v4
			_ = _rhs44
			var _rhs45 _dafny.Sequence = (Companion_D2_.Create_DC3_(_189_v117)).Dtor_cf3()
			_ = _rhs45
			var _lhs40 _dafny.Array = _196_v124
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_196_v124), 0))
			_ = _lhs41
			var _lhs42 *GlobalState = globalState
			_ = _lhs42
			_40_v5 = _rhs41
			(_lhs40).ArraySet1(_rhs42, (_lhs41).Int())
			_197_v125 = _rhs43
			_lhs42.F6 = _rhs44
			_189_v117 = _rhs45
			var _198_v126 _dafny.Sequence
			_ = _198_v126
			_198_v126 = _dafny.SeqOf(_34_v0)
			var _199_v127 D3
			_ = _199_v127
			_199_v127 = Companion_D3_.Create_DC6_(_39_v4)
			var _200_v128 _dafny.Sequence
			_ = _200_v128
			_200_v128 = _dafny.SeqOf(_199_v127)
			var _201_v129 D4
			_ = _201_v129
			_201_v129 = Companion_D4_.Create_DC10_(_39_v4, _40_v5, _39_v4, _35_v1.F15, _200_v128)
			var _202_v130 D3
			_ = _202_v130
			_202_v130 = Companion_D3_.Create_DC7_(Companion_D3_.Create_DC6_(_39_v4))
			var _203_v131 D3
			_ = _203_v131
			_203_v131 = Companion_D3_.Create_DC7_(_202_v130)
			var _204_v132 D3
			_ = _204_v132
			_204_v132 = Companion_D3_.Create_DC7_(_202_v130)
			var _pat_let_tv4 = _202_v130
			_ = _pat_let_tv4
			var _205_v133 _dafny.Set
			_ = _205_v133
			_205_v133 = _dafny.SetOf(_204_v132, func(_pat_let3_0 D3) D3 {
				return func(_206_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let4_0 D3) D3 {
						return func(_207_dt__update_hcf8_h0 D3) D3 {
							return Companion_D3_.Create_DC7_(_207_dt__update_hcf8_h0)
						}(_pat_let4_0)
					}(_pat_let_tv4)
				}(_pat_let3_0)
			}(_204_v132), _204_v132, _204_v132)
			var _rhs46 bool = (_201_v129).Dtor_cf16()
			_ = _rhs46
			var _rhs47 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(45), (_39_v4).Minus(_dafny.IntOfUint32((_189_v117).Cardinality())))
			_ = _rhs47
			var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_189_v117, (Companion_Default___.SafeIndex(((_205_v133).Union(_205_v133)).Cardinality(), _dafny.IntOfUint32((_189_v117).Cardinality()))).Uint32(), _40_v5)
			_ = _rhs48
			var _rhs49 *C0 = _35_v1
			_ = _rhs49
			var _rhs50 _dafny.Sequence = _198_v126
			_ = _rhs50
			var _lhs43 *GlobalState = globalState
			_ = _lhs43
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			_lhs43.F8 = _rhs46
			_lhs44.F0 = _rhs47
			_189_v117 = _rhs48
			_35_v1 = _rhs49
			_198_v126 = _rhs50
			var _208_v134 D4
			_ = _208_v134
			_208_v134 = Companion_D4_.Create_DC8_(_34_v0)
			(globalState).F10 = (_208_v134).Dtor_cf9()
		} else {
			var _209_v135 *C0
			_ = _209_v135
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__(false)
			_209_v135 = _nw36
			var _210_v136 _dafny.Map
			_ = _210_v136
			_210_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(globalState), _35_v1.F15)
			var _211_v137 _dafny.Set
			_ = _211_v137
			_211_v137 = _dafny.SetOf(_40_v5)
			_210_v136 = (_210_v136).Update(_209_v135.F15, (_211_v137).IsProperSubsetOf(_211_v137))
			var _212_v138 _dafny.Map
			_ = _212_v138
			_212_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v4, _209_v135)
			var _213_v139 _dafny.Sequence
			_ = _213_v139
			_213_v139 = _dafny.UnicodeSeqOfUtf8Bytes("medr")
			var _214_v140 _dafny.Map
			_ = _214_v140
			_214_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if !(_35_v1.F15) {
					return _39_v4
				}
				return _39_v4
			})(), (func() *C0 {
				if (_212_v138).Contains(_dafny.IntOfUint32((_213_v139).Cardinality())) {
					return (_212_v138).Get(_dafny.IntOfUint32((_213_v139).Cardinality())).(*C0)
				}
				return _35_v1
			})())
			_212_v138 = (_212_v138).Update(_39_v4, _35_v1)
			var _215_v141 _dafny.Sequence
			_ = _215_v141
			_215_v141 = _dafny.SeqOf(_209_v135.F15)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_178_v108), 0))
			_ = _index45
			(_178_v108).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_215_v141, _215_v141), (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_178_v108), 0))
			_ = _index46
			var _rhs51 _dafny.Sequence = _213_v139
			_ = _rhs51
			var _rhs52 _dafny.Int = _dafny.IntOfInt64(959)
			_ = _rhs52
			var _rhs53 bool = _35_v1.F15
			_ = _rhs53
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			var _lhs46 _dafny.Array = _178_v108
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_178_v108), 0))
			_ = _lhs47
			_213_v139 = _rhs51
			_lhs45.F0 = _rhs52
			(_lhs46).ArraySet1(_rhs53, (_lhs47).Int())
			var _216_v142 _dafny.Set
			_ = _216_v142
			_216_v142 = _dafny.SetOf(_34_v0)
			var _217_v143 _dafny.Array
			_ = _217_v143
			var _nwElement0_14 _dafny.Int = _39_v4
			_ = _nwElement0_14
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(8))
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_14, 0)
			(_nw37).ArraySet1(Companion_Default___.Fm0(true, _39_v4, (_215_v141).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_209_v135.F15, _39_v4, (_215_v141).Select((Companion_Default___.SafeIndex((_216_v142).Cardinality(), _dafny.IntOfUint32((_215_v141).Cardinality()))).Uint32()).(bool), globalState), _dafny.IntOfUint32((_215_v141).Cardinality()))).Uint32()).(bool), globalState), 1)
			(_nw37).ArraySet1((_39_v4).Plus(_dafny.IntOfUint32((_213_v139).Cardinality())), 2)
			(_nw37).ArraySet1(_39_v4, 3)
			(_nw37).ArraySet1(Companion_Default___.SafeModuloInt(_39_v4, _39_v4), 4)
			(_nw37).ArraySet1(_39_v4, 5)
			(_nw37).ArraySet1((func() _dafny.Int {
				if _35_v1.F15 {
					return _39_v4
				}
				return _39_v4
			})(), 6)
			(_nw37).ArraySet1(_39_v4, 7)
			_217_v143 = _nw37
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_217_v143), 0))
			_ = _index47
			(_217_v143).ArraySet1(Companion_Default___.SafeDivisionInt(_39_v4, _39_v4), (_index47).Int())
			var _218_v144 _dafny.Map
			_ = _218_v144
			_218_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v4, _39_v4)
			var _219_v146 D4
			_ = _219_v146
			_219_v146 = Companion_D4_.Create_DC8_(true)
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_217_v143), 0))
			_ = _index48
			var _rhs54 bool = !(_218_v144).Equals(func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-450), _dafny.IntOfInt64(868))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _220_v145 _dafny.Int
					_220_v145 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(-450)).Cmp(_220_v145) <= 0) && ((_220_v145).Cmp(_dafny.IntOfInt64(868)) < 0) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_220_v145, _39_v4), (_dafny.MultiSetOf(_39_v4, _39_v4)).Cardinality())
					}
				}
				return _coll7.ToMap()
			}())
			_ = _rhs54
			var _rhs55 _dafny.Int = Companion_Default___.SafeDivisionInt(_39_v4, Companion_Default___.SafeModuloInt(_39_v4, _39_v4))
			_ = _rhs55
			var _rhs56 bool = (_35_v1.F15) || ((_219_v146).Dtor_cf9())
			_ = _rhs56
			var _rhs57 _dafny.Int = Companion_Default___.SafeModuloInt(_39_v4, (Companion_Default___.Fm14(globalState)).Cardinality())
			_ = _rhs57
			var _lhs48 *GlobalState = globalState
			_ = _lhs48
			var _lhs49 _dafny.Array = _217_v143
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_217_v143), 0))
			_ = _lhs50
			var _lhs51 *GlobalState = globalState
			_ = _lhs51
			var _lhs52 *GlobalState = globalState
			_ = _lhs52
			_lhs48.F1 = _rhs54
			(_lhs49).ArraySet1(_rhs55, (_lhs50).Int())
			_lhs51.F8 = _rhs56
			_lhs52.F6 = _rhs57
		}
		var _221_v147 _dafny.Array
		_ = _221_v147
		var _nw38 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
		_ = _nw38
		_221_v147 = _nw38
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_221_v147), 0))
		_ = _index49
		(_221_v147).ArraySet1(_35_v1, (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_221_v147), 0))
		_ = _index50
		(_221_v147).ArraySet1(_35_v1, (_index50).Int())
	}
	var _222_v148 _dafny.Sequence
	_ = _222_v148
	_222_v148 = _dafny.UnicodeSeqOfUtf8Bytes("usstt")
	_39_v4 = (_39_v4).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_222_v148, _222_v148)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _223_v0 _dafny.Array
	_ = _223_v0
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_7
	var _nw39 _dafny.Array
	_ = _nw39
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw39 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) bool = func(_224_i0 _dafny.Int) bool {
			return !(!(!(true)))
		}
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw39 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw39).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw39).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_223_v0 = _nw39
	var _225_v1 _dafny.Array
	_ = _225_v1
	_225_v1 = _223_v0
	var _226_v2 bool
	_ = _226_v2
	_226_v2 = true
	var _227_v3 _dafny.Map
	_ = _227_v3
	_227_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v2, true)
	var _228_globalState *GlobalState
	_ = _228_globalState
	var _nw40 *GlobalState = New_GlobalState_()
	_ = _nw40
	_nw40.Ctor__(_dafny.IntOfInt64(527), false, true, _dafny.IntOfInt64(871), _dafny.IntOfInt64(272), _dafny.IntOfInt64(844), _dafny.IntOfInt64(508), true, true, _223_v0, false, (_225_v1), (_227_v3).Merge(_227_v3), _dafny.IntOfInt64(357), false)
	_228_globalState = _nw40
	var _229_v4 _dafny.Int
	_ = _229_v4
	_229_v4 = _dafny.IntOfInt64(-669)
	(_228_globalState).F6 = Companion_Default___.SafeDivisionInt(_229_v4, _dafny.IntOfInt64(-539))
	var _230_v5 _dafny.Map
	_ = _230_v5
	_230_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v2, Companion_Default___.Fm0(true, _229_v4, _226_v2, _228_globalState))
	if (_226_v2) == ((_230_v5).Contains(_226_v2)) {
		_226_v2 = _226_v2
		if _226_v2 {
			var _231_v6 _dafny.Sequence
			_ = _231_v6
			_231_v6 = _dafny.SeqOf(_226_v2)
			_226_v2 = _dafny.Companion_Sequence_.Contains(_231_v6, _226_v2)
			var _232_v7 _dafny.Sequence
			_ = _232_v7
			_232_v7 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v0, _229_v4))
			var _233_v8 _dafny.Map
			_ = _233_v8
			_233_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_232_v7).Select((Companion_Default___.SafeIndex(_229_v4, _dafny.IntOfUint32((_232_v7).Cardinality()))).Uint32()).(_dafny.Map), false)
			var _234_v9 _dafny.CodePoint
			_ = _234_v9
			_234_v9 = _dafny.CodePoint('h')
			var _rhs58 bool = _226_v2
			_ = _rhs58
			var _rhs59 _dafny.Map = _233_v8
			_ = _rhs59
			var _rhs60 _dafny.CodePoint = _234_v9
			_ = _rhs60
			var _rhs61 _dafny.Int = _229_v4
			_ = _rhs61
			var _lhs53 *GlobalState = _228_globalState
			_ = _lhs53
			var _lhs54 *GlobalState = _228_globalState
			_ = _lhs54
			_lhs53.F1 = _rhs58
			_233_v8 = _rhs59
			_234_v9 = _rhs60
			_lhs54.F6 = _rhs61
			var _235_v10 *C0
			_ = _235_v10
			var _nw41 *C0 = New_C0_()
			_ = _nw41
			_nw41.Ctor__(Companion_Default___.Fm2(_228_globalState))
			_235_v10 = _nw41
			var _236_v11 _dafny.Sequence
			_ = _236_v11
			_236_v11 = _dafny.UnicodeSeqOfUtf8Bytes("m")
			var _237_v12 _dafny.Set
			_ = _237_v12
			_237_v12 = _dafny.SetOf(_dafny.IntOfInt64(893), (_dafny.IntOfUint32((_236_v11).Cardinality())).Plus(_229_v4), _229_v4, _229_v4, _229_v4)
			var _238_v13 _dafny.Map
			_ = _238_v13
			_238_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v10, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_239_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_240_i1 _dafny.Int) _dafny.CodePoint {
					return _239_v9
				}
			})(_234_v9)))).Cardinality()))
			_237_v12 = _dafny.SetOf((_238_v13).Cardinality(), _229_v4)
			(_228_globalState).F0 = _229_v4
		} else {
			var _241_v14 D1
			_ = _241_v14
			_241_v14 = Companion_D1_.Create_DC1_(_dafny.IntOfInt64(-329))
			(_228_globalState).F6 = (_241_v14).Dtor_cf1()
			var _242_v15 _dafny.Sequence
			_ = _242_v15
			_242_v15 = _dafny.UnicodeSeqOfUtf8Bytes("pwxho")
			var _243_v16 _dafny.Map
			_ = _243_v16
			_243_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v4, _226_v2)
			var _244_v17 D2
			_ = _244_v17
			_244_v17 = Companion_D2_.Create_DC3_(_dafny.Companion_Sequence_.Update(_242_v15, (Companion_Default___.SafeIndex(_229_v4, _dafny.IntOfUint32((_242_v15).Cardinality()))).Uint32(), _dafny.CodePoint('b')))
			var _245_v18 _dafny.Sequence
			_ = _245_v18
			_245_v18 = _dafny.SeqOf(_226_v2)
			var _rhs62 _dafny.Sequence = (_244_v17).Dtor_cf3()
			_ = _rhs62
			var _rhs63 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_245_v18, _245_v18)).Cardinality())
			_ = _rhs63
			var _rhs64 bool = (!(_226_v2)) == (!(_226_v2))
			_ = _rhs64
			var _rhs65 _dafny.Map = _243_v16
			_ = _rhs65
			var _lhs55 *GlobalState = _228_globalState
			_ = _lhs55
			var _lhs56 *GlobalState = _228_globalState
			_ = _lhs56
			_242_v15 = _rhs62
			_lhs55.F6 = _rhs63
			_lhs56.F8 = _rhs64
			_243_v16 = _rhs65
			var _246_v19 _dafny.Sequence
			_ = _246_v19
			_246_v19 = _dafny.SeqOf(_229_v4)
			var _247_v20 _dafny.CodePoint
			_ = _247_v20
			_247_v20 = _dafny.CodePoint('h')
			var _248_v21 _dafny.Map
			_ = _248_v21
			_248_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v20, _229_v4)
			var _249_v22 _dafny.Array
			_ = _249_v22
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw42
			_249_v22 = _nw42
			var _250_v23 _dafny.Map
			_ = _250_v23
			_250_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_249_v22, _229_v4)
			_246_v19 = _dafny.Companion_Sequence_.Concatenate(_246_v19, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(943))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_251_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_248_v21).Cardinality(), (_227_v3).Cardinality())).Cardinality()), (_250_v23).Cardinality()), _246_v19))
			var _252_v24 *C0
			_ = _252_v24
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__(!(_226_v2) || (_226_v2))
			_252_v24 = _nw43
			(_228_globalState).F8 = _252_v24.F15
		}
		if _226_v2 {
			var _253_v25 _dafny.Sequence
			_ = _253_v25
			_253_v25 = _dafny.UnicodeSeqOfUtf8Bytes("xygbbejn")
			var _254_v26 _dafny.Map
			_ = _254_v26
			_254_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v25, _226_v2)
			var _255_v27 *C0
			_ = _255_v27
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__((func() bool {
				if (_254_v26).Contains(_253_v25) {
					return (_254_v26).Get(_253_v25).(bool)
				}
				return _226_v2
			})())
			_255_v27 = _nw44
			var _256_v28 _dafny.Map
			_ = _256_v28
			_256_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v27, _255_v27.F15)
			_256_v28 = (_256_v28).Update(_255_v27, Companion_Default___.Fm2(_228_globalState))
			Companion_Default___.M0(_228_globalState)
			var _257_v29 _dafny.Set
			_ = _257_v29
			_257_v29 = _dafny.SetOf(!(_226_v2) || (!(_226_v2)))
			_257_v29 = (Companion_Default___.Fm3(_229_v4, _228_globalState)).Difference((_257_v29).Union(_dafny.SetOf(_255_v27.F15)))
			(_228_globalState).F8 = _255_v27.F15
			var _258_v30 _dafny.Array
			_ = _258_v30
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_8
			var _nw45 _dafny.Array
			_ = _nw45
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw45 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = (func(_259_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_260_i3 _dafny.Int) _dafny.Int {
						return (_260_i3).Plus(_259_v4)
					}
				})(_229_v4)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw45 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw45).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw45).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_258_v30 = _nw45
			var _261_v31 _dafny.MultiSet
			_ = _261_v31
			_261_v31 = _dafny.MultiSetOf(_258_v30)
			var _262_v32 _dafny.Sequence
			_ = _262_v32
			_262_v32 = _dafny.SeqOf(_261_v31, _261_v31)
			var _263_v33 _dafny.Sequence
			_ = _263_v33
			_263_v33 = _dafny.SeqOf(_229_v4, _229_v4, _229_v4, ((_262_v32).Select((Companion_Default___.SafeIndex(_229_v4, _dafny.IntOfUint32((_262_v32).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), _229_v4)
			_263_v33 = _263_v33
		} else {
			var _264_v34 _dafny.Array
			_ = _264_v34
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw46
			_264_v34 = _nw46
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_264_v34), 0))
			_ = _index51
			(_264_v34).ArraySet1(_223_v0, (_index51).Int())
			var _265_v35 _dafny.Set
			_ = _265_v35
			_265_v35 = _dafny.SetOf(true)
			var _266_v36 _dafny.Set
			_ = _266_v36
			_266_v36 = _dafny.SetOf(Companion_Default___.Fm0(_226_v2, _229_v4, _226_v2, _228_globalState), (_229_v4).Times((_265_v35).Cardinality()), _229_v4, (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_230_v5).Cardinality(), _229_v4)))
			var _267_v37 _dafny.Sequence
			_ = _267_v37
			_267_v37 = _dafny.UnicodeSeqOfUtf8Bytes("xd")
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_264_v34), 0))
			_ = _index52
			var _rhs66 _dafny.Array = _223_v0
			_ = _rhs66
			var _rhs67 _dafny.Int = _229_v4
			_ = _rhs67
			var _rhs68 _dafny.Set = (Companion_Default___.Fm4(_226_v2, _229_v4, _228_globalState)).Union((_266_v36).Intersection(_266_v36))
			_ = _rhs68
			var _rhs69 bool = _226_v2
			_ = _rhs69
			var _rhs70 bool = (func() bool {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_267_v37, _267_v37) {
					return _226_v2
				}
				return _226_v2
			})()
			_ = _rhs70
			var _lhs57 _dafny.Array = _264_v34
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_264_v34), 0))
			_ = _lhs58
			var _lhs59 *GlobalState = _228_globalState
			_ = _lhs59
			var _lhs60 *GlobalState = _228_globalState
			_ = _lhs60
			var _lhs61 *GlobalState = _228_globalState
			_ = _lhs61
			(_lhs57).ArraySet1(_rhs66, (_lhs58).Int())
			_lhs59.F0 = _rhs67
			_266_v36 = _rhs68
			_lhs60.F10 = _rhs69
			_lhs61.F8 = _rhs70
			var _268_v38 *C0
			_ = _268_v38
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__(_226_v2)
			_268_v38 = _nw47
			var _269_v39 _dafny.Array
			_ = _269_v39
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw48
			_269_v39 = _nw48
			var _270_v40 _dafny.Map
			_ = _270_v40
			_270_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v4, _229_v4)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_269_v39), 0))
			_ = _index53
			(_269_v39).ArraySet1((_270_v40).Cardinality(), (_index53).Int())
			var _271_v41 D2
			_ = _271_v41
			_271_v41 = Companion_D2_.Create_DC4_(_dafny.IntOfInt64(613), _229_v4)
			var _272_v42 _dafny.Sequence
			_ = _272_v42
			_272_v42 = _dafny.SeqOf(_226_v2, _226_v2, _268_v38.F15)
			var _273_v43 _dafny.Map
			_ = _273_v43
			_273_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_271_v41, _272_v42)
			var _274_v44 _dafny.MultiSet
			_ = _274_v44
			_274_v44 = _dafny.MultiSetOf(_dafny.SeqOf(_268_v38.F15, _268_v38.F15, _268_v38.F15), _272_v42)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_269_v39), 0))
			_ = _index54
			var _rhs71 *C0 = _268_v38
			_ = _rhs71
			var _rhs72 _dafny.Array = _269_v39
			_ = _rhs72
			var _rhs73 _dafny.Int = ((_274_v44).Cardinality()).Minus(_229_v4)
			_ = _rhs73
			var _rhs74 _dafny.Map = _273_v43
			_ = _rhs74
			var _lhs62 _dafny.Array = _269_v39
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_269_v39), 0))
			_ = _lhs63
			_268_v38 = _rhs71
			_269_v39 = _rhs72
			(_lhs62).ArraySet1(_rhs73, (_lhs63).Int())
			_273_v43 = _rhs74
			(_228_globalState).F1 = Companion_Default___.Fm2(_228_globalState)
			Companion_Default___.M0(_228_globalState)
			var _275_v45 _dafny.Array
			_ = _275_v45
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_9
			var _nw49 _dafny.Array
			_ = _nw49
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw49 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = (func(_276_v42 _dafny.Sequence, _277_v39 _dafny.Array, _278_v38 *C0) func(_dafny.Int) _dafny.Sequence {
					return func(_279_i4 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_276_v42, _276_v42), (Companion_Default___.SafeIndex((_277_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_277_v39), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_276_v42, _276_v42)).Cardinality()))).Uint32(), _278_v38.F15)
					}
				})(_272_v42, _269_v39, _268_v38)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw49 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw49).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw49).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_275_v45 = _nw49
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_275_v45), 0))
			_ = _index55
			(_275_v45).ArraySet1(_272_v42, (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_275_v45), 0))
			_ = _index56
			(_275_v45).ArraySet1(_272_v42, (_index56).Int())
		}
		var _280_v46 D2
		_ = _280_v46
		_280_v46 = Companion_D2_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_281_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		})))
		var _282_v47 _dafny.Sequence
		_ = _282_v47
		_282_v47 = _dafny.UnicodeSeqOfUtf8Bytes("yqrgy")
		var _pat_let_tv5 = _282_v47
		_ = _pat_let_tv5
		var _source3 D2 = func(_pat_let5_0 D2) D2 {
			return func(_283_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let6_0 _dafny.Sequence) D2 {
					return func(_284_dt__update_hcf3_h0 _dafny.Sequence) D2 {
						return Companion_D2_.Create_DC3_(_284_dt__update_hcf3_h0)
					}(_pat_let6_0)
				}(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("b"), _pat_let_tv5))
			}(_pat_let5_0)
		}(_280_v46)
		_ = _source3
		if _source3.Is_DC4() {
			var _285___mcc_h0 _dafny.Int = _source3.Get_().(D2_DC4).Cf4
			_ = _285___mcc_h0
			var _286___mcc_h1 _dafny.Int = _source3.Get_().(D2_DC4).Cf5
			_ = _286___mcc_h1
			var _287_cf5 _dafny.Int = _286___mcc_h1
			_ = _287_cf5
			var _288_cf4 _dafny.Int = _285___mcc_h0
			_ = _288_cf4
			var _289_v48 _dafny.CodePoint
			_ = _289_v48
			_289_v48 = _dafny.CodePoint('d')
			var _rhs75 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(654), _dafny.IntOfInt64(-318))
			_ = _rhs75
			var _rhs76 bool = _dafny.Companion_Sequence_.Contains(_282_v47, _289_v48)
			_ = _rhs76
			var _rhs77 bool = !((_229_v4).Cmp(_288_cf4) <= 0)
			_ = _rhs77
			var _rhs78 bool = _226_v2
			_ = _rhs78
			var _rhs79 _dafny.Int = _288_cf4
			_ = _rhs79
			var _lhs64 *GlobalState = _228_globalState
			_ = _lhs64
			var _lhs65 *GlobalState = _228_globalState
			_ = _lhs65
			var _lhs66 *GlobalState = _228_globalState
			_ = _lhs66
			var _lhs67 *GlobalState = _228_globalState
			_ = _lhs67
			var _lhs68 *GlobalState = _228_globalState
			_ = _lhs68
			_lhs64.F4 = _rhs75
			_lhs65.F1 = _rhs76
			_lhs66.F10 = _rhs77
			_lhs67.F1 = _rhs78
			_lhs68.F6 = _rhs79
			var _290_v49 _dafny.Sequence
			_ = _290_v49
			_290_v49 = _dafny.SeqOf(_223_v0)
			_290_v49 = _dafny.SeqOf(_223_v0, (_225_v1), _223_v0)
			(_228_globalState).F6 = _288_cf4
			var _291_v50 D1
			_ = _291_v50
			_291_v50 = Companion_D1_.Create_DC2_(_288_cf4)
			var _292_v51 _dafny.Map
			_ = _292_v51
			_292_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_282_v47).Cardinality()), _dafny.IntOfUint32((_282_v47).Cardinality()))
			var _293_v52 _dafny.Map
			_ = _293_v52
			_293_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_291_v50)).Cardinality(), _292_v51)
			var _294_v53 _dafny.Set
			_ = _294_v53
			_294_v53 = _dafny.SetOf(_226_v2, _226_v2)
			var _295_v54 _dafny.Array
			_ = _295_v54
			var _nwElement0_15 _dafny.Int = _287_cf5
			_ = _nwElement0_15
			var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(14))
			_ = _nw50
			(_nw50).ArraySet1(_nwElement0_15, 0)
			(_nw50).ArraySet1(_dafny.IntOfInt64(-32), 1)
			(_nw50).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(893), _287_cf5), 2)
			(_nw50).ArraySet1(_288_cf4, 3)
			(_nw50).ArraySet1(Companion_Default___.Fm0(_226_v2, _287_cf5, Companion_Default___.Fm2(_228_globalState), _228_globalState), 4)
			(_nw50).ArraySet1(_287_cf5, 5)
			(_nw50).ArraySet1(Companion_Default___.Fm0(_226_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_296_cf5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_297_i6 _dafny.Int) _dafny.Int {
					return _296_cf5
				}
			})(_287_cf5)))).Cardinality()), _226_v2, _228_globalState), 6)
			(_nw50).ArraySet1(_288_cf4, 7)
			(_nw50).ArraySet1(_229_v4, 8)
			(_nw50).ArraySet1(Companion_Default___.SafeDivisionInt(_229_v4, _229_v4), 9)
			(_nw50).ArraySet1((_293_v52).Cardinality(), 10)
			(_nw50).ArraySet1(_229_v4, 11)
			(_nw50).ArraySet1((_287_cf5).Plus((_294_v53).Cardinality()), 12)
			(_nw50).ArraySet1(_229_v4, 13)
			_295_v54 = _nw50
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_295_v54), 0))
			_ = _index57
			(_295_v54).ArraySet1(_229_v4, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_295_v54), 0))
			_ = _index58
			(_295_v54).ArraySet1((_288_cf4).Times(_288_cf4), (_index58).Int())
		} else {
			var _298___mcc_h2 _dafny.Sequence = _source3.Get_().(D2_DC3).Cf3
			_ = _298___mcc_h2
			var _299_cf3 _dafny.Sequence = _298___mcc_h2
			_ = _299_cf3
			(_228_globalState).F4 = _229_v4
			var _300_v55 _dafny.Map
			_ = _300_v55
			_300_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_226_v2, _229_v4, _226_v2, _228_globalState), _226_v2)
			var _301_v56 _dafny.Map
			_ = _301_v56
			_301_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v2, _300_v55)
			var _302_v57 _dafny.Array
			_ = _302_v57
			var _nwElement0_16 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_228_globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v4, !(_226_v2)))).Merge(_301_v56)
			_ = _nwElement0_16
			var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(17))
			_ = _nw51
			(_nw51).ArraySet1(_nwElement0_16, 0)
			(_nw51).ArraySet1(_301_v56, 1)
			(_nw51).ArraySet1(_301_v56, 2)
			(_nw51).ArraySet1(_301_v56, 3)
			(_nw51).ArraySet1(_301_v56, 4)
			(_nw51).ArraySet1(_301_v56, 5)
			(_nw51).ArraySet1((_301_v56).Merge(_301_v56), 6)
			(_nw51).ArraySet1(_301_v56, 7)
			(_nw51).ArraySet1(_301_v56, 8)
			(_nw51).ArraySet1(_301_v56, 9)
			(_nw51).ArraySet1(_301_v56, 10)
			(_nw51).ArraySet1(Companion_Default___.Fm5(_228_globalState), 11)
			(_nw51).ArraySet1((_301_v56).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v2, _300_v55)), 12)
			(_nw51).ArraySet1(_301_v56, 13)
			(_nw51).ArraySet1((Companion_Default___.Fm5(_228_globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v2, _300_v55)), 14)
			(_nw51).ArraySet1(_301_v56, 15)
			(_nw51).ArraySet1(_301_v56, 16)
			_302_v57 = _nw51
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_302_v57), 0))
			_ = _index59
			(_302_v57).ArraySet1(_301_v56, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_302_v57), 0))
			_ = _index60
			(_302_v57).ArraySet1(((_301_v56).Merge(_301_v56)).Merge(_301_v56), (_index60).Int())
			(_228_globalState).F0 = _dafny.IntOfInt64(-334)
			var _303_v58 *C0
			_ = _303_v58
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__(_226_v2)
			_303_v58 = _nw52
		}
		var _304_v59 D1
		_ = _304_v59
		_304_v59 = Companion_D1_.Create_DC2_(_229_v4)
		(_228_globalState).F8 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_282_v47, Companion_Default___.Fm6(_229_v4, _226_v2, _304_v59, _226_v2, _228_globalState)), _282_v47)
	} else {
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_223_v0), 0))
		_ = _index61
		(_223_v0).ArraySet1(_226_v2, (_index61).Int())
		var _305_v60 _dafny.Sequence
		_ = _305_v60
		_305_v60 = _dafny.SeqOf(_226_v2)
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_223_v0), 0))
		_ = _index62
		(_223_v0).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
			if _226_v2 {
				return _305_v60
			}
			return _dafny.SeqOf(Companion_Default___.Fm2(_228_globalState), _226_v2)
		})(), _305_v60), (_index62).Int())
		(_228_globalState).F6 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-760), (_dafny.Zero).Minus(_229_v4))).Times((func() _dafny.Int {
			if (_223_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_223_v0), 0))).Int()).(bool) {
				return _229_v4
			}
			return _229_v4
		})())
		var _306_v61 _dafny.Sequence
		_ = _306_v61
		_306_v61 = _dafny.UnicodeSeqOfUtf8Bytes("dqluto")
		_306_v61 = _306_v61
		var _307_v62 _dafny.Sequence
		_ = _307_v62
		_307_v62 = _dafny.SeqOf(_229_v4)
		var _308_v63 _dafny.Map
		_ = _308_v63
		_308_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_307_v62).Select((Companion_Default___.SafeIndex(_229_v4, _dafny.IntOfUint32((_307_v62).Cardinality()))).Uint32()).(_dafny.Int), (_223_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_223_v0), 0))).Int()).(bool))
		var _309_v64 D3
		_ = _309_v64
		_309_v64 = Companion_D3_.Create_DC5_(_305_v60)
		_308_v63 = (_308_v63).Update((_dafny.Zero).Minus((_229_v4).Plus((_dafny.Zero).Minus(_229_v4))), _dafny.Companion_Sequence_.IsPrefixOf(_305_v60, (_309_v64).Dtor_cf6()))
		(_228_globalState).F4 = (_229_v4).Minus(_229_v4)
	}
	(_228_globalState).F1 = _226_v2
	var _310_i7 _dafny.Int
	_ = _310_i7
	_310_i7 = _dafny.Zero
	{
		for (_229_v4).Cmp(_dafny.IntOfInt64(439)) != 0 {
			{
				if (_310_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_310_i7 = (_310_i7).Plus(_dafny.One)
				var _311_v65 _dafny.Sequence
				_ = _311_v65
				_311_v65 = _dafny.UnicodeSeqOfUtf8Bytes("l")
				(_228_globalState).F8 = !_dafny.Companion_Sequence_.Equal(_311_v65, _dafny.UnicodeSeqOfUtf8Bytes("avcf"))
				var _312_v66 *C0
				_ = _312_v66
				var _nw53 *C0 = New_C0_()
				_ = _nw53
				_nw53.Ctor__((_226_v2) || (_226_v2))
				_312_v66 = _nw53
				_312_v66 = _312_v66
				_227_v3 = (_227_v3).Update(_226_v2, _312_v66.F15)
				var _313_v67 _dafny.Array
				_ = _313_v67
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_10
				var _nw54 _dafny.Array
				_ = _nw54
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw54 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Sequence = (func(_314_v4 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_315_i8 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_314_v4), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(159))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg17 _dafny.Int) interface{} {
									return coer17(arg17)
								}
							}(func(_316_i9 _dafny.Int) _dafny.Int {
								return (_dafny.MultiSetOf(_dafny.CodePoint('o'), _dafny.CodePoint('t'), _dafny.CodePoint('v'), _dafny.CodePoint('n'), _dafny.CodePoint('m'))).Cardinality()
							})))
						}
					})(_229_v4)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw54 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw54).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw54).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_313_v67 = _nw54
				var _317_v68 _dafny.Sequence
				_ = _317_v68
				_317_v68 = _dafny.SeqOf(_229_v4, _229_v4, _229_v4, _229_v4, _229_v4)
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_313_v67), 0))
				_ = _index63
				(_313_v67).ArraySet1(_317_v68, (_index63).Int())
				var _318_v69 _dafny.Sequence
				_ = _318_v69
				_318_v69 = _dafny.SeqOf(_312_v66.F15, _312_v66.F15, Companion_Default___.Fm2(_228_globalState), _226_v2)
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_313_v67), 0))
				_ = _index64
				(_313_v67).ArraySet1(Companion_Default___.Fm7((_229_v4).Plus(_229_v4), true, _318_v69, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhvsj")).Cardinality()), _228_globalState), (_index64).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _319_v70 *C0
	_ = _319_v70
	var _nw55 *C0 = New_C0_()
	_ = _nw55
	_nw55.Ctor__(false)
	_319_v70 = _nw55
	var _320_v71 D4
	_ = _320_v71
	_320_v71 = Companion_D4_.Create_DC9_(_223_v0, _319_v70, !(_319_v70.F15))
	var _321_v72 *C0
	_ = _321_v72
	var _nw56 *C0 = New_C0_()
	_ = _nw56
	_nw56.Ctor__((_320_v71).Dtor_cf12())
	_321_v72 = _nw56
	var _322_v73 _dafny.Array
	_ = _322_v73
	var _len0_11 _dafny.Int = _dafny.IntOfInt64(21)
	_ = _len0_11
	var _nw57 _dafny.Array
	_ = _nw57
	if _len0_11.Cmp(_dafny.Zero) == 0 {
		_nw57 = _dafny.NewArray(_len0_11)
	} else {
		var _init11 func(_dafny.Int) _dafny.Int = (func(_323_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_324_i10 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_324_i10, _323_v4)
			}
		})(_229_v4)
		_ = _init11
		var _element0_11 = _init11(_dafny.Zero)
		_ = _element0_11
		_nw57 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
		(_nw57).ArraySet1(_element0_11, 0)
		var _nativeLen0_11 = (_len0_11).Int()
		_ = _nativeLen0_11
		for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
			(_nw57).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
		}
	}
	_322_v73 = _nw57
	var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _index65
	(_322_v73).ArraySet1(_229_v4, (_index65).Int())
	var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _index66
	(_322_v73).ArraySet1(Companion_Default___.Fm0((false) || (_319_v70.F15), _229_v4, _321_v72.F15, _228_globalState), (_index66).Int())
	var _325_v74 _dafny.CodePoint
	_ = _325_v74
	_325_v74 = _dafny.CodePoint('f')
	var _326_v75 _dafny.Map
	_ = _326_v75
	_326_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_321_v72.F15, _325_v74)
	var _pat_let_tv6 = _322_v73
	_ = _pat_let_tv6
	var _pat_let_tv7 = _322_v73
	_ = _pat_let_tv7
	var _327_v76 D4
	_ = _327_v76
	_327_v76 = Companion_D4_.Create_DC10_(_229_v4, _325_v74, _229_v4, _226_v2, _dafny.SeqOf(func(_pat_let7_0 D3) D3 {
		return func(_328_dt__update__tmp_h1 D3) D3 {
			return func(_pat_let8_0 _dafny.Int) D3 {
				return func(_329_dt__update_hcf7_h0 _dafny.Int) D3 {
					return Companion_D3_.Create_DC6_(_329_dt__update_hcf7_h0)
				}(_pat_let8_0)
			}((_pat_let_tv7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_pat_let_tv6), 0))).Int()).(_dafny.Int))
		}(_pat_let7_0)
	}(Companion_D3_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kjuqikydf")).Cardinality()))), Companion_Default___.Fm8((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_319_v70.F15, (_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int))).Cardinality(), Companion_Default___.Fm2(_228_globalState))).Cardinality(), _226_v2, _228_globalState)))
	var _330_v77 D3
	_ = _330_v77
	_330_v77 = Companion_D3_.Create_DC5_(_dafny.SeqOf(_321_v72.F15))
	var _pat_let_tv8 = _322_v73
	_ = _pat_let_tv8
	var _pat_let_tv9 = _322_v73
	_ = _pat_let_tv9
	var _pat_let_tv10 = _229_v4
	_ = _pat_let_tv10
	var _pat_let_tv11 = _229_v4
	_ = _pat_let_tv11
	var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _index67
	var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _index68
	var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _index69
	var _rhs80 _dafny.Int = (_229_v4).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.MultiSetFromSeq(_dafny.SeqOf(_225_v1, _225_v1, _225_v1)))).Cardinality())
	_ = _rhs80
	var _rhs81 _dafny.Int = (_229_v4).Minus((_229_v4).Plus((_326_v75).Cardinality()))
	_ = _rhs81
	var _rhs82 bool = _321_v72.F15
	_ = _rhs82
	var _rhs83 _dafny.Int = (_327_v76).Dtor_cf13()
	_ = _rhs83
	var _rhs84 _dafny.Int = func(_source4 D3) _dafny.Int {
		if _source4.Is_DC6() {
			var _331___mcc_h3 _dafny.Int = _source4.Get_().(D3_DC6).Cf7
			_ = _331___mcc_h3
			var _332_cf7 _dafny.Int = _331___mcc_h3
			_ = _332_cf7
			return ((_pat_let_tv9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_pat_let_tv8), 0))).Int()).(_dafny.Int)).Times(_332_cf7)
		} else if _source4.Is_DC5() {
			var _333___mcc_h4 _dafny.Sequence = _source4.Get_().(D3_DC5).Cf6
			_ = _333___mcc_h4
			var _334_cf6 _dafny.Sequence = _333___mcc_h4
			_ = _334_cf6
			return _pat_let_tv10
		} else {
			var _335___mcc_h5 D3 = _source4.Get_().(D3_DC7).Cf8
			_ = _335___mcc_h5
			var _336_cf8 D3 = _335___mcc_h5
			_ = _336_cf8
			return (_dafny.Zero).Minus(_pat_let_tv11)
		}
	}(_330_v77)
	_ = _rhs84
	var _lhs69 _dafny.Array = _322_v73
	_ = _lhs69
	var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _lhs70
	var _lhs71 _dafny.Array = _322_v73
	_ = _lhs71
	var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _lhs72
	var _lhs73 *GlobalState = _228_globalState
	_ = _lhs73
	var _lhs74 _dafny.Array = _322_v73
	_ = _lhs74
	var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _lhs75
	(_lhs69).ArraySet1(_rhs80, (_lhs70).Int())
	(_lhs71).ArraySet1(_rhs81, (_lhs72).Int())
	_lhs73.F1 = _rhs82
	_229_v4 = _rhs83
	(_lhs74).ArraySet1(_rhs84, (_lhs75).Int())
	(_228_globalState).F10 = (_229_v4).Cmp(_229_v4) >= 0
	(_228_globalState).F6 = ((_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int), _229_v4)).Cardinality()))
	var _hi0 _dafny.Int = _229_v4
	_ = _hi0
	for _337_i11 := _dafny.IntOfInt64(685); _337_i11.Cmp(_hi0) < 0; _337_i11 = _337_i11.Plus(_dafny.One) {
		var _338_v78 _dafny.Map
		_ = _338_v78
		_338_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_i11, _dafny.IntOfInt64(-16))
		var _339_v79 _dafny.Map
		_ = _339_v79
		_339_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v74, _319_v70.F15)
		var _340_v80 _dafny.MultiSet
		_ = _340_v80
		_340_v80 = _dafny.MultiSetOf((_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int), _229_v4, _229_v4, (_339_v79).Cardinality())
		var _341_v81 _dafny.Map
		_ = _341_v81
		_341_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_319_v70, _dafny.IntOfInt64(159))
		var _342_v82 _dafny.Array
		_ = _342_v82
		var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
		_ = _nw58
		_342_v82 = _nw58
		var _343_v83 _dafny.Sequence
		_ = _343_v83
		_343_v83 = _dafny.SeqOf(_321_v72.F15, _319_v70.F15)
		var _344_v84 _dafny.Map
		_ = _344_v84
		_344_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v82, _dafny.IntOfUint32((_343_v83).Cardinality()))
		var _345_v85 _dafny.Map
		_ = _345_v85
		_345_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_338_v78).Contains(_229_v4) {
				return (_338_v78).Get(_229_v4).(_dafny.Int)
			}
			return (func() _dafny.Int {
				if (_340_v80).Contains(_dafny.IntOfInt64(8)) {
					return (_340_v80).Multiplicity(_dafny.IntOfInt64(8))
				}
				return _337_i11
			})()
		})(), ((_341_v81).Update(_319_v70, (_dafny.Zero).Minus((_344_v84).Cardinality()))).Merge(_341_v81))
		_345_v85 = (_345_v85).Update((_dafny.Zero).Minus((_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int)), (func() _dafny.Map {
			if _319_v70.F15 {
				return _341_v81
			}
			return _341_v81
		})())
		Companion_Default___.M0(_228_globalState)
		_340_v80 = _340_v80
		var _346_v86 _dafny.Sequence
		_ = _346_v86
		_346_v86 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(824))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_347_v74 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_348_i13 _dafny.Int) _dafny.CodePoint {
				return _347_v74
			}
		})(_325_v74))))
		var _349_v87 _dafny.Sequence
		_ = _349_v87
		_349_v87 = _dafny.UnicodeSeqOfUtf8Bytes("kqs")
		var _350_v88 _dafny.Set
		_ = _350_v88
		_350_v88 = _dafny.SetOf(_226_v2, _319_v70.F15, _319_v70.F15, _321_v72.F15, _319_v70.F15)
		var _351_v89 D1
		_ = _351_v89
		_351_v89 = Companion_D1_.Create_DC2_((_350_v88).Cardinality())
		var _352_v90 _dafny.Array
		_ = _352_v90
		var _nwElement0_17 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(687))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_353_v74 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_354_i12 _dafny.Int) _dafny.CodePoint {
				return _353_v74
			}
		})(_325_v74)))
		_ = _nwElement0_17
		var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(11))
		_ = _nw59
		(_nw59).ArraySet1(_nwElement0_17, 0)
		(_nw59).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_346_v86).Select((Companion_Default___.SafeIndex(_337_i11, _dafny.IntOfUint32((_346_v86).Cardinality()))).Uint32()).(_dafny.Sequence), _349_v87), 1)
		(_nw59).ArraySet1(_349_v87, 2)
		(_nw59).ArraySet1(_349_v87, 3)
		(_nw59).ArraySet1(_349_v87, 4)
		(_nw59).ArraySet1(_349_v87, 5)
		(_nw59).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(847))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_355_v74 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_356_i14 _dafny.Int) _dafny.CodePoint {
				return _355_v74
			}
		})(_325_v74))), 6)
		(_nw59).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_349_v87, _dafny.Companion_Sequence_.Update(_349_v87, (Companion_Default___.SafeIndex(_229_v4, _dafny.IntOfUint32((_349_v87).Cardinality()))).Uint32(), _325_v74)), 7)
		(_nw59).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_349_v87, _349_v87), 8)
		(_nw59).ArraySet1(_349_v87, 9)
		(_nw59).ArraySet1(Companion_Default___.Fm6(_229_v4, (_320_v71).Dtor_cf12(), _351_v89, _226_v2, _228_globalState), 10)
		_352_v90 = _nw59
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_352_v90), 0))
		_ = _index70
		(_352_v90).ArraySet1(_349_v87, (_index70).Int())
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_352_v90), 0))
		_ = _index71
		var _rhs85 bool = _319_v70.F15
		_ = _rhs85
		var _rhs86 bool = _319_v70.F15
		_ = _rhs86
		var _rhs87 bool = (_226_v2) || (!((_337_i11).Cmp(_337_i11) <= 0))
		_ = _rhs87
		var _rhs88 _dafny.Sequence = _349_v87
		_ = _rhs88
		var _rhs89 _dafny.Int = (((_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int)).Plus((_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int))).Times(_229_v4)
		_ = _rhs89
		var _lhs76 *C0 = _321_v72
		_ = _lhs76
		var _lhs77 *GlobalState = _228_globalState
		_ = _lhs77
		var _lhs78 *GlobalState = _228_globalState
		_ = _lhs78
		var _lhs79 _dafny.Array = _352_v90
		_ = _lhs79
		var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_352_v90), 0))
		_ = _lhs80
		var _lhs81 *GlobalState = _228_globalState
		_ = _lhs81
		_lhs76.F15 = _rhs85
		_lhs77.F10 = _rhs86
		_lhs78.F10 = _rhs87
		(_lhs79).ArraySet1(_rhs88, (_lhs80).Int())
		_lhs81.F0 = _rhs89
	}
	var _357_v91 _dafny.Sequence
	_ = _357_v91
	_357_v91 = _dafny.UnicodeSeqOfUtf8Bytes("ymky")
	var _hi1 _dafny.Int = Companion_Default___.Fm0(_319_v70.F15, _dafny.IntOfUint32((_357_v91).Cardinality()), Companion_Default___.Fm2(_228_globalState), _228_globalState)
	_ = _hi1
	for _358_i15 := Companion_Default___.SafeModuloInt(_229_v4, (_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int)); _358_i15.Cmp(_hi1) < 0; _358_i15 = _358_i15.Plus(_dafny.One) {
		var _359_v92 _dafny.Map
		_ = _359_v92
		_359_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v2, _321_v72)
		var _360_v93 _dafny.Set
		_ = _360_v93
		_360_v93 = _dafny.SetOf((_359_v92).Update(_319_v70.F15, _321_v72), _359_v92)
		(_228_globalState).F6 = (_360_v93).Cardinality()
		_322_v73 = _322_v73
		Companion_Default___.M0(_228_globalState)
		var _361_v94 _dafny.Sequence
		_ = _361_v94
		_361_v94 = _dafny.SeqOf(_319_v70.F15)
		var _362_v95 _dafny.Sequence
		_ = _362_v95
		_362_v95 = _dafny.SeqOf((_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int))
		var _363_v96 D1
		_ = _363_v96
		_363_v96 = Companion_D1_.Create_DC2_((_dafny.MultiSetFromSeq(_362_v95)).Cardinality())
		var _pat_let_tv12 = _321_v72
		_ = _pat_let_tv12
		var _pat_let_tv13 = _229_v4
		_ = _pat_let_tv13
		var _pat_let_tv14 = _321_v72
		_ = _pat_let_tv14
		var _pat_let_tv15 = _228_globalState
		_ = _pat_let_tv15
		_357_v91 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm6(_dafny.IntOfInt64(882), (_361_v94).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_357_v91).Cardinality()), (_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int), (_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int), (_230_v5).Cardinality())).Cardinality()), _dafny.IntOfUint32((_361_v94).Cardinality()))).Uint32()).(bool), func(_pat_let9_0 D1) D1 {
			return func(_364_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let10_0 _dafny.Int) D1 {
					return func(_365_dt__update_hcf2_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC2_(_365_dt__update_hcf2_h0)
					}(_pat_let10_0)
				}(Companion_Default___.Fm0(_pat_let_tv12.F15, _pat_let_tv13, _pat_let_tv14.F15, _pat_let_tv15))
			}(_pat_let9_0)
		}(_363_v96), _226_v2, _228_globalState), _357_v91)
	}
	var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _index72
	var _rhs90 _dafny.Int = (_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int)
	_ = _rhs90
	var _rhs91 bool = !(_226_v2) || (_319_v70.F15)
	_ = _rhs91
	var _lhs82 _dafny.Array = _322_v73
	_ = _lhs82
	var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))
	_ = _lhs83
	(_lhs82).ArraySet1(_rhs90, (_lhs83).Int())
	_226_v2 = _rhs91
	var _366_v97 _dafny.MultiSet
	_ = _366_v97
	_366_v97 = _dafny.MultiSetOf(_319_v70.F15, false)
	var _367_v98 *C0
	_ = _367_v98
	var _nw60 *C0 = New_C0_()
	_ = _nw60
	_nw60.Ctor__((_366_v97).IsSubsetOf(_366_v97))
	_367_v98 = _nw60
	var _368_v99 _dafny.Sequence
	_ = _368_v99
	_368_v99 = _dafny.SeqOf(_321_v72)
	var _369_v100 _dafny.Sequence
	_ = _369_v100
	_369_v100 = _dafny.SeqOf(_dafny.IntOfUint32((_368_v99).Cardinality()))
	_325_v74 = Companion_Default___.Fm9(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_369_v100).Cardinality()), (_322_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_322_v73), 0))).Int()).(_dafny.Int)), _dafny.UnicodeSeqOfUtf8Bytes("xxgmqx"), _229_v4, _228_globalState)
	(_228_globalState).F10 = !(!(_367_v98.F15))
	_357_v91 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sbo"), _357_v91), (Companion_Default___.SafeIndex(_229_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sbo"), _357_v91)).Cardinality()))).Uint32(), _dafny.CodePoint('e')), _357_v91)
	_dafny.Print((_223_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v0).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v1).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_226_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_227_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F9).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState.F11).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F12()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_229_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(121))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_310_i7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_319_v70.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v71).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_320_v71).Dtor_cf11().F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_320_v71).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_321_v72.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v73).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_325_v74)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_326_v75).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('f'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_327_v76).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_327_v76).Dtor_cf14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_327_v76).Dtor_cf15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_327_v76).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_327_v76).Dtor_cf17(), _dafny.SeqOf(Companion_D3_.Create_DC6_(_dafny.IntOfInt64(121)), Companion_D3_.Create_DC6_(_dafny.IntOfInt64(522)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_330_v77).Dtor_cf6(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_357_v91.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_366_v97).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_367_v98.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_368_v99).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_369_v100, _dafny.SeqOf(_dafny.One)))
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
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Array) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D0) Dtor_cf0() _dafny.Array {
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
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
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

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
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
	Cf4 _dafny.Int
	Cf5 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.Int, Cf5 _dafny.Int) D2 {
	return D2{D2_DC4{Cf4, Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC3 struct {
	Cf3 _dafny.Sequence
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf3 _dafny.Sequence) D2 {
	return D2{D2_DC3{Cf3}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf3() _dafny.Sequence {
	return _this.Get_().(D2_DC3).Cf3
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + data.Cf3.VerbatimString(true) + ")"
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
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf7 _dafny.Int
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf7 _dafny.Int) D3 {
	return D3{D3_DC6{Cf7}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC5 struct {
	Cf6 _dafny.Sequence
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf6 _dafny.Sequence) D3 {
	return D3{D3_DC5{Cf6}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

type D3_DC7 struct {
	Cf8 D3
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf8 D3) D3 {
	return D3{D3_DC7{Cf8}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC6_(_dafny.Zero)
}

func (_this D3) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf7
}

func (_this D3) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D3_DC5).Cf6
}

func (_this D3) Dtor_cf8() D3 {
	return _this.Get_().(D3_DC7).Cf8
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf6.Equals(data2.Cf6)
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf8.Equals(data2.Cf8)
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

type D4_DC9 struct {
	Cf10 _dafny.Array
	Cf11 *C0
	Cf12 bool
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf10 _dafny.Array, Cf11 *C0, Cf12 bool) D4 {
	return D4{D4_DC9{Cf10, Cf11, Cf12}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC10 struct {
	Cf13 _dafny.Int
	Cf14 _dafny.CodePoint
	Cf15 _dafny.Int
	Cf16 bool
	Cf17 _dafny.Sequence
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf13 _dafny.Int, Cf14 _dafny.CodePoint, Cf15 _dafny.Int, Cf16 bool, Cf17 _dafny.Sequence) D4 {
	return D4{D4_DC10{Cf13, Cf14, Cf15, Cf16, Cf17}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC11 struct {
	Cf18 _dafny.Int
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf18 _dafny.Int) D4 {
	return D4{D4_DC11{Cf18}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC8 struct {
	Cf9 bool
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf9 bool) D4 {
	return D4{D4_DC8{Cf9}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC12 struct {
	Cf19 D4
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf19 D4) D4 {
	return D4{D4_DC12{Cf19}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC9_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), (*C0)(nil), false)
}

func (_this D4) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf10
}

func (_this D4) Dtor_cf11() *C0 {
	return _this.Get_().(D4_DC9).Cf11
}

func (_this D4) Dtor_cf12() bool {
	return _this.Get_().(D4_DC9).Cf12
}

func (_this D4) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf13
}

func (_this D4) Dtor_cf14() _dafny.CodePoint {
	return _this.Get_().(D4_DC10).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) Dtor_cf16() bool {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf18
}

func (_this D4) Dtor_cf9() bool {
	return _this.Get_().(D4_DC8).Cf9
}

func (_this D4) Dtor_cf19() D4 {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16 && data1.Cf17.Equals(data2.Cf17)
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf9 == data2.Cf9
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf19.Equals(data2.Cf19)
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
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_() D5 {
	return D5{D5_DC14{}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf20 _dafny.Map
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf20 _dafny.Map) D5 {
	return D5{D5_DC13{Cf20}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_()
}

func (_this D5) Dtor_cf20() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf20) + ")"
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
			_, ok := other.Get_().(D5_DC14)
			return ok
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
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
	Cf23 *C0
	Cf24 _dafny.Int
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 bool, Cf23 *C0, Cf24 _dafny.Int) D6 {
	return D6{D6_DC16{Cf22, Cf23, Cf24}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf21 _dafny.Array
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf21 _dafny.Array) D6 {
	return D6{D6_DC15{Cf21}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(false, (*C0)(nil), _dafny.Zero)
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) Dtor_cf23() *C0 {
	return _this.Get_().(D6_DC16).Cf23
}

func (_this D6) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf24
}

func (_this D6) Dtor_cf21() _dafny.Array {
	return _this.Get_().(D6_DC15).Cf21
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0
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

type D7_DC17 struct {
	Cf25 _dafny.Array
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf25 _dafny.Array) D7 {
	return D7{D7_DC17{Cf25}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_()
}

func (_this D7) Dtor_cf25() _dafny.Array {
	return _this.Get_().(D7_DC17).Cf25
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf25) + ")"
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
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf25 == data2.Cf25
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

type D8_DC20 struct {
	Cf27 _dafny.Int
	Cf28 _dafny.Int
	Cf29 _dafny.Int
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf27 _dafny.Int, Cf28 _dafny.Int, Cf29 _dafny.Int) D8 {
	return D8{D8_DC20{Cf27, Cf28, Cf29}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC21 struct {
	Cf30 bool
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf30 bool) D8 {
	return D8{D8_DC21{Cf30}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC19 struct {
	Cf26 _dafny.Set
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf26 _dafny.Set) D8 {
	return D8{D8_DC19{Cf26}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC20_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D8) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D8_DC20).Cf27
}

func (_this D8) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D8_DC20).Cf28
}

func (_this D8) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D8_DC20).Cf29
}

func (_this D8) Dtor_cf30() bool {
	return _this.Get_().(D8_DC21).Cf30
}

func (_this D8) Dtor_cf26() _dafny.Set {
	return _this.Get_().(D8_DC19).Cf26
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf30 == data2.Cf30
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F1   bool
	F4   _dafny.Int
	F6   _dafny.Int
	F8   bool
	F9   _dafny.Array
	F10  bool
	F11  _dafny.Array
	_f2  bool
	_f3  _dafny.Int
	_f5  _dafny.Int
	_f7  bool
	_f12 _dafny.Map
	_f13 _dafny.Int
	_f14 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F1 = false
	_this.F4 = _dafny.Zero
	_this.F6 = _dafny.Zero
	_this.F8 = false
	_this.F9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F10 = false
	_this.F11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = false
	_this._f3 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f7 = false
	_this._f12 = _dafny.EmptyMap
	_this._f13 = _dafny.Zero
	_this._f14 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 bool, f3 _dafny.Int, f4 _dafny.Int, f5 _dafny.Int, f6 _dafny.Int, f7 bool, f8 bool, f9 _dafny.Array, f10 bool, f11 _dafny.Array, f12 _dafny.Map, f13 _dafny.Int, f14 bool) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this).F9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F12() _dafny.Map {
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

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F15 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F15 = false
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

func (_this *C0) Ctor__(f15 bool) {
	{
		(_this).F15 = f15
	}
}
func (_this *C0) Fm1(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _dafny.IntOfInt64(-178))).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dbuyqaqah")).Cardinality())), _this.F15)
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
