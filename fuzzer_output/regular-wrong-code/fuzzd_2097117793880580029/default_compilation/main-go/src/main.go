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
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D2 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqOf(true, true, false, true)).Cardinality()))
	_ = _source0
	if _source0.Is_DC5() {
		var _0___mcc_h0 bool = _source0.Get_().(D2_DC5).Cf6
		_ = _0___mcc_h0
		var _1___mcc_h1 bool = _source0.Get_().(D2_DC5).Cf7
		_ = _1___mcc_h1
		var _2___mcc_h2 _dafny.Set = _source0.Get_().(D2_DC5).Cf8
		_ = _2___mcc_h2
		var _3_cf8 _dafny.Set = _2___mcc_h2
		_ = _3_cf8
		var _4_cf7 bool = _1___mcc_h1
		_ = _4_cf7
		var _5_cf6 bool = _0___mcc_h0
		_ = _5_cf6
		return _dafny.SeqOf(false)
	} else if _source0.Is_DC6() {
		var _6___mcc_h3 _dafny.Int = _source0.Get_().(D2_DC6).Cf9
		_ = _6___mcc_h3
		var _7_cf9 _dafny.Int = _6___mcc_h3
		_ = _7_cf9
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true))
	} else if _source0.Is_DC7() {
		var _8___mcc_h4 _dafny.Int = _source0.Get_().(D2_DC7).Cf10
		_ = _8___mcc_h4
		var _9___mcc_h5 _dafny.Int = _source0.Get_().(D2_DC7).Cf11
		_ = _9___mcc_h5
		var _10_cf11 _dafny.Int = _9___mcc_h5
		_ = _10_cf11
		var _11_cf10 _dafny.Int = _8___mcc_h4
		_ = _11_cf10
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true))
	} else {
		var _12___mcc_h6 _dafny.MultiSet = _source0.Get_().(D2_DC4).Cf5
		_ = _12___mcc_h6
		var _13_cf5 _dafny.MultiSet = _12___mcc_h6
		_ = _13_cf5
		return _dafny.SeqOf(true, false, true)
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(-960), _dafny.IntOfInt64(274))).Cardinality(), _dafny.IntOfInt64(714)))).Difference(_dafny.SetOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rbbbp")).Cardinality()))))).Union((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(311), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yckvpg")).Cardinality()), (_dafny.SetOf(Companion_D6_.Create_DC19_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eshn")).Cardinality()))).Cardinality()), _dafny.IntOfInt64(869))).Cardinality(), _dafny.IntOfInt64(101), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-108))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality())), Companion_D6_.Create_DC19_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(950))).Cardinality(), _dafny.IntOfInt64(121), (_dafny.SetOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(385), _dafny.IntOfInt64(991))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _14_v0 _dafny.Int
			_14_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(385)).Cmp(_14_v0) <= 0) && ((_14_v0).Cmp(_dafny.IntOfInt64(991)) < 0) {
				_coll0.Add((_14_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vlcha")).Cardinality())), _dafny.IntOfInt64(726))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(365), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hnjtc")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)).Cardinality())).Cardinality())).Cardinality())).Cardinality())).Cardinality()), Companion_D6_.Create_DC19_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), (func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(476), _dafny.IntOfInt64(250))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _15_v1 _dafny.Int
			_15_v1 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(476)).Cmp(_15_v1) <= 0) && ((_15_v1).Cmp(_dafny.IntOfInt64(250)) < 0) {
				_coll1.Add((_15_v1).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ge")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))
			}
		}
		return _coll1.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))).Cardinality()), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-390), _dafny.IntOfInt64(638))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _16_v2 _dafny.Int
			_16_v2 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(-390)).Cmp(_16_v2) <= 0) && ((_16_v2).Cmp(_dafny.IntOfInt64(638)) < 0) {
				_coll2.Add((_16_v2).Times(_dafny.IntOfInt64(581)), _dafny.CodePoint('u'))
			}
		}
		return _coll2.ToMap()
	}()).Cardinality(), (_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(903), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mi")).Cardinality()), _dafny.IntOfInt64(-912))).Cardinality(), _dafny.IntOfInt64(272))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality())).Cardinality()), _dafny.IntOfInt64(948)))).Cardinality(), _dafny.IntOfInt64(863)))).Intersection(func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-420)))))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _17_v3 _dafny.Sequence
			_17_v3 = interface{}(_compr_3).(_dafny.Sequence)
			if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-420)))))).Contains(_17_v3) {
				_coll3.Add(_17_v3)
			}
		}
		return _coll3.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rf")).Cardinality()))).Cardinality())).Cardinality()), true)
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(951), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-444))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_18_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
		}))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_4).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-444))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_18_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
			})), _19_v0) {
				_coll4.Add((_19_v0).Plus((_dafny.SetOf(true)).Cardinality()), _dafny.IntOfInt64(390))
			}
		}
		return _coll4.ToMap()
	}()).Cardinality())).Cardinality()))).Times(((Companion_D5_.Create_DC16_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bmlkxdin")).Cardinality()), true, true)).Dtor_cf22()).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), _dafny.CodePoint('f'))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) D1 {
	var _source1 D5 = Companion_D5_.Create_DC15_(true, false)
	_ = _source1
	if _source1.Is_DC15() {
		var _20___mcc_h0 bool = _source1.Get_().(D5_DC15).Cf20
		_ = _20___mcc_h0
		var _21___mcc_h1 bool = _source1.Get_().(D5_DC15).Cf21
		_ = _21___mcc_h1
		var _22_cf21 bool = _21___mcc_h1
		_ = _22_cf21
		var _23_cf20 bool = _20___mcc_h0
		_ = _23_cf20
		return Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cpt")).Cardinality()), _23_cf20)
	} else if _source1.Is_DC16() {
		var _24___mcc_h2 _dafny.Int = _source1.Get_().(D5_DC16).Cf22
		_ = _24___mcc_h2
		var _25___mcc_h3 bool = _source1.Get_().(D5_DC16).Cf23
		_ = _25___mcc_h3
		var _26___mcc_h4 bool = _source1.Get_().(D5_DC16).Cf24
		_ = _26___mcc_h4
		var _27_cf24 bool = _26___mcc_h4
		_ = _27_cf24
		var _28_cf23 bool = _25___mcc_h3
		_ = _28_cf23
		var _29_cf22 _dafny.Int = _24___mcc_h2
		_ = _29_cf22
		return Companion_D1_.Create_DC2_((_dafny.SetOf(_27_cf24, _28_cf23)).Cardinality(), _27_cf24)
	} else {
		var _30___mcc_h5 _dafny.Sequence = _source1.Get_().(D5_DC14).Cf19
		_ = _30___mcc_h5
		var _31_cf19 _dafny.Sequence = _30___mcc_h5
		_ = _31_cf19
		return Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-957), true)
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Map, p3 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ggyklnkw")).Cardinality()), _dafny.IntOfInt64(133))
}
func (_static *CompanionStruct_Default___) Fm10(p0 D2, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false)).Difference((_dafny.SetOf(!(!(!(false))))).Intersection(_dafny.SetOf(false, true, !(!(false)), true, true)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) bool {
	if !(!((func() bool {
		if false {
			return true
		}
		return false
	})())) {
		return !(true)
	} else {
		return false
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("suuupmma"), _dafny.UnicodeSeqOfUtf8Bytes("ogyyc"))).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("bgrjnfjl")))).Intersection(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("hjljvpj"), _dafny.UnicodeSeqOfUtf8Bytes("imkj"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(208))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_32_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _33_v0 _dafny.Sequence
			_33_v0 = interface{}(_compr_5).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("hjljvpj"), _dafny.UnicodeSeqOfUtf8Bytes("imkj"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(208))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_32_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))).Contains(_33_v0) {
				_coll5.Add(_33_v0)
			}
		}
		return _coll5.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.IntOfInt64(762)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gdm")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-952))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_34_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D5_.Create_DC16_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D2_.Create_DC5_(true, false, _dafny.SetOf(_dafny.IntOfInt64(-109), _dafny.IntOfInt64(518)))).Dtor_cf6()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-434)))).Cardinality()))).Cardinality())).Keys().Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _35_v0 _dafny.CodePoint
				_35_v0 = interface{}(_compr_7).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D2_.Create_DC5_(true, false, _dafny.SetOf(_dafny.IntOfInt64(-109), _dafny.IntOfInt64(518)))).Dtor_cf6()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-434)))).Cardinality()))).Cardinality())).Contains(_35_v0) {
					_coll7.Add(_35_v0)
				}
			}
			return _coll7.ToSet()
		}()).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _36_v1 _dafny.CodePoint
			_36_v1 = interface{}(_compr_6).(_dafny.CodePoint)
			if (func() _dafny.Set {
				var _coll8 = _dafny.NewBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D2_.Create_DC5_(true, false, _dafny.SetOf(_dafny.IntOfInt64(-109), _dafny.IntOfInt64(518)))).Dtor_cf6()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-434)))).Cardinality()))).Cardinality())).Keys().Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _37_v0 _dafny.CodePoint
					_37_v0 = interface{}(_compr_8).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D2_.Create_DC5_(true, false, _dafny.SetOf(_dafny.IntOfInt64(-109), _dafny.IntOfInt64(518)))).Dtor_cf6()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-434)))).Cardinality()))).Cardinality())).Contains(_37_v0) {
						_coll8.Add(_37_v0)
					}
				}
				return _coll8.ToSet()
			}()).Contains(_36_v1) {
				_coll6.Add(_36_v1)
			}
		}
		return _coll6.ToSet()
	}()).Cardinality())).Cardinality(), true, false)).Dtor_cf22(), true)).Cardinality()).Cmp(_dafny.IntOfInt64(-981)) > 0, (_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(35), _dafny.IntOfInt64(349))))).IsProperSubsetOf(_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(837), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gcdsylhn")).Cardinality())))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.MultiSet, globalState *GlobalState) (_dafny.Set, bool, _dafny.Map) {
	var r0 _dafny.Set = _dafny.EmptySet
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.Map = _dafny.EmptyMap
	_ = r2
	var _38_v0 _dafny.Int
	_ = _38_v0
	_38_v0 = _dafny.IntOfInt64(353)
	var _39_v1 _dafny.CodePoint
	_ = _39_v1
	_39_v1 = _dafny.CodePoint('d')
	var _40_v2 _dafny.Map
	_ = _40_v2
	_40_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(_dafny.IntOfInt64(136), _38_v0, globalState), _39_v1)
	_40_v2 = (_40_v2).Update(_38_v0, _39_v1)
	var _41_v3 D2
	_ = _41_v3
	_41_v3 = Companion_D2_.Create_DC6_(_38_v0)
	var _42_v4 *C0
	_ = _42_v4
	var _nw0 *C0 = New_C0_()
	_ = _nw0
	_nw0.Ctor__((_41_v3).Dtor_cf9())
	_42_v4 = _nw0
	var _43_v5 _dafny.Map
	_ = _43_v5
	_43_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm11(_dafny.IntOfInt64(641), globalState), _42_v4)
	var _44_v6 _dafny.Array
	_ = _44_v6
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
	_ = _nw1
	_44_v6 = _nw1
	var _45_v8 _dafny.Array
	_ = _45_v8
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Map = (func(_46_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
			return func(_47_i0 _dafny.Int) _dafny.Map {
				return func() _dafny.Map {
					var _coll9 = _dafny.NewMapBuilder()
					_ = _coll9
					for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(826), _dafny.IntOfInt64(649))); ; {
						_compr_9, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _48_v7 _dafny.Int
						_48_v7 = interface{}(_compr_9).(_dafny.Int)
						if ((_dafny.IntOfInt64(826)).Cmp(_48_v7) <= 0) && ((_48_v7).Cmp(_dafny.IntOfInt64(649)) < 0) {
							_coll9.Add(Companion_Default___.SafeModuloInt(_48_v7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _46_v1)).Cardinality()), false)
						}
					}
					return _coll9.ToMap()
				}()
			}
		})(_39_v1)
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
	_45_v8 = _nw2
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_44_v6), 0))
	_ = _index0
	(_44_v6).ArraySet1(_45_v8, (_index0).Int())
	var _49_v9 bool
	_ = _49_v9
	_49_v9 = true
	var _50_v10 _dafny.Sequence
	_ = _50_v10
	_50_v10 = _dafny.UnicodeSeqOfUtf8Bytes("kngri")
	var _51_v11 D4
	_ = _51_v11
	_51_v11 = Companion_D4_.Create_DC10_(_45_v8)
	var _pat_let_tv0 = _45_v8
	_ = _pat_let_tv0
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_44_v6), 0))
	_ = _index1
	var _rhs0 _dafny.Int = ((_dafny.SetOf(!(_49_v9), _49_v9, _49_v9)).Cardinality()).Minus((_38_v0).Minus(Companion_Default___.Fm7((_dafny.Zero).Minus((_42_v4).F22()), (_dafny.Zero).Minus(_38_v0), globalState)))
	_ = _rhs0
	var _rhs1 _dafny.Map = _43_v5
	_ = _rhs1
	var _rhs2 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.Fm7((_42_v4).F22(), _38_v0, globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_50_v10, _50_v10)).Cardinality()))
	_ = _rhs2
	var _rhs3 _dafny.Array = (func(_pat_let0_0 D4) D4 {
		return func(_52_dt__update__tmp_h0 D4) D4 {
			return func(_pat_let1_0 _dafny.Array) D4 {
				return func(_53_dt__update_hcf14_h0 _dafny.Array) D4 {
					return Companion_D4_.Create_DC10_(_53_dt__update_hcf14_h0)
				}(_pat_let1_0)
			}(_pat_let_tv0)
		}(_pat_let0_0)
	}(_51_v11)).Dtor_cf14()
	_ = _rhs3
	var _rhs4 bool = _49_v9
	_ = _rhs4
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 _dafny.Array = _44_v6
	_ = _lhs1
	var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_44_v6), 0))
	_ = _lhs2
	var _lhs3 *GlobalState = globalState
	_ = _lhs3
	_38_v0 = _rhs0
	_43_v5 = _rhs1
	_lhs0.F3 = _rhs2
	(_lhs1).ArraySet1(_rhs3, (_lhs2).Int())
	_lhs3.F14 = _rhs4
	_50_v10 = _50_v10
	var _54_v12 _dafny.Map
	_ = _54_v12
	_54_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v0, _49_v9)
	var _55_v13 _dafny.Map
	_ = _55_v13
	_55_v13 = _54_v12
	var _56_v14 _dafny.Map
	_ = _56_v14
	_56_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, (_55_v13).Cardinality())
	var _57_v15 _dafny.Set
	_ = _57_v15
	_57_v15 = _dafny.SetOf((Companion_Default___.Fm12((_42_v4).F22(), _38_v0, _50_v10, globalState)).Cardinality(), (_42_v4).F22(), (func() _dafny.Int {
		if (_56_v14).Contains(_49_v9) {
			return (_56_v14).Get(_49_v9).(_dafny.Int)
		}
		return (_42_v4).F22()
	})(), _38_v0)
	var _58_i1 _dafny.Int
	_ = _58_i1
	_58_i1 = _dafny.Zero
	{
		for !((_57_v15).IsProperSubsetOf(_dafny.SetOf(_38_v0))) {
			{
				if (_58_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_58_i1 = (_58_i1).Plus(_dafny.One)
				var _59_v16 _dafny.Sequence
				_ = _59_v16
				_59_v16 = _dafny.SeqOf(_43_v5, _43_v5)
				(globalState).F14 = ((((_59_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.IntOfUint32((_59_v16).Cardinality()))).Uint32()).(_dafny.Map)).Merge((_43_v5).Update(false, _42_v4))).Cardinality()).Cmp((_42_v4).F22()) >= 0
				var _60_v17 _dafny.Sequence
				_ = _60_v17
				_60_v17 = _dafny.SeqOf(_49_v9, _49_v9)
				var _61_v18 *C1
				_ = _61_v18
				var _nw3 *C1 = New_C1_()
				_ = _nw3
				_nw3.Ctor__(_dafny.IntOfInt64(929), _60_v17)
				_61_v18 = _nw3
				var _62_v19 _dafny.Set
				_ = _62_v19
				_62_v19 = _dafny.SetOf(_61_v18, _61_v18)
				_62_v19 = _62_v19
				var _63_v20 _dafny.Map
				_ = _63_v20
				_63_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm13(_39_v1, _49_v9, globalState), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(552), (_42_v4).F22()))
				var _64_v21 _dafny.Sequence
				_ = _64_v21
				_64_v21 = _dafny.SeqOf((_42_v4).F22())
				_63_v20 = (_63_v20).Update(_64_v21, (_42_v4).F22())
				var _65_v22 _dafny.Array
				_ = _65_v22
				var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw4
				_65_v22 = _nw4
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_65_v22), 0))
				_ = _index2
				(_65_v22).ArraySet1(_50_v10, (_index2).Int())
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_65_v22), 0))
				_ = _index3
				(_65_v22).ArraySet1(_50_v10, (_index3).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	if !(_49_v9) {
		if _49_v9 {
			var _66_v23 _dafny.Sequence
			_ = _66_v23
			_66_v23 = _dafny.SeqOf(_49_v9)
			var _67_v24 *C1
			_ = _67_v24
			var _nw5 *C1 = New_C1_()
			_ = _nw5
			_nw5.Ctor__((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_68_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality())).Plus((_42_v4).F22()), _66_v23)
			_67_v24 = _nw5
			var _rhs5 _dafny.Int = (func() _dafny.Int {
				if _49_v9 {
					return (_42_v4).F22()
				}
				return (_42_v4).F22()
			})()
			_ = _rhs5
			var _rhs6 _dafny.Int = (_42_v4).F22()
			_ = _rhs6
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			_lhs4.F3 = _rhs5
			_38_v0 = _rhs6
			var _69_v25 D3
			_ = _69_v25
			_69_v25 = Companion_D3_.Create_DC8_((Companion_D3_.Create_DC8_(_67_v24)).Dtor_cf12())
			var _70_v26 _dafny.MultiSet
			_ = _70_v26
			_70_v26 = _dafny.MultiSetOf(_69_v25)
			var _71_v27 _dafny.Array
			_ = _71_v27
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw6
			_71_v27 = _nw6
			var _72_v28 _dafny.Array
			_ = _72_v28
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_1
			var _nw7 _dafny.Array
			_ = _nw7
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw7 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_73_v9 bool) func(_dafny.Int) bool {
					return func(_74_i3 _dafny.Int) bool {
						return _73_v9
					}
				})(_49_v9)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw7 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw7).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw7).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_72_v28 = _nw7
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_71_v27), 0))
			_ = _index4
			(_71_v27).ArraySet1((func() _dafny.Array {
				if _49_v9 {
					return _72_v28
				}
				return _72_v28
			})(), (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_71_v27), 0))
			_ = _index5
			var _rhs7 _dafny.MultiSet = _dafny.MultiSetOf(_69_v25)
			_ = _rhs7
			var _rhs8 _dafny.Array = _72_v28
			_ = _rhs8
			var _rhs9 bool = !(_49_v9) || (!(_49_v9))
			_ = _rhs9
			var _rhs10 bool = !_dafny.Companion_Sequence_.Contains(_50_v10, _39_v1)
			_ = _rhs10
			var _lhs5 _dafny.Array = _71_v27
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_71_v27), 0))
			_ = _lhs6
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			_70_v26 = _rhs7
			(_lhs5).ArraySet1(_rhs8, (_lhs6).Int())
			_lhs7.F14 = _rhs9
			_lhs8.F14 = _rhs10
			(globalState).F7 = _49_v9
			var _75_v30 D2
			_ = _75_v30
			_75_v30 = Companion_D2_.Create_DC7_((_dafny.Zero).Minus((_42_v4).F22()), (func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(667), _dafny.IntOfInt64(551))); ; {
					_compr_10, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _76_v29 _dafny.Int
					_76_v29 = interface{}(_compr_10).(_dafny.Int)
					if ((_dafny.IntOfInt64(667)).Cmp(_76_v29) <= 0) && ((_76_v29).Cmp(_dafny.IntOfInt64(551)) < 0) {
						_coll10.Add(Companion_Default___.SafeDivisionInt(_76_v29, (_42_v4).F22()), _49_v9)
					}
				}
				return _coll10.ToMap()
			}()).Cardinality())
			var _77_v31 _dafny.Map
			_ = _77_v31
			_77_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, _39_v1)).Update(_49_v9, _39_v1), _66_v23)
			var _78_v32 _dafny.Set
			_ = _78_v32
			_78_v32 = _dafny.SetOf(_49_v9)
			var _79_v33 _dafny.Map
			_ = _79_v33
			_79_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, _78_v32)
			var _80_v34 _dafny.Array
			_ = _80_v34
			var _nwElement0_0 _dafny.Int = (_75_v30).Dtor_cf10()
			_ = _nwElement0_0
			var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(5))
			_ = _nw8
			(_nw8).ArraySet1(_nwElement0_0, 0)
			(_nw8).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm14(globalState), _50_v10)).Cardinality()), 1)
			(_nw8).ArraySet1(((_42_v4).F22()).Times((_42_v4).F22()), 2)
			(_nw8).ArraySet1(((_67_v24).Fm1(_77_v31, globalState)).Times(_dafny.IntOfInt64(351)), 3)
			(_nw8).ArraySet1((_79_v33).Cardinality(), 4)
			_80_v34 = _nw8
			var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw9
			_80_v34 = _nw9
		} else {
			var _81_v35 _dafny.Sequence
			_ = _81_v35
			_81_v35 = _dafny.SeqOf(_49_v9)
			var _82_v36 *C1
			_ = _82_v36
			var _nw10 *C1 = New_C1_()
			_ = _nw10
			_nw10.Ctor__((_42_v4).F22(), _81_v35)
			_82_v36 = _nw10
			var _83_v37 _dafny.Sequence
			_ = _83_v37
			_83_v37 = _dafny.SeqOf(_82_v36)
			var _84_v38 _dafny.Sequence
			_ = _84_v38
			_84_v38 = _dafny.SeqOf((_82_v36).F20(), (_42_v4).F22(), (_42_v4).F22(), (_42_v4).F22())
			var _85_v39 _dafny.Sequence
			_ = _85_v39
			_85_v39 = _dafny.SeqOf(_84_v38, _84_v38)
			var _86_v40 _dafny.Sequence
			_ = _86_v40
			_86_v40 = _dafny.SeqOf(_50_v10)
			var _87_v41 _dafny.Array
			_ = _87_v41
			var _nwElement0_1 _dafny.Int = _dafny.IntOfUint32((_50_v10).Cardinality())
			_ = _nwElement0_1
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(15))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_1, 0)
			(_nw11).ArraySet1(_38_v0, 1)
			(_nw11).ArraySet1(_dafny.IntOfInt64(-302), 2)
			(_nw11).ArraySet1(_dafny.IntOfUint32((_83_v37).Cardinality()), 3)
			(_nw11).ArraySet1((_82_v36).F20(), 4)
			(_nw11).ArraySet1((_42_v4).F22(), 5)
			(_nw11).ArraySet1((_42_v4).F22(), 6)
			(_nw11).ArraySet1((_42_v4).F22(), 7)
			(_nw11).ArraySet1(((_42_v4).F22()).Minus(_dafny.IntOfInt64(531)), 8)
			(_nw11).ArraySet1(_38_v0, 9)
			(_nw11).ArraySet1(_38_v0, 10)
			(_nw11).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_81_v35).Cardinality())), 11)
			(_nw11).ArraySet1((_42_v4).F22(), 12)
			(_nw11).ArraySet1((_dafny.IntOfUint32((_85_v39).Cardinality())).Minus(_dafny.IntOfUint32((_86_v40).Cardinality())), 13)
			(_nw11).ArraySet1((_dafny.Zero).Minus(((_dafny.Zero).Minus((_42_v4).F22())).Minus(_dafny.IntOfInt64(366))), 14)
			_87_v41 = _nw11
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_87_v41), 0))
			_ = _index6
			(_87_v41).ArraySet1(_dafny.IntOfUint32((_50_v10).Cardinality()), (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_87_v41), 0))
			_ = _index7
			(_87_v41).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm7(Companion_Default___.SafeDivisionInt((_82_v36).F20(), (_42_v4).F22()), (_42_v4).F22(), globalState)), (_index7).Int())
			var _88_v42 _dafny.Array
			_ = _88_v42
			var _nw12 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw12
			_88_v42 = _nw12
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_88_v42), 0))
			_ = _index8
			(_88_v42).ArraySet1(_49_v9, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_88_v42), 0))
			_ = _index9
			(_88_v42).ArraySet1((_82_v36).Fm0(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lsnskptul"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(893))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_89_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_90_i4 _dafny.Int) _dafny.CodePoint {
					return _89_v1
				}
			})(_39_v1)))), (_42_v4).F22(), (_42_v4).F22(), _dafny.IntOfInt64(994), globalState), (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_88_v42), 0))
			_ = _index10
			(_88_v42).ArraySet1(!(_49_v9), (_index10).Int())
			var _91_v43 _dafny.Map
			_ = _91_v43
			_91_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, (_88_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_88_v42), 0))).Int()).(bool))
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_87_v41), 0))
			_ = _index11
			(_87_v41).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_41_v3).Dtor_cf9(), (_82_v36).F20()), Companion_Default___.SafeDivisionInt((_82_v36).F20(), (_91_v43).Cardinality())), (_index11).Int())
			(globalState).F11 = (_81_v35).Select((Companion_Default___.SafeIndex(_38_v0, _dafny.IntOfUint32((_81_v35).Cardinality()))).Uint32()).(bool)
		}
		var _92_v44 _dafny.Array
		_ = _92_v44
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw13
		_92_v44 = _nw13
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_92_v44), 0))
		_ = _index12
		(_92_v44).ArraySet1(_38_v0, (_index12).Int())
		var _93_v45 _dafny.MultiSet
		_ = _93_v45
		_93_v45 = _dafny.MultiSetOf(_49_v9)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_92_v44), 0))
		_ = _index13
		(_92_v44).ArraySet1((func() _dafny.Int {
			if _49_v9 {
				return (_42_v4).F22()
			}
			return Companion_Default___.SafeDivisionInt((_93_v45).Cardinality(), _38_v0)
		})(), (_index13).Int())
		(globalState).F14 = _49_v9
		(globalState).F3 = (_92_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_92_v44), 0))).Int()).(_dafny.Int)
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_92_v44), 0))
		_ = _index14
		(_92_v44).ArraySet1(((_42_v4).F22()).Plus((_92_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_92_v44), 0))).Int()).(_dafny.Int)), (_index14).Int())
	} else {
		var _94_v46 D1
		_ = _94_v46
		_94_v46 = Companion_D1_.Create_DC2_((_42_v4).F22(), _49_v9)
		(globalState).F5 = (func() bool {
			if _49_v9 {
				return (_94_v46).Dtor_cf3()
			}
			return _49_v9
		})()
		(globalState).F14 = Companion_Default___.Fm11((_42_v4).F22(), globalState)
		var _95_v47 _dafny.Sequence
		_ = _95_v47
		_95_v47 = _dafny.SeqOf(true, _49_v9)
		var _96_v48 *C1
		_ = _96_v48
		var _nw14 *C1 = New_C1_()
		_ = _nw14
		_nw14.Ctor__(_38_v0, _95_v47)
		_96_v48 = _nw14
		var _97_v49 _dafny.Map
		_ = _97_v49
		_97_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v48, _dafny.SeqOf((_96_v48).F20(), (_42_v4).F22()))
		var _98_v50 _dafny.Sequence
		_ = _98_v50
		_98_v50 = _dafny.SeqOf(_dafny.IntOfInt64(-91), (_42_v4).F22())
		if (func() bool {
			if !_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
				if (_97_v49).Contains(_96_v48) {
					return (_97_v49).Get(_96_v48).(_dafny.Sequence)
				}
				return _98_v50
			})(), _38_v0) {
				return ((_96_v48).F20()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ihliafrys")).Cardinality())) > 0
			}
			return _49_v9
		})() {
			var _99_v51 _dafny.Map
			_ = _99_v51
			_99_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_42_v4, (_96_v48).F20())
			_38_v0 = (_99_v51).Cardinality()
			var _100_v52 _dafny.Sequence
			_ = _100_v52
			_100_v52 = _dafny.SeqOf(_50_v10, _50_v10)
			var _101_v53 _dafny.MultiSet
			_ = _101_v53
			_101_v53 = _dafny.MultiSetOf(_39_v1)
			var _102_v54 D6
			_ = _102_v54
			_102_v54 = Companion_D6_.Create_DC17_(_39_v1)
			var _rhs11 _dafny.Int = (_dafny.Zero).Minus(((_42_v4).F22()).Plus(_dafny.IntOfInt64(830)))
			_ = _rhs11
			var _rhs12 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((Companion_D5_.Create_DC14_((_100_v52).Select((Companion_Default___.SafeIndex((_101_v53).Cardinality(), _dafny.IntOfUint32((_100_v52).Cardinality()))).Uint32()).(_dafny.Sequence))).Dtor_cf19(), _dafny.Companion_Sequence_.Update(_50_v10, (Companion_Default___.SafeIndex(_38_v0, _dafny.IntOfUint32((_50_v10).Cardinality()))).Uint32(), (_102_v54).Dtor_cf25()))
			_ = _rhs12
			var _rhs13 _dafny.Int = (_42_v4).F22()
			_ = _rhs13
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			var _lhs10 *GlobalState = globalState
			_ = _lhs10
			_lhs9.F3 = _rhs11
			_50_v10 = _rhs12
			_lhs10.F3 = _rhs13
			var _103_v55 _dafny.MultiSet
			_ = _103_v55
			_103_v55 = _dafny.MultiSetOf(_42_v4, _42_v4)
			var _104_v56 _dafny.Map
			_ = _104_v56
			_104_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, true)
			var _105_v58 _dafny.Sequence
			_ = _105_v58
			_105_v58 = _dafny.SeqOf(_95_v47)
			var _106_v59 _dafny.Map
			_ = _106_v59
			_106_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v56, func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_105_v58).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _107_v57 _dafny.Sequence
					_107_v57 = interface{}(_compr_11).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_105_v58, _107_v57) {
						_coll11.Add(_107_v57, (_96_v48).F20())
					}
				}
				return _coll11.ToMap()
			}())
			var _108_v60 _dafny.Map
			_ = _108_v60
			_108_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_96_v48).F21(), _dafny.IntOfInt64(-797))
			var _109_v61 _dafny.MultiSet
			_ = _109_v61
			_109_v61 = _dafny.MultiSetOf((func() _dafny.Map {
				if (_106_v59).Contains(Companion_Default___.Fm15((_42_v4).F22(), _50_v10, false, (_42_v4).F22(), globalState)) {
					return (_106_v59).Get(Companion_Default___.Fm15((_42_v4).F22(), _50_v10, false, (_42_v4).F22(), globalState)).(_dafny.Map)
				}
				return _108_v60
			})(), _108_v60)
			var _110_v62 _dafny.Array
			_ = _110_v62
			var _nwElement0_2 _dafny.Int = (_96_v48).F20()
			_ = _nwElement0_2
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(26))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_2, 0)
			(_nw15).ArraySet1(_dafny.IntOfUint32((_50_v10).Cardinality()), 1)
			(_nw15).ArraySet1((_103_v55).Cardinality(), 2)
			(_nw15).ArraySet1(_dafny.IntOfUint32((_50_v10).Cardinality()), 3)
			(_nw15).ArraySet1((_dafny.Zero).Minus((_96_v48).F20()), 4)
			(_nw15).ArraySet1(_dafny.IntOfInt64(-971), 5)
			(_nw15).ArraySet1((_dafny.Zero).Minus(_38_v0), 6)
			(_nw15).ArraySet1((_42_v4).F22(), 7)
			(_nw15).ArraySet1((_96_v48).F20(), 8)
			(_nw15).ArraySet1((_42_v4).F22(), 9)
			(_nw15).ArraySet1(_dafny.IntOfInt64(570), 10)
			(_nw15).ArraySet1(_dafny.IntOfInt64(-802), 11)
			(_nw15).ArraySet1((_96_v48).F20(), 12)
			(_nw15).ArraySet1((_42_v4).F22(), 13)
			(_nw15).ArraySet1((_96_v48).F20(), 14)
			(_nw15).ArraySet1((_42_v4).F22(), 15)
			(_nw15).ArraySet1((_42_v4).F22(), 16)
			(_nw15).ArraySet1((func() _dafny.Int {
				if (_109_v61).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v47, _dafny.IntOfInt64(980))) {
					return (_109_v61).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v47, _dafny.IntOfInt64(980)))
				}
				return (_42_v4).F22()
			})(), 17)
			(_nw15).ArraySet1((_dafny.Zero).Minus((_42_v4).F22()), 18)
			(_nw15).ArraySet1((_42_v4).F22(), 19)
			(_nw15).ArraySet1((_96_v48).F20(), 20)
			(_nw15).ArraySet1((_96_v48).F20(), 21)
			(_nw15).ArraySet1((_56_v14).Cardinality(), 22)
			(_nw15).ArraySet1((_42_v4).F22(), 23)
			(_nw15).ArraySet1((_42_v4).F22(), 24)
			(_nw15).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_56_v14).Contains(_49_v9) {
					return (_56_v14).Get(_49_v9).(_dafny.Int)
				}
				return (_96_v48).F20()
			})()), 25)
			_110_v62 = _nw15
			var _111_v63 _dafny.Map
			_ = _111_v63
			_111_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, _110_v62)
			_111_v63 = (_111_v63).Update(_49_v9, _110_v62)
			var _112_v64 _dafny.MultiSet
			_ = _112_v64
			_112_v64 = _dafny.MultiSetOf((_96_v48).F20())
			_112_v64 = _dafny.MultiSetOf(_38_v0, (_42_v4).F22(), (_96_v48).F20(), (_96_v48).F20(), _38_v0)
			var _113_v65 _dafny.Sequence
			_ = _113_v65
			_113_v65 = _dafny.SeqOf(_42_v4, _42_v4, _42_v4, _42_v4)
			var _114_v66 *C0
			_ = _114_v66
			_114_v66 = _42_v4
			var _115_v67 _dafny.Array
			_ = _115_v67
			var _nwElement0_3 *C0 = (_113_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_50_v10).Cardinality()), _dafny.IntOfUint32((_113_v65).Cardinality()))).Uint32()).(*C0)
			_ = _nwElement0_3
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(23))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_3, 0)
			(_nw16).ArraySet1((_114_v66), 1)
			(_nw16).ArraySet1(_42_v4, 2)
			(_nw16).ArraySet1(_42_v4, 3)
			(_nw16).ArraySet1(_42_v4, 4)
			(_nw16).ArraySet1(_42_v4, 5)
			(_nw16).ArraySet1(_42_v4, 6)
			(_nw16).ArraySet1(_42_v4, 7)
			(_nw16).ArraySet1((func() *C0 {
				if _49_v9 {
					return _42_v4
				}
				return _42_v4
			})(), 8)
			(_nw16).ArraySet1(_42_v4, 9)
			(_nw16).ArraySet1(_42_v4, 10)
			(_nw16).ArraySet1(_42_v4, 11)
			(_nw16).ArraySet1(_42_v4, 12)
			(_nw16).ArraySet1(_42_v4, 13)
			(_nw16).ArraySet1(_42_v4, 14)
			(_nw16).ArraySet1(_42_v4, 15)
			(_nw16).ArraySet1((_113_v65).Select((Companion_Default___.SafeIndex((_42_v4).F22(), _dafny.IntOfUint32((_113_v65).Cardinality()))).Uint32()).(*C0), 16)
			(_nw16).ArraySet1(_42_v4, 17)
			(_nw16).ArraySet1(_42_v4, 18)
			(_nw16).ArraySet1(_42_v4, 19)
			(_nw16).ArraySet1(_42_v4, 20)
			(_nw16).ArraySet1(_42_v4, 21)
			(_nw16).ArraySet1(_42_v4, 22)
			_115_v67 = _nw16
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_115_v67), 0))
			_ = _index15
			(_115_v67).ArraySet1(_42_v4, (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_115_v67), 0))
			_ = _index16
			(_115_v67).ArraySet1(_42_v4, (_index16).Int())
		} else {
			var _116_v68 _dafny.Array
			_ = _116_v68
			var _nw17 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw17
			_116_v68 = _nw17
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_116_v68), 0))
			_ = _index17
			(_116_v68).ArraySet1(_49_v9, (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_116_v68), 0))
			_ = _index18
			(_116_v68).ArraySet1((_38_v0).Cmp((_dafny.Zero).Minus((_96_v48).F20())) < 0, (_index18).Int())
			var _117_v69 _dafny.Array
			_ = _117_v69
			var _nwElement0_4 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-659))
			_ = _nwElement0_4
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.One)
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_4, 0)
			_117_v69 = _nw18
			var _118_v70 _dafny.MultiSet
			_ = _118_v70
			_118_v70 = _dafny.MultiSetOf((_116_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_116_v68), 0))).Int()).(bool), (_116_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_116_v68), 0))).Int()).(bool))
			var _119_v71 _dafny.Map
			_ = _119_v71
			_119_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, (_42_v4).Fm4((_118_v70).Cardinality(), globalState))
			var _120_v72 _dafny.Map
			_ = _120_v72
			_120_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v71, (_96_v48).F21())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_117_v69), 0))
			_ = _index19
			(_117_v69).ArraySet1((_96_v48).Fm1(_120_v72, globalState), (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_117_v69), 0))
			_ = _index20
			(_117_v69).ArraySet1((_96_v48).F20(), (_index20).Int())
			(globalState).F3 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if _49_v9 {
					return _dafny.IntOfInt64(-49)
				}
				return _dafny.IntOfInt64(290)
			})(), (_42_v4).F22())
			var _121_v73 D5
			_ = _121_v73
			_121_v73 = Companion_D5_.Create_DC14_(_50_v10)
			var _122_v74 _dafny.Map
			_ = _122_v74
			_122_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_116_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_116_v68), 0))).Int()).(bool), _dafny.UnicodeSeqOfUtf8Bytes("btkawurs"))
			var _123_v75 _dafny.Array
			_ = _123_v75
			var _nwElement0_5 _dafny.Sequence = _50_v10
			_ = _nwElement0_5
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(19))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_5, 0)
			(_nw19).ArraySet1(_50_v10, 1)
			(_nw19).ArraySet1(_50_v10, 2)
			(_nw19).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_121_v73).Dtor_cf19(), _50_v10), 3)
			(_nw19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_50_v10, _50_v10), 4)
			(_nw19).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_50_v10, _50_v10), (Companion_Default___.SafeIndex((_98_v50).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.IntOfUint32((_98_v50).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_50_v10, _50_v10)).Cardinality()))).Uint32(), _39_v1), 5)
			(_nw19).ArraySet1(_50_v10, 6)
			(_nw19).ArraySet1(_50_v10, 7)
			(_nw19).ArraySet1(_50_v10, 8)
			(_nw19).ArraySet1(_50_v10, 9)
			(_nw19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kspjtwcah"), Companion_Default___.Fm14(globalState)), 10)
			(_nw19).ArraySet1(_50_v10, 11)
			(_nw19).ArraySet1(_dafny.Companion_Sequence_.Update(_50_v10, (Companion_Default___.SafeIndex((_42_v4).F22(), _dafny.IntOfUint32((_50_v10).Cardinality()))).Uint32(), _dafny.CodePoint('j')), 12)
			(_nw19).ArraySet1(_50_v10, 13)
			(_nw19).ArraySet1(_50_v10, 14)
			(_nw19).ArraySet1(_50_v10, 15)
			(_nw19).ArraySet1(_50_v10, 16)
			(_nw19).ArraySet1(_50_v10, 17)
			(_nw19).ArraySet1((func() _dafny.Sequence {
				if (_122_v74).Contains(_49_v9) {
					return (_122_v74).Get(_49_v9).(_dafny.Sequence)
				}
				return _50_v10
			})(), 18)
			_123_v75 = _nw19
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_123_v75), 0))
			_ = _index21
			(_123_v75).ArraySet1(_50_v10, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_123_v75), 0))
			_ = _index22
			(_123_v75).ArraySet1((func() _dafny.Sequence {
				if false {
					return _50_v10
				}
				return _50_v10
			})(), (_index22).Int())
			var _124_v76 _dafny.Map
			_ = _124_v76
			_124_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v69, _49_v9)
			var _125_v77 _dafny.Map
			_ = _125_v77
			_125_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_117_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_117_v69), 0))).Int()).(_dafny.Int), (_42_v4).F22()), _124_v76)
			_125_v77 = (_125_v77).Update((_96_v48).F20(), _124_v76)
		}
		var _126_v78 D8
		_ = _126_v78
		_126_v78 = Companion_D8_.Create_DC21_((_96_v48).F21())
		var _127_v79 _dafny.Map
		_ = _127_v79
		_127_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_49_v9), _49_v9)
		var _128_v80 _dafny.Set
		_ = _128_v80
		_128_v80 = _dafny.SetOf(_42_v4)
		var _129_v81 _dafny.Map
		_ = _129_v81
		_129_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_127_v79).Contains(Companion_Default___.Fm11((_128_v80).Cardinality(), globalState)) {
				return (_127_v79).Get(Companion_Default___.Fm11((_128_v80).Cardinality(), globalState)).(bool)
			}
			return !(_49_v9)
		})(), _57_v15)
		var _130_v82 _dafny.Map
		_ = _130_v82
		_130_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_126_v78).Dtor_cf34(), (func() _dafny.Set {
			if (_129_v81).Contains(_49_v9) {
				return (_129_v81).Get(_49_v9).(_dafny.Set)
			}
			return _dafny.SetOf(_dafny.IntOfUint32((_95_v47).Cardinality()))
		})())
		_130_v82 = (_130_v82).Update(_dafny.Companion_Sequence_.Update(_95_v47, (Companion_Default___.SafeIndex((_42_v4).F22(), _dafny.IntOfUint32((_95_v47).Cardinality()))).Uint32(), (func() bool {
			if (_127_v79).Contains(_49_v9) {
				return (_127_v79).Get(_49_v9).(bool)
			}
			return _49_v9
		})()), _57_v15)
		var _131_v83 _dafny.Array
		_ = _131_v83
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_2
		var _nw20 _dafny.Array
		_ = _nw20
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw20 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_132_v15 _dafny.Set) func(_dafny.Int) _dafny.Int {
				return func(_133_i5 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_133_i5, (_132_v15).Cardinality())
				}
			})(_57_v15)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw20 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw20).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw20).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_131_v83 = _nw20
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_131_v83), 0))
		_ = _index23
		(_131_v83).ArraySet1((_42_v4).F22(), (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_131_v83), 0))
		_ = _index24
		var _rhs14 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_42_v4).F22(), (_42_v4).F22()), _dafny.IntOfInt64(5))
		_ = _rhs14
		var _rhs15 _dafny.Int = (_96_v48).F20()
		_ = _rhs15
		var _lhs11 *GlobalState = globalState
		_ = _lhs11
		var _lhs12 _dafny.Array = _131_v83
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_131_v83), 0))
		_ = _lhs13
		_lhs11.F3 = _rhs14
		(_lhs12).ArraySet1(_rhs15, (_lhs13).Int())
	}
	var _134_v84 _dafny.MultiSet
	_ = _134_v84
	_134_v84 = _dafny.MultiSetOf(_50_v10, _50_v10)
	var _hi0 _dafny.Int = (_42_v4).F22()
	_ = _hi0
	for _135_i6 := (_134_v84).Cardinality(); _135_i6.Cmp(_hi0) < 0; _135_i6 = _135_i6.Plus(_dafny.One) {
		var _136_v85 _dafny.Array
		_ = _136_v85
		var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw21
		_136_v85 = _nw21
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_136_v85), 0))
		_ = _index25
		(_136_v85).ArraySet1((_42_v4).F22(), (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_136_v85), 0))
		_ = _index26
		(_136_v85).ArraySet1(Companion_Default___.SafeModuloInt(_135_i6, (_42_v4).F22()), (_index26).Int())
		var _137_v86 _dafny.Sequence
		_ = _137_v86
		_137_v86 = _dafny.SeqOf(_135_i6)
		var _138_v87 _dafny.Map
		_ = _138_v87
		_138_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v9, _137_v86)
		var _139_v88 *C0
		_ = _139_v88
		var _nw22 *C0 = New_C0_()
		_ = _nw22
		_nw22.Ctor__(Companion_Default___.SafeModuloInt((_138_v87).Cardinality(), _135_i6))
		_139_v88 = _nw22
		var _140_v89 _dafny.Array
		_ = _140_v89
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_3
		var _nw23 _dafny.Array
		_ = _nw23
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw23 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = func(_141_i7 _dafny.Int) bool {
				return true
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw23 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw23).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw23).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_140_v89 = _nw23
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_140_v89), 0))
		_ = _index27
		(_140_v89).ArraySet1(_49_v9, (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_140_v89), 0))
		_ = _index28
		(_140_v89).ArraySet1((_49_v9) || (_49_v9), (_index28).Int())
		var _142_v90 *C1
		_ = _142_v90
		var _nw24 *C1 = New_C1_()
		_ = _nw24
		_nw24.Ctor__(_135_i6, _dafny.SeqOf(true))
		_142_v90 = _nw24
		var _143_v91 _dafny.Map
		_ = _143_v91
		_143_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_142_v90), true)
		var _144_v92 D3
		_ = _144_v92
		_144_v92 = Companion_D3_.Create_DC8_(_142_v90)
		_143_v91 = (_143_v91).Update(_144_v92, (func() bool {
			if (_140_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_140_v89), 0))).Int()).(bool) {
				return true
			}
			return (_140_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_140_v89), 0))).Int()).(bool)
		})())
	}
	var _145_v93 _dafny.Set
	_ = _145_v93
	_145_v93 = _dafny.SetOf(!(_49_v9), _49_v9)
	r0 = _145_v93
	var _146_v94 _dafny.Array
	_ = _146_v94
	var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
	_ = _nw25
	_146_v94 = _nw25
	var _147_v95 _dafny.Sequence
	_ = _147_v95
	_147_v95 = _dafny.SeqOf(_dafny.IntOfInt64(580))
	var _148_v96 _dafny.Array
	_ = _148_v96
	var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
	_ = _nw26
	_148_v96 = _nw26
	var _149_v97 D8
	_ = _149_v97
	_149_v97 = Companion_D8_.Create_DC22_(_146_v94, _147_v95, _38_v0, _148_v96, _49_v9)
	r1 = (func() bool {
		if (_149_v97).Dtor_cf39() {
			return _49_v9
		}
		return _49_v9
	})()
	r2 = _54_v12
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _150_v0 bool
	_ = _150_v0
	_150_v0 = false
	var _151_v1 _dafny.Array
	_ = _151_v1
	var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
	_ = _nw27
	_151_v1 = _nw27
	var _152_v2 _dafny.Int
	_ = _152_v2
	_152_v2 = _dafny.IntOfInt64(204)
	var _153_v3 _dafny.Array
	_ = _153_v3
	var _nw28 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
	_ = _nw28
	_153_v3 = _nw28
	var _154_v4 _dafny.Map
	_ = _154_v4
	_154_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v2, _153_v3)
	var _155_v5 _dafny.Sequence
	_ = _155_v5
	_155_v5 = _dafny.SeqOf(true)
	var _156_v6 _dafny.Map
	_ = _156_v6
	_156_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v0, _152_v2)
	var _157_v7 _dafny.Sequence
	_ = _157_v7
	_157_v7 = _dafny.SeqOf(_150_v0, (_155_v5).Select((Companion_Default___.SafeIndex((_156_v6).Cardinality(), _dafny.IntOfUint32((_155_v5).Cardinality()))).Uint32()).(bool))
	var _158_v8 _dafny.Set
	_ = _158_v8
	_158_v8 = _dafny.SetOf(_152_v2)
	var _159_v9 _dafny.Set
	_ = _159_v9
	_159_v9 = _dafny.SetOf(_150_v0)
	var _160_v10 _dafny.Array
	_ = _160_v10
	var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
	_ = _nw29
	_160_v10 = _nw29
	var _161_globalState *GlobalState
	_ = _161_globalState
	var _nw30 *GlobalState = New_GlobalState_()
	_ = _nw30
	_nw30.Ctor__((func() _dafny.Array {
		if _150_v0 {
			return _151_v1
		}
		return _151_v1
	})(), _dafny.IntOfInt64(872), false, _dafny.IntOfInt64(91), _dafny.IntOfInt64(751), true, _dafny.IntOfInt64(-494), true, _dafny.IntOfInt64(204), _154_v4, false, true, _157_v7, _158_v8, false, _dafny.IntOfInt64(-442), _159_v9, _dafny.IntOfInt64(341), _dafny.IntOfInt64(699), _160_v10)
	_161_globalState = _nw30
	(_161_globalState).F3 = (((func() _dafny.Set {
		if _150_v0 {
			return _159_v9
		}
		return _159_v9
	})()).Intersection(_159_v9)).Cardinality()
	var _162_v11 _dafny.Array
	_ = _162_v11
	var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
	_ = _nw31
	_162_v11 = _nw31
	var _163_v12 _dafny.Sequence
	_ = _163_v12
	_163_v12 = _dafny.UnicodeSeqOfUtf8Bytes("skhwrlhb")
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))
	_ = _index29
	(_162_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_163_v12, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-478))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}((func(_164_v12 _dafny.Sequence, _165_v2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
		return func(_166_i0 _dafny.Int) _dafny.CodePoint {
			return (_164_v12).Select((Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_164_v12).Cardinality()))).Uint32()).(_dafny.CodePoint)
		}
	})(_163_v12, _152_v2)))), (_index29).Int())
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))
	_ = _index30
	(_162_v11).ArraySet1(_163_v12, (_index30).Int())
	if _150_v0 {
		var _167_v13 _dafny.Array
		_ = _167_v13
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_4
		var _nw32 _dafny.Array
		_ = _nw32
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw32 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Int = func(_168_i1 _dafny.Int) _dafny.Int {
				return (_168_i1).Times(_dafny.IntOfInt64(553))
			}
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw32 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw32).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw32).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_167_v13 = _nw32
		var _169_v14 _dafny.Sequence
		_ = _169_v14
		_169_v14 = _dafny.SeqOf(_163_v12)
		var _rhs16 _dafny.Array = _167_v13
		_ = _rhs16
		var _rhs17 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_169_v14).Cardinality()), (_152_v2).Plus(_152_v2)))
		_ = _rhs17
		var _rhs18 bool = !(true)
		_ = _rhs18
		var _lhs14 *GlobalState = _161_globalState
		_ = _lhs14
		_167_v13 = _rhs16
		_lhs14.F3 = _rhs17
		_150_v0 = _rhs18
		var _170_v15 _dafny.MultiSet
		_ = _170_v15
		_170_v15 = _dafny.MultiSetOf(_152_v2)
		var _171_v16 _dafny.Set
		_ = _171_v16
		var _172_v17 bool
		_ = _172_v17
		var _173_v18 _dafny.Map
		_ = _173_v18
		var _out0 _dafny.Set
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 _dafny.Map
		_ = _out2
		_out0, _out1, _out2 = Companion_Default___.M0(_170_v15, _161_globalState)
		_171_v16 = _out0
		_172_v17 = _out1
		_173_v18 = _out2
		var _174_v19 _dafny.Sequence
		_ = _174_v19
		_174_v19 = _dafny.SeqOf(_152_v2)
		var _175_v20 _dafny.Map
		_ = _175_v20
		_175_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v19, _152_v2)
		var _176_v21 _dafny.Map
		_ = _176_v21
		_176_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(32)).Plus(_dafny.IntOfInt64(-306)), (func() _dafny.Int {
			if (_175_v20).Contains(_dafny.SeqOf(_152_v2)) {
				return (_175_v20).Get(_dafny.SeqOf(_152_v2)).(_dafny.Int)
			}
			return _152_v2
		})())
		_176_v21 = (_176_v21).Update((_dafny.Zero).Minus(_152_v2), _dafny.IntOfUint32(((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence)).Cardinality()))
		var _177_v22 _dafny.Map
		_ = _177_v22
		_177_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_150_v0), _150_v0)
		_156_v6 = (_156_v6).Update((_152_v2).Cmp((_177_v22).Cardinality()) < 0, _152_v2)
		var _178_v23 *C1
		_ = _178_v23
		var _nw33 *C1 = New_C1_()
		_ = _nw33
		_nw33.Ctor__(_dafny.IntOfInt64(362), _157_v7)
		_178_v23 = _nw33
	} else {
		_152_v2 = _152_v2
		var _179_v24 *C1
		_ = _179_v24
		var _nw34 *C1 = New_C1_()
		_ = _nw34
		_nw34.Ctor__(_152_v2, _dafny.SeqOf(_150_v0))
		_179_v24 = _nw34
		var _180_v25 _dafny.Map
		_ = _180_v25
		_180_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v24, _dafny.IntOfInt64(-39))
		var _181_v26 _dafny.MultiSet
		_ = _181_v26
		_181_v26 = _dafny.MultiSetOf((_180_v25).Cardinality(), (_159_v9).Cardinality(), (_179_v24).F20(), _152_v2)
		var _182_v27 _dafny.Set
		_ = _182_v27
		var _183_v28 bool
		_ = _183_v28
		var _184_v29 _dafny.Map
		_ = _184_v29
		var _out3 _dafny.Set
		_ = _out3
		var _out4 bool
		_ = _out4
		var _out5 _dafny.Map
		_ = _out5
		_out3, _out4, _out5 = Companion_Default___.M0(_181_v26, _161_globalState)
		_182_v27 = _out3
		_183_v28 = _out4
		_184_v29 = _out5
		_152_v2 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_179_v24).F20(), (_179_v24).F20()), (_159_v9).Cardinality())
		(_161_globalState).F11 = (Companion_Default___.SafeDivisionInt(_152_v2, (_179_v24).F20())).Cmp((_179_v24).F20()) != 0
		var _185_v30 _dafny.Array
		_ = _185_v30
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_5
		var _nw35 _dafny.Array
		_ = _nw35
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw35 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = func(_186_i2 _dafny.Int) _dafny.Int {
				return (_186_i2).Times(_dafny.IntOfInt64(-99))
			}
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
		_185_v30 = _nw35
		(_161_globalState).F7 = (_179_v24).Fm0(_163_v12, (_179_v24).F20(), Companion_Default___.SafeDivisionInt(_152_v2, (_dafny.SetOf(_185_v30, _185_v30, _185_v30, _185_v30, _185_v30)).Cardinality()), _152_v2, _161_globalState)
	}
	var _187_v31 _dafny.MultiSet
	_ = _187_v31
	_187_v31 = _dafny.MultiSetOf(_150_v0, _150_v0, _150_v0)
	var _188_i3 _dafny.Int
	_ = _188_i3
	_188_i3 = _dafny.Zero
	{
		for (_155_v5).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_187_v31).Contains(_150_v0) {
				return (_187_v31).Multiplicity(_150_v0)
			}
			return _152_v2
		})(), _dafny.IntOfUint32((_155_v5).Cardinality()))).Uint32()).(bool) {
			{
				if (_188_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_188_i3 = (_188_i3).Plus(_dafny.One)
				var _189_v33 _dafny.Sequence
				_ = _189_v33
				_189_v33 = _dafny.SeqOf(_152_v2, _152_v2)
				var _190_v34 _dafny.Map
				_ = _190_v34
				_190_v34 = func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate((_189_v33).Elements()); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _191_v32 _dafny.Int
						_191_v32 = interface{}(_compr_12).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_189_v33, _191_v32) {
							_coll12.Add((_191_v32).Plus(_152_v2), _150_v0)
						}
					}
					return _coll12.ToMap()
				}()
				_190_v34 = Companion_Default___.Fm6((_dafny.Zero).Minus(Companion_Default___.Fm7((func() _dafny.Int {
					if (_187_v31).Contains(_150_v0) {
						return (_187_v31).Multiplicity(_150_v0)
					}
					return _152_v2
				})(), _152_v2, _161_globalState)), _161_globalState)
				var _192_v35 _dafny.Map
				_ = _192_v35
				_192_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_190_v34, _dafny.UnicodeSeqOfUtf8Bytes("syahspiq"))
				_192_v35 = (_192_v35).Update(_190_v34, _163_v12)
				var _193_v36 _dafny.Array
				_ = _193_v36
				var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw36
				_193_v36 = _nw36
				_193_v36 = _193_v36
				(_161_globalState).F14 = (Companion_Default___.SafeModuloInt(_152_v2, (_dafny.Zero).Minus(_152_v2))).Cmp((Companion_Default___.Fm8(_161_globalState)).Dtor_cf2()) == 0
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _194_v37 _dafny.Array
	_ = _194_v37
	var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
	_ = _nw37
	_194_v37 = _nw37
	var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))
	_ = _index31
	(_194_v37).ArraySet1(_152_v2, (_index31).Int())
	var _195_v38 _dafny.Sequence
	_ = _195_v38
	_195_v38 = _dafny.SeqOf(_163_v12, _dafny.UnicodeSeqOfUtf8Bytes("qsf"), _dafny.Companion_Sequence_.Concatenate((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence), _163_v12), _163_v12, _163_v12)
	var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))
	_ = _index32
	var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))
	_ = _index33
	var _rhs19 _dafny.Sequence = (_195_v38).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_152_v2), _dafny.IntOfUint32((_195_v38).Cardinality()))).Uint32()).(_dafny.Sequence)
	_ = _rhs19
	var _rhs20 _dafny.Int = Companion_Default___.SafeDivisionInt(_152_v2, _152_v2)
	_ = _rhs20
	var _lhs15 _dafny.Array = _162_v11
	_ = _lhs15
	var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))
	_ = _lhs16
	var _lhs17 _dafny.Array = _194_v37
	_ = _lhs17
	var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))
	_ = _lhs18
	(_lhs15).ArraySet1(_rhs19, (_lhs16).Int())
	(_lhs17).ArraySet1(_rhs20, (_lhs18).Int())
	var _196_v39 D1
	_ = _196_v39
	_196_v39 = Companion_D1_.Create_DC2_(_152_v2, _150_v0)
	var _pat_let_tv1 = _150_v0
	_ = _pat_let_tv1
	var _pat_let_tv2 = _194_v37
	_ = _pat_let_tv2
	var _pat_let_tv3 = _194_v37
	_ = _pat_let_tv3
	var _pat_let_tv4 = _152_v2
	_ = _pat_let_tv4
	var _pat_let_tv5 = _152_v2
	_ = _pat_let_tv5
	var _pat_let_tv6 = _194_v37
	_ = _pat_let_tv6
	var _pat_let_tv7 = _194_v37
	_ = _pat_let_tv7
	(_161_globalState).F7 = func(_source2 D1) bool {
		if _source2.Is_DC2() {
			var _197___mcc_h0 _dafny.Int = _source2.Get_().(D1_DC2).Cf2
			_ = _197___mcc_h0
			var _198___mcc_h1 bool = _source2.Get_().(D1_DC2).Cf3
			_ = _198___mcc_h1
			var _199_cf3 bool = _198___mcc_h1
			_ = _199_cf3
			var _200_cf2 _dafny.Int = _197___mcc_h0
			_ = _200_cf2
			return !(_pat_let_tv1) || (_199_cf3)
		} else if _source2.Is_DC1() {
			var _201___mcc_h2 _dafny.Int = _source2.Get_().(D1_DC1).Cf1
			_ = _201___mcc_h2
			var _202_cf1 _dafny.Int = _201___mcc_h2
			_ = _202_cf1
			return ((_pat_let_tv3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_pat_let_tv2), 0))).Int()).(_dafny.Int)).Cmp(_pat_let_tv4) <= 0
		} else {
			var _203___mcc_h3 D1 = _source2.Get_().(D1_DC3).Cf4
			_ = _203___mcc_h3
			var _204_cf4 D1 = _203___mcc_h3
			_ = _204_cf4
			return (_pat_let_tv5).Cmp((_pat_let_tv7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_pat_let_tv6), 0))).Int()).(_dafny.Int)) >= 0
		}
	}(_196_v39)
	var _205_v40 _dafny.Map
	_ = _205_v40
	_205_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int), _150_v0)
	var _206_v41 _dafny.Map
	_ = _206_v41
	_206_v41 = _205_v40
	var _pat_let_tv8 = _152_v2
	_ = _pat_let_tv8
	(_161_globalState).F3 = func(_source3 _dafny.Map) _dafny.Int {
		var _207___mcc_h4 _dafny.Map = _source3
		_ = _207___mcc_h4
		var _208_cf0 _dafny.Map = _207___mcc_h4
		_ = _208_cf0
		return _pat_let_tv8
	}(_206_v41)
	var _209_v42 _dafny.CodePoint
	_ = _209_v42
	_209_v42 = _dafny.CodePoint('d')
	var _210_v43 _dafny.Set
	_ = _210_v43
	_210_v43 = _dafny.SetOf(_209_v42, _209_v42)
	var _211_v44 _dafny.Map
	_ = _211_v44
	_211_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v2, _210_v43)
	var _212_v45 _dafny.MultiSet
	_ = _212_v45
	_212_v45 = _dafny.MultiSetOf((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int), (_159_v9).Cardinality(), (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(26), _dafny.IntOfUint32(((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence)).Cardinality()))
	(_161_globalState).F3 = (Companion_Default___.SafeModuloInt(_152_v2, (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int))).Times(Companion_Default___.Fm7(((func() _dafny.Set {
		if (_211_v44).Contains((func() _dafny.Int {
			if (_187_v31).Contains(_150_v0) {
				return (_187_v31).Multiplicity(_150_v0)
			}
			return (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int)
		})()) {
			return (_211_v44).Get((func() _dafny.Int {
				if (_187_v31).Contains(_150_v0) {
					return (_187_v31).Multiplicity(_150_v0)
				}
				return (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int)
			})()).(_dafny.Set)
		}
		return _210_v43
	})()).Cardinality(), (func() _dafny.Int {
		if (_212_v45).Contains(_dafny.IntOfUint32(((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence)).Cardinality())) {
			return (_212_v45).Multiplicity(_dafny.IntOfUint32(((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence)).Cardinality()))
		}
		return (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int)
	})(), _161_globalState))
	var _213_v46 _dafny.Map
	_ = _213_v46
	_213_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v0, _212_v45)
	var _214_v47 _dafny.Sequence
	_ = _214_v47
	_214_v47 = _dafny.SeqOf(_152_v2)
	var _215_v48 D2
	_ = _215_v48
	_215_v48 = Companion_D2_.Create_DC4_(_dafny.MultiSetFromSeq(_214_v47))
	var _pat_let_tv9 = _152_v2
	_ = _pat_let_tv9
	var _216_v49 _dafny.Sequence
	_ = _216_v49
	_216_v49 = _dafny.SeqOf((func(_pat_let2_0 D2) D2 {
		return func(_217_dt__update__tmp_h0 D2) D2 {
			return func(_pat_let3_0 _dafny.MultiSet) D2 {
				return func(_218_dt__update_hcf5_h0 _dafny.MultiSet) D2 {
					return Companion_D2_.Create_DC4_(_218_dt__update_hcf5_h0)
				}(_pat_let3_0)
			}(_dafny.MultiSetOf(_pat_let_tv9, _dafny.IntOfInt64(765)))
		}(_pat_let2_0)
	}(_215_v48)).Dtor_cf5(), _212_v45)
	var _219_v50 _dafny.Map
	_ = _219_v50
	_219_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _163_v12)
	(_161_globalState).F14 = !((_212_v45).Union(_212_v45)).Equals((func() _dafny.MultiSet {
		if (_213_v46).Contains(_150_v0) {
			return (_213_v46).Get(_150_v0).(_dafny.MultiSet)
		}
		return (_216_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_219_v50).Contains(_150_v0) {
				return (_219_v50).Get(_150_v0).(_dafny.Sequence)
			}
			return (_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence)
		})()).Cardinality()), _dafny.IntOfUint32((_216_v49).Cardinality()))).Uint32()).(_dafny.MultiSet)
	})())
	var _hi1 _dafny.Int = (func() _dafny.Int {
		if (_156_v6).Contains(_150_v0) {
			return (_156_v6).Get(_150_v0).(_dafny.Int)
		}
		return _152_v2
	})()
	_ = _hi1
	for _220_i4 := (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int); _220_i4.Cmp(_hi1) < 0; _220_i4 = _220_i4.Plus(_dafny.One) {
		var _221_v51 _dafny.Map
		_ = _221_v51
		_221_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_205_v40, (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int))
		var _222_v52 _dafny.Sequence
		_ = _222_v52
		_222_v52 = _dafny.SeqOf(_221_v51)
		var _223_v53 _dafny.Map
		_ = _223_v53
		_223_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v2, _152_v2)
		var _224_v54 _dafny.Set
		_ = _224_v54
		var _225_v55 bool
		_ = _225_v55
		var _226_v56 _dafny.Map
		_ = _226_v56
		var _out6 _dafny.Set
		_ = _out6
		var _out7 bool
		_ = _out7
		var _out8 _dafny.Map
		_ = _out8
		_out6, _out7, _out8 = Companion_Default___.M0(Companion_Default___.Fm9(false, (_163_v12).Select((Companion_Default___.SafeIndex((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_163_v12).Cardinality()))).Uint32()).(_dafny.CodePoint), (_222_v52).Select((Companion_Default___.SafeIndex(_152_v2, _dafny.IntOfUint32((_222_v52).Cardinality()))).Uint32()).(_dafny.Map), _223_v53, _161_globalState), _161_globalState)
		_224_v54 = _out6
		_225_v55 = _out7
		_226_v56 = _out8
		(_161_globalState).F3 = (_dafny.Zero).Minus(Companion_Default___.Fm7(_152_v2, _dafny.IntOfUint32(((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence)).Cardinality()), _161_globalState))
		(_161_globalState).F7 = !(_150_v0)
		var _227_v57 _dafny.Map
		_ = _227_v57
		_227_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("skjwiqufa")).Cardinality())).Cmp((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int)) != 0)
		_227_v57 = (_227_v57).Update(_150_v0, _150_v0)
	}
	var _228_v58 D1
	_ = _228_v58
	_228_v58 = Companion_D1_.Create_DC2_(Companion_Default___.Fm7(_152_v2, _152_v2, _161_globalState), (_155_v5).Select((Companion_Default___.SafeIndex((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_155_v5).Cardinality()))).Uint32()).(bool))
	var _229_v59 D1
	_ = _229_v59
	_229_v59 = Companion_D1_.Create_DC3_(_228_v58)
	var _230_v60 _dafny.Map
	_ = _230_v60
	_230_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v59, _209_v42)
	var _231_v61 _dafny.Map
	_ = _231_v61
	_231_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v60, (_150_v0) || (true))
	_231_v61 = (_231_v61).Update(_230_v60, (func() bool {
		if _150_v0 {
			return _150_v0
		}
		return _150_v0
	})())
	var _232_v62 _dafny.Map
	_ = _232_v62
	_232_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_209_v42, _194_v37)
	var _hi2 _dafny.Int = (_232_v62).Cardinality()
	_ = _hi2
	for _233_i5 := _152_v2; _233_i5.Cmp(_hi2) < 0; _233_i5 = _233_i5.Plus(_dafny.One) {
		var _234_v63 *C1
		_ = _234_v63
		var _nw38 *C1 = New_C1_()
		_ = _nw38
		_nw38.Ctor__(_152_v2, _dafny.Companion_Sequence_.Concatenate(_157_v7, _157_v7))
		_234_v63 = _nw38
		(_161_globalState).F3 = (_233_i5).Times((_234_v63).F20())
		(_161_globalState).F11 = _150_v0
		var _235_v64 D1
		_ = _235_v64
		_235_v64 = Companion_D1_.Create_DC1_(_233_i5)
		var _source4 D1 = _235_v64
		_ = _source4
		if _source4.Is_DC2() {
			var _236___mcc_h5 _dafny.Int = _source4.Get_().(D1_DC2).Cf2
			_ = _236___mcc_h5
			var _237___mcc_h6 bool = _source4.Get_().(D1_DC2).Cf3
			_ = _237___mcc_h6
			var _238_cf3 bool = _237___mcc_h6
			_ = _238_cf3
			var _239_cf2 _dafny.Int = _236___mcc_h5
			_ = _239_cf2
			var _240_v65 D2
			_ = _240_v65
			_240_v65 = Companion_D2_.Create_DC6_((_234_v63).F20())
			var _241_v66 _dafny.Map
			_ = _241_v66
			_241_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_234_v63, (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int))
			var _242_v67 _dafny.Sequence
			_ = _242_v67
			_242_v67 = _dafny.SeqOf(_241_v66, _241_v66)
			var _243_v68 _dafny.Map
			_ = _243_v68
			_243_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_152_v2), _209_v42)
			var _244_v69 _dafny.Sequence
			_ = _244_v69
			_244_v69 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_242_v67).Cardinality()), _209_v42), _243_v68, _243_v68, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_234_v63).F20(), _209_v42))
			var _pat_let_tv10 = _244_v69
			_ = _pat_let_tv10
			var _245_v70 _dafny.Array
			_ = _245_v70
			var _nwElement0_6 _dafny.Set = _159_v9
			_ = _nwElement0_6
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(18))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_6, 0)
			(_nw39).ArraySet1(_159_v9, 1)
			(_nw39).ArraySet1(_159_v9, 2)
			(_nw39).ArraySet1(_159_v9, 3)
			(_nw39).ArraySet1(_159_v9, 4)
			(_nw39).ArraySet1(Companion_Default___.Fm10(_240_v65, _161_globalState), 5)
			(_nw39).ArraySet1(_159_v9, 6)
			(_nw39).ArraySet1((_159_v9).Difference(_159_v9), 7)
			(_nw39).ArraySet1(Companion_Default___.Fm10(func(_pat_let4_0 D2) D2 {
				return func(_246_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let5_0 _dafny.Int) D2 {
						return func(_247_dt__update_hcf9_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC6_(_247_dt__update_hcf9_h0)
						}(_pat_let5_0)
					}((_dafny.MultiSetFromSeq(_pat_let_tv10)).Cardinality())
				}(_pat_let4_0)
			}(_240_v65), _161_globalState), 8)
			(_nw39).ArraySet1(_dafny.SetOf(_150_v0, _150_v0, _150_v0), 9)
			(_nw39).ArraySet1((_159_v9).Difference(_159_v9), 10)
			(_nw39).ArraySet1(_159_v9, 11)
			(_nw39).ArraySet1(_159_v9, 12)
			(_nw39).ArraySet1(_159_v9, 13)
			(_nw39).ArraySet1(_159_v9, 14)
			(_nw39).ArraySet1(_159_v9, 15)
			(_nw39).ArraySet1(_159_v9, 16)
			(_nw39).ArraySet1(Companion_Default___.Fm10(_240_v65, _161_globalState), 17)
			_245_v70 = _nw39
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_245_v70), 0))
			_ = _index34
			(_245_v70).ArraySet1((_dafny.SetOf(!(_238_cf3))).Difference(_159_v9), (_index34).Int())
			var _248_v71 _dafny.Array
			_ = _248_v71
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw40
			_248_v71 = _nw40
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_248_v71), 0))
			_ = _index35
			(_248_v71).ArraySet1(_162_v11, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))
			_ = _index36
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_245_v70), 0))
			_ = _index37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_248_v71), 0))
			_ = _index38
			var _rhs21 _dafny.Int = (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int)
			_ = _rhs21
			var _rhs22 _dafny.Int = ((_dafny.Zero).Minus((_234_v63).F20())).Plus(_dafny.IntOfInt64(107))
			_ = _rhs22
			var _rhs23 _dafny.Set = _159_v9
			_ = _rhs23
			var _rhs24 _dafny.Array = _162_v11
			_ = _rhs24
			var _lhs19 *GlobalState = _161_globalState
			_ = _lhs19
			var _lhs20 _dafny.Array = _194_v37
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))
			_ = _lhs21
			var _lhs22 _dafny.Array = _245_v70
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_245_v70), 0))
			_ = _lhs23
			var _lhs24 _dafny.Array = _248_v71
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_248_v71), 0))
			_ = _lhs25
			_lhs19.F3 = _rhs21
			(_lhs20).ArraySet1(_rhs22, (_lhs21).Int())
			(_lhs22).ArraySet1(_rhs23, (_lhs23).Int())
			(_lhs24).ArraySet1(_rhs24, (_lhs25).Int())
			_156_v6 = ((_156_v6).Update(_150_v0, _239_cf2)).Merge(_156_v6)
			(_161_globalState).F7 = _150_v0
			(_161_globalState).F14 = (_234_v63).Fm0((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence), (_dafny.IntOfInt64(17)).Times(_152_v2), _152_v2, ((_212_v45).Cardinality()).Plus((_234_v63).F20()), _161_globalState)
		} else if _source4.Is_DC1() {
			var _249___mcc_h7 _dafny.Int = _source4.Get_().(D1_DC1).Cf1
			_ = _249___mcc_h7
			var _250_cf1 _dafny.Int = _249___mcc_h7
			_ = _250_cf1
			_215_v48 = _215_v48
			var _251_v72 _dafny.Map
			_ = _251_v72
			_251_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqOf(_153_v3, _153_v3)).Cardinality())), _214_v47)
			_251_v72 = (_251_v72).Update(_235_v64, _dafny.Companion_Sequence_.Concatenate(_214_v47, _214_v47))
			(_161_globalState).F11 = !((_159_v9).Intersection(_159_v9)).Equals((_159_v9).Union(_159_v9))
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))
			_ = _index39
			(_162_v11).ArraySet1((_162_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_162_v11), 0))).Int()).(_dafny.Sequence), (_index39).Int())
		} else {
			var _252___mcc_h8 D1 = _source4.Get_().(D1_DC3).Cf4
			_ = _252___mcc_h8
			var _253_cf4 D1 = _252___mcc_h8
			_ = _253_cf4
			_205_v40 = _205_v40
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))
			_ = _index40
			var _rhs25 bool = _150_v0
			_ = _rhs25
			var _rhs26 _dafny.Int = _233_i5
			_ = _rhs26
			var _rhs27 bool = (_150_v0) && (!_dafny.Companion_Sequence_.Contains(_163_v12, _209_v42))
			_ = _rhs27
			var _rhs28 _dafny.Int = (_dafny.Zero).Minus(((_205_v40).Cardinality()).Times(_dafny.IntOfUint32((_163_v12).Cardinality())))
			_ = _rhs28
			var _lhs26 *GlobalState = _161_globalState
			_ = _lhs26
			var _lhs27 _dafny.Array = _194_v37
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))
			_ = _lhs28
			var _lhs29 *GlobalState = _161_globalState
			_ = _lhs29
			_lhs26.F7 = _rhs25
			(_lhs27).ArraySet1(_rhs26, (_lhs28).Int())
			_150_v0 = _rhs27
			_lhs29.F3 = _rhs28
			_157_v7 = _155_v5
			_157_v7 = _157_v7
		}
	}
	var _254_v73 _dafny.Array
	_ = _254_v73
	var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
	_ = _nw41
	_254_v73 = _nw41
	_254_v73 = _254_v73
	(_161_globalState).F3 = ((_187_v31).Intersection((_187_v31).Difference(_187_v31))).Cardinality()
	var _255_v74 _dafny.Set
	_ = _255_v74
	var _256_v75 bool
	_ = _256_v75
	var _257_v76 _dafny.Map
	_ = _257_v76
	var _out9 _dafny.Set
	_ = _out9
	var _out10 bool
	_ = _out10
	var _out11 _dafny.Map
	_ = _out11
	_out9, _out10, _out11 = Companion_Default___.M0(_dafny.MultiSetOf(_dafny.IntOfInt64(857), (_dafny.Zero).Minus((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int)), (_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_195_v38).Cardinality()), _152_v2), _161_globalState)
	_255_v74 = _out9
	_256_v75 = _out10
	_257_v76 = _out11
	var _258_v77 *C1
	_ = _258_v77
	var _nw42 *C1 = New_C1_()
	_ = _nw42
	_nw42.Ctor__(Companion_Default___.Fm7((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-177), _161_globalState), _155_v5)
	_258_v77 = _nw42
	var _259_v78 D3
	_ = _259_v78
	_259_v78 = Companion_D3_.Create_DC8_(_258_v77)
	var _260_v79 _dafny.Set
	_ = _260_v79
	_260_v79 = _dafny.SetOf((_259_v78).Dtor_cf12())
	var _rhs29 bool = true
	_ = _rhs29
	var _rhs30 _dafny.Int = ((_194_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_194_v37), 0))).Int()).(_dafny.Int)).Plus(_152_v2)
	_ = _rhs30
	var _rhs31 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
		if (_187_v31).Contains(_150_v0) {
			return (_187_v31).Multiplicity(_150_v0)
		}
		return (_260_v79).Cardinality()
	})())
	_ = _rhs31
	var _lhs30 *GlobalState = _161_globalState
	_ = _lhs30
	_150_v0 = _rhs29
	_152_v2 = _rhs30
	_lhs30.F3 = _rhs31
	_dafny.Print(_150_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_152_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v4).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_155_v5, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(204))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_157_v7, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v8).Equals(_dafny.SetOf(_dafny.IntOfInt64(204))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v9).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_161_globalState).F9()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_161_globalState).F12(), _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_161_globalState).F13()).Equals(_dafny.SetOf(_dafny.IntOfInt64(204))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_161_globalState).F16()).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_v11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_163_v12.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v31).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_188_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_194_v37).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_195_v38, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("skhwrlhb"), _dafny.UnicodeSeqOfUtf8Bytes("qsf"), _dafny.UnicodeSeqOfUtf8Bytes("skhwrlhbskhwrlhb"), _dafny.UnicodeSeqOfUtf8Bytes("skhwrlhb"), _dafny.UnicodeSeqOfUtf8Bytes("skhwrlhb"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_196_v39).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_196_v39).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v40).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v41).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_209_v42)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_210_v43).Equals(_dafny.SetOf(_dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_211_v44).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.SetOf(_dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v45).Equals(_dafny.MultiSetOf(_dafny.Zero, _dafny.Zero, _dafny.One, _dafny.IntOfInt64(26), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_213_v46).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(_dafny.Zero, _dafny.Zero, _dafny.One, _dafny.IntOfInt64(26), _dafny.IntOfInt64(8)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_214_v47, _dafny.SeqOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_215_v48).Dtor_cf5()).Equals(_dafny.MultiSetOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_216_v49, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.Zero, _dafny.IntOfInt64(765)), _dafny.MultiSetOf(_dafny.Zero, _dafny.Zero, _dafny.One, _dafny.IntOfInt64(26), _dafny.IntOfInt64(8)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v50).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("skhwrlhb"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_v58).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_v58).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_229_v59).Dtor_cf4()).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_229_v59).Dtor_cf4()).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v60).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(7608), true)), _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_231_v61).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(7608), true)), _dafny.CodePoint('d')), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_232_v62).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v74).Equals(_dafny.SetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_256_v75)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v76).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(7257), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v77).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_258_v77).F21(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_v78).Dtor_cf12()).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_259_v78).Dtor_cf12()).F21(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v79).Cardinality())
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
	Cf0 _dafny.Map
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Map) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D0) Dtor_cf0() _dafny.Map {
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

type D1_DC2 struct {
	Cf2 _dafny.Int
	Cf3 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 bool) D1 {
	return D1{D1_DC2{Cf2, Cf3}}
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

type D1_DC3 struct {
	Cf4 D1
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 D1) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, false)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf4() D1 {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4.Equals(data2.Cf4)
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
	Cf6 bool
	Cf7 bool
	Cf8 _dafny.Set
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf6 bool, Cf7 bool, Cf8 _dafny.Set) D2 {
	return D2{D2_DC5{Cf6, Cf7, Cf8}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC6 struct {
	Cf9 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf9 _dafny.Int) D2 {
	return D2{D2_DC6{Cf9}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf10 _dafny.Int
	Cf11 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf10 _dafny.Int, Cf11 _dafny.Int) D2 {
	return D2{D2_DC7{Cf10, Cf11}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC4 struct {
	Cf5 _dafny.MultiSet
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 _dafny.MultiSet) D2 {
	return D2{D2_DC4{Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(false, false, _dafny.EmptySet)
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Set {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf10
}

func (_this D2) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf5() _dafny.MultiSet {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7 && data1.Cf8.Equals(data2.Cf8)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf5.Equals(data2.Cf5)
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
	Cf13 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf13 _dafny.Int) D3 {
	return D3{D3_DC9{Cf13}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC8 struct {
	Cf12 *C1
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf12 *C1) D3 {
	return D3{D3_DC8{Cf12}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.Zero)
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf13
}

func (_this D3) Dtor_cf12() *C1 {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf12 == data2.Cf12
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
	Cf15 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf15 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf15}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf16 bool
	Cf17 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf16 bool, Cf17 _dafny.Int) D4 {
	return D4{D4_DC12{Cf16, Cf17}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC10 struct {
	Cf14 _dafny.Array
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf14 _dafny.Array) D4 {
	return D4{D4_DC10{Cf14}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC13 struct {
	Cf18 D4
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf18 D4) D4 {
	return D4{D4_DC13{Cf18}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.EmptySeq)
}

func (_this D4) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf15
}

func (_this D4) Dtor_cf16() bool {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf17
}

func (_this D4) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D4_DC10).Cf14
}

func (_this D4) Dtor_cf18() D4 {
	return _this.Get_().(D4_DC13).Cf18
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf15.Equals(data2.Cf15)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf14 == data2.Cf14
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 bool
	Cf21 bool
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf20 bool, Cf21 bool) D5 {
	return D5{D5_DC15{Cf20, Cf21}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC16 struct {
	Cf22 _dafny.Int
	Cf23 bool
	Cf24 bool
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf22 _dafny.Int, Cf23 bool, Cf24 bool) D5 {
	return D5{D5_DC16{Cf22, Cf23, Cf24}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC14 struct {
	Cf19 _dafny.Sequence
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf19 _dafny.Sequence) D5 {
	return D5{D5_DC14{Cf19}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(false, false)
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC15).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC15).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf22
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC16).Cf23
}

func (_this D5) Dtor_cf24() bool {
	return _this.Get_().(D5_DC16).Cf24
}

func (_this D5) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D5_DC14).Cf19
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + data.Cf19.VerbatimString(true) + ")"
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
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23 == data2.Cf23 && data1.Cf24 == data2.Cf24
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	Cf26 bool
	Cf27 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf26 bool, Cf27 _dafny.Int) D6 {
	return D6{D6_DC18{Cf26, Cf27}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC19 struct {
	Cf28 bool
	Cf29 _dafny.Int
	Cf30 _dafny.Int
	Cf31 _dafny.Int
	Cf32 _dafny.Int
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf28 bool, Cf29 _dafny.Int, Cf30 _dafny.Int, Cf31 _dafny.Int, Cf32 _dafny.Int) D6 {
	return D6{D6_DC19{Cf28, Cf29, Cf30, Cf31, Cf32}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC17 struct {
	Cf25 _dafny.CodePoint
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf25 _dafny.CodePoint) D6 {
	return D6{D6_DC17{Cf25}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC18_(false, _dafny.Zero)
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC18).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf27
}

func (_this D6) Dtor_cf28() bool {
	return _this.Get_().(D6_DC19).Cf28
}

func (_this D6) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf29
}

func (_this D6) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf30
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf32
}

func (_this D6) Dtor_cf25() _dafny.CodePoint {
	return _this.Get_().(D6_DC17).Cf25
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32.Cmp(data2.Cf32) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf25 == data2.Cf25
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
	Cf33 *C0
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf33 *C0) D7 {
	return D7{D7_DC20{Cf33}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() *C0 {
	return (*C0)(nil)
}

func (_this D7) Dtor_cf33() *C0 {
	return _this.Get_().(D7_DC20).Cf33
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf33) + ")"
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
			return ok && data1.Cf33 == data2.Cf33
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
	Cf35 _dafny.Array
	Cf36 _dafny.Sequence
	Cf37 _dafny.Int
	Cf38 _dafny.Array
	Cf39 bool
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf35 _dafny.Array, Cf36 _dafny.Sequence, Cf37 _dafny.Int, Cf38 _dafny.Array, Cf39 bool) D8 {
	return D8{D8_DC22{Cf35, Cf36, Cf37, Cf38, Cf39}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC23 struct {
	Cf40 *C1
	Cf41 bool
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf40 *C1, Cf41 bool) D8 {
	return D8{D8_DC23{Cf40, Cf41}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf42 bool
	Cf43 D1
	Cf44 _dafny.Int
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf42 bool, Cf43 D1, Cf44 _dafny.Int) D8 {
	return D8{D8_DC24{Cf42, Cf43, Cf44}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC21 struct {
	Cf34 _dafny.Sequence
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf34 _dafny.Sequence) D8 {
	return D8{D8_DC21{Cf34}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC25 struct {
	Cf45 D8
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf45 D8) D8 {
	return D8{D8_DC25{Cf45}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC22_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D8) Dtor_cf35() _dafny.Array {
	return _this.Get_().(D8_DC22).Cf35
}

func (_this D8) Dtor_cf36() _dafny.Sequence {
	return _this.Get_().(D8_DC22).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D8_DC22).Cf37
}

func (_this D8) Dtor_cf38() _dafny.Array {
	return _this.Get_().(D8_DC22).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC22).Cf39
}

func (_this D8) Dtor_cf40() *C1 {
	return _this.Get_().(D8_DC23).Cf40
}

func (_this D8) Dtor_cf41() bool {
	return _this.Get_().(D8_DC23).Cf41
}

func (_this D8) Dtor_cf42() bool {
	return _this.Get_().(D8_DC24).Cf42
}

func (_this D8) Dtor_cf43() D1 {
	return _this.Get_().(D8_DC24).Cf43
}

func (_this D8) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf44
}

func (_this D8) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D8_DC21).Cf34
}

func (_this D8) Dtor_cf45() D8 {
	return _this.Get_().(D8_DC25).Cf45
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf45) + ")"
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
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36.Equals(data2.Cf36) && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41 == data2.Cf41
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf42 == data2.Cf42 && data1.Cf43.Equals(data2.Cf43) && data1.Cf44.Cmp(data2.Cf44) == 0
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf45.Equals(data2.Cf45)
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
	F3   _dafny.Int
	F5   bool
	F7   bool
	F11  bool
	F14  bool
	_f0  _dafny.Array
	_f1  _dafny.Int
	_f2  bool
	_f4  _dafny.Int
	_f6  _dafny.Int
	_f8  _dafny.Int
	_f9  _dafny.Map
	_f10 bool
	_f12 _dafny.Sequence
	_f13 _dafny.Set
	_f15 _dafny.Int
	_f16 _dafny.Set
	_f17 _dafny.Int
	_f18 _dafny.Int
	_f19 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = _dafny.Zero
	_this.F5 = false
	_this.F7 = false
	_this.F11 = false
	_this.F14 = false
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.Zero
	_this._f2 = false
	_this._f4 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.EmptyMap
	_this._f10 = false
	_this._f12 = _dafny.EmptySeq
	_this._f13 = _dafny.EmptySet
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.EmptySet
	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.Zero
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Int, f2 bool, f3 _dafny.Int, f4 _dafny.Int, f5 bool, f6 _dafny.Int, f7 bool, f8 _dafny.Int, f9 _dafny.Map, f10 bool, f11 bool, f12 _dafny.Sequence, f13 _dafny.Set, f14 bool, f15 _dafny.Int, f16 _dafny.Set, f17 _dafny.Int, f18 _dafny.Int, f19 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Map {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() _dafny.Sequence {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Set {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Set {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Int {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() _dafny.Array {
	{
		return _this._f19
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f22 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f22 = _dafny.Zero
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

func (_this *C0) Ctor__(f22 _dafny.Int) {
	{
		(_this)._f22 = f22
	}
}
func (_this *C0) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), (_dafny.MultiSetOf(true, !(!(true)))).Contains(false))
	}
}
func (_this *C0) Fm4(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('g')
	}
}
func (_this *C0) F22() _dafny.Int {
	{
		return _this._f22
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f20 _dafny.Int
	_f21 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f20 = _dafny.Zero
	_this._f21 = _dafny.EmptySeq
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

func (_this *C1) Ctor__(f20 _dafny.Int, f21 _dafny.Sequence) {
	{
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C1) Fm0(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(false) || (((_this).F21()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_261_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()), _dafny.IntOfUint32(((_this).F21()).Cardinality()))).Uint32()).(bool))
	}
}
func (_this *C1) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F20()).Minus((_this).F20())
	}
}
func (_this *C1) M1(globalState *GlobalState) {
	{
		var _262_v0 bool
		_ = _262_v0
		_262_v0 = false
		var _263_v1 _dafny.Map
		_ = _263_v1
		_263_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _262_v0)
		var _264_v3 _dafny.Sequence
		_ = _264_v3
		_264_v3 = _dafny.SeqOf(_dafny.IntOfInt64(50))
		var _265_v5 _dafny.Map
		_ = _265_v5
		_265_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F21()).Cardinality()), (_this).F20())
		var _266_v6 _dafny.Map
		_ = _266_v6
		_266_v6 = func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_265_v5).Keys().Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _267_v4 _dafny.Int
				_267_v4 = interface{}(_compr_13).(_dafny.Int)
				if (_265_v5).Contains(_267_v4) {
					_coll13.Add(Companion_Default___.SafeDivisionInt(_267_v4, (_this).F20()), _262_v0)
				}
			}
			return _coll13.ToMap()
		}()
		var _268_v7 _dafny.Map
		_ = _268_v7
		_268_v7 = (_266_v6)
		var _269_v9 _dafny.Sequence
		_ = _269_v9
		_269_v9 = _dafny.SeqOf(func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(207), _dafny.IntOfInt64(401))); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _270_v8 _dafny.Int
				_270_v8 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(207)).Cmp(_270_v8) <= 0) && ((_270_v8).Cmp(_dafny.IntOfInt64(401)) < 0) {
					_coll14.Add(Companion_Default___.SafeDivisionInt(_270_v8, (_this).F20()), _262_v0)
				}
			}
			return _coll14.ToMap()
		}())
		var _271_v10 _dafny.Array
		_ = _271_v10
		var _nwElement0_7 _dafny.Map = (_263_v1).Merge(_263_v1)
		_ = _nwElement0_7
		var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(20))
		_ = _nw43
		(_nw43).ArraySet1(_nwElement0_7, 0)
		(_nw43).ArraySet1((_263_v1).Merge(_263_v1), 1)
		(_nw43).ArraySet1(func() _dafny.Map {
			var _coll15 = _dafny.NewMapBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate((_264_v3).Elements()); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _272_v2 _dafny.Int
				_272_v2 = interface{}(_compr_15).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_264_v3, _272_v2) {
					_coll15.Add((_272_v2).Plus((_this).F20()), _262_v0)
				}
			}
			return _coll15.ToMap()
		}(), 2)
		(_nw43).ArraySet1((_263_v1).Merge((_263_v1).Update((_this).F20(), _262_v0)), 3)
		(_nw43).ArraySet1(_263_v1, 4)
		(_nw43).ArraySet1(_263_v1, 5)
		(_nw43).ArraySet1(_263_v1, 6)
		(_nw43).ArraySet1((_263_v1).Update((_this).F20(), true), 7)
		(_nw43).ArraySet1(_263_v1, 8)
		(_nw43).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(69), _262_v0), 9)
		(_nw43).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _262_v0), 10)
		(_nw43).ArraySet1(_263_v1, 11)
		(_nw43).ArraySet1(_263_v1, 12)
		(_nw43).ArraySet1((_268_v7), 13)
		(_nw43).ArraySet1((_269_v9).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_269_v9).Cardinality()))).Uint32()).(_dafny.Map), 14)
		(_nw43).ArraySet1(_263_v1, 15)
		(_nw43).ArraySet1(_263_v1, 16)
		(_nw43).ArraySet1(_263_v1, 17)
		(_nw43).ArraySet1((_263_v1).Merge(_263_v1), 18)
		(_nw43).ArraySet1(_263_v1, 19)
		_271_v10 = _nw43
		var _273_v11 _dafny.Sequence
		_ = _273_v11
		_273_v11 = _dafny.UnicodeSeqOfUtf8Bytes("mjg")
		var _rhs32 _dafny.Array = _271_v10
		_ = _rhs32
		var _rhs33 _dafny.Int = (_this).F20()
		_ = _rhs33
		var _rhs34 bool = !((_this).Fm0(_273_v11, (_this).F20(), (_this).F20(), (_this).F20(), globalState))
		_ = _rhs34
		var _lhs31 *GlobalState = globalState
		_ = _lhs31
		var _lhs32 *GlobalState = globalState
		_ = _lhs32
		_271_v10 = _rhs32
		_lhs31.F3 = _rhs33
		_lhs32.F5 = _rhs34
		var _274_v12 _dafny.Array
		_ = _274_v12
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_6
		var _nw44 _dafny.Array
		_ = _nw44
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw44 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Map = (func(_275_v0 bool) func(_dafny.Int) _dafny.Map {
				return func(_276_i0 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_275_v0, _275_v0)
				}
			})(_262_v0)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw44 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw44).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw44).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_274_v12 = _nw44
		var _277_v13 _dafny.Sequence
		_ = _277_v13
		_277_v13 = _dafny.SeqOf(_274_v12)
		var _278_v14 _dafny.Map
		_ = _278_v14
		_278_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v0, _dafny.IntOfInt64(468))
		var _279_v15 _dafny.Map
		_ = _279_v15
		_279_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v13).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v0, _262_v0)).Cardinality(), _dafny.IntOfUint32((_277_v13).Cardinality()))).Uint32()).(_dafny.Array), (_278_v14).Merge(_278_v14))
		_279_v15 = (_279_v15).Update(_274_v12, _278_v14)
		var _280_i1 _dafny.Int
		_ = _280_i1
		_280_i1 = _dafny.Zero
		{
			for _262_v0 {
				{
					if (_280_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_280_i1 = (_280_i1).Plus(_dafny.One)
					(globalState).F3 = (_dafny.Zero).Minus((_this).F20())
					(globalState).F7 = _262_v0
					var _281_v16 _dafny.CodePoint
					_ = _281_v16
					_281_v16 = _dafny.CodePoint('w')
					(globalState).F5 = !(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("axeqcyxa"), _281_v16)) || (_262_v0)
					var _282_v17 _dafny.MultiSet
					_ = _282_v17
					_282_v17 = _dafny.MultiSetOf(_262_v0, true)
					(globalState).F3 = (func() _dafny.Int {
						if _262_v0 {
							return (_dafny.Zero).Minus(_dafny.IntOfUint32((_273_v11).Cardinality()))
						}
						return (func() _dafny.Int {
							if (_282_v17).Contains(true) {
								return (_282_v17).Multiplicity(true)
							}
							return (_this).F20()
						})()
					})()
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _283_v18 _dafny.MultiSet
		_ = _283_v18
		_283_v18 = _dafny.MultiSetOf((_this).F20())
		var _284_v19 _dafny.Set
		_ = _284_v19
		var _285_v20 bool
		_ = _285_v20
		var _286_v21 _dafny.Map
		_ = _286_v21
		var _out12 _dafny.Set
		_ = _out12
		var _out13 bool
		_ = _out13
		var _out14 _dafny.Map
		_ = _out14
		_out12, _out13, _out14 = Companion_Default___.M0(_283_v18, globalState)
		_284_v19 = _out12
		_285_v20 = _out13
		_286_v21 = _out14
		if _285_v20 {
			var _287_v22 _dafny.Map
			_ = _287_v22
			_287_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_285_v20, _262_v0)
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_274_v12), 0))
			_ = _index41
			(_274_v12).ArraySet1(_287_v22, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_274_v12), 0))
			_ = _index42
			(_274_v12).ArraySet1(((_287_v22).Update(_285_v20, _285_v20)).Merge(_287_v22), (_index42).Int())
			_268_v7 = _266_v6
			var _288_v23 _dafny.CodePoint
			_ = _288_v23
			_288_v23 = _dafny.CodePoint('s')
			var _289_v24 _dafny.Map
			_ = _289_v24
			_289_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v0, _288_v23)
			(globalState).F3 = (_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_289_v24, Companion_Default___.Fm2(_262_v0, (_278_v14).Cardinality(), globalState)), globalState)
			var _source5 _dafny.Map = _266_v6
			_ = _source5
			var _290___mcc_h0 _dafny.Map = _source5
			_ = _290___mcc_h0
			var _291_cf0 _dafny.Map = _290___mcc_h0
			_ = _291_cf0
			_273_v11 = _273_v11
			var _292_v25 *C0
			_ = _292_v25
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__((_this).F20())
			_292_v25 = _nw45
			var _293_v26 _dafny.Array
			_ = _293_v26
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw46
			_293_v26 = _nw46
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_293_v26), 0))
			_ = _index43
			(_293_v26).ArraySet1(_285_v20, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_293_v26), 0))
			_ = _index44
			(_293_v26).ArraySet1(!((((_292_v25).F22()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_294_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			}))).Cardinality()))).Cmp(_dafny.IntOfInt64(-390)) > 0), (_index44).Int())
			var _295_v27 _dafny.Sequence
			_ = _295_v27
			_295_v27 = _dafny.SeqOf(_292_v25, _292_v25, _292_v25)
			var _rhs35 *C0 = (_295_v27).Select((Companion_Default___.SafeIndex(((_292_v25).F22()).Minus((_292_v25).F22()), _dafny.IntOfUint32((_295_v27).Cardinality()))).Uint32()).(*C0)
			_ = _rhs35
			var _rhs36 bool = !((_this).Fm0(_dafny.UnicodeSeqOfUtf8Bytes("rdmtla"), ((_292_v25).F22()).Plus((_dafny.Zero).Minus((_292_v25).F22())), (_this).F20(), _dafny.IntOfInt64(-502), globalState))
			_ = _rhs36
			var _rhs37 bool = true
			_ = _rhs37
			var _rhs38 bool = _262_v0
			_ = _rhs38
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			_292_v25 = _rhs35
			_lhs33.F7 = _rhs36
			_lhs34.F5 = _rhs37
			_285_v20 = _rhs38
			_286_v21 = (func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(99), _dafny.IntOfInt64(-193))); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _296_v28 _dafny.Int
					_296_v28 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(99)).Cmp(_296_v28) <= 0) && ((_296_v28).Cmp(_dafny.IntOfInt64(-193)) < 0) {
						_coll16.Add((_296_v28).Plus((_this).F20()), _262_v0)
					}
				}
				return _coll16.ToMap()
			}()).Merge((_269_v9).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_269_v9).Cardinality()))).Uint32()).(_dafny.Map))
		} else {
			_285_v20 = _262_v0
			if !(_285_v20) || ((func() bool {
				if _285_v20 {
					return true
				}
				return _285_v20
			})()) {
				var _297_v29 _dafny.Array
				_ = _297_v29
				var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw47
				_297_v29 = _nw47
				var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw48
				_297_v29 = _nw48
				(globalState).F11 = !(_262_v0) || (((_this).F21()).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm5((_this).F20(), _285_v20, (_this).F20(), globalState)).Cardinality(), _dafny.IntOfUint32(((_this).F21()).Cardinality()))).Uint32()).(bool))
				_297_v29 = _297_v29
				(globalState).F3 = (_this).F20()
				var _298_v30 _dafny.CodePoint
				_ = _298_v30
				_298_v30 = _dafny.CodePoint('f')
				var _299_v31 _dafny.Map
				_ = _299_v31
				_299_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_298_v30, _285_v20)
				(globalState).F11 = (func() bool {
					if (_299_v31).Contains(_298_v30) {
						return (_299_v31).Get(_298_v30).(bool)
					}
					return (_283_v18).IsDisjointFrom(_283_v18)
				})()
			} else {
				(globalState).F3 = (_this).F20()
				var _300_v32 *C0
				_ = _300_v32
				var _nw49 *C0 = New_C0_()
				_ = _nw49
				_nw49.Ctor__((_this).F20())
				_300_v32 = _nw49
				(globalState).F14 = ((_284_v19).Intersection(_284_v19)).IsSubsetOf(_284_v19)
				var _301_v33 _dafny.CodePoint
				_ = _301_v33
				_301_v33 = _dafny.CodePoint('s')
				_301_v33 = _301_v33
				_273_v11 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_273_v11, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_273_v11).Cardinality()))).Uint32(), _301_v33), _273_v11)
			}
			(globalState).F5 = _285_v20
			var _302_v34 _dafny.Sequence
			_ = _302_v34
			_302_v34 = _dafny.SeqOf(_266_v6)
			var _rhs39 bool = _285_v20
			_ = _rhs39
			var _rhs40 bool = (_dafny.Companion_Sequence_.Contains(_302_v34, _268_v7)) && (_285_v20)
			_ = _rhs40
			var _lhs35 *GlobalState = globalState
			_ = _lhs35
			_262_v0 = _rhs39
			_lhs35.F7 = _rhs40
			(globalState).F3 = Companion_Default___.SafeModuloInt((_this).F20(), (_this).F20())
		}
		var _source6 _dafny.Map = _263_v1
		_ = _source6
		var _303___mcc_h1 _dafny.Map = _source6
		_ = _303___mcc_h1
		var _304_cf0 _dafny.Map = _303___mcc_h1
		_ = _304_cf0
		var _305_v35 *C0
		_ = _305_v35
		var _nw50 *C0 = New_C0_()
		_ = _nw50
		_nw50.Ctor__((_dafny.Zero).Minus((_this).F20()))
		_305_v35 = _nw50
		var _306_v36 _dafny.Array
		_ = _306_v36
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_7
		var _nw51 _dafny.Array
		_ = _nw51
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw51 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Int = func(_307_i3 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_307_i3, _dafny.IntOfInt64(-681))
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
		_306_v36 = _nw51
		var _308_v37 _dafny.Sequence
		_ = _308_v37
		_308_v37 = _dafny.SeqOf(_306_v36)
		var _309_v38 _dafny.Map
		_ = _309_v38
		_309_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v18, (_308_v37).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_308_v37).Cardinality()))).Uint32()).(_dafny.Array))
		var _310_v39 _dafny.Array
		_ = _310_v39
		var _nwElement0_8 _dafny.Array = _306_v36
		_ = _nwElement0_8
		var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(8))
		_ = _nw52
		(_nw52).ArraySet1(_nwElement0_8, 0)
		(_nw52).ArraySet1((func() _dafny.Array {
			if (_309_v38).Contains(_dafny.MultiSetFromSeq(_264_v3)) {
				return (_309_v38).Get(_dafny.MultiSetFromSeq(_264_v3)).(_dafny.Array)
			}
			return _306_v36
		})(), 1)
		(_nw52).ArraySet1(_306_v36, 2)
		(_nw52).ArraySet1(_306_v36, 3)
		(_nw52).ArraySet1(_306_v36, 4)
		(_nw52).ArraySet1(_306_v36, 5)
		(_nw52).ArraySet1(_306_v36, 6)
		(_nw52).ArraySet1((_308_v37).Select((Companion_Default___.SafeIndex((_305_v35).F22(), _dafny.IntOfUint32((_308_v37).Cardinality()))).Uint32()).(_dafny.Array), 7)
		_310_v39 = _nw52
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))
		_ = _index45
		(_310_v39).ArraySet1(_306_v36, (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))
		_ = _index46
		var _rhs41 _dafny.Array = _306_v36
		_ = _rhs41
		var _rhs42 _dafny.Map = _268_v7
		_ = _rhs42
		var _rhs43 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_264_v3).Cardinality()))
		_ = _rhs43
		var _lhs36 _dafny.Array = _310_v39
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))
		_ = _lhs37
		var _lhs38 *GlobalState = globalState
		_ = _lhs38
		(_lhs36).ArraySet1(_rhs41, (_lhs37).Int())
		_266_v6 = _rhs42
		_lhs38.F3 = _rhs43
		var _311_v40 _dafny.Array
		_ = _311_v40
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
		_ = _nw53
		_311_v40 = _nw53
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_311_v40), 0))
		_ = _index47
		(_311_v40).ArraySet1(_266_v6, (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_311_v40), 0))
		_ = _index48
		(_311_v40).ArraySet1(_266_v6, (_index48).Int())
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_310_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))).Int()))
		_ = _arr0
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_dafny.ArrayCastTo((_310_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))).Int()))), 0))
		_ = _index49
		(_arr0).ArraySet1((_this).F20(), (_index49).Int())
		var _312_v41 _dafny.Set
		_ = _312_v41
		_312_v41 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_262_v0)).Cardinality()))
		var _313_v42 _dafny.Map
		_ = _313_v42
		_313_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v6, _312_v41)
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_310_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))).Int()))
		_ = _arr1
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_dafny.ArrayCastTo((_310_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))).Int()))), 0))
		_ = _index50
		var _rhs44 _dafny.Int = (_this).F20()
		_ = _rhs44
		var _rhs45 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_this).F20()).Times((_this).F20()), (_305_v35).F22())))
		_ = _rhs45
		var _rhs46 _dafny.Map = _313_v42
		_ = _rhs46
		var _rhs47 _dafny.Set = _312_v41
		_ = _rhs47
		var _lhs39 *GlobalState = globalState
		_ = _lhs39
		var _lhs40 _dafny.Array = _dafny.ArrayCastTo((_310_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))).Int()))
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_dafny.ArrayCastTo((_310_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_310_v39), 0))).Int()))), 0))
		_ = _lhs41
		_lhs39.F3 = _rhs44
		(_lhs40).ArraySet1(_rhs45, (_lhs41).Int())
		_313_v42 = _rhs46
		_312_v41 = _rhs47
	}
}
func (_this *C1) F20() _dafny.Int {
	{
		return _this._f20
	}
}
func (_this *C1) F21() _dafny.Sequence {
	{
		return _this._f21
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
