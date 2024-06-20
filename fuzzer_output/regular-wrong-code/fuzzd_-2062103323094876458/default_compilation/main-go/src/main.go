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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-120), false)).Cardinality(), (_dafny.IntOfInt64(-352)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(873))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
		if false {
			return _dafny.UnicodeSeqOfUtf8Bytes("ewqpfrw")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("pfb")
	})(), _dafny.CodePoint('k'))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_(!(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(175))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	})), _dafny.UnicodeSeqOfUtf8Bytes("ol"))).Cardinality())) != 0))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("k"), _dafny.UnicodeSeqOfUtf8Bytes("pgy"))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('k')
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfInt64(560), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mng")).Cardinality())), _dafny.IntOfInt64(987)), (func() _dafny.Int {
		if true {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()
		}
		return _dafny.IntOfInt64(-276)
	})(), _dafny.IntOfInt64(203))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.SetOf(!(false))).Contains(true), !(true) || (true), true)
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Map {
	if true {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false, true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(true, true, true, false, false)))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SeqOf(false))
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if false {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(636))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_2_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fqfptov")).Cardinality())
			}))
		}
		return _dafny.SeqOf(_dafny.IntOfInt64(-151))
	})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(890), false)).Cardinality(), _dafny.IntOfInt64(505)), _dafny.SeqOf(_dafny.IntOfInt64(768))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D5 {
	var _source0 D5 = Companion_D5_.Create_DC13_(_dafny.IntOfInt64(879), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), !(false))
	_ = _source0
	if _source0.Is_DC13() {
		var _3___mcc_h0 _dafny.Int = _source0.Get_().(D5_DC13).Cf19
		_ = _3___mcc_h0
		var _4___mcc_h1 _dafny.Int = _source0.Get_().(D5_DC13).Cf20
		_ = _4___mcc_h1
		var _5___mcc_h2 bool = _source0.Get_().(D5_DC13).Cf21
		_ = _5___mcc_h2
		var _6_cf21 bool = _5___mcc_h2
		_ = _6_cf21
		var _7_cf20 _dafny.Int = _4___mcc_h1
		_ = _7_cf20
		var _8_cf19 _dafny.Int = _3___mcc_h0
		_ = _8_cf19
		return Companion_D5_.Create_DC12_(_dafny.UnicodeSeqOfUtf8Bytes("msdl"))
	} else {
		var _9___mcc_h3 _dafny.Sequence = _source0.Get_().(D5_DC12).Cf18
		_ = _9___mcc_h3
		var _10_cf18 _dafny.Sequence = _9___mcc_h3
		_ = _10_cf18
		if true {
			return Companion_D5_.Create_DC12_(_10_cf18)
		} else {
			return Companion_D5_.Create_DC12_(_10_cf18)
		}
	}
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.Map {
	var _source1 D5 = Companion_D5_.Create_DC12_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})))
	_ = _source1
	if _source1.Is_DC13() {
		var _12___mcc_h0 _dafny.Int = _source1.Get_().(D5_DC13).Cf19
		_ = _12___mcc_h0
		var _13___mcc_h1 _dafny.Int = _source1.Get_().(D5_DC13).Cf20
		_ = _13___mcc_h1
		var _14___mcc_h2 bool = _source1.Get_().(D5_DC13).Cf21
		_ = _14___mcc_h2
		var _15_cf21 bool = _14___mcc_h2
		_ = _15_cf21
		var _16_cf20 _dafny.Int = _13___mcc_h1
		_ = _16_cf20
		var _17_cf19 _dafny.Int = _12___mcc_h0
		_ = _17_cf19
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_cf20, _15_cf21), _17_cf19)).Merge(func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.SetOf(_15_cf21)).Cardinality())).Cardinality(), _15_cf21), _15_cf21)).Keys().Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _18_v0 _dafny.Map
				_18_v0 = interface{}(_compr_0).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.SetOf(_15_cf21)).Cardinality())).Cardinality(), _15_cf21), _15_cf21)).Contains(_18_v0) {
					_coll0.Add(_18_v0, _16_cf20)
				}
			}
			return _coll0.ToMap()
		}())
	} else {
		var _19___mcc_h3 _dafny.Sequence = _source1.Get_().(D5_DC12).Cf18
		_ = _19___mcc_h3
		var _20_cf18 _dafny.Sequence = _19___mcc_h3
		_ = _20_cf18
		return (func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(857), _dafny.IntOfInt64(-21))); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _21_v2 _dafny.Int
					_21_v2 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(857)).Cmp(_21_v2) <= 0) && ((_21_v2).Cmp(_dafny.IntOfInt64(-21)) < 0) {
						_coll2.Add((_21_v2).Times(_dafny.IntOfInt64(290)), false)
					}
				}
				return _coll2.ToMap()
			}(), true)).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _22_v1 _dafny.Map
				_22_v1 = interface{}(_compr_1).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll3 = _dafny.NewMapBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(857), _dafny.IntOfInt64(-21))); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _21_v2 _dafny.Int
						_21_v2 = interface{}(_compr_3).(_dafny.Int)
						if ((_dafny.IntOfInt64(857)).Cmp(_21_v2) <= 0) && ((_21_v2).Cmp(_dafny.IntOfInt64(-21)) < 0) {
							_coll3.Add((_21_v2).Times(_dafny.IntOfInt64(290)), false)
						}
					}
					return _coll3.ToMap()
				}(), true)).Contains(_22_v1) {
					_coll1.Add(_22_v1, _dafny.IntOfInt64(542))
				}
			}
			return _coll1.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(327))).Cardinality(), !(true)), _dafny.IntOfInt64(610)))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true, false, (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(322), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(558))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_23_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))).Cardinality()))).Cardinality())).Cardinality())).Cmp((_dafny.SetOf(_dafny.IntOfInt64(76), _dafny.IntOfInt64(112))).Cardinality()) > 0)
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 _dafny.Map, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(508), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('o'))).Cardinality())).Cardinality())))).Cardinality()), _dafny.MultiSetFromSeq(_dafny.SeqOf(false, true, true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_24_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
	})), _dafny.MultiSetFromSeq(_dafny.SeqOf(false))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!(!(false))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC13_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qkgiwheqr")).Cardinality()), (_dafny.IntOfInt64(303)).Times(_dafny.IntOfInt64(832)), !(true) || (true))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Equals(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)))), _dafny.CodePoint('u'))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Map, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var _25_v0 _dafny.Set
	_ = _25_v0
	_25_v0 = _dafny.SetOf(p0)
	var _26_v1 _dafny.Array
	_ = _26_v1
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
	_ = _nw0
	_26_v1 = _nw0
	var _27_v2 _dafny.Array
	_ = _27_v2
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_0
	var _nw1 _dafny.Array
	_ = _nw1
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw1 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Sequence = func(_28_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(_dafny.IntOfInt64(468))
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
	_27_v2 = _nw1
	var _29_v3 _dafny.Map
	_ = _29_v3
	_29_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
	var _30_v4 _dafny.Sequence
	_ = _30_v4
	_30_v4 = _dafny.UnicodeSeqOfUtf8Bytes("knce")
	var _31_v5 _dafny.Map
	_ = _31_v5
	_31_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _30_v4)
	var _32_v6 _dafny.Array
	_ = _32_v6
	var _nwElement0_0 bool = false
	_ = _nwElement0_0
	var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(22))
	_ = _nw2
	(_nw2).ArraySet1(_nwElement0_0, 0)
	(_nw2).ArraySet1((Companion_D9_.Create_DC26_(p0, _25_v0, _26_v1, _27_v2)).Dtor_cf51(), 1)
	(_nw2).ArraySet1(p0, 2)
	(_nw2).ArraySet1(p0, 3)
	(_nw2).ArraySet1((func() bool {
		if (_29_v3).Contains(Companion_Default___.Fm1(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(176))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_33_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), globalState)) {
			return (_29_v3).Get(Companion_Default___.Fm1(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(176))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_34_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), globalState)).(bool)
		}
		return !(p0)
	})(), 4)
	(_nw2).ArraySet1(p0, 5)
	(_nw2).ArraySet1(p0, 6)
	(_nw2).ArraySet1(p0, 7)
	(_nw2).ArraySet1(p0, 8)
	(_nw2).ArraySet1(p0, 9)
	(_nw2).ArraySet1(p0, 10)
	(_nw2).ArraySet1(Companion_Default___.Fm1(p0, (func() _dafny.Sequence {
		if (_31_v5).Contains(p0) {
			return (_31_v5).Get(p0).(_dafny.Sequence)
		}
		return _30_v4
	})(), globalState), 11)
	(_nw2).ArraySet1(p0, 12)
	(_nw2).ArraySet1(p0, 13)
	(_nw2).ArraySet1(p0, 14)
	(_nw2).ArraySet1(false, 15)
	(_nw2).ArraySet1(p0, 16)
	(_nw2).ArraySet1(p0, 17)
	(_nw2).ArraySet1(p0, 18)
	(_nw2).ArraySet1(p0, 19)
	(_nw2).ArraySet1(p0, 20)
	(_nw2).ArraySet1(p0, 21)
	_32_v6 = _nw2
	var _35_v7 _dafny.Sequence
	_ = _35_v7
	_35_v7 = _dafny.SeqOf(_32_v6, _32_v6)
	var _36_v8 _dafny.Int
	_ = _36_v8
	_36_v8 = _dafny.Zero
	var _37_v9 _dafny.Map
	_ = _37_v9
	_37_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
		if p0 {
			return p0
		}
		return p0
	})()), _dafny.Companion_Sequence_.Update(_35_v7, (Companion_Default___.SafeIndex(_36_v8, _dafny.IntOfUint32((_35_v7).Cardinality()))).Uint32(), _32_v6))
	_37_v9 = (_37_v9).Update(!(p0), _35_v7)
	var _38_v10 _dafny.Sequence
	_ = _38_v10
	_38_v10 = _dafny.SeqOf(_36_v8)
	var _39_v11 D5
	_ = _39_v11
	_39_v11 = Companion_D5_.Create_DC13_(_36_v8, (_dafny.Zero).Minus(_36_v8), p0)
	var _pat_let_tv0 = _39_v11
	_ = _pat_let_tv0
	var _pat_let_tv1 = p0
	_ = _pat_let_tv1
	var _pat_let_tv2 = p0
	_ = _pat_let_tv2
	(globalState).F2 = (func(_source2 D2) _dafny.Set {
		if _source2.Is_DC5() {
			var _40___mcc_h0 bool = _source2.Get_().(D2_DC5).Cf5
			_ = _40___mcc_h0
			var _41___mcc_h1 _dafny.Int = _source2.Get_().(D2_DC5).Cf6
			_ = _41___mcc_h1
			var _42___mcc_h2 _dafny.Array = _source2.Get_().(D2_DC5).Cf7
			_ = _42___mcc_h2
			var _43___mcc_h3 _dafny.MultiSet = _source2.Get_().(D2_DC5).Cf8
			_ = _43___mcc_h3
			var _44_cf8 _dafny.MultiSet = _43___mcc_h3
			_ = _44_cf8
			var _45_cf7 _dafny.Array = _42___mcc_h2
			_ = _45_cf7
			var _46_cf6 _dafny.Int = _41___mcc_h1
			_ = _46_cf6
			var _47_cf5 bool = _40___mcc_h0
			_ = _47_cf5
			return _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_cf5, _dafny.CodePoint('a')))
		} else if _source2.Is_DC6() {
			var _48___mcc_h4 _dafny.Int = _source2.Get_().(D2_DC6).Cf9
			_ = _48___mcc_h4
			var _49___mcc_h5 bool = _source2.Get_().(D2_DC6).Cf10
			_ = _49___mcc_h5
			var _50___mcc_h6 bool = _source2.Get_().(D2_DC6).Cf11
			_ = _50___mcc_h6
			var _51___mcc_h7 _dafny.Int = _source2.Get_().(D2_DC6).Cf12
			_ = _51___mcc_h7
			var _52_cf12 _dafny.Int = _51___mcc_h7
			_ = _52_cf12
			var _53_cf11 bool = _50___mcc_h6
			_ = _53_cf11
			var _54_cf10 bool = _49___mcc_h5
			_ = _54_cf10
			var _55_cf9 _dafny.Int = _48___mcc_h4
			_ = _55_cf9
			return (func() _dafny.Set {
				var _coll4 = _dafny.NewBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((func() _dafny.Map {
					var _coll5 = _dafny.NewMapBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_cf10, _dafny.CodePoint('e')), _52_cf12)).Keys().Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _56_v12 _dafny.Map
						_56_v12 = interface{}(_compr_5).(_dafny.Map)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_cf10, _dafny.CodePoint('e')), _52_cf12)).Contains(_56_v12) {
							_coll5.Add(_56_v12, _dafny.CodePoint('t'))
						}
					}
					return _coll5.ToMap()
				}()).Keys().Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _57_v13 _dafny.Map
					_57_v13 = interface{}(_compr_4).(_dafny.Map)
					if (func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_cf10, _dafny.CodePoint('e')), _52_cf12)).Keys().Elements()); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _56_v12 _dafny.Map
							_56_v12 = interface{}(_compr_6).(_dafny.Map)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_cf10, _dafny.CodePoint('e')), _52_cf12)).Contains(_56_v12) {
								_coll6.Add(_56_v12, _dafny.CodePoint('t'))
							}
						}
						return _coll6.ToMap()
					}()).Contains(_57_v13) {
						_coll4.Add(_57_v13)
					}
				}
				return _coll4.ToSet()
			}()).Intersection(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('k')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_cf10, _dafny.CodePoint('h')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv0).Dtor_cf21(), _dafny.CodePoint('f'))))
		} else if _source2.Is_DC4() {
			var _58___mcc_h8 _dafny.Array = _source2.Get_().(D2_DC4).Cf4
			_ = _58___mcc_h8
			var _59_cf4 _dafny.Array = _58___mcc_h8
			_ = _59_cf4
			return _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv1, _dafny.CodePoint('a')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('b')))
		} else {
			var _60___mcc_h9 D2 = _source2.Get_().(D2_DC7).Cf13
			_ = _60___mcc_h9
			var _61_cf13 D2 = _60___mcc_h9
			_ = _61_cf13
			return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv2, _dafny.CodePoint('x')))).Difference(func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_62_i2 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('h'))
				}))).Elements()); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _63_v14 _dafny.Map
					_63_v14 = interface{}(_compr_7).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}(func(_62_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('h'))
					})), _63_v14) {
						_coll7.Add(_63_v14)
					}
				}
				return _coll7.ToSet()
			}())
		}
	}(Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_38_v10).Cardinality()), (_39_v11).Dtor_cf21(), p0, _36_v8))).Cardinality()
	var _64_v15 _dafny.Int
	_ = _64_v15
	_64_v15 = _36_v8
	var _65_v16 *C3
	_ = _65_v16
	var _nw3 *C3 = New_C3_()
	_ = _nw3
	_nw3.Ctor__(_30_v4, _64_v15, _dafny.IntOfInt64(-140))
	_65_v16 = _nw3
	var _66_v17 _dafny.Array
	_ = _66_v17
	var _nwElement0_1 *C3 = _65_v16
	_ = _nwElement0_1
	var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(12))
	_ = _nw4
	(_nw4).ArraySet1(_nwElement0_1, 0)
	(_nw4).ArraySet1(_65_v16, 1)
	(_nw4).ArraySet1(_65_v16, 2)
	(_nw4).ArraySet1(_65_v16, 3)
	(_nw4).ArraySet1(_65_v16, 4)
	(_nw4).ArraySet1(_65_v16, 5)
	(_nw4).ArraySet1(_65_v16, 6)
	(_nw4).ArraySet1(_65_v16, 7)
	(_nw4).ArraySet1(_65_v16, 8)
	(_nw4).ArraySet1(_65_v16, 9)
	(_nw4).ArraySet1(_65_v16, 10)
	(_nw4).ArraySet1(_65_v16, 11)
	_66_v17 = _nw4
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_66_v17), 0))
	_ = _index0
	(_66_v17).ArraySet1(_65_v16, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_66_v17), 0))
	_ = _index1
	var _rhs0 _dafny.Int = (_36_v8).Minus((_dafny.IntOfUint32(((_65_v16).F19()).Cardinality())).Plus((_dafny.Zero).Minus(_36_v8)))
	_ = _rhs0
	var _rhs1 *C3 = _65_v16
	_ = _rhs1
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 _dafny.Array = _66_v17
	_ = _lhs1
	var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_66_v17), 0))
	_ = _lhs2
	_lhs0.F2 = _rhs0
	(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
	var _hi0 _dafny.Int = _36_v8
	_ = _hi0
	for _67_i3 := _36_v8; _67_i3.Cmp(_hi0) < 0; _67_i3 = _67_i3.Plus(_dafny.One) {
		if p0 {
			var _68_v18 *C2
			_ = _68_v18
			var _nw5 *C2 = New_C2_()
			_ = _nw5
			_nw5.Ctor__()
			_68_v18 = _nw5
			var _69_v19 _dafny.MultiSet
			_ = _69_v19
			_69_v19 = _dafny.MultiSetOf(p0)
			var _70_v20 _dafny.Map
			_ = _70_v20
			_70_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v19, _39_v11)
			_70_v20 = (_70_v20).Update(_69_v19, Companion_D5_.Create_DC13_(_67_i3, _dafny.IntOfUint32((_38_v10).Cardinality()), p0))
			(globalState).F9 = _67_i3
			(globalState).F0 = false
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_32_v6), 0))
			_ = _index2
			(_32_v6).ArraySet1(p0, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_32_v6), 0))
			_ = _index3
			(_32_v6).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ph"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(602))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_71_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))), (_index3).Int())
		} else {
			var _72_v21 T0
			_ = _72_v21
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__(_64_v15, _dafny.IntOfInt64(625))
			_72_v21 = _nw6
			var _73_v22 _dafny.Sequence
			_ = _73_v22
			_73_v22 = _dafny.SeqOf(_72_v21, _72_v21, _72_v21)
			var _74_v23 _dafny.Sequence
			_ = _74_v23
			_74_v23 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_73_v22, (Companion_Default___.SafeIndex((_72_v21).F18(), _dafny.IntOfUint32((_73_v22).Cardinality()))).Uint32(), _72_v21))
			var _75_v24 _dafny.Array
			_ = _75_v24
			var _nwElement0_2 _dafny.Sequence = _74_v23
			_ = _nwElement0_2
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(14))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_2, 0)
			(_nw7).ArraySet1(_dafny.SeqOf(_73_v22, _73_v22), 1)
			(_nw7).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_74_v23, _74_v23), 2)
			(_nw7).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_74_v23, _74_v23), 3)
			(_nw7).ArraySet1(_74_v23, 4)
			(_nw7).ArraySet1((func() _dafny.Sequence {
				if p0 {
					return _74_v23
				}
				return _74_v23
			})(), 5)
			(_nw7).ArraySet1(_74_v23, 6)
			(_nw7).ArraySet1(_74_v23, 7)
			(_nw7).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_74_v23, _74_v23), 8)
			(_nw7).ArraySet1(_74_v23, 9)
			(_nw7).ArraySet1(_74_v23, 10)
			(_nw7).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_74_v23, _dafny.SeqOf(_73_v22)), 11)
			(_nw7).ArraySet1(_74_v23, 12)
			(_nw7).ArraySet1(_74_v23, 13)
			_75_v24 = _nw7
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_75_v24), 0))
			_ = _index4
			(_75_v24).ArraySet1(_74_v23, (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_75_v24), 0))
			_ = _index5
			(_75_v24).ArraySet1(_dafny.Companion_Sequence_.Update(_74_v23, (Companion_Default___.SafeIndex(_67_i3, _dafny.IntOfUint32((_74_v23).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_73_v22, _73_v22)), (_index5).Int())
			var _76_v25 _dafny.Map
			_ = _76_v25
			_76_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_38_v10, Companion_Default___.Fm11(_67_i3, p0, (_72_v21).F18(), globalState)), _36_v8)
			var _77_v26 _dafny.Map
			_ = _77_v26
			_77_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _36_v8)
			var _78_v27 _dafny.Sequence
			_ = _78_v27
			_78_v27 = _dafny.SeqOf(_77_v26)
			var _79_v28 _dafny.Set
			_ = _79_v28
			_79_v28 = _dafny.SetOf(_36_v8)
			var _80_v29 _dafny.Map
			_ = _80_v29
			_80_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v8, (_78_v27).Select((Companion_Default___.SafeIndex((_79_v28).Cardinality(), _dafny.IntOfUint32((_78_v27).Cardinality()))).Uint32()).(_dafny.Map))
			var _81_v30 _dafny.MultiSet
			_ = _81_v30
			_81_v30 = _dafny.MultiSetOf(_dafny.IntOfInt64(503))
			_76_v25 = (_76_v25).Update(_38_v10, ((_dafny.Zero).Minus(((func() _dafny.Map {
				if (_80_v29).Contains(_dafny.IntOfUint32(((_65_v16).F19()).Cardinality())) {
					return (_80_v29).Get(_dafny.IntOfUint32(((_65_v16).F19()).Cardinality())).(_dafny.Map)
				}
				return _77_v26
			})()).Cardinality())).Minus((_81_v30).Cardinality()))
			var _82_v31 _dafny.Array
			_ = _82_v31
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_1
			var _nw8 _dafny.Array
			_ = _nw8
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw8 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_83_p0 bool) func(_dafny.Int) _dafny.Int {
					return func(_84_i5 _dafny.Int) _dafny.Int {
						return (_84_i5).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_83_p0)).Cardinality()))
					}
				})(p0)
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
			_82_v31 = _nw8
			var _85_v32 D4
			_ = _85_v32
			_85_v32 = Companion_D4_.Create_DC10_(_82_v31)
			_85_v32 = _85_v32
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_82_v31), 0))
			_ = _index6
			(_82_v31).ArraySet1(_36_v8, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_82_v31), 0))
			_ = _index7
			(_82_v31).ArraySet1(_36_v8, (_index7).Int())
			var _86_v33 _dafny.Sequence
			_ = _86_v33
			_86_v33 = _dafny.SeqOf(p0)
			var _87_v34 _dafny.CodePoint
			_ = _87_v34
			_87_v34 = _dafny.CodePoint('i')
			var _rhs2 _dafny.Int = (_82_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_82_v31), 0))).Int()).(_dafny.Int)
			_ = _rhs2
			var _rhs3 _dafny.Sequence = _dafny.SeqOf(_dafny.Companion_Sequence_.Contains((_65_v16).F19(), _87_v34), p0)
			_ = _rhs3
			var _rhs4 _dafny.Int = Companion_Default___.SafeDivisionInt((_82_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_82_v31), 0))).Int()).(_dafny.Int), (_72_v21).F18())
			_ = _rhs4
			var _rhs5 bool = p0
			_ = _rhs5
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			_lhs3.F2 = _rhs2
			_86_v33 = _rhs3
			_lhs4.F9 = _rhs4
			_lhs5.F3 = _rhs5
		}
		var _88_v35 *C2
		_ = _88_v35
		var _nw9 *C2 = New_C2_()
		_ = _nw9
		_nw9.Ctor__()
		_88_v35 = _nw9
		(globalState).F2 = _36_v8
		var _89_v36 _dafny.CodePoint
		_ = _89_v36
		_89_v36 = _dafny.CodePoint('n')
		var _90_v37 *C1
		_ = _90_v37
		var _nw10 *C1 = New_C1_()
		_ = _nw10
		_nw10.Ctor__(_89_v36)
		_90_v37 = _nw10
		var _91_v38 _dafny.Sequence
		_ = _91_v38
		_91_v38 = _dafny.SeqOf(_90_v37, _90_v37, _90_v37)
		var _rhs6 _dafny.Sequence = _91_v38
		_ = _rhs6
		var _rhs7 bool = (_67_i3).Cmp(((_38_v10).Select((Companion_Default___.SafeIndex(_36_v8, _dafny.IntOfUint32((_38_v10).Cardinality()))).Uint32()).(_dafny.Int)).Times(_67_i3)) < 0
		_ = _rhs7
		var _rhs8 bool = _dafny.Companion_Sequence_.IsProperPrefixOf((_65_v16).F19(), _dafny.Companion_Sequence_.Concatenate((_65_v16).F19(), (_65_v16).F19()))
		_ = _rhs8
		var _rhs9 _dafny.Int = _67_i3
		_ = _rhs9
		var _lhs6 *GlobalState = globalState
		_ = _lhs6
		var _lhs7 *GlobalState = globalState
		_ = _lhs7
		var _lhs8 *GlobalState = globalState
		_ = _lhs8
		_91_v38 = _rhs6
		_lhs6.F0 = _rhs7
		_lhs7.F10 = _rhs8
		_lhs8.F9 = _rhs9
	}
	var _92_v39 _dafny.MultiSet
	_ = _92_v39
	_92_v39 = _dafny.MultiSetOf(p0, p0, p0)
	if (_92_v39).IsProperSubsetOf(_dafny.MultiSetOf(p0, p0)) {
		var _93_v40 _dafny.Array
		_ = _93_v40
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_2
		var _nw11 _dafny.Array
		_ = _nw11
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw11 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_94_v8 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_95_i6 _dafny.Int) _dafny.Int {
					return (_95_i6).Plus(_94_v8)
				}
			})(_36_v8)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw11 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw11).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw11).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_93_v40 = _nw11
		_93_v40 = _93_v40
		var _96_v41 _dafny.Map
		_ = _96_v41
		_96_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v8, (_dafny.Zero).Minus(_36_v8))
		_96_v41 = (_96_v41).Update(_36_v8, _36_v8)
		(globalState).F9 = ((_dafny.IntOfInt64(-709)).Times(_36_v8)).Minus((_36_v8).Times(_36_v8))
		(globalState).F9 = _36_v8
		var _97_v42 D6
		_ = _97_v42
		_97_v42 = Companion_D6_.Create_DC16_(_36_v8, p0, _36_v8, _36_v8, p0)
		(globalState).F9 = ((_97_v42).Dtor_cf28()).Plus(_dafny.IntOfInt64(-774))
	} else {
		var _98_v43 _dafny.Array
		_ = _98_v43
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_3
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Int = (func(_99_v8 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_100_i7 _dafny.Int) _dafny.Int {
					return (_100_i7).Minus(_99_v8)
				}
			})(_36_v8)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw12 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw12).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw12).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_98_v43 = _nw12
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw13
		_98_v43 = _nw13
		var _101_v44 _dafny.MultiSet
		_ = _101_v44
		_101_v44 = _dafny.MultiSetOf(_36_v8, Companion_Default___.Fm0(globalState))
		var _102_v45 _dafny.Set
		_ = _102_v45
		_102_v45 = _dafny.SetOf(_36_v8)
		var _103_v46 *C0
		_ = _103_v46
		var _nw14 *C0 = New_C0_()
		_ = _nw14
		_nw14.Ctor__(_36_v8, (_102_v45).Cardinality())
		_103_v46 = _nw14
		if (!(!(!(p0)))) && (!(p0)) {
			var _104_v47 *C0
			_ = _104_v47
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(_64_v15, _dafny.IntOfInt64(178))
			_104_v47 = _nw15
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_98_v43), 0))
			_ = _index8
			(_98_v43).ArraySet1((func() _dafny.Int {
				if p0 {
					return _36_v8
				}
				return _36_v8
			})(), (_index8).Int())
			var _105_v48 _dafny.Map
			_ = _105_v48
			_105_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _29_v3)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_98_v43), 0))
			_ = _index9
			var _rhs10 *C0 = _103_v46
			_ = _rhs10
			var _rhs11 bool = true
			_ = _rhs11
			var _rhs12 _dafny.Int = (_dafny.Zero).Minus(((_105_v48).Cardinality()).Plus(_36_v8))
			_ = _rhs12
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			var _lhs10 _dafny.Array = _98_v43
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_98_v43), 0))
			_ = _lhs11
			_104_v47 = _rhs10
			_lhs9.F0 = _rhs11
			(_lhs10).ArraySet1(_rhs12, (_lhs11).Int())
			var _106_v49 *C0
			_ = _106_v49
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__(_64_v15, (_98_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_98_v43), 0))).Int()).(_dafny.Int))
			_106_v49 = _nw16
			(globalState).F3 = (func() bool {
				if (_29_v3).Contains(p0) {
					return (_29_v3).Get(p0).(bool)
				}
				return false
			})()
			(globalState).F2 = (_98_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_98_v43), 0))).Int()).(_dafny.Int)
			var _107_v50 _dafny.Sequence
			_ = _107_v50
			_107_v50 = _dafny.SeqOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("vckfvwohs"), (_65_v16).F19()))
			_107_v50 = _107_v50
		} else {
			(globalState).F2 = (_dafny.Zero).Minus(((_36_v8).Times(_36_v8)).Minus(Companion_Default___.SafeDivisionInt(_36_v8, _36_v8)))
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_26_v1), 0))
			_ = _index10
			(_26_v1).ArraySet1(p0, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_26_v1), 0))
			_ = _index11
			(_26_v1).ArraySet1(p0, (_index11).Int())
			var _108_v51 D11
			_ = _108_v51
			_108_v51 = Companion_D11_.Create_DC32_(_36_v8)
			var _109_v52 _dafny.CodePoint
			_ = _109_v52
			_109_v52 = _dafny.CodePoint('q')
			var _110_v53 *C1
			_ = _110_v53
			var _nw17 *C1 = New_C1_()
			_ = _nw17
			_nw17.Ctor__(_109_v52)
			_110_v53 = _nw17
			var _rhs13 _dafny.Int = Companion_Default___.SafeDivisionInt(_36_v8, _36_v8)
			_ = _rhs13
			var _rhs14 bool = (_26_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_26_v1), 0))).Int()).(bool)
			_ = _rhs14
			var _rhs15 D11 = _108_v51
			_ = _rhs15
			var _rhs16 *C1 = _110_v53
			_ = _rhs16
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			_36_v8 = _rhs13
			_lhs12.F0 = _rhs14
			_108_v51 = _rhs15
			_110_v53 = _rhs16
			var _111_v54 _dafny.Map
			_ = _111_v54
			_111_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v53, _103_v46)
			var _112_v55 _dafny.Sequence
			_ = _112_v55
			_112_v55 = _dafny.SeqOf(_103_v46)
			var _113_v56 _dafny.Array
			_ = _113_v56
			var _nwElement0_3 *C0 = (func() *C0 {
				if p0 {
					return _103_v46
				}
				return _103_v46
			})()
			_ = _nwElement0_3
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(10))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_3, 0)
			(_nw18).ArraySet1((func() *C0 {
				if (_111_v54).Contains(_110_v53) {
					return (_111_v54).Get(_110_v53).(*C0)
				}
				return _103_v46
			})(), 1)
			(_nw18).ArraySet1(_103_v46, 2)
			(_nw18).ArraySet1(_103_v46, 3)
			(_nw18).ArraySet1(_103_v46, 4)
			(_nw18).ArraySet1(_103_v46, 5)
			(_nw18).ArraySet1((_112_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_38_v10).Cardinality()), _dafny.IntOfUint32((_112_v55).Cardinality()))).Uint32()).(*C0), 6)
			(_nw18).ArraySet1(_103_v46, 7)
			(_nw18).ArraySet1(_103_v46, 8)
			(_nw18).ArraySet1(_103_v46, 9)
			_113_v56 = _nw18
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_113_v56), 0))
			_ = _index12
			(_113_v56).ArraySet1(_103_v46, (_index12).Int())
			var _114_v57 D2
			_ = _114_v57
			_114_v57 = Companion_D2_.Create_DC6_(_36_v8, true, p0, _36_v8)
			var _115_v58 _dafny.Sequence
			_ = _115_v58
			_115_v58 = _dafny.SeqOf(_114_v57)
			var _116_v59 _dafny.MultiSet
			_ = _116_v59
			_116_v59 = _dafny.MultiSetOf(_39_v11, _39_v11, Companion_Default___.Fm17(_dafny.IntOfUint32((_30_v4).Cardinality()), _115_v58, globalState))
			var _117_v60 D1
			_ = _117_v60
			_117_v60 = Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_((_26_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_26_v1), 0))).Int()).(bool)))
			var _118_v61 D6
			_ = _118_v61
			_118_v61 = Companion_D6_.Create_DC15_(_36_v8, _36_v8, Companion_Default___.Fm11((_dafny.Zero).Minus(_36_v8), false, _dafny.IntOfInt64(218), globalState), _36_v8, _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()))))
			var _pat_let_tv3 = p0
			_ = _pat_let_tv3
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_113_v56), 0))
			_ = _index13
			var _rhs17 *C0 = _103_v46
			_ = _rhs17
			var _rhs18 _dafny.MultiSet = (_116_v59).Union((_116_v59).Difference(_116_v59))
			_ = _rhs18
			var _rhs19 _dafny.Int = (Companion_Default___.Fm18(_36_v8, globalState)).Cardinality()
			_ = _rhs19
			var _rhs20 _dafny.Int = (_dafny.Zero).Minus((_118_v61).Dtor_cf24())
			_ = _rhs20
			var _rhs21 D1 = func(_pat_let0_0 D1) D1 {
				return func(_119_dt__update__tmp_h3 D1) D1 {
					return func(_pat_let1_0 D1) D1 {
						return func(_120_dt__update_hcf3_h0 D1) D1 {
							return Companion_D1_.Create_DC3_(_120_dt__update_hcf3_h0)
						}(_pat_let1_0)
					}(Companion_D1_.Create_DC2_(_pat_let_tv3))
				}(_pat_let0_0)
			}(_117_v60)
			_ = _rhs21
			var _lhs13 _dafny.Array = _113_v56
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_113_v56), 0))
			_ = _lhs14
			var _lhs15 *GlobalState = globalState
			_ = _lhs15
			var _lhs16 *GlobalState = globalState
			_ = _lhs16
			(_lhs13).ArraySet1(_rhs17, (_lhs14).Int())
			_116_v59 = _rhs18
			_lhs15.F9 = _rhs19
			_lhs16.F9 = _rhs20
			_117_v60 = _rhs21
			var _121_v62 _dafny.Map
			_ = _121_v62
			_121_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v52, !(!(p0)))
			_121_v62 = (_121_v62).Update((_110_v53).F20(), p0)
		}
		(globalState).F10 = p0
		var _122_v63 *C2
		_ = _122_v63
		var _nw19 *C2 = New_C2_()
		_ = _nw19
		_nw19.Ctor__()
		_122_v63 = _nw19
		var _123_v64 _dafny.MultiSet
		_ = _123_v64
		_123_v64 = _dafny.MultiSetOf(_122_v63, _122_v63, _122_v63, _122_v63, _122_v63)
		_123_v64 = _123_v64
	}
	var _124_v65 _dafny.CodePoint
	_ = _124_v65
	_124_v65 = _dafny.CodePoint('j')
	var _125_v66 *C0
	_ = _125_v66
	var _nw20 *C0 = New_C0_()
	_ = _nw20
	_nw20.Ctor__(_64_v15, _36_v8)
	_125_v66 = _nw20
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_32_v6), 0))
	_ = _index14
	(_32_v6).ArraySet1(p0, (_index14).Int())
	var _126_v67 _dafny.Array
	_ = _126_v67
	var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
	_ = _nw21
	_126_v67 = _nw21
	var _127_v69 _dafny.Sequence
	_ = _127_v69
	_127_v69 = _dafny.SeqOf(func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(784), _dafny.IntOfInt64(243))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _128_v68 _dafny.Int
			_128_v68 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(784)).Cmp(_128_v68) <= 0) && ((_128_v68).Cmp(_dafny.IntOfInt64(243)) < 0) {
				_coll8.Add(Companion_Default___.SafeDivisionInt(_128_v68, _36_v8))
			}
		}
		return _coll8.ToSet()
	}())
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_126_v67), 0))
	_ = _index15
	(_126_v67).ArraySet1((_127_v69).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.IntOfUint32((_127_v69).Cardinality()))).Uint32()).(_dafny.Set), (_index15).Int())
	var _129_v70 _dafny.Map
	_ = _129_v70
	_129_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _36_v8)
	var _130_v71 _dafny.MultiSet
	_ = _130_v71
	_130_v71 = _dafny.MultiSetOf(_125_v66)
	var _131_v72 *C2
	_ = _131_v72
	var _nw22 *C2 = New_C2_()
	_ = _nw22
	_nw22.Ctor__()
	_131_v72 = _nw22
	var _132_v73 _dafny.Map
	_ = _132_v73
	_132_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _131_v72)
	var _133_v74 _dafny.Set
	_ = _133_v74
	_133_v74 = _dafny.SetOf(_36_v8, (_132_v73).Cardinality(), Companion_Default___.Fm0(globalState))
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_32_v6), 0))
	_ = _index16
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_126_v67), 0))
	_ = _index17
	var _rhs22 _dafny.Sequence = _30_v4
	_ = _rhs22
	var _rhs23 _dafny.CodePoint = ((_65_v16).F19()).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_129_v70).Cardinality(), (_130_v71).Cardinality()), _dafny.IntOfUint32(((_65_v16).F19()).Cardinality()))).Uint32()).(_dafny.CodePoint)
	_ = _rhs23
	var _rhs24 *C0 = (func() *C0 {
		if Companion_Default___.Fm1(p0, _30_v4, globalState) {
			return _125_v66
		}
		return _125_v66
	})()
	_ = _rhs24
	var _rhs25 bool = p0
	_ = _rhs25
	var _rhs26 _dafny.Set = ((_133_v74).Intersection(_133_v74)).Intersection((_133_v74).Intersection(_133_v74))
	_ = _rhs26
	var _lhs17 _dafny.Array = _32_v6
	_ = _lhs17
	var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_32_v6), 0))
	_ = _lhs18
	var _lhs19 _dafny.Array = _126_v67
	_ = _lhs19
	var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_126_v67), 0))
	_ = _lhs20
	_30_v4 = _rhs22
	_124_v65 = _rhs23
	_125_v66 = _rhs24
	(_lhs17).ArraySet1(_rhs25, (_lhs18).Int())
	(_lhs19).ArraySet1(_rhs26, (_lhs20).Int())
	r0 = (func() _dafny.Int {
		if _dafny.Companion_Sequence_.Contains(_30_v4, _dafny.CodePoint('g')) {
			return _36_v8
		}
		return _36_v8
	})()
	var _134_v75 _dafny.Map
	_ = _134_v75
	_134_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_32_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_32_v6), 0))).Int()).(bool), _124_v65)
	var _135_v76 _dafny.Map
	_ = _135_v76
	_135_v76 = _134_v75
	var _136_v77 _dafny.Map
	_ = _136_v77
	_136_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v8, (_134_v75))
	r1 = ((_136_v77).Update(Companion_Default___.SafeModuloInt(_36_v8, _36_v8), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _124_v65))).Cardinality()
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _137_v0 bool
	_ = _137_v0
	_137_v0 = true
	var _138_v1 _dafny.Array
	_ = _138_v1
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
	_ = _len0_4
	var _nw23 _dafny.Array
	_ = _nw23
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw23 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) bool = (func(_139_v0 bool) func(_dafny.Int) bool {
			return func(_140_i0 _dafny.Int) bool {
				return _139_v0
			}
		})(_137_v0)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw23 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw23).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw23).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_138_v1 = _nw23
	var _141_v2 _dafny.Int
	_ = _141_v2
	_141_v2 = _dafny.IntOfInt64(213)
	var _142_v3 _dafny.Map
	_ = _142_v3
	_142_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, (_dafny.Zero).Minus(_141_v2))
	var _143_v4 _dafny.Sequence
	_ = _143_v4
	_143_v4 = _dafny.UnicodeSeqOfUtf8Bytes("l")
	var _144_v5 _dafny.Map
	_ = _144_v5
	_144_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _137_v0)
	var _145_globalState *GlobalState
	_ = _145_globalState
	var _nw24 *GlobalState = New_GlobalState_()
	_ = _nw24
	_nw24.Ctor__(true, _dafny.SeqOf(_137_v0), _dafny.IntOfInt64(172), false, true, _dafny.CodePoint('f'), _dafny.IntOfInt64(669), _138_v1, _dafny.MultiSetOf(_141_v2, _141_v2), _dafny.IntOfInt64(257), false, (func() _dafny.Array {
		if _137_v0 {
			return _138_v1
		}
		return _138_v1
	})(), _142_v3, _143_v4, false, (_144_v5).Merge(_144_v5), true)
	_145_globalState = _nw24
	if _137_v0 {
		_142_v3 = (_142_v3).Update(_137_v0, _141_v2)
		var _146_v6 _dafny.Array
		_ = _146_v6
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_5
		var _nw25 _dafny.Array
		_ = _nw25
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw25 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Sequence = func(_147_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(false)
			}
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw25 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw25).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw25).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_146_v6 = _nw25
		var _148_v7 _dafny.Sequence
		_ = _148_v7
		_148_v7 = _dafny.SeqOf(_137_v0)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_146_v6), 0))
		_ = _index18
		(_146_v6).ArraySet1(_148_v7, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_146_v6), 0))
		_ = _index19
		(_146_v6).ArraySet1(_dafny.Companion_Sequence_.Update(_148_v7, (Companion_Default___.SafeIndex((func() _dafny.Int {
			if _137_v0 {
				return (_dafny.Zero).Minus(_141_v2)
			}
			return _141_v2
		})(), _dafny.IntOfUint32((_148_v7).Cardinality()))).Uint32(), (_148_v7).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_141_v2), _dafny.IntOfUint32((_148_v7).Cardinality()))).Uint32()).(bool)), (_index19).Int())
		var _149_v8 _dafny.Map
		_ = _149_v8
		_149_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, _141_v2)
		var _150_v9 _dafny.Map
		_ = _150_v9
		_150_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v8, _137_v0)
		_141_v2 = (_141_v2).Minus((_150_v9).Cardinality())
		var _151_v10 _dafny.Set
		_ = _151_v10
		_151_v10 = _dafny.SetOf(_141_v2)
		var _152_v11 _dafny.Sequence
		_ = _152_v11
		_152_v11 = _dafny.SeqOf(_dafny.IntOfInt64(313))
		var _153_v12 _dafny.Sequence
		_ = _153_v12
		_153_v12 = _dafny.SeqOf(_dafny.IntOfUint32((_152_v11).Cardinality()))
		var _154_v13 _dafny.Array
		_ = _154_v13
		var _nwElement0_4 _dafny.Int = _141_v2
		_ = _nwElement0_4
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
		_ = _nw26
		(_nw26).ArraySet1(_nwElement0_4, 0)
		(_nw26).ArraySet1(_dafny.IntOfInt64(77), 1)
		(_nw26).ArraySet1((_dafny.IntOfUint32(((_146_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_146_v6), 0))).Int()).(_dafny.Sequence)).Cardinality())).Times(_141_v2), 2)
		(_nw26).ArraySet1(_141_v2, 3)
		(_nw26).ArraySet1((Companion_Default___.Fm0(_145_globalState)).Times((_dafny.Zero).Minus(_141_v2)), 4)
		(_nw26).ArraySet1(_141_v2, 5)
		(_nw26).ArraySet1(Companion_Default___.Fm0(_145_globalState), 6)
		(_nw26).ArraySet1(_141_v2, 7)
		(_nw26).ArraySet1(_141_v2, 8)
		(_nw26).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(355), _141_v2), 9)
		(_nw26).ArraySet1(((_151_v10).Cardinality()).Plus(_141_v2), 10)
		(_nw26).ArraySet1(_141_v2, 11)
		(_nw26).ArraySet1(_141_v2, 12)
		(_nw26).ArraySet1(_dafny.IntOfInt64(336), 13)
		(_nw26).ArraySet1(_141_v2, 14)
		(_nw26).ArraySet1((Companion_Default___.Fm0(_145_globalState)).Minus(_141_v2), 15)
		(_nw26).ArraySet1(_141_v2, 16)
		(_nw26).ArraySet1(_141_v2, 17)
		(_nw26).ArraySet1(_141_v2, 18)
		(_nw26).ArraySet1(Companion_Default___.SafeDivisionInt(_141_v2, _141_v2), 19)
		(_nw26).ArraySet1(_141_v2, 20)
		(_nw26).ArraySet1((func() _dafny.Int {
			if _137_v0 {
				return _141_v2
			}
			return _141_v2
		})(), 21)
		(_nw26).ArraySet1(_dafny.IntOfUint32((_143_v4).Cardinality()), 22)
		(_nw26).ArraySet1(_141_v2, 23)
		(_nw26).ArraySet1(Companion_Default___.SafeDivisionInt(_141_v2, _dafny.IntOfUint32((_152_v11).Cardinality())), 24)
		(_nw26).ArraySet1(_141_v2, 25)
		(_nw26).ArraySet1(_141_v2, 26)
		(_nw26).ArraySet1(Companion_Default___.Fm0(_145_globalState), 27)
		_154_v13 = _nw26
		var _155_v14 _dafny.Map
		_ = _155_v14
		_155_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-657), _154_v13)
		var _rhs27 _dafny.Array = _138_v1
		_ = _rhs27
		var _rhs28 _dafny.Int = ((_141_v2).Times(_141_v2)).Minus(Companion_Default___.SafeDivisionInt(_141_v2, _141_v2))
		_ = _rhs28
		var _rhs29 _dafny.Array = (func() _dafny.Array {
			if (_155_v14).Contains(Companion_Default___.SafeModuloInt(_141_v2, _141_v2)) {
				return (_155_v14).Get(Companion_Default___.SafeModuloInt(_141_v2, _141_v2)).(_dafny.Array)
			}
			return _154_v13
		})()
		_ = _rhs29
		var _lhs21 *GlobalState = _145_globalState
		_ = _lhs21
		_138_v1 = _rhs27
		_lhs21.F9 = _rhs28
		_154_v13 = _rhs29
		var _156_v15 _dafny.Array
		_ = _156_v15
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw27
		_156_v15 = _nw27
		_156_v15 = _156_v15
	} else {
		var _157_v16 _dafny.CodePoint
		_ = _157_v16
		_157_v16 = _dafny.CodePoint('r')
		var _158_v17 _dafny.Map
		_ = _158_v17
		_158_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v16, _141_v2)
		(_145_globalState).F2 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v17, _137_v0)).Update(_158_v17, _137_v0)).Cardinality()
		var _159_v18 _dafny.Sequence
		_ = _159_v18
		_159_v18 = _dafny.SeqOf(_137_v0)
		_159_v18 = _dafny.SeqOf(!(_137_v0))
		var _160_v20 _dafny.MultiSet
		_ = _160_v20
		_160_v20 = _dafny.MultiSetOf(_141_v2, _141_v2, _141_v2, _141_v2, _141_v2)
		var _161_v21 _dafny.MultiSet
		_ = _161_v21
		_161_v21 = _dafny.MultiSetOf(_137_v0, _137_v0)
		var _162_v22 _dafny.Int
		_ = _162_v22
		var _163_v23 _dafny.Int
		_ = _163_v23
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		_out0, _out1 = Companion_Default___.M0(_dafny.Companion_Sequence_.IsPrefixOf(_143_v4, _143_v4), (func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_160_v20).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _164_v19 _dafny.Int
				_164_v19 = interface{}(_compr_9).(_dafny.Int)
				if (_160_v20).Contains(_164_v19) {
					_coll9.Add(Companion_Default___.SafeModuloInt(_164_v19, _141_v2), _137_v0)
				}
			}
			return _coll9.ToMap()
		}()).Update((_161_v21).Cardinality(), (_159_v18).Select((Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_159_v18).Cardinality()))).Uint32()).(bool)), _145_globalState)
		_162_v22 = _out0
		_163_v23 = _out1
		var _165_v24 _dafny.Sequence
		_ = _165_v24
		_165_v24 = _dafny.SeqOf(_143_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_166_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_167_i2 _dafny.Int) _dafny.CodePoint {
				return _166_v16
			}
		})(_157_v16))), _143_v4, _143_v4)
		if Companion_Default___.Fm1(false, (_165_v24).Select((Companion_Default___.SafeIndex(_162_v22, _dafny.IntOfUint32((_165_v24).Cardinality()))).Uint32()).(_dafny.Sequence), _145_globalState) {
			_157_v16 = _157_v16
			(_145_globalState).F3 = (Companion_Default___.Fm0(_145_globalState)).Cmp(_141_v2) > 0
			var _168_v25 _dafny.Int
			_ = _168_v25
			_168_v25 = _163_v23
			var _169_v26 _dafny.Map
			_ = _169_v26
			_169_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_168_v25), !(_137_v0))
			var _170_v27 _dafny.Int
			_ = _170_v27
			var _171_v28 _dafny.Int
			_ = _171_v28
			var _out2 _dafny.Int
			_ = _out2
			var _out3 _dafny.Int
			_ = _out3
			_out2, _out3 = Companion_Default___.M0(_137_v0, _169_v26, _145_globalState)
			_170_v27 = _out2
			_171_v28 = _out3
			(_145_globalState).F9 = _171_v28
			var _172_v29 _dafny.Map
			_ = _172_v29
			_172_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_159_v18, _dafny.SeqOf(false)), _137_v0)
			var _173_v30 _dafny.Sequence
			_ = _173_v30
			_173_v30 = _dafny.SeqOf(_170_v27)
			var _174_v31 _dafny.MultiSet
			_ = _174_v31
			_174_v31 = _dafny.MultiSetOf(_173_v30)
			(_145_globalState).F3 = !((func() bool {
				if (_172_v29).Contains(_159_v18) {
					return (_172_v29).Get(_159_v18).(bool)
				}
				return (_dafny.MultiSetOf(_173_v30, _173_v30)).IsSubsetOf(_174_v31)
			})())
		} else {
			_157_v16 = _157_v16
			_159_v18 = _dafny.Companion_Sequence_.Concatenate(_159_v18, _159_v18)
			var _175_v32 *C1
			_ = _175_v32
			var _nw28 *C1 = New_C1_()
			_ = _nw28
			_nw28.Ctor__(_157_v16)
			_175_v32 = _nw28
			var _176_v33 _dafny.Map
			_ = _176_v33
			_176_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _159_v18)
			var _177_v34 _dafny.Set
			_ = _177_v34
			_177_v34 = _dafny.SetOf(_159_v18, _159_v18, _159_v18, _159_v18, (func() _dafny.Sequence {
				if (_176_v33).Contains((_159_v18).Select((Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_159_v18).Cardinality()))).Uint32()).(bool)) {
					return (_176_v33).Get((_159_v18).Select((Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_159_v18).Cardinality()))).Uint32()).(bool)).(_dafny.Sequence)
				}
				return _159_v18
			})())
			(_145_globalState).F10 = (_177_v34).IsProperSubsetOf((_177_v34).Union(_177_v34))
			_142_v3 = (_142_v3).Update((func() bool {
				if _137_v0 {
					return _137_v0
				}
				return _137_v0
			})(), _163_v23)
		}
		var _rhs30 bool = !(!(true)) || (true)
		_ = _rhs30
		var _rhs31 bool = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_159_v18, _dafny.SeqOf(_137_v0, _137_v0, _137_v0)))).IsDisjointFrom(_161_v21)
		_ = _rhs31
		var _rhs32 bool = (_141_v2).Cmp(_162_v22) <= 0
		_ = _rhs32
		var _lhs22 *GlobalState = _145_globalState
		_ = _lhs22
		var _lhs23 *GlobalState = _145_globalState
		_ = _lhs23
		var _lhs24 *GlobalState = _145_globalState
		_ = _lhs24
		_lhs22.F10 = _rhs30
		_lhs23.F0 = _rhs31
		_lhs24.F0 = _rhs32
	}
	var _178_v35 _dafny.MultiSet
	_ = _178_v35
	_178_v35 = _dafny.MultiSetOf(_141_v2, _141_v2, _141_v2)
	var _179_v36 D6
	_ = _179_v36
	_179_v36 = Companion_D6_.Create_DC16_(_141_v2, _137_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality()), _141_v2, _137_v0)
	var _rhs33 _dafny.Int = (func() _dafny.Int {
		if (_178_v35).Contains((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_179_v36).Dtor_cf31(), _141_v2))) {
			return (_178_v35).Multiplicity((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_179_v36).Dtor_cf31(), _141_v2)))
		}
		return (_dafny.Zero).Minus(_141_v2)
	})()
	_ = _rhs33
	var _rhs34 _dafny.Int = _141_v2
	_ = _rhs34
	var _lhs25 *GlobalState = _145_globalState
	_ = _lhs25
	var _lhs26 *GlobalState = _145_globalState
	_ = _lhs26
	_lhs25.F9 = _rhs33
	_lhs26.F9 = _rhs34
	var _180_v37 _dafny.Sequence
	_ = _180_v37
	_180_v37 = _dafny.SeqOf(_dafny.IntOfInt64(-496), _141_v2, _141_v2)
	_180_v37 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_180_v37, _dafny.SeqOf(_141_v2)), _180_v37)
	var _181_v38 D1
	_ = _181_v38
	_181_v38 = Companion_D1_.Create_DC2_(false)
	var _182_v39 _dafny.MultiSet
	_ = _182_v39
	_182_v39 = _dafny.MultiSetOf(_137_v0)
	var _183_v40 _dafny.Set
	_ = _183_v40
	_183_v40 = _dafny.SetOf(_137_v0, _137_v0)
	var _pat_let_tv4 = _137_v0
	_ = _pat_let_tv4
	var _pat_let_tv5 = _137_v0
	_ = _pat_let_tv5
	var _pat_let_tv6 = _137_v0
	_ = _pat_let_tv6
	var _pat_let_tv7 = _145_globalState
	_ = _pat_let_tv7
	var _184_v41 _dafny.Array
	_ = _184_v41
	var _nwElement0_5 D1 = _181_v38
	_ = _nwElement0_5
	var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(20))
	_ = _nw29
	(_nw29).ArraySet1(_nwElement0_5, 0)
	(_nw29).ArraySet1(_181_v38, 1)
	(_nw29).ArraySet1(func(_pat_let2_0 D1) D1 {
		return func(_185_dt__update__tmp_h0 D1) D1 {
			return func(_pat_let3_0 bool) D1 {
				return func(_186_dt__update_hcf2_h0 bool) D1 {
					return Companion_D1_.Create_DC2_(_186_dt__update_hcf2_h0)
				}(_pat_let3_0)
			}(_pat_let_tv4)
		}(_pat_let2_0)
	}(_181_v38), 2)
	(_nw29).ArraySet1(_181_v38, 3)
	(_nw29).ArraySet1(_181_v38, 4)
	(_nw29).ArraySet1(_181_v38, 5)
	(_nw29).ArraySet1(Companion_D1_.Create_DC2_(_137_v0), 6)
	(_nw29).ArraySet1(Companion_D1_.Create_DC2_(_137_v0), 7)
	(_nw29).ArraySet1(_181_v38, 8)
	(_nw29).ArraySet1(_181_v38, 9)
	(_nw29).ArraySet1(_181_v38, 10)
	(_nw29).ArraySet1(_181_v38, 11)
	(_nw29).ArraySet1(func(_pat_let4_0 D1) D1 {
		return func(_187_dt__update__tmp_h1 D1) D1 {
			return func(_pat_let5_0 bool) D1 {
				return func(_188_dt__update_hcf2_h1 bool) D1 {
					return Companion_D1_.Create_DC2_(_188_dt__update_hcf2_h1)
				}(_pat_let5_0)
			}(_pat_let_tv5)
		}(_pat_let4_0)
	}(_181_v38), 12)
	(_nw29).ArraySet1(func(_pat_let6_0 D1) D1 {
		return func(_189_dt__update__tmp_h2 D1) D1 {
			return func(_pat_let7_0 bool) D1 {
				return func(_190_dt__update_hcf2_h2 bool) D1 {
					return Companion_D1_.Create_DC2_(_190_dt__update_hcf2_h2)
				}(_pat_let7_0)
			}(Companion_Default___.Fm1(_pat_let_tv6, _dafny.UnicodeSeqOfUtf8Bytes("jol"), _pat_let_tv7))
		}(_pat_let6_0)
	}(_181_v38), 13)
	(_nw29).ArraySet1(_181_v38, 14)
	(_nw29).ArraySet1(_181_v38, 15)
	(_nw29).ArraySet1(_181_v38, 16)
	(_nw29).ArraySet1(Companion_Default___.Fm5(_141_v2, (func() _dafny.Int {
		if (_182_v39).Contains(false) {
			return (_182_v39).Multiplicity(false)
		}
		return (_183_v40).Cardinality()
	})(), _145_globalState), 17)
	(_nw29).ArraySet1(_181_v38, 18)
	(_nw29).ArraySet1(_181_v38, 19)
	_184_v41 = _nw29
	var _191_v42 D2
	_ = _191_v42
	_191_v42 = Companion_D2_.Create_DC4_(_184_v41)
	_184_v41 = (func() _dafny.Array {
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_143_v4, _dafny.UnicodeSeqOfUtf8Bytes("e")) {
			return _184_v41
		}
		return (_191_v42).Dtor_cf4()
	})()
	var _192_v43 _dafny.CodePoint
	_ = _192_v43
	_192_v43 = _dafny.CodePoint('y')
	var _193_v44 *C1
	_ = _193_v44
	var _nw30 *C1 = New_C1_()
	_ = _nw30
	_nw30.Ctor__(_192_v43)
	_193_v44 = _nw30
	var _194_v45 D6
	_ = _194_v45
	_194_v45 = Companion_D6_.Create_DC14_(_193_v44)
	var _pat_let_tv8 = _137_v0
	_ = _pat_let_tv8
	var _pat_let_tv9 = _193_v44
	_ = _pat_let_tv9
	var _pat_let_tv10 = _193_v44
	_ = _pat_let_tv10
	var _source3 D6 = func(_pat_let8_0 D6) D6 {
		return func(_195_dt__update__tmp_h3 D6) D6 {
			return func(_pat_let9_0 *C1) D6 {
				return func(_196_dt__update_hcf22_h0 *C1) D6 {
					return Companion_D6_.Create_DC14_(_196_dt__update_hcf22_h0)
				}(_pat_let9_0)
			}((func() *C1 {
				if _pat_let_tv8 {
					return _pat_let_tv9
				}
				return _pat_let_tv10
			})())
		}(_pat_let8_0)
	}(_194_v45)
	_ = _source3
	if _source3.Is_DC15() {
		var _197___mcc_h0 _dafny.Int = _source3.Get_().(D6_DC15).Cf23
		_ = _197___mcc_h0
		var _198___mcc_h1 _dafny.Int = _source3.Get_().(D6_DC15).Cf24
		_ = _198___mcc_h1
		var _199___mcc_h2 _dafny.Sequence = _source3.Get_().(D6_DC15).Cf25
		_ = _199___mcc_h2
		var _200___mcc_h3 _dafny.Int = _source3.Get_().(D6_DC15).Cf26
		_ = _200___mcc_h3
		var _201___mcc_h4 _dafny.Set = _source3.Get_().(D6_DC15).Cf27
		_ = _201___mcc_h4
		var _202_cf27 _dafny.Set = _201___mcc_h4
		_ = _202_cf27
		var _203_cf26 _dafny.Int = _200___mcc_h3
		_ = _203_cf26
		var _204_cf25 _dafny.Sequence = _199___mcc_h2
		_ = _204_cf25
		var _205_cf24 _dafny.Int = _198___mcc_h1
		_ = _205_cf24
		var _206_cf23 _dafny.Int = _197___mcc_h0
		_ = _206_cf23
		var _207_v46 _dafny.MultiSet
		_ = _207_v46
		_207_v46 = _dafny.MultiSetOf((_193_v44).F20(), _192_v43, _192_v43)
		var _rhs35 _dafny.Int = _206_cf23
		_ = _rhs35
		var _rhs36 _dafny.Int = (func() _dafny.Int {
			if (_dafny.IntOfInt64(760)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_145_globalState), _206_cf23)).Cardinality()) >= 0 {
				return _dafny.IntOfInt64(446)
			}
			return (func() _dafny.Int {
				if _137_v0 {
					return (_207_v46).Cardinality()
				}
				return _206_cf23
			})()
		})()
		_ = _rhs36
		var _lhs27 *GlobalState = _145_globalState
		_ = _lhs27
		_lhs27.F9 = _rhs35
		_141_v2 = _rhs36
		var _208_v47 _dafny.Int
		_ = _208_v47
		_208_v47 = _dafny.IntOfUint32((_143_v4).Cardinality())
		var _209_v48 T0
		_ = _209_v48
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__(_208_v47, _206_cf23)
		_209_v48 = _nw31
		var _210_v49 _dafny.Array
		_ = _210_v49
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
		_ = _nw32
		_210_v49 = _nw32
		var _211_v50 *C0
		_ = _211_v50
		var _nw33 *C0 = New_C0_()
		_ = _nw33
		_nw33.Ctor__((_dafny.Zero).Minus(_203_cf26), (_209_v48).F18())
		_211_v50 = _nw33
		var _212_v51 _dafny.Map
		_ = _212_v51
		_212_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _211_v50)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_210_v49), 0))
		_ = _index20
		(_210_v49).ArraySet1((_212_v51).Merge(_212_v51), (_index20).Int())
		var _213_v52 _dafny.MultiSet
		_ = _213_v52
		_213_v52 = _dafny.MultiSetOf(_138_v1)
		var _214_v53 _dafny.Map
		_ = _214_v53
		_214_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_203_cf26, _203_cf26)
		var _215_v54 _dafny.Map
		_ = _215_v54
		_215_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _214_v53)
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_210_v49), 0))
		_ = _index21
		var _rhs37 T0 = _209_v48
		_ = _rhs37
		var _rhs38 _dafny.Sequence = Companion_Default___.Fm11(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_143_v4).Cardinality()), (_dafny.MultiSetFromSeq(_180_v37)).Cardinality()), _137_v0, (func() _dafny.Int {
			if (_213_v52).Contains(_138_v1) {
				return (_213_v52).Multiplicity(_138_v1)
			}
			return ((func() _dafny.Map {
				if (_215_v54).Contains(_137_v0) {
					return (_215_v54).Get(_137_v0).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_209_v48).F18(), _141_v2)
			})()).Cardinality()
		})(), _145_globalState)
		_ = _rhs38
		var _rhs39 *C1 = _193_v44
		_ = _rhs39
		var _rhs40 _dafny.Map = _212_v51
		_ = _rhs40
		var _lhs28 _dafny.Array = _210_v49
		_ = _lhs28
		var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_210_v49), 0))
		_ = _lhs29
		_209_v48 = _rhs37
		_204_cf25 = _rhs38
		_193_v44 = _rhs39
		(_lhs28).ArraySet1(_rhs40, (_lhs29).Int())
		var _216_v55 _dafny.Array
		_ = _216_v55
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw34
		_216_v55 = _nw34
		_216_v55 = _216_v55
		(_145_globalState).F0 = _137_v0
	} else if _source3.Is_DC16() {
		var _217___mcc_h5 _dafny.Int = _source3.Get_().(D6_DC16).Cf28
		_ = _217___mcc_h5
		var _218___mcc_h6 bool = _source3.Get_().(D6_DC16).Cf29
		_ = _218___mcc_h6
		var _219___mcc_h7 _dafny.Int = _source3.Get_().(D6_DC16).Cf30
		_ = _219___mcc_h7
		var _220___mcc_h8 _dafny.Int = _source3.Get_().(D6_DC16).Cf31
		_ = _220___mcc_h8
		var _221___mcc_h9 bool = _source3.Get_().(D6_DC16).Cf32
		_ = _221___mcc_h9
		var _222_cf32 bool = _221___mcc_h9
		_ = _222_cf32
		var _223_cf31 _dafny.Int = _220___mcc_h8
		_ = _223_cf31
		var _224_cf30 _dafny.Int = _219___mcc_h7
		_ = _224_cf30
		var _225_cf29 bool = _218___mcc_h6
		_ = _225_cf29
		var _226_cf28 _dafny.Int = _217___mcc_h5
		_ = _226_cf28
		var _227_v56 _dafny.Int
		_ = _227_v56
		var _228_v57 _dafny.MultiSet
		_ = _228_v57
		var _out4 _dafny.Int
		_ = _out4
		var _out5 _dafny.MultiSet
		_ = _out5
		_out4, _out5 = (_193_v44).M3(_137_v0, _145_globalState)
		_227_v56 = _out4
		_228_v57 = _out5
		var _229_v58 _dafny.Sequence
		_ = _229_v58
		_229_v58 = _dafny.SeqOf(_225_cf29)
		var _230_v59 _dafny.Map
		_ = _230_v59
		_230_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_cf28, Companion_Default___.Fm11(_224_cf30, _137_v0, _226_cf28, _145_globalState))
		var _231_v60 _dafny.Sequence
		_ = _231_v60
		_231_v60 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_180_v37, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-15), _dafny.IntOfUint32((_180_v37).Cardinality()))).Uint32(), _227_v56))
		var _232_v61 _dafny.Map
		_ = _232_v61
		_232_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_143_v4, _222_cf32)
		var _233_v62 _dafny.Array
		_ = _233_v62
		var _nwElement0_6 _dafny.Sequence = _180_v37
		_ = _nwElement0_6
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_6, 0)
		(_nw35).ArraySet1(_180_v37, 1)
		(_nw35).ArraySet1(_180_v37, 2)
		(_nw35).ArraySet1(_180_v37, 3)
		(_nw35).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_180_v37, _180_v37), 4)
		(_nw35).ArraySet1(_180_v37, 5)
		(_nw35).ArraySet1(_dafny.SeqOf(Companion_Default___.Fm0(_145_globalState), _dafny.IntOfInt64(295)), 6)
		(_nw35).ArraySet1(_dafny.SeqOf(_223_cf31, _dafny.IntOfInt64(173), (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_229_v58, (Companion_Default___.SafeIndex(_226_cf28, _dafny.IntOfUint32((_229_v58).Cardinality()))).Uint32(), _137_v0))).Cardinality(), _dafny.IntOfInt64(3)), 7)
		(_nw35).ArraySet1(_dafny.Companion_Sequence_.Update(_180_v37, (Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_180_v37).Cardinality()))).Uint32(), _226_cf28), 8)
		(_nw35).ArraySet1((func() _dafny.Sequence {
			if (_230_v59).Contains(_226_cf28) {
				return (_230_v59).Get(_226_cf28).(_dafny.Sequence)
			}
			return _180_v37
		})(), 9)
		(_nw35).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_180_v37, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(151))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_234_v40 _dafny.Set) func(_dafny.Int) _dafny.Int {
			return func(_235_i3 _dafny.Int) _dafny.Int {
				return (_234_v40).Cardinality()
			}
		})(_183_v40)))), 10)
		(_nw35).ArraySet1(_180_v37, 11)
		(_nw35).ArraySet1(_dafny.Companion_Sequence_.Update((_231_v60).Select((Companion_Default___.SafeIndex(_223_cf31, _dafny.IntOfUint32((_231_v60).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_232_v61).Cardinality(), _dafny.IntOfUint32(((_231_v60).Select((Companion_Default___.SafeIndex(_223_cf31, _dafny.IntOfUint32((_231_v60).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _224_cf30), 12)
		_233_v62 = _nw35
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_233_v62), 0))
		_ = _index22
		(_233_v62).ArraySet1(_180_v37, (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_233_v62), 0))
		_ = _index23
		(_233_v62).ArraySet1(_180_v37, (_index23).Int())
		var _236_v63 _dafny.Array
		_ = _236_v63
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw36
		_236_v63 = _nw36
		var _237_v64 D8
		_ = _237_v64
		_237_v64 = Companion_D8_.Create_DC23_(_226_cf28, _236_v63, _223_cf31)
		var _238_v65 _dafny.Set
		_ = _238_v65
		_238_v65 = _dafny.SetOf(_226_cf28)
		var _239_v66 _dafny.Map
		_ = _239_v66
		_239_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, _238_v65)
		var _240_v67 _dafny.Sequence
		_ = _240_v67
		_240_v67 = _dafny.SeqOf(_238_v65, (func() _dafny.Set {
			if (_239_v66).Contains((_144_v5).Cardinality()) {
				return (_239_v66).Get((_144_v5).Cardinality()).(_dafny.Set)
			}
			return _238_v65
		})())
		var _source4 D5 = Companion_Default___.Fm12((_237_v64).Dtor_cf46(), _240_v67, _224_cf30, _223_cf31, _145_globalState)
		_ = _source4
		if _source4.Is_DC13() {
			var _241___mcc_h15 _dafny.Int = _source4.Get_().(D5_DC13).Cf19
			_ = _241___mcc_h15
			var _242___mcc_h16 _dafny.Int = _source4.Get_().(D5_DC13).Cf20
			_ = _242___mcc_h16
			var _243___mcc_h17 bool = _source4.Get_().(D5_DC13).Cf21
			_ = _243___mcc_h17
			var _244_cf21 bool = _243___mcc_h17
			_ = _244_cf21
			var _245_cf20 _dafny.Int = _242___mcc_h16
			_ = _245_cf20
			var _246_cf19 _dafny.Int = _241___mcc_h15
			_ = _246_cf19
			(_145_globalState).F2 = _245_cf20
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_236_v63), 0))
			_ = _index24
			(_236_v63).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _142_v3)).Cardinality(), (_index24).Int())
			var _247_v68 _dafny.Array
			_ = _247_v68
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
			_ = _nw37
			_247_v68 = _nw37
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_247_v68), 0))
			_ = _index25
			(_247_v68).ArraySet1CodePoint((_193_v44).F20(), (_index25).Int())
			var _248_v69 _dafny.Map
			_ = _248_v69
			_248_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_cf31, _dafny.IntOfUint32((_229_v58).Cardinality()))
			var _249_v70 _dafny.Map
			_ = _249_v70
			_249_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v69, _143_v4)
			var _250_v71 _dafny.Map
			_ = _250_v71
			_250_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_cf20, _248_v69)
			var _251_v72 D5
			_ = _251_v72
			_251_v72 = Companion_D5_.Create_DC12_(_143_v4)
			var _252_v73 _dafny.Array
			_ = _252_v73
			var _nwElement0_7 _dafny.Sequence = _143_v4
			_ = _nwElement0_7
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(28))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_7, 0)
			(_nw38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sfvx"), _dafny.UnicodeSeqOfUtf8Bytes("tunb")), 1)
			(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("laijoy"), 2)
			(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("l"), 3)
			(_nw38).ArraySet1(_143_v4, 4)
			(_nw38).ArraySet1(_143_v4, 5)
			(_nw38).ArraySet1((func() _dafny.Sequence {
				if (_249_v70).Contains((func() _dafny.Map {
					if (_250_v71).Contains(_dafny.IntOfUint32(((_233_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_233_v62), 0))).Int()).(_dafny.Sequence)).Cardinality())) {
						return (_250_v71).Get(_dafny.IntOfUint32(((_233_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_233_v62), 0))).Int()).(_dafny.Sequence)).Cardinality())).(_dafny.Map)
					}
					return _248_v69
				})()) {
					return (_249_v70).Get((func() _dafny.Map {
						if (_250_v71).Contains(_dafny.IntOfUint32(((_233_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_233_v62), 0))).Int()).(_dafny.Sequence)).Cardinality())) {
							return (_250_v71).Get(_dafny.IntOfUint32(((_233_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_233_v62), 0))).Int()).(_dafny.Sequence)).Cardinality())).(_dafny.Map)
						}
						return _248_v69
					})()).(_dafny.Sequence)
				}
				return _143_v4
			})(), 6)
			(_nw38).ArraySet1(_dafny.Companion_Sequence_.Update(_143_v4, (Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_143_v4).Cardinality()))).Uint32(), (_193_v44).F20()), 7)
			(_nw38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_143_v4, _143_v4), 8)
			(_nw38).ArraySet1(_dafny.Companion_Sequence_.Update((_251_v72).Dtor_cf18(), (Companion_Default___.SafeIndex(_227_v56, _dafny.IntOfUint32(((_251_v72).Dtor_cf18()).Cardinality()))).Uint32(), _192_v43), 9)
			(_nw38).ArraySet1(_143_v4, 10)
			(_nw38).ArraySet1((func() _dafny.Sequence {
				if _225_cf29 {
					return _143_v4
				}
				return _143_v4
			})(), 11)
			(_nw38).ArraySet1((func() _dafny.Sequence {
				if Companion_Default___.Fm1(_225_cf29, _143_v4, _145_globalState) {
					return _dafny.Companion_Sequence_.Update(Companion_Default___.Fm6(_222_cf32, _145_globalState), (Companion_Default___.SafeIndex(_227_v56, _dafny.IntOfUint32((Companion_Default___.Fm6(_222_cf32, _145_globalState)).Cardinality()))).Uint32(), (_193_v44).F20())
				}
				return _143_v4
			})(), 12)
			(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vrxbpbgci"), 13)
			(_nw38).ArraySet1(_143_v4, 14)
			(_nw38).ArraySet1(_143_v4, 15)
			(_nw38).ArraySet1(_143_v4, 16)
			(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uci"), 17)
			(_nw38).ArraySet1(_143_v4, 18)
			(_nw38).ArraySet1(_143_v4, 19)
			(_nw38).ArraySet1((Companion_D5_.Create_DC12_(_143_v4)).Dtor_cf18(), 20)
			(_nw38).ArraySet1(_143_v4, 21)
			(_nw38).ArraySet1(_143_v4, 22)
			(_nw38).ArraySet1(_143_v4, 23)
			(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("smf"), 24)
			(_nw38).ArraySet1(_143_v4, 25)
			(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hrnoni"), 26)
			(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("abfgyma"), 27)
			_252_v73 = _nw38
			var _253_v74 _dafny.Set
			_ = _253_v74
			_253_v74 = _dafny.SetOf(_229_v58, _dafny.SeqOf(_225_cf29, _137_v0, true))
			var _254_v75 D9
			_ = _254_v75
			_254_v75 = Companion_D9_.Create_DC24_(_252_v73)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_236_v63), 0))
			_ = _index26
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_247_v68), 0))
			_ = _index27
			var _rhs41 _dafny.Int = (func() _dafny.Int {
				if (_142_v3).Contains((_253_v74).IsProperSubsetOf(_253_v74)) {
					return (_142_v3).Get((_253_v74).IsProperSubsetOf(_253_v74)).(_dafny.Int)
				}
				return _224_cf30
			})()
			_ = _rhs41
			var _rhs42 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(811)).Minus((_183_v40).Cardinality()), _226_cf28)
			_ = _rhs42
			var _rhs43 bool = _137_v0
			_ = _rhs43
			var _rhs44 _dafny.CodePoint = (_193_v44).F20()
			_ = _rhs44
			var _rhs45 _dafny.Array = (_254_v75).Dtor_cf49()
			_ = _rhs45
			var _lhs30 _dafny.Array = _236_v63
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_236_v63), 0))
			_ = _lhs31
			var _lhs32 *GlobalState = _145_globalState
			_ = _lhs32
			var _lhs33 *GlobalState = _145_globalState
			_ = _lhs33
			var _lhs34 _dafny.Array = _247_v68
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_247_v68), 0))
			_ = _lhs35
			(_lhs30).ArraySet1(_rhs41, (_lhs31).Int())
			_lhs32.F2 = _rhs42
			_lhs33.F3 = _rhs43
			(_lhs34).ArraySet1CodePoint(_rhs44, (_lhs35).Int())
			_252_v73 = _rhs45
			var _255_v76 *C2
			_ = _255_v76
			var _nw39 *C2 = New_C2_()
			_ = _nw39
			_nw39.Ctor__()
			_255_v76 = _nw39
			var _256_v77 D8
			_ = _256_v77
			_256_v77 = Companion_D8_.Create_DC22_(_255_v76)
			var _257_v78 _dafny.Int
			_ = _257_v78
			_257_v78 = _227_v56
			var _258_v79 *C0
			_ = _258_v79
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__(_257_v78, _224_cf30)
			_258_v79 = _nw40
			var _rhs46 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_229_v58, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_141_v2, (_236_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_236_v63), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_229_v58).Cardinality()))).Uint32(), true)
			_ = _rhs46
			var _rhs47 bool = _222_cf32
			_ = _rhs47
			var _rhs48 D8 = _256_v77
			_ = _rhs48
			var _rhs49 *C0 = _258_v79
			_ = _rhs49
			var _lhs36 *GlobalState = _145_globalState
			_ = _lhs36
			_229_v58 = _rhs46
			_lhs36.F10 = _rhs47
			_256_v77 = _rhs48
			_258_v79 = _rhs49
			var _259_v80 D4
			_ = _259_v80
			_259_v80 = Companion_D4_.Create_DC9_(_258_v79)
			_259_v80 = _259_v80
		} else {
			var _260___mcc_h18 _dafny.Sequence = _source4.Get_().(D5_DC12).Cf18
			_ = _260___mcc_h18
			var _261_cf18 _dafny.Sequence = _260___mcc_h18
			_ = _261_cf18
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_233_v62), 0))
			_ = _index28
			(_233_v62).ArraySet1(_180_v37, (_index28).Int())
			var _262_v81 _dafny.Map
			_ = _262_v81
			_262_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-665), _225_cf29)
			var _263_v82 _dafny.Map
			_ = _263_v82
			_263_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v81, _141_v2)
			var _264_v84 _dafny.MultiSet
			_ = _264_v84
			_264_v84 = _dafny.MultiSetOf(_262_v81)
			var _265_v86 D10
			_ = _265_v86
			_265_v86 = Companion_D10_.Create_DC28_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v81, (_dafny.Zero).Minus(_224_cf30)))
			var _266_v88 _dafny.Array
			_ = _266_v88
			var _nwElement0_8 _dafny.Map = Companion_Default___.Fm13(_145_globalState)
			_ = _nwElement0_8
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(27))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_8, 0)
			(_nw41).ArraySet1((_263_v82).Merge(_263_v82), 1)
			(_nw41).ArraySet1(_263_v82, 2)
			(_nw41).ArraySet1(_263_v82, 3)
			(_nw41).ArraySet1(_263_v82, 4)
			(_nw41).ArraySet1((_263_v82).Merge(_263_v82), 5)
			(_nw41).ArraySet1((func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter10 := _dafny.Iterate((_264_v84).Elements()); ; {
					_compr_10, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _267_v83 _dafny.Map
					_267_v83 = interface{}(_compr_10).(_dafny.Map)
					if (_264_v84).Contains(_267_v83) {
						_coll10.Add(_267_v83, _226_cf28)
					}
				}
				return _coll10.ToMap()
			}()).Update(_262_v81, (Companion_Default___.Fm8(_226_cf28, _dafny.IntOfInt64(-844), _145_globalState)).Cardinality()), 6)
			(_nw41).ArraySet1((_263_v82).Merge(_263_v82), 7)
			(_nw41).ArraySet1(_263_v82, 8)
			(_nw41).ArraySet1((func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_264_v84).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _268_v85 _dafny.Map
					_268_v85 = interface{}(_compr_11).(_dafny.Map)
					if (_264_v84).Contains(_268_v85) {
						_coll11.Add(_268_v85, _224_cf30)
					}
				}
				return _coll11.ToMap()
			}()).Update(_262_v81, _223_cf31), 9)
			(_nw41).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v81, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uavoowfky")).Cardinality())), 10)
			(_nw41).ArraySet1(_263_v82, 11)
			(_nw41).ArraySet1((_263_v82).Merge(_263_v82), 12)
			(_nw41).ArraySet1(((_265_v86).Dtor_cf56()).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v56, _137_v0), _141_v2), 13)
			(_nw41).ArraySet1(_263_v82, 14)
			(_nw41).ArraySet1(_263_v82, 15)
			(_nw41).ArraySet1(_263_v82, 16)
			(_nw41).ArraySet1(_263_v82, 17)
			(_nw41).ArraySet1((_263_v82).Merge(_263_v82), 18)
			(_nw41).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_238_v65).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _269_v87 _dafny.Int
					_269_v87 = interface{}(_compr_12).(_dafny.Int)
					if (_238_v65).Contains(_269_v87) {
						_coll12.Add((_269_v87).Minus(_224_cf30), _222_cf32)
					}
				}
				return _coll12.ToMap()
			}(), _224_cf30), 19)
			(_nw41).ArraySet1(_263_v82, 20)
			(_nw41).ArraySet1(_263_v82, 21)
			(_nw41).ArraySet1(_263_v82, 22)
			(_nw41).ArraySet1(_263_v82, 23)
			(_nw41).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_cf31, _222_cf32), _226_cf28), 24)
			(_nw41).ArraySet1((_263_v82).Merge(_263_v82), 25)
			(_nw41).ArraySet1((_263_v82).Merge(_263_v82), 26)
			_266_v88 = _nw41
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_266_v88), 0))
			_ = _index29
			(_266_v88).ArraySet1(_263_v82, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_266_v88), 0))
			_ = _index30
			var _rhs50 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v81, _141_v2)).Merge(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_263_v82).Keys().Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _270_v89 _dafny.Map
					_270_v89 = interface{}(_compr_13).(_dafny.Map)
					if (_263_v82).Contains(_270_v89) {
						_coll13.Add(_270_v89, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_143_v4, _dafny.IntOfInt64(-269))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())
					}
				}
				return _coll13.ToMap()
			}())).Merge(_263_v82)
			_ = _rhs50
			var _rhs51 _dafny.Int = _226_cf28
			_ = _rhs51
			var _lhs37 _dafny.Array = _266_v88
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_266_v88), 0))
			_ = _lhs38
			(_lhs37).ArraySet1(_rhs50, (_lhs38).Int())
			_223_cf31 = _rhs51
			var _271_v90 D2
			_ = _271_v90
			_271_v90 = Companion_D2_.Create_DC6_(_dafny.IntOfInt64(-696), _222_cf32, _222_cf32, _227_v56)
			_271_v90 = _271_v90
			_226_cf28 = (_141_v2).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v1, _236_v63)).Cardinality())
		}
		var _272_v91 *C2
		_ = _272_v91
		var _nw42 *C2 = New_C2_()
		_ = _nw42
		_nw42.Ctor__()
		_272_v91 = _nw42
		_272_v91 = _272_v91
	} else if _source3.Is_DC17() {
		var _273___mcc_h10 _dafny.CodePoint = _source3.Get_().(D6_DC17).Cf33
		_ = _273___mcc_h10
		var _274___mcc_h11 bool = _source3.Get_().(D6_DC17).Cf34
		_ = _274___mcc_h11
		var _275___mcc_h12 bool = _source3.Get_().(D6_DC17).Cf35
		_ = _275___mcc_h12
		var _276___mcc_h13 _dafny.CodePoint = _source3.Get_().(D6_DC17).Cf36
		_ = _276___mcc_h13
		var _277_cf36 _dafny.CodePoint = _276___mcc_h13
		_ = _277_cf36
		var _278_cf35 bool = _275___mcc_h12
		_ = _278_cf35
		var _279_cf34 bool = _274___mcc_h11
		_ = _279_cf34
		var _280_cf33 _dafny.CodePoint = _273___mcc_h10
		_ = _280_cf33
		(_145_globalState).F2 = (_141_v2).Times((_dafny.IntOfInt64(-556)).Plus(_141_v2))
		_279_cf34 = false
		var _281_v92 _dafny.Sequence
		_ = _281_v92
		_281_v92 = _dafny.SeqOf(_180_v37, _180_v37, _dafny.SeqOf(_dafny.IntOfUint32((_143_v4).Cardinality())), _180_v37)
		(_145_globalState).F10 = !(_dafny.Companion_Sequence_.IsProperPrefixOf((_281_v92).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-944), _dafny.IntOfUint32((_281_v92).Cardinality()))).Uint32()).(_dafny.Sequence), _180_v37))
		(_145_globalState).F0 = true
	} else {
		var _282___mcc_h14 *C1 = _source3.Get_().(D6_DC14).Cf22
		_ = _282___mcc_h14
		var _283_cf22 *C1 = _282___mcc_h14
		_ = _283_cf22
		(_145_globalState).F0 = true
		(_145_globalState).F9 = _141_v2
		var _284_v93 D9
		_ = _284_v93
		_284_v93 = Companion_D9_.Create_DC27_(_141_v2)
		var _pat_let_tv11 = _141_v2
		_ = _pat_let_tv11
		if !(func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_dafny.SeqOf(_141_v2)).Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _285_v94 _dafny.Int
				_285_v94 = interface{}(_compr_14).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_141_v2), _285_v94) {
					_coll14.Add((_285_v94).Times(_141_v2), _141_v2)
				}
			}
			return _coll14.ToMap()
		}()).Contains((_141_v2).Plus((func(_pat_let10_0 D9) D9 {
			return func(_286_dt__update__tmp_h4 D9) D9 {
				return func(_pat_let11_0 _dafny.Int) D9 {
					return func(_287_dt__update_hcf55_h0 _dafny.Int) D9 {
						return Companion_D9_.Create_DC27_(_287_dt__update_hcf55_h0)
					}(_pat_let11_0)
				}(_pat_let_tv11)
			}(_pat_let10_0)
		}(_284_v93)).Dtor_cf55())) {
			var _288_v95 _dafny.Sequence
			_ = _288_v95
			_288_v95 = _dafny.SeqOf(_143_v4, _143_v4, _143_v4, _143_v4, _dafny.UnicodeSeqOfUtf8Bytes("w"))
			var _289_v96 _dafny.Map
			_ = _289_v96
			_289_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_288_v95).Select((Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_288_v95).Cardinality()))).Uint32()).(_dafny.Sequence), !(_137_v0))
			_289_v96 = _289_v96
			var _290_v97 _dafny.Int
			_ = _290_v97
			var _291_v98 _dafny.MultiSet
			_ = _291_v98
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.MultiSet
			_ = _out7
			_out6, _out7 = (_283_cf22).M3(_137_v0, _145_globalState)
			_290_v97 = _out6
			_291_v98 = _out7
			var _292_v99 _dafny.Int
			_ = _292_v99
			_292_v99 = (_dafny.Zero).Minus(_290_v97)
			var _293_v100 T0
			_ = _293_v100
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__(_292_v99, _dafny.IntOfUint32((_143_v4).Cardinality()))
			_293_v100 = _nw43
			_293_v100 = _293_v100
			(_145_globalState).F0 = _137_v0
			var _294_v101 _dafny.Sequence
			_ = _294_v101
			_294_v101 = _dafny.SeqOf(_137_v0)
			(_145_globalState).F3 = _dafny.Companion_Sequence_.IsPrefixOf(_294_v101, _dafny.SeqOf(_137_v0))
		} else {
			var _295_v102 *C2
			_ = _295_v102
			var _nw44 *C2 = New_C2_()
			_ = _nw44
			_nw44.Ctor__()
			_295_v102 = _nw44
			var _nw45 *C2 = New_C2_()
			_ = _nw45
			_nw45.Ctor__()
			_295_v102 = _nw45
			var _296_v103 *C1
			_ = _296_v103
			var _nw46 *C1 = New_C1_()
			_ = _nw46
			_nw46.Ctor__((_283_cf22).F20())
			_296_v103 = _nw46
			var _297_v104 _dafny.Int
			_ = _297_v104
			var _298_v105 _dafny.MultiSet
			_ = _298_v105
			var _out8 _dafny.Int
			_ = _out8
			var _out9 _dafny.MultiSet
			_ = _out9
			_out8, _out9 = (_283_cf22).M3(!(_137_v0), _145_globalState)
			_297_v104 = _out8
			_298_v105 = _out9
			(_145_globalState).F10 = _137_v0
			(_145_globalState).F9 = _141_v2
		}
		var _299_v106 _dafny.Array
		_ = _299_v106
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_6
		var _nw47 _dafny.Array
		_ = _nw47
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw47 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = (func(_300_v0 bool) func(_dafny.Int) _dafny.Int {
				return func(_301_i4 _dafny.Int) _dafny.Int {
					return (_301_i4).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_300_v0)).Cardinality()))
				}
			})(_137_v0)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw47 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw47).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw47).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_299_v106 = _nw47
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_299_v106), 0))
		_ = _index31
		(_299_v106).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(170), _141_v2), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_299_v106), 0))
		_ = _index32
		(_299_v106).ArraySet1((_183_v40).Cardinality(), (_index32).Int())
	}
	var _302_v107 D9
	_ = _302_v107
	_302_v107 = Companion_D9_.Create_DC25_(_141_v2)
	var _pat_let_tv12 = _141_v2
	_ = _pat_let_tv12
	var _303_v108 _dafny.Array
	_ = _303_v108
	var _nwElement0_9 D9 = _302_v107
	_ = _nwElement0_9
	var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(6))
	_ = _nw48
	(_nw48).ArraySet1(_nwElement0_9, 0)
	(_nw48).ArraySet1(func(_pat_let12_0 D9) D9 {
		return func(_304_dt__update__tmp_h5 D9) D9 {
			return func(_pat_let13_0 _dafny.Int) D9 {
				return func(_305_dt__update_hcf50_h0 _dafny.Int) D9 {
					return Companion_D9_.Create_DC25_(_305_dt__update_hcf50_h0)
				}(_pat_let13_0)
			}(_pat_let_tv12)
		}(_pat_let12_0)
	}(_302_v107), 1)
	(_nw48).ArraySet1(_302_v107, 2)
	(_nw48).ArraySet1((func() D9 {
		if _137_v0 {
			return _302_v107
		}
		return _302_v107
	})(), 3)
	(_nw48).ArraySet1(_302_v107, 4)
	(_nw48).ArraySet1(_302_v107, 5)
	_303_v108 = _nw48
	var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_303_v108), 0))
	_ = _index33
	(_303_v108).ArraySet1(_302_v107, (_index33).Int())
	var _pat_let_tv13 = _143_v4
	_ = _pat_let_tv13
	var _pat_let_tv14 = _143_v4
	_ = _pat_let_tv14
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_303_v108), 0))
	_ = _index34
	(_303_v108).ArraySet1(func(_pat_let14_0 D9) D9 {
		return func(_306_dt__update__tmp_h6 D9) D9 {
			return func(_pat_let15_0 _dafny.Int) D9 {
				return func(_307_dt__update_hcf50_h1 _dafny.Int) D9 {
					return Companion_D9_.Create_DC25_(_307_dt__update_hcf50_h1)
				}(_pat_let15_0)
			}(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_pat_let_tv13, _pat_let_tv14)).Cardinality()))
		}(_pat_let14_0)
	}(_302_v107), (_index34).Int())
	if (Companion_Default___.SafeDivisionInt(_141_v2, _141_v2)).Cmp(_dafny.IntOfUint32(((func() _dafny.Sequence {
		if _137_v0 {
			return _dafny.SeqOf(_141_v2)
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(453))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_308_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_309_i5 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_308_v2)
			}
		})(_141_v2)))
	})()).Cardinality())) != 0 {
		var _310_v109 _dafny.Map
		_ = _310_v109
		_310_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_193_v44).F20(), _137_v0)
		var _rhs52 _dafny.Map = _310_v109
		_ = _rhs52
		var _rhs53 bool = _137_v0
		_ = _rhs53
		var _lhs39 *GlobalState = _145_globalState
		_ = _lhs39
		_310_v109 = _rhs52
		_lhs39.F0 = _rhs53
		var _311_v110 _dafny.Sequence
		_ = _311_v110
		_311_v110 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_143_v4, _143_v4), _dafny.UnicodeSeqOfUtf8Bytes("k"), _143_v4)
		var _rhs54 _dafny.Sequence = _180_v37
		_ = _rhs54
		var _rhs55 bool = (func() bool {
			if (func() bool {
				if _137_v0 {
					return _137_v0
				}
				return false
			})() {
				return _137_v0
			}
			return (_137_v0) && (_137_v0)
		})()
		_ = _rhs55
		var _rhs56 _dafny.Sequence = _311_v110
		_ = _rhs56
		var _lhs40 *GlobalState = _145_globalState
		_ = _lhs40
		_180_v37 = _rhs54
		_lhs40.F0 = _rhs55
		_311_v110 = _rhs56
		_143_v4 = _143_v4
		var _312_v111 _dafny.Map
		_ = _312_v111
		_312_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, _192_v43)
		var _313_v112 _dafny.Sequence
		_ = _313_v112
		_313_v112 = _dafny.SeqOf(_137_v0, true, !(_137_v0), _137_v0)
		_183_v40 = Companion_Default___.Fm14((_141_v2).Cmp(_141_v2) > 0, (func() _dafny.CodePoint {
			if (_312_v111).Contains(_dafny.IntOfUint32((_313_v112).Cardinality())) {
				return (_312_v111).Get(_dafny.IntOfUint32((_313_v112).Cardinality())).(_dafny.CodePoint)
			}
			return Companion_Default___.Fm7((_193_v44).F20(), _137_v0, _145_globalState)
		})(), _137_v0, _145_globalState)
		if !(Companion_Default___.Fm14(_137_v0, _dafny.CodePoint('c'), !(_137_v0), _145_globalState)).Equals(_183_v40) {
			var _314_v113 D2
			_ = _314_v113
			_314_v113 = Companion_D2_.Create_DC4_(_184_v41)
			var _315_v114 D2
			_ = _315_v114
			_315_v114 = Companion_D2_.Create_DC7_(_314_v113)
			_315_v114 = _315_v114
			var _316_v115 _dafny.Array
			_ = _316_v115
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_7
			var _nw49 _dafny.Array
			_ = _nw49
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw49 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Map = (func(_317_v109 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_318_i6 _dafny.Int) _dafny.Map {
						return _317_v109
					}
				})(_310_v109)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw49 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw49).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw49).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_316_v115 = _nw49
			var _rhs57 _dafny.Int = _dafny.IntOfUint32((_143_v4).Cardinality())
			_ = _rhs57
			var _rhs58 _dafny.Array = _138_v1
			_ = _rhs58
			var _rhs59 _dafny.Array = _316_v115
			_ = _rhs59
			var _lhs41 *GlobalState = _145_globalState
			_ = _lhs41
			var _lhs42 *GlobalState = _145_globalState
			_ = _lhs42
			_lhs41.F9 = _rhs57
			_lhs42.F11 = _rhs58
			_316_v115 = _rhs59
			var _rhs60 bool = (_137_v0) == ((_141_v2).Cmp(_dafny.IntOfInt64(721)) > 0)
			_ = _rhs60
			var _rhs61 _dafny.CodePoint = Companion_Default___.Fm7((_193_v44).F20(), _137_v0, _145_globalState)
			_ = _rhs61
			_137_v0 = _rhs60
			_192_v43 = _rhs61
			var _319_v116 _dafny.Map
			_ = _319_v116
			_319_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(Companion_Default___.Fm1(_137_v0, _143_v4, _145_globalState), Companion_Default___.Fm6(_137_v0, _145_globalState), _145_globalState), _dafny.Companion_Sequence_.Concatenate(_143_v4, _dafny.UnicodeSeqOfUtf8Bytes("eph")))
			_319_v116 = (_319_v116).Update(_137_v0, _dafny.UnicodeSeqOfUtf8Bytes("aaf"))
			(_145_globalState).F10 = !((func() bool {
				if !(_137_v0) || (_137_v0) {
					return (_137_v0) == (_137_v0)
				}
				return _137_v0
			})())
		} else {
			(_145_globalState).F10 = _137_v0
			(_145_globalState).F2 = Companion_Default___.SafeModuloInt((Companion_Default___.Fm8(_141_v2, _dafny.IntOfInt64(873), _145_globalState)).Cardinality(), _141_v2)
			var _320_v117 *C2
			_ = _320_v117
			var _nw50 *C2 = New_C2_()
			_ = _nw50
			_nw50.Ctor__()
			_320_v117 = _nw50
			var _321_v118 _dafny.MultiSet
			_ = _321_v118
			_321_v118 = _dafny.MultiSetOf(_320_v117)
			(_145_globalState).F0 = (_dafny.MultiSetOf(_320_v117)).IsSubsetOf((_321_v118).Difference(_321_v118))
			var _322_v119 _dafny.Map
			_ = _322_v119
			_322_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm11(_141_v2, _137_v0, _141_v2, _145_globalState), ((_dafny.Zero).Minus(_141_v2)).Cmp(_141_v2) >= 0)
			_322_v119 = _322_v119
			var _323_v120 _dafny.Array
			_ = _323_v120
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_8
			var _nw51 _dafny.Array
			_ = _nw51
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw51 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = func(_324_i7 _dafny.Int) _dafny.Int {
					return (_324_i7).Minus((_dafny.SetOf(_dafny.CodePoint('n'))).Cardinality())
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw51 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw51).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw51).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_323_v120 = _nw51
			var _325_v121 D6
			_ = _325_v121
			_325_v121 = Companion_D6_.Create_DC15_(_dafny.IntOfUint32((_143_v4).Cardinality()), _141_v2, _180_v37, _141_v2, _dafny.SetOf(_141_v2))
			var _326_v122 _dafny.Int
			_ = _326_v122
			_326_v122 = (_325_v121).Dtor_cf23()
			var _327_v123 *C0
			_ = _327_v123
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__(_326_v122, _dafny.IntOfUint32((_143_v4).Cardinality()))
			_327_v123 = _nw52
			var _328_v124 _dafny.Map
			_ = _328_v124
			_328_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v123, _323_v120)
			var _329_v125 _dafny.Array
			_ = _329_v125
			var _nwElement0_10 _dafny.Array = _323_v120
			_ = _nwElement0_10
			var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(25))
			_ = _nw53
			(_nw53).ArraySet1(_nwElement0_10, 0)
			(_nw53).ArraySet1((func() _dafny.Array {
				if (_328_v124).Contains(_327_v123) {
					return (_328_v124).Get(_327_v123).(_dafny.Array)
				}
				return _323_v120
			})(), 1)
			(_nw53).ArraySet1(_323_v120, 2)
			(_nw53).ArraySet1(_323_v120, 3)
			(_nw53).ArraySet1(_323_v120, 4)
			(_nw53).ArraySet1(_323_v120, 5)
			(_nw53).ArraySet1(_323_v120, 6)
			(_nw53).ArraySet1(_323_v120, 7)
			(_nw53).ArraySet1(_323_v120, 8)
			(_nw53).ArraySet1(_323_v120, 9)
			(_nw53).ArraySet1(_323_v120, 10)
			(_nw53).ArraySet1(_323_v120, 11)
			(_nw53).ArraySet1(_323_v120, 12)
			(_nw53).ArraySet1(_323_v120, 13)
			(_nw53).ArraySet1(_323_v120, 14)
			(_nw53).ArraySet1(_323_v120, 15)
			(_nw53).ArraySet1(_323_v120, 16)
			(_nw53).ArraySet1(_323_v120, 17)
			(_nw53).ArraySet1(_323_v120, 18)
			(_nw53).ArraySet1(_323_v120, 19)
			(_nw53).ArraySet1(_323_v120, 20)
			(_nw53).ArraySet1(_323_v120, 21)
			(_nw53).ArraySet1(_323_v120, 22)
			(_nw53).ArraySet1(_323_v120, 23)
			(_nw53).ArraySet1(_323_v120, 24)
			_329_v125 = _nw53
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_329_v125), 0))
			_ = _index35
			(_329_v125).ArraySet1(_323_v120, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_329_v125), 0))
			_ = _index36
			(_329_v125).ArraySet1(_323_v120, (_index36).Int())
		}
	} else {
		(_145_globalState).F3 = _137_v0
		var _330_v126 _dafny.Map
		_ = _330_v126
		_330_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_141_v2, _dafny.IntOfInt64(681)), (_182_v39).Difference(_182_v39))
		var _331_v127 _dafny.Set
		_ = _331_v127
		_331_v127 = _dafny.SetOf(_141_v2)
		_330_v126 = Companion_Default___.Fm15(((_331_v127).Cardinality()).Cmp(_141_v2) == 0, (_142_v3).Update(true, _dafny.IntOfInt64(-178)), _137_v0, false, _145_globalState)
		var _332_v128 _dafny.Map
		_ = _332_v128
		_332_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, (Companion_D5_.Create_DC12_(_143_v4)).Dtor_cf18())
		_143_v4 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_332_v128).Contains(_141_v2) {
				return (_332_v128).Get(_141_v2).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Update(_143_v4, (Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_143_v4).Cardinality()))).Uint32(), (_193_v44).F20())
		})(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kr")).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_332_v128).Contains(_141_v2) {
				return (_332_v128).Get(_141_v2).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Update(_143_v4, (Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_143_v4).Cardinality()))).Uint32(), (_193_v44).F20())
		})()).Cardinality()))).Uint32(), _192_v43)
		(_145_globalState).F3 = _137_v0
		var _333_v129 _dafny.Array
		_ = _333_v129
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw54
		_333_v129 = _nw54
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_333_v129), 0))
		_ = _index37
		(_333_v129).ArraySet1(_138_v1, (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_333_v129), 0))
		_ = _index38
		(_333_v129).ArraySet1(_138_v1, (_index38).Int())
	}
	var _334_v130 _dafny.Set
	_ = _334_v130
	_334_v130 = _dafny.SetOf(_193_v44, _193_v44, _193_v44)
	(_145_globalState).F3 = (_334_v130).Contains(_193_v44)
	var _335_v131 _dafny.Set
	_ = _335_v131
	_335_v131 = _dafny.SetOf(_dafny.IntOfUint32((_143_v4).Cardinality()))
	if (_dafny.SetOf(_141_v2)).IsDisjointFrom(_335_v131) {
		var _336_v132 bool
		_ = _336_v132
		var _337_v133 _dafny.Int
		_ = _337_v133
		var _338_v134 bool
		_ = _338_v134
		var _339_v135 bool
		_ = _339_v135
		var _out10 bool
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		var _out12 bool
		_ = _out12
		var _out13 bool
		_ = _out13
		_out10, _out11, _out12, _out13 = (_193_v44).M4(_145_globalState)
		_336_v132 = _out10
		_337_v133 = _out11
		_338_v134 = _out12
		_339_v135 = _out13
		var _340_v136 _dafny.Int
		_ = _340_v136
		_340_v136 = _337_v133
		var _341_v137 *C0
		_ = _341_v137
		var _nw55 *C0 = New_C0_()
		_ = _nw55
		_nw55.Ctor__(_340_v136, _337_v133)
		_341_v137 = _nw55
		var _342_v138 _dafny.MultiSet
		_ = _342_v138
		_342_v138 = _dafny.MultiSetOf(_341_v137, _341_v137, _341_v137, _341_v137)
		var _343_v139 _dafny.MultiSet
		_ = _343_v139
		_343_v139 = _342_v138
		var _source5 _dafny.MultiSet = _342_v138
		_ = _source5
		var _344___mcc_h19 _dafny.MultiSet = _source5
		_ = _344___mcc_h19
		var _345_cf14 _dafny.MultiSet = _344___mcc_h19
		_ = _345_cf14
		var _346_v140 *C2
		_ = _346_v140
		var _nw56 *C2 = New_C2_()
		_ = _nw56
		_nw56.Ctor__()
		_346_v140 = _nw56
		var _347_v141 _dafny.Array
		_ = _347_v141
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
		_ = _nw57
		_347_v141 = _nw57
		var _348_v142 _dafny.Array
		_ = _348_v142
		var _nwElement0_11 _dafny.Int = _141_v2
		_ = _nwElement0_11
		var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(6))
		_ = _nw58
		(_nw58).ArraySet1(_nwElement0_11, 0)
		(_nw58).ArraySet1(Companion_Default___.Fm0(_145_globalState), 1)
		(_nw58).ArraySet1(_141_v2, 2)
		(_nw58).ArraySet1(_141_v2, 3)
		(_nw58).ArraySet1(_141_v2, 4)
		(_nw58).ArraySet1(_dafny.IntOfUint32((_143_v4).Cardinality()), 5)
		_348_v142 = _nw58
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_347_v141), 0))
		_ = _index39
		(_347_v141).ArraySet1(_348_v142, (_index39).Int())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_347_v141), 0))
		_ = _index40
		(_347_v141).ArraySet1(_348_v142, (_index40).Int())
		(_145_globalState).F2 = _337_v133
		_192_v43 = (_193_v44).F20()
		var _rhs62 _dafny.Array = _138_v1
		_ = _rhs62
		var _rhs63 bool = _137_v0
		_ = _rhs63
		var _lhs43 *GlobalState = _145_globalState
		_ = _lhs43
		_138_v1 = _rhs62
		_lhs43.F10 = _rhs63
		(_145_globalState).F2 = _337_v133
		var _349_v143 _dafny.Sequence
		_ = _349_v143
		_349_v143 = _dafny.SeqOf(_338_v134, _137_v0)
		var _350_v144 _dafny.Map
		_ = _350_v144
		_350_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_349_v143, (_181_v38).Dtor_cf2())
		var _351_v145 _dafny.Map
		_ = _351_v145
		_351_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_350_v144).Cardinality(), (func() _dafny.Int {
			if (_142_v3).Contains(_336_v132) {
				return (_142_v3).Get(_336_v132).(_dafny.Int)
			}
			return _337_v133
		})())
		_337_v133 = (func() _dafny.Int {
			if (_351_v145).Contains(_337_v133) {
				return (_351_v145).Get(_337_v133).(_dafny.Int)
			}
			return _337_v133
		})()
	} else {
		var _352_v146 _dafny.Array
		_ = _352_v146
		var _nwElement0_12 _dafny.Int = _141_v2
		_ = _nwElement0_12
		var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(22))
		_ = _nw59
		(_nw59).ArraySet1(_nwElement0_12, 0)
		(_nw59).ArraySet1(_141_v2, 1)
		(_nw59).ArraySet1(_141_v2, 2)
		(_nw59).ArraySet1(_141_v2, 3)
		(_nw59).ArraySet1(_141_v2, 4)
		(_nw59).ArraySet1(_141_v2, 5)
		(_nw59).ArraySet1(_141_v2, 6)
		(_nw59).ArraySet1(_141_v2, 7)
		(_nw59).ArraySet1(_141_v2, 8)
		(_nw59).ArraySet1(_141_v2, 9)
		(_nw59).ArraySet1(_141_v2, 10)
		(_nw59).ArraySet1(_141_v2, 11)
		(_nw59).ArraySet1(_141_v2, 12)
		(_nw59).ArraySet1(_141_v2, 13)
		(_nw59).ArraySet1(_141_v2, 14)
		(_nw59).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ify")).Cardinality()), 15)
		(_nw59).ArraySet1(_141_v2, 16)
		(_nw59).ArraySet1(_141_v2, 17)
		(_nw59).ArraySet1(_141_v2, 18)
		(_nw59).ArraySet1(_141_v2, 19)
		(_nw59).ArraySet1(_141_v2, 20)
		(_nw59).ArraySet1(_141_v2, 21)
		_352_v146 = _nw59
		var _353_v147 D4
		_ = _353_v147
		_353_v147 = Companion_D4_.Create_DC10_(_352_v146)
		var _354_v148 D4
		_ = _354_v148
		_354_v148 = Companion_D4_.Create_DC11_(_353_v147)
		var _pat_let_tv15 = _352_v146
		_ = _pat_let_tv15
		var _355_v149 _dafny.Map
		_ = _355_v149
		_355_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, func(_pat_let16_0 D4) D4 {
			return func(_356_dt__update__tmp_h8 D4) D4 {
				return func(_pat_let17_0 D4) D4 {
					return func(_357_dt__update_hcf17_h0 D4) D4 {
						return Companion_D4_.Create_DC11_(_357_dt__update_hcf17_h0)
					}(_pat_let17_0)
				}(Companion_D4_.Create_DC10_(_pat_let_tv15))
			}(_pat_let16_0)
		}(_354_v148))
		var _pat_let_tv16 = _353_v147
		_ = _pat_let_tv16
		var _358_v150 _dafny.Set
		_ = _358_v150
		_358_v150 = _dafny.SetOf(_355_v149, _355_v149, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, func(_pat_let18_0 D4) D4 {
			return func(_359_dt__update__tmp_h9 D4) D4 {
				return func(_pat_let19_0 D4) D4 {
					return func(_360_dt__update_hcf17_h1 D4) D4 {
						return Companion_D4_.Create_DC11_(_360_dt__update_hcf17_h1)
					}(_pat_let19_0)
				}(_pat_let_tv16)
			}(_pat_let18_0)
		}(_354_v148)), _355_v149)
		var _source6 D7 = Companion_D7_.Create_DC18_((_358_v150).Union(_dafny.SetOf(_355_v149, _355_v149, _355_v149, _355_v149)))
		_ = _source6
		if _source6.Is_DC19() {
			var _361___mcc_h20 _dafny.Set = _source6.Get_().(D7_DC19).Cf38
			_ = _361___mcc_h20
			var _362___mcc_h21 bool = _source6.Get_().(D7_DC19).Cf39
			_ = _362___mcc_h21
			var _363___mcc_h22 bool = _source6.Get_().(D7_DC19).Cf40
			_ = _363___mcc_h22
			var _364_cf40 bool = _363___mcc_h22
			_ = _364_cf40
			var _365_cf39 bool = _362___mcc_h21
			_ = _365_cf39
			var _366_cf38 _dafny.Set = _361___mcc_h20
			_ = _366_cf38
			var _367_v151 D2
			_ = _367_v151
			_367_v151 = Companion_D2_.Create_DC7_(Companion_D2_.Create_DC4_(_184_v41))
			var _368_v152 D2
			_ = _368_v152
			_368_v152 = Companion_D2_.Create_DC4_(_184_v41)
			_367_v151 = Companion_D2_.Create_DC7_(_368_v152)
			(_145_globalState).F9 = (func() _dafny.Int {
				if false {
					return (_141_v2).Plus(_141_v2)
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mject"), _143_v4)).Cardinality())
			})()
			var _369_v153 _dafny.Int
			_ = _369_v153
			_369_v153 = _141_v2
			var _370_v154 *C0
			_ = _370_v154
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__(_369_v153, _141_v2)
			_370_v154 = _nw60
			var _371_v155 _dafny.Array
			_ = _371_v155
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw61
			_371_v155 = _nw61
			var _372_v156 D9
			_ = _372_v156
			_372_v156 = Companion_D9_.Create_DC24_(_371_v155)
			var _373_v157 D9
			_ = _373_v157
			_373_v157 = Companion_D9_.Create_DC24_((_372_v156).Dtor_cf49())
			var _374_v158 _dafny.Sequence
			_ = _374_v158
			_374_v158 = _dafny.SeqOf(_365_cf39)
			var _rhs64 _dafny.Int = _dafny.IntOfInt64(373)
			_ = _rhs64
			var _rhs65 bool = Companion_Default___.Fm1(_365_cf39, _143_v4, _145_globalState)
			_ = _rhs65
			var _rhs66 _dafny.Int = Companion_Default___.Fm0(_145_globalState)
			_ = _rhs66
			var _rhs67 D9 = (func() D9 {
				if !((Companion_D2_.Create_DC5_(_137_v0, _dafny.IntOfUint32((_374_v158).Cardinality()), _138_v1, _178_v35)).Dtor_cf5()) {
					return (func() D9 {
						if _365_cf39 {
							return _372_v156
						}
						return _372_v156
					})()
				}
				return _373_v157
			})()
			_ = _rhs67
			var _lhs44 *GlobalState = _145_globalState
			_ = _lhs44
			var _lhs45 *GlobalState = _145_globalState
			_ = _lhs45
			_lhs44.F2 = _rhs64
			_365_cf39 = _rhs65
			_lhs45.F9 = _rhs66
			_372_v156 = _rhs67
		} else if _source6.Is_DC20() {
			var _375___mcc_h23 _dafny.Set = _source6.Get_().(D7_DC20).Cf41
			_ = _375___mcc_h23
			var _376___mcc_h24 D5 = _source6.Get_().(D7_DC20).Cf42
			_ = _376___mcc_h24
			var _377___mcc_h25 _dafny.Int = _source6.Get_().(D7_DC20).Cf43
			_ = _377___mcc_h25
			var _378_cf43 _dafny.Int = _377___mcc_h25
			_ = _378_cf43
			var _379_cf42 D5 = _376___mcc_h24
			_ = _379_cf42
			var _380_cf41 _dafny.Set = _375___mcc_h23
			_ = _380_cf41
			(_145_globalState).F0 = _137_v0
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_138_v1), 0))
			_ = _index41
			(_138_v1).ArraySet1(_137_v0, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_138_v1), 0))
			_ = _index42
			var _rhs68 bool = true
			_ = _rhs68
			var _rhs69 _dafny.Int = ((_378_cf43).Plus(_378_cf43)).Plus((_dafny.Zero).Minus(_141_v2))
			_ = _rhs69
			var _lhs46 _dafny.Array = _138_v1
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_138_v1), 0))
			_ = _lhs47
			(_lhs46).ArraySet1(_rhs68, (_lhs47).Int())
			_378_cf43 = _rhs69
			var _381_v159 _dafny.Int
			_ = _381_v159
			_381_v159 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(68))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_382_cf43 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_383_i8 _dafny.Int) _dafny.Int {
					return _382_cf43
				}
			})(_378_cf43)))).Cardinality())
			var _384_v160 _dafny.Map
			_ = _384_v160
			_384_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_cf43, true)
			var _385_v161 *C0
			_ = _385_v161
			var _nw62 *C0 = New_C0_()
			_ = _nw62
			_nw62.Ctor__(_381_v159, (_384_v160).Cardinality())
			_385_v161 = _nw62
			var _386_v162 D4
			_ = _386_v162
			_386_v162 = Companion_D4_.Create_DC9_(_385_v161)
			var _387_v163 _dafny.Map
			_ = _387_v163
			_387_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _385_v161)
			var _388_v164 _dafny.Array
			_ = _388_v164
			var _nwElement0_13 *C0 = _385_v161
			_ = _nwElement0_13
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(10))
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_13, 0)
			(_nw63).ArraySet1(_385_v161, 1)
			(_nw63).ArraySet1(_385_v161, 2)
			(_nw63).ArraySet1((_386_v162).Dtor_cf15(), 3)
			(_nw63).ArraySet1(_385_v161, 4)
			(_nw63).ArraySet1(_385_v161, 5)
			(_nw63).ArraySet1(_385_v161, 6)
			(_nw63).ArraySet1(_385_v161, 7)
			(_nw63).ArraySet1(_385_v161, 8)
			(_nw63).ArraySet1((func() *C0 {
				if (_387_v163).Contains((_138_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_138_v1), 0))).Int()).(bool)) {
					return (_387_v163).Get((_138_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_138_v1), 0))).Int()).(bool)).(*C0)
				}
				return _385_v161
			})(), 9)
			_388_v164 = _nw63
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_388_v164), 0))
			_ = _index43
			(_388_v164).ArraySet1(_385_v161, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_388_v164), 0))
			_ = _index44
			(_388_v164).ArraySet1(_385_v161, (_index44).Int())
			var _389_v165 _dafny.Int
			_ = _389_v165
			var _390_v166 _dafny.Int
			_ = _390_v166
			var _out14 _dafny.Int
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			_out14, _out15 = Companion_Default___.M0((_138_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_138_v1), 0))).Int()).(bool), _384_v160, _145_globalState)
			_389_v165 = _out14
			_390_v166 = _out15
		} else if _source6.Is_DC18() {
			var _391___mcc_h26 _dafny.Set = _source6.Get_().(D7_DC18).Cf37
			_ = _391___mcc_h26
			var _392_cf37 _dafny.Set = _391___mcc_h26
			_ = _392_cf37
			var _rhs70 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_143_v4, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_193_v44).F20(), (_193_v44).F20()), _143_v4))
			_ = _rhs70
			var _rhs71 _dafny.Int = (_141_v2).Times((_141_v2).Times(_141_v2))
			_ = _rhs71
			var _lhs48 *GlobalState = _145_globalState
			_ = _lhs48
			_143_v4 = _rhs70
			_lhs48.F9 = _rhs71
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_352_v146), 0))
			_ = _index45
			(_352_v146).ArraySet1(_141_v2, (_index45).Int())
			var _393_v167 _dafny.Array
			_ = _393_v167
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(Companion_D10_.Default(), _dafny.IntOfInt64(25))
			_ = _nw64
			_393_v167 = _nw64
			var _394_v168 D10
			_ = _394_v168
			_394_v168 = Companion_D10_.Create_DC29_()
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_393_v167), 0))
			_ = _index46
			(_393_v167).ArraySet1(_394_v168, (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_352_v146), 0))
			_ = _index47
			(_352_v146).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_145_globalState), (_335_v131).Cardinality()), (_index47).Int())
			var _395_v169 *C0
			_ = _395_v169
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__(_dafny.IntOfUint32((_180_v37).Cardinality()), _141_v2)
			_395_v169 = _nw65
			var _396_v170 D2
			_ = _396_v170
			_396_v170 = Companion_D2_.Create_DC6_(_141_v2, true, _137_v0, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_144_v5).Contains(_137_v0) {
					return (_144_v5).Get(_137_v0).(bool)
				}
				return true
			})(), _395_v169)).Cardinality()))
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_352_v146), 0))
			_ = _index48
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_393_v167), 0))
			_ = _index49
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_352_v146), 0))
			_ = _index50
			var _rhs72 _dafny.Int = (_141_v2).Plus(_141_v2)
			_ = _rhs72
			var _rhs73 bool = _137_v0
			_ = _rhs73
			var _rhs74 D10 = Companion_D10_.Create_DC29_()
			_ = _rhs74
			var _rhs75 _dafny.Int = Companion_Default___.Fm0(_145_globalState)
			_ = _rhs75
			var _rhs76 _dafny.Int = (_396_v170).Dtor_cf12()
			_ = _rhs76
			var _lhs49 _dafny.Array = _352_v146
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_352_v146), 0))
			_ = _lhs50
			var _lhs51 *GlobalState = _145_globalState
			_ = _lhs51
			var _lhs52 _dafny.Array = _393_v167
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_393_v167), 0))
			_ = _lhs53
			var _lhs54 *GlobalState = _145_globalState
			_ = _lhs54
			var _lhs55 _dafny.Array = _352_v146
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_352_v146), 0))
			_ = _lhs56
			(_lhs49).ArraySet1(_rhs72, (_lhs50).Int())
			_lhs51.F3 = _rhs73
			(_lhs52).ArraySet1(_rhs74, (_lhs53).Int())
			_lhs54.F9 = _rhs75
			(_lhs55).ArraySet1(_rhs76, (_lhs56).Int())
			var _397_v171 _dafny.Map
			_ = _397_v171
			_397_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, (_352_v146).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_352_v146), 0))).Int()).(_dafny.Int))
			_137_v0 = (func() bool {
				if !((_397_v171).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hiwhxnyyd")).Cardinality()))) {
					return _137_v0
				}
				return _137_v0
			})()
			var _398_v172 _dafny.Map
			_ = _398_v172
			_398_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_141_v2, _dafny.IntOfUint32((_143_v4).Cardinality())), _137_v0)
			_398_v172 = (_398_v172).Update(((_352_v146).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_352_v146), 0))).Int()).(_dafny.Int)).Minus((_352_v146).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_352_v146), 0))).Int()).(_dafny.Int)), (func() bool {
				if _137_v0 {
					return _137_v0
				}
				return _137_v0
			})())
		} else {
			var _399___mcc_h27 D7 = _source6.Get_().(D7_DC21).Cf44
			_ = _399___mcc_h27
			var _400_cf44 D7 = _399___mcc_h27
			_ = _400_cf44
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_138_v1), 0))
			_ = _index51
			(_138_v1).ArraySet1((_141_v2).Cmp(_141_v2) != 0, (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_138_v1), 0))
			_ = _index52
			(_138_v1).ArraySet1(_137_v0, (_index52).Int())
			var _401_v173 _dafny.Int
			_ = _401_v173
			_401_v173 = (_dafny.Zero).Minus(_141_v2)
			var _402_v174 _dafny.MultiSet
			_ = _402_v174
			_402_v174 = _dafny.MultiSetOf((_193_v44).F20(), _192_v43)
			var _403_v175 T0
			_ = _403_v175
			var _nw66 *C0 = New_C0_()
			_ = _nw66
			_nw66.Ctor__(_401_v173, (_402_v174).Cardinality())
			_403_v175 = _nw66
			var _404_v176 _dafny.Map
			_ = _404_v176
			_404_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, _403_v175)
			var _405_v177 _dafny.Map
			_ = _405_v177
			_405_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), (_404_v176).Cardinality())
			var _rhs77 _dafny.Array = _352_v146
			_ = _rhs77
			var _rhs78 _dafny.Map = (_405_v177).Merge(_405_v177)
			_ = _rhs78
			_352_v146 = _rhs77
			_405_v177 = _rhs78
			_192_v43 = (_193_v44).F20()
			var _406_v178 _dafny.Array
			_ = _406_v178
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_9
			var _nw67 _dafny.Array
			_ = _nw67
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw67 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) bool = func(_407_i9 _dafny.Int) bool {
					return true
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw67 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw67).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw67).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_406_v178 = _nw67
			var _408_v179 D2
			_ = _408_v179
			_408_v179 = Companion_D2_.Create_DC5_(_137_v0, (_403_v175).F18(), _406_v178, _178_v35)
			var _409_v180 _dafny.Set
			_ = _409_v180
			_409_v180 = _dafny.SetOf(_408_v179)
			var _410_v181 _dafny.Map
			_ = _410_v181
			_410_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_409_v180).Difference(_409_v180), (_dafny.Zero).Minus((_180_v37).Select((Companion_Default___.SafeIndex((_403_v175).F18(), _dafny.IntOfUint32((_180_v37).Cardinality()))).Uint32()).(_dafny.Int)))
			_410_v181 = (_410_v181).Update(_409_v180, ((Companion_Default___.Fm16((_403_v175).F18(), (_138_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_138_v1), 0))).Int()).(bool), _dafny.CodePoint('b'), _141_v2, _145_globalState)).Merge(_144_v5)).Cardinality())
		}
		var _411_v182 _dafny.Sequence
		_ = _411_v182
		_411_v182 = _dafny.SeqOf(_143_v4, _143_v4)
		var _412_v183 _dafny.Array
		_ = _412_v183
		var _nwElement0_14 _dafny.Sequence = _143_v4
		_ = _nwElement0_14
		var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(28))
		_ = _nw68
		(_nw68).ArraySet1(_nwElement0_14, 0)
		(_nw68).ArraySet1(Companion_Default___.Fm6(_137_v0, _145_globalState), 1)
		(_nw68).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-743))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_413_v44 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_414_i10 _dafny.Int) _dafny.CodePoint {
				return (_413_v44).F20()
			}
		})(_193_v44))), 2)
		(_nw68).ArraySet1(_143_v4, 3)
		(_nw68).ArraySet1(_143_v4, 4)
		(_nw68).ArraySet1(_143_v4, 5)
		(_nw68).ArraySet1(_143_v4, 6)
		(_nw68).ArraySet1(Companion_Default___.Fm6(_137_v0, _145_globalState), 7)
		(_nw68).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_143_v4, _dafny.UnicodeSeqOfUtf8Bytes("l")), 8)
		(_nw68).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_143_v4, _143_v4), 9)
		(_nw68).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("veh"), 10)
		(_nw68).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(638))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_415_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_416_i11 _dafny.Int) _dafny.CodePoint {
				return _415_v43
			}
		})(_192_v43))), 11)
		(_nw68).ArraySet1(_143_v4, 12)
		(_nw68).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_143_v4, _143_v4), 13)
		(_nw68).ArraySet1(_143_v4, 14)
		(_nw68).ArraySet1(_143_v4, 15)
		(_nw68).ArraySet1(_dafny.Companion_Sequence_.Update(_143_v4, (Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_143_v4).Cardinality()))).Uint32(), (_193_v44).F20()), 16)
		(_nw68).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(489))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_417_v44 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_418_i12 _dafny.Int) _dafny.CodePoint {
				return (_417_v44).F20()
			}
		})(_193_v44))), 17)
		(_nw68).ArraySet1(_143_v4, 18)
		(_nw68).ArraySet1((_411_v182).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-360), _dafny.IntOfUint32((_411_v182).Cardinality()))).Uint32()).(_dafny.Sequence), 19)
		(_nw68).ArraySet1(_143_v4, 20)
		(_nw68).ArraySet1(_143_v4, 21)
		(_nw68).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("psmshitr"), 22)
		(_nw68).ArraySet1(_dafny.Companion_Sequence_.Update(_143_v4, (Companion_Default___.SafeIndex(_141_v2, _dafny.IntOfUint32((_143_v4).Cardinality()))).Uint32(), _dafny.CodePoint('f')), 23)
		(_nw68).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(478))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}(func(_419_i13 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		})), 24)
		(_nw68).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_143_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-912))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_420_i14 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))), 25)
		(_nw68).ArraySet1(_143_v4, 26)
		(_nw68).ArraySet1(_143_v4, 27)
		_412_v183 = _nw68
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_412_v183), 0))
		_ = _index53
		(_412_v183).ArraySet1(_143_v4, (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_412_v183), 0))
		_ = _index54
		(_412_v183).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(762))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_421_v44 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_422_i15 _dafny.Int) _dafny.CodePoint {
				return (_421_v44).F20()
			}
		})(_193_v44))), (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_138_v1), 0))
		_ = _index55
		(_138_v1).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_180_v37, _180_v37), (_index55).Int())
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_138_v1), 0))
		_ = _index56
		(_138_v1).ArraySet1((((_303_v108).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_303_v108), 0))).Int()).(D9)).Dtor_cf50()).Cmp(_141_v2) == 0, (_index56).Int())
		var _423_v184 D11
		_ = _423_v184
		_423_v184 = Companion_D11_.Create_DC31_(_182_v39)
		(_145_globalState).F2 = Companion_Default___.SafeDivisionInt((_141_v2).Times(((_423_v184).Dtor_cf58()).Cardinality()), _141_v2)
		var _424_v185 *C1
		_ = _424_v185
		var _nw69 *C1 = New_C1_()
		_ = _nw69
		_nw69.Ctor__(_192_v43)
		_424_v185 = _nw69
	}
	var _425_v186 _dafny.Array
	_ = _425_v186
	var _len0_10 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_10
	var _nw70 _dafny.Array
	_ = _nw70
	if _len0_10.Cmp(_dafny.Zero) == 0 {
		_nw70 = _dafny.NewArray(_len0_10)
	} else {
		var _init10 func(_dafny.Int) _dafny.Set = (func(_426_v131 _dafny.Set) func(_dafny.Int) _dafny.Set {
			return func(_427_i16 _dafny.Int) _dafny.Set {
				return _426_v131
			}
		})(_335_v131)
		_ = _init10
		var _element0_10 = _init10(_dafny.Zero)
		_ = _element0_10
		_nw70 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
		(_nw70).ArraySet1(_element0_10, 0)
		var _nativeLen0_10 = (_len0_10).Int()
		_ = _nativeLen0_10
		for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
			(_nw70).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
		}
	}
	_425_v186 = _nw70
	var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_425_v186), 0))
	_ = _index57
	(_425_v186).ArraySet1((_335_v131).Difference(_335_v131), (_index57).Int())
	var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_425_v186), 0))
	_ = _index58
	(_425_v186).ArraySet1(_335_v131, (_index58).Int())
	(_145_globalState).F0 = ((_141_v2).Minus(_141_v2)).Cmp(_141_v2) >= 0
	_142_v3 = (_142_v3).Merge((_142_v3).Merge(_142_v3))
	var _428_v187 D6
	_ = _428_v187
	_428_v187 = Companion_D6_.Create_DC15_(_dafny.IntOfInt64(818), _141_v2, _180_v37, _141_v2, _dafny.SetOf(_141_v2))
	var _hi1 _dafny.Int = _141_v2
	_ = _hi1
	for _429_i17 := ((_428_v187).Dtor_cf27()).Cardinality(); _429_i17.Cmp(_hi1) < 0; _429_i17 = _429_i17.Plus(_dafny.One) {
		var _rhs79 bool = _137_v0
		_ = _rhs79
		var _rhs80 bool = _137_v0
		_ = _rhs80
		var _rhs81 _dafny.Int = _429_i17
		_ = _rhs81
		var _lhs57 *GlobalState = _145_globalState
		_ = _lhs57
		var _lhs58 *GlobalState = _145_globalState
		_ = _lhs58
		_lhs57.F0 = _rhs79
		_lhs58.F0 = _rhs80
		_141_v2 = _rhs81
		var _430_v188 _dafny.Array
		_ = _430_v188
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_11
		var _nw71 _dafny.Array
		_ = _nw71
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw71 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = (func(_431_i17 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_432_i18 _dafny.Int) _dafny.Int {
					return (_432_i18).Times(_431_i17)
				}
			})(_429_i17)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw71 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw71).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw71).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_430_v188 = _nw71
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_430_v188), 0))
		_ = _index59
		(_430_v188).ArraySet1(_141_v2, (_index59).Int())
		var _433_v189 D4
		_ = _433_v189
		_433_v189 = Companion_D4_.Create_DC10_(_430_v188)
		var _pat_let_tv17 = _430_v188
		_ = _pat_let_tv17
		var _434_v190 _dafny.Array
		_ = _434_v190
		var _nwElement0_15 D4 = _433_v189
		_ = _nwElement0_15
		var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(13))
		_ = _nw72
		(_nw72).ArraySet1(_nwElement0_15, 0)
		(_nw72).ArraySet1(_433_v189, 1)
		(_nw72).ArraySet1(_433_v189, 2)
		(_nw72).ArraySet1(_433_v189, 3)
		(_nw72).ArraySet1(_433_v189, 4)
		(_nw72).ArraySet1(_433_v189, 5)
		(_nw72).ArraySet1(_433_v189, 6)
		(_nw72).ArraySet1(_433_v189, 7)
		(_nw72).ArraySet1(func(_pat_let20_0 D4) D4 {
			return func(_435_dt__update__tmp_h10 D4) D4 {
				return func(_pat_let21_0 _dafny.Array) D4 {
					return func(_436_dt__update_hcf16_h0 _dafny.Array) D4 {
						return Companion_D4_.Create_DC10_(_436_dt__update_hcf16_h0)
					}(_pat_let21_0)
				}(_pat_let_tv17)
			}(_pat_let20_0)
		}(_433_v189), 8)
		(_nw72).ArraySet1(_433_v189, 9)
		(_nw72).ArraySet1(_433_v189, 10)
		(_nw72).ArraySet1(Companion_D4_.Create_DC10_(_430_v188), 11)
		(_nw72).ArraySet1(_433_v189, 12)
		_434_v190 = _nw72
		var _437_v191 _dafny.Map
		_ = _437_v191
		_437_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_429_i17, (_dafny.Zero).Minus((_428_v187).Dtor_cf26()))
		var _438_v192 _dafny.Set
		_ = _438_v192
		_438_v192 = _dafny.SetOf(_434_v190, _434_v190)
		var _pat_let_tv18 = _438_v192
		_ = _pat_let_tv18
		var _pat_let_tv19 = _141_v2
		_ = _pat_let_tv19
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_430_v188), 0))
		_ = _index60
		(_430_v188).ArraySet1((func(_pat_let22_0 D7) D7 {
			return func(_439_dt__update__tmp_h11 D7) D7 {
				return func(_pat_let23_0 _dafny.Set) D7 {
					return func(_440_dt__update_hcf41_h0 _dafny.Set) D7 {
						return func(_pat_let24_0 _dafny.Int) D7 {
							return func(_441_dt__update_hcf43_h0 _dafny.Int) D7 {
								return Companion_D7_.Create_DC20_(_440_dt__update_hcf41_h0, (_439_dt__update__tmp_h11).Dtor_cf42(), _441_dt__update_hcf43_h0)
							}(_pat_let24_0)
						}(_pat_let_tv19)
					}(_pat_let23_0)
				}(_pat_let_tv18)
			}(_pat_let22_0)
		}(Companion_D7_.Create_DC20_(_dafny.SetOf(_434_v190, _434_v190), Companion_D5_.Create_DC13_(Companion_Default___.Fm0(_145_globalState), Companion_Default___.Fm0(_145_globalState), _137_v0), (_437_v191).Cardinality()))).Dtor_cf43(), (_index60).Int())
		var _442_v193 _dafny.Int
		_ = _442_v193
		_442_v193 = _429_i17
		var _443_v194 T0
		_ = _443_v194
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__(_442_v193, (_141_v2).Minus((_dafny.Zero).Minus(_141_v2)))
		_443_v194 = _nw73
		var _444_v195 _dafny.Map
		_ = _444_v195
		_444_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_137_v0, _dafny.UnicodeSeqOfUtf8Bytes("pdygc"), _145_globalState), _443_v194)
		_443_v194 = (func() T0 {
			if (_444_v195).Contains(_137_v0) {
				return (_444_v195).Get(_137_v0).(T0)
			}
			return _443_v194
		})()
		var _445_v196 D11
		_ = _445_v196
		_445_v196 = Companion_D11_.Create_DC31_(_182_v39)
		var _446_v197 D11
		_ = _446_v197
		_446_v197 = Companion_D11_.Create_DC33_(_445_v196)
		var _447_v198 D11
		_ = _447_v198
		_447_v198 = Companion_D11_.Create_DC33_(_446_v197)
		_447_v198 = _447_v198
	}
	for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_138_v1), 0))); ; {
		_guard_loop_0, _ok15 := _iter15()
		if !_ok15 {
			break
		}
		var _448_i19 _dafny.Int
		_448_i19 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_448_i19).Sign() != -1) && ((_448_i19).Cmp(_dafny.ArrayLen((_138_v1), 0)) < 0)) {
			(_138_v1).ArraySet1((_141_v2).Cmp(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, false)).Cardinality()).Plus(_dafny.IntOfUint32((_143_v4).Cardinality()))) != 0, (_448_i19).Int())
		}
	}
	var _449_v199 _dafny.Set
	_ = _449_v199
	_449_v199 = _dafny.SetOf(_143_v4, _143_v4, _dafny.UnicodeSeqOfUtf8Bytes("r"))
	_449_v199 = _449_v199
	var _hi2 _dafny.Int = _141_v2
	_ = _hi2
	for _450_i20 := Companion_Default___.SafeModuloInt(_141_v2, _141_v2); _450_i20.Cmp(_hi2) < 0; _450_i20 = _450_i20.Plus(_dafny.One) {
		(_145_globalState).F9 = _141_v2
		if !(_137_v0) {
			var _451_v200 *C3
			_ = _451_v200
			var _nw74 *C3 = New_C3_()
			_ = _nw74
			_nw74.Ctor__(_143_v4, _dafny.IntOfInt64(633), _141_v2)
			_451_v200 = _nw74
			var _452_v201 _dafny.Int
			_ = _452_v201
			_452_v201 = _450_i20
			var _453_v202 *C0
			_ = _453_v202
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__(_452_v201, (_451_v200).Fm3(_137_v0, _450_i20, _145_globalState))
			_453_v202 = _nw75
			var _454_v203 _dafny.MultiSet
			_ = _454_v203
			_454_v203 = _dafny.MultiSetOf(_453_v202, _453_v202, _453_v202, _453_v202)
			var _455_v204 _dafny.Array
			_ = _455_v204
			var _nwElement0_16 _dafny.Int = _141_v2
			_ = _nwElement0_16
			var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(11))
			_ = _nw76
			(_nw76).ArraySet1(_nwElement0_16, 0)
			(_nw76).ArraySet1(_450_i20, 1)
			(_nw76).ArraySet1(_450_i20, 2)
			(_nw76).ArraySet1((_454_v203).Cardinality(), 3)
			(_nw76).ArraySet1((_451_v200).Fm3(_137_v0, _141_v2, _145_globalState), 4)
			(_nw76).ArraySet1(_141_v2, 5)
			(_nw76).ArraySet1(_450_i20, 6)
			(_nw76).ArraySet1(_dafny.IntOfInt64(834), 7)
			(_nw76).ArraySet1(_141_v2, 8)
			(_nw76).ArraySet1(_141_v2, 9)
			(_nw76).ArraySet1(_dafny.IntOfInt64(-510), 10)
			_455_v204 = _nw76
			var _456_v205 D4
			_ = _456_v205
			_456_v205 = Companion_D4_.Create_DC10_(_455_v204)
			_456_v205 = Companion_D4_.Create_DC10_(_455_v204)
			var _457_v206 _dafny.MultiSet
			_ = _457_v206
			_457_v206 = _dafny.MultiSetOf(_138_v1, _138_v1, _138_v1, _138_v1, _138_v1)
			_457_v206 = _457_v206
			var _458_v207 _dafny.Array
			_ = _458_v207
			var _459_v208 bool
			_ = _459_v208
			var _out16 _dafny.Array
			_ = _out16
			var _out17 bool
			_ = _out17
			_out16, _out17 = (_451_v200).M1(_145_globalState)
			_458_v207 = _out16
			_459_v208 = _out17
			var _460_v209 _dafny.Map
			_ = _460_v209
			_460_v209 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v1, _183_v40)
			var _rhs82 _dafny.Array = _458_v207
			_ = _rhs82
			var _rhs83 _dafny.Int = (((_460_v209).Merge(_460_v209)).Update(_138_v1, (_183_v40).Union(_183_v40))).Cardinality()
			_ = _rhs83
			var _rhs84 *C3 = _451_v200
			_ = _rhs84
			_458_v207 = _rhs82
			_141_v2 = _rhs83
			_451_v200 = _rhs84
		} else {
			var _461_v210 *C2
			_ = _461_v210
			var _nw77 *C2 = New_C2_()
			_ = _nw77
			_nw77.Ctor__()
			_461_v210 = _nw77
			var _462_v211 _dafny.MultiSet
			_ = _462_v211
			_462_v211 = _dafny.MultiSetOf(_461_v210)
			var _463_v212 _dafny.Set
			_ = _463_v212
			_463_v212 = _dafny.SetOf(_462_v211, _462_v211)
			(_145_globalState).F0 = ((_463_v212).Difference(_463_v212)).IsProperSubsetOf(_463_v212)
			var _464_v213 _dafny.Map
			_ = _464_v213
			_464_v213 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_461_v210, false)
			var _465_v215 _dafny.Map
			_ = _465_v215
			_465_v215 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_141_v2), !(_137_v0))
			var _466_v216 _dafny.Set
			_ = _466_v216
			_466_v216 = _dafny.SetOf(_465_v215, _465_v215, _465_v215, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v2, _137_v0), _465_v215)
			var _467_v217 _dafny.Map
			_ = _467_v217
			_467_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_464_v213).Contains(_461_v210) {
					return (_464_v213).Get(_461_v210).(bool)
				}
				return _137_v0
			})(), Companion_D10_.Create_DC28_(func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter16 := _dafny.Iterate((_466_v216).Elements()); ; {
					_compr_15, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _468_v214 _dafny.Map
					_468_v214 = interface{}(_compr_15).(_dafny.Map)
					if (_466_v216).Contains(_468_v214) {
						_coll15.Add(_468_v214, _141_v2)
					}
				}
				return _coll15.ToMap()
			}()))
			var _469_v218 _dafny.Map
			_ = _469_v218
			_469_v218 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _467_v217)
			_469_v218 = _469_v218
			var _470_v219 _dafny.Int
			_ = _470_v219
			_470_v219 = _450_i20
			var _471_v220 T0
			_ = _471_v220
			var _nw78 *C0 = New_C0_()
			_ = _nw78
			_nw78.Ctor__(_470_v219, _450_i20)
			_471_v220 = _nw78
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__(_470_v219, (_471_v220).F18())
			_471_v220 = _nw79
			(_145_globalState).F2 = (func() _dafny.Int {
				if !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm6(_137_v0, _145_globalState), Companion_Default___.Fm6((_181_v38).Dtor_cf2(), _145_globalState)) {
					return Companion_Default___.SafeDivisionInt(_141_v2, _141_v2)
				}
				return (_471_v220).F18()
			})()
			_465_v215 = (_465_v215).Update((func() _dafny.Int {
				if _137_v0 {
					return _dafny.IntOfInt64(925)
				}
				return _dafny.IntOfUint32((_dafny.SeqOf(_137_v0, _137_v0, _137_v0)).Cardinality())
			})(), _137_v0)
		}
		(_145_globalState).F3 = _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm11(_450_i20, _137_v0, _450_i20, _145_globalState), _dafny.SeqOf(_450_i20, _450_i20))
		var _472_v221 _dafny.Array
		_ = _472_v221
		var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
		_ = _nw80
		_472_v221 = _nw80
		var _473_v222 _dafny.Set
		_ = _473_v222
		_473_v222 = _dafny.SetOf(_472_v221, _472_v221)
		var _474_v223 D7
		_ = _474_v223
		_474_v223 = Companion_D7_.Create_DC19_(_473_v222, _137_v0, _137_v0)
		var _475_v224 D7
		_ = _475_v224
		_475_v224 = Companion_D7_.Create_DC21_(_474_v223)
		_475_v224 = _475_v224
	}
	_dafny.Print(_137_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_141_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_142_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(213))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_143_v4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_144_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_145_globalState).F1(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F8).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(213), _dafny.IntOfInt64(213))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState.F11).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F12()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-213))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F13().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F15()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v35).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(212), _dafny.IntOfInt64(212), _dafny.IntOfInt64(212))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v36).Dtor_cf28())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v36).Dtor_cf29())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v36).Dtor_cf30())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v36).Dtor_cf31())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v36).Dtor_cf32())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_180_v37, _dafny.SeqOf(_dafny.IntOfInt64(-496), _dafny.IntOfInt64(212), _dafny.IntOfInt64(212), _dafny.IntOfInt64(212), _dafny.IntOfInt64(-496), _dafny.IntOfInt64(212), _dafny.IntOfInt64(212))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_v38).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v39).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_183_v40).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v41).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_191_v42).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_192_v43)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_193_v44).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_194_v45).Dtor_cf22()).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_302_v107).Dtor_cf50())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v108).ArrayGet1((_dafny.Zero).Int()).(D9)).Dtor_cf50())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v108).ArrayGet1((_dafny.One).Int()).(D9)).Dtor_cf50())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v108).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D9)).Dtor_cf50())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v108).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D9)).Dtor_cf50())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v108).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D9)).Dtor_cf50())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_303_v108).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D9)).Dtor_cf50())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_334_v130).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v131).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v186).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_428_v187).Dtor_cf23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_428_v187).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_428_v187).Dtor_cf25(), _dafny.SeqOf(_dafny.IntOfInt64(-496), _dafny.IntOfInt64(212), _dafny.IntOfInt64(212), _dafny.IntOfInt64(212), _dafny.IntOfInt64(-496), _dafny.IntOfInt64(212), _dafny.IntOfInt64(212))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_428_v187).Dtor_cf26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_428_v187).Dtor_cf27()).Equals(_dafny.SetOf(_dafny.IntOfInt64(212))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_449_v199).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("y"), _dafny.UnicodeSeqOfUtf8Bytes("r"))))
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

func (CompanionStruct_D0_) Default() _dafny.Int {
	return _dafny.Zero
}

func (_this D0) Dtor_cf0() _dafny.Int {
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
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
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
	Cf2 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 bool) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
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

type D1_DC3 struct {
	Cf3 D1
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 D1) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(false)
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf3() D1 {
	return _this.Get_().(D1_DC3).Cf3
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
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2 == data2.Cf2
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1 == data2.Cf1
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

type D2_DC5 struct {
	Cf5 bool
	Cf6 _dafny.Int
	Cf7 _dafny.Array
	Cf8 _dafny.MultiSet
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 bool, Cf6 _dafny.Int, Cf7 _dafny.Array, Cf8 _dafny.MultiSet) D2 {
	return D2{D2_DC5{Cf5, Cf6, Cf7, Cf8}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC6 struct {
	Cf9  _dafny.Int
	Cf10 bool
	Cf11 bool
	Cf12 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf9 _dafny.Int, Cf10 bool, Cf11 bool, Cf12 _dafny.Int) D2 {
	return D2{D2_DC6{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC4 struct {
	Cf4 _dafny.Array
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.Array) D2 {
	return D2{D2_DC4{Cf4}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC7 struct {
	Cf13 D2
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf13 D2) D2 {
	return D2{D2_DC7{Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMultiSet)
}

func (_this D2) Dtor_cf5() bool {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf8() _dafny.MultiSet {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) Dtor_cf13() D2 {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7 && data1.Cf8.Equals(data2.Cf8)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf4 == data2.Cf4
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf14 _dafny.MultiSet
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 _dafny.MultiSet) D3 {
	return D3{D3_DC8{Cf14}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D3) Dtor_cf14() _dafny.MultiSet {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ")"
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
	Cf16 _dafny.Array
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf16 _dafny.Array) D4 {
	return D4{D4_DC10{Cf16}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC9 struct {
	Cf15 *C0
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf15 *C0) D4 {
	return D4{D4_DC9{Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC11 struct {
	Cf17 D4
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 D4) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D4) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf15() *C0 {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) Dtor_cf17() D4 {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
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
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf16 == data2.Cf16
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17)
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
	Cf19 _dafny.Int
	Cf20 _dafny.Int
	Cf21 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf19 _dafny.Int, Cf20 _dafny.Int, Cf21 bool) D5 {
	return D5{D5_DC13{Cf19, Cf20, Cf21}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

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

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D5) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
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
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21
		}
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

type D6_DC15 struct {
	Cf23 _dafny.Int
	Cf24 _dafny.Int
	Cf25 _dafny.Sequence
	Cf26 _dafny.Int
	Cf27 _dafny.Set
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf23 _dafny.Int, Cf24 _dafny.Int, Cf25 _dafny.Sequence, Cf26 _dafny.Int, Cf27 _dafny.Set) D6 {
	return D6{D6_DC15{Cf23, Cf24, Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC16 struct {
	Cf28 _dafny.Int
	Cf29 bool
	Cf30 _dafny.Int
	Cf31 _dafny.Int
	Cf32 bool
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf28 _dafny.Int, Cf29 bool, Cf30 _dafny.Int, Cf31 _dafny.Int, Cf32 bool) D6 {
	return D6{D6_DC16{Cf28, Cf29, Cf30, Cf31, Cf32}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC17 struct {
	Cf33 _dafny.CodePoint
	Cf34 bool
	Cf35 bool
	Cf36 _dafny.CodePoint
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf33 _dafny.CodePoint, Cf34 bool, Cf35 bool, Cf36 _dafny.CodePoint) D6 {
	return D6{D6_DC17{Cf33, Cf34, Cf35, Cf36}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC14 struct {
	Cf22 *C1
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf22 *C1) D6 {
	return D6{D6_DC14{Cf22}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC15_(_dafny.Zero, _dafny.Zero, _dafny.EmptySeq, _dafny.Zero, _dafny.EmptySet)
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf23
}

func (_this D6) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D6_DC15).Cf25
}

func (_this D6) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Set {
	return _this.Get_().(D6_DC15).Cf27
}

func (_this D6) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf28
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC16).Cf29
}

func (_this D6) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf30
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC16).Cf32
}

func (_this D6) Dtor_cf33() _dafny.CodePoint {
	return _this.Get_().(D6_DC17).Cf33
}

func (_this D6) Dtor_cf34() bool {
	return _this.Get_().(D6_DC17).Cf34
}

func (_this D6) Dtor_cf35() bool {
	return _this.Get_().(D6_DC17).Cf35
}

func (_this D6) Dtor_cf36() _dafny.CodePoint {
	return _this.Get_().(D6_DC17).Cf36
}

func (_this D6) Dtor_cf22() *C1 {
	return _this.Get_().(D6_DC14).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Equals(data2.Cf25) && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27.Equals(data2.Cf27)
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29 == data2.Cf29 && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf33 == data2.Cf33 && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf22 == data2.Cf22
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

type D7_DC19 struct {
	Cf38 _dafny.Set
	Cf39 bool
	Cf40 bool
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf38 _dafny.Set, Cf39 bool, Cf40 bool) D7 {
	return D7{D7_DC19{Cf38, Cf39, Cf40}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
	Cf41 _dafny.Set
	Cf42 D5
	Cf43 _dafny.Int
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf41 _dafny.Set, Cf42 D5, Cf43 _dafny.Int) D7 {
	return D7{D7_DC20{Cf41, Cf42, Cf43}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC18 struct {
	Cf37 _dafny.Set
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf37 _dafny.Set) D7 {
	return D7{D7_DC18{Cf37}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC21 struct {
	Cf44 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf44 D7) D7 {
	return D7{D7_DC21{Cf44}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(_dafny.EmptySet, false, false)
}

func (_this D7) Dtor_cf38() _dafny.Set {
	return _this.Get_().(D7_DC19).Cf38
}

func (_this D7) Dtor_cf39() bool {
	return _this.Get_().(D7_DC19).Cf39
}

func (_this D7) Dtor_cf40() bool {
	return _this.Get_().(D7_DC19).Cf40
}

func (_this D7) Dtor_cf41() _dafny.Set {
	return _this.Get_().(D7_DC20).Cf41
}

func (_this D7) Dtor_cf42() D5 {
	return _this.Get_().(D7_DC20).Cf42
}

func (_this D7) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf43
}

func (_this D7) Dtor_cf37() _dafny.Set {
	return _this.Get_().(D7_DC18).Cf37
}

func (_this D7) Dtor_cf44() D7 {
	return _this.Get_().(D7_DC21).Cf44
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf38.Equals(data2.Cf38) && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf41.Equals(data2.Cf41) && data1.Cf42.Equals(data2.Cf42) && data1.Cf43.Cmp(data2.Cf43) == 0
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf37.Equals(data2.Cf37)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf44.Equals(data2.Cf44)
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
	Cf46 _dafny.Int
	Cf47 _dafny.Array
	Cf48 _dafny.Int
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf46 _dafny.Int, Cf47 _dafny.Array, Cf48 _dafny.Int) D8 {
	return D8{D8_DC23{Cf46, Cf47, Cf48}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC22 struct {
	Cf45 *C2
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf45 *C2) D8 {
	return D8{D8_DC22{Cf45}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D8) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf46
}

func (_this D8) Dtor_cf47() _dafny.Array {
	return _this.Get_().(D8_DC23).Cf47
}

func (_this D8) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf48
}

func (_this D8) Dtor_cf45() *C2 {
	return _this.Get_().(D8_DC22).Cf45
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf45) + ")"
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
			return ok && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47 == data2.Cf47 && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf45 == data2.Cf45
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

type D9_DC25 struct {
	Cf50 _dafny.Int
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf50 _dafny.Int) D9 {
	return D9{D9_DC25{Cf50}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC26 struct {
	Cf51 bool
	Cf52 _dafny.Set
	Cf53 _dafny.Array
	Cf54 _dafny.Array
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf51 bool, Cf52 _dafny.Set, Cf53 _dafny.Array, Cf54 _dafny.Array) D9 {
	return D9{D9_DC26{Cf51, Cf52, Cf53, Cf54}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC27 struct {
	Cf55 _dafny.Int
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf55 _dafny.Int) D9 {
	return D9{D9_DC27{Cf55}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC24 struct {
	Cf49 _dafny.Array
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf49 _dafny.Array) D9 {
	return D9{D9_DC24{Cf49}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC25_(_dafny.Zero)
}

func (_this D9) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf50
}

func (_this D9) Dtor_cf51() bool {
	return _this.Get_().(D9_DC26).Cf51
}

func (_this D9) Dtor_cf52() _dafny.Set {
	return _this.Get_().(D9_DC26).Cf52
}

func (_this D9) Dtor_cf53() _dafny.Array {
	return _this.Get_().(D9_DC26).Cf53
}

func (_this D9) Dtor_cf54() _dafny.Array {
	return _this.Get_().(D9_DC26).Cf54
}

func (_this D9) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D9_DC27).Cf55
}

func (_this D9) Dtor_cf49() _dafny.Array {
	return _this.Get_().(D9_DC24).Cf49
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf50.Cmp(data2.Cf50) == 0
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52.Equals(data2.Cf52) && data1.Cf53 == data2.Cf53 && data1.Cf54 == data2.Cf54
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf55.Cmp(data2.Cf55) == 0
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf49 == data2.Cf49
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
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_() D10 {
	return D10{D10_DC29{}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC28 struct {
	Cf56 _dafny.Map
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf56 _dafny.Map) D10 {
	return D10{D10_DC28{Cf56}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC30 struct {
	Cf57 D10
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf57 D10) D10 {
	return D10{D10_DC30{Cf57}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_()
}

func (_this D10) Dtor_cf56() _dafny.Map {
	return _this.Get_().(D10_DC28).Cf56
}

func (_this D10) Dtor_cf57() D10 {
	return _this.Get_().(D10_DC30).Cf57
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf57) + ")"
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
			_, ok := other.Get_().(D10_DC29)
			return ok
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf56.Equals(data2.Cf56)
		}
	case D10_DC30:
		{
			data2, ok := other.Get_().(D10_DC30)
			return ok && data1.Cf57.Equals(data2.Cf57)
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

type D11_DC32 struct {
	Cf59 _dafny.Int
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf59 _dafny.Int) D11 {
	return D11{D11_DC32{Cf59}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC31 struct {
	Cf58 _dafny.MultiSet
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf58 _dafny.MultiSet) D11 {
	return D11{D11_DC31{Cf58}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC33 struct {
	Cf60 D11
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf60 D11) D11 {
	return D11{D11_DC33{Cf60}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC32_(_dafny.Zero)
}

func (_this D11) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D11_DC32).Cf59
}

func (_this D11) Dtor_cf58() _dafny.MultiSet {
	return _this.Get_().(D11_DC31).Cf58
}

func (_this D11) Dtor_cf60() D11 {
	return _this.Get_().(D11_DC33).Cf60
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf59) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf58) + ")"
		}
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf59.Cmp(data2.Cf59) == 0
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf58.Equals(data2.Cf58)
		}
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf60.Equals(data2.Cf60)
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

// Definition of datatype D12
type D12 struct {
	Data_D12_
}

func (_this D12) Get_() Data_D12_ {
	return _this.Data_D12_
}

type Data_D12_ interface {
	isD12()
}

type CompanionStruct_D12_ struct {
}

var Companion_D12_ = CompanionStruct_D12_{}

type D12_DC34 struct {
	Cf61 _dafny.Map
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf61 _dafny.Map) D12 {
	return D12{D12_DC34{Cf61}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D12) Dtor_cf61() _dafny.Map {
	return _this.Get_().(D12_DC34).Cf61
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf61.Equals(data2.Cf61)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D12) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D12)
	return ok && _this.Equals(typed)
}

func Type_D12_() _dafny.TypeDescriptor {
	return type_D12_{}
}

type type_D12_ struct {
}

func (_this type_D12_) Default() interface{} {
	return Companion_D12_.Default()
}

func (_this type_D12_) String() string {
	return "main.D12"
}
func (_this D12) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D12{}

// End of datatype D12

// Definition of trait T0
type T0 interface {
	String() string
	Fm2(globalState *GlobalState) _dafny.Int
	F17() _dafny.Int
	F18() _dafny.Int
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
	F0   bool
	F2   _dafny.Int
	F3   bool
	F8   _dafny.MultiSet
	F9   _dafny.Int
	F10  bool
	F11  _dafny.Array
	_f1  _dafny.Sequence
	_f4  bool
	_f5  _dafny.CodePoint
	_f6  _dafny.Int
	_f7  _dafny.Array
	_f12 _dafny.Map
	_f13 _dafny.Sequence
	_f14 bool
	_f15 _dafny.Map
	_f16 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F2 = _dafny.Zero
	_this.F3 = false
	_this.F8 = _dafny.EmptyMultiSet
	_this.F9 = _dafny.Zero
	_this.F10 = false
	_this.F11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.EmptySeq
	_this._f4 = false
	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f12 = _dafny.EmptyMap
	_this._f13 = _dafny.EmptySeq
	_this._f14 = false
	_this._f15 = _dafny.EmptyMap
	_this._f16 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Sequence, f2 _dafny.Int, f3 bool, f4 bool, f5 _dafny.CodePoint, f6 _dafny.Int, f7 _dafny.Array, f8 _dafny.MultiSet, f9 _dafny.Int, f10 bool, f11 _dafny.Array, f12 _dafny.Map, f13 _dafny.Sequence, f14 bool, f15 _dafny.Map, f16 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this).F9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.CodePoint {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Array {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F12() _dafny.Map {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Sequence {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Map {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f17 _dafny.Int
	_f18 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.Zero
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

func (_this *C0) F17() _dafny.Int {
	return _this._f17
}
func (_this *C0) F18() _dafny.Int {
	return _this._f18
}
func (_this *C0) Ctor__(f17 _dafny.Int, f18 _dafny.Int) {
	{
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C0) Fm2(globalState *GlobalState) _dafny.Int {
	{
		var _source7 _dafny.Int = (_this).F17()
		_ = _source7
		var _476___mcc_h0 _dafny.Int = _source7
		_ = _476___mcc_h0
		var _477_cf0 _dafny.Int = _476___mcc_h0
		_ = _477_cf0
		return (_this).F17()
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f20 _dafny.CodePoint
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f20 = _dafny.CodePoint('D')
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

func (_this *C1) Ctor__(f20 _dafny.CodePoint) {
	{
		(_this)._f20 = f20
	}
}
func (_this *C1) M3(p0 bool, globalState *GlobalState) (_dafny.Int, _dafny.MultiSet) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var _478_i0 _dafny.Int
		_ = _478_i0
		_478_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_478_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_478_i0 = (_478_i0).Plus(_dafny.One)
					var _479_v0 _dafny.Array
					_ = _479_v0
					var _nw81 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
					_ = _nw81
					_479_v0 = _nw81
					var _480_v1 _dafny.Int
					_ = _480_v1
					_480_v1 = _dafny.IntOfInt64(990)
					var _481_v2 _dafny.Map
					_ = _481_v2
					_481_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
						if p0 {
							return _479_v0
						}
						return _479_v0
					})(), _480_v1)
					var _482_v3 _dafny.Sequence
					_ = _482_v3
					_482_v3 = _dafny.SeqOf((Companion_Default___.Fm5((_dafny.Zero).Minus(_480_v1), _dafny.IntOfInt64(121), globalState)).Dtor_cf2())
					(globalState).F2 = (func() _dafny.Int {
						if (_481_v2).Contains(_479_v0) {
							return (_481_v2).Get(_479_v0).(_dafny.Int)
						}
						return (_dafny.Zero).Minus(_dafny.IntOfUint32((_482_v3).Cardinality()))
					})()
					if p0 {
						var _483_v4 _dafny.Map
						_ = _483_v4
						_483_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_480_v1, (_480_v1).Minus(_480_v1))
						_483_v4 = (_483_v4).Update(_480_v1, (_dafny.IntOfInt64(304)).Minus((_dafny.Zero).Minus(_480_v1)))
						var _484_v5 bool
						_ = _484_v5
						var _485_v6 _dafny.Int
						_ = _485_v6
						var _486_v7 bool
						_ = _486_v7
						var _487_v8 bool
						_ = _487_v8
						var _out18 bool
						_ = _out18
						var _out19 _dafny.Int
						_ = _out19
						var _out20 bool
						_ = _out20
						var _out21 bool
						_ = _out21
						_out18, _out19, _out20, _out21 = (_this).M4(globalState)
						_484_v5 = _out18
						_485_v6 = _out19
						_486_v7 = _out20
						_487_v8 = _out21
						(globalState).F10 = !(func() _dafny.Set {
							var _coll16 = _dafny.NewBuilder()
							_ = _coll16
							for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(959), _dafny.IntOfInt64(519))); ; {
								_compr_16, _ok17 := _iter17()
								if !_ok17 {
									break
								}
								var _488_v9 _dafny.Int
								_488_v9 = interface{}(_compr_16).(_dafny.Int)
								if ((_dafny.IntOfInt64(959)).Cmp(_488_v9) <= 0) && ((_488_v9).Cmp(_dafny.IntOfInt64(519)) < 0) {
									_coll16.Add((_488_v9).Plus(_485_v6))
								}
							}
							return _coll16.ToSet()
						}()).Contains(_480_v1)
						_480_v1 = _485_v6
						var _489_v10 D1
						_ = _489_v10
						_489_v10 = Companion_D1_.Create_DC1_(_479_v0)
						var _490_v11 D1
						_ = _490_v11
						_490_v11 = Companion_D1_.Create_DC3_(_489_v10)
						var _491_v12 D1
						_ = _491_v12
						_491_v12 = Companion_D1_.Create_DC1_(_479_v0)
						var _492_v13 _dafny.Array
						_ = _492_v13
						var _nwElement0_17 D1 = _491_v12
						_ = _nwElement0_17
						var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(15))
						_ = _nw82
						(_nw82).ArraySet1(_nwElement0_17, 0)
						(_nw82).ArraySet1(_491_v12, 1)
						(_nw82).ArraySet1(_491_v12, 2)
						(_nw82).ArraySet1(_491_v12, 3)
						(_nw82).ArraySet1(_491_v12, 4)
						(_nw82).ArraySet1(_491_v12, 5)
						(_nw82).ArraySet1(_491_v12, 6)
						(_nw82).ArraySet1(Companion_D1_.Create_DC1_(_479_v0), 7)
						(_nw82).ArraySet1(_491_v12, 8)
						(_nw82).ArraySet1(_491_v12, 9)
						(_nw82).ArraySet1(_491_v12, 10)
						(_nw82).ArraySet1(_491_v12, 11)
						(_nw82).ArraySet1(_491_v12, 12)
						(_nw82).ArraySet1(_491_v12, 13)
						(_nw82).ArraySet1(_491_v12, 14)
						_492_v13 = _nw82
						var _493_v14 _dafny.Sequence
						_ = _493_v14
						_493_v14 = _dafny.SeqOf(_492_v13, _492_v13)
						var _494_v15 _dafny.Array
						_ = _494_v15
						var _nwElement0_18 _dafny.Array = _492_v13
						_ = _nwElement0_18
						var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(19))
						_ = _nw83
						(_nw83).ArraySet1(_nwElement0_18, 0)
						(_nw83).ArraySet1(_492_v13, 1)
						(_nw83).ArraySet1(_492_v13, 2)
						(_nw83).ArraySet1(_492_v13, 3)
						(_nw83).ArraySet1((func() _dafny.Array {
							if p0 {
								return _492_v13
							}
							return _492_v13
						})(), 4)
						(_nw83).ArraySet1(_492_v13, 5)
						(_nw83).ArraySet1(_492_v13, 6)
						(_nw83).ArraySet1(_492_v13, 7)
						(_nw83).ArraySet1(_492_v13, 8)
						(_nw83).ArraySet1((func() _dafny.Array {
							if _486_v7 {
								return _492_v13
							}
							return _492_v13
						})(), 9)
						(_nw83).ArraySet1(_492_v13, 10)
						(_nw83).ArraySet1(_492_v13, 11)
						(_nw83).ArraySet1(_492_v13, 12)
						(_nw83).ArraySet1(_492_v13, 13)
						(_nw83).ArraySet1((func() _dafny.Array {
							if p0 {
								return _492_v13
							}
							return _492_v13
						})(), 14)
						(_nw83).ArraySet1(_492_v13, 15)
						(_nw83).ArraySet1(_492_v13, 16)
						(_nw83).ArraySet1(_492_v13, 17)
						(_nw83).ArraySet1((_493_v14).Select((Companion_Default___.SafeIndex(_480_v1, _dafny.IntOfUint32((_493_v14).Cardinality()))).Uint32()).(_dafny.Array), 18)
						_494_v15 = _nw83
						var _495_v16 _dafny.Map
						_ = _495_v16
						_495_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_485_v6, _484_v5)
						var _496_v17 _dafny.Sequence
						_ = _496_v17
						_496_v17 = _dafny.UnicodeSeqOfUtf8Bytes("k")
						var _497_v18 _dafny.Set
						_ = _497_v18
						_497_v18 = _dafny.SetOf((_this).F20())
						var _498_v19 _dafny.Array
						_ = _498_v19
						var _nwElement0_19 _dafny.Int = _480_v1
						_ = _nwElement0_19
						var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(12))
						_ = _nw84
						(_nw84).ArraySet1(_nwElement0_19, 0)
						(_nw84).ArraySet1(_485_v6, 1)
						(_nw84).ArraySet1(_485_v6, 2)
						(_nw84).ArraySet1(_480_v1, 3)
						(_nw84).ArraySet1(_dafny.IntOfInt64(840), 4)
						(_nw84).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_480_v1)), 5)
						(_nw84).ArraySet1((_483_v4).Cardinality(), 6)
						(_nw84).ArraySet1(_dafny.IntOfUint32((_496_v17).Cardinality()), 7)
						(_nw84).ArraySet1((_497_v18).Cardinality(), 8)
						(_nw84).ArraySet1(_485_v6, 9)
						(_nw84).ArraySet1(_485_v6, 10)
						(_nw84).ArraySet1(Companion_Default___.Fm0(globalState), 11)
						_498_v19 = _nw84
						var _499_v20 _dafny.Map
						_ = _499_v20
						_499_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if (_495_v16).Contains(_485_v6) {
								return (_495_v16).Get(_485_v6).(bool)
							}
							return _484_v5
						})(), _498_v19)
						var _rhs85 D1 = _490_v11
						_ = _rhs85
						var _rhs86 bool = (_499_v20).Contains((_485_v6).Cmp(_480_v1) != 0)
						_ = _rhs86
						var _rhs87 _dafny.Array = _494_v15
						_ = _rhs87
						var _rhs88 bool = (_485_v6).Cmp(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_480_v1), _480_v1)) == 0
						_ = _rhs88
						var _rhs89 _dafny.Array = _479_v0
						_ = _rhs89
						_490_v11 = _rhs85
						_486_v7 = _rhs86
						_494_v15 = _rhs87
						_484_v5 = _rhs88
						_479_v0 = _rhs89
					} else {
						var _500_v21 _dafny.Sequence
						_ = _500_v21
						_500_v21 = _dafny.UnicodeSeqOfUtf8Bytes("dcfqu")
						var _rhs90 _dafny.Sequence = _500_v21
						_ = _rhs90
						var _rhs91 _dafny.Int = (func() _dafny.Int {
							if p0 {
								return (_480_v1).Minus(_480_v1)
							}
							return Companion_Default___.Fm0(globalState)
						})()
						_ = _rhs91
						var _lhs59 *GlobalState = globalState
						_ = _lhs59
						_500_v21 = _rhs90
						_lhs59.F2 = _rhs91
						var _501_v22 _dafny.Map
						_ = _501_v22
						_501_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_479_v0, p0)
						(globalState).F0 = !((_501_v22).Merge(_501_v22)).Equals(_501_v22)
						var _502_v23 _dafny.Int
						_ = _502_v23
						_502_v23 = (_dafny.Zero).Minus(Companion_Default___.Fm0(globalState))
						var _503_v24 *C0
						_ = _503_v24
						var _nw85 *C0 = New_C0_()
						_ = _nw85
						_nw85.Ctor__(_502_v23, _480_v1)
						_503_v24 = _nw85
						var _504_v25 *C0
						_ = _504_v25
						var _nw86 *C0 = New_C0_()
						_ = _nw86
						_nw86.Ctor__(_502_v23, _480_v1)
						_504_v25 = _nw86
						(globalState).F2 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_480_v1))
					}
					if (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-85), (_dafny.Zero).Minus(_dafny.IntOfInt64(-525)))).Cmp(_480_v1) < 0 {
						(globalState).F10 = (_482_v3).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_480_v1, (_dafny.Zero).Minus(_480_v1)), _dafny.IntOfUint32((_482_v3).Cardinality()))).Uint32()).(bool)
						var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_479_v0), 0))
						_ = _index61
						(_479_v0).ArraySet1(p0, (_index61).Int())
						var _505_v26 _dafny.Sequence
						_ = _505_v26
						_505_v26 = _dafny.UnicodeSeqOfUtf8Bytes("t")
						var _506_v27 _dafny.Map
						_ = _506_v27
						_506_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nogbhr"), (Companion_Default___.SafeIndex(_480_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nogbhr")).Cardinality()))).Uint32(), (_this).F20()), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm6(Companion_Default___.Fm1(p0, _505_v26, globalState), globalState), (Companion_Default___.SafeIndex(_480_v1, _dafny.IntOfUint32((Companion_Default___.Fm6(Companion_Default___.Fm1(p0, _505_v26, globalState), globalState)).Cardinality()))).Uint32(), Companion_Default___.Fm7((_this).F20(), p0, globalState)))).Cardinality()), (_dafny.SetOf(_480_v1)).IsDisjointFrom(Companion_Default___.Fm8((_dafny.MultiSetFromSeq(_505_v26)).Cardinality(), _480_v1, globalState)))
						var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_479_v0), 0))
						_ = _index62
						(_479_v0).ArraySet1(p0, (_index62).Int())
						var _507_v28 _dafny.Map
						_ = _507_v28
						_507_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
						var _508_v29 _dafny.Map
						_ = _508_v29
						_508_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_507_v28, (_dafny.Zero).Minus(_480_v1))
						var _509_v30 _dafny.MultiSet
						_ = _509_v30
						_509_v30 = _dafny.MultiSetOf(_480_v1)
						var _510_v31 _dafny.Set
						_ = _510_v31
						_510_v31 = _dafny.SetOf(_480_v1)
						var _511_v32 _dafny.Sequence
						_ = _511_v32
						_511_v32 = _dafny.SeqOf((_510_v31).Cardinality(), _480_v1, _dafny.IntOfUint32((_482_v3).Cardinality()), _480_v1, _480_v1)
						var _512_v34 _dafny.Map
						_ = _512_v34
						_512_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
							if (_509_v30).Contains(_480_v1) {
								return (_509_v30).Multiplicity(_480_v1)
							}
							return _dafny.IntOfUint32((_511_v32).Cardinality())
						})(), func() _dafny.Map {
							var _coll17 = _dafny.NewMapBuilder()
							_ = _coll17
							for _iter18 := _dafny.Iterate((_506_v27).Keys().Elements()); ; {
								_compr_17, _ok18 := _iter18()
								if !_ok18 {
									break
								}
								var _513_v33 _dafny.Int
								_513_v33 = interface{}(_compr_17).(_dafny.Int)
								if (_506_v27).Contains(_513_v33) {
									_coll17.Add((_513_v33).Times(_dafny.IntOfInt64(214)), p0)
								}
							}
							return _coll17.ToMap()
						}())
						var _514_v35 _dafny.Set
						_ = _514_v35
						_514_v35 = _dafny.SetOf((func() bool {
							if (_507_v28).Contains(p0) {
								return (_507_v28).Get(p0).(bool)
							}
							return p0
						})(), p0, p0)
						var _515_v36 D1
						_ = _515_v36
						_515_v36 = Companion_D1_.Create_DC2_(p0)
						var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_479_v0), 0))
						_ = _index63
						var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_479_v0), 0))
						_ = _index64
						var _rhs92 bool = !(_508_v29).Contains(_507_v28)
						_ = _rhs92
						var _rhs93 _dafny.Map = (func() _dafny.Map {
							if p0 {
								return _506_v27
							}
							return (func() _dafny.Map {
								if (_512_v34).Contains(_480_v1) {
									return (_512_v34).Get(_480_v1).(_dafny.Map)
								}
								return _506_v27
							})()
						})()
						_ = _rhs93
						var _rhs94 bool = p0
						_ = _rhs94
						var _rhs95 bool = !((_514_v35).Union(_dafny.SetOf((_515_v36).Dtor_cf2(), p0, p0))).Contains(p0)
						_ = _rhs95
						var _rhs96 _dafny.Int = Companion_Default___.Fm0(globalState)
						_ = _rhs96
						var _lhs60 _dafny.Array = _479_v0
						_ = _lhs60
						var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_479_v0), 0))
						_ = _lhs61
						var _lhs62 _dafny.Array = _479_v0
						_ = _lhs62
						var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_479_v0), 0))
						_ = _lhs63
						var _lhs64 *GlobalState = globalState
						_ = _lhs64
						var _lhs65 *GlobalState = globalState
						_ = _lhs65
						(_lhs60).ArraySet1(_rhs92, (_lhs61).Int())
						_506_v27 = _rhs93
						(_lhs62).ArraySet1(_rhs94, (_lhs63).Int())
						_lhs64.F0 = _rhs95
						_lhs65.F9 = _rhs96
						(globalState).F0 = (_515_v36).Dtor_cf2()
						var _516_v37 _dafny.Map
						_ = _516_v37
						_516_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), p0)
						_516_v37 = (_516_v37).Update((_this).F20(), p0)
						var _517_v38 _dafny.Set
						_ = _517_v38
						_517_v38 = _dafny.SetOf((_506_v27).Merge(_506_v27))
						_517_v38 = _517_v38
					} else {
						var _518_v39 _dafny.Int
						_ = _518_v39
						_518_v39 = _480_v1
						var _519_v40 *C0
						_ = _519_v40
						var _nw87 *C0 = New_C0_()
						_ = _nw87
						_nw87.Ctor__(_518_v39, _480_v1)
						_519_v40 = _nw87
						_519_v40 = _519_v40
						(globalState).F0 = p0
						(globalState).F9 = _480_v1
						(globalState).F2 = _480_v1
						var _520_v41 D1
						_ = _520_v41
						_520_v41 = Companion_D1_.Create_DC1_(_479_v0)
						var _521_v42 _dafny.Map
						_ = _521_v42
						_521_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _520_v41)
						var _522_v43 _dafny.Sequence
						_ = _522_v43
						_522_v43 = _dafny.UnicodeSeqOfUtf8Bytes("im")
						var _523_v44 _dafny.Array
						_ = _523_v44
						var _nwElement0_20 _dafny.Sequence = _522_v43
						_ = _nwElement0_20
						var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(2))
						_ = _nw88
						(_nw88).ArraySet1(_nwElement0_20, 0)
						(_nw88).ArraySet1(_522_v43, 1)
						_523_v44 = _nw88
						var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_523_v44), 0))
						_ = _index65
						(_523_v44).ArraySet1(_522_v43, (_index65).Int())
						var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_523_v44), 0))
						_ = _index66
						var _rhs97 bool = p0
						_ = _rhs97
						var _rhs98 _dafny.Int = Companion_Default___.SafeDivisionInt(_480_v1, _dafny.IntOfInt64(256))
						_ = _rhs98
						var _rhs99 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _520_v41)
						_ = _rhs99
						var _rhs100 _dafny.Sequence = Companion_Default___.Fm6(p0, globalState)
						_ = _rhs100
						var _lhs66 *GlobalState = globalState
						_ = _lhs66
						var _lhs67 *GlobalState = globalState
						_ = _lhs67
						var _lhs68 _dafny.Array = _523_v44
						_ = _lhs68
						var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_523_v44), 0))
						_ = _lhs69
						_lhs66.F3 = _rhs97
						_lhs67.F9 = _rhs98
						_521_v42 = _rhs99
						(_lhs68).ArraySet1(_rhs100, (_lhs69).Int())
					}
					var _524_v45 _dafny.Int
					_ = _524_v45
					_524_v45 = _480_v1
					var _525_v46 *C0
					_ = _525_v46
					var _nw89 *C0 = New_C0_()
					_ = _nw89
					_nw89.Ctor__(_524_v45, _480_v1)
					_525_v46 = _nw89
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		(globalState).F3 = p0
		var _526_v47 _dafny.Int
		_ = _526_v47
		_526_v47 = _dafny.IntOfInt64(825)
		r0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_526_v47), _dafny.IntOfInt64(127))
		var _527_i1 _dafny.Int
		_ = _527_i1
		_527_i1 = _dafny.Zero
		{
			for !(p0) {
				{
					if (_527_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_527_i1 = (_527_i1).Plus(_dafny.One)
					var _528_v48 _dafny.Map
					_ = _528_v48
					_528_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					_528_v48 = (_528_v48).Merge(_528_v48)
					(globalState).F10 = (!(p0)) && (((_dafny.Zero).Minus(_dafny.IntOfInt64(-389))).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(p0, p0, p0)).Cardinality())) > 0)
					var _529_v49 _dafny.Sequence
					_ = _529_v49
					_529_v49 = _dafny.UnicodeSeqOfUtf8Bytes("tht")
					_529_v49 = _dafny.Companion_Sequence_.Concatenate(_529_v49, _529_v49)
					var _530_v50 _dafny.Int
					_ = _530_v50
					_530_v50 = _526_v47
					var _531_v51 _dafny.Map
					_ = _531_v51
					_531_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_530_v50), Companion_Default___.Fm1(p0, _529_v49, globalState))
					var _532_v52 _dafny.Int
					_ = _532_v52
					var _533_v53 _dafny.Int
					_ = _533_v53
					var _out22 _dafny.Int
					_ = _out22
					var _out23 _dafny.Int
					_ = _out23
					_out22, _out23 = Companion_Default___.M0(p0, _531_v51, globalState)
					_532_v52 = _out22
					_533_v53 = _out23
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _534_v54 _dafny.Array
		_ = _534_v54
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw90
		_534_v54 = _nw90
		for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_534_v54), 0))); ; {
			_guard_loop_1, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _535_i2 _dafny.Int
			_535_i2 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_535_i2).Sign() != -1) && ((_535_i2).Cmp(_dafny.ArrayLen((_534_v54), 0)) < 0)) {
				(_534_v54).ArraySet1(!(p0), (_535_i2).Int())
			}
		}
		(globalState).F2 = _526_v47
		var _536_v55 _dafny.Sequence
		_ = _536_v55
		_536_v55 = _dafny.UnicodeSeqOfUtf8Bytes("iakx")
		r0 = (func() _dafny.Int {
			if p0 {
				return _dafny.IntOfUint32((_536_v55).Cardinality())
			}
			return _526_v47
		})()
		var _537_v56 _dafny.MultiSet
		_ = _537_v56
		_537_v56 = _dafny.MultiSetOf(p0)
		var _538_v57 _dafny.Set
		_ = _538_v57
		_538_v57 = _dafny.SetOf(_526_v47)
		var _539_v58 _dafny.Sequence
		_ = _539_v58
		_539_v58 = _dafny.SeqOf(p0, p0, p0, Companion_Default___.Fm1(p0, _536_v55, globalState), p0)
		var _540_v59 _dafny.MultiSet
		_ = _540_v59
		_540_v59 = _dafny.MultiSetOf(_537_v56, _537_v56, (func() _dafny.MultiSet {
			if p0 {
				return Companion_Default___.Fm9((_538_v57).Cardinality(), p0, p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_539_v58).Cardinality()))), globalState)
			}
			return _dafny.MultiSetFromSeq(_539_v58)
		})(), (_537_v56).Update(p0, Companion_Default___.Abs((_538_v57).Cardinality())))
		r1 = _540_v59
		return r0, r1
	}
}
func (_this *C1) M4(globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _541_v0 bool
		_ = _541_v0
		_541_v0 = false
		if _541_v0 {
			var _542_v1 _dafny.Int
			_ = _542_v1
			_542_v1 = _dafny.IntOfInt64(785)
			var _543_v2 _dafny.Sequence
			_ = _543_v2
			_543_v2 = _dafny.SeqOf(_541_v0)
			var _544_v3 _dafny.MultiSet
			_ = _544_v3
			_544_v3 = _dafny.MultiSetOf(_541_v0, (_542_v1).Cmp(_542_v1) == 0, (_543_v2).Select((Companion_Default___.SafeIndex(_542_v1, _dafny.IntOfUint32((_543_v2).Cardinality()))).Uint32()).(bool))
			var _545_v4 _dafny.Sequence
			_ = _545_v4
			_545_v4 = _dafny.SeqOf(_542_v1, (_dafny.Zero).Minus(_542_v1))
			var _546_v5 _dafny.Sequence
			_ = _546_v5
			_546_v5 = _dafny.UnicodeSeqOfUtf8Bytes("hfyg")
			var _rhs101 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if _541_v0 {
					return _dafny.SeqOf(_541_v0, _541_v0)
				}
				return _543_v2
			})(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _541_v0 {
					return _dafny.SeqOf(_541_v0, _541_v0)
				}
				return _543_v2
			})()).Cardinality()))).Uint32(), _541_v0))).Union(_544_v3)
			_ = _rhs101
			var _rhs102 bool = _541_v0
			_ = _rhs102
			var _rhs103 _dafny.Int = (Companion_Default___.SafeDivisionInt(_542_v1, _dafny.IntOfUint32((_545_v4).Cardinality()))).Minus((_dafny.Zero).Minus((_542_v1).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_546_v5).Cardinality())))))
			_ = _rhs103
			var _lhs70 *GlobalState = globalState
			_ = _lhs70
			var _lhs71 *GlobalState = globalState
			_ = _lhs71
			_544_v3 = _rhs101
			_lhs70.F3 = _rhs102
			_lhs71.F2 = _rhs103
			var _547_v6 _dafny.Array
			_ = _547_v6
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(16))
			_ = _nw91
			_547_v6 = _nw91
			var _548_v7 D2
			_ = _548_v7
			_548_v7 = Companion_D2_.Create_DC4_(_547_v6)
			_547_v6 = (_548_v7).Dtor_cf4()
			var _549_v8 _dafny.Array
			_ = _549_v8
			var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw92
			_549_v8 = _nw92
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_549_v8), 0))
			_ = _index67
			(_549_v8).ArraySet1((_542_v1).Minus(_542_v1), (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_549_v8), 0))
			_ = _index68
			(_549_v8).ArraySet1(_542_v1, (_index68).Int())
			(globalState).F0 = false
			var _550_v9 _dafny.Array
			_ = _550_v9
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw93
			_550_v9 = _nw93
			var _551_v10 _dafny.Sequence
			_ = _551_v10
			_551_v10 = _dafny.SeqOf(_550_v9, _550_v9)
			var _552_v11 _dafny.Sequence
			_ = _552_v11
			_552_v11 = _dafny.SeqOf(_550_v9, (_551_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-315), _dafny.IntOfUint32((_551_v10).Cardinality()))).Uint32()).(_dafny.Array))
			var _553_v12 _dafny.Map
			_ = _553_v12
			_553_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _550_v9)
			var _554_v13 _dafny.Array
			_ = _554_v13
			var _nwElement0_21 _dafny.Array = _550_v9
			_ = _nwElement0_21
			var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(28))
			_ = _nw94
			(_nw94).ArraySet1(_nwElement0_21, 0)
			(_nw94).ArraySet1(_550_v9, 1)
			(_nw94).ArraySet1(_550_v9, 2)
			(_nw94).ArraySet1(_550_v9, 3)
			(_nw94).ArraySet1(_550_v9, 4)
			(_nw94).ArraySet1(_550_v9, 5)
			(_nw94).ArraySet1(_550_v9, 6)
			(_nw94).ArraySet1(_550_v9, 7)
			(_nw94).ArraySet1(_550_v9, 8)
			(_nw94).ArraySet1(_550_v9, 9)
			(_nw94).ArraySet1(_550_v9, 10)
			(_nw94).ArraySet1((_552_v11).Select((Companion_Default___.SafeIndex(_542_v1, _dafny.IntOfUint32((_552_v11).Cardinality()))).Uint32()).(_dafny.Array), 11)
			(_nw94).ArraySet1(_550_v9, 12)
			(_nw94).ArraySet1(_550_v9, 13)
			(_nw94).ArraySet1((func() _dafny.Array {
				if true {
					return _550_v9
				}
				return _550_v9
			})(), 14)
			(_nw94).ArraySet1(_550_v9, 15)
			(_nw94).ArraySet1(_550_v9, 16)
			(_nw94).ArraySet1(_550_v9, 17)
			(_nw94).ArraySet1(_550_v9, 18)
			(_nw94).ArraySet1(_550_v9, 19)
			(_nw94).ArraySet1((func() _dafny.Array {
				if true {
					return (_552_v11).Select((Companion_Default___.SafeIndex(_542_v1, _dafny.IntOfUint32((_552_v11).Cardinality()))).Uint32()).(_dafny.Array)
				}
				return _550_v9
			})(), 20)
			(_nw94).ArraySet1(_550_v9, 21)
			(_nw94).ArraySet1(_550_v9, 22)
			(_nw94).ArraySet1(_550_v9, 23)
			(_nw94).ArraySet1(_550_v9, 24)
			(_nw94).ArraySet1(_550_v9, 25)
			(_nw94).ArraySet1((func() _dafny.Array {
				if (_553_v12).Contains(_541_v0) {
					return (_553_v12).Get(_541_v0).(_dafny.Array)
				}
				return _550_v9
			})(), 26)
			(_nw94).ArraySet1(_550_v9, 27)
			_554_v13 = _nw94
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_554_v13), 0))
			_ = _index69
			(_554_v13).ArraySet1(_550_v9, (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_554_v13), 0))
			_ = _index70
			(_554_v13).ArraySet1(_550_v9, (_index70).Int())
		} else {
			var _555_v14 _dafny.Int
			_ = _555_v14
			_555_v14 = _dafny.IntOfInt64(-271)
			var _556_v15 _dafny.Sequence
			_ = _556_v15
			_556_v15 = _dafny.SeqOf(_555_v14, _555_v14, _555_v14)
			(globalState).F0 = (Companion_Default___.SafeModuloInt(_555_v14, _555_v14)).Cmp((_556_v15).Select((Companion_Default___.SafeIndex(_555_v14, _dafny.IntOfUint32((_556_v15).Cardinality()))).Uint32()).(_dafny.Int)) >= 0
			var _557_v16 _dafny.MultiSet
			_ = _557_v16
			_557_v16 = _dafny.MultiSetOf(_541_v0, _541_v0)
			var _558_v17 _dafny.Map
			_ = _558_v17
			_558_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _557_v16)
			var _559_v18 _dafny.Int
			_ = _559_v18
			_559_v18 = (_558_v17).Cardinality()
			var _source8 _dafny.Int = _559_v18
			_ = _source8
			var _560___mcc_h0 _dafny.Int = _source8
			_ = _560___mcc_h0
			var _561_cf0 _dafny.Int = _560___mcc_h0
			_ = _561_cf0
			var _562_v19 *C0
			_ = _562_v19
			var _nw95 *C0 = New_C0_()
			_ = _nw95
			_nw95.Ctor__(_559_v18, _dafny.IntOfInt64(90))
			_562_v19 = _nw95
			var _563_v20 _dafny.Map
			_ = _563_v20
			_563_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_541_v0), _dafny.IntOfInt64(-251))
			var _564_v21 _dafny.Sequence
			_ = _564_v21
			_564_v21 = _dafny.SeqOf(_563_v20)
			_564_v21 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(787))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_565_v20 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_566_i0 _dafny.Int) _dafny.Map {
					return _565_v20
				}
			})(_563_v20)))
			var _567_v22 _dafny.CodePoint
			_ = _567_v22
			_567_v22 = _dafny.CodePoint('h')
			var _568_v23 _dafny.Sequence
			_ = _568_v23
			_568_v23 = _dafny.UnicodeSeqOfUtf8Bytes("nntcghypc")
			var _569_v24 _dafny.Array
			_ = _569_v24
			var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw96
			_569_v24 = _nw96
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_569_v24), 0))
			_ = _index71
			(_569_v24).ArraySet1(_561_cf0, (_index71).Int())
			var _570_v25 _dafny.Array
			_ = _570_v25
			var _nw97 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw97
			_570_v25 = _nw97
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_570_v25), 0))
			_ = _index72
			(_570_v25).ArraySet1((_557_v16).IsSubsetOf(_557_v16), (_index72).Int())
			var _571_v26 _dafny.Sequence
			_ = _571_v26
			_571_v26 = _dafny.SeqOf(_570_v25)
			var _572_v27 _dafny.Map
			_ = _572_v27
			_572_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_571_v26, _567_v22)
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_569_v24), 0))
			_ = _index73
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_570_v25), 0))
			_ = _index74
			var _rhs104 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_572_v27).Contains(_571_v26) {
					return (_572_v27).Get(_571_v26).(_dafny.CodePoint)
				}
				return _567_v22
			})()
			_ = _rhs104
			var _rhs105 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(597))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_573_v23 _dafny.Sequence, _574_cf0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_575_i1 _dafny.Int) _dafny.CodePoint {
					return (_573_v23).Select((Companion_Default___.SafeIndex(_574_cf0, _dafny.IntOfUint32((_573_v23).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_568_v23, _561_cf0))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(657))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_576_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_577_i2 _dafny.Int) _dafny.CodePoint {
					return _576_v22
				}
			})(_567_v22))))
			_ = _rhs105
			var _rhs106 _dafny.Int = _561_cf0
			_ = _rhs106
			var _rhs107 bool = !(_541_v0)
			_ = _rhs107
			var _lhs72 _dafny.Array = _569_v24
			_ = _lhs72
			var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_569_v24), 0))
			_ = _lhs73
			var _lhs74 _dafny.Array = _570_v25
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_570_v25), 0))
			_ = _lhs75
			_567_v22 = _rhs104
			_568_v23 = _rhs105
			(_lhs72).ArraySet1(_rhs106, (_lhs73).Int())
			(_lhs74).ArraySet1(_rhs107, (_lhs75).Int())
			_563_v20 = (_563_v20).Update(Companion_Default___.Fm1(_541_v0, _568_v23, globalState), (_569_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_569_v24), 0))).Int()).(_dafny.Int))
			var _578_v28 *C0
			_ = _578_v28
			var _nw98 *C0 = New_C0_()
			_ = _nw98
			_nw98.Ctor__(_559_v18, _555_v14)
			_578_v28 = _nw98
			var _579_v29 _dafny.MultiSet
			_ = _579_v29
			_579_v29 = _dafny.MultiSetOf(_578_v28)
			var _580_v30 _dafny.MultiSet
			_ = _580_v30
			_580_v30 = _579_v29
			var _581_v31 _dafny.Map
			_ = _581_v31
			_581_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_555_v14, _555_v14)
			var _582_v32 _dafny.Array
			_ = _582_v32
			var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw99
			_582_v32 = _nw99
			var _583_v33 _dafny.MultiSet
			_ = _583_v33
			_583_v33 = _dafny.MultiSetOf((Companion_Default___.Fm8(_555_v14, _555_v14, globalState)).Cardinality(), _555_v14)
			var _584_v34 D2
			_ = _584_v34
			_584_v34 = Companion_D2_.Create_DC5_(_541_v0, (_581_v31).Cardinality(), _582_v32, _583_v33)
			var _585_v35 _dafny.Sequence
			_ = _585_v35
			_585_v35 = _dafny.SeqOf(_541_v0, !(_541_v0), _541_v0)
			var _586_v36 _dafny.Array
			_ = _586_v36
			var _nwElement0_22 _dafny.Sequence = _dafny.SeqOf((_584_v34).Dtor_cf5(), _541_v0)
			_ = _nwElement0_22
			var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(16))
			_ = _nw100
			(_nw100).ArraySet1(_nwElement0_22, 0)
			(_nw100).ArraySet1(_dafny.SeqOf(_541_v0, _541_v0), 1)
			(_nw100).ArraySet1(_585_v35, 2)
			(_nw100).ArraySet1(_585_v35, 3)
			(_nw100).ArraySet1(_585_v35, 4)
			(_nw100).ArraySet1(_585_v35, 5)
			(_nw100).ArraySet1(_585_v35, 6)
			(_nw100).ArraySet1(_585_v35, 7)
			(_nw100).ArraySet1(_585_v35, 8)
			(_nw100).ArraySet1(_585_v35, 9)
			(_nw100).ArraySet1(_585_v35, 10)
			(_nw100).ArraySet1(_585_v35, 11)
			(_nw100).ArraySet1(_585_v35, 12)
			(_nw100).ArraySet1(_585_v35, 13)
			(_nw100).ArraySet1(_585_v35, 14)
			(_nw100).ArraySet1(_585_v35, 15)
			_586_v36 = _nw100
			var _587_v37 _dafny.Map
			_ = _587_v37
			_587_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_580_v30).Union(_579_v29), _586_v36)
			_587_v37 = (_587_v37).Update(_579_v29, _586_v36)
			var _588_v38 _dafny.Sequence
			_ = _588_v38
			_588_v38 = _dafny.UnicodeSeqOfUtf8Bytes("eancxn")
			_588_v38 = _588_v38
			(globalState).F0 = ((_555_v14).Times(_555_v14)).Cmp((_dafny.Zero).Minus(_555_v14)) < 0
		}
		var _589_i3 _dafny.Int
		_ = _589_i3
		_589_i3 = _dafny.Zero
		{
			for true {
				{
					if (_589_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_589_i3 = (_589_i3).Plus(_dafny.One)
					var _590_v39 _dafny.Sequence
					_ = _590_v39
					_590_v39 = _dafny.SeqOf(_541_v0)
					var _591_v40 _dafny.Sequence
					_ = _591_v40
					_591_v40 = _dafny.UnicodeSeqOfUtf8Bytes("htdkpplcq")
					var _592_v41 _dafny.Array
					_ = _592_v41
					var _len0_12 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_12
					var _nw101 _dafny.Array
					_ = _nw101
					if _len0_12.Cmp(_dafny.Zero) == 0 {
						_nw101 = _dafny.NewArray(_len0_12)
					} else {
						var _init12 func(_dafny.Int) _dafny.Int = func(_593_i4 _dafny.Int) _dafny.Int {
							return (_593_i4).Plus(_dafny.IntOfInt64(221))
						}
						_ = _init12
						var _element0_12 = _init12(_dafny.Zero)
						_ = _element0_12
						_nw101 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
						(_nw101).ArraySet1(_element0_12, 0)
						var _nativeLen0_12 = (_len0_12).Int()
						_ = _nativeLen0_12
						for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
							(_nw101).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
						}
					}
					_592_v41 = _nw101
					var _594_v42 _dafny.Map
					_ = _594_v42
					_594_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _592_v41)
					if (_590_v39).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_591_v40).Cardinality())), (_594_v42).Cardinality()), _dafny.IntOfUint32((_590_v39).Cardinality()))).Uint32()).(bool) {
						var _595_v43 _dafny.Set
						_ = _595_v43
						_595_v43 = _dafny.SetOf(_541_v0)
						var _596_v44 _dafny.Int
						_ = _596_v44
						_596_v44 = _dafny.IntOfInt64(192)
						var _597_v45 _dafny.Map
						_ = _597_v45
						_597_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _596_v44)
						var _598_v46 _dafny.Int
						_ = _598_v46
						_598_v46 = (_597_v45).Cardinality()
						var _599_v47 _dafny.Map
						_ = _599_v47
						_599_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_595_v43, (_590_v39).Select((Companion_Default___.SafeIndex((_598_v46), _dafny.IntOfUint32((_590_v39).Cardinality()))).Uint32()).(bool))
						_599_v47 = _599_v47
						(globalState).F3 = (_541_v0) || (_541_v0)
						r3 = _541_v0
						(globalState).F10 = _541_v0
						var _600_v48 _dafny.Sequence
						_ = _600_v48
						_600_v48 = _dafny.SeqOf(_596_v44, _596_v44)
						var _601_v49 _dafny.MultiSet
						_ = _601_v49
						_601_v49 = _dafny.MultiSetOf((_dafny.SetOf((_this).F20(), (_591_v40).Select((Companion_Default___.SafeIndex(_596_v44, _dafny.IntOfUint32((_591_v40).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality())
						var _602_v50 _dafny.Array
						_ = _602_v50
						var _nw102 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
						_ = _nw102
						_602_v50 = _nw102
						var _603_v51 D2
						_ = _603_v51
						_603_v51 = Companion_D2_.Create_DC5_(_541_v0, _596_v44, _602_v50, _dafny.MultiSetOf(_596_v44))
						var _604_v52 _dafny.Array
						_ = _604_v52
						var _nwElement0_23 bool = (_601_v49).IsSubsetOf(_dafny.MultiSetFromSeq(_600_v48))
						_ = _nwElement0_23
						var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(5))
						_ = _nw103
						(_nw103).ArraySet1(_nwElement0_23, 0)
						(_nw103).ArraySet1((_590_v39).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_596_v44), _dafny.IntOfUint32((_590_v39).Cardinality()))).Uint32()).(bool), 1)
						(_nw103).ArraySet1(_541_v0, 2)
						(_nw103).ArraySet1(_541_v0, 3)
						(_nw103).ArraySet1((_603_v51).Dtor_cf5(), 4)
						_604_v52 = _nw103
						var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_604_v52), 0))
						_ = _index75
						(_604_v52).ArraySet1(_541_v0, (_index75).Int())
						var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_604_v52), 0))
						_ = _index76
						(_604_v52).ArraySet1(_541_v0, (_index76).Int())
					} else {
						_591_v40 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(764))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg24 _dafny.Int) interface{} {
								return coer24(arg24)
							}
						}(func(_605_i5 _dafny.Int) _dafny.CodePoint {
							return (_this).F20()
						})), _591_v40)
						var _606_v53 _dafny.Map
						_ = _606_v53
						_606_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_590_v39, _dafny.IntOfInt64(555))
						var _607_v54 _dafny.Int
						_ = _607_v54
						_607_v54 = (_606_v53).Cardinality()
						var _608_v55 *C0
						_ = _608_v55
						var _nw104 *C0 = New_C0_()
						_ = _nw104
						_nw104.Ctor__(_607_v54, _dafny.IntOfUint32((_591_v40).Cardinality()))
						_608_v55 = _nw104
						var _609_v56 _dafny.Sequence
						_ = _609_v56
						_609_v56 = _dafny.SeqOf(_608_v55, _608_v55, _608_v55)
						var _610_v57 _dafny.Int
						_ = _610_v57
						_610_v57 = _dafny.IntOfInt64(30)
						var _611_v58 D4
						_ = _611_v58
						_611_v58 = Companion_D4_.Create_DC9_(_608_v55)
						var _612_v59 _dafny.Array
						_ = _612_v59
						var _nwElement0_24 *C0 = _608_v55
						_ = _nwElement0_24
						var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(24))
						_ = _nw105
						(_nw105).ArraySet1(_nwElement0_24, 0)
						(_nw105).ArraySet1(_608_v55, 1)
						(_nw105).ArraySet1(_608_v55, 2)
						(_nw105).ArraySet1(_608_v55, 3)
						(_nw105).ArraySet1((_609_v56).Select((Companion_Default___.SafeIndex(_610_v57, _dafny.IntOfUint32((_609_v56).Cardinality()))).Uint32()).(*C0), 4)
						(_nw105).ArraySet1(_608_v55, 5)
						(_nw105).ArraySet1(_608_v55, 6)
						(_nw105).ArraySet1(_608_v55, 7)
						(_nw105).ArraySet1(_608_v55, 8)
						(_nw105).ArraySet1(_608_v55, 9)
						(_nw105).ArraySet1(_608_v55, 10)
						(_nw105).ArraySet1(_608_v55, 11)
						(_nw105).ArraySet1(_608_v55, 12)
						(_nw105).ArraySet1(_608_v55, 13)
						(_nw105).ArraySet1(_608_v55, 14)
						(_nw105).ArraySet1(_608_v55, 15)
						(_nw105).ArraySet1(_608_v55, 16)
						(_nw105).ArraySet1(_608_v55, 17)
						(_nw105).ArraySet1(_608_v55, 18)
						(_nw105).ArraySet1(_608_v55, 19)
						(_nw105).ArraySet1(_608_v55, 20)
						(_nw105).ArraySet1(_608_v55, 21)
						(_nw105).ArraySet1(_608_v55, 22)
						(_nw105).ArraySet1((_611_v58).Dtor_cf15(), 23)
						_612_v59 = _nw105
						var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_612_v59), 0))
						_ = _index77
						(_612_v59).ArraySet1(_608_v55, (_index77).Int())
						var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_612_v59), 0))
						_ = _index78
						(_612_v59).ArraySet1(_608_v55, (_index78).Int())
						(globalState).F2 = _610_v57
						var _613_v60 _dafny.Array
						_ = _613_v60
						var _nw106 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
						_ = _nw106
						_613_v60 = _nw106
						var _614_v61 _dafny.Map
						_ = _614_v61
						_614_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _613_v60)
						var _615_v62 _dafny.Array
						_ = _615_v62
						var _nwElement0_25 _dafny.Array = _613_v60
						_ = _nwElement0_25
						var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(17))
						_ = _nw107
						(_nw107).ArraySet1(_nwElement0_25, 0)
						(_nw107).ArraySet1(_613_v60, 1)
						(_nw107).ArraySet1(_613_v60, 2)
						(_nw107).ArraySet1(_613_v60, 3)
						(_nw107).ArraySet1(_613_v60, 4)
						(_nw107).ArraySet1(_613_v60, 5)
						(_nw107).ArraySet1(_613_v60, 6)
						(_nw107).ArraySet1((func() _dafny.Array {
							if (_614_v61).Contains(_541_v0) {
								return (_614_v61).Get(_541_v0).(_dafny.Array)
							}
							return _613_v60
						})(), 7)
						(_nw107).ArraySet1(_613_v60, 8)
						(_nw107).ArraySet1(_613_v60, 9)
						(_nw107).ArraySet1(_613_v60, 10)
						(_nw107).ArraySet1(_613_v60, 11)
						(_nw107).ArraySet1(_613_v60, 12)
						(_nw107).ArraySet1(_613_v60, 13)
						(_nw107).ArraySet1(_613_v60, 14)
						(_nw107).ArraySet1(_613_v60, 15)
						(_nw107).ArraySet1(_613_v60, 16)
						_615_v62 = _nw107
						var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_615_v62), 0))
						_ = _index79
						(_615_v62).ArraySet1(_613_v60, (_index79).Int())
						var _616_v63 _dafny.MultiSet
						_ = _616_v63
						_616_v63 = _dafny.MultiSetOf(_610_v57, _610_v57, _610_v57)
						var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_615_v62), 0))
						_ = _index80
						var _rhs108 _dafny.Int = (_616_v63).Cardinality()
						_ = _rhs108
						var _rhs109 _dafny.Array = _613_v60
						_ = _rhs109
						var _lhs76 _dafny.Array = _615_v62
						_ = _lhs76
						var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_615_v62), 0))
						_ = _lhs77
						_610_v57 = _rhs108
						(_lhs76).ArraySet1(_rhs109, (_lhs77).Int())
						(globalState).F2 = _610_v57
					}
					_541_v0 = _541_v0
					var _617_v64 _dafny.Array
					_ = _617_v64
					var _nwElement0_26 bool = _541_v0
					_ = _nwElement0_26
					var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(28))
					_ = _nw108
					(_nw108).ArraySet1(_nwElement0_26, 0)
					(_nw108).ArraySet1(_541_v0, 1)
					(_nw108).ArraySet1(_541_v0, 2)
					(_nw108).ArraySet1(_541_v0, 3)
					(_nw108).ArraySet1(Companion_Default___.Fm1(true, _591_v40, globalState), 4)
					(_nw108).ArraySet1(_541_v0, 5)
					(_nw108).ArraySet1(Companion_Default___.Fm1(_541_v0, _591_v40, globalState), 6)
					(_nw108).ArraySet1(_541_v0, 7)
					(_nw108).ArraySet1(_541_v0, 8)
					(_nw108).ArraySet1(_541_v0, 9)
					(_nw108).ArraySet1(!(_541_v0), 10)
					(_nw108).ArraySet1(_541_v0, 11)
					(_nw108).ArraySet1(_541_v0, 12)
					(_nw108).ArraySet1(_541_v0, 13)
					(_nw108).ArraySet1(_541_v0, 14)
					(_nw108).ArraySet1(_541_v0, 15)
					(_nw108).ArraySet1(true, 16)
					(_nw108).ArraySet1(_541_v0, 17)
					(_nw108).ArraySet1(_541_v0, 18)
					(_nw108).ArraySet1(_541_v0, 19)
					(_nw108).ArraySet1(_541_v0, 20)
					(_nw108).ArraySet1(_541_v0, 21)
					(_nw108).ArraySet1(_541_v0, 22)
					(_nw108).ArraySet1(false, 23)
					(_nw108).ArraySet1(_541_v0, 24)
					(_nw108).ArraySet1(_541_v0, 25)
					(_nw108).ArraySet1(_541_v0, 26)
					(_nw108).ArraySet1(Companion_Default___.Fm1(_541_v0, _dafny.Companion_Sequence_.Update(_591_v40, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.IntOfUint32((_591_v40).Cardinality()))).Uint32(), (_this).F20()), globalState), 27)
					_617_v64 = _nw108
					var _618_v65 D1
					_ = _618_v65
					_618_v65 = Companion_D1_.Create_DC1_(_617_v64)
					var _619_v66 _dafny.Sequence
					_ = _619_v66
					_619_v66 = _dafny.SeqOf(_617_v64)
					var _620_v67 _dafny.Array
					_ = _620_v67
					var _nwElement0_27 _dafny.Array = _617_v64
					_ = _nwElement0_27
					var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(25))
					_ = _nw109
					(_nw109).ArraySet1(_nwElement0_27, 0)
					(_nw109).ArraySet1(_617_v64, 1)
					(_nw109).ArraySet1(_617_v64, 2)
					(_nw109).ArraySet1(_617_v64, 3)
					(_nw109).ArraySet1(_617_v64, 4)
					(_nw109).ArraySet1(_617_v64, 5)
					(_nw109).ArraySet1(_617_v64, 6)
					(_nw109).ArraySet1((func() _dafny.Array {
						if _541_v0 {
							return _617_v64
						}
						return _617_v64
					})(), 7)
					(_nw109).ArraySet1(_617_v64, 8)
					(_nw109).ArraySet1(_617_v64, 9)
					(_nw109).ArraySet1(_617_v64, 10)
					(_nw109).ArraySet1(_617_v64, 11)
					(_nw109).ArraySet1(_617_v64, 12)
					(_nw109).ArraySet1(_617_v64, 13)
					(_nw109).ArraySet1(_617_v64, 14)
					(_nw109).ArraySet1(_617_v64, 15)
					(_nw109).ArraySet1(_617_v64, 16)
					(_nw109).ArraySet1(_617_v64, 17)
					(_nw109).ArraySet1(_617_v64, 18)
					(_nw109).ArraySet1(_617_v64, 19)
					(_nw109).ArraySet1(_617_v64, 20)
					(_nw109).ArraySet1((_618_v65).Dtor_cf1(), 21)
					(_nw109).ArraySet1(_617_v64, 22)
					(_nw109).ArraySet1((func() _dafny.Array {
						if _541_v0 {
							return _617_v64
						}
						return (_619_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_590_v39).Cardinality()), _dafny.IntOfUint32((_619_v66).Cardinality()))).Uint32()).(_dafny.Array)
					})(), 23)
					(_nw109).ArraySet1(_617_v64, 24)
					_620_v67 = _nw109
					var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
					_ = _nw110
					_620_v67 = _nw110
					var _621_v68 _dafny.Int
					_ = _621_v68
					_621_v68 = _dafny.IntOfInt64(14)
					r2 = (Companion_Default___.SafeDivisionInt(_621_v68, (_dafny.SetOf(_591_v40, _591_v40)).Cardinality())).Cmp(_dafny.IntOfUint32((_590_v39).Cardinality())) > 0
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _622_v69 _dafny.Array
		_ = _622_v69
		var _nw111 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw111
		_622_v69 = _nw111
		var _623_v70 D1
		_ = _623_v70
		_623_v70 = Companion_D1_.Create_DC1_(_622_v69)
		var _624_v71 _dafny.Map
		_ = _624_v71
		_624_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_623_v70, _dafny.IntOfInt64(421))
		var _625_v72 _dafny.Int
		_ = _625_v72
		_625_v72 = _dafny.IntOfInt64(662)
		var _626_v73 _dafny.Map
		_ = _626_v73
		_626_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_624_v71).Contains(_623_v70) {
				return (_624_v71).Get(_623_v70).(_dafny.Int)
			}
			return (_dafny.Zero).Minus(_625_v72)
		})(), _625_v72)
		var _627_v74 _dafny.Sequence
		_ = _627_v74
		_627_v74 = _dafny.SeqOf((_this).F20())
		_626_v73 = (_626_v73).Update((_dafny.IntOfUint32((_627_v74).Cardinality())).Minus(_625_v72), _dafny.IntOfInt64(-351))
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))
		_ = _index81
		(_622_v69).ArraySet1(_541_v0, (_index81).Int())
		var _628_v75 _dafny.Sequence
		_ = _628_v75
		_628_v75 = _dafny.SeqOf(_625_v72)
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))
		_ = _index82
		(_622_v69).ArraySet1(!(false) || (_dafny.Companion_Sequence_.IsProperPrefixOf(_628_v75, _628_v75)), (_index82).Int())
		var _629_v76 _dafny.Sequence
		_ = _629_v76
		_629_v76 = _dafny.SeqOf((_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool))
		var _630_v77 D1
		_ = _630_v77
		_630_v77 = Companion_D1_.Create_DC2_(!((_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool)) || ((_629_v76).Select((Companion_Default___.SafeIndex(_625_v72, _dafny.IntOfUint32((_629_v76).Cardinality()))).Uint32()).(bool)))
		_630_v77 = func(_pat_let25_0 D1) D1 {
			return func(_631_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let26_0 bool) D1 {
					return func(_632_dt__update_hcf2_h0 bool) D1 {
						return Companion_D1_.Create_DC2_(_632_dt__update_hcf2_h0)
					}(_pat_let26_0)
				}(false)
			}(_pat_let25_0)
		}(_630_v77)
		var _633_v78 _dafny.Map
		_ = _633_v78
		_633_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_625_v72), (_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool))
		var _634_v79 _dafny.Int
		_ = _634_v79
		var _635_v80 _dafny.Int
		_ = _635_v80
		var _out24 _dafny.Int
		_ = _out24
		var _out25 _dafny.Int
		_ = _out25
		_out24, _out25 = Companion_Default___.M0((func() bool {
			if (_629_v76).Select((Companion_Default___.SafeIndex((_628_v75).Select((Companion_Default___.SafeIndex(_625_v72, _dafny.IntOfUint32((_628_v75).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_629_v76).Cardinality()))).Uint32()).(bool) {
				return (_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool)
			}
			return (_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool)
		})(), (_633_v78).Update(_625_v72, _541_v0), globalState)
		_634_v79 = _out24
		_635_v80 = _out25
		r0 = !((_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool))
		var _636_v81 D5
		_ = _636_v81
		_636_v81 = Companion_D5_.Create_DC12_(_627_v74)
		r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_636_v81).Dtor_cf18(), _dafny.Companion_Sequence_.Concatenate(_627_v74, _627_v74))).Cardinality())
		r2 = Companion_Default___.Fm1(_541_v0, _627_v74, globalState)
		r3 = !(((_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool)) == ((_622_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_622_v69), 0))).Int()).(bool))) || (!(_541_v0))
		return r0, r1, r2, r3
	}
}
func (_this *C1) F20() _dafny.CodePoint {
	{
		return _this._f20
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
func (_this *C2) M2(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _637_v0 _dafny.Int
		_ = _637_v0
		_637_v0 = _dafny.IntOfInt64(793)
		(globalState).F3 = (_637_v0).Cmp(_637_v0) == 0
		var _638_v1 _dafny.Int
		_ = _638_v1
		_638_v1 = _dafny.IntOfInt64(84)
		var _639_v2 _dafny.CodePoint
		_ = _639_v2
		_639_v2 = _dafny.CodePoint('f')
		var _640_v3 *C1
		_ = _640_v3
		var _nw112 *C1 = New_C1_()
		_ = _nw112
		_nw112.Ctor__(_639_v2)
		_640_v3 = _nw112
		var _641_v4 _dafny.MultiSet
		_ = _641_v4
		_641_v4 = _dafny.MultiSetOf(_640_v3)
		var _642_v5 _dafny.MultiSet
		_ = _642_v5
		_642_v5 = _dafny.MultiSetOf(_dafny.IntOfInt64(906))
		var _643_v6 *C0
		_ = _643_v6
		var _nw113 *C0 = New_C0_()
		_ = _nw113
		_nw113.Ctor__(_638_v1, ((_dafny.SetOf(_641_v4)).Cardinality()).Plus((func() _dafny.Int {
			if (_642_v5).Contains(_637_v0) {
				return (_642_v5).Multiplicity(_637_v0)
			}
			return _dafny.IntOfInt64(302)
		})()))
		_643_v6 = _nw113
		var _644_v7 bool
		_ = _644_v7
		_644_v7 = true
		var _645_v8 _dafny.Sequence
		_ = _645_v8
		_645_v8 = _dafny.UnicodeSeqOfUtf8Bytes("rfvwqmbea")
		var _646_v9 D5
		_ = _646_v9
		_646_v9 = Companion_D5_.Create_DC13_(_637_v0, _637_v0, Companion_Default___.Fm1(!(_644_v7), _645_v8, globalState))
		(globalState).F3 = func(_source9 D5) bool {
			if _source9.Is_DC13() {
				var _647___mcc_h0 _dafny.Int = _source9.Get_().(D5_DC13).Cf19
				_ = _647___mcc_h0
				var _648___mcc_h1 _dafny.Int = _source9.Get_().(D5_DC13).Cf20
				_ = _648___mcc_h1
				var _649___mcc_h2 bool = _source9.Get_().(D5_DC13).Cf21
				_ = _649___mcc_h2
				var _650_cf21 bool = _649___mcc_h2
				_ = _650_cf21
				var _651_cf20 _dafny.Int = _648___mcc_h1
				_ = _651_cf20
				var _652_cf19 _dafny.Int = _647___mcc_h0
				_ = _652_cf19
				return false
			} else {
				var _653___mcc_h3 _dafny.Sequence = _source9.Get_().(D5_DC12).Cf18
				_ = _653___mcc_h3
				var _654_cf18 _dafny.Sequence = _653___mcc_h3
				_ = _654_cf18
				return false
			}
		}(_646_v9)
		var _655_v10 _dafny.Array
		_ = _655_v10
		var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw114
		_655_v10 = _nw114
		for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_655_v10), 0))); ; {
			_guard_loop_2, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _656_i0 _dafny.Int
			_656_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_656_i0).Sign() != -1) && ((_656_i0).Cmp(_dafny.ArrayLen((_655_v10), 0)) < 0)) {
				(_655_v10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_644_v7), _dafny.SeqOf(_644_v7)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_644_v7, _644_v7), _dafny.SeqOf(_644_v7))), (_656_i0).Int())
			}
		}
		var _657_v11 D5
		_ = _657_v11
		_657_v11 = Companion_D5_.Create_DC12_(_645_v8)
		var _pat_let_tv20 = _644_v7
		_ = _pat_let_tv20
		if func(_source10 D5) bool {
			if _source10.Is_DC13() {
				var _658___mcc_h4 _dafny.Int = _source10.Get_().(D5_DC13).Cf19
				_ = _658___mcc_h4
				var _659___mcc_h5 _dafny.Int = _source10.Get_().(D5_DC13).Cf20
				_ = _659___mcc_h5
				var _660___mcc_h6 bool = _source10.Get_().(D5_DC13).Cf21
				_ = _660___mcc_h6
				var _661_cf21 bool = _660___mcc_h6
				_ = _661_cf21
				var _662_cf20 _dafny.Int = _659___mcc_h5
				_ = _662_cf20
				var _663_cf19 _dafny.Int = _658___mcc_h4
				_ = _663_cf19
				return false
			} else {
				var _664___mcc_h7 _dafny.Sequence = _source10.Get_().(D5_DC12).Cf18
				_ = _664___mcc_h7
				var _665_cf18 _dafny.Sequence = _664___mcc_h7
				_ = _665_cf18
				return !(_pat_let_tv20)
			}
		}(_657_v11) {
			_644_v7 = _644_v7
			var _666_v12 _dafny.Set
			_ = _666_v12
			_666_v12 = _dafny.SetOf(!(_644_v7), _644_v7, _644_v7)
			var _667_v13 _dafny.Array
			_ = _667_v13
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_13
			var _nw115 _dafny.Array
			_ = _nw115
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw115 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) bool = func(_668_i1 _dafny.Int) bool {
					return false
				}
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw115 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw115).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw115).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_667_v13 = _nw115
			var _669_v14 _dafny.Map
			_ = _669_v14
			_669_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if _644_v7 {
					return (_666_v12).Cardinality()
				}
				return _637_v0
			})(), _667_v13)
			_669_v14 = (_669_v14).Update(_dafny.IntOfInt64(377), _667_v13)
			if ((_637_v0).Plus(_637_v0)).Cmp(_637_v0) > 0 {
				(globalState).F2 = _637_v0
				(globalState).F10 = !(_dafny.Companion_Sequence_.Contains(_645_v8, (func() _dafny.CodePoint {
					if _644_v7 {
						return _639_v2
					}
					return (_640_v3).F20()
				})()))
				(globalState).F3 = (Companion_Default___.SafeDivisionInt(_637_v0, (_dafny.Zero).Minus(_637_v0))).Cmp(_637_v0) < 0
				r0 = _637_v0
				var _670_v15 D6
				_ = _670_v15
				_670_v15 = Companion_D6_.Create_DC14_(_640_v3)
				var _671_v16 _dafny.Array
				_ = _671_v16
				var _nwElement0_28 *C1 = (func() *C1 {
					if _644_v7 {
						return _640_v3
					}
					return (_670_v15).Dtor_cf22()
				})()
				_ = _nwElement0_28
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(9))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_28, 0)
				(_nw116).ArraySet1(_640_v3, 1)
				(_nw116).ArraySet1(_640_v3, 2)
				(_nw116).ArraySet1(_640_v3, 3)
				(_nw116).ArraySet1(_640_v3, 4)
				(_nw116).ArraySet1(_640_v3, 5)
				(_nw116).ArraySet1(_640_v3, 6)
				(_nw116).ArraySet1(_640_v3, 7)
				(_nw116).ArraySet1(_640_v3, 8)
				_671_v16 = _nw116
				_671_v16 = _671_v16
			} else {
				var _672_v17 _dafny.Array
				_ = _672_v17
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw117
				_672_v17 = _nw117
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_672_v17), 0))
				_ = _index83
				(_672_v17).ArraySet1((func() _dafny.Sequence {
					if _644_v7 {
						return _dafny.Companion_Sequence_.Update(_645_v8, (Companion_Default___.SafeIndex(_637_v0, _dafny.IntOfUint32((_645_v8).Cardinality()))).Uint32(), _639_v2)
					}
					return _645_v8
				})(), (_index83).Int())
				var _673_v18 _dafny.Array
				_ = _673_v18
				var _nw118 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
				_ = _nw118
				_673_v18 = _nw118
				var _674_v19 _dafny.MultiSet
				_ = _674_v19
				_674_v19 = _dafny.MultiSetOf(_673_v18)
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_672_v17), 0))
				_ = _index84
				var _rhs110 bool = (func() bool {
					if _644_v7 {
						return !(_644_v7)
					}
					return _644_v7
				})()
				_ = _rhs110
				var _rhs111 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(143))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}((func(_675_v3 *C1) func(_dafny.Int) _dafny.CodePoint {
					return func(_676_i2 _dafny.Int) _dafny.CodePoint {
						return (_675_v3).F20()
					}
				})(_640_v3))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_637_v0), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(143))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}((func(_677_v3 *C1) func(_dafny.Int) _dafny.CodePoint {
					return func(_678_i2 _dafny.Int) _dafny.CodePoint {
						return (_677_v3).F20()
					}
				})(_640_v3)))).Cardinality()))).Uint32(), _dafny.CodePoint('h'))
				_ = _rhs111
				var _rhs112 bool = (_674_v19).IsSubsetOf((_674_v19).Union(_674_v19))
				_ = _rhs112
				var _rhs113 _dafny.Int = _637_v0
				_ = _rhs113
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				var _lhs79 _dafny.Array = _672_v17
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_672_v17), 0))
				_ = _lhs80
				var _lhs81 *GlobalState = globalState
				_ = _lhs81
				_lhs78.F0 = _rhs110
				(_lhs79).ArraySet1(_rhs111, (_lhs80).Int())
				_lhs81.F3 = _rhs112
				r0 = _rhs113
				var _679_v20 _dafny.Map
				_ = _679_v20
				_679_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_637_v0, (_642_v5).Cardinality()), _644_v7)
				var _680_v21 _dafny.Set
				_ = _680_v21
				_680_v21 = _dafny.SetOf(_dafny.CodePoint('e'))
				_679_v20 = (_679_v20).Update(((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState))).Plus(_637_v0), (_680_v21).IsProperSubsetOf(_680_v21))
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_667_v13), 0))
				_ = _index85
				(_667_v13).ArraySet1(true, (_index85).Int())
				var _681_v22 _dafny.MultiSet
				_ = _681_v22
				_681_v22 = _dafny.MultiSetOf(_644_v7, _644_v7)
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_667_v13), 0))
				_ = _index86
				(_667_v13).ArraySet1((func() bool {
					if !(!(_644_v7)) || (_644_v7) {
						return (_681_v22).IsProperSubsetOf(_681_v22)
					}
					return _644_v7
				})(), (_index86).Int())
				_644_v7 = !(_dafny.Companion_Sequence_.Equal((_672_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_672_v17), 0))).Int()).(_dafny.Sequence), (_672_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_672_v17), 0))).Int()).(_dafny.Sequence)))
				var _682_v23 _dafny.Map
				_ = _682_v23
				_682_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(606), _639_v2)
				var _683_v24 _dafny.Sequence
				_ = _683_v24
				_683_v24 = _dafny.SeqOf(_682_v23)
				(globalState).F0 = (_637_v0).Cmp(((_683_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.IntOfUint32((_683_v24).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()) <= 0
			}
			var _684_v25 _dafny.Array
			_ = _684_v25
			var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(6))
			_ = _nw119
			_684_v25 = _nw119
			var _685_v26 _dafny.Set
			_ = _685_v26
			_685_v26 = _dafny.SetOf(_667_v13)
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_684_v25), 0))
			_ = _index87
			(_684_v25).ArraySet1(_685_v26, (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_684_v25), 0))
			_ = _index88
			(_684_v25).ArraySet1((_685_v26).Difference(_685_v26), (_index88).Int())
			(globalState).F0 = _644_v7
		} else {
			(globalState).F2 = _637_v0
			var _686_v27 _dafny.Map
			_ = _686_v27
			_686_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_645_v8).Cardinality()), _640_v3)
			var _687_v28 _dafny.Map
			_ = _687_v28
			_687_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_686_v27, !(_644_v7))
			_687_v28 = (_687_v28).Update(_686_v27, false)
			if _644_v7 {
				var _688_v29 *C1
				_ = _688_v29
				var _nw120 *C1 = New_C1_()
				_ = _nw120
				_nw120.Ctor__((func() _dafny.CodePoint {
					if _644_v7 {
						return (_640_v3).F20()
					}
					return _dafny.CodePoint('m')
				})())
				_688_v29 = _nw120
				(globalState).F9 = (Companion_Default___.SafeModuloInt(_637_v0, (Companion_Default___.Fm10(globalState)).Cardinality())).Minus((_637_v0).Plus(_dafny.IntOfUint32((_645_v8).Cardinality())))
				var _689_v30 _dafny.Array
				_ = _689_v30
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_14
				var _nw121 _dafny.Array
				_ = _nw121
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw121 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) bool = func(_690_i3 _dafny.Int) bool {
						return true
					}
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw121 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw121).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw121).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_689_v30 = _nw121
				var _691_v31 D2
				_ = _691_v31
				_691_v31 = Companion_D2_.Create_DC5_(_644_v7, _dafny.IntOfInt64(324), _689_v30, _642_v5)
				var _692_v32 _dafny.MultiSet
				_ = _692_v32
				_692_v32 = _dafny.MultiSetOf(_691_v31)
				var _693_v33 _dafny.Map
				_ = _693_v33
				_693_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_644_v7, _dafny.UnicodeSeqOfUtf8Bytes("saqvvnypo"))
				var _694_v34 _dafny.Sequence
				_ = _694_v34
				_694_v34 = _dafny.SeqOf(_692_v32, (_692_v32).Update(_691_v31, Companion_Default___.Abs((_dafny.Zero).Minus((_693_v33).Cardinality()))))
				_694_v34 = _694_v34
				(globalState).F0 = !(!(_644_v7)) || (_644_v7)
				(globalState).F2 = Companion_Default___.Fm0(globalState)
			} else {
				var _695_v35 *C1
				_ = _695_v35
				var _nw122 *C1 = New_C1_()
				_ = _nw122
				_nw122.Ctor__((_640_v3).F20())
				_695_v35 = _nw122
				var _696_v36 _dafny.MultiSet
				_ = _696_v36
				_696_v36 = _dafny.MultiSetOf(_645_v8)
				var _697_v37 _dafny.Sequence
				_ = _697_v37
				_697_v37 = _dafny.SeqOf(_646_v9, Companion_D5_.Create_DC13_(_637_v0, (_696_v36).Cardinality(), _644_v7))
				var _698_v38 _dafny.Sequence
				_ = _698_v38
				_698_v38 = _dafny.SeqOf(_644_v7, _dafny.Companion_Sequence_.IsPrefixOf(_697_v37, _dafny.Companion_Sequence_.Update(_697_v37, (Companion_Default___.SafeIndex(_637_v0, _dafny.IntOfUint32((_697_v37).Cardinality()))).Uint32(), _646_v9)))
				(globalState).F10 = (_698_v38).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_644_v7, false)).Cardinality()), _637_v0), _dafny.IntOfUint32((_698_v38).Cardinality()))).Uint32()).(bool)
				(globalState).F2 = _637_v0
				var _699_v39 _dafny.Map
				_ = _699_v39
				_699_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_657_v11, _dafny.CodePoint('d'))
				var _700_v40 _dafny.Array
				_ = _700_v40
				var _nwElement0_29 _dafny.CodePoint = _639_v2
				_ = _nwElement0_29
				var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(27))
				_ = _nw123
				(_nw123).ArraySet1CodePoint(_nwElement0_29, 0)
				(_nw123).ArraySet1CodePoint(Companion_Default___.Fm7(_639_v2, Companion_Default___.Fm1(_644_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(295))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}((func(_701_v3 *C1) func(_dafny.Int) _dafny.CodePoint {
					return func(_702_i4 _dafny.Int) _dafny.CodePoint {
						return (_701_v3).F20()
					}
				})(_640_v3))), globalState), globalState), 1)
				(_nw123).ArraySet1CodePoint((_695_v35).F20(), 2)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 3)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 4)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 5)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 6)
				(_nw123).ArraySet1CodePoint((_695_v35).F20(), 7)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 8)
				(_nw123).ArraySet1CodePoint(_dafny.CodePoint('b'), 9)
				(_nw123).ArraySet1CodePoint(_639_v2, 10)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 11)
				(_nw123).ArraySet1CodePoint(_639_v2, 12)
				(_nw123).ArraySet1CodePoint(_639_v2, 13)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 14)
				(_nw123).ArraySet1CodePoint(_639_v2, 15)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 16)
				(_nw123).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_699_v39).Contains(_657_v11) {
						return (_699_v39).Get(_657_v11).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('f')
				})(), 17)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 18)
				(_nw123).ArraySet1CodePoint((_695_v35).F20(), 19)
				(_nw123).ArraySet1CodePoint((_695_v35).F20(), 20)
				(_nw123).ArraySet1CodePoint(_639_v2, 21)
				(_nw123).ArraySet1CodePoint(_639_v2, 22)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 23)
				(_nw123).ArraySet1CodePoint(_dafny.CodePoint('j'), 24)
				(_nw123).ArraySet1CodePoint((_640_v3).F20(), 25)
				(_nw123).ArraySet1CodePoint(_639_v2, 26)
				_700_v40 = _nw123
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_700_v40), 0))
				_ = _index89
				(_700_v40).ArraySet1CodePoint((_695_v35).F20(), (_index89).Int())
				var _703_v41 _dafny.Array
				_ = _703_v41
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
				_ = _nw124
				_703_v41 = _nw124
				var _704_v42 _dafny.Map
				_ = _704_v42
				_704_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_703_v41, (_640_v3).F20())
				var _705_v43 _dafny.Sequence
				_ = _705_v43
				_705_v43 = _dafny.SeqOf(_703_v41)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_700_v40), 0))
				_ = _index90
				(_700_v40).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_704_v42).Contains((_705_v43).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_637_v0), _dafny.IntOfUint32((_705_v43).Cardinality()))).Uint32()).(_dafny.Array)) {
						return (_704_v42).Get((_705_v43).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_637_v0), _dafny.IntOfUint32((_705_v43).Cardinality()))).Uint32()).(_dafny.Array)).(_dafny.CodePoint)
					}
					return (_640_v3).F20()
				})(), (_index90).Int())
				var _706_v44 _dafny.Array
				_ = _706_v44
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_15
				var _nw125 _dafny.Array
				_ = _nw125
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw125 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) bool = (func(_707_v7 bool) func(_dafny.Int) bool {
						return func(_708_i5 _dafny.Int) bool {
							return _707_v7
						}
					})(_644_v7)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw125 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw125).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw125).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_706_v44 = _nw125
				var _709_v45 _dafny.MultiSet
				_ = _709_v45
				_709_v45 = _dafny.MultiSetOf(_dafny.SeqOf(_637_v0))
				var _710_v46 _dafny.Sequence
				_ = _710_v46
				_710_v46 = _dafny.SeqOf(_709_v45)
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_706_v44), 0))
				_ = _index91
				(_706_v44).ArraySet1(((_710_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_640_v3).F20())).Cardinality()), _dafny.IntOfUint32((_710_v46).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_709_v45), (_index91).Int())
				var _711_v47 D6
				_ = _711_v47
				_711_v47 = Companion_D6_.Create_DC17_((_700_v40).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_700_v40), 0))).Int()), _644_v7, !(_644_v7), (_640_v3).F20())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_706_v44), 0))
				_ = _index92
				(_706_v44).ArraySet1((_711_v47).Dtor_cf34(), (_index92).Int())
			}
			(globalState).F2 = _dafny.IntOfInt64(89)
			var _712_v48 *C1
			_ = _712_v48
			var _nw126 *C1 = New_C1_()
			_ = _nw126
			_nw126.Ctor__(_dafny.CodePoint('d'))
			_712_v48 = _nw126
		}
		var _713_v49 _dafny.Set
		_ = _713_v49
		_713_v49 = _dafny.SetOf(_644_v7)
		var _714_v50 D6
		_ = _714_v50
		_714_v50 = Companion_D6_.Create_DC16_(_637_v0, _644_v7, _dafny.IntOfUint32((_645_v8).Cardinality()), (_713_v49).Cardinality(), _644_v7)
		if !(_644_v7) || ((func() bool {
			if _644_v7 {
				return _644_v7
			}
			return (_714_v50).Dtor_cf32()
		})()) {
			var _715_v51 _dafny.Sequence
			_ = _715_v51
			_715_v51 = _dafny.SeqOf((_dafny.IntOfUint32((_dafny.SeqOf(_644_v7, _644_v7)).Cardinality())).Minus(_637_v0), _637_v0)
			_715_v51 = _715_v51
			_644_v7 = _644_v7
			var _716_v52 *C1
			_ = _716_v52
			var _nw127 *C1 = New_C1_()
			_ = _nw127
			_nw127.Ctor__(_dafny.CodePoint('k'))
			_716_v52 = _nw127
			(globalState).F3 = _644_v7
			_639_v2 = (_716_v52).F20()
		} else {
			var _717_v53 *C0
			_ = _717_v53
			var _nw128 *C0 = New_C0_()
			_ = _nw128
			_nw128.Ctor__(_638_v1, _637_v0)
			_717_v53 = _nw128
			var _718_v54 _dafny.Map
			_ = _718_v54
			_718_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_637_v0, _644_v7)
			var _719_v55 _dafny.Sequence
			_ = _719_v55
			_719_v55 = _dafny.SeqOf(_644_v7, _644_v7)
			var _720_v56 _dafny.Array
			_ = _720_v56
			var _nwElement0_30 bool = (_718_v54).Equals(_718_v54)
			_ = _nwElement0_30
			var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(25))
			_ = _nw129
			(_nw129).ArraySet1(_nwElement0_30, 0)
			(_nw129).ArraySet1(_644_v7, 1)
			(_nw129).ArraySet1((_637_v0).Cmp(_637_v0) <= 0, 2)
			(_nw129).ArraySet1(!(_644_v7), 3)
			(_nw129).ArraySet1(_644_v7, 4)
			(_nw129).ArraySet1(_644_v7, 5)
			(_nw129).ArraySet1((_719_v55).Select((Companion_Default___.SafeIndex(_637_v0, _dafny.IntOfUint32((_719_v55).Cardinality()))).Uint32()).(bool), 6)
			(_nw129).ArraySet1(_644_v7, 7)
			(_nw129).ArraySet1((func() bool {
				if !(false) {
					return _644_v7
				}
				return _644_v7
			})(), 8)
			(_nw129).ArraySet1(Companion_Default___.Fm1(_644_v7, Companion_Default___.Fm6(!(_644_v7), globalState), globalState), 9)
			(_nw129).ArraySet1(_644_v7, 10)
			(_nw129).ArraySet1(_644_v7, 11)
			(_nw129).ArraySet1(((_713_v49).Cardinality()).Cmp(_637_v0) >= 0, 12)
			(_nw129).ArraySet1(!(_642_v5).Equals(_642_v5), 13)
			(_nw129).ArraySet1(_644_v7, 14)
			(_nw129).ArraySet1(_644_v7, 15)
			(_nw129).ArraySet1(_644_v7, 16)
			(_nw129).ArraySet1(_644_v7, 17)
			(_nw129).ArraySet1(_644_v7, 18)
			(_nw129).ArraySet1(true, 19)
			(_nw129).ArraySet1(_644_v7, 20)
			(_nw129).ArraySet1((_714_v50).Dtor_cf29(), 21)
			(_nw129).ArraySet1(_644_v7, 22)
			(_nw129).ArraySet1(_644_v7, 23)
			(_nw129).ArraySet1(_644_v7, 24)
			_720_v56 = _nw129
			var _rhs114 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_645_v8, _645_v8)
			_ = _rhs114
			var _rhs115 *C0 = _717_v53
			_ = _rhs115
			var _rhs116 _dafny.Sequence = _645_v8
			_ = _rhs116
			var _rhs117 _dafny.Array = _720_v56
			_ = _rhs117
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			_645_v8 = _rhs114
			_717_v53 = _rhs115
			_645_v8 = _rhs116
			_lhs82.F11 = _rhs117
			(globalState).F9 = _637_v0
			var _721_v57 _dafny.Map
			_ = _721_v57
			_721_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_644_v7, _644_v7)
			var _pat_let_tv21 = _719_v55
			_ = _pat_let_tv21
			var _722_v58 _dafny.Sequence
			_ = _722_v58
			_722_v58 = _dafny.SeqOf(func(_pat_let27_0 D5) D5 {
				return func(_723_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let28_0 _dafny.Int) D5 {
						return func(_724_dt__update_hcf19_h0 _dafny.Int) D5 {
							return Companion_D5_.Create_DC13_(_724_dt__update_hcf19_h0, (_723_dt__update__tmp_h0).Dtor_cf20(), (_723_dt__update__tmp_h0).Dtor_cf21())
						}(_pat_let28_0)
					}(_dafny.IntOfUint32((_pat_let_tv21).Cardinality()))
				}(_pat_let27_0)
			}(_646_v9), _646_v9)
			_721_v57 = (_721_v57).Update(Companion_Default___.Fm1(_644_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(269))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_725_v3 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_726_i6 _dafny.Int) _dafny.CodePoint {
					return (_725_v3).F20()
				}
			})(_640_v3))), globalState), _dafny.Companion_Sequence_.IsProperPrefixOf(_722_v58, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer29 func(_dafny.Int) D5) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_727_v9 D5) func(_dafny.Int) D5 {
				return func(_728_i7 _dafny.Int) D5 {
					return _727_v9
				}
			})(_646_v9)))))
			var _729_v59 D1
			_ = _729_v59
			_729_v59 = Companion_D1_.Create_DC2_(_644_v7)
			var _730_v60 D1
			_ = _730_v60
			_730_v60 = Companion_D1_.Create_DC3_(_729_v59)
			_730_v60 = _730_v60
			if _644_v7 {
				(globalState).F10 = (func() bool {
					if _644_v7 {
						return (func() bool {
							if _644_v7 {
								return _644_v7
							}
							return Companion_Default___.Fm1(_644_v7, _dafny.UnicodeSeqOfUtf8Bytes("xtbk"), globalState)
						})()
					}
					return _644_v7
				})()
				var _731_v61 _dafny.Array
				_ = _731_v61
				var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw130
				_731_v61 = _nw130
				_731_v61 = _731_v61
				(globalState).F3 = _644_v7
				var _732_v62 _dafny.Map
				_ = _732_v62
				_732_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_721_v57).Cardinality()), _637_v0)
				_732_v62 = _732_v62
				(globalState).F2 = (_637_v0).Minus(_637_v0)
			} else {
				(globalState).F10 = _644_v7
				var _733_v63 _dafny.Array
				_ = _733_v63
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw131
				_733_v63 = _nw131
				var _734_v64 D4
				_ = _734_v64
				_734_v64 = Companion_D4_.Create_DC10_(_733_v63)
				var _735_v65 D4
				_ = _735_v65
				_735_v65 = Companion_D4_.Create_DC11_(_734_v64)
				var _736_v66 _dafny.Sequence
				_ = _736_v66
				_736_v66 = _dafny.SeqOf(_637_v0, _637_v0)
				var _737_v67 _dafny.Map
				_ = _737_v67
				_737_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_735_v65, _736_v66)
				var _738_v68 _dafny.Map
				_ = _738_v68
				_738_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_737_v67).Contains(_735_v65) {
						return (_737_v67).Get(_735_v65).(_dafny.Sequence)
					}
					return _736_v66
				})(), Companion_Default___.SafeModuloInt(_637_v0, _637_v0))
				_738_v68 = (_738_v68).Update((func() _dafny.Sequence {
					if _644_v7 {
						return _736_v66
					}
					return _736_v66
				})(), _dafny.IntOfInt64(329))
				_730_v60 = (func() D1 {
					if (_644_v7) || (_644_v7) {
						return _730_v60
					}
					return _730_v60
				})()
				_637_v0 = _637_v0
				var _739_v69 _dafny.Map
				_ = _739_v69
				_739_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v56, _644_v7)
				var _740_v71 _dafny.Set
				_ = _740_v71
				_740_v71 = _dafny.SetOf(_637_v0, _637_v0)
				(globalState).F0 = (func() bool {
					if (_739_v69).Contains(_720_v56) {
						return (_739_v69).Get(_720_v56).(bool)
					}
					return (_740_v71).IsProperSubsetOf(func() _dafny.Set {
						var _coll18 = _dafny.NewBuilder()
						_ = _coll18
						for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-694), _dafny.IntOfInt64(-854))); ; {
							_compr_18, _ok21 := _iter21()
							if !_ok21 {
								break
							}
							var _741_v70 _dafny.Int
							_741_v70 = interface{}(_compr_18).(_dafny.Int)
							if ((_dafny.IntOfInt64(-694)).Cmp(_741_v70) <= 0) && ((_741_v70).Cmp(_dafny.IntOfInt64(-854)) < 0) {
								_coll18.Add((_741_v70).Minus(_637_v0))
							}
						}
						return _coll18.ToSet()
					}())
				})()
			}
		}
		r0 = (_dafny.Zero).Minus((_637_v0).Minus((_637_v0).Times(_637_v0)))
		return r0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f17 _dafny.Int
	_f18 _dafny.Int
	_f19 _dafny.Sequence
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.Zero
	_this._f19 = _dafny.EmptySeq
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

func (_this *C3) F17() _dafny.Int {
	return _this._f17
}
func (_this *C3) F18() _dafny.Int {
	return _this._f18
}
func (_this *C3) Ctor__(f19 _dafny.Sequence, f17 _dafny.Int, f18 _dafny.Int) {
	{
		(_this)._f19 = f19
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C3) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_this).F17()
	}
}
func (_this *C3) Fm3(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F18()
	}
}
func (_this *C3) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_this).F18(), ((_this).F18()).Plus((_this).F18()))
	}
}
func (_this *C3) M1(globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _742_v0 _dafny.Sequence
		_ = _742_v0
		_742_v0 = _dafny.SeqOf((_this).F18())
		var _hi3 _dafny.Int = ((_this).F18()).Times(_dafny.IntOfUint32(((_this).F19()).Cardinality()))
		_ = _hi3
		for _743_i0 := (_dafny.Zero).Minus((_742_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-233), _dafny.IntOfUint32((_742_v0).Cardinality()))).Uint32()).(_dafny.Int)); _743_i0.Cmp(_hi3) < 0; _743_i0 = _743_i0.Plus(_dafny.One) {
			var _source11 _dafny.Int = (_this).F17()
			_ = _source11
			var _744___mcc_h0 _dafny.Int = _source11
			_ = _744___mcc_h0
			var _745_cf0 _dafny.Int = _744___mcc_h0
			_ = _745_cf0
			(globalState).F2 = _743_i0
			var _746_v1 bool
			_ = _746_v1
			_746_v1 = true
			var _747_v2 _dafny.Map
			_ = _747_v2
			_747_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
				if _746_v1 {
					return _746_v1
				}
				return _746_v1
			})()), _745_cf0)
			(globalState).F2 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_747_v2).Contains((_743_i0).Cmp(_dafny.IntOfUint32(((_this).F19()).Cardinality())) <= 0) {
					return (_747_v2).Get((_743_i0).Cmp(_dafny.IntOfUint32(((_this).F19()).Cardinality())) <= 0).(_dafny.Int)
				}
				return _dafny.IntOfInt64(720)
			})())
			(globalState).F9 = _dafny.IntOfUint32(((_this).F19()).Cardinality())
			var _748_v3 _dafny.Array
			_ = _748_v3
			var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw132
			_748_v3 = _nw132
			var _749_v4 _dafny.Map
			_ = _749_v4
			_749_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_746_v1, _746_v1)
			var _750_v5 _dafny.Array
			_ = _750_v5
			var _nwElement0_31 bool = _746_v1
			_ = _nwElement0_31
			var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(6))
			_ = _nw133
			(_nw133).ArraySet1(_nwElement0_31, 0)
			(_nw133).ArraySet1((func() bool {
				if (_749_v4).Contains(!(_746_v1)) {
					return (_749_v4).Get(!(_746_v1)).(bool)
				}
				return _746_v1
			})(), 1)
			(_nw133).ArraySet1(_746_v1, 2)
			(_nw133).ArraySet1(Companion_Default___.Fm1(true, _dafny.UnicodeSeqOfUtf8Bytes("ti"), globalState), 3)
			(_nw133).ArraySet1(_746_v1, 4)
			(_nw133).ArraySet1(true, 5)
			_750_v5 = _nw133
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_748_v3), 0))
			_ = _index93
			(_748_v3).ArraySet1(_750_v5, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_750_v5), 0))
			_ = _index94
			(_750_v5).ArraySet1(true, (_index94).Int())
			var _751_v6 D1
			_ = _751_v6
			_751_v6 = Companion_D1_.Create_DC1_(_750_v5)
			var _752_v7 _dafny.Sequence
			_ = _752_v7
			_752_v7 = _dafny.SeqOf(_746_v1)
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_748_v3), 0))
			_ = _index95
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_750_v5), 0))
			_ = _index96
			var _rhs118 _dafny.Array = (_751_v6).Dtor_cf1()
			_ = _rhs118
			var _rhs119 _dafny.Int = _743_i0
			_ = _rhs119
			var _rhs120 bool = (func() bool {
				if (_752_v7).Select((Companion_Default___.SafeIndex(_743_i0, _dafny.IntOfUint32((_752_v7).Cardinality()))).Uint32()).(bool) {
					return _746_v1
				}
				return _746_v1
			})()
			_ = _rhs120
			var _rhs121 bool = _746_v1
			_ = _rhs121
			var _lhs83 _dafny.Array = _748_v3
			_ = _lhs83
			var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_748_v3), 0))
			_ = _lhs84
			var _lhs85 *GlobalState = globalState
			_ = _lhs85
			var _lhs86 *GlobalState = globalState
			_ = _lhs86
			var _lhs87 _dafny.Array = _750_v5
			_ = _lhs87
			var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_750_v5), 0))
			_ = _lhs88
			(_lhs83).ArraySet1(_rhs118, (_lhs84).Int())
			_lhs85.F2 = _rhs119
			_lhs86.F0 = _rhs120
			(_lhs87).ArraySet1(_rhs121, (_lhs88).Int())
			var _753_v8 _dafny.Array
			_ = _753_v8
			var _nw134 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw134
			_753_v8 = _nw134
			var _754_v9 D1
			_ = _754_v9
			_754_v9 = Companion_D1_.Create_DC1_(_753_v8)
			var _source12 D1 = _754_v9
			_ = _source12
			if _source12.Is_DC2() {
				var _755___mcc_h1 bool = _source12.Get_().(D1_DC2).Cf2
				_ = _755___mcc_h1
				var _756_cf2 bool = _755___mcc_h1
				_ = _756_cf2
				var _757_v10 _dafny.CodePoint
				_ = _757_v10
				_757_v10 = _dafny.CodePoint('x')
				var _758_v11 *C1
				_ = _758_v11
				var _nw135 *C1 = New_C1_()
				_ = _nw135
				_nw135.Ctor__(_757_v10)
				_758_v11 = _nw135
				(globalState).F9 = (_this).F18()
				var _759_v12 *C0
				_ = _759_v12
				var _nw136 *C0 = New_C0_()
				_ = _nw136
				_nw136.Ctor__((_this).F17(), (_dafny.Zero).Minus((_this).F18()))
				_759_v12 = _nw136
				var _760_v13 _dafny.Map
				_ = _760_v13
				_760_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_i0, _dafny.IntOfInt64(549))
				var _761_v14 _dafny.Map
				_ = _761_v14
				_761_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_756_cf2, (_this).F18())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_753_v8), 0))
				_ = _index97
				(_753_v8).ArraySet1(!(_dafny.SetOf(_743_i0, (_this).F18())).Equals(Companion_Default___.Fm8((func() _dafny.Int {
					if (_761_v14).Contains(_756_cf2) {
						return (_761_v14).Get(_756_cf2).(_dafny.Int)
					}
					return _743_i0
				})(), _743_i0, globalState)), (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_753_v8), 0))
				_ = _index98
				var _rhs122 _dafny.CodePoint = _757_v10
				_ = _rhs122
				var _rhs123 _dafny.Int = _743_i0
				_ = _rhs123
				var _rhs124 _dafny.Map = (_760_v13).Merge(_760_v13)
				_ = _rhs124
				var _rhs125 bool = (_743_i0).Cmp((_this).F18()) < 0
				_ = _rhs125
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 _dafny.Array = _753_v8
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_753_v8), 0))
				_ = _lhs91
				_757_v10 = _rhs122
				_lhs89.F9 = _rhs123
				_760_v13 = _rhs124
				(_lhs90).ArraySet1(_rhs125, (_lhs91).Int())
			} else if _source12.Is_DC1() {
				var _762___mcc_h2 _dafny.Array = _source12.Get_().(D1_DC1).Cf1
				_ = _762___mcc_h2
				var _763_cf1 _dafny.Array = _762___mcc_h2
				_ = _763_cf1
				var _764_v15 bool
				_ = _764_v15
				_764_v15 = false
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_753_v8), 0))
				_ = _index99
				(_753_v8).ArraySet1(_764_v15, (_index99).Int())
				var _765_v16 T0
				_ = _765_v16
				var _nw137 *C0 = New_C0_()
				_ = _nw137
				_nw137.Ctor__((_this).F18(), _743_i0)
				_765_v16 = _nw137
				var _766_v17 _dafny.Set
				_ = _766_v17
				_766_v17 = _dafny.SetOf(_765_v16, _765_v16)
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_753_v8), 0))
				_ = _index100
				(_753_v8).ArraySet1((_766_v17).IsProperSubsetOf((_766_v17).Difference(_766_v17)), (_index100).Int())
				var _767_v18 _dafny.Array
				_ = _767_v18
				var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw138
				_767_v18 = _nw138
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_767_v18), 0))
				_ = _index101
				(_767_v18).ArraySet1(_dafny.IntOfInt64(874), (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_753_v8), 0))
				_ = _index102
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_767_v18), 0))
				_ = _index103
				var _rhs126 bool = (_753_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_753_v8), 0))).Int()).(bool)
				_ = _rhs126
				var _rhs127 _dafny.Int = (_this).F18()
				_ = _rhs127
				var _lhs92 _dafny.Array = _753_v8
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_753_v8), 0))
				_ = _lhs93
				var _lhs94 _dafny.Array = _767_v18
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_767_v18), 0))
				_ = _lhs95
				(_lhs92).ArraySet1(_rhs126, (_lhs93).Int())
				(_lhs94).ArraySet1(_rhs127, (_lhs95).Int())
				_765_v16 = _765_v16
				var _768_v19 _dafny.Set
				_ = _768_v19
				_768_v19 = _dafny.SetOf((_this).F18(), _743_i0)
				var _769_v20 _dafny.Map
				_ = _769_v20
				_769_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_768_v19, _763_cf1)
				_769_v20 = (_769_v20).Update(_768_v19, _763_cf1)
			} else {
				var _770___mcc_h3 D1 = _source12.Get_().(D1_DC3).Cf3
				_ = _770___mcc_h3
				var _771_cf3 D1 = _770___mcc_h3
				_ = _771_cf3
				var _772_v21 bool
				_ = _772_v21
				_772_v21 = true
				var _773_v22 _dafny.CodePoint
				_ = _773_v22
				_773_v22 = _dafny.CodePoint('t')
				var _774_v23 *C1
				_ = _774_v23
				var _nw139 *C1 = New_C1_()
				_ = _nw139
				_nw139.Ctor__(_773_v22)
				_774_v23 = _nw139
				var _775_v24 _dafny.Sequence
				_ = _775_v24
				_775_v24 = _dafny.SeqOf((func() *C1 {
					if _772_v21 {
						return _774_v23
					}
					return _774_v23
				})(), _774_v23, _774_v23, (func() *C1 {
					if _772_v21 {
						return _774_v23
					}
					return _774_v23
				})(), _774_v23)
				_775_v24 = _dafny.Companion_Sequence_.Concatenate(_775_v24, _775_v24)
				var _776_v25 *C0
				_ = _776_v25
				var _nw140 *C0 = New_C0_()
				_ = _nw140
				_nw140.Ctor__(_dafny.IntOfUint32(((_this).F19()).Cardinality()), (_this).F18())
				_776_v25 = _nw140
				var _777_v26 D4
				_ = _777_v26
				_777_v26 = Companion_D4_.Create_DC9_(_776_v25)
				var _778_v27 D4
				_ = _778_v27
				_778_v27 = Companion_D4_.Create_DC11_(_777_v26)
				var _779_v28 D4
				_ = _779_v28
				_779_v28 = Companion_D4_.Create_DC11_(_778_v27)
				var _780_v29 _dafny.Map
				_ = _780_v29
				_780_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_772_v21, _779_v28)
				var _781_v30 _dafny.Set
				_ = _781_v30
				_781_v30 = _dafny.SetOf(_780_v29)
				var _782_v31 D7
				_ = _782_v31
				_782_v31 = Companion_D7_.Create_DC18_(_781_v30)
				_781_v30 = (_782_v31).Dtor_cf37()
				var _783_v32 _dafny.Int
				_ = _783_v32
				var _784_v33 _dafny.MultiSet
				_ = _784_v33
				var _out26 _dafny.Int
				_ = _out26
				var _out27 _dafny.MultiSet
				_ = _out27
				_out26, _out27 = (_774_v23).M3(((_this).F18()).Cmp((_this).F18()) == 0, globalState)
				_783_v32 = _out26
				_784_v33 = _out27
				var _785_v34 *C2
				_ = _785_v34
				var _nw141 *C2 = New_C2_()
				_ = _nw141
				_nw141.Ctor__()
				_785_v34 = _nw141
				var _786_v35 _dafny.Array
				_ = _786_v35
				var _nwElement0_32 *C2 = _785_v34
				_ = _nwElement0_32
				var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(13))
				_ = _nw142
				(_nw142).ArraySet1(_nwElement0_32, 0)
				(_nw142).ArraySet1(_785_v34, 1)
				(_nw142).ArraySet1(_785_v34, 2)
				(_nw142).ArraySet1(_785_v34, 3)
				(_nw142).ArraySet1((Companion_D8_.Create_DC22_(_785_v34)).Dtor_cf45(), 4)
				(_nw142).ArraySet1(_785_v34, 5)
				(_nw142).ArraySet1((func() *C2 {
					if _772_v21 {
						return _785_v34
					}
					return _785_v34
				})(), 6)
				(_nw142).ArraySet1(_785_v34, 7)
				(_nw142).ArraySet1(_785_v34, 8)
				(_nw142).ArraySet1(_785_v34, 9)
				(_nw142).ArraySet1(_785_v34, 10)
				(_nw142).ArraySet1(_785_v34, 11)
				(_nw142).ArraySet1(_785_v34, 12)
				_786_v35 = _nw142
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_786_v35), 0))
				_ = _index104
				(_786_v35).ArraySet1(_785_v34, (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_786_v35), 0))
				_ = _index105
				(_786_v35).ArraySet1(_785_v34, (_index105).Int())
			}
			var _787_v36 _dafny.Array
			_ = _787_v36
			var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
			_ = _nw143
			_787_v36 = _nw143
			var _788_v37 bool
			_ = _788_v37
			_788_v37 = false
			var _789_v38 _dafny.Map
			_ = _789_v38
			_789_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F19()).Cardinality()), _788_v37)
			var _790_v40 _dafny.Set
			_ = _790_v40
			_790_v40 = _dafny.SetOf((_dafny.Zero).Minus((_this).F18()), _743_i0)
			var _791_v41 _dafny.Sequence
			_ = _791_v41
			_791_v41 = _dafny.SeqOf(_789_v38, _789_v38, func() _dafny.Map {
				var _coll19 = _dafny.NewMapBuilder()
				_ = _coll19
				for _iter22 := _dafny.Iterate((_790_v40).Elements()); ; {
					_compr_19, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _792_v39 _dafny.Int
					_792_v39 = interface{}(_compr_19).(_dafny.Int)
					if (_790_v40).Contains(_792_v39) {
						_coll19.Add(Companion_Default___.SafeDivisionInt(_792_v39, _743_i0), _788_v37)
					}
				}
				return _coll19.ToMap()
			}())
			var _793_v42 _dafny.Map
			_ = _793_v42
			_793_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v37, (_this).Fm4(_dafny.IntOfUint32((_791_v41).Cardinality()), (_this).F19(), _742_v0, globalState))
			var _794_v43 _dafny.Array
			_ = _794_v43
			var _nw144 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
			_ = _nw144
			_794_v43 = _nw144
			var _795_v44 _dafny.MultiSet
			_ = _795_v44
			_795_v44 = _dafny.MultiSetOf(_794_v43, _794_v43, _794_v43, _794_v43, _794_v43)
			var _796_v45 _dafny.MultiSet
			_ = _796_v45
			_796_v45 = _dafny.MultiSetOf((_793_v42).Cardinality(), (func() _dafny.Int {
				if (_795_v44).Contains(_794_v43) {
					return (_795_v44).Multiplicity(_794_v43)
				}
				return (_this).F18()
			})())
			var _797_v46 D2
			_ = _797_v46
			_797_v46 = Companion_D2_.Create_DC5_(_788_v37, (_dafny.Zero).Minus((_this).F18()), _753_v8, _796_v45)
			var _798_v47 _dafny.Set
			_ = _798_v47
			_798_v47 = _dafny.SetOf(Companion_D2_.Create_DC5_(_788_v37, _743_i0, _753_v8, _796_v45), _797_v46)
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_787_v36), 0))
			_ = _index106
			(_787_v36).ArraySet1(_798_v47, (_index106).Int())
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_787_v36), 0))
			_ = _index107
			(_787_v36).ArraySet1(_798_v47, (_index107).Int())
			var _799_v48 T0
			_ = _799_v48
			var _nw145 *C0 = New_C0_()
			_ = _nw145
			_nw145.Ctor__((_this).F17(), (_this).F18())
			_799_v48 = _nw145
			var _800_v49 _dafny.Map
			_ = _800_v49
			_800_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_799_v48, _788_v37)
			(globalState).F3 = !((_800_v49).Merge(_800_v49)).Equals((_800_v49).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_799_v48, _788_v37)))
		}
		var _hi4 _dafny.Int = _dafny.IntOfInt64(-109)
		_ = _hi4
		for _801_i1 := (_this).F18(); _801_i1.Cmp(_hi4) < 0; _801_i1 = _801_i1.Plus(_dafny.One) {
			var _802_v50 bool
			_ = _802_v50
			_802_v50 = false
			(globalState).F2 = ((func() _dafny.Int {
				if _802_v50 {
					return (_this).Fm3(_802_v50, (_this).F18(), globalState)
				}
				return (_this).F18()
			})()).Plus((_this).F18())
			var _803_v51 _dafny.CodePoint
			_ = _803_v51
			_803_v51 = _dafny.CodePoint('x')
			_803_v51 = _803_v51
			(globalState).F9 = (_801_i1).Plus((_dafny.Zero).Minus((_742_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F19()).Cardinality()), _dafny.IntOfUint32((_742_v0).Cardinality()))).Uint32()).(_dafny.Int)))
			_802_v50 = _802_v50
		}
		var _804_v52 _dafny.Array
		_ = _804_v52
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_16
		var _nw146 _dafny.Array
		_ = _nw146
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw146 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Int = func(_805_i2 _dafny.Int) _dafny.Int {
				return (_805_i2).Plus(_dafny.IntOfInt64(-107))
			}
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw146 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw146).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw146).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_804_v52 = _nw146
		var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))
		_ = _index108
		(_804_v52).ArraySet1((_this).F18(), (_index108).Int())
		var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))
		_ = _index109
		(_804_v52).ArraySet1(((_this).F18()).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState))), (_index109).Int())
		var _806_v53 _dafny.Array
		_ = _806_v53
		var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
		_ = _nw147
		_806_v53 = _nw147
		var _807_v54 _dafny.Array
		_ = _807_v54
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_17
		var _nw148 _dafny.Array
		_ = _nw148
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw148 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) bool = func(_808_i3 _dafny.Int) bool {
				return !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("alyvhe"), (_this).F19())
			}
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw148 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw148).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw148).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_807_v54 = _nw148
		var _809_v55 bool
		_ = _809_v55
		_809_v55 = true
		var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))
		_ = _index110
		(_807_v54).ArraySet1(_809_v55, (_index110).Int())
		var _810_v56 _dafny.MultiSet
		_ = _810_v56
		_810_v56 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F19()).Cardinality()), _809_v55)).Cardinality())
		var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))
		_ = _index111
		var _rhs128 _dafny.Array = _806_v53
		_ = _rhs128
		var _rhs129 bool = (_810_v56).IsSubsetOf(_810_v56)
		_ = _rhs129
		var _lhs96 _dafny.Array = _807_v54
		_ = _lhs96
		var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))
		_ = _lhs97
		_806_v53 = _rhs128
		(_lhs96).ArraySet1(_rhs129, (_lhs97).Int())
		var _811_v57 _dafny.Sequence
		_ = _811_v57
		_811_v57 = _dafny.UnicodeSeqOfUtf8Bytes("baxh")
		var _812_v58 _dafny.MultiSet
		_ = _812_v58
		_812_v58 = _dafny.MultiSetOf(_804_v52)
		var _rhs130 _dafny.Sequence = (_this).F19()
		_ = _rhs130
		var _rhs131 _dafny.Int = (_this).F18()
		_ = _rhs131
		var _rhs132 _dafny.MultiSet = _812_v58
		_ = _rhs132
		var _rhs133 bool = (func() bool {
			if (_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool) {
				return (_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool)
			}
			return (_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool)
		})()
		_ = _rhs133
		var _lhs98 *GlobalState = globalState
		_ = _lhs98
		var _lhs99 *GlobalState = globalState
		_ = _lhs99
		_811_v57 = _rhs130
		_lhs98.F9 = _rhs131
		_812_v58 = _rhs132
		_lhs99.F0 = _rhs133
		var _813_i4 _dafny.Int
		_ = _813_i4
		_813_i4 = _dafny.Zero
		{
			for _809_v55 {
				{
					if (_813_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_813_i4 = (_813_i4).Plus(_dafny.One)
					if _809_v55 {
						var _814_v59 _dafny.Map
						_ = _814_v59
						_814_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool))
						var _815_v60 _dafny.Int
						_ = _815_v60
						var _816_v61 _dafny.Int
						_ = _816_v61
						var _out28 _dafny.Int
						_ = _out28
						var _out29 _dafny.Int
						_ = _out29
						_out28, _out29 = Companion_Default___.M0(_809_v55, _814_v59, globalState)
						_815_v60 = _out28
						_816_v61 = _out29
						var _817_v63 _dafny.Set
						_ = _817_v63
						_817_v63 = _dafny.SetOf(_814_v59)
						var _818_v64 D6
						_ = _818_v64
						_818_v64 = Companion_D6_.Create_DC15_((func() _dafny.Map {
							var _coll20 = _dafny.NewMapBuilder()
							_ = _coll20
							for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-322), _dafny.IntOfInt64(616))); ; {
								_compr_20, _ok23 := _iter23()
								if !_ok23 {
									break
								}
								var _819_v62 _dafny.Int
								_819_v62 = interface{}(_compr_20).(_dafny.Int)
								if ((_dafny.IntOfInt64(-322)).Cmp(_819_v62) <= 0) && ((_819_v62).Cmp(_dafny.IntOfInt64(616)) < 0) {
									_coll20.Add((_819_v62).Plus((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)), (_dafny.Zero).Minus(_816_v61))
								}
							}
							return _coll20.ToMap()
						}()).Cardinality(), _816_v61, _dafny.SeqOf((_this).F18()), (_this).F18(), _dafny.SetOf(_815_v60, (_817_v63).Cardinality(), (_dafny.Zero).Minus((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)), _815_v60, _816_v61))
						var _820_v65 _dafny.CodePoint
						_ = _820_v65
						_820_v65 = _dafny.CodePoint('c')
						var _821_v66 _dafny.Map
						_ = _821_v66
						_821_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool), (_this).F18())
						var _rhs134 D6 = _818_v64
						_ = _rhs134
						var _rhs135 bool = !((_821_v66).Merge(_821_v66)).Equals((_821_v66).Merge(_821_v66))
						_ = _rhs135
						var _rhs136 _dafny.CodePoint = _820_v65
						_ = _rhs136
						var _rhs137 _dafny.Int = (_816_v61).Times((func() _dafny.Int {
							if (_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool) {
								return _815_v60
							}
							return _815_v60
						})())
						_ = _rhs137
						var _rhs138 bool = Companion_Default___.Fm1(Companion_Default___.Fm1((_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool), _811_v57, globalState), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("g"), (Companion_Default___.SafeIndex((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality()))).Uint32(), _820_v65), globalState)
						_ = _rhs138
						var _lhs100 *GlobalState = globalState
						_ = _lhs100
						var _lhs101 *GlobalState = globalState
						_ = _lhs101
						_818_v64 = _rhs134
						_lhs100.F10 = _rhs135
						_820_v65 = _rhs136
						_815_v60 = _rhs137
						_lhs101.F3 = _rhs138
						var _822_v67 _dafny.Array
						_ = _822_v67
						var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
						_ = _nw149
						_822_v67 = _nw149
						var _823_v68 *C1
						_ = _823_v68
						var _nw150 *C1 = New_C1_()
						_ = _nw150
						_nw150.Ctor__(_820_v65)
						_823_v68 = _nw150
						var _824_v69 _dafny.Map
						_ = _824_v69
						_824_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_v68, (_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int))
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_822_v67), 0))
						_ = _index112
						(_822_v67).ArraySet1(_824_v69, (_index112).Int())
						var _825_v70 D1
						_ = _825_v70
						_825_v70 = Companion_D1_.Create_DC1_(_807_v54)
						var _826_v71 _dafny.Sequence
						_ = _826_v71
						_826_v71 = _dafny.SeqOf(_825_v70)
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))
						_ = _index113
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_822_v67), 0))
						_ = _index114
						var _rhs139 _dafny.Int = (_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)
						_ = _rhs139
						var _rhs140 _dafny.Int = (Companion_Default___.Fm0(globalState)).Times(((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_826_v71).Cardinality())))
						_ = _rhs140
						var _rhs141 bool = (_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool)
						_ = _rhs141
						var _rhs142 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_v68, _dafny.IntOfInt64(670))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_v68, _dafny.IntOfInt64(143))).Merge(_824_v69))
						_ = _rhs142
						var _rhs143 _dafny.Int = (_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)
						_ = _rhs143
						var _lhs102 *GlobalState = globalState
						_ = _lhs102
						var _lhs103 _dafny.Array = _807_v54
						_ = _lhs103
						var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))
						_ = _lhs104
						var _lhs105 _dafny.Array = _822_v67
						_ = _lhs105
						var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_822_v67), 0))
						_ = _lhs106
						var _lhs107 *GlobalState = globalState
						_ = _lhs107
						_lhs102.F9 = _rhs139
						_816_v61 = _rhs140
						(_lhs103).ArraySet1(_rhs141, (_lhs104).Int())
						(_lhs105).ArraySet1(_rhs142, (_lhs106).Int())
						_lhs107.F2 = _rhs143
						(globalState).F3 = (_816_v61).Cmp((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)) <= 0
						(globalState).F9 = _dafny.IntOfUint32(((_this).F19()).Cardinality())
					} else {
						(globalState).F2 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-132), (_this).F18())
						(globalState).F2 = (_this).F18()
						_811_v57 = (_this).F19()
						_742_v0 = _742_v0
						var _827_v72 _dafny.Map
						_ = _827_v72
						_827_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_this).F19()).Cardinality()), (_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)))
						_827_v72 = (_827_v72).Update((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(55))
					}
					(globalState).F9 = ((_this).F18()).Minus((func() _dafny.Int {
						if _809_v55 {
							return (_this).Fm3(_809_v55, (_this).F18(), globalState)
						}
						return _dafny.IntOfInt64(850)
					})())
					var _828_v73 _dafny.Sequence
					_ = _828_v73
					_828_v73 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg30 _dafny.Int) interface{} {
							return coer30(arg30)
						}
					}(func(_829_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					})), _dafny.UnicodeSeqOfUtf8Bytes("bj"), _811_v57, _811_v57, _811_v57)
					var _830_v74 _dafny.Sequence
					_ = _830_v74
					_830_v74 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg31 _dafny.Int) interface{} {
							return coer31(arg31)
						}
					}(func(_831_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					})), (_this).F19(), (_this).F19(), (_828_v73).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_828_v73).Cardinality()))).Uint32()).(_dafny.Sequence))
					var _832_v75 _dafny.Sequence
					_ = _832_v75
					_832_v75 = _dafny.SeqOf((_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool))
					var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))
					_ = _index115
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))
					_ = _index116
					var _rhs144 _dafny.Int = (_this).F18()
					_ = _rhs144
					var _rhs145 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_828_v73).Select((Companion_Default___.SafeIndex((_this).Fm3((_832_v75).Select((Companion_Default___.SafeIndex((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_832_v75).Cardinality()))).Uint32()).(bool), _dafny.IntOfUint32((_832_v75).Cardinality()), globalState), _dafny.IntOfUint32((_828_v73).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int))
					_ = _rhs145
					var _lhs108 _dafny.Array = _804_v52
					_ = _lhs108
					var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))
					_ = _lhs109
					var _lhs110 _dafny.Array = _804_v52
					_ = _lhs110
					var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))
					_ = _lhs111
					(_lhs108).ArraySet1(_rhs144, (_lhs109).Int())
					(_lhs110).ArraySet1(_rhs145, (_lhs111).Int())
					(globalState).F9 = ((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus((_804_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_804_v52), 0))).Int()).(_dafny.Int)))
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		r0 = _804_v52
		r1 = (_807_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_807_v54), 0))).Int()).(bool)
		return r0, r1
	}
}
func (_this *C3) F19() _dafny.Sequence {
	{
		return _this._f19
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
