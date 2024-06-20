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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ibq"), _dafny.UnicodeSeqOfUtf8Bytes("qjdrauoj")), Companion_Default___.SafeDivisionInt((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(163))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(68)
		})))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(163))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_0_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(68)
			})))).Contains(_1_v0) {
				_coll0.Add((_1_v0).Plus(_dafny.IntOfInt64(-976)), !(false))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(120)))
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('h')
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('i')), _dafny.SeqOf(_dafny.CodePoint('n'), _dafny.CodePoint('o'))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('a')), _dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('d'), _dafny.CodePoint('d'))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _dafny.SeqOf(_dafny.CodePoint('g')))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('a'))))).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v0 _dafny.CodePoint
			_2_v0 = interface{}(_compr_1).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Contains(_2_v0) {
				_coll1.Add(_2_v0, _dafny.SeqOf(_dafny.CodePoint('p')))
			}
		}
		return _coll1.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Set, globalState *GlobalState) bool {
	var _source0 D3 = Companion_D3_.Create_DC8_(_dafny.CodePoint('h'), !(false), (Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ndasii")).Cardinality()))).Dtor_cf1(), _dafny.UnicodeSeqOfUtf8Bytes("aqwpe"))
	_ = _source0
	if _source0.Is_DC8() {
		var _3___mcc_h0 _dafny.CodePoint = _source0.Get_().(D3_DC8).Cf17
		_ = _3___mcc_h0
		var _4___mcc_h1 bool = _source0.Get_().(D3_DC8).Cf18
		_ = _4___mcc_h1
		var _5___mcc_h2 bool = _source0.Get_().(D3_DC8).Cf19
		_ = _5___mcc_h2
		var _6___mcc_h3 _dafny.Sequence = _source0.Get_().(D3_DC8).Cf20
		_ = _6___mcc_h3
		var _7_cf20 _dafny.Sequence = _6___mcc_h3
		_ = _7_cf20
		var _8_cf19 bool = _5___mcc_h2
		_ = _8_cf19
		var _9_cf18 bool = _4___mcc_h1
		_ = _9_cf18
		var _10_cf17 _dafny.CodePoint = _3___mcc_h0
		_ = _10_cf17
		return !(_9_cf18)
	} else if _source0.Is_DC7() {
		var _11___mcc_h4 _dafny.Sequence = _source0.Get_().(D3_DC7).Cf16
		_ = _11___mcc_h4
		var _12_cf16 _dafny.Sequence = _11___mcc_h4
		_ = _12_cf16
		return (false) == (true)
	} else {
		var _13___mcc_h5 D3 = _source0.Get_().(D3_DC9).Cf21
		_ = _13___mcc_h5
		var _14_cf21 D3 = _13___mcc_h5
		_ = _14_cf21
		return (true) || (true)
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(247), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-345), _dafny.IntOfInt64(863))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xdtj")).Cardinality())), _dafny.IntOfInt64(-684))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality())), false)).Cardinality()))).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32(((Companion_D6_.Create_DC17_((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(998), _dafny.IntOfInt64(270))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _15_v0 _dafny.Int
			_15_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(998)).Cmp(_15_v0) <= 0) && ((_15_v0).Cmp(_dafny.IntOfInt64(270)) < 0) {
				_coll2.Add((_15_v0).Minus(_dafny.IntOfInt64(322)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(true)).Cardinality())).Cardinality())
			}
		}
		return _coll2.ToMap()
	}()).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(653))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), false)).Dtor_cf39()).Cardinality()))) >= 0, !(true) || (!(false)))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(!(false)), !(!(true)))).Union(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(true, true)
		}
		return _dafny.SeqOf(false, true)
	})(), _dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) Fm16(p0 D0, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(9))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC9_(Companion_D3_.Create_DC9_(Companion_D3_.Create_DC8_(_dafny.CodePoint('k'), false, true, _dafny.UnicodeSeqOfUtf8Bytes("bvdjclmpx")))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC9_(Companion_D3_.Create_DC9_(Companion_D3_.Create_DC9_(Companion_D3_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(92))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))))))))).Cardinality(), (_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(!(true))))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if !(!(true)) {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), false)
	})()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Int {
	if true {
		return Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(717), (_dafny.SetOf(false)).Cardinality())).Cardinality(), _dafny.IntOfInt64(59))
	} else {
		return (_dafny.MultiSetOf(false, false, false, false, true)).Cardinality()
	}
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return ((func() D1 {
		if true {
			return Companion_D1_.Create_DC4_(!(false), true, _dafny.IntOfInt64(747), _dafny.SeqOf(_dafny.IntOfInt64(883), _dafny.IntOfInt64(416)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('u'))).Cardinality()))
		}
		return Companion_D1_.Create_DC4_(false, true, _dafny.IntOfInt64(-294), (Companion_D1_.Create_DC4_(false, false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality(), false)).Cardinality(), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mpil")).Cardinality()), _dafny.IntOfInt64(275), _dafny.IntOfInt64(715), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-940))), _dafny.IntOfInt64(17))).Dtor_cf8(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(306), !(true))).Cardinality(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(230))).Cardinality()))).Cardinality())
	})()).Dtor_cf8()
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) D0 {
	var _source1 D3 = Companion_D3_.Create_DC8_(_dafny.CodePoint('t'), true, false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})))
	_ = _source1
	if _source1.Is_DC8() {
		var _19___mcc_h0 _dafny.CodePoint = _source1.Get_().(D3_DC8).Cf17
		_ = _19___mcc_h0
		var _20___mcc_h1 bool = _source1.Get_().(D3_DC8).Cf18
		_ = _20___mcc_h1
		var _21___mcc_h2 bool = _source1.Get_().(D3_DC8).Cf19
		_ = _21___mcc_h2
		var _22___mcc_h3 _dafny.Sequence = _source1.Get_().(D3_DC8).Cf20
		_ = _22___mcc_h3
		var _23_cf20 _dafny.Sequence = _22___mcc_h3
		_ = _23_cf20
		var _24_cf19 bool = _21___mcc_h2
		_ = _24_cf19
		var _25_cf18 bool = _20___mcc_h1
		_ = _25_cf18
		var _26_cf17 _dafny.CodePoint = _19___mcc_h0
		_ = _26_cf17
		return Companion_D0_.Create_DC0_((_23_cf20).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.IntOfUint32((_23_cf20).Cardinality()))).Uint32()).(_dafny.CodePoint))
	} else if _source1.Is_DC7() {
		var _27___mcc_h4 _dafny.Sequence = _source1.Get_().(D3_DC7).Cf16
		_ = _27___mcc_h4
		var _28_cf16 _dafny.Sequence = _27___mcc_h4
		_ = _28_cf16
		return Companion_D0_.Create_DC0_(_dafny.CodePoint('d'))
	} else {
		var _29___mcc_h5 D3 = _source1.Get_().(D3_DC9).Cf21
		_ = _29___mcc_h5
		var _30_cf21 D3 = _29___mcc_h5
		_ = _30_cf21
		return Companion_D0_.Create_DC0_(_dafny.CodePoint('e'))
	}
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("icl"), _dafny.UnicodeSeqOfUtf8Bytes("x")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(522))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_31_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(202))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_32_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_(_dafny.CodePoint('w'), false, ((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-362), true)).Cardinality())).Cmp((_dafny.Zero).Minus((_dafny.MultiSetOf(!(!(false)), true, false)).Cardinality())) >= 0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ykgprxn"), _dafny.UnicodeSeqOfUtf8Bytes("bkcd")))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(!(false), true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(!(false)))))
}
func (_static *CompanionStruct_Default___) Fm25(p0 D9, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(501), _dafny.UnicodeSeqOfUtf8Bytes("lifhwcw"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(360), _dafny.UnicodeSeqOfUtf8Bytes("altifx")))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(161))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_33_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-977))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_34_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))
	})))).Intersection(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ujto"), _dafny.UnicodeSeqOfUtf8Bytes("pfcmr"), _dafny.UnicodeSeqOfUtf8Bytes("jnkits"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(70))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_35_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))))).Difference((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ynytpecsu"))
		}
		return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("x"))
	})())
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, p3 _dafny.MultiSet, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('k'))).Cardinality()))).Minus(_dafny.IntOfInt64(4)), false, (func() _dafny.Int {
		if false {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfInt64(-758)))).Cardinality()
		}
		return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetOf(_dafny.IntOfInt64(715))).Cardinality())).Cardinality())
	})(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_36_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(764)
	}))).Cardinality()), _dafny.IntOfInt64(938), _dafny.IntOfInt64(739), (_dafny.MultiSetOf(!(true))).Cardinality(), _dafny.Zero)).IsSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(352), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wfi")).Cardinality()), _dafny.IntOfInt64(763))).Cardinality())))).Cardinality()), _dafny.IntOfInt64(980), (func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(12), _dafny.IntOfInt64(21))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _37_v0 _dafny.Int
			_37_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(12)).Cmp(_37_v0) <= 0) && ((_37_v0).Cmp(_dafny.IntOfInt64(21)) < 0) {
				_coll3.Add((_37_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(279))).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('c'))).Cardinality())
			}
		}
		return _coll3.ToMap()
	}()).Cardinality())), _dafny.IntOfInt64(-549))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(465))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(895)))).Merge((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(855))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(677))
	})())
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(242))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_39_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jcgiuvj")).Cardinality())
	}))).Cardinality()))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-791), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("s"))).Cardinality()), _dafny.IntOfInt64(-678))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-166))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_40_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Map, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() _dafny.Int {
		if !(!(true)) {
			return (_dafny.MultiSetOf(_dafny.IntOfInt64(657))).Cardinality()
		}
		return _dafny.IntOfInt64(769)
	})())
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source2 D8 = Companion_D8_.Create_DC19_(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(473), _dafny.IntOfInt64(374))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _41_v0 _dafny.Int
			_41_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(473)).Cmp(_41_v0) <= 0) && ((_41_v0).Cmp(_dafny.IntOfInt64(374)) < 0) {
				_coll4.Add((_41_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jwdxy")).Cardinality())), false)
			}
		}
		return _coll4.ToMap()
	}())
	_ = _source2
	if _source2.Is_DC20() {
		var _42___mcc_h0 _dafny.CodePoint = _source2.Get_().(D8_DC20).Cf43
		_ = _42___mcc_h0
		var _43___mcc_h1 bool = _source2.Get_().(D8_DC20).Cf44
		_ = _43___mcc_h1
		var _44___mcc_h2 bool = _source2.Get_().(D8_DC20).Cf45
		_ = _44___mcc_h2
		var _45___mcc_h3 _dafny.Map = _source2.Get_().(D8_DC20).Cf46
		_ = _45___mcc_h3
		var _46___mcc_h4 bool = _source2.Get_().(D8_DC20).Cf47
		_ = _46___mcc_h4
		var _47_cf47 bool = _46___mcc_h4
		_ = _47_cf47
		var _48_cf46 _dafny.Map = _45___mcc_h3
		_ = _48_cf46
		var _49_cf45 bool = _44___mcc_h2
		_ = _49_cf45
		var _50_cf44 bool = _43___mcc_h1
		_ = _50_cf44
		var _51_cf43 _dafny.CodePoint = _42___mcc_h0
		_ = _51_cf43
		return _dafny.SeqOf(_47_cf47)
	} else {
		var _52___mcc_h5 _dafny.Map = _source2.Get_().(D8_DC19).Cf42
		_ = _52___mcc_h5
		var _53_cf42 _dafny.Map = _52___mcc_h5
		_ = _53_cf42
		return _dafny.SeqOf(false, false)
	}
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sqdajpdkh")).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(311), !(false)))).Merge((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(229), _dafny.IntOfInt64(-103))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _54_v0 _dafny.Int
			_54_v0 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(229)).Cmp(_54_v0) <= 0) && ((_54_v0).Cmp(_dafny.IntOfInt64(-103)) < 0) {
				_coll5.Add(Companion_Default___.SafeModuloInt(_54_v0, _dafny.IntOfInt64(-349)), true)
			}
		}
		return _coll5.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.CodePoint('j'))).Cardinality(), false)))
}
func (_static *CompanionStruct_Default___) Fm35(globalState *GlobalState) D6 {
	if false {
		if true {
			return Companion_D6_.Create_DC17_(_dafny.IntOfInt64(121), _dafny.UnicodeSeqOfUtf8Bytes("h"), (Companion_D9_.Create_DC22_(_dafny.IntOfInt64(19), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(173))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_55_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			}))).Cardinality()), false))).Dtor_cf50())
		} else {
			return Companion_D6_.Create_DC17_(_dafny.IntOfInt64(119), _dafny.UnicodeSeqOfUtf8Bytes("ndg"), false)
		}
	} else {
		return Companion_D6_.Create_DC17_(_dafny.IntOfInt64(695), _dafny.UnicodeSeqOfUtf8Bytes("stgojmq"), true)
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Sequence, p1 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.MultiSet, _dafny.Set, _dafny.Int) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
	_ = r1
	var r2 _dafny.Set = _dafny.EmptySet
	_ = r2
	var r3 _dafny.Int = _dafny.Zero
	_ = r3
	var _56_v0 bool
	_ = _56_v0
	_56_v0 = false
	var _57_i0 _dafny.Int
	_ = _57_i0
	_57_i0 = _dafny.Zero
	{
		for _56_v0 {
			{
				if (_57_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_57_i0 = (_57_i0).Plus(_dafny.One)
				var _58_v1 _dafny.Array
				_ = _58_v1
				var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw0
				_58_v1 = _nw0
				var _59_v2 _dafny.Int
				_ = _59_v2
				_59_v2 = _dafny.IntOfInt64(849)
				var _60_v3 _dafny.Set
				_ = _60_v3
				_60_v3 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_59_v2, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), !(_56_v0))).Cardinality()), _59_v2, _59_v2)
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_58_v1), 0))
				_ = _index0
				(_58_v1).ArraySet1((_59_v2).Times((_60_v3).Cardinality()), (_index0).Int())
				var _61_v4 _dafny.Sequence
				_ = _61_v4
				_61_v4 = _dafny.UnicodeSeqOfUtf8Bytes("bcuys")
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_58_v1), 0))
				_ = _index1
				(_58_v1).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_61_v4, _61_v4), _dafny.Companion_Sequence_.Concatenate(_61_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(829))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}(func(_62_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))))).Cardinality()), (_index1).Int())
				r3 = _59_v2
				r0 = _56_v0
				r0 = _56_v0
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _63_v5 _dafny.Sequence
	_ = _63_v5
	_63_v5 = _dafny.UnicodeSeqOfUtf8Bytes("btljib")
	var _64_v6 _dafny.Int
	_ = _64_v6
	_64_v6 = _dafny.IntOfInt64(-243)
	var _65_v7 _dafny.Map
	_ = _65_v7
	_65_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v6, !(_56_v0))
	var _66_v8 D8
	_ = _66_v8
	_66_v8 = Companion_D8_.Create_DC19_(_65_v7)
	var _pat_let_tv0 = _63_v5
	_ = _pat_let_tv0
	var _pat_let_tv1 = _65_v7
	_ = _pat_let_tv1
	_63_v5 = func(_source3 D8) _dafny.Sequence {
		if _source3.Is_DC20() {
			var _67___mcc_h0 _dafny.CodePoint = _source3.Get_().(D8_DC20).Cf43
			_ = _67___mcc_h0
			var _68___mcc_h1 bool = _source3.Get_().(D8_DC20).Cf44
			_ = _68___mcc_h1
			var _69___mcc_h2 bool = _source3.Get_().(D8_DC20).Cf45
			_ = _69___mcc_h2
			var _70___mcc_h3 _dafny.Map = _source3.Get_().(D8_DC20).Cf46
			_ = _70___mcc_h3
			var _71___mcc_h4 bool = _source3.Get_().(D8_DC20).Cf47
			_ = _71___mcc_h4
			var _72_cf47 bool = _71___mcc_h4
			_ = _72_cf47
			var _73_cf46 _dafny.Map = _70___mcc_h3
			_ = _73_cf46
			var _74_cf45 bool = _69___mcc_h2
			_ = _74_cf45
			var _75_cf44 bool = _68___mcc_h1
			_ = _75_cf44
			var _76_cf43 _dafny.CodePoint = _67___mcc_h0
			_ = _76_cf43
			return _dafny.UnicodeSeqOfUtf8Bytes("iuyjsjqt")
		} else {
			var _77___mcc_h5 _dafny.Map = _source3.Get_().(D8_DC19).Cf42
			_ = _77___mcc_h5
			var _78_cf42 _dafny.Map = _77___mcc_h5
			_ = _78_cf42
			return _pat_let_tv0
		}
	}(func(_pat_let0_0 D8) D8 {
		return func(_79_dt__update__tmp_h0 D8) D8 {
			return func(_pat_let1_0 _dafny.Map) D8 {
				return func(_80_dt__update_hcf42_h0 _dafny.Map) D8 {
					return Companion_D8_.Create_DC19_(_80_dt__update_hcf42_h0)
				}(_pat_let1_0)
			}(_pat_let_tv1)
		}(_pat_let0_0)
	}(_66_v8))
	var _81_v9 *C0
	_ = _81_v9
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__()
	_81_v9 = _nw1
	var _82_v10 _dafny.MultiSet
	_ = _82_v10
	_82_v10 = _dafny.MultiSetOf(_56_v0, _56_v0)
	r1 = ((_dafny.MultiSetOf(!(_56_v0))).Intersection((_dafny.MultiSetFromSeq(p1)).Update(_56_v0, Companion_Default___.Abs(_64_v6)))).Union((_82_v10).Union(_dafny.MultiSetOf(!(_56_v0))))
	r3 = (_64_v6).Plus((_dafny.Zero).Minus(_64_v6))
	r3 = (Companion_Default___.Fm35(globalState)).Dtor_cf38()
	r0 = (_64_v6).Cmp(_64_v6) < 0
	r1 = _82_v10
	var _83_v11 D3
	_ = _83_v11
	_83_v11 = Companion_D3_.Create_DC7_(_63_v5)
	var _pat_let_tv2 = p0
	_ = _pat_let_tv2
	var _pat_let_tv3 = _64_v6
	_ = _pat_let_tv3
	var _pat_let_tv4 = p0
	_ = _pat_let_tv4
	var _pat_let_tv5 = p1
	_ = _pat_let_tv5
	var _pat_let_tv6 = _64_v6
	_ = _pat_let_tv6
	var _pat_let_tv7 = p1
	_ = _pat_let_tv7
	var _pat_let_tv8 = _56_v0
	_ = _pat_let_tv8
	var _pat_let_tv9 = _56_v0
	_ = _pat_let_tv9
	var _pat_let_tv10 = _56_v0
	_ = _pat_let_tv10
	var _pat_let_tv11 = _56_v0
	_ = _pat_let_tv11
	var _pat_let_tv12 = _56_v0
	_ = _pat_let_tv12
	r2 = func(_source4 D3) _dafny.Set {
		if _source4.Is_DC8() {
			var _84___mcc_h6 _dafny.CodePoint = _source4.Get_().(D3_DC8).Cf17
			_ = _84___mcc_h6
			var _85___mcc_h7 bool = _source4.Get_().(D3_DC8).Cf18
			_ = _85___mcc_h7
			var _86___mcc_h8 bool = _source4.Get_().(D3_DC8).Cf19
			_ = _86___mcc_h8
			var _87___mcc_h9 _dafny.Sequence = _source4.Get_().(D3_DC8).Cf20
			_ = _87___mcc_h9
			var _88_cf20 _dafny.Sequence = _87___mcc_h9
			_ = _88_cf20
			var _89_cf19 bool = _86___mcc_h8
			_ = _89_cf19
			var _90_cf18 bool = _85___mcc_h7
			_ = _90_cf18
			var _91_cf17 _dafny.CodePoint = _84___mcc_h6
			_ = _91_cf17
			return (_dafny.SetOf(!(_90_cf18), _90_cf18, (_pat_let_tv2).Select((Companion_Default___.SafeIndex(_pat_let_tv3, _dafny.IntOfUint32((_pat_let_tv4).Cardinality()))).Uint32()).(bool))).Union(_dafny.SetOf((_pat_let_tv5).Select((Companion_Default___.SafeIndex(_pat_let_tv6, _dafny.IntOfUint32((_pat_let_tv7).Cardinality()))).Uint32()).(bool), _90_cf18))
		} else if _source4.Is_DC7() {
			var _92___mcc_h10 _dafny.Sequence = _source4.Get_().(D3_DC7).Cf16
			_ = _92___mcc_h10
			var _93_cf16 _dafny.Sequence = _92___mcc_h10
			_ = _93_cf16
			return (_dafny.SetOf(_pat_let_tv8)).Difference(_dafny.SetOf(_pat_let_tv9))
		} else {
			var _94___mcc_h11 D3 = _source4.Get_().(D3_DC9).Cf21
			_ = _94___mcc_h11
			var _95_cf21 D3 = _94___mcc_h11
			_ = _95_cf21
			return (_dafny.SetOf(_pat_let_tv10)).Intersection(_dafny.SetOf(_pat_let_tv11, _pat_let_tv12))
		}
	}(_83_v11)
	var _96_v12 _dafny.CodePoint
	_ = _96_v12
	_96_v12 = _dafny.CodePoint('w')
	var _97_v13 _dafny.Map
	_ = _97_v13
	_97_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v6, _64_v6)
	var _98_v14 _dafny.Sequence
	_ = _98_v14
	_98_v14 = _dafny.SeqOf((_97_v13).Cardinality(), _64_v6)
	var _99_v15 _dafny.MultiSet
	_ = _99_v15
	_99_v15 = _dafny.MultiSetOf(_64_v6)
	r3 = (func() _dafny.Int {
		if (_64_v6).Cmp(_64_v6) <= 0 {
			return Companion_Default___.Fm19(_96_v12, _56_v0, globalState)
		}
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_98_v14, _dafny.SeqOf((_99_v15).Cardinality()))).Cardinality())
	})()
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _100_v0 bool
	_ = _100_v0
	_100_v0 = false
	var _101_v1 _dafny.Array
	_ = _101_v1
	var _nwElement0_0 bool = _100_v0
	_ = _nwElement0_0
	var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(26))
	_ = _nw2
	(_nw2).ArraySet1(_nwElement0_0, 0)
	(_nw2).ArraySet1(_100_v0, 1)
	(_nw2).ArraySet1(_100_v0, 2)
	(_nw2).ArraySet1(_100_v0, 3)
	(_nw2).ArraySet1(_100_v0, 4)
	(_nw2).ArraySet1(_100_v0, 5)
	(_nw2).ArraySet1(_100_v0, 6)
	(_nw2).ArraySet1(_100_v0, 7)
	(_nw2).ArraySet1(_100_v0, 8)
	(_nw2).ArraySet1(_100_v0, 9)
	(_nw2).ArraySet1(_100_v0, 10)
	(_nw2).ArraySet1(false, 11)
	(_nw2).ArraySet1(true, 12)
	(_nw2).ArraySet1(!(_100_v0), 13)
	(_nw2).ArraySet1(!(_100_v0), 14)
	(_nw2).ArraySet1(_100_v0, 15)
	(_nw2).ArraySet1(false, 16)
	(_nw2).ArraySet1(_100_v0, 17)
	(_nw2).ArraySet1(false, 18)
	(_nw2).ArraySet1(_100_v0, 19)
	(_nw2).ArraySet1(_100_v0, 20)
	(_nw2).ArraySet1(_100_v0, 21)
	(_nw2).ArraySet1(_100_v0, 22)
	(_nw2).ArraySet1(_100_v0, 23)
	(_nw2).ArraySet1(_100_v0, 24)
	(_nw2).ArraySet1(_100_v0, 25)
	_101_v1 = _nw2
	var _102_v2 _dafny.Int
	_ = _102_v2
	_102_v2 = _dafny.IntOfInt64(599)
	var _103_v3 _dafny.MultiSet
	_ = _103_v3
	_103_v3 = _dafny.MultiSetOf(_102_v2)
	var _104_v4 _dafny.Sequence
	_ = _104_v4
	_104_v4 = _dafny.SeqOf(_102_v2)
	var _105_globalState *GlobalState
	_ = _105_globalState
	var _nw3 *GlobalState = New_GlobalState_()
	_ = _nw3
	_nw3.Ctor__(false, _dafny.IntOfInt64(38), _dafny.IntOfInt64(586), false, _101_v1, _101_v1, _103_v3, _104_v4, true, _dafny.IntOfInt64(249))
	_105_globalState = _nw3
	var _106_v5 _dafny.Map
	_ = _106_v5
	_106_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-550), _dafny.IntOfUint32((_104_v4).Cardinality())))
	_106_v5 = _106_v5
	var _107_v6 _dafny.Sequence
	_ = _107_v6
	_107_v6 = _dafny.SeqOf(_100_v0, _100_v0)
	var _108_v7 bool
	_ = _108_v7
	var _109_v8 _dafny.MultiSet
	_ = _109_v8
	var _110_v9 _dafny.Set
	_ = _110_v9
	var _111_v10 _dafny.Int
	_ = _111_v10
	var _out0 bool
	_ = _out0
	var _out1 _dafny.MultiSet
	_ = _out1
	var _out2 _dafny.Set
	_ = _out2
	var _out3 _dafny.Int
	_ = _out3
	_out0, _out1, _out2, _out3 = Companion_Default___.M0(_107_v6, _107_v6, _105_globalState)
	_108_v7 = _out0
	_109_v8 = _out1
	_110_v9 = _out2
	_111_v10 = _out3
	_111_v10 = _102_v2
	var _112_v11 _dafny.Set
	_ = _112_v11
	_112_v11 = _dafny.SetOf(_102_v2, _111_v10, _111_v10, _111_v10, _102_v2)
	_112_v11 = _112_v11
	var _113_v12 _dafny.Array
	_ = _113_v12
	var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
	_ = _nw4
	_113_v12 = _nw4
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
	_ = _index2
	(_113_v12).ArraySet1((_dafny.Zero).Minus((_102_v2).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_104_v4).Cardinality())))), (_index2).Int())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
	_ = _index3
	(_101_v1).ArraySet1(!(_108_v7) || (_108_v7), (_index3).Int())
	var _114_v13 _dafny.Sequence
	_ = _114_v13
	_114_v13 = _dafny.UnicodeSeqOfUtf8Bytes("fe")
	var _115_v14 _dafny.Map
	_ = _115_v14
	_115_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v13, _102_v2)
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
	_ = _index4
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
	_ = _index5
	var _rhs0 bool = !(_112_v11).Equals(_112_v11)
	_ = _rhs0
	var _rhs1 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex(_102_v2, _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _dafny.CodePoint('b')), _114_v13)).Cardinality())
	_ = _rhs1
	var _rhs2 _dafny.Int = _111_v10
	_ = _rhs2
	var _rhs3 bool = _100_v0
	_ = _rhs3
	var _rhs4 _dafny.Map = (func() _dafny.Map {
		if _108_v7 {
			return _115_v14
		}
		return Companion_Default___.Fm0(_105_globalState)
	})()
	_ = _rhs4
	var _lhs0 _dafny.Array = _113_v12
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
	_ = _lhs1
	var _lhs2 _dafny.Array = _101_v1
	_ = _lhs2
	var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
	_ = _lhs3
	_100_v0 = _rhs0
	_102_v2 = _rhs1
	(_lhs0).ArraySet1(_rhs2, (_lhs1).Int())
	(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
	_115_v14 = _rhs4
	if (_112_v11).IsSubsetOf((_112_v11).Intersection(_112_v11)) {
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
		_ = _index6
		(_101_v1).ArraySet1((((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(496))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_116_v12 _dafny.Array) func(_dafny.Int) _dafny.Int {
			return func(_117_i0 _dafny.Int) _dafny.Int {
				return (_116_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_116_v12), 0))).Int()).(_dafny.Int)
			}
		})(_113_v12))))).Update(_102_v2, Companion_Default___.Abs(_111_v10))).Update((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_dafny.Zero).Minus(_102_v2)))).IsProperSubsetOf(_103_v3), (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
		_ = _index7
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
		_ = _index8
		var _rhs5 _dafny.Int = (Companion_Default___.SafeDivisionInt((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_114_v13).Cardinality()))).Times((_dafny.Zero).Minus((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int)))
		_ = _rhs5
		var _rhs6 _dafny.Int = Companion_Default___.SafeModuloInt((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), (_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int))
		_ = _rhs6
		var _rhs7 _dafny.Sequence = _114_v13
		_ = _rhs7
		var _lhs4 _dafny.Array = _113_v12
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
		_ = _lhs5
		var _lhs6 _dafny.Array = _113_v12
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
		_ = _lhs7
		(_lhs4).ArraySet1(_rhs5, (_lhs5).Int())
		(_lhs6).ArraySet1(_rhs6, (_lhs7).Int())
		_114_v13 = _rhs7
		var _118_v15 _dafny.Map
		_ = _118_v15
		_118_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v7, (_101_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))).Int()).(bool))
		var _119_v16 _dafny.Map
		_ = _119_v16
		_119_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_104_v4).Select((Companion_Default___.SafeIndex((_118_v15).Cardinality(), _dafny.IntOfUint32((_104_v4).Cardinality()))).Uint32()).(_dafny.Int), _102_v2)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
		_ = _index9
		var _rhs8 _dafny.MultiSet = _103_v3
		_ = _rhs8
		var _rhs9 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_102_v2).Times((_119_v16).Cardinality()), _dafny.IntOfInt64(373)))
		_ = _rhs9
		var _rhs10 _dafny.Sequence = _114_v13
		_ = _rhs10
		var _rhs11 _dafny.Int = _dafny.IntOfInt64(192)
		_ = _rhs11
		var _lhs8 _dafny.Array = _113_v12
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
		_ = _lhs9
		_103_v3 = _rhs8
		(_lhs8).ArraySet1(_rhs9, (_lhs9).Int())
		_114_v13 = _rhs10
		_102_v2 = _rhs11
		_114_v13 = _114_v13
		var _120_v18 D0
		_ = _120_v18
		_120_v18 = Companion_D0_.Create_DC0_(_dafny.CodePoint('v'))
		var _121_v19 _dafny.CodePoint
		_ = _121_v19
		_121_v19 = _dafny.CodePoint('x')
		var _122_v20 _dafny.Map
		_ = _122_v20
		_122_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v19, _114_v13)
		var _123_v22 _dafny.Set
		_ = _123_v22
		_123_v22 = _dafny.SetOf(_121_v19)
		var _124_v24 _dafny.Map
		_ = _124_v24
		_124_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v19, _118_v15)
		var _125_v26 _dafny.Array
		_ = _125_v26
		var _nwElement0_1 _dafny.Map = func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_114_v13).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _126_v17 _dafny.CodePoint
				_126_v17 = interface{}(_compr_6).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_114_v13, _126_v17) {
					_coll6.Add(_126_v17, _dafny.SeqOf(_dafny.CodePoint('p')))
				}
			}
			return _coll6.ToMap()
		}()
		_ = _nwElement0_1
		var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(29))
		_ = _nw5
		(_nw5).ArraySet1(_nwElement0_1, 0)
		(_nw5).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_105_globalState), _114_v13)).Update((_120_v18).Dtor_cf0(), Companion_Default___.Fm2(_100_v0, (_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.CodePoint('q'), (_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _105_globalState)), 1)
		(_nw5).ArraySet1(_122_v20, 2)
		(_nw5).ArraySet1(_122_v20, 3)
		(_nw5).ArraySet1(Companion_Default___.Fm3(_100_v0, _105_globalState), 4)
		(_nw5).ArraySet1(_122_v20, 5)
		(_nw5).ArraySet1(_122_v20, 6)
		(_nw5).ArraySet1(_122_v20, 7)
		(_nw5).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v19, _dafny.SeqOf(Companion_Default___.Fm1(_105_globalState))), 8)
		(_nw5).ArraySet1((_122_v20).Merge(_122_v20), 9)
		(_nw5).ArraySet1(func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_123_v22).Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _127_v21 _dafny.CodePoint
				_127_v21 = interface{}(_compr_7).(_dafny.CodePoint)
				if (_123_v22).Contains(_127_v21) {
					_coll7.Add(_127_v21, _dafny.SeqOf(_121_v19))
				}
			}
			return _coll7.ToMap()
		}(), 10)
		(_nw5).ArraySet1(Companion_Default___.Fm3(_100_v0, _105_globalState), 11)
		(_nw5).ArraySet1(_122_v20, 12)
		(_nw5).ArraySet1(_122_v20, 13)
		(_nw5).ArraySet1((_122_v20).Update(_121_v19, _114_v13), 14)
		(_nw5).ArraySet1(func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_124_v24).Keys().Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _128_v23 _dafny.CodePoint
				_128_v23 = interface{}(_compr_8).(_dafny.CodePoint)
				if (_124_v24).Contains(_128_v23) {
					_coll8.Add(_128_v23, _114_v13)
				}
			}
			return _coll8.ToMap()
		}(), 15)
		(_nw5).ArraySet1(_122_v20, 16)
		(_nw5).ArraySet1(_122_v20, 17)
		(_nw5).ArraySet1(_122_v20, 18)
		(_nw5).ArraySet1(_122_v20, 19)
		(_nw5).ArraySet1(_122_v20, 20)
		(_nw5).ArraySet1((_122_v20).Merge(_122_v20), 21)
		(_nw5).ArraySet1((func() _dafny.Map {
			if _100_v0 {
				return _122_v20
			}
			return func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate((_123_v22).Elements()); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _129_v25 _dafny.CodePoint
					_129_v25 = interface{}(_compr_9).(_dafny.CodePoint)
					if (_123_v22).Contains(_129_v25) {
						_coll9.Add(_129_v25, _114_v13)
					}
				}
				return _coll9.ToMap()
			}()
		})(), 22)
		(_nw5).ArraySet1((_122_v20).Update(_121_v19, _114_v13), 23)
		(_nw5).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v19, _114_v13), 24)
		(_nw5).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_120_v18).Dtor_cf0(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_130_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_131_i1 _dafny.Int) _dafny.CodePoint {
				return _130_v19
			}
		})(_121_v19)))), 25)
		(_nw5).ArraySet1(Companion_Default___.Fm3((_101_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))).Int()).(bool), _105_globalState), 26)
		(_nw5).ArraySet1(_122_v20, 27)
		(_nw5).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v19, _dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex(_102_v2, _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _121_v19)), 28)
		_125_v26 = _nw5
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_125_v26), 0))
		_ = _index10
		(_125_v26).ArraySet1(_122_v20, (_index10).Int())
		var _132_v27 _dafny.Sequence
		_ = _132_v27
		_132_v27 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v19, _dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _121_v19)), _122_v20, _122_v20, _122_v20, (_122_v20).Merge(_122_v20))
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_125_v26), 0))
		_ = _index11
		(_125_v26).ArraySet1((_132_v27).Select((Companion_Default___.SafeIndex(_102_v2, _dafny.IntOfUint32((_132_v27).Cardinality()))).Uint32()).(_dafny.Map), (_index11).Int())
	} else {
		var _133_v28 _dafny.Map
		_ = _133_v28
		_133_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v1, (_101_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))).Int()).(bool))
		_133_v28 = (_133_v28).Update(_101_v1, !(!(_100_v0) || (_108_v7)))
		_109_v8 = _109_v8
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
		_ = _index12
		(_101_v1).ArraySet1((_109_v8).IsProperSubsetOf(_109_v8), (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
		_ = _index13
		(_101_v1).ArraySet1(!(Companion_Default___.Fm4(_112_v11, _105_globalState)), (_index13).Int())
		var _134_v29 D6
		_ = _134_v29
		_134_v29 = Companion_D6_.Create_DC17_(_102_v2, _114_v13, _100_v0)
		var _135_v30 *C1
		_ = _135_v30
		var _nw6 *C1 = New_C1_()
		_ = _nw6
		_nw6.Ctor__((_dafny.IntOfInt64(-225)).Cmp((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int)) <= 0, (_134_v29).Dtor_cf40(), _109_v8)
		_135_v30 = _nw6
	}
	var _136_v31 bool
	_ = _136_v31
	var _137_v32 _dafny.MultiSet
	_ = _137_v32
	var _138_v33 _dafny.Set
	_ = _138_v33
	var _139_v34 _dafny.Int
	_ = _139_v34
	var _out4 bool
	_ = _out4
	var _out5 _dafny.MultiSet
	_ = _out5
	var _out6 _dafny.Set
	_ = _out6
	var _out7 _dafny.Int
	_ = _out7
	_out4, _out5, _out6, _out7 = Companion_Default___.M0(_107_v6, (func() _dafny.Sequence {
		if _108_v7 {
			return _107_v6
		}
		return _107_v6
	})(), _105_globalState)
	_136_v31 = _out4
	_137_v32 = _out5
	_138_v33 = _out6
	_139_v34 = _out7
	var _140_v35 _dafny.Array
	_ = _140_v35
	var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
	_ = _nw7
	_140_v35 = _nw7
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_140_v35), 0))
	_ = _index14
	(_140_v35).ArraySet1(_113_v12, (_index14).Int())
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_140_v35), 0))
	_ = _index15
	(_140_v35).ArraySet1(_113_v12, (_index15).Int())
	var _141_v36 D9
	_ = _141_v36
	_141_v36 = Companion_D9_.Create_DC23_(_139_v34)
	var _142_v37 _dafny.Map
	_ = _142_v37
	_142_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _102_v2)
	var _hi0 _dafny.Int = (func() _dafny.Int {
		if (_142_v37).Contains(_111_v10) {
			return (_142_v37).Get(_111_v10).(_dafny.Int)
		}
		return _111_v10
	})()
	_ = _hi0
	for _143_i2 := (_141_v36).Dtor_cf52(); _143_i2.Cmp(_hi0) < 0; _143_i2 = _143_i2.Plus(_dafny.One) {
		var _144_v38 *C5
		_ = _144_v38
		var _nw8 *C5 = New_C5_()
		_ = _nw8
		_nw8.Ctor__(_100_v0, _109_v8)
		_144_v38 = _nw8
		var _145_v39 _dafny.Map
		_ = _145_v39
		_145_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v38, _108_v7)
		var _146_v40 T0
		_ = _146_v40
		var _nw9 *C2 = New_C2_()
		_ = _nw9
		_nw9.Ctor__(!((_144_v38).F11()), _143_i2, _109_v8)
		_146_v40 = _nw9
		var _147_v41 _dafny.Array
		_ = _147_v41
		var _nwElement0_2 T0 = _146_v40
		_ = _nwElement0_2
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(8))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_2, 0)
		(_nw10).ArraySet1(_146_v40, 1)
		(_nw10).ArraySet1(_146_v40, 2)
		(_nw10).ArraySet1(_146_v40, 3)
		(_nw10).ArraySet1(_146_v40, 4)
		(_nw10).ArraySet1(_146_v40, 5)
		(_nw10).ArraySet1(_146_v40, 6)
		(_nw10).ArraySet1(_146_v40, 7)
		_147_v41 = _nw10
		var _148_v42 _dafny.Map
		_ = _148_v42
		_148_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_147_v41, _108_v7)
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
		_ = _index16
		(_101_v1).ArraySet1(!((func() bool {
			if (_145_v39).Contains((_144_v38)) {
				return (_145_v39).Get((_144_v38)).(bool)
			}
			return (func() bool {
				if (_148_v42).Contains(_147_v41) {
					return (_148_v42).Get(_147_v41).(bool)
				}
				return true
			})()
		})()) || ((_144_v38).F11()), (_index16).Int())
		var _149_v43 bool
		_ = _149_v43
		var _150_v44 _dafny.MultiSet
		_ = _150_v44
		var _151_v45 _dafny.Set
		_ = _151_v45
		var _152_v46 _dafny.Int
		_ = _152_v46
		var _out8 bool
		_ = _out8
		var _out9 _dafny.MultiSet
		_ = _out9
		var _out10 _dafny.Set
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out8, _out9, _out10, _out11 = Companion_Default___.M0(_dafny.SeqOf((_101_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))).Int()).(bool)), _107_v6, _105_globalState)
		_149_v43 = _out8
		_150_v44 = _out9
		_151_v45 = _out10
		_152_v46 = _out11
		var _153_v47 _dafny.Sequence
		_ = _153_v47
		_153_v47 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_107_v6))
		var _154_v48 _dafny.CodePoint
		_ = _154_v48
		_154_v48 = _dafny.CodePoint('p')
		var _155_v49 _dafny.Map
		_ = _155_v49
		_155_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_144_v38).Fm5(_151_v45, _dafny.IntOfUint32((_153_v47).Cardinality()), _152_v46, _105_globalState), _143_i2, _154_v48, _152_v46, _105_globalState), _149_v43)
		var _156_v50 _dafny.Map
		_ = _156_v50
		_156_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v43, _108_v7)
		_155_v49 = (_155_v49).Update(Companion_Default___.Fm2((func() bool {
			if (_156_v50).Contains((_144_v38).F11()) {
				return (_156_v50).Get((_144_v38).F11()).(bool)
			}
			return _149_v43
		})(), _102_v2, _154_v48, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ivq")).Cardinality()), _105_globalState), _136_v31)
		var _157_v51 *C0
		_ = _157_v51
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__()
		_157_v51 = _nw11
		_157_v51 = _157_v51
	}
	var _158_v52 _dafny.CodePoint
	_ = _158_v52
	_158_v52 = _dafny.CodePoint('f')
	var _159_v53 _dafny.Sequence
	_ = _159_v53
	_159_v53 = _dafny.SeqOf(_114_v13, _114_v13, Companion_Default___.Fm2(_100_v0, _111_v10, _158_v52, (_142_v37).Cardinality(), _105_globalState), _114_v13)
	var _160_v54 _dafny.Array
	_ = _160_v54
	var _nwElement0_3 _dafny.Sequence = (_159_v53).Select((Companion_Default___.SafeIndex(_111_v10, _dafny.IntOfUint32((_159_v53).Cardinality()))).Uint32()).(_dafny.Sequence)
	_ = _nwElement0_3
	var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(24))
	_ = _nw12
	(_nw12).ArraySet1(_nwElement0_3, 0)
	(_nw12).ArraySet1(_114_v13, 1)
	(_nw12).ArraySet1(_114_v13, 2)
	(_nw12).ArraySet1(_114_v13, 3)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_114_v13, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _158_v52), (Companion_Default___.SafeIndex(_111_v10, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _158_v52)).Cardinality()))).Uint32(), (_114_v13).Select((Companion_Default___.SafeIndex(_102_v2, _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32()).(_dafny.CodePoint))), (Companion_Default___.SafeIndex((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_114_v13, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _158_v52), (Companion_Default___.SafeIndex(_111_v10, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _158_v52)).Cardinality()))).Uint32(), (_114_v13).Select((Companion_Default___.SafeIndex(_102_v2, _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32()).(_dafny.CodePoint)))).Cardinality()))).Uint32(), Companion_Default___.Fm1(_105_globalState)), 4)
	(_nw12).ArraySet1(_114_v13, 5)
	(_nw12).ArraySet1(_114_v13, 6)
	(_nw12).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}((func(_161_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_162_i3 _dafny.Int) _dafny.CodePoint {
			return _161_v52
		}
	})(_158_v52))), 7)
	(_nw12).ArraySet1(Companion_Default___.Fm2(_108_v7, _111_v10, _158_v52, _139_v34, _105_globalState), 8)
	(_nw12).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("muykribt"), 9)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_114_v13, _dafny.UnicodeSeqOfUtf8Bytes("atwbgc")), 10)
	(_nw12).ArraySet1(_114_v13, 11)
	(_nw12).ArraySet1(_114_v13, 12)
	(_nw12).ArraySet1(_114_v13, 13)
	(_nw12).ArraySet1(_114_v13, 14)
	(_nw12).ArraySet1(_114_v13, 15)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_114_v13, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("lx"), (Companion_Default___.SafeIndex(_111_v10, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lx")).Cardinality()))).Uint32(), _158_v52)), 16)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex(_111_v10, _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _158_v52), 17)
	(_nw12).ArraySet1(_114_v13, 18)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_114_v13, _114_v13), 19)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_114_v13, _dafny.UnicodeSeqOfUtf8Bytes("yyix")), 20)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_114_v13, _dafny.UnicodeSeqOfUtf8Bytes("hkxekp")), 21)
	(_nw12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_114_v13, Companion_Default___.Fm2((_101_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))).Int()).(bool), _dafny.IntOfInt64(228), _158_v52, _102_v2, _105_globalState)), 22)
	(_nw12).ArraySet1((func() _dafny.Sequence {
		if _100_v0 {
			return _114_v13
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-159))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_163_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_164_i4 _dafny.Int) _dafny.CodePoint {
				return _163_v52
			}
		})(_158_v52)))
	})(), 23)
	_160_v54 = _nw12
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_160_v54), 0))
	_ = _index17
	(_160_v54).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hythm"), (_index17).Int())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_160_v54), 0))
	_ = _index18
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
	_ = _index19
	var _rhs12 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_114_v13, Companion_Default___.Fm2(_100_v0, (Companion_D0_.Create_DC1_(_100_v0, (_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int))).Dtor_cf2(), _158_v52, _dafny.IntOfInt64(432), _105_globalState))
	_ = _rhs12
	var _rhs13 _dafny.Int = (_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int)
	_ = _rhs13
	var _lhs10 _dafny.Array = _160_v54
	_ = _lhs10
	var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_160_v54), 0))
	_ = _lhs11
	var _lhs12 _dafny.Array = _113_v12
	_ = _lhs12
	var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))
	_ = _lhs13
	(_lhs10).ArraySet1(_rhs12, (_lhs11).Int())
	(_lhs12).ArraySet1(_rhs13, (_lhs13).Int())
	var _hi1 _dafny.Int = (_dafny.Zero).Minus((_102_v2).Times(_111_v10))
	_ = _hi1
	for _165_i5 := (_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int); _165_i5.Cmp(_hi1) < 0; _165_i5 = _165_i5.Plus(_dafny.One) {
		var _166_v55 _dafny.Array
		_ = _166_v55
		var _nw13 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
		_ = _nw13
		_166_v55 = _nw13
		var _167_v56 _dafny.Map
		_ = _167_v56
		_167_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(366), _166_v55)
		_167_v56 = _167_v56
		var _168_v57 _dafny.Array
		_ = _168_v57
		var _nw14 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
		_ = _nw14
		_168_v57 = _nw14
		var _nw15 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
		_ = _nw15
		_168_v57 = _nw15
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))
		_ = _index20
		(_101_v1).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_107_v6, _dafny.Companion_Sequence_.Update(_107_v6, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.IntOfUint32((_107_v6).Cardinality()))).Uint32(), _136_v31)), _107_v6), (_index20).Int())
		var _169_v58 _dafny.Map
		_ = _169_v58
		_169_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_i5, (_101_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_101_v1), 0))).Int()).(bool))
		_102_v2 = (_111_v10).Minus((_169_v58).Cardinality())
	}
	_106_v5 = (_106_v5).Update(!_dafny.Companion_Sequence_.Contains(_107_v6, _108_v7), (_dafny.Zero).Minus((_113_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_113_v12), 0))).Int()).(_dafny.Int)))
	_112_v11 = _112_v11
	var _170_v59 T1
	_ = _170_v59
	var _nw16 *C3 = New_C3_()
	_ = _nw16
	_nw16.Ctor__(_136_v31)
	_170_v59 = _nw16
	var _171_v60 _dafny.Sequence
	_ = _171_v60
	_171_v60 = _dafny.SeqOf(_170_v59, _170_v59, _170_v59, _170_v59, _170_v59)
	_171_v60 = _171_v60
	var _172_v61 _dafny.MultiSet
	_ = _172_v61
	_172_v61 = _137_v32
	var _pat_let_tv13 = _113_v12
	_ = _pat_let_tv13
	var _pat_let_tv14 = _113_v12
	_ = _pat_let_tv14
	_108_v7 = func(_source5 _dafny.MultiSet) bool {
		var _173___mcc_h0 _dafny.MultiSet = _source5
		_ = _173___mcc_h0
		var _174_cf41 _dafny.MultiSet = _173___mcc_h0
		_ = _174_cf41
		return (_dafny.IntOfInt64(381)).Cmp((_pat_let_tv14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_pat_let_tv13), 0))).Int()).(_dafny.Int)) != 0
	}(_172_v61)
	var _175_v62 bool
	_ = _175_v62
	var _176_v63 _dafny.MultiSet
	_ = _176_v63
	var _177_v64 _dafny.Set
	_ = _177_v64
	var _178_v65 _dafny.Int
	_ = _178_v65
	var _out12 bool
	_ = _out12
	var _out13 _dafny.MultiSet
	_ = _out13
	var _out14 _dafny.Set
	_ = _out14
	var _out15 _dafny.Int
	_ = _out15
	_out12, _out13, _out14, _out15 = Companion_Default___.M0(_107_v6, _107_v6, _105_globalState)
	_175_v62 = _out12
	_176_v63 = _out13
	_177_v64 = _out14
	_178_v65 = _out15
	_dafny.Print(_100_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v1).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_v3).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(599))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_104_v4, _dafny.SeqOf(_dafny.IntOfInt64(599))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F5).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F6()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(599))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_105_globalState).F7(), _dafny.SeqOf(_dafny.IntOfInt64(599))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.Zero).UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_107_v6, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_108_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v8).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_110_v9).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_111_v10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_112_v11).Equals(_dafny.SetOf(_dafny.IntOfInt64(599))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v12).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_v13.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_115_v14).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ibqqjdrauoj"), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_136_v31)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v32).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v33).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_139_v34)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_140_v35).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v36).Dtor_cf52())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_142_v37).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(192))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_158_v52)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_159_v53, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fe"), _dafny.UnicodeSeqOfUtf8Bytes("fe"), _dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('i'), _dafny.CodePoint('n'), _dafny.CodePoint('o'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('d'), _dafny.CodePoint('d')), _dafny.UnicodeSeqOfUtf8Bytes("fe"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_160_v54).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_160_v54).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('i'), _dafny.CodePoint('n'), _dafny.CodePoint('o'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('d'), _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v54).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_160_v54).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'), _dafny.CodePoint('f'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_171_v60).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v61).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_v62)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_176_v63).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_177_v64).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_178_v65)
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
	Cf2 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.CodePoint
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.CodePoint) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf3 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf3 D0) D0 {
	return D0{D0_DC2{Cf3}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf0() _dafny.CodePoint {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf5 bool
	Cf6 bool
	Cf7 _dafny.Int
	Cf8 _dafny.Sequence
	Cf9 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 bool, Cf6 bool, Cf7 _dafny.Int, Cf8 _dafny.Sequence, Cf9 _dafny.Int) D1 {
	return D1{D1_DC4{Cf5, Cf6, Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf10 _dafny.Int
	Cf11 bool
	Cf12 _dafny.Int
	Cf13 bool
	Cf14 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf10 _dafny.Int, Cf11 bool, Cf12 _dafny.Int, Cf13 bool, Cf14 _dafny.Int) D1 {
	return D1{D1_DC5{Cf10, Cf11, Cf12, Cf13, Cf14}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf4 _dafny.Array
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.Array) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, false, _dafny.Zero, _dafny.EmptySeq, _dafny.Zero)
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf11() bool {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) Dtor_cf13() bool {
	return _this.Get_().(D1_DC5).Cf13
}

func (_this D1) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf14
}

func (_this D1) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Equals(data2.Cf8) && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4 == data2.Cf4
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
	Cf15 _dafny.MultiSet
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf15 _dafny.MultiSet) D2 {
	return D2{D2_DC6{Cf15}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D2) Dtor_cf15() _dafny.MultiSet {
	return _this.Get_().(D2_DC6).Cf15
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf17 _dafny.CodePoint
	Cf18 bool
	Cf19 bool
	Cf20 _dafny.Sequence
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf17 _dafny.CodePoint, Cf18 bool, Cf19 bool, Cf20 _dafny.Sequence) D3 {
	return D3{D3_DC8{Cf17, Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf16 _dafny.Sequence
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf16 _dafny.Sequence) D3 {
	return D3{D3_DC7{Cf16}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC9 struct {
	Cf21 D3
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf21 D3) D3 {
	return D3{D3_DC9{Cf21}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.CodePoint('D'), false, false, _dafny.EmptySeq)
}

func (_this D3) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D3_DC8).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC8).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC8).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf20
}

func (_this D3) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D3_DC7).Cf16
}

func (_this D3) Dtor_cf21() D3 {
	return _this.Get_().(D3_DC9).Cf21
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + data.Cf20.VerbatimString(true) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + data.Cf16.VerbatimString(true) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19 && data1.Cf20.Equals(data2.Cf20)
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf16.Equals(data2.Cf16)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf21.Equals(data2.Cf21)
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

type D4_DC11 struct {
	Cf23 D0
	Cf24 _dafny.Array
	Cf25 _dafny.Int
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf23 D0, Cf24 _dafny.Array, Cf25 _dafny.Int) D4 {
	return D4{D4_DC11{Cf23, Cf24, Cf25}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf22 *C0
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf22 *C0) D4 {
	return D4{D4_DC10{Cf22}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(Companion_D0_.Default(), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D4) Dtor_cf23() D0 {
	return _this.Get_().(D4_DC11).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf25
}

func (_this D4) Dtor_cf22() *C0 {
	return _this.Get_().(D4_DC10).Cf22
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf23.Equals(data2.Cf23) && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf22 == data2.Cf22
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
	Cf27 _dafny.Int
	Cf28 _dafny.Array
	Cf29 _dafny.Array
	Cf30 _dafny.Int
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf27 _dafny.Int, Cf28 _dafny.Array, Cf29 _dafny.Array, Cf30 _dafny.Int) D5 {
	return D5{D5_DC13{Cf27, Cf28, Cf29, Cf30}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
	Cf31 _dafny.Array
	Cf32 _dafny.Int
	Cf33 _dafny.Int
	Cf34 _dafny.Map
	Cf35 _dafny.Set
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf31 _dafny.Array, Cf32 _dafny.Int, Cf33 _dafny.Int, Cf34 _dafny.Map, Cf35 _dafny.Set) D5 {
	return D5{D5_DC14{Cf31, Cf32, Cf33, Cf34, Cf35}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC12 struct {
	Cf26 _dafny.Set
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf26 _dafny.Set) D5 {
	return D5{D5_DC12{Cf26}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC15 struct {
	Cf36 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf36 D5) D5 {
	return D5{D5_DC15{Cf36}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D5) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf28
}

func (_this D5) Dtor_cf29() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf29
}

func (_this D5) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf30
}

func (_this D5) Dtor_cf31() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf31
}

func (_this D5) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf32
}

func (_this D5) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf33
}

func (_this D5) Dtor_cf34() _dafny.Map {
	return _this.Get_().(D5_DC14).Cf34
}

func (_this D5) Dtor_cf35() _dafny.Set {
	return _this.Get_().(D5_DC14).Cf35
}

func (_this D5) Dtor_cf26() _dafny.Set {
	return _this.Get_().(D5_DC12).Cf26
}

func (_this D5) Dtor_cf36() D5 {
	return _this.Get_().(D5_DC15).Cf36
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29 && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34.Equals(data2.Cf34) && data1.Cf35.Equals(data2.Cf35)
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf26.Equals(data2.Cf26)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf36.Equals(data2.Cf36)
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
	Cf38 _dafny.Int
	Cf39 _dafny.Sequence
	Cf40 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf38 _dafny.Int, Cf39 _dafny.Sequence, Cf40 bool) D6 {
	return D6{D6_DC17{Cf38, Cf39, Cf40}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf37 _dafny.Set
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf37 _dafny.Set) D6 {
	return D6{D6_DC16{Cf37}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(_dafny.Zero, _dafny.EmptySeq, false)
}

func (_this D6) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf38
}

func (_this D6) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf39
}

func (_this D6) Dtor_cf40() bool {
	return _this.Get_().(D6_DC17).Cf40
}

func (_this D6) Dtor_cf37() _dafny.Set {
	return _this.Get_().(D6_DC16).Cf37
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf38) + ", " + data.Cf39.VerbatimString(true) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf37) + ")"
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
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39.Equals(data2.Cf39) && data1.Cf40 == data2.Cf40
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf37.Equals(data2.Cf37)
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
	Cf41 _dafny.MultiSet
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf41 _dafny.MultiSet) D7 {
	return D7{D7_DC18{Cf41}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D7) Dtor_cf41() _dafny.MultiSet {
	return _this.Get_().(D7_DC18).Cf41
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf41) + ")"
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
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf41.Equals(data2.Cf41)
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
	Cf43 _dafny.CodePoint
	Cf44 bool
	Cf45 bool
	Cf46 _dafny.Map
	Cf47 bool
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf43 _dafny.CodePoint, Cf44 bool, Cf45 bool, Cf46 _dafny.Map, Cf47 bool) D8 {
	return D8{D8_DC20{Cf43, Cf44, Cf45, Cf46, Cf47}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC19 struct {
	Cf42 _dafny.Map
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf42 _dafny.Map) D8 {
	return D8{D8_DC19{Cf42}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC20_(_dafny.CodePoint('D'), false, false, _dafny.EmptyMap, false)
}

func (_this D8) Dtor_cf43() _dafny.CodePoint {
	return _this.Get_().(D8_DC20).Cf43
}

func (_this D8) Dtor_cf44() bool {
	return _this.Get_().(D8_DC20).Cf44
}

func (_this D8) Dtor_cf45() bool {
	return _this.Get_().(D8_DC20).Cf45
}

func (_this D8) Dtor_cf46() _dafny.Map {
	return _this.Get_().(D8_DC20).Cf46
}

func (_this D8) Dtor_cf47() bool {
	return _this.Get_().(D8_DC20).Cf47
}

func (_this D8) Dtor_cf42() _dafny.Map {
	return _this.Get_().(D8_DC19).Cf42
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf43 == data2.Cf43 && data1.Cf44 == data2.Cf44 && data1.Cf45 == data2.Cf45 && data1.Cf46.Equals(data2.Cf46) && data1.Cf47 == data2.Cf47
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
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

type D9_DC22 struct {
	Cf49 _dafny.Int
	Cf50 bool
	Cf51 _dafny.Map
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf49 _dafny.Int, Cf50 bool, Cf51 _dafny.Map) D9 {
	return D9{D9_DC22{Cf49, Cf50, Cf51}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC23 struct {
	Cf52 _dafny.Int
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf52 _dafny.Int) D9 {
	return D9{D9_DC23{Cf52}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC24 struct {
	Cf53 _dafny.Int
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf53 _dafny.Int) D9 {
	return D9{D9_DC24{Cf53}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC21 struct {
	Cf48 _dafny.Sequence
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf48 _dafny.Sequence) D9 {
	return D9{D9_DC21{Cf48}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC22_(_dafny.Zero, false, _dafny.EmptyMap)
}

func (_this D9) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D9_DC22).Cf49
}

func (_this D9) Dtor_cf50() bool {
	return _this.Get_().(D9_DC22).Cf50
}

func (_this D9) Dtor_cf51() _dafny.Map {
	return _this.Get_().(D9_DC22).Cf51
}

func (_this D9) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf52
}

func (_this D9) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf53
}

func (_this D9) Dtor_cf48() _dafny.Sequence {
	return _this.Get_().(D9_DC21).Cf48
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf52) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf53) + ")"
		}
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf49.Cmp(data2.Cf49) == 0 && data1.Cf50 == data2.Cf50 && data1.Cf51.Equals(data2.Cf51)
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf52.Cmp(data2.Cf52) == 0
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf53.Cmp(data2.Cf53) == 0
		}
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf48.Equals(data2.Cf48)
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

type D10_DC25 struct {
	Cf54 *C5
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf54 *C5) D10 {
	return D10{D10_DC25{Cf54}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

func (CompanionStruct_D10_) Default() *C5 {
	return (*C5)(nil)
}

func (_this D10) Dtor_cf54() *C5 {
	return _this.Get_().(D10_DC25).Cf54
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC25:
		{
			return "D10.DC25" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
			return ok && data1.Cf54 == data2.Cf54
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

// Definition of trait T0
type T0 interface {
	String() string
	F10() _dafny.MultiSet
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
	M1(p0 _dafny.CodePoint, p1 _dafny.Int, p2 D0, globalState *GlobalState)
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

// Definition of trait T2
type T2 interface {
	String() string
	Fm5(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool
	Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int
}
type CompanionStruct_T2_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T2_ = CompanionStruct_T2_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T2_) CastTo_(x interface{}) T2 {
	var t T2
	t, _ = x.(T2)
	return t
}

// End of trait T2

// Definition of class GlobalState
type GlobalState struct {
	F1  _dafny.Int
	F5  _dafny.Array
	_f0 bool
	_f2 _dafny.Int
	_f3 bool
	_f4 _dafny.Array
	_f6 _dafny.MultiSet
	_f7 _dafny.Sequence
	_f8 bool
	_f9 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this.F5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = false
	_this._f2 = _dafny.Zero
	_this._f3 = false
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.EmptyMultiSet
	_this._f7 = _dafny.EmptySeq
	_this._f8 = false
	_this._f9 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.Int, f3 bool, f4 _dafny.Array, f5 _dafny.Array, f6 _dafny.MultiSet, f7 _dafny.Sequence, f8 bool, f9 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Array {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.MultiSet {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Sequence {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_}
}

var _ T2 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__() {
	{
	}
}
func (_this *C0) Fm5(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C0) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.SetOf(false, true)).Cardinality(), _dafny.IntOfInt64(-766)), _dafny.SeqOf(_dafny.IntOfInt64(131)))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("igq")).Cardinality()))
	}
}
func (_this *C0) Fm13(p0 bool, globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate((Companion_D3_.Create_DC8_(_dafny.CodePoint('h'), true, false, _dafny.UnicodeSeqOfUtf8Bytes("xnohburd"))).Dtor_cf20(), _dafny.UnicodeSeqOfUtf8Bytes("wami")), _dafny.CodePoint('c'))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f10 _dafny.MultiSet
	_f14 bool
	_f15 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f10 = _dafny.EmptyMultiSet
	_this._f14 = false
	_this._f15 = false
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

func (_this *C1) F10() _dafny.MultiSet {
	return _this._f10
}
func (_this *C1) Ctor__(f14 bool, f15 bool, f10 _dafny.MultiSet) {
	{
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f10 = f10
	}
}
func (_this *C1) M5(globalState *GlobalState) (_dafny.Array, _dafny.Int, _dafny.Map) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var _179_v0 bool
		_ = _179_v0
		_179_v0 = false
		_179_v0 = _179_v0
		var _180_v1 *C0
		_ = _180_v1
		var _nw17 *C0 = New_C0_()
		_ = _nw17
		_nw17.Ctor__()
		_180_v1 = _nw17
		var _181_v2 D4
		_ = _181_v2
		_181_v2 = Companion_D4_.Create_DC10_(_180_v1)
		var _182_v3 _dafny.Int
		_ = _182_v3
		_182_v3 = _dafny.IntOfInt64(979)
		var _183_v4 _dafny.CodePoint
		_ = _183_v4
		_183_v4 = _dafny.CodePoint('a')
		var _rhs14 *C0 = (_181_v2).Dtor_cf22()
		_ = _rhs14
		var _rhs15 bool = _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm2(_179_v0, _182_v3, _183_v4, _182_v3, globalState), _dafny.UnicodeSeqOfUtf8Bytes("w"))
		_ = _rhs15
		_180_v1 = _rhs14
		_179_v0 = _rhs15
		var _184_v5 _dafny.Sequence
		_ = _184_v5
		_184_v5 = _dafny.SeqOf(_182_v3)
		var _hi2 _dafny.Int = (_184_v5).Select((Companion_Default___.SafeIndex(_182_v3, _dafny.IntOfUint32((_184_v5).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi2
		for _185_i0 := _182_v3; _185_i0.Cmp(_hi2) < 0; _185_i0 = _185_i0.Plus(_dafny.One) {
			var _186_v6 _dafny.Array
			_ = _186_v6
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw18
			_186_v6 = _nw18
			var _187_v7 _dafny.Map
			_ = _187_v7
			_187_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _186_v6)
			var _188_v8 _dafny.Array
			_ = _188_v8
			var _nwElement0_4 _dafny.Array = _186_v6
			_ = _nwElement0_4
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(4))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_4, 0)
			(_nw19).ArraySet1(_186_v6, 1)
			(_nw19).ArraySet1((func() _dafny.Array {
				if (_187_v7).Contains(_179_v0) {
					return (_187_v7).Get(_179_v0).(_dafny.Array)
				}
				return _186_v6
			})(), 2)
			(_nw19).ArraySet1(_186_v6, 3)
			_188_v8 = _nw19
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_188_v8), 0))
			_ = _index21
			(_188_v8).ArraySet1(_186_v6, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_188_v8), 0))
			_ = _index22
			(_188_v8).ArraySet1((func() _dafny.Array {
				if false {
					return _186_v6
				}
				return _186_v6
			})(), (_index22).Int())
			var _189_v9 _dafny.Array
			_ = _189_v9
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_0
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_190_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_191_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_190_v4), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg20 _dafny.Int) interface{} {
								return coer20(arg20)
							}
						}((func(_192_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_193_i2 _dafny.Int) _dafny.CodePoint {
								return _192_v4
							}
						})(_190_v4))))
					}
				})(_183_v4)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw20 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw20).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw20).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_189_v9 = _nw20
			var _194_v10 _dafny.Sequence
			_ = _194_v10
			_194_v10 = _dafny.SeqOf(_183_v4)
			var _195_v11 _dafny.MultiSet
			_ = _195_v11
			_195_v11 = _dafny.MultiSetOf(_185_i0, _182_v3, _dafny.IntOfUint32((_194_v10).Cardinality()), _185_i0, _dafny.IntOfUint32((_184_v5).Cardinality()))
			var _196_v12 _dafny.Sequence
			_ = _196_v12
			_196_v12 = _dafny.SeqOf((_194_v10).Select((Companion_Default___.SafeIndex((_195_v11).Cardinality(), _dafny.IntOfUint32((_194_v10).Cardinality()))).Uint32()).(_dafny.CodePoint), _183_v4, _183_v4, _183_v4, _dafny.CodePoint('s'))
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_189_v9), 0))
			_ = _index23
			(_189_v9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_196_v12, _196_v12), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_189_v9), 0))
			_ = _index24
			(_189_v9).ArraySet1(_196_v12, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_189_v9), 0))
			_ = _index25
			(_189_v9).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-892))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_197_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_198_i3 _dafny.Int) _dafny.CodePoint {
					return _197_v4
				}
			})(_183_v4))), (_index25).Int())
			var _199_v13 *C0
			_ = _199_v13
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__()
			_199_v13 = _nw21
		}
		var _200_v14 _dafny.Set
		_ = _200_v14
		_200_v14 = _dafny.SetOf(_182_v3)
		var _201_v16 _dafny.Sequence
		_ = _201_v16
		_201_v16 = _dafny.SeqOf(_200_v14, _dafny.SetOf(_dafny.IntOfUint32((_184_v5).Cardinality()), _182_v3))
		var _202_v17 _dafny.Array
		_ = _202_v17
		var _nwElement0_5 _dafny.Set = _200_v14
		_ = _nwElement0_5
		var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(17))
		_ = _nw22
		(_nw22).ArraySet1(_nwElement0_5, 0)
		(_nw22).ArraySet1((_dafny.SetOf(_182_v3, (_180_v1).Fm6(_179_v0, _182_v3, globalState))).Union(_200_v14), 1)
		(_nw22).ArraySet1(_200_v14, 2)
		(_nw22).ArraySet1(_200_v14, 3)
		(_nw22).ArraySet1(_200_v14, 4)
		(_nw22).ArraySet1(_200_v14, 5)
		(_nw22).ArraySet1((_200_v14).Union(_200_v14), 6)
		(_nw22).ArraySet1(func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(35), _dafny.IntOfInt64(104))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _203_v15 _dafny.Int
				_203_v15 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(35)).Cmp(_203_v15) <= 0) && ((_203_v15).Cmp(_dafny.IntOfInt64(104)) < 0) {
					_coll10.Add(Companion_Default___.SafeModuloInt(_203_v15, (func() _dafny.Int {
						if ((_this).F10()).Contains((_this).F14()) {
							return ((_this).F10()).Multiplicity((_this).F14())
						}
						return (_dafny.SetOf((_this).F14())).Cardinality()
					})()))
				}
			}
			return _coll10.ToSet()
		}(), 7)
		(_nw22).ArraySet1(_200_v14, 8)
		(_nw22).ArraySet1((_200_v14).Union(_200_v14), 9)
		(_nw22).ArraySet1(_200_v14, 10)
		(_nw22).ArraySet1(_200_v14, 11)
		(_nw22).ArraySet1((_201_v16).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_182_v3), _dafny.IntOfUint32((_201_v16).Cardinality()))).Uint32()).(_dafny.Set), 12)
		(_nw22).ArraySet1(_200_v14, 13)
		(_nw22).ArraySet1(_200_v14, 14)
		(_nw22).ArraySet1(_200_v14, 15)
		(_nw22).ArraySet1(_200_v14, 16)
		_202_v17 = _nw22
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_202_v17), 0))); ; {
			_guard_loop_0, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _204_i4 _dafny.Int
			_204_i4 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_204_i4).Sign() != -1) && ((_204_i4).Cmp(_dafny.ArrayLen((_202_v17), 0)) < 0)) {
				(_202_v17).ArraySet1(_200_v14, (_204_i4).Int())
			}
		}
		var _205_i5 _dafny.Int
		_ = _205_i5
		_205_i5 = _dafny.Zero
		{
			for Companion_Default___.Fm4((func() _dafny.Set {
				if (_this).F14() {
					return _200_v14
				}
				return _200_v14
			})(), globalState) {
				{
					if (_205_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_205_i5 = (_205_i5).Plus(_dafny.One)
					r1 = _dafny.IntOfInt64(618)
					if (_this).F15() {
						var _206_v18 T2
						_ = _206_v18
						var _nw23 *C0 = New_C0_()
						_ = _nw23
						_nw23.Ctor__()
						_206_v18 = _nw23
						_206_v18 = _206_v18
						var _207_v19 _dafny.Sequence
						_ = _207_v19
						_207_v19 = _dafny.UnicodeSeqOfUtf8Bytes("efljrb")
						var _208_v20 _dafny.Map
						_ = _208_v20
						_208_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v19, true)
						var _209_v22 _dafny.Sequence
						_ = _209_v22
						_209_v22 = _dafny.SeqOf(_207_v19)
						var _210_v23 _dafny.Map
						_ = _210_v23
						_210_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F15())
						var _211_v24 _dafny.Sequence
						_ = _211_v24
						_211_v24 = _dafny.SeqOf(_210_v23, _210_v23, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v0, (_this).F15()), _210_v23, _210_v23)
						var _212_v25 _dafny.Sequence
						_ = _212_v25
						_212_v25 = _dafny.SeqOf((_211_v24).Select((Companion_Default___.SafeIndex(_182_v3, _dafny.IntOfUint32((_211_v24).Cardinality()))).Uint32()).(_dafny.Map))
						var _213_v26 _dafny.Map
						_ = _213_v26
						_213_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_211_v24).Select((Companion_Default___.SafeIndex(_182_v3, _dafny.IntOfUint32((_211_v24).Cardinality()))).Uint32()).(_dafny.Map)).Update(_179_v0, (_this).F14()), _dafny.IntOfInt64(433))
						var _214_v28 _dafny.Array
						_ = _214_v28
						var _nwElement0_6 _dafny.Int = ((_208_v20).Merge(func() _dafny.Map {
							var _coll11 = _dafny.NewMapBuilder()
							_ = _coll11
							for _iter12 := _dafny.Iterate((_209_v22).Elements()); ; {
								_compr_11, _ok12 := _iter12()
								if !_ok12 {
									break
								}
								var _215_v21 _dafny.Sequence
								_215_v21 = interface{}(_compr_11).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_209_v22, _215_v21) {
									_coll11.Add(_215_v21, _179_v0)
								}
							}
							return _coll11.ToMap()
						}())).Cardinality()
						_ = _nwElement0_6
						var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(27))
						_ = _nw24
						(_nw24).ArraySet1(_nwElement0_6, 0)
						(_nw24).ArraySet1(_182_v3, 1)
						(_nw24).ArraySet1(_182_v3, 2)
						(_nw24).ArraySet1((_182_v3).Times(_182_v3), 3)
						(_nw24).ArraySet1(((_this).F10()).Cardinality(), 4)
						(_nw24).ArraySet1((_206_v18).Fm6((_this).F14(), _dafny.IntOfUint32((_207_v19).Cardinality()), globalState), 5)
						(_nw24).ArraySet1(_182_v3, 6)
						(_nw24).ArraySet1(_182_v3, 7)
						(_nw24).ArraySet1(_182_v3, 8)
						(_nw24).ArraySet1(_182_v3, 9)
						(_nw24).ArraySet1(_182_v3, 10)
						(_nw24).ArraySet1(_dafny.IntOfUint32((_207_v19).Cardinality()), 11)
						(_nw24).ArraySet1(_dafny.IntOfInt64(350), 12)
						(_nw24).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(439))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg22 _dafny.Int) interface{} {
								return coer22(arg22)
							}
						}((func(_216_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_217_i6 _dafny.Int) _dafny.CodePoint {
								return _216_v4
							}
						})(_183_v4)))).Cardinality()), 13)
						(_nw24).ArraySet1(_182_v3, 14)
						(_nw24).ArraySet1(_182_v3, 15)
						(_nw24).ArraySet1((_180_v1).Fm6(!((_this).F15()), _182_v3, globalState), 16)
						(_nw24).ArraySet1((_200_v14).Cardinality(), 17)
						(_nw24).ArraySet1(_182_v3, 18)
						(_nw24).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_184_v5, (Companion_Default___.SafeIndex(_182_v3, _dafny.IntOfUint32((_184_v5).Cardinality()))).Uint32(), _182_v3)).Cardinality()), _182_v3), 19)
						(_nw24).ArraySet1((_180_v1).Fm6((_this).F15(), _182_v3, globalState), 20)
						(_nw24).ArraySet1(_182_v3, 21)
						(_nw24).ArraySet1(_182_v3, 22)
						(_nw24).ArraySet1(Companion_Default___.SafeModuloInt(_182_v3, _182_v3), 23)
						(_nw24).ArraySet1((_182_v3).Plus(_182_v3), 24)
						(_nw24).ArraySet1((func() _dafny.Int {
							if (_213_v26).Contains(_210_v23) {
								return (_213_v26).Get(_210_v23).(_dafny.Int)
							}
							return (_dafny.Zero).Minus((func() _dafny.Set {
								var _coll12 = _dafny.NewBuilder()
								_ = _coll12
								for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(218), _dafny.IntOfInt64(629))); ; {
									_compr_12, _ok13 := _iter13()
									if !_ok13 {
										break
									}
									var _218_v27 _dafny.Int
									_218_v27 = interface{}(_compr_12).(_dafny.Int)
									if ((_dafny.IntOfInt64(218)).Cmp(_218_v27) <= 0) && ((_218_v27).Cmp(_dafny.IntOfInt64(629)) < 0) {
										_coll12.Add((_218_v27).Times(_182_v3))
									}
								}
								return _coll12.ToSet()
							}()).Cardinality())
						})(), 25)
						(_nw24).ArraySet1(_182_v3, 26)
						_214_v28 = _nw24
						var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_214_v28), 0))
						_ = _index26
						(_214_v28).ArraySet1(_dafny.IntOfInt64(883), (_index26).Int())
						var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_214_v28), 0))
						_ = _index27
						(_214_v28).ArraySet1(_182_v3, (_index27).Int())
						_179_v0 = false
						var _219_v29 _dafny.Array
						_ = _219_v29
						var _len0_1 _dafny.Int = _dafny.IntOfInt64(9)
						_ = _len0_1
						var _nw25 _dafny.Array
						_ = _nw25
						if _len0_1.Cmp(_dafny.Zero) == 0 {
							_nw25 = _dafny.NewArray(_len0_1)
						} else {
							var _init1 func(_dafny.Int) bool = func(_220_i7 _dafny.Int) bool {
								return (_this).F15()
							}
							_ = _init1
							var _element0_1 = _init1(_dafny.Zero)
							_ = _element0_1
							_nw25 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
							(_nw25).ArraySet1(_element0_1, 0)
							var _nativeLen0_1 = (_len0_1).Int()
							_ = _nativeLen0_1
							for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
								(_nw25).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
							}
						}
						_219_v29 = _nw25
						r0 = _219_v29
						var _rhs16 _dafny.Int = ((func() _dafny.Int {
							if (_this).F15() {
								return (_214_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_214_v28), 0))).Int()).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_207_v19).Cardinality())
						})()).Plus(_182_v3)
						_ = _rhs16
						var _rhs17 bool = (_this).F14()
						_ = _rhs17
						var _lhs14 *GlobalState = globalState
						_ = _lhs14
						_lhs14.F1 = _rhs16
						_179_v0 = _rhs17
					} else {
						var _221_v30 _dafny.Array
						_ = _221_v30
						var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
						_ = _nw26
						_221_v30 = _nw26
						var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))
						_ = _index28
						(_221_v30).ArraySet1(Companion_Default___.SafeDivisionInt(_182_v3, ((_this).F10()).Cardinality()), (_index28).Int())
						var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))
						_ = _index29
						(_221_v30).ArraySet1((func() _dafny.Int {
							if ((_this).F10()).Contains((_182_v3).Cmp(_182_v3) == 0) {
								return ((_this).F10()).Multiplicity((_182_v3).Cmp(_182_v3) == 0)
							}
							return _182_v3
						})(), (_index29).Int())
						var _222_v31 _dafny.Sequence
						_ = _222_v31
						_222_v31 = _dafny.SeqOf((_this).F14(), (_this).F14())
						var _rhs18 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_222_v31, _222_v31)
						_ = _rhs18
						var _rhs19 bool = _179_v0
						_ = _rhs19
						var _rhs20 _dafny.Int = _dafny.IntOfInt64(490)
						_ = _rhs20
						var _lhs15 *GlobalState = globalState
						_ = _lhs15
						_222_v31 = _rhs18
						_179_v0 = _rhs19
						_lhs15.F1 = _rhs20
						var _223_v32 D1
						_ = _223_v32
						_223_v32 = Companion_D1_.Create_DC5_(_182_v3, (_this).F14(), _dafny.IntOfInt64(166), !(_179_v0) || (_179_v0), _182_v3)
						var _224_v33 _dafny.Sequence
						_ = _224_v33
						_224_v33 = _dafny.UnicodeSeqOfUtf8Bytes("ko")
						_223_v32 = Companion_D1_.Create_DC5_((_221_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))).Int()).(_dafny.Int), (_this).F14(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(791), (_221_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("scqxnyhq"), _224_v33), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_222_v31, _222_v31)).Cardinality()))
						var _225_v34 _dafny.MultiSet
						_ = _225_v34
						_225_v34 = _dafny.MultiSetOf(_dafny.IntOfInt64(330))
						_182_v3 = (func() _dafny.Int {
							if (_225_v34).Contains((_221_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))).Int()).(_dafny.Int)) {
								return (_225_v34).Multiplicity((_221_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))).Int()).(_dafny.Int))
							}
							return (_221_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))).Int()).(_dafny.Int)
						})()
						var _226_v35 _dafny.Array
						_ = _226_v35
						var _len0_2 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_2
						var _nw27 _dafny.Array
						_ = _nw27
						if _len0_2.Cmp(_dafny.Zero) == 0 {
							_nw27 = _dafny.NewArray(_len0_2)
						} else {
							var _init2 func(_dafny.Int) bool = func(_227_i8 _dafny.Int) bool {
								return (_this).F14()
							}
							_ = _init2
							var _element0_2 = _init2(_dafny.Zero)
							_ = _element0_2
							_nw27 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
							(_nw27).ArraySet1(_element0_2, 0)
							var _nativeLen0_2 = (_len0_2).Int()
							_ = _nativeLen0_2
							for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
								(_nw27).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
							}
						}
						_226_v35 = _nw27
						var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_226_v35), 0))
						_ = _index30
						(_226_v35).ArraySet1(_179_v0, (_index30).Int())
						var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_226_v35), 0))
						_ = _index31
						var _rhs21 bool = _179_v0
						_ = _rhs21
						var _rhs22 bool = !((_this).F14())
						_ = _rhs22
						var _rhs23 _dafny.Sequence = _224_v33
						_ = _rhs23
						var _rhs24 _dafny.Int = (_221_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))).Int()).(_dafny.Int)
						_ = _rhs24
						var _rhs25 _dafny.Int = (_221_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_221_v30), 0))).Int()).(_dafny.Int)
						_ = _rhs25
						var _lhs16 _dafny.Array = _226_v35
						_ = _lhs16
						var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_226_v35), 0))
						_ = _lhs17
						var _lhs18 *GlobalState = globalState
						_ = _lhs18
						var _lhs19 *GlobalState = globalState
						_ = _lhs19
						(_lhs16).ArraySet1(_rhs21, (_lhs17).Int())
						_179_v0 = _rhs22
						_224_v33 = _rhs23
						_lhs18.F1 = _rhs24
						_lhs19.F1 = _rhs25
					}
					var _228_v36 _dafny.MultiSet
					_ = _228_v36
					_228_v36 = _dafny.MultiSetFromSeq(_184_v5)
					var _source6 _dafny.MultiSet = _228_v36
					_ = _source6
					var _229___mcc_h0 _dafny.MultiSet = _source6
					_ = _229___mcc_h0
					var _230_cf15 _dafny.MultiSet = _229___mcc_h0
					_ = _230_cf15
					var _231_v37 *C0
					_ = _231_v37
					var _nw28 *C0 = New_C0_()
					_ = _nw28
					_nw28.Ctor__()
					_231_v37 = _nw28
					var _232_v38 *C0
					_ = _232_v38
					var _nw29 *C0 = New_C0_()
					_ = _nw29
					_nw29.Ctor__()
					_232_v38 = _nw29
					var _233_v39 _dafny.Array
					_ = _233_v39
					var _nwElement0_7 bool = (_this).F15()
					_ = _nwElement0_7
					var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(20))
					_ = _nw30
					(_nw30).ArraySet1(_nwElement0_7, 0)
					(_nw30).ArraySet1((_this).F15(), 1)
					(_nw30).ArraySet1(_179_v0, 2)
					(_nw30).ArraySet1(_179_v0, 3)
					(_nw30).ArraySet1((_this).F15(), 4)
					(_nw30).ArraySet1((_this).F14(), 5)
					(_nw30).ArraySet1(Companion_Default___.Fm4(_200_v14, globalState), 6)
					(_nw30).ArraySet1(false, 7)
					(_nw30).ArraySet1(_179_v0, 8)
					(_nw30).ArraySet1(_179_v0, 9)
					(_nw30).ArraySet1((_this).F15(), 10)
					(_nw30).ArraySet1((_this).F15(), 11)
					(_nw30).ArraySet1(true, 12)
					(_nw30).ArraySet1(!(true), 13)
					(_nw30).ArraySet1((_this).F14(), 14)
					(_nw30).ArraySet1((_this).F15(), 15)
					(_nw30).ArraySet1((_this).F15(), 16)
					(_nw30).ArraySet1((_this).F15(), 17)
					(_nw30).ArraySet1(_179_v0, 18)
					(_nw30).ArraySet1((_this).F15(), 19)
					_233_v39 = _nw30
					var _234_v40 _dafny.MultiSet
					_ = _234_v40
					_234_v40 = _dafny.MultiSetOf(_233_v39)
					_179_v0 = ((_234_v40).Intersection(_234_v40)).IsDisjointFrom((_234_v40).Difference(_234_v40))
					(globalState).F1 = (_182_v3).Plus(_182_v3)
					var _rhs26 bool = !(_179_v0)
					_ = _rhs26
					var _rhs27 bool = (_this).F14()
					_ = _rhs27
					var _rhs28 bool = (_dafny.IntOfInt64(-259)).Cmp((_180_v1).Fm6((_this).F14(), _182_v3, globalState)) != 0
					_ = _rhs28
					var _rhs29 _dafny.Int = _dafny.IntOfInt64(2)
					_ = _rhs29
					_179_v0 = _rhs26
					_179_v0 = _rhs27
					_179_v0 = _rhs28
					r1 = _rhs29
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _hi3 _dafny.Int = _182_v3
		_ = _hi3
		for _235_i9 := (_182_v3).Times(_182_v3); _235_i9.Cmp(_hi3) < 0; _235_i9 = _235_i9.Plus(_dafny.One) {
			var _236_v41 _dafny.Map
			_ = _236_v41
			_236_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v3, _235_i9)
			_236_v41 = (_236_v41).Update((_dafny.Zero).Minus((_235_i9).Minus(_dafny.IntOfInt64(341))), _235_i9)
			var _237_v42 _dafny.Sequence
			_ = _237_v42
			_237_v42 = _dafny.UnicodeSeqOfUtf8Bytes("vkio")
			var _238_v43 _dafny.Sequence
			_ = _238_v43
			_238_v43 = _dafny.SeqOf((_this).F14())
			var _239_v44 _dafny.Array
			_ = _239_v44
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_3
			var _nw31 _dafny.Array
			_ = _nw31
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw31 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = func(_240_i10 _dafny.Int) bool {
					return (_this).F15()
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw31 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw31).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw31).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_239_v44 = _nw31
			var _241_v45 _dafny.Map
			_ = _241_v45
			_241_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm1(globalState), _183_v4, _183_v4)).Cardinality()), _239_v44)
			var _242_v48 _dafny.Array
			_ = _242_v48
			var _nwElement0_8 bool = (_this).F14()
			_ = _nwElement0_8
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_8, 0)
			(_nw32).ArraySet1((_this).F14(), 1)
			(_nw32).ArraySet1((_this).F14(), 2)
			(_nw32).ArraySet1((_this).F15(), 3)
			(_nw32).ArraySet1(((_this).F10()).IsSubsetOf((_this).F10()), 4)
			(_nw32).ArraySet1(_179_v0, 5)
			(_nw32).ArraySet1(((_this).F14()) || ((_this).F14()), 6)
			(_nw32).ArraySet1((_this).F15(), 7)
			(_nw32).ArraySet1(_179_v0, 8)
			(_nw32).ArraySet1((_182_v3).Cmp(_235_i9) >= 0, 9)
			(_nw32).ArraySet1(_179_v0, 10)
			(_nw32).ArraySet1((_this).F15(), 11)
			(_nw32).ArraySet1(!_dafny.Companion_Sequence_.Equal(_237_v42, _237_v42), 12)
			(_nw32).ArraySet1(!((_this).F15()), 13)
			(_nw32).ArraySet1((_180_v1).Fm5(Companion_Default___.Fm14((_this).F14(), globalState), _235_i9, _235_i9, globalState), 14)
			(_nw32).ArraySet1(!((_this).F14()), 15)
			(_nw32).ArraySet1(_dafny.Companion_Sequence_.Equal(_238_v43, _dafny.Companion_Sequence_.Update(_238_v43, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_238_v43).Cardinality()), _dafny.IntOfUint32((_238_v43).Cardinality()))).Uint32(), (_this).F15())), 16)
			(_nw32).ArraySet1((_this).F15(), 17)
			(_nw32).ArraySet1(true, 18)
			(_nw32).ArraySet1(((_180_v1).Fm13((_this).F15(), globalState)) == (!((_this).F15())), 19)
			(_nw32).ArraySet1((_this).F14(), 20)
			(_nw32).ArraySet1((_this).F15(), 21)
			(_nw32).ArraySet1((_235_i9).Cmp(_dafny.IntOfInt64(-262)) > 0, 22)
			(_nw32).ArraySet1((_235_i9).Cmp((_241_v45).Cardinality()) <= 0, 23)
			(_nw32).ArraySet1((func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(708), _dafny.IntOfInt64(526))); ; {
					_compr_13, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _243_v46 _dafny.Int
					_243_v46 = interface{}(_compr_13).(_dafny.Int)
					if ((_dafny.IntOfInt64(708)).Cmp(_243_v46) <= 0) && ((_243_v46).Cmp(_dafny.IntOfInt64(526)) < 0) {
						_coll13.Add(Companion_Default___.SafeDivisionInt(_243_v46, _182_v3))
					}
				}
				return _coll13.ToSet()
			}()).IsDisjointFrom(func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(184), _dafny.IntOfInt64(772))); ; {
					_compr_14, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _244_v47 _dafny.Int
					_244_v47 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(184)).Cmp(_244_v47) <= 0) && ((_244_v47).Cmp(_dafny.IntOfInt64(772)) < 0) {
						_coll14.Add((_244_v47).Times(_182_v3))
					}
				}
				return _coll14.ToSet()
			}()), 24)
			_242_v48 = _nw32
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_242_v48), 0))
			_ = _index32
			(_242_v48).ArraySet1(_179_v0, (_index32).Int())
			var _245_v49 _dafny.Map
			_ = _245_v49
			_245_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v0, _237_v42)
			var _246_v50 _dafny.MultiSet
			_ = _246_v50
			_246_v50 = _dafny.MultiSetOf((_245_v49).Cardinality())
			var _247_v51 _dafny.Map
			_ = _247_v51
			_247_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_246_v50).Cardinality())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_242_v48), 0))
			_ = _index33
			(_242_v48).ArraySet1(!((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(842))).Merge(_247_v51)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _235_i9)), (_index33).Int())
			if (_this).F14() {
				_237_v42 = _237_v42
				_237_v42 = (func(_pat_let2_0 D3) D3 {
					return func(_248_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let3_0 _dafny.Sequence) D3 {
							return func(_250_dt__update_hcf16_h0 _dafny.Sequence) D3 {
								return Companion_D3_.Create_DC7_(_250_dt__update_hcf16_h0)
							}(_pat_let3_0)
						}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(228))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg23 _dafny.Int) interface{} {
								return coer23(arg23)
							}
						}(func(_249_i11 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('s')
						})))
					}(_pat_let2_0)
				}(Companion_D3_.Create_DC7_(_237_v42))).Dtor_cf16()
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_242_v48), 0))
				_ = _index34
				(_242_v48).ArraySet1((_this).F15(), (_index34).Int())
				var _251_v52 _dafny.Array
				_ = _251_v52
				var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
				_ = _nw33
				_251_v52 = _nw33
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_251_v52), 0))
				_ = _index35
				(_251_v52).ArraySet1(_235_i9, (_index35).Int())
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_251_v52), 0))
				_ = _index36
				var _rhs30 _dafny.Int = _182_v3
				_ = _rhs30
				var _rhs31 _dafny.Int = _182_v3
				_ = _rhs31
				var _lhs20 _dafny.Array = _251_v52
				_ = _lhs20
				var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_251_v52), 0))
				_ = _lhs21
				r1 = _rhs30
				(_lhs20).ArraySet1(_rhs31, (_lhs21).Int())
				var _252_v53 _dafny.Set
				_ = _252_v53
				_252_v53 = _dafny.SetOf(_179_v0)
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_251_v52), 0))
				_ = _index37
				(_251_v52).ArraySet1((_dafny.Zero).Minus((_184_v5).Select((Companion_Default___.SafeIndex((_252_v53).Cardinality(), _dafny.IntOfUint32((_184_v5).Cardinality()))).Uint32()).(_dafny.Int)), (_index37).Int())
			} else {
				var _253_v54 _dafny.MultiSet
				_ = _253_v54
				_253_v54 = _dafny.MultiSetOf((func() _dafny.Sequence {
					if (_this).F14() {
						return _237_v42
					}
					return _237_v42
				})())
				_253_v54 = (_253_v54).Intersection(((_253_v54).Update(_237_v42, Companion_Default___.Abs(_dafny.IntOfUint32((_238_v43).Cardinality())))).Union(_253_v54))
				var _254_v55 _dafny.Array
				_ = _254_v55
				var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw34
				_254_v55 = _nw34
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_254_v55), 0))
				_ = _index38
				(_254_v55).ArraySet1(_235_i9, (_index38).Int())
				var _255_v56 _dafny.Map
				_ = _255_v56
				_255_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v3, _179_v0)
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_254_v55), 0))
				_ = _index39
				var _rhs32 _dafny.Int = _182_v3
				_ = _rhs32
				var _rhs33 _dafny.Int = (_180_v1).Fm6((func() bool {
					if (_255_v56).Contains(_182_v3) {
						return (_255_v56).Get(_182_v3).(bool)
					}
					return (_this).F15()
				})(), _182_v3, globalState)
				_ = _rhs33
				var _lhs22 *GlobalState = globalState
				_ = _lhs22
				var _lhs23 _dafny.Array = _254_v55
				_ = _lhs23
				var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_254_v55), 0))
				_ = _lhs24
				_lhs22.F1 = _rhs32
				(_lhs23).ArraySet1(_rhs33, (_lhs24).Int())
				_182_v3 = _dafny.IntOfUint32((_237_v42).Cardinality())
				var _256_v57 _dafny.MultiSet
				_ = _256_v57
				_256_v57 = _dafny.MultiSetOf(_183_v4, _183_v4, _183_v4)
				var _257_v58 D0
				_ = _257_v58
				_257_v58 = Companion_D0_.Create_DC1_((_this).F15(), (_256_v57).Cardinality())
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_242_v48), 0))
				_ = _index40
				(_242_v48).ArraySet1((_257_v58).Dtor_cf1(), (_index40).Int())
				var _258_v59 *C0
				_ = _258_v59
				var _nw35 *C0 = New_C0_()
				_ = _nw35
				_nw35.Ctor__()
				_258_v59 = _nw35
			}
			_238_v43 = _dafny.Companion_Sequence_.Concatenate(_238_v43, _238_v43)
		}
		var _259_v60 _dafny.Array
		_ = _259_v60
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_4
		var _nw36 _dafny.Array
		_ = _nw36
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw36 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) bool = (func(_260_v0 bool) func(_dafny.Int) bool {
				return func(_261_i12 _dafny.Int) bool {
					return (_dafny.SetOf(true, _260_v0)).IsProperSubsetOf(_dafny.SetOf(_260_v0))
				}
			})(_179_v0)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw36 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw36).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw36).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_259_v60 = _nw36
		r0 = _259_v60
		r1 = _dafny.IntOfInt64(34)
		var _262_v61 _dafny.Map
		_ = _262_v61
		_262_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F15())
		r2 = ((_262_v61).Merge(_262_v61)).Merge(_262_v61)
		return r0, r1, r2
	}
}
func (_this *C1) M6(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map, *C0) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 *C0 = (*C0)(nil)
		_ = r2
		var _263_v1 _dafny.Sequence
		_ = _263_v1
		_263_v1 = _dafny.SeqOf(Companion_Default___.Fm4(func() _dafny.Set {
			var _coll15 = _dafny.NewBuilder()
			_ = _coll15
			for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-679), _dafny.IntOfInt64(733))); ; {
				_compr_15, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _264_v0 _dafny.Int
				_264_v0 = interface{}(_compr_15).(_dafny.Int)
				if ((_dafny.IntOfInt64(-679)).Cmp(_264_v0) <= 0) && ((_264_v0).Cmp(_dafny.IntOfInt64(733)) < 0) {
					_coll15.Add((_264_v0).Times(p1))
				}
			}
			return _coll15.ToSet()
		}(), globalState))
		if (_263_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool) {
			var _265_v2 _dafny.Array
			_ = _265_v2
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_5
			var _nw37 _dafny.Array
			_ = _nw37
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw37 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Map = (func(_266_p1 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_267_i0 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_p1, false)).Update(_266_p1, (_this).F14())
					}
				})(p1)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw37 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw37).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw37).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_265_v2 = _nw37
			_265_v2 = (func() _dafny.Array {
				if !((_this).F15()) {
					return _265_v2
				}
				return _265_v2
			})()
			var _268_v3 _dafny.Array
			_ = _268_v3
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw38
			_268_v3 = _nw38
			var _269_v4 _dafny.Sequence
			_ = _269_v4
			_269_v4 = _dafny.UnicodeSeqOfUtf8Bytes("xcoyciqnu")
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_268_v3), 0))
			_ = _index41
			(_268_v3).ArraySet1(_269_v4, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_268_v3), 0))
			_ = _index42
			(_268_v3).ArraySet1(_269_v4, (_index42).Int())
			var _270_v5 _dafny.Map
			_ = _270_v5
			_270_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			_270_v5 = (_270_v5).Merge(_270_v5)
			var _271_v6 _dafny.Sequence
			_ = _271_v6
			_271_v6 = _dafny.SeqOf((_268_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_268_v3), 0))).Int()).(_dafny.Sequence), (_268_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_268_v3), 0))).Int()).(_dafny.Sequence))
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_271_v6, _271_v6), _271_v6) {
				var _272_v7 _dafny.MultiSet
				_ = _272_v7
				_272_v7 = _dafny.MultiSetOf((_this).F14(), (_this).F15(), (_this).F14(), (_this).F15(), (_this).F15())
				var _273_v8 bool
				_ = _273_v8
				_273_v8 = false
				var _274_v9 _dafny.Set
				_ = _274_v9
				_274_v9 = _dafny.SetOf(_dafny.Companion_Sequence_.Equal((_268_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_268_v3), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("twrc")))
				var _275_v10 D5
				_ = _275_v10
				_275_v10 = Companion_D5_.Create_DC12_(_274_v9)
				var _rhs34 _dafny.MultiSet = _dafny.MultiSetFromSeq(Companion_Default___.Fm15(_dafny.CodePoint('w'), _dafny.IntOfInt64(767), globalState))
				_ = _rhs34
				var _rhs35 bool = !((_this).F14()) || ((_this).F14())
				_ = _rhs35
				var _rhs36 _dafny.Set = (_275_v10).Dtor_cf26()
				_ = _rhs36
				_272_v7 = _rhs34
				_273_v8 = _rhs35
				_274_v9 = _rhs36
				var _276_v11 _dafny.CodePoint
				_ = _276_v11
				_276_v11 = _dafny.CodePoint('d')
				var _277_v12 _dafny.Array
				_ = _277_v12
				var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw39
				_277_v12 = _nw39
				var _278_v13 _dafny.Array
				_ = _278_v13
				var _nwElement0_9 _dafny.Int = (_dafny.Zero).Minus(p1)
				_ = _nwElement0_9
				var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(11))
				_ = _nw40
				(_nw40).ArraySet1(_nwElement0_9, 0)
				(_nw40).ArraySet1((p1).Minus(p1), 1)
				(_nw40).ArraySet1(p1, 2)
				(_nw40).ArraySet1(p1, 3)
				(_nw40).ArraySet1(p1, 4)
				(_nw40).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(_273_v8, (_dafny.Zero).Minus(p1), _276_v11, p1, globalState), _269_v4)).Cardinality()), 5)
				(_nw40).ArraySet1(p1, 6)
				(_nw40).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), 7)
				(_nw40).ArraySet1(p1, 8)
				(_nw40).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_273_v8, _277_v12)).Cardinality(), 9)
				(_nw40).ArraySet1(p1, 10)
				_278_v13 = _nw40
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_278_v13), 0))
				_ = _index43
				(_278_v13).ArraySet1(p1, (_index43).Int())
				var _279_v14 D0
				_ = _279_v14
				_279_v14 = Companion_D0_.Create_DC1_((Companion_D1_.Create_DC5_(p1, (_this).F15(), p1, _273_v8, p1)).Dtor_cf11(), p1)
				var _280_v15 _dafny.Array
				_ = _280_v15
				var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw41
				_280_v15 = _nw41
				var _281_v16 D4
				_ = _281_v16
				_281_v16 = Companion_D4_.Create_DC11_(_279_v14, _280_v15, p1)
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_278_v13), 0))
				_ = _index44
				(_278_v13).ArraySet1((_281_v16).Dtor_cf25(), (_index44).Int())
				var _282_v17 *C0
				_ = _282_v17
				var _nw42 *C0 = New_C0_()
				_ = _nw42
				_nw42.Ctor__()
				_282_v17 = _nw42
				_276_v11 = _276_v11
				var _283_v18 _dafny.Sequence
				_ = _283_v18
				_283_v18 = _dafny.SeqOf((_dafny.Zero).Minus((_278_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_278_v13), 0))).Int()).(_dafny.Int)))
				(globalState).F1 = (((_this).F10()).Cardinality()).Minus(Companion_Default___.SafeModuloInt((_278_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_278_v13), 0))).Int()).(_dafny.Int), (_283_v18).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_283_v18).Cardinality()))).Uint32()).(_dafny.Int)))
			} else {
				var _284_v19 D1
				_ = _284_v19
				_284_v19 = Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("o")).Cardinality()), (_this).F14(), p1, (_this).F15(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F14())).Cardinality())
				var _285_v20 _dafny.Map
				_ = _285_v20
				_285_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_284_v19, (_this).F15())
				var _pat_let_tv15 = p1
				_ = _pat_let_tv15
				var _pat_let_tv16 = p1
				_ = _pat_let_tv16
				_285_v20 = (_285_v20).Update(func(_pat_let4_0 D1) D1 {
					return func(_286_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let5_0 _dafny.Int) D1 {
							return func(_287_dt__update_hcf10_h0 _dafny.Int) D1 {
								return func(_pat_let6_0 _dafny.Int) D1 {
									return func(_288_dt__update_hcf14_h0 _dafny.Int) D1 {
										return func(_pat_let7_0 bool) D1 {
											return func(_289_dt__update_hcf11_h0 bool) D1 {
												return Companion_D1_.Create_DC5_(_287_dt__update_hcf10_h0, _289_dt__update_hcf11_h0, (_286_dt__update__tmp_h0).Dtor_cf12(), (_286_dt__update__tmp_h0).Dtor_cf13(), _288_dt__update_hcf14_h0)
											}(_pat_let7_0)
										}((_this).F14())
									}(_pat_let6_0)
								}(_pat_let_tv16)
							}(_pat_let5_0)
						}(_pat_let_tv15)
					}(_pat_let4_0)
				}(_284_v19), (_this).F14())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_268_v3), 0))
				_ = _index45
				(_268_v3).ArraySet1((_268_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_268_v3), 0))).Int()).(_dafny.Sequence), (_index45).Int())
				var _290_v21 _dafny.Sequence
				_ = _290_v21
				_290_v21 = _dafny.SeqOf(_263_v1)
				var _291_v22 _dafny.CodePoint
				_ = _291_v22
				_291_v22 = _dafny.CodePoint('m')
				var _292_v23 _dafny.Map
				_ = _292_v23
				_292_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v22, (_this).F15())
				var _293_v24 bool
				_ = _293_v24
				var _294_v25 _dafny.MultiSet
				_ = _294_v25
				var _295_v26 _dafny.Set
				_ = _295_v26
				var _296_v27 _dafny.Int
				_ = _296_v27
				var _out16 bool
				_ = _out16
				var _out17 _dafny.MultiSet
				_ = _out17
				var _out18 _dafny.Set
				_ = _out18
				var _out19 _dafny.Int
				_ = _out19
				_out16, _out17, _out18, _out19 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_263_v1, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32(), false), _263_v1), (_290_v21).Select((Companion_Default___.SafeIndex((_292_v23).Cardinality(), _dafny.IntOfUint32((_290_v21).Cardinality()))).Uint32()).(_dafny.Sequence), globalState)
				_293_v24 = _out16
				_294_v25 = _out17
				_295_v26 = _out18
				_296_v27 = _out19
				_269_v4 = _dafny.Companion_Sequence_.Update(_269_v4, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_269_v4).Cardinality()), ((_this).F10()).Cardinality())), _dafny.IntOfUint32((_269_v4).Cardinality()))).Uint32(), _291_v22)
				_293_v24 = _293_v24
			}
			var _297_v28 *C0
			_ = _297_v28
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__()
			_297_v28 = _nw43
		} else {
			var _298_v29 *C0
			_ = _298_v29
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__()
			_298_v29 = _nw44
			var _299_v30 _dafny.Map
			_ = _299_v30
			_299_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F15())
			var _300_v31 _dafny.Set
			_ = _300_v31
			_300_v31 = _dafny.SetOf(p1)
			_299_v30 = (_299_v30).Update(Companion_Default___.SafeModuloInt(p1, p1), (func() bool {
				if !((_this).F14()) {
					return (_this).F14()
				}
				return Companion_Default___.Fm4(_300_v31, globalState)
			})())
			var _301_v32 bool
			_ = _301_v32
			_301_v32 = false
			_301_v32 = (_this).F15()
			var _302_v33 _dafny.CodePoint
			_ = _302_v33
			_302_v33 = _dafny.CodePoint('o')
			var _303_v34 _dafny.Set
			_ = _303_v34
			_303_v34 = _dafny.SetOf(_302_v33, _302_v33, _302_v33, _302_v33, _302_v33)
			var _304_v35 _dafny.Map
			_ = _304_v35
			_304_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(426), p1)
			var _305_v36 _dafny.Array
			_ = _305_v36
			var _nwElement0_10 _dafny.Int = p1
			_ = _nwElement0_10
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(20))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_10, 0)
			(_nw45).ArraySet1(p1, 1)
			(_nw45).ArraySet1(p1, 2)
			(_nw45).ArraySet1((_298_v29).Fm6(_301_v32, p1, globalState), 3)
			(_nw45).ArraySet1(p1, 4)
			(_nw45).ArraySet1(p1, 5)
			(_nw45).ArraySet1(p1, 6)
			(_nw45).ArraySet1(p1, 7)
			(_nw45).ArraySet1((func() _dafny.Int {
				if (_304_v35).Contains(p1) {
					return (_304_v35).Get(p1).(_dafny.Int)
				}
				return p1
			})(), 8)
			(_nw45).ArraySet1(_dafny.IntOfInt64(746), 9)
			(_nw45).ArraySet1(p1, 10)
			(_nw45).ArraySet1((_dafny.Zero).Minus((_298_v29).Fm6((_this).F15(), p1, globalState)), 11)
			(_nw45).ArraySet1(p1, 12)
			(_nw45).ArraySet1(p1, 13)
			(_nw45).ArraySet1((_298_v29).Fm6((_this).F14(), (_dafny.Zero).Minus(p1), globalState), 14)
			(_nw45).ArraySet1(_dafny.IntOfUint32((_263_v1).Cardinality()), 15)
			(_nw45).ArraySet1(p1, 16)
			(_nw45).ArraySet1(p1, 17)
			(_nw45).ArraySet1(p1, 18)
			(_nw45).ArraySet1(p1, 19)
			_305_v36 = _nw45
			var _306_v37 D0
			_ = _306_v37
			_306_v37 = Companion_D0_.Create_DC1_(!((_this).F14()), p1)
			var _307_v38 _dafny.Map
			_ = _307_v38
			_307_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _306_v37)
			var _308_v39 _dafny.Map
			_ = _308_v39
			_308_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v36, (_this).F15())
			var _309_v40 _dafny.Sequence
			_ = _309_v40
			_309_v40 = _dafny.SeqOf(_307_v38, _307_v38)
			var _pat_let_tv17 = _301_v32
			_ = _pat_let_tv17
			var _310_v41 _dafny.Array
			_ = _310_v41
			var _nwElement0_11 _dafny.Map = _307_v38
			_ = _nwElement0_11
			var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(15))
			_ = _nw46
			(_nw46).ArraySet1(_nwElement0_11, 0)
			(_nw46).ArraySet1(_307_v38, 1)
			(_nw46).ArraySet1(_307_v38, 2)
			(_nw46).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_308_v39).Contains(_305_v36) {
					return (_308_v39).Get(_305_v36).(bool)
				}
				return (_this).F15()
			})(), _306_v37), 3)
			(_nw46).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v32, func(_pat_let8_0 D0) D0 {
				return func(_311_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let9_0 bool) D0 {
						return func(_312_dt__update_hcf1_h0 bool) D0 {
							return Companion_D0_.Create_DC1_(_312_dt__update_hcf1_h0, (_311_dt__update__tmp_h1).Dtor_cf2())
						}(_pat_let9_0)
					}(_pat_let_tv17)
				}(_pat_let8_0)
			}(_306_v37)), 4)
			(_nw46).ArraySet1((_307_v38).Update(_301_v32, _306_v37), 5)
			(_nw46).ArraySet1(_307_v38, 6)
			(_nw46).ArraySet1(_307_v38, 7)
			(_nw46).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v32, _306_v37)).Update(_301_v32, _306_v37)).Update((_this).F15(), _306_v37), 8)
			(_nw46).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v32, _306_v37), 9)
			(_nw46).ArraySet1((_309_v40).Select((Companion_Default___.SafeIndex((_304_v35).Cardinality(), _dafny.IntOfUint32((_309_v40).Cardinality()))).Uint32()).(_dafny.Map), 10)
			(_nw46).ArraySet1(_307_v38, 11)
			(_nw46).ArraySet1(_307_v38, 12)
			(_nw46).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v32, Companion_D0_.Create_DC1_((_this).F14(), p1)), 13)
			(_nw46).ArraySet1(_307_v38, 14)
			_310_v41 = _nw46
			var _313_v42 D5
			_ = _313_v42
			_313_v42 = Companion_D5_.Create_DC13_((_dafny.Zero).Minus((p1).Times((_303_v34).Cardinality())), _305_v36, _310_v41, p1)
			var _source7 D5 = _313_v42
			_ = _source7
			if _source7.Is_DC13() {
				var _314___mcc_h0 _dafny.Int = _source7.Get_().(D5_DC13).Cf27
				_ = _314___mcc_h0
				var _315___mcc_h1 _dafny.Array = _source7.Get_().(D5_DC13).Cf28
				_ = _315___mcc_h1
				var _316___mcc_h2 _dafny.Array = _source7.Get_().(D5_DC13).Cf29
				_ = _316___mcc_h2
				var _317___mcc_h3 _dafny.Int = _source7.Get_().(D5_DC13).Cf30
				_ = _317___mcc_h3
				var _318_cf30 _dafny.Int = _317___mcc_h3
				_ = _318_cf30
				var _319_cf29 _dafny.Array = _316___mcc_h2
				_ = _319_cf29
				var _320_cf28 _dafny.Array = _315___mcc_h1
				_ = _320_cf28
				var _321_cf27 _dafny.Int = _314___mcc_h0
				_ = _321_cf27
				var _322_v43 _dafny.Array
				_ = _322_v43
				var _nwElement0_12 bool = (_this).F15()
				_ = _nwElement0_12
				var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(17))
				_ = _nw47
				(_nw47).ArraySet1(_nwElement0_12, 0)
				(_nw47).ArraySet1(!(_301_v32), 1)
				(_nw47).ArraySet1(false, 2)
				(_nw47).ArraySet1(_301_v32, 3)
				(_nw47).ArraySet1(_301_v32, 4)
				(_nw47).ArraySet1(_301_v32, 5)
				(_nw47).ArraySet1((_this).F14(), 6)
				(_nw47).ArraySet1(((_this).F15()) && (!(_301_v32)), 7)
				(_nw47).ArraySet1((_this).F15(), 8)
				(_nw47).ArraySet1((_this).F15(), 9)
				(_nw47).ArraySet1((_this).F14(), 10)
				(_nw47).ArraySet1(_301_v32, 11)
				(_nw47).ArraySet1((_this).F14(), 12)
				(_nw47).ArraySet1(_301_v32, 13)
				(_nw47).ArraySet1((_this).F14(), 14)
				(_nw47).ArraySet1((_this).F15(), 15)
				(_nw47).ArraySet1(_301_v32, 16)
				_322_v43 = _nw47
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_322_v43), 0))
				_ = _index46
				(_322_v43).ArraySet1((_this).F14(), (_index46).Int())
				var _323_v44 _dafny.Sequence
				_ = _323_v44
				_323_v44 = _dafny.UnicodeSeqOfUtf8Bytes("t")
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_322_v43), 0))
				_ = _index47
				(_322_v43).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_323_v44, _323_v44), (_index47).Int())
				var _324_v45 _dafny.Array
				_ = _324_v45
				var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
				_ = _nw48
				_324_v45 = _nw48
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_324_v45), 0))
				_ = _index48
				(_324_v45).ArraySet1CodePoint(_dafny.CodePoint('c'), (_index48).Int())
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_324_v45), 0))
				_ = _index49
				(_324_v45).ArraySet1CodePoint(_302_v33, (_index49).Int())
				var _325_v46 _dafny.Map
				_ = _325_v46
				_325_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_263_v1, true)
				_299_v30 = (_299_v30).Update(_321_cf27, (func() bool {
					if (_325_v46).Contains(_263_v1) {
						return (_325_v46).Get(_263_v1).(bool)
					}
					return _301_v32
				})())
				var _326_v47 _dafny.Array
				_ = _326_v47
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_6
				var _nw49 _dafny.Array
				_ = _nw49
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw49 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Sequence = func(_327_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_this).F15())
					}
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw49 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw49).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw49).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_326_v47 = _nw49
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_326_v47), 0))
				_ = _index50
				(_326_v47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_263_v1, _263_v1), (_index50).Int())
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_326_v47), 0))
				_ = _index51
				(_326_v47).ArraySet1(_263_v1, (_index51).Int())
			} else if _source7.Is_DC14() {
				var _328___mcc_h4 _dafny.Array = _source7.Get_().(D5_DC14).Cf31
				_ = _328___mcc_h4
				var _329___mcc_h5 _dafny.Int = _source7.Get_().(D5_DC14).Cf32
				_ = _329___mcc_h5
				var _330___mcc_h6 _dafny.Int = _source7.Get_().(D5_DC14).Cf33
				_ = _330___mcc_h6
				var _331___mcc_h7 _dafny.Map = _source7.Get_().(D5_DC14).Cf34
				_ = _331___mcc_h7
				var _332___mcc_h8 _dafny.Set = _source7.Get_().(D5_DC14).Cf35
				_ = _332___mcc_h8
				var _333_cf35 _dafny.Set = _332___mcc_h8
				_ = _333_cf35
				var _334_cf34 _dafny.Map = _331___mcc_h7
				_ = _334_cf34
				var _335_cf33 _dafny.Int = _330___mcc_h6
				_ = _335_cf33
				var _336_cf32 _dafny.Int = _329___mcc_h5
				_ = _336_cf32
				var _337_cf31 _dafny.Array = _328___mcc_h4
				_ = _337_cf31
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_305_v36), 0))
				_ = _index52
				(_305_v36).ArraySet1((func() _dafny.Int {
					if (_this).F14() {
						return _336_cf32
					}
					return (_300_v31).Cardinality()
				})(), (_index52).Int())
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_305_v36), 0))
				_ = _index53
				(_305_v36).ArraySet1((p1).Times(_dafny.IntOfInt64(-398)), (_index53).Int())
				_300_v31 = (_300_v31).Union(func() _dafny.Set {
					var _coll16 = _dafny.NewBuilder()
					_ = _coll16
					for _iter17 := _dafny.Iterate((_300_v31).Elements()); ; {
						_compr_16, _ok17 := _iter17()
						if !_ok17 {
							break
						}
						var _338_v48 _dafny.Int
						_338_v48 = interface{}(_compr_16).(_dafny.Int)
						if (_300_v31).Contains(_338_v48) {
							_coll16.Add((_338_v48).Times(_dafny.IntOfInt64(721)))
						}
					}
					return _coll16.ToSet()
				}())
				_301_v32 = true
				var _339_v49 _dafny.Array
				_ = _339_v49
				var _nw50 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(15))
				_ = _nw50
				_339_v49 = _nw50
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_339_v49), 0))
				_ = _index54
				(_339_v49).ArraySet1(_306_v37, (_index54).Int())
				var _340_v50 D0
				_ = _340_v50
				_340_v50 = Companion_D0_.Create_DC0_(_302_v33)
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_339_v49), 0))
				_ = _index55
				(_339_v49).ArraySet1(Companion_Default___.Fm16(_340_v50, globalState), (_index55).Int())
			} else if _source7.Is_DC12() {
				var _341___mcc_h9 _dafny.Set = _source7.Get_().(D5_DC12).Cf26
				_ = _341___mcc_h9
				var _342_cf26 _dafny.Set = _341___mcc_h9
				_ = _342_cf26
				var _343_v51 _dafny.Sequence
				_ = _343_v51
				_343_v51 = _dafny.UnicodeSeqOfUtf8Bytes("uhkhhuwaq")
				_343_v51 = _343_v51
				(globalState).F1 = ((_dafny.Zero).Minus(p1)).Times(p1)
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_305_v36), 0))
				_ = _index56
				(_305_v36).ArraySet1((_300_v31).Cardinality(), (_index56).Int())
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_305_v36), 0))
				_ = _index57
				var _rhs37 bool = (_this).F14()
				_ = _rhs37
				var _rhs38 _dafny.Int = ((_dafny.MultiSetOf(false)).Difference(((_this).F10()).Difference(_dafny.MultiSetOf(Companion_Default___.Fm4(_300_v31, globalState))))).Cardinality()
				_ = _rhs38
				var _lhs25 _dafny.Array = _305_v36
				_ = _lhs25
				var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_305_v36), 0))
				_ = _lhs26
				_301_v32 = _rhs37
				(_lhs25).ArraySet1(_rhs38, (_lhs26).Int())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_305_v36), 0))
				_ = _index58
				(_305_v36).ArraySet1(p1, (_index58).Int())
			} else {
				var _344___mcc_h10 D5 = _source7.Get_().(D5_DC15).Cf36
				_ = _344___mcc_h10
				var _345_cf36 D5 = _344___mcc_h10
				_ = _345_cf36
				var _346_v52 _dafny.Array
				_ = _346_v52
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_7
				var _nw51 _dafny.Array
				_ = _nw51
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw51 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) bool = func(_347_i2 _dafny.Int) bool {
						return (_this).F14()
					}
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw51 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw51).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw51).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_346_v52 = _nw51
				(globalState).F5 = (func() _dafny.Array {
					if _301_v32 {
						return (func() _dafny.Array {
							if (_this).F15() {
								return _346_v52
							}
							return _346_v52
						})()
					}
					return _346_v52
				})()
				(globalState).F1 = p1
				var _348_v53 _dafny.Sequence
				_ = _348_v53
				_348_v53 = _dafny.UnicodeSeqOfUtf8Bytes("jfclgutmi")
				_348_v53 = _348_v53
				_301_v32 = ((p1).Cmp(p1) != 0) && ((_263_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool))
			}
			var _349_v54 _dafny.Sequence
			_ = _349_v54
			_349_v54 = _dafny.UnicodeSeqOfUtf8Bytes("iexqde")
			var _350_v55 _dafny.MultiSet
			_ = _350_v55
			_350_v55 = _dafny.MultiSetOf(p1, p1, _dafny.IntOfUint32((_349_v54).Cardinality()))
			var _351_v56 _dafny.Set
			_ = _351_v56
			_351_v56 = _dafny.SetOf((_this).F15())
			var _352_v57 _dafny.Sequence
			_ = _352_v57
			_352_v57 = _dafny.SeqOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(p1, p1, (func() _dafny.Int {
				if (_304_v35).Contains(p1) {
					return (_304_v35).Get(p1).(_dafny.Int)
				}
				return (_351_v56).Cardinality()
			})(), p1)).Cardinality()), p1), (_351_v56).Cardinality(), p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(globalState), _349_v54)).Cardinality(), p1)
			_350_v55 = _dafny.MultiSetFromSeq(_352_v57)
		}
		(globalState).F1 = p1
		var _353_v58 _dafny.Map
		_ = _353_v58
		_353_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _263_v1)
		var _354_v59 _dafny.Set
		_ = _354_v59
		_354_v59 = _dafny.SetOf(p1, p1)
		var _355_v60 _dafny.CodePoint
		_ = _355_v60
		_355_v60 = _dafny.CodePoint('i')
		var _356_v61 _dafny.Sequence
		_ = _356_v61
		_356_v61 = _dafny.UnicodeSeqOfUtf8Bytes("r")
		var _357_v62 _dafny.MultiSet
		_ = _357_v62
		_357_v62 = _dafny.MultiSetOf(p1, p1)
		var _rhs39 _dafny.Int = Companion_Default___.SafeModuloInt(p1, p1)
		_ = _rhs39
		var _rhs40 _dafny.Sequence = (func() _dafny.Sequence {
			if Companion_Default___.Fm4(_354_v59, globalState) {
				return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm15(_355_v60, p1, globalState), _263_v1)
			}
			return _263_v1
		})()
		_ = _rhs40
		var _rhs41 _dafny.Map = (_353_v58).Update((_this).F14(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_263_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_356_v61).Cardinality()), _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32(), (_this).F14()), (Companion_Default___.SafeIndex((_357_v62).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_263_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_356_v61).Cardinality()), _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32(), (_this).F14())).Cardinality()))).Uint32(), (_this).F15()), _263_v1))
		_ = _rhs41
		var _lhs27 *GlobalState = globalState
		_ = _lhs27
		_lhs27.F1 = _rhs39
		_263_v1 = _rhs40
		_353_v58 = _rhs41
		var _358_v63 _dafny.Set
		_ = _358_v63
		_358_v63 = _dafny.SetOf((_this).F15())
		var _hi4 _dafny.Int = _dafny.IntOfInt64(427)
		_ = _hi4
		for _359_i3 := ((_358_v63).Difference(_358_v63)).Cardinality(); _359_i3.Cmp(_hi4) < 0; _359_i3 = _359_i3.Plus(_dafny.One) {
			var _360_v64 *C0
			_ = _360_v64
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__()
			_360_v64 = _nw52
			var _361_v65 D4
			_ = _361_v65
			_361_v65 = Companion_D4_.Create_DC10_(_360_v64)
			var _362_v67 D6
			_ = _362_v67
			_362_v67 = Companion_D6_.Create_DC16_(func() _dafny.Set {
				var _coll17 = _dafny.NewBuilder()
				_ = _coll17
				for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(490), _dafny.IntOfInt64(676))); ; {
					_compr_17, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _363_v66 _dafny.Int
					_363_v66 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(490)).Cmp(_363_v66) <= 0) && ((_363_v66).Cmp(_dafny.IntOfInt64(676)) < 0) {
						_coll17.Add((_363_v66).Minus(_dafny.IntOfUint32((_356_v61).Cardinality())))
					}
				}
				return _coll17.ToSet()
			}())
			_361_v65 = (func() D4 {
				if Companion_Default___.Fm4((_362_v67).Dtor_cf37(), globalState) {
					return _361_v65
				}
				return _361_v65
			})()
			var _364_v68 *C0
			_ = _364_v68
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__()
			_364_v68 = _nw53
			if (_this).F14() {
				_263_v1 = _263_v1
				var _365_v69 _dafny.Array
				_ = _365_v69
				var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
				_ = _nw54
				_365_v69 = _nw54
				var _366_v70 _dafny.Map
				_ = _366_v70
				_366_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ey"))).Cardinality()), (_this).F10())
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_365_v69), 0))
				_ = _index59
				(_365_v69).ArraySet1(_366_v70, (_index59).Int())
				var _367_v71 _dafny.Sequence
				_ = _367_v71
				_367_v71 = _dafny.SeqOf(p1, _359_i3)
				var _368_v72 D1
				_ = _368_v72
				_368_v72 = Companion_D1_.Create_DC4_((_this).F15(), false, _359_i3, _367_v71, p1)
				var _369_v74 _dafny.Array
				_ = _369_v74
				var _nwElement0_13 bool = Companion_Default___.Fm4(func() _dafny.Set {
					var _coll18 = _dafny.NewBuilder()
					_ = _coll18
					for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(771), _dafny.IntOfInt64(-659))); ; {
						_compr_18, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _370_v73 _dafny.Int
						_370_v73 = interface{}(_compr_18).(_dafny.Int)
						if ((_dafny.IntOfInt64(771)).Cmp(_370_v73) <= 0) && ((_370_v73).Cmp(_dafny.IntOfInt64(-659)) < 0) {
							_coll18.Add((_370_v73).Times(_359_i3))
						}
					}
					return _coll18.ToSet()
				}(), globalState)
				_ = _nwElement0_13
				var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(14))
				_ = _nw55
				(_nw55).ArraySet1(_nwElement0_13, 0)
				(_nw55).ArraySet1(!((_this).F14()), 1)
				(_nw55).ArraySet1(!((_this).F14()), 2)
				(_nw55).ArraySet1((_this).F15(), 3)
				(_nw55).ArraySet1((_this).F14(), 4)
				(_nw55).ArraySet1((_this).F15(), 5)
				(_nw55).ArraySet1((_this).F14(), 6)
				(_nw55).ArraySet1((_this).F15(), 7)
				(_nw55).ArraySet1((_this).F15(), 8)
				(_nw55).ArraySet1((_this).F14(), 9)
				(_nw55).ArraySet1((_this).F14(), 10)
				(_nw55).ArraySet1((_this).F15(), 11)
				(_nw55).ArraySet1((_this).F14(), 12)
				(_nw55).ArraySet1((_this).F14(), 13)
				_369_v74 = _nw55
				var _371_v75 _dafny.Map
				_ = _371_v75
				_371_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F14()), _369_v74)
				var _372_v76 _dafny.Sequence
				_ = _372_v76
				_372_v76 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_368_v72).Dtor_cf6(), _369_v74), _371_v75, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F15()), _369_v74), _371_v75)
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_365_v69), 0))
				_ = _index60
				(_365_v69).ArraySet1(Companion_Default___.Fm17(_dafny.MultiSetOf(_355_v60), ((_372_v76).Select((Companion_Default___.SafeIndex(_359_i3, _dafny.IntOfUint32((_372_v76).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), globalState), (_index60).Int())
				var _373_v77 _dafny.Map
				_ = _373_v77
				_373_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jtyk"), _356_v61), (_this).F14())
				_373_v77 = (_373_v77).Update((_this).F15(), (p1).Cmp(_359_i3) != 0)
				var _374_v78 bool
				_ = _374_v78
				_374_v78 = false
				_374_v78 = !((func() bool {
					if (_373_v77).Contains((_this).F15()) {
						return (_373_v77).Get((_this).F15()).(bool)
					}
					return (_354_v59).IsSubsetOf(_354_v59)
				})())
				var _rhs42 _dafny.Int = p1
				_ = _rhs42
				var _rhs43 _dafny.Sequence = _263_v1
				_ = _rhs43
				var _lhs28 *GlobalState = globalState
				_ = _lhs28
				_lhs28.F1 = _rhs42
				_263_v1 = _rhs43
			} else {
				(globalState).F1 = (((_this).F10()).Union((_dafny.MultiSetFromSeq(_263_v1)).Union((_this).F10()))).Cardinality()
				var _375_v79 *C0
				_ = _375_v79
				var _nw56 *C0 = New_C0_()
				_ = _nw56
				_nw56.Ctor__()
				_375_v79 = _nw56
				var _376_v80 *C0
				_ = _376_v80
				var _nw57 *C0 = New_C0_()
				_ = _nw57
				_nw57.Ctor__()
				_376_v80 = _nw57
				(globalState).F1 = (func() _dafny.Int {
					if (_357_v62).Contains(_359_i3) {
						return (_357_v62).Multiplicity(_359_i3)
					}
					return _359_i3
				})()
				var _377_v82 _dafny.Array
				_ = _377_v82
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_8
				var _nw58 _dafny.Array
				_ = _nw58
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw58 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Int = (func(_378_i3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_379_i4 _dafny.Int) _dafny.Int {
							return (_379_i4).Times((func() _dafny.Map {
								var _coll19 = _dafny.NewMapBuilder()
								_ = _coll19
								for _iter20 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F15()), _378_i3)).Keys().Elements()); ; {
									_compr_19, _ok20 := _iter20()
									if !_ok20 {
										break
									}
									var _380_v81 _dafny.Map
									_380_v81 = interface{}(_compr_19).(_dafny.Map)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F15()), _378_i3)).Contains(_380_v81) {
										_coll19.Add(_380_v81, (_this).F14())
									}
								}
								return _coll19.ToMap()
							}()).Cardinality())
						}
					})(_359_i3)
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
				_377_v82 = _nw58
				var _381_v83 _dafny.Map
				_ = _381_v83
				_381_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_359_i3, ((_this).F10()).Cardinality())
				var _382_v84 _dafny.Map
				_ = _382_v84
				_382_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_381_v83).Cardinality())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_377_v82), 0))
				_ = _index61
				(_377_v82).ArraySet1(((_381_v83).Update(p1, (func() _dafny.Int {
					if (_381_v83).Contains(_359_i3) {
						return (_381_v83).Get(_359_i3).(_dafny.Int)
					}
					return (_381_v83).Cardinality()
				})())).Cardinality(), (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_377_v82), 0))
				_ = _index62
				(_377_v82).ArraySet1(((_360_v64).Fm6((_this).F15(), p1, globalState)).Times(_dafny.IntOfUint32((_356_v61).Cardinality())), (_index62).Int())
			}
			var _383_v85 bool
			_ = _383_v85
			_383_v85 = true
			_383_v85 = !((func() bool {
				if _383_v85 {
					return true
				}
				return _383_v85
			})())
		}
		var _384_v86 D0
		_ = _384_v86
		_384_v86 = Companion_D0_.Create_DC0_(_355_v60)
		if func(_source8 D0) bool {
			if _source8.Is_DC1() {
				var _385___mcc_h11 bool = _source8.Get_().(D0_DC1).Cf1
				_ = _385___mcc_h11
				var _386___mcc_h12 _dafny.Int = _source8.Get_().(D0_DC1).Cf2
				_ = _386___mcc_h12
				var _387_cf2 _dafny.Int = _386___mcc_h12
				_ = _387_cf2
				var _388_cf1 bool = _385___mcc_h11
				_ = _388_cf1
				return false
			} else if _source8.Is_DC0() {
				var _389___mcc_h13 _dafny.CodePoint = _source8.Get_().(D0_DC0).Cf0
				_ = _389___mcc_h13
				var _390_cf0 _dafny.CodePoint = _389___mcc_h13
				_ = _390_cf0
				return (_this).F15()
			} else {
				var _391___mcc_h14 D0 = _source8.Get_().(D0_DC2).Cf3
				_ = _391___mcc_h14
				var _392_cf3 D0 = _391___mcc_h14
				_ = _392_cf3
				return true
			}
		}(_384_v86) {
			if _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wny"), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("q"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality()))).Uint32(), _355_v60)), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wny"), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("q"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality()))).Uint32(), _355_v60))).Cardinality()))).Uint32(), _355_v60), _356_v61) {
				var _393_v87 *C0
				_ = _393_v87
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__()
				_393_v87 = _nw59
				var _394_v88 _dafny.Map
				_ = _394_v88
				_394_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _393_v87)
				var _395_v89 _dafny.Array
				_ = _395_v89
				var _nwElement0_14 *C0 = _393_v87
				_ = _nwElement0_14
				var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(28))
				_ = _nw60
				(_nw60).ArraySet1(_nwElement0_14, 0)
				(_nw60).ArraySet1((func() *C0 {
					if (_this).F15() {
						return _393_v87
					}
					return _393_v87
				})(), 1)
				(_nw60).ArraySet1(_393_v87, 2)
				(_nw60).ArraySet1(_393_v87, 3)
				(_nw60).ArraySet1(_393_v87, 4)
				(_nw60).ArraySet1((func() *C0 {
					if (_394_v88).Contains(p1) {
						return (_394_v88).Get(p1).(*C0)
					}
					return _393_v87
				})(), 5)
				(_nw60).ArraySet1(_393_v87, 6)
				(_nw60).ArraySet1(_393_v87, 7)
				(_nw60).ArraySet1(_393_v87, 8)
				(_nw60).ArraySet1(_393_v87, 9)
				(_nw60).ArraySet1(_393_v87, 10)
				(_nw60).ArraySet1(_393_v87, 11)
				(_nw60).ArraySet1(_393_v87, 12)
				(_nw60).ArraySet1(_393_v87, 13)
				(_nw60).ArraySet1(_393_v87, 14)
				(_nw60).ArraySet1(_393_v87, 15)
				(_nw60).ArraySet1(_393_v87, 16)
				(_nw60).ArraySet1(_393_v87, 17)
				(_nw60).ArraySet1(_393_v87, 18)
				(_nw60).ArraySet1(_393_v87, 19)
				(_nw60).ArraySet1(_393_v87, 20)
				(_nw60).ArraySet1(_393_v87, 21)
				(_nw60).ArraySet1(_393_v87, 22)
				(_nw60).ArraySet1((func() *C0 {
					if (_this).F15() {
						return _393_v87
					}
					return _393_v87
				})(), 23)
				(_nw60).ArraySet1(_393_v87, 24)
				(_nw60).ArraySet1(_393_v87, 25)
				(_nw60).ArraySet1(_393_v87, 26)
				(_nw60).ArraySet1(_393_v87, 27)
				_395_v89 = _nw60
				_395_v89 = _395_v89
				var _396_v90 _dafny.Sequence
				_ = _396_v90
				_396_v90 = _dafny.SeqOf(p1, p1)
				var _397_v91 _dafny.Sequence
				_ = _397_v91
				_397_v91 = _dafny.SeqOf(_358_v63, _358_v63, _358_v63, _358_v63, _358_v63)
				var _398_v94 _dafny.Array
				_ = _398_v94
				var _nwElement0_15 _dafny.Int = (p1).Times(p1)
				_ = _nwElement0_15
				var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(20))
				_ = _nw61
				(_nw61).ArraySet1(_nwElement0_15, 0)
				(_nw61).ArraySet1((_dafny.Zero).Minus(p1), 1)
				(_nw61).ArraySet1(p1, 2)
				(_nw61).ArraySet1(p1, 3)
				(_nw61).ArraySet1((_396_v90).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_396_v90).Cardinality()))).Uint32()).(_dafny.Int), 4)
				(_nw61).ArraySet1(p1, 5)
				(_nw61).ArraySet1(p1, 6)
				(_nw61).ArraySet1((_dafny.Zero).Minus((Companion_Default___.Fm16(_384_v86, globalState)).Dtor_cf2()), 7)
				(_nw61).ArraySet1(((_397_v91).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_397_v91).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), 8)
				(_nw61).ArraySet1(p1, 9)
				(_nw61).ArraySet1((_dafny.Zero).Minus(p1), 10)
				(_nw61).ArraySet1((func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(861), _dafny.IntOfInt64(-298))); ; {
						_compr_20, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _399_v92 _dafny.Int
						_399_v92 = interface{}(_compr_20).(_dafny.Int)
						if ((_dafny.IntOfInt64(861)).Cmp(_399_v92) <= 0) && ((_399_v92).Cmp(_dafny.IntOfInt64(-298)) < 0) {
							_coll20.Add((_399_v92).Plus(p1), (Companion_D0_.Create_DC1_((_this).F15(), p1)).Dtor_cf2())
						}
					}
					return _coll20.ToMap()
				}()).Cardinality(), 11)
				(_nw61).ArraySet1(p1, 12)
				(_nw61).ArraySet1(p1, 13)
				(_nw61).ArraySet1(p1, 14)
				(_nw61).ArraySet1((_dafny.Zero).Minus(((_this).F10()).Cardinality()), 15)
				(_nw61).ArraySet1(_dafny.IntOfUint32((_356_v61).Cardinality()), 16)
				(_nw61).ArraySet1(Companion_Default___.SafeDivisionInt(p1, (Companion_Default___.Fm18((_this).F15(), func() _dafny.Map {
					var _coll21 = _dafny.NewMapBuilder()
					_ = _coll21
					for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(830), _dafny.IntOfInt64(26))); ; {
						_compr_21, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _400_v93 _dafny.Int
						_400_v93 = interface{}(_compr_21).(_dafny.Int)
						if ((_dafny.IntOfInt64(830)).Cmp(_400_v93) <= 0) && ((_400_v93).Cmp(_dafny.IntOfInt64(26)) < 0) {
							_coll21.Add((_400_v93).Times((_358_v63).Cardinality()), (_this).F14())
						}
					}
					return _coll21.ToMap()
				}(), globalState)).Cardinality()), 17)
				(_nw61).ArraySet1((p1).Minus(p1), 18)
				(_nw61).ArraySet1((_393_v87).Fm6((_this).F15(), p1, globalState), 19)
				_398_v94 = _nw61
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_398_v94), 0))
				_ = _index63
				(_398_v94).ArraySet1(p1, (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_398_v94), 0))
				_ = _index64
				var _rhs44 _dafny.CodePoint = _355_v60
				_ = _rhs44
				var _rhs45 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_396_v90, _396_v90)).Cardinality())).Plus(_dafny.IntOfInt64(-840))
				_ = _rhs45
				var _rhs46 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if (_this).F15() {
						return (_dafny.Zero).Minus(_dafny.IntOfInt64(-26))
					}
					return p1
				})()))
				_ = _rhs46
				var _rhs47 _dafny.Int = (func() _dafny.Int {
					if true {
						return p1
					}
					return p1
				})()
				_ = _rhs47
				var _rhs48 _dafny.Int = p1
				_ = _rhs48
				var _lhs29 _dafny.Array = _398_v94
				_ = _lhs29
				var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_398_v94), 0))
				_ = _lhs30
				var _lhs31 *GlobalState = globalState
				_ = _lhs31
				var _lhs32 *GlobalState = globalState
				_ = _lhs32
				var _lhs33 *GlobalState = globalState
				_ = _lhs33
				_355_v60 = _rhs44
				(_lhs29).ArraySet1(_rhs45, (_lhs30).Int())
				_lhs31.F1 = _rhs46
				_lhs32.F1 = _rhs47
				_lhs33.F1 = _rhs48
				var _401_v95 bool
				_ = _401_v95
				_401_v95 = false
				var _402_v96 _dafny.Array
				_ = _402_v96
				var _nw62 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw62
				_402_v96 = _nw62
				var _403_v97 _dafny.Array
				_ = _403_v97
				var _nwElement0_16 _dafny.Array = _402_v96
				_ = _nwElement0_16
				var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(13))
				_ = _nw63
				(_nw63).ArraySet1(_nwElement0_16, 0)
				(_nw63).ArraySet1(_402_v96, 1)
				(_nw63).ArraySet1(_402_v96, 2)
				(_nw63).ArraySet1(_402_v96, 3)
				(_nw63).ArraySet1(_402_v96, 4)
				(_nw63).ArraySet1(_402_v96, 5)
				(_nw63).ArraySet1(_402_v96, 6)
				(_nw63).ArraySet1(_402_v96, 7)
				(_nw63).ArraySet1(_402_v96, 8)
				(_nw63).ArraySet1(_402_v96, 9)
				(_nw63).ArraySet1(_402_v96, 10)
				(_nw63).ArraySet1(_402_v96, 11)
				(_nw63).ArraySet1(_402_v96, 12)
				_403_v97 = _nw63
				var _404_v98 _dafny.Array
				_ = _404_v98
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_9
				var _nw64 _dafny.Array
				_ = _nw64
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw64 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.CodePoint = (func(_405_v60 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_406_i5 _dafny.Int) _dafny.CodePoint {
							return _405_v60
						}
					})(_355_v60)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw64 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw64).ArraySet1CodePoint(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw64).ArraySet1CodePoint(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_404_v98 = _nw64
				var _407_v99 _dafny.Sequence
				_ = _407_v99
				_407_v99 = _dafny.SeqOf(_404_v98, _404_v98, _404_v98)
				var _rhs49 bool = (_dafny.SetOf((_this).F15(), (_this).F14())).IsProperSubsetOf(_358_v63)
				_ = _rhs49
				var _rhs50 _dafny.Int = p1
				_ = _rhs50
				var _rhs51 _dafny.Array = _403_v97
				_ = _rhs51
				var _rhs52 bool = (_393_v87).Fm5(_358_v63, _dafny.IntOfUint32((_407_v99).Cardinality()), (p1).Plus(p1), globalState)
				_ = _rhs52
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				_401_v95 = _rhs49
				_lhs34.F1 = _rhs50
				_403_v97 = _rhs51
				_401_v95 = _rhs52
				var _408_v100 _dafny.Array
				_ = _408_v100
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_10
				var _nw65 _dafny.Array
				_ = _nw65
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw65 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Sequence = (func(_409_v61 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_410_i6 _dafny.Int) _dafny.Sequence {
							return _409_v61
						}
					})(_356_v61)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw65 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw65).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw65).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_408_v100 = _nw65
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_403_v97), 0))
				_ = _index65
				(_403_v97).ArraySet1(_402_v96, (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_403_v97), 0))
				_ = _index66
				var _rhs53 _dafny.Array = _402_v96
				_ = _rhs53
				var _rhs54 _dafny.Array = _408_v100
				_ = _rhs54
				var _rhs55 _dafny.Array = _402_v96
				_ = _rhs55
				var _lhs35 *GlobalState = globalState
				_ = _lhs35
				var _lhs36 _dafny.Array = _403_v97
				_ = _lhs36
				var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_403_v97), 0))
				_ = _lhs37
				_lhs35.F5 = _rhs53
				_408_v100 = _rhs54
				(_lhs36).ArraySet1(_rhs55, (_lhs37).Int())
			} else {
				var _411_v101 _dafny.Array
				_ = _411_v101
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_11
				var _nw66 _dafny.Array
				_ = _nw66
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw66 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) bool = func(_412_i7 _dafny.Int) bool {
						return (_this).F15()
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
				_411_v101 = _nw66
				var _413_v102 _dafny.MultiSet
				_ = _413_v102
				_413_v102 = _dafny.MultiSetOf(_411_v101, _411_v101)
				(globalState).F1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(p1, ((_413_v102).Update(_411_v101, Companion_Default___.Abs(p1))).Cardinality()), p1)
				var _414_v103 bool
				_ = _414_v103
				_414_v103 = false
				_414_v103 = (_this).F14()
				var _415_v104 _dafny.Map
				_ = _415_v104
				_415_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), Companion_Default___.Fm19(_355_v60, (_this).F14(), globalState))).Cardinality())
				var _416_v105 _dafny.Map
				_ = _416_v105
				_416_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _417_v106 _dafny.Sequence
				_ = _417_v106
				_417_v106 = _dafny.SeqOf(p1)
				var _418_v107 _dafny.Array
				_ = _418_v107
				var _nwElement0_17 _dafny.Int = p1
				_ = _nwElement0_17
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(29))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_17, 0)
				(_nw67).ArraySet1(p1, 1)
				(_nw67).ArraySet1(_dafny.IntOfInt64(-308), 2)
				(_nw67).ArraySet1(p1, 3)
				(_nw67).ArraySet1(p1, 4)
				(_nw67).ArraySet1(p1, 5)
				(_nw67).ArraySet1(p1, 6)
				(_nw67).ArraySet1(_dafny.IntOfUint32((_356_v61).Cardinality()), 7)
				(_nw67).ArraySet1(p1, 8)
				(_nw67).ArraySet1(p1, 9)
				(_nw67).ArraySet1((_354_v59).Cardinality(), 10)
				(_nw67).ArraySet1(p1, 11)
				(_nw67).ArraySet1(p1, 12)
				(_nw67).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).F14())).Cardinality()), 13)
				(_nw67).ArraySet1(p1, 14)
				(_nw67).ArraySet1(((_this).F10()).Cardinality(), 15)
				(_nw67).ArraySet1(p1, 16)
				(_nw67).ArraySet1((_415_v104).Cardinality(), 17)
				(_nw67).ArraySet1(p1, 18)
				(_nw67).ArraySet1(p1, 19)
				(_nw67).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(308))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}((func(_419_v60 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_420_i8 _dafny.Int) _dafny.CodePoint {
						return _419_v60
					}
				})(_355_v60)))).Cardinality()), 20)
				(_nw67).ArraySet1(p1, 21)
				(_nw67).ArraySet1(_dafny.IntOfUint32((_356_v61).Cardinality()), 22)
				(_nw67).ArraySet1(_dafny.IntOfUint32((_356_v61).Cardinality()), 23)
				(_nw67).ArraySet1(_dafny.IntOfUint32((_263_v1).Cardinality()), 24)
				(_nw67).ArraySet1(p1, 25)
				(_nw67).ArraySet1((func() _dafny.Int {
					if (_416_v105).Contains(p1) {
						return (_416_v105).Get(p1).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(Companion_Default___.Fm19(_355_v60, (_this).F14(), globalState))
				})(), 26)
				(_nw67).ArraySet1(p1, 27)
				(_nw67).ArraySet1(_dafny.IntOfUint32((_417_v106).Cardinality()), 28)
				_418_v107 = _nw67
				var _421_v108 _dafny.MultiSet
				_ = _421_v108
				_421_v108 = _dafny.MultiSetOf(_418_v107)
				var _422_v109 _dafny.Map
				_ = _422_v109
				_422_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(p1))).Cardinality()), _421_v108)
				var _423_v110 _dafny.Map
				_ = _423_v110
				_423_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _418_v107)
				var _424_v111 _dafny.Array
				_ = _424_v111
				var _nwElement0_18 _dafny.MultiSet = _421_v108
				_ = _nwElement0_18
				var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(9))
				_ = _nw68
				(_nw68).ArraySet1(_nwElement0_18, 0)
				(_nw68).ArraySet1(_421_v108, 1)
				(_nw68).ArraySet1(_dafny.MultiSetOf(_418_v107), 2)
				(_nw68).ArraySet1(_421_v108, 3)
				(_nw68).ArraySet1((func() _dafny.MultiSet {
					if (_422_v109).Contains(p1) {
						return (_422_v109).Get(p1).(_dafny.MultiSet)
					}
					return _421_v108
				})(), 4)
				(_nw68).ArraySet1(_421_v108, 5)
				(_nw68).ArraySet1(_421_v108, 6)
				(_nw68).ArraySet1(_dafny.MultiSetOf((func() _dafny.Array {
					if (_423_v110).Contains(_414_v103) {
						return (_423_v110).Get(_414_v103).(_dafny.Array)
					}
					return _418_v107
				})()), 7)
				(_nw68).ArraySet1(_421_v108, 8)
				_424_v111 = _nw68
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_424_v111), 0))
				_ = _index67
				(_424_v111).ArraySet1((_421_v108).Union(_421_v108), (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_424_v111), 0))
				_ = _index68
				(_424_v111).ArraySet1(_421_v108, (_index68).Int())
				_356_v61 = _dafny.UnicodeSeqOfUtf8Bytes("t")
				(globalState).F1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-725), (_415_v104).Cardinality())
			}
			var _425_v112 bool
			_ = _425_v112
			_425_v112 = true
			var _426_v113 D6
			_ = _426_v113
			_426_v113 = Companion_D6_.Create_DC17_(p1, _356_v61, (_this).F14())
			_425_v112 = !(((func() D6 {
				if (_263_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool) {
					return _426_v113
				}
				return _426_v113
			})()).Dtor_cf40())
			var _427_v114 _dafny.Map
			_ = _427_v114
			_427_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool), p1)
			_427_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_425_v112), (p1).Minus(p1))
			var _428_v116 D5
			_ = _428_v116
			_428_v116 = Companion_D5_.Create_DC12_(_358_v63)
			var _rhs56 _dafny.Int = Companion_Default___.Fm19(_355_v60, (_this).F15(), globalState)
			_ = _rhs56
			var _rhs57 _dafny.Int = p1
			_ = _rhs57
			var _rhs58 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter23 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _428_v116)).Update(p1, _428_v116)).Keys().Elements()); ; {
					_compr_22, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _429_v115 _dafny.Int
					_429_v115 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _428_v116)).Update(p1, _428_v116)).Contains(_429_v115) {
						_coll22.Add(Companion_Default___.SafeModuloInt(_429_v115, p1), (func() _dafny.Map {
							if _425_v112 {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_263_v1).Cardinality()), (_this).F14())
							}
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F14())
						})())
					}
				}
				return _coll22.ToMap()
			}()).Cardinality())
			_ = _rhs58
			var _rhs59 _dafny.Sequence = (func() _dafny.Sequence {
				if (_this).F14() {
					return _dafny.Companion_Sequence_.Concatenate(_356_v61, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}(func(_430_i9 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('t')
					})))
				}
				return _356_v61
			})()
			_ = _rhs59
			var _lhs38 *GlobalState = globalState
			_ = _lhs38
			var _lhs39 *GlobalState = globalState
			_ = _lhs39
			var _lhs40 *GlobalState = globalState
			_ = _lhs40
			_lhs38.F1 = _rhs56
			_lhs39.F1 = _rhs57
			_lhs40.F1 = _rhs58
			_356_v61 = _rhs59
			if (_this).F15() {
				_425_v112 = _425_v112
				_425_v112 = (_this).F15()
				var _431_v117 bool
				_ = _431_v117
				var _432_v118 _dafny.MultiSet
				_ = _432_v118
				var _433_v119 _dafny.Set
				_ = _433_v119
				var _434_v120 _dafny.Int
				_ = _434_v120
				var _out20 bool
				_ = _out20
				var _out21 _dafny.MultiSet
				_ = _out21
				var _out22 _dafny.Set
				_ = _out22
				var _out23 _dafny.Int
				_ = _out23
				_out20, _out21, _out22, _out23 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_263_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32(), (_this).F15()), _263_v1), _263_v1, globalState)
				_431_v117 = _out20
				_432_v118 = _out21
				_433_v119 = _out22
				_434_v120 = _out23
				_425_v112 = _431_v117
				var _435_v121 _dafny.Map
				_ = _435_v121
				_435_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(381))).Cardinality(), (_this).F14())
				var _436_v122 _dafny.MultiSet
				_ = _436_v122
				_436_v122 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_((_this).F15(), (_435_v121).Cardinality()))
				var _437_v123 D0
				_ = _437_v123
				_437_v123 = Companion_D0_.Create_DC1_((_this).F14(), p1)
				(globalState).F1 = Companion_Default___.SafeDivisionInt(_434_v120, (func() _dafny.Int {
					if (_436_v122).Contains(_437_v123) {
						return (_436_v122).Multiplicity(_437_v123)
					}
					return _434_v120
				})())
			} else {
				var _438_v124 *C0
				_ = _438_v124
				var _nw69 *C0 = New_C0_()
				_ = _nw69
				_nw69.Ctor__()
				_438_v124 = _nw69
				var _439_v125 _dafny.Array
				_ = _439_v125
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
				_ = _nw70
				_439_v125 = _nw70
				var _440_v126 _dafny.Array
				_ = _440_v126
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw71
				_440_v126 = _nw71
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_439_v125), 0))
				_ = _index69
				(_439_v125).ArraySet1(_440_v126, (_index69).Int())
				var _441_v128 _dafny.Map
				_ = _441_v128
				_441_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _442_v129 _dafny.Map
				_ = _442_v129
				_442_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll23 = _dafny.NewMapBuilder()
					_ = _coll23
					for _iter24 := _dafny.Iterate((_441_v128).Keys().Elements()); ; {
						_compr_23, _ok24 := _iter24()
						if !_ok24 {
							break
						}
						var _443_v127 _dafny.Int
						_443_v127 = interface{}(_compr_23).(_dafny.Int)
						if (_441_v128).Contains(_443_v127) {
							_coll23.Add(Companion_Default___.SafeModuloInt(_443_v127, p1), p1)
						}
					}
					return _coll23.ToMap()
				}()).Merge(_441_v128), _440_v126)
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_439_v125), 0))
				_ = _index70
				(_439_v125).ArraySet1((func() _dafny.Array {
					if (_442_v129).Contains(_441_v128) {
						return (_442_v129).Get(_441_v128).(_dafny.Array)
					}
					return _440_v126
				})(), (_index70).Int())
				var _444_v130 _dafny.Sequence
				_ = _444_v130
				_444_v130 = _dafny.SeqOf(_356_v61)
				var _445_v131 _dafny.Map
				_ = _445_v131
				_445_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_444_v130).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_444_v130).Cardinality()))).Uint32()).(_dafny.Sequence))
				_356_v61 = (func() _dafny.Sequence {
					if (_445_v131).Contains((_this).F15()) {
						return (_445_v131).Get((_this).F15()).(_dafny.Sequence)
					}
					return _356_v61
				})()
				_425_v112 = (_this).F15()
				_425_v112 = (_this).F14()
			}
		} else {
			var _446_v132 _dafny.Array
			_ = _446_v132
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_12
			var _nw72 _dafny.Array
			_ = _nw72
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw72 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Int = (func(_447_v61 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_448_i10 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_448_i10, _dafny.IntOfUint32((_447_v61).Cardinality()))
					}
				})(_356_v61)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw72 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw72).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw72).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_446_v132 = _nw72
			_446_v132 = _446_v132
			var _449_v133 bool
			_ = _449_v133
			_449_v133 = true
			var _450_v134 _dafny.Map
			_ = _450_v134
			_450_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F10()).Cardinality(), false)
			var _451_v135 _dafny.Map
			_ = _451_v135
			_451_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_450_v134, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality()))
			var _452_v136 _dafny.Array
			_ = _452_v136
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_13
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) bool = (func(_453_v1 _dafny.Sequence, _454_p1 _dafny.Int) func(_dafny.Int) bool {
					return func(_455_i11 _dafny.Int) bool {
						return (_453_v1).Select((Companion_Default___.SafeIndex(_454_p1, _dafny.IntOfUint32((_453_v1).Cardinality()))).Uint32()).(bool)
					}
				})(_263_v1, p1)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw73 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw73).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw73).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_452_v136 = _nw73
			var _456_v137 _dafny.MultiSet
			_ = _456_v137
			_456_v137 = _dafny.MultiSetOf(_452_v136)
			_449_v133 = (p1).Cmp(((_451_v135).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _449_v133), (func() _dafny.Int {
				if (_456_v137).Contains(_452_v136) {
					return (_456_v137).Multiplicity(_452_v136)
				}
				return (_dafny.Zero).Minus(p1)
			})())).Cardinality()) >= 0
			_449_v133 = !((_this).F14()) || ((_this).F15())
			_449_v133 = (_this).F14()
			if (_this).F15() {
				_449_v133 = (_this).F15()
				(globalState).F1 = ((_dafny.IntOfInt64(549)).Minus(p1)).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p1), p1))
				var _457_v138 _dafny.Map
				_ = _457_v138
				_457_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F15())
				var _458_v139 _dafny.Map
				_ = _458_v139
				_458_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_457_v138).Cardinality(), (_357_v62).Difference(_dafny.MultiSetOf(p1)))
				var _459_v140 _dafny.MultiSet
				_ = _459_v140
				_459_v140 = (_this).F10()
				var _460_v141 *C0
				_ = _460_v141
				var _nw74 *C0 = New_C0_()
				_ = _nw74
				_nw74.Ctor__()
				_460_v141 = _nw74
				var _rhs60 _dafny.Int = (p1).Times((_450_v134).Cardinality())
				_ = _rhs60
				var _rhs61 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _357_v62)
				_ = _rhs61
				var _rhs62 bool = !((_this).F10()).Equals((_this).F10())
				_ = _rhs62
				var _rhs63 *C0 = _460_v141
				_ = _rhs63
				var _rhs64 _dafny.Int = ((((_this).F10()).Cardinality()).Times(p1)).Times((_dafny.Zero).Minus(p1))
				_ = _rhs64
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 *GlobalState = globalState
				_ = _lhs42
				_lhs41.F1 = _rhs60
				_458_v139 = _rhs61
				_449_v133 = _rhs62
				r2 = _rhs63
				_lhs42.F1 = _rhs64
				var _461_v142 _dafny.Array
				_ = _461_v142
				var _nw75 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
				_ = _nw75
				_461_v142 = _nw75
				_461_v142 = _461_v142
				var _462_v143 D1
				_ = _462_v143
				_462_v143 = Companion_D1_.Create_DC4_((_this).F14(), false, p1, Companion_Default___.Fm20(p1, _dafny.IntOfInt64(713), globalState), _dafny.IntOfInt64(-757))
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_452_v136), 0))
				_ = _index71
				(_452_v136).ArraySet1((_462_v143).Dtor_cf6(), (_index71).Int())
				var _463_v144 _dafny.Map
				_ = _463_v144
				_463_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_449_v133, p1)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_452_v136), 0))
				_ = _index72
				(_452_v136).ArraySet1(!((func() bool {
					if (_450_v134).Contains((_463_v144).Cardinality()) {
						return (_450_v134).Get((_463_v144).Cardinality()).(bool)
					}
					return (_this).F15()
				})()) || ((func() bool {
					if !(!((_263_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool))) {
						return !(!(!(true)))
					}
					return (_this).F15()
				})()), (_index72).Int())
			} else {
				var _464_v145 _dafny.Array
				_ = _464_v145
				var _nwElement0_19 _dafny.MultiSet = _dafny.MultiSetOf(_452_v136, _452_v136, _452_v136)
				_ = _nwElement0_19
				var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(3))
				_ = _nw76
				(_nw76).ArraySet1(_nwElement0_19, 0)
				(_nw76).ArraySet1(_dafny.MultiSetOf(_452_v136, _452_v136), 1)
				(_nw76).ArraySet1(_456_v137, 2)
				_464_v145 = _nw76
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_464_v145), 0))
				_ = _index73
				(_464_v145).ArraySet1((_456_v137).Difference(_456_v137), (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_464_v145), 0))
				_ = _index74
				(_464_v145).ArraySet1(_456_v137, (_index74).Int())
				var _465_v146 D1
				_ = _465_v146
				_465_v146 = Companion_D1_.Create_DC4_(_449_v133, (_this).F14(), p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(81))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}(func(_466_i12 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-93)
				})), p1)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_446_v132), 0))
				_ = _index75
				(_446_v132).ArraySet1((_465_v146).Dtor_cf9(), (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_446_v132), 0))
				_ = _index76
				(_446_v132).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), (_index76).Int())
				_449_v133 = !((_this).F14())
				var _467_v147 *C0
				_ = _467_v147
				var _nw77 *C0 = New_C0_()
				_ = _nw77
				_nw77.Ctor__()
				_467_v147 = _nw77
				var _468_v148 *C0
				_ = _468_v148
				var _nw78 *C0 = New_C0_()
				_ = _nw78
				_nw78.Ctor__()
				_468_v148 = _nw78
			}
		}
		if (_this).F15() {
			if (_263_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool) {
				var _469_v149 bool
				_ = _469_v149
				_469_v149 = false
				_469_v149 = ((p1).Times(p1)).Cmp(p1) >= 0
				var _470_v150 *C0
				_ = _470_v150
				var _nw79 *C0 = New_C0_()
				_ = _nw79
				_nw79.Ctor__()
				_470_v150 = _nw79
				(globalState).F1 = p1
				var _471_v151 *C0
				_ = _471_v151
				var _nw80 *C0 = New_C0_()
				_ = _nw80
				_nw80.Ctor__()
				_471_v151 = _nw80
				_356_v61 = _dafny.Companion_Sequence_.Concatenate(_356_v61, _356_v61)
			} else {
				var _472_v152 _dafny.Array
				_ = _472_v152
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_14
				var _nw81 _dafny.Array
				_ = _nw81
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw81 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) _dafny.Int = func(_473_i13 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_473_i13, _dafny.IntOfInt64(86))
					}
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw81 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw81).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw81).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_472_v152 = _nw81
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_472_v152), 0))
				_ = _index77
				(_472_v152).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm19(_355_v60, (_this).F15(), globalState), (_dafny.Zero).Minus(p1)), (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_472_v152), 0))
				_ = _index78
				(_472_v152).ArraySet1(Companion_Default___.Fm19(_355_v60, (_this).F15(), globalState), (_index78).Int())
				var _474_v153 _dafny.Array
				_ = _474_v153
				var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
				_ = _nw82
				_474_v153 = _nw82
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_474_v153), 0))
				_ = _index79
				(_474_v153).ArraySet1(_dafny.SeqOf(!((_this).F14())), (_index79).Int())
				var _475_v154 bool
				_ = _475_v154
				_475_v154 = true
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_474_v153), 0))
				_ = _index80
				var _rhs65 _dafny.Sequence = _263_v1
				_ = _rhs65
				var _rhs66 bool = Companion_Default___.Fm4(_354_v59, globalState)
				_ = _rhs66
				var _rhs67 bool = !((Companion_Default___.Fm19(_355_v60, (_this).F14(), globalState)).Cmp(p1) > 0)
				_ = _rhs67
				var _lhs43 _dafny.Array = _474_v153
				_ = _lhs43
				var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_474_v153), 0))
				_ = _lhs44
				(_lhs43).ArraySet1(_rhs65, (_lhs44).Int())
				_475_v154 = _rhs66
				_475_v154 = _rhs67
				var _476_v155 *C0
				_ = _476_v155
				var _nw83 *C0 = New_C0_()
				_ = _nw83
				_nw83.Ctor__()
				_476_v155 = _nw83
				var _477_v156 *C0
				_ = _477_v156
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__()
				_477_v156 = _nw84
				_356_v61 = Companion_Default___.Fm2(_475_v154, (_472_v152).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_472_v152), 0))).Int()).(_dafny.Int), _355_v60, (_472_v152).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_472_v152), 0))).Int()).(_dafny.Int), globalState)
			}
			var _478_v157 _dafny.Array
			_ = _478_v157
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw85
			_478_v157 = _nw85
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_478_v157), 0))
			_ = _index81
			(_478_v157).ArraySet1((_this).F15(), (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_478_v157), 0))
			_ = _index82
			(_478_v157).ArraySet1(false, (_index82).Int())
			var _479_v158 _dafny.Array
			_ = _479_v158
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_15
			var _nw86 _dafny.Array
			_ = _nw86
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw86 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.CodePoint = (func(_480_v60 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_481_i14 _dafny.Int) _dafny.CodePoint {
						return _480_v60
					}
				})(_355_v60)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw86 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw86).ArraySet1CodePoint(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw86).ArraySet1CodePoint(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_479_v158 = _nw86
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_479_v158), 0))
			_ = _index83
			(_479_v158).ArraySet1CodePoint(_355_v60, (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_479_v158), 0))
			_ = _index84
			(_479_v158).ArraySet1CodePoint(Companion_Default___.Fm1(globalState), (_index84).Int())
			var _482_v159 _dafny.Array
			_ = _482_v159
			var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw87
			_482_v159 = _nw87
			var _483_v160 _dafny.Map
			_ = _483_v160
			_483_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_482_v159, p1)
			_483_v160 = ((_483_v160).Merge(_483_v160)).Merge(_483_v160)
			var _484_v161 _dafny.Map
			_ = _484_v161
			_484_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F15())
			var _485_v163 D8
			_ = _485_v163
			_485_v163 = Companion_D8_.Create_DC19_(_484_v161)
			var _486_v164 _dafny.Set
			_ = _486_v164
			_486_v164 = _dafny.SetOf(_484_v161, _484_v161, (func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(976), _dafny.IntOfInt64(896))); ; {
					_compr_24, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _487_v162 _dafny.Int
					_487_v162 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(976)).Cmp(_487_v162) <= 0) && ((_487_v162).Cmp(_dafny.IntOfInt64(896)) < 0) {
						_coll24.Add((_487_v162).Minus(p1), true)
					}
				}
				return _coll24.ToMap()
			}()).Merge((_485_v163).Dtor_cf42()), _484_v161)
			var _488_v165 _dafny.Sequence
			_ = _488_v165
			_488_v165 = _dafny.SeqOf(_486_v164)
			_486_v164 = (_486_v164).Union((_488_v165).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_488_v165).Cardinality()))).Uint32()).(_dafny.Set))
		} else {
			if Companion_Default___.Fm4((func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(52), _dafny.IntOfInt64(-560))); ; {
					_compr_25, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _489_v166 _dafny.Int
					_489_v166 = interface{}(_compr_25).(_dafny.Int)
					if ((_dafny.IntOfInt64(52)).Cmp(_489_v166) <= 0) && ((_489_v166).Cmp(_dafny.IntOfInt64(-560)) < 0) {
						_coll25.Add(Companion_Default___.SafeModuloInt(_489_v166, _dafny.IntOfUint32((_dafny.SeqOf(p1, p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(385))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg27 _dafny.Int) interface{} {
								return coer27(arg27)
							}
						}((func(_490_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
							return func(_491_i15 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _dafny.MultiSetFromSeq(_490_v1))
							}
						})(_263_v1)))).Cardinality()))).Cardinality())).Cardinality())))
					}
				}
				return _coll25.ToSet()
			}()).Difference(_354_v59), globalState) {
				(globalState).F1 = p1
				var _492_v167 _dafny.Array
				_ = _492_v167
				var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
				_ = _nw88
				_492_v167 = _nw88
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_492_v167), 0))
				_ = _index85
				(_492_v167).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_356_v61, _356_v61), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_356_v61, _356_v61)).Cardinality()))).Uint32(), _355_v60), (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_492_v167), 0))
				_ = _index86
				(_492_v167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(285))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}((func(_493_v60 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_494_i16 _dafny.Int) _dafny.CodePoint {
						return _493_v60
					}
				})(_355_v60))), _dafny.Companion_Sequence_.Update(_356_v61, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_356_v61).Cardinality()))).Uint32(), _355_v60)), (_index86).Int())
				var _495_v168 bool
				_ = _495_v168
				_495_v168 = true
				_495_v168 = (_this).F14()
				(globalState).F1 = (_dafny.Zero).Minus((_dafny.Zero).Minus(p1))
				var _496_v169 _dafny.Map
				_ = _496_v169
				_496_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_263_v1, _263_v1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_356_v61).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_263_v1, _263_v1)).Cardinality()))).Uint32(), (_this).F14()), (p1).Times(p1))
				var _497_v170 D3
				_ = _497_v170
				_497_v170 = Companion_D3_.Create_DC7_((_492_v167).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_492_v167), 0))).Int()).(_dafny.Sequence))
				var _498_v171 _dafny.Map
				_ = _498_v171
				_498_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_497_v170, false)
				var _499_v172 _dafny.Map
				_ = _499_v172
				_499_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), !((_this).F14()))
				var _500_v173 _dafny.Map
				_ = _500_v173
				_500_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC1_((_this).F14(), p1)).Dtor_cf1(), _499_v172)
				var _rhs68 _dafny.Int = (func() _dafny.Int {
					if (_496_v169).Contains(_263_v1) {
						return (_496_v169).Get(_263_v1).(_dafny.Int)
					}
					return p1
				})()
				_ = _rhs68
				var _rhs69 bool = (((func() _dafny.Map {
					if (_500_v173).Contains(_495_v168) {
						return (_500_v173).Get(_495_v168).(_dafny.Map)
					}
					return _499_v172
				})()).Merge(_499_v172)).Contains((func() bool {
					if (_498_v171).Contains(_497_v170) {
						return (_498_v171).Get(_497_v170).(bool)
					}
					return (func() bool {
						if (_499_v172).Contains((_this).F15()) {
							return (_499_v172).Get((_this).F15()).(bool)
						}
						return (_this).F14()
					})()
				})())
				_ = _rhs69
				var _rhs70 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, Companion_Default___.Fm19(_355_v60, (_this).F15(), globalState))
				_ = _rhs70
				var _lhs45 *GlobalState = globalState
				_ = _lhs45
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				_lhs45.F1 = _rhs68
				_495_v168 = _rhs69
				_lhs46.F1 = _rhs70
			} else {
				var _501_v174 _dafny.MultiSet
				_ = _501_v174
				_501_v174 = (_this).F10()
				_501_v174 = _dafny.MultiSetOf((_263_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool))
				var _502_v175 _dafny.Array
				_ = _502_v175
				var _nwElement0_20 D0 = _384_v86
				_ = _nwElement0_20
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(3))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_20, 0)
				(_nw89).ArraySet1(Companion_Default___.Fm21(globalState), 1)
				(_nw89).ArraySet1(_384_v86, 2)
				_502_v175 = _nw89
				var _503_v176 _dafny.Map
				_ = _503_v176
				_503_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_502_v175, (_263_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ale")).Cardinality()), _dafny.IntOfUint32((_263_v1).Cardinality()))).Uint32()).(bool))
				var _504_v177 _dafny.Sequence
				_ = _504_v177
				_504_v177 = _dafny.SeqOf(_354_v59, _354_v59, _354_v59)
				var _505_v178 _dafny.Map
				_ = _505_v178
				_505_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_504_v177).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_504_v177).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _263_v1)
				var _506_v179 _dafny.Map
				_ = _506_v179
				_506_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_505_v178, _502_v175)
				_503_v176 = (_503_v176).Update((func() _dafny.Array {
					if (_506_v179).Contains(_505_v178) {
						return (_506_v179).Get(_505_v178).(_dafny.Array)
					}
					return _502_v175
				})(), _dafny.Companion_Sequence_.Equal(_356_v61, _356_v61))
				var _507_v180 bool
				_ = _507_v180
				_507_v180 = false
				var _508_v181 _dafny.Array
				_ = _508_v181
				var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw90
				_508_v181 = _nw90
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _index87
				(_508_v181).ArraySet1(p1, (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _index88
				var _rhs71 bool = false
				_ = _rhs71
				var _rhs72 _dafny.Int = p1
				_ = _rhs72
				var _lhs47 _dafny.Array = _508_v181
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _lhs48
				_507_v180 = _rhs71
				(_lhs47).ArraySet1(_rhs72, (_lhs48).Int())
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _index89
				(_508_v181).ArraySet1(Companion_Default___.SafeDivisionInt(((_508_v181).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))).Int()).(_dafny.Int)).Times((_508_v181).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))).Int()).(_dafny.Int)), Companion_Default___.Fm19(Companion_Default___.Fm1(globalState), (_this).F15(), globalState)), (_index89).Int())
				var _509_v182 D1
				_ = _509_v182
				_509_v182 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(-460), (_this).F15(), (_508_v181).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))).Int()).(_dafny.Int), (_this).F15(), p1)
				var _510_v184 _dafny.Sequence
				_ = _510_v184
				_510_v184 = _dafny.SeqOf((_508_v181).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))).Int()).(_dafny.Int))
				var _511_v185 _dafny.Sequence
				_ = _511_v185
				_511_v185 = _dafny.SeqOf(_509_v182, _509_v182, _509_v182, Companion_D1_.Create_DC5_((_358_v63).Cardinality(), false, p1, true, _dafny.IntOfUint32((_510_v184).Cardinality())))
				var _512_v188 _dafny.Set
				_ = _512_v188
				_512_v188 = _dafny.SetOf(_509_v182)
				var _513_v189 _dafny.Map
				_ = _513_v189
				_513_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_507_v180, Companion_Default___.SafeModuloInt(p1, p1))
				var _514_v190 _dafny.Array
				_ = _514_v190
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw91
				_514_v190 = _nw91
				var _515_v191 D1
				_ = _515_v191
				_515_v191 = Companion_D1_.Create_DC3_(_514_v190)
				var _516_v192 _dafny.Set
				_ = _516_v192
				_516_v192 = _dafny.SetOf(_515_v191, _515_v191)
				var _517_v193 _dafny.Map
				_ = _517_v193
				_517_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_516_v192, p1)
				var _518_v194 _dafny.Sequence
				_ = _518_v194
				_518_v194 = _dafny.SeqOf(_517_v193, _517_v193, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_516_v192, _dafny.IntOfUint32((_263_v1).Cardinality()))).Update(_dafny.SetOf(_515_v191), Companion_Default___.Fm19(_355_v60, (_this).F15(), globalState)))
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _index90
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _index91
				var _rhs73 bool = ((_dafny.SetOf(_509_v182, Companion_D1_.Create_DC5_(Companion_Default___.Fm19(_355_v60, (_this).F15(), globalState), false, (_508_v181).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))).Int()).(_dafny.Int), (_this).F14(), (_354_v59).Cardinality()))).Difference(func() _dafny.Set {
					var _coll26 = _dafny.NewBuilder()
					_ = _coll26
					for _iter27 := _dafny.Iterate((func() _dafny.Map {
						var _coll27 = _dafny.NewMapBuilder()
						_ = _coll27
						for _iter28 := _dafny.Iterate((func() _dafny.Set {
							var _coll28 = _dafny.NewBuilder()
							_ = _coll28
							for _iter29 := _dafny.Iterate((_511_v185).Elements()); ; {
								_compr_28, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _519_v186 D1
								_519_v186 = interface{}(_compr_28).(D1)
								if _dafny.Companion_Sequence_.Contains(_511_v185, _519_v186) {
									_coll28.Add(_519_v186)
								}
							}
							return _coll28.ToSet()
						}()).Elements()); ; {
							_compr_27, _ok28 := _iter28()
							if !_ok28 {
								break
							}
							var _520_v183 D1
							_520_v183 = interface{}(_compr_27).(D1)
							if (func() _dafny.Set {
								var _coll29 = _dafny.NewBuilder()
								_ = _coll29
								for _iter30 := _dafny.Iterate((_511_v185).Elements()); ; {
									_compr_29, _ok30 := _iter30()
									if !_ok30 {
										break
									}
									var _521_v186 D1
									_521_v186 = interface{}(_compr_29).(D1)
									if _dafny.Companion_Sequence_.Contains(_511_v185, _521_v186) {
										_coll29.Add(_521_v186)
									}
								}
								return _coll29.ToSet()
							}()).Contains(_520_v183) {
								_coll27.Add(_520_v183, (_this).F15())
							}
						}
						return _coll27.ToMap()
					}()).Keys().Elements()); ; {
						_compr_26, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _522_v187 D1
						_522_v187 = interface{}(_compr_26).(D1)
						if (func() _dafny.Map {
							var _coll30 = _dafny.NewMapBuilder()
							_ = _coll30
							for _iter31 := _dafny.Iterate((func() _dafny.Set {
								var _coll31 = _dafny.NewBuilder()
								_ = _coll31
								for _iter32 := _dafny.Iterate((_511_v185).Elements()); ; {
									_compr_31, _ok32 := _iter32()
									if !_ok32 {
										break
									}
									var _523_v186 D1
									_523_v186 = interface{}(_compr_31).(D1)
									if _dafny.Companion_Sequence_.Contains(_511_v185, _523_v186) {
										_coll31.Add(_523_v186)
									}
								}
								return _coll31.ToSet()
							}()).Elements()); ; {
								_compr_30, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _520_v183 D1
								_520_v183 = interface{}(_compr_30).(D1)
								if (func() _dafny.Set {
									var _coll32 = _dafny.NewBuilder()
									_ = _coll32
									for _iter33 := _dafny.Iterate((_511_v185).Elements()); ; {
										_compr_32, _ok33 := _iter33()
										if !_ok33 {
											break
										}
										var _524_v186 D1
										_524_v186 = interface{}(_compr_32).(D1)
										if _dafny.Companion_Sequence_.Contains(_511_v185, _524_v186) {
											_coll32.Add(_524_v186)
										}
									}
									return _coll32.ToSet()
								}()).Contains(_520_v183) {
									_coll30.Add(_520_v183, (_this).F15())
								}
							}
							return _coll30.ToMap()
						}()).Contains(_522_v187) {
							_coll26.Add(_522_v187)
						}
					}
					return _coll26.ToSet()
				}())).IsDisjointFrom(_512_v188)
				_ = _rhs73
				var _rhs74 _dafny.Int = (func() _dafny.Int {
					if (_513_v189).Contains((_this).F15()) {
						return (_513_v189).Get((_this).F15()).(_dafny.Int)
					}
					return p1
				})()
				_ = _rhs74
				var _rhs75 bool = !(true)
				_ = _rhs75
				var _rhs76 _dafny.Int = ((_518_v194).Select((Companion_Default___.SafeIndex((_508_v181).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_518_v194).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()
				_ = _rhs76
				var _lhs49 _dafny.Array = _508_v181
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _lhs50
				var _lhs51 _dafny.Array = _508_v181
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_508_v181), 0))
				_ = _lhs52
				_507_v180 = _rhs73
				(_lhs49).ArraySet1(_rhs74, (_lhs50).Int())
				_507_v180 = _rhs75
				(_lhs51).ArraySet1(_rhs76, (_lhs52).Int())
			}
			var _525_v195 _dafny.Set
			_ = _525_v195
			_525_v195 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("nthxycdvu"))
			var _526_v196 _dafny.Map
			_ = _526_v196
			_526_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _525_v195)
			var _527_v197 _dafny.Array
			_ = _527_v197
			var _nwElement0_21 _dafny.Set = _525_v195
			_ = _nwElement0_21
			var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(9))
			_ = _nw92
			(_nw92).ArraySet1(_nwElement0_21, 0)
			(_nw92).ArraySet1(_525_v195, 1)
			(_nw92).ArraySet1(_525_v195, 2)
			(_nw92).ArraySet1(_525_v195, 3)
			(_nw92).ArraySet1(_525_v195, 4)
			(_nw92).ArraySet1((func() _dafny.Set {
				if (_526_v196).Contains(p1) {
					return (_526_v196).Get(p1).(_dafny.Set)
				}
				return _525_v195
			})(), 5)
			(_nw92).ArraySet1(_dafny.SetOf(_356_v61, _356_v61), 6)
			(_nw92).ArraySet1((func() _dafny.Set {
				if (_526_v196).Contains(p1) {
					return (_526_v196).Get(p1).(_dafny.Set)
				}
				return _525_v195
			})(), 7)
			(_nw92).ArraySet1((_525_v195).Difference(_525_v195), 8)
			_527_v197 = _nw92
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_527_v197), 0))
			_ = _index92
			(_527_v197).ArraySet1((_dafny.SetOf(_356_v61)).Difference(_525_v195), (_index92).Int())
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_527_v197), 0))
			_ = _index93
			(_527_v197).ArraySet1(Companion_Default___.Fm22(globalState), (_index93).Int())
			var _528_v198 bool
			_ = _528_v198
			_528_v198 = false
			_528_v198 = !(((p1).Cmp(p1) == 0) || (_dafny.Companion_Sequence_.Contains(_263_v1, _528_v198)))
			var _529_v199 T2
			_ = _529_v199
			var _nw93 *C0 = New_C0_()
			_ = _nw93
			_nw93.Ctor__()
			_529_v199 = _nw93
			_529_v199 = _529_v199
			var _530_v200 _dafny.Array
			_ = _530_v200
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_16
			var _nw94 _dafny.Array
			_ = _nw94
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw94 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) bool = (func(_531_v1 _dafny.Sequence, _532_v61 _dafny.Sequence) func(_dafny.Int) bool {
					return func(_533_i17 _dafny.Int) bool {
						return (_531_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_532_v61).Cardinality()), _dafny.IntOfUint32((_531_v1).Cardinality()))).Uint32()).(bool)
					}
				})(_263_v1, _356_v61)
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
			_530_v200 = _nw94
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_530_v200), 0))
			_ = _index94
			(_530_v200).ArraySet1(_528_v198, (_index94).Int())
			var _534_v201 _dafny.Map
			_ = _534_v201
			_534_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm19(_355_v60, true, globalState), _dafny.IntOfInt64(824)), !((_this).F14()))
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_530_v200), 0))
			_ = _index95
			(_530_v200).ArraySet1(!((func() bool {
				if (_534_v201).Contains((_dafny.Zero).Minus(Companion_Default___.Fm19(_355_v60, _528_v198, globalState))) {
					return (_534_v201).Get((_dafny.Zero).Minus(Companion_Default___.Fm19(_355_v60, _528_v198, globalState))).(bool)
				}
				return _528_v198
			})()), (_index95).Int())
		}
		var _535_v202 _dafny.Array
		_ = _535_v202
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_17
		var _nw95 _dafny.Array
		_ = _nw95
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw95 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) bool = (func(_536_p1 _dafny.Int) func(_dafny.Int) bool {
				return func(_537_i18 _dafny.Int) bool {
					return ((_dafny.Zero).Minus(_536_p1)).Cmp(_536_p1) > 0
				}
			})(p1)
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
		_535_v202 = _nw95
		r0 = _535_v202
		var _538_v204 _dafny.Map
		_ = _538_v204
		_538_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F14())
		r1 = (func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(171), _dafny.IntOfInt64(252))); ; {
				_compr_33, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _539_v203 _dafny.Int
				_539_v203 = interface{}(_compr_33).(_dafny.Int)
				if ((_dafny.IntOfInt64(171)).Cmp(_539_v203) <= 0) && ((_539_v203).Cmp(_dafny.IntOfInt64(252)) < 0) {
					_coll33.Add((_539_v203).Plus(_dafny.IntOfInt64(429)), (_this).F14())
				}
			}
			return _coll33.ToMap()
		}()).Merge(_538_v204)
		var _540_v205 *C0
		_ = _540_v205
		var _nw96 *C0 = New_C0_()
		_ = _nw96
		_nw96.Ctor__()
		_540_v205 = _nw96
		r2 = _540_v205
		return r0, r1, r2
	}
}
func (_this *C1) F14() bool {
	{
		return _this._f14
	}
}
func (_this *C1) F15() bool {
	{
		return _this._f15
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f10 _dafny.MultiSet
	F12  bool
	_f13 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f10 = _dafny.EmptyMultiSet
	_this.F12 = false
	_this._f13 = _dafny.Zero
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

func (_this *C2) F10() _dafny.MultiSet {
	return _this._f10
}
func (_this *C2) Ctor__(f12 bool, f13 _dafny.Int, f10 _dafny.MultiSet) {
	{
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this)._f10 = f10
	}
}
func (_this *C2) Fm10(p0 D0, p1 D0, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('n')
	}
}
func (_this *C2) Fm11(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		if !(_this.F12) {
			return (_this).F13()
		} else {
			return (_this).F13()
		}
	}
}
func (_this *C2) M4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _541_v0 _dafny.Sequence
		_ = _541_v0
		_541_v0 = _dafny.UnicodeSeqOfUtf8Bytes("fsvmc")
		(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-614))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_542_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})), _541_v0)).Cardinality())
		var _543_v1 _dafny.Array
		_ = _543_v1
		var _nw97 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw97
		_543_v1 = _nw97
		var _544_v2 _dafny.Map
		_ = _544_v2
		_544_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_543_v1, Companion_Default___.Fm12(globalState))
		var _545_v3 _dafny.Sequence
		_ = _545_v3
		_545_v3 = _dafny.SeqOf(!(_this.F12))
		_544_v2 = (_544_v2).Update(_543_v1, _545_v3)
		for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_543_v1), 0))); ; {
			_guard_loop_1, _ok35 := _iter35()
			if !_ok35 {
				break
			}
			var _546_i1 _dafny.Int
			_546_i1 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_546_i1).Sign() != -1) && ((_546_i1).Cmp(_dafny.ArrayLen((_543_v1), 0)) < 0)) {
				(_543_v1).ArraySet1(_this.F12, (_546_i1).Int())
			}
		}
		var _547_i2 _dafny.Int
		_ = _547_i2
		_547_i2 = _dafny.Zero
		{
			for _this.F12 {
				{
					if (_547_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_547_i2 = (_547_i2).Plus(_dafny.One)
					var _548_v4 _dafny.Set
					_ = _548_v4
					_548_v4 = _dafny.SetOf(_dafny.IntOfUint32((_545_v3).Cardinality()), p1)
					var _549_v5 _dafny.Set
					_ = _549_v5
					_549_v5 = _dafny.SetOf((_548_v4).Cardinality(), (_this).F13())
					if (_dafny.SetOf(p0, p1)).IsDisjointFrom(_549_v5) {
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))
						_ = _index96
						(_543_v1).ArraySet1(_this.F12, (_index96).Int())
						var _550_v6 _dafny.Set
						_ = _550_v6
						_550_v6 = _dafny.SetOf(Companion_Default___.Fm4(_549_v5, globalState))
						var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))
						_ = _index97
						(_543_v1).ArraySet1((_dafny.SetOf(Companion_Default___.Fm4(_dafny.SetOf(p1, p1, p1), globalState), _this.F12, _this.F12, _this.F12)).IsDisjointFrom(_550_v6), (_index97).Int())
						var _551_v7 D0
						_ = _551_v7
						_551_v7 = Companion_D0_.Create_DC1_((_543_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))).Int()).(bool), p0)
						var _552_v8 D0
						_ = _552_v8
						_552_v8 = Companion_D0_.Create_DC2_(_551_v7)
						var _553_v9 _dafny.MultiSet
						_ = _553_v9
						_553_v9 = _dafny.MultiSetOf(_552_v8)
						var _554_v10 _dafny.Map
						_ = _554_v10
						_554_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v9, p2)
						_554_v10 = (_554_v10).Update(_553_v9, p2)
						var _555_v11 _dafny.MultiSet
						_ = _555_v11
						_555_v11 = _dafny.MultiSetOf(p1, p1)
						var _556_v12 _dafny.Map
						_ = _556_v12
						_556_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p1, p1, (_this).Fm11(_dafny.MultiSetFromSeq(p2), globalState), (_this).F13())).Intersection(_549_v5), ((_this).F10()).Update(_this.F12, Companion_Default___.Abs((_555_v11).Cardinality())))
						_556_v12 = (_556_v12).Update((_dafny.SetOf(p1, p0)).Difference(_548_v4), ((_this).F10()).Update((_543_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))).Int()).(bool), Companion_Default___.Abs(p0)))
						var _557_v13 _dafny.CodePoint
						_ = _557_v13
						_557_v13 = _dafny.CodePoint('t')
						var _558_v14 _dafny.Sequence
						_ = _558_v14
						_558_v14 = _dafny.SeqOf((_555_v11).Update((_this).F13(), Companion_Default___.Abs((_this).F13())), _dafny.MultiSetFromSeq(p2), _555_v11)
						var _559_v15 _dafny.Array
						_ = _559_v15
						var _nwElement0_22 _dafny.Sequence = _541_v0
						_ = _nwElement0_22
						var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(17))
						_ = _nw98
						(_nw98).ArraySet1(_nwElement0_22, 0)
						(_nw98).ArraySet1(_541_v0, 1)
						(_nw98).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(450))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg30 _dafny.Int) interface{} {
								return coer30(arg30)
							}
						}(func(_560_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('q')
						})), 2)
						(_nw98).ArraySet1(Companion_Default___.Fm2((_543_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))).Int()).(bool), p0, _557_v13, ((_558_v14).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_558_v14).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), globalState), 3)
						(_nw98).ArraySet1(_541_v0, 4)
						(_nw98).ArraySet1(_dafny.Companion_Sequence_.Update(_541_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_541_v0).Cardinality()), _dafny.IntOfUint32((_541_v0).Cardinality()))).Uint32(), _557_v13), 5)
						(_nw98).ArraySet1(_541_v0, 6)
						(_nw98).ArraySet1(_541_v0, 7)
						(_nw98).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_541_v0, _dafny.UnicodeSeqOfUtf8Bytes("qmb")), 8)
						(_nw98).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jqnofwwqn"), _541_v0), 9)
						(_nw98).ArraySet1(_541_v0, 10)
						(_nw98).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("majd"), 11)
						(_nw98).ArraySet1(_541_v0, 12)
						(_nw98).ArraySet1(_541_v0, 13)
						(_nw98).ArraySet1(_541_v0, 14)
						(_nw98).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_541_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-611))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}((func(_561_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_562_i4 _dafny.Int) _dafny.CodePoint {
								return _561_v13
							}
						})(_557_v13)))), 15)
						(_nw98).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-701))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg32 _dafny.Int) interface{} {
								return coer32(arg32)
							}
						}((func(_563_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_564_i5 _dafny.Int) _dafny.CodePoint {
								return _563_v13
							}
						})(_557_v13))), 16)
						_559_v15 = _nw98
						var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_559_v15), 0))
						_ = _index98
						(_559_v15).ArraySet1(_541_v0, (_index98).Int())
						var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_559_v15), 0))
						_ = _index99
						(_559_v15).ArraySet1(_541_v0, (_index99).Int())
						var _565_v16 _dafny.Array
						_ = _565_v16
						var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
						_ = _nw99
						_565_v16 = _nw99
						var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_565_v16), 0))
						_ = _index100
						(_565_v16).ArraySet1((func() _dafny.Int {
							if (_543_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))).Int()).(bool) {
								return (p2).Select((Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int)
							}
							return p1
						})(), (_index100).Int())
						var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))
						_ = _index101
						var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_565_v16), 0))
						_ = _index102
						var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_559_v15), 0))
						_ = _index103
						var _rhs77 _dafny.Int = (_this).F13()
						_ = _rhs77
						var _rhs78 bool = _this.F12
						_ = _rhs78
						var _rhs79 _dafny.Int = (p0).Times(p0)
						_ = _rhs79
						var _rhs80 _dafny.Sequence = (_559_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_559_v15), 0))).Int()).(_dafny.Sequence)
						_ = _rhs80
						var _lhs53 *GlobalState = globalState
						_ = _lhs53
						var _lhs54 _dafny.Array = _543_v1
						_ = _lhs54
						var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_543_v1), 0))
						_ = _lhs55
						var _lhs56 _dafny.Array = _565_v16
						_ = _lhs56
						var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_565_v16), 0))
						_ = _lhs57
						var _lhs58 _dafny.Array = _559_v15
						_ = _lhs58
						var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_559_v15), 0))
						_ = _lhs59
						_lhs53.F1 = _rhs77
						(_lhs54).ArraySet1(_rhs78, (_lhs55).Int())
						(_lhs56).ArraySet1(_rhs79, (_lhs57).Int())
						(_lhs58).ArraySet1(_rhs80, (_lhs59).Int())
					} else {
						var _566_v17 _dafny.Map
						_ = _566_v17
						_566_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(386), p0)
						_566_v17 = _566_v17
						_548_v4 = _548_v4
						(_this).F12 = _this.F12
						var _567_v18 D1
						_ = _567_v18
						_567_v18 = Companion_D1_.Create_DC3_(_543_v1)
						var _568_v19 _dafny.Array
						_ = _568_v19
						var _nwElement0_23 _dafny.Array = _543_v1
						_ = _nwElement0_23
						var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(19))
						_ = _nw100
						(_nw100).ArraySet1(_nwElement0_23, 0)
						(_nw100).ArraySet1(_543_v1, 1)
						(_nw100).ArraySet1(_543_v1, 2)
						(_nw100).ArraySet1(_543_v1, 3)
						(_nw100).ArraySet1(_543_v1, 4)
						(_nw100).ArraySet1(_543_v1, 5)
						(_nw100).ArraySet1((_567_v18).Dtor_cf4(), 6)
						(_nw100).ArraySet1(_543_v1, 7)
						(_nw100).ArraySet1(_543_v1, 8)
						(_nw100).ArraySet1(_543_v1, 9)
						(_nw100).ArraySet1(_543_v1, 10)
						(_nw100).ArraySet1(_543_v1, 11)
						(_nw100).ArraySet1(_543_v1, 12)
						(_nw100).ArraySet1(_543_v1, 13)
						(_nw100).ArraySet1(_543_v1, 14)
						(_nw100).ArraySet1(_543_v1, 15)
						(_nw100).ArraySet1(_543_v1, 16)
						(_nw100).ArraySet1(_543_v1, 17)
						(_nw100).ArraySet1(_543_v1, 18)
						_568_v19 = _nw100
						var _569_v20 _dafny.Map
						_ = _569_v20
						_569_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _568_v19)
						_569_v20 = (_569_v20).Update(_this.F12, _568_v19)
						var _570_v21 _dafny.Map
						_ = _570_v21
						_570_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_567_v18, _this.F12)
						var _571_v22 _dafny.MultiSet
						_ = _571_v22
						_571_v22 = _dafny.MultiSetOf((_this).F13(), (_this).F13())
						var _572_v23 _dafny.CodePoint
						_ = _572_v23
						_572_v23 = _dafny.CodePoint('c')
						var _573_v24 D1
						_ = _573_v24
						_573_v24 = Companion_D1_.Create_DC4_(_this.F12, _this.F12, (_570_v21).Cardinality(), _dafny.SeqOf(_dafny.IntOfUint32((_541_v0).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm2(_this.F12, (_571_v22).Cardinality(), _572_v23, p0, globalState)).Cardinality())), (_this).Fm11(_dafny.MultiSetOf(_dafny.IntOfUint32((_541_v0).Cardinality())), globalState))
						var _574_v25 _dafny.Sequence
						_ = _574_v25
						_574_v25 = _dafny.SeqOf(_573_v24, _573_v24, _573_v24, _573_v24, _573_v24)
						var _575_v26 _dafny.Map
						_ = _575_v26
						_575_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, _574_v25)
						var _576_v27 D0
						_ = _576_v27
						_576_v27 = Companion_D0_.Create_DC1_(!(((_575_v26).Cardinality()).Cmp((_this).F13()) != 0), _dafny.IntOfInt64(156))
						var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_543_v1), 0))
						_ = _index104
						(_543_v1).ArraySet1(_this.F12, (_index104).Int())
						var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_568_v19), 0))
						_ = _index105
						(_568_v19).ArraySet1(_543_v1, (_index105).Int())
						var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_543_v1), 0))
						_ = _index106
						var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_568_v19), 0))
						_ = _index107
						var _rhs81 D0 = _576_v27
						_ = _rhs81
						var _rhs82 bool = (func() bool {
							if _this.F12 {
								return !((_571_v22).IsSubsetOf(_571_v22))
							}
							return _this.F12
						})()
						_ = _rhs82
						var _rhs83 _dafny.Int = Companion_Default___.SafeDivisionInt(((_this).F13()).Times((_this).Fm11(_571_v22, globalState)), ((func() _dafny.Map {
							if _this.F12 {
								return _566_v17
							}
							return _566_v17
						})()).Cardinality())
						_ = _rhs83
						var _rhs84 _dafny.Array = _543_v1
						_ = _rhs84
						var _lhs60 _dafny.Array = _543_v1
						_ = _lhs60
						var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_543_v1), 0))
						_ = _lhs61
						var _lhs62 *GlobalState = globalState
						_ = _lhs62
						var _lhs63 _dafny.Array = _568_v19
						_ = _lhs63
						var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_568_v19), 0))
						_ = _lhs64
						_576_v27 = _rhs81
						(_lhs60).ArraySet1(_rhs82, (_lhs61).Int())
						_lhs62.F1 = _rhs83
						(_lhs63).ArraySet1(_rhs84, (_lhs64).Int())
					}
					var _577_v28 D1
					_ = _577_v28
					_577_v28 = Companion_D1_.Create_DC3_(_543_v1)
					var _578_v29 _dafny.Array
					_ = _578_v29
					var _nwElement0_24 D1 = _577_v28
					_ = _nwElement0_24
					var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(3))
					_ = _nw101
					(_nw101).ArraySet1(_nwElement0_24, 0)
					(_nw101).ArraySet1(_577_v28, 1)
					(_nw101).ArraySet1(_577_v28, 2)
					_578_v29 = _nw101
					var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_578_v29), 0))
					_ = _index108
					(_578_v29).ArraySet1(_577_v28, (_index108).Int())
					var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_578_v29), 0))
					_ = _index109
					(_578_v29).ArraySet1(_577_v28, (_index109).Int())
					if _this.F12 {
						var _579_v30 _dafny.CodePoint
						_ = _579_v30
						_579_v30 = _dafny.CodePoint('n')
						var _580_v31 _dafny.Set
						_ = _580_v31
						_580_v31 = _dafny.SetOf(false, _this.F12)
						var _581_v32 _dafny.Set
						_ = _581_v32
						_581_v32 = _dafny.SetOf(Companion_Default___.Fm2(_this.F12, (_580_v31).Cardinality(), _579_v30, p1, globalState), _dafny.UnicodeSeqOfUtf8Bytes("h"))
						var _582_v33 D0
						_ = _582_v33
						_582_v33 = Companion_D0_.Create_DC1_((p0).Cmp(_dafny.IntOfInt64(-711)) >= 0, (_581_v32).Cardinality())
						var _rhs85 _dafny.Int = (_this).F13()
						_ = _rhs85
						var _rhs86 _dafny.CodePoint = _dafny.CodePoint('c')
						_ = _rhs86
						var _rhs87 D0 = _582_v33
						_ = _rhs87
						var _lhs65 *GlobalState = globalState
						_ = _lhs65
						_lhs65.F1 = _rhs85
						_579_v30 = _rhs86
						_582_v33 = _rhs87
						var _rhs88 _dafny.Int = (_this).F13()
						_ = _rhs88
						var _rhs89 _dafny.Int = p1
						_ = _rhs89
						var _lhs66 *GlobalState = globalState
						_ = _lhs66
						var _lhs67 *GlobalState = globalState
						_ = _lhs67
						_lhs66.F1 = _rhs88
						_lhs67.F1 = _rhs89
						var _583_v34 _dafny.Map
						_ = _583_v34
						_583_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-505), _this.F12)
						var _584_v35 _dafny.Map
						_ = _584_v35
						_584_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, (_this).F13())
						_583_v34 = (_583_v34).Update(((_584_v35).Update(_this.F12, p0)).Cardinality(), _this.F12)
						var _585_v36 _dafny.Array
						_ = _585_v36
						var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
						_ = _nw102
						_585_v36 = _nw102
						var _len0_18 _dafny.Int = _dafny.IntOfInt64(13)
						_ = _len0_18
						var _nw103 _dafny.Array
						_ = _nw103
						if _len0_18.Cmp(_dafny.Zero) == 0 {
							_nw103 = _dafny.NewArray(_len0_18)
						} else {
							var _init18 func(_dafny.Int) _dafny.MultiSet = func(_586_i6 _dafny.Int) _dafny.MultiSet {
								return _dafny.MultiSetOf((_this).F13(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F12)).Cardinality()))).Cardinality())
							}
							_ = _init18
							var _element0_18 = _init18(_dafny.Zero)
							_ = _element0_18
							_nw103 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
							(_nw103).ArraySet1(_element0_18, 0)
							var _nativeLen0_18 = (_len0_18).Int()
							_ = _nativeLen0_18
							for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
								(_nw103).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
							}
						}
						_585_v36 = _nw103
						(_this).F12 = _this.F12
					} else {
						var _587_v37 _dafny.Array
						_ = _587_v37
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(27)
						_ = _len0_19
						var _nw104 _dafny.Array
						_ = _nw104
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw104 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) _dafny.Sequence = (func(_588_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_589_i7 _dafny.Int) _dafny.Sequence {
									return _588_v3
								}
							})(_545_v3)
							_ = _init19
							var _element0_19 = _init19(_dafny.Zero)
							_ = _element0_19
							_nw104 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
							(_nw104).ArraySet1(_element0_19, 0)
							var _nativeLen0_19 = (_len0_19).Int()
							_ = _nativeLen0_19
							for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
								(_nw104).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
							}
						}
						_587_v37 = _nw104
						var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_587_v37), 0))
						_ = _index110
						(_587_v37).ArraySet1(_545_v3, (_index110).Int())
						var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_587_v37), 0))
						_ = _index111
						(_587_v37).ArraySet1(_545_v3, (_index111).Int())
						(globalState).F1 = p1
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_578_v29), 0))
						_ = _index112
						(_578_v29).ArraySet1((_578_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_578_v29), 0))).Int()).(D1), (_index112).Int())
						(globalState).F1 = _dafny.IntOfInt64(66)
						var _590_v38 _dafny.MultiSet
						_ = _590_v38
						_590_v38 = _dafny.MultiSetOf(_this.F12, !(_this.F12), _this.F12, _this.F12)
						var _591_v39 _dafny.MultiSet
						_ = _591_v39
						_591_v39 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("micx")).Cardinality())), p0)
						var _592_v40 _dafny.Map
						_ = _592_v40
						_592_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm11(_591_v39, globalState), _this.F12)
						var _593_v41 D0
						_ = _593_v41
						_593_v41 = Companion_D0_.Create_DC1_(_this.F12, p0)
						var _rhs90 _dafny.MultiSet = ((_590_v38).Union(_dafny.MultiSetFromSeq(Companion_Default___.Fm12(globalState)))).Intersection(_590_v38)
						_ = _rhs90
						var _rhs91 _dafny.Array = (func() _dafny.Array {
							if (func() bool {
								if !((func() bool {
									if (_592_v40).Contains((_593_v41).Dtor_cf2()) {
										return (_592_v40).Get((_593_v41).Dtor_cf2()).(bool)
									}
									return _this.F12
								})()) {
									return false
								}
								return _this.F12
							})() {
								return _543_v1
							}
							return _543_v1
						})()
						_ = _rhs91
						var _lhs68 *GlobalState = globalState
						_ = _lhs68
						_590_v38 = _rhs90
						_lhs68.F5 = _rhs91
					}
					if _dafny.Companion_Sequence_.IsProperPrefixOf(_541_v0, _dafny.UnicodeSeqOfUtf8Bytes("hmmeoafg")) {
						var _594_v42 _dafny.CodePoint
						_ = _594_v42
						_594_v42 = _dafny.CodePoint('g')
						var _595_v43 D0
						_ = _595_v43
						_595_v43 = Companion_D0_.Create_DC0_(_594_v42)
						var _596_v44 D0
						_ = _596_v44
						_596_v44 = Companion_D0_.Create_DC1_(!(_this.F12), (_dafny.Zero).Minus((p2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int)))
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_543_v1), 0))
						_ = _index113
						(_543_v1).ArraySet1(true, (_index113).Int())
						var _597_v45 _dafny.MultiSet
						_ = _597_v45
						_597_v45 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(756), (_this).F13(), (_this).F13(), (p2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int))
						var _598_v46 _dafny.MultiSet
						_ = _598_v46
						_598_v46 = _597_v45
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_543_v1), 0))
						_ = _index114
						var _rhs92 _dafny.Int = (_this).Fm11((_598_v46).Union(_597_v45), globalState)
						_ = _rhs92
						var _rhs93 D0 = _595_v43
						_ = _rhs93
						var _rhs94 D0 = func(_pat_let10_0 D0) D0 {
							return func(_601_dt__update__tmp_h1 D0) D0 {
								return func(_pat_let13_0 bool) D0 {
									return func(_602_dt__update_hcf1_h1 bool) D0 {
										return Companion_D0_.Create_DC1_(_602_dt__update_hcf1_h1, (_601_dt__update__tmp_h1).Dtor_cf2())
									}(_pat_let13_0)
								}(_this.F12)
							}(_pat_let10_0)
						}(func(_pat_let11_0 D0) D0 {
							return func(_599_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let12_0 bool) D0 {
									return func(_600_dt__update_hcf1_h0 bool) D0 {
										return Companion_D0_.Create_DC1_(_600_dt__update_hcf1_h0, (_599_dt__update__tmp_h0).Dtor_cf2())
									}(_pat_let12_0)
								}(_this.F12)
							}(_pat_let11_0)
						}(Companion_D0_.Create_DC1_(_this.F12, p1)))
						_ = _rhs94
						var _rhs95 bool = _this.F12
						_ = _rhs95
						var _lhs69 *GlobalState = globalState
						_ = _lhs69
						var _lhs70 _dafny.Array = _543_v1
						_ = _lhs70
						var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_543_v1), 0))
						_ = _lhs71
						_lhs69.F1 = _rhs92
						_595_v43 = _rhs93
						_596_v44 = _rhs94
						(_lhs70).ArraySet1(_rhs95, (_lhs71).Int())
						var _603_v47 _dafny.Sequence
						_ = _603_v47
						_603_v47 = _dafny.SeqOf(_541_v0, _541_v0, _541_v0, _541_v0)
						(globalState).F1 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32(((_603_v47).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_603_v47).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), _dafny.IntOfInt64(910))
						_541_v0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bvmvcibbo"), _541_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_541_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hlqwmddvn")).Cardinality()), _dafny.IntOfUint32((_541_v0).Cardinality()))).Uint32(), _594_v42), _541_v0)), (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bvmvcibbo"), _541_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_541_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hlqwmddvn")).Cardinality()), _dafny.IntOfUint32((_541_v0).Cardinality()))).Uint32(), _594_v42), _541_v0))).Cardinality()))).Uint32(), _594_v42)
						(globalState).F1 = p0
						var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_543_v1), 0))
						_ = _index115
						(_543_v1).ArraySet1(_this.F12, (_index115).Int())
					} else {
						(globalState).F1 = p1
						(_this).F12 = Companion_Default___.Fm4(_549_v5, globalState)
						var _604_v48 _dafny.Map
						_ = _604_v48
						_604_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, Companion_Default___.Fm4(_548_v4, globalState))
						_604_v48 = (_604_v48).Update(_this.F12, _this.F12)
						(_this).F12 = _this.F12
						var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_543_v1), 0))
						_ = _index116
						(_543_v1).ArraySet1(_this.F12, (_index116).Int())
						var _605_v49 _dafny.CodePoint
						_ = _605_v49
						_605_v49 = _dafny.CodePoint('q')
						var _606_v50 _dafny.Array
						_ = _606_v50
						var _nwElement0_25 _dafny.CodePoint = _605_v49
						_ = _nwElement0_25
						var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(22))
						_ = _nw105
						(_nw105).ArraySet1CodePoint(_nwElement0_25, 0)
						(_nw105).ArraySet1CodePoint(_605_v49, 1)
						(_nw105).ArraySet1CodePoint(_605_v49, 2)
						(_nw105).ArraySet1CodePoint(_605_v49, 3)
						(_nw105).ArraySet1CodePoint(_605_v49, 4)
						(_nw105).ArraySet1CodePoint(_605_v49, 5)
						(_nw105).ArraySet1CodePoint(_605_v49, 6)
						(_nw105).ArraySet1CodePoint(_605_v49, 7)
						(_nw105).ArraySet1CodePoint(_605_v49, 8)
						(_nw105).ArraySet1CodePoint(_605_v49, 9)
						(_nw105).ArraySet1CodePoint(_605_v49, 10)
						(_nw105).ArraySet1CodePoint(_605_v49, 11)
						(_nw105).ArraySet1CodePoint(_605_v49, 12)
						(_nw105).ArraySet1CodePoint(_605_v49, 13)
						(_nw105).ArraySet1CodePoint(_605_v49, 14)
						(_nw105).ArraySet1CodePoint(_605_v49, 15)
						(_nw105).ArraySet1CodePoint(_605_v49, 16)
						(_nw105).ArraySet1CodePoint(_605_v49, 17)
						(_nw105).ArraySet1CodePoint(_605_v49, 18)
						(_nw105).ArraySet1CodePoint(_605_v49, 19)
						(_nw105).ArraySet1CodePoint(_605_v49, 20)
						(_nw105).ArraySet1CodePoint(_605_v49, 21)
						_606_v50 = _nw105
						var _607_v51 _dafny.Map
						_ = _607_v51
						_607_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((p2).Cardinality()))
						var _608_v52 _dafny.Map
						_ = _608_v52
						_608_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_606_v50, (func() _dafny.Int {
							if (_607_v51).Contains(_this.F12) {
								return (_607_v51).Get(_this.F12).(_dafny.Int)
							}
							return p0
						})())
						var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_543_v1), 0))
						_ = _index117
						var _rhs96 bool = (true) && (!(_this.F12))
						_ = _rhs96
						var _rhs97 bool = _this.F12
						_ = _rhs97
						var _rhs98 _dafny.Map = _608_v52
						_ = _rhs98
						var _rhs99 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}((func(_609_v49 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_610_i8 _dafny.Int) _dafny.CodePoint {
								return _609_v49
							}
						})(_605_v49))), _541_v0), _dafny.UnicodeSeqOfUtf8Bytes("anvdwavn")), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_this.F12)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}((func(_611_v49 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_612_i8 _dafny.Int) _dafny.CodePoint {
								return _611_v49
							}
						})(_605_v49))), _541_v0), _dafny.UnicodeSeqOfUtf8Bytes("anvdwavn"))).Cardinality()))).Uint32(), _605_v49)
						_ = _rhs99
						var _rhs100 bool = _this.F12
						_ = _rhs100
						var _lhs72 *C2 = _this
						_ = _lhs72
						var _lhs73 _dafny.Array = _543_v1
						_ = _lhs73
						var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_543_v1), 0))
						_ = _lhs74
						var _lhs75 *C2 = _this
						_ = _lhs75
						_lhs72.F12 = _rhs96
						(_lhs73).ArraySet1(_rhs97, (_lhs74).Int())
						_608_v52 = _rhs98
						_541_v0 = _rhs99
						_lhs75.F12 = _rhs100
					}
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _613_v53 T2
		_ = _613_v53
		var _nw106 *C0 = New_C0_()
		_ = _nw106
		_nw106.Ctor__()
		_613_v53 = _nw106
		var _614_v54 _dafny.Sequence
		_ = _614_v54
		_614_v54 = _dafny.SeqOf(_613_v53, _613_v53, _613_v53, _613_v53, _613_v53)
		var _615_v55 _dafny.Map
		_ = _615_v55
		_615_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, _this.F12)
		var _616_v56 _dafny.MultiSet
		_ = _616_v56
		_616_v56 = _dafny.MultiSetOf((_dafny.Zero).Minus((_615_v55).Cardinality()), p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F12)).Cardinality(), (_this).F13())
		var _617_v57 _dafny.MultiSet
		_ = _617_v57
		_617_v57 = _dafny.MultiSetOf(_dafny.IntOfInt64(623), (_616_v56).Cardinality(), (_this).F13())
		var _618_v58 _dafny.Map
		_ = _618_v58
		_618_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _613_v53)
		var _619_v59 _dafny.Map
		_ = _619_v59
		_619_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(802), _613_v53)
		var _620_v60 _dafny.Map
		_ = _620_v60
		_620_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, _543_v1)
		var _621_v61 _dafny.Array
		_ = _621_v61
		var _nwElement0_26 T2 = _613_v53
		_ = _nwElement0_26
		var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(20))
		_ = _nw107
		(_nw107).ArraySet1(_nwElement0_26, 0)
		(_nw107).ArraySet1(_613_v53, 1)
		(_nw107).ArraySet1(_613_v53, 2)
		(_nw107).ArraySet1(_613_v53, 3)
		(_nw107).ArraySet1(_613_v53, 4)
		(_nw107).ArraySet1((_614_v54).Select((Companion_Default___.SafeIndex((_this).Fm11(_616_v56, globalState), _dafny.IntOfUint32((_614_v54).Cardinality()))).Uint32()).(T2), 5)
		(_nw107).ArraySet1(_613_v53, 6)
		(_nw107).ArraySet1(_613_v53, 7)
		(_nw107).ArraySet1(_613_v53, 8)
		(_nw107).ArraySet1((func() T2 {
			if (_618_v58).Contains(_this.F12) {
				return (_618_v58).Get(_this.F12).(T2)
			}
			return _613_v53
		})(), 9)
		(_nw107).ArraySet1(_613_v53, 10)
		(_nw107).ArraySet1(_613_v53, 11)
		(_nw107).ArraySet1(_613_v53, 12)
		(_nw107).ArraySet1(_613_v53, 13)
		(_nw107).ArraySet1(_613_v53, 14)
		(_nw107).ArraySet1(_613_v53, 15)
		(_nw107).ArraySet1(_613_v53, 16)
		(_nw107).ArraySet1(_613_v53, 17)
		(_nw107).ArraySet1(_613_v53, 18)
		(_nw107).ArraySet1((func() T2 {
			if (_619_v59).Contains((_620_v60).Cardinality()) {
				return (_619_v59).Get((_620_v60).Cardinality()).(T2)
			}
			return _613_v53
		})(), 19)
		_621_v61 = _nw107
		var _622_v62 _dafny.Sequence
		_ = _622_v62
		_622_v62 = _dafny.SeqOf(_621_v61, _621_v61, _621_v61, _621_v61)
		_622_v62 = _dafny.Companion_Sequence_.Concatenate(_622_v62, _dafny.Companion_Sequence_.Concatenate(_622_v62, _622_v62))
		var _623_v63 *C0
		_ = _623_v63
		var _nw108 *C0 = New_C0_()
		_ = _nw108
		_nw108.Ctor__()
		_623_v63 = _nw108
		var _624_v64 T0
		_ = _624_v64
		var _nw109 *C1 = New_C1_()
		_ = _nw109
		_nw109.Ctor__(_this.F12, _this.F12, (_this).F10())
		_624_v64 = _nw109
		var _625_v65 _dafny.Set
		_ = _625_v65
		_625_v65 = _dafny.SetOf(_624_v64)
		var _626_v66 _dafny.Map
		_ = _626_v66
		_626_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_625_v65, false)
		r0 = (_626_v66).Merge(_626_v66)
		return r0
	}
}
func (_this *C2) F13() _dafny.Int {
	{
		return _this._f13
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f16 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f16 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__(f16 bool) {
	{
		(_this)._f16 = f16
	}
}
func (_this *C3) Fm30(globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_this).F16())).Cardinality(), (_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-822))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_627_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-739), (_this).F16())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nfqbkcer")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(956), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _dafny.IntOfInt64(949))).Cardinality())), Companion_D9_.Create_DC23_(_dafny.IntOfInt64(48)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
			var _coll34 = _dafny.NewBuilder()
			_ = _coll34
			for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-12), _dafny.IntOfInt64(-592))); ; {
				_compr_34, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _628_v0 _dafny.Int
				_628_v0 = interface{}(_compr_34).(_dafny.Int)
				if ((_dafny.IntOfInt64(-12)).Cmp(_628_v0) <= 0) && ((_628_v0).Cmp(_dafny.IntOfInt64(-592)) < 0) {
					_coll34.Add((_628_v0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ypcu")).Cardinality())), !((_this).F16()))).Cardinality()))
				}
			}
			return _coll34.ToSet()
		}(), Companion_D9_.Create_DC23_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(304), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uqqfvcsyh")).Cardinality()))).Cardinality())))
	}
}
func (_this *C3) Fm31(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(806)
	}
}
func (_this *C3) M1(p0 _dafny.CodePoint, p1 _dafny.Int, p2 D0, globalState *GlobalState) {
	{
		var _629_v0 _dafny.Set
		_ = _629_v0
		_629_v0 = _dafny.SetOf(p1, _dafny.IntOfInt64(-38))
		var _630_v1 _dafny.Sequence
		_ = _630_v1
		_630_v1 = _dafny.SeqOf(p1)
		(_this).M7(Companion_Default___.Fm4(_629_v0, globalState), _630_v1, globalState)
		var _631_v2 _dafny.Array
		_ = _631_v2
		var _nw110 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw110
		_631_v2 = _nw110
		for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_631_v2), 0))); ; {
			_guard_loop_2, _ok37 := _iter37()
			if !_ok37 {
				break
			}
			var _632_i0 _dafny.Int
			_632_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_632_i0).Sign() != -1) && ((_632_i0).Cmp(_dafny.ArrayLen((_631_v2), 0)) < 0)) {
				(_631_v2).ArraySet1((_this).F16(), (_632_i0).Int())
			}
		}
		var _633_v3 _dafny.Set
		_ = _633_v3
		_633_v3 = _dafny.SetOf(false)
		var _634_v4 _dafny.Sequence
		_ = _634_v4
		_634_v4 = _dafny.SeqOf(!((_this).F16()))
		var _635_v5 *C1
		_ = _635_v5
		var _nw111 *C1 = New_C1_()
		_ = _nw111
		_nw111.Ctor__((_this).F16(), (_this).F16(), _dafny.MultiSetFromSeq(_634_v4))
		_635_v5 = _nw111
		var _636_v6 _dafny.Sequence
		_ = _636_v6
		_636_v6 = _dafny.SeqOf(_635_v5, _635_v5)
		var _637_v7 _dafny.MultiSet
		_ = _637_v7
		_637_v7 = _dafny.MultiSetOf(p1, (_633_v3).Cardinality(), p1, _dafny.IntOfUint32((_636_v6).Cardinality()), p1)
		var _638_v8 _dafny.Array
		_ = _638_v8
		var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
		_ = _nw112
		_638_v8 = _nw112
		var _639_v9 _dafny.Map
		_ = _639_v9
		_639_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_637_v7).Difference(_637_v7), _638_v8)
		_639_v9 = (_639_v9).Update(_637_v7, _638_v8)
		if (_dafny.IntOfInt64(793)).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-839))) <= 0 {
			(globalState).F1 = p1
			var _640_v10 _dafny.Sequence
			_ = _640_v10
			_640_v10 = _dafny.UnicodeSeqOfUtf8Bytes("cxakbjhpc")
			var _641_v11 D1
			_ = _641_v11
			_641_v11 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(748), (_635_v5).F15(), _dafny.IntOfUint32((_640_v10).Cardinality()), (_this).F16(), p1)
			var _642_v12 _dafny.Map
			_ = _642_v12
			_642_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _643_v13 _dafny.Map
			_ = _643_v13
			_643_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_635_v5).F15())
			var _644_v14 _dafny.Array
			_ = _644_v14
			var _nwElement0_27 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_640_v10, _640_v10)
			_ = _nwElement0_27
			var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(11))
			_ = _nw113
			(_nw113).ArraySet1(_nwElement0_27, 0)
			(_nw113).ArraySet1(_640_v10, 1)
			(_nw113).ArraySet1(_640_v10, 2)
			(_nw113).ArraySet1(_640_v10, 3)
			(_nw113).ArraySet1(_640_v10, 4)
			(_nw113).ArraySet1((func() _dafny.Sequence {
				if (_635_v5).F14() {
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_640_v10, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm32((_641_v11).Dtor_cf14(), p0, _642_v12, (_635_v5).F15(), globalState)).Cardinality()), _dafny.IntOfUint32((_640_v10).Cardinality()))).Uint32(), p0), (Companion_Default___.SafeIndex((_643_v13).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_640_v10, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm32((_641_v11).Dtor_cf14(), p0, _642_v12, (_635_v5).F15(), globalState)).Cardinality()), _dafny.IntOfUint32((_640_v10).Cardinality()))).Uint32(), p0)).Cardinality()))).Uint32(), p0)
				}
				return _640_v10
			})(), 5)
			(_nw113).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_640_v10, _640_v10), 6)
			(_nw113).ArraySet1((func() _dafny.Sequence {
				if (_635_v5).F14() {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(555))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg36 _dafny.Int) interface{} {
							return coer36(arg36)
						}
					}((func(_645_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_646_i1 _dafny.Int) _dafny.CodePoint {
							return _645_p0
						}
					})(p0)))
				}
				return _640_v10
			})(), 7)
			(_nw113).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_640_v10, _640_v10), 8)
			(_nw113).ArraySet1(_640_v10, 9)
			(_nw113).ArraySet1(_640_v10, 10)
			_644_v14 = _nw113
			_644_v14 = _644_v14
			var _647_v15 _dafny.Array
			_ = _647_v15
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_20
			var _nw114 _dafny.Array
			_ = _nw114
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw114 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Int = (func(_648_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_649_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_649_i2, _648_p1)
					}
				})(p1)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw114 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw114).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw114).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_647_v15 = _nw114
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_647_v15), 0))
			_ = _index118
			(_647_v15).ArraySet1(p1, (_index118).Int())
			var _650_v16 _dafny.Sequence
			_ = _650_v16
			_650_v16 = _dafny.SeqOf(_640_v10)
			var _651_v17 _dafny.MultiSet
			_ = _651_v17
			_651_v17 = _dafny.MultiSetOf((_635_v5).F15(), (_635_v5).F14(), !(false))
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_647_v15), 0))
			_ = _index119
			var _rhs101 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_this).F16() {
					return _dafny.IntOfInt64(71)
				}
				return _dafny.IntOfUint32((_650_v16).Cardinality())
			})(), (p1).Times((_643_v13).Cardinality())))
			_ = _rhs101
			var _rhs102 _dafny.Int = p1
			_ = _rhs102
			var _rhs103 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm33((func() _dafny.Int {
				if (_651_v17).Contains((_635_v5).F15()) {
					return (_651_v17).Multiplicity((_635_v5).F15())
				}
				return p1
			})(), (p1).Cmp(p1) > 0, (_635_v5).F15(), p1, globalState)).Cardinality())
			_ = _rhs103
			var _rhs104 _dafny.Int = (func() _dafny.Int {
				if (_637_v7).Contains(p1) {
					return (_637_v7).Multiplicity(p1)
				}
				return (p1).Times(_dafny.IntOfUint32((_634_v4).Cardinality()))
			})()
			_ = _rhs104
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 _dafny.Array = _647_v15
			_ = _lhs78
			var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_647_v15), 0))
			_ = _lhs79
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			_lhs76.F1 = _rhs101
			_lhs77.F1 = _rhs102
			(_lhs78).ArraySet1(_rhs103, (_lhs79).Int())
			_lhs80.F1 = _rhs104
			_629_v0 = _629_v0
			(globalState).F1 = (_647_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_647_v15), 0))).Int()).(_dafny.Int)
		} else {
			(globalState).F1 = (func() _dafny.Int {
				if (_637_v7).Contains(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F16() {
						return _dafny.SeqOf((_this).F16())
					}
					return _634_v4
				})()).Cardinality())) {
					return (_637_v7).Multiplicity(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_this).F16() {
							return _dafny.SeqOf((_this).F16())
						}
						return _634_v4
					})()).Cardinality()))
				}
				return p1
			})()
			var _652_v18 bool
			_ = _652_v18
			_652_v18 = false
			var _rhs105 _dafny.Int = p1
			_ = _rhs105
			var _rhs106 bool = (_635_v5).F15()
			_ = _rhs106
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			_lhs81.F1 = _rhs105
			_652_v18 = _rhs106
			var _653_v19 _dafny.Array
			_ = _653_v19
			var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
			_ = _nw115
			_653_v19 = _nw115
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_21
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Sequence = func(_654_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dbl"), _dafny.UnicodeSeqOfUtf8Bytes("x"))
				}
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw116 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw116).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw116).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_653_v19 = _nw116
			var _655_v20 _dafny.MultiSet
			_ = _655_v20
			_655_v20 = _dafny.MultiSetOf((_635_v5).F14())
			_655_v20 = (_655_v20).Difference(_655_v20)
			var _656_v21 _dafny.Sequence
			_ = _656_v21
			_656_v21 = _dafny.UnicodeSeqOfUtf8Bytes("xwcdnmpw")
			var _657_v22 _dafny.Map
			_ = _657_v22
			_657_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_this).F16())).Cardinality(), _656_v21)
			var _rhs107 bool = (p1).Cmp(p1) == 0
			_ = _rhs107
			var _rhs108 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Plus(p1), _dafny.UnicodeSeqOfUtf8Bytes("oumom"))
			_ = _rhs108
			_652_v18 = _rhs107
			_657_v22 = _rhs108
		}
		var _hi5 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, p1)
		_ = _hi5
		for _658_i4 := Companion_Default___.SafeModuloInt(p1, p1); _658_i4.Cmp(_hi5) < 0; _658_i4 = _658_i4.Plus(_dafny.One) {
			var _659_v23 _dafny.Array
			_ = _659_v23
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_22
			var _nw117 _dafny.Array
			_ = _nw117
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw117 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = (func(_660_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_661_i5 _dafny.Int) _dafny.Int {
						return (_661_i5).Minus(_660_p1)
					}
				})(p1)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw117 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw117).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw117).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_659_v23 = _nw117
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_659_v23), 0))
			_ = _index120
			(_659_v23).ArraySet1(_658_i4, (_index120).Int())
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_659_v23), 0))
			_ = _index121
			(_659_v23).ArraySet1(p1, (_index121).Int())
			var _662_v24 _dafny.CodePoint
			_ = _662_v24
			_662_v24 = _dafny.CodePoint('h')
			var _663_v25 _dafny.Sequence
			_ = _663_v25
			_663_v25 = _dafny.UnicodeSeqOfUtf8Bytes("qs")
			var _664_v26 _dafny.Map
			_ = _664_v26
			_664_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_dafny.IntOfUint32((_663_v25).Cardinality())).Cmp(_658_i4) >= 0), (_635_v5).F15())
			var _665_v27 _dafny.Map
			_ = _665_v27
			_665_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F16())
			var _666_v28 D9
			_ = _666_v28
			_666_v28 = Companion_D9_.Create_DC22_(p1, (_this).F16(), _665_v27)
			var _667_v29 _dafny.Map
			_ = _667_v29
			_667_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_635_v5).F15(), _665_v27)
			var _668_v30 _dafny.Sequence
			_ = _668_v30
			_668_v30 = _dafny.SeqOf(_665_v27, Companion_Default___.Fm34((_666_v28).Dtor_cf50(), _634_v4, (_630_v1).Select((Companion_Default___.SafeIndex((_629_v0).Cardinality(), _dafny.IntOfUint32((_630_v1).Cardinality()))).Uint32()).(_dafny.Int), (_635_v5).F15(), globalState), _665_v27, _665_v27, (func() _dafny.Map {
				if (_667_v29).Contains((_635_v5).F14()) {
					return (_667_v29).Get((_635_v5).F14()).(_dafny.Map)
				}
				return (func() _dafny.Map {
					if (_667_v29).Contains((_635_v5).F15()) {
						return (_667_v29).Get((_635_v5).F15()).(_dafny.Map)
					}
					return _665_v27
				})()
			})())
			var _669_v31 _dafny.Map
			_ = _669_v31
			_669_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_635_v5).F14(), p1)
			var _670_v32 D8
			_ = _670_v32
			_670_v32 = Companion_D8_.Create_DC20_(_662_v24, (_635_v5).F15(), Companion_Default___.Fm4(_dafny.SetOf((_659_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_659_v23), 0))).Int()).(_dafny.Int), _658_i4, p1), globalState), _669_v31, (_635_v5).F14())
			var _671_v33 _dafny.Map
			_ = _671_v33
			_671_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_670_v32, p1)
			var _672_v34 _dafny.Map
			_ = _672_v34
			_672_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC20_(_662_v24, (_635_v5).F15(), (_635_v5).F15(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_635_v5).F15(), p1), (_this).F16()), _658_i4))
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_659_v23), 0))
			_ = _index122
			var _rhs109 _dafny.Int = (_dafny.Zero).Minus((_this).Fm31(_662_v24, (_this).F16(), (func() _dafny.Map {
				if (_635_v5).F15() {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_659_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_659_v23), 0))).Int()).(_dafny.Int), _671_v33)
				}
				return _672_v34
			})(), globalState))
			_ = _rhs109
			var _rhs110 _dafny.CodePoint = Companion_Default___.Fm1(globalState)
			_ = _rhs110
			var _rhs111 _dafny.Map = (_664_v26).Merge(_664_v26)
			_ = _rhs111
			var _rhs112 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_668_v30, _668_v30), _668_v30)
			_ = _rhs112
			var _lhs82 _dafny.Array = _659_v23
			_ = _lhs82
			var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_659_v23), 0))
			_ = _lhs83
			(_lhs82).ArraySet1(_rhs109, (_lhs83).Int())
			_662_v24 = _rhs110
			_664_v26 = _rhs111
			_668_v30 = _rhs112
			_663_v25 = _663_v25
			var _673_v35 bool
			_ = _673_v35
			_673_v35 = true
			_673_v35 = (_this).F16()
		}
		var _674_v36 _dafny.Array
		_ = _674_v36
		var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw118
		_674_v36 = _nw118
		var _675_v37 _dafny.Sequence
		_ = _675_v37
		_675_v37 = _dafny.UnicodeSeqOfUtf8Bytes("nqgf")
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_674_v36), 0))
		_ = _index123
		(_674_v36).ArraySet1(_dafny.IntOfUint32((_675_v37).Cardinality()), (_index123).Int())
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_674_v36), 0))
		_ = _index124
		(_674_v36).ArraySet1(p1, (_index124).Int())
	}
}
func (_this *C3) M7(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) {
	{
		var _676_v0 _dafny.Int
		_ = _676_v0
		_676_v0 = _dafny.IntOfInt64(53)
		var _677_v1 _dafny.Map
		_ = _677_v1
		_677_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_676_v0).Plus(_676_v0), _676_v0)
		_677_v1 = (_677_v1).Update(_676_v0, _676_v0)
		var _hi6 _dafny.Int = _676_v0
		_ = _hi6
		for _678_i0 := Companion_Default___.SafeDivisionInt(_676_v0, _676_v0); _678_i0.Cmp(_hi6) < 0; _678_i0 = _678_i0.Plus(_dafny.One) {
			_676_v0 = Companion_Default___.SafeModuloInt(_678_i0, _dafny.IntOfInt64(689))
			var _679_v2 _dafny.Set
			_ = _679_v2
			_679_v2 = _dafny.SetOf((_678_i0).Minus(_676_v0), _dafny.IntOfInt64(597), _678_i0, _678_i0)
			(globalState).F1 = (_679_v2).Cardinality()
			if (_this).F16() {
				(globalState).F1 = Companion_Default___.SafeModuloInt(_678_i0, _676_v0)
				(globalState).F1 = _676_v0
				(globalState).F1 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_678_i0), (_dafny.Zero).Minus(((func() _dafny.Set {
					if p0 {
						return _679_v2
					}
					return _dafny.SetOf(_678_i0)
				})()).Cardinality()))
				var _680_v3 _dafny.MultiSet
				_ = _680_v3
				_680_v3 = _dafny.MultiSetOf(p0)
				var _681_v4 *C1
				_ = _681_v4
				var _nw119 *C1 = New_C1_()
				_ = _nw119
				_nw119.Ctor__(false, p0, _680_v3)
				_681_v4 = _nw119
				var _682_v5 _dafny.Sequence
				_ = _682_v5
				_682_v5 = _dafny.SeqOf(_681_v4, _681_v4, _681_v4)
				_676_v0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_682_v5).Cardinality())), (_dafny.Zero).Minus(_678_i0))
				var _683_v6 _dafny.Array
				_ = _683_v6
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_23
				var _nw120 _dafny.Array
				_ = _nw120
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw120 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Sequence = func(_684_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("wruaxewg")
					}
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw120 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw120).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw120).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_683_v6 = _nw120
				var _685_v7 _dafny.Array
				_ = _685_v7
				var _nwElement0_28 _dafny.Int = _676_v0
				_ = _nwElement0_28
				var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(12))
				_ = _nw121
				(_nw121).ArraySet1(_nwElement0_28, 0)
				(_nw121).ArraySet1(_dafny.IntOfInt64(174), 1)
				(_nw121).ArraySet1(_676_v0, 2)
				(_nw121).ArraySet1(_676_v0, 3)
				(_nw121).ArraySet1(_678_i0, 4)
				(_nw121).ArraySet1(_678_i0, 5)
				(_nw121).ArraySet1(_676_v0, 6)
				(_nw121).ArraySet1(_678_i0, 7)
				(_nw121).ArraySet1(_dafny.IntOfInt64(687), 8)
				(_nw121).ArraySet1(_678_i0, 9)
				(_nw121).ArraySet1(_678_i0, 10)
				(_nw121).ArraySet1(_676_v0, 11)
				_685_v7 = _nw121
				var _686_v8 _dafny.Array
				_ = _686_v8
				var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
				_ = _nw122
				_686_v8 = _nw122
				var _687_v9 D5
				_ = _687_v9
				_687_v9 = Companion_D5_.Create_DC13_(_678_i0, _685_v7, _686_v8, _678_i0)
				var _688_v10 _dafny.Map
				_ = _688_v10
				_688_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_683_v6, _687_v9)
				_688_v10 = (_688_v10).Update(_683_v6, _687_v9)
			} else {
				var _689_v11 D9
				_ = _689_v11
				_689_v11 = Companion_D9_.Create_DC24_(_678_i0)
				var _690_v12 _dafny.Map
				_ = _690_v12
				_690_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_689_v11, (_this).F16())
				(globalState).F1 = (p1).Select((Companion_Default___.SafeIndex((_690_v12).Cardinality(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.Int)
				var _691_v13 *C0
				_ = _691_v13
				var _nw123 *C0 = New_C0_()
				_ = _nw123
				_nw123.Ctor__()
				_691_v13 = _nw123
				_691_v13 = _691_v13
				var _692_v14 _dafny.MultiSet
				_ = _692_v14
				_692_v14 = _dafny.MultiSetOf(p0)
				_692_v14 = _692_v14
				var _693_v15 *C2
				_ = _693_v15
				var _nw124 *C2 = New_C2_()
				_ = _nw124
				_nw124.Ctor__(p0, _676_v0, _692_v14)
				_693_v15 = _nw124
				(globalState).F1 = _676_v0
			}
			var _694_v16 *C0
			_ = _694_v16
			var _nw125 *C0 = New_C0_()
			_ = _nw125
			_nw125.Ctor__()
			_694_v16 = _nw125
		}
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("edunij"), _dafny.UnicodeSeqOfUtf8Bytes("rmlmct"))).Cardinality())
		_ = _hi7
		for _695_i2 := _676_v0; _695_i2.Cmp(_hi7) < 0; _695_i2 = _695_i2.Plus(_dafny.One) {
			var _696_v17 _dafny.Set
			_ = _696_v17
			_696_v17 = _dafny.SetOf(!(true), (_this).F16())
			var _697_v18 D0
			_ = _697_v18
			_697_v18 = Companion_D0_.Create_DC1_(p0, _676_v0)
			var _698_v19 _dafny.Array
			_ = _698_v19
			var _nwElement0_29 _dafny.Int = _695_i2
			_ = _nwElement0_29
			var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(4))
			_ = _nw126
			(_nw126).ArraySet1(_nwElement0_29, 0)
			(_nw126).ArraySet1(_695_i2, 1)
			(_nw126).ArraySet1((_dafny.Zero).Minus((_696_v17).Cardinality()), 2)
			(_nw126).ArraySet1((_697_v18).Dtor_cf2(), 3)
			_698_v19 = _nw126
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_698_v19), 0))
			_ = _index125
			(_698_v19).ArraySet1(_676_v0, (_index125).Int())
			var _699_v20 _dafny.Sequence
			_ = _699_v20
			_699_v20 = _dafny.SeqOf(p0, (_this).F16())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_698_v19), 0))
			_ = _index126
			(_698_v19).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_this).F16() {
					return _699_v20
				}
				return _dafny.Companion_Sequence_.Update(_699_v20, (Companion_Default___.SafeIndex(_695_i2, _dafny.IntOfUint32((_699_v20).Cardinality()))).Uint32(), (_this).F16())
			})()).Cardinality()))).Plus(_676_v0), (_index126).Int())
			var _700_v21 _dafny.Array
			_ = _700_v21
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_24
			var _nw127 _dafny.Array
			_ = _nw127
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw127 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Map = (func(_701_v0 _dafny.Int, _702_i2 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_703_i3 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_701_v0, _702_i2)
					}
				})(_676_v0, _695_i2)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw127 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw127).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw127).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_700_v21 = _nw127
			_700_v21 = _700_v21
			var _704_v22 _dafny.MultiSet
			_ = _704_v22
			_704_v22 = _dafny.MultiSetOf(_676_v0)
			_676_v0 = (func() _dafny.Int {
				if (_704_v22).Contains((p1).Select((Companion_Default___.SafeIndex(_676_v0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.Int)) {
					return (_704_v22).Multiplicity((p1).Select((Companion_Default___.SafeIndex(_676_v0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.Int))
				}
				return Companion_Default___.SafeModuloInt(_695_i2, _dafny.IntOfInt64(409))
			})()
			var _705_v23 _dafny.CodePoint
			_ = _705_v23
			_705_v23 = _dafny.CodePoint('m')
			var _706_v24 _dafny.Sequence
			_ = _706_v24
			_706_v24 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_this).F16()))
			var _707_v25 _dafny.Map
			_ = _707_v25
			_707_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _676_v0)
			var _708_v26 D8
			_ = _708_v26
			_708_v26 = Companion_D8_.Create_DC20_(_705_v23, (_this).F16(), !((_this).F16()), _707_v25, p0)
			var _709_v27 _dafny.Map
			_ = _709_v27
			_709_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_708_v26, _676_v0)
			var _710_v28 _dafny.Map
			_ = _710_v28
			_710_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_706_v24).Cardinality()), _709_v27)
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_698_v19), 0))
			_ = _index127
			(_698_v19).ArraySet1((_this).Fm31(_705_v23, (_dafny.IntOfUint32((p1).Cardinality())).Cmp((_698_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_698_v19), 0))).Int()).(_dafny.Int)) >= 0, _710_v28, globalState), (_index127).Int())
		}
		var _711_v29 _dafny.MultiSet
		_ = _711_v29
		_711_v29 = _dafny.MultiSetOf(_676_v0, _676_v0)
		var _712_v30 _dafny.Array
		_ = _712_v30
		var _nwElement0_30 _dafny.MultiSet = _711_v29
		_ = _nwElement0_30
		var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(4))
		_ = _nw128
		(_nw128).ArraySet1(_nwElement0_30, 0)
		(_nw128).ArraySet1(_711_v29, 1)
		(_nw128).ArraySet1((_711_v29).Difference(_dafny.MultiSetFromSeq(p1)), 2)
		(_nw128).ArraySet1(_dafny.MultiSetFromSeq(p1), 3)
		_712_v30 = _nw128
		_712_v30 = _712_v30
		var _713_v31 _dafny.Map
		_ = _713_v31
		_713_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_677_v1, _676_v0)
		var _714_v32 bool
		_ = _714_v32
		_714_v32 = true
		var _715_v33 _dafny.Sequence
		_ = _715_v33
		_715_v33 = _dafny.SeqOf(p0)
		var _716_v34 _dafny.Sequence
		_ = _716_v34
		_716_v34 = _dafny.SeqOf(_715_v33)
		var _717_v35 _dafny.Set
		_ = _717_v35
		_717_v35 = _dafny.SetOf(p0)
		var _rhs113 _dafny.Map = (_713_v31).Update(((_677_v1).Update(_dafny.IntOfUint32(((_716_v34).Select((Companion_Default___.SafeIndex(_676_v0, _dafny.IntOfUint32((_716_v34).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _676_v0)).Merge(_677_v1), ((_717_v35).Intersection(_717_v35)).Cardinality())
		_ = _rhs113
		var _rhs114 bool = (_this).F16()
		_ = _rhs114
		_713_v31 = _rhs113
		_714_v32 = _rhs114
		var _718_v36 _dafny.CodePoint
		_ = _718_v36
		_718_v36 = _dafny.CodePoint('l')
		var _719_v37 _dafny.Array
		_ = _719_v37
		var _nw129 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw129
		_719_v37 = _nw129
		var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_719_v37), 0))
		_ = _index128
		(_719_v37).ArraySet1(_714_v32, (_index128).Int())
		var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_719_v37), 0))
		_ = _index129
		var _rhs115 _dafny.CodePoint = Companion_Default___.Fm1(globalState)
		_ = _rhs115
		var _rhs116 bool = p0
		_ = _rhs116
		var _lhs84 _dafny.Array = _719_v37
		_ = _lhs84
		var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_719_v37), 0))
		_ = _lhs85
		_718_v36 = _rhs115
		(_lhs84).ArraySet1(_rhs116, (_lhs85).Int())
	}
}
func (_this *C3) M8(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _720_v0 bool
		_ = _720_v0
		_720_v0 = false
		_720_v0 = (true) && ((func() bool {
			if true {
				return p2
			}
			return p2
		})())
		var _721_v1 _dafny.Sequence
		_ = _721_v1
		_721_v1 = _dafny.UnicodeSeqOfUtf8Bytes("le")
		var _722_v2 _dafny.Map
		_ = _722_v2
		_722_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_721_v1).Cardinality()))
		var _723_v3 _dafny.Set
		_ = _723_v3
		_723_v3 = _dafny.SetOf((_722_v2).Cardinality())
		var _724_v4 _dafny.MultiSet
		_ = _724_v4
		_724_v4 = _dafny.MultiSetOf(p3)
		var _725_v5 _dafny.MultiSet
		_ = _725_v5
		_725_v5 = _dafny.MultiSetOf(p0, (_this).F16())
		var _726_v6 *C1
		_ = _726_v6
		var _nw130 *C1 = New_C1_()
		_ = _nw130
		_nw130.Ctor__((p2) && (Companion_Default___.Fm4(_723_v3, globalState)), (_724_v4).IsProperSubsetOf(_724_v4), (_725_v5).Update((_this).F16(), Companion_Default___.Abs((_dafny.Zero).Minus(p3))))
		_726_v6 = _nw130
		if Companion_Default___.Fm4(_723_v3, globalState) {
			var _727_v7 _dafny.Array
			_ = _727_v7
			var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw131
			_727_v7 = _nw131
			var _728_v8 _dafny.Map
			_ = _728_v8
			_728_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, false)
			var _729_v9 _dafny.Map
			_ = _729_v9
			_729_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (func() bool {
				if (_728_v8).Contains(_dafny.IntOfInt64(593)) {
					return (_728_v8).Get(_dafny.IntOfInt64(593)).(bool)
				}
				return true
			})())
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))
			_ = _index130
			(_727_v7).ArraySet1(((_729_v9).Merge((_729_v9).Update(!(p2), p2))).Cardinality(), (_index130).Int())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))
			_ = _index131
			(_727_v7).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_721_v1, _721_v1)).Cardinality()), p1), (_index131).Int())
			var _730_v10 _dafny.Sequence
			_ = _730_v10
			_730_v10 = _dafny.SeqOf((_726_v6).F14(), ((_727_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))).Int()).(_dafny.Int)).Cmp(p1) >= 0, (_726_v6).F14(), !((_this).F16()), (_726_v6).F14())
			_730_v10 = (func() _dafny.Sequence {
				if false {
					return _730_v10
				}
				return (func() _dafny.Sequence {
					if (_726_v6).F14() {
						return _730_v10
					}
					return _730_v10
				})()
			})()
			var _731_v11 D8
			_ = _731_v11
			_731_v11 = Companion_D8_.Create_DC19_(_728_v8)
			var _732_v12 _dafny.Sequence
			_ = _732_v12
			_732_v12 = _dafny.SeqOf(_731_v11, _731_v11)
			var _733_v13 _dafny.Sequence
			_ = _733_v13
			_733_v13 = _dafny.SeqOf(_732_v12, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-57))).Uint32(), func(coer37 func(_dafny.Int) D8) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_734_v11 D8) func(_dafny.Int) D8 {
				return func(_735_i0 _dafny.Int) D8 {
					return _734_v11
				}
			})(_731_v11))), _732_v12, _732_v12)
			var _736_v15 _dafny.Set
			_ = _736_v15
			_736_v15 = _dafny.SetOf(_732_v12)
			if ((_736_v15).Union(_736_v15)).IsProperSubsetOf(func() _dafny.Set {
				var _coll35 = _dafny.NewBuilder()
				_ = _coll35
				for _iter38 := _dafny.Iterate((_733_v13).Elements()); ; {
					_compr_35, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _737_v14 _dafny.Sequence
					_737_v14 = interface{}(_compr_35).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_733_v13, _737_v14) {
						_coll35.Add(_737_v14)
					}
				}
				return _coll35.ToSet()
			}()) {
				var _738_v16 *C1
				_ = _738_v16
				var _nw132 *C1 = New_C1_()
				_ = _nw132
				_nw132.Ctor__(p2, (_726_v6).F14(), _dafny.MultiSetOf(p2, (_726_v6).F14(), p2))
				_738_v16 = _nw132
				(globalState).F1 = (p3).Times((p3).Times((_727_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))).Int()).(_dafny.Int)))
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))
				_ = _index132
				(_727_v7).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p3)), (_index132).Int())
				r0 = (p1).Plus(p3)
				_720_v0 = (_738_v16).F15()
			} else {
				_725_v5 = _725_v5
				_720_v0 = (_726_v6).F14()
				var _739_v17 _dafny.Array
				_ = _739_v17
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
				_ = _nw133
				_739_v17 = _nw133
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_739_v17), 0))
				_ = _index133
				(_739_v17).ArraySet1(_721_v1, (_index133).Int())
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_739_v17), 0))
				_ = _index134
				(_739_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_721_v1, _dafny.UnicodeSeqOfUtf8Bytes("v")), _721_v1), (_index134).Int())
				_720_v0 = p0
				var _nw134 *C1 = New_C1_()
				_ = _nw134
				_nw134.Ctor__((_726_v6).F14(), (_730_v10).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_730_v10).Cardinality()))).Uint32()).(bool), _725_v5)
				_726_v6 = _nw134
			}
			var _740_v18 D0
			_ = _740_v18
			_740_v18 = Companion_D0_.Create_DC1_((_this).F16(), (_727_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))).Int()).(_dafny.Int))
			var _741_v20 _dafny.Map
			_ = _741_v20
			_741_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(395), (_727_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))).Int()).(_dafny.Int))
			var _742_v21 _dafny.MultiSet
			_ = _742_v21
			_742_v21 = _dafny.MultiSetOf(func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter39 := _dafny.Iterate((_741_v20).Keys().Elements()); ; {
					_compr_36, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _743_v19 _dafny.Int
					_743_v19 = interface{}(_compr_36).(_dafny.Int)
					if (_741_v20).Contains(_743_v19) {
						_coll36.Add(Companion_Default___.SafeModuloInt(_743_v19, p1), _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))
					}
				}
				return _coll36.ToMap()
			}())
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_727_v7), 0))
			_ = _index135
			(_727_v7).ArraySet1(((_740_v18).Dtor_cf2()).Plus((func() _dafny.Int {
				if (_742_v21).Contains(_741_v20) {
					return (_742_v21).Multiplicity(_741_v20)
				}
				return p3
			})()), (_index135).Int())
			var _744_v22 *C0
			_ = _744_v22
			var _nw135 *C0 = New_C0_()
			_ = _nw135
			_nw135.Ctor__()
			_744_v22 = _nw135
		} else {
			var _745_v23 _dafny.Array
			_ = _745_v23
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_25
			var _nw136 _dafny.Array
			_ = _nw136
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw136 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Map = (func(_746_p0 bool, _747_v3 _dafny.Set) func(_dafny.Int) _dafny.Map {
					return func(_748_i1 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_746_p0, (_747_v3).Cardinality())
					}
				})(p0, _723_v3)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw136 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw136).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw136).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_745_v23 = _nw136
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_745_v23), 0))
			_ = _index136
			(_745_v23).ArraySet1(_722_v2, (_index136).Int())
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_745_v23), 0))
			_ = _index137
			(_745_v23).ArraySet1(_722_v2, (_index137).Int())
			r0 = (p3).Plus(_dafny.IntOfInt64(714))
			var _749_v24 T2
			_ = _749_v24
			var _nw137 *C0 = New_C0_()
			_ = _nw137
			_nw137.Ctor__()
			_749_v24 = _nw137
			var _750_v25 _dafny.Map
			_ = _750_v25
			_750_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_749_v24, true)
			_750_v25 = (_750_v25).Update(_749_v24, p2)
			var _751_v26 _dafny.Map
			_ = _751_v26
			_751_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cmp(_dafny.IntOfInt64(952)) > 0, (_726_v6).F14())
			_751_v26 = (_751_v26).Update(!(_dafny.Companion_Sequence_.Equal(_721_v1, _721_v1)), _720_v0)
			if (p0) == ((_726_v6).F15()) {
				_720_v0 = (_726_v6).F15()
				(globalState).F1 = p1
				var _752_v27 _dafny.Map
				_ = _752_v27
				_752_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_726_v6).F14())
				var _753_v28 _dafny.Sequence
				_ = _753_v28
				_753_v28 = _dafny.SeqOf(p1, p3)
				var _754_v29 D1
				_ = _754_v29
				_754_v29 = Companion_D1_.Create_DC4_(p2, (_726_v6).F15(), p1, _753_v28, p3)
				var _755_v30 _dafny.Sequence
				_ = _755_v30
				_755_v30 = _dafny.SeqOf(p0, (_726_v6).F14())
				var _756_v31 _dafny.Array
				_ = _756_v31
				var _nwElement0_31 _dafny.Int = ((_752_v27).Cardinality()).Minus(p1)
				_ = _nwElement0_31
				var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(26))
				_ = _nw138
				(_nw138).ArraySet1(_nwElement0_31, 0)
				(_nw138).ArraySet1(p3, 1)
				(_nw138).ArraySet1(((_751_v26).Merge(_751_v26)).Cardinality(), 2)
				(_nw138).ArraySet1(p1, 3)
				(_nw138).ArraySet1((_725_v5).Cardinality(), 4)
				(_nw138).ArraySet1(p1, 5)
				(_nw138).ArraySet1((_754_v29).Dtor_cf9(), 6)
				(_nw138).ArraySet1((_749_v24).Fm6(p2, p1, globalState), 7)
				(_nw138).ArraySet1(p1, 8)
				(_nw138).ArraySet1((func() _dafny.Int {
					if (_724_v4).Contains(p3) {
						return (_724_v4).Multiplicity(p3)
					}
					return p3
				})(), 9)
				(_nw138).ArraySet1(p1, 10)
				(_nw138).ArraySet1(((_dafny.Zero).Minus(p1)).Times((_725_v5).Cardinality()), 11)
				(_nw138).ArraySet1(p3, 12)
				(_nw138).ArraySet1(p1, 13)
				(_nw138).ArraySet1(_dafny.IntOfUint32((_755_v30).Cardinality()), 14)
				(_nw138).ArraySet1(p3, 15)
				(_nw138).ArraySet1((_751_v26).Cardinality(), 16)
				(_nw138).ArraySet1(p3, 17)
				(_nw138).ArraySet1(p1, 18)
				(_nw138).ArraySet1(p3, 19)
				(_nw138).ArraySet1(p3, 20)
				(_nw138).ArraySet1(p1, 21)
				(_nw138).ArraySet1((p1).Plus(_dafny.IntOfUint32((_721_v1).Cardinality())), 22)
				(_nw138).ArraySet1(p1, 23)
				(_nw138).ArraySet1(p3, 24)
				(_nw138).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), 25)
				_756_v31 = _nw138
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_756_v31), 0))
				_ = _index138
				(_756_v31).ArraySet1(p1, (_index138).Int())
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_756_v31), 0))
				_ = _index139
				(_756_v31).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p3), (_index139).Int())
				_720_v0 = (_this).F16()
				_720_v0 = ((_dafny.Zero).Minus((_756_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_756_v31), 0))).Int()).(_dafny.Int))).Cmp((_723_v3).Cardinality()) < 0
			} else {
				_726_v6 = _726_v6
				(globalState).F1 = p3
				var _757_v32 _dafny.Sequence
				_ = _757_v32
				_757_v32 = _dafny.SeqOf(_dafny.IntOfUint32((_721_v1).Cardinality()))
				var _758_v33 _dafny.Sequence
				_ = _758_v33
				_758_v33 = _dafny.SeqOf(p2)
				var _759_v34 *C2
				_ = _759_v34
				var _nw139 *C2 = New_C2_()
				_ = _nw139
				_nw139.Ctor__(_720_v0, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-699), _dafny.IntOfUint32((_757_v32).Cardinality())), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_758_v33, _758_v33)))
				_759_v34 = _nw139
				var _760_v35 _dafny.Array
				_ = _760_v35
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_26
				var _nw140 _dafny.Array
				_ = _nw140
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw140 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Int = (func(_761_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_762_i2 _dafny.Int) _dafny.Int {
							return (_762_i2).Plus(_761_p3)
						}
					})(p3)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw140 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw140).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw140).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_760_v35 = _nw140
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_760_v35), 0))
				_ = _index140
				(_760_v35).ArraySet1((p1).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("et")).Cardinality())), (_index140).Int())
				var _763_v36 _dafny.CodePoint
				_ = _763_v36
				_763_v36 = _dafny.CodePoint('f')
				var _764_v37 D0
				_ = _764_v37
				_764_v37 = Companion_D0_.Create_DC0_(_763_v36)
				var _765_v38 D0
				_ = _765_v38
				_765_v38 = Companion_D0_.Create_DC1_((_726_v6).F15(), p3)
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_760_v35), 0))
				_ = _index141
				(_760_v35).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("r"), _721_v1), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p3, p3), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("r"), _721_v1)).Cardinality()))).Uint32(), (_759_v34).Fm10(_764_v37, Companion_D0_.Create_DC2_(_765_v38), globalState))).Cardinality()), (_index141).Int())
				r0 = (_759_v34).F13()
			}
		}
		var _766_v39 _dafny.Sequence
		_ = _766_v39
		_766_v39 = _dafny.SeqOf(p2)
		var _767_v40 *C0
		_ = _767_v40
		var _nw141 *C0 = New_C0_()
		_ = _nw141
		_nw141.Ctor__()
		_767_v40 = _nw141
		var _768_v41 _dafny.Sequence
		_ = _768_v41
		_768_v41 = _dafny.SeqOf(_767_v40)
		var _769_v42 _dafny.Sequence
		_ = _769_v42
		_769_v42 = _dafny.SeqOf(_768_v41)
		var _770_v43 _dafny.Map
		_ = _770_v43
		_770_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_766_v39, !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_769_v42, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_769_v42).Cardinality()))).Uint32(), _768_v41), _768_v41))
		_770_v43 = (_770_v43).Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm33(p3, p0, (_726_v6).F15(), p3, globalState), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p2), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))).Uint32(), (_767_v40).Fm13(p2, globalState)), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p2), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))).Uint32(), (_767_v40).Fm13(p2, globalState))).Cardinality()))).Uint32(), p0)), _720_v0)
		var _771_i3 _dafny.Int
		_ = _771_i3
		_771_i3 = _dafny.Zero
		{
			for (_726_v6).F14() {
				{
					if (_771_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_771_i3 = (_771_i3).Plus(_dafny.One)
					var _772_v44 _dafny.Sequence
					_ = _772_v44
					_772_v44 = _dafny.SeqOf(_721_v1)
					(globalState).F1 = _dafny.IntOfUint32((_772_v44).Cardinality())
					var _773_v45 *C2
					_ = _773_v45
					var _nw142 *C2 = New_C2_()
					_ = _nw142
					_nw142.Ctor__((_726_v6).F15(), _dafny.IntOfInt64(677), _dafny.MultiSetFromSeq(_766_v39))
					_773_v45 = _nw142
					var _774_v46 _dafny.Array
					_ = _774_v46
					var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
					_ = _nw143
					_774_v46 = _nw143
					var _775_v47 _dafny.Set
					_ = _775_v47
					_775_v47 = _dafny.SetOf((_767_v40).Fm13(_773_v45.F12, globalState))
					var _776_v48 _dafny.Sequence
					_ = _776_v48
					_776_v48 = _dafny.SeqOf((_773_v45).F13(), (_775_v47).Cardinality(), p3)
					var _777_v49 _dafny.Array
					_ = _777_v49
					var _nwElement0_32 _dafny.Int = p3
					_ = _nwElement0_32
					var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(8))
					_ = _nw144
					(_nw144).ArraySet1(_nwElement0_32, 0)
					(_nw144).ArraySet1(p1, 1)
					(_nw144).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_766_v39).Cardinality())), 2)
					(_nw144).ArraySet1(p1, 3)
					(_nw144).ArraySet1(_dafny.IntOfUint32((_776_v48).Cardinality()), 4)
					(_nw144).ArraySet1((_773_v45).F13(), 5)
					(_nw144).ArraySet1((_773_v45).F13(), 6)
					(_nw144).ArraySet1((_723_v3).Cardinality(), 7)
					_777_v49 = _nw144
					var _778_v50 _dafny.Map
					_ = _778_v50
					_778_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v46, _777_v49)
					_778_v50 = _778_v50
					_720_v0 = (_767_v40).Fm5((_775_v47).Intersection(_775_v47), p3, (_722_v2).Cardinality(), globalState)
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _779_v51 *C1
		_ = _779_v51
		var _nw145 *C1 = New_C1_()
		_ = _nw145
		_nw145.Ctor__(p2, (_this).F16(), (_725_v5).Update(p0, Companion_Default___.Abs(p3)))
		_779_v51 = _nw145
		r0 = p1
		return r0
	}
}
func (_this *C3) F16() bool {
	{
		return _this._f16
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	dummy byte
}

func New_C4_() *C4 {
	_this := C4{}

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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__() {
	{
	}
}
func (_this *C4) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(921)
	}
}
func (_this *C4) M1(p0 _dafny.CodePoint, p1 _dafny.Int, p2 D0, globalState *GlobalState) {
	{
		var _780_v0 bool
		_ = _780_v0
		_780_v0 = false
		var _781_i0 _dafny.Int
		_ = _781_i0
		_781_i0 = _dafny.Zero
		{
			for _780_v0 {
				{
					if (_781_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_781_i0 = (_781_i0).Plus(_dafny.One)
					var _782_v1 _dafny.Set
					_ = _782_v1
					_782_v1 = _dafny.SetOf((_dafny.Zero).Minus(p1), (p1).Minus(p1))
					(globalState).F1 = (_782_v1).Cardinality()
					var _783_v2 _dafny.Array
					_ = _783_v2
					var _len0_27 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_27
					var _nw146 _dafny.Array
					_ = _nw146
					if _len0_27.Cmp(_dafny.Zero) == 0 {
						_nw146 = _dafny.NewArray(_len0_27)
					} else {
						var _init27 func(_dafny.Int) bool = (func(_784_v0 bool) func(_dafny.Int) bool {
							return func(_785_i1 _dafny.Int) bool {
								return _784_v0
							}
						})(_780_v0)
						_ = _init27
						var _element0_27 = _init27(_dafny.Zero)
						_ = _element0_27
						_nw146 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
						(_nw146).ArraySet1(_element0_27, 0)
						var _nativeLen0_27 = (_len0_27).Int()
						_ = _nativeLen0_27
						for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
							(_nw146).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
						}
					}
					_783_v2 = _nw146
					var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_783_v2), 0))
					_ = _index142
					(_783_v2).ArraySet1(true, (_index142).Int())
					var _786_v3 _dafny.MultiSet
					_ = _786_v3
					_786_v3 = _dafny.MultiSetOf(_780_v0, _780_v0)
					var _787_v4 _dafny.Map
					_ = _787_v4
					_787_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_786_v3).IsSubsetOf(_786_v3))
					var _788_v5 _dafny.Sequence
					_ = _788_v5
					_788_v5 = _dafny.UnicodeSeqOfUtf8Bytes("skejjki")
					var _789_v6 _dafny.Map
					_ = _789_v6
					_789_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_780_v0, _788_v5)
					var _790_v7 _dafny.MultiSet
					_ = _790_v7
					_790_v7 = _dafny.MultiSetOf(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_789_v6).Contains(!(!((p2).Dtor_cf1()))) {
							return (_789_v6).Get(!(!((p2).Dtor_cf1()))).(_dafny.Sequence)
						}
						return _788_v5
					})()).Cardinality()))
					var _791_v8 _dafny.Array
					_ = _791_v8
					var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
					_ = _nw147
					_791_v8 = _nw147
					var _792_v9 _dafny.Sequence
					_ = _792_v9
					_792_v9 = _dafny.SeqOf(_791_v8)
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_783_v2), 0))
					_ = _index143
					var _rhs117 _dafny.Int = p1
					_ = _rhs117
					var _rhs118 _dafny.Int = (func() _dafny.Int {
						if (_790_v7).Contains((p1).Times(p1)) {
							return (_790_v7).Multiplicity((p1).Times(p1))
						}
						return _dafny.IntOfUint32((_792_v9).Cardinality())
					})()
					_ = _rhs118
					var _rhs119 bool = _780_v0
					_ = _rhs119
					var _rhs120 _dafny.Map = _787_v4
					_ = _rhs120
					var _lhs86 *GlobalState = globalState
					_ = _lhs86
					var _lhs87 *GlobalState = globalState
					_ = _lhs87
					var _lhs88 _dafny.Array = _783_v2
					_ = _lhs88
					var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_783_v2), 0))
					_ = _lhs89
					_lhs86.F1 = _rhs117
					_lhs87.F1 = _rhs118
					(_lhs88).ArraySet1(_rhs119, (_lhs89).Int())
					_787_v4 = _rhs120
					var _793_v10 *C1
					_ = _793_v10
					var _nw148 *C1 = New_C1_()
					_ = _nw148
					_nw148.Ctor__((_783_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_783_v2), 0))).Int()).(bool), (_783_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_783_v2), 0))).Int()).(bool), _786_v3)
					_793_v10 = _nw148
					var _794_v11 _dafny.Map
					_ = _794_v11
					_794_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("eo"), !((_793_v10).F14()))
					_782_v1 = (func() _dafny.Set {
						if (func() bool {
							if (_794_v11).Contains(_dafny.UnicodeSeqOfUtf8Bytes("ei")) {
								return (_794_v11).Get(_dafny.UnicodeSeqOfUtf8Bytes("ei")).(bool)
							}
							return (_783_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_783_v2), 0))).Int()).(bool)
						})() {
							return _782_v1
						}
						return _782_v1
					})()
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _795_v12 _dafny.Sequence
		_ = _795_v12
		_795_v12 = _dafny.SeqOf(!(_780_v0))
		var _796_v13 _dafny.Map
		_ = _796_v13
		_796_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_780_v0, _795_v12)
		var _797_v14 _dafny.MultiSet
		_ = _797_v14
		_797_v14 = _dafny.MultiSetOf(p1)
		var _798_v15 _dafny.Sequence
		_ = _798_v15
		_798_v15 = _dafny.UnicodeSeqOfUtf8Bytes("ycsjd")
		var _799_v16 _dafny.Array
		_ = _799_v16
		var _nwElement0_33 _dafny.Int = p1
		_ = _nwElement0_33
		var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(14))
		_ = _nw149
		(_nw149).ArraySet1(_nwElement0_33, 0)
		(_nw149).ArraySet1((_dafny.Zero).Minus((_796_v13).Cardinality()), 1)
		(_nw149).ArraySet1((_797_v14).Cardinality(), 2)
		(_nw149).ArraySet1(p1, 3)
		(_nw149).ArraySet1(p1, 4)
		(_nw149).ArraySet1(p1, 5)
		(_nw149).ArraySet1((_dafny.Zero).Minus(p1), 6)
		(_nw149).ArraySet1(p1, 7)
		(_nw149).ArraySet1(p1, 8)
		(_nw149).ArraySet1(p1, 9)
		(_nw149).ArraySet1((_this).Fm9(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(256))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}((func(_800_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_801_i2 _dafny.Int) _dafny.CodePoint {
				return _800_p0
			}
		})(p0))), p1, globalState), 10)
		(_nw149).ArraySet1(_dafny.IntOfUint32((_798_v15).Cardinality()), 11)
		(_nw149).ArraySet1(p1, 12)
		(_nw149).ArraySet1((p1).Times(p1), 13)
		_799_v16 = _nw149
		var _802_v17 _dafny.MultiSet
		_ = _802_v17
		_802_v17 = _dafny.MultiSetOf(_780_v0)
		var _803_v18 *C1
		_ = _803_v18
		var _nw150 *C1 = New_C1_()
		_ = _nw150
		_nw150.Ctor__(_780_v0, _780_v0, _802_v17)
		_803_v18 = _nw150
		var _804_v19 _dafny.Map
		_ = _804_v19
		_804_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_780_v0, _803_v18)
		var _805_v20 _dafny.Map
		_ = _805_v20
		_805_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _804_v19)
		var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))
		_ = _index144
		(_799_v16).ArraySet1(Companion_Default___.SafeModuloInt((_805_v20).Cardinality(), p1), (_index144).Int())
		var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))
		_ = _index145
		(_799_v16).ArraySet1((p1).Times(p1), (_index145).Int())
		var _806_v21 _dafny.Array
		_ = _806_v21
		var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw151
		_806_v21 = _nw151
		var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_806_v21), 0))
		_ = _index146
		(_806_v21).ArraySet1(_795_v12, (_index146).Int())
		var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_806_v21), 0))
		_ = _index147
		(_806_v21).ArraySet1(_795_v12, (_index147).Int())
		var _hi8 _dafny.Int = p1
		_ = _hi8
		for _807_i3 := (_this).Fm9(_798_v15, (_this).Fm9(_798_v15, (_dafny.Zero).Minus((_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int)), globalState), globalState); _807_i3.Cmp(_hi8) < 0; _807_i3 = _807_i3.Plus(_dafny.One) {
			_780_v0 = !(!((false) || ((_803_v18).F15())))
			var _808_v22 _dafny.Sequence
			_ = _808_v22
			_808_v22 = _dafny.SeqOf(_807_i3, (_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int), (_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int))
			var _809_v23 _dafny.Sequence
			_ = _809_v23
			_809_v23 = _dafny.SeqOf(_808_v22, _808_v22, _808_v22)
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))
			_ = _index148
			var _rhs121 _dafny.Sequence = (_809_v23).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_809_v23).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs121
			var _rhs122 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32(((_806_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_806_v21), 0))).Int()).(_dafny.Sequence)).Cardinality())).Times(p1), (_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int))
			_ = _rhs122
			var _lhs90 _dafny.Array = _799_v16
			_ = _lhs90
			var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))
			_ = _lhs91
			_808_v22 = _rhs121
			(_lhs90).ArraySet1(_rhs122, (_lhs91).Int())
			var _810_v24 _dafny.Array
			_ = _810_v24
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_28
			var _nw152 _dafny.Array
			_ = _nw152
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw152 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) bool = (func(_811_i3 _dafny.Int, _812_v16 _dafny.Array) func(_dafny.Int) bool {
					return func(_813_i4 _dafny.Int) bool {
						return (_811_i3).Cmp((_812_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_812_v16), 0))).Int()).(_dafny.Int)) == 0
					}
				})(_807_i3, _799_v16)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw152 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw152).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw152).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_810_v24 = _nw152
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_810_v24), 0))
			_ = _index149
			(_810_v24).ArraySet1((_803_v18).F14(), (_index149).Int())
			var _814_v25 _dafny.Set
			_ = _814_v25
			_814_v25 = _dafny.SetOf((_dafny.Zero).Minus((_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int)))
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_810_v24), 0))
			_ = _index150
			(_810_v24).ArraySet1(Companion_Default___.Fm4(_814_v25, globalState), (_index150).Int())
			var _815_v26 _dafny.Map
			_ = _815_v26
			_815_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _798_v15)
			var _816_v27 D6
			_ = _816_v27
			_816_v27 = Companion_D6_.Create_DC17_((_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("hetkxtwq"), (_803_v18).F15())
			var _817_v28 _dafny.Map
			_ = _817_v28
			_817_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_803_v18).F15())
			var _818_v29 D6
			_ = _818_v29
			_818_v29 = Companion_D6_.Create_DC17_((_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int), (_816_v27).Dtor_cf39(), (func() bool {
				if (_817_v28).Contains((_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int)) {
					return (_817_v28).Get((_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int)).(bool)
				}
				return _780_v0
			})())
			(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_807_i3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_815_v26).Contains(p0) {
					return (_815_v26).Get(p0).(_dafny.Sequence)
				}
				return _798_v15
			})(), (_818_v29).Dtor_cf39())).Cardinality())))
		}
		var _ingredients0 = _dafny.NewBuilder()
		_ = _ingredients0
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_799_v16), 0))); ; {
			_guard_loop_3, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _819_i5 _dafny.Int
			_819_i5 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_819_i5).Sign() != -1) && ((_819_i5).Cmp(_dafny.ArrayLen((_799_v16), 0)) < 0)) {
				_ingredients0.Add(_dafny.TupleOf(_799_v16, (_819_i5).Int(), Companion_Default___.SafeModuloInt(_819_i5, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_799_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_799_v16), 0))).Int()).(_dafny.Int), (_803_v18).F15())).Merge((Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_803_v18).F15()))).Dtor_cf42())).Cardinality())))
			}
		}
		for _iter41 := _dafny.Iterate(_ingredients0); ; {
			_tup0, _ok41 := _iter41()
			if !_ok41 {
				break
			}
			(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
		}
		var _pat_let_tv18 = p1
		_ = _pat_let_tv18
		var _pat_let_tv19 = p1
		_ = _pat_let_tv19
		var _pat_let_tv20 = _797_v14
		_ = _pat_let_tv20
		var _pat_let_tv21 = _797_v14
		_ = _pat_let_tv21
		var _rhs123 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm19(p0, _780_v0, globalState), Companion_Default___.SafeDivisionInt(p1, p1))
		_ = _rhs123
		var _rhs124 _dafny.Array = (func() _dafny.Array {
			if (_803_v18).F14() {
				return _799_v16
			}
			return _799_v16
		})()
		_ = _rhs124
		var _rhs125 bool = !((_803_v18).F15()) || ((_803_v18).F14())
		_ = _rhs125
		var _rhs126 _dafny.MultiSet = func(_source9 D0) _dafny.MultiSet {
			if _source9.Is_DC1() {
				var _820___mcc_h0 bool = _source9.Get_().(D0_DC1).Cf1
				_ = _820___mcc_h0
				var _821___mcc_h1 _dafny.Int = _source9.Get_().(D0_DC1).Cf2
				_ = _821___mcc_h1
				var _822_cf2 _dafny.Int = _821___mcc_h1
				_ = _822_cf2
				var _823_cf1 bool = _820___mcc_h0
				_ = _823_cf1
				return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv18), _dafny.SeqOf(_pat_let_tv19, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(176))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}(func(_824_i6 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-747)
				}))).Cardinality()))))
			} else if _source9.Is_DC0() {
				var _825___mcc_h2 _dafny.CodePoint = _source9.Get_().(D0_DC0).Cf0
				_ = _825___mcc_h2
				var _826_cf0 _dafny.CodePoint = _825___mcc_h2
				_ = _826_cf0
				return _pat_let_tv20
			} else {
				var _827___mcc_h3 D0 = _source9.Get_().(D0_DC2).Cf3
				_ = _827___mcc_h3
				var _828_cf3 D0 = _827___mcc_h3
				_ = _828_cf3
				return _pat_let_tv21
			}
		}(Companion_Default___.Fm21(globalState))
		_ = _rhs126
		var _lhs92 *GlobalState = globalState
		_ = _lhs92
		_lhs92.F1 = _rhs123
		_799_v16 = _rhs124
		_780_v0 = _rhs125
		_797_v14 = _rhs126
	}
}
func (_this *C4) M2(globalState *GlobalState) (_dafny.Set, _dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _829_v0 bool
		_ = _829_v0
		_829_v0 = false
		var _830_i0 _dafny.Int
		_ = _830_i0
		_830_i0 = _dafny.Zero
		{
			for _829_v0 {
				{
					if (_830_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_830_i0 = (_830_i0).Plus(_dafny.One)
					var _831_v1 _dafny.Array
					_ = _831_v1
					var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
					_ = _nw153
					_831_v1 = _nw153
					var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))
					_ = _index151
					(_831_v1).ArraySet1(_829_v0, (_index151).Int())
					var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))
					_ = _index152
					(_831_v1).ArraySet1(_829_v0, (_index152).Int())
					var _832_v2 _dafny.Int
					_ = _832_v2
					_832_v2 = _dafny.IntOfInt64(809)
					var _833_v3 _dafny.Array
					_ = _833_v3
					var _nwElement0_34 _dafny.Int = _832_v2
					_ = _nwElement0_34
					var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(4))
					_ = _nw154
					(_nw154).ArraySet1(_nwElement0_34, 0)
					(_nw154).ArraySet1((_dafny.Zero).Minus(_832_v2), 1)
					(_nw154).ArraySet1(_dafny.IntOfInt64(315), 2)
					(_nw154).ArraySet1(_832_v2, 3)
					_833_v3 = _nw154
					var _834_v4 _dafny.Array
					_ = _834_v4
					var _nwElement0_35 _dafny.Array = _833_v3
					_ = _nwElement0_35
					var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(2))
					_ = _nw155
					(_nw155).ArraySet1(_nwElement0_35, 0)
					(_nw155).ArraySet1((func() _dafny.Array {
						if _829_v0 {
							return _833_v3
						}
						return _833_v3
					})(), 1)
					_834_v4 = _nw155
					_834_v4 = _834_v4
					var _835_v5 _dafny.Map
					_ = _835_v5
					_835_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _832_v2)
					var _836_v6 _dafny.Set
					_ = _836_v6
					_836_v6 = _dafny.SetOf(_832_v2, _832_v2, (((_835_v5).Update((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _dafny.IntOfInt64(642))).Update((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _832_v2)).Cardinality())
					if !(Companion_Default___.Fm4(_836_v6, globalState)) {
						var _837_v7 bool
						_ = _837_v7
						var _838_v8 _dafny.Array
						_ = _838_v8
						var _839_v9 _dafny.Int
						_ = _839_v9
						var _840_v10 _dafny.Int
						_ = _840_v10
						var _out24 bool
						_ = _out24
						var _out25 _dafny.Array
						_ = _out25
						var _out26 _dafny.Int
						_ = _out26
						var _out27 _dafny.Int
						_ = _out27
						_out24, _out25, _out26, _out27 = (_this).M3(globalState)
						_837_v7 = _out24
						_838_v8 = _out25
						_839_v9 = _out26
						_840_v10 = _out27
						var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
						_ = _nw156
						_833_v3 = _nw156
						var _841_v11 _dafny.CodePoint
						_ = _841_v11
						_841_v11 = _dafny.CodePoint('j')
						var _842_v12 D3
						_ = _842_v12
						_842_v12 = Companion_D3_.Create_DC8_(_841_v11, _829_v0, Companion_Default___.Fm4(_836_v6, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(662))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg40 _dafny.Int) interface{} {
								return coer40(arg40)
							}
						}((func(_843_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_844_i1 _dafny.Int) _dafny.CodePoint {
								return _843_v11
							}
						})(_841_v11))))
						_829_v0 = (_842_v12).Dtor_cf19()
						var _845_v13 _dafny.Map
						_ = _845_v13
						_845_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), !(_829_v0))
						_837_v7 = !(!((func() bool {
							if (_845_v13).Contains(_837_v7) {
								return (_845_v13).Get(_837_v7).(bool)
							}
							return _837_v7
						})()) || (_829_v0))
						_833_v3 = _833_v3
					} else {
						var _846_v14 _dafny.MultiSet
						_ = _846_v14
						_846_v14 = _dafny.MultiSetOf((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), (_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), (_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _829_v0, (_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool))
						var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))
						_ = _index153
						(_831_v1).ArraySet1(((_846_v14).Cardinality()).Cmp(_832_v2) < 0, (_index153).Int())
						var _847_v15 _dafny.Map
						_ = _847_v15
						_847_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_832_v2), (_832_v2).Plus(_832_v2))
						var _848_v16 _dafny.Sequence
						_ = _848_v16
						_848_v16 = _dafny.SeqOf(_835_v5, _835_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _dafny.IntOfInt64(127)))
						var _849_v17 _dafny.MultiSet
						_ = _849_v17
						_849_v17 = _dafny.MultiSetOf(_832_v2)
						_847_v15 = (_847_v15).Update((((_848_v16).Select((Companion_Default___.SafeIndex(_832_v2, _dafny.IntOfUint32((_848_v16).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Plus((func() _dafny.Int {
							if (_849_v17).Contains(_832_v2) {
								return (_849_v17).Multiplicity(_832_v2)
							}
							return _832_v2
						})()), (_832_v2).Plus(_832_v2))
						var _850_v18 _dafny.CodePoint
						_ = _850_v18
						_850_v18 = _dafny.CodePoint('k')
						_850_v18 = _dafny.CodePoint('i')
						var _851_v19 _dafny.Sequence
						_ = _851_v19
						_851_v19 = _dafny.UnicodeSeqOfUtf8Bytes("sby")
						var _852_v20 _dafny.Map
						_ = _852_v20
						_852_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_834_v4, _851_v19)
						var _853_v21 _dafny.Sequence
						_ = _853_v21
						_853_v21 = _dafny.SeqOf(_851_v19)
						var _854_v22 _dafny.Sequence
						_ = _854_v22
						_854_v22 = _dafny.SeqOf(_851_v19, _851_v19, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-274))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg41 _dafny.Int) interface{} {
								return coer41(arg41)
							}
						}((func(_855_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_856_i5 _dafny.Int) _dafny.CodePoint {
								return _855_v18
							}
						})(_850_v18))), (_853_v21).Select((Companion_Default___.SafeIndex(_832_v2, _dafny.IntOfUint32((_853_v21).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm2(_829_v0, _dafny.IntOfUint32((_851_v19).Cardinality()), _850_v18, _832_v2, globalState))
						var _857_v23 _dafny.Array
						_ = _857_v23
						var _nwElement0_36 _dafny.Sequence = _851_v19
						_ = _nwElement0_36
						var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(24))
						_ = _nw157
						(_nw157).ArraySet1(_nwElement0_36, 0)
						(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(757))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg42 _dafny.Int) interface{} {
								return coer42(arg42)
							}
						}((func(_858_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_859_i2 _dafny.Int) _dafny.CodePoint {
								return _858_v18
							}
						})(_850_v18))), _851_v19), 1)
						(_nw157).ArraySet1(_851_v19, 2)
						(_nw157).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("llcxt"), 3)
						(_nw157).ArraySet1(_851_v19, 4)
						(_nw157).ArraySet1((func() _dafny.Sequence {
							if (_852_v20).Contains(_834_v4) {
								return (_852_v20).Get(_834_v4).(_dafny.Sequence)
							}
							return _851_v19
						})(), 5)
						(_nw157).ArraySet1(_851_v19, 6)
						(_nw157).ArraySet1(_851_v19, 7)
						(_nw157).ArraySet1(_851_v19, 8)
						(_nw157).ArraySet1(_851_v19, 9)
						(_nw157).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(377))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}((func(_860_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_861_i3 _dafny.Int) _dafny.CodePoint {
								return _860_v18
							}
						})(_850_v18))), 10)
						(_nw157).ArraySet1(_dafny.Companion_Sequence_.Update(_851_v19, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.IntOfUint32((_851_v19).Cardinality()))).Uint32(), _850_v18), 11)
						(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_851_v19, Companion_Default___.Fm2((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _832_v2, _850_v18, _832_v2, globalState)), 12)
						(_nw157).ArraySet1(_851_v19, 13)
						(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_851_v19, _851_v19), 14)
						(_nw157).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(914))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}((func(_862_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_863_i4 _dafny.Int) _dafny.CodePoint {
								return _862_v18
							}
						})(_850_v18))), 15)
						(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_851_v19, _851_v19), 16)
						(_nw157).ArraySet1(_851_v19, 17)
						(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aqevycure"), _851_v19), 18)
						(_nw157).ArraySet1(Companion_Default___.Fm2(_829_v0, _832_v2, _850_v18, (_dafny.Zero).Minus(_dafny.IntOfUint32((_851_v19).Cardinality())), globalState), 19)
						(_nw157).ArraySet1(_dafny.Companion_Sequence_.Update(_851_v19, (Companion_Default___.SafeIndex(_832_v2, _dafny.IntOfUint32((_851_v19).Cardinality()))).Uint32(), _850_v18), 20)
						(_nw157).ArraySet1((_853_v21).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm19(_850_v18, true, globalState), _dafny.IntOfUint32((_853_v21).Cardinality()))).Uint32()).(_dafny.Sequence), 21)
						(_nw157).ArraySet1(_851_v19, 22)
						(_nw157).ArraySet1(_851_v19, 23)
						_857_v23 = _nw157
						var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_857_v23), 0))
						_ = _index154
						(_857_v23).ArraySet1(_dafny.Companion_Sequence_.Update(_851_v19, (Companion_Default___.SafeIndex(_832_v2, _dafny.IntOfUint32((_851_v19).Cardinality()))).Uint32(), _850_v18), (_index154).Int())
						var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_857_v23), 0))
						_ = _index155
						(_857_v23).ArraySet1(_851_v19, (_index155).Int())
						var _864_v24 D3
						_ = _864_v24
						_864_v24 = Companion_D3_.Create_DC8_(_850_v18, (_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), (_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), (_857_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_857_v23), 0))).Int()).(_dafny.Sequence))
						var _865_v25 D3
						_ = _865_v25
						_865_v25 = Companion_D3_.Create_DC9_(_864_v24)
						_865_v25 = _865_v25
					}
					var _866_v26 _dafny.Sequence
					_ = _866_v26
					_866_v26 = _dafny.UnicodeSeqOfUtf8Bytes("oxvklsjlo")
					var _867_v27 _dafny.CodePoint
					_ = _867_v27
					_867_v27 = _dafny.CodePoint('u')
					var _868_v28 _dafny.Map
					_ = _868_v28
					_868_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, _867_v27)
					var _869_v29 _dafny.Map
					_ = _869_v29
					_869_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), (_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool))
					var _870_v30 _dafny.Map
					_ = _870_v30
					_870_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-195), _829_v0)
					var _871_v31 _dafny.Map
					_ = _871_v31
					_871_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_832_v2, (_870_v30).Cardinality())
					var _872_v32 _dafny.Array
					_ = _872_v32
					var _nwElement0_37 _dafny.Sequence = _866_v26
					_ = _nwElement0_37
					var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(19))
					_ = _nw158
					(_nw158).ArraySet1(_nwElement0_37, 0)
					(_nw158).ArraySet1(_866_v26, 1)
					(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_866_v26, _866_v26), 2)
					(_nw158).ArraySet1(_866_v26, 3)
					(_nw158).ArraySet1(Companion_Default___.Fm2((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _832_v2, (func() _dafny.CodePoint {
						if (_868_v28).Contains((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool)) {
							return (_868_v28).Get((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool)).(_dafny.CodePoint)
						}
						return _867_v27
					})(), (_869_v29).Cardinality(), globalState), 4)
					(_nw158).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("le"), 5)
					(_nw158).ArraySet1(_866_v26, 6)
					(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_866_v26, (Companion_Default___.SafeIndex(_832_v2, _dafny.IntOfUint32((_866_v26).Cardinality()))).Uint32(), _867_v27), _866_v26), 7)
					(_nw158).ArraySet1(_866_v26, 8)
					(_nw158).ArraySet1(Companion_Default___.Fm2((_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), _832_v2, _dafny.CodePoint('i'), _832_v2, globalState), 9)
					(_nw158).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}((func(_873_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_874_i6 _dafny.Int) _dafny.CodePoint {
							return _873_v27
						}
					})(_867_v27))), 10)
					(_nw158).ArraySet1(_866_v26, 11)
					(_nw158).ArraySet1(_866_v26, 12)
					(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_866_v26, _866_v26), 13)
					(_nw158).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(579))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}((func(_875_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_876_i7 _dafny.Int) _dafny.CodePoint {
							return _875_v27
						}
					})(_867_v27))), 14)
					(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate((Companion_Default___.Fm23(_867_v27, (_871_v31).Cardinality(), _829_v0, (_831_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_831_v1), 0))).Int()).(bool), globalState)).Dtor_cf20(), _866_v26), 15)
					(_nw158).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rkcfltsn"), 16)
					(_nw158).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(818))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}((func(_877_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_878_i8 _dafny.Int) _dafny.CodePoint {
							return _877_v27
						}
					})(_867_v27))), 17)
					(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_866_v26, _866_v26), 18)
					_872_v32 = _nw158
					var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_872_v32), 0))
					_ = _index156
					(_872_v32).ArraySet1(_866_v26, (_index156).Int())
					var _879_v33 _dafny.Sequence
					_ = _879_v33
					_879_v33 = _dafny.SeqOf(_866_v26)
					var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_872_v32), 0))
					_ = _index157
					(_872_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_879_v33).Select((Companion_Default___.SafeIndex(_832_v2, _dafny.IntOfUint32((_879_v33).Cardinality()))).Uint32()).(_dafny.Sequence), _866_v26), (_index157).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _880_i9 _dafny.Int
		_ = _880_i9
		_880_i9 = _dafny.Zero
		{
			for !(_829_v0) {
				{
					if (_880_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_880_i9 = (_880_i9).Plus(_dafny.One)
					var _881_v34 _dafny.Int
					_ = _881_v34
					_881_v34 = _dafny.IntOfInt64(522)
					var _882_v35 _dafny.MultiSet
					_ = _882_v35
					_882_v35 = _dafny.MultiSetOf(_829_v0)
					var _883_v36 *C1
					_ = _883_v36
					var _nw159 *C1 = New_C1_()
					_ = _nw159
					_nw159.Ctor__(_829_v0, _829_v0, _882_v35)
					_883_v36 = _nw159
					var _884_v37 _dafny.Map
					_ = _884_v37
					_884_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, _883_v36)
					(globalState).F1 = Companion_Default___.SafeDivisionInt(_881_v34, (func() _dafny.Int {
						if _829_v0 {
							return (_884_v37).Cardinality()
						}
						return _881_v34
					})())
					var _885_v38 _dafny.MultiSet
					_ = _885_v38
					_885_v38 = _dafny.MultiSetOf(_881_v34)
					var _886_v39 _dafny.MultiSet
					_ = _886_v39
					_886_v39 = _885_v38
					var _887_v40 _dafny.Map
					_ = _887_v40
					_887_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_886_v39, (_881_v34).Cmp((_885_v38).Cardinality()) != 0)
					var _888_v41 _dafny.Set
					_ = _888_v41
					_888_v41 = _dafny.SetOf(_881_v34, (_882_v35).Cardinality())
					var _889_v42 _dafny.Map
					_ = _889_v42
					_889_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_881_v34, Companion_Default___.Fm4(_888_v41, globalState))
					var _890_v43 _dafny.Map
					_ = _890_v43
					_890_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(929))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg48 _dafny.Int) interface{} {
							return coer48(arg48)
						}
					}(func(_891_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					})), _881_v34)
					_887_v40 = (_887_v40).Update(_885_v38, (func() bool {
						if (_889_v42).Contains((_890_v43).Cardinality()) {
							return (_889_v42).Get((_890_v43).Cardinality()).(bool)
						}
						return _829_v0
					})())
					_829_v0 = !(_829_v0)
					var _892_v44 *C0
					_ = _892_v44
					var _nw160 *C0 = New_C0_()
					_ = _nw160
					_nw160.Ctor__()
					_892_v44 = _nw160
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _893_v45 _dafny.Sequence
		_ = _893_v45
		_893_v45 = _dafny.UnicodeSeqOfUtf8Bytes("evwq")
		var _894_v46 _dafny.MultiSet
		_ = _894_v46
		_894_v46 = _dafny.MultiSetOf(_893_v45)
		var _895_v47 _dafny.Int
		_ = _895_v47
		_895_v47 = _dafny.IntOfInt64(502)
		var _896_v48 _dafny.Sequence
		_ = _896_v48
		_896_v48 = _dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_894_v46).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_893_v45).Cardinality())))), _dafny.IntOfInt64(485), Companion_Default___.Fm19(Companion_Default___.Fm1(globalState), _829_v0, globalState), _895_v47)
		_896_v48 = _dafny.Companion_Sequence_.Concatenate(_896_v48, _896_v48)
		var _897_i11 _dafny.Int
		_ = _897_i11
		_897_i11 = _dafny.Zero
		{
			for true {
				{
					if (_897_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_897_i11 = (_897_i11).Plus(_dafny.One)
					var _898_v49 _dafny.Map
					_ = _898_v49
					_898_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v47, _829_v0)
					var _899_v50 _dafny.Sequence
					_ = _899_v50
					_899_v50 = _dafny.SeqOf(_829_v0)
					var _900_v51 _dafny.MultiSet
					_ = _900_v51
					_900_v51 = _dafny.MultiSetOf(_829_v0, _829_v0)
					var _901_v52 _dafny.Map
					_ = _901_v52
					_901_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, _900_v51)
					var _902_v53 _dafny.Map
					_ = _902_v53
					_902_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v0, Companion_Default___.Fm4(_dafny.SetOf(_dafny.IntOfInt64(-482)), globalState))
					var _903_v54 _dafny.MultiSet
					_ = _903_v54
					_903_v54 = _dafny.MultiSetOf(_895_v47, _dafny.IntOfUint32((_dafny.SeqOf(_829_v0)).Cardinality()))
					var _904_v55 _dafny.Set
					_ = _904_v55
					_904_v55 = _dafny.SetOf(_895_v47, _895_v47, _895_v47, _895_v47, (_903_v54).Cardinality())
					var _905_v56 _dafny.Map
					_ = _905_v56
					_905_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_902_v53, _dafny.MultiSetOf((_904_v55).Cardinality(), _895_v47))
					var _906_v57 _dafny.Map
					_ = _906_v57
					_906_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_905_v56).Cardinality(), (_898_v49).Cardinality())
					var _907_v58 _dafny.Array
					_ = _907_v58
					var _nwElement0_38 bool = (func() bool {
						if (_898_v49).Contains(_895_v47) {
							return (_898_v49).Get(_895_v47).(bool)
						}
						return _829_v0
					})()
					_ = _nwElement0_38
					var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(19))
					_ = _nw161
					(_nw161).ArraySet1(_nwElement0_38, 0)
					(_nw161).ArraySet1(_829_v0, 1)
					(_nw161).ArraySet1(_829_v0, 2)
					(_nw161).ArraySet1(!(_dafny.Companion_Sequence_.Contains(_899_v50, _829_v0)), 3)
					(_nw161).ArraySet1(_829_v0, 4)
					(_nw161).ArraySet1(_829_v0, 5)
					(_nw161).ArraySet1((Companion_Default___.Fm24(_dafny.IntOfUint32((_899_v50).Cardinality()), _829_v0, _898_v49, _895_v47, globalState)).IsSubsetOf((func() _dafny.MultiSet {
						if (_901_v52).Contains(_829_v0) {
							return (_901_v52).Get(_829_v0).(_dafny.MultiSet)
						}
						return _900_v51
					})()), 6)
					(_nw161).ArraySet1(false, 7)
					(_nw161).ArraySet1(_829_v0, 8)
					(_nw161).ArraySet1(_829_v0, 9)
					(_nw161).ArraySet1(true, 10)
					(_nw161).ArraySet1(_829_v0, 11)
					(_nw161).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(863), _895_v47)).Equals(_906_v57), 12)
					(_nw161).ArraySet1(_829_v0, 13)
					(_nw161).ArraySet1(!(true) || (_829_v0), 14)
					(_nw161).ArraySet1(_829_v0, 15)
					(_nw161).ArraySet1((_dafny.IntOfInt64(206)).Cmp(_895_v47) < 0, 16)
					(_nw161).ArraySet1(_829_v0, 17)
					(_nw161).ArraySet1(_829_v0, 18)
					_907_v58 = _nw161
					var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_907_v58), 0))
					_ = _index158
					(_907_v58).ArraySet1(!((_895_v47).Cmp((_902_v53).Cardinality()) > 0), (_index158).Int())
					var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_907_v58), 0))
					_ = _index159
					(_907_v58).ArraySet1((func() bool {
						if _829_v0 {
							return _829_v0
						}
						return _829_v0
					})(), (_index159).Int())
					var _908_v59 _dafny.CodePoint
					_ = _908_v59
					_908_v59 = _dafny.CodePoint('b')
					var _909_v60 D3
					_ = _909_v60
					_909_v60 = Companion_D3_.Create_DC8_(_908_v59, _829_v0, _829_v0, _893_v45)
					var _910_v61 _dafny.Array
					_ = _910_v61
					var _nwElement0_39 _dafny.Sequence = _893_v45
					_ = _nwElement0_39
					var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(21))
					_ = _nw162
					(_nw162).ArraySet1(_nwElement0_39, 0)
					(_nw162).ArraySet1(_893_v45, 1)
					(_nw162).ArraySet1(_893_v45, 2)
					(_nw162).ArraySet1(_893_v45, 3)
					(_nw162).ArraySet1(_893_v45, 4)
					(_nw162).ArraySet1(_893_v45, 5)
					(_nw162).ArraySet1(_893_v45, 6)
					(_nw162).ArraySet1(_893_v45, 7)
					(_nw162).ArraySet1(_893_v45, 8)
					(_nw162).ArraySet1(Companion_Default___.Fm2(_829_v0, _895_v47, _908_v59, _895_v47, globalState), 9)
					(_nw162).ArraySet1(_893_v45, 10)
					(_nw162).ArraySet1(_dafny.Companion_Sequence_.Update(_893_v45, (Companion_Default___.SafeIndex(_895_v47, _dafny.IntOfUint32((_893_v45).Cardinality()))).Uint32(), _908_v59), 11)
					(_nw162).ArraySet1(_893_v45, 12)
					(_nw162).ArraySet1(_893_v45, 13)
					(_nw162).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(308))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}((func(_911_v59 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_912_i12 _dafny.Int) _dafny.CodePoint {
							return _911_v59
						}
					})(_908_v59))), 14)
					(_nw162).ArraySet1((_909_v60).Dtor_cf20(), 15)
					(_nw162).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fbkhxmvx"), 16)
					(_nw162).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ooxi"), 17)
					(_nw162).ArraySet1(_893_v45, 18)
					(_nw162).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vvrnb"), 19)
					(_nw162).ArraySet1(_893_v45, 20)
					_910_v61 = _nw162
					var _913_v62 _dafny.Map
					_ = _913_v62
					_913_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v61, (_899_v50).Select((Companion_Default___.SafeIndex(_895_v47, _dafny.IntOfUint32((_899_v50).Cardinality()))).Uint32()).(bool))
					_829_v0 = (func() bool {
						if (_913_v62).Contains((func() _dafny.Array {
							if _829_v0 {
								return _910_v61
							}
							return _910_v61
						})()) {
							return (_913_v62).Get((func() _dafny.Array {
								if _829_v0 {
									return _910_v61
								}
								return _910_v61
							})()).(bool)
						}
						return _829_v0
					})()
					var _914_v63 *C2
					_ = _914_v63
					var _nw163 *C2 = New_C2_()
					_ = _nw163
					_nw163.Ctor__(_829_v0, (_902_v53).Cardinality(), _900_v51)
					_914_v63 = _nw163
					var _915_v64 *C0
					_ = _915_v64
					var _nw164 *C0 = New_C0_()
					_ = _nw164
					_nw164.Ctor__()
					_915_v64 = _nw164
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _916_v65 _dafny.Array
		_ = _916_v65
		var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
		_ = _nw165
		_916_v65 = _nw165
		for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_916_v65), 0))); ; {
			_guard_loop_4, _ok42 := _iter42()
			if !_ok42 {
				break
			}
			var _917_i13 _dafny.Int
			_917_i13 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_917_i13).Sign() != -1) && ((_917_i13).Cmp(_dafny.ArrayLen((_916_v65), 0)) < 0)) {
				(_916_v65).ArraySet1(_893_v45, (_917_i13).Int())
			}
		}
		var _918_v66 _dafny.MultiSet
		_ = _918_v66
		_918_v66 = _dafny.MultiSetOf(_829_v0, false, _829_v0, _829_v0, _829_v0)
		var _919_v67 _dafny.CodePoint
		_ = _919_v67
		_919_v67 = _dafny.CodePoint('b')
		var _920_v68 _dafny.Set
		_ = _920_v68
		_920_v68 = _dafny.SetOf(_829_v0)
		var _921_v69 _dafny.Array
		_ = _921_v69
		var _nwElement0_40 _dafny.Int = _895_v47
		_ = _nwElement0_40
		var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(28))
		_ = _nw166
		(_nw166).ArraySet1(_nwElement0_40, 0)
		(_nw166).ArraySet1(_895_v47, 1)
		(_nw166).ArraySet1(_895_v47, 2)
		(_nw166).ArraySet1((_918_v66).Cardinality(), 3)
		(_nw166).ArraySet1(_895_v47, 4)
		(_nw166).ArraySet1(_895_v47, 5)
		(_nw166).ArraySet1(_895_v47, 6)
		(_nw166).ArraySet1(_895_v47, 7)
		(_nw166).ArraySet1((func() _dafny.Int {
			if (_918_v66).Contains(_829_v0) {
				return (_918_v66).Multiplicity(_829_v0)
			}
			return _895_v47
		})(), 8)
		(_nw166).ArraySet1(_895_v47, 9)
		(_nw166).ArraySet1(Companion_Default___.Fm19(_919_v67, _829_v0, globalState), 10)
		(_nw166).ArraySet1(_895_v47, 11)
		(_nw166).ArraySet1(_895_v47, 12)
		(_nw166).ArraySet1(_895_v47, 13)
		(_nw166).ArraySet1((_this).Fm9(_893_v45, (_dafny.Zero).Minus(_895_v47), globalState), 14)
		(_nw166).ArraySet1(_dafny.IntOfInt64(240), 15)
		(_nw166).ArraySet1(_895_v47, 16)
		(_nw166).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_896_v48).Select((Companion_Default___.SafeIndex(_895_v47, _dafny.IntOfUint32((_896_v48).Cardinality()))).Uint32()).(_dafny.Int))), 17)
		(_nw166).ArraySet1(_895_v47, 18)
		(_nw166).ArraySet1(_895_v47, 19)
		(_nw166).ArraySet1(_895_v47, 20)
		(_nw166).ArraySet1(_895_v47, 21)
		(_nw166).ArraySet1(_895_v47, 22)
		(_nw166).ArraySet1(_895_v47, 23)
		(_nw166).ArraySet1(_895_v47, 24)
		(_nw166).ArraySet1((_920_v68).Cardinality(), 25)
		(_nw166).ArraySet1(_895_v47, 26)
		(_nw166).ArraySet1(_895_v47, 27)
		_921_v69 = _nw166
		var _922_v70 _dafny.Array
		_ = _922_v70
		var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
		_ = _nw167
		_922_v70 = _nw167
		var _source10 D5 = Companion_D5_.Create_DC13_(_895_v47, _921_v69, _922_v70, (_dafny.IntOfInt64(200)).Minus(_dafny.IntOfUint32((_893_v45).Cardinality())))
		_ = _source10
		if _source10.Is_DC13() {
			var _923___mcc_h0 _dafny.Int = _source10.Get_().(D5_DC13).Cf27
			_ = _923___mcc_h0
			var _924___mcc_h1 _dafny.Array = _source10.Get_().(D5_DC13).Cf28
			_ = _924___mcc_h1
			var _925___mcc_h2 _dafny.Array = _source10.Get_().(D5_DC13).Cf29
			_ = _925___mcc_h2
			var _926___mcc_h3 _dafny.Int = _source10.Get_().(D5_DC13).Cf30
			_ = _926___mcc_h3
			var _927_cf30 _dafny.Int = _926___mcc_h3
			_ = _927_cf30
			var _928_cf29 _dafny.Array = _925___mcc_h2
			_ = _928_cf29
			var _929_cf28 _dafny.Array = _924___mcc_h1
			_ = _929_cf28
			var _930_cf27 _dafny.Int = _923___mcc_h0
			_ = _930_cf27
			var _931_v71 D9
			_ = _931_v71
			_931_v71 = Companion_D9_.Create_DC21_(_dafny.SeqOf(_829_v0))
			var _932_v72 _dafny.Sequence
			_ = _932_v72
			_932_v72 = _dafny.SeqOf(_829_v0)
			var _933_v73 bool
			_ = _933_v73
			var _934_v74 _dafny.MultiSet
			_ = _934_v74
			var _935_v75 _dafny.Set
			_ = _935_v75
			var _936_v76 _dafny.Int
			_ = _936_v76
			var _out28 bool
			_ = _out28
			var _out29 _dafny.MultiSet
			_ = _out29
			var _out30 _dafny.Set
			_ = _out30
			var _out31 _dafny.Int
			_ = _out31
			_out28, _out29, _out30, _out31 = Companion_Default___.M0((_931_v71).Dtor_cf48(), _932_v72, globalState)
			_933_v73 = _out28
			_934_v74 = _out29
			_935_v75 = _out30
			_936_v76 = _out31
			var _937_v77 _dafny.Map
			_ = _937_v77
			_937_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_930_cf27, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("shpmce"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-490), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("shpmce")).Cardinality()))).Uint32(), _dafny.CodePoint('q')))
			var _938_v78 D9
			_ = _938_v78
			_938_v78 = Companion_D9_.Create_DC24_(_895_v47)
			_937_v77 = Companion_Default___.Fm25(_938_v78, globalState)
			_933_v73 = !((((_dafny.Zero).Minus(_dafny.IntOfUint32((_893_v45).Cardinality()))).Plus((_dafny.Zero).Minus(_930_cf27))).Cmp(_930_cf27) == 0)
			_927_cf30 = _936_v76
		} else if _source10.Is_DC14() {
			var _939___mcc_h4 _dafny.Array = _source10.Get_().(D5_DC14).Cf31
			_ = _939___mcc_h4
			var _940___mcc_h5 _dafny.Int = _source10.Get_().(D5_DC14).Cf32
			_ = _940___mcc_h5
			var _941___mcc_h6 _dafny.Int = _source10.Get_().(D5_DC14).Cf33
			_ = _941___mcc_h6
			var _942___mcc_h7 _dafny.Map = _source10.Get_().(D5_DC14).Cf34
			_ = _942___mcc_h7
			var _943___mcc_h8 _dafny.Set = _source10.Get_().(D5_DC14).Cf35
			_ = _943___mcc_h8
			var _944_cf35 _dafny.Set = _943___mcc_h8
			_ = _944_cf35
			var _945_cf34 _dafny.Map = _942___mcc_h7
			_ = _945_cf34
			var _946_cf33 _dafny.Int = _941___mcc_h6
			_ = _946_cf33
			var _947_cf32 _dafny.Int = _940___mcc_h5
			_ = _947_cf32
			var _948_cf31 _dafny.Array = _939___mcc_h4
			_ = _948_cf31
			_896_v48 = _dafny.SeqOf((_895_v47).Plus(_895_v47), _895_v47)
			var _949_v79 T0
			_ = _949_v79
			var _nw168 *C1 = New_C1_()
			_ = _nw168
			_nw168.Ctor__(_829_v0, _829_v0, _dafny.MultiSetOf(_829_v0))
			_949_v79 = _nw168
			var _950_v80 _dafny.Array
			_ = _950_v80
			var _nwElement0_41 T0 = _949_v79
			_ = _nwElement0_41
			var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(15))
			_ = _nw169
			(_nw169).ArraySet1(_nwElement0_41, 0)
			(_nw169).ArraySet1(_949_v79, 1)
			(_nw169).ArraySet1(_949_v79, 2)
			(_nw169).ArraySet1(_949_v79, 3)
			(_nw169).ArraySet1(_949_v79, 4)
			(_nw169).ArraySet1(_949_v79, 5)
			(_nw169).ArraySet1((func() T0 {
				if !(false) {
					return _949_v79
				}
				return _949_v79
			})(), 6)
			(_nw169).ArraySet1(_949_v79, 7)
			(_nw169).ArraySet1(_949_v79, 8)
			(_nw169).ArraySet1(_949_v79, 9)
			(_nw169).ArraySet1(_949_v79, 10)
			(_nw169).ArraySet1(_949_v79, 11)
			(_nw169).ArraySet1(_949_v79, 12)
			(_nw169).ArraySet1(_949_v79, 13)
			(_nw169).ArraySet1(_949_v79, 14)
			_950_v80 = _nw169
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_950_v80), 0))
			_ = _index160
			(_950_v80).ArraySet1(_949_v79, (_index160).Int())
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_950_v80), 0))
			_ = _index161
			var _nw170 *C1 = New_C1_()
			_ = _nw170
			_nw170.Ctor__(_829_v0, _dafny.Companion_Sequence_.IsPrefixOf(_893_v45, _893_v45), _918_v66)
			(_950_v80).ArraySet1(_nw170, (_index161).Int())
			var _951_v81 _dafny.Set
			_ = _951_v81
			_951_v81 = _dafny.SetOf(_895_v47, _947_cf32)
			var _952_v82 _dafny.Map
			_ = _952_v82
			_952_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_951_v81).Cardinality()), _829_v0)
			var _953_v83 _dafny.MultiSet
			_ = _953_v83
			_953_v83 = _dafny.MultiSetOf(_895_v47, _946_cf33, _946_cf33, (_952_v82).Cardinality())
			var _954_v84 _dafny.Sequence
			_ = _954_v84
			_954_v84 = _dafny.SeqOf((_953_v83).Difference(_dafny.MultiSetOf(_895_v47)), (_dafny.MultiSetOf(_946_cf33)).Difference(_953_v83), (_953_v83).Union(_953_v83), _dafny.MultiSetFromSeq(_896_v48))
			var _955_v85 _dafny.Sequence
			_ = _955_v85
			_955_v85 = _dafny.SeqOf(_954_v84, _954_v84, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(390))).Uint32(), func(coer50 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_956_v83 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_957_i14 _dafny.Int) _dafny.MultiSet {
					return _956_v83
				}
			})(_953_v83))), _954_v84, _dafny.SeqOf(_953_v83))
			var _958_v86 _dafny.Map
			_ = _958_v86
			_958_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_953_v83).Cardinality(), _946_cf33)
			_954_v84 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_954_v84, _954_v84), (_955_v85).Select((Companion_Default___.SafeIndex((_958_v86).Cardinality(), _dafny.IntOfUint32((_955_v85).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_921_v69), 0))
			_ = _index162
			(_921_v69).ArraySet1(Companion_Default___.Fm19(Companion_Default___.Fm1(globalState), _829_v0, globalState), (_index162).Int())
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_921_v69), 0))
			_ = _index163
			(_921_v69).ArraySet1(Companion_Default___.Fm19(_919_v67, !(Companion_Default___.Fm4(_951_v81, globalState)), globalState), (_index163).Int())
		} else if _source10.Is_DC12() {
			var _959___mcc_h9 _dafny.Set = _source10.Get_().(D5_DC12).Cf26
			_ = _959___mcc_h9
			var _960_cf26 _dafny.Set = _959___mcc_h9
			_ = _960_cf26
			var _961_v87 _dafny.Map
			_ = _961_v87
			_961_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v47, _895_v47)
			var _rhs127 bool = !(_829_v0)
			_ = _rhs127
			var _rhs128 _dafny.Map = _961_v87
			_ = _rhs128
			var _rhs129 _dafny.CodePoint = _919_v67
			_ = _rhs129
			_829_v0 = _rhs127
			_961_v87 = _rhs128
			_919_v67 = _rhs129
			var _rhs130 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_895_v47, _895_v47))
			_ = _rhs130
			var _rhs131 _dafny.Int = (_895_v47).Minus(_dafny.IntOfUint32((_893_v45).Cardinality()))
			_ = _rhs131
			var _rhs132 _dafny.Int = _dafny.IntOfInt64(561)
			_ = _rhs132
			var _lhs93 *GlobalState = globalState
			_ = _lhs93
			r1 = _rhs130
			_lhs93.F1 = _rhs131
			r1 = _rhs132
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_916_v65), 0))
			_ = _index164
			(_916_v65).ArraySet1(_dafny.Companion_Sequence_.Update(_893_v45, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_895_v47), _dafny.IntOfUint32((_893_v45).Cardinality()))).Uint32(), _919_v67), (_index164).Int())
			var _962_v88 _dafny.Array
			_ = _962_v88
			var _nw171 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw171
			_962_v88 = _nw171
			var _963_v89 _dafny.Sequence
			_ = _963_v89
			_963_v89 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(440))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}(func(_964_i15 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			})))
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_916_v65), 0))
			_ = _index165
			var _rhs133 _dafny.Array = _962_v88
			_ = _rhs133
			var _rhs134 _dafny.Sequence = (_963_v89).Select((Companion_Default___.SafeIndex(_895_v47, _dafny.IntOfUint32((_963_v89).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs134
			var _rhs135 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_893_v45, _dafny.SeqOf(_919_v67))
			_ = _rhs135
			var _rhs136 _dafny.Sequence = _893_v45
			_ = _rhs136
			var _rhs137 bool = !(!(_829_v0))
			_ = _rhs137
			var _lhs94 *GlobalState = globalState
			_ = _lhs94
			var _lhs95 _dafny.Array = _916_v65
			_ = _lhs95
			var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_916_v65), 0))
			_ = _lhs96
			_lhs94.F5 = _rhs133
			(_lhs95).ArraySet1(_rhs134, (_lhs96).Int())
			r2 = _rhs135
			_893_v45 = _rhs136
			_829_v0 = _rhs137
			var _965_v90 _dafny.Array
			_ = _965_v90
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_29
			var _nw172 _dafny.Array
			_ = _nw172
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw172 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.Sequence = (func(_966_v0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_967_i16 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_966_v0, _966_v0, _966_v0, !(true))
					}
				})(_829_v0)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw172 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw172).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw172).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_965_v90 = _nw172
			var _968_v91 _dafny.Sequence
			_ = _968_v91
			_968_v91 = _dafny.SeqOf(!(true))
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_965_v90), 0))
			_ = _index166
			(_965_v90).ArraySet1(_968_v91, (_index166).Int())
			var _969_v92 _dafny.Array
			_ = _969_v92
			var _nw173 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw173
			_969_v92 = _nw173
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_969_v92), 0))
			_ = _index167
			(_969_v92).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_896_v48, _829_v0)).Update(Companion_Default___.Fm20(_895_v47, _895_v47, globalState), _829_v0), (_index167).Int())
			var _970_v93 _dafny.Set
			_ = _970_v93
			_970_v93 = _dafny.SetOf(_895_v47, _895_v47)
			var _971_v94 _dafny.MultiSet
			_ = _971_v94
			_971_v94 = _dafny.MultiSetOf((_dafny.SetOf(_829_v0)).Cardinality(), _895_v47, _895_v47, (_this).Fm9((_916_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_916_v65), 0))).Int()).(_dafny.Sequence), _dafny.IntOfInt64(169), globalState))
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_965_v90), 0))
			_ = _index168
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_969_v92), 0))
			_ = _index169
			var _rhs138 _dafny.Sequence = _968_v91
			_ = _rhs138
			var _rhs139 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_970_v93).Cardinality()), ((_971_v94).Cardinality()).Cmp(_dafny.IntOfInt64(424)) <= 0)
			_ = _rhs139
			var _rhs140 bool = !(_829_v0)
			_ = _rhs140
			var _lhs97 _dafny.Array = _965_v90
			_ = _lhs97
			var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_965_v90), 0))
			_ = _lhs98
			var _lhs99 _dafny.Array = _969_v92
			_ = _lhs99
			var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_969_v92), 0))
			_ = _lhs100
			(_lhs97).ArraySet1(_rhs138, (_lhs98).Int())
			(_lhs99).ArraySet1(_rhs139, (_lhs100).Int())
			_829_v0 = _rhs140
		} else {
			var _972___mcc_h10 D5 = _source10.Get_().(D5_DC15).Cf36
			_ = _972___mcc_h10
			var _973_cf36 D5 = _972___mcc_h10
			_ = _973_cf36
			var _974_v95 *C2
			_ = _974_v95
			var _nw174 *C2 = New_C2_()
			_ = _nw174
			_nw174.Ctor__(_829_v0, _895_v47, _dafny.MultiSetOf(false))
			_974_v95 = _nw174
			_974_v95 = _974_v95
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_921_v69), 0))
			_ = _index170
			(_921_v69).ArraySet1(Companion_Default___.SafeDivisionInt(_895_v47, _895_v47), (_index170).Int())
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_921_v69), 0))
			_ = _index171
			var _rhs141 _dafny.Int = _895_v47
			_ = _rhs141
			var _rhs142 _dafny.Int = (_dafny.Zero).Minus((_974_v95).F13())
			_ = _rhs142
			var _rhs143 bool = (!(false)) == ((_895_v47).Cmp(_dafny.IntOfInt64(-387)) > 0)
			_ = _rhs143
			var _rhs144 bool = !(_829_v0)
			_ = _rhs144
			var _lhs101 _dafny.Array = _921_v69
			_ = _lhs101
			var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_921_v69), 0))
			_ = _lhs102
			var _lhs103 *C2 = _974_v95
			_ = _lhs103
			(_lhs101).ArraySet1(_rhs141, (_lhs102).Int())
			r1 = _rhs142
			_829_v0 = _rhs143
			_lhs103.F12 = _rhs144
			var _975_v96 _dafny.Sequence
			_ = _975_v96
			_975_v96 = _dafny.SeqOf(_829_v0)
			var _976_v97 _dafny.Set
			_ = _976_v97
			_976_v97 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_975_v96, (Companion_Default___.SafeIndex((_921_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_921_v69), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_975_v96).Cardinality()))).Uint32(), _974_v95.F12)).Cardinality()), (_921_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_921_v69), 0))).Int()).(_dafny.Int))
			var _977_v98 _dafny.Array
			_ = _977_v98
			var _nwElement0_42 bool = _829_v0
			_ = _nwElement0_42
			var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(11))
			_ = _nw175
			(_nw175).ArraySet1(_nwElement0_42, 0)
			(_nw175).ArraySet1(_974_v95.F12, 1)
			(_nw175).ArraySet1(Companion_Default___.Fm4(_976_v97, globalState), 2)
			(_nw175).ArraySet1(!((_974_v95.F12) && (true)), 3)
			(_nw175).ArraySet1(_974_v95.F12, 4)
			(_nw175).ArraySet1(_974_v95.F12, 5)
			(_nw175).ArraySet1((func() bool {
				if _974_v95.F12 {
					return _974_v95.F12
				}
				return true
			})(), 6)
			(_nw175).ArraySet1(!(_829_v0), 7)
			(_nw175).ArraySet1(_974_v95.F12, 8)
			(_nw175).ArraySet1(!((func() bool {
				if _974_v95.F12 {
					return (_975_v96).Select((Companion_Default___.SafeIndex((_974_v95).F13(), _dafny.IntOfUint32((_975_v96).Cardinality()))).Uint32()).(bool)
				}
				return _974_v95.F12
			})()), 9)
			(_nw175).ArraySet1(((_921_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_921_v69), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(575)) <= 0, 10)
			_977_v98 = _nw175
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_977_v98), 0))
			_ = _index172
			(_977_v98).ArraySet1(_829_v0, (_index172).Int())
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_977_v98), 0))
			_ = _index173
			(_977_v98).ArraySet1(((_920_v68).Union(_920_v68)).IsDisjointFrom(_920_v68), (_index173).Int())
			(_974_v95).F12 = false
		}
		var _978_v99 _dafny.Sequence
		_ = _978_v99
		_978_v99 = _dafny.SeqOf(_829_v0, _829_v0)
		var _979_v100 _dafny.Map
		_ = _979_v100
		_979_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v47, _895_v47)
		var _980_v101 _dafny.Set
		_ = _980_v101
		_980_v101 = _dafny.SetOf(_895_v47, _dafny.IntOfUint32((_896_v48).Cardinality()), _895_v47)
		var _981_v102 _dafny.Set
		_ = _981_v102
		_981_v102 = _dafny.SetOf(_895_v47, Companion_Default___.Fm19(_919_v67, (_978_v99).Select((Companion_Default___.SafeIndex(_895_v47, _dafny.IntOfUint32((_978_v99).Cardinality()))).Uint32()).(bool), globalState), (_895_v47).Minus(_895_v47), (func() _dafny.Int {
			if (_979_v100).Contains((_this).Fm9(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}((func(_982_v67 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_983_i17 _dafny.Int) _dafny.CodePoint {
					return _982_v67
				}
			})(_919_v67))), (Companion_D6_.Create_DC17_(_895_v47, _893_v45, Companion_Default___.Fm4(_980_v101, globalState))).Dtor_cf38(), globalState)) {
				return (_979_v100).Get((_this).Fm9(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_984_v67 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_985_i17 _dafny.Int) _dafny.CodePoint {
						return _984_v67
					}
				})(_919_v67))), (Companion_D6_.Create_DC17_(_895_v47, _893_v45, Companion_Default___.Fm4(_980_v101, globalState))).Dtor_cf38(), globalState)).(_dafny.Int)
			}
			return _dafny.IntOfInt64(153)
		})())
		r0 = _981_v102
		r1 = _895_v47
		r2 = _dafny.Companion_Sequence_.Concatenate(_893_v45, Companion_Default___.Fm2(_829_v0, _895_v47, _919_v67, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v47, _829_v0)).Cardinality(), globalState))
		return r0, r1, r2
	}
}
func (_this *C4) M3(globalState *GlobalState) (bool, _dafny.Array, _dafny.Int, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		if false {
			var _986_v0 _dafny.Sequence
			_ = _986_v0
			_986_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dtaw")
			var _987_v1 _dafny.Set
			_ = _987_v1
			_987_v1 = _dafny.SetOf(_986_v0)
			var _988_v2 _dafny.MultiSet
			_ = _988_v2
			_988_v2 = _dafny.MultiSetOf(_986_v0)
			var _989_v3 _dafny.Sequence
			_ = _989_v3
			_989_v3 = _dafny.SeqOf(_988_v2)
			var _990_v4 _dafny.Int
			_ = _990_v4
			_990_v4 = _dafny.IntOfInt64(794)
			_987_v1 = func() _dafny.Set {
				var _coll37 = _dafny.NewBuilder()
				_ = _coll37
				for _iter43 := _dafny.Iterate(((_989_v3).Select((Companion_Default___.SafeIndex(_990_v4, _dafny.IntOfUint32((_989_v3).Cardinality()))).Uint32()).(_dafny.MultiSet)).Elements()); ; {
					_compr_37, _ok43 := _iter43()
					if !_ok43 {
						break
					}
					var _991_v5 _dafny.Sequence
					_991_v5 = interface{}(_compr_37).(_dafny.Sequence)
					if ((_989_v3).Select((Companion_Default___.SafeIndex(_990_v4, _dafny.IntOfUint32((_989_v3).Cardinality()))).Uint32()).(_dafny.MultiSet)).Contains(_991_v5) {
						_coll37.Add(_991_v5)
					}
				}
				return _coll37.ToSet()
			}()
			var _992_v6 _dafny.Map
			_ = _992_v6
			_992_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_986_v0).Cardinality()), _990_v4)
			_992_v6 = (_992_v6).Update((_dafny.Zero).Minus((_990_v4).Plus(_990_v4)), (_dafny.Zero).Minus(_990_v4))
			var _993_v7 bool
			_ = _993_v7
			_993_v7 = true
			var _994_v8 _dafny.Sequence
			_ = _994_v8
			_994_v8 = _dafny.SeqOf(true, _993_v7)
			var _995_v9 _dafny.CodePoint
			_ = _995_v9
			_995_v9 = _dafny.CodePoint('s')
			(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if ((_this).Fm9(_986_v0, _990_v4, globalState)).Cmp(_990_v4) > 0 {
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_986_v0, _986_v0), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_994_v8).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_986_v0, _986_v0)).Cardinality()))).Uint32(), _dafny.CodePoint('t'))
				}
				return _986_v0
			})(), (Companion_Default___.SafeIndex(_990_v4, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if ((_this).Fm9(_986_v0, _990_v4, globalState)).Cmp(_990_v4) > 0 {
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_986_v0, _986_v0), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_994_v8).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_986_v0, _986_v0)).Cardinality()))).Uint32(), _dafny.CodePoint('t'))
				}
				return _986_v0
			})()).Cardinality()))).Uint32(), _995_v9)).Cardinality())
			_994_v8 = _994_v8
			var _996_v10 _dafny.Map
			_ = _996_v10
			_996_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_993_v7, _995_v9)
			_996_v10 = (_996_v10).Update(_993_v7, _995_v9)
		} else {
			var _997_v11 _dafny.Int
			_ = _997_v11
			_997_v11 = _dafny.IntOfInt64(-824)
			(globalState).F1 = (_997_v11).Times(_dafny.IntOfInt64(-559))
			var _998_v12 _dafny.CodePoint
			_ = _998_v12
			_998_v12 = _dafny.CodePoint('u')
			var _999_v13 bool
			_ = _999_v13
			_999_v13 = true
			var _1000_v14 _dafny.MultiSet
			_ = _1000_v14
			_1000_v14 = _dafny.MultiSetOf(_999_v13, !(true))
			var _1001_v15 _dafny.Set
			_ = _1001_v15
			_1001_v15 = _dafny.SetOf((_1000_v14).Cardinality(), _997_v11)
			var _1002_v16 _dafny.Map
			_ = _1002_v16
			_1002_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v13, _997_v11)
			var _1003_v17 D8
			_ = _1003_v17
			_1003_v17 = Companion_D8_.Create_DC20_(_998_v12, Companion_Default___.Fm4(_1001_v15, globalState), _999_v13, _1002_v16, _999_v13)
			var _1004_v18 _dafny.Map
			_ = _1004_v18
			_1004_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1003_v17).Dtor_cf43(), _998_v12)
			var _1005_v19 _dafny.Sequence
			_ = _1005_v19
			_1005_v19 = _dafny.SeqOf(_998_v12, _dafny.CodePoint('g'), _998_v12)
			_1004_v18 = (_1004_v18).Update(_998_v12, (_1005_v19).Select((Companion_Default___.SafeIndex(_997_v11, _dafny.IntOfUint32((_1005_v19).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _1006_v20 _dafny.Array
			_ = _1006_v20
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_30
			var _nw176 _dafny.Array
			_ = _nw176
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw176 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.CodePoint = (func(_1007_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1008_i0 _dafny.Int) _dafny.CodePoint {
						return _1007_v12
					}
				})(_998_v12)
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw176 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw176).ArraySet1CodePoint(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw176).ArraySet1CodePoint(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1006_v20 = _nw176
			_1006_v20 = _1006_v20
			r0 = !(_999_v13) || (_999_v13)
			var _1009_v21 _dafny.Map
			_ = _1009_v21
			_1009_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_997_v11, false)
			var _rhs145 _dafny.Int = _997_v11
			_ = _rhs145
			var _rhs146 bool = _999_v13
			_ = _rhs146
			var _rhs147 bool = Companion_Default___.Fm4(_1001_v15, globalState)
			_ = _rhs147
			var _rhs148 _dafny.Int = ((_1009_v21).Cardinality()).Times((_997_v11).Plus(_997_v11))
			_ = _rhs148
			var _rhs149 _dafny.Int = Companion_Default___.SafeDivisionInt(_997_v11, _997_v11)
			_ = _rhs149
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			_lhs104.F1 = _rhs145
			_999_v13 = _rhs146
			_999_v13 = _rhs147
			r3 = _rhs148
			r3 = _rhs149
		}
		var _1010_v22 bool
		_ = _1010_v22
		_1010_v22 = true
		var _1011_v23 _dafny.MultiSet
		_ = _1011_v23
		_1011_v23 = _dafny.MultiSetOf(_1010_v22)
		var _1012_v24 _dafny.MultiSet
		_ = _1012_v24
		_1012_v24 = _1011_v23
		var _source11 _dafny.MultiSet = _1012_v24
		_ = _source11
		var _1013___mcc_h0 _dafny.MultiSet = _source11
		_ = _1013___mcc_h0
		var _1014_cf41 _dafny.MultiSet = _1013___mcc_h0
		_ = _1014_cf41
		var _1015_v25 _dafny.Sequence
		_ = _1015_v25
		_1015_v25 = _dafny.SeqOf(false, _1010_v22)
		var _1016_v26 bool
		_ = _1016_v26
		var _1017_v27 _dafny.MultiSet
		_ = _1017_v27
		var _1018_v28 _dafny.Set
		_ = _1018_v28
		var _1019_v29 _dafny.Int
		_ = _1019_v29
		var _out32 bool
		_ = _out32
		var _out33 _dafny.MultiSet
		_ = _out33
		var _out34 _dafny.Set
		_ = _out34
		var _out35 _dafny.Int
		_ = _out35
		_out32, _out33, _out34, _out35 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_1015_v25, _1015_v25), _1015_v25, globalState)
		_1016_v26 = _out32
		_1017_v27 = _out33
		_1018_v28 = _out34
		_1019_v29 = _out35
		var _1020_v30 *C0
		_ = _1020_v30
		var _nw177 *C0 = New_C0_()
		_ = _nw177
		_nw177.Ctor__()
		_1020_v30 = _nw177
		_1019_v29 = _1019_v29
		var _1021_v31 _dafny.CodePoint
		_ = _1021_v31
		_1021_v31 = _dafny.CodePoint('t')
		(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.Fm19(_1021_v31, _1016_v26, globalState))
		var _1022_v32 _dafny.Int
		_ = _1022_v32
		_1022_v32 = _dafny.IntOfInt64(-895)
		var _1023_v33 _dafny.Map
		_ = _1023_v33
		_1023_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1022_v32, _1010_v22)
		var _1024_v34 _dafny.Array
		_ = _1024_v34
		var _nwElement0_43 bool = !(_1010_v22)
		_ = _nwElement0_43
		var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(5))
		_ = _nw178
		(_nw178).ArraySet1(_nwElement0_43, 0)
		(_nw178).ArraySet1((func() bool {
			if (_1023_v33).Contains(_1022_v32) {
				return (_1023_v33).Get(_1022_v32).(bool)
			}
			return _1010_v22
		})(), 1)
		(_nw178).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_1022_v32)).Cardinality())).Cmp(_1022_v32) > 0, 2)
		(_nw178).ArraySet1(_1010_v22, 3)
		(_nw178).ArraySet1(_1010_v22, 4)
		_1024_v34 = _nw178
		var _1025_v35 _dafny.Sequence
		_ = _1025_v35
		_1025_v35 = _dafny.UnicodeSeqOfUtf8Bytes("l")
		var _1026_v36 _dafny.MultiSet
		_ = _1026_v36
		_1026_v36 = _dafny.MultiSetOf(_1025_v35)
		var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
		_ = _index174
		(_1024_v34).ArraySet1((_1026_v36).IsDisjointFrom(Companion_Default___.Fm26(globalState)), (_index174).Int())
		var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
		_ = _index175
		(_1024_v34).ArraySet1(_1010_v22, (_index175).Int())
		var _1027_v37 _dafny.Set
		_ = _1027_v37
		_1027_v37 = _dafny.SetOf(_1022_v32, (_dafny.Zero).Minus(_1022_v32))
		var _1028_v38 _dafny.Sequence
		_ = _1028_v38
		_1028_v38 = _dafny.SeqOf(Companion_Default___.Fm4(_1027_v37, globalState))
		var _1029_v39 D9
		_ = _1029_v39
		_1029_v39 = Companion_D9_.Create_DC21_(_1028_v38)
		var _pat_let_tv22 = _1028_v38
		_ = _pat_let_tv22
		var _source12 D9 = func(_pat_let14_0 D9) D9 {
			return func(_1030_dt__update__tmp_h0 D9) D9 {
				return func(_pat_let15_0 _dafny.Sequence) D9 {
					return func(_1031_dt__update_hcf48_h0 _dafny.Sequence) D9 {
						return Companion_D9_.Create_DC21_(_1031_dt__update_hcf48_h0)
					}(_pat_let15_0)
				}(_pat_let_tv22)
			}(_pat_let14_0)
		}(_1029_v39)
		_ = _source12
		if _source12.Is_DC22() {
			var _1032___mcc_h1 _dafny.Int = _source12.Get_().(D9_DC22).Cf49
			_ = _1032___mcc_h1
			var _1033___mcc_h2 bool = _source12.Get_().(D9_DC22).Cf50
			_ = _1033___mcc_h2
			var _1034___mcc_h3 _dafny.Map = _source12.Get_().(D9_DC22).Cf51
			_ = _1034___mcc_h3
			var _1035_cf51 _dafny.Map = _1034___mcc_h3
			_ = _1035_cf51
			var _1036_cf50 bool = _1033___mcc_h2
			_ = _1036_cf50
			var _1037_cf49 _dafny.Int = _1032___mcc_h1
			_ = _1037_cf49
			var _1038_v40 *C2
			_ = _1038_v40
			var _nw179 *C2 = New_C2_()
			_ = _nw179
			_nw179.Ctor__(false, (_1037_cf49).Plus(_dafny.IntOfInt64(506)), (_1011_v23).Difference(_1011_v23))
			_1038_v40 = _nw179
			var _1039_v41 _dafny.Array
			_ = _1039_v41
			var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw180
			_1039_v41 = _nw180
			var _1040_v42 _dafny.Map
			_ = _1040_v42
			_1040_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_cf49, (func() _dafny.Int {
				if (_1011_v23).Contains((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool)) {
					return (_1011_v23).Multiplicity((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool))
				}
				return _1022_v32
			})())
			var _1041_v43 _dafny.CodePoint
			_ = _1041_v43
			_1041_v43 = _dafny.CodePoint('v')
			var _1042_v44 _dafny.Set
			_ = _1042_v44
			_1042_v44 = _dafny.SetOf(_1024_v34)
			var _1043_v45 _dafny.MultiSet
			_ = _1043_v45
			_1043_v45 = _dafny.MultiSetOf((_1038_v40).F13())
			var _1044_v46 _dafny.Map
			_ = _1044_v46
			_1044_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1028_v38).Select((Companion_Default___.SafeIndex(_1022_v32, _dafny.IntOfUint32((_1028_v38).Cardinality()))).Uint32()).(bool), (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool))
			var _1045_v47 _dafny.Map
			_ = _1045_v47
			_1045_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_v35, Companion_Default___.Fm27(_1041_v43, _dafny.IntOfInt64(-714), _1038_v40.F12, Companion_Default___.Fm24(_1037_cf49, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), _1023_v33, _1037_cf49, globalState), globalState))
			var _1046_v48 _dafny.Array
			_ = _1046_v48
			var _nwElement0_44 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_1036_cf50, true, _1010_v22)).Cardinality())
			_ = _nwElement0_44
			var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(16))
			_ = _nw181
			(_nw181).ArraySet1(_nwElement0_44, 0)
			(_nw181).ArraySet1((func() _dafny.Int {
				if (_1040_v42).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_1041_v43)).Cardinality())) {
					return (_1040_v42).Get(_dafny.IntOfUint32((_dafny.SeqOf(_1041_v43)).Cardinality())).(_dafny.Int)
				}
				return (_1042_v44).Cardinality()
			})(), 1)
			(_nw181).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vtmvqyk")).Cardinality())).Minus(_1022_v32)), 2)
			(_nw181).ArraySet1((_dafny.IntOfUint32((_1025_v35).Cardinality())).Plus(_dafny.IntOfUint32((_1025_v35).Cardinality())), 3)
			(_nw181).ArraySet1(_1037_cf49, 4)
			(_nw181).ArraySet1(_dafny.IntOfInt64(720), 5)
			(_nw181).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_1043_v45).Contains((_1038_v40).F13()) {
					return (_1043_v45).Multiplicity((_1038_v40).F13())
				}
				return (_1044_v46).Cardinality()
			})(), _1037_cf49), 6)
			(_nw181).ArraySet1(_dafny.IntOfInt64(-4), 7)
			(_nw181).ArraySet1(Companion_Default___.SafeModuloInt(_1037_cf49, _1037_cf49), 8)
			(_nw181).ArraySet1((_1038_v40).F13(), 9)
			(_nw181).ArraySet1(_1022_v32, 10)
			(_nw181).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-734), (_dafny.Zero).Minus((_1038_v40).Fm11(_1043_v45, globalState))), 11)
			(_nw181).ArraySet1((_1038_v40).F13(), 12)
			(_nw181).ArraySet1((_1045_v47).Cardinality(), 13)
			(_nw181).ArraySet1((_1022_v32).Times(_dafny.IntOfUint32((_1025_v35).Cardinality())), 14)
			(_nw181).ArraySet1(_dafny.IntOfUint32((_1028_v38).Cardinality()), 15)
			_1046_v48 = _nw181
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1046_v48), 0))
			_ = _index176
			(_1046_v48).ArraySet1((_dafny.Zero).Minus((_1038_v40).F13()), (_index176).Int())
			var _1047_v49 _dafny.Map
			_ = _1047_v49
			_1047_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1010_v22, Companion_Default___.Fm28((_1038_v40).Fm11(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1037_cf49)), globalState), globalState))
			var _1048_v50 _dafny.Map
			_ = _1048_v50
			_1048_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1022_v32)
			var _1049_v51 _dafny.Map
			_ = _1049_v51
			_1049_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if (_1047_v49).Contains(true) {
					return (_1047_v49).Get(true).(_dafny.Map)
				}
				return _1048_v50
			})(), _1039_v41)
			var _1050_v52 _dafny.Sequence
			_ = _1050_v52
			_1050_v52 = _dafny.SeqOf((func() _dafny.Map {
				if _1036_cf50 {
					return _1048_v50
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1037_cf49)
			})(), _1048_v50)
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1046_v48), 0))
			_ = _index177
			var _rhs150 _dafny.Array = (func() _dafny.Array {
				if (_1049_v51).Contains(_1048_v50) {
					return (_1049_v51).Get(_1048_v50).(_dafny.Array)
				}
				return _1039_v41
			})()
			_ = _rhs150
			var _rhs151 _dafny.Int = _dafny.IntOfUint32((_1050_v52).Cardinality())
			_ = _rhs151
			var _lhs105 _dafny.Array = _1046_v48
			_ = _lhs105
			var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1046_v48), 0))
			_ = _lhs106
			_1039_v41 = _rhs150
			(_lhs105).ArraySet1(_rhs151, (_lhs106).Int())
			_1010_v22 = _1010_v22
			var _1051_v53 *C0
			_ = _1051_v53
			var _nw182 *C0 = New_C0_()
			_ = _nw182
			_nw182.Ctor__()
			_1051_v53 = _nw182
			var _1052_v54 _dafny.Map
			_ = _1052_v54
			_1052_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1051_v53, false)
			var _1053_v55 _dafny.Map
			_ = _1053_v55
			_1053_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1028_v38, _dafny.Companion_Sequence_.Update(_1028_v38, (Companion_Default___.SafeIndex((_1052_v54).Cardinality(), _dafny.IntOfUint32((_1028_v38).Cardinality()))).Uint32(), _1010_v22)), (_dafny.Zero).Minus((_1046_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1046_v48), 0))).Int()).(_dafny.Int)))
			_1053_v55 = (_1053_v55).Update(_1028_v38, _1022_v32)
		} else if _source12.Is_DC23() {
			var _1054___mcc_h4 _dafny.Int = _source12.Get_().(D9_DC23).Cf52
			_ = _1054___mcc_h4
			var _1055_cf52 _dafny.Int = _1054___mcc_h4
			_ = _1055_cf52
			r3 = _1055_cf52
			var _1056_v56 _dafny.Sequence
			_ = _1056_v56
			_1056_v56 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate((_1029_v39).Dtor_cf48(), _dafny.SeqOf((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool))))
			var _rhs152 bool = _1010_v22
			_ = _rhs152
			var _rhs153 _dafny.Sequence = _1056_v56
			_ = _rhs153
			r0 = _rhs152
			_1056_v56 = _rhs153
			var _1057_v57 _dafny.MultiSet
			_ = _1057_v57
			_1057_v57 = _dafny.MultiSetOf(_1022_v32, _1022_v32, _1022_v32)
			var _1058_v58 _dafny.Sequence
			_ = _1058_v58
			_1058_v58 = _dafny.SeqOf(_1055_cf52)
			if (_dafny.MultiSetFromSeq(_1058_v58)).IsProperSubsetOf((_dafny.MultiSetFromSeq(Companion_Default___.Fm20(_1055_cf52, _1055_cf52, globalState))).Intersection(_1057_v57)) {
				var _1059_v59 _dafny.CodePoint
				_ = _1059_v59
				_1059_v59 = _dafny.CodePoint('n')
				var _1060_v60 D3
				_ = _1060_v60
				_1060_v60 = Companion_D3_.Create_DC8_(_1059_v59, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), _1010_v22, _1025_v35)
				var _pat_let_tv23 = _1024_v34
				_ = _pat_let_tv23
				var _pat_let_tv24 = _1024_v34
				_ = _pat_let_tv24
				var _pat_let_tv25 = _1059_v59
				_ = _pat_let_tv25
				_1060_v60 = func(_pat_let16_0 D3) D3 {
					return func(_1061_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let17_0 bool) D3 {
							return func(_1062_dt__update_hcf18_h0 bool) D3 {
								return func(_pat_let18_0 _dafny.CodePoint) D3 {
									return func(_1063_dt__update_hcf17_h0 _dafny.CodePoint) D3 {
										return Companion_D3_.Create_DC8_(_1063_dt__update_hcf17_h0, _1062_dt__update_hcf18_h0, (_1061_dt__update__tmp_h1).Dtor_cf19(), (_1061_dt__update__tmp_h1).Dtor_cf20())
									}(_pat_let18_0)
								}(_pat_let_tv25)
							}(_pat_let17_0)
						}((_pat_let_tv24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_pat_let_tv23), 0))).Int()).(bool))
					}(_pat_let16_0)
				}(_1060_v60)
				var _1064_v61 *C0
				_ = _1064_v61
				var _nw183 *C0 = New_C0_()
				_ = _nw183
				_nw183.Ctor__()
				_1064_v61 = _nw183
				var _1065_v62 _dafny.Map
				_ = _1065_v62
				_1065_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), _1055_cf52)
				var _1066_v63 _dafny.Map
				_ = _1066_v63
				_1066_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-903), _1024_v34)
				var _1067_v64 _dafny.Array
				_ = _1067_v64
				var _nwElement0_45 _dafny.Int = ((_1065_v62).Cardinality()).Plus(_1022_v32)
				_ = _nwElement0_45
				var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(6))
				_ = _nw184
				(_nw184).ArraySet1(_nwElement0_45, 0)
				(_nw184).ArraySet1((_1066_v63).Cardinality(), 1)
				(_nw184).ArraySet1(_1022_v32, 2)
				(_nw184).ArraySet1((func() _dafny.Int {
					if (_1057_v57).Contains(_1022_v32) {
						return (_1057_v57).Multiplicity(_1022_v32)
					}
					return _1055_cf52
				})(), 3)
				(_nw184).ArraySet1(((_1027_v37).Cardinality()).Times(_1055_cf52), 4)
				(_nw184).ArraySet1((_dafny.Zero).Minus(_1022_v32), 5)
				_1067_v64 = _nw184
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_1067_v64), 0))
				_ = _index178
				(_1067_v64).ArraySet1((func() _dafny.Int {
					if (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool) {
						return _1022_v32
					}
					return (_dafny.Zero).Minus((_1011_v23).Cardinality())
				})(), (_index178).Int())
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_1067_v64), 0))
				_ = _index179
				var _rhs154 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(956))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1068_v59 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1069_i1 _dafny.Int) _dafny.CodePoint {
						return _1068_v59
					}
				})(_1059_v59)))
				_ = _rhs154
				var _rhs155 _dafny.Int = _1022_v32
				_ = _rhs155
				var _lhs107 _dafny.Array = _1067_v64
				_ = _lhs107
				var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_1067_v64), 0))
				_ = _lhs108
				_1025_v35 = _rhs154
				(_lhs107).ArraySet1(_rhs155, (_lhs108).Int())
				var _1070_v65 *C2
				_ = _1070_v65
				var _nw185 *C2 = New_C2_()
				_ = _nw185
				_nw185.Ctor__((_1022_v32).Cmp((_1067_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_1067_v64), 0))).Int()).(_dafny.Int)) >= 0, _dafny.IntOfInt64(373), (_1012_v24))
				_1070_v65 = _nw185
				var _1071_v66 _dafny.Map
				_ = _1071_v66
				_1071_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1070_v65.F12, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool))
				_1010_v22 = (_1010_v22) == ((func() bool {
					if (_1071_v66).Contains(true) {
						return (_1071_v66).Get(true).(bool)
					}
					return _1070_v65.F12
				})())
			} else {
				(globalState).F1 = ((_dafny.Zero).Minus(_1022_v32)).Times(_1022_v32)
				var _1072_v67 _dafny.Set
				_ = _1072_v67
				_1072_v67 = _dafny.SetOf(_1025_v35)
				r3 = (_dafny.Zero).Minus((_1022_v32).Plus(((_1058_v58).Select((Companion_Default___.SafeIndex((_1072_v67).Cardinality(), _dafny.IntOfUint32((_1058_v58).Cardinality()))).Uint32()).(_dafny.Int)).Times(_1055_cf52)))
				var _len0_31 _dafny.Int = _dafny.One
				_ = _len0_31
				var _nw186 _dafny.Array
				_ = _nw186
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw186 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) bool = (func(_1073_v34 _dafny.Array) func(_dafny.Int) bool {
						return func(_1074_i2 _dafny.Int) bool {
							return (_1073_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1073_v34), 0))).Int()).(bool)
						}
					})(_1024_v34)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw186 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw186).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw186).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1024_v34 = _nw186
				var _1075_v68 _dafny.CodePoint
				_ = _1075_v68
				_1075_v68 = _dafny.CodePoint('a')
				r3 = Companion_Default___.Fm19(_1075_v68, _1010_v22, globalState)
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
				_ = _index180
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
				_ = _index181
				var _rhs156 _dafny.Int = (_1055_cf52).Plus(_1022_v32)
				_ = _rhs156
				var _rhs157 bool = (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool)
				_ = _rhs157
				var _rhs158 _dafny.Int = _1055_cf52
				_ = _rhs158
				var _rhs159 bool = (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool)
				_ = _rhs159
				var _rhs160 bool = (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool)
				_ = _rhs160
				var _lhs109 _dafny.Array = _1024_v34
				_ = _lhs109
				var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
				_ = _lhs110
				var _lhs111 _dafny.Array = _1024_v34
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
				_ = _lhs112
				_1022_v32 = _rhs156
				(_lhs109).ArraySet1(_rhs157, (_lhs110).Int())
				r2 = _rhs158
				(_lhs111).ArraySet1(_rhs159, (_lhs112).Int())
				_1010_v22 = _rhs160
			}
			var _1076_v69 _dafny.Array
			_ = _1076_v69
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_32
			var _nw187 _dafny.Array
			_ = _nw187
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw187 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Sequence = (func(_1077_v58 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1078_i3 _dafny.Int) _dafny.Sequence {
						return _1077_v58
					}
				})(_1058_v58)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw187 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw187).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw187).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1076_v69 = _nw187
			var _1079_v70 D6
			_ = _1079_v70
			_1079_v70 = Companion_D6_.Create_DC17_(_dafny.IntOfInt64(276), _1025_v35, _1010_v22)
			var _1080_v71 _dafny.Map
			_ = _1080_v71
			_1080_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1028_v38, _1010_v22)
			var _1081_v72 _dafny.Map
			_ = _1081_v72
			_1081_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1055_cf52, _1080_v71)
			var _rhs161 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("csulwnn")
			_ = _rhs161
			var _rhs162 _dafny.Sequence = (_1079_v70).Dtor_cf39()
			_ = _rhs162
			var _rhs163 _dafny.Array = _1076_v69
			_ = _rhs163
			var _rhs164 bool = (_dafny.IntOfUint32((_1058_v58).Cardinality())).Cmp(_dafny.IntOfUint32((_1058_v58).Cardinality())) > 0
			_ = _rhs164
			var _rhs165 bool = !(((func() _dafny.Map {
				if (_1081_v72).Contains(_1022_v32) {
					return (_1081_v72).Get(_1022_v32).(_dafny.Map)
				}
				return _1080_v71
			})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1028_v38, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool)))).Contains((func() _dafny.Sequence {
				if (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool) {
					return _dafny.SeqOf(_1010_v22, _1010_v22, true, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), false)
				}
				return _1028_v38
			})())
			_ = _rhs165
			_1025_v35 = _rhs161
			_1025_v35 = _rhs162
			_1076_v69 = _rhs163
			r0 = _rhs164
			r0 = _rhs165
		} else if _source12.Is_DC24() {
			var _1082___mcc_h5 _dafny.Int = _source12.Get_().(D9_DC24).Cf53
			_ = _1082___mcc_h5
			var _1083_cf53 _dafny.Int = _1082___mcc_h5
			_ = _1083_cf53
			var _1084_v73 *C1
			_ = _1084_v73
			var _nw188 *C1 = New_C1_()
			_ = _nw188
			_nw188.Ctor__(Companion_Default___.Fm4(_dafny.SetOf(_1022_v32, _dafny.IntOfUint32((_1028_v38).Cardinality())), globalState), _1010_v22, (_1011_v23).Intersection(_dafny.MultiSetFromSeq(_1028_v38)))
			_1084_v73 = _nw188
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
			_ = _index182
			(_1024_v34).ArraySet1((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), (_index182).Int())
			r0 = (_1084_v73).F15()
			var _1085_v74 _dafny.Sequence
			_ = _1085_v74
			_1085_v74 = _dafny.SeqOf((_1027_v37).Intersection(_1027_v37), _1027_v37)
			_1027_v37 = (_1085_v74).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1028_v38).Cardinality()), _dafny.IntOfUint32((_1085_v74).Cardinality()))).Uint32()).(_dafny.Set)
		} else {
			var _1086___mcc_h6 _dafny.Sequence = _source12.Get_().(D9_DC21).Cf48
			_ = _1086___mcc_h6
			var _1087_cf48 _dafny.Sequence = _1086___mcc_h6
			_ = _1087_cf48
			r0 = _1010_v22
			_1025_v35 = _dafny.UnicodeSeqOfUtf8Bytes("rd")
			var _1088_v75 _dafny.CodePoint
			_ = _1088_v75
			_1088_v75 = _dafny.CodePoint('o')
			_1010_v22 = !(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1025_v35, (Companion_Default___.SafeIndex(_1022_v32, _dafny.IntOfUint32((_1025_v35).Cardinality()))).Uint32(), _1088_v75), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-326))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}(func(_1089_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))), (Companion_Default___.SafeIndex(_1022_v32, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1025_v35, (Companion_Default___.SafeIndex(_1022_v32, _dafny.IntOfUint32((_1025_v35).Cardinality()))).Uint32(), _1088_v75), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-326))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}(func(_1090_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})))).Cardinality()))).Uint32(), _1088_v75), _1088_v75))
			var _1091_v76 _dafny.Map
			_ = _1091_v76
			_1091_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), Companion_Default___.Fm4(_1027_v37, globalState))
			var _1092_v77 _dafny.Map
			_ = _1092_v77
			_1092_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1022_v32, _1091_v76)
			var _1093_v78 _dafny.Set
			_ = _1093_v78
			_1093_v78 = _dafny.SetOf((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool))
			_1091_v76 = (func() _dafny.Map {
				if (_1092_v77).Contains(Companion_Default___.SafeModuloInt((_this).Fm9(_1025_v35, (_1093_v78).Cardinality(), globalState), _dafny.IntOfInt64(990))) {
					return (_1092_v77).Get(Companion_Default___.SafeModuloInt((_this).Fm9(_1025_v35, (_1093_v78).Cardinality(), globalState), _dafny.IntOfInt64(990))).(_dafny.Map)
				}
				return (Companion_Default___.Fm18((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), _1023_v33, globalState)).Update(_1010_v22, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool))
			})()
		}
		if _1010_v22 {
			_1022_v32 = _1022_v32
			var _1094_v79 _dafny.MultiSet
			_ = _1094_v79
			_1094_v79 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1022_v32), _1022_v32, _1022_v32)
			_1094_v79 = ((Companion_Default___.Fm29(_dafny.IntOfInt64(314), _1010_v22, false, globalState)).Update(_1022_v32, Companion_Default___.Abs(_dafny.IntOfUint32((_1025_v35).Cardinality())))).Union((func() _dafny.MultiSet {
				if _1010_v22 {
					return _1094_v79
				}
				return _1094_v79
			})())
			var _1095_v80 _dafny.CodePoint
			_ = _1095_v80
			_1095_v80 = _dafny.CodePoint('b')
			var _1096_v81 D3
			_ = _1096_v81
			_1096_v81 = Companion_D3_.Create_DC8_(_1095_v80, true, _1010_v22, Companion_Default___.Fm2(!(!(!(_1010_v22))), _1022_v32, _1095_v80, _1022_v32, globalState))
			var _1097_v82 _dafny.Sequence
			_ = _1097_v82
			_1097_v82 = _dafny.SeqOf(_1024_v34)
			var _1098_v83 _dafny.Set
			_ = _1098_v83
			_1098_v83 = _dafny.SetOf(_1010_v22, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool))
			(globalState).F1 = (_this).Fm9(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_1096_v81).Dtor_cf20(), (Companion_Default___.SafeIndex(_1022_v32, _dafny.IntOfUint32(((_1096_v81).Dtor_cf20()).Cardinality()))).Uint32(), Companion_Default___.Fm1(globalState)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1097_v82).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1096_v81).Dtor_cf20(), (Companion_Default___.SafeIndex(_1022_v32, _dafny.IntOfUint32(((_1096_v81).Dtor_cf20()).Cardinality()))).Uint32(), Companion_Default___.Fm1(globalState))).Cardinality()))).Uint32(), _1095_v80), ((_1098_v83).Cardinality()).Plus(_1022_v32), globalState)
			var _1099_v84 T1
			_ = _1099_v84
			var _nw189 *C3 = New_C3_()
			_ = _nw189
			_nw189.Ctor__(_1010_v22)
			_1099_v84 = _nw189
			var _1100_v85 _dafny.Map
			_ = _1100_v85
			_1100_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1099_v84, (Companion_Default___.Fm4(_1027_v37, globalState)) == (false))
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
			_ = _index183
			(_1024_v34).ArraySet1((func() bool {
				if (_1100_v85).Contains(_1099_v84) {
					return (_1100_v85).Get(_1099_v84).(bool)
				}
				return _1010_v22
			})(), (_index183).Int())
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
			_ = _index184
			(_1024_v34).ArraySet1(!(Companion_Default___.Fm4(_1027_v37, globalState)), (_index184).Int())
		} else {
			var _1101_v86 _dafny.Array
			_ = _1101_v86
			var _nwElement0_46 _dafny.Sequence = _1025_v35
			_ = _nwElement0_46
			var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(3))
			_ = _nw190
			(_nw190).ArraySet1(_nwElement0_46, 0)
			(_nw190).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("rxgwn"), (Companion_Default___.SafeIndex((_1011_v23).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rxgwn")).Cardinality()))).Uint32(), (_1025_v35).Select((Companion_Default___.SafeIndex(_1022_v32, _dafny.IntOfUint32((_1025_v35).Cardinality()))).Uint32()).(_dafny.CodePoint)), _1025_v35), 1)
			(_nw190).ArraySet1(_1025_v35, 2)
			_1101_v86 = _nw190
			var _1102_v87 _dafny.Array
			_ = _1102_v87
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw191
			_1102_v87 = _nw191
			var _1103_v88 _dafny.Array
			_ = _1103_v88
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_33
			var _nw192 _dafny.Array
			_ = _nw192
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw192 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) _dafny.Set = (func(_1104_v37 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_1105_i5 _dafny.Int) _dafny.Set {
						return _1104_v37
					}
				})(_1027_v37)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw192 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw192).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw192).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1103_v88 = _nw192
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_1103_v88), 0))
			_ = _index185
			(_1103_v88).ArraySet1(_1027_v37, (_index185).Int())
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
			_ = _index186
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_1103_v88), 0))
			_ = _index187
			var _rhs166 bool = (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool)
			_ = _rhs166
			var _rhs167 _dafny.Array = _1101_v86
			_ = _rhs167
			var _rhs168 _dafny.Int = _1022_v32
			_ = _rhs168
			var _rhs169 _dafny.Array = _1102_v87
			_ = _rhs169
			var _rhs170 _dafny.Set = _1027_v37
			_ = _rhs170
			var _lhs113 _dafny.Array = _1024_v34
			_ = _lhs113
			var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))
			_ = _lhs114
			var _lhs115 *GlobalState = globalState
			_ = _lhs115
			var _lhs116 _dafny.Array = _1103_v88
			_ = _lhs116
			var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_1103_v88), 0))
			_ = _lhs117
			(_lhs113).ArraySet1(_rhs166, (_lhs114).Int())
			_1101_v86 = _rhs167
			_lhs115.F1 = _rhs168
			_1102_v87 = _rhs169
			(_lhs116).ArraySet1(_rhs170, (_lhs117).Int())
			var _1106_v89 _dafny.Map
			_ = _1106_v89
			_1106_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1024_v34, _1022_v32)
			_1106_v89 = (_1106_v89).Update(_1024_v34, (_dafny.Zero).Minus((_1022_v32).Minus(_1022_v32)))
			(globalState).F1 = _1022_v32
			var _1107_v90 _dafny.Set
			_ = _1107_v90
			_1107_v90 = _dafny.SetOf(_1027_v37)
			var _1108_v91 _dafny.Array
			_ = _1108_v91
			var _nwElement0_47 _dafny.Int = (_dafny.Zero).Minus(_1022_v32)
			_ = _nwElement0_47
			var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(10))
			_ = _nw193
			(_nw193).ArraySet1(_nwElement0_47, 0)
			(_nw193).ArraySet1(_1022_v32, 1)
			(_nw193).ArraySet1(_1022_v32, 2)
			(_nw193).ArraySet1(_1022_v32, 3)
			(_nw193).ArraySet1(_1022_v32, 4)
			(_nw193).ArraySet1((_1106_v89).Cardinality(), 5)
			(_nw193).ArraySet1(((func() _dafny.Set {
				if (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool) {
					return _1107_v90
				}
				return _1107_v90
			})()).Cardinality(), 6)
			(_nw193).ArraySet1(_1022_v32, 7)
			(_nw193).ArraySet1(_1022_v32, 8)
			(_nw193).ArraySet1((_dafny.IntOfInt64(847)).Plus(_1022_v32), 9)
			_1108_v91 = _nw193
			_1108_v91 = _1108_v91
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1101_v86), 0))
			_ = _index188
			(_1101_v86).ArraySet1(_1025_v35, (_index188).Int())
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1101_v86), 0))
			_ = _index189
			(_1101_v86).ArraySet1(_1025_v35, (_index189).Int())
		}
		var _1109_v93 _dafny.Array
		_ = _1109_v93
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_34
		var _nw194 _dafny.Array
		_ = _nw194
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw194 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) _dafny.Set = (func(_1110_v32 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_1111_i7 _dafny.Int) _dafny.Set {
					return func() _dafny.Set {
						var _coll38 = _dafny.NewBuilder()
						_ = _coll38
						for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(555), _dafny.IntOfInt64(410))); ; {
							_compr_38, _ok44 := _iter44()
							if !_ok44 {
								break
							}
							var _1112_v92 _dafny.Int
							_1112_v92 = interface{}(_compr_38).(_dafny.Int)
							if ((_dafny.IntOfInt64(555)).Cmp(_1112_v92) <= 0) && ((_1112_v92).Cmp(_dafny.IntOfInt64(410)) < 0) {
								_coll38.Add((_1112_v92).Times(_1110_v32))
							}
						}
						return _coll38.ToSet()
					}()
				}
			})(_1022_v32)
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw194 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw194).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw194).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1109_v93 = _nw194
		for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1109_v93), 0))); ; {
			_guard_loop_5, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _1113_i6 _dafny.Int
			_1113_i6 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1113_i6).Sign() != -1) && ((_1113_i6).Cmp(_dafny.ArrayLen((_1109_v93), 0)) < 0)) {
				(_1109_v93).ArraySet1((_1027_v37).Intersection(_1027_v37), (_1113_i6).Int())
			}
		}
		var _1114_v94 D0
		_ = _1114_v94
		_1114_v94 = Companion_D0_.Create_DC0_(_dafny.CodePoint('e'))
		var _pat_let_tv26 = _1024_v34
		_ = _pat_let_tv26
		var _pat_let_tv27 = _1024_v34
		_ = _pat_let_tv27
		var _pat_let_tv28 = _1025_v35
		_ = _pat_let_tv28
		r0 = !(func(_source13 D0) bool {
			if _source13.Is_DC1() {
				var _1115___mcc_h7 bool = _source13.Get_().(D0_DC1).Cf1
				_ = _1115___mcc_h7
				var _1116___mcc_h8 _dafny.Int = _source13.Get_().(D0_DC1).Cf2
				_ = _1116___mcc_h8
				var _1117_cf2 _dafny.Int = _1116___mcc_h8
				_ = _1117_cf2
				var _1118_cf1 bool = _1115___mcc_h7
				_ = _1118_cf1
				return !(_1118_cf1)
			} else if _source13.Is_DC0() {
				var _1119___mcc_h9 _dafny.CodePoint = _source13.Get_().(D0_DC0).Cf0
				_ = _1119___mcc_h9
				var _1120_cf0 _dafny.CodePoint = _1119___mcc_h9
				_ = _1120_cf0
				return !(!((_pat_let_tv27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_pat_let_tv26), 0))).Int()).(bool)))
			} else {
				var _1121___mcc_h10 D0 = _source13.Get_().(D0_DC2).Cf3
				_ = _1121___mcc_h10
				var _1122_cf3 D0 = _1121___mcc_h10
				_ = _1122_cf3
				return !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("xq"), _pat_let_tv28)
			}
		}(Companion_D0_.Create_DC2_(_1114_v94)))
		var _nwElement0_48 _dafny.Array = _1024_v34
		_ = _nwElement0_48
		var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(4))
		_ = _nw195
		(_nw195).ArraySet1(_nwElement0_48, 0)
		(_nw195).ArraySet1(_1024_v34, 1)
		(_nw195).ArraySet1(_1024_v34, 2)
		(_nw195).ArraySet1(_1024_v34, 3)
		r1 = _nw195
		r2 = _1022_v32
		var _1123_v95 _dafny.CodePoint
		_ = _1123_v95
		_1123_v95 = _dafny.CodePoint('w')
		var _1124_v96 _dafny.Map
		_ = _1124_v96
		_1124_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), Companion_Default___.Fm19(_1123_v95, (_1024_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_1024_v34), 0))).Int()).(bool), globalState))
		r3 = (_1022_v32).Minus(Companion_Default___.SafeModuloInt(_1022_v32, (func() _dafny.Int {
			if (_1124_v96).Contains(true) {
				return (_1124_v96).Get(true).(_dafny.Int)
			}
			return _1022_v32
		})()))
		return r0, r1, r2, r3
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f10 _dafny.MultiSet
	_f11 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f10 = _dafny.EmptyMultiSet
	_this._f11 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_, Companion_T2_.TraitID_}
}

var _ T0 = &C5{}
var _ T1 = &C5{}
var _ T2 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F10() _dafny.MultiSet {
	return _this._f10
}
func (_this *C5) Ctor__(f11 bool, f10 _dafny.MultiSet) {
	{
		(_this)._f11 = f11
		(_this)._f10 = f10
	}
}
func (_this *C5) Fm5(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C5) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}(func(_1125_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})), _dafny.UnicodeSeqOfUtf8Bytes("pqoiu"))).Cardinality())
	}
}
func (_this *C5) Fm7(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C5) M1(p0 _dafny.CodePoint, p1 _dafny.Int, p2 D0, globalState *GlobalState) {
	{
		var _1126_v0 _dafny.Sequence
		_ = _1126_v0
		_1126_v0 = _dafny.UnicodeSeqOfUtf8Bytes("wbyaw")
		var _1127_v1 D0
		_ = _1127_v1
		_1127_v1 = Companion_D0_.Create_DC0_(p0)
		var _1128_v2 _dafny.Map
		_ = _1128_v2
		_1128_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _1129_v3 _dafny.MultiSet
		_ = _1129_v3
		_1129_v3 = _dafny.MultiSetOf(_dafny.IntOfInt64(721))
		var _1130_v4 _dafny.Set
		_ = _1130_v4
		_1130_v4 = _dafny.SetOf((_1128_v2).Cardinality(), (Companion_D0_.Create_DC1_((_this).F11(), (_1129_v3).Cardinality())).Dtor_cf2())
		var _1131_v5 _dafny.Sequence
		_ = _1131_v5
		_1131_v5 = _dafny.SeqOf((_this).F11())
		var _1132_v6 _dafny.Sequence
		_ = _1132_v6
		_1132_v6 = _dafny.SeqOf(p1, p1, p1)
		var _1133_v7 _dafny.Set
		_ = _1133_v7
		_1133_v7 = _dafny.SetOf((_this).F11())
		var _1134_v8 _dafny.Array
		_ = _1134_v8
		var _len0_35 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_35
		var _nw196 _dafny.Array
		_ = _nw196
		if _len0_35.Cmp(_dafny.Zero) == 0 {
			_nw196 = _dafny.NewArray(_len0_35)
		} else {
			var _init35 func(_dafny.Int) _dafny.Int = (func(_1135_p2 D0) func(_dafny.Int) _dafny.Int {
				return func(_1136_i0 _dafny.Int) _dafny.Int {
					return (_1136_i0).Minus((_1135_p2).Dtor_cf2())
				}
			})(p2)
			_ = _init35
			var _element0_35 = _init35(_dafny.Zero)
			_ = _element0_35
			_nw196 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
			(_nw196).ArraySet1(_element0_35, 0)
			var _nativeLen0_35 = (_len0_35).Int()
			_ = _nativeLen0_35
			for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
				(_nw196).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
			}
		}
		_1134_v8 = _nw196
		var _1137_v9 _dafny.Map
		_ = _1137_v9
		_1137_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1134_v8, p1)
		var _1138_v10 _dafny.Array
		_ = _1138_v10
		var _nwElement0_49 _dafny.Int = Companion_Default___.SafeModuloInt((_this).Fm6((_this).F11(), _dafny.IntOfInt64(657), globalState), _dafny.IntOfInt64(641))
		_ = _nwElement0_49
		var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(19))
		_ = _nw197
		(_nw197).ArraySet1(_nwElement0_49, 0)
		(_nw197).ArraySet1(p1, 1)
		(_nw197).ArraySet1(_dafny.IntOfInt64(-842), 2)
		(_nw197).ArraySet1(p1, 3)
		(_nw197).ArraySet1(p1, 4)
		(_nw197).ArraySet1(p1, 5)
		(_nw197).ArraySet1((_this).Fm6(Companion_Default___.Fm4(_1130_v4, globalState), p1, globalState), 6)
		(_nw197).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1131_v5, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1131_v5).Cardinality()))).Uint32(), (_this).F11())).Cardinality()), (_1132_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1132_v6).Cardinality()))).Uint32()).(_dafny.Int)), 7)
		(_nw197).ArraySet1(p1, 8)
		(_nw197).ArraySet1((_1133_v7).Cardinality(), 9)
		(_nw197).ArraySet1((p1).Times((_dafny.Zero).Minus(p1)), 10)
		(_nw197).ArraySet1((_this).Fm6(false, p1, globalState), 11)
		(_nw197).ArraySet1((_1137_v9).Cardinality(), 12)
		(_nw197).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1126_v0, _1126_v0)).Cardinality()), 13)
		(_nw197).ArraySet1((p1).Times(_dafny.IntOfInt64(407)), 14)
		(_nw197).ArraySet1(p1, 15)
		(_nw197).ArraySet1(p1, 16)
		(_nw197).ArraySet1(_dafny.IntOfInt64(954), 17)
		(_nw197).ArraySet1(p1, 18)
		_1138_v10 = _nw197
		var _1139_v11 D0
		_ = _1139_v11
		_1139_v11 = Companion_D0_.Create_DC1_((_this).F11(), p1)
		var _1140_v12 D0
		_ = _1140_v12
		_1140_v12 = Companion_D0_.Create_DC2_(_1139_v11)
		var _1141_v13 _dafny.Map
		_ = _1141_v13
		_1141_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _rhs171 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1126_v0, _1126_v0)
		_ = _rhs171
		var _rhs172 D0 = _1127_v1
		_ = _rhs172
		var _rhs173 _dafny.Array = _1134_v8
		_ = _rhs173
		var _rhs174 D0 = _1140_v12
		_ = _rhs174
		var _rhs175 _dafny.Map = _1141_v13
		_ = _rhs175
		_1126_v0 = _rhs171
		_1127_v1 = _rhs172
		_1138_v10 = _rhs173
		_1140_v12 = _rhs174
		_1141_v13 = _rhs175
		if _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
			if (p2).Dtor_cf1() {
				return _1132_v6
			}
			return _1132_v6
		})(), Companion_Default___.Fm8(Companion_Default___.Fm4(_1130_v4, globalState), globalState)) {
			var _1142_v14 bool
			_ = _1142_v14
			_1142_v14 = false
			_1142_v14 = !((_dafny.IntOfInt64(423)).Cmp(p1) > 0)
			var _1143_v15 bool
			_ = _1143_v15
			var _1144_v16 _dafny.MultiSet
			_ = _1144_v16
			var _1145_v17 _dafny.Set
			_ = _1145_v17
			var _1146_v18 _dafny.Int
			_ = _1146_v18
			var _out36 bool
			_ = _out36
			var _out37 _dafny.MultiSet
			_ = _out37
			var _out38 _dafny.Set
			_ = _out38
			var _out39 _dafny.Int
			_ = _out39
			_out36, _out37, _out38, _out39 = Companion_Default___.M0(_dafny.SeqOf(true, !(!(true))), _1131_v5, globalState)
			_1143_v15 = _out36
			_1144_v16 = _out37
			_1145_v17 = _out38
			_1146_v18 = _out39
			_1146_v18 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32((_1126_v0).Cardinality())).Minus(_1146_v18), _1146_v18))
			_1142_v14 = (_1146_v18).Cmp(p1) != 0
			var _1147_v19 _dafny.Map
			_ = _1147_v19
			_1147_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1146_v18, (_this).F11())
			(globalState).F1 = (_dafny.Zero).Minus((_1132_v6).Select((Companion_Default___.SafeIndex((_1146_v18).Times((_1147_v19).Cardinality()), _dafny.IntOfUint32((_1132_v6).Cardinality()))).Uint32()).(_dafny.Int))
		} else {
			var _1148_v20 *C3
			_ = _1148_v20
			var _nw198 *C3 = New_C3_()
			_ = _nw198
			_nw198.Ctor__((_this).F11())
			_1148_v20 = _nw198
			var _1149_v21 bool
			_ = _1149_v21
			_1149_v21 = true
			var _rhs176 bool = _1149_v21
			_ = _rhs176
			var _rhs177 bool = (_this).F11()
			_ = _rhs177
			_1149_v21 = _rhs176
			_1149_v21 = _rhs177
			if _1149_v21 {
				var _1150_v22 _dafny.Map
				_ = _1150_v22
				_1150_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F11())
				var _1151_v24 _dafny.Map
				_ = _1151_v24
				_1151_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll39 = _dafny.NewMapBuilder()
					_ = _coll39
					for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(213), _dafny.IntOfInt64(198))); ; {
						_compr_39, _ok46 := _iter46()
						if !_ok46 {
							break
						}
						var _1152_v23 _dafny.Int
						_1152_v23 = interface{}(_compr_39).(_dafny.Int)
						if ((_dafny.IntOfInt64(213)).Cmp(_1152_v23) <= 0) && ((_1152_v23).Cmp(_dafny.IntOfInt64(198)) < 0) {
							_coll39.Add(Companion_Default___.SafeDivisionInt(_1152_v23, p1), p1)
						}
					}
					return _coll39.ToMap()
				}()).Cardinality(), _1149_v21)
				var _1153_v25 _dafny.Set
				_ = _1153_v25
				_1153_v25 = _dafny.SetOf((_1150_v22).Update((_1148_v20).F16(), false), Companion_Default___.Fm18(true, _1151_v24, globalState), (func() _dafny.Map {
					if _1149_v21 {
						return _1150_v22
					}
					return _1150_v22
				})())
				(globalState).F1 = (_1153_v25).Cardinality()
				var _1154_v26 _dafny.Array
				_ = _1154_v26
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_36
				var _nw199 _dafny.Array
				_ = _nw199
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw199 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Set = (func(_1155_v7 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_1156_i1 _dafny.Int) _dafny.Set {
							return (_1155_v7).Union(_dafny.SetOf(false))
						}
					})(_1133_v7)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw199 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw199).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw199).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1154_v26 = _nw199
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_1154_v26), 0))
				_ = _index190
				(_1154_v26).ArraySet1(_dafny.SetOf((_1148_v20).F16(), (_1148_v20).F16(), (_1148_v20).F16()), (_index190).Int())
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_1154_v26), 0))
				_ = _index191
				(_1154_v26).ArraySet1(_1133_v7, (_index191).Int())
				(globalState).F1 = (_dafny.IntOfInt64(483)).Minus(_dafny.IntOfUint32((_1126_v0).Cardinality()))
				var _1157_v27 _dafny.Array
				_ = _1157_v27
				var _nw200 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw200
				_1157_v27 = _nw200
				var _1158_v28 _dafny.Sequence
				_ = _1158_v28
				_1158_v28 = _dafny.SeqOf(_1157_v27, _1157_v27)
				var _1159_v29 _dafny.Sequence
				_ = _1159_v29
				_1159_v29 = _dafny.SeqOf((func() _dafny.Array {
					if (_1148_v20).F16() {
						return _1157_v27
					}
					return _1157_v27
				})(), (func() _dafny.Array {
					if (_1148_v20).F16() {
						return _1157_v27
					}
					return _1157_v27
				})(), _1157_v27, (_1158_v28).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1158_v28).Cardinality()))).Uint32()).(_dafny.Array))
				var _rhs178 _dafny.Int = p1
				_ = _rhs178
				var _rhs179 _dafny.Array = (_1158_v28).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_1158_v28).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs179
				var _rhs180 _dafny.Array = _1134_v8
				_ = _rhs180
				var _rhs181 _dafny.Sequence = _1126_v0
				_ = _rhs181
				var _lhs118 *GlobalState = globalState
				_ = _lhs118
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				_lhs118.F1 = _rhs178
				_lhs119.F5 = _rhs179
				_1134_v8 = _rhs180
				_1126_v0 = _rhs181
				(globalState).F1 = (Companion_Default___.Fm14((_1148_v20).F16(), globalState)).Cardinality()
			} else {
				var _1160_v30 _dafny.Array
				_ = _1160_v30
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_37
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) bool = (func(_1161_v20 *C3) func(_dafny.Int) bool {
						return func(_1162_i2 _dafny.Int) bool {
							return (_1161_v20).F16()
						}
					})(_1148_v20)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw201 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw201).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw201).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1160_v30 = _nw201
				var _1163_v31 D6
				_ = _1163_v31
				_1163_v31 = Companion_D6_.Create_DC17_(_dafny.IntOfUint32((_1131_v5).Cardinality()), _1126_v0, true)
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))
				_ = _index192
				(_1160_v30).ArraySet1((_1149_v21) && ((_1163_v31).Dtor_cf40()), (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1160_v30), 0))
				_ = _index193
				(_1160_v30).ArraySet1((_1148_v20).F16(), (_index193).Int())
				var _1164_v32 _dafny.Map
				_ = _1164_v32
				_1164_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1149_v21, p1)
				var _1165_v33 _dafny.Map
				_ = _1165_v33
				_1165_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F11())
				var _1166_v34 _dafny.Sequence
				_ = _1166_v34
				_1166_v34 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xrghkd"))
				var _1167_v35 *C2
				_ = _1167_v35
				var _nw202 *C2 = New_C2_()
				_ = _nw202
				_nw202.Ctor__((_this).F11(), p1, (_this).F10())
				_1167_v35 = _nw202
				var _1168_v36 _dafny.Map
				_ = _1168_v36
				_1168_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1148_v20).F16(), (_1148_v20).F16())).Cardinality(), (_1165_v33).Cardinality(), _dafny.IntOfUint32((_1166_v34).Cardinality())), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1148_v20).F16(), (_1148_v20).F16())).Cardinality(), (_1165_v33).Cardinality(), _dafny.IntOfUint32((_1166_v34).Cardinality()))).Cardinality()))).Uint32(), p1), _1167_v35)
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))
				_ = _index194
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1160_v30), 0))
				_ = _index195
				var _rhs182 bool = ((func() _dafny.Map {
					if _1149_v21 {
						return _1164_v32
					}
					return _1164_v32
				})()).Equals((_1164_v32).Merge(_1164_v32))
				_ = _rhs182
				var _rhs183 bool = (_1168_v36).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1132_v6, _1167_v35))
				_ = _rhs183
				var _lhs120 _dafny.Array = _1160_v30
				_ = _lhs120
				var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))
				_ = _lhs121
				var _lhs122 _dafny.Array = _1160_v30
				_ = _lhs122
				var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1160_v30), 0))
				_ = _lhs123
				(_lhs120).ArraySet1(_rhs182, (_lhs121).Int())
				(_lhs122).ArraySet1(_rhs183, (_lhs123).Int())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))
				_ = _index196
				var _rhs184 bool = ((false) == ((_1160_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))).Int()).(bool))) && ((_1148_v20).F16())
				_ = _rhs184
				var _rhs185 _dafny.Int = _dafny.IntOfInt64(730)
				_ = _rhs185
				var _lhs124 _dafny.Array = _1160_v30
				_ = _lhs124
				var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))
				_ = _lhs125
				var _lhs126 *GlobalState = globalState
				_ = _lhs126
				(_lhs124).ArraySet1(_rhs184, (_lhs125).Int())
				_lhs126.F1 = _rhs185
				var _1169_v37 _dafny.Array
				_ = _1169_v37
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
				_ = _nw203
				_1169_v37 = _nw203
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_1169_v37), 0))
				_ = _index197
				(_1169_v37).ArraySet1CodePoint(_dafny.CodePoint('s'), (_index197).Int())
				var _1170_v38 D8
				_ = _1170_v38
				_1170_v38 = Companion_D8_.Create_DC20_(p0, (_this).F11(), _1149_v21, _1164_v32, _1149_v21)
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_1169_v37), 0))
				_ = _index198
				var _rhs186 _dafny.Int = ((func() _dafny.Int {
					if (_1164_v32).Contains((_this).F11()) {
						return (_1164_v32).Get((_this).F11()).(_dafny.Int)
					}
					return p1
				})()).Minus(Companion_Default___.SafeDivisionInt(p1, (_1167_v35).F13()))
				_ = _rhs186
				var _rhs187 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_1148_v20).F16() {
						return p0
					}
					return (func() _dafny.CodePoint {
						if _1167_v35.F12 {
							return p0
						}
						return (_1170_v38).Dtor_cf43()
					})()
				})()
				_ = _rhs187
				var _rhs188 _dafny.Int = (_1167_v35).F13()
				_ = _rhs188
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				var _lhs128 _dafny.Array = _1169_v37
				_ = _lhs128
				var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_1169_v37), 0))
				_ = _lhs129
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				_lhs127.F1 = _rhs186
				(_lhs128).ArraySet1CodePoint(_rhs187, (_lhs129).Int())
				_lhs130.F1 = _rhs188
				(_1167_v35).F12 = (_1160_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))).Int()).(bool)
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1160_v30), 0))
				_ = _index199
				(_1160_v30).ArraySet1(_1149_v21, (_index199).Int())
			}
			if (_this).F11() {
				_1131_v5 = _1131_v5
				(globalState).F1 = (p1).Times(p1)
				_1126_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(260))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_1171_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1172_i3 _dafny.Int) _dafny.CodePoint {
						return _1171_p0
					}
				})(p0))), _1126_v0), _1126_v0)
				var _1173_v39 _dafny.Map
				_ = _1173_v39
				_1173_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1134_v8, !(_1149_v21) || ((_1148_v20).F16()))
				_1173_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1134_v8, (_1148_v20).F16())
				var _1174_v40 _dafny.Map
				_ = _1174_v40
				_1174_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1149_v21, p1)
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1134_v8), 0))
				_ = _index200
				(_1134_v8).ArraySet1(((_1174_v40).Update((_this).F11(), p1)).Cardinality(), (_index200).Int())
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1134_v8), 0))
				_ = _index201
				(_1134_v8).ArraySet1(p1, (_index201).Int())
			} else {
				var _1175_v41 _dafny.Array
				_ = _1175_v41
				var _nw204 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw204
				_1175_v41 = _nw204
				(globalState).F5 = _1175_v41
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_1175_v41), 0))
				_ = _index202
				(_1175_v41).ArraySet1((_1148_v20).F16(), (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_1175_v41), 0))
				_ = _index203
				(_1175_v41).ArraySet1(_1149_v21, (_index203).Int())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_1175_v41), 0))
				_ = _index204
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_1175_v41), 0))
				_ = _index205
				var _rhs189 _dafny.Int = (_dafny.IntOfInt64(322)).Minus(p1)
				_ = _rhs189
				var _rhs190 bool = (_1148_v20).F16()
				_ = _rhs190
				var _rhs191 _dafny.Int = (_dafny.IntOfInt64(618)).Plus(p1)
				_ = _rhs191
				var _rhs192 _dafny.Int = p1
				_ = _rhs192
				var _rhs193 bool = (_this).Fm5((_1133_v7).Union(_1133_v7), p1, (_1132_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1132_v6).Cardinality()))).Uint32()).(_dafny.Int), globalState)
				_ = _rhs193
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				var _lhs132 _dafny.Array = _1175_v41
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_1175_v41), 0))
				_ = _lhs133
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				var _lhs135 *GlobalState = globalState
				_ = _lhs135
				var _lhs136 _dafny.Array = _1175_v41
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_1175_v41), 0))
				_ = _lhs137
				_lhs131.F1 = _rhs189
				(_lhs132).ArraySet1(_rhs190, (_lhs133).Int())
				_lhs134.F1 = _rhs191
				_lhs135.F1 = _rhs192
				(_lhs136).ArraySet1(_rhs193, (_lhs137).Int())
				(globalState).F1 = Companion_Default___.SafeModuloInt(p1, p1)
				_1149_v21 = (_1148_v20).F16()
			}
			var _1176_v42 _dafny.Map
			_ = _1176_v42
			_1176_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1148_v20).F16())
			_1176_v42 = (_1176_v42).Merge(_1176_v42)
		}
		if (_this).F11() {
			var _1177_v43 _dafny.Array
			_ = _1177_v43
			var _nwElement0_50 bool = (_this).F11()
			_ = _nwElement0_50
			var _nw205 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(12))
			_ = _nw205
			(_nw205).ArraySet1(_nwElement0_50, 0)
			(_nw205).ArraySet1((_this).F11(), 1)
			(_nw205).ArraySet1((_this).F11(), 2)
			(_nw205).ArraySet1((_this).F11(), 3)
			(_nw205).ArraySet1((_this).F11(), 4)
			(_nw205).ArraySet1((_this).F11(), 5)
			(_nw205).ArraySet1((_this).F11(), 6)
			(_nw205).ArraySet1(!((_this).F11()), 7)
			(_nw205).ArraySet1(false, 8)
			(_nw205).ArraySet1((_this).F11(), 9)
			(_nw205).ArraySet1((_this).F11(), 10)
			(_nw205).ArraySet1(true, 11)
			_1177_v43 = _nw205
			var _1178_v44 _dafny.Array
			_ = _1178_v44
			var _nwElement0_51 _dafny.Array = (func() _dafny.Array {
				if (_this).F11() {
					return _1177_v43
				}
				return _1177_v43
			})()
			_ = _nwElement0_51
			var _nw206 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(5))
			_ = _nw206
			(_nw206).ArraySet1(_nwElement0_51, 0)
			(_nw206).ArraySet1(_1177_v43, 1)
			(_nw206).ArraySet1(_1177_v43, 2)
			(_nw206).ArraySet1(_1177_v43, 3)
			(_nw206).ArraySet1(_1177_v43, 4)
			_1178_v44 = _nw206
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))
			_ = _index206
			(_1178_v44).ArraySet1(_1177_v43, (_index206).Int())
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))
			_ = _index207
			(_1178_v44).ArraySet1(_1177_v43, (_index207).Int())
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1134_v8), 0))
			_ = _index208
			(_1134_v8).ArraySet1(p1, (_index208).Int())
			var _1179_v45 bool
			_ = _1179_v45
			_1179_v45 = false
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1134_v8), 0))
			_ = _index209
			var _rhs194 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1132_v6, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1, p1), _1132_v6))
			_ = _rhs194
			var _rhs195 _dafny.Int = p1
			_ = _rhs195
			var _rhs196 _dafny.Int = (_this).Fm6(_1179_v45, p1, globalState)
			_ = _rhs196
			var _rhs197 bool = (!(false)) && ((_this).F11())
			_ = _rhs197
			var _lhs138 *GlobalState = globalState
			_ = _lhs138
			var _lhs139 _dafny.Array = _1134_v8
			_ = _lhs139
			var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1134_v8), 0))
			_ = _lhs140
			_1132_v6 = _rhs194
			_lhs138.F1 = _rhs195
			(_lhs139).ArraySet1(_rhs196, (_lhs140).Int())
			_1179_v45 = _rhs197
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_1178_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))).Int()))
			_ = _arr0
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_dafny.ArrayCastTo((_1178_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))).Int()))), 0))
			_ = _index210
			(_arr0).ArraySet1(_1179_v45, (_index210).Int())
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_1178_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))).Int()))
			_ = _arr1
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_dafny.ArrayCastTo((_1178_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))).Int()))), 0))
			_ = _index211
			(_arr1).ArraySet1((Companion_Default___.Fm19(p0, _1179_v45, globalState)).Cmp(p1) >= 0, (_index211).Int())
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1178_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))).Int()))
			_ = _arr2
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_dafny.ArrayCastTo((_1178_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_1178_v44), 0))).Int()))), 0))
			_ = _index212
			(_arr2).ArraySet1((_this).F11(), (_index212).Int())
			var _1180_v46 _dafny.CodePoint
			_ = _1180_v46
			_1180_v46 = _dafny.CodePoint('x')
			_1180_v46 = p0
		} else {
			var _1181_v47 bool
			_ = _1181_v47
			_1181_v47 = true
			_1181_v47 = (p1).Cmp(p1) == 0
			if _1181_v47 {
				(globalState).F1 = (func() _dafny.Int {
					if !((_this).F11()) || (_1181_v47) {
						return (p1).Times(p1)
					}
					return p1
				})()
				_1181_v47 = !(false)
				_1181_v47 = !(_1181_v47)
				(globalState).F1 = p1
				var _1182_v48 T2
				_ = _1182_v48
				var _nw207 *C0 = New_C0_()
				_ = _nw207
				_nw207.Ctor__()
				_1182_v48 = _nw207
				_1182_v48 = _1182_v48
			} else {
				var _1183_v49 *C4
				_ = _1183_v49
				var _nw208 *C4 = New_C4_()
				_ = _nw208
				_nw208.Ctor__()
				_1183_v49 = _nw208
				var _1184_v50 D3
				_ = _1184_v50
				_1184_v50 = Companion_D3_.Create_DC8_(_dafny.CodePoint('m'), (_this).F11(), Companion_Default___.Fm4(_1130_v4, globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jhadxaocd"), _dafny.Companion_Sequence_.Update(_1126_v0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1126_v0).Cardinality()))).Uint32(), _dafny.CodePoint('k'))))
				_1184_v50 = _1184_v50
				(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p1, (_dafny.Zero).Minus((((_this).F10()).Difference(((_this).F10()).Update(_1181_v47, Companion_Default___.Abs(p1)))).Cardinality())))
				var _1185_v51 *C2
				_ = _1185_v51
				var _nw209 *C2 = New_C2_()
				_ = _nw209
				_nw209.Ctor__(!(_1130_v4).Contains(_dafny.IntOfUint32((_1132_v6).Cardinality())), (p1).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("kwiask"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1126_v0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1126_v0).Cardinality()))).Uint32(), p0)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kwiask")).Cardinality()))).Uint32(), _dafny.CodePoint('e'))).Cardinality())), (_this).F10())
				_1185_v51 = _nw209
				(globalState).F1 = (_1185_v51).F13()
			}
			var _1186_v52 *C2
			_ = _1186_v52
			var _nw210 *C2 = New_C2_()
			_ = _nw210
			_nw210.Ctor__((_1129_v3).IsProperSubsetOf(_1129_v3), p1, (_this).F10())
			_1186_v52 = _nw210
			var _1187_v53 _dafny.Array
			_ = _1187_v53
			var _nwElement0_52 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1126_v0, _1126_v0), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1126_v0, _1126_v0)).Cardinality()))).Uint32(), _dafny.CodePoint('d'))
			_ = _nwElement0_52
			var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(3))
			_ = _nw211
			(_nw211).ArraySet1(_nwElement0_52, 0)
			(_nw211).ArraySet1(Companion_Default___.Fm2(_1186_v52.F12, (_1130_v4).Cardinality(), p0, (_1186_v52).F13(), globalState), 1)
			(_nw211).ArraySet1(_dafny.Companion_Sequence_.Update(_1126_v0, (Companion_Default___.SafeIndex((_1186_v52).F13(), _dafny.IntOfUint32((_1126_v0).Cardinality()))).Uint32(), p0), 2)
			_1187_v53 = _nw211
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1187_v53), 0))
			_ = _index213
			(_1187_v53).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xonhdkbwi"), _1126_v0), (_index213).Int())
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1187_v53), 0))
			_ = _index214
			(_1187_v53).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(187))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}(func(_1188_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			})), (_index214).Int())
			var _1189_v54 _dafny.Map
			_ = _1189_v54
			_1189_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F11()), (func() _dafny.Array {
				if (_this).F11() {
					return _1138_v10
				}
				return _1138_v10
			})())
			_1138_v10 = (func() _dafny.Array {
				if (_1189_v54).Contains(_dafny.Companion_Sequence_.Equal(_1131_v5, _1131_v5)) {
					return (_1189_v54).Get(_dafny.Companion_Sequence_.Equal(_1131_v5, _1131_v5)).(_dafny.Array)
				}
				return _1134_v8
			})()
		}
		var _1190_v55 bool
		_ = _1190_v55
		_1190_v55 = false
		var _1191_v56 _dafny.Array
		_ = _1191_v56
		var _nw212 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw212
		_1191_v56 = _nw212
		var _1192_v57 _dafny.Map
		_ = _1192_v57
		_1192_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1191_v56, (_this).F11())
		var _1193_v58 _dafny.Map
		_ = _1193_v58
		_1193_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), p0)
		var _1194_v59 _dafny.Map
		_ = _1194_v59
		_1194_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1193_v58).Cardinality()).Plus(p1), _dafny.Companion_Sequence_.Update(_1126_v0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1126_v0).Cardinality()))).Uint32(), p0))
		var _1195_v60 _dafny.Map
		_ = _1195_v60
		_1195_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_1130_v4, globalState), _1126_v0)
		var _pat_let_tv29 = p1
		_ = _pat_let_tv29
		var _rhs198 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if _1190_v55 {
				return _dafny.SeqOf(p1)
			}
			return _1132_v6
		})(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _1190_v55 {
				return _dafny.SeqOf(p1)
			}
			return _1132_v6
		})()).Cardinality()))).Uint32(), p1), (Companion_Default___.SafeIndex((_1192_v57).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if _1190_v55 {
				return _dafny.SeqOf(p1)
			}
			return _1132_v6
		})(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _1190_v55 {
				return _dafny.SeqOf(p1)
			}
			return _1132_v6
		})()).Cardinality()))).Uint32(), p1)).Cardinality()))).Uint32(), (_1132_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1132_v6).Cardinality()))).Uint32()).(_dafny.Int)), (func() _dafny.Sequence {
			if (_this).F11() {
				return _1132_v6
			}
			return _1132_v6
		})())
		_ = _rhs198
		var _rhs199 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_1131_v5, Companion_Default___.Fm33(p1, (_this).F11(), (_this).F11(), p1, globalState))
		_ = _rhs199
		var _rhs200 bool = (func(_pat_let19_0 D0) D0 {
			return func(_1196_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let20_0 _dafny.Int) D0 {
					return func(_1197_dt__update_hcf2_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_((_1196_dt__update__tmp_h0).Dtor_cf1(), _1197_dt__update_hcf2_h0)
					}(_pat_let20_0)
				}(_pat_let_tv29)
			}(_pat_let19_0)
		}(p2)).Dtor_cf1()
		_ = _rhs200
		var _rhs201 _dafny.Sequence = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_1194_v59).Contains(p1) {
				return (_1194_v59).Get(p1).(_dafny.Sequence)
			}
			return (func() _dafny.Sequence {
				if (_1195_v60).Contains((_this).F11()) {
					return (_1195_v60).Get((_this).F11()).(_dafny.Sequence)
				}
				return _1126_v0
			})()
		})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1126_v0).Cardinality())), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1194_v59).Contains(p1) {
				return (_1194_v59).Get(p1).(_dafny.Sequence)
			}
			return (func() _dafny.Sequence {
				if (_1195_v60).Contains((_this).F11()) {
					return (_1195_v60).Get((_this).F11()).(_dafny.Sequence)
				}
				return _1126_v0
			})()
		})()).Cardinality()))).Uint32(), p0)
		_ = _rhs201
		var _rhs202 bool = !((_this).Fm7(p1, globalState))
		_ = _rhs202
		_1132_v6 = _rhs198
		_1190_v55 = _rhs199
		_1190_v55 = _rhs200
		_1126_v0 = _rhs201
		_1190_v55 = _rhs202
		var _1198_v61 _dafny.Array
		_ = _1198_v61
		var _len0_38 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_38
		var _nw213 _dafny.Array
		_ = _nw213
		if _len0_38.Cmp(_dafny.Zero) == 0 {
			_nw213 = _dafny.NewArray(_len0_38)
		} else {
			var _init38 func(_dafny.Int) _dafny.Set = (func(_1199_v4 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_1200_i6 _dafny.Int) _dafny.Set {
					return _1199_v4
				}
			})(_1130_v4)
			_ = _init38
			var _element0_38 = _init38(_dafny.Zero)
			_ = _element0_38
			_nw213 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
			(_nw213).ArraySet1(_element0_38, 0)
			var _nativeLen0_38 = (_len0_38).Int()
			_ = _nativeLen0_38
			for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
				(_nw213).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
			}
		}
		_1198_v61 = _nw213
		for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1198_v61), 0))); ; {
			_guard_loop_6, _ok47 := _iter47()
			if !_ok47 {
				break
			}
			var _1201_i5 _dafny.Int
			_1201_i5 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_1201_i5).Sign() != -1) && ((_1201_i5).Cmp(_dafny.ArrayLen((_1198_v61), 0)) < 0)) {
				(_1198_v61).ArraySet1((_dafny.SetOf(p1)).Intersection((_1130_v4).Intersection(func() _dafny.Set {
					var _coll40 = _dafny.NewBuilder()
					_ = _coll40
					for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(102), _dafny.IntOfInt64(717))); ; {
						_compr_40, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1202_v62 _dafny.Int
						_1202_v62 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(102)).Cmp(_1202_v62) <= 0) && ((_1202_v62).Cmp(_dafny.IntOfInt64(717)) < 0) {
							_coll40.Add(Companion_Default___.SafeDivisionInt(_1202_v62, _dafny.IntOfUint32((_1131_v5).Cardinality())))
						}
					}
					return _coll40.ToSet()
				}())), (_1201_i5).Int())
			}
		}
		var _1203_v63 _dafny.Array
		_ = _1203_v63
		var _nw214 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
		_ = _nw214
		_1203_v63 = _nw214
		var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1203_v63), 0))
		_ = _index215
		(_1203_v63).ArraySet1(_1133_v7, (_index215).Int())
		var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1203_v63), 0))
		_ = _index216
		(_1203_v63).ArraySet1(_1133_v7, (_index216).Int())
	}
}
func (_this *C5) F11() bool {
	{
		return _this._f11
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
