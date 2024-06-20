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
	return _dafny.IntOfInt64(-214)
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.MultiSetOf(_dafny.IntOfInt64(420), _dafny.IntOfUint32((_dafny.SeqOf(false, (Companion_D0_.Create_DC2_(_dafny.MultiSetOf(_dafny.CodePoint('k'), _dafny.CodePoint('h')), !(!(true)))).Dtor_cf6())).Cardinality()), _dafny.IntOfInt64(630), _dafny.IntOfInt64(-203), (_dafny.Zero).Minus((_dafny.MultiSetOf(false, true)).Cardinality()))).Cardinality())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-421))))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.IntOfInt64(-770)).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-551), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("apvmft")).Cardinality()))))) >= 0
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	if (func() bool {
		if true {
			return false
		}
		return false
	})() {
		return _dafny.SeqOf(_dafny.IntOfInt64(811))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(552))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(927), (_dafny.Zero).Minus((_dafny.MultiSetOf(true, false)).Cardinality()))).Cardinality())
		})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-559), _dafny.IntOfInt64(-695), _dafny.IntOfInt64(903))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	var _source0 D10 = Companion_D10_.Create_DC29_(false, _dafny.CodePoint('f'))
	_ = _source0
	if _source0.Is_DC29() {
		var _1___mcc_h0 bool = _source0.Get_().(D10_DC29).Cf42
		_ = _1___mcc_h0
		var _2___mcc_h1 _dafny.CodePoint = _source0.Get_().(D10_DC29).Cf43
		_ = _2___mcc_h1
		var _3_cf43 _dafny.CodePoint = _2___mcc_h1
		_ = _3_cf43
		var _4_cf42 bool = _1___mcc_h0
		_ = _4_cf42
		return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-730)), _dafny.SeqOf(_dafny.IntOfInt64(14), _dafny.IntOfInt64(18))))
	} else if _source0.Is_DC30() {
		var _5___mcc_h2 _dafny.Int = _source0.Get_().(D10_DC30).Cf44
		_ = _5___mcc_h2
		var _6___mcc_h3 _dafny.Array = _source0.Get_().(D10_DC30).Cf45
		_ = _6___mcc_h3
		var _7___mcc_h4 bool = _source0.Get_().(D10_DC30).Cf46
		_ = _7___mcc_h4
		var _8___mcc_h5 _dafny.Int = _source0.Get_().(D10_DC30).Cf47
		_ = _8___mcc_h5
		var _9___mcc_h6 _dafny.Int = _source0.Get_().(D10_DC30).Cf48
		_ = _9___mcc_h6
		var _10_cf48 _dafny.Int = _9___mcc_h6
		_ = _10_cf48
		var _11_cf47 _dafny.Int = _8___mcc_h5
		_ = _11_cf47
		var _12_cf46 bool = _7___mcc_h4
		_ = _12_cf46
		var _13_cf45 _dafny.Array = _6___mcc_h3
		_ = _13_cf45
		var _14_cf44 _dafny.Int = _5___mcc_h2
		_ = _14_cf44
		if !(_12_cf46) {
			return _dafny.MultiSetOf(_10_cf48, _dafny.IntOfInt64(212), _11_cf47)
		} else {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_12_cf46, _12_cf46)).Cardinality()), _dafny.IntOfInt64(848)))
		}
	} else {
		var _15___mcc_h7 _dafny.Map = _source0.Get_().(D10_DC28).Cf41
		_ = _15___mcc_h7
		var _16_cf41 _dafny.Map = _15___mcc_h7
		_ = _16_cf41
		return _dafny.MultiSetOf(_dafny.IntOfInt64(723))
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Difference(_dafny.SetOf(!(false)))).Intersection((_dafny.SetOf(true, true, true, true, false)).Union(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 D1, p2 D0, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return (func() _dafny.CodePoint {
			if true {
				return _dafny.CodePoint('k')
			}
			return _dafny.CodePoint('a')
		})()
	}))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	var _source1 D13 = Companion_D13_.Create_DC38_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_ = _source1
	if _source1.Is_DC39() {
		var _18___mcc_h0 _dafny.Map = _source1.Get_().(D13_DC39).Cf69
		_ = _18___mcc_h0
		var _19_cf69 _dafny.Map = _18___mcc_h0
		_ = _19_cf69
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-396)), _dafny.SeqOf(_dafny.IntOfInt64(367), _dafny.IntOfInt64(184), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(838), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-492))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_20_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-647)
		}))).Cardinality())))).Cardinality()))
	} else if _source1.Is_DC40() {
		var _21___mcc_h1 D8 = _source1.Get_().(D13_DC40).Cf70
		_ = _21___mcc_h1
		var _22___mcc_h2 _dafny.Int = _source1.Get_().(D13_DC40).Cf71
		_ = _22___mcc_h2
		var _23_cf71 _dafny.Int = _22___mcc_h2
		_ = _23_cf71
		var _24_cf70 D8 = _21___mcc_h1
		_ = _24_cf70
		return _dafny.SeqOf(_23_cf71)
	} else if _source1.Is_DC38() {
		var _25___mcc_h3 _dafny.Sequence = _source1.Get_().(D13_DC38).Cf68
		_ = _25___mcc_h3
		var _26_cf68 _dafny.Sequence = _25___mcc_h3
		_ = _26_cf68
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(557))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_27_i1 _dafny.Int) _dafny.Int {
			return (_dafny.SetOf(true)).Cardinality()
		}))
	} else {
		var _28___mcc_h4 D13 = _source1.Get_().(D13_DC41).Cf72
		_ = _28___mcc_h4
		var _29_cf72 D13 = _28___mcc_h4
		_ = _29_cf72
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-849))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_30_i2 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), _dafny.IntOfInt64(603))).Cardinality(), _dafny.IntOfInt64(497))).Cardinality()
		})), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), true))).Cardinality())), _dafny.IntOfInt64(212)))
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if (_dafny.IntOfInt64(331)).Cmp((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ao")).Cardinality())))).Cardinality()) == 0 {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), !(false))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-592))))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(104), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-845), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdjmrl")).Cardinality()))).Cardinality())).Cardinality()))).Cardinality(), _dafny.IntOfInt64(371))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Intersection(_dafny.MultiSetOf(true, !(true)))).Difference(_dafny.MultiSetOf(false))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return ((func() D16 {
		if !(false) {
			return Companion_D16_.Create_DC46_(_dafny.UnicodeSeqOfUtf8Bytes("wrs"), _dafny.SetOf(false), false)
		}
		return Companion_D16_.Create_DC46_(_dafny.UnicodeSeqOfUtf8Bytes("kbcxn"), _dafny.SetOf(false), true)
	})()).Dtor_cf79()
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(_dafny.SeqOf(false, false, !(!(false))), _dafny.SeqOf(true))
		}
		return _dafny.MultiSetOf(_dafny.SeqOf(true), _dafny.SeqOf(true))
	})()).Difference((_dafny.MultiSetOf((Companion_D6_.Create_DC17_(_dafny.SeqOf(true, true, false, !(false)))).Dtor_cf23(), _dafny.SeqOf(false, false), _dafny.SeqOf(false, true))).Intersection(_dafny.MultiSetOf(_dafny.SeqOf(false, false, false, false))))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), true)
}
func (_static *CompanionStruct_Default___) Fm18(p0 D0, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf((Companion_D8_.Create_DC23_(true, !(false), false, _dafny.CodePoint('x'))).Dtor_cf35())).Difference(_dafny.SetOf(_dafny.CodePoint('s')))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(513), (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()), _dafny.IntOfInt64(420), _dafny.IntOfInt64(698)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qa")).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-314)))).Cardinality())).Cardinality())).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(48), _dafny.IntOfInt64(460))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("lt")
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	if true {
		return _dafny.SetOf(Companion_D0_.Create_DC3_(), Companion_D0_.Create_DC3_(), Companion_D0_.Create_DC3_())
	} else {
		return _dafny.SetOf(Companion_D0_.Create_DC3_())
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true, true)).Difference(_dafny.SetOf(!(true)))).Union((_dafny.SetOf(!(true))).Difference(_dafny.SetOf(true, false, true)))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(210), _dafny.IntOfInt64(454))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(417))).Uint32(), func(coer5 func(_dafny.Int) D8) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_31_i0 _dafny.Int) D8 {
		return Companion_D8_.Create_DC23_(!(false), true, false, _dafny.CodePoint('i'))
	}))).Cardinality()), _dafny.IntOfInt64(794)))).Merge((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(134), _dafny.IntOfInt64(361))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _32_v0 _dafny.Int
			_32_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(134)).Cmp(_32_v0) <= 0) && ((_32_v0).Cmp(_dafny.IntOfInt64(361)) < 0) {
				_coll0.Add((_32_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vu")).Cardinality())))
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(666), _dafny.IntOfInt64(-331))))
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('h')
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(249))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_33_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(_dafny.IntOfInt64(428))
		}))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _34_v0 _dafny.Sequence
			_34_v0 = interface{}(_compr_1).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(249))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_33_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_dafny.IntOfInt64(428))
			})), _34_v0) {
				_coll1.Add(_34_v0)
			}
		}
		return _coll1.ToSet()
	}()).IsDisjointFrom(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-168))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_35_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false)).Cardinality()
	}))).Cardinality()), (_dafny.Zero).Minus((_dafny.SetOf(false, (Companion_D16_.Create_DC46_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_36_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	})), _dafny.SetOf(false), false)).Dtor_cf81())).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC24_((func() D8 {
		if true {
			return Companion_D8_.Create_DC23_(false, !(false), !(true), _dafny.CodePoint('f'))
		}
		return Companion_D8_.Create_DC23_(true, true, true, _dafny.CodePoint('t'))
	})())
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(186), false)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Cardinality(), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ogslplvme")).Cardinality()), true)))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	if ((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, false, false), true)).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _37_v0 _dafny.Set
			_37_v0 = interface{}(_compr_2).(_dafny.Set)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, false, false), true)).Contains(_37_v0) {
				_coll2.Add(_37_v0, _dafny.IntOfUint32((_dafny.SeqOf(true, true, false)).Cardinality()))
			}
		}
		return _coll2.ToMap()
	}()).Cardinality()).Cmp((Companion_D11_.Create_DC33_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfInt64(-513))).Dtor_cf54()) < 0 {
		return (func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))), _dafny.CodePoint('m'))).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _38_v1 _dafny.Sequence
				_38_v1 = interface{}(_compr_3).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))), _dafny.CodePoint('m'))).Contains(_38_v1) {
					_coll3.Add(_38_v1, _dafny.UnicodeSeqOfUtf8Bytes("mkjgpgh"))
				}
			}
			return _coll3.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-199))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(659))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_39_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		}))))
	} else {
		return (func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-778), _dafny.IntOfInt64(723)))).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _40_v2 _dafny.Sequence
				_40_v2 = interface{}(_compr_4).(_dafny.Sequence)
				if (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-778), _dafny.IntOfInt64(723)))).Contains(_40_v2) {
					_coll4.Add(_40_v2, _dafny.UnicodeSeqOfUtf8Bytes("ivuqtjs"))
				}
			}
			return _coll4.ToMap()
		}()).Merge(func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(196)))).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _41_v3 _dafny.Sequence
				_41_v3 = interface{}(_compr_5).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(196))), _41_v3) {
					_coll5.Add(_41_v3, _dafny.UnicodeSeqOfUtf8Bytes("qpbh"))
				}
			}
			return _coll5.ToMap()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) D6 {
	if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('g'))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('i'))) {
		return Companion_D6_.Create_DC18_(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("a")).Cardinality()))
	} else {
		return Companion_D6_.Create_DC18_(!(true), _dafny.IntOfInt64(760))
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D12 {
	return Companion_D12_.Create_DC37_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf((Companion_D12_.Create_DC36_(false, false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(486))).Cardinality()))).Dtor_cf61(), !(true))).Cardinality())).Cardinality(), false)).Cardinality(), _dafny.IntOfInt64(15))).Cardinality()).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC3_(), (_dafny.SetOf(true, !(false))).Cardinality())).Cardinality()) != 0, _dafny.IntOfInt64(713), _dafny.IntOfInt64(217))
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Set {
	var _source2 D12 = Companion_D12_.Create_DC36_(false, false, (Companion_D12_.Create_DC37_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true), true, _dafny.IntOfInt64(527), _dafny.IntOfInt64(89))).Dtor_cf67())
	_ = _source2
	if _source2.Is_DC35() {
		var _42___mcc_h0 _dafny.Array = _source2.Get_().(D12_DC35).Cf56
		_ = _42___mcc_h0
		var _43___mcc_h1 bool = _source2.Get_().(D12_DC35).Cf57
		_ = _43___mcc_h1
		var _44___mcc_h2 _dafny.Sequence = _source2.Get_().(D12_DC35).Cf58
		_ = _44___mcc_h2
		var _45___mcc_h3 _dafny.Sequence = _source2.Get_().(D12_DC35).Cf59
		_ = _45___mcc_h3
		var _46___mcc_h4 _dafny.Int = _source2.Get_().(D12_DC35).Cf60
		_ = _46___mcc_h4
		var _47_cf60 _dafny.Int = _46___mcc_h4
		_ = _47_cf60
		var _48_cf59 _dafny.Sequence = _45___mcc_h3
		_ = _48_cf59
		var _49_cf58 _dafny.Sequence = _44___mcc_h2
		_ = _49_cf58
		var _50_cf57 bool = _43___mcc_h1
		_ = _50_cf57
		var _51_cf56 _dafny.Array = _42___mcc_h0
		_ = _51_cf56
		return _dafny.SetOf(_dafny.IntOfInt64(401))
	} else if _source2.Is_DC36() {
		var _52___mcc_h5 bool = _source2.Get_().(D12_DC36).Cf61
		_ = _52___mcc_h5
		var _53___mcc_h6 bool = _source2.Get_().(D12_DC36).Cf62
		_ = _53___mcc_h6
		var _54___mcc_h7 _dafny.Int = _source2.Get_().(D12_DC36).Cf63
		_ = _54___mcc_h7
		var _55_cf63 _dafny.Int = _54___mcc_h7
		_ = _55_cf63
		var _56_cf62 bool = _53___mcc_h6
		_ = _56_cf62
		var _57_cf61 bool = _52___mcc_h5
		_ = _57_cf61
		return _dafny.SetOf((_dafny.SetOf(_56_cf62, _57_cf61)).Cardinality())
	} else if _source2.Is_DC37() {
		var _58___mcc_h8 _dafny.Map = _source2.Get_().(D12_DC37).Cf64
		_ = _58___mcc_h8
		var _59___mcc_h9 bool = _source2.Get_().(D12_DC37).Cf65
		_ = _59___mcc_h9
		var _60___mcc_h10 _dafny.Int = _source2.Get_().(D12_DC37).Cf66
		_ = _60___mcc_h10
		var _61___mcc_h11 _dafny.Int = _source2.Get_().(D12_DC37).Cf67
		_ = _61___mcc_h11
		var _62_cf67 _dafny.Int = _61___mcc_h11
		_ = _62_cf67
		var _63_cf66 _dafny.Int = _60___mcc_h10
		_ = _63_cf66
		var _64_cf65 bool = _59___mcc_h9
		_ = _64_cf65
		var _65_cf64 _dafny.Map = _58___mcc_h8
		_ = _65_cf64
		return _dafny.SetOf(_62_cf67)
	} else {
		var _66___mcc_h12 _dafny.Array = _source2.Get_().(D12_DC34).Cf55
		_ = _66___mcc_h12
		var _67_cf55 _dafny.Array = _66___mcc_h12
		_ = _67_cf55
		return (_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, (Companion_D16_.Create_DC46_(_dafny.UnicodeSeqOfUtf8Bytes("yoa"), _dafny.SetOf(false, true), false)).Dtor_cf81())).Cardinality())), _dafny.IntOfInt64(725))).Union((Companion_D4_.Create_DC11_(func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(784), _dafny.IntOfInt64(-950))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _68_v0 _dafny.Int
				_68_v0 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(784)).Cmp(_68_v0) <= 0) && ((_68_v0).Cmp(_dafny.IntOfInt64(-950)) < 0) {
					_coll6.Add((_68_v0).Minus((_dafny.MultiSetOf(_dafny.CodePoint('o'), _dafny.CodePoint('e'), _dafny.CodePoint('d'), _dafny.CodePoint('u'), _dafny.CodePoint('t'))).Cardinality()))
				}
			}
			return _coll6.ToSet()
		}())).Dtor_cf18())
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _69_v0 _dafny.Array
	_ = _69_v0
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(26)
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.CodePoint = func(_70_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw0).ArraySet1CodePoint(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw0).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_69_v0 = _nw0
	var _71_globalState *GlobalState
	_ = _71_globalState
	var _nw1 *GlobalState = New_GlobalState_()
	_ = _nw1
	_nw1.Ctor__(_69_v0, _dafny.IntOfInt64(-562), false)
	_71_globalState = _nw1
	var _72_v1 bool
	_ = _72_v1
	_72_v1 = false
	if _72_v1 {
		var _73_v2 _dafny.Array
		_ = _73_v2
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_1
		var _nw2 _dafny.Array
		_ = _nw2
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw2 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_74_v1 bool) func(_dafny.Int) bool {
				return func(_75_i1 _dafny.Int) bool {
					return _74_v1
				}
			})(_72_v1)
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
		_73_v2 = _nw2
		var _76_v3 _dafny.Int
		_ = _76_v3
		_76_v3 = _dafny.IntOfInt64(-419)
		var _77_v4 _dafny.Map
		_ = _77_v4
		_77_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v3, _73_v2)
		_73_v2 = (func() _dafny.Array {
			if (_77_v4).Contains(_76_v3) {
				return (_77_v4).Get(_76_v3).(_dafny.Array)
			}
			return _73_v2
		})()
		_76_v3 = _dafny.IntOfInt64(587)
		(_71_globalState).F2 = _72_v1
		var _78_v5 _dafny.CodePoint
		_ = _78_v5
		_78_v5 = _dafny.CodePoint('v')
		_78_v5 = _78_v5
		var _79_v6 _dafny.MultiSet
		_ = _79_v6
		_79_v6 = _dafny.MultiSetOf(_78_v5, _78_v5)
		var _80_v7 _dafny.Sequence
		_ = _80_v7
		_80_v7 = _dafny.SeqOf(_dafny.IntOfInt64(43), _dafny.IntOfInt64(206), (func() _dafny.Int {
			if _72_v1 {
				return _76_v3
			}
			return _dafny.IntOfInt64(862)
		})(), ((_79_v6).Update(_78_v5, Companion_Default___.Abs(_76_v3))).Cardinality(), _76_v3)
		var _81_v8 _dafny.Map
		_ = _81_v8
		_81_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(349), _72_v1)
		var _82_v9 D0
		_ = _82_v9
		_82_v9 = Companion_D0_.Create_DC2_(_79_v6, _72_v1)
		var _83_v10 _dafny.Map
		_ = _83_v10
		_83_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_81_v8).Cardinality(), (_82_v9).Dtor_cf6())
		var _84_v11 _dafny.Map
		_ = _84_v11
		_84_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v1, _81_v8)
		_80_v7 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_80_v7, _80_v7), _dafny.SeqOf(_76_v3, (_84_v11).Cardinality(), _76_v3, _dafny.IntOfInt64(140), _dafny.IntOfInt64(-53)))
	} else {
		var _85_v12 _dafny.Array
		_ = _85_v12
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw3
		_85_v12 = _nw3
		var _86_v13 _dafny.Int
		_ = _86_v13
		_86_v13 = _dafny.IntOfInt64(744)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))
		_ = _index0
		(_85_v12).ArraySet1(_86_v13, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))
		_ = _index1
		(_85_v12).ArraySet1(Companion_Default___.Fm0(_71_globalState), (_index1).Int())
		var _87_v14 _dafny.Map
		_ = _87_v14
		_87_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v13, Companion_Default___.Fm0(_71_globalState))
		var _88_v15 _dafny.Sequence
		_ = _88_v15
		_88_v15 = _dafny.UnicodeSeqOfUtf8Bytes("t")
		var _89_v16 _dafny.Map
		_ = _89_v16
		_89_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v13, _88_v15)
		var _90_v17 _dafny.Sequence
		_ = _90_v17
		_90_v17 = _dafny.SeqOf(_86_v13)
		var _91_v18 _dafny.Map
		_ = _91_v18
		_91_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int), (_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_89_v16).Cardinality(), (_90_v17).Select((Companion_Default___.SafeIndex((_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_90_v17).Cardinality()))).Uint32()).(_dafny.Int)))
		_87_v14 = (func() _dafny.Map {
			if (_91_v18).Contains(_86_v13) {
				return (_91_v18).Get(_86_v13).(_dafny.Map)
			}
			return (_87_v14).Update(_dafny.IntOfInt64(-505), _86_v13)
		})()
		var _92_v19 _dafny.Sequence
		_ = _92_v19
		_92_v19 = _dafny.SeqOf(_85_v12, _85_v12, _85_v12)
		var _93_v20 _dafny.Sequence
		_ = _93_v20
		_93_v20 = _dafny.SeqOf(!(_72_v1))
		var _94_v21 _dafny.Array
		_ = _94_v21
		var _nwElement0_0 _dafny.Array = _85_v12
		_ = _nwElement0_0
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(16))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_0, 0)
		(_nw4).ArraySet1(_85_v12, 1)
		(_nw4).ArraySet1((func() _dafny.Array {
			if _72_v1 {
				return _85_v12
			}
			return (_92_v19).Select((Companion_Default___.SafeIndex(_86_v13, _dafny.IntOfUint32((_92_v19).Cardinality()))).Uint32()).(_dafny.Array)
		})(), 2)
		(_nw4).ArraySet1(_85_v12, 3)
		(_nw4).ArraySet1(_85_v12, 4)
		(_nw4).ArraySet1(_85_v12, 5)
		(_nw4).ArraySet1(_85_v12, 6)
		(_nw4).ArraySet1(_85_v12, 7)
		(_nw4).ArraySet1(_85_v12, 8)
		(_nw4).ArraySet1(_85_v12, 9)
		(_nw4).ArraySet1(_85_v12, 10)
		(_nw4).ArraySet1((_92_v19).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_93_v20).Cardinality()), _dafny.IntOfUint32((_92_v19).Cardinality()))).Uint32()).(_dafny.Array), 11)
		(_nw4).ArraySet1(_85_v12, 12)
		(_nw4).ArraySet1(_85_v12, 13)
		(_nw4).ArraySet1(_85_v12, 14)
		(_nw4).ArraySet1(_85_v12, 15)
		_94_v21 = _nw4
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_94_v21), 0))
		_ = _index2
		(_94_v21).ArraySet1(_85_v12, (_index2).Int())
		var _95_v22 _dafny.Map
		_ = _95_v22
		_95_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int), false)
		var _96_v23 _dafny.CodePoint
		_ = _96_v23
		_96_v23 = _dafny.CodePoint('m')
		var _97_v24 _dafny.Sequence
		_ = _97_v24
		_97_v24 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_88_v15, (Companion_Default___.SafeIndex((_95_v22).Cardinality(), _dafny.IntOfUint32((_88_v15).Cardinality()))).Uint32(), _96_v23))
		var _98_v25 _dafny.Set
		_ = _98_v25
		_98_v25 = _dafny.SetOf(Companion_Default___.Fm0(_71_globalState), (_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int), _86_v13)
		var _99_v26 _dafny.MultiSet
		_ = _99_v26
		_99_v26 = _dafny.MultiSetOf(_72_v1)
		var _100_v27 _dafny.Map
		_ = _100_v27
		_100_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_99_v26).IsDisjointFrom(_dafny.MultiSetFromSeq(Companion_Default___.Fm1(_96_v23, _72_v1, _71_globalState))), !(Companion_Default___.Fm2(_86_v13, _71_globalState)))
		var _101_v28 _dafny.Array
		_ = _101_v28
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_2
		var _nw5 _dafny.Array
		_ = _nw5
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw5 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_102_v1 bool) func(_dafny.Int) bool {
				return func(_103_i2 _dafny.Int) bool {
					return _102_v1
				}
			})(_72_v1)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw5 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw5).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw5).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_101_v28 = _nw5
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_94_v21), 0))
		_ = _index3
		var _rhs0 bool = (_98_v25).Contains((_dafny.Zero).Minus((_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int)))
		_ = _rhs0
		var _rhs1 _dafny.Array = _85_v12
		_ = _rhs1
		var _rhs2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_97_v24, _dafny.Companion_Sequence_.Concatenate(_97_v24, _97_v24))
		_ = _rhs2
		var _rhs3 bool = (func() bool {
			if (_100_v27).Contains(!((Companion_Default___.Fm0(_71_globalState)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_99_v26).Cardinality(), _101_v28)).Cardinality()) != 0)) {
				return (_100_v27).Get(!((Companion_Default___.Fm0(_71_globalState)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_99_v26).Cardinality(), _101_v28)).Cardinality()) != 0)).(bool)
			}
			return _72_v1
		})()
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _88_v15
		_ = _rhs4
		var _lhs0 _dafny.Array = _94_v21
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_94_v21), 0))
		_ = _lhs1
		_72_v1 = _rhs0
		(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
		_97_v24 = _rhs2
		_72_v1 = _rhs3
		_88_v15 = _rhs4
		var _104_v29 *C1
		_ = _104_v29
		var _nw6 *C1 = New_C1_()
		_ = _nw6
		_nw6.Ctor__(_86_v13, Companion_Default___.SafeDivisionInt((_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int), _86_v13), Companion_Default___.Fm2(_dafny.IntOfInt64(132), _71_globalState))
		_104_v29 = _nw6
		var _105_v30 _dafny.Int
		_ = _105_v30
		var _out0 _dafny.Int
		_ = _out0
		_out0 = (_104_v29).M1((_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_85_v12), 0))).Int()).(_dafny.Int), _71_globalState)
		_105_v30 = _out0
	}
	var _106_v31 _dafny.Array
	_ = _106_v31
	var _nwElement0_1 bool = _72_v1
	_ = _nwElement0_1
	var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(11))
	_ = _nw7
	(_nw7).ArraySet1(_nwElement0_1, 0)
	(_nw7).ArraySet1(_72_v1, 1)
	(_nw7).ArraySet1(_72_v1, 2)
	(_nw7).ArraySet1(_72_v1, 3)
	(_nw7).ArraySet1(_72_v1, 4)
	(_nw7).ArraySet1(_72_v1, 5)
	(_nw7).ArraySet1(_72_v1, 6)
	(_nw7).ArraySet1(_72_v1, 7)
	(_nw7).ArraySet1(!(_72_v1), 8)
	(_nw7).ArraySet1(_72_v1, 9)
	(_nw7).ArraySet1(_72_v1, 10)
	_106_v31 = _nw7
	var _107_v32 _dafny.Int
	_ = _107_v32
	_107_v32 = _dafny.IntOfInt64(714)
	var _108_v33 _dafny.Set
	_ = _108_v33
	_108_v33 = _dafny.SetOf(_72_v1)
	var _109_v34 *C5
	_ = _109_v34
	var _nw8 *C5 = New_C5_()
	_ = _nw8
	_nw8.Ctor__(_106_v31, ((_dafny.Zero).Minus(_107_v32)).Minus((_108_v33).Cardinality()), (_107_v32).Minus(Companion_Default___.Fm0(_71_globalState)), (_108_v33).IsProperSubsetOf(_108_v33))
	_109_v34 = _nw8
	var _110_v35 _dafny.CodePoint
	_ = _110_v35
	_110_v35 = _dafny.CodePoint('s')
	var _111_v36 _dafny.Sequence
	_ = _111_v36
	_111_v36 = _dafny.UnicodeSeqOfUtf8Bytes("ndf")
	var _112_v37 _dafny.Array
	_ = _112_v37
	var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
	_ = _nw9
	_112_v37 = _nw9
	var _113_v38 *C3
	_ = _113_v38
	var _nw10 *C3 = New_C3_()
	_ = _nw10
	_nw10.Ctor__(_110_v35, _111_v36, _107_v32, _72_v1, (Companion_D10_.Create_DC30_(_dafny.IntOfUint32((Companion_Default___.Fm1(_110_v35, _72_v1, _71_globalState)).Cardinality()), _112_v37, _72_v1, _107_v32, _107_v32)).Dtor_cf44())
	_113_v38 = _nw10
	var _114_v39 _dafny.Map
	_ = _114_v39
	_114_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v38, _107_v32)
	var _115_v40 _dafny.Sequence
	_ = _115_v40
	_115_v40 = _dafny.SeqOf(_72_v1, _72_v1, _72_v1, _72_v1)
	var _116_v41 _dafny.Sequence
	_ = _116_v41
	_116_v41 = _dafny.SeqOf((_114_v39).Update(_113_v38, (_dafny.MultiSetFromSeq(_115_v40)).Cardinality()))
	var _hi0 _dafny.Int = (func() _dafny.Int {
		if _72_v1 {
			return _107_v32
		}
		return _107_v32
	})()
	_ = _hi0
	for _117_i3 := _dafny.IntOfUint32((_116_v41).Cardinality()); _117_i3.Cmp(_hi0) < 0; _117_i3 = _117_i3.Plus(_dafny.One) {
		var _118_v42 _dafny.Array
		_ = _118_v42
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
		_ = _nw11
		_118_v42 = _nw11
		var _119_v43 _dafny.Array
		_ = _119_v43
		var _nw12 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
		_ = _nw12
		_119_v43 = _nw12
		var _120_v44 _dafny.Map
		_ = _120_v44
		_120_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v32, _119_v43)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_118_v42), 0))
		_ = _index4
		(_118_v42).ArraySet1((_120_v44).Merge(_120_v44), (_index4).Int())
		var _121_v45 T1
		_ = _121_v45
		var _nw13 *C4 = New_C4_()
		_ = _nw13
		_nw13.Ctor__(_dafny.IntOfInt64(472), (Companion_Default___.Fm31(_71_globalState)).Cardinality(), _72_v1)
		_121_v45 = _nw13
		var _122_v46 _dafny.Map
		_ = _122_v46
		_122_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_117_i3, _dafny.IntOfUint32((_dafny.SeqOf(_121_v45, _121_v45)).Cardinality())), _120_v44)
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_118_v42), 0))
		_ = _index5
		(_118_v42).ArraySet1((func() _dafny.Map {
			if (_122_v46).Contains((_dafny.Zero).Minus((_107_v32).Plus(_117_i3))) {
				return (_122_v46).Get((_dafny.Zero).Minus((_107_v32).Plus(_117_i3))).(_dafny.Map)
			}
			return _120_v44
		})(), (_index5).Int())
		var _123_v47 _dafny.Map
		_ = _123_v47
		_123_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v36, _113_v38.F13)
		var _124_v48 _dafny.Array
		_ = _124_v48
		var _nwElement0_2 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("wiusmh")
		_ = _nwElement0_2
		var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(29))
		_ = _nw14
		(_nw14).ArraySet1(_nwElement0_2, 0)
		(_nw14).ArraySet1(Companion_Default___.Fm20((_121_v45).F3(), _71_globalState), 1)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(584))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_125_v38 *C3) func(_dafny.Int) _dafny.CodePoint {
			return func(_126_i4 _dafny.Int) _dafny.CodePoint {
				return (_125_v38).F12()
			}
		})(_113_v38))), _dafny.UnicodeSeqOfUtf8Bytes("tr")), 2)
		(_nw14).ArraySet1(_113_v38.F13, 3)
		(_nw14).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("bybpuk"), 4)
		(_nw14).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qkhmicrt"), 5)
		(_nw14).ArraySet1(_111_v36, 6)
		(_nw14).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-349))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_127_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_128_i5 _dafny.Int) _dafny.CodePoint {
				return _127_v35
			}
		})(_110_v35))), 7)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_111_v36, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(293))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_129_v38 *C3) func(_dafny.Int) _dafny.CodePoint {
			return func(_130_i6 _dafny.Int) _dafny.CodePoint {
				return (_129_v38).F12()
			}
		})(_113_v38))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(293))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_131_v38 *C3) func(_dafny.Int) _dafny.CodePoint {
			return func(_132_i6 _dafny.Int) _dafny.CodePoint {
				return (_131_v38).F12()
			}
		})(_113_v38)))).Cardinality()))).Uint32(), (Companion_D8_.Create_DC23_(_121_v45.F4(), _72_v1, true, (_113_v38).F12())).Dtor_cf35())), 8)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm20(_dafny.IntOfInt64(-245), _71_globalState), _111_v36), 9)
		(_nw14).ArraySet1(_113_v38.F13, 10)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_113_v38.F13, _111_v36), 11)
		(_nw14).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("crcifcx"), 12)
		(_nw14).ArraySet1(_113_v38.F13, 13)
		(_nw14).ArraySet1(_113_v38.F13, 14)
		(_nw14).ArraySet1(_113_v38.F13, 15)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_111_v36, _113_v38.F13), 16)
		(_nw14).ArraySet1(Companion_Default___.Fm20((_121_v45).F7(), _71_globalState), 17)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_113_v38.F13, _111_v36), 18)
		(_nw14).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jluu"), 19)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wagffgkjj"), _113_v38.F13), (Companion_Default___.SafeIndex(_107_v32, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wagffgkjj"), _113_v38.F13)).Cardinality()))).Uint32(), (_113_v38).F12()), 20)
		(_nw14).ArraySet1(_111_v36, 21)
		(_nw14).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("agrk"), 22)
		(_nw14).ArraySet1((func() _dafny.Sequence {
			if _72_v1 {
				return _113_v38.F13
			}
			return (func() _dafny.Sequence {
				if (_123_v47).Contains(_dafny.UnicodeSeqOfUtf8Bytes("x")) {
					return (_123_v47).Get(_dafny.UnicodeSeqOfUtf8Bytes("x")).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("k")
			})()
		})(), 23)
		(_nw14).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-793))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_133_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_134_i7 _dafny.Int) _dafny.CodePoint {
				return _133_v35
			}
		})(_110_v35))), 24)
		(_nw14).ArraySet1(_113_v38.F13, 25)
		(_nw14).ArraySet1(_113_v38.F13, 26)
		(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_111_v36, _113_v38.F13), 27)
		(_nw14).ArraySet1((func() _dafny.Sequence {
			if _121_v45.F4() {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(854))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}((func(_135_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_136_i8 _dafny.Int) _dafny.CodePoint {
						return _135_v35
					}
				})(_110_v35)))
			}
			return _111_v36
		})(), 28)
		_124_v48 = _nw14
		_124_v48 = _124_v48
		if !(_121_v45.F4()) || (!(true) || (_72_v1)) {
			var _137_v49 _dafny.Map
			_ = _137_v49
			_137_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_121_v45).F3()).Plus((_121_v45).F3()), _110_v35)
			_137_v49 = (_137_v49).Update((_113_v38).Fm12(!(_121_v45.F4()), _121_v45.F4(), _71_globalState), _110_v35)
			(_113_v38).M5(_71_globalState)
			(_71_globalState).F1 = _117_i3
			var _138_v50 _dafny.Array
			_ = _138_v50
			var _nw15 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
			_ = _nw15
			_138_v50 = _nw15
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_138_v50), 0))
			_ = _index6
			(_138_v50).ArraySet1(_121_v45, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_138_v50), 0))
			_ = _index7
			var _nw16 *C5 = New_C5_()
			_ = _nw16
			_nw16.Ctor__((_109_v34).F9(), (_121_v45).F3(), _dafny.IntOfInt64(-406), _121_v45.F4())
			(_138_v50).ArraySet1(_nw16, (_index7).Int())
			var _139_v51 *C7
			_ = _139_v51
			var _nw17 *C7 = New_C7_()
			_ = _nw17
			_nw17.Ctor__(_121_v45.F4(), (_121_v45).F7(), (_121_v45).F7(), (_108_v33).IsProperSubsetOf(_108_v33))
			_139_v51 = _nw17
		} else {
			(_121_v45).F4_set_(_121_v45.F4())
			var _140_v52 *C7
			_ = _140_v52
			var _nw18 *C7 = New_C7_()
			_ = _nw18
			_nw18.Ctor__(_121_v45.F4(), (_121_v45).F7(), _dafny.IntOfUint32((_111_v36).Cardinality()), _72_v1)
			_140_v52 = _nw18
			var _141_v53 _dafny.Sequence
			_ = _141_v53
			_141_v53 = _dafny.SeqOf(_140_v52, _140_v52, _140_v52)
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_112_v37), 0))
			_ = _index8
			(_112_v37).ArraySet1(_dafny.IntOfUint32((_141_v53).Cardinality()), (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_112_v37), 0))
			_ = _index9
			(_112_v37).ArraySet1((_140_v52).F6(), (_index9).Int())
			var _142_v54 _dafny.Set
			_ = _142_v54
			_142_v54 = _dafny.SetOf(_109_v34, _109_v34)
			_142_v54 = _142_v54
			var _143_v55 _dafny.Array
			_ = _143_v55
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw19
			_143_v55 = _nw19
			var _144_v56 T0
			_ = _144_v56
			var _nw20 *C2 = New_C2_()
			_ = _nw20
			_nw20.Ctor__((_121_v45).F7(), _143_v55, _107_v32, _72_v1)
			_144_v56 = _nw20
			var _145_v57 _dafny.Set
			_ = _145_v57
			_145_v57 = _dafny.SetOf(_144_v56)
			var _146_v58 _dafny.Map
			_ = _146_v58
			_146_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v45.F4(), (_145_v57).Cardinality())
			var _147_v59 _dafny.Map
			_ = _147_v59
			_147_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_146_v58).Merge(_146_v58)).Cardinality(), _121_v45.F4())
			var _rhs5 bool = !((func() bool {
				if (_147_v59).Contains((_dafny.Zero).Minus(_117_i3)) {
					return (_147_v59).Get((_dafny.Zero).Minus(_117_i3)).(bool)
				}
				return _121_v45.F4()
			})())
			_ = _rhs5
			var _rhs6 bool = _121_v45.F4()
			_ = _rhs6
			var _rhs7 _dafny.Int = (_121_v45).F7()
			_ = _rhs7
			var _lhs2 *C7 = _140_v52
			_ = _lhs2
			var _lhs3 *GlobalState = _71_globalState
			_ = _lhs3
			_72_v1 = _rhs5
			_lhs2.F5 = _rhs6
			_lhs3.F1 = _rhs7
			_143_v55 = _143_v55
		}
		var _148_v60 _dafny.Map
		_ = _148_v60
		_148_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v45.F4(), _111_v36)
		(_121_v45).F4_set_(((_121_v45).F3()).Cmp(((_113_v38).Fm12(Companion_Default___.Fm2(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_148_v60).Contains((_113_v38).Fm13(_71_globalState)) {
				return (_148_v60).Get((_113_v38).Fm13(_71_globalState)).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("epeew")
		})()).Cardinality()), _71_globalState), _72_v1, _71_globalState)).Minus((_121_v45).F7())) >= 0)
	}
	var _hi1 _dafny.Int = _107_v32
	_ = _hi1
	for _149_i9 := _107_v32; _149_i9.Cmp(_hi1) < 0; _149_i9 = _149_i9.Plus(_dafny.One) {
		var _150_v61 D8
		_ = _150_v61
		_150_v61 = Companion_D8_.Create_DC23_(_72_v1, !(_72_v1), false, _dafny.CodePoint('e'))
		var _151_v62 D8
		_ = _151_v62
		_151_v62 = Companion_D8_.Create_DC24_(_150_v61)
		var _152_v63 D8
		_ = _152_v63
		_152_v63 = Companion_D8_.Create_DC24_(_150_v61)
		var _153_v64 D13
		_ = _153_v64
		_153_v64 = Companion_D13_.Create_DC40_(_152_v63, (_dafny.Zero).Minus(_107_v32))
		var _154_v65 D12
		_ = _154_v65
		_154_v65 = Companion_D12_.Create_DC36_(_72_v1, !(_72_v1), Companion_Default___.Fm0(_71_globalState))
		var _155_v66 D12
		_ = _155_v66
		_155_v66 = Companion_D12_.Create_DC36_(_72_v1, _72_v1, (_154_v65).Dtor_cf63())
		var _rhs8 bool = (_107_v32).Cmp((_154_v65).Dtor_cf63()) >= 0
		_ = _rhs8
		var _rhs9 bool = _72_v1
		_ = _rhs9
		var _rhs10 D13 = _153_v64
		_ = _rhs10
		var _lhs4 *GlobalState = _71_globalState
		_ = _lhs4
		_72_v1 = _rhs8
		_lhs4.F2 = _rhs9
		_153_v64 = _rhs10
		var _156_v67 _dafny.MultiSet
		_ = _156_v67
		_156_v67 = _dafny.MultiSetOf((_113_v38).F12())
		var _157_v68 *C4
		_ = _157_v68
		var _nw21 *C4 = New_C4_()
		_ = _nw21
		_nw21.Ctor__(_107_v32, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_72_v1))).Cardinality(), (Companion_D0_.Create_DC2_(_156_v67, !(_72_v1))).Dtor_cf6())
		_157_v68 = _nw21
		var _rhs11 *C3 = _113_v38
		_ = _rhs11
		var _rhs12 *C4 = _157_v68
		_ = _rhs12
		_113_v38 = _rhs11
		_157_v68 = _rhs12
		_107_v32 = _107_v32
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw22
		_106_v31 = _nw22
	}
	var _158_v69 *C1
	_ = _158_v69
	var _nw23 *C1 = New_C1_()
	_ = _nw23
	_nw23.Ctor__(_dafny.IntOfUint32((_111_v36).Cardinality()), _dafny.IntOfInt64(840), (_115_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), _dafny.IntOfUint32((_115_v40).Cardinality()))).Uint32()).(bool))
	_158_v69 = _nw23
	var _159_v70 _dafny.Sequence
	_ = _159_v70
	_159_v70 = _dafny.SeqOf(_158_v69)
	var _160_v71 _dafny.Set
	_ = _160_v71
	_160_v71 = _dafny.SetOf(_158_v69, _158_v69, _158_v69, (_159_v70).Select((Companion_Default___.SafeIndex(_107_v32, _dafny.IntOfUint32((_159_v70).Cardinality()))).Uint32()).(*C1))
	_160_v71 = (_160_v71).Intersection(_160_v71)
	var _hi2 _dafny.Int = (_dafny.IntOfInt64(273)).Minus(_107_v32)
	_ = _hi2
	for _161_i10 := _107_v32; _161_i10.Cmp(_hi2) < 0; _161_i10 = _161_i10.Plus(_dafny.One) {
		var _162_v72 _dafny.MultiSet
		_ = _162_v72
		_162_v72 = _dafny.MultiSetOf(_72_v1, _72_v1)
		var _163_v73 _dafny.MultiSet
		_ = _163_v73
		_163_v73 = _dafny.MultiSetOf(_161_i10)
		(_113_v38).F13 = Companion_Default___.Fm15((_72_v1) || (_72_v1), Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_162_v72).Contains(!(_72_v1)) {
				return (_162_v72).Multiplicity(!(_72_v1))
			}
			return _dafny.IntOfInt64(-565)
		})(), _161_i10), (func() _dafny.Int {
			if (_163_v73).Contains(_161_i10) {
				return (_163_v73).Multiplicity(_161_i10)
			}
			return _107_v32
		})(), _71_globalState)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_106_v31), 0))
		_ = _index10
		(_106_v31).ArraySet1(_72_v1, (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_106_v31), 0))
		_ = _index11
		(_106_v31).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v1, _161_i10)).Cardinality()).Cmp(_dafny.IntOfInt64(-570)) < 0, (_index11).Int())
		var _164_v74 _dafny.Set
		_ = _164_v74
		_164_v74 = _dafny.SetOf(_113_v38, _113_v38, _113_v38, _113_v38, _113_v38)
		_164_v74 = _164_v74
		var _165_v75 _dafny.Array
		_ = _165_v75
		var _len0_3 _dafny.Int = _dafny.One
		_ = _len0_3
		var _nw24 _dafny.Array
		_ = _nw24
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw24 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_166_v31 _dafny.Array) func(_dafny.Int) bool {
				return func(_167_i11 _dafny.Int) bool {
					return !((_166_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_166_v31), 0))).Int()).(bool)) || ((_166_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_166_v31), 0))).Int()).(bool))
				}
			})(_106_v31)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw24 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw24).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw24).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_165_v75 = _nw24
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_69_v0), 0))
		_ = _index12
		(_69_v0).ArraySet1CodePoint((_113_v38).F12(), (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_69_v0), 0))
		_ = _index13
		var _rhs13 _dafny.Array = _165_v75
		_ = _rhs13
		var _rhs14 _dafny.Int = Companion_Default___.SafeModuloInt((_107_v32).Minus(_107_v32), _107_v32)
		_ = _rhs14
		var _rhs15 _dafny.Int = _161_i10
		_ = _rhs15
		var _rhs16 _dafny.CodePoint = (_113_v38).F12()
		_ = _rhs16
		var _rhs17 bool = _72_v1
		_ = _rhs17
		var _lhs5 *GlobalState = _71_globalState
		_ = _lhs5
		var _lhs6 *GlobalState = _71_globalState
		_ = _lhs6
		var _lhs7 _dafny.Array = _69_v0
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_69_v0), 0))
		_ = _lhs8
		_165_v75 = _rhs13
		_lhs5.F1 = _rhs14
		_lhs6.F1 = _rhs15
		(_lhs7).ArraySet1CodePoint(_rhs16, (_lhs8).Int())
		_72_v1 = _rhs17
	}
	_72_v1 = _72_v1
	var _168_v76 _dafny.Set
	_ = _168_v76
	_168_v76 = _dafny.SetOf(_107_v32)
	var _169_i12 _dafny.Int
	_ = _169_i12
	_169_i12 = _dafny.Zero
	{
		for !((_107_v32).Cmp(Companion_Default___.SafeModuloInt((_168_v76).Cardinality(), (_dafny.Zero).Minus(_107_v32))) != 0) {
			{
				if (_169_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_169_i12 = (_169_i12).Plus(_dafny.One)
				_107_v32 = _107_v32
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
				_ = _index14
				(_112_v37).ArraySet1(Companion_Default___.SafeDivisionInt(_107_v32, Companion_Default___.Fm0(_71_globalState)), (_index14).Int())
				var _170_v77 _dafny.Array
				_ = _170_v77
				var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw25
				_170_v77 = _nw25
				var _171_v78 _dafny.Array
				_ = _171_v78
				var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
				_ = _nw26
				_171_v78 = _nw26
				var _172_v79 _dafny.Set
				_ = _172_v79
				_172_v79 = _dafny.SetOf((_113_v38).F12(), _dafny.CodePoint('k'))
				var _173_v80 T1
				_ = _173_v80
				var _nw27 *C5 = New_C5_()
				_ = _nw27
				_nw27.Ctor__((_109_v34).F9(), _dafny.IntOfUint32((_113_v38.F13).Cardinality()), (_172_v79).Cardinality(), _72_v1)
				_173_v80 = _nw27
				var _174_v81 _dafny.Set
				_ = _174_v81
				_174_v81 = _dafny.SetOf(_173_v80)
				var _175_v82 _dafny.Map
				_ = _175_v82
				_175_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v1, _174_v81)
				var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_171_v78), 0))
				_ = _index15
				(_171_v78).ArraySet1(_175_v82, (_index15).Int())
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
				_ = _index16
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_171_v78), 0))
				_ = _index17
				var _rhs18 _dafny.Int = (_168_v76).Cardinality()
				_ = _rhs18
				var _rhs19 _dafny.Array = _170_v77
				_ = _rhs19
				var _rhs20 _dafny.Map = (((_175_v82).Update(_173_v80.F4(), _174_v81)).Merge(_175_v82)).Update(_72_v1, _174_v81)
				_ = _rhs20
				var _lhs9 _dafny.Array = _112_v37
				_ = _lhs9
				var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
				_ = _lhs10
				var _lhs11 _dafny.Array = _171_v78
				_ = _lhs11
				var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_171_v78), 0))
				_ = _lhs12
				(_lhs9).ArraySet1(_rhs18, (_lhs10).Int())
				_170_v77 = _rhs19
				(_lhs11).ArraySet1(_rhs20, (_lhs12).Int())
				if _173_v80.F4() {
					var _176_v83 _dafny.Array
					_ = _176_v83
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_4
					var _nw28 _dafny.Array
					_ = _nw28
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw28 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) _dafny.Int = (func(_177_v80 T1) func(_dafny.Int) _dafny.Int {
							return func(_178_i13 _dafny.Int) _dafny.Int {
								return (_178_i13).Times((_177_v80).F7())
							}
						})(_173_v80)
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw28 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw28).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw28).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_176_v83 = _nw28
					var _179_v84 _dafny.Array
					_ = _179_v84
					var _nw29 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
					_ = _nw29
					_179_v84 = _nw29
					var _180_v85 _dafny.Map
					_ = _180_v85
					_180_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v84, _113_v38)
					var _rhs21 _dafny.Array = _176_v83
					_ = _rhs21
					var _rhs22 *C3 = (func() *C3 {
						if (_180_v85).Contains(_179_v84) {
							return (_180_v85).Get(_179_v84).(*C3)
						}
						return _113_v38
					})()
					_ = _rhs22
					_176_v83 = _rhs21
					_113_v38 = _rhs22
					var _181_v86 *C2
					_ = _181_v86
					var _nw30 *C2 = New_C2_()
					_ = _nw30
					_nw30.Ctor__(_107_v32, _176_v83, (_173_v80).F7(), _173_v80.F4())
					_181_v86 = _nw30
					var _182_v87 _dafny.Map
					_ = _182_v87
					_182_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_173_v80.F4(), _173_v80.F4())
					var _183_v88 D12
					_ = _183_v88
					_183_v88 = Companion_D12_.Create_DC37_(_182_v87, _72_v1, (_112_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))).Int()).(_dafny.Int), (_173_v80).F7())
					var _184_v89 _dafny.Sequence
					_ = _184_v89
					_184_v89 = _dafny.SeqOf((_173_v80).F7(), Companion_Default___.SafeModuloInt((_173_v80).F7(), (_173_v80).F7()), (_183_v88).Dtor_cf66())
					var _185_v90 _dafny.Map
					_ = _185_v90
					_185_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_181_v86).F14(), _dafny.SeqOf(_173_v80.F4(), _173_v80.F4()))
					var _186_v91 _dafny.MultiSet
					_ = _186_v91
					_186_v91 = _dafny.MultiSetOf(_113_v38.F13, _111_v36)
					var _187_v92 _dafny.MultiSet
					_ = _187_v92
					_187_v92 = _dafny.MultiSetOf((_173_v80).F3())
					var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
					_ = _index18
					var _rhs23 *C2 = _181_v86
					_ = _rhs23
					var _rhs24 _dafny.Sequence = _dafny.SeqOf(((_173_v80).F7()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_185_v90).Contains((_173_v80).F7()) {
							return (_185_v90).Get((_173_v80).F7()).(_dafny.Sequence)
						}
						return _115_v40
					})()).Cardinality()))) == 0, ((_186_v91).Cardinality()).Cmp((_173_v80).F3()) != 0, _173_v80.F4())
					_ = _rhs24
					var _rhs25 _dafny.Int = Companion_Default___.Fm0(_71_globalState)
					_ = _rhs25
					var _rhs26 _dafny.Sequence = _184_v89
					_ = _rhs26
					var _rhs27 _dafny.Sequence = Companion_Default___.Fm4(_173_v80.F4(), _72_v1, Companion_Default___.Fm2((_dafny.Zero).Minus((_173_v80).F3()), _71_globalState), (_107_v32).Cmp((_187_v92).Cardinality()) < 0, _71_globalState)
					_ = _rhs27
					var _lhs13 _dafny.Array = _112_v37
					_ = _lhs13
					var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
					_ = _lhs14
					_181_v86 = _rhs23
					_115_v40 = _rhs24
					(_lhs13).ArraySet1(_rhs25, (_lhs14).Int())
					_184_v89 = _rhs26
					_184_v89 = _rhs27
					var _188_v93 _dafny.Map
					_ = _188_v93
					_188_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_173_v80).F7(), _173_v80.F4())
					(_71_globalState).F2 = (func() bool {
						if (_188_v93).Contains((_181_v86).F14()) {
							return (_188_v93).Get((_181_v86).F14()).(bool)
						}
						return _173_v80.F4()
					})()
					(_173_v80).F4_set_(Companion_Default___.Fm2((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_173_v80).F7(), (func() _dafny.Int {
						if (_187_v92).Contains((_173_v80).F7()) {
							return (_187_v92).Multiplicity((_173_v80).F7())
						}
						return (_173_v80).F3()
					})())), _71_globalState))
					var _189_v94 *C7
					_ = _189_v94
					var _nw31 *C7 = New_C7_()
					_ = _nw31
					_nw31.Ctor__(!(false), Companion_Default___.Fm0(_71_globalState), _107_v32, _173_v80.F4())
					_189_v94 = _nw31
					var _190_v95 _dafny.Array
					_ = _190_v95
					var _nwElement0_3 *C7 = _189_v94
					_ = _nwElement0_3
					var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(5))
					_ = _nw32
					(_nw32).ArraySet1(_nwElement0_3, 0)
					(_nw32).ArraySet1(_189_v94, 1)
					(_nw32).ArraySet1(_189_v94, 2)
					(_nw32).ArraySet1(_189_v94, 3)
					(_nw32).ArraySet1(_189_v94, 4)
					_190_v95 = _nw32
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_190_v95), 0))
					_ = _index19
					(_190_v95).ArraySet1(_189_v94, (_index19).Int())
					var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
					_ = _index20
					var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_190_v95), 0))
					_ = _index21
					var _rhs28 _dafny.Int = (_181_v86).F14()
					_ = _rhs28
					var _rhs29 _dafny.Int = Companion_Default___.Fm0(_71_globalState)
					_ = _rhs29
					var _rhs30 _dafny.Int = (_173_v80).F7()
					_ = _rhs30
					var _rhs31 _dafny.Array = (_181_v86).F15()
					_ = _rhs31
					var _rhs32 *C7 = _189_v94
					_ = _rhs32
					var _lhs15 *GlobalState = _71_globalState
					_ = _lhs15
					var _lhs16 _dafny.Array = _112_v37
					_ = _lhs16
					var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
					_ = _lhs17
					var _lhs18 *GlobalState = _71_globalState
					_ = _lhs18
					var _lhs19 _dafny.Array = _190_v95
					_ = _lhs19
					var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_190_v95), 0))
					_ = _lhs20
					_lhs15.F1 = _rhs28
					(_lhs16).ArraySet1(_rhs29, (_lhs17).Int())
					_lhs18.F1 = _rhs30
					_176_v83 = _rhs31
					(_lhs19).ArraySet1(_rhs32, (_lhs20).Int())
				} else {
					var _191_v96 _dafny.Map
					_ = _191_v96
					_191_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_173_v80).F7(), _69_v0)
					var _192_v97 _dafny.MultiSet
					_ = _192_v97
					_192_v97 = _dafny.MultiSetOf((func() _dafny.Array {
						if (_191_v96).Contains((_173_v80).F7()) {
							return (_191_v96).Get((_173_v80).F7()).(_dafny.Array)
						}
						return _69_v0
					})())
					_192_v97 = ((_192_v97).Difference(_dafny.MultiSetOf(_69_v0))).Union((_192_v97).Difference(_192_v97))
					var _193_v98 _dafny.Map
					_ = _193_v98
					_193_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v34, _111_v36)
					var _194_v99 _dafny.MultiSet
					_ = _194_v99
					_194_v99 = _dafny.MultiSetOf(true)
					var _195_v100 _dafny.Map
					_ = _195_v100
					_195_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_193_v98).Update(_109_v34, _113_v38.F13)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v34, _113_v38.F13)), ((_108_v33).Cardinality()).Times((_194_v99).Cardinality()))
					_195_v100 = (_195_v100).Update(_193_v98, (_173_v80).F3())
					var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_112_v37), 0))
					_ = _index22
					(_112_v37).ArraySet1((_113_v38).Fm12(Companion_Default___.Fm2(_107_v32, _71_globalState), _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jd"), _111_v36), _71_globalState), (_index22).Int())
					(_173_v80).F4_set_(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_111_v36, _113_v38.F13), (Companion_Default___.SafeIndex((_173_v80).F3(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_111_v36, _113_v38.F13)).Cardinality()))).Uint32(), (_113_v38).F12()), _111_v36))
					var _196_v101 D12
					_ = _196_v101
					_196_v101 = Companion_D12_.Create_DC35_((_109_v34).F9(), true, _dafny.SeqOf(_72_v1), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-514))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}((func(_197_v38 *C3) func(_dafny.Int) _dafny.CodePoint {
						return func(_198_i14 _dafny.Int) _dafny.CodePoint {
							return (_197_v38).F12()
						}
					})(_113_v38))), (_173_v80).F7())
					var _pat_let_tv0 = _113_v38
					_ = _pat_let_tv0
					var _pat_let_tv1 = _173_v80
					_ = _pat_let_tv1
					var _pat_let_tv2 = _115_v40
					_ = _pat_let_tv2
					var _pat_let_tv3 = _106_v31
					_ = _pat_let_tv3
					_111_v36 = (func(_pat_let0_0 D12) D12 {
						return func(_199_dt__update__tmp_h0 D12) D12 {
							return func(_pat_let1_0 _dafny.Sequence) D12 {
								return func(_200_dt__update_hcf59_h0 _dafny.Sequence) D12 {
									return func(_pat_let2_0 _dafny.Int) D12 {
										return func(_201_dt__update_hcf60_h0 _dafny.Int) D12 {
											return func(_pat_let3_0 _dafny.Sequence) D12 {
												return func(_202_dt__update_hcf58_h0 _dafny.Sequence) D12 {
													return func(_pat_let4_0 _dafny.Array) D12 {
														return func(_203_dt__update_hcf56_h0 _dafny.Array) D12 {
															return Companion_D12_.Create_DC35_(_203_dt__update_hcf56_h0, (_199_dt__update__tmp_h0).Dtor_cf57(), _202_dt__update_hcf58_h0, _200_dt__update_hcf59_h0, _201_dt__update_hcf60_h0)
														}(_pat_let4_0)
													}(_pat_let_tv3)
												}(_pat_let3_0)
											}(_pat_let_tv2)
										}(_pat_let2_0)
									}((_pat_let_tv1).F7())
								}(_pat_let1_0)
							}(_pat_let_tv0.F13)
						}(_pat_let0_0)
					}(_196_v101)).Dtor_cf59()
				}
				(_113_v38).M5(_71_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _204_v102 _dafny.Map
	_ = _204_v102
	_204_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v1, _110_v35)
	var _205_v103 _dafny.Sequence
	_ = _205_v103
	_205_v103 = _dafny.SeqOf(_107_v32)
	var _206_i15 _dafny.Int
	_ = _206_i15
	_206_i15 = _dafny.Zero
	{
		for !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-792))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_217_v32 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_218_i16 _dafny.Int) _dafny.Int {
				return _217_v32
			}
		})(_107_v32))), _205_v103)) || (_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_113_v38.F13, (Companion_Default___.SafeIndex(_107_v32, _dafny.IntOfUint32((_113_v38.F13).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
			if (_204_v102).Contains(_72_v1) {
				return (_204_v102).Get(_72_v1).(_dafny.CodePoint)
			}
			return (_113_v38).F12()
		})()), _111_v36)) {
			{
				if (_206_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_206_i15 = (_206_i15).Plus(_dafny.One)
				var _207_v104 _dafny.MultiSet
				_ = _207_v104
				_207_v104 = _dafny.MultiSetOf(_72_v1)
				(_71_globalState).F1 = ((_dafny.IntOfUint32((_113_v38.F13).Cardinality())).Minus((_dafny.Zero).Minus(_107_v32))).Minus((_207_v104).Cardinality())
				var _208_v105 _dafny.Map
				_ = _208_v105
				_208_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_72_v1) && (_72_v1), _72_v1)
				_208_v105 = (_208_v105).Update((_107_v32).Cmp(_107_v32) <= 0, _72_v1)
				var _209_v106 *C4
				_ = _209_v106
				var _nw33 *C4 = New_C4_()
				_ = _nw33
				_nw33.Ctor__((_207_v104).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_205_v103).Cardinality())), Companion_Default___.Fm2(_107_v32, _71_globalState))
				_209_v106 = _nw33
				var _210_v107 _dafny.Map
				_ = _210_v107
				_210_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v1, _209_v106)
				var _211_v108 _dafny.Sequence
				_ = _211_v108
				_211_v108 = _dafny.SeqOf(_209_v106, _209_v106)
				var _212_v109 _dafny.Array
				_ = _212_v109
				var _nwElement0_4 *C4 = _209_v106
				_ = _nwElement0_4
				var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(14))
				_ = _nw34
				(_nw34).ArraySet1(_nwElement0_4, 0)
				(_nw34).ArraySet1(_209_v106, 1)
				(_nw34).ArraySet1(_209_v106, 2)
				(_nw34).ArraySet1(_209_v106, 3)
				(_nw34).ArraySet1(_209_v106, 4)
				(_nw34).ArraySet1(_209_v106, 5)
				(_nw34).ArraySet1(_209_v106, 6)
				(_nw34).ArraySet1((func() *C4 {
					if (_210_v107).Contains(true) {
						return (_210_v107).Get(true).(*C4)
					}
					return _209_v106
				})(), 7)
				(_nw34).ArraySet1(_209_v106, 8)
				(_nw34).ArraySet1(_209_v106, 9)
				(_nw34).ArraySet1(_209_v106, 10)
				(_nw34).ArraySet1(_209_v106, 11)
				(_nw34).ArraySet1((_211_v108).Select((Companion_Default___.SafeIndex(_107_v32, _dafny.IntOfUint32((_211_v108).Cardinality()))).Uint32()).(*C4), 12)
				(_nw34).ArraySet1(_209_v106, 13)
				_212_v109 = _nw34
				var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_212_v109), 0))
				_ = _index23
				(_212_v109).ArraySet1(_209_v106, (_index23).Int())
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_212_v109), 0))
				_ = _index24
				(_212_v109).ArraySet1(_209_v106, (_index24).Int())
				var _213_v110 _dafny.MultiSet
				_ = _213_v110
				var _214_v111 _dafny.CodePoint
				_ = _214_v111
				var _215_v112 _dafny.Map
				_ = _215_v112
				var _216_v113 bool
				_ = _216_v113
				var _out1 _dafny.MultiSet
				_ = _out1
				var _out2 _dafny.CodePoint
				_ = _out2
				var _out3 _dafny.Map
				_ = _out3
				var _out4 bool
				_ = _out4
				_out1, _out2, _out3, _out4 = (_109_v34).M2((_72_v1) == (true), _112_v37, (_208_v105).Merge(_208_v105), _72_v1, _71_globalState)
				_213_v110 = _out1
				_214_v111 = _out2
				_215_v112 = _out3
				_216_v113 = _out4
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_106_v31), 0))); ; {
		_guard_loop_0, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _219_i17 _dafny.Int
		_219_i17 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_219_i17).Sign() != -1) && ((_219_i17).Cmp(_dafny.ArrayLen((_106_v31), 0)) < 0)) {
			(_106_v31).ArraySet1(!(!(!(_72_v1))), (_219_i17).Int())
		}
	}
	var _220_v114 _dafny.Array
	_ = _220_v114
	var _nwElement0_5 *C3 = _113_v38
	_ = _nwElement0_5
	var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(4))
	_ = _nw35
	(_nw35).ArraySet1(_nwElement0_5, 0)
	(_nw35).ArraySet1((func() *C3 {
		if _72_v1 {
			return _113_v38
		}
		return _113_v38
	})(), 1)
	(_nw35).ArraySet1(_113_v38, 2)
	(_nw35).ArraySet1(_113_v38, 3)
	_220_v114 = _nw35
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_220_v114), 0))
	_ = _index25
	(_220_v114).ArraySet1(_113_v38, (_index25).Int())
	var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_220_v114), 0))
	_ = _index26
	(_220_v114).ArraySet1(_113_v38, (_index26).Int())
	var _221_v115 _dafny.Array
	_ = _221_v115
	var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
	_ = _nw36
	_221_v115 = _nw36
	var _222_v116 T0
	_ = _222_v116
	var _nw37 *C3 = New_C3_()
	_ = _nw37
	_nw37.Ctor__((_113_v38).F12(), _111_v36, _dafny.IntOfUint32((_113_v38.F13).Cardinality()), true, _107_v32)
	_222_v116 = _nw37
	var _223_v117 *C2
	_ = _223_v117
	var _nw38 *C2 = New_C2_()
	_ = _nw38
	_nw38.Ctor__((_222_v116).F3(), _112_v37, (_222_v116).F3(), _72_v1)
	_223_v117 = _nw38
	var _224_v118 _dafny.Map
	_ = _224_v118
	_224_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v116, _223_v117)
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_221_v115), 0))
	_ = _index27
	(_221_v115).ArraySet1((_224_v118).Merge(_224_v118), (_index27).Int())
	var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_221_v115), 0))
	_ = _index28
	(_221_v115).ArraySet1(_224_v118, (_index28).Int())
	var _225_v119 _dafny.Sequence
	_ = _225_v119
	_225_v119 = _dafny.SeqOf(_112_v37, _112_v37)
	var _226_v120 _dafny.Sequence
	_ = _226_v120
	_226_v120 = _dafny.SeqOf(_225_v119, (Companion_D15_.Create_DC43_(_225_v119)).Dtor_cf74())
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
	_ = _index29
	(_106_v31).ArraySet1(!_dafny.Companion_Sequence_.Contains((_226_v120).Select((Companion_Default___.SafeIndex((_223_v117).F14(), _dafny.IntOfUint32((_226_v120).Cardinality()))).Uint32()).(_dafny.Sequence), (_223_v117).F15()), (_index29).Int())
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
	_ = _index30
	(_106_v31).ArraySet1(_72_v1, (_index30).Int())
	var _227_i18 _dafny.Int
	_ = _227_i18
	_227_i18 = _dafny.Zero
	{
		for _72_v1 {
			{
				if (_227_i18).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_227_i18 = (_227_i18).Plus(_dafny.One)
				var _228_v121 _dafny.Map
				_ = _228_v121
				_228_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v116.F4(), _222_v116.F4())
				_228_v121 = (_228_v121).Update((_106_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))).Int()).(bool), _222_v116.F4())
				var _229_v122 _dafny.MultiSet
				_ = _229_v122
				_229_v122 = _dafny.MultiSetOf(_107_v32)
				(_71_globalState).F2 = Companion_Default___.Fm2(((_229_v122).Intersection(_229_v122)).Cardinality(), _71_globalState)
				var _230_v123 *C0
				_ = _230_v123
				var _nw39 *C0 = New_C0_()
				_ = _nw39
				_nw39.Ctor__(_222_v116.F4(), (_109_v34).F9())
				_230_v123 = _nw39
				(_222_v116).F4_set_(!(_222_v116.F4()))
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_112_v37), 0))
	_ = _index31
	(_112_v37).ArraySet1((_222_v116).F3(), (_index31).Int())
	var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_112_v37), 0))
	_ = _index32
	(_112_v37).ArraySet1(_107_v32, (_index32).Int())
	var _231_v124 _dafny.Map
	_ = _231_v124
	_231_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v1, false)
	var _232_v125 D13
	_ = _232_v125
	_232_v125 = Companion_D13_.Create_DC38_(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(144))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}((func(_233_v116 T0, _234_v31 _dafny.Array) func(_dafny.Int) _dafny.Map {
		return func(_235_i19 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_233_v116.F4(), (_234_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_234_v31), 0))).Int()).(bool))
		}
	})(_222_v116, _106_v31))), (Companion_Default___.SafeIndex((_223_v117).F14(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(144))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}((func(_236_v116 T0, _237_v31 _dafny.Array) func(_dafny.Int) _dafny.Map {
		return func(_238_i19 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v116.F4(), (_237_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_237_v31), 0))).Int()).(bool))
		}
	})(_222_v116, _106_v31)))).Cardinality()))).Uint32(), _231_v124))
	var _source3 D13 = _232_v125
	_ = _source3
	if _source3.Is_DC39() {
		var _239___mcc_h0 _dafny.Map = _source3.Get_().(D13_DC39).Cf69
		_ = _239___mcc_h0
		var _240_cf69 _dafny.Map = _239___mcc_h0
		_ = _240_cf69
		(_71_globalState).F2 = _222_v116.F4()
		_111_v36 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm15((_106_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))).Int()).(bool), (_223_v117).F14(), (_222_v116).F3(), _71_globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-238))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_241_v38 *C3) func(_dafny.Int) _dafny.CodePoint {
			return func(_242_i20 _dafny.Int) _dafny.CodePoint {
				return (_241_v38).F12()
			}
		})(_113_v38))))
		var _243_v126 _dafny.MultiSet
		_ = _243_v126
		_243_v126 = _dafny.MultiSetOf(_222_v116.F4())
		var _244_v127 _dafny.Map
		_ = _244_v127
		_244_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_223_v117).F14(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sqq")).Cardinality()))
		(_71_globalState).F1 = (((func() _dafny.MultiSet {
			if (_113_v38).Fm13(_71_globalState) {
				return _243_v126
			}
			return _243_v126
		})()).Update((Companion_Default___.Fm2(_107_v32, _71_globalState)) && (_222_v116.F4()), Companion_Default___.Abs((func() _dafny.Int {
			if (_244_v127).Contains((_dafny.SetOf(_109_v34)).Cardinality()) {
				return (_244_v127).Get((_dafny.SetOf(_109_v34)).Cardinality()).(_dafny.Int)
			}
			return Companion_Default___.Fm0(_71_globalState)
		})()))).Cardinality()
		var _245_v128 T1
		_ = _245_v128
		var _nw40 *C3 = New_C3_()
		_ = _nw40
		_nw40.Ctor__((_113_v38).F12(), _113_v38.F13, (_168_v76).Cardinality(), _222_v116.F4(), (_223_v117).F14())
		_245_v128 = _nw40
		var _246_v129 _dafny.MultiSet
		_ = _246_v129
		_246_v129 = _dafny.MultiSetOf(_245_v128)
		var _247_v130 D15
		_ = _247_v130
		_247_v130 = Companion_D15_.Create_DC44_(_245_v128, _110_v35, _107_v32)
		var _248_v131 _dafny.Sequence
		_ = _248_v131
		_248_v131 = _dafny.SeqOf(_245_v128)
		_246_v129 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_247_v130).Dtor_cf75()), _248_v131))).Difference(_dafny.MultiSetFromSeq(_248_v131))
	} else if _source3.Is_DC40() {
		var _249___mcc_h1 D8 = _source3.Get_().(D13_DC40).Cf70
		_ = _249___mcc_h1
		var _250___mcc_h2 _dafny.Int = _source3.Get_().(D13_DC40).Cf71
		_ = _250___mcc_h2
		var _251_cf71 _dafny.Int = _250___mcc_h2
		_ = _251_cf71
		var _252_cf70 D8 = _249___mcc_h1
		_ = _252_cf70
		var _253_v132 _dafny.Sequence
		_ = _253_v132
		_253_v132 = _dafny.SeqOf(_115_v40, _dafny.SeqOf(_222_v116.F4(), _72_v1, (_106_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))).Int()).(bool), _222_v116.F4()))
		var _254_v133 *C6
		_ = _254_v133
		var _nw41 *C6 = New_C6_()
		_ = _nw41
		_nw41.Ctor__((_223_v117).F14(), _251_cf71, _222_v116.F4(), _dafny.IntOfUint32(((_253_v132).Select((Companion_Default___.SafeIndex(_107_v32, _dafny.IntOfUint32((_253_v132).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
		_254_v133 = _nw41
		var _255_v134 _dafny.Sequence
		_ = _255_v134
		_255_v134 = _dafny.SeqOf(_254_v133, _254_v133, _254_v133)
		var _256_v135 _dafny.MultiSet
		_ = _256_v135
		_256_v135 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_205_v103)).Cardinality(), (_223_v117).F14(), _dafny.IntOfInt64(844))
		var _257_v136 D3
		_ = _257_v136
		_257_v136 = Companion_D3_.Create_DC10_((_256_v135).Cardinality(), (_109_v34).F9(), (_dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Int {
			if (_256_v135).Contains((_223_v117).F14()) {
				return (_256_v135).Multiplicity((_223_v117).F14())
			}
			return (_112_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_112_v37), 0))).Int()).(_dafny.Int)
		})()), (_223_v117).F14(), (_222_v116).F3())).Cardinality())
		var _258_v137 _dafny.Array
		_ = _258_v137
		var _nwElement0_6 *C6 = _254_v133
		_ = _nwElement0_6
		var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(24))
		_ = _nw42
		(_nw42).ArraySet1(_nwElement0_6, 0)
		(_nw42).ArraySet1(_254_v133, 1)
		(_nw42).ArraySet1(_254_v133, 2)
		(_nw42).ArraySet1(_254_v133, 3)
		(_nw42).ArraySet1(_254_v133, 4)
		(_nw42).ArraySet1(_254_v133, 5)
		(_nw42).ArraySet1(_254_v133, 6)
		(_nw42).ArraySet1(_254_v133, 7)
		(_nw42).ArraySet1(_254_v133, 8)
		(_nw42).ArraySet1((func() *C6 {
			if _72_v1 {
				return _254_v133
			}
			return _254_v133
		})(), 9)
		(_nw42).ArraySet1((_255_v134).Select((Companion_Default___.SafeIndex((_257_v136).Dtor_cf15(), _dafny.IntOfUint32((_255_v134).Cardinality()))).Uint32()).(*C6), 10)
		(_nw42).ArraySet1(_254_v133, 11)
		(_nw42).ArraySet1(_254_v133, 12)
		(_nw42).ArraySet1(_254_v133, 13)
		(_nw42).ArraySet1(_254_v133, 14)
		(_nw42).ArraySet1(_254_v133, 15)
		(_nw42).ArraySet1(_254_v133, 16)
		(_nw42).ArraySet1(_254_v133, 17)
		(_nw42).ArraySet1(_254_v133, 18)
		(_nw42).ArraySet1(_254_v133, 19)
		(_nw42).ArraySet1(_254_v133, 20)
		(_nw42).ArraySet1(_254_v133, 21)
		(_nw42).ArraySet1(_254_v133, 22)
		(_nw42).ArraySet1(_254_v133, 23)
		_258_v137 = _nw42
		var _rhs33 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qxab"), _113_v38.F13)).Cardinality())
		_ = _rhs33
		var _rhs34 _dafny.Array = _258_v137
		_ = _rhs34
		_107_v32 = _rhs33
		_258_v137 = _rhs34
		var _259_v138 _dafny.Map
		_ = _259_v138
		_259_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_222_v116).F3())
		var _260_v139 D12
		_ = _260_v139
		_260_v139 = Companion_D12_.Create_DC37_(_231_v124, _72_v1, _dafny.IntOfUint32((_111_v36).Cardinality()), (_259_v138).Cardinality())
		var _261_v140 D1
		_ = _261_v140
		_261_v140 = Companion_D1_.Create_DC6_(_222_v116, _dafny.UnicodeSeqOfUtf8Bytes("pvomejghm"), (_106_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))).Int()).(bool), _222_v116.F4())
		var _source4 D12 = (func() D12 {
			if _222_v116.F4() {
				return _260_v139
			}
			return Companion_D12_.Create_DC37_(_231_v124, (_261_v140).Dtor_cf12(), (_222_v116).F3(), (_223_v117).F14())
		})()
		_ = _source4
		if _source4.Is_DC35() {
			var _262___mcc_h5 _dafny.Array = _source4.Get_().(D12_DC35).Cf56
			_ = _262___mcc_h5
			var _263___mcc_h6 bool = _source4.Get_().(D12_DC35).Cf57
			_ = _263___mcc_h6
			var _264___mcc_h7 _dafny.Sequence = _source4.Get_().(D12_DC35).Cf58
			_ = _264___mcc_h7
			var _265___mcc_h8 _dafny.Sequence = _source4.Get_().(D12_DC35).Cf59
			_ = _265___mcc_h8
			var _266___mcc_h9 _dafny.Int = _source4.Get_().(D12_DC35).Cf60
			_ = _266___mcc_h9
			var _267_cf60 _dafny.Int = _266___mcc_h9
			_ = _267_cf60
			var _268_cf59 _dafny.Sequence = _265___mcc_h8
			_ = _268_cf59
			var _269_cf58 _dafny.Sequence = _264___mcc_h7
			_ = _269_cf58
			var _270_cf57 bool = _263___mcc_h6
			_ = _270_cf57
			var _271_cf56 _dafny.Array = _262___mcc_h5
			_ = _271_cf56
			var _272_v141 *C4
			_ = _272_v141
			var _nw43 *C4 = New_C4_()
			_ = _nw43
			_nw43.Ctor__((_108_v33).Cardinality(), (_254_v133).F8(), _222_v116.F4())
			_272_v141 = _nw43
			var _273_v142 _dafny.Map
			_ = _273_v142
			_273_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v32, _272_v141)
			var _274_v143 _dafny.Map
			_ = _274_v143
			_274_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() *C4 {
				if _222_v116.F4() {
					return (func() *C4 {
						if (_273_v142).Contains(_251_cf71) {
							return (_273_v142).Get(_251_cf71).(*C4)
						}
						return _272_v141
					})()
				}
				return _272_v141
			})())
			_274_v143 = _274_v143
			var _275_v144 _dafny.Array
			_ = _275_v144
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_5
			var _nw44 _dafny.Array
			_ = _nw44
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw44 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Sequence = (func(_276_v40 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_277_i21 _dafny.Int) _dafny.Sequence {
						return _276_v40
					}
				})(_115_v40)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw44 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw44).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw44).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_275_v144 = _nw44
			var _278_v145 _dafny.Map
			_ = _278_v145
			_278_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v37, (_254_v133).F8())
			(_223_v117).M6(_275_v144, _251_cf71, ((_278_v145).Cardinality()).Plus((_223_v117).F14()), _71_globalState)
			_231_v124 = (_231_v124).Update(!(!(true)), _222_v116.F4())
			var _279_v146 *C7
			_ = _279_v146
			var _nw45 *C7 = New_C7_()
			_ = _nw45
			_nw45.Ctor__(_270_cf57, Companion_Default___.Fm0(_71_globalState), ((_168_v76).Cardinality()).Times((_223_v117).F14()), true)
			_279_v146 = _nw45
		} else if _source4.Is_DC36() {
			var _280___mcc_h10 bool = _source4.Get_().(D12_DC36).Cf61
			_ = _280___mcc_h10
			var _281___mcc_h11 bool = _source4.Get_().(D12_DC36).Cf62
			_ = _281___mcc_h11
			var _282___mcc_h12 _dafny.Int = _source4.Get_().(D12_DC36).Cf63
			_ = _282___mcc_h12
			var _283_cf63 _dafny.Int = _282___mcc_h12
			_ = _283_cf63
			var _284_cf62 bool = _281___mcc_h11
			_ = _284_cf62
			var _285_cf61 bool = _280___mcc_h10
			_ = _285_cf61
			var _286_v147 _dafny.Array
			_ = _286_v147
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw46
			_286_v147 = _nw46
			var _287_v148 *C0
			_ = _287_v148
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__(false, (_109_v34).F9())
			_287_v148 = _nw47
			var _288_v149 _dafny.Array
			_ = _288_v149
			var _nwElement0_7 *C0 = _287_v148
			_ = _nwElement0_7
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(3))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_7, 0)
			(_nw48).ArraySet1(_287_v148, 1)
			(_nw48).ArraySet1(_287_v148, 2)
			_288_v149 = _nw48
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_286_v147), 0))
			_ = _index33
			(_286_v147).ArraySet1(_288_v149, (_index33).Int())
			var _289_v150 D12
			_ = _289_v150
			_289_v150 = Companion_D12_.Create_DC35_((_109_v34).F9(), _72_v1, _115_v40, _113_v38.F13, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-843))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_290_v38 *C3) func(_dafny.Int) _dafny.CodePoint {
				return func(_291_i22 _dafny.Int) _dafny.CodePoint {
					return (_290_v38).F12()
				}
			})(_113_v38)))).Cardinality()))
			var _292_v151 D4
			_ = _292_v151
			_292_v151 = Companion_D4_.Create_DC11_(_dafny.SetOf(_dafny.IntOfInt64(737), (_289_v150).Dtor_cf60()))
			var _293_v152 _dafny.Map
			_ = _293_v152
			_293_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_292_v151, _256_v135)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_286_v147), 0))
			_ = _index34
			var _rhs35 _dafny.MultiSet = (func() _dafny.MultiSet {
				if (_293_v152).Contains(_292_v151) {
					return (_293_v152).Get(_292_v151).(_dafny.MultiSet)
				}
				return (_256_v135).Difference(_256_v135)
			})()
			_ = _rhs35
			var _rhs36 _dafny.Int = (_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_205_v103)).Cardinality())
			_ = _rhs36
			var _rhs37 _dafny.Sequence = _113_v38.F13
			_ = _rhs37
			var _rhs38 _dafny.Sequence = _115_v40
			_ = _rhs38
			var _rhs39 _dafny.Array = (func() _dafny.Array {
				if _222_v116.F4() {
					return _288_v149
				}
				return _288_v149
			})()
			_ = _rhs39
			var _lhs21 *GlobalState = _71_globalState
			_ = _lhs21
			var _lhs22 *C3 = _113_v38
			_ = _lhs22
			var _lhs23 _dafny.Array = _286_v147
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_286_v147), 0))
			_ = _lhs24
			_256_v135 = _rhs35
			_lhs21.F1 = _rhs36
			_lhs22.F13 = _rhs37
			_115_v40 = _rhs38
			(_lhs23).ArraySet1(_rhs39, (_lhs24).Int())
			var _294_v153 _dafny.Set
			_ = _294_v153
			_294_v153 = _dafny.SetOf(_112_v37, (_223_v117).F15(), (_223_v117).F15(), (_223_v117).F15())
			_294_v153 = ((_294_v153).Intersection(_dafny.SetOf((_223_v117).F15()))).Difference((_294_v153).Intersection(_294_v153))
			var _295_v155 _dafny.MultiSet
			_ = _295_v155
			_295_v155 = _dafny.MultiSetOf(_110_v35, (_113_v38).F12(), (_113_v38).F12(), (_113_v38).F12(), (_113_v38).F12())
			var _296_v156 _dafny.Sequence
			_ = _296_v156
			_296_v156 = _dafny.SeqOf(func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter8 := _dafny.Iterate((_295_v155).Elements()); ; {
					_compr_7, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _297_v154 _dafny.CodePoint
					_297_v154 = interface{}(_compr_7).(_dafny.CodePoint)
					if (_295_v155).Contains(_297_v154) {
						_coll7.Add(_297_v154, _287_v148.F10)
					}
				}
				return _coll7.ToMap()
			}())
			var _298_v157 _dafny.MultiSet
			_ = _298_v157
			_298_v157 = _dafny.MultiSetOf((_296_v156).Select((Companion_Default___.SafeIndex((_222_v116).F3(), _dafny.IntOfUint32((_296_v156).Cardinality()))).Uint32()).(_dafny.Map))
			var _299_v158 D10
			_ = _299_v158
			_299_v158 = Companion_D10_.Create_DC29_(!(_72_v1), (_113_v38).F12())
			var _300_v159 _dafny.Map
			_ = _300_v159
			_300_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_299_v158).Dtor_cf43(), _222_v116.F4())
			_251_cf71 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_298_v157).Contains(_300_v159) {
					return (_298_v157).Multiplicity(_300_v159)
				}
				return _107_v32
			})())
			(_287_v148).F10 = (((_dafny.Zero).Minus(Companion_Default___.Fm0(_71_globalState))).Plus(_dafny.IntOfInt64(690))).Cmp((_222_v116).F3()) != 0
		} else if _source4.Is_DC37() {
			var _301___mcc_h13 _dafny.Map = _source4.Get_().(D12_DC37).Cf64
			_ = _301___mcc_h13
			var _302___mcc_h14 bool = _source4.Get_().(D12_DC37).Cf65
			_ = _302___mcc_h14
			var _303___mcc_h15 _dafny.Int = _source4.Get_().(D12_DC37).Cf66
			_ = _303___mcc_h15
			var _304___mcc_h16 _dafny.Int = _source4.Get_().(D12_DC37).Cf67
			_ = _304___mcc_h16
			var _305_cf67 _dafny.Int = _304___mcc_h16
			_ = _305_cf67
			var _306_cf66 _dafny.Int = _303___mcc_h15
			_ = _306_cf66
			var _307_cf65 bool = _302___mcc_h14
			_ = _307_cf65
			var _308_cf64 _dafny.Map = _301___mcc_h13
			_ = _308_cf64
			_106_v31 = _106_v31
			(_71_globalState).F2 = Companion_Default___.Fm2(_107_v32, _71_globalState)
			(_222_v116).F4_set_(_222_v116.F4())
			var _309_v160 _dafny.Map
			_ = _309_v160
			_309_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v1, _113_v38.F13)
			var _310_v161 D16
			_ = _310_v161
			_310_v161 = Companion_D16_.Create_DC45_(_309_v160)
			var _311_v162 _dafny.Array
			_ = _311_v162
			var _nwElement0_8 _dafny.Map = _309_v160
			_ = _nwElement0_8
			var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(4))
			_ = _nw49
			(_nw49).ArraySet1(_nwElement0_8, 0)
			(_nw49).ArraySet1((_309_v160).Merge(_309_v160), 1)
			(_nw49).ArraySet1((_309_v160).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_307_cf65, _111_v36)).Update(false, _111_v36)), 2)
			(_nw49).ArraySet1((_310_v161).Dtor_cf78(), 3)
			_311_v162 = _nw49
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_311_v162), 0))
			_ = _index35
			(_311_v162).ArraySet1(((Companion_D16_.Create_DC45_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_72_v1), _111_v36))).Dtor_cf78()).Merge(_309_v160), (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_311_v162), 0))
			_ = _index36
			(_311_v162).ArraySet1(_309_v160, (_index36).Int())
		} else {
			var _312___mcc_h17 _dafny.Array = _source4.Get_().(D12_DC34).Cf55
			_ = _312___mcc_h17
			var _313_cf55 _dafny.Array = _312___mcc_h17
			_ = _313_cf55
			var _314_v163 _dafny.Array
			_ = _314_v163
			var _nw50 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
			_ = _nw50
			_314_v163 = _nw50
			var _315_v164 _dafny.Set
			_ = _315_v164
			_315_v164 = _dafny.SetOf(_314_v163, _314_v163, _314_v163)
			_315_v164 = (_315_v164).Union(_dafny.SetOf(_314_v163, _314_v163))
			(_71_globalState).F1 = _dafny.IntOfUint32((_113_v38.F13).Cardinality())
			var _316_v165 _dafny.MultiSet
			_ = _316_v165
			var _317_v166 _dafny.CodePoint
			_ = _317_v166
			var _318_v167 _dafny.Map
			_ = _318_v167
			var _319_v168 bool
			_ = _319_v168
			var _out5 _dafny.MultiSet
			_ = _out5
			var _out6 _dafny.CodePoint
			_ = _out6
			var _out7 _dafny.Map
			_ = _out7
			var _out8 bool
			_ = _out8
			_out5, _out6, _out7, _out8 = (_109_v34).M2((_106_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))).Int()).(bool), (_223_v117).F15(), (_231_v124).Merge(_231_v124), _72_v1, _71_globalState)
			_316_v165 = _out5
			_317_v166 = _out6
			_318_v167 = _out7
			_319_v168 = _out8
			var _320_v169 _dafny.Int
			_ = _320_v169
			var _out9 _dafny.Int
			_ = _out9
			_out9 = (_109_v34).M1(_dafny.IntOfInt64(424), _71_globalState)
			_320_v169 = _out9
		}
		var _321_v170 *C5
		_ = _321_v170
		var _nw51 *C5 = New_C5_()
		_ = _nw51
		_nw51.Ctor__((_109_v34).F9(), (_223_v117).F14(), (_223_v117).F14(), _72_v1)
		_321_v170 = _nw51
		if _222_v116.F4() {
			var _322_v171 _dafny.Array
			_ = _322_v171
			var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw52
			_322_v171 = _nw52
			var _323_v172 _dafny.Array
			_ = _323_v172
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(4))
			_ = _nw53
			_323_v172 = _nw53
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_322_v171), 0))
			_ = _index37
			(_322_v171).ArraySet1((func() _dafny.Array {
				if false {
					return _323_v172
				}
				return _323_v172
			})(), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_322_v171), 0))
			_ = _index38
			(_322_v171).ArraySet1(_323_v172, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_112_v37), 0))
			_ = _index39
			(_112_v37).ArraySet1(_dafny.IntOfInt64(185), (_index39).Int())
			var _324_v173 *C0
			_ = _324_v173
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__(_222_v116.F4(), (_109_v34).F9())
			_324_v173 = _nw54
			_324_v173 = _324_v173
			var _325_v174 T1
			_ = _325_v174
			var _nw55 *C5 = New_C5_()
			_ = _nw55
			_nw55.Ctor__((_109_v34).F9(), _107_v32, _107_v32, true)
			_325_v174 = _nw55
			var _326_v175 _dafny.Sequence
			_ = _326_v175
			_326_v175 = _dafny.SeqOf(_325_v174, _325_v174)
			var _327_v176 _dafny.Map
			_ = _327_v176
			_327_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_326_v175, (_254_v133).F8())
			_327_v176 = _327_v176
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
			_ = _index40
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
			_ = _index41
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
			_ = _index42
			var _rhs40 bool = _325_v174.F4()
			_ = _rhs40
			var _rhs41 bool = (false) || ((_324_v173).Fm5((_254_v133).F8(), (_325_v174).F3(), _110_v35, _72_v1, _71_globalState))
			_ = _rhs41
			var _rhs42 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_205_v103, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-229))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_328_v38 *C3) func(_dafny.Int) _dafny.Int {
				return func(_329_i23 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_328_v38.F13).Cardinality())
				}
			})(_113_v38)))), _205_v103)).Cardinality())
			_ = _rhs42
			var _rhs43 bool = _222_v116.F4()
			_ = _rhs43
			var _lhs25 _dafny.Array = _106_v31
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
			_ = _lhs26
			var _lhs27 _dafny.Array = _106_v31
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
			_ = _lhs28
			var _lhs29 *GlobalState = _71_globalState
			_ = _lhs29
			var _lhs30 _dafny.Array = _106_v31
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
			_ = _lhs31
			(_lhs25).ArraySet1(_rhs40, (_lhs26).Int())
			(_lhs27).ArraySet1(_rhs41, (_lhs28).Int())
			_lhs29.F1 = _rhs42
			(_lhs30).ArraySet1(_rhs43, (_lhs31).Int())
		} else {
			_72_v1 = _222_v116.F4()
			var _330_v177 *C6
			_ = _330_v177
			var _nw56 *C6 = New_C6_()
			_ = _nw56
			_nw56.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32((_111_v36).Cardinality())), (_223_v117).F14(), _222_v116.F4(), _251_cf71)
			_330_v177 = _nw56
			var _331_v178 _dafny.Array
			_ = _331_v178
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_6
			var _nw57 _dafny.Array
			_ = _nw57
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw57 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Sequence = (func(_332_v177 *C6) func(_dafny.Int) _dafny.Sequence {
					return func(_333_i24 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_332_v177).F8())
					}
				})(_330_v177)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw57 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw57).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw57).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_331_v178 = _nw57
			var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw58
			_331_v178 = _nw58
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_112_v37), 0))
			_ = _index43
			(_112_v37).ArraySet1((_222_v116).F3(), (_index43).Int())
			var _334_v179 _dafny.Map
			_ = _334_v179
			_334_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_109_v34).F9(), _dafny.SetOf(_106_v31))
			var _335_v180 _dafny.Set
			_ = _335_v180
			_335_v180 = _dafny.SetOf((_109_v34).F9(), (_109_v34).F9())
			_334_v179 = (_334_v179).Update((_109_v34).F9(), _335_v180)
		}
	} else if _source3.Is_DC38() {
		var _336___mcc_h3 _dafny.Sequence = _source3.Get_().(D13_DC38).Cf68
		_ = _336___mcc_h3
		var _337_cf68 _dafny.Sequence = _336___mcc_h3
		_ = _337_cf68
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_112_v37), 0))
		_ = _index44
		(_112_v37).ArraySet1((_222_v116).F3(), (_index44).Int())
		var _338_v181 _dafny.Map
		_ = _338_v181
		_338_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v117, _dafny.IntOfInt64(-124))
		_338_v181 = (_338_v181).Update(_223_v117, (_112_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_112_v37), 0))).Int()).(_dafny.Int))
		var _339_v182 *C7
		_ = _339_v182
		var _nw59 *C7 = New_C7_()
		_ = _nw59
		_nw59.Ctor__(_222_v116.F4(), _107_v32, (_223_v117).F14(), _72_v1)
		_339_v182 = _nw59
		var _340_v183 _dafny.Map
		_ = _340_v183
		_340_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v182, _115_v40)
		var _341_v184 _dafny.MultiSet
		_ = _341_v184
		_341_v184 = _dafny.MultiSetOf(_dafny.SeqOf(_72_v1, _72_v1), (func() _dafny.Sequence {
			if (_340_v183).Contains(_339_v182) {
				return (_340_v183).Get(_339_v182).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Update(_115_v40, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_111_v36).Cardinality()), _dafny.IntOfUint32((_115_v40).Cardinality()))).Uint32(), false)
		})(), _115_v40)
		_341_v184 = ((_341_v184).Update(_115_v40, Companion_Default___.Abs((_168_v76).Cardinality()))).Union((Companion_D17_.Create_DC48_(_dafny.MultiSetOf(_dafny.SeqOf(_222_v116.F4())))).Dtor_cf83())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))
		_ = _index45
		(_106_v31).ArraySet1(!((_106_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))).Int()).(bool)) || (_339_v182.F5), (_index45).Int())
	} else {
		var _342___mcc_h4 D13 = _source3.Get_().(D13_DC41).Cf72
		_ = _342___mcc_h4
		var _343_cf72 D13 = _342___mcc_h4
		_ = _343_cf72
		_112_v37 = (_223_v117).F15()
		(_71_globalState).F1 = (_222_v116).F3()
		(_71_globalState).F2 = _72_v1
		var _344_v185 _dafny.Set
		_ = _344_v185
		_344_v185 = _dafny.SetOf(_69_v0)
		var _rhs44 bool = !((_344_v185).IsProperSubsetOf((func() _dafny.Set {
			if (_106_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_106_v31), 0))).Int()).(bool) {
				return _344_v185
			}
			return _344_v185
		})()))
		_ = _rhs44
		var _rhs45 bool = _72_v1
		_ = _rhs45
		var _rhs46 _dafny.Sequence = _113_v38.F13
		_ = _rhs46
		var _lhs32 *GlobalState = _71_globalState
		_ = _lhs32
		var _lhs33 *C3 = _113_v38
		_ = _lhs33
		_lhs32.F2 = _rhs44
		_72_v1 = _rhs45
		_lhs33.F13 = _rhs46
	}
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(22)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(23)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(24)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v0).ArrayGet1CodePoint((_dafny.IntOfInt64(25)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(22)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(23)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(24)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(25)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_72_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v31).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_107_v32)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_108_v33).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_109_v34).F9()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_110_v35)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_111_v36.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_112_v37).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v38).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_113_v38.F13.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v39).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_115_v40, _dafny.SeqOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_116_v41).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_159_v70).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v71).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v76).Equals(_dafny.SetOf(_dafny.IntOfInt64(714))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_169_i12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_204_v102).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_205_v103, _dafny.SeqOf(_dafny.IntOfInt64(714))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_206_i15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.Zero).Int()).(*C3)).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.Zero).Int()).(*C3).F13.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.Zero).Int()).(*C3)).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.Zero).Int()).(*C3).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.Zero).Int()).(*C3)).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.One).Int()).(*C3)).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.One).Int()).(*C3).F13.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.One).Int()).(*C3)).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.One).Int()).(*C3).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.One).Int()).(*C3)).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C3)).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C3).F13.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C3)).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C3).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C3)).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C3)).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C3).F13.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C3)).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v114).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C3).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_220_v114).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C3)).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_221_v115).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v116).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_222_v116.F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v117).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_223_v117).F15()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v118).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_225_v119).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_226_v120).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_227_i18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_231_v124).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_232_v125).Dtor_cf68(), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))))
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
	Cf1 bool
	Cf2 bool
	Cf3 _dafny.CodePoint
	Cf4 _dafny.Array
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 bool, Cf3 _dafny.CodePoint, Cf4 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 _dafny.MultiSet
	Cf6 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 _dafny.MultiSet, Cf6 bool) D0 {
	return D0{D0_DC2{Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_() D0 {
	return D0{D0_DC3{}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

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

type D0_DC4 struct {
	Cf7 D0
}

func (D0_DC4) isD0() {}

func (CompanionStruct_D0_) Create_DC4_(Cf7 D0) D0 {
	return D0{D0_DC4{Cf7}}
}

func (_this D0) Is_DC4() bool {
	_, ok := _this.Get_().(D0_DC4)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, false, _dafny.CodePoint('D'), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.MultiSet {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf7() D0 {
	return _this.Get_().(D0_DC4).Cf7
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC4:
		{
			return "D0.DC4" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3 && data1.Cf4 == data2.Cf4
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Equals(data2.Cf5) && data1.Cf6 == data2.Cf6
		}
	case D0_DC3:
		{
			_, ok := other.Get_().(D0_DC3)
			return ok
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC4:
		{
			data2, ok := other.Get_().(D0_DC4)
			return ok && data1.Cf7.Equals(data2.Cf7)
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

type D1_DC6 struct {
	Cf9  T0
	Cf10 _dafny.Sequence
	Cf11 bool
	Cf12 bool
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf9 T0, Cf10 _dafny.Sequence, Cf11 bool, Cf12 bool) D1 {
	return D1{D1_DC6{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC7 struct {
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_() D1 {
	return D1{D1_DC7{}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

type D1_DC5 struct {
	Cf8 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf8 _dafny.Int) D1 {
	return D1{D1_DC5{Cf8}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC6_((T0)(nil), _dafny.EmptySeq, false, false)
}

func (_this D1) Dtor_cf9() T0 {
	return _this.Get_().(D1_DC6).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D1_DC6).Cf10
}

func (_this D1) Dtor_cf11() bool {
	return _this.Get_().(D1_DC6).Cf11
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC6).Cf12
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf9) + ", " + data.Cf10.VerbatimString(true) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf8) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && _dafny.AreEqual(data1.Cf9, data2.Cf9) && data1.Cf10.Equals(data2.Cf10) && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12
		}
	case D1_DC7:
		{
			_, ok := other.Get_().(D1_DC7)
			return ok
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0
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
	Cf13 _dafny.Array
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 _dafny.Array) D2 {
	return D2{D2_DC8{Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf13) + ")"
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
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf13 == data2.Cf13
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
	Cf15 _dafny.Int
	Cf16 _dafny.Array
	Cf17 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf15 _dafny.Int, Cf16 _dafny.Array, Cf17 _dafny.Int) D3 {
	return D3{D3_DC10{Cf15, Cf16, Cf17}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf14 _dafny.Array
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf14 _dafny.Array) D3 {
	return D3{D3_DC9{Cf14}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D3_DC9).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf14 == data2.Cf14
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

type D4_DC11 struct {
	Cf18 _dafny.Set
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf18 _dafny.Set) D4 {
	return D4{D4_DC11{Cf18}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC13 struct {
	Cf19 D4
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf19 D4) D4 {
	return D4{D4_DC13{Cf19}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_()
}

func (_this D4) Dtor_cf18() _dafny.Set {
	return _this.Get_().(D4_DC11).Cf18
}

func (_this D4) Dtor_cf19() D4 {
	return _this.Get_().(D4_DC13).Cf19
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf19) + ")"
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
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
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

type D5_DC15 struct {
	Cf21 D0
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf21 D0) D5 {
	return D5{D5_DC15{Cf21}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC14 struct {
	Cf20 _dafny.Map
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf20 _dafny.Map) D5 {
	return D5{D5_DC14{Cf20}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC16 struct {
	Cf22 D5
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf22 D5) D5 {
	return D5{D5_DC16{Cf22}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(Companion_D0_.Default())
}

func (_this D5) Dtor_cf21() D0 {
	return _this.Get_().(D5_DC15).Cf21
}

func (_this D5) Dtor_cf20() _dafny.Map {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf22() D5 {
	return _this.Get_().(D5_DC16).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf20.Equals(data2.Cf20)
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
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

type D6_DC18 struct {
	Cf24 bool
	Cf25 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf24 bool, Cf25 _dafny.Int) D6 {
	return D6{D6_DC18{Cf24, Cf25}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC19 struct {
	Cf26 _dafny.Int
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf26 _dafny.Int) D6 {
	return D6{D6_DC19{Cf26}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC17 struct {
	Cf23 _dafny.Sequence
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf23 _dafny.Sequence) D6 {
	return D6{D6_DC17{Cf23}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC18_(false, _dafny.Zero)
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC18).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf25
}

func (_this D6) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf26
}

func (_this D6) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf23
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
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

type D7_DC21 struct {
	Cf28 _dafny.Array
	Cf29 bool
	Cf30 bool
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf28 _dafny.Array, Cf29 bool, Cf30 bool) D7 {
	return D7{D7_DC21{Cf28, Cf29, Cf30}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC20 struct {
	Cf27 _dafny.Map
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf27 _dafny.Map) D7 {
	return D7{D7_DC20{Cf27}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC21_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, false)
}

func (_this D7) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D7_DC21).Cf28
}

func (_this D7) Dtor_cf29() bool {
	return _this.Get_().(D7_DC21).Cf29
}

func (_this D7) Dtor_cf30() bool {
	return _this.Get_().(D7_DC21).Cf30
}

func (_this D7) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D7_DC20).Cf27
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf27.Equals(data2.Cf27)
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
	Cf32 bool
	Cf33 bool
	Cf34 bool
	Cf35 _dafny.CodePoint
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf32 bool, Cf33 bool, Cf34 bool, Cf35 _dafny.CodePoint) D8 {
	return D8{D8_DC23{Cf32, Cf33, Cf34, Cf35}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC22 struct {
	Cf31 _dafny.Map
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf31 _dafny.Map) D8 {
	return D8{D8_DC22{Cf31}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC24 struct {
	Cf36 D8
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf36 D8) D8 {
	return D8{D8_DC24{Cf36}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(false, false, false, _dafny.CodePoint('D'))
}

func (_this D8) Dtor_cf32() bool {
	return _this.Get_().(D8_DC23).Cf32
}

func (_this D8) Dtor_cf33() bool {
	return _this.Get_().(D8_DC23).Cf33
}

func (_this D8) Dtor_cf34() bool {
	return _this.Get_().(D8_DC23).Cf34
}

func (_this D8) Dtor_cf35() _dafny.CodePoint {
	return _this.Get_().(D8_DC23).Cf35
}

func (_this D8) Dtor_cf31() _dafny.Map {
	return _this.Get_().(D8_DC22).Cf31
}

func (_this D8) Dtor_cf36() D8 {
	return _this.Get_().(D8_DC24).Cf36
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf31) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33 && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf31.Equals(data2.Cf31)
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
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

type D9_DC26 struct {
	Cf38 _dafny.Int
	Cf39 bool
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf38 _dafny.Int, Cf39 bool) D9 {
	return D9{D9_DC26{Cf38, Cf39}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC25 struct {
	Cf37 _dafny.Set
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf37 _dafny.Set) D9 {
	return D9{D9_DC25{Cf37}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC27 struct {
	Cf40 D9
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf40 D9) D9 {
	return D9{D9_DC27{Cf40}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC26_(_dafny.Zero, false)
}

func (_this D9) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D9_DC26).Cf38
}

func (_this D9) Dtor_cf39() bool {
	return _this.Get_().(D9_DC26).Cf39
}

func (_this D9) Dtor_cf37() _dafny.Set {
	return _this.Get_().(D9_DC25).Cf37
}

func (_this D9) Dtor_cf40() D9 {
	return _this.Get_().(D9_DC27).Cf40
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39 == data2.Cf39
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf37.Equals(data2.Cf37)
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf40.Equals(data2.Cf40)
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
	Cf42 bool
	Cf43 _dafny.CodePoint
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf42 bool, Cf43 _dafny.CodePoint) D10 {
	return D10{D10_DC29{Cf42, Cf43}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC30 struct {
	Cf44 _dafny.Int
	Cf45 _dafny.Array
	Cf46 bool
	Cf47 _dafny.Int
	Cf48 _dafny.Int
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf44 _dafny.Int, Cf45 _dafny.Array, Cf46 bool, Cf47 _dafny.Int, Cf48 _dafny.Int) D10 {
	return D10{D10_DC30{Cf44, Cf45, Cf46, Cf47, Cf48}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

type D10_DC28 struct {
	Cf41 _dafny.Map
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf41 _dafny.Map) D10 {
	return D10{D10_DC28{Cf41}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(false, _dafny.CodePoint('D'))
}

func (_this D10) Dtor_cf42() bool {
	return _this.Get_().(D10_DC29).Cf42
}

func (_this D10) Dtor_cf43() _dafny.CodePoint {
	return _this.Get_().(D10_DC29).Cf43
}

func (_this D10) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D10_DC30).Cf44
}

func (_this D10) Dtor_cf45() _dafny.Array {
	return _this.Get_().(D10_DC30).Cf45
}

func (_this D10) Dtor_cf46() bool {
	return _this.Get_().(D10_DC30).Cf46
}

func (_this D10) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D10_DC30).Cf47
}

func (_this D10) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D10_DC30).Cf48
}

func (_this D10) Dtor_cf41() _dafny.Map {
	return _this.Get_().(D10_DC28).Cf41
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf42 == data2.Cf42 && data1.Cf43 == data2.Cf43
		}
	case D10_DC30:
		{
			data2, ok := other.Get_().(D10_DC30)
			return ok && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45 == data2.Cf45 && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf41.Equals(data2.Cf41)
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
	Cf50 bool
	Cf51 _dafny.Int
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf50 bool, Cf51 _dafny.Int) D11 {
	return D11{D11_DC32{Cf50, Cf51}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC33 struct {
	Cf52 bool
	Cf53 _dafny.Int
	Cf54 _dafny.Int
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf52 bool, Cf53 _dafny.Int, Cf54 _dafny.Int) D11 {
	return D11{D11_DC33{Cf52, Cf53, Cf54}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

type D11_DC31 struct {
	Cf49 _dafny.Sequence
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf49 _dafny.Sequence) D11 {
	return D11{D11_DC31{Cf49}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC32_(false, _dafny.Zero)
}

func (_this D11) Dtor_cf50() bool {
	return _this.Get_().(D11_DC32).Cf50
}

func (_this D11) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D11_DC32).Cf51
}

func (_this D11) Dtor_cf52() bool {
	return _this.Get_().(D11_DC33).Cf52
}

func (_this D11) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D11_DC33).Cf53
}

func (_this D11) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D11_DC33).Cf54
}

func (_this D11) Dtor_cf49() _dafny.Sequence {
	return _this.Get_().(D11_DC31).Cf49
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ")"
		}
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf49) + ")"
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
			return ok && data1.Cf50 == data2.Cf50 && data1.Cf51.Cmp(data2.Cf51) == 0
		}
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D12_DC35 struct {
	Cf56 _dafny.Array
	Cf57 bool
	Cf58 _dafny.Sequence
	Cf59 _dafny.Sequence
	Cf60 _dafny.Int
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_(Cf56 _dafny.Array, Cf57 bool, Cf58 _dafny.Sequence, Cf59 _dafny.Sequence, Cf60 _dafny.Int) D12 {
	return D12{D12_DC35{Cf56, Cf57, Cf58, Cf59, Cf60}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

type D12_DC36 struct {
	Cf61 bool
	Cf62 bool
	Cf63 _dafny.Int
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf61 bool, Cf62 bool, Cf63 _dafny.Int) D12 {
	return D12{D12_DC36{Cf61, Cf62, Cf63}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC37 struct {
	Cf64 _dafny.Map
	Cf65 bool
	Cf66 _dafny.Int
	Cf67 _dafny.Int
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_(Cf64 _dafny.Map, Cf65 bool, Cf66 _dafny.Int, Cf67 _dafny.Int) D12 {
	return D12{D12_DC37{Cf64, Cf65, Cf66, Cf67}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

type D12_DC34 struct {
	Cf55 _dafny.Array
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf55 _dafny.Array) D12 {
	return D12{D12_DC34{Cf55}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC35_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.Zero)
}

func (_this D12) Dtor_cf56() _dafny.Array {
	return _this.Get_().(D12_DC35).Cf56
}

func (_this D12) Dtor_cf57() bool {
	return _this.Get_().(D12_DC35).Cf57
}

func (_this D12) Dtor_cf58() _dafny.Sequence {
	return _this.Get_().(D12_DC35).Cf58
}

func (_this D12) Dtor_cf59() _dafny.Sequence {
	return _this.Get_().(D12_DC35).Cf59
}

func (_this D12) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D12_DC35).Cf60
}

func (_this D12) Dtor_cf61() bool {
	return _this.Get_().(D12_DC36).Cf61
}

func (_this D12) Dtor_cf62() bool {
	return _this.Get_().(D12_DC36).Cf62
}

func (_this D12) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf63
}

func (_this D12) Dtor_cf64() _dafny.Map {
	return _this.Get_().(D12_DC37).Cf64
}

func (_this D12) Dtor_cf65() bool {
	return _this.Get_().(D12_DC37).Cf65
}

func (_this D12) Dtor_cf66() _dafny.Int {
	return _this.Get_().(D12_DC37).Cf66
}

func (_this D12) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D12_DC37).Cf67
}

func (_this D12) Dtor_cf55() _dafny.Array {
	return _this.Get_().(D12_DC34).Cf55
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC35:
		{
			return "D12.DC35" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + data.Cf59.VerbatimString(true) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ")"
		}
	case D12_DC37:
		{
			return "D12.DC37" + "(" + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC35:
		{
			data2, ok := other.Get_().(D12_DC35)
			return ok && data1.Cf56 == data2.Cf56 && data1.Cf57 == data2.Cf57 && data1.Cf58.Equals(data2.Cf58) && data1.Cf59.Equals(data2.Cf59) && data1.Cf60.Cmp(data2.Cf60) == 0
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf61 == data2.Cf61 && data1.Cf62 == data2.Cf62 && data1.Cf63.Cmp(data2.Cf63) == 0
		}
	case D12_DC37:
		{
			data2, ok := other.Get_().(D12_DC37)
			return ok && data1.Cf64.Equals(data2.Cf64) && data1.Cf65 == data2.Cf65 && data1.Cf66.Cmp(data2.Cf66) == 0 && data1.Cf67.Cmp(data2.Cf67) == 0
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf55 == data2.Cf55
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

type D13_DC39 struct {
	Cf69 _dafny.Map
}

func (D13_DC39) isD13() {}

func (CompanionStruct_D13_) Create_DC39_(Cf69 _dafny.Map) D13 {
	return D13{D13_DC39{Cf69}}
}

func (_this D13) Is_DC39() bool {
	_, ok := _this.Get_().(D13_DC39)
	return ok
}

type D13_DC40 struct {
	Cf70 D8
	Cf71 _dafny.Int
}

func (D13_DC40) isD13() {}

func (CompanionStruct_D13_) Create_DC40_(Cf70 D8, Cf71 _dafny.Int) D13 {
	return D13{D13_DC40{Cf70, Cf71}}
}

func (_this D13) Is_DC40() bool {
	_, ok := _this.Get_().(D13_DC40)
	return ok
}

type D13_DC38 struct {
	Cf68 _dafny.Sequence
}

func (D13_DC38) isD13() {}

func (CompanionStruct_D13_) Create_DC38_(Cf68 _dafny.Sequence) D13 {
	return D13{D13_DC38{Cf68}}
}

func (_this D13) Is_DC38() bool {
	_, ok := _this.Get_().(D13_DC38)
	return ok
}

type D13_DC41 struct {
	Cf72 D13
}

func (D13_DC41) isD13() {}

func (CompanionStruct_D13_) Create_DC41_(Cf72 D13) D13 {
	return D13{D13_DC41{Cf72}}
}

func (_this D13) Is_DC41() bool {
	_, ok := _this.Get_().(D13_DC41)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC39_(_dafny.EmptyMap)
}

func (_this D13) Dtor_cf69() _dafny.Map {
	return _this.Get_().(D13_DC39).Cf69
}

func (_this D13) Dtor_cf70() D8 {
	return _this.Get_().(D13_DC40).Cf70
}

func (_this D13) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D13_DC40).Cf71
}

func (_this D13) Dtor_cf68() _dafny.Sequence {
	return _this.Get_().(D13_DC38).Cf68
}

func (_this D13) Dtor_cf72() D13 {
	return _this.Get_().(D13_DC41).Cf72
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC39:
		{
			return "D13.DC39" + "(" + _dafny.String(data.Cf69) + ")"
		}
	case D13_DC40:
		{
			return "D13.DC40" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D13_DC38:
		{
			return "D13.DC38" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D13_DC41:
		{
			return "D13.DC41" + "(" + _dafny.String(data.Cf72) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC39:
		{
			data2, ok := other.Get_().(D13_DC39)
			return ok && data1.Cf69.Equals(data2.Cf69)
		}
	case D13_DC40:
		{
			data2, ok := other.Get_().(D13_DC40)
			return ok && data1.Cf70.Equals(data2.Cf70) && data1.Cf71.Cmp(data2.Cf71) == 0
		}
	case D13_DC38:
		{
			data2, ok := other.Get_().(D13_DC38)
			return ok && data1.Cf68.Equals(data2.Cf68)
		}
	case D13_DC41:
		{
			data2, ok := other.Get_().(D13_DC41)
			return ok && data1.Cf72.Equals(data2.Cf72)
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

type D14_DC42 struct {
	Cf73 _dafny.Sequence
}

func (D14_DC42) isD14() {}

func (CompanionStruct_D14_) Create_DC42_(Cf73 _dafny.Sequence) D14 {
	return D14{D14_DC42{Cf73}}
}

func (_this D14) Is_DC42() bool {
	_, ok := _this.Get_().(D14_DC42)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D14) Dtor_cf73() _dafny.Sequence {
	return _this.Get_().(D14_DC42).Cf73
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC42:
		{
			return "D14.DC42" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC42:
		{
			data2, ok := other.Get_().(D14_DC42)
			return ok && data1.Cf73.Equals(data2.Cf73)
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

type D15_DC44 struct {
	Cf75 T1
	Cf76 _dafny.CodePoint
	Cf77 _dafny.Int
}

func (D15_DC44) isD15() {}

func (CompanionStruct_D15_) Create_DC44_(Cf75 T1, Cf76 _dafny.CodePoint, Cf77 _dafny.Int) D15 {
	return D15{D15_DC44{Cf75, Cf76, Cf77}}
}

func (_this D15) Is_DC44() bool {
	_, ok := _this.Get_().(D15_DC44)
	return ok
}

type D15_DC43 struct {
	Cf74 _dafny.Sequence
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_(Cf74 _dafny.Sequence) D15 {
	return D15{D15_DC43{Cf74}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC44_((T1)(nil), _dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D15) Dtor_cf75() T1 {
	return _this.Get_().(D15_DC44).Cf75
}

func (_this D15) Dtor_cf76() _dafny.CodePoint {
	return _this.Get_().(D15_DC44).Cf76
}

func (_this D15) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D15_DC44).Cf77
}

func (_this D15) Dtor_cf74() _dafny.Sequence {
	return _this.Get_().(D15_DC43).Cf74
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC44:
		{
			return "D15.DC44" + "(" + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ")"
		}
	case D15_DC43:
		{
			return "D15.DC43" + "(" + _dafny.String(data.Cf74) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC44:
		{
			data2, ok := other.Get_().(D15_DC44)
			return ok && _dafny.AreEqual(data1.Cf75, data2.Cf75) && data1.Cf76 == data2.Cf76 && data1.Cf77.Cmp(data2.Cf77) == 0
		}
	case D15_DC43:
		{
			data2, ok := other.Get_().(D15_DC43)
			return ok && data1.Cf74.Equals(data2.Cf74)
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

type D16_DC46 struct {
	Cf79 _dafny.Sequence
	Cf80 _dafny.Set
	Cf81 bool
}

func (D16_DC46) isD16() {}

func (CompanionStruct_D16_) Create_DC46_(Cf79 _dafny.Sequence, Cf80 _dafny.Set, Cf81 bool) D16 {
	return D16{D16_DC46{Cf79, Cf80, Cf81}}
}

func (_this D16) Is_DC46() bool {
	_, ok := _this.Get_().(D16_DC46)
	return ok
}

type D16_DC45 struct {
	Cf78 _dafny.Map
}

func (D16_DC45) isD16() {}

func (CompanionStruct_D16_) Create_DC45_(Cf78 _dafny.Map) D16 {
	return D16{D16_DC45{Cf78}}
}

func (_this D16) Is_DC45() bool {
	_, ok := _this.Get_().(D16_DC45)
	return ok
}

type D16_DC47 struct {
	Cf82 D16
}

func (D16_DC47) isD16() {}

func (CompanionStruct_D16_) Create_DC47_(Cf82 D16) D16 {
	return D16{D16_DC47{Cf82}}
}

func (_this D16) Is_DC47() bool {
	_, ok := _this.Get_().(D16_DC47)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC46_(_dafny.EmptySeq, _dafny.EmptySet, false)
}

func (_this D16) Dtor_cf79() _dafny.Sequence {
	return _this.Get_().(D16_DC46).Cf79
}

func (_this D16) Dtor_cf80() _dafny.Set {
	return _this.Get_().(D16_DC46).Cf80
}

func (_this D16) Dtor_cf81() bool {
	return _this.Get_().(D16_DC46).Cf81
}

func (_this D16) Dtor_cf78() _dafny.Map {
	return _this.Get_().(D16_DC45).Cf78
}

func (_this D16) Dtor_cf82() D16 {
	return _this.Get_().(D16_DC47).Cf82
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC46:
		{
			return "D16.DC46" + "(" + data.Cf79.VerbatimString(true) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D16_DC45:
		{
			return "D16.DC45" + "(" + _dafny.String(data.Cf78) + ")"
		}
	case D16_DC47:
		{
			return "D16.DC47" + "(" + _dafny.String(data.Cf82) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC46:
		{
			data2, ok := other.Get_().(D16_DC46)
			return ok && data1.Cf79.Equals(data2.Cf79) && data1.Cf80.Equals(data2.Cf80) && data1.Cf81 == data2.Cf81
		}
	case D16_DC45:
		{
			data2, ok := other.Get_().(D16_DC45)
			return ok && data1.Cf78.Equals(data2.Cf78)
		}
	case D16_DC47:
		{
			data2, ok := other.Get_().(D16_DC47)
			return ok && data1.Cf82.Equals(data2.Cf82)
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

type D17_DC49 struct {
	Cf84 _dafny.Int
	Cf85 _dafny.Int
	Cf86 *C6
}

func (D17_DC49) isD17() {}

func (CompanionStruct_D17_) Create_DC49_(Cf84 _dafny.Int, Cf85 _dafny.Int, Cf86 *C6) D17 {
	return D17{D17_DC49{Cf84, Cf85, Cf86}}
}

func (_this D17) Is_DC49() bool {
	_, ok := _this.Get_().(D17_DC49)
	return ok
}

type D17_DC50 struct {
	Cf87 *C3
	Cf88 bool
	Cf89 *C0
	Cf90 _dafny.Int
	Cf91 bool
}

func (D17_DC50) isD17() {}

func (CompanionStruct_D17_) Create_DC50_(Cf87 *C3, Cf88 bool, Cf89 *C0, Cf90 _dafny.Int, Cf91 bool) D17 {
	return D17{D17_DC50{Cf87, Cf88, Cf89, Cf90, Cf91}}
}

func (_this D17) Is_DC50() bool {
	_, ok := _this.Get_().(D17_DC50)
	return ok
}

type D17_DC48 struct {
	Cf83 _dafny.MultiSet
}

func (D17_DC48) isD17() {}

func (CompanionStruct_D17_) Create_DC48_(Cf83 _dafny.MultiSet) D17 {
	return D17{D17_DC48{Cf83}}
}

func (_this D17) Is_DC48() bool {
	_, ok := _this.Get_().(D17_DC48)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC49_(_dafny.Zero, _dafny.Zero, (*C6)(nil))
}

func (_this D17) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D17_DC49).Cf84
}

func (_this D17) Dtor_cf85() _dafny.Int {
	return _this.Get_().(D17_DC49).Cf85
}

func (_this D17) Dtor_cf86() *C6 {
	return _this.Get_().(D17_DC49).Cf86
}

func (_this D17) Dtor_cf87() *C3 {
	return _this.Get_().(D17_DC50).Cf87
}

func (_this D17) Dtor_cf88() bool {
	return _this.Get_().(D17_DC50).Cf88
}

func (_this D17) Dtor_cf89() *C0 {
	return _this.Get_().(D17_DC50).Cf89
}

func (_this D17) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D17_DC50).Cf90
}

func (_this D17) Dtor_cf91() bool {
	return _this.Get_().(D17_DC50).Cf91
}

func (_this D17) Dtor_cf83() _dafny.MultiSet {
	return _this.Get_().(D17_DC48).Cf83
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC49:
		{
			return "D17.DC49" + "(" + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ")"
		}
	case D17_DC50:
		{
			return "D17.DC50" + "(" + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ")"
		}
	case D17_DC48:
		{
			return "D17.DC48" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC49:
		{
			data2, ok := other.Get_().(D17_DC49)
			return ok && data1.Cf84.Cmp(data2.Cf84) == 0 && data1.Cf85.Cmp(data2.Cf85) == 0 && data1.Cf86 == data2.Cf86
		}
	case D17_DC50:
		{
			data2, ok := other.Get_().(D17_DC50)
			return ok && data1.Cf87 == data2.Cf87 && data1.Cf88 == data2.Cf88 && data1.Cf89 == data2.Cf89 && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91 == data2.Cf91
		}
	case D17_DC48:
		{
			data2, ok := other.Get_().(D17_DC48)
			return ok && data1.Cf83.Equals(data2.Cf83)
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

type D18_DC51 struct {
	Cf92 _dafny.Map
}

func (D18_DC51) isD18() {}

func (CompanionStruct_D18_) Create_DC51_(Cf92 _dafny.Map) D18 {
	return D18{D18_DC51{Cf92}}
}

func (_this D18) Is_DC51() bool {
	_, ok := _this.Get_().(D18_DC51)
	return ok
}

func (CompanionStruct_D18_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D18) Dtor_cf92() _dafny.Map {
	return _this.Get_().(D18_DC51).Cf92
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC51:
		{
			return "D18.DC51" + "(" + _dafny.String(data.Cf92) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC51:
		{
			data2, ok := other.Get_().(D18_DC51)
			return ok && data1.Cf92.Equals(data2.Cf92)
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

type D19_DC52 struct {
	Cf93 _dafny.Map
}

func (D19_DC52) isD19() {}

func (CompanionStruct_D19_) Create_DC52_(Cf93 _dafny.Map) D19 {
	return D19{D19_DC52{Cf93}}
}

func (_this D19) Is_DC52() bool {
	_, ok := _this.Get_().(D19_DC52)
	return ok
}

func (CompanionStruct_D19_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D19) Dtor_cf93() _dafny.Map {
	return _this.Get_().(D19_DC52).Cf93
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC52:
		{
			return "D19.DC52" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC52:
		{
			data2, ok := other.Get_().(D19_DC52)
			return ok && data1.Cf93.Equals(data2.Cf93)
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

// Definition of trait T0
type T0 interface {
	String() string
	F4() bool
	F4_set_(value bool)
	F3() _dafny.Int
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
	F4() bool
	F4_set_(value bool)
	F3() _dafny.Int
	M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int
	F7() _dafny.Int
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
	F1  _dafny.Int
	F2  bool
	_f0 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this.F2 = false
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Int, f2 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F10  bool
	_f11 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F10 = false
	_this._f11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) Ctor__(f10 bool, f11 _dafny.Array) {
	{
		(_this).F10 = f10
		(_this)._f11 = f11
	}
}
func (_this *C0) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cdnmimuc")).Cardinality())).Cmp((_dafny.SetOf(!(_this.F10))).Cardinality()) != 0
	}
}
func (_this *C0) F11() _dafny.Array {
	{
		return _this._f11
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f4 bool
	_f7 _dafny.Int
	_f3 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f4 = false
	_this._f7 = _dafny.Zero
	_this._f3 = _dafny.Zero
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

func (_this *C1) F4() bool {
	return _this._f4
}
func (_this *C1) F4_set_(value bool) {
	_this._f4 = value
}
func (_this *C1) F7() _dafny.Int {
	return _this._f7
}
func (_this *C1) F3() _dafny.Int {
	return _this._f3
}
func (_this *C1) Ctor__(f7 _dafny.Int, f3 _dafny.Int, f4 bool) {
	{
		(_this)._f7 = f7
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C1) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		if !(_this.F4()) {
			var _345_v0 _dafny.Array
			_ = _345_v0
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw60
			_345_v0 = _nw60
			var _346_v1 *C0
			_ = _346_v1
			var _nw61 *C0 = New_C0_()
			_ = _nw61
			_nw61.Ctor__(!(Companion_Default___.Fm2((_this).F7(), globalState)), _345_v0)
			_346_v1 = _nw61
			var _347_v2 *C0
			_ = _347_v2
			var _nw62 *C0 = New_C0_()
			_ = _nw62
			_nw62.Ctor__(_346_v1.F10, (_346_v1).F11())
			_347_v2 = _nw62
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_345_v0), 0))
			_ = _index46
			(_345_v0).ArraySet1((func() bool {
				if _this.F4() {
					return _347_v2.F10
				}
				return _this.F4()
			})(), (_index46).Int())
			var _348_v3 _dafny.Sequence
			_ = _348_v3
			_348_v3 = _dafny.SeqOf(_346_v1.F10, _346_v1.F10)
			var _349_v4 _dafny.Sequence
			_ = _349_v4
			_349_v4 = _dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm19(p0, globalState)).Cardinality()))
			var _350_v5 _dafny.Sequence
			_ = _350_v5
			_350_v5 = _dafny.UnicodeSeqOfUtf8Bytes("hxpnn")
			var _351_v6 D3
			_ = _351_v6
			_351_v6 = Companion_D3_.Create_DC10_(_dafny.IntOfUint32((_348_v3).Cardinality()), (_346_v1).F11(), (_349_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_349_v4, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_350_v5).Cardinality())), _dafny.IntOfUint32((_349_v4).Cardinality()))).Uint32(), p0)).Cardinality()), _dafny.IntOfUint32((_349_v4).Cardinality()))).Uint32()).(_dafny.Int))
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_345_v0), 0))
			_ = _index47
			var _rhs47 _dafny.Int = (_351_v6).Dtor_cf17()
			_ = _rhs47
			var _rhs48 bool = !(_this.F4())
			_ = _rhs48
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			var _lhs35 _dafny.Array = _345_v0
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_345_v0), 0))
			_ = _lhs36
			_lhs34.F1 = _rhs47
			(_lhs35).ArraySet1(_rhs48, (_lhs36).Int())
			var _352_v7 _dafny.CodePoint
			_ = _352_v7
			_352_v7 = _dafny.CodePoint('j')
			_352_v7 = _dafny.CodePoint('c')
			(globalState).F1 = (_this).F7()
		} else {
			var _353_v8 _dafny.Sequence
			_ = _353_v8
			_353_v8 = _dafny.UnicodeSeqOfUtf8Bytes("xpxqoyf")
			var _354_v9 _dafny.Array
			_ = _354_v9
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_7
			var _nw63 _dafny.Array
			_ = _nw63
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw63 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = func(_355_i0 _dafny.Int) bool {
					return _this.F4()
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw63 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw63).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw63).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_354_v9 = _nw63
			var _356_v10 D3
			_ = _356_v10
			_356_v10 = Companion_D3_.Create_DC10_(_dafny.IntOfUint32((_353_v8).Cardinality()), _354_v9, (_this).F7())
			_353_v8 = Companion_Default___.Fm20(((_356_v10).Dtor_cf15()).Times((_this).F3()), globalState)
			(globalState).F1 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F3()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _this.F4())).Cardinality())
			if _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ukt"), Companion_Default___.Fm20((_dafny.Zero).Minus(_dafny.IntOfUint32((_353_v8).Cardinality())), globalState)) {
				(globalState).F1 = (_this).F3()
				var _357_v11 _dafny.MultiSet
				_ = _357_v11
				_357_v11 = _dafny.MultiSetOf(_this.F4())
				_357_v11 = _357_v11
				(globalState).F2 = (func() bool {
					if _this.F4() {
						return (p0).Cmp((_dafny.Zero).Minus((_this).F7())) > 0
					}
					return false
				})()
				(globalState).F1 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(791), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _this.F4())).Cardinality())
				_356_v10 = _356_v10
			} else {
				(globalState).F2 = Companion_Default___.Fm2(Companion_Default___.SafeModuloInt(p0, (_this).F3()), globalState)
				var _358_v12 D0
				_ = _358_v12
				_358_v12 = Companion_D0_.Create_DC0_(_this.F4())
				var _359_v13 _dafny.MultiSet
				_ = _359_v13
				_359_v13 = _dafny.MultiSetOf(_this.F4(), _this.F4(), _this.F4(), (_358_v12).Dtor_cf0(), _this.F4())
				var _360_v14 _dafny.Map
				_ = _360_v14
				_360_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_353_v8, _359_v13)
				_360_v14 = (_360_v14).Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-476))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}(func(_361_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})), (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-476))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}(func(_362_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				}))).Cardinality()))).Uint32(), _dafny.CodePoint('q')), _359_v13)
				(globalState).F1 = (_dafny.Zero).Minus((_this).F7())
				var _363_v15 _dafny.Array
				_ = _363_v15
				var _nwElement0_9 _dafny.Int = p0
				_ = _nwElement0_9
				var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(11))
				_ = _nw64
				(_nw64).ArraySet1(_nwElement0_9, 0)
				(_nw64).ArraySet1((_this).F3(), 1)
				(_nw64).ArraySet1((_this).F7(), 2)
				(_nw64).ArraySet1(Companion_Default___.Fm0(globalState), 3)
				(_nw64).ArraySet1((_this).F3(), 4)
				(_nw64).ArraySet1(_dafny.IntOfUint32((_353_v8).Cardinality()), 5)
				(_nw64).ArraySet1(p0, 6)
				(_nw64).ArraySet1((_dafny.Zero).Minus((_this).F7()), 7)
				(_nw64).ArraySet1((_this).F7(), 8)
				(_nw64).ArraySet1(p0, 9)
				(_nw64).ArraySet1(_dafny.IntOfInt64(-233), 10)
				_363_v15 = _nw64
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_363_v15), 0))
				_ = _index48
				(_363_v15).ArraySet1(((_dafny.Zero).Minus((_this).F7())).Times(p0), (_index48).Int())
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_363_v15), 0))
				_ = _index49
				(_363_v15).ArraySet1(Companion_Default___.Fm0(globalState), (_index49).Int())
				(globalState).F2 = true
			}
			var _364_v16 _dafny.Set
			_ = _364_v16
			_364_v16 = _dafny.SetOf(p0)
			r0 = (_364_v16).Cardinality()
			_354_v9 = _354_v9
		}
		var _hi3 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(338), (_this).F3())
		_ = _hi3
		for _365_i2 := _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(40))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_378_i3 _dafny.Int) _dafny.Int {
			return (_this).F3()
		}))).Cardinality()); _365_i2.Cmp(_hi3) < 0; _365_i2 = _365_i2.Plus(_dafny.One) {
			var _366_v17 _dafny.Array
			_ = _366_v17
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_8
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = func(_367_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kjyvte"), _dafny.UnicodeSeqOfUtf8Bytes("ydnlt"))
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw65 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw65).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw65).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_366_v17 = _nw65
			var _368_v18 _dafny.Sequence
			_ = _368_v18
			_368_v18 = _dafny.UnicodeSeqOfUtf8Bytes("wjj")
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_366_v17), 0))
			_ = _index50
			(_366_v17).ArraySet1(_368_v18, (_index50).Int())
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_366_v17), 0))
			_ = _index51
			(_366_v17).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hyb"), (_index51).Int())
			var _369_v19 _dafny.Array
			_ = _369_v19
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_9
			var _nw66 _dafny.Array
			_ = _nw66
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw66 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = func(_370_i5 _dafny.Int) _dafny.Int {
					return (_370_i5).Times(_dafny.IntOfUint32((_dafny.SeqOf(_this.F4())).Cardinality()))
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw66 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw66).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw66).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_369_v19 = _nw66
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_369_v19), 0))
			_ = _index52
			(_369_v19).ArraySet1(_365_i2, (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_369_v19), 0))
			_ = _index53
			(_369_v19).ArraySet1((_this).F3(), (_index53).Int())
			r0 = Companion_Default___.Fm0(globalState)
			var _371_v20 _dafny.Array
			_ = _371_v20
			var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
			_ = _nw67
			_371_v20 = _nw67
			var _372_v21 _dafny.CodePoint
			_ = _372_v21
			_372_v21 = _dafny.CodePoint('u')
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_371_v20), 0))
			_ = _index54
			(_371_v20).ArraySet1CodePoint(_372_v21, (_index54).Int())
			var _373_v22 _dafny.Map
			_ = _373_v22
			_373_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F4())
			var _374_v23 _dafny.Sequence
			_ = _374_v23
			_374_v23 = _dafny.SeqOf(_373_v22)
			var _375_v24 _dafny.Sequence
			_ = _375_v24
			_375_v24 = _dafny.SeqOf(Companion_Default___.Fm21(_374_v23, globalState))
			var _376_v25 _dafny.Sequence
			_ = _376_v25
			_376_v25 = _dafny.SeqOf(_this.F4())
			var _377_v26 _dafny.Sequence
			_ = _377_v26
			_377_v26 = _dafny.SeqOf(_dafny.IntOfUint32((_376_v25).Cardinality()))
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_371_v20), 0))
			_ = _index55
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_369_v19), 0))
			_ = _index56
			var _rhs49 bool = _this.F4()
			_ = _rhs49
			var _rhs50 _dafny.CodePoint = _372_v21
			_ = _rhs50
			var _rhs51 _dafny.Int = ((_375_v24).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F7(), _dafny.IntOfUint32((_377_v26).Cardinality())), _dafny.IntOfUint32((_375_v24).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
			_ = _rhs51
			var _lhs37 *GlobalState = globalState
			_ = _lhs37
			var _lhs38 _dafny.Array = _371_v20
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_371_v20), 0))
			_ = _lhs39
			var _lhs40 _dafny.Array = _369_v19
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_369_v19), 0))
			_ = _lhs41
			_lhs37.F2 = _rhs49
			(_lhs38).ArraySet1CodePoint(_rhs50, (_lhs39).Int())
			(_lhs40).ArraySet1(_rhs51, (_lhs41).Int())
		}
		var _379_v27 _dafny.Sequence
		_ = _379_v27
		_379_v27 = _dafny.UnicodeSeqOfUtf8Bytes("s")
		_379_v27 = _dafny.Companion_Sequence_.Concatenate(_379_v27, _dafny.Companion_Sequence_.Concatenate(_379_v27, _dafny.UnicodeSeqOfUtf8Bytes("asujtxoi")))
		var _380_v28 _dafny.Map
		_ = _380_v28
		_380_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_379_v27).Cardinality()), _dafny.IntOfInt64(235))
		var _381_v29 _dafny.Sequence
		_ = _381_v29
		_381_v29 = _dafny.SeqOf(_380_v28, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _dafny.IntOfUint32((_379_v27).Cardinality())), _380_v28)
		var _hi4 _dafny.Int = ((_this).F3()).Plus(_dafny.IntOfUint32((_381_v29).Cardinality()))
		_ = _hi4
		for _382_i6 := Companion_Default___.SafeDivisionInt((_this).F7(), _dafny.IntOfInt64(357)); _382_i6.Cmp(_hi4) < 0; _382_i6 = _382_i6.Plus(_dafny.One) {
			var _383_v30 _dafny.MultiSet
			_ = _383_v30
			_383_v30 = _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)))
			(_this).F4_set_((func() bool {
				if (_383_v30).IsSubsetOf(_383_v30) {
					return _this.F4()
				}
				return ((_this).F7()).Cmp(Companion_Default___.Fm0(globalState)) < 0
			})())
			if Companion_Default___.Fm2((_this).F3(), globalState) {
				var _384_v31 _dafny.Array
				_ = _384_v31
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_10
				var _nw68 _dafny.Array
				_ = _nw68
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw68 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = func(_385_i7 _dafny.Int) bool {
						return _this.F4()
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw68 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw68).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw68).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_384_v31 = _nw68
				var _386_v32 _dafny.Array
				_ = _386_v32
				var _nwElement0_10 _dafny.Array = _384_v31
				_ = _nwElement0_10
				var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(11))
				_ = _nw69
				(_nw69).ArraySet1(_nwElement0_10, 0)
				(_nw69).ArraySet1(_384_v31, 1)
				(_nw69).ArraySet1(_384_v31, 2)
				(_nw69).ArraySet1(_384_v31, 3)
				(_nw69).ArraySet1(_384_v31, 4)
				(_nw69).ArraySet1(_384_v31, 5)
				(_nw69).ArraySet1(_384_v31, 6)
				(_nw69).ArraySet1((func() _dafny.Array {
					if _this.F4() {
						return _384_v31
					}
					return _384_v31
				})(), 7)
				(_nw69).ArraySet1(_384_v31, 8)
				(_nw69).ArraySet1(_384_v31, 9)
				(_nw69).ArraySet1(_384_v31, 10)
				_386_v32 = _nw69
				var _387_v33 _dafny.Array
				_ = _387_v33
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
				_ = _nw70
				_387_v33 = _nw70
				var _388_v34 _dafny.CodePoint
				_ = _388_v34
				_388_v34 = _dafny.CodePoint('r')
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_387_v33), 0))
				_ = _index57
				(_387_v33).ArraySet1CodePoint(_388_v34, (_index57).Int())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_387_v33), 0))
				_ = _index58
				var _rhs52 _dafny.Array = _386_v32
				_ = _rhs52
				var _rhs53 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(globalState), _382_i6)
				_ = _rhs53
				var _rhs54 bool = _this.F4()
				_ = _rhs54
				var _rhs55 bool = !(_this.F4())
				_ = _rhs55
				var _rhs56 _dafny.CodePoint = _388_v34
				_ = _rhs56
				var _lhs42 *GlobalState = globalState
				_ = _lhs42
				var _lhs43 *C1 = _this
				_ = _lhs43
				var _lhs44 *C1 = _this
				_ = _lhs44
				var _lhs45 _dafny.Array = _387_v33
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_387_v33), 0))
				_ = _lhs46
				_386_v32 = _rhs52
				_lhs42.F1 = _rhs53
				_lhs43.F4_set_(_rhs54)
				_lhs44.F4_set_(_rhs55)
				(_lhs45).ArraySet1CodePoint(_rhs56, (_lhs46).Int())
				var _389_v35 _dafny.Sequence
				_ = _389_v35
				_389_v35 = _dafny.SeqOf(_379_v27)
				var _390_v36 _dafny.Map
				_ = _390_v36
				_390_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _389_v35)
				_390_v36 = (_390_v36).Update((_dafny.Zero).Minus(p0), _dafny.Companion_Sequence_.Concatenate(_389_v35, _389_v35))
				var _391_v37 _dafny.Array
				_ = _391_v37
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_11
				var _nw71 _dafny.Array
				_ = _nw71
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw71 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.Int = (func(_392_i6 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_393_i8 _dafny.Int) _dafny.Int {
							return (_393_i8).Minus(_392_i6)
						}
					})(_382_i6)
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
				_391_v37 = _nw71
				_391_v37 = _391_v37
				(globalState).F1 = (_this).F3()
				var _394_v38 _dafny.Sequence
				_ = _394_v38
				_394_v38 = _dafny.SeqOf(_this.F4(), _this.F4(), _this.F4())
				var _395_v39 _dafny.MultiSet
				_ = _395_v39
				_395_v39 = _dafny.MultiSetOf((_387_v33).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_387_v33), 0))).Int()))
				var _396_v40 D0
				_ = _396_v40
				_396_v40 = Companion_D0_.Create_DC2_(_395_v39, _this.F4())
				var _397_v41 D0
				_ = _397_v41
				_397_v41 = Companion_D0_.Create_DC2_((_396_v40).Dtor_cf5(), _this.F4())
				var _398_v42 *C0
				_ = _398_v42
				var _nw72 *C0 = New_C0_()
				_ = _nw72
				_nw72.Ctor__(!((_396_v40).Dtor_cf6()), _384_v31)
				_398_v42 = _nw72
				var _399_v43 D6
				_ = _399_v43
				_399_v43 = Companion_D6_.Create_DC17_(_394_v38)
				var _rhs57 _dafny.Int = (_this).F3()
				_ = _rhs57
				var _rhs58 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_399_v43).Dtor_cf23(), _394_v38)
				_ = _rhs58
				var _rhs59 bool = _this.F4()
				_ = _rhs59
				var _rhs60 _dafny.Int = p0
				_ = _rhs60
				var _rhs61 *C0 = _398_v42
				_ = _rhs61
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				var _lhs48 *C1 = _this
				_ = _lhs48
				var _lhs49 *GlobalState = globalState
				_ = _lhs49
				_lhs47.F1 = _rhs57
				_394_v38 = _rhs58
				_lhs48.F4_set_(_rhs59)
				_lhs49.F1 = _rhs60
				_398_v42 = _rhs61
			} else {
				var _400_v44 _dafny.Array
				_ = _400_v44
				var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
				_ = _nw73
				_400_v44 = _nw73
				_400_v44 = _400_v44
				(_this).F4_set_(_dafny.Companion_Sequence_.IsPrefixOf(_379_v27, _379_v27))
				var _401_v45 _dafny.CodePoint
				_ = _401_v45
				_401_v45 = _dafny.CodePoint('o')
				_401_v45 = _401_v45
				(globalState).F2 = Companion_Default___.Fm2(_dafny.IntOfInt64(-802), globalState)
				var _402_v46 _dafny.Array
				_ = _402_v46
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw74
				_402_v46 = _nw74
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_402_v46), 0))
				_ = _index59
				(_402_v46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_379_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(940))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}((func(_403_v45 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_404_i9 _dafny.Int) _dafny.CodePoint {
						return _403_v45
					}
				})(_401_v45)))), (_index59).Int())
				var _405_v47 D1
				_ = _405_v47
				_405_v47 = Companion_D1_.Create_DC5_(_382_i6)
				var _406_v48 _dafny.Map
				_ = _406_v48
				_406_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_405_v47, _dafny.IntOfInt64(220))
				var _407_v49 _dafny.Map
				_ = _407_v49
				_407_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F4())
				var _408_v50 _dafny.Map
				_ = _408_v50
				_408_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_407_v49, _382_i6)
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_402_v46), 0))
				_ = _index60
				var _rhs62 _dafny.Int = (func() _dafny.Int {
					if (_406_v48).Contains(_405_v47) {
						return (_406_v48).Get(_405_v47).(_dafny.Int)
					}
					return (_408_v50).Cardinality()
				})()
				_ = _rhs62
				var _rhs63 _dafny.Sequence = Companion_Default___.Fm20((_this).F3(), globalState)
				_ = _rhs63
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				var _lhs51 _dafny.Array = _402_v46
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_402_v46), 0))
				_ = _lhs52
				_lhs50.F1 = _rhs62
				(_lhs51).ArraySet1(_rhs63, (_lhs52).Int())
			}
			(globalState).F2 = _this.F4()
			(globalState).F2 = true
		}
		var _409_i10 _dafny.Int
		_ = _409_i10
		_409_i10 = _dafny.Zero
		{
			for !(_this.F4()) || (_this.F4()) {
				{
					if (_409_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_409_i10 = (_409_i10).Plus(_dafny.One)
					var _410_v51 _dafny.Array
					_ = _410_v51
					var _nwElement0_11 bool = _this.F4()
					_ = _nwElement0_11
					var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(29))
					_ = _nw75
					(_nw75).ArraySet1(_nwElement0_11, 0)
					(_nw75).ArraySet1(_this.F4(), 1)
					(_nw75).ArraySet1(_this.F4(), 2)
					(_nw75).ArraySet1(_this.F4(), 3)
					(_nw75).ArraySet1(true, 4)
					(_nw75).ArraySet1(_this.F4(), 5)
					(_nw75).ArraySet1(_this.F4(), 6)
					(_nw75).ArraySet1(_this.F4(), 7)
					(_nw75).ArraySet1(_this.F4(), 8)
					(_nw75).ArraySet1(_this.F4(), 9)
					(_nw75).ArraySet1(_this.F4(), 10)
					(_nw75).ArraySet1(_this.F4(), 11)
					(_nw75).ArraySet1(_this.F4(), 12)
					(_nw75).ArraySet1(_this.F4(), 13)
					(_nw75).ArraySet1(_this.F4(), 14)
					(_nw75).ArraySet1(true, 15)
					(_nw75).ArraySet1(_this.F4(), 16)
					(_nw75).ArraySet1(_this.F4(), 17)
					(_nw75).ArraySet1(_this.F4(), 18)
					(_nw75).ArraySet1(_this.F4(), 19)
					(_nw75).ArraySet1(_this.F4(), 20)
					(_nw75).ArraySet1(true, 21)
					(_nw75).ArraySet1(Companion_Default___.Fm2((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_379_v27).Cardinality()))), globalState), 22)
					(_nw75).ArraySet1(_this.F4(), 23)
					(_nw75).ArraySet1(_this.F4(), 24)
					(_nw75).ArraySet1(_this.F4(), 25)
					(_nw75).ArraySet1(Companion_Default___.Fm2(p0, globalState), 26)
					(_nw75).ArraySet1(_this.F4(), 27)
					(_nw75).ArraySet1(true, 28)
					_410_v51 = _nw75
					var _411_v52 *C0
					_ = _411_v52
					var _nw76 *C0 = New_C0_()
					_ = _nw76
					_nw76.Ctor__(_this.F4(), _410_v51)
					_411_v52 = _nw76
					var _412_v53 _dafny.Set
					_ = _412_v53
					_412_v53 = _dafny.SetOf(p0, (_this).F7(), (_this).F3())
					var _413_v54 _dafny.Map
					_ = _413_v54
					_413_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F7()).Cmp(p0) < 0, _379_v27)
					var _414_v55 _dafny.CodePoint
					_ = _414_v55
					_414_v55 = _dafny.CodePoint('f')
					var _rhs64 _dafny.Sequence = (func() _dafny.Sequence {
						if (_413_v54).Contains((func() bool {
							if Companion_Default___.Fm2((_this).F7(), globalState) {
								return (_411_v52).Fm5(p0, (_this).F7(), _414_v55, _411_v52.F10, globalState)
							}
							return _this.F4()
						})()) {
							return (_413_v54).Get((func() bool {
								if Companion_Default___.Fm2((_this).F7(), globalState) {
									return (_411_v52).Fm5(p0, (_this).F7(), _414_v55, _411_v52.F10, globalState)
								}
								return _this.F4()
							})()).(_dafny.Sequence)
						}
						return _379_v27
					})()
					_ = _rhs64
					var _rhs65 _dafny.Set = _412_v53
					_ = _rhs65
					var _rhs66 _dafny.Map = ((_381_v29).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_381_v29).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_380_v28)
					_ = _rhs66
					var _rhs67 _dafny.Array = (func() _dafny.Array {
						if _this.F4() {
							return _410_v51
						}
						return _410_v51
					})()
					_ = _rhs67
					_379_v27 = _rhs64
					_412_v53 = _rhs65
					_380_v28 = _rhs66
					_410_v51 = _rhs67
					var _415_v56 _dafny.Sequence
					_ = _415_v56
					_415_v56 = _dafny.SeqOf(_411_v52.F10, _this.F4())
					var _416_v57 _dafny.Array
					_ = _416_v57
					var _nwElement0_12 _dafny.Sequence = _415_v56
					_ = _nwElement0_12
					var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(4))
					_ = _nw77
					(_nw77).ArraySet1(_nwElement0_12, 0)
					(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_415_v56, _dafny.SeqOf(_411_v52.F10)), 1)
					(_nw77).ArraySet1(_415_v56, 2)
					(_nw77).ArraySet1(_415_v56, 3)
					_416_v57 = _nw77
					var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_416_v57), 0))
					_ = _index61
					(_416_v57).ArraySet1(_415_v56, (_index61).Int())
					var _417_v58 _dafny.Array
					_ = _417_v58
					var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
					_ = _nw78
					_417_v58 = _nw78
					var _418_v59 _dafny.Set
					_ = _418_v59
					_418_v59 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("vfxj"))
					var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_417_v58), 0))
					_ = _index62
					(_417_v58).ArraySet1(((_this).F3()).Times((_418_v59).Cardinality()), (_index62).Int())
					var _419_v60 _dafny.Sequence
					_ = _419_v60
					_419_v60 = _dafny.SeqOf(_412_v53, _412_v53)
					var _420_v61 _dafny.Map
					_ = _420_v61
					_420_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_410_v51, _414_v55)
					var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_416_v57), 0))
					_ = _index63
					var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_417_v58), 0))
					_ = _index64
					var _rhs68 _dafny.Sequence = _415_v56
					_ = _rhs68
					var _rhs69 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_419_v60, _419_v60)).Cardinality())
					_ = _rhs69
					var _rhs70 _dafny.Int = (_420_v61).Cardinality()
					_ = _rhs70
					var _rhs71 _dafny.Set = (_dafny.SetOf(p0, p0)).Intersection(_412_v53)
					_ = _rhs71
					var _lhs53 _dafny.Array = _416_v57
					_ = _lhs53
					var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_416_v57), 0))
					_ = _lhs54
					var _lhs55 *GlobalState = globalState
					_ = _lhs55
					var _lhs56 _dafny.Array = _417_v58
					_ = _lhs56
					var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_417_v58), 0))
					_ = _lhs57
					(_lhs53).ArraySet1(_rhs68, (_lhs54).Int())
					_lhs55.F1 = _rhs69
					(_lhs56).ArraySet1(_rhs70, (_lhs57).Int())
					_412_v53 = _rhs71
					(globalState).F1 = (_this).F3()
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _421_v62 _dafny.Set
		_ = _421_v62
		_421_v62 = _dafny.SetOf(!(_this.F4()))
		var _422_i11 _dafny.Int
		_ = _422_i11
		_422_i11 = _dafny.Zero
		{
			for (_421_v62).IsProperSubsetOf((Companion_Default___.Fm22(_this.F4(), _this.F4(), globalState)).Intersection(_dafny.SetOf(_this.F4()))) {
				{
					if (_422_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_422_i11 = (_422_i11).Plus(_dafny.One)
					var _423_v63 _dafny.Array
					_ = _423_v63
					var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
					_ = _nw79
					_423_v63 = _nw79
					var _424_v64 _dafny.Array
					_ = _424_v64
					var _nwElement0_13 _dafny.Array = _423_v63
					_ = _nwElement0_13
					var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(11))
					_ = _nw80
					(_nw80).ArraySet1(_nwElement0_13, 0)
					(_nw80).ArraySet1(_423_v63, 1)
					(_nw80).ArraySet1(_423_v63, 2)
					(_nw80).ArraySet1(_423_v63, 3)
					(_nw80).ArraySet1(_423_v63, 4)
					(_nw80).ArraySet1(_423_v63, 5)
					(_nw80).ArraySet1(_423_v63, 6)
					(_nw80).ArraySet1(_423_v63, 7)
					(_nw80).ArraySet1(_423_v63, 8)
					(_nw80).ArraySet1((func() _dafny.Array {
						if _this.F4() {
							return _423_v63
						}
						return _423_v63
					})(), 9)
					(_nw80).ArraySet1(_423_v63, 10)
					_424_v64 = _nw80
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_424_v64), 0))
					_ = _index65
					(_424_v64).ArraySet1(_423_v63, (_index65).Int())
					var _425_v65 _dafny.Array
					_ = _425_v65
					var _len0_12 _dafny.Int = _dafny.One
					_ = _len0_12
					var _nw81 _dafny.Array
					_ = _nw81
					if _len0_12.Cmp(_dafny.Zero) == 0 {
						_nw81 = _dafny.NewArray(_len0_12)
					} else {
						var _init12 func(_dafny.Int) _dafny.Map = (func(_426_v27 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
							return func(_427_i12 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _426_v27)
							}
						})(_379_v27)
						_ = _init12
						var _element0_12 = _init12(_dafny.Zero)
						_ = _element0_12
						_nw81 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
						(_nw81).ArraySet1(_element0_12, 0)
						var _nativeLen0_12 = (_len0_12).Int()
						_ = _nativeLen0_12
						for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
							(_nw81).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
						}
					}
					_425_v65 = _nw81
					var _428_v66 D7
					_ = _428_v66
					_428_v66 = Companion_D7_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _dafny.UnicodeSeqOfUtf8Bytes("tfjfsyru")))
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_425_v65), 0))
					_ = _index66
					(_425_v65).ArraySet1((_428_v66).Dtor_cf27(), (_index66).Int())
					var _429_v67 _dafny.CodePoint
					_ = _429_v67
					_429_v67 = _dafny.CodePoint('y')
					var _430_v68 _dafny.Map
					_ = _430_v68
					_430_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _379_v27)
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_424_v64), 0))
					_ = _index67
					var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_425_v65), 0))
					_ = _index68
					var _rhs72 _dafny.Array = _423_v63
					_ = _rhs72
					var _rhs73 _dafny.Map = _430_v68
					_ = _rhs73
					var _rhs74 _dafny.Int = (_this).F3()
					_ = _rhs74
					var _rhs75 _dafny.CodePoint = _429_v67
					_ = _rhs75
					var _lhs58 _dafny.Array = _424_v64
					_ = _lhs58
					var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_424_v64), 0))
					_ = _lhs59
					var _lhs60 _dafny.Array = _425_v65
					_ = _lhs60
					var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_425_v65), 0))
					_ = _lhs61
					var _lhs62 *GlobalState = globalState
					_ = _lhs62
					(_lhs58).ArraySet1(_rhs72, (_lhs59).Int())
					(_lhs60).ArraySet1(_rhs73, (_lhs61).Int())
					_lhs62.F1 = _rhs74
					_429_v67 = _rhs75
					r0 = Companion_Default___.Fm0(globalState)
					var _431_v69 _dafny.Array
					_ = _431_v69
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_13
					var _nw82 _dafny.Array
					_ = _nw82
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw82 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) bool = func(_432_i13 _dafny.Int) bool {
							return _this.F4()
						}
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw82 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw82).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw82).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_431_v69 = _nw82
					var _433_v70 _dafny.Map
					_ = _433_v70
					_433_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F4())
					var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
					_ = _index69
					(_431_v69).ArraySet1((func() bool {
						if (_433_v70).Contains(_this.F4()) {
							return (_433_v70).Get(_this.F4()).(bool)
						}
						return _this.F4()
					})(), (_index69).Int())
					var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
					_ = _index70
					(_431_v69).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-39))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}((func(_434_v67 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_435_i14 _dafny.Int) _dafny.CodePoint {
							return _434_v67
						}
					})(_429_v67))), _429_v67), (_index70).Int())
					var _436_v71 D0
					_ = _436_v71
					_436_v71 = Companion_D0_.Create_DC3_()
					var _source5 D0 = _436_v71
					_ = _source5
					if _source5.Is_DC1() {
						var _437___mcc_h0 bool = _source5.Get_().(D0_DC1).Cf1
						_ = _437___mcc_h0
						var _438___mcc_h1 bool = _source5.Get_().(D0_DC1).Cf2
						_ = _438___mcc_h1
						var _439___mcc_h2 _dafny.CodePoint = _source5.Get_().(D0_DC1).Cf3
						_ = _439___mcc_h2
						var _440___mcc_h3 _dafny.Array = _source5.Get_().(D0_DC1).Cf4
						_ = _440___mcc_h3
						var _441_cf4 _dafny.Array = _440___mcc_h3
						_ = _441_cf4
						var _442_cf3 _dafny.CodePoint = _439___mcc_h2
						_ = _442_cf3
						var _443_cf2 bool = _438___mcc_h1
						_ = _443_cf2
						var _444_cf1 bool = _437___mcc_h0
						_ = _444_cf1
						var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_423_v63), 0))
						_ = _index71
						(_423_v63).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cnff")).Cardinality()), (_index71).Int())
						var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_423_v63), 0))
						_ = _index72
						var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _index73
						var _rhs76 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(909))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg29 _dafny.Int) interface{} {
								return coer29(arg29)
							}
						}(func(_445_i15 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('t')
						})), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("hmi"), (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hmi")).Cardinality()))).Uint32(), _429_v67))).Cardinality())
						_ = _rhs76
						var _rhs77 bool = (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)
						_ = _rhs77
						var _rhs78 bool = (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)
						_ = _rhs78
						var _lhs63 _dafny.Array = _423_v63
						_ = _lhs63
						var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_423_v63), 0))
						_ = _lhs64
						var _lhs65 _dafny.Array = _431_v69
						_ = _lhs65
						var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _lhs66
						var _lhs67 *GlobalState = globalState
						_ = _lhs67
						(_lhs63).ArraySet1(_rhs76, (_lhs64).Int())
						(_lhs65).ArraySet1(_rhs77, (_lhs66).Int())
						_lhs67.F2 = _rhs78
						var _446_v72 _dafny.Sequence
						_ = _446_v72
						_446_v72 = _dafny.SeqOf(((_dafny.Zero).Minus((func() _dafny.Int {
							if (_380_v28).Contains(p0) {
								return (_380_v28).Get(p0).(_dafny.Int)
							}
							return p0
						})())).Cmp((_423_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_423_v63), 0))).Int()).(_dafny.Int)) <= 0, !(_443_cf2))
						var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _index74
						(_431_v69).ArraySet1((_446_v72).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((p0).Minus((_this).F3())), _dafny.IntOfUint32((_446_v72).Cardinality()))).Uint32()).(bool), (_index74).Int())
						var _447_v73 D6
						_ = _447_v73
						_447_v73 = Companion_D6_.Create_DC17_(Companion_Default___.Fm1(_429_v67, _443_cf2, globalState))
						var _448_v74 D5
						_ = _448_v74
						_448_v74 = Companion_D5_.Create_DC14_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_443_cf2, _dafny.MultiSetFromSeq((_447_v73).Dtor_cf23())))
						_448_v74 = _448_v74
						var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _index75
						(_431_v69).ArraySet1(false, (_index75).Int())
					} else if _source5.Is_DC2() {
						var _449___mcc_h4 _dafny.MultiSet = _source5.Get_().(D0_DC2).Cf5
						_ = _449___mcc_h4
						var _450___mcc_h5 bool = _source5.Get_().(D0_DC2).Cf6
						_ = _450___mcc_h5
						var _451_cf6 bool = _450___mcc_h5
						_ = _451_cf6
						var _452_cf5 _dafny.MultiSet = _449___mcc_h4
						_ = _452_cf5
						var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _index76
						(_431_v69).ArraySet1(!((_dafny.IntOfUint32((_379_v27).Cardinality())).Cmp((_this).F3()) < 0), (_index76).Int())
						(globalState).F2 = !((_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)) || (_451_cf6)
						_379_v27 = _379_v27
						var _453_v75 _dafny.Sequence
						_ = _453_v75
						_453_v75 = _dafny.SeqOf((func() bool {
							if (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool) {
								return _this.F4()
							}
							return (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)
						})())
						var _454_v76 _dafny.Sequence
						_ = _454_v76
						_454_v76 = _dafny.SeqOf(_453_v75, _453_v75, _dafny.SeqOf(_this.F4()), _453_v75)
						var _455_v77 _dafny.MultiSet
						_ = _455_v77
						_455_v77 = _dafny.MultiSetOf((_this).F3(), _dafny.IntOfUint32((_453_v75).Cardinality()))
						_453_v75 = (_454_v76).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
							if (_455_v77).Contains((_this).F7()) {
								return (_455_v77).Multiplicity((_this).F7())
							}
							return p0
						})(), (_this).F3())), _dafny.IntOfUint32((_454_v76).Cardinality()))).Uint32()).(_dafny.Sequence)
					} else if _source5.Is_DC3() {
						var _456_v78 _dafny.Array
						_ = _456_v78
						var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
						_ = _nw83
						_456_v78 = _nw83
						var _457_v79 _dafny.Array
						_ = _457_v79
						var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
						_ = _nw84
						_457_v79 = _nw84
						var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_456_v78), 0))
						_ = _index77
						(_456_v78).ArraySet1(_457_v79, (_index77).Int())
						var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_456_v78), 0))
						_ = _index78
						var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _index79
						var _rhs79 _dafny.Array = _457_v79
						_ = _rhs79
						var _rhs80 _dafny.CodePoint = _429_v67
						_ = _rhs80
						var _rhs81 bool = ((_421_v62).IsDisjointFrom(_421_v62)) && (_this.F4())
						_ = _rhs81
						var _lhs68 _dafny.Array = _456_v78
						_ = _lhs68
						var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_456_v78), 0))
						_ = _lhs69
						var _lhs70 _dafny.Array = _431_v69
						_ = _lhs70
						var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _lhs71
						(_lhs68).ArraySet1(_rhs79, (_lhs69).Int())
						_429_v67 = _rhs80
						(_lhs70).ArraySet1(_rhs81, (_lhs71).Int())
						var _458_v80 _dafny.Map
						_ = _458_v80
						_458_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_379_v27, _431_v69)
						_458_v80 = (_458_v80).Update(_dafny.Companion_Sequence_.Concatenate(_379_v27, _379_v27), _431_v69)
						var _459_v81 _dafny.Sequence
						_ = _459_v81
						_459_v81 = _dafny.SeqOf(_dafny.IntOfInt64(560))
						var _460_v82 _dafny.Map
						_ = _460_v82
						_460_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_459_v81).Cardinality()), !((_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)) || ((_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)))
						_460_v82 = (_460_v82).Update((_this).F3(), (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool))
						var _461_v83 _dafny.MultiSet
						_ = _461_v83
						_461_v83 = _dafny.MultiSetOf((_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool), (p0).Cmp(p0) != 0, (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool))
						_461_v83 = _461_v83
					} else if _source5.Is_DC0() {
						var _462___mcc_h6 bool = _source5.Get_().(D0_DC0).Cf0
						_ = _462___mcc_h6
						var _463_cf0 bool = _462___mcc_h6
						_ = _463_cf0
						_379_v27 = _dafny.Companion_Sequence_.Concatenate(_379_v27, _dafny.UnicodeSeqOfUtf8Bytes("uclqne"))
						var _464_v84 _dafny.Array
						_ = _464_v84
						_464_v84 = _431_v69
						var _465_v85 *C0
						_ = _465_v85
						var _nw85 *C0 = New_C0_()
						_ = _nw85
						_nw85.Ctor__(Companion_Default___.Fm2(p0, globalState), (_464_v84))
						_465_v85 = _nw85
						var _466_v86 _dafny.Map
						_ = _466_v86
						_466_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v85, (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool))
						_466_v86 = (_466_v86).Update(_465_v85, false)
						var _467_v87 _dafny.MultiSet
						_ = _467_v87
						_467_v87 = _dafny.MultiSetOf(_429_v67)
						var _468_v88 D0
						_ = _468_v88
						_468_v88 = Companion_D0_.Create_DC2_(_467_v87, (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool))
						(_465_v85).F10 = ((func() D0 {
							if _463_cf0 {
								return _468_v88
							}
							return _468_v88
						})()).Dtor_cf6()
						var _469_v89 _dafny.MultiSet
						_ = _469_v89
						_469_v89 = _dafny.MultiSetOf(false, Companion_Default___.Fm2((_this).F7(), globalState), _465_v85.F10)
						var _470_v90 _dafny.Sequence
						_ = _470_v90
						_470_v90 = _dafny.SeqOf(_469_v89, _469_v89, (_469_v89).Union(_469_v89), ((_469_v89).Update(_this.F4(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F7())))).Difference(_469_v89), (_dafny.MultiSetOf(_465_v85.F10)).Difference(_469_v89))
						var _471_v91 _dafny.MultiSet
						_ = _471_v91
						_471_v91 = _dafny.MultiSetOf(p0)
						_470_v90 = _dafny.Companion_Sequence_.Update(_470_v90, (Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_471_v91).Contains(p0) {
								return (_471_v91).Multiplicity(p0)
							}
							return (_this).F3()
						})(), _dafny.IntOfUint32((_470_v90).Cardinality()))).Uint32(), _469_v89)
					} else {
						var _472___mcc_h7 D0 = _source5.Get_().(D0_DC4).Cf7
						_ = _472___mcc_h7
						var _473_cf7 D0 = _472___mcc_h7
						_ = _473_cf7
						var _474_v92 _dafny.Sequence
						_ = _474_v92
						_474_v92 = _dafny.SeqOf(!(false))
						(globalState).F2 = (_this.F4()) || (!_dafny.Companion_Sequence_.Contains(_474_v92, (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)))
						var _475_v93 _dafny.Sequence
						_ = _475_v93
						_475_v93 = _dafny.SeqOf((_dafny.SetOf((_this).F7(), (_this).F7())).Cardinality())
						var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _index80
						(_431_v69).ArraySet1((_dafny.IntOfInt64(563)).Cmp((_475_v93).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_475_v93).Cardinality()))).Uint32()).(_dafny.Int)) < 0, (_index80).Int())
						(globalState).F2 = (_431_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))).Int()).(bool)
						var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_431_v69), 0))
						_ = _index81
						(_431_v69).ArraySet1(_this.F4(), (_index81).Int())
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r0 = (_this).F7()
		return r0
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f4  bool
	_f3  _dafny.Int
	_f14 _dafny.Int
	_f15 _dafny.Array
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f4 = false
	_this._f3 = _dafny.Zero
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F4() bool {
	return _this._f4
}
func (_this *C2) F4_set_(value bool) {
	_this._f4 = value
}
func (_this *C2) F3() _dafny.Int {
	return _this._f3
}
func (_this *C2) Ctor__(f14 _dafny.Int, f15 _dafny.Array, f3 _dafny.Int, f4 bool) {
	{
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C2) M6(p0 _dafny.Array, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) {
	{
		if _this.F4() {
			(_this).F4_set_(_this.F4())
			(globalState).F2 = (p2).Cmp(p1) >= 0
			if true {
				var _476_v0 _dafny.Array
				_ = _476_v0
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_14
				var _nw86 _dafny.Array
				_ = _nw86
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw86 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) bool = func(_477_i0 _dafny.Int) bool {
						return (Companion_D0_.Create_DC0_(false)).Dtor_cf0()
					}
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
				_476_v0 = _nw86
				var _478_v1 *C0
				_ = _478_v1
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__(_this.F4(), _476_v0)
				_478_v1 = _nw87
				var _479_v2 *C0
				_ = _479_v2
				var _nw88 *C0 = New_C0_()
				_ = _nw88
				_nw88.Ctor__(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajxdmcbun")).Cardinality()))).Cmp((_this).F3()) == 0, _476_v0)
				_479_v2 = _nw88
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index82
				((_this).F15()).ArraySet1(p1, (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index83
				((_this).F15()).ArraySet1((_this).F3(), (_index83).Int())
				var _480_v3 _dafny.Sequence
				_ = _480_v3
				_480_v3 = _dafny.UnicodeSeqOfUtf8Bytes("di")
				var _481_v4 _dafny.Map
				_ = _481_v4
				_481_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.Companion_Sequence_.Concatenate(_480_v3, _480_v3))
				_481_v4 = (_481_v4).Update(_this.F4(), _dafny.UnicodeSeqOfUtf8Bytes("oeaya"))
				var _482_v5 _dafny.Sequence
				_ = _482_v5
				_482_v5 = _dafny.SeqOf(_this.F4(), _478_v1.F10, _479_v2.F10)
				var _483_v6 _dafny.Map
				_ = _483_v6
				_483_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_482_v5)).Cardinality(), (p2).Plus((_this).F14()))
				var _484_v7 _dafny.MultiSet
				_ = _484_v7
				_484_v7 = _dafny.MultiSetOf((_this).F14())
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index84
				var _rhs82 _dafny.Int = (func() _dafny.Int {
					if (_483_v6).Contains((func() _dafny.Int {
						if _478_v1.F10 {
							return _dafny.IntOfInt64(435)
						}
						return (_dafny.Zero).Minus((_this).F14())
					})()) {
						return (_483_v6).Get((func() _dafny.Int {
							if _478_v1.F10 {
								return _dafny.IntOfInt64(435)
							}
							return (_dafny.Zero).Minus((_this).F14())
						})()).(_dafny.Int)
					}
					return p1
				})()
				_ = _rhs82
				var _rhs83 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F14(), ((_dafny.Zero).Minus(p2)).Times(p2))
				_ = _rhs83
				var _rhs84 bool = ((_484_v7).Intersection(_484_v7)).IsProperSubsetOf(_484_v7)
				_ = _rhs84
				var _lhs72 _dafny.Array = (_this).F15()
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _lhs73
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				var _lhs75 *C0 = _479_v2
				_ = _lhs75
				(_lhs72).ArraySet1(_rhs82, (_lhs73).Int())
				_lhs74.F1 = _rhs83
				_lhs75.F10 = _rhs84
			} else {
				var _485_v8 _dafny.Sequence
				_ = _485_v8
				_485_v8 = _dafny.SeqOf(_dafny.IntOfInt64(-723), (_this).F14(), (_this).F3(), p2, (_this).F3())
				(globalState).F1 = (_485_v8).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_485_v8).Cardinality()))).Uint32()).(_dafny.Int)
				(globalState).F1 = p1
				var _486_v9 _dafny.Array
				_ = _486_v9
				var _nw89 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw89
				_486_v9 = _nw89
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_486_v9), 0))
				_ = _index85
				(_486_v9).ArraySet1(_this.F4(), (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_486_v9), 0))
				_ = _index86
				(_486_v9).ArraySet1(_this.F4(), (_index86).Int())
				var _487_v10 _dafny.Map
				_ = _487_v10
				_487_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_486_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_486_v9), 0))).Int()).(bool), (_this).F3())
				(globalState).F1 = ((func() _dafny.Map {
					if (_486_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_486_v9), 0))).Int()).(bool) {
						return _487_v10
					}
					return _487_v10
				})()).Cardinality()
				var _488_v11 *C0
				_ = _488_v11
				var _nw90 *C0 = New_C0_()
				_ = _nw90
				_nw90.Ctor__(((_this).F3()).Cmp(p1) > 0, _486_v9)
				_488_v11 = _nw90
			}
			var _489_v12 _dafny.Array
			_ = _489_v12
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
			_ = _nw91
			_489_v12 = _nw91
			var _490_v13 _dafny.Array
			_ = _490_v13
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_15
			var _nw92 _dafny.Array
			_ = _nw92
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw92 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) bool = func(_491_i1 _dafny.Int) bool {
					return _this.F4()
				}
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw92 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw92).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw92).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_490_v13 = _nw92
			var _492_v14 _dafny.MultiSet
			_ = _492_v14
			_492_v14 = _dafny.MultiSetOf(_490_v13, _490_v13)
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_489_v12), 0))
			_ = _index87
			(_489_v12).ArraySet1(_492_v14, (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_489_v12), 0))
			_ = _index88
			(_489_v12).ArraySet1(_dafny.MultiSetOf(_490_v13), (_index88).Int())
			(globalState).F2 = false
		} else {
			var _493_v15 _dafny.Array
			_ = _493_v15
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
			_ = _nw93
			_493_v15 = _nw93
			var _494_v16 _dafny.MultiSet
			_ = _494_v16
			_494_v16 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F14()))
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_493_v15), 0))
			_ = _index89
			(_493_v15).ArraySet1(_494_v16, (_index89).Int())
			var _495_v17 _dafny.CodePoint
			_ = _495_v17
			_495_v17 = _dafny.CodePoint('f')
			var _496_v18 _dafny.Map
			_ = _496_v18
			_496_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _495_v17)
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_493_v15), 0))
			_ = _index90
			(_493_v15).ArraySet1((_494_v16).Union(_dafny.MultiSetOf((_496_v18).Cardinality())), (_index90).Int())
			var _rhs85 _dafny.Int = (_this).F14()
			_ = _rhs85
			var _rhs86 bool = _this.F4()
			_ = _rhs86
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			_lhs76.F1 = _rhs85
			_lhs77.F2 = _rhs86
			var _497_v19 _dafny.Array
			_ = _497_v19
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_16
			var _nw94 _dafny.Array
			_ = _nw94
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw94 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Map = func(_498_i2 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _dafny.MultiSetOf(_this.F4(), _this.F4(), _this.F4()))
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw94 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw94).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw94).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_497_v19 = _nw94
			var _499_v20 _dafny.MultiSet
			_ = _499_v20
			_499_v20 = _dafny.MultiSetOf(_this.F4(), false)
			var _500_v21 _dafny.Map
			_ = _500_v21
			_500_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_this).F14(), globalState), _499_v20)
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_497_v19), 0))
			_ = _index91
			(_497_v19).ArraySet1(_500_v21, (_index91).Int())
			var _501_v22 D5
			_ = _501_v22
			_501_v22 = Companion_D5_.Create_DC14_(_500_v21)
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_497_v19), 0))
			_ = _index92
			(_497_v19).ArraySet1(((func() D5 {
				if _this.F4() {
					return _501_v22
				}
				return _501_v22
			})()).Dtor_cf20(), (_index92).Int())
			(globalState).F2 = !(_this.F4())
			var _502_v23 _dafny.Sequence
			_ = _502_v23
			_502_v23 = _dafny.UnicodeSeqOfUtf8Bytes("arajbrv")
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_502_v23, _502_v23) {
				var _503_v24 _dafny.Array
				_ = _503_v24
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_17
				var _nw95 _dafny.Array
				_ = _nw95
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw95 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) bool = func(_504_i3 _dafny.Int) bool {
						return !(_this.F4())
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw95 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw95).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw95).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_503_v24 = _nw95
				var _505_v25 _dafny.Array
				_ = _505_v25
				var _nwElement0_14 _dafny.Array = (func() _dafny.Array {
					if _this.F4() {
						return _503_v24
					}
					return _503_v24
				})()
				_ = _nwElement0_14
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(25))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_14, 0)
				(_nw96).ArraySet1(_503_v24, 1)
				(_nw96).ArraySet1(_503_v24, 2)
				(_nw96).ArraySet1(_503_v24, 3)
				(_nw96).ArraySet1(_503_v24, 4)
				(_nw96).ArraySet1(_503_v24, 5)
				(_nw96).ArraySet1(_503_v24, 6)
				(_nw96).ArraySet1(_503_v24, 7)
				(_nw96).ArraySet1(_503_v24, 8)
				(_nw96).ArraySet1(_503_v24, 9)
				(_nw96).ArraySet1(_503_v24, 10)
				(_nw96).ArraySet1(_503_v24, 11)
				(_nw96).ArraySet1(_503_v24, 12)
				(_nw96).ArraySet1(_503_v24, 13)
				(_nw96).ArraySet1((func() _dafny.Array {
					if !(_this.F4()) {
						return _503_v24
					}
					return _503_v24
				})(), 14)
				(_nw96).ArraySet1(_503_v24, 15)
				(_nw96).ArraySet1(_503_v24, 16)
				(_nw96).ArraySet1(_503_v24, 17)
				(_nw96).ArraySet1(_503_v24, 18)
				(_nw96).ArraySet1(_503_v24, 19)
				(_nw96).ArraySet1(_503_v24, 20)
				(_nw96).ArraySet1(_503_v24, 21)
				(_nw96).ArraySet1(_503_v24, 22)
				(_nw96).ArraySet1(_503_v24, 23)
				(_nw96).ArraySet1(_503_v24, 24)
				_505_v25 = _nw96
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_505_v25), 0))
				_ = _index93
				(_505_v25).ArraySet1(_503_v24, (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_505_v25), 0))
				_ = _index94
				var _rhs87 _dafny.Array = _503_v24
				_ = _rhs87
				var _rhs88 bool = Companion_Default___.Fm2(((_this).F14()).Plus(p1), globalState)
				_ = _rhs88
				var _lhs78 _dafny.Array = _505_v25
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_505_v25), 0))
				_ = _lhs79
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				(_lhs78).ArraySet1(_rhs87, (_lhs79).Int())
				_lhs80.F2 = _rhs88
				(globalState).F1 = (p2).Times(p2)
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_493_v15), 0))
				_ = _index95
				var _rhs89 _dafny.MultiSet = (_493_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_493_v15), 0))).Int()).(_dafny.MultiSet)
				_ = _rhs89
				var _rhs90 bool = _this.F4()
				_ = _rhs90
				var _lhs81 _dafny.Array = _493_v15
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_493_v15), 0))
				_ = _lhs82
				var _lhs83 *GlobalState = globalState
				_ = _lhs83
				(_lhs81).ArraySet1(_rhs89, (_lhs82).Int())
				_lhs83.F2 = _rhs90
				_502_v23 = _502_v23
				var _506_v26 _dafny.Set
				_ = _506_v26
				_506_v26 = _dafny.SetOf(((_493_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_493_v15), 0))).Int()).(_dafny.MultiSet)).IsProperSubsetOf(_494_v16), _this.F4(), _this.F4(), _this.F4())
				var _507_v27 _dafny.Sequence
				_ = _507_v27
				_507_v27 = _dafny.SeqOf(_506_v26)
				var _508_v28 _dafny.Sequence
				_ = _508_v28
				_508_v28 = _dafny.SeqOf(p2, p2)
				_506_v26 = (_507_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_508_v28).Cardinality()), _dafny.IntOfUint32((_507_v27).Cardinality()))).Uint32()).(_dafny.Set)
			} else {
				var _509_v29 _dafny.Set
				_ = _509_v29
				_509_v29 = _dafny.SetOf((_this).F15(), (_this).F15(), (_this).F15(), (_this).F15(), (_this).F15())
				var _510_v30 _dafny.Map
				_ = _510_v30
				_510_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _509_v29)
				_510_v30 = (_510_v30).Update(false, _509_v29)
				var _511_v31 _dafny.Map
				_ = _511_v31
				_511_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(214), (p1).Cmp(p2) == 0)
				_511_v31 = (_511_v31).Update(p1, _this.F4())
				(globalState).F1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_this).F14(), (_this).F14()), (_this).F14())
				var _512_v32 _dafny.Array
				_ = _512_v32
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_18
				var _nw97 _dafny.Array
				_ = _nw97
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw97 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) D5 = (func(_513_v21 _dafny.Map) func(_dafny.Int) D5 {
						return func(_514_i4 _dafny.Int) D5 {
							return Companion_D5_.Create_DC16_(Companion_D5_.Create_DC16_(Companion_D5_.Create_DC16_(Companion_D5_.Create_DC16_(Companion_D5_.Create_DC14_((_513_v21).Update(_this.F4(), _dafny.MultiSetOf(_this.F4(), _this.F4())))))))
						}
					})(_500_v21)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw97 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw97).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw97).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_512_v32 = _nw97
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index96
				((_this).F15()).ArraySet1((_499_v20).Cardinality(), (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index97
				var _rhs91 _dafny.Array = _512_v32
				_ = _rhs91
				var _rhs92 bool = !(_this.F4()) || (_this.F4())
				_ = _rhs92
				var _rhs93 bool = _this.F4()
				_ = _rhs93
				var _rhs94 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F14(), p1)
				_ = _rhs94
				var _rhs95 _dafny.Int = ((_dafny.Zero).Minus((_this).F14())).Plus(((_this).F14()).Minus(p1))
				_ = _rhs95
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 _dafny.Array = (_this).F15()
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _lhs88
				_512_v32 = _rhs91
				_lhs84.F2 = _rhs92
				_lhs85.F2 = _rhs93
				_lhs86.F1 = _rhs94
				(_lhs87).ArraySet1(_rhs95, (_lhs88).Int())
				var _515_v33 _dafny.Array
				_ = _515_v33
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_19
				var _nw98 _dafny.Array
				_ = _nw98
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw98 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) bool = func(_516_i5 _dafny.Int) bool {
						return (_this.F4()) == (_this.F4())
					}
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw98 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw98).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw98).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_515_v33 = _nw98
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw99
				_515_v33 = _nw99
			}
		}
		var _517_i6 _dafny.Int
		_ = _517_i6
		_517_i6 = _dafny.Zero
		{
			for Companion_Default___.Fm2((_dafny.Zero).Minus((_this).F14()), globalState) {
				{
					if (_517_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_517_i6 = (_517_i6).Plus(_dafny.One)
					var _518_v34 _dafny.CodePoint
					_ = _518_v34
					_518_v34 = _dafny.CodePoint('x')
					var _519_v35 _dafny.MultiSet
					_ = _519_v35
					_519_v35 = _dafny.MultiSetOf(_518_v34, _518_v34)
					var _520_v36 D0
					_ = _520_v36
					_520_v36 = Companion_D0_.Create_DC2_(_519_v35, _this.F4())
					var _521_v37 D0
					_ = _521_v37
					_521_v37 = Companion_D0_.Create_DC4_(_520_v36)
					var _rhs96 _dafny.Int = p2
					_ = _rhs96
					var _rhs97 bool = (Companion_Default___.Fm0(globalState)).Cmp(_dafny.IntOfInt64(776)) == 0
					_ = _rhs97
					var _rhs98 bool = _this.F4()
					_ = _rhs98
					var _rhs99 D0 = (func() D0 {
						if _this.F4() {
							return _521_v37
						}
						return _521_v37
					})()
					_ = _rhs99
					var _lhs89 *GlobalState = globalState
					_ = _lhs89
					var _lhs90 *GlobalState = globalState
					_ = _lhs90
					var _lhs91 *GlobalState = globalState
					_ = _lhs91
					_lhs89.F1 = _rhs96
					_lhs90.F2 = _rhs97
					_lhs91.F2 = _rhs98
					_521_v37 = _rhs99
					var _522_v38 _dafny.Array
					_ = _522_v38
					var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
					_ = _nw100
					_522_v38 = _nw100
					var _523_v39 _dafny.Array
					_ = _523_v39
					var _nwElement0_15 bool = _this.F4()
					_ = _nwElement0_15
					var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(18))
					_ = _nw101
					(_nw101).ArraySet1(_nwElement0_15, 0)
					(_nw101).ArraySet1(_this.F4(), 1)
					(_nw101).ArraySet1(_this.F4(), 2)
					(_nw101).ArraySet1(_this.F4(), 3)
					(_nw101).ArraySet1(_this.F4(), 4)
					(_nw101).ArraySet1(_this.F4(), 5)
					(_nw101).ArraySet1(_this.F4(), 6)
					(_nw101).ArraySet1(_this.F4(), 7)
					(_nw101).ArraySet1(_this.F4(), 8)
					(_nw101).ArraySet1(_this.F4(), 9)
					(_nw101).ArraySet1(_this.F4(), 10)
					(_nw101).ArraySet1(_this.F4(), 11)
					(_nw101).ArraySet1(!(_this.F4()), 12)
					(_nw101).ArraySet1(_this.F4(), 13)
					(_nw101).ArraySet1(_this.F4(), 14)
					(_nw101).ArraySet1(_this.F4(), 15)
					(_nw101).ArraySet1(_this.F4(), 16)
					(_nw101).ArraySet1(Companion_Default___.Fm2(p1, globalState), 17)
					_523_v39 = _nw101
					var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_522_v38), 0))
					_ = _index98
					(_522_v38).ArraySet1(_523_v39, (_index98).Int())
					var _524_v40 _dafny.Sequence
					_ = _524_v40
					_524_v40 = _dafny.SeqOf((func() bool {
						if _this.F4() {
							return _this.F4()
						}
						return _this.F4()
					})(), false, _this.F4(), _this.F4())
					var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))
					_ = _index99
					(_523_v39).ArraySet1(_this.F4(), (_index99).Int())
					var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_522_v38), 0))
					_ = _index100
					var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))
					_ = _index101
					var _rhs100 _dafny.Array = _523_v39
					_ = _rhs100
					var _rhs101 _dafny.Sequence = _524_v40
					_ = _rhs101
					var _rhs102 bool = (p1).Cmp((_this).F3()) == 0
					_ = _rhs102
					var _lhs92 _dafny.Array = _522_v38
					_ = _lhs92
					var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_522_v38), 0))
					_ = _lhs93
					var _lhs94 _dafny.Array = _523_v39
					_ = _lhs94
					var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))
					_ = _lhs95
					(_lhs92).ArraySet1(_rhs100, (_lhs93).Int())
					_524_v40 = _rhs101
					(_lhs94).ArraySet1(_rhs102, (_lhs95).Int())
					var _525_v41 _dafny.Map
					_ = _525_v41
					_525_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_523_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))).Int()).(bool), (_this).F3())
					var _526_v42 _dafny.MultiSet
					_ = _526_v42
					_526_v42 = _dafny.MultiSetOf(_dafny.IntOfInt64(850))
					var _527_v43 _dafny.MultiSet
					_ = _527_v43
					_527_v43 = _dafny.MultiSetOf(false, (_523_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))).Int()).(bool))
					var _528_v44 _dafny.Map
					_ = _528_v44
					_528_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_519_v35).Cardinality(), _this.F4())
					var _529_v45 _dafny.Map
					_ = _529_v45
					_529_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_528_v44).Cardinality())
					var _530_v46 _dafny.Array
					_ = _530_v46
					var _nwElement0_16 _dafny.Int = Companion_Default___.SafeModuloInt(p1, (_this).F3())
					_ = _nwElement0_16
					var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(25))
					_ = _nw102
					(_nw102).ArraySet1(_nwElement0_16, 0)
					(_nw102).ArraySet1(((_525_v41).Cardinality()).Plus(_dafny.IntOfInt64(688)), 1)
					(_nw102).ArraySet1(p1, 2)
					(_nw102).ArraySet1((p1).Times((_this).F3()), 3)
					(_nw102).ArraySet1(((_this).F14()).Times((_this).F14()), 4)
					(_nw102).ArraySet1(Companion_Default___.Fm0(globalState), 5)
					(_nw102).ArraySet1(Companion_Default___.SafeModuloInt((_this).F3(), (_this).F3()), 6)
					(_nw102).ArraySet1((p2).Minus(_dafny.IntOfInt64(-388)), 7)
					(_nw102).ArraySet1((_dafny.Zero).Minus((_this).F14()), 8)
					(_nw102).ArraySet1(((Companion_D1_.Create_DC5_(Companion_Default___.Fm0(globalState))).Dtor_cf8()).Minus(p1), 9)
					(_nw102).ArraySet1((func() _dafny.Int {
						if (_526_v42).Contains(p2) {
							return (_526_v42).Multiplicity(p2)
						}
						return (_this).F14()
					})(), 10)
					(_nw102).ArraySet1((_this).F14(), 11)
					(_nw102).ArraySet1((_this).F14(), 12)
					(_nw102).ArraySet1(p2, 13)
					(_nw102).ArraySet1(p2, 14)
					(_nw102).ArraySet1(p2, 15)
					(_nw102).ArraySet1((_this).F3(), 16)
					(_nw102).ArraySet1(p2, 17)
					(_nw102).ArraySet1(p1, 18)
					(_nw102).ArraySet1(p1, 19)
					(_nw102).ArraySet1((_this).F3(), 20)
					(_nw102).ArraySet1(p1, 21)
					(_nw102).ArraySet1((func() _dafny.Int {
						if (_527_v43).Contains((_523_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))).Int()).(bool)) {
							return (_527_v43).Multiplicity((_523_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))).Int()).(bool))
						}
						return (_529_v45).Cardinality()
					})(), 22)
					(_nw102).ArraySet1((_this).F3(), 23)
					(_nw102).ArraySet1(_dafny.IntOfInt64(-440), 24)
					_530_v46 = _nw102
					var _531_v47 _dafny.Array
					_ = _531_v47
					var _nw103 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
					_ = _nw103
					_531_v47 = _nw103
					var _532_v48 D8
					_ = _532_v48
					_532_v48 = Companion_D8_.Create_DC22_(_528_v44)
					var _533_v49 T1
					_ = _533_v49
					var _nw104 *C1 = New_C1_()
					_ = _nw104
					_nw104.Ctor__((_this).F3(), ((_532_v48).Dtor_cf31()).Cardinality(), false)
					_533_v49 = _nw104
					var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_531_v47), 0))
					_ = _index102
					(_531_v47).ArraySet1(_533_v49, (_index102).Int())
					var _534_v50 _dafny.Map
					_ = _534_v50
					_534_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_529_v45).Update(p2, (_this).F3()), _530_v46)
					var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_531_v47), 0))
					_ = _index103
					var _rhs103 bool = (_523_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_523_v39), 0))).Int()).(bool)
					_ = _rhs103
					var _rhs104 _dafny.Int = p1
					_ = _rhs104
					var _rhs105 _dafny.Array = (func() _dafny.Array {
						if (_534_v50).Contains((Companion_Default___.Fm23(globalState)).Update((_533_v49).F7(), (_dafny.Zero).Minus((_533_v49).F7()))) {
							return (_534_v50).Get((Companion_Default___.Fm23(globalState)).Update((_533_v49).F7(), (_dafny.Zero).Minus((_533_v49).F7()))).(_dafny.Array)
						}
						return _530_v46
					})()
					_ = _rhs105
					var _rhs106 T1 = _533_v49
					_ = _rhs106
					var _rhs107 _dafny.Int = (_533_v49).F3()
					_ = _rhs107
					var _lhs96 *C2 = _this
					_ = _lhs96
					var _lhs97 *GlobalState = globalState
					_ = _lhs97
					var _lhs98 _dafny.Array = _531_v47
					_ = _lhs98
					var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_531_v47), 0))
					_ = _lhs99
					var _lhs100 *GlobalState = globalState
					_ = _lhs100
					_lhs96.F4_set_(_rhs103)
					_lhs97.F1 = _rhs104
					_530_v46 = _rhs105
					(_lhs98).ArraySet1(_rhs106, (_lhs99).Int())
					_lhs100.F1 = _rhs107
					var _535_v51 _dafny.Array
					_ = _535_v51
					var _nwElement0_17 _dafny.CodePoint = _518_v34
					_ = _nwElement0_17
					var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(17))
					_ = _nw105
					(_nw105).ArraySet1CodePoint(_nwElement0_17, 0)
					(_nw105).ArraySet1CodePoint(_518_v34, 1)
					(_nw105).ArraySet1CodePoint(_518_v34, 2)
					(_nw105).ArraySet1CodePoint(_518_v34, 3)
					(_nw105).ArraySet1CodePoint(_518_v34, 4)
					(_nw105).ArraySet1CodePoint(_518_v34, 5)
					(_nw105).ArraySet1CodePoint(_518_v34, 6)
					(_nw105).ArraySet1CodePoint(_518_v34, 7)
					(_nw105).ArraySet1CodePoint((func() _dafny.CodePoint {
						if _this.F4() {
							return _dafny.CodePoint('k')
						}
						return _518_v34
					})(), 8)
					(_nw105).ArraySet1CodePoint(_518_v34, 9)
					(_nw105).ArraySet1CodePoint(_518_v34, 10)
					(_nw105).ArraySet1CodePoint(_518_v34, 11)
					(_nw105).ArraySet1CodePoint(_518_v34, 12)
					(_nw105).ArraySet1CodePoint(_518_v34, 13)
					(_nw105).ArraySet1CodePoint(_518_v34, 14)
					(_nw105).ArraySet1CodePoint(_518_v34, 15)
					(_nw105).ArraySet1CodePoint(_518_v34, 16)
					_535_v51 = _nw105
					var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_535_v51), 0))
					_ = _index104
					(_535_v51).ArraySet1CodePoint(_518_v34, (_index104).Int())
					var _536_v52 _dafny.Map
					_ = _536_v52
					_536_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p1, p1), _dafny.CodePoint('n'))
					var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_535_v51), 0))
					_ = _index105
					(_535_v51).ArraySet1CodePoint((func() _dafny.CodePoint {
						if (_536_v52).Contains(_dafny.IntOfInt64(693)) {
							return (_536_v52).Get(_dafny.IntOfInt64(693)).(_dafny.CodePoint)
						}
						return _dafny.CodePoint('g')
					})(), (_index105).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _537_v53 _dafny.Array
		_ = _537_v53
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_20
		var _nw106 _dafny.Array
		_ = _nw106
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw106 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) bool = func(_538_i8 _dafny.Int) bool {
				return false
			}
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw106 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw106).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw106).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_537_v53 = _nw106
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_537_v53), 0))); ; {
			_guard_loop_1, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _539_i7 _dafny.Int
			_539_i7 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_539_i7).Sign() != -1) && ((_539_i7).Cmp(_dafny.ArrayLen((_537_v53), 0)) < 0)) {
				(_537_v53).ArraySet1(_this.F4(), (_539_i7).Int())
			}
		}
		var _540_v54 _dafny.Sequence
		_ = _540_v54
		_540_v54 = _dafny.SeqOf(_this.F4(), _this.F4())
		var _541_v55 D7
		_ = _541_v55
		_541_v55 = Companion_D7_.Create_DC21_(_537_v53, _dafny.Companion_Sequence_.IsProperPrefixOf(_540_v54, _dafny.SeqOf(_this.F4(), _this.F4(), _this.F4(), _this.F4())), !(_this.F4()))
		_541_v55 = _541_v55
		var _542_v56 _dafny.Map
		_ = _542_v56
		_542_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), (p1).Plus(p2))
		var _543_v57 _dafny.CodePoint
		_ = _543_v57
		_543_v57 = _dafny.CodePoint('v')
		var _544_v58 _dafny.MultiSet
		_ = _544_v58
		_544_v58 = _dafny.MultiSetOf(_dafny.CodePoint('w'), _543_v57, _dafny.CodePoint('r'))
		_542_v56 = (_542_v56).Update((Companion_D0_.Create_DC2_(_544_v58, _this.F4())).Dtor_cf6(), Companion_Default___.SafeDivisionInt((_this).F3(), (_this).F3()))
		if !((_this.F4()) || (true)) || (_this.F4()) {
			var _545_v59 D8
			_ = _545_v59
			_545_v59 = Companion_D8_.Create_DC23_(_this.F4(), _this.F4(), _this.F4(), _543_v57)
			var _546_v60 _dafny.Sequence
			_ = _546_v60
			_546_v60 = _dafny.UnicodeSeqOfUtf8Bytes("dakcdomv")
			var _pat_let_tv4 = globalState
			_ = _pat_let_tv4
			var _rhs108 D8 = func(_pat_let5_0 D8) D8 {
				return func(_547_dt__update__tmp_h0 D8) D8 {
					return func(_pat_let6_0 _dafny.CodePoint) D8 {
						return func(_548_dt__update_hcf35_h0 _dafny.CodePoint) D8 {
							return func(_pat_let7_0 bool) D8 {
								return func(_549_dt__update_hcf32_h0 bool) D8 {
									return Companion_D8_.Create_DC23_(_549_dt__update_hcf32_h0, (_547_dt__update__tmp_h0).Dtor_cf33(), (_547_dt__update__tmp_h0).Dtor_cf34(), _548_dt__update_hcf35_h0)
								}(_pat_let7_0)
							}(_this.F4())
						}(_pat_let6_0)
					}(Companion_Default___.Fm24(_pat_let_tv4))
				}(_pat_let5_0)
			}((func() D8 {
				if true {
					return _545_v59
				}
				return _545_v59
			})())
			_ = _rhs108
			var _rhs109 bool = _this.F4()
			_ = _rhs109
			var _rhs110 bool = _dafny.Companion_Sequence_.Contains(_546_v60, _543_v57)
			_ = _rhs110
			var _lhs101 *GlobalState = globalState
			_ = _lhs101
			var _lhs102 *GlobalState = globalState
			_ = _lhs102
			_545_v59 = _rhs108
			_lhs101.F2 = _rhs109
			_lhs102.F2 = _rhs110
			_543_v57 = _543_v57
			var _550_v61 _dafny.Map
			_ = _550_v61
			_550_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), (_this).F14())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _index106
			((_this).F15()).ArraySet1((func() _dafny.Int {
				if (_550_v61).Contains((_this).F3()) {
					return (_550_v61).Get((_this).F3()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus((_this).F3())
			})(), (_index106).Int())
			var _551_v62 _dafny.Set
			_ = _551_v62
			_551_v62 = _dafny.SetOf(_this.F4())
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _index107
			((_this).F15()).ArraySet1((p1).Plus(((_551_v62).Union(_551_v62)).Cardinality()), (_index107).Int())
			_540_v54 = _540_v54
			var _552_v63 _dafny.Sequence
			_ = _552_v63
			_552_v63 = _dafny.SeqOf(p1, p2)
			(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_552_v63, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(222))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_553_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_554_i9 _dafny.Int) _dafny.Int {
					return _553_p2
				}
			})(p2))), _552_v63), (Companion_Default___.SafeIndex(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(222))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_555_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_556_i9 _dafny.Int) _dafny.Int {
					return _555_p2
				}
			})(p2))), _552_v63)).Cardinality()))).Uint32(), p1))).Cardinality())
		} else {
			var _557_v64 _dafny.Map
			_ = _557_v64
			_557_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_540_v54, (_this).F3())
			(globalState).F1 = (func() _dafny.Int {
				if (_557_v64).Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_540_v54, _540_v54), (Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_540_v54, _540_v54)).Cardinality()))).Uint32(), true)) {
					return (_557_v64).Get(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_540_v54, _540_v54), (Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_540_v54, _540_v54)).Cardinality()))).Uint32(), true)).(_dafny.Int)
				}
				return p1
			})()
			var _558_v65 D8
			_ = _558_v65
			_558_v65 = Companion_D8_.Create_DC23_(_this.F4(), _this.F4(), _this.F4(), _543_v57)
			var _559_v66 D8
			_ = _559_v66
			_559_v66 = Companion_D8_.Create_DC24_(_558_v65)
			var _560_v67 D8
			_ = _560_v67
			_560_v67 = Companion_D8_.Create_DC24_(_559_v66)
			var _561_v68 D8
			_ = _561_v68
			_561_v68 = Companion_D8_.Create_DC24_(_560_v67)
			var _562_v69 _dafny.Map
			_ = _562_v69
			_562_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _561_v68)
			_562_v69 = ((_562_v69).Merge(_562_v69)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _561_v68))
			var _563_v70 *C1
			_ = _563_v70
			var _nw107 *C1 = New_C1_()
			_ = _nw107
			_nw107.Ctor__((_this).F14(), p2, _this.F4())
			_563_v70 = _nw107
			var _564_v71 _dafny.Sequence
			_ = _564_v71
			_564_v71 = _dafny.UnicodeSeqOfUtf8Bytes("axxyv")
			_564_v71 = _564_v71
			var _565_v72 _dafny.Array
			_ = _565_v72
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_21
			var _nw108 _dafny.Array
			_ = _nw108
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw108 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Sequence = (func(_566_v58 _dafny.MultiSet) func(_dafny.Int) _dafny.Sequence {
					return func(_567_i10 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_566_v58).Cardinality())
					}
				})(_544_v58)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw108 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw108).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw108).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_565_v72 = _nw108
			var _568_v73 _dafny.Set
			_ = _568_v73
			_568_v73 = _dafny.SetOf((_this).F3())
			var _569_v74 _dafny.Map
			_ = _569_v74
			_569_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
			var _rhs111 _dafny.Int = _dafny.IntOfInt64(-860)
			_ = _rhs111
			var _rhs112 bool = !((_this.F4()) && ((_568_v73).Equals(_568_v73)))
			_ = _rhs112
			var _rhs113 bool = Companion_Default___.Fm2(((_569_v74).Merge((_569_v74).Update(_this.F4(), _this.F4()))).Cardinality(), globalState)
			_ = _rhs113
			var _rhs114 bool = Companion_Default___.Fm2(p2, globalState)
			_ = _rhs114
			var _rhs115 _dafny.Array = _565_v72
			_ = _rhs115
			var _lhs103 *GlobalState = globalState
			_ = _lhs103
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			var _lhs105 *GlobalState = globalState
			_ = _lhs105
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			_lhs103.F1 = _rhs111
			_lhs104.F2 = _rhs112
			_lhs105.F2 = _rhs113
			_lhs106.F2 = _rhs114
			_565_v72 = _rhs115
		}
	}
}
func (_this *C2) M7(p0 _dafny.Int, p1 _dafny.Array, globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		r0 = (_this).F14()
		var _570_v0 _dafny.Sequence
		_ = _570_v0
		_570_v0 = _dafny.SeqOf(_dafny.IntOfInt64(-638))
		var _571_v1 _dafny.Sequence
		_ = _571_v1
		_571_v1 = _dafny.SeqOf(_dafny.IntOfUint32((_570_v0).Cardinality()), (p0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F3())).Cardinality()), (_dafny.IntOfInt64(-673)).Plus((_this).F14()))
		_571_v1 = _570_v0
		var _hi5 _dafny.Int = Companion_Default___.Fm0(globalState)
		_ = _hi5
		for _572_i0 := (_this).F14(); _572_i0.Cmp(_hi5) < 0; _572_i0 = _572_i0.Plus(_dafny.One) {
			var _573_v2 _dafny.CodePoint
			_ = _573_v2
			_573_v2 = _dafny.CodePoint('i')
			var _574_v3 _dafny.Set
			_ = _574_v3
			_574_v3 = _dafny.SetOf(_573_v2)
			var _575_v4 _dafny.Array
			_ = _575_v4
			var _nwElement0_18 _dafny.Set = _574_v3
			_ = _nwElement0_18
			var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(6))
			_ = _nw109
			(_nw109).ArraySet1(_nwElement0_18, 0)
			(_nw109).ArraySet1(_574_v3, 1)
			(_nw109).ArraySet1((_574_v3).Intersection(_574_v3), 2)
			(_nw109).ArraySet1(_574_v3, 3)
			(_nw109).ArraySet1(_dafny.SetOf(Companion_Default___.Fm24(globalState), _573_v2, _573_v2), 4)
			(_nw109).ArraySet1((func() _dafny.Set {
				if _this.F4() {
					return _574_v3
				}
				return _574_v3
			})(), 5)
			_575_v4 = _nw109
			_575_v4 = _575_v4
			var _source6 D0 = Companion_Default___.Fm25((func() bool {
				if _this.F4() {
					return Companion_Default___.Fm2(_dafny.IntOfInt64(-606), globalState)
				}
				return _this.F4()
			})(), globalState)
			_ = _source6
			if _source6.Is_DC1() {
				var _576___mcc_h0 bool = _source6.Get_().(D0_DC1).Cf1
				_ = _576___mcc_h0
				var _577___mcc_h1 bool = _source6.Get_().(D0_DC1).Cf2
				_ = _577___mcc_h1
				var _578___mcc_h2 _dafny.CodePoint = _source6.Get_().(D0_DC1).Cf3
				_ = _578___mcc_h2
				var _579___mcc_h3 _dafny.Array = _source6.Get_().(D0_DC1).Cf4
				_ = _579___mcc_h3
				var _580_cf4 _dafny.Array = _579___mcc_h3
				_ = _580_cf4
				var _581_cf3 _dafny.CodePoint = _578___mcc_h2
				_ = _581_cf3
				var _582_cf2 bool = _577___mcc_h1
				_ = _582_cf2
				var _583_cf1 bool = _576___mcc_h0
				_ = _583_cf1
				var _584_v5 D0
				_ = _584_v5
				_584_v5 = Companion_D0_.Create_DC1_(false, _583_cf1, _581_cf3, _580_cf4)
				(globalState).F1 = _dafny.IntOfUint32((_dafny.SeqOf((_584_v5).Dtor_cf3())).Cardinality())
				var _585_v6 _dafny.Map
				_ = _585_v6
				_585_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F4())
				var _586_v7 _dafny.Set
				_ = _586_v7
				_586_v7 = _dafny.SetOf(_583_cf1)
				var _587_v8 _dafny.Map
				_ = _587_v8
				_587_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_583_cf1, _586_v7)
				var _588_v9 _dafny.Sequence
				_ = _588_v9
				_588_v9 = _dafny.SeqOf(_586_v7, _dafny.SetOf(false))
				var _589_v10 _dafny.Sequence
				_ = _589_v10
				_589_v10 = _dafny.UnicodeSeqOfUtf8Bytes("ltfp")
				var _590_v11 _dafny.MultiSet
				_ = _590_v11
				_590_v11 = _dafny.MultiSetOf(_this.F4(), true)
				var _591_v12 _dafny.Map
				_ = _591_v12
				_591_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_582_cf2, _590_v11)
				var _592_v13 _dafny.Array
				_ = _592_v13
				var _nwElement0_19 bool = _this.F4()
				_ = _nwElement0_19
				var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(14))
				_ = _nw110
				(_nw110).ArraySet1(_nwElement0_19, 0)
				(_nw110).ArraySet1(_this.F4(), 1)
				(_nw110).ArraySet1((func() bool {
					if (_585_v6).Contains(Companion_Default___.Fm2((_this).F3(), globalState)) {
						return (_585_v6).Get(Companion_Default___.Fm2((_this).F3(), globalState)).(bool)
					}
					return _this.F4()
				})(), 2)
				(_nw110).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm20(((func() _dafny.Set {
					if (_587_v8).Contains(Companion_Default___.Fm2(p0, globalState)) {
						return (_587_v8).Get(Companion_Default___.Fm2(p0, globalState)).(_dafny.Set)
					}
					return (_588_v9).Select((Companion_Default___.SafeIndex(_572_i0, _dafny.IntOfUint32((_588_v9).Cardinality()))).Uint32()).(_dafny.Set)
				})()).Cardinality(), globalState), _589_v10), 3)
				(_nw110).ArraySet1((_this.F4()) || (Companion_Default___.Fm2((_591_v12).Cardinality(), globalState)), 4)
				(_nw110).ArraySet1(false, 5)
				(_nw110).ArraySet1(_582_cf2, 6)
				(_nw110).ArraySet1(!(_this.F4()), 7)
				(_nw110).ArraySet1(((_this).F14()).Cmp(_572_i0) < 0, 8)
				(_nw110).ArraySet1((func() bool {
					if _582_cf2 {
						return _583_cf1
					}
					return _582_cf2
				})(), 9)
				(_nw110).ArraySet1((func() bool {
					if (_585_v6).Contains(_582_cf2) {
						return (_585_v6).Get(_582_cf2).(bool)
					}
					return _582_cf2
				})(), 10)
				(_nw110).ArraySet1(!(_this.F4()), 11)
				(_nw110).ArraySet1(_this.F4(), 12)
				(_nw110).ArraySet1((Companion_Default___.Fm0(globalState)).Cmp(p0) != 0, 13)
				_592_v13 = _nw110
				var _593_v14 _dafny.MultiSet
				_ = _593_v14
				_593_v14 = _dafny.MultiSetOf(_590_v11, _590_v11)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_592_v13), 0))
				_ = _index108
				(_592_v13).ArraySet1((_593_v14).Equals(_dafny.MultiSetFromSeq(_dafny.SeqOf(_590_v11))), (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_592_v13), 0))
				_ = _index109
				(_592_v13).ArraySet1((_572_i0).Cmp(_572_i0) < 0, (_index109).Int())
				var _594_v15 *C0
				_ = _594_v15
				var _nw111 *C0 = New_C0_()
				_ = _nw111
				_nw111.Ctor__((_592_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_592_v13), 0))).Int()).(bool), _592_v13)
				_594_v15 = _nw111
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index110
				((_this).F15()).ArraySet1((_this).F3(), (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index111
				var _rhs116 _dafny.Array = (_this).F15()
				_ = _rhs116
				var _rhs117 _dafny.Int = ((_dafny.Zero).Minus((_this).F14())).Minus((_dafny.Zero).Minus((_this).F14()))
				_ = _rhs117
				var _lhs107 _dafny.Array = (_this).F15()
				_ = _lhs107
				var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _lhs108
				_580_cf4 = _rhs116
				(_lhs107).ArraySet1(_rhs117, (_lhs108).Int())
			} else if _source6.Is_DC2() {
				var _595___mcc_h4 _dafny.MultiSet = _source6.Get_().(D0_DC2).Cf5
				_ = _595___mcc_h4
				var _596___mcc_h5 bool = _source6.Get_().(D0_DC2).Cf6
				_ = _596___mcc_h5
				var _597_cf6 bool = _596___mcc_h5
				_ = _597_cf6
				var _598_cf5 _dafny.MultiSet = _595___mcc_h4
				_ = _598_cf5
				var _599_v16 _dafny.Array
				_ = _599_v16
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
				_ = _nw112
				_599_v16 = _nw112
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_599_v16), 0))
				_ = _index112
				(_599_v16).ArraySet1((func() _dafny.Array {
					if _597_cf6 {
						return (_this).F15()
					}
					return (_this).F15()
				})(), (_index112).Int())
				var _600_v17 _dafny.Array
				_ = _600_v17
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(26))
				_ = _nw113
				_600_v17 = _nw113
				var _601_v18 D8
				_ = _601_v18
				_601_v18 = Companion_D8_.Create_DC23_(false, _597_cf6, _this.F4(), _573_v2)
				var _602_v19 D8
				_ = _602_v19
				_602_v19 = Companion_D8_.Create_DC24_(_601_v18)
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_600_v17), 0))
				_ = _index113
				(_600_v17).ArraySet1(_602_v19, (_index113).Int())
				var _603_v20 _dafny.MultiSet
				_ = _603_v20
				_603_v20 = _dafny.MultiSetOf(_572_i0, (_this).F14())
				var _604_v21 _dafny.Sequence
				_ = _604_v21
				_604_v21 = _dafny.SeqOf(_597_cf6, _597_cf6)
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_599_v16), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_600_v17), 0))
				_ = _index115
				var _rhs118 _dafny.Array = (_this).F15()
				_ = _rhs118
				var _rhs119 D8 = Companion_Default___.Fm26(globalState)
				_ = _rhs119
				var _rhs120 _dafny.MultiSet = (_603_v20).Update(_dafny.IntOfUint32((_604_v21).Cardinality()), Companion_Default___.Abs((_this).F3()))
				_ = _rhs120
				var _lhs109 _dafny.Array = _599_v16
				_ = _lhs109
				var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_599_v16), 0))
				_ = _lhs110
				var _lhs111 _dafny.Array = _600_v17
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_600_v17), 0))
				_ = _lhs112
				(_lhs109).ArraySet1(_rhs118, (_lhs110).Int())
				(_lhs111).ArraySet1(_rhs119, (_lhs112).Int())
				_603_v20 = _rhs120
				var _605_v22 _dafny.Map
				_ = _605_v22
				_605_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_572_i0, _this.F4())
				_605_v22 = (_605_v22).Update((_dafny.Zero).Minus((_this).F3()), _this.F4())
				r0 = _572_i0
				r0 = (_this).F3()
			} else if _source6.Is_DC3() {
				var _606_v23 _dafny.Sequence
				_ = _606_v23
				_606_v23 = _dafny.UnicodeSeqOfUtf8Bytes("af")
				var _607_v24 _dafny.Map
				_ = _607_v24
				_607_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_572_i0).Cmp(_dafny.IntOfInt64(-271)) >= 0, _dafny.SeqOf(_606_v23))
				var _608_v25 _dafny.Sequence
				_ = _608_v25
				_608_v25 = _dafny.SeqOf(_this.F4())
				var _609_v26 _dafny.Sequence
				_ = _609_v26
				_609_v26 = _dafny.SeqOf(_606_v23)
				_607_v24 = (_607_v24).Update(!((_this.F4()) == ((_608_v25).Select((Companion_Default___.SafeIndex(_572_i0, _dafny.IntOfUint32((_608_v25).Cardinality()))).Uint32()).(bool))), _609_v26)
				var _610_v27 _dafny.Array
				_ = _610_v27
				var _nwElement0_20 bool = _this.F4()
				_ = _nwElement0_20
				var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(17))
				_ = _nw114
				(_nw114).ArraySet1(_nwElement0_20, 0)
				(_nw114).ArraySet1(true, 1)
				(_nw114).ArraySet1(_this.F4(), 2)
				(_nw114).ArraySet1(_this.F4(), 3)
				(_nw114).ArraySet1(!(_this.F4()), 4)
				(_nw114).ArraySet1(_this.F4(), 5)
				(_nw114).ArraySet1(Companion_Default___.Fm2(_572_i0, globalState), 6)
				(_nw114).ArraySet1(_this.F4(), 7)
				(_nw114).ArraySet1(true, 8)
				(_nw114).ArraySet1(_this.F4(), 9)
				(_nw114).ArraySet1(_this.F4(), 10)
				(_nw114).ArraySet1(false, 11)
				(_nw114).ArraySet1(true, 12)
				(_nw114).ArraySet1(false, 13)
				(_nw114).ArraySet1(_this.F4(), 14)
				(_nw114).ArraySet1(_this.F4(), 15)
				(_nw114).ArraySet1(_this.F4(), 16)
				_610_v27 = _nw114
				var _611_v28 _dafny.Map
				_ = _611_v28
				_611_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _610_v27)
				_611_v28 = (_611_v28).Update(true, _610_v27)
				var _612_v29 *C0
				_ = _612_v29
				var _nw115 *C0 = New_C0_()
				_ = _nw115
				_nw115.Ctor__(_this.F4(), _610_v27)
				_612_v29 = _nw115
				(_this).F4_set_(_612_v29.F10)
			} else if _source6.Is_DC0() {
				var _613___mcc_h6 bool = _source6.Get_().(D0_DC0).Cf0
				_ = _613___mcc_h6
				var _614_cf0 bool = _613___mcc_h6
				_ = _614_cf0
				var _615_v30 _dafny.Sequence
				_ = _615_v30
				_615_v30 = _dafny.UnicodeSeqOfUtf8Bytes("gjdgtgnld")
				var _616_v31 _dafny.Map
				_ = _616_v31
				_616_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v30, _572_i0)
				var _617_v32 _dafny.Map
				_ = _617_v32
				_617_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_cf0, false)
				_616_v31 = (_616_v31).Update(_615_v30, (_dafny.MultiSetOf(_614_cf0, (func() bool {
					if (_617_v32).Contains(_this.F4()) {
						return (_617_v32).Get(_this.F4()).(bool)
					}
					return true
				})())).Cardinality())
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index116
				((_this).F15()).ArraySet1((_this).F14(), (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index117
				((_this).F15()).ArraySet1((_this).F14(), (_index117).Int())
				var _618_v33 _dafny.Array
				_ = _618_v33
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
				_ = _nw116
				_618_v33 = _nw116
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_618_v33), 0))
				_ = _index118
				(_618_v33).ArraySet1(_615_v30, (_index118).Int())
				var _619_v34 _dafny.Map
				_ = _619_v34
				_619_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(_dafny.Int))
				var _620_v36 _dafny.Map
				_ = _620_v36
				_620_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), (func() _dafny.Set {
					var _coll8 = _dafny.NewBuilder()
					_ = _coll8
					for _iter10 := _dafny.Iterate((_619_v34).Keys().Elements()); ; {
						_compr_8, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _621_v35 _dafny.Int
						_621_v35 = interface{}(_compr_8).(_dafny.Int)
						if (_619_v34).Contains(_621_v35) {
							_coll8.Add((_621_v35).Times(_dafny.IntOfInt64(-756)))
						}
					}
					return _coll8.ToSet()
				}()).Cardinality())
				var _622_v37 _dafny.MultiSet
				_ = _622_v37
				_622_v37 = _dafny.MultiSetOf(_620_v36, _620_v36)
				var _623_v38 _dafny.Array
				_ = _623_v38
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw117
				_623_v38 = _nw117
				var _624_v39 *C0
				_ = _624_v39
				var _nw118 *C0 = New_C0_()
				_ = _nw118
				_nw118.Ctor__(_614_cf0, _623_v38)
				_624_v39 = _nw118
				var _625_v40 _dafny.Map
				_ = _625_v40
				_625_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_624_v39, _615_v30)
				var _626_v41 _dafny.Map
				_ = _626_v41
				_626_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_622_v37).Update(_620_v36, Companion_Default___.Abs((_this).F3())), (func() _dafny.Sequence {
					if (_625_v40).Contains(_624_v39) {
						return (_625_v40).Get(_624_v39).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("tpoiv")
				})())
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_618_v33), 0))
				_ = _index119
				(_618_v33).ArraySet1((func() _dafny.Sequence {
					if (_626_v41).Contains(_622_v37) {
						return (_626_v41).Get(_622_v37).(_dafny.Sequence)
					}
					return _615_v30
				})(), (_index119).Int())
				_570_v0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_627_i1 _dafny.Int) _dafny.Int {
					return (_this).F3()
				}))
			} else {
				var _628___mcc_h7 D0 = _source6.Get_().(D0_DC4).Cf7
				_ = _628___mcc_h7
				var _629_cf7 D0 = _628___mcc_h7
				_ = _629_cf7
				r1 = Companion_Default___.Fm2((_dafny.Zero).Minus((_this).F3()), globalState)
				var _630_v42 _dafny.Map
				_ = _630_v42
				_630_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), true)
				_630_v42 = (_630_v42).Update((_this).F3(), !((p0).Cmp((_this).F14()) > 0))
				var _631_v43 _dafny.Sequence
				_ = _631_v43
				_631_v43 = _dafny.UnicodeSeqOfUtf8Bytes("viwcggtyo")
				var _632_v44 _dafny.Set
				_ = _632_v44
				_632_v44 = _dafny.SetOf(p0, _dafny.IntOfUint32((_631_v43).Cardinality()))
				var _633_v45 _dafny.MultiSet
				_ = _633_v45
				_633_v45 = _dafny.MultiSetOf((_this).F14())
				r0 = ((_632_v44).Cardinality()).Plus((func() _dafny.Int {
					if _this.F4() {
						return (_this).F3()
					}
					return (_633_v45).Cardinality()
				})())
				var _634_v46 _dafny.Array
				_ = _634_v46
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_22
				var _nw119 _dafny.Array
				_ = _nw119
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw119 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) bool = func(_635_i2 _dafny.Int) bool {
						return _this.F4()
					}
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw119 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw119).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw119).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_634_v46 = _nw119
				var _636_v47 D7
				_ = _636_v47
				_636_v47 = Companion_D7_.Create_DC21_(_634_v46, _this.F4(), _this.F4())
				var _pat_let_tv5 = _634_v46
				_ = _pat_let_tv5
				var _637_v48 _dafny.Array
				_ = _637_v48
				var _nwElement0_21 D7 = Companion_D7_.Create_DC21_(_634_v46, _this.F4(), _this.F4())
				_ = _nwElement0_21
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(15))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_21, 0)
				(_nw120).ArraySet1(_636_v47, 1)
				(_nw120).ArraySet1(_636_v47, 2)
				(_nw120).ArraySet1(_636_v47, 3)
				(_nw120).ArraySet1(_636_v47, 4)
				(_nw120).ArraySet1(_636_v47, 5)
				(_nw120).ArraySet1(_636_v47, 6)
				(_nw120).ArraySet1(_636_v47, 7)
				(_nw120).ArraySet1(_636_v47, 8)
				(_nw120).ArraySet1(_636_v47, 9)
				(_nw120).ArraySet1(func(_pat_let8_0 D7) D7 {
					return func(_638_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let9_0 bool) D7 {
							return func(_639_dt__update_hcf29_h0 bool) D7 {
								return Companion_D7_.Create_DC21_((_638_dt__update__tmp_h0).Dtor_cf28(), _639_dt__update_hcf29_h0, (_638_dt__update__tmp_h0).Dtor_cf30())
							}(_pat_let9_0)
						}(_this.F4())
					}(_pat_let8_0)
				}(_636_v47), 10)
				(_nw120).ArraySet1(Companion_D7_.Create_DC21_(_634_v46, _this.F4(), _this.F4()), 11)
				(_nw120).ArraySet1(func(_pat_let10_0 D7) D7 {
					return func(_640_dt__update__tmp_h1 D7) D7 {
						return func(_pat_let11_0 _dafny.Array) D7 {
							return func(_641_dt__update_hcf28_h0 _dafny.Array) D7 {
								return Companion_D7_.Create_DC21_(_641_dt__update_hcf28_h0, (_640_dt__update__tmp_h1).Dtor_cf29(), (_640_dt__update__tmp_h1).Dtor_cf30())
							}(_pat_let11_0)
						}(_pat_let_tv5)
					}(_pat_let10_0)
				}(Companion_D7_.Create_DC21_(_634_v46, false, _this.F4())), 12)
				(_nw120).ArraySet1(_636_v47, 13)
				(_nw120).ArraySet1(Companion_D7_.Create_DC21_(_634_v46, _this.F4(), _this.F4()), 14)
				_637_v48 = _nw120
				var _642_v49 _dafny.Set
				_ = _642_v49
				_642_v49 = _dafny.SetOf(_637_v48, _637_v48)
				var _643_v50 _dafny.Sequence
				_ = _643_v50
				_643_v50 = _dafny.SeqOf(_dafny.SetOf(_637_v48, _637_v48))
				_642_v49 = ((_643_v50).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_643_v50).Cardinality()))).Uint32()).(_dafny.Set)).Union(_642_v49)
			}
			var _644_v52 D7
			_ = _644_v52
			_644_v52 = Companion_D7_.Create_DC20_(func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(353), _dafny.IntOfInt64(82))); ; {
					_compr_9, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _645_v51 _dafny.Int
					_645_v51 = interface{}(_compr_9).(_dafny.Int)
					if ((_dafny.IntOfInt64(353)).Cmp(_645_v51) <= 0) && ((_645_v51).Cmp(_dafny.IntOfInt64(82)) < 0) {
						_coll9.Add(Companion_Default___.SafeModuloInt(_645_v51, p0), _dafny.UnicodeSeqOfUtf8Bytes("xoulfc"))
					}
				}
				return _coll9.ToMap()
			}())
			var _source7 D7 = _644_v52
			_ = _source7
			if _source7.Is_DC21() {
				var _646___mcc_h8 _dafny.Array = _source7.Get_().(D7_DC21).Cf28
				_ = _646___mcc_h8
				var _647___mcc_h9 bool = _source7.Get_().(D7_DC21).Cf29
				_ = _647___mcc_h9
				var _648___mcc_h10 bool = _source7.Get_().(D7_DC21).Cf30
				_ = _648___mcc_h10
				var _649_cf30 bool = _648___mcc_h10
				_ = _649_cf30
				var _650_cf29 bool = _647___mcc_h9
				_ = _650_cf29
				var _651_cf28 _dafny.Array = _646___mcc_h8
				_ = _651_cf28
				(_this).F4_set_(_650_cf29)
				var _652_v53 _dafny.Map
				_ = _652_v53
				_652_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index120
				((_this).F15()).ArraySet1((p0).Plus((_652_v53).Cardinality()), (_index120).Int())
				var _653_v54 D7
				_ = _653_v54
				_653_v54 = Companion_D7_.Create_DC21_(_651_cf28, _649_cf30, !(_650_cf29))
				var _654_v55 _dafny.Map
				_ = _654_v55
				_654_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _572_i0)
				var _655_v56 D3
				_ = _655_v56
				_655_v56 = Companion_D3_.Create_DC10_((_572_i0).Plus((_dafny.Zero).Minus(_572_i0)), (_653_v54).Dtor_cf28(), (func() _dafny.Int {
					if (_654_v55).Contains(_649_cf30) {
						return (_654_v55).Get(_649_cf30).(_dafny.Int)
					}
					return (_this).F14()
				})())
				var _656_v57 _dafny.Sequence
				_ = _656_v57
				_656_v57 = _dafny.SeqOf(_651_cf28)
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((p1), 0))
				_ = _index121
				(p1).ArraySet1((_656_v57).Select((Companion_Default___.SafeIndex((_this).F3(), _dafny.IntOfUint32((_656_v57).Cardinality()))).Uint32()).(_dafny.Array), (_index121).Int())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index122
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((p1), 0))
				_ = _index123
				var _rhs121 _dafny.Int = _dafny.IntOfInt64(-163)
				_ = _rhs121
				var _rhs122 bool = (func() bool {
					if Companion_Default___.Fm2((_this).F3(), globalState) {
						return Companion_Default___.Fm2(_572_i0, globalState)
					}
					return true
				})()
				_ = _rhs122
				var _rhs123 _dafny.Int = p0
				_ = _rhs123
				var _rhs124 D3 = _655_v56
				_ = _rhs124
				var _rhs125 _dafny.Array = _651_cf28
				_ = _rhs125
				var _lhs113 *GlobalState = globalState
				_ = _lhs113
				var _lhs114 *GlobalState = globalState
				_ = _lhs114
				var _lhs115 _dafny.Array = (_this).F15()
				_ = _lhs115
				var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _lhs116
				var _lhs117 _dafny.Array = p1
				_ = _lhs117
				var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((p1), 0))
				_ = _lhs118
				_lhs113.F1 = _rhs121
				_lhs114.F2 = _rhs122
				(_lhs115).ArraySet1(_rhs123, (_lhs116).Int())
				_655_v56 = _rhs124
				(_lhs117).ArraySet1(_rhs125, (_lhs118).Int())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index124
				((_this).F15()).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(globalState), (func() _dafny.Int {
					if _650_cf29 {
						return ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(_dafny.Int)
					}
					return (_this).F14()
				})()), (_index124).Int())
				var _657_v58 *C0
				_ = _657_v58
				var _nw121 *C0 = New_C0_()
				_ = _nw121
				_nw121.Ctor__(_649_cf30, _dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((p1), 0))).Int())))
				_657_v58 = _nw121
			} else {
				var _658___mcc_h11 _dafny.Map = _source7.Get_().(D7_DC20).Cf27
				_ = _658___mcc_h11
				var _659_cf27 _dafny.Map = _658___mcc_h11
				_ = _659_cf27
				var _660_v59 _dafny.Sequence
				_ = _660_v59
				_660_v59 = _dafny.SeqOf(_this.F4(), false)
				r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_660_v59, _dafny.Companion_Sequence_.Concatenate(_660_v59, _660_v59))).Cardinality())
				var _661_v60 _dafny.Sequence
				_ = _661_v60
				_661_v60 = _dafny.UnicodeSeqOfUtf8Bytes("yqnnbe")
				r0 = _dafny.IntOfUint32((_661_v60).Cardinality())
				var _662_v61 _dafny.Map
				_ = _662_v61
				_662_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F4())
				var _663_v62 D6
				_ = _663_v62
				_663_v62 = Companion_D6_.Create_DC18_(_this.F4(), p0)
				var _664_v63 _dafny.Set
				_ = _664_v63
				_664_v63 = _dafny.SetOf(_572_i0, _572_i0, (_662_v61).Cardinality(), (_663_v62).Dtor_cf25(), _572_i0)
				var _665_v64 _dafny.Array
				_ = _665_v64
				var _nwElement0_22 bool = _this.F4()
				_ = _nwElement0_22
				var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(15))
				_ = _nw122
				(_nw122).ArraySet1(_nwElement0_22, 0)
				(_nw122).ArraySet1(((_664_v63).Cardinality()).Cmp(_dafny.IntOfInt64(656)) < 0, 1)
				(_nw122).ArraySet1(_this.F4(), 2)
				(_nw122).ArraySet1((_662_v61).Contains(Companion_Default___.Fm2((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState))), globalState)), 3)
				(_nw122).ArraySet1((func() bool {
					if _this.F4() {
						return _this.F4()
					}
					return !(true)
				})(), 4)
				(_nw122).ArraySet1(((_this).F3()).Cmp(p0) <= 0, 5)
				(_nw122).ArraySet1(_this.F4(), 6)
				(_nw122).ArraySet1(_this.F4(), 7)
				(_nw122).ArraySet1(!(false), 8)
				(_nw122).ArraySet1((_662_v61).Contains(!(_this.F4())), 9)
				(_nw122).ArraySet1(_this.F4(), 10)
				(_nw122).ArraySet1(_this.F4(), 11)
				(_nw122).ArraySet1(_this.F4(), 12)
				(_nw122).ArraySet1(_this.F4(), 13)
				(_nw122).ArraySet1(_this.F4(), 14)
				_665_v64 = _nw122
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_665_v64), 0))
				_ = _index125
				(_665_v64).ArraySet1(_this.F4(), (_index125).Int())
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_665_v64), 0))
				_ = _index126
				(_665_v64).ArraySet1(_this.F4(), (_index126).Int())
				_665_v64 = _665_v64
			}
			var _666_v65 _dafny.Sequence
			_ = _666_v65
			_666_v65 = _dafny.UnicodeSeqOfUtf8Bytes("png")
			var _667_v66 _dafny.Map
			_ = _667_v66
			_667_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_666_v65).Cardinality()), _this.F4())
			var _668_v67 *C1
			_ = _668_v67
			var _nw123 *C1 = New_C1_()
			_ = _nw123
			_nw123.Ctor__((_dafny.Zero).Minus(_572_i0), (_667_v66).Cardinality(), _this.F4())
			_668_v67 = _nw123
		}
		var _669_v68 _dafny.CodePoint
		_ = _669_v68
		_669_v68 = _dafny.CodePoint('l')
		var _pat_let_tv6 = _669_v68
		_ = _pat_let_tv6
		var _pat_let_tv7 = _669_v68
		_ = _pat_let_tv7
		var _pat_let_tv8 = _669_v68
		_ = _pat_let_tv8
		_669_v68 = func(_source8 D1) _dafny.CodePoint {
			if _source8.Is_DC6() {
				var _670___mcc_h12 T0 = _source8.Get_().(D1_DC6).Cf9
				_ = _670___mcc_h12
				var _671___mcc_h13 _dafny.Sequence = _source8.Get_().(D1_DC6).Cf10
				_ = _671___mcc_h13
				var _672___mcc_h14 bool = _source8.Get_().(D1_DC6).Cf11
				_ = _672___mcc_h14
				var _673___mcc_h15 bool = _source8.Get_().(D1_DC6).Cf12
				_ = _673___mcc_h15
				var _674_cf12 bool = _673___mcc_h15
				_ = _674_cf12
				var _675_cf11 bool = _672___mcc_h14
				_ = _675_cf11
				var _676_cf10 _dafny.Sequence = _671___mcc_h13
				_ = _676_cf10
				var _677_cf9 T0 = _670___mcc_h12
				_ = _677_cf9
				if _674_cf12 {
					return _pat_let_tv6
				} else {
					return _dafny.CodePoint('c')
				}
			} else if _source8.Is_DC7() {
				return _pat_let_tv7
			} else {
				var _678___mcc_h16 _dafny.Int = _source8.Get_().(D1_DC5).Cf8
				_ = _678___mcc_h16
				var _679_cf8 _dafny.Int = _678___mcc_h16
				_ = _679_cf8
				return _pat_let_tv8
			}
		}(Companion_D1_.Create_DC5_((_this).F14()))
		r0 = (p0).Times((_this).F14())
		(globalState).F2 = _this.F4()
		r0 = (func() _dafny.Int {
			if (_dafny.IntOfInt64(906)).Cmp(p0) < 0 {
				return (_this).F14()
			}
			return Companion_Default___.SafeDivisionInt((_this).F14(), (_dafny.Zero).Minus((_this).F3()))
		})()
		r1 = Companion_Default___.Fm2(Companion_Default___.SafeModuloInt((_this).F14(), p0), globalState)
		return r0, r1
	}
}
func (_this *C2) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *C2) F15() _dafny.Array {
	{
		return _this._f15
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f4  bool
	_f7  _dafny.Int
	_f3  _dafny.Int
	F13  _dafny.Sequence
	_f12 _dafny.CodePoint
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f4 = false
	_this._f7 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this.F13 = _dafny.EmptySeq
	_this._f12 = _dafny.CodePoint('D')
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C3{}
var _ T1 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F4() bool {
	return _this._f4
}
func (_this *C3) F4_set_(value bool) {
	_this._f4 = value
}
func (_this *C3) F7() _dafny.Int {
	return _this._f7
}
func (_this *C3) F3() _dafny.Int {
	return _this._f3
}
func (_this *C3) Ctor__(f12 _dafny.CodePoint, f13 _dafny.Sequence, f3 _dafny.Int, f4 bool, f7 _dafny.Int) {
	{
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f7 = f7
	}
}
func (_this *C3) Fm12(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_this).F7(), _dafny.IntOfInt64(14))
	}
}
func (_this *C3) Fm13(globalState *GlobalState) bool {
	{
		return _this.F4()
	}
}
func (_this *C3) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _680_v0 _dafny.Set
		_ = _680_v0
		_680_v0 = _dafny.SetOf(((_this).F7()).Minus((_this).F7()), (_this).F7(), _dafny.IntOfUint32((_this.F13).Cardinality()))
		_680_v0 = _dafny.SetOf(Companion_Default___.SafeDivisionInt((_this).F3(), (_this).F7()))
		var _681_v1 _dafny.Map
		_ = _681_v1
		_681_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_Default___.Fm14(!(_this.F4()), (_this).Fm13(globalState), globalState)), _this.F4())
		var _682_v2 _dafny.MultiSet
		_ = _682_v2
		_682_v2 = _dafny.MultiSetOf(true)
		(globalState).F2 = (func() bool {
			if (_681_v1).Contains(_dafny.SeqOf(_dafny.MultiSetOf(_this.F4()), _682_v2)) {
				return (_681_v1).Get(_dafny.SeqOf(_dafny.MultiSetOf(_this.F4()), _682_v2)).(bool)
			}
			return ((_this).F3()).Cmp(p0) < 0
		})()
		var _683_i0 _dafny.Int
		_ = _683_i0
		_683_i0 = _dafny.Zero
		{
			for _this.F4() {
				{
					if (_683_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_683_i0 = (_683_i0).Plus(_dafny.One)
					(_this).M5(globalState)
					(globalState).F2 = !((_this).Fm13(globalState))
					if _this.F4() {
						var _684_v3 D0
						_ = _684_v3
						_684_v3 = Companion_D0_.Create_DC0_(_this.F4())
						var _685_v4 _dafny.Array
						_ = _685_v4
						var _nwElement0_23 bool = !(false)
						_ = _nwElement0_23
						var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(23))
						_ = _nw124
						(_nw124).ArraySet1(_nwElement0_23, 0)
						(_nw124).ArraySet1(_this.F4(), 1)
						(_nw124).ArraySet1((_684_v3).Dtor_cf0(), 2)
						(_nw124).ArraySet1(_this.F4(), 3)
						(_nw124).ArraySet1(_this.F4(), 4)
						(_nw124).ArraySet1(_this.F4(), 5)
						(_nw124).ArraySet1(_this.F4(), 6)
						(_nw124).ArraySet1(_this.F4(), 7)
						(_nw124).ArraySet1(_this.F4(), 8)
						(_nw124).ArraySet1(_this.F4(), 9)
						(_nw124).ArraySet1(_this.F4(), 10)
						(_nw124).ArraySet1(_this.F4(), 11)
						(_nw124).ArraySet1(false, 12)
						(_nw124).ArraySet1(_this.F4(), 13)
						(_nw124).ArraySet1(_this.F4(), 14)
						(_nw124).ArraySet1(_this.F4(), 15)
						(_nw124).ArraySet1(_this.F4(), 16)
						(_nw124).ArraySet1(false, 17)
						(_nw124).ArraySet1(false, 18)
						(_nw124).ArraySet1(_this.F4(), 19)
						(_nw124).ArraySet1(_this.F4(), 20)
						(_nw124).ArraySet1(false, 21)
						(_nw124).ArraySet1(_this.F4(), 22)
						_685_v4 = _nw124
						var _686_v5 *C0
						_ = _686_v5
						var _nw125 *C0 = New_C0_()
						_ = _nw125
						_nw125.Ctor__((func() bool {
							if !(_this.F4()) {
								return _this.F4()
							}
							return _this.F4()
						})(), _685_v4)
						_686_v5 = _nw125
						var _687_v6 _dafny.Sequence
						_ = _687_v6
						_687_v6 = _dafny.SeqOf((_this).F3())
						var _688_v7 *C0
						_ = _688_v7
						var _nw126 *C0 = New_C0_()
						_ = _nw126
						_nw126.Ctor__((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm15(!(_686_v5.F10), (_this).F3(), _dafny.IntOfUint32((_687_v6).Cardinality()), globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm15(!(_686_v5.F10), (_this).F3(), _dafny.IntOfUint32((_687_v6).Cardinality()), globalState)).Cardinality()))).Uint32(), (_this).F12())).Cardinality())).Cmp(p0) == 0, (_686_v5).F11())
						_688_v7 = _nw126
						var _689_v8 D4
						_ = _689_v8
						_689_v8 = Companion_D4_.Create_DC11_(_680_v0)
						var _690_v9 D4
						_ = _690_v9
						_690_v9 = Companion_D4_.Create_DC11_((_689_v8).Dtor_cf18())
						_680_v0 = (_689_v8).Dtor_cf18()
						var _691_v10 _dafny.Array
						_ = _691_v10
						var _len0_23 _dafny.Int = _dafny.IntOfInt64(5)
						_ = _len0_23
						var _nw127 _dafny.Array
						_ = _nw127
						if _len0_23.Cmp(_dafny.Zero) == 0 {
							_nw127 = _dafny.NewArray(_len0_23)
						} else {
							var _init23 func(_dafny.Int) _dafny.Int = (func(_692_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_693_i1 _dafny.Int) _dafny.Int {
									return (_693_i1).Times(_692_p0)
								}
							})(p0)
							_ = _init23
							var _element0_23 = _init23(_dafny.Zero)
							_ = _element0_23
							_nw127 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
							(_nw127).ArraySet1(_element0_23, 0)
							var _nativeLen0_23 = (_len0_23).Int()
							_ = _nativeLen0_23
							for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
								(_nw127).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
							}
						}
						_691_v10 = _nw127
						_691_v10 = _691_v10
						var _694_v11 _dafny.MultiSet
						_ = _694_v11
						_694_v11 = _dafny.MultiSetOf((_688_v7).F11(), (_688_v7).F11(), (_688_v7).F11(), (_686_v5).F11())
						var _695_v12 _dafny.Map
						_ = _695_v12
						_695_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_688_v7.F10, (_688_v7).Fm5(_dafny.IntOfInt64(292), p0, (_this).F12(), _688_v7.F10, globalState))
						var _696_v13 _dafny.Map
						_ = _696_v13
						_696_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_695_v12).Cardinality(), _686_v5.F10)
						var _rhs126 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F3(), ((_696_v13).Cardinality()).Minus(p0))
						_ = _rhs126
						var _rhs127 _dafny.MultiSet = (_694_v11).Difference(_694_v11)
						_ = _rhs127
						var _rhs128 _dafny.Int = (_this).F7()
						_ = _rhs128
						var _rhs129 _dafny.Array = _691_v10
						_ = _rhs129
						var _rhs130 _dafny.Sequence = _687_v6
						_ = _rhs130
						var _lhs119 *GlobalState = globalState
						_ = _lhs119
						r0 = _rhs126
						_694_v11 = _rhs127
						_lhs119.F1 = _rhs128
						_691_v10 = _rhs129
						_687_v6 = _rhs130
					} else {
						var _697_v14 _dafny.Array
						_ = _697_v14
						var _nwElement0_24 _dafny.CodePoint = (_this).F12()
						_ = _nwElement0_24
						var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(7))
						_ = _nw128
						(_nw128).ArraySet1CodePoint(_nwElement0_24, 0)
						(_nw128).ArraySet1CodePoint((_this).F12(), 1)
						(_nw128).ArraySet1CodePoint((_this).F12(), 2)
						(_nw128).ArraySet1CodePoint((_this).F12(), 3)
						(_nw128).ArraySet1CodePoint((_this).F12(), 4)
						(_nw128).ArraySet1CodePoint((_this).F12(), 5)
						(_nw128).ArraySet1CodePoint((_this).F12(), 6)
						_697_v14 = _nw128
						var _698_v15 _dafny.Sequence
						_ = _698_v15
						_698_v15 = _dafny.SeqOf(_697_v14)
						var _699_v16 _dafny.MultiSet
						_ = _699_v16
						_699_v16 = _dafny.MultiSetOf((_this).F3())
						var _700_v17 _dafny.Array
						_ = _700_v17
						var _nwElement0_25 _dafny.Array = _697_v14
						_ = _nwElement0_25
						var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(13))
						_ = _nw129
						(_nw129).ArraySet1(_nwElement0_25, 0)
						(_nw129).ArraySet1(_697_v14, 1)
						(_nw129).ArraySet1(_697_v14, 2)
						(_nw129).ArraySet1((_698_v15).Select((Companion_Default___.SafeIndex((_699_v16).Cardinality(), _dafny.IntOfUint32((_698_v15).Cardinality()))).Uint32()).(_dafny.Array), 3)
						(_nw129).ArraySet1(_697_v14, 4)
						(_nw129).ArraySet1((_698_v15).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_698_v15).Cardinality()))).Uint32()).(_dafny.Array), 5)
						(_nw129).ArraySet1(_697_v14, 6)
						(_nw129).ArraySet1(_697_v14, 7)
						(_nw129).ArraySet1(_697_v14, 8)
						(_nw129).ArraySet1(_697_v14, 9)
						(_nw129).ArraySet1(_697_v14, 10)
						(_nw129).ArraySet1(_697_v14, 11)
						(_nw129).ArraySet1(_697_v14, 12)
						_700_v17 = _nw129
						var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_700_v17), 0))
						_ = _index127
						(_700_v17).ArraySet1(_697_v14, (_index127).Int())
						var _701_v18 D3
						_ = _701_v18
						_701_v18 = Companion_D3_.Create_DC9_(_697_v14)
						var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_700_v17), 0))
						_ = _index128
						var _rhs131 _dafny.Array = (_701_v18).Dtor_cf14()
						_ = _rhs131
						var _rhs132 bool = (_this.F4()) || (_this.F4())
						_ = _rhs132
						var _lhs120 _dafny.Array = _700_v17
						_ = _lhs120
						var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_700_v17), 0))
						_ = _lhs121
						var _lhs122 *GlobalState = globalState
						_ = _lhs122
						(_lhs120).ArraySet1(_rhs131, (_lhs121).Int())
						_lhs122.F2 = _rhs132
						var _702_v19 _dafny.Array
						_ = _702_v19
						var _nw130 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
						_ = _nw130
						_702_v19 = _nw130
						var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_702_v19), 0))
						_ = _index129
						(_702_v19).ArraySet1(_this.F4(), (_index129).Int())
						var _703_v20 _dafny.Sequence
						_ = _703_v20
						_703_v20 = _dafny.SeqOf(p0)
						var _704_v21 D1
						_ = _704_v21
						_704_v21 = Companion_D1_.Create_DC5_((_this).F7())
						var _705_v22 _dafny.Map
						_ = _705_v22
						_705_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v21, (_this).F3())
						var _706_v23 D1
						_ = _706_v23
						_706_v23 = Companion_D1_.Create_DC5_((func() _dafny.Int {
							if (_705_v22).Contains(_704_v21) {
								return (_705_v22).Get(_704_v21).(_dafny.Int)
							}
							return (_this).F7()
						})())
						var _707_v24 _dafny.Sequence
						_ = _707_v24
						_707_v24 = _dafny.SeqOf(_this.F4(), _this.F4(), _this.F4())
						var _708_v25 _dafny.Map
						_ = _708_v25
						_708_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_707_v24).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_707_v24).Cardinality()))).Uint32()).(bool), p0)
						var _709_v26 _dafny.Map
						_ = _709_v26
						_709_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _708_v25)
						var _710_v27 _dafny.Array
						_ = _710_v27
						var _nwElement0_26 _dafny.Int = (_this).F7()
						_ = _nwElement0_26
						var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(21))
						_ = _nw131
						(_nw131).ArraySet1(_nwElement0_26, 0)
						(_nw131).ArraySet1((_this).F3(), 1)
						(_nw131).ArraySet1((_this).F7(), 2)
						(_nw131).ArraySet1((_this).F3(), 3)
						(_nw131).ArraySet1(p0, 4)
						(_nw131).ArraySet1(((_this).F7()).Minus(_dafny.IntOfUint32((_703_v20).Cardinality())), 5)
						(_nw131).ArraySet1(p0, 6)
						(_nw131).ArraySet1(((_this).F7()).Minus(_dafny.IntOfInt64(166)), 7)
						(_nw131).ArraySet1((_this).F7(), 8)
						(_nw131).ArraySet1((_704_v21).Dtor_cf8(), 9)
						(_nw131).ArraySet1(Companion_Default___.Fm0(globalState), 10)
						(_nw131).ArraySet1(_dafny.IntOfInt64(-367), 11)
						(_nw131).ArraySet1((_dafny.IntOfInt64(407)).Times((_this).F3()), 12)
						(_nw131).ArraySet1((_709_v26).Cardinality(), 13)
						(_nw131).ArraySet1(Companion_Default___.SafeModuloInt((Companion_D1_.Create_DC5_((_this).F7())).Dtor_cf8(), p0), 14)
						(_nw131).ArraySet1(_dafny.IntOfInt64(-272), 15)
						(_nw131).ArraySet1(p0, 16)
						(_nw131).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_707_v24).Cardinality()), (_this).F3()), 17)
						(_nw131).ArraySet1((_this).F7(), 18)
						(_nw131).ArraySet1(Companion_Default___.SafeModuloInt((_this).F7(), (_this).F7()), 19)
						(_nw131).ArraySet1((_this).F3(), 20)
						_710_v27 = _nw131
						var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_702_v19), 0))
						_ = _index130
						var _rhs133 bool = _this.F4()
						_ = _rhs133
						var _rhs134 _dafny.Array = _710_v27
						_ = _rhs134
						var _lhs123 _dafny.Array = _702_v19
						_ = _lhs123
						var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_702_v19), 0))
						_ = _lhs124
						(_lhs123).ArraySet1(_rhs133, (_lhs124).Int())
						_710_v27 = _rhs134
						var _711_v28 *C0
						_ = _711_v28
						var _nw132 *C0 = New_C0_()
						_ = _nw132
						_nw132.Ctor__((_702_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_702_v19), 0))).Int()).(bool), _702_v19)
						_711_v28 = _nw132
						(globalState).F1 = p0
						var _712_v29 _dafny.Array
						_ = _712_v29
						var _len0_24 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_24
						var _nw133 _dafny.Array
						_ = _nw133
						if _len0_24.Cmp(_dafny.Zero) == 0 {
							_nw133 = _dafny.NewArray(_len0_24)
						} else {
							var _init24 func(_dafny.Int) _dafny.Set = (func(_713_v0 _dafny.Set) func(_dafny.Int) _dafny.Set {
								return func(_714_i2 _dafny.Int) _dafny.Set {
									return _713_v0
								}
							})(_680_v0)
							_ = _init24
							var _element0_24 = _init24(_dafny.Zero)
							_ = _element0_24
							_nw133 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
							(_nw133).ArraySet1(_element0_24, 0)
							var _nativeLen0_24 = (_len0_24).Int()
							_ = _nativeLen0_24
							for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
								(_nw133).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
							}
						}
						_712_v29 = _nw133
						var _715_v30 _dafny.Array
						_ = _715_v30
						var _nw134 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(3))
						_ = _nw134
						_715_v30 = _nw134
						var _716_v31 D3
						_ = _716_v31
						_716_v31 = Companion_D3_.Create_DC10_((_dafny.Zero).Minus(p0), (_711_v28).F11(), (_this).F7())
						var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_715_v30), 0))
						_ = _index131
						(_715_v30).ArraySet1(_716_v31, (_index131).Int())
						var _717_v32 _dafny.Sequence
						_ = _717_v32
						_717_v32 = _dafny.SeqOf(_this.F13)
						var _718_v33 _dafny.MultiSet
						_ = _718_v33
						_718_v33 = _dafny.MultiSetOf(_this.F13, _this.F13)
						var _pat_let_tv9 = _702_v19
						_ = _pat_let_tv9
						var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_702_v19), 0))
						_ = _index132
						var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_715_v30), 0))
						_ = _index133
						var _rhs135 _dafny.Array = _712_v29
						_ = _rhs135
						var _rhs136 bool = ((_this).F7()).Cmp(((func() _dafny.MultiSet {
							if true {
								return _dafny.MultiSetFromSeq(_717_v32)
							}
							return _718_v33
						})()).Cardinality()) >= 0
						_ = _rhs136
						var _rhs137 bool = (p0).Cmp(Companion_Default___.SafeModuloInt(p0, (_this).F7())) == 0
						_ = _rhs137
						var _rhs138 D3 = func(_pat_let12_0 D3) D3 {
							return func(_719_dt__update__tmp_h0 D3) D3 {
								return func(_pat_let13_0 _dafny.Int) D3 {
									return func(_720_dt__update_hcf15_h0 _dafny.Int) D3 {
										return func(_pat_let14_0 _dafny.Array) D3 {
											return func(_721_dt__update_hcf16_h0 _dafny.Array) D3 {
												return Companion_D3_.Create_DC10_(_720_dt__update_hcf15_h0, _721_dt__update_hcf16_h0, (_719_dt__update__tmp_h0).Dtor_cf17())
											}(_pat_let14_0)
										}(_pat_let_tv9)
									}(_pat_let13_0)
								}((_dafny.IntOfInt64(997)).Plus((_this).F3()))
							}(_pat_let12_0)
						}(_716_v31)
						_ = _rhs138
						var _lhs125 _dafny.Array = _702_v19
						_ = _lhs125
						var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_702_v19), 0))
						_ = _lhs126
						var _lhs127 *GlobalState = globalState
						_ = _lhs127
						var _lhs128 _dafny.Array = _715_v30
						_ = _lhs128
						var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_715_v30), 0))
						_ = _lhs129
						_712_v29 = _rhs135
						(_lhs125).ArraySet1(_rhs136, (_lhs126).Int())
						_lhs127.F2 = _rhs137
						(_lhs128).ArraySet1(_rhs138, (_lhs129).Int())
					}
					var _722_v34 _dafny.Array
					_ = _722_v34
					var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
					_ = _nw135
					_722_v34 = _nw135
					var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_722_v34), 0))
					_ = _index134
					(_722_v34).ArraySet1(p0, (_index134).Int())
					var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_722_v34), 0))
					_ = _index135
					(_722_v34).ArraySet1((_this).F7(), (_index135).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _723_i3 _dafny.Int
		_ = _723_i3
		_723_i3 = _dafny.Zero
		{
			for _this.F4() {
				{
					if (_723_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_723_i3 = (_723_i3).Plus(_dafny.One)
					var _724_v35 _dafny.Sequence
					_ = _724_v35
					_724_v35 = _dafny.SeqOf(p0)
					var _725_v36 _dafny.Map
					_ = _725_v36
					_725_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F4()), _724_v35)
					_725_v36 = (_725_v36).Update(!(_this.F4()), _724_v35)
					_680_v0 = _680_v0
					var _726_v37 _dafny.Map
					_ = _726_v37
					_726_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _dafny.IntOfInt64(637))
					var _727_v38 _dafny.Array
					_ = _727_v38
					var _nwElement0_27 bool = (_726_v37).Contains(!(_this.F4()))
					_ = _nwElement0_27
					var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(5))
					_ = _nw136
					(_nw136).ArraySet1(_nwElement0_27, 0)
					(_nw136).ArraySet1(true, 1)
					(_nw136).ArraySet1(_this.F4(), 2)
					(_nw136).ArraySet1(true, 3)
					(_nw136).ArraySet1(_this.F4(), 4)
					_727_v38 = _nw136
					var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_727_v38), 0))
					_ = _index136
					(_727_v38).ArraySet1(!(_this.F4()), (_index136).Int())
					var _728_v39 _dafny.Map
					_ = _728_v39
					_728_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(546), (_this).F3())
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_727_v38), 0))
					_ = _index137
					(_727_v38).ArraySet1(((_this).F7()).Cmp((_728_v39).Cardinality()) == 0, (_index137).Int())
					var _729_v40 *C0
					_ = _729_v40
					var _nw137 *C0 = New_C0_()
					_ = _nw137
					_nw137.Ctor__(((_this).F7()).Cmp((_this).F3()) >= 0, _727_v38)
					_729_v40 = _nw137
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		(globalState).F2 = false
		var _730_v41 D1
		_ = _730_v41
		_730_v41 = Companion_D1_.Create_DC7_()
		var _rhs139 _dafny.Int = (_this).Fm12(Companion_Default___.Fm2((_680_v0).Cardinality(), globalState), _this.F4(), globalState)
		_ = _rhs139
		var _rhs140 bool = func(_source9 D1) bool {
			if _source9.Is_DC6() {
				var _731___mcc_h0 T0 = _source9.Get_().(D1_DC6).Cf9
				_ = _731___mcc_h0
				var _732___mcc_h1 _dafny.Sequence = _source9.Get_().(D1_DC6).Cf10
				_ = _732___mcc_h1
				var _733___mcc_h2 bool = _source9.Get_().(D1_DC6).Cf11
				_ = _733___mcc_h2
				var _734___mcc_h3 bool = _source9.Get_().(D1_DC6).Cf12
				_ = _734___mcc_h3
				var _735_cf12 bool = _734___mcc_h3
				_ = _735_cf12
				var _736_cf11 bool = _733___mcc_h2
				_ = _736_cf11
				var _737_cf10 _dafny.Sequence = _732___mcc_h1
				_ = _737_cf10
				var _738_cf9 T0 = _731___mcc_h0
				_ = _738_cf9
				if !(false) {
					return false
				} else {
					return true
				}
			} else if _source9.Is_DC7() {
				return _this.F4()
			} else {
				var _739___mcc_h4 _dafny.Int = _source9.Get_().(D1_DC5).Cf8
				_ = _739___mcc_h4
				var _740_cf8 _dafny.Int = _739___mcc_h4
				_ = _740_cf8
				return (_dafny.MultiSetOf((_this).F3())).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_741_cf8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_742_i4 _dafny.Int) _dafny.Int {
						return _741_cf8
					}
				})(_740_cf8)))).Cardinality()))
			}
		}(_730_v41)
		_ = _rhs140
		var _rhs141 _dafny.Int = ((p0).Minus(_dafny.IntOfUint32((_this.F13).Cardinality()))).Times((_this).F3())
		_ = _rhs141
		var _lhs130 *GlobalState = globalState
		_ = _lhs130
		var _lhs131 *GlobalState = globalState
		_ = _lhs131
		var _lhs132 *GlobalState = globalState
		_ = _lhs132
		_lhs130.F1 = _rhs139
		_lhs131.F2 = _rhs140
		_lhs132.F1 = _rhs141
		var _743_v42 _dafny.Map
		_ = _743_v42
		_743_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_682_v2, (_this).F7())
		r0 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_743_v42).Contains(_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F4()))) {
				return (_743_v42).Get(_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F4()))).(_dafny.Int)
			}
			return (_this).F3()
		})(), (_680_v0).Cardinality()), (_this).F3())
		return r0
	}
}
func (_this *C3) M5(globalState *GlobalState) {
	{
		var _hi6 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F7(), (_this).F3())
		_ = _hi6
		for _744_i0 := (_this).F7(); _744_i0.Cmp(_hi6) < 0; _744_i0 = _744_i0.Plus(_dafny.One) {
			var _745_v0 _dafny.Array
			_ = _745_v0
			var _nw138 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw138
			_745_v0 = _nw138
			var _746_v1 *C0
			_ = _746_v1
			var _nw139 *C0 = New_C0_()
			_ = _nw139
			_nw139.Ctor__(!(((_dafny.Zero).Minus((_this).F7())).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_747_i1 _dafny.Int) _dafny.CodePoint {
				return (_this).F12()
			}))).Cardinality())) < 0), _745_v0)
			_746_v1 = _nw139
			(globalState).F1 = (_this).F3()
			var _748_v2 _dafny.Sequence
			_ = _748_v2
			_748_v2 = _dafny.SeqOf(_746_v1.F10)
			var _749_v3 _dafny.MultiSet
			_ = _749_v3
			_749_v3 = _dafny.MultiSetOf(_748_v2, _748_v2)
			(globalState).F1 = (((_749_v3).Intersection(_749_v3)).Difference((Companion_Default___.Fm16(_this.F4(), _this.F4(), globalState)).Difference(Companion_Default___.Fm16(_746_v1.F10, true, globalState)))).Cardinality()
			var _750_v4 _dafny.Map
			_ = _750_v4
			_750_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_746_v1.F10, _this.F4())
			var _751_v5 *C0
			_ = _751_v5
			var _nw140 *C0 = New_C0_()
			_ = _nw140
			_nw140.Ctor__((func() bool {
				if (_750_v4).Contains(_this.F4()) {
					return (_750_v4).Get(_this.F4()).(bool)
				}
				return _746_v1.F10
			})(), (_746_v1).F11())
			_751_v5 = _nw140
		}
		var _752_v6 _dafny.Array
		_ = _752_v6
		var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw141
		_752_v6 = _nw141
		var _753_v7 D0
		_ = _753_v7
		_753_v7 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_752_v6, _752_v6, _752_v6), _dafny.SeqOf(_752_v6)))
		var _source10 D0 = _753_v7
		_ = _source10
		if _source10.Is_DC1() {
			var _754___mcc_h0 bool = _source10.Get_().(D0_DC1).Cf1
			_ = _754___mcc_h0
			var _755___mcc_h1 bool = _source10.Get_().(D0_DC1).Cf2
			_ = _755___mcc_h1
			var _756___mcc_h2 _dafny.CodePoint = _source10.Get_().(D0_DC1).Cf3
			_ = _756___mcc_h2
			var _757___mcc_h3 _dafny.Array = _source10.Get_().(D0_DC1).Cf4
			_ = _757___mcc_h3
			var _758_cf4 _dafny.Array = _757___mcc_h3
			_ = _758_cf4
			var _759_cf3 _dafny.CodePoint = _756___mcc_h2
			_ = _759_cf3
			var _760_cf2 bool = _755___mcc_h1
			_ = _760_cf2
			var _761_cf1 bool = _754___mcc_h0
			_ = _761_cf1
			(globalState).F1 = _dafny.IntOfInt64(-77)
			var _762_v8 _dafny.Array
			_ = _762_v8
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_25
			var _nw142 _dafny.Array
			_ = _nw142
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw142 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) bool = (func(_763_cf1 bool) func(_dafny.Int) bool {
					return func(_764_i2 _dafny.Int) bool {
						return _763_cf1
					}
				})(_761_cf1)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw142 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw142).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw142).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_762_v8 = _nw142
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_762_v8), 0))
			_ = _index138
			(_762_v8).ArraySet1(_760_cf2, (_index138).Int())
			var _765_v9 _dafny.Sequence
			_ = _765_v9
			_765_v9 = _dafny.SeqOf(_752_v6)
			var _766_v10 _dafny.MultiSet
			_ = _766_v10
			_766_v10 = _dafny.MultiSetOf((_765_v9).Select((Companion_Default___.SafeIndex((_this).F3(), _dafny.IntOfUint32((_765_v9).Cardinality()))).Uint32()).(_dafny.Array))
			var _767_v11 _dafny.Map
			_ = _767_v11
			_767_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_760_cf2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_this.F13, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_766_v10).Cardinality()), _dafny.IntOfUint32((_this.F13).Cardinality()))).Uint32(), (_this).F12())).Cardinality()))
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_762_v8), 0))
			_ = _index139
			var _rhs142 bool = _761_cf1
			_ = _rhs142
			var _rhs143 _dafny.Map = _767_v11
			_ = _rhs143
			var _lhs133 _dafny.Array = _762_v8
			_ = _lhs133
			var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_762_v8), 0))
			_ = _lhs134
			(_lhs133).ArraySet1(_rhs142, (_lhs134).Int())
			_767_v11 = _rhs143
			(globalState).F2 = ((_this).F7()).Cmp(((_this).F7()).Minus((_this).F7())) <= 0
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_762_v8), 0))
			_ = _index140
			(_762_v8).ArraySet1(_760_cf2, (_index140).Int())
		} else if _source10.Is_DC2() {
			var _768___mcc_h4 _dafny.MultiSet = _source10.Get_().(D0_DC2).Cf5
			_ = _768___mcc_h4
			var _769___mcc_h5 bool = _source10.Get_().(D0_DC2).Cf6
			_ = _769___mcc_h5
			var _770_cf6 bool = _769___mcc_h5
			_ = _770_cf6
			var _771_cf5 _dafny.MultiSet = _768___mcc_h4
			_ = _771_cf5
			var _772_v12 _dafny.Sequence
			_ = _772_v12
			_772_v12 = _dafny.SeqOf(_770_cf6, _770_cf6)
			var _773_v13 _dafny.Sequence
			_ = _773_v13
			_773_v13 = _dafny.SeqOf(_dafny.SeqOf(_this.F4()), _dafny.Companion_Sequence_.Update(_772_v12, (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_772_v12).Cardinality()))).Uint32(), _this.F4()), _772_v12, _772_v12)
			var _774_v14 D0
			_ = _774_v14
			_774_v14 = Companion_D0_.Create_DC1_(false, _this.F4(), (_this).F12(), _752_v6)
			var _775_v15 _dafny.Map
			_ = _775_v15
			_775_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_773_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm15(_770_cf6, (_dafny.Zero).Minus((_this).F3()), _dafny.IntOfUint32((_this.F13).Cardinality()), globalState)).Cardinality()), _dafny.IntOfUint32((_773_v13).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_this.F13, _dafny.Companion_Sequence_.Update(_this.F13, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F3()), _dafny.IntOfUint32((_this.F13).Cardinality()))).Uint32(), (_774_v14).Dtor_cf3())))
			_775_v15 = _775_v15
			(_this).F4_set_(_770_cf6)
			var _776_v16 _dafny.Array
			_ = _776_v16
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_26
			var _nw143 _dafny.Array
			_ = _nw143
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw143 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) bool = (func(_777_cf6 bool) func(_dafny.Int) bool {
					return func(_778_i3 _dafny.Int) bool {
						return _777_cf6
					}
				})(_770_cf6)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw143 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw143).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw143).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_776_v16 = _nw143
			var _779_v17 *C0
			_ = _779_v17
			var _nw144 *C0 = New_C0_()
			_ = _nw144
			_nw144.Ctor__(_this.F4(), _776_v16)
			_779_v17 = _nw144
			var _780_v18 _dafny.Map
			_ = _780_v18
			_780_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _770_cf6)
			var _781_v20 _dafny.Sequence
			_ = _781_v20
			_781_v20 = _dafny.SeqOf(_780_v18, (func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter12 := _dafny.Iterate((_this.F13).Elements()); ; {
					_compr_10, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _782_v19 _dafny.CodePoint
					_782_v19 = interface{}(_compr_10).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_this.F13, _782_v19) {
						_coll10.Add(_782_v19, _779_v17.F10)
					}
				}
				return _coll10.ToMap()
			}()).Update((_this).F12(), _this.F4()), Companion_Default___.Fm17(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_779_v17.F10, _779_v17.F10), globalState), _780_v18, _780_v18)
			(globalState).F1 = (((_781_v20).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_781_v20).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Plus((_this).F3())
		} else if _source10.Is_DC3() {
			var _783_v21 _dafny.Array
			_ = _783_v21
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
			_ = _nw145
			_783_v21 = _nw145
			var _784_v22 D3
			_ = _784_v22
			_784_v22 = Companion_D3_.Create_DC9_(_783_v21)
			var _785_v23 _dafny.Sequence
			_ = _785_v23
			_785_v23 = _dafny.SeqOf(_784_v22, _784_v22, _784_v22, _784_v22)
			var _786_v24 _dafny.Array
			_ = _786_v24
			var _nwElement0_28 _dafny.Sequence = _785_v23
			_ = _nwElement0_28
			var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(29))
			_ = _nw146
			(_nw146).ArraySet1(_nwElement0_28, 0)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22), 1)
			(_nw146).ArraySet1(_785_v23, 2)
			(_nw146).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_785_v23, _785_v23), 3)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22, _784_v22, _784_v22), 4)
			(_nw146).ArraySet1(_785_v23, 5)
			(_nw146).ArraySet1(_785_v23, 6)
			(_nw146).ArraySet1(_785_v23, 7)
			(_nw146).ArraySet1(_785_v23, 8)
			(_nw146).ArraySet1(_785_v23, 9)
			(_nw146).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_785_v23, _785_v23), (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_785_v23, _785_v23)).Cardinality()))).Uint32(), _784_v22), 10)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22), 11)
			(_nw146).ArraySet1(_785_v23, 12)
			(_nw146).ArraySet1(_785_v23, 13)
			(_nw146).ArraySet1(_785_v23, 14)
			(_nw146).ArraySet1(_785_v23, 15)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22, _784_v22), 16)
			(_nw146).ArraySet1(_785_v23, 17)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22), 18)
			(_nw146).ArraySet1(_785_v23, 19)
			(_nw146).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_785_v23, _785_v23), 20)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22), 21)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22, _784_v22), 22)
			(_nw146).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_785_v23, _785_v23), 23)
			(_nw146).ArraySet1(_785_v23, 24)
			(_nw146).ArraySet1(_dafny.SeqOf(Companion_D3_.Create_DC9_(_783_v21)), 25)
			(_nw146).ArraySet1(_dafny.SeqOf(_784_v22, Companion_D3_.Create_DC9_(_783_v21), _784_v22), 26)
			(_nw146).ArraySet1(_785_v23, 27)
			(_nw146).ArraySet1(_785_v23, 28)
			_786_v24 = _nw146
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_786_v24), 0))
			_ = _index141
			(_786_v24).ArraySet1(_785_v23, (_index141).Int())
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_786_v24), 0))
			_ = _index142
			(_786_v24).ArraySet1((func() _dafny.Sequence {
				if _this.F4() {
					return _785_v23
				}
				return _dafny.Companion_Sequence_.Concatenate(_785_v23, _785_v23)
			})(), (_index142).Int())
			var _787_v25 _dafny.Map
			_ = _787_v25
			_787_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F7())
			var _788_v26 _dafny.Sequence
			_ = _788_v26
			_788_v26 = _dafny.SeqOf((_787_v25).Merge(_787_v25), _787_v25)
			var _789_v27 _dafny.Sequence
			_ = _789_v27
			_789_v27 = _dafny.SeqOf(false)
			var _790_v28 _dafny.Map
			_ = _790_v28
			_790_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), _this.F4())
			var _791_v30 _dafny.MultiSet
			_ = _791_v30
			_791_v30 = _dafny.MultiSetOf((_this).F12(), (_this).F12())
			var _792_v31 D0
			_ = _792_v31
			_792_v31 = Companion_D0_.Create_DC2_(_791_v30, (Companion_D0_.Create_DC2_(_791_v30, _this.F4())).Dtor_cf6())
			var _793_v32 T0
			_ = _793_v32
			var _nw147 *C2 = New_C2_()
			_ = _nw147
			_nw147.Ctor__(_dafny.IntOfInt64(704), _752_v6, (_this).F7(), _this.F4())
			_793_v32 = _nw147
			var _794_v33 _dafny.Set
			_ = _794_v33
			_794_v33 = _dafny.SetOf(false)
			var _795_v34 D8
			_ = _795_v34
			_795_v34 = Companion_D8_.Create_DC23_(_this.F4(), _793_v32.F4(), _793_v32.F4(), (_this).F12())
			var _796_v35 _dafny.Map
			_ = _796_v35
			_796_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_793_v32.F4(), _this.F4())
			var _797_v36 _dafny.Array
			_ = _797_v36
			var _nwElement0_29 bool = _this.F4()
			_ = _nwElement0_29
			var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(22))
			_ = _nw148
			(_nw148).ArraySet1(_nwElement0_29, 0)
			(_nw148).ArraySet1(_this.F4(), 1)
			(_nw148).ArraySet1(_this.F4(), 2)
			(_nw148).ArraySet1(_this.F4(), 3)
			(_nw148).ArraySet1((_this.F4()) && (true), 4)
			(_nw148).ArraySet1(_this.F4(), 5)
			(_nw148).ArraySet1((_789_v27).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_789_v27).Cardinality()))).Uint32()).(bool), 6)
			(_nw148).ArraySet1(!(_790_v28).Contains((_this).F7()), 7)
			(_nw148).ArraySet1(_this.F4(), 8)
			(_nw148).ArraySet1(_this.F4(), 9)
			(_nw148).ArraySet1(_this.F4(), 10)
			(_nw148).ArraySet1(((Companion_Default___.Fm14(_this.F4(), _this.F4(), globalState)).Cardinality()).Cmp(_dafny.IntOfInt64(-524)) == 0, 11)
			(_nw148).ArraySet1(_this.F4(), 12)
			(_nw148).ArraySet1((_dafny.IntOfInt64(-900)).Cmp(_dafny.IntOfInt64(-434)) >= 0, 13)
			(_nw148).ArraySet1((Companion_Default___.Fm18(_792_v31, (Companion_D1_.Create_DC6_(_793_v32, Companion_Default___.Fm20((_794_v33).Cardinality(), globalState), _this.F4(), _793_v32.F4())).Dtor_cf11(), _793_v32.F4(), globalState)).IsSubsetOf(func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter13 := _dafny.Iterate((_this.F13).Elements()); ; {
					_compr_11, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _798_v29 _dafny.CodePoint
					_798_v29 = interface{}(_compr_11).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_this.F13, _798_v29) {
						_coll11.Add(_798_v29)
					}
				}
				return _coll11.ToSet()
			}()), 14)
			(_nw148).ArraySet1(_793_v32.F4(), 15)
			(_nw148).ArraySet1(_793_v32.F4(), 16)
			(_nw148).ArraySet1((_795_v34).Dtor_cf32(), 17)
			(_nw148).ArraySet1(_793_v32.F4(), 18)
			(_nw148).ArraySet1(_this.F4(), 19)
			(_nw148).ArraySet1((func() bool {
				if (_796_v35).Contains(_this.F4()) {
					return (_796_v35).Get(_this.F4()).(bool)
				}
				return _this.F4()
			})(), 20)
			(_nw148).ArraySet1(_this.F4(), 21)
			_797_v36 = _nw148
			var _799_v37 _dafny.Map
			_ = _799_v37
			_799_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F13)
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_797_v36), 0))
			_ = _index143
			(_797_v36).ArraySet1(((_799_v37).Cardinality()).Cmp((_794_v33).Cardinality()) >= 0, (_index143).Int())
			var _800_v38 _dafny.Sequence
			_ = _800_v38
			_800_v38 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kmrehf")).Cardinality()))
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_797_v36), 0))
			_ = _index144
			var _rhs144 _dafny.Int = ((func() _dafny.Int {
				if (_795_v34).Dtor_cf32() {
					return (_this).F7()
				}
				return _dafny.IntOfInt64(900)
			})()).Times((Companion_Default___.Fm27(globalState)).Cardinality())
			_ = _rhs144
			var _rhs145 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_788_v26, (Companion_Default___.SafeIndex((_793_v32).F3(), _dafny.IntOfUint32((_788_v26).Cardinality()))).Uint32(), _787_v25)
			_ = _rhs145
			var _rhs146 bool = _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm15(_this.F4(), _dafny.IntOfInt64(318), (_dafny.Zero).Minus(_dafny.IntOfUint32((_800_v38).Cardinality())), globalState), (_this.F13).Select((Companion_Default___.SafeIndex((_793_v32).F3(), _dafny.IntOfUint32((_this.F13).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_ = _rhs146
			var _lhs135 *GlobalState = globalState
			_ = _lhs135
			var _lhs136 _dafny.Array = _797_v36
			_ = _lhs136
			var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_797_v36), 0))
			_ = _lhs137
			_lhs135.F1 = _rhs144
			_788_v26 = _rhs145
			(_lhs136).ArraySet1(_rhs146, (_lhs137).Int())
			(_this).F4_set_((_this).Fm13(globalState))
			var _801_v39 _dafny.CodePoint
			_ = _801_v39
			_801_v39 = _dafny.CodePoint('l')
			_801_v39 = (_this).F12()
		} else if _source10.Is_DC0() {
			var _802___mcc_h6 bool = _source10.Get_().(D0_DC0).Cf0
			_ = _802___mcc_h6
			var _803_cf0 bool = _802___mcc_h6
			_ = _803_cf0
			var _804_v40 _dafny.Array
			_ = _804_v40
			var _nw149 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw149
			_804_v40 = _nw149
			var _805_v41 _dafny.Map
			_ = _805_v41
			_805_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_804_v40, _dafny.IntOfUint32((_dafny.SeqOf((_this).F7())).Cardinality()))
			_805_v41 = (_805_v41).Update(_804_v40, (_this).F3())
			if (_this).Fm13(globalState) {
				var _806_v42 D3
				_ = _806_v42
				_806_v42 = Companion_D3_.Create_DC10_((_this).F3(), _804_v40, ((_this).F7()).Times((_this).F3()))
				_806_v42 = _806_v42
				var _807_v43 _dafny.Sequence
				_ = _807_v43
				_807_v43 = _dafny.SeqOf((Companion_D0_.Create_DC0_(_803_cf0)).Dtor_cf0())
				var _808_v44 _dafny.Map
				_ = _808_v44
				_808_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), Companion_Default___.Fm2(_dafny.IntOfUint32((_807_v43).Cardinality()), globalState))
				var _809_v45 *C2
				_ = _809_v45
				var _nw150 *C2 = New_C2_()
				_ = _nw150
				_nw150.Ctor__((_dafny.Zero).Minus((_this).Fm12(_803_cf0, true, globalState)), _752_v6, ((_808_v44).Merge(_808_v44)).Cardinality(), _this.F4())
				_809_v45 = _nw150
				(globalState).F2 = (!(_803_cf0)) || (_803_cf0)
				var _810_v46 _dafny.Map
				_ = _810_v46
				_810_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _803_cf0)
				var _811_v47 D6
				_ = _811_v47
				_811_v47 = Companion_D6_.Create_DC18_((func() bool {
					if (_810_v46).Contains(_this.F4()) {
						return (_810_v46).Get(_this.F4()).(bool)
					}
					return _803_cf0
				})(), (_this).F7())
				var _812_v48 _dafny.Map
				_ = _812_v48
				_812_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v45).F14(), Companion_Default___.Fm2((_811_v47).Dtor_cf25(), globalState))
				var _813_v49 _dafny.Sequence
				_ = _813_v49
				_813_v49 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()))
				var _rhs147 _dafny.Map = (func() _dafny.Map {
					if !_dafny.Companion_Sequence_.Equal(_813_v49, _813_v49) {
						return _812_v48
					}
					return _812_v48
				})()
				_ = _rhs147
				var _rhs148 _dafny.Int = (_809_v45).F14()
				_ = _rhs148
				var _rhs149 _dafny.Int = Companion_Default___.SafeDivisionInt((_809_v45).F14(), (_this).F7())
				_ = _rhs149
				var _rhs150 _dafny.Int = _dafny.IntOfInt64(128)
				_ = _rhs150
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				var _lhs140 *GlobalState = globalState
				_ = _lhs140
				_812_v48 = _rhs147
				_lhs138.F1 = _rhs148
				_lhs139.F1 = _rhs149
				_lhs140.F1 = _rhs150
				(_this).F13 = _dafny.Companion_Sequence_.Concatenate(_this.F13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(659))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}(func(_814_i4 _dafny.Int) _dafny.CodePoint {
					return (_this).F12()
				})))
			} else {
				var _815_v50 _dafny.Sequence
				_ = _815_v50
				_815_v50 = _dafny.SeqOf(!(_803_cf0), _803_cf0)
				(globalState).F2 = (_815_v50).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F13).Cardinality()), _dafny.IntOfUint32((_815_v50).Cardinality()))).Uint32()).(bool)
				var _816_v51 _dafny.Set
				_ = _816_v51
				_816_v51 = _dafny.SetOf(true)
				var _817_v52 D9
				_ = _817_v52
				_817_v52 = Companion_D9_.Create_DC25_(_816_v51)
				var _rhs151 _dafny.Set = ((_817_v52).Dtor_cf37()).Difference(_816_v51)
				_ = _rhs151
				var _rhs152 bool = ((_this).Fm13(globalState)) || (_this.F4())
				_ = _rhs152
				var _lhs141 *GlobalState = globalState
				_ = _lhs141
				_816_v51 = _rhs151
				_lhs141.F2 = _rhs152
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_804_v40), 0))
				_ = _index145
				(_804_v40).ArraySet1(!(_803_cf0), (_index145).Int())
				var _818_v53 _dafny.MultiSet
				_ = _818_v53
				_818_v53 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gfobxgt")).Cardinality()))
				var _819_v54 _dafny.MultiSet
				_ = _819_v54
				_819_v54 = _dafny.MultiSetOf(Companion_Default___.Fm15(false, (_dafny.Zero).Minus((_this).F3()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vbacq")).Cardinality()), globalState), _this.F13)
				var _820_v55 _dafny.Set
				_ = _820_v55
				_820_v55 = _dafny.SetOf((_this).F12(), (_this).F12(), (_this).F12())
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_804_v40), 0))
				_ = _index146
				(_804_v40).ArraySet1(!(!((_818_v53).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.SeqOf((_819_v54).Cardinality(), (_820_v55).Cardinality()))))) || (_this.F4()), (_index146).Int())
				(globalState).F1 = (_this).F3()
				(globalState).F2 = _this.F4()
			}
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_804_v40), 0))
			_ = _index147
			(_804_v40).ArraySet1((Companion_Default___.Fm0(globalState)).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mibpnhp")).Cardinality())) > 0, (_index147).Int())
			var _821_v56 _dafny.MultiSet
			_ = _821_v56
			_821_v56 = _dafny.MultiSetOf((_this).F3(), (_this).F3(), (_this).F7(), (_this).F7())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_804_v40), 0))
			_ = _index148
			(_804_v40).ArraySet1((_821_v56).IsDisjointFrom(_821_v56), (_index148).Int())
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_752_v6), 0))
			_ = _index149
			(_752_v6).ArraySet1(Companion_Default___.SafeModuloInt((_this).F3(), (_this).F7()), (_index149).Int())
			var _822_v57 _dafny.Map
			_ = _822_v57
			_822_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _821_v56)
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_752_v6), 0))
			_ = _index150
			(_752_v6).ArraySet1(((_822_v57).Merge(_822_v57)).Cardinality(), (_index150).Int())
		} else {
			var _823___mcc_h7 D0 = _source10.Get_().(D0_DC4).Cf7
			_ = _823___mcc_h7
			var _824_cf7 D0 = _823___mcc_h7
			_ = _824_cf7
			var _825_v58 _dafny.Map
			_ = _825_v58
			_825_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), !(_dafny.Companion_Sequence_.Contains(_this.F13, (_this).F12())))
			var _826_v59 D9
			_ = _826_v59
			_826_v59 = Companion_D9_.Create_DC26_((_this).F3(), true)
			_825_v58 = (_825_v58).Update((_this).F7(), (_826_v59).Dtor_cf39())
			(globalState).F2 = _this.F4()
			var _827_v60 D8
			_ = _827_v60
			_827_v60 = Companion_D8_.Create_DC23_(_this.F4(), false, true, (_this).F12())
			var _828_v61 *C2
			_ = _828_v61
			var _nw151 *C2 = New_C2_()
			_ = _nw151
			_nw151.Ctor__((_this).F3(), _752_v6, ((_this).F3()).Minus((_this).F7()), (_827_v60).Dtor_cf33())
			_828_v61 = _nw151
			var _829_v62 _dafny.Set
			_ = _829_v62
			_829_v62 = _dafny.SetOf(false, _this.F4(), _this.F4())
			var _830_v63 _dafny.Array
			_ = _830_v63
			var _nwElement0_30 bool = _this.F4()
			_ = _nwElement0_30
			var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(22))
			_ = _nw152
			(_nw152).ArraySet1(_nwElement0_30, 0)
			(_nw152).ArraySet1(_this.F4(), 1)
			(_nw152).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_828_v61).F14()), (_this).F3()), 2)
			(_nw152).ArraySet1(_this.F4(), 3)
			(_nw152).ArraySet1(!(_this.F4()), 4)
			(_nw152).ArraySet1((_this.F4()) && (_this.F4()), 5)
			(_nw152).ArraySet1(((_828_v61).F14()).Cmp(_dafny.IntOfInt64(759)) > 0, 6)
			(_nw152).ArraySet1((func() bool {
				if false {
					return _this.F4()
				}
				return _this.F4()
			})(), 7)
			(_nw152).ArraySet1(_this.F4(), 8)
			(_nw152).ArraySet1(((_this).F7()).Cmp((_this).F3()) != 0, 9)
			(_nw152).ArraySet1(_this.F4(), 10)
			(_nw152).ArraySet1(_this.F4(), 11)
			(_nw152).ArraySet1(_this.F4(), 12)
			(_nw152).ArraySet1(!(false), 13)
			(_nw152).ArraySet1(_this.F4(), 14)
			(_nw152).ArraySet1(_this.F4(), 15)
			(_nw152).ArraySet1(_this.F4(), 16)
			(_nw152).ArraySet1((_dafny.SetOf(_this.F4())).IsDisjointFrom(_829_v62), 17)
			(_nw152).ArraySet1(_this.F4(), 18)
			(_nw152).ArraySet1(_this.F4(), 19)
			(_nw152).ArraySet1(_this.F4(), 20)
			(_nw152).ArraySet1((_this).Fm13(globalState), 21)
			_830_v63 = _nw152
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_830_v63), 0))
			_ = _index151
			(_830_v63).ArraySet1(_this.F4(), (_index151).Int())
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_830_v63), 0))
			_ = _index152
			(_830_v63).ArraySet1((func() bool {
				if (func() bool {
					if _this.F4() {
						return false
					}
					return !(_this.F4())
				})() {
					return !((!(true)) && (_this.F4()))
				}
				return _this.F4()
			})(), (_index152).Int())
		}
		var _831_v64 _dafny.CodePoint
		_ = _831_v64
		_831_v64 = _dafny.CodePoint('g')
		var _832_v65 D8
		_ = _832_v65
		_832_v65 = Companion_D8_.Create_DC23_(_this.F4(), _this.F4(), _this.F4(), _831_v64)
		_831_v64 = (_832_v65).Dtor_cf35()
		var _hi7 _dafny.Int = (_dafny.Zero).Minus((_this).F3())
		_ = _hi7
		for _833_i5 := (_this).F7(); _833_i5.Cmp(_hi7) < 0; _833_i5 = _833_i5.Plus(_dafny.One) {
			var _834_v66 _dafny.Array
			_ = _834_v66
			var _nw153 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(21))
			_ = _nw153
			_834_v66 = _nw153
			var _835_v67 _dafny.Map
			_ = _835_v67
			_835_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F4())
			var _836_v68 D6
			_ = _836_v68
			_836_v68 = Companion_D6_.Create_DC19_((_835_v67).Cardinality())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_834_v66), 0))
			_ = _index153
			(_834_v66).ArraySet1(_836_v68, (_index153).Int())
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_834_v66), 0))
			_ = _index154
			(_834_v66).ArraySet1(func(_pat_let15_0 D6) D6 {
				return func(_837_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let16_0 _dafny.Int) D6 {
						return func(_838_dt__update_hcf26_h0 _dafny.Int) D6 {
							return Companion_D6_.Create_DC19_(_838_dt__update_hcf26_h0)
						}(_pat_let16_0)
					}(_833_i5)
				}(_pat_let15_0)
			}(_836_v68), (_index154).Int())
			var _839_v69 _dafny.Map
			_ = _839_v69
			_839_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm28(!(_this.F4()), (_this).F3(), _this.F4(), globalState), (_this).F3())
			var _840_v70 _dafny.Sequence
			_ = _840_v70
			_840_v70 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F7()))
			var _841_v71 _dafny.Map
			_ = _841_v71
			_841_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_840_v70, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-752))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_842_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_843_i6 _dafny.Int) _dafny.CodePoint {
					return _842_v64
				}
			})(_831_v64))))
			_839_v69 = (_839_v69).Update((Companion_D10_.Create_DC28_(_841_v71)).Dtor_cf41(), _833_i5)
			(globalState).F2 = !((_833_i5).Cmp((_this).F7()) >= 0) || (_this.F4())
			var _844_v72 _dafny.Array
			_ = _844_v72
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_27
			var _nw154 _dafny.Array
			_ = _nw154
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw154 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) bool = func(_845_i7 _dafny.Int) bool {
					return _this.F4()
				}
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
			_844_v72 = _nw154
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_844_v72), 0))
			_ = _index155
			(_844_v72).ArraySet1(_this.F4(), (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_844_v72), 0))
			_ = _index156
			(_844_v72).ArraySet1(_this.F4(), (_index156).Int())
		}
		var _846_v73 D0
		_ = _846_v73
		_846_v73 = Companion_D0_.Create_DC3_()
		var _source11 D0 = _846_v73
		_ = _source11
		if _source11.Is_DC1() {
			var _847___mcc_h8 bool = _source11.Get_().(D0_DC1).Cf1
			_ = _847___mcc_h8
			var _848___mcc_h9 bool = _source11.Get_().(D0_DC1).Cf2
			_ = _848___mcc_h9
			var _849___mcc_h10 _dafny.CodePoint = _source11.Get_().(D0_DC1).Cf3
			_ = _849___mcc_h10
			var _850___mcc_h11 _dafny.Array = _source11.Get_().(D0_DC1).Cf4
			_ = _850___mcc_h11
			var _851_cf4 _dafny.Array = _850___mcc_h11
			_ = _851_cf4
			var _852_cf3 _dafny.CodePoint = _849___mcc_h10
			_ = _852_cf3
			var _853_cf2 bool = _848___mcc_h9
			_ = _853_cf2
			var _854_cf1 bool = _847___mcc_h8
			_ = _854_cf1
			(globalState).F1 = (func() _dafny.Int {
				if false {
					return (_this).F7()
				}
				return _dafny.IntOfInt64(338)
			})()
			(_this).F4_set_(_854_cf1)
			var _855_v74 _dafny.Map
			_ = _855_v74
			_855_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_852_cf3, Companion_D9_.Create_DC26_((_this).F7(), false))
			var _856_v75 _dafny.Set
			_ = _856_v75
			_856_v75 = _dafny.SetOf((_this).F7(), (_855_v74).Cardinality())
			(_this).F13 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ymicw"), _this.F13), (Companion_Default___.SafeIndex((_856_v75).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ymicw"), _this.F13)).Cardinality()))).Uint32(), _831_v64)
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_752_v6), 0))
			_ = _index157
			(_752_v6).ArraySet1((func() _dafny.Int {
				if _853_cf2 {
					return (_this).F3()
				}
				return _dafny.IntOfInt64(476)
			})(), (_index157).Int())
			var _857_v76 _dafny.Array
			_ = _857_v76
			var _len0_28 _dafny.Int = _dafny.One
			_ = _len0_28
			var _nw155 _dafny.Array
			_ = _nw155
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw155 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) bool = (func(_858_cf1 bool) func(_dafny.Int) bool {
					return func(_859_i8 _dafny.Int) bool {
						return _858_cf1
					}
				})(_854_cf1)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw155 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw155).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw155).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_857_v76 = _nw155
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_857_v76), 0))
			_ = _index158
			(_857_v76).ArraySet1(_this.F4(), (_index158).Int())
			var _860_v77 _dafny.Sequence
			_ = _860_v77
			_860_v77 = _dafny.SeqOf(!(true))
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_752_v6), 0))
			_ = _index159
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_857_v76), 0))
			_ = _index160
			var _rhs153 bool = !((_860_v77).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F3()), _dafny.IntOfUint32((_860_v77).Cardinality()))).Uint32()).(bool)) || (_853_cf2)
			_ = _rhs153
			var _rhs154 _dafny.Int = _dafny.IntOfInt64(52)
			_ = _rhs154
			var _rhs155 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dcb"), _this.F13), _this.F13)
			_ = _rhs155
			var _rhs156 bool = !((_860_v77).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("athho")).Cardinality()), (_this).F3()), _dafny.IntOfUint32((_860_v77).Cardinality()))).Uint32()).(bool))
			_ = _rhs156
			var _rhs157 bool = _853_cf2
			_ = _rhs157
			var _lhs142 *C3 = _this
			_ = _lhs142
			var _lhs143 _dafny.Array = _752_v6
			_ = _lhs143
			var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_752_v6), 0))
			_ = _lhs144
			var _lhs145 *C3 = _this
			_ = _lhs145
			var _lhs146 _dafny.Array = _857_v76
			_ = _lhs146
			var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_857_v76), 0))
			_ = _lhs147
			var _lhs148 *GlobalState = globalState
			_ = _lhs148
			_lhs142.F4_set_(_rhs153)
			(_lhs143).ArraySet1(_rhs154, (_lhs144).Int())
			_lhs145.F4_set_(_rhs155)
			(_lhs146).ArraySet1(_rhs156, (_lhs147).Int())
			_lhs148.F2 = _rhs157
		} else if _source11.Is_DC2() {
			var _861___mcc_h12 _dafny.MultiSet = _source11.Get_().(D0_DC2).Cf5
			_ = _861___mcc_h12
			var _862___mcc_h13 bool = _source11.Get_().(D0_DC2).Cf6
			_ = _862___mcc_h13
			var _863_cf6 bool = _862___mcc_h13
			_ = _863_cf6
			var _864_cf5 _dafny.MultiSet = _861___mcc_h12
			_ = _864_cf5
			(globalState).F1 = Companion_Default___.SafeDivisionInt((_this).F3(), (_this).F3())
			(_this).F4_set_((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F13, _this.F13)).Cardinality())).Cmp((_this).F3()) != 0)
			var _865_v78 _dafny.Array
			_ = _865_v78
			var _nw156 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw156
			_865_v78 = _nw156
			var _866_v79 *C0
			_ = _866_v79
			var _nw157 *C0 = New_C0_()
			_ = _nw157
			_nw157.Ctor__(!(_863_cf6), _865_v78)
			_866_v79 = _nw157
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_865_v78), 0))
			_ = _index161
			(_865_v78).ArraySet1(_863_cf6, (_index161).Int())
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_865_v78), 0))
			_ = _index162
			(_865_v78).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(728))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_867_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_868_i9 _dafny.Int) _dafny.CodePoint {
					return _867_v64
				}
			})(_831_v64))), _this.F13), (_index162).Int())
		} else if _source11.Is_DC3() {
			(globalState).F1 = (_dafny.Zero).Minus((_this).F3())
			var _869_v80 _dafny.MultiSet
			_ = _869_v80
			_869_v80 = _dafny.MultiSetOf((_this).F7(), (_this).F7(), (_this).F3(), _dafny.IntOfInt64(551))
			var _870_v81 _dafny.MultiSet
			_ = _870_v81
			_870_v81 = _dafny.MultiSetOf((_869_v80).Cardinality())
			(globalState).F2 = (_870_v81).IsSubsetOf(_870_v81)
			var _871_v82 _dafny.Sequence
			_ = _871_v82
			_871_v82 = _dafny.SeqOf(_this.F4(), false, _this.F4(), _this.F4())
			var _872_v83 _dafny.MultiSet
			_ = _872_v83
			_872_v83 = _dafny.MultiSetOf((_871_v82).Select((Companion_Default___.SafeIndex((_this).F3(), _dafny.IntOfUint32((_871_v82).Cardinality()))).Uint32()).(bool), _this.F4())
			(globalState).F1 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(((_872_v83).Difference(_872_v83)).Cardinality()), (_this).F3())
			(globalState).F1 = (func() _dafny.Int {
				if (_dafny.IntOfInt64(561)).Cmp((_this).F7()) == 0 {
					return _dafny.IntOfInt64(279)
				}
				return (_this).F7()
			})()
		} else if _source11.Is_DC0() {
			var _873___mcc_h14 bool = _source11.Get_().(D0_DC0).Cf0
			_ = _873___mcc_h14
			var _874_cf0 bool = _873___mcc_h14
			_ = _874_cf0
			var _rhs158 bool = _874_cf0
			_ = _rhs158
			var _rhs159 bool = Companion_Default___.Fm2((_this).F3(), globalState)
			_ = _rhs159
			var _lhs149 *GlobalState = globalState
			_ = _lhs149
			var _lhs150 *C3 = _this
			_ = _lhs150
			_lhs149.F2 = _rhs158
			_lhs150.F4_set_(_rhs159)
			var _pat_let_tv10 = _831_v64
			_ = _pat_let_tv10
			var _pat_let_tv11 = _874_cf0
			_ = _pat_let_tv11
			var _875_v84 _dafny.Array
			_ = _875_v84
			var _nwElement0_31 D8 = _832_v65
			_ = _nwElement0_31
			var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(13))
			_ = _nw158
			(_nw158).ArraySet1(_nwElement0_31, 0)
			(_nw158).ArraySet1(Companion_D8_.Create_DC23_(_this.F4(), _this.F4(), true, (_this).F12()), 1)
			(_nw158).ArraySet1(_832_v65, 2)
			(_nw158).ArraySet1(Companion_D8_.Create_DC23_(_this.F4(), true, _this.F4(), _831_v64), 3)
			(_nw158).ArraySet1(_832_v65, 4)
			(_nw158).ArraySet1(_832_v65, 5)
			(_nw158).ArraySet1(Companion_D8_.Create_DC23_(_this.F4(), _874_cf0, _this.F4(), _831_v64), 6)
			(_nw158).ArraySet1(_832_v65, 7)
			(_nw158).ArraySet1(Companion_D8_.Create_DC23_(_874_cf0, _874_cf0, _this.F4(), (_this).F12()), 8)
			(_nw158).ArraySet1(_832_v65, 9)
			(_nw158).ArraySet1(_832_v65, 10)
			(_nw158).ArraySet1(_832_v65, 11)
			(_nw158).ArraySet1(func(_pat_let17_0 D8) D8 {
				return func(_876_dt__update__tmp_h1 D8) D8 {
					return func(_pat_let18_0 _dafny.CodePoint) D8 {
						return func(_877_dt__update_hcf35_h0 _dafny.CodePoint) D8 {
							return func(_pat_let19_0 bool) D8 {
								return func(_878_dt__update_hcf32_h0 bool) D8 {
									return Companion_D8_.Create_DC23_(_878_dt__update_hcf32_h0, (_876_dt__update__tmp_h1).Dtor_cf33(), (_876_dt__update__tmp_h1).Dtor_cf34(), _877_dt__update_hcf35_h0)
								}(_pat_let19_0)
							}(_pat_let_tv11)
						}(_pat_let18_0)
					}(_pat_let_tv10)
				}(_pat_let17_0)
			}(_832_v65), 12)
			_875_v84 = _nw158
			var _879_v85 _dafny.Map
			_ = _879_v85
			_879_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _875_v84)
			_879_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(638), _875_v84)
			(_this).F4_set_(true)
			var _880_v86 _dafny.MultiSet
			_ = _880_v86
			_880_v86 = _dafny.MultiSetOf(_this.F4(), _this.F4(), _874_cf0, false, _this.F4())
			var _881_v87 _dafny.MultiSet
			_ = _881_v87
			_881_v87 = _dafny.MultiSetOf(_dafny.IntOfUint32((_this.F13).Cardinality()), (_this).F7(), Companion_Default___.SafeModuloInt((_this).F3(), (_880_v86).Cardinality()), (_dafny.Zero).Minus((_this).F7()))
			var _882_v88 _dafny.Sequence
			_ = _882_v88
			_882_v88 = _dafny.SeqOf((_this).F3(), (_this).F3(), (_this).F3())
			var _883_v89 D11
			_ = _883_v89
			_883_v89 = Companion_D11_.Create_DC31_(_882_v88)
			_881_v87 = _dafny.MultiSetFromSeq((_883_v89).Dtor_cf49())
		} else {
			var _884___mcc_h15 D0 = _source11.Get_().(D0_DC4).Cf7
			_ = _884___mcc_h15
			var _885_cf7 D0 = _884___mcc_h15
			_ = _885_cf7
			var _886_v91 _dafny.Set
			_ = _886_v91
			_886_v91 = _dafny.SetOf((_this).F3())
			(globalState).F2 = (_886_v91).IsSubsetOf(func() _dafny.Set {
				var _coll12 = _dafny.NewBuilder()
				_ = _coll12
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(819), _dafny.IntOfInt64(503))); ; {
					_compr_12, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _887_v90 _dafny.Int
					_887_v90 = interface{}(_compr_12).(_dafny.Int)
					if ((_dafny.IntOfInt64(819)).Cmp(_887_v90) <= 0) && ((_887_v90).Cmp(_dafny.IntOfInt64(503)) < 0) {
						_coll12.Add(Companion_Default___.SafeModuloInt(_887_v90, (_this).F3()))
					}
				}
				return _coll12.ToSet()
			}())
			var _888_v92 _dafny.Map
			_ = _888_v92
			_888_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("upfnlrddl")).Cardinality()), _this.F4())
			var _889_v94 _dafny.Map
			_ = _889_v94
			_889_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (_886_v91).Cardinality())
			var _890_v97 _dafny.Sequence
			_ = _890_v97
			_890_v97 = _dafny.SeqOf(_dafny.IntOfInt64(44), (_this).F3())
			var _891_v98 _dafny.Sequence
			_ = _891_v98
			_891_v98 = _dafny.SeqOf(_this.F4())
			var _892_v100 _dafny.Array
			_ = _892_v100
			var _nwElement0_32 _dafny.Map = (_888_v92).Merge(_888_v92)
			_ = _nwElement0_32
			var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(28))
			_ = _nw159
			(_nw159).ArraySet1(_nwElement0_32, 0)
			(_nw159).ArraySet1((_888_v92).Update(Companion_Default___.Fm0(globalState), _this.F4()), 1)
			(_nw159).ArraySet1(_888_v92, 2)
			(_nw159).ArraySet1(_888_v92, 3)
			(_nw159).ArraySet1(((_888_v92).Update((_this).F3(), _this.F4())).Merge(_888_v92), 4)
			(_nw159).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _this.F4()), 5)
			(_nw159).ArraySet1(_888_v92, 6)
			(_nw159).ArraySet1((_888_v92).Update((_this).F7(), _this.F4()), 7)
			(_nw159).ArraySet1((_888_v92).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _this.F4())), 8)
			(_nw159).ArraySet1(_888_v92, 9)
			(_nw159).ArraySet1(_888_v92, 10)
			(_nw159).ArraySet1(_888_v92, 11)
			(_nw159).ArraySet1((_888_v92).Merge(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-833), _dafny.IntOfInt64(-986))); ; {
					_compr_13, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _893_v93 _dafny.Int
					_893_v93 = interface{}(_compr_13).(_dafny.Int)
					if ((_dafny.IntOfInt64(-833)).Cmp(_893_v93) <= 0) && ((_893_v93).Cmp(_dafny.IntOfInt64(-986)) < 0) {
						_coll13.Add((_893_v93).Plus((_this).F7()), _this.F4())
					}
				}
				return _coll13.ToMap()
			}()), 12)
			(_nw159).ArraySet1(_888_v92, 13)
			(_nw159).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_889_v94).Contains(_831_v64) {
					return (_889_v94).Get(_831_v64).(_dafny.Int)
				}
				return (_this).F7()
			})(), _this.F4()), 14)
			(_nw159).ArraySet1(_888_v92, 15)
			(_nw159).ArraySet1((func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(626), _dafny.IntOfInt64(259))); ; {
					_compr_14, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _894_v95 _dafny.Int
					_894_v95 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(626)).Cmp(_894_v95) <= 0) && ((_894_v95).Cmp(_dafny.IntOfInt64(259)) < 0) {
						_coll14.Add((_894_v95).Minus((_this).F7()), _this.F4())
					}
				}
				return _coll14.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(607), _this.F4())), 16)
			(_nw159).ArraySet1(_888_v92, 17)
			(_nw159).ArraySet1(_888_v92, 18)
			(_nw159).ArraySet1(_888_v92, 19)
			(_nw159).ArraySet1(_888_v92, 20)
			(_nw159).ArraySet1(func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter17 := _dafny.Iterate((_890_v97).Elements()); ; {
					_compr_15, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _895_v96 _dafny.Int
					_895_v96 = interface{}(_compr_15).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_890_v97, _895_v96) {
						_coll15.Add(Companion_Default___.SafeModuloInt(_895_v96, (_this).F3()), _this.F4())
					}
				}
				return _coll15.ToMap()
			}(), 21)
			(_nw159).ArraySet1((_888_v92).Merge(_888_v92), 22)
			(_nw159).ArraySet1((_888_v92).Merge(_888_v92), 23)
			(_nw159).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), !(_this.F4())), 24)
			(_nw159).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _this.F4())).Update((_this).F3(), false)).Merge(_888_v92), 25)
			(_nw159).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(211), (_891_v98).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_891_v98).Cardinality()))).Uint32()).(bool)), 26)
			(_nw159).ArraySet1((_888_v92).Merge(func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(937), _dafny.IntOfInt64(258))); ; {
					_compr_16, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _896_v99 _dafny.Int
					_896_v99 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(937)).Cmp(_896_v99) <= 0) && ((_896_v99).Cmp(_dafny.IntOfInt64(258)) < 0) {
						_coll16.Add((_896_v99).Minus((_this).F7()), _this.F4())
					}
				}
				return _coll16.ToMap()
			}()), 27)
			_892_v100 = _nw159
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_892_v100), 0))
			_ = _index163
			(_892_v100).ArraySet1(_888_v92, (_index163).Int())
			var _897_v101 _dafny.MultiSet
			_ = _897_v101
			_897_v101 = _dafny.MultiSetOf(_this.F4(), _this.F4())
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_892_v100), 0))
			_ = _index164
			var _rhs160 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_dafny.IntOfUint32((_890_v97).Cardinality())).Cmp((_this).F7()) > 0)
			_ = _rhs160
			var _rhs161 bool = (_897_v101).IsProperSubsetOf(_897_v101)
			_ = _rhs161
			var _lhs151 _dafny.Array = _892_v100
			_ = _lhs151
			var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_892_v100), 0))
			_ = _lhs152
			var _lhs153 *GlobalState = globalState
			_ = _lhs153
			(_lhs151).ArraySet1(_rhs160, (_lhs152).Int())
			_lhs153.F2 = _rhs161
			(globalState).F2 = _this.F4()
			var _898_v102 *C1
			_ = _898_v102
			var _nw160 *C1 = New_C1_()
			_ = _nw160
			_nw160.Ctor__((_this).F7(), _dafny.IntOfInt64(885), Companion_Default___.Fm2((_this).F3(), globalState))
			_898_v102 = _nw160
		}
		(_this).F4_set_(true)
	}
}
func (_this *C3) F12() _dafny.CodePoint {
	{
		return _this._f12
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f4 bool
	_f7 _dafny.Int
	_f3 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f4 = false
	_this._f7 = _dafny.Zero
	_this._f3 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F4() bool {
	return _this._f4
}
func (_this *C4) F4_set_(value bool) {
	_this._f4 = value
}
func (_this *C4) F7() _dafny.Int {
	return _this._f7
}
func (_this *C4) F3() _dafny.Int {
	return _this._f3
}
func (_this *C4) Ctor__(f7 _dafny.Int, f3 _dafny.Int, f4 bool) {
	{
		(_this)._f7 = f7
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C4) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		(globalState).F2 = _this.F4()
		var _899_v0 _dafny.Array
		_ = _899_v0
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_29
		var _nw161 _dafny.Array
		_ = _nw161
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw161 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) bool = func(_900_i0 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(!(_this.F4())), _dafny.SeqOf(!(_this.F4()), _this.F4()))
			}
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw161 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw161).ArraySet1(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw161).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_899_v0 = _nw161
		var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))
		_ = _index165
		(_899_v0).ArraySet1(_this.F4(), (_index165).Int())
		var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))
		_ = _index166
		(_899_v0).ArraySet1(true, (_index166).Int())
		var _hi8 _dafny.Int = _dafny.IntOfInt64(972)
		_ = _hi8
		for _901_i1 := (_this).F7(); _901_i1.Cmp(_hi8) < 0; _901_i1 = _901_i1.Plus(_dafny.One) {
			var _902_v1 _dafny.Array
			_ = _902_v1
			var _nw162 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw162
			_902_v1 = _nw162
			var _903_v2 _dafny.Set
			_ = _903_v2
			_903_v2 = _dafny.SetOf(_902_v1)
			var _904_v3 _dafny.Array
			_ = _904_v3
			var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw163
			_904_v3 = _nw163
			var _905_v4 _dafny.Sequence
			_ = _905_v4
			_905_v4 = _dafny.SeqOf((_this).F3(), p0, (_this).F3(), _901_i1)
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_904_v3), 0))
			_ = _index167
			(_904_v3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_905_v4, _905_v4), (_index167).Int())
			var _906_v5 _dafny.Map
			_ = _906_v5
			_906_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState)), _903_v2)
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_904_v3), 0))
			_ = _index168
			var _rhs162 bool = _this.F4()
			_ = _rhs162
			var _rhs163 _dafny.Set = (func() _dafny.Set {
				if (_906_v5).Contains(p0) {
					return (_906_v5).Get(p0).(_dafny.Set)
				}
				return _903_v2
			})()
			_ = _rhs163
			var _rhs164 _dafny.Sequence = _905_v4
			_ = _rhs164
			var _lhs154 *GlobalState = globalState
			_ = _lhs154
			var _lhs155 _dafny.Array = _904_v3
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_904_v3), 0))
			_ = _lhs156
			_lhs154.F2 = _rhs162
			_903_v2 = _rhs163
			(_lhs155).ArraySet1(_rhs164, (_lhs156).Int())
			var _907_v6 _dafny.Sequence
			_ = _907_v6
			_907_v6 = _dafny.UnicodeSeqOfUtf8Bytes("g")
			(globalState).F1 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_907_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-923))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}(func(_908_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))).Cardinality()))
			var _909_v7 *C0
			_ = _909_v7
			var _nw164 *C0 = New_C0_()
			_ = _nw164
			_nw164.Ctor__(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("dmaqyb"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(486))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg39 _dafny.Int) interface{} {
					return coer39(arg39)
				}
			}(func(_910_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}))), _899_v0)
			_909_v7 = _nw164
			var _911_v8 _dafny.MultiSet
			_ = _911_v8
			_911_v8 = _dafny.MultiSetOf(_901_i1)
			var _912_v9 _dafny.CodePoint
			_ = _912_v9
			_912_v9 = _dafny.CodePoint('a')
			var _913_v10 _dafny.MultiSet
			_ = _913_v10
			_913_v10 = _dafny.MultiSetOf(_912_v9)
			var _914_v11 D0
			_ = _914_v11
			_914_v11 = Companion_D0_.Create_DC2_(_913_v10, _909_v7.F10)
			var _915_v12 _dafny.Map
			_ = _915_v12
			_915_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v11, (_899_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))).Int()).(bool))
			var _pat_let_tv12 = _913_v10
			_ = _pat_let_tv12
			var _pat_let_tv13 = _913_v10
			_ = _pat_let_tv13
			if (_911_v8).IsDisjointFrom(Companion_Default___.Fm6((func() bool {
				if (_915_v12).Contains(func(_pat_let20_0 D0) D0 {
					return func(_916_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let21_0 _dafny.MultiSet) D0 {
							return func(_917_dt__update_hcf5_h0 _dafny.MultiSet) D0 {
								return Companion_D0_.Create_DC2_(_917_dt__update_hcf5_h0, (_916_dt__update__tmp_h0).Dtor_cf6())
							}(_pat_let21_0)
						}(_pat_let_tv12)
					}(_pat_let20_0)
				}(Companion_D0_.Create_DC2_(_913_v10, true))) {
					return (_915_v12).Get(func(_pat_let22_0 D0) D0 {
						return func(_918_dt__update__tmp_h1 D0) D0 {
							return func(_pat_let23_0 _dafny.MultiSet) D0 {
								return func(_919_dt__update_hcf5_h1 _dafny.MultiSet) D0 {
									return Companion_D0_.Create_DC2_(_919_dt__update_hcf5_h1, (_918_dt__update__tmp_h1).Dtor_cf6())
								}(_pat_let23_0)
							}(_pat_let_tv13)
						}(_pat_let22_0)
					}(Companion_D0_.Create_DC2_(_913_v10, true))).(bool)
				}
				return _this.F4()
			})(), (_899_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))).Int()).(bool), globalState)) {
				var _920_v13 D0
				_ = _920_v13
				_920_v13 = Companion_D0_.Create_DC1_(!(_this.F4()), _909_v7.F10, _912_v9, _902_v1)
				var _921_v14 _dafny.Sequence
				_ = _921_v14
				_921_v14 = _dafny.SeqOf((_920_v13).Dtor_cf1())
				var _922_v15 _dafny.Map
				_ = _922_v15
				_922_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_921_v14, _907_v6)
				var _923_v16 _dafny.Set
				_ = _923_v16
				_923_v16 = _dafny.SetOf(_901_i1)
				var _924_v18 _dafny.Map
				_ = _924_v18
				_924_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_907_v6, (_this).F7())
				var _rhs165 _dafny.CodePoint = (func() _dafny.CodePoint {
					if !(_this.F4()) || (_909_v7.F10) {
						return _912_v9
					}
					return _912_v9
				})()
				_ = _rhs165
				var _rhs166 _dafny.Map = (_922_v15).Merge((_922_v15).Merge(_922_v15))
				_ = _rhs166
				var _rhs167 _dafny.Set = _923_v16
				_ = _rhs167
				var _rhs168 _dafny.Int = (Companion_Default___.Fm7(func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(792), _dafny.IntOfInt64(-411))); ; {
						_compr_17, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _925_v17 _dafny.Int
						_925_v17 = interface{}(_compr_17).(_dafny.Int)
						if ((_dafny.IntOfInt64(792)).Cmp(_925_v17) <= 0) && ((_925_v17).Cmp(_dafny.IntOfInt64(-411)) < 0) {
							_coll17.Add((_925_v17).Times((func() _dafny.Int {
								if (_911_v8).Contains((_this).F3()) {
									return (_911_v8).Multiplicity((_this).F3())
								}
								return (_this).F7()
							})()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_i1, _909_v7.F10)).Cardinality())
						}
					}
					return _coll17.ToMap()
				}(), (_this).F3(), _924_v18, globalState)).Cardinality()
				_ = _rhs168
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				_912_v9 = _rhs165
				_922_v15 = _rhs166
				_923_v16 = _rhs167
				_lhs157.F1 = _rhs168
				(globalState).F1 = Companion_Default___.SafeModuloInt((_this).F7(), p0)
				var _926_v19 _dafny.Map
				_ = _926_v19
				_926_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_909_v7.F10, (_this).F7())
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_902_v1), 0))
				_ = _index169
				(_902_v1).ArraySet1((_926_v19).Cardinality(), (_index169).Int())
				var _927_v20 D1
				_ = _927_v20
				_927_v20 = Companion_D1_.Create_DC7_()
				var _928_v21 _dafny.Map
				_ = _928_v21
				_928_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC7_(), _this.F4())
				var _929_v22 _dafny.MultiSet
				_ = _929_v22
				_929_v22 = _dafny.MultiSetOf(_this.F4(), _909_v7.F10, _909_v7.F10)
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_902_v1), 0))
				_ = _index170
				var _rhs169 bool = !(_909_v7.F10)
				_ = _rhs169
				var _rhs170 _dafny.Int = (_this).F7()
				_ = _rhs170
				var _rhs171 bool = (_928_v21).Contains(_927_v20)
				_ = _rhs171
				var _rhs172 _dafny.Int = ((_dafny.IntOfUint32((_907_v6).Cardinality())).Times((_929_v22).Cardinality())).Minus((_this).F7())
				_ = _rhs172
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				var _lhs161 _dafny.Array = _902_v1
				_ = _lhs161
				var _lhs162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_902_v1), 0))
				_ = _lhs162
				_lhs158.F2 = _rhs169
				_lhs159.F1 = _rhs170
				_lhs160.F2 = _rhs171
				(_lhs161).ArraySet1(_rhs172, (_lhs162).Int())
				var _930_v23 *C0
				_ = _930_v23
				var _nw165 *C0 = New_C0_()
				_ = _nw165
				_nw165.Ctor__((_this.F4()) || (_this.F4()), (_909_v7).F11())
				_930_v23 = _nw165
				var _931_v24 D0
				_ = _931_v24
				_931_v24 = Companion_D0_.Create_DC3_()
				_931_v24 = _931_v24
			} else {
				(globalState).F2 = (_899_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))).Int()).(bool)
				(globalState).F1 = (_this).F7()
				var _932_v25 _dafny.Map
				_ = _932_v25
				_932_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_909_v7.F10, !(_this.F4()))
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_902_v1), 0))
				_ = _index171
				(_902_v1).ArraySet1((_901_i1).Times((_this).F7()), (_index171).Int())
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_902_v1), 0))
				_ = _index172
				var _rhs173 bool = _909_v7.F10
				_ = _rhs173
				var _rhs174 _dafny.Map = _932_v25
				_ = _rhs174
				var _rhs175 _dafny.Int = (func() _dafny.Set {
					var _coll18 = _dafny.NewBuilder()
					_ = _coll18
					for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(149), _dafny.IntOfInt64(861))); ; {
						_compr_18, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _933_v26 _dafny.Int
						_933_v26 = interface{}(_compr_18).(_dafny.Int)
						if ((_dafny.IntOfInt64(149)).Cmp(_933_v26) <= 0) && ((_933_v26).Cmp(_dafny.IntOfInt64(861)) < 0) {
							_coll18.Add(Companion_Default___.SafeModuloInt(_933_v26, _901_i1))
						}
					}
					return _coll18.ToSet()
				}()).Cardinality()
				_ = _rhs175
				var _rhs176 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(p0, (_this).F3())).Times(Companion_Default___.Fm0(globalState)))
				_ = _rhs176
				var _lhs163 *C4 = _this
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				var _lhs165 _dafny.Array = _902_v1
				_ = _lhs165
				var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_902_v1), 0))
				_ = _lhs166
				_lhs163.F4_set_(_rhs173)
				_932_v25 = _rhs174
				_lhs164.F1 = _rhs175
				(_lhs165).ArraySet1(_rhs176, (_lhs166).Int())
				var _934_v27 D0
				_ = _934_v27
				_934_v27 = Companion_D0_.Create_DC0_(_909_v7.F10)
				_907_v6 = Companion_Default___.Fm8(false, Companion_D1_.Create_DC7_(), _934_v27, p0, globalState)
				(globalState).F2 = !(true)
			}
		}
		var _935_v28 _dafny.Sequence
		_ = _935_v28
		_935_v28 = _dafny.SeqOf(_dafny.IntOfInt64(231), (_dafny.Zero).Minus((_dafny.SetOf((_899_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))).Int()).(bool), _this.F4())).Cardinality()))
		r0 = (_935_v28).Select((Companion_Default___.SafeIndex((p0).Minus((_this).F3()), _dafny.IntOfUint32((_935_v28).Cardinality()))).Uint32()).(_dafny.Int)
		var _936_v29 _dafny.Array
		_ = _936_v29
		_936_v29 = _899_v0
		var _937_v30 _dafny.Sequence
		_ = _937_v30
		_937_v30 = _dafny.SeqOf((_936_v29))
		var _938_v31 _dafny.Array
		_ = _938_v31
		var _nwElement0_33 _dafny.Array = _899_v0
		_ = _nwElement0_33
		var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(19))
		_ = _nw166
		(_nw166).ArraySet1(_nwElement0_33, 0)
		(_nw166).ArraySet1((func() _dafny.Array {
			if (_899_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))).Int()).(bool) {
				return _899_v0
			}
			return _899_v0
		})(), 1)
		(_nw166).ArraySet1(_899_v0, 2)
		(_nw166).ArraySet1(_899_v0, 3)
		(_nw166).ArraySet1(_899_v0, 4)
		(_nw166).ArraySet1(_899_v0, 5)
		(_nw166).ArraySet1(_899_v0, 6)
		(_nw166).ArraySet1(_899_v0, 7)
		(_nw166).ArraySet1(_899_v0, 8)
		(_nw166).ArraySet1(_899_v0, 9)
		(_nw166).ArraySet1(_899_v0, 10)
		(_nw166).ArraySet1(_899_v0, 11)
		(_nw166).ArraySet1(_899_v0, 12)
		(_nw166).ArraySet1(_899_v0, 13)
		(_nw166).ArraySet1(_899_v0, 14)
		(_nw166).ArraySet1(_899_v0, 15)
		(_nw166).ArraySet1((_937_v30).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_937_v30).Cardinality()))).Uint32()).(_dafny.Array), 16)
		(_nw166).ArraySet1((func() _dafny.Array {
			if (_899_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))).Int()).(bool) {
				return _899_v0
			}
			return _899_v0
		})(), 17)
		(_nw166).ArraySet1(_899_v0, 18)
		_938_v31 = _nw166
		var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_938_v31), 0))
		_ = _index173
		(_938_v31).ArraySet1(_899_v0, (_index173).Int())
		var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_938_v31), 0))
		_ = _index174
		(_938_v31).ArraySet1(_899_v0, (_index174).Int())
		var _939_v32 *C0
		_ = _939_v32
		var _nw167 *C0 = New_C0_()
		_ = _nw167
		_nw167.Ctor__((_899_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_899_v0), 0))).Int()).(bool), (_937_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.IntOfUint32((_937_v30).Cardinality()))).Uint32()).(_dafny.Array))
		_939_v32 = _nw167
		var _940_v33 D0
		_ = _940_v33
		_940_v33 = Companion_D0_.Create_DC0_(_939_v32.F10)
		var _pat_let_tv14 = p0
		_ = _pat_let_tv14
		var _pat_let_tv15 = _899_v0
		_ = _pat_let_tv15
		var _pat_let_tv16 = _899_v0
		_ = _pat_let_tv16
		r0 = (_dafny.Zero).Minus(func(_source12 D0) _dafny.Int {
			if _source12.Is_DC1() {
				var _941___mcc_h0 bool = _source12.Get_().(D0_DC1).Cf1
				_ = _941___mcc_h0
				var _942___mcc_h1 bool = _source12.Get_().(D0_DC1).Cf2
				_ = _942___mcc_h1
				var _943___mcc_h2 _dafny.CodePoint = _source12.Get_().(D0_DC1).Cf3
				_ = _943___mcc_h2
				var _944___mcc_h3 _dafny.Array = _source12.Get_().(D0_DC1).Cf4
				_ = _944___mcc_h3
				var _945_cf4 _dafny.Array = _944___mcc_h3
				_ = _945_cf4
				var _946_cf3 _dafny.CodePoint = _943___mcc_h2
				_ = _946_cf3
				var _947_cf2 bool = _942___mcc_h1
				_ = _947_cf2
				var _948_cf1 bool = _941___mcc_h0
				_ = _948_cf1
				return (_pat_let_tv14).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_pat_let_tv15), 0))).Int()).(bool), true)).Cardinality())
			} else if _source12.Is_DC2() {
				var _949___mcc_h4 _dafny.MultiSet = _source12.Get_().(D0_DC2).Cf5
				_ = _949___mcc_h4
				var _950___mcc_h5 bool = _source12.Get_().(D0_DC2).Cf6
				_ = _950___mcc_h5
				var _951_cf6 bool = _950___mcc_h5
				_ = _951_cf6
				var _952_cf5 _dafny.MultiSet = _949___mcc_h4
				_ = _952_cf5
				return (_this).F7()
			} else if _source12.Is_DC3() {
				return _dafny.IntOfInt64(480)
			} else if _source12.Is_DC0() {
				var _953___mcc_h6 bool = _source12.Get_().(D0_DC0).Cf0
				_ = _953___mcc_h6
				var _954_cf0 bool = _953___mcc_h6
				_ = _954_cf0
				return (_this).F7()
			} else {
				var _955___mcc_h7 D0 = _source12.Get_().(D0_DC4).Cf7
				_ = _955___mcc_h7
				var _956_cf7 D0 = _955___mcc_h7
				_ = _956_cf7
				return (_this).F7()
			}
		}(_940_v33))
		return r0
	}
}
func (_this *C4) M4(p0 _dafny.Sequence, globalState *GlobalState) {
	{
		var _957_i0 _dafny.Int
		_ = _957_i0
		_957_i0 = _dafny.Zero
		{
			for _this.F4() {
				{
					if (_957_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_957_i0 = (_957_i0).Plus(_dafny.One)
					var _958_v0 _dafny.CodePoint
					_ = _958_v0
					_958_v0 = _dafny.CodePoint('w')
					var _959_v1 _dafny.Map
					_ = _959_v1
					_959_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F4()), _958_v0)
					var _960_v2 _dafny.Array
					_ = _960_v2
					var _nwElement0_34 _dafny.CodePoint = _958_v0
					_ = _nwElement0_34
					var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(6))
					_ = _nw168
					(_nw168).ArraySet1CodePoint(_nwElement0_34, 0)
					(_nw168).ArraySet1CodePoint(_958_v0, 1)
					(_nw168).ArraySet1CodePoint(_958_v0, 2)
					(_nw168).ArraySet1CodePoint((func() _dafny.CodePoint {
						if (_959_v1).Contains(Companion_Default___.Fm2(_dafny.IntOfInt64(866), globalState)) {
							return (_959_v1).Get(Companion_Default___.Fm2(_dafny.IntOfInt64(866), globalState)).(_dafny.CodePoint)
						}
						return _958_v0
					})(), 3)
					(_nw168).ArraySet1CodePoint(_958_v0, 4)
					(_nw168).ArraySet1CodePoint(_958_v0, 5)
					_960_v2 = _nw168
					var _961_v3 _dafny.Sequence
					_ = _961_v3
					_961_v3 = _dafny.SeqOf(_960_v2, _960_v2)
					var _962_v4 _dafny.Map
					_ = _962_v4
					_962_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), (_this).F7())
					var _963_v5 _dafny.MultiSet
					_ = _963_v5
					_963_v5 = _dafny.MultiSetOf(Companion_Default___.Fm0(globalState), (_this).F3(), (_this).F3(), (_this).F3())
					var _964_v6 D3
					_ = _964_v6
					_964_v6 = Companion_D3_.Create_DC9_(_960_v2)
					var _965_v7 _dafny.Map
					_ = _965_v7
					_965_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hynbt")).Cardinality()), _960_v2)
					var _966_v8 _dafny.Array
					_ = _966_v8
					var _nwElement0_35 _dafny.Array = _960_v2
					_ = _nwElement0_35
					var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(22))
					_ = _nw169
					(_nw169).ArraySet1(_nwElement0_35, 0)
					(_nw169).ArraySet1(_960_v2, 1)
					(_nw169).ArraySet1(_960_v2, 2)
					(_nw169).ArraySet1(_960_v2, 3)
					(_nw169).ArraySet1((func() _dafny.Array {
						if _this.F4() {
							return _960_v2
						}
						return _960_v2
					})(), 4)
					(_nw169).ArraySet1(_960_v2, 5)
					(_nw169).ArraySet1(_960_v2, 6)
					(_nw169).ArraySet1((_961_v3).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm9((p0).Select((Companion_Default___.SafeIndex((_962_v4).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint), _this.F4(), globalState)).Cardinality(), _dafny.IntOfUint32((_961_v3).Cardinality()))).Uint32()).(_dafny.Array), 7)
					(_nw169).ArraySet1((_961_v3).Select((Companion_Default___.SafeIndex((_963_v5).Cardinality(), _dafny.IntOfUint32((_961_v3).Cardinality()))).Uint32()).(_dafny.Array), 8)
					(_nw169).ArraySet1((_964_v6).Dtor_cf14(), 9)
					(_nw169).ArraySet1(_960_v2, 10)
					(_nw169).ArraySet1(_960_v2, 11)
					(_nw169).ArraySet1(_960_v2, 12)
					(_nw169).ArraySet1(_960_v2, 13)
					(_nw169).ArraySet1(_960_v2, 14)
					(_nw169).ArraySet1(_960_v2, 15)
					(_nw169).ArraySet1(_960_v2, 16)
					(_nw169).ArraySet1(_960_v2, 17)
					(_nw169).ArraySet1((func() _dafny.Array {
						if _this.F4() {
							return (func() _dafny.Array {
								if (_965_v7).Contains((_this).F3()) {
									return (_965_v7).Get((_this).F3()).(_dafny.Array)
								}
								return _960_v2
							})()
						}
						return _960_v2
					})(), 18)
					(_nw169).ArraySet1(_960_v2, 19)
					(_nw169).ArraySet1(_960_v2, 20)
					(_nw169).ArraySet1(_960_v2, 21)
					_966_v8 = _nw169
					_966_v8 = _966_v8
					(globalState).F1 = ((_this).F7()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F3()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm10(_this.F4(), _this.F4(), (_this).F3(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), p0), globalState)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_this).F3())).Cardinality()))).Uint32(), _dafny.IntOfInt64(730))).Cardinality()))
					(_this).F4_set_(!(_this.F4()))
					if !(true) {
						(globalState).F2 = false
						(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-964), (_this).F3()), (_this).F3()))
						(globalState).F2 = _this.F4()
						(_this).F4_set_(false)
						(globalState).F2 = true
					} else {
						var _967_v9 _dafny.Sequence
						_ = _967_v9
						_967_v9 = _dafny.UnicodeSeqOfUtf8Bytes("evx")
						_967_v9 = p0
						var _968_v10 _dafny.Array
						_ = _968_v10
						var _nw170 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
						_ = _nw170
						_968_v10 = _nw170
						var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_968_v10), 0))
						_ = _index175
						(_968_v10).ArraySet1(_this.F4(), (_index175).Int())
						var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_968_v10), 0))
						_ = _index176
						(_968_v10).ArraySet1(_this.F4(), (_index176).Int())
						var _969_v11 _dafny.Set
						_ = _969_v11
						_969_v11 = _dafny.SetOf((_this).F7())
						var _970_v13 _dafny.Sequence
						_ = _970_v13
						_970_v13 = _dafny.SeqOf((func() _dafny.Map {
							var _coll19 = _dafny.NewMapBuilder()
							_ = _coll19
							for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(309), _dafny.IntOfInt64(185))); ; {
								_compr_19, _ok21 := _iter21()
								if !_ok21 {
									break
								}
								var _971_v12 _dafny.Int
								_971_v12 = interface{}(_compr_19).(_dafny.Int)
								if ((_dafny.IntOfInt64(309)).Cmp(_971_v12) <= 0) && ((_971_v12).Cmp(_dafny.IntOfInt64(185)) < 0) {
									_coll19.Add((_971_v12).Plus(_dafny.IntOfInt64(514)), true)
								}
							}
							return _coll19.ToMap()
						}()).Cardinality(), (_this).F3())
						var _972_v14 _dafny.Array
						_ = _972_v14
						var _nwElement0_36 _dafny.Int = (_969_v11).Cardinality()
						_ = _nwElement0_36
						var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(16))
						_ = _nw171
						(_nw171).ArraySet1(_nwElement0_36, 0)
						(_nw171).ArraySet1(_dafny.IntOfInt64(613), 1)
						(_nw171).ArraySet1((_this).F3(), 2)
						(_nw171).ArraySet1((_this).F3(), 3)
						(_nw171).ArraySet1(Companion_Default___.Fm0(globalState), 4)
						(_nw171).ArraySet1((_this).F3(), 5)
						(_nw171).ArraySet1((_this).F3(), 6)
						(_nw171).ArraySet1(_dafny.IntOfInt64(373), 7)
						(_nw171).ArraySet1(_dafny.IntOfUint32((_970_v13).Cardinality()), 8)
						(_nw171).ArraySet1((_this).F7(), 9)
						(_nw171).ArraySet1((_969_v11).Cardinality(), 10)
						(_nw171).ArraySet1(_dafny.IntOfInt64(473), 11)
						(_nw171).ArraySet1((_this).F7(), 12)
						(_nw171).ArraySet1((_this).F7(), 13)
						(_nw171).ArraySet1((_this).F3(), 14)
						(_nw171).ArraySet1((_this).F7(), 15)
						_972_v14 = _nw171
						var _973_v15 _dafny.Array
						_ = _973_v15
						var _nwElement0_37 _dafny.Array = _972_v14
						_ = _nwElement0_37
						var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(18))
						_ = _nw172
						(_nw172).ArraySet1(_nwElement0_37, 0)
						(_nw172).ArraySet1(_972_v14, 1)
						(_nw172).ArraySet1(_972_v14, 2)
						(_nw172).ArraySet1(_972_v14, 3)
						(_nw172).ArraySet1(_972_v14, 4)
						(_nw172).ArraySet1(_972_v14, 5)
						(_nw172).ArraySet1(_972_v14, 6)
						(_nw172).ArraySet1(_972_v14, 7)
						(_nw172).ArraySet1(_972_v14, 8)
						(_nw172).ArraySet1(_972_v14, 9)
						(_nw172).ArraySet1(_972_v14, 10)
						(_nw172).ArraySet1(_972_v14, 11)
						(_nw172).ArraySet1(_972_v14, 12)
						(_nw172).ArraySet1(_972_v14, 13)
						(_nw172).ArraySet1(_972_v14, 14)
						(_nw172).ArraySet1(_972_v14, 15)
						(_nw172).ArraySet1(_972_v14, 16)
						(_nw172).ArraySet1(_972_v14, 17)
						_973_v15 = _nw172
						var _974_v16 _dafny.Map
						_ = _974_v16
						_974_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _973_v15)
						_974_v16 = (_974_v16).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()), _973_v15)
						var _975_v17 _dafny.Map
						_ = _975_v17
						_975_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _967_v9)
						(globalState).F1 = (func() _dafny.Int {
							if !_dafny.Companion_Sequence_.Equal(_970_v13, _dafny.SeqOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
								if (_975_v17).Contains(_this.F4()) {
									return (_975_v17).Get(_this.F4()).(_dafny.Sequence)
								}
								return _dafny.UnicodeSeqOfUtf8Bytes("eu")
							})()).Cardinality()))) {
								return (_dafny.Zero).Minus((_this).F3())
							}
							return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(122), (_dafny.Zero).Minus((_this).F3()))
						})()
						var _976_v18 *C0
						_ = _976_v18
						var _nw173 *C0 = New_C0_()
						_ = _nw173
						_nw173.Ctor__(_this.F4(), _968_v10)
						_976_v18 = _nw173
					}
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _977_v19 _dafny.CodePoint
		_ = _977_v19
		_977_v19 = _dafny.CodePoint('k')
		_977_v19 = _977_v19
		var _978_v20 _dafny.Map
		_ = _978_v20
		_978_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F3()).Times((_this).F3()), _this.F4())
		var _979_v21 _dafny.Set
		_ = _979_v21
		_979_v21 = _dafny.SetOf(_this.F4(), _this.F4())
		var _980_v22 _dafny.MultiSet
		_ = _980_v22
		_980_v22 = _dafny.MultiSetOf((_this).F3(), (_dafny.Zero).Minus((_979_v21).Cardinality()), (_this).F3(), (_this).F7(), (_dafny.Zero).Minus((_this).F3()))
		if (func() bool {
			if (_978_v20).Contains((_980_v22).Cardinality()) {
				return (_978_v20).Get((_980_v22).Cardinality()).(bool)
			}
			return _this.F4()
		})() {
			var _981_v23 _dafny.Array
			_ = _981_v23
			var _nw174 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw174
			_981_v23 = _nw174
			var _982_v24 *C0
			_ = _982_v24
			var _nw175 *C0 = New_C0_()
			_ = _nw175
			_nw175.Ctor__(_this.F4(), _981_v23)
			_982_v24 = _nw175
			var _983_v25 _dafny.Map
			_ = _983_v25
			_983_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), (_this).F7())
			_983_v25 = (_983_v25).Update((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState)), ((_979_v21).Union(_979_v21)).Cardinality())
			var _984_v26 _dafny.Map
			_ = _984_v26
			_984_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_982_v24.F10, p0)
			(globalState).F1 = (_dafny.IntOfUint32((Companion_Default___.Fm10(_982_v24.F10, _this.F4(), (_this).F7(), _984_v26, globalState)).Cardinality())).Minus((_this).F3())
			(_this).F4_set_(true)
			(globalState).F1 = (_this).F7()
		} else {
			var _985_v27 _dafny.Map
			_ = _985_v27
			_985_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F4())
			_985_v27 = (_985_v27).Update(_this.F4(), _this.F4())
			var _986_v28 _dafny.Array
			_ = _986_v28
			var _nwElement0_38 _dafny.Int = (_this).F3()
			_ = _nwElement0_38
			var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(2))
			_ = _nw176
			(_nw176).ArraySet1(_nwElement0_38, 0)
			(_nw176).ArraySet1((_this).F7(), 1)
			_986_v28 = _nw176
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_986_v28), 0))
			_ = _index177
			(_986_v28).ArraySet1(Companion_Default___.Fm0(globalState), (_index177).Int())
			var _987_v29 _dafny.Map
			_ = _987_v29
			_987_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7())
			var _988_v30 _dafny.Sequence
			_ = _988_v30
			_988_v30 = _dafny.SeqOf(_987_v29)
			var _989_v31 _dafny.Map
			_ = _989_v31
			_989_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F4())
			var _990_v32 _dafny.Array
			_ = _990_v32
			var _nwElement0_39 bool = _this.F4()
			_ = _nwElement0_39
			var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(10))
			_ = _nw177
			(_nw177).ArraySet1(_nwElement0_39, 0)
			(_nw177).ArraySet1(true, 1)
			(_nw177).ArraySet1(_this.F4(), 2)
			(_nw177).ArraySet1(_this.F4(), 3)
			(_nw177).ArraySet1(true, 4)
			(_nw177).ArraySet1(_this.F4(), 5)
			(_nw177).ArraySet1(_this.F4(), 6)
			(_nw177).ArraySet1((func() bool {
				if (_989_v31).Contains(p0) {
					return (_989_v31).Get(p0).(bool)
				}
				return _this.F4()
			})(), 7)
			(_nw177).ArraySet1(_this.F4(), 8)
			(_nw177).ArraySet1(_this.F4(), 9)
			_990_v32 = _nw177
			var _991_v33 *C0
			_ = _991_v33
			var _nw178 *C0 = New_C0_()
			_ = _nw178
			_nw178.Ctor__(_this.F4(), _990_v32)
			_991_v33 = _nw178
			var _992_v34 _dafny.Map
			_ = _992_v34
			_992_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _991_v33)
			var _993_v35 _dafny.Sequence
			_ = _993_v35
			_993_v35 = _dafny.SeqOf(_992_v34, _992_v34)
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_986_v28), 0))
			_ = _index178
			(_986_v28).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if !((_988_v30).Select((Companion_Default___.SafeIndex((_this).F3(), _dafny.IntOfUint32((_988_v30).Cardinality()))).Uint32()).(_dafny.Map)).Equals((Companion_Default___.Fm11((_this).F7(), globalState)).Update((_this).F3(), (_this).F7())) {
					return _993_v35
				}
				return _993_v35
			})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if _991_v33.F10 {
					return (_this).F3()
				}
				return (_this).F7()
			})()), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if !((_988_v30).Select((Companion_Default___.SafeIndex((_this).F3(), _dafny.IntOfUint32((_988_v30).Cardinality()))).Uint32()).(_dafny.Map)).Equals((Companion_Default___.Fm11((_this).F7(), globalState)).Update((_this).F3(), (_this).F7())) {
					return _993_v35
				}
				return _993_v35
			})()).Cardinality()))).Uint32(), _992_v34)).Cardinality()), (_index178).Int())
			var _994_v36 D11
			_ = _994_v36
			_994_v36 = Companion_D11_.Create_DC32_(_991_v33.F10, (_this).F3())
			var _995_v37 T1
			_ = _995_v37
			var _nw179 *C3 = New_C3_()
			_ = _nw179
			_nw179.Ctor__(_977_v19, p0, (_994_v36).Dtor_cf51(), _this.F4(), (_this).F7())
			_995_v37 = _nw179
			var _996_v38 _dafny.Array
			_ = _996_v38
			var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
			_ = _nw180
			_996_v38 = _nw180
			var _997_v39 _dafny.Sequence
			_ = _997_v39
			_997_v39 = _dafny.SeqOf((_this).F7(), (_995_v37).F7())
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_996_v38), 0))
			_ = _index179
			(_996_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_997_v39, _997_v39), (_index179).Int())
			var _998_v40 D1
			_ = _998_v40
			_998_v40 = Companion_D1_.Create_DC7_()
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_986_v28), 0))
			_ = _index180
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_996_v38), 0))
			_ = _index181
			var _rhs177 _dafny.Int = ((_this).F7()).Minus((_this).F7())
			_ = _rhs177
			var _rhs178 T1 = _995_v37
			_ = _rhs178
			var _rhs179 _dafny.Sequence = _997_v39
			_ = _rhs179
			var _rhs180 D1 = _998_v40
			_ = _rhs180
			var _lhs167 _dafny.Array = _986_v28
			_ = _lhs167
			var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_986_v28), 0))
			_ = _lhs168
			var _lhs169 _dafny.Array = _996_v38
			_ = _lhs169
			var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_996_v38), 0))
			_ = _lhs170
			(_lhs167).ArraySet1(_rhs177, (_lhs168).Int())
			_995_v37 = _rhs178
			(_lhs169).ArraySet1(_rhs179, (_lhs170).Int())
			_998_v40 = _rhs180
			var _999_v41 D7
			_ = _999_v41
			_999_v41 = Companion_D7_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_986_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_986_v28), 0))).Int()).(_dafny.Int), p0))
			var _1000_v42 _dafny.Set
			_ = _1000_v42
			_1000_v42 = _dafny.SetOf((_995_v37).F3())
			var _1001_v43 _dafny.Map
			_ = _1001_v43
			_1001_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1000_v42).Cardinality(), p0)
			var _pat_let_tv17 = _1001_v43
			_ = _pat_let_tv17
			_999_v41 = func(_pat_let24_0 D7) D7 {
				return func(_1002_dt__update__tmp_h0 D7) D7 {
					return func(_pat_let25_0 _dafny.Map) D7 {
						return func(_1003_dt__update_hcf27_h0 _dafny.Map) D7 {
							return Companion_D7_.Create_DC20_(_1003_dt__update_hcf27_h0)
						}(_pat_let25_0)
					}(_pat_let_tv17)
				}(_pat_let24_0)
			}(_999_v41)
			var _1004_v44 _dafny.Sequence
			_ = _1004_v44
			_1004_v44 = _dafny.SeqOf(p0, p0, p0)
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(p0, p0, p0, _dafny.UnicodeSeqOfUtf8Bytes("gjq")), _1004_v44) {
				var _1005_v45 _dafny.MultiSet
				_ = _1005_v45
				_1005_v45 = _dafny.MultiSetOf(_986_v28)
				var _1006_v46 _dafny.Map
				_ = _1006_v46
				_1006_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1005_v45, (_986_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_986_v28), 0))).Int()).(_dafny.Int))
				_1006_v46 = (_1006_v46).Update((_1005_v45).Union(_1005_v45), (_995_v37).F3())
				(_991_v33).F10 = _991_v33.F10
				(_991_v33).F10 = !(_991_v33.F10)
				var _1007_v47 _dafny.MultiSet
				_ = _1007_v47
				_1007_v47 = _dafny.MultiSetOf(_997_v39)
				(globalState).F1 = (_dafny.Zero).Minus((_1007_v47).Cardinality())
				(globalState).F1 = (_995_v37).F3()
			} else {
				(_991_v33).F10 = Companion_Default___.Fm2((_995_v37).F7(), globalState)
				var _1008_v48 _dafny.Sequence
				_ = _1008_v48
				_1008_v48 = _dafny.UnicodeSeqOfUtf8Bytes("gxc")
				var _rhs181 bool = !(_991_v33.F10)
				_ = _rhs181
				var _rhs182 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p0, _1008_v48), _dafny.Companion_Sequence_.Concatenate(_1008_v48, _1008_v48))
				_ = _rhs182
				var _lhs171 T1 = _995_v37
				_ = _lhs171
				_lhs171.F4_set_(_rhs181)
				_1008_v48 = _rhs182
				var _1009_v49 _dafny.MultiSet
				_ = _1009_v49
				_1009_v49 = _dafny.MultiSetOf(_985_v27, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F4()), true)).Update(_995_v37.F4(), _995_v37.F4()), _985_v27)
				var _1010_v50 _dafny.Map
				_ = _1010_v50
				_1010_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F3(), (_986_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_986_v28), 0))).Int()).(_dafny.Int)), _1009_v49)
				_1010_v50 = (_1010_v50).Update((_dafny.Zero).Minus((_995_v37).F3()), _1009_v49)
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen(((_991_v33).F11()), 0))
				_ = _index182
				((_991_v33).F11()).ArraySet1(_this.F4(), (_index182).Int())
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen(((_991_v33).F11()), 0))
				_ = _index183
				var _rhs183 bool = _995_v37.F4()
				_ = _rhs183
				var _rhs184 bool = ((_995_v37).F3()).Cmp((_995_v37).F3()) >= 0
				_ = _rhs184
				var _lhs172 *GlobalState = globalState
				_ = _lhs172
				var _lhs173 _dafny.Array = (_991_v33).F11()
				_ = _lhs173
				var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen(((_991_v33).F11()), 0))
				_ = _lhs174
				_lhs172.F2 = _rhs183
				(_lhs173).ArraySet1(_rhs184, (_lhs174).Int())
				var _1011_v51 _dafny.Sequence
				_ = _1011_v51
				_1011_v51 = _dafny.SeqOf((false) && (_this.F4()))
				_1011_v51 = _1011_v51
			}
		}
		(_this).F4_set_(_this.F4())
		(globalState).F1 = (_this).F7()
		var _1012_v52 _dafny.Array
		_ = _1012_v52
		var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
		_ = _nw181
		_1012_v52 = _nw181
		_1012_v52 = _1012_v52
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f4 bool
	_f7 _dafny.Int
	_f3 _dafny.Int
	_f9 _dafny.Array
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f4 = false
	_this._f7 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C5) F4() bool {
	return _this._f4
}
func (_this *C5) F4_set_(value bool) {
	_this._f4 = value
}
func (_this *C5) F7() _dafny.Int {
	return _this._f7
}
func (_this *C5) F3() _dafny.Int {
	return _this._f3
}
func (_this *C5) Ctor__(f9 _dafny.Array, f7 _dafny.Int, f3 _dafny.Int, f4 bool) {
	{
		(_this)._f9 = f9
		(_this)._f7 = f7
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C5) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1013_v0 _dafny.Array
		_ = _1013_v0
		var _nw182 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
		_ = _nw182
		_1013_v0 = _nw182
		var _1014_v1 _dafny.Array
		_ = _1014_v1
		var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw183
		_1014_v1 = _nw183
		var _1015_v2 _dafny.Sequence
		_ = _1015_v2
		_1015_v2 = _dafny.SeqOf(_1013_v0, _1013_v0)
		var _rhs185 _dafny.Array = (_1015_v2).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(p0, (_this).F3(), (_this).F3())).Cardinality(), _dafny.IntOfUint32((_1015_v2).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs185
		var _rhs186 _dafny.Array = _1014_v1
		_ = _rhs186
		var _rhs187 _dafny.Int = ((func() _dafny.Int {
			if _this.F4() {
				return p0
			}
			return (_this).F7()
		})()).Times(p0)
		_ = _rhs187
		var _lhs175 *GlobalState = globalState
		_ = _lhs175
		_1013_v0 = _rhs185
		_1014_v1 = _rhs186
		_lhs175.F1 = _rhs187
		var _hi9 _dafny.Int = ((_this).F3()).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ptajbtlxg")).Cardinality()))
		_ = _hi9
		for _1016_i0 := (_this).F3(); _1016_i0.Cmp(_hi9) < 0; _1016_i0 = _1016_i0.Plus(_dafny.One) {
			var _1017_v3 _dafny.Array
			_ = _1017_v3
			var _nw184 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw184
			_1017_v3 = _nw184
			_1017_v3 = _1017_v3
			var _1018_v4 _dafny.Sequence
			_ = _1018_v4
			_1018_v4 = _dafny.SeqOf(p0, _1016_i0, (_this).F3())
			var _1019_v5 _dafny.Set
			_ = _1019_v5
			_1019_v5 = _dafny.SetOf(_1018_v4, _1018_v4)
			var _1020_v6 _dafny.Map
			_ = _1020_v6
			_1020_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1016_i0, _dafny.SetOf(_this.F4(), !(false)))
			var _1021_v7 _dafny.Sequence
			_ = _1021_v7
			_1021_v7 = _dafny.UnicodeSeqOfUtf8Bytes("efq")
			var _1022_v8 _dafny.Array
			_ = _1022_v8
			var _nwElement0_40 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(606))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}((func(_1023_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1024_i1 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_1023_i0)
				}
			})(_1016_i0)))).Cardinality())
			_ = _nwElement0_40
			var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(17))
			_ = _nw185
			(_nw185).ArraySet1(_nwElement0_40, 0)
			(_nw185).ArraySet1(p0, 1)
			(_nw185).ArraySet1(p0, 2)
			(_nw185).ArraySet1((_1019_v5).Cardinality(), 3)
			(_nw185).ArraySet1(_dafny.IntOfInt64(-369), 4)
			(_nw185).ArraySet1((_this).F3(), 5)
			(_nw185).ArraySet1((_this).F3(), 6)
			(_nw185).ArraySet1(p0, 7)
			(_nw185).ArraySet1(p0, 8)
			(_nw185).ArraySet1(p0, 9)
			(_nw185).ArraySet1((_1020_v6).Cardinality(), 10)
			(_nw185).ArraySet1(_dafny.IntOfInt64(687), 11)
			(_nw185).ArraySet1((_this).F7(), 12)
			(_nw185).ArraySet1(_1016_i0, 13)
			(_nw185).ArraySet1((_this).F3(), 14)
			(_nw185).ArraySet1((_this).F3(), 15)
			(_nw185).ArraySet1(_dafny.IntOfUint32((_1021_v7).Cardinality()), 16)
			_1022_v8 = _nw185
			var _1025_v9 *C2
			_ = _1025_v9
			var _nw186 *C2 = New_C2_()
			_ = _nw186
			_nw186.Ctor__((_1016_i0).Plus(_1016_i0), _1022_v8, (_this).F7(), _this.F4())
			_1025_v9 = _nw186
			var _1026_v10 _dafny.Sequence
			_ = _1026_v10
			_1026_v10 = _dafny.SeqOf(_this.F4())
			var _1027_v11 _dafny.Map
			_ = _1027_v11
			_1027_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1026_v10).Cardinality()), _1021_v7)
			var _1028_v12 _dafny.Sequence
			_ = _1028_v12
			_1028_v12 = _dafny.SeqOf(_this.F4(), _this.F4(), _this.F4(), !(_1027_v11).Equals(_1027_v11))
			var _1029_v13 _dafny.Sequence
			_ = _1029_v13
			_1029_v13 = _dafny.SeqOf(_1028_v12)
			_1028_v12 = (_1029_v13).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_1029_v13).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _1030_v14 _dafny.Array
			_ = _1030_v14
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_30
			var _nw187 _dafny.Array
			_ = _nw187
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw187 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) D6 = func(_1031_i2 _dafny.Int) D6 {
					return Companion_D6_.Create_DC18_(false, (_this).F7())
				}
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw187 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw187).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw187).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1030_v14 = _nw187
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1030_v14), 0))
			_ = _index184
			(_1030_v14).ArraySet1(Companion_Default___.Fm29(globalState), (_index184).Int())
			var _1032_v15 D6
			_ = _1032_v15
			_1032_v15 = Companion_D6_.Create_DC18_(!(!((func() bool {
				if _this.F4() {
					return _this.F4()
				}
				return _this.F4()
			})())), ((_1025_v9).F14()).Minus(p0))
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1030_v14), 0))
			_ = _index185
			(_1030_v14).ArraySet1(_1032_v15, (_index185).Int())
		}
		var _1033_v16 _dafny.CodePoint
		_ = _1033_v16
		_1033_v16 = _dafny.CodePoint('s')
		var _1034_v17 _dafny.MultiSet
		_ = _1034_v17
		_1034_v17 = _dafny.MultiSetOf(_1033_v16)
		var _1035_v18 _dafny.Sequence
		_ = _1035_v18
		_1035_v18 = _dafny.SeqOf((_this).F7())
		var _1036_v19 _dafny.Map
		_ = _1036_v19
		_1036_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1035_v18, _1033_v16)
		var _1037_v20 _dafny.Map
		_ = _1037_v20
		_1037_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1036_v19).Cardinality(), _this.F4())
		var _1038_v21 _dafny.MultiSet
		_ = _1038_v21
		_1038_v21 = _dafny.MultiSetOf((func() _dafny.Int {
			if (_1034_v17).Contains(_1033_v16) {
				return (_1034_v17).Multiplicity(_1033_v16)
			}
			return (_this).F7()
		})(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), (func() bool {
			if (_1037_v20).Contains(_dafny.IntOfInt64(-496)) {
				return (_1037_v20).Get(_dafny.IntOfInt64(-496)).(bool)
			}
			return _this.F4()
		})())).Cardinality())
		var _1039_i3 _dafny.Int
		_ = _1039_i3
		_1039_i3 = _dafny.Zero
		{
			for (_1038_v21).IsDisjointFrom(_1038_v21) {
				{
					if (_1039_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1039_i3 = (_1039_i3).Plus(_dafny.One)
					var _1040_v22 _dafny.Set
					_ = _1040_v22
					_1040_v22 = _dafny.SetOf(_this.F4(), _this.F4())
					var _1041_v23 D6
					_ = _1041_v23
					_1041_v23 = Companion_D6_.Create_DC19_((_1040_v22).Cardinality())
					var _pat_let_tv18 = p0
					_ = _pat_let_tv18
					var _1042_v24 _dafny.Set
					_ = _1042_v24
					_1042_v24 = _dafny.SetOf(_1041_v23, _1041_v23, _1041_v23, func(_pat_let26_0 D6) D6 {
						return func(_1043_dt__update__tmp_h0 D6) D6 {
							return func(_pat_let27_0 _dafny.Int) D6 {
								return func(_1044_dt__update_hcf26_h0 _dafny.Int) D6 {
									return Companion_D6_.Create_DC19_(_1044_dt__update_hcf26_h0)
								}(_pat_let27_0)
							}((_dafny.Zero).Minus(_pat_let_tv18))
						}(_pat_let26_0)
					}(_1041_v23), Companion_D6_.Create_DC19_((_this).F3()))
					_1042_v24 = _1042_v24
					var _1045_v25 _dafny.Map
					_ = _1045_v25
					_1045_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), ((_dafny.Zero).Minus(p0)).Plus(_dafny.IntOfInt64(829)))
					_1045_v25 = (_1045_v25).Update(((_dafny.Zero).Minus((_this).F7())).Cmp(p0) <= 0, (_this).F3())
					r0 = p0
					var _1046_v26 _dafny.Sequence
					_ = _1046_v26
					_1046_v26 = _dafny.UnicodeSeqOfUtf8Bytes("rkni")
					var _1047_v27 _dafny.Map
					_ = _1047_v27
					_1047_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1046_v26)
					_1047_v27 = (_1047_v27).Update((_this).F3(), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
						if (_1047_v27).Contains(p0) {
							return (_1047_v27).Get(p0).(_dafny.Sequence)
						}
						return _1046_v26
					})(), _1046_v26))
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		(_this).F4_set_(_this.F4())
		var _1048_v28 _dafny.Array
		_ = _1048_v28
		var _len0_31 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_31
		var _nw188 _dafny.Array
		_ = _nw188
		if _len0_31.Cmp(_dafny.Zero) == 0 {
			_nw188 = _dafny.NewArray(_len0_31)
		} else {
			var _init31 func(_dafny.Int) _dafny.Int = func(_1049_i4 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_1049_i4, _dafny.IntOfInt64(453))
			}
			_ = _init31
			var _element0_31 = _init31(_dafny.Zero)
			_ = _element0_31
			_nw188 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
			(_nw188).ArraySet1(_element0_31, 0)
			var _nativeLen0_31 = (_len0_31).Int()
			_ = _nativeLen0_31
			for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
				(_nw188).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
			}
		}
		_1048_v28 = _nw188
		var _1050_v29 _dafny.Set
		_ = _1050_v29
		_1050_v29 = _dafny.SetOf((_this).F3())
		var _1051_v30 _dafny.Sequence
		_ = _1051_v30
		_1051_v30 = _dafny.UnicodeSeqOfUtf8Bytes("hnjykaw")
		var _nwElement0_41 _dafny.Int = p0
		_ = _nwElement0_41
		var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(8))
		_ = _nw189
		(_nw189).ArraySet1(_nwElement0_41, 0)
		(_nw189).ArraySet1(p0, 1)
		(_nw189).ArraySet1(p0, 2)
		(_nw189).ArraySet1((_this).F7(), 3)
		(_nw189).ArraySet1(Companion_Default___.SafeDivisionInt((_1050_v29).Cardinality(), _dafny.IntOfUint32((_1051_v30).Cardinality())), 4)
		(_nw189).ArraySet1(((_this).F3()).Times((_this).F3()), 5)
		(_nw189).ArraySet1(_dafny.IntOfInt64(327), 6)
		(_nw189).ArraySet1((_this).F7(), 7)
		_1048_v28 = _nw189
		var _rhs188 _dafny.Int = (_this).F7()
		_ = _rhs188
		var _rhs189 bool = _this.F4()
		_ = _rhs189
		var _rhs190 _dafny.Sequence = _1051_v30
		_ = _rhs190
		var _lhs176 *GlobalState = globalState
		_ = _lhs176
		var _lhs177 *GlobalState = globalState
		_ = _lhs177
		_lhs176.F1 = _rhs188
		_lhs177.F2 = _rhs189
		_1051_v30 = _rhs190
		r0 = _dafny.IntOfInt64(195)
		return r0
	}
}
func (_this *C5) M2(p0 bool, p1 _dafny.Array, p2 _dafny.Map, p3 bool, globalState *GlobalState) (_dafny.MultiSet, _dafny.CodePoint, _dafny.Map, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 bool = false
		_ = r3
		var _1052_v0 _dafny.Sequence
		_ = _1052_v0
		_1052_v0 = _dafny.SeqOf(_this.F4())
		var _1053_v1 D1
		_ = _1053_v1
		_1053_v1 = Companion_D1_.Create_DC7_()
		var _1054_v2 D0
		_ = _1054_v2
		_1054_v2 = Companion_D0_.Create_DC0_(_this.F4())
		(globalState).F1 = (((_this).F7()).Minus(_dafny.IntOfUint32((_1052_v0).Cardinality()))).Times(_dafny.IntOfUint32((Companion_Default___.Fm8(_this.F4(), _1053_v1, _1054_v2, (_this).F3(), globalState)).Cardinality()))
		var _1055_v3 _dafny.Array
		_ = _1055_v3
		var _len0_32 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_32
		var _nw190 _dafny.Array
		_ = _nw190
		if _len0_32.Cmp(_dafny.Zero) == 0 {
			_nw190 = _dafny.NewArray(_len0_32)
		} else {
			var _init32 func(_dafny.Int) bool = (func(_1056_p0 bool) func(_dafny.Int) bool {
				return func(_1057_i1 _dafny.Int) bool {
					return _1056_p0
				}
			})(p0)
			_ = _init32
			var _element0_32 = _init32(_dafny.Zero)
			_ = _element0_32
			_nw190 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
			(_nw190).ArraySet1(_element0_32, 0)
			var _nativeLen0_32 = (_len0_32).Int()
			_ = _nativeLen0_32
			for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
				(_nw190).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
			}
		}
		_1055_v3 = _nw190
		for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1055_v3), 0))); ; {
			_guard_loop_2, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _1058_i0 _dafny.Int
			_1058_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1058_i0).Sign() != -1) && ((_1058_i0).Cmp(_dafny.ArrayLen((_1055_v3), 0)) < 0)) {
				(_1055_v3).ArraySet1((_1052_v0).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), p3)).Cardinality(), _dafny.IntOfUint32((_1052_v0).Cardinality()))).Uint32()).(bool), (_1058_i0).Int())
			}
		}
		var _1059_v4 _dafny.Sequence
		_ = _1059_v4
		_1059_v4 = _dafny.UnicodeSeqOfUtf8Bytes("lgks")
		var _1060_v5 _dafny.CodePoint
		_ = _1060_v5
		_1060_v5 = _dafny.CodePoint('f')
		var _1061_v6 T0
		_ = _1061_v6
		var _nw191 *C3 = New_C3_()
		_ = _nw191
		_nw191.Ctor__(_1060_v5, _1059_v4, (_this).F7(), _this.F4(), (_this).F3())
		_1061_v6 = _nw191
		var _1062_v7 D1
		_ = _1062_v7
		_1062_v7 = Companion_D1_.Create_DC6_(_1061_v6, _1059_v4, _1061_v6.F4(), p3)
		var _rhs191 _dafny.Sequence = _1059_v4
		_ = _rhs191
		var _rhs192 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1062_v7).Dtor_cf10(), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("uuopatwnv"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uuopatwnv")).Cardinality()))).Uint32(), _1060_v5))
		_ = _rhs192
		_1059_v4 = _rhs191
		_1059_v4 = _rhs192
		var _rhs193 _dafny.Int = (_1061_v6).F3()
		_ = _rhs193
		var _rhs194 _dafny.Sequence = _1052_v0
		_ = _rhs194
		var _rhs195 _dafny.Int = (_1061_v6).F3()
		_ = _rhs195
		var _lhs178 *GlobalState = globalState
		_ = _lhs178
		var _lhs179 *GlobalState = globalState
		_ = _lhs179
		_lhs178.F1 = _rhs193
		_1052_v0 = _rhs194
		_lhs179.F1 = _rhs195
		var _1063_i2 _dafny.Int
		_ = _1063_i2
		_1063_i2 = _dafny.Zero
		{
			for _this.F4() {
				{
					if (_1063_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1063_i2 = (_1063_i2).Plus(_dafny.One)
					var _1064_v8 _dafny.Map
					_ = _1064_v8
					_1064_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1061_v6).F3(), _1061_v6.F4())
					var _1065_v9 D8
					_ = _1065_v9
					_1065_v9 = Companion_D8_.Create_DC22_(_1064_v8)
					var _1066_v10 _dafny.Array
					_ = _1066_v10
					var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
					_ = _nw192
					_1066_v10 = _nw192
					var _1067_v11 _dafny.Sequence
					_ = _1067_v11
					_1067_v11 = _dafny.SeqOf(_1066_v10)
					var _1068_v12 D12
					_ = _1068_v12
					_1068_v12 = Companion_D12_.Create_DC34_(_1066_v10)
					var _1069_v13 _dafny.Array
					_ = _1069_v13
					var _nwElement0_42 _dafny.Array = _1066_v10
					_ = _nwElement0_42
					var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(24))
					_ = _nw193
					(_nw193).ArraySet1(_nwElement0_42, 0)
					(_nw193).ArraySet1(_1066_v10, 1)
					(_nw193).ArraySet1(_1066_v10, 2)
					(_nw193).ArraySet1(_1066_v10, 3)
					(_nw193).ArraySet1((_1067_v11).Select((Companion_Default___.SafeIndex((_1061_v6).F3(), _dafny.IntOfUint32((_1067_v11).Cardinality()))).Uint32()).(_dafny.Array), 4)
					(_nw193).ArraySet1(_1066_v10, 5)
					(_nw193).ArraySet1((func() _dafny.Array {
						if _1061_v6.F4() {
							return _1066_v10
						}
						return _1066_v10
					})(), 6)
					(_nw193).ArraySet1(_1066_v10, 7)
					(_nw193).ArraySet1(_1066_v10, 8)
					(_nw193).ArraySet1(_1066_v10, 9)
					(_nw193).ArraySet1(_1066_v10, 10)
					(_nw193).ArraySet1(_1066_v10, 11)
					(_nw193).ArraySet1(_1066_v10, 12)
					(_nw193).ArraySet1(_1066_v10, 13)
					(_nw193).ArraySet1(_1066_v10, 14)
					(_nw193).ArraySet1(_1066_v10, 15)
					(_nw193).ArraySet1((Companion_D12_.Create_DC34_(_1066_v10)).Dtor_cf55(), 16)
					(_nw193).ArraySet1(_1066_v10, 17)
					(_nw193).ArraySet1((_1068_v12).Dtor_cf55(), 18)
					(_nw193).ArraySet1(_1066_v10, 19)
					(_nw193).ArraySet1(_1066_v10, 20)
					(_nw193).ArraySet1(_1066_v10, 21)
					(_nw193).ArraySet1(_1066_v10, 22)
					(_nw193).ArraySet1(_1066_v10, 23)
					_1069_v13 = _nw193
					var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_1069_v13), 0))
					_ = _index186
					(_1069_v13).ArraySet1(_1066_v10, (_index186).Int())
					var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_1055_v3), 0))
					_ = _index187
					(_1055_v3).ArraySet1(p3, (_index187).Int())
					var _1070_v14 _dafny.MultiSet
					_ = _1070_v14
					_1070_v14 = _dafny.MultiSetOf((_this).F3(), (_dafny.Zero).Minus((_this).F7()))
					var _pat_let_tv19 = _1064_v8
					_ = _pat_let_tv19
					var _pat_let_tv20 = _1061_v6
					_ = _pat_let_tv20
					var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_1069_v13), 0))
					_ = _index188
					var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_1055_v3), 0))
					_ = _index189
					var _rhs196 D8 = func(_pat_let28_0 D8) D8 {
						return func(_1071_dt__update__tmp_h0 D8) D8 {
							return func(_pat_let29_0 _dafny.Map) D8 {
								return func(_1072_dt__update_hcf31_h0 _dafny.Map) D8 {
									return Companion_D8_.Create_DC22_(_1072_dt__update_hcf31_h0)
								}(_pat_let29_0)
							}((_pat_let_tv19).Update((_pat_let_tv20).F3(), _this.F4()))
						}(_pat_let28_0)
					}(_1065_v9)
					_ = _rhs196
					var _rhs197 _dafny.Array = _1066_v10
					_ = _rhs197
					var _rhs198 bool = !(((_1070_v14).Cardinality()).Cmp((func() _dafny.Int {
						if _1061_v6.F4() {
							return (_this).F7()
						}
						return Companion_Default___.Fm0(globalState)
					})()) == 0)
					_ = _rhs198
					var _rhs199 bool = (func() bool {
						if p3 {
							return p3
						}
						return p3
					})()
					_ = _rhs199
					var _lhs180 _dafny.Array = _1069_v13
					_ = _lhs180
					var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_1069_v13), 0))
					_ = _lhs181
					var _lhs182 _dafny.Array = _1055_v3
					_ = _lhs182
					var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_1055_v3), 0))
					_ = _lhs183
					var _lhs184 T0 = _1061_v6
					_ = _lhs184
					_1065_v9 = _rhs196
					(_lhs180).ArraySet1(_rhs197, (_lhs181).Int())
					(_lhs182).ArraySet1(_rhs198, (_lhs183).Int())
					_lhs184.F4_set_(_rhs199)
					_1059_v4 = _dafny.UnicodeSeqOfUtf8Bytes("ncpxl")
					var _1073_v15 _dafny.Sequence
					_ = _1073_v15
					_1073_v15 = _dafny.SeqOf(p2)
					var _1074_v16 D13
					_ = _1074_v16
					_1074_v16 = Companion_D13_.Create_DC38_(_1073_v15)
					var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_1055_v3), 0))
					_ = _index190
					(_1055_v3).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((_1074_v16).Dtor_cf68(), _dafny.SeqOf(p2)), (_index190).Int())
					var _1075_v17 bool
					_ = _1075_v17
					var _1076_v18 _dafny.Int
					_ = _1076_v18
					var _1077_v19 _dafny.Sequence
					_ = _1077_v19
					var _1078_v20 _dafny.Int
					_ = _1078_v20
					var _out10 bool
					_ = _out10
					var _out11 _dafny.Int
					_ = _out11
					var _out12 _dafny.Sequence
					_ = _out12
					var _out13 _dafny.Int
					_ = _out13
					_out10, _out11, _out12, _out13 = (_this).M3((Companion_D13_.Create_DC39_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_1061_v6).F3(), globalState), (_this).F7()))).Dtor_cf69(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1059_v4, _1059_v4)).Cardinality()), globalState)
					_1075_v17 = _out10
					_1076_v18 = _out11
					_1077_v19 = _out12
					_1078_v20 = _out13
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1079_v21 _dafny.Sequence
		_ = _1079_v21
		_1079_v21 = _dafny.SeqOf(_dafny.IntOfInt64(534))
		var _1080_v22 _dafny.Map
		_ = _1080_v22
		_1080_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1079_v21, _1059_v4)
		var _1081_v23 D10
		_ = _1081_v23
		_1081_v23 = Companion_D10_.Create_DC28_(_1080_v22)
		var _pat_let_tv21 = _1059_v4
		_ = _pat_let_tv21
		var _pat_let_tv22 = _1060_v5
		_ = _pat_let_tv22
		var _pat_let_tv23 = _1080_v22
		_ = _pat_let_tv23
		(globalState).F1 = _dafny.IntOfUint32((func(_source13 D10) _dafny.Sequence {
			if _source13.Is_DC29() {
				var _1082___mcc_h0 bool = _source13.Get_().(D10_DC29).Cf42
				_ = _1082___mcc_h0
				var _1083___mcc_h1 _dafny.CodePoint = _source13.Get_().(D10_DC29).Cf43
				_ = _1083___mcc_h1
				var _1084_cf43 _dafny.CodePoint = _1083___mcc_h1
				_ = _1084_cf43
				var _1085_cf42 bool = _1082___mcc_h0
				_ = _1085_cf42
				return _pat_let_tv21
			} else if _source13.Is_DC30() {
				var _1086___mcc_h2 _dafny.Int = _source13.Get_().(D10_DC30).Cf44
				_ = _1086___mcc_h2
				var _1087___mcc_h3 _dafny.Array = _source13.Get_().(D10_DC30).Cf45
				_ = _1087___mcc_h3
				var _1088___mcc_h4 bool = _source13.Get_().(D10_DC30).Cf46
				_ = _1088___mcc_h4
				var _1089___mcc_h5 _dafny.Int = _source13.Get_().(D10_DC30).Cf47
				_ = _1089___mcc_h5
				var _1090___mcc_h6 _dafny.Int = _source13.Get_().(D10_DC30).Cf48
				_ = _1090___mcc_h6
				var _1091_cf48 _dafny.Int = _1090___mcc_h6
				_ = _1091_cf48
				var _1092_cf47 _dafny.Int = _1089___mcc_h5
				_ = _1092_cf47
				var _1093_cf46 bool = _1088___mcc_h4
				_ = _1093_cf46
				var _1094_cf45 _dafny.Array = _1087___mcc_h3
				_ = _1094_cf45
				var _1095_cf44 _dafny.Int = _1086___mcc_h2
				_ = _1095_cf44
				return _dafny.UnicodeSeqOfUtf8Bytes("whfgm")
			} else {
				var _1096___mcc_h7 _dafny.Map = _source13.Get_().(D10_DC28).Cf41
				_ = _1096___mcc_h7
				var _1097_cf41 _dafny.Map = _1096___mcc_h7
				_ = _1097_cf41
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(894))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_1098_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1099_i3 _dafny.Int) _dafny.CodePoint {
						return _1098_v5
					}
				})(_pat_let_tv22)))
			}
		}(func(_pat_let30_0 D10) D10 {
			return func(_1100_dt__update__tmp_h1 D10) D10 {
				return func(_pat_let31_0 _dafny.Map) D10 {
					return func(_1101_dt__update_hcf41_h0 _dafny.Map) D10 {
						return Companion_D10_.Create_DC28_(_1101_dt__update_hcf41_h0)
					}(_pat_let31_0)
				}(_pat_let_tv23)
			}(_pat_let30_0)
		}(_1081_v23))).Cardinality())
		var _1102_v24 D12
		_ = _1102_v24
		_1102_v24 = Companion_D12_.Create_DC36_(p0, _1061_v6.F4(), (_1061_v6).F3())
		var _1103_v25 _dafny.MultiSet
		_ = _1103_v25
		_1103_v25 = _dafny.MultiSetOf((func() bool {
			if (_1102_v24).Dtor_cf61() {
				return _this.F4()
			}
			return p3
		})())
		r0 = _1103_v25
		var _1104_v26 _dafny.MultiSet
		_ = _1104_v26
		_1104_v26 = _dafny.MultiSetOf((_1103_v25).Cardinality())
		r1 = (func() _dafny.CodePoint {
			if !(_1104_v26).Equals(_1104_v26) {
				return _1060_v5
			}
			return _dafny.CodePoint('o')
		})()
		r2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1060_v5)
		r3 = _this.F4()
		return r0, r1, r2, r3
	}
}
func (_this *C5) M3(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _hi10 _dafny.Int = Companion_Default___.Fm0(globalState)
		_ = _hi10
		for _1105_i0 := (_dafny.Zero).Minus((_this).F3()); _1105_i0.Cmp(_hi10) < 0; _1105_i0 = _1105_i0.Plus(_dafny.One) {
			var _1106_v0 _dafny.Array
			_ = _1106_v0
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw194
			_1106_v0 = _nw194
			var _rhs200 _dafny.Array = _1106_v0
			_ = _rhs200
			var _rhs201 _dafny.Int = ((_this).F3()).Plus(((Companion_Default___.Fm30(_this.F4(), (_this).F7(), (_this).F3(), globalState)).Dtor_cf66()).Minus((_this).F3()))
			_ = _rhs201
			_1106_v0 = _rhs200
			r1 = _rhs201
			var _1107_v1 _dafny.Sequence
			_ = _1107_v1
			_1107_v1 = _dafny.SeqOf(_this.F4())
			var _1108_v2 _dafny.Sequence
			_ = _1108_v2
			_1108_v2 = _dafny.UnicodeSeqOfUtf8Bytes("epm")
			var _1109_v3 _dafny.Sequence
			_ = _1109_v3
			_1109_v3 = _dafny.SeqOf(true, (_1107_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1108_v2).Cardinality()), _dafny.IntOfUint32((_1107_v1).Cardinality()))).Uint32()).(bool))
			_1107_v1 = _1107_v1
			var _1110_v4 _dafny.Set
			_ = _1110_v4
			_1110_v4 = _dafny.SetOf(_this.F4(), _this.F4(), _this.F4(), _this.F4(), _this.F4())
			var _1111_v5 _dafny.Sequence
			_ = _1111_v5
			_1111_v5 = _dafny.SeqOf((_dafny.IntOfUint32((_dafny.SeqOf((_this).F9())).Cardinality())).Minus((_this).F3()))
			var _rhs202 _dafny.Set = ((_1110_v4).Difference(_1110_v4)).Difference(_1110_v4)
			_ = _rhs202
			var _rhs203 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_1105_i0).Times(_1105_i0)))
			_ = _rhs203
			var _rhs204 bool = (_this.F4()) || (true)
			_ = _rhs204
			var _rhs205 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1111_v5, _1111_v5), _1111_v5)
			_ = _rhs205
			var _lhs185 *GlobalState = globalState
			_ = _lhs185
			var _lhs186 *GlobalState = globalState
			_ = _lhs186
			_1110_v4 = _rhs202
			_lhs185.F1 = _rhs203
			_lhs186.F2 = _rhs204
			_1111_v5 = _rhs205
			var _1112_v6 _dafny.Array
			_ = _1112_v6
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_33
			var _nw195 _dafny.Array
			_ = _nw195
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw195 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) _dafny.Int = (func(_1113_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_1114_i1 _dafny.Int) _dafny.Int {
						return (_1114_i1).Plus(_dafny.IntOfUint32((_1113_v2).Cardinality()))
					}
				})(_1108_v2)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw195 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw195).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw195).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1112_v6 = _nw195
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1112_v6), 0))
			_ = _index191
			(_1112_v6).ArraySet1(((_this).F3()).Times((_this).F3()), (_index191).Int())
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1112_v6), 0))
			_ = _index192
			var _rhs206 _dafny.Int = Companion_Default___.Fm0(globalState)
			_ = _rhs206
			var _rhs207 bool = (_1110_v4).IsSubsetOf(_dafny.SetOf(_this.F4()))
			_ = _rhs207
			var _rhs208 _dafny.Int = (_dafny.Zero).Minus((_1105_i0).Minus(p1))
			_ = _rhs208
			var _lhs187 *GlobalState = globalState
			_ = _lhs187
			var _lhs188 _dafny.Array = _1112_v6
			_ = _lhs188
			var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1112_v6), 0))
			_ = _lhs189
			_lhs187.F1 = _rhs206
			r0 = _rhs207
			(_lhs188).ArraySet1(_rhs208, (_lhs189).Int())
		}
		var _hi11 _dafny.Int = (_this).F3()
		_ = _hi11
		for _1115_i2 := p1; _1115_i2.Cmp(_hi11) < 0; _1115_i2 = _1115_i2.Plus(_dafny.One) {
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))
			_ = _index193
			((_this).F9()).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hpexysybx")).Cardinality())).Cmp((_this).F7()) <= 0, (_index193).Int())
			var _1116_v7 _dafny.MultiSet
			_ = _1116_v7
			_1116_v7 = _dafny.MultiSetOf(_this.F4(), _this.F4(), _this.F4())
			var _1117_v8 _dafny.Map
			_ = _1117_v8
			_1117_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F7(), (_this).F7()), Companion_Default___.Fm2(_1115_i2, globalState))
			var _1118_v9 _dafny.Sequence
			_ = _1118_v9
			_1118_v9 = _dafny.UnicodeSeqOfUtf8Bytes("he")
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))
			_ = _index194
			var _rhs209 bool = (_1116_v7).IsSubsetOf(Companion_Default___.Fm14(_this.F4(), false, globalState))
			_ = _rhs209
			var _rhs210 _dafny.Int = (_this).F7()
			_ = _rhs210
			var _rhs211 bool = (func() bool {
				if (_1117_v8).Contains((_this).F7()) {
					return (_1117_v8).Get((_this).F7()).(bool)
				}
				return !(_this.F4())
			})()
			_ = _rhs211
			var _rhs212 _dafny.Int = ((_this).F7()).Times(_dafny.IntOfUint32((_1118_v9).Cardinality()))
			_ = _rhs212
			var _rhs213 _dafny.Int = ((func() _dafny.Int {
				if _this.F4() {
					return (_this).F7()
				}
				return p1
			})()).Minus((_dafny.Zero).Minus((_this).F3()))
			_ = _rhs213
			var _lhs190 _dafny.Array = (_this).F9()
			_ = _lhs190
			var _lhs191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))
			_ = _lhs191
			var _lhs192 *GlobalState = globalState
			_ = _lhs192
			var _lhs193 *GlobalState = globalState
			_ = _lhs193
			(_lhs190).ArraySet1(_rhs209, (_lhs191).Int())
			r1 = _rhs210
			_lhs192.F2 = _rhs211
			_lhs193.F1 = _rhs212
			r3 = _rhs213
			var _1119_v10 _dafny.Array
			_ = _1119_v10
			var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw196
			_1119_v10 = _nw196
			var _1120_v11 D6
			_ = _1120_v11
			_1120_v11 = Companion_D6_.Create_DC18_(((_this).F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))).Int()).(bool), _dafny.IntOfUint32((_1118_v9).Cardinality()))
			var _1121_v12 _dafny.Array
			_ = _1121_v12
			var _nwElement0_43 bool = ((_this).F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))).Int()).(bool)
			_ = _nwElement0_43
			var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(5))
			_ = _nw197
			(_nw197).ArraySet1(_nwElement0_43, 0)
			(_nw197).ArraySet1(_this.F4(), 1)
			(_nw197).ArraySet1(_this.F4(), 2)
			(_nw197).ArraySet1((_1120_v11).Dtor_cf24(), 3)
			(_nw197).ArraySet1(!(Companion_Default___.Fm2(p1, globalState)), 4)
			_1121_v12 = _nw197
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_1119_v10), 0))
			_ = _index195
			(_1119_v10).ArraySet1(_1121_v12, (_index195).Int())
			var _1122_v13 _dafny.Sequence
			_ = _1122_v13
			_1122_v13 = _dafny.SeqOf(true, ((_this).F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))).Int()).(bool), ((_this).F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))).Int()).(bool), !(((_this).F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F9()), 0))).Int()).(bool)))
			var _1123_v14 D12
			_ = _1123_v14
			_1123_v14 = Companion_D12_.Create_DC35_(_1121_v12, _this.F4(), _1122_v13, _dafny.UnicodeSeqOfUtf8Bytes("tgykad"), p1)
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_1119_v10), 0))
			_ = _index196
			(_1119_v10).ArraySet1((_1123_v14).Dtor_cf56(), (_index196).Int())
			r1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_this).F7(), (_this).F3()), (_dafny.Zero).Minus(p1))
			var _1124_v15 _dafny.Map
			_ = _1124_v15
			_1124_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), false)
			var _1125_v16 _dafny.Array
			_ = _1125_v16
			var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
			_ = _nw198
			_1125_v16 = _nw198
			var _1126_v17 D12
			_ = _1126_v17
			_1126_v17 = Companion_D12_.Create_DC34_(_1125_v16)
			var _1127_v18 _dafny.Map
			_ = _1127_v18
			_1127_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1124_v15).Cardinality(), _1126_v17)
			var _1128_v19 _dafny.MultiSet
			_ = _1128_v19
			_1128_v19 = _dafny.MultiSetOf(_1127_v18)
			_1128_v19 = (func() _dafny.MultiSet {
				if false {
					return (func() _dafny.MultiSet {
						if true {
							return _1128_v19
						}
						return _1128_v19
					})()
				}
				return _1128_v19
			})()
		}
		var _1129_v20 _dafny.Map
		_ = _1129_v20
		_1129_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.UnicodeSeqOfUtf8Bytes("q"))
		var _1130_v21 D7
		_ = _1130_v21
		_1130_v21 = Companion_D7_.Create_DC20_(_1129_v20)
		var _1131_v22 _dafny.MultiSet
		_ = _1131_v22
		_1131_v22 = _dafny.MultiSetOf(_1130_v21, _1130_v21)
		if ((_1131_v22).Union(_1131_v22)).Contains((func() D7 {
			if _this.F4() {
				return Companion_D7_.Create_DC20_(_1129_v20)
			}
			return _1130_v21
		})()) {
			var _1132_v23 _dafny.CodePoint
			_ = _1132_v23
			_1132_v23 = _dafny.CodePoint('v')
			var _1133_v24 _dafny.Set
			_ = _1133_v24
			_1133_v24 = _dafny.SetOf(_1132_v23)
			var _1134_v25 _dafny.Map
			_ = _1134_v25
			_1134_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1133_v24, _1132_v23)
			var _1135_v26 *C1
			_ = _1135_v26
			var _nw199 *C1 = New_C1_()
			_ = _nw199
			_nw199.Ctor__((_1134_v25).Cardinality(), (_this).F3(), _this.F4())
			_1135_v26 = _nw199
			var _1136_v27 _dafny.Set
			_ = _1136_v27
			_1136_v27 = _dafny.SetOf(_1135_v26, _1135_v26, _1135_v26, _1135_v26)
			_1136_v27 = (_dafny.SetOf(_1135_v26, _1135_v26, _1135_v26)).Union(_dafny.SetOf(_1135_v26))
			var _1137_v28 *C4
			_ = _1137_v28
			var _nw200 *C4 = New_C4_()
			_ = _nw200
			_nw200.Ctor__((_this).F3(), (_this).F3(), (p1).Cmp((_this).F7()) >= 0)
			_1137_v28 = _nw200
			r1 = (_this).F7()
			var _1138_v29 _dafny.Array
			_ = _1138_v29
			var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
			_ = _nw201
			_1138_v29 = _nw201
			var _1139_v30 _dafny.Sequence
			_ = _1139_v30
			_1139_v30 = _dafny.SeqOf(_this.F4())
			var _1140_v31 _dafny.Sequence
			_ = _1140_v31
			_1140_v31 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _1141_v32 D12
			_ = _1141_v32
			_1141_v32 = Companion_D12_.Create_DC35_((_this).F9(), true, _1139_v30, _1140_v31, (_this).F3())
			var _1142_v33 _dafny.Sequence
			_ = _1142_v33
			_1142_v33 = _dafny.SeqOf((_1141_v32).Dtor_cf57())
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_1138_v29), 0))
			_ = _index197
			(_1138_v29).ArraySet1(_1139_v30, (_index197).Int())
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_1138_v29), 0))
			_ = _index198
			(_1138_v29).ArraySet1(_1139_v30, (_index198).Int())
			var _1143_v34 _dafny.Array
			_ = _1143_v34
			var _nw202 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw202
			_1143_v34 = _nw202
			_1143_v34 = (_this).F9()
		} else {
			var _1144_v35 _dafny.Set
			_ = _1144_v35
			_1144_v35 = _dafny.SetOf(_this.F4())
			var _1145_v36 _dafny.Sequence
			_ = _1145_v36
			_1145_v36 = _dafny.UnicodeSeqOfUtf8Bytes("tjqkmu")
			var _rhs214 _dafny.Set = ((_1144_v35).Intersection(_dafny.SetOf(false, _this.F4()))).Union(_1144_v35)
			_ = _rhs214
			var _rhs215 bool = ((func() _dafny.Int {
				if _this.F4() {
					return p1
				}
				return (_this).F7()
			})()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1145_v36, _dafny.UnicodeSeqOfUtf8Bytes("spyelrt"))).Cardinality())) <= 0
			_ = _rhs215
			var _lhs194 *C5 = _this
			_ = _lhs194
			_1144_v35 = _rhs214
			_lhs194.F4_set_(_rhs215)
			var _1146_v37 _dafny.Map
			_ = _1146_v37
			_1146_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _this.F4())
			_1146_v37 = (_1146_v37).Update(_this.F4(), ((_this).F3()).Cmp(p1) == 0)
			var _1147_v38 _dafny.MultiSet
			_ = _1147_v38
			_1147_v38 = _dafny.MultiSetOf((_this).F7(), (_dafny.IntOfInt64(362)).Plus((_this).F3()), (_this).F7())
			_1147_v38 = ((_1147_v38).Update((_dafny.Zero).Minus((_this).F3()), Companion_Default___.Abs(p1))).Update((func() _dafny.Int {
				if (p0).Contains(_this.F4()) {
					return (p0).Get(_this.F4()).(_dafny.Int)
				}
				return p1
			})(), Companion_Default___.Abs((_dafny.Zero).Minus((p1).Minus(p1))))
			var _1148_v39 _dafny.CodePoint
			_ = _1148_v39
			_1148_v39 = _dafny.CodePoint('o')
			var _1149_v40 _dafny.Map
			_ = _1149_v40
			_1149_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1145_v36, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1145_v36).Cardinality()))).Uint32(), _1148_v39), (_this).F7())
			var _1150_v41 D12
			_ = _1150_v41
			_1150_v41 = Companion_D12_.Create_DC37_(_1146_v37, _this.F4(), (func() _dafny.Int {
				if (_1149_v40).Contains(_dafny.UnicodeSeqOfUtf8Bytes("cmrcf")) {
					return (_1149_v40).Get(_dafny.UnicodeSeqOfUtf8Bytes("cmrcf")).(_dafny.Int)
				}
				return (_this).F3()
			})(), (_this).F3())
			r3 = (_1150_v41).Dtor_cf67()
			var _1151_v42 _dafny.Map
			_ = _1151_v42
			_1151_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F3()).Plus(_dafny.IntOfInt64(-416)), (_this).F7())
			_1151_v42 = (_1151_v42).Update((_this).F3(), (_this).F3())
		}
		var _1152_v43 _dafny.Sequence
		_ = _1152_v43
		_1152_v43 = _dafny.SeqOf(_this.F4())
		(globalState).F1 = (func() _dafny.Int {
			if (_1152_v43).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_1152_v43).Cardinality()))).Uint32()).(bool) {
				return Companion_Default___.SafeModuloInt(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-567))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}(func(_1153_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality()))
			}
			return _dafny.IntOfInt64(762)
		})()
		var _1154_v44 *C4
		_ = _1154_v44
		var _nw203 *C4 = New_C4_()
		_ = _nw203
		_nw203.Ctor__((_dafny.Zero).Minus(((_this).F7()).Minus((_this).F3())), (_dafny.Zero).Minus(Companion_Default___.Fm0(globalState)), Companion_Default___.Fm2(_dafny.IntOfUint32((_1152_v43).Cardinality()), globalState))
		_1154_v44 = _nw203
		var _1155_v45 _dafny.Array
		_ = _1155_v45
		var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw204
		_1155_v45 = _nw204
		_1155_v45 = _1155_v45
		r0 = _this.F4()
		var _1156_v46 _dafny.Sequence
		_ = _1156_v46
		_1156_v46 = _dafny.UnicodeSeqOfUtf8Bytes("gtmwd")
		r1 = _dafny.IntOfUint32((_1156_v46).Cardinality())
		r2 = Companion_Default___.Fm15(_this.F4(), Companion_Default___.SafeModuloInt(p1, (_this).F3()), _dafny.IntOfInt64(639), globalState)
		r3 = (_dafny.Zero).Minus((_this).F3())
		return r0, r1, r2, r3
	}
}
func (_this *C5) F9() _dafny.Array {
	{
		return _this._f9
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f4 bool
	_f7 _dafny.Int
	_f3 _dafny.Int
	_f8 _dafny.Int
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f4 = false
	_this._f7 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f8 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C6{}
var _ T1 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F4() bool {
	return _this._f4
}
func (_this *C6) F4_set_(value bool) {
	_this._f4 = value
}
func (_this *C6) F7() _dafny.Int {
	return _this._f7
}
func (_this *C6) F3() _dafny.Int {
	return _this._f3
}
func (_this *C6) Ctor__(f8 _dafny.Int, f3 _dafny.Int, f4 bool, f7 _dafny.Int) {
	{
		(_this)._f8 = f8
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f7 = f7
	}
}
func (_this *C6) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		if _this.F4() {
			var _1157_v0 _dafny.Sequence
			_ = _1157_v0
			_1157_v0 = _dafny.SeqOf(_this.F4())
			var _1158_v1 _dafny.Map
			_ = _1158_v1
			_1158_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _dafny.IntOfUint32((_1157_v0).Cardinality()))
			_1158_v1 = (_1158_v1).Update((_this.F4()) && (_this.F4()), (func() _dafny.Int {
				if _this.F4() {
					return (_this).F3()
				}
				return (_this).F3()
			})())
			var _1159_v2 *C4
			_ = _1159_v2
			var _nw205 *C4 = New_C4_()
			_ = _nw205
			_nw205.Ctor__(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qvoflbfav")).Cardinality())), ((_this).F3()).Times((_this).F3()), _this.F4())
			_1159_v2 = _nw205
			var _1160_v3 _dafny.Array
			_ = _1160_v3
			var _nw206 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw206
			_1160_v3 = _nw206
			var _1161_v4 _dafny.Array
			_ = _1161_v4
			var _nwElement0_44 _dafny.Array = _1160_v3
			_ = _nwElement0_44
			var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(4))
			_ = _nw207
			(_nw207).ArraySet1(_nwElement0_44, 0)
			(_nw207).ArraySet1(_1160_v3, 1)
			(_nw207).ArraySet1(_1160_v3, 2)
			(_nw207).ArraySet1(_1160_v3, 3)
			_1161_v4 = _nw207
			_1161_v4 = _1161_v4
			(globalState).F1 = ((_this).F7()).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F8())))
			var _1162_v5 *C1
			_ = _1162_v5
			var _nw208 *C1 = New_C1_()
			_ = _nw208
			_nw208.Ctor__(Companion_Default___.Fm0(globalState), (_this).F8(), _this.F4())
			_1162_v5 = _nw208
		} else {
			var _1163_v6 _dafny.Sequence
			_ = _1163_v6
			_1163_v6 = _dafny.UnicodeSeqOfUtf8Bytes("kbtb")
			var _1164_v7 _dafny.MultiSet
			_ = _1164_v7
			_1164_v7 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_1163_v6, _1163_v6), _1163_v6)
			var _1165_v8 _dafny.Sequence
			_ = _1165_v8
			_1165_v8 = _dafny.SeqOf(_1163_v6, _1163_v6)
			_1164_v7 = ((func() _dafny.MultiSet {
				if _this.F4() {
					return _1164_v7
				}
				return (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("oeofpknod"), _1163_v6, _dafny.UnicodeSeqOfUtf8Bytes("ou"), _1163_v6)).Update(_1163_v6, Companion_Default___.Abs(p0))
			})()).Difference((_1164_v7).Intersection((_dafny.MultiSetFromSeq(_1165_v8)).Update(_1163_v6, Companion_Default___.Abs(p0))))
			var _1166_v9 _dafny.Array
			_ = _1166_v9
			var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw209
			_1166_v9 = _nw209
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1166_v9), 0))
			_ = _index199
			(_1166_v9).ArraySet1(_this.F4(), (_index199).Int())
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1166_v9), 0))
			_ = _index200
			(_1166_v9).ArraySet1(_this.F4(), (_index200).Int())
			var _1167_v10 _dafny.CodePoint
			_ = _1167_v10
			_1167_v10 = _dafny.CodePoint('o')
			_1167_v10 = _1167_v10
			var _1168_v11 _dafny.Array
			_ = _1168_v11
			var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw210
			_1168_v11 = _nw210
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1168_v11), 0))
			_ = _index201
			(_1168_v11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1166_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1166_v9), 0))).Int()).(bool)), (_index201).Int())
			var _1169_v12 _dafny.Map
			_ = _1169_v12
			_1169_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), Companion_Default___.Fm2(p0, globalState))
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1168_v11), 0))
			_ = _index202
			(_1168_v11).ArraySet1((_1169_v12).Merge(_1169_v12), (_index202).Int())
			var _rhs216 _dafny.Int = (func() _dafny.Int {
				if (_1166_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1166_v9), 0))).Int()).(bool) {
					return (_this).F3()
				}
				return (_this).F7()
			})()
			_ = _rhs216
			var _rhs217 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1163_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_1170_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1171_i0 _dafny.Int) _dafny.CodePoint {
					return _1170_v10
				}
			})(_1167_v10)))), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F8(), p0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1163_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_1172_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1173_i0 _dafny.Int) _dafny.CodePoint {
					return _1172_v10
				}
			})(_1167_v10))))).Cardinality()))).Uint32(), _1167_v10)
			_ = _rhs217
			var _rhs218 bool = (_1166_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1166_v9), 0))).Int()).(bool)
			_ = _rhs218
			var _rhs219 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm20((_this).F8(), globalState), _1163_v6)
			_ = _rhs219
			var _rhs220 bool = true
			_ = _rhs220
			var _lhs195 *GlobalState = globalState
			_ = _lhs195
			var _lhs196 *GlobalState = globalState
			_ = _lhs196
			var _lhs197 *GlobalState = globalState
			_ = _lhs197
			_lhs195.F1 = _rhs216
			_1163_v6 = _rhs217
			_lhs196.F2 = _rhs218
			_1163_v6 = _rhs219
			_lhs197.F2 = _rhs220
		}
		var _1174_v13 _dafny.Array
		_ = _1174_v13
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_34
		var _nw211 _dafny.Array
		_ = _nw211
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw211 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) bool = func(_1175_i1 _dafny.Int) bool {
				return _this.F4()
			}
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw211 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw211).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw211).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1174_v13 = _nw211
		var _1176_v14 *C5
		_ = _1176_v14
		var _nw212 *C5 = New_C5_()
		_ = _nw212
		_nw212.Ctor__(_1174_v13, (func() _dafny.Int {
			if !(_this.F4()) {
				return (_this).F8()
			}
			return (_this).F8()
		})(), (_this).F7(), !(_this.F4()))
		_1176_v14 = _nw212
		var _1177_v15 _dafny.Sequence
		_ = _1177_v15
		_1177_v15 = _dafny.UnicodeSeqOfUtf8Bytes("euk")
		var _1178_v16 _dafny.Sequence
		_ = _1178_v16
		_1178_v16 = _dafny.SeqOf(_1177_v15, _1177_v15)
		if !(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_1178_v16, _1178_v16), _1177_v15)) {
			var _1179_v17 _dafny.Map
			_ = _1179_v17
			_1179_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F4())
			var _1180_v18 _dafny.Set
			_ = _1180_v18
			_1180_v18 = _dafny.SetOf((_1179_v17).Cardinality(), (_this).F8(), (_dafny.SetOf((_this).F7(), (_this).F7())).Cardinality())
			var _1181_v19 D4
			_ = _1181_v19
			_1181_v19 = Companion_D4_.Create_DC11_(_1180_v18)
			var _source14 D4 = _1181_v19
			_ = _source14
			if _source14.Is_DC12() {
				var _1182_v20 _dafny.Array
				_ = _1182_v20
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_35
				var _nw213 _dafny.Array
				_ = _nw213
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw213 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.Int = func(_1183_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1183_i2, (_this).F3())
					}
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw213 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw213).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw213).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1182_v20 = _nw213
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1182_v20), 0))
				_ = _index203
				(_1182_v20).ArraySet1((_this).F7(), (_index203).Int())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1182_v20), 0))
				_ = _index204
				(_1182_v20).ArraySet1(Companion_Default___.Fm0(globalState), (_index204).Int())
				var _1184_v21 *C2
				_ = _1184_v21
				var _nw214 *C2 = New_C2_()
				_ = _nw214
				_nw214.Ctor__(Companion_Default___.SafeModuloInt((_this).F3(), _dafny.IntOfInt64(-240)), _1182_v20, (_this).F3(), _this.F4())
				_1184_v21 = _nw214
				(globalState).F2 = _this.F4()
				(globalState).F1 = _dafny.IntOfInt64(-741)
			} else if _source14.Is_DC11() {
				var _1185___mcc_h0 _dafny.Set = _source14.Get_().(D4_DC11).Cf18
				_ = _1185___mcc_h0
				var _1186_cf18 _dafny.Set = _1185___mcc_h0
				_ = _1186_cf18
				(globalState).F1 = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jacm")).Cardinality())).Minus(Companion_Default___.SafeModuloInt((_this).F8(), (_this).F7()))
				_1177_v15 = _dafny.UnicodeSeqOfUtf8Bytes("cmeqw")
				var _1187_v22 _dafny.Set
				_ = _1187_v22
				_1187_v22 = _dafny.SetOf(_this.F4())
				var _1188_v23 _dafny.MultiSet
				_ = _1188_v23
				_1188_v23 = _dafny.MultiSetOf((_this).F3(), p0, (_1187_v22).Cardinality())
				var _1189_v24 _dafny.Map
				_ = _1189_v24
				_1189_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _this.F4())
				var _1190_v25 _dafny.Sequence
				_ = _1190_v25
				_1190_v25 = _dafny.SeqOf(_this.F4(), _this.F4(), (func() bool {
					if (_1189_v24).Contains(p0) {
						return (_1189_v24).Get(p0).(bool)
					}
					return _this.F4()
				})())
				var _1191_v26 D3
				_ = _1191_v26
				_1191_v26 = Companion_D3_.Create_DC10_(_dafny.IntOfUint32((_1177_v15).Cardinality()), (_1176_v14).F9(), (func() _dafny.Int {
					if (_1188_v23).Contains((_dafny.Zero).Minus((_this).F8())) {
						return (_1188_v23).Multiplicity((_dafny.Zero).Minus((_this).F8()))
					}
					return _dafny.IntOfUint32((_1190_v25).Cardinality())
				})())
				var _1192_v28 *C5
				_ = _1192_v28
				var _nw215 *C5 = New_C5_()
				_ = _nw215
				_nw215.Ctor__((func() _dafny.Array {
					if _this.F4() {
						return (_1191_v26).Dtor_cf16()
					}
					return (_1176_v14).F9()
				})(), (func() _dafny.Set {
					var _coll20 = _dafny.NewBuilder()
					_ = _coll20
					for _iter23 := _dafny.Iterate((_1189_v24).Keys().Elements()); ; {
						_compr_20, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _1193_v27 _dafny.Int
						_1193_v27 = interface{}(_compr_20).(_dafny.Int)
						if (_1189_v24).Contains(_1193_v27) {
							_coll20.Add((_1193_v27).Plus(_dafny.IntOfInt64(-533)))
						}
					}
					return _coll20.ToSet()
				}()).Cardinality(), (_dafny.MultiSetFromSeq(_1190_v25)).Cardinality(), Companion_Default___.Fm2((_this).F8(), globalState))
				_1192_v28 = _nw215
				var _1194_v29 _dafny.Map
				_ = _1194_v29
				_1194_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _1174_v13)
				var _1195_v30 *C0
				_ = _1195_v30
				var _nw216 *C0 = New_C0_()
				_ = _nw216
				_nw216.Ctor__(_this.F4(), (func() _dafny.Array {
					if (_1194_v29).Contains(_this.F4()) {
						return (_1194_v29).Get(_this.F4()).(_dafny.Array)
					}
					return (_1176_v14).F9()
				})())
				_1195_v30 = _nw216
			} else {
				var _1196___mcc_h1 D4 = _source14.Get_().(D4_DC13).Cf19
				_ = _1196___mcc_h1
				var _1197_cf19 D4 = _1196___mcc_h1
				_ = _1197_cf19
				var _1198_v31 _dafny.Array
				_ = _1198_v31
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_36
				var _nw217 _dafny.Array
				_ = _nw217
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw217 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = func(_1199_i3 _dafny.Int) _dafny.Int {
						return (_1199_i3).Times(_dafny.IntOfInt64(862))
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw217 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw217).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw217).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1198_v31 = _nw217
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))
				_ = _index205
				(_1198_v31).ArraySet1(p0, (_index205).Int())
				var _1200_v32 _dafny.CodePoint
				_ = _1200_v32
				_1200_v32 = _dafny.CodePoint('j')
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))
				_ = _index206
				var _rhs221 _dafny.Int = Companion_Default___.SafeModuloInt(((_dafny.Zero).Minus((_this).F8())).Plus((_this).F7()), (_dafny.Zero).Minus((Companion_Default___.Fm9(_1200_v32, _this.F4(), globalState)).Cardinality()))
				_ = _rhs221
				var _rhs222 bool = !(_this.F4())
				_ = _rhs222
				var _lhs198 _dafny.Array = _1198_v31
				_ = _lhs198
				var _lhs199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))
				_ = _lhs199
				var _lhs200 *GlobalState = globalState
				_ = _lhs200
				(_lhs198).ArraySet1(_rhs221, (_lhs199).Int())
				_lhs200.F2 = _rhs222
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))
				_ = _index207
				(_1198_v31).ArraySet1((func() _dafny.Int {
					if _this.F4() {
						return p0
					}
					return _dafny.IntOfInt64(819)
				})(), (_index207).Int())
				var _1201_v33 _dafny.Sequence
				_ = _1201_v33
				_1201_v33 = _dafny.SeqOf(((_this).F3()).Minus(p0), (_this).F3(), (_this).F8())
				var _1202_v34 _dafny.Sequence
				_ = _1202_v34
				_1202_v34 = _dafny.SeqOf(_this.F4())
				var _1203_v35 _dafny.Map
				_ = _1203_v35
				_1203_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _1200_v32)
				var _1204_v36 _dafny.Set
				_ = _1204_v36
				_1204_v36 = _dafny.SetOf(_1200_v32, (func() _dafny.CodePoint {
					if (_1203_v35).Contains((_1198_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))).Int()).(_dafny.Int)) {
						return (_1203_v35).Get((_1198_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))).Int()).(_dafny.Int)).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('k')
				})(), _1200_v32)
				var _1205_v37 _dafny.Map
				_ = _1205_v37
				_1205_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1204_v36).Cardinality(), _this.F4())
				var _1206_v38 *C1
				_ = _1206_v38
				var _nw218 *C1 = New_C1_()
				_ = _nw218
				_nw218.Ctor__((_this).F3(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(542))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}(func(_1207_i4 _dafny.Int) _dafny.Int {
					return (_this).F8()
				}))).Cardinality()), _this.F4())
				_1206_v38 = _nw218
				var _1208_v39 _dafny.Map
				_ = _1208_v39
				_1208_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(233), _1206_v38)
				var _1209_v40 _dafny.MultiSet
				_ = _1209_v40
				_1209_v40 = _dafny.MultiSetOf((func() *C1 {
					if _this.F4() {
						return _1206_v38
					}
					return _1206_v38
				})(), _1206_v38, _1206_v38, (func() *C1 {
					if (_1208_v39).Contains((_this).F8()) {
						return (_1208_v39).Get((_this).F8()).(*C1)
					}
					return _1206_v38
				})(), _1206_v38)
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))
				_ = _index208
				var _rhs223 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1201_v33, (Companion_Default___.SafeIndex(((func() _dafny.Map {
					if _this.F4() {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1202_v34).Cardinality()), _this.F4())
					}
					return _1205_v37
				})()).Cardinality(), _dafny.IntOfUint32((_1201_v33).Cardinality()))).Uint32(), p0)
				_ = _rhs223
				var _rhs224 _dafny.Int = (func() _dafny.Int {
					if (_1209_v40).Contains((func() *C1 {
						if _this.F4() {
							return _1206_v38
						}
						return _1206_v38
					})()) {
						return (_1209_v40).Multiplicity((func() *C1 {
							if _this.F4() {
								return _1206_v38
							}
							return _1206_v38
						})())
					}
					return Companion_Default___.SafeModuloInt(p0, (_this).F8())
				})()
				_ = _rhs224
				var _rhs225 bool = (_1202_v34).Select((Companion_Default___.SafeIndex(((_this).F3()).Plus(Companion_Default___.Fm0(globalState)), _dafny.IntOfUint32((_1202_v34).Cardinality()))).Uint32()).(bool)
				_ = _rhs225
				var _rhs226 bool = _this.F4()
				_ = _rhs226
				var _rhs227 _dafny.Int = (_1198_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))).Int()).(_dafny.Int)
				_ = _rhs227
				var _lhs201 _dafny.Array = _1198_v31
				_ = _lhs201
				var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1198_v31), 0))
				_ = _lhs202
				var _lhs203 *GlobalState = globalState
				_ = _lhs203
				var _lhs204 *C6 = _this
				_ = _lhs204
				var _lhs205 *GlobalState = globalState
				_ = _lhs205
				_1201_v33 = _rhs223
				(_lhs201).ArraySet1(_rhs224, (_lhs202).Int())
				_lhs203.F2 = _rhs225
				_lhs204.F4_set_(_rhs226)
				_lhs205.F1 = _rhs227
				_1198_v31 = _1198_v31
			}
			var _1210_v41 _dafny.Map
			_ = _1210_v41
			_1210_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F4()) || (!(_this.F4())), (_this).F3())
			var _1211_v42 *C4
			_ = _1211_v42
			var _nw219 *C4 = New_C4_()
			_ = _nw219
			_nw219.Ctor__(_dafny.IntOfInt64(-621), (_this).F7(), _this.F4())
			_1211_v42 = _nw219
			var _1212_v43 _dafny.Sequence
			_ = _1212_v43
			_1212_v43 = _dafny.SeqOf(_1211_v42, _1211_v42)
			var _1213_v44 _dafny.Map
			_ = _1213_v44
			_1213_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F4()), _1212_v43)
			var _1214_v45 _dafny.Sequence
			_ = _1214_v45
			_1214_v45 = (func() _dafny.Sequence {
				if (_1213_v44).Contains(true) {
					return (_1213_v44).Get(true).(_dafny.Sequence)
				}
				return _1212_v43
			})()
			_1210_v41 = (_1210_v41).Update(_this.F4(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1214_v45), _1212_v43)).Cardinality()))
			var _1215_v46 _dafny.CodePoint
			_ = _1215_v46
			_1215_v46 = _dafny.CodePoint('f')
			var _1216_v47 _dafny.Sequence
			_ = _1216_v47
			_1216_v47 = _dafny.SeqOf(_this.F4())
			var _1217_v48 *C3
			_ = _1217_v48
			var _nw220 *C3 = New_C3_()
			_ = _nw220
			_nw220.Ctor__(_1215_v46, _dafny.UnicodeSeqOfUtf8Bytes("iqrkkklva"), (_this).F8(), _this.F4(), _dafny.IntOfUint32((_1216_v47).Cardinality()))
			_1217_v48 = _nw220
			var _1218_v49 _dafny.Map
			_ = _1218_v49
			_1218_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1217_v48, _this.F4())
			var _1219_v50 _dafny.Map
			_ = _1219_v50
			_1219_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _1218_v49)
			_1218_v49 = ((_1218_v49).Merge(_1218_v49)).Merge(((func() _dafny.Map {
				if (_1219_v50).Contains(_this.F4()) {
					return (_1219_v50).Get(_this.F4()).(_dafny.Map)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1217_v48, _this.F4())).Update(_1217_v48, _this.F4())
			})()).Merge(_1218_v49))
			var _1220_v51 _dafny.Array
			_ = _1220_v51
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_37
			var _nw221 _dafny.Array
			_ = _nw221
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw221 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Int = func(_1221_i5 _dafny.Int) _dafny.Int {
					return (_1221_i5).Minus((_this).F8())
				}
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw221 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw221).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw221).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1220_v51 = _nw221
			var _1222_v52 _dafny.MultiSet
			_ = _1222_v52
			_1222_v52 = _dafny.MultiSetOf(_1220_v51, _1220_v51)
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_1220_v51), 0))
			_ = _index209
			(_1220_v51).ArraySet1((_1222_v52).Cardinality(), (_index209).Int())
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_1220_v51), 0))
			_ = _index210
			(_1220_v51).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(46), (_this).F7()), (_index210).Int())
			var _1223_v53 _dafny.Set
			_ = _1223_v53
			_1223_v53 = _dafny.SetOf(_this.F4())
			var _1224_v54 _dafny.Map
			_ = _1224_v54
			_1224_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _1223_v53)
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_1220_v51), 0))
			_ = _index211
			(_1220_v51).ArraySet1(((((func() _dafny.Set {
				if (_1224_v54).Contains(p0) {
					return (_1224_v54).Get(p0).(_dafny.Set)
				}
				return _1223_v53
			})()).Difference(_1223_v53)).Difference(_1223_v53)).Cardinality(), (_index211).Int())
		} else {
			var _1225_v55 _dafny.Sequence
			_ = _1225_v55
			_1225_v55 = _dafny.SeqOf((_1176_v14).F9(), (_1176_v14).F9(), (_1176_v14).F9())
			_1225_v55 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1225_v55, (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_1225_v55).Cardinality()))).Uint32(), (_1176_v14).F9()), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1225_v55, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1225_v55).Cardinality()))).Uint32(), _1174_v13), _1225_v55))
			var _1226_v56 _dafny.Map
			_ = _1226_v56
			_1226_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F4()), _this.F4())
			_1226_v56 = _1226_v56
			var _1227_v57 _dafny.Sequence
			_ = _1227_v57
			_1227_v57 = _dafny.SeqOf(_this.F4(), _this.F4())
			var _1228_v58 _dafny.Set
			_ = _1228_v58
			_1228_v58 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1227_v57, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1227_v57).Cardinality()))).Uint32(), _this.F4()), _1227_v57), _1227_v57, _1227_v57, _1227_v57)
			var _1229_v59 _dafny.Sequence
			_ = _1229_v59
			_1229_v59 = _dafny.SeqOf(_dafny.SetOf(_1227_v57), (_1228_v58).Union(_dafny.SetOf(_1227_v57)))
			_1228_v58 = (_1229_v59).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1229_v59).Cardinality()))).Uint32()).(_dafny.Set)
			var _1230_v60 _dafny.MultiSet
			_ = _1230_v60
			_1230_v60 = _dafny.MultiSetOf(_this.F4())
			var _1231_v61 _dafny.Set
			_ = _1231_v61
			_1231_v61 = _dafny.SetOf((_1230_v60).Cardinality(), _dafny.IntOfInt64(818))
			_1177_v15 = Companion_Default___.Fm15(((_1231_v61).Cardinality()).Cmp((_this).F8()) < 0, (_this).F8(), (_this).F8(), globalState)
			r0 = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_this.F4()) || (_this.F4()) {
					return (_this).F7()
				}
				return (_dafny.Zero).Minus(_dafny.IntOfInt64(-893))
			})()))
		}
		(globalState).F2 = _this.F4()
		var _1232_v62 D4
		_ = _1232_v62
		_1232_v62 = Companion_D4_.Create_DC12_()
		var _1233_v63 _dafny.Set
		_ = _1233_v63
		_1233_v63 = _dafny.SetOf(_this.F4())
		var _1234_v64 _dafny.Sequence
		_ = _1234_v64
		_1234_v64 = _dafny.SeqOf(_this.F4())
		var _1235_v65 _dafny.Sequence
		_ = _1235_v65
		_1235_v65 = _dafny.SeqOf(p0)
		var _1236_v66 _dafny.Map
		_ = _1236_v66
		_1236_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1235_v65).Cardinality()), (_this).F8())
		var _1237_v67 D3
		_ = _1237_v67
		_1237_v67 = Companion_D3_.Create_DC10_((_this).F7(), (_1176_v14).F9(), (_dafny.Zero).Minus((_this).F3()))
		var _1238_v68 _dafny.Array
		_ = _1238_v68
		var _nwElement0_45 _dafny.Int = (_this).F3()
		_ = _nwElement0_45
		var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(24))
		_ = _nw222
		(_nw222).ArraySet1(_nwElement0_45, 0)
		(_nw222).ArraySet1(p0, 1)
		(_nw222).ArraySet1((_this).F8(), 2)
		(_nw222).ArraySet1(_dafny.IntOfInt64(688), 3)
		(_nw222).ArraySet1(_dafny.IntOfUint32((_1177_v15).Cardinality()), 4)
		(_nw222).ArraySet1((_1233_v63).Cardinality(), 5)
		(_nw222).ArraySet1(p0, 6)
		(_nw222).ArraySet1((_this).F8(), 7)
		(_nw222).ArraySet1((_this).F7(), 8)
		(_nw222).ArraySet1((_this).F8(), 9)
		(_nw222).ArraySet1(_dafny.IntOfUint32((_1234_v64).Cardinality()), 10)
		(_nw222).ArraySet1(Companion_Default___.Fm0(globalState), 11)
		(_nw222).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(831))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_1239_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))).Cardinality()), 12)
		(_nw222).ArraySet1((_dafny.MultiSetFromSeq(_1235_v65)).Cardinality(), 13)
		(_nw222).ArraySet1(((_1236_v66).Update(p0, p0)).Cardinality(), 14)
		(_nw222).ArraySet1(p0, 15)
		(_nw222).ArraySet1((_1237_v67).Dtor_cf17(), 16)
		(_nw222).ArraySet1((Companion_Default___.Fm31(globalState)).Cardinality(), 17)
		(_nw222).ArraySet1((_this).F3(), 18)
		(_nw222).ArraySet1(p0, 19)
		(_nw222).ArraySet1((_this).F8(), 20)
		(_nw222).ArraySet1(p0, 21)
		(_nw222).ArraySet1((_this).F8(), 22)
		(_nw222).ArraySet1((_this).F3(), 23)
		_1238_v68 = _nw222
		var _1240_v69 _dafny.Sequence
		_ = _1240_v69
		_1240_v69 = _dafny.SeqOf(_1238_v68)
		var _1241_v70 *C2
		_ = _1241_v70
		var _nw223 *C2 = New_C2_()
		_ = _nw223
		_nw223.Ctor__(p0, _1238_v68, (_this).F7(), _this.F4())
		_1241_v70 = _nw223
		var _1242_v71 _dafny.Sequence
		_ = _1242_v71
		_1242_v71 = _dafny.SeqOf(_1241_v70, _1241_v70, _1241_v70)
		var _1243_v72 _dafny.Sequence
		_ = _1243_v72
		_1243_v72 = _dafny.SeqOf(_1240_v69)
		var _1244_v73 _dafny.CodePoint
		_ = _1244_v73
		_1244_v73 = _dafny.CodePoint('b')
		var _1245_v74 D8
		_ = _1245_v74
		_1245_v74 = Companion_D8_.Create_DC23_(_this.F4(), _this.F4(), _this.F4(), _1244_v73)
		var _1246_v75 D0
		_ = _1246_v75
		_1246_v75 = Companion_D0_.Create_DC1_(_this.F4(), _this.F4(), _1244_v73, (_1241_v70).F15())
		var _pat_let_tv24 = _1235_v65
		_ = _pat_let_tv24
		var _pat_let_tv25 = _1235_v65
		_ = _pat_let_tv25
		var _pat_let_tv26 = _1235_v65
		_ = _pat_let_tv26
		var _pat_let_tv27 = _1241_v70
		_ = _pat_let_tv27
		var _pat_let_tv28 = _1177_v15
		_ = _pat_let_tv28
		var _pat_let_tv29 = _1241_v70
		_ = _pat_let_tv29
		var _pat_let_tv30 = _1244_v73
		_ = _pat_let_tv30
		var _pat_let_tv31 = _1244_v73
		_ = _pat_let_tv31
		var _rhs228 D4 = (func() D4 {
			if _this.F4() {
				return _1232_v62
			}
			return Companion_D4_.Create_DC12_()
		})()
		_ = _rhs228
		var _rhs229 _dafny.Sequence = (_1243_v72).Select((Companion_Default___.SafeIndex((_1241_v70).F14(), _dafny.IntOfUint32((_1243_v72).Cardinality()))).Uint32()).(_dafny.Sequence)
		_ = _rhs229
		var _rhs230 _dafny.Sequence = func(_source15 D8) _dafny.Sequence {
			if _source15.Is_DC23() {
				var _1247___mcc_h2 bool = _source15.Get_().(D8_DC23).Cf32
				_ = _1247___mcc_h2
				var _1248___mcc_h3 bool = _source15.Get_().(D8_DC23).Cf33
				_ = _1248___mcc_h3
				var _1249___mcc_h4 bool = _source15.Get_().(D8_DC23).Cf34
				_ = _1249___mcc_h4
				var _1250___mcc_h5 _dafny.CodePoint = _source15.Get_().(D8_DC23).Cf35
				_ = _1250___mcc_h5
				var _1251_cf35 _dafny.CodePoint = _1250___mcc_h5
				_ = _1251_cf35
				var _1252_cf34 bool = _1249___mcc_h4
				_ = _1252_cf34
				var _1253_cf33 bool = _1248___mcc_h3
				_ = _1253_cf33
				var _1254_cf32 bool = _1247___mcc_h2
				_ = _1254_cf32
				return _dafny.Companion_Sequence_.Update(_pat_let_tv24, (Companion_Default___.SafeIndex((_dafny.MultiSetOf((Companion_D0_.Create_DC0_(_1252_cf34)).Dtor_cf0())).Cardinality(), _dafny.IntOfUint32((_pat_let_tv25).Cardinality()))).Uint32(), (_this).F8())
			} else if _source15.Is_DC22() {
				var _1255___mcc_h6 _dafny.Map = _source15.Get_().(D8_DC22).Cf31
				_ = _1255___mcc_h6
				var _1256_cf31 _dafny.Map = _1255___mcc_h6
				_ = _1256_cf31
				return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv26, _dafny.SeqOf((_dafny.Zero).Minus((_pat_let_tv27).F14()), (_dafny.MultiSetOf(_this.F4())).Cardinality()))
			} else {
				var _1257___mcc_h7 D8 = _source15.Get_().(D8_DC24).Cf36
				_ = _1257___mcc_h7
				var _1258_cf36 D8 = _1257___mcc_h7
				_ = _1258_cf36
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(724))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_1259_v15 _dafny.Sequence, _1260_v70 *C2, _1261_v73 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
					return func(_1262_i7 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1259_v15, (Companion_Default___.SafeIndex((_1260_v70).F14(), _dafny.IntOfUint32((_1259_v15).Cardinality()))).Uint32(), _1261_v73)).Cardinality())
					}
				})(_pat_let_tv28, _pat_let_tv29, _pat_let_tv30)))
			}
		}(_1245_v74)
		_ = _rhs230
		var _rhs231 bool = (func(_pat_let32_0 D0) D0 {
			return func(_1263_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let33_0 _dafny.CodePoint) D0 {
					return func(_1264_dt__update_hcf3_h0 _dafny.CodePoint) D0 {
						return Companion_D0_.Create_DC1_((_1263_dt__update__tmp_h0).Dtor_cf1(), (_1263_dt__update__tmp_h0).Dtor_cf2(), _1264_dt__update_hcf3_h0, (_1263_dt__update__tmp_h0).Dtor_cf4())
					}(_pat_let33_0)
				}(_pat_let_tv31)
			}(_pat_let32_0)
		}(_1246_v75)).Dtor_cf2()
		_ = _rhs231
		var _rhs232 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1242_v71, _1242_v71), _1242_v71)
		_ = _rhs232
		var _lhs206 *GlobalState = globalState
		_ = _lhs206
		_1232_v62 = _rhs228
		_1240_v69 = _rhs229
		_1235_v65 = _rhs230
		_lhs206.F2 = _rhs231
		_1242_v71 = _rhs232
		var _1265_v76 _dafny.Map
		_ = _1265_v76
		_1265_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_1234_v64).Cardinality()))
		var _1266_v77 _dafny.MultiSet
		_ = _1266_v77
		_1266_v77 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1235_v65).Cardinality()), (func() _dafny.Int {
			if (_1265_v76).Contains(_this.F4()) {
				return (_1265_v76).Get(_this.F4()).(_dafny.Int)
			}
			return (_this).F3()
		})(), (_this).F7(), (_this).F8())
		var _1267_v78 *C3
		_ = _1267_v78
		var _nw224 *C3 = New_C3_()
		_ = _nw224
		_nw224.Ctor__(_1244_v73, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-483))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}((func(_1268_v73 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1269_i8 _dafny.Int) _dafny.CodePoint {
				return _1268_v73
			}
		})(_1244_v73))), _dafny.IntOfUint32((_1177_v15).Cardinality()), _this.F4(), (((_1266_v77).Update(p0, Companion_Default___.Abs(p0))).Update((_this).F8(), Companion_Default___.Abs(_dafny.IntOfInt64(-529)))).Cardinality())
		_1267_v78 = _nw224
		r0 = (_1241_v70).F14()
		return r0
	}
}
func (_this *C6) F8() _dafny.Int {
	{
		return _this._f8
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f4 bool
	_f3 _dafny.Int
	F5  bool
	_f6 _dafny.Int
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f4 = false
	_this._f3 = _dafny.Zero
	_this.F5 = false
	_this._f6 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F4() bool {
	return _this._f4
}
func (_this *C7) F4_set_(value bool) {
	_this._f4 = value
}
func (_this *C7) F3() _dafny.Int {
	return _this._f3
}
func (_this *C7) Ctor__(f5 bool, f6 _dafny.Int, f3 _dafny.Int, f4 bool) {
	{
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C7) Fm3(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('f')
	}
}
func (_this *C7) M0(globalState *GlobalState) (_dafny.Sequence, _dafny.Int) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		(_this).F5 = _this.F5
		var _1270_v0 D1
		_ = _1270_v0
		_1270_v0 = Companion_D1_.Create_DC5_((_this).F3())
		var _hi12 _dafny.Int = ((_this).F3()).Plus((_1270_v0).Dtor_cf8())
		_ = _hi12
		for _1271_i0 := (_this).F6(); _1271_i0.Cmp(_hi12) < 0; _1271_i0 = _1271_i0.Plus(_dafny.One) {
			var _1272_v1 _dafny.Sequence
			_ = _1272_v1
			_1272_v1 = _dafny.SeqOf(_this.F4())
			var _1273_v2 _dafny.Sequence
			_ = _1273_v2
			_1273_v2 = _dafny.SeqOf(!((_1272_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.IntOfUint32((_1272_v1).Cardinality()))).Uint32()).(bool)), _this.F5, _this.F4(), _this.F4(), _this.F4())
			var _1274_v3 _dafny.Array
			_ = _1274_v3
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_38
			var _nw225 _dafny.Array
			_ = _nw225
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw225 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) bool = func(_1275_i1 _dafny.Int) bool {
					return _this.F4()
				}
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw225 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw225).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw225).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1274_v3 = _nw225
			var _1276_v4 _dafny.Map
			_ = _1276_v4
			_1276_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1274_v3, _this.F4())
			var _1277_v5 _dafny.Sequence
			_ = _1277_v5
			_1277_v5 = _dafny.SeqOf((_this).F6())
			var _1278_v6 D0
			_ = _1278_v6
			_1278_v6 = Companion_D0_.Create_DC0_(_this.F4())
			var _1279_v7 _dafny.Sequence
			_ = _1279_v7
			_1279_v7 = _dafny.UnicodeSeqOfUtf8Bytes("wyblf")
			var _1280_v8 _dafny.MultiSet
			_ = _1280_v8
			_1280_v8 = _dafny.MultiSetOf(_this.F5)
			var _1281_v9 _dafny.Array
			_ = _1281_v9
			var _nwElement0_46 bool = _this.F5
			_ = _nwElement0_46
			var _nw226 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(24))
			_ = _nw226
			(_nw226).ArraySet1(_nwElement0_46, 0)
			(_nw226).ArraySet1((_1272_v1).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1272_v1).Cardinality()))).Uint32()).(bool), 1)
			(_nw226).ArraySet1(_this.F4(), 2)
			(_nw226).ArraySet1(_this.F5, 3)
			(_nw226).ArraySet1(_this.F4(), 4)
			(_nw226).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm4(_this.F5, (func() bool {
				if (_1276_v4).Contains(_1274_v3) {
					return (_1276_v4).Get(_1274_v3).(bool)
				}
				return _this.F5
			})(), _this.F5, _this.F4(), globalState), _1277_v5), 5)
			(_nw226).ArraySet1(_this.F5, 6)
			(_nw226).ArraySet1(((_this).F3()).Cmp(_1271_i0) != 0, 7)
			(_nw226).ArraySet1(_this.F4(), 8)
			(_nw226).ArraySet1((_1278_v6).Dtor_cf0(), 9)
			(_nw226).ArraySet1(_this.F4(), 10)
			(_nw226).ArraySet1(!(_this.F4()), 11)
			(_nw226).ArraySet1(_this.F5, 12)
			(_nw226).ArraySet1(_this.F5, 13)
			(_nw226).ArraySet1(_this.F4(), 14)
			(_nw226).ArraySet1(((_this).F6()).Cmp(_dafny.IntOfUint32((_1279_v7).Cardinality())) < 0, 15)
			(_nw226).ArraySet1(_this.F5, 16)
			(_nw226).ArraySet1(_this.F5, 17)
			(_nw226).ArraySet1((_dafny.MultiSetOf(_this.F5, _this.F4())).IsDisjointFrom(_1280_v8), 18)
			(_nw226).ArraySet1(_this.F4(), 19)
			(_nw226).ArraySet1(!(_this.F5) || (_this.F5), 20)
			(_nw226).ArraySet1(_this.F4(), 21)
			(_nw226).ArraySet1(_this.F4(), 22)
			(_nw226).ArraySet1(!(!(!(_this.F4()))), 23)
			_1281_v9 = _nw226
			var _1282_v10 _dafny.Sequence
			_ = _1282_v10
			_1282_v10 = _dafny.SeqOf(_1281_v9, _1274_v3, _1281_v9)
			_1281_v9 = (_1282_v10).Select((Companion_Default___.SafeIndex((_this).F3(), _dafny.IntOfUint32((_1282_v10).Cardinality()))).Uint32()).(_dafny.Array)
			var _1283_v11 _dafny.Map
			_ = _1283_v11
			_1283_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1271_i0), (_this).F6()), _dafny.UnicodeSeqOfUtf8Bytes("u"))
			_1283_v11 = (_1283_v11).Update((_this).F3(), _1279_v7)
			(globalState).F1 = (_this).F6()
			var _1284_v12 *C6
			_ = _1284_v12
			var _nw227 *C6 = New_C6_()
			_ = _nw227
			_nw227.Ctor__(Companion_Default___.SafeDivisionInt((_this).F3(), (_dafny.Zero).Minus((_this).F3())), (_this).F3(), true, (_this).F6())
			_1284_v12 = _nw227
		}
		var _1285_v13 _dafny.Array
		_ = _1285_v13
		var _len0_39 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_39
		var _nw228 _dafny.Array
		_ = _nw228
		if _len0_39.Cmp(_dafny.Zero) == 0 {
			_nw228 = _dafny.NewArray(_len0_39)
		} else {
			var _init39 func(_dafny.Int) _dafny.Int = func(_1286_i3 _dafny.Int) _dafny.Int {
				return (_1286_i3).Plus((_this).F3())
			}
			_ = _init39
			var _element0_39 = _init39(_dafny.Zero)
			_ = _element0_39
			_nw228 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
			(_nw228).ArraySet1(_element0_39, 0)
			var _nativeLen0_39 = (_len0_39).Int()
			_ = _nativeLen0_39
			for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
				(_nw228).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
			}
		}
		_1285_v13 = _nw228
		for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1285_v13), 0))); ; {
			_guard_loop_3, _ok24 := _iter24()
			if !_ok24 {
				break
			}
			var _1287_i2 _dafny.Int
			_1287_i2 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1287_i2).Sign() != -1) && ((_1287_i2).Cmp(_dafny.ArrayLen((_1285_v13), 0)) < 0)) {
				(_1285_v13).ArraySet1(Companion_Default___.SafeModuloInt(_1287_i2, _dafny.IntOfInt64(947)), (_1287_i2).Int())
			}
		}
		var _1288_v14 _dafny.Sequence
		_ = _1288_v14
		_1288_v14 = _dafny.SeqOf(_this.F4())
		var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_1285_v13), 0))
		_ = _index212
		(_1285_v13).ArraySet1(_dafny.IntOfUint32((_1288_v14).Cardinality()), (_index212).Int())
		var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_1285_v13), 0))
		_ = _index213
		(_1285_v13).ArraySet1(Companion_Default___.Fm0(globalState), (_index213).Int())
		(_this).F4_set_(_this.F4())
		(_this).F4_set_((_this.F4()) && (_this.F4()))
		r0 = _1288_v14
		r1 = (_this).F3()
		return r0, r1
	}
}
func (_this *C7) F6() _dafny.Int {
	{
		return _this._f6
	}
}

// End of class C7
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
