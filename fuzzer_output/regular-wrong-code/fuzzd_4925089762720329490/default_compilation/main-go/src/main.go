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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Map, globalState *GlobalState) bool {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-229)), !(true))).Contains((_dafny.MultiSetOf(_dafny.IntOfInt64(49))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-112))))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (Companion_Default___.SafeDivisionInt((_dafny.SetOf(true)).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-576)))).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(604), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(471), true))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("byj")).Cardinality()), _dafny.IntOfInt64(-962))).Cardinality(), _dafny.IntOfInt64(-999)), _dafny.SeqOf(_dafny.IntOfInt64(227))), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(599))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('w')
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jyepokid"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gl"), _dafny.UnicodeSeqOfUtf8Bytes("djdolp")))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) D3 {
	var _source0 D0 = (func() D0 {
		if true {
			return Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(822), _dafny.IntOfInt64(444)), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))))
		}
		return Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(634))))
	})()
	_ = _source0
	if _source0.Is_DC1() {
		var _0___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _0___mcc_h0
		var _1_cf1 _dafny.Int = _0___mcc_h0
		_ = _1_cf1
		if false {
			return Companion_D3_.Create_DC11_(_dafny.CodePoint('q'), !(true), !(true))
		} else {
			return Companion_D3_.Create_DC11_(_dafny.CodePoint('m'), !(true), true)
		}
	} else if _source0.Is_DC2() {
		var _2___mcc_h1 bool = _source0.Get_().(D0_DC2).Cf2
		_ = _2___mcc_h1
		var _3___mcc_h2 _dafny.Set = _source0.Get_().(D0_DC2).Cf3
		_ = _3___mcc_h2
		var _4___mcc_h3 bool = _source0.Get_().(D0_DC2).Cf4
		_ = _4___mcc_h3
		var _5___mcc_h4 bool = _source0.Get_().(D0_DC2).Cf5
		_ = _5___mcc_h4
		var _6___mcc_h5 _dafny.Int = _source0.Get_().(D0_DC2).Cf6
		_ = _6___mcc_h5
		var _7_cf6 _dafny.Int = _6___mcc_h5
		_ = _7_cf6
		var _8_cf5 bool = _5___mcc_h4
		_ = _8_cf5
		var _9_cf4 bool = _4___mcc_h3
		_ = _9_cf4
		var _10_cf3 _dafny.Set = _3___mcc_h2
		_ = _10_cf3
		var _11_cf2 bool = _2___mcc_h1
		_ = _11_cf2
		return Companion_D3_.Create_DC11_(_dafny.CodePoint('a'), _9_cf4, _11_cf2)
	} else if _source0.Is_DC3() {
		var _12___mcc_h6 _dafny.Sequence = _source0.Get_().(D0_DC3).Cf7
		_ = _12___mcc_h6
		var _13_cf7 _dafny.Sequence = _12___mcc_h6
		_ = _13_cf7
		return Companion_D3_.Create_DC11_(_dafny.CodePoint('f'), false, true)
	} else {
		var _14___mcc_h7 _dafny.Sequence = _source0.Get_().(D0_DC0).Cf0
		_ = _14___mcc_h7
		var _15_cf0 _dafny.Sequence = _14___mcc_h7
		_ = _15_cf0
		return Companion_D3_.Create_DC11_(_dafny.CodePoint('t'), false, true)
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 D0, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(312))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_16_i0 _dafny.Int) _dafny.Int {
		return (func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((func() _dafny.Map {
				var _coll1 = _dafny.NewMapBuilder()
				_ = _coll1
				for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-198), _dafny.IntOfInt64(-371))); ; {
					_compr_1, _ok1 := _iter1()
					if !_ok1 {
						break
					}
					var _17_v1 _dafny.Int
					_17_v1 = interface{}(_compr_1).(_dafny.Int)
					if ((_dafny.IntOfInt64(-198)).Cmp(_17_v1) <= 0) && ((_17_v1).Cmp(_dafny.IntOfInt64(-371)) < 0) {
						_coll1.Add((_17_v1).Plus((_dafny.MultiSetOf(_dafny.IntOfInt64(-110))).Cardinality()), !(false))
					}
				}
				return _coll1.ToMap()
			}()).Keys().Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _18_v0 _dafny.Int
				_18_v0 = interface{}(_compr_0).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll2 = _dafny.NewMapBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-198), _dafny.IntOfInt64(-371))); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _17_v1 _dafny.Int
						_17_v1 = interface{}(_compr_2).(_dafny.Int)
						if ((_dafny.IntOfInt64(-198)).Cmp(_17_v1) <= 0) && ((_17_v1).Cmp(_dafny.IntOfInt64(-371)) < 0) {
							_coll2.Add((_17_v1).Plus((_dafny.MultiSetOf(_dafny.IntOfInt64(-110))).Cardinality()), !(false))
						}
					}
					return _coll2.ToMap()
				}()).Contains(_18_v0) {
					_coll0.Add((_18_v0).Minus(_dafny.IntOfInt64(873)), true)
				}
			}
			return _coll0.ToMap()
		}()).Cardinality()
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-962))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_19_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(987))).Cardinality())
	}))), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-915))))
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(302), _dafny.IntOfInt64(656))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _20_v0 _dafny.Int
			_20_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(302)).Cmp(_20_v0) <= 0) && ((_20_v0).Cmp(_dafny.IntOfInt64(656)) < 0) {
				_coll3.Add((_20_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bkdjgk")).Cardinality())), (_dafny.MultiSetOf((Companion_D2_.Create_DC6_(false)).Dtor_cf10())).Cardinality())
			}
		}
		return _coll3.ToMap()
	}()), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false), !(false))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(587)), _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-336))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_21_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(921)
	})))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(531), _dafny.IntOfInt64(263)), _dafny.MultiSetOf(_dafny.IntOfInt64(245), _dafny.IntOfInt64(950)))))
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-122)))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(178))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality()))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetFromSeq((Companion_D7_.Create_DC19_(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(151))).Cardinality()), _dafny.CodePoint('v'))).Cardinality(), true)).Cardinality()))).Dtor_cf34())).Cardinality()))), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vqbthl")).Cardinality()))).Union((Companion_D4_.Create_DC14_(_dafny.MultiSetOf(_dafny.IntOfInt64(119), _dafny.IntOfInt64(-456)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ty")).Cardinality()), Companion_D2_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("ehtttxx")))).Dtor_cf28()), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lotvixcay")).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(508), _dafny.IntOfInt64(-985)))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) D2 {
	var _source1 D7 = Companion_D7_.Create_DC21_(true, true)
	_ = _source1
	if _source1.Is_DC20() {
		var _23___mcc_h0 bool = _source1.Get_().(D7_DC20).Cf35
		_ = _23___mcc_h0
		var _24___mcc_h1 _dafny.Sequence = _source1.Get_().(D7_DC20).Cf36
		_ = _24___mcc_h1
		var _25___mcc_h2 _dafny.Int = _source1.Get_().(D7_DC20).Cf37
		_ = _25___mcc_h2
		var _26___mcc_h3 bool = _source1.Get_().(D7_DC20).Cf38
		_ = _26___mcc_h3
		var _27___mcc_h4 _dafny.Int = _source1.Get_().(D7_DC20).Cf39
		_ = _27___mcc_h4
		var _28_cf39 _dafny.Int = _27___mcc_h4
		_ = _28_cf39
		var _29_cf38 bool = _26___mcc_h3
		_ = _29_cf38
		var _30_cf37 _dafny.Int = _25___mcc_h2
		_ = _30_cf37
		var _31_cf36 _dafny.Sequence = _24___mcc_h1
		_ = _31_cf36
		var _32_cf35 bool = _23___mcc_h0
		_ = _32_cf35
		return Companion_D2_.Create_DC6_(_32_cf35)
	} else if _source1.Is_DC21() {
		var _33___mcc_h5 bool = _source1.Get_().(D7_DC21).Cf40
		_ = _33___mcc_h5
		var _34___mcc_h6 bool = _source1.Get_().(D7_DC21).Cf41
		_ = _34___mcc_h6
		var _35_cf41 bool = _34___mcc_h6
		_ = _35_cf41
		var _36_cf40 bool = _33___mcc_h5
		_ = _36_cf40
		if _35_cf41 {
			return Companion_D2_.Create_DC6_(false)
		} else {
			return Companion_D2_.Create_DC6_(_36_cf40)
		}
	} else {
		var _37___mcc_h7 _dafny.Sequence = _source1.Get_().(D7_DC19).Cf34
		_ = _37___mcc_h7
		var _38_cf34 _dafny.Sequence = _37___mcc_h7
		_ = _38_cf34
		return Companion_D2_.Create_DC6_(false)
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(897), false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.IntOfInt64(-664)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _dafny.IntOfInt64(441)))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(426))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_39_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(500)
	}))))), Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(410), _dafny.IntOfInt64(803))))), (func() D0 {
		if false {
			return Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))))
		}
		return Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(723))).Cardinality())))
	})())
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.MultiSetOf(false))).Cardinality(), _dafny.IntOfInt64(-497))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(130), _dafny.IntOfInt64(-588))
	})()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(811))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_40_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))).Cardinality()), (_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(false, true, false)
}
func (_static *CompanionStruct_Default___) M1(p0 bool, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 bool = false
	_ = r1
	var _41_v0 _dafny.Map
	_ = _41_v0
	_41_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(p1)).Cmp(p1) > 0, p0)
	_41_v0 = (_41_v0).Update(p0, p0)
	var _42_v1 _dafny.Sequence
	_ = _42_v1
	_42_v1 = _dafny.SeqOf(_dafny.SeqOf(p0))
	var _43_v2 D0
	_ = _43_v2
	_43_v2 = Companion_D0_.Create_DC3_((_42_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.IntOfUint32((_42_v1).Cardinality()))).Uint32()).(_dafny.Sequence))
	var _rhs0 D0 = (func() D0 {
		if p0 {
			return _43_v2
		}
		return (func() D0 {
			if p0 {
				return _43_v2
			}
			return _43_v2
		})()
	})()
	_ = _rhs0
	var _rhs1 _dafny.Int = (p1).Times(p1)
	_ = _rhs1
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	_43_v2 = _rhs0
	_lhs0.F2 = _rhs1
	var _44_v3 _dafny.Array
	_ = _44_v3
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
	_ = _nw0
	_44_v3 = _nw0
	var _45_v4 _dafny.Set
	_ = _45_v4
	_45_v4 = _dafny.SetOf(_44_v3)
	var _46_v5 *C0
	_ = _46_v5
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__()
	_46_v5 = _nw1
	var _47_v6 _dafny.Sequence
	_ = _47_v6
	_47_v6 = _dafny.SeqOf(_46_v5)
	var _48_v7 _dafny.CodePoint
	_ = _48_v7
	_48_v7 = _dafny.CodePoint('o')
	var _49_v8 _dafny.MultiSet
	_ = _49_v8
	_49_v8 = _dafny.MultiSetOf(_48_v7, _48_v7, _48_v7)
	var _50_v9 _dafny.Map
	_ = _50_v9
	_50_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v8, p1)
	var _51_v10 D0
	_ = _51_v10
	_51_v10 = Companion_D0_.Create_DC2_(p0, _45_v4, !(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_47_v6, _47_v6))), p0, (_50_v9).Cardinality())
	var _source2 D0 = _51_v10
	_ = _source2
	if _source2.Is_DC1() {
		var _52___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
		_ = _52___mcc_h0
		var _53_cf1 _dafny.Int = _52___mcc_h0
		_ = _53_cf1
		var _54_v11 _dafny.Array
		_ = _54_v11
		var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
		_ = _nw2
		_54_v11 = _nw2
		var _55_v12 D6
		_ = _55_v12
		_55_v12 = Companion_D6_.Create_DC16_(_44_v3)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_54_v11), 0))
		_ = _index0
		(_54_v11).ArraySet1((_55_v12).Dtor_cf32(), (_index0).Int())
		var _56_v13 _dafny.Sequence
		_ = _56_v13
		_56_v13 = _dafny.SeqOf(_44_v3)
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_54_v11), 0))
		_ = _index1
		(_54_v11).ArraySet1((func() _dafny.Array {
			if p0 {
				return _44_v3
			}
			return (_56_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.IntOfUint32((_56_v13).Cardinality()))).Uint32()).(_dafny.Array)
		})(), (_index1).Int())
		var _57_v14 _dafny.Array
		_ = _57_v14
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_0
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_58_p0 bool) func(_dafny.Int) bool {
				return func(_59_i0 _dafny.Int) bool {
					return !(_58_p0)
				}
			})(p0)
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
		_57_v14 = _nw3
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))
		_ = _index2
		(_57_v14).ArraySet1(p0, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))
		_ = _index3
		(_57_v14).ArraySet1(p0, (_index3).Int())
		if (_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool) {
			var _60_v15 *C0
			_ = _60_v15
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__()
			_60_v15 = _nw4
			(globalState).F8 = _53_cf1
			var _61_v16 _dafny.Set
			_ = _61_v16
			_61_v16 = _dafny.SetOf((_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool))
			var _62_v17 *C1
			_ = _62_v17
			var _nw5 *C1 = New_C1_()
			_ = _nw5
			_nw5.Ctor__(_61_v16, _53_cf1)
			_62_v17 = _nw5
			var _63_v18 _dafny.Array
			_ = _63_v18
			var _nwElement0_0 *C1 = _62_v17
			_ = _nwElement0_0
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(15))
			_ = _nw6
			(_nw6).ArraySet1(_nwElement0_0, 0)
			(_nw6).ArraySet1(_62_v17, 1)
			(_nw6).ArraySet1(_62_v17, 2)
			(_nw6).ArraySet1(_62_v17, 3)
			(_nw6).ArraySet1(_62_v17, 4)
			(_nw6).ArraySet1(_62_v17, 5)
			(_nw6).ArraySet1(_62_v17, 6)
			(_nw6).ArraySet1(_62_v17, 7)
			(_nw6).ArraySet1(_62_v17, 8)
			(_nw6).ArraySet1(_62_v17, 9)
			(_nw6).ArraySet1(_62_v17, 10)
			(_nw6).ArraySet1(_62_v17, 11)
			(_nw6).ArraySet1(_62_v17, 12)
			(_nw6).ArraySet1(_62_v17, 13)
			(_nw6).ArraySet1(_62_v17, 14)
			_63_v18 = _nw6
			var _64_v19 _dafny.Array
			_ = _64_v19
			var _nwElement0_1 _dafny.Array = _63_v18
			_ = _nwElement0_1
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(8))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_1, 0)
			(_nw7).ArraySet1(_63_v18, 1)
			(_nw7).ArraySet1(_63_v18, 2)
			(_nw7).ArraySet1(_63_v18, 3)
			(_nw7).ArraySet1(_63_v18, 4)
			(_nw7).ArraySet1(_63_v18, 5)
			(_nw7).ArraySet1(_63_v18, 6)
			(_nw7).ArraySet1(_63_v18, 7)
			_64_v19 = _nw7
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_64_v19), 0))
			_ = _index4
			(_64_v19).ArraySet1(_63_v18, (_index4).Int())
			var _65_v20 _dafny.Sequence
			_ = _65_v20
			_65_v20 = _dafny.SeqOf(p0)
			var _66_v21 _dafny.Set
			_ = _66_v21
			_66_v21 = _dafny.SetOf(_65_v20)
			var _67_v22 _dafny.Map
			_ = _67_v22
			_67_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _63_v18)
			var _68_v23 _dafny.Sequence
			_ = _68_v23
			_68_v23 = _dafny.SeqOf(_63_v18, (func() _dafny.Array {
				if (_67_v22).Contains((_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool)) {
					return (_67_v22).Get((_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _63_v18
			})())
			var _69_v24 _dafny.MultiSet
			_ = _69_v24
			_69_v24 = _dafny.MultiSetOf(_62_v17.F27)
			var _70_v25 _dafny.Sequence
			_ = _70_v25
			_70_v25 = _dafny.SeqOf((_68_v23).Select((Companion_Default___.SafeIndex((Companion_D3_.Create_DC10_(p0, (_69_v24).Cardinality(), p0)).Dtor_cf17(), _dafny.IntOfUint32((_68_v23).Cardinality()))).Uint32()).(_dafny.Array), _63_v18, _63_v18)
			var _71_v26 T0
			_ = _71_v26
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__()
			_71_v26 = _nw8
			var _72_v27 _dafny.Sequence
			_ = _72_v27
			_72_v27 = _dafny.SeqOf(_71_v26)
			var _73_v28 _dafny.MultiSet
			_ = _73_v28
			_73_v28 = _dafny.MultiSetOf((_72_v27).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_72_v27).Cardinality()))).Uint32()).(T0), _71_v26)
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))
			_ = _index5
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_64_v19), 0))
			_ = _index6
			var _rhs2 bool = (_66_v21).IsProperSubsetOf(_dafny.SetOf(_65_v20))
			_ = _rhs2
			var _rhs3 bool = (_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool)
			_ = _rhs3
			var _rhs4 _dafny.Array = (_70_v25).Select((Companion_Default___.SafeIndex((_73_v28).Cardinality(), _dafny.IntOfUint32((_70_v25).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs4
			var _lhs1 *GlobalState = globalState
			_ = _lhs1
			var _lhs2 _dafny.Array = _57_v14
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))
			_ = _lhs3
			var _lhs4 _dafny.Array = _64_v19
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_64_v19), 0))
			_ = _lhs5
			_lhs1.F23 = _rhs2
			(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
			(_lhs4).ArraySet1(_rhs4, (_lhs5).Int())
			(globalState).F2 = p1
			(globalState).F23 = !((_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool))
		} else {
			var _74_v29 _dafny.Sequence
			_ = _74_v29
			_74_v29 = _dafny.SeqOf((_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool), (_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool))
			var _75_v30 _dafny.Sequence
			_ = _75_v30
			_75_v30 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			(globalState).F19 = Companion_Default___.Fm1((_74_v29).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_75_v30).Cardinality()), _dafny.IntOfUint32((_74_v29).Cardinality()))).Uint32()).(bool), _48_v7, _dafny.IntOfUint32((_75_v30).Cardinality()), globalState)
			var _76_v31 *C0
			_ = _76_v31
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__()
			_76_v31 = _nw9
			(globalState).F23 = (_dafny.IntOfInt64(474)).Cmp((p1).Times(_dafny.IntOfInt64(572))) >= 0
			var _77_v32 _dafny.MultiSet
			_ = _77_v32
			_77_v32 = _dafny.MultiSetOf(p1)
			var _78_v33 _dafny.Sequence
			_ = _78_v33
			_78_v33 = _dafny.SeqOf(_77_v32, _77_v32)
			var _79_v34 D0
			_ = _79_v34
			_79_v34 = Companion_D0_.Create_DC0_(_78_v33)
			var _80_v35 _dafny.MultiSet
			_ = _80_v35
			_80_v35 = _dafny.MultiSetOf(_79_v34, _79_v34)
			var _81_v36 _dafny.MultiSet
			_ = _81_v36
			_81_v36 = _80_v35
			var _82_v37 _dafny.Set
			_ = _82_v37
			_82_v37 = _dafny.SetOf((_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool))
			var _83_v38 T0
			_ = _83_v38
			var _nw10 *C1 = New_C1_()
			_ = _nw10
			_nw10.Ctor__(_82_v37, p1)
			_83_v38 = _nw10
			var _84_v39 _dafny.Map
			_ = _84_v39
			_84_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_54_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_54_v11), 0))).Int())), _dafny.IntOfUint32((_74_v29).Cardinality()))
			var _85_v40 _dafny.Map
			_ = _85_v40
			_85_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v38, (func() _dafny.Int {
				if (_84_v39).Contains(_dafny.ArrayCastTo((_54_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_54_v11), 0))).Int()))) {
					return (_84_v39).Get(_dafny.ArrayCastTo((_54_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_54_v11), 0))).Int()))).(_dafny.Int)
				}
				return _53_cf1
			})())
			var _86_v41 _dafny.Sequence
			_ = _86_v41
			_86_v41 = _dafny.SeqOf(_85_v40)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))
			_ = _index7
			var _rhs5 _dafny.MultiSet = (_80_v35).Intersection(_80_v35)
			_ = _rhs5
			var _rhs6 bool = (_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool)
			_ = _rhs6
			var _rhs7 _dafny.Int = _dafny.IntOfUint32((_86_v41).Cardinality())
			_ = _rhs7
			var _rhs8 _dafny.CodePoint = _48_v7
			_ = _rhs8
			var _lhs6 _dafny.Array = _57_v14
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			_81_v36 = _rhs5
			(_lhs6).ArraySet1(_rhs6, (_lhs7).Int())
			_lhs8.F10 = _rhs7
			_48_v7 = _rhs8
			var _87_v42 _dafny.Map
			_ = _87_v42
			_87_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() bool {
				if (_41_v0).Contains(p0) {
					return (_41_v0).Get(p0).(bool)
				}
				return p0
			})())
			var _88_v43 _dafny.Sequence
			_ = _88_v43
			_88_v43 = _dafny.SeqOf(Companion_Default___.Fm20((func() bool {
				if (_87_v42).Contains(p1) {
					return (_87_v42).Get(p1).(bool)
				}
				return (_57_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))).Int()).(bool)
			})(), globalState))
			var _89_v44 _dafny.Map
			_ = _89_v44
			_89_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_44_v3, p0)
			var _90_v45 _dafny.Map
			_ = _90_v45
			_90_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_88_v43).Select((Companion_Default___.SafeIndex((_89_v44).Cardinality(), _dafny.IntOfUint32((_88_v43).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.ArrayCastTo((_54_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_54_v11), 0))).Int())))
			(globalState).F11 = (func() _dafny.Array {
				if (_90_v45).Contains(Companion_Default___.SafeDivisionInt((_76_v31).Fm8(_82_v37, globalState), p1)) {
					return (_90_v45).Get(Companion_Default___.SafeDivisionInt((_76_v31).Fm8(_82_v37, globalState), p1)).(_dafny.Array)
				}
				return _44_v3
			})()
		}
		var _91_v46 _dafny.MultiSet
		_ = _91_v46
		_91_v46 = _dafny.MultiSetOf(_dafny.ArrayCastTo((_54_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_54_v11), 0))).Int())), _44_v3)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_57_v14), 0))
		_ = _index8
		(_57_v14).ArraySet1((_91_v46).IsSubsetOf(_91_v46), (_index8).Int())
	} else if _source2.Is_DC2() {
		var _92___mcc_h1 bool = _source2.Get_().(D0_DC2).Cf2
		_ = _92___mcc_h1
		var _93___mcc_h2 _dafny.Set = _source2.Get_().(D0_DC2).Cf3
		_ = _93___mcc_h2
		var _94___mcc_h3 bool = _source2.Get_().(D0_DC2).Cf4
		_ = _94___mcc_h3
		var _95___mcc_h4 bool = _source2.Get_().(D0_DC2).Cf5
		_ = _95___mcc_h4
		var _96___mcc_h5 _dafny.Int = _source2.Get_().(D0_DC2).Cf6
		_ = _96___mcc_h5
		var _97_cf6 _dafny.Int = _96___mcc_h5
		_ = _97_cf6
		var _98_cf5 bool = _95___mcc_h4
		_ = _98_cf5
		var _99_cf4 bool = _94___mcc_h3
		_ = _99_cf4
		var _100_cf3 _dafny.Set = _93___mcc_h2
		_ = _100_cf3
		var _101_cf2 bool = _92___mcc_h1
		_ = _101_cf2
		r1 = _99_cf4
		var _102_v47 T0
		_ = _102_v47
		var _nw11 *C1 = New_C1_()
		_ = _nw11
		_nw11.Ctor__(Companion_Default___.Fm21(_98_cf5, true, globalState), _97_cf6)
		_102_v47 = _nw11
		var _103_v48 _dafny.Map
		_ = _103_v48
		_103_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v7, _98_cf5)
		var _104_v49 _dafny.Sequence
		_ = _104_v49
		_104_v49 = _dafny.UnicodeSeqOfUtf8Bytes("lpfwbpp")
		var _105_v50 _dafny.Map
		_ = _105_v50
		_105_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v48, _104_v49)
		_102_v47 = (func() T0 {
			if !_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
				if (_105_v50).Contains(_103_v48) {
					return (_105_v50).Get(_103_v48).(_dafny.Sequence)
				}
				return _104_v49
			})(), _48_v7) {
				return _102_v47
			}
			return _102_v47
		})()
		_49_v8 = _dafny.MultiSetFromSeq(_104_v49)
		if !(_101_cf2) {
			(globalState).F19 = p1
			var _106_v51 _dafny.Set
			_ = _106_v51
			_106_v51 = _dafny.SetOf(_99_cf4, !(_98_cf5), _99_cf4, p0)
			var _107_v52 _dafny.Sequence
			_ = _107_v52
			_107_v52 = _dafny.SeqOf(_dafny.SetOf(_101_cf2), _106_v51, _106_v51)
			var _108_v53 _dafny.Map
			_ = _108_v53
			_108_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_107_v52).Cardinality()))
			var _109_v54 _dafny.Map
			_ = _109_v54
			_109_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_cf6, (_108_v53).Cardinality())
			var _110_v55 _dafny.Sequence
			_ = _110_v55
			_110_v55 = _dafny.SeqOf(_dafny.IntOfInt64(206), (func() _dafny.Int {
				if (_109_v54).Contains(p1) {
					return (_109_v54).Get(p1).(_dafny.Int)
				}
				return p1
			})())
			_110_v55 = _110_v55
			_48_v7 = _48_v7
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_44_v3), 0))
			_ = _index9
			(_44_v3).ArraySet1(_97_cf6, (_index9).Int())
			var _111_v56 _dafny.MultiSet
			_ = _111_v56
			_111_v56 = _dafny.MultiSetOf(_101_cf2)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_44_v3), 0))
			_ = _index10
			var _rhs9 _dafny.Int = _97_cf6
			_ = _rhs9
			var _rhs10 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mjpyri")).Cardinality())).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_112_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			})), _104_v49), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if (_111_v56).Contains(_98_cf5) {
					return (_111_v56).Multiplicity(_98_cf5)
				}
				return _97_cf6
			})()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_113_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			})), _104_v49)).Cardinality()))).Uint32(), _48_v7)).Cardinality()))
			_ = _rhs10
			var _rhs11 bool = true
			_ = _rhs11
			var _rhs12 _dafny.Int = (_dafny.Zero).Minus((_97_cf6).Plus((_dafny.Zero).Minus(_97_cf6)))
			_ = _rhs12
			var _lhs9 _dafny.Array = _44_v3
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_44_v3), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			(_lhs9).ArraySet1(_rhs9, (_lhs10).Int())
			_lhs11.F5 = _rhs10
			r1 = _rhs11
			_lhs12.F10 = _rhs12
			var _114_v57 *C0
			_ = _114_v57
			var _nw12 *C0 = New_C0_()
			_ = _nw12
			_nw12.Ctor__()
			_114_v57 = _nw12
		} else {
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_44_v3), 0))
			_ = _index11
			(_44_v3).ArraySet1(p1, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_44_v3), 0))
			_ = _index12
			(_44_v3).ArraySet1(p1, (_index12).Int())
			var _115_v58 _dafny.MultiSet
			_ = _115_v58
			_115_v58 = _dafny.MultiSetOf(_104_v49)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_44_v3), 0))
			_ = _index13
			(_44_v3).ArraySet1((func() _dafny.Int {
				if (_115_v58).Contains(_dafny.Companion_Sequence_.Update(_104_v49, (Companion_Default___.SafeIndex(_97_cf6, _dafny.IntOfUint32((_104_v49).Cardinality()))).Uint32(), _48_v7)) {
					return (_115_v58).Multiplicity(_dafny.Companion_Sequence_.Update(_104_v49, (Companion_Default___.SafeIndex(_97_cf6, _dafny.IntOfUint32((_104_v49).Cardinality()))).Uint32(), _48_v7))
				}
				return p1
			})(), (_index13).Int())
			var _116_v59 _dafny.Array
			_ = _116_v59
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw13
			_116_v59 = _nw13
			var _117_v60 _dafny.Map
			_ = _117_v60
			_117_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v59, p1)
			var _118_v61 _dafny.Map
			_ = _118_v61
			_118_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_cf4, _117_v60)
			(globalState).F23 = (func() bool {
				if (_41_v0).Contains((p1).Cmp(_97_cf6) < 0) {
					return (_41_v0).Get((p1).Cmp(_97_cf6) < 0).(bool)
				}
				return (_117_v60).Equals((func() _dafny.Map {
					if (_118_v61).Contains(_98_cf5) {
						return (_118_v61).Get(_98_cf5).(_dafny.Map)
					}
					return _117_v60
				})())
			})()
			var _119_v62 _dafny.Set
			_ = _119_v62
			_119_v62 = _dafny.SetOf(_98_cf5, _101_cf2, !(_98_cf5))
			var _120_v63 _dafny.MultiSet
			_ = _120_v63
			_120_v63 = _dafny.MultiSetOf(_dafny.IntOfUint32((_104_v49).Cardinality()))
			var _121_v64 _dafny.Sequence
			_ = _121_v64
			_121_v64 = _dafny.SeqOf(_120_v63)
			var _122_v65 D0
			_ = _122_v65
			_122_v65 = Companion_D0_.Create_DC0_(_121_v64)
			var _123_v66 _dafny.MultiSet
			_ = _123_v66
			_123_v66 = _dafny.MultiSetOf(_122_v65)
			var _124_v67 _dafny.MultiSet
			_ = _124_v67
			_124_v67 = _123_v66
			var _125_v68 _dafny.Map
			_ = _125_v68
			_125_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_46_v5).Fm8(_119_v62, globalState)).Times((_dafny.Zero).Minus((_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int))), _124_v67)
			_125_v68 = _125_v68
			(globalState).F22 = _99_cf4
		}
	} else if _source2.Is_DC3() {
		var _126___mcc_h6 _dafny.Sequence = _source2.Get_().(D0_DC3).Cf7
		_ = _126___mcc_h6
		var _127_cf7 _dafny.Sequence = _126___mcc_h6
		_ = _127_cf7
		var _128_v69 _dafny.MultiSet
		_ = _128_v69
		_128_v69 = _dafny.MultiSetOf(_44_v3)
		var _129_v70 _dafny.Sequence
		_ = _129_v70
		_129_v70 = _dafny.SeqOf(p1)
		var _130_v71 _dafny.Sequence
		_ = _130_v71
		_130_v71 = _dafny.SeqOf(Companion_Default___.SafeModuloInt((_128_v69).Cardinality(), (_129_v70).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_129_v70).Cardinality()))).Uint32()).(_dafny.Int)), p1)
		_129_v70 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1), _130_v71)
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_44_v3), 0))
		_ = _index14
		(_44_v3).ArraySet1(p1, (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_44_v3), 0))
		_ = _index15
		(_44_v3).ArraySet1(p1, (_index15).Int())
		var _131_v72 _dafny.MultiSet
		_ = _131_v72
		_131_v72 = _dafny.MultiSetOf(Companion_Default___.Fm0(p0, _41_v0, globalState))
		var _132_v73 D2
		_ = _132_v73
		_132_v73 = Companion_D2_.Create_DC8_(p1, p0, (func() _dafny.Int {
			if (_131_v72).Contains(p0) {
				return (_131_v72).Multiplicity(p0)
			}
			return (_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int)
		})())
		_132_v73 = _132_v73
		_46_v5 = _46_v5
	} else {
		var _133___mcc_h7 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf0
		_ = _133___mcc_h7
		var _134_cf0 _dafny.Sequence = _133___mcc_h7
		_ = _134_cf0
		(globalState).F25 = p0
		var _rhs13 bool = !(p0)
		_ = _rhs13
		var _rhs14 _dafny.Int = p1
		_ = _rhs14
		var _lhs13 *GlobalState = globalState
		_ = _lhs13
		r1 = _rhs13
		_lhs13.F2 = _rhs14
		var _135_v74 D0
		_ = _135_v74
		_135_v74 = Companion_D0_.Create_DC1_(p1)
		var _136_v75 T0
		_ = _136_v75
		var _nw14 *C0 = New_C0_()
		_ = _nw14
		_nw14.Ctor__()
		_136_v75 = _nw14
		var _137_v76 D3
		_ = _137_v76
		_137_v76 = Companion_D3_.Create_DC12_(p0, _135_v74, p0, p0, _136_v75)
		var _138_v77 _dafny.Map
		_ = _138_v77
		_138_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v76, p1)
		_138_v77 = (_138_v77).Update(_137_v76, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()))
		var _139_v78 _dafny.Sequence
		_ = _139_v78
		_139_v78 = _dafny.UnicodeSeqOfUtf8Bytes("akpsj")
		_139_v78 = _dafny.Companion_Sequence_.Concatenate(_139_v78, _139_v78)
	}
	var _hi0 _dafny.Int = p1
	_ = _hi0
	for _140_i2 := ((_41_v0).Update(p0, p0)).Cardinality(); _140_i2.Cmp(_hi0) < 0; _140_i2 = _140_i2.Plus(_dafny.One) {
		var _141_v79 D3
		_ = _141_v79
		_141_v79 = Companion_D3_.Create_DC10_(p0, _140_i2, p0)
		r1 = Companion_Default___.Fm0((_141_v79).Dtor_cf16(), _41_v0, globalState)
		if (_dafny.IntOfInt64(106)).Cmp(_140_i2) != 0 {
			var _142_v80 _dafny.MultiSet
			_ = _142_v80
			_142_v80 = _dafny.MultiSetOf(p0, p0)
			var _143_v81 _dafny.Sequence
			_ = _143_v81
			_143_v81 = _dafny.SeqOf(p0)
			var _rhs15 bool = !(!(Companion_Default___.Fm0(!(p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_Default___.Fm0(true, _41_v0, globalState)), globalState)))
			_ = _rhs15
			var _rhs16 _dafny.Int = _140_i2
			_ = _rhs16
			var _rhs17 _dafny.MultiSet = _142_v80
			_ = _rhs17
			var _rhs18 bool = (_143_v81).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_143_v81).Cardinality()))).Uint32()).(bool)
			_ = _rhs18
			var _lhs14 *GlobalState = globalState
			_ = _lhs14
			var _lhs15 *GlobalState = globalState
			_ = _lhs15
			var _lhs16 *GlobalState = globalState
			_ = _lhs16
			var _lhs17 *GlobalState = globalState
			_ = _lhs17
			_lhs14.F25 = _rhs15
			_lhs15.F19 = _rhs16
			_lhs16.F16 = _rhs17
			_lhs17.F23 = _rhs18
			var _144_v82 _dafny.Set
			_ = _144_v82
			_144_v82 = _dafny.SetOf(p0, p0, p0, p0, p0)
			var _145_v84 _dafny.Set
			_ = _145_v84
			_145_v84 = _dafny.SetOf(_140_i2, p1, _140_i2)
			var _146_v85 _dafny.Map
			_ = _146_v85
			_146_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v82, (func() _dafny.Set {
				var _coll4 = _dafny.NewBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(20), _dafny.IntOfInt64(792))); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _147_v83 _dafny.Int
					_147_v83 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(20)).Cmp(_147_v83) <= 0) && ((_147_v83).Cmp(_dafny.IntOfInt64(792)) < 0) {
						_coll4.Add((_147_v83).Plus(p1))
					}
				}
				return _coll4.ToSet()
			}()).Union(_145_v84))
			var _148_v86 T0
			_ = _148_v86
			var _nw15 *C1 = New_C1_()
			_ = _nw15
			_nw15.Ctor__(_dafny.SetOf(p0), _140_i2)
			_148_v86 = _nw15
			var _149_v87 _dafny.MultiSet
			_ = _149_v87
			_149_v87 = _dafny.MultiSetOf(_148_v86, _148_v86)
			var _rhs19 bool = (_dafny.MultiSetOf(_148_v86, _148_v86)).IsSubsetOf((_149_v87).Intersection(_149_v87))
			_ = _rhs19
			var _rhs20 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v82, (_145_v84).Difference(_145_v84))
			_ = _rhs20
			var _rhs21 bool = Companion_Default___.Fm0(Companion_Default___.Fm0(false, _41_v0, globalState), _41_v0, globalState)
			_ = _rhs21
			var _rhs22 _dafny.Int = p1
			_ = _rhs22
			var _lhs18 *GlobalState = globalState
			_ = _lhs18
			var _lhs19 *GlobalState = globalState
			_ = _lhs19
			var _lhs20 *GlobalState = globalState
			_ = _lhs20
			_lhs18.F22 = _rhs19
			_146_v85 = _rhs20
			_lhs19.F22 = _rhs21
			_lhs20.F15 = _rhs22
			var _150_v88 _dafny.Map
			_ = _150_v88
			_150_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_148_v86, (func() T0 {
				if p0 {
					return _148_v86
				}
				return _148_v86
			})())
			_150_v88 = (_150_v88).Update(_148_v86, _148_v86)
			var _151_v89 D0
			_ = _151_v89
			_151_v89 = Companion_D0_.Create_DC1_(p1)
			var _152_v90 _dafny.Sequence
			_ = _152_v90
			_152_v90 = _dafny.SeqOf(_dafny.IntOfInt64(800), _140_i2)
			var _153_v91 _dafny.Map
			_ = _153_v91
			_153_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0) && (!(Companion_Default___.Fm0(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), globalState))), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm12((_145_v84).Cardinality(), _151_v89, globalState), _152_v90)).Cardinality()))
			_153_v91 = (_153_v91).Update((func() bool {
				if !(p0) {
					return true
				}
				return Companion_Default___.Fm0(p0, _41_v0, globalState)
			})(), Companion_Default___.SafeDivisionInt((Companion_Default___.Fm21(p0, p0, globalState)).Cardinality(), p1))
			var _154_v92 _dafny.Sequence
			_ = _154_v92
			_154_v92 = _dafny.UnicodeSeqOfUtf8Bytes("ldritmxh")
			_154_v92 = _dafny.UnicodeSeqOfUtf8Bytes("oocr")
		} else {
			(globalState).F5 = (func() _dafny.Int {
				if p0 {
					return _140_i2
				}
				return _140_i2
			})()
			var _155_v93 _dafny.Map
			_ = _155_v93
			_155_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _156_v95 _dafny.Sequence
			_ = _156_v95
			_156_v95 = _dafny.UnicodeSeqOfUtf8Bytes("b")
			var _157_v96 _dafny.Map
			_ = _157_v96
			_157_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_156_v95).Cardinality()), p1)
			var _rhs23 _dafny.Map = (func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_157_v96).Keys().Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _158_v94 _dafny.Int
					_158_v94 = interface{}(_compr_5).(_dafny.Int)
					if (_157_v96).Contains(_158_v94) {
						_coll5.Add((_158_v94).Minus(p1), p0)
					}
				}
				return _coll5.ToMap()
			}()).Merge(_155_v93)
			_ = _rhs23
			var _rhs24 _dafny.CodePoint = _48_v7
			_ = _rhs24
			_155_v93 = _rhs23
			_48_v7 = _rhs24
			(globalState).F22 = p0
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))
			_ = _index16
			(_44_v3).ArraySet1((p1).Times(_dafny.IntOfInt64(695)), (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))
			_ = _index17
			(_44_v3).ArraySet1(_140_i2, (_index17).Int())
			var _159_v97 _dafny.Set
			_ = _159_v97
			_159_v97 = _dafny.SetOf(p0)
			var _160_v98 *C1
			_ = _160_v98
			var _nw16 *C1 = New_C1_()
			_ = _nw16
			_nw16.Ctor__(_159_v97, p1)
			_160_v98 = _nw16
			var _161_v99 _dafny.Map
			_ = _161_v99
			_161_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v98, _dafny.IntOfUint32((_156_v95).Cardinality()))
			var _162_v100 _dafny.Map
			_ = _162_v100
			_162_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v98, false)
			var _163_v101 _dafny.Sequence
			_ = _163_v101
			_163_v101 = _dafny.SeqOf(_140_i2, (_162_v100).Cardinality(), (_dafny.Zero).Minus(p1))
			var _164_v102 _dafny.Sequence
			_ = _164_v102
			_164_v102 = _dafny.SeqOf(false, p0)
			var _165_v103 D0
			_ = _165_v103
			_165_v103 = Companion_D0_.Create_DC1_((_dafny.MultiSetFromSeq(_164_v102)).Cardinality())
			var _166_v104 _dafny.Map
			_ = _166_v104
			_166_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v95, p0)
			var _pat_let_tv0 = p0
			_ = _pat_let_tv0
			var _pat_let_tv1 = _48_v7
			_ = _pat_let_tv1
			var _pat_let_tv2 = _44_v3
			_ = _pat_let_tv2
			var _pat_let_tv3 = _44_v3
			_ = _pat_let_tv3
			var _pat_let_tv4 = globalState
			_ = _pat_let_tv4
			var _167_v105 _dafny.Array
			_ = _167_v105
			var _nwElement0_2 _dafny.Int = (_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int)
			_ = _nwElement0_2
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(20))
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_2, 0)
			(_nw17).ArraySet1(_140_i2, 1)
			(_nw17).ArraySet1((func() _dafny.Int {
				if true {
					return _dafny.IntOfInt64(952)
				}
				return _140_i2
			})(), 2)
			(_nw17).ArraySet1((_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int), 3)
			(_nw17).ArraySet1(((_161_v99).Cardinality()).Plus((_dafny.Zero).Minus((_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int))), 4)
			(_nw17).ArraySet1(p1, 5)
			(_nw17).ArraySet1((_163_v101).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_163_v101).Cardinality()))).Uint32()).(_dafny.Int), 6)
			(_nw17).ArraySet1((_dafny.Zero).Minus(_160_v98.F27), 7)
			(_nw17).ArraySet1((_dafny.Zero).Minus(_160_v98.F27), 8)
			(_nw17).ArraySet1((_dafny.Zero).Minus(_160_v98.F27), 9)
			(_nw17).ArraySet1((_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int), 10)
			(_nw17).ArraySet1((func() _dafny.Int {
				if !(Companion_Default___.Fm0(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), globalState)) {
					return (_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int)
				}
				return (_dafny.MultiSetOf((_164_v102).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_140_i2), _dafny.IntOfUint32((_164_v102).Cardinality()))).Uint32()).(bool), p0, p0, p0)).Cardinality()
			})(), 11)
			(_nw17).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xrsul")).Cardinality()), 12)
			(_nw17).ArraySet1((_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int), 13)
			(_nw17).ArraySet1(_dafny.IntOfUint32((_156_v95).Cardinality()), 14)
			(_nw17).ArraySet1((func() _dafny.Int {
				if p0 {
					return _160_v98.F27
				}
				return Companion_Default___.Fm1(p0, _48_v7, _160_v98.F27, globalState)
			})(), 15)
			(_nw17).ArraySet1(p1, 16)
			(_nw17).ArraySet1(((func(_pat_let0_0 D0) D0 {
				return func(_168_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let1_0 _dafny.Int) D0 {
						return func(_169_dt__update_hcf1_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC1_(_169_dt__update_hcf1_h0)
						}(_pat_let1_0)
					}(Companion_Default___.Fm1(_pat_let_tv0, _pat_let_tv1, (_pat_let_tv3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_pat_let_tv2), 0))).Int()).(_dafny.Int), _pat_let_tv4))
				}(_pat_let0_0)
			}(_165_v103)).Dtor_cf1()).Plus(((_dafny.MultiSetOf(p0)).Update((func() bool {
				if (_166_v104).Contains(_156_v95) {
					return (_166_v104).Get(_156_v95).(bool)
				}
				return p0
			})(), Companion_Default___.Abs(_140_i2))).Cardinality()), 17)
			(_nw17).ArraySet1((_dafny.Zero).Minus(_140_i2), 18)
			(_nw17).ArraySet1((_44_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_44_v3), 0))).Int()).(_dafny.Int), 19)
			_167_v105 = _nw17
			(globalState).F11 = _167_v105
		}
		r0 = p1
		var _170_v106 *C0
		_ = _170_v106
		var _nw18 *C0 = New_C0_()
		_ = _nw18
		_nw18.Ctor__()
		_170_v106 = _nw18
	}
	var _171_v107 *C1
	_ = _171_v107
	var _nw19 *C1 = New_C1_()
	_ = _nw19
	_nw19.Ctor__(Companion_Default___.Fm21(p0, !(p0), globalState), p1)
	_171_v107 = _nw19
	var _172_v108 _dafny.Set
	_ = _172_v108
	_172_v108 = _dafny.SetOf(_171_v107, _171_v107)
	_172_v108 = (_dafny.SetOf(_171_v107)).Intersection(_172_v108)
	var _173_v109 _dafny.Int
	_ = _173_v109
	var _174_v110 _dafny.Int
	_ = _174_v110
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	_out0, _out1 = (_171_v107).M0(p0, globalState)
	_173_v109 = _out0
	_174_v110 = _out1
	var _175_v111 _dafny.MultiSet
	_ = _175_v111
	_175_v111 = _dafny.MultiSetOf(true)
	r0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((func() _dafny.Int {
		if p0 {
			return Companion_Default___.Fm1(p0, _48_v7, _173_v109, globalState)
		}
		return (func() _dafny.Int {
			if (_175_v111).Contains(p0) {
				return (_175_v111).Multiplicity(p0)
			}
			return p1
		})()
	})()), _173_v109)
	r1 = p0
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _176_v0 bool
	_ = _176_v0
	_176_v0 = true
	var _177_v1 _dafny.Int
	_ = _177_v1
	_177_v1 = _dafny.IntOfInt64(17)
	var _178_v2 _dafny.Sequence
	_ = _178_v2
	_178_v2 = _dafny.SeqOf(_177_v1)
	var _179_v3 _dafny.Map
	_ = _179_v3
	_179_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v0, _178_v2)
	var _180_v4 _dafny.Array
	_ = _180_v4
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_1
	var _nw20 _dafny.Array
	_ = _nw20
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw20 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = (func(_181_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_182_i1 _dafny.Int) _dafny.Int {
				return (_182_i1).Plus(_181_v1)
			}
		})(_177_v1)
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
	_180_v4 = _nw20
	var _183_v5 _dafny.Sequence
	_ = _183_v5
	_183_v5 = _dafny.UnicodeSeqOfUtf8Bytes("cb")
	var _184_v6 _dafny.MultiSet
	_ = _184_v6
	_184_v6 = _dafny.MultiSetOf(_183_v5)
	var _185_v8 _dafny.MultiSet
	_ = _185_v8
	_185_v8 = _dafny.MultiSetOf(_176_v0, _176_v0, _176_v0)
	var _186_v9 _dafny.Array
	_ = _186_v9
	var _nwElement0_3 bool = _176_v0
	_ = _nwElement0_3
	var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(11))
	_ = _nw21
	(_nw21).ArraySet1(_nwElement0_3, 0)
	(_nw21).ArraySet1(false, 1)
	(_nw21).ArraySet1(_176_v0, 2)
	(_nw21).ArraySet1(_176_v0, 3)
	(_nw21).ArraySet1(_176_v0, 4)
	(_nw21).ArraySet1(_176_v0, 5)
	(_nw21).ArraySet1(_176_v0, 6)
	(_nw21).ArraySet1(_176_v0, 7)
	(_nw21).ArraySet1(_176_v0, 8)
	(_nw21).ArraySet1(_176_v0, 9)
	(_nw21).ArraySet1(_176_v0, 10)
	_186_v9 = _nw21
	var _187_v10 _dafny.CodePoint
	_ = _187_v10
	_187_v10 = _dafny.CodePoint('u')
	var _188_v11 _dafny.MultiSet
	_ = _188_v11
	_188_v11 = _dafny.MultiSetOf(_187_v10, _187_v10, _dafny.CodePoint('f'), _187_v10)
	var _189_v12 _dafny.Map
	_ = _189_v12
	_189_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v1, _dafny.MultiSetOf(_187_v10))
	var _190_globalState *GlobalState
	_ = _190_globalState
	var _nw22 *GlobalState = New_GlobalState_()
	_ = _nw22
	_nw22.Ctor__(_179_v3, _dafny.IntOfInt64(201), _dafny.IntOfInt64(-912), _dafny.IntOfInt64(-807), false, _dafny.IntOfInt64(990), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(144))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_191_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), _dafny.IntOfInt64(150), _dafny.IntOfInt64(944), true, _dafny.IntOfInt64(334), _180_v4, func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_184_v6).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _192_v7 _dafny.Sequence
			_192_v7 = interface{}(_compr_6).(_dafny.Sequence)
			if (_184_v6).Contains(_192_v7) {
				_coll6.Add(_192_v7)
			}
		}
		return _coll6.ToSet()
	}(), _dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), true, _dafny.IntOfInt64(850), _185_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_193_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})), _186_v9, _dafny.IntOfInt64(-870), _183_v5, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v1, _188_v11)).Merge(_189_v12), false, false, _183_v5, true)
	_190_globalState = _nw22
	var _194_v13 _dafny.Sequence
	_ = _194_v13
	_194_v13 = _dafny.SeqOf(_180_v4, _180_v4, _180_v4, _180_v4)
	_194_v13 = _194_v13
	var _195_v14 _dafny.Map
	_ = _195_v14
	_195_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_177_v1), _176_v0)
	var _196_v15 _dafny.Map
	_ = _196_v15
	_196_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v0, _176_v0)
	if !(Companion_Default___.Fm0((func() bool {
		if (_195_v14).Contains(_177_v1) {
			return (_195_v14).Get(_177_v1).(bool)
		}
		return _176_v0
	})(), (_196_v15).Merge(_196_v15), _190_globalState)) {
		(_190_globalState).F10 = (_dafny.Zero).Minus(_177_v1)
		(_190_globalState).F22 = !(_176_v0) || (_176_v0)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))
		_ = _index18
		(_180_v4).ArraySet1(_dafny.IntOfInt64(945), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))
		_ = _index19
		(_180_v4).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(669), _177_v1), (_index19).Int())
		if _176_v0 {
			(_190_globalState).F2 = ((_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int)).Times((_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int))
			var _197_v16 _dafny.Sequence
			_ = _197_v16
			_197_v16 = _dafny.SeqOf(_176_v0)
			var _198_v17 _dafny.MultiSet
			_ = _198_v17
			_198_v17 = _dafny.MultiSetOf(_dafny.IntOfUint32((_178_v2).Cardinality()), _dafny.IntOfUint32((_197_v16).Cardinality()), _177_v1)
			var _199_v18 _dafny.MultiSet
			_ = _199_v18
			_199_v18 = _dafny.MultiSetOf(_198_v17)
			var _200_v19 _dafny.MultiSet
			_ = _200_v19
			_200_v19 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_199_v18).Contains(_198_v17) {
					return (_199_v18).Multiplicity(_198_v17)
				}
				return _177_v1
			})())
			var _201_v20 _dafny.Map
			_ = _201_v20
			_201_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v0, (_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int))
			_176_v0 = ((func() _dafny.Int {
				if (_185_v8).Contains(_176_v0) {
					return (_185_v8).Multiplicity(_176_v0)
				}
				return (_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int)
			})()).Cmp(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_198_v17).Contains(Companion_Default___.Fm1(_176_v0, _187_v10, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_178_v2)).Cardinality(), (_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int))).Cardinality(), _190_globalState)) {
					return (_198_v17).Multiplicity(Companion_Default___.Fm1(_176_v0, _187_v10, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_178_v2)).Cardinality(), (_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int))).Cardinality(), _190_globalState))
				}
				return (_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int)
			})(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_201_v20).Cardinality(), false)).Cardinality()))) <= 0
			(_190_globalState).F2 = Companion_Default___.SafeModuloInt(_177_v1, (_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int))
			var _202_v21 _dafny.Sequence
			_ = _202_v21
			_202_v21 = _dafny.SeqOf((_200_v19).Union(_198_v17), (_dafny.MultiSetFromSeq(_178_v2)).Intersection(Companion_Default___.Fm2(Companion_Default___.Fm1(_176_v0, Companion_Default___.Fm3(_177_v1, _176_v0, _190_globalState), _177_v1, _190_globalState), _176_v0, _190_globalState)))
			var _203_v22 D0
			_ = _203_v22
			_203_v22 = Companion_D0_.Create_DC0_(_202_v21)
			_202_v21 = (_203_v22).Dtor_cf0()
			_183_v5 = _183_v5
		} else {
			_183_v5 = _183_v5
			var _204_v23 *C0
			_ = _204_v23
			var _nw23 *C0 = New_C0_()
			_ = _nw23
			_nw23.Ctor__()
			_204_v23 = _nw23
			var _205_v24 _dafny.Array
			_ = _205_v24
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_2
			var _nw24 _dafny.Array
			_ = _nw24
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw24 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) D3 = (func(_206_v10 _dafny.CodePoint, _207_v0 bool) func(_dafny.Int) D3 {
					return func(_208_i3 _dafny.Int) D3 {
						return Companion_D3_.Create_DC11_(_206_v10, _207_v0, true)
					}
				})(_187_v10, _176_v0)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw24 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw24).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw24).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_205_v24 = _nw24
			var _209_v25 D3
			_ = _209_v25
			_209_v25 = Companion_D3_.Create_DC11_(_187_v10, _176_v0, _176_v0)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_205_v24), 0))
			_ = _index20
			(_205_v24).ArraySet1(_209_v25, (_index20).Int())
			var _pat_let_tv5 = _180_v4
			_ = _pat_let_tv5
			var _pat_let_tv6 = _180_v4
			_ = _pat_let_tv6
			var _pat_let_tv7 = _176_v0
			_ = _pat_let_tv7
			var _pat_let_tv8 = _190_globalState
			_ = _pat_let_tv8
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_205_v24), 0))
			_ = _index21
			(_205_v24).ArraySet1(func(_pat_let2_0 D3) D3 {
				return func(_210_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let3_0 _dafny.CodePoint) D3 {
						return func(_211_dt__update_hcf19_h0 _dafny.CodePoint) D3 {
							return Companion_D3_.Create_DC11_(_211_dt__update_hcf19_h0, (_210_dt__update__tmp_h0).Dtor_cf20(), (_210_dt__update__tmp_h0).Dtor_cf21())
						}(_pat_let3_0)
					}(Companion_Default___.Fm3((_pat_let_tv6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_pat_let_tv5), 0))).Int()).(_dafny.Int), _pat_let_tv7, _pat_let_tv8))
				}(_pat_let2_0)
			}(_209_v25), (_index21).Int())
			var _212_v26 _dafny.CodePoint
			_ = _212_v26
			_212_v26 = _187_v10
			_212_v26 = _187_v10
			(_190_globalState).F19 = (func() _dafny.Int {
				if (_185_v8).Contains((false) && (!(_176_v0))) {
					return (_185_v8).Multiplicity((false) && (!(_176_v0)))
				}
				return _177_v1
			})()
		}
		var _213_v27 _dafny.Sequence
		_ = _213_v27
		_213_v27 = _dafny.SeqOf(_dafny.SetOf(_176_v0, _176_v0), _dafny.SetOf(_176_v0, _176_v0))
		var _214_v28 *C1
		_ = _214_v28
		var _nw25 *C1 = New_C1_()
		_ = _nw25
		_nw25.Ctor__((_213_v27).Select((Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_213_v27).Cardinality()))).Uint32()).(_dafny.Set), _177_v1)
		_214_v28 = _nw25
		var _215_v29 _dafny.Sequence
		_ = _215_v29
		_215_v29 = _dafny.SeqOf(_176_v0, !(_176_v0), (Companion_D3_.Create_DC11_(_187_v10, _176_v0, _176_v0)).Dtor_cf20())
		var _216_v30 _dafny.Map
		_ = _216_v30
		_216_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v29, _214_v28)
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))
		_ = _index22
		var _rhs25 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_178_v2).Cardinality()), _177_v1)
		_ = _rhs25
		var _rhs26 *C1 = (func() *C1 {
			if (_216_v30).Contains(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if false {
					return _215_v29
				}
				return _dafny.SeqOf(_176_v0, _176_v0, _176_v0)
			})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_177_v1), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if false {
					return _215_v29
				}
				return _dafny.SeqOf(_176_v0, _176_v0, _176_v0)
			})()).Cardinality()))).Uint32(), _176_v0)) {
				return (_216_v30).Get(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if false {
						return _215_v29
					}
					return _dafny.SeqOf(_176_v0, _176_v0, _176_v0)
				})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_177_v1), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if false {
						return _215_v29
					}
					return _dafny.SeqOf(_176_v0, _176_v0, _176_v0)
				})()).Cardinality()))).Uint32(), _176_v0)).(*C1)
			}
			return _214_v28
		})()
		_ = _rhs26
		var _lhs21 _dafny.Array = _180_v4
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_180_v4), 0))
		_ = _lhs22
		(_lhs21).ArraySet1(_rhs25, (_lhs22).Int())
		_214_v28 = _rhs26
	} else {
		if (_177_v1).Cmp((_185_v8).Cardinality()) > 0 {
			var _217_v31 T0
			_ = _217_v31
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__()
			_217_v31 = _nw26
			var _218_v32 _dafny.Sequence
			_ = _218_v32
			_218_v32 = _dafny.SeqOf(_dafny.MultiSetOf(_177_v1, _177_v1))
			var _219_v33 D0
			_ = _219_v33
			_219_v33 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Concatenate(_218_v32, _218_v32))
			var _220_v34 _dafny.Sequence
			_ = _220_v34
			_220_v34 = _dafny.SeqOf(true)
			var _rhs27 _dafny.Int = (func() _dafny.Int {
				if false {
					return _177_v1
				}
				return (func() _dafny.Int {
					if (_220_v34).Select((Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_220_v34).Cardinality()))).Uint32()).(bool) {
						return _177_v1
					}
					return _177_v1
				})()
			})()
			_ = _rhs27
			var _rhs28 _dafny.Int = (_177_v1).Plus(_177_v1)
			_ = _rhs28
			var _rhs29 T0 = _217_v31
			_ = _rhs29
			var _rhs30 bool = _176_v0
			_ = _rhs30
			var _rhs31 D0 = _219_v33
			_ = _rhs31
			var _lhs23 *GlobalState = _190_globalState
			_ = _lhs23
			var _lhs24 *GlobalState = _190_globalState
			_ = _lhs24
			var _lhs25 *GlobalState = _190_globalState
			_ = _lhs25
			_lhs23.F10 = _rhs27
			_lhs24.F19 = _rhs28
			_217_v31 = _rhs29
			_lhs25.F23 = _rhs30
			_219_v33 = _rhs31
			_217_v31 = _217_v31
			_195_v14 = (_195_v14).Update(_177_v1, _176_v0)
			var _221_v35 _dafny.Set
			_ = _221_v35
			_221_v35 = _dafny.SetOf(_176_v0, _176_v0)
			var _222_v36 *C1
			_ = _222_v36
			var _nw27 *C1 = New_C1_()
			_ = _nw27
			_nw27.Ctor__(_221_v35, (_177_v1).Times(_dafny.IntOfInt64(107)))
			_222_v36 = _nw27
			var _223_v37 _dafny.Map
			_ = _223_v37
			_223_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v36, _187_v10)
			_223_v37 = (_223_v37).Update(_222_v36, _dafny.CodePoint('x'))
		} else {
			var _224_v38 _dafny.Array
			_ = _224_v38
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_3
			var _nw28 _dafny.Array
			_ = _nw28
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw28 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_225_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_226_i4 _dafny.Int) _dafny.Sequence {
						return _225_v2
					}
				})(_178_v2)
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
			_224_v38 = _nw28
			var _227_v39 _dafny.Set
			_ = _227_v39
			_227_v39 = _dafny.SetOf(Companion_Default___.SafeDivisionInt(_177_v1, _177_v1))
			var _228_v40 _dafny.Set
			_ = _228_v40
			_228_v40 = _dafny.SetOf(_176_v0, _176_v0, _176_v0, _176_v0)
			var _229_v41 *C0
			_ = _229_v41
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__()
			_229_v41 = _nw29
			var _230_v42 _dafny.Map
			_ = _230_v42
			_230_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v41, _176_v0)
			var _231_v43 *C1
			_ = _231_v43
			var _nw30 *C1 = New_C1_()
			_ = _nw30
			_nw30.Ctor__((_228_v40).Union(_dafny.SetOf((func() bool {
				if (_230_v42).Contains(_229_v41) {
					return (_230_v42).Get(_229_v41).(bool)
				}
				return _176_v0
			})(), _176_v0, !(_176_v0))), (_229_v41).Fm4(_177_v1, _176_v0, _190_globalState))
			_231_v43 = _nw30
			var _232_v44 _dafny.Sequence
			_ = _232_v44
			_232_v44 = _dafny.SeqOf(_231_v43)
			var _rhs32 _dafny.Array = _224_v38
			_ = _rhs32
			var _rhs33 _dafny.Set = _dafny.SetOf(_231_v43.F27, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_232_v44).Cardinality()), (_178_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_183_v5).Cardinality()), _dafny.IntOfUint32((_178_v2).Cardinality()))).Uint32()).(_dafny.Int)), _177_v1)
			_ = _rhs33
			var _rhs34 *C1 = _231_v43
			_ = _rhs34
			var _rhs35 _dafny.Int = _231_v43.F27
			_ = _rhs35
			var _lhs26 *GlobalState = _190_globalState
			_ = _lhs26
			_224_v38 = _rhs32
			_227_v39 = _rhs33
			_231_v43 = _rhs34
			_lhs26.F19 = _rhs35
			var _233_v45 _dafny.Array
			_ = _233_v45
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_4
			var _nw31 _dafny.Array
			_ = _nw31
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw31 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) D4 = (func(_234_v8 _dafny.MultiSet) func(_dafny.Int) D4 {
					return func(_235_i5 _dafny.Int) D4 {
						return Companion_D4_.Create_DC13_(_234_v8)
					}
				})(_185_v8)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw31 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw31).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw31).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_233_v45 = _nw31
			var _236_v46 T0
			_ = _236_v46
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__((_231_v43).F26(), _177_v1)
			_236_v46 = _nw32
			var _237_v47 _dafny.Sequence
			_ = _237_v47
			_237_v47 = _dafny.SeqOf(_236_v46, _236_v46)
			var _238_v48 _dafny.Map
			_ = _238_v48
			_238_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_237_v47).Select((Companion_Default___.SafeIndex((_178_v2).Select((Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_178_v2).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_237_v47).Cardinality()))).Uint32()).(T0), _177_v1)
			var _239_v49 _dafny.Map
			_ = _239_v49
			_239_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v0, _233_v45)
			var _rhs36 _dafny.Array = (func() _dafny.Array {
				if (_239_v49).Contains((func() bool {
					if _176_v0 {
						return false
					}
					return _176_v0
				})()) {
					return (_239_v49).Get((func() bool {
						if _176_v0 {
							return false
						}
						return _176_v0
					})()).(_dafny.Array)
				}
				return _233_v45
			})()
			_ = _rhs36
			var _rhs37 _dafny.Map = _238_v48
			_ = _rhs37
			var _rhs38 _dafny.Int = _231_v43.F27
			_ = _rhs38
			var _lhs27 *GlobalState = _190_globalState
			_ = _lhs27
			_233_v45 = _rhs36
			_238_v48 = _rhs37
			_lhs27.F15 = _rhs38
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_180_v4), 0))
			_ = _index23
			(_180_v4).ArraySet1(_231_v43.F27, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_180_v4), 0))
			_ = _index24
			(_180_v4).ArraySet1(_231_v43.F27, (_index24).Int())
			var _240_v50 _dafny.Array
			_ = _240_v50
			var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
			_ = _nw33
			_240_v50 = _nw33
			var _241_v51 _dafny.Map
			_ = _241_v51
			_241_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v0, _177_v1)
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_240_v50), 0))
			_ = _index25
			(_240_v50).ArraySet1(_241_v51, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_240_v50), 0))
			_ = _index26
			(_240_v50).ArraySet1(_241_v51, (_index26).Int())
			var _242_v52 D0
			_ = _242_v52
			_242_v52 = Companion_D0_.Create_DC1_((_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int))
			var _243_v53 D3
			_ = _243_v53
			_243_v53 = Companion_D3_.Create_DC12_(false, _242_v52, _176_v0, _176_v0, _236_v46)
			var _244_v54 _dafny.Array
			_ = _244_v54
			var _nwElement0_4 T0 = _236_v46
			_ = _nwElement0_4
			var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(19))
			_ = _nw34
			(_nw34).ArraySet1(_nwElement0_4, 0)
			(_nw34).ArraySet1(_236_v46, 1)
			(_nw34).ArraySet1(_236_v46, 2)
			(_nw34).ArraySet1(_236_v46, 3)
			(_nw34).ArraySet1(_236_v46, 4)
			(_nw34).ArraySet1(_236_v46, 5)
			(_nw34).ArraySet1(_236_v46, 6)
			(_nw34).ArraySet1((_243_v53).Dtor_cf26(), 7)
			(_nw34).ArraySet1(_236_v46, 8)
			(_nw34).ArraySet1(_236_v46, 9)
			(_nw34).ArraySet1(_236_v46, 10)
			(_nw34).ArraySet1(_236_v46, 11)
			(_nw34).ArraySet1((func() T0 {
				if _176_v0 {
					return (Companion_D3_.Create_DC12_(_176_v0, _242_v52, _176_v0, _176_v0, _236_v46)).Dtor_cf26()
				}
				return _236_v46
			})(), 12)
			(_nw34).ArraySet1(_236_v46, 13)
			(_nw34).ArraySet1(_236_v46, 14)
			(_nw34).ArraySet1(_236_v46, 15)
			(_nw34).ArraySet1(_236_v46, 16)
			(_nw34).ArraySet1(_236_v46, 17)
			(_nw34).ArraySet1(_236_v46, 18)
			_244_v54 = _nw34
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_244_v54), 0))
			_ = _index27
			(_244_v54).ArraySet1(_236_v46, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_244_v54), 0))
			_ = _index28
			(_244_v54).ArraySet1(_236_v46, (_index28).Int())
		}
		var _245_v55 D2
		_ = _245_v55
		_245_v55 = Companion_D2_.Create_DC8_(_dafny.IntOfUint32((_183_v5).Cardinality()), _176_v0, _177_v1)
		if (_245_v55).Dtor_cf13() {
			var _246_v56 D3
			_ = _246_v56
			_246_v56 = Companion_D3_.Create_DC11_(_187_v10, _176_v0, _176_v0)
			var _pat_let_tv9 = _176_v0
			_ = _pat_let_tv9
			var _247_v57 _dafny.Sequence
			_ = _247_v57
			_247_v57 = _dafny.SeqOf(func(_pat_let4_0 D3) D3 {
				return func(_248_dt__update__tmp_h2 D3) D3 {
					return func(_pat_let5_0 bool) D3 {
						return func(_249_dt__update_hcf21_h0 bool) D3 {
							return Companion_D3_.Create_DC11_((_248_dt__update__tmp_h2).Dtor_cf19(), (_248_dt__update__tmp_h2).Dtor_cf20(), _249_dt__update_hcf21_h0)
						}(_pat_let5_0)
					}(_pat_let_tv9)
				}(_pat_let4_0)
			}(_246_v56), Companion_Default___.Fm11(_177_v1, _190_globalState), Companion_D3_.Create_DC11_(_187_v10, _176_v0, _176_v0))
			(_190_globalState).F5 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_247_v57, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(876))).Uint32(), func(coer10 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_250_v10 _dafny.CodePoint, _251_v0 bool) func(_dafny.Int) D3 {
				return func(_252_i6 _dafny.Int) D3 {
					return Companion_D3_.Create_DC11_(_250_v10, _251_v0, _251_v0)
				}
			})(_187_v10, _176_v0))))).Cardinality())
			(_190_globalState).F25 = _176_v0
			(_190_globalState).F2 = (_177_v1).Plus(Companion_Default___.SafeDivisionInt(_177_v1, _177_v1))
			var _253_v58 _dafny.Array
			_ = _253_v58
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_5
			var _nw35 _dafny.Array
			_ = _nw35
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw35 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) D3 = (func(_254_v10 _dafny.CodePoint, _255_v0 bool) func(_dafny.Int) D3 {
					return func(_256_i7 _dafny.Int) D3 {
						return Companion_D3_.Create_DC11_(_254_v10, !(_255_v0), !(!(!(_255_v0))))
					}
				})(_187_v10, _176_v0)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw35 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw35).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw35).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_253_v58 = _nw35
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_253_v58), 0))
			_ = _index29
			(_253_v58).ArraySet1(_246_v56, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_253_v58), 0))
			_ = _index30
			(_253_v58).ArraySet1(Companion_Default___.Fm11(_177_v1, _190_globalState), (_index30).Int())
			var _257_v59 D0
			_ = _257_v59
			_257_v59 = Companion_D0_.Create_DC1_(_177_v1)
			var _258_v60 _dafny.Set
			_ = _258_v60
			_258_v60 = _dafny.SetOf(_176_v0, _176_v0)
			var _259_v61 T0
			_ = _259_v61
			var _nw36 *C1 = New_C1_()
			_ = _nw36
			_nw36.Ctor__(_258_v60, _177_v1)
			_259_v61 = _nw36
			var _260_v62 D3
			_ = _260_v62
			_260_v62 = Companion_D3_.Create_DC12_(_176_v0, _257_v59, _176_v0, _176_v0, _259_v61)
			var _261_v63 _dafny.Map
			_ = _261_v63
			_261_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v62, _177_v1)
			var _262_v64 _dafny.Map
			_ = _262_v64
			_262_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v1, _261_v63)
			(_190_globalState).F22 = !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v62, _177_v1)).Equals((func() _dafny.Map {
				if (_262_v64).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_178_v2).Cardinality()))) {
					return (_262_v64).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_178_v2).Cardinality()))).(_dafny.Map)
				}
				return _261_v63
			})()))
		} else {
			var _263_v65 _dafny.Set
			_ = _263_v65
			_263_v65 = _dafny.SetOf(_176_v0)
			var _264_v66 *C1
			_ = _264_v66
			var _nw37 *C1 = New_C1_()
			_ = _nw37
			_nw37.Ctor__(_263_v65, _177_v1)
			_264_v66 = _nw37
			var _265_v67 D0
			_ = _265_v67
			_265_v67 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-492))
			var _266_v68 _dafny.Array
			_ = _266_v68
			var _nwElement0_5 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(28))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}(func(_267_i8 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cwxonbmhb")).Cardinality())
			}))
			_ = _nwElement0_5
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(13))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_5, 0)
			(_nw38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_178_v2, Companion_Default___.Fm12(_177_v1, _265_v67, _190_globalState)), 1)
			(_nw38).ArraySet1(_178_v2, 2)
			(_nw38).ArraySet1(_dafny.SeqOf(_264_v66.F27), 3)
			(_nw38).ArraySet1(_178_v2, 4)
			(_nw38).ArraySet1(_178_v2, 5)
			(_nw38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_178_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(762))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_268_v8 _dafny.MultiSet, _269_v0 bool, _270_v66 *C1) func(_dafny.Int) _dafny.Int {
				return func(_271_i9 _dafny.Int) _dafny.Int {
					return ((_268_v8).Update(_269_v0, Companion_Default___.Abs((_dafny.Zero).Minus(_270_v66.F27)))).Cardinality()
				}
			})(_185_v8, _176_v0, _264_v66)))), 6)
			(_nw38).ArraySet1(_178_v2, 7)
			(_nw38).ArraySet1(_178_v2, 8)
			(_nw38).ArraySet1((func() _dafny.Sequence {
				if _176_v0 {
					return _178_v2
				}
				return _178_v2
			})(), 9)
			(_nw38).ArraySet1(_178_v2, 10)
			(_nw38).ArraySet1(_178_v2, 11)
			(_nw38).ArraySet1(_178_v2, 12)
			_266_v68 = _nw38
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_266_v68), 0))
			_ = _index31
			(_266_v68).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_178_v2, _178_v2), (_index31).Int())
			var _272_v69 _dafny.Set
			_ = _272_v69
			_272_v69 = _dafny.SetOf(_264_v66.F27, _177_v1)
			var _273_v70 _dafny.Sequence
			_ = _273_v70
			_273_v70 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_178_v2, (Companion_Default___.SafeIndex((_272_v69).Cardinality(), _dafny.IntOfUint32((_178_v2).Cardinality()))).Uint32(), _264_v66.F27), (Companion_Default___.SafeIndex(_264_v66.F27, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_178_v2, (Companion_Default___.SafeIndex((_272_v69).Cardinality(), _dafny.IntOfUint32((_178_v2).Cardinality()))).Uint32(), _264_v66.F27)).Cardinality()))).Uint32(), _177_v1), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-788))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_274_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_275_i10 _dafny.Int) _dafny.Int {
					return _274_v1
				}
			})(_177_v1))), _178_v2, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(377))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_276_v66 *C1) func(_dafny.Int) _dafny.Int {
				return func(_277_i11 _dafny.Int) _dafny.Int {
					return _276_v66.F27
				}
			})(_264_v66))), _178_v2))
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_266_v68), 0))
			_ = _index32
			(_266_v68).ArraySet1((_273_v70).Select((Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_273_v70).Cardinality()))).Uint32()).(_dafny.Sequence), (_index32).Int())
			(_190_globalState).F15 = _264_v66.F27
			var _278_v71 D2
			_ = _278_v71
			_278_v71 = Companion_D2_.Create_DC7_(_176_v0)
			var _rhs39 D2 = _278_v71
			_ = _rhs39
			var _rhs40 _dafny.Int = (_264_v66.F27).Times(_177_v1)
			_ = _rhs40
			var _lhs28 *GlobalState = _190_globalState
			_ = _lhs28
			_278_v71 = _rhs39
			_lhs28.F8 = _rhs40
			(_190_globalState).F22 = _dafny.Companion_Sequence_.IsPrefixOf(_183_v5, _183_v5)
		}
		(_190_globalState).F10 = Companion_Default___.Fm1(_176_v0, _187_v10, _177_v1, _190_globalState)
		(_190_globalState).F19 = _177_v1
		if (_185_v8).IsDisjointFrom(_dafny.MultiSetOf(_176_v0, _176_v0, true)) {
			var _279_v72 _dafny.Map
			_ = _279_v72
			_279_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_183_v5).Cardinality()), _177_v1)
			(_190_globalState).F23 = !_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm13(_190_globalState), _279_v72)
			var _280_v73 _dafny.Map
			_ = _280_v73
			_280_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v8, _dafny.IntOfInt64(363))
			_280_v73 = (_280_v73).Update(_dafny.MultiSetOf(_176_v0), _177_v1)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_180_v4), 0))
			_ = _index33
			(_180_v4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), (Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5)).Cardinality()))).Uint32(), _187_v10)).Cardinality()), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_180_v4), 0))
			_ = _index34
			(_180_v4).ArraySet1(_177_v1, (_index34).Int())
			var _281_v74 _dafny.Map
			_ = _281_v74
			_281_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_177_v1)), _176_v0)
			var _rhs41 _dafny.Map = (_281_v74).Merge(_281_v74)
			_ = _rhs41
			var _rhs42 _dafny.Int = _dafny.IntOfInt64(465)
			_ = _rhs42
			var _rhs43 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_178_v2, _178_v2), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(23))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_282_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_283_i12 _dafny.Int) _dafny.Int {
					return _282_v1
				}
			})(_177_v1))))
			_ = _rhs43
			var _rhs44 bool = Companion_Default___.Fm0((_176_v0) == (_176_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v0, Companion_Default___.Fm0(_176_v0, _196_v15, _190_globalState)), _190_globalState)
			_ = _rhs44
			var _rhs45 _dafny.Int = (_180_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_180_v4), 0))).Int()).(_dafny.Int)
			_ = _rhs45
			var _lhs29 *GlobalState = _190_globalState
			_ = _lhs29
			var _lhs30 *GlobalState = _190_globalState
			_ = _lhs30
			var _lhs31 *GlobalState = _190_globalState
			_ = _lhs31
			_281_v74 = _rhs41
			_lhs29.F15 = _rhs42
			_178_v2 = _rhs43
			_lhs30.F25 = _rhs44
			_lhs31.F2 = _rhs45
			var _284_v75 _dafny.Set
			_ = _284_v75
			_284_v75 = _dafny.SetOf(_176_v0)
			var _285_v76 _dafny.Array
			_ = _285_v76
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw39
			_285_v76 = _nw39
			var _286_v77 _dafny.Set
			_ = _286_v77
			_286_v77 = _dafny.SetOf(_285_v76)
			var _287_v78 _dafny.Sequence
			_ = _287_v78
			_287_v78 = _dafny.SeqOf(_176_v0)
			var _288_v79 D0
			_ = _288_v79
			_288_v79 = Companion_D0_.Create_DC2_(_176_v0, _286_v77, !(_176_v0), _176_v0, _dafny.IntOfUint32((_287_v78).Cardinality()))
			var _289_v80 *C1
			_ = _289_v80
			var _nw40 *C1 = New_C1_()
			_ = _nw40
			_nw40.Ctor__(_284_v75, (_288_v79).Dtor_cf6())
			_289_v80 = _nw40
			var _290_v81 _dafny.Set
			_ = _290_v81
			_290_v81 = _dafny.SetOf(_289_v80)
			(_190_globalState).F25 = ((_290_v81).Cardinality()).Cmp(_dafny.IntOfInt64(549)) <= 0
		} else {
			_178_v2 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(762))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_291_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_292_i13 _dafny.Int) _dafny.Int {
					return _291_v1
				}
			})(_177_v1)))
			var _293_v82 _dafny.Set
			_ = _293_v82
			_293_v82 = _dafny.SetOf(_176_v0, _176_v0, _176_v0)
			var _294_v83 _dafny.CodePoint
			_ = _294_v83
			_294_v83 = _187_v10
			var _295_v84 _dafny.Sequence
			_ = _295_v84
			_295_v84 = _dafny.SeqOf(_176_v0, _176_v0)
			var _296_v85 _dafny.Set
			_ = _296_v85
			_296_v85 = _dafny.SetOf(_177_v1)
			var _rhs46 _dafny.Array = _180_v4
			_ = _rhs46
			var _rhs47 bool = Companion_Default___.Fm0((true) == (Companion_Default___.Fm0(_176_v0, _196_v15, _190_globalState)), (_196_v15).Merge(_196_v15), _190_globalState)
			_ = _rhs47
			var _rhs48 bool = (_185_v8).IsSubsetOf((_dafny.MultiSetFromSeq(_295_v84)).Union(_185_v8))
			_ = _rhs48
			var _rhs49 _dafny.Set = _dafny.SetOf(_176_v0, _176_v0, !(_176_v0), (_295_v84).Select((Companion_Default___.SafeIndex((_296_v85).Cardinality(), _dafny.IntOfUint32((_295_v84).Cardinality()))).Uint32()).(bool))
			_ = _rhs49
			var _rhs50 _dafny.CodePoint = _187_v10
			_ = _rhs50
			var _lhs32 *GlobalState = _190_globalState
			_ = _lhs32
			var _lhs33 *GlobalState = _190_globalState
			_ = _lhs33
			_lhs32.F11 = _rhs46
			_lhs33.F23 = _rhs47
			_176_v0 = _rhs48
			_293_v82 = _rhs49
			_294_v83 = _rhs50
			var _297_v86 _dafny.Array
			_ = _297_v86
			var _nwElement0_6 _dafny.CodePoint = _187_v10
			_ = _nwElement0_6
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(9))
			_ = _nw41
			(_nw41).ArraySet1CodePoint(_nwElement0_6, 0)
			(_nw41).ArraySet1CodePoint(_187_v10, 1)
			(_nw41).ArraySet1CodePoint(_187_v10, 2)
			(_nw41).ArraySet1CodePoint(_187_v10, 3)
			(_nw41).ArraySet1CodePoint(_187_v10, 4)
			(_nw41).ArraySet1CodePoint(_187_v10, 5)
			(_nw41).ArraySet1CodePoint(_dafny.CodePoint('k'), 6)
			(_nw41).ArraySet1CodePoint(_187_v10, 7)
			(_nw41).ArraySet1CodePoint(_dafny.CodePoint('u'), 8)
			_297_v86 = _nw41
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_297_v86), 0))
			_ = _index35
			(_297_v86).ArraySet1CodePoint(_dafny.CodePoint('h'), (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_297_v86), 0))
			_ = _index36
			(_297_v86).ArraySet1CodePoint(_187_v10, (_index36).Int())
			(_190_globalState).F22 = _176_v0
			var _298_v87 _dafny.MultiSet
			_ = _298_v87
			_298_v87 = _dafny.MultiSetOf(_177_v1)
			var _299_v88 _dafny.Sequence
			_ = _299_v88
			_299_v88 = _dafny.SeqOf(_298_v87, _298_v87, _298_v87)
			var _300_v89 D0
			_ = _300_v89
			_300_v89 = Companion_D0_.Create_DC0_(_299_v88)
			_300_v89 = Companion_Default___.Fm14(_190_globalState)
		}
	}
	(_190_globalState).F22 = !((func() _dafny.Map {
		if false {
			return _196_v15
		}
		return _196_v15
	})()).Equals(_196_v15)
	var _301_v90 _dafny.Set
	_ = _301_v90
	_301_v90 = _dafny.SetOf(_177_v1, _177_v1)
	if (_301_v90).IsSubsetOf((_301_v90).Difference(_301_v90)) {
		if !(_176_v0) || (_176_v0) {
			var _302_v91 _dafny.MultiSet
			_ = _302_v91
			_302_v91 = _dafny.MultiSetOf(_177_v1)
			var _303_v92 D2
			_ = _303_v92
			_303_v92 = Companion_D2_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("tvlnsoaai"))
			var _304_v93 D4
			_ = _304_v93
			_304_v93 = Companion_D4_.Create_DC14_(_302_v91, _177_v1, _303_v92)
			(_190_globalState).F2 = (func() _dafny.Int {
				if _176_v0 {
					return _177_v1
				}
				return (_304_v93).Dtor_cf29()
			})()
			(_190_globalState).F22 = Companion_Default___.Fm0(_176_v0, _196_v15, _190_globalState)
			var _305_v94 *C0
			_ = _305_v94
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__()
			_305_v94 = _nw42
			var _306_v95 _dafny.Map
			_ = _306_v95
			_306_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-221))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_307_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_308_i14 _dafny.Int) _dafny.CodePoint {
					return _307_v10
				}
			})(_187_v10))), _dafny.Companion_Sequence_.Concatenate(_183_v5, _dafny.UnicodeSeqOfUtf8Bytes("jkvceykyp")))
			_306_v95 = (_306_v95).Update(_183_v5, _183_v5)
			var _309_v96 _dafny.Map
			_ = _309_v96
			_309_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v1, _dafny.UnicodeSeqOfUtf8Bytes("tgejkifdl"))
			var _310_v97 _dafny.Map
			_ = _310_v97
			_310_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_309_v96, _177_v1)
			_310_v97 = (_310_v97).Update(_309_v96, _177_v1)
		} else {
			_187_v10 = _dafny.CodePoint('w')
			(_190_globalState).F8 = Companion_Default___.SafeDivisionInt(_177_v1, _177_v1)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_186_v9), 0))
			_ = _index37
			(_186_v9).ArraySet1(_176_v0, (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_186_v9), 0))
			_ = _index38
			(_186_v9).ArraySet1(!(_176_v0) || (_176_v0), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_186_v9), 0))
			_ = _index39
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_186_v9), 0))
			_ = _index40
			var _rhs51 bool = _176_v0
			_ = _rhs51
			var _rhs52 _dafny.CodePoint = _187_v10
			_ = _rhs52
			var _rhs53 bool = _176_v0
			_ = _rhs53
			var _rhs54 bool = _176_v0
			_ = _rhs54
			var _lhs34 *GlobalState = _190_globalState
			_ = _lhs34
			var _lhs35 _dafny.Array = _186_v9
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_186_v9), 0))
			_ = _lhs36
			var _lhs37 _dafny.Array = _186_v9
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_186_v9), 0))
			_ = _lhs38
			_lhs34.F23 = _rhs51
			_187_v10 = _rhs52
			(_lhs35).ArraySet1(_rhs53, (_lhs36).Int())
			(_lhs37).ArraySet1(_rhs54, (_lhs38).Int())
			var _311_v98 _dafny.Int
			_ = _311_v98
			var _312_v99 bool
			_ = _312_v99
			var _out2 _dafny.Int
			_ = _out2
			var _out3 bool
			_ = _out3
			_out2, _out3 = Companion_Default___.M1((func() bool {
				if (_196_v15).Contains((_186_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_186_v9), 0))).Int()).(bool)) {
					return (_196_v15).Get((_186_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_186_v9), 0))).Int()).(bool)).(bool)
				}
				return _176_v0
			})(), Companion_Default___.SafeDivisionInt(_177_v1, _dafny.IntOfInt64(140)), _190_globalState)
			_311_v98 = _out2
			_312_v99 = _out3
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_186_v9), 0))
			_ = _index41
			(_186_v9).ArraySet1((Companion_Default___.Fm15(_190_globalState)).IsDisjointFrom(_dafny.SetOf(_311_v98, _177_v1, _177_v1, (_dafny.Zero).Minus(_177_v1))), (_index41).Int())
		}
		var _313_v100 *C0
		_ = _313_v100
		var _nw43 *C0 = New_C0_()
		_ = _nw43
		_nw43.Ctor__()
		_313_v100 = _nw43
		var _314_v101 _dafny.Sequence
		_ = _314_v101
		_314_v101 = _dafny.SeqOf((func() bool {
			if _176_v0 {
				return _176_v0
			}
			return _176_v0
		})(), _176_v0, !(_176_v0))
		var _315_v102 _dafny.Map
		_ = _315_v102
		_315_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v1, _177_v1)
		var _rhs55 _dafny.Int = (_313_v100).Fm4((func() _dafny.Int {
			if (_315_v102).Contains(_177_v1) {
				return (_315_v102).Get(_177_v1).(_dafny.Int)
			}
			return _177_v1
		})(), (_177_v1).Cmp(_dafny.IntOfUint32((_314_v101).Cardinality())) != 0, _190_globalState)
		_ = _rhs55
		var _rhs56 _dafny.Int = (_dafny.Zero).Minus(_177_v1)
		_ = _rhs56
		var _rhs57 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_314_v101, (Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_314_v101).Cardinality()))).Uint32(), !(_195_v14).Contains(_177_v1))
		_ = _rhs57
		var _rhs58 _dafny.Int = _177_v1
		_ = _rhs58
		var _lhs39 *GlobalState = _190_globalState
		_ = _lhs39
		var _lhs40 *GlobalState = _190_globalState
		_ = _lhs40
		var _lhs41 *GlobalState = _190_globalState
		_ = _lhs41
		_lhs39.F10 = _rhs55
		_lhs40.F8 = _rhs56
		_314_v101 = _rhs57
		_lhs41.F19 = _rhs58
		var _316_v103 _dafny.Set
		_ = _316_v103
		_316_v103 = _dafny.SetOf(_313_v100, _313_v100, _313_v100, _313_v100, _313_v100)
		var _317_v104 _dafny.MultiSet
		_ = _317_v104
		_317_v104 = _dafny.MultiSetOf(_316_v103)
		var _318_v105 _dafny.Sequence
		_ = _318_v105
		_318_v105 = _dafny.SeqOf(_316_v103)
		_317_v104 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_318_v105, (Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_318_v105).Cardinality()))).Uint32(), _316_v103))
		(_190_globalState).F8 = _dafny.IntOfUint32((_183_v5).Cardinality())
	} else {
		_183_v5 = _183_v5
		var _319_v106 D3
		_ = _319_v106
		_319_v106 = Companion_D3_.Create_DC10_((_dafny.IntOfUint32((_dafny.SeqOf(_177_v1)).Cardinality())).Cmp(_177_v1) <= 0, (_dafny.Zero).Minus(_177_v1), _176_v0)
		_319_v106 = _319_v106
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_186_v9), 0))
		_ = _index42
		(_186_v9).ArraySet1(_176_v0, (_index42).Int())
		var _320_v107 D2
		_ = _320_v107
		_320_v107 = Companion_D2_.Create_DC8_(_177_v1, _176_v0, _177_v1)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_186_v9), 0))
		_ = _index43
		var _rhs59 bool = (_176_v0) == (_176_v0)
		_ = _rhs59
		var _rhs60 _dafny.Int = Companion_Default___.Fm1((func() bool {
			if _176_v0 {
				return _176_v0
			}
			return _176_v0
		})(), _187_v10, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_183_v5).Cardinality()), _177_v1), _190_globalState)
		_ = _rhs60
		var _rhs61 _dafny.Int = (_320_v107).Dtor_cf12()
		_ = _rhs61
		var _lhs42 _dafny.Array = _186_v9
		_ = _lhs42
		var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_186_v9), 0))
		_ = _lhs43
		var _lhs44 *GlobalState = _190_globalState
		_ = _lhs44
		(_lhs42).ArraySet1(_rhs59, (_lhs43).Int())
		_lhs44.F2 = _rhs60
		_177_v1 = _rhs61
		var _321_v108 _dafny.Map
		_ = _321_v108
		_321_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(685), _180_v4)
		_321_v108 = (_321_v108).Update((_dafny.IntOfUint32((_183_v5).Cardinality())).Minus(_177_v1), _180_v4)
		var _322_v109 _dafny.Set
		_ = _322_v109
		_322_v109 = _dafny.SetOf((_186_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_186_v9), 0))).Int()).(bool))
		var _323_v110 *C1
		_ = _323_v110
		var _nw44 *C1 = New_C1_()
		_ = _nw44
		_nw44.Ctor__(_322_v109, _177_v1)
		_323_v110 = _nw44
	}
	var _324_v111 _dafny.Int
	_ = _324_v111
	var _325_v112 bool
	_ = _325_v112
	var _out4 _dafny.Int
	_ = _out4
	var _out5 bool
	_ = _out5
	_out4, _out5 = Companion_Default___.M1(false, _dafny.IntOfInt64(292), _190_globalState)
	_324_v111 = _out4
	_325_v112 = _out5
	if _176_v0 {
		(_190_globalState).F23 = _176_v0
		var _326_v113 _dafny.MultiSet
		_ = _326_v113
		_326_v113 = _dafny.MultiSetOf(_177_v1, _177_v1)
		_326_v113 = (_326_v113).Union(_dafny.MultiSetOf(_324_v111, _177_v1))
		if (_301_v90).IsProperSubsetOf(_301_v90) {
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_180_v4), 0))
			_ = _index44
			(_180_v4).ArraySet1(_324_v111, (_index44).Int())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_180_v4), 0))
			_ = _index45
			(_180_v4).ArraySet1((_177_v1).Minus((func() _dafny.Int {
				if _176_v0 {
					return _dafny.IntOfInt64(714)
				}
				return (_dafny.Zero).Minus(_324_v111)
			})()), (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_186_v9), 0))
			_ = _index46
			(_186_v9).ArraySet1((Companion_Default___.Fm15(_190_globalState)).IsProperSubsetOf(func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(722), _dafny.IntOfInt64(354))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _327_v114 _dafny.Int
					_327_v114 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(722)).Cmp(_327_v114) <= 0) && ((_327_v114).Cmp(_dafny.IntOfInt64(354)) < 0) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_327_v114, _dafny.IntOfInt64(-455)))
					}
				}
				return _coll7.ToSet()
			}()), (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_186_v9), 0))
			_ = _index47
			(_186_v9).ArraySet1(_325_v112, (_index47).Int())
			_196_v15 = (_196_v15).Update(Companion_Default___.Fm0(_325_v112, _196_v15, _190_globalState), _176_v0)
			var _328_v115 _dafny.Set
			_ = _328_v115
			_328_v115 = _dafny.SetOf(_325_v112)
			var _329_v116 *C1
			_ = _329_v116
			var _nw45 *C1 = New_C1_()
			_ = _nw45
			_nw45.Ctor__(_328_v115, (_328_v115).Cardinality())
			_329_v116 = _nw45
			var _330_v117 _dafny.MultiSet
			_ = _330_v117
			_330_v117 = _dafny.MultiSetOf(_329_v116)
			var _331_v118 *C1
			_ = _331_v118
			var _nw46 *C1 = New_C1_()
			_ = _nw46
			_nw46.Ctor__(_dafny.SetOf(_176_v0), (func() _dafny.Int {
				if (_330_v117).Contains(_329_v116) {
					return (_330_v117).Multiplicity(_329_v116)
				}
				return _177_v1
			})())
			_331_v118 = _nw46
			(_190_globalState).F2 = _329_v116.F27
		} else {
			(_190_globalState).F23 = _176_v0
			_186_v9 = _186_v9
			var _332_v119 *C0
			_ = _332_v119
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__()
			_332_v119 = _nw47
			_186_v9 = _186_v9
			var _333_v120 _dafny.Int
			_ = _333_v120
			var _334_v121 bool
			_ = _334_v121
			var _out6 _dafny.Int
			_ = _out6
			var _out7 bool
			_ = _out7
			_out6, _out7 = Companion_Default___.M1(_176_v0, (_dafny.IntOfInt64(-429)).Plus(_324_v111), _190_globalState)
			_333_v120 = _out6
			_334_v121 = _out7
		}
		(_190_globalState).F23 = _325_v112
		var _335_v122 _dafny.Int
		_ = _335_v122
		var _336_v123 bool
		_ = _336_v123
		var _out8 _dafny.Int
		_ = _out8
		var _out9 bool
		_ = _out9
		_out8, _out9 = Companion_Default___.M1(_176_v0, (_dafny.MultiSetOf(_324_v111)).Cardinality(), _190_globalState)
		_335_v122 = _out8
		_336_v123 = _out9
	} else {
		var _337_v124 D0
		_ = _337_v124
		_337_v124 = Companion_D0_.Create_DC1_(_177_v1)
		(_190_globalState).F5 = Companion_Default___.Fm1(_325_v112, _187_v10, (_337_v124).Dtor_cf1(), _190_globalState)
		_195_v14 = (_195_v14).Update(_dafny.IntOfInt64(438), _176_v0)
		var _338_v125 _dafny.Map
		_ = _338_v125
		_338_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v4, (_dafny.Zero).Minus(((_185_v8).Intersection(_185_v8)).Cardinality()))
		_338_v125 = (_338_v125).Update(_180_v4, _324_v111)
		_186_v9 = _186_v9
		var _339_v126 _dafny.Int
		_ = _339_v126
		var _340_v127 bool
		_ = _340_v127
		var _out10 _dafny.Int
		_ = _out10
		var _out11 bool
		_ = _out11
		_out10, _out11 = Companion_Default___.M1(_325_v112, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _176_v0 {
				return Companion_Default___.Fm10(_177_v1, _dafny.IntOfInt64(288), _176_v0, _176_v0, _190_globalState)
			}
			return _183_v5
		})()).Cardinality()), _190_globalState)
		_339_v126 = _out10
		_340_v127 = _out11
	}
	var _341_v128 _dafny.Sequence
	_ = _341_v128
	_341_v128 = _dafny.SeqOf(_dafny.MultiSetOf(_176_v0, _325_v112, _176_v0, _325_v112, _325_v112))
	var _342_v129 _dafny.Set
	_ = _342_v129
	_342_v129 = _dafny.SetOf(_185_v8, _185_v8, _185_v8, (_341_v128).Select((Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_341_v128).Cardinality()))).Uint32()).(_dafny.MultiSet), _185_v8)
	var _343_v130 _dafny.Map
	_ = _343_v130
	_343_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v8, (_324_v111).Cmp((_342_v129).Cardinality()) == 0)
	_343_v130 = (_343_v130).Update(_185_v8, _176_v0)
	_183_v5 = _183_v5
	var _344_v131 D3
	_ = _344_v131
	_344_v131 = Companion_D3_.Create_DC10_(_176_v0, _177_v1, false)
	var _345_v132 _dafny.Sequence
	_ = _345_v132
	_345_v132 = _dafny.SeqOf(true)
	var _346_v133 _dafny.Int
	_ = _346_v133
	var _347_v134 bool
	_ = _347_v134
	var _out12 _dafny.Int
	_ = _out12
	var _out13 bool
	_ = _out13
	_out12, _out13 = Companion_Default___.M1(!((_344_v131).Dtor_cf16()), ((_dafny.MultiSetFromSeq(_345_v132)).Intersection(_185_v8)).Cardinality(), _190_globalState)
	_346_v133 = _out12
	_347_v134 = _out13
	_347_v134 = _347_v134
	var _348_v135 _dafny.MultiSet
	_ = _348_v135
	_348_v135 = _dafny.MultiSetOf(_dafny.IntOfInt64(60))
	var _349_v136 _dafny.Map
	_ = _349_v136
	_349_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_348_v135, _346_v133)
	var _rhs62 _dafny.Map = _349_v136
	_ = _rhs62
	var _rhs63 _dafny.Int = Companion_Default___.SafeModuloInt((_177_v1).Times(_324_v111), (_324_v111).Plus(_324_v111))
	_ = _rhs63
	var _rhs64 _dafny.Int = _346_v133
	_ = _rhs64
	var _rhs65 _dafny.Int = (_dafny.Zero).Minus(_177_v1)
	_ = _rhs65
	var _lhs45 *GlobalState = _190_globalState
	_ = _lhs45
	var _lhs46 *GlobalState = _190_globalState
	_ = _lhs46
	_349_v136 = _rhs62
	_324_v111 = _rhs63
	_lhs45.F19 = _rhs64
	_lhs46.F5 = _rhs65
	if _347_v134 {
		var _350_v137 _dafny.Set
		_ = _350_v137
		_350_v137 = _dafny.SetOf(_176_v0)
		var _351_v138 *C1
		_ = _351_v138
		var _nw48 *C1 = New_C1_()
		_ = _nw48
		_nw48.Ctor__((_350_v137).Union(_350_v137), _346_v133)
		_351_v138 = _nw48
		(_190_globalState).F25 = Companion_Default___.Fm0(_325_v112, (_196_v15).Merge(_196_v15), _190_globalState)
		var _352_v139 _dafny.Array
		_ = _352_v139
		var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
		_ = _nw49
		_352_v139 = _nw49
		var _353_v140 _dafny.Array
		_ = _353_v140
		var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
		_ = _nw50
		_353_v140 = _nw50
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_352_v139), 0))
		_ = _index48
		(_352_v139).ArraySet1(_353_v140, (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_352_v139), 0))
		_ = _index49
		(_352_v139).ArraySet1(_353_v140, (_index49).Int())
		var _354_v141 _dafny.Int
		_ = _354_v141
		var _355_v142 _dafny.Int
		_ = _355_v142
		var _out14 _dafny.Int
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		_out14, _out15 = (_351_v138).M0((func() bool {
			if !(_347_v134) {
				return _176_v0
			}
			return _176_v0
		})(), _190_globalState)
		_354_v141 = _out14
		_355_v142 = _out15
		(_190_globalState).F23 = _347_v134
	} else {
		(_190_globalState).F5 = _346_v133
		var _356_v143 D2
		_ = _356_v143
		_356_v143 = Companion_D2_.Create_DC8_(_346_v133, _325_v112, _dafny.IntOfUint32((_183_v5).Cardinality()))
		var _357_v144 _dafny.Map
		_ = _357_v144
		_357_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v112, _177_v1)
		var _rhs66 _dafny.Int = _346_v133
		_ = _rhs66
		var _rhs67 _dafny.Sequence = _dafny.SeqOf(_347_v134, _325_v112, (func() bool {
			if Companion_Default___.Fm0(_347_v134, _196_v15, _190_globalState) {
				return _347_v134
			}
			return !(_347_v134)
		})(), (func() bool {
			if (_356_v143).Dtor_cf13() {
				return _325_v112
			}
			return _325_v112
		})(), (_346_v133).Cmp(Companion_Default___.Fm1(_325_v112, _dafny.CodePoint('w'), (_357_v144).Cardinality(), _190_globalState)) < 0)
		_ = _rhs67
		var _rhs68 bool = _176_v0
		_ = _rhs68
		var _lhs47 *GlobalState = _190_globalState
		_ = _lhs47
		_lhs47.F5 = _rhs66
		_345_v132 = _rhs67
		_176_v0 = _rhs68
		var _358_v145 _dafny.Set
		_ = _358_v145
		_358_v145 = _dafny.SetOf(false, true)
		var _359_v146 T0
		_ = _359_v146
		var _nw51 *C1 = New_C1_()
		_ = _nw51
		_nw51.Ctor__(_358_v145, _177_v1)
		_359_v146 = _nw51
		var _pat_let_tv10 = _183_v5
		_ = _pat_let_tv10
		var _360_v147 D3
		_ = _360_v147
		_360_v147 = Companion_D3_.Create_DC12_(_325_v112, func(_pat_let6_0 D0) D0 {
			return func(_361_dt__update__tmp_h4 D0) D0 {
				return func(_pat_let7_0 _dafny.Int) D0 {
					return func(_362_dt__update_hcf1_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_(_362_dt__update_hcf1_h0)
					}(_pat_let7_0)
				}(_dafny.IntOfUint32((_pat_let_tv10).Cardinality()))
			}(_pat_let6_0)
		}(Companion_D0_.Create_DC1_(_346_v133)), (_344_v131).Dtor_cf18(), false, _359_v146)
		_325_v112 = (_360_v147).Dtor_cf22()
		var _363_v149 _dafny.Map
		_ = _363_v149
		_363_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_177_v1, _346_v133)), (func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(967), _dafny.IntOfInt64(310))); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _364_v148 _dafny.Int
				_364_v148 = interface{}(_compr_8).(_dafny.Int)
				if ((_dafny.IntOfInt64(967)).Cmp(_364_v148) <= 0) && ((_364_v148).Cmp(_dafny.IntOfInt64(310)) < 0) {
					_coll8.Add((_364_v148).Plus(_346_v133), _347_v134)
				}
			}
			return _coll8.ToMap()
		}()).Cardinality())
		var _365_v150 _dafny.Sequence
		_ = _365_v150
		_365_v150 = _dafny.SeqOf(_183_v5)
		_363_v149 = (_363_v149).Update(_177_v1, _dafny.IntOfUint32((_365_v150).Cardinality()))
		var _366_v151 *C1
		_ = _366_v151
		var _nw52 *C1 = New_C1_()
		_ = _nw52
		_nw52.Ctor__((func() _dafny.Set {
			if true {
				return _358_v145
			}
			return _dafny.SetOf(_347_v134, _176_v0)
		})(), _177_v1)
		_366_v151 = _nw52
	}
	var _367_v152 D0
	_ = _367_v152
	_367_v152 = Companion_D0_.Create_DC0_(Companion_Default___.Fm16(_177_v1, _190_globalState))
	var _pat_let_tv11 = _187_v10
	_ = _pat_let_tv11
	var _pat_let_tv12 = _187_v10
	_ = _pat_let_tv12
	var _pat_let_tv13 = _187_v10
	_ = _pat_let_tv13
	var _pat_let_tv14 = _325_v112
	_ = _pat_let_tv14
	var _pat_let_tv15 = _187_v10
	_ = _pat_let_tv15
	var _pat_let_tv16 = _187_v10
	_ = _pat_let_tv16
	var _source3 _dafny.CodePoint = func(_source4 D0) _dafny.CodePoint {
		if _source4.Is_DC1() {
			var _368___mcc_h1 _dafny.Int = _source4.Get_().(D0_DC1).Cf1
			_ = _368___mcc_h1
			var _369_cf1 _dafny.Int = _368___mcc_h1
			_ = _369_cf1
			return _pat_let_tv11
		} else if _source4.Is_DC2() {
			var _370___mcc_h2 bool = _source4.Get_().(D0_DC2).Cf2
			_ = _370___mcc_h2
			var _371___mcc_h3 _dafny.Set = _source4.Get_().(D0_DC2).Cf3
			_ = _371___mcc_h3
			var _372___mcc_h4 bool = _source4.Get_().(D0_DC2).Cf4
			_ = _372___mcc_h4
			var _373___mcc_h5 bool = _source4.Get_().(D0_DC2).Cf5
			_ = _373___mcc_h5
			var _374___mcc_h6 _dafny.Int = _source4.Get_().(D0_DC2).Cf6
			_ = _374___mcc_h6
			var _375_cf6 _dafny.Int = _374___mcc_h6
			_ = _375_cf6
			var _376_cf5 bool = _373___mcc_h5
			_ = _376_cf5
			var _377_cf4 bool = _372___mcc_h4
			_ = _377_cf4
			var _378_cf3 _dafny.Set = _371___mcc_h3
			_ = _378_cf3
			var _379_cf2 bool = _370___mcc_h2
			_ = _379_cf2
			return _pat_let_tv12
		} else if _source4.Is_DC3() {
			var _380___mcc_h7 _dafny.Sequence = _source4.Get_().(D0_DC3).Cf7
			_ = _380___mcc_h7
			var _381_cf7 _dafny.Sequence = _380___mcc_h7
			_ = _381_cf7
			return _pat_let_tv13
		} else {
			var _382___mcc_h8 _dafny.Sequence = _source4.Get_().(D0_DC0).Cf0
			_ = _382___mcc_h8
			var _383_cf0 _dafny.Sequence = _382___mcc_h8
			_ = _383_cf0
			if _pat_let_tv14 {
				return _pat_let_tv15
			} else {
				return _pat_let_tv16
			}
		}
	}(_367_v152)
	_ = _source3
	var _384___mcc_h0 _dafny.CodePoint = _source3
	_ = _384___mcc_h0
	var _385_cf8 _dafny.CodePoint = _384___mcc_h0
	_ = _385_cf8
	var _386_v153 _dafny.Map
	_ = _386_v153
	_386_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _325_v112)
	_386_v153 = (_386_v153).Update(_385_cf8, !(true) || (_176_v0))
	var _387_v154 _dafny.Set
	_ = _387_v154
	_387_v154 = _dafny.SetOf((Companion_D2_.Create_DC7_(_347_v134)).Dtor_cf11())
	var _388_v155 *C1
	_ = _388_v155
	var _nw53 *C1 = New_C1_()
	_ = _nw53
	_nw53.Ctor__(_387_v154, _177_v1)
	_388_v155 = _nw53
	var _389_v156 _dafny.Int
	_ = _389_v156
	var _390_v157 bool
	_ = _390_v157
	var _out16 _dafny.Int
	_ = _out16
	var _out17 bool
	_ = _out17
	_out16, _out17 = Companion_Default___.M1(_325_v112, _324_v111, _190_globalState)
	_389_v156 = _out16
	_390_v157 = _out17
	(_190_globalState).F5 = _346_v133
	if true {
		var _391_v158 D2
		_ = _391_v158
		_391_v158 = Companion_D2_.Create_DC5_(_183_v5)
		var _392_v159 _dafny.Sequence
		_ = _392_v159
		_392_v159 = _dafny.SeqOf(_183_v5, _183_v5, _183_v5)
		var _393_v160 _dafny.Array
		_ = _393_v160
		var _nwElement0_7 _dafny.Sequence = _183_v5
		_ = _nwElement0_7
		var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(12))
		_ = _nw54
		(_nw54).ArraySet1(_nwElement0_7, 0)
		(_nw54).ArraySet1(_183_v5, 1)
		(_nw54).ArraySet1(_183_v5, 2)
		(_nw54).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(929))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_394_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_395_i15 _dafny.Int) _dafny.CodePoint {
				return _394_v10
			}
		})(_187_v10))), 3)
		(_nw54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), 4)
		(_nw54).ArraySet1(_183_v5, 5)
		(_nw54).ArraySet1(_183_v5, 6)
		(_nw54).ArraySet1(_183_v5, 7)
		(_nw54).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_391_v158).Dtor_cf9(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_396_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_397_i16 _dafny.Int) _dafny.CodePoint {
				return _396_v10
			}
		})(_187_v10)))), (Companion_Default___.SafeIndex(_346_v133, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_391_v158).Dtor_cf9(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_398_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_399_i16 _dafny.Int) _dafny.CodePoint {
				return _398_v10
			}
		})(_187_v10))))).Cardinality()))).Uint32(), _dafny.CodePoint('w')), 8)
		(_nw54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), 9)
		(_nw54).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_392_v159).Select((Companion_Default___.SafeIndex(_324_v111, _dafny.IntOfUint32((_392_v159).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_178_v2).Cardinality()), _dafny.IntOfUint32(((_392_v159).Select((Companion_Default___.SafeIndex(_324_v111, _dafny.IntOfUint32((_392_v159).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _187_v10), (Companion_Default___.SafeIndex(_177_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_392_v159).Select((Companion_Default___.SafeIndex(_324_v111, _dafny.IntOfUint32((_392_v159).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_178_v2).Cardinality()), _dafny.IntOfUint32(((_392_v159).Select((Companion_Default___.SafeIndex(_324_v111, _dafny.IntOfUint32((_392_v159).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _187_v10)).Cardinality()))).Uint32(), _187_v10), 10)
		(_nw54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), 11)
		_393_v160 = _nw54
		_393_v160 = _393_v160
		var _400_v161 _dafny.Array
		_ = _400_v161
		var _nw55 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
		_ = _nw55
		_400_v161 = _nw55
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_186_v9), 0))
		_ = _index50
		(_186_v9).ArraySet1(_347_v134, (_index50).Int())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_186_v9), 0))
		_ = _index51
		var _rhs69 _dafny.Array = _400_v161
		_ = _rhs69
		var _rhs70 bool = true
		_ = _rhs70
		var _rhs71 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_183_v5, (Companion_Default___.SafeIndex(_346_v133, _dafny.IntOfUint32((_183_v5).Cardinality()))).Uint32(), _187_v10)
		_ = _rhs71
		var _rhs72 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), _183_v5)
		_ = _rhs72
		var _lhs48 _dafny.Array = _186_v9
		_ = _lhs48
		var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_186_v9), 0))
		_ = _lhs49
		_400_v161 = _rhs69
		(_lhs48).ArraySet1(_rhs70, (_lhs49).Int())
		_183_v5 = _rhs71
		_183_v5 = _rhs72
		var _401_v162 _dafny.Array
		_ = _401_v162
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_6
		var _nw56 _dafny.Array
		_ = _nw56
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw56 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Sequence = (func(_402_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_403_i17 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_402_v2, _402_v2)
				}
			})(_178_v2)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw56 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw56).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw56).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_401_v162 = _nw56
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_401_v162), 0))
		_ = _index52
		(_401_v162).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(70))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_404_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_405_i18 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_404_v1)
			}
		})(_177_v1))), (_index52).Int())
		var _406_v163 _dafny.Set
		_ = _406_v163
		_406_v163 = _dafny.SetOf(_347_v134, _347_v134)
		var _407_v164 *C1
		_ = _407_v164
		var _nw57 *C1 = New_C1_()
		_ = _nw57
		_nw57.Ctor__(_406_v163, _177_v1)
		_407_v164 = _nw57
		var _408_v165 _dafny.Map
		_ = _408_v165
		_408_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_v111, _407_v164)
		var _409_v166 _dafny.Sequence
		_ = _409_v166
		_409_v166 = _dafny.SeqOf(_408_v165)
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_401_v162), 0))
		_ = _index53
		(_401_v162).ArraySet1(_dafny.SeqOf(_324_v111, _dafny.IntOfUint32((_409_v166).Cardinality())), (_index53).Int())
		_407_v164 = _407_v164
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), _183_v5) {
			var _410_v167 _dafny.Int
			_ = _410_v167
			var _411_v168 bool
			_ = _411_v168
			var _out18 _dafny.Int
			_ = _out18
			var _out19 bool
			_ = _out19
			_out18, _out19 = Companion_Default___.M1((_347_v134) || (_176_v0), _177_v1, _190_globalState)
			_410_v167 = _out18
			_411_v168 = _out19
			_347_v134 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate((_401_v162).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_401_v162), 0))).Int()).(_dafny.Sequence), _178_v2), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_410_v167, _dafny.IntOfInt64(270)), _178_v2), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_345_v132).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_410_v167, _dafny.IntOfInt64(270)), _178_v2)).Cardinality()))).Uint32(), _324_v111))
			var _412_v169 T0
			_ = _412_v169
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__()
			_412_v169 = _nw58
			var _413_v170 _dafny.Sequence
			_ = _413_v170
			_413_v170 = _dafny.SeqOf(_412_v169)
			_413_v170 = _413_v170
			(_190_globalState).F19 = (_dafny.Zero).Minus((_177_v1).Minus(_dafny.IntOfInt64(254)))
			(_190_globalState).F25 = _176_v0
		} else {
			var _414_v171 _dafny.Map
			_ = _414_v171
			_414_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v0, _dafny.Companion_Sequence_.Update(_345_v132, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.IntOfUint32((_345_v132).Cardinality()))).Uint32(), _325_v112))
			var _415_v172 _dafny.Map
			_ = _415_v172
			_415_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_178_v2).Cardinality()), (_414_v171).Merge(_414_v171))
			_415_v172 = (_415_v172).Update(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_v111, !((_186_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_186_v9), 0))).Int()).(bool)))).Update(_177_v1, _347_v134)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_347_v134, _345_v132))
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_186_v9), 0))
			_ = _index54
			(_186_v9).ArraySet1(!((_407_v164.F27).Cmp(_346_v133) <= 0), (_index54).Int())
			var _416_v173 _dafny.Array
			_ = _416_v173
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
			_ = _nw59
			_416_v173 = _nw59
			_416_v173 = _416_v173
			var _417_v174 *C0
			_ = _417_v174
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__()
			_417_v174 = _nw60
			var _418_v175 _dafny.Sequence
			_ = _418_v175
			_418_v175 = _dafny.SeqOf(_417_v174)
			var _419_v176 _dafny.Array
			_ = _419_v176
			var _nwElement0_8 *C0 = _417_v174
			_ = _nwElement0_8
			var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
			_ = _nw61
			(_nw61).ArraySet1(_nwElement0_8, 0)
			(_nw61).ArraySet1(_417_v174, 1)
			(_nw61).ArraySet1(_417_v174, 2)
			(_nw61).ArraySet1(_417_v174, 3)
			(_nw61).ArraySet1(_417_v174, 4)
			(_nw61).ArraySet1(_417_v174, 5)
			(_nw61).ArraySet1(_417_v174, 6)
			(_nw61).ArraySet1(_417_v174, 7)
			(_nw61).ArraySet1(_417_v174, 8)
			(_nw61).ArraySet1(_417_v174, 9)
			(_nw61).ArraySet1(_417_v174, 10)
			(_nw61).ArraySet1(_417_v174, 11)
			(_nw61).ArraySet1(_417_v174, 12)
			(_nw61).ArraySet1(_417_v174, 13)
			(_nw61).ArraySet1(_417_v174, 14)
			(_nw61).ArraySet1(_417_v174, 15)
			(_nw61).ArraySet1((func() *C0 {
				if _347_v134 {
					return _417_v174
				}
				return _417_v174
			})(), 16)
			(_nw61).ArraySet1(_417_v174, 17)
			(_nw61).ArraySet1((_418_v175).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_345_v132).Cardinality()), _dafny.IntOfUint32((_418_v175).Cardinality()))).Uint32()).(*C0), 18)
			(_nw61).ArraySet1(_417_v174, 19)
			(_nw61).ArraySet1(_417_v174, 20)
			(_nw61).ArraySet1(_417_v174, 21)
			(_nw61).ArraySet1(_417_v174, 22)
			(_nw61).ArraySet1(_417_v174, 23)
			(_nw61).ArraySet1(_417_v174, 24)
			_419_v176 = _nw61
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_419_v176), 0))
			_ = _index55
			(_419_v176).ArraySet1(_417_v174, (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_419_v176), 0))
			_ = _index56
			var _rhs73 _dafny.Int = _324_v111
			_ = _rhs73
			var _rhs74 _dafny.Int = (_346_v133).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_418_v175).Cardinality()), _dafny.IntOfUint32((_178_v2).Cardinality()), _dafny.IntOfInt64(-511), _407_v164.F27, _324_v111)).Cardinality()))
			_ = _rhs74
			var _rhs75 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(563))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_420_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_421_i19 _dafny.Int) _dafny.CodePoint {
					return _420_v10
				}
			})(_187_v10)))
			_ = _rhs75
			var _rhs76 *C0 = _417_v174
			_ = _rhs76
			var _rhs77 bool = _347_v134
			_ = _rhs77
			var _lhs50 *GlobalState = _190_globalState
			_ = _lhs50
			var _lhs51 *C1 = _407_v164
			_ = _lhs51
			var _lhs52 _dafny.Array = _419_v176
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_419_v176), 0))
			_ = _lhs53
			var _lhs54 *GlobalState = _190_globalState
			_ = _lhs54
			_lhs50.F8 = _rhs73
			_lhs51.F27 = _rhs74
			_183_v5 = _rhs75
			(_lhs52).ArraySet1(_rhs76, (_lhs53).Int())
			_lhs54.F25 = _rhs77
			var _422_v177 _dafny.Int
			_ = _422_v177
			var _423_v178 bool
			_ = _423_v178
			var _out20 _dafny.Int
			_ = _out20
			var _out21 bool
			_ = _out21
			_out20, _out21 = Companion_Default___.M1(!((_185_v8).IsProperSubsetOf(_185_v8)), (_177_v1).Plus(_407_v164.F27), _190_globalState)
			_422_v177 = _out20
			_423_v178 = _out21
		}
	} else {
		var _424_v179 _dafny.Set
		_ = _424_v179
		_424_v179 = _dafny.SetOf(!(true), _176_v0, _176_v0)
		var _425_v180 *C1
		_ = _425_v180
		var _nw62 *C1 = New_C1_()
		_ = _nw62
		_nw62.Ctor__(_424_v179, _324_v111)
		_425_v180 = _nw62
		var _426_v181 _dafny.MultiSet
		_ = _426_v181
		_426_v181 = _dafny.MultiSetOf(_425_v180, _425_v180)
		if ((_426_v181).Difference(_426_v181)).IsDisjointFrom(_426_v181) {
			var _427_v182 _dafny.Sequence
			_ = _427_v182
			_427_v182 = _dafny.SeqOf((_348_v135).Intersection(_348_v135))
			_427_v182 = _427_v182
			var _428_v183 *C0
			_ = _428_v183
			var _nw63 *C0 = New_C0_()
			_ = _nw63
			_nw63.Ctor__()
			_428_v183 = _nw63
			_186_v9 = _186_v9
			var _429_v184 _dafny.Map
			_ = _429_v184
			_429_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_183_v5, (Companion_Default___.SafeIndex(_425_v180.F27, _dafny.IntOfUint32((_183_v5).Cardinality()))).Uint32(), _187_v10), _176_v0)
			(_190_globalState).F22 = (func() bool {
				if (_429_v184).Contains(_183_v5) {
					return (_429_v184).Get(_183_v5).(bool)
				}
				return _325_v112
			})()
			var _430_v185 *C0
			_ = _430_v185
			var _nw64 *C0 = New_C0_()
			_ = _nw64
			_nw64.Ctor__()
			_430_v185 = _nw64
		} else {
			var _431_v186 _dafny.Map
			_ = _431_v186
			_431_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(625), (func() _dafny.Int {
				if (_348_v135).Contains(((_425_v180).F26()).Cardinality()) {
					return (_348_v135).Multiplicity(((_425_v180).F26()).Cardinality())
				}
				return _324_v111
			})()), _dafny.Companion_Sequence_.Update(_183_v5, (Companion_Default___.SafeIndex(_324_v111, _dafny.IntOfUint32((_183_v5).Cardinality()))).Uint32(), _187_v10))
			_431_v186 = _431_v186
			var _432_v187 _dafny.Array
			_ = _432_v187
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_7
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.MultiSet = (func(_433_v8 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_434_i20 _dafny.Int) _dafny.MultiSet {
						return _433_v8
					}
				})(_185_v8)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw65 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw65).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw65).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_432_v187 = _nw65
			var _435_v188 _dafny.Map
			_ = _435_v188
			_435_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_432_v187, _425_v180.F27)
			_435_v188 = (_435_v188).Update(_432_v187, (_324_v111).Plus(_324_v111))
			_183_v5 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_183_v5, _183_v5), _183_v5)
			var _436_v189 D2
			_ = _436_v189
			_436_v189 = Companion_D2_.Create_DC6_((_425_v180.F27).Cmp(_177_v1) == 0)
			var _rhs78 bool = _176_v0
			_ = _rhs78
			var _rhs79 D2 = Companion_Default___.Fm17(_190_globalState)
			_ = _rhs79
			var _lhs55 *GlobalState = _190_globalState
			_ = _lhs55
			_lhs55.F23 = _rhs78
			_436_v189 = _rhs79
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_180_v4), 0))
			_ = _index57
			(_180_v4).ArraySet1(_346_v133, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_180_v4), 0))
			_ = _index58
			(_180_v4).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm1(true, _187_v10, _346_v133, _190_globalState)), (_index58).Int())
		}
		var _rhs80 _dafny.Int = _177_v1
		_ = _rhs80
		var _rhs81 _dafny.Int = (_324_v111).Times(_324_v111)
		_ = _rhs81
		var _lhs56 *GlobalState = _190_globalState
		_ = _lhs56
		var _lhs57 *C1 = _425_v180
		_ = _lhs57
		_lhs56.F8 = _rhs80
		_lhs57.F27 = _rhs81
		var _437_v190 _dafny.Map
		_ = _437_v190
		_437_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_425_v180.F27).Plus(_346_v133))
		var _438_v191 _dafny.Map
		_ = _438_v191
		_438_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_176_v0), _324_v111)
		var _439_v192 _dafny.Sequence
		_ = _439_v192
		_439_v192 = _dafny.SeqOf(_425_v180, _425_v180)
		_437_v190 = Companion_Default___.Fm18((func() _dafny.Int {
			if (_438_v191).Contains(_176_v0) {
				return (_438_v191).Get(_176_v0).(_dafny.Int)
			}
			return _425_v180.F27
		})(), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_425_v180), _439_v192), _190_globalState)
		var _440_v193 D0
		_ = _440_v193
		_440_v193 = Companion_D0_.Create_DC1_((_195_v14).Cardinality())
		if !(!((_dafny.IntOfUint32((Companion_Default___.Fm12(_dafny.IntOfUint32((_183_v5).Cardinality()), _440_v193, _190_globalState)).Cardinality())).Cmp((_196_v15).Cardinality()) >= 0)) || (_347_v134) {
			var _441_v194 D2
			_ = _441_v194
			_441_v194 = Companion_D2_.Create_DC5_(_183_v5)
			var _442_v195 D4
			_ = _442_v195
			_442_v195 = Companion_D4_.Create_DC14_(_348_v135, _425_v180.F27, _441_v194)
			(_190_globalState).F23 = (_324_v111).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(605), (_442_v195).Dtor_cf29()))) == 0
			var _443_v196 _dafny.Int
			_ = _443_v196
			var _444_v197 bool
			_ = _444_v197
			var _out22 _dafny.Int
			_ = _out22
			var _out23 bool
			_ = _out23
			_out22, _out23 = Companion_Default___.M1(_347_v134, _425_v180.F27, _190_globalState)
			_443_v196 = _out22
			_444_v197 = _out23
			var _445_v198 T0
			_ = _445_v198
			var _nw66 *C0 = New_C0_()
			_ = _nw66
			_nw66.Ctor__()
			_445_v198 = _nw66
			var _rhs82 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_348_v135).Contains(_346_v133) {
					return (_348_v135).Multiplicity(_346_v133)
				}
				return _425_v180.F27
			})(), _324_v111)))
			_ = _rhs82
			var _rhs83 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rfeampgsi"), _183_v5), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(573))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_446_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_447_i21 _dafny.Int) _dafny.CodePoint {
					return _446_v10
				}
			})(_187_v10))), (_441_v194).Dtor_cf9()))
			_ = _rhs83
			var _rhs84 _dafny.Int = (_177_v1).Plus(_dafny.IntOfInt64(-503))
			_ = _rhs84
			var _rhs85 T0 = _445_v198
			_ = _rhs85
			var _rhs86 T0 = _445_v198
			_ = _rhs86
			var _lhs58 *GlobalState = _190_globalState
			_ = _lhs58
			_177_v1 = _rhs82
			_183_v5 = _rhs83
			_lhs58.F19 = _rhs84
			_445_v198 = _rhs85
			_445_v198 = _rhs86
			(_190_globalState).F22 = (_425_v180).Fm6(_dafny.SeqOf(_367_v152), _dafny.Companion_Sequence_.Concatenate(_183_v5, _dafny.UnicodeSeqOfUtf8Bytes("u")), _185_v8, _190_globalState)
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_186_v9), 0))
			_ = _index59
			(_186_v9).ArraySet1(!(!_dafny.Companion_Sequence_.Contains(_183_v5, _187_v10)), (_index59).Int())
			var _448_v199 _dafny.Map
			_ = _448_v199
			_448_v199 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_348_v135, _178_v2)
			var _449_v200 D2
			_ = _449_v200
			_449_v200 = Companion_D2_.Create_DC8_(_324_v111, _176_v0, _346_v133)
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_186_v9), 0))
			_ = _index60
			var _rhs87 _dafny.CodePoint = _187_v10
			_ = _rhs87
			var _rhs88 _dafny.Int = _324_v111
			_ = _rhs88
			var _rhs89 _dafny.Int = ((func() _dafny.Map {
				if _325_v112 {
					return _448_v199
				}
				return _448_v199
			})()).Cardinality()
			_ = _rhs89
			var _rhs90 bool = (_449_v200).Dtor_cf13()
			_ = _rhs90
			var _rhs91 _dafny.Int = _177_v1
			_ = _rhs91
			var _lhs59 *GlobalState = _190_globalState
			_ = _lhs59
			var _lhs60 _dafny.Array = _186_v9
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_186_v9), 0))
			_ = _lhs61
			var _lhs62 *GlobalState = _190_globalState
			_ = _lhs62
			_187_v10 = _rhs87
			_324_v111 = _rhs88
			_lhs59.F2 = _rhs89
			(_lhs60).ArraySet1(_rhs90, (_lhs61).Int())
			_lhs62.F10 = _rhs91
		} else {
			var _rhs92 *C1 = _425_v180
			_ = _rhs92
			var _rhs93 _dafny.Sequence = _341_v128
			_ = _rhs93
			_425_v180 = _rhs92
			_341_v128 = _rhs93
			var _450_v201 D2
			_ = _450_v201
			_450_v201 = Companion_D2_.Create_DC8_(_346_v133, _325_v112, _425_v180.F27)
			var _451_v202 _dafny.Int
			_ = _451_v202
			var _452_v203 bool
			_ = _452_v203
			var _out24 _dafny.Int
			_ = _out24
			var _out25 bool
			_ = _out25
			_out24, _out25 = Companion_Default___.M1(false, (_450_v201).Dtor_cf12(), _190_globalState)
			_451_v202 = _out24
			_452_v203 = _out25
			(_190_globalState).F2 = _dafny.IntOfUint32((_183_v5).Cardinality())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_186_v9), 0))
			_ = _index61
			(_186_v9).ArraySet1(_176_v0, (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_186_v9), 0))
			_ = _index62
			(_186_v9).ArraySet1(_452_v203, (_index62).Int())
			_195_v14 = (_195_v14).Merge(_195_v14)
		}
		var _453_v204 _dafny.MultiSet
		_ = _453_v204
		_453_v204 = _dafny.MultiSetOf(_367_v152)
		var _454_v205 _dafny.Map
		_ = _454_v205
		_454_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_v180, _453_v204)
		var _455_v206 _dafny.Sequence
		_ = _455_v206
		_455_v206 = _dafny.SeqOf(((_dafny.MultiSetFromSeq(_178_v2)).Update(_425_v180.F27, Companion_Default___.Abs(_324_v111))).Update(_425_v180.F27, Companion_Default___.Abs(_dafny.IntOfInt64(154))), _348_v135)
		var _pat_let_tv17 = _455_v206
		_ = _pat_let_tv17
		var _456_v207 _dafny.MultiSet
		_ = _456_v207
		_456_v207 = _dafny.MultiSetOf(func(_pat_let8_0 D0) D0 {
			return func(_457_dt__update__tmp_h5 D0) D0 {
				return func(_pat_let9_0 _dafny.Sequence) D0 {
					return func(_458_dt__update_hcf0_h0 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_458_dt__update_hcf0_h0)
					}(_pat_let9_0)
				}(_pat_let_tv17)
			}(_pat_let8_0)
		}(_367_v152), _367_v152)
		var _459_v208 _dafny.Sequence
		_ = _459_v208
		_459_v208 = _dafny.SeqOf(Companion_Default___.Fm19(_190_globalState), (_456_v207))
		_454_v205 = (_454_v205).Update(_425_v180, (_459_v208).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xbifwaa")).Cardinality()), _dafny.IntOfUint32((_459_v208).Cardinality()))).Uint32()).(_dafny.MultiSet))
	}
	var _460_v209 T0
	_ = _460_v209
	var _nw67 *C0 = New_C0_()
	_ = _nw67
	_nw67.Ctor__()
	_460_v209 = _nw67
	_460_v209 = _460_v209
	var _461_v210 D3
	_ = _461_v210
	_461_v210 = Companion_D3_.Create_DC11_(_dafny.CodePoint('r'), false, _176_v0)
	var _hi1 _dafny.Int = _177_v1
	_ = _hi1
	for _462_i22 := (Companion_Default___.Fm1(!((_461_v210).Dtor_cf20()), _187_v10, (_dafny.MultiSetOf(_346_v133)).Cardinality(), _190_globalState)).Minus(_324_v111); _462_i22.Cmp(_hi1) < 0; _462_i22 = _462_i22.Plus(_dafny.One) {
		var _463_v211 *C0
		_ = _463_v211
		var _nw68 *C0 = New_C0_()
		_ = _nw68
		_nw68.Ctor__()
		_463_v211 = _nw68
		_463_v211 = _463_v211
		var _464_v212 _dafny.Int
		_ = _464_v212
		var _465_v213 bool
		_ = _465_v213
		var _out26 _dafny.Int
		_ = _out26
		var _out27 bool
		_ = _out27
		_out26, _out27 = Companion_Default___.M1(_176_v0, (_344_v131).Dtor_cf17(), _190_globalState)
		_464_v212 = _out26
		_465_v213 = _out27
		var _466_v214 _dafny.Set
		_ = _466_v214
		_466_v214 = _dafny.SetOf(_465_v213)
		var _467_v215 *C1
		_ = _467_v215
		var _nw69 *C1 = New_C1_()
		_ = _nw69
		_nw69.Ctor__(_466_v214, _dafny.IntOfUint32((_183_v5).Cardinality()))
		_467_v215 = _nw69
		_346_v133 = (_dafny.Zero).Minus((_dafny.IntOfInt64(578)).Plus(_177_v1))
	}
	_dafny.Print(_176_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_177_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_178_v2, _dafny.SeqOf(_dafny.IntOfInt64(17))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(17)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_183_v5.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v6).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("cb"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_v8).Equals(_dafny.MultiSetOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_187_v10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v11).Equals(_dafny.MultiSetOf(_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('f'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(17), _dafny.MultiSetOf(_dafny.CodePoint('u')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(17)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_190_globalState).F6(), _dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F11).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F12()).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("cb"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F13().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState.F16).Equals(_dafny.MultiSetOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_190_globalState).F17(), _dafny.SeqOf(_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F20().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_190_globalState).F21()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(17), _dafny.MultiSetOf(_dafny.CodePoint('u')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_globalState).F24().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_globalState.F25)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_194_v13).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v14).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-17), true).UpdateUnsafe(_dafny.IntOfInt64(438), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_196_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_301_v90).Equals(_dafny.SetOf(_dafny.IntOfInt64(17))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_324_v111)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_325_v112)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_341_v128, _dafny.SeqOf(_dafny.MultiSetOf(false, false, false, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_342_v129).Equals(_dafny.SetOf(_dafny.MultiSetOf(true, true, true), _dafny.MultiSetOf(false, false, false, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_343_v130).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, true, true), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_344_v131).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_344_v131).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_344_v131).Dtor_cf18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_345_v132, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_346_v133)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_347_v134)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_348_v135).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(60))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v136).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(60)), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_367_v152).Dtor_cf0(), _dafny.SeqOf(_dafny.MultiSetOf(), _dafny.MultiSetOf(_dafny.IntOfInt64(6), _dafny.IntOfInt64(119), _dafny.IntOfInt64(-456)), _dafny.MultiSetOf(_dafny.IntOfInt64(9)), _dafny.MultiSetOf(_dafny.IntOfInt64(508), _dafny.IntOfInt64(-985)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_461_v210).Dtor_cf19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_461_v210).Dtor_cf20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_461_v210).Dtor_cf21())
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 bool
	Cf3 _dafny.Set
	Cf4 bool
	Cf5 bool
	Cf6 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool, Cf3 _dafny.Set, Cf4 bool, Cf5 bool, Cf6 _dafny.Int) D0 {
	return D0{D0_DC2{Cf2, Cf3, Cf4, Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf7 _dafny.Sequence
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf7 _dafny.Sequence) D0 {
	return D0{D0_DC3{Cf7}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Sequence) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Set {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() bool {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D0_DC3).Cf7
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Equals(data2.Cf3) && data1.Cf4 == data2.Cf4 && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf7.Equals(data2.Cf7)
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
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
	Cf8 _dafny.CodePoint
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.CodePoint) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D1) Dtor_cf8() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf8 == data2.Cf8
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
	Cf10 bool
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 bool) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf11 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 bool) D2 {
	return D2{D2_DC7{Cf11}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf12 _dafny.Int
	Cf13 bool
	Cf14 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf12 _dafny.Int, Cf13 bool, Cf14 _dafny.Int) D2 {
	return D2{D2_DC8{Cf12, Cf13, Cf14}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC5 struct {
	Cf9 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf9 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(false)
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + data.Cf9.VerbatimString(true) + ")"
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
			return ok && data1.Cf10 == data2.Cf10
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf11 == data2.Cf11
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf9.Equals(data2.Cf9)
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

type D3_DC10 struct {
	Cf16 bool
	Cf17 _dafny.Int
	Cf18 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 bool, Cf17 _dafny.Int, Cf18 bool) D3 {
	return D3{D3_DC10{Cf16, Cf17, Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf19 _dafny.CodePoint
	Cf20 bool
	Cf21 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf19 _dafny.CodePoint, Cf20 bool, Cf21 bool) D3 {
	return D3{D3_DC11{Cf19, Cf20, Cf21}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC12 struct {
	Cf22 bool
	Cf23 D0
	Cf24 bool
	Cf25 bool
	Cf26 T0
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf22 bool, Cf23 D0, Cf24 bool, Cf25 bool, Cf26 T0) D3 {
	return D3{D3_DC12{Cf22, Cf23, Cf24, Cf25, Cf26}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC9 struct {
	Cf15 _dafny.Map
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf15 _dafny.Map) D3 {
	return D3{D3_DC9{Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, _dafny.Zero, false)
}

func (_this D3) Dtor_cf16() bool {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf19() _dafny.CodePoint {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC11).Cf20
}

func (_this D3) Dtor_cf21() bool {
	return _this.Get_().(D3_DC11).Cf21
}

func (_this D3) Dtor_cf22() bool {
	return _this.Get_().(D3_DC12).Cf22
}

func (_this D3) Dtor_cf23() D0 {
	return _this.Get_().(D3_DC12).Cf23
}

func (_this D3) Dtor_cf24() bool {
	return _this.Get_().(D3_DC12).Cf24
}

func (_this D3) Dtor_cf25() bool {
	return _this.Get_().(D3_DC12).Cf25
}

func (_this D3) Dtor_cf26() T0 {
	return _this.Get_().(D3_DC12).Cf26
}

func (_this D3) Dtor_cf15() _dafny.Map {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18 == data2.Cf18
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf22 == data2.Cf22 && data1.Cf23.Equals(data2.Cf23) && data1.Cf24 == data2.Cf24 && data1.Cf25 == data2.Cf25 && _dafny.AreEqual(data1.Cf26, data2.Cf26)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf15.Equals(data2.Cf15)
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

type D4_DC14 struct {
	Cf28 _dafny.MultiSet
	Cf29 _dafny.Int
	Cf30 D2
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf28 _dafny.MultiSet, Cf29 _dafny.Int, Cf30 D2) D4 {
	return D4{D4_DC14{Cf28, Cf29, Cf30}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC13 struct {
	Cf27 _dafny.MultiSet
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf27 _dafny.MultiSet) D4 {
	return D4{D4_DC13{Cf27}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_(_dafny.EmptyMultiSet, _dafny.Zero, Companion_D2_.Default())
}

func (_this D4) Dtor_cf28() _dafny.MultiSet {
	return _this.Get_().(D4_DC14).Cf28
}

func (_this D4) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf29
}

func (_this D4) Dtor_cf30() D2 {
	return _this.Get_().(D4_DC14).Cf30
}

func (_this D4) Dtor_cf27() _dafny.MultiSet {
	return _this.Get_().(D4_DC13).Cf27
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf28.Equals(data2.Cf28) && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30.Equals(data2.Cf30)
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf27.Equals(data2.Cf27)
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
	Cf31 _dafny.MultiSet
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf31 _dafny.MultiSet) D5 {
	return D5{D5_DC15{Cf31}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D5) Dtor_cf31() _dafny.MultiSet {
	return _this.Get_().(D5_DC15).Cf31
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf31.Equals(data2.Cf31)
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

type D6_DC17 struct {
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_() D6 {
	return D6{D6_DC17{}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf32 _dafny.Array
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf32 _dafny.Array) D6 {
	return D6{D6_DC16{Cf32}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC18 struct {
	Cf33 D6
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf33 D6) D6 {
	return D6{D6_DC18{Cf33}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_()
}

func (_this D6) Dtor_cf32() _dafny.Array {
	return _this.Get_().(D6_DC16).Cf32
}

func (_this D6) Dtor_cf33() D6 {
	return _this.Get_().(D6_DC18).Cf33
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC17:
		{
			_, ok := other.Get_().(D6_DC17)
			return ok
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf32 == data2.Cf32
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf33.Equals(data2.Cf33)
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
	Cf35 bool
	Cf36 _dafny.Sequence
	Cf37 _dafny.Int
	Cf38 bool
	Cf39 _dafny.Int
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf35 bool, Cf36 _dafny.Sequence, Cf37 _dafny.Int, Cf38 bool, Cf39 _dafny.Int) D7 {
	return D7{D7_DC20{Cf35, Cf36, Cf37, Cf38, Cf39}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC21 struct {
	Cf40 bool
	Cf41 bool
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf40 bool, Cf41 bool) D7 {
	return D7{D7_DC21{Cf40, Cf41}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC19 struct {
	Cf34 _dafny.Sequence
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf34 _dafny.Sequence) D7 {
	return D7{D7_DC19{Cf34}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(false, _dafny.EmptySeq, _dafny.Zero, false, _dafny.Zero)
}

func (_this D7) Dtor_cf35() bool {
	return _this.Get_().(D7_DC20).Cf35
}

func (_this D7) Dtor_cf36() _dafny.Sequence {
	return _this.Get_().(D7_DC20).Cf36
}

func (_this D7) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf37
}

func (_this D7) Dtor_cf38() bool {
	return _this.Get_().(D7_DC20).Cf38
}

func (_this D7) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf39
}

func (_this D7) Dtor_cf40() bool {
	return _this.Get_().(D7_DC21).Cf40
}

func (_this D7) Dtor_cf41() bool {
	return _this.Get_().(D7_DC21).Cf41
}

func (_this D7) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D7_DC19).Cf34
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf35) + ", " + data.Cf36.VerbatimString(true) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf34) + ")"
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
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36.Equals(data2.Cf36) && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38 == data2.Cf38 && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41 == data2.Cf41
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D8_DC22 struct {
	Cf42 _dafny.Map
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf42 _dafny.Map) D8 {
	return D8{D8_DC22{Cf42}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D8) Dtor_cf42() _dafny.Map {
	return _this.Get_().(D8_DC22).Cf42
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf42.Equals(data2.Cf42)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm4(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int
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
	F2   _dafny.Int
	F5   _dafny.Int
	F8   _dafny.Int
	F10  _dafny.Int
	F11  _dafny.Array
	F15  _dafny.Int
	F16  _dafny.MultiSet
	F19  _dafny.Int
	F22  bool
	F23  bool
	F25  bool
	_f0  _dafny.Map
	_f1  _dafny.Int
	_f3  _dafny.Int
	_f4  bool
	_f6  _dafny.Sequence
	_f7  _dafny.Int
	_f9  bool
	_f12 _dafny.Set
	_f13 _dafny.Sequence
	_f14 bool
	_f17 _dafny.Sequence
	_f18 _dafny.Array
	_f20 _dafny.Sequence
	_f21 _dafny.Map
	_f24 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.Zero
	_this.F5 = _dafny.Zero
	_this.F8 = _dafny.Zero
	_this.F10 = _dafny.Zero
	_this.F11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F15 = _dafny.Zero
	_this.F16 = _dafny.EmptyMultiSet
	_this.F19 = _dafny.Zero
	_this.F22 = false
	_this.F23 = false
	_this.F25 = false
	_this._f0 = _dafny.EmptyMap
	_this._f1 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f4 = false
	_this._f6 = _dafny.EmptySeq
	_this._f7 = _dafny.Zero
	_this._f9 = false
	_this._f12 = _dafny.EmptySet
	_this._f13 = _dafny.EmptySeq
	_this._f14 = false
	_this._f17 = _dafny.EmptySeq
	_this._f18 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f20 = _dafny.EmptySeq
	_this._f21 = _dafny.EmptyMap
	_this._f24 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Int, f4 bool, f5 _dafny.Int, f6 _dafny.Sequence, f7 _dafny.Int, f8 _dafny.Int, f9 bool, f10 _dafny.Int, f11 _dafny.Array, f12 _dafny.Set, f13 _dafny.Sequence, f14 bool, f15 _dafny.Int, f16 _dafny.MultiSet, f17 _dafny.Sequence, f18 _dafny.Array, f19 _dafny.Int, f20 _dafny.Sequence, f21 _dafny.Map, f22 bool, f23 bool, f24 _dafny.Sequence, f25 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this).F22 = f22
		(_this).F23 = f23
		(_this)._f24 = f24
		(_this).F25 = f25
	}
}
func (_this *GlobalState) F0() _dafny.Map {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F12() _dafny.Set {
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
func (_this *GlobalState) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Array {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F20() _dafny.Sequence {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() _dafny.Map {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F24() _dafny.Sequence {
	{
		return _this._f24
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__() {
	{
	}
}
func (_this *C0) Fm4(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(152))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_468_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		}))).Cardinality()), _dafny.IntOfInt64(860))).Minus(_dafny.IntOfInt64(661))
	}
}
func (_this *C0) Fm7(p0 bool, p1 _dafny.Set, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-108)), _dafny.SeqOf((Companion_D0_.Create_DC1_(_dafny.IntOfInt64(160))).Dtor_cf1(), _dafny.IntOfInt64(51), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(574))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()))).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _469_v0 _dafny.Int
				_469_v0 = interface{}(_compr_9).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-108)), _dafny.SeqOf((Companion_D0_.Create_DC1_(_dafny.IntOfInt64(160))).Dtor_cf1(), _dafny.IntOfInt64(51), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(574))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())), _469_v0) {
					_coll9.Add((_469_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nhtgllpi")).Cardinality())), !(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(742))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}(func(_470_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('n')
					}))).Cardinality()))).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sntppqvt")).Cardinality())) > 0))
				}
			}
			return _coll9.ToMap()
		}()
	}
}
func (_this *C0) Fm8(p0 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32(((func() _dafny.Sequence {
			if true {
				return _dafny.SeqOf(_dafny.IntOfInt64(-934), _dafny.IntOfInt64(520))
			}
			return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("papx")).Cardinality())), false)).Cardinality())
		})()).Cardinality())).Plus((_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer26 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_471_i0 _dafny.Int) D0 {
			return Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ua")).Cardinality()))
		}))).Cardinality()))).Minus(_dafny.IntOfInt64(826))))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F27  _dafny.Int
	_f26 _dafny.Set
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F27 = _dafny.Zero
	_this._f26 = _dafny.EmptySet
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

func (_this *C1) Ctor__(f26 _dafny.Set, f27 _dafny.Int) {
	{
		(_this)._f26 = f26
		(_this).F27 = f27
	}
}
func (_this *C1) Fm4(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _this.F27
	}
}
func (_this *C1) Fm5(globalState *GlobalState) _dafny.CodePoint {
	{
		return ((func() _dafny.CodePoint {
			if true {
				return (_dafny.CodePoint('s'))
			}
			return _dafny.CodePoint('h')
		})())
	}
}
func (_this *C1) Fm6(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return ((_this.F27).Plus(_this.F27)).Cmp(_this.F27) <= 0
	}
}
func (_this *C1) M0(p0 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _472_v0 _dafny.CodePoint
		_ = _472_v0
		_472_v0 = _dafny.CodePoint('x')
		var _473_v1 _dafny.Map
		_ = _473_v1
		_473_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, _472_v0)
		_473_v1 = (_473_v1).Update(Companion_Default___.Fm1(true, _472_v0, _this.F27, globalState), _472_v0)
		var _474_v2 _dafny.Array
		_ = _474_v2
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw70
		_474_v2 = _nw70
		var _475_v3 _dafny.Set
		_ = _475_v3
		_475_v3 = _dafny.SetOf(_474_v2, _474_v2)
		if ((_475_v3).Intersection(_475_v3)).IsSubsetOf(_475_v3) {
			var _476_v4 _dafny.MultiSet
			_ = _476_v4
			_476_v4 = _dafny.MultiSetOf(_dafny.IntOfInt64(860))
			var _477_v5 _dafny.Sequence
			_ = _477_v5
			_477_v5 = _dafny.SeqOf(_476_v4, _476_v4)
			var _478_v6 D0
			_ = _478_v6
			_478_v6 = Companion_D0_.Create_DC0_(_477_v5)
			var _479_v7 _dafny.Sequence
			_ = _479_v7
			_479_v7 = _dafny.UnicodeSeqOfUtf8Bytes("ypmgf")
			var _480_v8 _dafny.MultiSet
			_ = _480_v8
			_480_v8 = _dafny.MultiSetOf(p0, p0)
			var _481_v9 _dafny.Map
			_ = _481_v9
			_481_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _482_v10 _dafny.Map
			_ = _482_v10
			_482_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4(_this.F27, Companion_Default___.Fm0((_this).Fm6(_dafny.SeqOf(_478_v6), _479_v7, _480_v8, globalState), _481_v9, globalState), globalState), p0)
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_474_v2), 0))
			_ = _index63
			(_474_v2).ArraySet1(_this.F27, (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_474_v2), 0))
			_ = _index64
			var _rhs94 bool = !(!(_dafny.SetOf(_this.F27)).Contains(_dafny.IntOfInt64(931))) || (p0)
			_ = _rhs94
			var _rhs95 _dafny.Int = (func() _dafny.Int {
				if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F27, _this.F27, (_480_v8).Cardinality()))).Equals(_476_v4) {
					return _this.F27
				}
				return _this.F27
			})()
			_ = _rhs95
			var _rhs96 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, p0)
			_ = _rhs96
			var _rhs97 _dafny.Int = _dafny.IntOfInt64(-345)
			_ = _rhs97
			var _lhs63 *GlobalState = globalState
			_ = _lhs63
			var _lhs64 *GlobalState = globalState
			_ = _lhs64
			var _lhs65 _dafny.Array = _474_v2
			_ = _lhs65
			var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_474_v2), 0))
			_ = _lhs66
			_lhs63.F23 = _rhs94
			_lhs64.F19 = _rhs95
			_482_v10 = _rhs96
			(_lhs65).ArraySet1(_rhs97, (_lhs66).Int())
			var _483_v11 _dafny.Array
			_ = _483_v11
			var _nwElement0_9 _dafny.Sequence = _479_v7
			_ = _nwElement0_9
			var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(5))
			_ = _nw71
			(_nw71).ArraySet1(_nwElement0_9, 0)
			(_nw71).ArraySet1(_479_v7, 1)
			(_nw71).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("wtybqap"), (Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wtybqap")).Cardinality()))).Uint32(), _472_v0), 2)
			(_nw71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_479_v7, _dafny.Companion_Sequence_.Update(_479_v7, (Companion_Default___.SafeIndex((_474_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_474_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_479_v7).Cardinality()))).Uint32(), _472_v0)), 3)
			(_nw71).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vmqrsdga"), 4)
			_483_v11 = _nw71
			_483_v11 = _483_v11
			_479_v7 = _479_v7
			(globalState).F23 = p0
			var _484_v12 _dafny.Map
			_ = _484_v12
			_484_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, _dafny.IntOfUint32((_479_v7).Cardinality()))
			_484_v12 = (_484_v12).Update(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(154))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_485_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_486_i0 _dafny.Int) _dafny.CodePoint {
					return _485_v0
				}
			})(_472_v0)))).Cardinality()), ((_this).F26()).Cardinality())
		} else {
			var _487_v13 D0
			_ = _487_v13
			_487_v13 = Companion_D0_.Create_DC2_(p0, _475_v3, p0, p0, _this.F27)
			(globalState).F22 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_487_v13).Dtor_cf6(), _this.F27)).Cardinality()).Cmp((_dafny.Zero).Minus(_this.F27)) == 0
			var _488_v14 _dafny.Array
			_ = _488_v14
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw72
			_488_v14 = _nw72
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_488_v14), 0))
			_ = _index65
			(_488_v14).ArraySet1(p0, (_index65).Int())
			var _489_v15 _dafny.Sequence
			_ = _489_v15
			_489_v15 = _dafny.UnicodeSeqOfUtf8Bytes("np")
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_488_v14), 0))
			_ = _index66
			var _rhs98 _dafny.Array = _488_v14
			_ = _rhs98
			var _rhs99 bool = (_dafny.IntOfUint32((_489_v15).Cardinality())).Cmp((_475_v3).Cardinality()) >= 0
			_ = _rhs99
			var _lhs67 _dafny.Array = _488_v14
			_ = _lhs67
			var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_488_v14), 0))
			_ = _lhs68
			_488_v14 = _rhs98
			(_lhs67).ArraySet1(_rhs99, (_lhs68).Int())
			(_this).F27 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.Fm1((_488_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_488_v14), 0))).Int()).(bool), _472_v0, _this.F27, globalState)), _this.F27)
			var _490_v16 _dafny.Array
			_ = _490_v16
			var _nwElement0_10 _dafny.CodePoint = _472_v0
			_ = _nwElement0_10
			var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
			_ = _nw73
			(_nw73).ArraySet1CodePoint(_nwElement0_10, 0)
			(_nw73).ArraySet1CodePoint(_472_v0, 1)
			(_nw73).ArraySet1CodePoint(_472_v0, 2)
			(_nw73).ArraySet1CodePoint(_472_v0, 3)
			(_nw73).ArraySet1CodePoint((_this).Fm5(globalState), 4)
			(_nw73).ArraySet1CodePoint(_472_v0, 5)
			(_nw73).ArraySet1CodePoint(_472_v0, 6)
			(_nw73).ArraySet1CodePoint(_472_v0, 7)
			(_nw73).ArraySet1CodePoint(_472_v0, 8)
			(_nw73).ArraySet1CodePoint(_472_v0, 9)
			(_nw73).ArraySet1CodePoint(_472_v0, 10)
			(_nw73).ArraySet1CodePoint(_472_v0, 11)
			(_nw73).ArraySet1CodePoint(_472_v0, 12)
			_490_v16 = _nw73
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_490_v16), 0))
			_ = _index67
			(_490_v16).ArraySet1CodePoint(_472_v0, (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_490_v16), 0))
			_ = _index68
			(_490_v16).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _dafny.CodePoint('l')
				}
				return _472_v0
			})(), (_index68).Int())
			var _491_v17 _dafny.MultiSet
			_ = _491_v17
			_491_v17 = _dafny.MultiSetOf(!(p0), (_488_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_488_v14), 0))).Int()).(bool), p0)
			(globalState).F15 = (func() _dafny.Int {
				if (true) || (p0) {
					return _this.F27
				}
				return (func() _dafny.Int {
					if (_491_v17).Contains(!((_488_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_488_v14), 0))).Int()).(bool))) {
						return (_491_v17).Multiplicity(!((_488_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_488_v14), 0))).Int()).(bool)))
					}
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_489_v15, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sc")).Cardinality()), _dafny.IntOfUint32((_489_v15).Cardinality()))).Uint32(), _472_v0)).Cardinality())
				})()
			})()
		}
		if p0 {
			var _492_v18 _dafny.Map
			_ = _492_v18
			_492_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-390))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_493_i1 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F27))
			}))))
			var _494_v19 _dafny.Sequence
			_ = _494_v19
			_494_v19 = _dafny.SeqOf(_dafny.IntOfInt64(48))
			var _495_v20 _dafny.Set
			_ = _495_v20
			_495_v20 = _dafny.SetOf(_494_v19)
			_492_v18 = (_492_v18).Update(p0, _495_v20)
			var _496_v21 _dafny.Sequence
			_ = _496_v21
			_496_v21 = _dafny.UnicodeSeqOfUtf8Bytes("uflnveram")
			(globalState).F8 = (_this).Fm4(_dafny.IntOfUint32((_496_v21).Cardinality()), p0, globalState)
			var _497_v22 _dafny.MultiSet
			_ = _497_v22
			_497_v22 = _dafny.MultiSetOf(_this.F27)
			var _498_v23 _dafny.MultiSet
			_ = _498_v23
			_498_v23 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_497_v22).Contains(_this.F27) {
					return (_497_v22).Multiplicity(_this.F27)
				}
				return _dafny.IntOfUint32((_496_v21).Cardinality())
			})())
			var _499_v24 _dafny.Sequence
			_ = _499_v24
			_499_v24 = _dafny.SeqOf(_498_v23)
			var _500_v25 D0
			_ = _500_v25
			_500_v25 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Concatenate(_499_v24, _dafny.SeqOf((_499_v24).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_499_v24).Cardinality()))).Uint32()).(_dafny.MultiSet))))
			var _source5 D0 = _500_v25
			_ = _source5
			if _source5.Is_DC1() {
				var _501___mcc_h0 _dafny.Int = _source5.Get_().(D0_DC1).Cf1
				_ = _501___mcc_h0
				var _502_cf1 _dafny.Int = _501___mcc_h0
				_ = _502_cf1
				var _503_v26 _dafny.Array
				_ = _503_v26
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw74
				_503_v26 = _nw74
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))
				_ = _index69
				(_503_v26).ArraySet1(p0, (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))
				_ = _index70
				(_503_v26).ArraySet1(p0, (_index70).Int())
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))
				_ = _index71
				var _rhs100 bool = !((_503_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))).Int()).(bool))
				_ = _rhs100
				var _rhs101 bool = (_503_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))).Int()).(bool)
				_ = _rhs101
				var _lhs69 _dafny.Array = _503_v26
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))
				_ = _lhs70
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				(_lhs69).ArraySet1(_rhs100, (_lhs70).Int())
				_lhs71.F25 = _rhs101
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))
				_ = _index72
				(_503_v26).ArraySet1((func() bool {
					if (p0) && ((_503_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))).Int()).(bool)) {
						return (_503_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_503_v26), 0))).Int()).(bool)
					}
					return p0
				})(), (_index72).Int())
				var _504_v27 T0
				_ = _504_v27
				var _nw75 *C0 = New_C0_()
				_ = _nw75
				_nw75.Ctor__()
				_504_v27 = _nw75
				_504_v27 = _504_v27
			} else if _source5.Is_DC2() {
				var _505___mcc_h1 bool = _source5.Get_().(D0_DC2).Cf2
				_ = _505___mcc_h1
				var _506___mcc_h2 _dafny.Set = _source5.Get_().(D0_DC2).Cf3
				_ = _506___mcc_h2
				var _507___mcc_h3 bool = _source5.Get_().(D0_DC2).Cf4
				_ = _507___mcc_h3
				var _508___mcc_h4 bool = _source5.Get_().(D0_DC2).Cf5
				_ = _508___mcc_h4
				var _509___mcc_h5 _dafny.Int = _source5.Get_().(D0_DC2).Cf6
				_ = _509___mcc_h5
				var _510_cf6 _dafny.Int = _509___mcc_h5
				_ = _510_cf6
				var _511_cf5 bool = _508___mcc_h4
				_ = _511_cf5
				var _512_cf4 bool = _507___mcc_h3
				_ = _512_cf4
				var _513_cf3 _dafny.Set = _506___mcc_h2
				_ = _513_cf3
				var _514_cf2 bool = _505___mcc_h1
				_ = _514_cf2
				var _515_v28 _dafny.Map
				_ = _515_v28
				_515_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _496_v21)
				var _516_v29 _dafny.Map
				_ = _516_v29
				_516_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_cf2, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_515_v28).Contains(_511_cf5) {
						return (_515_v28).Get(_511_cf5).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("g")
				})()).Cardinality()))
				var _517_v30 _dafny.Map
				_ = _517_v30
				_517_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}((func(_518_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_519_i2 _dafny.Int) _dafny.CodePoint {
						return _518_v0
					}
				})(_472_v0))), (Companion_Default___.SafeIndex((_516_v29).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_520_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_521_i2 _dafny.Int) _dafny.CodePoint {
						return _520_v0
					}
				})(_472_v0)))).Cardinality()))).Uint32(), _472_v0), _496_v21)).Cardinality()))), _510_cf6)
				_517_v30 = (_517_v30).Update((_dafny.IntOfInt64(611)).Minus(_dafny.IntOfUint32((_494_v19).Cardinality())), _dafny.IntOfInt64(-674))
				var _522_v31 _dafny.Sequence
				_ = _522_v31
				_522_v31 = _dafny.SeqOf(_500_v25, _500_v25)
				var _523_v32 _dafny.MultiSet
				_ = _523_v32
				_523_v32 = _dafny.MultiSetOf(false, _511_cf5, false)
				var _524_v33 _dafny.Map
				_ = _524_v33
				_524_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_cf2, (_this).Fm6(_522_v31, _496_v21, _523_v32, globalState))
				var _525_v34 _dafny.Sequence
				_ = _525_v34
				_525_v34 = _dafny.SeqOf(_474_v2)
				(globalState).F2 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_524_v33).Cardinality()), (_this.F27).Times(_dafny.IntOfUint32((_525_v34).Cardinality())))
				(globalState).F10 = _this.F27
				(globalState).F22 = Companion_Default___.Fm0(!((_510_cf6).Cmp(_510_cf6) == 0), Companion_Default___.Fm9(_dafny.IntOfInt64(836), globalState), globalState)
			} else if _source5.Is_DC3() {
				var _526___mcc_h6 _dafny.Sequence = _source5.Get_().(D0_DC3).Cf7
				_ = _526___mcc_h6
				var _527_cf7 _dafny.Sequence = _526___mcc_h6
				_ = _527_cf7
				(globalState).F2 = _dafny.IntOfInt64(744)
				(globalState).F2 = _this.F27
				var _528_v35 *C0
				_ = _528_v35
				var _nw76 *C0 = New_C0_()
				_ = _nw76
				_nw76.Ctor__()
				_528_v35 = _nw76
				var _529_v36 _dafny.Sequence
				_ = _529_v36
				_529_v36 = _dafny.SeqOf(_494_v19, _494_v19, (func() _dafny.Sequence {
					if !(p0) {
						return _494_v19
					}
					return _dafny.SeqOf(_this.F27)
				})(), _494_v19)
				(globalState).F5 = _dafny.IntOfUint32((_529_v36).Cardinality())
			} else {
				var _530___mcc_h7 _dafny.Sequence = _source5.Get_().(D0_DC0).Cf0
				_ = _530___mcc_h7
				var _531_cf0 _dafny.Sequence = _530___mcc_h7
				_ = _531_cf0
				var _532_v37 *C0
				_ = _532_v37
				var _nw77 *C0 = New_C0_()
				_ = _nw77
				_nw77.Ctor__()
				_532_v37 = _nw77
				var _533_v38 _dafny.Map
				_ = _533_v38
				_533_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				_533_v38 = (_533_v38).Update(true, p0)
				var _534_v39 _dafny.CodePoint
				_ = _534_v39
				_534_v39 = _dafny.CodePoint('p')
				var _535_v40 _dafny.Map
				_ = _535_v40
				_535_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _472_v0)
				var _rhs102 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_535_v40).Contains(p0) {
						return (_535_v40).Get(p0).(_dafny.CodePoint)
					}
					return _472_v0
				})()
				_ = _rhs102
				var _rhs103 _dafny.MultiSet = _497_v22
				_ = _rhs103
				_534_v39 = _rhs102
				_498_v23 = _rhs103
				var _536_v41 _dafny.MultiSet
				_ = _536_v41
				_536_v41 = _dafny.MultiSetOf(_472_v0, _472_v0, _472_v0, _472_v0, _472_v0)
				var _rhs104 bool = p0
				_ = _rhs104
				var _rhs105 bool = p0
				_ = _rhs105
				var _rhs106 _dafny.MultiSet = _536_v41
				_ = _rhs106
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				var _lhs73 *GlobalState = globalState
				_ = _lhs73
				_lhs72.F25 = _rhs104
				_lhs73.F25 = _rhs105
				_536_v41 = _rhs106
			}
			var _537_v42 _dafny.Array
			_ = _537_v42
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw78
			_537_v42 = _nw78
			var _538_v43 _dafny.Set
			_ = _538_v43
			_538_v43 = _dafny.SetOf(_537_v42)
			(globalState).F23 = ((_538_v43).Union(_538_v43)).IsDisjointFrom(_538_v43)
			var _539_v44 _dafny.MultiSet
			_ = _539_v44
			_539_v44 = _dafny.MultiSetOf(p0)
			var _540_v45 _dafny.Map
			_ = _540_v45
			_540_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Int {
				if p0 {
					return (_539_v44).Cardinality()
				}
				return _this.F27
			})())
			_540_v45 = (_540_v45).Update((_497_v22).IsDisjointFrom(_497_v22), _this.F27)
		} else {
			if p0 {
				var _541_v46 _dafny.Array
				_ = _541_v46
				var _nw79 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
				_ = _nw79
				_541_v46 = _nw79
				var _542_v47 _dafny.Map
				_ = _542_v47
				_542_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _541_v46)
				var _543_v48 _dafny.MultiSet
				_ = _543_v48
				_543_v48 = _dafny.MultiSetOf((_542_v47).Cardinality())
				_543_v48 = (_543_v48).Union(_543_v48)
				var _544_v49 *C0
				_ = _544_v49
				var _nw80 *C0 = New_C0_()
				_ = _nw80
				_nw80.Ctor__()
				_544_v49 = _nw80
				var _545_v50 _dafny.Sequence
				_ = _545_v50
				_545_v50 = _dafny.UnicodeSeqOfUtf8Bytes("l")
				_545_v50 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(416))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_546_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_547_i3 _dafny.Int) _dafny.CodePoint {
						return _546_v0
					}
				})(_472_v0))), (Companion_D2_.Create_DC5_(_545_v50)).Dtor_cf9())
				var _548_v51 *C0
				_ = _548_v51
				var _nw81 *C0 = New_C0_()
				_ = _nw81
				_nw81.Ctor__()
				_548_v51 = _nw81
				var _549_v52 _dafny.Map
				_ = _549_v52
				_549_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _550_v53 _dafny.Map
				_ = _550_v53
				_550_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_549_v52).Merge(_549_v52)).Cardinality(), (_this.F27).Plus(_this.F27))
				_550_v53 = (_550_v53).Update(_dafny.IntOfInt64(792), _this.F27)
			} else {
				var _551_v54 _dafny.MultiSet
				_ = _551_v54
				_551_v54 = _dafny.MultiSetOf(_this.F27, _this.F27, _this.F27)
				(globalState).F25 = Companion_Default___.Fm0((_551_v54).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F27))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), globalState)
				var _552_v55 _dafny.Array
				_ = _552_v55
				var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw82
				_552_v55 = _nw82
				var _553_v56 _dafny.Array
				_ = _553_v56
				var _nw83 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw83
				_553_v56 = _nw83
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_552_v55), 0))
				_ = _index73
				(_552_v55).ArraySet1(_553_v56, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_552_v55), 0))
				_ = _index74
				(_552_v55).ArraySet1(_553_v56, (_index74).Int())
				var _554_v57 _dafny.MultiSet
				_ = _554_v57
				_554_v57 = _dafny.MultiSetOf(p0)
				(globalState).F16 = ((_dafny.MultiSetOf(false, false)).Difference(_554_v57)).Union(_554_v57)
				(globalState).F22 = p0
				var _555_v58 _dafny.Sequence
				_ = _555_v58
				_555_v58 = _dafny.SeqOf(p0)
				var _556_v59 _dafny.Sequence
				_ = _556_v59
				_556_v59 = _dafny.SeqOf(_this.F27)
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_552_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_552_v55), 0))).Int()))
				_ = _arr0
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_dafny.ArrayCastTo((_552_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_552_v55), 0))).Int()))), 0))
				_ = _index75
				(_arr0).ArraySet1((_555_v58).Select((Companion_Default___.SafeIndex((_556_v59).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_556_v59).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_555_v58).Cardinality()))).Uint32()).(bool), (_index75).Int())
				var _557_v60 _dafny.Map
				_ = _557_v60
				_557_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _472_v0)
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_552_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_552_v55), 0))).Int()))
				_ = _arr1
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_dafny.ArrayCastTo((_552_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_552_v55), 0))).Int()))), 0))
				_ = _index76
				(_arr1).ArraySet1(!(_557_v60).Equals((Companion_D3_.Create_DC9_(_557_v60)).Dtor_cf15()), (_index76).Int())
			}
			var _558_v61 _dafny.MultiSet
			_ = _558_v61
			_558_v61 = _dafny.MultiSetOf(p0, p0, p0, p0)
			var _559_v62 D4
			_ = _559_v62
			_559_v62 = Companion_D4_.Create_DC13_(_558_v61)
			(globalState).F16 = (_559_v62).Dtor_cf27()
			var _560_v63 _dafny.Sequence
			_ = _560_v63
			_560_v63 = _dafny.SeqOf(p0)
			var _561_v64 _dafny.Sequence
			_ = _561_v64
			_561_v64 = _dafny.SeqOf((_560_v63).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_560_v63).Cardinality()))).Uint32()).(bool), p0, p0)
			var _562_v65 D2
			_ = _562_v65
			_562_v65 = Companion_D2_.Create_DC7_(p0)
			_561_v64 = (func() _dafny.Sequence {
				if (_562_v65).Dtor_cf11() {
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_560_v63, (Companion_Default___.SafeIndex(Companion_Default___.Fm1(p0, _472_v0, _this.F27, globalState), _dafny.IntOfUint32((_560_v63).Cardinality()))).Uint32(), false), (Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_560_v63, (Companion_Default___.SafeIndex(Companion_Default___.Fm1(p0, _472_v0, _this.F27, globalState), _dafny.IntOfUint32((_560_v63).Cardinality()))).Uint32(), false)).Cardinality()))).Uint32(), p0)
				}
				return _561_v64
			})()
			var _563_v66 _dafny.Set
			_ = _563_v66
			_563_v66 = _dafny.SetOf(_this.F27)
			var _564_v67 _dafny.Map
			_ = _564_v67
			_564_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _565_v68 _dafny.Sequence
			_ = _565_v68
			_565_v68 = _dafny.UnicodeSeqOfUtf8Bytes("eelid")
			var _566_v69 _dafny.Sequence
			_ = _566_v69
			_566_v69 = _dafny.SeqOf(_565_v68)
			var _567_v70 _dafny.Sequence
			_ = _567_v70
			_567_v70 = _dafny.SeqOf(_dafny.IntOfUint32(((_566_v69).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_566_v69).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			var _568_v71 _dafny.Map
			_ = _568_v71
			_568_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F27, Companion_Default___.Fm1(p0, _472_v0, _this.F27, globalState), (_563_v66).Cardinality(), _dafny.IntOfInt64(-710), _dafny.IntOfUint32((Companion_Default___.Fm10(_this.F27, _this.F27, p0, Companion_Default___.Fm0(p0, _564_v67, globalState), globalState)).Cardinality())), _567_v70), Companion_D3_.Create_DC10_(p0, _this.F27, p0))
			_568_v71 = _568_v71
			(globalState).F8 = (_dafny.Zero).Minus(_this.F27)
		}
		var _569_v72 _dafny.Array
		_ = _569_v72
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_8
		var _nw84 _dafny.Array
		_ = _nw84
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw84 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Map = (func(_570_p0 bool) func(_dafny.Int) _dafny.Map {
				return func(_571_i4 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(806), _570_p0)
				}
			})(p0)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw84 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw84).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw84).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_569_v72 = _nw84
		var _572_v73 _dafny.Map
		_ = _572_v73
		_572_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, p0)
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_569_v72), 0))
		_ = _index77
		(_569_v72).ArraySet1(_572_v73, (_index77).Int())
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_569_v72), 0))
		_ = _index78
		(_569_v72).ArraySet1((_572_v73).Merge((_572_v73).Merge(_572_v73)), (_index78).Int())
		var _573_v74 *C0
		_ = _573_v74
		var _nw85 *C0 = New_C0_()
		_ = _nw85
		_nw85.Ctor__()
		_573_v74 = _nw85
		(globalState).F23 = p0
		r0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F27, _this.F27))
		r1 = _this.F27
		return r0, r1
	}
}
func (_this *C1) F26() _dafny.Set {
	{
		return _this._f26
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
