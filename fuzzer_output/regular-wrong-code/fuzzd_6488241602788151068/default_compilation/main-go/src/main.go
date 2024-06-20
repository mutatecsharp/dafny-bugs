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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(-249)
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((false) || (!(true)), (_dafny.IntOfInt64(252)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-707))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality())) != 0, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('m'))).Cardinality()).Cmp(_dafny.IntOfInt64(176)) >= 0)
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) bool {
	return (Companion_D0_.Create_DC1_((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(961))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(594))).Cardinality())
	}))).Cardinality()), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false)).Cardinality())).Cardinality())).Cardinality(), _dafny.IntOfInt64(-495), false)).Dtor_cf3()
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('g')), _dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('u'), _dafny.CodePoint('r'), _dafny.CodePoint('f'), _dafny.CodePoint('u'))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _dafny.SeqOf(_dafny.CodePoint('i'))))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(!(true)))), _dafny.UnicodeSeqOfUtf8Bytes("raqnsyr"))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('m')
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	var _source0 _dafny.Set = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ixyg")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("auue")).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.CodePoint('c'))).Cardinality(), (_dafny.MultiSetOf(true, false, false, false, (Companion_D0_.Create_DC2_(false, !(true), _dafny.IntOfInt64(362), true)).Dtor_cf5())).Cardinality()))
	_ = _source0
	var _3___mcc_h0 _dafny.Set = _source0
	_ = _3___mcc_h0
	var _4_cf35 _dafny.Set = _3___mcc_h0
	_ = _4_cf35
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SeqOf(_dafny.IntOfInt64(124), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(331))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _dafny.UnicodeSeqOfUtf8Bytes("gwyvjb"))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _6_v0 _dafny.Sequence
			_6_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(331))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})), _dafny.UnicodeSeqOfUtf8Bytes("gwyvjb"))).Contains(_6_v0) {
				_coll0.Add(_6_v0)
			}
		}
		return _coll0.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_7_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-230)
	})))).Cardinality()), _dafny.IntOfInt64(-762)), ((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kl")).Cardinality()), _dafny.IntOfInt64(977), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-882))).Cardinality()), _dafny.IntOfInt64(983))).Cardinality()))).Cardinality())).Plus(_dafny.IntOfInt64(993)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	var _source1 D2 = Companion_D2_.Create_DC8_()
	_ = _source1
	if _source1.Is_DC8() {
		return (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(116), _dafny.IntOfInt64(-258), _dafny.IntOfInt64(647), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hklxydp")).Cardinality())))).Difference(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((Companion_D9_.Create_DC27_(_dafny.SetOf(_dafny.IntOfInt64(673)))).Dtor_cf37()).Cardinality(), true)).Cardinality()))
	} else if _source1.Is_DC9() {
		var _8___mcc_h0 _dafny.Int = _source1.Get_().(D2_DC9).Cf13
		_ = _8___mcc_h0
		var _9___mcc_h1 *C0 = _source1.Get_().(D2_DC9).Cf14
		_ = _9___mcc_h1
		var _10___mcc_h2 *C0 = _source1.Get_().(D2_DC9).Cf15
		_ = _10___mcc_h2
		var _11_cf15 *C0 = _10___mcc_h2
		_ = _11_cf15
		var _12_cf14 *C0 = _9___mcc_h1
		_ = _12_cf14
		var _13_cf13 _dafny.Int = _8___mcc_h0
		_ = _13_cf13
		return (_dafny.SetOf(_11_cf15.F19, _13_cf13, _11_cf15.F19, (_dafny.Zero).Minus(_12_cf14.F19))).Union(_dafny.SetOf(_11_cf15.F19, _11_cf15.F19))
	} else if _source1.Is_DC7() {
		var _14___mcc_h3 _dafny.Array = _source1.Get_().(D2_DC7).Cf12
		_ = _14___mcc_h3
		var _15_cf12 _dafny.Array = _14___mcc_h3
		_ = _15_cf12
		return (func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_16_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(631)
			}))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _17_v0 _dafny.Int
				_17_v0 = interface{}(_compr_1).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_16_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(631)
				})), _17_v0) {
					_coll1.Add(Companion_Default___.SafeDivisionInt(_17_v0, (_dafny.SetOf(true, true, false)).Cardinality()))
				}
			}
			return _coll1.ToSet()
		}()).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nxglsvqlj")).Cardinality())))
	} else {
		var _18___mcc_h4 D2 = _source1.Get_().(D2_DC10).Cf16
		_ = _18___mcc_h4
		var _19_cf16 D2 = _18___mcc_h4
		_ = _19_cf16
		return _dafny.SetOf(_dafny.IntOfInt64(134), _dafny.IntOfInt64(321))
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("erpx")).Cardinality())), _dafny.IntOfInt64(36), (false) && (true))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC19_(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("bpwtnk"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(870))
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(true, !(false), !(true), true, !(false)))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("slbcnr")).Cardinality()), !(false))).Cardinality(), _dafny.IntOfInt64(953))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _21_v0 _dafny.Int
			_21_v0 = interface{}(_compr_2).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("slbcnr")).Cardinality()), !(false))).Cardinality(), _dafny.IntOfInt64(953))).Contains(_21_v0) {
				_coll2.Add((_21_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())), (_dafny.SetOf(_dafny.IntOfInt64(-705), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())
			}
		}
		return _coll2.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rihcohcjh")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true, (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-453))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality())).Cmp(_dafny.IntOfInt64(70)) != 0, (func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(!(false)), _dafny.SeqOf(false), _dafny.SeqOf(false, false), _dafny.SeqOf(true, true), _dafny.SeqOf(true))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _23_v0 _dafny.Sequence
			_23_v0 = interface{}(_compr_3).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(!(false)), _dafny.SeqOf(false), _dafny.SeqOf(false, false), _dafny.SeqOf(true, true), _dafny.SeqOf(true))).Contains(_23_v0) {
				_coll3.Add(_23_v0, Companion_D0_.Create_DC0_(_dafny.IntOfInt64(566)))
			}
		}
		return _coll3.ToMap()
	}()).Equals(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(true))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _24_v1 _dafny.Sequence
			_24_v1 = interface{}(_compr_4).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(true)), _24_v1) {
				_coll4.Add(_24_v1, Companion_D0_.Create_DC0_(_dafny.IntOfInt64(421)))
			}
		}
		return _coll4.ToMap()
	}()), (_dafny.MultiSetOf(true, false, !(false))).IsDisjointFrom(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 bool, globalState *GlobalState) D6 {
	var _source2 _dafny.Set = (func() _dafny.Set {
		if false {
			return _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(722), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-185))).Cardinality())))
		}
		return _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_25_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality()), false)).Cardinality(), _dafny.IntOfInt64(265)))
	})()
	_ = _source2
	var _26___mcc_h0 _dafny.Set = _source2
	_ = _26___mcc_h0
	var _27_cf35 _dafny.Set = _26___mcc_h0
	_ = _27_cf35
	return Companion_D6_.Create_DC23_((Companion_D0_.Create_DC2_(!(true), !(true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wmttkoul")).Cardinality()), false)).Dtor_cf4(), _dafny.IntOfInt64(136), _dafny.IntOfInt64(541))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, _dafny.Int, _dafny.MultiSet) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Sequence = _dafny.EmptySeq
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var r3 _dafny.MultiSet = _dafny.EmptyMultiSet
	_ = r3
	var _28_v0 _dafny.MultiSet
	_ = _28_v0
	_28_v0 = _dafny.MultiSetOf(p0, p0, false)
	var _29_v1 _dafny.Sequence
	_ = _29_v1
	_29_v1 = _dafny.SeqOf(_28_v0)
	var _30_v2 _dafny.Sequence
	_ = _30_v2
	_30_v2 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_28_v0, _28_v0), _29_v1), _dafny.Companion_Sequence_.Concatenate(_29_v1, _29_v1))
	var _31_v3 _dafny.Sequence
	_ = _31_v3
	_31_v3 = _dafny.UnicodeSeqOfUtf8Bytes("rg")
	var _pat_let_tv0 = p2
	_ = _pat_let_tv0
	var _pat_let_tv1 = p2
	_ = _pat_let_tv1
	var _rhs0 _dafny.Int = _dafny.IntOfUint32(((_30_v2).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_30_v2).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
	_ = _rhs0
	var _rhs1 _dafny.Sequence = _31_v3
	_ = _rhs1
	var _rhs2 _dafny.Int = func(_source3 D4) _dafny.Int {
		if _source3.Is_DC15() {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality())
		} else if _source3.Is_DC14() {
			var _32___mcc_h0 _dafny.Sequence = _source3.Get_().(D4_DC14).Cf21
			_ = _32___mcc_h0
			var _33_cf21 _dafny.Sequence = _32___mcc_h0
			_ = _33_cf21
			return _pat_let_tv0
		} else {
			var _34___mcc_h1 D4 = _source3.Get_().(D4_DC16).Cf22
			_ = _34___mcc_h1
			var _35_cf22 D4 = _34___mcc_h1
			_ = _35_cf22
			return _pat_let_tv1
		}
	}(Companion_D4_.Create_DC15_())
	_ = _rhs2
	var _rhs3 bool = p1
	_ = _rhs3
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 *GlobalState = globalState
	_ = _lhs1
	var _lhs2 *GlobalState = globalState
	_ = _lhs2
	_lhs0.F12 = _rhs0
	r1 = _rhs1
	_lhs1.F2 = _rhs2
	_lhs2.F3 = _rhs3
	var _36_v4 _dafny.CodePoint
	_ = _36_v4
	_36_v4 = _dafny.CodePoint('b')
	var _37_v5 D5
	_ = _37_v5
	_37_v5 = Companion_D5_.Create_DC17_(_36_v4)
	var _pat_let_tv2 = p2
	_ = _pat_let_tv2
	var _pat_let_tv3 = p2
	_ = _pat_let_tv3
	var _pat_let_tv4 = p2
	_ = _pat_let_tv4
	var _pat_let_tv5 = p2
	_ = _pat_let_tv5
	(globalState).F2 = func(_source4 D5) _dafny.Int {
		if _source4.Is_DC18() {
			var _38___mcc_h2 _dafny.Int = _source4.Get_().(D5_DC18).Cf24
			_ = _38___mcc_h2
			var _39___mcc_h3 bool = _source4.Get_().(D5_DC18).Cf25
			_ = _39___mcc_h3
			var _40___mcc_h4 _dafny.Array = _source4.Get_().(D5_DC18).Cf26
			_ = _40___mcc_h4
			var _41___mcc_h5 bool = _source4.Get_().(D5_DC18).Cf27
			_ = _41___mcc_h5
			var _42_cf27 bool = _41___mcc_h5
			_ = _42_cf27
			var _43_cf26 _dafny.Array = _40___mcc_h4
			_ = _43_cf26
			var _44_cf25 bool = _39___mcc_h3
			_ = _44_cf25
			var _45_cf24 _dafny.Int = _38___mcc_h2
			_ = _45_cf24
			return _pat_let_tv2
		} else if _source4.Is_DC19() {
			var _46___mcc_h6 bool = _source4.Get_().(D5_DC19).Cf28
			_ = _46___mcc_h6
			var _47_cf28 bool = _46___mcc_h6
			_ = _47_cf28
			return _pat_let_tv3
		} else if _source4.Is_DC17() {
			var _48___mcc_h7 _dafny.CodePoint = _source4.Get_().(D5_DC17).Cf23
			_ = _48___mcc_h7
			var _49_cf23 _dafny.CodePoint = _48___mcc_h7
			_ = _49_cf23
			return _pat_let_tv4
		} else {
			var _50___mcc_h8 D5 = _source4.Get_().(D5_DC20).Cf29
			_ = _50___mcc_h8
			var _51_cf29 D5 = _50___mcc_h8
			_ = _51_cf29
			return (_dafny.Zero).Minus(_pat_let_tv5)
		}
	}(_37_v5)
	if (p1) && (p1) {
		var _52_v6 _dafny.Array
		_ = _52_v6
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_53_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_54_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_54_i0, _53_p2)
				}
			})(p2)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw0).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw0).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_52_v6 = _nw0
		var _55_v7 _dafny.Sequence
		_ = _55_v7
		_55_v7 = _dafny.SeqOf(p0)
		var _56_v8 _dafny.Map
		_ = _56_v8
		_56_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _55_v7)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_52_v6), 0))
		_ = _index0
		(_52_v6).ArraySet1((_56_v8).Cardinality(), (_index0).Int())
		var _57_v9 _dafny.Sequence
		_ = _57_v9
		_57_v9 = _dafny.SeqOf(p2)
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_52_v6), 0))
		_ = _index1
		(_52_v6).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_57_v9).Cardinality()), p2), (_index1).Int())
		var _58_v10 _dafny.Sequence
		_ = _58_v10
		_58_v10 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_59_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_60_i2 _dafny.Int) _dafny.Int {
				return _59_p2
			}
		})(p2))))
		var _61_v11 _dafny.Sequence
		_ = _61_v11
		_61_v11 = _dafny.SeqOf(_58_v10, _58_v10)
		var _62_v12 _dafny.Map
		_ = _62_v12
		_62_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfInt64(-913))
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_52_v6), 0))
		_ = _index2
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_52_v6), 0))
		_ = _index3
		var _rhs4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(812))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_63_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})), _31_v3)
		_ = _rhs4
		var _rhs5 bool = !_dafny.Companion_Sequence_.Contains(_57_v9, p2)
		_ = _rhs5
		var _rhs6 _dafny.Int = (_dafny.IntOfUint32(((_61_v11).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_61_v11).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_62_v12).Cardinality(), (_57_v9).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_57_v9).Cardinality()))).Uint32()).(_dafny.Int))))
		_ = _rhs6
		var _rhs7 _dafny.Int = p2
		_ = _rhs7
		var _lhs3 *GlobalState = globalState
		_ = _lhs3
		var _lhs4 _dafny.Array = _52_v6
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_52_v6), 0))
		_ = _lhs5
		var _lhs6 _dafny.Array = _52_v6
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_52_v6), 0))
		_ = _lhs7
		r1 = _rhs4
		_lhs3.F3 = _rhs5
		(_lhs4).ArraySet1(_rhs6, (_lhs5).Int())
		(_lhs6).ArraySet1(_rhs7, (_lhs7).Int())
		_36_v4 = _36_v4
		(globalState).F3 = p0
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_52_v6), 0))
		_ = _index4
		(_52_v6).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfUint32((_31_v3).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bgihoex")).Cardinality())), (_index4).Int())
		var _64_v13 _dafny.Array
		_ = _64_v13
		var _nwElement0_0 bool = p1
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.One)
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		_64_v13 = _nw1
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_64_v13), 0))
		_ = _index5
		(_64_v13).ArraySet1(Companion_Default___.Fm2(p0, globalState), (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_64_v13), 0))
		_ = _index6
		(_64_v13).ArraySet1(Companion_Default___.Fm2((_55_v7).Select((Companion_Default___.SafeIndex((_52_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_52_v6), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_55_v7).Cardinality()))).Uint32()).(bool), globalState), (_index6).Int())
	} else {
		var _65_v14 _dafny.Map
		_ = _65_v14
		_65_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p3)
		_65_v14 = (_65_v14).Update(p3, p3)
		var _66_v15 *C0
		_ = _66_v15
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__(_dafny.IntOfInt64(536))
		_66_v15 = _nw2
		var _67_v16 _dafny.Sequence
		_ = _67_v16
		_67_v16 = _dafny.SeqOf(_66_v15)
		var _68_v17 D4
		_ = _68_v17
		_68_v17 = Companion_D4_.Create_DC14_(_67_v16)
		var _69_v18 _dafny.Map
		_ = _69_v18
		_69_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC22_(), _68_v17)
		var _70_v19 _dafny.Array
		_ = _70_v19
		var _nwElement0_1 _dafny.Map = _69_v18
		_ = _nwElement0_1
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(3))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_1, 0)
		(_nw3).ArraySet1((_69_v18).Merge(_69_v18), 1)
		(_nw3).ArraySet1(_69_v18, 2)
		_70_v19 = _nw3
		var _71_v20 D6
		_ = _71_v20
		_71_v20 = Companion_D6_.Create_DC22_()
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_70_v19), 0))
		_ = _index7
		(_70_v19).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v20, _68_v17), (_index7).Int())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_70_v19), 0))
		_ = _index8
		(_70_v19).ArraySet1(_69_v18, (_index8).Int())
		var _72_v21 _dafny.Array
		_ = _72_v21
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw4
		_72_v21 = _nw4
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_72_v21), 0))
		_ = _index9
		(_72_v21).ArraySet1(p1, (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_72_v21), 0))
		_ = _index10
		(_72_v21).ArraySet1(p1, (_index10).Int())
		var _73_v22 _dafny.Set
		_ = _73_v22
		_73_v22 = _dafny.SetOf(p0)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_72_v21), 0))
		_ = _index11
		(_72_v21).ArraySet1((((Companion_Default___.Fm17(globalState)).Intersection(_73_v22)).Cardinality()).Cmp((func() _dafny.Int {
			if p0 {
				return _66_v15.F19
			}
			return (_dafny.Zero).Minus(_66_v15.F19)
		})()) >= 0, (_index11).Int())
		var _74_v23 _dafny.Sequence
		_ = _74_v23
		_74_v23 = _dafny.SeqOf(_66_v15.F19)
		var _75_v24 *C0
		_ = _75_v24
		var _nw5 *C0 = New_C0_()
		_ = _nw5
		_nw5.Ctor__((_74_v23).Select((Companion_Default___.SafeIndex(_66_v15.F19, _dafny.IntOfUint32((_74_v23).Cardinality()))).Uint32()).(_dafny.Int))
		_75_v24 = _nw5
	}
	var _76_v25 _dafny.Map
	_ = _76_v25
	_76_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
	var _hi0 _dafny.Int = p2
	_ = _hi0
	for _77_i3 := (_76_v25).Cardinality(); _77_i3.Cmp(_hi0) < 0; _77_i3 = _77_i3.Plus(_dafny.One) {
		var _78_v26 _dafny.Array
		_ = _78_v26
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw6
		_78_v26 = _nw6
		var _79_v27 _dafny.MultiSet
		_ = _79_v27
		_79_v27 = _dafny.MultiSetOf(_78_v26, _78_v26, _78_v26)
		(globalState).F3 = ((_79_v27).Intersection(_79_v27)).IsSubsetOf(_79_v27)
		_78_v26 = _78_v26
		(globalState).F2 = p2
		var _80_v28 _dafny.Array
		_ = _80_v28
		var _nwElement0_2 bool = (p2).Cmp(_77_i3) == 0
		_ = _nwElement0_2
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(5))
		_ = _nw7
		(_nw7).ArraySet1(_nwElement0_2, 0)
		(_nw7).ArraySet1(p0, 1)
		(_nw7).ArraySet1((Companion_Default___.Fm2(!(p1), globalState)) && (Companion_Default___.Fm2(Companion_Default___.Fm2(true, globalState), globalState)), 2)
		(_nw7).ArraySet1(p1, 3)
		(_nw7).ArraySet1((p1) && (p1), 4)
		_80_v28 = _nw7
		var _81_v29 *C0
		_ = _81_v29
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__((_dafny.IntOfInt64(343)).Times(p2))
		_81_v29 = _nw8
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_80_v28), 0))
		_ = _index12
		(_80_v28).ArraySet1((_77_i3).Cmp((_dafny.Zero).Minus(_77_i3)) != 0, (_index12).Int())
		var _82_v30 _dafny.Sequence
		_ = _82_v30
		_82_v30 = _dafny.SeqOf(p0)
		var _83_v31 _dafny.Map
		_ = _83_v31
		_83_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), (_dafny.Zero).Minus(_77_i3))
		var _84_v32 _dafny.Sequence
		_ = _84_v32
		_84_v32 = _dafny.SeqOf(_77_i3)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_80_v28), 0))
		_ = _index13
		var _rhs8 _dafny.Array = _80_v28
		_ = _rhs8
		var _rhs9 *C0 = _81_v29
		_ = _rhs9
		var _rhs10 bool = ((_dafny.SetOf((_82_v30).Select((Companion_Default___.SafeIndex(_77_i3, _dafny.IntOfUint32((_82_v30).Cardinality()))).Uint32()).(bool), (_82_v30).Select((Companion_Default___.SafeIndex(_81_v29.F19, _dafny.IntOfUint32((_82_v30).Cardinality()))).Uint32()).(bool))).Union(_dafny.SetOf(p1))).IsProperSubsetOf(Companion_Default___.Fm17(globalState))
		_ = _rhs10
		var _rhs11 bool = (_81_v29.F19).Cmp((func() _dafny.Int {
			if (_83_v31).Contains(p0) {
				return (_83_v31).Get(p0).(_dafny.Int)
			}
			return _77_i3
		})()) >= 0
		_ = _rhs11
		var _rhs12 bool = ((_dafny.IntOfUint32((_84_v32).Cardinality())).Cmp(_dafny.IntOfUint32((_31_v3).Cardinality())) < 0) && (p1)
		_ = _rhs12
		var _lhs8 _dafny.Array = _80_v28
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_80_v28), 0))
		_ = _lhs9
		var _lhs10 *GlobalState = globalState
		_ = _lhs10
		var _lhs11 *GlobalState = globalState
		_ = _lhs11
		_80_v28 = _rhs8
		_81_v29 = _rhs9
		(_lhs8).ArraySet1(_rhs10, (_lhs9).Int())
		_lhs10.F3 = _rhs11
		_lhs11.F3 = _rhs12
	}
	var _pat_let_tv6 = p3
	_ = _pat_let_tv6
	var _pat_let_tv7 = p3
	_ = _pat_let_tv7
	var _pat_let_tv8 = p1
	_ = _pat_let_tv8
	var _pat_let_tv9 = p0
	_ = _pat_let_tv9
	var _pat_let_tv10 = p1
	_ = _pat_let_tv10
	var _pat_let_tv11 = p0
	_ = _pat_let_tv11
	var _pat_let_tv12 = p1
	_ = _pat_let_tv12
	var _pat_let_tv13 = p3
	_ = _pat_let_tv13
	var _pat_let_tv14 = p0
	_ = _pat_let_tv14
	if func(_source5 D0) bool {
		if _source5.Is_DC1() {
			var _85___mcc_h9 _dafny.Int = _source5.Get_().(D0_DC1).Cf1
			_ = _85___mcc_h9
			var _86___mcc_h10 _dafny.Int = _source5.Get_().(D0_DC1).Cf2
			_ = _86___mcc_h10
			var _87___mcc_h11 bool = _source5.Get_().(D0_DC1).Cf3
			_ = _87___mcc_h11
			var _88_cf3 bool = _87___mcc_h11
			_ = _88_cf3
			var _89_cf2 _dafny.Int = _86___mcc_h10
			_ = _89_cf2
			var _90_cf1 _dafny.Int = _85___mcc_h9
			_ = _90_cf1
			return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_pat_let_tv6), _dafny.SeqOf(_pat_let_tv7, _pat_let_tv8, _pat_let_tv9, _pat_let_tv10))
		} else if _source5.Is_DC2() {
			var _91___mcc_h12 bool = _source5.Get_().(D0_DC2).Cf4
			_ = _91___mcc_h12
			var _92___mcc_h13 bool = _source5.Get_().(D0_DC2).Cf5
			_ = _92___mcc_h13
			var _93___mcc_h14 _dafny.Int = _source5.Get_().(D0_DC2).Cf6
			_ = _93___mcc_h14
			var _94___mcc_h15 bool = _source5.Get_().(D0_DC2).Cf7
			_ = _94___mcc_h15
			var _95_cf7 bool = _94___mcc_h15
			_ = _95_cf7
			var _96_cf6 _dafny.Int = _93___mcc_h14
			_ = _96_cf6
			var _97_cf5 bool = _92___mcc_h13
			_ = _97_cf5
			var _98_cf4 bool = _91___mcc_h12
			_ = _98_cf4
			return _pat_let_tv11
		} else {
			var _99___mcc_h16 _dafny.Int = _source5.Get_().(D0_DC0).Cf0
			_ = _99___mcc_h16
			var _100_cf0 _dafny.Int = _99___mcc_h16
			_ = _100_cf0
			if _pat_let_tv12 {
				return _pat_let_tv13
			} else {
				return _pat_let_tv14
			}
		}
	}(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-878))) {
		var _101_v33 *C0
		_ = _101_v33
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__(p2)
		_101_v33 = _nw9
		var _102_v34 _dafny.Sequence
		_ = _102_v34
		_102_v34 = _dafny.SeqOf(_101_v33)
		var _103_v35 D4
		_ = _103_v35
		_103_v35 = Companion_D4_.Create_DC14_(_102_v34)
		var _104_v36 D4
		_ = _104_v36
		_104_v36 = Companion_D4_.Create_DC16_(_103_v35)
		var _pat_let_tv15 = _103_v35
		_ = _pat_let_tv15
		var _105_v37 _dafny.Array
		_ = _105_v37
		var _nwElement0_3 D4 = _104_v36
		_ = _nwElement0_3
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(19))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_3, 0)
		(_nw10).ArraySet1(_104_v36, 1)
		(_nw10).ArraySet1(_104_v36, 2)
		(_nw10).ArraySet1(_104_v36, 3)
		(_nw10).ArraySet1(_104_v36, 4)
		(_nw10).ArraySet1(_104_v36, 5)
		(_nw10).ArraySet1(_104_v36, 6)
		(_nw10).ArraySet1(_104_v36, 7)
		(_nw10).ArraySet1(_104_v36, 8)
		(_nw10).ArraySet1(_104_v36, 9)
		(_nw10).ArraySet1(_104_v36, 10)
		(_nw10).ArraySet1(_104_v36, 11)
		(_nw10).ArraySet1(_104_v36, 12)
		(_nw10).ArraySet1(func(_pat_let0_0 D4) D4 {
			return func(_106_dt__update__tmp_h0 D4) D4 {
				return func(_pat_let1_0 D4) D4 {
					return func(_107_dt__update_hcf22_h0 D4) D4 {
						return Companion_D4_.Create_DC16_(_107_dt__update_hcf22_h0)
					}(_pat_let1_0)
				}(_pat_let_tv15)
			}(_pat_let0_0)
		}(_104_v36), 13)
		(_nw10).ArraySet1(_104_v36, 14)
		(_nw10).ArraySet1(_104_v36, 15)
		(_nw10).ArraySet1(Companion_D4_.Create_DC16_(_103_v35), 16)
		(_nw10).ArraySet1(Companion_D4_.Create_DC16_(_103_v35), 17)
		(_nw10).ArraySet1(_104_v36, 18)
		_105_v37 = _nw10
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_105_v37), 0))
		_ = _index14
		(_105_v37).ArraySet1(_104_v36, (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_105_v37), 0))
		_ = _index15
		(_105_v37).ArraySet1(_104_v36, (_index15).Int())
		r0 = _101_v33.F19
		var _108_v38 *C0
		_ = _108_v38
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__(_101_v33.F19)
		_108_v38 = _nw11
		var _109_v39 _dafny.Sequence
		_ = _109_v39
		_109_v39 = _dafny.SeqOf(p3, Companion_Default___.Fm2(p3, globalState))
		var _110_v40 _dafny.Map
		_ = _110_v40
		_110_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfUint32((Companion_Default___.Fm14(p1, globalState)).Cardinality()))
		if (_109_v39).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_110_v40).Contains(Companion_Default___.Fm2(p0, globalState)) {
				return (_110_v40).Get(Companion_Default___.Fm2(p0, globalState)).(_dafny.Int)
			}
			return _dafny.IntOfInt64(-950)
		})(), _dafny.IntOfUint32((_109_v39).Cardinality()))).Uint32()).(bool) {
			r1 = Companion_Default___.Fm5(_dafny.IntOfUint32((_31_v3).Cardinality()), globalState)
			var _rhs13 bool = !((_28_v0).IsProperSubsetOf((_28_v0).Update(p3, Companion_Default___.Abs(p2))))
			_ = _rhs13
			var _rhs14 *C0 = (func() *C0 {
				if !((_dafny.IntOfInt64(384)).Cmp(_101_v33.F19) < 0) {
					return _101_v33
				}
				return _108_v38
			})()
			_ = _rhs14
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			_lhs12.F3 = _rhs13
			_108_v38 = _rhs14
			var _111_v41 _dafny.Array
			_ = _111_v41
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_1
			var _nw12 _dafny.Array
			_ = _nw12
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw12 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_112_p0 bool) func(_dafny.Int) bool {
					return func(_113_i4 _dafny.Int) bool {
						return _112_p0
					}
				})(p0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw12 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw12).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw12).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_111_v41 = _nw12
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))
			_ = _index16
			(_111_v41).ArraySet1((func() bool {
				if p1 {
					return false
				}
				return p3
			})(), (_index16).Int())
			var _114_v42 _dafny.Map
			_ = _114_v42
			_114_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v4, (_108_v38.F19).Cmp(_dafny.IntOfUint32((_31_v3).Cardinality())) <= 0)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))
			_ = _index17
			(_111_v41).ArraySet1((func() bool {
				if (_114_v42).Contains(_36_v4) {
					return (_114_v42).Get(_36_v4).(bool)
				}
				return p0
			})(), (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))
			_ = _index18
			(_111_v41).ArraySet1(!((_76_v25).Equals(_76_v25)) || ((_28_v0).IsSubsetOf((_dafny.MultiSetOf((_111_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))).Int()).(bool), !(true))).Update((_111_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))).Int()).(bool), Companion_Default___.Abs(p2)))), (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))
			_ = _index19
			var _rhs15 _dafny.Int = _101_v33.F19
			_ = _rhs15
			var _rhs16 bool = Companion_Default___.Fm2((_111_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))).Int()).(bool), globalState)
			_ = _rhs16
			var _lhs13 *GlobalState = globalState
			_ = _lhs13
			var _lhs14 _dafny.Array = _111_v41
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_111_v41), 0))
			_ = _lhs15
			_lhs13.F2 = _rhs15
			(_lhs14).ArraySet1(_rhs16, (_lhs15).Int())
		} else {
			var _115_v43 D6
			_ = _115_v43
			_115_v43 = Companion_D6_.Create_DC23_(p3, _101_v33.F19, _101_v33.F19)
			var _116_v44 _dafny.Map
			_ = _116_v44
			_116_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_115_v43).Dtor_cf31(), _108_v38)
			var _117_v45 _dafny.Array
			_ = _117_v45
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw13
			_117_v45 = _nw13
			var _118_v46 D5
			_ = _118_v46
			_118_v46 = Companion_D5_.Create_DC18_(_108_v38.F19, p1, _117_v45, p1)
			var _119_v47 _dafny.Array
			_ = _119_v47
			var _nwElement0_4 *C0 = _108_v38
			_ = _nwElement0_4
			var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(27))
			_ = _nw14
			(_nw14).ArraySet1(_nwElement0_4, 0)
			(_nw14).ArraySet1(_108_v38, 1)
			(_nw14).ArraySet1(_108_v38, 2)
			(_nw14).ArraySet1(_101_v33, 3)
			(_nw14).ArraySet1(_108_v38, 4)
			(_nw14).ArraySet1(_108_v38, 5)
			(_nw14).ArraySet1(_101_v33, 6)
			(_nw14).ArraySet1(_108_v38, 7)
			(_nw14).ArraySet1(_108_v38, 8)
			(_nw14).ArraySet1(_108_v38, 9)
			(_nw14).ArraySet1(_108_v38, 10)
			(_nw14).ArraySet1((func() *C0 {
				if (_116_v44).Contains((_118_v46).Dtor_cf25()) {
					return (_116_v44).Get((_118_v46).Dtor_cf25()).(*C0)
				}
				return _101_v33
			})(), 11)
			(_nw14).ArraySet1(_108_v38, 12)
			(_nw14).ArraySet1(_108_v38, 13)
			(_nw14).ArraySet1(_108_v38, 14)
			(_nw14).ArraySet1(_101_v33, 15)
			(_nw14).ArraySet1(_108_v38, 16)
			(_nw14).ArraySet1(_108_v38, 17)
			(_nw14).ArraySet1(_101_v33, 18)
			(_nw14).ArraySet1(_108_v38, 19)
			(_nw14).ArraySet1(_108_v38, 20)
			(_nw14).ArraySet1(_108_v38, 21)
			(_nw14).ArraySet1(_101_v33, 22)
			(_nw14).ArraySet1(_108_v38, 23)
			(_nw14).ArraySet1(_101_v33, 24)
			(_nw14).ArraySet1(_108_v38, 25)
			(_nw14).ArraySet1(_101_v33, 26)
			_119_v47 = _nw14
			var _120_v48 _dafny.Map
			_ = _120_v48
			_120_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v38.F19, _119_v47)
			_120_v48 = (_120_v48).Update(_108_v38.F19, _119_v47)
			(globalState).F3 = p1
			var _121_v49 *C0
			_ = _121_v49
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__((_dafny.Zero).Minus((_dafny.Zero).Minus((_101_v33).Fm3(globalState))))
			_121_v49 = _nw15
			var _122_v50 *C0
			_ = _122_v50
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__(_108_v38.F19)
			_122_v50 = _nw16
			(globalState).F3 = !(p1) || (p0)
		}
		var _123_v51 _dafny.MultiSet
		_ = _123_v51
		_123_v51 = _dafny.MultiSetOf(_dafny.IntOfInt64(103), _108_v38.F19)
		_123_v51 = _123_v51
	} else {
		var _124_v52 _dafny.Sequence
		_ = _124_v52
		_124_v52 = _dafny.SeqOf(p3)
		var _125_v53 *C0
		_ = _125_v53
		var _nw17 *C0 = New_C0_()
		_ = _nw17
		_nw17.Ctor__(Companion_Default___.Fm0(_dafny.IntOfUint32((_124_v52).Cardinality()), _36_v4, _dafny.SeqOf(_31_v3, _31_v3, _31_v3, _31_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_126_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_127_i5 _dafny.Int) _dafny.CodePoint {
				return _126_v4
			}
		})(_36_v4)))), p2, globalState))
		_125_v53 = _nw17
		var _128_v54 *C0
		_ = _128_v54
		var _nw18 *C0 = New_C0_()
		_ = _nw18
		_nw18.Ctor__(_125_v53.F19)
		_128_v54 = _nw18
		var _129_v55 _dafny.Set
		_ = _129_v55
		_129_v55 = _dafny.SetOf(_128_v54.F19)
		if (_129_v55).IsDisjointFrom((Companion_Default___.Fm11(_128_v54.F19, globalState)).Union(_129_v55)) {
			var _130_v56 _dafny.Set
			_ = _130_v56
			_130_v56 = _dafny.SetOf(p0)
			var _131_v57 _dafny.Map
			_ = _131_v57
			_131_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v4, (_130_v56).Cardinality())
			var _rhs17 *C0 = _128_v54
			_ = _rhs17
			var _rhs18 bool = ((_dafny.Zero).Minus((func() _dafny.Int {
				if (_131_v57).Contains(_36_v4) {
					return (_131_v57).Get(_36_v4).(_dafny.Int)
				}
				return _125_v53.F19
			})())).Cmp(_128_v54.F19) != 0
			_ = _rhs18
			var _rhs19 _dafny.Int = _128_v54.F19
			_ = _rhs19
			var _lhs16 *GlobalState = globalState
			_ = _lhs16
			var _lhs17 *GlobalState = globalState
			_ = _lhs17
			_128_v54 = _rhs17
			_lhs16.F3 = _rhs18
			_lhs17.F8 = _rhs19
			var _132_v58 *C0
			_ = _132_v58
			var _nw19 *C0 = New_C0_()
			_ = _nw19
			_nw19.Ctor__((_128_v54.F19).Plus(_125_v53.F19))
			_132_v58 = _nw19
			var _133_v59 _dafny.Map
			_ = _133_v59
			_133_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _128_v54)
			_133_v59 = (_133_v59).Merge(_133_v59)
			(globalState).F3 = ((_dafny.IntOfUint32((_31_v3).Cardinality())).Cmp(_dafny.IntOfInt64(-942)) >= 0) && (p0)
			(globalState).F2 = _125_v53.F19
		} else {
			(globalState).F8 = _125_v53.F19
			var _134_v60 D6
			_ = _134_v60
			_134_v60 = Companion_D6_.Create_DC23_(p1, _125_v53.F19, _125_v53.F19)
			var _135_v61 _dafny.MultiSet
			_ = _135_v61
			_135_v61 = _dafny.MultiSetOf(p2, _dafny.IntOfUint32((_31_v3).Cardinality()), _125_v53.F19)
			var _136_v62 _dafny.Map
			_ = _136_v62
			_136_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v54.F19, p3)
			var _137_v63 _dafny.Map
			_ = _137_v63
			_137_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_136_v62).Contains(_125_v53.F19) {
					return (_136_v62).Get(_125_v53.F19).(bool)
				}
				return p0
			})(), p0)
			var _pat_let_tv16 = _128_v54
			_ = _pat_let_tv16
			var _pat_let_tv17 = _135_v61
			_ = _pat_let_tv17
			var _138_v64 _dafny.Array
			_ = _138_v64
			var _nwElement0_5 D6 = _134_v60
			_ = _nwElement0_5
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(29))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_5, 0)
			(_nw20).ArraySet1(_134_v60, 1)
			(_nw20).ArraySet1(_134_v60, 2)
			(_nw20).ArraySet1(_134_v60, 3)
			(_nw20).ArraySet1(_134_v60, 4)
			(_nw20).ArraySet1(_134_v60, 5)
			(_nw20).ArraySet1(_134_v60, 6)
			(_nw20).ArraySet1(_134_v60, 7)
			(_nw20).ArraySet1(_134_v60, 8)
			(_nw20).ArraySet1(_134_v60, 9)
			(_nw20).ArraySet1(func(_pat_let2_0 D6) D6 {
				return func(_139_dt__update__tmp_h1 D6) D6 {
					return func(_pat_let3_0 _dafny.Int) D6 {
						return func(_140_dt__update_hcf33_h0 _dafny.Int) D6 {
							return Companion_D6_.Create_DC23_((_139_dt__update__tmp_h1).Dtor_cf31(), (_139_dt__update__tmp_h1).Dtor_cf32(), _140_dt__update_hcf33_h0)
						}(_pat_let3_0)
					}(_pat_let_tv16.F19)
				}(_pat_let2_0)
			}(Companion_D6_.Create_DC23_(p3, _dafny.IntOfUint32((_31_v3).Cardinality()), _125_v53.F19)), 10)
			(_nw20).ArraySet1(_134_v60, 11)
			(_nw20).ArraySet1(func(_pat_let4_0 D6) D6 {
				return func(_141_dt__update__tmp_h2 D6) D6 {
					return func(_pat_let5_0 _dafny.Int) D6 {
						return func(_142_dt__update_hcf32_h0 _dafny.Int) D6 {
							return Companion_D6_.Create_DC23_((_141_dt__update__tmp_h2).Dtor_cf31(), _142_dt__update_hcf32_h0, (_141_dt__update__tmp_h2).Dtor_cf33())
						}(_pat_let5_0)
					}((_pat_let_tv17).Cardinality())
				}(_pat_let4_0)
			}(_134_v60), 12)
			(_nw20).ArraySet1(_134_v60, 13)
			(_nw20).ArraySet1(Companion_D6_.Create_DC23_(p1, _125_v53.F19, _125_v53.F19), 14)
			(_nw20).ArraySet1(Companion_D6_.Create_DC23_((func() bool {
				if (_137_v63).Contains(p0) {
					return (_137_v63).Get(p0).(bool)
				}
				return p0
			})(), _dafny.IntOfInt64(-328), (_dafny.Zero).Minus(_128_v54.F19)), 15)
			(_nw20).ArraySet1((func() D6 {
				if p0 {
					return _134_v60
				}
				return _134_v60
			})(), 16)
			(_nw20).ArraySet1((func() D6 {
				if false {
					return _134_v60
				}
				return _134_v60
			})(), 17)
			(_nw20).ArraySet1(Companion_D6_.Create_DC23_(p0, p2, _125_v53.F19), 18)
			(_nw20).ArraySet1(_134_v60, 19)
			(_nw20).ArraySet1(_134_v60, 20)
			(_nw20).ArraySet1(_134_v60, 21)
			(_nw20).ArraySet1(_134_v60, 22)
			(_nw20).ArraySet1(_134_v60, 23)
			(_nw20).ArraySet1(_134_v60, 24)
			(_nw20).ArraySet1((func() D6 {
				if p3 {
					return _134_v60
				}
				return _134_v60
			})(), 25)
			(_nw20).ArraySet1(Companion_Default___.Fm18(_128_v54.F19, p1, globalState), 26)
			(_nw20).ArraySet1(_134_v60, 27)
			(_nw20).ArraySet1(_134_v60, 28)
			_138_v64 = _nw20
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_138_v64), 0))
			_ = _index20
			(_138_v64).ArraySet1((func() D6 {
				if p3 {
					return _134_v60
				}
				return Companion_D6_.Create_DC23_(p0, _128_v54.F19, (func() _dafny.Int {
					if (_28_v0).Contains((func() bool {
						if (_136_v62).Contains(_125_v53.F19) {
							return (_136_v62).Get(_125_v53.F19).(bool)
						}
						return !(p0)
					})()) {
						return (_28_v0).Multiplicity((func() bool {
							if (_136_v62).Contains(_125_v53.F19) {
								return (_136_v62).Get(_125_v53.F19).(bool)
							}
							return !(p0)
						})())
					}
					return _128_v54.F19
				})())
			})(), (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_138_v64), 0))
			_ = _index21
			(_138_v64).ArraySet1(Companion_Default___.Fm18(_125_v53.F19, false, globalState), (_index21).Int())
			(globalState).F3 = _dafny.Companion_Sequence_.IsProperPrefixOf(_31_v3, _dafny.UnicodeSeqOfUtf8Bytes("hnidaeeo"))
			var _143_v65 _dafny.Map
			_ = _143_v65
			_143_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_dafny.Zero).Minus(_125_v53.F19))
			_143_v65 = (_143_v65).Update(p1, _125_v53.F19)
			var _144_v66 _dafny.Array
			_ = _144_v66
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw21
			_144_v66 = _nw21
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_144_v66), 0))
			_ = _index22
			(_144_v66).ArraySet1(_128_v54.F19, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_144_v66), 0))
			_ = _index23
			(_144_v66).ArraySet1(_128_v54.F19, (_index23).Int())
		}
		var _145_v67 _dafny.Array
		_ = _145_v67
		var _nw22 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
		_ = _nw22
		_145_v67 = _nw22
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_145_v67), 0))
		_ = _index24
		(_145_v67).ArraySet1(_125_v53, (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_145_v67), 0))
		_ = _index25
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(115), (_dafny.Zero).Minus(p2)))
		(_145_v67).ArraySet1(_nw23, (_index25).Int())
		var _146_v68 _dafny.Map
		_ = _146_v68
		_146_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(198))
		var _147_v69 _dafny.Sequence
		_ = _147_v69
		_147_v69 = _dafny.SeqOf(_31_v3)
		_146_v68 = (_146_v68).Update(((_125_v53).Fm4(p2, _128_v54.F19, p0, p0, globalState)).Cmp(_128_v54.F19) != 0, _dafny.IntOfUint32(((_147_v69).Select((Companion_Default___.SafeIndex(_125_v53.F19, _dafny.IntOfUint32((_147_v69).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
	}
	var _148_v70 _dafny.Sequence
	_ = _148_v70
	_148_v70 = _dafny.SeqOf(false, p0)
	var _149_v71 _dafny.Map
	_ = _149_v71
	_149_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
	var _150_v72 _dafny.MultiSet
	_ = _150_v72
	_150_v72 = _dafny.MultiSetOf(_149_v71, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3))
	var _151_v73 _dafny.Map
	_ = _151_v73
	_151_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_150_v72).Cardinality())
	var _152_v74 _dafny.Map
	_ = _152_v74
	_152_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_151_v73).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)))
	var _rhs20 _dafny.Sequence = _148_v70
	_ = _rhs20
	var _rhs21 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_31_v3, _dafny.Companion_Sequence_.Concatenate(_31_v3, _31_v3))
	_ = _rhs21
	var _rhs22 _dafny.Map = _152_v74
	_ = _rhs22
	var _rhs23 D5 = _37_v5
	_ = _rhs23
	_148_v70 = _rhs20
	r1 = _rhs21
	_152_v74 = _rhs22
	_37_v5 = _rhs23
	var _153_v75 *C0
	_ = _153_v75
	var _nw24 *C0 = New_C0_()
	_ = _nw24
	_nw24.Ctor__(p2)
	_153_v75 = _nw24
	r0 = (_dafny.IntOfUint32((_dafny.SeqOf(_153_v75)).Cardinality())).Times(p2)
	r1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_31_v3, _dafny.UnicodeSeqOfUtf8Bytes("ey")), Companion_Default___.Fm5(_153_v75.F19, globalState))
	r2 = _153_v75.F19
	var _154_v76 _dafny.Set
	_ = _154_v76
	_154_v76 = _dafny.SetOf(p0, true, p1)
	r3 = (_28_v0).Update((p2).Cmp((_154_v76).Cardinality()) < 0, Companion_Default___.Abs(Companion_Default___.SafeModuloInt(p2, p2)))
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _155_v0 bool
	_ = _155_v0
	_155_v0 = false
	var _156_v1 _dafny.Sequence
	_ = _156_v1
	_156_v1 = _dafny.SeqOf(_155_v0)
	var _157_globalState *GlobalState
	_ = _157_globalState
	var _nw25 *GlobalState = New_GlobalState_()
	_ = _nw25
	_nw25.Ctor__(_dafny.IntOfInt64(-782), false, _dafny.IntOfInt64(351), false, _dafny.IntOfInt64(-139), _dafny.IntOfInt64(973), false, _dafny.IntOfInt64(-608), _dafny.IntOfInt64(117), _dafny.IntOfInt64(-409), _dafny.UnicodeSeqOfUtf8Bytes("adqetglt"), false, _dafny.IntOfInt64(-956), true, _dafny.IntOfInt64(26), false, _dafny.IntOfInt64(957), _dafny.IntOfInt64(182), _156_v1)
	_157_globalState = _nw25
	var _158_v2 _dafny.Int
	_ = _158_v2
	_158_v2 = _dafny.IntOfInt64(270)
	var _159_v3 _dafny.MultiSet
	_ = _159_v3
	_159_v3 = _dafny.MultiSetOf(_155_v0)
	var _hi1 _dafny.Int = (_dafny.Zero).Minus((_158_v2).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_159_v3).Cardinality()))))
	_ = _hi1
	for _160_i0 := _dafny.IntOfInt64(115); _160_i0.Cmp(_hi1) < 0; _160_i0 = _160_i0.Plus(_dafny.One) {
		var _161_v4 _dafny.CodePoint
		_ = _161_v4
		_161_v4 = _dafny.CodePoint('f')
		var _162_v5 _dafny.Sequence
		_ = _162_v5
		_162_v5 = _dafny.UnicodeSeqOfUtf8Bytes("oiwhio")
		var _163_v6 _dafny.Array
		_ = _163_v6
		var _nwElement0_6 _dafny.Sequence = _162_v5
		_ = _nwElement0_6
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(3))
		_ = _nw26
		(_nw26).ArraySet1(_nwElement0_6, 0)
		(_nw26).ArraySet1(_162_v5, 1)
		(_nw26).ArraySet1(_162_v5, 2)
		_163_v6 = _nw26
		var _164_v7 _dafny.Set
		_ = _164_v7
		_164_v7 = _dafny.SetOf(_155_v0)
		var _165_v8 _dafny.Sequence
		_ = _165_v8
		_165_v8 = _dafny.SeqOf(_162_v5)
		var _166_v9 _dafny.Map
		_ = _166_v9
		_166_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_164_v7).Cardinality(), _161_v4, _165_v8, _dafny.IntOfInt64(591), _157_globalState), _163_v6)
		var _167_v10 _dafny.Map
		_ = _167_v10
		_167_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, _155_v0)
		var _168_v11 _dafny.Array
		_ = _168_v11
		var _nwElement0_7 _dafny.Array = _163_v6
		_ = _nwElement0_7
		var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(19))
		_ = _nw27
		(_nw27).ArraySet1(_nwElement0_7, 0)
		(_nw27).ArraySet1(_163_v6, 1)
		(_nw27).ArraySet1(_163_v6, 2)
		(_nw27).ArraySet1(_163_v6, 3)
		(_nw27).ArraySet1(_163_v6, 4)
		(_nw27).ArraySet1(_163_v6, 5)
		(_nw27).ArraySet1(_163_v6, 6)
		(_nw27).ArraySet1(_163_v6, 7)
		(_nw27).ArraySet1(_163_v6, 8)
		(_nw27).ArraySet1(_163_v6, 9)
		(_nw27).ArraySet1((func() _dafny.Array {
			if (_166_v9).Contains((_167_v10).Cardinality()) {
				return (_166_v9).Get((_167_v10).Cardinality()).(_dafny.Array)
			}
			return _163_v6
		})(), 10)
		(_nw27).ArraySet1((func() _dafny.Array {
			if _155_v0 {
				return _163_v6
			}
			return _163_v6
		})(), 11)
		(_nw27).ArraySet1(_163_v6, 12)
		(_nw27).ArraySet1(_163_v6, 13)
		(_nw27).ArraySet1(_163_v6, 14)
		(_nw27).ArraySet1(_163_v6, 15)
		(_nw27).ArraySet1(_163_v6, 16)
		(_nw27).ArraySet1(_163_v6, 17)
		(_nw27).ArraySet1(_163_v6, 18)
		_168_v11 = _nw27
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_168_v11), 0))
		_ = _index26
		(_168_v11).ArraySet1(_163_v6, (_index26).Int())
		var _169_v12 _dafny.Array
		_ = _169_v12
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(26))
		_ = _nw28
		_169_v12 = _nw28
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_169_v12), 0))
		_ = _index27
		(_169_v12).ArraySet1(_164_v7, (_index27).Int())
		var _170_v13 _dafny.Map
		_ = _170_v13
		_170_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, _164_v7)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_168_v11), 0))
		_ = _index28
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_169_v12), 0))
		_ = _index29
		var _rhs24 _dafny.CodePoint = _dafny.CodePoint('r')
		_ = _rhs24
		var _rhs25 _dafny.Array = _163_v6
		_ = _rhs25
		var _rhs26 bool = _155_v0
		_ = _rhs26
		var _rhs27 _dafny.Set = (_164_v7).Intersection((func() _dafny.Set {
			if (_170_v13).Contains(_160_i0) {
				return (_170_v13).Get(_160_i0).(_dafny.Set)
			}
			return _164_v7
		})())
		_ = _rhs27
		var _lhs18 _dafny.Array = _168_v11
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_168_v11), 0))
		_ = _lhs19
		var _lhs20 *GlobalState = _157_globalState
		_ = _lhs20
		var _lhs21 _dafny.Array = _169_v12
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_169_v12), 0))
		_ = _lhs22
		_161_v4 = _rhs24
		(_lhs18).ArraySet1(_rhs25, (_lhs19).Int())
		_lhs20.F3 = _rhs26
		(_lhs21).ArraySet1(_rhs27, (_lhs22).Int())
		var _171_v14 _dafny.Sequence
		_ = _171_v14
		_171_v14 = _dafny.SeqOf(_dafny.IntOfInt64(172))
		var _172_v15 _dafny.Map
		_ = _172_v15
		_172_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_171_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-542), _dafny.IntOfUint32((_171_v14).Cardinality()))).Uint32()).(_dafny.Int), _160_i0)
		var _rhs28 _dafny.Int = _158_v2
		_ = _rhs28
		var _rhs29 _dafny.Int = (_dafny.Zero).Minus((((_dafny.Zero).Minus((func() _dafny.Int {
			if (_172_v15).Contains(_158_v2) {
				return (_172_v15).Get(_158_v2).(_dafny.Int)
			}
			return _158_v2
		})())).Plus(_158_v2)).Times(_160_i0))
		_ = _rhs29
		var _rhs30 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gydcso")).Cardinality()))).Minus(_dafny.IntOfInt64(-566))
		_ = _rhs30
		var _rhs31 _dafny.Int = _158_v2
		_ = _rhs31
		var _lhs23 *GlobalState = _157_globalState
		_ = _lhs23
		var _lhs24 *GlobalState = _157_globalState
		_ = _lhs24
		var _lhs25 *GlobalState = _157_globalState
		_ = _lhs25
		var _lhs26 *GlobalState = _157_globalState
		_ = _lhs26
		_lhs23.F12 = _rhs28
		_lhs24.F8 = _rhs29
		_lhs25.F2 = _rhs30
		_lhs26.F12 = _rhs31
		var _173_v16 _dafny.Array
		_ = _173_v16
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_2
		var _nw29 _dafny.Array
		_ = _nw29
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw29 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Set = (func(_174_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
				return func(_175_i1 _dafny.Int) _dafny.Set {
					return (_dafny.SetOf(_174_v5)).Difference(_dafny.SetOf(_174_v5, _dafny.UnicodeSeqOfUtf8Bytes("kvf")))
				}
			})(_162_v5)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw29 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw29).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw29).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_173_v16 = _nw29
		var _176_v17 _dafny.Set
		_ = _176_v17
		_176_v17 = _dafny.SetOf(_162_v5, _162_v5)
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_173_v16), 0))
		_ = _index30
		(_173_v16).ArraySet1(_176_v17, (_index30).Int())
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_173_v16), 0))
		_ = _index31
		(_173_v16).ArraySet1(((_176_v17).Intersection(_176_v17)).Intersection(_176_v17), (_index31).Int())
		(_157_globalState).F8 = ((_dafny.Zero).Minus(_158_v2)).Minus(_158_v2)
	}
	var _177_v18 _dafny.Sequence
	_ = _177_v18
	_177_v18 = _dafny.UnicodeSeqOfUtf8Bytes("gcqnkipj")
	_177_v18 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(650))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_178_i2 _dafny.Int) _dafny.CodePoint {
		return (func() _dafny.CodePoint {
			if true {
				return _dafny.CodePoint('g')
			}
			return _dafny.CodePoint('x')
		})()
	}))
	if _155_v0 {
		var _179_v19 _dafny.Int
		_ = _179_v19
		var _180_v20 _dafny.Sequence
		_ = _180_v20
		var _181_v21 _dafny.Int
		_ = _181_v21
		var _182_v22 _dafny.MultiSet
		_ = _182_v22
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.Sequence
		_ = _out1
		var _out2 _dafny.Int
		_ = _out2
		var _out3 _dafny.MultiSet
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M0(_155_v0, _155_v0, _158_v2, _155_v0, _157_globalState)
		_179_v19 = _out0
		_180_v20 = _out1
		_181_v21 = _out2
		_182_v22 = _out3
		_155_v0 = !(_155_v0)
		var _183_v23 _dafny.CodePoint
		_ = _183_v23
		_183_v23 = _dafny.CodePoint('q')
		var _184_v24 _dafny.Map
		_ = _184_v24
		_184_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, _183_v23)
		var _185_v25 _dafny.Set
		_ = _185_v25
		_185_v25 = _dafny.SetOf(_184_v24)
		var _186_v26 D0
		_ = _186_v26
		_186_v26 = Companion_D0_.Create_DC0_(_158_v2)
		var _pat_let_tv18 = _181_v21
		_ = _pat_let_tv18
		var _187_v27 _dafny.Array
		_ = _187_v27
		var _nwElement0_8 _dafny.Int = (_185_v25).Cardinality()
		_ = _nwElement0_8
		var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(5))
		_ = _nw30
		(_nw30).ArraySet1(_nwElement0_8, 0)
		(_nw30).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_177_v18).Cardinality())), 1)
		(_nw30).ArraySet1(_179_v19, 2)
		(_nw30).ArraySet1(_dafny.IntOfInt64(694), 3)
		(_nw30).ArraySet1((func(_pat_let6_0 D0) D0 {
			return func(_188_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let7_0 _dafny.Int) D0 {
					return func(_189_dt__update_hcf0_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_189_dt__update_hcf0_h0)
					}(_pat_let7_0)
				}(_pat_let_tv18)
			}(_pat_let6_0)
		}(_186_v26)).Dtor_cf0(), 4)
		_187_v27 = _nw30
		var _190_v28 _dafny.MultiSet
		_ = _190_v28
		_190_v28 = _dafny.MultiSetOf(_187_v27)
		_190_v28 = (_190_v28).Intersection(_190_v28)
		var _191_v29 _dafny.Map
		_ = _191_v29
		_191_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, _182_v22)
		var _192_v30 D0
		_ = _192_v30
		_192_v30 = Companion_D0_.Create_DC2_(_155_v0, (_179_v19).Cmp(_158_v2) >= 0, _181_v21, true)
		var _rhs32 _dafny.Int = _179_v19
		_ = _rhs32
		var _rhs33 _dafny.Map = (_191_v29).Update(_179_v19, _dafny.MultiSetOf(_155_v0))
		_ = _rhs33
		var _rhs34 D0 = _192_v30
		_ = _rhs34
		var _lhs27 *GlobalState = _157_globalState
		_ = _lhs27
		_lhs27.F2 = _rhs32
		_191_v29 = _rhs33
		_192_v30 = _rhs34
		var _193_v31 _dafny.Array
		_ = _193_v31
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
		_ = _nw31
		_193_v31 = _nw31
		var _194_v32 _dafny.Sequence
		_ = _194_v32
		_194_v32 = _dafny.SeqOf(_177_v18, _177_v18)
		var _195_v33 _dafny.Array
		_ = _195_v33
		var _nwElement0_9 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_196_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_197_i3 _dafny.Int) _dafny.CodePoint {
				return _196_v23
			}
		})(_183_v23))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.Fm0(_179_v19, _dafny.CodePoint('r'), _194_v32, _179_v19, _157_globalState)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_198_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_199_i3 _dafny.Int) _dafny.CodePoint {
				return _198_v23
			}
		})(_183_v23)))).Cardinality()))).Uint32(), _dafny.CodePoint('h'))
		_ = _nwElement0_9
		var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(23))
		_ = _nw32
		(_nw32).ArraySet1(_nwElement0_9, 0)
		(_nw32).ArraySet1(_177_v18, 1)
		(_nw32).ArraySet1(_177_v18, 2)
		(_nw32).ArraySet1(_180_v20, 3)
		(_nw32).ArraySet1(_180_v20, 4)
		(_nw32).ArraySet1(_177_v18, 5)
		(_nw32).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(238))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_200_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_201_i4 _dafny.Int) _dafny.CodePoint {
				return _200_v23
			}
		})(_183_v23))), 6)
		(_nw32).ArraySet1(_177_v18, 7)
		(_nw32).ArraySet1(_177_v18, 8)
		(_nw32).ArraySet1(_180_v20, 9)
		(_nw32).ArraySet1(_177_v18, 10)
		(_nw32).ArraySet1(_180_v20, 11)
		(_nw32).ArraySet1(_180_v20, 12)
		(_nw32).ArraySet1(_177_v18, 13)
		(_nw32).ArraySet1(_180_v20, 14)
		(_nw32).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xeumplyd"), 15)
		(_nw32).ArraySet1(_177_v18, 16)
		(_nw32).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(550))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_202_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_203_i5 _dafny.Int) _dafny.CodePoint {
				return _202_v23
			}
		})(_183_v23))), 17)
		(_nw32).ArraySet1(_177_v18, 18)
		(_nw32).ArraySet1(_177_v18, 19)
		(_nw32).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(885))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_204_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_205_i6 _dafny.Int) _dafny.CodePoint {
				return _204_v23
			}
		})(_183_v23))), 20)
		(_nw32).ArraySet1(_177_v18, 21)
		(_nw32).ArraySet1(_180_v20, 22)
		_195_v33 = _nw32
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_193_v31), 0))
		_ = _index32
		(_193_v31).ArraySet1(_195_v33, (_index32).Int())
		var _206_v34 _dafny.MultiSet
		_ = _206_v34
		_206_v34 = _dafny.MultiSetOf(_186_v26)
		var _207_v35 _dafny.MultiSet
		_ = _207_v35
		_207_v35 = _dafny.MultiSetOf(_179_v19, _181_v21)
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_187_v27), 0))
		_ = _index33
		(_187_v27).ArraySet1(Companion_Default___.SafeModuloInt(((_206_v34).Update(_186_v26, Companion_Default___.Abs((_182_v22).Cardinality()))).Cardinality(), Companion_Default___.Fm0(_158_v2, _183_v23, _194_v32, (_207_v35).Cardinality(), _157_globalState)), (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_193_v31), 0))
		_ = _index34
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_187_v27), 0))
		_ = _index35
		var _rhs35 _dafny.Array = _187_v27
		_ = _rhs35
		var _rhs36 _dafny.Int = _158_v2
		_ = _rhs36
		var _rhs37 _dafny.Array = _195_v33
		_ = _rhs37
		var _rhs38 _dafny.Int = (_dafny.Zero).Minus(_158_v2)
		_ = _rhs38
		var _lhs28 *GlobalState = _157_globalState
		_ = _lhs28
		var _lhs29 _dafny.Array = _193_v31
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_193_v31), 0))
		_ = _lhs30
		var _lhs31 _dafny.Array = _187_v27
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_187_v27), 0))
		_ = _lhs32
		_187_v27 = _rhs35
		_lhs28.F0 = _rhs36
		(_lhs29).ArraySet1(_rhs37, (_lhs30).Int())
		(_lhs31).ArraySet1(_rhs38, (_lhs32).Int())
	} else {
		var _208_v36 _dafny.CodePoint
		_ = _208_v36
		_208_v36 = _dafny.CodePoint('s')
		_156_v1 = Companion_Default___.Fm1(_208_v36, _155_v0, _157_globalState)
		if Companion_Default___.Fm2((_155_v0) == (_155_v0), _157_globalState) {
			var _209_v37 *C0
			_ = _209_v37
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__((_158_v2).Plus((_dafny.Zero).Minus(_158_v2)))
			_209_v37 = _nw33
			var _210_v38 _dafny.Sequence
			_ = _210_v38
			_210_v38 = _dafny.SeqOf(_177_v18)
			var _211_v39 _dafny.Set
			_ = _211_v39
			_211_v39 = _dafny.SetOf(_155_v0)
			(_157_globalState).F0 = Companion_Default___.Fm0((_209_v37.F19).Times(_209_v37.F19), _208_v36, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_208_v36), _177_v18), _210_v38), ((_211_v39).Intersection(_211_v39)).Cardinality(), _157_globalState)
			var _212_v40 _dafny.Int
			_ = _212_v40
			var _213_v41 _dafny.Sequence
			_ = _213_v41
			var _214_v42 _dafny.Int
			_ = _214_v42
			var _215_v43 _dafny.MultiSet
			_ = _215_v43
			var _out4 _dafny.Int
			_ = _out4
			var _out5 _dafny.Sequence
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.MultiSet
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0(_155_v0, _dafny.Companion_Sequence_.IsPrefixOf(_177_v18, _177_v18), _209_v37.F19, _155_v0, _157_globalState)
			_212_v40 = _out4
			_213_v41 = _out5
			_214_v42 = _out6
			_215_v43 = _out7
			_158_v2 = (_dafny.Zero).Minus(_212_v40)
			var _216_v44 _dafny.Array
			_ = _216_v44
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw34
			_216_v44 = _nw34
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_216_v44), 0))
			_ = _index36
			(_216_v44).ArraySet1(true, (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_216_v44), 0))
			_ = _index37
			(_216_v44).ArraySet1(_155_v0, (_index37).Int())
		} else {
			var _217_v45 _dafny.Map
			_ = _217_v45
			_217_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, (_155_v0) || (_155_v0))
			_217_v45 = (_217_v45).Update(_158_v2, Companion_Default___.Fm2(false, _157_globalState))
			var _218_v46 _dafny.Array
			_ = _218_v46
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw35
			_218_v46 = _nw35
			_218_v46 = _218_v46
			var _219_v47 _dafny.Set
			_ = _219_v47
			_219_v47 = _dafny.SetOf((func() bool {
				if _155_v0 {
					return _155_v0
				}
				return _155_v0
			})())
			var _220_v48 *C0
			_ = _220_v48
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__(_158_v2)
			_220_v48 = _nw36
			var _221_v49 _dafny.Map
			_ = _221_v49
			_221_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_208_v36, Companion_Default___.Fm2(_155_v0, _157_globalState))
			var _222_v50 _dafny.Set
			_ = _222_v50
			_222_v50 = _dafny.SetOf(_220_v48, _220_v48)
			var _223_v51 _dafny.Map
			_ = _223_v51
			_223_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, _208_v36)
			var _224_v52 _dafny.Sequence
			_ = _224_v52
			_224_v52 = _dafny.SeqOf(_223_v51, _223_v51)
			var _225_v53 _dafny.Map
			_ = _225_v53
			_225_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _dafny.IntOfUint32((_156_v1).Cardinality()))
			var _226_v54 _dafny.Set
			_ = _226_v54
			_226_v54 = _dafny.SetOf((_224_v52).Select((Companion_Default___.SafeIndex(_158_v2, _dafny.IntOfUint32((_224_v52).Cardinality()))).Uint32()).(_dafny.Map), (_223_v51).Update((_225_v53).Cardinality(), _208_v36), (_223_v51).Update(_dafny.IntOfInt64(-407), _208_v36))
			var _227_v55 _dafny.Sequence
			_ = _227_v55
			_227_v55 = _dafny.SeqOf((_dafny.Zero).Minus(_158_v2))
			var _228_v56 D0
			_ = _228_v56
			_228_v56 = Companion_D0_.Create_DC2_(!(!(_155_v0)), _155_v0, _158_v2, _155_v0)
			var _229_v57 _dafny.Array
			_ = _229_v57
			var _nwElement0_10 bool = _155_v0
			_ = _nwElement0_10
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(16))
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_10, 0)
			(_nw37).ArraySet1(Companion_Default___.Fm2((func() bool {
				if (_221_v49).Contains(_208_v36) {
					return (_221_v49).Get(_208_v36).(bool)
				}
				return _155_v0
			})(), _157_globalState), 1)
			(_nw37).ArraySet1((_222_v50).IsProperSubsetOf(_222_v50), 2)
			(_nw37).ArraySet1((_226_v54).Contains(_223_v51), 3)
			(_nw37).ArraySet1((_156_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_227_v55).Cardinality()), _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool), 4)
			(_nw37).ArraySet1(_155_v0, 5)
			(_nw37).ArraySet1(_155_v0, 6)
			(_nw37).ArraySet1((_228_v56).Dtor_cf4(), 7)
			(_nw37).ArraySet1(!((_155_v0) == (_155_v0)), 8)
			(_nw37).ArraySet1(!_dafny.Companion_Sequence_.Equal(_227_v55, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_230_v48 *C0) func(_dafny.Int) _dafny.Int {
				return func(_231_i7 _dafny.Int) _dafny.Int {
					return _230_v48.F19
				}
			})(_220_v48)))), 9)
			(_nw37).ArraySet1(_155_v0, 10)
			(_nw37).ArraySet1((_155_v0) || (_155_v0), 11)
			(_nw37).ArraySet1(_155_v0, 12)
			(_nw37).ArraySet1(true, 13)
			(_nw37).ArraySet1(_155_v0, 14)
			(_nw37).ArraySet1(_155_v0, 15)
			_229_v57 = _nw37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_229_v57), 0))
			_ = _index38
			(_229_v57).ArraySet1(_155_v0, (_index38).Int())
			var _232_v58 _dafny.Sequence
			_ = _232_v58
			_232_v58 = _dafny.SeqOf(_dafny.SeqOf(_208_v36), _177_v18, _177_v18)
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_229_v57), 0))
			_ = _index39
			var _rhs39 _dafny.Set = ((_219_v47).Intersection(_219_v47)).Union(_219_v47)
			_ = _rhs39
			var _rhs40 _dafny.Int = (_158_v2).Minus(Companion_Default___.Fm0(_220_v48.F19, _208_v36, _232_v58, _dafny.IntOfUint32((_177_v18).Cardinality()), _157_globalState))
			_ = _rhs40
			var _rhs41 *C0 = _220_v48
			_ = _rhs41
			var _rhs42 bool = !((_220_v48.F19).Cmp((_225_v53).Cardinality()) == 0) || (_155_v0)
			_ = _rhs42
			var _lhs33 *GlobalState = _157_globalState
			_ = _lhs33
			var _lhs34 _dafny.Array = _229_v57
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_229_v57), 0))
			_ = _lhs35
			_219_v47 = _rhs39
			_lhs33.F2 = _rhs40
			_220_v48 = _rhs41
			(_lhs34).ArraySet1(_rhs42, (_lhs35).Int())
			var _233_v59 _dafny.Int
			_ = _233_v59
			var _234_v60 _dafny.Sequence
			_ = _234_v60
			var _235_v61 _dafny.Int
			_ = _235_v61
			var _236_v62 _dafny.MultiSet
			_ = _236_v62
			var _out8 _dafny.Int
			_ = _out8
			var _out9 _dafny.Sequence
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.MultiSet
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0((_220_v48.F19).Cmp(_158_v2) < 0, !_dafny.Companion_Sequence_.Contains(_156_v1, Companion_Default___.Fm2(_155_v0, _157_globalState)), (_dafny.Zero).Minus(_220_v48.F19), _155_v0, _157_globalState)
			_233_v59 = _out8
			_234_v60 = _out9
			_235_v61 = _out10
			_236_v62 = _out11
			(_157_globalState).F3 = (_229_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_229_v57), 0))).Int()).(bool)
		}
		(_157_globalState).F2 = _158_v2
		(_157_globalState).F2 = _158_v2
		(_157_globalState).F3 = _155_v0
	}
	var _237_v63 _dafny.Map
	_ = _237_v63
	_237_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_155_v0, _155_v0, _155_v0, _155_v0), (_dafny.IntOfUint32((_177_v18).Cardinality())).Cmp(_158_v2) == 0)
	var _238_v64 _dafny.Sequence
	_ = _238_v64
	_238_v64 = _dafny.SeqOf(_158_v2)
	var _239_v65 *C0
	_ = _239_v65
	var _nw38 *C0 = New_C0_()
	_ = _nw38
	_nw38.Ctor__(_158_v2)
	_239_v65 = _nw38
	var _240_v66 _dafny.Map
	_ = _240_v66
	_240_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, _239_v65)
	var _241_v67 _dafny.Map
	_ = _241_v67
	_241_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_240_v66).Cardinality(), _158_v2)
	var _242_v68 *C0
	_ = _242_v68
	var _nw39 *C0 = New_C0_()
	_ = _nw39
	_nw39.Ctor__((_238_v64).Select((Companion_Default___.SafeIndex((_241_v67).Cardinality(), _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int))
	_242_v68 = _nw39
	var _243_v69 _dafny.Map
	_ = _243_v69
	_243_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v65, _155_v0)
	_237_v63 = (_237_v63).Update(_dafny.SetOf(!((func() bool {
		if (_243_v69).Contains(_239_v65) {
			return (_243_v69).Get(_239_v65).(bool)
		}
		return _155_v0
	})())), _155_v0)
	_155_v0 = _155_v0
	var _244_i8 _dafny.Int
	_ = _244_i8
	_244_i8 = _dafny.Zero
	{
		for _155_v0 {
			{
				if (_244_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_244_i8 = (_244_i8).Plus(_dafny.One)
				if _155_v0 {
					var _245_v70 D0
					_ = _245_v70
					_245_v70 = Companion_D0_.Create_DC2_(true, _155_v0, _242_v68.F19, false)
					var _246_v71 _dafny.Array
					_ = _246_v71
					var _nwElement0_11 bool = _155_v0
					_ = _nwElement0_11
					var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(27))
					_ = _nw40
					(_nw40).ArraySet1(_nwElement0_11, 0)
					(_nw40).ArraySet1(_155_v0, 1)
					(_nw40).ArraySet1(_155_v0, 2)
					(_nw40).ArraySet1(_155_v0, 3)
					(_nw40).ArraySet1(_155_v0, 4)
					(_nw40).ArraySet1((_156_v1).Select((Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool), 5)
					(_nw40).ArraySet1(_155_v0, 6)
					(_nw40).ArraySet1(_155_v0, 7)
					(_nw40).ArraySet1(Companion_Default___.Fm2(!(_155_v0), _157_globalState), 8)
					(_nw40).ArraySet1(_155_v0, 9)
					(_nw40).ArraySet1(_155_v0, 10)
					(_nw40).ArraySet1(_155_v0, 11)
					(_nw40).ArraySet1(_155_v0, 12)
					(_nw40).ArraySet1((_245_v70).Dtor_cf4(), 13)
					(_nw40).ArraySet1(_155_v0, 14)
					(_nw40).ArraySet1(!(_155_v0), 15)
					(_nw40).ArraySet1(_155_v0, 16)
					(_nw40).ArraySet1(_155_v0, 17)
					(_nw40).ArraySet1(_155_v0, 18)
					(_nw40).ArraySet1(_155_v0, 19)
					(_nw40).ArraySet1(_155_v0, 20)
					(_nw40).ArraySet1(_155_v0, 21)
					(_nw40).ArraySet1(_155_v0, 22)
					(_nw40).ArraySet1(!(_155_v0), 23)
					(_nw40).ArraySet1(_155_v0, 24)
					(_nw40).ArraySet1(_155_v0, 25)
					(_nw40).ArraySet1(true, 26)
					_246_v71 = _nw40
					var _247_v72 _dafny.Map
					_ = _247_v72
					_247_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_246_v71, _155_v0)
					_247_v72 = (_247_v72).Update(_246_v71, _155_v0)
					var _248_v73 _dafny.Array
					_ = _248_v73
					var _nw41 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
					_ = _nw41
					_248_v73 = _nw41
					var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_248_v73), 0))
					_ = _index40
					(_248_v73).ArraySet1(_239_v65, (_index40).Int())
					var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_248_v73), 0))
					_ = _index41
					(_248_v73).ArraySet1(_242_v68, (_index41).Int())
					var _249_v74 _dafny.Array
					_ = _249_v74
					var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
					_ = _nw42
					_249_v74 = _nw42
					var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_249_v74), 0))
					_ = _index42
					(_249_v74).ArraySet1(_158_v2, (_index42).Int())
					var _250_v75 _dafny.Array
					_ = _250_v75
					var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
					_ = _nw43
					_250_v75 = _nw43
					var _251_v76 D0
					_ = _251_v76
					_251_v76 = Companion_D0_.Create_DC0_(_239_v65.F19)
					var _252_v77 _dafny.Sequence
					_ = _252_v77
					_252_v77 = _dafny.SeqOf(_156_v1)
					var _253_v78 _dafny.MultiSet
					_ = _253_v78
					_253_v78 = _dafny.MultiSetOf(_156_v1, (_252_v77).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_252_v77).Cardinality()))).Uint32()).(_dafny.Sequence), _156_v1)
					var _254_v79 _dafny.CodePoint
					_ = _254_v79
					_254_v79 = _dafny.CodePoint('s')
					var _255_v80 _dafny.Sequence
					_ = _255_v80
					_255_v80 = _dafny.SeqOf(Companion_Default___.Fm5(_158_v2, _157_globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(382))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg21 _dafny.Int) interface{} {
							return coer21(arg21)
						}
					}((func(_256_v79 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_257_i9 _dafny.Int) _dafny.CodePoint {
							return _256_v79
						}
					})(_254_v79))), _177_v18)
					var _258_v81 _dafny.Map
					_ = _258_v81
					_258_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v73, _250_v75)
					var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_249_v74), 0))
					_ = _index43
					var _rhs43 bool = !((func() bool {
						if _155_v0 {
							return _155_v0
						}
						return _155_v0
					})()) || (_155_v0)
					_ = _rhs43
					var _rhs44 _dafny.Int = _242_v68.F19
					_ = _rhs44
					var _rhs45 _dafny.Int = Companion_Default___.Fm0(((_dafny.MultiSetFromSeq(_252_v77)).Intersection(_253_v78)).Cardinality(), _254_v79, _255_v80, _239_v65.F19, _157_globalState)
					_ = _rhs45
					var _rhs46 _dafny.Array = (func() _dafny.Array {
						if (_258_v81).Contains(_248_v73) {
							return (_258_v81).Get(_248_v73).(_dafny.Array)
						}
						return _250_v75
					})()
					_ = _rhs46
					var _rhs47 D0 = _251_v76
					_ = _rhs47
					var _lhs36 *GlobalState = _157_globalState
					_ = _lhs36
					var _lhs37 *C0 = _242_v68
					_ = _lhs37
					var _lhs38 _dafny.Array = _249_v74
					_ = _lhs38
					var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_249_v74), 0))
					_ = _lhs39
					_lhs36.F3 = _rhs43
					_lhs37.F19 = _rhs44
					(_lhs38).ArraySet1(_rhs45, (_lhs39).Int())
					_250_v75 = _rhs46
					_251_v76 = _rhs47
					(_239_v65).F19 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_242_v68.F19))
					var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_246_v71), 0))
					_ = _index44
					(_246_v71).ArraySet1(_155_v0, (_index44).Int())
					var _259_v82 _dafny.Array
					_ = _259_v82
					var _len0_3 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_3
					var _nw44 _dafny.Array
					_ = _nw44
					if _len0_3.Cmp(_dafny.Zero) == 0 {
						_nw44 = _dafny.NewArray(_len0_3)
					} else {
						var _init3 func(_dafny.Int) _dafny.CodePoint = (func(_260_v79 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_261_i10 _dafny.Int) _dafny.CodePoint {
								return _260_v79
							}
						})(_254_v79)
						_ = _init3
						var _element0_3 = _init3(_dafny.Zero)
						_ = _element0_3
						_nw44 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
						(_nw44).ArraySet1CodePoint(_element0_3, 0)
						var _nativeLen0_3 = (_len0_3).Int()
						_ = _nativeLen0_3
						for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
							(_nw44).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
						}
					}
					_259_v82 = _nw44
					var _262_v83 _dafny.MultiSet
					_ = _262_v83
					_262_v83 = _dafny.MultiSetOf(_259_v82, _259_v82)
					var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_246_v71), 0))
					_ = _index45
					var _rhs48 _dafny.Int = _158_v2
					_ = _rhs48
					var _rhs49 bool = (func() bool {
						if _155_v0 {
							return ((_262_v83).Cardinality()).Cmp(_239_v65.F19) <= 0
						}
						return _155_v0
					})()
					_ = _rhs49
					var _rhs50 bool = (_242_v68.F19).Cmp(_242_v68.F19) >= 0
					_ = _rhs50
					var _lhs40 *GlobalState = _157_globalState
					_ = _lhs40
					var _lhs41 _dafny.Array = _246_v71
					_ = _lhs41
					var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_246_v71), 0))
					_ = _lhs42
					_lhs40.F8 = _rhs48
					(_lhs41).ArraySet1(_rhs49, (_lhs42).Int())
					_155_v0 = _rhs50
				} else {
					_155_v0 = false
					var _263_v84 _dafny.Array
					_ = _263_v84
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_4
					var _nw45 _dafny.Array
					_ = _nw45
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw45 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) bool = (func(_264_v0 bool) func(_dafny.Int) bool {
							return func(_265_i11 _dafny.Int) bool {
								return _264_v0
							}
						})(_155_v0)
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw45 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw45).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw45).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_263_v84 = _nw45
					var _266_v85 _dafny.Sequence
					_ = _266_v85
					_266_v85 = _dafny.SeqOf(_263_v84, _263_v84)
					var _267_v86 _dafny.MultiSet
					_ = _267_v86
					_267_v86 = _dafny.MultiSetOf(_263_v84, _263_v84, (_266_v85).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_177_v18).Cardinality()), _dafny.IntOfUint32((_266_v85).Cardinality()))).Uint32()).(_dafny.Array))
					_267_v86 = (_267_v86).Union((_267_v86).Update(_263_v84, Companion_Default___.Abs(_239_v65.F19)))
					var _268_v87 _dafny.Map
					_ = _268_v87
					_268_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _155_v0)
					var _269_v88 _dafny.Map
					_ = _269_v88
					_269_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v87, _dafny.UnicodeSeqOfUtf8Bytes("p"))
					var _270_v89 D0
					_ = _270_v89
					_270_v89 = Companion_D0_.Create_DC2_(_155_v0, _155_v0, _242_v68.F19, _155_v0)
					var _271_v90 _dafny.Map
					_ = _271_v90
					_271_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_270_v89, _270_v89), _155_v0)
					var _272_v91 _dafny.Sequence
					_ = _272_v91
					_272_v91 = _dafny.SeqOf(_269_v88)
					var _273_v92 D1
					_ = _273_v92
					_273_v92 = Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("v"))
					var _274_v94 _dafny.Sequence
					_ = _274_v94
					_274_v94 = _dafny.SeqOf(_268_v87, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, (func() bool {
						if (_268_v87).Contains(true) {
							return (_268_v87).Get(true).(bool)
						}
						return _155_v0
					})()))
					var _275_v96 _dafny.Sequence
					_ = _275_v96
					_275_v96 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qbbrxwsjm"), _177_v18)
					var _276_v97 _dafny.Array
					_ = _276_v97
					var _nwElement0_12 _dafny.Map = (_269_v88).Merge(_269_v88)
					_ = _nwElement0_12
					var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(26))
					_ = _nw46
					(_nw46).ArraySet1(_nwElement0_12, 0)
					(_nw46).ArraySet1(_269_v88, 1)
					(_nw46).ArraySet1((_269_v88).Merge(_269_v88), 2)
					(_nw46).ArraySet1(Companion_Default___.Fm6(_239_v65.F19, _158_v2, _155_v0, _271_v90, _157_globalState), 3)
					(_nw46).ArraySet1(_269_v88, 4)
					(_nw46).ArraySet1((_269_v88).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, Companion_Default___.Fm2(_155_v0, _157_globalState)), _177_v18), 5)
					(_nw46).ArraySet1(_269_v88, 6)
					(_nw46).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v87, _dafny.UnicodeSeqOfUtf8Bytes("yv")), 7)
					(_nw46).ArraySet1(_269_v88, 8)
					(_nw46).ArraySet1((_269_v88).Merge(_269_v88), 9)
					(_nw46).ArraySet1(_269_v88, 10)
					(_nw46).ArraySet1((_272_v91).Select((Companion_Default___.SafeIndex(_158_v2, _dafny.IntOfUint32((_272_v91).Cardinality()))).Uint32()).(_dafny.Map), 11)
					(_nw46).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v87, (_273_v92).Dtor_cf8())).Merge(func() _dafny.Map {
						var _coll5 = _dafny.NewMapBuilder()
						_ = _coll5
						for _iter5 := _dafny.Iterate((_dafny.MultiSetFromSeq(_274_v94)).Elements()); ; {
							_compr_5, _ok5 := _iter5()
							if !_ok5 {
								break
							}
							var _277_v93 _dafny.Map
							_277_v93 = interface{}(_compr_5).(_dafny.Map)
							if (_dafny.MultiSetFromSeq(_274_v94)).Contains(_277_v93) {
								_coll5.Add(_277_v93, _177_v18)
							}
						}
						return _coll5.ToMap()
					}()), 12)
					(_nw46).ArraySet1((_269_v88).Merge(_269_v88), 13)
					(_nw46).ArraySet1(func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate((_274_v94).Elements()); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _278_v95 _dafny.Map
							_278_v95 = interface{}(_compr_6).(_dafny.Map)
							if _dafny.Companion_Sequence_.Contains(_274_v94, _278_v95) {
								_coll6.Add(_278_v95, _177_v18)
							}
						}
						return _coll6.ToMap()
					}(), 14)
					(_nw46).ArraySet1(_269_v88, 15)
					(_nw46).ArraySet1(((_269_v88).Update(_268_v87, (_275_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_238_v64).Cardinality()), _dafny.IntOfUint32((_275_v96).Cardinality()))).Uint32()).(_dafny.Sequence))).Merge(_269_v88), 16)
					(_nw46).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v87, _177_v18), 17)
					(_nw46).ArraySet1(_269_v88, 18)
					(_nw46).ArraySet1((_272_v91).Select((Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_272_v91).Cardinality()))).Uint32()).(_dafny.Map), 19)
					(_nw46).ArraySet1(_269_v88, 20)
					(_nw46).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v87, _dafny.Companion_Sequence_.Update(_177_v18, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-222), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32(), (_177_v18).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint))), 21)
					(_nw46).ArraySet1(_269_v88, 22)
					(_nw46).ArraySet1(_269_v88, 23)
					(_nw46).ArraySet1(_269_v88, 24)
					(_nw46).ArraySet1(_269_v88, 25)
					_276_v97 = _nw46
					var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_276_v97), 0))
					_ = _index46
					(_276_v97).ArraySet1((_269_v88).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v87, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(37))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}(func(_279_i12 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('c')
					})))), (_index46).Int())
					var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_276_v97), 0))
					_ = _index47
					(_276_v97).ArraySet1(_269_v88, (_index47).Int())
					var _280_v98 _dafny.Array
					_ = _280_v98
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_5
					var _nw47 _dafny.Array
					_ = _nw47
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw47 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) _dafny.Set = (func(_281_v0 bool) func(_dafny.Int) _dafny.Set {
							return func(_282_i14 _dafny.Int) _dafny.Set {
								return _dafny.SetOf(_281_v0, _281_v0, false)
							}
						})(_155_v0)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw47 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw47).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw47).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_280_v98 = _nw47
					var _283_v99 _dafny.Map
					_ = _283_v99
					_283_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v98, _177_v18)
					var _284_v100 _dafny.Array
					_ = _284_v100
					var _nwElement0_13 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-94))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg23 _dafny.Int) interface{} {
							return coer23(arg23)
						}
					}(func(_285_i13 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('w')
					}))
					_ = _nwElement0_13
					var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(27))
					_ = _nw48
					(_nw48).ArraySet1(_nwElement0_13, 0)
					(_nw48).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tjdvjrwn"), 1)
					(_nw48).ArraySet1((func() _dafny.Sequence {
						if _155_v0 {
							return _177_v18
						}
						return _177_v18
					})(), 2)
					(_nw48).ArraySet1((func() _dafny.Sequence {
						if (_283_v99).Contains(_280_v98) {
							return (_283_v99).Get(_280_v98).(_dafny.Sequence)
						}
						return _177_v18
					})(), 3)
					(_nw48).ArraySet1(Companion_Default___.Fm5(_242_v68.F19, _157_globalState), 4)
					(_nw48).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-417))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg24 _dafny.Int) interface{} {
							return coer24(arg24)
						}
					}(func(_286_i15 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('y')
					})), 5)
					(_nw48).ArraySet1(_177_v18, 6)
					(_nw48).ArraySet1(_177_v18, 7)
					(_nw48).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vxgk"), 8)
					(_nw48).ArraySet1(_177_v18, 9)
					(_nw48).ArraySet1(_177_v18, 10)
					(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vb"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-384))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}(func(_287_i16 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('u')
					}))), 11)
					(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_177_v18, _177_v18), 12)
					(_nw48).ArraySet1(_177_v18, 13)
					(_nw48).ArraySet1(_177_v18, 14)
					(_nw48).ArraySet1(_177_v18, 15)
					(_nw48).ArraySet1(_177_v18, 16)
					(_nw48).ArraySet1(_177_v18, 17)
					(_nw48).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("a"), 18)
					(_nw48).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(340))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}(func(_288_i17 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					})), 19)
					(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_177_v18, _dafny.UnicodeSeqOfUtf8Bytes("kkieswsc")), 20)
					(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_177_v18, _177_v18), 21)
					(_nw48).ArraySet1(_177_v18, 22)
					(_nw48).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(551))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg27 _dafny.Int) interface{} {
							return coer27(arg27)
						}
					}(func(_289_i18 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					})), 23)
					(_nw48).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uscna"), 24)
					(_nw48).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_275_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.IntOfUint32((_275_v96).Cardinality()))).Uint32()).(_dafny.Sequence), _177_v18), 25)
					(_nw48).ArraySet1(_177_v18, 26)
					_284_v100 = _nw48
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_284_v100), 0))
					_ = _index48
					(_284_v100).ArraySet1((func() _dafny.Sequence {
						if _155_v0 {
							return _dafny.UnicodeSeqOfUtf8Bytes("btynhwo")
						}
						return _177_v18
					})(), (_index48).Int())
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_284_v100), 0))
					_ = _index49
					(_284_v100).ArraySet1(_177_v18, (_index49).Int())
					var _rhs51 _dafny.Int = _242_v68.F19
					_ = _rhs51
					var _rhs52 _dafny.Int = (func() _dafny.Int {
						if (_241_v67).Contains(_dafny.IntOfUint32((_238_v64).Cardinality())) {
							return (_241_v67).Get(_dafny.IntOfUint32((_238_v64).Cardinality())).(_dafny.Int)
						}
						return _242_v68.F19
					})()
					_ = _rhs52
					var _rhs53 bool = !(_155_v0)
					_ = _rhs53
					var _rhs54 _dafny.Int = _dafny.IntOfUint32(((_284_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_284_v100), 0))).Int()).(_dafny.Sequence)).Cardinality())
					_ = _rhs54
					var _rhs55 _dafny.Int = ((_242_v68.F19).Times(_242_v68.F19)).Plus((func() _dafny.Int {
						if (_159_v3).Contains(_155_v0) {
							return (_159_v3).Multiplicity(_155_v0)
						}
						return _239_v65.F19
					})())
					_ = _rhs55
					var _lhs43 *C0 = _239_v65
					_ = _lhs43
					var _lhs44 *GlobalState = _157_globalState
					_ = _lhs44
					var _lhs45 *GlobalState = _157_globalState
					_ = _lhs45
					var _lhs46 *GlobalState = _157_globalState
					_ = _lhs46
					var _lhs47 *GlobalState = _157_globalState
					_ = _lhs47
					_lhs43.F19 = _rhs51
					_lhs44.F0 = _rhs52
					_lhs45.F3 = _rhs53
					_lhs46.F8 = _rhs54
					_lhs47.F12 = _rhs55
				}
				(_157_globalState).F3 = _155_v0
				var _290_v101 D1
				_ = _290_v101
				_290_v101 = Companion_D1_.Create_DC5_((!(_155_v0)) && (_155_v0))
				var _source6 D1 = _290_v101
				_ = _source6
				if _source6.Is_DC4() {
					var _291___mcc_h0 *C0 = _source6.Get_().(D1_DC4).Cf9
					_ = _291___mcc_h0
					var _292_cf9 *C0 = _291___mcc_h0
					_ = _292_cf9
					_156_v1 = _156_v1
					_159_v3 = (_159_v3).Difference(_159_v3)
					(_239_v65).F19 = (_dafny.Zero).Minus(_239_v65.F19)
					var _293_v102 _dafny.CodePoint
					_ = _293_v102
					_293_v102 = _dafny.CodePoint('s')
					var _294_v103 _dafny.Map
					_ = _294_v103
					_294_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _238_v64)
					var _rhs56 _dafny.CodePoint = (func() _dafny.CodePoint {
						if _155_v0 {
							return _dafny.CodePoint('a')
						}
						return Companion_Default___.Fm7(_155_v0, _239_v65.F19, _157_globalState)
					})()
					_ = _rhs56
					var _rhs57 bool = _155_v0
					_ = _rhs57
					var _rhs58 _dafny.Int = (_242_v68).Fm4((_dafny.Zero).Minus(_292_cf9.F19), _158_v2, _155_v0, !(_294_v103).Equals(Companion_Default___.Fm8(_dafny.IntOfUint32((_177_v18).Cardinality()), _292_cf9.F19, _155_v0, _157_globalState)), _157_globalState)
					_ = _rhs58
					var _lhs48 *GlobalState = _157_globalState
					_ = _lhs48
					var _lhs49 *GlobalState = _157_globalState
					_ = _lhs49
					_293_v102 = _rhs56
					_lhs48.F3 = _rhs57
					_lhs49.F8 = _rhs58
				} else if _source6.Is_DC5() {
					var _295___mcc_h1 bool = _source6.Get_().(D1_DC5).Cf10
					_ = _295___mcc_h1
					var _296_cf10 bool = _295___mcc_h1
					_ = _296_cf10
					(_157_globalState).F3 = _155_v0
					_241_v67 = (_241_v67).Update(_239_v65.F19, Companion_Default___.SafeDivisionInt(_242_v68.F19, _242_v68.F19))
					var _297_v104 _dafny.Array
					_ = _297_v104
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(26)
					_ = _len0_6
					var _nw49 _dafny.Array
					_ = _nw49
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw49 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.CodePoint = func(_298_i19 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('i')
						}
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw49 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw49).ArraySet1CodePoint(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw49).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_297_v104 = _nw49
					var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_297_v104), 0))
					_ = _index50
					(_297_v104).ArraySet1CodePoint(_dafny.CodePoint('i'), (_index50).Int())
					var _299_v105 _dafny.CodePoint
					_ = _299_v105
					_299_v105 = _dafny.CodePoint('l')
					var _300_v106 _dafny.Set
					_ = _300_v106
					_300_v106 = _dafny.SetOf(_299_v105, _299_v105, _299_v105)
					var _301_v107 _dafny.MultiSet
					_ = _301_v107
					_301_v107 = _dafny.MultiSetOf(_299_v105, _299_v105)
					var _302_v108 _dafny.Sequence
					_ = _302_v108
					_302_v108 = _dafny.SeqOf(_177_v18, Companion_Default___.Fm5((func() _dafny.Int {
						if (_301_v107).Contains(_299_v105) {
							return (_301_v107).Multiplicity(_299_v105)
						}
						return _242_v68.F19
					})(), _157_globalState))
					var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_297_v104), 0))
					_ = _index51
					var _rhs59 bool = !(_296_cf10) || (!(_300_v106).Equals(_dafny.SetOf(_299_v105, _299_v105, _299_v105)))
					_ = _rhs59
					var _rhs60 _dafny.CodePoint = _299_v105
					_ = _rhs60
					var _rhs61 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_156_v1, _156_v1), _dafny.SeqOf(_155_v0))
					_ = _rhs61
					var _rhs62 bool = !_dafny.Companion_Sequence_.Equal(_302_v108, _dafny.Companion_Sequence_.Concatenate(_302_v108, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("owfqc"))))
					_ = _rhs62
					var _lhs50 *GlobalState = _157_globalState
					_ = _lhs50
					var _lhs51 _dafny.Array = _297_v104
					_ = _lhs51
					var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_297_v104), 0))
					_ = _lhs52
					var _lhs53 *GlobalState = _157_globalState
					_ = _lhs53
					_lhs50.F3 = _rhs59
					(_lhs51).ArraySet1CodePoint(_rhs60, (_lhs52).Int())
					_lhs53.F3 = _rhs61
					_296_cf10 = _rhs62
					var _303_v110 _dafny.Array
					_ = _303_v110
					var _nwElement0_14 _dafny.Int = (func() _dafny.Map {
						var _coll7 = _dafny.NewMapBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate((_238_v64).Elements()); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _304_v109 _dafny.Int
							_304_v109 = interface{}(_compr_7).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_238_v64, _304_v109) {
								_coll7.Add(Companion_Default___.SafeDivisionInt(_304_v109, (_dafny.Zero).Minus(_158_v2)), (Companion_D0_.Create_DC1_(_242_v68.F19, _dafny.IntOfInt64(24), true)).Dtor_cf1())
							}
						}
						return _coll7.ToMap()
					}()).Cardinality()
					_ = _nwElement0_14
					var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(3))
					_ = _nw50
					(_nw50).ArraySet1(_nwElement0_14, 0)
					(_nw50).ArraySet1(_242_v68.F19, 1)
					(_nw50).ArraySet1(_239_v65.F19, 2)
					_303_v110 = _nw50
					var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_303_v110), 0))
					_ = _index52
					(_303_v110).ArraySet1(_158_v2, (_index52).Int())
					var _305_v111 _dafny.MultiSet
					_ = _305_v111
					_305_v111 = _dafny.MultiSetOf(_241_v67)
					var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_303_v110), 0))
					_ = _index53
					(_303_v110).ArraySet1((_238_v64).Select((Companion_Default___.SafeIndex((_305_v111).Cardinality(), _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int), (_index53).Int())
				} else if _source6.Is_DC3() {
					var _306___mcc_h2 _dafny.Sequence = _source6.Get_().(D1_DC3).Cf8
					_ = _306___mcc_h2
					var _307_cf8 _dafny.Sequence = _306___mcc_h2
					_ = _307_cf8
					var _308_v112 D0
					_ = _308_v112
					_308_v112 = Companion_D0_.Create_DC2_(!(_155_v0), _155_v0, _242_v68.F19, _155_v0)
					var _309_v113 _dafny.Array
					_ = _309_v113
					var _nwElement0_15 D0 = _308_v112
					_ = _nwElement0_15
					var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.One)
					_ = _nw51
					(_nw51).ArraySet1(_nwElement0_15, 0)
					_309_v113 = _nw51
					var _310_v114 D1
					_ = _310_v114
					_310_v114 = Companion_D1_.Create_DC3_(_177_v18)
					var _311_v115 _dafny.Map
					_ = _311_v115
					_311_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_309_v113, _310_v114)
					_311_v115 = ((_311_v115).Merge(_311_v115)).Merge(_311_v115)
					(_157_globalState).F3 = !((_242_v68.F19).Cmp((_239_v65).Fm3(_157_globalState)) == 0)
					(_157_globalState).F3 = Companion_Default___.Fm2(_155_v0, _157_globalState)
					(_157_globalState).F8 = _239_v65.F19
				} else {
					var _312___mcc_h3 D1 = _source6.Get_().(D1_DC6).Cf11
					_ = _312___mcc_h3
					var _313_cf11 D1 = _312___mcc_h3
					_ = _313_cf11
					var _314_v116 *C0
					_ = _314_v116
					var _nw52 *C0 = New_C0_()
					_ = _nw52
					_nw52.Ctor__(_239_v65.F19)
					_314_v116 = _nw52
					var _315_v117 _dafny.Map
					_ = _315_v117
					_315_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(true, _157_globalState), _242_v68.F19)
					_315_v117 = (_315_v117).Update(true, _242_v68.F19)
					_238_v64 = _238_v64
					var _316_v118 _dafny.Array
					_ = _316_v118
					var _len0_7 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_7
					var _nw53 _dafny.Array
					_ = _nw53
					if _len0_7.Cmp(_dafny.Zero) == 0 {
						_nw53 = _dafny.NewArray(_len0_7)
					} else {
						var _init7 func(_dafny.Int) _dafny.Int = func(_317_i20 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_317_i20, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nffg")).Cardinality())))
						}
						_ = _init7
						var _element0_7 = _init7(_dafny.Zero)
						_ = _element0_7
						_nw53 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
						(_nw53).ArraySet1(_element0_7, 0)
						var _nativeLen0_7 = (_len0_7).Int()
						_ = _nativeLen0_7
						for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
							(_nw53).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
						}
					}
					_316_v118 = _nw53
					var _318_v119 _dafny.MultiSet
					_ = _318_v119
					_318_v119 = _dafny.MultiSetOf(_316_v118)
					var _319_v120 _dafny.Map
					_ = _319_v120
					_319_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_315_v117).Cardinality(), _155_v0)
					var _320_v121 D0
					_ = _320_v121
					_320_v121 = Companion_D0_.Create_DC1_((_319_v120).Cardinality(), _242_v68.F19, _155_v0)
					var _321_v122 _dafny.Int
					_ = _321_v122
					var _322_v123 _dafny.Sequence
					_ = _322_v123
					var _323_v124 _dafny.Int
					_ = _323_v124
					var _324_v125 _dafny.MultiSet
					_ = _324_v125
					var _out12 _dafny.Int
					_ = _out12
					var _out13 _dafny.Sequence
					_ = _out13
					var _out14 _dafny.Int
					_ = _out14
					var _out15 _dafny.MultiSet
					_ = _out15
					_out12, _out13, _out14, _out15 = Companion_Default___.M0((_239_v65.F19).Cmp((_318_v119).Cardinality()) >= 0, _155_v0, (_159_v3).Cardinality(), ((_dafny.Zero).Minus((_242_v68).Fm4(_dafny.IntOfUint32((_177_v18).Cardinality()), _239_v65.F19, (_320_v121).Dtor_cf3(), _155_v0, _157_globalState))).Cmp(_314_v116.F19) == 0, _157_globalState)
					_321_v122 = _out12
					_322_v123 = _out13
					_323_v124 = _out14
					_324_v125 = _out15
				}
				(_157_globalState).F0 = _239_v65.F19
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _325_v126 _dafny.Set
	_ = _325_v126
	_325_v126 = _dafny.SetOf(_155_v0, _155_v0)
	var _hi2 _dafny.Int = _239_v65.F19
	_ = _hi2
	for _326_i21 := (func() _dafny.Int {
		if !(_155_v0) {
			return (_dafny.Zero).Minus(_158_v2)
		}
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_242_v68.F19), (_325_v126).Cardinality())).Cardinality()
	})(); _326_i21.Cmp(_hi2) < 0; _326_i21 = _326_i21.Plus(_dafny.One) {
		_177_v18 = _177_v18
		var _327_v127 D1
		_ = _327_v127
		_327_v127 = Companion_D1_.Create_DC5_(_155_v0)
		var _328_v128 _dafny.Array
		_ = _328_v128
		var _nwElement0_16 bool = (_327_v127).Dtor_cf10()
		_ = _nwElement0_16
		var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(18))
		_ = _nw54
		(_nw54).ArraySet1(_nwElement0_16, 0)
		(_nw54).ArraySet1(_155_v0, 1)
		(_nw54).ArraySet1(false, 2)
		(_nw54).ArraySet1(_155_v0, 3)
		(_nw54).ArraySet1(_155_v0, 4)
		(_nw54).ArraySet1(true, 5)
		(_nw54).ArraySet1(_155_v0, 6)
		(_nw54).ArraySet1(_155_v0, 7)
		(_nw54).ArraySet1(_155_v0, 8)
		(_nw54).ArraySet1(_155_v0, 9)
		(_nw54).ArraySet1(!(_155_v0), 10)
		(_nw54).ArraySet1(Companion_Default___.Fm2(_155_v0, _157_globalState), 11)
		(_nw54).ArraySet1(true, 12)
		(_nw54).ArraySet1(_155_v0, 13)
		(_nw54).ArraySet1(_155_v0, 14)
		(_nw54).ArraySet1(!(_155_v0), 15)
		(_nw54).ArraySet1(_155_v0, 16)
		(_nw54).ArraySet1(Companion_Default___.Fm2(_155_v0, _157_globalState), 17)
		_328_v128 = _nw54
		var _329_v129 _dafny.Array
		_ = _329_v129
		var _nwElement0_17 _dafny.Array = (func() _dafny.Array {
			if _155_v0 {
				return _328_v128
			}
			return _328_v128
		})()
		_ = _nwElement0_17
		var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(2))
		_ = _nw55
		(_nw55).ArraySet1(_nwElement0_17, 0)
		(_nw55).ArraySet1(_328_v128, 1)
		_329_v129 = _nw55
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_329_v129), 0))
		_ = _index54
		(_329_v129).ArraySet1(_328_v128, (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_329_v129), 0))
		_ = _index55
		(_329_v129).ArraySet1((Companion_D2_.Create_DC7_(_328_v128)).Dtor_cf12(), (_index55).Int())
		var _330_v130 _dafny.Sequence
		_ = _330_v130
		_330_v130 = _dafny.SeqOf(_159_v3, _159_v3)
		var _331_v131 *C0
		_ = _331_v131
		var _nw56 *C0 = New_C0_()
		_ = _nw56
		_nw56.Ctor__(_dafny.IntOfUint32((_330_v130).Cardinality()))
		_331_v131 = _nw56
		var _332_v132 D2
		_ = _332_v132
		_332_v132 = Companion_D2_.Create_DC8_()
		var _source7 D2 = _332_v132
		_ = _source7
		if _source7.Is_DC8() {
			var _333_v133 _dafny.Array
			_ = _333_v133
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_8
			var _nw57 _dafny.Array
			_ = _nw57
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw57 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = func(_334_i22 _dafny.Int) _dafny.Int {
					return (_334_i22).Plus(_dafny.IntOfInt64(469))
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw57 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw57).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw57).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_333_v133 = _nw57
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_333_v133), 0))
			_ = _index56
			(_333_v133).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pyjrottw")).Cardinality()), (_index56).Int())
			var _335_v134 _dafny.Set
			_ = _335_v134
			_335_v134 = _dafny.SetOf(_dafny.IntOfUint32((Companion_Default___.Fm5((_241_v67).Cardinality(), _157_globalState)).Cardinality()))
			var _336_v135 _dafny.Map
			_ = _336_v135
			_336_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v68.F19, _155_v0)
			var _337_v136 D3
			_ = _337_v136
			_337_v136 = Companion_D3_.Create_DC11_(_336_v135)
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_333_v133), 0))
			_ = _index57
			var _rhs63 _dafny.Int = ((_335_v134).Difference(_335_v134)).Cardinality()
			_ = _rhs63
			var _rhs64 _dafny.Int = ((_238_v64).Select((Companion_Default___.SafeIndex(_331_v131.F19, _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int)).Plus(((_337_v136).Dtor_cf17()).Cardinality())
			_ = _rhs64
			var _rhs65 _dafny.Int = _239_v65.F19
			_ = _rhs65
			var _lhs54 _dafny.Array = _333_v133
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_333_v133), 0))
			_ = _lhs55
			var _lhs56 *GlobalState = _157_globalState
			_ = _lhs56
			_158_v2 = _rhs63
			(_lhs54).ArraySet1(_rhs64, (_lhs55).Int())
			_lhs56.F0 = _rhs65
			_155_v0 = _155_v0
			var _338_v137 _dafny.Map
			_ = _338_v137
			_338_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _155_v0)
			var _339_v138 _dafny.Int
			_ = _339_v138
			var _340_v139 _dafny.Sequence
			_ = _340_v139
			var _341_v140 _dafny.Int
			_ = _341_v140
			var _342_v141 _dafny.MultiSet
			_ = _342_v141
			var _out16 _dafny.Int
			_ = _out16
			var _out17 _dafny.Sequence
			_ = _out17
			var _out18 _dafny.Int
			_ = _out18
			var _out19 _dafny.MultiSet
			_ = _out19
			_out16, _out17, _out18, _out19 = Companion_Default___.M0(_155_v0, !(_338_v137).Contains(_155_v0), _326_i21, (func() bool {
				if (_338_v137).Contains(_155_v0) {
					return (_338_v137).Get(_155_v0).(bool)
				}
				return _155_v0
			})(), _157_globalState)
			_339_v138 = _out16
			_340_v139 = _out17
			_341_v140 = _out18
			_342_v141 = _out19
			_340_v139 = _177_v18
		} else if _source7.Is_DC9() {
			var _343___mcc_h4 _dafny.Int = _source7.Get_().(D2_DC9).Cf13
			_ = _343___mcc_h4
			var _344___mcc_h5 *C0 = _source7.Get_().(D2_DC9).Cf14
			_ = _344___mcc_h5
			var _345___mcc_h6 *C0 = _source7.Get_().(D2_DC9).Cf15
			_ = _345___mcc_h6
			var _346_cf15 *C0 = _345___mcc_h6
			_ = _346_cf15
			var _347_cf14 *C0 = _344___mcc_h5
			_ = _347_cf14
			var _348_cf13 _dafny.Int = _343___mcc_h4
			_ = _348_cf13
			(_157_globalState).F0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(763), _239_v65.F19)).Plus((_dafny.Zero).Minus(_331_v131.F19))
			_155_v0 = false
			var _349_v142 _dafny.CodePoint
			_ = _349_v142
			_349_v142 = _dafny.CodePoint('x')
			var _350_v143 _dafny.Map
			_ = _350_v143
			_350_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _349_v142)
			_350_v143 = (_350_v143).Update(_155_v0, _349_v142)
			var _351_v144 D1
			_ = _351_v144
			_351_v144 = Companion_D1_.Create_DC4_(_347_cf14)
			var _pat_let_tv19 = _242_v68
			_ = _pat_let_tv19
			var _pat_let_tv20 = _347_cf14
			_ = _pat_let_tv20
			var _352_v145 _dafny.Array
			_ = _352_v145
			var _nwElement0_18 D1 = Companion_D1_.Create_DC4_(_331_v131)
			_ = _nwElement0_18
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(25))
			_ = _nw58
			(_nw58).ArraySet1(_nwElement0_18, 0)
			(_nw58).ArraySet1((func() D1 {
				if _155_v0 {
					return _351_v144
				}
				return Companion_D1_.Create_DC4_(_242_v68)
			})(), 1)
			(_nw58).ArraySet1(_351_v144, 2)
			(_nw58).ArraySet1(_351_v144, 3)
			(_nw58).ArraySet1(_351_v144, 4)
			(_nw58).ArraySet1(Companion_D1_.Create_DC4_(_239_v65), 5)
			(_nw58).ArraySet1(_351_v144, 6)
			(_nw58).ArraySet1(_351_v144, 7)
			(_nw58).ArraySet1(_351_v144, 8)
			(_nw58).ArraySet1(_351_v144, 9)
			(_nw58).ArraySet1(_351_v144, 10)
			(_nw58).ArraySet1((func() D1 {
				if _155_v0 {
					return _351_v144
				}
				return func(_pat_let8_0 D1) D1 {
					return func(_353_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let9_0 *C0) D1 {
							return func(_354_dt__update_hcf9_h0 *C0) D1 {
								return Companion_D1_.Create_DC4_(_354_dt__update_hcf9_h0)
							}(_pat_let9_0)
						}(_pat_let_tv19)
					}(_pat_let8_0)
				}(_351_v144)
			})(), 11)
			(_nw58).ArraySet1(Companion_D1_.Create_DC4_(_242_v68), 12)
			(_nw58).ArraySet1(_351_v144, 13)
			(_nw58).ArraySet1(_351_v144, 14)
			(_nw58).ArraySet1(_351_v144, 15)
			(_nw58).ArraySet1(_351_v144, 16)
			(_nw58).ArraySet1(_351_v144, 17)
			(_nw58).ArraySet1(_351_v144, 18)
			(_nw58).ArraySet1(_351_v144, 19)
			(_nw58).ArraySet1(_351_v144, 20)
			(_nw58).ArraySet1(func(_pat_let10_0 D1) D1 {
				return func(_355_dt__update__tmp_h2 D1) D1 {
					return func(_pat_let11_0 *C0) D1 {
						return func(_356_dt__update_hcf9_h1 *C0) D1 {
							return Companion_D1_.Create_DC4_(_356_dt__update_hcf9_h1)
						}(_pat_let11_0)
					}(_pat_let_tv20)
				}(_pat_let10_0)
			}(_351_v144), 21)
			(_nw58).ArraySet1(Companion_D1_.Create_DC4_(_331_v131), 22)
			(_nw58).ArraySet1(Companion_D1_.Create_DC4_(_242_v68), 23)
			(_nw58).ArraySet1(_351_v144, 24)
			_352_v145 = _nw58
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_352_v145), 0))
			_ = _index58
			(_352_v145).ArraySet1(_351_v144, (_index58).Int())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_352_v145), 0))
			_ = _index59
			var _rhs66 D1 = _351_v144
			_ = _rhs66
			var _rhs67 _dafny.Int = (_dafny.Zero).Minus(_326_i21)
			_ = _rhs67
			var _rhs68 _dafny.Int = (Companion_D0_.Create_DC2_(_155_v0, _155_v0, _242_v68.F19, _155_v0)).Dtor_cf6()
			_ = _rhs68
			var _rhs69 _dafny.Int = (_239_v65.F19).Times(_242_v68.F19)
			_ = _rhs69
			var _lhs57 _dafny.Array = _352_v145
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_352_v145), 0))
			_ = _lhs58
			var _lhs59 *C0 = _239_v65
			_ = _lhs59
			var _lhs60 *C0 = _242_v68
			_ = _lhs60
			var _lhs61 *GlobalState = _157_globalState
			_ = _lhs61
			(_lhs57).ArraySet1(_rhs66, (_lhs58).Int())
			_lhs59.F19 = _rhs67
			_lhs60.F19 = _rhs68
			_lhs61.F2 = _rhs69
		} else if _source7.Is_DC7() {
			var _357___mcc_h7 _dafny.Array = _source7.Get_().(D2_DC7).Cf12
			_ = _357___mcc_h7
			var _358_cf12 _dafny.Array = _357___mcc_h7
			_ = _358_cf12
			var _359_v146 _dafny.Array
			_ = _359_v146
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
			_ = _nw59
			_359_v146 = _nw59
			var _360_v147 D2
			_ = _360_v147
			_360_v147 = Companion_D2_.Create_DC9_(_242_v68.F19, _331_v131, _239_v65)
			var _361_v148 _dafny.Sequence
			_ = _361_v148
			_361_v148 = _dafny.SeqOf(_360_v147)
			var _362_v149 _dafny.Map
			_ = _362_v149
			_362_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v68.F19, Companion_D2_.Create_DC10_((_361_v148).Select((Companion_Default___.SafeIndex(_331_v131.F19, _dafny.IntOfUint32((_361_v148).Cardinality()))).Uint32()).(D2)))
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_359_v146), 0))
			_ = _index60
			(_359_v146).ArraySet1(_362_v149, (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_359_v146), 0))
			_ = _index61
			(_359_v146).ArraySet1((_362_v149).Merge(_362_v149), (_index61).Int())
			var _363_v150 _dafny.MultiSet
			_ = _363_v150
			_363_v150 = _dafny.MultiSetOf(_239_v65, _239_v65)
			var _364_v151 _dafny.Map
			_ = _364_v151
			_364_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_363_v150, _326_i21)
			var _365_v152 _dafny.Map
			_ = _365_v152
			_365_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _155_v0)
			var _366_v153 _dafny.Sequence
			_ = _366_v153
			_366_v153 = _dafny.SeqOf(_242_v68)
			var _367_v154 D4
			_ = _367_v154
			_367_v154 = Companion_D4_.Create_DC14_(_366_v153)
			var _368_v155 _dafny.MultiSet
			_ = _368_v155
			_368_v155 = _dafny.MultiSetOf(_158_v2, _239_v65.F19)
			var _369_v156 _dafny.Array
			_ = _369_v156
			var _nwElement0_19 _dafny.Map = _364_v151
			_ = _nwElement0_19
			var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(8))
			_ = _nw60
			(_nw60).ArraySet1(_nwElement0_19, 0)
			(_nw60).ArraySet1((_364_v151).Merge(((_364_v151).Update(_363_v150, (_365_v152).Cardinality())).Update(_363_v150, (_238_v64).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int))), 1)
			(_nw60).ArraySet1(_364_v151, 2)
			(_nw60).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq((_367_v154).Dtor_cf21()), _326_i21), 3)
			(_nw60).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_363_v150, _239_v65.F19)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_363_v150, _326_i21)), 4)
			(_nw60).ArraySet1((_364_v151).Update(_363_v150, (_dafny.Zero).Minus((_368_v155).Cardinality())), 5)
			(_nw60).ArraySet1(_364_v151, 6)
			(_nw60).ArraySet1(_364_v151, 7)
			_369_v156 = _nw60
			var _370_v157 _dafny.Map
			_ = _370_v157
			_370_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _331_v131.F19)
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_369_v156), 0))
			_ = _index62
			(_369_v156).ArraySet1((_364_v151).Update(_363_v150, (_370_v157).Cardinality()), (_index62).Int())
			var _371_v158 _dafny.Array
			_ = _371_v158
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_9
			var _nw61 _dafny.Array
			_ = _nw61
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw61 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = (func(_372_v68 *C0) func(_dafny.Int) _dafny.Int {
					return func(_373_i23 _dafny.Int) _dafny.Int {
						return (_373_i23).Minus(_372_v68.F19)
					}
				})(_242_v68)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw61 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw61).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw61).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_371_v158 = _nw61
			var _374_v159 _dafny.Map
			_ = _374_v159
			_374_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_241_v67).Merge(_241_v67), _371_v158)
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_369_v156), 0))
			_ = _index63
			var _rhs70 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64)).Cardinality())
			_ = _rhs70
			var _rhs71 _dafny.Map = _364_v151
			_ = _rhs71
			var _rhs72 _dafny.Array = (func() _dafny.Array {
				if (_374_v159).Contains(func() _dafny.Map {
					var _coll8 = _dafny.NewMapBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(333), _dafny.IntOfInt64(713))); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _375_v160 _dafny.Int
						_375_v160 = interface{}(_compr_8).(_dafny.Int)
						if ((_dafny.IntOfInt64(333)).Cmp(_375_v160) <= 0) && ((_375_v160).Cmp(_dafny.IntOfInt64(713)) < 0) {
							_coll8.Add((_375_v160).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_177_v18).Cardinality()))), _dafny.IntOfInt64(719))
						}
					}
					return _coll8.ToMap()
				}()) {
					return (_374_v159).Get(func() _dafny.Map {
						var _coll9 = _dafny.NewMapBuilder()
						_ = _coll9
						for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(333), _dafny.IntOfInt64(713))); ; {
							_compr_9, _ok9 := _iter9()
							if !_ok9 {
								break
							}
							var _376_v160 _dafny.Int
							_376_v160 = interface{}(_compr_9).(_dafny.Int)
							if ((_dafny.IntOfInt64(333)).Cmp(_376_v160) <= 0) && ((_376_v160).Cmp(_dafny.IntOfInt64(713)) < 0) {
								_coll9.Add((_376_v160).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_177_v18).Cardinality()))), _dafny.IntOfInt64(719))
							}
						}
						return _coll9.ToMap()
					}()).(_dafny.Array)
				}
				return _371_v158
			})()
			_ = _rhs72
			var _lhs62 *GlobalState = _157_globalState
			_ = _lhs62
			var _lhs63 _dafny.Array = _369_v156
			_ = _lhs63
			var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_369_v156), 0))
			_ = _lhs64
			_lhs62.F12 = _rhs70
			(_lhs63).ArraySet1(_rhs71, (_lhs64).Int())
			_371_v158 = _rhs72
			(_157_globalState).F8 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(701), _239_v65.F19)
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_328_v128), 0))
			_ = _index64
			(_328_v128).ArraySet1(_155_v0, (_index64).Int())
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_328_v128), 0))
			_ = _index65
			(_328_v128).ArraySet1(_155_v0, (_index65).Int())
		} else {
			var _377___mcc_h8 D2 = _source7.Get_().(D2_DC10).Cf16
			_ = _377___mcc_h8
			var _378_cf16 D2 = _377___mcc_h8
			_ = _378_cf16
			(_157_globalState).F3 = (func() bool {
				if Companion_Default___.Fm2(_155_v0, _157_globalState) {
					return _155_v0
				}
				return (_242_v68.F19).Cmp(_331_v131.F19) == 0
			})()
			var _379_v161 _dafny.CodePoint
			_ = _379_v161
			_379_v161 = _dafny.CodePoint('e')
			var _380_v162 D5
			_ = _380_v162
			_380_v162 = Companion_D5_.Create_DC17_((_177_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_177_v18).Cardinality()), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_379_v161 = (_380_v162).Dtor_cf23()
			var _381_v163 _dafny.Map
			_ = _381_v163
			_381_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, true)
			var _382_v164 _dafny.MultiSet
			_ = _382_v164
			_382_v164 = _dafny.MultiSetOf((_381_v163).Cardinality())
			var _rhs73 *C0 = _242_v68
			_ = _rhs73
			var _rhs74 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_379_v161, _379_v161, (_177_v18).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_382_v164).Contains(_239_v65.F19) {
					return (_382_v164).Multiplicity(_239_v65.F19)
				}
				return (_159_v3).Cardinality()
			})(), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint), (func() _dafny.CodePoint {
				if _155_v0 {
					return _dafny.CodePoint('j')
				}
				return _dafny.CodePoint('t')
			})(), _379_v161), (Companion_Default___.SafeIndex((_239_v65.F19).Times(_326_i21), _dafny.IntOfUint32((_dafny.SeqOf(_379_v161, _379_v161, (_177_v18).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_382_v164).Contains(_239_v65.F19) {
					return (_382_v164).Multiplicity(_239_v65.F19)
				}
				return (_159_v3).Cardinality()
			})(), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint), (func() _dafny.CodePoint {
				if _155_v0 {
					return _dafny.CodePoint('j')
				}
				return _dafny.CodePoint('t')
			})(), _379_v161)).Cardinality()))).Uint32(), _379_v161), (Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_379_v161, _379_v161, (_177_v18).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_382_v164).Contains(_239_v65.F19) {
					return (_382_v164).Multiplicity(_239_v65.F19)
				}
				return (_159_v3).Cardinality()
			})(), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint), (func() _dafny.CodePoint {
				if _155_v0 {
					return _dafny.CodePoint('j')
				}
				return _dafny.CodePoint('t')
			})(), _379_v161), (Companion_Default___.SafeIndex((_239_v65.F19).Times(_326_i21), _dafny.IntOfUint32((_dafny.SeqOf(_379_v161, _379_v161, (_177_v18).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_382_v164).Contains(_239_v65.F19) {
					return (_382_v164).Multiplicity(_239_v65.F19)
				}
				return (_159_v3).Cardinality()
			})(), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint), (func() _dafny.CodePoint {
				if _155_v0 {
					return _dafny.CodePoint('j')
				}
				return _dafny.CodePoint('t')
			})(), _379_v161)).Cardinality()))).Uint32(), _379_v161)).Cardinality()))).Uint32(), _379_v161)
			_ = _rhs74
			_331_v131 = _rhs73
			_177_v18 = _rhs74
			var _383_v165 _dafny.Array
			_ = _383_v165
			var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw62
			_383_v165 = _nw62
			var _384_v166 _dafny.MultiSet
			_ = _384_v166
			_384_v166 = _dafny.MultiSetOf(_dafny.SeqOf(_155_v0, _155_v0), _156_v1)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_383_v165), 0))
			_ = _index66
			(_383_v165).ArraySet1((_384_v166).Cardinality(), (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_383_v165), 0))
			_ = _index67
			(_383_v165).ArraySet1(((_158_v2).Times(_331_v131.F19)).Minus(_326_i21), (_index67).Int())
		}
	}
	var _385_v167 _dafny.Array
	_ = _385_v167
	var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
	_ = _nw63
	_385_v167 = _nw63
	var _386_v168 D5
	_ = _386_v168
	_386_v168 = Companion_D5_.Create_DC18_(_242_v68.F19, _155_v0, _385_v167, _155_v0)
	var _387_v169 D5
	_ = _387_v169
	_387_v169 = Companion_D5_.Create_DC20_(_386_v168)
	var _388_v170 D5
	_ = _388_v170
	_388_v170 = Companion_D5_.Create_DC20_(_387_v169)
	var _source8 D5 = _388_v170
	_ = _source8
	if _source8.Is_DC18() {
		var _389___mcc_h9 _dafny.Int = _source8.Get_().(D5_DC18).Cf24
		_ = _389___mcc_h9
		var _390___mcc_h10 bool = _source8.Get_().(D5_DC18).Cf25
		_ = _390___mcc_h10
		var _391___mcc_h11 _dafny.Array = _source8.Get_().(D5_DC18).Cf26
		_ = _391___mcc_h11
		var _392___mcc_h12 bool = _source8.Get_().(D5_DC18).Cf27
		_ = _392___mcc_h12
		var _393_cf27 bool = _392___mcc_h12
		_ = _393_cf27
		var _394_cf26 _dafny.Array = _391___mcc_h11
		_ = _394_cf26
		var _395_cf25 bool = _390___mcc_h10
		_ = _395_cf25
		var _396_cf24 _dafny.Int = _389___mcc_h9
		_ = _396_cf24
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_394_cf26), 0))
		_ = _index68
		(_394_cf26).ArraySet1(_395_cf25, (_index68).Int())
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_394_cf26), 0))
		_ = _index69
		(_394_cf26).ArraySet1(Companion_Default___.Fm2(!_dafny.Companion_Sequence_.Equal(_156_v1, _dafny.SeqOf(_395_cf25)), _157_globalState), (_index69).Int())
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_394_cf26), 0))
		_ = _index70
		var _rhs75 _dafny.Int = _239_v65.F19
		_ = _rhs75
		var _rhs76 bool = !(!(_155_v0) || (_393_cf27))
		_ = _rhs76
		var _rhs77 _dafny.Int = ((_325_v126).Cardinality()).Times(_242_v68.F19)
		_ = _rhs77
		var _rhs78 bool = _395_cf25
		_ = _rhs78
		var _lhs65 *GlobalState = _157_globalState
		_ = _lhs65
		var _lhs66 _dafny.Array = _394_cf26
		_ = _lhs66
		var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_394_cf26), 0))
		_ = _lhs67
		var _lhs68 *GlobalState = _157_globalState
		_ = _lhs68
		_lhs65.F0 = _rhs75
		(_lhs66).ArraySet1(_rhs76, (_lhs67).Int())
		_lhs68.F12 = _rhs77
		_155_v0 = _rhs78
		var _397_v171 _dafny.MultiSet
		_ = _397_v171
		_397_v171 = _dafny.MultiSetOf((_240_v66).Merge(_240_v66))
		_397_v171 = _397_v171
		var _398_v172 D2
		_ = _398_v172
		_398_v172 = Companion_D2_.Create_DC8_()
		var _source9 D2 = _398_v172
		_ = _source9
		if _source9.Is_DC8() {
			(_157_globalState).F3 = (_241_v67).Equals((func() _dafny.Map {
				if _155_v0 {
					return func() _dafny.Map {
						var _coll10 = _dafny.NewMapBuilder()
						_ = _coll10
						for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(143), _dafny.IntOfInt64(-7))); ; {
							_compr_10, _ok10 := _iter10()
							if !_ok10 {
								break
							}
							var _399_v173 _dafny.Int
							_399_v173 = interface{}(_compr_10).(_dafny.Int)
							if ((_dafny.IntOfInt64(143)).Cmp(_399_v173) <= 0) && ((_399_v173).Cmp(_dafny.IntOfInt64(-7)) < 0) {
								_coll10.Add((_399_v173).Times(_242_v68.F19), _396_cf24)
							}
						}
						return _coll10.ToMap()
					}()
				}
				return _241_v67
			})())
			(_157_globalState).F0 = _dafny.IntOfInt64(174)
			_177_v18 = (func() _dafny.Sequence {
				if (_394_cf26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_394_cf26), 0))).Int()).(bool) {
					return _177_v18
				}
				return _177_v18
			})()
			var _400_v174 _dafny.CodePoint
			_ = _400_v174
			_400_v174 = _dafny.CodePoint('l')
			_400_v174 = _400_v174
		} else if _source9.Is_DC9() {
			var _401___mcc_h16 _dafny.Int = _source9.Get_().(D2_DC9).Cf13
			_ = _401___mcc_h16
			var _402___mcc_h17 *C0 = _source9.Get_().(D2_DC9).Cf14
			_ = _402___mcc_h17
			var _403___mcc_h18 *C0 = _source9.Get_().(D2_DC9).Cf15
			_ = _403___mcc_h18
			var _404_cf15 *C0 = _403___mcc_h18
			_ = _404_cf15
			var _405_cf14 *C0 = _402___mcc_h17
			_ = _405_cf14
			var _406_cf13 _dafny.Int = _401___mcc_h16
			_ = _406_cf13
			(_239_v65).F19 = _406_cf13
			(_239_v65).F19 = Companion_Default___.SafeModuloInt(_239_v65.F19, (_242_v68).Fm3(_157_globalState))
			var _407_v175 _dafny.Sequence
			_ = _407_v175
			_407_v175 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_156_v1, _156_v1), _156_v1, _156_v1, _dafny.SeqOf(_155_v0, _395_cf25, _395_cf25), _156_v1)
			_407_v175 = _dafny.Companion_Sequence_.Concatenate(_407_v175, _407_v175)
			var _408_v176 _dafny.Array
			_ = _408_v176
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
			_ = _nw64
			_408_v176 = _nw64
			var _409_v177 _dafny.CodePoint
			_ = _409_v177
			_409_v177 = _dafny.CodePoint('n')
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_408_v176), 0))
			_ = _index71
			(_408_v176).ArraySet1CodePoint(_409_v177, (_index71).Int())
			var _410_v178 _dafny.Map
			_ = _410_v178
			_410_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v177, _155_v0), _393_cf27)
			var _411_v179 D1
			_ = _411_v179
			_411_v179 = Companion_D1_.Create_DC5_(_395_cf25)
			var _412_v180 _dafny.Map
			_ = _412_v180
			_412_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v177, (_411_v179).Dtor_cf10())
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_408_v176), 0))
			_ = _index72
			(_408_v176).ArraySet1CodePoint(Companion_Default___.Fm7((func() bool {
				if (_410_v178).Contains(_412_v180) {
					return (_410_v178).Get(_412_v180).(bool)
				}
				return !(_393_cf27)
			})(), _dafny.IntOfInt64(318), _157_globalState), (_index72).Int())
		} else if _source9.Is_DC7() {
			var _413___mcc_h19 _dafny.Array = _source9.Get_().(D2_DC7).Cf12
			_ = _413___mcc_h19
			var _414_cf12 _dafny.Array = _413___mcc_h19
			_ = _414_cf12
			_155_v0 = (_155_v0) == (_155_v0)
			(_242_v68).F19 = ((_238_v64).Select((Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int)).Times(_dafny.IntOfUint32((_dafny.SeqOf(_242_v68.F19)).Cardinality()))
			var _415_v181 _dafny.Int
			_ = _415_v181
			var _416_v182 _dafny.Sequence
			_ = _416_v182
			var _417_v183 _dafny.Int
			_ = _417_v183
			var _418_v184 _dafny.MultiSet
			_ = _418_v184
			var _out20 _dafny.Int
			_ = _out20
			var _out21 _dafny.Sequence
			_ = _out21
			var _out22 _dafny.Int
			_ = _out22
			var _out23 _dafny.MultiSet
			_ = _out23
			_out20, _out21, _out22, _out23 = Companion_Default___.M0(_155_v0, _395_cf25, _239_v65.F19, (func() bool {
				if false {
					return (_156_v1).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool)
				}
				return _393_cf27
			})(), _157_globalState)
			_415_v181 = _out20
			_416_v182 = _out21
			_417_v183 = _out22
			_418_v184 = _out23
			var _419_v185 _dafny.CodePoint
			_ = _419_v185
			_419_v185 = _dafny.CodePoint('t')
			var _420_v186 _dafny.Map
			_ = _420_v186
			_420_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v65, _416_v182)
			var _421_v187 _dafny.Sequence
			_ = _421_v187
			_421_v187 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(340))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_422_v185 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_423_i24 _dafny.Int) _dafny.CodePoint {
					return _422_v185
				}
			})(_419_v185))), _177_v18)
			var _424_v188 _dafny.Array
			_ = _424_v188
			var _nwElement0_20 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_177_v18, _416_v182), (Companion_Default___.SafeIndex(_415_v181, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_177_v18, _416_v182)).Cardinality()))).Uint32(), _419_v185)
			_ = _nwElement0_20
			var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(21))
			_ = _nw65
			(_nw65).ArraySet1(_nwElement0_20, 0)
			(_nw65).ArraySet1(_416_v182, 1)
			(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_177_v18, _416_v182), 2)
			(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rcptp"), 3)
			(_nw65).ArraySet1((func() _dafny.Sequence {
				if (_420_v186).Contains(_242_v68) {
					return (_420_v186).Get(_242_v68).(_dafny.Sequence)
				}
				return _177_v18
			})(), 4)
			(_nw65).ArraySet1((_421_v187).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_421_v187).Cardinality()))).Uint32()).(_dafny.Sequence), 5)
			(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bbyhlyjx"), _416_v182), 6)
			(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_416_v182, _416_v182), 7)
			(_nw65).ArraySet1(_416_v182, 8)
			(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_242_v68.F19, _157_globalState), _dafny.UnicodeSeqOfUtf8Bytes("xa")), 9)
			(_nw65).ArraySet1(_177_v18, 10)
			(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tgjqjb"), 11)
			(_nw65).ArraySet1(_177_v18, 12)
			(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_177_v18, _177_v18), 13)
			(_nw65).ArraySet1(_416_v182, 14)
			(_nw65).ArraySet1(_177_v18, 15)
			(_nw65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(576))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_425_v185 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_426_i25 _dafny.Int) _dafny.CodePoint {
					return _425_v185
				}
			})(_419_v185))), 16)
			(_nw65).ArraySet1(_416_v182, 17)
			(_nw65).ArraySet1(Companion_Default___.Fm5(_239_v65.F19, _157_globalState), 18)
			(_nw65).ArraySet1(_dafny.Companion_Sequence_.Update(_177_v18, (Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32(), _419_v185), 19)
			(_nw65).ArraySet1(_177_v18, 20)
			_424_v188 = _nw65
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_424_v188), 0))
			_ = _index73
			(_424_v188).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tb"), (_index73).Int())
			var _427_v189 _dafny.Array
			_ = _427_v189
			var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw66
			_427_v189 = _nw66
			var _428_v190 _dafny.Map
			_ = _428_v190
			_428_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_cf25, _427_v189)
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_424_v188), 0))
			_ = _index74
			var _rhs79 _dafny.Sequence = _416_v182
			_ = _rhs79
			var _rhs80 _dafny.Array = _414_cf12
			_ = _rhs80
			var _rhs81 bool = (func() bool {
				if false {
					return _393_cf27
				}
				return (_394_cf26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_394_cf26), 0))).Int()).(bool)
			})()
			_ = _rhs81
			var _rhs82 _dafny.Array = (func() _dafny.Array {
				if (_428_v190).Contains(_395_cf25) {
					return (_428_v190).Get(_395_cf25).(_dafny.Array)
				}
				return _427_v189
			})()
			_ = _rhs82
			var _lhs69 _dafny.Array = _424_v188
			_ = _lhs69
			var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_424_v188), 0))
			_ = _lhs70
			(_lhs69).ArraySet1(_rhs79, (_lhs70).Int())
			_414_cf12 = _rhs80
			_155_v0 = _rhs81
			_427_v189 = _rhs82
		} else {
			var _429___mcc_h20 D2 = _source9.Get_().(D2_DC10).Cf16
			_ = _429___mcc_h20
			var _430_cf16 D2 = _429___mcc_h20
			_ = _430_cf16
			(_157_globalState).F3 = _393_cf27
			(_157_globalState).F0 = _239_v65.F19
			_395_cf25 = ((_156_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool)) && (true)
			var _431_v191 _dafny.MultiSet
			_ = _431_v191
			_431_v191 = _dafny.MultiSetOf(_156_v1, _156_v1, _156_v1)
			_431_v191 = (_431_v191).Intersection((_431_v191).Intersection(_431_v191))
		}
	} else if _source8.Is_DC19() {
		var _432___mcc_h13 bool = _source8.Get_().(D5_DC19).Cf28
		_ = _432___mcc_h13
		var _433_cf28 bool = _432___mcc_h13
		_ = _433_cf28
		var _434_v192 _dafny.CodePoint
		_ = _434_v192
		_434_v192 = _dafny.CodePoint('w')
		_434_v192 = (_177_v18).Select((Companion_Default___.SafeIndex(_158_v2, _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint)
		var _435_v193 _dafny.MultiSet
		_ = _435_v193
		_435_v193 = _dafny.MultiSetOf(_158_v2)
		if (_435_v193).IsProperSubsetOf(_435_v193) {
			_325_v126 = _325_v126
			var _436_v194 *C0
			_ = _436_v194
			var _nw67 *C0 = New_C0_()
			_ = _nw67
			_nw67.Ctor__(_242_v68.F19)
			_436_v194 = _nw67
			var _437_v195 _dafny.Set
			_ = _437_v195
			_437_v195 = _dafny.SetOf(_177_v18)
			_437_v195 = Companion_Default___.Fm9((_dafny.Zero).Minus(_242_v68.F19), _dafny.CodePoint('u'), Companion_Default___.Fm2(_155_v0, _157_globalState), _157_globalState)
			_239_v65 = _239_v65
			_433_cf28 = _155_v0
		} else {
			(_157_globalState).F3 = ((_433_cf28) || (_155_v0)) || (false)
			(_157_globalState).F3 = _433_cf28
			var _438_v196 _dafny.Int
			_ = _438_v196
			var _439_v197 _dafny.Sequence
			_ = _439_v197
			var _440_v198 _dafny.Int
			_ = _440_v198
			var _441_v199 _dafny.MultiSet
			_ = _441_v199
			var _out24 _dafny.Int
			_ = _out24
			var _out25 _dafny.Sequence
			_ = _out25
			var _out26 _dafny.Int
			_ = _out26
			var _out27 _dafny.MultiSet
			_ = _out27
			_out24, _out25, _out26, _out27 = Companion_Default___.M0(_433_cf28, !_dafny.Companion_Sequence_.Contains(_238_v64, _239_v65.F19), _239_v65.F19, (_242_v68.F19).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_242_v68.F19), _434_v192)).Cardinality()) < 0, _157_globalState)
			_438_v196 = _out24
			_439_v197 = _out25
			_440_v198 = _out26
			_441_v199 = _out27
			var _442_v200 _dafny.Array
			_ = _442_v200
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
			_ = _nw68
			_442_v200 = _nw68
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_442_v200), 0))
			_ = _index75
			(_442_v200).ArraySet1(_241_v67, (_index75).Int())
			var _443_v201 _dafny.Map
			_ = _443_v201
			_443_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, (_241_v67).Merge(_241_v67))
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_442_v200), 0))
			_ = _index76
			(_442_v200).ArraySet1((func() _dafny.Map {
				if (_443_v201).Contains(_438_v196) {
					return (_443_v201).Get(_438_v196).(_dafny.Map)
				}
				return _241_v67
			})(), (_index76).Int())
			_385_v167 = _385_v167
		}
		var _444_v202 _dafny.Map
		_ = _444_v202
		_444_v202 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _dafny.SetOf(_433_cf28))
		var _445_v203 D0
		_ = _445_v203
		_445_v203 = Companion_D0_.Create_DC1_(((func() _dafny.Set {
			if (_444_v202).Contains(true) {
				return (_444_v202).Get(true).(_dafny.Set)
			}
			return _325_v126
		})()).Cardinality(), _242_v68.F19, _155_v0)
		var _446_v204 *C0
		_ = _446_v204
		var _nw69 *C0 = New_C0_()
		_ = _nw69
		_nw69.Ctor__((_445_v203).Dtor_cf1())
		_446_v204 = _nw69
		var _447_v205 _dafny.Array
		_ = _447_v205
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
		_ = _nw70
		_447_v205 = _nw70
		var _448_v206 _dafny.MultiSet
		_ = _448_v206
		_448_v206 = _dafny.MultiSetOf(_241_v67)
		var _449_v207 _dafny.Map
		_ = _449_v207
		_449_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_447_v205, _448_v206)
		_449_v207 = (_449_v207).Update(_447_v205, (_448_v206).Update(Companion_Default___.Fm10(_157_globalState), Companion_Default___.Abs(_158_v2)))
	} else if _source8.Is_DC17() {
		var _450___mcc_h14 _dafny.CodePoint = _source8.Get_().(D5_DC17).Cf23
		_ = _450___mcc_h14
		var _451_cf23 _dafny.CodePoint = _450___mcc_h14
		_ = _451_cf23
		_158_v2 = (_242_v68.F19).Minus(_242_v68.F19)
		var _452_v208 _dafny.Map
		_ = _452_v208
		_452_v208 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, true)
		_155_v0 = (func() bool {
			if (_452_v208).Contains((_159_v3).Cardinality()) {
				return (_452_v208).Get((_159_v3).Cardinality()).(bool)
			}
			return _155_v0
		})()
		var _453_v209 _dafny.MultiSet
		_ = _453_v209
		_453_v209 = _dafny.MultiSetOf((_dafny.Zero).Minus(_242_v68.F19))
		if (Companion_D0_.Create_DC1_(_158_v2, (func() _dafny.Int {
			if (_453_v209).Contains(_dafny.IntOfUint32((_177_v18).Cardinality())) {
				return (_453_v209).Multiplicity(_dafny.IntOfUint32((_177_v18).Cardinality()))
			}
			return _239_v65.F19
		})(), _155_v0)).Dtor_cf3() {
			_238_v64 = _238_v64
			(_157_globalState).F3 = _155_v0
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_385_v167), 0))
			_ = _index77
			(_385_v167).ArraySet1(!(_155_v0), (_index77).Int())
			var _454_v210 _dafny.Array
			_ = _454_v210
			var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw71
			_454_v210 = _nw71
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_454_v210), 0))
			_ = _index78
			(_454_v210).ArraySet1(Companion_Default___.SafeModuloInt(_239_v65.F19, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_455_cf23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_456_i26 _dafny.Int) _dafny.CodePoint {
					return _455_cf23
				}
			})(_451_cf23)))).Cardinality())), (_index78).Int())
			var _457_v211 _dafny.Set
			_ = _457_v211
			_457_v211 = _dafny.SetOf(_242_v68)
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_385_v167), 0))
			_ = _index79
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_454_v210), 0))
			_ = _index80
			var _rhs83 bool = true
			_ = _rhs83
			var _rhs84 bool = !(Companion_Default___.Fm2(_155_v0, _157_globalState))
			_ = _rhs84
			var _rhs85 _dafny.Int = Companion_Default___.SafeModuloInt(_239_v65.F19, _dafny.IntOfInt64(-237))
			_ = _rhs85
			var _rhs86 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_177_v18, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_177_v18).Cardinality()), _dafny.IntOfInt64(-470))), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32(), _451_cf23), (Companion_Default___.SafeIndex((_453_v209).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_177_v18, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_177_v18).Cardinality()), _dafny.IntOfInt64(-470))), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32(), _451_cf23)).Cardinality()))).Uint32(), (_177_v18).Select((Companion_Default___.SafeIndex((_453_v209).Cardinality(), _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_ = _rhs86
			var _rhs87 bool = (_457_v211).IsDisjointFrom(_457_v211)
			_ = _rhs87
			var _lhs71 _dafny.Array = _385_v167
			_ = _lhs71
			var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_385_v167), 0))
			_ = _lhs72
			var _lhs73 *GlobalState = _157_globalState
			_ = _lhs73
			var _lhs74 _dafny.Array = _454_v210
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_454_v210), 0))
			_ = _lhs75
			(_lhs71).ArraySet1(_rhs83, (_lhs72).Int())
			_lhs73.F3 = _rhs84
			(_lhs74).ArraySet1(_rhs85, (_lhs75).Int())
			_177_v18 = _rhs86
			_155_v0 = _rhs87
			(_157_globalState).F3 = !((func() bool {
				if _155_v0 {
					return _155_v0
				}
				return (_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool)
			})())
			var _458_v212 D2
			_ = _458_v212
			_458_v212 = Companion_D2_.Create_DC9_(_239_v65.F19, _242_v68, _242_v68)
			var _459_v213 D2
			_ = _459_v213
			_459_v213 = Companion_D2_.Create_DC10_(_458_v212)
			var _pat_let_tv21 = _458_v212
			_ = _pat_let_tv21
			_459_v213 = func(_pat_let12_0 D2) D2 {
				return func(_460_dt__update__tmp_h3 D2) D2 {
					return func(_pat_let13_0 D2) D2 {
						return func(_461_dt__update_hcf16_h0 D2) D2 {
							return Companion_D2_.Create_DC10_(_461_dt__update_hcf16_h0)
						}(_pat_let13_0)
					}(_pat_let_tv21)
				}(_pat_let12_0)
			}(_459_v213)
		} else {
			var _462_v214 _dafny.Array
			_ = _462_v214
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_10
			var _nw72 _dafny.Array
			_ = _nw72
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw72 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Int = (func(_463_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_464_i27 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_464_i27, (_dafny.Zero).Minus(_463_v2))
					}
				})(_158_v2)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw72 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw72).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw72).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_462_v214 = _nw72
			_462_v214 = _462_v214
			var _465_v215 _dafny.Int
			_ = _465_v215
			var _466_v216 _dafny.Sequence
			_ = _466_v216
			var _467_v217 _dafny.Int
			_ = _467_v217
			var _468_v218 _dafny.MultiSet
			_ = _468_v218
			var _out28 _dafny.Int
			_ = _out28
			var _out29 _dafny.Sequence
			_ = _out29
			var _out30 _dafny.Int
			_ = _out30
			var _out31 _dafny.MultiSet
			_ = _out31
			_out28, _out29, _out30, _out31 = Companion_Default___.M0(false, _155_v0, ((_452_v208).Merge(_452_v208)).Cardinality(), false, _157_globalState)
			_465_v215 = _out28
			_466_v216 = _out29
			_467_v217 = _out30
			_468_v218 = _out31
			var _469_v219 _dafny.Set
			_ = _469_v219
			_469_v219 = _dafny.SetOf(_467_v217, _dafny.IntOfUint32((_466_v216).Cardinality()))
			var _470_v220 _dafny.Sequence
			_ = _470_v220
			_470_v220 = _dafny.SeqOf(Companion_Default___.Fm5(_239_v65.F19, _157_globalState), _177_v18)
			(_157_globalState).F3 = ((_469_v219).Cardinality()).Cmp(Companion_Default___.Fm0(_242_v68.F19, _451_cf23, _470_v220, _239_v65.F19, _157_globalState)) <= 0
			_241_v67 = (_241_v67).Update(_dafny.IntOfUint32((_177_v18).Cardinality()), _239_v65.F19)
			_385_v167 = _385_v167
		}
		var _471_v221 _dafny.Array
		_ = _471_v221
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_11
		var _nw73 _dafny.Array
		_ = _nw73
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw73 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = (func(_472_v65 *C0) func(_dafny.Int) _dafny.Int {
				return func(_473_i28 _dafny.Int) _dafny.Int {
					return (_473_i28).Plus(_472_v65.F19)
				}
			})(_239_v65)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw73 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw73).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw73).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_471_v221 = _nw73
		_471_v221 = _471_v221
	} else {
		var _474___mcc_h15 D5 = _source8.Get_().(D5_DC20).Cf29
		_ = _474___mcc_h15
		var _475_cf29 D5 = _474___mcc_h15
		_ = _475_cf29
		var _476_v222 _dafny.MultiSet
		_ = _476_v222
		_476_v222 = _dafny.MultiSetOf(_239_v65)
		_476_v222 = _476_v222
		(_157_globalState).F3 = _dafny.Companion_Sequence_.Equal(_177_v18, _dafny.UnicodeSeqOfUtf8Bytes("eevhh"))
		var _477_v223 D2
		_ = _477_v223
		_477_v223 = Companion_D2_.Create_DC9_(_239_v65.F19, _242_v68, _242_v68)
		var _nw74 *C0 = New_C0_()
		_ = _nw74
		_nw74.Ctor__((_477_v223).Dtor_cf13())
		_242_v68 = _nw74
		var _478_v224 _dafny.CodePoint
		_ = _478_v224
		_478_v224 = _dafny.CodePoint('i')
		var _479_v225 _dafny.Sequence
		_ = _479_v225
		_479_v225 = _dafny.SeqOf(_177_v18, _dafny.SeqOf(_478_v224, _478_v224))
		var _480_v226 _dafny.Map
		_ = _480_v226
		_480_v226 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(777))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_481_v224 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_482_i29 _dafny.Int) _dafny.CodePoint {
				return _481_v224
			}
		})(_478_v224))), _242_v68.F19)
		_478_v224 = Companion_Default___.Fm7(_155_v0, Companion_Default___.Fm0(_239_v65.F19, _478_v224, _479_v225, (func() _dafny.Int {
			if (_480_v226).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(889))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}((func(_483_v224 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_484_i30 _dafny.Int) _dafny.CodePoint {
					return _483_v224
				}
			})(_478_v224)))) {
				return (_480_v226).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(889))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_485_v224 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_486_i30 _dafny.Int) _dafny.CodePoint {
						return _485_v224
					}
				})(_478_v224)))).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_dafny.SeqOf(_155_v0, _155_v0)).Cardinality())
		})(), _157_globalState), _157_globalState)
	}
	var _487_v227 _dafny.CodePoint
	_ = _487_v227
	_487_v227 = _dafny.CodePoint('s')
	var _488_v228 _dafny.Array
	_ = _488_v228
	var _nwElement0_21 _dafny.CodePoint = _487_v227
	_ = _nwElement0_21
	var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(4))
	_ = _nw75
	(_nw75).ArraySet1CodePoint(_nwElement0_21, 0)
	(_nw75).ArraySet1CodePoint(_487_v227, 1)
	(_nw75).ArraySet1CodePoint((_177_v18).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
	(_nw75).ArraySet1CodePoint((func() _dafny.CodePoint {
		if false {
			return _487_v227
		}
		return _487_v227
	})(), 3)
	_488_v228 = _nw75
	var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))
	_ = _index81
	(_488_v228).ArraySet1CodePoint(Companion_Default___.Fm7(true, _dafny.IntOfInt64(157), _157_globalState), (_index81).Int())
	var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))
	_ = _index82
	(_488_v228).ArraySet1CodePoint(_487_v227, (_index82).Int())
	_158_v2 = _158_v2
	var _hi3 _dafny.Int = _158_v2
	_ = _hi3
	for _489_i31 := (_dafny.IntOfUint32((_238_v64).Cardinality())).Plus(_158_v2); _489_i31.Cmp(_hi3) < 0; _489_i31 = _489_i31.Plus(_dafny.One) {
		_177_v18 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-856))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}((func(_490_v228 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
			return func(_491_i32 _dafny.Int) _dafny.CodePoint {
				return (_490_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_490_v228), 0))).Int())
			}
		})(_488_v228)))
		if _155_v0 {
			var _492_v229 _dafny.Sequence
			_ = _492_v229
			_492_v229 = _dafny.SeqOf(_177_v18)
			var _493_v230 D0
			_ = _493_v230
			_493_v230 = Companion_D0_.Create_DC1_(Companion_Default___.Fm0(_242_v68.F19, _dafny.CodePoint('m'), _492_v229, _158_v2, _157_globalState), (_159_v3).Cardinality(), _155_v0)
			var _494_v231 _dafny.Map
			_ = _494_v231
			_494_v231 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC17_((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int())), _493_v230)
			var _495_v233 D5
			_ = _495_v233
			_495_v233 = Companion_D5_.Create_DC17_((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()))
			var _pat_let_tv22 = _487_v227
			_ = _pat_let_tv22
			var _496_v234 _dafny.Set
			_ = _496_v234
			_496_v234 = _dafny.SetOf(_495_v233, _495_v233, func(_pat_let14_0 D5) D5 {
				return func(_497_dt__update__tmp_h4 D5) D5 {
					return func(_pat_let15_0 _dafny.CodePoint) D5 {
						return func(_498_dt__update_hcf23_h0 _dafny.CodePoint) D5 {
							return Companion_D5_.Create_DC17_(_498_dt__update_hcf23_h0)
						}(_pat_let15_0)
					}(_pat_let_tv22)
				}(_pat_let14_0)
			}(_495_v233))
			var _499_v235 _dafny.Sequence
			_ = _499_v235
			_499_v235 = _dafny.SeqOf(_494_v231, func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_496_v234).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _500_v232 D5
					_500_v232 = interface{}(_compr_11).(D5)
					if (_496_v234).Contains(_500_v232) {
						_coll11.Add(_500_v232, _493_v230)
					}
				}
				return _coll11.ToMap()
			}(), (_494_v231).Merge(_494_v231))
			_499_v235 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_499_v235, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(204))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}((func(_501_v233 D5, _502_v230 D0) func(_dafny.Int) _dafny.Map {
				return func(_503_i33 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_501_v233, _502_v230)
				}
			})(_495_v233, _493_v230)))), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_499_v235, _499_v235), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_499_v235, _499_v235)).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_495_v233, _493_v230)))
			(_239_v65).F19 = _158_v2
			var _504_v236 D5
			_ = _504_v236
			_504_v236 = Companion_D5_.Create_DC18_(_158_v2, _155_v0, _385_v167, _155_v0)
			var _505_v237 _dafny.Map
			_ = _505_v237
			_505_v237 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_504_v236).Dtor_cf26(), (func() bool {
				if _155_v0 {
					return _155_v0
				}
				return true
			})())
			var _506_v238 _dafny.Map
			_ = _506_v238
			_506_v238 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_507_v68 *C0) func(_dafny.Int) _dafny.Int {
				return func(_508_i34 _dafny.Int) _dafny.Int {
					return _507_v68.F19
				}
			})(_242_v68))))
			var _509_v239 _dafny.MultiSet
			_ = _509_v239
			_509_v239 = _dafny.MultiSetOf((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()), _487_v227)
			var _510_v240 _dafny.Array
			_ = _510_v240
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_12
			var _nw76 _dafny.Array
			_ = _nw76
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw76 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Sequence = func(_511_i35 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("fvlqdof")
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw76 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw76).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw76).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_510_v240 = _nw76
			var _512_v241 _dafny.Map
			_ = _512_v241
			_512_v241 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_510_v240, _155_v0)
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))
			_ = _index83
			var _rhs88 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_385_v167, (_156_v1).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool))
			_ = _rhs88
			var _rhs89 _dafny.CodePoint = (func() _dafny.CodePoint {
				if _155_v0 {
					return _dafny.CodePoint('h')
				}
				return (_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int())
			})()
			_ = _rhs89
			var _rhs90 bool = _dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
				if (_506_v238).Contains(_155_v0) {
					return (_506_v238).Get(_155_v0).(_dafny.Sequence)
				}
				return _dafny.SeqOf((_dafny.Zero).Minus(_242_v68.F19))
			})(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64), (Companion_Default___.SafeIndex((_238_v64).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if (_509_v239).Contains((_177_v18).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint)) {
					return (_509_v239).Multiplicity((_177_v18).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_177_v18).Cardinality()))).Uint32()).(_dafny.CodePoint))
				}
				return (_dafny.Zero).Minus(_489_i31)
			})()), _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64)).Cardinality()))).Uint32(), (_dafny.Zero).Minus((_159_v3).Cardinality())))
			_ = _rhs90
			var _rhs91 bool = (func() bool {
				if (_512_v241).Contains(_510_v240) {
					return (_512_v241).Get(_510_v240).(bool)
				}
				return _155_v0
			})()
			_ = _rhs91
			var _lhs76 _dafny.Array = _488_v228
			_ = _lhs76
			var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))
			_ = _lhs77
			var _lhs78 *GlobalState = _157_globalState
			_ = _lhs78
			var _lhs79 *GlobalState = _157_globalState
			_ = _lhs79
			_505_v237 = _rhs88
			(_lhs76).ArraySet1CodePoint(_rhs89, (_lhs77).Int())
			_lhs78.F3 = _rhs90
			_lhs79.F3 = _rhs91
			(_157_globalState).F8 = _239_v65.F19
			var _513_v242 _dafny.MultiSet
			_ = _513_v242
			_513_v242 = _dafny.MultiSetOf(_239_v65.F19)
			var _514_v243 D0
			_ = _514_v243
			_514_v243 = Companion_D0_.Create_DC0_((_242_v68).Fm4(_dafny.IntOfUint32((_238_v64).Cardinality()), _242_v68.F19, _155_v0, _155_v0, _157_globalState))
			var _515_v244 _dafny.Array
			_ = _515_v244
			var _nwElement0_22 _dafny.Int = (_238_v64).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _nwElement0_22
			var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(8))
			_ = _nw77
			(_nw77).ArraySet1(_nwElement0_22, 0)
			(_nw77).ArraySet1((_242_v68).Fm3(_157_globalState), 1)
			(_nw77).ArraySet1(_158_v2, 2)
			(_nw77).ArraySet1((_159_v3).Cardinality(), 3)
			(_nw77).ArraySet1(_dafny.IntOfUint32((_177_v18).Cardinality()), 4)
			(_nw77).ArraySet1(_239_v65.F19, 5)
			(_nw77).ArraySet1(((_dafny.MultiSetFromSeq(_238_v64)).Difference(_513_v242)).Cardinality(), 6)
			(_nw77).ArraySet1((func() _dafny.Int {
				if _155_v0 {
					return _489_i31
				}
				return (_514_v243).Dtor_cf0()
			})(), 7)
			_515_v244 = _nw77
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_515_v244), 0))
			_ = _index84
			(_515_v244).ArraySet1(_158_v2, (_index84).Int())
			var _516_v245 _dafny.MultiSet
			_ = _516_v245
			_516_v245 = _dafny.MultiSetOf(_242_v68)
			var _517_v246 _dafny.Map
			_ = _517_v246
			_517_v246 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_516_v245, _489_i31)
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_515_v244), 0))
			_ = _index85
			(_515_v244).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_517_v246).Contains(_516_v245) {
					return (_517_v246).Get(_516_v245).(_dafny.Int)
				}
				return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_238_v64).Cardinality()), _158_v2)
			})()), (_index85).Int())
		} else {
			var _518_v247 _dafny.Array
			_ = _518_v247
			var _nwElement0_23 _dafny.Int = _239_v65.F19
			_ = _nwElement0_23
			var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(10))
			_ = _nw78
			(_nw78).ArraySet1(_nwElement0_23, 0)
			(_nw78).ArraySet1(_489_i31, 1)
			(_nw78).ArraySet1(_239_v65.F19, 2)
			(_nw78).ArraySet1(_242_v68.F19, 3)
			(_nw78).ArraySet1((_dafny.IntOfInt64(586)).Plus(_242_v68.F19), 4)
			(_nw78).ArraySet1(_239_v65.F19, 5)
			(_nw78).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_156_v1, _156_v1)).Cardinality()), 6)
			(_nw78).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_489_i31, _239_v65.F19)).Cardinality()), 7)
			(_nw78).ArraySet1(_239_v65.F19, 8)
			(_nw78).ArraySet1(Companion_Default___.SafeModuloInt(_239_v65.F19, _242_v68.F19), 9)
			_518_v247 = _nw78
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_518_v247), 0))
			_ = _index86
			(_518_v247).ArraySet1((_dafny.Zero).Minus(_239_v65.F19), (_index86).Int())
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_518_v247), 0))
			_ = _index87
			(_518_v247).ArraySet1(Companion_Default___.SafeModuloInt(_489_i31, (_dafny.Zero).Minus((_489_i31).Plus(_dafny.IntOfInt64(-138)))), (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_385_v167), 0))
			_ = _index88
			(_385_v167).ArraySet1(_155_v0, (_index88).Int())
			var _519_v248 _dafny.Sequence
			_ = _519_v248
			_519_v248 = _dafny.SeqOf(_239_v65)
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_385_v167), 0))
			_ = _index89
			(_385_v167).ArraySet1((_156_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_519_v248).Cardinality()), _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool), (_index89).Int())
			(_157_globalState).F8 = (_dafny.IntOfUint32((_dafny.SeqOf(_155_v0)).Cardinality())).Plus((_238_v64).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int))
			(_157_globalState).F3 = !((func() bool {
				if (_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool) {
					return (_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool)
				}
				return (_155_v0) && (!((_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool)))
			})())
			var _520_v249 *C0
			_ = _520_v249
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__(_242_v68.F19)
			_520_v249 = _nw79
		}
		var _521_v250 _dafny.Map
		_ = _521_v250
		_521_v250 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v68.F19, (_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()))
		var _522_v251 D5
		_ = _522_v251
		_522_v251 = Companion_D5_.Create_DC18_((_521_v250).Cardinality(), Companion_Default___.Fm2(_155_v0, _157_globalState), _385_v167, !(_155_v0))
		var _source10 D5 = _522_v251
		_ = _source10
		if _source10.Is_DC18() {
			var _523___mcc_h21 _dafny.Int = _source10.Get_().(D5_DC18).Cf24
			_ = _523___mcc_h21
			var _524___mcc_h22 bool = _source10.Get_().(D5_DC18).Cf25
			_ = _524___mcc_h22
			var _525___mcc_h23 _dafny.Array = _source10.Get_().(D5_DC18).Cf26
			_ = _525___mcc_h23
			var _526___mcc_h24 bool = _source10.Get_().(D5_DC18).Cf27
			_ = _526___mcc_h24
			var _527_cf27 bool = _526___mcc_h24
			_ = _527_cf27
			var _528_cf26 _dafny.Array = _525___mcc_h23
			_ = _528_cf26
			var _529_cf25 bool = _524___mcc_h22
			_ = _529_cf25
			var _530_cf24 _dafny.Int = _523___mcc_h21
			_ = _530_cf24
			var _531_v252 _dafny.Map
			_ = _531_v252
			_531_v252 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v2, _527_cf27)
			_531_v252 = (_531_v252).Update((_239_v65).Fm3(_157_globalState), _155_v0)
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_528_cf26), 0))
			_ = _index90
			(_528_cf26).ArraySet1(_529_cf25, (_index90).Int())
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_385_v167), 0))
			_ = _index91
			(_385_v167).ArraySet1(Companion_Default___.Fm2(true, _157_globalState), (_index91).Int())
			var _532_v253 _dafny.Array
			_ = _532_v253
			var _nwElement0_24 _dafny.Sequence = _177_v18
			_ = _nwElement0_24
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(7))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_24, 0)
			(_nw80).ArraySet1(_177_v18, 1)
			(_nw80).ArraySet1(_177_v18, 2)
			(_nw80).ArraySet1(_177_v18, 3)
			(_nw80).ArraySet1(_177_v18, 4)
			(_nw80).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rqsphj"), 5)
			(_nw80).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(579))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_533_v227 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_534_i36 _dafny.Int) _dafny.CodePoint {
					return _533_v227
				}
			})(_487_v227))), 6)
			_532_v253 = _nw80
			var _535_v254 D3
			_ = _535_v254
			_535_v254 = Companion_D3_.Create_DC12_(_532_v253, _529_cf25)
			var _pat_let_tv23 = _527_cf27
			_ = _pat_let_tv23
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_528_cf26), 0))
			_ = _index92
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_385_v167), 0))
			_ = _index93
			var _rhs92 bool = (func() bool {
				if (_531_v252).Contains((_238_v64).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int)) {
					return (_531_v252).Get((_238_v64).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int)).(bool)
				}
				return _529_cf25
			})()
			_ = _rhs92
			var _rhs93 bool = (Companion_Default___.Fm2(!(_155_v0), _157_globalState)) == (!_dafny.Companion_Sequence_.Contains(_177_v18, _487_v227))
			_ = _rhs93
			var _rhs94 _dafny.Int = (_239_v65).Fm4(_242_v68.F19, (_239_v65).Fm4(_239_v65.F19, _242_v68.F19, (func(_pat_let16_0 D3) D3 {
				return func(_536_dt__update__tmp_h5 D3) D3 {
					return func(_pat_let17_0 bool) D3 {
						return func(_537_dt__update_hcf19_h0 bool) D3 {
							return Companion_D3_.Create_DC12_((_536_dt__update__tmp_h5).Dtor_cf18(), _537_dt__update_hcf19_h0)
						}(_pat_let17_0)
					}(_pat_let_tv23)
				}(_pat_let16_0)
			}(_535_v254)).Dtor_cf19(), _155_v0, _157_globalState), _529_cf25, _155_v0, _157_globalState)
			_ = _rhs94
			var _rhs95 bool = _155_v0
			_ = _rhs95
			var _rhs96 bool = _529_cf25
			_ = _rhs96
			var _lhs80 _dafny.Array = _528_cf26
			_ = _lhs80
			var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_528_cf26), 0))
			_ = _lhs81
			var _lhs82 *C0 = _242_v68
			_ = _lhs82
			var _lhs83 _dafny.Array = _385_v167
			_ = _lhs83
			var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_385_v167), 0))
			_ = _lhs84
			(_lhs80).ArraySet1(_rhs92, (_lhs81).Int())
			_155_v0 = _rhs93
			_lhs82.F19 = _rhs94
			(_lhs83).ArraySet1(_rhs95, (_lhs84).Int())
			_529_cf25 = _rhs96
			(_157_globalState).F2 = _239_v65.F19
			_240_v66 = _240_v66
		} else if _source10.Is_DC19() {
			var _538___mcc_h25 bool = _source10.Get_().(D5_DC19).Cf28
			_ = _538___mcc_h25
			var _539_cf28 bool = _538___mcc_h25
			_ = _539_cf28
			var _540_v255 _dafny.Map
			_ = _540_v255
			_540_v255 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(!(_539_cf28), _157_globalState), _242_v68.F19)
			var _541_v256 _dafny.Array
			_ = _541_v256
			var _nwElement0_25 _dafny.Int = Companion_Default___.SafeModuloInt(_489_i31, _242_v68.F19)
			_ = _nwElement0_25
			var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(13))
			_ = _nw81
			(_nw81).ArraySet1(_nwElement0_25, 0)
			(_nw81).ArraySet1((_dafny.Zero).Minus(_489_i31), 1)
			(_nw81).ArraySet1(_dafny.IntOfInt64(777), 2)
			(_nw81).ArraySet1(_239_v65.F19, 3)
			(_nw81).ArraySet1((_489_i31).Minus(_489_i31), 4)
			(_nw81).ArraySet1(_239_v65.F19, 5)
			(_nw81).ArraySet1((func() _dafny.Int {
				if (_540_v255).Contains(_155_v0) {
					return (_540_v255).Get(_155_v0).(_dafny.Int)
				}
				return _242_v68.F19
			})(), 6)
			(_nw81).ArraySet1((_239_v65).Fm4(_242_v68.F19, (_dafny.Zero).Minus(_242_v68.F19), _155_v0, _155_v0, _157_globalState), 7)
			(_nw81).ArraySet1(_158_v2, 8)
			(_nw81).ArraySet1(_dafny.IntOfInt64(324), 9)
			(_nw81).ArraySet1(_239_v65.F19, 10)
			(_nw81).ArraySet1((_dafny.IntOfInt64(55)).Plus(_239_v65.F19), 11)
			(_nw81).ArraySet1((func() _dafny.Int {
				if _155_v0 {
					return _239_v65.F19
				}
				return _239_v65.F19
			})(), 12)
			_541_v256 = _nw81
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_541_v256), 0))
			_ = _index94
			(_541_v256).ArraySet1(((_540_v255).Merge(_540_v255)).Cardinality(), (_index94).Int())
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_541_v256), 0))
			_ = _index95
			(_541_v256).ArraySet1(_242_v68.F19, (_index95).Int())
			var _542_v257 _dafny.Map
			_ = _542_v257
			_542_v257 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v65.F19, _539_cf28)
			_542_v257 = (_542_v257).Update(_dafny.IntOfInt64(223), (func() bool {
				if _539_cf28 {
					return _539_cf28
				}
				return _155_v0
			})())
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_385_v167), 0))
			_ = _index96
			(_385_v167).ArraySet1((_242_v68.F19).Cmp(_242_v68.F19) < 0, (_index96).Int())
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_385_v167), 0))
			_ = _index97
			(_385_v167).ArraySet1((true) == (_539_cf28), (_index97).Int())
			var _543_v258 _dafny.Array
			_ = _543_v258
			var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw82
			_543_v258 = _nw82
			var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw83
			_543_v258 = _nw83
		} else if _source10.Is_DC17() {
			var _544___mcc_h26 _dafny.CodePoint = _source10.Get_().(D5_DC17).Cf23
			_ = _544___mcc_h26
			var _545_cf23 _dafny.CodePoint = _544___mcc_h26
			_ = _545_cf23
			var _546_v259 _dafny.Map
			_ = _546_v259
			_546_v259 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("j"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_242_v68.F19), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality()))).Uint32(), _dafny.CodePoint('b')), !(_155_v0))
			_546_v259 = (func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_546_v259).Keys().Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _547_v260 _dafny.Sequence
					_547_v260 = interface{}(_compr_12).(_dafny.Sequence)
					if (_546_v259).Contains(_547_v260) {
						_coll12.Add(_547_v260, _155_v0)
					}
				}
				return _coll12.ToMap()
			}()).Merge((_546_v259).Update(_dafny.UnicodeSeqOfUtf8Bytes("mnlnnbyv"), !(!(_155_v0))))
			var _548_v261 _dafny.Int
			_ = _548_v261
			var _549_v262 _dafny.Sequence
			_ = _549_v262
			var _550_v263 _dafny.Int
			_ = _550_v263
			var _551_v264 _dafny.MultiSet
			_ = _551_v264
			var _out32 _dafny.Int
			_ = _out32
			var _out33 _dafny.Sequence
			_ = _out33
			var _out34 _dafny.Int
			_ = _out34
			var _out35 _dafny.MultiSet
			_ = _out35
			_out32, _out33, _out34, _out35 = Companion_Default___.M0(_dafny.Companion_Sequence_.Contains(_238_v64, _239_v65.F19), _155_v0, _489_i31, !(_155_v0), _157_globalState)
			_548_v261 = _out32
			_549_v262 = _out33
			_550_v263 = _out34
			_551_v264 = _out35
			var _552_v265 _dafny.Map
			_ = _552_v265
			_552_v265 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_548_v261, _155_v0)
			var _553_v266 _dafny.Map
			_ = _553_v266
			_553_v266 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_552_v265).Update(_158_v2, true), (!(true)) && (_155_v0))
			_553_v266 = (_553_v266).Merge(_553_v266)
			var _554_v267 _dafny.Set
			_ = _554_v267
			_554_v267 = _dafny.SetOf(_385_v167)
			_554_v267 = ((_554_v267).Difference(_554_v267)).Union(_554_v267)
		} else {
			var _555___mcc_h27 D5 = _source10.Get_().(D5_DC20).Cf29
			_ = _555___mcc_h27
			var _556_cf29 D5 = _555___mcc_h27
			_ = _556_cf29
			var _557_v268 _dafny.Int
			_ = _557_v268
			var _558_v269 _dafny.Sequence
			_ = _558_v269
			var _559_v270 _dafny.Int
			_ = _559_v270
			var _560_v271 _dafny.MultiSet
			_ = _560_v271
			var _out36 _dafny.Int
			_ = _out36
			var _out37 _dafny.Sequence
			_ = _out37
			var _out38 _dafny.Int
			_ = _out38
			var _out39 _dafny.MultiSet
			_ = _out39
			_out36, _out37, _out38, _out39 = Companion_Default___.M0(!_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm5((_dafny.Zero).Minus(_489_i31), _157_globalState), _dafny.UnicodeSeqOfUtf8Bytes("pmxvhnl")), Companion_Default___.Fm2(_155_v0, _157_globalState), _239_v65.F19, _155_v0, _157_globalState)
			_557_v268 = _out36
			_558_v269 = _out37
			_559_v270 = _out38
			_560_v271 = _out39
			var _561_v272 D2
			_ = _561_v272
			_561_v272 = Companion_D2_.Create_DC9_(_239_v65.F19, _242_v68, _242_v68)
			var _562_v273 D2
			_ = _562_v273
			_562_v273 = Companion_D2_.Create_DC10_(_561_v272)
			var _563_v274 _dafny.Map
			_ = _563_v274
			_563_v274 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v65.F19, _155_v0)
			var _564_v275 D3
			_ = _564_v275
			_564_v275 = Companion_D3_.Create_DC11_(_563_v274)
			var _565_v276 _dafny.MultiSet
			_ = _565_v276
			_565_v276 = _dafny.MultiSetOf(_564_v275, Companion_D3_.Create_DC11_(_563_v274))
			var _rhs97 D2 = _562_v273
			_ = _rhs97
			var _rhs98 _dafny.Int = (func() _dafny.Int {
				if (_155_v0) || (_155_v0) {
					return _dafny.IntOfInt64(676)
				}
				return _559_v270
			})()
			_ = _rhs98
			var _rhs99 _dafny.MultiSet = (_565_v276).Update(_564_v275, Companion_Default___.Abs((_dafny.Zero).Minus(_239_v65.F19)))
			_ = _rhs99
			var _lhs85 *C0 = _239_v65
			_ = _lhs85
			_562_v273 = _rhs97
			_lhs85.F19 = _rhs98
			_565_v276 = _rhs99
			_239_v65 = (func() *C0 {
				if _155_v0 {
					return (func() *C0 {
						if _155_v0 {
							return _239_v65
						}
						return _239_v65
					})()
				}
				return _239_v65
			})()
			var _566_v277 _dafny.Array
			_ = _566_v277
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_13
			var _nw84 _dafny.Array
			_ = _nw84
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw84 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Int = func(_567_i37 _dafny.Int) _dafny.Int {
					return (_567_i37).Times(_dafny.IntOfInt64(936))
				}
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw84 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw84).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw84).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_566_v277 = _nw84
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_566_v277), 0))
			_ = _index98
			(_566_v277).ArraySet1(_158_v2, (_index98).Int())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_566_v277), 0))
			_ = _index99
			(_566_v277).ArraySet1(_242_v68.F19, (_index99).Int())
			var _568_v278 _dafny.Map
			_ = _568_v278
			_568_v278 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_558_v269, _155_v0)
			var _569_v279 _dafny.Set
			_ = _569_v279
			_569_v279 = _dafny.SetOf((_238_v64).Select((Companion_Default___.SafeIndex((_568_v278).Cardinality(), _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.Int))
			var _570_v280 _dafny.Map
			_ = _570_v280
			_570_v280 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, (_158_v2).Minus(_557_v268))
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_566_v277), 0))
			_ = _index100
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_566_v277), 0))
			_ = _index101
			var _rhs100 _dafny.Int = _158_v2
			_ = _rhs100
			var _rhs101 _dafny.Int = (func() _dafny.Int {
				if (_570_v280).Contains(true) {
					return (_570_v280).Get(true).(_dafny.Int)
				}
				return _242_v68.F19
			})()
			_ = _rhs101
			var _rhs102 bool = _155_v0
			_ = _rhs102
			var _rhs103 _dafny.Int = _489_i31
			_ = _rhs103
			var _rhs104 _dafny.Set = _569_v279
			_ = _rhs104
			var _lhs86 _dafny.Array = _566_v277
			_ = _lhs86
			var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_566_v277), 0))
			_ = _lhs87
			var _lhs88 *GlobalState = _157_globalState
			_ = _lhs88
			var _lhs89 *GlobalState = _157_globalState
			_ = _lhs89
			var _lhs90 _dafny.Array = _566_v277
			_ = _lhs90
			var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_566_v277), 0))
			_ = _lhs91
			(_lhs86).ArraySet1(_rhs100, (_lhs87).Int())
			_lhs88.F8 = _rhs101
			_lhs89.F3 = _rhs102
			(_lhs90).ArraySet1(_rhs103, (_lhs91).Int())
			_569_v279 = _rhs104
		}
		(_157_globalState).F8 = (_242_v68.F19).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}((func(_571_v227 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_572_i38 _dafny.Int) _dafny.CodePoint {
				return _571_v227
			}
		})(_487_v227)))).Cardinality()))
	}
	var _573_v281 D0
	_ = _573_v281
	_573_v281 = Companion_D0_.Create_DC1_((_dafny.MultiSetOf(_155_v0)).Cardinality(), _158_v2, !(_155_v0))
	var _source11 D0 = (func() D0 {
		if !(_155_v0) {
			return (func() D0 {
				if _155_v0 {
					return _573_v281
				}
				return _573_v281
			})()
		}
		return _573_v281
	})()
	_ = _source11
	if _source11.Is_DC1() {
		var _574___mcc_h28 _dafny.Int = _source11.Get_().(D0_DC1).Cf1
		_ = _574___mcc_h28
		var _575___mcc_h29 _dafny.Int = _source11.Get_().(D0_DC1).Cf2
		_ = _575___mcc_h29
		var _576___mcc_h30 bool = _source11.Get_().(D0_DC1).Cf3
		_ = _576___mcc_h30
		var _577_cf3 bool = _576___mcc_h30
		_ = _577_cf3
		var _578_cf2 _dafny.Int = _575___mcc_h29
		_ = _578_cf2
		var _579_cf1 _dafny.Int = _574___mcc_h28
		_ = _579_cf1
		var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_385_v167), 0))
		_ = _index102
		(_385_v167).ArraySet1(_577_cf3, (_index102).Int())
		var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_385_v167), 0))
		_ = _index103
		(_385_v167).ArraySet1(_577_cf3, (_index103).Int())
		if _577_cf3 {
			var _580_v282 _dafny.Array
			_ = _580_v282
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
			_ = _nw85
			_580_v282 = _nw85
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_580_v282), 0))
			_ = _index104
			(_580_v282).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()), _155_v0), (_index104).Int())
			var _581_v283 _dafny.Sequence
			_ = _581_v283
			_581_v283 = _dafny.SeqOf(_242_v68, _242_v68)
			var _582_v284 _dafny.Map
			_ = _582_v284
			_582_v284 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()), ((_dafny.Zero).Minus(_579_cf1)).Cmp(_dafny.IntOfUint32((_581_v283).Cardinality())) < 0)
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_580_v282), 0))
			_ = _index105
			(_580_v282).ArraySet1(_582_v284, (_index105).Int())
			(_242_v68).F19 = Companion_Default___.SafeDivisionInt(_239_v65.F19, _239_v65.F19)
			(_157_globalState).F3 = _155_v0
			var _583_v285 _dafny.Array
			_ = _583_v285
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_14
			var _nw86 _dafny.Array
			_ = _nw86
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw86 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = (func(_584_v68 *C0, _585_v0 bool, _586_v167 _dafny.Array) func(_dafny.Int) bool {
					return func(_587_i39 _dafny.Int) bool {
						return (_584_v68.F19).Cmp((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_585_v0, (_586_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_586_v167), 0))).Int()).(bool))).Cardinality())) > 0
					}
				})(_242_v68, _155_v0, _385_v167)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw86 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw86).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw86).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_583_v285 = _nw86
			var _nwElement0_26 bool = (_dafny.IntOfInt64(-48)).Cmp(_242_v68.F19) == 0
			_ = _nwElement0_26
			var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(4))
			_ = _nw87
			(_nw87).ArraySet1(_nwElement0_26, 0)
			(_nw87).ArraySet1(_155_v0, 1)
			(_nw87).ArraySet1((_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool), 2)
			(_nw87).ArraySet1((_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool), 3)
			_583_v285 = _nw87
			var _588_v286 _dafny.Array
			_ = _588_v286
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
			_ = _nw88
			_588_v286 = _nw88
			var _589_v287 D5
			_ = _589_v287
			_589_v287 = Companion_D5_.Create_DC17_(_dafny.CodePoint('v'))
			var _590_v288 _dafny.Map
			_ = _590_v288
			_590_v288 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v65.F19, (_589_v287).Dtor_cf23())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_588_v286), 0))
			_ = _index106
			(_588_v286).ArraySet1(_590_v288, (_index106).Int())
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_588_v286), 0))
			_ = _index107
			var _rhs105 _dafny.Map = (_590_v288).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(893), _487_v227))
			_ = _rhs105
			var _rhs106 *C0 = _242_v68
			_ = _rhs106
			var _lhs92 _dafny.Array = _588_v286
			_ = _lhs92
			var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_588_v286), 0))
			_ = _lhs93
			(_lhs92).ArraySet1(_rhs105, (_lhs93).Int())
			_242_v68 = _rhs106
		} else {
			_158_v2 = _158_v2
			(_157_globalState).F3 = _155_v0
			var _591_v289 _dafny.MultiSet
			_ = _591_v289
			_591_v289 = _dafny.MultiSetOf(_242_v68)
			var _592_v290 D1
			_ = _592_v290
			_592_v290 = Companion_D1_.Create_DC4_(_242_v68)
			var _593_v291 _dafny.Sequence
			_ = _593_v291
			_593_v291 = _dafny.SeqOf((_592_v290).Dtor_cf9())
			var _594_v292 _dafny.Array
			_ = _594_v292
			var _nwElement0_27 _dafny.MultiSet = _591_v289
			_ = _nwElement0_27
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(14))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_27, 0)
			(_nw89).ArraySet1((_591_v289).Difference(_dafny.MultiSetOf(_242_v68, _239_v65, _239_v65, _242_v68)), 1)
			(_nw89).ArraySet1((_591_v289).Union(_dafny.MultiSetFromSeq(_593_v291)), 2)
			(_nw89).ArraySet1(_591_v289, 3)
			(_nw89).ArraySet1(_591_v289, 4)
			(_nw89).ArraySet1(_591_v289, 5)
			(_nw89).ArraySet1((_dafny.MultiSetOf(_239_v65)).Intersection(_591_v289), 6)
			(_nw89).ArraySet1((_591_v289).Difference(_591_v289), 7)
			(_nw89).ArraySet1((_591_v289).Intersection((_dafny.MultiSetOf(_242_v68)).Update(_242_v68, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ihibjhg")).Cardinality())))), 8)
			(_nw89).ArraySet1(_591_v289, 9)
			(_nw89).ArraySet1((_591_v289).Difference(_591_v289), 10)
			(_nw89).ArraySet1((_591_v289).Union(_591_v289), 11)
			(_nw89).ArraySet1(_591_v289, 12)
			(_nw89).ArraySet1(_591_v289, 13)
			_594_v292 = _nw89
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_594_v292), 0))
			_ = _index108
			(_594_v292).ArraySet1(_591_v289, (_index108).Int())
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_594_v292), 0))
			_ = _index109
			(_594_v292).ArraySet1((_591_v289).Update(_239_v65, Companion_Default___.Abs(_578_cf2)), (_index109).Int())
			var _595_v293 _dafny.Array
			_ = _595_v293
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_15
			var _nw90 _dafny.Array
			_ = _nw90
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw90 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Map = (func(_596_v281 D0, _597_v167 _dafny.Array) func(_dafny.Int) _dafny.Map {
					return func(_598_i40 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_596_v281, (_597_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_597_v167), 0))).Int()).(bool))
					}
				})(_573_v281, _385_v167)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw90 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw90).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw90).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_595_v293 = _nw90
			var _599_v294 _dafny.Map
			_ = _599_v294
			_599_v294 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_573_v281, _155_v0)
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_595_v293), 0))
			_ = _index110
			(_595_v293).ArraySet1(_599_v294, (_index110).Int())
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_595_v293), 0))
			_ = _index111
			(_595_v293).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_573_v281, _155_v0)).Merge(_599_v294), (_index111).Int())
			var _600_v295 _dafny.MultiSet
			_ = _600_v295
			_600_v295 = _dafny.MultiSetOf(_dafny.IntOfUint32((_156_v1).Cardinality()), _242_v68.F19)
			var _601_v296 _dafny.Map
			_ = _601_v296
			_601_v296 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC20_(_386_v168), (_600_v295).Difference(_dafny.MultiSetOf(_579_cf1)))
			var _602_v297 _dafny.Sequence
			_ = _602_v297
			_602_v297 = _dafny.SeqOf(_386_v168, _387_v169)
			_601_v296 = (_601_v296).Update(Companion_D5_.Create_DC20_((_602_v297).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-900), _dafny.IntOfUint32((_602_v297).Cardinality()))).Uint32()).(D5)), _600_v295)
		}
		var _603_v298 _dafny.MultiSet
		_ = _603_v298
		_603_v298 = _dafny.MultiSetOf(_239_v65.F19)
		(_157_globalState).F8 = (_dafny.Zero).Minus((_242_v68).Fm4(Companion_Default___.SafeModuloInt((_603_v298).Cardinality(), _239_v65.F19), _dafny.IntOfInt64(-138), _155_v0, _577_cf3, _157_globalState))
		(_157_globalState).F3 = (_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool)
	} else if _source11.Is_DC2() {
		var _604___mcc_h31 bool = _source11.Get_().(D0_DC2).Cf4
		_ = _604___mcc_h31
		var _605___mcc_h32 bool = _source11.Get_().(D0_DC2).Cf5
		_ = _605___mcc_h32
		var _606___mcc_h33 _dafny.Int = _source11.Get_().(D0_DC2).Cf6
		_ = _606___mcc_h33
		var _607___mcc_h34 bool = _source11.Get_().(D0_DC2).Cf7
		_ = _607___mcc_h34
		var _608_cf7 bool = _607___mcc_h34
		_ = _608_cf7
		var _609_cf6 _dafny.Int = _606___mcc_h33
		_ = _609_cf6
		var _610_cf5 bool = _605___mcc_h32
		_ = _610_cf5
		var _611_cf4 bool = _604___mcc_h31
		_ = _611_cf4
		var _612_v299 _dafny.MultiSet
		_ = _612_v299
		_612_v299 = _dafny.MultiSetOf(_241_v67)
		var _613_v300 D0
		_ = _613_v300
		_613_v300 = Companion_D0_.Create_DC0_(_239_v65.F19)
		var _614_v301 _dafny.Array
		_ = _614_v301
		var _nwElement0_28 *C0 = _242_v68
		_ = _nwElement0_28
		var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(17))
		_ = _nw91
		(_nw91).ArraySet1(_nwElement0_28, 0)
		(_nw91).ArraySet1(_242_v68, 1)
		(_nw91).ArraySet1(_242_v68, 2)
		(_nw91).ArraySet1(_239_v65, 3)
		(_nw91).ArraySet1(_242_v68, 4)
		(_nw91).ArraySet1(_242_v68, 5)
		(_nw91).ArraySet1(_239_v65, 6)
		(_nw91).ArraySet1(_239_v65, 7)
		(_nw91).ArraySet1(_242_v68, 8)
		(_nw91).ArraySet1(_242_v68, 9)
		(_nw91).ArraySet1(_239_v65, 10)
		(_nw91).ArraySet1(_242_v68, 11)
		(_nw91).ArraySet1(_239_v65, 12)
		(_nw91).ArraySet1(_239_v65, 13)
		(_nw91).ArraySet1(_242_v68, 14)
		(_nw91).ArraySet1(_239_v65, 15)
		(_nw91).ArraySet1(_242_v68, 16)
		_614_v301 = _nw91
		var _615_v302 _dafny.Map
		_ = _615_v302
		_615_v302 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_v301, _239_v65.F19)
		var _616_v303 _dafny.Map
		_ = _616_v303
		_616_v303 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_cf4, (_241_v67).Cardinality())
		var _rhs107 bool = true
		_ = _rhs107
		var _rhs108 bool = (_242_v68.F19).Cmp(((_612_v299).Update(_241_v67, Companion_Default___.Abs(_242_v68.F19))).Cardinality()) != 0
		_ = _rhs108
		var _rhs109 bool = (_239_v65.F19).Cmp((_613_v300).Dtor_cf0()) < 0
		_ = _rhs109
		var _rhs110 _dafny.Int = Companion_Default___.SafeDivisionInt((_239_v65.F19).Minus((_615_v302).Cardinality()), (_616_v303).Cardinality())
		_ = _rhs110
		_610_cf5 = _rhs107
		_610_cf5 = _rhs108
		_155_v0 = _rhs109
		_609_cf6 = _rhs110
		if false {
			var _617_v304 _dafny.Sequence
			_ = _617_v304
			_617_v304 = _dafny.SeqOf(_dafny.SeqOf((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int())))
			(_242_v68).F19 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_158_v2, (_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()), _617_v304, _239_v65.F19, _157_globalState), _158_v2)
			var _618_v305 _dafny.Array
			_ = _618_v305
			var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
			_ = _nw92
			_618_v305 = _nw92
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_618_v305), 0))
			_ = _index112
			(_618_v305).ArraySet1(_385_v167, (_index112).Int())
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_618_v305), 0))
			_ = _index113
			(_618_v305).ArraySet1(_385_v167, (_index113).Int())
			var _619_v306 _dafny.Array
			_ = _619_v306
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
			_ = _nw93
			_619_v306 = _nw93
			var _620_v307 _dafny.MultiSet
			_ = _620_v307
			_620_v307 = _dafny.MultiSetOf(_242_v68, _239_v65, _239_v65, _239_v65, _242_v68)
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_619_v306), 0))
			_ = _index114
			(_619_v306).ArraySet1((_620_v307).Difference(_dafny.MultiSetOf(_242_v68, _242_v68)), (_index114).Int())
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_619_v306), 0))
			_ = _index115
			(_619_v306).ArraySet1((_620_v307).Union(_620_v307), (_index115).Int())
			var _621_v308 _dafny.Int
			_ = _621_v308
			var _622_v309 _dafny.Sequence
			_ = _622_v309
			var _623_v310 _dafny.Int
			_ = _623_v310
			var _624_v311 _dafny.MultiSet
			_ = _624_v311
			var _out40 _dafny.Int
			_ = _out40
			var _out41 _dafny.Sequence
			_ = _out41
			var _out42 _dafny.Int
			_ = _out42
			var _out43 _dafny.MultiSet
			_ = _out43
			_out40, _out41, _out42, _out43 = Companion_Default___.M0(_610_cf5, !((func() bool {
				if _611_cf4 {
					return _608_cf7
				}
				return _608_cf7
			})()), _dafny.IntOfInt64(602), (_159_v3).IsSubsetOf(_159_v3), _157_globalState)
			_621_v308 = _out40
			_622_v309 = _out41
			_623_v310 = _out42
			_624_v311 = _out43
			_177_v18 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_622_v309, _dafny.UnicodeSeqOfUtf8Bytes("mhymh")), (Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_622_v309, _dafny.UnicodeSeqOfUtf8Bytes("mhymh"))).Cardinality()))).Uint32(), _487_v227), Companion_Default___.Fm5(_158_v2, _157_globalState))
		} else {
			_608_cf7 = !(_608_cf7)
			var _625_v312 _dafny.Map
			_ = _625_v312
			_625_v312 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v18, _610_cf5)
			_625_v312 = (_625_v312).Update(_177_v18, _611_cf4)
			var _626_v313 _dafny.Map
			_ = _626_v313
			_626_v313 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v1, _610_cf5)
			(_157_globalState).F0 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _239_v65.F19)).Update((func() bool {
				if (_626_v313).Contains(_156_v1) {
					return (_626_v313).Get(_156_v1).(bool)
				}
				return _611_cf4
			})(), _239_v65.F19)).Cardinality(), _609_cf6), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hryki")).Cardinality()))
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_614_v301), 0))
			_ = _index116
			(_614_v301).ArraySet1(_239_v65, (_index116).Int())
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_614_v301), 0))
			_ = _index117
			(_614_v301).ArraySet1(_242_v68, (_index117).Int())
			var _627_v314 _dafny.Int
			_ = _627_v314
			var _628_v315 _dafny.Sequence
			_ = _628_v315
			var _629_v316 _dafny.Int
			_ = _629_v316
			var _630_v317 _dafny.MultiSet
			_ = _630_v317
			var _out44 _dafny.Int
			_ = _out44
			var _out45 _dafny.Sequence
			_ = _out45
			var _out46 _dafny.Int
			_ = _out46
			var _out47 _dafny.MultiSet
			_ = _out47
			_out44, _out45, _out46, _out47 = Companion_Default___.M0((_242_v68.F19).Cmp(_239_v65.F19) <= 0, (_156_v1).Select((Companion_Default___.SafeIndex(_609_cf6, _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool), _239_v65.F19, false, _157_globalState)
			_627_v314 = _out44
			_628_v315 = _out45
			_629_v316 = _out46
			_630_v317 = _out47
		}
		if _610_cf5 {
			var _631_v318 _dafny.Array
			_ = _631_v318
			var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw94
			_631_v318 = _nw94
			var _632_v319 _dafny.Map
			_ = _632_v319
			_632_v319 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_cf4, _608_cf7)
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_631_v318), 0))
			_ = _index118
			(_631_v318).ArraySet1(_632_v319, (_index118).Int())
			var _633_v320 _dafny.Sequence
			_ = _633_v320
			_633_v320 = _dafny.SeqOf(_385_v167, _385_v167)
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_631_v318), 0))
			_ = _index119
			var _rhs111 _dafny.Map = (_632_v319).Merge((_632_v319).Merge(_632_v319))
			_ = _rhs111
			var _rhs112 _dafny.Int = (Companion_Default___.SafeModuloInt(_242_v68.F19, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_156_v1, (Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32(), _155_v0)).Cardinality()))).Plus((func() _dafny.Int {
				if _608_cf7 {
					return _239_v65.F19
				}
				return _242_v68.F19
			})())
			_ = _rhs112
			var _rhs113 _dafny.Int = (func() _dafny.Int {
				if Companion_Default___.Fm2((_573_v281).Dtor_cf3(), _157_globalState) {
					return _dafny.IntOfUint32((_177_v18).Cardinality())
				}
				return (_dafny.Zero).Minus(_242_v68.F19)
			})()
			_ = _rhs113
			var _rhs114 bool = _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_385_v167), (func() _dafny.Sequence {
				if _608_cf7 {
					return _633_v320
				}
				return _633_v320
			})())
			_ = _rhs114
			var _lhs94 _dafny.Array = _631_v318
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_631_v318), 0))
			_ = _lhs95
			var _lhs96 *GlobalState = _157_globalState
			_ = _lhs96
			var _lhs97 *GlobalState = _157_globalState
			_ = _lhs97
			(_lhs94).ArraySet1(_rhs111, (_lhs95).Int())
			_lhs96.F2 = _rhs112
			_lhs97.F0 = _rhs113
			_155_v0 = _rhs114
			var _634_v321 _dafny.Set
			_ = _634_v321
			_634_v321 = _dafny.SetOf(_614_v301)
			_634_v321 = (_634_v321).Union(_634_v321)
			var _635_v322 _dafny.Int
			_ = _635_v322
			var _636_v323 _dafny.Sequence
			_ = _636_v323
			var _637_v324 _dafny.Int
			_ = _637_v324
			var _638_v325 _dafny.MultiSet
			_ = _638_v325
			var _out48 _dafny.Int
			_ = _out48
			var _out49 _dafny.Sequence
			_ = _out49
			var _out50 _dafny.Int
			_ = _out50
			var _out51 _dafny.MultiSet
			_ = _out51
			_out48, _out49, _out50, _out51 = Companion_Default___.M0(_611_cf4, _610_cf5, _239_v65.F19, !(_608_cf7), _157_globalState)
			_635_v322 = _out48
			_636_v323 = _out49
			_637_v324 = _out50
			_638_v325 = _out51
			var _639_v326 D0
			_ = _639_v326
			_639_v326 = Companion_D0_.Create_DC2_(_611_cf4, _608_cf7, _239_v65.F19, _610_cf5)
			var _640_v327 _dafny.Int
			_ = _640_v327
			var _641_v328 _dafny.Sequence
			_ = _641_v328
			var _642_v329 _dafny.Int
			_ = _642_v329
			var _643_v330 _dafny.MultiSet
			_ = _643_v330
			var _out52 _dafny.Int
			_ = _out52
			var _out53 _dafny.Sequence
			_ = _out53
			var _out54 _dafny.Int
			_ = _out54
			var _out55 _dafny.MultiSet
			_ = _out55
			_out52, _out53, _out54, _out55 = Companion_Default___.M0((_639_v326).Dtor_cf5(), Companion_Default___.Fm2(true, _157_globalState), _609_cf6, _610_cf5, _157_globalState)
			_640_v327 = _out52
			_641_v328 = _out53
			_642_v329 = _out54
			_643_v330 = _out55
			_642_v329 = _637_v324
		} else {
			var _644_v331 _dafny.Array
			_ = _644_v331
			var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
			_ = _nw95
			_644_v331 = _nw95
			var _645_v332 _dafny.Map
			_ = _645_v332
			_645_v332 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v18, _611_cf4)
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_644_v331), 0))
			_ = _index120
			(_644_v331).ArraySet1(_645_v332, (_index120).Int())
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_644_v331), 0))
			_ = _index121
			(_644_v331).ArraySet1(_645_v332, (_index121).Int())
			_487_v227 = _487_v227
			var _646_v333 D2
			_ = _646_v333
			_646_v333 = Companion_D2_.Create_DC9_(_242_v68.F19, _242_v68, _239_v65)
			var _647_v334 D1
			_ = _647_v334
			_647_v334 = Companion_D1_.Create_DC3_(Companion_Default___.Fm5(_239_v65.F19, _157_globalState))
			var _648_v335 _dafny.Set
			_ = _648_v335
			_648_v335 = _dafny.SetOf(_647_v334)
			var _649_v336 _dafny.Map
			_ = _649_v336
			_649_v336 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let18_0 D2) D2 {
				return func(_650_dt__update__tmp_h6 D2) D2 {
					return func(_pat_let19_0 _dafny.Int) D2 {
						return func(_651_dt__update_hcf13_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC9_(_651_dt__update_hcf13_h0, (_650_dt__update__tmp_h6).Dtor_cf14(), (_650_dt__update__tmp_h6).Dtor_cf15())
						}(_pat_let19_0)
					}(_dafny.IntOfInt64(961))
				}(_pat_let18_0)
			}(_646_v333), (_648_v335).Difference(_dafny.SetOf(_647_v334, Companion_D1_.Create_DC3_(_177_v18), _647_v334)))
			_649_v336 = (_649_v336).Update(Companion_D2_.Create_DC9_(_609_cf6, _239_v65, _239_v65), (_648_v335).Difference(_648_v335))
			_616_v303 = (_616_v303).Update(_611_cf4, _242_v68.F19)
			var _652_v337 _dafny.Array
			_ = _652_v337
			var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw96
			_652_v337 = _nw96
			var _653_v338 _dafny.Sequence
			_ = _653_v338
			_653_v338 = _dafny.SeqOf(_177_v18, _177_v18)
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_652_v337), 0))
			_ = _index122
			(_652_v337).ArraySet1(_dafny.IntOfUint32((_653_v338).Cardinality()), (_index122).Int())
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_652_v337), 0))
			_ = _index123
			(_652_v337).ArraySet1(_239_v65.F19, (_index123).Int())
		}
		var _654_v339 _dafny.Sequence
		_ = _654_v339
		_654_v339 = _dafny.SeqOf(_242_v68, _239_v65, _242_v68, _242_v68, _239_v65)
		var _655_v340 D4
		_ = _655_v340
		_655_v340 = Companion_D4_.Create_DC14_(_654_v339)
		var _656_v341 _dafny.Set
		_ = _656_v341
		_656_v341 = _dafny.SetOf(_238_v64, _238_v64)
		var _rhs115 bool = !(_656_v341).Equals(func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqOf(_238_v64)).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _657_v342 _dafny.Sequence
				_657_v342 = interface{}(_compr_13).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_238_v64), _657_v342) {
					_coll13.Add(_657_v342)
				}
			}
			return _coll13.ToSet()
		}())
		_ = _rhs115
		var _rhs116 D4 = _655_v340
		_ = _rhs116
		var _rhs117 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(633))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}((func(_658_v227 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_659_i41 _dafny.Int) _dafny.CodePoint {
				return _658_v227
			}
		})(_487_v227)))
		_ = _rhs117
		var _rhs118 bool = (_242_v68.F19).Cmp(_609_cf6) >= 0
		_ = _rhs118
		var _lhs98 *GlobalState = _157_globalState
		_ = _lhs98
		var _lhs99 *GlobalState = _157_globalState
		_ = _lhs99
		_lhs98.F3 = _rhs115
		_655_v340 = _rhs116
		_177_v18 = _rhs117
		_lhs99.F3 = _rhs118
	} else {
		var _660___mcc_h35 _dafny.Int = _source11.Get_().(D0_DC0).Cf0
		_ = _660___mcc_h35
		var _661_cf0 _dafny.Int = _660___mcc_h35
		_ = _661_cf0
		var _662_v343 *C0
		_ = _662_v343
		var _nw97 *C0 = New_C0_()
		_ = _nw97
		_nw97.Ctor__(_661_cf0)
		_662_v343 = _nw97
		var _663_v344 _dafny.Array
		_ = _663_v344
		var _nw98 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw98
		_663_v344 = _nw98
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_663_v344), 0))
		_ = _index124
		(_663_v344).ArraySet1(_242_v68, (_index124).Int())
		var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_663_v344), 0))
		_ = _index125
		(_663_v344).ArraySet1(_239_v65, (_index125).Int())
		var _664_v345 _dafny.MultiSet
		_ = _664_v345
		_664_v345 = _dafny.MultiSetOf(_dafny.IntOfInt64(-949))
		if ((_662_v343.F19).Cmp(_242_v68.F19) == 0) == ((_664_v345).IsProperSubsetOf(_664_v345)) {
			var _665_v346 _dafny.Set
			_ = _665_v346
			_665_v346 = _dafny.SetOf(_662_v343.F19, _662_v343.F19, _662_v343.F19, _239_v65.F19)
			var _666_v349 _dafny.MultiSet
			_ = _666_v349
			_666_v349 = _dafny.MultiSetOf(_665_v346, (_665_v346).Union(_665_v346), func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(964), _dafny.IntOfInt64(-496))); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _667_v347 _dafny.Int
					_667_v347 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(964)).Cmp(_667_v347) <= 0) && ((_667_v347).Cmp(_dafny.IntOfInt64(-496)) < 0) {
						_coll14.Add((_667_v347).Times(_662_v343.F19))
					}
				}
				return _coll14.ToSet()
			}(), Companion_Default___.Fm11(_242_v68.F19, _157_globalState), func() _dafny.Set {
				var _coll15 = _dafny.NewBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-993), _dafny.IntOfInt64(-987))); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _668_v348 _dafny.Int
					_668_v348 = interface{}(_compr_15).(_dafny.Int)
					if ((_dafny.IntOfInt64(-993)).Cmp(_668_v348) <= 0) && ((_668_v348).Cmp(_dafny.IntOfInt64(-987)) < 0) {
						_coll15.Add((_668_v348).Times(_dafny.IntOfUint32((_177_v18).Cardinality())))
					}
				}
				return _coll15.ToSet()
			}())
			var _669_v350 _dafny.Sequence
			_ = _669_v350
			_669_v350 = _dafny.SeqOf(_666_v349)
			_666_v349 = (_666_v349).Difference(((_669_v350).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_669_v350).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference(_666_v349))
			var _670_v351 _dafny.Array
			_ = _670_v351
			var _nwElement0_29 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("fnbj")
			_ = _nwElement0_29
			var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(11))
			_ = _nw99
			(_nw99).ArraySet1(_nwElement0_29, 0)
			(_nw99).ArraySet1(_177_v18, 1)
			(_nw99).ArraySet1(_177_v18, 2)
			(_nw99).ArraySet1(_177_v18, 3)
			(_nw99).ArraySet1(_177_v18, 4)
			(_nw99).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("omisqr"), 5)
			(_nw99).ArraySet1(_177_v18, 6)
			(_nw99).ArraySet1(_177_v18, 7)
			(_nw99).ArraySet1(_177_v18, 8)
			(_nw99).ArraySet1(_177_v18, 9)
			(_nw99).ArraySet1(_177_v18, 10)
			_670_v351 = _nw99
			var _671_v352 D3
			_ = _671_v352
			_671_v352 = Companion_D3_.Create_DC12_(_670_v351, _155_v0)
			var _672_v353 D1
			_ = _672_v353
			_672_v353 = Companion_D1_.Create_DC5_(false)
			var _673_v354 _dafny.Int
			_ = _673_v354
			var _674_v355 _dafny.Sequence
			_ = _674_v355
			var _675_v356 _dafny.Int
			_ = _675_v356
			var _676_v357 _dafny.MultiSet
			_ = _676_v357
			var _out56 _dafny.Int
			_ = _out56
			var _out57 _dafny.Sequence
			_ = _out57
			var _out58 _dafny.Int
			_ = _out58
			var _out59 _dafny.MultiSet
			_ = _out59
			_out56, _out57, _out58, _out59 = Companion_Default___.M0((_156_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool), _155_v0, (_239_v65).Fm4(_662_v343.F19, _239_v65.F19, (_671_v352).Dtor_cf19(), (_672_v353).Dtor_cf10(), _157_globalState), (_662_v343.F19).Cmp((_242_v68).Fm4(_158_v2, (_dafny.Zero).Minus((_dafny.Zero).Minus(_242_v68.F19)), _155_v0, _155_v0, _157_globalState)) < 0, _157_globalState)
			_673_v354 = _out56
			_674_v355 = _out57
			_675_v356 = _out58
			_676_v357 = _out59
			var _677_v358 *C0
			_ = _677_v358
			var _nw100 *C0 = New_C0_()
			_ = _nw100
			_nw100.Ctor__(_242_v68.F19)
			_677_v358 = _nw100
			var _678_v359 *C0
			_ = _678_v359
			var _nw101 *C0 = New_C0_()
			_ = _nw101
			_nw101.Ctor__(_239_v65.F19)
			_678_v359 = _nw101
			var _679_v360 _dafny.Array
			_ = _679_v360
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw102
			_679_v360 = _nw102
			_679_v360 = _679_v360
		} else {
			(_157_globalState).F3 = _155_v0
			_662_v343 = _239_v65
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_385_v167), 0))
			_ = _index126
			(_385_v167).ArraySet1(_155_v0, (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_385_v167), 0))
			_ = _index127
			(_385_v167).ArraySet1((_156_v1).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if _155_v0 {
					return _242_v68.F19
				}
				return _242_v68.F19
			})(), _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool), (_index127).Int())
			_664_v345 = _664_v345
			var _680_v361 _dafny.Sequence
			_ = _680_v361
			_680_v361 = _dafny.SeqOf(_177_v18, _dafny.UnicodeSeqOfUtf8Bytes("c"), _177_v18)
			var _681_v362 D1
			_ = _681_v362
			_681_v362 = Companion_D1_.Create_DC3_((_680_v361).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_239_v65.F19), _dafny.IntOfUint32((_680_v361).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _rhs119 D1 = _681_v362
			_ = _rhs119
			var _rhs120 _dafny.Sequence = Companion_Default___.Fm5(_158_v2, _157_globalState)
			_ = _rhs120
			_681_v362 = _rhs119
			_177_v18 = _rhs120
		}
		(_157_globalState).F3 = !(!(!(_155_v0)) || (_155_v0))
	}
	var _682_v363 _dafny.Array
	_ = _682_v363
	var _nwElement0_30 D0 = _573_v281
	_ = _nwElement0_30
	var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(9))
	_ = _nw103
	(_nw103).ArraySet1(_nwElement0_30, 0)
	(_nw103).ArraySet1(Companion_Default___.Fm12(_239_v65.F19, _242_v68.F19, _242_v68.F19, _155_v0, _157_globalState), 1)
	(_nw103).ArraySet1(_573_v281, 2)
	(_nw103).ArraySet1(Companion_D0_.Create_DC1_(_242_v68.F19, _158_v2, _155_v0), 3)
	(_nw103).ArraySet1(Companion_D0_.Create_DC1_((_239_v65).Fm3(_157_globalState), _239_v65.F19, _155_v0), 4)
	(_nw103).ArraySet1(_573_v281, 5)
	(_nw103).ArraySet1(_573_v281, 6)
	(_nw103).ArraySet1(_573_v281, 7)
	(_nw103).ArraySet1(_573_v281, 8)
	_682_v363 = _nw103
	var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_682_v363), 0))
	_ = _index128
	(_682_v363).ArraySet1(_573_v281, (_index128).Int())
	var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_682_v363), 0))
	_ = _index129
	(_682_v363).ArraySet1(_573_v281, (_index129).Int())
	var _683_v365 _dafny.Map
	_ = _683_v365
	_683_v365 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_385_v167, func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-50), _dafny.IntOfInt64(413))); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _684_v364 _dafny.Int
			_684_v364 = interface{}(_compr_16).(_dafny.Int)
			if ((_dafny.IntOfInt64(-50)).Cmp(_684_v364) <= 0) && ((_684_v364).Cmp(_dafny.IntOfInt64(413)) < 0) {
				_coll16.Add(Companion_Default___.SafeDivisionInt(_684_v364, _dafny.IntOfUint32((_156_v1).Cardinality())), _242_v68.F19)
			}
		}
		return _coll16.ToMap()
	}())
	_155_v0 = (_683_v365).Contains(_385_v167)
	var _685_v366 _dafny.Map
	_ = _685_v366
	_685_v366 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_177_v18).Cardinality()), _159_v3)
	var _686_i42 _dafny.Int
	_ = _686_i42
	_686_i42 = _dafny.Zero
	{
		for (_685_v366).Contains(_242_v68.F19) {
			{
				if (_686_i42).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_686_i42 = (_686_i42).Plus(_dafny.One)
				var _687_v367 _dafny.Map
				_ = _687_v367
				_687_v367 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v68.F19, _155_v0)
				var _688_v368 D3
				_ = _688_v368
				_688_v368 = Companion_D3_.Create_DC11_(_687_v367)
				var _689_v369 D3
				_ = _689_v369
				_689_v369 = Companion_D3_.Create_DC13_(_688_v368)
				var _690_v370 D3
				_ = _690_v370
				_690_v370 = Companion_D3_.Create_DC13_((_689_v369).Dtor_cf20())
				var _source12 D3 = _690_v370
				_ = _source12
				if _source12.Is_DC12() {
					var _691___mcc_h36 _dafny.Array = _source12.Get_().(D3_DC12).Cf18
					_ = _691___mcc_h36
					var _692___mcc_h37 bool = _source12.Get_().(D3_DC12).Cf19
					_ = _692___mcc_h37
					var _693_cf19 bool = _692___mcc_h37
					_ = _693_cf19
					var _694_cf18 _dafny.Array = _691___mcc_h36
					_ = _694_cf18
					var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_385_v167), 0))
					_ = _index130
					(_385_v167).ArraySet1((_693_cf19) || (_693_cf19), (_index130).Int())
					var _695_v371 _dafny.Set
					_ = _695_v371
					_695_v371 = _dafny.SetOf((_dafny.Zero).Minus(_242_v68.F19))
					var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_385_v167), 0))
					_ = _index131
					(_385_v167).ArraySet1((func() bool {
						if (_242_v68.F19).Cmp((_695_v371).Cardinality()) >= 0 {
							return true
						}
						return _155_v0
					})(), (_index131).Int())
					(_157_globalState).F8 = _dafny.IntOfUint32((_177_v18).Cardinality())
					var _696_v372 _dafny.Map
					_ = _696_v372
					_696_v372 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _388_v170)
					(_157_globalState).F2 = ((_696_v372).Update((_242_v68.F19).Cmp(_158_v2) != 0, _388_v170)).Cardinality()
					var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_385_v167), 0))
					_ = _index132
					(_385_v167).ArraySet1((_242_v68.F19).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_177_v18).Cardinality()), _239_v65.F19)) < 0, (_index132).Int())
				} else if _source12.Is_DC11() {
					var _697___mcc_h38 _dafny.Map = _source12.Get_().(D3_DC11).Cf17
					_ = _697___mcc_h38
					var _698_cf17 _dafny.Map = _697___mcc_h38
					_ = _698_cf17
					_158_v2 = _dafny.IntOfInt64(554)
					_158_v2 = _239_v65.F19
					var _699_v373 D5
					_ = _699_v373
					_699_v373 = Companion_D5_.Create_DC19_(_155_v0)
					var _pat_let_tv24 = _155_v0
					_ = _pat_let_tv24
					var _700_v374 _dafny.Sequence
					_ = _700_v374
					_700_v374 = _dafny.SeqOf(_699_v373, func(_pat_let20_0 D5) D5 {
						return func(_701_dt__update__tmp_h7 D5) D5 {
							return func(_pat_let21_0 bool) D5 {
								return func(_702_dt__update_hcf28_h0 bool) D5 {
									return Companion_D5_.Create_DC19_(_702_dt__update_hcf28_h0)
								}(_pat_let21_0)
							}(_pat_let_tv24)
						}(_pat_let20_0)
					}(_699_v373))
					var _703_v375 _dafny.Array
					_ = _703_v375
					var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
					_ = _nw104
					_703_v375 = _nw104
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_703_v375), 0))
					_ = _index133
					(_703_v375).ArraySet1(_239_v65.F19, (_index133).Int())
					var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_703_v375), 0))
					_ = _index134
					var _rhs121 _dafny.Sequence = _dafny.SeqOf(_699_v373, Companion_Default___.Fm13((_dafny.Zero).Minus(_242_v68.F19), _157_globalState), _699_v373)
					_ = _rhs121
					var _rhs122 _dafny.Int = (_242_v68.F19).Times(_158_v2)
					_ = _rhs122
					var _rhs123 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64), _dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64))
					_ = _rhs123
					var _rhs124 bool = _155_v0
					_ = _rhs124
					var _lhs100 _dafny.Array = _703_v375
					_ = _lhs100
					var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_703_v375), 0))
					_ = _lhs101
					_700_v374 = _rhs121
					(_lhs100).ArraySet1(_rhs122, (_lhs101).Int())
					_238_v64 = _rhs123
					_155_v0 = _rhs124
					var _704_v376 _dafny.MultiSet
					_ = _704_v376
					_704_v376 = _dafny.MultiSetOf(_242_v68.F19, _239_v65.F19, _242_v68.F19, _242_v68.F19)
					(_157_globalState).F3 = (_704_v376).Contains((_242_v68.F19).Times((_239_v65).Fm4(_239_v65.F19, _239_v65.F19, _155_v0, _155_v0, _157_globalState)))
				} else {
					var _705___mcc_h39 D3 = _source12.Get_().(D3_DC13).Cf20
					_ = _705___mcc_h39
					var _706_cf20 D3 = _705___mcc_h39
					_ = _706_cf20
					var _707_v377 _dafny.Int
					_ = _707_v377
					var _708_v378 _dafny.Sequence
					_ = _708_v378
					var _709_v379 _dafny.Int
					_ = _709_v379
					var _710_v380 _dafny.MultiSet
					_ = _710_v380
					var _out60 _dafny.Int
					_ = _out60
					var _out61 _dafny.Sequence
					_ = _out61
					var _out62 _dafny.Int
					_ = _out62
					var _out63 _dafny.MultiSet
					_ = _out63
					_out60, _out61, _out62, _out63 = Companion_Default___.M0((_155_v0) == (_155_v0), _155_v0, _242_v68.F19, _155_v0, _157_globalState)
					_707_v377 = _out60
					_708_v378 = _out61
					_709_v379 = _out62
					_710_v380 = _out63
					_155_v0 = (_155_v0) == (_155_v0)
					var _711_v381 _dafny.Set
					_ = _711_v381
					_711_v381 = _dafny.SetOf(_709_v379, _dafny.IntOfInt64(-70))
					var _712_v382 _dafny.Array
					_ = _712_v382
					var _nwElement0_31 _dafny.Sequence = _238_v64
					_ = _nwElement0_31
					var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(26))
					_ = _nw105
					(_nw105).ArraySet1(_nwElement0_31, 0)
					(_nw105).ArraySet1((func() _dafny.Sequence {
						if _155_v0 {
							return _238_v64
						}
						return _238_v64
					})(), 1)
					(_nw105).ArraySet1(_238_v64, 2)
					(_nw105).ArraySet1(_238_v64, 3)
					(_nw105).ArraySet1(_238_v64, 4)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64), 5)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64), 6)
					(_nw105).ArraySet1(_238_v64, 7)
					(_nw105).ArraySet1(_238_v64, 8)
					(_nw105).ArraySet1(_238_v64, 9)
					(_nw105).ArraySet1(_238_v64, 10)
					(_nw105).ArraySet1(_238_v64, 11)
					(_nw105).ArraySet1(_238_v64, 12)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}((func(_713_v68 *C0) func(_dafny.Int) _dafny.Int {
						return func(_714_i43 _dafny.Int) _dafny.Int {
							return _713_v68.F19
						}
					})(_242_v68))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_242_v68).Fm4(_242_v68.F19, _dafny.IntOfInt64(248), _155_v0, true, _157_globalState)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg41 _dafny.Int) interface{} {
							return coer41(arg41)
						}
					}((func(_715_v68 *C0) func(_dafny.Int) _dafny.Int {
						return func(_716_i43 _dafny.Int) _dafny.Int {
							return _715_v68.F19
						}
					})(_242_v68)))).Cardinality()))).Uint32(), _158_v2), 13)
					(_nw105).ArraySet1(_238_v64, 14)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64), 15)
					(_nw105).ArraySet1(_238_v64, 16)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(809))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}((func(_717_v68 *C0) func(_dafny.Int) _dafny.Int {
						return func(_718_i44 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus(_717_v68.F19)
						}
					})(_242_v68))), (Companion_Default___.SafeIndex(_158_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(809))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg43 _dafny.Int) interface{} {
							return coer43(arg43)
						}
					}((func(_719_v68 *C0) func(_dafny.Int) _dafny.Int {
						return func(_720_i44 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus(_719_v68.F19)
						}
					})(_242_v68)))).Cardinality()))).Uint32(), _239_v65.F19), 17)
					(_nw105).ArraySet1(_238_v64, 18)
					(_nw105).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(302))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg44 _dafny.Int) interface{} {
							return coer44(arg44)
						}
					}((func(_721_v68 *C0) func(_dafny.Int) _dafny.Int {
						return func(_722_i45 _dafny.Int) _dafny.Int {
							return _721_v68.F19
						}
					})(_242_v68))), 19)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_242_v68.F19, (_711_v381).Cardinality()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.IntOfUint32((_dafny.SeqOf(_242_v68.F19, (_711_v381).Cardinality())).Cardinality()))).Uint32(), _709_v379), 20)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_238_v64, _dafny.SeqOf(_707_v377)), 21)
					(_nw105).ArraySet1(_238_v64, 22)
					(_nw105).ArraySet1(_238_v64, 23)
					(_nw105).ArraySet1(Companion_Default___.Fm14(_155_v0, _157_globalState), 24)
					(_nw105).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(750), _242_v68.F19), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_155_v0)).Cardinality()), _239_v65.F19)), 25)
					_712_v382 = _nw105
					var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_712_v382), 0))
					_ = _index135
					(_712_v382).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_238_v64, _238_v64), (_index135).Int())
					var _723_v383 _dafny.Array
					_ = _723_v383
					var _len0_16 _dafny.Int = _dafny.IntOfInt64(26)
					_ = _len0_16
					var _nw106 _dafny.Array
					_ = _nw106
					if _len0_16.Cmp(_dafny.Zero) == 0 {
						_nw106 = _dafny.NewArray(_len0_16)
					} else {
						var _init16 func(_dafny.Int) _dafny.Int = (func(_724_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_725_i46 _dafny.Int) _dafny.Int {
								return (_725_i46).Times(_724_v2)
							}
						})(_158_v2)
						_ = _init16
						var _element0_16 = _init16(_dafny.Zero)
						_ = _element0_16
						_nw106 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
						(_nw106).ArraySet1(_element0_16, 0)
						var _nativeLen0_16 = (_len0_16).Int()
						_ = _nativeLen0_16
						for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
							(_nw106).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
						}
					}
					_723_v383 = _nw106
					var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_723_v383), 0))
					_ = _index136
					(_723_v383).ArraySet1(_239_v65.F19, (_index136).Int())
					var _726_v384 _dafny.Map
					_ = _726_v384
					_726_v384 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _dafny.SeqOf(_239_v65.F19))
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_712_v382), 0))
					_ = _index137
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_723_v383), 0))
					_ = _index138
					var _rhs125 bool = _dafny.Companion_Sequence_.Contains(_156_v1, _155_v0)
					_ = _rhs125
					var _rhs126 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_238_v64, _dafny.SeqOf(_158_v2)), _238_v64)
					_ = _rhs126
					var _rhs127 _dafny.Set = _325_v126
					_ = _rhs127
					var _rhs128 _dafny.Int = (_242_v68).Fm4((_726_v384).Cardinality(), _239_v65.F19, ((_239_v65).Fm3(_157_globalState)).Cmp(_242_v68.F19) == 0, _155_v0, _157_globalState)
					_ = _rhs128
					var _lhs102 *GlobalState = _157_globalState
					_ = _lhs102
					var _lhs103 _dafny.Array = _712_v382
					_ = _lhs103
					var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_712_v382), 0))
					_ = _lhs104
					var _lhs105 _dafny.Array = _723_v383
					_ = _lhs105
					var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_723_v383), 0))
					_ = _lhs106
					_lhs102.F3 = _rhs125
					(_lhs103).ArraySet1(_rhs126, (_lhs104).Int())
					_325_v126 = _rhs127
					(_lhs105).ArraySet1(_rhs128, (_lhs106).Int())
					(_157_globalState).F3 = Companion_Default___.Fm2(!(false), _157_globalState)
				}
				if _155_v0 {
					var _rhs129 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-882))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}((func(_727_v227 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_728_i47 _dafny.Int) _dafny.CodePoint {
							return _727_v227
						}
					})(_487_v227))), (Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-882))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}((func(_729_v227 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_730_i47 _dafny.Int) _dafny.CodePoint {
							return _729_v227
						}
					})(_487_v227)))).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
						if (_156_v1).Select((Companion_Default___.SafeIndex(_242_v68.F19, _dafny.IntOfUint32((_156_v1).Cardinality()))).Uint32()).(bool) {
							return (_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int())
						}
						return _487_v227
					})())
					_ = _rhs129
					var _rhs130 _dafny.CodePoint = _487_v227
					_ = _rhs130
					var _rhs131 bool = ((func() _dafny.Int {
						if (_241_v67).Contains(_239_v65.F19) {
							return (_241_v67).Get(_239_v65.F19).(_dafny.Int)
						}
						return _239_v65.F19
					})()).Cmp(_242_v68.F19) <= 0
					_ = _rhs131
					var _rhs132 _dafny.Int = _242_v68.F19
					_ = _rhs132
					var _rhs133 *C0 = _239_v65
					_ = _rhs133
					var _lhs107 *GlobalState = _157_globalState
					_ = _lhs107
					_177_v18 = _rhs129
					_487_v227 = _rhs130
					_155_v0 = _rhs131
					_lhs107.F2 = _rhs132
					_242_v68 = _rhs133
					var _nw107 *C0 = New_C0_()
					_ = _nw107
					_nw107.Ctor__(_239_v65.F19)
					_239_v65 = _nw107
					var _731_v385 _dafny.Map
					_ = _731_v385
					_731_v385 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _487_v227)
					_731_v385 = (_731_v385).Update(_155_v0, _dafny.CodePoint('m'))
					var _732_v386 _dafny.Array
					_ = _732_v386
					var _len0_17 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_17
					var _nw108 _dafny.Array
					_ = _nw108
					if _len0_17.Cmp(_dafny.Zero) == 0 {
						_nw108 = _dafny.NewArray(_len0_17)
					} else {
						var _init17 func(_dafny.Int) _dafny.Int = (func(_733_v65 *C0) func(_dafny.Int) _dafny.Int {
							return func(_734_i48 _dafny.Int) _dafny.Int {
								return (_734_i48).Minus(_733_v65.F19)
							}
						})(_239_v65)
						_ = _init17
						var _element0_17 = _init17(_dafny.Zero)
						_ = _element0_17
						_nw108 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
						(_nw108).ArraySet1(_element0_17, 0)
						var _nativeLen0_17 = (_len0_17).Int()
						_ = _nativeLen0_17
						for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
							(_nw108).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
						}
					}
					_732_v386 = _nw108
					var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_732_v386), 0))
					_ = _index139
					(_732_v386).ArraySet1(_242_v68.F19, (_index139).Int())
					var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_732_v386), 0))
					_ = _index140
					var _rhs134 _dafny.Sequence = _238_v64
					_ = _rhs134
					var _rhs135 *C0 = _239_v65
					_ = _rhs135
					var _rhs136 _dafny.Int = _242_v68.F19
					_ = _rhs136
					var _rhs137 _dafny.Int = _dafny.IntOfInt64(742)
					_ = _rhs137
					var _rhs138 _dafny.CodePoint = (func() _dafny.CodePoint {
						if (_731_v385).Contains(_155_v0) {
							return (_731_v385).Get(_155_v0).(_dafny.CodePoint)
						}
						return _dafny.CodePoint('e')
					})()
					_ = _rhs138
					var _lhs108 *GlobalState = _157_globalState
					_ = _lhs108
					var _lhs109 _dafny.Array = _732_v386
					_ = _lhs109
					var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_732_v386), 0))
					_ = _lhs110
					_238_v64 = _rhs134
					_242_v68 = _rhs135
					_lhs108.F8 = _rhs136
					(_lhs109).ArraySet1(_rhs137, (_lhs110).Int())
					_487_v227 = _rhs138
					(_239_v65).F19 = _239_v65.F19
				} else {
					var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_385_v167), 0))
					_ = _index141
					(_385_v167).ArraySet1(_dafny.Companion_Sequence_.Contains(_177_v18, _487_v227), (_index141).Int())
					var _735_v387 _dafny.Set
					_ = _735_v387
					_735_v387 = _dafny.SetOf((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()), _dafny.CodePoint('p'))
					var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_385_v167), 0))
					_ = _index142
					(_385_v167).ArraySet1((_155_v0) == ((_735_v387).IsSubsetOf(_735_v387)), (_index142).Int())
					var _736_v388 *C0
					_ = _736_v388
					var _nw109 *C0 = New_C0_()
					_ = _nw109
					_nw109.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_177_v18).Cardinality()), _239_v65.F19))
					_736_v388 = _nw109
					var _737_v389 _dafny.Set
					_ = _737_v389
					_737_v389 = _dafny.SetOf(_dafny.IntOfInt64(512), _dafny.IntOfInt64(-517), _242_v68.F19)
					var _rhs139 _dafny.Set = (func() _dafny.Set {
						if !(_155_v0) {
							return Companion_Default___.Fm11(_239_v65.F19, _157_globalState)
						}
						return func() _dafny.Set {
							var _coll17 = _dafny.NewBuilder()
							_ = _coll17
							for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(779), _dafny.IntOfInt64(398))); ; {
								_compr_17, _ok17 := _iter17()
								if !_ok17 {
									break
								}
								var _738_v390 _dafny.Int
								_738_v390 = interface{}(_compr_17).(_dafny.Int)
								if ((_dafny.IntOfInt64(779)).Cmp(_738_v390) <= 0) && ((_738_v390).Cmp(_dafny.IntOfInt64(398)) < 0) {
									_coll17.Add(Companion_Default___.SafeModuloInt(_738_v390, _239_v65.F19))
								}
							}
							return _coll17.ToSet()
						}()
					})()
					_ = _rhs139
					var _rhs140 _dafny.Int = (_dafny.Zero).Minus((_239_v65.F19).Minus(_dafny.IntOfInt64(168)))
					_ = _rhs140
					var _rhs141 _dafny.Int = ((_dafny.Zero).Minus(_242_v68.F19)).Minus((_239_v65).Fm3(_157_globalState))
					_ = _rhs141
					var _rhs142 bool = _155_v0
					_ = _rhs142
					var _lhs111 *C0 = _242_v68
					_ = _lhs111
					var _lhs112 *GlobalState = _157_globalState
					_ = _lhs112
					var _lhs113 *GlobalState = _157_globalState
					_ = _lhs113
					_737_v389 = _rhs139
					_lhs111.F19 = _rhs140
					_lhs112.F12 = _rhs141
					_lhs113.F3 = _rhs142
					var _739_v391 *C0
					_ = _739_v391
					var _nw110 *C0 = New_C0_()
					_ = _nw110
					_nw110.Ctor__((func() _dafny.Int {
						if (_385_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_385_v167), 0))).Int()).(bool) {
							return _dafny.IntOfInt64(264)
						}
						return _242_v68.F19
					})())
					_739_v391 = _nw110
					(_736_v388).F19 = _158_v2
				}
				var _740_v392 D5
				_ = _740_v392
				_740_v392 = Companion_D5_.Create_DC17_(_487_v227)
				var _741_v393 _dafny.Sequence
				_ = _741_v393
				_741_v393 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(458))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_742_v227 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_743_i49 _dafny.Int) _dafny.CodePoint {
						return _742_v227
					}
				})(_487_v227))), _177_v18)
				var _744_v394 D6
				_ = _744_v394
				_744_v394 = Companion_D6_.Create_DC21_(_741_v393)
				var _pat_let_tv25 = _741_v393
				_ = _pat_let_tv25
				_687_v367 = (_687_v367).Update((_239_v65.F19).Plus(Companion_Default___.Fm0(_242_v68.F19, (_740_v392).Dtor_cf23(), (func(_pat_let22_0 D6) D6 {
					return func(_745_dt__update__tmp_h8 D6) D6 {
						return func(_pat_let23_0 _dafny.Sequence) D6 {
							return func(_746_dt__update_hcf30_h0 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC21_(_746_dt__update_hcf30_h0)
							}(_pat_let23_0)
						}(_pat_let_tv25)
					}(_pat_let22_0)
				}(_744_v394)).Dtor_cf30(), (_dafny.Zero).Minus((_239_v65).Fm4(_239_v65.F19, _dafny.IntOfUint32((_177_v18).Cardinality()), _155_v0, (func() bool {
					if (_687_v367).Contains(_239_v65.F19) {
						return (_687_v367).Get(_239_v65.F19).(bool)
					}
					return _155_v0
				})(), _157_globalState)), _157_globalState)), (_158_v2).Cmp(_239_v65.F19) >= 0)
				if _155_v0 {
					var _747_v395 _dafny.Int
					_ = _747_v395
					var _748_v396 _dafny.Sequence
					_ = _748_v396
					var _749_v397 _dafny.Int
					_ = _749_v397
					var _750_v398 _dafny.MultiSet
					_ = _750_v398
					var _out64 _dafny.Int
					_ = _out64
					var _out65 _dafny.Sequence
					_ = _out65
					var _out66 _dafny.Int
					_ = _out66
					var _out67 _dafny.MultiSet
					_ = _out67
					_out64, _out65, _out66, _out67 = Companion_Default___.M0(false, true, _239_v65.F19, _155_v0, _157_globalState)
					_747_v395 = _out64
					_748_v396 = _out65
					_749_v397 = _out66
					_750_v398 = _out67
					var _751_v399 _dafny.MultiSet
					_ = _751_v399
					_751_v399 = _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_747_v395)), _242_v68.F19)
					var _752_v400 _dafny.Map
					_ = _752_v400
					_752_v400 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _751_v399)
					var _753_v401 _dafny.Array
					_ = _753_v401
					var _nwElement0_32 _dafny.MultiSet = (_751_v399).Intersection(_751_v399)
					_ = _nwElement0_32
					var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(13))
					_ = _nw111
					(_nw111).ArraySet1(_nwElement0_32, 0)
					(_nw111).ArraySet1((func() _dafny.MultiSet {
						if (_752_v400).Contains(_155_v0) {
							return (_752_v400).Get(_155_v0).(_dafny.MultiSet)
						}
						return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfInt64(-279)), (Companion_Default___.SafeIndex(_239_v65.F19, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-279))).Cardinality()))).Uint32(), _239_v65.F19))
					})(), 1)
					(_nw111).ArraySet1(_751_v399, 2)
					(_nw111).ArraySet1(_dafny.MultiSetOf((_dafny.Zero).Minus(_747_v395)), 3)
					(_nw111).ArraySet1(((_751_v399).Update(_239_v65.F19, Companion_Default___.Abs(_239_v65.F19))).Union(Companion_Default___.Fm15(_157_globalState)), 4)
					(_nw111).ArraySet1(_751_v399, 5)
					(_nw111).ArraySet1(_751_v399, 6)
					(_nw111).ArraySet1(_751_v399, 7)
					(_nw111).ArraySet1(((func() _dafny.MultiSet {
						if (_752_v400).Contains(_155_v0) {
							return (_752_v400).Get(_155_v0).(_dafny.MultiSet)
						}
						return _751_v399
					})()).Difference(_751_v399), 8)
					(_nw111).ArraySet1((_751_v399).Difference(_dafny.MultiSetOf(_242_v68.F19, _239_v65.F19, _239_v65.F19, _158_v2, (_dafny.Zero).Minus(_158_v2))), 9)
					(_nw111).ArraySet1(Companion_Default___.Fm15(_157_globalState), 10)
					(_nw111).ArraySet1(_751_v399, 11)
					(_nw111).ArraySet1(_751_v399, 12)
					_753_v401 = _nw111
					_753_v401 = _753_v401
					var _754_v402 _dafny.Map
					_ = _754_v402
					_754_v402 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v0, _242_v68.F19)
					var _755_v403 _dafny.Set
					_ = _755_v403
					_755_v403 = _dafny.SetOf((_754_v402).Merge(_754_v402))
					var _756_v404 _dafny.Array
					_ = _756_v404
					var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
					_ = _nw112
					_756_v404 = _nw112
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_756_v404), 0))
					_ = _index143
					(_756_v404).ArraySet1((func() _dafny.Sequence {
						if !(_155_v0) {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(32))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg48 _dafny.Int) interface{} {
									return coer48(arg48)
								}
							}((func(_757_v228 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
								return func(_758_i50 _dafny.Int) _dafny.CodePoint {
									return (_757_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_757_v228), 0))).Int())
								}
							})(_488_v228)))
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("mamgipjj")
					})(), (_index143).Int())
					var _759_v405 _dafny.Array
					_ = _759_v405
					var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
					_ = _nw113
					_759_v405 = _nw113
					var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_759_v405), 0))
					_ = _index144
					(_759_v405).ArraySet1(_dafny.IntOfInt64(-902), (_index144).Int())
					var _760_v406 _dafny.Sequence
					_ = _760_v406
					_760_v406 = _dafny.SeqOf(_239_v65, _242_v68, _239_v65)
					var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_756_v404), 0))
					_ = _index145
					var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_759_v405), 0))
					_ = _index146
					var _rhs143 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_748_v396, Companion_Default___.Fm5(_dafny.IntOfUint32((_760_v406).Cardinality()), _157_globalState))
					_ = _rhs143
					var _rhs144 _dafny.Set = (_755_v403).Intersection(_755_v403)
					_ = _rhs144
					var _rhs145 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_177_v18, _748_v396), (Companion_Default___.SafeIndex((_239_v65.F19).Plus(_749_v397), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_177_v18, _748_v396)).Cardinality()))).Uint32(), (_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()))
					_ = _rhs145
					var _rhs146 _dafny.Int = _158_v2
					_ = _rhs146
					var _lhs114 _dafny.Array = _756_v404
					_ = _lhs114
					var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_756_v404), 0))
					_ = _lhs115
					var _lhs116 _dafny.Array = _759_v405
					_ = _lhs116
					var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_759_v405), 0))
					_ = _lhs117
					_748_v396 = _rhs143
					_755_v403 = _rhs144
					(_lhs114).ArraySet1(_rhs145, (_lhs115).Int())
					(_lhs116).ArraySet1(_rhs146, (_lhs117).Int())
					var _761_v407 _dafny.Map
					_ = _761_v407
					_761_v407 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_159_v3).Difference(_750_v398), _242_v68.F19)
					var _762_v408 _dafny.Array
					_ = _762_v408
					var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
					_ = _nw114
					_762_v408 = _nw114
					var _rhs147 _dafny.Int = (_dafny.IntOfInt64(-575)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(420))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}((func(_763_v227 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_764_i51 _dafny.Int) _dafny.CodePoint {
							return _763_v227
						}
					})(_487_v227)))).Cardinality()))
					_ = _rhs147
					var _rhs148 _dafny.Map = _761_v407
					_ = _rhs148
					var _rhs149 bool = _155_v0
					_ = _rhs149
					var _rhs150 *C0 = _239_v65
					_ = _rhs150
					var _rhs151 _dafny.Array = _762_v408
					_ = _rhs151
					var _lhs118 *GlobalState = _157_globalState
					_ = _lhs118
					var _lhs119 *GlobalState = _157_globalState
					_ = _lhs119
					_lhs118.F12 = _rhs147
					_761_v407 = _rhs148
					_lhs119.F3 = _rhs149
					_242_v68 = _rhs150
					_762_v408 = _rhs151
					(_157_globalState).F3 = (_242_v68.F19).Cmp(_239_v65.F19) < 0
				} else {
					(_239_v65).F19 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_177_v18, _177_v18)).Cardinality()), _242_v68.F19)
					_241_v67 = Companion_Default___.Fm10(_157_globalState)
					_155_v0 = _155_v0
					var _765_v410 _dafny.Array
					_ = _765_v410
					var _len0_18 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_18
					var _nw115 _dafny.Array
					_ = _nw115
					if _len0_18.Cmp(_dafny.Zero) == 0 {
						_nw115 = _dafny.NewArray(_len0_18)
					} else {
						var _init18 func(_dafny.Int) _dafny.Set = (func(_766_v67 _dafny.Map) func(_dafny.Int) _dafny.Set {
							return func(_767_i52 _dafny.Int) _dafny.Set {
								return (_dafny.SetOf(_766_v67)).Difference(func() _dafny.Set {
									var _coll18 = _dafny.NewBuilder()
									_ = _coll18
									for _iter18 := _dafny.Iterate((_dafny.SetOf(_766_v67)).Elements()); ; {
										_compr_18, _ok18 := _iter18()
										if !_ok18 {
											break
										}
										var _768_v409 _dafny.Map
										_768_v409 = interface{}(_compr_18).(_dafny.Map)
										if (_dafny.SetOf(_766_v67)).Contains(_768_v409) {
											_coll18.Add(_768_v409)
										}
									}
									return _coll18.ToSet()
								}())
							}
						})(_241_v67)
						_ = _init18
						var _element0_18 = _init18(_dafny.Zero)
						_ = _element0_18
						_nw115 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
						(_nw115).ArraySet1(_element0_18, 0)
						var _nativeLen0_18 = (_len0_18).Int()
						_ = _nativeLen0_18
						for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
							(_nw115).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
						}
					}
					_765_v410 = _nw115
					var _769_v411 D1
					_ = _769_v411
					_769_v411 = Companion_D1_.Create_DC5_(Companion_Default___.Fm2(_155_v0, _157_globalState))
					var _770_v412 _dafny.Map
					_ = _770_v412
					_770_v412 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v67, _769_v411)
					var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_765_v410), 0))
					_ = _index147
					(_765_v410).ArraySet1((func() _dafny.Set {
						var _coll19 = _dafny.NewBuilder()
						_ = _coll19
						for _iter19 := _dafny.Iterate((_770_v412).Keys().Elements()); ; {
							_compr_19, _ok19 := _iter19()
							if !_ok19 {
								break
							}
							var _771_v413 _dafny.Map
							_771_v413 = interface{}(_compr_19).(_dafny.Map)
							if (_770_v412).Contains(_771_v413) {
								_coll19.Add(_771_v413)
							}
						}
						return _coll19.ToSet()
					}()).Difference(Companion_Default___.Fm16(_157_globalState)), (_index147).Int())
					var _772_v414 _dafny.Set
					_ = _772_v414
					_772_v414 = _dafny.SetOf(_241_v67)
					var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_765_v410), 0))
					_ = _index148
					(_765_v410).ArraySet1(_772_v414, (_index148).Int())
					var _773_v415 _dafny.Array
					_ = _773_v415
					var _nw116 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
					_ = _nw116
					_773_v415 = _nw116
					var _774_v418 _dafny.Array
					_ = _774_v418
					var _nwElement0_33 _dafny.Map = _241_v67
					_ = _nwElement0_33
					var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(21))
					_ = _nw117
					(_nw117).ArraySet1(_nwElement0_33, 0)
					(_nw117).ArraySet1(_241_v67, 1)
					(_nw117).ArraySet1(func() _dafny.Map {
						var _coll20 = _dafny.NewMapBuilder()
						_ = _coll20
						for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(203), _dafny.IntOfInt64(-284))); ; {
							_compr_20, _ok20 := _iter20()
							if !_ok20 {
								break
							}
							var _775_v416 _dafny.Int
							_775_v416 = interface{}(_compr_20).(_dafny.Int)
							if ((_dafny.IntOfInt64(203)).Cmp(_775_v416) <= 0) && ((_775_v416).Cmp(_dafny.IntOfInt64(-284)) < 0) {
								_coll20.Add((_775_v416).Minus(_158_v2), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_488_v228).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_488_v228), 0))).Int()), !(_155_v0))).Cardinality())
							}
						}
						return _coll20.ToMap()
					}(), 2)
					(_nw117).ArraySet1(_241_v67, 3)
					(_nw117).ArraySet1(_241_v67, 4)
					(_nw117).ArraySet1(_241_v67, 5)
					(_nw117).ArraySet1(_241_v67, 6)
					(_nw117).ArraySet1(_241_v67, 7)
					(_nw117).ArraySet1(_241_v67, 8)
					(_nw117).ArraySet1(_241_v67, 9)
					(_nw117).ArraySet1(_241_v67, 10)
					(_nw117).ArraySet1(_241_v67, 11)
					(_nw117).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v68.F19, _dafny.IntOfInt64(908)), 12)
					(_nw117).ArraySet1(_241_v67, 13)
					(_nw117).ArraySet1(_241_v67, 14)
					(_nw117).ArraySet1(_241_v67, 15)
					(_nw117).ArraySet1(_241_v67, 16)
					(_nw117).ArraySet1(_241_v67, 17)
					(_nw117).ArraySet1(_241_v67, 18)
					(_nw117).ArraySet1((_241_v67).Update(_242_v68.F19, (_dafny.Zero).Minus(_239_v65.F19)), 19)
					(_nw117).ArraySet1(func() _dafny.Map {
						var _coll21 = _dafny.NewMapBuilder()
						_ = _coll21
						for _iter21 := _dafny.Iterate((_241_v67).Keys().Elements()); ; {
							_compr_21, _ok21 := _iter21()
							if !_ok21 {
								break
							}
							var _776_v417 _dafny.Int
							_776_v417 = interface{}(_compr_21).(_dafny.Int)
							if (_241_v67).Contains(_776_v417) {
								_coll21.Add((_776_v417).Minus(_242_v68.F19), (_159_v3).Cardinality())
							}
						}
						return _coll21.ToMap()
					}(), 20)
					_774_v418 = _nw117
					var _777_v419 _dafny.Map
					_ = _777_v419
					_777_v419 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v418, _773_v415)
					_773_v415 = (func() _dafny.Array {
						if (_777_v419).Contains(_774_v418) {
							return (_777_v419).Get(_774_v418).(_dafny.Array)
						}
						return _773_v415
					})()
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_385_v167), 0))
	_ = _index149
	(_385_v167).ArraySet1(_155_v0, (_index149).Int())
	var _778_v420 _dafny.Array
	_ = _778_v420
	var _nw118 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(29))
	_ = _nw118
	_778_v420 = _nw118
	var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_385_v167), 0))
	_ = _index150
	var _rhs152 bool = (_dafny.IntOfInt64(519)).Cmp(Companion_Default___.SafeModuloInt(_239_v65.F19, _dafny.IntOfUint32((_dafny.SeqOf(!(_155_v0), _155_v0)).Cardinality()))) != 0
	_ = _rhs152
	var _rhs153 _dafny.Array = (func() _dafny.Array {
		if _155_v0 {
			return _778_v420
		}
		return _778_v420
	})()
	_ = _rhs153
	var _lhs120 _dafny.Array = _385_v167
	_ = _lhs120
	var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_385_v167), 0))
	_ = _lhs121
	(_lhs120).ArraySet1(_rhs152, (_lhs121).Int())
	_778_v420 = _rhs153
	_dafny.Print(_155_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_156_v1, _dafny.SeqOf(false, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F10().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_157_globalState).F18(), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_158_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v3).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_177_v18, _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v63).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false), false).UpdateUnsafe(_dafny.SetOf(true), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_238_v64, _dafny.SeqOf(_dafny.IntOfInt64(270))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_239_v65.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_v66).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v67).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(270))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_242_v68.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_243_v69).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_244_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_325_v126).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v167).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v167).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_386_v168).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_386_v168).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_386_v168).Dtor_cf26()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_386_v168).Dtor_cf26()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_386_v168).Dtor_cf27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_387_v169).Dtor_cf29()).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_387_v169).Dtor_cf29()).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_387_v169).Dtor_cf29()).Dtor_cf26()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_387_v169).Dtor_cf29()).Dtor_cf26()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_387_v169).Dtor_cf29()).Dtor_cf27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_388_v170).Dtor_cf29()).Dtor_cf29()).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_388_v170).Dtor_cf29()).Dtor_cf29()).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_388_v170).Dtor_cf29()).Dtor_cf29()).Dtor_cf26()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_388_v170).Dtor_cf29()).Dtor_cf29()).Dtor_cf26()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_388_v170).Dtor_cf29()).Dtor_cf29()).Dtor_cf27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_487_v227)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_488_v228).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_488_v228).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_488_v228).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_488_v228).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_573_v281).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_573_v281).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_573_v281).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D0)).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_682_v363).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D0)).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_683_v365).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_685_v366).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(650), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_686_i42)
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
	Cf2 _dafny.Int
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 bool
	Cf5 bool
	Cf6 _dafny.Int
	Cf7 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 bool, Cf5 bool, Cf6 _dafny.Int, Cf7 bool) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6, Cf7}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
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

func (_this D0) Dtor_cf7() bool {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7
		}
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

type D1_DC4 struct {
	Cf9 *C0
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf9 *C0) D1 {
	return D1{D1_DC4{Cf9}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf10 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf10 bool) D1 {
	return D1{D1_DC5{Cf10}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf8 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf8 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf8}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC6 struct {
	Cf11 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf11 D1) D1 {
	return D1{D1_DC6{Cf11}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_((*C0)(nil))
}

func (_this D1) Dtor_cf9() *C0 {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) Dtor_cf11() D1 {
	return _this.Get_().(D1_DC6).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + data.Cf8.VerbatimString(true) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf9 == data2.Cf9
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf10 == data2.Cf10
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf8.Equals(data2.Cf8)
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D2_DC8 struct {
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_() D2 {
	return D2{D2_DC8{}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf13 _dafny.Int
	Cf14 *C0
	Cf15 *C0
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf13 _dafny.Int, Cf14 *C0, Cf15 *C0) D2 {
	return D2{D2_DC9{Cf13, Cf14, Cf15}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC7 struct {
	Cf12 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 _dafny.Array) D2 {
	return D2{D2_DC7{Cf12}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC10 struct {
	Cf16 D2
}

func (D2_DC10) isD2() {}

func (CompanionStruct_D2_) Create_DC10_(Cf16 D2) D2 {
	return D2{D2_DC10{Cf16}}
}

func (_this D2) Is_DC10() bool {
	_, ok := _this.Get_().(D2_DC10)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_()
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf13
}

func (_this D2) Dtor_cf14() *C0 {
	return _this.Get_().(D2_DC9).Cf14
}

func (_this D2) Dtor_cf15() *C0 {
	return _this.Get_().(D2_DC9).Cf15
}

func (_this D2) Dtor_cf12() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf16() D2 {
	return _this.Get_().(D2_DC10).Cf16
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC10:
		{
			return "D2.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC8:
		{
			_, ok := other.Get_().(D2_DC8)
			return ok
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf12 == data2.Cf12
		}
	case D2_DC10:
		{
			data2, ok := other.Get_().(D2_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D3_DC12 struct {
	Cf18 _dafny.Array
	Cf19 bool
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf18 _dafny.Array, Cf19 bool) D3 {
	return D3{D3_DC12{Cf18, Cf19}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC11 struct {
	Cf17 _dafny.Map
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf17 _dafny.Map) D3 {
	return D3{D3_DC11{Cf17}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC13 struct {
	Cf20 D3
}

func (D3_DC13) isD3() {}

func (CompanionStruct_D3_) Create_DC13_(Cf20 D3) D3 {
	return D3{D3_DC13{Cf20}}
}

func (_this D3) Is_DC13() bool {
	_, ok := _this.Get_().(D3_DC13)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC12_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D3) Dtor_cf18() _dafny.Array {
	return _this.Get_().(D3_DC12).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC12).Cf19
}

func (_this D3) Dtor_cf17() _dafny.Map {
	return _this.Get_().(D3_DC11).Cf17
}

func (_this D3) Dtor_cf20() D3 {
	return _this.Get_().(D3_DC13).Cf20
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC13:
		{
			return "D3.DC13" + "(" + _dafny.String(data.Cf20) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	case D3_DC13:
		{
			data2, ok := other.Get_().(D3_DC13)
			return ok && data1.Cf20.Equals(data2.Cf20)
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

type D4_DC15 struct {
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_() D4 {
	return D4{D4_DC15{}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC14 struct {
	Cf21 _dafny.Sequence
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf21 _dafny.Sequence) D4 {
	return D4{D4_DC14{Cf21}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC16 struct {
	Cf22 D4
}

func (D4_DC16) isD4() {}

func (CompanionStruct_D4_) Create_DC16_(Cf22 D4) D4 {
	return D4{D4_DC16{Cf22}}
}

func (_this D4) Is_DC16() bool {
	_, ok := _this.Get_().(D4_DC16)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC15_()
}

func (_this D4) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf22() D4 {
	return _this.Get_().(D4_DC16).Cf22
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC15:
		{
			return "D4.DC15"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC16:
		{
			return "D4.DC16" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC15:
		{
			_, ok := other.Get_().(D4_DC15)
			return ok
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D4_DC16:
		{
			data2, ok := other.Get_().(D4_DC16)
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

type D5_DC18 struct {
	Cf24 _dafny.Int
	Cf25 bool
	Cf26 _dafny.Array
	Cf27 bool
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf24 _dafny.Int, Cf25 bool, Cf26 _dafny.Array, Cf27 bool) D5 {
	return D5{D5_DC18{Cf24, Cf25, Cf26, Cf27}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC19 struct {
	Cf28 bool
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf28 bool) D5 {
	return D5{D5_DC19{Cf28}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

type D5_DC17 struct {
	Cf23 _dafny.CodePoint
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf23 _dafny.CodePoint) D5 {
	return D5{D5_DC17{Cf23}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC20 struct {
	Cf29 D5
}

func (D5_DC20) isD5() {}

func (CompanionStruct_D5_) Create_DC20_(Cf29 D5) D5 {
	return D5{D5_DC20{Cf29}}
}

func (_this D5) Is_DC20() bool {
	_, ok := _this.Get_().(D5_DC20)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC18_(_dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D5) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf24
}

func (_this D5) Dtor_cf25() bool {
	return _this.Get_().(D5_DC18).Cf25
}

func (_this D5) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D5_DC18).Cf26
}

func (_this D5) Dtor_cf27() bool {
	return _this.Get_().(D5_DC18).Cf27
}

func (_this D5) Dtor_cf28() bool {
	return _this.Get_().(D5_DC19).Cf28
}

func (_this D5) Dtor_cf23() _dafny.CodePoint {
	return _this.Get_().(D5_DC17).Cf23
}

func (_this D5) Dtor_cf29() D5 {
	return _this.Get_().(D5_DC20).Cf29
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC20:
		{
			return "D5.DC20" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf28 == data2.Cf28
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D5_DC20:
		{
			data2, ok := other.Get_().(D5_DC20)
			return ok && data1.Cf29.Equals(data2.Cf29)
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
}

func (D6_DC22) isD6() {}

func (CompanionStruct_D6_) Create_DC22_() D6 {
	return D6{D6_DC22{}}
}

func (_this D6) Is_DC22() bool {
	_, ok := _this.Get_().(D6_DC22)
	return ok
}

type D6_DC23 struct {
	Cf31 bool
	Cf32 _dafny.Int
	Cf33 _dafny.Int
}

func (D6_DC23) isD6() {}

func (CompanionStruct_D6_) Create_DC23_(Cf31 bool, Cf32 _dafny.Int, Cf33 _dafny.Int) D6 {
	return D6{D6_DC23{Cf31, Cf32, Cf33}}
}

func (_this D6) Is_DC23() bool {
	_, ok := _this.Get_().(D6_DC23)
	return ok
}

type D6_DC21 struct {
	Cf30 _dafny.Sequence
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf30 _dafny.Sequence) D6 {
	return D6{D6_DC21{Cf30}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

type D6_DC24 struct {
	Cf34 D6
}

func (D6_DC24) isD6() {}

func (CompanionStruct_D6_) Create_DC24_(Cf34 D6) D6 {
	return D6{D6_DC24{Cf34}}
}

func (_this D6) Is_DC24() bool {
	_, ok := _this.Get_().(D6_DC24)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC22_()
}

func (_this D6) Dtor_cf31() bool {
	return _this.Get_().(D6_DC23).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D6_DC23).Cf32
}

func (_this D6) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D6_DC23).Cf33
}

func (_this D6) Dtor_cf30() _dafny.Sequence {
	return _this.Get_().(D6_DC21).Cf30
}

func (_this D6) Dtor_cf34() D6 {
	return _this.Get_().(D6_DC24).Cf34
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC22:
		{
			return "D6.DC22"
		}
	case D6_DC23:
		{
			return "D6.DC23" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC21:
		{
			return "D6.DC21" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC24:
		{
			return "D6.DC24" + "(" + _dafny.String(data.Cf34) + ")"
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
			_, ok := other.Get_().(D6_DC22)
			return ok
		}
	case D6_DC23:
		{
			data2, ok := other.Get_().(D6_DC23)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf30.Equals(data2.Cf30)
		}
	case D6_DC24:
		{
			data2, ok := other.Get_().(D6_DC24)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D7_DC25 struct {
	Cf35 _dafny.Set
}

func (D7_DC25) isD7() {}

func (CompanionStruct_D7_) Create_DC25_(Cf35 _dafny.Set) D7 {
	return D7{D7_DC25{Cf35}}
}

func (_this D7) Is_DC25() bool {
	_, ok := _this.Get_().(D7_DC25)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D7) Dtor_cf35() _dafny.Set {
	return _this.Get_().(D7_DC25).Cf35
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC25:
		{
			return "D7.DC25" + "(" + _dafny.String(data.Cf35) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC25:
		{
			data2, ok := other.Get_().(D7_DC25)
			return ok && data1.Cf35.Equals(data2.Cf35)
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

type D8_DC26 struct {
	Cf36 _dafny.Set
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf36 _dafny.Set) D8 {
	return D8{D8_DC26{Cf36}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D8) Dtor_cf36() _dafny.Set {
	return _this.Get_().(D8_DC26).Cf36
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf36) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC26:
		{
			data2, ok := other.Get_().(D8_DC26)
			return ok && data1.Cf36.Equals(data2.Cf36)
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

type D9_DC28 struct {
	Cf38 _dafny.Int
	Cf39 _dafny.Int
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf38 _dafny.Int, Cf39 _dafny.Int) D9 {
	return D9{D9_DC28{Cf38, Cf39}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

type D9_DC29 struct {
	Cf40 _dafny.Int
	Cf41 _dafny.Int
}

func (D9_DC29) isD9() {}

func (CompanionStruct_D9_) Create_DC29_(Cf40 _dafny.Int, Cf41 _dafny.Int) D9 {
	return D9{D9_DC29{Cf40, Cf41}}
}

func (_this D9) Is_DC29() bool {
	_, ok := _this.Get_().(D9_DC29)
	return ok
}

type D9_DC27 struct {
	Cf37 _dafny.Set
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf37 _dafny.Set) D9 {
	return D9{D9_DC27{Cf37}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC30 struct {
	Cf42 D9
}

func (D9_DC30) isD9() {}

func (CompanionStruct_D9_) Create_DC30_(Cf42 D9) D9 {
	return D9{D9_DC30{Cf42}}
}

func (_this D9) Is_DC30() bool {
	_, ok := _this.Get_().(D9_DC30)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC28_(_dafny.Zero, _dafny.Zero)
}

func (_this D9) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D9_DC28).Cf38
}

func (_this D9) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D9_DC28).Cf39
}

func (_this D9) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D9_DC29).Cf40
}

func (_this D9) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D9_DC29).Cf41
}

func (_this D9) Dtor_cf37() _dafny.Set {
	return _this.Get_().(D9_DC27).Cf37
}

func (_this D9) Dtor_cf42() D9 {
	return _this.Get_().(D9_DC30).Cf42
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D9_DC29:
		{
			return "D9.DC29" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D9_DC30:
		{
			return "D9.DC30" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC28:
		{
			data2, ok := other.Get_().(D9_DC28)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D9_DC29:
		{
			data2, ok := other.Get_().(D9_DC29)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Cmp(data2.Cf41) == 0
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf37.Equals(data2.Cf37)
		}
	case D9_DC30:
		{
			data2, ok := other.Get_().(D9_DC30)
			return ok && data1.Cf42.Equals(data2.Cf42)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F2   _dafny.Int
	F3   bool
	F8   _dafny.Int
	F12  _dafny.Int
	_f1  bool
	_f4  _dafny.Int
	_f5  _dafny.Int
	_f6  bool
	_f7  _dafny.Int
	_f9  _dafny.Int
	_f10 _dafny.Sequence
	_f11 bool
	_f13 bool
	_f14 _dafny.Int
	_f15 bool
	_f16 _dafny.Int
	_f17 _dafny.Int
	_f18 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F2 = _dafny.Zero
	_this.F3 = false
	_this.F8 = _dafny.Zero
	_this.F12 = _dafny.Zero
	_this._f1 = false
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this._f7 = _dafny.Zero
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.EmptySeq
	_this._f11 = false
	_this._f13 = false
	_this._f14 = _dafny.Zero
	_this._f15 = false
	_this._f16 = _dafny.Zero
	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 _dafny.Int, f3 bool, f4 _dafny.Int, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 _dafny.Int, f9 _dafny.Int, f10 _dafny.Sequence, f11 bool, f12 _dafny.Int, f13 bool, f14 _dafny.Int, f15 bool, f16 _dafny.Int, f17 _dafny.Int, f18 _dafny.Sequence) {
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
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
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
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Sequence {
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
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() bool {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Sequence {
	{
		return _this._f18
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F19 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F19 = _dafny.Zero
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

func (_this *C0) Ctor__(f19 _dafny.Int) {
	{
		(_this).F19 = f19
	}
}
func (_this *C0) Fm3(globalState *GlobalState) _dafny.Int {
	{
		return _this.F19
	}
}
func (_this *C0) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus(_this.F19)).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false)).Cardinality())).Cardinality()))
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
