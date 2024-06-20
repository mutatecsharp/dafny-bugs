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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	return ((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-842), _dafny.IntOfInt64(432))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-842)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(432)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_0_v0, _dafny.IntOfInt64(308)), _dafny.IntOfInt64(166))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Plus(_dafny.IntOfUint32(((Companion_D13_.Create_DC26_(_dafny.UnicodeSeqOfUtf8Bytes("gtqau"), true, true)).Dtor_cf33()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("juk"), _dafny.UnicodeSeqOfUtf8Bytes("liutef")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(902))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})), _dafny.UnicodeSeqOfUtf8Bytes("olrx")))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Map, globalState *GlobalState) bool {
	return !((func() bool {
		if false {
			return true
		}
		return (func() bool {
			if false {
				return false
			}
			return true
		})()
	})())
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cacr")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfInt64(792))).Cardinality(), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(581), false))).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(310), _dafny.IntOfInt64(336))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(310)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(336)) < 0) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_2_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-856))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}))).Cardinality())), (Companion_D24_.Create_DC64_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg1 _dafny.Int) interface{} {
						return coer1(arg1)
					}
				}(func(_3_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				}))).Cardinality()), true, false)).Dtor_cf93())
			}
		}
		return _coll1.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D0 {
	if (func() bool {
		if true {
			return false
		}
		return true
	})() {
		return Companion_D0_.Create_DC0_(!(true))
	} else {
		return Companion_D0_.Create_DC1_(Companion_D0_.Create_DC0_(false))
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D20 = Companion_D20_.Create_DC49_(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(500), _dafny.IntOfInt64(743))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(500)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(743)) < 0) {
				_coll2.Add(Companion_Default___.SafeDivisionInt(_5_v0, _dafny.IntOfInt64(893)), _dafny.CodePoint('v'))
			}
		}
		return _coll2.ToMap()
	}())
	_ = _source0
	if _source0.Is_DC50() {
		var _6___mcc_h0 _dafny.Int = _source0.Get_().(D20_DC50).Cf68
		_ = _6___mcc_h0
		var _7___mcc_h1 bool = _source0.Get_().(D20_DC50).Cf69
		_ = _7___mcc_h1
		var _8___mcc_h2 bool = _source0.Get_().(D20_DC50).Cf70
		_ = _8___mcc_h2
		var _9_cf70 bool = _8___mcc_h2
		_ = _9_cf70
		var _10_cf69 bool = _7___mcc_h1
		_ = _10_cf69
		var _11_cf68 _dafny.Int = _6___mcc_h0
		_ = _11_cf68
		return _dafny.SeqOf(_dafny.MultiSetOf(_10_cf69), _dafny.MultiSetOf(false), _dafny.MultiSetOf(_9_cf70))
	} else if _source0.Is_DC51() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(false)), _dafny.SeqOf(_dafny.MultiSetOf(true, true)))
	} else {
		var _12___mcc_h3 _dafny.Map = _source0.Get_().(D20_DC49).Cf67
		_ = _12___mcc_h3
		var _13_cf67 _dafny.Map = _12___mcc_h3
		_ = _13_cf67
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(true), _dafny.MultiSetOf(true)), _dafny.SeqOf(_dafny.MultiSetOf(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(872))).Cardinality(), (_dafny.MultiSetOf(_dafny.CodePoint('v'), _dafny.CodePoint('n'))).Cardinality(), _dafny.IntOfInt64(-229), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(358)), _dafny.SeqOf(_dafny.IntOfInt64(636))))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(864), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mci")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 bool, p2 D4, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.CodePoint('h'), (Companion_D3_.Create_DC6_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(743))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})), _dafny.CodePoint('c'), false)).Dtor_cf7())).Intersection(_dafny.MultiSetOf(_dafny.CodePoint('w'), _dafny.CodePoint('q')))).Union(_dafny.MultiSetOf(_dafny.CodePoint('a')))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if true {
			return _dafny.SetOf(false)
		}
		return _dafny.SetOf(true)
	})()).Difference((_dafny.SetOf(true)).Intersection(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('b')
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 bool, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC7_((_dafny.Zero).Minus((Companion_D19_.Create_DC47_(Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), true)), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(352))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_15_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D17_.Create_DC39_(_dafny.IntOfInt64(838), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(353))).Cardinality(), (func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tynvdcxoj")).Cardinality())))).Cardinality())).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _16_v0 _dafny.CodePoint
				_16_v0 = interface{}(_compr_3).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tynvdcxoj")).Cardinality())))).Cardinality())).Contains(_16_v0) {
					_coll3.Add(_16_v0, true)
				}
			}
			return _coll3.ToMap()
		}()).Cardinality()), !(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), Companion_D13_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-51), (_dafny.MultiSetOf(_dafny.IntOfInt64(52))).Cardinality()))), _dafny.IntOfInt64(358)))).Cardinality()
	}))).Cardinality()))).Cardinality(), false)).Dtor_cf63()))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfjbncjt")).Cardinality()))).Cardinality())).Cmp(_dafny.IntOfInt64(-845)) >= 0 {
		return _dafny.SeqOf(false)
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Union(_dafny.SetOf(true))).Intersection(_dafny.SetOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(171))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_17_i0 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-380), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hf")).Cardinality())))).Cardinality()
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_18_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-136)), _dafny.SeqOf(_dafny.IntOfInt64(-24)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(906))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xtfitcgkw")).Cardinality())), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(44))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_19_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-784)
	}))).Cardinality()))).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(844))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_20_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vpemvormm")).Cardinality())
	})))
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) D13 {
	if !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(true), _dafny.SeqOf(false, true))) {
		return Companion_D13_.Create_DC26_(_dafny.UnicodeSeqOfUtf8Bytes("vmuja"), !(!(false)), false)
	} else {
		return Companion_D13_.Create_DC26_((Companion_D13_.Create_DC26_(_dafny.UnicodeSeqOfUtf8Bytes("lhaqxlf"), true, false)).Dtor_cf33(), true, true)
	}
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), !(true)))).Intersection(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), !(true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false)))).Intersection((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), true), func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('c'))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _21_v0 _dafny.CodePoint
			_21_v0 = interface{}(_compr_4).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('c'))).Contains(_21_v0) {
				_coll4.Add(_21_v0, false)
			}
		}
		return _coll4.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false))).Union(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(922))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_22_i0 _dafny.Int) _dafny.Map {
		return func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), true)).Keys().Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _23_v1 _dafny.CodePoint
				_23_v1 = interface{}(_compr_5).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), true)).Contains(_23_v1) {
					_coll5.Add(_23_v1, false)
				}
			}
			return _coll5.ToMap()
		}()
	})))))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("qs"), _dafny.CodePoint('p'), (_dafny.MultiSetOf(false)).IsProperSubsetOf(_dafny.MultiSetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.CodePoint, p1 _dafny.Map, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(8), _dafny.IntOfInt64(496), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-312), (_dafny.MultiSetOf(_dafny.IntOfInt64(582), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fre")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(784))).Cardinality()))).Cardinality())).Difference(_dafny.SetOf((_dafny.Zero).Minus((_dafny.SetOf(true, false)).Cardinality()), _dafny.IntOfInt64(386)))).Difference(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(false, true)).Cardinality())).Cardinality()), _dafny.IntOfInt64(190), _dafny.IntOfInt64(-913)))
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.SetOf(false, false, true)).IsSubsetOf(_dafny.SetOf(true)) {
		return _dafny.CodePoint('v')
	} else {
		return _dafny.CodePoint('p')
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Set, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	var _source1 D18 = Companion_D18_.Create_DC42_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
	_ = _source1
	if _source1.Is_DC43() {
		return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(642))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_24_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality()), _dafny.IntOfInt64(-885), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), _dafny.IntOfInt64(538), _dafny.IntOfInt64(588))).Cardinality())
	} else if _source1.Is_DC44() {
		return _dafny.MultiSetOf(_dafny.IntOfInt64(158), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
	} else {
		var _25___mcc_h0 _dafny.Map = _source1.Get_().(D18_DC42).Cf58
		_ = _25___mcc_h0
		var _26_cf58 _dafny.Map = _25___mcc_h0
		_ = _26_cf58
		return _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-301)), _dafny.IntOfInt64(-858))
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(915), _dafny.IntOfInt64(103))).Intersection(_dafny.SetOf(_dafny.IntOfInt64(833)))).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_27_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rsvm")).Cardinality())
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Intersection(_dafny.SetOf(true))).Intersection(_dafny.SetOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D14 {
	if (true) == (true) {
		return Companion_D14_.Create_DC28_(_dafny.SetOf(_dafny.IntOfInt64(396)))
	} else if true {
		return Companion_D14_.Create_DC28_(func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(88), _dafny.IntOfInt64(346))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _28_v0 _dafny.Int
				_28_v0 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(88)).Cmp(_28_v0) <= 0) && ((_28_v0).Cmp(_dafny.IntOfInt64(346)) < 0) {
					_coll6.Add((_28_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(610))).Cardinality())))
				}
			}
			return _coll6.ToSet()
		}())
	} else {
		return Companion_D14_.Create_DC28_(func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(896), _dafny.IntOfInt64(614))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _29_v1 _dafny.Int
				_29_v1 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(896)).Cmp(_29_v1) <= 0) && ((_29_v1).Cmp(_dafny.IntOfInt64(614)) < 0) {
					_coll7.Add((_29_v1).Minus(_dafny.IntOfInt64(-280)))
				}
			}
			return _coll7.ToSet()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) _dafny.MultiSet {
	return (Companion_D25_.Create_DC67_(_dafny.MultiSetOf(_dafny.IntOfInt64(-892), _dafny.IntOfInt64(188)))).Dtor_cf98()
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('u')
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source2 D21 = Companion_D21_.Create_DC54_(_dafny.IntOfInt64(262), _dafny.IntOfInt64(529), !(!(true)))
	_ = _source2
	if _source2.Is_DC53() {
		var _30___mcc_h0 _dafny.Array = _source2.Get_().(D21_DC53).Cf72
		_ = _30___mcc_h0
		var _31___mcc_h1 bool = _source2.Get_().(D21_DC53).Cf73
		_ = _31___mcc_h1
		var _32___mcc_h2 _dafny.MultiSet = _source2.Get_().(D21_DC53).Cf74
		_ = _32___mcc_h2
		var _33___mcc_h3 _dafny.Array = _source2.Get_().(D21_DC53).Cf75
		_ = _33___mcc_h3
		var _34_cf75 _dafny.Array = _33___mcc_h3
		_ = _34_cf75
		var _35_cf74 _dafny.MultiSet = _32___mcc_h2
		_ = _35_cf74
		var _36_cf73 bool = _31___mcc_h1
		_ = _36_cf73
		var _37_cf72 _dafny.Array = _30___mcc_h0
		_ = _37_cf72
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_cf73, _dafny.IntOfInt64(449))
	} else if _source2.Is_DC54() {
		var _38___mcc_h4 _dafny.Int = _source2.Get_().(D21_DC54).Cf76
		_ = _38___mcc_h4
		var _39___mcc_h5 _dafny.Int = _source2.Get_().(D21_DC54).Cf77
		_ = _39___mcc_h5
		var _40___mcc_h6 bool = _source2.Get_().(D21_DC54).Cf78
		_ = _40___mcc_h6
		var _41_cf78 bool = _40___mcc_h6
		_ = _41_cf78
		var _42_cf77 _dafny.Int = _39___mcc_h5
		_ = _42_cf77
		var _43_cf76 _dafny.Int = _38___mcc_h4
		_ = _43_cf76
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_cf78, _42_cf77)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_cf78, _42_cf77))
	} else if _source2.Is_DC52() {
		var _44___mcc_h7 _dafny.Map = _source2.Get_().(D21_DC52).Cf71
		_ = _44___mcc_h7
		var _45_cf71 _dafny.Map = _44___mcc_h7
		_ = _45_cf71
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('b'))).Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _46_v0 _dafny.CodePoint
					_46_v0 = interface{}(_compr_8).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('b')), _46_v0) {
						_coll8.Add(_46_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-442))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg12 _dafny.Int) interface{} {
								return coer12(arg12)
							}
						}(func(_47_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('c')
						}))).Cardinality()))
					}
				}
				return _coll8.ToMap()
			}()).Cardinality())
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(278))
		}
	} else {
		var _48___mcc_h8 D21 = _source2.Get_().(D21_DC55).Cf79
		_ = _48___mcc_h8
		var _49_cf79 D21 = _48___mcc_h8
		_ = _49_cf79
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("abldcs")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(542)))
	}
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-663))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_50_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})))
}
func (_static *CompanionStruct_Default___) Fm37(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-504)))
}
func (_static *CompanionStruct_Default___) Fm38(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SetOf(_dafny.CodePoint('i'), _dafny.CodePoint('g')), _dafny.SetOf(_dafny.CodePoint('v')), _dafny.SetOf(_dafny.CodePoint('p'), _dafny.CodePoint('y'), _dafny.CodePoint('w'), _dafny.CodePoint('e')))).Intersection(func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('a')), _dafny.SetOf(_dafny.CodePoint('j')))).Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _51_v0 _dafny.Set
				_51_v0 = interface{}(_compr_10).(_dafny.Set)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('a')), _dafny.SetOf(_dafny.CodePoint('j'))), _51_v0) {
					_coll10.Add(_51_v0)
				}
			}
			return _coll10.ToSet()
		}()).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _52_v1 _dafny.Set
			_52_v1 = interface{}(_compr_9).(_dafny.Set)
			if (func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('a')), _dafny.SetOf(_dafny.CodePoint('j')))).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _53_v0 _dafny.Set
					_53_v0 = interface{}(_compr_11).(_dafny.Set)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('a')), _dafny.SetOf(_dafny.CodePoint('j'))), _53_v0) {
						_coll11.Add(_53_v0)
					}
				}
				return _coll11.ToSet()
			}()).Contains(_52_v1) {
				_coll9.Add(_52_v1)
			}
		}
		return _coll9.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-361))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_54_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(336))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_55_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(476)
	})))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(759))).Cardinality())), _dafny.SeqOf(!(true)))).Merge(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-784)), _dafny.SeqOf(_dafny.IntOfInt64(508), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('o'), _dafny.CodePoint('l'))).Cardinality()), _dafny.IntOfInt64(486))).Cardinality())))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _56_v0 _dafny.Sequence
			_56_v0 = interface{}(_compr_12).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-784)), _dafny.SeqOf(_dafny.IntOfInt64(508), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('o'), _dafny.CodePoint('l'))).Cardinality()), _dafny.IntOfInt64(486))).Cardinality()))), _56_v0) {
				_coll12.Add(_56_v0, _dafny.SeqOf(false))
			}
		}
		return _coll12.ToMap()
	}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()))).Cardinality()), _dafny.SeqOf(true, false))).Merge(func() _dafny.Map {
		var _coll13 = _dafny.NewMapBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(317)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(267))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_57_i2 _dafny.Int) _dafny.Int {
			return (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(960))).Cardinality())).Cardinality()
		})))).Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _58_v1 _dafny.Sequence
			_58_v1 = interface{}(_compr_13).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(317)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(267))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}(func(_57_i2 _dafny.Int) _dafny.Int {
				return (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(960))).Cardinality())).Cardinality()
			})))).Contains(_58_v1) {
				_coll13.Add(_58_v1, _dafny.SeqOf(true))
			}
		}
		return _coll13.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) D16 {
	return Companion_D16_.Create_DC36_(true, (_dafny.MultiSetOf(false, false)).Union(_dafny.MultiSetOf(true)), (_dafny.IntOfInt64(-574)).Cmp(_dafny.IntOfInt64(499)) != 0)
}
func (_static *CompanionStruct_Default___) Fm41(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(true))
}
func (_static *CompanionStruct_Default___) Fm42(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC10_((_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_D20_.Create_DC50_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uh")).Cardinality()))).Cardinality()), false, true)).Dtor_cf68())), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("owettajm")).Cardinality())).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(true))).Cardinality())), (func() bool {
		if !(false) {
			return false
		}
		return !(!(false))
	})(), !(true), (_dafny.MultiSetOf(false)).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm43(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), Companion_D4_.Create_DC10_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("memjruqa")).Cardinality()), _dafny.IntOfInt64(-160), true, true, (_dafny.MultiSetOf(_dafny.IntOfInt64(404), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(885), func() _dafny.Set {
		var _coll14 = _dafny.NewBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate(((Companion_D26_.Create_DC71_(_dafny.MultiSetOf(_dafny.CodePoint('p')))).Dtor_cf100()).Elements()); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _59_v0 _dafny.CodePoint
			_59_v0 = interface{}(_compr_14).(_dafny.CodePoint)
			if ((Companion_D26_.Create_DC71_(_dafny.MultiSetOf(_dafny.CodePoint('p')))).Dtor_cf100()).Contains(_59_v0) {
				_coll14.Add(_59_v0)
			}
		}
		return _coll14.ToSet()
	}())).Cardinality(), _dafny.IntOfInt64(955), _dafny.IntOfInt64(236))).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D4_.Create_DC10_(_dafny.IntOfInt64(285), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("exg")).Cardinality())), true, true, _dafny.IntOfInt64(113)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D4_.Create_DC10_(_dafny.IntOfInt64(655), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()), true, true, _dafny.IntOfInt64(849)))))
}
func (_static *CompanionStruct_Default___) Fm44(globalState *GlobalState) D13 {
	if !(true) {
		return Companion_D13_.Create_DC27_(true)
	} else {
		return Companion_D13_.Create_DC27_(true)
	}
}
func (_static *CompanionStruct_Default___) Fm45(globalState *GlobalState) D13 {
	if false {
		return Companion_D13_.Create_DC25_(func() _dafny.Map {
			var _coll15 = _dafny.NewMapBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(190), _dafny.IntOfInt64(-939))); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _60_v0 _dafny.Int
				_60_v0 = interface{}(_compr_15).(_dafny.Int)
				if ((_dafny.IntOfInt64(190)).Cmp(_60_v0) <= 0) && ((_60_v0).Cmp(_dafny.IntOfInt64(-939)) < 0) {
					_coll15.Add((_60_v0).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(743), _dafny.IntOfInt64(-433))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(226), _dafny.IntOfInt64(423))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(622))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}(func(_61_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					}))).Cardinality())))).Cardinality()), _dafny.IntOfInt64(-545))
				}
			}
			return _coll15.ToMap()
		}())
	} else {
		return Companion_D13_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())), _dafny.IntOfInt64(652)))
	}
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(323))).Uint32(), func(coer19 func(_dafny.Int) D14) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_62_i0 _dafny.Int) D14 {
		return Companion_D14_.Create_DC30_(_dafny.SeqOf(_dafny.IntOfInt64(691), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-832))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_63_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xxqknse")).Cardinality())), _dafny.SeqOf(false))
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-332))).Uint32(), func(coer21 func(_dafny.Int) D14) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_64_i2 _dafny.Int) D14 {
		return Companion_D14_.Create_DC30_(_dafny.SeqOf(_dafny.IntOfInt64(94)), _dafny.SeqOf(!(false), true, true))
	})))
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		if true {
			return _dafny.SetOf(true, false)
		}
		return _dafny.SetOf(false)
	})()
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-648)).Minus(_dafny.IntOfInt64(583)), (func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('l'))).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _65_v0 _dafny.CodePoint
			_65_v0 = interface{}(_compr_16).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('l')), _65_v0) {
				_coll16.Add(_65_v0, true)
			}
		}
		return _coll16.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false)))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) D19 {
	return Companion_D19_.Create_DC46_(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D23_.Create_DC60_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!(false))))), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("unbxcg")).Cardinality()))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D23_.Create_DC60_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(false)))), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(366))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_66_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))).Cardinality()))), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(Companion_D14_.Create_DC28_(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))))).Cardinality(), false)).Cardinality()).Times(_dafny.IntOfInt64(711)))
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("obwgo")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(40))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_67_i0 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(false, !(true))).Cardinality()
	}))))
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(428))).Cardinality()))).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dgjnixnjt")).Cardinality()), _dafny.Zero)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(532), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("umwrjbo")).Cardinality()))).Cardinality(), true)).Cardinality())).Cardinality()), true)), _dafny.SeqOf(func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(620), _dafny.IntOfInt64(233))); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _68_v0 _dafny.Int
			_68_v0 = interface{}(_compr_17).(_dafny.Int)
			if ((_dafny.IntOfInt64(620)).Cmp(_68_v0) <= 0) && ((_68_v0).Cmp(_dafny.IntOfInt64(233)) < 0) {
				_coll17.Add((_68_v0).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-232))), !(true))
			}
		}
		return _coll17.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.SeqOf(!(!(false)), false, true))).Intersection((_dafny.MultiSetOf(!(false))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(false))))
}
func (_static *CompanionStruct_Default___) Fm53(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if (Companion_D23_.Create_DC61_(_dafny.SeqOf(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('b'))).Cardinality(), false)).Dtor_cf88() {
			return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("yuclgsvm"), _dafny.UnicodeSeqOfUtf8Bytes("l"))
		}
		return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ndjpjd"), _dafny.UnicodeSeqOfUtf8Bytes("n"), _dafny.UnicodeSeqOfUtf8Bytes("jgu"), _dafny.UnicodeSeqOfUtf8Bytes("p"))
	})()).Intersection(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(986))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_69_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})), _dafny.UnicodeSeqOfUtf8Bytes("clwvk")))
}
func (_static *CompanionStruct_Default___) Fm54(globalState *GlobalState) _dafny.Map {
	return (Companion_D13_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_70_i0 _dafny.Int) _dafny.CodePoint {
		return (Companion_D3_.Create_DC5_(_dafny.CodePoint('n'))).Dtor_cf5()
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))).Dtor_cf32()
}
func (_static *CompanionStruct_Default___) Fm55(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("fvrqpnj"), _dafny.UnicodeSeqOfUtf8Bytes("db"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_71_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), _dafny.UnicodeSeqOfUtf8Bytes("hs"))).Difference((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sfwcvbt"))).Intersection(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("dntsofjge"), _dafny.UnicodeSeqOfUtf8Bytes("ctwuqpd"))))
}
func (_static *CompanionStruct_Default___) Fm56(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(Companion_D6_.Create_DC13_(_dafny.IntOfInt64(613), _dafny.IntOfInt64(12)))).Difference(_dafny.MultiSetOf(Companion_D6_.Create_DC13_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("itggvjgu")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(36))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}(func(_72_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()))))).Union(_dafny.MultiSetOf(Companion_D6_.Create_DC13_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wmqdtf")).Cardinality()), _dafny.IntOfInt64(-359)), Companion_D6_.Create_DC13_((_dafny.Zero).Minus((_dafny.SetOf(true, false, true)).Cardinality()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(false))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm57(globalState *GlobalState) D16 {
	if (_dafny.MultiSetOf(_dafny.IntOfInt64(564))).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(519), (_dafny.SetOf(_dafny.IntOfInt64(43))).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(934), _dafny.IntOfInt64(879))).Cardinality())) {
		return Companion_D16_.Create_DC35_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(-867))).Cardinality(), (Companion_D22_.Create_DC57_(_dafny.IntOfInt64(-174), (_dafny.SetOf(false)).Cardinality(), false)).Dtor_cf81()), _dafny.SeqOf(true)))
	} else {
		return Companion_D16_.Create_DC35_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(553), true)).Cardinality()), _dafny.SeqOf(false, false)))
	}
}
func (_static *CompanionStruct_Default___) Fm58(p0 _dafny.Int, globalState *GlobalState) D12 {
	return Companion_D12_.Create_DC24_(Companion_D12_.Create_DC22_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm59(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-859))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}(func(_73_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mjaaqd"), _dafny.UnicodeSeqOfUtf8Bytes("bojv"))
	}))
}
func (_static *CompanionStruct_Default___) Fm60(globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll18 = _dafny.NewMapBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wodnahax")).Cardinality()), _dafny.CodePoint('s')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
			var _coll19 = _dafny.NewBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(805), _dafny.IntOfInt64(429))); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _74_v1 _dafny.Int
				_74_v1 = interface{}(_compr_19).(_dafny.Int)
				if ((_dafny.IntOfInt64(805)).Cmp(_74_v1) <= 0) && ((_74_v1).Cmp(_dafny.IntOfInt64(429)) < 0) {
					_coll19.Add((_74_v1).Plus(_dafny.IntOfInt64(-564)))
				}
			}
			return _coll19.ToSet()
		}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-498), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())).Cardinality())).Cardinality(), _dafny.CodePoint('q')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(131), _dafny.IntOfInt64(105))).Cardinality()), _dafny.CodePoint('f')), func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(754), _dafny.IntOfInt64(47))); ; {
				_compr_20, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _75_v2 _dafny.Int
				_75_v2 = interface{}(_compr_20).(_dafny.Int)
				if ((_dafny.IntOfInt64(754)).Cmp(_75_v2) <= 0) && ((_75_v2).Cmp(_dafny.IntOfInt64(47)) < 0) {
					_coll20.Add((_75_v2).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())), _dafny.CodePoint('d'))
				}
			}
			return _coll20.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-540), _dafny.CodePoint('o')))).Elements()); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _76_v0 _dafny.Map
			_76_v0 = interface{}(_compr_18).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wodnahax")).Cardinality()), _dafny.CodePoint('s')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				var _coll21 = _dafny.NewBuilder()
				_ = _coll21
				for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(805), _dafny.IntOfInt64(429))); ; {
					_compr_21, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _77_v1 _dafny.Int
					_77_v1 = interface{}(_compr_21).(_dafny.Int)
					if ((_dafny.IntOfInt64(805)).Cmp(_77_v1) <= 0) && ((_77_v1).Cmp(_dafny.IntOfInt64(429)) < 0) {
						_coll21.Add((_77_v1).Plus(_dafny.IntOfInt64(-564)))
					}
				}
				return _coll21.ToSet()
			}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-498), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())).Cardinality())).Cardinality(), _dafny.CodePoint('q')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(131), _dafny.IntOfInt64(105))).Cardinality()), _dafny.CodePoint('f')), func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(754), _dafny.IntOfInt64(47))); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _75_v2 _dafny.Int
					_75_v2 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.IntOfInt64(754)).Cmp(_75_v2) <= 0) && ((_75_v2).Cmp(_dafny.IntOfInt64(47)) < 0) {
						_coll22.Add((_75_v2).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())), _dafny.CodePoint('d'))
					}
				}
				return _coll22.ToMap()
			}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-540), _dafny.CodePoint('o'))), _76_v0) {
				_coll18.Add(_76_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)))
			}
		}
		return _coll18.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, !(true)))).Cardinality(), _dafny.CodePoint('a')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm61(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(455))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}(func(_78_i0 _dafny.Int) _dafny.Map {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))
	})))
}
func (_static *CompanionStruct_Default___) Fm62(p0 _dafny.Int, globalState *GlobalState) D18 {
	if true {
		return Companion_D18_.Create_DC43_()
	} else {
		return Companion_D18_.Create_DC43_()
	}
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) {
	var _79_v0 _dafny.Sequence
	_ = _79_v0
	_79_v0 = _dafny.SeqOf(p0)
	var _80_v1 _dafny.Sequence
	_ = _80_v1
	_80_v1 = _dafny.SeqOf(_79_v0)
	var _81_i0 _dafny.Int
	_ = _81_i0
	_81_i0 = _dafny.Zero
	{
		for _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate((_80_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_80_v1).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(p0, p0)), _79_v0) {
			{
				if (_81_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_81_i0 = (_81_i0).Plus(_dafny.One)
				var _82_v2 _dafny.CodePoint
				_ = _82_v2
				_82_v2 = _dafny.CodePoint('f')
				var _83_v3 D3
				_ = _83_v3
				_83_v3 = Companion_D3_.Create_DC5_(_82_v2)
				var _84_v4 _dafny.Map
				_ = _84_v4
				_84_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _83_v3)
				_84_v4 = (_84_v4).Update(p3, _83_v3)
				var _85_v5 _dafny.Array
				_ = _85_v5
				var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
				_ = _nw0
				_85_v5 = _nw0
				var _86_v6 _dafny.Array
				_ = _86_v6
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_0
				var _nw1 _dafny.Array
				_ = _nw1
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw1 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) _dafny.Sequence = func(_87_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(261))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg30 _dafny.Int) interface{} {
								return coer30(arg30)
							}
						}(func(_88_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.UnicodeSeqOfUtf8Bytes("rphheln")
						}))
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
				_86_v6 = _nw1
				var _89_v7 _dafny.Sequence
				_ = _89_v7
				_89_v7 = _dafny.UnicodeSeqOfUtf8Bytes("hne")
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_86_v6), 0))
				_ = _index0
				(_86_v6).ArraySet1(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("lgtdubw"), _89_v7, _89_v7), (_index0).Int())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_85_v5), 0))
				_ = _index1
				(_85_v5).ArraySet1((_79_v0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_79_v0).Cardinality()))).Uint32()).(bool), (_index1).Int())
				var _90_v8 _dafny.Sequence
				_ = _90_v8
				_90_v8 = _dafny.SeqOf(_89_v7)
				var _91_v9 _dafny.MultiSet
				_ = _91_v9
				_91_v9 = _dafny.MultiSetOf(!(p0), p0)
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_86_v6), 0))
				_ = _index2
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_85_v5), 0))
				_ = _index3
				var _rhs0 _dafny.Array = _85_v5
				_ = _rhs0
				var _rhs1 _dafny.Int = p1
				_ = _rhs1
				var _rhs2 _dafny.Int = (p1).Times((p1).Plus(_dafny.IntOfUint32((_79_v0).Cardinality())))
				_ = _rhs2
				var _rhs3 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_89_v7), _90_v8), (func() _dafny.Sequence {
					if p0 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}((func(_92_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_93_i3 _dafny.Int) _dafny.Sequence {
								return _92_v7
							}
						})(_89_v7)))
					}
					return _dafny.Companion_Sequence_.Update(_90_v8, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_90_v8).Cardinality()))).Uint32(), _89_v7)
				})()), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_89_v7), _90_v8), (func() _dafny.Sequence {
					if p0 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg32 _dafny.Int) interface{} {
								return coer32(arg32)
							}
						}((func(_94_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_95_i3 _dafny.Int) _dafny.Sequence {
								return _94_v7
							}
						})(_89_v7)))
					}
					return _dafny.Companion_Sequence_.Update(_90_v8, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_90_v8).Cardinality()))).Uint32(), _89_v7)
				})())).Cardinality()))).Uint32(), (func() _dafny.Sequence {
					if p0 {
						return _89_v7
					}
					return _89_v7
				})())
				_ = _rhs3
				var _rhs4 bool = ((_dafny.MultiSetOf(p0)).Difference(_91_v9)).IsSubsetOf(_91_v9)
				_ = _rhs4
				var _lhs0 *GlobalState = globalState
				_ = _lhs0
				var _lhs1 *GlobalState = globalState
				_ = _lhs1
				var _lhs2 _dafny.Array = _86_v6
				_ = _lhs2
				var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_86_v6), 0))
				_ = _lhs3
				var _lhs4 _dafny.Array = _85_v5
				_ = _lhs4
				var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_85_v5), 0))
				_ = _lhs5
				_85_v5 = _rhs0
				_lhs0.F7 = _rhs1
				_lhs1.F7 = _rhs2
				(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
				(_lhs4).ArraySet1(_rhs4, (_lhs5).Int())
				(globalState).F13 = ((_85_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_85_v5), 0))).Int()).(bool)) == ((_79_v0).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_79_v0).Cardinality()))).Uint32()).(bool))
				var _96_v10 _dafny.Sequence
				_ = _96_v10
				_96_v10 = _dafny.SeqOf(p1)
				(globalState).F16 = (_96_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.IntOfUint32((_96_v10).Cardinality()))).Uint32()).(_dafny.Int)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	(globalState).F16 = p1
	var _97_v11 _dafny.Array
	_ = _97_v11
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
	_ = _nw2
	_97_v11 = _nw2
	var _98_v13 _dafny.Sequence
	_ = _98_v13
	_98_v13 = _dafny.SeqOf(Companion_Default___.Fm1(_dafny.IntOfInt64(712), globalState))
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_97_v11), 0))
	_ = _index4
	(_97_v11).ArraySet1(func() _dafny.Map {
		var _coll23 = _dafny.NewMapBuilder()
		_ = _coll23
		for _iter23 := _dafny.Iterate((_98_v13).Elements()); ; {
			_compr_23, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _99_v12 _dafny.Sequence
			_99_v12 = interface{}(_compr_23).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_98_v13, _99_v12) {
				_coll23.Add(_99_v12, p1)
			}
		}
		return _coll23.ToMap()
	}(), (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_97_v11), 0))
	_ = _index5
	(_97_v11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg33 _dafny.Int) interface{} {
			return coer33(arg33)
		}
	}(func(_100_i4 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	})), p1), (_index5).Int())
	var _pat_let_tv0 = p1
	_ = _pat_let_tv0
	var _pat_let_tv1 = p0
	_ = _pat_let_tv1
	var _pat_let_tv2 = p0
	_ = _pat_let_tv2
	var _source3 D1 = func(_source4 D0) D1 {
		if _source4.Is_DC0() {
			var _101___mcc_h2 bool = _source4.Get_().(D0_DC0).Cf0
			_ = _101___mcc_h2
			var _102_cf0 bool = _101___mcc_h2
			_ = _102_cf0
			return Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv0, _pat_let_tv1))
		} else {
			var _103___mcc_h3 D0 = _source4.Get_().(D0_DC1).Cf1
			_ = _103___mcc_h3
			var _104_cf1 D0 = _103___mcc_h3
			_ = _104_cf1
			return Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(845), _pat_let_tv2))
		}
	}(Companion_D0_.Create_DC0_(p0))
	_ = _source3
	if _source3.Is_DC3() {
		var _105___mcc_h0 bool = _source3.Get_().(D1_DC3).Cf3
		_ = _105___mcc_h0
		var _106_cf3 bool = _105___mcc_h0
		_ = _106_cf3
		var _107_v14 _dafny.Array
		_ = _107_v14
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw3
		_107_v14 = _nw3
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
		_ = _index6
		(_107_v14).ArraySet1(_106_cf3, (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
		_ = _index7
		(_107_v14).ArraySet1(_106_cf3, (_index7).Int())
		if (_107_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))).Int()).(bool) {
			var _108_v15 _dafny.Array
			_ = _108_v15
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
			_ = _nw4
			_108_v15 = _nw4
			var _109_v16 _dafny.CodePoint
			_ = _109_v16
			_109_v16 = _dafny.CodePoint('c')
			var _110_v17 _dafny.Sequence
			_ = _110_v17
			_110_v17 = _dafny.UnicodeSeqOfUtf8Bytes("tmcphea")
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_108_v15), 0))
			_ = _index8
			(_108_v15).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(320))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_111_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(320))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_112_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			}))).Cardinality()))).Uint32(), _109_v16), _110_v17), (_index8).Int())
			var _113_v18 _dafny.Set
			_ = _113_v18
			_113_v18 = _dafny.SetOf(p0)
			var _114_v19 _dafny.Map
			_ = _114_v19
			_114_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_110_v17).Cardinality()), (_113_v18).Cardinality())
			var _115_v20 _dafny.Map
			_ = _115_v20
			_115_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
			var _116_v21 _dafny.Set
			_ = _116_v21
			_116_v21 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_115_v20, globalState), p3)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3))
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_108_v15), 0))
			_ = _index9
			(_108_v15).ArraySet1((func() _dafny.Sequence {
				if (_dafny.SetOf(_114_v19)).Equals(_116_v21) {
					return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("on"), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("on")).Cardinality()))).Uint32(), _109_v16)
				}
				return _110_v17
			})(), (_index9).Int())
			_115_v20 = (_115_v20).Update(_106_cf3, _106_cf3)
			(globalState).F20 = (func() _dafny.Sequence {
				if (Companion_Default___.Fm0(p1, p3, (_115_v20).Cardinality(), (_107_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))).Int()).(bool), globalState)).Cmp(p1) < 0 {
					return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yootrwdif")).Cardinality()), globalState), (_108_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_108_v15), 0))).Int()).(_dafny.Sequence))
				}
				return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_110_v17, _dafny.UnicodeSeqOfUtf8Bytes("r")), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_110_v17, _dafny.UnicodeSeqOfUtf8Bytes("r"))).Cardinality()))).Uint32(), _109_v16)
			})()
			(globalState).F16 = (_dafny.Zero).Minus(p3)
			(globalState).F2 = _106_cf3
		} else {
			_106_cf3 = (func() bool {
				if p0 {
					return _106_cf3
				}
				return false
			})()
			(globalState).F0 = p0
			(globalState).F2 = !(false)
			var _117_v22 _dafny.Map
			_ = _117_v22
			_117_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(418)).Times((_dafny.Zero).Minus(p3)), p3)
			var _118_v23 _dafny.Array
			_ = _118_v23
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_1
			var _nw5 _dafny.Array
			_ = _nw5
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw5 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) D0 = (func(_119_p0 bool) func(_dafny.Int) D0 {
					return func(_120_i6 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_119_p0)
					}
				})(p0)
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
			_118_v23 = _nw5
			var _121_v24 _dafny.Map
			_ = _121_v24
			_121_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), !(p0))
			var _122_v25 D0
			_ = _122_v25
			_122_v25 = Companion_D0_.Create_DC0_((func() bool {
				if (_121_v24).Contains(_dafny.IntOfInt64(147)) {
					return (_121_v24).Get(_dafny.IntOfInt64(147)).(bool)
				}
				return p0
			})())
			var _pat_let_tv3 = _106_cf3
			_ = _pat_let_tv3
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_118_v23), 0))
			_ = _index10
			(_118_v23).ArraySet1(func(_pat_let0_0 D0) D0 {
				return func(_123_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let1_0 bool) D0 {
						return func(_124_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_124_dt__update_hcf0_h0)
						}(_pat_let1_0)
					}(_pat_let_tv3)
				}(_pat_let0_0)
			}(_122_v25), (_index10).Int())
			var _125_v27 _dafny.Array
			_ = _125_v27
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_2
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_126_v24 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_127_i7 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_127_i7, (func() _dafny.Set {
							var _coll24 = _dafny.NewBuilder()
							_ = _coll24
							for _iter24 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _126_v24)).Keys().Elements()); ; {
								_compr_24, _ok24 := _iter24()
								if !_ok24 {
									break
								}
								var _128_v26 _dafny.CodePoint
								_128_v26 = interface{}(_compr_24).(_dafny.CodePoint)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _126_v24)).Contains(_128_v26) {
									_coll24.Add(_128_v26)
								}
							}
							return _coll24.ToSet()
						}()).Cardinality())
					}
				})(_121_v24)
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
			_125_v27 = _nw6
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_125_v27), 0))
			_ = _index11
			(_125_v27).ArraySet1((_dafny.Zero).Minus(p1), (_index11).Int())
			var _129_v28 _dafny.Sequence
			_ = _129_v28
			_129_v28 = _dafny.UnicodeSeqOfUtf8Bytes("ylg")
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_118_v23), 0))
			_ = _index12
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
			_ = _index13
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_125_v27), 0))
			_ = _index14
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
			_ = _index15
			var _rhs5 _dafny.Map = (_117_v22).Update((p3).Minus(_dafny.IntOfInt64(-542)), p3)
			_ = _rhs5
			var _rhs6 D0 = _122_v25
			_ = _rhs6
			var _rhs7 bool = p0
			_ = _rhs7
			var _rhs8 _dafny.Int = (p1).Minus(p3)
			_ = _rhs8
			var _rhs9 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-830))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_130_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), _129_v28)
			_ = _rhs9
			var _lhs6 _dafny.Array = _118_v23
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_118_v23), 0))
			_ = _lhs7
			var _lhs8 _dafny.Array = _107_v14
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
			_ = _lhs9
			var _lhs10 _dafny.Array = _125_v27
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_125_v27), 0))
			_ = _lhs11
			var _lhs12 _dafny.Array = _107_v14
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
			_ = _lhs13
			_117_v22 = _rhs5
			(_lhs6).ArraySet1(_rhs6, (_lhs7).Int())
			(_lhs8).ArraySet1(_rhs7, (_lhs9).Int())
			(_lhs10).ArraySet1(_rhs8, (_lhs11).Int())
			(_lhs12).ArraySet1(_rhs9, (_lhs13).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_125_v27), 0))
			_ = _index17
			var _rhs10 bool = _106_cf3
			_ = _rhs10
			var _rhs11 _dafny.Int = (_125_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_125_v27), 0))).Int()).(_dafny.Int)
			_ = _rhs11
			var _rhs12 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p3), p1)
			_ = _rhs12
			var _rhs13 _dafny.Array = _125_v27
			_ = _rhs13
			var _lhs14 _dafny.Array = _107_v14
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
			_ = _lhs15
			var _lhs16 _dafny.Array = _125_v27
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_125_v27), 0))
			_ = _lhs17
			var _lhs18 *GlobalState = globalState
			_ = _lhs18
			(_lhs14).ArraySet1(_rhs10, (_lhs15).Int())
			(_lhs16).ArraySet1(_rhs11, (_lhs17).Int())
			_lhs18.F7 = _rhs12
			_125_v27 = _rhs13
		}
		var _131_v29 _dafny.Array
		_ = _131_v29
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw7
		_131_v29 = _nw7
		var _132_v30 _dafny.Sequence
		_ = _132_v30
		_132_v30 = _dafny.UnicodeSeqOfUtf8Bytes("x")
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_131_v29), 0))
		_ = _index18
		(_131_v29).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(p1, globalState), _132_v30), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_131_v29), 0))
		_ = _index19
		(_131_v29).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(90))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}(func(_133_i9 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})), _132_v30), (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_107_v14), 0))
		_ = _index20
		(_107_v14).ArraySet1(true, (_index20).Int())
	} else {
		var _134___mcc_h1 _dafny.Map = _source3.Get_().(D1_DC2).Cf2
		_ = _134___mcc_h1
		var _135_cf2 _dafny.Map = _134___mcc_h1
		_ = _135_cf2
		(globalState).F16 = p3
		var _136_v31 _dafny.CodePoint
		_ = _136_v31
		_136_v31 = _dafny.CodePoint('q')
		var _137_v32 _dafny.Map
		_ = _137_v32
		_137_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p3).Plus(p1), _136_v31)
		_136_v31 = (func() _dafny.CodePoint {
			if (_137_v32).Contains(p3) {
				return (_137_v32).Get(p3).(_dafny.CodePoint)
			}
			return _136_v31
		})()
		var _138_v33 _dafny.Map
		_ = _138_v33
		_138_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _139_v34 T0
		_ = _139_v34
		var _nw8 *C3 = New_C3_()
		_ = _nw8
		_nw8.Ctor__()
		_139_v34 = _nw8
		var _140_v35 _dafny.Array
		_ = _140_v35
		var _nwElement0_0 T0 = _139_v34
		_ = _nwElement0_0
		var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(18))
		_ = _nw9
		(_nw9).ArraySet1(_nwElement0_0, 0)
		(_nw9).ArraySet1(_139_v34, 1)
		(_nw9).ArraySet1(_139_v34, 2)
		(_nw9).ArraySet1(_139_v34, 3)
		(_nw9).ArraySet1(_139_v34, 4)
		(_nw9).ArraySet1(_139_v34, 5)
		(_nw9).ArraySet1(_139_v34, 6)
		(_nw9).ArraySet1(_139_v34, 7)
		(_nw9).ArraySet1(_139_v34, 8)
		(_nw9).ArraySet1(_139_v34, 9)
		(_nw9).ArraySet1(_139_v34, 10)
		(_nw9).ArraySet1(_139_v34, 11)
		(_nw9).ArraySet1(_139_v34, 12)
		(_nw9).ArraySet1(_139_v34, 13)
		(_nw9).ArraySet1(_139_v34, 14)
		(_nw9).ArraySet1(_139_v34, 15)
		(_nw9).ArraySet1(_139_v34, 16)
		(_nw9).ArraySet1(_139_v34, 17)
		_140_v35 = _nw9
		var _141_v36 _dafny.Map
		_ = _141_v36
		_141_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v35, p0)
		var _142_v37 _dafny.Map
		_ = _142_v37
		_142_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_138_v33, globalState), Companion_D22_.Create_DC56_(_141_v36))
		var _143_v38 _dafny.Sequence
		_ = _143_v38
		_143_v38 = _dafny.SeqOf(_142_v37)
		var _144_v39 _dafny.Sequence
		_ = _144_v39
		_144_v39 = _dafny.UnicodeSeqOfUtf8Bytes("frletcac")
		var _145_v40 *C4
		_ = _145_v40
		var _nw10 *C4 = New_C4_()
		_ = _nw10
		_nw10.Ctor__(!(!((_143_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_144_v39).Cardinality()), _dafny.IntOfUint32((_143_v38).Cardinality()))).Uint32()).(_dafny.Map)).Equals(_142_v37)), p0)
		_145_v40 = _nw10
		_135_cf2 = (_135_cf2).Update((func() _dafny.Int {
			if _145_v40.F25 {
				return p3
			}
			return Companion_Default___.Fm0(_dafny.IntOfUint32((_79_v0).Cardinality()), _dafny.IntOfInt64(936), p3, _145_v40.F25, globalState)
		})(), false)
	}
	var _146_v41 *C2
	_ = _146_v41
	var _nw11 *C2 = New_C2_()
	_ = _nw11
	_nw11.Ctor__(p1)
	_146_v41 = _nw11
	var _147_v42 D18
	_ = _147_v42
	_147_v42 = Companion_D18_.Create_DC43_()
	var _148_v43 _dafny.CodePoint
	_ = _148_v43
	_148_v43 = _dafny.CodePoint('l')
	var _149_v44 _dafny.Map
	_ = _149_v44
	_149_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _148_v43)
	var _150_v45 _dafny.Map
	_ = _150_v45
	_150_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_149_v44).Cardinality())
	_147_v42 = Companion_Default___.Fm62(((_150_v45).Merge(_150_v45)).Cardinality(), globalState)
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _151_v0 _dafny.Array
	_ = _151_v0
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_3
	var _nw12 _dafny.Array
	_ = _nw12
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw12 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Int = func(_152_i0 _dafny.Int) _dafny.Int {
			return (_152_i0).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(423), (_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false), !(!(true))))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-866), _dafny.IntOfInt64(259))).Cardinality())).Cardinality())
		}
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
	_151_v0 = _nw12
	var _153_v1 _dafny.Sequence
	_ = _153_v1
	_153_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ucjarmpr")
	var _154_v2 bool
	_ = _154_v2
	_154_v2 = false
	var _155_v3 _dafny.Set
	_ = _155_v3
	_155_v3 = _dafny.SetOf(_153_v1, _153_v1)
	var _156_v5 _dafny.Int
	_ = _156_v5
	_156_v5 = _dafny.IntOfInt64(710)
	var _157_v6 _dafny.CodePoint
	_ = _157_v6
	_157_v6 = _dafny.CodePoint('t')
	var _158_v7 _dafny.Array
	_ = _158_v7
	var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
	_ = _nw13
	_158_v7 = _nw13
	var _159_globalState *GlobalState
	_ = _159_globalState
	var _nw14 *GlobalState = New_GlobalState_()
	_ = _nw14
	_nw14.Ctor__(true, _151_v0, false, _dafny.IntOfInt64(746), true, false, false, _dafny.IntOfInt64(729), _153_v1, _dafny.IntOfInt64(-908), (func() _dafny.Set {
		if _154_v2 {
			return _155_v3
		}
		return func() _dafny.Set {
			var _coll25 = _dafny.NewBuilder()
			_ = _coll25
			for _iter25 := _dafny.Iterate((_155_v3).Elements()); ; {
				_compr_25, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _160_v4 _dafny.Sequence
				_160_v4 = interface{}(_compr_25).(_dafny.Sequence)
				if (_155_v3).Contains(_160_v4) {
					_coll25.Add(_160_v4)
				}
			}
			return _coll25.ToSet()
		}()
	})(), _dafny.IntOfInt64(-466), _dafny.IntOfInt64(794), true, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_153_v1, (Companion_Default___.SafeIndex(_156_v5, _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32(), _157_v6), _153_v1), _dafny.IntOfInt64(52), _dafny.IntOfInt64(251), false, _dafny.IntOfInt64(953), _153_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cesqe"), _153_v1), false, _158_v7)
	_159_globalState = _nw14
	var _161_v8 _dafny.Sequence
	_ = _161_v8
	_161_v8 = _dafny.SeqOf(_156_v5)
	_156_v5 = ((_dafny.Zero).Minus(_dafny.IntOfInt64(-504))).Minus((_161_v8).Select((Companion_Default___.SafeIndex((_161_v8).Select((Companion_Default___.SafeIndex(_156_v5, _dafny.IntOfUint32((_161_v8).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_161_v8).Cardinality()))).Uint32()).(_dafny.Int))
	var _162_v9 _dafny.Array
	_ = _162_v9
	var _nw15 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
	_ = _nw15
	_162_v9 = _nw15
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))
	_ = _index21
	(_162_v9).ArraySet1(_154_v2, (_index21).Int())
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))
	_ = _index22
	(_162_v9).ArraySet1(_154_v2, (_index22).Int())
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))
	_ = _index23
	(_162_v9).ArraySet1((_154_v2) && ((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)), (_index23).Int())
	if _154_v2 {
		(_159_globalState).F16 = _156_v5
		var _163_v10 _dafny.Map
		_ = _163_v10
		_163_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.IntOfInt64(727), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("faepiih")).Cardinality()), _156_v5, _154_v2, _159_globalState), _154_v2)).Cardinality())
		var _164_v11 _dafny.Map
		_ = _164_v11
		_164_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v1, (_163_v10).Cardinality())
		var _165_v12 _dafny.Sequence
		_ = _165_v12
		_165_v12 = _dafny.SeqOf(_154_v2)
		var _166_v13 _dafny.Map
		_ = _166_v13
		_166_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v12, Companion_Default___.Fm1(_156_v5, _159_globalState))
		var _167_v14 _dafny.Sequence
		_ = _167_v14
		_167_v14 = _dafny.SeqOf(_165_v12, _165_v12, _165_v12)
		Companion_Default___.M0((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool), (_164_v11).Cardinality(), (_166_v13).Update((_167_v14).Select((Companion_Default___.SafeIndex(_156_v5, _dafny.IntOfUint32((_167_v14).Cardinality()))).Uint32()).(_dafny.Sequence), _153_v1), (func() _dafny.Int {
			if _154_v2 {
				return _156_v5
			}
			return _156_v5
		})(), _159_globalState)
		(_159_globalState).F7 = (_161_v8).Select((Companion_Default___.SafeIndex(_156_v5, _dafny.IntOfUint32((_161_v8).Cardinality()))).Uint32()).(_dafny.Int)
		var _168_v15 D0
		_ = _168_v15
		_168_v15 = Companion_D0_.Create_DC0_(_154_v2)
		var _169_v16 D0
		_ = _169_v16
		_169_v16 = Companion_D0_.Create_DC1_(Companion_D0_.Create_DC1_(_168_v15))
		var _170_v17 D0
		_ = _170_v17
		_170_v17 = Companion_D0_.Create_DC1_(_169_v16)
		var _source5 D0 = _170_v17
		_ = _source5
		if _source5.Is_DC0() {
			var _171___mcc_h0 bool = _source5.Get_().(D0_DC0).Cf0
			_ = _171___mcc_h0
			var _172_cf0 bool = _171___mcc_h0
			_ = _172_cf0
			var _173_v18 _dafny.Map
			_ = _173_v18
			_173_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v2, _154_v2)
			var _174_v19 _dafny.Map
			_ = _174_v19
			_174_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_cf0, (func() bool {
				if (_173_v18).Contains(_154_v2) {
					return (_173_v18).Get(_154_v2).(bool)
				}
				return _154_v2
			})())
			var _175_v20 _dafny.Map
			_ = _175_v20
			_175_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v5, (func() bool {
				if (_174_v19).Contains(_154_v2) {
					return (_174_v19).Get(_154_v2).(bool)
				}
				return _172_cf0
			})())
			var _176_v21 _dafny.Set
			_ = _176_v21
			_176_v21 = _dafny.SetOf(_156_v5, (_175_v20).Cardinality())
			Companion_Default___.M0((_176_v21).IsProperSubsetOf(_176_v21), _156_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v12, _153_v1), _156_v5, _159_globalState)
			var _177_v22 _dafny.Array
			_ = _177_v22
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_4
			var _nw16 _dafny.Array
			_ = _nw16
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw16 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Map = (func(_178_v20 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_179_i1 _dafny.Int) _dafny.Map {
						return (Companion_D1_.Create_DC2_(_178_v20)).Dtor_cf2()
					}
				})(_175_v20)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw16 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw16).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw16).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_177_v22 = _nw16
			var _180_v23 _dafny.Map
			_ = _180_v23
			_180_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v5, _156_v5)
			var _181_v24 _dafny.Sequence
			_ = _181_v24
			_181_v24 = _dafny.SeqOf(_180_v23)
			var _182_v25 _dafny.Sequence
			_ = _182_v25
			_182_v25 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_181_v24, _181_v24))
			var _183_v26 _dafny.MultiSet
			_ = _183_v26
			_183_v26 = _dafny.MultiSetOf(_151_v0)
			var _184_v27 _dafny.Map
			_ = _184_v27
			_184_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_cf0, _177_v22)
			var _rhs14 _dafny.Int = (_dafny.MultiSetFromSeq((_182_v25).Select((Companion_Default___.SafeIndex(_156_v5, _dafny.IntOfUint32((_182_v25).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()
			_ = _rhs14
			var _rhs15 bool = Companion_Default___.Fm2(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v2, (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)), _159_globalState)
			_ = _rhs15
			var _rhs16 _dafny.Array = (func() _dafny.Array {
				if (_165_v12).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_183_v26).Contains(_151_v0) {
						return (_183_v26).Multiplicity(_151_v0)
					}
					return _156_v5
				})(), _dafny.IntOfUint32((_165_v12).Cardinality()))).Uint32()).(bool) {
					return _177_v22
				}
				return (func() _dafny.Array {
					if (_184_v27).Contains(_154_v2) {
						return (_184_v27).Get(_154_v2).(_dafny.Array)
					}
					return _177_v22
				})()
			})()
			_ = _rhs16
			var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt((_156_v5).Plus(_156_v5), _156_v5)
			_ = _rhs17
			var _rhs18 bool = (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)
			_ = _rhs18
			var _lhs19 *GlobalState = _159_globalState
			_ = _lhs19
			var _lhs20 *GlobalState = _159_globalState
			_ = _lhs20
			var _lhs21 *GlobalState = _159_globalState
			_ = _lhs21
			var _lhs22 *GlobalState = _159_globalState
			_ = _lhs22
			_lhs19.F7 = _rhs14
			_lhs20.F2 = _rhs15
			_177_v22 = _rhs16
			_lhs21.F16 = _rhs17
			_lhs22.F2 = _rhs18
			var _185_v28 _dafny.MultiSet
			_ = _185_v28
			_185_v28 = _dafny.MultiSetOf(_173_v18, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool), _154_v2))
			var _186_v29 _dafny.MultiSet
			_ = _186_v29
			_186_v29 = _dafny.MultiSetOf(_156_v5, _156_v5)
			(_159_globalState).F16 = Companion_Default___.Fm0(_156_v5, (_185_v28).Cardinality(), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_156_v5, _156_v5)), (_186_v29).IsProperSubsetOf(_186_v29), _159_globalState)
			(_159_globalState).F13 = _dafny.Companion_Sequence_.IsPrefixOf(_153_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("urqdgor"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-451))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}((func(_187_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_188_i2 _dafny.Int) _dafny.CodePoint {
					return _187_v6
				}
			})(_157_v6)))))
		} else {
			var _189___mcc_h1 D0 = _source5.Get_().(D0_DC1).Cf1
			_ = _189___mcc_h1
			var _190_cf1 D0 = _189___mcc_h1
			_ = _190_cf1
			(_159_globalState).F21 = _154_v2
			_156_v5 = _dafny.IntOfInt64(439)
			var _191_v30 _dafny.Map
			_ = _191_v30
			_191_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool), (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))
			Companion_Default___.M0(Companion_Default___.Fm2(_191_v30, _159_globalState), _156_v5, _166_v13, _dafny.IntOfInt64(-198), _159_globalState)
			var _192_v31 D1
			_ = _192_v31
			_192_v31 = Companion_D1_.Create_DC3_((_154_v2) || ((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)))
			_192_v31 = _192_v31
		}
		var _193_v32 D0
		_ = _193_v32
		_193_v32 = Companion_D0_.Create_DC0_((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))
		_193_v32 = _193_v32
	} else {
		(_159_globalState).F21 = true
		_153_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bsi"), _153_v1)
		_154_v2 = false
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_151_v0), 0))
		_ = _index24
		(_151_v0).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-193)), (_index24).Int())
		var _194_v33 _dafny.Set
		_ = _194_v33
		_194_v33 = _dafny.SetOf(_154_v2, _154_v2)
		var _195_v34 _dafny.Map
		_ = _195_v34
		_195_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_194_v33).Cardinality(), (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))
		var _196_v36 D1
		_ = _196_v36
		_196_v36 = Companion_D1_.Create_DC2_((_195_v34).Merge(func() _dafny.Map {
			var _coll26 = _dafny.NewMapBuilder()
			_ = _coll26
			for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(929), _dafny.IntOfInt64(635))); ; {
				_compr_26, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _197_v35 _dafny.Int
				_197_v35 = interface{}(_compr_26).(_dafny.Int)
				if ((_dafny.IntOfInt64(929)).Cmp(_197_v35) <= 0) && ((_197_v35).Cmp(_dafny.IntOfInt64(635)) < 0) {
					_coll26.Add(Companion_Default___.SafeDivisionInt(_197_v35, _156_v5), (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))
				}
			}
			return _coll26.ToMap()
		}()))
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_151_v0), 0))
		_ = _index25
		var _rhs19 bool = (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)
		_ = _rhs19
		var _rhs20 _dafny.Int = (func() _dafny.Int {
			if (func() bool {
				if _154_v2 {
					return _154_v2
				}
				return (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)
			})() {
				return (_dafny.MultiSetOf((_dafny.Zero).Minus(_156_v5))).Cardinality()
			}
			return _156_v5
		})()
		_ = _rhs20
		var _rhs21 D1 = _196_v36
		_ = _rhs21
		var _lhs23 *GlobalState = _159_globalState
		_ = _lhs23
		var _lhs24 _dafny.Array = _151_v0
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_151_v0), 0))
		_ = _lhs25
		_lhs23.F21 = _rhs19
		(_lhs24).ArraySet1(_rhs20, (_lhs25).Int())
		_196_v36 = _rhs21
		var _198_v37 _dafny.Array
		_ = _198_v37
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_5
		var _nw17 _dafny.Array
		_ = _nw17
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw17 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Sequence = (func(_199_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_200_i3 _dafny.Int) _dafny.Sequence {
					return _199_v1
				}
			})(_153_v1)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw17 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw17).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw17).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_198_v37 = _nw17
		var _201_v38 _dafny.Map
		_ = _201_v38
		_201_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v37, _dafny.Companion_Sequence_.Concatenate(_153_v1, _153_v1))
		_201_v38 = (_201_v38).Update(_198_v37, _153_v1)
	}
	var _202_v39 D0
	_ = _202_v39
	_202_v39 = Companion_D0_.Create_DC0_(_154_v2)
	var _203_v40 _dafny.Sequence
	_ = _203_v40
	_203_v40 = _dafny.SeqOf((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool), false, _154_v2, false)
	var _204_v41 _dafny.Map
	_ = _204_v41
	_204_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_203_v40, _dafny.UnicodeSeqOfUtf8Bytes("ncojfqse"))
	Companion_Default___.M0((_202_v39).Dtor_cf0(), _156_v5, _204_v41, _156_v5, _159_globalState)
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_151_v0), 0))); ; {
		_guard_loop_0, _ok27 := _iter27()
		if !_ok27 {
			break
		}
		var _205_i4 _dafny.Int
		_205_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_205_i4).Sign() != -1) && ((_205_i4).Cmp(_dafny.ArrayLen((_151_v0), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_151_v0, (_205_i4).Int(), (_205_i4).Plus((func() _dafny.Int {
				if (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool) {
					return _156_v5
				}
				return (_dafny.Zero).Minus(_156_v5)
			})())))
		}
	}
	for _iter28 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok28 := _iter28()
		if !_ok28 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	Companion_Default___.M0((func() bool {
		if false {
			return _154_v2
		}
		return true
	})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_156_v5), _161_v8)).Cardinality()), (_204_v41).Merge(_204_v41), _156_v5, _159_globalState)
	var _206_v42 _dafny.Map
	_ = _206_v42
	_206_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v2, _156_v5)
	var _207_v43 _dafny.Map
	_ = _207_v43
	_207_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_153_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-461))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg39 _dafny.Int) interface{} {
			return coer39(arg39)
		}
	}((func(_208_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_209_i5 _dafny.Int) _dafny.CodePoint {
			return _208_v6
		}
	})(_157_v6)))), !((_206_v42).Contains((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))))
	_207_v43 = _207_v43
	(_159_globalState).F13 = _154_v2
	var _210_v44 _dafny.Map
	_ = _210_v44
	_210_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v5, (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))
	var _211_v45 D1
	_ = _211_v45
	_211_v45 = Companion_D1_.Create_DC2_(_210_v44)
	var _212_v46 _dafny.Set
	_ = _212_v46
	_212_v46 = _dafny.SetOf(_156_v5)
	_211_v45 = Companion_D1_.Create_DC2_(Companion_Default___.Fm3(_156_v5, false, _153_v1, (_212_v46).Cardinality(), _159_globalState))
	(_159_globalState).F16 = _156_v5
	var _213_v47 _dafny.MultiSet
	_ = _213_v47
	_213_v47 = _dafny.MultiSetOf(_156_v5)
	(_159_globalState).F13 = (_213_v47).IsDisjointFrom(_213_v47)
	var _source6 D0 = Companion_D0_.Create_DC1_(Companion_Default___.Fm4(Companion_Default___.Fm0(_156_v5, _dafny.IntOfInt64(593), _156_v5, _154_v2, _159_globalState), _156_v5, _159_globalState))
	_ = _source6
	if _source6.Is_DC0() {
		var _214___mcc_h2 bool = _source6.Get_().(D0_DC0).Cf0
		_ = _214___mcc_h2
		var _215_cf0 bool = _214___mcc_h2
		_ = _215_cf0
		var _pat_let_tv4 = _210_v44
		_ = _pat_let_tv4
		var _source7 D1 = func(_pat_let2_0 D1) D1 {
			return func(_216_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let3_0 _dafny.Map) D1 {
					return func(_217_dt__update_hcf2_h0 _dafny.Map) D1 {
						return Companion_D1_.Create_DC2_(_217_dt__update_hcf2_h0)
					}(_pat_let3_0)
				}(_pat_let_tv4)
			}(_pat_let2_0)
		}(_211_v45)
		_ = _source7
		if _source7.Is_DC3() {
			var _218___mcc_h4 bool = _source7.Get_().(D1_DC3).Cf3
			_ = _218___mcc_h4
			var _219_cf3 bool = _218___mcc_h4
			_ = _219_cf3
			_151_v0 = _151_v0
			_157_v6 = (func() _dafny.CodePoint {
				if _154_v2 {
					return _157_v6
				}
				return _157_v6
			})()
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw18
			_162_v9 = _nw18
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_151_v0), 0))
			_ = _index26
			(_151_v0).ArraySet1(_156_v5, (_index26).Int())
			var _220_v48 _dafny.MultiSet
			_ = _220_v48
			_220_v48 = _dafny.MultiSetOf((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_151_v0), 0))
			_ = _index27
			(_151_v0).ArraySet1(Companion_Default___.SafeDivisionInt(_156_v5, Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_156_v5, (_220_v48).Cardinality(), _156_v5, (func() bool {
				if (_210_v44).Contains(_dafny.IntOfUint32((_153_v1).Cardinality())) {
					return (_210_v44).Get(_dafny.IntOfUint32((_153_v1).Cardinality())).(bool)
				}
				return (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)
			})(), _159_globalState), (_dafny.Zero).Minus(_156_v5))), (_index27).Int())
		} else {
			var _221___mcc_h5 _dafny.Map = _source7.Get_().(D1_DC2).Cf2
			_ = _221___mcc_h5
			var _222_cf2 _dafny.Map = _221___mcc_h5
			_ = _222_cf2
			_203_v40 = _dafny.Companion_Sequence_.Concatenate(_203_v40, _dafny.SeqOf((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)))
			_161_v8 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_156_v5, _156_v5), _161_v8)
			(_159_globalState).F13 = !(!(_212_v46).Equals(func() _dafny.Set {
				var _coll27 = _dafny.NewBuilder()
				_ = _coll27
				for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(375), _dafny.IntOfInt64(820))); ; {
					_compr_27, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _223_v49 _dafny.Int
					_223_v49 = interface{}(_compr_27).(_dafny.Int)
					if ((_dafny.IntOfInt64(375)).Cmp(_223_v49) <= 0) && ((_223_v49).Cmp(_dafny.IntOfInt64(820)) < 0) {
						_coll27.Add((_223_v49).Minus(_dafny.IntOfUint32((_203_v40).Cardinality())))
					}
				}
				return _coll27.ToSet()
			}()))
			var _224_v50 _dafny.MultiSet
			_ = _224_v50
			_224_v50 = _dafny.MultiSetOf(_151_v0)
			var _225_v51 _dafny.Map
			_ = _225_v51
			_225_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_cf0, (_224_v50))
			var _226_v52 _dafny.MultiSet
			_ = _226_v52
			_226_v52 = _dafny.MultiSetOf(_151_v0)
			_225_v51 = (_225_v51).Update(_154_v2, _226_v52)
		}
		(_159_globalState).F13 = _154_v2
		_151_v0 = _151_v0
		_157_v6 = _dafny.CodePoint('e')
	} else {
		var _227___mcc_h3 D0 = _source6.Get_().(D0_DC1).Cf1
		_ = _227___mcc_h3
		var _228_cf1 D0 = _227___mcc_h3
		_ = _228_cf1
		var _229_v53 _dafny.Array
		_ = _229_v53
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
		_ = _nw19
		_229_v53 = _nw19
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_229_v53), 0))
		_ = _index28
		(_229_v53).ArraySet1(_151_v0, (_index28).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_229_v53), 0))
		_ = _index29
		(_229_v53).ArraySet1(_151_v0, (_index29).Int())
		Companion_Default___.M0((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool), _dafny.IntOfInt64(-852), _204_v41, Companion_Default___.SafeModuloInt(_156_v5, _156_v5), _159_globalState)
		(_159_globalState).F0 = (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)
		var _230_v54 _dafny.Map
		_ = _230_v54
		_230_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v6, _154_v2)
		_203_v40 = _dafny.SeqOf((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool), (func() bool {
			if (_230_v54).Contains(_157_v6) {
				return (_230_v54).Get(_157_v6).(bool)
			}
			return false
		})())
	}
	var _231_i6 _dafny.Int
	_ = _231_i6
	_231_i6 = _dafny.Zero
	{
		for (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool) {
			{
				if (_231_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_231_i6 = (_231_i6).Plus(_dafny.One)
				var _pat_let_tv5 = _210_v44
				_ = _pat_let_tv5
				_211_v45 = func(_pat_let4_0 D1) D1 {
					return func(_232_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let5_0 _dafny.Map) D1 {
							return func(_233_dt__update_hcf2_h1 _dafny.Map) D1 {
								return Companion_D1_.Create_DC2_(_233_dt__update_hcf2_h1)
							}(_pat_let5_0)
						}(_pat_let_tv5)
					}(_pat_let4_0)
				}(_211_v45)
				_211_v45 = _211_v45
				var _234_v55 _dafny.Set
				_ = _234_v55
				_234_v55 = _dafny.SetOf(_154_v2, (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool))
				_156_v5 = (_234_v55).Cardinality()
				(_159_globalState).F16 = ((_dafny.Zero).Minus(_156_v5)).Minus((_dafny.IntOfInt64(148)).Plus((_dafny.Zero).Minus(_156_v5)))
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _235_v56 _dafny.Map
	_ = _235_v56
	_235_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v9, _156_v5)
	var _236_v57 D3
	_ = _236_v57
	_236_v57 = Companion_D3_.Create_DC6_(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("yswqe"), (Companion_Default___.SafeIndex(_156_v5, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yswqe")).Cardinality()))).Uint32(), _157_v6), _157_v6, _154_v2)
	var _237_v58 _dafny.Array
	_ = _237_v58
	var _nwElement0_1 _dafny.Map = _235_v56
	_ = _nwElement0_1
	var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(21))
	_ = _nw20
	(_nw20).ArraySet1(_nwElement0_1, 0)
	(_nw20).ArraySet1(_235_v56, 1)
	(_nw20).ArraySet1((_235_v56).Update(_162_v9, _156_v5), 2)
	(_nw20).ArraySet1((_235_v56).Merge(_235_v56), 3)
	(_nw20).ArraySet1(_235_v56, 4)
	(_nw20).ArraySet1(_235_v56, 5)
	(_nw20).ArraySet1((func() _dafny.Map {
		if (_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool) {
			return _235_v56
		}
		return _235_v56
	})(), 6)
	(_nw20).ArraySet1((_235_v56).Merge((_235_v56).Update(_162_v9, (_dafny.MultiSetOf((_236_v57).Dtor_cf7())).Cardinality())), 7)
	(_nw20).ArraySet1((_235_v56).Merge(_235_v56), 8)
	(_nw20).ArraySet1(_235_v56, 9)
	(_nw20).ArraySet1(_235_v56, 10)
	(_nw20).ArraySet1(_235_v56, 11)
	(_nw20).ArraySet1(_235_v56, 12)
	(_nw20).ArraySet1(_235_v56, 13)
	(_nw20).ArraySet1((_235_v56).Update(_162_v9, _156_v5), 14)
	(_nw20).ArraySet1(_235_v56, 15)
	(_nw20).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v9, _156_v5), 16)
	(_nw20).ArraySet1((_235_v56).Merge(_235_v56), 17)
	(_nw20).ArraySet1(_235_v56, 18)
	(_nw20).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v9, _156_v5), 19)
	(_nw20).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v9, _dafny.IntOfInt64(129)), 20)
	_237_v58 = _nw20
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_237_v58), 0))
	_ = _index30
	(_237_v58).ArraySet1(_235_v56, (_index30).Int())
	var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_237_v58), 0))
	_ = _index31
	var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))
	_ = _index32
	var _rhs22 _dafny.Map = _235_v56
	_ = _rhs22
	var _rhs23 bool = !(_154_v2)
	_ = _rhs23
	var _rhs24 _dafny.Int = _156_v5
	_ = _rhs24
	var _lhs26 _dafny.Array = _237_v58
	_ = _lhs26
	var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_237_v58), 0))
	_ = _lhs27
	var _lhs28 _dafny.Array = _162_v9
	_ = _lhs28
	var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))
	_ = _lhs29
	(_lhs26).ArraySet1(_rhs22, (_lhs27).Int())
	(_lhs28).ArraySet1(_rhs23, (_lhs29).Int())
	_156_v5 = _rhs24
	_154_v2 = ((_162_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_162_v9), 0))).Int()).(bool)) && (_154_v2)
	_dafny.Print((_151_v0).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v0).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_153_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_154_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v3).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ucjarmpr"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_156_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState.F10).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ucjarmpr"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F14().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F19.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F20.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_161_v8, _dafny.SeqOf(_dafny.IntOfInt64(710))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_202_v39).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_203_v40, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_204_v41).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, false, false, false), _dafny.UnicodeSeqOfUtf8Bytes("ncojfqse"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v42).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-206))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_207_v43).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("bsiucjarmprttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt"), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_210_v44).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-206), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_211_v45).Dtor_cf2()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), false).UpdateUnsafe(_dafny.IntOfInt64(581), false).UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v46).Equals(_dafny.SetOf(_dafny.IntOfInt64(-206))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_213_v47).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-206))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_231_i6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v56).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v57).Dtor_cf6().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v57).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v57).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_237_v58).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Cardinality())
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

type D0_DC1 struct {
	Cf1 D0
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 D0) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(false)
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() D0 {
	return _this.Get_().(D0_DC1).Cf1
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
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
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
	Cf3 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 bool) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf2 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Map) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(false)
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf2() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf3 == data2.Cf3
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2.Equals(data2.Cf2)
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
	Cf4 _dafny.MultiSet
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.MultiSet) D2 {
	return D2{D2_DC4{Cf4}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D2) Dtor_cf4() _dafny.MultiSet {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf4) + ")"
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

type D3_DC6 struct {
	Cf6 _dafny.Sequence
	Cf7 _dafny.CodePoint
	Cf8 bool
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf6 _dafny.Sequence, Cf7 _dafny.CodePoint, Cf8 bool) D3 {
	return D3{D3_DC6{Cf6, Cf7, Cf8}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC5 struct {
	Cf5 _dafny.CodePoint
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf5 _dafny.CodePoint) D3 {
	return D3{D3_DC5{Cf5}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC6_(_dafny.EmptySeq, _dafny.CodePoint('D'), false)
}

func (_this D3) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D3_DC6).Cf6
}

func (_this D3) Dtor_cf7() _dafny.CodePoint {
	return _this.Get_().(D3_DC6).Cf7
}

func (_this D3) Dtor_cf8() bool {
	return _this.Get_().(D3_DC6).Cf8
}

func (_this D3) Dtor_cf5() _dafny.CodePoint {
	return _this.Get_().(D3_DC5).Cf5
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + data.Cf6.VerbatimString(true) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf5 == data2.Cf5
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

type D4_DC8 struct {
	Cf10 bool
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf10 bool) D4 {
	return D4{D4_DC8{Cf10}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC9 struct {
	Cf11 _dafny.Int
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf11 _dafny.Int) D4 {
	return D4{D4_DC9{Cf11}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC10 struct {
	Cf12 _dafny.Int
	Cf13 _dafny.Int
	Cf14 bool
	Cf15 bool
	Cf16 _dafny.Int
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf12 _dafny.Int, Cf13 _dafny.Int, Cf14 bool, Cf15 bool, Cf16 _dafny.Int) D4 {
	return D4{D4_DC10{Cf12, Cf13, Cf14, Cf15, Cf16}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC7 struct {
	Cf9 _dafny.Int
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf9 _dafny.Int) D4 {
	return D4{D4_DC7{Cf9}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC8_(false)
}

func (_this D4) Dtor_cf10() bool {
	return _this.Get_().(D4_DC8).Cf10
}

func (_this D4) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf11
}

func (_this D4) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf12
}

func (_this D4) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf13
}

func (_this D4) Dtor_cf14() bool {
	return _this.Get_().(D4_DC10).Cf14
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D4_DC7).Cf9
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf9) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf10 == data2.Cf10
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0
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

type D5_DC11 struct {
	Cf17 _dafny.Array
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf17 _dafny.Array) D5 {
	return D5{D5_DC11{Cf17}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D5_DC11).Cf17
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf17) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf17 == data2.Cf17
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

type D6_DC13 struct {
	Cf19 _dafny.Int
	Cf20 _dafny.Int
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf19 _dafny.Int, Cf20 _dafny.Int) D6 {
	return D6{D6_DC13{Cf19, Cf20}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC12 struct {
	Cf18 _dafny.Sequence
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf18 _dafny.Sequence) D6 {
	return D6{D6_DC12{Cf18}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC14 struct {
	Cf21 D6
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf21 D6) D6 {
	return D6{D6_DC14{Cf21}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC13_(_dafny.Zero, _dafny.Zero)
}

func (_this D6) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf19
}

func (_this D6) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf20
}

func (_this D6) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D6_DC12).Cf18
}

func (_this D6) Dtor_cf21() D6 {
	return _this.Get_().(D6_DC14).Cf21
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf21) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf23 _dafny.Array
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf23 _dafny.Array) D7 {
	return D7{D7_DC16{Cf23}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC15 struct {
	Cf22 _dafny.Sequence
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf22 _dafny.Sequence) D7 {
	return D7{D7_DC15{Cf22}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC17 struct {
	Cf24 D7
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf24 D7) D7 {
	return D7{D7_DC17{Cf24}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC16_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D7) Dtor_cf23() _dafny.Array {
	return _this.Get_().(D7_DC16).Cf23
}

func (_this D7) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D7_DC15).Cf22
}

func (_this D7) Dtor_cf24() D7 {
	return _this.Get_().(D7_DC17).Cf24
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf22) + ")"
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
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf22.Equals(data2.Cf22)
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

type D8_DC18 struct {
	Cf25 _dafny.Sequence
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf25 _dafny.Sequence) D8 {
	return D8{D8_DC18{Cf25}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D8_DC18).Cf25
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf26 _dafny.MultiSet
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf26 _dafny.MultiSet) D9 {
	return D9{D9_DC19{Cf26}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D9) Dtor_cf26() _dafny.MultiSet {
	return _this.Get_().(D9_DC19).Cf26
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D10_DC20 struct {
	Cf27 _dafny.Map
}

func (D10_DC20) isD10() {}

func (CompanionStruct_D10_) Create_DC20_(Cf27 _dafny.Map) D10 {
	return D10{D10_DC20{Cf27}}
}

func (_this D10) Is_DC20() bool {
	_, ok := _this.Get_().(D10_DC20)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D10) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D10_DC20).Cf27
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC20:
		{
			return "D10.DC20" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC20:
		{
			data2, ok := other.Get_().(D10_DC20)
			return ok && data1.Cf27.Equals(data2.Cf27)
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

type D11_DC21 struct {
	Cf28 _dafny.Set
}

func (D11_DC21) isD11() {}

func (CompanionStruct_D11_) Create_DC21_(Cf28 _dafny.Set) D11 {
	return D11{D11_DC21{Cf28}}
}

func (_this D11) Is_DC21() bool {
	_, ok := _this.Get_().(D11_DC21)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D11) Dtor_cf28() _dafny.Set {
	return _this.Get_().(D11_DC21).Cf28
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC21:
		{
			return "D11.DC21" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC21:
		{
			data2, ok := other.Get_().(D11_DC21)
			return ok && data1.Cf28.Equals(data2.Cf28)
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

type D12_DC23 struct {
	Cf30 _dafny.Int
}

func (D12_DC23) isD12() {}

func (CompanionStruct_D12_) Create_DC23_(Cf30 _dafny.Int) D12 {
	return D12{D12_DC23{Cf30}}
}

func (_this D12) Is_DC23() bool {
	_, ok := _this.Get_().(D12_DC23)
	return ok
}

type D12_DC22 struct {
	Cf29 _dafny.Map
}

func (D12_DC22) isD12() {}

func (CompanionStruct_D12_) Create_DC22_(Cf29 _dafny.Map) D12 {
	return D12{D12_DC22{Cf29}}
}

func (_this D12) Is_DC22() bool {
	_, ok := _this.Get_().(D12_DC22)
	return ok
}

type D12_DC24 struct {
	Cf31 D12
}

func (D12_DC24) isD12() {}

func (CompanionStruct_D12_) Create_DC24_(Cf31 D12) D12 {
	return D12{D12_DC24{Cf31}}
}

func (_this D12) Is_DC24() bool {
	_, ok := _this.Get_().(D12_DC24)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC23_(_dafny.Zero)
}

func (_this D12) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D12_DC23).Cf30
}

func (_this D12) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D12_DC22).Cf29
}

func (_this D12) Dtor_cf31() D12 {
	return _this.Get_().(D12_DC24).Cf31
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC23:
		{
			return "D12.DC23" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D12_DC22:
		{
			return "D12.DC22" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D12_DC24:
		{
			return "D12.DC24" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC23:
		{
			data2, ok := other.Get_().(D12_DC23)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D12_DC22:
		{
			data2, ok := other.Get_().(D12_DC22)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D12_DC24:
		{
			data2, ok := other.Get_().(D12_DC24)
			return ok && data1.Cf31.Equals(data2.Cf31)
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

// Definition of datatype D13
type D13 struct {
	Data_D13_
}

func (_this D13) Get_() Data_D13_ {
	return _this.Data_D13_
}

type Data_D13_ interface {
	isD13()
}

type CompanionStruct_D13_ struct {
}

var Companion_D13_ = CompanionStruct_D13_{}

type D13_DC26 struct {
	Cf33 _dafny.Sequence
	Cf34 bool
	Cf35 bool
}

func (D13_DC26) isD13() {}

func (CompanionStruct_D13_) Create_DC26_(Cf33 _dafny.Sequence, Cf34 bool, Cf35 bool) D13 {
	return D13{D13_DC26{Cf33, Cf34, Cf35}}
}

func (_this D13) Is_DC26() bool {
	_, ok := _this.Get_().(D13_DC26)
	return ok
}

type D13_DC27 struct {
	Cf36 bool
}

func (D13_DC27) isD13() {}

func (CompanionStruct_D13_) Create_DC27_(Cf36 bool) D13 {
	return D13{D13_DC27{Cf36}}
}

func (_this D13) Is_DC27() bool {
	_, ok := _this.Get_().(D13_DC27)
	return ok
}

type D13_DC25 struct {
	Cf32 _dafny.Map
}

func (D13_DC25) isD13() {}

func (CompanionStruct_D13_) Create_DC25_(Cf32 _dafny.Map) D13 {
	return D13{D13_DC25{Cf32}}
}

func (_this D13) Is_DC25() bool {
	_, ok := _this.Get_().(D13_DC25)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC26_(_dafny.EmptySeq, false, false)
}

func (_this D13) Dtor_cf33() _dafny.Sequence {
	return _this.Get_().(D13_DC26).Cf33
}

func (_this D13) Dtor_cf34() bool {
	return _this.Get_().(D13_DC26).Cf34
}

func (_this D13) Dtor_cf35() bool {
	return _this.Get_().(D13_DC26).Cf35
}

func (_this D13) Dtor_cf36() bool {
	return _this.Get_().(D13_DC27).Cf36
}

func (_this D13) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D13_DC25).Cf32
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC26:
		{
			return "D13.DC26" + "(" + data.Cf33.VerbatimString(true) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D13_DC27:
		{
			return "D13.DC27" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D13_DC25:
		{
			return "D13.DC25" + "(" + _dafny.String(data.Cf32) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC26:
		{
			data2, ok := other.Get_().(D13_DC26)
			return ok && data1.Cf33.Equals(data2.Cf33) && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35
		}
	case D13_DC27:
		{
			data2, ok := other.Get_().(D13_DC27)
			return ok && data1.Cf36 == data2.Cf36
		}
	case D13_DC25:
		{
			data2, ok := other.Get_().(D13_DC25)
			return ok && data1.Cf32.Equals(data2.Cf32)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D13) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D13)
	return ok && _this.Equals(typed)
}

func Type_D13_() _dafny.TypeDescriptor {
	return type_D13_{}
}

type type_D13_ struct {
}

func (_this type_D13_) Default() interface{} {
	return Companion_D13_.Default()
}

func (_this type_D13_) String() string {
	return "main.D13"
}
func (_this D13) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D13{}

// End of datatype D13

// Definition of datatype D14
type D14 struct {
	Data_D14_
}

func (_this D14) Get_() Data_D14_ {
	return _this.Data_D14_
}

type Data_D14_ interface {
	isD14()
}

type CompanionStruct_D14_ struct {
}

var Companion_D14_ = CompanionStruct_D14_{}

type D14_DC29 struct {
	Cf38 _dafny.Int
	Cf39 _dafny.Map
}

func (D14_DC29) isD14() {}

func (CompanionStruct_D14_) Create_DC29_(Cf38 _dafny.Int, Cf39 _dafny.Map) D14 {
	return D14{D14_DC29{Cf38, Cf39}}
}

func (_this D14) Is_DC29() bool {
	_, ok := _this.Get_().(D14_DC29)
	return ok
}

type D14_DC30 struct {
	Cf40 _dafny.Sequence
	Cf41 _dafny.Sequence
}

func (D14_DC30) isD14() {}

func (CompanionStruct_D14_) Create_DC30_(Cf40 _dafny.Sequence, Cf41 _dafny.Sequence) D14 {
	return D14{D14_DC30{Cf40, Cf41}}
}

func (_this D14) Is_DC30() bool {
	_, ok := _this.Get_().(D14_DC30)
	return ok
}

type D14_DC28 struct {
	Cf37 _dafny.Set
}

func (D14_DC28) isD14() {}

func (CompanionStruct_D14_) Create_DC28_(Cf37 _dafny.Set) D14 {
	return D14{D14_DC28{Cf37}}
}

func (_this D14) Is_DC28() bool {
	_, ok := _this.Get_().(D14_DC28)
	return ok
}

type D14_DC31 struct {
	Cf42 D14
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf42 D14) D14 {
	return D14{D14_DC31{Cf42}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC29_(_dafny.Zero, _dafny.EmptyMap)
}

func (_this D14) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D14_DC29).Cf38
}

func (_this D14) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D14_DC29).Cf39
}

func (_this D14) Dtor_cf40() _dafny.Sequence {
	return _this.Get_().(D14_DC30).Cf40
}

func (_this D14) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D14_DC30).Cf41
}

func (_this D14) Dtor_cf37() _dafny.Set {
	return _this.Get_().(D14_DC28).Cf37
}

func (_this D14) Dtor_cf42() D14 {
	return _this.Get_().(D14_DC31).Cf42
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC29:
		{
			return "D14.DC29" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D14_DC30:
		{
			return "D14.DC30" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D14_DC28:
		{
			return "D14.DC28" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC29:
		{
			data2, ok := other.Get_().(D14_DC29)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39.Equals(data2.Cf39)
		}
	case D14_DC30:
		{
			data2, ok := other.Get_().(D14_DC30)
			return ok && data1.Cf40.Equals(data2.Cf40) && data1.Cf41.Equals(data2.Cf41)
		}
	case D14_DC28:
		{
			data2, ok := other.Get_().(D14_DC28)
			return ok && data1.Cf37.Equals(data2.Cf37)
		}
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D14) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D14)
	return ok && _this.Equals(typed)
}

func Type_D14_() _dafny.TypeDescriptor {
	return type_D14_{}
}

type type_D14_ struct {
}

func (_this type_D14_) Default() interface{} {
	return Companion_D14_.Default()
}

func (_this type_D14_) String() string {
	return "main.D14"
}
func (_this D14) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D14{}

// End of datatype D14

// Definition of datatype D15
type D15 struct {
	Data_D15_
}

func (_this D15) Get_() Data_D15_ {
	return _this.Data_D15_
}

type Data_D15_ interface {
	isD15()
}

type CompanionStruct_D15_ struct {
}

var Companion_D15_ = CompanionStruct_D15_{}

type D15_DC33 struct {
	Cf44 _dafny.Sequence
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf44 _dafny.Sequence) D15 {
	return D15{D15_DC33{Cf44}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

type D15_DC32 struct {
	Cf43 _dafny.Array
}

func (D15_DC32) isD15() {}

func (CompanionStruct_D15_) Create_DC32_(Cf43 _dafny.Array) D15 {
	return D15{D15_DC32{Cf43}}
}

func (_this D15) Is_DC32() bool {
	_, ok := _this.Get_().(D15_DC32)
	return ok
}

type D15_DC34 struct {
	Cf45 D15
}

func (D15_DC34) isD15() {}

func (CompanionStruct_D15_) Create_DC34_(Cf45 D15) D15 {
	return D15{D15_DC34{Cf45}}
}

func (_this D15) Is_DC34() bool {
	_, ok := _this.Get_().(D15_DC34)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC33_(_dafny.EmptySeq)
}

func (_this D15) Dtor_cf44() _dafny.Sequence {
	return _this.Get_().(D15_DC33).Cf44
}

func (_this D15) Dtor_cf43() _dafny.Array {
	return _this.Get_().(D15_DC32).Cf43
}

func (_this D15) Dtor_cf45() D15 {
	return _this.Get_().(D15_DC34).Cf45
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC33:
		{
			return "D15.DC33" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D15_DC32:
		{
			return "D15.DC32" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D15_DC34:
		{
			return "D15.DC34" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf44.Equals(data2.Cf44)
		}
	case D15_DC32:
		{
			data2, ok := other.Get_().(D15_DC32)
			return ok && data1.Cf43 == data2.Cf43
		}
	case D15_DC34:
		{
			data2, ok := other.Get_().(D15_DC34)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D15) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D15)
	return ok && _this.Equals(typed)
}

func Type_D15_() _dafny.TypeDescriptor {
	return type_D15_{}
}

type type_D15_ struct {
}

func (_this type_D15_) Default() interface{} {
	return Companion_D15_.Default()
}

func (_this type_D15_) String() string {
	return "main.D15"
}
func (_this D15) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D15{}

// End of datatype D15

// Definition of datatype D16
type D16 struct {
	Data_D16_
}

func (_this D16) Get_() Data_D16_ {
	return _this.Data_D16_
}

type Data_D16_ interface {
	isD16()
}

type CompanionStruct_D16_ struct {
}

var Companion_D16_ = CompanionStruct_D16_{}

type D16_DC36 struct {
	Cf47 bool
	Cf48 _dafny.MultiSet
	Cf49 bool
}

func (D16_DC36) isD16() {}

func (CompanionStruct_D16_) Create_DC36_(Cf47 bool, Cf48 _dafny.MultiSet, Cf49 bool) D16 {
	return D16{D16_DC36{Cf47, Cf48, Cf49}}
}

func (_this D16) Is_DC36() bool {
	_, ok := _this.Get_().(D16_DC36)
	return ok
}

type D16_DC35 struct {
	Cf46 _dafny.Map
}

func (D16_DC35) isD16() {}

func (CompanionStruct_D16_) Create_DC35_(Cf46 _dafny.Map) D16 {
	return D16{D16_DC35{Cf46}}
}

func (_this D16) Is_DC35() bool {
	_, ok := _this.Get_().(D16_DC35)
	return ok
}

type D16_DC37 struct {
	Cf50 D16
}

func (D16_DC37) isD16() {}

func (CompanionStruct_D16_) Create_DC37_(Cf50 D16) D16 {
	return D16{D16_DC37{Cf50}}
}

func (_this D16) Is_DC37() bool {
	_, ok := _this.Get_().(D16_DC37)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC36_(false, _dafny.EmptyMultiSet, false)
}

func (_this D16) Dtor_cf47() bool {
	return _this.Get_().(D16_DC36).Cf47
}

func (_this D16) Dtor_cf48() _dafny.MultiSet {
	return _this.Get_().(D16_DC36).Cf48
}

func (_this D16) Dtor_cf49() bool {
	return _this.Get_().(D16_DC36).Cf49
}

func (_this D16) Dtor_cf46() _dafny.Map {
	return _this.Get_().(D16_DC35).Cf46
}

func (_this D16) Dtor_cf50() D16 {
	return _this.Get_().(D16_DC37).Cf50
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC36:
		{
			return "D16.DC36" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D16_DC35:
		{
			return "D16.DC35" + "(" + _dafny.String(data.Cf46) + ")"
		}
	case D16_DC37:
		{
			return "D16.DC37" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC36:
		{
			data2, ok := other.Get_().(D16_DC36)
			return ok && data1.Cf47 == data2.Cf47 && data1.Cf48.Equals(data2.Cf48) && data1.Cf49 == data2.Cf49
		}
	case D16_DC35:
		{
			data2, ok := other.Get_().(D16_DC35)
			return ok && data1.Cf46.Equals(data2.Cf46)
		}
	case D16_DC37:
		{
			data2, ok := other.Get_().(D16_DC37)
			return ok && data1.Cf50.Equals(data2.Cf50)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D16) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D16)
	return ok && _this.Equals(typed)
}

func Type_D16_() _dafny.TypeDescriptor {
	return type_D16_{}
}

type type_D16_ struct {
}

func (_this type_D16_) Default() interface{} {
	return Companion_D16_.Default()
}

func (_this type_D16_) String() string {
	return "main.D16"
}
func (_this D16) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D16{}

// End of datatype D16

// Definition of datatype D17
type D17 struct {
	Data_D17_
}

func (_this D17) Get_() Data_D17_ {
	return _this.Data_D17_
}

type Data_D17_ interface {
	isD17()
}

type CompanionStruct_D17_ struct {
}

var Companion_D17_ = CompanionStruct_D17_{}

type D17_DC39 struct {
	Cf52 _dafny.Int
	Cf53 _dafny.Map
	Cf54 bool
	Cf55 _dafny.Map
	Cf56 _dafny.Int
}

func (D17_DC39) isD17() {}

func (CompanionStruct_D17_) Create_DC39_(Cf52 _dafny.Int, Cf53 _dafny.Map, Cf54 bool, Cf55 _dafny.Map, Cf56 _dafny.Int) D17 {
	return D17{D17_DC39{Cf52, Cf53, Cf54, Cf55, Cf56}}
}

func (_this D17) Is_DC39() bool {
	_, ok := _this.Get_().(D17_DC39)
	return ok
}

type D17_DC40 struct {
}

func (D17_DC40) isD17() {}

func (CompanionStruct_D17_) Create_DC40_() D17 {
	return D17{D17_DC40{}}
}

func (_this D17) Is_DC40() bool {
	_, ok := _this.Get_().(D17_DC40)
	return ok
}

type D17_DC41 struct {
	Cf57 _dafny.Set
}

func (D17_DC41) isD17() {}

func (CompanionStruct_D17_) Create_DC41_(Cf57 _dafny.Set) D17 {
	return D17{D17_DC41{Cf57}}
}

func (_this D17) Is_DC41() bool {
	_, ok := _this.Get_().(D17_DC41)
	return ok
}

type D17_DC38 struct {
	Cf51 _dafny.Array
}

func (D17_DC38) isD17() {}

func (CompanionStruct_D17_) Create_DC38_(Cf51 _dafny.Array) D17 {
	return D17{D17_DC38{Cf51}}
}

func (_this D17) Is_DC38() bool {
	_, ok := _this.Get_().(D17_DC38)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC39_(_dafny.Zero, _dafny.EmptyMap, false, _dafny.EmptyMap, _dafny.Zero)
}

func (_this D17) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D17_DC39).Cf52
}

func (_this D17) Dtor_cf53() _dafny.Map {
	return _this.Get_().(D17_DC39).Cf53
}

func (_this D17) Dtor_cf54() bool {
	return _this.Get_().(D17_DC39).Cf54
}

func (_this D17) Dtor_cf55() _dafny.Map {
	return _this.Get_().(D17_DC39).Cf55
}

func (_this D17) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D17_DC39).Cf56
}

func (_this D17) Dtor_cf57() _dafny.Set {
	return _this.Get_().(D17_DC41).Cf57
}

func (_this D17) Dtor_cf51() _dafny.Array {
	return _this.Get_().(D17_DC38).Cf51
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC39:
		{
			return "D17.DC39" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D17_DC40:
		{
			return "D17.DC40"
		}
	case D17_DC41:
		{
			return "D17.DC41" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D17_DC38:
		{
			return "D17.DC38" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC39:
		{
			data2, ok := other.Get_().(D17_DC39)
			return ok && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53.Equals(data2.Cf53) && data1.Cf54 == data2.Cf54 && data1.Cf55.Equals(data2.Cf55) && data1.Cf56.Cmp(data2.Cf56) == 0
		}
	case D17_DC40:
		{
			_, ok := other.Get_().(D17_DC40)
			return ok
		}
	case D17_DC41:
		{
			data2, ok := other.Get_().(D17_DC41)
			return ok && data1.Cf57.Equals(data2.Cf57)
		}
	case D17_DC38:
		{
			data2, ok := other.Get_().(D17_DC38)
			return ok && data1.Cf51 == data2.Cf51
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D17) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D17)
	return ok && _this.Equals(typed)
}

func Type_D17_() _dafny.TypeDescriptor {
	return type_D17_{}
}

type type_D17_ struct {
}

func (_this type_D17_) Default() interface{} {
	return Companion_D17_.Default()
}

func (_this type_D17_) String() string {
	return "main.D17"
}
func (_this D17) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D17{}

// End of datatype D17

// Definition of datatype D18
type D18 struct {
	Data_D18_
}

func (_this D18) Get_() Data_D18_ {
	return _this.Data_D18_
}

type Data_D18_ interface {
	isD18()
}

type CompanionStruct_D18_ struct {
}

var Companion_D18_ = CompanionStruct_D18_{}

type D18_DC43 struct {
}

func (D18_DC43) isD18() {}

func (CompanionStruct_D18_) Create_DC43_() D18 {
	return D18{D18_DC43{}}
}

func (_this D18) Is_DC43() bool {
	_, ok := _this.Get_().(D18_DC43)
	return ok
}

type D18_DC44 struct {
}

func (D18_DC44) isD18() {}

func (CompanionStruct_D18_) Create_DC44_() D18 {
	return D18{D18_DC44{}}
}

func (_this D18) Is_DC44() bool {
	_, ok := _this.Get_().(D18_DC44)
	return ok
}

type D18_DC42 struct {
	Cf58 _dafny.Map
}

func (D18_DC42) isD18() {}

func (CompanionStruct_D18_) Create_DC42_(Cf58 _dafny.Map) D18 {
	return D18{D18_DC42{Cf58}}
}

func (_this D18) Is_DC42() bool {
	_, ok := _this.Get_().(D18_DC42)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC43_()
}

func (_this D18) Dtor_cf58() _dafny.Map {
	return _this.Get_().(D18_DC42).Cf58
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC43:
		{
			return "D18.DC43"
		}
	case D18_DC44:
		{
			return "D18.DC44"
		}
	case D18_DC42:
		{
			return "D18.DC42" + "(" + _dafny.String(data.Cf58) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC43:
		{
			_, ok := other.Get_().(D18_DC43)
			return ok
		}
	case D18_DC44:
		{
			_, ok := other.Get_().(D18_DC44)
			return ok
		}
	case D18_DC42:
		{
			data2, ok := other.Get_().(D18_DC42)
			return ok && data1.Cf58.Equals(data2.Cf58)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D18) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D18)
	return ok && _this.Equals(typed)
}

func Type_D18_() _dafny.TypeDescriptor {
	return type_D18_{}
}

type type_D18_ struct {
}

func (_this type_D18_) Default() interface{} {
	return Companion_D18_.Default()
}

func (_this type_D18_) String() string {
	return "main.D18"
}
func (_this D18) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D18{}

// End of datatype D18

// Definition of datatype D19
type D19 struct {
	Data_D19_
}

func (_this D19) Get_() Data_D19_ {
	return _this.Data_D19_
}

type Data_D19_ interface {
	isD19()
}

type CompanionStruct_D19_ struct {
}

var Companion_D19_ = CompanionStruct_D19_{}

type D19_DC46 struct {
	Cf60 bool
	Cf61 _dafny.Int
}

func (D19_DC46) isD19() {}

func (CompanionStruct_D19_) Create_DC46_(Cf60 bool, Cf61 _dafny.Int) D19 {
	return D19{D19_DC46{Cf60, Cf61}}
}

func (_this D19) Is_DC46() bool {
	_, ok := _this.Get_().(D19_DC46)
	return ok
}

type D19_DC47 struct {
	Cf62 D1
	Cf63 _dafny.Int
	Cf64 bool
}

func (D19_DC47) isD19() {}

func (CompanionStruct_D19_) Create_DC47_(Cf62 D1, Cf63 _dafny.Int, Cf64 bool) D19 {
	return D19{D19_DC47{Cf62, Cf63, Cf64}}
}

func (_this D19) Is_DC47() bool {
	_, ok := _this.Get_().(D19_DC47)
	return ok
}

type D19_DC48 struct {
	Cf65 _dafny.Int
	Cf66 bool
}

func (D19_DC48) isD19() {}

func (CompanionStruct_D19_) Create_DC48_(Cf65 _dafny.Int, Cf66 bool) D19 {
	return D19{D19_DC48{Cf65, Cf66}}
}

func (_this D19) Is_DC48() bool {
	_, ok := _this.Get_().(D19_DC48)
	return ok
}

type D19_DC45 struct {
	Cf59 _dafny.MultiSet
}

func (D19_DC45) isD19() {}

func (CompanionStruct_D19_) Create_DC45_(Cf59 _dafny.MultiSet) D19 {
	return D19{D19_DC45{Cf59}}
}

func (_this D19) Is_DC45() bool {
	_, ok := _this.Get_().(D19_DC45)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC46_(false, _dafny.Zero)
}

func (_this D19) Dtor_cf60() bool {
	return _this.Get_().(D19_DC46).Cf60
}

func (_this D19) Dtor_cf61() _dafny.Int {
	return _this.Get_().(D19_DC46).Cf61
}

func (_this D19) Dtor_cf62() D1 {
	return _this.Get_().(D19_DC47).Cf62
}

func (_this D19) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D19_DC47).Cf63
}

func (_this D19) Dtor_cf64() bool {
	return _this.Get_().(D19_DC47).Cf64
}

func (_this D19) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D19_DC48).Cf65
}

func (_this D19) Dtor_cf66() bool {
	return _this.Get_().(D19_DC48).Cf66
}

func (_this D19) Dtor_cf59() _dafny.MultiSet {
	return _this.Get_().(D19_DC45).Cf59
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC46:
		{
			return "D19.DC46" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D19_DC47:
		{
			return "D19.DC47" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D19_DC48:
		{
			return "D19.DC48" + "(" + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ")"
		}
	case D19_DC45:
		{
			return "D19.DC45" + "(" + _dafny.String(data.Cf59) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC46:
		{
			data2, ok := other.Get_().(D19_DC46)
			return ok && data1.Cf60 == data2.Cf60 && data1.Cf61.Cmp(data2.Cf61) == 0
		}
	case D19_DC47:
		{
			data2, ok := other.Get_().(D19_DC47)
			return ok && data1.Cf62.Equals(data2.Cf62) && data1.Cf63.Cmp(data2.Cf63) == 0 && data1.Cf64 == data2.Cf64
		}
	case D19_DC48:
		{
			data2, ok := other.Get_().(D19_DC48)
			return ok && data1.Cf65.Cmp(data2.Cf65) == 0 && data1.Cf66 == data2.Cf66
		}
	case D19_DC45:
		{
			data2, ok := other.Get_().(D19_DC45)
			return ok && data1.Cf59.Equals(data2.Cf59)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D19) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D19)
	return ok && _this.Equals(typed)
}

func Type_D19_() _dafny.TypeDescriptor {
	return type_D19_{}
}

type type_D19_ struct {
}

func (_this type_D19_) Default() interface{} {
	return Companion_D19_.Default()
}

func (_this type_D19_) String() string {
	return "main.D19"
}
func (_this D19) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D19{}

// End of datatype D19

// Definition of datatype D20
type D20 struct {
	Data_D20_
}

func (_this D20) Get_() Data_D20_ {
	return _this.Data_D20_
}

type Data_D20_ interface {
	isD20()
}

type CompanionStruct_D20_ struct {
}

var Companion_D20_ = CompanionStruct_D20_{}

type D20_DC50 struct {
	Cf68 _dafny.Int
	Cf69 bool
	Cf70 bool
}

func (D20_DC50) isD20() {}

func (CompanionStruct_D20_) Create_DC50_(Cf68 _dafny.Int, Cf69 bool, Cf70 bool) D20 {
	return D20{D20_DC50{Cf68, Cf69, Cf70}}
}

func (_this D20) Is_DC50() bool {
	_, ok := _this.Get_().(D20_DC50)
	return ok
}

type D20_DC51 struct {
}

func (D20_DC51) isD20() {}

func (CompanionStruct_D20_) Create_DC51_() D20 {
	return D20{D20_DC51{}}
}

func (_this D20) Is_DC51() bool {
	_, ok := _this.Get_().(D20_DC51)
	return ok
}

type D20_DC49 struct {
	Cf67 _dafny.Map
}

func (D20_DC49) isD20() {}

func (CompanionStruct_D20_) Create_DC49_(Cf67 _dafny.Map) D20 {
	return D20{D20_DC49{Cf67}}
}

func (_this D20) Is_DC49() bool {
	_, ok := _this.Get_().(D20_DC49)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC50_(_dafny.Zero, false, false)
}

func (_this D20) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D20_DC50).Cf68
}

func (_this D20) Dtor_cf69() bool {
	return _this.Get_().(D20_DC50).Cf69
}

func (_this D20) Dtor_cf70() bool {
	return _this.Get_().(D20_DC50).Cf70
}

func (_this D20) Dtor_cf67() _dafny.Map {
	return _this.Get_().(D20_DC49).Cf67
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC50:
		{
			return "D20.DC50" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ")"
		}
	case D20_DC51:
		{
			return "D20.DC51"
		}
	case D20_DC49:
		{
			return "D20.DC49" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC50:
		{
			data2, ok := other.Get_().(D20_DC50)
			return ok && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69 == data2.Cf69 && data1.Cf70 == data2.Cf70
		}
	case D20_DC51:
		{
			_, ok := other.Get_().(D20_DC51)
			return ok
		}
	case D20_DC49:
		{
			data2, ok := other.Get_().(D20_DC49)
			return ok && data1.Cf67.Equals(data2.Cf67)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D20) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D20)
	return ok && _this.Equals(typed)
}

func Type_D20_() _dafny.TypeDescriptor {
	return type_D20_{}
}

type type_D20_ struct {
}

func (_this type_D20_) Default() interface{} {
	return Companion_D20_.Default()
}

func (_this type_D20_) String() string {
	return "main.D20"
}
func (_this D20) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D20{}

// End of datatype D20

// Definition of datatype D21
type D21 struct {
	Data_D21_
}

func (_this D21) Get_() Data_D21_ {
	return _this.Data_D21_
}

type Data_D21_ interface {
	isD21()
}

type CompanionStruct_D21_ struct {
}

var Companion_D21_ = CompanionStruct_D21_{}

type D21_DC53 struct {
	Cf72 _dafny.Array
	Cf73 bool
	Cf74 _dafny.MultiSet
	Cf75 _dafny.Array
}

func (D21_DC53) isD21() {}

func (CompanionStruct_D21_) Create_DC53_(Cf72 _dafny.Array, Cf73 bool, Cf74 _dafny.MultiSet, Cf75 _dafny.Array) D21 {
	return D21{D21_DC53{Cf72, Cf73, Cf74, Cf75}}
}

func (_this D21) Is_DC53() bool {
	_, ok := _this.Get_().(D21_DC53)
	return ok
}

type D21_DC54 struct {
	Cf76 _dafny.Int
	Cf77 _dafny.Int
	Cf78 bool
}

func (D21_DC54) isD21() {}

func (CompanionStruct_D21_) Create_DC54_(Cf76 _dafny.Int, Cf77 _dafny.Int, Cf78 bool) D21 {
	return D21{D21_DC54{Cf76, Cf77, Cf78}}
}

func (_this D21) Is_DC54() bool {
	_, ok := _this.Get_().(D21_DC54)
	return ok
}

type D21_DC52 struct {
	Cf71 _dafny.Map
}

func (D21_DC52) isD21() {}

func (CompanionStruct_D21_) Create_DC52_(Cf71 _dafny.Map) D21 {
	return D21{D21_DC52{Cf71}}
}

func (_this D21) Is_DC52() bool {
	_, ok := _this.Get_().(D21_DC52)
	return ok
}

type D21_DC55 struct {
	Cf79 D21
}

func (D21_DC55) isD21() {}

func (CompanionStruct_D21_) Create_DC55_(Cf79 D21) D21 {
	return D21{D21_DC55{Cf79}}
}

func (_this D21) Is_DC55() bool {
	_, ok := _this.Get_().(D21_DC55)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC53_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.EmptyMultiSet, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D21) Dtor_cf72() _dafny.Array {
	return _this.Get_().(D21_DC53).Cf72
}

func (_this D21) Dtor_cf73() bool {
	return _this.Get_().(D21_DC53).Cf73
}

func (_this D21) Dtor_cf74() _dafny.MultiSet {
	return _this.Get_().(D21_DC53).Cf74
}

func (_this D21) Dtor_cf75() _dafny.Array {
	return _this.Get_().(D21_DC53).Cf75
}

func (_this D21) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D21_DC54).Cf76
}

func (_this D21) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D21_DC54).Cf77
}

func (_this D21) Dtor_cf78() bool {
	return _this.Get_().(D21_DC54).Cf78
}

func (_this D21) Dtor_cf71() _dafny.Map {
	return _this.Get_().(D21_DC52).Cf71
}

func (_this D21) Dtor_cf79() D21 {
	return _this.Get_().(D21_DC55).Cf79
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC53:
		{
			return "D21.DC53" + "(" + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ")"
		}
	case D21_DC54:
		{
			return "D21.DC54" + "(" + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ")"
		}
	case D21_DC52:
		{
			return "D21.DC52" + "(" + _dafny.String(data.Cf71) + ")"
		}
	case D21_DC55:
		{
			return "D21.DC55" + "(" + _dafny.String(data.Cf79) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC53:
		{
			data2, ok := other.Get_().(D21_DC53)
			return ok && data1.Cf72 == data2.Cf72 && data1.Cf73 == data2.Cf73 && data1.Cf74.Equals(data2.Cf74) && data1.Cf75 == data2.Cf75
		}
	case D21_DC54:
		{
			data2, ok := other.Get_().(D21_DC54)
			return ok && data1.Cf76.Cmp(data2.Cf76) == 0 && data1.Cf77.Cmp(data2.Cf77) == 0 && data1.Cf78 == data2.Cf78
		}
	case D21_DC52:
		{
			data2, ok := other.Get_().(D21_DC52)
			return ok && data1.Cf71.Equals(data2.Cf71)
		}
	case D21_DC55:
		{
			data2, ok := other.Get_().(D21_DC55)
			return ok && data1.Cf79.Equals(data2.Cf79)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D21) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D21)
	return ok && _this.Equals(typed)
}

func Type_D21_() _dafny.TypeDescriptor {
	return type_D21_{}
}

type type_D21_ struct {
}

func (_this type_D21_) Default() interface{} {
	return Companion_D21_.Default()
}

func (_this type_D21_) String() string {
	return "main.D21"
}
func (_this D21) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D21{}

// End of datatype D21

// Definition of datatype D22
type D22 struct {
	Data_D22_
}

func (_this D22) Get_() Data_D22_ {
	return _this.Data_D22_
}

type Data_D22_ interface {
	isD22()
}

type CompanionStruct_D22_ struct {
}

var Companion_D22_ = CompanionStruct_D22_{}

type D22_DC57 struct {
	Cf81 _dafny.Int
	Cf82 _dafny.Int
	Cf83 bool
}

func (D22_DC57) isD22() {}

func (CompanionStruct_D22_) Create_DC57_(Cf81 _dafny.Int, Cf82 _dafny.Int, Cf83 bool) D22 {
	return D22{D22_DC57{Cf81, Cf82, Cf83}}
}

func (_this D22) Is_DC57() bool {
	_, ok := _this.Get_().(D22_DC57)
	return ok
}

type D22_DC58 struct {
}

func (D22_DC58) isD22() {}

func (CompanionStruct_D22_) Create_DC58_() D22 {
	return D22{D22_DC58{}}
}

func (_this D22) Is_DC58() bool {
	_, ok := _this.Get_().(D22_DC58)
	return ok
}

type D22_DC56 struct {
	Cf80 _dafny.Map
}

func (D22_DC56) isD22() {}

func (CompanionStruct_D22_) Create_DC56_(Cf80 _dafny.Map) D22 {
	return D22{D22_DC56{Cf80}}
}

func (_this D22) Is_DC56() bool {
	_, ok := _this.Get_().(D22_DC56)
	return ok
}

type D22_DC59 struct {
	Cf84 D22
}

func (D22_DC59) isD22() {}

func (CompanionStruct_D22_) Create_DC59_(Cf84 D22) D22 {
	return D22{D22_DC59{Cf84}}
}

func (_this D22) Is_DC59() bool {
	_, ok := _this.Get_().(D22_DC59)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC57_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D22) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D22_DC57).Cf81
}

func (_this D22) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D22_DC57).Cf82
}

func (_this D22) Dtor_cf83() bool {
	return _this.Get_().(D22_DC57).Cf83
}

func (_this D22) Dtor_cf80() _dafny.Map {
	return _this.Get_().(D22_DC56).Cf80
}

func (_this D22) Dtor_cf84() D22 {
	return _this.Get_().(D22_DC59).Cf84
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC57:
		{
			return "D22.DC57" + "(" + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ", " + _dafny.String(data.Cf83) + ")"
		}
	case D22_DC58:
		{
			return "D22.DC58"
		}
	case D22_DC56:
		{
			return "D22.DC56" + "(" + _dafny.String(data.Cf80) + ")"
		}
	case D22_DC59:
		{
			return "D22.DC59" + "(" + _dafny.String(data.Cf84) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC57:
		{
			data2, ok := other.Get_().(D22_DC57)
			return ok && data1.Cf81.Cmp(data2.Cf81) == 0 && data1.Cf82.Cmp(data2.Cf82) == 0 && data1.Cf83 == data2.Cf83
		}
	case D22_DC58:
		{
			_, ok := other.Get_().(D22_DC58)
			return ok
		}
	case D22_DC56:
		{
			data2, ok := other.Get_().(D22_DC56)
			return ok && data1.Cf80.Equals(data2.Cf80)
		}
	case D22_DC59:
		{
			data2, ok := other.Get_().(D22_DC59)
			return ok && data1.Cf84.Equals(data2.Cf84)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D22) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D22)
	return ok && _this.Equals(typed)
}

func Type_D22_() _dafny.TypeDescriptor {
	return type_D22_{}
}

type type_D22_ struct {
}

func (_this type_D22_) Default() interface{} {
	return Companion_D22_.Default()
}

func (_this type_D22_) String() string {
	return "main.D22"
}
func (_this D22) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D22{}

// End of datatype D22

// Definition of datatype D23
type D23 struct {
	Data_D23_
}

func (_this D23) Get_() Data_D23_ {
	return _this.Data_D23_
}

type Data_D23_ interface {
	isD23()
}

type CompanionStruct_D23_ struct {
}

var Companion_D23_ = CompanionStruct_D23_{}

type D23_DC61 struct {
	Cf86 _dafny.Sequence
	Cf87 _dafny.Int
	Cf88 bool
}

func (D23_DC61) isD23() {}

func (CompanionStruct_D23_) Create_DC61_(Cf86 _dafny.Sequence, Cf87 _dafny.Int, Cf88 bool) D23 {
	return D23{D23_DC61{Cf86, Cf87, Cf88}}
}

func (_this D23) Is_DC61() bool {
	_, ok := _this.Get_().(D23_DC61)
	return ok
}

type D23_DC60 struct {
	Cf85 _dafny.Map
}

func (D23_DC60) isD23() {}

func (CompanionStruct_D23_) Create_DC60_(Cf85 _dafny.Map) D23 {
	return D23{D23_DC60{Cf85}}
}

func (_this D23) Is_DC60() bool {
	_, ok := _this.Get_().(D23_DC60)
	return ok
}

type D23_DC62 struct {
	Cf89 D23
}

func (D23_DC62) isD23() {}

func (CompanionStruct_D23_) Create_DC62_(Cf89 D23) D23 {
	return D23{D23_DC62{Cf89}}
}

func (_this D23) Is_DC62() bool {
	_, ok := _this.Get_().(D23_DC62)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC61_(_dafny.EmptySeq, _dafny.Zero, false)
}

func (_this D23) Dtor_cf86() _dafny.Sequence {
	return _this.Get_().(D23_DC61).Cf86
}

func (_this D23) Dtor_cf87() _dafny.Int {
	return _this.Get_().(D23_DC61).Cf87
}

func (_this D23) Dtor_cf88() bool {
	return _this.Get_().(D23_DC61).Cf88
}

func (_this D23) Dtor_cf85() _dafny.Map {
	return _this.Get_().(D23_DC60).Cf85
}

func (_this D23) Dtor_cf89() D23 {
	return _this.Get_().(D23_DC62).Cf89
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC61:
		{
			return "D23.DC61" + "(" + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ")"
		}
	case D23_DC60:
		{
			return "D23.DC60" + "(" + _dafny.String(data.Cf85) + ")"
		}
	case D23_DC62:
		{
			return "D23.DC62" + "(" + _dafny.String(data.Cf89) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC61:
		{
			data2, ok := other.Get_().(D23_DC61)
			return ok && data1.Cf86.Equals(data2.Cf86) && data1.Cf87.Cmp(data2.Cf87) == 0 && data1.Cf88 == data2.Cf88
		}
	case D23_DC60:
		{
			data2, ok := other.Get_().(D23_DC60)
			return ok && data1.Cf85.Equals(data2.Cf85)
		}
	case D23_DC62:
		{
			data2, ok := other.Get_().(D23_DC62)
			return ok && data1.Cf89.Equals(data2.Cf89)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D23) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D23)
	return ok && _this.Equals(typed)
}

func Type_D23_() _dafny.TypeDescriptor {
	return type_D23_{}
}

type type_D23_ struct {
}

func (_this type_D23_) Default() interface{} {
	return Companion_D23_.Default()
}

func (_this type_D23_) String() string {
	return "main.D23"
}
func (_this D23) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D23{}

// End of datatype D23

// Definition of datatype D24
type D24 struct {
	Data_D24_
}

func (_this D24) Get_() Data_D24_ {
	return _this.Data_D24_
}

type Data_D24_ interface {
	isD24()
}

type CompanionStruct_D24_ struct {
}

var Companion_D24_ = CompanionStruct_D24_{}

type D24_DC64 struct {
	Cf91 _dafny.Int
	Cf92 bool
	Cf93 bool
}

func (D24_DC64) isD24() {}

func (CompanionStruct_D24_) Create_DC64_(Cf91 _dafny.Int, Cf92 bool, Cf93 bool) D24 {
	return D24{D24_DC64{Cf91, Cf92, Cf93}}
}

func (_this D24) Is_DC64() bool {
	_, ok := _this.Get_().(D24_DC64)
	return ok
}

type D24_DC65 struct {
	Cf94 _dafny.Int
	Cf95 _dafny.Int
	Cf96 _dafny.Int
}

func (D24_DC65) isD24() {}

func (CompanionStruct_D24_) Create_DC65_(Cf94 _dafny.Int, Cf95 _dafny.Int, Cf96 _dafny.Int) D24 {
	return D24{D24_DC65{Cf94, Cf95, Cf96}}
}

func (_this D24) Is_DC65() bool {
	_, ok := _this.Get_().(D24_DC65)
	return ok
}

type D24_DC63 struct {
	Cf90 T0
}

func (D24_DC63) isD24() {}

func (CompanionStruct_D24_) Create_DC63_(Cf90 T0) D24 {
	return D24{D24_DC63{Cf90}}
}

func (_this D24) Is_DC63() bool {
	_, ok := _this.Get_().(D24_DC63)
	return ok
}

type D24_DC66 struct {
	Cf97 D24
}

func (D24_DC66) isD24() {}

func (CompanionStruct_D24_) Create_DC66_(Cf97 D24) D24 {
	return D24{D24_DC66{Cf97}}
}

func (_this D24) Is_DC66() bool {
	_, ok := _this.Get_().(D24_DC66)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC64_(_dafny.Zero, false, false)
}

func (_this D24) Dtor_cf91() _dafny.Int {
	return _this.Get_().(D24_DC64).Cf91
}

func (_this D24) Dtor_cf92() bool {
	return _this.Get_().(D24_DC64).Cf92
}

func (_this D24) Dtor_cf93() bool {
	return _this.Get_().(D24_DC64).Cf93
}

func (_this D24) Dtor_cf94() _dafny.Int {
	return _this.Get_().(D24_DC65).Cf94
}

func (_this D24) Dtor_cf95() _dafny.Int {
	return _this.Get_().(D24_DC65).Cf95
}

func (_this D24) Dtor_cf96() _dafny.Int {
	return _this.Get_().(D24_DC65).Cf96
}

func (_this D24) Dtor_cf90() T0 {
	return _this.Get_().(D24_DC63).Cf90
}

func (_this D24) Dtor_cf97() D24 {
	return _this.Get_().(D24_DC66).Cf97
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC64:
		{
			return "D24.DC64" + "(" + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ")"
		}
	case D24_DC65:
		{
			return "D24.DC65" + "(" + _dafny.String(data.Cf94) + ", " + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ")"
		}
	case D24_DC63:
		{
			return "D24.DC63" + "(" + _dafny.String(data.Cf90) + ")"
		}
	case D24_DC66:
		{
			return "D24.DC66" + "(" + _dafny.String(data.Cf97) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC64:
		{
			data2, ok := other.Get_().(D24_DC64)
			return ok && data1.Cf91.Cmp(data2.Cf91) == 0 && data1.Cf92 == data2.Cf92 && data1.Cf93 == data2.Cf93
		}
	case D24_DC65:
		{
			data2, ok := other.Get_().(D24_DC65)
			return ok && data1.Cf94.Cmp(data2.Cf94) == 0 && data1.Cf95.Cmp(data2.Cf95) == 0 && data1.Cf96.Cmp(data2.Cf96) == 0
		}
	case D24_DC63:
		{
			data2, ok := other.Get_().(D24_DC63)
			return ok && _dafny.AreEqual(data1.Cf90, data2.Cf90)
		}
	case D24_DC66:
		{
			data2, ok := other.Get_().(D24_DC66)
			return ok && data1.Cf97.Equals(data2.Cf97)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D24) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D24)
	return ok && _this.Equals(typed)
}

func Type_D24_() _dafny.TypeDescriptor {
	return type_D24_{}
}

type type_D24_ struct {
}

func (_this type_D24_) Default() interface{} {
	return Companion_D24_.Default()
}

func (_this type_D24_) String() string {
	return "main.D24"
}
func (_this D24) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D24{}

// End of datatype D24

// Definition of datatype D25
type D25 struct {
	Data_D25_
}

func (_this D25) Get_() Data_D25_ {
	return _this.Data_D25_
}

type Data_D25_ interface {
	isD25()
}

type CompanionStruct_D25_ struct {
}

var Companion_D25_ = CompanionStruct_D25_{}

type D25_DC68 struct {
}

func (D25_DC68) isD25() {}

func (CompanionStruct_D25_) Create_DC68_() D25 {
	return D25{D25_DC68{}}
}

func (_this D25) Is_DC68() bool {
	_, ok := _this.Get_().(D25_DC68)
	return ok
}

type D25_DC69 struct {
}

func (D25_DC69) isD25() {}

func (CompanionStruct_D25_) Create_DC69_() D25 {
	return D25{D25_DC69{}}
}

func (_this D25) Is_DC69() bool {
	_, ok := _this.Get_().(D25_DC69)
	return ok
}

type D25_DC67 struct {
	Cf98 _dafny.MultiSet
}

func (D25_DC67) isD25() {}

func (CompanionStruct_D25_) Create_DC67_(Cf98 _dafny.MultiSet) D25 {
	return D25{D25_DC67{Cf98}}
}

func (_this D25) Is_DC67() bool {
	_, ok := _this.Get_().(D25_DC67)
	return ok
}

type D25_DC70 struct {
	Cf99 D25
}

func (D25_DC70) isD25() {}

func (CompanionStruct_D25_) Create_DC70_(Cf99 D25) D25 {
	return D25{D25_DC70{Cf99}}
}

func (_this D25) Is_DC70() bool {
	_, ok := _this.Get_().(D25_DC70)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC68_()
}

func (_this D25) Dtor_cf98() _dafny.MultiSet {
	return _this.Get_().(D25_DC67).Cf98
}

func (_this D25) Dtor_cf99() D25 {
	return _this.Get_().(D25_DC70).Cf99
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC68:
		{
			return "D25.DC68"
		}
	case D25_DC69:
		{
			return "D25.DC69"
		}
	case D25_DC67:
		{
			return "D25.DC67" + "(" + _dafny.String(data.Cf98) + ")"
		}
	case D25_DC70:
		{
			return "D25.DC70" + "(" + _dafny.String(data.Cf99) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC68:
		{
			_, ok := other.Get_().(D25_DC68)
			return ok
		}
	case D25_DC69:
		{
			_, ok := other.Get_().(D25_DC69)
			return ok
		}
	case D25_DC67:
		{
			data2, ok := other.Get_().(D25_DC67)
			return ok && data1.Cf98.Equals(data2.Cf98)
		}
	case D25_DC70:
		{
			data2, ok := other.Get_().(D25_DC70)
			return ok && data1.Cf99.Equals(data2.Cf99)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D25) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D25)
	return ok && _this.Equals(typed)
}

func Type_D25_() _dafny.TypeDescriptor {
	return type_D25_{}
}

type type_D25_ struct {
}

func (_this type_D25_) Default() interface{} {
	return Companion_D25_.Default()
}

func (_this type_D25_) String() string {
	return "main.D25"
}
func (_this D25) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D25{}

// End of datatype D25

// Definition of datatype D26
type D26 struct {
	Data_D26_
}

func (_this D26) Get_() Data_D26_ {
	return _this.Data_D26_
}

type Data_D26_ interface {
	isD26()
}

type CompanionStruct_D26_ struct {
}

var Companion_D26_ = CompanionStruct_D26_{}

type D26_DC72 struct {
	Cf101 _dafny.Int
}

func (D26_DC72) isD26() {}

func (CompanionStruct_D26_) Create_DC72_(Cf101 _dafny.Int) D26 {
	return D26{D26_DC72{Cf101}}
}

func (_this D26) Is_DC72() bool {
	_, ok := _this.Get_().(D26_DC72)
	return ok
}

type D26_DC73 struct {
	Cf102 bool
	Cf103 bool
}

func (D26_DC73) isD26() {}

func (CompanionStruct_D26_) Create_DC73_(Cf102 bool, Cf103 bool) D26 {
	return D26{D26_DC73{Cf102, Cf103}}
}

func (_this D26) Is_DC73() bool {
	_, ok := _this.Get_().(D26_DC73)
	return ok
}

type D26_DC71 struct {
	Cf100 _dafny.MultiSet
}

func (D26_DC71) isD26() {}

func (CompanionStruct_D26_) Create_DC71_(Cf100 _dafny.MultiSet) D26 {
	return D26{D26_DC71{Cf100}}
}

func (_this D26) Is_DC71() bool {
	_, ok := _this.Get_().(D26_DC71)
	return ok
}

type D26_DC74 struct {
	Cf104 D26
}

func (D26_DC74) isD26() {}

func (CompanionStruct_D26_) Create_DC74_(Cf104 D26) D26 {
	return D26{D26_DC74{Cf104}}
}

func (_this D26) Is_DC74() bool {
	_, ok := _this.Get_().(D26_DC74)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC72_(_dafny.Zero)
}

func (_this D26) Dtor_cf101() _dafny.Int {
	return _this.Get_().(D26_DC72).Cf101
}

func (_this D26) Dtor_cf102() bool {
	return _this.Get_().(D26_DC73).Cf102
}

func (_this D26) Dtor_cf103() bool {
	return _this.Get_().(D26_DC73).Cf103
}

func (_this D26) Dtor_cf100() _dafny.MultiSet {
	return _this.Get_().(D26_DC71).Cf100
}

func (_this D26) Dtor_cf104() D26 {
	return _this.Get_().(D26_DC74).Cf104
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC72:
		{
			return "D26.DC72" + "(" + _dafny.String(data.Cf101) + ")"
		}
	case D26_DC73:
		{
			return "D26.DC73" + "(" + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ")"
		}
	case D26_DC71:
		{
			return "D26.DC71" + "(" + _dafny.String(data.Cf100) + ")"
		}
	case D26_DC74:
		{
			return "D26.DC74" + "(" + _dafny.String(data.Cf104) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC72:
		{
			data2, ok := other.Get_().(D26_DC72)
			return ok && data1.Cf101.Cmp(data2.Cf101) == 0
		}
	case D26_DC73:
		{
			data2, ok := other.Get_().(D26_DC73)
			return ok && data1.Cf102 == data2.Cf102 && data1.Cf103 == data2.Cf103
		}
	case D26_DC71:
		{
			data2, ok := other.Get_().(D26_DC71)
			return ok && data1.Cf100.Equals(data2.Cf100)
		}
	case D26_DC74:
		{
			data2, ok := other.Get_().(D26_DC74)
			return ok && data1.Cf104.Equals(data2.Cf104)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D26) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D26)
	return ok && _this.Equals(typed)
}

func Type_D26_() _dafny.TypeDescriptor {
	return type_D26_{}
}

type type_D26_ struct {
}

func (_this type_D26_) Default() interface{} {
	return Companion_D26_.Default()
}

func (_this type_D26_) String() string {
	return "main.D26"
}
func (_this D26) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D26{}

// End of datatype D26

// Definition of trait T0
type T0 interface {
	String() string
	Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map
	Fm6(p0 _dafny.Int, globalState *GlobalState) bool
	M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool)
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
	Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map
	Fm6(p0 _dafny.Int, globalState *GlobalState) bool
	M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool)
	Fm7(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool
	M2(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Array)
	M3(p0 _dafny.Array, globalState *GlobalState)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   bool
	F2   bool
	F7   _dafny.Int
	F8   _dafny.Sequence
	F10  _dafny.Set
	F13  bool
	F16  _dafny.Int
	F19  _dafny.Sequence
	F20  _dafny.Sequence
	F21  bool
	F22  _dafny.Array
	_f1  _dafny.Array
	_f3  _dafny.Int
	_f4  bool
	_f5  bool
	_f6  bool
	_f9  _dafny.Int
	_f11 _dafny.Int
	_f12 _dafny.Int
	_f14 _dafny.Sequence
	_f15 _dafny.Int
	_f17 bool
	_f18 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F2 = false
	_this.F7 = _dafny.Zero
	_this.F8 = _dafny.EmptySeq
	_this.F10 = _dafny.EmptySet
	_this.F13 = false
	_this.F16 = _dafny.Zero
	_this.F19 = _dafny.EmptySeq
	_this.F20 = _dafny.EmptySeq
	_this.F21 = false
	_this.F22 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.Zero
	_this._f4 = false
	_this._f5 = false
	_this._f6 = false
	_this._f9 = _dafny.Zero
	_this._f11 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f14 = _dafny.EmptySeq
	_this._f15 = _dafny.Zero
	_this._f17 = false
	_this._f18 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Array, f2 bool, f3 _dafny.Int, f4 bool, f5 bool, f6 bool, f7 _dafny.Int, f8 _dafny.Sequence, f9 _dafny.Int, f10 _dafny.Set, f11 _dafny.Int, f12 _dafny.Int, f13 bool, f14 _dafny.Sequence, f15 _dafny.Int, f16 _dafny.Int, f17 bool, f18 _dafny.Int, f19 _dafny.Sequence, f20 _dafny.Sequence, f21 bool, f22 _dafny.Array) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this).F20 = f20
		(_this).F21 = f21
		(_this).F22 = f22
	}
}
func (_this *GlobalState) F1() _dafny.Array {
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
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F11() _dafny.Int {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Sequence {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Int {
	{
		return _this._f18
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F27 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F27 = false
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

func (_this *C0) Ctor__(f27 bool) {
	{
		(_this).F27 = f27
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F29  _dafny.Int
	_f30 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F29 = _dafny.Zero
	_this._f30 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__(f29 _dafny.Int, f30 _dafny.Sequence) {
	{
		(_this).F29 = f29
		(_this)._f30 = f30
	}
}
func (_this *C1) Fm7(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C1) Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((false) == (!(true)), _dafny.UnicodeSeqOfUtf8Bytes("ctdvjyvbi"))
	}
}
func (_this *C1) Fm6(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this.F29).Cmp(_this.F29) == 0
	}
}
func (_this *C1) Fm29(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.MultiSetOf(false, false, true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))).Cardinality()
	}
}
func (_this *C1) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		if p1 {
			if p1 {
				(globalState).F2 = p1
				var _238_v0 *C0
				_ = _238_v0
				var _nw21 *C0 = New_C0_()
				_ = _nw21
				_nw21.Ctor__(!(true))
				_238_v0 = _nw21
				var _239_v1 *C0
				_ = _239_v1
				var _nw22 *C0 = New_C0_()
				_ = _nw22
				_nw22.Ctor__(p1)
				_239_v1 = _nw22
				var _240_v2 _dafny.CodePoint
				_ = _240_v2
				_240_v2 = _dafny.CodePoint('l')
				var _241_v3 _dafny.Sequence
				_ = _241_v3
				_241_v3 = _dafny.SeqOf(_240_v2)
				var _242_v4 _dafny.MultiSet
				_ = _242_v4
				_242_v4 = _dafny.MultiSetOf(_241_v3)
				var _243_v5 _dafny.Sequence
				_ = _243_v5
				_243_v5 = _dafny.SeqOf(Companion_Default___.SafeModuloInt(p0, (_242_v4).Cardinality()))
				var _244_v6 _dafny.Array
				_ = _244_v6
				var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
				_ = _nw23
				_244_v6 = _nw23
				var _245_v7 _dafny.Array
				_ = _245_v7
				var _nw24 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw24
				_245_v7 = _nw24
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_244_v6), 0))
				_ = _index33
				(_244_v6).ArraySet1(_245_v7, (_index33).Int())
				var _246_v8 _dafny.Map
				_ = _246_v8
				_246_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _239_v1.F27)
				var _247_v9 _dafny.Array
				_ = _247_v9
				var _nwElement0_2 _dafny.Int = (Companion_Default___.Fm0(p3, _dafny.IntOfInt64(438), _dafny.IntOfUint32((_241_v3).Cardinality()), p1, globalState)).Minus((_246_v8).Cardinality())
				_ = _nwElement0_2
				var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.One)
				_ = _nw25
				(_nw25).ArraySet1(_nwElement0_2, 0)
				_247_v9 = _nw25
				var _248_v10 _dafny.Map
				_ = _248_v10
				_248_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _238_v0.F27)
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_244_v6), 0))
				_ = _index34
				var _rhs25 _dafny.Sequence = _dafny.SeqOf(p3, (_248_v10).Cardinality(), p3)
				_ = _rhs25
				var _rhs26 _dafny.Array = _245_v7
				_ = _rhs26
				var _rhs27 _dafny.Int = p3
				_ = _rhs27
				var _rhs28 _dafny.Array = _247_v9
				_ = _rhs28
				var _lhs30 _dafny.Array = _244_v6
				_ = _lhs30
				var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_244_v6), 0))
				_ = _lhs31
				var _lhs32 *GlobalState = globalState
				_ = _lhs32
				_243_v5 = _rhs25
				(_lhs30).ArraySet1(_rhs26, (_lhs31).Int())
				_lhs32.F7 = _rhs27
				_247_v9 = _rhs28
				(globalState).F21 = _238_v0.F27
			} else {
				var _249_v11 _dafny.Array
				_ = _249_v11
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_6
				var _nw26 _dafny.Array
				_ = _nw26
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw26 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Int = func(_250_i0 _dafny.Int) _dafny.Int {
						return (_250_i0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ks")).Cardinality()))
					}
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
				_249_v11 = _nw26
				var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw27
				_249_v11 = _nw27
				(globalState).F0 = p1
				var _251_v12 _dafny.Map
				_ = _251_v12
				_251_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
				(globalState).F7 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(p0, (_251_v12).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), true)).Cardinality())
				var _252_v13 _dafny.Map
				_ = _252_v13
				_252_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _253_v14 _dafny.MultiSet
				_ = _253_v14
				_253_v14 = _dafny.MultiSetOf(_this.F29, (Companion_Default___.Fm0(p0, p3, _dafny.IntOfInt64(758), p1, globalState)).Times((_252_v13).Cardinality()), p0)
				_253_v14 = _253_v14
				var _254_v15 D3
				_ = _254_v15
				_254_v15 = Companion_D3_.Create_DC5_(_dafny.CodePoint('m'))
				var _255_v16 _dafny.MultiSet
				_ = _255_v16
				_255_v16 = _dafny.MultiSetOf(_254_v15)
				var _256_v17 _dafny.Set
				_ = _256_v17
				_256_v17 = _dafny.SetOf(p3)
				var _257_v18 _dafny.Sequence
				_ = _257_v18
				_257_v18 = _dafny.SeqOf(p1, !((Companion_Default___.Fm30(p1, _this.F29, (_dafny.Zero).Minus(_this.F29), _255_v16, globalState)).Equals(_256_v17)), p1, p1)
				_257_v18 = _257_v18
			}
			var _258_v19 _dafny.Map
			_ = _258_v19
			_258_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
			var _rhs29 _dafny.Int = ((func() _dafny.Map {
				if p1 {
					return _258_v19
				}
				return _258_v19
			})()).Cardinality()
			_ = _rhs29
			var _rhs30 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm31(p0, !(p1) || (p1), globalState)).Cardinality())
			_ = _rhs30
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			_lhs33.F16 = _rhs29
			r1 = _rhs30
			if p1 {
				(globalState).F2 = p1
				var _259_v20 _dafny.Sequence
				_ = _259_v20
				_259_v20 = _dafny.SeqOf(p1)
				var _260_v21 _dafny.Map
				_ = _260_v21
				_260_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _259_v20)
				var _261_v22 _dafny.Map
				_ = _261_v22
				_261_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_260_v21).Contains(_dafny.IntOfInt64(142)) {
						return (_260_v21).Get(_dafny.IntOfInt64(142)).(_dafny.Sequence)
					}
					return _259_v20
				})(), Companion_Default___.Fm1(_this.F29, globalState))
				Companion_Default___.M0(p1, p0, _261_v22, (_this).Fm29(globalState), globalState)
				var _262_v23 _dafny.Map
				_ = _262_v23
				_262_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_259_v20, _dafny.SeqOf(p1, p1, p1, p1, p1)), false)
				var _263_v24 _dafny.Set
				_ = _263_v24
				_263_v24 = _dafny.SetOf(Companion_Default___.Fm0(p3, p3, p3, p1, globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_259_v20, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_259_v20).Cardinality()))).Uint32(), p1)).Cardinality()), _this.F29, _this.F29)
				var _264_v25 D14
				_ = _264_v25
				_264_v25 = Companion_D14_.Create_DC28_(_263_v24)
				(globalState).F2 = (func() bool {
					if (_262_v23).Contains(_259_v20) {
						return (_262_v23).Get(_259_v20).(bool)
					}
					return ((_264_v25).Dtor_cf37()).IsProperSubsetOf(_263_v24)
				})()
				var _265_v26 _dafny.Sequence
				_ = _265_v26
				_265_v26 = _dafny.UnicodeSeqOfUtf8Bytes("ifljbrg")
				var _266_v27 _dafny.Sequence
				_ = _266_v27
				_266_v27 = _dafny.SeqOf(_this.F29, _dafny.IntOfUint32((_265_v26).Cardinality()))
				(globalState).F7 = Companion_Default___.Fm0(p0, (_266_v27).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_266_v27).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(392), p1, globalState)
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((p2), 0))
				_ = _index35
				(p2).ArraySet1(p0, (_index35).Int())
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((p2), 0))
				_ = _index36
				(p2).ArraySet1(p3, (_index36).Int())
			} else {
				var _267_v28 _dafny.Sequence
				_ = _267_v28
				_267_v28 = _dafny.UnicodeSeqOfUtf8Bytes("gtrohbu")
				var _268_v29 _dafny.Map
				_ = _268_v29
				_268_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _267_v28)
				var _269_v30 _dafny.CodePoint
				_ = _269_v30
				_269_v30 = _dafny.CodePoint('w')
				(globalState).F8 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_268_v29).Contains(p1) {
						return (_268_v29).Get(p1).(_dafny.Sequence)
					}
					return _267_v28
				})(), (Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_268_v29).Contains(p1) {
						return (_268_v29).Get(p1).(_dafny.Sequence)
					}
					return _267_v28
				})()).Cardinality()))).Uint32(), _269_v30), _267_v28)
				var _270_v31 _dafny.Array
				_ = _270_v31
				var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
				_ = _nw28
				_270_v31 = _nw28
				var _271_v32 _dafny.Map
				_ = _271_v32
				_271_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.SeqOf(_this.F29, _dafny.IntOfUint32((_267_v28).Cardinality())))
				var _272_v33 _dafny.Sequence
				_ = _272_v33
				_272_v33 = _dafny.SeqOf(p1)
				var _273_v34 D12
				_ = _273_v34
				_273_v34 = Companion_D12_.Create_DC23_(_dafny.IntOfInt64(528))
				var _274_v35 _dafny.Map
				_ = _274_v35
				_274_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _275_v36 _dafny.Sequence
				_ = _275_v36
				_275_v36 = _dafny.SeqOf(p0, (_273_v34).Dtor_cf30(), (_274_v35).Cardinality())
				var _276_v37 _dafny.Set
				_ = _276_v37
				_276_v37 = _dafny.SetOf(p0, p0, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_271_v32).Contains(_dafny.IntOfUint32((_272_v33).Cardinality())) {
						return (_271_v32).Get(_dafny.IntOfUint32((_272_v33).Cardinality())).(_dafny.Sequence)
					}
					return _dafny.SeqOf((_275_v36).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_275_v36).Cardinality()))).Uint32()).(_dafny.Int), p0)
				})()).Cardinality()))
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_270_v31), 0))
				_ = _index37
				(_270_v31).ArraySet1(_276_v37, (_index37).Int())
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_270_v31), 0))
				_ = _index38
				(_270_v31).ArraySet1(_276_v37, (_index38).Int())
				var _277_v38 _dafny.Array
				_ = _277_v38
				var _nw29 _dafny.Array = _dafny.NewArrayWithValue(Companion_D14_.Default(), _dafny.IntOfInt64(3))
				_ = _nw29
				_277_v38 = _nw29
				var _278_v39 D14
				_ = _278_v39
				_278_v39 = Companion_D14_.Create_DC28_((_270_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_270_v31), 0))).Int()).(_dafny.Set))
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_277_v38), 0))
				_ = _index39
				(_277_v38).ArraySet1(_278_v39, (_index39).Int())
				var _279_v40 _dafny.Map
				_ = _279_v40
				_279_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_274_v35, _274_v35, _274_v35, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm6(p0, globalState), true), _274_v35)).IsDisjointFrom(_dafny.SetOf((_274_v35).Update(p1, p1), _274_v35)), p2)
				var _280_v41 _dafny.MultiSet
				_ = _280_v41
				_280_v41 = _dafny.MultiSetOf(p1, (_272_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.IntOfUint32((_272_v33).Cardinality()))).Uint32()).(bool))
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_277_v38), 0))
				_ = _index40
				var _rhs31 bool = p1
				_ = _rhs31
				var _rhs32 D14 = Companion_Default___.Fm32(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p1, false), !(p1)), (_280_v41).Cardinality(), p3, globalState)
				_ = _rhs32
				var _rhs33 _dafny.Map = (_279_v40).Merge(_279_v40)
				_ = _rhs33
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				var _lhs35 _dafny.Array = _277_v38
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_277_v38), 0))
				_ = _lhs36
				_lhs34.F2 = _rhs31
				(_lhs35).ArraySet1(_rhs32, (_lhs36).Int())
				_279_v40 = _rhs33
				(globalState).F2 = p1
				var _281_v42 _dafny.Set
				_ = _281_v42
				_281_v42 = _dafny.SetOf(p1, p1, (_this).Fm7(p1, p3, p1, globalState))
				var _282_v43 _dafny.MultiSet
				_ = _282_v43
				_282_v43 = _dafny.MultiSetOf(_281_v42)
				var _283_v44 _dafny.Set
				_ = _283_v44
				_283_v44 = _281_v42
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((p2), 0))
				_ = _index41
				(p2).ArraySet1((func() _dafny.Int {
					if (_282_v43).Contains(_283_v44) {
						return (_282_v43).Multiplicity(_283_v44)
					}
					return p0
				})(), (_index41).Int())
				var _284_v45 _dafny.MultiSet
				_ = _284_v45
				_284_v45 = _dafny.MultiSetOf((p3).Plus(_this.F29), _this.F29, (_this.F29).Plus(p0))
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((p2), 0))
				_ = _index42
				(p2).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_284_v45).Contains(p3) {
						return (_284_v45).Multiplicity(p3)
					}
					return Companion_Default___.SafeDivisionInt((func() _dafny.Int {
						if (_258_v19).Contains(p1) {
							return (_258_v19).Get(p1).(_dafny.Int)
						}
						return p3
					})(), _this.F29)
				})()), (_index42).Int())
			}
			(globalState).F13 = (_dafny.IntOfUint32((_dafny.SeqOf(p0, (_258_v19).Cardinality(), _this.F29, p3, _this.F29)).Cardinality())).Cmp(p0) <= 0
			var _285_v46 _dafny.CodePoint
			_ = _285_v46
			_285_v46 = _dafny.CodePoint('t')
			var _286_v47 _dafny.Sequence
			_ = _286_v47
			_286_v47 = _dafny.SeqOf(_285_v46, _285_v46)
			r1 = Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}(func(_287_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})), _286_v47)).Cardinality()))
		} else {
			var _288_v48 _dafny.Sequence
			_ = _288_v48
			_288_v48 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			var _289_v49 _dafny.CodePoint
			_ = _289_v49
			_289_v49 = _dafny.CodePoint('i')
			var _290_v50 D3
			_ = _290_v50
			_290_v50 = Companion_D3_.Create_DC6_(_288_v48, _289_v49, p1)
			(globalState).F7 = _dafny.IntOfUint32((_dafny.SeqOf(_290_v50)).Cardinality())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((p2), 0))
			_ = _index43
			(p2).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p3)), (_index43).Int())
			var _291_v51 D4
			_ = _291_v51
			_291_v51 = Companion_D4_.Create_DC9_(p3)
			var _292_v52 _dafny.Sequence
			_ = _292_v52
			_292_v52 = _dafny.SeqOf(p1, p1, p1)
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((p2), 0))
			_ = _index44
			(p2).ArraySet1((p0).Times(((_291_v51).Dtor_cf11()).Plus(_dafny.IntOfUint32((_292_v52).Cardinality()))), (_index44).Int())
			var _293_v53 _dafny.Map
			_ = _293_v53
			_293_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
			(globalState).F21 = Companion_Default___.Fm2((_293_v53).Merge(_293_v53), globalState)
			(globalState).F7 = (_dafny.Zero).Minus((_this.F29).Minus((_dafny.Zero).Minus((p0).Times(_dafny.IntOfInt64(549)))))
			var _294_v54 _dafny.Array
			_ = _294_v54
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw30
			_294_v54 = _nw30
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_294_v54), 0))
			_ = _index45
			(_294_v54).ArraySet1(!(p1), (_index45).Int())
			var _295_v55 _dafny.Sequence
			_ = _295_v55
			_295_v55 = _dafny.SeqOf((_dafny.Zero).Minus((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)))
			var _296_v56 _dafny.Map
			_ = _296_v56
			_296_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p3)
			var _297_v57 _dafny.Map
			_ = _297_v57
			_297_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_288_v48, _dafny.MultiSetOf((_295_v55).Select((Companion_Default___.SafeIndex((_296_v56).Cardinality(), _dafny.IntOfUint32((_295_v55).Cardinality()))).Uint32()).(_dafny.Int), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)))
			var _298_v58 _dafny.MultiSet
			_ = _298_v58
			_298_v58 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_288_v48).Cardinality()))
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_294_v54), 0))
			_ = _index46
			var _rhs34 _dafny.Int = ((func() _dafny.MultiSet {
				if (_297_v57).Contains(_288_v48) {
					return (_297_v57).Get(_288_v48).(_dafny.MultiSet)
				}
				return _298_v58
			})()).Cardinality()
			_ = _rhs34
			var _rhs35 bool = p1
			_ = _rhs35
			var _lhs37 *C1 = _this
			_ = _lhs37
			var _lhs38 _dafny.Array = _294_v54
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_294_v54), 0))
			_ = _lhs39
			_lhs37.F29 = _rhs34
			(_lhs38).ArraySet1(_rhs35, (_lhs39).Int())
		}
		var _299_v59 _dafny.Map
		_ = _299_v59
		_299_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		_299_v59 = (_299_v59).Update(p2, p0)
		var _300_v60 _dafny.Sequence
		_ = _300_v60
		_300_v60 = _dafny.SeqOf(_this.F29)
		var _301_v61 *C0
		_ = _301_v61
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__(p1)
		_301_v61 = _nw31
		var _302_v62 _dafny.Sequence
		_ = _302_v62
		_302_v62 = _dafny.SeqOf(p1, p1, p1, p1)
		var _303_v63 _dafny.Map
		_ = _303_v63
		_303_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_302_v62).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_302_v62).Cardinality()), _dafny.IntOfUint32((_302_v62).Cardinality()))).Uint32()).(bool), _301_v61.F27)
		var _304_v64 _dafny.Map
		_ = _304_v64
		_304_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v61, _303_v63)
		var _305_v65 _dafny.Sequence
		_ = _305_v65
		_305_v65 = _dafny.UnicodeSeqOfUtf8Bytes("fgkj")
		var _306_v66 _dafny.Map
		_ = _306_v66
		_306_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_302_v62, _305_v65)
		Companion_Default___.M0(p1, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_300_v60).Cardinality()), Companion_Default___.Fm0(_dafny.IntOfInt64(-680), ((func() _dafny.Map {
			if (_304_v64).Contains(_301_v61) {
				return (_304_v64).Get(_301_v61).(_dafny.Map)
			}
			return _303_v63
		})()).Cardinality(), p3, _301_v61.F27, globalState)), _306_v66, _dafny.IntOfUint32((_300_v60).Cardinality()), globalState)
		var _307_v68 _dafny.CodePoint
		_ = _307_v68
		_307_v68 = _dafny.CodePoint('w')
		var _308_v69 _dafny.Map
		_ = _308_v69
		_308_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _307_v68)
		var _309_v70 _dafny.Map
		_ = _309_v70
		_309_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, (_308_v69).Cardinality())
		var _310_v72 _dafny.Map
		_ = _310_v72
		_310_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Set {
			var _coll28 = _dafny.NewBuilder()
			_ = _coll28
			for _iter30 := _dafny.Iterate((_dafny.SeqOf(_309_v70, _309_v70)).Elements()); ; {
				_compr_28, _ok30 := _iter30()
				if !_ok30 {
					break
				}
				var _311_v71 _dafny.Map
				_311_v71 = interface{}(_compr_28).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_309_v70, _309_v70), _311_v71) {
					_coll28.Add(_311_v71)
				}
			}
			return _coll28.ToSet()
		}()).Cardinality())
		if (((func() _dafny.Map {
			var _coll29 = _dafny.NewMapBuilder()
			_ = _coll29
			for _iter31 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_310_v72).Contains(p3) {
					return (_310_v72).Get(p3).(_dafny.Int)
				}
				return p0
			})(), p0)).Keys().Elements()); ; {
				_compr_29, _ok31 := _iter31()
				if !_ok31 {
					break
				}
				var _312_v67 _dafny.Int
				_312_v67 = interface{}(_compr_29).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_310_v72).Contains(p3) {
						return (_310_v72).Get(p3).(_dafny.Int)
					}
					return p0
				})(), p0)).Contains(_312_v67) {
					_coll29.Add(Companion_Default___.SafeModuloInt(_312_v67, (_300_v60).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_300_v60).Cardinality()))).Uint32()).(_dafny.Int)), p0)
				}
			}
			return _coll29.ToMap()
		}()).Update(_this.F29, p0)).Cardinality()).Cmp(p3) <= 0 {
			(_this).F29 = p3
			var _313_v73 _dafny.Array
			_ = _313_v73
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw32
			_313_v73 = _nw32
			var _314_v74 _dafny.MultiSet
			_ = _314_v74
			_314_v74 = _dafny.MultiSetOf(p3, p0)
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_313_v73), 0))
			_ = _index47
			(_313_v73).ArraySet1((_314_v74).IsProperSubsetOf(Companion_Default___.Fm33(globalState)), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_313_v73), 0))
			_ = _index48
			(_313_v73).ArraySet1(((_this).Fm29(globalState)).Cmp(((func() _dafny.Int {
				if (_309_v70).Contains(p0) {
					return (_309_v70).Get(p0).(_dafny.Int)
				}
				return p3
			})()).Plus(_this.F29)) >= 0, (_index48).Int())
			(globalState).F13 = p1
			var _315_v75 _dafny.MultiSet
			_ = _315_v75
			_315_v75 = _dafny.MultiSetOf(_301_v61.F27, (_313_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_313_v73), 0))).Int()).(bool))
			if ((_315_v75).Update(p1, Companion_Default___.Abs((_dafny.Zero).Minus(p0)))).IsSubsetOf(_315_v75) {
				r1 = _this.F29
				var _316_v76 _dafny.Array
				_ = _316_v76
				var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(4))
				_ = _nw33
				_316_v76 = _nw33
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_316_v76), 0))
				_ = _index49
				(_316_v76).ArraySet1CodePoint(_307_v68, (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_316_v76), 0))
				_ = _index50
				(_316_v76).ArraySet1CodePoint(Companion_Default___.Fm34(globalState), (_index50).Int())
				var _317_v77 _dafny.MultiSet
				_ = _317_v77
				_317_v77 = _dafny.MultiSetOf(_dafny.SeqOf(true, p1), _302_v62)
				(globalState).F0 = !((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_302_v62, _302_v62, _dafny.SeqOf(p1, p1, false), _302_v62), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_300_v60).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_302_v62, _302_v62, _dafny.SeqOf(p1, p1, false), _302_v62)).Cardinality()))).Uint32(), _302_v62))).IsProperSubsetOf(_317_v77))
				Companion_Default___.M0(!_dafny.Companion_Sequence_.Contains(_305_v65, _307_v68), (p3).Plus(_this.F29), _306_v66, (_315_v75).Cardinality(), globalState)
				(globalState).F21 = !((_this.F29).Cmp((_this).Fm29(globalState)) != 0)
			} else {
				var _nw34 *C0 = New_C0_()
				_ = _nw34
				_nw34.Ctor__(p1)
				_301_v61 = _nw34
				var _318_v78 _dafny.Array
				_ = _318_v78
				var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
				_ = _nw35
				_318_v78 = _nw35
				var _319_v79 D15
				_ = _319_v79
				_319_v79 = Companion_D15_.Create_DC32_(_318_v78)
				_318_v78 = (_319_v79).Dtor_cf43()
				var _320_v80 D14
				_ = _320_v80
				_320_v80 = Companion_D14_.Create_DC28_(_dafny.SetOf(_this.F29))
				var _321_v81 D7
				_ = _321_v81
				_321_v81 = Companion_D7_.Create_DC15_(_300_v60)
				var _322_v82 D7
				_ = _322_v82
				_322_v82 = Companion_D7_.Create_DC17_(_321_v81)
				var _323_v83 D7
				_ = _323_v83
				_323_v83 = Companion_D7_.Create_DC17_(_322_v82)
				var _324_v84 D7
				_ = _324_v84
				_324_v84 = Companion_D7_.Create_DC17_(_322_v82)
				var _325_v85 D7
				_ = _325_v85
				_325_v85 = Companion_D7_.Create_DC17_(_322_v82)
				var _326_v86 D7
				_ = _326_v86
				_326_v86 = Companion_D7_.Create_DC17_(_323_v83)
				var _327_v87 _dafny.Map
				_ = _327_v87
				_327_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v80, _326_v86)
				var _328_v88 _dafny.Set
				_ = _328_v88
				_328_v88 = _dafny.SetOf((_dafny.MultiSetFromSeq(_300_v60)).Cardinality())
				_327_v87 = (_327_v87).Update(Companion_D14_.Create_DC28_(_328_v88), Companion_D7_.Create_DC17_(_324_v84))
				var _329_v89 _dafny.Array
				_ = _329_v89
				var _nwElement0_3 _dafny.Sequence = _305_v65
				_ = _nwElement0_3
				var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(21))
				_ = _nw36
				(_nw36).ArraySet1(_nwElement0_3, 0)
				(_nw36).ArraySet1(_305_v65, 1)
				(_nw36).ArraySet1(_305_v65, 2)
				(_nw36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_305_v65, _305_v65), 3)
				(_nw36).ArraySet1(_dafny.SeqOf(_307_v68, _307_v68, _307_v68), 4)
				(_nw36).ArraySet1(_305_v65, 5)
				(_nw36).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_330_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_331_i2 _dafny.Int) _dafny.CodePoint {
						return _330_v68
					}
				})(_307_v68))), 6)
				(_nw36).ArraySet1(_305_v65, 7)
				(_nw36).ArraySet1(_dafny.SeqOf(_307_v68, _307_v68, _307_v68, _307_v68, _307_v68), 8)
				(_nw36).ArraySet1(_305_v65, 9)
				(_nw36).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(631))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_332_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_333_i3 _dafny.Int) _dafny.CodePoint {
						return _332_v68
					}
				})(_307_v68))), 10)
				(_nw36).ArraySet1(_305_v65, 11)
				(_nw36).ArraySet1(_305_v65, 12)
				(_nw36).ArraySet1(_305_v65, 13)
				(_nw36).ArraySet1(_305_v65, 14)
				(_nw36).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-368))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_334_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_335_i4 _dafny.Int) _dafny.CodePoint {
						return _334_v68
					}
				})(_307_v68))), 15)
				(_nw36).ArraySet1(_305_v65, 16)
				(_nw36).ArraySet1(_305_v65, 17)
				(_nw36).ArraySet1(_305_v65, 18)
				(_nw36).ArraySet1(_dafny.SeqOf(_307_v68), 19)
				(_nw36).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-300))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_336_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_337_i5 _dafny.Int) _dafny.CodePoint {
						return _336_v68
					}
				})(_307_v68))), 20)
				_329_v89 = _nw36
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_329_v89), 0))
				_ = _index51
				(_329_v89).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_338_v68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_339_i6 _dafny.Int) _dafny.CodePoint {
						return _338_v68
					}
				})(_307_v68))), (_index51).Int())
				var _340_v90 D3
				_ = _340_v90
				_340_v90 = Companion_D3_.Create_DC5_(_307_v68)
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_329_v89), 0))
				_ = _index52
				(_329_v89).ArraySet1((func() _dafny.Sequence {
					if _301_v61.F27 {
						return _305_v65
					}
					return _dafny.SeqOf((_340_v90).Dtor_cf5())
				})(), (_index52).Int())
				var _341_v91 *C0
				_ = _341_v91
				var _nw37 *C0 = New_C0_()
				_ = _nw37
				_nw37.Ctor__(_301_v61.F27)
				_341_v91 = _nw37
			}
			var _342_v92 _dafny.Sequence
			_ = _342_v92
			_342_v92 = _dafny.SeqOf((func() _dafny.Array {
				if _301_v61.F27 {
					return _313_v73
				}
				return _313_v73
			})(), _313_v73)
			_313_v73 = (_342_v92).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_342_v92).Cardinality()))).Uint32()).(_dafny.Array)
		} else {
			_307_v68 = _307_v68
			var _343_v93 *C0
			_ = _343_v93
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__(_301_v61.F27)
			_343_v93 = _nw38
			var _344_v94 _dafny.MultiSet
			_ = _344_v94
			_344_v94 = _dafny.MultiSetOf(p3, p3)
			(globalState).F2 = (Companion_Default___.Fm0(_dafny.IntOfInt64(268), p3, (_344_v94).Cardinality(), _301_v61.F27, globalState)).Cmp((_dafny.IntOfInt64(36)).Plus(_this.F29)) < 0
			var _345_v95 _dafny.Map
			_ = _345_v95
			_345_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v61.F27, p0)
			var _346_v96 _dafny.Map
			_ = _346_v96
			_346_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if _343_v93.F27 {
					return _345_v95
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), p3)
			})(), p3)
			var _347_v97 _dafny.Map
			_ = _347_v97
			_347_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_D14_.Create_DC30_(_300_v60, _302_v62))
			_346_v96 = (_346_v96).Update((Companion_Default___.Fm35(p1, p3, globalState)).Merge(_345_v95), (_300_v60).Select((Companion_Default___.SafeIndex((_347_v97).Cardinality(), _dafny.IntOfUint32((_300_v60).Cardinality()))).Uint32()).(_dafny.Int))
			(globalState).F7 = p3
		}
		var _348_v98 _dafny.Map
		_ = _348_v98
		_348_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _301_v61.F27)
		var _349_v99 _dafny.Map
		_ = _349_v99
		_349_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, (_this).Fm7(!((func() bool {
			if (_348_v98).Contains(_this.F29) {
				return (_348_v98).Get(_this.F29).(bool)
			}
			return _301_v61.F27
		})()), p0, _301_v61.F27, globalState))
		(globalState).F0 = ((_349_v99).Cardinality()).Cmp(p3) == 0
		var _350_v100 _dafny.Map
		_ = _350_v100
		_350_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v61.F27, p3)
		var _351_v101 _dafny.MultiSet
		_ = _351_v101
		_351_v101 = _dafny.MultiSetOf(p2)
		var _352_v102 _dafny.MultiSet
		_ = _352_v102
		_352_v102 = _dafny.MultiSetOf(p0, p3, _this.F29)
		Companion_Default___.M0(false, p0, Companion_Default___.Fm36((_350_v100).Cardinality(), p0, (func() _dafny.Int {
			if (_351_v101).Contains(p2) {
				return (_351_v101).Multiplicity(p2)
			}
			return p0
		})(), globalState), (_352_v102).Cardinality(), globalState)
		r0 = p1
		var _353_v103 _dafny.MultiSet
		_ = _353_v103
		_353_v103 = _dafny.MultiSetOf(_301_v61.F27, p1)
		r1 = (func() _dafny.Int {
			if (_353_v103).Contains(_301_v61.F27) {
				return (_353_v103).Multiplicity(_301_v61.F27)
			}
			return _dafny.IntOfInt64(359)
		})()
		var _354_v104 _dafny.Array
		_ = _354_v104
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
		_ = _nw39
		_354_v104 = _nw39
		r2 = _354_v104
		return r0, r1, r2
	}
}
func (_this *C1) M3(p0 _dafny.Array, globalState *GlobalState) {
	{
		var _355_v0 _dafny.Map
		_ = _355_v0
		_355_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _this.F29)
		var _356_v1 _dafny.Sequence
		_ = _356_v1
		_356_v1 = _dafny.UnicodeSeqOfUtf8Bytes("kx")
		var _357_v2 _dafny.Map
		_ = _357_v2
		_357_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_355_v0, _356_v1)
		var _358_v3 _dafny.Array
		_ = _358_v3
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_7
		var _nw40 _dafny.Array
		_ = _nw40
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw40 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) bool = func(_359_i1 _dafny.Int) bool {
				return false
			}
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw40 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw40).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw40).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_358_v3 = _nw40
		(globalState).F0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _358_v3)).Contains((_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_357_v2).Contains(_355_v0) {
				return (_357_v2).Get(_355_v0).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-463))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_360_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))
		})()).Cardinality())).Plus(_this.F29))
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))
		_ = _index53
		(p0).ArraySet1(_this.F29, (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))
		_ = _index54
		(p0).ArraySet1((_dafny.IntOfUint32((_356_v1).Cardinality())).Minus(_this.F29), (_index54).Int())
		var _361_v4 _dafny.MultiSet
		_ = _361_v4
		_361_v4 = _dafny.MultiSetOf(_358_v3)
		var _362_i2 _dafny.Int
		_ = _362_i2
		_362_i2 = _dafny.Zero
		{
			for ((_361_v4).Cardinality()).Cmp(_dafny.IntOfInt64(523)) == 0 {
				{
					if (_362_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_362_i2 = (_362_i2).Plus(_dafny.One)
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))
					_ = _index55
					(p0).ArraySet1(((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Plus(_this.F29), (_index55).Int())
					var _363_v5 _dafny.CodePoint
					_ = _363_v5
					_363_v5 = _dafny.CodePoint('d')
					_363_v5 = _363_v5
					var _364_v6 bool
					_ = _364_v6
					_364_v6 = true
					var _365_v7 *C0
					_ = _365_v7
					var _nw41 *C0 = New_C0_()
					_ = _nw41
					_nw41.Ctor__(!(_364_v6))
					_365_v7 = _nw41
					var _366_v8 _dafny.MultiSet
					_ = _366_v8
					_366_v8 = _dafny.MultiSetOf(_365_v7.F27)
					var _367_v9 _dafny.Map
					_ = _367_v9
					_367_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if (_366_v8).Contains(_364_v6) {
							return (_366_v8).Multiplicity(_364_v6)
						}
						return (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)
					})(), _365_v7.F27)
					var _368_v10 _dafny.Sequence
					_ = _368_v10
					_368_v10 = _dafny.SeqOf(_this.F29, ((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Times((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)), ((_367_v9).Cardinality()).Plus(_this.F29), _this.F29, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
					_368_v10 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm37(globalState), Companion_Default___.Fm37(globalState))
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _369_v11 _dafny.Array
		_ = _369_v11
		var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw42
		_369_v11 = _nw42
		for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_369_v11), 0))); ; {
			_guard_loop_1, _ok32 := _iter32()
			if !_ok32 {
				break
			}
			var _370_i3 _dafny.Int
			_370_i3 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_370_i3).Sign() != -1) && ((_370_i3).Cmp(_dafny.ArrayLen((_369_v11), 0)) < 0)) {
				(_369_v11).ArraySet1((_370_i3).Minus((_this.F29).Minus(_this.F29)), (_370_i3).Int())
			}
		}
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_358_v3), 0))); ; {
			_guard_loop_2, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _371_i4 _dafny.Int
			_371_i4 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_371_i4).Sign() != -1) && ((_371_i4).Cmp(_dafny.ArrayLen((_358_v3), 0)) < 0)) {
				(_358_v3).ArraySet1(true, (_371_i4).Int())
			}
		}
		var _372_v12 bool
		_ = _372_v12
		_372_v12 = false
		var _373_v13 _dafny.Sequence
		_ = _373_v13
		_373_v13 = _dafny.SeqOf(false, _372_v12, _372_v12, (_this).Fm6(_dafny.IntOfInt64(-244), globalState), _372_v12)
		var _374_v14 _dafny.Map
		_ = _374_v14
		_374_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_372_v12), _373_v13)
		var _hi0 _dafny.Int = ((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Times(_this.F29)
		_ = _hi0
		for _375_i5 := _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_374_v14).Contains(_372_v12) {
				return (_374_v14).Get(_372_v12).(_dafny.Sequence)
			}
			return _dafny.SeqOf(_372_v12)
		})()).Cardinality()); _375_i5.Cmp(_hi0) < 0; _375_i5 = _375_i5.Plus(_dafny.One) {
			var _376_v15 _dafny.CodePoint
			_ = _376_v15
			_376_v15 = _dafny.CodePoint('e')
			var _377_v16 _dafny.Map
			_ = _377_v16
			_377_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_375_i5, _376_v15)
			_377_v16 = (_377_v16).Update((_dafny.Zero).Minus((_dafny.IntOfUint32((_373_v13).Cardinality())).Times(_dafny.IntOfInt64(-703))), _376_v15)
			var _378_v17 D12
			_ = _378_v17
			_378_v17 = Companion_D12_.Create_DC23_(_dafny.IntOfUint32((_356_v1).Cardinality()))
			(globalState).F7 = (_378_v17).Dtor_cf30()
			var _379_v18 _dafny.Sequence
			_ = _379_v18
			_379_v18 = _dafny.SeqOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(651), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _this.F29)
			var _380_v19 _dafny.Map
			_ = _380_v19
			_380_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_379_v18, _369_v11)
			var _381_v20 _dafny.Map
			_ = _381_v20
			_381_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_380_v19).Contains(_379_v18) {
					return (_380_v19).Get(_379_v18).(_dafny.Array)
				}
				return _369_v11
			})(), _dafny.SeqOf(_379_v18, _379_v18))
			var _382_v21 _dafny.Sequence
			_ = _382_v21
			_382_v21 = _dafny.SeqOf(_379_v18)
			_381_v20 = (_381_v20).Update(_369_v11, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_383_p0 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
				return func(_384_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_383_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_383_p0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(368), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jcxjkvp")).Cardinality()))
				}
			})(p0))), _382_v21))
			var _385_v22 _dafny.MultiSet
			_ = _385_v22
			_385_v22 = _dafny.MultiSetOf(_379_v18, _379_v18, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-816))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}((func(_386_i5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_387_i7 _dafny.Int) _dafny.Int {
					return _386_i5
				}
			})(_375_i5))), _dafny.SeqOf(_375_i5, (_dafny.Zero).Minus(_375_i5)), _dafny.Companion_Sequence_.Update(_379_v18, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.IntOfUint32((_379_v18).Cardinality()))).Uint32(), _dafny.IntOfInt64(206)))
			var _388_v23 _dafny.MultiSet
			_ = _388_v23
			_388_v23 = _dafny.MultiSetOf(_356_v1)
			var _rhs36 bool = ((_355_v0).Cardinality()).Cmp(_dafny.IntOfInt64(509)) == 0
			_ = _rhs36
			var _rhs37 bool = ((func() _dafny.Int {
				if (_385_v22).Contains(_379_v18) {
					return (_385_v22).Multiplicity(_379_v18)
				}
				return (func() _dafny.Int {
					if (_388_v23).Contains(_356_v1) {
						return (_388_v23).Multiplicity(_356_v1)
					}
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ilycvcw")).Cardinality()))
				})()
			})()).Cmp((_dafny.IntOfInt64(33)).Minus(_this.F29)) >= 0
			_ = _rhs37
			var _lhs40 *GlobalState = globalState
			_ = _lhs40
			var _lhs41 *GlobalState = globalState
			_ = _lhs41
			_lhs40.F21 = _rhs36
			_lhs41.F13 = _rhs37
		}
	}
}
func (_this *C1) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _389_v0 bool
		_ = _389_v0
		_389_v0 = false
		(globalState).F0 = (func() bool {
			if _389_v0 {
				return (p0).Cmp(p0) != 0
			}
			return _389_v0
		})()
		if (_this).Fm7(_389_v0, (func() _dafny.Int {
			if _389_v0 {
				return _this.F29
			}
			return p0
		})(), (p0).Cmp((_dafny.MultiSetOf(p0)).Cardinality()) == 0, globalState) {
			var _390_v1 _dafny.MultiSet
			_ = _390_v1
			_390_v1 = _dafny.MultiSetOf(_this.F29)
			var _391_v2 _dafny.Map
			_ = _391_v2
			_391_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_390_v1, _389_v0)
			var _392_v3 _dafny.Set
			_ = _392_v3
			_392_v3 = _dafny.SetOf(false)
			var _393_v4 _dafny.Sequence
			_ = _393_v4
			_393_v4 = _dafny.UnicodeSeqOfUtf8Bytes("by")
			var _394_v5 _dafny.Sequence
			_ = _394_v5
			_394_v5 = _dafny.SeqOf((func() _dafny.Int {
				if (_390_v1).Contains((_392_v3).Cardinality()) {
					return (_390_v1).Multiplicity((_392_v3).Cardinality())
				}
				return p0
			})(), p0, _dafny.IntOfUint32((_393_v4).Cardinality()))
			var _source8 _dafny.Map = (_391_v2).Update(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_394_v5, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.IntOfUint32((_394_v5).Cardinality()))).Uint32(), _this.F29)), _389_v0)
			_ = _source8
			var _395___mcc_h0 _dafny.Map = _source8
			_ = _395___mcc_h0
			var _396_cf27 _dafny.Map = _395___mcc_h0
			_ = _396_cf27
			(globalState).F13 = _389_v0
			(globalState).F7 = p0
			var _397_v6 *C0
			_ = _397_v6
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__(_389_v0)
			_397_v6 = _nw43
			(globalState).F21 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_393_v4, _dafny.UnicodeSeqOfUtf8Bytes("sbarm")), _393_v4)
			var _398_v7 D13
			_ = _398_v7
			_398_v7 = Companion_D13_.Create_DC26_(_393_v4, _389_v0, _389_v0)
			var _399_v8 _dafny.Set
			_ = _399_v8
			_399_v8 = _dafny.SetOf(_393_v4, (_398_v7).Dtor_cf33(), _393_v4)
			(globalState).F10 = (_399_v8).Intersection(_dafny.SetOf(_393_v4, _393_v4))
			var _400_v9 *C0
			_ = _400_v9
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__(_389_v0)
			_400_v9 = _nw44
			(globalState).F7 = _this.F29
			var _401_v10 _dafny.CodePoint
			_ = _401_v10
			_401_v10 = _dafny.CodePoint('l')
			var _402_v11 _dafny.Set
			_ = _402_v11
			_402_v11 = _dafny.SetOf(_401_v10)
			var _403_v12 _dafny.Set
			_ = _403_v12
			_403_v12 = _dafny.SetOf((_402_v11).Union(_dafny.SetOf(_401_v10)))
			_403_v12 = Companion_Default___.Fm38(globalState)
		} else {
			if false {
				(globalState).F7 = (func() _dafny.Int {
					if _389_v0 {
						return _dafny.IntOfInt64(366)
					}
					return p0
				})()
				var _404_v13 _dafny.Array
				_ = _404_v13
				var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
				_ = _nw45
				_404_v13 = _nw45
				var _405_v14 _dafny.CodePoint
				_ = _405_v14
				_405_v14 = _dafny.CodePoint('b')
				var _406_v15 _dafny.Map
				_ = _406_v15
				_406_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_389_v0, _405_v14)
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_404_v13), 0))
				_ = _index56
				(_404_v13).ArraySet1(_406_v15, (_index56).Int())
				var _407_v16 _dafny.Map
				_ = _407_v16
				_407_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _389_v0)
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_404_v13), 0))
				_ = _index57
				(_404_v13).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_407_v16, globalState), _405_v14)).Merge(_406_v15), (_index57).Int())
				var _408_v17 D7
				_ = _408_v17
				_408_v17 = Companion_D7_.Create_DC15_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(279))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}(func(_409_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kueylnvh")).Cardinality())
				})))
				var _410_v18 D7
				_ = _410_v18
				_410_v18 = Companion_D7_.Create_DC17_(_408_v17)
				var _411_v19 _dafny.Sequence
				_ = _411_v19
				_411_v19 = _dafny.SeqOf(_410_v18)
				(globalState).F21 = _dafny.Companion_Sequence_.IsProperPrefixOf(_411_v19, _dafny.Companion_Sequence_.Concatenate(_411_v19, _411_v19))
				var _412_v20 _dafny.Array
				_ = _412_v20
				var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw46
				_412_v20 = _nw46
				r0 = _412_v20
				(globalState).F0 = _389_v0
			} else {
				(globalState).F16 = _dafny.IntOfInt64(748)
				(globalState).F7 = _this.F29
				var _413_v21 _dafny.Array
				_ = _413_v21
				var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw47
				_413_v21 = _nw47
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_413_v21), 0))
				_ = _index58
				(_413_v21).ArraySet1((_dafny.Zero).Minus(_this.F29), (_index58).Int())
				var _414_v22 _dafny.Map
				_ = _414_v22
				_414_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_389_v0, p0)
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_413_v21), 0))
				_ = _index59
				(_413_v21).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_414_v22).Contains(_389_v0) {
						return (_414_v22).Get(_389_v0).(_dafny.Int)
					}
					return _this.F29
				})(), p0), (_index59).Int())
				var _415_v23 _dafny.MultiSet
				_ = _415_v23
				_415_v23 = _dafny.MultiSetOf(_389_v0)
				var _416_v24 _dafny.Sequence
				_ = _416_v24
				_416_v24 = _dafny.SeqOf(_415_v23)
				var _417_v25 _dafny.MultiSet
				_ = _417_v25
				_417_v25 = (_416_v24).Select((Companion_Default___.SafeIndex((_414_v22).Cardinality(), _dafny.IntOfUint32((_416_v24).Cardinality()))).Uint32()).(_dafny.MultiSet)
				_417_v25 = _417_v25
				(globalState).F13 = (_389_v0) || (_389_v0)
			}
			var _418_v26 *C0
			_ = _418_v26
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__(_389_v0)
			_418_v26 = _nw48
			(globalState).F7 = p0
			var _419_v27 _dafny.Map
			_ = _419_v27
			_419_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_418_v26.F27, (_this).Fm29(globalState))
			var _420_v28 _dafny.MultiSet
			_ = _420_v28
			_420_v28 = _dafny.MultiSetOf(_419_v27)
			var _421_v29 _dafny.Map
			_ = _421_v29
			_421_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_this.F29)).Cardinality(), _418_v26.F27)
			var _422_v30 _dafny.Sequence
			_ = _422_v30
			_422_v30 = _dafny.SeqOf((_421_v29).Cardinality(), _this.F29)
			(globalState).F2 = (_420_v28).IsSubsetOf((_420_v28).Difference(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetFromSeq(_422_v30)).Cardinality()), _419_v27, _419_v27, _419_v27, _419_v27)))
			(globalState).F7 = p0
		}
		var _423_v31 *C0
		_ = _423_v31
		var _nw49 *C0 = New_C0_()
		_ = _nw49
		_nw49.Ctor__(_389_v0)
		_423_v31 = _nw49
		var _424_i1 _dafny.Int
		_ = _424_i1
		_424_i1 = _dafny.Zero
		{
			for (_this.F29).Cmp((_this).Fm29(globalState)) == 0 {
				{
					if (_424_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_424_i1 = (_424_i1).Plus(_dafny.One)
					(globalState).F13 = !(_389_v0)
					var _425_v32 _dafny.Array
					_ = _425_v32
					var _len0_8 _dafny.Int = _dafny.IntOfInt64(2)
					_ = _len0_8
					var _nw50 _dafny.Array
					_ = _nw50
					if _len0_8.Cmp(_dafny.Zero) == 0 {
						_nw50 = _dafny.NewArray(_len0_8)
					} else {
						var _init8 func(_dafny.Int) _dafny.Int = (func(_426_v0 bool, _427_v31 *C0) func(_dafny.Int) _dafny.Int {
							return func(_428_i2 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_428_i2, (_dafny.MultiSetOf(_426_v0, !(_427_v31.F27))).Cardinality())
							}
						})(_389_v0, _423_v31)
						_ = _init8
						var _element0_8 = _init8(_dafny.Zero)
						_ = _element0_8
						_nw50 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
						(_nw50).ArraySet1(_element0_8, 0)
						var _nativeLen0_8 = (_len0_8).Int()
						_ = _nativeLen0_8
						for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
							(_nw50).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
						}
					}
					_425_v32 = _nw50
					var _429_v33 _dafny.Set
					_ = _429_v33
					_429_v33 = _dafny.SetOf(_425_v32, _425_v32)
					var _430_v34 _dafny.Sequence
					_ = _430_v34
					_430_v34 = _dafny.SeqOf(_389_v0, _389_v0, _423_v31.F27)
					var _431_v35 _dafny.Sequence
					_ = _431_v35
					_431_v35 = _dafny.UnicodeSeqOfUtf8Bytes("xjuxav")
					var _432_v36 _dafny.Array
					_ = _432_v36
					var _nwElement0_4 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm1((_429_v33).Cardinality(), globalState)).Cardinality())
					_ = _nwElement0_4
					var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(11))
					_ = _nw51
					(_nw51).ArraySet1(_nwElement0_4, 0)
					(_nw51).ArraySet1(p0, 1)
					(_nw51).ArraySet1(_dafny.IntOfUint32((_430_v34).Cardinality()), 2)
					(_nw51).ArraySet1((_dafny.IntOfUint32((_431_v35).Cardinality())).Plus(_dafny.IntOfInt64(-830)), 3)
					(_nw51).ArraySet1(p0, 4)
					(_nw51).ArraySet1((_dafny.IntOfUint32((_431_v35).Cardinality())).Minus(_dafny.IntOfUint32((_431_v35).Cardinality())), 5)
					(_nw51).ArraySet1(p0, 6)
					(_nw51).ArraySet1(_this.F29, 7)
					(_nw51).ArraySet1(_this.F29, 8)
					(_nw51).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfInt64(-110), p0, _this.F29, _389_v0, globalState), 9)
					(_nw51).ArraySet1(_dafny.IntOfInt64(-680), 10)
					_432_v36 = _nw51
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_425_v32), 0))
					_ = _index60
					(_425_v32).ArraySet1(_this.F29, (_index60).Int())
					var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_425_v32), 0))
					_ = _index61
					(_425_v32).ArraySet1(_this.F29, (_index61).Int())
					var _433_v37 _dafny.Map
					_ = _433_v37
					_433_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_425_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_425_v32), 0))).Int()).(_dafny.Int), _423_v31.F27)
					var _434_v38 D1
					_ = _434_v38
					_434_v38 = Companion_D1_.Create_DC2_((_433_v37).Merge(_433_v37))
					_434_v38 = Companion_D1_.Create_DC2_(_433_v37)
					var _435_v39 _dafny.Map
					_ = _435_v39
					_435_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_423_v31.F27, _this.F29)
					_435_v39 = ((_435_v39).Merge(_435_v39)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_389_v0, _dafny.IntOfInt64(811)))
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _436_v40 *C0
		_ = _436_v40
		var _nw52 *C0 = New_C0_()
		_ = _nw52
		_nw52.Ctor__(_423_v31.F27)
		_436_v40 = _nw52
		var _437_v41 _dafny.Map
		_ = _437_v41
		_437_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_389_v0), _436_v40.F27)
		var _rhs38 _dafny.Int = (_this).Fm29(globalState)
		_ = _rhs38
		var _rhs39 _dafny.Map = _437_v41
		_ = _rhs39
		var _lhs42 *C1 = _this
		_ = _lhs42
		_lhs42.F29 = _rhs38
		_437_v41 = _rhs39
		var _438_v42 _dafny.Array
		_ = _438_v42
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw53
		_438_v42 = _nw53
		r0 = _438_v42
		r1 = _436_v40.F27
		return r0, r1
	}
}
func (_this *C1) F30() _dafny.Sequence {
	{
		return _this._f30
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f28 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f28 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f28 _dafny.Int) {
	{
		(_this)._f28 = f28
	}
}
func (_this *C2) Fm7(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return (((_this).F28()).Cmp((_dafny.Zero).Minus((_this).F28())) > 0) == ((func() bool {
			if !(!(true)) {
				return false
			}
			return false
		})())
	}
}
func (_this *C2) Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var _source9 D1 = Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-844), false))
		_ = _source9
		if _source9.Is_DC3() {
			var _439___mcc_h0 bool = _source9.Get_().(D1_DC3).Cf3
			_ = _439___mcc_h0
			var _440_cf3 bool = _439___mcc_h0
			_ = _440_cf3
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_440_cf3, _dafny.UnicodeSeqOfUtf8Bytes("lcwihbpu"))
		} else {
			var _441___mcc_h1 _dafny.Map = _source9.Get_().(D1_DC2).Cf2
			_ = _441___mcc_h1
			var _442_cf2 _dafny.Map = _441___mcc_h1
			_ = _442_cf2
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.UnicodeSeqOfUtf8Bytes("duyckw"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.UnicodeSeqOfUtf8Bytes("mtfgtk")))
		}
	}
}
func (_this *C2) Fm6(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C2) Fm17(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		if !(true) || (false) {
			return (_this).F28()
		} else {
			return (_this).F28()
		}
	}
}
func (_this *C2) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _443_v0 _dafny.Sequence
		_ = _443_v0
		_443_v0 = _dafny.UnicodeSeqOfUtf8Bytes("klcc")
		var _hi1 _dafny.Int = _dafny.IntOfUint32((_443_v0).Cardinality())
		_ = _hi1
		for _444_i0 := (_this).F28(); _444_i0.Cmp(_hi1) < 0; _444_i0 = _444_i0.Plus(_dafny.One) {
			(globalState).F21 = p1
			var _445_v1 *C0
			_ = _445_v1
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__(p1)
			_445_v1 = _nw54
			var _446_v2 _dafny.MultiSet
			_ = _446_v2
			_446_v2 = _dafny.MultiSetOf(_445_v1.F27)
			var _447_v3 _dafny.MultiSet
			_ = _447_v3
			_447_v3 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt((_446_v2).Cardinality(), (_this).F28()), _444_i0)
			var _448_v4 D4
			_ = _448_v4
			_448_v4 = Companion_D4_.Create_DC9_((_dafny.Zero).Minus(p0))
			var _rhs40 bool = _445_v1.F27
			_ = _rhs40
			var _rhs41 _dafny.Int = (func() _dafny.Int {
				if (_447_v3).Contains((_this).F28()) {
					return (_447_v3).Multiplicity((_this).F28())
				}
				return (_dafny.Zero).Minus((func() _dafny.Int {
					if (_447_v3).Contains(p0) {
						return (_447_v3).Multiplicity(p0)
					}
					return p3
				})())
			})()
			_ = _rhs41
			var _rhs42 _dafny.Int = ((_448_v4).Dtor_cf11()).Plus(p3)
			_ = _rhs42
			var _rhs43 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, _444_i0))
			_ = _rhs43
			var _lhs43 *C0 = _445_v1
			_ = _lhs43
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			_lhs43.F27 = _rhs40
			_lhs44.F7 = _rhs41
			_lhs45.F7 = _rhs42
			_lhs46.F7 = _rhs43
			var _449_v5 _dafny.Sequence
			_ = _449_v5
			_449_v5 = _dafny.SeqOf(p1)
			var _450_v6 _dafny.Sequence
			_ = _450_v6
			_450_v6 = _449_v5
			var _451_v7 _dafny.CodePoint
			_ = _451_v7
			_451_v7 = _dafny.CodePoint('q')
			var _452_v8 _dafny.Array
			_ = _452_v8
			var _nwElement0_5 _dafny.Sequence = _449_v5
			_ = _nwElement0_5
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(19))
			_ = _nw55
			(_nw55).ArraySet1(_nwElement0_5, 0)
			(_nw55).ArraySet1(_449_v5, 1)
			(_nw55).ArraySet1(_dafny.Companion_Sequence_.Update(_449_v5, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.IntOfUint32((_449_v5).Cardinality()))).Uint32(), _445_v1.F27), 2)
			(_nw55).ArraySet1(_449_v5, 3)
			(_nw55).ArraySet1((_450_v6), 4)
			(_nw55).ArraySet1(_dafny.SeqOf(_445_v1.F27), 5)
			(_nw55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_449_v5, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm18(_445_v1.F27, _dafny.CodePoint('p'), _445_v1.F27, globalState), (Companion_Default___.SafeIndex(_444_i0, _dafny.IntOfUint32((Companion_Default___.Fm18(_445_v1.F27, _dafny.CodePoint('p'), _445_v1.F27, globalState)).Cardinality()))).Uint32(), _445_v1.F27)), 6)
			(_nw55).ArraySet1(_449_v5, 7)
			(_nw55).ArraySet1(_449_v5, 8)
			(_nw55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_449_v5, _449_v5), 9)
			(_nw55).ArraySet1(_449_v5, 10)
			(_nw55).ArraySet1(_dafny.SeqOf(_445_v1.F27), 11)
			(_nw55).ArraySet1(_449_v5, 12)
			(_nw55).ArraySet1(_dafny.SeqOf(p1, _445_v1.F27, p1, _445_v1.F27), 13)
			(_nw55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm18(!(_445_v1.F27), _451_v7, _445_v1.F27, globalState), _449_v5), 14)
			(_nw55).ArraySet1(_dafny.Companion_Sequence_.Update(_449_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_449_v5).Cardinality()))).Uint32(), _445_v1.F27), 15)
			(_nw55).ArraySet1(_dafny.SeqOf(_445_v1.F27), 16)
			(_nw55).ArraySet1(_449_v5, 17)
			(_nw55).ArraySet1(_449_v5, 18)
			_452_v8 = _nw55
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_452_v8), 0))
			_ = _index62
			(_452_v8).ArraySet1(_449_v5, (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_452_v8), 0))
			_ = _index63
			(_452_v8).ArraySet1(_449_v5, (_index63).Int())
		}
		var _453_v9 _dafny.Array
		_ = _453_v9
		var _nw56 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(20))
		_ = _nw56
		_453_v9 = _nw56
		for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_453_v9), 0))); ; {
			_guard_loop_3, _ok34 := _iter34()
			if !_ok34 {
				break
			}
			var _454_i1 _dafny.Int
			_454_i1 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_454_i1).Sign() != -1) && ((_454_i1).Cmp(_dafny.ArrayLen((_453_v9), 0)) < 0)) {
				(_453_v9).ArraySet1(Companion_D3_.Create_DC6_(_443_v0, _dafny.CodePoint('v'), p1), (_454_i1).Int())
			}
		}
		var _455_v10 D1
		_ = _455_v10
		_455_v10 = Companion_D1_.Create_DC3_(p1)
		var _456_v11 _dafny.Map
		_ = _456_v11
		_456_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_455_v10).Dtor_cf3())
		var _457_v12 D4
		_ = _457_v12
		_457_v12 = Companion_D4_.Create_DC7_((_456_v11).Cardinality())
		var _458_v13 _dafny.Set
		_ = _458_v13
		_458_v13 = _dafny.SetOf(false)
		var _459_v14 _dafny.Map
		_ = _459_v14
		_459_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-94))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg50 _dafny.Int) interface{} {
				return coer50(arg50)
			}
		}(func(_460_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})), _458_v13)
		var _461_v15 _dafny.Sequence
		_ = _461_v15
		_461_v15 = _dafny.SeqOf(p0, p0)
		var _462_v16 _dafny.Sequence
		_ = _462_v16
		_462_v16 = _dafny.SeqOf(p1, p1, p1)
		var _463_v17 _dafny.Array
		_ = _463_v17
		var _nwElement0_6 _dafny.Set = (func() _dafny.Set {
			if (_459_v14).Contains(_443_v0) {
				return (_459_v14).Get(_443_v0).(_dafny.Set)
			}
			return _458_v13
		})()
		_ = _nwElement0_6
		var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw57
		(_nw57).ArraySet1(_nwElement0_6, 0)
		(_nw57).ArraySet1((_dafny.SetOf(p1)).Intersection(_458_v13), 1)
		(_nw57).ArraySet1((_458_v13).Intersection(_458_v13), 2)
		(_nw57).ArraySet1(_458_v13, 3)
		(_nw57).ArraySet1(_458_v13, 4)
		(_nw57).ArraySet1(_458_v13, 5)
		(_nw57).ArraySet1(_458_v13, 6)
		(_nw57).ArraySet1(_458_v13, 7)
		(_nw57).ArraySet1(Companion_Default___.Fm19(_dafny.IntOfUint32((_461_v15).Cardinality()), (_462_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.IntOfUint32((_462_v16).Cardinality()))).Uint32()).(bool), p1, _dafny.IntOfInt64(-900), globalState), 8)
		(_nw57).ArraySet1((_458_v13).Union(_dafny.SetOf(p1)), 9)
		(_nw57).ArraySet1(_458_v13, 10)
		(_nw57).ArraySet1((_458_v13).Union(_458_v13), 11)
		(_nw57).ArraySet1((_458_v13).Intersection(_dafny.SetOf(p1, p1, p1, p1, p1)), 12)
		_463_v17 = _nw57
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_463_v17), 0))
		_ = _index64
		(_463_v17).ArraySet1(_458_v13, (_index64).Int())
		var _464_v18 *C0
		_ = _464_v18
		var _nw58 *C0 = New_C0_()
		_ = _nw58
		_nw58.Ctor__((p3).Cmp((_this).F28()) < 0)
		_464_v18 = _nw58
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_463_v17), 0))
		_ = _index65
		var _rhs44 _dafny.Int = (((_this).F28()).Times(p3)).Times(Companion_Default___.SafeDivisionInt(p3, _dafny.IntOfUint32((_461_v15).Cardinality())))
		_ = _rhs44
		var _rhs45 bool = p1
		_ = _rhs45
		var _rhs46 D4 = _457_v12
		_ = _rhs46
		var _rhs47 _dafny.Set = _458_v13
		_ = _rhs47
		var _rhs48 *C0 = _464_v18
		_ = _rhs48
		var _lhs47 *GlobalState = globalState
		_ = _lhs47
		var _lhs48 *GlobalState = globalState
		_ = _lhs48
		var _lhs49 _dafny.Array = _463_v17
		_ = _lhs49
		var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_463_v17), 0))
		_ = _lhs50
		_lhs47.F7 = _rhs44
		_lhs48.F21 = _rhs45
		_457_v12 = _rhs46
		(_lhs49).ArraySet1(_rhs47, (_lhs50).Int())
		_464_v18 = _rhs48
		(_464_v18).F27 = _464_v18.F27
		var _465_v19 _dafny.Array
		_ = _465_v19
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_9
		var _nw59 _dafny.Array
		_ = _nw59
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw59 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) bool = (func(_466_v18 *C0) func(_dafny.Int) bool {
				return func(_467_i3 _dafny.Int) bool {
					return (func() bool {
						if _466_v18.F27 {
							return _466_v18.F27
						}
						return _466_v18.F27
					})()
				}
			})(_464_v18)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw59 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw59).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw59).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_465_v19 = _nw59
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_465_v19), 0))
		_ = _index66
		(_465_v19).ArraySet1(p1, (_index66).Int())
		var _468_v20 _dafny.CodePoint
		_ = _468_v20
		_468_v20 = _dafny.CodePoint('a')
		var _469_v21 _dafny.Sequence
		_ = _469_v21
		_469_v21 = Companion_Default___.Fm18(_464_v18.F27, _468_v20, false, globalState)
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_465_v19), 0))
		_ = _index67
		var _rhs49 _dafny.Int = (_this).Fm17(_464_v18.F27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(518))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}(func(_470_i4 _dafny.Int) _dafny.Int {
			return (_this).F28()
		})), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(191))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg52 _dafny.Int) interface{} {
				return coer52(arg52)
			}
		}(func(_471_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()), _469_v21, globalState)
		_ = _rhs49
		var _rhs50 bool = (_464_v18.F27) || (_464_v18.F27)
		_ = _rhs50
		var _lhs51 *GlobalState = globalState
		_ = _lhs51
		var _lhs52 _dafny.Array = _465_v19
		_ = _lhs52
		var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_465_v19), 0))
		_ = _lhs53
		_lhs51.F7 = _rhs49
		(_lhs52).ArraySet1(_rhs50, (_lhs53).Int())
		var _472_v22 _dafny.Array
		_ = _472_v22
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_10
		var _nw60 _dafny.Array
		_ = _nw60
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw60 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) D1 = (func(_473_v10 D1) func(_dafny.Int) D1 {
				return func(_474_i7 _dafny.Int) D1 {
					return _473_v10
				}
			})(_455_v10)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw60 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw60).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw60).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_472_v22 = _nw60
		for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_472_v22), 0))); ; {
			_guard_loop_4, _ok35 := _iter35()
			if !_ok35 {
				break
			}
			var _475_i6 _dafny.Int
			_475_i6 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_475_i6).Sign() != -1) && ((_475_i6).Cmp(_dafny.ArrayLen((_472_v22), 0)) < 0)) {
				(_472_v22).ArraySet1((func() D1 {
					if (func() bool {
						if (_456_v11).Contains((_dafny.SetOf(p0, (_this).F28())).Cardinality()) {
							return (_456_v11).Get((_dafny.SetOf(p0, (_this).F28())).Cardinality()).(bool)
						}
						return _464_v18.F27
					})() {
						return _455_v10
					}
					return _455_v10
				})(), (_475_i6).Int())
			}
		}
		var _476_v23 _dafny.Map
		_ = _476_v23
		_476_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		r0 = !(Companion_Default___.Fm2((_476_v23).Update(_464_v18.F27, _464_v18.F27), globalState))
		var _pat_let_tv6 = p3
		_ = _pat_let_tv6
		var _pat_let_tv7 = _465_v19
		_ = _pat_let_tv7
		var _pat_let_tv8 = _465_v19
		_ = _pat_let_tv8
		var _pat_let_tv9 = p3
		_ = _pat_let_tv9
		var _pat_let_tv10 = p3
		_ = _pat_let_tv10
		r1 = func(_source10 D4) _dafny.Int {
			if _source10.Is_DC8() {
				var _477___mcc_h0 bool = _source10.Get_().(D4_DC8).Cf10
				_ = _477___mcc_h0
				var _478_cf10 bool = _477___mcc_h0
				_ = _478_cf10
				return _pat_let_tv6
			} else if _source10.Is_DC9() {
				var _479___mcc_h1 _dafny.Int = _source10.Get_().(D4_DC9).Cf11
				_ = _479___mcc_h1
				var _480_cf11 _dafny.Int = _479___mcc_h1
				_ = _480_cf11
				return ((Companion_D12_.Create_DC22_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_pat_let_tv7), 0))).Int()).(bool), _pat_let_tv9))).Dtor_cf29()).Cardinality()
			} else if _source10.Is_DC10() {
				var _481___mcc_h2 _dafny.Int = _source10.Get_().(D4_DC10).Cf12
				_ = _481___mcc_h2
				var _482___mcc_h3 _dafny.Int = _source10.Get_().(D4_DC10).Cf13
				_ = _482___mcc_h3
				var _483___mcc_h4 bool = _source10.Get_().(D4_DC10).Cf14
				_ = _483___mcc_h4
				var _484___mcc_h5 bool = _source10.Get_().(D4_DC10).Cf15
				_ = _484___mcc_h5
				var _485___mcc_h6 _dafny.Int = _source10.Get_().(D4_DC10).Cf16
				_ = _485___mcc_h6
				var _486_cf16 _dafny.Int = _485___mcc_h6
				_ = _486_cf16
				var _487_cf15 bool = _484___mcc_h5
				_ = _487_cf15
				var _488_cf14 bool = _483___mcc_h4
				_ = _488_cf14
				var _489_cf13 _dafny.Int = _482___mcc_h3
				_ = _489_cf13
				var _490_cf12 _dafny.Int = _481___mcc_h2
				_ = _490_cf12
				return (_pat_let_tv10).Plus(_489_cf13)
			} else {
				var _491___mcc_h7 _dafny.Int = _source10.Get_().(D4_DC7).Cf9
				_ = _491___mcc_h7
				var _492_cf9 _dafny.Int = _491___mcc_h7
				_ = _492_cf9
				return (_this).F28()
			}
		}(Companion_D4_.Create_DC7_(_dafny.IntOfInt64(406)))
		var _493_v24 _dafny.Array
		_ = _493_v24
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
		_ = _nw61
		_493_v24 = _nw61
		r2 = _493_v24
		return r0, r1, r2
	}
}
func (_this *C2) M3(p0 _dafny.Array, globalState *GlobalState) {
	{
		(globalState).F16 = _dafny.IntOfInt64(644)
		var _494_v0 D12
		_ = _494_v0
		_494_v0 = Companion_D12_.Create_DC23_(_dafny.IntOfInt64(7))
		var _source11 D12 = func(_pat_let6_0 D12) D12 {
			return func(_495_dt__update__tmp_h0 D12) D12 {
				return func(_pat_let7_0 _dafny.Int) D12 {
					return func(_496_dt__update_hcf30_h0 _dafny.Int) D12 {
						return Companion_D12_.Create_DC23_(_496_dt__update_hcf30_h0)
					}(_pat_let7_0)
				}(((_this).F28()).Times((_this).F28()))
			}(_pat_let6_0)
		}(_494_v0)
		_ = _source11
		if _source11.Is_DC23() {
			var _497___mcc_h0 _dafny.Int = _source11.Get_().(D12_DC23).Cf30
			_ = _497___mcc_h0
			var _498_cf30 _dafny.Int = _497___mcc_h0
			_ = _498_cf30
			var _499_v1 bool
			_ = _499_v1
			_499_v1 = true
			var _500_v2 _dafny.CodePoint
			_ = _500_v2
			_500_v2 = _dafny.CodePoint('n')
			var _501_v3 _dafny.Map
			_ = _501_v3
			_501_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_499_v1), _500_v2)
			var _502_v4 _dafny.Map
			_ = _502_v4
			_502_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_498_cf30, (_501_v3).Cardinality())
			var _503_v5 _dafny.Map
			_ = _503_v5
			_503_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_499_v1, (_502_v4).Cardinality())
			if !((_503_v5).Update(Companion_Default___.Fm2(Companion_Default___.Fm20(_498_cf30, _499_v1, _499_v1, _499_v1, globalState), globalState), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bkpcvyger")).Cardinality()))).Equals(_503_v5) {
				var _504_v6 _dafny.MultiSet
				_ = _504_v6
				_504_v6 = _dafny.MultiSetOf(_499_v1)
				_499_v1 = ((_504_v6).Difference(_dafny.MultiSetOf(_499_v1))).IsProperSubsetOf(_504_v6)
				var _505_v7 _dafny.Array
				_ = _505_v7
				var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw62
				_505_v7 = _nw62
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_505_v7), 0))
				_ = _index68
				(_505_v7).ArraySet1(p0, (_index68).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_505_v7), 0))
				_ = _index69
				(_505_v7).ArraySet1(p0, (_index69).Int())
				var _506_v8 _dafny.Array
				_ = _506_v8
				var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(6))
				_ = _nw63
				_506_v8 = _nw63
				var _507_v9 _dafny.Map
				_ = _507_v9
				_507_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_499_v1, _499_v1)
				var _508_v10 _dafny.Set
				_ = _508_v10
				_508_v10 = _dafny.SetOf(_507_v9)
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_506_v8), 0))
				_ = _index70
				(_506_v8).ArraySet1(_508_v10, (_index70).Int())
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_506_v8), 0))
				_ = _index71
				(_506_v8).ArraySet1(_508_v10, (_index71).Int())
				var _509_v11 _dafny.Set
				_ = _509_v11
				_509_v11 = _dafny.SetOf(_499_v1)
				var _510_v12 _dafny.Map
				_ = _510_v12
				_510_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_509_v11).Cardinality(), _499_v1)
				(globalState).F2 = !(_510_v12).Contains(_498_cf30)
				var _511_v13 _dafny.Sequence
				_ = _511_v13
				_511_v13 = _dafny.UnicodeSeqOfUtf8Bytes("aaacch")
				var _512_v14 _dafny.Sequence
				_ = _512_v14
				_512_v14 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_511_v13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(966))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_513_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_514_i0 _dafny.Int) _dafny.CodePoint {
						return _513_v2
					}
				})(_500_v2)))), _511_v13, _511_v13)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_505_v7), 0))
				_ = _index72
				var _rhs51 bool = _499_v1
				_ = _rhs51
				var _rhs52 bool = !(_499_v1)
				_ = _rhs52
				var _rhs53 _dafny.Array = _dafny.ArrayCastTo((_505_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_505_v7), 0))).Int()))
				_ = _rhs53
				var _rhs54 _dafny.Sequence = (_512_v14).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F28()), _dafny.IntOfUint32((_512_v14).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs54
				var _rhs55 _dafny.Int = _498_cf30
				_ = _rhs55
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				var _lhs55 *GlobalState = globalState
				_ = _lhs55
				var _lhs56 _dafny.Array = _505_v7
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_505_v7), 0))
				_ = _lhs57
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				var _lhs59 *GlobalState = globalState
				_ = _lhs59
				_lhs54.F0 = _rhs51
				_lhs55.F0 = _rhs52
				(_lhs56).ArraySet1(_rhs53, (_lhs57).Int())
				_lhs58.F19 = _rhs54
				_lhs59.F16 = _rhs55
			} else {
				_500_v2 = _dafny.CodePoint('a')
				(_this).M8(!(!(!(_499_v1))), _499_v1, _499_v1, globalState)
				var _515_v15 _dafny.Sequence
				_ = _515_v15
				_515_v15 = _dafny.SeqOf(_499_v1)
				var _rhs56 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_498_cf30, (_this).F28()))))
				_ = _rhs56
				var _rhs57 _dafny.Int = _dafny.IntOfUint32((_515_v15).Cardinality())
				_ = _rhs57
				var _lhs60 *GlobalState = globalState
				_ = _lhs60
				var _lhs61 *GlobalState = globalState
				_ = _lhs61
				_lhs60.F7 = _rhs56
				_lhs61.F7 = _rhs57
				var _516_v16 _dafny.Sequence
				_ = _516_v16
				_516_v16 = _dafny.UnicodeSeqOfUtf8Bytes("olkogre")
				var _517_v17 _dafny.Map
				_ = _517_v17
				_517_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_516_v16, (Companion_Default___.SafeIndex(_498_cf30, _dafny.IntOfUint32((_516_v16).Cardinality()))).Uint32(), _500_v2), (func() bool {
					if false {
						return !(_499_v1)
					}
					return false
				})())
				var _rhs58 bool = _499_v1
				_ = _rhs58
				var _rhs59 _dafny.Map = _517_v17
				_ = _rhs59
				var _rhs60 bool = _499_v1
				_ = _rhs60
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				var _lhs63 *GlobalState = globalState
				_ = _lhs63
				_lhs62.F0 = _rhs58
				_517_v17 = _rhs59
				_lhs63.F2 = _rhs60
				(globalState).F16 = (_this).F28()
			}
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((p0), 0))
			_ = _index73
			(p0).ArraySet1(_498_cf30, (_index73).Int())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((p0), 0))
			_ = _index74
			var _rhs61 _dafny.Int = (_this).F28()
			_ = _rhs61
			var _rhs62 _dafny.Int = (_dafny.IntOfInt64(-20)).Minus(_498_cf30)
			_ = _rhs62
			var _rhs63 _dafny.CodePoint = _500_v2
			_ = _rhs63
			var _lhs64 _dafny.Array = p0
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((p0), 0))
			_ = _lhs65
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			(_lhs64).ArraySet1(_rhs61, (_lhs65).Int())
			_lhs66.F7 = _rhs62
			_500_v2 = _rhs63
			var _518_v18 _dafny.Array
			_ = _518_v18
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw64
			_518_v18 = _nw64
			var _519_v19 _dafny.MultiSet
			_ = _519_v19
			_519_v19 = _dafny.MultiSetOf(_518_v18)
			var _520_v20 _dafny.MultiSet
			_ = _520_v20
			_520_v20 = _519_v19
			var _521_v21 _dafny.MultiSet
			_ = _521_v21
			_521_v21 = (_520_v20)
			var _source12 _dafny.MultiSet = _520_v20
			_ = _source12
			var _522___mcc_h3 _dafny.MultiSet = _source12
			_ = _522___mcc_h3
			var _523_cf4 _dafny.MultiSet = _522___mcc_h3
			_ = _523_cf4
			var _524_v22 *C0
			_ = _524_v22
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__(!(_499_v1))
			_524_v22 = _nw65
			(globalState).F16 = _498_cf30
			(globalState).F16 = (_this).F28()
			var _525_v23 _dafny.Sequence
			_ = _525_v23
			_525_v23 = _dafny.UnicodeSeqOfUtf8Bytes("am")
			var _526_v24 _dafny.Set
			_ = _526_v24
			_526_v24 = _dafny.SetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-250), (_this).F28()), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_this).F28(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_525_v23, _525_v23), (Companion_Default___.SafeIndex(_498_cf30, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_525_v23, _525_v23)).Cardinality()))).Uint32(), _500_v2)).Cardinality()))
			_526_v24 = ((_526_v24).Difference(_526_v24)).Difference(_526_v24)
			var _527_v25 _dafny.Set
			_ = _527_v25
			_527_v25 = _dafny.SetOf(_498_cf30)
			var _528_v26 _dafny.Sequence
			_ = _528_v26
			_528_v26 = _dafny.SeqOf(_499_v1, true)
			var _529_v27 _dafny.Sequence
			_ = _529_v27
			_529_v27 = _dafny.SeqOf(_dafny.IntOfUint32((_528_v26).Cardinality()), _498_cf30)
			var _530_v28 _dafny.MultiSet
			_ = _530_v28
			_530_v28 = _dafny.MultiSetOf(_498_cf30)
			var _531_v29 _dafny.MultiSet
			_ = _531_v29
			_531_v29 = _dafny.MultiSetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_529_v27).Select((Companion_Default___.SafeIndex((_530_v28).Cardinality(), _dafny.IntOfUint32((_529_v27).Cardinality()))).Uint32()).(_dafny.Int))
			_527_v25 = ((_527_v25).Intersection(_dafny.SetOf(_498_cf30, (_this).F28(), _498_cf30, (_530_v28).Cardinality()))).Union(_527_v25)
		} else if _source11.Is_DC22() {
			var _532___mcc_h1 _dafny.Map = _source11.Get_().(D12_DC22).Cf29
			_ = _532___mcc_h1
			var _533_cf29 _dafny.Map = _532___mcc_h1
			_ = _533_cf29
			var _534_v30 _dafny.Array
			_ = _534_v30
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_11
			var _nw66 _dafny.Array
			_ = _nw66
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw66 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Sequence = func(_535_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("liqnc"), _dafny.UnicodeSeqOfUtf8Bytes("mjrvidhd"))
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw66 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw66).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw66).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_534_v30 = _nw66
			_534_v30 = _534_v30
			var _536_v31 _dafny.Map
			_ = _536_v31
			_536_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F28()), (_this).F28())
			var _537_v32 bool
			_ = _537_v32
			_537_v32 = true
			var _538_v34 D13
			_ = _538_v34
			_538_v34 = Companion_D13_.Create_DC25_(func() _dafny.Map {
				var _coll30 = _dafny.NewMapBuilder()
				_ = _coll30
				for _iter36 := _dafny.Iterate((Companion_Default___.Fm21(globalState)).Elements()); ; {
					_compr_30, _ok36 := _iter36()
					if !_ok36 {
						break
					}
					var _539_v33 _dafny.Int
					_539_v33 = interface{}(_compr_30).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm21(globalState), _539_v33) {
						_coll30.Add(Companion_Default___.SafeModuloInt(_539_v33, (_this).F28()), (_this).F28())
					}
				}
				return _coll30.ToMap()
			}())
			var _540_v35 _dafny.Map
			_ = _540_v35
			_540_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((false) && (_537_v32), (_538_v34).Dtor_cf32())
			_536_v31 = (func() _dafny.Map {
				if (_540_v35).Contains(_537_v32) {
					return (_540_v35).Get(_537_v32).(_dafny.Map)
				}
				return _536_v31
			})()
			var _541_v36 _dafny.Array
			_ = _541_v36
			var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(3))
			_ = _nw67
			_541_v36 = _nw67
			var _542_v37 _dafny.MultiSet
			_ = _542_v37
			_542_v37 = _dafny.MultiSetOf(p0, p0, p0)
			var _543_v38 _dafny.MultiSet
			_ = _543_v38
			_543_v38 = _542_v37
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_541_v36), 0))
			_ = _index75
			(_541_v36).ArraySet1(_543_v38, (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_541_v36), 0))
			_ = _index76
			(_541_v36).ArraySet1(_543_v38, (_index76).Int())
			if true {
				var _544_v39 _dafny.Sequence
				_ = _544_v39
				_544_v39 = _dafny.SeqOf(_dafny.IntOfInt64(-367))
				var _545_v40 _dafny.Sequence
				_ = _545_v40
				_545_v40 = _dafny.SeqOf(_537_v32, _537_v32)
				var _546_v41 _dafny.Sequence
				_ = _546_v41
				_546_v41 = _545_v40
				(globalState).F7 = (func() _dafny.Int {
					if _537_v32 {
						return (_this).F28()
					}
					return Companion_Default___.Fm0((_this).F28(), (_this).Fm17(_537_v32, _544_v39, (_this).F28(), _546_v41, globalState), (_this).F28(), _537_v32, globalState)
				})()
				var _547_v42 _dafny.CodePoint
				_ = _547_v42
				_547_v42 = _dafny.CodePoint('n')
				var _548_v43 _dafny.Map
				_ = _548_v43
				_548_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F28()), _547_v42)
				var _549_v44 _dafny.Sequence
				_ = _549_v44
				_549_v44 = _dafny.UnicodeSeqOfUtf8Bytes("mg")
				_548_v43 = (_548_v43).Update(_dafny.IntOfInt64(366), (_549_v44).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_549_v44).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _550_v45 *C0
				_ = _550_v45
				var _nw68 *C0 = New_C0_()
				_ = _nw68
				_nw68.Ctor__(_537_v32)
				_550_v45 = _nw68
				_545_v40 = _545_v40
				var _551_v46 _dafny.Map
				_ = _551_v46
				_551_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _537_v32)
				_551_v46 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _537_v32)).Merge(_551_v46)
			} else {
				var _552_v47 *C0
				_ = _552_v47
				var _nw69 *C0 = New_C0_()
				_ = _nw69
				_nw69.Ctor__(_537_v32)
				_552_v47 = _nw69
				var _553_v48 _dafny.MultiSet
				_ = _553_v48
				_553_v48 = _dafny.MultiSetOf(_552_v47)
				(globalState).F7 = (((_this).F28()).Times((_this).F28())).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_553_v48).Contains(_552_v47) {
						return (_553_v48).Multiplicity(_552_v47)
					}
					return (_this).F28()
				})(), (_this).F28()))
				var _554_v49 _dafny.Sequence
				_ = _554_v49
				_554_v49 = _dafny.SeqOf(false)
				var _555_v50 _dafny.Sequence
				_ = _555_v50
				_555_v50 = _dafny.UnicodeSeqOfUtf8Bytes("v")
				var _556_v51 _dafny.CodePoint
				_ = _556_v51
				_556_v51 = _dafny.CodePoint('s')
				var _557_v52 D3
				_ = _557_v52
				_557_v52 = Companion_D3_.Create_DC6_(_555_v50, _556_v51, _552_v47.F27)
				var _558_v53 _dafny.Map
				_ = _558_v53
				_558_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_554_v49, (_557_v52).Dtor_cf6())
				Companion_Default___.M0(_552_v47.F27, (_this).F28(), (func() _dafny.Map {
					if _552_v47.F27 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_554_v49, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_554_v49).Cardinality()))).Uint32(), _552_v47.F27), _555_v50)
					}
					return _558_v53
				})(), (_this).F28(), globalState)
				(globalState).F16 = (_this).F28()
				var _559_v54 _dafny.Array
				_ = _559_v54
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
				_ = _nw70
				_559_v54 = _nw70
				_559_v54 = _559_v54
				(globalState).F21 = false
			}
		} else {
			var _560___mcc_h2 D12 = _source11.Get_().(D12_DC24).Cf31
			_ = _560___mcc_h2
			var _561_cf31 D12 = _560___mcc_h2
			_ = _561_cf31
			var _562_v55 _dafny.Map
			_ = _562_v55
			_562_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F28())
			var _563_v56 bool
			_ = _563_v56
			_563_v56 = false
			_562_v55 = (_562_v55).Update(_563_v56, (func() _dafny.Int {
				if _563_v56 {
					return (_this).F28()
				}
				return (_this).F28()
			})())
			(_this).M8(_563_v56, _563_v56, _563_v56, globalState)
			var _564_v57 _dafny.Sequence
			_ = _564_v57
			_564_v57 = _dafny.UnicodeSeqOfUtf8Bytes("npq")
			(globalState).F13 = _dafny.Companion_Sequence_.IsProperPrefixOf(_564_v57, _564_v57)
			var _565_v58 _dafny.CodePoint
			_ = _565_v58
			_565_v58 = _dafny.CodePoint('d')
			var _566_v59 _dafny.Sequence
			_ = _566_v59
			_566_v59 = _dafny.SeqOf(_563_v56)
			var _567_v60 D13
			_ = _567_v60
			_567_v60 = Companion_D13_.Create_DC26_(_564_v57, (Companion_D3_.Create_DC6_(_dafny.Companion_Sequence_.Update(_564_v57, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_564_v57).Cardinality()))).Uint32(), _dafny.CodePoint('e')), _565_v58, _563_v56)).Dtor_cf8(), (_566_v59).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_566_v59).Cardinality()))).Uint32()).(bool))
			if (_567_v60).Dtor_cf35() {
				var _568_v61 _dafny.Sequence
				_ = _568_v61
				_568_v61 = _dafny.SeqOf((_this).F28())
				var _569_v62 _dafny.Sequence
				_ = _569_v62
				_569_v62 = _dafny.SeqOf(_568_v61, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}(func(_570_i2 _dafny.Int) _dafny.Int {
					return (_this).F28()
				})))
				var _571_v63 _dafny.Sequence
				_ = _571_v63
				_571_v63 = _dafny.SeqOf(_569_v62, _569_v62)
				var _572_v64 _dafny.Array
				_ = _572_v64
				var _nwElement0_7 _dafny.Sequence = (_571_v63).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_571_v63).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _nwElement0_7
				var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(7))
				_ = _nw71
				(_nw71).ArraySet1(_nwElement0_7, 0)
				(_nw71).ArraySet1(_569_v62, 1)
				(_nw71).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-179))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_573_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_574_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-177))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg56 _dafny.Int) interface{} {
								return coer56(arg56)
							}
						}((func(_575_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_576_i4 _dafny.Int) _dafny.CodePoint {
								return _575_v58
							}
						})(_573_v58)))).Cardinality())), (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-177))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg57 _dafny.Int) interface{} {
								return coer57(arg57)
							}
						}((func(_577_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_578_i4 _dafny.Int) _dafny.CodePoint {
								return _577_v58
							}
						})(_573_v58)))).Cardinality()))).Cardinality()))).Uint32(), (_this).F28())
					}
				})(_565_v58))), 2)
				(_nw71).ArraySet1(_569_v62, 3)
				(_nw71).ArraySet1(_569_v62, 4)
				(_nw71).ArraySet1(_569_v62, 5)
				(_nw71).ArraySet1((func() _dafny.Sequence {
					if _563_v56 {
						return _569_v62
					}
					return _dafny.Companion_Sequence_.Update(Companion_Default___.Fm22(_563_v56, _563_v56, (_566_v59).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_566_v59).Cardinality()))).Uint32()).(bool), globalState), (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((Companion_Default___.Fm22(_563_v56, _563_v56, (_566_v59).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_566_v59).Cardinality()))).Uint32()).(bool), globalState)).Cardinality()))).Uint32(), _568_v61)
				})(), 6)
				_572_v64 = _nw71
				var _579_v65 _dafny.Map
				_ = _579_v65
				_579_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_566_v59).Cardinality())), _569_v62)
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_572_v64), 0))
				_ = _index77
				(_572_v64).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_569_v62, (func() _dafny.Sequence {
					if (_579_v65).Contains((_this).F28()) {
						return (_579_v65).Get((_this).F28()).(_dafny.Sequence)
					}
					return _569_v62
				})()), (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_572_v64), 0))
				_ = _index78
				(_572_v64).ArraySet1(_569_v62, (_index78).Int())
				(globalState).F13 = _563_v56
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((p0), 0))
				_ = _index79
				(p0).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(976), (_this).F28()), (_index79).Int())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((p0), 0))
				_ = _index80
				(p0).ArraySet1((_dafny.Zero).Minus((_this).F28()), (_index80).Int())
				var _580_v66 _dafny.MultiSet
				_ = _580_v66
				_580_v66 = _dafny.MultiSetOf(_565_v58, (_564_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.IntOfUint32((_564_v57).Cardinality()))).Uint32()).(_dafny.CodePoint), (func() _dafny.CodePoint {
					if _563_v56 {
						return _565_v58
					}
					return _565_v58
				})(), _565_v58)
				(globalState).F16 = (func() _dafny.Int {
					if (_580_v66).Contains((_564_v57).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_564_v57).Cardinality()))).Uint32()).(_dafny.CodePoint)) {
						return (_580_v66).Multiplicity((_564_v57).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_564_v57).Cardinality()))).Uint32()).(_dafny.CodePoint))
					}
					return Companion_Default___.SafeDivisionInt((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_this).F28())
				})()
				(globalState).F21 = !(((_this).F28()).Cmp((_this).F28()) != 0)
			} else {
				(globalState).F2 = _563_v56
				var _581_v67 _dafny.Set
				_ = _581_v67
				_581_v67 = _dafny.SetOf(_565_v58, _565_v58)
				var _582_v68 _dafny.Sequence
				_ = _582_v68
				_582_v68 = _dafny.SeqOf(_581_v67)
				(globalState).F2 = ((_582_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.IntOfUint32((_582_v68).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_dafny.CodePoint('w'))
				var _583_v69 _dafny.Sequence
				_ = _583_v69
				_583_v69 = _dafny.SeqOf((_this).F28())
				var _584_v70 _dafny.Set
				_ = _584_v70
				_584_v70 = _dafny.SetOf(_563_v56)
				var _585_v71 _dafny.Map
				_ = _585_v71
				_585_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _584_v70)
				var _586_v72 _dafny.Array
				_ = _586_v72
				var _nwElement0_8 _dafny.Set = _584_v70
				_ = _nwElement0_8
				var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(9))
				_ = _nw72
				(_nw72).ArraySet1(_nwElement0_8, 0)
				(_nw72).ArraySet1(_dafny.SetOf(_563_v56, _563_v56), 1)
				(_nw72).ArraySet1(_dafny.SetOf(_563_v56), 2)
				(_nw72).ArraySet1((func() _dafny.Set {
					if (_585_v71).Contains((_this).F28()) {
						return (_585_v71).Get((_this).F28()).(_dafny.Set)
					}
					return _584_v70
				})(), 3)
				(_nw72).ArraySet1(_584_v70, 4)
				(_nw72).ArraySet1(_584_v70, 5)
				(_nw72).ArraySet1(_584_v70, 6)
				(_nw72).ArraySet1(_584_v70, 7)
				(_nw72).ArraySet1(_584_v70, 8)
				_586_v72 = _nw72
				var _587_v73 D7
				_ = _587_v73
				_587_v73 = Companion_D7_.Create_DC16_(_586_v72)
				var _588_v74 _dafny.Array
				_ = _588_v74
				var _nwElement0_9 D7 = _587_v73
				_ = _nwElement0_9
				var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(13))
				_ = _nw73
				(_nw73).ArraySet1(_nwElement0_9, 0)
				(_nw73).ArraySet1(Companion_D7_.Create_DC16_(_586_v72), 1)
				(_nw73).ArraySet1(_587_v73, 2)
				(_nw73).ArraySet1(_587_v73, 3)
				(_nw73).ArraySet1(_587_v73, 4)
				(_nw73).ArraySet1(_587_v73, 5)
				(_nw73).ArraySet1(_587_v73, 6)
				(_nw73).ArraySet1(Companion_D7_.Create_DC16_(_586_v72), 7)
				(_nw73).ArraySet1(_587_v73, 8)
				(_nw73).ArraySet1(_587_v73, 9)
				(_nw73).ArraySet1(_587_v73, 10)
				(_nw73).ArraySet1(Companion_D7_.Create_DC16_(_586_v72), 11)
				(_nw73).ArraySet1(_587_v73, 12)
				_588_v74 = _nw73
				var _589_v75 _dafny.Sequence
				_ = _589_v75
				_589_v75 = _dafny.SeqOf(_588_v74, _588_v74)
				var _590_v76 _dafny.Map
				_ = _590_v76
				_590_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_v56, _563_v56)
				var _591_v77 _dafny.MultiSet
				_ = _591_v77
				_591_v77 = _dafny.MultiSetOf((_590_v76).Cardinality(), (_this).F28())
				_583_v69 = _dafny.SeqOf((_this).F28(), _dafny.IntOfUint32((_589_v75).Cardinality()), ((_591_v77).Intersection(_dafny.MultiSetOf((_this).F28(), _dafny.IntOfUint32((_564_v57).Cardinality())))).Cardinality(), (_583_v69).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_583_v69).Cardinality()))).Uint32()).(_dafny.Int))
				var _592_v78 *C0
				_ = _592_v78
				var _nw74 *C0 = New_C0_()
				_ = _nw74
				_nw74.Ctor__(_563_v56)
				_592_v78 = _nw74
				_583_v69 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_566_v59).Cardinality()), (_this).F28()), _dafny.SeqOf(_dafny.IntOfInt64(-630))), _583_v69)
			}
		}
		if true {
			var _593_v79 _dafny.Array
			_ = _593_v79
			var _nw75 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw75
			_593_v79 = _nw75
			var _594_v80 bool
			_ = _594_v80
			_594_v80 = false
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_593_v79), 0))
			_ = _index81
			(_593_v79).ArraySet1(_594_v80, (_index81).Int())
			var _595_v81 _dafny.Sequence
			_ = _595_v81
			_595_v81 = _dafny.UnicodeSeqOfUtf8Bytes("erwiuf")
			var _596_v82 _dafny.Map
			_ = _596_v82
			_596_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_594_v80, _595_v81)
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_593_v79), 0))
			_ = _index82
			(_593_v79).ArraySet1((Companion_Default___.Fm23(_594_v80, (func() _dafny.Sequence {
				if (_596_v82).Contains(true) {
					return (_596_v82).Get(true).(_dafny.Sequence)
				}
				return _595_v81
			})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(50))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}(func(_597_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			})), _594_v80, globalState)).Dtor_cf34(), (_index82).Int())
			var _598_v83 _dafny.Sequence
			_ = _598_v83
			_598_v83 = _dafny.SeqOf(_dafny.IntOfUint32((_595_v81).Cardinality()), (_this).F28())
			var _599_v84 _dafny.Set
			_ = _599_v84
			_599_v84 = _dafny.SetOf((_598_v83).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_598_v83).Cardinality()))).Uint32()).(_dafny.Int), (_this).F28(), (_dafny.Zero).Minus((_this).F28()))
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((p0), 0))
			_ = _index83
			(p0).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_599_v84).Cardinality())), (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((p0), 0))
			_ = _index84
			(p0).ArraySet1(((_this).F28()).Plus((_dafny.Zero).Minus((_this).F28())), (_index84).Int())
			var _600_v85 _dafny.Map
			_ = _600_v85
			_600_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_593_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_593_v79), 0))).Int()).(bool), _593_v79)
			var _601_v86 _dafny.Map
			_ = _601_v86
			_601_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_600_v85, _dafny.IntOfInt64(-210))
			_601_v86 = (_601_v86).Update((_600_v85).Merge(_600_v85), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
			var _602_v87 _dafny.Map
			_ = _602_v87
			_602_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_594_v80, (_dafny.Zero).Minus((_this).F28()))
			var _603_v88 _dafny.Map
			_ = _603_v88
			_603_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_602_v87).Cardinality())
			_603_v88 = (_603_v88).Update((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
			var _604_v89 *C0
			_ = _604_v89
			var _nw76 *C0 = New_C0_()
			_ = _nw76
			_nw76.Ctor__((_593_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_593_v79), 0))).Int()).(bool))
			_604_v89 = _nw76
		} else {
			var _605_v90 bool
			_ = _605_v90
			_605_v90 = true
			var _606_v91 _dafny.Map
			_ = _606_v91
			_606_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v90, (_this).F28())).Cardinality(), _dafny.CodePoint('i'))
			var _607_v92 _dafny.CodePoint
			_ = _607_v92
			_607_v92 = _dafny.CodePoint('a')
			var _608_v93 _dafny.Map
			_ = _608_v93
			_608_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v90, _605_v90)
			var _609_v94 _dafny.Map
			_ = _609_v94
			_609_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cdoyd")).Cardinality()), _608_v93)
			var _610_v95 _dafny.Map
			_ = _610_v95
			_610_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_607_v92, _605_v90)
			var _611_v96 _dafny.MultiSet
			_ = _611_v96
			_611_v96 = _dafny.MultiSetOf(_610_v95)
			(globalState).F21 = !((Companion_Default___.Fm24((func() _dafny.CodePoint {
				if (_606_v91).Contains(_dafny.IntOfInt64(493)) {
					return (_606_v91).Get(_dafny.IntOfInt64(493)).(_dafny.CodePoint)
				}
				return _607_v92
			})(), (_609_v94).Cardinality(), _dafny.IntOfInt64(342), _dafny.IntOfInt64(724), globalState)).IsDisjointFrom(_611_v96))
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((p0), 0))
			_ = _index85
			(p0).ArraySet1((_this).F28(), (_index85).Int())
			var _612_v98 _dafny.Map
			_ = _612_v98
			_612_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll31 = _dafny.NewMapBuilder()
				_ = _coll31
				for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-121), _dafny.IntOfInt64(315))); ; {
					_compr_31, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _613_v97 _dafny.Int
					_613_v97 = interface{}(_compr_31).(_dafny.Int)
					if ((_dafny.IntOfInt64(-121)).Cmp(_613_v97) <= 0) && ((_613_v97).Cmp(_dafny.IntOfInt64(315)) < 0) {
						_coll31.Add((_613_v97).Minus((_this).F28()), false)
					}
				}
				return _coll31.ToMap()
			}()).Cardinality(), _605_v90)
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((p0), 0))
			_ = _index86
			(p0).ArraySet1(Companion_Default___.SafeDivisionInt(((_612_v98).Merge(_612_v98)).Cardinality(), (_this).F28()), (_index86).Int())
			(globalState).F21 = _605_v90
			(globalState).F0 = ((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Cmp((_this).F28()) > 0
			var _614_v99 _dafny.Sequence
			_ = _614_v99
			_614_v99 = _dafny.SeqOf(_605_v90, !(_605_v90))
			var _615_v100 _dafny.Map
			_ = _615_v100
			_615_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v90, _614_v99)
			_615_v100 = (_615_v100).Update(false, _614_v99)
		}
		var _616_v101 _dafny.Sequence
		_ = _616_v101
		_616_v101 = _dafny.UnicodeSeqOfUtf8Bytes("jeovgrwpv")
		var _617_v102 _dafny.MultiSet
		_ = _617_v102
		_617_v102 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(181))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}(func(_618_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), _616_v101, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg60 _dafny.Int) interface{} {
				return coer60(arg60)
			}
		}(func(_619_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		})))
		_617_v102 = ((_617_v102).Difference(_617_v102)).Union(_617_v102)
		var _620_v103 bool
		_ = _620_v103
		_620_v103 = true
		var _621_v104 _dafny.MultiSet
		_ = _621_v104
		_621_v104 = _dafny.MultiSetOf(_620_v103)
		var _622_v105 _dafny.Map
		_ = _622_v105
		_622_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(213), !(false))
		var _623_v106 _dafny.Sequence
		_ = _623_v106
		_623_v106 = _dafny.SeqOf(Companion_Default___.Fm2(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_620_v103, _620_v103), globalState))
		var _624_v107 _dafny.Sequence
		_ = _624_v107
		_624_v107 = _623_v106
		var _hi2 _dafny.Int = ((_this).F28()).Plus((_this).Fm17(_620_v103, _dafny.SeqOf((_this).F28()), _dafny.IntOfUint32((_623_v106).Cardinality()), _624_v107, globalState))
		_ = _hi2
		for _625_i8 := ((func() _dafny.Int {
			if (_621_v104).Contains(true) {
				return (_621_v104).Multiplicity(true)
			}
			return (func() _dafny.Int {
				if (_621_v104).Contains(_620_v103) {
					return (_621_v104).Multiplicity(_620_v103)
				}
				return (_622_v105).Cardinality()
			})()
		})()).Plus(_dafny.IntOfInt64(921)); _625_i8.Cmp(_hi2) < 0; _625_i8 = _625_i8.Plus(_dafny.One) {
			var _626_v108 _dafny.CodePoint
			_ = _626_v108
			_626_v108 = _dafny.CodePoint('o')
			(_this).M8(_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm1(_625_i8, globalState), _626_v108), _620_v103, _620_v103, globalState)
			var _627_v109 _dafny.Map
			_ = _627_v109
			_627_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_620_v103, (_this).F28())
			(globalState).F16 = Companion_Default___.SafeModuloInt(((_627_v109).Update(_620_v103, (_this).F28())).Cardinality(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(762), _625_i8))
			var _628_v110 D6
			_ = _628_v110
			_628_v110 = Companion_D6_.Create_DC13_((_this).F28(), _625_i8)
			var _629_v111 D6
			_ = _629_v111
			_629_v111 = Companion_D6_.Create_DC14_(_628_v110)
			var _source13 D6 = _629_v111
			_ = _source13
			if _source13.Is_DC13() {
				var _630___mcc_h4 _dafny.Int = _source13.Get_().(D6_DC13).Cf19
				_ = _630___mcc_h4
				var _631___mcc_h5 _dafny.Int = _source13.Get_().(D6_DC13).Cf20
				_ = _631___mcc_h5
				var _632_cf20 _dafny.Int = _631___mcc_h5
				_ = _632_cf20
				var _633_cf19 _dafny.Int = _630___mcc_h4
				_ = _633_cf19
				var _634_v112 _dafny.Sequence
				_ = _634_v112
				_634_v112 = _dafny.SeqOf(p0, p0)
				var _635_v113 D6
				_ = _635_v113
				_635_v113 = Companion_D6_.Create_DC12_(_634_v112)
				var _636_v114 _dafny.MultiSet
				_ = _636_v114
				_636_v114 = _dafny.MultiSetOf(_632_cf20)
				var _pat_let_tv11 = _634_v112
				_ = _pat_let_tv11
				var _pat_let_tv12 = _634_v112
				_ = _pat_let_tv12
				var _pat_let_tv13 = p0
				_ = _pat_let_tv13
				var _rhs64 _dafny.Int = _632_cf20
				_ = _rhs64
				var _rhs65 _dafny.Int = _632_cf20
				_ = _rhs65
				var _rhs66 bool = (_620_v103) == (_620_v103)
				_ = _rhs66
				var _rhs67 bool = (_dafny.MultiSetOf(_dafny.IntOfUint32((_616_v101).Cardinality()))).IsProperSubsetOf(_636_v114)
				_ = _rhs67
				var _rhs68 D6 = func(_pat_let8_0 D6) D6 {
					return func(_637_dt__update__tmp_h1 D6) D6 {
						return func(_pat_let9_0 _dafny.Sequence) D6 {
							return func(_638_dt__update_hcf18_h0 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC12_(_638_dt__update_hcf18_h0)
							}(_pat_let9_0)
						}(_dafny.Companion_Sequence_.Update(_pat_let_tv11, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.IntOfUint32((_pat_let_tv12).Cardinality()))).Uint32(), _pat_let_tv13))
					}(_pat_let8_0)
				}(Companion_D6_.Create_DC12_(_634_v112))
				_ = _rhs68
				var _lhs67 *GlobalState = globalState
				_ = _lhs67
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				_lhs67.F7 = _rhs64
				_lhs68.F7 = _rhs65
				_lhs69.F21 = _rhs66
				_lhs70.F21 = _rhs67
				_635_v113 = _rhs68
				(globalState).F13 = !(!(_620_v103) || ((_620_v103) == (_620_v103)))
				(globalState).F2 = _620_v103
				_626_v108 = _626_v108
			} else if _source13.Is_DC12() {
				var _639___mcc_h6 _dafny.Sequence = _source13.Get_().(D6_DC12).Cf18
				_ = _639___mcc_h6
				var _640_cf18 _dafny.Sequence = _639___mcc_h6
				_ = _640_cf18
				var _641_v115 _dafny.Set
				_ = _641_v115
				_641_v115 = _dafny.SetOf(_620_v103)
				var _642_v116 *C0
				_ = _642_v116
				var _nw77 *C0 = New_C0_()
				_ = _nw77
				_nw77.Ctor__((_dafny.SetOf(_620_v103, !(_620_v103), true)).IsSubsetOf(_641_v115))
				_642_v116 = _nw77
				_623_v106 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_642_v116.F27), _623_v106)
				(_this).M8((_621_v104).IsProperSubsetOf(_621_v104), _620_v103, _642_v116.F27, globalState)
				var _643_v117 _dafny.Array
				_ = _643_v117
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_12
				var _nw78 _dafny.Array
				_ = _nw78
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw78 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) bool = (func(_644_i8 _dafny.Int, _645_v105 _dafny.Map, _646_v103 bool) func(_dafny.Int) bool {
						return func(_647_i9 _dafny.Int) bool {
							return (func() bool {
								if (_645_v105).Contains(_644_i8) {
									return (_645_v105).Get(_644_i8).(bool)
								}
								return _646_v103
							})()
						}
					})(_625_i8, _622_v105, _620_v103)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw78 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw78).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw78).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_643_v117 = _nw78
				var _648_v118 _dafny.Map
				_ = _648_v118
				_648_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _623_v106)
				var _649_v119 _dafny.Sequence
				_ = _649_v119
				_649_v119 = _dafny.SeqOf((func() _dafny.Sequence {
					if (_648_v118).Contains((_dafny.Zero).Minus((_this).F28())) {
						return (_648_v118).Get((_dafny.Zero).Minus((_this).F28())).(_dafny.Sequence)
					}
					return _623_v106
				})())
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_643_v117), 0))
				_ = _index87
				(_643_v117).ArraySet1(_dafny.Companion_Sequence_.Contains((_649_v119).Select((Companion_Default___.SafeIndex(_625_i8, _dafny.IntOfUint32((_649_v119).Cardinality()))).Uint32()).(_dafny.Sequence), _620_v103), (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_643_v117), 0))
				_ = _index88
				(_643_v117).ArraySet1((func() bool {
					if _620_v103 {
						return (_620_v103) && (_620_v103)
					}
					return (_625_i8).Cmp((_this).F28()) <= 0
				})(), (_index88).Int())
			} else {
				var _650___mcc_h7 D6 = _source13.Get_().(D6_DC14).Cf21
				_ = _650___mcc_h7
				var _651_cf21 D6 = _650___mcc_h7
				_ = _651_cf21
				(globalState).F2 = (_620_v103) == (_620_v103)
				var _652_v120 D0
				_ = _652_v120
				_652_v120 = Companion_D0_.Create_DC0_(_620_v103)
				var _653_v121 _dafny.Sequence
				_ = _653_v121
				_653_v121 = _dafny.SeqOf((_dafny.MultiSetOf(!((_652_v120).Dtor_cf0()))).Update(!(_620_v103), Companion_Default___.Abs(_625_i8)))
				_621_v104 = (_653_v121).Select((Companion_Default___.SafeIndex(_625_i8, _dafny.IntOfUint32((_653_v121).Cardinality()))).Uint32()).(_dafny.MultiSet)
				var _654_v122 _dafny.Array
				_ = _654_v122
				var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw79
				_654_v122 = _nw79
				_654_v122 = (func() _dafny.Array {
					if ((_this).Fm6(_dafny.IntOfInt64(554), globalState)) && (_620_v103) {
						return _654_v122
					}
					return p0
				})()
				var _655_v123 _dafny.Array
				_ = _655_v123
				var _nw80 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw80
				_655_v123 = _nw80
				var _656_v124 _dafny.Map
				_ = _656_v124
				_656_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_655_v123, _625_i8)
				_656_v124 = (_656_v124).Update(_655_v123, _625_i8)
			}
			var _657_v125 *C0
			_ = _657_v125
			var _nw81 *C0 = New_C0_()
			_ = _nw81
			_nw81.Ctor__(_620_v103)
			_657_v125 = _nw81
		}
		var _658_v126 _dafny.Array
		_ = _658_v126
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
		_ = _nw82
		_658_v126 = _nw82
		var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_658_v126), 0))
		_ = _index89
		(_658_v126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_616_v101, Companion_Default___.Fm1((_this).F28(), globalState)), (_index89).Int())
		var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_658_v126), 0))
		_ = _index90
		(_658_v126).ArraySet1(_616_v101, (_index90).Int())
	}
}
func (_this *C2) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _659_v0 _dafny.Set
		_ = _659_v0
		_659_v0 = _dafny.SetOf(_dafny.IntOfInt64(939), p0, (_this).F28())
		var _660_v1 bool
		_ = _660_v1
		_660_v1 = true
		var _661_v2 _dafny.CodePoint
		_ = _661_v2
		_661_v2 = _dafny.CodePoint('a')
		var _662_v3 _dafny.Sequence
		_ = _662_v3
		_662_v3 = _dafny.UnicodeSeqOfUtf8Bytes("prnrf")
		var _663_v4 _dafny.Sequence
		_ = _663_v4
		_663_v4 = _dafny.SeqOf(!(_660_v1), Companion_Default___.Fm2(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v1, _660_v1), globalState))
		var _664_v5 D1
		_ = _664_v5
		_664_v5 = Companion_D1_.Create_DC3_(_660_v1)
		var _665_v6 _dafny.Set
		_ = _665_v6
		_665_v6 = _dafny.SetOf(true, _660_v1, _660_v1)
		var _666_v7 _dafny.Array
		_ = _666_v7
		var _nwElement0_10 bool = !(((_659_v0).Cardinality()).Cmp((_this).F28()) <= 0)
		_ = _nwElement0_10
		var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(16))
		_ = _nw83
		(_nw83).ArraySet1(_nwElement0_10, 0)
		(_nw83).ArraySet1(_660_v1, 1)
		(_nw83).ArraySet1(!_dafny.Companion_Sequence_.Contains(_662_v3, _661_v2), 2)
		(_nw83).ArraySet1((func() bool {
			if _660_v1 {
				return (_663_v4).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_663_v4).Cardinality()))).Uint32()).(bool)
			}
			return (_664_v5).Dtor_cf3()
		})(), 3)
		(_nw83).ArraySet1((_660_v1) == (_660_v1), 4)
		(_nw83).ArraySet1(_660_v1, 5)
		(_nw83).ArraySet1(_660_v1, 6)
		(_nw83).ArraySet1((_this).Fm6((_665_v6).Cardinality(), globalState), 7)
		(_nw83).ArraySet1(_660_v1, 8)
		(_nw83).ArraySet1(!(true) || (_660_v1), 9)
		(_nw83).ArraySet1(_660_v1, 10)
		(_nw83).ArraySet1(_660_v1, 11)
		(_nw83).ArraySet1(_660_v1, 12)
		(_nw83).ArraySet1(_660_v1, 13)
		(_nw83).ArraySet1(_660_v1, 14)
		(_nw83).ArraySet1(_660_v1, 15)
		_666_v7 = _nw83
		for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_666_v7), 0))); ; {
			_guard_loop_5, _ok38 := _iter38()
			if !_ok38 {
				break
			}
			var _667_i0 _dafny.Int
			_667_i0 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_667_i0).Sign() != -1) && ((_667_i0).Cmp(_dafny.ArrayLen((_666_v7), 0)) < 0)) {
				(_666_v7).ArraySet1(_660_v1, (_667_i0).Int())
			}
		}
		var _668_v8 _dafny.Map
		_ = _668_v8
		_668_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
		var _669_v9 _dafny.Array
		_ = _669_v9
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_13
		var _nw84 _dafny.Array
		_ = _nw84
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw84 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.CodePoint = (func(_670_v3 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
				return func(_671_i1 _dafny.Int) _dafny.CodePoint {
					return (_670_v3).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_670_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_662_v3)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw84 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw84).ArraySet1CodePoint(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw84).ArraySet1CodePoint(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_669_v9 = _nw84
		var _672_v10 _dafny.Map
		_ = _672_v10
		_672_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _669_v9)
		_668_v8 = (_668_v8).Update(p0, ((_672_v10).Cardinality()).Plus((_this).F28()))
		if _660_v1 {
			var _673_v11 _dafny.Array
			_ = _673_v11
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_14
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = (func(_674_v1 bool, _675_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_676_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_676_i2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_674_v1, _675_p0)).Cardinality())
					}
				})(_660_v1, p0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw85 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw85).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw85).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_673_v11 = _nw85
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_673_v11), 0))
			_ = _index91
			(_673_v11).ArraySet1(p0, (_index91).Int())
			var _677_v12 _dafny.Sequence
			_ = _677_v12
			_677_v12 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-535))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}((func(_678_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_679_i3 _dafny.Int) _dafny.CodePoint {
					return _678_v2
				}
			})(_661_v2))), _dafny.UnicodeSeqOfUtf8Bytes("lwkghweu"), _662_v3)
			var _680_v13 _dafny.Sequence
			_ = _680_v13
			_680_v13 = _dafny.SeqOf((_677_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.IntOfUint32((_677_v12).Cardinality()))).Uint32()).(_dafny.Sequence), _662_v3, _662_v3, _dafny.UnicodeSeqOfUtf8Bytes("hpd"))
			var _681_v14 _dafny.Map
			_ = _681_v14
			_681_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v4, _666_v7)
			var _682_v15 _dafny.Map
			_ = _682_v15
			_682_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v8, _661_v2)
			var _683_v17 _dafny.Set
			_ = _683_v17
			_683_v17 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus((_dafny.MultiSetOf(_660_v1, _660_v1, false, _660_v1, true)).Cardinality())), _668_v8)
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_673_v11), 0))
			_ = _index92
			var _rhs69 _dafny.Int = Companion_Default___.SafeModuloInt(p0, (_this).F28())
			_ = _rhs69
			var _rhs70 bool = (func() _dafny.Set {
				var _coll32 = _dafny.NewBuilder()
				_ = _coll32
				for _iter39 := _dafny.Iterate((_682_v15).Keys().Elements()); ; {
					_compr_32, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _684_v16 _dafny.Map
					_684_v16 = interface{}(_compr_32).(_dafny.Map)
					if (_682_v15).Contains(_684_v16) {
						_coll32.Add(_684_v16)
					}
				}
				return _coll32.ToSet()
			}()).IsDisjointFrom(_683_v17)
			_ = _rhs70
			var _rhs71 _dafny.Sequence = _677_v12
			_ = _rhs71
			var _rhs72 _dafny.Map = ((_681_v14).Merge(_681_v14)).Merge(_681_v14)
			_ = _rhs72
			var _lhs71 _dafny.Array = _673_v11
			_ = _lhs71
			var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_673_v11), 0))
			_ = _lhs72
			var _lhs73 *GlobalState = globalState
			_ = _lhs73
			(_lhs71).ArraySet1(_rhs69, (_lhs72).Int())
			_lhs73.F21 = _rhs70
			_677_v12 = _rhs71
			_681_v14 = _rhs72
			_660_v1 = false
			(globalState).F13 = !(_660_v1)
			(globalState).F21 = _660_v1
			var _685_v18 _dafny.Array
			_ = _685_v18
			var _nwElement0_11 _dafny.Sequence = _662_v3
			_ = _nwElement0_11
			var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(10))
			_ = _nw86
			(_nw86).ArraySet1(_nwElement0_11, 0)
			(_nw86).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((_662_v3).Cardinality()), globalState), 1)
			(_nw86).ArraySet1(_662_v3, 2)
			(_nw86).ArraySet1((Companion_Default___.Fm25(_660_v1, (_this).F28(), _660_v1, _dafny.IntOfInt64(842), globalState)).Dtor_cf6(), 3)
			(_nw86).ArraySet1(_662_v3, 4)
			(_nw86).ArraySet1(_662_v3, 5)
			(_nw86).ArraySet1(_662_v3, 6)
			(_nw86).ArraySet1(_dafny.Companion_Sequence_.Update(_662_v3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_662_v3).Cardinality()))).Uint32(), _661_v2), 7)
			(_nw86).ArraySet1(_662_v3, 8)
			(_nw86).ArraySet1(_662_v3, 9)
			_685_v18 = _nw86
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_685_v18), 0))
			_ = _index93
			(_685_v18).ArraySet1(_662_v3, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_685_v18), 0))
			_ = _index94
			(_685_v18).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_662_v3, _662_v3), (_index94).Int())
		} else {
			var _686_v19 _dafny.Array
			_ = _686_v19
			var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw87
			_686_v19 = _nw87
			var _687_v20 _dafny.MultiSet
			_ = _687_v20
			_687_v20 = _dafny.MultiSetOf(_661_v2)
			var _688_v21 _dafny.Map
			_ = _688_v21
			_688_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v1, (_687_v20).Cardinality())
			var _689_v22 _dafny.Map
			_ = _689_v22
			_689_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_688_v21, _660_v1)
			var _690_v23 _dafny.Map
			_ = _690_v23
			_690_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v1, _686_v19)
			var _691_v24 _dafny.Array
			_ = _691_v24
			var _nwElement0_12 _dafny.Array = _686_v19
			_ = _nwElement0_12
			var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(12))
			_ = _nw88
			(_nw88).ArraySet1(_nwElement0_12, 0)
			(_nw88).ArraySet1(_686_v19, 1)
			(_nw88).ArraySet1(_686_v19, 2)
			(_nw88).ArraySet1(_686_v19, 3)
			(_nw88).ArraySet1(_686_v19, 4)
			(_nw88).ArraySet1(_686_v19, 5)
			(_nw88).ArraySet1((func() _dafny.Array {
				if (_this).Fm7(false, p0, (func() bool {
					if (_689_v22).Contains(_688_v21) {
						return (_689_v22).Get(_688_v21).(bool)
					}
					return _660_v1
				})(), globalState) {
					return _686_v19
				}
				return _686_v19
			})(), 6)
			(_nw88).ArraySet1(_686_v19, 7)
			(_nw88).ArraySet1(_686_v19, 8)
			(_nw88).ArraySet1((func() _dafny.Array {
				if (_690_v23).Contains(_660_v1) {
					return (_690_v23).Get(_660_v1).(_dafny.Array)
				}
				return _686_v19
			})(), 9)
			(_nw88).ArraySet1(_686_v19, 10)
			(_nw88).ArraySet1(_686_v19, 11)
			_691_v24 = _nw88
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))
			_ = _index95
			(_691_v24).ArraySet1(_686_v19, (_index95).Int())
			var _692_v25 _dafny.Sequence
			_ = _692_v25
			_692_v25 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_663_v4))
			var _693_v26 _dafny.Sequence
			_ = _693_v26
			_693_v26 = _dafny.SeqOf(_dafny.IntOfInt64(445))
			var _694_v27 _dafny.Map
			_ = _694_v27
			_694_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(553), _693_v26)
			var _695_v28 _dafny.MultiSet
			_ = _695_v28
			_695_v28 = _dafny.MultiSetOf(true)
			var _696_v29 D4
			_ = _696_v29
			_696_v29 = Companion_D4_.Create_DC10_(p0, (_this).F28(), (_this).Fm7(_660_v1, (_this).F28(), _660_v1, globalState), _660_v1, p0)
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))
			_ = _index96
			var _rhs73 _dafny.Array = (func() _dafny.Array {
				if _660_v1 {
					return (func() _dafny.Array {
						if _660_v1 {
							return _686_v19
						}
						return _686_v19
					})()
				}
				return _686_v19
			})()
			_ = _rhs73
			var _rhs74 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_692_v25, (Companion_Default___.SafeIndex((_694_v27).Cardinality(), _dafny.IntOfUint32((_692_v25).Cardinality()))).Uint32(), (_dafny.MultiSetFromSeq(_663_v4)).Difference(_695_v28)), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_692_v25, (Companion_Default___.SafeIndex((_694_v27).Cardinality(), _dafny.IntOfUint32((_692_v25).Cardinality()))).Uint32(), (_dafny.MultiSetFromSeq(_663_v4)).Difference(_695_v28))).Cardinality()))).Uint32(), (_695_v28).Update(_660_v1, Companion_Default___.Abs((_this).F28())))
			_ = _rhs74
			var _rhs75 _dafny.Int = (_696_v29).Dtor_cf16()
			_ = _rhs75
			var _rhs76 bool = _660_v1
			_ = _rhs76
			var _lhs74 _dafny.Array = _691_v24
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))
			_ = _lhs75
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			(_lhs74).ArraySet1(_rhs73, (_lhs75).Int())
			_692_v25 = _rhs74
			_lhs76.F7 = _rhs75
			_lhs77.F0 = _rhs76
			var _697_v30 *C0
			_ = _697_v30
			var _nw89 *C0 = New_C0_()
			_ = _nw89
			_nw89.Ctor__(true)
			_697_v30 = _nw89
			var _698_v31 *C0
			_ = _698_v31
			var _nw90 *C0 = New_C0_()
			_ = _nw90
			_nw90.Ctor__(((_this).F28()).Cmp(p0) >= 0)
			_698_v31 = _nw90
			if _660_v1 {
				var _699_v32 _dafny.Map
				_ = _699_v32
				_699_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cmp((_this).F28()) == 0, (_this).Fm6(_dafny.IntOfUint32((_662_v3).Cardinality()), globalState))
				_699_v32 = (_699_v32).Update((_660_v1) || (_698_v31.F27), _698_v31.F27)
				r1 = _698_v31.F27
				var _700_v33 _dafny.Array
				_ = _700_v33
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw91
				_700_v33 = _nw91
				var _701_v34 _dafny.Map
				_ = _701_v34
				_701_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm6((_this).F28(), globalState))
				var _702_v37 _dafny.Array
				_ = _702_v37
				var _nwElement0_13 _dafny.Set = _659_v0
				_ = _nwElement0_13
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(7))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_13, 0)
				(_nw92).ArraySet1(_dafny.SetOf((_701_v34).Cardinality()), 1)
				(_nw92).ArraySet1(_659_v0, 2)
				(_nw92).ArraySet1(func() _dafny.Set {
					var _coll33 = _dafny.NewBuilder()
					_ = _coll33
					for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(630), _dafny.IntOfInt64(607))); ; {
						_compr_33, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _703_v35 _dafny.Int
						_703_v35 = interface{}(_compr_33).(_dafny.Int)
						if ((_dafny.IntOfInt64(630)).Cmp(_703_v35) <= 0) && ((_703_v35).Cmp(_dafny.IntOfInt64(607)) < 0) {
							_coll33.Add((_703_v35).Minus((_this).F28()))
						}
					}
					return _coll33.ToSet()
				}(), 3)
				(_nw92).ArraySet1(_659_v0, 4)
				(_nw92).ArraySet1(_659_v0, 5)
				(_nw92).ArraySet1(func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(910), _dafny.IntOfInt64(974))); ; {
						_compr_34, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _704_v36 _dafny.Int
						_704_v36 = interface{}(_compr_34).(_dafny.Int)
						if ((_dafny.IntOfInt64(910)).Cmp(_704_v36) <= 0) && ((_704_v36).Cmp(_dafny.IntOfInt64(974)) < 0) {
							_coll34.Add((_704_v36).Minus((_this).F28()))
						}
					}
					return _coll34.ToSet()
				}(), 6)
				_702_v37 = _nw92
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_700_v33), 0))
				_ = _index97
				(_700_v33).ArraySet1(_702_v37, (_index97).Int())
				var _705_v38 _dafny.Sequence
				_ = _705_v38
				_705_v38 = _dafny.SeqOf(_701_v34, _701_v34)
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_700_v33), 0))
				_ = _index98
				var _rhs77 _dafny.Array = _666_v7
				_ = _rhs77
				var _rhs78 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_705_v38, _705_v38)).Cardinality())
				_ = _rhs78
				var _rhs79 _dafny.Array = (func() _dafny.Array {
					if _698_v31.F27 {
						return _702_v37
					}
					return _702_v37
				})()
				_ = _rhs79
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				var _lhs79 _dafny.Array = _700_v33
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_700_v33), 0))
				_ = _lhs80
				_666_v7 = _rhs77
				_lhs78.F7 = _rhs78
				(_lhs79).ArraySet1(_rhs79, (_lhs80).Int())
				_666_v7 = _666_v7
				var _706_v39 _dafny.Map
				_ = _706_v39
				_706_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_698_v31.F27), (_this).F28())
				_706_v39 = (_706_v39).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_698_v31.F27), _663_v4), (_this).F28())
			} else {
				(globalState).F2 = true
				_693_v26 = _693_v26
				(globalState).F7 = _dafny.IntOfInt64(519)
				_666_v7 = _666_v7
				var _707_v40 _dafny.Map
				_ = _707_v40
				_707_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_697_v30.F27, _661_v2)
				var _708_v41 _dafny.Map
				_ = _708_v41
				_708_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_662_v3, (_707_v40).Cardinality())
				_708_v41 = (_708_v41).Update(_dafny.UnicodeSeqOfUtf8Bytes("vvrdpbh"), p0)
			}
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_691_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))).Int()))
			_ = _arr0
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_dafny.ArrayCastTo((_691_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))).Int()))), 0))
			_ = _index99
			(_arr0).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F28(), p0), (_index99).Int())
			var _709_v42 D6
			_ = _709_v42
			_709_v42 = Companion_D6_.Create_DC13_((_this).F28(), (_dafny.MultiSetOf(_698_v31.F27)).Cardinality())
			var _710_v43 D6
			_ = _710_v43
			_710_v43 = Companion_D6_.Create_DC14_(_709_v42)
			var _711_v44 D6
			_ = _711_v44
			_711_v44 = Companion_D6_.Create_DC14_(_709_v42)
			var _712_v45 _dafny.Sequence
			_ = _712_v45
			_712_v45 = _dafny.SeqOf(_662_v3)
			var _713_v46 _dafny.Sequence
			_ = _713_v46
			_713_v46 = _dafny.SeqOf((_712_v45).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_712_v45).Cardinality()))).Uint32()).(_dafny.Sequence), _662_v3, _662_v3, _dafny.UnicodeSeqOfUtf8Bytes("ikcvsk"))
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_691_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))).Int()))
			_ = _arr1
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_dafny.ArrayCastTo((_691_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))).Int()))), 0))
			_ = _index100
			var _rhs80 _dafny.Int = (_dafny.Zero).Minus((_this).F28())
			_ = _rhs80
			var _rhs81 D6 = _711_v44
			_ = _rhs81
			var _rhs82 _dafny.Int = Companion_Default___.SafeModuloInt(p0, p0)
			_ = _rhs82
			var _rhs83 _dafny.Sequence = (_713_v46).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_713_v46).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs83
			var _lhs81 _dafny.Array = _dafny.ArrayCastTo((_691_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))).Int()))
			_ = _lhs81
			var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_dafny.ArrayCastTo((_691_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_691_v24), 0))).Int()))), 0))
			_ = _lhs82
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			var _lhs84 *GlobalState = globalState
			_ = _lhs84
			(_lhs81).ArraySet1(_rhs80, (_lhs82).Int())
			_711_v44 = _rhs81
			_lhs83.F7 = _rhs82
			_lhs84.F20 = _rhs83
		}
		_668_v8 = (_668_v8).Update((_this).F28(), p0)
		var _hi3 _dafny.Int = (_this).F28()
		_ = _hi3
		for _714_i4 := p0; _714_i4.Cmp(_hi3) < 0; _714_i4 = _714_i4.Plus(_dafny.One) {
			var _715_v47 _dafny.Array
			_ = _715_v47
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
			_ = _nw93
			_715_v47 = _nw93
			var _716_v48 _dafny.Map
			_ = _716_v48
			_716_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _660_v1)
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_715_v47), 0))
			_ = _index101
			(_715_v47).ArraySet1((_716_v48).Merge(_716_v48), (_index101).Int())
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_715_v47), 0))
			_ = _index102
			(_715_v47).ArraySet1((_716_v48).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_714_i4, _660_v1)), (_index102).Int())
			(globalState).F21 = (false) && (_660_v1)
			var _717_v49 *C0
			_ = _717_v49
			var _nw94 *C0 = New_C0_()
			_ = _nw94
			_nw94.Ctor__(_660_v1)
			_717_v49 = _nw94
			(globalState).F19 = _dafny.Companion_Sequence_.Concatenate(_662_v3, _662_v3)
		}
		var _718_v50 _dafny.Map
		_ = _718_v50
		_718_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
		if (_659_v0).IsProperSubsetOf(Companion_Default___.Fm26(_661_v2, _718_v50, globalState)) {
			var _719_v51 _dafny.Map
			_ = _719_v51
			_719_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(_662_v3, _661_v2), Companion_Default___.Fm27(!(_660_v1), _660_v1, globalState))
			_719_v51 = (_719_v51).Update(_660_v1, (func() _dafny.CodePoint {
				if _660_v1 {
					return _dafny.CodePoint('q')
				}
				return _661_v2
			})())
			(globalState).F8 = _662_v3
			var _720_v52 _dafny.MultiSet
			_ = _720_v52
			_720_v52 = _dafny.MultiSetOf(_660_v1, _660_v1)
			var _721_v53 *C0
			_ = _721_v53
			var _nw95 *C0 = New_C0_()
			_ = _nw95
			_nw95.Ctor__(((_this).F28()).Cmp((_720_v52).Cardinality()) == 0)
			_721_v53 = _nw95
			var _722_v54 *C0
			_ = _722_v54
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__(!((_this).Fm6((_this).F28(), globalState)))
			_722_v54 = _nw96
			var _723_v55 *C0
			_ = _723_v55
			var _nw97 *C0 = New_C0_()
			_ = _nw97
			_nw97.Ctor__(_721_v53.F27)
			_723_v55 = _nw97
		} else {
			if _660_v1 {
				(globalState).F16 = (_dafny.Zero).Minus(p0)
				var _724_v56 _dafny.Array
				_ = _724_v56
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw98
				_724_v56 = _nw98
				var _725_v57 _dafny.Map
				_ = _725_v57
				_725_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_724_v56, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _660_v1 {
						return _662_v3
					}
					return _662_v3
				})()).Cardinality()))
				_725_v57 = (_725_v57).Update(_724_v56, (_dafny.IntOfInt64(-586)).Times(p0))
				var _726_v58 bool
				_ = _726_v58
				var _727_v59 bool
				_ = _727_v59
				var _728_v60 bool
				_ = _728_v60
				var _729_v61 bool
				_ = _729_v61
				var _out0 bool
				_ = _out0
				var _out1 bool
				_ = _out1
				var _out2 bool
				_ = _out2
				var _out3 bool
				_ = _out3
				_out0, _out1, _out2, _out3 = (_this).M9(globalState)
				_726_v58 = _out0
				_727_v59 = _out1
				_728_v60 = _out2
				_729_v61 = _out3
				_668_v8 = ((_668_v8).Merge(func() _dafny.Map {
					var _coll35 = _dafny.NewMapBuilder()
					_ = _coll35
					for _iter42 := _dafny.Iterate((_659_v0).Elements()); ; {
						_compr_35, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _730_v62 _dafny.Int
						_730_v62 = interface{}(_compr_35).(_dafny.Int)
						if (_659_v0).Contains(_730_v62) {
							_coll35.Add((_730_v62).Minus((_this).F28()), (func() _dafny.Int {
								if (_668_v8).Contains((_dafny.Zero).Minus((_this).F28())) {
									return (_668_v8).Get((_dafny.Zero).Minus((_this).F28())).(_dafny.Int)
								}
								return (_this).F28()
							})())
						}
					}
					return _coll35.ToMap()
				}())).Merge((func() _dafny.Map {
					if false {
						return _668_v8
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(493), _dafny.IntOfUint32((_662_v3).Cardinality()))
				})())
				(globalState).F13 = Companion_Default___.Fm2(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_v58, _727_v59), globalState)
			} else {
				var _731_v63 D4
				_ = _731_v63
				_731_v63 = Companion_D4_.Create_DC7_((_this).F28())
				var _732_v64 _dafny.MultiSet
				_ = _732_v64
				_732_v64 = _dafny.MultiSetOf(!(_660_v1))
				var _733_v65 _dafny.Array
				_ = _733_v65
				var _nwElement0_14 _dafny.Int = p0
				_ = _nwElement0_14
				var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(9))
				_ = _nw99
				(_nw99).ArraySet1(_nwElement0_14, 0)
				(_nw99).ArraySet1((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_731_v63).Dtor_cf9(), true)).Cardinality()), 1)
				(_nw99).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(986))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_734_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_735_i5 _dafny.Int) _dafny.CodePoint {
						return _734_v2
					}
				})(_661_v2)))).Cardinality())), 2)
				(_nw99).ArraySet1((_this).F28(), 3)
				(_nw99).ArraySet1(p0, 4)
				(_nw99).ArraySet1(p0, 5)
				(_nw99).ArraySet1((_732_v64).Cardinality(), 6)
				(_nw99).ArraySet1(p0, 7)
				(_nw99).ArraySet1((_this).F28(), 8)
				_733_v65 = _nw99
				r0 = (func() _dafny.Array {
					if _660_v1 {
						return _733_v65
					}
					return _733_v65
				})()
				var _736_v66 _dafny.Map
				_ = _736_v66
				_736_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v1, _660_v1)
				_736_v66 = (_736_v66).Update(!(_660_v1) || (_660_v1), _660_v1)
				(globalState).F7 = (_this).F28()
				var _737_v67 *C0
				_ = _737_v67
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__(_660_v1)
				_737_v67 = _nw100
				var _738_v68 _dafny.Sequence
				_ = _738_v68
				_738_v68 = _dafny.SeqOf(_dafny.IntOfInt64(977), (_this).F28())
				var _739_v69 _dafny.Sequence
				_ = _739_v69
				_739_v69 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(989))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}(func(_740_i8 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("efk")).Cardinality())
				})), (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(989))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}(func(_741_i8 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("efk")).Cardinality())
				}))).Cardinality()))).Uint32(), (_this).F28()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(396))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_742_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_743_i9 _dafny.Int) _dafny.Int {
						return _742_p0
					}
				})(p0))))
				var _744_v70 _dafny.Map
				_ = _744_v70
				_744_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v1, (_739_v69).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_738_v68).Cardinality()), _dafny.IntOfUint32((_739_v69).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _745_v71 _dafny.Array
				_ = _745_v71
				var _nwElement0_15 _dafny.Sequence = _738_v68
				_ = _nwElement0_15
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(22))
				_ = _nw101
				(_nw101).ArraySet1(_nwElement0_15, 0)
				(_nw101).ArraySet1(_738_v68, 1)
				(_nw101).ArraySet1((func() _dafny.Sequence {
					if _660_v1 {
						return _738_v68
					}
					return _738_v68
				})(), 2)
				(_nw101).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_738_v68, _738_v68), 3)
				(_nw101).ArraySet1(_738_v68, 4)
				(_nw101).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_738_v68, _738_v68), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_738_v68, _738_v68)).Cardinality()))).Uint32(), (func() _dafny.Int {
					if (_668_v8).Contains(p0) {
						return (_668_v8).Get(p0).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_662_v3).Cardinality())
				})()), 5)
				(_nw101).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_746_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_747_i6 _dafny.Int) _dafny.Int {
						return _746_p0
					}
				})(p0))), 6)
				(_nw101).ArraySet1(_738_v68, 7)
				(_nw101).ArraySet1(_738_v68, 8)
				(_nw101).ArraySet1(_738_v68, 9)
				(_nw101).ArraySet1(_738_v68, 10)
				(_nw101).ArraySet1(_dafny.Companion_Sequence_.Update(_738_v68, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_738_v68).Cardinality()))).Uint32(), (_this).F28()), 11)
				(_nw101).ArraySet1(_738_v68, 12)
				(_nw101).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_738_v68, _738_v68), 13)
				(_nw101).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(471))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_748_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_749_i7 _dafny.Int) _dafny.Int {
						return _748_p0
					}
				})(p0))), _738_v68), 14)
				(_nw101).ArraySet1(_738_v68, 15)
				(_nw101).ArraySet1(_738_v68, 16)
				(_nw101).ArraySet1((func() _dafny.Sequence {
					if (_744_v70).Contains(Companion_Default___.Fm2(_736_v66, globalState)) {
						return (_744_v70).Get(Companion_Default___.Fm2(_736_v66, globalState)).(_dafny.Sequence)
					}
					return _738_v68
				})(), 17)
				(_nw101).ArraySet1(Companion_Default___.Fm21(globalState), 18)
				(_nw101).ArraySet1((func() _dafny.Sequence {
					if _660_v1 {
						return _738_v68
					}
					return _738_v68
				})(), 19)
				(_nw101).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_738_v68, _738_v68), 20)
				(_nw101).ArraySet1(_738_v68, 21)
				_745_v71 = _nw101
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_745_v71), 0))
				_ = _index103
				(_745_v71).ArraySet1((_739_v69).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F28()), _dafny.IntOfUint32((_739_v69).Cardinality()))).Uint32()).(_dafny.Sequence), (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_745_v71), 0))
				_ = _index104
				(_745_v71).ArraySet1(_738_v68, (_index104).Int())
			}
			var _750_v72 _dafny.Array
			_ = _750_v72
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_15
			var _nw102 _dafny.Array
			_ = _nw102
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw102 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Int = func(_751_i10 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_751_i10, (_this).F28())
				}
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw102 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw102).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw102).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_750_v72 = _nw102
			var _752_v73 _dafny.Map
			_ = _752_v73
			_752_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _750_v72)
			r0 = (func() _dafny.Array {
				if (_752_v73).Contains(p0) {
					return (_752_v73).Get(p0).(_dafny.Array)
				}
				return _750_v72
			})()
			var _753_v74 bool
			_ = _753_v74
			var _754_v75 bool
			_ = _754_v75
			var _755_v76 bool
			_ = _755_v76
			var _756_v77 bool
			_ = _756_v77
			var _out4 bool
			_ = _out4
			var _out5 bool
			_ = _out5
			var _out6 bool
			_ = _out6
			var _out7 bool
			_ = _out7
			_out4, _out5, _out6, _out7 = (_this).M9(globalState)
			_753_v74 = _out4
			_754_v75 = _out5
			_755_v76 = _out6
			_756_v77 = _out7
			if (_663_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_663_v4).Cardinality()))).Uint32()).(bool) {
				_661_v2 = _661_v2
				var _757_v78 *C0
				_ = _757_v78
				var _nw103 *C0 = New_C0_()
				_ = _nw103
				_nw103.Ctor__(_756_v77)
				_757_v78 = _nw103
				_663_v4 = _dafny.SeqOf((_this).Fm6(_dafny.IntOfInt64(-80), globalState), _660_v1)
				var _758_v79 _dafny.Array
				_ = _758_v79
				var _nwElement0_16 _dafny.Set = (_659_v0).Difference(_659_v0)
				_ = _nwElement0_16
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(2))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_16, 0)
				(_nw104).ArraySet1(_659_v0, 1)
				_758_v79 = _nw104
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_758_v79), 0))
				_ = _index105
				(_758_v79).ArraySet1(_659_v0, (_index105).Int())
				var _759_v80 _dafny.Array
				_ = _759_v80
				var _nw105 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
				_ = _nw105
				_759_v80 = _nw105
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_759_v80), 0))
				_ = _index106
				(_759_v80).ArraySet1(_750_v72, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_758_v79), 0))
				_ = _index107
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_759_v80), 0))
				_ = _index108
				var _rhs84 _dafny.Set = _659_v0
				_ = _rhs84
				var _rhs85 _dafny.Array = _750_v72
				_ = _rhs85
				var _rhs86 bool = _755_v76
				_ = _rhs86
				var _lhs85 _dafny.Array = _758_v79
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_758_v79), 0))
				_ = _lhs86
				var _lhs87 _dafny.Array = _759_v80
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_759_v80), 0))
				_ = _lhs88
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				(_lhs85).ArraySet1(_rhs84, (_lhs86).Int())
				(_lhs87).ArraySet1(_rhs85, (_lhs88).Int())
				_lhs89.F2 = _rhs86
				var _760_v81 _dafny.Map
				_ = _760_v81
				_760_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v1, p0)
				var _761_v82 _dafny.MultiSet
				_ = _761_v82
				_761_v82 = _dafny.MultiSetOf((func() _dafny.Int {
					if (_760_v81).Contains(true) {
						return (_760_v81).Get(true).(_dafny.Int)
					}
					return p0
				})())
				var _762_v83 _dafny.MultiSet
				_ = _762_v83
				_762_v83 = _dafny.MultiSetOf((_761_v82).Cardinality(), (_dafny.Zero).Minus(p0))
				var _763_v84 _dafny.Sequence
				_ = _763_v84
				_763_v84 = _dafny.SeqOf(_750_v72)
				var _764_v85 _dafny.Sequence
				_ = _764_v85
				_764_v85 = _dafny.SeqOf((_763_v84).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_763_v84).Cardinality()))).Uint32()).(_dafny.Array))
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_759_v80), 0))
				_ = _index109
				var _rhs87 _dafny.Array = _666_v7
				_ = _rhs87
				var _rhs88 bool = (Companion_Default___.Fm28((_758_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_758_v79), 0))).Int()).(_dafny.Set), p0, _753_v74, true, globalState)).IsSubsetOf((_761_v82).Update(_dafny.IntOfUint32((_662_v3).Cardinality()), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xnvuuooh")).Cardinality()))))
				_ = _rhs88
				var _rhs89 _dafny.Array = (_763_v84).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F28()), _dafny.IntOfUint32((_763_v84).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs89
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				var _lhs91 _dafny.Array = _759_v80
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_759_v80), 0))
				_ = _lhs92
				_666_v7 = _rhs87
				_lhs90.F2 = _rhs88
				(_lhs91).ArraySet1(_rhs89, (_lhs92).Int())
			} else {
				(globalState).F13 = (func() bool {
					if !(!(_755_v76)) {
						return (_dafny.IntOfInt64(479)).Cmp((_dafny.Zero).Minus((_this).F28())) >= 0
					}
					return _754_v75
				})()
				var _765_v86 _dafny.Sequence
				_ = _765_v86
				_765_v86 = _dafny.SeqOf(_718_v50, _718_v50)
				var _766_v87 T1
				_ = _766_v87
				var _nw106 *C1 = New_C1_()
				_ = _nw106
				_nw106.Ctor__(p0, _765_v86)
				_766_v87 = _nw106
				var _767_v88 _dafny.Set
				_ = _767_v88
				_767_v88 = _dafny.SetOf(_766_v87)
				var _768_v89 _dafny.Sequence
				_ = _768_v89
				_768_v89 = _dafny.SeqOf(p0)
				var _769_v90 _dafny.Sequence
				_ = _769_v90
				_769_v90 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_663_v4, (Companion_Default___.SafeIndex((_668_v8).Cardinality(), _dafny.IntOfUint32((_663_v4).Cardinality()))).Uint32(), _753_v74), (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_663_v4, (Companion_Default___.SafeIndex((_668_v8).Cardinality(), _dafny.IntOfUint32((_663_v4).Cardinality()))).Uint32(), _753_v74)).Cardinality()))).Uint32(), _754_v75)
				var _770_v91 _dafny.MultiSet
				_ = _770_v91
				_770_v91 = _dafny.MultiSetOf((p0).Cmp((_767_v88).Cardinality()) < 0, !(((_this).Fm17(!(_753_v74), _768_v89, p0, _769_v90, globalState)).Cmp((_dafny.MultiSetOf(p0)).Cardinality()) >= 0))
				var _771_v92 _dafny.Map
				_ = _771_v92
				_771_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v1, p0)
				(globalState).F16 = (func() _dafny.Int {
					if (_770_v91).Contains(_660_v1) {
						return (_770_v91).Multiplicity(_660_v1)
					}
					return (func() _dafny.Int {
						if (_771_v92).Contains(_753_v74) {
							return (_771_v92).Get(_753_v74).(_dafny.Int)
						}
						return p0
					})()
				})()
				_666_v7 = _666_v7
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_750_v72), 0))
				_ = _index110
				(_750_v72).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfUint32((_662_v3).Cardinality()), (_this).F28(), p0, _755_v76, globalState), (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_750_v72), 0))
				_ = _index111
				(_750_v72).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if !(_754_v75) {
						return _dafny.IntOfUint32((_662_v3).Cardinality())
					}
					return (Companion_Default___.Fm35(_755_v76, (_this).F28(), globalState)).Cardinality()
				})(), _dafny.IntOfUint32((_662_v3).Cardinality())), (_index111).Int())
				(globalState).F19 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_662_v3, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_662_v3).Cardinality()))).Uint32(), _661_v2), _662_v3), _dafny.Companion_Sequence_.Concatenate(_662_v3, _662_v3))
			}
			if _756_v77 {
				_659_v0 = _659_v0
				var _772_v93 _dafny.Sequence
				_ = _772_v93
				_772_v93 = _dafny.SeqOf(_662_v3)
				(globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-857))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}((func(_773_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_774_i11 _dafny.Int) _dafny.CodePoint {
						return _773_v2
					}
				})(_661_v2))), (_772_v93).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_772_v93).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _775_v94 _dafny.Sequence
				_ = _775_v94
				_775_v94 = _dafny.SeqOf(_718_v50)
				var _776_v95 *C1
				_ = _776_v95
				var _nw107 *C1 = New_C1_()
				_ = _nw107
				_nw107.Ctor__((_this).F28(), _775_v94)
				_776_v95 = _nw107
				_666_v7 = _666_v7
				(globalState).F0 = true
			} else {
				var _777_v96 _dafny.Map
				_ = _777_v96
				_777_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_756_v77, _755_v76)
				var _778_v98 _dafny.MultiSet
				_ = _778_v98
				_778_v98 = _dafny.MultiSetOf(_756_v77, Companion_Default___.Fm2(_777_v96, globalState), _755_v76)
				var _779_v99 T1
				_ = _779_v99
				var _nw108 *C1 = New_C1_()
				_ = _nw108
				_nw108.Ctor__((_777_v96).Cardinality(), _dafny.SeqOf(func() _dafny.Map {
					var _coll36 = _dafny.NewMapBuilder()
					_ = _coll36
					for _iter43 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Int {
						if (_778_v98).Contains(_660_v1) {
							return (_778_v98).Multiplicity(_660_v1)
						}
						return p0
					})()))).Elements()); ; {
						_compr_36, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _780_v97 _dafny.Int
						_780_v97 = interface{}(_compr_36).(_dafny.Int)
						if (_dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Int {
							if (_778_v98).Contains(_660_v1) {
								return (_778_v98).Multiplicity(_660_v1)
							}
							return p0
						})()))).Contains(_780_v97) {
							_coll36.Add(Companion_Default___.SafeDivisionInt(_780_v97, (_this).F28()), _660_v1)
						}
					}
					return _coll36.ToMap()
				}()))
				_779_v99 = _nw108
				var _781_v100 _dafny.Map
				_ = _781_v100
				_781_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_755_v76), _779_v99)
				_781_v100 = (_781_v100).Merge(_781_v100)
				var _rhs90 _dafny.Int = p0
				_ = _rhs90
				var _rhs91 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_662_v3).Cardinality()))
				_ = _rhs91
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				_lhs93.F16 = _rhs90
				_lhs94.F16 = _rhs91
				_660_v1 = _756_v77
				_754_v75 = _660_v1
				var _782_v101 *C0
				_ = _782_v101
				var _nw109 *C0 = New_C0_()
				_ = _nw109
				_nw109.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_662_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-673))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg69 _dafny.Int) interface{} {
						return coer69(arg69)
					}
				}((func(_783_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_784_i12 _dafny.Int) _dafny.CodePoint {
						return _783_v2
					}
				})(_661_v2)))))
				_782_v101 = _nw109
			}
		}
		var _785_v102 _dafny.Array
		_ = _785_v102
		var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw110
		_785_v102 = _nw110
		r0 = _785_v102
		r1 = _660_v1
		return r0, r1
	}
}
func (_this *C2) M8(p0 bool, p1 bool, p2 bool, globalState *GlobalState) {
	{
		var _786_v0 _dafny.Sequence
		_ = _786_v0
		_786_v0 = _dafny.SeqOf(p2, p0, p1, p0)
		_786_v0 = _dafny.Companion_Sequence_.Concatenate(_786_v0, _786_v0)
		var _787_v1 _dafny.Map
		_ = _787_v1
		_787_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())
		var _788_v2 _dafny.Sequence
		_ = _788_v2
		_788_v2 = _dafny.UnicodeSeqOfUtf8Bytes("asmgvid")
		var _789_v3 _dafny.Array
		_ = _789_v3
		var _nwElement0_17 _dafny.Int = (_this).F28()
		_ = _nwElement0_17
		var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(8))
		_ = _nw111
		(_nw111).ArraySet1(_nwElement0_17, 0)
		(_nw111).ArraySet1((_this).F28(), 1)
		(_nw111).ArraySet1((func() _dafny.Int {
			if (_787_v1).Contains(p0) {
				return (_787_v1).Get(p0).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_788_v2).Cardinality())
		})(), 2)
		(_nw111).ArraySet1(_dafny.IntOfInt64(-74), 3)
		(_nw111).ArraySet1((_this).F28(), 4)
		(_nw111).ArraySet1((_this).F28(), 5)
		(_nw111).ArraySet1((_this).F28(), 6)
		(_nw111).ArraySet1((_this).F28(), 7)
		_789_v3 = _nw111
		var _790_v4 _dafny.Sequence
		_ = _790_v4
		_790_v4 = _dafny.SeqOf(_789_v3, _789_v3)
		var _791_v5 D6
		_ = _791_v5
		_791_v5 = Companion_D6_.Create_DC12_(_790_v4)
		var _pat_let_tv14 = _790_v4
		_ = _pat_let_tv14
		var _792_v6 _dafny.Map
		_ = _792_v6
		_792_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let10_0 D6) D6 {
			return func(_793_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let11_0 _dafny.Sequence) D6 {
					return func(_794_dt__update_hcf18_h0 _dafny.Sequence) D6 {
						return Companion_D6_.Create_DC12_(_794_dt__update_hcf18_h0)
					}(_pat_let11_0)
				}(_pat_let_tv14)
			}(_pat_let10_0)
		}(_791_v5), (_this).F28())
		var _795_v7 _dafny.Sequence
		_ = _795_v7
		_795_v7 = _dafny.SeqOf(_792_v6)
		var _796_v8 _dafny.Map
		_ = _796_v8
		_796_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(177), p1)
		var _797_v9 _dafny.Sequence
		_ = _797_v9
		_797_v9 = _dafny.SeqOf(_dafny.IntOfInt64(-739), (_this).F28(), (_this).F28())
		var _798_v10 _dafny.Array
		_ = _798_v10
		var _nwElement0_18 _dafny.Map = _792_v6
		_ = _nwElement0_18
		var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(19))
		_ = _nw112
		(_nw112).ArraySet1(_nwElement0_18, 0)
		(_nw112).ArraySet1(_792_v6, 1)
		(_nw112).ArraySet1(((_795_v7).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_795_v7).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_792_v6), 2)
		(_nw112).ArraySet1(_792_v6, 3)
		(_nw112).ArraySet1(_792_v6, 4)
		(_nw112).ArraySet1((_792_v6).Update(_791_v5, (_this).F28()), 5)
		(_nw112).ArraySet1((_792_v6).Merge(_792_v6), 6)
		(_nw112).ArraySet1(((_795_v7).Select((Companion_Default___.SafeIndex((_796_v8).Cardinality(), _dafny.IntOfUint32((_795_v7).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_792_v6), 7)
		(_nw112).ArraySet1((_792_v6).Merge(_792_v6), 8)
		(_nw112).ArraySet1(_792_v6, 9)
		(_nw112).ArraySet1(_792_v6, 10)
		(_nw112).ArraySet1((_792_v6).Update(_791_v5, (_this).F28()), 11)
		(_nw112).ArraySet1(_792_v6, 12)
		(_nw112).ArraySet1((_792_v6).Update(_791_v5, _dafny.IntOfUint32((_797_v9).Cardinality())), 13)
		(_nw112).ArraySet1(_792_v6, 14)
		(_nw112).ArraySet1((_792_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_791_v5, _dafny.IntOfUint32((_786_v0).Cardinality()))), 15)
		(_nw112).ArraySet1(_792_v6, 16)
		(_nw112).ArraySet1((_792_v6).Merge(_792_v6), 17)
		(_nw112).ArraySet1(_792_v6, 18)
		_798_v10 = _nw112
		_798_v10 = _798_v10
		var _799_v11 _dafny.Array
		_ = _799_v11
		var _nwElement0_19 bool = p2
		_ = _nwElement0_19
		var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(21))
		_ = _nw113
		(_nw113).ArraySet1(_nwElement0_19, 0)
		(_nw113).ArraySet1(p1, 1)
		(_nw113).ArraySet1((_dafny.IntOfInt64(511)).Cmp((_dafny.Zero).Minus((_this).F28())) >= 0, 2)
		(_nw113).ArraySet1(p0, 3)
		(_nw113).ArraySet1(p0, 4)
		(_nw113).ArraySet1(!(true), 5)
		(_nw113).ArraySet1(p0, 6)
		(_nw113).ArraySet1(p1, 7)
		(_nw113).ArraySet1(p0, 8)
		(_nw113).ArraySet1(p1, 9)
		(_nw113).ArraySet1(p1, 10)
		(_nw113).ArraySet1(p2, 11)
		(_nw113).ArraySet1(p1, 12)
		(_nw113).ArraySet1(p1, 13)
		(_nw113).ArraySet1(p2, 14)
		(_nw113).ArraySet1(p0, 15)
		(_nw113).ArraySet1(true, 16)
		(_nw113).ArraySet1((func() bool {
			if p2 {
				return !(p2)
			}
			return p1
		})(), 17)
		(_nw113).ArraySet1(!(!(p2)), 18)
		(_nw113).ArraySet1(p2, 19)
		(_nw113).ArraySet1(p2, 20)
		_799_v11 = _nw113
		for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_799_v11), 0))); ; {
			_guard_loop_6, _ok44 := _iter44()
			if !_ok44 {
				break
			}
			var _800_i0 _dafny.Int
			_800_i0 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_800_i0).Sign() != -1) && ((_800_i0).Cmp(_dafny.ArrayLen((_799_v11), 0)) < 0)) {
				(_799_v11).ArraySet1(!(((_this).F28()).Cmp((_this).F28()) > 0), (_800_i0).Int())
			}
		}
		var _801_v12 D12
		_ = _801_v12
		_801_v12 = Companion_D12_.Create_DC23_((_this).F28())
		var _802_v13 D12
		_ = _802_v13
		_802_v13 = Companion_D12_.Create_DC24_(_801_v12)
		var _803_v14 D12
		_ = _803_v14
		_803_v14 = Companion_D12_.Create_DC24_(_802_v13)
		var _804_i1 _dafny.Int
		_ = _804_i1
		_804_i1 = _dafny.Zero
		{
			var _pat_let_tv15 = p1
			_ = _pat_let_tv15
			var _pat_let_tv16 = p2
			_ = _pat_let_tv16
			var _pat_let_tv17 = p0
			_ = _pat_let_tv17
			for func(_source14 D12) bool {
				if _source14.Is_DC23() {
					var _820___mcc_h0 _dafny.Int = _source14.Get_().(D12_DC23).Cf30
					_ = _820___mcc_h0
					var _821_cf30 _dafny.Int = _820___mcc_h0
					_ = _821_cf30
					return _pat_let_tv15
				} else if _source14.Is_DC22() {
					var _822___mcc_h1 _dafny.Map = _source14.Get_().(D12_DC22).Cf29
					_ = _822___mcc_h1
					var _823_cf29 _dafny.Map = _822___mcc_h1
					_ = _823_cf29
					return _pat_let_tv16
				} else {
					var _824___mcc_h2 D12 = _source14.Get_().(D12_DC24).Cf31
					_ = _824___mcc_h2
					var _825_cf31 D12 = _824___mcc_h2
					_ = _825_cf31
					return !(_pat_let_tv17)
				}
			}(_803_v14) {
				{
					if (_804_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_804_i1 = (_804_i1).Plus(_dafny.One)
					(globalState).F21 = p0
					if p1 {
						(globalState).F7 = ((_this).F28()).Plus(((_this).F28()).Minus((_this).F28()))
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_799_v11), 0))
						_ = _index112
						(_799_v11).ArraySet1(p0, (_index112).Int())
						var _805_v15 D1
						_ = _805_v15
						_805_v15 = Companion_D1_.Create_DC3_(p2)
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_799_v11), 0))
						_ = _index113
						(_799_v11).ArraySet1((_805_v15).Dtor_cf3(), (_index113).Int())
						(globalState).F21 = false
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_789_v3), 0))
						_ = _index114
						(_789_v3).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F28(), (_this).F28()), (_index114).Int())
						var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_789_v3), 0))
						_ = _index115
						(_789_v3).ArraySet1(Companion_Default___.SafeDivisionInt(((_this).F28()).Times((_this).F28()), ((_dafny.Zero).Minus((_this).F28())).Plus((_787_v1).Cardinality())), (_index115).Int())
						(globalState).F21 = false
					} else {
						(globalState).F16 = (_this).F28()
						var _806_v16 *C0
						_ = _806_v16
						var _nw114 *C0 = New_C0_()
						_ = _nw114
						_nw114.Ctor__(p0)
						_806_v16 = _nw114
						(globalState).F7 = Companion_Default___.SafeModuloInt((_this).F28(), (_this).F28())
						var _807_v17 _dafny.Set
						_ = _807_v17
						_807_v17 = _dafny.SetOf((_this).F28())
						var _808_v18 _dafny.Map
						_ = _808_v18
						_808_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_807_v17, _dafny.SetOf(p2, p0))
						var _809_v19 _dafny.Set
						_ = _809_v19
						_809_v19 = _dafny.SetOf(_806_v16.F27, false, p1, p2, p1)
						_808_v18 = (_808_v18).Update(_807_v17, _809_v19)
						var _810_v20 _dafny.Array
						_ = _810_v20
						var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
						_ = _nw115
						_810_v20 = _nw115
						var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_810_v20), 0))
						_ = _index116
						(_810_v20).ArraySet1(_786_v0, (_index116).Int())
						var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_810_v20), 0))
						_ = _index117
						(_810_v20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_786_v0, _786_v0), (_index117).Int())
					}
					var _811_v21 _dafny.Array
					_ = _811_v21
					var _len0_16 _dafny.Int = _dafny.IntOfInt64(11)
					_ = _len0_16
					var _nw116 _dafny.Array
					_ = _nw116
					if _len0_16.Cmp(_dafny.Zero) == 0 {
						_nw116 = _dafny.NewArray(_len0_16)
					} else {
						var _init16 func(_dafny.Int) _dafny.Set = (func(_812_p0 bool, _813_p1 bool) func(_dafny.Int) _dafny.Set {
							return func(_814_i2 _dafny.Int) _dafny.Set {
								return _dafny.SetOf(_812_p0, _813_p1, !(_812_p0), _813_p1, _812_p0)
							}
						})(p0, p1)
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
					_811_v21 = _nw116
					var _815_v22 D7
					_ = _815_v22
					_815_v22 = Companion_D7_.Create_DC16_(_811_v21)
					_815_v22 = Companion_D7_.Create_DC16_(_811_v21)
					var _816_v23 _dafny.CodePoint
					_ = _816_v23
					_816_v23 = _dafny.CodePoint('b')
					var _817_v24 _dafny.Map
					_ = _817_v24
					_817_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
					var _818_v25 _dafny.Set
					_ = _818_v25
					_818_v25 = _dafny.SetOf(p2)
					var _819_v26 _dafny.Map
					_ = _819_v26
					_819_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _818_v25)
					var _rhs92 _dafny.CodePoint = _816_v23
					_ = _rhs92
					var _rhs93 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F28(), Companion_Default___.Fm0((_817_v24).Cardinality(), Companion_Default___.Fm0((_this).F28(), _dafny.IntOfInt64(350), ((func() _dafny.Set {
						if (_819_v26).Contains(p2) {
							return (_819_v26).Get(p2).(_dafny.Set)
						}
						return _818_v25
					})()).Cardinality(), p1, globalState), (_this).F28(), p0, globalState))))
					_ = _rhs93
					var _rhs94 bool = p2
					_ = _rhs94
					var _rhs95 _dafny.Int = (_this).F28()
					_ = _rhs95
					var _lhs95 *GlobalState = globalState
					_ = _lhs95
					var _lhs96 *GlobalState = globalState
					_ = _lhs96
					var _lhs97 *GlobalState = globalState
					_ = _lhs97
					_816_v23 = _rhs92
					_lhs95.F7 = _rhs93
					_lhs96.F2 = _rhs94
					_lhs97.F16 = _rhs95
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _826_v27 _dafny.Array
		_ = _826_v27
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_17
		var _nw117 _dafny.Array
		_ = _nw117
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw117 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.MultiSet = (func(_827_p2 bool, _828_p0 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_829_i3 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(!(_827_p2), _827_p2, _828_p0)
				}
			})(p2, p0)
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
		_826_v27 = _nw117
		var _830_v28 _dafny.MultiSet
		_ = _830_v28
		_830_v28 = _dafny.MultiSetOf(p2)
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_826_v27), 0))
		_ = _index118
		(_826_v27).ArraySet1(_830_v28, (_index118).Int())
		var _831_v29 _dafny.Map
		_ = _831_v29
		_831_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
		var _832_v30 _dafny.MultiSet
		_ = _832_v30
		_832_v30 = _dafny.MultiSetOf(Companion_Default___.Fm34(globalState))
		var _833_v31 _dafny.Sequence
		_ = _833_v31
		_833_v31 = _dafny.SeqOf(_832_v30)
		var _834_v32 _dafny.Map
		_ = _834_v32
		_834_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _788_v2)
		var _835_v33 _dafny.Map
		_ = _835_v33
		_835_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_834_v32).Cardinality(), _dafny.IntOfUint32((_788_v2).Cardinality()))
		var _836_v34 _dafny.Set
		_ = _836_v34
		_836_v34 = _dafny.SetOf((_this).F28(), (_this).F28(), ((_796_v8).Cardinality()).Plus(_dafny.IntOfUint32((_833_v31).Cardinality())), Companion_Default___.Fm0((_835_v33).Cardinality(), (_this).F28(), (_this).F28(), p0, globalState))
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_826_v27), 0))
		_ = _index119
		var _rhs96 _dafny.MultiSet = _830_v28
		_ = _rhs96
		var _rhs97 _dafny.Map = (_831_v29).Merge(_831_v29)
		_ = _rhs97
		var _rhs98 _dafny.Set = (_836_v34).Intersection(_dafny.SetOf((_this).F28(), (_this).F28()))
		_ = _rhs98
		var _lhs98 _dafny.Array = _826_v27
		_ = _lhs98
		var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_826_v27), 0))
		_ = _lhs99
		(_lhs98).ArraySet1(_rhs96, (_lhs99).Int())
		_831_v29 = _rhs97
		_836_v34 = _rhs98
		(globalState).F2 = (_this).Fm7(((_this).F28()).Cmp((_this).F28()) <= 0, (_this).F28(), p1, globalState)
	}
}
func (_this *C2) M9(globalState *GlobalState) (bool, bool, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _837_v0 _dafny.Sequence
		_ = _837_v0
		_837_v0 = _dafny.SeqOf((_this).F28())
		var _838_i0 _dafny.Int
		_ = _838_i0
		_838_i0 = _dafny.Zero
		{
			for !(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_837_v0, _dafny.Companion_Sequence_.Update(_837_v0, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_837_v0).Cardinality()))).Uint32(), (_this).F28())), Companion_Default___.Fm37(globalState))) {
				{
					if (_838_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_838_i0 = (_838_i0).Plus(_dafny.One)
					var _839_v1 _dafny.Array
					_ = _839_v1
					var _len0_18 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_18
					var _nw118 _dafny.Array
					_ = _nw118
					if _len0_18.Cmp(_dafny.Zero) == 0 {
						_nw118 = _dafny.NewArray(_len0_18)
					} else {
						var _init18 func(_dafny.Int) _dafny.CodePoint = func(_840_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('o')
						}
						_ = _init18
						var _element0_18 = _init18(_dafny.Zero)
						_ = _element0_18
						_nw118 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
						(_nw118).ArraySet1CodePoint(_element0_18, 0)
						var _nativeLen0_18 = (_len0_18).Int()
						_ = _nativeLen0_18
						for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
							(_nw118).ArraySet1CodePoint(_init18(_dafny.IntOf(_i0_18)), _i0_18)
						}
					}
					_839_v1 = _nw118
					var _841_v2 _dafny.CodePoint
					_ = _841_v2
					_841_v2 = _dafny.CodePoint('h')
					var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))
					_ = _index120
					(_839_v1).ArraySet1CodePoint(_841_v2, (_index120).Int())
					var _842_v3 bool
					_ = _842_v3
					_842_v3 = false
					var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))
					_ = _index121
					(_839_v1).ArraySet1CodePoint(Companion_Default___.Fm27(((_this).F28()).Cmp((_this).F28()) == 0, _842_v3, globalState), (_index121).Int())
					var _843_v4 _dafny.Map
					_ = _843_v4
					_843_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v3, _842_v3)
					var _844_v5 _dafny.Sequence
					_ = _844_v5
					_844_v5 = _dafny.SeqOf((func() bool {
						if (_843_v4).Contains(_842_v3) {
							return (_843_v4).Get(_842_v3).(bool)
						}
						return _842_v3
					})(), _842_v3)
					var _845_v6 _dafny.Sequence
					_ = _845_v6
					_845_v6 = _844_v5
					var _846_v7 _dafny.Map
					_ = _846_v7
					_846_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v3, _844_v5)
					var _847_v8 D14
					_ = _847_v8
					_847_v8 = Companion_D14_.Create_DC29_((_this).Fm17(_842_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(884))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg70 _dafny.Int) interface{} {
							return coer70(arg70)
						}
					}((func(_848_v3 bool) func(_dafny.Int) _dafny.Int {
						return func(_849_i2 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.SeqOf(_848_v3, false)).Cardinality())
						}
					})(_842_v3))), (_this).F28(), _845_v6, globalState), _846_v7)
					(globalState).F16 = (_dafny.Zero).Minus((_847_v8).Dtor_cf38())
					var _source15 D14 = Companion_D14_.Create_DC30_(_837_v0, _dafny.Companion_Sequence_.Update(_844_v5, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_844_v5).Cardinality()))).Uint32(), !(_842_v3)))
					_ = _source15
					if _source15.Is_DC29() {
						var _850___mcc_h0 _dafny.Int = _source15.Get_().(D14_DC29).Cf38
						_ = _850___mcc_h0
						var _851___mcc_h1 _dafny.Map = _source15.Get_().(D14_DC29).Cf39
						_ = _851___mcc_h1
						var _852_cf39 _dafny.Map = _851___mcc_h1
						_ = _852_cf39
						var _853_cf38 _dafny.Int = _850___mcc_h0
						_ = _853_cf38
						var _854_v9 _dafny.MultiSet
						_ = _854_v9
						_854_v9 = _dafny.MultiSetOf(!(_842_v3) || (_842_v3))
						var _855_v10 _dafny.Map
						_ = _855_v10
						_855_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
						var _856_v11 D13
						_ = _856_v11
						_856_v11 = Companion_D13_.Create_DC25_(_855_v10)
						var _857_v12 _dafny.Set
						_ = _857_v12
						_857_v12 = _dafny.SetOf(_856_v11)
						var _858_v13 _dafny.Map
						_ = _858_v13
						_858_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_857_v12, _841_v2)
						var _rhs99 _dafny.Int = (_this).F28()
						_ = _rhs99
						var _rhs100 _dafny.Int = _dafny.IntOfInt64(209)
						_ = _rhs100
						var _rhs101 _dafny.Int = (func() _dafny.Int {
							if (_854_v9).Contains(_842_v3) {
								return (_854_v9).Multiplicity(_842_v3)
							}
							return ((_858_v13).Update(_857_v12, _841_v2)).Cardinality()
						})()
						_ = _rhs101
						var _rhs102 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm0((_853_cf38).Plus((_this).F28()), _dafny.IntOfInt64(-485), (func() _dafny.Int {
							if !(_842_v3) {
								return (_dafny.MultiSetFromSeq(_837_v0)).Cardinality()
							}
							return _853_cf38
						})(), _842_v3, globalState))
						_ = _rhs102
						var _lhs100 *GlobalState = globalState
						_ = _lhs100
						var _lhs101 *GlobalState = globalState
						_ = _lhs101
						var _lhs102 *GlobalState = globalState
						_ = _lhs102
						_lhs100.F16 = _rhs99
						_853_cf38 = _rhs100
						_lhs101.F7 = _rhs101
						_lhs102.F7 = _rhs102
						var _859_v14 _dafny.Array
						_ = _859_v14
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(11)
						_ = _len0_19
						var _nw119 _dafny.Array
						_ = _nw119
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw119 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) _dafny.Int = (func(_860_v9 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
								return func(_861_i3 _dafny.Int) _dafny.Int {
									return (_861_i3).Times((_860_v9).Cardinality())
								}
							})(_854_v9)
							_ = _init19
							var _element0_19 = _init19(_dafny.Zero)
							_ = _element0_19
							_nw119 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
							(_nw119).ArraySet1(_element0_19, 0)
							var _nativeLen0_19 = (_len0_19).Int()
							_ = _nativeLen0_19
							for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
								(_nw119).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
							}
						}
						_859_v14 = _nw119
						var _862_v15 _dafny.MultiSet
						_ = _862_v15
						_862_v15 = _dafny.MultiSetOf(_853_cf38)
						var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))
						_ = _index122
						(_859_v14).ArraySet1((_853_cf38).Minus((_862_v15).Cardinality()), (_index122).Int())
						var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))
						_ = _index123
						(_859_v14).ArraySet1(Companion_Default___.SafeModuloInt(_853_cf38, _853_cf38), (_index123).Int())
						_837_v0 = _837_v0
						var _863_v16 _dafny.Map
						_ = _863_v16
						_863_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_837_v0, (Companion_Default___.SafeIndex((_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_837_v0).Cardinality()))).Uint32(), _853_cf38), (Companion_Default___.SafeIndex((_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_837_v0, (Companion_Default___.SafeIndex((_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_837_v0).Cardinality()))).Uint32(), _853_cf38)).Cardinality()))).Uint32(), (_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int)), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-205)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_837_v0, (Companion_Default___.SafeIndex((_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_837_v0).Cardinality()))).Uint32(), _853_cf38), (Companion_Default___.SafeIndex((_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_837_v0, (Companion_Default___.SafeIndex((_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_837_v0).Cardinality()))).Uint32(), _853_cf38)).Cardinality()))).Uint32(), (_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int))).Cardinality()))).Uint32(), (_this).F28()), _844_v5)
						var _864_v17 _dafny.Sequence
						_ = _864_v17
						_864_v17 = _dafny.UnicodeSeqOfUtf8Bytes("gsaqxlayd")
						var _865_v18 _dafny.Sequence
						_ = _865_v18
						_865_v18 = _dafny.SeqOf(_863_v16, _863_v16, _863_v16, ((_863_v16).Update(_837_v0, _844_v5)).Update(_837_v0, _844_v5))
						var _866_v19 _dafny.Set
						_ = _866_v19
						_866_v19 = _dafny.SetOf(_842_v3)
						var _867_v21 D4
						_ = _867_v21
						_867_v21 = Companion_D4_.Create_DC9_(_dafny.IntOfUint32((_837_v0).Cardinality()))
						var _868_v22 _dafny.Map
						_ = _868_v22
						_868_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm21(globalState), _867_v21)
						var _869_v23 D16
						_ = _869_v23
						_869_v23 = Companion_D16_.Create_DC35_(_863_v16)
						var _870_v25 _dafny.Sequence
						_ = _870_v25
						_870_v25 = _dafny.SeqOf(_837_v0)
						var _871_v28 _dafny.MultiSet
						_ = _871_v28
						_871_v28 = _dafny.MultiSetOf(_837_v0)
						var _pat_let_tv18 = _863_v16
						_ = _pat_let_tv18
						var _872_v29 _dafny.Array
						_ = _872_v29
						var _nwElement0_20 _dafny.Map = _863_v16
						_ = _nwElement0_20
						var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(27))
						_ = _nw120
						(_nw120).ArraySet1(_nwElement0_20, 0)
						(_nw120).ArraySet1(_863_v16, 1)
						(_nw120).ArraySet1(_863_v16, 2)
						(_nw120).ArraySet1(_863_v16, 3)
						(_nw120).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v0, _dafny.SeqOf(_842_v3, false, _842_v3)), 4)
						(_nw120).ArraySet1(((_863_v16).Update(_dafny.SeqOf(_dafny.IntOfUint32((_864_v17).Cardinality()), _dafny.IntOfInt64(205)), _844_v5)).Merge(_863_v16), 5)
						(_nw120).ArraySet1(_863_v16, 6)
						(_nw120).ArraySet1((_865_v18).Select((Companion_Default___.SafeIndex((_866_v19).Cardinality(), _dafny.IntOfUint32((_865_v18).Cardinality()))).Uint32()).(_dafny.Map), 7)
						(_nw120).ArraySet1(_863_v16, 8)
						(_nw120).ArraySet1(_863_v16, 9)
						(_nw120).ArraySet1(_863_v16, 10)
						(_nw120).ArraySet1(_863_v16, 11)
						(_nw120).ArraySet1(_863_v16, 12)
						(_nw120).ArraySet1(Companion_Default___.Fm39((_854_v9).Cardinality(), _842_v3, (_this).F28(), globalState), 13)
						(_nw120).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v0, _844_v5), 14)
						(_nw120).ArraySet1(_863_v16, 15)
						(_nw120).ArraySet1((func() _dafny.Map {
							var _coll37 = _dafny.NewMapBuilder()
							_ = _coll37
							for _iter45 := _dafny.Iterate((_868_v22).Keys().Elements()); ; {
								_compr_37, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _873_v20 _dafny.Sequence
								_873_v20 = interface{}(_compr_37).(_dafny.Sequence)
								if (_868_v22).Contains(_873_v20) {
									_coll37.Add(_873_v20, _dafny.SeqOf(_842_v3))
								}
							}
							return _coll37.ToMap()
						}()).Merge(_863_v16), 16)
						(_nw120).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v0, _844_v5), 17)
						(_nw120).ArraySet1((func(_pat_let12_0 D16) D16 {
							return func(_874_dt__update__tmp_h0 D16) D16 {
								return func(_pat_let13_0 _dafny.Map) D16 {
									return func(_875_dt__update_hcf46_h0 _dafny.Map) D16 {
										return Companion_D16_.Create_DC35_(_875_dt__update_hcf46_h0)
									}(_pat_let13_0)
								}(_pat_let_tv18)
							}(_pat_let12_0)
						}(_869_v23)).Dtor_cf46(), 18)
						(_nw120).ArraySet1(_863_v16, 19)
						(_nw120).ArraySet1((func() _dafny.Map {
							var _coll38 = _dafny.NewMapBuilder()
							_ = _coll38
							for _iter46 := _dafny.Iterate((_870_v25).Elements()); ; {
								_compr_38, _ok46 := _iter46()
								if !_ok46 {
									break
								}
								var _876_v24 _dafny.Sequence
								_876_v24 = interface{}(_compr_38).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_870_v25, _876_v24) {
									_coll38.Add(_876_v24, _844_v5)
								}
							}
							return _coll38.ToMap()
						}()).Merge(_863_v16), 20)
						(_nw120).ArraySet1(_863_v16, 21)
						(_nw120).ArraySet1(_863_v16, 22)
						(_nw120).ArraySet1(_863_v16, 23)
						(_nw120).ArraySet1(_863_v16, 24)
						(_nw120).ArraySet1((func() _dafny.Map {
							var _coll39 = _dafny.NewMapBuilder()
							_ = _coll39
							for _iter47 := _dafny.Iterate((_870_v25).Elements()); ; {
								_compr_39, _ok47 := _iter47()
								if !_ok47 {
									break
								}
								var _877_v26 _dafny.Sequence
								_877_v26 = interface{}(_compr_39).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_870_v25, _877_v26) {
									_coll39.Add(_877_v26, _844_v5)
								}
							}
							return _coll39.ToMap()
						}()).Update(_837_v0, _844_v5), 25)
						(_nw120).ArraySet1((func() _dafny.Map {
							var _coll40 = _dafny.NewMapBuilder()
							_ = _coll40
							for _iter48 := _dafny.Iterate((_871_v28).Elements()); ; {
								_compr_40, _ok48 := _iter48()
								if !_ok48 {
									break
								}
								var _878_v27 _dafny.Sequence
								_878_v27 = interface{}(_compr_40).(_dafny.Sequence)
								if (_871_v28).Contains(_878_v27) {
									_coll40.Add(_878_v27, _dafny.SeqOf(_842_v3, false))
								}
							}
							return _coll40.ToMap()
						}()).Merge(_863_v16), 26)
						_872_v29 = _nw120
						var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_872_v29), 0))
						_ = _index124
						(_872_v29).ArraySet1(_863_v16, (_index124).Int())
						var _879_v30 *C0
						_ = _879_v30
						var _nw121 *C0 = New_C0_()
						_ = _nw121
						_nw121.Ctor__(false)
						_879_v30 = _nw121
						var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_872_v29), 0))
						_ = _index125
						var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))
						_ = _index126
						var _rhs103 _dafny.Map = (_869_v23).Dtor_cf46()
						_ = _rhs103
						var _rhs104 _dafny.CodePoint = _841_v2
						_ = _rhs104
						var _rhs105 _dafny.Int = (_859_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))).Int()).(_dafny.Int)
						_ = _rhs105
						var _rhs106 bool = ((_this).F28()).Cmp(Companion_Default___.SafeDivisionInt((_this).F28(), _dafny.IntOfUint32((_864_v17).Cardinality()))) != 0
						_ = _rhs106
						var _rhs107 *C0 = _879_v30
						_ = _rhs107
						var _lhs103 _dafny.Array = _872_v29
						_ = _lhs103
						var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_872_v29), 0))
						_ = _lhs104
						var _lhs105 _dafny.Array = _859_v14
						_ = _lhs105
						var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_859_v14), 0))
						_ = _lhs106
						(_lhs103).ArraySet1(_rhs103, (_lhs104).Int())
						_841_v2 = _rhs104
						(_lhs105).ArraySet1(_rhs105, (_lhs106).Int())
						r1 = _rhs106
						_879_v30 = _rhs107
					} else if _source15.Is_DC30() {
						var _880___mcc_h2 _dafny.Sequence = _source15.Get_().(D14_DC30).Cf40
						_ = _880___mcc_h2
						var _881___mcc_h3 _dafny.Sequence = _source15.Get_().(D14_DC30).Cf41
						_ = _881___mcc_h3
						var _882_cf41 _dafny.Sequence = _881___mcc_h3
						_ = _882_cf41
						var _883_cf40 _dafny.Sequence = _880___mcc_h2
						_ = _883_cf40
						var _884_v31 _dafny.Map
						_ = _884_v31
						_884_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_882_cf41).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F28()), _dafny.IntOfUint32((_882_cf41).Cardinality()))).Uint32()).(bool), (_this).Fm6((_this).F28(), globalState)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg71 _dafny.Int) interface{} {
								return coer71(arg71)
							}
						}((func(_885_v1 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
							return func(_886_i4 _dafny.Int) _dafny.CodePoint {
								return (_885_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_885_v1), 0))).Int())
							}
						})(_839_v1))))
						var _887_v32 _dafny.Sequence
						_ = _887_v32
						_887_v32 = _dafny.UnicodeSeqOfUtf8Bytes("bfsvcv")
						var _888_v33 _dafny.Map
						_ = _888_v33
						_888_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
						Companion_Default___.M0(!(_842_v3) || (_842_v3), _dafny.IntOfInt64(-369), (_884_v31).Merge((_884_v31).Update(_882_cf41, _887_v32)), (_dafny.IntOfInt64(811)).Minus((_888_v33).Cardinality()), globalState)
						(globalState).F7 = (_dafny.Zero).Minus((_this).F28())
						(globalState).F16 = (_this).F28()
						var _889_v34 _dafny.Array
						_ = _889_v34
						var _nw122 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
						_ = _nw122
						_889_v34 = _nw122
						var _890_v35 _dafny.Array
						_ = _890_v35
						var _nwElement0_21 _dafny.Array = _889_v34
						_ = _nwElement0_21
						var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(18))
						_ = _nw123
						(_nw123).ArraySet1(_nwElement0_21, 0)
						(_nw123).ArraySet1((func() _dafny.Array {
							if _842_v3 {
								return _889_v34
							}
							return _889_v34
						})(), 1)
						(_nw123).ArraySet1(_889_v34, 2)
						(_nw123).ArraySet1(_889_v34, 3)
						(_nw123).ArraySet1(_889_v34, 4)
						(_nw123).ArraySet1(_889_v34, 5)
						(_nw123).ArraySet1(_889_v34, 6)
						(_nw123).ArraySet1(_889_v34, 7)
						(_nw123).ArraySet1(_889_v34, 8)
						(_nw123).ArraySet1(_889_v34, 9)
						(_nw123).ArraySet1(_889_v34, 10)
						(_nw123).ArraySet1(_889_v34, 11)
						(_nw123).ArraySet1((func() _dafny.Array {
							if _842_v3 {
								return _889_v34
							}
							return _889_v34
						})(), 12)
						(_nw123).ArraySet1(_889_v34, 13)
						(_nw123).ArraySet1(_889_v34, 14)
						(_nw123).ArraySet1(_889_v34, 15)
						(_nw123).ArraySet1(_889_v34, 16)
						(_nw123).ArraySet1(_889_v34, 17)
						_890_v35 = _nw123
						var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_890_v35), 0))
						_ = _index127
						(_890_v35).ArraySet1(_889_v34, (_index127).Int())
						var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_890_v35), 0))
						_ = _index128
						(_890_v35).ArraySet1(_889_v34, (_index128).Int())
					} else if _source15.Is_DC28() {
						var _891___mcc_h4 _dafny.Set = _source15.Get_().(D14_DC28).Cf37
						_ = _891___mcc_h4
						var _892_cf37 _dafny.Set = _891___mcc_h4
						_ = _892_cf37
						(globalState).F7 = ((_this).F28()).Times(Companion_Default___.SafeModuloInt((_this).F28(), _dafny.IntOfInt64(300)))
						var _893_v36 _dafny.Set
						_ = _893_v36
						_893_v36 = _dafny.SetOf(_842_v3)
						var _894_v37 _dafny.Array
						_ = _894_v37
						var _nw124 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
						_ = _nw124
						_894_v37 = _nw124
						var _895_v38 _dafny.Map
						_ = _895_v38
						_895_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_894_v37, _893_v36)
						var _896_v39 _dafny.Array
						_ = _896_v39
						var _nwElement0_22 _dafny.Set = (_893_v36).Difference(_893_v36)
						_ = _nwElement0_22
						var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(16))
						_ = _nw125
						(_nw125).ArraySet1(_nwElement0_22, 0)
						(_nw125).ArraySet1(_893_v36, 1)
						(_nw125).ArraySet1(_893_v36, 2)
						(_nw125).ArraySet1((_893_v36).Difference(_893_v36), 3)
						(_nw125).ArraySet1(_893_v36, 4)
						(_nw125).ArraySet1(_dafny.SetOf((func() bool {
							if (_843_v4).Contains(_842_v3) {
								return (_843_v4).Get(_842_v3).(bool)
							}
							return _842_v3
						})()), 5)
						(_nw125).ArraySet1(Companion_Default___.Fm19(_dafny.IntOfInt64(365), _842_v3, _842_v3, (_this).F28(), globalState), 6)
						(_nw125).ArraySet1((_893_v36).Difference(_893_v36), 7)
						(_nw125).ArraySet1((func() _dafny.Set {
							if _842_v3 {
								return _893_v36
							}
							return _893_v36
						})(), 8)
						(_nw125).ArraySet1(_893_v36, 9)
						(_nw125).ArraySet1(_dafny.SetOf(_842_v3, (_this).Fm7(_842_v3, (_this).F28(), _842_v3, globalState)), 10)
						(_nw125).ArraySet1(_893_v36, 11)
						(_nw125).ArraySet1((_893_v36).Union(_893_v36), 12)
						(_nw125).ArraySet1(_893_v36, 13)
						(_nw125).ArraySet1((_893_v36).Difference(_893_v36), 14)
						(_nw125).ArraySet1((func() _dafny.Set {
							if (_895_v38).Contains(_894_v37) {
								return (_895_v38).Get(_894_v37).(_dafny.Set)
							}
							return _893_v36
						})(), 15)
						_896_v39 = _nw125
						var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_896_v39), 0))
						_ = _index129
						(_896_v39).ArraySet1((_893_v36).Intersection(_893_v36), (_index129).Int())
						var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_896_v39), 0))
						_ = _index130
						(_896_v39).ArraySet1(_893_v36, (_index130).Int())
						var _897_v40 _dafny.Sequence
						_ = _897_v40
						_897_v40 = _dafny.SeqOf(_837_v0)
						var _898_v41 _dafny.Array
						_ = _898_v41
						var _nwElement0_23 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_837_v0, _dafny.SeqOf((_892_cf37).Cardinality(), (_this).F28(), (_this).F28()))
						_ = _nwElement0_23
						var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(4))
						_ = _nw126
						(_nw126).ArraySet1(_nwElement0_23, 0)
						(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_837_v0, _837_v0), 1)
						(_nw126).ArraySet1(_837_v0, 2)
						(_nw126).ArraySet1((_897_v40).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_897_v40).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
						_898_v41 = _nw126
						var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((_898_v41), 0))
						_ = _index131
						(_898_v41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm21(globalState), _837_v0), (_index131).Int())
						var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(791), _dafny.ArrayLen((_898_v41), 0))
						_ = _index132
						(_898_v41).ArraySet1(_837_v0, (_index132).Int())
						var _899_v42 *C0
						_ = _899_v42
						var _nw127 *C0 = New_C0_()
						_ = _nw127
						_nw127.Ctor__((_dafny.IntOfInt64(78)).Cmp((_this).F28()) == 0)
						_899_v42 = _nw127
						_899_v42 = _899_v42
					} else {
						var _900___mcc_h5 D14 = _source15.Get_().(D14_DC31).Cf42
						_ = _900___mcc_h5
						var _901_cf42 D14 = _900___mcc_h5
						_ = _901_cf42
						var _902_v43 D15
						_ = _902_v43
						_902_v43 = Companion_D15_.Create_DC33_(_837_v0)
						var _903_v44 D7
						_ = _903_v44
						_903_v44 = Companion_D7_.Create_DC15_((_902_v43).Dtor_cf44())
						var _904_v45 D7
						_ = _904_v45
						_904_v45 = Companion_D7_.Create_DC17_(_903_v44)
						_904_v45 = _904_v45
						var _905_v46 _dafny.Map
						_ = _905_v46
						_905_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v3, (_dafny.Zero).Minus((_this).F28()))
						var _rhs108 _dafny.Int = (_837_v0).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_837_v0).Cardinality()))).Uint32()).(_dafny.Int)
						_ = _rhs108
						var _rhs109 _dafny.Int = (func() _dafny.Int {
							if (_905_v46).Contains(true) {
								return (_905_v46).Get(true).(_dafny.Int)
							}
							return (_this).F28()
						})()
						_ = _rhs109
						var _lhs107 *GlobalState = globalState
						_ = _lhs107
						var _lhs108 *GlobalState = globalState
						_ = _lhs108
						_lhs107.F7 = _rhs108
						_lhs108.F16 = _rhs109
						var _906_v47 _dafny.Map
						_ = _906_v47
						_906_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F28()), _844_v5)
						var _907_v48 D16
						_ = _907_v48
						_907_v48 = Companion_D16_.Create_DC35_(_906_v47)
						var _908_v49 D16
						_ = _908_v49
						_908_v49 = Companion_D16_.Create_DC37_(_907_v48)
						_908_v49 = Companion_D16_.Create_DC37_(_907_v48)
						var _909_v50 D12
						_ = _909_v50
						_909_v50 = Companion_D12_.Create_DC22_(_905_v46)
						var _910_v51 _dafny.Set
						_ = _910_v51
						_910_v51 = _dafny.SetOf(_909_v50)
						var _911_v52 _dafny.Map
						_ = _911_v52
						_911_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_909_v50, _842_v3)
						_910_v51 = ((_dafny.SetOf(Companion_D12_.Create_DC22_(_905_v46))).Union(_910_v51)).Difference((func() _dafny.Set {
							var _coll41 = _dafny.NewBuilder()
							_ = _coll41
							for _iter49 := _dafny.Iterate((_911_v52).Keys().Elements()); ; {
								_compr_41, _ok49 := _iter49()
								if !_ok49 {
									break
								}
								var _912_v53 D12
								_912_v53 = interface{}(_compr_41).(D12)
								if (_911_v52).Contains(_912_v53) {
									_coll41.Add(_912_v53)
								}
							}
							return _coll41.ToSet()
						}()).Difference(_910_v51))
					}
					var _913_v55 D16
					_ = _913_v55
					_913_v55 = Companion_D16_.Create_DC37_(Companion_D16_.Create_DC35_(func() _dafny.Map {
						var _coll42 = _dafny.NewMapBuilder()
						_ = _coll42
						for _iter50 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v0, _842_v3)).Keys().Elements()); ; {
							_compr_42, _ok50 := _iter50()
							if !_ok50 {
								break
							}
							var _914_v54 _dafny.Sequence
							_914_v54 = interface{}(_compr_42).(_dafny.Sequence)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v0, _842_v3)).Contains(_914_v54) {
								_coll42.Add(_914_v54, _844_v5)
							}
						}
						return _coll42.ToMap()
					}()))
					var _915_v56 _dafny.MultiSet
					_ = _915_v56
					_915_v56 = _dafny.MultiSetOf(_913_v55, _913_v55, _913_v55)
					var _916_v57 _dafny.Map
					_ = _916_v57
					_916_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v3, (_915_v56).Union(_915_v56))
					var _917_v58 _dafny.Array
					_ = _917_v58
					var _len0_20 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_20
					var _nw128 _dafny.Array
					_ = _nw128
					if _len0_20.Cmp(_dafny.Zero) == 0 {
						_nw128 = _dafny.NewArray(_len0_20)
					} else {
						var _init20 func(_dafny.Int) bool = (func(_918_v3 bool) func(_dafny.Int) bool {
							return func(_919_i5 _dafny.Int) bool {
								return _918_v3
							}
						})(_842_v3)
						_ = _init20
						var _element0_20 = _init20(_dafny.Zero)
						_ = _element0_20
						_nw128 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
						(_nw128).ArraySet1(_element0_20, 0)
						var _nativeLen0_20 = (_len0_20).Int()
						_ = _nativeLen0_20
						for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
							(_nw128).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
						}
					}
					_917_v58 = _nw128
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_917_v58), 0))
					_ = _index133
					(_917_v58).ArraySet1(!(!(!(_842_v3))), (_index133).Int())
					var _920_v59 _dafny.Sequence
					_ = _920_v59
					_920_v59 = _dafny.UnicodeSeqOfUtf8Bytes("ji")
					var _921_v60 _dafny.Map
					_ = _921_v60
					_921_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v3, (_this).F28())
					var _922_v61 _dafny.Array
					_ = _922_v61
					var _nwElement0_24 _dafny.Sequence = _920_v59
					_ = _nwElement0_24
					var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(10))
					_ = _nw129
					(_nw129).ArraySet1(_nwElement0_24, 0)
					(_nw129).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(947))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg72 _dafny.Int) interface{} {
							return coer72(arg72)
						}
					}(func(_923_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					})), 1)
					(_nw129).ArraySet1(_920_v59, 2)
					(_nw129).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_920_v59, _920_v59), 3)
					(_nw129).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(949))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg73 _dafny.Int) interface{} {
							return coer73(arg73)
						}
					}((func(_924_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_925_i7 _dafny.Int) _dafny.CodePoint {
							return _924_v2
						}
					})(_841_v2))), 4)
					(_nw129).ArraySet1(_920_v59, 5)
					(_nw129).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_920_v59, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg74 _dafny.Int) interface{} {
							return coer74(arg74)
						}
					}((func(_926_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_927_i8 _dafny.Int) _dafny.CodePoint {
							return _926_v2
						}
					})(_841_v2)))), 6)
					(_nw129).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_920_v59, (Companion_Default___.SafeIndex((_921_v60).Cardinality(), _dafny.IntOfUint32((_920_v59).Cardinality()))).Uint32(), (_839_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))).Int())), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_844_v5).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_920_v59, (Companion_Default___.SafeIndex((_921_v60).Cardinality(), _dafny.IntOfUint32((_920_v59).Cardinality()))).Uint32(), (_839_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))).Int()))).Cardinality()))).Uint32(), _dafny.CodePoint('d')), 7)
					(_nw129).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xry"), 8)
					(_nw129).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg75 _dafny.Int) interface{} {
							return coer75(arg75)
						}
					}((func(_928_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_929_i9 _dafny.Int) _dafny.CodePoint {
							return _928_v2
						}
					})(_841_v2))), 9)
					_922_v61 = _nw129
					var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_922_v61), 0))
					_ = _index134
					(_922_v61).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg76 _dafny.Int) interface{} {
							return coer76(arg76)
						}
					}(func(_930_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('w')
					})), (_index134).Int())
					var _931_v62 D13
					_ = _931_v62
					_931_v62 = Companion_D13_.Create_DC26_(_920_v59, _842_v3, _842_v3)
					var _pat_let_tv19 = _842_v3
					_ = _pat_let_tv19
					var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))
					_ = _index135
					var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_917_v58), 0))
					_ = _index136
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_922_v61), 0))
					_ = _index137
					var _rhs110 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let14_0 D13) D13 {
						return func(_932_dt__update__tmp_h1 D13) D13 {
							return func(_pat_let15_0 bool) D13 {
								return func(_933_dt__update_hcf34_h0 bool) D13 {
									return Companion_D13_.Create_DC26_((_932_dt__update__tmp_h1).Dtor_cf33(), _933_dt__update_hcf34_h0, (_932_dt__update__tmp_h1).Dtor_cf35())
								}(_pat_let15_0)
							}(_pat_let_tv19)
						}(_pat_let14_0)
					}(_931_v62)).Dtor_cf34(), _915_v56)
					_ = _rhs110
					var _rhs111 _dafny.CodePoint = (func() _dafny.CodePoint {
						if (_842_v3) && (_842_v3) {
							return (_839_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))).Int())
						}
						return (_839_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))).Int())
					})()
					_ = _rhs111
					var _rhs112 bool = (Companion_Default___.Fm2(_843_v4, globalState)) && ((_931_v62).Dtor_cf35())
					_ = _rhs112
					var _rhs113 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_920_v59, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_920_v59).Cardinality()))).Uint32(), (_839_v1).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))).Int()))
					_ = _rhs113
					var _lhs109 _dafny.Array = _839_v1
					_ = _lhs109
					var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_839_v1), 0))
					_ = _lhs110
					var _lhs111 _dafny.Array = _917_v58
					_ = _lhs111
					var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_917_v58), 0))
					_ = _lhs112
					var _lhs113 _dafny.Array = _922_v61
					_ = _lhs113
					var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_922_v61), 0))
					_ = _lhs114
					_916_v57 = _rhs110
					(_lhs109).ArraySet1CodePoint(_rhs111, (_lhs110).Int())
					(_lhs111).ArraySet1(_rhs112, (_lhs112).Int())
					(_lhs113).ArraySet1(_rhs113, (_lhs114).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _934_v63 bool
		_ = _934_v63
		_934_v63 = true
		var _935_v64 _dafny.Sequence
		_ = _935_v64
		_935_v64 = _dafny.SeqOf(_934_v63, _934_v63)
		var _936_v65 _dafny.Map
		_ = _936_v65
		_936_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v0, _935_v64)
		var _937_v66 D16
		_ = _937_v66
		_937_v66 = Companion_D16_.Create_DC35_(_936_v65)
		var _938_v67 _dafny.Array
		_ = _938_v67
		var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
		_ = _nw130
		_938_v67 = _nw130
		var _939_v68 _dafny.MultiSet
		_ = _939_v68
		_939_v68 = _dafny.MultiSetOf(_934_v63, false)
		var _940_v69 D16
		_ = _940_v69
		_940_v69 = Companion_D16_.Create_DC36_(true, _939_v68, _934_v63)
		var _pat_let_tv20 = _934_v63
		_ = _pat_let_tv20
		var _pat_let_tv21 = _934_v63
		_ = _pat_let_tv21
		var _941_v70 _dafny.Array
		_ = _941_v70
		var _nwElement0_25 D16 = _940_v69
		_ = _nwElement0_25
		var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(15))
		_ = _nw131
		(_nw131).ArraySet1(_nwElement0_25, 0)
		(_nw131).ArraySet1(_940_v69, 1)
		(_nw131).ArraySet1(_940_v69, 2)
		(_nw131).ArraySet1(_940_v69, 3)
		(_nw131).ArraySet1(_940_v69, 4)
		(_nw131).ArraySet1(Companion_D16_.Create_DC36_(_934_v63, _939_v68, _934_v63), 5)
		(_nw131).ArraySet1(_940_v69, 6)
		(_nw131).ArraySet1(_940_v69, 7)
		(_nw131).ArraySet1(Companion_Default___.Fm40(globalState), 8)
		(_nw131).ArraySet1(_940_v69, 9)
		(_nw131).ArraySet1(func(_pat_let16_0 D16) D16 {
			return func(_942_dt__update__tmp_h2 D16) D16 {
				return func(_pat_let17_0 bool) D16 {
					return func(_943_dt__update_hcf49_h0 bool) D16 {
						return Companion_D16_.Create_DC36_((_942_dt__update__tmp_h2).Dtor_cf47(), (_942_dt__update__tmp_h2).Dtor_cf48(), _943_dt__update_hcf49_h0)
					}(_pat_let17_0)
				}(_pat_let_tv20)
			}(_pat_let16_0)
		}(_940_v69), 10)
		(_nw131).ArraySet1(func(_pat_let18_0 D16) D16 {
			return func(_944_dt__update__tmp_h3 D16) D16 {
				return func(_pat_let19_0 bool) D16 {
					return func(_945_dt__update_hcf47_h0 bool) D16 {
						return Companion_D16_.Create_DC36_(_945_dt__update_hcf47_h0, (_944_dt__update__tmp_h3).Dtor_cf48(), (_944_dt__update__tmp_h3).Dtor_cf49())
					}(_pat_let19_0)
				}(_pat_let_tv21)
			}(_pat_let18_0)
		}(_940_v69), 11)
		(_nw131).ArraySet1(_940_v69, 12)
		(_nw131).ArraySet1(Companion_Default___.Fm40(globalState), 13)
		(_nw131).ArraySet1(_940_v69, 14)
		_941_v70 = _nw131
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_938_v67), 0))
		_ = _index138
		(_938_v67).ArraySet1(_941_v70, (_index138).Int())
		var _946_v71 D17
		_ = _946_v71
		_946_v71 = Companion_D17_.Create_DC38_(_941_v70)
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_938_v67), 0))
		_ = _index139
		var _rhs114 D16 = _937_v66
		_ = _rhs114
		var _rhs115 _dafny.Int = (_this).F28()
		_ = _rhs115
		var _rhs116 _dafny.Array = (_946_v71).Dtor_cf51()
		_ = _rhs116
		var _lhs115 *GlobalState = globalState
		_ = _lhs115
		var _lhs116 _dafny.Array = _938_v67
		_ = _lhs116
		var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_938_v67), 0))
		_ = _lhs117
		_937_v66 = _rhs114
		_lhs115.F7 = _rhs115
		(_lhs116).ArraySet1(_rhs116, (_lhs117).Int())
		var _947_v72 _dafny.CodePoint
		_ = _947_v72
		_947_v72 = _dafny.CodePoint('p')
		_947_v72 = _947_v72
		(globalState).F16 = (_this).F28()
		(globalState).F7 = (_dafny.SetOf(_934_v63, (func() bool {
			if _934_v63 {
				return _934_v63
			}
			return _934_v63
		})(), _934_v63, _934_v63)).Cardinality()
		var _948_v73 D15
		_ = _948_v73
		_948_v73 = Companion_D15_.Create_DC33_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F28()), (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F28())).Cardinality()))).Uint32(), (_this).F28()))
		var _949_v74 D15
		_ = _949_v74
		_949_v74 = Companion_D15_.Create_DC34_(_948_v73)
		var _950_v75 _dafny.Map
		_ = _950_v75
		_950_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_949_v74, Companion_Default___.Fm0(_dafny.IntOfInt64(-164), (_this).F28(), (_this).F28(), true, globalState))
		r1 = !(_934_v63) || (((_950_v75).Cardinality()).Cmp((_this).F28()) != 0)
		r0 = _934_v63
		r1 = !(_934_v63)
		var _951_v76 D4
		_ = _951_v76
		_951_v76 = Companion_D4_.Create_DC9_((_this).F28())
		r2 = (_this).Fm6((_951_v76).Dtor_cf11(), globalState)
		var _952_v77 _dafny.Map
		_ = _952_v77
		_952_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _934_v63)
		var _953_v78 _dafny.Sequence
		_ = _953_v78
		_953_v78 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _934_v63), _952_v77)
		var _954_v79 T1
		_ = _954_v79
		var _nw132 *C1 = New_C1_()
		_ = _nw132
		_nw132.Ctor__((_this).F28(), _953_v78)
		_954_v79 = _nw132
		var _955_v80 _dafny.Sequence
		_ = _955_v80
		_955_v80 = _dafny.SeqOf(_954_v79)
		r3 = !(!((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_955_v80, _955_v80)).Cardinality())).Cmp(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-779)), (_this).F28())) < 0))
		return r0, r1, r2, r3
	}
}
func (_this *C2) F28() _dafny.Int {
	{
		return _this._f28
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	dummy byte
}

func New_C3_() *C3 {
	_this := C3{}

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

func (_this *C3) Ctor__() {
	{
	}
}
func (_this *C3) Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(523)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false))).Cardinality())) == 0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(992))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg77 _dafny.Int) interface{} {
				return coer77(arg77)
			}
		}(func(_956_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))
	}
}
func (_this *C3) Fm6(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C3) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _957_v0 _dafny.Map
		_ = _957_v0
		_957_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _958_v1 _dafny.Map
		_ = _958_v1
		_958_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_957_v0).Cardinality(), p0)
		var _959_v2 _dafny.MultiSet
		_ = _959_v2
		_959_v2 = _dafny.MultiSetOf(_957_v0, _957_v0)
		var _960_v4 bool
		_ = _960_v4
		_960_v4 = false
		var _961_v5 _dafny.MultiSet
		_ = _961_v5
		_961_v5 = _dafny.MultiSetOf(true, _960_v4)
		var _962_v6 _dafny.Map
		_ = _962_v6
		_962_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((func() _dafny.Set {
			var _coll43 = _dafny.NewBuilder()
			_ = _coll43
			for _iter51 := _dafny.Iterate((_959_v2).Elements()); ; {
				_compr_43, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _963_v3 _dafny.Map
				_963_v3 = interface{}(_compr_43).(_dafny.Map)
				if (_959_v2).Contains(_963_v3) {
					_coll43.Add(_963_v3)
				}
			}
			return _coll43.ToSet()
		}()).Cardinality()).Times(p0)), (_961_v5).Union(_961_v5))
		var _964_v7 _dafny.Sequence
		_ = _964_v7
		_964_v7 = _dafny.SeqOf(_960_v4)
		var _965_v8 _dafny.Sequence
		_ = _965_v8
		_965_v8 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_964_v7))
		_962_v6 = (_962_v6).Update(p0, (_965_v8).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_965_v8).Cardinality()))).Uint32()).(_dafny.MultiSet))
		if _960_v4 {
			var _966_v9 _dafny.Array
			_ = _966_v9
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_21
			var _nw133 _dafny.Array
			_ = _nw133
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw133 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_967_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_968_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_968_i0, _967_p0)
					}
				})(p0)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw133 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw133).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw133).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_966_v9 = _nw133
			var _969_v10 _dafny.Sequence
			_ = _969_v10
			_969_v10 = _dafny.SeqOf(_966_v9, _966_v9)
			var _970_v11 _dafny.Map
			_ = _970_v11
			_970_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v4, _960_v4)
			var _971_v12 D4
			_ = _971_v12
			_971_v12 = Companion_D4_.Create_DC10_(p0, (_970_v11).Cardinality(), _960_v4, _960_v4, p0)
			var _972_v13 _dafny.Sequence
			_ = _972_v13
			_972_v13 = _dafny.UnicodeSeqOfUtf8Bytes("eofajvq")
			var _973_v14 D6
			_ = _973_v14
			_973_v14 = Companion_D6_.Create_DC12_(_969_v10)
			var _974_v15 _dafny.Sequence
			_ = _974_v15
			_974_v15 = _dafny.SeqOf((_973_v14).Dtor_cf18(), _969_v10, _dafny.Companion_Sequence_.Concatenate(_969_v10, _969_v10))
			var _rhs117 bool = (_971_v12).Dtor_cf14()
			_ = _rhs117
			var _rhs118 bool = !(((_dafny.IntOfUint32((_972_v13).Cardinality())).Times((_970_v11).Cardinality())).Cmp(p0) <= 0)
			_ = _rhs118
			var _rhs119 _dafny.Sequence = (_974_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_974_v15).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs119
			var _lhs118 *GlobalState = globalState
			_ = _lhs118
			r1 = _rhs117
			_lhs118.F0 = _rhs118
			_969_v10 = _rhs119
			(globalState).F16 = p0
			var _975_v16 _dafny.Array
			_ = _975_v16
			var _nw134 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw134
			_975_v16 = _nw134
			var _976_v17 _dafny.Set
			_ = _976_v17
			_976_v17 = _dafny.SetOf(_975_v16, _975_v16, _975_v16)
			var _rhs120 _dafny.Int = p0
			_ = _rhs120
			var _rhs121 _dafny.Int = (func() _dafny.Int {
				if _960_v4 {
					return p0
				}
				return (_dafny.Zero).Minus((_dafny.IntOfInt64(420)).Times(p0))
			})()
			_ = _rhs121
			var _rhs122 _dafny.Sequence = _964_v7
			_ = _rhs122
			var _rhs123 _dafny.Set = _976_v17
			_ = _rhs123
			var _rhs124 _dafny.Array = _966_v9
			_ = _rhs124
			var _lhs119 *GlobalState = globalState
			_ = _lhs119
			var _lhs120 *GlobalState = globalState
			_ = _lhs120
			_lhs119.F16 = _rhs120
			_lhs120.F7 = _rhs121
			_964_v7 = _rhs122
			_976_v17 = _rhs123
			r0 = _rhs124
			var _977_v18 *C0
			_ = _977_v18
			var _nw135 *C0 = New_C0_()
			_ = _nw135
			_nw135.Ctor__(_960_v4)
			_977_v18 = _nw135
			var _pat_let_tv22 = _969_v10
			_ = _pat_let_tv22
			var _pat_let_tv23 = _969_v10
			_ = _pat_let_tv23
			var _pat_let_tv24 = p0
			_ = _pat_let_tv24
			var _pat_let_tv25 = _969_v10
			_ = _pat_let_tv25
			var _pat_let_tv26 = _969_v10
			_ = _pat_let_tv26
			var _pat_let_tv27 = _966_v9
			_ = _pat_let_tv27
			_973_v14 = func(_pat_let20_0 D6) D6 {
				return func(_978_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let21_0 _dafny.Sequence) D6 {
						return func(_979_dt__update_hcf18_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC12_(_979_dt__update_hcf18_h0)
						}(_pat_let21_0)
					}(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv22, _pat_let_tv23), (Companion_Default___.SafeIndex(_pat_let_tv24, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_pat_let_tv25, _pat_let_tv26)).Cardinality()))).Uint32(), _pat_let_tv27))
				}(_pat_let20_0)
			}(_973_v14)
		} else {
			(globalState).F7 = _dafny.IntOfInt64(639)
			var _980_v19 _dafny.Map
			_ = _980_v19
			_980_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v4, (p0).Cmp(Companion_Default___.Fm0(p0, p0, p0, _960_v4, globalState)) >= 0)
			_980_v19 = (_980_v19).Update(_960_v4, _960_v4)
			(globalState).F7 = (_dafny.Zero).Minus(p0)
			(globalState).F21 = (p0).Cmp(p0) > 0
			var _981_v20 _dafny.Set
			_ = _981_v20
			_981_v20 = _dafny.SetOf(_960_v4, !(_960_v4) || (false), _960_v4, (func() bool {
				if _960_v4 {
					return _960_v4
				}
				return _960_v4
			})(), _960_v4)
			_981_v20 = _dafny.SetOf(_960_v4)
		}
		(globalState).F21 = _960_v4
		var _982_v21 D0
		_ = _982_v21
		_982_v21 = Companion_D0_.Create_DC0_(_960_v4)
		var _pat_let_tv28 = _960_v4
		_ = _pat_let_tv28
		var _pat_let_tv29 = _960_v4
		_ = _pat_let_tv29
		if func(_source16 D0) bool {
			if _source16.Is_DC0() {
				var _983___mcc_h0 bool = _source16.Get_().(D0_DC0).Cf0
				_ = _983___mcc_h0
				var _984_cf0 bool = _983___mcc_h0
				_ = _984_cf0
				if _pat_let_tv28 {
					return _984_cf0
				} else {
					return _pat_let_tv29
				}
			} else {
				var _985___mcc_h1 D0 = _source16.Get_().(D0_DC1).Cf1
				_ = _985___mcc_h1
				var _986_cf1 D0 = _985___mcc_h1
				_ = _986_cf1
				return true
			}
		}(_982_v21) {
			var _987_v22 _dafny.CodePoint
			_ = _987_v22
			_987_v22 = _dafny.CodePoint('e')
			_987_v22 = _987_v22
			var _988_v23 *C0
			_ = _988_v23
			var _nw136 *C0 = New_C0_()
			_ = _nw136
			_nw136.Ctor__(_960_v4)
			_988_v23 = _nw136
			(globalState).F2 = (true) == (_988_v23.F27)
			var _989_v24 *C0
			_ = _989_v24
			var _nw137 *C0 = New_C0_()
			_ = _nw137
			_nw137.Ctor__((_988_v23.F27) == (_988_v23.F27))
			_989_v24 = _nw137
			var _990_v25 _dafny.Array
			_ = _990_v25
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_22
			var _nw138 _dafny.Array
			_ = _nw138
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw138 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = (func(_991_v23 *C0, _992_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_993_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_993_i1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v23.F27, Companion_D4_.Create_DC9_(_992_p0))).Cardinality())
					}
				})(_988_v23, p0)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw138 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw138).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw138).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_990_v25 = _nw138
			var _994_v26 _dafny.Sequence
			_ = _994_v26
			_994_v26 = _dafny.SeqOf(p0, p0)
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_990_v25), 0))
			_ = _index140
			(_990_v25).ArraySet1((_dafny.IntOfUint32((_994_v26).Cardinality())).Minus(p0), (_index140).Int())
			var _995_v27 _dafny.Set
			_ = _995_v27
			_995_v27 = _dafny.SetOf(_960_v4)
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_990_v25), 0))
			_ = _index141
			var _rhs125 bool = (_995_v27).IsSubsetOf((_995_v27).Difference(_995_v27))
			_ = _rhs125
			var _rhs126 _dafny.Int = Companion_Default___.Fm0((p0).Minus(p0), (_dafny.MultiSetOf(true)).Cardinality(), p0, _989_v24.F27, globalState)
			_ = _rhs126
			var _rhs127 bool = _988_v23.F27
			_ = _rhs127
			var _lhs121 *GlobalState = globalState
			_ = _lhs121
			var _lhs122 _dafny.Array = _990_v25
			_ = _lhs122
			var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_990_v25), 0))
			_ = _lhs123
			var _lhs124 *GlobalState = globalState
			_ = _lhs124
			_lhs121.F21 = _rhs125
			(_lhs122).ArraySet1(_rhs126, (_lhs123).Int())
			_lhs124.F2 = _rhs127
		} else {
			var _996_v28 _dafny.Sequence
			_ = _996_v28
			_996_v28 = _dafny.UnicodeSeqOfUtf8Bytes("tsxxruaee")
			(globalState).F19 = _dafny.Companion_Sequence_.Concatenate(_996_v28, _dafny.Companion_Sequence_.Concatenate(_996_v28, _996_v28))
			var _997_v29 _dafny.MultiSet
			_ = _997_v29
			_997_v29 = _dafny.MultiSetOf(p0, p0, p0, p0, p0)
			var _998_v30 _dafny.Sequence
			_ = _998_v30
			_998_v30 = _dafny.SeqOf(p0, p0)
			var _999_v31 D7
			_ = _999_v31
			_999_v31 = Companion_D7_.Create_DC15_(_998_v30)
			var _1000_v32 D7
			_ = _1000_v32
			_1000_v32 = Companion_D7_.Create_DC15_((_999_v31).Dtor_cf22())
			var _1001_v33 _dafny.Array
			_ = _1001_v33
			var _nwElement0_26 _dafny.MultiSet = _997_v29
			_ = _nwElement0_26
			var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(22))
			_ = _nw139
			(_nw139).ArraySet1(_nwElement0_26, 0)
			(_nw139).ArraySet1(_997_v29, 1)
			(_nw139).ArraySet1(_997_v29, 2)
			(_nw139).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(p0)), 3)
			(_nw139).ArraySet1(_997_v29, 4)
			(_nw139).ArraySet1(_dafny.MultiSetOf(p0, _dafny.IntOfInt64(790)), 5)
			(_nw139).ArraySet1(_997_v29, 6)
			(_nw139).ArraySet1(_997_v29, 7)
			(_nw139).ArraySet1(_dafny.MultiSetOf(p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), p0, _dafny.IntOfUint32((_996_v28).Cardinality())), 8)
			(_nw139).ArraySet1((_dafny.MultiSetFromSeq((_999_v31).Dtor_cf22())).Union(_997_v29), 9)
			(_nw139).ArraySet1((func() _dafny.MultiSet {
				if _960_v4 {
					return _997_v29
				}
				return _997_v29
			})(), 10)
			(_nw139).ArraySet1(_997_v29, 11)
			(_nw139).ArraySet1((_997_v29).Union(_997_v29), 12)
			(_nw139).ArraySet1((_997_v29).Difference(_997_v29), 13)
			(_nw139).ArraySet1(_997_v29, 14)
			(_nw139).ArraySet1((_997_v29).Intersection(_997_v29), 15)
			(_nw139).ArraySet1(_dafny.MultiSetOf(p0), 16)
			(_nw139).ArraySet1(_997_v29, 17)
			(_nw139).ArraySet1(_997_v29, 18)
			(_nw139).ArraySet1(((_997_v29).Update(p0, Companion_Default___.Abs(_dafny.IntOfUint32((_996_v28).Cardinality())))).Difference(_997_v29), 19)
			(_nw139).ArraySet1(_997_v29, 20)
			(_nw139).ArraySet1((Companion_Default___.Fm12(_996_v28, _dafny.IntOfInt64(453), _960_v4, globalState)).Union(_dafny.MultiSetFromSeq(_998_v30)), 21)
			_1001_v33 = _nw139
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_1001_v33), 0))
			_ = _index142
			(_1001_v33).ArraySet1((_997_v29).Intersection(_997_v29), (_index142).Int())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_1001_v33), 0))
			_ = _index143
			(_1001_v33).ArraySet1(_dafny.MultiSetFromSeq(_998_v30), (_index143).Int())
			(globalState).F2 = _960_v4
			var _1002_v34 _dafny.Array
			_ = _1002_v34
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_23
			var _nw140 _dafny.Array
			_ = _nw140
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw140 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.Int = (func(_1003_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1004_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1004_i2, _1003_p0)
					}
				})(p0)
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw140 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw140).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw140).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_1002_v34 = _nw140
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_1002_v34), 0))
			_ = _index144
			(_1002_v34).ArraySet1(p0, (_index144).Int())
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_1002_v34), 0))
			_ = _index145
			(_1002_v34).ArraySet1(_dafny.IntOfInt64(62), (_index145).Int())
			(globalState).F13 = _960_v4
		}
		var _1005_v35 _dafny.CodePoint
		_ = _1005_v35
		_1005_v35 = _dafny.CodePoint('b')
		var _1006_v36 _dafny.MultiSet
		_ = _1006_v36
		_1006_v36 = _dafny.MultiSetOf(_1005_v35, _dafny.CodePoint('c'), _1005_v35, _dafny.CodePoint('v'))
		var _1007_v37 _dafny.Sequence
		_ = _1007_v37
		_1007_v37 = _dafny.UnicodeSeqOfUtf8Bytes("xysyc")
		var _1008_v38 D4
		_ = _1008_v38
		_1008_v38 = Companion_D4_.Create_DC10_(p0, _dafny.IntOfInt64(30), !(_960_v4), _960_v4, (_dafny.Zero).Minus(p0))
		var _1009_v39 _dafny.Sequence
		_ = _1009_v39
		_1009_v39 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_1007_v37), _1006_v36)
		var _1010_v40 _dafny.Array
		_ = _1010_v40
		var _nwElement0_27 _dafny.MultiSet = _1006_v36
		_ = _nwElement0_27
		var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(12))
		_ = _nw141
		(_nw141).ArraySet1(_nwElement0_27, 0)
		(_nw141).ArraySet1(_1006_v36, 1)
		(_nw141).ArraySet1(_1006_v36, 2)
		(_nw141).ArraySet1(_1006_v36, 3)
		(_nw141).ArraySet1(_1006_v36, 4)
		(_nw141).ArraySet1(Companion_Default___.Fm13(_dafny.IntOfUint32((_1007_v37).Cardinality()), !(_960_v4), _1008_v38, _960_v4, globalState), 5)
		(_nw141).ArraySet1(_1006_v36, 6)
		(_nw141).ArraySet1(_1006_v36, 7)
		(_nw141).ArraySet1((_1006_v36).Union((_1009_v39).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1009_v39).Cardinality()))).Uint32()).(_dafny.MultiSet)), 8)
		(_nw141).ArraySet1((_1006_v36).Intersection(_1006_v36), 9)
		(_nw141).ArraySet1((func() _dafny.MultiSet {
			if _960_v4 {
				return _1006_v36
			}
			return _1006_v36
		})(), 10)
		(_nw141).ArraySet1(_1006_v36, 11)
		_1010_v40 = _nw141
		for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1010_v40), 0))); ; {
			_guard_loop_7, _ok52 := _iter52()
			if !_ok52 {
				break
			}
			var _1011_i3 _dafny.Int
			_1011_i3 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_1011_i3).Sign() != -1) && ((_1011_i3).Cmp(_dafny.ArrayLen((_1010_v40), 0)) < 0)) {
				(_1010_v40).ArraySet1(_dafny.MultiSetOf(_1005_v35), (_1011_i3).Int())
			}
		}
		var _1012_v41 _dafny.Array
		_ = _1012_v41
		var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
		_ = _nw142
		_1012_v41 = _nw142
		var _1013_v42 D7
		_ = _1013_v42
		_1013_v42 = Companion_D7_.Create_DC16_(_1012_v41)
		var _1014_v43 D7
		_ = _1014_v43
		_1014_v43 = Companion_D7_.Create_DC17_(_1013_v42)
		var _1015_v44 D7
		_ = _1015_v44
		_1015_v44 = Companion_D7_.Create_DC17_(_1014_v43)
		var _source17 D7 = _1015_v44
		_ = _source17
		if _source17.Is_DC16() {
			var _1016___mcc_h2 _dafny.Array = _source17.Get_().(D7_DC16).Cf23
			_ = _1016___mcc_h2
			var _1017_cf23 _dafny.Array = _1016___mcc_h2
			_ = _1017_cf23
			var _1018_v45 _dafny.Map
			_ = _1018_v45
			_1018_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v37, p0)
			_1018_v45 = (_1018_v45).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jyay"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jyay")).Cardinality()))).Uint32(), _1005_v35), _1007_v37), p0)
			var _1019_v46 D6
			_ = _1019_v46
			_1019_v46 = Companion_D6_.Create_DC13_(p0, p0)
			(globalState).F21 = (_964_v7).Select((Companion_Default___.SafeIndex((_1019_v46).Dtor_cf20(), _dafny.IntOfUint32((_964_v7).Cardinality()))).Uint32()).(bool)
			var _1020_v47 _dafny.Map
			_ = _1020_v47
			_1020_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v4, _dafny.UnicodeSeqOfUtf8Bytes("eaxlkubrh"))
			var _1021_v48 _dafny.Map
			_ = _1021_v48
			_1021_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_1020_v47).Contains(_960_v4) {
					return (_1020_v47).Get(_960_v4).(_dafny.Sequence)
				}
				return _1007_v37
			})(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1020_v47).Contains(_960_v4) {
					return (_1020_v47).Get(_960_v4).(_dafny.Sequence)
				}
				return _1007_v37
			})()).Cardinality()))).Uint32(), _1005_v35), _1007_v37), (_964_v7).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_dafny.IntOfInt64(-147), p0, p0, _960_v4, globalState), _dafny.IntOfUint32((_964_v7).Cardinality()))).Uint32()).(bool))
			_1021_v48 = (_1021_v48).Update(_1007_v37, (p0).Cmp(p0) >= 0)
			if _960_v4 {
				(globalState).F2 = _960_v4
				var _1022_v49 _dafny.Array
				_ = _1022_v49
				var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw143
				_1022_v49 = _nw143
				var _1023_v50 _dafny.Array
				_ = _1023_v50
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_24
				var _nw144 _dafny.Array
				_ = _nw144
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw144 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Int = (func(_1024_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1025_i4 _dafny.Int) _dafny.Int {
							return (_1025_i4).Minus(_1024_p0)
						}
					})(p0)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw144 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw144).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw144).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1023_v50 = _nw144
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1022_v49), 0))
				_ = _index146
				(_1022_v49).ArraySet1(_1023_v50, (_index146).Int())
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_1022_v49), 0))
				_ = _index147
				(_1022_v49).ArraySet1(_1023_v50, (_index147).Int())
				(globalState).F21 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nxdew"), _1007_v37), _1007_v37)
				var _1026_v51 _dafny.Set
				_ = _1026_v51
				_1026_v51 = _dafny.SetOf(_960_v4, _960_v4, _960_v4)
				var _1027_v52 _dafny.Set
				_ = _1027_v52
				_1027_v52 = _dafny.SetOf(!(_1026_v51).Equals(Companion_Default___.Fm14((_dafny.Zero).Minus(p0), _960_v4, globalState)), _960_v4, _960_v4, (p0).Cmp(_dafny.IntOfInt64(580)) >= 0)
				_1027_v52 = (_dafny.SetOf(_960_v4, _960_v4, _960_v4)).Union((_dafny.SetOf(!(_960_v4), _960_v4, _960_v4)).Intersection(_dafny.SetOf(_960_v4)))
				var _1028_v53 _dafny.MultiSet
				_ = _1028_v53
				_1028_v53 = _dafny.MultiSetOf((_dafny.MultiSetOf(p0, _dafny.IntOfInt64(-127))).Cardinality())
				var _1029_v54 _dafny.Sequence
				_ = _1029_v54
				_1029_v54 = _dafny.SeqOf(_1023_v50)
				var _1030_v55 _dafny.Set
				_ = _1030_v55
				_1030_v55 = _dafny.SetOf(p0, _dafny.IntOfInt64(318))
				var _1031_v56 _dafny.Map
				_ = _1031_v56
				_1031_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1028_v53, (_dafny.SetOf(p0, _dafny.IntOfUint32((_1029_v54).Cardinality()), Companion_Default___.Fm0((_dafny.Zero).Minus(p0), p0, p0, _960_v4, globalState), p0, p0)).IsDisjointFrom(_1030_v55))
				var _1032_v57 _dafny.Sequence
				_ = _1032_v57
				_1032_v57 = _964_v7
				var _1033_v58 *C0
				_ = _1033_v58
				var _nw145 *C0 = New_C0_()
				_ = _nw145
				_nw145.Ctor__(_960_v4)
				_1033_v58 = _nw145
				var _1034_v59 _dafny.Map
				_ = _1034_v59
				_1034_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1033_v58, _1028_v53)
				_1031_v56 = (_1031_v56).Update(_dafny.MultiSetOf(_dafny.IntOfUint32((_1032_v57).Cardinality()), p0, (_1034_v59).Cardinality()), (_this).Fm6(_dafny.IntOfUint32((_1007_v37).Cardinality()), globalState))
			} else {
				var _1035_v60 _dafny.Array
				_ = _1035_v60
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_25
				var _nw146 _dafny.Array
				_ = _nw146
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw146 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Int = (func(_1036_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1037_i5 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1037_i5, (Companion_D4_.Create_DC9_(_1036_p0)).Dtor_cf11())
						}
					})(p0)
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw146 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw146).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw146).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1035_v60 = _nw146
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1035_v60), 0))
				_ = _index148
				(_1035_v60).ArraySet1(p0, (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1035_v60), 0))
				_ = _index149
				(_1035_v60).ArraySet1(p0, (_index149).Int())
				var _1038_v61 _dafny.Sequence
				_ = _1038_v61
				_1038_v61 = _dafny.SeqOf((_1008_v38).Dtor_cf16())
				_1005_v35 = Companion_Default___.Fm15(_dafny.Companion_Sequence_.Concatenate(_1038_v61, _1038_v61), p0, globalState)
				var _1039_v62 _dafny.Map
				_ = _1039_v62
				_1039_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v4, _1035_v60)
				var _1040_v63 _dafny.Map
				_ = _1040_v63
				_1040_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1005_v35, p0)
				var _1041_v64 _dafny.MultiSet
				_ = _1041_v64
				_1041_v64 = _dafny.MultiSetOf(_dafny.IntOfInt64(372), (func() _dafny.Int {
					if (_1040_v63).Contains(_1005_v35) {
						return (_1040_v63).Get(_1005_v35).(_dafny.Int)
					}
					return (_1035_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1035_v60), 0))).Int()).(_dafny.Int)
				})())
				var _1042_v65 _dafny.Map
				_ = _1042_v65
				_1042_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v4, _1041_v64)
				var _rhs128 bool = (Companion_Default___.Fm0(p0, (_1042_v65).Cardinality(), (_1035_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1035_v60), 0))).Int()).(_dafny.Int), _960_v4, globalState)).Cmp(Companion_Default___.SafeModuloInt((_961_v5).Cardinality(), _dafny.IntOfUint32((_1007_v37).Cardinality()))) >= 0
				_ = _rhs128
				var _rhs129 _dafny.MultiSet = _961_v5
				_ = _rhs129
				var _rhs130 _dafny.Map = (_1039_v62).Merge(_1039_v62)
				_ = _rhs130
				var _lhs125 *GlobalState = globalState
				_ = _lhs125
				_lhs125.F21 = _rhs128
				_961_v5 = _rhs129
				_1039_v62 = _rhs130
				(globalState).F7 = p0
				(globalState).F0 = _960_v4
			}
		} else if _source17.Is_DC15() {
			var _1043___mcc_h3 _dafny.Sequence = _source17.Get_().(D7_DC15).Cf22
			_ = _1043___mcc_h3
			var _1044_cf22 _dafny.Sequence = _1043___mcc_h3
			_ = _1044_cf22
			var _1045_v66 _dafny.MultiSet
			_ = _1045_v66
			_1045_v66 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg78 _dafny.Int) interface{} {
					return coer78(arg78)
				}
			}((func(_1046_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1047_i6 _dafny.Int) _dafny.CodePoint {
					return _1046_v35
				}
			})(_1005_v35))), _dafny.UnicodeSeqOfUtf8Bytes("tbsqpray"), _1007_v37)
			(globalState).F7 = (_1045_v66).Cardinality()
			(globalState).F0 = (_964_v7).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_964_v7).Cardinality()))).Uint32()).(bool)
			(globalState).F7 = p0
			var _1048_v67 _dafny.Array
			_ = _1048_v67
			var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw147
			_1048_v67 = _nw147
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_1048_v67), 0))
			_ = _index150
			(_1048_v67).ArraySet1(p0, (_index150).Int())
			var _1049_v68 _dafny.Array
			_ = _1049_v68
			var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
			_ = _nw148
			_1049_v68 = _nw148
			var _1050_v69 _dafny.Map
			_ = _1050_v69
			_1050_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_960_v4), _1048_v67)
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1049_v68), 0))
			_ = _index151
			(_1049_v68).ArraySet1(_1050_v69, (_index151).Int())
			var _1051_v70 _dafny.Sequence
			_ = _1051_v70
			_1051_v70 = _dafny.SeqOf(_1050_v69)
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_1048_v67), 0))
			_ = _index152
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1049_v68), 0))
			_ = _index153
			var _rhs131 _dafny.Int = ((Companion_Default___.Fm14(p0, _960_v4, globalState)).Cardinality()).Plus((_dafny.Zero).Minus(p0))
			_ = _rhs131
			var _rhs132 _dafny.Sequence = _1044_cf22
			_ = _rhs132
			var _rhs133 _dafny.Map = (_1051_v70).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(p0, p0, p0, _960_v4, globalState), _dafny.IntOfUint32((_1051_v70).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs133
			var _lhs126 _dafny.Array = _1048_v67
			_ = _lhs126
			var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_1048_v67), 0))
			_ = _lhs127
			var _lhs128 _dafny.Array = _1049_v68
			_ = _lhs128
			var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1049_v68), 0))
			_ = _lhs129
			(_lhs126).ArraySet1(_rhs131, (_lhs127).Int())
			_1044_cf22 = _rhs132
			(_lhs128).ArraySet1(_rhs133, (_lhs129).Int())
		} else {
			var _1052___mcc_h4 D7 = _source17.Get_().(D7_DC17).Cf24
			_ = _1052___mcc_h4
			var _1053_cf24 D7 = _1052___mcc_h4
			_ = _1053_cf24
			if _960_v4 {
				(globalState).F21 = _dafny.Companion_Sequence_.IsPrefixOf(_1007_v37, _1007_v37)
				var _1054_v71 _dafny.Map
				_ = _1054_v71
				_1054_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v4, _960_v4)
				var _1055_v72 _dafny.Set
				_ = _1055_v72
				_1055_v72 = _dafny.SetOf(_1054_v71)
				var _1056_v73 *C0
				_ = _1056_v73
				var _nw149 *C0 = New_C0_()
				_ = _nw149
				_nw149.Ctor__(_960_v4)
				_1056_v73 = _nw149
				var _1057_v74 _dafny.Array
				_ = _1057_v74
				var _nwElement0_28 *C0 = _1056_v73
				_ = _nwElement0_28
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(24))
				_ = _nw150
				(_nw150).ArraySet1(_nwElement0_28, 0)
				(_nw150).ArraySet1(_1056_v73, 1)
				(_nw150).ArraySet1(_1056_v73, 2)
				(_nw150).ArraySet1(_1056_v73, 3)
				(_nw150).ArraySet1(_1056_v73, 4)
				(_nw150).ArraySet1(_1056_v73, 5)
				(_nw150).ArraySet1(_1056_v73, 6)
				(_nw150).ArraySet1(_1056_v73, 7)
				(_nw150).ArraySet1(_1056_v73, 8)
				(_nw150).ArraySet1(_1056_v73, 9)
				(_nw150).ArraySet1(_1056_v73, 10)
				(_nw150).ArraySet1(_1056_v73, 11)
				(_nw150).ArraySet1(_1056_v73, 12)
				(_nw150).ArraySet1(_1056_v73, 13)
				(_nw150).ArraySet1(_1056_v73, 14)
				(_nw150).ArraySet1(_1056_v73, 15)
				(_nw150).ArraySet1(_1056_v73, 16)
				(_nw150).ArraySet1(_1056_v73, 17)
				(_nw150).ArraySet1(_1056_v73, 18)
				(_nw150).ArraySet1(_1056_v73, 19)
				(_nw150).ArraySet1(_1056_v73, 20)
				(_nw150).ArraySet1(_1056_v73, 21)
				(_nw150).ArraySet1(_1056_v73, 22)
				(_nw150).ArraySet1(_1056_v73, 23)
				_1057_v74 = _nw150
				var _1058_v75 _dafny.Sequence
				_ = _1058_v75
				_1058_v75 = _dafny.SeqOf(_1057_v74, _1057_v74)
				var _rhs134 _dafny.Set = (_1055_v72).Intersection((_1055_v72).Intersection(_1055_v72))
				_ = _rhs134
				var _rhs135 _dafny.Sequence = _1007_v37
				_ = _rhs135
				var _rhs136 bool = _960_v4
				_ = _rhs136
				var _rhs137 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1057_v74, _1057_v74), _1058_v75), _1058_v75)).Cardinality())
				_ = _rhs137
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				var _lhs132 *GlobalState = globalState
				_ = _lhs132
				_1055_v72 = _rhs134
				_lhs130.F8 = _rhs135
				_lhs131.F2 = _rhs136
				_lhs132.F7 = _rhs137
				(globalState).F13 = _1056_v73.F27
				var _1059_v76 _dafny.Array
				_ = _1059_v76
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_26
				var _nw151 _dafny.Array
				_ = _nw151
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw151 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = (func(_1060_v73 *C0) func(_dafny.Int) bool {
						return func(_1061_i7 _dafny.Int) bool {
							return _1060_v73.F27
						}
					})(_1056_v73)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw151 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw151).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw151).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1059_v76 = _nw151
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1059_v76), 0))
				_ = _index154
				(_1059_v76).ArraySet1(Companion_Default___.Fm2(_1054_v71, globalState), (_index154).Int())
				var _1062_v77 _dafny.Set
				_ = _1062_v77
				_1062_v77 = _dafny.SetOf(p0)
				var _1063_v78 _dafny.MultiSet
				_ = _1063_v78
				_1063_v78 = _dafny.MultiSetOf(_1062_v77, (_1062_v77).Union(_dafny.SetOf(p0)), (_dafny.SetOf(p0, p0)).Union(_1062_v77), _1062_v77)
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1059_v76), 0))
				_ = _index155
				var _rhs138 _dafny.Int = _dafny.IntOfInt64(899)
				_ = _rhs138
				var _rhs139 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_1007_v37, _1007_v37)
				_ = _rhs139
				var _rhs140 _dafny.MultiSet = _1063_v78
				_ = _rhs140
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				var _lhs134 _dafny.Array = _1059_v76
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1059_v76), 0))
				_ = _lhs135
				_lhs133.F16 = _rhs138
				(_lhs134).ArraySet1(_rhs139, (_lhs135).Int())
				_1063_v78 = _rhs140
				var _1064_v79 *C0
				_ = _1064_v79
				var _nw152 *C0 = New_C0_()
				_ = _nw152
				_nw152.Ctor__(_960_v4)
				_1064_v79 = _nw152
			} else {
				_1008_v38 = _1008_v38
				(globalState).F13 = _960_v4
				(globalState).F8 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("d"), _1007_v37)
				_1005_v35 = _1005_v35
				var _1065_v80 _dafny.Sequence
				_ = _1065_v80
				_1065_v80 = _dafny.SeqOf(_964_v7, _964_v7)
				(globalState).F7 = Companion_Default___.Fm0(_dafny.IntOfUint32((_1065_v80).Cardinality()), (p0).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(94))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}((func(_1066_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1067_i8 _dafny.Int) _dafny.Int {
						return _1066_p0
					}
				})(p0)))).Cardinality())), (func() _dafny.Int {
					if (_961_v5).Contains(false) {
						return (_961_v5).Multiplicity(false)
					}
					return p0
				})(), !((_960_v4) || (_960_v4)), globalState)
			}
			var _1068_v81 *C0
			_ = _1068_v81
			var _nw153 *C0 = New_C0_()
			_ = _nw153
			_nw153.Ctor__((_960_v4) || (_960_v4))
			_1068_v81 = _nw153
			var _1069_v82 _dafny.Array
			_ = _1069_v82
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_27
			var _nw154 _dafny.Array
			_ = _nw154
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw154 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Int = (func(_1070_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1071_i9 _dafny.Int) _dafny.Int {
						return (_1071_i9).Minus(_1070_p0)
					}
				})(p0)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw154 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw154).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw154).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1069_v82 = _nw154
			var _1072_v83 _dafny.Set
			_ = _1072_v83
			_1072_v83 = _dafny.SetOf(_1069_v82)
			var _1073_v84 _dafny.Array
			_ = _1073_v84
			var _nw155 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw155
			_1073_v84 = _nw155
			var _rhs141 _dafny.Set = _1072_v83
			_ = _rhs141
			var _rhs142 _dafny.CodePoint = _1005_v35
			_ = _rhs142
			var _rhs143 _dafny.Array = _1073_v84
			_ = _rhs143
			_1072_v83 = _rhs141
			_1005_v35 = _rhs142
			_1073_v84 = _rhs143
			var _1074_v85 _dafny.Array
			_ = _1074_v85
			var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
			_ = _nw156
			_1074_v85 = _nw156
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_1074_v85), 0))
			_ = _index156
			(_1074_v85).ArraySet1((_961_v5).Intersection(_961_v5), (_index156).Int())
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_1074_v85), 0))
			_ = _index157
			(_1074_v85).ArraySet1(_961_v5, (_index157).Int())
		}
		var _1075_v86 _dafny.Array
		_ = _1075_v86
		var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw157
		_1075_v86 = _nw157
		r0 = (func() _dafny.Array {
			if _960_v4 {
				return _1075_v86
			}
			return _1075_v86
		})()
		r1 = _960_v4
		return r0, r1
	}
}
func (_this *C3) M7(p0 _dafny.Array, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (D1, _dafny.Int) {
	{
		var r0 D1 = Companion_D1_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1076_v0 _dafny.Array
		_ = _1076_v0
		var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw158
		_1076_v0 = _nw158
		var _1077_v1 _dafny.Array
		_ = _1077_v1
		var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw159
		_1077_v1 = _nw159
		var _1078_v2 _dafny.Map
		_ = _1078_v2
		_1078_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1076_v0, _1077_v1)
		_1078_v2 = (_1078_v2).Update(_1076_v0, _1077_v1)
		var _hi4 _dafny.Int = p3
		_ = _hi4
		for _1079_i0 := _dafny.IntOfInt64(896); _1079_i0.Cmp(_hi4) < 0; _1079_i0 = _1079_i0.Plus(_dafny.One) {
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((p0), 0))
			_ = _index158
			(p0).ArraySet1CodePoint(_dafny.CodePoint('k'), (_index158).Int())
			var _1080_v3 _dafny.CodePoint
			_ = _1080_v3
			_1080_v3 = _dafny.CodePoint('o')
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((p0), 0))
			_ = _index159
			(p0).ArraySet1CodePoint(_1080_v3, (_index159).Int())
			var _1081_v4 bool
			_ = _1081_v4
			_1081_v4 = false
			if _1081_v4 {
				(globalState).F2 = _1081_v4
				(globalState).F16 = _dafny.IntOfInt64(-201)
				var _1082_v5 _dafny.MultiSet
				_ = _1082_v5
				_1082_v5 = _dafny.MultiSetOf(_dafny.IntOfInt64(463), _dafny.IntOfInt64(279), p2)
				_1082_v5 = _1082_v5
				var _1083_v6 *C0
				_ = _1083_v6
				var _nw160 *C0 = New_C0_()
				_ = _nw160
				_nw160.Ctor__((func() bool {
					if !(_1081_v4) {
						return _1081_v4
					}
					return _1081_v4
				})())
				_1083_v6 = _nw160
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((p0), 0))
				_ = _index160
				(p0).ArraySet1CodePoint(_dafny.CodePoint('n'), (_index160).Int())
			} else {
				var _1084_v7 _dafny.Array
				_ = _1084_v7
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_28
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) bool = (func(_1085_p3 _dafny.Int, _1086_p1 _dafny.Int, _1087_v4 bool, _1088_p2 _dafny.Int) func(_dafny.Int) bool {
						return func(_1089_i1 _dafny.Int) bool {
							return !(false) || ((Companion_D4_.Create_DC10_(_1085_p3, _1086_p1, !(_1087_v4), _1087_v4, (_dafny.Zero).Minus(_1088_p2))).Dtor_cf15())
						}
					})(p3, p1, _1081_v4, p2)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw161 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw161).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw161).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1084_v7 = _nw161
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1084_v7), 0))
				_ = _index161
				(_1084_v7).ArraySet1((_1079_i0).Cmp(_1079_i0) < 0, (_index161).Int())
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1084_v7), 0))
				_ = _index162
				var _rhs144 bool = _1081_v4
				_ = _rhs144
				var _rhs145 bool = (Companion_D0_.Create_DC0_(_1081_v4)).Dtor_cf0()
				_ = _rhs145
				var _rhs146 bool = _1081_v4
				_ = _rhs146
				var _lhs136 *GlobalState = globalState
				_ = _lhs136
				var _lhs137 _dafny.Array = _1084_v7
				_ = _lhs137
				var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1084_v7), 0))
				_ = _lhs138
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				_lhs136.F2 = _rhs144
				(_lhs137).ArraySet1(_rhs145, (_lhs138).Int())
				_lhs139.F0 = _rhs146
				var _1090_v8 _dafny.MultiSet
				_ = _1090_v8
				_1090_v8 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1084_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1084_v7), 0))).Int()).(bool), p3)).Cardinality(), _dafny.IntOfInt64(879), (_dafny.SetOf(false, _1081_v4)).Cardinality(), p2, p3)
				var _1091_v9 _dafny.Sequence
				_ = _1091_v9
				_1091_v9 = _dafny.SeqOf(_1090_v8, _1090_v8)
				var _1092_v10 _dafny.Map
				_ = _1092_v10
				_1092_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1091_v9).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1091_v9).Cardinality()))).Uint32()).(_dafny.MultiSet), _1081_v4)
				var _1093_v11 _dafny.Map
				_ = _1093_v11
				_1093_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1081_v4, _1081_v4)
				var _1094_v12 _dafny.Map
				_ = _1094_v12
				_1094_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1093_v11).Contains(_1081_v4) {
						return (_1093_v11).Get(_1081_v4).(bool)
					}
					return (_1084_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1084_v7), 0))).Int()).(bool)
				})(), _1081_v4)
				var _1095_v13 *C0
				_ = _1095_v13
				var _nw162 *C0 = New_C0_()
				_ = _nw162
				_nw162.Ctor__(((func() bool {
					if (_1092_v10).Contains(_dafny.MultiSetOf(p2, p3)) {
						return (_1092_v10).Get(_dafny.MultiSetOf(p2, p3)).(bool)
					}
					return _1081_v4
				})()) || (Companion_Default___.Fm2(_1093_v11, globalState)))
				_1095_v13 = _nw162
				var _1096_v14 _dafny.Array
				_ = _1096_v14
				var _nw163 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(9))
				_ = _nw163
				_1096_v14 = _nw163
				var _1097_v15 _dafny.MultiSet
				_ = _1097_v15
				_1097_v15 = _dafny.MultiSetOf(_1095_v13)
				var _1098_v16 D4
				_ = _1098_v16
				_1098_v16 = Companion_D4_.Create_DC7_((_1097_v15).Cardinality())
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1096_v14), 0))
				_ = _index163
				(_1096_v14).ArraySet1(_1098_v16, (_index163).Int())
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1096_v14), 0))
				_ = _index164
				(_1096_v14).ArraySet1(Companion_Default___.Fm16(p1, _1081_v4, globalState), (_index164).Int())
				(globalState).F7 = p2
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1084_v7), 0))
				_ = _index165
				(_1084_v7).ArraySet1((_1084_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1084_v7), 0))).Int()).(bool), (_index165).Int())
			}
			(globalState).F0 = _1081_v4
			var _1099_v17 _dafny.Set
			_ = _1099_v17
			_1099_v17 = _dafny.SetOf(p1, p2)
			if (_1099_v17).IsDisjointFrom((_1099_v17).Union(_1099_v17)) {
				var _1100_v18 _dafny.Sequence
				_ = _1100_v18
				_1100_v18 = _dafny.SeqOf(_1079_i0)
				r1 = (p1).Times((_dafny.IntOfUint32((_1100_v18).Cardinality())).Plus(p2))
				var _1101_v19 _dafny.Sequence
				_ = _1101_v19
				_1101_v19 = _dafny.SeqOf(true)
				(globalState).F16 = _dafny.IntOfUint32((_1101_v19).Cardinality())
				var _1102_v20 _dafny.MultiSet
				_ = _1102_v20
				_1102_v20 = _dafny.MultiSetOf(Companion_Default___.Fm15(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg80 _dafny.Int) interface{} {
						return coer80(arg80)
					}
				}((func(_1103_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1104_i2 _dafny.Int) _dafny.Int {
						return _1103_p2
					}
				})(p2))), (_1100_v18).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1100_v18).Cardinality()))).Uint32()).(_dafny.Int), globalState))
				_1102_v20 = _1102_v20
				var _1105_v21 _dafny.Array
				_ = _1105_v21
				var _nw164 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(7))
				_ = _nw164
				_1105_v21 = _nw164
				var _1106_v22 _dafny.Set
				_ = _1106_v22
				_1106_v22 = _dafny.SetOf(_1081_v4)
				var _1107_v23 D6
				_ = _1107_v23
				_1107_v23 = Companion_D6_.Create_DC13_((_1106_v22).Cardinality(), (_dafny.SetOf(p1)).Cardinality())
				var _1108_v24 D6
				_ = _1108_v24
				_1108_v24 = Companion_D6_.Create_DC14_(_1107_v23)
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1105_v21), 0))
				_ = _index166
				(_1105_v21).ArraySet1(_1108_v24, (_index166).Int())
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1105_v21), 0))
				_ = _index167
				(_1105_v21).ArraySet1(_1108_v24, (_index167).Int())
				(globalState).F16 = p3
			} else {
				var _1109_v25 _dafny.Sequence
				_ = _1109_v25
				_1109_v25 = _dafny.SeqOf(_1081_v4, _1081_v4, _1081_v4)
				(globalState).F21 = (_1109_v25).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1109_v25).Cardinality()))).Uint32()).(bool)
				var _1110_v26 _dafny.MultiSet
				_ = _1110_v26
				_1110_v26 = _dafny.MultiSetOf(_1079_i0)
				var _1111_v27 _dafny.Map
				_ = _1111_v27
				_1111_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1110_v26, _1081_v4)
				var _1112_v28 _dafny.Map
				_ = _1112_v28
				_1112_v28 = _1111_v27
				(globalState).F2 = ((_1111_v27).Cardinality()).Cmp(Companion_Default___.SafeModuloInt(p2, p1)) == 0
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((p0), 0))
				_ = _index168
				(p0).ArraySet1CodePoint((p0).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((p0), 0))).Int()), (_index168).Int())
				_1080_v3 = _1080_v3
				(globalState).F0 = false
			}
		}
		var _1113_v29 bool
		_ = _1113_v29
		_1113_v29 = false
		var _1114_i3 _dafny.Int
		_ = _1114_i3
		_1114_i3 = _dafny.Zero
		{
			for !(_1113_v29) {
				{
					if (_1114_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1114_i3 = (_1114_i3).Plus(_dafny.One)
					(globalState).F13 = (func() bool {
						if _1113_v29 {
							return _1113_v29
						}
						return (func() bool {
							if _1113_v29 {
								return _1113_v29
							}
							return true
						})()
					})()
					var _1115_v30 _dafny.Array
					_ = _1115_v30
					var _nw165 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(14))
					_ = _nw165
					_1115_v30 = _nw165
					var _1116_v31 D4
					_ = _1116_v31
					_1116_v31 = Companion_D4_.Create_DC8_(_1113_v29)
					var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1115_v30), 0))
					_ = _index169
					(_1115_v30).ArraySet1(_1116_v31, (_index169).Int())
					var _1117_v32 _dafny.CodePoint
					_ = _1117_v32
					_1117_v32 = _dafny.CodePoint('d')
					var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1115_v30), 0))
					_ = _index170
					var _rhs147 D4 = _1116_v31
					_ = _rhs147
					var _rhs148 _dafny.CodePoint = _1117_v32
					_ = _rhs148
					var _lhs140 _dafny.Array = _1115_v30
					_ = _lhs140
					var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1115_v30), 0))
					_ = _lhs141
					(_lhs140).ArraySet1(_rhs147, (_lhs141).Int())
					_1117_v32 = _rhs148
					var _1118_v33 _dafny.Sequence
					_ = _1118_v33
					_1118_v33 = _dafny.SeqOf(_1113_v29)
					var _1119_v34 _dafny.Set
					_ = _1119_v34
					_1119_v34 = _dafny.SetOf(_1118_v33, _1118_v33)
					if (_1113_v29) || ((_1119_v34).IsSubsetOf(_1119_v34)) {
						var _1120_v35 *C0
						_ = _1120_v35
						var _nw166 *C0 = New_C0_()
						_ = _nw166
						_nw166.Ctor__((_dafny.IntOfInt64(-32)).Cmp((_dafny.Zero).Minus(p2)) >= 0)
						_1120_v35 = _nw166
						var _1121_v36 _dafny.Array
						_ = _1121_v36
						var _nw167 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
						_ = _nw167
						_1121_v36 = _nw167
						_1121_v36 = _1121_v36
						_1076_v0 = _1076_v0
						var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_1121_v36), 0))
						_ = _index171
						(_1121_v36).ArraySet1((_1120_v35.F27) || (_1120_v35.F27), (_index171).Int())
						var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_1121_v36), 0))
						_ = _index172
						(_1121_v36).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1118_v33, _1118_v33), (_index172).Int())
						_1117_v32 = _1117_v32
					} else {
						var _1122_v37 _dafny.Map
						_ = _1122_v37
						_1122_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v29, _1113_v29)
						var _1123_v38 _dafny.Map
						_ = _1123_v38
						_1123_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_1122_v37).Cardinality()).Minus(p2)), !(_1113_v29))
						var _1124_v39 _dafny.Map
						_ = _1124_v39
						_1124_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v29, p1)
						var _1125_v40 _dafny.Sequence
						_ = _1125_v40
						_1125_v40 = _dafny.UnicodeSeqOfUtf8Bytes("vn")
						var _1126_v41 _dafny.Set
						_ = _1126_v41
						_1126_v41 = _dafny.SetOf(_dafny.IntOfInt64(-813))
						var _1127_v43 _dafny.Sequence
						_ = _1127_v43
						_1127_v43 = _dafny.SeqOf(func() _dafny.Map {
							var _coll44 = _dafny.NewMapBuilder()
							_ = _coll44
							for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(316), _dafny.IntOfInt64(965))); ; {
								_compr_44, _ok53 := _iter53()
								if !_ok53 {
									break
								}
								var _1128_v42 _dafny.Int
								_1128_v42 = interface{}(_compr_44).(_dafny.Int)
								if ((_dafny.IntOfInt64(316)).Cmp(_1128_v42) <= 0) && ((_1128_v42).Cmp(_dafny.IntOfInt64(965)) < 0) {
									_coll44.Add(Companion_Default___.SafeModuloInt(_1128_v42, p1), true)
								}
							}
							return _coll44.ToMap()
						}())
						var _1129_v44 T1
						_ = _1129_v44
						var _nw168 *C1 = New_C1_()
						_ = _nw168
						_nw168.Ctor__(p2, _1127_v43)
						_1129_v44 = _nw168
						var _1130_v45 _dafny.Sequence
						_ = _1130_v45
						_1130_v45 = _dafny.SeqOf(_1129_v44)
						var _1131_v46 _dafny.MultiSet
						_ = _1131_v46
						_1131_v46 = _dafny.MultiSetOf(_1129_v44, (_1130_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-289), _dafny.IntOfUint32((_1130_v45).Cardinality()))).Uint32()).(T1))
						var _rhs149 bool = (func() bool {
							if (_1123_v38).Contains(p3) {
								return (_1123_v38).Get(p3).(bool)
							}
							return _1113_v29
						})()
						_ = _rhs149
						var _rhs150 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
							if (_1124_v39).Contains((_this).Fm6(p3, globalState)) {
								return (_1124_v39).Get((_this).Fm6(p3, globalState)).(_dafny.Int)
							}
							return Companion_Default___.Fm0(_dafny.IntOfInt64(187), _dafny.IntOfInt64(-322), p3, _1113_v29, globalState)
						})())
						_ = _rhs150
						var _rhs151 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1125_v40, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1125_v40).Cardinality()))).Uint32(), _1117_v32)
						_ = _rhs151
						var _rhs152 bool = (p2).Cmp(((_1126_v41).Cardinality()).Minus((_dafny.Zero).Minus(p2))) > 0
						_ = _rhs152
						var _rhs153 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
							if (_1131_v46).Contains(_1129_v44) {
								return (_1131_v46).Multiplicity(_1129_v44)
							}
							return p3
						})(), p1)
						_ = _rhs153
						var _lhs142 *GlobalState = globalState
						_ = _lhs142
						var _lhs143 *GlobalState = globalState
						_ = _lhs143
						var _lhs144 *GlobalState = globalState
						_ = _lhs144
						var _lhs145 *GlobalState = globalState
						_ = _lhs145
						_lhs142.F13 = _rhs149
						r1 = _rhs150
						_lhs143.F20 = _rhs151
						_lhs144.F0 = _rhs152
						_lhs145.F7 = _rhs153
						r1 = p2
						var _1132_v47 _dafny.Set
						_ = _1132_v47
						_1132_v47 = _dafny.SetOf(p0)
						var _1133_v48 _dafny.MultiSet
						_ = _1133_v48
						_1133_v48 = _dafny.MultiSetOf((_1132_v47).Cardinality())
						var _rhs154 _dafny.Int = (func() _dafny.Int {
							if _1113_v29 {
								return _dafny.IntOfInt64(260)
							}
							return Companion_Default___.SafeDivisionInt((func() _dafny.Int {
								if (_1133_v48).Contains(p2) {
									return (_1133_v48).Multiplicity(p2)
								}
								return p2
							})(), p1)
						})()
						_ = _rhs154
						var _rhs155 _dafny.Int = Companion_Default___.Fm0(((_dafny.Zero).Minus(_dafny.IntOfInt64(-766))).Plus(p1), Companion_Default___.Fm0(p1, p3, p1, false, globalState), ((_1124_v39).Merge(_1124_v39)).Cardinality(), !(_1113_v29) || (_1113_v29), globalState)
						_ = _rhs155
						var _lhs146 *GlobalState = globalState
						_ = _lhs146
						var _lhs147 *GlobalState = globalState
						_ = _lhs147
						_lhs146.F16 = _rhs154
						_lhs147.F16 = _rhs155
						var _1134_v49 *C1
						_ = _1134_v49
						var _nw169 *C1 = New_C1_()
						_ = _nw169
						_nw169.Ctor__(p1, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1123_v38, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1113_v29), _1123_v38, (_1123_v38).Update(p3, _1113_v29)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1125_v40).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_1123_v38, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1113_v29), _1123_v38, (_1123_v38).Update(p3, _1113_v29))).Cardinality()))).Uint32(), _1123_v38), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1123_v38, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1113_v29), _1123_v38, (_1123_v38).Update(p3, _1113_v29)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1125_v40).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_1123_v38, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1113_v29), _1123_v38, (_1123_v38).Update(p3, _1113_v29))).Cardinality()))).Uint32(), _1123_v38)).Cardinality()))).Uint32(), _1123_v38))
						_1134_v49 = _nw169
						var _1135_v50 _dafny.Array
						_ = _1135_v50
						var _len0_29 _dafny.Int = _dafny.IntOfInt64(5)
						_ = _len0_29
						var _nw170 _dafny.Array
						_ = _nw170
						if _len0_29.Cmp(_dafny.Zero) == 0 {
							_nw170 = _dafny.NewArray(_len0_29)
						} else {
							var _init29 func(_dafny.Int) bool = func(_1136_i4 _dafny.Int) bool {
								return true
							}
							_ = _init29
							var _element0_29 = _init29(_dafny.Zero)
							_ = _element0_29
							_nw170 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
							(_nw170).ArraySet1(_element0_29, 0)
							var _nativeLen0_29 = (_len0_29).Int()
							_ = _nativeLen0_29
							for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
								(_nw170).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
							}
						}
						_1135_v50 = _nw170
						var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1135_v50), 0))
						_ = _index173
						(_1135_v50).ArraySet1(_1113_v29, (_index173).Int())
						var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1135_v50), 0))
						_ = _index174
						(_1135_v50).ArraySet1(_1113_v29, (_index174).Int())
					}
					if _1113_v29 {
						var _1137_v51 _dafny.Sequence
						_ = _1137_v51
						_1137_v51 = _dafny.SeqOf(p1)
						var _1138_v52 D15
						_ = _1138_v52
						_1138_v52 = Companion_D15_.Create_DC33_(_dafny.Companion_Sequence_.Concatenate(_1137_v51, _1137_v51))
						_1138_v52 = _1138_v52
						(globalState).F2 = _1113_v29
						(globalState).F7 = p1
						var _1139_v53 _dafny.Array
						_ = _1139_v53
						var _len0_30 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_30
						var _nw171 _dafny.Array
						_ = _nw171
						if _len0_30.Cmp(_dafny.Zero) == 0 {
							_nw171 = _dafny.NewArray(_len0_30)
						} else {
							var _init30 func(_dafny.Int) _dafny.CodePoint = (func(_1140_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1141_i5 _dafny.Int) _dafny.CodePoint {
									return _1140_v32
								}
							})(_1117_v32)
							_ = _init30
							var _element0_30 = _init30(_dafny.Zero)
							_ = _element0_30
							_nw171 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
							(_nw171).ArraySet1CodePoint(_element0_30, 0)
							var _nativeLen0_30 = (_len0_30).Int()
							_ = _nativeLen0_30
							for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
								(_nw171).ArraySet1CodePoint(_init30(_dafny.IntOf(_i0_30)), _i0_30)
							}
						}
						_1139_v53 = _nw171
						_1139_v53 = p0
						(globalState).F16 = p2
					} else {
						var _1142_v54 _dafny.Map
						_ = _1142_v54
						_1142_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1113_v29)
						var _1143_v55 _dafny.Sequence
						_ = _1143_v55
						_1143_v55 = _dafny.SeqOf(_1142_v54)
						var _1144_v56 *C1
						_ = _1144_v56
						var _nw172 *C1 = New_C1_()
						_ = _nw172
						_nw172.Ctor__(Companion_Default___.SafeDivisionInt(p3, p3), _dafny.Companion_Sequence_.Concatenate(_1143_v55, _1143_v55))
						_1144_v56 = _nw172
						var _1145_v57 D12
						_ = _1145_v57
						_1145_v57 = Companion_D12_.Create_DC23_(p1)
						var _1146_v58 _dafny.Sequence
						_ = _1146_v58
						_1146_v58 = _dafny.SeqOf((_1145_v57).Dtor_cf30(), _1144_v56.F29)
						var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1076_v0), 0))
						_ = _index175
						(_1076_v0).ArraySet1(_dafny.IntOfUint32((_1146_v58).Cardinality()), (_index175).Int())
						var _1147_v59 _dafny.Sequence
						_ = _1147_v59
						_1147_v59 = _dafny.UnicodeSeqOfUtf8Bytes("ehgphob")
						var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1076_v0), 0))
						_ = _index176
						(_1076_v0).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1147_v59).Cardinality()), _dafny.IntOfInt64(-723)), (_index176).Int())
						var _1148_v60 *C0
						_ = _1148_v60
						var _nw173 *C0 = New_C0_()
						_ = _nw173
						_nw173.Ctor__(_1113_v29)
						_1148_v60 = _nw173
						var _1149_v62 _dafny.Array
						_ = _1149_v62
						var _len0_31 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_31
						var _nw174 _dafny.Array
						_ = _nw174
						if _len0_31.Cmp(_dafny.Zero) == 0 {
							_nw174 = _dafny.NewArray(_len0_31)
						} else {
							var _init31 func(_dafny.Int) _dafny.Map = (func(_1150_v56 *C1, _1151_p1 _dafny.Int, _1152_p3 _dafny.Int, _1153_v59 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
								return func(_1154_i6 _dafny.Int) _dafny.Map {
									return (func() _dafny.Map {
										var _coll45 = _dafny.NewMapBuilder()
										_ = _coll45
										for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-631), _dafny.IntOfInt64(508))); ; {
											_compr_45, _ok54 := _iter54()
											if !_ok54 {
												break
											}
											var _1155_v61 _dafny.Int
											_1155_v61 = interface{}(_compr_45).(_dafny.Int)
											if ((_dafny.IntOfInt64(-631)).Cmp(_1155_v61) <= 0) && ((_1155_v61).Cmp(_dafny.IntOfInt64(508)) < 0) {
												_coll45.Add((_1155_v61).Minus(_1151_p1), _dafny.MultiSetOf(_1152_p3))
											}
										}
										return _coll45.ToMap()
									}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v56.F29, _dafny.MultiSetOf(_dafny.IntOfInt64(887), _dafny.IntOfUint32((_1153_v59).Cardinality()))))
								}
							})(_1144_v56, p1, p3, _1147_v59)
							_ = _init31
							var _element0_31 = _init31(_dafny.Zero)
							_ = _element0_31
							_nw174 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
							(_nw174).ArraySet1(_element0_31, 0)
							var _nativeLen0_31 = (_len0_31).Int()
							_ = _nativeLen0_31
							for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
								(_nw174).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
							}
						}
						_1149_v62 = _nw174
						var _1156_v63 _dafny.MultiSet
						_ = _1156_v63
						_1156_v63 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1146_v58).Cardinality()))
						var _1157_v64 _dafny.Map
						_ = _1157_v64
						_1157_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer81 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg81 _dafny.Int) interface{} {
								return coer81(arg81)
							}
						}((func(_1158_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1159_i7 _dafny.Int) _dafny.Int {
								return _1158_p1
							}
						})(p1))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1156_v63)).Update(_dafny.IntOfInt64(884), _1156_v63))
						var _1160_v65 _dafny.Sequence
						_ = _1160_v65
						_1160_v65 = _1118_v33
						var _1161_v66 _dafny.Map
						_ = _1161_v66
						_1161_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.MultiSetOf(_dafny.IntOfUint32((_1160_v65).Cardinality())))
						var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1149_v62), 0))
						_ = _index177
						(_1149_v62).ArraySet1((func() _dafny.Map {
							if (_1157_v64).Contains(_1146_v58) {
								return (_1157_v64).Get(_1146_v58).(_dafny.Map)
							}
							return _1161_v66
						})(), (_index177).Int())
						var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1149_v62), 0))
						_ = _index178
						(_1149_v62).ArraySet1(func() _dafny.Map {
							var _coll46 = _dafny.NewMapBuilder()
							_ = _coll46
							for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(3), _dafny.IntOfInt64(524))); ; {
								_compr_46, _ok55 := _iter55()
								if !_ok55 {
									break
								}
								var _1162_v67 _dafny.Int
								_1162_v67 = interface{}(_compr_46).(_dafny.Int)
								if ((_dafny.IntOfInt64(3)).Cmp(_1162_v67) <= 0) && ((_1162_v67).Cmp(_dafny.IntOfInt64(524)) < 0) {
									_coll46.Add(Companion_Default___.SafeDivisionInt(_1162_v67, (_dafny.Zero).Minus(p2)), _1156_v63)
								}
							}
							return _coll46.ToMap()
						}(), (_index178).Int())
						(globalState).F21 = _1148_v60.F27
					}
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1163_v68 *C2
		_ = _1163_v68
		var _nw175 *C2 = New_C2_()
		_ = _nw175
		_nw175.Ctor__(p3)
		_1163_v68 = _nw175
		var _1164_v69 _dafny.Map
		_ = _1164_v69
		_1164_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1163_v68).F28(), _1113_v29)
		var _1165_v70 _dafny.Sequence
		_ = _1165_v70
		_1165_v70 = _dafny.UnicodeSeqOfUtf8Bytes("dlb")
		var _1166_v71 _dafny.Sequence
		_ = _1166_v71
		_1166_v71 = _dafny.SeqOf(_1164_v69)
		_1164_v69 = ((Companion_Default___.Fm3(p2, _1113_v29, _1165_v70, (_1163_v68).F28(), globalState)).Merge((_1166_v71).Select((Companion_Default___.SafeIndex((_1163_v68).F28(), _dafny.IntOfUint32((_1166_v71).Cardinality()))).Uint32()).(_dafny.Map))).Merge(_1164_v69)
		for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1076_v0), 0))); ; {
			_guard_loop_8, _ok56 := _iter56()
			if !_ok56 {
				break
			}
			var _1167_i8 _dafny.Int
			_1167_i8 = interface{}(_guard_loop_8).(_dafny.Int)
			if (true) && (((_1167_i8).Sign() != -1) && ((_1167_i8).Cmp(_dafny.ArrayLen((_1076_v0), 0)) < 0)) {
				(_1076_v0).ArraySet1((_1167_i8).Minus(p2), (_1167_i8).Int())
			}
		}
		r0 = Companion_D1_.Create_DC2_(_1164_v69)
		r1 = Companion_Default___.Fm0(p3, (_1163_v68).F28(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(372))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}((func(_1168_v29 bool) func(_dafny.Int) _dafny.Sequence {
			return func(_1169_i9 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_1168_v29)
			}
		})(_1113_v29)))).Cardinality()), _1113_v29, globalState)
		return r0, r1
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	F25  bool
	_f26 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this.F25 = false
	_this._f26 = false
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__(f25 bool, f26 bool) {
	{
		(_this).F25 = f25
		(_this)._f26 = f26
	}
}
func (_this *C4) Fm9(p0 bool, globalState *GlobalState) bool {
	{
		var _source18 D3 = Companion_D3_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("tfdl"), _dafny.CodePoint('r'), _this.F25)
		_ = _source18
		if _source18.Is_DC6() {
			var _1170___mcc_h0 _dafny.Sequence = _source18.Get_().(D3_DC6).Cf6
			_ = _1170___mcc_h0
			var _1171___mcc_h1 _dafny.CodePoint = _source18.Get_().(D3_DC6).Cf7
			_ = _1171___mcc_h1
			var _1172___mcc_h2 bool = _source18.Get_().(D3_DC6).Cf8
			_ = _1172___mcc_h2
			var _1173_cf8 bool = _1172___mcc_h2
			_ = _1173_cf8
			var _1174_cf7 _dafny.CodePoint = _1171___mcc_h1
			_ = _1174_cf7
			var _1175_cf6 _dafny.Sequence = _1170___mcc_h0
			_ = _1175_cf6
			return true
		} else {
			var _1176___mcc_h3 _dafny.CodePoint = _source18.Get_().(D3_DC5).Cf5
			_ = _1176___mcc_h3
			var _1177_cf5 _dafny.CodePoint = _1176___mcc_h3
			_ = _1177_cf5
			return (_this).F26()
		}
	}
}
func (_this *C4) Fm10(p0 _dafny.Int, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			var _coll47 = _dafny.NewMapBuilder()
			_ = _coll47
			for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-511), _dafny.IntOfInt64(-748))); ; {
				_compr_47, _ok57 := _iter57()
				if !_ok57 {
					break
				}
				var _1178_v0 _dafny.Int
				_1178_v0 = interface{}(_compr_47).(_dafny.Int)
				if ((_dafny.IntOfInt64(-511)).Cmp(_1178_v0) <= 0) && ((_1178_v0).Cmp(_dafny.IntOfInt64(-748)) < 0) {
					_coll47.Add((_1178_v0).Times((_dafny.SetOf(_dafny.IntOfInt64(-345))).Cardinality()), _dafny.CodePoint('g'))
				}
			}
			return _coll47.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(387), _dafny.CodePoint('s')))).Merge((func() _dafny.Map {
			var _coll48 = _dafny.NewMapBuilder()
			_ = _coll48
			for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-284), _dafny.IntOfInt64(557))); ; {
				_compr_48, _ok58 := _iter58()
				if !_ok58 {
					break
				}
				var _1179_v1 _dafny.Int
				_1179_v1 = interface{}(_compr_48).(_dafny.Int)
				if ((_dafny.IntOfInt64(-284)).Cmp(_1179_v1) <= 0) && ((_1179_v1).Cmp(_dafny.IntOfInt64(557)) < 0) {
					_coll48.Add((_1179_v1).Plus(_dafny.IntOfInt64(882)), _dafny.CodePoint('a'))
				}
			}
			return _coll48.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_this.F25)).Cardinality()), _dafny.CodePoint('p'))))
	}
}
func (_this *C4) M5(p0 bool, globalState *GlobalState) {
	{
		var _1180_v0 _dafny.Array
		_ = _1180_v0
		var _nwElement0_29 bool = p0
		_ = _nwElement0_29
		var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(5))
		_ = _nw176
		(_nw176).ArraySet1(_nwElement0_29, 0)
		(_nw176).ArraySet1(_this.F25, 1)
		(_nw176).ArraySet1(false, 2)
		(_nw176).ArraySet1((_this).F26(), 3)
		(_nw176).ArraySet1(_this.F25, 4)
		_1180_v0 = _nw176
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))
		_ = _index179
		(_1180_v0).ArraySet1(!(p0), (_index179).Int())
		var _1181_v1 _dafny.Sequence
		_ = _1181_v1
		_1181_v1 = _dafny.SeqOf(_1180_v0, _1180_v0)
		var _1182_v2 _dafny.Int
		_ = _1182_v2
		_1182_v2 = _dafny.IntOfInt64(683)
		var _1183_v3 _dafny.Map
		_ = _1183_v3
		_1183_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1181_v1).Select((Companion_Default___.SafeIndex(_1182_v2, _dafny.IntOfUint32((_1181_v1).Cardinality()))).Uint32()).(_dafny.Array)), !(true))
		var _1184_v4 _dafny.Map
		_ = _1184_v4
		_1184_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1182_v2, _1182_v2)
		var _1185_v5 _dafny.MultiSet
		_ = _1185_v5
		_1185_v5 = _dafny.MultiSetOf(_1182_v2, _dafny.IntOfInt64(-508), (_1184_v4).Cardinality())
		var _1186_v6 _dafny.Sequence
		_ = _1186_v6
		_1186_v6 = _dafny.SeqOf(p0)
		var _1187_v7 _dafny.Map
		_ = _1187_v7
		_1187_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(768))
		var _1188_v8 _dafny.MultiSet
		_ = _1188_v8
		_1188_v8 = _dafny.MultiSetOf(_this.F25, p0)
		var _1189_v9 _dafny.Sequence
		_ = _1189_v9
		_1189_v9 = _dafny.SeqOf(_1188_v8, _1188_v8)
		var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))
		_ = _index180
		var _rhs156 bool = (func() bool {
			if (_1183_v3).Contains(_1180_v0) {
				return (_1183_v3).Get(_1180_v0).(bool)
			}
			return !(false)
		})()
		_ = _rhs156
		var _rhs157 _dafny.Int = Companion_Default___.SafeDivisionInt(((func() _dafny.Int {
			if (_1185_v5).Contains(_dafny.IntOfUint32((_1186_v6).Cardinality())) {
				return (_1185_v5).Multiplicity(_dafny.IntOfUint32((_1186_v6).Cardinality()))
			}
			return _1182_v2
		})()).Plus((_1187_v7).Cardinality()), (func() _dafny.Int {
			if (_1185_v5).Contains(_1182_v2) {
				return (_1185_v5).Multiplicity(_1182_v2)
			}
			return _1182_v2
		})())
		_ = _rhs157
		var _rhs158 bool = (_1186_v6).Select((Companion_Default___.SafeIndex(_1182_v2, _dafny.IntOfUint32((_1186_v6).Cardinality()))).Uint32()).(bool)
		_ = _rhs158
		var _rhs159 bool = _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm11(_1182_v2, globalState), _1189_v9)
		_ = _rhs159
		var _lhs148 *C4 = _this
		_ = _lhs148
		var _lhs149 *GlobalState = globalState
		_ = _lhs149
		var _lhs150 *C4 = _this
		_ = _lhs150
		var _lhs151 _dafny.Array = _1180_v0
		_ = _lhs151
		var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))
		_ = _lhs152
		_lhs148.F25 = _rhs156
		_lhs149.F7 = _rhs157
		_lhs150.F25 = _rhs158
		(_lhs151).ArraySet1(_rhs159, (_lhs152).Int())
		for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1180_v0), 0))); ; {
			_guard_loop_9, _ok59 := _iter59()
			if !_ok59 {
				break
			}
			var _1190_i0 _dafny.Int
			_1190_i0 = interface{}(_guard_loop_9).(_dafny.Int)
			if (true) && (((_1190_i0).Sign() != -1) && ((_1190_i0).Cmp(_dafny.ArrayLen((_1180_v0), 0)) < 0)) {
				(_1180_v0).ArraySet1(p0, (_1190_i0).Int())
			}
		}
		var _hi5 _dafny.Int = _1182_v2
		_ = _hi5
		for _1191_i1 := _1182_v2; _1191_i1.Cmp(_hi5) < 0; _1191_i1 = _1191_i1.Plus(_dafny.One) {
			var _1192_v10 _dafny.Array
			_ = _1192_v10
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_32
			var _nw177 _dafny.Array
			_ = _nw177
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw177 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Int = (func(_1193_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1194_i2 _dafny.Int) _dafny.Int {
						return (_1194_i2).Plus(_1193_i1)
					}
				})(_1191_i1)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw177 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw177).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw177).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1192_v10 = _nw177
			var _1195_v11 _dafny.Sequence
			_ = _1195_v11
			_1195_v11 = _dafny.SeqOf(_1191_i1)
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1192_v10), 0))
			_ = _index181
			(_1192_v10).ArraySet1(((_1195_v11).Select((Companion_Default___.SafeIndex(_1182_v2, _dafny.IntOfUint32((_1195_v11).Cardinality()))).Uint32()).(_dafny.Int)).Plus((_1184_v4).Cardinality()), (_index181).Int())
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1192_v10), 0))
			_ = _index182
			(_1192_v10).ArraySet1((func() _dafny.Int {
				if (_1188_v8).Contains((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool)) {
					return (_1188_v8).Multiplicity((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool))
				}
				return _1191_i1
			})(), (_index182).Int())
			var _1196_v12 _dafny.Map
			_ = _1196_v12
			_1196_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1180_v0, Companion_Default___.SafeModuloInt(_1182_v2, _dafny.IntOfInt64(475)))
			var _1197_v13 _dafny.CodePoint
			_ = _1197_v13
			_1197_v13 = _dafny.CodePoint('v')
			var _1198_v14 _dafny.Map
			_ = _1198_v14
			_1198_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1191_i1).Times(_1182_v2), (_1196_v12).Update(_1180_v0, _1191_i1))
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))
			_ = _index183
			var _rhs160 bool = false
			_ = _rhs160
			var _rhs161 _dafny.Map = (func() _dafny.Map {
				if (_1198_v14).Contains((_1192_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1192_v10), 0))).Int()).(_dafny.Int)) {
					return (_1198_v14).Get((_1192_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1192_v10), 0))).Int()).(_dafny.Int)).(_dafny.Map)
				}
				return (func() _dafny.Map {
					if (_1198_v14).Contains(_1182_v2) {
						return (_1198_v14).Get(_1182_v2).(_dafny.Map)
					}
					return _1196_v12
				})()
			})()
			_ = _rhs161
			var _rhs162 _dafny.Int = ((func() _dafny.Int {
				if (_1185_v5).Contains(_dafny.IntOfInt64(11)) {
					return (_1185_v5).Multiplicity(_dafny.IntOfInt64(11))
				}
				return _dafny.IntOfInt64(-590)
			})()).Times(_1182_v2)
			_ = _rhs162
			var _rhs163 bool = (_1186_v6).Select((Companion_Default___.SafeIndex((_1192_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1192_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1186_v6).Cardinality()))).Uint32()).(bool)
			_ = _rhs163
			var _rhs164 _dafny.CodePoint = _1197_v13
			_ = _rhs164
			var _lhs153 _dafny.Array = _1180_v0
			_ = _lhs153
			var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))
			_ = _lhs154
			var _lhs155 *GlobalState = globalState
			_ = _lhs155
			var _lhs156 *GlobalState = globalState
			_ = _lhs156
			(_lhs153).ArraySet1(_rhs160, (_lhs154).Int())
			_1196_v12 = _rhs161
			_lhs155.F7 = _rhs162
			_lhs156.F21 = _rhs163
			_1197_v13 = _rhs164
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1192_v10), 0))
			_ = _index184
			(_1192_v10).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(28), _1182_v2)), (_index184).Int())
			_1188_v8 = ((_dafny.MultiSetOf((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool), (_this).Fm9(p0, globalState))).Update(_this.F25, Companion_Default___.Abs(_1191_i1))).Difference((_1188_v8).Union(_1188_v8))
		}
		var _hi6 _dafny.Int = (_dafny.Zero).Minus(_1182_v2)
		_ = _hi6
		for _1199_i3 := Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(529), _1182_v2); _1199_i3.Cmp(_hi6) < 0; _1199_i3 = _1199_i3.Plus(_dafny.One) {
			var _1200_v15 _dafny.Map
			_ = _1200_v15
			_1200_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1182_v2).Plus(_1199_i3), (_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool))
			_1200_v15 = (_1200_v15).Update(_1182_v2, p0)
			var _1201_v16 _dafny.MultiSet
			_ = _1201_v16
			_1201_v16 = _dafny.MultiSetOf(_1180_v0)
			var _1202_v17 *C2
			_ = _1202_v17
			var _nw178 *C2 = New_C2_()
			_ = _nw178
			_nw178.Ctor__((_1201_v16).Cardinality())
			_1202_v17 = _nw178
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))
			_ = _index185
			(_1180_v0).ArraySet1((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool), (_index185).Int())
			var _1203_v18 D13
			_ = _1203_v18
			_1203_v18 = Companion_D13_.Create_DC27_(!(_this.F25))
			var _source19 D13 = (func() D13 {
				if _this.F25 {
					return _1203_v18
				}
				return _1203_v18
			})()
			_ = _source19
			if _source19.Is_DC26() {
				var _1204___mcc_h0 _dafny.Sequence = _source19.Get_().(D13_DC26).Cf33
				_ = _1204___mcc_h0
				var _1205___mcc_h1 bool = _source19.Get_().(D13_DC26).Cf34
				_ = _1205___mcc_h1
				var _1206___mcc_h2 bool = _source19.Get_().(D13_DC26).Cf35
				_ = _1206___mcc_h2
				var _1207_cf35 bool = _1206___mcc_h2
				_ = _1207_cf35
				var _1208_cf34 bool = _1205___mcc_h1
				_ = _1208_cf34
				var _1209_cf33 _dafny.Sequence = _1204___mcc_h0
				_ = _1209_cf33
				var _1210_v19 _dafny.CodePoint
				_ = _1210_v19
				_1210_v19 = _dafny.CodePoint('f')
				(globalState).F2 = !((_1186_v6).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_1202_v17).F28(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1202_v17).F28(), _1210_v19)).Cardinality()), _dafny.IntOfUint32((_1186_v6).Cardinality()))).Uint32()).(bool))
				var _1211_v20 *C2
				_ = _1211_v20
				var _nw179 *C2 = New_C2_()
				_ = _nw179
				_nw179.Ctor__((_dafny.IntOfInt64(758)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tmdic")).Cardinality())))
				_1211_v20 = _nw179
				var _1212_v21 _dafny.Array
				_ = _1212_v21
				var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw180
				_1212_v21 = _nw180
				(_1211_v20).M3(_1212_v21, globalState)
				var _1213_v22 _dafny.Sequence
				_ = _1213_v22
				_1213_v22 = _dafny.SeqOf(_1182_v2)
				var _1214_v23 _dafny.Map
				_ = _1214_v23
				_1214_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("knj"), _dafny.Companion_Sequence_.Concatenate(_1213_v22, _1213_v22))
				_1214_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_cf33, _1213_v22)
			} else if _source19.Is_DC27() {
				var _1215___mcc_h3 bool = _source19.Get_().(D13_DC27).Cf36
				_ = _1215___mcc_h3
				var _1216_cf36 bool = _1215___mcc_h3
				_ = _1216_cf36
				(globalState).F13 = _1216_cf36
				(globalState).F7 = ((_1202_v17).F28()).Times(_1182_v2)
				var _1217_v24 _dafny.Sequence
				_ = _1217_v24
				_1217_v24 = _dafny.SeqOf(_1182_v2, (_1202_v17).F28())
				var _1218_v25 _dafny.Sequence
				_ = _1218_v25
				_1218_v25 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer83 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1219_i3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1220_i4 _dafny.Int) _dafny.Int {
						return _1219_i3
					}
				})(_1199_i3))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-680))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_1221_v17 *C2) func(_dafny.Int) _dafny.Int {
					return func(_1222_i5 _dafny.Int) _dafny.Int {
						return (_1221_v17).F28()
					}
				})(_1202_v17)))), (func() _dafny.Sequence {
					if _this.F25 {
						return _1217_v24
					}
					return _1217_v24
				})(), _1217_v24)
				var _1223_v26 _dafny.Map
				_ = _1223_v26
				_1223_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1200_v15).Cardinality(), _1186_v6)
				var _1224_v27 _dafny.CodePoint
				_ = _1224_v27
				_1224_v27 = _dafny.CodePoint('p')
				var _1225_v28 _dafny.Map
				_ = _1225_v28
				_1225_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool), _this.F25)
				var _1226_v29 _dafny.Array
				_ = _1226_v29
				var _nwElement0_30 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1186_v6, (Companion_Default___.SafeIndex(((_1223_v26).Update(Companion_Default___.Fm0(_1182_v2, _1182_v2, (_1202_v17).F28(), false, globalState), _1186_v6)).Cardinality(), _dafny.IntOfUint32((_1186_v6).Cardinality()))).Uint32(), p0)
				_ = _nwElement0_30
				var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(19))
				_ = _nw181
				(_nw181).ArraySet1(_nwElement0_30, 0)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm18(_1216_cf36, _1224_v27, _this.F25, globalState), Companion_Default___.Fm18((_this).F26(), _1224_v27, p0, globalState)), 1)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1186_v6, _dafny.SeqOf(true, Companion_Default___.Fm2(_1225_v28, globalState), true)), 2)
				(_nw181).ArraySet1(_1186_v6, 3)
				(_nw181).ArraySet1(_1186_v6, 4)
				(_nw181).ArraySet1(_1186_v6, 5)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1186_v6, _1186_v6), 6)
				(_nw181).ArraySet1(_1186_v6, 7)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool), !((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool))), Companion_Default___.Fm18(_1216_cf36, _1224_v27, p0, globalState)), 8)
				(_nw181).ArraySet1(_1186_v6, 9)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool)), _1186_v6), 10)
				(_nw181).ArraySet1(_1186_v6, 11)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1186_v6, _dafny.SeqOf((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool))), 12)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1186_v6, _1186_v6), 13)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1186_v6, _1186_v6), 14)
				(_nw181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1186_v6, _1186_v6), 15)
				(_nw181).ArraySet1(_1186_v6, 16)
				(_nw181).ArraySet1(_dafny.SeqOf(_this.F25), 17)
				(_nw181).ArraySet1(_dafny.SeqOf(p0), 18)
				_1226_v29 = _nw181
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1226_v29), 0))
				_ = _index186
				(_1226_v29).ArraySet1(_1186_v6, (_index186).Int())
				var _1227_v30 _dafny.Sequence
				_ = _1227_v30
				_1227_v30 = _dafny.UnicodeSeqOfUtf8Bytes("ytqvm")
				var _1228_v31 _dafny.Set
				_ = _1228_v31
				_1228_v31 = _dafny.SetOf(_1227_v30, _1227_v30)
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1226_v29), 0))
				_ = _index187
				var _rhs165 bool = false
				_ = _rhs165
				var _rhs166 bool = _this.F25
				_ = _rhs166
				var _rhs167 _dafny.Sequence = Companion_Default___.Fm22(Companion_Default___.Fm2(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1216_cf36, (_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool)), globalState), (_1228_v31).IsDisjointFrom(_1228_v31), p0, globalState)
				_ = _rhs167
				var _rhs168 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1186_v6, _1186_v6)
				_ = _rhs168
				var _lhs157 *C4 = _this
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				var _lhs159 _dafny.Array = _1226_v29
				_ = _lhs159
				var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1226_v29), 0))
				_ = _lhs160
				_lhs157.F25 = _rhs165
				_lhs158.F13 = _rhs166
				_1218_v25 = _rhs167
				(_lhs159).ArraySet1(_rhs168, (_lhs160).Int())
				var _1229_v32 _dafny.Array
				_ = _1229_v32
				var _nwElement0_31 _dafny.Int = _dafny.IntOfInt64(770)
				_ = _nwElement0_31
				var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(15))
				_ = _nw182
				(_nw182).ArraySet1(_nwElement0_31, 0)
				(_nw182).ArraySet1(_1182_v2, 1)
				(_nw182).ArraySet1((_1202_v17).F28(), 2)
				(_nw182).ArraySet1(_dafny.IntOfUint32((_1227_v30).Cardinality()), 3)
				(_nw182).ArraySet1((_1202_v17).F28(), 4)
				(_nw182).ArraySet1((_1202_v17).F28(), 5)
				(_nw182).ArraySet1((_1202_v17).F28(), 6)
				(_nw182).ArraySet1(_1199_i3, 7)
				(_nw182).ArraySet1(_1182_v2, 8)
				(_nw182).ArraySet1(_dafny.IntOfInt64(140), 9)
				(_nw182).ArraySet1(_1182_v2, 10)
				(_nw182).ArraySet1((_1202_v17).F28(), 11)
				(_nw182).ArraySet1(_1199_i3, 12)
				(_nw182).ArraySet1(_1199_i3, 13)
				(_nw182).ArraySet1((_1202_v17).F28(), 14)
				_1229_v32 = _nw182
				var _1230_v33 _dafny.Set
				_ = _1230_v33
				_1230_v33 = _dafny.SetOf(_1229_v32, _1229_v32, _1229_v32)
				(globalState).F13 = (func() bool {
					if !((func() bool {
						if (_1225_v28).Contains(true) {
							return (_1225_v28).Get(true).(bool)
						}
						return _this.F25
					})()) {
						return (_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool)
					}
					return (_1230_v33).Contains(_1229_v32)
				})()
			} else {
				var _1231___mcc_h4 _dafny.Map = _source19.Get_().(D13_DC25).Cf32
				_ = _1231___mcc_h4
				var _1232_cf32 _dafny.Map = _1231___mcc_h4
				_ = _1232_cf32
				var _1233_v34 _dafny.Array
				_ = _1233_v34
				var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw183
				_1233_v34 = _nw183
				var _1234_v35 _dafny.Array
				_ = _1234_v35
				var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw184
				_1234_v35 = _nw184
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1233_v34), 0))
				_ = _index188
				(_1233_v34).ArraySet1(_1234_v35, (_index188).Int())
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_1234_v35), 0))
				_ = _index189
				(_1234_v35).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1185_v5).Contains((_1202_v17).F28()) {
						return (_1185_v5).Multiplicity((_1202_v17).F28())
					}
					return _1199_i3
				})()), (_index189).Int())
				var _1235_v36 _dafny.Array
				_ = _1235_v36
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_33
				var _nw185 _dafny.Array
				_ = _nw185
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw185 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Sequence = func(_1236_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(887))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg85 _dafny.Int) interface{} {
								return coer85(arg85)
							}
						}(func(_1237_i7 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('q')
						})), _dafny.UnicodeSeqOfUtf8Bytes("yxeguqbr"))
					}
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw185 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw185).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw185).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1235_v36 = _nw185
				var _1238_v37 _dafny.Sequence
				_ = _1238_v37
				_1238_v37 = _dafny.UnicodeSeqOfUtf8Bytes("w")
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1235_v36), 0))
				_ = _index190
				(_1235_v36).ArraySet1(_1238_v37, (_index190).Int())
				var _1239_v38 _dafny.Map
				_ = _1239_v38
				_1239_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1234_v35, _dafny.CodePoint('d'))
				var _1240_v39 _dafny.Map
				_ = _1240_v39
				_1240_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, (_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool))
				var _1241_v40 _dafny.Set
				_ = _1241_v40
				_1241_v40 = _dafny.SetOf(_1234_v35)
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1233_v34), 0))
				_ = _index191
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_1234_v35), 0))
				_ = _index192
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1235_v36), 0))
				_ = _index193
				var _rhs169 _dafny.Array = _1234_v35
				_ = _rhs169
				var _rhs170 _dafny.Int = (_dafny.Zero).Minus(((_1239_v38).Merge(_1239_v38)).Cardinality())
				_ = _rhs170
				var _rhs171 bool = Companion_Default___.Fm2(_1240_v39, globalState)
				_ = _rhs171
				var _rhs172 _dafny.Int = (((_1241_v40).Union(_1241_v40)).Intersection(_1241_v40)).Cardinality()
				_ = _rhs172
				var _rhs173 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("iaka"), _1238_v37)
				_ = _rhs173
				var _lhs161 _dafny.Array = _1233_v34
				_ = _lhs161
				var _lhs162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1233_v34), 0))
				_ = _lhs162
				var _lhs163 *GlobalState = globalState
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				var _lhs165 _dafny.Array = _1234_v35
				_ = _lhs165
				var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_1234_v35), 0))
				_ = _lhs166
				var _lhs167 _dafny.Array = _1235_v36
				_ = _lhs167
				var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1235_v36), 0))
				_ = _lhs168
				(_lhs161).ArraySet1(_rhs169, (_lhs162).Int())
				_lhs163.F16 = _rhs170
				_lhs164.F2 = _rhs171
				(_lhs165).ArraySet1(_rhs172, (_lhs166).Int())
				(_lhs167).ArraySet1(_rhs173, (_lhs168).Int())
				var _1242_v41 _dafny.Set
				_ = _1242_v41
				_1242_v41 = _dafny.SetOf((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool), p0)
				var _1243_v42 _dafny.Map
				_ = _1243_v42
				_1243_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v41, !(_this.F25) || ((_1180_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1180_v0), 0))).Int()).(bool)))
				_1243_v42 = (_1243_v42).Update((_1242_v41).Intersection(_1242_v41), p0)
				(globalState).F7 = Companion_Default___.SafeDivisionInt((_1202_v17).F28(), (_1202_v17).F28())
				var _1244_v43 *C2
				_ = _1244_v43
				var _nw186 *C2 = New_C2_()
				_ = _nw186
				_nw186.Ctor__((_1234_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_1234_v35), 0))).Int()).(_dafny.Int))
				_1244_v43 = _nw186
			}
		}
		_1180_v0 = (func() _dafny.Array {
			if false {
				return _1180_v0
			}
			return _1180_v0
		})()
		var _1245_v44 _dafny.Map
		_ = _1245_v44
		_1245_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _rhs174 bool = (_1186_v6).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(160)).Times(_1182_v2), _dafny.IntOfUint32((_1186_v6).Cardinality()))).Uint32()).(bool)
		_ = _rhs174
		var _rhs175 _dafny.MultiSet = _dafny.MultiSetOf(p0, (_1182_v2).Cmp((_1245_v44).Cardinality()) < 0, p0)
		_ = _rhs175
		var _lhs169 *GlobalState = globalState
		_ = _lhs169
		_lhs169.F21 = _rhs174
		_1188_v8 = _rhs175
	}
}
func (_this *C4) M6(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _1246_v0 _dafny.Sequence
		_ = _1246_v0
		_1246_v0 = _dafny.SeqOf(Companion_Default___.Fm0(p3, _dafny.IntOfInt64(950), p3, _this.F25, globalState))
		var _1247_v1 _dafny.Array
		_ = _1247_v1
		var _len0_34 _dafny.Int = _dafny.One
		_ = _len0_34
		var _nw187 _dafny.Array
		_ = _nw187
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw187 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) bool = func(_1248_i0 _dafny.Int) bool {
				return (func() bool {
					if (_this).F26() {
						return _this.F25
					}
					return (_this).F26()
				})()
			}
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw187 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw187).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw187).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1247_v1 = _nw187
		var _1249_v2 _dafny.Map
		_ = _1249_v2
		_1249_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _this.F25)
		var _1250_v5 D1
		_ = _1250_v5
		_1250_v5 = Companion_D1_.Create_DC2_(func() _dafny.Map {
			var _coll49 = _dafny.NewMapBuilder()
			_ = _coll49
			for _iter60 := _dafny.Iterate((_1246_v0).Elements()); ; {
				_compr_49, _ok60 := _iter60()
				if !_ok60 {
					break
				}
				var _1251_v4 _dafny.Int
				_1251_v4 = interface{}(_compr_49).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_1246_v0, _1251_v4) {
					_coll49.Add((_1251_v4).Plus(p3), _this.F25)
				}
			}
			return _coll49.ToMap()
		}())
		var _1252_v7 _dafny.Sequence
		_ = _1252_v7
		_1252_v7 = _dafny.SeqOf((_this).F26(), (_this).F26())
		var _1253_v8 _dafny.Map
		_ = _1253_v8
		_1253_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1249_v2)
		var _1254_v10 _dafny.Map
		_ = _1254_v10
		_1254_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), true)
		var _1255_v13 _dafny.Array
		_ = _1255_v13
		var _nwElement0_32 _dafny.Map = _1249_v2
		_ = _nwElement0_32
		var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(23))
		_ = _nw188
		(_nw188).ArraySet1(_nwElement0_32, 0)
		(_nw188).ArraySet1(_1249_v2, 1)
		(_nw188).ArraySet1(_1249_v2, 2)
		(_nw188).ArraySet1(func() _dafny.Map {
			var _coll50 = _dafny.NewMapBuilder()
			_ = _coll50
			for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(306), _dafny.IntOfInt64(956))); ; {
				_compr_50, _ok61 := _iter61()
				if !_ok61 {
					break
				}
				var _1256_v3 _dafny.Int
				_1256_v3 = interface{}(_compr_50).(_dafny.Int)
				if ((_dafny.IntOfInt64(306)).Cmp(_1256_v3) <= 0) && ((_1256_v3).Cmp(_dafny.IntOfInt64(956)) < 0) {
					_coll50.Add(Companion_Default___.SafeModuloInt(_1256_v3, p3), _this.F25)
				}
			}
			return _coll50.ToMap()
		}(), 3)
		(_nw188).ArraySet1((_1250_v5).Dtor_cf2(), 4)
		(_nw188).ArraySet1((_1249_v2).Merge(func() _dafny.Map {
			var _coll51 = _dafny.NewMapBuilder()
			_ = _coll51
			for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(766), _dafny.IntOfInt64(466))); ; {
				_compr_51, _ok62 := _iter62()
				if !_ok62 {
					break
				}
				var _1257_v6 _dafny.Int
				_1257_v6 = interface{}(_compr_51).(_dafny.Int)
				if ((_dafny.IntOfInt64(766)).Cmp(_1257_v6) <= 0) && ((_1257_v6).Cmp(_dafny.IntOfInt64(466)) < 0) {
					_coll51.Add(Companion_Default___.SafeModuloInt(_1257_v6, p3), _this.F25)
				}
			}
			return _coll51.ToMap()
		}()), 5)
		(_nw188).ArraySet1((_1249_v2).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_1252_v7).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1252_v7).Cardinality()))).Uint32()).(bool))), 6)
		(_nw188).ArraySet1((_1249_v2).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F26())), 7)
		(_nw188).ArraySet1(_1249_v2, 8)
		(_nw188).ArraySet1(_1249_v2, 9)
		(_nw188).ArraySet1(_1249_v2, 10)
		(_nw188).ArraySet1(((func() _dafny.Map {
			if (_1253_v8).Contains(p3) {
				return (_1253_v8).Get(p3).(_dafny.Map)
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F26())).Update(p3, true)
		})()).Merge(func() _dafny.Map {
			var _coll52 = _dafny.NewMapBuilder()
			_ = _coll52
			for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(658), _dafny.IntOfInt64(762))); ; {
				_compr_52, _ok63 := _iter63()
				if !_ok63 {
					break
				}
				var _1258_v9 _dafny.Int
				_1258_v9 = interface{}(_compr_52).(_dafny.Int)
				if ((_dafny.IntOfInt64(658)).Cmp(_1258_v9) <= 0) && ((_1258_v9).Cmp(_dafny.IntOfInt64(762)) < 0) {
					_coll52.Add(Companion_Default___.SafeDivisionInt(_1258_v9, p3), (_this).F26())
				}
			}
			return _coll52.ToMap()
		}()), 11)
		(_nw188).ArraySet1(_1249_v2, 12)
		(_nw188).ArraySet1(_1249_v2, 13)
		(_nw188).ArraySet1(_1249_v2, 14)
		(_nw188).ArraySet1(_1249_v2, 15)
		(_nw188).ArraySet1(_1249_v2, 16)
		(_nw188).ArraySet1(_1249_v2, 17)
		(_nw188).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1254_v10).Cardinality(), true), 18)
		(_nw188).ArraySet1(_1249_v2, 19)
		(_nw188).ArraySet1(_1249_v2, 20)
		(_nw188).ArraySet1(func() _dafny.Map {
			var _coll53 = _dafny.NewMapBuilder()
			_ = _coll53
			for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(444), _dafny.IntOfInt64(576))); ; {
				_compr_53, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _1259_v11 _dafny.Int
				_1259_v11 = interface{}(_compr_53).(_dafny.Int)
				if ((_dafny.IntOfInt64(444)).Cmp(_1259_v11) <= 0) && ((_1259_v11).Cmp(_dafny.IntOfInt64(576)) < 0) {
					_coll53.Add(Companion_Default___.SafeModuloInt(_1259_v11, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, (_this).F26())).Cardinality()), _this.F25)
				}
			}
			return _coll53.ToMap()
		}(), 21)
		(_nw188).ArraySet1(func() _dafny.Map {
			var _coll54 = _dafny.NewMapBuilder()
			_ = _coll54
			for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(788), _dafny.IntOfInt64(326))); ; {
				_compr_54, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _1260_v12 _dafny.Int
				_1260_v12 = interface{}(_compr_54).(_dafny.Int)
				if ((_dafny.IntOfInt64(788)).Cmp(_1260_v12) <= 0) && ((_1260_v12).Cmp(_dafny.IntOfInt64(326)) < 0) {
					_coll54.Add(Companion_Default___.SafeModuloInt(_1260_v12, p3), _this.F25)
				}
			}
			return _coll54.ToMap()
		}(), 22)
		_1255_v13 = _nw188
		var _1261_v14 _dafny.MultiSet
		_ = _1261_v14
		_1261_v14 = _dafny.MultiSetOf(!(_this.F25) || (!((_this).F26())), (_this).Fm9(_this.F25, globalState), (_this).F26(), !(Companion_Default___.Fm2(_1254_v10, globalState)), ((func() bool {
			if (_1249_v2).Contains(p3) {
				return (_1249_v2).Get(p3).(bool)
			}
			return (_this).F26()
		})()) || (_this.F25))
		var _1262_v15 _dafny.Sequence
		_ = _1262_v15
		_1262_v15 = _dafny.SeqOf(_1247_v1)
		var _rhs176 _dafny.Sequence = _1246_v0
		_ = _rhs176
		var _rhs177 _dafny.Array = (_1262_v15).Select((Companion_Default___.SafeIndex((p3).Plus(p3), _dafny.IntOfUint32((_1262_v15).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs177
		var _rhs178 _dafny.Array = _1255_v13
		_ = _rhs178
		var _rhs179 _dafny.MultiSet = _1261_v14
		_ = _rhs179
		_1246_v0 = _rhs176
		_1247_v1 = _rhs177
		_1255_v13 = _rhs178
		_1261_v14 = _rhs179
		var _1263_v16 D12
		_ = _1263_v16
		_1263_v16 = Companion_D12_.Create_DC23_((_dafny.Zero).Minus(p3))
		var _1264_v17 _dafny.Map
		_ = _1264_v17
		_1264_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, _1263_v16)
		_1264_v17 = _1264_v17
		r1 = p3
		var _1265_v18 _dafny.Array
		_ = _1265_v18
		var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
		_ = _nw189
		_1265_v18 = _nw189
		_1265_v18 = p0
		var _1266_i1 _dafny.Int
		_ = _1266_i1
		_1266_i1 = _dafny.Zero
		{
			for (_this).Fm9((_this.F25) || (_this.F25), globalState) {
				{
					if (_1266_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1266_i1 = (_1266_i1).Plus(_dafny.One)
					var _1267_v19 *C3
					_ = _1267_v19
					var _nw190 *C3 = New_C3_()
					_ = _nw190
					_nw190.Ctor__()
					_1267_v19 = _nw190
					if !_dafny.Companion_Sequence_.Contains(p1, _dafny.CodePoint('u')) {
						_1247_v1 = _1247_v1
						var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_1247_v1), 0))
						_ = _index194
						(_1247_v1).ArraySet1(_this.F25, (_index194).Int())
						var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_1247_v1), 0))
						_ = _index195
						var _rhs180 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("iongp"), _dafny.Companion_Sequence_.Concatenate(p1, p1))
						_ = _rhs180
						var _rhs181 bool = (_this).F26()
						_ = _rhs181
						var _rhs182 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("osab"))).Cardinality())
						_ = _rhs182
						var _lhs170 *C4 = _this
						_ = _lhs170
						var _lhs171 _dafny.Array = _1247_v1
						_ = _lhs171
						var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_1247_v1), 0))
						_ = _lhs172
						var _lhs173 *GlobalState = globalState
						_ = _lhs173
						_lhs170.F25 = _rhs180
						(_lhs171).ArraySet1(_rhs181, (_lhs172).Int())
						_lhs173.F16 = _rhs182
						var _1268_v20 _dafny.CodePoint
						_ = _1268_v20
						_1268_v20 = _dafny.CodePoint('q')
						var _1269_v21 _dafny.Map
						_ = _1269_v21
						_1269_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1247_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_1247_v1), 0))).Int()).(bool)), Companion_Default___.Fm0(p3, _dafny.IntOfInt64(169), p3, false, globalState))
						var _1270_v22 _dafny.Map
						_ = _1270_v22
						_1270_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1268_v20, (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(((_1269_v21).Merge(_1269_v21)).Cardinality()))))
						_1270_v22 = (_1270_v22).Update(_1268_v20, (_dafny.IntOfInt64(-25)).Minus((func() _dafny.Int {
							if (_1269_v21).Contains((_1247_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_1247_v1), 0))).Int()).(bool)) {
								return (_1269_v21).Get((_1247_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_1247_v1), 0))).Int()).(bool)).(_dafny.Int)
							}
							return p3
						})()))
						var _1271_v23 _dafny.Array
						_ = _1271_v23
						var _len0_35 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_35
						var _nw191 _dafny.Array
						_ = _nw191
						if _len0_35.Cmp(_dafny.Zero) == 0 {
							_nw191 = _dafny.NewArray(_len0_35)
						} else {
							var _init35 func(_dafny.Int) _dafny.Map = (func(_1272_v21 _dafny.Map) func(_dafny.Int) _dafny.Map {
								return func(_1273_i2 _dafny.Int) _dafny.Map {
									return _1272_v21
								}
							})(_1269_v21)
							_ = _init35
							var _element0_35 = _init35(_dafny.Zero)
							_ = _element0_35
							_nw191 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
							(_nw191).ArraySet1(_element0_35, 0)
							var _nativeLen0_35 = (_len0_35).Int()
							_ = _nativeLen0_35
							for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
								(_nw191).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
							}
						}
						_1271_v23 = _nw191
						_1271_v23 = _1271_v23
						var _1274_v24 _dafny.Array
						_ = _1274_v24
						var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
						_ = _nw192
						_1274_v24 = _nw192
						_1274_v24 = _1274_v24
					} else {
						var _1275_v25 _dafny.Map
						_ = _1275_v25
						_1275_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (p3).Times((_dafny.Zero).Minus(p3)))
						var _1276_v27 _dafny.Map
						_ = _1276_v27
						_1276_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
							var _coll55 = _dafny.NewBuilder()
							_ = _coll55
							for _iter66 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-718), _dafny.IntOfInt64(-647))); ; {
								_compr_55, _ok66 := _iter66()
								if !_ok66 {
									break
								}
								var _1277_v26 _dafny.Int
								_1277_v26 = interface{}(_compr_55).(_dafny.Int)
								if ((_dafny.IntOfInt64(-718)).Cmp(_1277_v26) <= 0) && ((_1277_v26).Cmp(_dafny.IntOfInt64(-647)) < 0) {
									_coll55.Add(Companion_Default___.SafeDivisionInt(_1277_v26, p3))
								}
							}
							return _coll55.ToSet()
						}(), _this.F25)
						_1275_v25 = (_1275_v25).Update((_1276_v27).Cardinality(), p3)
						var _1278_v28 D14
						_ = _1278_v28
						_1278_v28 = Companion_D14_.Create_DC30_(_dafny.Companion_Sequence_.Concatenate(_1246_v0, _1246_v0), _dafny.Companion_Sequence_.Concatenate(_1252_v7, _1252_v7))
						_1278_v28 = _1278_v28
						var _1279_v29 _dafny.Array
						_ = _1279_v29
						var _nwElement0_33 _dafny.Array = p2
						_ = _nwElement0_33
						var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(7))
						_ = _nw193
						(_nw193).ArraySet1(_nwElement0_33, 0)
						(_nw193).ArraySet1(p2, 1)
						(_nw193).ArraySet1(p0, 2)
						(_nw193).ArraySet1(p2, 3)
						(_nw193).ArraySet1(p2, 4)
						(_nw193).ArraySet1(p0, 5)
						(_nw193).ArraySet1(p2, 6)
						_1279_v29 = _nw193
						var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1279_v29), 0))
						_ = _index196
						(_1279_v29).ArraySet1(p2, (_index196).Int())
						var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1279_v29), 0))
						_ = _index197
						(_1279_v29).ArraySet1(_1265_v18, (_index197).Int())
						var _1280_v30 _dafny.Array
						_ = _1280_v30
						var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
						_ = _nw194
						_1280_v30 = _nw194
						var _1281_v31 _dafny.CodePoint
						_ = _1281_v31
						_1281_v31 = _dafny.CodePoint('r')
						var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1280_v30), 0))
						_ = _index198
						(_1280_v30).ArraySet1CodePoint(_1281_v31, (_index198).Int())
						var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1280_v30), 0))
						_ = _index199
						(_1280_v30).ArraySet1CodePoint(_1281_v31, (_index199).Int())
						var _1282_v33 _dafny.Sequence
						_ = _1282_v33
						_1282_v33 = _dafny.SeqOf(func() _dafny.Map {
							var _coll56 = _dafny.NewMapBuilder()
							_ = _coll56
							for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-929), _dafny.IntOfInt64(930))); ; {
								_compr_56, _ok67 := _iter67()
								if !_ok67 {
									break
								}
								var _1283_v32 _dafny.Int
								_1283_v32 = interface{}(_compr_56).(_dafny.Int)
								if ((_dafny.IntOfInt64(-929)).Cmp(_1283_v32) <= 0) && ((_1283_v32).Cmp(_dafny.IntOfInt64(930)) < 0) {
									_coll56.Add((_1283_v32).Times((_dafny.SetOf(p3, p3)).Cardinality()), _this.F25)
								}
							}
							return _coll56.ToMap()
						}())
						var _1284_v34 T1
						_ = _1284_v34
						var _nw195 *C1 = New_C1_()
						_ = _nw195
						_nw195.Ctor__(_dafny.IntOfUint32((p1).Cardinality()), _1282_v33)
						_1284_v34 = _nw195
						var _1285_v35 _dafny.Map
						_ = _1285_v35
						_1285_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1284_v34, p3)
						_1285_v35 = (_1285_v35).Update(_1284_v34, p3)
					}
					(globalState).F13 = _this.F25
					var _1286_v36 _dafny.Array
					_ = _1286_v36
					var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
					_ = _nw196
					_1286_v36 = _nw196
					var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1286_v36), 0))
					_ = _index200
					(_1286_v36).ArraySet1(p1, (_index200).Int())
					var _1287_v37 _dafny.CodePoint
					_ = _1287_v37
					_1287_v37 = _dafny.CodePoint('d')
					var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1286_v36), 0))
					_ = _index201
					(_1286_v36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p1, _dafny.SeqOf(_1287_v37, _1287_v37)), p1), (_index201).Int())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1288_i3 _dafny.Int
		_ = _1288_i3
		_1288_i3 = _dafny.Zero
		{
			for (func() bool {
				if _this.F25 {
					return (_this).F26()
				}
				return true
			})() {
				{
					if (_1288_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1288_i3 = (_1288_i3).Plus(_dafny.One)
					(globalState).F21 = _this.F25
					var _1289_v39 _dafny.Sequence
					_ = _1289_v39
					_1289_v39 = _dafny.SeqOf(func() _dafny.Map {
						var _coll57 = _dafny.NewMapBuilder()
						_ = _coll57
						for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(888), _dafny.IntOfInt64(272))); ; {
							_compr_57, _ok68 := _iter68()
							if !_ok68 {
								break
							}
							var _1290_v38 _dafny.Int
							_1290_v38 = interface{}(_compr_57).(_dafny.Int)
							if ((_dafny.IntOfInt64(888)).Cmp(_1290_v38) <= 0) && ((_1290_v38).Cmp(_dafny.IntOfInt64(272)) < 0) {
								_coll57.Add(Companion_Default___.SafeDivisionInt(_1290_v38, p3), (_this).F26())
							}
						}
						return _coll57.ToMap()
					}())
					var _1291_v40 *C1
					_ = _1291_v40
					var _nw197 *C1 = New_C1_()
					_ = _nw197
					_nw197.Ctor__(p3, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(483))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg86 _dafny.Int) interface{} {
							return coer86(arg86)
						}
					}((func(_1292_v2 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_1293_i4 _dafny.Int) _dafny.Map {
							return _1292_v2
						}
					})(_1249_v2))), _1289_v39))
					_1291_v40 = _nw197
					r1 = (_dafny.IntOfInt64(-433)).Times(Companion_Default___.SafeDivisionInt(p3, p3))
					var _1294_v41 _dafny.Set
					_ = _1294_v41
					_1294_v41 = _dafny.SetOf((_this).F26(), _this.F25)
					var _1295_v42 D4
					_ = _1295_v42
					_1295_v42 = Companion_D4_.Create_DC7_((_1294_v41).Cardinality())
					(_1291_v40).F29 = (_1295_v42).Dtor_cf9()
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		r0 = !(_this.F25)
		r1 = (p3).Times((_dafny.IntOfUint32((_1252_v7).Cardinality())).Plus(p3))
		var _1296_v43 D18
		_ = _1296_v43
		_1296_v43 = Companion_D18_.Create_DC42_(_1254_v10)
		r2 = Companion_Default___.Fm2((_1296_v43).Dtor_cf58(), globalState)
		var _1297_v44 _dafny.Set
		_ = _1297_v44
		_1297_v44 = _dafny.SetOf(p2, p2)
		r3 = (_1297_v44).IsProperSubsetOf(_1297_v44)
		return r0, r1, r2, r3
	}
}
func (_this *C4) F26() bool {
	{
		return _this._f26
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f24 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f24 = false
	return &_this
}

type CompanionStruct_C5_ struct {
}

var Companion_C5_ = CompanionStruct_C5_{}

func (_this *C5) Equals(other *C5) bool {
	return _this == other
}

func (_this *C5) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C5)
	return ok && _this.Equals(other)
}

func (*C5) String() string {
	return "_module.C5"
}

func Type_C5_() _dafny.TypeDescriptor {
	return type_C5_{}
}

type type_C5_ struct {
}

func (_this type_C5_) Default() interface{} {
	return (*C5)(nil)
}

func (_this type_C5_) String() string {
	return "main.C5"
}
func (_this *C5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f24 bool) {
	{
		(_this)._f24 = f24
	}
}
func (_this *C5) Fm7(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C5) Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("wx"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("mjsnrtion")))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(378))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg87 _dafny.Int) interface{} {
				return coer87(arg87)
			}
		}(func(_1298_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("sjvmhb"))))
	}
}
func (_this *C5) Fm6(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		var _source20 D3 = Companion_D3_.Create_DC5_(_dafny.CodePoint('x'))
		_ = _source20
		if _source20.Is_DC6() {
			var _1299___mcc_h0 _dafny.Sequence = _source20.Get_().(D3_DC6).Cf6
			_ = _1299___mcc_h0
			var _1300___mcc_h1 _dafny.CodePoint = _source20.Get_().(D3_DC6).Cf7
			_ = _1300___mcc_h1
			var _1301___mcc_h2 bool = _source20.Get_().(D3_DC6).Cf8
			_ = _1301___mcc_h2
			var _1302_cf8 bool = _1301___mcc_h2
			_ = _1302_cf8
			var _1303_cf7 _dafny.CodePoint = _1300___mcc_h1
			_ = _1303_cf7
			var _1304_cf6 _dafny.Sequence = _1299___mcc_h0
			_ = _1304_cf6
			return (_this).F24()
		} else {
			var _1305___mcc_h3 _dafny.CodePoint = _source20.Get_().(D3_DC5).Cf5
			_ = _1305___mcc_h3
			var _1306_cf5 _dafny.CodePoint = _1305___mcc_h3
			_ = _1306_cf5
			return ((_this).F24()) || ((_this).F24())
		}
	}
}
func (_this *C5) Fm8(p0 bool, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.CodePoint('t'))).Cardinality()), (_this).F24())).Cardinality()).Minus(_dafny.IntOfInt64(-530))).Times(_dafny.IntOfInt64(86)))
	}
}
func (_this *C5) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _1307_v0 _dafny.Array
		_ = _1307_v0
		var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw198
		_1307_v0 = _nw198
		for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1307_v0), 0))); ; {
			_guard_loop_10, _ok69 := _iter69()
			if !_ok69 {
				break
			}
			var _1308_i0 _dafny.Int
			_1308_i0 = interface{}(_guard_loop_10).(_dafny.Int)
			if (true) && (((_1308_i0).Sign() != -1) && ((_1308_i0).Cmp(_dafny.ArrayLen((_1307_v0), 0)) < 0)) {
				(_1307_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ckiixp"), _dafny.UnicodeSeqOfUtf8Bytes("ogsr")), (_1308_i0).Int())
			}
		}
		var _1309_v1 D4
		_ = _1309_v1
		_1309_v1 = Companion_D4_.Create_DC7_(p3)
		(globalState).F13 = (_this).Fm7(p1, (p3).Minus(p3), (_dafny.IntOfInt64(215)).Cmp((_1309_v1).Dtor_cf9()) != 0, globalState)
		if (_dafny.IntOfInt64(852)).Cmp(p0) != 0 {
			var _1310_v2 _dafny.Map
			_ = _1310_v2
			_1310_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p1)
			(globalState).F0 = Companion_Default___.Fm2(_1310_v2, globalState)
			var _1311_v3 _dafny.Map
			_ = _1311_v3
			_1311_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _1312_v4 _dafny.Set
			_ = _1312_v4
			_1312_v4 = _dafny.SetOf((_this).Fm6((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v3, Companion_Default___.Fm0(p3, p0, p3, p1, globalState))).Cardinality(), globalState))
			r1 = (p3).Minus(((_dafny.SetOf((_this).F24(), p1)).Difference(_1312_v4)).Cardinality())
			(globalState).F7 = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(412)).Times(p3), p0)
			var _1313_v5 _dafny.Array
			_ = _1313_v5
			var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw199
			_1313_v5 = _nw199
			var _1314_v6 _dafny.Sequence
			_ = _1314_v6
			_1314_v6 = _dafny.UnicodeSeqOfUtf8Bytes("bqijp")
			var _1315_v7 _dafny.CodePoint
			_ = _1315_v7
			_1315_v7 = _dafny.CodePoint('s')
			var _1316_v8 _dafny.Sequence
			_ = _1316_v8
			_1316_v8 = _dafny.SeqOf(!(true))
			var _rhs183 _dafny.Int = Companion_Default___.SafeDivisionInt((_1311_v3).Cardinality(), p0)
			_ = _rhs183
			var _rhs184 bool = !((func() bool {
				if p1 {
					return p1
				}
				return p1
			})()) || ((_this).F24())
			_ = _rhs184
			var _rhs185 bool = (_this).Fm7(!(p1) || (p1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dcbbvmlo"), _dafny.Companion_Sequence_.Update(_1314_v6, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1314_v6).Cardinality()))).Uint32(), _1315_v7))).Cardinality()), true, globalState)
			_ = _rhs185
			var _rhs186 bool = (_1316_v8).Select((Companion_Default___.SafeIndex((_1309_v1).Dtor_cf9(), _dafny.IntOfUint32((_1316_v8).Cardinality()))).Uint32()).(bool)
			_ = _rhs186
			var _rhs187 _dafny.Array = _1313_v5
			_ = _rhs187
			var _lhs174 *GlobalState = globalState
			_ = _lhs174
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			var _lhs176 *GlobalState = globalState
			_ = _lhs176
			_lhs174.F16 = _rhs183
			_lhs175.F0 = _rhs184
			r0 = _rhs185
			_lhs176.F0 = _rhs186
			_1313_v5 = _rhs187
			var _1317_v9 *C3
			_ = _1317_v9
			var _nw200 *C3 = New_C3_()
			_ = _nw200
			_nw200.Ctor__()
			_1317_v9 = _nw200
		} else {
			var _1318_v10 _dafny.Array
			_ = _1318_v10
			var _nwElement0_34 _dafny.Array = p2
			_ = _nwElement0_34
			var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(8))
			_ = _nw201
			(_nw201).ArraySet1(_nwElement0_34, 0)
			(_nw201).ArraySet1(p2, 1)
			(_nw201).ArraySet1(p2, 2)
			(_nw201).ArraySet1(p2, 3)
			(_nw201).ArraySet1(p2, 4)
			(_nw201).ArraySet1(p2, 5)
			(_nw201).ArraySet1(p2, 6)
			(_nw201).ArraySet1(p2, 7)
			_1318_v10 = _nw201
			_1318_v10 = _1318_v10
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((p2), 0))
			_ = _index202
			(p2).ArraySet1(p3, (_index202).Int())
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((p2), 0))
			_ = _index203
			var _rhs188 _dafny.Int = (p3).Minus(p0)
			_ = _rhs188
			var _rhs189 _dafny.Int = p3
			_ = _rhs189
			var _rhs190 _dafny.Int = p3
			_ = _rhs190
			var _lhs177 *GlobalState = globalState
			_ = _lhs177
			var _lhs178 *GlobalState = globalState
			_ = _lhs178
			var _lhs179 _dafny.Array = p2
			_ = _lhs179
			var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((p2), 0))
			_ = _lhs180
			_lhs177.F16 = _rhs188
			_lhs178.F16 = _rhs189
			(_lhs179).ArraySet1(_rhs190, (_lhs180).Int())
			var _1319_v11 _dafny.CodePoint
			_ = _1319_v11
			_1319_v11 = _dafny.CodePoint('f')
			_1319_v11 = _dafny.CodePoint('v')
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((p2), 0))
			_ = _index204
			(p2).ArraySet1(p3, (_index204).Int())
			var _1320_v12 _dafny.Array
			_ = _1320_v12
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_36
			var _nw202 _dafny.Array
			_ = _nw202
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw202 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) D12 = (func(_1321_p1 bool, _1322_p3 _dafny.Int) func(_dafny.Int) D12 {
					return func(_1323_i1 _dafny.Int) D12 {
						return Companion_D12_.Create_DC22_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1321_p1, _1322_p3))
					}
				})(p1, p3)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw202 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw202).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw202).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1320_v12 = _nw202
			var _1324_v13 _dafny.Map
			_ = _1324_v13
			_1324_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), _1320_v12)
			_1324_v13 = (_1324_v13).Update(p0, _1320_v12)
		}
		if !(p1) {
			var _1325_v14 _dafny.Sequence
			_ = _1325_v14
			_1325_v14 = _dafny.UnicodeSeqOfUtf8Bytes("gus")
			(globalState).F19 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1325_v14, _1325_v14), _1325_v14)
			var _1326_v15 _dafny.Array
			_ = _1326_v15
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_37
			var _nw203 _dafny.Array
			_ = _nw203
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw203 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Sequence = (func(_1327_p0 _dafny.Int, _1328_p3 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1329_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_1327_p0, _1327_p0, _1327_p0, _dafny.IntOfInt64(-190), _1328_p3)
					}
				})(p0, p3)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw203 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw203).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw203).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1326_v15 = _nw203
			var _1330_v16 D15
			_ = _1330_v16
			_1330_v16 = Companion_D15_.Create_DC32_(_1326_v15)
			var _1331_v17 D15
			_ = _1331_v17
			_1331_v17 = Companion_D15_.Create_DC34_(_1330_v16)
			var _1332_v18 _dafny.Map
			_ = _1332_v18
			_1332_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1331_v17)
			_1332_v18 = (_1332_v18).Update(p1, _1331_v17)
			var _1333_v19 _dafny.CodePoint
			_ = _1333_v19
			_1333_v19 = _dafny.CodePoint('n')
			var _1334_v20 _dafny.Map
			_ = _1334_v20
			_1334_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1325_v14, _dafny.Companion_Sequence_.Update(_1325_v14, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1325_v14).Cardinality()))).Uint32(), _1333_v19))).Cardinality()), p3)
			var _1335_v21 _dafny.Sequence
			_ = _1335_v21
			_1335_v21 = _dafny.SeqOf(_1334_v20)
			_1334_v20 = ((_1334_v20).Merge((_1335_v21).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1335_v21).Cardinality()))).Uint32()).(_dafny.Map))).Merge(_1334_v20)
			r1 = p0
			(globalState).F13 = p1
		} else {
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((p2), 0))
			_ = _index205
			(p2).ArraySet1(p3, (_index205).Int())
			var _1336_v22 _dafny.Set
			_ = _1336_v22
			_1336_v22 = _dafny.SetOf(p2, p2, p2)
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((p2), 0))
			_ = _index206
			(p2).ArraySet1((_1336_v22).Cardinality(), (_index206).Int())
			var _1337_v23 _dafny.Sequence
			_ = _1337_v23
			_1337_v23 = _dafny.UnicodeSeqOfUtf8Bytes("coillmrw")
			var _1338_v24 _dafny.MultiSet
			_ = _1338_v24
			_1338_v24 = _dafny.MultiSetOf(_1337_v23)
			var _1339_v25 _dafny.Map
			_ = _1339_v25
			_1339_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('l'))
			var _1340_v26 _dafny.CodePoint
			_ = _1340_v26
			_1340_v26 = _dafny.CodePoint('k')
			var _1341_v27 _dafny.MultiSet
			_ = _1341_v27
			_1341_v27 = _dafny.MultiSetOf((func() _dafny.Int {
				if !((_this).F24()) {
					return (_1338_v24).Cardinality()
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(183))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}(func(_1342_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				})), (Companion_Default___.SafeIndex((_this).Fm8((_this).F24(), (_1339_v25).Update(p3, _1340_v26), _1337_v23, globalState), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(183))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg89 _dafny.Int) interface{} {
						return coer89(arg89)
					}
				}(func(_1343_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))).Cardinality()))).Uint32(), Companion_Default___.Fm27((_this).F24(), !(false), globalState))).Cardinality())
			})())
			var _1344_v28 _dafny.Map
			_ = _1344_v28
			_1344_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), p0)
			var _1345_v29 _dafny.Map
			_ = _1345_v29
			_1345_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F24())
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((p2), 0))
			_ = _index207
			(p2).ArraySet1((func() _dafny.Int {
				if (_1341_v27).Contains(p3) {
					return (_1341_v27).Multiplicity(p3)
				}
				return (func() _dafny.Int {
					if (_1344_v28).Contains(_dafny.IntOfInt64(881)) {
						return (_1344_v28).Get(_dafny.IntOfInt64(881)).(_dafny.Int)
					}
					return (_1345_v29).Cardinality()
				})()
			})(), (_index207).Int())
			var _1346_v30 _dafny.Array
			_ = _1346_v30
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_38
			var _nw204 _dafny.Array
			_ = _nw204
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw204 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) bool = (func(_1347_p0 _dafny.Int, _1348_p2 _dafny.Array) func(_dafny.Int) bool {
					return func(_1349_i4 _dafny.Int) bool {
						return ((_dafny.Zero).Minus(_1347_p0)).Cmp((_1348_p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1348_p2), 0))).Int()).(_dafny.Int)) <= 0
					}
				})(p0, p2)
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw204 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw204).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw204).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1346_v30 = _nw204
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1346_v30), 0))
			_ = _index208
			(_1346_v30).ArraySet1(p1, (_index208).Int())
			var _1350_v31 _dafny.Map
			_ = _1350_v31
			_1350_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(488))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg90 _dafny.Int) interface{} {
					return coer90(arg90)
				}
			}((func(_1351_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1352_i5 _dafny.Int) _dafny.CodePoint {
					return _1351_v26
				}
			})(_1340_v26))), p0)
			var _1353_v32 *C4
			_ = _1353_v32
			var _nw205 *C4 = New_C4_()
			_ = _nw205
			_nw205.Ctor__(false, p1)
			_1353_v32 = _nw205
			var _1354_v33 _dafny.Sequence
			_ = _1354_v33
			_1354_v33 = _dafny.SeqOf((_1353_v32).F26())
			var _1355_v34 _dafny.Sequence
			_ = _1355_v34
			_1355_v34 = _dafny.SeqOf(p1, (_1354_v33).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1354_v33).Cardinality()))).Uint32()).(bool), p1)
			var _1356_v35 _dafny.Sequence
			_ = _1356_v35
			_1356_v35 = _dafny.SeqOf(p1, _1353_v32.F25)
			var _1357_v36 _dafny.Map
			_ = _1357_v36
			_1357_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1356_v35, _1340_v26)
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1346_v30), 0))
			_ = _index209
			var _rhs191 bool = true
			_ = _rhs191
			var _rhs192 _dafny.Array = _1346_v30
			_ = _rhs192
			var _rhs193 _dafny.Int = (func() _dafny.Int {
				if p1 {
					return ((_1350_v31).Cardinality()).Minus((_dafny.Zero).Minus((_dafny.SetOf(_1353_v32, _1353_v32, _1353_v32, _1353_v32)).Cardinality()))
				}
				return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1354_v33).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))
			})()
			_ = _rhs193
			var _rhs194 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_1357_v36).Contains(Companion_Default___.Fm41(!(true), _dafny.SetOf(Companion_Default___.Fm2(_1345_v29, globalState), true, (_1353_v32).F26(), _1353_v32.F25), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), p3, globalState)) {
					return (_1357_v36).Get(Companion_Default___.Fm41(!(true), _dafny.SetOf(Companion_Default___.Fm2(_1345_v29, globalState), true, (_1353_v32).F26(), _1353_v32.F25), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), p3, globalState)).(_dafny.CodePoint)
				}
				return _1340_v26
			})()
			_ = _rhs194
			var _lhs181 _dafny.Array = _1346_v30
			_ = _lhs181
			var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1346_v30), 0))
			_ = _lhs182
			var _lhs183 *GlobalState = globalState
			_ = _lhs183
			(_lhs181).ArraySet1(_rhs191, (_lhs182).Int())
			_1346_v30 = _rhs192
			_lhs183.F7 = _rhs193
			_1340_v26 = _rhs194
			(globalState).F8 = _1337_v23
			var _1358_v37 _dafny.Array
			_ = _1358_v37
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_39
			var _nw206 _dafny.Array
			_ = _nw206
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw206 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Map = (func(_1359_v32 *C4) func(_dafny.Int) _dafny.Map {
					return func(_1360_i6 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1359_v32.F25)
					}
				})(_1353_v32)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw206 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw206).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw206).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1358_v37 = _nw206
			_1358_v37 = (func() _dafny.Array {
				if (_1353_v32).F26() {
					return _1358_v37
				}
				return _1358_v37
			})()
		}
		if !(!(p1)) || (!(p1)) {
			var _1361_v38 _dafny.CodePoint
			_ = _1361_v38
			_1361_v38 = _dafny.CodePoint('m')
			_1361_v38 = _1361_v38
			var _1362_v39 _dafny.Sequence
			_ = _1362_v39
			_1362_v39 = _dafny.UnicodeSeqOfUtf8Bytes("i")
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1307_v0), 0))
			_ = _index210
			(_1307_v0).ArraySet1(_1362_v39, (_index210).Int())
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1307_v0), 0))
			_ = _index211
			(_1307_v0).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("woo"), (_index211).Int())
			var _1363_v40 _dafny.Array
			_ = _1363_v40
			var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw207
			_1363_v40 = _nw207
			_1363_v40 = _1363_v40
			var _1364_v41 _dafny.Map
			_ = _1364_v41
			_1364_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1361_v38, p0)
			(globalState).F19 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1362_v39, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1364_v41).Contains(_1361_v38) {
					return (_1364_v41).Get(_1361_v38).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(p3)
			})(), _dafny.IntOfUint32((_1362_v39).Cardinality()))).Uint32(), _1361_v38), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1362_v39, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1364_v41).Contains(_1361_v38) {
					return (_1364_v41).Get(_1361_v38).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(p3)
			})(), _dafny.IntOfUint32((_1362_v39).Cardinality()))).Uint32(), _1361_v38)).Cardinality()))).Uint32(), _1361_v38)
			(globalState).F13 = (_this).F24()
		} else {
			var _1365_v42 _dafny.Array
			_ = _1365_v42
			var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw208
			_1365_v42 = _nw208
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1365_v42), 0))
			_ = _index212
			(_1365_v42).ArraySet1(p2, (_index212).Int())
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((p2), 0))
			_ = _index213
			(p2).ArraySet1(p0, (_index213).Int())
			var _1366_v43 _dafny.Sequence
			_ = _1366_v43
			_1366_v43 = _dafny.UnicodeSeqOfUtf8Bytes("krdj")
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1365_v42), 0))
			_ = _index214
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((p2), 0))
			_ = _index215
			var _rhs195 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs195
			var _rhs196 _dafny.Array = p2
			_ = _rhs196
			var _rhs197 _dafny.Int = (func() _dafny.Int {
				if p1 {
					return _dafny.IntOfUint32((_1366_v43).Cardinality())
				}
				return p3
			})()
			_ = _rhs197
			var _lhs184 _dafny.Array = _1365_v42
			_ = _lhs184
			var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1365_v42), 0))
			_ = _lhs185
			var _lhs186 _dafny.Array = p2
			_ = _lhs186
			var _lhs187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((p2), 0))
			_ = _lhs187
			r1 = _rhs195
			(_lhs184).ArraySet1(_rhs196, (_lhs185).Int())
			(_lhs186).ArraySet1(_rhs197, (_lhs187).Int())
			var _1367_v44 _dafny.Array
			_ = _1367_v44
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_40
			var _nw209 _dafny.Array
			_ = _nw209
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw209 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) bool = (func(_1368_p1 bool) func(_dafny.Int) bool {
					return func(_1369_i7 _dafny.Int) bool {
						return _1368_p1
					}
				})(p1)
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw209 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw209).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw209).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1367_v44 = _nw209
			var _1370_v45 _dafny.Array
			_ = _1370_v45
			var _nwElement0_35 _dafny.Array = _1367_v44
			_ = _nwElement0_35
			var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(11))
			_ = _nw210
			(_nw210).ArraySet1(_nwElement0_35, 0)
			(_nw210).ArraySet1(_1367_v44, 1)
			(_nw210).ArraySet1(_1367_v44, 2)
			(_nw210).ArraySet1(_1367_v44, 3)
			(_nw210).ArraySet1(_1367_v44, 4)
			(_nw210).ArraySet1(_1367_v44, 5)
			(_nw210).ArraySet1(_1367_v44, 6)
			(_nw210).ArraySet1(_1367_v44, 7)
			(_nw210).ArraySet1(_1367_v44, 8)
			(_nw210).ArraySet1(_1367_v44, 9)
			(_nw210).ArraySet1(_1367_v44, 10)
			_1370_v45 = _nw210
			var _1371_v46 _dafny.Array
			_ = _1371_v46
			_1371_v46 = _1367_v44
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1370_v45), 0))
			_ = _index216
			(_1370_v45).ArraySet1((_1371_v46), (_index216).Int())
			var _1372_v47 _dafny.Map
			_ = _1372_v47
			_1372_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F24())
			var _1373_v48 D18
			_ = _1373_v48
			_1373_v48 = Companion_D18_.Create_DC42_((_1372_v47).Update(p1, (_this).F24()))
			var _1374_v49 _dafny.Sequence
			_ = _1374_v49
			_1374_v49 = _dafny.SeqOf((_this).F24(), p1, true, p1)
			var _1375_v50 _dafny.Sequence
			_ = _1375_v50
			_1375_v50 = _dafny.SeqOf((_1372_v47).Merge((_1373_v48).Dtor_cf58()), _1372_v47, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F24())).Update((Companion_Default___.Fm25(true, p0, p1, _dafny.IntOfUint32((_1374_v49).Cardinality()), globalState)).Dtor_cf8(), (_this).F24()))
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1370_v45), 0))
			_ = _index217
			var _rhs198 _dafny.Array = (func() _dafny.Array {
				if p1 {
					return _1367_v44
				}
				return _1367_v44
			})()
			_ = _rhs198
			var _rhs199 bool = p1
			_ = _rhs199
			var _rhs200 bool = p1
			_ = _rhs200
			var _rhs201 bool = !(_dafny.Companion_Sequence_.Equal(_1366_v43, _1366_v43))
			_ = _rhs201
			var _rhs202 _dafny.Map = (_1375_v50).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1375_v50).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs202
			var _lhs188 _dafny.Array = _1370_v45
			_ = _lhs188
			var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1370_v45), 0))
			_ = _lhs189
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			var _lhs191 *GlobalState = globalState
			_ = _lhs191
			var _lhs192 *GlobalState = globalState
			_ = _lhs192
			(_lhs188).ArraySet1(_rhs198, (_lhs189).Int())
			_lhs190.F13 = _rhs199
			_lhs191.F13 = _rhs200
			_lhs192.F2 = _rhs201
			_1372_v47 = _rhs202
			(globalState).F0 = (_this).Fm6(p3, globalState)
			var _1376_v51 _dafny.Map
			_ = _1376_v51
			_1376_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
			var _1377_v52 D4
			_ = _1377_v52
			_1377_v52 = Companion_D4_.Create_DC10_((_1376_v51).Cardinality(), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), (_this).F24(), p1, (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
			var _1378_v53 _dafny.Sequence
			_ = _1378_v53
			_1378_v53 = _dafny.SeqOf(_1377_v52, Companion_Default___.Fm42(globalState))
			var _1379_v54 _dafny.MultiSet
			_ = _1379_v54
			_1379_v54 = _dafny.MultiSetOf(p1)
			var _1380_v55 D4
			_ = _1380_v55
			_1380_v55 = Companion_D4_.Create_DC10_(_dafny.IntOfUint32((_1378_v53).Cardinality()), p3, p1, (Companion_D16_.Create_DC36_(!(p1), _1379_v54, p1)).Dtor_cf49(), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
			var _1381_v56 _dafny.Map
			_ = _1381_v56
			_1381_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1380_v55)
			var _1382_v57 _dafny.Map
			_ = _1382_v57
			_1382_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1381_v56)
			_1382_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm43((_this).F24(), p3, _dafny.IntOfUint32((_1366_v43).Cardinality()), p1, globalState))
			var _1383_v58 _dafny.Map
			_ = _1383_v58
			_1383_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.ArrayCastTo((_1365_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1365_v42), 0))).Int())))
			_1383_v58 = (_1383_v58).Update((_this).F24(), _dafny.ArrayCastTo((_1365_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1365_v42), 0))).Int())))
		}
		(globalState).F16 = p3
		r0 = !(((func() bool {
			if (_this).F24() {
				return false
			}
			return p1
		})()) == ((_this).F24()))
		r1 = (_dafny.Zero).Minus(p0)
		var _len0_41 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_41
		var _nw211 _dafny.Array
		_ = _nw211
		if _len0_41.Cmp(_dafny.Zero) == 0 {
			_nw211 = _dafny.NewArray(_len0_41)
		} else {
			var _init41 func(_dafny.Int) _dafny.Map = func(_1384_i8 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), Companion_D3_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("ysrfvc"), _dafny.CodePoint('w'), false))
			}
			_ = _init41
			var _element0_41 = _init41(_dafny.Zero)
			_ = _element0_41
			_nw211 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
			(_nw211).ArraySet1(_element0_41, 0)
			var _nativeLen0_41 = (_len0_41).Int()
			_ = _nativeLen0_41
			for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
				(_nw211).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
			}
		}
		r2 = _nw211
		return r0, r1, r2
	}
}
func (_this *C5) M3(p0 _dafny.Array, globalState *GlobalState) {
	{
		var _1385_v0 _dafny.Int
		_ = _1385_v0
		_1385_v0 = _dafny.IntOfInt64(48)
		var _1386_v1 _dafny.Set
		_ = _1386_v1
		_1386_v1 = _dafny.SetOf(_1385_v0)
		var _1387_v2 *C4
		_ = _1387_v2
		var _nw212 *C4 = New_C4_()
		_ = _nw212
		_nw212.Ctor__(!(_1386_v1).Equals(_1386_v1), (_this).F24())
		_1387_v2 = _nw212
		var _1388_v3 _dafny.MultiSet
		_ = _1388_v3
		_1388_v3 = _dafny.MultiSetOf((_this).F24())
		var _1389_v4 D16
		_ = _1389_v4
		_1389_v4 = Companion_D16_.Create_DC36_((_1387_v2).F26(), _1388_v3, (_1387_v2).F26())
		var _source21 D16 = _1389_v4
		_ = _source21
		if _source21.Is_DC36() {
			var _1390___mcc_h0 bool = _source21.Get_().(D16_DC36).Cf47
			_ = _1390___mcc_h0
			var _1391___mcc_h1 _dafny.MultiSet = _source21.Get_().(D16_DC36).Cf48
			_ = _1391___mcc_h1
			var _1392___mcc_h2 bool = _source21.Get_().(D16_DC36).Cf49
			_ = _1392___mcc_h2
			var _1393_cf49 bool = _1392___mcc_h2
			_ = _1393_cf49
			var _1394_cf48 _dafny.MultiSet = _1391___mcc_h1
			_ = _1394_cf48
			var _1395_cf47 bool = _1390___mcc_h0
			_ = _1395_cf47
			(globalState).F16 = _1385_v0
			var _1396_v5 _dafny.Sequence
			_ = _1396_v5
			_1396_v5 = _dafny.SeqOf(_1393_cf49, _1395_cf47)
			_1394_cf48 = ((func() _dafny.MultiSet {
				if _1395_cf47 {
					return _1388_v3
				}
				return _1394_cf48
			})()).Difference(_dafny.MultiSetFromSeq(_1396_v5))
			var _1397_v6 *C0
			_ = _1397_v6
			var _nw213 *C0 = New_C0_()
			_ = _nw213
			_nw213.Ctor__((_1396_v5).Select((Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1396_v5).Cardinality()))).Uint32()).(bool))
			_1397_v6 = _nw213
			var _1398_v7 *C0
			_ = _1398_v7
			var _nw214 *C0 = New_C0_()
			_ = _nw214
			_nw214.Ctor__((_1385_v0).Cmp(_dafny.IntOfInt64(74)) <= 0)
			_1398_v7 = _nw214
		} else if _source21.Is_DC35() {
			var _1399___mcc_h3 _dafny.Map = _source21.Get_().(D16_DC35).Cf46
			_ = _1399___mcc_h3
			var _1400_cf46 _dafny.Map = _1399___mcc_h3
			_ = _1400_cf46
			var _1401_v8 _dafny.Array
			_ = _1401_v8
			var _nw215 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(13))
			_ = _nw215
			_1401_v8 = _nw215
			var _1402_v9 D0
			_ = _1402_v9
			_1402_v9 = Companion_D0_.Create_DC1_(Companion_Default___.Fm4(_dafny.IntOfInt64(140), _1385_v0, globalState))
			var _1403_v10 _dafny.Sequence
			_ = _1403_v10
			_1403_v10 = _dafny.SeqOf(_1402_v9, _1402_v9, Companion_D0_.Create_DC0_(_1387_v2.F25))
			var _1404_v11 D0
			_ = _1404_v11
			_1404_v11 = Companion_D0_.Create_DC1_((_1403_v10).Select((Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1403_v10).Cardinality()))).Uint32()).(D0))
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1401_v8), 0))
			_ = _index218
			(_1401_v8).ArraySet1(_1404_v11, (_index218).Int())
			var _pat_let_tv30 = _1402_v9
			_ = _pat_let_tv30
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1401_v8), 0))
			_ = _index219
			(_1401_v8).ArraySet1(func(_pat_let22_0 D0) D0 {
				return func(_1405_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let23_0 D0) D0 {
						return func(_1406_dt__update_hcf1_h0 D0) D0 {
							return Companion_D0_.Create_DC1_(_1406_dt__update_hcf1_h0)
						}(_pat_let23_0)
					}(_pat_let_tv30)
				}(_pat_let22_0)
			}(_1404_v11), (_index219).Int())
			var _1407_v12 _dafny.Array
			_ = _1407_v12
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_42
			var _nw216 _dafny.Array
			_ = _nw216
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw216 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) _dafny.Sequence = (func(_1408_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1409_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_1408_v0)
					}
				})(_1385_v0)
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw216 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw216).ArraySet1(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw216).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_1407_v12 = _nw216
			var _1410_v13 D15
			_ = _1410_v13
			_1410_v13 = Companion_D15_.Create_DC32_(_1407_v12)
			var _1411_v14 _dafny.MultiSet
			_ = _1411_v14
			_1411_v14 = _dafny.MultiSetOf(_1410_v13)
			if !(!((_1411_v14).Update(_1410_v13, Companion_Default___.Abs(_1385_v0))).Equals(_1411_v14)) {
				var _1412_v16 _dafny.Sequence
				_ = _1412_v16
				_1412_v16 = _dafny.SeqOf((_1387_v2).F26(), !((_1387_v2).F26()))
				var _1413_v17 _dafny.Sequence
				_ = _1413_v17
				_1413_v17 = _dafny.SeqOf(p0)
				var _1414_v18 _dafny.Array
				_ = _1414_v18
				var _nwElement0_36 bool = _1387_v2.F25
				_ = _nwElement0_36
				var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(20))
				_ = _nw217
				(_nw217).ArraySet1(_nwElement0_36, 0)
				(_nw217).ArraySet1(((_1387_v2).F26()) == ((_this).F24()), 1)
				(_nw217).ArraySet1(_1387_v2.F25, 2)
				(_nw217).ArraySet1((_1387_v2).F26(), 3)
				(_nw217).ArraySet1((_1387_v2).F26(), 4)
				(_nw217).ArraySet1((func() _dafny.Set {
					var _coll58 = _dafny.NewBuilder()
					_ = _coll58
					for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(100), _dafny.IntOfInt64(-738))); ; {
						_compr_58, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _1415_v15 _dafny.Int
						_1415_v15 = interface{}(_compr_58).(_dafny.Int)
						if ((_dafny.IntOfInt64(100)).Cmp(_1415_v15) <= 0) && ((_1415_v15).Cmp(_dafny.IntOfInt64(-738)) < 0) {
							_coll58.Add((_1415_v15).Minus(_1385_v0))
						}
					}
					return _coll58.ToSet()
				}()).IsSubsetOf(_1386_v1), 5)
				(_nw217).ArraySet1(_1387_v2.F25, 6)
				(_nw217).ArraySet1(_1387_v2.F25, 7)
				(_nw217).ArraySet1(!((func() bool {
					if (_this).Fm7((_1387_v2).F26(), _1385_v0, _1387_v2.F25, globalState) {
						return (_1387_v2).F26()
					}
					return (_1387_v2).F26()
				})()), 8)
				(_nw217).ArraySet1((_1387_v2).F26(), 9)
				(_nw217).ArraySet1((_1387_v2).F26(), 10)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1412_v16, _dafny.Companion_Sequence_.Update(_1412_v16, (Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1412_v16).Cardinality()))).Uint32(), false)), 11)
				(_nw217).ArraySet1((_1387_v2).F26(), 12)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1413_v17, _1413_v17), 13)
				(_nw217).ArraySet1((_1387_v2).F26(), 14)
				(_nw217).ArraySet1(_1387_v2.F25, 15)
				(_nw217).ArraySet1((_1387_v2).F26(), 16)
				(_nw217).ArraySet1(_1387_v2.F25, 17)
				(_nw217).ArraySet1((_1387_v2).F26(), 18)
				(_nw217).ArraySet1((_1387_v2).F26(), 19)
				_1414_v18 = _nw217
				_1414_v18 = _1414_v18
				var _1416_v19 _dafny.Sequence
				_ = _1416_v19
				_1416_v19 = _dafny.UnicodeSeqOfUtf8Bytes("lb")
				(globalState).F20 = _1416_v19
				(globalState).F7 = _1385_v0
				var _1417_v20 _dafny.MultiSet
				_ = _1417_v20
				_1417_v20 = _dafny.MultiSetOf(_1385_v0)
				var _1418_v21 _dafny.Sequence
				_ = _1418_v21
				_1418_v21 = _dafny.SeqOf((_1388_v3).Cardinality(), _1385_v0)
				(globalState).F16 = (Companion_Default___.Fm0((_1417_v20).Cardinality(), _1385_v0, _dafny.IntOfUint32((_1418_v21).Cardinality()), (_1387_v2).F26(), globalState)).Times(_1385_v0)
				var _1419_v22 *C4
				_ = _1419_v22
				var _nw218 *C4 = New_C4_()
				_ = _nw218
				_nw218.Ctor__((_1387_v2).F26(), !((_1387_v2).F26()))
				_1419_v22 = _nw218
			} else {
				(_1387_v2).M5(_1387_v2.F25, globalState)
				var _1420_v23 _dafny.Sequence
				_ = _1420_v23
				_1420_v23 = _dafny.SeqOf(_1387_v2.F25)
				var _1421_v24 _dafny.Set
				_ = _1421_v24
				_1421_v24 = _dafny.SetOf(_1387_v2.F25, (_this).F24(), _1387_v2.F25)
				var _1422_v25 _dafny.Array
				_ = _1422_v25
				var _nwElement0_37 bool = _1387_v2.F25
				_ = _nwElement0_37
				var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(14))
				_ = _nw219
				(_nw219).ArraySet1(_nwElement0_37, 0)
				(_nw219).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1420_v23, _1420_v23), 1)
				(_nw219).ArraySet1(false, 2)
				(_nw219).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1420_v23, (_1387_v2).F26()), 3)
				(_nw219).ArraySet1(!((_1387_v2).F26()), 4)
				(_nw219).ArraySet1((_1387_v2).F26(), 5)
				(_nw219).ArraySet1((_1387_v2).F26(), 6)
				(_nw219).ArraySet1((_1387_v2).F26(), 7)
				(_nw219).ArraySet1((func() bool {
					if _1387_v2.F25 {
						return (_1387_v2).F26()
					}
					return true
				})(), 8)
				(_nw219).ArraySet1(_1387_v2.F25, 9)
				(_nw219).ArraySet1((_1421_v24).Equals(_1421_v24), 10)
				(_nw219).ArraySet1((_1385_v0).Cmp(_1385_v0) != 0, 11)
				(_nw219).ArraySet1(false, 12)
				(_nw219).ArraySet1((_this).F24(), 13)
				_1422_v25 = _nw219
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1422_v25), 0))
				_ = _index220
				(_1422_v25).ArraySet1((_this).F24(), (_index220).Int())
				var _1423_v26 _dafny.Sequence
				_ = _1423_v26
				_1423_v26 = _dafny.SeqOf(_1385_v0)
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1422_v25), 0))
				_ = _index221
				(_1422_v25).ArraySet1((((_dafny.Zero).Minus(_1385_v0)).Cmp(_1385_v0) == 0) || (_dafny.Companion_Sequence_.IsPrefixOf(_1423_v26, _1423_v26)), (_index221).Int())
				var _1424_v27 _dafny.Map
				_ = _1424_v27
				_1424_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1385_v0, (_this).F24())
				var _1425_v28 _dafny.Map
				_ = _1425_v28
				_1425_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1387_v2).F26(), false)
				var _rhs203 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.SetOf(_1387_v2.F25)).Intersection(_1421_v24)).Cardinality(), (_1422_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1422_v25), 0))).Int()).(bool))
				_ = _rhs203
				var _rhs204 _dafny.Map = ((_1425_v28).Update(_1387_v2.F25, (_1387_v2).F26())).Merge((_1425_v28).Merge(_1425_v28))
				_ = _rhs204
				var _rhs205 _dafny.Int = (_1385_v0).Minus(_dafny.IntOfInt64(155))
				_ = _rhs205
				var _rhs206 _dafny.Int = _1385_v0
				_ = _rhs206
				var _rhs207 bool = (_1387_v2).F26()
				_ = _rhs207
				var _lhs193 *GlobalState = globalState
				_ = _lhs193
				var _lhs194 *GlobalState = globalState
				_ = _lhs194
				var _lhs195 *GlobalState = globalState
				_ = _lhs195
				_1424_v27 = _rhs203
				_1425_v28 = _rhs204
				_lhs193.F16 = _rhs205
				_lhs194.F7 = _rhs206
				_lhs195.F2 = _rhs207
				var _1426_v29 _dafny.Sequence
				_ = _1426_v29
				_1426_v29 = _dafny.UnicodeSeqOfUtf8Bytes("udixpjur")
				var _1427_v30 _dafny.CodePoint
				_ = _1427_v30
				_1427_v30 = _dafny.CodePoint('m')
				(globalState).F8 = _dafny.Companion_Sequence_.Update(_1426_v29, (Companion_Default___.SafeIndex((_1421_v24).Cardinality(), _dafny.IntOfUint32((_1426_v29).Cardinality()))).Uint32(), _1427_v30)
				var _1428_v31 _dafny.Map
				_ = _1428_v31
				_1428_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm32((_1422_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1422_v25), 0))).Int()).(bool), _1385_v0, _dafny.IntOfUint32((_1423_v26).Cardinality()), globalState), _1387_v2.F25)
				var _1429_v32 D14
				_ = _1429_v32
				_1429_v32 = Companion_D14_.Create_DC28_(_dafny.SetOf(_dafny.IntOfInt64(717), _1385_v0))
				_1428_v31 = (_1428_v31).Update((func() D14 {
					if (_1422_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_1422_v25), 0))).Int()).(bool) {
						return _1429_v32
					}
					return _1429_v32
				})(), (Companion_Default___.Fm44(globalState)).Dtor_cf36())
			}
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((p0), 0))
			_ = _index222
			(p0).ArraySet1(_1385_v0, (_index222).Int())
			var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((p0), 0))
			_ = _index223
			(p0).ArraySet1(_1385_v0, (_index223).Int())
			var _1430_v33 _dafny.Map
			_ = _1430_v33
			_1430_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1387_v2).F26(), _dafny.IntOfInt64(262))
			var _1431_v34 _dafny.Sequence
			_ = _1431_v34
			_1431_v34 = _dafny.SeqOf(_1430_v33)
			var _1432_v35 _dafny.Sequence
			_ = _1432_v35
			_1432_v35 = _dafny.SeqOf(true)
			var _1433_v36 _dafny.Set
			_ = _1433_v36
			_1433_v36 = _dafny.SetOf(_1387_v2.F25, (_1387_v2).F26())
			var _1434_v37 _dafny.Array
			_ = _1434_v37
			var _nwElement0_38 bool = (_1387_v2).F26()
			_ = _nwElement0_38
			var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(9))
			_ = _nw220
			(_nw220).ArraySet1(_nwElement0_38, 0)
			(_nw220).ArraySet1(_1387_v2.F25, 1)
			(_nw220).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_1430_v33), _1431_v34), 2)
			(_nw220).ArraySet1(_1387_v2.F25, 3)
			(_nw220).ArraySet1(_1387_v2.F25, 4)
			(_nw220).ArraySet1((_1387_v2.F25) || (_1387_v2.F25), 5)
			(_nw220).ArraySet1((_this).F24(), 6)
			(_nw220).ArraySet1(!((_1387_v2).F26()) || ((_1387_v2).F26()), 7)
			(_nw220).ArraySet1((_dafny.IntOfUint32((_1432_v35).Cardinality())).Cmp((_1433_v36).Cardinality()) != 0, 8)
			_1434_v37 = _nw220
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1434_v37), 0))
			_ = _index224
			(_1434_v37).ArraySet1(_1387_v2.F25, (_index224).Int())
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1434_v37), 0))
			_ = _index225
			(_1434_v37).ArraySet1((_1387_v2).F26(), (_index225).Int())
		} else {
			var _1435___mcc_h4 D16 = _source21.Get_().(D16_DC37).Cf50
			_ = _1435___mcc_h4
			var _1436_cf50 D16 = _1435___mcc_h4
			_ = _1436_cf50
			var _1437_v38 _dafny.Sequence
			_ = _1437_v38
			_1437_v38 = _dafny.UnicodeSeqOfUtf8Bytes("vuuysgyyb")
			var _1438_v39 _dafny.Sequence
			_ = _1438_v39
			_1438_v39 = _dafny.SeqOf(_1387_v2.F25)
			var _1439_v40 _dafny.MultiSet
			_ = _1439_v40
			_1439_v40 = _dafny.MultiSetOf((_dafny.IntOfUint32((_1437_v38).Cardinality())).Plus((func() _dafny.Int {
				if (_1388_v3).Contains(_1387_v2.F25) {
					return (_1388_v3).Multiplicity(_1387_v2.F25)
				}
				return _1385_v0
			})()), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_1438_v39)).Cardinality()), _dafny.IntOfUint32((_1437_v38).Cardinality())))
			_1439_v40 = _dafny.MultiSetOf(_1385_v0, (_1385_v0).Times(_dafny.IntOfInt64(306)), Companion_Default___.Fm0(_1385_v0, (_dafny.Zero).Minus(_1385_v0), _1385_v0, !(_1387_v2.F25), globalState), _dafny.IntOfInt64(241))
			(globalState).F13 = _1387_v2.F25
			(globalState).F13 = (_1387_v2).F26()
			var _nw221 *C4 = New_C4_()
			_ = _nw221
			_nw221.Ctor__((_1387_v2).F26(), _1387_v2.F25)
			_1387_v2 = _nw221
		}
		var _1440_v41 _dafny.Map
		_ = _1440_v41
		_1440_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1387_v2).F26(), (_this).F24())
		if Companion_Default___.Fm2(_1440_v41, globalState) {
			var _1441_v42 _dafny.Array
			_ = _1441_v42
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_43
			var _nw222 _dafny.Array
			_ = _nw222
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw222 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.CodePoint = func(_1442_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw222 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw222).ArraySet1CodePoint(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw222).ArraySet1CodePoint(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1441_v42 = _nw222
			_1441_v42 = _1441_v42
			var _1443_v43 _dafny.MultiSet
			_ = _1443_v43
			_1443_v43 = _dafny.MultiSetOf(_1385_v0, _1385_v0)
			if ((func() _dafny.Int {
				if (_1387_v2).F26() {
					return _1385_v0
				}
				return (func() _dafny.Int {
					if (_1443_v43).Contains(_1385_v0) {
						return (_1443_v43).Multiplicity(_1385_v0)
					}
					return _1385_v0
				})()
			})()).Cmp(_dafny.IntOfInt64(161)) <= 0 {
				var _1444_v44 _dafny.Sequence
				_ = _1444_v44
				_1444_v44 = _dafny.SeqOf(_1388_v3)
				(globalState).F16 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1444_v44).Cardinality()), _1385_v0)
				(_1387_v2).F25 = _1387_v2.F25
				var _1445_v45 _dafny.Map
				_ = _1445_v45
				_1445_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_1385_v0, _dafny.IntOfInt64(380)), _1385_v0)
				var _1446_v46 _dafny.MultiSet
				_ = _1446_v46
				_1446_v46 = _dafny.MultiSetOf(p0)
				_1445_v45 = (_1445_v45).Update(_1385_v0, (_1385_v0).Minus((_1446_v46).Cardinality()))
				var _1447_v47 _dafny.Array
				_ = _1447_v47
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_44
				var _nw223 _dafny.Array
				_ = _nw223
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw223 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) _dafny.Sequence = func(_1448_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("hkapgn")
					}
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw223 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw223).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw223).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1447_v47 = _nw223
				var _1449_v48 _dafny.CodePoint
				_ = _1449_v48
				_1449_v48 = _dafny.CodePoint('r')
				var _1450_v49 _dafny.Map
				_ = _1450_v49
				_1450_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v47, _1449_v48)
				_1450_v49 = (_1450_v49).Update(_1447_v47, _1449_v48)
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p0), 0))
				_ = _index226
				(p0).ArraySet1(_dafny.IntOfInt64(-806), (_index226).Int())
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p0), 0))
				_ = _index227
				(p0).ArraySet1(_1385_v0, (_index227).Int())
			} else {
				var _1451_v50 _dafny.Map
				_ = _1451_v50
				_1451_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1385_v0, (_1387_v2).F26())
				var _1452_v52 _dafny.Sequence
				_ = _1452_v52
				_1452_v52 = _dafny.SeqOf(_1451_v50, _1451_v50, func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter71 := _dafny.Iterate((_1386_v1).Elements()); ; {
						_compr_59, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _1453_v51 _dafny.Int
						_1453_v51 = interface{}(_compr_59).(_dafny.Int)
						if (_1386_v1).Contains(_1453_v51) {
							_coll59.Add((_1453_v51).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_1387_v2.F25)).Cardinality())), _1387_v2.F25)
						}
					}
					return _coll59.ToMap()
				}())
				var _1454_v53 *C1
				_ = _1454_v53
				var _nw224 *C1 = New_C1_()
				_ = _nw224
				_nw224.Ctor__((_1385_v0).Times(_1385_v0), (func() _dafny.Sequence {
					if _1387_v2.F25 {
						return _dafny.SeqOf(_1451_v50, _1451_v50, _1451_v50)
					}
					return _1452_v52
				})())
				_1454_v53 = _nw224
				var _1455_v54 _dafny.Sequence
				_ = _1455_v54
				_1455_v54 = _dafny.SeqOf(p0)
				_1455_v54 = _1455_v54
				_1388_v3 = _1388_v3
				(globalState).F0 = (_1387_v2).F26()
				var _1456_v55 *C2
				_ = _1456_v55
				var _nw225 *C2 = New_C2_()
				_ = _nw225
				_nw225.Ctor__(_1385_v0)
				_1456_v55 = _nw225
			}
			var _source22 D13 = Companion_Default___.Fm45(globalState)
			_ = _source22
			if _source22.Is_DC26() {
				var _1457___mcc_h5 _dafny.Sequence = _source22.Get_().(D13_DC26).Cf33
				_ = _1457___mcc_h5
				var _1458___mcc_h6 bool = _source22.Get_().(D13_DC26).Cf34
				_ = _1458___mcc_h6
				var _1459___mcc_h7 bool = _source22.Get_().(D13_DC26).Cf35
				_ = _1459___mcc_h7
				var _1460_cf35 bool = _1459___mcc_h7
				_ = _1460_cf35
				var _1461_cf34 bool = _1458___mcc_h6
				_ = _1461_cf34
				var _1462_cf33 _dafny.Sequence = _1457___mcc_h5
				_ = _1462_cf33
				(globalState).F16 = _1385_v0
				(globalState).F13 = (_1387_v2).F26()
				var _1463_v56 _dafny.CodePoint
				_ = _1463_v56
				_1463_v56 = _dafny.CodePoint('w')
				_1463_v56 = (_1462_cf33).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1385_v0), _dafny.IntOfUint32((_1462_cf33).Cardinality()))).Uint32()).(_dafny.CodePoint)
				var _1464_v57 _dafny.Sequence
				_ = _1464_v57
				_1464_v57 = _dafny.SeqOf((_dafny.IntOfInt64(305)).Times(_1385_v0))
				_1464_v57 = _dafny.Companion_Sequence_.Concatenate(_1464_v57, _dafny.Companion_Sequence_.Concatenate(_1464_v57, _1464_v57))
			} else if _source22.Is_DC27() {
				var _1465___mcc_h8 bool = _source22.Get_().(D13_DC27).Cf36
				_ = _1465___mcc_h8
				var _1466_cf36 bool = _1465___mcc_h8
				_ = _1466_cf36
				(globalState).F0 = !((_1388_v3).Contains((_1387_v2).F26()))
				var _1467_v58 _dafny.Array
				_ = _1467_v58
				var _nw226 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw226
				_1467_v58 = _nw226
				var _1468_v59 _dafny.MultiSet
				_ = _1468_v59
				_1468_v59 = _dafny.MultiSetOf(_1467_v58, _1467_v58, _1467_v58)
				var _1469_v60 _dafny.Sequence
				_ = _1469_v60
				_1469_v60 = _dafny.SeqOf(_dafny.MultiSetOf(_1467_v58))
				_1468_v59 = (_1469_v60).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if _1387_v2.F25 {
						return _1385_v0
					}
					return _1385_v0
				})(), _dafny.IntOfUint32((_1469_v60).Cardinality()))).Uint32()).(_dafny.MultiSet)
				var _1470_v61 _dafny.Map
				_ = _1470_v61
				_1470_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1385_v0, (_1387_v2).F26())
				var _1471_v62 _dafny.Sequence
				_ = _1471_v62
				_1471_v62 = _dafny.SeqOf(Companion_Default___.Fm2(_1440_v41, globalState), (func() bool {
					if (_1470_v61).Contains(_1385_v0) {
						return (_1470_v61).Get(_1385_v0).(bool)
					}
					return (_this).F24()
				})(), (_1385_v0).Cmp(_1385_v0) < 0)
				var _rhs208 bool = (_1471_v62).Select((Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1471_v62).Cardinality()))).Uint32()).(bool)
				_ = _rhs208
				var _rhs209 _dafny.Int = (_dafny.Zero).Minus(_1385_v0)
				_ = _rhs209
				var _rhs210 _dafny.Array = _1467_v58
				_ = _rhs210
				var _rhs211 *C4 = _1387_v2
				_ = _rhs211
				var _rhs212 _dafny.Int = _1385_v0
				_ = _rhs212
				var _lhs196 *C4 = _1387_v2
				_ = _lhs196
				var _lhs197 *GlobalState = globalState
				_ = _lhs197
				var _lhs198 *GlobalState = globalState
				_ = _lhs198
				_lhs196.F25 = _rhs208
				_lhs197.F7 = _rhs209
				_1467_v58 = _rhs210
				_1387_v2 = _rhs211
				_lhs198.F7 = _rhs212
				var _1472_v63 _dafny.CodePoint
				_ = _1472_v63
				_1472_v63 = _dafny.CodePoint('c')
				_1472_v63 = _1472_v63
			} else {
				var _1473___mcc_h9 _dafny.Map = _source22.Get_().(D13_DC25).Cf32
				_ = _1473___mcc_h9
				var _1474_cf32 _dafny.Map = _1473___mcc_h9
				_ = _1474_cf32
				var _1475_v64 _dafny.Array
				_ = _1475_v64
				var _nw227 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw227
				_1475_v64 = _nw227
				_1475_v64 = _1475_v64
				var _1476_v65 _dafny.Map
				_ = _1476_v65
				_1476_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1387_v2.F25, _1474_cf32)
				_1476_v65 = (_1476_v65).Update((_1387_v2).F26(), _1474_cf32)
				var _1477_v66 _dafny.Array
				_ = _1477_v66
				var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
				_ = _nw228
				_1477_v66 = _nw228
				var _rhs213 bool = _1387_v2.F25
				_ = _rhs213
				var _rhs214 _dafny.Array = _1477_v66
				_ = _rhs214
				var _rhs215 bool = (_1387_v2).F26()
				_ = _rhs215
				var _rhs216 bool = _1387_v2.F25
				_ = _rhs216
				var _lhs199 *GlobalState = globalState
				_ = _lhs199
				var _lhs200 *C4 = _1387_v2
				_ = _lhs200
				var _lhs201 *GlobalState = globalState
				_ = _lhs201
				_lhs199.F2 = _rhs213
				_1477_v66 = _rhs214
				_lhs200.F25 = _rhs215
				_lhs201.F0 = _rhs216
				var _1478_v67 *C2
				_ = _1478_v67
				var _nw229 *C2 = New_C2_()
				_ = _nw229
				_nw229.Ctor__((_dafny.Zero).Minus(_1385_v0))
				_1478_v67 = _nw229
			}
			var _1479_v68 _dafny.CodePoint
			_ = _1479_v68
			_1479_v68 = _dafny.CodePoint('q')
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_1441_v42), 0))
			_ = _index228
			(_1441_v42).ArraySet1CodePoint(_1479_v68, (_index228).Int())
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_1441_v42), 0))
			_ = _index229
			(_1441_v42).ArraySet1CodePoint(_1479_v68, (_index229).Int())
			var _1480_v69 _dafny.Sequence
			_ = _1480_v69
			_1480_v69 = _dafny.SeqOf((_1387_v2).F26())
			var _1481_v70 D14
			_ = _1481_v70
			_1481_v70 = Companion_D14_.Create_DC30_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg91 _dafny.Int) interface{} {
					return coer91(arg91)
				}
			}(func(_1482_i3 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(832)
			})), _1480_v69)
			var _1483_v71 _dafny.Array
			_ = _1483_v71
			var _nwElement0_39 _dafny.Sequence = _1480_v69
			_ = _nwElement0_39
			var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(20))
			_ = _nw230
			(_nw230).ArraySet1(_nwElement0_39, 0)
			(_nw230).ArraySet1(_1480_v69, 1)
			(_nw230).ArraySet1(Companion_Default___.Fm18(_1387_v2.F25, _dafny.CodePoint('p'), (_1387_v2).F26(), globalState), 2)
			(_nw230).ArraySet1(_dafny.SeqOf(_1387_v2.F25), 3)
			(_nw230).ArraySet1(_dafny.SeqOf(true, !(false), (_this).F24(), (_1387_v2).F26(), true), 4)
			(_nw230).ArraySet1(_1480_v69, 5)
			(_nw230).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1387_v2.F25), _1480_v69), 6)
			(_nw230).ArraySet1(_dafny.SeqOf((_this).F24()), 7)
			(_nw230).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1480_v69, _1480_v69), 8)
			(_nw230).ArraySet1(_1480_v69, 9)
			(_nw230).ArraySet1(_1480_v69, 10)
			(_nw230).ArraySet1(_1480_v69, 11)
			(_nw230).ArraySet1(_1480_v69, 12)
			(_nw230).ArraySet1(_1480_v69, 13)
			(_nw230).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1480_v69, _dafny.SeqOf(_1387_v2.F25, _1387_v2.F25)), 14)
			(_nw230).ArraySet1((func() _dafny.Sequence {
				if true {
					return Companion_Default___.Fm18(Companion_Default___.Fm2(_1440_v41, globalState), _1479_v68, _1387_v2.F25, globalState)
				}
				return (_1481_v70).Dtor_cf41()
			})(), 15)
			(_nw230).ArraySet1(_dafny.SeqOf(_1387_v2.F25), 16)
			(_nw230).ArraySet1(_1480_v69, 17)
			(_nw230).ArraySet1(_1480_v69, 18)
			(_nw230).ArraySet1(_1480_v69, 19)
			_1483_v71 = _nw230
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1483_v71), 0))
			_ = _index230
			(_1483_v71).ArraySet1(_1480_v69, (_index230).Int())
			var _1484_v72 _dafny.Set
			_ = _1484_v72
			_1484_v72 = _dafny.SetOf((_1387_v2).F26())
			var _1485_v73 _dafny.Set
			_ = _1485_v73
			_1485_v73 = (_1484_v72).Difference(_dafny.SetOf((_1387_v2).F26(), _1387_v2.F25))
			var _1486_v74 _dafny.Sequence
			_ = _1486_v74
			_1486_v74 = _dafny.UnicodeSeqOfUtf8Bytes("sgpusvt")
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1483_v71), 0))
			_ = _index231
			var _rhs217 _dafny.Sequence = _1486_v74
			_ = _rhs217
			var _rhs218 _dafny.Int = _dafny.IntOfInt64(-46)
			_ = _rhs218
			var _rhs219 _dafny.Sequence = _1480_v69
			_ = _rhs219
			var _rhs220 _dafny.Set = _1485_v73
			_ = _rhs220
			var _rhs221 bool = _1387_v2.F25
			_ = _rhs221
			var _lhs202 *GlobalState = globalState
			_ = _lhs202
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 _dafny.Array = _1483_v71
			_ = _lhs204
			var _lhs205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1483_v71), 0))
			_ = _lhs205
			var _lhs206 *GlobalState = globalState
			_ = _lhs206
			_lhs202.F20 = _rhs217
			_lhs203.F7 = _rhs218
			(_lhs204).ArraySet1(_rhs219, (_lhs205).Int())
			_1485_v73 = _rhs220
			_lhs206.F21 = _rhs221
		} else {
			var _1487_v75 _dafny.Sequence
			_ = _1487_v75
			_1487_v75 = _dafny.UnicodeSeqOfUtf8Bytes("bbe")
			var _1488_v76 _dafny.Map
			_ = _1488_v76
			_1488_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_1487_v75, _1487_v75), _1385_v0)
			var _rhs222 _dafny.Int = Companion_Default___.SafeModuloInt(_1385_v0, _1385_v0)
			_ = _rhs222
			var _rhs223 bool = (_1387_v2).F26()
			_ = _rhs223
			var _rhs224 _dafny.Map = (_1488_v76).Merge(_1488_v76)
			_ = _rhs224
			var _lhs207 *GlobalState = globalState
			_ = _lhs207
			var _lhs208 *C4 = _1387_v2
			_ = _lhs208
			_lhs207.F7 = _rhs222
			_lhs208.F25 = _rhs223
			_1488_v76 = _rhs224
			var _1489_v77 _dafny.Array
			_ = _1489_v77
			var _len0_45 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_45
			var _nw231 _dafny.Array
			_ = _nw231
			if _len0_45.Cmp(_dafny.Zero) == 0 {
				_nw231 = _dafny.NewArray(_len0_45)
			} else {
				var _init45 func(_dafny.Int) D13 = (func(_1490_v2 *C4) func(_dafny.Int) D13 {
					return func(_1491_i4 _dafny.Int) D13 {
						return (func() D13 {
							if true {
								return Companion_D13_.Create_DC27_(_1490_v2.F25)
							}
							return Companion_D13_.Create_DC27_(_1490_v2.F25)
						})()
					}
				})(_1387_v2)
				_ = _init45
				var _element0_45 = _init45(_dafny.Zero)
				_ = _element0_45
				_nw231 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
				(_nw231).ArraySet1(_element0_45, 0)
				var _nativeLen0_45 = (_len0_45).Int()
				_ = _nativeLen0_45
				for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
					(_nw231).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
				}
			}
			_1489_v77 = _nw231
			var _1492_v78 D13
			_ = _1492_v78
			_1492_v78 = Companion_D13_.Create_DC27_(_1387_v2.F25)
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_1489_v77), 0))
			_ = _index232
			(_1489_v77).ArraySet1(_1492_v78, (_index232).Int())
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_1489_v77), 0))
			_ = _index233
			(_1489_v77).ArraySet1(_1492_v78, (_index233).Int())
			var _1493_v79 _dafny.Sequence
			_ = _1493_v79
			_1493_v79 = _dafny.SeqOf(_1387_v2.F25)
			var _1494_v80 _dafny.Sequence
			_ = _1494_v80
			_1494_v80 = _dafny.SeqOf(_dafny.IntOfUint32((_1493_v79).Cardinality()), _1385_v0)
			var _1495_v81 _dafny.MultiSet
			_ = _1495_v81
			_1495_v81 = _dafny.MultiSetOf(_1385_v0)
			var _1496_v82 _dafny.Map
			_ = _1496_v82
			_1496_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1494_v80).Cardinality()), _1495_v81)
			var _1497_v83 _dafny.Map
			_ = _1497_v83
			_1497_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1385_v0, (func() _dafny.MultiSet {
				if (_1496_v82).Contains(_1385_v0) {
					return (_1496_v82).Get(_1385_v0).(_dafny.MultiSet)
				}
				return _1495_v81
			})())
			var _1498_v84 _dafny.Set
			_ = _1498_v84
			_1498_v84 = _dafny.SetOf(_1493_v79)
			var _1499_v85 _dafny.Array
			_ = _1499_v85
			var _nwElement0_40 bool = (_1495_v81).IsSubsetOf((func() _dafny.MultiSet {
				if (_1497_v83).Contains(_1385_v0) {
					return (_1497_v83).Get(_1385_v0).(_dafny.MultiSet)
				}
				return _1495_v81
			})())
			_ = _nwElement0_40
			var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(15))
			_ = _nw232
			(_nw232).ArraySet1(_nwElement0_40, 0)
			(_nw232).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1493_v79, _dafny.SeqOf((_this).F24())), 1)
			(_nw232).ArraySet1((_1387_v2).F26(), 2)
			(_nw232).ArraySet1((_1385_v0).Cmp(_1385_v0) <= 0, 3)
			(_nw232).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_1387_v2.F25, _1387_v2.F25, _1387_v2.F25), _dafny.SeqOf(_1387_v2.F25, (_1387_v2).F26(), _1387_v2.F25)), 4)
			(_nw232).ArraySet1((!((_1387_v2).F26())) || (false), 5)
			(_nw232).ArraySet1((_this).F24(), 6)
			(_nw232).ArraySet1((_1387_v2).F26(), 7)
			(_nw232).ArraySet1((func() bool {
				if _1387_v2.F25 {
					return (_this).F24()
				}
				return (_this).F24()
			})(), 8)
			(_nw232).ArraySet1((_1493_v79).Select((Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1493_v79).Cardinality()))).Uint32()).(bool), 9)
			(_nw232).ArraySet1((_1387_v2).F26(), 10)
			(_nw232).ArraySet1((_1387_v2).F26(), 11)
			(_nw232).ArraySet1((_1498_v84).Equals(_1498_v84), 12)
			(_nw232).ArraySet1((_1387_v2).F26(), 13)
			(_nw232).ArraySet1((_1387_v2).F26(), 14)
			_1499_v85 = _nw232
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))
			_ = _index234
			(p0).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1385_v0), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-100))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg92 _dafny.Int) interface{} {
					return coer92(arg92)
				}
			}((func(_1500_v75 _dafny.Sequence, _1501_v2 *C4, _1502_v76 _dafny.Map, _1503_v0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_1504_i5 _dafny.Int) _dafny.CodePoint {
					return (_1500_v75).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_1502_v76).Contains(_1501_v2.F25) {
							return (_1502_v76).Get(_1501_v2.F25).(_dafny.Int)
						}
						return _1503_v0
					})(), _dafny.IntOfUint32((_1500_v75).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_1487_v75, _1387_v2, _1488_v76, _1385_v0)))).Cardinality())), (_index234).Int())
			var _1505_v86 _dafny.Sequence
			_ = _1505_v86
			_1505_v86 = _dafny.SeqOf(_1487_v75, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(441))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}(func(_1506_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})))
			var _1507_v87 _dafny.MultiSet
			_ = _1507_v87
			_1507_v87 = _dafny.MultiSetOf(_1499_v85)
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))
			_ = _index235
			var _rhs225 bool = !(((_1507_v87).Union((Companion_D19_.Create_DC45_(_1507_v87)).Dtor_cf59())).IsSubsetOf(_1507_v87))
			_ = _rhs225
			var _rhs226 _dafny.Array = _1499_v85
			_ = _rhs226
			var _rhs227 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1487_v75).Cardinality()), _1385_v0)
			_ = _rhs227
			var _rhs228 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1505_v86, _1505_v86), _dafny.SeqOf(_1487_v75, _1487_v75, _1487_v75))
			_ = _rhs228
			var _lhs209 *GlobalState = globalState
			_ = _lhs209
			var _lhs210 _dafny.Array = p0
			_ = _lhs210
			var _lhs211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))
			_ = _lhs211
			_lhs209.F13 = _rhs225
			_1499_v85 = _rhs226
			(_lhs210).ArraySet1(_rhs227, (_lhs211).Int())
			_1505_v86 = _rhs228
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1499_v85), 0))
			_ = _index236
			(_1499_v85).ArraySet1(_1387_v2.F25, (_index236).Int())
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1499_v85), 0))
			_ = _index237
			(_1499_v85).ArraySet1((_1385_v0).Cmp((_1385_v0).Times(_1385_v0)) < 0, (_index237).Int())
			var _1508_v88 _dafny.Map
			_ = _1508_v88
			_1508_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1487_v75).Cardinality()), _dafny.CodePoint('r'))
			var _1509_v89 _dafny.Set
			_ = _1509_v89
			_1509_v89 = _dafny.SetOf(_1387_v2.F25)
			if ((_1385_v0).Times((func() _dafny.Int {
				if (_1495_v81).Contains((_this).Fm8(true, _1508_v88, _dafny.Companion_Sequence_.Update(_1487_v75, (Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1487_v75).Cardinality()))).Uint32(), Companion_Default___.Fm34(globalState)), globalState)) {
					return (_1495_v81).Multiplicity((_this).Fm8(true, _1508_v88, _dafny.Companion_Sequence_.Update(_1487_v75, (Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1487_v75).Cardinality()))).Uint32(), Companion_Default___.Fm34(globalState)), globalState))
				}
				return (_1509_v89).Cardinality()
			})())).Cmp((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)) == 0 {
				var _1510_v90 D4
				_ = _1510_v90
				_1510_v90 = Companion_D4_.Create_DC10_(_1385_v0, (_dafny.Zero).Minus(_1385_v0), _1387_v2.F25, false, (_1488_v76).Cardinality())
				(globalState).F2 = (_1510_v90).Dtor_cf14()
				var _1511_v91 _dafny.Map
				_ = _1511_v91
				_1511_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1387_v2).F26(), _1487_v75)
				var _1512_v92 _dafny.CodePoint
				_ = _1512_v92
				_1512_v92 = _dafny.CodePoint('y')
				var _1513_v93 _dafny.Array
				_ = _1513_v93
				var _nwElement0_41 _dafny.Sequence = _1487_v75
				_ = _nwElement0_41
				var _nw233 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(29))
				_ = _nw233
				(_nw233).ArraySet1(_nwElement0_41, 0)
				(_nw233).ArraySet1((func() _dafny.Sequence {
					if (_1511_v91).Contains((_1387_v2).F26()) {
						return (_1511_v91).Get((_1387_v2).F26()).(_dafny.Sequence)
					}
					return _1487_v75
				})(), 1)
				(_nw233).ArraySet1(Companion_Default___.Fm1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), globalState), 2)
				(_nw233).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hyp"), 3)
				(_nw233).ArraySet1(_1487_v75, 4)
				(_nw233).ArraySet1(_1487_v75, 5)
				(_nw233).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-529))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg94 _dafny.Int) interface{} {
						return coer94(arg94)
					}
				}(func(_1514_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				})), 6)
				(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rdnqhsuh"), _dafny.UnicodeSeqOfUtf8Bytes("fuy")), 7)
				(_nw233).ArraySet1(_1487_v75, 8)
				(_nw233).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-316))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}((func(_1515_v75 _dafny.Sequence, _1516_p0 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_1517_i8 _dafny.Int) _dafny.CodePoint {
						return (_1515_v75).Select((Companion_Default___.SafeIndex((_1516_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1516_p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1515_v75).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_1487_v75, p0))), 9)
				(_nw233).ArraySet1(_1487_v75, 10)
				(_nw233).ArraySet1(_1487_v75, 11)
				(_nw233).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cepxkgleo"), 12)
				(_nw233).ArraySet1(_dafny.Companion_Sequence_.Update(_1487_v75, (Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1487_v75).Cardinality()))).Uint32(), _dafny.CodePoint('d')), 13)
				(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1487_v75, (_1505_v86).Select((Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1505_v86).Cardinality()))).Uint32()).(_dafny.Sequence)), 14)
				(_nw233).ArraySet1(_1487_v75, 15)
				(_nw233).ArraySet1(_1487_v75, 16)
				(_nw233).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}(func(_1518_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})), 17)
				(_nw233).ArraySet1(_1487_v75, 18)
				(_nw233).ArraySet1(_dafny.Companion_Sequence_.Update(_1487_v75, (Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1487_v75).Cardinality()))).Uint32(), _1512_v92), 19)
				(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1487_v75, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(807))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}((func(_1519_v92 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1520_i10 _dafny.Int) _dafny.CodePoint {
						return _1519_v92
					}
				})(_1512_v92)))), 20)
				(_nw233).ArraySet1(_1487_v75, 21)
				(_nw233).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(885))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}(func(_1521_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				})), 22)
				(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate((Companion_D13_.Create_DC26_(_1487_v75, _1387_v2.F25, _1387_v2.F25)).Dtor_cf33(), _1487_v75), 23)
				(_nw233).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ctaxiqtj"), 24)
				(_nw233).ArraySet1(_1487_v75, 25)
				(_nw233).ArraySet1(_1487_v75, 26)
				(_nw233).ArraySet1(_1487_v75, 27)
				(_nw233).ArraySet1(_1487_v75, 28)
				_1513_v93 = _nw233
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1513_v93), 0))
				_ = _index238
				(_1513_v93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rckcmrpbm"), _1487_v75), (_index238).Int())
				var _1522_v94 _dafny.Map
				_ = _1522_v94
				_1522_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1385_v0, _1487_v75)
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1513_v93), 0))
				_ = _index239
				(_1513_v93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("srrgrldpd"), (func() _dafny.Sequence {
					if (_1522_v94).Contains(_1385_v0) {
						return (_1522_v94).Get(_1385_v0).(_dafny.Sequence)
					}
					return _1487_v75
				})()), (_index239).Int())
				(globalState).F16 = Companion_Default___.SafeModuloInt((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
				var _1523_v95 _dafny.Array
				_ = _1523_v95
				var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
				_ = _nw234
				_1523_v95 = _nw234
				var _1524_v96 _dafny.Map
				_ = _1524_v96
				_1524_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1523_v95)
				_1524_v96 = (_1524_v96).Update((_1387_v2).F26(), _1523_v95)
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))
				_ = _index240
				(p0).ArraySet1((Companion_Default___.SafeModuloInt(_1385_v0, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))).Plus(_1385_v0), (_index240).Int())
			} else {
				var _1525_v97 D14
				_ = _1525_v97
				_1525_v97 = Companion_D14_.Create_DC30_(_1494_v80, _1493_v79)
				var _1526_v98 _dafny.MultiSet
				_ = _1526_v98
				_1526_v98 = _dafny.MultiSetOf(_dafny.SeqOf(Companion_D14_.Create_DC30_(_1494_v80, _1493_v79), _1525_v97, _1525_v97))
				var _1527_v99 _dafny.CodePoint
				_ = _1527_v99
				_1527_v99 = _dafny.CodePoint('x')
				_1526_v98 = (func() _dafny.MultiSet {
					if false {
						return (_1526_v98).Difference(_1526_v98)
					}
					return ((_1526_v98).Update(Companion_Default___.Fm46(_1387_v2.F25, _1527_v99, globalState), Companion_Default___.Abs(_dafny.IntOfInt64(843)))).Difference(_1526_v98)
				})()
				_1488_v76 = (_1488_v76).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_1493_v79, _1493_v79), _dafny.IntOfInt64(931))
				var _1528_v100 *C2
				_ = _1528_v100
				var _nw235 *C2 = New_C2_()
				_ = _nw235
				_nw235.Ctor__((func() _dafny.Int {
					if (_1499_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1499_v85), 0))).Int()).(bool) {
						return _dafny.IntOfInt64(614)
					}
					return (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)
				})())
				_1528_v100 = _nw235
				_1528_v100 = _1528_v100
				(globalState).F13 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_1487_v75, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1493_v79).Cardinality()), _dafny.IntOfUint32((_1487_v75).Cardinality()))).Uint32(), _1527_v99), _dafny.UnicodeSeqOfUtf8Bytes("vjgipyft"))
				(_1387_v2).F25 = (_1386_v1).IsSubsetOf(_1386_v1)
			}
		}
		var _1529_v101 _dafny.Sequence
		_ = _1529_v101
		_1529_v101 = _dafny.UnicodeSeqOfUtf8Bytes("ckhhh")
		var _1530_v102 _dafny.Map
		_ = _1530_v102
		_1530_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1529_v101, (_this).F24())
		var _hi7 _dafny.Int = _1385_v0
		_ = _hi7
		for _1531_i12 := (_1530_v102).Cardinality(); _1531_i12.Cmp(_hi7) < 0; _1531_i12 = _1531_i12.Plus(_dafny.One) {
			_1386_v1 = (_1386_v1).Difference(_dafny.SetOf((_dafny.Zero).Minus(_1385_v0)))
			(_1387_v2).F25 = (_1387_v2).F26()
			if _1387_v2.F25 {
				var _1532_v103 _dafny.Map
				_ = _1532_v103
				_1532_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1385_v0, (_1387_v2).F26())
				var _1533_v104 D1
				_ = _1533_v104
				_1533_v104 = Companion_D1_.Create_DC2_(_1532_v103)
				_1533_v104 = _1533_v104
				var _1534_v105 _dafny.Array
				_ = _1534_v105
				var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw236
				_1534_v105 = _nw236
				var _1535_v106 _dafny.Array
				_ = _1535_v106
				var _nwElement0_42 bool = _1387_v2.F25
				_ = _nwElement0_42
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(27))
				_ = _nw237
				(_nw237).ArraySet1(_nwElement0_42, 0)
				(_nw237).ArraySet1((_1387_v2).F26(), 1)
				(_nw237).ArraySet1(_1387_v2.F25, 2)
				(_nw237).ArraySet1(_1387_v2.F25, 3)
				(_nw237).ArraySet1((_1387_v2).F26(), 4)
				(_nw237).ArraySet1((_this).F24(), 5)
				(_nw237).ArraySet1((_1387_v2).F26(), 6)
				(_nw237).ArraySet1(!(_1387_v2.F25), 7)
				(_nw237).ArraySet1(!(_1387_v2.F25), 8)
				(_nw237).ArraySet1((_1387_v2).F26(), 9)
				(_nw237).ArraySet1((_1387_v2).F26(), 10)
				(_nw237).ArraySet1((_this).Fm7(_1387_v2.F25, _1531_i12, false, globalState), 11)
				(_nw237).ArraySet1(_1387_v2.F25, 12)
				(_nw237).ArraySet1((_1387_v2).F26(), 13)
				(_nw237).ArraySet1((_1387_v2).F26(), 14)
				(_nw237).ArraySet1((_1387_v2).F26(), 15)
				(_nw237).ArraySet1((_this).F24(), 16)
				(_nw237).ArraySet1((_this).F24(), 17)
				(_nw237).ArraySet1((_1387_v2).F26(), 18)
				(_nw237).ArraySet1((_1387_v2).F26(), 19)
				(_nw237).ArraySet1((_1387_v2).F26(), 20)
				(_nw237).ArraySet1(!(_1387_v2.F25), 21)
				(_nw237).ArraySet1((_this).Fm6(_dafny.IntOfUint32((_1529_v101).Cardinality()), globalState), 22)
				(_nw237).ArraySet1(_1387_v2.F25, 23)
				(_nw237).ArraySet1(!((_1387_v2).F26()), 24)
				(_nw237).ArraySet1((_1387_v2).F26(), 25)
				(_nw237).ArraySet1((_1387_v2).F26(), 26)
				_1535_v106 = _nw237
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_1534_v105), 0))
				_ = _index241
				(_1534_v105).ArraySet1(_1535_v106, (_index241).Int())
				var _1536_v107 _dafny.CodePoint
				_ = _1536_v107
				_1536_v107 = _dafny.CodePoint('o')
				var _1537_v108 D13
				_ = _1537_v108
				_1537_v108 = Companion_D13_.Create_DC26_(_1529_v101, (_1387_v2).Fm9(_1387_v2.F25, globalState), _1387_v2.F25)
				var _1538_v109 D19
				_ = _1538_v109
				_1538_v109 = Companion_D19_.Create_DC46_(_1387_v2.F25, _1385_v0)
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_1534_v105), 0))
				_ = _index242
				var _rhs229 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate((_1537_v108).Dtor_cf33(), _dafny.Companion_Sequence_.Update(_1529_v101, (Companion_Default___.SafeIndex(_1531_i12, _dafny.IntOfUint32((_1529_v101).Cardinality()))).Uint32(), _1536_v107)), _1536_v107)
				_ = _rhs229
				var _rhs230 bool = (_1387_v2).F26()
				_ = _rhs230
				var _rhs231 _dafny.Int = (func() _dafny.Int {
					if _1387_v2.F25 {
						return (func() _dafny.Int {
							if (_1387_v2).F26() {
								return _1531_i12
							}
							return (_1538_v109).Dtor_cf61()
						})()
					}
					return _1385_v0
				})()
				_ = _rhs231
				var _rhs232 _dafny.Array = _1535_v106
				_ = _rhs232
				var _lhs212 *GlobalState = globalState
				_ = _lhs212
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				var _lhs215 _dafny.Array = _1534_v105
				_ = _lhs215
				var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_1534_v105), 0))
				_ = _lhs216
				_lhs212.F2 = _rhs229
				_lhs213.F0 = _rhs230
				_lhs214.F7 = _rhs231
				(_lhs215).ArraySet1(_rhs232, (_lhs216).Int())
				var _1539_v110 _dafny.Map
				_ = _1539_v110
				_1539_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1387_v2).F26(), _1385_v0)
				var _1540_v111 _dafny.Sequence
				_ = _1540_v111
				_1540_v111 = _dafny.SeqOf(Companion_Default___.Fm1(_1531_i12, globalState), _1529_v101)
				var _1541_v112 _dafny.Map
				_ = _1541_v112
				_1541_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1385_v0).Times((_1539_v110).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1540_v111, (Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1540_v111).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("je")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xydhcd"))))).Cardinality())))
				var _1542_v113 D19
				_ = _1542_v113
				_1542_v113 = Companion_D19_.Create_DC47_(_1533_v104, _dafny.IntOfInt64(-809), true)
				var _1543_v114 _dafny.Map
				_ = _1543_v114
				_1543_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_1534_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_1534_v105), 0))).Int())), _1387_v2.F25)
				(globalState).F16 = (func() _dafny.Int {
					if (_1541_v112).Contains((_1542_v113).Dtor_cf63()) {
						return (_1541_v112).Get((_1542_v113).Dtor_cf63()).(_dafny.Int)
					}
					return (_dafny.Zero).Minus((func() _dafny.Int {
						if (_1387_v2).F26() {
							return (_dafny.Zero).Minus((_1440_v41).Cardinality())
						}
						return (_1543_v114).Cardinality()
					})())
				})()
				var _1544_v115 _dafny.Array
				_ = _1544_v115
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw238
				_1544_v115 = _nw238
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_46
				var _nw239 _dafny.Array
				_ = _nw239
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw239 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Int = (func(_1545_i12 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1546_i13 _dafny.Int) _dafny.Int {
							return (_1546_i13).Plus(_1545_i12)
						}
					})(_1531_i12)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw239 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw239).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw239).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1544_v115 = _nw239
				var _1547_v116 *C0
				_ = _1547_v116
				var _nw240 *C0 = New_C0_()
				_ = _nw240
				_nw240.Ctor__(Companion_Default___.Fm2(_1440_v41, globalState))
				_1547_v116 = _nw240
				var _1548_v117 _dafny.Map
				_ = _1548_v117
				_1548_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1387_v2.F25, _1547_v116)
				_1548_v117 = (_1548_v117).Update(_1547_v116.F27, _1547_v116)
			} else {
				var _1549_v118 _dafny.Array
				_ = _1549_v118
				var _nw241 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(23))
				_ = _nw241
				_1549_v118 = _nw241
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1549_v118), 0))
				_ = _index243
				(_1549_v118).ArraySet1(Companion_D6_.Create_DC12_(_dafny.SeqOf(p0)), (_index243).Int())
				var _1550_v119 _dafny.CodePoint
				_ = _1550_v119
				_1550_v119 = _dafny.CodePoint('w')
				var _1551_v120 _dafny.Sequence
				_ = _1551_v120
				_1551_v120 = _dafny.SeqOf(p0, p0)
				var _1552_v121 D6
				_ = _1552_v121
				_1552_v121 = Companion_D6_.Create_DC12_(_1551_v120)
				var _1553_v122 _dafny.Sequence
				_ = _1553_v122
				_1553_v122 = _dafny.SeqOf(_1385_v0)
				var _1554_v123 _dafny.Map
				_ = _1554_v123
				_1554_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1553_v122).Select((Companion_Default___.SafeIndex(_1531_i12, _dafny.IntOfUint32((_1553_v122).Cardinality()))).Uint32()).(_dafny.Int), _1387_v2.F25)
				var _1555_v124 _dafny.Set
				_ = _1555_v124
				_1555_v124 = _dafny.SetOf(_1550_v119)
				var _1556_v125 D4
				_ = _1556_v125
				_1556_v125 = Companion_D4_.Create_DC10_(_1385_v0, (_dafny.MultiSetOf((_dafny.Zero).Minus((_1555_v124).Cardinality()), _1531_i12)).Cardinality(), _1387_v2.F25, _1387_v2.F25, _1531_i12)
				var _1557_v126 *C2
				_ = _1557_v126
				var _nw242 *C2 = New_C2_()
				_ = _nw242
				_nw242.Ctor__((_1554_v123).Cardinality())
				_1557_v126 = _nw242
				var _1558_v127 _dafny.Sequence
				_ = _1558_v127
				_1558_v127 = _dafny.SeqOf((_1387_v2).F26())
				var _1559_v128 _dafny.Sequence
				_ = _1559_v128
				_1559_v128 = _dafny.SeqOf(_1558_v127)
				var _pat_let_tv31 = _1385_v0
				_ = _pat_let_tv31
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1549_v118), 0))
				_ = _index244
				var _rhs233 D6 = (func() D6 {
					if !_dafny.Companion_Sequence_.Contains(_1529_v101, _1550_v119) {
						return _1552_v121
					}
					return _1552_v121
				})()
				_ = _rhs233
				var _rhs234 bool = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ys")).Cardinality())).Cmp((_dafny.Zero).Minus((_this).Fm8((func() bool {
					if (_1554_v123).Contains(_1531_i12) {
						return (_1554_v123).Get(_1531_i12).(bool)
					}
					return (_1387_v2).F26()
				})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let24_0 D4) D4 {
					return func(_1560_dt__update__tmp_h1 D4) D4 {
						return func(_pat_let25_0 _dafny.Int) D4 {
							return func(_1561_dt__update_hcf13_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC10_((_1560_dt__update__tmp_h1).Dtor_cf12(), _1561_dt__update_hcf13_h0, (_1560_dt__update__tmp_h1).Dtor_cf14(), (_1560_dt__update__tmp_h1).Dtor_cf15(), (_1560_dt__update__tmp_h1).Dtor_cf16())
							}(_pat_let25_0)
						}(_pat_let_tv31)
					}(_pat_let24_0)
				}(_1556_v125)).Dtor_cf12(), _1550_v119), _1529_v101, globalState))) < 0
				_ = _rhs234
				var _rhs235 _dafny.Int = _1531_i12
				_ = _rhs235
				var _rhs236 bool = (func() bool {
					if (_1440_v41).Contains((_dafny.SetOf(_1557_v126, _1557_v126, _1557_v126)).IsProperSubsetOf(_dafny.SetOf(_1557_v126))) {
						return (_1440_v41).Get((_dafny.SetOf(_1557_v126, _1557_v126, _1557_v126)).IsProperSubsetOf(_dafny.SetOf(_1557_v126))).(bool)
					}
					return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_1558_v127, _1558_v127), _1559_v128)
				})()
				_ = _rhs236
				var _rhs237 bool = _1387_v2.F25
				_ = _rhs237
				var _lhs217 _dafny.Array = _1549_v118
				_ = _lhs217
				var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1549_v118), 0))
				_ = _lhs218
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				var _lhs220 *GlobalState = globalState
				_ = _lhs220
				var _lhs221 *GlobalState = globalState
				_ = _lhs221
				var _lhs222 *GlobalState = globalState
				_ = _lhs222
				(_lhs217).ArraySet1(_rhs233, (_lhs218).Int())
				_lhs219.F2 = _rhs234
				_lhs220.F7 = _rhs235
				_lhs221.F2 = _rhs236
				_lhs222.F13 = _rhs237
				(globalState).F21 = _1387_v2.F25
				_1385_v0 = _1531_i12
				var _1562_v129 _dafny.Set
				_ = _1562_v129
				_1562_v129 = _dafny.SetOf((_1387_v2).F26(), (_1387_v2).F26())
				var _1563_v130 _dafny.Sequence
				_ = _1563_v130
				_1563_v130 = _dafny.SeqOf(_1562_v129)
				var _1564_v131 _dafny.MultiSet
				_ = _1564_v131
				_1564_v131 = _dafny.MultiSetOf(((_1563_v130).Select((Companion_Default___.SafeIndex(_1531_i12, _dafny.IntOfUint32((_1563_v130).Cardinality()))).Uint32()).(_dafny.Set)).Intersection(_1562_v129), _1562_v129, _dafny.SetOf(!(Companion_Default___.Fm2(_1440_v41, globalState))))
				_1564_v131 = _dafny.MultiSetOf(_1562_v129)
				var _1565_v132 _dafny.Array
				_ = _1565_v132
				var _nw243 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw243
				_1565_v132 = _nw243
				var _1566_v133 _dafny.Map
				_ = _1566_v133
				_1566_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1531_i12, _dafny.CodePoint('c'))
				var _1567_v134 _dafny.Map
				_ = _1567_v134
				_1567_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1565_v132, (_this).Fm6((_this).Fm8((_1387_v2).F26(), _1566_v133, _1529_v101, globalState), globalState))
				_1567_v134 = (_1567_v134).Update(_1565_v132, !(_1387_v2.F25))
			}
			(globalState).F2 = _1387_v2.F25
		}
		var _1568_v135 _dafny.Set
		_ = _1568_v135
		_1568_v135 = _dafny.SetOf((_this).F24(), _1387_v2.F25, (_this).F24(), true)
		var _1569_v136 _dafny.Map
		_ = _1569_v136
		_1569_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1568_v135).Cardinality())
		var _hi8 _dafny.Int = ((_1569_v136).Merge(_1569_v136)).Cardinality()
		_ = _hi8
		for _1570_i14 := (_1385_v0).Plus(_1385_v0); _1570_i14.Cmp(_hi8) < 0; _1570_i14 = _1570_i14.Plus(_dafny.One) {
			var _1571_v137 _dafny.Set
			_ = _1571_v137
			_1571_v137 = (_1568_v135).Intersection(Companion_Default___.Fm31(_1385_v0, (_1387_v2).F26(), globalState))
			_1571_v137 = _1568_v135
			var _1572_v138 D13
			_ = _1572_v138
			_1572_v138 = Companion_D13_.Create_DC27_(false)
			_1572_v138 = (func() D13 {
				if (_1387_v2).F26() {
					return _1572_v138
				}
				return _1572_v138
			})()
			var _1573_v139 T0
			_ = _1573_v139
			var _nw244 *C3 = New_C3_()
			_ = _nw244
			_nw244.Ctor__()
			_1573_v139 = _nw244
			var _nw245 *C3 = New_C3_()
			_ = _nw245
			_nw245.Ctor__()
			_1573_v139 = _nw245
			var _1574_v140 _dafny.Array
			_ = _1574_v140
			var _nw246 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw246
			_1574_v140 = _nw246
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1574_v140), 0))
			_ = _index245
			(_1574_v140).ArraySet1(false, (_index245).Int())
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1574_v140), 0))
			_ = _index246
			var _rhs238 bool = (_1387_v2).F26()
			_ = _rhs238
			var _rhs239 _dafny.Int = _1385_v0
			_ = _rhs239
			var _rhs240 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1570_i14))
			_ = _rhs240
			var _rhs241 bool = !(false)
			_ = _rhs241
			var _lhs223 *GlobalState = globalState
			_ = _lhs223
			var _lhs224 *GlobalState = globalState
			_ = _lhs224
			var _lhs225 *GlobalState = globalState
			_ = _lhs225
			var _lhs226 _dafny.Array = _1574_v140
			_ = _lhs226
			var _lhs227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1574_v140), 0))
			_ = _lhs227
			_lhs223.F13 = _rhs238
			_lhs224.F7 = _rhs239
			_lhs225.F16 = _rhs240
			(_lhs226).ArraySet1(_rhs241, (_lhs227).Int())
		}
		var _hi9 _dafny.Int = _1385_v0
		_ = _hi9
		for _1575_i15 := _dafny.IntOfInt64(169); _1575_i15.Cmp(_hi9) < 0; _1575_i15 = _1575_i15.Plus(_dafny.One) {
			if (_1387_v2).F26() {
				var _1576_v141 _dafny.Set
				_ = _1576_v141
				_1576_v141 = _1568_v135
				var _1577_v142 _dafny.Array
				_ = _1577_v142
				var _nwElement0_43 bool = (_dafny.SetOf((_1387_v2).F26(), true)).IsProperSubsetOf(_1568_v135)
				_ = _nwElement0_43
				var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(2))
				_ = _nw247
				(_nw247).ArraySet1(_nwElement0_43, 0)
				(_nw247).ArraySet1((_1387_v2).F26(), 1)
				_1577_v142 = _nw247
				var _1578_v143 _dafny.Sequence
				_ = _1578_v143
				_1578_v143 = _dafny.SeqOf((_1387_v2).F26(), _1387_v2.F25)
				var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_1577_v142), 0))
				_ = _index247
				(_1577_v142).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1578_v143, _1578_v143), (_index247).Int())
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_1577_v142), 0))
				_ = _index248
				var _rhs242 _dafny.Set = (func() _dafny.Set {
					if (_1385_v0).Cmp(_1575_i15) >= 0 {
						return _1576_v141
					}
					return _1576_v141
				})()
				_ = _rhs242
				var _rhs243 _dafny.Set = (func() _dafny.Set {
					var _coll60 = _dafny.NewBuilder()
					_ = _coll60
					for _iter72 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(872), _dafny.IntOfInt64(204))); ; {
						_compr_60, _ok72 := _iter72()
						if !_ok72 {
							break
						}
						var _1579_v144 _dafny.Int
						_1579_v144 = interface{}(_compr_60).(_dafny.Int)
						if ((_dafny.IntOfInt64(872)).Cmp(_1579_v144) <= 0) && ((_1579_v144).Cmp(_dafny.IntOfInt64(204)) < 0) {
							_coll60.Add((_1579_v144).Minus(_1575_i15))
						}
					}
					return _coll60.ToSet()
				}()).Union((_dafny.SetOf(_1575_i15)).Union(_1386_v1))
				_ = _rhs243
				var _rhs244 bool = (_this).F24()
				_ = _rhs244
				var _rhs245 bool = _1387_v2.F25
				_ = _rhs245
				var _lhs228 *C4 = _1387_v2
				_ = _lhs228
				var _lhs229 _dafny.Array = _1577_v142
				_ = _lhs229
				var _lhs230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_1577_v142), 0))
				_ = _lhs230
				_1576_v141 = _rhs242
				_1386_v1 = _rhs243
				_lhs228.F25 = _rhs244
				(_lhs229).ArraySet1(_rhs245, (_lhs230).Int())
				(globalState).F2 = (func() bool {
					if (_1387_v2).F26() {
						return (_this).F24()
					}
					return (_1577_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_1577_v142), 0))).Int()).(bool)
				})()
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_1577_v142), 0))
				_ = _index249
				(_1577_v142).ArraySet1((func() bool {
					if !(_1387_v2.F25) {
						return (_1385_v0).Cmp((_1386_v1).Cardinality()) != 0
					}
					return ((_1568_v135).Cardinality()).Cmp(_1575_i15) != 0
				})(), (_index249).Int())
				var _1580_v145 _dafny.Sequence
				_ = _1580_v145
				_1580_v145 = _dafny.SeqOf(_dafny.IntOfUint32((_1529_v101).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1387_v2.F25, _1385_v0)).Cardinality(), _1575_i15, _1385_v0, _1575_i15)
				_1440_v41 = (_1440_v41).Update(_1387_v2.F25, _dafny.Companion_Sequence_.Contains(_1580_v145, (_dafny.Zero).Minus(_1385_v0)))
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((p0), 0))
				_ = _index250
				(p0).ArraySet1((_dafny.Zero).Minus(_1575_i15), (_index250).Int())
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((p0), 0))
				_ = _index251
				(p0).ArraySet1(_1385_v0, (_index251).Int())
			} else {
				var _1581_v146 _dafny.Sequence
				_ = _1581_v146
				_1581_v146 = _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(867)))
				var _rhs246 _dafny.Set = ((_1568_v135).Union(_1568_v135)).Intersection(Companion_Default___.Fm14(_1575_i15, (_1387_v2).F26(), globalState))
				_ = _rhs246
				var _rhs247 _dafny.Set = (_1581_v146).Select((Companion_Default___.SafeIndex(_1385_v0, _dafny.IntOfUint32((_1581_v146).Cardinality()))).Uint32()).(_dafny.Set)
				_ = _rhs247
				_1568_v135 = _rhs246
				_1386_v1 = _rhs247
				(globalState).F7 = ((func() _dafny.Int {
					if (_1569_v136).Contains((_1387_v2).F26()) {
						return (_1569_v136).Get((_1387_v2).F26()).(_dafny.Int)
					}
					return _1575_i15
				})()).Times((func() _dafny.Int {
					if _1387_v2.F25 {
						return _1385_v0
					}
					return _1575_i15
				})())
				(globalState).F2 = _1387_v2.F25
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((p0), 0))
				_ = _index252
				(p0).ArraySet1(_1575_i15, (_index252).Int())
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((p0), 0))
				_ = _index253
				(p0).ArraySet1(Companion_Default___.SafeModuloInt(_1385_v0, _1385_v0), (_index253).Int())
				(_1387_v2).M5(Companion_Default___.Fm2(_1440_v41, globalState), globalState)
			}
			(globalState).F7 = (_1385_v0).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(559), _1385_v0))
			var _1582_v147 _dafny.Array
			_ = _1582_v147
			var _nwElement0_44 _dafny.Map = _1440_v41
			_ = _nwElement0_44
			var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(22))
			_ = _nw248
			(_nw248).ArraySet1(_nwElement0_44, 0)
			(_nw248).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1387_v2.F25, true)).Merge(_1440_v41), 1)
			(_nw248).ArraySet1((Companion_Default___.Fm20(_1385_v0, (_this).Fm6(_1575_i15, globalState), (_1387_v2).F26(), _1387_v2.F25, globalState)).Update((_1387_v2).F26(), (_this).Fm7(_1387_v2.F25, (_dafny.Zero).Minus(_1385_v0), !(!((_1387_v2).F26())), globalState)), 2)
			(_nw248).ArraySet1((_1440_v41).Update(!((_this).F24()), (_1387_v2).F26()), 3)
			(_nw248).ArraySet1(_1440_v41, 4)
			(_nw248).ArraySet1((_1440_v41).Merge(_1440_v41), 5)
			(_nw248).ArraySet1(_1440_v41, 6)
			(_nw248).ArraySet1(_1440_v41, 7)
			(_nw248).ArraySet1(_1440_v41, 8)
			(_nw248).ArraySet1(_1440_v41, 9)
			(_nw248).ArraySet1((_1440_v41).Merge(_1440_v41), 10)
			(_nw248).ArraySet1(_1440_v41, 11)
			(_nw248).ArraySet1(_1440_v41, 12)
			(_nw248).ArraySet1(_1440_v41, 13)
			(_nw248).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1387_v2).F26(), false), 14)
			(_nw248).ArraySet1((_1440_v41).Merge(_1440_v41), 15)
			(_nw248).ArraySet1((_1440_v41).Merge(_1440_v41), 16)
			(_nw248).ArraySet1((_1440_v41).Merge(_1440_v41), 17)
			(_nw248).ArraySet1((_1440_v41).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1387_v2).F26()), _1387_v2.F25)), 18)
			(_nw248).ArraySet1(_1440_v41, 19)
			(_nw248).ArraySet1(_1440_v41, 20)
			(_nw248).ArraySet1((_1440_v41).Merge(_1440_v41), 21)
			_1582_v147 = _nw248
			var _1583_v148 _dafny.Array
			_ = _1583_v148
			var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(13))
			_ = _nw249
			_1583_v148 = _nw249
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1583_v148), 0))
			_ = _index254
			(_1583_v148).ArraySet1((_1568_v135).Union(_1568_v135), (_index254).Int())
			var _1584_v149 _dafny.Sequence
			_ = _1584_v149
			_1584_v149 = _dafny.SeqOf(_1575_i15)
			var _1585_v150 _dafny.Sequence
			_ = _1585_v150
			_1585_v150 = _dafny.SeqOf((_this).F24())
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1583_v148), 0))
			_ = _index255
			var _rhs248 _dafny.Array = _1582_v147
			_ = _rhs248
			var _rhs249 bool = _1387_v2.F25
			_ = _rhs249
			var _rhs250 bool = (_1585_v150).Select((Companion_Default___.SafeIndex(_1575_i15, _dafny.IntOfUint32((_1585_v150).Cardinality()))).Uint32()).(bool)
			_ = _rhs250
			var _rhs251 _dafny.Set = _1568_v135
			_ = _rhs251
			var _rhs252 _dafny.Sequence = _1584_v149
			_ = _rhs252
			var _lhs231 *GlobalState = globalState
			_ = _lhs231
			var _lhs232 *GlobalState = globalState
			_ = _lhs232
			var _lhs233 _dafny.Array = _1583_v148
			_ = _lhs233
			var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1583_v148), 0))
			_ = _lhs234
			_1582_v147 = _rhs248
			_lhs231.F0 = _rhs249
			_lhs232.F2 = _rhs250
			(_lhs233).ArraySet1(_rhs251, (_lhs234).Int())
			_1584_v149 = _rhs252
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((p0), 0))
			_ = _index256
			(p0).ArraySet1(_1385_v0, (_index256).Int())
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((p0), 0))
			_ = _index257
			var _rhs253 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1385_v0, _dafny.IntOfInt64(471)))
			_ = _rhs253
			var _rhs254 bool = Companion_Default___.Fm2(_1440_v41, globalState)
			_ = _rhs254
			var _lhs235 _dafny.Array = p0
			_ = _lhs235
			var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((p0), 0))
			_ = _lhs236
			var _lhs237 *GlobalState = globalState
			_ = _lhs237
			(_lhs235).ArraySet1(_rhs253, (_lhs236).Int())
			_lhs237.F21 = _rhs254
		}
	}
}
func (_this *C5) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _1586_v0 _dafny.Sequence
		_ = _1586_v0
		_1586_v0 = _dafny.SeqOf(p0)
		var _1587_v1 _dafny.Sequence
		_ = _1587_v1
		_1587_v1 = _dafny.SeqOf((_1586_v0).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1586_v0).Cardinality()))).Uint32()).(_dafny.Int))
		var _hi10 _dafny.Int = p0
		_ = _hi10
		for _1588_i0 := (_dafny.IntOfUint32((_1587_v1).Cardinality())).Minus(p0); _1588_i0.Cmp(_hi10) < 0; _1588_i0 = _1588_i0.Plus(_dafny.One) {
			(globalState).F21 = (_this).F24()
			var _1589_v2 *C3
			_ = _1589_v2
			var _nw250 *C3 = New_C3_()
			_ = _nw250
			_nw250.Ctor__()
			_1589_v2 = _nw250
			var _1590_v3 *C2
			_ = _1590_v3
			var _nw251 *C2 = New_C2_()
			_ = _nw251
			_nw251.Ctor__(p0)
			_1590_v3 = _nw251
			if true {
				(globalState).F0 = true
				var _1591_v4 _dafny.CodePoint
				_ = _1591_v4
				_1591_v4 = _dafny.CodePoint('k')
				var _1592_v5 _dafny.MultiSet
				_ = _1592_v5
				_1592_v5 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1591_v4)).Cardinality())
				var _1593_v6 _dafny.Set
				_ = _1593_v6
				_1593_v6 = _dafny.SetOf(Companion_Default___.SafeDivisionInt((_1590_v3).F28(), Companion_Default___.Fm0((_1590_v3).F28(), (_1592_v5).Cardinality(), _dafny.IntOfUint32((_1587_v1).Cardinality()), (_this).F24(), globalState)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm21(globalState)).Cardinality()), (_1590_v3).F28()), _1586_v0)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1586_v0, _1587_v1)).Cardinality()))
				var _1594_v7 _dafny.Map
				_ = _1594_v7
				_1594_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1590_v3).F28(), _1591_v4)
				var _1595_v8 _dafny.Sequence
				_ = _1595_v8
				_1595_v8 = _dafny.UnicodeSeqOfUtf8Bytes("gffxeber")
				var _1596_v9 _dafny.Map
				_ = _1596_v9
				_1596_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1590_v3).F28(), (_this).F24())
				var _1597_v11 _dafny.Sequence
				_ = _1597_v11
				_1597_v11 = _dafny.SeqOf((_this).F24(), (_this).F24())
				var _1598_v12 _dafny.Sequence
				_ = _1598_v12
				_1598_v12 = _1597_v11
				var _rhs255 bool = (_this).F24()
				_ = _rhs255
				var _rhs256 bool = (p0).Cmp((p0).Minus((_this).Fm8(false, _1594_v7, _1595_v8, globalState))) <= 0
				_ = _rhs256
				var _rhs257 _dafny.Set = ((_1593_v6).Union(_1593_v6)).Union(func() _dafny.Set {
					var _coll61 = _dafny.NewBuilder()
					_ = _coll61
					for _iter73 := _dafny.Iterate((_1596_v9).Keys().Elements()); ; {
						_compr_61, _ok73 := _iter73()
						if !_ok73 {
							break
						}
						var _1599_v10 _dafny.Int
						_1599_v10 = interface{}(_compr_61).(_dafny.Int)
						if (_1596_v9).Contains(_1599_v10) {
							_coll61.Add(Companion_Default___.SafeModuloInt(_1599_v10, (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg99 _dafny.Int) interface{} {
									return coer99(arg99)
								}
							}(func(_1600_i1 _dafny.Int) _dafny.Int {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cftvrnjd")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(96), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.IntOfInt64(379))).Cardinality())).Cardinality())).Cardinality()
							})))).Cardinality()))
						}
					}
					return _coll61.ToSet()
				}())
				_ = _rhs257
				var _rhs258 _dafny.Int = ((_1590_v3).Fm17((_this).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(187))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}(func(_1601_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality())
				})), _1588_i0, _1598_v12, globalState)).Plus(Companion_Default___.SafeModuloInt((_1590_v3).F28(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)).Cardinality()))
				_ = _rhs258
				var _lhs238 *GlobalState = globalState
				_ = _lhs238
				var _lhs239 *GlobalState = globalState
				_ = _lhs239
				_lhs238.F21 = _rhs255
				r1 = _rhs256
				_1593_v6 = _rhs257
				_lhs239.F7 = _rhs258
				var _1602_v13 _dafny.Array
				_ = _1602_v13
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_47
				var _nw252 _dafny.Array
				_ = _nw252
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw252 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Sequence = (func(_1603_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1604_i3 _dafny.Int) _dafny.Sequence {
							return _1603_v8
						}
					})(_1595_v8)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw252 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw252).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw252).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1602_v13 = _nw252
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1602_v13), 0))
				_ = _index258
				(_1602_v13).ArraySet1(_1595_v8, (_index258).Int())
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1602_v13), 0))
				_ = _index259
				(_1602_v13).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(777))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1605_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1606_i4 _dafny.Int) _dafny.CodePoint {
						return _1605_v4
					}
				})(_1591_v4))), (_index259).Int())
				var _1607_v14 _dafny.Array
				_ = _1607_v14
				var _nw253 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw253
				_1607_v14 = _nw253
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_1607_v14), 0))
				_ = _index260
				(_1607_v14).ArraySet1((_this).F24(), (_index260).Int())
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_1607_v14), 0))
				_ = _index261
				(_1607_v14).ArraySet1((_dafny.IntOfUint32((_1586_v0).Cardinality())).Cmp(_dafny.IntOfInt64(-113)) <= 0, (_index261).Int())
				var _1608_v16 _dafny.MultiSet
				_ = _1608_v16
				_1608_v16 = _dafny.MultiSetOf(_dafny.SeqOf(p0))
				var _1609_v17 D16
				_ = _1609_v17
				_1609_v17 = Companion_D16_.Create_DC35_(func() _dafny.Map {
					var _coll62 = _dafny.NewMapBuilder()
					_ = _coll62
					for _iter74 := _dafny.Iterate((_1608_v16).Elements()); ; {
						_compr_62, _ok74 := _iter74()
						if !_ok74 {
							break
						}
						var _1610_v15 _dafny.Sequence
						_1610_v15 = interface{}(_compr_62).(_dafny.Sequence)
						if (_1608_v16).Contains(_1610_v15) {
							_coll62.Add(_1610_v15, _1597_v11)
						}
					}
					return _coll62.ToMap()
				}())
				var _1611_v18 D16
				_ = _1611_v18
				_1611_v18 = Companion_D16_.Create_DC37_(_1609_v17)
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1602_v13), 0))
				_ = _index262
				var _rhs259 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1595_v8, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_1588_i0, globalState), _1595_v8))
				_ = _rhs259
				var _rhs260 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_1602_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1602_v13), 0))).Int()).(_dafny.Sequence), _1595_v8), _1595_v8)
				_ = _rhs260
				var _rhs261 _dafny.CodePoint = _1591_v4
				_ = _rhs261
				var _rhs262 D16 = _1611_v18
				_ = _rhs262
				var _lhs240 *GlobalState = globalState
				_ = _lhs240
				var _lhs241 _dafny.Array = _1602_v13
				_ = _lhs241
				var _lhs242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1602_v13), 0))
				_ = _lhs242
				_lhs240.F8 = _rhs259
				(_lhs241).ArraySet1(_rhs260, (_lhs242).Int())
				_1591_v4 = _rhs261
				_1611_v18 = _rhs262
			} else {
				var _1612_v19 _dafny.Array
				_ = _1612_v19
				var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
				_ = _nw254
				_1612_v19 = _nw254
				var _1613_v20 _dafny.Map
				_ = _1613_v20
				_1613_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1612_v19)
				(globalState).F22 = (func() _dafny.Array {
					if (_1613_v20).Contains(Companion_Default___.Fm0((_1590_v3).F28(), _1588_i0, p0, (_this).F24(), globalState)) {
						return (_1613_v20).Get(Companion_Default___.Fm0((_1590_v3).F28(), _1588_i0, p0, (_this).F24(), globalState)).(_dafny.Array)
					}
					return _1612_v19
				})()
				var _1614_v21 _dafny.Array
				_ = _1614_v21
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw255
				_1614_v21 = _nw255
				var _1615_v22 _dafny.Sequence
				_ = _1615_v22
				_1615_v22 = _dafny.SeqOf((_this).F24())
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1614_v21), 0))
				_ = _index263
				(_1614_v21).ArraySet1(_1615_v22, (_index263).Int())
				var _1616_v23 _dafny.Array
				_ = _1616_v23
				var _nw256 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw256
				_1616_v23 = _nw256
				var _1617_v24 _dafny.Sequence
				_ = _1617_v24
				_1617_v24 = _dafny.SeqOf(_1616_v23, _1616_v23, _1616_v23, _1616_v23, _1616_v23)
				var _1618_v25 _dafny.Set
				_ = _1618_v25
				_1618_v25 = _dafny.SetOf(_1588_i0)
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1614_v21), 0))
				_ = _index264
				var _rhs263 _dafny.Array = (_1617_v24).Select((Companion_Default___.SafeIndex(_1588_i0, _dafny.IntOfUint32((_1617_v24).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs263
				var _rhs264 bool = (_1618_v25).IsDisjointFrom((Companion_D14_.Create_DC28_(_1618_v25)).Dtor_cf37())
				_ = _rhs264
				var _rhs265 _dafny.Int = p0
				_ = _rhs265
				var _rhs266 _dafny.Sequence = _1615_v22
				_ = _rhs266
				var _lhs243 *GlobalState = globalState
				_ = _lhs243
				var _lhs244 *GlobalState = globalState
				_ = _lhs244
				var _lhs245 _dafny.Array = _1614_v21
				_ = _lhs245
				var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1614_v21), 0))
				_ = _lhs246
				r0 = _rhs263
				_lhs243.F13 = _rhs264
				_lhs244.F7 = _rhs265
				(_lhs245).ArraySet1(_rhs266, (_lhs246).Int())
				var _1619_v26 _dafny.Map
				_ = _1619_v26
				_1619_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1590_v3).F28()).Times(_dafny.IntOfInt64(899)), !(((_this).F24()) || ((_this).F24())))
				_1619_v26 = (func() _dafny.Map {
					if !((_this).F24()) || ((_this).F24()) {
						return func() _dafny.Map {
							var _coll63 = _dafny.NewMapBuilder()
							_ = _coll63
							for _iter75 := _dafny.Iterate((func() _dafny.Set {
								var _coll64 = _dafny.NewBuilder()
								_ = _coll64
								for _iter76 := _dafny.Iterate((_dafny.MultiSetOf((_1590_v3).F28(), _1588_i0)).Elements()); ; {
									_compr_64, _ok76 := _iter76()
									if !_ok76 {
										break
									}
									var _1620_v28 _dafny.Int
									_1620_v28 = interface{}(_compr_64).(_dafny.Int)
									if (_dafny.MultiSetOf((_1590_v3).F28(), _1588_i0)).Contains(_1620_v28) {
										_coll64.Add((_1620_v28).Minus(_dafny.IntOfInt64(227)))
									}
								}
								return _coll64.ToSet()
							}()).Elements()); ; {
								_compr_63, _ok75 := _iter75()
								if !_ok75 {
									break
								}
								var _1621_v27 _dafny.Int
								_1621_v27 = interface{}(_compr_63).(_dafny.Int)
								if (func() _dafny.Set {
									var _coll65 = _dafny.NewBuilder()
									_ = _coll65
									for _iter77 := _dafny.Iterate((_dafny.MultiSetOf((_1590_v3).F28(), _1588_i0)).Elements()); ; {
										_compr_65, _ok77 := _iter77()
										if !_ok77 {
											break
										}
										var _1622_v28 _dafny.Int
										_1622_v28 = interface{}(_compr_65).(_dafny.Int)
										if (_dafny.MultiSetOf((_1590_v3).F28(), _1588_i0)).Contains(_1622_v28) {
											_coll65.Add((_1622_v28).Minus(_dafny.IntOfInt64(227)))
										}
									}
									return _coll65.ToSet()
								}()).Contains(_1621_v27) {
									_coll63.Add((_1621_v27).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1590_v3).F28(), _1588_i0)).Cardinality()), (_this).F24())
								}
							}
							return _coll63.ToMap()
						}()
					}
					return (_1619_v26).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1590_v3).F28(), (_this).F24()))
				})()
				var _1623_v29 _dafny.Sequence
				_ = _1623_v29
				_1623_v29 = _dafny.UnicodeSeqOfUtf8Bytes("ccmntf")
				var _1624_v30 _dafny.Map
				_ = _1624_v30
				_1624_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1614_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1614_v21), 0))).Int()).(_dafny.Sequence), _1623_v29)
				var _1625_v31 _dafny.MultiSet
				_ = _1625_v31
				_1625_v31 = _dafny.MultiSetOf(p0, (_1590_v3).F28())
				Companion_Default___.M0(((_1614_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1614_v21), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_1590_v3).F28(), _dafny.IntOfUint32(((_1614_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1614_v21), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(bool), _1588_i0, _1624_v30, ((_1625_v31).Difference(_1625_v31)).Cardinality(), globalState)
				(globalState).F16 = _1588_i0
			}
		}
		var _1626_v32 _dafny.CodePoint
		_ = _1626_v32
		_1626_v32 = _dafny.CodePoint('c')
		_1626_v32 = _1626_v32
		(globalState).F0 = !(!((_this).F24()) || ((_this).F24()))
		var _1627_v33 _dafny.Map
		_ = _1627_v33
		_1627_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
		var _rhs267 _dafny.Int = p0
		_ = _rhs267
		var _rhs268 _dafny.Int = p0
		_ = _rhs268
		var _rhs269 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, ((_1627_v33).Merge(_1627_v33)).Cardinality())
		_ = _rhs269
		var _lhs247 *GlobalState = globalState
		_ = _lhs247
		var _lhs248 *GlobalState = globalState
		_ = _lhs248
		var _lhs249 *GlobalState = globalState
		_ = _lhs249
		_lhs247.F16 = _rhs267
		_lhs248.F16 = _rhs268
		_lhs249.F16 = _rhs269
		var _1628_v34 _dafny.Array
		_ = _1628_v34
		var _len0_48 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_48
		var _nw257 _dafny.Array
		_ = _nw257
		if _len0_48.Cmp(_dafny.Zero) == 0 {
			_nw257 = _dafny.NewArray(_len0_48)
		} else {
			var _init48 func(_dafny.Int) _dafny.Int = (func(_1629_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1630_i5 _dafny.Int) _dafny.Int {
					return (_1630_i5).Minus(_1629_p0)
				}
			})(p0)
			_ = _init48
			var _element0_48 = _init48(_dafny.Zero)
			_ = _element0_48
			_nw257 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
			(_nw257).ArraySet1(_element0_48, 0)
			var _nativeLen0_48 = (_len0_48).Int()
			_ = _nativeLen0_48
			for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
				(_nw257).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
			}
		}
		_1628_v34 = _nw257
		r0 = _1628_v34
		if (_this).F24() {
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))
			_ = _index265
			(_1628_v34).ArraySet1(p0, (_index265).Int())
			var _1631_v35 _dafny.Array
			_ = _1631_v35
			var _nw258 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw258
			_1631_v35 = _nw258
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))
			_ = _index266
			(_1631_v35).ArraySet1((_this).F24(), (_index266).Int())
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))
			_ = _index267
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))
			_ = _index268
			var _rhs270 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-510))
			_ = _rhs270
			var _rhs271 bool = !((_this).F24())
			_ = _rhs271
			var _rhs272 _dafny.Int = (p0).Minus(p0)
			_ = _rhs272
			var _rhs273 _dafny.Int = (p0).Times((p0).Minus(p0))
			_ = _rhs273
			var _lhs250 _dafny.Array = _1628_v34
			_ = _lhs250
			var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))
			_ = _lhs251
			var _lhs252 _dafny.Array = _1631_v35
			_ = _lhs252
			var _lhs253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))
			_ = _lhs253
			var _lhs254 *GlobalState = globalState
			_ = _lhs254
			var _lhs255 *GlobalState = globalState
			_ = _lhs255
			(_lhs250).ArraySet1(_rhs270, (_lhs251).Int())
			(_lhs252).ArraySet1(_rhs271, (_lhs253).Int())
			_lhs254.F16 = _rhs272
			_lhs255.F16 = _rhs273
			var _1632_v36 _dafny.Map
			_ = _1632_v36
			_1632_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F24()), _dafny.IntOfInt64(-656))
			(globalState).F7 = (func() _dafny.Int {
				if (_1632_v36).Contains(((_this).F24()) == ((_1631_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))).Int()).(bool))) {
					return (_1632_v36).Get(((_this).F24()) == ((_1631_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))).Int()).(bool))).(_dafny.Int)
				}
				return p0
			})()
			var _1633_v37 _dafny.Set
			_ = _1633_v37
			_1633_v37 = _dafny.SetOf(_dafny.IntOfInt64(533))
			var _1634_v38 _dafny.Sequence
			_ = _1634_v38
			_1634_v38 = _dafny.UnicodeSeqOfUtf8Bytes("ospctjgb")
			var _1635_v39 _dafny.Array
			_ = _1635_v39
			var _nwElement0_45 _dafny.Int = (_dafny.Zero).Minus((_1633_v37).Cardinality())
			_ = _nwElement0_45
			var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(19))
			_ = _nw259
			(_nw259).ArraySet1(_nwElement0_45, 0)
			(_nw259).ArraySet1(_dafny.IntOfUint32((_1634_v38).Cardinality()), 1)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 2)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 3)
			(_nw259).ArraySet1((_dafny.Zero).Minus(p0), 4)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 5)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 6)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 7)
			(_nw259).ArraySet1((_dafny.Zero).Minus((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int)), 8)
			(_nw259).ArraySet1(_dafny.IntOfUint32((_1634_v38).Cardinality()), 9)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 10)
			(_nw259).ArraySet1((_this).Fm8((_this).F24(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(692), _1626_v32)).Update(_dafny.IntOfInt64(-65), _1626_v32), _1634_v38, globalState), 11)
			(_nw259).ArraySet1(p0, 12)
			(_nw259).ArraySet1(_dafny.IntOfInt64(-377), 13)
			(_nw259).ArraySet1(_dafny.IntOfInt64(143), 14)
			(_nw259).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1631_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))).Int()).(bool)), (_this).F24())).Cardinality(), 15)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 16)
			(_nw259).ArraySet1((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), 17)
			(_nw259).ArraySet1(p0, 18)
			_1635_v39 = _nw259
			var _1636_v40 D6
			_ = _1636_v40
			_1636_v40 = Companion_D6_.Create_DC12_(_dafny.SeqOf(_1635_v39))
			var _1637_v42 _dafny.Map
			_ = _1637_v42
			_1637_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1626_v32, (_1631_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))).Int()).(bool))
			var _1638_v43 _dafny.Map
			_ = _1638_v43
			_1638_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), _1637_v42)
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))
			_ = _index269
			var _rhs274 bool = (_1631_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))).Int()).(bool)
			_ = _rhs274
			var _rhs275 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(231))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}((func(_1639_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1640_i6 _dafny.Int) _dafny.CodePoint {
					return _1639_v32
				}
			})(_1626_v32)))
			_ = _rhs275
			var _rhs276 bool = !(!((func() _dafny.Map {
				var _coll66 = _dafny.NewMapBuilder()
				_ = _coll66
				for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-222), _dafny.IntOfInt64(403))); ; {
					_compr_66, _ok78 := _iter78()
					if !_ok78 {
						break
					}
					var _1641_v41 _dafny.Int
					_1641_v41 = interface{}(_compr_66).(_dafny.Int)
					if ((_dafny.IntOfInt64(-222)).Cmp(_1641_v41) <= 0) && ((_1641_v41).Cmp(_dafny.IntOfInt64(403)) < 0) {
						_coll66.Add(Companion_Default___.SafeDivisionInt(_1641_v41, (_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1626_v32, (_1631_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))).Int()).(bool)))
					}
				}
				return _coll66.ToMap()
			}()).Merge(_1638_v43)).Equals(Companion_Default___.Fm48(_1626_v32, (_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int), globalState)))
			_ = _rhs276
			var _rhs277 D6 = _1636_v40
			_ = _rhs277
			var _rhs278 bool = !((_this).F24()) || ((_1631_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))).Int()).(bool))
			_ = _rhs278
			var _lhs256 _dafny.Array = _1631_v35
			_ = _lhs256
			var _lhs257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1631_v35), 0))
			_ = _lhs257
			var _lhs258 *GlobalState = globalState
			_ = _lhs258
			var _lhs259 *GlobalState = globalState
			_ = _lhs259
			var _lhs260 *GlobalState = globalState
			_ = _lhs260
			(_lhs256).ArraySet1(_rhs274, (_lhs257).Int())
			_lhs258.F20 = _rhs275
			_lhs259.F0 = _rhs276
			_1636_v40 = _rhs277
			_lhs260.F13 = _rhs278
			(globalState).F13 = !(true) || ((p0).Cmp((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int)) <= 0)
			var _1642_v44 *C2
			_ = _1642_v44
			var _nw260 *C2 = New_C2_()
			_ = _nw260
			_nw260.Ctor__((_1628_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1628_v34), 0))).Int()).(_dafny.Int))
			_1642_v44 = _nw260
			_1642_v44 = _1642_v44
		} else {
			var _1643_v45 _dafny.Map
			_ = _1643_v45
			_1643_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm7((_this).F24(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_this).F24())).Cardinality(), (_this).F24(), globalState), (_this).F24())
			var _1644_v46 _dafny.Set
			_ = _1644_v46
			_1644_v46 = _dafny.SetOf(Companion_Default___.Fm2(_1643_v45, globalState))
			var _1645_v47 _dafny.Map
			_ = _1645_v47
			_1645_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1644_v46, Companion_Default___.Fm18((_this).F24(), _1626_v32, (_this).Fm6(p0, globalState), globalState))
			var _1646_v48 _dafny.Sequence
			_ = _1646_v48
			_1646_v48 = _dafny.SeqOf(true)
			_1645_v47 = (_1645_v47).Update(Companion_Default___.Fm14(p0, (_this).F24(), globalState), _1646_v48)
			var _1647_v49 _dafny.Array
			_ = _1647_v49
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw261
			_1647_v49 = _nw261
			_1647_v49 = _1647_v49
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_1628_v34), 0))
			_ = _index270
			(_1628_v34).ArraySet1(p0, (_index270).Int())
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_1628_v34), 0))
			_ = _index271
			(_1628_v34).ArraySet1((((_dafny.SetOf((_this).F24(), true)).Difference(_1644_v46)).Difference(_1644_v46)).Cardinality(), (_index271).Int())
			_1646_v48 = _1646_v48
			var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw262
			r0 = _nw262
		}
		r0 = _1628_v34
		r1 = (_this).F24()
		return r0, r1
	}
}
func (_this *C5) F24() bool {
	{
		return _this._f24
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	dummy byte
}

func New_C6_() *C6 {
	_this := C6{}

	return &_this
}

type CompanionStruct_C6_ struct {
}

var Companion_C6_ = CompanionStruct_C6_{}

func (_this *C6) Equals(other *C6) bool {
	return _this == other
}

func (_this *C6) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C6)
	return ok && _this.Equals(other)
}

func (*C6) String() string {
	return "_module.C6"
}

func Type_C6_() _dafny.TypeDescriptor {
	return type_C6_{}
}

type type_C6_ struct {
}

func (_this type_C6_) Default() interface{} {
	return (*C6)(nil)
}

func (_this type_C6_) String() string {
	return "main.C6"
}
func (_this *C6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__() {
	{
	}
}
func (_this *C6) Fm7(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C6) Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			if false {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("xdbnkvr"))
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(585))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg103 _dafny.Int) interface{} {
					return coer103(arg103)
				}
			}(func(_1648_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			})))
		})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("bgtt")))
	}
}
func (_this *C6) Fm6(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C6) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		(globalState).F13 = p1
		(globalState).F16 = _dafny.IntOfInt64(-429)
		var _hi11 _dafny.Int = p3
		_ = _hi11
		for _1649_i0 := p3; _1649_i0.Cmp(_hi11) < 0; _1649_i0 = _1649_i0.Plus(_dafny.One) {
			var _1650_v0 _dafny.MultiSet
			_ = _1650_v0
			_1650_v0 = _dafny.MultiSetOf(_dafny.IntOfInt64(307))
			var _1651_v1 _dafny.MultiSet
			_ = _1651_v1
			_1651_v1 = _dafny.MultiSetOf(_1650_v0)
			(globalState).F0 = ((_dafny.MultiSetOf(_1650_v0, _1650_v0, _1650_v0, (_1650_v0).Update(_1649_i0, Companion_Default___.Abs(_dafny.IntOfInt64(-346))))).IsProperSubsetOf((_1651_v1).Update(_1650_v0, Companion_Default___.Abs(p0)))) == (p1)
			var _1652_v2 _dafny.Sequence
			_ = _1652_v2
			_1652_v2 = _dafny.SeqOf(!(p1), p1)
			var _1653_v3 *C0
			_ = _1653_v3
			var _nw263 *C0 = New_C0_()
			_ = _nw263
			_nw263.Ctor__((_1652_v2).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1652_v2).Cardinality()))).Uint32()).(bool))
			_1653_v3 = _nw263
			var _1654_v4 _dafny.Sequence
			_ = _1654_v4
			_1654_v4 = _dafny.SeqOf(p3)
			if (_dafny.IntOfUint32((_1654_v4).Cardinality())).Cmp((func() _dafny.Int {
				if p1 {
					return _1649_i0
				}
				return _1649_i0
			})()) >= 0 {
				(globalState).F7 = p0
				var _1655_v5 _dafny.Array
				_ = _1655_v5
				var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
				_ = _nw264
				_1655_v5 = _nw264
				var _1656_v6 _dafny.Sequence
				_ = _1656_v6
				_1656_v6 = _dafny.UnicodeSeqOfUtf8Bytes("odwrbay")
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1655_v5), 0))
				_ = _index272
				(_1655_v5).ArraySet1(_1656_v6, (_index272).Int())
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1655_v5), 0))
				_ = _index273
				(_1655_v5).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfInt64(151), globalState), (_index273).Int())
				(globalState).F16 = _1649_i0
				var _1657_v7 _dafny.Array
				_ = _1657_v7
				var _nw265 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw265
				_1657_v7 = _nw265
				_1657_v7 = _1657_v7
				var _1658_v8 _dafny.Sequence
				_ = _1658_v8
				_1658_v8 = _dafny.SeqOf((_1655_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1655_v5), 0))).Int()).(_dafny.Sequence), (_1655_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1655_v5), 0))).Int()).(_dafny.Sequence), (_1655_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1655_v5), 0))).Int()).(_dafny.Sequence), (_1655_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1655_v5), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm1(_1649_i0, globalState))
				var _1659_v9 _dafny.MultiSet
				_ = _1659_v9
				_1659_v9 = _dafny.MultiSetOf(_1653_v3.F27)
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1655_v5), 0))
				_ = _index274
				(_1655_v5).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1658_v8).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1654_v4).Select((Companion_Default___.SafeIndex((_1659_v9).Cardinality(), _dafny.IntOfUint32((_1654_v4).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfUint32((_1658_v8).Cardinality()))).Uint32()).(_dafny.Sequence), _1656_v6), (_index274).Int())
			} else {
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p2), 0))
				_ = _index275
				(p2).ArraySet1((_1654_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1652_v2).Cardinality()), _dafny.IntOfUint32((_1654_v4).Cardinality()))).Uint32()).(_dafny.Int), (_index275).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p2), 0))
				_ = _index276
				(p2).ArraySet1(p3, (_index276).Int())
				(_1653_v3).F27 = _1653_v3.F27
				var _1660_v10 _dafny.Set
				_ = _1660_v10
				_1660_v10 = _dafny.SetOf(_1653_v3.F27, (_this).Fm6(p3, globalState))
				var _1661_v11 _dafny.Array
				_ = _1661_v11
				var _nwElement0_46 _dafny.Set = _1660_v10
				_ = _nwElement0_46
				var _nw266 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(15))
				_ = _nw266
				(_nw266).ArraySet1(_nwElement0_46, 0)
				(_nw266).ArraySet1(_1660_v10, 1)
				(_nw266).ArraySet1(_1660_v10, 2)
				(_nw266).ArraySet1(_1660_v10, 3)
				(_nw266).ArraySet1(_1660_v10, 4)
				(_nw266).ArraySet1(_1660_v10, 5)
				(_nw266).ArraySet1(_1660_v10, 6)
				(_nw266).ArraySet1(_dafny.SetOf(_1653_v3.F27, _1653_v3.F27, _1653_v3.F27), 7)
				(_nw266).ArraySet1(_1660_v10, 8)
				(_nw266).ArraySet1(_1660_v10, 9)
				(_nw266).ArraySet1(_1660_v10, 10)
				(_nw266).ArraySet1(_dafny.SetOf(_1653_v3.F27), 11)
				(_nw266).ArraySet1(_1660_v10, 12)
				(_nw266).ArraySet1(_1660_v10, 13)
				(_nw266).ArraySet1(_1660_v10, 14)
				_1661_v11 = _nw266
				var _1662_v12 D7
				_ = _1662_v12
				_1662_v12 = Companion_D7_.Create_DC16_((func() _dafny.Array {
					if _1653_v3.F27 {
						return _1661_v11
					}
					return _1661_v11
				})())
				_1662_v12 = _1662_v12
				var _1663_v13 *C4
				_ = _1663_v13
				var _nw267 *C4 = New_C4_()
				_ = _nw267
				_nw267.Ctor__(p1, _1653_v3.F27)
				_1663_v13 = _nw267
				var _1664_v14 _dafny.Map
				_ = _1664_v14
				_1664_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1653_v3.F27, _1663_v13)
				var _1665_v15 _dafny.Map
				_ = _1665_v15
				_1665_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C4 {
					if (_1664_v14).Contains(_1663_v13.F25) {
						return (_1664_v14).Get(_1663_v13.F25).(*C4)
					}
					return _1663_v13
				})(), p0)
				_1665_v15 = (_1665_v15).Update(_1663_v13, (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(_dafny.IntOfInt64(511), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), p1, globalState))))
				var _1666_v16 _dafny.Map
				_ = _1666_v16
				_1666_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), !(_1653_v3.F27))
				var _1667_v17 _dafny.Array
				_ = _1667_v17
				var _nw268 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
				_ = _nw268
				_1667_v17 = _nw268
				var _1668_v18 _dafny.Map
				_ = _1668_v18
				_1668_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1666_v16, _1667_v17)
				_1668_v18 = (_1668_v18).Update(_1666_v16, _1667_v17)
			}
			var _1669_v19 _dafny.Map
			_ = _1669_v19
			_1669_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1649_i0, _1653_v3.F27)
			var _1670_v20 _dafny.Map
			_ = _1670_v20
			_1670_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1669_v19, p2)
			_1670_v20 = ((_1670_v20).Merge((_1670_v20).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1649_i0, p1), p2))).Merge(_1670_v20)
		}
		var _1671_v21 _dafny.Map
		_ = _1671_v21
		_1671_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
		(globalState).F2 = (Companion_D19_.Create_DC47_(Companion_D1_.Create_DC2_(_1671_v21), p3, p1)).Dtor_cf64()
		(globalState).F2 = (p1) == (p1)
		var _1672_v22 _dafny.Array
		_ = _1672_v22
		var _nwElement0_47 bool = (func() bool {
			if p1 {
				return false
			}
			return !(p1)
		})()
		_ = _nwElement0_47
		var _nw269 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(3))
		_ = _nw269
		(_nw269).ArraySet1(_nwElement0_47, 0)
		(_nw269).ArraySet1(p1, 1)
		(_nw269).ArraySet1(p1, 2)
		_1672_v22 = _nw269
		var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1672_v22), 0))
		_ = _index277
		(_1672_v22).ArraySet1(!(!(p1)) || (true), (_index277).Int())
		var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1672_v22), 0))
		_ = _index278
		(_1672_v22).ArraySet1(p1, (_index278).Int())
		var _1673_v23 _dafny.CodePoint
		_ = _1673_v23
		_1673_v23 = _dafny.CodePoint('e')
		var _1674_v24 _dafny.Sequence
		_ = _1674_v24
		_1674_v24 = _dafny.SeqOf(p0, _dafny.IntOfInt64(325))
		var _1675_v25 _dafny.Map
		_ = _1675_v25
		_1675_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1673_v23, p1)
		var _1676_v26 _dafny.MultiSet
		_ = _1676_v26
		_1676_v26 = _dafny.MultiSetOf((Companion_Default___.Fm49(_1673_v23, _1674_v24, _dafny.IntOfInt64(813), _1673_v23, globalState)).Dtor_cf60(), (func() bool {
			if (_1675_v25).Contains(_1673_v23) {
				return (_1675_v25).Get(_1673_v23).(bool)
			}
			return p1
		})(), p1)
		var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((p2), 0))
		_ = _index279
		(p2).ArraySet1((_1676_v26).Cardinality(), (_index279).Int())
		var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1672_v22), 0))
		_ = _index280
		var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1672_v22), 0))
		_ = _index281
		var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((p2), 0))
		_ = _index282
		var _rhs279 bool = p1
		_ = _rhs279
		var _rhs280 bool = (p0).Cmp(p0) < 0
		_ = _rhs280
		var _rhs281 _dafny.MultiSet = _1676_v26
		_ = _rhs281
		var _rhs282 _dafny.Int = p3
		_ = _rhs282
		var _lhs261 _dafny.Array = _1672_v22
		_ = _lhs261
		var _lhs262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1672_v22), 0))
		_ = _lhs262
		var _lhs263 _dafny.Array = _1672_v22
		_ = _lhs263
		var _lhs264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1672_v22), 0))
		_ = _lhs264
		var _lhs265 _dafny.Array = p2
		_ = _lhs265
		var _lhs266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((p2), 0))
		_ = _lhs266
		(_lhs261).ArraySet1(_rhs279, (_lhs262).Int())
		(_lhs263).ArraySet1(_rhs280, (_lhs264).Int())
		_1676_v26 = _rhs281
		(_lhs265).ArraySet1(_rhs282, (_lhs266).Int())
		r0 = ((p1) == (p1)) || (p1)
		r1 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_1672_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1672_v22), 0))).Int()).(bool) {
				return p0
			}
			return (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)
		})(), _dafny.IntOfInt64(133))
		var _1677_v27 _dafny.Array
		_ = _1677_v27
		var _len0_49 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_49
		var _nw270 _dafny.Array
		_ = _nw270
		if _len0_49.Cmp(_dafny.Zero) == 0 {
			_nw270 = _dafny.NewArray(_len0_49)
		} else {
			var _init49 func(_dafny.Int) _dafny.Map = (func(_1678_v22 _dafny.Array, _1679_v23 _dafny.CodePoint, _1680_p1 bool) func(_dafny.Int) _dafny.Map {
				return func(_1681_i1 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1678_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1678_v22), 0))).Int()).(bool), Companion_D3_.Create_DC6_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(490))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg104 _dafny.Int) interface{} {
							return coer104(arg104)
						}
					}((func(_1682_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1683_i2 _dafny.Int) _dafny.CodePoint {
							return _1682_v23
						}
					})(_1679_v23))), _dafny.CodePoint('b'), _1680_p1))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1680_p1, Companion_D3_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("whrnsnix"), _1679_v23, (_1678_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_1678_v22), 0))).Int()).(bool))))
				}
			})(_1672_v22, _1673_v23, p1)
			_ = _init49
			var _element0_49 = _init49(_dafny.Zero)
			_ = _element0_49
			_nw270 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
			(_nw270).ArraySet1(_element0_49, 0)
			var _nativeLen0_49 = (_len0_49).Int()
			_ = _nativeLen0_49
			for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
				(_nw270).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
			}
		}
		_1677_v27 = _nw270
		r2 = _1677_v27
		return r0, r1, r2
	}
}
func (_this *C6) M3(p0 _dafny.Array, globalState *GlobalState) {
	{
		var _1684_v0 _dafny.Int
		_ = _1684_v0
		_1684_v0 = _dafny.IntOfInt64(685)
		var _1685_v1 *C5
		_ = _1685_v1
		var _nw271 *C5 = New_C5_()
		_ = _nw271
		_nw271.Ctor__((_1684_v0).Cmp(_1684_v0) <= 0)
		_1685_v1 = _nw271
		var _1686_v2 *C3
		_ = _1686_v2
		var _nw272 *C3 = New_C3_()
		_ = _nw272
		_nw272.Ctor__()
		_1686_v2 = _nw272
		(globalState).F0 = true
		var _1687_v3 _dafny.Map
		_ = _1687_v3
		_1687_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), false)
		var _1688_v4 _dafny.Map
		_ = _1688_v4
		_1688_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_1687_v3, globalState), Companion_Default___.Fm0(_1684_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("brysun")).Cardinality()), _1684_v0, (_1685_v1).F24(), globalState))
		var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))
		_ = _index283
		(p0).ArraySet1((_1688_v4).Cardinality(), (_index283).Int())
		var _1689_v5 D19
		_ = _1689_v5
		_1689_v5 = Companion_D19_.Create_DC48_(_1684_v0, (_1685_v1).F24())
		var _1690_v6 _dafny.Sequence
		_ = _1690_v6
		_1690_v6 = _dafny.SeqOf((_1685_v1).F24(), (_1685_v1).F24())
		var _1691_v7 _dafny.MultiSet
		_ = _1691_v7
		_1691_v7 = _dafny.MultiSetOf(_1687_v3)
		var _1692_v8 _dafny.Map
		_ = _1692_v8
		_1692_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1691_v7).Cardinality(), _1684_v0)
		var _1693_v9 _dafny.Set
		_ = _1693_v9
		_1693_v9 = _dafny.SetOf(_1692_v8)
		var _1694_v10 _dafny.Map
		_ = _1694_v10
		_1694_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1693_v9).Cardinality(), (_1685_v1).F24())
		var _1695_v11 _dafny.Array
		_ = _1695_v11
		var _nwElement0_48 bool = (func() bool {
			if !((_1685_v1).F24()) {
				return !((_1685_v1).F24())
			}
			return (_1685_v1).F24()
		})()
		_ = _nwElement0_48
		var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(29))
		_ = _nw273
		(_nw273).ArraySet1(_nwElement0_48, 0)
		(_nw273).ArraySet1((_1689_v5).Dtor_cf66(), 1)
		(_nw273).ArraySet1((_1685_v1).F24(), 2)
		(_nw273).ArraySet1((_1685_v1).F24(), 3)
		(_nw273).ArraySet1((_1685_v1).F24(), 4)
		(_nw273).ArraySet1((_1685_v1).F24(), 5)
		(_nw273).ArraySet1(!((_1685_v1).F24()), 6)
		(_nw273).ArraySet1((_1685_v1).F24(), 7)
		(_nw273).ArraySet1((func() bool {
			if (func() bool {
				if (_1687_v3).Contains((_1685_v1).F24()) {
					return (_1687_v3).Get((_1685_v1).F24()).(bool)
				}
				return (_1685_v1).F24()
			})() {
				return (_1685_v1).F24()
			}
			return false
		})(), 8)
		(_nw273).ArraySet1((_1685_v1).F24(), 9)
		(_nw273).ArraySet1((_1685_v1).F24(), 10)
		(_nw273).ArraySet1((_1685_v1).F24(), 11)
		(_nw273).ArraySet1((_1685_v1).F24(), 12)
		(_nw273).ArraySet1((_1685_v1).F24(), 13)
		(_nw273).ArraySet1((_1685_v1).F24(), 14)
		(_nw273).ArraySet1((_1685_v1).F24(), 15)
		(_nw273).ArraySet1(false, 16)
		(_nw273).ArraySet1(!((_1685_v1).F24()), 17)
		(_nw273).ArraySet1(!(_dafny.Companion_Sequence_.IsPrefixOf(_1690_v6, _1690_v6)), 18)
		(_nw273).ArraySet1((_1685_v1).F24(), 19)
		(_nw273).ArraySet1((_1685_v1).F24(), 20)
		(_nw273).ArraySet1((_1685_v1).F24(), 21)
		(_nw273).ArraySet1((_1685_v1).F24(), 22)
		(_nw273).ArraySet1((func() bool {
			if false {
				return (_1685_v1).F24()
			}
			return (_1685_v1).F24()
		})(), 23)
		(_nw273).ArraySet1((!((_1685_v1).F24())) || ((_1685_v1).F24()), 24)
		(_nw273).ArraySet1((func() bool {
			if (_1694_v10).Contains(_1684_v0) {
				return (_1694_v10).Get(_1684_v0).(bool)
			}
			return false
		})(), 25)
		(_nw273).ArraySet1(!(false), 26)
		(_nw273).ArraySet1(false, 27)
		(_nw273).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf((_1685_v1).F24())).Cardinality())).Cmp(_1684_v0) <= 0, 28)
		_1695_v11 = _nw273
		var _1696_v12 _dafny.CodePoint
		_ = _1696_v12
		_1696_v12 = _dafny.CodePoint('n')
		var _1697_v13 D20
		_ = _1697_v13
		_1697_v13 = Companion_D20_.Create_DC49_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1684_v0, _1696_v12))
		var _1698_v14 _dafny.Map
		_ = _1698_v14
		_1698_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1685_v1).F24()), (_1697_v13).Dtor_cf67())
		var _1699_v15 _dafny.Sequence
		_ = _1699_v15
		_1699_v15 = _dafny.UnicodeSeqOfUtf8Bytes("snhrb")
		var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))
		_ = _index284
		var _rhs283 _dafny.Int = Companion_Default___.SafeDivisionInt((_1684_v0).Times(_dafny.IntOfInt64(229)), (_dafny.Zero).Minus((_1685_v1).Fm8((_1690_v6).Select((Companion_Default___.SafeIndex(_1684_v0, _dafny.IntOfUint32((_1690_v6).Cardinality()))).Uint32()).(bool), (func() _dafny.Map {
			if (_1698_v14).Contains((_1685_v1).F24()) {
				return (_1698_v14).Get((_1685_v1).F24()).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1684_v0, _1696_v12)
		})(), _1699_v15, globalState)))
		_ = _rhs283
		var _rhs284 bool = (func() bool {
			if (_1687_v3).Contains((_1685_v1).F24()) {
				return (_1687_v3).Get((_1685_v1).F24()).(bool)
			}
			return (_this).Fm6(_1684_v0, globalState)
		})()
		_ = _rhs284
		var _rhs285 _dafny.Int = (_1684_v0).Minus(_1684_v0)
		_ = _rhs285
		var _rhs286 _dafny.Array = _1695_v11
		_ = _rhs286
		var _lhs267 _dafny.Array = p0
		_ = _lhs267
		var _lhs268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))
		_ = _lhs268
		var _lhs269 *GlobalState = globalState
		_ = _lhs269
		var _lhs270 *GlobalState = globalState
		_ = _lhs270
		(_lhs267).ArraySet1(_rhs283, (_lhs268).Int())
		_lhs269.F2 = _rhs284
		_lhs270.F16 = _rhs285
		_1695_v11 = _rhs286
		var _1700_i0 _dafny.Int
		_ = _1700_i0
		_1700_i0 = _dafny.Zero
		{
			for (_1685_v1).F24() {
				{
					if (_1700_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1700_i0 = (_1700_i0).Plus(_dafny.One)
					(globalState).F16 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
						if (_1685_v1).F24() {
							return _1684_v0
						}
						return _1684_v0
					})(), _1684_v0)
					_1694_v10 = (_1694_v10).Update(_dafny.IntOfInt64(145), (_1685_v1).F24())
					var _1701_v16 _dafny.Array
					_ = _1701_v16
					var _len0_50 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_50
					var _nw274 _dafny.Array
					_ = _nw274
					if _len0_50.Cmp(_dafny.Zero) == 0 {
						_nw274 = _dafny.NewArray(_len0_50)
					} else {
						var _init50 func(_dafny.Int) D16 = (func(_1702_v1 *C5, _1703_v6 _dafny.Sequence) func(_dafny.Int) D16 {
							return func(_1704_i1 _dafny.Int) D16 {
								return Companion_D16_.Create_DC36_((_1702_v1).F24(), _dafny.MultiSetFromSeq(_1703_v6), (_1702_v1).F24())
							}
						})(_1685_v1, _1690_v6)
						_ = _init50
						var _element0_50 = _init50(_dafny.Zero)
						_ = _element0_50
						_nw274 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
						(_nw274).ArraySet1(_element0_50, 0)
						var _nativeLen0_50 = (_len0_50).Int()
						_ = _nativeLen0_50
						for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
							(_nw274).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
						}
					}
					_1701_v16 = _nw274
					var _1705_v17 D17
					_ = _1705_v17
					_1705_v17 = Companion_D17_.Create_DC38_(_1701_v16)
					var _1706_v18 _dafny.Map
					_ = _1706_v18
					_1706_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, _1688_v4)
					var _1707_v19 _dafny.Sequence
					_ = _1707_v19
					_1707_v19 = _dafny.SeqOf(_1706_v18, _1706_v18)
					var _1708_v20 _dafny.Map
					_ = _1708_v20
					_1708_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1684_v0, _1688_v4)
					var _1709_v21 _dafny.Map
					_ = _1709_v21
					_1709_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), Companion_Default___.Fm35((_1685_v1).F24(), _dafny.IntOfUint32((_1699_v15).Cardinality()), globalState))
					var _1710_v22 _dafny.MultiSet
					_ = _1710_v22
					_1710_v22 = _dafny.MultiSetOf(true, (_1685_v1).F24(), (_1685_v1).F24(), (_1685_v1).F24(), (_1685_v1).F24())
					var _1711_v23 _dafny.Array
					_ = _1711_v23
					var _nwElement0_49 _dafny.Map = (_1706_v18).Merge(_1706_v18)
					_ = _nwElement0_49
					var _nw275 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(29))
					_ = _nw275
					(_nw275).ArraySet1(_nwElement0_49, 0)
					(_nw275).ArraySet1(_1706_v18, 1)
					(_nw275).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, (_1688_v4).Update((_1685_v1).F24(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))), 2)
					(_nw275).ArraySet1(_1706_v18, 3)
					(_nw275).ArraySet1((_1706_v18).Update(_1705_v17, _1688_v4), 4)
					(_nw275).ArraySet1((_1706_v18).Merge(_1706_v18), 5)
					(_nw275).ArraySet1((_1707_v19).Select((Companion_Default___.SafeIndex(_1684_v0, _dafny.IntOfUint32((_1707_v19).Cardinality()))).Uint32()).(_dafny.Map), 6)
					(_nw275).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D17_.Create_DC38_(_1701_v16), _1688_v4), 7)
					(_nw275).ArraySet1(_1706_v18, 8)
					(_nw275).ArraySet1((_1706_v18).Merge(_1706_v18), 9)
					(_nw275).ArraySet1(_1706_v18, 10)
					(_nw275).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, _1688_v4), 11)
					(_nw275).ArraySet1(_1706_v18, 12)
					(_nw275).ArraySet1((func() _dafny.Map {
						if (_1685_v1).F24() {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, _1688_v4)
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, (func() _dafny.Map {
							if (_1708_v20).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)) {
								return (_1708_v20).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).(_dafny.Map)
							}
							return _1688_v4
						})())
					})(), 13)
					(_nw275).ArraySet1(_1706_v18, 14)
					(_nw275).ArraySet1(_1706_v18, 15)
					(_nw275).ArraySet1(_1706_v18, 16)
					(_nw275).ArraySet1(_1706_v18, 17)
					(_nw275).ArraySet1(_1706_v18, 18)
					(_nw275).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, _1688_v4)).Update(_1705_v17, (_1688_v4).Update((_1685_v1).F24(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)))).Merge(_1706_v18), 19)
					(_nw275).ArraySet1((_1706_v18).Update(Companion_D17_.Create_DC38_(_1701_v16), (func() _dafny.Map {
						if (_1709_v21).Contains((_1685_v1).F24()) {
							return (_1709_v21).Get((_1685_v1).F24()).(_dafny.Map)
						}
						return _1688_v4
					})()), 20)
					(_nw275).ArraySet1(_1706_v18, 21)
					(_nw275).ArraySet1(_1706_v18, 22)
					(_nw275).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), (_dafny.Zero).Minus((_1710_v22).Cardinality()))), 23)
					(_nw275).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v17, _1688_v4)).Merge(_1706_v18), 24)
					(_nw275).ArraySet1((_1706_v18).Update(_1705_v17, _1688_v4), 25)
					(_nw275).ArraySet1((_1706_v18).Merge(_1706_v18), 26)
					(_nw275).ArraySet1((_1706_v18).Update(Companion_D17_.Create_DC38_((_1705_v17).Dtor_cf51()), _1688_v4), 27)
					(_nw275).ArraySet1(_1706_v18, 28)
					_1711_v23 = _nw275
					var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1711_v23), 0))
					_ = _index285
					(_1711_v23).ArraySet1(_1706_v18, (_index285).Int())
					var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1711_v23), 0))
					_ = _index286
					(_1711_v23).ArraySet1((_1707_v19).Select((Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1707_v19).Cardinality()))).Uint32()).(_dafny.Map), (_index286).Int())
					(globalState).F19 = _1699_v15
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _hi12 _dafny.Int = _dafny.IntOfUint32((_1699_v15).Cardinality())
		_ = _hi12
		for _1712_i2 := ((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfInt64(211)); _1712_i2.Cmp(_hi12) < 0; _1712_i2 = _1712_i2.Plus(_dafny.One) {
			var _1713_v24 _dafny.Map
			_ = _1713_v24
			_1713_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(318))).Uint32(), func(coer105 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_1714_v15 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1715_i3 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_1714_v15).Cardinality())
				}
			})(_1699_v15))))
			var _1716_v25 _dafny.Map
			_ = _1716_v25
			_1716_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), _1713_v24)
			var _1717_v26 _dafny.Set
			_ = _1717_v26
			_1717_v26 = _dafny.SetOf((_1685_v1).F24())
			var _1718_v27 _dafny.Sequence
			_ = _1718_v27
			_1718_v27 = _dafny.SeqOf((_1688_v4).Cardinality(), (_1717_v26).Cardinality(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
			var _1719_v28 _dafny.Array
			_ = _1719_v28
			var _nwElement0_50 _dafny.Map = _1713_v24
			_ = _nwElement0_50
			var _nw276 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(14))
			_ = _nw276
			(_nw276).ArraySet1(_nwElement0_50, 0)
			(_nw276).ArraySet1(_1713_v24, 1)
			(_nw276).ArraySet1((func() _dafny.Map {
				if (_1716_v25).Contains((_1685_v1).F24()) {
					return (_1716_v25).Get((_1685_v1).F24()).(_dafny.Map)
				}
				return _1713_v24
			})(), 2)
			(_nw276).ArraySet1(_1713_v24, 3)
			(_nw276).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), _1718_v27), 4)
			(_nw276).ArraySet1((_1713_v24).Merge(_1713_v24), 5)
			(_nw276).ArraySet1((_1713_v24).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), _1718_v27)), 6)
			(_nw276).ArraySet1((_1713_v24).Merge(_1713_v24), 7)
			(_nw276).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), _1718_v27), 8)
			(_nw276).ArraySet1((Companion_Default___.Fm50(_1712_i2, _1712_i2, globalState)).Merge(_1713_v24), 9)
			(_nw276).ArraySet1(_1713_v24, 10)
			(_nw276).ArraySet1((_1713_v24).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1685_v1).F24(), _1718_v27)), 11)
			(_nw276).ArraySet1(_1713_v24, 12)
			(_nw276).ArraySet1(_1713_v24, 13)
			_1719_v28 = _nw276
			var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1719_v28), 0))
			_ = _index287
			(_1719_v28).ArraySet1(_1713_v24, (_index287).Int())
			var _1720_v29 _dafny.Map
			_ = _1720_v29
			_1720_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1699_v15, _1692_v8)
			var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1719_v28), 0))
			_ = _index288
			var _rhs287 _dafny.CodePoint = _1696_v12
			_ = _rhs287
			var _rhs288 bool = (_1685_v1).F24()
			_ = _rhs288
			var _rhs289 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1685_v1).F24()), _1718_v27)
			_ = _rhs289
			var _rhs290 bool = !((_1685_v1).F24())
			_ = _rhs290
			var _rhs291 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(((_1720_v29).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1699_v15, (Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1699_v15).Cardinality()))).Uint32(), _dafny.CodePoint('i')), _1692_v8)).Merge(_1720_v29))).Cardinality()))
			_ = _rhs291
			var _lhs271 *GlobalState = globalState
			_ = _lhs271
			var _lhs272 _dafny.Array = _1719_v28
			_ = _lhs272
			var _lhs273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1719_v28), 0))
			_ = _lhs273
			var _lhs274 *GlobalState = globalState
			_ = _lhs274
			var _lhs275 *GlobalState = globalState
			_ = _lhs275
			_1696_v12 = _rhs287
			_lhs271.F2 = _rhs288
			(_lhs272).ArraySet1(_rhs289, (_lhs273).Int())
			_lhs274.F21 = _rhs290
			_lhs275.F7 = _rhs291
			var _1721_v30 _dafny.Sequence
			_ = _1721_v30
			_1721_v30 = _dafny.SeqOf(_1694_v10)
			var _1722_v31 *C1
			_ = _1722_v31
			var _nw277 *C1 = New_C1_()
			_ = _nw277
			_nw277.Ctor__((func() _dafny.Int {
				if (_1685_v1).F24() {
					return _1712_i2
				}
				return _dafny.IntOfUint32((_1699_v15).Cardinality())
			})(), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm51(_dafny.CodePoint('i'), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), globalState), _1721_v30))
			_1722_v31 = _nw277
			var _1723_v32 _dafny.Map
			_ = _1723_v32
			_1723_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm34(globalState), (_1687_v3).Cardinality())
			_1723_v32 = (_1723_v32).Update(_1696_v12, _1722_v31.F29)
			var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1695_v11), 0))
			_ = _index289
			(_1695_v11).ArraySet1((_1685_v1).F24(), (_index289).Int())
			var _1724_v33 _dafny.Sequence
			_ = _1724_v33
			_1724_v33 = _dafny.SeqOf(p0)
			var _1725_v34 D6
			_ = _1725_v34
			_1725_v34 = Companion_D6_.Create_DC12_(_dafny.Companion_Sequence_.Update(_1724_v33, (Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1724_v33).Cardinality()))).Uint32(), p0))
			var _1726_v35 D6
			_ = _1726_v35
			_1726_v35 = Companion_D6_.Create_DC14_(_1725_v34)
			var _1727_v36 D6
			_ = _1727_v36
			_1727_v36 = Companion_D6_.Create_DC14_(_1725_v34)
			var _1728_v37 D6
			_ = _1728_v37
			_1728_v37 = Companion_D6_.Create_DC14_(_1726_v35)
			var _1729_v38 _dafny.Map
			_ = _1729_v38
			_1729_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1722_v31.F29, p0)
			var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1695_v11), 0))
			_ = _index290
			var _rhs292 _dafny.Int = ((_1722_v31).Fm29(globalState)).Plus(((_1694_v10).Cardinality()).Plus(_1712_i2))
			_ = _rhs292
			var _rhs293 bool = (_1729_v38).Contains((func() _dafny.Int {
				if (_1685_v1).F24() {
					return (_dafny.Zero).Minus(_1722_v31.F29)
				}
				return _1684_v0
			})())
			_ = _rhs293
			var _rhs294 D6 = Companion_D6_.Create_DC14_(_1725_v34)
			_ = _rhs294
			var _rhs295 _dafny.Sequence = _1699_v15
			_ = _rhs295
			var _rhs296 _dafny.Int = _1722_v31.F29
			_ = _rhs296
			var _lhs276 *C1 = _1722_v31
			_ = _lhs276
			var _lhs277 _dafny.Array = _1695_v11
			_ = _lhs277
			var _lhs278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1695_v11), 0))
			_ = _lhs278
			var _lhs279 *GlobalState = globalState
			_ = _lhs279
			var _lhs280 *GlobalState = globalState
			_ = _lhs280
			_lhs276.F29 = _rhs292
			(_lhs277).ArraySet1(_rhs293, (_lhs278).Int())
			_1728_v37 = _rhs294
			_lhs279.F20 = _rhs295
			_lhs280.F7 = _rhs296
		}
	}
}
func (_this *C6) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _1730_v0 bool
		_ = _1730_v0
		_1730_v0 = false
		var _1731_v1 _dafny.Map
		_ = _1731_v1
		_1731_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1730_v0, _1730_v0)
		_1731_v1 = (_1731_v1).Update(_1730_v0, Companion_Default___.Fm2((_1731_v1).Update(_1730_v0, !(_1730_v0)), globalState))
		var _1732_i0 _dafny.Int
		_ = _1732_i0
		_1732_i0 = _dafny.Zero
		{
			for _1730_v0 {
				{
					if (_1732_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1732_i0 = (_1732_i0).Plus(_dafny.One)
					var _1733_v2 _dafny.CodePoint
					_ = _1733_v2
					_1733_v2 = _dafny.CodePoint('t')
					var _rhs297 bool = _1730_v0
					_ = _rhs297
					var _rhs298 _dafny.CodePoint = _dafny.CodePoint('k')
					_ = _rhs298
					var _lhs281 *GlobalState = globalState
					_ = _lhs281
					_lhs281.F0 = _rhs297
					_1733_v2 = _rhs298
					if true {
						var _1734_v3 _dafny.MultiSet
						_ = _1734_v3
						_1734_v3 = _dafny.MultiSetOf((_1731_v1).Cardinality())
						(globalState).F2 = ((_1734_v3).IsDisjointFrom(_1734_v3)) || (_1730_v0)
						var _1735_v5 _dafny.Set
						_ = _1735_v5
						_1735_v5 = _dafny.SetOf((func() _dafny.Map {
							var _coll67 = _dafny.NewMapBuilder()
							_ = _coll67
							for _iter79 := _dafny.Iterate((_1734_v3).Elements()); ; {
								_compr_67, _ok79 := _iter79()
								if !_ok79 {
									break
								}
								var _1736_v4 _dafny.Int
								_1736_v4 = interface{}(_compr_67).(_dafny.Int)
								if (_1734_v3).Contains(_1736_v4) {
									_coll67.Add((_1736_v4).Plus(p0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(351))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
										return func(arg106 _dafny.Int) interface{} {
											return coer106(arg106)
										}
									}((func(_1737_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
										return func(_1738_i1 _dafny.Int) _dafny.Int {
											return _1737_p0
										}
									})(p0))))
								}
							}
							return _coll67.ToMap()
						}()).Cardinality(), p0)
						var _1739_v6 _dafny.MultiSet
						_ = _1739_v6
						_1739_v6 = _dafny.MultiSetOf(_1730_v0, (_1735_v5).IsSubsetOf(_1735_v5))
						var _1740_v7 _dafny.Map
						_ = _1740_v7
						_1740_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1739_v6).Update(_1730_v0, Companion_Default___.Abs(p0)))
						var _1741_v8 _dafny.Array
						_ = _1741_v8
						var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
						_ = _nw278
						_1741_v8 = _nw278
						var _1742_v9 _dafny.Sequence
						_ = _1742_v9
						_1742_v9 = _dafny.SeqOf(_1730_v0)
						var _1743_v10 _dafny.Map
						_ = _1743_v10
						_1743_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_v8, _1742_v9)
						_1739_v6 = ((func() _dafny.MultiSet {
							if (_1740_v7).Contains(p0) {
								return (_1740_v7).Get(p0).(_dafny.MultiSet)
							}
							return _1739_v6
						})()).Union(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
							if (_1743_v10).Contains(_1741_v8) {
								return (_1743_v10).Get(_1741_v8).(_dafny.Sequence)
							}
							return _1742_v9
						})()))
						var _1744_v11 *C0
						_ = _1744_v11
						var _nw279 *C0 = New_C0_()
						_ = _nw279
						_nw279.Ctor__(_1730_v0)
						_1744_v11 = _nw279
						_1744_v11 = _1744_v11
						var _1745_v12 _dafny.Array
						_ = _1745_v12
						var _nw280 _dafny.Array = _dafny.NewArray(_dafny.One)
						_ = _nw280
						_1745_v12 = _nw280
						var _1746_v13 _dafny.Sequence
						_ = _1746_v13
						_1746_v13 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eoog")).Cardinality()))
						var _1747_v14 _dafny.Sequence
						_ = _1747_v14
						_1747_v14 = _dafny.SeqOf(_1746_v13)
						var _1748_v15 _dafny.Map
						_ = _1748_v15
						_1748_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1742_v9, (_1747_v14).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1747_v14).Cardinality()))).Uint32()).(_dafny.Sequence))
						var _1749_v16 D21
						_ = _1749_v16
						_1749_v16 = Companion_D21_.Create_DC52_(_1748_v15)
						var _1750_v17 *C2
						_ = _1750_v17
						var _nw281 *C2 = New_C2_()
						_ = _nw281
						_nw281.Ctor__(((_1749_v16).Dtor_cf71()).Cardinality())
						_1750_v17 = _nw281
						var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_1745_v12), 0))
						_ = _index291
						(_1745_v12).ArraySet1(_1750_v17, (_index291).Int())
						var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_1745_v12), 0))
						_ = _index292
						(_1745_v12).ArraySet1(_1750_v17, (_index292).Int())
						(globalState).F19 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-484))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg107 _dafny.Int) interface{} {
								return coer107(arg107)
							}
						}((func(_1751_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1752_i2 _dafny.Int) _dafny.CodePoint {
								return _1751_v2
							}
						})(_1733_v2)))
					} else {
						var _1753_v18 _dafny.Array
						_ = _1753_v18
						var _nw282 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
						_ = _nw282
						_1753_v18 = _nw282
						r0 = _1753_v18
						var _1754_v19 _dafny.Set
						_ = _1754_v19
						_1754_v19 = _dafny.SetOf(p0, p0)
						var _1755_v20 D14
						_ = _1755_v20
						_1755_v20 = Companion_D14_.Create_DC28_(_1754_v19)
						var _1756_v21 _dafny.Array
						_ = _1756_v21
						var _nwElement0_51 bool = _1730_v0
						_ = _nwElement0_51
						var _nw283 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(25))
						_ = _nw283
						(_nw283).ArraySet1(_nwElement0_51, 0)
						(_nw283).ArraySet1(_1730_v0, 1)
						(_nw283).ArraySet1(!(_1730_v0), 2)
						(_nw283).ArraySet1(_1730_v0, 3)
						(_nw283).ArraySet1(!(_1730_v0), 4)
						(_nw283).ArraySet1((_1754_v19).IsProperSubsetOf((_1755_v20).Dtor_cf37()), 5)
						(_nw283).ArraySet1((p0).Cmp(p0) > 0, 6)
						(_nw283).ArraySet1(_1730_v0, 7)
						(_nw283).ArraySet1(_1730_v0, 8)
						(_nw283).ArraySet1((!(true)) == (false), 9)
						(_nw283).ArraySet1(true, 10)
						(_nw283).ArraySet1(!(true), 11)
						(_nw283).ArraySet1(_1730_v0, 12)
						(_nw283).ArraySet1(_1730_v0, 13)
						(_nw283).ArraySet1(!(_1730_v0) || (_1730_v0), 14)
						(_nw283).ArraySet1(_1730_v0, 15)
						(_nw283).ArraySet1(_1730_v0, 16)
						(_nw283).ArraySet1(_1730_v0, 17)
						(_nw283).ArraySet1(_1730_v0, 18)
						(_nw283).ArraySet1(false, 19)
						(_nw283).ArraySet1(_1730_v0, 20)
						(_nw283).ArraySet1(_1730_v0, 21)
						(_nw283).ArraySet1(_1730_v0, 22)
						(_nw283).ArraySet1(_1730_v0, 23)
						(_nw283).ArraySet1(_1730_v0, 24)
						_1756_v21 = _nw283
						var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1756_v21), 0))
						_ = _index293
						(_1756_v21).ArraySet1(_1730_v0, (_index293).Int())
						var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1756_v21), 0))
						_ = _index294
						(_1756_v21).ArraySet1(Companion_Default___.Fm2(_1731_v1, globalState), (_index294).Int())
						var _1757_v22 T0
						_ = _1757_v22
						var _nw284 *C3 = New_C3_()
						_ = _nw284
						_nw284.Ctor__()
						_1757_v22 = _nw284
						var _1758_v23 _dafny.Map
						_ = _1758_v23
						_1758_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1730_v0, _1757_v22)
						var _1759_v24 _dafny.Sequence
						_ = _1759_v24
						_1759_v24 = _dafny.SeqOf(_1758_v23)
						var _1760_v25 _dafny.Sequence
						_ = _1760_v25
						_1760_v25 = _dafny.SeqOf((_1759_v24).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1759_v24).Cardinality()))).Uint32()).(_dafny.Map))
						var _1761_v26 _dafny.Sequence
						_ = _1761_v26
						_1761_v26 = _dafny.SeqOf(false)
						var _1762_v27 _dafny.Map
						_ = _1762_v27
						_1762_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1761_v26, _dafny.UnicodeSeqOfUtf8Bytes("ulcd"))
						var _1763_v28 _dafny.Sequence
						_ = _1763_v28
						_1763_v28 = _dafny.SeqOf(Companion_Default___.Fm27((_1756_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1756_v21), 0))).Int()).(bool), (_1756_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1756_v21), 0))).Int()).(bool), globalState))
						var _1764_v29 _dafny.Set
						_ = _1764_v29
						_1764_v29 = _dafny.SetOf(_1763_v28, _1763_v28, _1763_v28)
						Companion_Default___.M0((p0).Cmp(p0) == 0, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1760_v25).Cardinality()), p0), _1762_v27, ((_1764_v29).Difference(_1764_v29)).Cardinality(), globalState)
						var _1765_v33 _dafny.Sequence
						_ = _1765_v33
						_1765_v33 = _dafny.SeqOf(p0)
						var _1766_v34 _dafny.Map
						_ = _1766_v34
						_1766_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(225))
						var _1767_v35 _dafny.Sequence
						_ = _1767_v35
						_1767_v35 = _dafny.SeqOf(func() _dafny.Map {
							var _coll68 = _dafny.NewMapBuilder()
							_ = _coll68
							for _iter80 := _dafny.Iterate((_1765_v33).Elements()); ; {
								_compr_68, _ok80 := _iter80()
								if !_ok80 {
									break
								}
								var _1768_v32 _dafny.Int
								_1768_v32 = interface{}(_compr_68).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_1765_v33, _1768_v32) {
									_coll68.Add(Companion_Default___.SafeModuloInt(_1768_v32, p0), _dafny.IntOfInt64(251))
								}
							}
							return _coll68.ToMap()
						}(), _1766_v34)
						var _1769_v36 _dafny.Map
						_ = _1769_v36
						_1769_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
							var _coll69 = _dafny.NewMapBuilder()
							_ = _coll69
							for _iter81 := _dafny.Iterate((func() _dafny.Map {
								var _coll70 = _dafny.NewMapBuilder()
								_ = _coll70
								for _iter82 := _dafny.Iterate((_1767_v35).Elements()); ; {
									_compr_70, _ok82 := _iter82()
									if !_ok82 {
										break
									}
									var _1770_v31 _dafny.Map
									_1770_v31 = interface{}(_compr_70).(_dafny.Map)
									if _dafny.Companion_Sequence_.Contains(_1767_v35, _1770_v31) {
										_coll70.Add(_1770_v31, p0)
									}
								}
								return _coll70.ToMap()
							}()).Keys().Elements()); ; {
								_compr_69, _ok81 := _iter81()
								if !_ok81 {
									break
								}
								var _1771_v30 _dafny.Map
								_1771_v30 = interface{}(_compr_69).(_dafny.Map)
								if (func() _dafny.Map {
									var _coll71 = _dafny.NewMapBuilder()
									_ = _coll71
									for _iter83 := _dafny.Iterate((_1767_v35).Elements()); ; {
										_compr_71, _ok83 := _iter83()
										if !_ok83 {
											break
										}
										var _1770_v31 _dafny.Map
										_1770_v31 = interface{}(_compr_71).(_dafny.Map)
										if _dafny.Companion_Sequence_.Contains(_1767_v35, _1770_v31) {
											_coll71.Add(_1770_v31, p0)
										}
									}
									return _coll71.ToMap()
								}()).Contains(_1771_v30) {
									_coll69.Add(_1771_v30, _1730_v0)
								}
							}
							return _coll69.ToMap()
						}(), true)
						var _1772_v38 _dafny.Map
						_ = _1772_v38
						_1772_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
							var _coll72 = _dafny.NewMapBuilder()
							_ = _coll72
							for _iter84 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(83), _dafny.IntOfInt64(288))); ; {
								_compr_72, _ok84 := _iter84()
								if !_ok84 {
									break
								}
								var _1773_v37 _dafny.Int
								_1773_v37 = interface{}(_compr_72).(_dafny.Int)
								if ((_dafny.IntOfInt64(83)).Cmp(_1773_v37) <= 0) && ((_1773_v37).Cmp(_dafny.IntOfInt64(288)) < 0) {
									_coll72.Add((_1773_v37).Plus(p0), p0)
								}
							}
							return _coll72.ToMap()
						}(), (_1756_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1756_v21), 0))).Int()).(bool))
						_1769_v36 = (_1769_v36).Update(_1772_v38, _1730_v0)
						var _1774_v39 *C2
						_ = _1774_v39
						var _nw285 *C2 = New_C2_()
						_ = _nw285
						_nw285.Ctor__(p0)
						_1774_v39 = _nw285
					}
					var _1775_v40 _dafny.Map
					_ = _1775_v40
					_1775_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1730_v0), _dafny.UnicodeSeqOfUtf8Bytes("l"))
					var _1776_v41 _dafny.Sequence
					_ = _1776_v41
					_1776_v41 = _dafny.SeqOf(_1775_v40)
					Companion_Default___.M0(!(_1730_v0) || (_1730_v0), p0, (_1776_v41).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1776_v41).Cardinality()))).Uint32()).(_dafny.Map), (p0).Minus(p0), globalState)
					var _1777_v42 _dafny.Sequence
					_ = _1777_v42
					_1777_v42 = _dafny.UnicodeSeqOfUtf8Bytes("t")
					var _1778_v43 _dafny.Map
					_ = _1778_v43
					_1778_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1730_v0), _1777_v42)
					(globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_1777_v42, (func() _dafny.Sequence {
						if (_1778_v43).Contains(_1730_v0) {
							return (_1778_v43).Get(_1730_v0).(_dafny.Sequence)
						}
						return _1777_v42
					})())
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1779_v44 _dafny.Sequence
		_ = _1779_v44
		_1779_v44 = _dafny.UnicodeSeqOfUtf8Bytes("aapa")
		var _1780_v45 _dafny.Map
		_ = _1780_v45
		_1780_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1730_v0)
		var _1781_v46 _dafny.Map
		_ = _1781_v46
		_1781_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1780_v45).Cardinality(), p0)
		var _hi13 _dafny.Int = p0
		_ = _hi13
		for _1782_i3 := Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1779_v44).Cardinality()), Companion_Default___.Fm0(p0, (_1781_v46).Cardinality(), p0, _1730_v0, globalState)); _1782_i3.Cmp(_hi13) < 0; _1782_i3 = _1782_i3.Plus(_dafny.One) {
			var _1783_v47 T0
			_ = _1783_v47
			var _nw286 *C3 = New_C3_()
			_ = _nw286
			_nw286.Ctor__()
			_1783_v47 = _nw286
			_1783_v47 = _1783_v47
			var _1784_v48 _dafny.Array
			_ = _1784_v48
			var _len0_51 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_51
			var _nw287 _dafny.Array
			_ = _nw287
			if _len0_51.Cmp(_dafny.Zero) == 0 {
				_nw287 = _dafny.NewArray(_len0_51)
			} else {
				var _init51 func(_dafny.Int) bool = (func(_1785_v44 _dafny.Sequence) func(_dafny.Int) bool {
					return func(_1786_i4 _dafny.Int) bool {
						return _dafny.Companion_Sequence_.Contains(_1785_v44, _dafny.CodePoint('n'))
					}
				})(_1779_v44)
				_ = _init51
				var _element0_51 = _init51(_dafny.Zero)
				_ = _element0_51
				_nw287 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
				(_nw287).ArraySet1(_element0_51, 0)
				var _nativeLen0_51 = (_len0_51).Int()
				_ = _nativeLen0_51
				for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
					(_nw287).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
				}
			}
			_1784_v48 = _nw287
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1784_v48), 0))
			_ = _index295
			(_1784_v48).ArraySet1(_1730_v0, (_index295).Int())
			var _1787_v49 _dafny.Array
			_ = _1787_v49
			var _nw288 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
			_ = _nw288
			_1787_v49 = _nw288
			var _1788_v50 _dafny.Map
			_ = _1788_v50
			_1788_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1787_v49, _1730_v0)
			var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1784_v48), 0))
			_ = _index296
			var _rhs299 _dafny.Int = _1782_i3
			_ = _rhs299
			var _rhs300 bool = !(_1788_v50).Equals((_1788_v50).Merge((Companion_D22_.Create_DC56_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1787_v49, false))).Dtor_cf80()))
			_ = _rhs300
			var _lhs282 *GlobalState = globalState
			_ = _lhs282
			var _lhs283 _dafny.Array = _1784_v48
			_ = _lhs283
			var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1784_v48), 0))
			_ = _lhs284
			_lhs282.F16 = _rhs299
			(_lhs283).ArraySet1(_rhs300, (_lhs284).Int())
			var _1789_v51 _dafny.Map
			_ = _1789_v51
			_1789_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1730_v0, _1782_i3)
			var _1790_v52 _dafny.MultiSet
			_ = _1790_v52
			_1790_v52 = _dafny.MultiSetOf(_1782_i3)
			var _1791_v53 _dafny.Set
			_ = _1791_v53
			_1791_v53 = _dafny.SetOf(p0, p0, p0, p0)
			var _1792_v54 _dafny.Sequence
			_ = _1792_v54
			_1792_v54 = _dafny.SeqOf((_1784_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1784_v48), 0))).Int()).(bool), (_1784_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1784_v48), 0))).Int()).(bool))
			var _1793_v55 D4
			_ = _1793_v55
			_1793_v55 = Companion_D4_.Create_DC9_(_1782_i3)
			var _1794_v56 _dafny.Sequence
			_ = _1794_v56
			_1794_v56 = _dafny.SeqOf(p0, _dafny.IntOfInt64(515))
			var _1795_v57 D4
			_ = _1795_v57
			_1795_v57 = Companion_D4_.Create_DC10_(_dafny.IntOfUint32((_1792_v54).Cardinality()), (_1731_v1).Cardinality(), _1730_v0, (_1784_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1784_v48), 0))).Int()).(bool), _1782_i3)
			var _1796_v58 _dafny.CodePoint
			_ = _1796_v58
			_1796_v58 = _dafny.CodePoint('k')
			var _1797_v59 _dafny.MultiSet
			_ = _1797_v59
			_1797_v59 = _dafny.MultiSetOf(_1796_v58, _1796_v58, _1796_v58, _1796_v58)
			var _1798_v60 _dafny.Set
			_ = _1798_v60
			_1798_v60 = _dafny.SetOf(_1730_v0, _1730_v0, _1730_v0)
			var _1799_v61 _dafny.Array
			_ = _1799_v61
			var _nwElement0_52 _dafny.Int = _1782_i3
			_ = _nwElement0_52
			var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(24))
			_ = _nw289
			(_nw289).ArraySet1(_nwElement0_52, 0)
			(_nw289).ArraySet1(_1782_i3, 1)
			(_nw289).ArraySet1((_1781_v46).Cardinality(), 2)
			(_nw289).ArraySet1(((_1789_v51).Cardinality()).Plus((_1790_v52).Cardinality()), 3)
			(_nw289).ArraySet1(p0, 4)
			(_nw289).ArraySet1(Companion_Default___.SafeModuloInt(p0, (_1791_v53).Cardinality()), 5)
			(_nw289).ArraySet1((p0).Times(_dafny.IntOfUint32((_1792_v54).Cardinality())), 6)
			(_nw289).ArraySet1(_dafny.IntOfInt64(848), 7)
			(_nw289).ArraySet1(p0, 8)
			(_nw289).ArraySet1((_1793_v55).Dtor_cf11(), 9)
			(_nw289).ArraySet1(((_1781_v46).Cardinality()).Minus(_1782_i3), 10)
			(_nw289).ArraySet1((_1794_v56).Select((Companion_Default___.SafeIndex(_1782_i3, _dafny.IntOfUint32((_1794_v56).Cardinality()))).Uint32()).(_dafny.Int), 11)
			(_nw289).ArraySet1((_1795_v57).Dtor_cf12(), 12)
			(_nw289).ArraySet1(p0, 13)
			(_nw289).ArraySet1((p0).Minus((_1797_v59).Cardinality()), 14)
			(_nw289).ArraySet1((func() _dafny.Int {
				if (_1790_v52).Contains(_1782_i3) {
					return (_1790_v52).Multiplicity(_1782_i3)
				}
				return p0
			})(), 15)
			(_nw289).ArraySet1(_1782_i3, 16)
			(_nw289).ArraySet1(_1782_i3, 17)
			(_nw289).ArraySet1(p0, 18)
			(_nw289).ArraySet1((func() _dafny.Int {
				if (_1781_v46).Contains((_1798_v60).Cardinality()) {
					return (_1781_v46).Get((_1798_v60).Cardinality()).(_dafny.Int)
				}
				return p0
			})(), 19)
			(_nw289).ArraySet1(_1782_i3, 20)
			(_nw289).ArraySet1(_1782_i3, 21)
			(_nw289).ArraySet1(Companion_Default___.Fm0(p0, _dafny.IntOfInt64(827), p0, !((_1784_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1784_v48), 0))).Int()).(bool)), globalState), 22)
			(_nw289).ArraySet1(_dafny.IntOfInt64(-744), 23)
			_1799_v61 = _nw289
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_1799_v61), 0))
			_ = _index297
			(_1799_v61).ArraySet1(_1782_i3, (_index297).Int())
			var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_1799_v61), 0))
			_ = _index298
			(_1799_v61).ArraySet1(_dafny.IntOfInt64(-85), (_index298).Int())
			_1789_v51 = (_1789_v51).Update(!(true), (_1799_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_1799_v61), 0))).Int()).(_dafny.Int))
		}
		var _1800_v62 *C0
		_ = _1800_v62
		var _nw290 *C0 = New_C0_()
		_ = _nw290
		_nw290.Ctor__(_1730_v0)
		_1800_v62 = _nw290
		(globalState).F7 = p0
		var _1801_i5 _dafny.Int
		_ = _1801_i5
		_1801_i5 = _dafny.Zero
		{
			for _1800_v62.F27 {
				{
					if (_1801_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1801_i5 = (_1801_i5).Plus(_dafny.One)
					(globalState).F7 = p0
					var _1802_v63 _dafny.MultiSet
					_ = _1802_v63
					_1802_v63 = _dafny.MultiSetOf(_1800_v62.F27)
					_1802_v63 = Companion_Default___.Fm52(_1800_v62.F27, p0, globalState)
					var _1803_v64 _dafny.Array
					_ = _1803_v64
					var _len0_52 _dafny.Int = _dafny.IntOfInt64(29)
					_ = _len0_52
					var _nw291 _dafny.Array
					_ = _nw291
					if _len0_52.Cmp(_dafny.Zero) == 0 {
						_nw291 = _dafny.NewArray(_len0_52)
					} else {
						var _init52 func(_dafny.Int) _dafny.Set = (func(_1804_v44 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
							return func(_1805_i6 _dafny.Int) _dafny.Set {
								return _dafny.SetOf(_1804_v44, _1804_v44)
							}
						})(_1779_v44)
						_ = _init52
						var _element0_52 = _init52(_dafny.Zero)
						_ = _element0_52
						_nw291 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
						(_nw291).ArraySet1(_element0_52, 0)
						var _nativeLen0_52 = (_len0_52).Int()
						_ = _nativeLen0_52
						for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
							(_nw291).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
						}
					}
					_1803_v64 = _nw291
					var _1806_v65 _dafny.Set
					_ = _1806_v65
					_1806_v65 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ycoxy"))
					var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1803_v64), 0))
					_ = _index299
					(_1803_v64).ArraySet1((_1806_v65).Intersection(func() _dafny.Set {
						var _coll73 = _dafny.NewBuilder()
						_ = _coll73
						for _iter85 := _dafny.Iterate((_1806_v65).Elements()); ; {
							_compr_73, _ok85 := _iter85()
							if !_ok85 {
								break
							}
							var _1807_v66 _dafny.Sequence
							_1807_v66 = interface{}(_compr_73).(_dafny.Sequence)
							if (_1806_v65).Contains(_1807_v66) {
								_coll73.Add(_1807_v66)
							}
						}
						return _coll73.ToSet()
					}()), (_index299).Int())
					var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1803_v64), 0))
					_ = _index300
					(_1803_v64).ArraySet1(_1806_v65, (_index300).Int())
					var _1808_v67 _dafny.CodePoint
					_ = _1808_v67
					_1808_v67 = _dafny.CodePoint('b')
					_1808_v67 = _1808_v67
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1809_v68 _dafny.Array
		_ = _1809_v68
		var _nwElement0_53 _dafny.Int = (_dafny.Zero).Minus(p0)
		_ = _nwElement0_53
		var _nw292 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.One)
		_ = _nw292
		(_nw292).ArraySet1(_nwElement0_53, 0)
		_1809_v68 = _nw292
		r0 = _1809_v68
		r1 = (func() bool {
			if _1730_v0 {
				return _1800_v62.F27
			}
			return _1800_v62.F27
		})()
		return r0, r1
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f23 _dafny.Map
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f23 = _dafny.EmptyMap
	return &_this
}

type CompanionStruct_C7_ struct {
}

var Companion_C7_ = CompanionStruct_C7_{}

func (_this *C7) Equals(other *C7) bool {
	return _this == other
}

func (_this *C7) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C7)
	return ok && _this.Equals(other)
}

func (*C7) String() string {
	return "_module.C7"
}

func Type_C7_() _dafny.TypeDescriptor {
	return type_C7_{}
}

type type_C7_ struct {
}

func (_this type_C7_) Default() interface{} {
	return (*C7)(nil)
}

func (_this type_C7_) String() string {
	return "main.C7"
}
func (_this *C7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C7{}
var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) Ctor__(f23 _dafny.Map) {
	{
		(_this)._f23 = f23
	}
}
func (_this *C7) Fm7(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C7) Fm5(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if false {
				return !(true)
			}
			return true
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uqyvs"), _dafny.UnicodeSeqOfUtf8Bytes("egwtgbjvl")))
	}
}
func (_this *C7) Fm6(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(!(!(true)) || ((_dafny.SetOf(_dafny.IntOfInt64(-747))).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfInt64(158), _dafny.IntOfInt64(-959)))))
	}
}
func (_this *C7) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _1810_v0 *C6
		_ = _1810_v0
		var _nw293 *C6 = New_C6_()
		_ = _nw293
		_nw293.Ctor__()
		_1810_v0 = _nw293
		var _hi14 _dafny.Int = p3
		_ = _hi14
		for _1811_i0 := p3; _1811_i0.Cmp(_hi14) < 0; _1811_i0 = _1811_i0.Plus(_dafny.One) {
			(globalState).F16 = (_dafny.IntOfInt64(914)).Minus((_1811_i0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(p3)).Cardinality())))
			var _1812_v1 _dafny.Array
			_ = _1812_v1
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_53
			var _nw294 _dafny.Array
			_ = _nw294
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw294 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) _dafny.Sequence = func(_1813_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("rhc")
				}
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw294 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw294).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw294).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_1812_v1 = _nw294
			_1812_v1 = (func() _dafny.Array {
				if (_1811_i0).Cmp(_1811_i0) < 0 {
					return _1812_v1
				}
				return _1812_v1
			})()
			var _1814_v2 T1
			_ = _1814_v2
			var _nw295 *C5 = New_C5_()
			_ = _nw295
			_nw295.Ctor__(p1)
			_1814_v2 = _nw295
			_1814_v2 = _1814_v2
			var _1815_v4 _dafny.Array
			_ = _1815_v4
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_54
			var _nw296 _dafny.Array
			_ = _nw296
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw296 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) _dafny.Sequence = (func(_1816_p1 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_1817_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(func() _dafny.Map {
							var _coll74 = _dafny.NewMapBuilder()
							_ = _coll74
							for _iter86 := _dafny.Iterate((_dafny.SetOf(_dafny.SetOf(_1816_p1), _dafny.SetOf(_1816_p1))).Elements()); ; {
								_compr_74, _ok86 := _iter86()
								if !_ok86 {
									break
								}
								var _1818_v3 _dafny.Set
								_1818_v3 = interface{}(_compr_74).(_dafny.Set)
								if (_dafny.SetOf(_dafny.SetOf(_1816_p1), _dafny.SetOf(_1816_p1))).Contains(_1818_v3) {
									_coll74.Add(_1818_v3, _1816_p1)
								}
							}
							return _coll74.ToMap()
						}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_1816_p1), _1816_p1))
					}
				})(p1)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw296 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw296).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw296).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_1815_v4 = _nw296
			var _1819_v5 _dafny.Set
			_ = _1819_v5
			_1819_v5 = _dafny.SetOf(false, p1)
			var _1820_v6 _dafny.Map
			_ = _1820_v6
			_1820_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1819_v5, p1)
			var _1821_v7 _dafny.Sequence
			_ = _1821_v7
			_1821_v7 = _dafny.SeqOf(_1820_v6)
			var _1822_v8 _dafny.Set
			_ = _1822_v8
			_1822_v8 = _1819_v5
			var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_1815_v4), 0))
			_ = _index301
			(_1815_v4).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1821_v7, _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1822_v8, p1)).Update(_1822_v8, false))), (_index301).Int())
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_1815_v4), 0))
			_ = _index302
			(_1815_v4).ArraySet1(_1821_v7, (_index302).Int())
		}
		var _1823_v9 *C2
		_ = _1823_v9
		var _nw297 *C2 = New_C2_()
		_ = _nw297
		_nw297.Ctor__((_dafny.Zero).Minus(p0))
		_1823_v9 = _nw297
		var _1824_v10 _dafny.Set
		_ = _1824_v10
		_1824_v10 = _dafny.SetOf(_1823_v9, _1823_v9, _1823_v9)
		if (_1824_v10).IsSubsetOf(_dafny.SetOf(_1823_v9, _1823_v9)) {
			(globalState).F16 = p3
			var _1825_v11 _dafny.Sequence
			_ = _1825_v11
			_1825_v11 = _dafny.SeqOf((_1823_v9).F28())
			var _1826_v12 *C0
			_ = _1826_v12
			var _nw298 *C0 = New_C0_()
			_ = _nw298
			_nw298.Ctor__(((_1823_v9).F28()).Cmp(Companion_Default___.Fm0(p0, _dafny.IntOfUint32((_1825_v11).Cardinality()), Companion_Default___.Fm0(p3, p3, p0, p1, globalState), p1, globalState)) != 0)
			_1826_v12 = _nw298
			var _1827_v13 _dafny.Array
			_ = _1827_v13
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_55
			var _nw299 _dafny.Array
			_ = _nw299
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw299 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) bool = (func(_1828_v12 *C0, _1829_p1 bool) func(_dafny.Int) bool {
					return func(_1830_i3 _dafny.Int) bool {
						return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_1828_v12.F27))).IsDisjointFrom(_dafny.MultiSetOf(_1829_p1, !(_1828_v12.F27)))
					}
				})(_1826_v12, p1)
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw299 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw299).ArraySet1(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw299).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_1827_v13 = _nw299
			var _1831_v14 *C6
			_ = _1831_v14
			var _nw300 *C6 = New_C6_()
			_ = _nw300
			_nw300.Ctor__()
			_1831_v14 = _nw300
			var _1832_v15 _dafny.Sequence
			_ = _1832_v15
			_1832_v15 = _dafny.SeqOf(_1826_v12.F27)
			var _1833_v16 _dafny.MultiSet
			_ = _1833_v16
			_1833_v16 = _dafny.MultiSetOf((func() _dafny.Sequence {
				if _1826_v12.F27 {
					return _1832_v15
				}
				return _1832_v15
			})())
			var _1834_v17 _dafny.Sequence
			_ = _1834_v17
			_1834_v17 = _dafny.UnicodeSeqOfUtf8Bytes("mdrtqm")
			var _rhs301 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_1823_v9).F28()))
			_ = _rhs301
			var _rhs302 _dafny.Int = _dafny.IntOfUint32((_1834_v17).Cardinality())
			_ = _rhs302
			var _rhs303 _dafny.Array = _1827_v13
			_ = _rhs303
			var _rhs304 *C6 = _1831_v14
			_ = _rhs304
			var _rhs305 _dafny.MultiSet = _1833_v16
			_ = _rhs305
			var _lhs285 *GlobalState = globalState
			_ = _lhs285
			r1 = _rhs301
			_lhs285.F7 = _rhs302
			_1827_v13 = _rhs303
			_1831_v14 = _rhs304
			_1833_v16 = _rhs305
			var _1835_v18 _dafny.Map
			_ = _1835_v18
			_1835_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _1836_v19 *C1
			_ = _1836_v19
			var _nw301 *C1 = New_C1_()
			_ = _nw301
			_nw301.Ctor__(((_1823_v9).F28()).Plus(Companion_Default___.Fm0(p3, (_1823_v9).F28(), (_1823_v9).F28(), true, globalState)), _dafny.SeqOf(_1835_v18, _1835_v18, _1835_v18))
			_1836_v19 = _nw301
			(globalState).F7 = (_dafny.Zero).Minus((_1823_v9).Fm17(false, _1825_v11, p3, _1832_v15, globalState))
		} else {
			var _1837_v20 D17
			_ = _1837_v20
			_1837_v20 = Companion_D17_.Create_DC40_()
			var _source23 D17 = _1837_v20
			_ = _source23
			if _source23.Is_DC39() {
				var _1838___mcc_h0 _dafny.Int = _source23.Get_().(D17_DC39).Cf52
				_ = _1838___mcc_h0
				var _1839___mcc_h1 _dafny.Map = _source23.Get_().(D17_DC39).Cf53
				_ = _1839___mcc_h1
				var _1840___mcc_h2 bool = _source23.Get_().(D17_DC39).Cf54
				_ = _1840___mcc_h2
				var _1841___mcc_h3 _dafny.Map = _source23.Get_().(D17_DC39).Cf55
				_ = _1841___mcc_h3
				var _1842___mcc_h4 _dafny.Int = _source23.Get_().(D17_DC39).Cf56
				_ = _1842___mcc_h4
				var _1843_cf56 _dafny.Int = _1842___mcc_h4
				_ = _1843_cf56
				var _1844_cf55 _dafny.Map = _1841___mcc_h3
				_ = _1844_cf55
				var _1845_cf54 bool = _1840___mcc_h2
				_ = _1845_cf54
				var _1846_cf53 _dafny.Map = _1839___mcc_h1
				_ = _1846_cf53
				var _1847_cf52 _dafny.Int = _1838___mcc_h0
				_ = _1847_cf52
				_1847_cf52 = _dafny.IntOfInt64(-36)
				var _1848_v21 _dafny.Map
				_ = _1848_v21
				_1848_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
				var _1849_v22 _dafny.Sequence
				_ = _1849_v22
				_1849_v22 = _dafny.SeqOf(_1848_v21, _1848_v21)
				(globalState).F13 = Companion_Default___.Fm2((_1849_v22).Select((Companion_Default___.SafeIndex(_1847_cf52, _dafny.IntOfUint32((_1849_v22).Cardinality()))).Uint32()).(_dafny.Map), globalState)
				var _1850_v23 _dafny.MultiSet
				_ = _1850_v23
				_1850_v23 = _dafny.MultiSetOf(!(p1), p1)
				var _1851_v24 *C2
				_ = _1851_v24
				var _nw302 *C2 = New_C2_()
				_ = _nw302
				_nw302.Ctor__((_1850_v23).Cardinality())
				_1851_v24 = _nw302
				var _1852_v25 _dafny.Sequence
				_ = _1852_v25
				_1852_v25 = _dafny.SeqOf(_1847_cf52, _dafny.IntOfInt64(955), (_1823_v9).F28(), (_1850_v23).Cardinality())
				(globalState).F0 = _dafny.Companion_Sequence_.Equal(_1852_v25, _1852_v25)
			} else if _source23.Is_DC40() {
				var _1853_v26 _dafny.Sequence
				_ = _1853_v26
				_1853_v26 = _dafny.SeqOf(p1, p1)
				var _1854_v27 _dafny.Sequence
				_ = _1854_v27
				_1854_v27 = _1853_v26
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((p2), 0))
				_ = _index303
				(p2).ArraySet1((_1823_v9).Fm17(p1, _dafny.SeqOf((_1823_v9).F28()), (_1823_v9).F28(), _1854_v27, globalState), (_index303).Int())
				var _1855_v28 _dafny.Sequence
				_ = _1855_v28
				_1855_v28 = _dafny.UnicodeSeqOfUtf8Bytes("ug")
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((p2), 0))
				_ = _index304
				(p2).ArraySet1(_dafny.IntOfUint32((_1855_v28).Cardinality()), (_index304).Int())
				(globalState).F16 = p3
				var _1856_v29 *C5
				_ = _1856_v29
				var _nw303 *C5 = New_C5_()
				_ = _nw303
				_nw303.Ctor__((_this).Fm7(p1, (_dafny.Zero).Minus((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)), p1, globalState))
				_1856_v29 = _nw303
				var _1857_v30 _dafny.CodePoint
				_ = _1857_v30
				_1857_v30 = _dafny.CodePoint('u')
				var _1858_v31 _dafny.Array
				_ = _1858_v31
				var _nw304 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
				_ = _nw304
				_1858_v31 = _nw304
				var _1859_v32 _dafny.Map
				_ = _1859_v32
				_1859_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1856_v29).F24())
				var _1860_v33 _dafny.Map
				_ = _1860_v33
				_1860_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1856_v29).F24(), _1859_v32)
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1858_v31), 0))
				_ = _index305
				(_1858_v31).ArraySet1((Companion_D23_.Create_DC60_(_1860_v33)).Dtor_cf85(), (_index305).Int())
				var _1861_v34 _dafny.Set
				_ = _1861_v34
				_1861_v34 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("uvqcohcw"))
				var _1862_v35 D0
				_ = _1862_v35
				_1862_v35 = Companion_D0_.Create_DC0_(p1)
				var _1863_v36 _dafny.Set
				_ = _1863_v36
				_1863_v36 = _dafny.SetOf(_1857_v30)
				var _1864_v37 _dafny.MultiSet
				_ = _1864_v37
				_1864_v37 = _dafny.MultiSetOf(_dafny.IntOfInt64(648))
				var _1865_v38 _dafny.Map
				_ = _1865_v38
				_1865_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1823_v9).F28(), p3)
				var _1866_v39 _dafny.Map
				_ = _1866_v39
				_1866_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1865_v38, p1)
				var _1867_v40 _dafny.Map
				_ = _1867_v40
				_1867_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1856_v29).F24(), (Companion_Default___.Fm54(globalState)).Cardinality())
				var _1868_v41 _dafny.MultiSet
				_ = _1868_v41
				_1868_v41 = _dafny.MultiSetOf(true)
				var _1869_v42 _dafny.Sequence
				_ = _1869_v42
				_1869_v42 = _dafny.SeqOf((_1868_v41).Cardinality())
				var _1870_v43 D15
				_ = _1870_v43
				_1870_v43 = Companion_D15_.Create_DC33_(_1869_v42)
				var _1871_v44 _dafny.Array
				_ = _1871_v44
				var _nwElement0_54 bool = (_1856_v29).F24()
				_ = _nwElement0_54
				var _nw305 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(29))
				_ = _nw305
				(_nw305).ArraySet1(_nwElement0_54, 0)
				(_nw305).ArraySet1((_1861_v34).IsDisjointFrom(Companion_Default___.Fm53(p1, globalState)), 1)
				(_nw305).ArraySet1(!((_1856_v29).F24()), 2)
				(_nw305).ArraySet1(p1, 3)
				(_nw305).ArraySet1((_1856_v29).F24(), 4)
				(_nw305).ArraySet1(p1, 5)
				(_nw305).ArraySet1(!(p1), 6)
				(_nw305).ArraySet1(false, 7)
				(_nw305).ArraySet1((_1856_v29).F24(), 8)
				(_nw305).ArraySet1(p1, 9)
				(_nw305).ArraySet1(!((_1862_v35).Dtor_cf0()), 10)
				(_nw305).ArraySet1(p1, 11)
				(_nw305).ArraySet1((_1863_v36).IsSubsetOf(_1863_v36), 12)
				(_nw305).ArraySet1((_1856_v29).F24(), 13)
				(_nw305).ArraySet1((_1856_v29).Fm6(_dafny.IntOfUint32((_1853_v26).Cardinality()), globalState), 14)
				(_nw305).ArraySet1(!((_1864_v37).IsDisjointFrom(_1864_v37)), 15)
				(_nw305).ArraySet1((_1856_v29).F24(), 16)
				(_nw305).ArraySet1(false, 17)
				(_nw305).ArraySet1((_1856_v29).F24(), 18)
				(_nw305).ArraySet1(p1, 19)
				(_nw305).ArraySet1((_1856_v29).F24(), 20)
				(_nw305).ArraySet1((_1856_v29).F24(), 21)
				(_nw305).ArraySet1((_1866_v39).Contains(_1865_v38), 22)
				(_nw305).ArraySet1(p1, 23)
				(_nw305).ArraySet1(!(p1), 24)
				(_nw305).ArraySet1((_1856_v29).F24(), 25)
				(_nw305).ArraySet1(((_1867_v40).Cardinality()).Cmp((_1823_v9).F28()) != 0, 26)
				(_nw305).ArraySet1((_this).Fm6((_1823_v9).Fm17(!(p1), (_1870_v43).Dtor_cf44(), (_1823_v9).F28(), _1854_v27, globalState), globalState), 27)
				(_nw305).ArraySet1(true, 28)
				_1871_v44 = _nw305
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1871_v44), 0))
				_ = _index306
				(_1871_v44).ArraySet1((_1856_v29).F24(), (_index306).Int())
				var _1872_v45 _dafny.Sequence
				_ = _1872_v45
				_1872_v45 = _dafny.SeqOf(_1865_v38)
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1858_v31), 0))
				_ = _index307
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1871_v44), 0))
				_ = _index308
				var _rhs306 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_1856_v29).F24() {
						return _1857_v30
					}
					return _1857_v30
				})()
				_ = _rhs306
				var _rhs307 _dafny.Map = (_1860_v33).Merge((func() _dafny.Map {
					if p1 {
						return _1860_v33
					}
					return _1860_v33
				})())
				_ = _rhs307
				var _rhs308 bool = !(!((_1856_v29).F24()))
				_ = _rhs308
				var _rhs309 _dafny.Int = Companion_Default___.SafeDivisionInt((_1869_v42).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_1872_v45)).Cardinality(), _dafny.IntOfUint32((_1869_v42).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Fm0((_1869_v42).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1855_v28).Cardinality()), _dafny.IntOfUint32((_1869_v42).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(431), p0, Companion_Default___.Fm2(_1859_v32, globalState), globalState))
				_ = _rhs309
				var _rhs310 _dafny.Int = (_dafny.Zero).Minus((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
				_ = _rhs310
				var _lhs286 _dafny.Array = _1858_v31
				_ = _lhs286
				var _lhs287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1858_v31), 0))
				_ = _lhs287
				var _lhs288 _dafny.Array = _1871_v44
				_ = _lhs288
				var _lhs289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1871_v44), 0))
				_ = _lhs289
				var _lhs290 *GlobalState = globalState
				_ = _lhs290
				_1857_v30 = _rhs306
				(_lhs286).ArraySet1(_rhs307, (_lhs287).Int())
				(_lhs288).ArraySet1(_rhs308, (_lhs289).Int())
				r1 = _rhs309
				_lhs290.F7 = _rhs310
			} else if _source23.Is_DC41() {
				var _1873___mcc_h5 _dafny.Set = _source23.Get_().(D17_DC41).Cf57
				_ = _1873___mcc_h5
				var _1874_cf57 _dafny.Set = _1873___mcc_h5
				_ = _1874_cf57
				(globalState).F7 = _dafny.IntOfInt64(577)
				var _1875_v46 _dafny.Map
				_ = _1875_v46
				_1875_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-300), p1)
				var _1876_v47 _dafny.Sequence
				_ = _1876_v47
				_1876_v47 = _dafny.SeqOf(_1875_v46)
				var _1877_v48 *C1
				_ = _1877_v48
				var _nw306 *C1 = New_C1_()
				_ = _nw306
				_nw306.Ctor__(p0, _1876_v47)
				_1877_v48 = _nw306
				(globalState).F21 = p1
				var _1878_v49 D19
				_ = _1878_v49
				_1878_v49 = Companion_D19_.Create_DC48_(p0, p1)
				r1 = (_1878_v49).Dtor_cf65()
			} else {
				var _1879___mcc_h6 _dafny.Array = _source23.Get_().(D17_DC38).Cf51
				_ = _1879___mcc_h6
				var _1880_cf51 _dafny.Array = _1879___mcc_h6
				_ = _1880_cf51
				var _1881_v50 _dafny.Sequence
				_ = _1881_v50
				_1881_v50 = _dafny.SeqOf(p1, p1, p1)
				var _1882_v51 _dafny.MultiSet
				_ = _1882_v51
				_1882_v51 = _dafny.MultiSetOf(p1)
				(globalState).F2 = ((_1881_v50).Select((Companion_Default___.SafeIndex((_1882_v51).Cardinality(), _dafny.IntOfUint32((_1881_v50).Cardinality()))).Uint32()).(bool)) == (p1)
				var _1883_v52 _dafny.MultiSet
				_ = _1883_v52
				_1883_v52 = _dafny.MultiSetOf(Companion_Default___.Fm0(p3, p0, (_1823_v9).F28(), p1, globalState), (_1823_v9).F28(), (_1823_v9).F28(), (_1823_v9).F28(), (_1823_v9).F28())
				(globalState).F2 = (_1883_v52).IsDisjointFrom(_dafny.MultiSetOf((_1823_v9).F28()))
				var _1884_v53 _dafny.CodePoint
				_ = _1884_v53
				_1884_v53 = _dafny.CodePoint('n')
				_1884_v53 = Companion_Default___.Fm34(globalState)
				var _1885_v54 _dafny.Array
				_ = _1885_v54
				var _nw307 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw307
				_1885_v54 = _nw307
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_1885_v54), 0))
				_ = _index309
				(_1885_v54).ArraySet1(!(p1) || (p1), (_index309).Int())
				var _1886_v55 _dafny.Set
				_ = _1886_v55
				_1886_v55 = _dafny.SetOf(true)
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_1885_v54), 0))
				_ = _index310
				(_1885_v54).ArraySet1((func() bool {
					if !(p1) {
						return (_1886_v55).IsSubsetOf(_1886_v55)
					}
					return !(p1)
				})(), (_index310).Int())
			}
			var _1887_v56 _dafny.CodePoint
			_ = _1887_v56
			_1887_v56 = _dafny.CodePoint('g')
			var _1888_v57 _dafny.Array
			_ = _1888_v57
			var _nwElement0_55 _dafny.CodePoint = _1887_v56
			_ = _nwElement0_55
			var _nw308 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(13))
			_ = _nw308
			(_nw308).ArraySet1CodePoint(_nwElement0_55, 0)
			(_nw308).ArraySet1CodePoint(_1887_v56, 1)
			(_nw308).ArraySet1CodePoint(_1887_v56, 2)
			(_nw308).ArraySet1CodePoint(_1887_v56, 3)
			(_nw308).ArraySet1CodePoint(_1887_v56, 4)
			(_nw308).ArraySet1CodePoint(_1887_v56, 5)
			(_nw308).ArraySet1CodePoint(_dafny.CodePoint('k'), 6)
			(_nw308).ArraySet1CodePoint(_dafny.CodePoint('f'), 7)
			(_nw308).ArraySet1CodePoint(_1887_v56, 8)
			(_nw308).ArraySet1CodePoint(_1887_v56, 9)
			(_nw308).ArraySet1CodePoint(_1887_v56, 10)
			(_nw308).ArraySet1CodePoint(_1887_v56, 11)
			(_nw308).ArraySet1CodePoint(_1887_v56, 12)
			_1888_v57 = _nw308
			var _1889_v58 _dafny.Map
			_ = _1889_v58
			_1889_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1888_v57)
			_1888_v57 = (func() _dafny.Array {
				if (_1889_v58).Contains((_1823_v9).F28()) {
					return (_1889_v58).Get((_1823_v9).F28()).(_dafny.Array)
				}
				return _1888_v57
			})()
			var _1890_v59 _dafny.Sequence
			_ = _1890_v59
			_1890_v59 = _dafny.SeqOf(p1)
			var _1891_v60 _dafny.Sequence
			_ = _1891_v60
			_1891_v60 = _dafny.UnicodeSeqOfUtf8Bytes("sbxhia")
			var _1892_v61 _dafny.MultiSet
			_ = _1892_v61
			_1892_v61 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1891_v60).Cardinality()), p3)
			var _1893_v62 _dafny.Map
			_ = _1893_v62
			_1893_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_1890_v59).Cardinality())).Times((_1892_v61).Cardinality()), _1887_v56)
			var _rhs311 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_1893_v62).Contains(Companion_Default___.SafeModuloInt((_1823_v9).F28(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1891_v60).Cardinality()), (_1823_v9).F28())).Cardinality()))) {
					return (_1893_v62).Get(Companion_Default___.SafeModuloInt((_1823_v9).F28(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1891_v60).Cardinality()), (_1823_v9).F28())).Cardinality()))).(_dafny.CodePoint)
				}
				return _1887_v56
			})()
			_ = _rhs311
			var _rhs312 bool = (_1823_v9).Fm6((_1823_v9).F28(), globalState)
			_ = _rhs312
			var _lhs291 *GlobalState = globalState
			_ = _lhs291
			_1887_v56 = _rhs311
			_lhs291.F2 = _rhs312
			var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((p2), 0))
			_ = _index311
			(p2).ArraySet1(p3, (_index311).Int())
			var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((p2), 0))
			_ = _index312
			(p2).ArraySet1(_dafny.IntOfUint32((_1891_v60).Cardinality()), (_index312).Int())
			r1 = _dafny.IntOfInt64(794)
		}
		var _1894_v63 *C2
		_ = _1894_v63
		var _nw309 *C2 = New_C2_()
		_ = _nw309
		_nw309.Ctor__(((_1823_v9).F28()).Plus(p0))
		_1894_v63 = _nw309
		var _1895_v64 _dafny.Map
		_ = _1895_v64
		_1895_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1823_v9).F28())
		var _1896_v65 _dafny.Set
		_ = _1896_v65
		_1896_v65 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0), _1895_v64)
		var _1897_v66 _dafny.Map
		_ = _1897_v66
		_1897_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1896_v65).IsSubsetOf(_1896_v65))
		var _1898_i4 _dafny.Int
		_ = _1898_i4
		_1898_i4 = _dafny.Zero
		{
			for !((func() bool {
				if (_1897_v66).Contains(p1) {
					return (_1897_v66).Get(p1).(bool)
				}
				return true
			})()) {
				{
					if (_1898_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1898_i4 = (_1898_i4).Plus(_dafny.One)
					var _1899_v67 *C3
					_ = _1899_v67
					var _nw310 *C3 = New_C3_()
					_ = _nw310
					_nw310.Ctor__()
					_1899_v67 = _nw310
					(globalState).F16 = _dafny.IntOfUint32((Companion_Default___.Fm1((_1823_v9).F28(), globalState)).Cardinality())
					var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((p2), 0))
					_ = _index313
					(p2).ArraySet1(p0, (_index313).Int())
					var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((p2), 0))
					_ = _index314
					(p2).ArraySet1(p0, (_index314).Int())
					var _1900_v68 _dafny.Array
					_ = _1900_v68
					var _nw311 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
					_ = _nw311
					_1900_v68 = _nw311
					var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1900_v68), 0))
					_ = _index315
					(_1900_v68).ArraySet1(p1, (_index315).Int())
					var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1900_v68), 0))
					_ = _index316
					(_1900_v68).ArraySet1(p1, (_index316).Int())
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		var _1901_v69 _dafny.Sequence
		_ = _1901_v69
		_1901_v69 = _dafny.SeqOf(p1, true, (p1) && (p1))
		var _1902_v70 _dafny.Sequence
		_ = _1902_v70
		_1902_v70 = _dafny.SeqOf((_1894_v63).F28(), _dafny.IntOfInt64(773))
		var _1903_v71 _dafny.Sequence
		_ = _1903_v71
		_1903_v71 = _dafny.UnicodeSeqOfUtf8Bytes("ygs")
		var _1904_i5 _dafny.Int
		_ = _1904_i5
		_1904_i5 = _dafny.Zero
		{
			for (_1901_v69).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_1823_v9).F28()), (_1902_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1903_v71).Cardinality()), _dafny.IntOfUint32((_1902_v70).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfUint32((_1901_v69).Cardinality()))).Uint32()).(bool) {
				{
					if (_1904_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_1904_i5 = (_1904_i5).Plus(_dafny.One)
					var _1905_v72 _dafny.Sequence
					_ = _1905_v72
					_1905_v72 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qxmym"), _1903_v71, _1903_v71, Companion_Default___.Fm1(_dafny.IntOfUint32((_1903_v71).Cardinality()), globalState), _1903_v71)
					var _1906_v73 _dafny.MultiSet
					_ = _1906_v73
					_1906_v73 = _dafny.MultiSetOf((_1905_v72).Select((Companion_Default___.SafeIndex((_1823_v9).F28(), _dafny.IntOfUint32((_1905_v72).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg108 _dafny.Int) interface{} {
							return coer108(arg108)
						}
					}(func(_1907_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					})), _1903_v71)
					var _1908_v74 _dafny.Map
					_ = _1908_v74
					_1908_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1823_v9).F28(), Companion_Default___.Fm55((_1894_v63).F28(), p1, p1, globalState))
					var _1909_v75 _dafny.Map
					_ = _1909_v75
					_1909_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1901_v69).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1901_v69).Cardinality()))).Uint32()).(bool), (func() _dafny.MultiSet {
						if (_1908_v74).Contains(p3) {
							return (_1908_v74).Get(p3).(_dafny.MultiSet)
						}
						return _1906_v73
					})())
					var _1910_v76 _dafny.Sequence
					_ = _1910_v76
					_1910_v76 = _dafny.SeqOf((_1906_v73).Update(_1903_v71, Companion_Default___.Abs((_1894_v63).F28())), _dafny.MultiSetFromSeq(_dafny.SeqOf(_1903_v71)))
					var _1911_v77 _dafny.Array
					_ = _1911_v77
					var _nwElement0_56 _dafny.MultiSet = _dafny.MultiSetOf(_1903_v71, _1903_v71)
					_ = _nwElement0_56
					var _nw312 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(17))
					_ = _nw312
					(_nw312).ArraySet1(_nwElement0_56, 0)
					(_nw312).ArraySet1(_1906_v73, 1)
					(_nw312).ArraySet1(_1906_v73, 2)
					(_nw312).ArraySet1((func() _dafny.MultiSet {
						if p1 {
							return _1906_v73
						}
						return _dafny.MultiSetFromSeq(_1905_v72)
					})(), 3)
					(_nw312).ArraySet1((_1906_v73).Intersection(_1906_v73), 4)
					(_nw312).ArraySet1(_1906_v73, 5)
					(_nw312).ArraySet1(_dafny.MultiSetOf(_1903_v71, _dafny.UnicodeSeqOfUtf8Bytes("biloj")), 6)
					(_nw312).ArraySet1(_1906_v73, 7)
					(_nw312).ArraySet1(_1906_v73, 8)
					(_nw312).ArraySet1(_1906_v73, 9)
					(_nw312).ArraySet1(_1906_v73, 10)
					(_nw312).ArraySet1((func() _dafny.MultiSet {
						if (_1909_v75).Contains(p1) {
							return (_1909_v75).Get(p1).(_dafny.MultiSet)
						}
						return _1906_v73
					})(), 11)
					(_nw312).ArraySet1((_1910_v76).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.IntOfUint32((_1910_v76).Cardinality()))).Uint32()).(_dafny.MultiSet), 12)
					(_nw312).ArraySet1(_1906_v73, 13)
					(_nw312).ArraySet1(_dafny.MultiSetOf(_1903_v71), 14)
					(_nw312).ArraySet1((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-724))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg109 _dafny.Int) interface{} {
							return coer109(arg109)
						}
					}(func(_1912_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('t')
					})))).Update(_1903_v71, Companion_Default___.Abs((_1894_v63).F28())), 15)
					(_nw312).ArraySet1(_1906_v73, 16)
					_1911_v77 = _nw312
					var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1911_v77), 0))
					_ = _index317
					(_1911_v77).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1903_v71), _1905_v72)), (_index317).Int())
					var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1911_v77), 0))
					_ = _index318
					(_1911_v77).ArraySet1(_1906_v73, (_index318).Int())
					(globalState).F16 = ((_dafny.Zero).Minus((_1894_v63).F28())).Plus(((_1894_v63).F28()).Plus((_1894_v63).F28()))
					(globalState).F13 = !(p1)
					var _1913_v78 T0
					_ = _1913_v78
					var _nw313 *C3 = New_C3_()
					_ = _nw313
					_nw313.Ctor__()
					_1913_v78 = _nw313
					var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((p2), 0))
					_ = _index319
					(p2).ArraySet1((_1894_v63).F28(), (_index319).Int())
					var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((p2), 0))
					_ = _index320
					var _rhs313 T0 = _1913_v78
					_ = _rhs313
					var _rhs314 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((_1823_v9).F28(), p0), _dafny.SeqOf(p3))
					_ = _rhs314
					var _rhs315 _dafny.Int = p3
					_ = _rhs315
					var _rhs316 bool = p1
					_ = _rhs316
					var _rhs317 _dafny.Int = (_1823_v9).F28()
					_ = _rhs317
					var _lhs292 *GlobalState = globalState
					_ = _lhs292
					var _lhs293 _dafny.Array = p2
					_ = _lhs293
					var _lhs294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((p2), 0))
					_ = _lhs294
					var _lhs295 *GlobalState = globalState
					_ = _lhs295
					var _lhs296 *GlobalState = globalState
					_ = _lhs296
					_1913_v78 = _rhs313
					_lhs292.F2 = _rhs314
					(_lhs293).ArraySet1(_rhs315, (_lhs294).Int())
					_lhs295.F13 = _rhs316
					_lhs296.F7 = _rhs317
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		var _1914_v79 *C0
		_ = _1914_v79
		var _nw314 *C0 = New_C0_()
		_ = _nw314
		_nw314.Ctor__(p1)
		_1914_v79 = _nw314
		var _1915_v80 _dafny.Sequence
		_ = _1915_v80
		_1915_v80 = _dafny.SeqOf(_1914_v79)
		var _1916_v81 _dafny.Map
		_ = _1916_v81
		_1916_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1915_v80).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1903_v71).Cardinality()), _dafny.IntOfUint32((_1915_v80).Cardinality()))).Uint32()).(*C0), (_1823_v9).F28())
		r0 = (_1901_v69).Select((Companion_Default___.SafeIndex((_1916_v81).Cardinality(), _dafny.IntOfUint32((_1901_v69).Cardinality()))).Uint32()).(bool)
		r1 = ((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1894_v63).F28())).Cardinality())).Plus((_1894_v63).F28())
		var _len0_56 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_56
		var _nw315 _dafny.Array
		_ = _nw315
		if _len0_56.Cmp(_dafny.Zero) == 0 {
			_nw315 = _dafny.NewArray(_len0_56)
		} else {
			var _init56 func(_dafny.Int) _dafny.Map = (func(_1917_v79 *C0) func(_dafny.Int) _dafny.Map {
				return func(_1918_i8 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1917_v79.F27, Companion_D3_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("twmsayjd"), _dafny.CodePoint('c'), _1917_v79.F27))
				}
			})(_1914_v79)
			_ = _init56
			var _element0_56 = _init56(_dafny.Zero)
			_ = _element0_56
			_nw315 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
			(_nw315).ArraySet1(_element0_56, 0)
			var _nativeLen0_56 = (_len0_56).Int()
			_ = _nativeLen0_56
			for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
				(_nw315).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
			}
		}
		r2 = _nw315
		return r0, r1, r2
	}
}
func (_this *C7) M3(p0 _dafny.Array, globalState *GlobalState) {
	{
		var _1919_v0 bool
		_ = _1919_v0
		_1919_v0 = true
		var _1920_i0 _dafny.Int
		_ = _1920_i0
		_1920_i0 = _dafny.Zero
		{
			for !(_1919_v0) {
				{
					if (_1920_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_1920_i0 = (_1920_i0).Plus(_dafny.One)
					if _1919_v0 {
						var _1921_v1 _dafny.Array
						_ = _1921_v1
						var _len0_57 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_57
						var _nw316 _dafny.Array
						_ = _nw316
						if _len0_57.Cmp(_dafny.Zero) == 0 {
							_nw316 = _dafny.NewArray(_len0_57)
						} else {
							var _init57 func(_dafny.Int) bool = (func(_1922_v0 bool) func(_dafny.Int) bool {
								return func(_1923_i1 _dafny.Int) bool {
									return _1922_v0
								}
							})(_1919_v0)
							_ = _init57
							var _element0_57 = _init57(_dafny.Zero)
							_ = _element0_57
							_nw316 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
							(_nw316).ArraySet1(_element0_57, 0)
							var _nativeLen0_57 = (_len0_57).Int()
							_ = _nativeLen0_57
							for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
								(_nw316).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
							}
						}
						_1921_v1 = _nw316
						var _1924_v2 _dafny.Sequence
						_ = _1924_v2
						_1924_v2 = _dafny.SeqOf(_1921_v1)
						var _1925_v3 _dafny.Int
						_ = _1925_v3
						_1925_v3 = _dafny.IntOfInt64(533)
						_1921_v1 = (_1924_v2).Select((Companion_Default___.SafeIndex(_1925_v3, _dafny.IntOfUint32((_1924_v2).Cardinality()))).Uint32()).(_dafny.Array)
						var _1926_v4 _dafny.Array
						_ = _1926_v4
						var _len0_58 _dafny.Int = _dafny.IntOfInt64(20)
						_ = _len0_58
						var _nw317 _dafny.Array
						_ = _nw317
						if _len0_58.Cmp(_dafny.Zero) == 0 {
							_nw317 = _dafny.NewArray(_len0_58)
						} else {
							var _init58 func(_dafny.Int) D15 = func(_1927_i2 _dafny.Int) D15 {
								return Companion_D15_.Create_DC33_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bauwltxil")).Cardinality())))
							}
							_ = _init58
							var _element0_58 = _init58(_dafny.Zero)
							_ = _element0_58
							_nw317 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
							(_nw317).ArraySet1(_element0_58, 0)
							var _nativeLen0_58 = (_len0_58).Int()
							_ = _nativeLen0_58
							for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
								(_nw317).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
							}
						}
						_1926_v4 = _nw317
						var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1926_v4), 0))
						_ = _index321
						(_1926_v4).ArraySet1(Companion_D15_.Create_DC33_(Companion_Default___.Fm21(globalState)), (_index321).Int())
						var _1928_v5 _dafny.Sequence
						_ = _1928_v5
						_1928_v5 = _dafny.SeqOf(_1919_v0)
						var _1929_v6 _dafny.Map
						_ = _1929_v6
						_1929_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D14_.Create_DC29_(_1925_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1928_v5)), _dafny.IntOfInt64(-262))
						var _1930_v7 _dafny.Map
						_ = _1930_v7
						_1930_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1919_v0, p0)
						var _1931_v8 _dafny.Sequence
						_ = _1931_v8
						_1931_v8 = _dafny.SeqOf(_1925_v3, (_1929_v6).Cardinality(), (_1930_v7).Cardinality())
						var _1932_v9 D15
						_ = _1932_v9
						_1932_v9 = Companion_D15_.Create_DC33_(_1931_v8)
						var _pat_let_tv32 = _1931_v8
						_ = _pat_let_tv32
						var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1926_v4), 0))
						_ = _index322
						var _rhs318 bool = _1919_v0
						_ = _rhs318
						var _rhs319 bool = (_1925_v3).Cmp(_1925_v3) < 0
						_ = _rhs319
						var _rhs320 D15 = func(_pat_let26_0 D15) D15 {
							return func(_1933_dt__update__tmp_h0 D15) D15 {
								return func(_pat_let27_0 _dafny.Sequence) D15 {
									return func(_1934_dt__update_hcf44_h0 _dafny.Sequence) D15 {
										return Companion_D15_.Create_DC33_(_1934_dt__update_hcf44_h0)
									}(_pat_let27_0)
								}(_pat_let_tv32)
							}(_pat_let26_0)
						}(_1932_v9)
						_ = _rhs320
						var _rhs321 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1925_v3, Companion_Default___.SafeDivisionInt(_1925_v3, _1925_v3)))
						_ = _rhs321
						var _lhs297 *GlobalState = globalState
						_ = _lhs297
						var _lhs298 *GlobalState = globalState
						_ = _lhs298
						var _lhs299 _dafny.Array = _1926_v4
						_ = _lhs299
						var _lhs300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1926_v4), 0))
						_ = _lhs300
						var _lhs301 *GlobalState = globalState
						_ = _lhs301
						_lhs297.F13 = _rhs318
						_lhs298.F21 = _rhs319
						(_lhs299).ArraySet1(_rhs320, (_lhs300).Int())
						_lhs301.F16 = _rhs321
						var _1935_v10 _dafny.Set
						_ = _1935_v10
						_1935_v10 = _dafny.SetOf(_1925_v3)
						var _1936_v11 *C5
						_ = _1936_v11
						var _nw318 *C5 = New_C5_()
						_ = _nw318
						_nw318.Ctor__((_1935_v10).IsProperSubsetOf(_dafny.SetOf(_1925_v3)))
						_1936_v11 = _nw318
						_1919_v0 = (_1936_v11).F24()
						var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1921_v1), 0))
						_ = _index323
						(_1921_v1).ArraySet1((_1928_v5).Select((Companion_Default___.SafeIndex(_1925_v3, _dafny.IntOfUint32((_1928_v5).Cardinality()))).Uint32()).(bool), (_index323).Int())
						var _1937_v12 _dafny.Set
						_ = _1937_v12
						_1937_v12 = _dafny.SetOf(_1931_v8)
						var _1938_v13 _dafny.Sequence
						_ = _1938_v13
						_1938_v13 = _dafny.SeqOf(_1931_v8, _1931_v8, _1931_v8)
						var _1939_v15 _dafny.MultiSet
						_ = _1939_v15
						_1939_v15 = _dafny.MultiSetOf(_1919_v0, _1919_v0, (_1936_v11).F24())
						var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1921_v1), 0))
						_ = _index324
						var _rhs322 bool = (_1936_v11).F24()
						_ = _rhs322
						var _rhs323 bool = ((_1937_v12).Difference(func() _dafny.Set {
							var _coll75 = _dafny.NewBuilder()
							_ = _coll75
							for _iter87 := _dafny.Iterate((_1938_v13).Elements()); ; {
								_compr_75, _ok87 := _iter87()
								if !_ok87 {
									break
								}
								var _1940_v14 _dafny.Sequence
								_1940_v14 = interface{}(_compr_75).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_1938_v13, _1940_v14) {
									_coll75.Add(_1940_v14)
								}
							}
							return _coll75.ToSet()
						}())).IsProperSubsetOf(_dafny.SetOf(_dafny.Companion_Sequence_.Update(_1931_v8, (Companion_Default___.SafeIndex(_1925_v3, _dafny.IntOfUint32((_1931_v8).Cardinality()))).Uint32(), _1925_v3)))
						_ = _rhs323
						var _rhs324 bool = (_1939_v15).IsDisjointFrom(_1939_v15)
						_ = _rhs324
						var _lhs302 *GlobalState = globalState
						_ = _lhs302
						var _lhs303 _dafny.Array = _1921_v1
						_ = _lhs303
						var _lhs304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1921_v1), 0))
						_ = _lhs304
						var _lhs305 *GlobalState = globalState
						_ = _lhs305
						_lhs302.F2 = _rhs322
						(_lhs303).ArraySet1(_rhs323, (_lhs304).Int())
						_lhs305.F13 = _rhs324
					} else {
						var _1941_v16 *C4
						_ = _1941_v16
						var _nw319 *C4 = New_C4_()
						_ = _nw319
						_nw319.Ctor__(_1919_v0, _1919_v0)
						_1941_v16 = _nw319
						var _1942_v17 _dafny.Array
						_ = _1942_v17
						var _nw320 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(16))
						_ = _nw320
						_1942_v17 = _nw320
						var _1943_v18 _dafny.Set
						_ = _1943_v18
						_1943_v18 = _dafny.SetOf(_dafny.IntOfInt64(738))
						var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_1942_v17), 0))
						_ = _index325
						(_1942_v17).ArraySet1(_1943_v18, (_index325).Int())
						var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_1942_v17), 0))
						_ = _index326
						(_1942_v17).ArraySet1(_1943_v18, (_index326).Int())
						var _1944_v19 _dafny.Array
						_ = _1944_v19
						var _nw321 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
						_ = _nw321
						_1944_v19 = _nw321
						var _1945_v20 _dafny.CodePoint
						_ = _1945_v20
						_1945_v20 = _dafny.CodePoint('f')
						var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1944_v19), 0))
						_ = _index327
						(_1944_v19).ArraySet1CodePoint(_1945_v20, (_index327).Int())
						var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1944_v19), 0))
						_ = _index328
						(_1944_v19).ArraySet1CodePoint(_1945_v20, (_index328).Int())
						var _1946_v21 _dafny.Int
						_ = _1946_v21
						_1946_v21 = _dafny.IntOfInt64(486)
						var _1947_v22 _dafny.Map
						_ = _1947_v22
						_1947_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1946_v21, (_1944_v19).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1944_v19), 0))).Int()))
						var _1948_v23 _dafny.Sequence
						_ = _1948_v23
						_1948_v23 = _dafny.SeqOf(_1919_v0)
						var _1949_v24 _dafny.Sequence
						_ = _1949_v24
						_1949_v24 = _dafny.UnicodeSeqOfUtf8Bytes("nwdg")
						var _1950_v25 _dafny.Sequence
						_ = _1950_v25
						_1950_v25 = _dafny.SeqOf(_1949_v24, _1949_v24)
						var _1951_v26 _dafny.Map
						_ = _1951_v26
						_1951_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1950_v25).Select((Companion_Default___.SafeIndex(_1946_v21, _dafny.IntOfUint32((_1950_v25).Cardinality()))).Uint32()).(_dafny.Sequence), (_1941_v16).F26())
						var _1952_v27 *C0
						_ = _1952_v27
						var _nw322 *C0 = New_C0_()
						_ = _nw322
						_nw322.Ctor__(_1941_v16.F25)
						_1952_v27 = _nw322
						var _1953_v28 _dafny.Map
						_ = _1953_v28
						_1953_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1952_v27, _1919_v0)
						var _1954_v29 _dafny.Map
						_ = _1954_v29
						_1954_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1952_v27.F27, !(_1941_v16.F25))
						var _1955_v30 _dafny.Array
						_ = _1955_v30
						var _nwElement0_57 bool = ((_1947_v22).Cardinality()).Cmp(Companion_Default___.Fm0(_1946_v21, _1946_v21, Companion_Default___.Fm0(_1946_v21, ((_1942_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_1942_v17), 0))).Int()).(_dafny.Set)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(19))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg110 _dafny.Int) interface{} {
								return coer110(arg110)
							}
						}((func(_1956_v21 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1957_i3 _dafny.Int) _dafny.Int {
								return _1956_v21
							}
						})(_1946_v21)))).Cardinality()), (_1948_v23).Select((Companion_Default___.SafeIndex(_1946_v21, _dafny.IntOfUint32((_1948_v23).Cardinality()))).Uint32()).(bool), globalState), (_1941_v16).F26(), globalState)) <= 0
						_ = _nwElement0_57
						var _nw323 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(7))
						_ = _nw323
						(_nw323).ArraySet1(_nwElement0_57, 0)
						(_nw323).ArraySet1((func() bool {
							if (_1951_v26).Contains(_1949_v24) {
								return (_1951_v26).Get(_1949_v24).(bool)
							}
							return _1919_v0
						})(), 1)
						(_nw323).ArraySet1(!((!((func() bool {
							if (_1953_v28).Contains(_1952_v27) {
								return (_1953_v28).Get(_1952_v27).(bool)
							}
							return _1941_v16.F25
						})())) == (false)), 2)
						(_nw323).ArraySet1(Companion_Default___.Fm2(_1954_v29, globalState), 3)
						(_nw323).ArraySet1((true) && (!(false)), 4)
						(_nw323).ArraySet1((_1919_v0) == (!(_1919_v0)), 5)
						(_nw323).ArraySet1(false, 6)
						_1955_v30 = _nw323
						var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1955_v30), 0))
						_ = _index329
						(_1955_v30).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1949_v24, _1945_v20), (_index329).Int())
						var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1955_v30), 0))
						_ = _index330
						(_1955_v30).ArraySet1((_1941_v16).F26(), (_index330).Int())
						var _1958_v31 _dafny.Array
						_ = _1958_v31
						var _nw324 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
						_ = _nw324
						_1958_v31 = _nw324
						_1958_v31 = _1958_v31
					}
					var _1959_v32 _dafny.Int
					_ = _1959_v32
					_1959_v32 = _dafny.IntOfInt64(678)
					(globalState).F7 = (_dafny.Zero).Minus(_1959_v32)
					var _1960_v33 _dafny.CodePoint
					_ = _1960_v33
					_1960_v33 = _dafny.CodePoint('p')
					var _1961_v34 _dafny.Map
					_ = _1961_v34
					_1961_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1919_v0, _1919_v0)
					var _1962_v35 _dafny.Array
					_ = _1962_v35
					var _nwElement0_58 _dafny.Sequence = Companion_Default___.Fm18(true, _1960_v33, false, globalState)
					_ = _nwElement0_58
					var _nw325 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(2))
					_ = _nw325
					(_nw325).ArraySet1(_nwElement0_58, 0)
					(_nw325).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1919_v0), Companion_Default___.Fm18(true, _1960_v33, Companion_Default___.Fm2(_1961_v34, globalState), globalState)), 1)
					_1962_v35 = _nw325
					var _1963_v36 _dafny.Sequence
					_ = _1963_v36
					_1963_v36 = _dafny.SeqOf(_1919_v0, !(_1919_v0))
					var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1962_v35), 0))
					_ = _index331
					(_1962_v35).ArraySet1(_1963_v36, (_index331).Int())
					var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1962_v35), 0))
					_ = _index332
					(_1962_v35).ArraySet1(_dafny.SeqOf(_1919_v0, _1919_v0), (_index332).Int())
					if _1919_v0 {
						var _1964_v37 _dafny.Sequence
						_ = _1964_v37
						_1964_v37 = _dafny.UnicodeSeqOfUtf8Bytes("dhdknns")
						var _1965_v38 _dafny.Map
						_ = _1965_v38
						_1965_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1919_v0, (_1962_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1962_v35), 0))).Int()).(_dafny.Sequence))
						var _1966_v39 _dafny.Map
						_ = _1966_v39
						_1966_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1959_v32, _1959_v32)
						var _1967_v41 D13
						_ = _1967_v41
						_1967_v41 = Companion_D13_.Create_DC25_(_1966_v39)
						var _1968_v42 _dafny.Map
						_ = _1968_v42
						_1968_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1963_v36, (Companion_Default___.SafeIndex((func() _dafny.Map {
							var _coll76 = _dafny.NewMapBuilder()
							_ = _coll76
							for _iter88 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(368), _dafny.IntOfInt64(628))); ; {
								_compr_76, _ok88 := _iter88()
								if !_ok88 {
									break
								}
								var _1969_v40 _dafny.Int
								_1969_v40 = interface{}(_compr_76).(_dafny.Int)
								if ((_dafny.IntOfInt64(368)).Cmp(_1969_v40) <= 0) && ((_1969_v40).Cmp(_dafny.IntOfInt64(628)) < 0) {
									_coll76.Add((_1969_v40).Minus(_1959_v32), _1919_v0)
								}
							}
							return _coll76.ToMap()
						}()).Cardinality(), _dafny.IntOfUint32((_1963_v36).Cardinality()))).Uint32(), _1919_v0)).Cardinality()), _1967_v41)
						var _1970_v43 D17
						_ = _1970_v43
						_1970_v43 = Companion_D17_.Create_DC39_(_dafny.IntOfInt64(4), _1966_v39, _1919_v0, _1968_v42, _1959_v32)
						var _1971_v44 _dafny.Map
						_ = _1971_v44
						_1971_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1959_v32, _1919_v0)
						var _1972_v45 _dafny.Sequence
						_ = _1972_v45
						_1972_v45 = _dafny.SeqOf(_1971_v44)
						var _1973_v46 *C1
						_ = _1973_v46
						var _nw326 *C1 = New_C1_()
						_ = _nw326
						_nw326.Ctor__(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_dafny.IntOfUint32((_1964_v37).Cardinality()), (_1965_v38).Cardinality(), _1959_v32, (_1970_v43).Dtor_cf54(), globalState), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cfemxm")).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1971_v44), _1972_v45))
						_1973_v46 = _nw326
						var _1974_v47 _dafny.Sequence
						_ = _1974_v47
						_1974_v47 = _dafny.SeqOf(_1959_v32)
						var _1975_v48 _dafny.Array
						_ = _1975_v48
						var _len0_59 _dafny.Int = _dafny.One
						_ = _len0_59
						var _nw327 _dafny.Array
						_ = _nw327
						if _len0_59.Cmp(_dafny.Zero) == 0 {
							_nw327 = _dafny.NewArray(_len0_59)
						} else {
							var _init59 func(_dafny.Int) bool = (func(_1976_v0 bool) func(_dafny.Int) bool {
								return func(_1977_i4 _dafny.Int) bool {
									return _1976_v0
								}
							})(_1919_v0)
							_ = _init59
							var _element0_59 = _init59(_dafny.Zero)
							_ = _element0_59
							_nw327 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
							(_nw327).ArraySet1(_element0_59, 0)
							var _nativeLen0_59 = (_len0_59).Int()
							_ = _nativeLen0_59
							for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
								(_nw327).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
							}
						}
						_1975_v48 = _nw327
						var _1978_v49 _dafny.Array
						_ = _1978_v49
						_1978_v49 = _1975_v48
						var _1979_v50 _dafny.Array
						_ = _1979_v50
						var _nwElement0_59 _dafny.Array = _1975_v48
						_ = _nwElement0_59
						var _nw328 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(9))
						_ = _nw328
						(_nw328).ArraySet1(_nwElement0_59, 0)
						(_nw328).ArraySet1(_1975_v48, 1)
						(_nw328).ArraySet1(_1975_v48, 2)
						(_nw328).ArraySet1(_1975_v48, 3)
						(_nw328).ArraySet1((_1978_v49), 4)
						(_nw328).ArraySet1(_1975_v48, 5)
						(_nw328).ArraySet1(_1975_v48, 6)
						(_nw328).ArraySet1(_1975_v48, 7)
						(_nw328).ArraySet1(_1975_v48, 8)
						_1979_v50 = _nw328
						var _1980_v51 _dafny.Map
						_ = _1980_v51
						_1980_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_1973_v46.F29, _dafny.IntOfInt64(310)), _1974_v47), _1979_v50)
						_1980_v51 = (_1980_v51).Update((_1973_v46.F29).Cmp(_dafny.IntOfInt64(-382)) <= 0, _1979_v50)
						var _1981_v52 D18
						_ = _1981_v52
						_1981_v52 = Companion_D18_.Create_DC43_()
						var _1982_v53 _dafny.Map
						_ = _1982_v53
						_1982_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1959_v32, _1981_v52)
						(globalState).F2 = !(_1982_v53).Contains((func() _dafny.Int {
							if _1919_v0 {
								return _dafny.IntOfInt64(467)
							}
							return _1959_v32
						})())
						var _1983_v54 _dafny.Array
						_ = _1983_v54
						var _nw329 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
						_ = _nw329
						_1983_v54 = _nw329
						_1983_v54 = _1983_v54
						var _1984_v55 *C3
						_ = _1984_v55
						var _nw330 *C3 = New_C3_()
						_ = _nw330
						_nw330.Ctor__()
						_1984_v55 = _nw330
					} else {
						var _1985_v56 _dafny.Sequence
						_ = _1985_v56
						_1985_v56 = _dafny.UnicodeSeqOfUtf8Bytes("ap")
						var _1986_v57 D13
						_ = _1986_v57
						_1986_v57 = Companion_D13_.Create_DC26_(_1985_v56, _1919_v0, _1919_v0)
						var _1987_v58 _dafny.Map
						_ = _1987_v58
						_1987_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1986_v57).Dtor_cf33(), _1919_v0)
						_1987_v58 = (_1987_v58).Update(_1985_v56, !(true))
						var _1988_v59 _dafny.Map
						_ = _1988_v59
						_1988_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1959_v32, _1919_v0)
						var _1989_v60 _dafny.Sequence
						_ = _1989_v60
						_1989_v60 = _dafny.SeqOf(_1988_v59)
						var _1990_v61 *C1
						_ = _1990_v61
						var _nw331 *C1 = New_C1_()
						_ = _nw331
						_nw331.Ctor__(_dafny.IntOfUint32((_1985_v56).Cardinality()), _1989_v60)
						_1990_v61 = _nw331
						var _1991_v62 _dafny.Sequence
						_ = _1991_v62
						_1991_v62 = _dafny.SeqOf(_1990_v61, _1990_v61, _1990_v61, _1990_v61, _1990_v61)
						var _1992_v63 _dafny.Map
						_ = _1992_v63
						_1992_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1963_v36, _1985_v56)
						Companion_Default___.M0(_dafny.Companion_Sequence_.Equal(_1991_v62, _1991_v62), _1959_v32, (_1992_v63).Update((_1962_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1962_v35), 0))).Int()).(_dafny.Sequence), _1985_v56), (_1961_v34).Cardinality(), globalState)
						(globalState).F19 = _1985_v56
						var _1993_v64 D0
						_ = _1993_v64
						_1993_v64 = Companion_D0_.Create_DC0_(Companion_Default___.Fm2(_1961_v34, globalState))
						var _1994_v65 D0
						_ = _1994_v65
						_1994_v65 = Companion_D0_.Create_DC1_(_1993_v64)
						var _1995_v66 _dafny.MultiSet
						_ = _1995_v66
						_1995_v66 = _dafny.MultiSetOf(_1994_v65, _1994_v65, _1994_v65)
						(globalState).F16 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-782)), _1959_v32), (_dafny.Zero).Minus((func() _dafny.Int {
							if (_1995_v66).Contains(_1994_v65) {
								return (_1995_v66).Multiplicity(_1994_v65)
							}
							return _1990_v61.F29
						})()))
						var _1996_v67 _dafny.Map
						_ = _1996_v67
						_1996_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1919_v0, _1990_v61.F29)
						var _1997_v68 _dafny.Map
						_ = _1997_v68
						_1997_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm56(_1990_v61.F29, (_1996_v67).Cardinality(), _1919_v0, globalState), (_1990_v61).Fm29(globalState))
						_1997_v68 = _1997_v68
					}
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _1998_v69 _dafny.Int
		_ = _1998_v69
		_1998_v69 = _dafny.IntOfInt64(816)
		var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))
		_ = _index333
		(p0).ArraySet1(_1998_v69, (_index333).Int())
		var _1999_v70 _dafny.Array
		_ = _1999_v70
		var _nw332 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(22))
		_ = _nw332
		_1999_v70 = _nw332
		var _2000_v71 D7
		_ = _2000_v71
		_2000_v71 = Companion_D7_.Create_DC16_(_1999_v70)
		var _2001_v72 _dafny.MultiSet
		_ = _2001_v72
		_2001_v72 = _dafny.MultiSetOf(_2000_v71)
		var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))
		_ = _index334
		(p0).ArraySet1(((_2001_v72).Update(_2000_v71, Companion_Default___.Abs(_1998_v69))).Cardinality(), (_index334).Int())
		var _2002_v73 *C5
		_ = _2002_v73
		var _nw333 *C5 = New_C5_()
		_ = _nw333
		_nw333.Ctor__(_1919_v0)
		_2002_v73 = _nw333
		var _2003_v74 _dafny.MultiSet
		_ = _2003_v74
		_2003_v74 = _dafny.MultiSetOf(p0)
		var _source24 _dafny.MultiSet = _2003_v74
		_ = _source24
		var _2004___mcc_h0 _dafny.MultiSet = _source24
		_ = _2004___mcc_h0
		var _2005_cf4 _dafny.MultiSet = _2004___mcc_h0
		_ = _2005_cf4
		var _2006_v75 _dafny.CodePoint
		_ = _2006_v75
		_2006_v75 = _dafny.CodePoint('x')
		var _2007_v76 _dafny.Sequence
		_ = _2007_v76
		_2007_v76 = _dafny.SeqOf(_2006_v75)
		_1998_v69 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_1998_v69, _dafny.IntOfUint32((_2007_v76).Cardinality())), _1998_v69)
		var _2008_v77 _dafny.Array
		_ = _2008_v77
		var _nw334 _dafny.Array = _dafny.NewArrayWithValue(Companion_D19_.Default(), _dafny.IntOfInt64(22))
		_ = _nw334
		_2008_v77 = _nw334
		var _2009_v78 _dafny.Array
		_ = _2009_v78
		var _len0_60 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_60
		var _nw335 _dafny.Array
		_ = _nw335
		if _len0_60.Cmp(_dafny.Zero) == 0 {
			_nw335 = _dafny.NewArray(_len0_60)
		} else {
			var _init60 func(_dafny.Int) bool = (func(_2010_v73 *C5) func(_dafny.Int) bool {
				return func(_2011_i5 _dafny.Int) bool {
					return (_2010_v73).F24()
				}
			})(_2002_v73)
			_ = _init60
			var _element0_60 = _init60(_dafny.Zero)
			_ = _element0_60
			_nw335 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
			(_nw335).ArraySet1(_element0_60, 0)
			var _nativeLen0_60 = (_len0_60).Int()
			_ = _nativeLen0_60
			for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
				(_nw335).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
			}
		}
		_2009_v78 = _nw335
		var _2012_v79 D19
		_ = _2012_v79
		_2012_v79 = Companion_D19_.Create_DC45_(_dafny.MultiSetOf(_2009_v78))
		var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_2008_v77), 0))
		_ = _index335
		(_2008_v77).ArraySet1(_2012_v79, (_index335).Int())
		var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_2008_v77), 0))
		_ = _index336
		(_2008_v77).ArraySet1(_2012_v79, (_index336).Int())
		var _2013_v80 _dafny.MultiSet
		_ = _2013_v80
		_2013_v80 = _dafny.MultiSetOf((_2002_v73).F24())
		var _rhs325 _dafny.MultiSet = _2013_v80
		_ = _rhs325
		var _rhs326 bool = true
		_ = _rhs326
		var _lhs306 *GlobalState = globalState
		_ = _lhs306
		_2013_v80 = _rhs325
		_lhs306.F2 = _rhs326
		(globalState).F13 = (_2002_v73).F24()
		var _pat_let_tv33 = _1998_v69
		_ = _pat_let_tv33
		var _pat_let_tv34 = _2002_v73
		_ = _pat_let_tv34
		var _pat_let_tv35 = _1998_v69
		_ = _pat_let_tv35
		var _pat_let_tv36 = _1998_v69
		_ = _pat_let_tv36
		var _pat_let_tv37 = _1998_v69
		_ = _pat_let_tv37
		(globalState).F7 = func(_source25 D16) _dafny.Int {
			if _source25.Is_DC36() {
				var _2014___mcc_h1 bool = _source25.Get_().(D16_DC36).Cf47
				_ = _2014___mcc_h1
				var _2015___mcc_h2 _dafny.MultiSet = _source25.Get_().(D16_DC36).Cf48
				_ = _2015___mcc_h2
				var _2016___mcc_h3 bool = _source25.Get_().(D16_DC36).Cf49
				_ = _2016___mcc_h3
				var _2017_cf49 bool = _2016___mcc_h3
				_ = _2017_cf49
				var _2018_cf48 _dafny.MultiSet = _2015___mcc_h2
				_ = _2018_cf48
				var _2019_cf47 bool = _2014___mcc_h1
				_ = _2019_cf47
				return (_pat_let_tv33).Times(_dafny.IntOfInt64(-220))
			} else if _source25.Is_DC35() {
				var _2020___mcc_h4 _dafny.Map = _source25.Get_().(D16_DC35).Cf46
				_ = _2020___mcc_h4
				var _2021_cf46 _dafny.Map = _2020___mcc_h4
				_ = _2021_cf46
				if (_pat_let_tv34).F24() {
					return _pat_let_tv35
				} else {
					return _dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv36)).Cardinality())
				}
			} else {
				var _2022___mcc_h5 D16 = _source25.Get_().(D16_DC37).Cf50
				_ = _2022___mcc_h5
				var _2023_cf50 D16 = _2022___mcc_h5
				_ = _2023_cf50
				return ((_dafny.SetOf((_dafny.Zero).Minus(_pat_let_tv37))).Cardinality()).Minus(_dafny.IntOfInt64(405))
			}
		}(Companion_Default___.Fm57(globalState))
		var _2024_v81 _dafny.CodePoint
		_ = _2024_v81
		_2024_v81 = _dafny.CodePoint('y')
		_2024_v81 = _dafny.CodePoint('c')
	}
}
func (_this *C7) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _2025_v0 bool
		_ = _2025_v0
		_2025_v0 = true
		var _2026_v1 _dafny.Sequence
		_ = _2026_v1
		_2026_v1 = _dafny.SeqOf(_2025_v0)
		var _2027_v2 _dafny.MultiSet
		_ = _2027_v2
		_2027_v2 = _dafny.MultiSetOf(p0, p0)
		var _2028_v3 *C4
		_ = _2028_v3
		var _nw336 *C4 = New_C4_()
		_ = _nw336
		_nw336.Ctor__(((_2026_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2026_v1).Cardinality()))).Uint32()).(bool)) && (_2025_v0), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2027_v2, p0)).Contains(_2027_v2))
		_2028_v3 = _nw336
		if (_2028_v3).F26() {
			var _2029_v4 _dafny.Map
			_ = _2029_v4
			_2029_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iuyqqd")).Cardinality()), p0)
			var _2030_v5 _dafny.Map
			_ = _2030_v5
			_2030_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm45(globalState))
			var _2031_v6 _dafny.MultiSet
			_ = _2031_v6
			_2031_v6 = _dafny.MultiSetOf(true, _2028_v3.F25, _2025_v0)
			var _2032_v7 D17
			_ = _2032_v7
			_2032_v7 = Companion_D17_.Create_DC39_(p0, (_2029_v4).Merge(Companion_Default___.Fm54(globalState)), !((_2028_v3).F26()), _2030_v5, (_2031_v6).Cardinality())
			_2032_v7 = _2032_v7
			var _2033_v8 _dafny.Set
			_ = _2033_v8
			_2033_v8 = _dafny.SetOf(_dafny.IntOfInt64(-285))
			var _2034_v9 _dafny.Map
			_ = _2034_v9
			_2034_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2033_v8, !(_2028_v3.F25))
			var _2035_v10 _dafny.Array
			_ = _2035_v10
			var _nwElement0_60 bool = (_2025_v0) || ((_2028_v3).F26())
			_ = _nwElement0_60
			var _nw337 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(11))
			_ = _nw337
			(_nw337).ArraySet1(_nwElement0_60, 0)
			(_nw337).ArraySet1(false, 1)
			(_nw337).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2033_v8, p0)).Contains(_2033_v8), 2)
			(_nw337).ArraySet1((_2028_v3).F26(), 3)
			(_nw337).ArraySet1(false, 4)
			(_nw337).ArraySet1(_2028_v3.F25, 5)
			(_nw337).ArraySet1((func() bool {
				if (_2034_v9).Contains(_2033_v8) {
					return (_2034_v9).Get(_2033_v8).(bool)
				}
				return _2025_v0
			})(), 6)
			(_nw337).ArraySet1(_2028_v3.F25, 7)
			(_nw337).ArraySet1(_2028_v3.F25, 8)
			(_nw337).ArraySet1((_dafny.SetOf(p0)).IsProperSubsetOf(_2033_v8), 9)
			(_nw337).ArraySet1(_2028_v3.F25, 10)
			_2035_v10 = _nw337
			_2035_v10 = _2035_v10
			var _2036_v11 *C0
			_ = _2036_v11
			var _nw338 *C0 = New_C0_()
			_ = _nw338
			_nw338.Ctor__((_2028_v3).F26())
			_2036_v11 = _nw338
			var _2037_v12 _dafny.Map
			_ = _2037_v12
			_2037_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2036_v11, _2028_v3.F25)
			var _2038_v13 _dafny.Map
			_ = _2038_v13
			_2038_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2037_v12).Merge(_2037_v12), _2036_v11.F27)
			var _2039_v14 _dafny.Map
			_ = _2039_v14
			_2039_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2037_v12)
			var _2040_v15 _dafny.Map
			_ = _2040_v15
			_2040_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2028_v3.F25)
			var _2041_v16 _dafny.Sequence
			_ = _2041_v16
			_2041_v16 = _dafny.SeqOf(_dafny.IntOfUint32((_2026_v1).Cardinality()), (_2040_v15).Cardinality())
			var _2042_v17 _dafny.Map
			_ = _2042_v17
			_2042_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2026_v1, _2041_v16)
			var _2043_v18 D21
			_ = _2043_v18
			_2043_v18 = Companion_D21_.Create_DC52_(_2042_v17)
			var _pat_let_tv38 = _2042_v17
			_ = _pat_let_tv38
			var _2044_v19 _dafny.Map
			_ = _2044_v19
			_2044_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let28_0 D21) D21 {
				return func(_2045_dt__update__tmp_h0 D21) D21 {
					return func(_pat_let29_0 _dafny.Map) D21 {
						return func(_2046_dt__update_hcf71_h0 _dafny.Map) D21 {
							return Companion_D21_.Create_DC52_(_2046_dt__update_hcf71_h0)
						}(_pat_let29_0)
					}(_pat_let_tv38)
				}(_pat_let28_0)
			}(_2043_v18), _2036_v11)
			var _2047_v20 _dafny.Set
			_ = _2047_v20
			_2047_v20 = _dafny.SetOf((func() *C0 {
				if (_2044_v19).Contains(_2043_v18) {
					return (_2044_v19).Get(_2043_v18).(*C0)
				}
				return _2036_v11
			})(), _2036_v11)
			var _2048_v21 _dafny.Map
			_ = _2048_v21
			_2048_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2036_v11.F27, _2047_v20)
			_2038_v13 = (_2038_v13).Update((_2037_v12).Merge((func() _dafny.Map {
				if (_2039_v14).Contains(p0) {
					return (_2039_v14).Get(p0).(_dafny.Map)
				}
				return _2037_v12
			})()), ((func() _dafny.Set {
				if (_2048_v21).Contains(_2036_v11.F27) {
					return (_2048_v21).Get(_2036_v11.F27).(_dafny.Set)
				}
				return _2047_v20
			})()).IsSubsetOf(_2047_v20))
			var _2049_v22 _dafny.Array
			_ = _2049_v22
			var _nw339 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw339
			_2049_v22 = _nw339
			var _2050_v23 _dafny.Sequence
			_ = _2050_v23
			_2050_v23 = _dafny.SeqOf(_2049_v22)
			var _2051_v24 D6
			_ = _2051_v24
			_2051_v24 = Companion_D6_.Create_DC12_(_2050_v23)
			var _2052_v25 D6
			_ = _2052_v25
			_2052_v25 = Companion_D6_.Create_DC14_(_2051_v24)
			var _2053_v26 D0
			_ = _2053_v26
			_2053_v26 = Companion_D0_.Create_DC0_(_2028_v3.F25)
			var _2054_v27 _dafny.Map
			_ = _2054_v27
			_2054_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2053_v26, p0)
			var _2055_v28 _dafny.CodePoint
			_ = _2055_v28
			_2055_v28 = _dafny.CodePoint('n')
			var _2056_v29 _dafny.Set
			_ = _2056_v29
			_2056_v29 = _dafny.SetOf(_2055_v28)
			var _2057_v30 _dafny.MultiSet
			_ = _2057_v30
			_2057_v30 = _dafny.MultiSetOf(_2056_v29, _dafny.SetOf(_dafny.CodePoint('x')), _dafny.SetOf(_2055_v28))
			var _2058_v31 _dafny.MultiSet
			_ = _2058_v31
			_2058_v31 = _dafny.MultiSetOf(_2035_v10)
			var _rhs327 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(970), p0, (_2054_v27).Cardinality(), p0, (_dafny.Zero).Minus(Companion_Default___.Fm0(p0, (_dafny.Zero).Minus(p0), p0, (_this).Fm7(false, p0, _2028_v3.F25, globalState), globalState))), (p0).Plus(_dafny.IntOfInt64(691)))
			_ = _rhs327
			var _rhs328 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_2057_v30).Contains(_2056_v29) {
					return (_2057_v30).Multiplicity(_2056_v29)
				}
				return (p0).Minus(p0)
			})())
			_ = _rhs328
			var _rhs329 D6 = _2052_v25
			_ = _rhs329
			var _rhs330 _dafny.Int = p0
			_ = _rhs330
			var _rhs331 _dafny.Int = (func() _dafny.Int {
				if (_2058_v31).Contains(_2035_v10) {
					return (_2058_v31).Multiplicity(_2035_v10)
				}
				return _dafny.IntOfInt64(384)
			})()
			_ = _rhs331
			var _lhs307 *GlobalState = globalState
			_ = _lhs307
			var _lhs308 *GlobalState = globalState
			_ = _lhs308
			var _lhs309 *GlobalState = globalState
			_ = _lhs309
			var _lhs310 *GlobalState = globalState
			_ = _lhs310
			_lhs307.F21 = _rhs327
			_lhs308.F16 = _rhs328
			_2052_v25 = _rhs329
			_lhs309.F7 = _rhs330
			_lhs310.F7 = _rhs331
			var _2059_v32 _dafny.Sequence
			_ = _2059_v32
			_2059_v32 = _dafny.UnicodeSeqOfUtf8Bytes("as")
			(globalState).F0 = (func() bool {
				if !(!(_2028_v3.F25) || (_2028_v3.F25)) {
					return _dafny.Companion_Sequence_.IsPrefixOf(_2059_v32, _2059_v32)
				}
				return (_2028_v3).F26()
			})()
		} else {
			var _2060_v33 _dafny.Array
			_ = _2060_v33
			var _len0_61 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_61
			var _nw340 _dafny.Array
			_ = _nw340
			if _len0_61.Cmp(_dafny.Zero) == 0 {
				_nw340 = _dafny.NewArray(_len0_61)
			} else {
				var _init61 func(_dafny.Int) _dafny.MultiSet = (func(_2061_v2 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_2062_i0 _dafny.Int) _dafny.MultiSet {
						return _2061_v2
					}
				})(_2027_v2)
				_ = _init61
				var _element0_61 = _init61(_dafny.Zero)
				_ = _element0_61
				_nw340 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
				(_nw340).ArraySet1(_element0_61, 0)
				var _nativeLen0_61 = (_len0_61).Int()
				_ = _nativeLen0_61
				for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
					(_nw340).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
				}
			}
			_2060_v33 = _nw340
			_2060_v33 = _2060_v33
			var _2063_v34 _dafny.Array
			_ = _2063_v34
			var _nw341 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(11))
			_ = _nw341
			_2063_v34 = _nw341
			var _2064_v35 D12
			_ = _2064_v35
			_2064_v35 = Companion_D12_.Create_DC23_(p0)
			var _2065_v36 D12
			_ = _2065_v36
			_2065_v36 = Companion_D12_.Create_DC24_(_2064_v35)
			var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_2063_v34), 0))
			_ = _index337
			(_2063_v34).ArraySet1(_2065_v36, (_index337).Int())
			var _pat_let_tv39 = _2064_v35
			_ = _pat_let_tv39
			var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_2063_v34), 0))
			_ = _index338
			(_2063_v34).ArraySet1(func(_pat_let30_0 D12) D12 {
				return func(_2066_dt__update__tmp_h1 D12) D12 {
					return func(_pat_let31_0 D12) D12 {
						return func(_2067_dt__update_hcf31_h0 D12) D12 {
							return Companion_D12_.Create_DC24_(_2067_dt__update_hcf31_h0)
						}(_pat_let31_0)
					}(_pat_let_tv39)
				}(_pat_let30_0)
			}(Companion_D12_.Create_DC24_(_2064_v35)), (_index338).Int())
			var _2068_v37 *C6
			_ = _2068_v37
			var _nw342 *C6 = New_C6_()
			_ = _nw342
			_nw342.Ctor__()
			_2068_v37 = _nw342
			var _2069_v38 _dafny.Set
			_ = _2069_v38
			_2069_v38 = _dafny.SetOf(_2028_v3.F25, (_2028_v3).F26())
			(_2028_v3).M5((_2069_v38).IsDisjointFrom(_2069_v38), globalState)
			var _2070_v39 _dafny.Map
			_ = _2070_v39
			_2070_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2028_v3).F26(), (_2028_v3).F26())
			(globalState).F21 = Companion_Default___.Fm2(_2070_v39, globalState)
		}
		var _2071_v40 _dafny.CodePoint
		_ = _2071_v40
		_2071_v40 = _dafny.CodePoint('c')
		var _2072_v41 _dafny.Map
		_ = _2072_v41
		_2072_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2071_v40, !(_2025_v0))
		_2072_v41 = (_2072_v41).Update(_2071_v40, _2025_v0)
		if (p0).Cmp(_dafny.IntOfInt64(306)) == 0 {
			(globalState).F16 = _dafny.IntOfInt64(836)
			var _2073_v42 _dafny.Map
			_ = _2073_v42
			_2073_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2027_v2, _2028_v3.F25)
			var _2074_v43 _dafny.Map
			_ = _2074_v43
			_2074_v43 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2027_v2, (_2028_v3).F26())).Merge(_2073_v42)
			var _2075_v44 D12
			_ = _2075_v44
			_2075_v44 = Companion_D12_.Create_DC23_(p0)
			var _2076_v45 D12
			_ = _2076_v45
			_2076_v45 = Companion_D12_.Create_DC24_(_2075_v44)
			var _rhs332 _dafny.Map = _2074_v43
			_ = _rhs332
			var _rhs333 D12 = Companion_Default___.Fm58(p0, globalState)
			_ = _rhs333
			_2074_v43 = _rhs332
			_2076_v45 = _rhs333
			var _2077_v46 _dafny.Array
			_ = _2077_v46
			var _nw343 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw343
			_2077_v46 = _nw343
			_2077_v46 = _2077_v46
			var _2078_v47 T0
			_ = _2078_v47
			var _nw344 *C3 = New_C3_()
			_ = _nw344
			_nw344.Ctor__()
			_2078_v47 = _nw344
			var _2079_v48 _dafny.Sequence
			_ = _2079_v48
			_2079_v48 = _dafny.SeqOf(_2078_v47)
			var _2080_v49 _dafny.Sequence
			_ = _2080_v49
			_2080_v49 = _dafny.UnicodeSeqOfUtf8Bytes("ggly")
			var _2081_v50 D24
			_ = _2081_v50
			_2081_v50 = Companion_D24_.Create_DC63_(_2078_v47)
			var _2082_v51 _dafny.Array
			_ = _2082_v51
			var _nwElement0_61 T0 = (_2079_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2080_v49).Cardinality()), _dafny.IntOfUint32((_2079_v48).Cardinality()))).Uint32()).(T0)
			_ = _nwElement0_61
			var _nw345 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(18))
			_ = _nw345
			(_nw345).ArraySet1(_nwElement0_61, 0)
			(_nw345).ArraySet1(_2078_v47, 1)
			(_nw345).ArraySet1(_2078_v47, 2)
			(_nw345).ArraySet1(_2078_v47, 3)
			(_nw345).ArraySet1(_2078_v47, 4)
			(_nw345).ArraySet1(_2078_v47, 5)
			(_nw345).ArraySet1(_2078_v47, 6)
			(_nw345).ArraySet1((_2081_v50).Dtor_cf90(), 7)
			(_nw345).ArraySet1(_2078_v47, 8)
			(_nw345).ArraySet1((func() T0 {
				if _2028_v3.F25 {
					return _2078_v47
				}
				return _2078_v47
			})(), 9)
			(_nw345).ArraySet1(_2078_v47, 10)
			(_nw345).ArraySet1((_2081_v50).Dtor_cf90(), 11)
			(_nw345).ArraySet1(_2078_v47, 12)
			(_nw345).ArraySet1(_2078_v47, 13)
			(_nw345).ArraySet1(_2078_v47, 14)
			(_nw345).ArraySet1(_2078_v47, 15)
			(_nw345).ArraySet1(_2078_v47, 16)
			(_nw345).ArraySet1((func() T0 {
				if !(_2028_v3.F25) {
					return _2078_v47
				}
				return _2078_v47
			})(), 17)
			_2082_v51 = _nw345
			var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_2082_v51), 0))
			_ = _index339
			(_2082_v51).ArraySet1(_2078_v47, (_index339).Int())
			var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_2082_v51), 0))
			_ = _index340
			(_2082_v51).ArraySet1(_2078_v47, (_index340).Int())
			if !(!(false)) {
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_2077_v46), 0))
				_ = _index341
				(_2077_v46).ArraySet1(!(_2028_v3.F25) || ((_2028_v3).F26()), (_index341).Int())
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_2077_v46), 0))
				_ = _index342
				(_2077_v46).ArraySet1(!(!(_2028_v3.F25)), (_index342).Int())
				(globalState).F0 = (_dafny.Companion_Sequence_.IsPrefixOf(_2080_v49, _2080_v49)) == ((_2028_v3).F26())
				var _2083_v52 _dafny.Map
				_ = _2083_v52
				_2083_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2025_v0, p0)
				var _2084_v53 _dafny.Set
				_ = _2084_v53
				_2084_v53 = _dafny.SetOf(_dafny.IntOfUint32((_2080_v49).Cardinality()))
				var _2085_v54 _dafny.Set
				_ = _2085_v54
				_2085_v54 = _dafny.SetOf(_dafny.SetOf(p0, p0))
				var _rhs334 _dafny.Sequence = (func() _dafny.Sequence {
					if !(_2084_v53).Equals(_2084_v53) {
						return _2026_v1
					}
					return _2026_v1
				})()
				_ = _rhs334
				var _rhs335 _dafny.Int = Companion_Default___.SafeModuloInt(p0, Companion_Default___.Fm0(p0, p0, p0, (_2028_v3).F26(), globalState))
				_ = _rhs335
				var _rhs336 _dafny.Map = Companion_Default___.Fm35(_2025_v0, (_2085_v54).Cardinality(), globalState)
				_ = _rhs336
				var _rhs337 _dafny.Int = p0
				_ = _rhs337
				var _rhs338 _dafny.Sequence = _2080_v49
				_ = _rhs338
				var _lhs311 *GlobalState = globalState
				_ = _lhs311
				var _lhs312 *GlobalState = globalState
				_ = _lhs312
				var _lhs313 *GlobalState = globalState
				_ = _lhs313
				_2026_v1 = _rhs334
				_lhs311.F16 = _rhs335
				_2083_v52 = _rhs336
				_lhs312.F16 = _rhs337
				_lhs313.F8 = _rhs338
				(globalState).F7 = (_dafny.IntOfInt64(179)).Plus(p0)
				var _2086_v55 _dafny.Array
				_ = _2086_v55
				var _len0_62 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_62
				var _nw346 _dafny.Array
				_ = _nw346
				if _len0_62.Cmp(_dafny.Zero) == 0 {
					_nw346 = _dafny.NewArray(_len0_62)
				} else {
					var _init62 func(_dafny.Int) _dafny.Sequence = (func(_2087_v49 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2088_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_2087_v49, _2087_v49)
						}
					})(_2080_v49)
					_ = _init62
					var _element0_62 = _init62(_dafny.Zero)
					_ = _element0_62
					_nw346 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
					(_nw346).ArraySet1(_element0_62, 0)
					var _nativeLen0_62 = (_len0_62).Int()
					_ = _nativeLen0_62
					for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
						(_nw346).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
					}
				}
				_2086_v55 = _nw346
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_2086_v55), 0))
				_ = _index343
				(_2086_v55).ArraySet1(_dafny.SeqOf(_2080_v49), (_index343).Int())
				var _2089_v56 _dafny.Sequence
				_ = _2089_v56
				_2089_v56 = _dafny.SeqOf(_2080_v49, _2080_v49, _2080_v49, _2080_v49)
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_2086_v55), 0))
				_ = _index344
				(_2086_v55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2089_v56, Companion_Default___.Fm59((_2028_v3).F26(), globalState)), (_index344).Int())
			} else {
				var _2090_v57 _dafny.MultiSet
				_ = _2090_v57
				_2090_v57 = _dafny.MultiSetOf(_2025_v0)
				var _2091_v58 _dafny.Array
				_ = _2091_v58
				var _nwElement0_62 _dafny.Int = p0
				_ = _nwElement0_62
				var _nw347 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.One)
				_ = _nw347
				(_nw347).ArraySet1(_nwElement0_62, 0)
				_2091_v58 = _nw347
				var _2092_v59 D21
				_ = _2092_v59
				_2092_v59 = Companion_D21_.Create_DC53_(_2077_v46, false, ((_2090_v57).Update(_2025_v0, Companion_Default___.Abs(_dafny.IntOfUint32((_2080_v49).Cardinality())))).Difference(_2090_v57), _2091_v58)
				_2092_v59 = (func() D21 {
					if !((_2028_v3).F26()) || (_2028_v3.F25) {
						return _2092_v59
					}
					return _2092_v59
				})()
				(globalState).F16 = p0
				var _2093_v60 _dafny.Array
				_ = _2093_v60
				var _nw348 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw348
				_2093_v60 = _nw348
				_2093_v60 = _2093_v60
				(globalState).F8 = _dafny.UnicodeSeqOfUtf8Bytes("ydtter")
				var _2094_v61 *C0
				_ = _2094_v61
				var _nw349 *C0 = New_C0_()
				_ = _nw349
				_nw349.Ctor__(_2025_v0)
				_2094_v61 = _nw349
			}
		} else {
			var _2095_v62 _dafny.Array
			_ = _2095_v62
			var _len0_63 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_63
			var _nw350 _dafny.Array
			_ = _nw350
			if _len0_63.Cmp(_dafny.Zero) == 0 {
				_nw350 = _dafny.NewArray(_len0_63)
			} else {
				var _init63 func(_dafny.Int) _dafny.Int = (func(_2096_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2097_i2 _dafny.Int) _dafny.Int {
						return (_2097_i2).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2096_p0, _2096_p0)).Cardinality())
					}
				})(p0)
				_ = _init63
				var _element0_63 = _init63(_dafny.Zero)
				_ = _element0_63
				_nw350 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
				(_nw350).ArraySet1(_element0_63, 0)
				var _nativeLen0_63 = (_len0_63).Int()
				_ = _nativeLen0_63
				for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
					(_nw350).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
				}
			}
			_2095_v62 = _nw350
			var _2098_v63 _dafny.Sequence
			_ = _2098_v63
			var _out8 _dafny.Sequence
			_ = _out8
			_out8 = (_this).M4((_dafny.MultiSetOf(_2095_v62, _2095_v62)).Cardinality(), _2025_v0, _dafny.IntOfInt64(934), globalState)
			_2098_v63 = _out8
			var _2099_v64 _dafny.Sequence
			_ = _2099_v64
			_2099_v64 = _dafny.SeqOf(_2095_v62)
			r0 = (_2099_v64).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2099_v64).Cardinality()))).Uint32()).(_dafny.Array)
			var _2100_v65 _dafny.Map
			_ = _2100_v65
			_2100_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('v'))
			var _2101_v66 _dafny.Map
			_ = _2101_v66
			_2101_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2028_v3.F25, _2028_v3.F25)
			var _2102_v67 _dafny.Map
			_ = _2102_v67
			_2102_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2100_v65).Merge(_2100_v65), _2101_v66)
			_2102_v67 = Companion_Default___.Fm60(globalState)
			var _2103_v68 _dafny.Set
			_ = _2103_v68
			_2103_v68 = _dafny.SetOf(_2028_v3.F25)
			var _2104_v69 _dafny.Map
			_ = _2104_v69
			_2104_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2103_v68).Cardinality())
			var _2105_v70 _dafny.Sequence
			_ = _2105_v70
			_2105_v70 = _dafny.SeqOf(_2104_v69)
			var _rhs339 bool = (_2028_v3).F26()
			_ = _rhs339
			var _rhs340 _dafny.Int = (p0).Minus(p0)
			_ = _rhs340
			var _rhs341 _dafny.Int = (_dafny.IntOfUint32((_2105_v70).Cardinality())).Minus(p0)
			_ = _rhs341
			var _lhs314 *C4 = _2028_v3
			_ = _lhs314
			var _lhs315 *GlobalState = globalState
			_ = _lhs315
			var _lhs316 *GlobalState = globalState
			_ = _lhs316
			_lhs314.F25 = _rhs339
			_lhs315.F16 = _rhs340
			_lhs316.F16 = _rhs341
			var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_2095_v62), 0))
			_ = _index345
			(_2095_v62).ArraySet1(((_2027_v2).Intersection(_2027_v2)).Cardinality(), (_index345).Int())
			var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_2095_v62), 0))
			_ = _index346
			(_2095_v62).ArraySet1(p0, (_index346).Int())
		}
		var _2106_v71 _dafny.MultiSet
		_ = _2106_v71
		_2106_v71 = _dafny.MultiSetOf(_2071_v40, _2071_v40, Companion_Default___.Fm15(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-90))).Uint32(), func(coer111 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg111 _dafny.Int) interface{} {
				return coer111(arg111)
			}
		}(func(_2107_i3 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(214)
		})), p0, globalState))
		var _2108_v72 _dafny.Sequence
		_ = _2108_v72
		_2108_v72 = _dafny.UnicodeSeqOfUtf8Bytes("h")
		var _2109_v73 _dafny.MultiSet
		_ = _2109_v73
		_2109_v73 = _dafny.MultiSetOf((_2028_v3).F26())
		var _rhs342 _dafny.MultiSet = ((_dafny.MultiSetOf((_2108_v72).Select((Companion_Default___.SafeIndex((_2109_v73).Cardinality(), _dafny.IntOfUint32((_2108_v72).Cardinality()))).Uint32()).(_dafny.CodePoint))).Update(_2071_v40, Companion_Default___.Abs(p0))).Union((_2106_v71).Update(_2071_v40, Companion_Default___.Abs(_dafny.IntOfUint32((Companion_Default___.Fm1(_dafny.IntOfInt64(502), globalState)).Cardinality()))))
		_ = _rhs342
		var _rhs343 bool = _2028_v3.F25
		_ = _rhs343
		var _rhs344 _dafny.CodePoint = (func() _dafny.CodePoint {
			if !((_dafny.IntOfInt64(706)).Cmp(_dafny.IntOfInt64(212)) < 0) {
				return _dafny.CodePoint('x')
			}
			return _2071_v40
		})()
		_ = _rhs344
		var _rhs345 _dafny.Int = p0
		_ = _rhs345
		var _lhs317 *C4 = _2028_v3
		_ = _lhs317
		var _lhs318 *GlobalState = globalState
		_ = _lhs318
		_2106_v71 = _rhs342
		_lhs317.F25 = _rhs343
		_2071_v40 = _rhs344
		_lhs318.F7 = _rhs345
		var _2110_v74 _dafny.Array
		_ = _2110_v74
		var _nw351 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw351
		_2110_v74 = _nw351
		for _iter89 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2110_v74), 0))); ; {
			_guard_loop_11, _ok89 := _iter89()
			if !_ok89 {
				break
			}
			var _2111_i4 _dafny.Int
			_2111_i4 = interface{}(_guard_loop_11).(_dafny.Int)
			if (true) && (((_2111_i4).Sign() != -1) && ((_2111_i4).Cmp(_dafny.ArrayLen((_2110_v74), 0)) < 0)) {
				(_2110_v74).ArraySet1((_2111_i4).Minus(p0), (_2111_i4).Int())
			}
		}
		r0 = _2110_v74
		var _2112_v75 _dafny.Set
		_ = _2112_v75
		_2112_v75 = _dafny.SetOf(_2110_v74, _2110_v74)
		r1 = !((p0).Cmp((_2112_v75).Cardinality()) != 0)
		return r0, r1
	}
}
func (_this *C7) M4(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _2113_v0 _dafny.Map
		_ = _2113_v0
		_2113_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
		var _2114_v1 _dafny.Set
		_ = _2114_v1
		_2114_v1 = _dafny.SetOf(p1)
		var _2115_v2 _dafny.CodePoint
		_ = _2115_v2
		_2115_v2 = _dafny.CodePoint('d')
		var _2116_v3 _dafny.Sequence
		_ = _2116_v3
		_2116_v3 = _dafny.UnicodeSeqOfUtf8Bytes("yjscrou")
		var _2117_v4 _dafny.Array
		_ = _2117_v4
		var _nwElement0_63 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-831))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg112 _dafny.Int) interface{} {
				return coer112(arg112)
			}
		}(func(_2118_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))).Cardinality())
		_ = _nwElement0_63
		var _nw352 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(16))
		_ = _nw352
		(_nw352).ArraySet1(_nwElement0_63, 0)
		(_nw352).ArraySet1(p2, 1)
		(_nw352).ArraySet1((_dafny.Zero).Minus(p2), 2)
		(_nw352).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), Companion_Default___.Fm0(p2, p2, p2, p1, globalState)), 3)
		(_nw352).ArraySet1(((_2113_v0).Cardinality()).Times(p2), 4)
		(_nw352).ArraySet1(p0, 5)
		(_nw352).ArraySet1(p2, 6)
		(_nw352).ArraySet1((_dafny.Zero).Minus(p0), 7)
		(_nw352).ArraySet1(((_dafny.SetOf(false)).Union(_2114_v1)).Cardinality(), 8)
		(_nw352).ArraySet1(_dafny.IntOfInt64(359), 9)
		(_nw352).ArraySet1(p2, 10)
		(_nw352).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2115_v2, _2115_v2)).Cardinality()).Minus(_dafny.IntOfInt64(303)), 11)
		(_nw352).ArraySet1(p0, 12)
		(_nw352).ArraySet1(p0, 13)
		(_nw352).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_2116_v3).Cardinality())), 14)
		(_nw352).ArraySet1(p0, 15)
		_2117_v4 = _nw352
		var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))
		_ = _index347
		(_2117_v4).ArraySet1(p0, (_index347).Int())
		var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))
		_ = _index348
		(_2117_v4).ArraySet1(Companion_Default___.Fm0(p2, p2, p0, (func() bool {
			if p1 {
				return p1
			}
			return p1
		})(), globalState), (_index348).Int())
		var _2119_v5 _dafny.Sequence
		_ = _2119_v5
		_2119_v5 = _dafny.SeqOf(_2117_v4, _2117_v4, _2117_v4, _2117_v4, _2117_v4)
		var _2120_v6 D6
		_ = _2120_v6
		_2120_v6 = Companion_D6_.Create_DC12_(_2119_v5)
		var _pat_let_tv40 = _2117_v4
		_ = _pat_let_tv40
		_2119_v5 = (func() _dafny.Sequence {
			if p1 {
				return _dafny.Companion_Sequence_.Concatenate(_2119_v5, _2119_v5)
			}
			return (func(_pat_let32_0 D6) D6 {
				return func(_2121_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let33_0 _dafny.Sequence) D6 {
						return func(_2122_dt__update_hcf18_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC12_(_2122_dt__update_hcf18_h0)
						}(_pat_let33_0)
					}(_dafny.SeqOf(_pat_let_tv40))
				}(_pat_let32_0)
			}(_2120_v6)).Dtor_cf18()
		})()
		(globalState).F7 = (func() _dafny.Int {
			if p1 {
				return (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)
			}
			return p0
		})()
		var _2123_v7 *C6
		_ = _2123_v7
		var _nw353 *C6 = New_C6_()
		_ = _nw353
		_nw353.Ctor__()
		_2123_v7 = _nw353
		_2123_v7 = _2123_v7
		for _iter90 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2117_v4), 0))); ; {
			_guard_loop_12, _ok90 := _iter90()
			if !_ok90 {
				break
			}
			var _2124_i1 _dafny.Int
			_2124_i1 = interface{}(_guard_loop_12).(_dafny.Int)
			if (true) && (((_2124_i1).Sign() != -1) && ((_2124_i1).Cmp(_dafny.ArrayLen((_2117_v4), 0)) < 0)) {
				(_2117_v4).ArraySet1((_2124_i1).Plus(_dafny.IntOfUint32((_2116_v3).Cardinality())), (_2124_i1).Int())
			}
		}
		var _pat_let_tv41 = _2119_v5
		_ = _pat_let_tv41
		var _source26 D6 = func(_pat_let34_0 D6) D6 {
			return func(_2125_dt__update__tmp_h1 D6) D6 {
				return func(_pat_let35_0 _dafny.Sequence) D6 {
					return func(_2126_dt__update_hcf18_h1 _dafny.Sequence) D6 {
						return Companion_D6_.Create_DC12_(_2126_dt__update_hcf18_h1)
					}(_pat_let35_0)
				}(_pat_let_tv41)
			}(_pat_let34_0)
		}(_2120_v6)
		_ = _source26
		if _source26.Is_DC13() {
			var _2127___mcc_h0 _dafny.Int = _source26.Get_().(D6_DC13).Cf19
			_ = _2127___mcc_h0
			var _2128___mcc_h1 _dafny.Int = _source26.Get_().(D6_DC13).Cf20
			_ = _2128___mcc_h1
			var _2129_cf20 _dafny.Int = _2128___mcc_h1
			_ = _2129_cf20
			var _2130_cf19 _dafny.Int = _2127___mcc_h0
			_ = _2130_cf19
			var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))
			_ = _index349
			(_2117_v4).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), (_index349).Int())
			(globalState).F2 = false
			var _2131_v8 _dafny.Sequence
			_ = _2131_v8
			_2131_v8 = _dafny.SeqOf(p1, p1, true)
			var _2132_v9 _dafny.Sequence
			_ = _2132_v9
			_2132_v9 = _dafny.SeqOf(_2131_v8)
			var _2133_v10 _dafny.Sequence
			_ = _2133_v10
			_2133_v10 = _dafny.SeqOf(_dafny.IntOfUint32(((_2132_v9).Select((Companion_Default___.SafeIndex(_2130_cf19, _dafny.IntOfUint32((_2132_v9).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfInt64(-542), (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int))
			var _2134_v11 _dafny.Map
			_ = _2134_v11
			_2134_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2133_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2116_v3).Cardinality()), _dafny.IntOfUint32((_2133_v10).Cardinality()))).Uint32()).(_dafny.Int), p2)
			var _2135_v12 _dafny.Array
			_ = _2135_v12
			var _len0_64 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_64
			var _nw354 _dafny.Array
			_ = _nw354
			if _len0_64.Cmp(_dafny.Zero) == 0 {
				_nw354 = _dafny.NewArray(_len0_64)
			} else {
				var _init64 func(_dafny.Int) _dafny.CodePoint = (func(_2136_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2137_i2 _dafny.Int) _dafny.CodePoint {
						return _2136_v2
					}
				})(_2115_v2)
				_ = _init64
				var _element0_64 = _init64(_dafny.Zero)
				_ = _element0_64
				_nw354 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
				(_nw354).ArraySet1CodePoint(_element0_64, 0)
				var _nativeLen0_64 = (_len0_64).Int()
				_ = _nativeLen0_64
				for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
					(_nw354).ArraySet1CodePoint(_init64(_dafny.IntOf(_i0_64)), _i0_64)
				}
			}
			_2135_v12 = _nw354
			var _2138_v13 _dafny.Map
			_ = _2138_v13
			_2138_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2135_v12)
			_2134_v11 = (_2134_v11).Update((_2138_v13).Cardinality(), p2)
			var _2139_v14 _dafny.MultiSet
			_ = _2139_v14
			_2139_v14 = _dafny.MultiSetOf(p1, p1, p1)
			var _2140_v15 _dafny.Map
			_ = _2140_v15
			_2140_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_2139_v14).Union(_2139_v14))
			_2140_v15 = (_2140_v15).Update((func() bool {
				if p1 {
					return true
				}
				return p1
			})(), Companion_Default___.Fm52(p1, Companion_Default___.Fm0(_dafny.IntOfUint32((_2116_v3).Cardinality()), (_dafny.Zero).Minus(_2130_cf19), p2, (_2131_v8).Select((Companion_Default___.SafeIndex(_2130_cf19, _dafny.IntOfUint32((_2131_v8).Cardinality()))).Uint32()).(bool), globalState), globalState))
		} else if _source26.Is_DC12() {
			var _2141___mcc_h2 _dafny.Sequence = _source26.Get_().(D6_DC12).Cf18
			_ = _2141___mcc_h2
			var _2142_cf18 _dafny.Sequence = _2141___mcc_h2
			_ = _2142_cf18
			(globalState).F2 = (func() bool {
				if p1 {
					return p1
				}
				return p1
			})()
			var _2143_v16 _dafny.Array
			_ = _2143_v16
			var _nwElement0_64 bool = p1
			_ = _nwElement0_64
			var _nw355 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(25))
			_ = _nw355
			(_nw355).ArraySet1(_nwElement0_64, 0)
			(_nw355).ArraySet1(p1, 1)
			(_nw355).ArraySet1(p1, 2)
			(_nw355).ArraySet1(p1, 3)
			(_nw355).ArraySet1((p2).Cmp((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)) <= 0, 4)
			(_nw355).ArraySet1((p1) || (p1), 5)
			(_nw355).ArraySet1(Companion_Default___.Fm2(_2113_v0, globalState), 6)
			(_nw355).ArraySet1(p1, 7)
			(_nw355).ArraySet1(p1, 8)
			(_nw355).ArraySet1(p1, 9)
			(_nw355).ArraySet1(p1, 10)
			(_nw355).ArraySet1(p1, 11)
			(_nw355).ArraySet1(true, 12)
			(_nw355).ArraySet1(p1, 13)
			(_nw355).ArraySet1(p1, 14)
			(_nw355).ArraySet1(p1, 15)
			(_nw355).ArraySet1(p1, 16)
			(_nw355).ArraySet1(p1, 17)
			(_nw355).ArraySet1(p1, 18)
			(_nw355).ArraySet1(!(p1), 19)
			(_nw355).ArraySet1(false, 20)
			(_nw355).ArraySet1(p1, 21)
			(_nw355).ArraySet1(p1, 22)
			(_nw355).ArraySet1(p1, 23)
			(_nw355).ArraySet1(p1, 24)
			_2143_v16 = _nw355
			_2143_v16 = _2143_v16
			var _2144_v17 _dafny.Sequence
			_ = _2144_v17
			_2144_v17 = _dafny.SeqOf(_2116_v3, _2116_v3)
			var _2145_v18 _dafny.Map
			_ = _2145_v18
			_2145_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(102), p1)
			var _2146_v19 _dafny.Map
			_ = _2146_v19
			_2146_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), _2145_v18)
			var _2147_v20 _dafny.Map
			_ = _2147_v20
			_2147_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, ((func() _dafny.Map {
				if (_2146_v19).Contains(_dafny.IntOfInt64(170)) {
					return (_2146_v19).Get(_dafny.IntOfInt64(170)).(_dafny.Map)
				}
				return _2145_v18
			})()).Cardinality())).Cardinality(), (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int))
			var _2148_v21 D3
			_ = _2148_v21
			_2148_v21 = Companion_D3_.Create_DC6_(_2116_v3, _2115_v2, p1)
			var _2149_v22 _dafny.Array
			_ = _2149_v22
			var _nwElement0_65 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_2144_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.IntOfUint32((_2144_v17).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("hesfue"))
			_ = _nwElement0_65
			var _nw356 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(26))
			_ = _nw356
			(_nw356).ArraySet1(_nwElement0_65, 0)
			(_nw356).ArraySet1((func() _dafny.Sequence {
				if p1 {
					return _2116_v3
				}
				return _dafny.Companion_Sequence_.Update(_2116_v3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2116_v3).Cardinality()))).Uint32(), _2115_v2)
			})(), 1)
			(_nw356).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-666))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg113 _dafny.Int) interface{} {
					return coer113(arg113)
				}
			}((func(_2150_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2151_i3 _dafny.Int) _dafny.CodePoint {
					return _2150_v2
				}
			})(_2115_v2))), 2)
			(_nw356).ArraySet1(_2116_v3, 3)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_dafny.IntOfUint32((_2116_v3).Cardinality()), globalState), _2116_v3), 4)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2116_v3, _2116_v3), 5)
			(_nw356).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("thegxd"), 6)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2116_v3, _2116_v3), 7)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("npsshuqfe"), (Companion_Default___.SafeIndex((_2147_v20).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("npsshuqfe")).Cardinality()))).Uint32(), _2115_v2), 8)
			(_nw356).ArraySet1(_2116_v3, 9)
			(_nw356).ArraySet1(_2116_v3, 10)
			(_nw356).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cshpf"), 11)
			(_nw356).ArraySet1(_2116_v3, 12)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-731))).Uint32(), func(coer114 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg114 _dafny.Int) interface{} {
					return coer114(arg114)
				}
			}((func(_2152_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2153_i4 _dafny.Int) _dafny.CodePoint {
					return _2152_v2
				}
			})(_2115_v2))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-731))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg115 _dafny.Int) interface{} {
					return coer115(arg115)
				}
			}((func(_2154_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2155_i4 _dafny.Int) _dafny.CodePoint {
					return _2154_v2
				}
			})(_2115_v2)))).Cardinality()))).Uint32(), _2115_v2), 13)
			(_nw356).ArraySet1(_2116_v3, 14)
			(_nw356).ArraySet1(_2116_v3, 15)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2116_v3, _2116_v3), 16)
			(_nw356).ArraySet1((_2144_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kiiymoo")).Cardinality()), _dafny.IntOfUint32((_2144_v17).Cardinality()))).Uint32()).(_dafny.Sequence), 17)
			(_nw356).ArraySet1(_2116_v3, 18)
			(_nw356).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("svnce"), 19)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(p0, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm1(p0, globalState)).Cardinality()))).Uint32(), _dafny.CodePoint('d')), _dafny.UnicodeSeqOfUtf8Bytes("xkvbkkxeo")), 20)
			(_nw356).ArraySet1(_2116_v3, 21)
			(_nw356).ArraySet1((_2148_v21).Dtor_cf6(), 22)
			(_nw356).ArraySet1(_dafny.Companion_Sequence_.Update(_2116_v3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2116_v3).Cardinality()))).Uint32(), _2115_v2), 23)
			(_nw356).ArraySet1(_2116_v3, 24)
			(_nw356).ArraySet1(_2116_v3, 25)
			_2149_v22 = _nw356
			var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_2149_v22), 0))
			_ = _index350
			(_2149_v22).ArraySet1(_2116_v3, (_index350).Int())
			var _2156_v23 _dafny.Map
			_ = _2156_v23
			_2156_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Times((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.Concatenate(_2116_v3, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2116_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.IntOfUint32((_2116_v3).Cardinality()))).Uint32(), _2115_v2), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2116_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.IntOfUint32((_2116_v3).Cardinality()))).Uint32(), _2115_v2)).Cardinality()))).Uint32(), _2115_v2)))
			var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_2149_v22), 0))
			_ = _index351
			(_2149_v22).ArraySet1((func() _dafny.Sequence {
				if (_2156_v23).Contains((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)) {
					return (_2156_v23).Get((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
				}
				return _2116_v3
			})(), (_index351).Int())
			(globalState).F2 = p1
		} else {
			var _2157___mcc_h3 D6 = _source26.Get_().(D6_DC14).Cf21
			_ = _2157___mcc_h3
			var _2158_cf21 D6 = _2157___mcc_h3
			_ = _2158_cf21
			if !(p1) {
				(globalState).F20 = _dafny.Companion_Sequence_.Concatenate(_2116_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("isojrib"), _2116_v3))
				(globalState).F16 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfInt64(-215))))
				var _2159_v24 _dafny.Map
				_ = _2159_v24
				_2159_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				var _2160_v25 _dafny.Map
				_ = _2160_v25
				_2160_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(p0, p0, (_dafny.Zero).Minus(p2), !(p1), globalState), _2116_v3)
				_2159_v24 = (_2159_v24).Update((_2160_v25).Cardinality(), p1)
				var _2161_v26 _dafny.MultiSet
				_ = _2161_v26
				_2161_v26 = _dafny.MultiSetOf(p1, p1)
				var _2162_v27 _dafny.Sequence
				_ = _2162_v27
				_2162_v27 = _dafny.SeqOf(_2161_v26)
				var _2163_v28 _dafny.Map
				_ = _2163_v28
				_2163_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_2162_v27).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2162_v27).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_2161_v26), _dafny.SeqOf(p1, false))
				var _2164_v29 _dafny.Sequence
				_ = _2164_v29
				_2164_v29 = _dafny.SeqOf(p1, p1)
				_2163_v28 = (_2163_v28).Update(p1, _2164_v29)
				var _2165_v30 _dafny.Map
				_ = _2165_v30
				_2165_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2117_v4, _2123_v7)
				(globalState).F16 = Companion_Default___.Fm0((func() _dafny.Int {
					if (_2161_v26).Contains(p1) {
						return (_2161_v26).Multiplicity(p1)
					}
					return (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)
				})(), Companion_Default___.Fm0(Companion_Default___.Fm0(p2, p0, (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), true, globalState), (_2159_v24).Cardinality(), (_2165_v30).Cardinality(), p1, globalState), Companion_Default___.Fm0((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), p1, globalState), p1, globalState)
			} else {
				var _2166_v31 _dafny.MultiSet
				_ = _2166_v31
				_2166_v31 = _dafny.MultiSetOf(p1, p1)
				var _2167_v32 _dafny.Map
				_ = _2167_v32
				_2167_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2117_v4, !(_2166_v31).Contains(p1))
				_2167_v32 = (_2167_v32).Update(_2117_v4, p1)
				var _2168_v33 _dafny.Map
				_ = _2168_v33
				_2168_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(267), p2)
				var _2169_v34 _dafny.Sequence
				_ = _2169_v34
				_2169_v34 = _dafny.SeqOf((func() _dafny.Int {
					if (_2168_v33).Contains(((_2168_v33).Update((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), p0)).Cardinality()) {
						return (_2168_v33).Get(((_2168_v33).Update((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), p0)).Cardinality()).(_dafny.Int)
					}
					return p2
				})(), (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), (_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int))
				var _2170_v35 _dafny.Map
				_ = _2170_v35
				_2170_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2169_v34).Select((Companion_Default___.SafeIndex((_2114_v1).Cardinality(), _dafny.IntOfUint32((_2169_v34).Cardinality()))).Uint32()).(_dafny.Int), p1)
				_2170_v35 = (_2170_v35).Update((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), p1)
				var _2171_v36 *C4
				_ = _2171_v36
				var _nw357 *C4 = New_C4_()
				_ = _nw357
				_nw357.Ctor__(p1, p1)
				_2171_v36 = _nw357
				var _2172_v37 _dafny.Set
				_ = _2172_v37
				_2172_v37 = _dafny.SetOf(_2171_v36)
				_2172_v37 = _2172_v37
				(globalState).F16 = (_2169_v34).Select((Companion_Default___.SafeIndex((p0).Minus(p0), _dafny.IntOfUint32((_2169_v34).Cardinality()))).Uint32()).(_dafny.Int)
				var _2173_v38 _dafny.Array
				_ = _2173_v38
				var _nw358 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
				_ = _nw358
				_2173_v38 = _nw358
				var _2174_v40 _dafny.Map
				_ = _2174_v40
				_2174_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2115_v2, false)
				var _2175_v41 D22
				_ = _2175_v41
				_2175_v41 = Companion_D22_.Create_DC59_(Companion_D22_.Create_DC57_((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(767), _2171_v36.F25))
				var _2176_v42 _dafny.Map
				_ = _2176_v42
				_2176_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll77 = _dafny.NewMapBuilder()
					_ = _coll77
					for _iter91 := _dafny.Iterate((_2174_v40).Keys().Elements()); ; {
						_compr_77, _ok91 := _iter91()
						if !_ok91 {
							break
						}
						var _2177_v39 _dafny.CodePoint
						_2177_v39 = interface{}(_compr_77).(_dafny.CodePoint)
						if (_2174_v40).Contains(_2177_v39) {
							_coll77.Add(_2177_v39, (_2171_v36).F26())
						}
					}
					return _coll77.ToMap()
				}(), _2175_v41)
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_2173_v38), 0))
				_ = _index352
				(_2173_v38).ArraySet1((_2176_v42).Merge(_2176_v42), (_index352).Int())
				var _2178_v43 _dafny.Sequence
				_ = _2178_v43
				_2178_v43 = _dafny.SeqOf((_2171_v36).F26())
				var _2179_v44 _dafny.Map
				_ = _2179_v44
				_2179_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2171_v36).F26(), _2178_v43)
				var _2180_v45 D16
				_ = _2180_v45
				_2180_v45 = Companion_D16_.Create_DC36_((_2171_v36).F26(), (_2166_v31).Union(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
					if (_2179_v44).Contains((_this).Fm7(p1, _dafny.IntOfUint32((_2178_v43).Cardinality()), _2171_v36.F25, globalState)) {
						return (_2179_v44).Get((_this).Fm7(p1, _dafny.IntOfUint32((_2178_v43).Cardinality()), _2171_v36.F25, globalState)).(_dafny.Sequence)
					}
					return _2178_v43
				})())), _2171_v36.F25)
				var _2181_v46 _dafny.Array
				_ = _2181_v46
				var _nw359 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw359
				_2181_v46 = _nw359
				var _2182_v47 _dafny.Sequence
				_ = _2182_v47
				_2182_v47 = _dafny.SeqOf(_2181_v46)
				var _2183_v48 _dafny.Array
				_ = _2183_v48
				var _nwElement0_66 _dafny.Array = _2181_v46
				_ = _nwElement0_66
				var _nw360 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(26))
				_ = _nw360
				(_nw360).ArraySet1(_nwElement0_66, 0)
				(_nw360).ArraySet1(_2181_v46, 1)
				(_nw360).ArraySet1(_2181_v46, 2)
				(_nw360).ArraySet1(_2181_v46, 3)
				(_nw360).ArraySet1(_2181_v46, 4)
				(_nw360).ArraySet1(_2181_v46, 5)
				(_nw360).ArraySet1(_2181_v46, 6)
				(_nw360).ArraySet1(_2181_v46, 7)
				(_nw360).ArraySet1(_2181_v46, 8)
				(_nw360).ArraySet1(_2181_v46, 9)
				(_nw360).ArraySet1(_2181_v46, 10)
				(_nw360).ArraySet1(_2181_v46, 11)
				(_nw360).ArraySet1(_2181_v46, 12)
				(_nw360).ArraySet1(_2181_v46, 13)
				(_nw360).ArraySet1(_2181_v46, 14)
				(_nw360).ArraySet1(_2181_v46, 15)
				(_nw360).ArraySet1(_2181_v46, 16)
				(_nw360).ArraySet1(_2181_v46, 17)
				(_nw360).ArraySet1((_2182_v47).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2182_v47).Cardinality()))).Uint32()).(_dafny.Array), 18)
				(_nw360).ArraySet1(_2181_v46, 19)
				(_nw360).ArraySet1(_2181_v46, 20)
				(_nw360).ArraySet1(_2181_v46, 21)
				(_nw360).ArraySet1(_2181_v46, 22)
				(_nw360).ArraySet1(_2181_v46, 23)
				(_nw360).ArraySet1(_2181_v46, 24)
				(_nw360).ArraySet1(_2181_v46, 25)
				_2183_v48 = _nw360
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_2183_v48), 0))
				_ = _index353
				(_2183_v48).ArraySet1(_2181_v46, (_index353).Int())
				var _pat_let_tv42 = _2178_v43
				_ = _pat_let_tv42
				var _pat_let_tv43 = _2117_v4
				_ = _pat_let_tv43
				var _pat_let_tv44 = _2117_v4
				_ = _pat_let_tv44
				var _pat_let_tv45 = _2178_v43
				_ = _pat_let_tv45
				var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_2173_v38), 0))
				_ = _index354
				var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_2183_v48), 0))
				_ = _index355
				var _rhs346 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2174_v40, _2175_v41)).Merge(_2176_v42)
				_ = _rhs346
				var _rhs347 _dafny.Sequence = _2178_v43
				_ = _rhs347
				var _rhs348 _dafny.CodePoint = (_2116_v3).Select((Companion_Default___.SafeIndex((p0).Minus(p2), _dafny.IntOfUint32((_2116_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs348
				var _rhs349 D16 = func(_pat_let36_0 D16) D16 {
					return func(_2184_dt__update__tmp_h2 D16) D16 {
						return func(_pat_let37_0 bool) D16 {
							return func(_2185_dt__update_hcf49_h0 bool) D16 {
								return Companion_D16_.Create_DC36_((_2184_dt__update__tmp_h2).Dtor_cf47(), (_2184_dt__update__tmp_h2).Dtor_cf48(), _2185_dt__update_hcf49_h0)
							}(_pat_let37_0)
						}((_pat_let_tv42).Select((Companion_Default___.SafeIndex((_pat_let_tv44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_pat_let_tv43), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_pat_let_tv45).Cardinality()))).Uint32()).(bool))
					}(_pat_let36_0)
				}(Companion_D16_.Create_DC36_(!((_2171_v36).F26()), _2166_v31, _2171_v36.F25))
				_ = _rhs349
				var _rhs350 _dafny.Array = _2181_v46
				_ = _rhs350
				var _lhs319 _dafny.Array = _2173_v38
				_ = _lhs319
				var _lhs320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_2173_v38), 0))
				_ = _lhs320
				var _lhs321 _dafny.Array = _2183_v48
				_ = _lhs321
				var _lhs322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_2183_v48), 0))
				_ = _lhs322
				(_lhs319).ArraySet1(_rhs346, (_lhs320).Int())
				_2178_v43 = _rhs347
				_2115_v2 = _rhs348
				_2180_v45 = _rhs349
				(_lhs321).ArraySet1(_rhs350, (_lhs322).Int())
			}
			var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))
			_ = _index356
			(_2117_v4).ArraySet1(Companion_Default___.SafeModuloInt((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), Companion_Default___.SafeModuloInt((func() _dafny.Map {
				var _coll78 = _dafny.NewMapBuilder()
				_ = _coll78
				for _iter92 := _dafny.Iterate((_dafny.SeqOf(_2116_v3)).Elements()); ; {
					_compr_78, _ok92 := _iter92()
					if !_ok92 {
						break
					}
					var _2186_v49 _dafny.Sequence
					_2186_v49 = interface{}(_compr_78).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_2116_v3), _2186_v49) {
						_coll78.Add(_2186_v49, _2115_v2)
					}
				}
				return _coll78.ToMap()
			}()).Cardinality(), p2)), (_index356).Int())
			var _2187_v50 _dafny.MultiSet
			_ = _2187_v50
			_2187_v50 = _dafny.MultiSetOf(_2117_v4, _2117_v4)
			var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))
			_ = _index357
			(_2117_v4).ArraySet1((_dafny.Zero).Minus(((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)).Plus(((_2187_v50).Difference(_2187_v50)).Cardinality())), (_index357).Int())
			if (_2123_v7).Fm7(!(p1) || (!(!(p1))), p0, p1, globalState) {
				var _2188_v51 T0
				_ = _2188_v51
				var _nw361 *C3 = New_C3_()
				_ = _nw361
				_nw361.Ctor__()
				_2188_v51 = _nw361
				_2188_v51 = _2188_v51
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))
				_ = _index358
				(_2117_v4).ArraySet1(p0, (_index358).Int())
				var _2189_v52 _dafny.Set
				_ = _2189_v52
				_2189_v52 = _dafny.SetOf(p2, p2)
				var _2190_v53 _dafny.Map
				_ = _2190_v53
				_2190_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int)))
				var _2191_v54 _dafny.Map
				_ = _2191_v54
				_2191_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_2116_v3).Cardinality()))
				var _2192_v55 _dafny.Sequence
				_ = _2192_v55
				_2192_v55 = _dafny.SeqOf((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), (_2191_v54).Cardinality())
				var _2193_v56 _dafny.Map
				_ = _2193_v56
				_2193_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_2190_v53).Contains(p1) {
						return (_2190_v53).Get(p1).(_dafny.Int)
					}
					return (_2192_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2116_v3).Cardinality()), _dafny.IntOfUint32((_2192_v55).Cardinality()))).Uint32()).(_dafny.Int)
				})(), !(p1))
				_2189_v52 = Companion_Default___.Fm26(_2115_v2, _2193_v56, globalState)
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))
				_ = _index359
				(_2117_v4).ArraySet1(((Companion_Default___.Fm61(_dafny.IntOfInt64(748), _2115_v2, _2115_v2, globalState)).Update(((_2113_v0).Update(true, p1)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Update(true, p1)), Companion_Default___.Abs(p0))).Cardinality(), (_index359).Int())
				var _2194_v57 _dafny.MultiSet
				_ = _2194_v57
				_2194_v57 = _dafny.MultiSetOf(p0)
				var _2195_v58 _dafny.Map
				_ = _2195_v58
				_2195_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2113_v0)
				var _2196_v59 D23
				_ = _2196_v59
				_2196_v59 = Companion_D23_.Create_DC60_(_2195_v58)
				var _2197_v60 _dafny.Set
				_ = _2197_v60
				_2197_v60 = _dafny.SetOf(_2196_v59, _2196_v59, _2196_v59)
				var _2198_v61 *C2
				_ = _2198_v61
				var _nw362 *C2 = New_C2_()
				_ = _nw362
				_nw362.Ctor__(((_2197_v60).Union(_2197_v60)).Cardinality())
				_2198_v61 = _nw362
				var _rhs351 _dafny.MultiSet = (_2194_v57).Difference((_2194_v57).Union(_2194_v57))
				_ = _rhs351
				var _rhs352 bool = !(!(p1))
				_ = _rhs352
				var _rhs353 _dafny.MultiSet = (_2194_v57).Difference((Companion_D25_.Create_DC67_(_2194_v57)).Dtor_cf98())
				_ = _rhs353
				var _rhs354 *C2 = _2198_v61
				_ = _rhs354
				var _lhs323 *GlobalState = globalState
				_ = _lhs323
				_2194_v57 = _rhs351
				_lhs323.F21 = _rhs352
				_2194_v57 = _rhs353
				_2198_v61 = _rhs354
			} else {
				_2115_v2 = _2115_v2
				_2113_v0 = (_2113_v0).Update(p1, p1)
				var _2199_v62 _dafny.Array
				_ = _2199_v62
				var _nw363 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
				_ = _nw363
				_2199_v62 = _nw363
				var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_2199_v62), 0))
				_ = _index360
				(_2199_v62).ArraySet1(_2117_v4, (_index360).Int())
				var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_2199_v62), 0))
				_ = _index361
				(_2199_v62).ArraySet1(_2117_v4, (_index361).Int())
				(globalState).F13 = p1
				var _2200_v63 _dafny.Array
				_ = _2200_v63
				var _nw364 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw364
				_2200_v63 = _nw364
				_2200_v63 = _2200_v63
			}
		}
		r0 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if p1 {
				return _2116_v3
			}
			return _dafny.Companion_Sequence_.Update(_2116_v3, (Companion_Default___.SafeIndex((_2117_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_2117_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2116_v3).Cardinality()))).Uint32(), _dafny.CodePoint('n'))
		})(), _2116_v3)
		return r0
	}
}
func (_this *C7) F23() _dafny.Map {
	{
		return _this._f23
	}
}

// End of class C7
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
