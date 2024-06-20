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
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) bool {
	return (((_dafny.MultiSetOf(!(true), true)).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality()))).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-257))))) < 0
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) _dafny.Int {
	return (((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(98), _dafny.IntOfInt64(163))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(98)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(163)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_0_v0, _dafny.IntOfInt64(131)), false)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-722), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(393), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-391))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality()), false), func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(122), _dafny.IntOfInt64(265))).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(122), _dafny.IntOfInt64(265))).Contains(_2_v1) {
				_coll1.Add((_2_v1).Minus(_dafny.IntOfInt64(798)), true)
			}
		}
		return _coll1.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-223))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality()), true))).Cardinality()))).Minus(_dafny.IntOfInt64(117))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), true)).Cardinality())).Cardinality())).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_2).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), true)).Cardinality())).Cardinality()), _4_v0) {
				_coll2.Add((_4_v0).Times(_dafny.IntOfInt64(802)), _dafny.IntOfInt64(-596))
			}
		}
		return _coll2.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(567), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(true, true), _dafny.SeqOf(false, false), _dafny.SeqOf(true))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(837))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_5_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(47))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("okfva")).Cardinality()))))).Cardinality(), _dafny.IntOfInt64(-567), _dafny.IntOfInt64(792), _dafny.IntOfInt64(-57))).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-769))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_6_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(985)
	}))).Cardinality()), _dafny.IntOfInt64(-713))) {
		return _dafny.CodePoint('l')
	} else {
		return _dafny.CodePoint('k')
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(!(false))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC9_(_dafny.CodePoint('k')), false)).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _7_v0 D2
			_7_v0 = interface{}(_compr_3).(D2)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC9_(_dafny.CodePoint('k')), false)).Contains(_7_v0) {
				_coll3.Add(_7_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ev"), _dafny.UnicodeSeqOfUtf8Bytes("sbujxi"))).Cardinality()))
			}
		}
		return _coll3.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	if true {
		return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Union(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_D2_.Create_DC10_(false, !(false))).Dtor_cf20())))
	} else {
		return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Union(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!((_dafny.IntOfInt64(-677)).Cmp(_dafny.IntOfInt64(772)) > 0), (_dafny.SetOf(_dafny.MultiSetOf(false, !(false)))).IsProperSubsetOf((Companion_D6_.Create_DC21_(_dafny.SetOf(_dafny.MultiSetOf(false, !(true))))).Dtor_cf42()))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _8_globalState *GlobalState
	_ = _8_globalState
	var _nw0 *GlobalState = New_GlobalState_()
	_ = _nw0
	_nw0.Ctor__(_dafny.IntOfInt64(470), _dafny.IntOfInt64(-748), true, _dafny.IntOfInt64(220), _dafny.IntOfInt64(551), false)
	_8_globalState = _nw0
	var _9_v0 _dafny.CodePoint
	_ = _9_v0
	_9_v0 = _dafny.CodePoint('e')
	var _10_v1 *C1
	_ = _10_v1
	var _nw1 *C1 = New_C1_()
	_ = _nw1
	_nw1.Ctor__(_9_v0, false)
	_10_v1 = _nw1
	(_8_globalState).F1 = (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-275))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}((func(_11_v1 *C1) func(_dafny.Int) _dafny.Sequence {
		return func(_12_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf((_11_v1).F7())
		}
	})(_10_v1))))).Cardinality()
	var _13_v2 D1
	_ = _13_v2
	_13_v2 = Companion_D1_.Create_DC6_(_dafny.CodePoint('o'), Companion_Default___.Fm6(_8_globalState))
	var _14_v3 _dafny.Sequence
	_ = _14_v3
	_14_v3 = _dafny.SeqOf(Companion_D1_.Create_DC6_(_9_v0, (_10_v1).F6()), (func() D1 {
		if false {
			return _13_v2
		}
		return _13_v2
	})())
	_14_v3 = _dafny.SeqOf(_13_v2, _13_v2, _13_v2, _13_v2, _13_v2)
	var _15_v4 _dafny.Int
	_ = _15_v4
	_15_v4 = _dafny.IntOfInt64(22)
	var _16_v5 _dafny.Map
	_ = _16_v5
	_16_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v4, _15_v4)
	var _17_v6 _dafny.Map
	_ = _17_v6
	_17_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_10_v1).F7()), true)
	var _18_v7 _dafny.Array
	_ = _18_v7
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = (func(_19_v1 *C1) func(_dafny.Int) bool {
			return func(_20_i1 _dafny.Int) bool {
				return (_19_v1).F7()
			}
		})(_10_v1)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw2 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw2).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw2).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_18_v7 = _nw2
	var _21_v8 D0
	_ = _21_v8
	_21_v8 = Companion_D0_.Create_DC2_(_16_v5, _17_v6, _18_v7)
	var _pat_let_tv0 = _18_v7
	_ = _pat_let_tv0
	var _source0 D0 = func(_pat_let0_0 D0) D0 {
		return func(_22_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 _dafny.Array) D0 {
				return func(_23_dt__update_hcf8_h0 _dafny.Array) D0 {
					return Companion_D0_.Create_DC2_((_22_dt__update__tmp_h0).Dtor_cf6(), (_22_dt__update__tmp_h0).Dtor_cf7(), _23_dt__update_hcf8_h0)
				}(_pat_let1_0)
			}(_pat_let_tv0)
		}(_pat_let0_0)
	}(_21_v8)
	_ = _source0
	if _source0.Is_DC1() {
		var _24___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _24___mcc_h0
		var _25___mcc_h1 _dafny.Set = _source0.Get_().(D0_DC1).Cf2
		_ = _25___mcc_h1
		var _26___mcc_h2 _dafny.Int = _source0.Get_().(D0_DC1).Cf3
		_ = _26___mcc_h2
		var _27___mcc_h3 _dafny.Array = _source0.Get_().(D0_DC1).Cf4
		_ = _27___mcc_h3
		var _28___mcc_h4 _dafny.CodePoint = _source0.Get_().(D0_DC1).Cf5
		_ = _28___mcc_h4
		var _29_cf5 _dafny.CodePoint = _28___mcc_h4
		_ = _29_cf5
		var _30_cf4 _dafny.Array = _27___mcc_h3
		_ = _30_cf4
		var _31_cf3 _dafny.Int = _26___mcc_h2
		_ = _31_cf3
		var _32_cf2 _dafny.Set = _25___mcc_h1
		_ = _32_cf2
		var _33_cf1 _dafny.Int = _24___mcc_h0
		_ = _33_cf1
		var _nwElement0_0 bool = (_10_v1).F7()
		_ = _nwElement0_0
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(2))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_0, 0)
		(_nw3).ArraySet1((Companion_Default___.Fm7(_dafny.IntOfInt64(758), _15_v4, false, _8_globalState)).Dtor_cf10(), 1)
		_18_v7 = _nw3
		var _34_v9 _dafny.Map
		_ = _34_v9
		_34_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() *C1 {
			if (_10_v1).F7() {
				return _10_v1
			}
			return _10_v1
		})())
		_34_v9 = (_34_v9).Update((_10_v1).F7(), _10_v1)
		_16_v5 = (_16_v5).Update(_15_v4, _15_v4)
		var _35_v10 _dafny.MultiSet
		_ = _35_v10
		_35_v10 = _dafny.MultiSetOf(_33_cf1)
		(_8_globalState).F5 = ((_35_v10).IsSubsetOf(_35_v10)) == (Companion_Default___.Fm1(_8_globalState))
	} else if _source0.Is_DC2() {
		var _36___mcc_h5 _dafny.Map = _source0.Get_().(D0_DC2).Cf6
		_ = _36___mcc_h5
		var _37___mcc_h6 _dafny.Map = _source0.Get_().(D0_DC2).Cf7
		_ = _37___mcc_h6
		var _38___mcc_h7 _dafny.Array = _source0.Get_().(D0_DC2).Cf8
		_ = _38___mcc_h7
		var _39_cf8 _dafny.Array = _38___mcc_h7
		_ = _39_cf8
		var _40_cf7 _dafny.Map = _37___mcc_h6
		_ = _40_cf7
		var _41_cf6 _dafny.Map = _36___mcc_h5
		_ = _41_cf6
		var _42_v11 _dafny.Sequence
		_ = _42_v11
		_42_v11 = _dafny.SeqOf(_15_v4, _15_v4, _15_v4)
		var _43_v12 _dafny.Map
		_ = _43_v12
		_43_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_42_v11, _15_v4)
		var _44_v13 _dafny.Int
		_ = _44_v13
		var _45_v14 _dafny.Int
		_ = _45_v14
		var _46_v15 _dafny.Array
		_ = _46_v15
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		var _out2 _dafny.Array
		_ = _out2
		_out0, _out1, _out2 = (_10_v1).M0((func() _dafny.Int {
			if (_43_v12).Contains(_42_v11) {
				return (_43_v12).Get(_42_v11).(_dafny.Int)
			}
			return _dafny.IntOfInt64(105)
		})(), _15_v4, _8_globalState)
		_44_v13 = _out0
		_45_v14 = _out1
		_46_v15 = _out2
		var _47_v16 _dafny.Sequence
		_ = _47_v16
		_47_v16 = _dafny.UnicodeSeqOfUtf8Bytes("olipx")
		_47_v16 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(301))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_48_v1 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_49_i2 _dafny.Int) _dafny.CodePoint {
				return (_48_v1).F6()
			}
		})(_10_v1))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(434))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}((func(_50_v1 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_51_i3 _dafny.Int) _dafny.CodePoint {
				return (_50_v1).F6()
			}
		})(_10_v1))))
		(_8_globalState).F5 = (_10_v1).F7()
		var _52_v17 D0
		_ = _52_v17
		_52_v17 = Companion_D0_.Create_DC3_(_45_v14)
		var _53_v18 D1
		_ = _53_v18
		_53_v18 = Companion_D1_.Create_DC5_((_52_v17).Dtor_cf9(), _41_cf6, _dafny.IntOfUint32((_47_v16).Cardinality()))
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_46_v15), 0))
		_ = _index0
		(_46_v15).ArraySet1((_53_v18).Dtor_cf11(), (_index0).Int())
		var _54_v19 _dafny.Set
		_ = _54_v19
		_54_v19 = _dafny.SetOf((_10_v1).F7())
		var _55_v20 _dafny.Array
		_ = _55_v20
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_1
		var _nw4 _dafny.Array
		_ = _nw4
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw4 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Sequence = (func(_56_v1 *C1) func(_dafny.Int) _dafny.Sequence {
				return func(_57_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_56_v1).F7())
				}
			})(_10_v1)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw4 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw4).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw4).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_55_v20 = _nw4
		var _58_v21 D0
		_ = _58_v21
		_58_v21 = Companion_D0_.Create_DC1_(_15_v4, _54_v19, _45_v14, _55_v20, _9_v0)
		var _59_v22 _dafny.Map
		_ = _59_v22
		_59_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_58_v21).Dtor_cf5(), !((_10_v1).F7()))
		var _60_v23 _dafny.Sequence
		_ = _60_v23
		_60_v23 = _dafny.SeqOf((_59_v22).Update((_10_v1).F6(), (_10_v1).F7()))
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_46_v15), 0))
		_ = _index1
		(_46_v15).ArraySet1(_dafny.IntOfUint32((_60_v23).Cardinality()), (_index1).Int())
	} else if _source0.Is_DC3() {
		var _61___mcc_h8 _dafny.Int = _source0.Get_().(D0_DC3).Cf9
		_ = _61___mcc_h8
		var _62_cf9 _dafny.Int = _61___mcc_h8
		_ = _62_cf9
		var _63_v24 _dafny.Array
		_ = _63_v24
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw5
		_63_v24 = _nw5
		var _64_v25 _dafny.Sequence
		_ = _64_v25
		_64_v25 = _dafny.SeqOf(_63_v24)
		_64_v25 = _dafny.Companion_Sequence_.Concatenate(_64_v25, _dafny.Companion_Sequence_.Update(_64_v25, (Companion_Default___.SafeIndex(_15_v4, _dafny.IntOfUint32((_64_v25).Cardinality()))).Uint32(), _63_v24))
		(_8_globalState).F5 = !((_10_v1).F7())
		var _65_v26 _dafny.Set
		_ = _65_v26
		_65_v26 = _dafny.SetOf((Companion_D2_.Create_DC8_(_10_v1)).Dtor_cf17())
		if (_65_v26).IsDisjointFrom((_65_v26).Union(_65_v26)) {
			var _66_v27 _dafny.Array
			_ = _66_v27
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_2
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Map = (func(_67_v1 *C1, _68_v4 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_69_i5 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC9_((_67_v1).F6()), _68_v4)
					}
				})(_10_v1, _15_v4)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw6 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw6).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw6).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_66_v27 = _nw6
			var _70_v28 _dafny.Map
			_ = _70_v28
			_70_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC9_((_10_v1).F6()), _15_v4)
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_66_v27), 0))
			_ = _index2
			(_66_v27).ArraySet1((func() _dafny.Map {
				if !((_10_v1).F7()) {
					return _70_v28
				}
				return _70_v28
			})(), (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_66_v27), 0))
			_ = _index3
			(_66_v27).ArraySet1((_70_v28).Merge(Companion_Default___.Fm8(_15_v4, (_10_v1).F7(), (_10_v1).F7(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_10_v1).F7()), _62_cf9), _8_globalState)), (_index3).Int())
			var _rhs0 *C1 = _10_v1
			_ = _rhs0
			var _rhs1 bool = (_10_v1).F7()
			_ = _rhs1
			var _lhs0 *GlobalState = _8_globalState
			_ = _lhs0
			_10_v1 = _rhs0
			_lhs0.F5 = _rhs1
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_18_v7), 0))
			_ = _index4
			(_18_v7).ArraySet1((_10_v1).F7(), (_index4).Int())
			var _71_v29 _dafny.Set
			_ = _71_v29
			_71_v29 = _dafny.SetOf(!((_10_v1).F7()), ((_10_v1).F7()) || ((_10_v1).F7()), (_10_v1).F7())
			var _72_v30 _dafny.Map
			_ = _72_v30
			_72_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v5, !((_10_v1).F7()))
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_18_v7), 0))
			_ = _index5
			var _rhs2 bool = ((_72_v30).Cardinality()).Cmp(_15_v4) >= 0
			_ = _rhs2
			var _rhs3 _dafny.Set = _71_v29
			_ = _rhs3
			var _lhs1 _dafny.Array = _18_v7
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_18_v7), 0))
			_ = _lhs2
			(_lhs1).ArraySet1(_rhs2, (_lhs2).Int())
			_71_v29 = _rhs3
			var _73_v31 *C0
			_ = _73_v31
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__(_16_v5, (_18_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_18_v7), 0))).Int()).(bool))
			_73_v31 = _nw7
			_73_v31 = _73_v31
			var _74_v32 _dafny.Array
			_ = _74_v32
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw8
			_74_v32 = _nw8
			var _75_v33 _dafny.Sequence
			_ = _75_v33
			_75_v33 = _dafny.UnicodeSeqOfUtf8Bytes("bhphawf")
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_74_v32), 0))
			_ = _index6
			(_74_v32).ArraySet1(_75_v33, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_74_v32), 0))
			_ = _index7
			(_74_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_75_v33, _75_v33), (_index7).Int())
		} else {
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_18_v7), 0))
			_ = _index8
			(_18_v7).ArraySet1((_10_v1).F7(), (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_18_v7), 0))
			_ = _index9
			(_18_v7).ArraySet1((func() bool {
				if (_10_v1).F7() {
					return (_62_cf9).Cmp(_dafny.IntOfInt64(-190)) == 0
				}
				return Companion_Default___.Fm1(_8_globalState)
			})(), (_index9).Int())
			_10_v1 = _10_v1
			var _76_v34 *C1
			_ = _76_v34
			var _nw9 *C1 = New_C1_()
			_ = _nw9
			_nw9.Ctor__((_10_v1).F6(), (_10_v1).F7())
			_76_v34 = _nw9
			var _77_v35 _dafny.Sequence
			_ = _77_v35
			_77_v35 = _dafny.SeqOf(_15_v4, _62_cf9)
			(_8_globalState).F5 = !(!((_10_v1).F7()) || (_dafny.Companion_Sequence_.IsPrefixOf(_77_v35, _77_v35)))
			var _78_v36 D2
			_ = _78_v36
			_78_v36 = Companion_D2_.Create_DC8_(_10_v1)
			var _79_v37 _dafny.Array
			_ = _79_v37
			var _nwElement0_1 *C1 = (_78_v36).Dtor_cf17()
			_ = _nwElement0_1
			var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(14))
			_ = _nw10
			(_nw10).ArraySet1(_nwElement0_1, 0)
			(_nw10).ArraySet1(_10_v1, 1)
			(_nw10).ArraySet1(_10_v1, 2)
			(_nw10).ArraySet1(_76_v34, 3)
			(_nw10).ArraySet1(_76_v34, 4)
			(_nw10).ArraySet1(_76_v34, 5)
			(_nw10).ArraySet1(_10_v1, 6)
			(_nw10).ArraySet1(_76_v34, 7)
			(_nw10).ArraySet1(_76_v34, 8)
			(_nw10).ArraySet1(_10_v1, 9)
			(_nw10).ArraySet1((func() *C1 {
				if (_10_v1).F7() {
					return _10_v1
				}
				return _10_v1
			})(), 10)
			(_nw10).ArraySet1(_10_v1, 11)
			(_nw10).ArraySet1(_76_v34, 12)
			(_nw10).ArraySet1(_76_v34, 13)
			_79_v37 = _nw10
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_79_v37), 0))
			_ = _index10
			(_79_v37).ArraySet1(_10_v1, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_79_v37), 0))
			_ = _index11
			(_79_v37).ArraySet1(_10_v1, (_index11).Int())
		}
		var _80_v38 _dafny.Sequence
		_ = _80_v38
		_80_v38 = _dafny.UnicodeSeqOfUtf8Bytes("qsq")
		_15_v4 = (_15_v4).Minus(_dafny.IntOfUint32((_80_v38).Cardinality()))
	} else {
		var _81___mcc_h9 _dafny.Array = _source0.Get_().(D0_DC0).Cf0
		_ = _81___mcc_h9
		var _82_cf0 _dafny.Array = _81___mcc_h9
		_ = _82_cf0
		var _83_v39 _dafny.Array
		_ = _83_v39
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_3
		var _nw11 _dafny.Array
		_ = _nw11
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw11 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Int = (func(_84_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_85_i6 _dafny.Int) _dafny.Int {
					return (_85_i6).Plus(_84_v4)
				}
			})(_15_v4)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw11 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw11).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw11).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_83_v39 = _nw11
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_83_v39), 0))
		_ = _index12
		(_83_v39).ArraySet1(_15_v4, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_83_v39), 0))
		_ = _index13
		var _rhs4 _dafny.Int = _15_v4
		_ = _rhs4
		var _rhs5 _dafny.Int = _dafny.IntOfInt64(280)
		_ = _rhs5
		var _lhs3 _dafny.Array = _83_v39
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_83_v39), 0))
		_ = _lhs4
		var _lhs5 *GlobalState = _8_globalState
		_ = _lhs5
		(_lhs3).ArraySet1(_rhs4, (_lhs4).Int())
		_lhs5.F1 = _rhs5
		var _86_v40 D0
		_ = _86_v40
		_86_v40 = Companion_D0_.Create_DC0_(_18_v7)
		var _87_v41 _dafny.Map
		_ = _87_v41
		_87_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v40, (_83_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_83_v39), 0))).Int()).(_dafny.Int))
		var _88_v42 _dafny.Sequence
		_ = _88_v42
		_88_v42 = _dafny.UnicodeSeqOfUtf8Bytes("htil")
		var _89_v43 D1
		_ = _89_v43
		_89_v43 = Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_88_v42).Cardinality()), _16_v5, (_83_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_83_v39), 0))).Int()).(_dafny.Int))
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_83_v39), 0))
		_ = _index14
		(_83_v39).ArraySet1((func() _dafny.Int {
			if (_10_v1).F7() {
				return Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if (_87_v41).Contains(_86_v40) {
						return (_87_v41).Get(_86_v40).(_dafny.Int)
					}
					return (_83_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_83_v39), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfInt64(623))
			}
			return (_dafny.Zero).Minus((_89_v43).Dtor_cf13())
		})(), (_index14).Int())
		_18_v7 = _18_v7
		_17_v6 = (_17_v6).Update((_10_v1).F7(), (_10_v1).F7())
	}
	var _hi0 _dafny.Int = _15_v4
	_ = _hi0
	for _90_i7 := _15_v4; _90_i7.Cmp(_hi0) < 0; _90_i7 = _90_i7.Plus(_dafny.One) {
		(_8_globalState).F1 = (_90_i7).Minus(Companion_Default___.Fm2(false, _8_globalState))
		var _91_v44 _dafny.Array
		_ = _91_v44
		var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
		_ = _nw12
		_91_v44 = _nw12
		var _92_v45 _dafny.Sequence
		_ = _92_v45
		_92_v45 = _dafny.UnicodeSeqOfUtf8Bytes("a")
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_91_v44), 0))
		_ = _index15
		(_91_v44).ArraySet1((_15_v4).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_92_v45).Cardinality()))), (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_91_v44), 0))
		_ = _index16
		(_91_v44).ArraySet1(_15_v4, (_index16).Int())
		var _93_v46 _dafny.Map
		_ = _93_v46
		_93_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ekadg"), _15_v4)
		_93_v46 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("tpsq"), (_91_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_91_v44), 0))).Int()).(_dafny.Int))).Merge(_93_v46)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_18_v7), 0))
		_ = _index17
		(_18_v7).ArraySet1(!(false), (_index17).Int())
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_18_v7), 0))
		_ = _index18
		var _rhs6 bool = (_10_v1).F7()
		_ = _rhs6
		var _rhs7 bool = (_10_v1).F7()
		_ = _rhs7
		var _rhs8 _dafny.CodePoint = (_10_v1).F6()
		_ = _rhs8
		var _lhs6 _dafny.Array = _18_v7
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_18_v7), 0))
		_ = _lhs7
		var _lhs8 *GlobalState = _8_globalState
		_ = _lhs8
		(_lhs6).ArraySet1(_rhs6, (_lhs7).Int())
		_lhs8.F5 = _rhs7
		_9_v0 = _rhs8
	}
	for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_18_v7), 0))); ; {
		_guard_loop_0, _ok4 := _iter4()
		if !_ok4 {
			break
		}
		var _94_i8 _dafny.Int
		_94_i8 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_94_i8).Sign() != -1) && ((_94_i8).Cmp(_dafny.ArrayLen((_18_v7), 0)) < 0)) {
			(_18_v7).ArraySet1((_10_v1).F7(), (_94_i8).Int())
		}
	}
	var _95_v47 _dafny.Sequence
	_ = _95_v47
	_95_v47 = _dafny.UnicodeSeqOfUtf8Bytes("fyyyndfmr")
	_95_v47 = (func() _dafny.Sequence {
		if (_10_v1).F7() {
			return _95_v47
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("vintro")
	})()
	var _hi1 _dafny.Int = _dafny.IntOfInt64(320)
	_ = _hi1
	for _96_i9 := _15_v4; _96_i9.Cmp(_hi1) < 0; _96_i9 = _96_i9.Plus(_dafny.One) {
		var _97_v48 _dafny.Int
		_ = _97_v48
		var _98_v49 _dafny.Int
		_ = _98_v49
		var _99_v50 _dafny.Array
		_ = _99_v50
		var _out3 _dafny.Int
		_ = _out3
		var _out4 _dafny.Int
		_ = _out4
		var _out5 _dafny.Array
		_ = _out5
		_out3, _out4, _out5 = (_10_v1).M0(_15_v4, _15_v4, _8_globalState)
		_97_v48 = _out3
		_98_v49 = _out4
		_99_v50 = _out5
		var _100_v51 D2
		_ = _100_v51
		_100_v51 = Companion_D2_.Create_DC8_(_10_v1)
		var _101_v52 D2
		_ = _101_v52
		_101_v52 = Companion_D2_.Create_DC12_(_100_v51)
		var _source1 D2 = _101_v52
		_ = _source1
		if _source1.Is_DC9() {
			var _102___mcc_h10 _dafny.CodePoint = _source1.Get_().(D2_DC9).Cf18
			_ = _102___mcc_h10
			var _103_cf18 _dafny.CodePoint = _102___mcc_h10
			_ = _103_cf18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_99_v50), 0))
			_ = _index19
			(_99_v50).ArraySet1((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_10_v1).F7(), _96_i9)).Cardinality()), (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_99_v50), 0))
			_ = _index20
			(_99_v50).ArraySet1(_96_i9, (_index20).Int())
			var _104_v53 _dafny.Array
			_ = _104_v53
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw13
			_104_v53 = _nw13
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_104_v53), 0))
			_ = _index21
			(_104_v53).ArraySet1(_18_v7, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_104_v53), 0))
			_ = _index22
			(_104_v53).ArraySet1(_18_v7, (_index22).Int())
			var _105_v54 D1
			_ = _105_v54
			_105_v54 = Companion_D1_.Create_DC7_((_10_v1).F6())
			var _106_v55 _dafny.Map
			_ = _106_v55
			_106_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v54, (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pb")).Cardinality())), (_dafny.Zero).Minus(_98_v49))))
			_106_v55 = _106_v55
			_16_v5 = (_16_v5).Update((_15_v4).Plus(_96_i9), _dafny.IntOfInt64(321))
		} else if _source1.Is_DC10() {
			var _107___mcc_h11 bool = _source1.Get_().(D2_DC10).Cf19
			_ = _107___mcc_h11
			var _108___mcc_h12 bool = _source1.Get_().(D2_DC10).Cf20
			_ = _108___mcc_h12
			var _109_cf20 bool = _108___mcc_h12
			_ = _109_cf20
			var _110_cf19 bool = _107___mcc_h11
			_ = _110_cf19
			var _111_v56 _dafny.Map
			_ = _111_v56
			_111_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _97_v48)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_99_v50), 0))
			_ = _index23
			(_99_v50).ArraySet1(((_111_v56).Update((_10_v1).F7(), _97_v48)).Cardinality(), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_99_v50), 0))
			_ = _index24
			(_99_v50).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_95_v47).Cardinality()), _98_v49), (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_18_v7), 0))
			_ = _index25
			(_18_v7).ArraySet1(_110_cf19, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_18_v7), 0))
			_ = _index26
			(_18_v7).ArraySet1(!(true) || (!_dafny.Companion_Sequence_.Contains(_95_v47, _9_v0)), (_index26).Int())
			var _112_v57 _dafny.Map
			_ = _112_v57
			_112_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm2((_10_v1).F7(), _8_globalState)).Cmp((_99_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_99_v50), 0))).Int()).(_dafny.Int)) != 0, _dafny.SetOf((_18_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_18_v7), 0))).Int()).(bool)))
			var _113_v58 D1
			_ = _113_v58
			_113_v58 = Companion_D1_.Create_DC4_(_110_cf19)
			var _114_v59 _dafny.Map
			_ = _114_v59
			_114_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v58, _dafny.UnicodeSeqOfUtf8Bytes("g"))
			var _115_v60 _dafny.Sequence
			_ = _115_v60
			_115_v60 = _dafny.SeqOf(_98_v49)
			var _116_v61 _dafny.Set
			_ = _116_v61
			_116_v61 = _dafny.SetOf(_109_cf20, true, Companion_Default___.Fm1(_8_globalState), _110_cf19, _110_cf19)
			var _rhs9 bool = !_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
				if (_114_v59).Contains(Companion_D1_.Create_DC4_(_110_cf19)) {
					return (_114_v59).Get(Companion_D1_.Create_DC4_(_110_cf19)).(_dafny.Sequence)
				}
				return _95_v47
			})(), _95_v47)
			_ = _rhs9
			var _rhs10 _dafny.Int = (_dafny.Zero).Minus(_98_v49)
			_ = _rhs10
			var _rhs11 _dafny.Map = (_112_v57).Update(_dafny.Companion_Sequence_.Contains(_115_v60, (_99_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_99_v50), 0))).Int()).(_dafny.Int)), (_dafny.SetOf(_110_cf19, _109_cf20, _110_cf19)).Intersection(_116_v61))
			_ = _rhs11
			var _rhs12 *C1 = _10_v1
			_ = _rhs12
			var _lhs9 *GlobalState = _8_globalState
			_ = _lhs9
			_110_cf19 = _rhs9
			_lhs9.F1 = _rhs10
			_112_v57 = _rhs11
			_10_v1 = _rhs12
			var _117_v62 _dafny.Map
			_ = _117_v62
			_117_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v49, _dafny.MultiSetOf(_10_v1, _10_v1))
			var _118_v63 D2
			_ = _118_v63
			_118_v63 = Companion_D2_.Create_DC8_(_10_v1)
			var _119_v64 _dafny.MultiSet
			_ = _119_v64
			_119_v64 = _dafny.MultiSetOf(_10_v1, _10_v1, _10_v1, (_118_v63).Dtor_cf17(), _10_v1)
			var _120_v65 _dafny.Int
			_ = _120_v65
			var _121_v66 _dafny.Int
			_ = _121_v66
			var _122_v67 _dafny.Array
			_ = _122_v67
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			var _out8 _dafny.Array
			_ = _out8
			_out6, _out7, _out8 = (_10_v1).M0(_dafny.IntOfUint32((_115_v60).Cardinality()), ((func() _dafny.MultiSet {
				if (_117_v62).Contains(_98_v49) {
					return (_117_v62).Get(_98_v49).(_dafny.MultiSet)
				}
				return _119_v64
			})()).Cardinality(), _8_globalState)
			_120_v65 = _out6
			_121_v66 = _out7
			_122_v67 = _out8
		} else if _source1.Is_DC11() {
			var _123___mcc_h13 _dafny.Int = _source1.Get_().(D2_DC11).Cf21
			_ = _123___mcc_h13
			var _124___mcc_h14 _dafny.Int = _source1.Get_().(D2_DC11).Cf22
			_ = _124___mcc_h14
			var _125_cf22 _dafny.Int = _124___mcc_h14
			_ = _125_cf22
			var _126_cf21 _dafny.Int = _123___mcc_h13
			_ = _126_cf21
			var _rhs13 _dafny.Int = _15_v4
			_ = _rhs13
			var _rhs14 _dafny.Array = _99_v50
			_ = _rhs14
			_126_cf21 = _rhs13
			_99_v50 = _rhs14
			(_8_globalState).F5 = true
			(_8_globalState).F5 = (_10_v1).F7()
			var _127_v68 _dafny.Map
			_ = _127_v68
			_127_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v49, true)
			_127_v68 = (_127_v68).Update(_dafny.IntOfInt64(-12), false)
		} else if _source1.Is_DC8() {
			var _128___mcc_h15 *C1 = _source1.Get_().(D2_DC8).Cf17
			_ = _128___mcc_h15
			var _129_cf17 *C1 = _128___mcc_h15
			_ = _129_cf17
			_18_v7 = _18_v7
			(_8_globalState).F1 = _97_v48
			var _130_v69 _dafny.Map
			_ = _130_v69
			_130_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_95_v47).Cardinality()), !((_10_v1).F7()))
			var _131_v70 _dafny.Int
			_ = _131_v70
			var _132_v71 _dafny.Int
			_ = _132_v71
			var _133_v72 _dafny.Array
			_ = _133_v72
			var _out9 _dafny.Int
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Array
			_ = _out11
			_out9, _out10, _out11 = (_10_v1).M0(_98_v49, Companion_Default___.SafeModuloInt((_130_v69).Cardinality(), _dafny.IntOfInt64(440)), _8_globalState)
			_131_v70 = _out9
			_132_v71 = _out10
			_133_v72 = _out11
			(_8_globalState).F5 = (_10_v1).F7()
		} else {
			var _134___mcc_h16 D2 = _source1.Get_().(D2_DC12).Cf23
			_ = _134___mcc_h16
			var _135_cf23 D2 = _134___mcc_h16
			_ = _135_cf23
			(_8_globalState).F5 = !(Companion_Default___.Fm1(_8_globalState))
			var _136_v73 *C0
			_ = _136_v73
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_16_v5, (_10_v1).F7())
			_136_v73 = _nw14
			var _137_v74 _dafny.MultiSet
			_ = _137_v74
			_137_v74 = _dafny.MultiSetOf(_136_v73, _136_v73, _136_v73, _136_v73)
			var _138_v75 _dafny.Sequence
			_ = _138_v75
			_138_v75 = _dafny.SeqOf((_137_v74).Cardinality())
			var _139_v76 _dafny.Sequence
			_ = _139_v76
			_139_v76 = _dafny.SeqOf(_138_v75)
			var _140_v77 _dafny.Set
			_ = _140_v77
			_140_v77 = _dafny.SetOf(_97_v48, _dafny.IntOfUint32((_139_v76).Cardinality()))
			var _rhs15 bool = (_10_v1).F7()
			_ = _rhs15
			var _rhs16 _dafny.Set = _140_v77
			_ = _rhs16
			var _lhs10 *GlobalState = _8_globalState
			_ = _lhs10
			_lhs10.F5 = _rhs15
			_140_v77 = _rhs16
			_9_v0 = _9_v0
			var _141_v78 _dafny.MultiSet
			_ = _141_v78
			_141_v78 = _dafny.MultiSetOf(_18_v7, _18_v7)
			var _142_v79 _dafny.Map
			_ = _142_v79
			_142_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v49, (_dafny.MultiSetOf(_18_v7)).Update(_18_v7, Companion_Default___.Abs(_98_v49)))
			var _143_v80 _dafny.Array
			_ = _143_v80
			var _nwElement0_2 _dafny.MultiSet = _141_v78
			_ = _nwElement0_2
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(14))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_2, 0)
			(_nw15).ArraySet1(_141_v78, 1)
			(_nw15).ArraySet1(_141_v78, 2)
			(_nw15).ArraySet1(_141_v78, 3)
			(_nw15).ArraySet1((_141_v78).Update(_18_v7, Companion_Default___.Abs(_97_v48)), 4)
			(_nw15).ArraySet1(_141_v78, 5)
			(_nw15).ArraySet1((_141_v78).Difference(_141_v78), 6)
			(_nw15).ArraySet1((_141_v78).Union(_141_v78), 7)
			(_nw15).ArraySet1(_141_v78, 8)
			(_nw15).ArraySet1(_141_v78, 9)
			(_nw15).ArraySet1(_141_v78, 10)
			(_nw15).ArraySet1((func() _dafny.MultiSet {
				if (_142_v79).Contains(_97_v48) {
					return (_142_v79).Get(_97_v48).(_dafny.MultiSet)
				}
				return _141_v78
			})(), 11)
			(_nw15).ArraySet1((_141_v78).Difference(_141_v78), 12)
			(_nw15).ArraySet1(_dafny.MultiSetOf(_18_v7), 13)
			_143_v80 = _nw15
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_143_v80), 0))
			_ = _index27
			(_143_v80).ArraySet1(((_dafny.MultiSetOf(_18_v7)).Update(_18_v7, Companion_Default___.Abs(_96_i9))).Update(_18_v7, Companion_Default___.Abs(_dafny.IntOfUint32((_95_v47).Cardinality()))), (_index27).Int())
			var _144_v81 _dafny.Sequence
			_ = _144_v81
			_144_v81 = _dafny.SeqOf(_18_v7)
			var _145_v82 _dafny.Sequence
			_ = _145_v82
			_145_v82 = _dafny.SeqOf(_141_v78, _141_v78, _141_v78, _141_v78, _dafny.MultiSetOf(_18_v7, _18_v7, (_144_v81).Select((Companion_Default___.SafeIndex(_15_v4, _dafny.IntOfUint32((_144_v81).Cardinality()))).Uint32()).(_dafny.Array), _18_v7, _18_v7))
			var _146_v83 _dafny.Map
			_ = _146_v83
			_146_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_10_v1).F7(), _15_v4)
			var _147_v84 _dafny.Map
			_ = _147_v84
			_147_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v77, _146_v83)
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_143_v80), 0))
			_ = _index28
			(_143_v80).ArraySet1(((_145_v82).Select((Companion_Default___.SafeIndex(((func() _dafny.Map {
				if (_147_v84).Contains(_140_v77) {
					return (_147_v84).Get(_140_v77).(_dafny.Map)
				}
				return _146_v83
			})()).Cardinality(), _dafny.IntOfUint32((_145_v82).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference(_141_v78), (_index28).Int())
		}
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_18_v7), 0))
		_ = _index29
		(_18_v7).ArraySet1((_10_v1).F7(), (_index29).Int())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_18_v7), 0))
		_ = _index30
		(_18_v7).ArraySet1(true, (_index30).Int())
		var _148_v85 *C1
		_ = _148_v85
		var _nw16 *C1 = New_C1_()
		_ = _nw16
		_nw16.Ctor__(_dafny.CodePoint('b'), false)
		_148_v85 = _nw16
	}
	var _149_i10 _dafny.Int
	_ = _149_i10
	_149_i10 = _dafny.Zero
	{
		for (_10_v1).F7() {
			{
				if (_149_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_149_i10 = (_149_i10).Plus(_dafny.One)
				if (_10_v1).F7() {
					(_8_globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_15_v4, _15_v4))
					var _150_v86 D2
					_ = _150_v86
					_150_v86 = Companion_D2_.Create_DC9_((_10_v1).F6())
					_150_v86 = _150_v86
					var _151_v87 _dafny.Int
					_ = _151_v87
					var _152_v88 _dafny.Int
					_ = _152_v88
					var _153_v89 _dafny.Array
					_ = _153_v89
					var _out12 _dafny.Int
					_ = _out12
					var _out13 _dafny.Int
					_ = _out13
					var _out14 _dafny.Array
					_ = _out14
					_out12, _out13, _out14 = (_10_v1).M0(_15_v4, _15_v4, _8_globalState)
					_151_v87 = _out12
					_152_v88 = _out13
					_153_v89 = _out14
					var _154_v90 _dafny.Map
					_ = _154_v90
					_154_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v7, _15_v4)
					_154_v90 = (_154_v90).Update(_18_v7, _15_v4)
					_95_v47 = (_10_v1).Fm0(_dafny.UnicodeSeqOfUtf8Bytes("arwyvlusb"), (_10_v1).F6(), _8_globalState)
				} else {
					_9_v0 = (_95_v47).Select((Companion_Default___.SafeIndex(_15_v4, _dafny.IntOfUint32((_95_v47).Cardinality()))).Uint32()).(_dafny.CodePoint)
					var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_18_v7), 0))
					_ = _index31
					(_18_v7).ArraySet1((_10_v1).F7(), (_index31).Int())
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_18_v7), 0))
					_ = _index32
					var _rhs17 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_95_v47, _95_v47), _dafny.UnicodeSeqOfUtf8Bytes("hgpygk"))
					_ = _rhs17
					var _rhs18 bool = (_10_v1).F7()
					_ = _rhs18
					var _rhs19 bool = (_10_v1).F7()
					_ = _rhs19
					var _lhs11 *GlobalState = _8_globalState
					_ = _lhs11
					var _lhs12 _dafny.Array = _18_v7
					_ = _lhs12
					var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_18_v7), 0))
					_ = _lhs13
					_95_v47 = _rhs17
					_lhs11.F5 = _rhs18
					(_lhs12).ArraySet1(_rhs19, (_lhs13).Int())
					(_8_globalState).F5 = !((_15_v4).Cmp(_15_v4) < 0)
					var _155_v91 _dafny.Array
					_ = _155_v91
					var _nwElement0_3 _dafny.CodePoint = _9_v0
					_ = _nwElement0_3
					var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(2))
					_ = _nw17
					(_nw17).ArraySet1CodePoint(_nwElement0_3, 0)
					(_nw17).ArraySet1CodePoint(_9_v0, 1)
					_155_v91 = _nw17
					var _156_v92 _dafny.Sequence
					_ = _156_v92
					_156_v92 = _dafny.SeqOf(_155_v91)
					var _157_v93 D3
					_ = _157_v93
					_157_v93 = Companion_D3_.Create_DC13_(_155_v91)
					var _158_v94 _dafny.Array
					_ = _158_v94
					var _nwElement0_4 _dafny.Array = _155_v91
					_ = _nwElement0_4
					var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(26))
					_ = _nw18
					(_nw18).ArraySet1(_nwElement0_4, 0)
					(_nw18).ArraySet1(_155_v91, 1)
					(_nw18).ArraySet1(_155_v91, 2)
					(_nw18).ArraySet1(_155_v91, 3)
					(_nw18).ArraySet1(_155_v91, 4)
					(_nw18).ArraySet1(_155_v91, 5)
					(_nw18).ArraySet1(_155_v91, 6)
					(_nw18).ArraySet1(_155_v91, 7)
					(_nw18).ArraySet1(_155_v91, 8)
					(_nw18).ArraySet1((func() _dafny.Array {
						if (_10_v1).F7() {
							return _155_v91
						}
						return _155_v91
					})(), 9)
					(_nw18).ArraySet1(_155_v91, 10)
					(_nw18).ArraySet1((_156_v92).Select((Companion_Default___.SafeIndex(_15_v4, _dafny.IntOfUint32((_156_v92).Cardinality()))).Uint32()).(_dafny.Array), 11)
					(_nw18).ArraySet1(_155_v91, 12)
					(_nw18).ArraySet1(_155_v91, 13)
					(_nw18).ArraySet1(_155_v91, 14)
					(_nw18).ArraySet1(_155_v91, 15)
					(_nw18).ArraySet1(_155_v91, 16)
					(_nw18).ArraySet1(_155_v91, 17)
					(_nw18).ArraySet1(_155_v91, 18)
					(_nw18).ArraySet1((_157_v93).Dtor_cf24(), 19)
					(_nw18).ArraySet1(_155_v91, 20)
					(_nw18).ArraySet1(_155_v91, 21)
					(_nw18).ArraySet1(_155_v91, 22)
					(_nw18).ArraySet1(_155_v91, 23)
					(_nw18).ArraySet1(_155_v91, 24)
					(_nw18).ArraySet1(_155_v91, 25)
					_158_v94 = _nw18
					_158_v94 = _158_v94
					var _159_v95 _dafny.MultiSet
					_ = _159_v95
					_159_v95 = _dafny.MultiSetOf(_15_v4)
					var _160_v96 _dafny.Array
					_ = _160_v96
					var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
					_ = _nw19
					_160_v96 = _nw19
					var _161_v97 _dafny.Array
					_ = _161_v97
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(3)
					_ = _len0_4
					var _nw20 _dafny.Array
					_ = _nw20
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw20 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) D2 = (func(_162_v1 *C1) func(_dafny.Int) D2 {
							return func(_163_i11 _dafny.Int) D2 {
								return Companion_D2_.Create_DC9_((_162_v1).F6())
							}
						})(_10_v1)
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
					_161_v97 = _nw20
					var _164_v98 _dafny.Sequence
					_ = _164_v98
					_164_v98 = _dafny.SeqOf(_161_v97, _161_v97)
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_160_v96), 0))
					_ = _index33
					(_160_v96).ArraySet1(_164_v98, (_index33).Int())
					var _165_v99 _dafny.Set
					_ = _165_v99
					_165_v99 = _dafny.SetOf((_18_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_18_v7), 0))).Int()).(bool))
					var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_18_v7), 0))
					_ = _index34
					var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_160_v96), 0))
					_ = _index35
					var _rhs20 bool = (_18_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_18_v7), 0))).Int()).(bool)
					_ = _rhs20
					var _rhs21 _dafny.MultiSet = _159_v95
					_ = _rhs21
					var _rhs22 _dafny.Int = Companion_Default___.SafeModuloInt(_15_v4, _15_v4)
					_ = _rhs22
					var _rhs23 _dafny.Sequence = _dafny.SeqOf(_161_v97)
					_ = _rhs23
					var _rhs24 bool = (_dafny.SetOf((_10_v1).F7())).IsProperSubsetOf((func() _dafny.Set {
						if (_10_v1).F7() {
							return _165_v99
						}
						return _165_v99
					})())
					_ = _rhs24
					var _lhs14 _dafny.Array = _18_v7
					_ = _lhs14
					var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_18_v7), 0))
					_ = _lhs15
					var _lhs16 *GlobalState = _8_globalState
					_ = _lhs16
					var _lhs17 _dafny.Array = _160_v96
					_ = _lhs17
					var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_160_v96), 0))
					_ = _lhs18
					var _lhs19 *GlobalState = _8_globalState
					_ = _lhs19
					(_lhs14).ArraySet1(_rhs20, (_lhs15).Int())
					_159_v95 = _rhs21
					_lhs16.F1 = _rhs22
					(_lhs17).ArraySet1(_rhs23, (_lhs18).Int())
					_lhs19.F5 = _rhs24
				}
				if _dafny.Companion_Sequence_.Equal(_95_v47, _95_v47) {
					var _166_v100 _dafny.Map
					_ = _166_v100
					_166_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v4, (_10_v1).F7())
					_166_v100 = (_166_v100).Update(Companion_Default___.SafeDivisionInt(_15_v4, (_dafny.MultiSetOf(_95_v47)).Cardinality()), (_10_v1).F7())
					var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_18_v7), 0))
					_ = _index36
					(_18_v7).ArraySet1((_10_v1).F7(), (_index36).Int())
					var _167_v101 _dafny.Array
					_ = _167_v101
					var _nw21 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
					_ = _nw21
					_167_v101 = _nw21
					var _168_v102 _dafny.Map
					_ = _168_v102
					_168_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v101, (_10_v1).F7())
					var _169_v103 _dafny.MultiSet
					_ = _169_v103
					_169_v103 = _dafny.MultiSetOf(_15_v4, _15_v4, _15_v4)
					var _170_v104 _dafny.Set
					_ = _170_v104
					_170_v104 = _dafny.SetOf(!((_10_v1).F7()), (_10_v1).F7())
					var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_18_v7), 0))
					_ = _index37
					var _rhs25 bool = (_10_v1).F7()
					_ = _rhs25
					var _rhs26 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-735))
					_ = _rhs26
					var _rhs27 _dafny.Int = ((_168_v102).Update(_167_v101, (_dafny.MultiSetOf(_15_v4, _15_v4, (_170_v104).Cardinality(), _15_v4, _15_v4)).IsProperSubsetOf(_169_v103))).Cardinality()
					_ = _rhs27
					var _lhs20 _dafny.Array = _18_v7
					_ = _lhs20
					var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_18_v7), 0))
					_ = _lhs21
					var _lhs22 *GlobalState = _8_globalState
					_ = _lhs22
					var _lhs23 *GlobalState = _8_globalState
					_ = _lhs23
					(_lhs20).ArraySet1(_rhs25, (_lhs21).Int())
					_lhs22.F1 = _rhs26
					_lhs23.F0 = _rhs27
					var _171_v105 _dafny.Array
					_ = _171_v105
					var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
					_ = _nw22
					_171_v105 = _nw22
					var _172_v106 _dafny.Map
					_ = _172_v106
					_172_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v105, (_10_v1).F7())
					_172_v106 = (_172_v106).Update(_171_v105, (_10_v1).F7())
					var _173_v107 _dafny.Map
					_ = _173_v107
					_173_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_10_v1, _dafny.IntOfInt64(806))
					var _174_v108 _dafny.Sequence
					_ = _174_v108
					_174_v108 = _dafny.SeqOf((_10_v1).F7(), Companion_Default___.Fm1(_8_globalState), (_10_v1).F7(), !((_10_v1).F7()), (_18_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_18_v7), 0))).Int()).(bool))
					var _175_v109 *C0
					_ = _175_v109
					var _nw23 *C0 = New_C0_()
					_ = _nw23
					_nw23.Ctor__((_16_v5).Update((func() _dafny.Int {
						if (_173_v107).Contains(_10_v1) {
							return (_173_v107).Get(_10_v1).(_dafny.Int)
						}
						return _dafny.IntOfUint32((_174_v108).Cardinality())
					})(), _15_v4), (_18_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_18_v7), 0))).Int()).(bool))
					_175_v109 = _nw23
					_15_v4 = _15_v4
				} else {
					var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_18_v7), 0))
					_ = _index38
					(_18_v7).ArraySet1(Companion_Default___.Fm1(_8_globalState), (_index38).Int())
					var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_18_v7), 0))
					_ = _index39
					(_18_v7).ArraySet1((_10_v1).F7(), (_index39).Int())
					var _176_v110 _dafny.Map
					_ = _176_v110
					_176_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v4, _95_v47)
					var _177_v111 _dafny.Map
					_ = _177_v111
					_177_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-45), _176_v110)
					_177_v111 = (_177_v111).Update(_15_v4, _176_v110)
					var _178_v112 D1
					_ = _178_v112
					_178_v112 = Companion_D1_.Create_DC4_((_10_v1).F7())
					var _pat_let_tv1 = _18_v7
					_ = _pat_let_tv1
					var _pat_let_tv2 = _18_v7
					_ = _pat_let_tv2
					_178_v112 = func(_pat_let2_0 D1) D1 {
						return func(_179_dt__update__tmp_h1 D1) D1 {
							return func(_pat_let3_0 bool) D1 {
								return func(_180_dt__update_hcf10_h0 bool) D1 {
									return Companion_D1_.Create_DC4_(_180_dt__update_hcf10_h0)
								}(_pat_let3_0)
							}((_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(bool))
						}(_pat_let2_0)
					}(_178_v112)
					(_8_globalState).F0 = (_dafny.Zero).Minus(_15_v4)
					var _181_v113 *C1
					_ = _181_v113
					var _nw24 *C1 = New_C1_()
					_ = _nw24
					_nw24.Ctor__(_dafny.CodePoint('j'), (_18_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_18_v7), 0))).Int()).(bool))
					_181_v113 = _nw24
				}
				var _182_v114 _dafny.Sequence
				_ = _182_v114
				_182_v114 = _dafny.SeqOf(Companion_Default___.Fm2((_10_v1).F7(), _8_globalState), _dafny.IntOfInt64(555))
				var _183_v116 _dafny.Set
				_ = _183_v116
				_183_v116 = _dafny.SetOf(_15_v4)
				(_8_globalState).F5 = (_183_v116).IsProperSubsetOf(func() _dafny.Set {
					var _coll4 = _dafny.NewBuilder()
					_ = _coll4
					for _iter5 := _dafny.Iterate((_182_v114).Elements()); ; {
						_compr_4, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _184_v115 _dafny.Int
						_184_v115 = interface{}(_compr_4).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_182_v114, _184_v115) {
							_coll4.Add((_184_v115).Minus(_dafny.IntOfInt64(-668)))
						}
					}
					return _coll4.ToSet()
				}())
				var _185_v117 _dafny.Array
				_ = _185_v117
				var _nwElement0_5 _dafny.Sequence = _95_v47
				_ = _nwElement0_5
				var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(2))
				_ = _nw25
				(_nw25).ArraySet1(_nwElement0_5, 0)
				(_nw25).ArraySet1((func() _dafny.Sequence {
					if !((_10_v1).F7()) {
						return _dafny.UnicodeSeqOfUtf8Bytes("tdaboijva")
					}
					return _95_v47
				})(), 1)
				_185_v117 = _nw25
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_185_v117), 0))
				_ = _index40
				(_185_v117).ArraySet1(_95_v47, (_index40).Int())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_185_v117), 0))
				_ = _index41
				(_185_v117).ArraySet1((_10_v1).Fm0((_10_v1).Fm0(_dafny.UnicodeSeqOfUtf8Bytes("gqains"), _9_v0, _8_globalState), (_10_v1).F6(), _8_globalState), (_index41).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _186_v118 _dafny.MultiSet
	_ = _186_v118
	_186_v118 = _dafny.MultiSetOf(_17_v6)
	var _187_v119 _dafny.Array
	_ = _187_v119
	var _nwElement0_6 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-948))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}((func(_188_v1 *C1) func(_dafny.Int) _dafny.CodePoint {
		return func(_189_i12 _dafny.Int) _dafny.CodePoint {
			return (_188_v1).F6()
		}
	})(_10_v1)))).Cardinality())
	_ = _nwElement0_6
	var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(11))
	_ = _nw26
	(_nw26).ArraySet1(_nwElement0_6, 0)
	(_nw26).ArraySet1(_15_v4, 1)
	(_nw26).ArraySet1(_15_v4, 2)
	(_nw26).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(621), _dafny.IntOfInt64(665)), 3)
	(_nw26).ArraySet1((_17_v6).Cardinality(), 4)
	(_nw26).ArraySet1(Companion_Default___.Fm2((func() bool {
		if (_17_v6).Contains(false) {
			return (_17_v6).Get(false).(bool)
		}
		return (_10_v1).F7()
	})(), _8_globalState), 5)
	(_nw26).ArraySet1(_15_v4, 6)
	(_nw26).ArraySet1(_15_v4, 7)
	(_nw26).ArraySet1(_15_v4, 8)
	(_nw26).ArraySet1(_15_v4, 9)
	(_nw26).ArraySet1(((Companion_Default___.Fm9((_10_v1).F7(), (_10_v1).F7(), _8_globalState)).Difference(_186_v118)).Cardinality(), 10)
	_187_v119 = _nw26
	var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
	_ = _index42
	(_187_v119).ArraySet1(_15_v4, (_index42).Int())
	var _190_v120 _dafny.MultiSet
	_ = _190_v120
	_190_v120 = _dafny.MultiSetOf(_15_v4, _15_v4)
	var _191_v121 _dafny.Sequence
	_ = _191_v121
	_191_v121 = _dafny.SeqOf(_dafny.IntOfUint32((_95_v47).Cardinality()))
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
	_ = _index43
	var _rhs28 _dafny.Int = _15_v4
	_ = _rhs28
	var _rhs29 _dafny.Int = (func() _dafny.Int {
		if (_190_v120).Contains(_dafny.IntOfInt64(-694)) {
			return (_190_v120).Multiplicity(_dafny.IntOfInt64(-694))
		}
		return (_191_v121).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2((_10_v1).F7(), _8_globalState), _dafny.IntOfUint32((_191_v121).Cardinality()))).Uint32()).(_dafny.Int)
	})()
	_ = _rhs29
	var _rhs30 _dafny.Int = _15_v4
	_ = _rhs30
	var _rhs31 bool = !(!((_10_v1).F7()) || (!((_10_v1).F7())))
	_ = _rhs31
	var _lhs24 _dafny.Array = _187_v119
	_ = _lhs24
	var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
	_ = _lhs25
	var _lhs26 *GlobalState = _8_globalState
	_ = _lhs26
	(_lhs24).ArraySet1(_rhs28, (_lhs25).Int())
	_15_v4 = _rhs29
	_15_v4 = _rhs30
	_lhs26.F5 = _rhs31
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
	_ = _index44
	(_187_v119).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(872))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}((func(_192_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_193_i13 _dafny.Int) _dafny.CodePoint {
			return _192_v0
		}
	})(_9_v0)))).Cardinality()), (_index44).Int())
	var _194_v122 _dafny.Sequence
	_ = _194_v122
	_194_v122 = _dafny.SeqOf(!((_10_v1).F7()) || ((_10_v1).F7()))
	(_8_globalState).F5 = (_194_v122).Select((Companion_Default___.SafeIndex(_15_v4, _dafny.IntOfUint32((_194_v122).Cardinality()))).Uint32()).(bool)
	var _195_v123 _dafny.Map
	_ = _195_v123
	_195_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_10_v1).F7(), (_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int))
	var _196_v124 D0
	_ = _196_v124
	_196_v124 = Companion_D0_.Create_DC3_(_15_v4)
	var _197_v125 _dafny.Map
	_ = _197_v125
	_197_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_194_v122).Select((Companion_Default___.SafeIndex(((_190_v120).Update((func() _dafny.Int {
		if (_195_v123).Contains(!(!((_10_v1).F7()))) {
			return (_195_v123).Get(!(!((_10_v1).F7()))).(_dafny.Int)
		}
		return _15_v4
	})(), Companion_Default___.Abs((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)))).Cardinality(), _dafny.IntOfUint32((_194_v122).Cardinality()))).Uint32()).(bool), _196_v124)
	_197_v125 = (_197_v125).Update((_10_v1).F7(), _196_v124)
	if (_dafny.IntOfUint32((_95_v47).Cardinality())).Cmp(_15_v4) <= 0 {
		var _198_v126 _dafny.Set
		_ = _198_v126
		_198_v126 = _dafny.SetOf(_dafny.IntOfInt64(132), _15_v4)
		_198_v126 = (_198_v126).Intersection((func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(928), _dafny.IntOfInt64(795))); ; {
				_compr_5, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _199_v127 _dafny.Int
				_199_v127 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(928)).Cmp(_199_v127) <= 0) && ((_199_v127).Cmp(_dafny.IntOfInt64(795)) < 0) {
					_coll5.Add((_199_v127).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(347))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}((func(_200_v119 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_201_i14 _dafny.Int) _dafny.Int {
							return (_200_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_200_v119), 0))).Int()).(_dafny.Int)
						}
					})(_187_v119)))).Cardinality())))
				}
			}
			return _coll5.ToSet()
		}()).Difference(_198_v126))
		(_8_globalState).F5 = !((_10_v1).F7())
		var _202_v128 _dafny.Sequence
		_ = _202_v128
		_202_v128 = _dafny.SeqOf(_16_v5)
		var _203_v130 _dafny.Map
		_ = _203_v130
		_203_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int), (_10_v1).F7())
		var _204_v131 _dafny.Sequence
		_ = _204_v131
		_204_v131 = _dafny.SeqOf(_203_v130)
		var _205_v132 *C0
		_ = _205_v132
		var _nw27 *C0 = New_C0_()
		_ = _nw27
		_nw27.Ctor__(((_202_v128).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_191_v121).Cardinality()), _dafny.IntOfUint32((_202_v128).Cardinality()))).Uint32()).(_dafny.Map)).Update(_15_v4, (func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter7 := _dafny.Iterate((_204_v131).Elements()); ; {
				_compr_6, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _206_v129 _dafny.Map
				_206_v129 = interface{}(_compr_6).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_204_v131, _206_v129) {
					_coll6.Add(_206_v129, (func() bool {
						if (_17_v6).Contains(true) {
							return (_17_v6).Get(true).(bool)
						}
						return (_10_v1).F7()
					})())
				}
			}
			return _coll6.ToMap()
		}()).Cardinality()), (_10_v1).F7())
		_205_v132 = _nw27
		var _207_v133 D1
		_ = _207_v133
		_207_v133 = Companion_D1_.Create_DC5_(Companion_Default___.SafeDivisionInt((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int), _15_v4), (_205_v132.F8).Merge(_16_v5), (func() _dafny.Int {
			if (_205_v132.F8).Contains(_15_v4) {
				return (_205_v132.F8).Get(_15_v4).(_dafny.Int)
			}
			return _15_v4
		})())
		var _source2 D1 = _207_v133
		_ = _source2
		if _source2.Is_DC5() {
			var _208___mcc_h17 _dafny.Int = _source2.Get_().(D1_DC5).Cf11
			_ = _208___mcc_h17
			var _209___mcc_h18 _dafny.Map = _source2.Get_().(D1_DC5).Cf12
			_ = _209___mcc_h18
			var _210___mcc_h19 _dafny.Int = _source2.Get_().(D1_DC5).Cf13
			_ = _210___mcc_h19
			var _211_cf13 _dafny.Int = _210___mcc_h19
			_ = _211_cf13
			var _212_cf12 _dafny.Map = _209___mcc_h18
			_ = _212_cf12
			var _213_cf11 _dafny.Int = _208___mcc_h17
			_ = _213_cf11
			var _214_v134 _dafny.Int
			_ = _214_v134
			var _215_v135 _dafny.Int
			_ = _215_v135
			var _216_v136 _dafny.Array
			_ = _216_v136
			var _out15 _dafny.Int
			_ = _out15
			var _out16 _dafny.Int
			_ = _out16
			var _out17 _dafny.Array
			_ = _out17
			_out15, _out16, _out17 = (_10_v1).M0((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int), _213_cf11, _8_globalState)
			_214_v134 = _out15
			_215_v135 = _out16
			_216_v136 = _out17
			var _217_v137 *C0
			_ = _217_v137
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_191_v121).Cardinality()), _dafny.IntOfUint32((_95_v47).Cardinality()))).Merge(_212_cf12), true)
			_217_v137 = _nw28
			var _218_v138 _dafny.MultiSet
			_ = _218_v138
			_218_v138 = _dafny.MultiSetOf((_10_v1).F6())
			_218_v138 = _218_v138
			_9_v0 = (_95_v47).Select((Companion_Default___.SafeIndex(_213_cf11, _dafny.IntOfUint32((_95_v47).Cardinality()))).Uint32()).(_dafny.CodePoint)
		} else if _source2.Is_DC6() {
			var _219___mcc_h20 _dafny.CodePoint = _source2.Get_().(D1_DC6).Cf14
			_ = _219___mcc_h20
			var _220___mcc_h21 _dafny.CodePoint = _source2.Get_().(D1_DC6).Cf15
			_ = _220___mcc_h21
			var _221_cf15 _dafny.CodePoint = _220___mcc_h21
			_ = _221_cf15
			var _222_cf14 _dafny.CodePoint = _219___mcc_h20
			_ = _222_cf14
			var _rhs32 bool = (_205_v132).F9()
			_ = _rhs32
			var _rhs33 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(596), _15_v4)
			_ = _rhs33
			var _lhs27 *GlobalState = _8_globalState
			_ = _lhs27
			var _lhs28 *GlobalState = _8_globalState
			_ = _lhs28
			_lhs27.F5 = _rhs32
			_lhs28.F1 = _rhs33
			var _223_v139 _dafny.Array
			_ = _223_v139
			var _nwElement0_7 *C0 = _205_v132
			_ = _nwElement0_7
			var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(12))
			_ = _nw29
			(_nw29).ArraySet1(_nwElement0_7, 0)
			(_nw29).ArraySet1(_205_v132, 1)
			(_nw29).ArraySet1(_205_v132, 2)
			(_nw29).ArraySet1(_205_v132, 3)
			(_nw29).ArraySet1(_205_v132, 4)
			(_nw29).ArraySet1(_205_v132, 5)
			(_nw29).ArraySet1((func() *C0 {
				if true {
					return _205_v132
				}
				return _205_v132
			})(), 6)
			(_nw29).ArraySet1(_205_v132, 7)
			(_nw29).ArraySet1(_205_v132, 8)
			(_nw29).ArraySet1(_205_v132, 9)
			(_nw29).ArraySet1(_205_v132, 10)
			(_nw29).ArraySet1(_205_v132, 11)
			_223_v139 = _nw29
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_223_v139), 0))
			_ = _index45
			(_223_v139).ArraySet1(_205_v132, (_index45).Int())
			var _224_v140 _dafny.Array
			_ = _224_v140
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
			_ = _nw30
			_224_v140 = _nw30
			var _225_v141 _dafny.Array
			_ = _225_v141
			var _nwElement0_8 _dafny.Array = _224_v140
			_ = _nwElement0_8
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(17))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_8, 0)
			(_nw31).ArraySet1(_224_v140, 1)
			(_nw31).ArraySet1(_224_v140, 2)
			(_nw31).ArraySet1(_224_v140, 3)
			(_nw31).ArraySet1(_224_v140, 4)
			(_nw31).ArraySet1(_224_v140, 5)
			(_nw31).ArraySet1(_224_v140, 6)
			(_nw31).ArraySet1(_224_v140, 7)
			(_nw31).ArraySet1(_224_v140, 8)
			(_nw31).ArraySet1((func() _dafny.Array {
				if (_10_v1).F7() {
					return _224_v140
				}
				return _224_v140
			})(), 9)
			(_nw31).ArraySet1(_224_v140, 10)
			(_nw31).ArraySet1(_224_v140, 11)
			(_nw31).ArraySet1(_224_v140, 12)
			(_nw31).ArraySet1(_224_v140, 13)
			(_nw31).ArraySet1(_224_v140, 14)
			(_nw31).ArraySet1(_224_v140, 15)
			(_nw31).ArraySet1(_224_v140, 16)
			_225_v141 = _nw31
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_223_v139), 0))
			_ = _index46
			var _rhs34 *C0 = _205_v132
			_ = _rhs34
			var _rhs35 _dafny.Int = (_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)
			_ = _rhs35
			var _rhs36 _dafny.Array = _225_v141
			_ = _rhs36
			var _rhs37 _dafny.Int = (_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)
			_ = _rhs37
			var _rhs38 bool = !((_10_v1).F7())
			_ = _rhs38
			var _lhs29 _dafny.Array = _223_v139
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_223_v139), 0))
			_ = _lhs30
			var _lhs31 *GlobalState = _8_globalState
			_ = _lhs31
			var _lhs32 *GlobalState = _8_globalState
			_ = _lhs32
			var _lhs33 *GlobalState = _8_globalState
			_ = _lhs33
			(_lhs29).ArraySet1(_rhs34, (_lhs30).Int())
			_lhs31.F1 = _rhs35
			_225_v141 = _rhs36
			_lhs32.F1 = _rhs37
			_lhs33.F5 = _rhs38
			_221_cf15 = (_10_v1).F6()
			(_8_globalState).F5 = ((_205_v132).F9()) || ((_10_v1).F7())
		} else if _source2.Is_DC7() {
			var _226___mcc_h22 _dafny.CodePoint = _source2.Get_().(D1_DC7).Cf16
			_ = _226___mcc_h22
			var _227_cf16 _dafny.CodePoint = _226___mcc_h22
			_ = _227_cf16
			var _228_v142 _dafny.Int
			_ = _228_v142
			var _229_v143 _dafny.Int
			_ = _229_v143
			var _230_v144 _dafny.Array
			_ = _230_v144
			var _out18 _dafny.Int
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			var _out20 _dafny.Array
			_ = _out20
			_out18, _out19, _out20 = (_10_v1).M0((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int), _15_v4, _8_globalState)
			_228_v142 = _out18
			_229_v143 = _out19
			_230_v144 = _out20
			var _231_v145 _dafny.Map
			_ = _231_v145
			_231_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_10_v1).F7(), (_10_v1).F7(), (_10_v1).F7(), (_10_v1).F7()), _15_v4)
			_231_v145 = (_231_v145).Update(Companion_Default___.Fm10(_dafny.IntOfInt64(105), _dafny.SetOf((_10_v1).F7(), false, (_10_v1).F7(), (_205_v132).F9(), !(!(!((_10_v1).F7())))), _8_globalState), (func() _dafny.Int {
				if !((_10_v1).F7()) {
					return _15_v4
				}
				return (_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)
			})())
			var _232_v146 D2
			_ = _232_v146
			_232_v146 = Companion_D2_.Create_DC10_((_205_v132).F9(), (_10_v1).F7())
			_232_v146 = Companion_D2_.Create_DC10_((_205_v132).F9(), (_10_v1).F7())
			var _233_v147 _dafny.MultiSet
			_ = _233_v147
			_233_v147 = _dafny.MultiSetOf((_10_v1).F7(), (_205_v132).F9())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_18_v7), 0))
			_ = _index47
			(_18_v7).ArraySet1(((_dafny.Zero).Minus((_233_v147).Cardinality())).Cmp(_15_v4) >= 0, (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_18_v7), 0))
			_ = _index48
			var _rhs39 bool = ((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)).Cmp((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)) != 0
			_ = _rhs39
			var _rhs40 bool = (_205_v132).F9()
			_ = _rhs40
			var _rhs41 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_228_v142, (_228_v142).Minus((_195_v123).Cardinality())))
			_ = _rhs41
			var _lhs34 *GlobalState = _8_globalState
			_ = _lhs34
			var _lhs35 _dafny.Array = _18_v7
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_18_v7), 0))
			_ = _lhs36
			var _lhs37 *GlobalState = _8_globalState
			_ = _lhs37
			_lhs34.F5 = _rhs39
			(_lhs35).ArraySet1(_rhs40, (_lhs36).Int())
			_lhs37.F0 = _rhs41
		} else {
			var _234___mcc_h23 bool = _source2.Get_().(D1_DC4).Cf10
			_ = _234___mcc_h23
			var _235_cf10 bool = _234___mcc_h23
			_ = _235_cf10
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_18_v7), 0))
			_ = _index49
			(_18_v7).ArraySet1(_235_cf10, (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_18_v7), 0))
			_ = _index50
			(_18_v7).ArraySet1(_235_cf10, (_index50).Int())
			var _236_v148 D2
			_ = _236_v148
			_236_v148 = Companion_D2_.Create_DC10_(true, (_205_v132).Fm4(_8_globalState))
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_18_v7), 0))
			_ = _index51
			(_18_v7).ArraySet1((_236_v148).Dtor_cf20(), (_index51).Int())
			_9_v0 = (_10_v1).F6()
			(_8_globalState).F5 = (_235_cf10) == ((_205_v132).F9())
		}
		var _237_v149 _dafny.Array
		_ = _237_v149
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw32
		_237_v149 = _nw32
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_237_v149), 0))
		_ = _index52
		(_237_v149).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_95_v47, _95_v47), (_index52).Int())
		var _238_v150 D2
		_ = _238_v150
		_238_v150 = Companion_D2_.Create_DC10_((_10_v1).F7(), (_10_v1).F7())
		var _239_v151 D3
		_ = _239_v151
		_239_v151 = Companion_D3_.Create_DC14_(_10_v1, (_10_v1).F6(), _10_v1, _205_v132, (_205_v132).F9())
		var _240_v152 _dafny.Set
		_ = _240_v152
		_240_v152 = _dafny.SetOf(Companion_D3_.Create_DC14_(_10_v1, (_10_v1).F6(), _10_v1, _205_v132, (_205_v132).F9()), _239_v151)
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
		_ = _index53
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_237_v149), 0))
		_ = _index54
		var _rhs42 bool = !((!(!((_238_v150).Dtor_cf19()))) && ((_240_v152).IsDisjointFrom(_240_v152)))
		_ = _rhs42
		var _rhs43 _dafny.Int = (_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)
		_ = _rhs43
		var _rhs44 _dafny.Sequence = _95_v47
		_ = _rhs44
		var _lhs38 *GlobalState = _8_globalState
		_ = _lhs38
		var _lhs39 _dafny.Array = _187_v119
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
		_ = _lhs40
		var _lhs41 _dafny.Array = _237_v149
		_ = _lhs41
		var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_237_v149), 0))
		_ = _lhs42
		_lhs38.F5 = _rhs42
		(_lhs39).ArraySet1(_rhs43, (_lhs40).Int())
		(_lhs41).ArraySet1(_rhs44, (_lhs42).Int())
	} else {
		var _241_v153 _dafny.Sequence
		_ = _241_v153
		_241_v153 = _dafny.SeqOf(_95_v47, _dafny.Companion_Sequence_.Concatenate(_95_v47, _95_v47), _95_v47)
		var _242_v154 D4
		_ = _242_v154
		_242_v154 = Companion_D4_.Create_DC16_(_241_v153)
		_241_v153 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_241_v153, _dafny.Companion_Sequence_.Update((_242_v154).Dtor_cf31(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32(((_242_v154).Dtor_cf31()).Cardinality()))).Uint32(), _95_v47)), _241_v153)
		var _243_v155 D5
		_ = _243_v155
		_243_v155 = Companion_D5_.Create_DC18_(_187_v119)
		_187_v119 = (_243_v155).Dtor_cf37()
		var _244_v156 _dafny.Int
		_ = _244_v156
		var _245_v157 _dafny.Int
		_ = _245_v157
		var _246_v158 _dafny.Array
		_ = _246_v158
		var _out21 _dafny.Int
		_ = _out21
		var _out22 _dafny.Int
		_ = _out22
		var _out23 _dafny.Array
		_ = _out23
		_out21, _out22, _out23 = (_10_v1).M0((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
			if (_195_v123).Contains((_10_v1).F7()) {
				return (_195_v123).Get((_10_v1).F7()).(_dafny.Int)
			}
			return _15_v4
		})(), _8_globalState)
		_244_v156 = _out21
		_245_v157 = _out22
		_246_v158 = _out23
		(_8_globalState).F5 = (_10_v1).F7()
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_18_v7), 0))
		_ = _index55
		(_18_v7).ArraySet1((_10_v1).F7(), (_index55).Int())
		var _247_v159 *C0
		_ = _247_v159
		var _nw33 *C0 = New_C0_()
		_ = _nw33
		_nw33.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int), _245_v157), (_10_v1).F7())
		_247_v159 = _nw33
		var _248_v160 _dafny.Sequence
		_ = _248_v160
		_248_v160 = _dafny.SeqOf(_247_v159, _247_v159, _247_v159)
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_18_v7), 0))
		_ = _index56
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
		_ = _index57
		var _rhs45 _dafny.Int = _245_v157
		_ = _rhs45
		var _rhs46 _dafny.CodePoint = Companion_Default___.Fm6(_8_globalState)
		_ = _rhs46
		var _rhs47 bool = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_248_v160).Cardinality()))).Cmp(_15_v4) < 0
		_ = _rhs47
		var _rhs48 _dafny.Int = (_245_v157).Minus(_245_v157)
		_ = _rhs48
		var _rhs49 _dafny.Int = _dafny.IntOfInt64(943)
		_ = _rhs49
		var _lhs43 *GlobalState = _8_globalState
		_ = _lhs43
		var _lhs44 _dafny.Array = _18_v7
		_ = _lhs44
		var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_18_v7), 0))
		_ = _lhs45
		var _lhs46 _dafny.Array = _187_v119
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))
		_ = _lhs47
		_lhs43.F1 = _rhs45
		_9_v0 = _rhs46
		(_lhs44).ArraySet1(_rhs47, (_lhs45).Int())
		_244_v156 = _rhs48
		(_lhs46).ArraySet1(_rhs49, (_lhs47).Int())
	}
	var _249_v161 _dafny.Set
	_ = _249_v161
	_249_v161 = _dafny.SetOf(_dafny.IntOfInt64(110), (_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int))
	var _250_v162 _dafny.MultiSet
	_ = _250_v162
	_250_v162 = _dafny.MultiSetOf(_95_v47)
	var _251_v163 _dafny.Int
	_ = _251_v163
	var _252_v164 _dafny.Int
	_ = _252_v164
	var _253_v165 _dafny.Array
	_ = _253_v165
	var _out24 _dafny.Int
	_ = _out24
	var _out25 _dafny.Int
	_ = _out25
	var _out26 _dafny.Array
	_ = _out26
	_out24, _out25, _out26 = (_10_v1).M0(((_249_v161).Cardinality()).Minus((_250_v162).Cardinality()), (func() _dafny.Int {
		if (_10_v1).F7() {
			return (_187_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_187_v119), 0))).Int()).(_dafny.Int)
		}
		return _15_v4
	})(), _8_globalState)
	_251_v163 = _out24
	_252_v164 = _out25
	_253_v165 = _out26
	_252_v164 = _dafny.IntOfInt64(749)
	_dafny.Print(_8_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_8_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_8_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_8_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_8_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_8_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_9_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_10_v1).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_10_v1).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_13_v2).Dtor_cf14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_13_v2).Dtor_cf15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_14_v3, _dafny.SeqOf(Companion_D1_.Create_DC6_(_dafny.CodePoint('o'), _dafny.CodePoint('k')), Companion_D1_.Create_DC6_(_dafny.CodePoint('o'), _dafny.CodePoint('k')), Companion_D1_.Create_DC6_(_dafny.CodePoint('o'), _dafny.CodePoint('k')), Companion_D1_.Create_DC6_(_dafny.CodePoint('o'), _dafny.CodePoint('k')), Companion_D1_.Create_DC6_(_dafny.CodePoint('o'), _dafny.CodePoint('k')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_15_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_16_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(22), _dafny.IntOfInt64(22))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_17_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v7).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf6()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(22), _dafny.IntOfInt64(22))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf7()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_21_v8).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_95_v47.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_149_i10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v118).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v119).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v120).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(22), _dafny.IntOfInt64(22))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_191_v121, _dafny.SeqOf(_dafny.IntOfInt64(6))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_194_v122, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v123).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(872))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_196_v124).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v125).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D0_.Create_DC3_(_dafny.IntOfInt64(22))).UpdateUnsafe(false, Companion_D0_.Create_DC3_(_dafny.IntOfInt64(22)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_249_v161).Equals(_dafny.SetOf(_dafny.IntOfInt64(110), _dafny.IntOfInt64(872))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v162).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("vintro"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_251_v163)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_252_v164)
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
	Cf1 _dafny.Int
	Cf2 _dafny.Set
	Cf3 _dafny.Int
	Cf4 _dafny.Array
	Cf5 _dafny.CodePoint
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Set, Cf3 _dafny.Int, Cf4 _dafny.Array, Cf5 _dafny.CodePoint) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf6 _dafny.Map
	Cf7 _dafny.Map
	Cf8 _dafny.Array
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf6 _dafny.Map, Cf7 _dafny.Map, Cf8 _dafny.Array) D0 {
	return D0{D0_DC2{Cf6, Cf7, Cf8}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf9 _dafny.Int
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf9 _dafny.Int) D0 {
	return D0{D0_DC3{Cf9}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.EmptySet, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.CodePoint('D'))
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Set {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Map {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D0_DC2).Cf8
}

func (_this D0) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D0_DC3).Cf9
}

func (_this D0) Dtor_cf0() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf9) + ")"
		}
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
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Equals(data2.Cf2) && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4 && data1.Cf5 == data2.Cf5
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7.Equals(data2.Cf7) && data1.Cf8 == data2.Cf8
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0
		}
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

type D1_DC5 struct {
	Cf11 _dafny.Int
	Cf12 _dafny.Map
	Cf13 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf11 _dafny.Int, Cf12 _dafny.Map, Cf13 _dafny.Int) D1 {
	return D1{D1_DC5{Cf11, Cf12, Cf13}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf14 _dafny.CodePoint
	Cf15 _dafny.CodePoint
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf14 _dafny.CodePoint, Cf15 _dafny.CodePoint) D1 {
	return D1{D1_DC6{Cf14, Cf15}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC7 struct {
	Cf16 _dafny.CodePoint
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_(Cf16 _dafny.CodePoint) D1 {
	return D1{D1_DC7{Cf16}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

type D1_DC4 struct {
	Cf10 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf10 bool) D1 {
	return D1{D1_DC4{Cf10}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(_dafny.Zero, _dafny.EmptyMap, _dafny.Zero)
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf13
}

func (_this D1) Dtor_cf14() _dafny.CodePoint {
	return _this.Get_().(D1_DC6).Cf14
}

func (_this D1) Dtor_cf15() _dafny.CodePoint {
	return _this.Get_().(D1_DC6).Cf15
}

func (_this D1) Dtor_cf16() _dafny.CodePoint {
	return _this.Get_().(D1_DC7).Cf16
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12.Equals(data2.Cf12) && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
		}
	case D1_DC7:
		{
			data2, ok := other.Get_().(D1_DC7)
			return ok && data1.Cf16 == data2.Cf16
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf10 == data2.Cf10
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

type D2_DC9 struct {
	Cf18 _dafny.CodePoint
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf18 _dafny.CodePoint) D2 {
	return D2{D2_DC9{Cf18}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC10 struct {
	Cf19 bool
	Cf20 bool
}

func (D2_DC10) isD2() {}

func (CompanionStruct_D2_) Create_DC10_(Cf19 bool, Cf20 bool) D2 {
	return D2{D2_DC10{Cf19, Cf20}}
}

func (_this D2) Is_DC10() bool {
	_, ok := _this.Get_().(D2_DC10)
	return ok
}

type D2_DC11 struct {
	Cf21 _dafny.Int
	Cf22 _dafny.Int
}

func (D2_DC11) isD2() {}

func (CompanionStruct_D2_) Create_DC11_(Cf21 _dafny.Int, Cf22 _dafny.Int) D2 {
	return D2{D2_DC11{Cf21, Cf22}}
}

func (_this D2) Is_DC11() bool {
	_, ok := _this.Get_().(D2_DC11)
	return ok
}

type D2_DC8 struct {
	Cf17 *C1
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf17 *C1) D2 {
	return D2{D2_DC8{Cf17}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC12 struct {
	Cf23 D2
}

func (D2_DC12) isD2() {}

func (CompanionStruct_D2_) Create_DC12_(Cf23 D2) D2 {
	return D2{D2_DC12{Cf23}}
}

func (_this D2) Is_DC12() bool {
	_, ok := _this.Get_().(D2_DC12)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC9_(_dafny.CodePoint('D'))
}

func (_this D2) Dtor_cf18() _dafny.CodePoint {
	return _this.Get_().(D2_DC9).Cf18
}

func (_this D2) Dtor_cf19() bool {
	return _this.Get_().(D2_DC10).Cf19
}

func (_this D2) Dtor_cf20() bool {
	return _this.Get_().(D2_DC10).Cf20
}

func (_this D2) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D2_DC11).Cf21
}

func (_this D2) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D2_DC11).Cf22
}

func (_this D2) Dtor_cf17() *C1 {
	return _this.Get_().(D2_DC8).Cf17
}

func (_this D2) Dtor_cf23() D2 {
	return _this.Get_().(D2_DC12).Cf23
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D2_DC10:
		{
			return "D2.DC10" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D2_DC11:
		{
			return "D2.DC11" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC12:
		{
			return "D2.DC12" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf18 == data2.Cf18
		}
	case D2_DC10:
		{
			data2, ok := other.Get_().(D2_DC10)
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20
		}
	case D2_DC11:
		{
			data2, ok := other.Get_().(D2_DC11)
			return ok && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf17 == data2.Cf17
		}
	case D2_DC12:
		{
			data2, ok := other.Get_().(D2_DC12)
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D3_DC14 struct {
	Cf25 *C1
	Cf26 _dafny.CodePoint
	Cf27 *C1
	Cf28 *C0
	Cf29 bool
}

func (D3_DC14) isD3() {}

func (CompanionStruct_D3_) Create_DC14_(Cf25 *C1, Cf26 _dafny.CodePoint, Cf27 *C1, Cf28 *C0, Cf29 bool) D3 {
	return D3{D3_DC14{Cf25, Cf26, Cf27, Cf28, Cf29}}
}

func (_this D3) Is_DC14() bool {
	_, ok := _this.Get_().(D3_DC14)
	return ok
}

type D3_DC13 struct {
	Cf24 _dafny.Array
}

func (D3_DC13) isD3() {}

func (CompanionStruct_D3_) Create_DC13_(Cf24 _dafny.Array) D3 {
	return D3{D3_DC13{Cf24}}
}

func (_this D3) Is_DC13() bool {
	_, ok := _this.Get_().(D3_DC13)
	return ok
}

type D3_DC15 struct {
	Cf30 D3
}

func (D3_DC15) isD3() {}

func (CompanionStruct_D3_) Create_DC15_(Cf30 D3) D3 {
	return D3{D3_DC15{Cf30}}
}

func (_this D3) Is_DC15() bool {
	_, ok := _this.Get_().(D3_DC15)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC14_((*C1)(nil), _dafny.CodePoint('D'), (*C1)(nil), (*C0)(nil), false)
}

func (_this D3) Dtor_cf25() *C1 {
	return _this.Get_().(D3_DC14).Cf25
}

func (_this D3) Dtor_cf26() _dafny.CodePoint {
	return _this.Get_().(D3_DC14).Cf26
}

func (_this D3) Dtor_cf27() *C1 {
	return _this.Get_().(D3_DC14).Cf27
}

func (_this D3) Dtor_cf28() *C0 {
	return _this.Get_().(D3_DC14).Cf28
}

func (_this D3) Dtor_cf29() bool {
	return _this.Get_().(D3_DC14).Cf29
}

func (_this D3) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D3_DC13).Cf24
}

func (_this D3) Dtor_cf30() D3 {
	return _this.Get_().(D3_DC15).Cf30
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC14:
		{
			return "D3.DC14" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D3_DC13:
		{
			return "D3.DC13" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D3_DC15:
		{
			return "D3.DC15" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC14:
		{
			data2, ok := other.Get_().(D3_DC14)
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29
		}
	case D3_DC13:
		{
			data2, ok := other.Get_().(D3_DC13)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D3_DC15:
		{
			data2, ok := other.Get_().(D3_DC15)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D4_DC17 struct {
	Cf32 _dafny.Int
	Cf33 bool
	Cf34 _dafny.Int
	Cf35 _dafny.Int
	Cf36 *C1
}

func (D4_DC17) isD4() {}

func (CompanionStruct_D4_) Create_DC17_(Cf32 _dafny.Int, Cf33 bool, Cf34 _dafny.Int, Cf35 _dafny.Int, Cf36 *C1) D4 {
	return D4{D4_DC17{Cf32, Cf33, Cf34, Cf35, Cf36}}
}

func (_this D4) Is_DC17() bool {
	_, ok := _this.Get_().(D4_DC17)
	return ok
}

type D4_DC16 struct {
	Cf31 _dafny.Sequence
}

func (D4_DC16) isD4() {}

func (CompanionStruct_D4_) Create_DC16_(Cf31 _dafny.Sequence) D4 {
	return D4{D4_DC16{Cf31}}
}

func (_this D4) Is_DC16() bool {
	_, ok := _this.Get_().(D4_DC16)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC17_(_dafny.Zero, false, _dafny.Zero, _dafny.Zero, (*C1)(nil))
}

func (_this D4) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D4_DC17).Cf32
}

func (_this D4) Dtor_cf33() bool {
	return _this.Get_().(D4_DC17).Cf33
}

func (_this D4) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D4_DC17).Cf34
}

func (_this D4) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D4_DC17).Cf35
}

func (_this D4) Dtor_cf36() *C1 {
	return _this.Get_().(D4_DC17).Cf36
}

func (_this D4) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D4_DC16).Cf31
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC17:
		{
			return "D4.DC17" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D4_DC16:
		{
			return "D4.DC16" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC17:
		{
			data2, ok := other.Get_().(D4_DC17)
			return ok && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36 == data2.Cf36
		}
	case D4_DC16:
		{
			data2, ok := other.Get_().(D4_DC16)
			return ok && data1.Cf31.Equals(data2.Cf31)
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

type D5_DC19 struct {
	Cf38 bool
	Cf39 bool
	Cf40 bool
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf38 bool, Cf39 bool, Cf40 bool) D5 {
	return D5{D5_DC19{Cf38, Cf39, Cf40}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

type D5_DC18 struct {
	Cf37 _dafny.Array
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf37 _dafny.Array) D5 {
	return D5{D5_DC18{Cf37}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC20 struct {
	Cf41 D5
}

func (D5_DC20) isD5() {}

func (CompanionStruct_D5_) Create_DC20_(Cf41 D5) D5 {
	return D5{D5_DC20{Cf41}}
}

func (_this D5) Is_DC20() bool {
	_, ok := _this.Get_().(D5_DC20)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC19_(false, false, false)
}

func (_this D5) Dtor_cf38() bool {
	return _this.Get_().(D5_DC19).Cf38
}

func (_this D5) Dtor_cf39() bool {
	return _this.Get_().(D5_DC19).Cf39
}

func (_this D5) Dtor_cf40() bool {
	return _this.Get_().(D5_DC19).Cf40
}

func (_this D5) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D5_DC18).Cf37
}

func (_this D5) Dtor_cf41() D5 {
	return _this.Get_().(D5_DC20).Cf41
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D5_DC20:
		{
			return "D5.DC20" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf37 == data2.Cf37
		}
	case D5_DC20:
		{
			data2, ok := other.Get_().(D5_DC20)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D6_DC22 struct {
	Cf43 bool
}

func (D6_DC22) isD6() {}

func (CompanionStruct_D6_) Create_DC22_(Cf43 bool) D6 {
	return D6{D6_DC22{Cf43}}
}

func (_this D6) Is_DC22() bool {
	_, ok := _this.Get_().(D6_DC22)
	return ok
}

type D6_DC23 struct {
	Cf44 _dafny.Int
	Cf45 _dafny.Sequence
	Cf46 _dafny.Int
}

func (D6_DC23) isD6() {}

func (CompanionStruct_D6_) Create_DC23_(Cf44 _dafny.Int, Cf45 _dafny.Sequence, Cf46 _dafny.Int) D6 {
	return D6{D6_DC23{Cf44, Cf45, Cf46}}
}

func (_this D6) Is_DC23() bool {
	_, ok := _this.Get_().(D6_DC23)
	return ok
}

type D6_DC21 struct {
	Cf42 _dafny.Set
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf42 _dafny.Set) D6 {
	return D6{D6_DC21{Cf42}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

type D6_DC24 struct {
	Cf47 D6
}

func (D6_DC24) isD6() {}

func (CompanionStruct_D6_) Create_DC24_(Cf47 D6) D6 {
	return D6{D6_DC24{Cf47}}
}

func (_this D6) Is_DC24() bool {
	_, ok := _this.Get_().(D6_DC24)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC22_(false)
}

func (_this D6) Dtor_cf43() bool {
	return _this.Get_().(D6_DC22).Cf43
}

func (_this D6) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D6_DC23).Cf44
}

func (_this D6) Dtor_cf45() _dafny.Sequence {
	return _this.Get_().(D6_DC23).Cf45
}

func (_this D6) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D6_DC23).Cf46
}

func (_this D6) Dtor_cf42() _dafny.Set {
	return _this.Get_().(D6_DC21).Cf42
}

func (_this D6) Dtor_cf47() D6 {
	return _this.Get_().(D6_DC24).Cf47
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC22:
		{
			return "D6.DC22" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D6_DC23:
		{
			return "D6.DC23" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D6_DC21:
		{
			return "D6.DC21" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D6_DC24:
		{
			return "D6.DC24" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC22:
		{
			data2, ok := other.Get_().(D6_DC22)
			return ok && data1.Cf43 == data2.Cf43
		}
	case D6_DC23:
		{
			data2, ok := other.Get_().(D6_DC23)
			return ok && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45.Equals(data2.Cf45) && data1.Cf46.Cmp(data2.Cf46) == 0
		}
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D6_DC24:
		{
			data2, ok := other.Get_().(D6_DC24)
			return ok && data1.Cf47.Equals(data2.Cf47)
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
	F0  _dafny.Int
	F1  _dafny.Int
	F5  bool
	_f2 bool
	_f3 _dafny.Int
	_f4 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F1 = _dafny.Zero
	_this.F5 = false
	_this._f2 = false
	_this._f3 = _dafny.Zero
	_this._f4 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 bool, f3 _dafny.Int, f4 _dafny.Int, f5 bool) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
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
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F8  _dafny.Map
	_f9 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F8 = _dafny.EmptyMap
	_this._f9 = false
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

func (_this *C0) Ctor__(f8 _dafny.Map, f9 bool) {
	{
		(_this).F8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C0) Fm3(p0 _dafny.CodePoint, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus(_dafny.IntOfInt64(-152))).Minus((_dafny.IntOfInt64(851)).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.IntOfInt64(940))).Cardinality()))
	}
}
func (_this *C0) Fm4(globalState *GlobalState) bool {
	{
		return ((_dafny.MultiSetOf(_dafny.SeqOf(true), _dafny.SeqOf((_this).F9()), _dafny.SeqOf((_this).F9(), (_this).F9()))).Union(_dafny.MultiSetOf(_dafny.SeqOf(false, (_this).F9()), _dafny.SeqOf((_this).F9())))).IsDisjointFrom((_dafny.MultiSetOf(_dafny.SeqOf((_this).F9()), _dafny.SeqOf((_this).F9()))).Difference(_dafny.MultiSetOf(_dafny.SeqOf((_this).F9()), _dafny.SeqOf((_this).F9()), _dafny.SeqOf((_this).F9(), false, (_this).F9()))))
	}
}
func (_this *C0) F9() bool {
	{
		return _this._f9
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f6 _dafny.CodePoint
	_f7 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f6 = _dafny.CodePoint('D')
	_this._f7 = false
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

func (_this *C1) Ctor__(f6 _dafny.CodePoint, f7 bool) {
	{
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C1) Fm0(p0 _dafny.Sequence, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-188))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_254_i0 _dafny.Int) _dafny.CodePoint {
			return (_this).F6()
		})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qkbbvoi"), _dafny.UnicodeSeqOfUtf8Bytes("dcitg")))
	}
}
func (_this *C1) M0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		if (_this).F7() {
			var _255_v0 _dafny.Array
			_ = _255_v0
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_5
			var _nw34 _dafny.Array
			_ = _nw34
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw34 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Map = (func(_256_p1 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_257_i0 _dafny.Int) _dafny.Map {
						return (func() _dafny.Map {
							if (_this).F7() {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_256_p1, _dafny.SeqOf((_this).F7()))
							}
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_256_p1, _dafny.SeqOf((_this).F7(), (_this).F7()))
						})()
					}
				})(p1)
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
			_255_v0 = _nw34
			var _258_v1 _dafny.Sequence
			_ = _258_v1
			_258_v1 = _dafny.SeqOf((_this).F7())
			var _259_v2 _dafny.Map
			_ = _259_v2
			_259_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _258_v1)
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_255_v0), 0))
			_ = _index58
			(_255_v0).ArraySet1(_259_v2, (_index58).Int())
			var _260_v3 _dafny.Array
			_ = _260_v3
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw35
			_260_v3 = _nw35
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_260_v3), 0))
			_ = _index59
			(_260_v3).ArraySet1(p1, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_255_v0), 0))
			_ = _index60
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_260_v3), 0))
			_ = _index61
			var _rhs50 _dafny.Int = Companion_Default___.SafeModuloInt(p1, _dafny.IntOfUint32((_258_v1).Cardinality()))
			_ = _rhs50
			var _rhs51 _dafny.Map = (func() _dafny.Map {
				if Companion_Default___.Fm1(globalState) {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _258_v1)
				}
				return (_259_v2).Merge(_259_v2)
			})()
			_ = _rhs51
			var _rhs52 _dafny.Int = _dafny.IntOfInt64(48)
			_ = _rhs52
			var _rhs53 _dafny.Int = _dafny.IntOfInt64(682)
			_ = _rhs53
			var _lhs48 *GlobalState = globalState
			_ = _lhs48
			var _lhs49 _dafny.Array = _255_v0
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_255_v0), 0))
			_ = _lhs50
			var _lhs51 _dafny.Array = _260_v3
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_260_v3), 0))
			_ = _lhs52
			_lhs48.F1 = _rhs50
			(_lhs49).ArraySet1(_rhs51, (_lhs50).Int())
			r1 = _rhs52
			(_lhs51).ArraySet1(_rhs53, (_lhs52).Int())
			(globalState).F5 = (func() bool {
				if false {
					return (_this).F7()
				}
				return (_this).F7()
			})()
			var _261_v4 _dafny.Sequence
			_ = _261_v4
			_261_v4 = _dafny.UnicodeSeqOfUtf8Bytes("cfkjko")
			var _262_v5 _dafny.Map
			_ = _262_v5
			_262_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_261_v4, (_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int))
			_262_v5 = (_262_v5).Update(_261_v4, (_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int))
			var _263_v6 _dafny.Map
			_ = _263_v6
			_263_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), p1)
			var _264_v7 _dafny.Map
			_ = _264_v7
			_264_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_261_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7()))
			var _265_v8 _dafny.Map
			_ = _265_v8
			_265_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F7()), (_this).F7())
			(globalState).F1 = (func() _dafny.Int {
				if !(_263_v6).Contains((_this).F7()) {
					return Companion_Default___.Fm2(Companion_Default___.Fm1(globalState), globalState)
				}
				return (func() _dafny.Int {
					if (_this).F7() {
						return ((func() _dafny.Map {
							if (_264_v7).Contains(_261_v4) {
								return (_264_v7).Get(_261_v4).(_dafny.Map)
							}
							return _265_v8
						})()).Cardinality()
					}
					return (_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int)
				})()
			})()
			var _266_v9 _dafny.Map
			_ = _266_v9
			_266_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), ((_this).F7()) || ((_this).F7()))
			_266_v9 = (_266_v9).Update((_this).F6(), false)
		} else {
			var _267_v10 _dafny.Map
			_ = _267_v10
			_267_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_this).F7(), globalState), (_this).F7())
			var _268_v11 _dafny.MultiSet
			_ = _268_v11
			_268_v11 = _dafny.MultiSetOf((_this).F7())
			var _269_v12 *C0
			_ = _269_v12
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__(Companion_Default___.Fm5((_267_v10).Cardinality(), (_this).F7(), (_this).F7(), globalState), (_dafny.MultiSetOf((_this).F7(), (_this).F7())).Equals(_268_v11))
			_269_v12 = _nw36
			var _270_v13 _dafny.Array
			_ = _270_v13
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw37
			_270_v13 = _nw37
			var _271_v14 D0
			_ = _271_v14
			_271_v14 = Companion_D0_.Create_DC0_(_270_v13)
			_270_v13 = (_271_v14).Dtor_cf0()
			var _272_v15 _dafny.Map
			_ = _272_v15
			_272_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_269_v12).F9(), p0)
			_272_v15 = (_272_v15).Update((_269_v12).F9(), Companion_Default___.Fm2((_this).F7(), globalState))
			var _273_v16 _dafny.CodePoint
			_ = _273_v16
			_273_v16 = _dafny.CodePoint('q')
			_273_v16 = (func() _dafny.CodePoint {
				if !((_269_v12).F9()) {
					return (_this).F6()
				}
				return (_this).F6()
			})()
			var _274_v17 _dafny.Array
			_ = _274_v17
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(18))
			_ = _nw38
			_274_v17 = _nw38
			var _275_v18 _dafny.Set
			_ = _275_v18
			_275_v18 = _dafny.SetOf(_270_v13, _270_v13)
			var _276_v19 _dafny.Map
			_ = _276_v19
			_276_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_275_v18, false)
			_274_v17 = (func() _dafny.Array {
				if (func() bool {
					if (_269_v12).F9() {
						return (func() bool {
							if (_276_v19).Contains(_275_v18) {
								return (_276_v19).Get(_275_v18).(bool)
							}
							return (_269_v12).Fm4(globalState)
						})()
					}
					return (_this).F7()
				})() {
					return _274_v17
				}
				return _274_v17
			})()
		}
		var _277_v20 _dafny.Sequence
		_ = _277_v20
		_277_v20 = _dafny.UnicodeSeqOfUtf8Bytes("swfjwfa")
		var _278_v21 D1
		_ = _278_v21
		_278_v21 = Companion_D1_.Create_DC4_(false)
		var _279_v22 _dafny.Set
		_ = _279_v22
		_279_v22 = _dafny.SetOf(p0, _dafny.IntOfUint32((_277_v20).Cardinality()), p1)
		var _280_v23 _dafny.MultiSet
		_ = _280_v23
		_280_v23 = _dafny.MultiSetOf((_279_v22).Cardinality())
		var _281_v24 _dafny.Array
		_ = _281_v24
		var _nwElement0_9 bool = (_this).F7()
		_ = _nwElement0_9
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(15))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_9, 0)
		(_nw39).ArraySet1(!((_this).F7()), 1)
		(_nw39).ArraySet1((p1).Cmp(p0) > 0, 2)
		(_nw39).ArraySet1(!(_dafny.Companion_Sequence_.Equal(_277_v20, _277_v20)), 3)
		(_nw39).ArraySet1(!((_278_v21).Dtor_cf10()), 4)
		(_nw39).ArraySet1((_this).F7(), 5)
		(_nw39).ArraySet1(!(Companion_Default___.Fm1(globalState)), 6)
		(_nw39).ArraySet1((_this).F7(), 7)
		(_nw39).ArraySet1(!((_this).F7()), 8)
		(_nw39).ArraySet1((_280_v23).Equals(_280_v23), 9)
		(_nw39).ArraySet1(Companion_Default___.Fm1(globalState), 10)
		(_nw39).ArraySet1((_this).F7(), 11)
		(_nw39).ArraySet1((_this).F7(), 12)
		(_nw39).ArraySet1((_this).F7(), 13)
		(_nw39).ArraySet1((_this).F7(), 14)
		_281_v24 = _nw39
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_281_v24), 0))
		_ = _index62
		(_281_v24).ArraySet1((p0).Cmp(Companion_Default___.Fm2((_this).F7(), globalState)) < 0, (_index62).Int())
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_281_v24), 0))
		_ = _index63
		(_281_v24).ArraySet1((_this).F7(), (_index63).Int())
		(globalState).F5 = (_this).F7()
		var _282_v25 _dafny.Map
		_ = _282_v25
		_282_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-673), p1)
		var _283_v26 _dafny.Set
		_ = _283_v26
		_283_v26 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_277_v20, _277_v20), (Companion_Default___.SafeIndex((_282_v25).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_277_v20, _277_v20)).Cardinality()))).Uint32(), _dafny.CodePoint('h')))
		r1 = (_283_v26).Cardinality()
		var _284_v27 _dafny.Set
		_ = _284_v27
		_284_v27 = _dafny.SetOf((_this).F7(), (_this).F7(), (_this).F7())
		var _285_v28 _dafny.Array
		_ = _285_v28
		var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
		_ = _nw40
		_285_v28 = _nw40
		var _286_v29 D0
		_ = _286_v29
		_286_v29 = Companion_D0_.Create_DC1_(p1, _284_v27, p0, _285_v28, (_this).F6())
		var _287_v30 _dafny.Array
		_ = _287_v30
		var _nwElement0_10 D0 = _286_v29
		_ = _nwElement0_10
		var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(19))
		_ = _nw41
		(_nw41).ArraySet1(_nwElement0_10, 0)
		(_nw41).ArraySet1(_286_v29, 1)
		(_nw41).ArraySet1(_286_v29, 2)
		(_nw41).ArraySet1(_286_v29, 3)
		(_nw41).ArraySet1(_286_v29, 4)
		(_nw41).ArraySet1(_286_v29, 5)
		(_nw41).ArraySet1(_286_v29, 6)
		(_nw41).ArraySet1(_286_v29, 7)
		(_nw41).ArraySet1(_286_v29, 8)
		(_nw41).ArraySet1(_286_v29, 9)
		(_nw41).ArraySet1(_286_v29, 10)
		(_nw41).ArraySet1(_286_v29, 11)
		(_nw41).ArraySet1(_286_v29, 12)
		(_nw41).ArraySet1(_286_v29, 13)
		(_nw41).ArraySet1(_286_v29, 14)
		(_nw41).ArraySet1(_286_v29, 15)
		(_nw41).ArraySet1(_286_v29, 16)
		(_nw41).ArraySet1(_286_v29, 17)
		(_nw41).ArraySet1(_286_v29, 18)
		_287_v30 = _nw41
		var _288_v31 _dafny.Map
		_ = _288_v31
		_288_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_287_v30, (_this).F7())
		_288_v31 = (_288_v31).Update(_287_v30, (_281_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_281_v24), 0))).Int()).(bool))
		(globalState).F5 = (_281_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_281_v24), 0))).Int()).(bool)
		var _289_v32 _dafny.Map
		_ = _289_v32
		_289_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(192))
		r0 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_281_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_281_v24), 0))).Int()).(bool) {
				return (_dafny.SetOf(true)).Cardinality()
			}
			return p1
		})(), (_289_v32).Cardinality())
		var _290_v33 *C0
		_ = _290_v33
		var _nw42 *C0 = New_C0_()
		_ = _nw42
		_nw42.Ctor__(_282_v25, !((_this).F7()))
		_290_v33 = _nw42
		var _291_v34 _dafny.Array
		_ = _291_v34
		var _nwElement0_11 *C0 = _290_v33
		_ = _nwElement0_11
		var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(14))
		_ = _nw43
		(_nw43).ArraySet1(_nwElement0_11, 0)
		(_nw43).ArraySet1(_290_v33, 1)
		(_nw43).ArraySet1(_290_v33, 2)
		(_nw43).ArraySet1(_290_v33, 3)
		(_nw43).ArraySet1(_290_v33, 4)
		(_nw43).ArraySet1(_290_v33, 5)
		(_nw43).ArraySet1(_290_v33, 6)
		(_nw43).ArraySet1(_290_v33, 7)
		(_nw43).ArraySet1(_290_v33, 8)
		(_nw43).ArraySet1(_290_v33, 9)
		(_nw43).ArraySet1(_290_v33, 10)
		(_nw43).ArraySet1(_290_v33, 11)
		(_nw43).ArraySet1(_290_v33, 12)
		(_nw43).ArraySet1(_290_v33, 13)
		_291_v34 = _nw43
		var _292_v35 _dafny.Map
		_ = _292_v35
		_292_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
		r1 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if (_this).F7() {
				return _291_v34
			}
			return _291_v34
		})(), (func() bool {
			if (func() bool {
				if (_292_v35).Contains((_this).F7()) {
					return (_292_v35).Get((_this).F7()).(bool)
				}
				return (_this).F7()
			})() {
				return (_this).F7()
			}
			return true
		})())).Cardinality()
		var _293_v36 _dafny.Array
		_ = _293_v36
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw44
		_293_v36 = _nw44
		r2 = _293_v36
		return r0, r1, r2
	}
}
func (_this *C1) F6() _dafny.CodePoint {
	{
		return _this._f6
	}
}
func (_this *C1) F7() bool {
	{
		return _this._f7
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
