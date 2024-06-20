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
func (_static *CompanionStruct_Default___) Fm3(p0 bool, globalState *GlobalState) _dafny.Int {
	return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()).Plus(_dafny.IntOfInt64(618))).Times(((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(887)
		}))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_0_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(887)
			})), _1_v0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_1_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lll")).Cardinality())), true)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-226))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D3 = (func() D3 {
		if !(!(false)) {
			return Companion_D3_.Create_DC11_(Companion_D3_.Create_DC10_(_dafny.CodePoint('y')))
		}
		return Companion_D3_.Create_DC11_(Companion_D3_.Create_DC10_(_dafny.CodePoint('u')))
	})()
	_ = _source0
	if _source0.Is_DC9() {
		var _3___mcc_h0 _dafny.Int = _source0.Get_().(D3_DC9).Cf8
		_ = _3___mcc_h0
		var _4___mcc_h1 bool = _source0.Get_().(D3_DC9).Cf9
		_ = _4___mcc_h1
		var _5___mcc_h2 _dafny.Int = _source0.Get_().(D3_DC9).Cf10
		_ = _5___mcc_h2
		var _6___mcc_h3 bool = _source0.Get_().(D3_DC9).Cf11
		_ = _6___mcc_h3
		var _7_cf11 bool = _6___mcc_h3
		_ = _7_cf11
		var _8_cf10 _dafny.Int = _5___mcc_h2
		_ = _8_cf10
		var _9_cf9 bool = _4___mcc_h1
		_ = _9_cf9
		var _10_cf8 _dafny.Int = _3___mcc_h0
		_ = _10_cf8
		return _dafny.CodePoint('o')
	} else if _source0.Is_DC10() {
		var _11___mcc_h4 _dafny.CodePoint = _source0.Get_().(D3_DC10).Cf12
		_ = _11___mcc_h4
		var _12_cf12 _dafny.CodePoint = _11___mcc_h4
		_ = _12_cf12
		return _12_cf12
	} else if _source0.Is_DC8() {
		var _13___mcc_h5 _dafny.Sequence = _source0.Get_().(D3_DC8).Cf7
		_ = _13___mcc_h5
		var _14_cf7 _dafny.Sequence = _13___mcc_h5
		_ = _14_cf7
		return _dafny.CodePoint('b')
	} else {
		var _15___mcc_h6 D3 = _source0.Get_().(D3_DC11).Cf13
		_ = _15___mcc_h6
		var _16_cf13 D3 = _15___mcc_h6
		_ = _16_cf13
		return _dafny.CodePoint('h')
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Map, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(!(true))), _dafny.SeqOf(true, true))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nxm"), _dafny.UnicodeSeqOfUtf8Bytes("uaxq"))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Map, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) D1 {
	if (_dafny.MultiSetOf(_dafny.IntOfInt64(668))).Equals(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-644), (_dafny.MultiSetOf(false, true, !(false), false)).Cardinality())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), true)).Cardinality(), (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-321), _dafny.IntOfInt64(102), _dafny.IntOfInt64(974), _dafny.IntOfInt64(250))).Cardinality()), _dafny.CodePoint('h')))).Cardinality())) {
		if true {
			return Companion_D1_.Create_DC4_(false)
		} else {
			return Companion_D1_.Create_DC4_(false)
		}
	} else if (Companion_D7_.Create_DC21_(false, _dafny.CodePoint('o'), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-439))).Uint32(), func(coer3 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_17_i0 _dafny.Int) D0 {
		return Companion_D0_.Create_DC1_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(361))).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()))
	}))).Cardinality())))).Dtor_cf30() {
		return Companion_D1_.Create_DC4_(true)
	} else {
		return Companion_D1_.Create_DC4_(false)
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(false), !(!(true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(false), true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(true), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(true), true)))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false)).Intersection(_dafny.SetOf(true, true, false, true, true))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(871), false)).Merge((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(896), _dafny.IntOfInt64(351))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(896)).Cmp(_18_v0) <= 0) && ((_18_v0).Cmp(_dafny.IntOfInt64(351)) < 0) {
				_coll1.Add((_18_v0).Times((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-24))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_19_i0 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()
				})))).Cardinality()), true)
			}
		}
		return _coll1.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(821), false)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-758), _dafny.IntOfInt64(774))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _20_v0 _dafny.Int
			_20_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(-758)).Cmp(_20_v0) <= 0) && ((_20_v0).Cmp(_dafny.IntOfInt64(774)) < 0) {
				_coll2.Add(Companion_Default___.SafeDivisionInt(_20_v0, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(531))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-896))).Cardinality()))).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fixlptlef")).Cardinality()))
			}
		}
		return _coll2.ToMap()
	}()).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(49)))).Cardinality(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(829), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(748), false)).Cardinality())).Cardinality()).Minus(_dafny.IntOfInt64(795)))
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, p2 bool, globalState *GlobalState) bool {
	if true {
		return (_dafny.MultiSetOf(true, true, !(!(true)), true, true)).IsSubsetOf(_dafny.MultiSetOf(false))
	} else {
		return ((_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(748), _dafny.IntOfInt64(26))).Cardinality())).Cmp(_dafny.IntOfInt64(761)) >= 0
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(!(false)), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(true), _dafny.SeqOf(true, !(!(false)))), false)
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Map {
	if !((Companion_D7_.Create_DC21_(true, _dafny.CodePoint('s'), _dafny.MultiSetOf(_dafny.IntOfInt64(756)))).Dtor_cf30()) {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nfbfbvbsb")).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(453), false))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(514), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jdxlfgub")).Cardinality())), false))).Merge(func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-111), _dafny.IntOfInt64(887))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _21_v0 _dafny.Int
				_21_v0 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(-111)).Cmp(_21_v0) <= 0) && ((_21_v0).Cmp(_dafny.IntOfInt64(887)) < 0) {
					_coll3.Add(Companion_Default___.SafeModuloInt(_21_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()), !(true)))
				}
			}
			return _coll3.ToMap()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uxeiish")).Cardinality()), _dafny.IntOfInt64(755))).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-863))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_22_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fgghpfsb")).Cardinality()))).Cardinality())), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(344), _dafny.IntOfInt64(194))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _23_v0 _dafny.Int
			_23_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(344)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(194)) < 0) {
				_coll4.Add((_23_v0).Times(_dafny.IntOfInt64(-608)))
			}
		}
		return _coll4.ToSet()
	}()).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true, !(false))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.CodePoint('a'), _dafny.CodePoint('u'), _dafny.CodePoint('c'))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) D1 {
	var _source1 D3 = Companion_D3_.Create_DC8_(_dafny.UnicodeSeqOfUtf8Bytes("ycfygtdb"))
	_ = _source1
	if _source1.Is_DC9() {
		var _24___mcc_h0 _dafny.Int = _source1.Get_().(D3_DC9).Cf8
		_ = _24___mcc_h0
		var _25___mcc_h1 bool = _source1.Get_().(D3_DC9).Cf9
		_ = _25___mcc_h1
		var _26___mcc_h2 _dafny.Int = _source1.Get_().(D3_DC9).Cf10
		_ = _26___mcc_h2
		var _27___mcc_h3 bool = _source1.Get_().(D3_DC9).Cf11
		_ = _27___mcc_h3
		var _28_cf11 bool = _27___mcc_h3
		_ = _28_cf11
		var _29_cf10 _dafny.Int = _26___mcc_h2
		_ = _29_cf10
		var _30_cf9 bool = _25___mcc_h1
		_ = _30_cf9
		var _31_cf8 _dafny.Int = _24___mcc_h0
		_ = _31_cf8
		return Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC4_(_30_cf9)))
	} else if _source1.Is_DC10() {
		var _32___mcc_h4 _dafny.CodePoint = _source1.Get_().(D3_DC10).Cf12
		_ = _32___mcc_h4
		var _33_cf12 _dafny.CodePoint = _32___mcc_h4
		_ = _33_cf12
		return Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(!(true))))
	} else if _source1.Is_DC8() {
		var _34___mcc_h5 _dafny.Sequence = _source1.Get_().(D3_DC8).Cf7
		_ = _34___mcc_h5
		var _35_cf7 _dafny.Sequence = _34___mcc_h5
		_ = _35_cf7
		return Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC4_(false))))
	} else {
		var _36___mcc_h6 D3 = _source1.Get_().(D3_DC11).Cf13
		_ = _36___mcc_h6
		var _37_cf13 D3 = _36___mcc_h6
		_ = _37_cf13
		return Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(!(true)))))
	}
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 bool, globalState *GlobalState) D1 {
	if !(false) || ((Companion_D7_.Create_DC21_(false, _dafny.CodePoint('f'), _dafny.MultiSetOf((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(914)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfInt64(355), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lloh")).Cardinality()), _dafny.IntOfInt64(384))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-576)))).Dtor_cf30()) {
		return Companion_D1_.Create_DC3_(!(!(false)))
	} else {
		return Companion_D1_.Create_DC3_(false)
	}
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cytq"), _dafny.UnicodeSeqOfUtf8Bytes("r")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("kcqqyceot")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("omkpbmjjq"))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.CodePoint('g'))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _dafny.CodePoint('d'))).Merge(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('x'), _dafny.CodePoint('m'))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _38_v0 _dafny.CodePoint
			_38_v0 = interface{}(_compr_5).(_dafny.CodePoint)
			if (_dafny.MultiSetOf(_dafny.CodePoint('x'), _dafny.CodePoint('m'))).Contains(_38_v0) {
				_coll5.Add(_38_v0, _dafny.CodePoint('p'))
			}
		}
		return _coll5.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) M5(p0 _dafny.Sequence, p1 _dafny.Set, globalState *GlobalState) (bool, _dafny.Map) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Map = _dafny.EmptyMap
	_ = r1
	var _39_v0 _dafny.Sequence
	_ = _39_v0
	_39_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ojcommo")
	var _40_v1 _dafny.Set
	_ = _40_v1
	_40_v1 = _dafny.SetOf(_39_v0)
	var _41_i0 _dafny.Int
	_ = _41_i0
	_41_i0 = _dafny.Zero
	{
		for (_dafny.IntOfInt64(457)).Cmp((_40_v1).Cardinality()) > 0 {
			{
				if (_41_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_41_i0 = (_41_i0).Plus(_dafny.One)
				var _42_v2 _dafny.Array
				_ = _42_v2
				var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw0
				_42_v2 = _nw0
				var _43_v3 bool
				_ = _43_v3
				_43_v3 = true
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))
				_ = _index0
				(_42_v2).ArraySet1(Companion_Default___.Fm12(_43_v3, true, _43_v3, globalState), (_index0).Int())
				var _44_v4 _dafny.Array
				_ = _44_v4
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_0
				var _nw1 _dafny.Array
				_ = _nw1
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw1 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) _dafny.Int = func(_45_i1 _dafny.Int) _dafny.Int {
						return (_45_i1).Plus(_dafny.IntOfInt64(364))
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
				_44_v4 = _nw1
				var _46_v5 _dafny.Int
				_ = _46_v5
				_46_v5 = _dafny.IntOfInt64(521)
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
				_ = _index1
				(_44_v4).ArraySet1(_46_v5, (_index1).Int())
				var _47_v6 D3
				_ = _47_v6
				_47_v6 = Companion_D3_.Create_DC8_(_39_v0)
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))
				_ = _index2
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
				_ = _index3
				var _rhs0 bool = (Companion_D1_.Create_DC3_(_43_v3)).Dtor_cf3()
				_ = _rhs0
				var _rhs1 _dafny.Int = _46_v5
				_ = _rhs1
				var _rhs2 D3 = _47_v6
				_ = _rhs2
				var _lhs0 _dafny.Array = _42_v2
				_ = _lhs0
				var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))
				_ = _lhs1
				var _lhs2 _dafny.Array = _44_v4
				_ = _lhs2
				var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
				_ = _lhs3
				(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
				(_lhs2).ArraySet1(_rhs1, (_lhs3).Int())
				_47_v6 = _rhs2
				if true {
					var _48_v7 _dafny.Array
					_ = _48_v7
					var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
					_ = _nw2
					_48_v7 = _nw2
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_48_v7), 0))
					_ = _index4
					(_48_v7).ArraySet1(_44_v4, (_index4).Int())
					var _49_v8 _dafny.Set
					_ = _49_v8
					_49_v8 = _dafny.SetOf((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), _43_v3)
					var _50_v10 _dafny.Map
					_ = _50_v10
					_50_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate((p1).Elements()); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _51_v9 _dafny.Int
							_51_v9 = interface{}(_compr_6).(_dafny.Int)
							if (p1).Contains(_51_v9) {
								_coll6.Add((_51_v9).Minus((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)), _43_v3)
							}
						}
						return _coll6.ToMap()
					}()).Cardinality(), _43_v3)
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_48_v7), 0))
					_ = _index5
					(_48_v7).ArraySet1((func() _dafny.Array {
						if (_49_v8).IsSubsetOf(_49_v8) {
							return _44_v4
						}
						return (func() _dafny.Array {
							if (func() bool {
								if (_50_v10).Contains((p0).Select((Companion_Default___.SafeIndex(_46_v5, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)) {
									return (_50_v10).Get((p0).Select((Companion_Default___.SafeIndex(_46_v5, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)).(bool)
								}
								return _43_v3
							})() {
								return _44_v4
							}
							return _44_v4
						})()
					})(), (_index5).Int())
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index6
					(_44_v4).ArraySet1((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), (_index6).Int())
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index7
					(_44_v4).ArraySet1((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), (_index7).Int())
					var _52_v11 _dafny.Map
					_ = _52_v11
					_52_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_42_v2, _46_v5)
					_52_v11 = ((_52_v11).Merge(_52_v11)).Merge(_52_v11)
					_46_v5 = Companion_Default___.SafeDivisionInt((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), Companion_Default___.Fm3(!(false), globalState))
				} else {
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index8
					(_44_v4).ArraySet1((_dafny.Zero).Minus((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)), (_index8).Int())
					var _53_v12 T0
					_ = _53_v12
					var _nw3 *C2 = New_C2_()
					_ = _nw3
					_nw3.Ctor__((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), _46_v5, (_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int))
					_53_v12 = _nw3
					var _54_v13 _dafny.Map
					_ = _54_v13
					_54_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v5, _53_v12)
					var _55_v14 _dafny.Array
					_ = _55_v14
					var _nwElement0_0 T0 = _53_v12
					_ = _nwElement0_0
					var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(28))
					_ = _nw4
					(_nw4).ArraySet1(_nwElement0_0, 0)
					(_nw4).ArraySet1(_53_v12, 1)
					(_nw4).ArraySet1(_53_v12, 2)
					(_nw4).ArraySet1(_53_v12, 3)
					(_nw4).ArraySet1(_53_v12, 4)
					(_nw4).ArraySet1(_53_v12, 5)
					(_nw4).ArraySet1(_53_v12, 6)
					(_nw4).ArraySet1(_53_v12, 7)
					(_nw4).ArraySet1(_53_v12, 8)
					(_nw4).ArraySet1(_53_v12, 9)
					(_nw4).ArraySet1(_53_v12, 10)
					(_nw4).ArraySet1(_53_v12, 11)
					(_nw4).ArraySet1(_53_v12, 12)
					(_nw4).ArraySet1(_53_v12, 13)
					(_nw4).ArraySet1((func() T0 {
						if (_54_v13).Contains(_dafny.IntOfInt64(-282)) {
							return (_54_v13).Get(_dafny.IntOfInt64(-282)).(T0)
						}
						return _53_v12
					})(), 14)
					(_nw4).ArraySet1(_53_v12, 15)
					(_nw4).ArraySet1(_53_v12, 16)
					(_nw4).ArraySet1(_53_v12, 17)
					(_nw4).ArraySet1(_53_v12, 18)
					(_nw4).ArraySet1(_53_v12, 19)
					(_nw4).ArraySet1(_53_v12, 20)
					(_nw4).ArraySet1(_53_v12, 21)
					(_nw4).ArraySet1(_53_v12, 22)
					(_nw4).ArraySet1(_53_v12, 23)
					(_nw4).ArraySet1(_53_v12, 24)
					(_nw4).ArraySet1(_53_v12, 25)
					(_nw4).ArraySet1(_53_v12, 26)
					(_nw4).ArraySet1(_53_v12, 27)
					_55_v14 = _nw4
					var _56_v15 _dafny.Map
					_ = _56_v15
					_56_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v14, (_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool))
					var _57_v16 _dafny.Map
					_ = _57_v16
					_57_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v15, (p0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int))
					_57_v16 = (_57_v16).Update((_56_v15).Merge(_56_v15), (_53_v12).F4())
					var _58_v17 _dafny.Sequence
					_ = _58_v17
					_58_v17 = _dafny.SeqOf(_39_v0)
					var _59_v18 _dafny.Array
					_ = _59_v18
					var _nwElement0_1 _dafny.Sequence = _39_v0
					_ = _nwElement0_1
					var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(12))
					_ = _nw5
					(_nw5).ArraySet1(_nwElement0_1, 0)
					(_nw5).ArraySet1((_58_v17).Select((Companion_Default___.SafeIndex((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_58_v17).Cardinality()))).Uint32()).(_dafny.Sequence), 1)
					(_nw5).ArraySet1(_39_v0, 2)
					(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_39_v0, _39_v0), 3)
					(_nw5).ArraySet1(_39_v0, 4)
					(_nw5).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-331))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg6 _dafny.Int) interface{} {
							return coer6(arg6)
						}
					}(func(_60_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('u')
					})), 5)
					(_nw5).ArraySet1(_39_v0, 6)
					(_nw5).ArraySet1(_39_v0, 7)
					(_nw5).ArraySet1(_39_v0, 8)
					(_nw5).ArraySet1(_39_v0, 9)
					(_nw5).ArraySet1(_39_v0, 10)
					(_nw5).ArraySet1(_39_v0, 11)
					_59_v18 = _nw5
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_59_v18), 0))
					_ = _index9
					(_59_v18).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("owumuar"), (_index9).Int())
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))
					_ = _index10
					var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_59_v18), 0))
					_ = _index11
					var _rhs3 bool = _43_v3
					_ = _rhs3
					var _rhs4 bool = _43_v3
					_ = _rhs4
					var _rhs5 _dafny.Int = (_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)
					_ = _rhs5
					var _rhs6 _dafny.Sequence = _39_v0
					_ = _rhs6
					var _lhs4 *GlobalState = globalState
					_ = _lhs4
					var _lhs5 _dafny.Array = _42_v2
					_ = _lhs5
					var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))
					_ = _lhs6
					var _lhs7 T0 = _53_v12
					_ = _lhs7
					var _lhs8 _dafny.Array = _59_v18
					_ = _lhs8
					var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_59_v18), 0))
					_ = _lhs9
					_lhs4.F2 = _rhs3
					(_lhs5).ArraySet1(_rhs4, (_lhs6).Int())
					_lhs7.F5_set_(_rhs5)
					(_lhs8).ArraySet1(_rhs6, (_lhs9).Int())
					var _61_v19 *C1
					_ = _61_v19
					var _nw6 *C1 = New_C1_()
					_ = _nw6
					_nw6.Ctor__(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_53_v12.F5()), _53_v12.F5()), _53_v12.F5())
					_61_v19 = _nw6
					var _62_v20 *C2
					_ = _62_v20
					var _nw7 *C2 = New_C2_()
					_ = _nw7
					_nw7.Ctor__((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_43_v3), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cbblxtt")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_43_v3)).Cardinality()))).Uint32(), _43_v3)).Cardinality())).Cmp((_53_v12).F4()) != 0, _46_v5, _53_v12.F5())
					_62_v20 = _nw7
				}
				var _63_v21 _dafny.Map
				_ = _63_v21
				_63_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_43_v3, false)).Cardinality(), _dafny.IntOfInt64(-512)), _43_v3)
				var _64_v22 _dafny.Map
				_ = _64_v22
				_64_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), globalState), _46_v5)
				if (func() bool {
					if (_63_v21).Contains(_64_v22) {
						return (_63_v21).Get(_64_v22).(bool)
					}
					return _43_v3
				})() {
					var _65_v23 _dafny.Sequence
					_ = _65_v23
					_65_v23 = _dafny.SeqOf((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), _43_v3, (_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), (_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), (_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool))
					var _66_v24 _dafny.Map
					_ = _66_v24
					_66_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm20((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), _39_v0, _65_v23, globalState), (_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool))
					var _67_v25 _dafny.Sequence
					_ = _67_v25
					_67_v25 = _dafny.SeqOf(_39_v0, _39_v0, _39_v0)
					_66_v24 = (_66_v24).Update(_67_v25, (_65_v23).Select((Companion_Default___.SafeIndex((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_65_v23).Cardinality()))).Uint32()).(bool))
					var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))
					_ = _index12
					(_42_v2).ArraySet1(false, (_index12).Int())
					var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index13
					var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index14
					var _rhs7 _dafny.Int = (_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)
					_ = _rhs7
					var _rhs8 _dafny.Int = (p0).Select((Companion_Default___.SafeIndex((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _rhs8
					var _lhs10 _dafny.Array = _44_v4
					_ = _lhs10
					var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _lhs11
					var _lhs12 _dafny.Array = _44_v4
					_ = _lhs12
					var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _lhs13
					(_lhs10).ArraySet1(_rhs7, (_lhs11).Int())
					(_lhs12).ArraySet1(_rhs8, (_lhs13).Int())
					var _68_v26 *C1
					_ = _68_v26
					var _nw8 *C1 = New_C1_()
					_ = _nw8
					_nw8.Ctor__((_46_v5).Times((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)), (_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int))
					_68_v26 = _nw8
					var _69_v27 *C3
					_ = _69_v27
					var _nw9 *C3 = New_C3_()
					_ = _nw9
					_nw9.Ctor__((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool))
					_69_v27 = _nw9
				} else {
					_39_v0 = _39_v0
					var _70_v28 _dafny.MultiSet
					_ = _70_v28
					_70_v28 = _dafny.MultiSetOf((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool))
					var _71_v29 _dafny.Sequence
					_ = _71_v29
					_71_v29 = _dafny.SeqOf((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool))
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index15
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index16
					var _rhs9 _dafny.Int = Companion_Default___.SafeDivisionInt((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_70_v28).Cardinality()))
					_ = _rhs9
					var _rhs10 bool = (_71_v29).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), (_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_71_v29).Cardinality()))).Uint32()).(bool)
					_ = _rhs10
					var _rhs11 _dafny.Int = _dafny.IntOfInt64(487)
					_ = _rhs11
					var _lhs14 _dafny.Array = _44_v4
					_ = _lhs14
					var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _lhs15
					var _lhs16 *GlobalState = globalState
					_ = _lhs16
					var _lhs17 _dafny.Array = _44_v4
					_ = _lhs17
					var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _lhs18
					(_lhs14).ArraySet1(_rhs9, (_lhs15).Int())
					_lhs16.F2 = _rhs10
					(_lhs17).ArraySet1(_rhs11, (_lhs18).Int())
					var _72_v30 _dafny.Map
					_ = _72_v30
					_72_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_43_v3, (_dafny.Zero).Minus((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)))
					var _73_v31 _dafny.Map
					_ = _73_v31
					_73_v31 = _72_v30
					_72_v30 = ((_72_v30).Update(_43_v3, _46_v5)).Merge((_73_v31))
					var _74_v32 _dafny.CodePoint
					_ = _74_v32
					_74_v32 = _dafny.CodePoint('f')
					var _75_v33 _dafny.Map
					_ = _75_v33
					_75_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(true, (_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int), Companion_Default___.Fm12((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), _43_v3, _43_v3, globalState), _dafny.IntOfUint32((_39_v0).Cardinality()), globalState), _74_v32)
					var _76_v34 _dafny.Map
					_ = _76_v34
					_76_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _43_v3)
					var _77_v35 _dafny.MultiSet
					_ = _77_v35
					_77_v35 = _dafny.MultiSetOf(_75_v33, Companion_Default___.Fm21(_46_v5, (func() bool {
						if (_76_v34).Contains(_43_v3) {
							return (_76_v34).Get(_43_v3).(bool)
						}
						return _43_v3
					})(), (_76_v34).Cardinality(), _46_v5, globalState), _75_v33, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v32, _74_v32))
					var _78_v36 *C2
					_ = _78_v36
					var _nw10 *C2 = New_C2_()
					_ = _nw10
					_nw10.Ctor__((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool), _46_v5, _46_v5)
					_78_v36 = _nw10
					var _79_v37 _dafny.Sequence
					_ = _79_v37
					_79_v37 = _dafny.SeqOf(_78_v36)
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _index17
					var _rhs12 _dafny.MultiSet = (_77_v35).Union((_77_v35).Update(_75_v33, Companion_Default___.Abs(_dafny.IntOfUint32((_79_v37).Cardinality()))))
					_ = _rhs12
					var _rhs13 _dafny.Int = (_dafny.Zero).Minus(_46_v5)
					_ = _rhs13
					var _rhs14 _dafny.Int = (_dafny.IntOfInt64(-531)).Plus(((_dafny.Zero).Minus((_dafny.Zero).Minus(_46_v5))).Plus((_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)))
					_ = _rhs14
					var _lhs19 _dafny.Array = _44_v4
					_ = _lhs19
					var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))
					_ = _lhs20
					_77_v35 = _rhs12
					_46_v5 = _rhs13
					(_lhs19).ArraySet1(_rhs14, (_lhs20).Int())
					var _80_v38 _dafny.Map
					_ = _80_v38
					_80_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D4_.Create_DC13_((_42_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_42_v2), 0))).Int()).(bool))).Dtor_cf15(), _dafny.MultiSetOf(_dafny.CodePoint('i')))
					_80_v38 = (_80_v38).Update(_43_v3, _dafny.MultiSetFromSeq(_39_v0))
				}
				_46_v5 = (_44_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_44_v4), 0))).Int()).(_dafny.Int)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	(globalState).F2 = false
	var _81_v39 bool
	_ = _81_v39
	_81_v39 = false
	r0 = _81_v39
	var _82_v40 _dafny.Array
	_ = _82_v40
	var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
	_ = _nw11
	_82_v40 = _nw11
	_82_v40 = _82_v40
	var _83_v41 _dafny.Int
	_ = _83_v41
	_83_v41 = _dafny.IntOfInt64(-496)
	var _84_v42 _dafny.Set
	_ = _84_v42
	_84_v42 = _dafny.SetOf(_81_v39, _81_v39)
	var _hi0 _dafny.Int = Companion_Default___.SafeModuloInt(_83_v41, (_84_v42).Cardinality())
	_ = _hi0
	for _85_i3 := _83_v41; _85_i3.Cmp(_hi0) < 0; _85_i3 = _85_i3.Plus(_dafny.One) {
		var _86_v43 _dafny.MultiSet
		_ = _86_v43
		_86_v43 = _dafny.MultiSetOf(_81_v39)
		var _87_v44 _dafny.Map
		_ = _87_v44
		_87_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_85_i3), _86_v43)
		var _88_v45 D6
		_ = _88_v45
		_88_v45 = Companion_D6_.Create_DC18_(_39_v0, _81_v39, (_dafny.Zero).Minus(_83_v41))
		var _89_v46 _dafny.Set
		_ = _89_v46
		_89_v46 = _dafny.SetOf((func() _dafny.MultiSet {
			if (_87_v44).Contains((_88_v45).Dtor_cf27()) {
				return (_87_v44).Get((_88_v45).Dtor_cf27()).(_dafny.MultiSet)
			}
			return _86_v43
		})())
		_89_v46 = ((_89_v46).Union(_dafny.SetOf(_86_v43))).Difference(_89_v46)
		var _90_v47 _dafny.Array
		_ = _90_v47
		var _nw12 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
		_ = _nw12
		_90_v47 = _nw12
		var _rhs15 _dafny.Array = _90_v47
		_ = _rhs15
		var _rhs16 bool = (func() bool {
			if !(_81_v39) {
				return _81_v39
			}
			return _81_v39
		})()
		_ = _rhs16
		var _lhs21 *GlobalState = globalState
		_ = _lhs21
		_90_v47 = _rhs15
		_lhs21.F2 = _rhs16
		var _91_v48 _dafny.Map
		_ = _91_v48
		_91_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v41, _39_v0)
		var _92_v49 _dafny.CodePoint
		_ = _92_v49
		_92_v49 = _dafny.CodePoint('l')
		_91_v48 = (_91_v48).Update(_85_i3, _dafny.Companion_Sequence_.Update(_39_v0, (Companion_Default___.SafeIndex(_83_v41, _dafny.IntOfUint32((_39_v0).Cardinality()))).Uint32(), _92_v49))
		var _93_v50 _dafny.Array
		_ = _93_v50
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw13
		_93_v50 = _nw13
		var _94_v51 _dafny.Sequence
		_ = _94_v51
		_94_v51 = _dafny.SeqOf(_81_v39, _81_v39)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_93_v50), 0))
		_ = _index18
		(_93_v50).ArraySet1((_dafny.IntOfUint32((_94_v51).Cardinality())).Minus(_85_i3), (_index18).Int())
		var _95_v52 _dafny.MultiSet
		_ = _95_v52
		_95_v52 = _dafny.MultiSetOf(_dafny.CodePoint('p'))
		var _96_v53 D3
		_ = _96_v53
		_96_v53 = Companion_D3_.Create_DC9_(_dafny.IntOfUint32((p0).Cardinality()), Companion_Default___.Fm12(_81_v39, _81_v39, _81_v39, globalState), (_95_v52).Cardinality(), _81_v39)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_93_v50), 0))
		_ = _index19
		(_93_v50).ArraySet1((func(_pat_let0_0 D3) D3 {
			return func(_97_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let1_0 bool) D3 {
					return func(_98_dt__update_hcf11_h0 bool) D3 {
						return Companion_D3_.Create_DC9_((_97_dt__update__tmp_h0).Dtor_cf8(), (_97_dt__update__tmp_h0).Dtor_cf9(), (_97_dt__update__tmp_h0).Dtor_cf10(), _98_dt__update_hcf11_h0)
					}(_pat_let1_0)
				}(false)
			}(_pat_let0_0)
		}(_96_v53)).Dtor_cf8(), (_index19).Int())
	}
	(globalState).F2 = _81_v39
	var _99_v54 _dafny.Map
	_ = _99_v54
	_99_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v41, !(_81_v39))
	var _100_v55 D0
	_ = _100_v55
	_100_v55 = Companion_D0_.Create_DC0_((_99_v54).Cardinality())
	var _pat_let_tv0 = _81_v39
	_ = _pat_let_tv0
	var _pat_let_tv1 = _81_v39
	_ = _pat_let_tv1
	var _pat_let_tv2 = _81_v39
	_ = _pat_let_tv2
	var _pat_let_tv3 = _83_v41
	_ = _pat_let_tv3
	r0 = func(_source2 D0) bool {
		if _source2.Is_DC1() {
			var _101___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
			_ = _101___mcc_h0
			var _102___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
			_ = _102___mcc_h1
			var _103_cf2 _dafny.Int = _102___mcc_h1
			_ = _103_cf2
			var _104_cf1 _dafny.Int = _101___mcc_h0
			_ = _104_cf1
			return _pat_let_tv0
		} else if _source2.Is_DC2() {
			return _pat_let_tv1
		} else {
			var _105___mcc_h2 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
			_ = _105___mcc_h2
			var _106_cf0 _dafny.Int = _105___mcc_h2
			_ = _106_cf0
			return _pat_let_tv2
		}
	}(func(_pat_let2_0 D0) D0 {
		return func(_107_dt__update__tmp_h1 D0) D0 {
			return func(_pat_let3_0 _dafny.Int) D0 {
				return func(_108_dt__update_hcf0_h0 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_108_dt__update_hcf0_h0)
				}(_pat_let3_0)
			}(_pat_let_tv3)
		}(_pat_let2_0)
	}(_100_v55))
	var _109_v56 _dafny.CodePoint
	_ = _109_v56
	_109_v56 = _dafny.CodePoint('e')
	var _110_v57 _dafny.Map
	_ = _110_v57
	_110_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_39_v0, _39_v0), (Companion_Default___.SafeIndex(_83_v41, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_39_v0, _39_v0)).Cardinality()))).Uint32(), _109_v56), (Companion_Default___.SafeIndex(_83_v41, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_39_v0, _39_v0), (Companion_Default___.SafeIndex(_83_v41, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_39_v0, _39_v0)).Cardinality()))).Uint32(), _109_v56)).Cardinality()))).Uint32(), _109_v56), Companion_Default___.Fm3(_81_v39, globalState))
	r1 = _110_v57
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _111_globalState *GlobalState
	_ = _111_globalState
	var _nw14 *GlobalState = New_GlobalState_()
	_ = _nw14
	_nw14.Ctor__(true, false, true)
	_111_globalState = _nw14
	var _112_v0 _dafny.CodePoint
	_ = _112_v0
	_112_v0 = _dafny.CodePoint('h')
	var _113_v1 _dafny.MultiSet
	_ = _113_v1
	_113_v1 = _dafny.MultiSetOf(_112_v0, _112_v0, _112_v0, _112_v0, _112_v0)
	var _114_v2 *C0
	_ = _114_v2
	var _nw15 *C0 = New_C0_()
	_ = _nw15
	_nw15.Ctor__(_113_v1)
	_114_v2 = _nw15
	var _115_v3 _dafny.Int
	_ = _115_v3
	_115_v3 = _dafny.IntOfInt64(295)
	_115_v3 = _dafny.IntOfInt64(-786)
	var _116_v4 _dafny.Array
	_ = _116_v4
	var _nw16 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
	_ = _nw16
	_116_v4 = _nw16
	var _117_v5 bool
	_ = _117_v5
	_117_v5 = true
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
	_ = _index20
	(_116_v4).ArraySet1(_117_v5, (_index20).Int())
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
	_ = _index21
	(_116_v4).ArraySet1((_115_v3).Cmp(_115_v3) >= 0, (_index21).Int())
	var _118_v6 _dafny.Array
	_ = _118_v6
	var _nw17 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(21))
	_ = _nw17
	_118_v6 = _nw17
	var _119_v7 _dafny.Set
	_ = _119_v7
	_119_v7 = _dafny.SetOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_118_v6), 0))
	_ = _index22
	(_118_v6).ArraySet1(Companion_D4_.Create_DC12_(_119_v7), (_index22).Int())
	var _120_v8 D4
	_ = _120_v8
	_120_v8 = Companion_D4_.Create_DC12_(_119_v7)
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_118_v6), 0))
	_ = _index23
	(_118_v6).ArraySet1(_120_v8, (_index23).Int())
	var _121_i0 _dafny.Int
	_ = _121_i0
	_121_i0 = _dafny.Zero
	{
		for false {
			{
				if (_121_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_121_i0 = (_121_i0).Plus(_dafny.One)
				var _122_v9 _dafny.Array
				_ = _122_v9
				var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_1
				var _nw18 _dafny.Array
				_ = _nw18
				if _len0_1.Cmp(_dafny.Zero) == 0 {
					_nw18 = _dafny.NewArray(_len0_1)
				} else {
					var _init1 func(_dafny.Int) _dafny.CodePoint = (func(_123_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_124_i1 _dafny.Int) _dafny.CodePoint {
							return _123_v0
						}
					})(_112_v0)
					_ = _init1
					var _element0_1 = _init1(_dafny.Zero)
					_ = _element0_1
					_nw18 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
					(_nw18).ArraySet1CodePoint(_element0_1, 0)
					var _nativeLen0_1 = (_len0_1).Int()
					_ = _nativeLen0_1
					for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
						(_nw18).ArraySet1CodePoint(_init1(_dafny.IntOf(_i0_1)), _i0_1)
					}
				}
				_122_v9 = _nw18
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_122_v9), 0))
				_ = _index24
				(_122_v9).ArraySet1CodePoint(_dafny.CodePoint('h'), (_index24).Int())
				var _125_v10 _dafny.Map
				_ = _125_v10
				_125_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v3, _115_v3)
				var _126_v11 _dafny.Sequence
				_ = _126_v11
				_126_v11 = _dafny.SeqOf((_125_v10).Cardinality())
				var _127_v12 _dafny.Sequence
				_ = _127_v12
				_127_v12 = _dafny.SeqOf(_117_v5, _117_v5)
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_122_v9), 0))
				_ = _index25
				var _rhs17 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_127_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_115_v3, _115_v3)).Cardinality()), _dafny.IntOfUint32((_127_v12).Cardinality()))).Uint32()).(bool) {
						return _dafny.CodePoint('t')
					}
					return _112_v0
				})()
				_ = _rhs17
				var _rhs18 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_126_v11, _126_v11)
				_ = _rhs18
				var _rhs19 bool = (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)
				_ = _rhs19
				var _lhs22 _dafny.Array = _122_v9
				_ = _lhs22
				var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_122_v9), 0))
				_ = _lhs23
				(_lhs22).ArraySet1CodePoint(_rhs17, (_lhs23).Int())
				_126_v11 = _rhs18
				_117_v5 = _rhs19
				var _128_v13 _dafny.Sequence
				_ = _128_v13
				_128_v13 = _dafny.UnicodeSeqOfUtf8Bytes("mpqme")
				_128_v13 = _128_v13
				var _129_v14 D2
				_ = _129_v14
				_129_v14 = Companion_D2_.Create_DC7_()
				_129_v14 = _129_v14
				var _130_v15 _dafny.Map
				_ = _130_v15
				_130_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_126_v11).Select((Companion_Default___.SafeIndex(_115_v3, _dafny.IntOfUint32((_126_v11).Cardinality()))).Uint32()).(_dafny.Int), _117_v5)
				(_111_globalState).F2 = !(!((func() _dafny.Map {
					if _117_v5 {
						return _130_v15
					}
					return _130_v15
				})()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v3, _117_v5)))
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_119_v7 = (_119_v7).Union(_119_v7)
	var _hi1 _dafny.Int = (_119_v7).Cardinality()
	_ = _hi1
	for _131_i2 := _115_v3; _131_i2.Cmp(_hi1) < 0; _131_i2 = _131_i2.Plus(_dafny.One) {
		var _132_v16 _dafny.MultiSet
		_ = _132_v16
		_132_v16 = _dafny.MultiSetOf(Companion_Default___.Fm17(_111_globalState), _113_v1)
		if !(!(((func() _dafny.Int {
			if (_132_v16).Contains((_114_v2).F7()) {
				return (_132_v16).Multiplicity((_114_v2).F7())
			}
			return _115_v3
		})()).Cmp(_131_i2) < 0)) {
			var _133_v17 _dafny.Sequence
			_ = _133_v17
			_133_v17 = _dafny.SeqOf(_131_i2, _131_i2)
			var _134_v18 _dafny.Sequence
			_ = _134_v18
			_134_v18 = _dafny.UnicodeSeqOfUtf8Bytes("nqlglyph")
			(_111_globalState).F2 = ((_133_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_134_v18).Cardinality()), _dafny.IntOfUint32((_133_v17).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_131_i2) <= 0
			var _135_v19 *C3
			_ = _135_v19
			var _nw19 *C3 = New_C3_()
			_ = _nw19
			_nw19.Ctor__((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
			_135_v19 = _nw19
			var _136_v20 D6
			_ = _136_v20
			_136_v20 = Companion_D6_.Create_DC17_(_135_v19)
			_135_v19 = (_136_v20).Dtor_cf24()
			(_111_globalState).F2 = _117_v5
			var _137_v21 _dafny.Array
			_ = _137_v21
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw20
			_137_v21 = _nw20
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_137_v21), 0))
			_ = _index26
			(_137_v21).ArraySet1(_dafny.IntOfUint32((_134_v18).Cardinality()), (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_137_v21), 0))
			_ = _index27
			(_137_v21).ArraySet1(Companion_Default___.Fm3((_115_v3).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_138_v18 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_139_i3 _dafny.Int) _dafny.Sequence {
					return _138_v18
				}
			})(_134_v18)))).Cardinality())) <= 0, _111_globalState), (_index27).Int())
			var _140_v22 _dafny.Map
			_ = _140_v22
			_140_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v0, _135_v19)
			_140_v22 = _140_v22
		} else {
			var _141_v23 _dafny.Sequence
			_ = _141_v23
			_141_v23 = _dafny.SeqOf(_dafny.IntOfInt64(867))
			var _142_v24 _dafny.Map
			_ = _142_v24
			_142_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_i2, _115_v3)
			var _143_v25 _dafny.Map
			_ = _143_v25
			_143_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_141_v23).Select((Companion_Default___.SafeIndex((_142_v24).Cardinality(), _dafny.IntOfUint32((_141_v23).Cardinality()))).Uint32()).(_dafny.Int), _116_v4)
			var _144_v26 _dafny.Sequence
			_ = _144_v26
			_144_v26 = _dafny.SeqOf(_116_v4)
			_116_v4 = (func() _dafny.Array {
				if (_143_v25).Contains(_131_i2) {
					return (_143_v25).Get(_131_i2).(_dafny.Array)
				}
				return (_144_v26).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm3((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _111_globalState), _dafny.IntOfUint32((_144_v26).Cardinality()))).Uint32()).(_dafny.Array)
			})()
			var _145_v27 _dafny.Map
			_ = _145_v27
			_145_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v5, _112_v0)
			_145_v27 = (_145_v27).Update(!((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)), _112_v0)
			var _146_v28 _dafny.Array
			_ = _146_v28
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw21
			_146_v28 = _nw21
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_146_v28), 0))
			_ = _index28
			(_146_v28).ArraySet1((_115_v3).Plus(_131_i2), (_index28).Int())
			var _147_v29 _dafny.Sequence
			_ = _147_v29
			_147_v29 = _dafny.UnicodeSeqOfUtf8Bytes("oasbjj")
			var _148_v30 _dafny.Map
			_ = _148_v30
			_148_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_147_v29, _dafny.IntOfInt64(692))
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_146_v28), 0))
			_ = _index29
			(_146_v28).ArraySet1((func() _dafny.Int {
				if (_148_v30).Contains(_147_v29) {
					return (_148_v30).Get(_147_v29).(_dafny.Int)
				}
				return _131_i2
			})(), (_index29).Int())
			var _rhs20 _dafny.CodePoint = (_147_v29).Select((Companion_Default___.SafeIndex(_115_v3, _dafny.IntOfUint32((_147_v29).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs20
			var _rhs21 _dafny.Int = _131_i2
			_ = _rhs21
			var _rhs22 _dafny.Int = ((_dafny.Zero).Minus(_131_i2)).Minus(_131_i2)
			_ = _rhs22
			_112_v0 = _rhs20
			_115_v3 = _rhs21
			_115_v3 = _rhs22
			(_111_globalState).F2 = (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)
		}
		var _149_v31 _dafny.Sequence
		_ = _149_v31
		_149_v31 = _dafny.UnicodeSeqOfUtf8Bytes("tvpces")
		var _150_v32 D6
		_ = _150_v32
		_150_v32 = Companion_D6_.Create_DC18_(_149_v31, _117_v5, _131_i2)
		_150_v32 = _150_v32
		var _151_v33 D3
		_ = _151_v33
		_151_v33 = Companion_D3_.Create_DC8_(_149_v31)
		_149_v31 = _dafny.Companion_Sequence_.Concatenate((_151_v33).Dtor_cf7(), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm6((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v5, _131_i2)).Cardinality(), (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _111_globalState), _149_v31))
		var _152_v34 _dafny.Array
		_ = _152_v34
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_2
		var _nw22 _dafny.Array
		_ = _nw22
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw22 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) D2 = (func(_153_v4 _dafny.Array) func(_dafny.Int) D2 {
				return func(_154_i4 _dafny.Int) D2 {
					return Companion_D2_.Create_DC6_(_dafny.SeqOf((_153_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_153_v4), 0))).Int()).(bool)))
				}
			})(_116_v4)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw22 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw22).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw22).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_152_v34 = _nw22
		var _155_v35 _dafny.Map
		_ = _155_v35
		_155_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v34, !(_117_v5))
		var _156_v36 _dafny.Array
		_ = _156_v36
		var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw23
		_156_v36 = _nw23
		var _157_v37 _dafny.Map
		_ = _157_v37
		_157_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v36, _155_v35)
		_155_v35 = (func() _dafny.Map {
			if (_157_v37).Contains(_156_v36) {
				return (_157_v37).Get(_156_v36).(_dafny.Map)
			}
			return _155_v35
		})()
	}
	_115_v3 = _115_v3
	_115_v3 = _115_v3
	var _158_v38 _dafny.Sequence
	_ = _158_v38
	_158_v38 = _dafny.SeqOf(_115_v3)
	var _159_v39 _dafny.Array
	_ = _159_v39
	var _nwElement0_2 _dafny.Int = _115_v3
	_ = _nwElement0_2
	var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(5))
	_ = _nw24
	(_nw24).ArraySet1(_nwElement0_2, 0)
	(_nw24).ArraySet1(_dafny.IntOfInt64(134), 1)
	(_nw24).ArraySet1(_115_v3, 2)
	(_nw24).ArraySet1((_dafny.Zero).Minus(_115_v3), 3)
	(_nw24).ArraySet1(_dafny.IntOfUint32((_158_v38).Cardinality()), 4)
	_159_v39 = _nw24
	var _160_v40 _dafny.Sequence
	_ = _160_v40
	_160_v40 = _dafny.SeqOf(_159_v39)
	_160_v40 = _160_v40
	var _161_v41 _dafny.Sequence
	_ = _161_v41
	_161_v41 = _dafny.UnicodeSeqOfUtf8Bytes("oayxq")
	var _162_v42 D0
	_ = _162_v42
	_162_v42 = Companion_D0_.Create_DC2_()
	var _pat_let_tv4 = _161_v41
	_ = _pat_let_tv4
	var _pat_let_tv5 = _161_v41
	_ = _pat_let_tv5
	var _pat_let_tv6 = _161_v41
	_ = _pat_let_tv6
	var _pat_let_tv7 = _161_v41
	_ = _pat_let_tv7
	_161_v41 = func(_source3 D0) _dafny.Sequence {
		if _source3.Is_DC1() {
			var _163___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC1).Cf1
			_ = _163___mcc_h0
			var _164___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
			_ = _164___mcc_h1
			var _165_cf2 _dafny.Int = _164___mcc_h1
			_ = _165_cf2
			var _166_cf1 _dafny.Int = _163___mcc_h0
			_ = _166_cf1
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("n"), _pat_let_tv4)
		} else if _source3.Is_DC2() {
			return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv5, _pat_let_tv6)
		} else {
			var _167___mcc_h2 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
			_ = _167___mcc_h2
			var _168_cf0 _dafny.Int = _167___mcc_h2
			_ = _168_cf0
			return _pat_let_tv7
		}
	}(_162_v42)
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
	_ = _index30
	(_116_v4).ArraySet1(!((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)), (_index30).Int())
	if _117_v5 {
		var _169_v43 _dafny.Map
		_ = _169_v43
		_169_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v3, _159_v39)
		_159_v39 = (func() _dafny.Array {
			if (_169_v43).Contains(Companion_Default___.SafeDivisionInt(_115_v3, _115_v3)) {
				return (_169_v43).Get(Companion_Default___.SafeDivisionInt(_115_v3, _115_v3)).(_dafny.Array)
			}
			return _159_v39
		})()
		var _170_v44 *C3
		_ = _170_v44
		var _nw25 *C3 = New_C3_()
		_ = _nw25
		_nw25.Ctor__((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
		_170_v44 = _nw25
		(_111_globalState).F2 = (_170_v44).F3()
		var _171_v45 _dafny.Sequence
		_ = _171_v45
		_171_v45 = _dafny.SeqOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), true)
		var _172_v46 _dafny.Array
		_ = _172_v46
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_3
		var _nw26 _dafny.Array
		_ = _nw26
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw26 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = (func(_173_v38 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_174_i5 _dafny.Int) _dafny.Sequence {
					return _173_v38
				}
			})(_158_v38)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw26 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw26).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw26).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_172_v46 = _nw26
		var _rhs23 _dafny.Sequence = _171_v45
		_ = _rhs23
		var _rhs24 _dafny.Array = _172_v46
		_ = _rhs24
		_171_v45 = _rhs23
		_172_v46 = _rhs24
		var _175_v47 D1
		_ = _175_v47
		_175_v47 = Companion_D1_.Create_DC4_(true)
		var _176_v48 D1
		_ = _176_v48
		_176_v48 = Companion_D1_.Create_DC5_(_175_v47)
		var _177_v49 D1
		_ = _177_v49
		_177_v49 = Companion_D1_.Create_DC5_(_175_v47)
		var _178_v50 D1
		_ = _178_v50
		_178_v50 = Companion_D1_.Create_DC5_(_176_v48)
		var _179_v51 *C1
		_ = _179_v51
		var _nw27 *C1 = New_C1_()
		_ = _nw27
		_nw27.Ctor__(_115_v3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v3, _117_v5)).Cardinality())
		_179_v51 = _nw27
		var _180_v52 _dafny.Sequence
		_ = _180_v52
		_180_v52 = _dafny.SeqOf(_179_v51, _179_v51, _179_v51, _179_v51)
		_178_v50 = Companion_Default___.Fm18(_dafny.IntOfUint32((_180_v52).Cardinality()), _111_globalState)
	} else {
		var _181_v53 _dafny.Array
		_ = _181_v53
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
		_ = _nw28
		_181_v53 = _nw28
		var _182_v54 T0
		_ = _182_v54
		var _nw29 *C2 = New_C2_()
		_ = _nw29
		_nw29.Ctor__(true, _115_v3, _115_v3)
		_182_v54 = _nw29
		var _183_v55 _dafny.Map
		_ = _183_v55
		_183_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v54, true)
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_181_v53), 0))
		_ = _index31
		(_181_v53).ArraySet1(_183_v55, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_181_v53), 0))
		_ = _index32
		(_181_v53).ArraySet1(_183_v55, (_index32).Int())
		_161_v41 = _161_v41
		var _184_v56 _dafny.Map
		_ = _184_v56
		_184_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_161_v41, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-346))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_185_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_186_i6 _dafny.Int) _dafny.CodePoint {
				return _185_v0
			}
		})(_112_v0)))), _117_v5)
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
		_ = _index33
		(_116_v4).ArraySet1((func() bool {
			if (_184_v56).Contains((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)) {
				return (_184_v56).Get((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)).(bool)
			}
			return _117_v5
		})(), (_index33).Int())
		var _187_v57 _dafny.Sequence
		_ = _187_v57
		_187_v57 = _dafny.SeqOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _117_v5, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _117_v5, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
		var _188_v58 _dafny.Map
		_ = _188_v58
		_188_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v7, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
		if (func() bool {
			if (_187_v57).Select((Companion_Default___.SafeIndex((_182_v54).F4(), _dafny.IntOfUint32((_187_v57).Cardinality()))).Uint32()).(bool) {
				return true
			}
			return (func() bool {
				if (_188_v58).Contains(Companion_Default___.Fm9((_182_v54).F4(), _117_v5, _dafny.IntOfUint32((_161_v41).Cardinality()), _111_globalState)) {
					return (_188_v58).Get(Companion_Default___.Fm9((_182_v54).F4(), _117_v5, _dafny.IntOfUint32((_161_v41).Cardinality()), _111_globalState)).(bool)
				}
				return _117_v5
			})()
		})() {
			var _189_v59 _dafny.MultiSet
			_ = _189_v59
			_189_v59 = _dafny.MultiSetOf(_159_v39, _159_v39, _159_v39, _159_v39)
			var _190_v60 _dafny.Sequence
			_ = _190_v60
			_190_v60 = _dafny.SeqOf(_189_v59, _189_v59, _189_v59)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
			_ = _index34
			var _rhs25 _dafny.CodePoint = _112_v0
			_ = _rhs25
			var _rhs26 _dafny.MultiSet = (_190_v60).Select((Companion_Default___.SafeIndex(((_182_v54).F4()).Plus((_182_v54).F4()), _dafny.IntOfUint32((_190_v60).Cardinality()))).Uint32()).(_dafny.MultiSet)
			_ = _rhs26
			var _rhs27 bool = !(_117_v5) || ((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
			_ = _rhs27
			var _rhs28 bool = _117_v5
			_ = _rhs28
			var _rhs29 _dafny.Int = _182_v54.F5()
			_ = _rhs29
			var _lhs24 _dafny.Array = _116_v4
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
			_ = _lhs25
			var _lhs26 T0 = _182_v54
			_ = _lhs26
			_112_v0 = _rhs25
			_189_v59 = _rhs26
			_117_v5 = _rhs27
			(_lhs24).ArraySet1(_rhs28, (_lhs25).Int())
			_lhs26.F5_set_(_rhs29)
			var _191_v61 *C3
			_ = _191_v61
			var _nw30 *C3 = New_C3_()
			_ = _nw30
			_nw30.Ctor__((_187_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_187_v57).Cardinality()), _dafny.IntOfUint32((_187_v57).Cardinality()))).Uint32()).(bool))
			_191_v61 = _nw30
			_191_v61 = _191_v61
			_115_v3 = (_dafny.Zero).Minus(_115_v3)
			var _192_v62 _dafny.Map
			_ = _192_v62
			_192_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_191_v61).F3(), _dafny.IntOfInt64(372))
			(_182_v54).F5_set_((func() _dafny.Int {
				if (_192_v62).Contains(_117_v5) {
					return (_192_v62).Get(_117_v5).(_dafny.Int)
				}
				return (_182_v54).F4()
			})())
			var _193_v63 _dafny.MultiSet
			_ = _193_v63
			_193_v63 = _dafny.MultiSetOf(_158_v38, _dafny.Companion_Sequence_.Concatenate(_158_v38, _158_v38), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_194_v54 T0, _195_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_196_i7 _dafny.Int) _dafny.Int {
					return (Companion_D0_.Create_DC1_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_194_v54).F4(), _195_v3)).Cardinality(), _194_v54.F5())).Dtor_cf2()
				}
			})(_182_v54, _115_v3))))
			(_182_v54).F5_set_((func() _dafny.Int {
				if (_193_v63).Contains((func() _dafny.Sequence {
					if (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool) {
						return _158_v38
					}
					return _dafny.SeqOf((_182_v54).F4(), _115_v3)
				})()) {
					return (_193_v63).Multiplicity((func() _dafny.Sequence {
						if (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool) {
							return _158_v38
						}
						return _dafny.SeqOf((_182_v54).F4(), _115_v3)
					})())
				}
				return _182_v54.F5()
			})())
		} else {
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
			_ = _index35
			(_116_v4).ArraySet1(Companion_Default___.Fm12(false, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _117_v5, _111_globalState), (_index35).Int())
			(_182_v54).F5_set_(_dafny.IntOfInt64(605))
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_159_v39), 0))
			_ = _index36
			(_159_v39).ArraySet1((_182_v54).F4(), (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_159_v39), 0))
			_ = _index37
			(_159_v39).ArraySet1(_182_v54.F5(), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_159_v39), 0))
			_ = _index38
			(_159_v39).ArraySet1((_115_v3).Plus(_115_v3), (_index38).Int())
			var _197_v64 _dafny.Array
			_ = _197_v64
			var _nwElement0_3 _dafny.Int = _115_v3
			_ = _nwElement0_3
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(5))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_3, 0)
			(_nw31).ArraySet1(_dafny.IntOfUint32((_158_v38).Cardinality()), 1)
			(_nw31).ArraySet1(_dafny.IntOfInt64(745), 2)
			(_nw31).ArraySet1((_182_v54).F4(), 3)
			(_nw31).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(728), (_182_v54).F4()), 4)
			_197_v64 = _nw31
			var _198_v65 _dafny.Array
			_ = _198_v65
			var _nw32 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
			_ = _nw32
			_198_v65 = _nw32
			var _199_v66 _dafny.Set
			_ = _199_v66
			_199_v66 = _dafny.SetOf(_198_v65)
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_159_v39), 0))
			_ = _index39
			var _rhs30 _dafny.Array = _197_v64
			_ = _rhs30
			var _rhs31 _dafny.Int = _115_v3
			_ = _rhs31
			var _rhs32 bool = !((_dafny.SetOf(_198_v65)).IsSubsetOf(_199_v66))
			_ = _rhs32
			var _lhs27 _dafny.Array = _159_v39
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_159_v39), 0))
			_ = _lhs28
			var _lhs29 *GlobalState = _111_globalState
			_ = _lhs29
			_197_v64 = _rhs30
			(_lhs27).ArraySet1(_rhs31, (_lhs28).Int())
			_lhs29.F2 = _rhs32
		}
		(_182_v54).F5_set_(_182_v54.F5())
	}
	var _source4 D0 = _162_v42
	_ = _source4
	if _source4.Is_DC1() {
		var _200___mcc_h3 _dafny.Int = _source4.Get_().(D0_DC1).Cf1
		_ = _200___mcc_h3
		var _201___mcc_h4 _dafny.Int = _source4.Get_().(D0_DC1).Cf2
		_ = _201___mcc_h4
		var _202_cf2 _dafny.Int = _201___mcc_h4
		_ = _202_cf2
		var _203_cf1 _dafny.Int = _200___mcc_h3
		_ = _203_cf1
		var _204_v67 _dafny.Map
		_ = _204_v67
		_204_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_117_v5), Companion_Default___.Fm12((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), !(_117_v5), _117_v5, _111_globalState)))
		var _205_v68 *C2
		_ = _205_v68
		var _nw33 *C2 = New_C2_()
		_ = _nw33
		_nw33.Ctor__(_117_v5, _202_cf2, _202_cf2)
		_205_v68 = _nw33
		var _206_v69 _dafny.Map
		_ = _206_v69
		_206_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v3, _205_v68)
		_204_v67 = (_204_v67).Update(!(!((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))), !((_dafny.IntOfUint32((_158_v38).Cardinality())).Cmp((_206_v69).Cardinality()) >= 0))
		_202_cf2 = _202_cf2
		var _207_v70 *C1
		_ = _207_v70
		var _nw34 *C1 = New_C1_()
		_ = _nw34
		_nw34.Ctor__(Companion_Default___.Fm3((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _111_globalState), _115_v3)
		_207_v70 = _nw34
		_207_v70 = _207_v70
		var _208_v71 _dafny.MultiSet
		_ = _208_v71
		_208_v71 = _dafny.MultiSetOf(_117_v5)
		var _209_v72 D6
		_ = _209_v72
		_209_v72 = Companion_D6_.Create_DC18_(_161_v41, _117_v5, ((_208_v71).Update(_117_v5, Companion_Default___.Abs(_115_v3))).Cardinality())
		var _210_v73 D6
		_ = _210_v73
		_210_v73 = Companion_D6_.Create_DC19_(_209_v72)
		_210_v73 = _210_v73
	} else if _source4.Is_DC2() {
		var _211_v74 *C1
		_ = _211_v74
		var _nw35 *C1 = New_C1_()
		_ = _nw35
		_nw35.Ctor__((_dafny.Zero).Minus(_115_v3), _dafny.IntOfInt64(933))
		_211_v74 = _nw35
		_211_v74 = _211_v74
		_115_v3 = (_115_v3).Plus(_115_v3)
		var _212_v75 _dafny.Array
		_ = _212_v75
		var _nw36 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw36
		_212_v75 = _nw36
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_212_v75), 0))
		_ = _index40
		(_212_v75).ArraySet1(_114_v2, (_index40).Int())
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_212_v75), 0))
		_ = _index41
		(_212_v75).ArraySet1(_114_v2, (_index41).Int())
		var _213_v76 _dafny.Array
		_ = _213_v76
		var _nw37 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
		_ = _nw37
		_213_v76 = _nw37
		_213_v76 = _213_v76
	} else {
		var _214___mcc_h5 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
		_ = _214___mcc_h5
		var _215_cf0 _dafny.Int = _214___mcc_h5
		_ = _215_cf0
		var _216_v77 _dafny.Map
		_ = _216_v77
		_216_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v3, _115_v3)
		var _217_v78 _dafny.Map
		_ = _217_v78
		_217_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_216_v77).Cardinality(), _115_v3)
		var _218_v79 D0
		_ = _218_v79
		_218_v79 = Companion_D0_.Create_DC0_((_217_v78).Cardinality())
		var _219_v80 _dafny.Sequence
		_ = _219_v80
		_219_v80 = _dafny.SeqOf(_218_v79)
		_219_v80 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_219_v80, _219_v80), (Companion_Default___.SafeIndex((_215_cf0).Plus(_215_cf0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_219_v80, _219_v80)).Cardinality()))).Uint32(), _218_v79)
		var _220_v81 T0
		_ = _220_v81
		var _nw38 *C1 = New_C1_()
		_ = _nw38
		_nw38.Ctor__((_dafny.Zero).Minus(Companion_Default___.Fm3((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _111_globalState)), (_119_v7).Cardinality())
		_220_v81 = _nw38
		var _221_v82 _dafny.Sequence
		_ = _221_v82
		_221_v82 = _dafny.SeqOf(_220_v81)
		var _222_v83 _dafny.Sequence
		_ = _222_v83
		_222_v83 = _dafny.SeqOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _117_v5)
		var _223_v84 _dafny.Set
		_ = _223_v84
		_223_v84 = _dafny.SetOf(_222_v83)
		var _224_v85 _dafny.Map
		_ = _224_v85
		_224_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(629), (_dafny.Zero).Minus(_115_v3)), (_221_v82).Select((Companion_Default___.SafeIndex((_223_v84).Cardinality(), _dafny.IntOfUint32((_221_v82).Cardinality()))).Uint32()).(T0))
		_224_v85 = ((_224_v85).Merge(_224_v85)).Merge(_224_v85)
		(_111_globalState).F2 = false
		var _225_v86 bool
		_ = _225_v86
		var _226_v87 bool
		_ = _226_v87
		var _227_v88 bool
		_ = _227_v88
		var _out0 bool
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 bool
		_ = _out2
		_out0, _out1, _out2 = (_220_v81).M1((func() _dafny.Array {
			if (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool) {
				return _159_v39
			}
			return _159_v39
		})(), _dafny.CodePoint('u'), _111_globalState)
		_225_v86 = _out0
		_226_v87 = _out1
		_227_v88 = _out2
	}
	var _source5 D3 = Companion_D3_.Create_DC8_(_161_v41)
	_ = _source5
	if _source5.Is_DC9() {
		var _228___mcc_h6 _dafny.Int = _source5.Get_().(D3_DC9).Cf8
		_ = _228___mcc_h6
		var _229___mcc_h7 bool = _source5.Get_().(D3_DC9).Cf9
		_ = _229___mcc_h7
		var _230___mcc_h8 _dafny.Int = _source5.Get_().(D3_DC9).Cf10
		_ = _230___mcc_h8
		var _231___mcc_h9 bool = _source5.Get_().(D3_DC9).Cf11
		_ = _231___mcc_h9
		var _232_cf11 bool = _231___mcc_h9
		_ = _232_cf11
		var _233_cf10 _dafny.Int = _230___mcc_h8
		_ = _233_cf10
		var _234_cf9 bool = _229___mcc_h7
		_ = _234_cf9
		var _235_cf8 _dafny.Int = _228___mcc_h6
		_ = _235_cf8
		_234_cf9 = (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)
		var _236_v89 _dafny.MultiSet
		_ = _236_v89
		_236_v89 = _dafny.MultiSetOf(_235_cf8, _dafny.IntOfInt64(-952))
		var _237_v90 _dafny.Set
		_ = _237_v90
		_237_v90 = _dafny.SetOf((func() _dafny.Int {
			if (_236_v89).Contains(_235_cf8) {
				return (_236_v89).Multiplicity(_235_cf8)
			}
			return _115_v3
		})(), _233_cf10, _dafny.IntOfInt64(-994), _115_v3)
		var _238_v91 bool
		_ = _238_v91
		var _239_v92 _dafny.Map
		_ = _239_v92
		var _out3 bool
		_ = _out3
		var _out4 _dafny.Map
		_ = _out4
		_out3, _out4 = Companion_Default___.M5(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_233_cf10), _158_v38), _237_v90, _111_globalState)
		_238_v91 = _out3
		_239_v92 = _out4
		var _240_v94 bool
		_ = _240_v94
		var _241_v95 _dafny.Map
		_ = _241_v95
		var _out5 bool
		_ = _out5
		var _out6 _dafny.Map
		_ = _out6
		_out5, _out6 = Companion_Default___.M5(_158_v38, func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-340), _dafny.IntOfInt64(994))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _242_v93 _dafny.Int
				_242_v93 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(-340)).Cmp(_242_v93) <= 0) && ((_242_v93).Cmp(_dafny.IntOfInt64(994)) < 0) {
					_coll7.Add((_242_v93).Minus(_235_cf8))
				}
			}
			return _coll7.ToSet()
		}(), _111_globalState)
		_240_v94 = _out5
		_241_v95 = _out6
		if !(_238_v91) {
			var _243_v96 bool
			_ = _243_v96
			var _244_v97 _dafny.Map
			_ = _244_v97
			var _out7 bool
			_ = _out7
			var _out8 _dafny.Map
			_ = _out8
			_out7, _out8 = Companion_Default___.M5(_158_v38, _237_v90, _111_globalState)
			_243_v96 = _out7
			_244_v97 = _out8
			var _rhs33 _dafny.Sequence = _dafny.SeqOf((func() _dafny.Array {
				if _234_cf9 {
					return _159_v39
				}
				return _159_v39
			})(), _159_v39, _159_v39, _159_v39)
			_ = _rhs33
			var _rhs34 _dafny.Int = _115_v3
			_ = _rhs34
			_160_v40 = _rhs33
			_235_cf8 = _rhs34
			(_111_globalState).F2 = !(_232_cf11) || (_117_v5)
			var _245_v98 D1
			_ = _245_v98
			_245_v98 = Companion_D1_.Create_DC3_(_234_cf9)
			var _pat_let_tv8 = _116_v4
			_ = _pat_let_tv8
			var _pat_let_tv9 = _116_v4
			_ = _pat_let_tv9
			var _246_v99 _dafny.MultiSet
			_ = _246_v99
			_246_v99 = _dafny.MultiSetOf(Companion_Default___.Fm19(_234_cf9, !((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)), _111_globalState), _245_v98, _245_v98, _245_v98, func(_pat_let4_0 D1) D1 {
				return func(_247_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let5_0 bool) D1 {
						return func(_248_dt__update_hcf3_h0 bool) D1 {
							return Companion_D1_.Create_DC3_(_248_dt__update_hcf3_h0)
						}(_pat_let5_0)
					}((_pat_let_tv9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_pat_let_tv8), 0))).Int()).(bool))
				}(_pat_let4_0)
			}(_245_v98))
			var _249_v100 _dafny.Array
			_ = _249_v100
			var _nwElement0_4 _dafny.MultiSet = _246_v99
			_ = _nwElement0_4
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(6))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_4, 0)
			(_nw39).ArraySet1(_246_v99, 1)
			(_nw39).ArraySet1((_246_v99).Intersection((_dafny.MultiSetOf(_245_v98)).Update(Companion_D1_.Create_DC3_(_240_v94), Companion_Default___.Abs(_233_cf10))), 2)
			(_nw39).ArraySet1(_246_v99, 3)
			(_nw39).ArraySet1(_246_v99, 4)
			(_nw39).ArraySet1(_246_v99, 5)
			_249_v100 = _nw39
			_249_v100 = _249_v100
			var _250_v101 D7
			_ = _250_v101
			_250_v101 = Companion_D7_.Create_DC20_(_158_v38)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
			_ = _index42
			var _rhs35 bool = _dafny.Companion_Sequence_.IsPrefixOf((_250_v101).Dtor_cf29(), _158_v38)
			_ = _rhs35
			var _rhs36 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm6(_117_v5, _233_cf10, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _111_globalState), _dafny.UnicodeSeqOfUtf8Bytes("ygn"))
			_ = _rhs36
			var _rhs37 bool = !((_233_cf10).Cmp(_115_v3) > 0)
			_ = _rhs37
			var _lhs30 *GlobalState = _111_globalState
			_ = _lhs30
			var _lhs31 _dafny.Array = _116_v4
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
			_ = _lhs32
			_lhs30.F2 = _rhs35
			_161_v41 = _rhs36
			(_lhs31).ArraySet1(_rhs37, (_lhs32).Int())
		} else {
			var _251_v102 _dafny.Sequence
			_ = _251_v102
			_251_v102 = _dafny.SeqOf(_161_v41)
			var _252_v103 _dafny.MultiSet
			_ = _252_v103
			_252_v103 = _dafny.MultiSetOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
			var _253_v104 D6
			_ = _253_v104
			_253_v104 = Companion_D6_.Create_DC18_(_161_v41, true, (_252_v103).Cardinality())
			var _254_v105 D6
			_ = _254_v105
			_254_v105 = Companion_D6_.Create_DC18_(_dafny.UnicodeSeqOfUtf8Bytes("q"), _238_v91, (_253_v104).Dtor_cf27())
			var _255_v106 _dafny.Sequence
			_ = _255_v106
			_255_v106 = _dafny.SeqOf(_161_v41, (_251_v102).Select((Companion_Default___.SafeIndex(_115_v3, _dafny.IntOfUint32((_251_v102).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate((_253_v104).Dtor_cf25(), _161_v41))
			_251_v102 = _255_v106
			var _256_v107 D0
			_ = _256_v107
			_256_v107 = Companion_D0_.Create_DC0_(_115_v3)
			_235_cf8 = (_235_cf8).Plus((_256_v107).Dtor_cf0())
			var _257_v108 *C3
			_ = _257_v108
			var _nw40 *C3 = New_C3_()
			_ = _nw40
			_nw40.Ctor__(_232_cf11)
			_257_v108 = _nw40
			var _258_v109 T0
			_ = _258_v109
			var _nw41 *C2 = New_C2_()
			_ = _nw41
			_nw41.Ctor__(true, (_236_v89).Cardinality(), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_161_v41).Cardinality()), _233_cf10))
			_258_v109 = _nw41
			var _259_v110 *C2
			_ = _259_v110
			var _nw42 *C2 = New_C2_()
			_ = _nw42
			_nw42.Ctor__((_257_v108).F3(), ((_258_v109).F4()).Times(_233_cf10), _dafny.IntOfInt64(685))
			_259_v110 = _nw42
			var _260_v111 _dafny.Map
			_ = _260_v111
			_260_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_232_cf11, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
			var _rhs38 bool = (_260_v111).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v5, _240_v94)).Merge(_260_v111))
			_ = _rhs38
			var _rhs39 *C3 = (func() *C3 {
				if _238_v91 {
					return _257_v108
				}
				return _257_v108
			})()
			_ = _rhs39
			var _rhs40 T0 = _258_v109
			_ = _rhs40
			var _rhs41 *C2 = _259_v110
			_ = _rhs41
			var _lhs33 *GlobalState = _111_globalState
			_ = _lhs33
			_lhs33.F2 = _rhs38
			_257_v108 = _rhs39
			_258_v109 = _rhs40
			_259_v110 = _rhs41
			_260_v111 = Companion_Default___.Fm14(_dafny.IntOfInt64(543), _111_globalState)
			_234_cf9 = (_257_v108).F3()
		}
	} else if _source5.Is_DC10() {
		var _261___mcc_h10 _dafny.CodePoint = _source5.Get_().(D3_DC10).Cf12
		_ = _261___mcc_h10
		var _262_cf12 _dafny.CodePoint = _261___mcc_h10
		_ = _262_cf12
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
		_ = _index43
		(_116_v4).ArraySet1(!((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)), (_index43).Int())
		var _263_v112 *C2
		_ = _263_v112
		var _nw43 *C2 = New_C2_()
		_ = _nw43
		_nw43.Ctor__((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _115_v3, _115_v3)
		_263_v112 = _nw43
		var _264_v113 *C3
		_ = _264_v113
		var _nw44 *C3 = New_C3_()
		_ = _nw44
		_nw44.Ctor__((_263_v112).F6())
		_264_v113 = _nw44
		var _265_v114 _dafny.Map
		_ = _265_v114
		_265_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v112).F6(), _264_v113)
		var _266_v115 _dafny.Map
		_ = _266_v115
		_266_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-744), (_263_v112).F6())
		var _267_v116 _dafny.Map
		_ = _267_v116
		_267_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C3 {
			if (_265_v114).Contains(_117_v5) {
				return (_265_v114).Get(_117_v5).(*C3)
			}
			return _264_v113
		})(), _266_v115)
		var _268_v118 _dafny.Sequence
		_ = _268_v118
		_268_v118 = _dafny.SeqOf(_158_v38)
		var _269_v119 _dafny.Sequence
		_ = _269_v119
		_269_v119 = _dafny.SeqOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), false, !(_117_v5), (_263_v112).F6())
		_267_v116 = (_267_v116).Update(_264_v113, func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate(((_268_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_269_v119).Cardinality()), _dafny.IntOfUint32((_268_v118).Cardinality()))).Uint32()).(_dafny.Sequence)).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _270_v117 _dafny.Int
				_270_v117 = interface{}(_compr_8).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains((_268_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_269_v119).Cardinality()), _dafny.IntOfUint32((_268_v118).Cardinality()))).Uint32()).(_dafny.Sequence), _270_v117) {
					_coll8.Add((_270_v117).Minus(_dafny.IntOfUint32((_161_v41).Cardinality())), (_263_v112).F6())
				}
			}
			return _coll8.ToMap()
		}())
		_115_v3 = Companion_Default___.SafeModuloInt(_115_v3, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_158_v38).Cardinality()), _115_v3))
	} else if _source5.Is_DC8() {
		var _271___mcc_h11 _dafny.Sequence = _source5.Get_().(D3_DC8).Cf7
		_ = _271___mcc_h11
		var _272_cf7 _dafny.Sequence = _271___mcc_h11
		_ = _272_cf7
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_159_v39), 0))
		_ = _index44
		(_159_v39).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-902))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}((func(_273_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_274_i8 _dafny.Int) _dafny.Int {
				return _273_v3
			}
		})(_115_v3)))).Cardinality()), (_index44).Int())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_159_v39), 0))
		_ = _index45
		(_159_v39).ArraySet1(Companion_Default___.SafeModuloInt(_115_v3, _115_v3), (_index45).Int())
		var _275_v120 _dafny.Sequence
		_ = _275_v120
		_275_v120 = _dafny.SeqOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
		(_111_globalState).F2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_275_v120, (Companion_Default___.SafeIndex(_115_v3, _dafny.IntOfUint32((_275_v120).Cardinality()))).Uint32(), (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool)), _275_v120)
		var _276_v121 D7
		_ = _276_v121
		_276_v121 = Companion_D7_.Create_DC22_(_dafny.IntOfUint32((_275_v120).Cardinality()))
		var _277_v122 D0
		_ = _277_v122
		_277_v122 = Companion_D0_.Create_DC0_((_276_v121).Dtor_cf33())
		_115_v3 = (_277_v122).Dtor_cf0()
		_158_v38 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_158_v38, _dafny.SeqOf((_dafny.SetOf((_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _117_v5)).Cardinality())), _158_v38)
	} else {
		var _278___mcc_h12 D3 = _source5.Get_().(D3_DC11).Cf13
		_ = _278___mcc_h12
		var _279_cf13 D3 = _278___mcc_h12
		_ = _279_cf13
		var _280_v123 _dafny.Map
		_ = _280_v123
		_280_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v5, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool))
		var _281_v124 _dafny.Map
		_ = _281_v124
		_281_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_280_v123).Update(_117_v5, _117_v5), _115_v3)
		var _282_v125 _dafny.Map
		_ = _282_v125
		_282_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v124, _159_v39)
		_282_v125 = (_282_v125).Update(_281_v124, _159_v39)
		var _283_v126 D8
		_ = _283_v126
		_283_v126 = Companion_D8_.Create_DC23_(_114_v2)
		_114_v2 = (_283_v126).Dtor_cf34()
		var _284_v127 *C1
		_ = _284_v127
		var _nw45 *C1 = New_C1_()
		_ = _nw45
		_nw45.Ctor__((_115_v3).Minus(_115_v3), (_115_v3).Times((_dafny.Zero).Minus(_115_v3)))
		_284_v127 = _nw45
		_117_v5 = _117_v5
	}
	var _hi2 _dafny.Int = _115_v3
	_ = _hi2
	for _285_i9 := (_158_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.IntOfUint32((_158_v38).Cardinality()))).Uint32()).(_dafny.Int); _285_i9.Cmp(_hi2) < 0; _285_i9 = _285_i9.Plus(_dafny.One) {
		var _286_v128 _dafny.Set
		_ = _286_v128
		_286_v128 = _dafny.SetOf(_115_v3)
		var _287_v129 bool
		_ = _287_v129
		var _288_v130 _dafny.Map
		_ = _288_v130
		var _out9 bool
		_ = _out9
		var _out10 _dafny.Map
		_ = _out10
		_out9, _out10 = Companion_Default___.M5(_158_v38, _286_v128, _111_globalState)
		_287_v129 = _out9
		_288_v130 = _out10
		var _289_v131 T0
		_ = _289_v131
		var _nw46 *C2 = New_C2_()
		_ = _nw46
		_nw46.Ctor__(false, _115_v3, _115_v3)
		_289_v131 = _nw46
		var _nw47 *C2 = New_C2_()
		_ = _nw47
		_nw47.Ctor__(_287_v129, _285_i9, _115_v3)
		_289_v131 = _nw47
		var _290_v132 *C3
		_ = _290_v132
		var _nw48 *C3 = New_C3_()
		_ = _nw48
		_nw48.Ctor__(Companion_Default___.Fm12(_117_v5, _117_v5, (_116_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))).Int()).(bool), _111_globalState))
		_290_v132 = _nw48
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
		_ = _index46
		var _rhs42 bool = _117_v5
		_ = _rhs42
		var _rhs43 bool = (((_dafny.SetOf(_290_v132, _290_v132)).Intersection(_dafny.SetOf(_290_v132))).Cardinality()).Cmp(_289_v131.F5()) < 0
		_ = _rhs43
		var _rhs44 bool = _117_v5
		_ = _rhs44
		var _rhs45 *C3 = _290_v132
		_ = _rhs45
		var _lhs34 _dafny.Array = _116_v4
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_116_v4), 0))
		_ = _lhs35
		var _lhs36 *GlobalState = _111_globalState
		_ = _lhs36
		_117_v5 = _rhs42
		(_lhs34).ArraySet1(_rhs43, (_lhs35).Int())
		_lhs36.F2 = _rhs44
		_290_v132 = _rhs45
		_115_v3 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-302), (_dafny.Zero).Minus((_dafny.Zero).Minus(_289_v131.F5()))), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_285_i9), (_289_v131).F4()))
	}
	_dafny.Print((_111_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_111_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_112_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v1).Equals(_dafny.MultiSetOf(_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_114_v2).F7()).Equals(_dafny.MultiSetOf(_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_115_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v4).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_117_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_118_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D4)).Dtor_cf14()).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v7).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v8).Dtor_cf14()).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_121_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_158_v38, _dafny.SeqOf(_dafny.Zero, _dafny.IntOfInt64(2), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v39).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v39).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v39).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v39).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v39).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_160_v40).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_v41.VerbatimString(false))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_() D0 {
	return D0{D0_DC2{}}
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
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0
		}
	case D0_DC2:
		{
			_, ok := other.Get_().(D0_DC2)
			return ok
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
	Cf4 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 bool) D1 {
	return D1{D1_DC4{Cf4}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

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

type D1_DC5 struct {
	Cf5 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf5 D1) D1 {
	return D1{D1_DC5{Cf5}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false)
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf5() D1 {
	return _this.Get_().(D1_DC5).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf4 == data2.Cf4
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3 == data2.Cf3
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf5.Equals(data2.Cf5)
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

type D2_DC7 struct {
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_() D2 {
	return D2{D2_DC7{}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
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

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_()
}

func (_this D2) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf6) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC7:
		{
			_, ok := other.Get_().(D2_DC7)
			return ok
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf6.Equals(data2.Cf6)
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

type D3_DC9 struct {
	Cf8  _dafny.Int
	Cf9  bool
	Cf10 _dafny.Int
	Cf11 bool
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf8 _dafny.Int, Cf9 bool, Cf10 _dafny.Int, Cf11 bool) D3 {
	return D3{D3_DC9{Cf8, Cf9, Cf10, Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf12 _dafny.CodePoint
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf12 _dafny.CodePoint) D3 {
	return D3{D3_DC10{Cf12}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf7 _dafny.Sequence
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf7 _dafny.Sequence) D3 {
	return D3{D3_DC8{Cf7}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC11 struct {
	Cf13 D3
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf13 D3) D3 {
	return D3{D3_DC11{Cf13}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.Zero, false, _dafny.Zero, false)
}

func (_this D3) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf8
}

func (_this D3) Dtor_cf9() bool {
	return _this.Get_().(D3_DC9).Cf9
}

func (_this D3) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) Dtor_cf12() _dafny.CodePoint {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf7
}

func (_this D3) Dtor_cf13() D3 {
	return _this.Get_().(D3_DC11).Cf13
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + data.Cf7.VerbatimString(true) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9 && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11 == data2.Cf11
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf12 == data2.Cf12
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf7.Equals(data2.Cf7)
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf13.Equals(data2.Cf13)
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

type D4_DC13 struct {
	Cf15 bool
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf15 bool) D4 {
	return D4{D4_DC13{Cf15}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf16 _dafny.Map
	Cf17 bool
	Cf18 _dafny.Sequence
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf16 _dafny.Map, Cf17 bool, Cf18 _dafny.Sequence) D4 {
	return D4{D4_DC14{Cf16, Cf17, Cf18}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC12 struct {
	Cf14 _dafny.Set
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf14 _dafny.Set) D4 {
	return D4{D4_DC12{Cf14}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(false)
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC13).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D4_DC14).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC14).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D4_DC14).Cf18
}

func (_this D4) Dtor_cf14() _dafny.Set {
	return _this.Get_().(D4_DC12).Cf14
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf14) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf16.Equals(data2.Cf16) && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D5_DC16 struct {
	Cf20 _dafny.Int
	Cf21 _dafny.MultiSet
	Cf22 _dafny.Sequence
	Cf23 D3
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf20 _dafny.Int, Cf21 _dafny.MultiSet, Cf22 _dafny.Sequence, Cf23 D3) D5 {
	return D5{D5_DC16{Cf20, Cf21, Cf22, Cf23}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC15 struct {
	Cf19 _dafny.Array
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf19 _dafny.Array) D5 {
	return D5{D5_DC15{Cf19}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC16_(_dafny.Zero, _dafny.EmptyMultiSet, _dafny.EmptySeq, Companion_D3_.Default())
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf20
}

func (_this D5) Dtor_cf21() _dafny.MultiSet {
	return _this.Get_().(D5_DC16).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D5_DC16).Cf22
}

func (_this D5) Dtor_cf23() D3 {
	return _this.Get_().(D5_DC16).Cf23
}

func (_this D5) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D5_DC15).Cf19
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Equals(data2.Cf21) && data1.Cf22.Equals(data2.Cf22) && data1.Cf23.Equals(data2.Cf23)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf19 == data2.Cf19
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
	Cf25 _dafny.Sequence
	Cf26 bool
	Cf27 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf25 _dafny.Sequence, Cf26 bool, Cf27 _dafny.Int) D6 {
	return D6{D6_DC18{Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC17 struct {
	Cf24 *C3
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf24 *C3) D6 {
	return D6{D6_DC17{Cf24}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC19 struct {
	Cf28 D6
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf28 D6) D6 {
	return D6{D6_DC19{Cf28}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC18_(_dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D6) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D6_DC18).Cf25
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC18).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf27
}

func (_this D6) Dtor_cf24() *C3 {
	return _this.Get_().(D6_DC17).Cf24
}

func (_this D6) Dtor_cf28() D6 {
	return _this.Get_().(D6_DC19).Cf28
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC18:
		{
			return "D6.DC18" + "(" + data.Cf25.VerbatimString(true) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf25.Equals(data2.Cf25) && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf28.Equals(data2.Cf28)
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
	Cf30 bool
	Cf31 _dafny.CodePoint
	Cf32 _dafny.MultiSet
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf30 bool, Cf31 _dafny.CodePoint, Cf32 _dafny.MultiSet) D7 {
	return D7{D7_DC21{Cf30, Cf31, Cf32}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC22 struct {
	Cf33 _dafny.Int
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf33 _dafny.Int) D7 {
	return D7{D7_DC22{Cf33}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC20 struct {
	Cf29 _dafny.Sequence
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf29 _dafny.Sequence) D7 {
	return D7{D7_DC20{Cf29}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC21_(false, _dafny.CodePoint('D'), _dafny.EmptyMultiSet)
}

func (_this D7) Dtor_cf30() bool {
	return _this.Get_().(D7_DC21).Cf30
}

func (_this D7) Dtor_cf31() _dafny.CodePoint {
	return _this.Get_().(D7_DC21).Cf31
}

func (_this D7) Dtor_cf32() _dafny.MultiSet {
	return _this.Get_().(D7_DC21).Cf32
}

func (_this D7) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf33
}

func (_this D7) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D7_DC20).Cf29
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31 && data1.Cf32.Equals(data2.Cf32)
		}
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf29.Equals(data2.Cf29)
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

type D8_DC24 struct {
	Cf35 _dafny.Int
	Cf36 _dafny.Array
	Cf37 _dafny.Array
	Cf38 _dafny.Int
	Cf39 bool
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf35 _dafny.Int, Cf36 _dafny.Array, Cf37 _dafny.Array, Cf38 _dafny.Int, Cf39 bool) D8 {
	return D8{D8_DC24{Cf35, Cf36, Cf37, Cf38, Cf39}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC23 struct {
	Cf34 *C0
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf34 *C0) D8 {
	return D8{D8_DC23{Cf34}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC25 struct {
	Cf40 D8
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf40 D8) D8 {
	return D8{D8_DC25{Cf40}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC24_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, false)
}

func (_this D8) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf35
}

func (_this D8) Dtor_cf36() _dafny.Array {
	return _this.Get_().(D8_DC24).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D8_DC24).Cf37
}

func (_this D8) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC24).Cf39
}

func (_this D8) Dtor_cf34() *C0 {
	return _this.Get_().(D8_DC23).Cf34
}

func (_this D8) Dtor_cf40() D8 {
	return _this.Get_().(D8_DC25).Cf40
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39 == data2.Cf39
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf34 == data2.Cf34
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf40.Equals(data2.Cf40)
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
	Cf41 _dafny.Map
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf41 _dafny.Map) D9 {
	return D9{D9_DC26{Cf41}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D9) Dtor_cf41() _dafny.Map {
	return _this.Get_().(D9_DC26).Cf41
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf41.Equals(data2.Cf41)
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

// Definition of trait T0
type T0 interface {
	String() string
	F5() _dafny.Int
	F5_set_(value _dafny.Int)
	M1(p0 _dafny.Array, p1 _dafny.CodePoint, globalState *GlobalState) (bool, bool, bool)
	F4() _dafny.Int
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
	F2  bool
	_f0 bool
	_f1 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = false
	_this._f0 = false
	_this._f1 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f7 _dafny.MultiSet
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f7 = _dafny.EmptyMultiSet
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

func (_this *C0) Ctor__(f7 _dafny.MultiSet) {
	{
		(_this)._f7 = f7
	}
}
func (_this *C0) F7() _dafny.MultiSet {
	{
		return _this._f7
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f5 _dafny.Int
	_f4 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f5 = _dafny.Zero
	_this._f4 = _dafny.Zero
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

func (_this *C1) F5() _dafny.Int {
	return _this._f5
}
func (_this *C1) F5_set_(value _dafny.Int) {
	_this._f5 = value
}
func (_this *C1) F4() _dafny.Int {
	return _this._f4
}
func (_this *C1) Ctor__(f4 _dafny.Int, f5 _dafny.Int) {
	{
		(_this)._f4 = f4
		(_this)._f5 = f5
	}
}
func (_this *C1) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Set, p3 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C1) Fm2(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())).Union(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F4())).Cardinality()))).IsProperSubsetOf((_dafny.MultiSetOf((_this).F4())).Difference(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())))
	}
}
func (_this *C1) M1(p0 _dafny.Array, p1 _dafny.CodePoint, globalState *GlobalState) (bool, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _291_v0 _dafny.Sequence
		_ = _291_v0
		_291_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ffyidwc")
		var _292_v1 bool
		_ = _292_v1
		_292_v1 = true
		var _293_v2 _dafny.Sequence
		_ = _293_v2
		_293_v2 = _dafny.SeqOf(_292_v1)
		(_this).F5_set_(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_291_v0, _dafny.UnicodeSeqOfUtf8Bytes("ykiwthwa"))).Cardinality()), (_dafny.Zero).Minus((_dafny.SetOf((_293_v2).Select((Companion_Default___.SafeIndex(_this.F5(), _dafny.IntOfUint32((_293_v2).Cardinality()))).Uint32()).(bool))).Cardinality())))
		var _294_v3 _dafny.Set
		_ = _294_v3
		_294_v3 = _dafny.SetOf(!(_292_v1))
		r2 = ((_294_v3).Difference(_294_v3)).IsDisjointFrom(_294_v3)
		(_this).F5_set_(_this.F5())
		r1 = _292_v1
		var _295_v4 _dafny.Map
		_ = _295_v4
		_295_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), !(true))
		if !((((_295_v4).Merge(_295_v4)).Cardinality()).Cmp(((_294_v3).Cardinality()).Minus(_this.F5())) >= 0) {
			var _rhs46 bool = (_dafny.IntOfUint32((_291_v0).Cardinality())).Cmp((_this).F4()) != 0
			_ = _rhs46
			var _rhs47 bool = (_this.F5()).Cmp(_this.F5()) >= 0
			_ = _rhs47
			r0 = _rhs46
			r1 = _rhs47
			var _296_v5 _dafny.Array
			_ = _296_v5
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw49
			_296_v5 = _nw49
			var _297_v6 _dafny.Sequence
			_ = _297_v6
			_297_v6 = _dafny.SeqOf(_296_v5, _296_v5, _296_v5, _296_v5)
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((p0), 0))
			_ = _index47
			(p0).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_297_v6).Cardinality()), _this.F5()), (_index47).Int())
			var _298_v7 _dafny.Sequence
			_ = _298_v7
			_298_v7 = _dafny.SeqOf((_this).F4())
			var _299_v8 _dafny.Map
			_ = _299_v8
			_299_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_298_v7, (Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_298_v7).Cardinality()))).Uint32(), _this.F5()), _dafny.SeqOf(_dafny.IntOfUint32((_291_v0).Cardinality()))), _this.F5())
			var _300_v9 D0
			_ = _300_v9
			_300_v9 = Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("khdwjltpl")).Cardinality()))).Cardinality())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((p0), 0))
			_ = _index48
			var _rhs48 _dafny.Int = (_this).F4()
			_ = _rhs48
			var _rhs49 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_300_v9).Dtor_cf0()), (_this).F4())
			_ = _rhs49
			var _rhs50 _dafny.Map = _299_v8
			_ = _rhs50
			var _lhs37 _dafny.Array = p0
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((p0), 0))
			_ = _lhs38
			var _lhs39 *C1 = _this
			_ = _lhs39
			(_lhs37).ArraySet1(_rhs48, (_lhs38).Int())
			_lhs39.F5_set_(_rhs49)
			_299_v8 = _rhs50
			(_this).F5_set_((_this).F4())
			var _301_v10 D1
			_ = _301_v10
			_301_v10 = Companion_D1_.Create_DC3_(_292_v1)
			var _302_v11 _dafny.Map
			_ = _302_v11
			_302_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_292_v1, (_this).F4())
			var _303_v12 _dafny.MultiSet
			_ = _303_v12
			_303_v12 = _dafny.MultiSetOf((_302_v11).Cardinality())
			var _304_v13 _dafny.Int
			_ = _304_v13
			var _305_v14 _dafny.Sequence
			_ = _305_v14
			var _306_v15 _dafny.MultiSet
			_ = _306_v15
			var _307_v16 _dafny.Sequence
			_ = _307_v16
			var _out11 _dafny.Int
			_ = _out11
			var _out12 _dafny.Sequence
			_ = _out12
			var _out13 _dafny.MultiSet
			_ = _out13
			var _out14 _dafny.Sequence
			_ = _out14
			_out11, _out12, _out13, _out14 = (_this).M4((_301_v10).Dtor_cf3(), (func() _dafny.Int {
				if (_303_v12).Contains(_dafny.IntOfInt64(369)) {
					return (_303_v12).Multiplicity(_dafny.IntOfInt64(369))
				}
				return Companion_Default___.Fm3(_292_v1, globalState)
			})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_308_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_309_i0 _dafny.Int) _dafny.CodePoint {
					return _308_p1
				}
			})(p1))), globalState)
			_304_v13 = _out11
			_305_v14 = _out12
			_306_v15 = _out13
			_307_v16 = _out14
			var _310_v17 _dafny.Map
			_ = _310_v17
			_310_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_304_v13, _304_v13), (_this.F5()).Plus(_dafny.IntOfInt64(933)))
			var _311_v18 _dafny.MultiSet
			_ = _311_v18
			_311_v18 = _dafny.MultiSetOf(false)
			_310_v17 = (_310_v17).Update(Companion_Default___.SafeModuloInt((_this).F4(), _this.F5()), ((_dafny.MultiSetOf(true, _292_v1, _292_v1, true, _292_v1)).Intersection((_311_v18).Update(_292_v1, Companion_Default___.Abs(_this.F5())))).Cardinality())
		} else {
			var _312_v19 _dafny.MultiSet
			_ = _312_v19
			_312_v19 = _dafny.MultiSetOf(_dafny.CodePoint('i'), p1)
			var _313_v20 *C0
			_ = _313_v20
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__(_312_v19)
			_313_v20 = _nw50
			var _314_v21 _dafny.Array
			_ = _314_v21
			var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
			_ = _nw51
			_314_v21 = _nw51
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_314_v21), 0))
			_ = _index49
			(_314_v21).ArraySet1CodePoint(p1, (_index49).Int())
			var _315_v22 _dafny.Sequence
			_ = _315_v22
			_315_v22 = _dafny.SeqOf((_this).F4())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_314_v21), 0))
			_ = _index50
			(_314_v21).ArraySet1CodePoint(Companion_Default___.Fm4((_this).Fm2(!(_292_v1), (_this).F4(), _292_v1, globalState), _this.F5(), _292_v1, Companion_Default___.SafeDivisionInt(_this.F5(), (_315_v22).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_315_v22).Cardinality()))).Uint32()).(_dafny.Int)), globalState), (_index50).Int())
			r2 = !(true)
			var _316_v23 _dafny.Sequence
			_ = _316_v23
			_316_v23 = _dafny.SeqOf(_315_v22)
			r0 = !_dafny.Companion_Sequence_.Equal(_316_v23, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(338))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_317_v22 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_318_i1 _dafny.Int) _dafny.Sequence {
					return _317_v22
				}
			})(_315_v22))))
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((p0), 0))
			_ = _index51
			(p0).ArraySet1(_this.F5(), (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((p0), 0))
			_ = _index52
			(p0).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm3(_292_v1, globalState))), (_index52).Int())
		}
		var _319_v24 D1
		_ = _319_v24
		_319_v24 = Companion_D1_.Create_DC3_(_292_v1)
		var _320_v25 _dafny.Map
		_ = _320_v25
		_320_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _292_v1)
		var _321_v26 _dafny.Sequence
		_ = _321_v26
		_321_v26 = _dafny.SeqOf(_dafny.IntOfInt64(160))
		var _322_v27 _dafny.MultiSet
		_ = _322_v27
		_322_v27 = _dafny.MultiSetOf(_this.F5(), _this.F5(), _dafny.IntOfUint32((_321_v26).Cardinality()), Companion_Default___.Fm3(_292_v1, globalState))
		var _323_v28 _dafny.Array
		_ = _323_v28
		var _nwElement0_5 bool = (func() bool {
			if _292_v1 {
				return _292_v1
			}
			return (_this).Fm2(_292_v1, _this.F5(), _292_v1, globalState)
		})()
		_ = _nwElement0_5
		var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
		_ = _nw52
		(_nw52).ArraySet1(_nwElement0_5, 0)
		(_nw52).ArraySet1(!(_292_v1) || (_292_v1), 1)
		(_nw52).ArraySet1(_292_v1, 2)
		(_nw52).ArraySet1(_292_v1, 3)
		(_nw52).ArraySet1(_292_v1, 4)
		(_nw52).ArraySet1(!(!((_319_v24).Dtor_cf3())), 5)
		(_nw52).ArraySet1((_this).Fm1(_292_v1, _this.F5(), _dafny.SetOf(false, _292_v1), (func() bool {
			if (_320_v25).Contains(_292_v1) {
				return (_320_v25).Get(_292_v1).(bool)
			}
			return _292_v1
		})(), globalState), 6)
		(_nw52).ArraySet1(_292_v1, 7)
		(_nw52).ArraySet1(_292_v1, 8)
		(_nw52).ArraySet1(_292_v1, 9)
		(_nw52).ArraySet1(_292_v1, 10)
		(_nw52).ArraySet1(_292_v1, 11)
		(_nw52).ArraySet1((_322_v27).IsSubsetOf(_322_v27), 12)
		(_nw52).ArraySet1(_292_v1, 13)
		(_nw52).ArraySet1(!((Companion_Default___.Fm3(_292_v1, globalState)).Cmp(_dafny.IntOfInt64(252)) != 0), 14)
		(_nw52).ArraySet1(_292_v1, 15)
		(_nw52).ArraySet1(_292_v1, 16)
		(_nw52).ArraySet1((true) || (_292_v1), 17)
		(_nw52).ArraySet1(!(_292_v1) || (_292_v1), 18)
		(_nw52).ArraySet1(_292_v1, 19)
		(_nw52).ArraySet1(_292_v1, 20)
		(_nw52).ArraySet1(_292_v1, 21)
		(_nw52).ArraySet1(_292_v1, 22)
		(_nw52).ArraySet1((_319_v24).Dtor_cf3(), 23)
		(_nw52).ArraySet1((_293_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.IntOfUint32((_293_v2).Cardinality()))).Uint32()).(bool), 24)
		_323_v28 = _nw52
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_323_v28), 0))
		_ = _index53
		(_323_v28).ArraySet1(!(_292_v1), (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_323_v28), 0))
		_ = _index54
		(_323_v28).ArraySet1(_292_v1, (_index54).Int())
		r0 = (_323_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_323_v28), 0))).Int()).(bool)
		var _324_v29 _dafny.Map
		_ = _324_v29
		_324_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F5(), _this.F5())
		r1 = ((_324_v29).Cardinality()).Cmp(_dafny.IntOfInt64(-967)) <= 0
		r2 = (func() bool {
			if (_323_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_323_v28), 0))).Int()).(bool) {
				return _292_v1
			}
			return _292_v1
		})()
		return r0, r1, r2
	}
}
func (_this *C1) M3(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _325_v0 bool
		_ = _325_v0
		_325_v0 = true
		(globalState).F2 = _325_v0
		if _325_v0 {
			var _326_v1 _dafny.MultiSet
			_ = _326_v1
			_326_v1 = _dafny.MultiSetOf(_dafny.CodePoint('x'))
			var _327_v2 *C0
			_ = _327_v2
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__(_326_v1)
			_327_v2 = _nw53
			var _328_v3 *C0
			_ = _328_v3
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__((_327_v2).F7())
			_328_v3 = _nw54
			var _329_v4 _dafny.Array
			_ = _329_v4
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw55
			_329_v4 = _nw55
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))
			_ = _index55
			(_329_v4).ArraySet1(!(_325_v0), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))
			_ = _index56
			(_329_v4).ArraySet1(_325_v0, (_index56).Int())
			_328_v3 = _328_v3
			var _330_v5 D0
			_ = _330_v5
			_330_v5 = Companion_D0_.Create_DC0_(_this.F5())
			var _source6 D0 = _330_v5
			_ = _source6
			if _source6.Is_DC1() {
				var _331___mcc_h0 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
				_ = _331___mcc_h0
				var _332___mcc_h1 _dafny.Int = _source6.Get_().(D0_DC1).Cf2
				_ = _332___mcc_h1
				var _333_cf2 _dafny.Int = _332___mcc_h1
				_ = _333_cf2
				var _334_cf1 _dafny.Int = _331___mcc_h0
				_ = _334_cf1
				var _335_v6 _dafny.Sequence
				_ = _335_v6
				_335_v6 = _dafny.SeqOf(_333_cf2)
				var _336_v7 _dafny.Sequence
				_ = _336_v7
				_336_v7 = _dafny.SeqOf((_dafny.IntOfUint32((_335_v6).Cardinality())).Cmp(_this.F5()) >= 0)
				var _337_v8 _dafny.Sequence
				_ = _337_v8
				_337_v8 = _dafny.UnicodeSeqOfUtf8Bytes("omrnrrxhq")
				_336_v7 = _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("hyq"), _337_v8))
				var _338_v9 _dafny.Set
				_ = _338_v9
				_338_v9 = _dafny.SetOf(_325_v0, (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool))
				var _339_v10 _dafny.Map
				_ = _339_v10
				_339_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_338_v9).IsProperSubsetOf(_338_v9))
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))
				_ = _index57
				(_329_v4).ArraySet1((func() bool {
					if (_339_v10).Contains(!((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool))) {
						return (_339_v10).Get(!((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool))).(bool)
					}
					return (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool)
				})(), (_index57).Int())
				var _340_v11 _dafny.Map
				_ = _340_v11
				_340_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool), _dafny.IntOfInt64(-72), (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool), globalState), _dafny.SeqOf(_dafny.IntOfInt64(134), _dafny.IntOfInt64(287)))
				_335_v6 = (func() _dafny.Sequence {
					if (_340_v11).Contains(_325_v0) {
						return (_340_v11).Get(_325_v0).(_dafny.Sequence)
					}
					return _335_v6
				})()
				(globalState).F2 = _325_v0
			} else if _source6.Is_DC2() {
				var _341_v12 _dafny.Array
				_ = _341_v12
				var _nw56 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(21))
				_ = _nw56
				_341_v12 = _nw56
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_341_v12), 0))
				_ = _index58
				(_341_v12).ArraySet1(Companion_D0_.Create_DC2_(), (_index58).Int())
				var _342_v13 _dafny.Map
				_ = _342_v13
				_342_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool)) || (_325_v0), _325_v0)
				var _343_v14 _dafny.Set
				_ = _343_v14
				_343_v14 = _dafny.SetOf(_325_v0)
				var _344_v16 _dafny.Sequence
				_ = _344_v16
				_344_v16 = _dafny.UnicodeSeqOfUtf8Bytes("cxo")
				var _345_v17 _dafny.Map
				_ = _345_v17
				_345_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _344_v16)
				var _346_v18 _dafny.Array
				_ = _346_v18
				var _nwElement0_6 _dafny.Sequence = p0
				_ = _nwElement0_6
				var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(22))
				_ = _nw57
				(_nw57).ArraySet1(_nwElement0_6, 0)
				(_nw57).ArraySet1(p0, 1)
				(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_325_v0, (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool)), _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _325_v0)), 2)
				(_nw57).ArraySet1(p0, 3)
				(_nw57).ArraySet1(_dafny.SeqOf(_325_v0), 4)
				(_nw57).ArraySet1(_dafny.SeqOf((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool), false), 5)
				(_nw57).ArraySet1(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_this.F5(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_this).Fm1(_325_v0, (_dafny.Zero).Minus((_this).F4()), _343_v14, _325_v0, globalState)), 6)
				(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 7)
				(_nw57).ArraySet1(p0, 8)
				(_nw57).ArraySet1(_dafny.SeqOf((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool), (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool)), 9)
				(_nw57).ArraySet1(_dafny.SeqOf((p0).Select((Companion_Default___.SafeIndex(_this.F5(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(bool)), 10)
				(_nw57).ArraySet1(_dafny.SeqOf(_325_v0), 11)
				(_nw57).ArraySet1(p0, 12)
				(_nw57).ArraySet1(Companion_Default___.Fm5(func() _dafny.Map {
					var _coll9 = _dafny.NewMapBuilder()
					_ = _coll9
					for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-360), _dafny.IntOfInt64(250))); ; {
						_compr_9, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _347_v15 _dafny.Int
						_347_v15 = interface{}(_compr_9).(_dafny.Int)
						if ((_dafny.IntOfInt64(-360)).Cmp(_347_v15) <= 0) && ((_347_v15).Cmp(_dafny.IntOfInt64(250)) < 0) {
							_coll9.Add((_347_v15).Plus(_this.F5()), _dafny.UnicodeSeqOfUtf8Bytes("dfrfshin"))
						}
					}
					return _coll9.ToMap()
				}(), _this.F5(), _325_v0, (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool), globalState), 13)
				(_nw57).ArraySet1(p0, 14)
				(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_345_v17, _this.F5(), (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool), _325_v0, globalState), p0), 15)
				(_nw57).ArraySet1(p0, 16)
				(_nw57).ArraySet1(p0, 17)
				(_nw57).ArraySet1(p0, 18)
				(_nw57).ArraySet1(p0, 19)
				(_nw57).ArraySet1(p0, 20)
				(_nw57).ArraySet1(p0, 21)
				_346_v18 = _nw57
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_346_v18), 0))
				_ = _index59
				(_346_v18).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), (_index59).Int())
				var _348_v19 _dafny.Map
				_ = _348_v19
				_348_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F4()), _this.F5())
				var _349_v20 D0
				_ = _349_v20
				_349_v20 = Companion_D0_.Create_DC2_()
				var _350_v21 _dafny.Map
				_ = _350_v21
				_350_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool))
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_341_v12), 0))
				_ = _index60
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_346_v18), 0))
				_ = _index61
				var _rhs51 bool = (_this.F5()).Cmp((_348_v19).Cardinality()) <= 0
				_ = _rhs51
				var _rhs52 D0 = _349_v20
				_ = _rhs52
				var _rhs53 _dafny.Map = _342_v13
				_ = _rhs53
				var _rhs54 _dafny.Sequence = p0
				_ = _rhs54
				var _rhs55 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt((_330_v5).Dtor_cf0(), _this.F5())).Plus(Companion_Default___.SafeModuloInt((_350_v21).Cardinality(), (_this).F4())))
				_ = _rhs55
				var _lhs40 *GlobalState = globalState
				_ = _lhs40
				var _lhs41 _dafny.Array = _341_v12
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_341_v12), 0))
				_ = _lhs42
				var _lhs43 _dafny.Array = _346_v18
				_ = _lhs43
				var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_346_v18), 0))
				_ = _lhs44
				_lhs40.F2 = _rhs51
				(_lhs41).ArraySet1(_rhs52, (_lhs42).Int())
				_342_v13 = _rhs53
				(_lhs43).ArraySet1(_rhs54, (_lhs44).Int())
				r1 = _rhs55
				_348_v19 = (_348_v19).Update(Companion_Default___.SafeModuloInt(_this.F5(), (func() _dafny.Set {
					var _coll10 = _dafny.NewBuilder()
					_ = _coll10
					for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(814), _dafny.IntOfInt64(257))); ; {
						_compr_10, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _351_v22 _dafny.Int
						_351_v22 = interface{}(_compr_10).(_dafny.Int)
						if ((_dafny.IntOfInt64(814)).Cmp(_351_v22) <= 0) && ((_351_v22).Cmp(_dafny.IntOfInt64(257)) < 0) {
							_coll10.Add((_351_v22).Plus(_this.F5()))
						}
					}
					return _coll10.ToSet()
				}()).Cardinality()), Companion_Default___.SafeDivisionInt(_this.F5(), (_this).F4()))
				var _352_v23 D1
				_ = _352_v23
				_352_v23 = Companion_D1_.Create_DC3_(!(_325_v0) || ((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool)))
				_352_v23 = _352_v23
				(_this).F5_set_((Companion_Default___.Fm3(_325_v0, globalState)).Minus((_this).F4()))
			} else {
				var _353___mcc_h2 _dafny.Int = _source6.Get_().(D0_DC0).Cf0
				_ = _353___mcc_h2
				var _354_cf0 _dafny.Int = _353___mcc_h2
				_ = _354_cf0
				_330_v5 = _330_v5
				var _355_v24 _dafny.Set
				_ = _355_v24
				_355_v24 = _dafny.SetOf((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool))
				(_this).F5_set_(Companion_Default___.Fm3(!(_355_v24).Contains((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool)), globalState))
				var _356_v25 _dafny.Map
				_ = _356_v25
				_356_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool), (_this).F4())
				_356_v25 = (_356_v25).Update(!((_329_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_329_v4), 0))).Int()).(bool)), (_this).F4())
				var _357_v26 _dafny.MultiSet
				_ = _357_v26
				_357_v26 = _dafny.MultiSetOf(_this.F5())
				var _358_v27 _dafny.Sequence
				_ = _358_v27
				_358_v27 = _dafny.UnicodeSeqOfUtf8Bytes("x")
				r0 = ((_dafny.Zero).Minus((func() _dafny.Int {
					if _325_v0 {
						return _this.F5()
					}
					return (_357_v26).Cardinality()
				})())).Plus(_dafny.IntOfUint32((_358_v27).Cardinality()))
			}
		} else {
			var _359_v28 _dafny.CodePoint
			_ = _359_v28
			_359_v28 = _dafny.CodePoint('e')
			var _360_v29 _dafny.Array
			_ = _360_v29
			var _nwElement0_7 _dafny.CodePoint = _359_v28
			_ = _nwElement0_7
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(3))
			_ = _nw58
			(_nw58).ArraySet1CodePoint(_nwElement0_7, 0)
			(_nw58).ArraySet1CodePoint(_359_v28, 1)
			(_nw58).ArraySet1CodePoint(_359_v28, 2)
			_360_v29 = _nw58
			var _rhs56 bool = _325_v0
			_ = _rhs56
			var _rhs57 bool = ((_this).F4()).Cmp((_this).F4()) == 0
			_ = _rhs57
			var _rhs58 _dafny.Int = Companion_Default___.Fm3(_dafny.Companion_Sequence_.IsProperPrefixOf(p0, p0), globalState)
			_ = _rhs58
			var _rhs59 _dafny.Array = _360_v29
			_ = _rhs59
			var _rhs60 _dafny.Int = (_this).F4()
			_ = _rhs60
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			_325_v0 = _rhs56
			_lhs45.F2 = _rhs57
			r0 = _rhs58
			_360_v29 = _rhs59
			r1 = _rhs60
			var _361_v30 _dafny.MultiSet
			_ = _361_v30
			_361_v30 = _dafny.MultiSetOf(_359_v28)
			var _362_v31 *C0
			_ = _362_v31
			var _nw59 *C0 = New_C0_()
			_ = _nw59
			_nw59.Ctor__(_361_v30)
			_362_v31 = _nw59
			var _363_v32 _dafny.Map
			_ = _363_v32
			_363_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F5())
			_363_v32 = (_363_v32).Update(_325_v0, _this.F5())
			var _364_v33 _dafny.Array
			_ = _364_v33
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw60
			_364_v33 = _nw60
			_364_v33 = _364_v33
			(globalState).F2 = _325_v0
		}
		var _365_v34 _dafny.CodePoint
		_ = _365_v34
		_365_v34 = _dafny.CodePoint('b')
		var _366_v35 _dafny.Sequence
		_ = _366_v35
		_366_v35 = _dafny.SeqOf(_365_v34)
		var _367_v36 *C0
		_ = _367_v36
		var _nw61 *C0 = New_C0_()
		_ = _nw61
		_nw61.Ctor__(_dafny.MultiSetFromSeq(_366_v35))
		_367_v36 = _nw61
		var _368_v37 _dafny.Array
		_ = _368_v37
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_4
		var _nw62 _dafny.Array
		_ = _nw62
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw62 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.CodePoint = (func(_369_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_370_i0 _dafny.Int) _dafny.CodePoint {
					return _369_v34
				}
			})(_365_v34)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw62 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw62).ArraySet1CodePoint(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw62).ArraySet1CodePoint(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_368_v37 = _nw62
		var _371_v38 _dafny.Map
		_ = _371_v38
		_371_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_365_v34, _368_v37)
		var _372_v39 _dafny.Sequence
		_ = _372_v39
		_372_v39 = _dafny.SeqOf((func() _dafny.Array {
			if (_371_v38).Contains(_365_v34) {
				return (_371_v38).Get(_365_v34).(_dafny.Array)
			}
			return _368_v37
		})(), _368_v37, _368_v37)
		if _dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
			if _325_v0 {
				return _372_v39
			}
			return _372_v39
		})(), _dafny.SeqOf(_368_v37, _368_v37)) {
			var _373_v40 _dafny.Array
			_ = _373_v40
			var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw63
			_373_v40 = _nw63
			var _374_v41 _dafny.Map
			_ = _374_v41
			_374_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v0, _373_v40)
			_374_v41 = (_374_v41).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v0, _373_v40))
			var _375_v42 _dafny.Sequence
			_ = _375_v42
			_375_v42 = _dafny.SeqOf(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dgn")).Cardinality()))
			(globalState).F2 = (Companion_Default___.SafeDivisionInt(_this.F5(), _dafny.IntOfUint32((_375_v42).Cardinality()))).Cmp((_this.F5()).Minus(_this.F5())) < 0
			var _376_v43 *C0
			_ = _376_v43
			var _nw64 *C0 = New_C0_()
			_ = _nw64
			_nw64.Ctor__(_dafny.MultiSetFromSeq(_366_v35))
			_376_v43 = _nw64
			var _377_v44 D0
			_ = _377_v44
			_377_v44 = Companion_D0_.Create_DC1_((_this.F5()).Plus(_this.F5()), _this.F5())
			var _source7 D0 = _377_v44
			_ = _source7
			if _source7.Is_DC1() {
				var _378___mcc_h3 _dafny.Int = _source7.Get_().(D0_DC1).Cf1
				_ = _378___mcc_h3
				var _379___mcc_h4 _dafny.Int = _source7.Get_().(D0_DC1).Cf2
				_ = _379___mcc_h4
				var _380_cf2 _dafny.Int = _379___mcc_h4
				_ = _380_cf2
				var _381_cf1 _dafny.Int = _378___mcc_h3
				_ = _381_cf1
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_373_v40), 0))
				_ = _index62
				(_373_v40).ArraySet1(_this.F5(), (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_373_v40), 0))
				_ = _index63
				(_373_v40).ArraySet1(_380_cf2, (_index63).Int())
				var _382_v45 _dafny.Int
				_ = _382_v45
				var _383_v46 _dafny.Sequence
				_ = _383_v46
				var _384_v47 _dafny.MultiSet
				_ = _384_v47
				var _385_v48 _dafny.Sequence
				_ = _385_v48
				var _out15 _dafny.Int
				_ = _out15
				var _out16 _dafny.Sequence
				_ = _out16
				var _out17 _dafny.MultiSet
				_ = _out17
				var _out18 _dafny.Sequence
				_ = _out18
				_out15, _out16, _out17, _out18 = (_this).M4(_325_v0, (_this).F4(), Companion_Default___.Fm6(_325_v0, _this.F5(), _325_v0, globalState), globalState)
				_382_v45 = _out15
				_383_v46 = _out16
				_384_v47 = _out17
				_385_v48 = _out18
				_374_v41 = _374_v41
				var _386_v49 _dafny.Array
				_ = _386_v49
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_5
				var _nw65 _dafny.Array
				_ = _nw65
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw65 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Int = (func(_387_cf1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_388_i1 _dafny.Int) _dafny.Int {
							return (_388_i1).Times(_387_cf1)
						}
					})(_381_cf1)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw65 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw65).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw65).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_386_v49 = _nw65
				_386_v49 = _386_v49
			} else if _source7.Is_DC2() {
				var _389_v50 _dafny.Map
				_ = _389_v50
				_389_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_325_v0) == (_325_v0), _325_v0)
				_389_v50 = (_389_v50).Update(!(_325_v0), _325_v0)
				_377_v44 = _377_v44
				_325_v0 = _325_v0
				_389_v50 = (_389_v50).Update(_325_v0, false)
			} else {
				var _390___mcc_h5 _dafny.Int = _source7.Get_().(D0_DC0).Cf0
				_ = _390___mcc_h5
				var _391_cf0 _dafny.Int = _390___mcc_h5
				_ = _391_cf0
				(globalState).F2 = _325_v0
				var _392_v51 _dafny.Map
				_ = _392_v51
				_392_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F4())
				var _393_v52 D1
				_ = _393_v52
				_393_v52 = Companion_D1_.Create_DC4_(_325_v0)
				var _pat_let_tv10 = _325_v0
				_ = _pat_let_tv10
				var _394_v53 _dafny.Set
				_ = _394_v53
				_394_v53 = _dafny.SetOf(Companion_Default___.Fm7(_325_v0, _392_v51, _365_v34, _325_v0, globalState), _393_v52, Companion_Default___.Fm7(_325_v0, _392_v51, _365_v34, _325_v0, globalState), _393_v52, func(_pat_let6_0 D1) D1 {
					return func(_395_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let7_0 bool) D1 {
							return func(_396_dt__update_hcf4_h0 bool) D1 {
								return Companion_D1_.Create_DC4_(_396_dt__update_hcf4_h0)
							}(_pat_let7_0)
						}(_pat_let_tv10)
					}(_pat_let6_0)
				}(_393_v52))
				_394_v53 = func() _dafny.Set {
					var _coll11 = _dafny.NewBuilder()
					_ = _coll11
					for _iter11 := _dafny.Iterate((Companion_Default___.Fm8(_365_v34, _325_v0, _391_cf0, _325_v0, globalState)).Keys().Elements()); ; {
						_compr_11, _ok11 := _iter11()
						if !_ok11 {
							break
						}
						var _397_v54 D1
						_397_v54 = interface{}(_compr_11).(D1)
						if (Companion_Default___.Fm8(_365_v34, _325_v0, _391_cf0, _325_v0, globalState)).Contains(_397_v54) {
							_coll11.Add(_397_v54)
						}
					}
					return _coll11.ToSet()
				}()
				var _398_v55 _dafny.Array
				_ = _398_v55
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_6
				var _nw66 _dafny.Array
				_ = _nw66
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw66 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) bool = (func(_399_v0 bool) func(_dafny.Int) bool {
						return func(_400_i2 _dafny.Int) bool {
							return _399_v0
						}
					})(_325_v0)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw66 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw66).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw66).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_398_v55 = _nw66
				var _401_v56 _dafny.Map
				_ = _401_v56
				_401_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_391_cf0, _365_v34)
				var _rhs61 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_401_v56).Contains((_this).F4()) {
						return (_401_v56).Get((_this).F4()).(_dafny.CodePoint)
					}
					return _365_v34
				})()
				_ = _rhs61
				var _rhs62 _dafny.CodePoint = _365_v34
				_ = _rhs62
				var _rhs63 bool = !(!(_325_v0) || (_325_v0)) || (!(_325_v0))
				_ = _rhs63
				var _rhs64 _dafny.Array = _398_v55
				_ = _rhs64
				_365_v34 = _rhs61
				_365_v34 = _rhs62
				_325_v0 = _rhs63
				_398_v55 = _rhs64
				(globalState).F2 = _dafny.Companion_Sequence_.Contains(_366_v35, _dafny.CodePoint('u'))
			}
			r0 = (_375_v42).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_dafny.MultiSetOf(_325_v0, _325_v0)).Cardinality()).Times((_this).F4())), _dafny.IntOfUint32((_375_v42).Cardinality()))).Uint32()).(_dafny.Int)
		} else {
			var _402_v57 _dafny.Array
			_ = _402_v57
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_7
			var _nw67 _dafny.Array
			_ = _nw67
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw67 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = (func(_403_v0 bool) func(_dafny.Int) bool {
					return func(_404_i3 _dafny.Int) bool {
						return !(_403_v0) || (_403_v0)
					}
				})(_325_v0)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw67 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw67).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw67).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_402_v57 = _nw67
			var _405_v58 _dafny.Array
			_ = _405_v58
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
			_ = _nw68
			_405_v58 = _nw68
			var _406_v59 _dafny.Set
			_ = _406_v59
			_406_v59 = _dafny.SetOf(_325_v0)
			var _rhs65 _dafny.Array = _402_v57
			_ = _rhs65
			var _rhs66 bool = !((_406_v59).Difference(_406_v59)).Equals(_406_v59)
			_ = _rhs66
			var _rhs67 _dafny.Array = _405_v58
			_ = _rhs67
			var _rhs68 bool = (_this).Fm2(_325_v0, _this.F5(), _325_v0, globalState)
			_ = _rhs68
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			_402_v57 = _rhs65
			_lhs46.F2 = _rhs66
			_405_v58 = _rhs67
			_325_v0 = _rhs68
			if true {
				var _407_v60 _dafny.Sequence
				_ = _407_v60
				_407_v60 = _dafny.SeqOf((_367_v36).F7())
				var _408_v61 *C0
				_ = _408_v61
				var _nw69 *C0 = New_C0_()
				_ = _nw69
				_nw69.Ctor__((_407_v60).Select((Companion_Default___.SafeIndex(_this.F5(), _dafny.IntOfUint32((_407_v60).Cardinality()))).Uint32()).(_dafny.MultiSet))
				_408_v61 = _nw69
				var _409_v62 _dafny.Array
				_ = _409_v62
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw70
				_409_v62 = _nw70
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw71
				_409_v62 = _nw71
				var _410_v63 *C0
				_ = _410_v63
				var _nw72 *C0 = New_C0_()
				_ = _nw72
				_nw72.Ctor__((_367_v36).F7())
				_410_v63 = _nw72
				r0 = ((_this).F4()).Plus((_this).F4())
				var _411_v64 D0
				_ = _411_v64
				_411_v64 = Companion_D0_.Create_DC0_(_this.F5())
				r0 = (_411_v64).Dtor_cf0()
			} else {
				var _412_v65 _dafny.Array
				_ = _412_v65
				var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw73
				_412_v65 = _nw73
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(826), _dafny.ArrayLen((_412_v65), 0))
				_ = _index64
				(_412_v65).ArraySet1(_this.F5(), (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(826), _dafny.ArrayLen((_412_v65), 0))
				_ = _index65
				(_412_v65).ArraySet1((_dafny.Zero).Minus((_this).F4()), (_index65).Int())
				(globalState).F2 = !(!(_325_v0))
				_402_v57 = _402_v57
				_412_v65 = _412_v65
				(globalState).F2 = _dafny.Companion_Sequence_.Equal(_366_v35, _366_v35)
			}
			var _413_v66 _dafny.MultiSet
			_ = _413_v66
			_413_v66 = _dafny.MultiSetOf(_325_v0)
			var _414_v67 _dafny.Sequence
			_ = _414_v67
			_414_v67 = _dafny.SeqOf(_366_v35)
			var _415_v68 _dafny.Map
			_ = _415_v68
			_415_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm6(_325_v0, (_this).F4(), _325_v0, globalState)).Cardinality()), false)
			var _416_v69 _dafny.Set
			_ = _416_v69
			_416_v69 = _dafny.SetOf(_415_v68, _415_v68, _415_v68)
			var _417_v70 _dafny.Map
			_ = _417_v70
			_417_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_413_v66).IsDisjointFrom(_413_v66), !(_dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm6(!(_325_v0), _this.F5(), _325_v0, globalState), (_414_v67).Select((Companion_Default___.SafeIndex((_416_v69).Cardinality(), _dafny.IntOfUint32((_414_v67).Cardinality()))).Uint32()).(_dafny.Sequence))))
			_417_v70 = (_417_v70).Update(_325_v0, _325_v0)
			var _418_v71 _dafny.MultiSet
			_ = _418_v71
			_418_v71 = _dafny.MultiSetOf(p0)
			var _rhs69 bool = (_325_v0) && (_325_v0)
			_ = _rhs69
			var _rhs70 _dafny.Int = ((_dafny.Zero).Minus((func() _dafny.Int {
				if (_418_v71).Contains(p0) {
					return (_418_v71).Multiplicity(p0)
				}
				return _this.F5()
			})())).Minus(_dafny.IntOfInt64(215))
			_ = _rhs70
			var _lhs47 *GlobalState = globalState
			_ = _lhs47
			_lhs47.F2 = _rhs69
			r0 = _rhs70
			(_this).F5_set_(Companion_Default___.SafeModuloInt((_this).F4(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _325_v0 {
					return _366_v35
				}
				return _dafny.Companion_Sequence_.Update(_366_v35, (Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_366_v35).Cardinality()))).Uint32(), _365_v34)
			})()).Cardinality())))
		}
		var _419_v72 _dafny.MultiSet
		_ = _419_v72
		_419_v72 = _dafny.MultiSetOf(_325_v0)
		_419_v72 = _419_v72
		var _420_v73 _dafny.Map
		_ = _420_v73
		_420_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
			if _325_v0 {
				return _366_v35
			}
			return _366_v35
		})(), _325_v0)
		_420_v73 = (_420_v73).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(584))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_421_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})), _dafny.UnicodeSeqOfUtf8Bytes("pu")), _325_v0)
		var _422_v74 _dafny.Map
		_ = _422_v74
		_422_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v0, _365_v34)
		r0 = Companion_Default___.SafeDivisionInt(((_422_v74).Cardinality()).Minus((_this).F4()), (_this).F4())
		r1 = (_this).F4()
		return r0, r1
	}
}
func (_this *C1) M4(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, _dafny.MultiSet, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _423_v0 _dafny.Sequence
		_ = _423_v0
		_423_v0 = _dafny.SeqOf(p1)
		var _424_v1 _dafny.MultiSet
		_ = _424_v1
		_424_v1 = _dafny.MultiSetOf((_dafny.SetOf(_423_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(658))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_425_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_426_i0 _dafny.Int) _dafny.Int {
				return _425_p1
			}
		})(p1))))).Cardinality())
		var _427_v2 _dafny.Sequence
		_ = _427_v2
		_427_v2 = _dafny.SeqOf((_424_v1).IsSubsetOf(_424_v1))
		(globalState).F2 = (_427_v2).Select((Companion_Default___.SafeIndex(_this.F5(), _dafny.IntOfUint32((_427_v2).Cardinality()))).Uint32()).(bool)
		var _428_v3 D1
		_ = _428_v3
		_428_v3 = Companion_D1_.Create_DC3_(p0)
		var _429_v4 _dafny.Set
		_ = _429_v4
		_429_v4 = _dafny.SetOf(Companion_D1_.Create_DC5_(_428_v3))
		var _430_v5 D1
		_ = _430_v5
		_430_v5 = Companion_D1_.Create_DC5_(_428_v3)
		_429_v4 = _dafny.SetOf(_430_v5, _430_v5)
		var _431_v6 _dafny.MultiSet
		_ = _431_v6
		_431_v6 = _dafny.MultiSetOf(p0)
		(globalState).F2 = (_this).Fm1(p0, _this.F5(), Companion_Default___.Fm9(p1, p0, (_431_v6).Cardinality(), globalState), (p1).Cmp(_this.F5()) >= 0, globalState)
		if p0 {
			var _432_v7 _dafny.Map
			_ = _432_v7
			_432_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _423_v0)
			var _433_v8 _dafny.Map
			_ = _433_v8
			_433_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _434_v9 _dafny.Set
			_ = _434_v9
			_434_v9 = _dafny.SetOf(p0)
			var _435_v10 D1
			_ = _435_v10
			_435_v10 = Companion_D1_.Create_DC4_(p0)
			var _436_v11 _dafny.Sequence
			_ = _436_v11
			_436_v11 = _dafny.SeqOf(_dafny.SetOf((_435_v10).Dtor_cf4(), p0), _434_v9, _434_v9)
			var _437_v12 _dafny.MultiSet
			_ = _437_v12
			_437_v12 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(431))).Uint32(), func(coer15 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_438_p0 bool) func(_dafny.Int) D1 {
				return func(_439_i1 _dafny.Int) D1 {
					return Companion_D1_.Create_DC4_(_438_p0)
				}
			})(p0))))
			var _440_v13 _dafny.Sequence
			_ = _440_v13
			_440_v13 = _dafny.SeqOf(_435_v10, _435_v10)
			var _441_v14 _dafny.Array
			_ = _441_v14
			var _nwElement0_8 _dafny.Int = p1
			_ = _nwElement0_8
			var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(21))
			_ = _nw74
			(_nw74).ArraySet1(_nwElement0_8, 0)
			(_nw74).ArraySet1((_dafny.Zero).Minus((p1).Plus((_dafny.Zero).Minus((_423_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.IntOfUint32((_423_v0).Cardinality()))).Uint32()).(_dafny.Int)))), 1)
			(_nw74).ArraySet1(((_432_v7).Cardinality()).Plus(_this.F5()), 2)
			(_nw74).ArraySet1((_433_v8).Cardinality(), 3)
			(_nw74).ArraySet1((_dafny.Zero).Minus(((_434_v9).Intersection((_436_v11).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_436_v11).Cardinality()))).Uint32()).(_dafny.Set))).Cardinality()), 4)
			(_nw74).ArraySet1((_dafny.Zero).Minus(p1), 5)
			(_nw74).ArraySet1((_this).F4(), 6)
			(_nw74).ArraySet1(p1, 7)
			(_nw74).ArraySet1((func() _dafny.Int {
				if (_437_v12).Contains(_440_v13) {
					return (_437_v12).Multiplicity(_440_v13)
				}
				return (_this).F4()
			})(), 8)
			(_nw74).ArraySet1(_this.F5(), 9)
			(_nw74).ArraySet1(p1, 10)
			(_nw74).ArraySet1(_this.F5(), 11)
			(_nw74).ArraySet1(_this.F5(), 12)
			(_nw74).ArraySet1(_dafny.IntOfInt64(-836), 13)
			(_nw74).ArraySet1(Companion_Default___.Fm3((_this).Fm1(!(p0), Companion_Default___.Fm3(p0, globalState), _434_v9, p0, globalState), globalState), 14)
			(_nw74).ArraySet1(_dafny.IntOfInt64(254), 15)
			(_nw74).ArraySet1((_this).F4(), 16)
			(_nw74).ArraySet1((_dafny.Zero).Minus((_this.F5()).Times(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))), 17)
			(_nw74).ArraySet1(p1, 18)
			(_nw74).ArraySet1((_423_v0).Select((Companion_Default___.SafeIndex(_this.F5(), _dafny.IntOfUint32((_423_v0).Cardinality()))).Uint32()).(_dafny.Int), 19)
			(_nw74).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_423_v0, _423_v0)).Cardinality()), 20)
			_441_v14 = _nw74
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_441_v14), 0))
			_ = _index66
			(_441_v14).ArraySet1((p1).Minus(_this.F5()), (_index66).Int())
			var _442_v15 _dafny.Array
			_ = _442_v15
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_8
			var _nw75 _dafny.Array
			_ = _nw75
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw75 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.CodePoint = (func(_443_p0 bool) func(_dafny.Int) _dafny.CodePoint {
					return func(_444_i2 _dafny.Int) _dafny.CodePoint {
						return (func() _dafny.CodePoint {
							if _443_p0 {
								return _dafny.CodePoint('v')
							}
							return _dafny.CodePoint('f')
						})()
					}
				})(p0)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw75 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw75).ArraySet1CodePoint(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw75).ArraySet1CodePoint(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_442_v15 = _nw75
			var _445_v16 _dafny.CodePoint
			_ = _445_v16
			_445_v16 = _dafny.CodePoint('s')
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_442_v15), 0))
			_ = _index67
			(_442_v15).ArraySet1CodePoint(_445_v16, (_index67).Int())
			var _446_v17 _dafny.Map
			_ = _446_v17
			_446_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p0)
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_441_v14), 0))
			_ = _index68
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_442_v15), 0))
			_ = _index69
			var _rhs71 _dafny.Int = (_dafny.Zero).Minus((_this).F4())
			_ = _rhs71
			var _rhs72 _dafny.CodePoint = Companion_Default___.Fm4(p0, _dafny.IntOfInt64(-652), !(p0) || (false), Companion_Default___.Fm3(!(false), globalState), globalState)
			_ = _rhs72
			var _rhs73 bool = !(p0) || (p0)
			_ = _rhs73
			var _rhs74 bool = !(p0)
			_ = _rhs74
			var _rhs75 _dafny.Map = (func() _dafny.Map {
				if p0 {
					return _446_v17
				}
				return (_446_v17).Merge(Companion_Default___.Fm10(_this.F5(), p0, (_this).F4(), p1, globalState))
			})()
			_ = _rhs75
			var _lhs48 _dafny.Array = _441_v14
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_441_v14), 0))
			_ = _lhs49
			var _lhs50 _dafny.Array = _442_v15
			_ = _lhs50
			var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_442_v15), 0))
			_ = _lhs51
			var _lhs52 *GlobalState = globalState
			_ = _lhs52
			var _lhs53 *GlobalState = globalState
			_ = _lhs53
			(_lhs48).ArraySet1(_rhs71, (_lhs49).Int())
			(_lhs50).ArraySet1CodePoint(_rhs72, (_lhs51).Int())
			_lhs52.F2 = _rhs73
			_lhs53.F2 = _rhs74
			_446_v17 = _rhs75
			var _447_v18 D0
			_ = _447_v18
			_447_v18 = Companion_D0_.Create_DC1_(_this.F5(), (_441_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_441_v14), 0))).Int()).(_dafny.Int))
			_447_v18 = _447_v18
			r0 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-676), p1), (_this).F4())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_441_v14), 0))
			_ = _index70
			(_441_v14).ArraySet1((_this).F4(), (_index70).Int())
			(globalState).F2 = p0
		} else {
			(globalState).F2 = (Companion_Default___.Fm11(_dafny.IntOfInt64(-981), p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F5())), globalState)).IsSubsetOf((_424_v1).Difference(_424_v1))
			(_this).F5_set_((((_this).F4()).Minus((_dafny.Zero).Minus(p1))).Minus(_this.F5()))
			(globalState).F2 = (p1).Cmp(_dafny.IntOfInt64(194)) < 0
			var _448_v19 _dafny.Map
			_ = _448_v19
			_448_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			if (_this).Fm2(p0, _this.F5(), (func() bool {
				if (_448_v19).Contains(p0) {
					return (_448_v19).Get(p0).(bool)
				}
				return p0
			})(), globalState) {
				var _449_v20 _dafny.Array
				_ = _449_v20
				var _nw76 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
				_ = _nw76
				_449_v20 = _nw76
				_449_v20 = _449_v20
				var _450_v21 _dafny.Array
				_ = _450_v21
				var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw77
				_450_v21 = _nw77
				var _rhs76 _dafny.Int = p1
				_ = _rhs76
				var _rhs77 _dafny.Array = _450_v21
				_ = _rhs77
				r0 = _rhs76
				_450_v21 = _rhs77
				r0 = p1
				_423_v0 = _423_v0
				var _451_v22 _dafny.Array
				_ = _451_v22
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_9
				var _nw78 _dafny.Array
				_ = _nw78
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw78 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Int = (func(_452_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_453_i3 _dafny.Int) _dafny.Int {
							return (_453_i3).Times(_452_p1)
						}
					})(p1)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw78 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw78).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw78).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_451_v22 = _nw78
				var _rhs78 _dafny.Int = (_this).F4()
				_ = _rhs78
				var _rhs79 _dafny.Array = _451_v22
				_ = _rhs79
				var _rhs80 bool = p0
				_ = _rhs80
				var _rhs81 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_424_v1).IsDisjointFrom(_424_v1) {
						return _dafny.IntOfInt64(-201)
					}
					return (_dafny.Zero).Minus(((_431_v6).Intersection(_431_v6)).Cardinality())
				})())
				_ = _rhs81
				var _lhs54 *C1 = _this
				_ = _lhs54
				var _lhs55 *GlobalState = globalState
				_ = _lhs55
				var _lhs56 *C1 = _this
				_ = _lhs56
				_lhs54.F5_set_(_rhs78)
				_451_v22 = _rhs79
				_lhs55.F2 = _rhs80
				_lhs56.F5_set_(_rhs81)
			} else {
				r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_427_v2, _427_v2)).Cardinality())
				var _454_v23 _dafny.Array
				_ = _454_v23
				var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw79
				_454_v23 = _nw79
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_454_v23), 0))
				_ = _index71
				(_454_v23).ArraySet1(_423_v0, (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_454_v23), 0))
				_ = _index72
				(_454_v23).ArraySet1(_dafny.Companion_Sequence_.Update(_423_v0, (Companion_Default___.SafeIndex(Companion_Default___.Fm3(p0, globalState), _dafny.IntOfUint32((_423_v0).Cardinality()))).Uint32(), p1), (_index72).Int())
				var _455_v24 D0
				_ = _455_v24
				_455_v24 = Companion_D0_.Create_DC0_(_this.F5())
				(_this).F5_set_((_455_v24).Dtor_cf0())
				var _456_v25 _dafny.Array
				_ = _456_v25
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_10
				var _nw80 _dafny.Array
				_ = _nw80
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw80 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = (func(_457_p0 bool) func(_dafny.Int) bool {
						return func(_458_i4 _dafny.Int) bool {
							return _457_p0
						}
					})(p0)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw80 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw80).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw80).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_456_v25 = _nw80
				var _459_v26 _dafny.Map
				_ = _459_v26
				_459_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_456_v25, _456_v25), p1)
				var _460_v27 _dafny.Map
				_ = _460_v27
				_460_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _456_v25)
				var _461_v28 _dafny.Sequence
				_ = _461_v28
				_461_v28 = _dafny.SeqOf((func() _dafny.Array {
					if (_460_v27).Contains(false) {
						return (_460_v27).Get(false).(_dafny.Array)
					}
					return _456_v25
				})(), _456_v25, _456_v25, _456_v25, _456_v25)
				var _462_v29 _dafny.Sequence
				_ = _462_v29
				_462_v29 = _dafny.SeqOf(_461_v28)
				(globalState).F2 = (p1).Cmp((func() _dafny.Int {
					if (_459_v26).Contains((_462_v29).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_462_v29).Cardinality()))).Uint32()).(_dafny.Sequence)) {
						return (_459_v26).Get((_462_v29).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_462_v29).Cardinality()))).Uint32()).(_dafny.Sequence)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((p2).Cardinality())
				})()) < 0
				var _463_v30 _dafny.Set
				_ = _463_v30
				_463_v30 = _dafny.SetOf(p0)
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_456_v25), 0))
				_ = _index73
				(_456_v25).ArraySet1((_this).Fm1(!((_this).Fm2(false, p1, p0, globalState)), (_dafny.Zero).Minus(p1), _463_v30, p0, globalState), (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_456_v25), 0))
				_ = _index74
				(_456_v25).ArraySet1(p0, (_index74).Int())
			}
			r0 = Companion_Default___.SafeDivisionInt((_this).F4(), (_this).F4())
		}
		(globalState).F2 = p0
		var _464_v31 _dafny.MultiSet
		_ = _464_v31
		_464_v31 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("k"))
		_464_v31 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("nxhuqgf"), p2)
		r0 = (_this).F4()
		r1 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_465_i5 _dafny.Int) _dafny.Int {
			return (_this).F4()
		}))
		r2 = _424_v1
		r3 = _427_v2
		return r0, r1, r2, r3
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f5 _dafny.Int
	_f4 _dafny.Int
	_f6 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f5 = _dafny.Zero
	_this._f4 = _dafny.Zero
	_this._f6 = false
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

func (_this *C2) F5() _dafny.Int {
	return _this._f5
}
func (_this *C2) F5_set_(value _dafny.Int) {
	_this._f5 = value
}
func (_this *C2) F4() _dafny.Int {
	return _this._f4
}
func (_this *C2) Ctor__(f6 bool, f4 _dafny.Int, f5 _dafny.Int) {
	{
		(_this)._f6 = f6
		(_this)._f4 = f4
		(_this)._f5 = f5
	}
}
func (_this *C2) Fm0(p0 bool, p1 _dafny.MultiSet, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())))
	}
}
func (_this *C2) M1(p0 _dafny.Array, p1 _dafny.CodePoint, globalState *GlobalState) (bool, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _466_v0 _dafny.Array
		_ = _466_v0
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw81
		_466_v0 = _nw81
		for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_466_v0), 0))); ; {
			_guard_loop_0, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _467_i0 _dafny.Int
			_467_i0 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_467_i0).Sign() != -1) && ((_467_i0).Cmp(_dafny.ArrayLen((_466_v0), 0)) < 0)) {
				(_466_v0).ArraySet1((_this).F6(), (_467_i0).Int())
			}
		}
		var _468_v1 _dafny.CodePoint
		_ = _468_v1
		_468_v1 = _dafny.CodePoint('o')
		_468_v1 = _468_v1
		(_this).F5_set_(_this.F5())
		var _rhs82 _dafny.Int = _this.F5()
		_ = _rhs82
		var _rhs83 bool = (_this).F6()
		_ = _rhs83
		var _lhs57 *C2 = _this
		_ = _lhs57
		var _lhs58 *GlobalState = globalState
		_ = _lhs58
		_lhs57.F5_set_(_rhs82)
		_lhs58.F2 = _rhs83
		var _469_v2 *C0
		_ = _469_v2
		var _nw82 *C0 = New_C0_()
		_ = _nw82
		_nw82.Ctor__(_dafny.MultiSetOf(_468_v1))
		_469_v2 = _nw82
		r0 = (_this).F6()
		var _470_v3 D1
		_ = _470_v3
		_470_v3 = Companion_D1_.Create_DC4_((_this).F6())
		r0 = !((_470_v3).Dtor_cf4())
		var _471_v4 _dafny.Sequence
		_ = _471_v4
		_471_v4 = _dafny.SeqOf(Companion_Default___.Fm12((_this).F6(), (_this).F6(), (_this).F6(), globalState))
		r1 = (_471_v4).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_471_v4).Cardinality()))).Uint32()).(bool)
		r2 = (_this).F6()
		return r0, r1, r2
	}
}
func (_this *C2) M2(p0 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		if false {
			r0 = (_this).F6()
			if ((_this).F6()) && ((_this).F6()) {
				r0 = !((_this).F6())
				var _472_v0 _dafny.Sequence
				_ = _472_v0
				_472_v0 = _dafny.SeqOf((_this).F6())
				var _473_v1 _dafny.MultiSet
				_ = _473_v1
				_473_v1 = _dafny.MultiSetOf(_472_v0)
				var _474_v2 D2
				_ = _474_v2
				_474_v2 = Companion_D2_.Create_DC6_(_472_v0)
				var _475_v3 _dafny.Map
				_ = _475_v3
				_475_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())
				var _476_v4 _dafny.Sequence
				_ = _476_v4
				_476_v4 = _dafny.UnicodeSeqOfUtf8Bytes("quepdvot")
				var _477_v5 D3
				_ = _477_v5
				_477_v5 = Companion_D3_.Create_DC8_(_476_v4)
				var _478_v6 _dafny.Set
				_ = _478_v6
				_478_v6 = _dafny.SetOf((_this).F6())
				var _479_v7 _dafny.CodePoint
				_ = _479_v7
				_479_v7 = _dafny.CodePoint('v')
				var _480_v8 _dafny.Array
				_ = _480_v8
				var _nwElement0_9 _dafny.Int = _this.F5()
				_ = _nwElement0_9
				var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(18))
				_ = _nw83
				(_nw83).ArraySet1(_nwElement0_9, 0)
				(_nw83).ArraySet1((func() _dafny.Int {
					if (_473_v1).Contains(_dafny.SeqOf(!((_this).F6()))) {
						return (_473_v1).Multiplicity(_dafny.SeqOf(!((_this).F6())))
					}
					return Companion_Default___.Fm3((_this).F6(), globalState)
				})(), 1)
				(_nw83).ArraySet1(_this.F5(), 2)
				(_nw83).ArraySet1(_dafny.IntOfUint32(((_474_v2).Dtor_cf6()).Cardinality()), 3)
				(_nw83).ArraySet1((_this).F4(), 4)
				(_nw83).ArraySet1((_475_v3).Cardinality(), 5)
				(_nw83).ArraySet1((_this).F4(), 6)
				(_nw83).ArraySet1(p0, 7)
				(_nw83).ArraySet1(_this.F5(), 8)
				(_nw83).ArraySet1(_this.F5(), 9)
				(_nw83).ArraySet1(_dafny.IntOfInt64(-265), 10)
				(_nw83).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_477_v5).Dtor_cf7(), (Companion_Default___.SafeIndex((_478_v6).Cardinality(), _dafny.IntOfUint32(((_477_v5).Dtor_cf7()).Cardinality()))).Uint32(), _479_v7)).Cardinality()), 11)
				(_nw83).ArraySet1((_this).F4(), 12)
				(_nw83).ArraySet1(p0, 13)
				(_nw83).ArraySet1((_this).F4(), 14)
				(_nw83).ArraySet1(p0, 15)
				(_nw83).ArraySet1(_this.F5(), 16)
				(_nw83).ArraySet1((_this).F4(), 17)
				_480_v8 = _nw83
				var _481_v9 _dafny.Array
				_ = _481_v9
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw84
				_481_v9 = _nw84
				var _482_v10 _dafny.Map
				_ = _482_v10
				_482_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_480_v8, _481_v9)
				var _483_v11 D1
				_ = _483_v11
				_483_v11 = Companion_D1_.Create_DC3_((_this).F6())
				var _484_v12 _dafny.Map
				_ = _484_v12
				_484_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_482_v10, !(((_this).F6()) == (Companion_Default___.Fm12((_483_v11).Dtor_cf3(), (_this).F6(), (_this).F6(), globalState))))
				var _485_v13 _dafny.Map
				_ = _485_v13
				_485_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
				_484_v12 = (_484_v12).Update((_482_v10).Update(_480_v8, _481_v9), (func() bool {
					if (_485_v13).Contains((_this).F6()) {
						return (_485_v13).Get((_this).F6()).(bool)
					}
					return (_this).F6()
				})())
				_475_v3 = (_475_v3).Update(_this.F5(), (_this).F6())
				(_this).F5_set_(p0)
				var _486_v14 _dafny.MultiSet
				_ = _486_v14
				_486_v14 = _dafny.MultiSetOf(_479_v7, _479_v7)
				var _487_v15 *C0
				_ = _487_v15
				var _nw85 *C0 = New_C0_()
				_ = _nw85
				_nw85.Ctor__((_486_v14).Difference(_dafny.MultiSetOf(_479_v7)))
				_487_v15 = _nw85
			} else {
				var _488_v16 _dafny.CodePoint
				_ = _488_v16
				_488_v16 = _dafny.CodePoint('v')
				var _489_v17 _dafny.MultiSet
				_ = _489_v17
				_489_v17 = _dafny.MultiSetOf(_488_v16)
				var _490_v18 *C0
				_ = _490_v18
				var _nw86 *C0 = New_C0_()
				_ = _nw86
				_nw86.Ctor__(_489_v17)
				_490_v18 = _nw86
				(globalState).F2 = (_this).F6()
				_488_v16 = _488_v16
				var _491_v19 _dafny.Sequence
				_ = _491_v19
				_491_v19 = _dafny.UnicodeSeqOfUtf8Bytes("ngmxo")
				var _492_v20 D3
				_ = _492_v20
				_492_v20 = Companion_D3_.Create_DC8_(_491_v19)
				_491_v19 = (_492_v20).Dtor_cf7()
				_491_v19 = Companion_Default___.Fm6((_this).F6(), _this.F5(), !((_this).F6()), globalState)
			}
			(_this).F5_set_((_dafny.Zero).Minus(Companion_Default___.Fm3((_this).F6(), globalState)))
			var _493_v21 _dafny.MultiSet
			_ = _493_v21
			_493_v21 = _dafny.MultiSetOf((_this).F6(), (_this).F6(), (_this).F6(), (_this).F6(), (_this).F6())
			var _494_v22 _dafny.Sequence
			_ = _494_v22
			_494_v22 = _dafny.SeqOf((_this).F4())
			var _495_v23 _dafny.CodePoint
			_ = _495_v23
			_495_v23 = _dafny.CodePoint('t')
			var _496_v24 _dafny.MultiSet
			_ = _496_v24
			_496_v24 = _dafny.MultiSetOf(_495_v23, _495_v23, _495_v23)
			var _497_v25 *C0
			_ = _497_v25
			var _nw87 *C0 = New_C0_()
			_ = _nw87
			_nw87.Ctor__(_496_v24)
			_497_v25 = _nw87
			var _498_v26 D0
			_ = _498_v26
			_498_v26 = Companion_D0_.Create_DC1_((_this).F4(), _dafny.IntOfInt64(469))
			var _499_v27 D0
			_ = _499_v27
			_499_v27 = Companion_D0_.Create_DC0_(p0)
			var _500_v28 D1
			_ = _500_v28
			_500_v28 = Companion_D1_.Create_DC4_((_this).F6())
			var _501_v29 D3
			_ = _501_v29
			_501_v29 = Companion_D3_.Create_DC9_(p0, (_this).F6(), _this.F5(), (_500_v28).Dtor_cf4())
			var _502_v30 _dafny.Array
			_ = _502_v30
			var _nwElement0_10 _dafny.Int = _dafny.IntOfInt64(356)
			_ = _nwElement0_10
			var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(19))
			_ = _nw88
			(_nw88).ArraySet1(_nwElement0_10, 0)
			(_nw88).ArraySet1(_this.F5(), 1)
			(_nw88).ArraySet1((_493_v21).Cardinality(), 2)
			(_nw88).ArraySet1(Companion_Default___.Fm3((_this).F6(), globalState), 3)
			(_nw88).ArraySet1(_this.F5(), 4)
			(_nw88).ArraySet1((_this).F4(), 5)
			(_nw88).ArraySet1((_dafny.Zero).Minus((_494_v22).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_497_v25)).Cardinality()), _dafny.IntOfUint32((_494_v22).Cardinality()))).Uint32()).(_dafny.Int)), 6)
			(_nw88).ArraySet1(_this.F5(), 7)
			(_nw88).ArraySet1(((_498_v26).Dtor_cf1()).Plus(_this.F5()), 8)
			(_nw88).ArraySet1(_dafny.IntOfInt64(195), 9)
			(_nw88).ArraySet1((_this).F4(), 10)
			(_nw88).ArraySet1(p0, 11)
			(_nw88).ArraySet1((_498_v26).Dtor_cf2(), 12)
			(_nw88).ArraySet1((_499_v27).Dtor_cf0(), 13)
			(_nw88).ArraySet1((_494_v22).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_494_v22).Cardinality()))).Uint32()).(_dafny.Int), 14)
			(_nw88).ArraySet1(_this.F5(), 15)
			(_nw88).ArraySet1((_501_v29).Dtor_cf10(), 16)
			(_nw88).ArraySet1(_this.F5(), 17)
			(_nw88).ArraySet1((_this).F4(), 18)
			_502_v30 = _nw88
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_502_v30), 0))
			_ = _index75
			(_502_v30).ArraySet1((_this).F4(), (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_502_v30), 0))
			_ = _index76
			(_502_v30).ArraySet1((_this).F4(), (_index76).Int())
			var _503_v31 _dafny.Sequence
			_ = _503_v31
			_503_v31 = _dafny.SeqOf((_this).F6(), (_this).F6())
			var _504_v32 _dafny.Map
			_ = _504_v32
			_504_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_502_v30, _503_v31)
			_493_v21 = _dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if (_504_v32).Contains(_502_v30) {
					return (_504_v32).Get(_502_v30).(_dafny.Sequence)
				}
				return _503_v31
			})())
		} else {
			(globalState).F2 = (_this).F6()
			var _505_v33 _dafny.CodePoint
			_ = _505_v33
			_505_v33 = _dafny.CodePoint('e')
			var _506_v34 _dafny.MultiSet
			_ = _506_v34
			_506_v34 = _dafny.MultiSetOf(_505_v33)
			var _507_v35 *C0
			_ = _507_v35
			var _nw89 *C0 = New_C0_()
			_ = _nw89
			_nw89.Ctor__(_506_v34)
			_507_v35 = _nw89
			var _508_v36 D1
			_ = _508_v36
			_508_v36 = Companion_D1_.Create_DC4_((_this).F6())
			var _509_v37 _dafny.Array
			_ = _509_v37
			var _nwElement0_11 D1 = _508_v36
			_ = _nwElement0_11
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(10))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_11, 0)
			(_nw90).ArraySet1(_508_v36, 1)
			(_nw90).ArraySet1(_508_v36, 2)
			(_nw90).ArraySet1(_508_v36, 3)
			(_nw90).ArraySet1(_508_v36, 4)
			(_nw90).ArraySet1(_508_v36, 5)
			(_nw90).ArraySet1(_508_v36, 6)
			(_nw90).ArraySet1(func(_pat_let8_0 D1) D1 {
				return func(_510_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let9_0 bool) D1 {
						return func(_511_dt__update_hcf4_h0 bool) D1 {
							return Companion_D1_.Create_DC4_(_511_dt__update_hcf4_h0)
						}(_pat_let9_0)
					}((_this).F6())
				}(_pat_let8_0)
			}(_508_v36), 7)
			(_nw90).ArraySet1(_508_v36, 8)
			(_nw90).ArraySet1(Companion_D1_.Create_DC4_((_this).F6()), 9)
			_509_v37 = _nw90
			var _512_v38 _dafny.Array
			_ = _512_v38
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw91
			_512_v38 = _nw91
			var _513_v39 _dafny.Sequence
			_ = _513_v39
			_513_v39 = _dafny.SeqOf(_512_v38)
			var _rhs84 *C0 = _507_v35
			_ = _rhs84
			var _rhs85 _dafny.Array = _509_v37
			_ = _rhs85
			var _rhs86 _dafny.Array = (_513_v39).Select((Companion_Default___.SafeIndex(((_this).F4()).Minus(Companion_Default___.Fm3((_this).F6(), globalState)), _dafny.IntOfUint32((_513_v39).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs86
			var _rhs87 bool = ((_this).F6()) && ((_this).F6())
			_ = _rhs87
			var _rhs88 _dafny.Int = ((_dafny.Zero).Minus((p0).Times((_this).F4()))).Minus((_this).F4())
			_ = _rhs88
			var _lhs59 *GlobalState = globalState
			_ = _lhs59
			var _lhs60 *C2 = _this
			_ = _lhs60
			_507_v35 = _rhs84
			_509_v37 = _rhs85
			_512_v38 = _rhs86
			_lhs59.F2 = _rhs87
			_lhs60.F5_set_(_rhs88)
			r0 = !(!((_this).F6()))
			var _514_v40 _dafny.Set
			_ = _514_v40
			_514_v40 = _dafny.SetOf((_this).F6())
			(globalState).F2 = !(_514_v40).Equals(_514_v40)
			var _515_v41 T0
			_ = _515_v41
			var _nw92 *C1 = New_C1_()
			_ = _nw92
			_nw92.Ctor__(p0, (_this).F4())
			_515_v41 = _nw92
			var _516_v42 _dafny.Array
			_ = _516_v42
			var _nwElement0_12 T0 = (func() T0 {
				if (_this).F6() {
					return _515_v41
				}
				return _515_v41
			})()
			_ = _nwElement0_12
			var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(7))
			_ = _nw93
			(_nw93).ArraySet1(_nwElement0_12, 0)
			(_nw93).ArraySet1(_515_v41, 1)
			(_nw93).ArraySet1(_515_v41, 2)
			(_nw93).ArraySet1(_515_v41, 3)
			(_nw93).ArraySet1(_515_v41, 4)
			(_nw93).ArraySet1(_515_v41, 5)
			(_nw93).ArraySet1(_515_v41, 6)
			_516_v42 = _nw93
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_516_v42), 0))
			_ = _index77
			(_516_v42).ArraySet1(_515_v41, (_index77).Int())
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_516_v42), 0))
			_ = _index78
			(_516_v42).ArraySet1(_515_v41, (_index78).Int())
		}
		var _517_v43 _dafny.Array
		_ = _517_v43
		var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
		_ = _nw94
		_517_v43 = _nw94
		var _518_v44 _dafny.Array
		_ = _518_v44
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_11
		var _nw95 _dafny.Array
		_ = _nw95
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw95 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = (func(_519_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_520_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_520_i0, _519_p0)
				}
			})(p0)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw95 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw95).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw95).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_518_v44 = _nw95
		var _521_v45 _dafny.CodePoint
		_ = _521_v45
		_521_v45 = _dafny.CodePoint('q')
		var _522_v46 _dafny.Map
		_ = _522_v46
		_522_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_518_v44, _521_v45)
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_517_v43), 0))
		_ = _index79
		(_517_v43).ArraySet1(_522_v46, (_index79).Int())
		var _523_v47 _dafny.Map
		_ = _523_v47
		_523_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-702))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_524_i1 _dafny.Int) _dafny.Int {
			return (_this).F4()
		})), _522_v46)
		var _525_v48 _dafny.Sequence
		_ = _525_v48
		_525_v48 = _dafny.UnicodeSeqOfUtf8Bytes("spwqw")
		var _526_v49 _dafny.Map
		_ = _526_v49
		_526_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _525_v48)
		var _527_v50 _dafny.Sequence
		_ = _527_v50
		_527_v50 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qsprbn")).Cardinality())), p0, (_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_526_v49).Contains((_this).F6()) {
				return (_526_v49).Get((_this).F6()).(_dafny.Sequence)
			}
			return _525_v48
		})()).Cardinality())), _dafny.IntOfInt64(703), (_this).F4())
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_517_v43), 0))
		_ = _index80
		(_517_v43).ArraySet1((func() _dafny.Map {
			if (_523_v47).Contains(_527_v50) {
				return (_523_v47).Get(_527_v50).(_dafny.Map)
			}
			return _522_v46
		})(), (_index80).Int())
		if (_this).F6() {
			var _528_v51 _dafny.Sequence
			_ = _528_v51
			_528_v51 = _dafny.SeqOf((_this).F6())
			var _529_v52 _dafny.Array
			_ = _529_v52
			var _nwElement0_13 bool = Companion_Default___.Fm12((_this).F6(), Companion_Default___.Fm12((_this).F6(), (_this).F6(), (_this).F6(), globalState), (_this).F6(), globalState)
			_ = _nwElement0_13
			var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(10))
			_ = _nw96
			(_nw96).ArraySet1(_nwElement0_13, 0)
			(_nw96).ArraySet1((_this).F6(), 1)
			(_nw96).ArraySet1((func(_pat_let10_0 D1) D1 {
				return func(_530_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let11_0 bool) D1 {
						return func(_531_dt__update_hcf3_h0 bool) D1 {
							return Companion_D1_.Create_DC3_(_531_dt__update_hcf3_h0)
						}(_pat_let11_0)
					}((_this).F6())
				}(_pat_let10_0)
			}(Companion_D1_.Create_DC3_((_this).F6()))).Dtor_cf3(), 2)
			(_nw96).ArraySet1(true, 3)
			(_nw96).ArraySet1((_this).F6(), 4)
			(_nw96).ArraySet1((_528_v51).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_528_v51).Cardinality()))).Uint32()).(bool), 5)
			(_nw96).ArraySet1((_this).F6(), 6)
			(_nw96).ArraySet1((_this).F6(), 7)
			(_nw96).ArraySet1((_this).F6(), 8)
			(_nw96).ArraySet1(!((_this).F6()) || (!((_this).F6())), 9)
			_529_v52 = _nw96
			var _532_v53 D1
			_ = _532_v53
			_532_v53 = Companion_D1_.Create_DC4_((_this).F6())
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
			_ = _index81
			(_529_v52).ArraySet1((_532_v53).Dtor_cf4(), (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
			_ = _index82
			var _rhs89 _dafny.Int = _this.F5()
			_ = _rhs89
			var _rhs90 _dafny.Array = _529_v52
			_ = _rhs90
			var _rhs91 bool = (_this).F6()
			_ = _rhs91
			var _lhs61 _dafny.Array = _529_v52
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
			_ = _lhs62
			r1 = _rhs89
			_529_v52 = _rhs90
			(_lhs61).ArraySet1(_rhs91, (_lhs62).Int())
			if (_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool) {
				var _533_v54 _dafny.MultiSet
				_ = _533_v54
				_533_v54 = _dafny.MultiSetOf(_521_v45)
				var _534_v55 _dafny.Sequence
				_ = _534_v55
				_534_v55 = _dafny.SeqOf(_533_v54)
				var _535_v56 *C0
				_ = _535_v56
				var _nw97 *C0 = New_C0_()
				_ = _nw97
				_nw97.Ctor__((_534_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.IntOfUint32((_534_v55).Cardinality()))).Uint32()).(_dafny.MultiSet))
				_535_v56 = _nw97
				var _536_v57 _dafny.Sequence
				_ = _536_v57
				_536_v57 = _dafny.SeqOf(_535_v56, _535_v56)
				var _537_v58 _dafny.Map
				_ = _537_v58
				_537_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool), (_536_v57).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_536_v57).Cardinality()))).Uint32()).(*C0))
				_537_v58 = (_537_v58).Update(true, _535_v56)
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
				_ = _index83
				(_529_v52).ArraySet1((_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool), (_index83).Int())
				r1 = (Companion_Default___.Fm3(!((_this).F6()), globalState)).Plus((_this).F4())
				var _538_v59 _dafny.MultiSet
				_ = _538_v59
				_538_v59 = _dafny.MultiSetOf(_535_v56)
				var _539_v60 _dafny.Map
				_ = _539_v60
				_539_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool) {
						return p0
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}((func(_540_v45 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_541_i2 _dafny.Int) _dafny.CodePoint {
							return _540_v45
						}
					})(_521_v45)))).Cardinality())
				})(), (_this).F4())
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
				_ = _index84
				var _rhs92 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F4()))
				_ = _rhs92
				var _rhs93 _dafny.MultiSet = _dafny.MultiSetOf(_535_v56, _535_v56, _535_v56, _535_v56, _535_v56)
				_ = _rhs93
				var _rhs94 bool = !((_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool))
				_ = _rhs94
				var _rhs95 _dafny.Map = (_539_v60).Merge(_539_v60)
				_ = _rhs95
				var _rhs96 bool = Companion_Default___.Fm12((_this.F5()).Cmp(p0) >= 0, (_this).F6(), (_this).F6(), globalState)
				_ = _rhs96
				var _lhs63 _dafny.Array = _529_v52
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
				_ = _lhs64
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				r1 = _rhs92
				_538_v59 = _rhs93
				(_lhs63).ArraySet1(_rhs94, (_lhs64).Int())
				_539_v60 = _rhs95
				_lhs65.F2 = _rhs96
				var _542_v61 _dafny.MultiSet
				_ = _542_v61
				_542_v61 = _dafny.MultiSetOf(_dafny.IntOfInt64(-358))
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
				_ = _index85
				(_529_v52).ArraySet1(((_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool)) || ((_542_v61).IsProperSubsetOf(_542_v61)), (_index85).Int())
			} else {
				(globalState).F2 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_525_v48, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_525_v48).Cardinality()))).Uint32(), _dafny.CodePoint('n')), Companion_Default___.Fm6((_this).F6(), p0, (_this).F6(), globalState))
				var _543_v62 _dafny.Map
				_ = _543_v62
				_543_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool))
				_543_v62 = (func() _dafny.Map {
					if (_this).F6() {
						return _543_v62
					}
					return _543_v62
				})()
				var _544_v63 _dafny.Array
				_ = _544_v63
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
				_ = _nw98
				_544_v63 = _nw98
				var _545_v64 _dafny.Set
				_ = _545_v64
				_545_v64 = _dafny.SetOf((_this).F4(), (_this).F4(), (_this).F4())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_544_v63), 0))
				_ = _index86
				(_544_v63).ArraySet1(_dafny.MultiSetOf(_545_v64), (_index86).Int())
				var _546_v65 _dafny.MultiSet
				_ = _546_v65
				_546_v65 = _dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfUint32((_525_v48).Cardinality()), _dafny.IntOfInt64(548), (Companion_D3_.Create_DC9_((_this).F4(), (_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool), p0, (_this).F6())).Dtor_cf8()))
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_544_v63), 0))
				_ = _index87
				(_544_v63).ArraySet1(_546_v65, (_index87).Int())
				_545_v64 = (_545_v64).Intersection(_545_v64)
				r1 = ((_this.F5()).Minus(_this.F5())).Plus(Companion_Default___.Fm3((_this).F6(), globalState))
			}
			if (_this).F6() {
				(_this).F5_set_((_this).F4())
				(_this).F5_set_((_dafny.Zero).Minus((_this).F4()))
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
				_ = _index88
				(_529_v52).ArraySet1((_this).F6(), (_index88).Int())
				_521_v45 = _521_v45
				(globalState).F2 = (_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool)
			} else {
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
				_ = _index89
				(_529_v52).ArraySet1((_this).F6(), (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))
				_ = _index90
				(_529_v52).ArraySet1((_this).F6(), (_index90).Int())
				(_this).F5_set_((_dafny.Zero).Minus(_this.F5()))
				var _547_v66 _dafny.Map
				_ = _547_v66
				_547_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(_dafny.SeqOf((_this).F6(), (_529_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_529_v52), 0))).Int()).(bool))), (_this).F6())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_518_v44), 0))
				_ = _index91
				(_518_v44).ArraySet1(((_dafny.Zero).Minus((_547_v66).Cardinality())).Times(_this.F5()), (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_518_v44), 0))
				_ = _index92
				(_518_v44).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_this).F4()), (_index92).Int())
				var _548_v67 T0
				_ = _548_v67
				var _nw99 *C1 = New_C1_()
				_ = _nw99
				_nw99.Ctor__((_this).F4(), (_this).F4())
				_548_v67 = _nw99
				var _549_v68 _dafny.Sequence
				_ = _549_v68
				_549_v68 = _dafny.SeqOf(_548_v67, _548_v67)
				var _550_v69 _dafny.Map
				_ = _550_v69
				_550_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.IntOfUint32((_549_v68).Cardinality()))
				var _551_v70 _dafny.Set
				_ = _551_v70
				_551_v70 = _dafny.SetOf((_550_v69).Cardinality())
				var _552_v72 _dafny.Sequence
				_ = _552_v72
				_552_v72 = _dafny.SeqOf(_551_v70, func() _dafny.Set {
					var _coll12 = _dafny.NewBuilder()
					_ = _coll12
					for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(390), _dafny.IntOfInt64(552))); ; {
						_compr_12, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _553_v71 _dafny.Int
						_553_v71 = interface{}(_compr_12).(_dafny.Int)
						if ((_dafny.IntOfInt64(390)).Cmp(_553_v71) <= 0) && ((_553_v71).Cmp(_dafny.IntOfInt64(552)) < 0) {
							_coll12.Add((_553_v71).Minus(_this.F5()))
						}
					}
					return _coll12.ToSet()
				}())
				(globalState).F2 = !(((_552_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jucpayd")).Cardinality()), _dafny.IntOfUint32((_552_v72).Cardinality()))).Uint32()).(_dafny.Set)).IsDisjointFrom(_551_v70))
			}
			(_this).F5_set_((_dafny.Zero).Minus(p0))
			r1 = (_this).F4()
		} else {
			r1 = (p0).Plus(((_this).F4()).Minus(_dafny.IntOfInt64(209)))
			var _554_v73 _dafny.MultiSet
			_ = _554_v73
			_554_v73 = _dafny.MultiSetOf(_521_v45)
			r1 = (((_554_v73).Difference(_554_v73)).Cardinality()).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(615), (_dafny.MultiSetOf((_this).F6(), (_this).F6())).Cardinality()))
			var _555_v74 D3
			_ = _555_v74
			_555_v74 = Companion_D3_.Create_DC9_((_this).F4(), (_this).F6(), (_this).F4(), (_this).F6())
			var _556_v75 _dafny.Sequence
			_ = _556_v75
			_556_v75 = _dafny.SeqOf(true, (_555_v74).Dtor_cf9())
			r0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_556_v75, _556_v75), _556_v75)
			var _557_v76 _dafny.Array
			_ = _557_v76
			var _nwElement0_14 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_this).F6() {
					return _521_v45
				}
				return _521_v45
			})()
			_ = _nwElement0_14
			var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(10))
			_ = _nw100
			(_nw100).ArraySet1CodePoint(_nwElement0_14, 0)
			(_nw100).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_this).F6() {
					return _521_v45
				}
				return _521_v45
			})(), 1)
			(_nw100).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_this).F6() {
					return (_525_v48).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_525_v48).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
				return _521_v45
			})(), 2)
			(_nw100).ArraySet1CodePoint(_521_v45, 3)
			(_nw100).ArraySet1CodePoint(_521_v45, 4)
			(_nw100).ArraySet1CodePoint(_521_v45, 5)
			(_nw100).ArraySet1CodePoint(_521_v45, 6)
			(_nw100).ArraySet1CodePoint(_521_v45, 7)
			(_nw100).ArraySet1CodePoint(_521_v45, 8)
			(_nw100).ArraySet1CodePoint(_521_v45, 9)
			_557_v76 = _nw100
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_557_v76), 0))
			_ = _index93
			(_557_v76).ArraySet1CodePoint(_521_v45, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_557_v76), 0))
			_ = _index94
			(_557_v76).ArraySet1CodePoint(_521_v45, (_index94).Int())
			(_this).F5_set_(Companion_Default___.Fm3((_this).F6(), globalState))
		}
		r0 = (_this).F6()
		r0 = (_this).F6()
		var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_518_v44), 0))
		_ = _index95
		(_518_v44).ArraySet1((_this.F5()).Times(p0), (_index95).Int())
		var _558_v77 _dafny.Array
		_ = _558_v77
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_12
		var _nw101 _dafny.Array
		_ = _nw101
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw101 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Set = (func(_559_v45 _dafny.CodePoint) func(_dafny.Int) _dafny.Set {
				return func(_560_i3 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg19 _dafny.Int) interface{} {
							return coer19(arg19)
						}
					}((func(_561_v45 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_562_i4 _dafny.Int) _dafny.CodePoint {
							return _561_v45
						}
					})(_559_v45)))).Cardinality()), (_this).F4(), _dafny.IntOfInt64(-661), (_this).F4())
				}
			})(_521_v45)
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
		_558_v77 = _nw101
		var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_518_v44), 0))
		_ = _index96
		var _rhs97 _dafny.Int = (_527_v50).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_527_v50).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs97
		var _rhs98 _dafny.Array = _558_v77
		_ = _rhs98
		var _lhs66 _dafny.Array = _518_v44
		_ = _lhs66
		var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_518_v44), 0))
		_ = _lhs67
		(_lhs66).ArraySet1(_rhs97, (_lhs67).Int())
		_558_v77 = _rhs98
		r0 = (_this).F6()
		r1 = (_this.F5()).Plus((_518_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_518_v44), 0))).Int()).(_dafny.Int))
		return r0, r1
	}
}
func (_this *C2) F6() bool {
	{
		return _this._f6
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f3 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f3 = false
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__(f3 bool) {
	{
		(_this)._f3 = f3
	}
}
func (_this *C3) M0(p0 _dafny.Array, p1 _dafny.Int, p2 bool, p3 _dafny.Map, globalState *GlobalState) (_dafny.Sequence, _dafny.Map) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _563_v0 *C2
		_ = _563_v0
		var _nw102 *C2 = New_C2_()
		_ = _nw102
		_nw102.Ctor__(Companion_Default___.Fm12(p2, p2, (_this).F3(), globalState), p1, p1)
		_563_v0 = _nw102
		var _564_v1 _dafny.Sequence
		_ = _564_v1
		_564_v1 = _dafny.UnicodeSeqOfUtf8Bytes("iqwgvcpna")
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_564_v1, _564_v1) {
			var _565_v2 _dafny.Sequence
			_ = _565_v2
			_565_v2 = _dafny.SeqOf(!((_563_v0).F6()), Companion_Default___.Fm12((_563_v0).F6(), p2, (_563_v0).F6(), globalState), true)
			(globalState).F2 = (_565_v2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_565_v2).Cardinality()))).Uint32()).(bool)
			if (func() bool {
				if p2 {
					return (_563_v0).F6()
				}
				return p2
			})() {
				var _566_v3 _dafny.Map
				_ = _566_v3
				_566_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), (_this).F3())
				var _567_v4 _dafny.MultiSet
				_ = _567_v4
				_567_v4 = _dafny.MultiSetOf(_566_v3)
				var _568_v5 _dafny.Sequence
				_ = _568_v5
				_568_v5 = _dafny.SeqOf((_dafny.Zero).Minus((_567_v4).Cardinality()))
				var _569_v6 _dafny.MultiSet
				_ = _569_v6
				_569_v6 = _dafny.MultiSetOf((_563_v0).F6(), (_563_v0).F6(), _dafny.Companion_Sequence_.IsProperPrefixOf(_568_v5, _568_v5), (p1).Cmp(_dafny.IntOfUint32((_564_v1).Cardinality())) <= 0, !((_563_v0).F6()))
				var _570_v7 _dafny.Sequence
				_ = _570_v7
				_570_v7 = _dafny.SeqOf(_565_v2)
				_569_v6 = (_569_v6).Union((_569_v6).Update(p2, Companion_Default___.Abs(_dafny.IntOfUint32((_570_v7).Cardinality()))))
				var _571_v8 _dafny.CodePoint
				_ = _571_v8
				_571_v8 = _dafny.CodePoint('r')
				_571_v8 = (func() _dafny.CodePoint {
					if p2 {
						return _571_v8
					}
					return _dafny.CodePoint('d')
				})()
				var _572_v9 _dafny.Map
				_ = _572_v9
				_572_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_571_v8, _564_v1)
				_572_v9 = (_572_v9).Update(_571_v8, _564_v1)
				var _573_v10 _dafny.Set
				_ = _573_v10
				_573_v10 = _dafny.SetOf(p2)
				(globalState).F2 = ((_573_v10).Difference(_573_v10)).IsDisjointFrom(_573_v10)
				var _574_v11 _dafny.Int
				_ = _574_v11
				_574_v11 = _dafny.IntOfInt64(-107)
				_574_v11 = p1
			} else {
				var _575_v12 _dafny.Array
				_ = _575_v12
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
				_ = _nw103
				_575_v12 = _nw103
				var _576_v13 _dafny.CodePoint
				_ = _576_v13
				_576_v13 = _dafny.CodePoint('p')
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_575_v12), 0))
				_ = _index97
				(_575_v12).ArraySet1CodePoint(_576_v13, (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_575_v12), 0))
				_ = _index98
				(_575_v12).ArraySet1CodePoint(_576_v13, (_index98).Int())
				var _577_v14 _dafny.Map
				_ = _577_v14
				_577_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_564_v1).Cardinality()))
				var _578_v15 _dafny.Map
				_ = _578_v15
				_578_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
				var _579_v16 _dafny.Set
				_ = _579_v16
				_579_v16 = _dafny.SetOf((_563_v0).F6())
				var _580_v17 D4
				_ = _580_v17
				_580_v17 = Companion_D4_.Create_DC12_(_579_v16)
				_577_v14 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm13((_563_v0).F6(), (_578_v15).Cardinality(), p1, globalState)).Cardinality(), ((_580_v17).Dtor_cf14()).Cardinality())).Merge(_577_v14)
				var _581_v18 _dafny.Int
				_ = _581_v18
				_581_v18 = _dafny.IntOfInt64(822)
				var _582_v19 _dafny.Map
				_ = _582_v19
				_582_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_563_v0).F6(), _581_v18)
				var _583_v20 _dafny.Sequence
				_ = _583_v20
				_583_v20 = _dafny.SeqOf(_581_v18)
				var _584_v21 T0
				_ = _584_v21
				var _nw104 *C2 = New_C2_()
				_ = _nw104
				_nw104.Ctor__((_563_v0).F6(), _dafny.IntOfUint32((_583_v20).Cardinality()), _581_v18)
				_584_v21 = _nw104
				var _585_v22 _dafny.Map
				_ = _585_v22
				_585_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_581_v18, _584_v21)
				_581_v18 = (func() _dafny.Int {
					if (_582_v19).Contains((_563_v0).F6()) {
						return (_582_v19).Get((_563_v0).F6()).(_dafny.Int)
					}
					return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm6(false, (_585_v22).Cardinality(), (_this).F3(), globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.IntOfUint32((Companion_Default___.Fm6(false, (_585_v22).Cardinality(), (_this).F3(), globalState)).Cardinality()))).Uint32(), (_575_v12).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_575_v12), 0))).Int()))).Cardinality())).Times(p1)
				})()
				var _586_v23 _dafny.Array
				_ = _586_v23
				var _nwElement0_15 bool = p2
				_ = _nwElement0_15
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(26))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_15, 0)
				(_nw105).ArraySet1((_563_v0).F6(), 1)
				(_nw105).ArraySet1((_563_v0).F6(), 2)
				(_nw105).ArraySet1((_563_v0).F6(), 3)
				(_nw105).ArraySet1((_563_v0).F6(), 4)
				(_nw105).ArraySet1((_563_v0).F6(), 5)
				(_nw105).ArraySet1((_563_v0).F6(), 6)
				(_nw105).ArraySet1((_563_v0).F6(), 7)
				(_nw105).ArraySet1((_563_v0).F6(), 8)
				(_nw105).ArraySet1(p2, 9)
				(_nw105).ArraySet1(p2, 10)
				(_nw105).ArraySet1((_563_v0).F6(), 11)
				(_nw105).ArraySet1(p2, 12)
				(_nw105).ArraySet1((_563_v0).F6(), 13)
				(_nw105).ArraySet1(p2, 14)
				(_nw105).ArraySet1((_563_v0).F6(), 15)
				(_nw105).ArraySet1(!(p2), 16)
				(_nw105).ArraySet1((_563_v0).F6(), 17)
				(_nw105).ArraySet1(p2, 18)
				(_nw105).ArraySet1((_this).F3(), 19)
				(_nw105).ArraySet1((_this).F3(), 20)
				(_nw105).ArraySet1((_563_v0).F6(), 21)
				(_nw105).ArraySet1((_563_v0).F6(), 22)
				(_nw105).ArraySet1((_563_v0).F6(), 23)
				(_nw105).ArraySet1((_563_v0).F6(), 24)
				(_nw105).ArraySet1((_563_v0).F6(), 25)
				_586_v23 = _nw105
				var _587_v24 _dafny.Sequence
				_ = _587_v24
				_587_v24 = _dafny.SeqOf(_586_v23, _586_v23)
				var _588_v25 _dafny.Sequence
				_ = _588_v25
				_588_v25 = _dafny.SeqOf(_587_v24)
				var _589_v26 _dafny.Sequence
				_ = _589_v26
				_589_v26 = _dafny.SeqOf(_587_v24, _dafny.Companion_Sequence_.Update((_588_v25).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_588_v25).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_584_v21).F4(), _dafny.IntOfUint32(((_588_v25).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_588_v25).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _586_v23))
				var _590_v27 _dafny.Map
				_ = _590_v27
				_590_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_576_v13, _dafny.IntOfUint32(((_589_v26).Select((Companion_Default___.SafeIndex(_584_v21.F5(), _dafny.IntOfUint32((_589_v26).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
				var _591_v28 _dafny.Array
				_ = _591_v28
				var _nwElement0_16 bool = p2
				_ = _nwElement0_16
				var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(3))
				_ = _nw106
				(_nw106).ArraySet1(_nwElement0_16, 0)
				(_nw106).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}((func(_592_v13 _dafny.CodePoint, _593_v20 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
					return func(_594_i0 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_592_v13, _dafny.IntOfUint32((_593_v20).Cardinality()))
					}
				})(_576_v13, _583_v20))), _590_v27), 1)
				(_nw106).ArraySet1(!(((_584_v21).F4()).Cmp(p1) != 0), 2)
				_591_v28 = _nw106
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_575_v12), 0))
				_ = _index99
				var _rhs99 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (func() bool {
						if (_563_v0).F6() {
							return (_563_v0).F6()
						}
						return p2
					})() {
						return (_564_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_564_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
					return _576_v13
				})()
				_ = _rhs99
				var _rhs100 _dafny.Array = (Companion_D5_.Create_DC15_(_591_v28)).Dtor_cf19()
				_ = _rhs100
				var _rhs101 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm3(p2, globalState))
				_ = _rhs101
				var _lhs68 _dafny.Array = _575_v12
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_575_v12), 0))
				_ = _lhs69
				(_lhs68).ArraySet1CodePoint(_rhs99, (_lhs69).Int())
				_586_v23 = _rhs100
				_581_v18 = _rhs101
				_585_v22 = (_585_v22).Update((p1).Plus((_dafny.Zero).Minus(p1)), _584_v21)
			}
			var _595_v29 _dafny.Int
			_ = _595_v29
			_595_v29 = _dafny.IntOfInt64(446)
			var _596_v30 _dafny.MultiSet
			_ = _596_v30
			_596_v30 = _dafny.MultiSetOf(!((_this).F3()))
			_595_v29 = (_596_v30).Cardinality()
			var _597_v31 _dafny.Sequence
			_ = _597_v31
			_597_v31 = _dafny.SeqOf((func() _dafny.Int {
				if (_563_v0).F6() {
					return p1
				}
				return p1
			})())
			_595_v29 = (_dafny.Zero).Minus((_597_v31).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_597_v31).Cardinality()))).Uint32()).(_dafny.Int))
			(globalState).F2 = !((p1).Cmp((func() _dafny.Int {
				if (_563_v0).F6() {
					return p1
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_564_v1).Cardinality()))
			})()) < 0)
		} else {
			var _598_v32 T0
			_ = _598_v32
			var _nw107 *C1 = New_C1_()
			_ = _nw107
			_nw107.Ctor__(p1, p1)
			_598_v32 = _nw107
			var _599_v33 _dafny.Map
			_ = _599_v33
			_599_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _598_v32)
			var _600_v34 _dafny.Array
			_ = _600_v34
			var _nw108 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw108
			_600_v34 = _nw108
			var _601_v35 _dafny.Sequence
			_ = _601_v35
			_601_v35 = _dafny.SeqOf(_600_v34)
			var _602_v36 D4
			_ = _602_v36
			_602_v36 = Companion_D4_.Create_DC14_((_599_v33).Merge(_599_v33), p2, _601_v35)
			var _pat_let_tv11 = _599_v33
			_ = _pat_let_tv11
			_602_v36 = func(_pat_let12_0 D4) D4 {
				return func(_603_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let13_0 _dafny.Map) D4 {
						return func(_604_dt__update_hcf16_h0 _dafny.Map) D4 {
							return func(_pat_let14_0 bool) D4 {
								return func(_605_dt__update_hcf17_h0 bool) D4 {
									return Companion_D4_.Create_DC14_(_604_dt__update_hcf16_h0, _605_dt__update_hcf17_h0, (_603_dt__update__tmp_h0).Dtor_cf18())
								}(_pat_let14_0)
							}((_this).F3())
						}(_pat_let13_0)
					}(_pat_let_tv11)
				}(_pat_let12_0)
			}(_602_v36)
			(_598_v32).F5_set_(_598_v32.F5())
			var _606_v37 _dafny.Array
			_ = _606_v37
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_13
			var _nw109 _dafny.Array
			_ = _nw109
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw109 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.MultiSet = (func(_607_v0 *C2, _608_p2 bool) func(_dafny.Int) _dafny.MultiSet {
					return func(_609_i1 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf((_607_v0).F6(), _608_p2)
					}
				})(_563_v0, p2)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw109 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw109).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw109).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_606_v37 = _nw109
			var _610_v38 _dafny.MultiSet
			_ = _610_v38
			_610_v38 = _dafny.MultiSetOf((_563_v0).F6())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_606_v37), 0))
			_ = _index100
			(_606_v37).ArraySet1((_610_v38).Intersection(_610_v38), (_index100).Int())
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_606_v37), 0))
			_ = _index101
			(_606_v37).ArraySet1(_dafny.MultiSetOf(!((_563_v0).F6()), (_this).F3()), (_index101).Int())
			var _611_v39 _dafny.Sequence
			_ = _611_v39
			_611_v39 = _dafny.SeqOf(p2)
			var _612_v40 D2
			_ = _612_v40
			_612_v40 = Companion_D2_.Create_DC6_(_611_v39)
			var _source8 D2 = _612_v40
			_ = _source8
			if _source8.Is_DC7() {
				var _613_v41 _dafny.Array
				_ = _613_v41
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw110
				_613_v41 = _nw110
				var _614_v42 _dafny.CodePoint
				_ = _614_v42
				_614_v42 = _dafny.CodePoint('e')
				var _615_v43 _dafny.Map
				_ = _615_v43
				_615_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_v42, Companion_Default___.Fm6(p2, _598_v32.F5(), !((_this).F3()), globalState))
				var _616_v44 *C1
				_ = _616_v44
				var _nw111 *C1 = New_C1_()
				_ = _nw111
				_nw111.Ctor__(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_615_v43).Contains(_614_v42) {
						return (_615_v43).Get(_614_v42).(_dafny.Sequence)
					}
					return _564_v1
				})()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(648))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}((func(_617_v32 T0, _618_v1 _dafny.Sequence, _619_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_620_i2 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(502), (_617_v32).F4(), _dafny.IntOfUint32((_618_v1).Cardinality()), _619_p1)).Cardinality())
					}
				})(_598_v32, _564_v1, p1)))).Cardinality()))
				_616_v44 = _nw111
				var _621_v45 _dafny.Map
				_ = _621_v45
				_621_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_616_v44, _598_v32.F5())
				var _rhs102 _dafny.Array = _600_v34
				_ = _rhs102
				var _rhs103 _dafny.Array = _613_v41
				_ = _rhs103
				var _rhs104 bool = ((func() _dafny.Int {
					if (_621_v45).Contains(_616_v44) {
						return (_621_v45).Get(_616_v44).(_dafny.Int)
					}
					return _dafny.IntOfInt64(372)
				})()).Cmp(p1) != 0
				_ = _rhs104
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				_600_v34 = _rhs102
				_613_v41 = _rhs103
				_lhs70.F2 = _rhs104
				var _622_v46 _dafny.Array
				_ = _622_v46
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(17))
				_ = _nw112
				_622_v46 = _nw112
				var _623_v47 _dafny.Map
				_ = _623_v47
				_623_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_622_v46, p1)
				var _624_v48 _dafny.Map
				_ = _624_v48
				_624_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_563_v0).F6(), p2)
				var _625_v49 _dafny.Map
				_ = _625_v49
				_625_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_623_v47).Merge(_623_v47), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_563_v0).F6(), p2)).Merge(_624_v48))
				_625_v49 = (_625_v49).Update(_623_v47, _624_v48)
				(globalState).F2 = (_563_v0).F6()
				var _626_v50 _dafny.Array
				_ = _626_v50
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
				_ = _nw113
				_626_v50 = _nw113
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_626_v50), 0))
				_ = _index102
				(_626_v50).ArraySet1((_624_v48).Update(true, false), (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_626_v50), 0))
				_ = _index103
				(_626_v50).ArraySet1((Companion_Default___.Fm14(_598_v32.F5(), globalState)).Merge(_624_v48), (_index103).Int())
			} else {
				var _627___mcc_h0 _dafny.Sequence = _source8.Get_().(D2_DC6).Cf6
				_ = _627___mcc_h0
				var _628_cf6 _dafny.Sequence = _627___mcc_h0
				_ = _628_cf6
				var _629_v51 *C2
				_ = _629_v51
				var _nw114 *C2 = New_C2_()
				_ = _nw114
				_nw114.Ctor__(p2, (_dafny.Zero).Minus(_598_v32.F5()), _dafny.IntOfInt64(645))
				_629_v51 = _nw114
				var _630_v52 D3
				_ = _630_v52
				_630_v52 = Companion_D3_.Create_DC8_(_564_v1)
				var _631_v53 _dafny.Map
				_ = _631_v53
				_631_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_564_v1).Cardinality())).Plus((_dafny.Zero).Minus(_598_v32.F5())), Companion_D5_.Create_DC16_(_598_v32.F5(), (_606_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_606_v37), 0))).Int()).(_dafny.MultiSet), _dafny.SeqOf(Companion_Default___.Fm12((_this).F3(), (_563_v0).F6(), (_629_v51).F6(), globalState), Companion_Default___.Fm12((_563_v0).F6(), p2, !((_this).F3()), globalState)), _630_v52))
				var _632_v54 D5
				_ = _632_v54
				_632_v54 = Companion_D5_.Create_DC16_(_dafny.IntOfUint32((_564_v1).Cardinality()), _dafny.MultiSetFromSeq(_611_v39), _628_cf6, _630_v52)
				_631_v53 = (_631_v53).Update((Companion_Default___.Fm15(globalState)).Cardinality(), _632_v54)
				var _633_v55 _dafny.CodePoint
				_ = _633_v55
				_633_v55 = _dafny.CodePoint('i')
				var _634_v56 *C0
				_ = _634_v56
				var _nw115 *C0 = New_C0_()
				_ = _nw115
				_nw115.Ctor__(_dafny.MultiSetOf(_633_v55))
				_634_v56 = _nw115
				var _635_v57 _dafny.Array
				_ = _635_v57
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw116
				_635_v57 = _nw116
				var _636_v58 _dafny.Sequence
				_ = _636_v58
				_636_v58 = _dafny.SeqOf(_635_v57)
				_636_v58 = _dafny.SeqOf(_635_v57, _635_v57, _635_v57, _635_v57, _635_v57)
			}
			(globalState).F2 = (_this).F3()
		}
		var _hi3 _dafny.Int = _dafny.IntOfUint32((_564_v1).Cardinality())
		_ = _hi3
		for _637_i3 := p1; _637_i3.Cmp(_hi3) < 0; _637_i3 = _637_i3.Plus(_dafny.One) {
			var _638_v59 _dafny.Array
			_ = _638_v59
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_14
			var _nw117 _dafny.Array
			_ = _nw117
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw117 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = func(_639_i4 _dafny.Int) bool {
					return !((_this).F3())
				}
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw117 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw117).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw117).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_638_v59 = _nw117
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_638_v59), 0))
			_ = _index104
			(_638_v59).ArraySet1((func() bool {
				if p2 {
					return (_this).F3()
				}
				return (_this).F3()
			})(), (_index104).Int())
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_638_v59), 0))
			_ = _index105
			(_638_v59).ArraySet1((p2) == ((_563_v0).F6()), (_index105).Int())
			var _640_v60 _dafny.CodePoint
			_ = _640_v60
			_640_v60 = _dafny.CodePoint('w')
			var _641_v61 *C0
			_ = _641_v61
			var _nw118 *C0 = New_C0_()
			_ = _nw118
			_nw118.Ctor__(_dafny.MultiSetOf(_dafny.CodePoint('y'), _640_v60))
			_641_v61 = _nw118
			var _642_v62 _dafny.MultiSet
			_ = _642_v62
			_642_v62 = _dafny.MultiSetOf(_641_v61, _641_v61)
			var _643_v63 *C1
			_ = _643_v63
			var _nw119 *C1 = New_C1_()
			_ = _nw119
			_nw119.Ctor__((_642_v62).Cardinality(), _dafny.IntOfInt64(831))
			_643_v63 = _nw119
			var _nw120 *C1 = New_C1_()
			_ = _nw120
			_nw120.Ctor__(Companion_Default___.Fm3((Companion_D1_.Create_DC4_(!(false))).Dtor_cf4(), globalState), _637_i3)
			_643_v63 = _nw120
			var _644_v64 _dafny.Int
			_ = _644_v64
			_644_v64 = _dafny.IntOfInt64(149)
			var _645_v65 _dafny.Set
			_ = _645_v65
			_645_v65 = _dafny.SetOf((_563_v0).F6())
			var _646_v66 _dafny.Set
			_ = _646_v66
			_646_v66 = _dafny.SetOf((_645_v65).Cardinality())
			var _647_v68 _dafny.Set
			_ = _647_v68
			_647_v68 = _dafny.SetOf(_646_v66, func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(595), _dafny.One)); ; {
					_compr_13, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _648_v67 _dafny.Int
					_648_v67 = interface{}(_compr_13).(_dafny.Int)
					if ((_dafny.IntOfInt64(595)).Cmp(_648_v67) <= 0) && ((_648_v67).Cmp(_dafny.One) < 0) {
						_coll13.Add(Companion_Default___.SafeDivisionInt(_648_v67, _dafny.IntOfInt64(-992)))
					}
				}
				return _coll13.ToSet()
			}())
			_644_v64 = ((_dafny.Zero).Minus((_647_v68).Cardinality())).Minus(_637_i3)
			var _649_v69 *C0
			_ = _649_v69
			var _nw121 *C0 = New_C0_()
			_ = _nw121
			_nw121.Ctor__((_641_v61).F7())
			_649_v69 = _nw121
		}
		var _650_v70 _dafny.Sequence
		_ = _650_v70
		_650_v70 = _dafny.SeqOf(true)
		var _651_v71 D1
		_ = _651_v71
		_651_v71 = Companion_D1_.Create_DC4_((_650_v70).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_650_v70).Cardinality()))).Uint32()).(bool))
		var _652_v72 D1
		_ = _652_v72
		_652_v72 = Companion_D1_.Create_DC5_(_651_v71)
		_652_v72 = _652_v72
		var _653_v73 _dafny.CodePoint
		_ = _653_v73
		_653_v73 = _dafny.CodePoint('y')
		var _654_v74 _dafny.Map
		_ = _654_v74
		_654_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _653_v73)
		_654_v74 = (_654_v74).Update(p2, _653_v73)
		var _hi4 _dafny.Int = p1
		_ = _hi4
		for _655_i5 := p1; _655_i5.Cmp(_hi4) < 0; _655_i5 = _655_i5.Plus(_dafny.One) {
			var _656_v75 _dafny.Map
			_ = _656_v75
			_656_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_655_i5, (_563_v0).F6())
			var _657_v76 _dafny.MultiSet
			_ = _657_v76
			_657_v76 = _dafny.MultiSetOf(_dafny.IntOfUint32((Companion_Default___.Fm16((_563_v0).F6(), Companion_Default___.Fm3((_563_v0).F6(), globalState), (_563_v0).F6(), globalState)).Cardinality()), _dafny.IntOfInt64(504), _655_i5)
			var _658_v77 *C1
			_ = _658_v77
			var _nw122 *C1 = New_C1_()
			_ = _nw122
			_nw122.Ctor__(_dafny.IntOfInt64(-609), ((_656_v75).Cardinality()).Minus((_657_v76).Cardinality()))
			_658_v77 = _nw122
			var _659_v78 _dafny.Map
			_ = _659_v78
			_659_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _655_i5)
			_659_v78 = (_659_v78).Update(true, (func() _dafny.Int {
				if (_563_v0).F6() {
					return p1
				}
				return p1
			})())
			var _660_v79 _dafny.Int
			_ = _660_v79
			_660_v79 = _dafny.IntOfInt64(35)
			_660_v79 = (_dafny.Zero).Minus(_660_v79)
			var _661_v80 D0
			_ = _661_v80
			_661_v80 = Companion_D0_.Create_DC2_()
			_661_v80 = _661_v80
		}
		var _662_v81 _dafny.Sequence
		_ = _662_v81
		_662_v81 = _dafny.SeqOf(_dafny.IntOfUint32((_564_v1).Cardinality()), (_dafny.Zero).Minus(p1))
		r0 = _dafny.Companion_Sequence_.Concatenate(_662_v81, _662_v81)
		var _663_v82 _dafny.Map
		_ = _663_v82
		_663_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _653_v73)
		var _664_v83 _dafny.Map
		_ = _664_v83
		_664_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v82, (_563_v0).F6())
		r1 = _664_v83
		return r0, r1
	}
}
func (_this *C3) F3() bool {
	{
		return _this._f3
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
